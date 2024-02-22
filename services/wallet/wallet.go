package wallet_sevice

import (
	"context"
	"fmt"
	"math/big"
	"time"
	"vrf/client"
	"vrf/keys/ethkey"

	"github.com/sirupsen/logrus"
)

type WalletService struct {
	Client client.Client
	key    ethkey.KeyV2
}

func NewWalletService(client *client.Client) *WalletService {
	logrus.Info("Starting Wallet Service ....\n")
	time.Sleep(1 * time.Second)
	return &WalletService{
		Client: *client,
	}
}

func (w *WalletService) CreateNewFTNKey() (ethkey.KeyV2, error) {
	logrus.Info("Generating New FTN Key ....\n")
	time.Sleep(1 * time.Second)
	key, err := ethkey.NewV2()
	w.key = key
	if err != nil {
		logrus.Errorf("failed to generate ETH key %v", err)
		return ethkey.KeyV2{}, err
	}
	return ethkey.NewV2()
}

func (w *WalletService) PrintWalletDetails() {
	logrus.Info("Succesfully generated FTN Wallet")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("FTN Address")
	fmt.Println(w.key.Address.Hex())
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Balance")
	balance, err := w.Client.BalanceAt(context.Background(), w.key.Address, nil)
	if err != nil {
		logrus.Fatal("Failed to get balance:", err)
	}
	balanceInEther := new(big.Float).SetInt(balance)
	etherValue := new(big.Float).Quo(balanceInEther, big.NewFloat(1e18))
	fmt.Printf("%.18f FTN\n", etherValue)
	fmt.Println()
}
