package application

// type Application struct {
// 	Client        *client.Client
// 	LogPoller     *logpoller.LogPoller
// 	WalletService *wallet_sevice.WalletService
// 	VRFService    *vrf_service.VRFService
// }

// func NewApplication() (*Application, error) {
// 	client, err := client.NewClient(os.Getenv("ETH_NODE_URL"))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create application : Err=<%v>", err)
// 	}
// 	replayFromBlock := uint64(1000000)
// 	if os.Getenv("EC_REPLAY_FROM_BLOCK") != "" {
// 		replayFromBlock, err = strconv.ParseUint(os.Getenv("EC_REPLAY_FROM_BLOCK"), 10, 64)
// 		if err != nil {
// 			fmt.Errorf("failed to get replay block number will use deafult value %v", err)
// 		}
// 	}
// 	logpoller, err := logpoller.NewLogPoller(client, common.HexToAddress(os.Getenv("EC_COORDINATOR_ADDRESS")), new(big.Int).SetUint64(replayFromBlock),wa)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create log poller : Err=<%v>", err)
// 	}
// 	return &Application{
// 		Client:        client,
// 		LogPoller:     logpoller,
// 		WalletService: nil,
// 		VRFService:    nil,
// 	}, nil
// }
