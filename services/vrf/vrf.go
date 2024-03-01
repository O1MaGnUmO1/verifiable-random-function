package vrf_service

import (
	"fmt"
	"math/big"
	"os"
	"time"
	"vrf/keys/secp256k1"
	"vrf/keys/vrfkey"
	"vrf/utils"

	"github.com/sirupsen/logrus"
)

type VRFKeyStore interface {
	GenerateProof(id string, seed *big.Int) (vrfkey.Proof, error)
}

type VRFService struct {
	key vrfkey.KeyV2
}

func NewVRFService() *VRFService {
	logrus.Info("Starting VRF Service ....\n")
	time.Sleep(1 * time.Second)
	return &VRFService{}
}

func (vrf *VRFService) GenerateNewVRFKey() (vrfkey.KeyV2, error) {
	key, err := vrfkey.GetKeyIfEnvSet()
	if err != nil {
		logrus.Infof("Key is not set trying to generate %v", err)
	}
	if !key.PublicKey.IsZero() {
		vrf.key = key
		return vrf.key, nil
	}
	logrus.Info("Generating New VRF Key ....")
	time.Sleep(1 * time.Second)
	key, err = vrfkey.NewV2()
	if err != nil {
		logrus.Errorf("failed to generate VRF Key %v", err)
		return vrfkey.KeyV2{}, err
	}
	vrf.key = key
	encJson, err := vrf.key.ToEncryptedJSON(os.Getenv("EC_VRF_KEY_PASSWORD"), utils.FastScryptParams)
	if err != nil {
		logrus.Errorf("failed to encrypt key %v", err)
		return vrfkey.KeyV2{}, err
	}
	currentWd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("failed to get current working directory %v", err)
		return vrfkey.KeyV2{}, err
	}
	if err := os.WriteFile("vrf_key.json", encJson, 0644); err != nil {
		logrus.Errorf("failed to write json %v", err)
		return vrfkey.KeyV2{}, err
	}
	logrus.Infof("file vrf_key.json is saved in %s", currentWd)

	return vrf.key, nil
}

func (vrf *VRFService) NewVRFKeyForTesting(seed *big.Int) vrfkey.KeyV2 {
	return vrfkey.MustNewV2XXXTestingOnly(seed)
}

func (vrf *VRFService) GetVrfKey() vrfkey.KeyV2 {
	return vrf.key
}

func (vrf *VRFService) PrintVrfKey() {
	logrus.Info("Succesfully generated VRF Key")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Public key")
	fmt.Println(vrf.key.PublicKey.String())
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Uncompressed")
	uncompressed, err := vrf.key.PublicKey.StringUncompressed()
	if err != nil {
		panic(err)
	}
	fmt.Println(uncompressed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Key Hash")
	fmt.Println(vrf.key.PublicKey.MustHash().Hex())
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()

	pki, err := secp256k1.NewPublicKeyFromHex(vrf.key.PublicKey.String())
	if err != nil {
		panic(err)
	}
	point, err := pki.Point()
	if err != nil {
		panic(err)
	}
	x, y := secp256k1.Coordinates(point)
	fmt.Println(x)
	fmt.Println(y)
}
