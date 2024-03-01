package headtracker

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
	"vrf/client"
	logpoller "vrf/logPoller"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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

// Start starts the head tracking process with enhanced error handling
func (ht *HeadTracker) Start(ctx context.Context) error {
	retryInterval := time.Second * 5 // Retry every 5 seconds
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := ht.subscribeToNewHead(ctx); err != nil {
				ht.logger.Errorf("Failed to subscribe to new heads: %v", err)
				// Check if the error is a connectivity issue.
				if err == rpc.ErrClientQuit || err == rpc.ErrSubscriptionQueueOverflow {
					// These errors indicate a problem with the Ethereum client or the subscription itself.
					ht.logger.Infof("Retrying subscription due to Ethereum client error in %v...", retryInterval)
					time.Sleep(retryInterval)
					continue // Retry the subscription
				} else {
					// For other types of errors, consider logging and handling them as needed.
					return fmt.Errorf("failed to subscribe to new heads with non-retryable error: %v", err)
				}
			}
			// Successfully subscribed
			return nil
		}
	}
}

// subscribeToNewHead attempts to subscribe to new block headers and process them
func (ht *HeadTracker) subscribeToNewHead(ctx context.Context) error {
	subs, err := ht.client.EthClient.SubscribeNewHead(ctx, ht.headers)
	if err != nil {
		return err // Return the error to be handled by the caller
	}

	ht.logger.Infoln("Successfully subscribed to new head")
	ht.sub = subs // Store the subscription for later use

	// Process incoming block headers

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				ht.sub.Unsubscribe() // Ensure the subscription is terminated
				return
			case err := <-ht.sub.Err():
				ht.logger.Errorf("Subscription error: %v", err)
				return // Exit the goroutine if the subscription faces an error
			case header := <-ht.headers:
				ht.processNewHead(header) // Process each new head
			}
		}
	}()
	wg.Wait()
	return nil
}

// processNewHead processes a new block header
func (ht *HeadTracker) processNewHead(header *types.Header) {
	ht.logger.WithFields(logrus.Fields{
		"Timestamp":   time.Now().UTC(),
		"Head Number": header.Number,
	}).Infoln("Received new head")
	ht.logPoller.SetLatestBlockNumber(header.Number.Uint64())

	// Process requests within this head using LogPoller
	if err := ht.logPoller.PollLogs(); err != nil {
		ht.logger.Errorf("Error processing requests: %v", err)
	}
}
