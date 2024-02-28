package logpoller

import (
	"bytes"
	"context"
	"fmt"
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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	LastBlockNumber uint64
	Unfulfilled     map[string]*coordinator.CoordinatorRandomWordsRequested
	Processed       map[string]bool
	rw              sync.RWMutex
	wg              sync.WaitGroup
	logger          *logrus.Logger
	Mu              sync.Mutex
	nonceMap        map[common.Address]uint64
	nonceMut        sync.Mutex
}

type LogState struct {
	LastProcessedBlock *big.Int `json:"last_processed_block"`
}

func (lp *LogPoller) getFromCache(key string) (*coordinator.CoordinatorRandomWordsRequested, bool) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	value, ok := lp.Unfulfilled[key]
	return value, ok
}

func (lp *LogPoller) addToCache(key string, value *coordinator.CoordinatorRandomWordsRequested) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	lp.Unfulfilled[key] = value
}

func (lp *LogPoller) deleteFromCache(key string) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	_, ok := lp.Unfulfilled[key]
	if ok {
		delete(lp.Unfulfilled, key)
	}
}

func (lp *LogPoller) setRequestProcessed(key string) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	lp.Processed[key] = true
}

func (lp *LogPoller) isProcessed(key string) bool {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	if processed, ok := lp.Processed[key]; ok {
		if processed {
			return true
		}
	}
	return false
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
		Unfulfilled:     make(map[string]*coordinator.CoordinatorRandomWordsRequested),
		logger:          logger,
		coord:           *coord,
		Processed:       make(map[string]bool),
		wallet_service:  &wallet_service,
		vrf_service:     vrf_service,
		Mu:              sync.Mutex{},
		rw:              sync.RWMutex{},
		wg:              sync.WaitGroup{},
		nonceMap:        make(map[common.Address]uint64),
		nonceMut:        sync.Mutex{},
	}, nil
}

// pollLogs polls Ethereum logs
func (lp *LogPoller) PollLogs() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignature},
		FromBlock: new(big.Int).SetUint64(lp.GetLatestBlockNumber() - uint64(200)),
		ToBlock:   new(big.Int).SetUint64(lp.GetLatestBlockNumber()),
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
		// Check if the request is already in unfulfilled map
		if !lp.isProcessed(req.RequestId.String()) {
			if _, ok := lp.getFromCache(req.RequestId.String()); !ok {
				lp.addToCache(req.RequestId.String(), req)
			}
		}
	}
	lp.WaitForConfirmations()
	return nil
}

func (lp *LogPoller) WaitForConfirmations() {
	lp.Mu.Lock()
	requests := make(map[string]*coordinator.CoordinatorRandomWordsRequested, len(lp.Unfulfilled))
	for id, req := range lp.Unfulfilled {
		requests[id] = req
	}
	lp.Mu.Unlock()
	for id, req := range requests {
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
		// Increment WaitGroup counter
		lp.wg.Add(1)
		// Start a goroutine for each request
		go func(id string, req *coordinator.CoordinatorRandomWordsRequested) {
			defer lp.wg.Done() // Decrement WaitGroup counter when goroutine finishes
			lp.ProcessRequest(id, req, lp.GetLatestBlockNumber())
		}(id, req)
	}

	// Wait for all goroutines to finish
	lp.wg.Wait()
}

func (lp *LogPoller) ProcessRequest(id string, req *coordinator.CoordinatorRandomWordsRequested, latestBlockNumber uint64) {
	if lp.isProcessed(id) {
		return
	}
	lp.deleteFromCache(id)
	if latestBlockNumber-req.Raw.BlockNumber >= uint64(req.MinimumRequestConfirmations) {
		if err := lp.decodeLog(req, lp.vrf_service.GetVrfKey()); err != nil {
			lp.logger.Errorf("Failed to decode log %v", err)
			lp.setRequestProcessed(id)
			return
		}
		lp.logger.WithFields(logrus.Fields{
			"min incoming confirmations": req.MinimumRequestConfirmations,
			"requestId":                  common.BytesToHash(req.RequestId.Bytes()),
		}).Info("finished handling request")
		lp.setRequestProcessed(id)

		return
	} // else {
	// 	// fmt.Printf("Waiting for %d confirmations. Current block: %d\n Request block: %d\n", req.MinimumRequestConfirmations, latestBlockNumber, req.Raw.BlockNumber)
	// }

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
	// msg := ethereum.CallMsg{
	// 	From: lp.wallet_service.Key.Address,
	// 	To:   &lp.contractAddress,
	// 	Data: b,
	// }

	// gas, err := lp.client.EthClient.EstimateGas(context.Background(), msg)
	// if err != nil {
	// 	fmt.Printf("failed to estimate gas %v", err)
	// 	return err
	// }

	// increased := float64(gas) * 1.1

	gas := uint64(1000000)

	lp.logger.WithFields(logrus.Fields{
		"From":      lp.wallet_service.Key.Address,
		"To":        lp.contractAddress,
		"Gas Limit": gas,
		// "Increased by 1.1": uint64(increased),
	}).Info("Estimate Gas")

	nonce, err := lp.GetNonce(lp.wallet_service.Key.Address)
	if err != nil {
		return fmt.Errorf("failed to get nonce %v", err)
	}

	chainId, _ := new(big.Int).SetString("4090", 10)
	// gasPrice, err := lp.client.EthClient.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return fmt.Errorf("failed to get gas price %v", err)
	// }

	gasPrice := big.NewInt(100)
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
	}).Infof("Sending transaction to request ID %s", common.BytesToHash(req.RequestId.Bytes()))

	receipt, err := bind.WaitMined(context.Background(), lp.client.EthClient, signedTx)
	if err != nil {
		logrus.Errorf("failed to mine tx %v", err)
		return err
	}
	lp.logger.WithFields(logrus.Fields{
		"Receipt": receipt,
	}).Info("Transaction sent")

	return nil
}

// getLastProcessedBlock returns the last processed block number
func (lp *LogPoller) GetLatestBlockNumber() uint64 {
	return lp.LastBlockNumber
}

// setLastProcessedBlock sets the last processed block number
func (lp *LogPoller) SetLatestBlockNumber(blockNumber uint64) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	lp.LastBlockNumber = blockNumber
}

// GetNonce fetches the next available nonce for the given account.
func (lp *LogPoller) GetNonce(account common.Address) (uint64, error) {
	lp.nonceMut.Lock()
	defer lp.nonceMut.Unlock()

	// Check if we have the nonce stored for the account
	if nonce, ok := lp.nonceMap[account]; ok {
		// Increment and update the nonce for the account
		lp.nonceMap[account] = nonce + 1
		return nonce, nil
	}

	// Fetch the current nonce from the network
	currentNonce, err := lp.client.EthClient.PendingNonceAt(context.Background(), account)
	if err != nil {
		return 0, err
	}

	// Store the fetched nonce
	lp.nonceMap[account] = currentNonce + 1

	return currentNonce, nil
}
