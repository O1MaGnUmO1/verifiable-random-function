package vrf_service

import (
	"fmt"
	"math/big"
	"testing"
	"vrf/keys/secp256k1"
	"vrf/keys/vrfkey"
	"vrf/keys/vrfkey/proof"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var vrfProofTemplate = `{
	pk: [
		%s,
		%s
	],
	gamma: [
		%s,
		%s
	],
	c: %s,
	s: %s,
	seed: %s,
	uWitness: %s,
	cGammaWitness: [
		%s,
		%s
	],
	sHashWitness: [
		%s,
		%s
	],
	zInv: %s
}
`

var rcTemplate = `{
	blockNum: %d,
	subId: %d,
	callbackGasLimit: %d,
	numWords: %d,
	sender: %s,
}
`

func TestVrfKeyX_Y(t *testing.T) {
	// keyHashString := "0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528"
	preSeedString := "73205574980244512779130235788648090841299858645147114963473582773838936145366"
	blockhashString := "0x783bec311ddd9a87b9d48b3afc9a1cf98051952bc59148108106078c6abe03e5"
	blockNum := uint64(9)
	senderString := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	secretKeyString := "10"
	subId := "1"
	callbackGasLimit := uint64(40000)
	numWords := uint64(5)
	// nativePayment := deployCmd.Bool("native-payment", false, "requestor of VRF request")

	// Generate V2Key from secret key.
	secretKey := decimal.RequireFromString(secretKeyString).BigInt()
	// secretKey := decimal.RequireFromString(secretKeyString).BigInt()
	key := vrfkey.MustNewV2XXXTestingOnly(secretKey)
	// if err != nil {
	// 	panic(err)
	// }
	uncompressed, err := key.PublicKey.StringUncompressed()
	if err != nil {
		panic(err)
	}
	pk := key.PublicKey
	pkh := pk.MustHash()
	fmt.Println("Compressed: ", pk.String())
	fmt.Println("Uncompressed: ", uncompressed)
	fmt.Println("Hash: ", pkh.String())

	//Parse big ints and hexes.
	// requestKeyHash := common.HexToHash(keyHashString)
	requestPreSeed := decimal.RequireFromString(preSeedString).BigInt()
	sender := common.HexToAddress(senderString)
	blockHash := common.HexToHash(blockhashString)
	// Ensure that the provided keyhash of the request matches the keyhash of the secret key.
	// if !bytes.Equal(requestKeyHash[:], pkh[:]) {
	// helpers.PanicErr(errors.New("invalid key hash"))
	// }

	//Generate proof.
	preSeed, err := proof.BigToSeed(requestPreSeed)
	if err != nil {
		(fmt.Errorf("unable to parse preseed: %w", err))
	}

	parsedSubId, ok := new(big.Int).SetString(subId, 10)
	if !ok {
		fmt.Errorf("unable to parse subID: %s %w", subId, err)
	}
	// extraArgs, err := extraargs.ExtraArgsV1(*nativePayment)
	// helpers.PanicErr(err)
	preSeedData := proof.PreSeedDataV2{
		PreSeed:          preSeed,
		BlockHash:        blockHash,
		BlockNum:         blockNum,
		SubId:            parsedSubId.Uint64(),
		CallbackGasLimit: uint32(callbackGasLimit),
		NumWords:         uint32(numWords),
		Sender:           sender,
	}
	finalSeed := proof.FinalSeedV2(preSeedData)
	p, err := key.GenerateProof(finalSeed)
	if err != nil {
		fmt.Errorf("unable to generate proof: %w", err)
	}
	onChainProof, rc, err := proof.GenerateProofResponseFromProofV2(p, preSeedData)
	if err != nil {
		fmt.Errorf("unable to generate proof response: %w", err)
	}
	_, err = p.VerifyVRFProof()
	if err != nil {
		panic(err)
	}

	//Print formatted VRF proof.
	fmt.Println("ON-CHAIN PROOF:")
	fmt.Printf(
		vrfProofTemplate,
		onChainProof.Pk[0],
		onChainProof.Pk[1],
		onChainProof.Gamma[0],
		onChainProof.Gamma[1],
		onChainProof.C,
		onChainProof.S,
		onChainProof.Seed,
		onChainProof.UWitness,
		onChainProof.CGammaWitness[0],
		onChainProof.CGammaWitness[1],
		onChainProof.SHashWitness[0],
		onChainProof.SHashWitness[1],
		onChainProof.ZInv,
	)

	// Print formatted request commitment.
	fmt.Println("\nREQUEST COMMITMENT:")
	fmt.Printf(
		rcTemplate,
		rc.BlockNum,
		rc.SubId,
		rc.CallbackGasLimit,
		rc.NumWords,
		rc.Sender,
	)

	pki, err := secp256k1.NewPublicKeyFromHex(pk.String())
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
