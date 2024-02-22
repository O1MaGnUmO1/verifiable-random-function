package main

import (
	"context"
	"fmt"
	"math/big"
	"vrf/client"
	logpoller "vrf/logPoller"

	"github.com/ethereum/go-ethereum/common"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	// Connect to the Ethereum node
	client, err := client.NewClient("https://rpc1.oasis.bahamutchain.com")
	if err != nil {
		fmt.Printf("failed to dial %v", err)
	}
	// Define the contract address
	contractAddress := common.HexToAddress("0x1f16092d5E87bE974454b3601B47f1F60d6C0C6F")

	lp, err := logpoller.NewLogPoller(client, contractAddress, big.NewInt(1000000))
	if err != nil {
		fmt.Printf("failed to create log poller %v", err)
	}
	fmt.Println("Starting log poller...")
	if err := lp.Start(context.Background()); err != nil {
		fmt.Errorf("failed to start log poller %v", err)
	}
}
