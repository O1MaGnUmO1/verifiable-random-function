package headtracker

import (
	"context"
	"fmt"
	"os"
	"time"
	"vrf/client"
	logpoller "vrf/logPoller"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

// CustomFormatter is a Logrus formatter that adds two newline characters to log entries.
type CustomFormatter struct {
	logrus.JSONFormatter
}

// Format formats the log entry and adds two newline characters to the end.
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

// HeadTracker tracks new heads on the Ethereum blockchain
type HeadTracker struct {
	client    *client.Client
	headers   chan *types.Header
	logger    *logrus.Logger
	sub       ethereum.Subscription
	logPoller *logpoller.LogPoller // Reference to LogPoller
}

// NewHeadTracker creates a new instance of HeadTracker
func NewHeadTracker(client *client.Client, logPoller *logpoller.LogPoller) *HeadTracker {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})
	return &HeadTracker{
		client:    client,
		headers:   make(chan *types.Header),
		logger:    logger,
		logPoller: logPoller, // Set the reference to LogPoller
	}
}

// Start starts the head tracking process
func (ht *HeadTracker) Start(ctx context.Context) error {
	var subs ethereum.Subscription
	var err error
	retryInterval := time.Second * 5 // Retry every 5 seconds
	maxRetries := 5                  // Maximum number of retries before failing
	retries := 0

	for {
		// Attempt to subscribe to new heads
		subs, err = ht.client.EthClient.SubscribeNewHead(ctx, ht.headers)
		if err != nil {
			ht.logger.Errorf("Failed to subscribe to heads on blockchain: %v", err)
			retries++
			if retries >= maxRetries {
				return fmt.Errorf("exceeded max retries for head subscription: %v", err)
			}
			ht.logger.Infof("Retrying subscription in %v...", retryInterval)
			select {
			case <-time.After(retryInterval):
				// Retry after the interval
				continue
			case <-ctx.Done():
				// Context cancelled, stop retrying
				return ctx.Err()
			}
		}
		break // Exit the loop if subscription is successful
	}

	ht.logger.Infoln("Successfully subscribed to new head")
	ht.sub = subs
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case header := <-ht.headers:
			ht.logger.WithFields(logrus.Fields{
				"Timestamp":   time.Now().UTC(),
				"Head Number": header.Number,
			}).Infoln("Received new head")
			ht.logPoller.SetLatestBlockNumber(header.Number.Uint64())
			// Process requests within this head using LogPoller
			go func() {
				err := ht.logPoller.PollLogs()
				if err != nil {
					ht.logger.Errorf("Error processing requests: %v", err)
				}
			}()
		}
	}
}
