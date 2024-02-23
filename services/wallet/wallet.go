package wallet_service

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"
	"vrf/client"
	"vrf/keys/ethkey"
	"vrf/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Balances struct {
	OldBalance *big.Int
	NewBalance *big.Int
}

type WalletService struct {
	Client    client.Client
	Key       ethkey.KeyV2
	BalanceCh chan Balances
}

func NewWalletService(client *client.Client) *WalletService {
	logrus.Info("Starting Wallet Service ....\n")
	time.Sleep(1 * time.Second)
	return &WalletService{
		Client:    *client,
		BalanceCh: make(chan Balances),
	}
}

func (w *WalletService) CreateNewFTNKey() (ethkey.KeyV2, error) {
	if err := os.Setenv("EC_FTN_KEY_JSON_PATH", "ftn_key.json"); err != nil {
		logrus.Errorf("failed to set EC_FTN_KEY_JSON_PATH env %v", err)
		return ethkey.KeyV2{}, err
	}
	ethKey, err := ethkey.GetKeyIfEnvSet()
	if err != nil {
		logrus.Infof("Key is not set trying to generate %v", err)
	}
	if ethKey.Address.Cmp(common.Address{}) != 0 {
		w.Key = ethKey
		return w.Key, nil
	}
	logrus.Info("Generating New FTN Key ....")
	time.Sleep(1 * time.Second)
	key, err := ethkey.NewV2()
	if err != nil {
		logrus.Errorf("failed to generate FTN key %v", err)
		return ethkey.KeyV2{}, err
	}
	w.Key = key
	if os.Getenv("EC_FTN_KEY_PASSWORD") != "" {
		encJson, err := w.Key.ToEncryptedJSON(os.Getenv("EC_FTN_KEY_PASSWORD"), utils.FastScryptParams)
		if err != nil {
			logrus.Errorf("failed to encrypt key %v", err)
			return ethkey.KeyV2{}, err
		}
		if err := os.WriteFile("ftn_key.json", encJson, 0644); err != nil {
			logrus.Errorf("failed to write json %v", err)
			return ethkey.KeyV2{}, err
		}
	}
	return w.Key, nil
}

func (w *WalletService) MonitorBalanceChanges() {
	// Get the initial balance
	prevBalance, err := getETHBalance(w.Client.EthClient, w.Key.Address)
	if err != nil {
		logrus.Errorf("failed to get balance %v", err)
		return
	}

	// Start monitoring balance changes
	go func() {
		for {
			// Get the current balance
			currBalance, err := getETHBalance(w.Client.EthClient, w.Key.Address)
			if err != nil {
				logrus.Errorf("failed to get current balance %v", err)
				return
			}

			// Check if the balance has changed
			if currBalance.Cmp(prevBalance) == 1 {
				// Send a balance change signal
				w.BalanceCh <- Balances{
					OldBalance: prevBalance,
					NewBalance: currBalance,
				}

				// Update the previous balance
				prevBalance = currBalance
			}
		}
	}()

	// Block indefinitely
	select {}

}

func getETHBalance(client *ethclient.Client, address common.Address) (*big.Int, error) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// weiToETH converts Wei to Ether and returns the result as a string
func WeiToETH(wei *big.Int) string {
	ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1000000000000000000))
	return ether.Text('f', 18) // Format with 18 decimal places
}

func (w *WalletService) PrintWalletDetails() {
	logrus.Info("Succesfully generated FTN Wallet")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("FTN Address")
	fmt.Println(w.Key.Address.Hex())
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Balance")
	balance, err := w.Client.BalanceAt(context.Background(), w.Key.Address, nil)
	if err != nil {
		logrus.Fatal("Failed to get balance:", err)
	}
	balanceInEther := new(big.Float).SetInt(balance)
	etherValue := new(big.Float).Quo(balanceInEther, big.NewFloat(1e18))
	fmt.Printf("%.18f FTN\n", etherValue)
	fmt.Println()
}
