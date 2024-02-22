package vrf_service

import (
	"fmt"
	"math/big"
	"time"
	"vrf/keys/vrfkey"

	"github.com/sirupsen/logrus"
)

type VRFService struct {
	key vrfkey.KeyV2
}

func NewVRFService() *VRFService {
	logrus.Info("Starting VRF Service ....\n")
	time.Sleep(1 * time.Second)
	return &VRFService{}
}

func (vrf *VRFService) GenerateNewVRFKey() (vrfkey.KeyV2, error) {
	logrus.Info("Generating New VRF Key ....\n")
	time.Sleep(1 * time.Second)
	key, err := vrfkey.NewV2()
	if err != nil {
		logrus.Errorf("failed to generate VRF Key %v", err)
		return vrfkey.KeyV2{}, err
	}
	vrf.key = key
	return key, nil
}

func (vrf *VRFService) NewVRFKeyForTesting(seed *big.Int) vrfkey.KeyV2 {
	return vrfkey.MustNewV2XXXTestingOnly(seed)
}

func (vrf *VRFService) PrintVrfKey() {
	fmt.Println("Succesfully generated VRF Key")
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
}
