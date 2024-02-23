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
	"time"
	"vrf/client"
	"vrf/coordinator"
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
	replayFromBlock *big.Int // Block number from which to start replaying logs
	logger          *logrus.Logger
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
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &LogPoller{
		client:          client,
		contractAddress: contractAddress,
		eventSignature:  []common.Hash{eventSignature, sig1},
		replayFromBlock: replayFromBlock,
		unfulfilled:     make(map[string]*coordinator.CoordinatorRandomWordsRequested),
		logger:          logger,
		coord:           *coord,
		wallet_service:  &wallet_service,
		vrf_service:     vrf_service,
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
			fromBlock = lp.replayFromBlock
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
	} else {
		fromBlock = new(big.Int).Add(lp.lastBlockNumber, big.NewInt(1))
	}

	latestBlockNumber, err := lp.client.GetLatestBlockNumber()
	if err != nil {
		return err
	}

	toBlock := latestBlockNumber

	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignature},
		FromBlock: fromBlock,
		ToBlock:   new(big.Int).SetUint64(toBlock),
	}

	logs, err := lp.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	for _, log := range logs {
		err := lp.decodeLog(log)
		if err != nil {
			fmt.Errorf("failed to decode log %v", err)
		}
	}

	if err := lp.setLastProcessedBlock(new(big.Int).SetUint64(latestBlockNumber)); err != nil {
		fmt.Printf("failed to set last processed block %v", err)
		return err
	}

	return nil
}

func (lp *LogPoller) decodeLog(log types.Log) error {
	req, err := lp.coord.ParseRandomWordsRequested(log)
	if err != nil {
		return fmt.Errorf("failed to parse request %v", err)
	}

	fmt.Printf("Block hash: %s\n", log.BlockHash.Hex())
	fmt.Printf("Block hash: %x\n", log.BlockNumber)

	key := lp.vrf_service.GetVrfKey()
	pkh := key.PublicKey.MustHash()
	if !bytes.Equal(req.KeyHash[:], pkh[:]) {
		logrus.WithFields(logrus.Fields{
			"expected Key Hash": pkh.Hex(),
			"Actual":            common.BytesToHash(req.KeyHash[:]),
			"Request ID":        req.RequestId,
		}).Info()
		return fmt.Errorf("key hashes are not equal")
	}

	if lp.unfulfilled[req.RequestId.String()] != nil {
		return fmt.Errorf("already fulfilled")
	}
	com, err := lp.coord.GetCommitment(nil, req.RequestId)
	if err != nil {
		return fmt.Errorf("failed to get commitment with request ID : %s Err: %v", common.BigToHash(req.RequestId).Hex(), err)
	}
	if utils.IsAllZeros(com) {
		lp.logger.WithFields(logrus.Fields{
			"Request ID": common.BigToHash(req.RequestId).Hex(),
		}).Info("Received already fulfilled request\n\n")
		fmt.Println()
		return err
	} else {
		lp.unfulfilled[req.RequestId.String()] = req
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
		fmt.Println()

		preSeed, err := proof.BigToSeed(req.PreSeed)
		if err != nil {
			return fmt.Errorf("failed to get preseed %v", err)
		}
		preSeedData := proof.PreSeedDataV2{
			PreSeed:          preSeed,
			BlockHash:        log.BlockHash,
			BlockNum:         log.BlockNumber,
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

		// parsedABI, err := abi.JSON(strings.NewReader(coordinator.CoordinatorABI))
		// if err != nil {
		// 	return fmt.Errorf("failed to parse abi %v", err)
		// }

		// b, err := parsedABI.Pack("fulfillRandomWords", onChainProof, rc)
		nonce, err := lp.client.EthClient.PendingNonceAt(context.Background(), lp.wallet_service.Key.Address)
		if err != nil {
			return fmt.Errorf("failed to get nonce %v", err)
		}
		// chainId, _ := new(big.Int).SetString("4090", 10)
		transactOPtx := bind.NewKeyedTransactor(lp.wallet_service.Key.ToEcdsaPrivKey())
		transactOPtx.Nonce = new(big.Int).SetUint64(nonce)
		// signerFn := InitializeSignerFn(chainID)
		// transactOPtx.Signer = signerFn
		tx, err := lp.coord.FulfillRandomWords(transactOPtx, onChainProof, rc)
		if err != nil {
			fmt.Printf("Fulfill error is %v", err)
			return fmt.Errorf("failed to fulfill random %v", err)
		}
		fmt.Printf("tx hash %s", tx.Hash())
		// chainId, _ := new(big.Int).SetString("4090", 10)
		// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), lp.wallet_service.Key.ToEcdsaPrivKey())
		// if err != nil {
		// 	return fmt.Errorf("failed to sign tx %v", err)
		// }
		// err = lp.client.EthClient.SendTransaction(context.Background(), signedTx)
		// if err != nil {
		// 	return err
		// }
		// logrus.Printf("Transaction sent: %s", signedTx.Hash().Hex())
	}

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
