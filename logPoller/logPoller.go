package logpoller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
	"vrf/client"
	"vrf/coordinator"
	"vrf/keys/vrfkey"
	"vrf/keys/vrfkey/proof"
	vrf_service "vrf/services/vrf"
	wallet_service "vrf/services/wallet"
	"vrf/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

// LogPoller represents an Ethereum log poller
type LogPoller struct {
	vrf_service     vrf_service.VRFService
	wallet_service  *wallet_service.WalletService
	client          *client.Client
	coord           coordinator.Coordinator
	contractAddress common.Address
	eventSignature  []common.Hash
	lastBlockNumber *big.Int
	unfulfilled     map[string]*coordinator.CoordinatorRandomWordsRequested
	processed       map[string]bool
	replayFromBlock *big.Int // Block number from which to start replaying logs
	logger          *logrus.Logger
	subscribed      bool
	sub             ethereum.Subscription
	m               sync.Mutex
	headers         chan *types.Header
}

type LogState struct {
	LastProcessedBlock *big.Int `json:"last_processed_block"`
}

// NewLogPoller creates a new instance of LogPoller
func NewLogPoller(client *client.Client, contractAddress common.Address, replayFromBlock *big.Int, wallet_service wallet_service.WalletService, vrf_service vrf_service.VRFService) (*LogPoller, error) {
	parsedABI, err := abi.JSON(strings.NewReader(coordinator.CoordinatorABI))
	if err != nil {
		return nil, err
	}
	coord, err := coordinator.NewCoordinator(contractAddress, client.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create coordinator %v", err)
	}

	eventSignature := parsedABI.Events["RandomWordsRequested"].ID
	if eventSignature == (common.Hash{}) {
		return nil, fmt.Errorf("event 'CoordinatorRandomWordsRequested' not found in contract ABI")
	}
	sig1 := parsedABI.Events["RandomWordsFulfilled"].ID
	if eventSignature == (common.Hash{}) {
		return nil, fmt.Errorf("event 'CoordinatorRandomWordsRequested' not found in contract ABI")
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})

	return &LogPoller{
		client:          client,
		contractAddress: contractAddress,
		eventSignature:  []common.Hash{eventSignature, sig1},
		replayFromBlock: replayFromBlock,
		unfulfilled:     make(map[string]*coordinator.CoordinatorRandomWordsRequested),
		logger:          logger,
		coord:           *coord,
		processed:       make(map[string]bool),
		wallet_service:  &wallet_service,
		vrf_service:     vrf_service,
		headers:         make(chan *types.Header),
	}, nil
}

// Start starts the log polling process
func (lp *LogPoller) Start(ctx context.Context) error {
	// go lp.wallet_service.MonitorBalanceChanges()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		// case newBalance := <-lp.wallet_service.BalanceCh:
		// 	logrus.WithFields(logrus.Fields{
		// 		"Previous Balance": wallet_service.WeiToETH((newBalance.OldBalance)) + "FTN",
		// 		"Current Balance":  wallet_service.WeiToETH(newBalance.NewBalance) + "FTN",
		// 		"Amount":           wallet_service.WeiToETH(newBalance.NewBalance.Sub(newBalance.NewBalance, newBalance.OldBalance)) + "FTN",
		// 	}).Info("Got Deposit")
		default:
			err := lp.pollLogs()
			if err != nil {
				log.Printf("Error polling logs: %v", err)
				// backoffDuration := time.Duration(retries) * time.Second
				// log.Printf("Retrying after %v", backoffDuration)
				time.Sleep(1 * time.Second)
			} else {
				// lp.logger.WithFields(logrus.Fields{
				// 	"Timestamp": time.Now().UTC(),
				// 	"Message":   "No New Logs to Read, Restarting ...",
				// }).Info("No New Logs")
				continue // Reset retry count on successful poll
			}
		}
	}
}

