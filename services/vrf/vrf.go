package vrf_service

import (
	"fmt"
	"math/big"
	"vrf/keys/vrfkey"
)

func GenerateNewVRFKey() (vrfkey.KeyV2, error) {
	return vrfkey.NewV2()
}

func NewVRFKeyForTesting(seed *big.Int) vrfkey.KeyV2 {
	return vrfkey.MustNewV2XXXTestingOnly(seed)
}

func PrintVrfKey(key vrfkey.KeyV2) {
	fmt.Printf("VRF key is %s", key.GoString())
}
