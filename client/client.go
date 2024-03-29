package client

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Client struct {
	EthClient *ethclient.Client
}

func (cl *Client) GetLatestBlockNumber() (uint64, error) {
	return cl.EthClient.BlockNumber(context.Background())
}

func (cl *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return cl.EthClient.FilterLogs(ctx, query)
}

func (cl *Client) BalanceAt(ctx context.Context, address common.Address, blockNum *big.Int) (*big.Int, error) {
	return cl.EthClient.BalanceAt(ctx, address, blockNum)
}

func NewClient() (*Client, error) {
	client, err := ethclient.Dial(os.Getenv("EC_NODE_WSURL"))
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return nil, err
	}
	return &Client{
		EthClient: client,
	}, nil
}

func (cl *Client) TryReconnect() error {
	for {
		cl.EthClient.Close()
		client, err := ethclient.Dial(os.Getenv("EC_NODE_WSURL"))
		if err == nil {
			fmt.Printf("Successfully connected to Ethereum node")
			cl.EthClient = client
			return nil
		}
		logrus.Errorf("Failed to connect to Ethereum node, Retrying ...")
	}
}