// pollLogs polls Ethereum logs
func (lp *LogPoller) pollLogs() error {
	var fromBlock *big.Int

	// Determine the block to start polling from
	if lp.lastBlockNumber == nil || lp.replayFromBlock != nil {
		if lp.replayFromBlock != nil {
			lp.replayFromBlock = nil // Reset replayFromBlock after using it
		} else {
			latestBlockNumber, err := lp.client.GetLatestBlockNumber()
			if err != nil {
				return err
			}
			fromBlock = lp.getLastProcessedBlock()
			if fromBlock.Cmp(new(big.Int).SetUint64(latestBlockNumber)) >= 0 {
				// Last processed block is already equal to or ahead of the latest block,
				// no need to poll logs
				return nil
			}
		}
	}

	latestBlockNumber, err := lp.client.GetLatestBlockNumber()
	if err != nil {
		return err
	}

	toBlock := latestBlockNumber

	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignature},
		FromBlock: new(big.Int).SetUint64(latestBlockNumber - uint64(200)),
		ToBlock:   new(big.Int).SetUint64(toBlock),
	}

	logs, err := lp.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
	for _, log := range logs {
		req, err := lp.coord.ParseRandomWordsRequested(log)
		if err != nil {
			// fmt.Printf("failed to parse request %v", err)
			continue
		}

		key := lp.vrf_service.GetVrfKey()
		pkh := key.PublicKey.MustHash()
		if !bytes.Equal(req.KeyHash[:], pkh[:]) {
			logrus.WithFields(logrus.Fields{
				"expected Key Hash": pkh.Hex(),
				"Actual":            common.BytesToHash(req.KeyHash[:]),
				"Request ID":        req.RequestId,
			}).Info()
			continue
		}

		latestBlockNumber, err = lp.client.GetLatestBlockNumber()
		if err != nil {
			return fmt.Errorf("failed to load last block %v", err)
		}
		com, err := lp.coord.GetCommitment(nil, req.RequestId)
		if err != nil {
			return fmt.Errorf("failed to get commitment with request ID : %s Err: %v", common.BigToHash(req.RequestId).Hex(), err)
		}
		if utils.IsAllZeros(com) {
			// lp.logger.WithFields(logrus.Fields{
			// 	"Request ID": common.BigToHash(req.RequestId).Hex(),
			// }).Info("Received already fulfilled request")
			// fmt.Println()
			continue
		}
		lp.m.Lock()
		// Check if the request is already in unfulfilled map
		if !lp.processed[req.RequestId.String()] {
			if _, exists := lp.unfulfilled[req.RequestId.String()]; !exists {
				lp.unfulfilled[req.RequestId.String()] = req
			}
			lp.m.Unlock()
		}
	}

	// Start WaitForConfirmations in a separate goroutine
	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := lp.WaitForConfirmations(done); err != nil {
			log.Printf("Error in WaitForConfirmations: %v", err)
		}
	}()

	if err := lp.setLastProcessedBlock(new(big.Int).SetUint64(latestBlockNumber)); err != nil {
		fmt.Printf("failed to set last processed block %v", err)
		return err
	}

	return nil
}

func (lp *LogPoller) WaitForConfirmations(done chan struct{}) error {
	if !lp.subscribed {
		subs, err := lp.client.EthClient.SubscribeNewHead(context.Background(), lp.headers)
		if err != nil {
			return err
		}
		lp.sub = subs
		lp.subscribed = true
	}

	for {
		select {
		case <-lp.sub.Err():
			return fmt.Errorf("failed to subscribe")
		case header := <-lp.headers:
			lp.logger.WithFields(logrus.Fields{
				"Timestamp":   time.Now().UTC(),
				"Head Number": header.Number,
			}).Infoln("Received new head")

			if len(lp.unfulfilled) == 0 {
				lp.logger.WithFields(logrus.Fields{
					"Timestamp": time.Now().UTC(),
				}).Infoln("No pending VRF requests")
				return nil
			}

			// Process each request in a separate goroutine
			for id, req := range lp.unfulfilled {
				lp.processRequest(id, req, header.Number.Uint64())
			}
		}
	}
}

func (lp *LogPoller) processRequest(id string, req *coordinator.CoordinatorRandomWordsRequested, latestBlockNumber uint64) {
	lp.logger.WithFields(logrus.Fields{
		"Timestamp":          time.Now().UTC(),
		"KeyHash":            common.BytesToHash(req.KeyHash[:]),
		"Request ID":         common.BytesToHash(req.RequestId.Bytes()),
		"Callback Gas Limit": req.CallbackGasLimit,
		"Min Incoming Confs": req.MinimumRequestConfirmations,
		"Num Words":          req.NumWords,
		"Sub ID":             req.SubId,
		"Sender":             req.Sender.Hex(),
	}).Info("Random Words Requested")

	lp.m.Lock()
	defer lp.m.Unlock()

	if _, ok := lp.unfulfilled[id]; ok {
		delete(lp.unfulfilled, id)
		go func(id string, req *coordinator.CoordinatorRandomWordsRequested) {
			if latestBlockNumber-req.Raw.BlockNumber >= uint64(req.MinimumRequestConfirmations) {
				if err := lp.decodeLog(req, lp.vrf_service.GetVrfKey()); err != nil {
					lp.logger.WithError(err).Error("Failed to decode log")
				}
				lp.processed[id] = true
				return
			} else {
				fmt.Printf("Waiting for %d confirmations. Current block: %d\n Request block: %d\n", req.MinimumRequestConfirmations, latestBlockNumber, req.Raw.BlockNumber)
			}
		}(id, req)
	}
}

