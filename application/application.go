package application

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"vrf/client"
	"vrf/headtracker"
	logpoller "vrf/logPoller"
	vrf_service "vrf/services/vrf"
	wallet_service "vrf/services/wallet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var replayFromBlock = uint64(1000000) // default number of replay blocks
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

type Application struct {
	Client        *client.Client
	LogPoller     *logpoller.LogPoller
	WalletService *wallet_service.WalletService
	VRFService    *vrf_service.VRFService
	HeadTracker   *headtracker.HeadTracker
	Logger        *logrus.Logger
}

func NewApplication() (*Application, error) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})
	client, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create application : Err=<%v>", err)
	}
	replayFromBlock := replayFromBlock
	if os.Getenv("EC_REPLAY_FROM_BLOCK") != "" {
		replayFromBlock, err = strconv.ParseUint(os.Getenv("EC_REPLAY_FROM_BLOCK"), 10, 64)
		if err != nil {
			logrus.Errorf("failed to get replay block number will use deafult value %v", err)
		}
	}
	walletService := wallet_service.NewWalletService(client)
	vrfService := vrf_service.NewVRFService()
	logpoller, err := logpoller.NewLogPoller(client, common.HexToAddress(os.Getenv("EC_COORDINATOR_ADDRESS")), replayFromBlock, walletService, vrfService)
	if err != nil {
		return nil, fmt.Errorf("failed to create log poller : Err=<%v>", err)
	}
	headtracker := headtracker.NewHeadTracker(client, logpoller)
	return &Application{
		Client:        client,
		LogPoller:     logpoller,
		WalletService: walletService,
		VRFService:    vrfService,
		HeadTracker:   headtracker,
		Logger:        logger,
	}, nil
}

func (app *Application) Run() {
	// Connect to the Ethereum node
	client, err := client.NewClient()
	if err != nil {
		app.Logger.Errorf("failed to dial %v", err)
	}
	_, err = app.WalletService.CreateNewFTNKey()
	if err != nil {
		app.Logger.Errorf("failed to create FTN Key %v", err)
	}
	app.WalletService.PrintWalletDetails()
	_, err = app.VRFService.GenerateNewVRFKey()
	if err != nil {
		app.Logger.Errorf("failed to generate VRF key %v", err)
	}
	app.VRFService.PrintVrfKey()
	ht := headtracker.NewHeadTracker(client, app.LogPoller)

	err = ht.Start(context.Background())
	if err != nil {
		fmt.Printf("error starting head tracker %v", err)
	}
}
