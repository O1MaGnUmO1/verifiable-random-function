package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

func NewClient(RPCUrl string) (*Client, error) {
	client, err := ethclient.Dial(RPCUrl)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return nil, err
	}
	return &Client{
		EthClient: client,
	}, nil
}
