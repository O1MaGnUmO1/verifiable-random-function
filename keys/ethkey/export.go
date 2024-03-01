package ethkey

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"vrf/utils"
)

type EncryptedEthKeyExport struct {
	KeyType string              `json:"keyType"`
	Address EIP55Address        `json:"address"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
}

func (key KeyV2) ToEncryptedJSON(password string, scryptParams utils.ScryptParams) (export []byte, err error) {
	// DEV: uuid is derived directly from the address, since it is not stored internally
	id, err := uuid.FromBytes(key.Address.Bytes()[:16])
	if err != nil {
		return nil, errors.Wrapf(err, "could not generate ethkey UUID")
	}
	dKey := &keystore.Key{
		Id:         id,
		Address:    key.Address,
		PrivateKey: key.privateKey,
	}
	return keystore.EncryptKey(dKey, password, scryptParams.N, scryptParams.P)
}

func GetKeyIfEnvSet() (KeyV2, error) {
	var jsonData []byte

	currentWd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("failed to get current working directory %v", err)
		return KeyV2{}, err
	}
	keyFilePath := filepath.Join(currentWd, "/ftn_key.json")
	if (utils.FileExists(keyFilePath) || os.Getenv("EC_FTN_KEY_JSON_PATH") != "") && os.Getenv("EC_FTN_KEY_PASSWORD") != "" {
		if os.Getenv("EC_FTN_KEY_JSON_PATH") != "" {
			jsonData, err = os.ReadFile(os.Getenv("EC_FTN_KEY_JSON_PATH"))
			if err != nil {
				logrus.Errorf("failed to read file %v", err)
				return KeyV2{}, err
			}
		} else {
			jsonData, err = os.ReadFile(keyFilePath)
			if err != nil {
				logrus.Errorf("failed to read file %v", err)
				return KeyV2{}, err
			}
		}
		dKey, err := keystore.DecryptKey(jsonData, os.Getenv("EC_FTN_KEY_PASSWORD"))
		if err != nil {
			return KeyV2{}, err
		}
		key := FromPrivateKey(dKey.PrivateKey)
		if err != nil {
			logrus.Errorf("failed get key from encrypted json %v", err)
			return KeyV2{}, err
		}
		return key, nil
	}
	errorMsg := "EC_FTN_KEY_PASSWORD is not set, if you want to save your keys please set it and re-run the program"
	logrus.Error(errorMsg)

	return KeyV2{}, fmt.Errorf(errorMsg)
}