func (lp *LogPoller) decodeLog(req *coordinator.CoordinatorRandomWordsRequested, key vrfkey.KeyV2) error {
	preSeed, err := proof.BigToSeed(req.PreSeed)
	if err != nil {
		return fmt.Errorf("failed to get preseed %v", err)
	}
	preSeedData := proof.PreSeedDataV2{
		PreSeed:          preSeed,
		BlockHash:        req.Raw.BlockHash,
		BlockNum:         req.Raw.BlockNumber,
		SubId:            req.SubId,
		CallbackGasLimit: req.CallbackGasLimit,
		NumWords:         req.NumWords,
		Sender:           req.Sender,
	}

	finalSeed := proof.FinalSeedV2(preSeedData)
	p, err := key.GenerateProof(finalSeed)
	if err != nil {
		return fmt.Errorf("failed to generate proof %v", err)
	}
	onChainProof, rc, err := proof.GenerateProofResponseFromProofV2(p, preSeedData)
	if err != nil {
		return fmt.Errorf("failed to generate proof responce %v", err)
	}
	ok, err := p.VerifyVRFProof()
	if !ok {
		panic("failed to verify proof")
	}
	if err != nil {
		panic(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(coordinator.CoordinatorABI))
	if err != nil {
		return fmt.Errorf("failed to parse abi %v", err)
	}

	b, err := parsedABI.Pack("fulfillRandomWords", onChainProof, rc)
	if err != nil {
		return fmt.Errorf("failed to pack abi %v", err)
	}
	msg := ethereum.CallMsg{
		From: lp.wallet_service.Key.Address,
		To:   &lp.contractAddress,
		Data: b,
	}

	gas, err := lp.client.EthClient.EstimateGas(context.Background(), msg)
	if err != nil {
		fmt.Printf("failed to estimate gas %v", err)
		return err
	}
	lp.logger.WithFields(logrus.Fields{
		"From":      lp.wallet_service.Key.Address,
		"To":        lp.contractAddress,
		"Gas Limit": gas,
	}).Info("Estimate Gas")
	nonce, err := lp.client.EthClient.PendingNonceAt(context.Background(), lp.wallet_service.Key.Address)
	if err != nil {
		return fmt.Errorf("failed to get nonce %v", err)
	}

	chainId, _ := new(big.Int).SetString("4090", 10)
	gasPrice, err := lp.client.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price %v", err)
	}
	tx := types.NewTransaction(nonce, lp.contractAddress, big.NewInt(0), gas, gasPrice, b)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), lp.wallet_service.Key.ToEcdsaPrivKey())
	if err != nil {
		return fmt.Errorf("failed to sign tx %v", err)
	}
	if err := lp.client.EthClient.SendTransaction(context.Background(), signedTx); err != nil {
		return fmt.Errorf("faield to send tx %v", err)
	}

	lp.logger.WithFields(logrus.Fields{
		"From":  lp.wallet_service.Key.Address,
		"To":    lp.contractAddress,
		"Nonce": nonce,
		"Data":  tx.Data(),
	}).Info("Sending transaction")

	// receipt, err := bind.WaitMined(context.Background(), lp.client.EthClient, tx)
	// if err != nil {
	// 	return fmt.Errorf("failed to mine transaction %v", err)
	// }
	// lp.logger.WithFields(logrus.Fields{
	// 	"Receipt": receipt,
	// }).Info("Transaction sent")

	return nil
}

// getLastProcessedBlock returns the last processed block number
func (lp *LogPoller) getLastProcessedBlock() *big.Int {
	// Load the last processed block number from a file
	state := &LogState{}
	data, err := ioutil.ReadFile("log_state.json")
	if err != nil {
		return nil // Error reading file, return nil
	}
	err = json.Unmarshal(data, state)
	if err != nil {
		return nil // Error unmarshalling JSON, return nil
	}
	return state.LastProcessedBlock
}

// setLastProcessedBlock sets the last processed block number
func (lp *LogPoller) setLastProcessedBlock(blockNumber *big.Int) error {
	// Serialize the block number to JSON
	state := &LogState{
		LastProcessedBlock: blockNumber,
	}
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	// Write the serialized JSON to a file
	err = ioutil.WriteFile("log_state.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
