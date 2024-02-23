package ethkey

import (
	"os"

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
	if (os.Getenv("EC_FTN_KEY_JSON_PATH")) != "" && os.Getenv("EC_FTN_KEY_PASSWORD") != "" {
		jsonData, err := os.ReadFile(os.Getenv("EC_FTN_KEY_JSON_PATH"))
		if err != nil {
			logrus.Errorf("failed to read file %v", err)
			return KeyV2{}, err
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
	return KeyV2{}, nil
}
