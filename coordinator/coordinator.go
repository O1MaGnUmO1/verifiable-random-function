// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coordinator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VRFCoordinatorV2FeeConfig is an auto generated low-level Go binding around an user-defined struct.
type VRFCoordinatorV2FeeConfig struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}

// VRFCoordinatorV2RequestCommitment is an auto generated low-level Go binding around an user-defined struct.
type VRFCoordinatorV2RequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
}

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

// CoordinatorMetaData contains all meta data concerning the Coordinator contract.
var CoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"InvalidBlockhash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier2\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier3\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier4\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier5\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structVRFCoordinatorV2.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSubId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier2\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier3\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier4\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier5\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"}],\"name\":\"getFeeTier\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_subscriptionConfigs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"requestedOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_subscriptions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier2\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier3\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier4\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"reqsForTier5\",\"type\":\"uint24\"}],\"internalType\":\"structVRFCoordinatorV2.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162004984380380620049848339810160408190526200003491620001a9565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e0565b5050506001600160601b0319606092831b8116608052911b1660a052620001e1565b6001600160a01b0381163314156200013b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001a457600080fd5b919050565b60008060408385031215620001bd57600080fd5b620001c8836200018c565b9150620001d8602084016200018c565b90509250929050565b60805160601c60a05160601c6147466200023e600039600081816105420152818161302101526130480152600081816102e20152818161109c01528181611c4f01528181612446015281816125460152612a6201526147466000f3fe608060405234801561001057600080fd5b506004361061021b5760003560e01c80637341c10c11610125578063af198b97116100ad578063d7ae1d301161007c578063d7ae1d30146105fb578063d7ca6f9b1461060e578063e72f6e3014610662578063e82ad7d414610675578063f2fde38b1461069857600080fd5b8063af198b9714610564578063c3f909d41461058f578063caf70c4a146105d5578063d2f9f9a7146105e857600080fd5b80639f87fad7116100f45780639f87fad7146104ec578063a21a23e4146104ff578063a47c769614610507578063a4c0ed361461052a578063ad1783611461053d57600080fd5b80637341c10c146104ad57806379ba5097146104c057806382359740146104c85780638da5cb5b146104db57600080fd5b806340d6bb82116101a857806364d51a2a1161017757806364d51a2a146103fd57806366316d8d146104055780636840c05e1461041857806369bcdb7d1461047a5780636f64f03f1461049a57600080fd5b806340d6bb82146103245780634cb48a54146103425780635d3b1d30146103555780635fbbc0d21461036857600080fd5b806308821d58116101ef57806308821d581461028d57806312b58349146102a057806315c48b84146102c25780631b6b6d23146102dd578063356dac711461031c57600080fd5b80620122911461022057806302bcc5b61461024057806304c357cb1461025557806306bfa63714610268575b600080fd5b6102286106ab565b60405161023793929190614334565b60405180910390f35b61025361024e3660046141e8565b610727565b005b610253610263366004614203565b61079e565b6005546001600160401b03165b6040516001600160401b039091168152602001610237565b61025361029b366004613f17565b6108e9565b600554600160401b90046001600160601b03165b604051908152602001610237565b6102ca60c881565b60405161ffff9091168152602001610237565b6103047f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610237565b600a546102b4565b61032d6101f481565b60405163ffffffff9091168152602001610237565b610253610350366004614097565b610a7d565b6102b4610363366004613f71565b610cd3565b600c546040805163ffffffff8084168252640100000000840481166020830152600160401b8404811692820192909252600160601b830482166060820152600160801b8304909116608082015262ffffff600160a01b8304811660a0830152600160b81b8304811660c0830152600160d01b8304811660e0830152600160e81b90920490911661010082015261012001610237565b6102ca606481565b610253610413366004613ed4565b610f9f565b6104536104263660046141e8565b6004602052600090815260409020546001600160601b03811690600160601b90046001600160401b031682565b604080516001600160601b0390931683526001600160401b03909116602083015201610237565b6102b46104883660046141b6565b60009081526009602052604090205490565b6102536104a8366004613e1a565b611178565b6102536104bb366004614203565b611277565b61025361142a565b6102536104d63660046141e8565b6114d4565b6000546001600160a01b0316610304565b6102536104fa366004614203565b611630565b610275611982565b61051a6105153660046141e8565b611b1f565b60405161023794939291906144d2565b610253610538366004613e4e565b611c19565b6103047f000000000000000000000000000000000000000000000000000000000000000081565b610577610572366004613fcf565b611de8565b6040516001600160601b039091168152602001610237565b600b546040805161ffff8316815263ffffffff62010000840481166020830152600160381b8404811692820192909252600160581b909204166060820152608001610237565b6102b46105e3366004613f33565b6121a7565b61032d6105f63660046141e8565b6121d7565b610253610609366004614203565b612354565b61064261061c3660046141e8565b600360205260009081526040902080546001909101546001600160a01b03918216911682565b604080516001600160a01b03938416815292909116602083015201610237565b610253610670366004613dff565b612426565b6106886106833660046141e8565b61260e565b6040519015158152602001610237565b6102536106a6366004613dff565b6127b9565b600b546007805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff1693919283919083018282801561071557602002820191906000526020600020905b815481526020019060010190808311610701575b50505050509050925092509250909192565b61072f6127ca565b6001600160401b0381166000908152600360205260409020546001600160a01b031661076e57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b03811660009081526003602052604090205461079b9082906001600160a01b031661281f565b50565b6001600160401b03821660009081526003602052604090205482906001600160a01b0316806107e057604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b0382161461081957604051636c51fda960e11b81526001600160a01b03821660048201526024015b60405180910390fd5b600b54600160301b900460ff16156108445760405163769dd35360e11b815260040160405180910390fd5b6001600160401b0384166000908152600360205260409020600101546001600160a01b038481169116146108e3576001600160401b03841660008181526003602090815260409182902060010180546001600160a01b0319166001600160a01b0388169081179091558251338152918201527f69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be91015b60405180910390a25b50505050565b6108f16127ca565b6040805180820182526000916109209190849060029083908390808284376000920191909152506121a7915050565b6000818152600660205260409020549091506001600160a01b03168061095c57604051631dfd6e1360e21b815260048101839052602401610810565b600082815260066020526040812080546001600160a01b03191690555b600754811015610a34578260078281548110610997576109976146e4565b90600052602060002001541415610a225760078054600091906109bc9060019061460d565b815481106109cc576109cc6146e4565b9060005260206000200154905080600783815481106109ed576109ed6146e4565b6000918252602090912001556007805480610a0a57610a0a6146ce565b60019003818190600052602060002001600090559055505b80610a2c8161464c565b915050610979565b50806001600160a01b03167f72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d83604051610a7091815260200190565b60405180910390a2505050565b610a856127ca565b60c861ffff87161115610abf5760405163539c34bb60e11b815261ffff871660048201819052602482015260c86044820152606401610810565b60008213610ae3576040516321ea67b360e11b815260048101839052602401610810565b6040805160a0808201835261ffff891680835263ffffffff89811660208086018290526000868801528a831660608088018290528b85166080988901819052600b805465ffffffffffff191690971762010000909502949094176effffffffffffffffff0000000000001916600160381b90920263ffffffff60581b191691909117600160581b909302929092179093558651600c80549489015189890151938a0151978a0151968a015160c08b015160e08c01516101008d015195881667ffffffffffffffff199099169890981764010000000093881693909302929092176fffffffffffffffff00000000000000001916600160401b9587169590950263ffffffff60601b191694909417600160601b988616989098029790971766ffffffffffffff60801b1916600160801b969094169590950262ffffff60a01b191692909217600160a01b62ffffff928316021765ffffffffffff60b81b1916600160b81b9582169590950262ffffff60d01b191694909417600160d01b92851692909202919091176001600160e81b0316600160e81b9390911692909202919091178155600a84905590517fc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb291610cc3918991899189918991899190614393565b60405180910390a1505050505050565b600b54600090600160301b900460ff1615610d015760405163769dd35360e11b815260040160405180910390fd5b6001600160401b0385166000908152600360205260409020546001600160a01b0316610d4057604051630fb532db60e11b815260040160405180910390fd5b3360009081526002602090815260408083206001600160401b03808a1685529252909120541680610d9557604051637800cff360e11b81526001600160401b0387166004820152336024820152604401610810565b600b5461ffff9081169086161080610db1575060c861ffff8616115b15610de857600b5460405163539c34bb60e11b815261ffff8088166004830152909116602482015260c86044820152606401610810565b600b5463ffffffff6201000090910481169085161115610e3657600b54604051637aebf00f60e11b815263ffffffff8087166004830152620100009092049091166024820152604401610810565b6101f463ffffffff84161115610e6f576040516311ce1afb60e21b815263ffffffff841660048201526101f46024820152604401610810565b6000610e7c82600161458d565b9050600080610e8d8a338b86612b96565b604080516020810184905243918101919091526001600160401b038c16606082015263ffffffff808b166080830152891660a08201523360c0820152919350915060e00160408051808303601f19018152828252805160209182012060008681526009835283902055848352820183905261ffff8a169082015263ffffffff80891660608301528716608082015233906001600160401b038b16908c907f63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a97729060a00160405180910390a4503360009081526002602090815260408083206001600160401b03808d168552925290912080549190931667ffffffffffffffff199091161790915591505095945050505050565b600b54600160301b900460ff1615610fca5760405163769dd35360e11b815260040160405180910390fd5b336000908152600860205260409020546001600160601b038083169116101561100657604051631e9acf1760e31b815260040160405180910390fd5b336000908152600860205260408120805483929061102e9084906001600160601b0316614624565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600560088282829054906101000a90046001600160601b03166110769190614624565b92506101000a8154816001600160601b0302191690836001600160601b031602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb83836040518363ffffffff1660e01b81526004016111059291906001600160a01b039290921682526001600160601b0316602082015260400190565b602060405180830381600087803b15801561111f57600080fd5b505af1158015611133573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111579190613f4f565b61117457604051631e9acf1760e31b815260040160405180910390fd5b5050565b6111806127ca565b6040805180820182526000916111af9190849060029083908390808284376000920191909152506121a7915050565b6000818152600660205260409020549091506001600160a01b0316156111eb57604051634a0b8fa760e01b815260048101829052602401610810565b600081815260066020908152604080832080546001600160a01b0319166001600160a01b0388169081179091556007805460018101825594527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b89101610a70565b6001600160401b03821660009081526003602052604090205482906001600160a01b0316806112b957604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146112ed57604051636c51fda960e11b81526001600160a01b0382166004820152602401610810565b600b54600160301b900460ff16156113185760405163769dd35360e11b815260040160405180910390fd5b6001600160401b03841660009081526003602052604090206002015460641415611355576040516305a48e0f60e01b815260040160405180910390fd5b6001600160a01b03831660009081526002602090815260408083206001600160401b038089168552925290912054161561138e576108e3565b6001600160a01b03831660008181526002602081815260408084206001600160401b038a16808652908352818520805467ffffffffffffffff19166001908117909155600384528286209094018054948501815585529382902090920180546001600160a01b03191685179055905192835290917f43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e091016108da565b6001546001600160a01b0316331461147d5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b6044820152606401610810565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600b54600160301b900460ff16156114ff5760405163769dd35360e11b815260040160405180910390fd5b6001600160401b0381166000908152600360205260409020546001600160a01b031661153e57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b0381166000908152600360205260409020600101546001600160a01b031633146115ab576001600160401b0381166000908152600360205260409081902060010154905163d084e97560e01b81526001600160a01b039091166004820152602401610810565b6001600160401b0381166000818152600360209081526040918290208054336001600160a01b0319808316821784556001909301805490931690925583516001600160a01b03909116808252928101919091529092917f6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0910160405180910390a25050565b6001600160401b03821660009081526003602052604090205482906001600160a01b03168061167257604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146116a657604051636c51fda960e11b81526001600160a01b0382166004820152602401610810565b600b54600160301b900460ff16156116d15760405163769dd35360e11b815260040160405180910390fd5b6116da8461260e565b156116f857604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b03831660009081526002602090815260408083206001600160401b0380891685529252909120541661175e57604051637800cff360e11b81526001600160401b03851660048201526001600160a01b0384166024820152604401610810565b6001600160401b0384166000908152600360209081526040808320600201805482518185028101850190935280835291929091908301828280156117cb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116117ad575b505050505090506000600182516117e2919061460d565b905060005b825181101561190957856001600160a01b031683828151811061180c5761180c6146e4565b60200260200101516001600160a01b031614156118f7576000838381518110611837576118376146e4565b6020026020010151905080600360008a6001600160401b03166001600160401b03168152602001908152602001600020600201838154811061187b5761187b6146e4565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092556001600160401b038a1681526003909152604090206002018054806118cf576118cf6146ce565b600082815260209020810160001990810180546001600160a01b031916905501905550611909565b806119018161464c565b9150506117e7565b506001600160a01b03851660008181526002602090815260408083206001600160401b038b1680855290835292819020805467ffffffffffffffff191690555192835290917f182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b91015b60405180910390a2505050505050565b600b54600090600160301b900460ff16156119b05760405163769dd35360e11b815260040160405180910390fd5b600580546001600160401b03169060006119c983614667565b82546101009290920a6001600160401b03818102199093169183160217909155600554169050600080604051908082528060200260200182016040528015611a1b578160200160208202803683370190505b50604080518082018252600080825260208083018281526001600160401b03888116808552600484528685209551865493516001600160601b039091166001600160a01b031994851617600160601b919093160291909117909455845160608101865233815280830184815281870188815295855260038452959093208351815483166001600160a01b0391821617825595516001820180549093169616959095179055915180519495509093611ad89260028501920190613b74565b50506040513381526001600160401b03841691507f464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf9060200160405180910390a250905090565b6001600160401b038116600090815260036020526040812054819081906060906001600160a01b0316611b6557604051630fb532db60e11b815260040160405180910390fd5b6001600160401b0380861660009081526004602090815260408083205460038352928190208054600290910180548351818602810186019094528084526001600160601b03861696600160601b909604909516946001600160a01b03909216939092918391830182828015611c0357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611be5575b5050505050905093509350935093509193509193565b600b54600160301b900460ff1615611c445760405163769dd35360e11b815260040160405180910390fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611c8d576040516344b0e3c360e01b815260040160405180910390fd5b60208114611cae57604051638129bbcd60e01b815260040160405180910390fd5b6000611cbc828401846141e8565b6001600160401b0381166000908152600360205260409020549091506001600160a01b0316611cfe57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b038116600090815260046020526040812080546001600160601b031691869190611d2f83856145b8565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600560088282829054906101000a90046001600160601b0316611d7791906145b8565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550816001600160401b03167fd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8828784611dd39190614575565b60408051928352602083019190915201611972565b600b54600090600160301b900460ff1615611e165760405163769dd35360e11b815260040160405180910390fd5b60005a90506000806000611e2a8787612c10565b9250925092506000866060015163ffffffff166001600160401b03811115611e5457611e546146fa565b604051908082528060200260200182016040528015611e7d578160200160208202803683370190505b50905060005b876060015163ffffffff16811015611ef15760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c828281518110611ed457611ed46146e4565b602090810291909101015280611ee98161464c565b915050611e83565b50600083815260096020526040808220829055518190631fe543e360e01b90611f209087908690602401614484565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600b805466ff0000000000001916600160301b179055908a015160808b0151919250600091611f889163ffffffff169084612dc6565b600b805466ff000000000000191690556020808c0180516001600160401b03908116600090815260049093526040808420549251821684529092208054939450600160601b918290048316936001939192600c92611fea92869290041661458d565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550600061203f8a600b600001600b9054906101000a900463ffffffff1663ffffffff16612039856121d7565b3a612e14565b6020808e01516001600160401b03166000908152600490915260409020549091506001600160601b038083169116101561208c57604051631e9acf1760e31b815260040160405180910390fd5b6020808d01516001600160401b0316600090815260049091526040812080548392906120c29084906001600160601b0316614624565b82546101009290920a6001600160601b0381810219909316918316021790915560008b8152600660209081526040808320546001600160a01b031683526008909152812080548594509092612119918591166145b8565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550877f7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e488838660405161218d939291909283526001600160601b039190911660208301521515604082015260600190565b60405180910390a299505050505050505050505b92915050565b6000816040516020016121ba9190614326565b604051602081830303815290604052805190602001209050919050565b6040805161012081018252600c5463ffffffff8082168352640100000000820481166020840152600160401b8204811693830193909352600160601b810483166060830152600160801b8104909216608082015262ffffff600160a01b8304811660a08301819052600160b81b8404821660c0840152600160d01b8404821660e0840152600160e81b909304166101008201526000916001600160401b03841611612283575192915050565b826001600160401b03168160a0015162ffffff161080156122b657508060c0015162ffffff16836001600160401b031611155b156122c5576020015192915050565b826001600160401b03168160c0015162ffffff161080156122f857508060e0015162ffffff16836001600160401b031611155b15612307576040015192915050565b826001600160401b03168160e0015162ffffff1610801561233b575080610100015162ffffff16836001600160401b031611155b1561234a576060015192915050565b6080015192915050565b6001600160401b03821660009081526003602052604090205482906001600160a01b03168061239657604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146123ca57604051636c51fda960e11b81526001600160a01b0382166004820152602401610810565b600b54600160301b900460ff16156123f55760405163769dd35360e11b815260040160405180910390fd5b6123fe8461260e565b1561241c57604051631685ecdd60e31b815260040160405180910390fd5b6108e3848461281f565b61242e6127ca565b6040516370a0823160e01b81523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a082319060240160206040518083038186803b15801561249057600080fd5b505afa1580156124a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c891906141cf565b600554909150600160401b90046001600160601b031681811115612509576040516354ced18160e11b81526004810182905260248101839052604401610810565b8181101561260957600061251d828461460d565b60405163a9059cbb60e01b81526001600160a01b038681166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb90604401602060405180830381600087803b15801561258c57600080fd5b505af11580156125a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c49190613f4f565b50604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600910160405180910390a1505b505050565b6001600160401b0381166000908152600360209081526040808320815160608101835281546001600160a01b03908116825260018301541681850152600282018054845181870281018701865281815287969395860193909291908301828280156126a257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612684575b505050505081525050905060005b8160400151518110156127af5760005b60075481101561279c576000612765600783815481106126e2576126e26146e4565b906000526020600020015485604001518581518110612703576127036146e4565b6020026020010151886002600089604001518981518110612726576127266146e4565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b03808f1683529352205416612b96565b50600081815260096020526040902054909150156127895750600195945050505050565b50806127948161464c565b9150506126c0565b50806127a78161464c565b9150506126b0565b5060009392505050565b6127c16127ca565b61079b81612eea565b6000546001600160a01b0316331461281d5760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b6044820152606401610810565b565b600b54600160301b900460ff161561284a5760405163769dd35360e11b815260040160405180910390fd5b6001600160401b0382166000908152600360209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152929593948601938301828280156128da57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116128bc575b505050919092525050506001600160401b0380851660009081526004602090815260408083208151808301909252546001600160601b038116808352600160601b909104909416918101919091529293505b8360400151518110156129ad576002600085604001518381518110612953576129536146e4565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b038a1682529092529020805467ffffffffffffffff19169055806129a58161464c565b91505061292c565b506001600160401b038516600090815260036020526040812080546001600160a01b031990811682556001820180549091169055906129ef6002830182613bd9565b50506001600160401b038516600090815260046020526040902080546001600160a01b031916905560058054829190600890612a3c908490600160401b90046001600160601b0316614624565b92506101000a8154816001600160601b0302191690836001600160601b031602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb85836001600160601b03166040518363ffffffff1660e01b8152600401612acb9291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015612ae557600080fd5b505af1158015612af9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b1d9190613f4f565b612b3a57604051631e9acf1760e31b815260040160405180910390fd5b604080516001600160a01b03861681526001600160601b03831660208201526001600160401b038716917fe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815910160405180910390a25050505050565b6040805160208082018790526001600160a01b0395909516818301526001600160401b039384166060820152919092166080808301919091528251808303909101815260a08201835280519084012060c082019490945260e080820185905282518083039091018152610100909101909152805191012091565b6000806000612c2285600001516121a7565b6000818152600660205260409020549093506001600160a01b031680612c5e57604051631dfd6e1360e21b815260048101859052602401610810565b6080860151604051612c7d918691602001918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600990935291205490935080612cc357604051631b44092560e11b815260040160405180910390fd5b85516020808801516040808a015160608b015160808c01519251612d2e968b9690959491019586526001600160401b03948516602087015292909316604085015263ffffffff90811660608501529190911660808301526001600160a01b031660a082015260c00190565b604051602081830303815290604052805190602001208114612d635760405163354a450b60e21b815260040160405180910390fd5b855160808801516040516001600160401b039092164091600091612d94918490602001918252602082015260400190565b6040516020818303038152906040528051906020012060001c9050612db98982612f94565b9450505050509250925092565b60005a611388811015612dd857600080fd5b611388810390508460408204820311612df057600080fd5b50823b612dfc57600080fd5b60008083516020850160008789f190505b9392505050565b600080612e1f612fff565b905060008113612e45576040516321ea67b360e11b815260048101829052602401610810565b6000815a612e538989614575565b612e5d919061460d565b612e6790866145ee565b612e7990670de0b6b3a76400006145ee565b612e8391906145da565b90506000612e9c63ffffffff871664e8d4a510006145ee565b9050612eb4816b033b2e3c9fd0803ce800000061460d565b821115612ed45760405163e80fa38160e01b815260040160405180910390fd5b612ede8183614575565b98975050505050505050565b6001600160a01b038116331415612f435760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610810565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612fc88360000151846020015185604001518660600151868860a001518960c001518a60e001518b610100015161310d565b60038360200151604051602001612fe0929190614470565b60408051601f1981840301815291905280516020909101209392505050565b600b54600090600160381b900463ffffffff1680151582806001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016156130e0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561309f57600080fd5b505afa1580156130b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130d7919061422d565b50945090925050505b8280156130fb57506130f2824261460d565b8463ffffffff16105b156131055750600a545b949350505050565b61311689613330565b6131625760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e2063757276650000000000006044820152606401610810565b61316b88613330565b6131af5760405162461bcd60e51b815260206004820152601560248201527467616d6d61206973206e6f74206f6e20637572766560581b6044820152606401610810565b6131b883613330565b6132045760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610810565b61320d82613330565b6132595760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610810565b613265878a88876133f3565b6132b15760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e657373000000000000006044820152606401610810565b60006132bd8a87613516565b905060006132d0898b878b86898961357a565b905060006132e1838d8d8a86613693565b9050808a146133225760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b210383937b7b360991b6044820152606401610810565b505050505050505050505050565b80516000906401000003d0191161337e5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420782d6f7264696e61746560701b6044820152606401610810565b60208201516401000003d019116133cc5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420792d6f7264696e61746560701b6044820152606401610810565b60208201516401000003d0199080096133ec8360005b60200201516136d3565b1492915050565b60006001600160a01b0382166134395760405162461bcd60e51b815260206004820152600b60248201526a626164207769746e65737360a81b6044820152606401610810565b60208401516000906001161561345057601c613453565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa1580156134ee573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b61351e613bf7565b61354b6001848460405160200161353793929190614305565b6040516020818303038152906040526136f7565b90505b61355781613330565b6121a15780516040805160208101929092526135739101613537565b905061354e565b613582613bf7565b825186516401000003d01990819006910614156135e15760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610810565b6135ec878988613746565b6136315760405162461bcd60e51b8152602060048201526016602482015275119a5c9cdd081b5d5b0818da1958dac819985a5b195960521b6044820152606401610810565b61363c848685613746565b6136885760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c65640000000000000000006044820152606401610810565b612ede86848461386e565b6000600286868685876040516020016136b1969594939291906142a6565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b6136ff613bf7565b61370882613935565b815261371d6137188260006133e2565b613970565b602082018190526002900660011415613741576020810180516401000003d0190390525b919050565b6000826137835760405162461bcd60e51b815260206004820152600b60248201526a3d32b9379039b1b0b630b960a91b6044820152606401610810565b835160208501516000906137999060029061468e565b156137a557601c6137a8565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa15801561381a573d6000803e3d6000fd5b5050506020604051035190506000866040516020016138399190614294565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b613876613bf7565b83516020808601518551918601516000938493849361389793909190613990565b919450925090506401000003d0198582096001146138f75760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a000000000000006044820152606401610810565b60405180604001604052806401000003d01980613916576139166146b8565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d01981106137415760408051602080820193909352815180820384018152908201909152805191012061393d565b60006121a18260026139896401000003d0196001614575565b901c613a70565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a08905060006139d083838585613b07565b90985090506139e188828e88613b2b565b90985090506139f288828c87613b2b565b90985090506000613a058d878b85613b2b565b9098509050613a1688828686613b07565b9098509050613a2788828e89613b2b565b9098509050818114613a5c576401000003d019818a0998506401000003d01982890997506401000003d0198183099650613a60565b8196505b5050505050509450945094915050565b600080613a7b613c15565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152613aad613c33565b60208160c0846005600019fa925082613afd5760405162461bcd60e51b81526020600482015260126024820152716269674d6f64457870206661696c7572652160701b6044820152606401610810565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215613bc9579160200282015b82811115613bc957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613b94565b50613bd5929150613c51565b5090565b508054600082559060005260206000209081019061079b9190613c51565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115613bd55760008155600101613c52565b80356001600160a01b038116811461374157600080fd5b80604081018310156121a157600080fd5b600082601f830112613c9f57600080fd5b604051604081018181106001600160401b0382111715613cc157613cc16146fa565b8060405250808385604086011115613cd857600080fd5b60005b6002811015613cfa578135835260209283019290910190600101613cdb565b509195945050505050565b600060a08284031215613d1757600080fd5b60405160a081018181106001600160401b0382111715613d3957613d396146fa565b604052905080613d4883613dce565b8152613d5660208401613dce565b6020820152613d6760408401613dba565b6040820152613d7860608401613dba565b6060820152613d8960808401613c66565b60808201525092915050565b803561ffff8116811461374157600080fd5b803562ffffff8116811461374157600080fd5b803563ffffffff8116811461374157600080fd5b80356001600160401b038116811461374157600080fd5b805169ffffffffffffffffffff8116811461374157600080fd5b600060208284031215613e1157600080fd5b612e0d82613c66565b60008060608385031215613e2d57600080fd5b613e3683613c66565b9150613e458460208501613c7d565b90509250929050565b60008060008060608587031215613e6457600080fd5b613e6d85613c66565b93506020850135925060408501356001600160401b0380821115613e9057600080fd5b818701915087601f830112613ea457600080fd5b813581811115613eb357600080fd5b886020828501011115613ec557600080fd5b95989497505060200194505050565b60008060408385031215613ee757600080fd5b613ef083613c66565b915060208301356001600160601b0381168114613f0c57600080fd5b809150509250929050565b600060408284031215613f2957600080fd5b612e0d8383613c7d565b600060408284031215613f4557600080fd5b612e0d8383613c8e565b600060208284031215613f6157600080fd5b81518015158114612e0d57600080fd5b600080600080600060a08688031215613f8957600080fd5b85359450613f9960208701613dce565b9350613fa760408701613d95565b9250613fb560608701613dba565b9150613fc360808701613dba565b90509295509295909350565b600080828403610240811215613fe457600080fd5b6101a080821215613ff457600080fd5b613ffc61454c565b91506140088686613c8e565b82526140178660408701613c8e565b60208301526080850135604083015260a0850135606083015260c0850135608083015261404660e08601613c66565b60a083015261010061405a87828801613c8e565b60c084015261406d876101408801613c8e565b60e0840152610180860135818401525081935061408c86828701613d05565b925050509250929050565b6000806000806000808688036101c08112156140b257600080fd5b6140bb88613d95565b96506140c960208901613dba565b95506140d760408901613dba565b94506140e560608901613dba565b93506080880135925061012080609f198301121561410257600080fd5b61410a61454c565b915061411860a08a01613dba565b825261412660c08a01613dba565b602083015261413760e08a01613dba565b604083015261010061414a818b01613dba565b606084015261415a828b01613dba565b608084015261416c6101408b01613da7565b60a084015261417e6101608b01613da7565b60c08401526141906101808b01613da7565b60e08401526141a26101a08b01613da7565b818401525050809150509295509295509295565b6000602082840312156141c857600080fd5b5035919050565b6000602082840312156141e157600080fd5b5051919050565b6000602082840312156141fa57600080fd5b612e0d82613dce565b6000806040838503121561421657600080fd5b61421f83613dce565b9150613e4560208401613c66565b600080600080600060a0868803121561424557600080fd5b61424e86613de5565b9450602086015193506040860151925060608601519150613fc360808701613de5565b8060005b60028110156108e3578151845260209384019390910190600101614275565b61429e8183614271565b604001919050565b8681526142b66020820187614271565b6142c36060820186614271565b6142d060a0820185614271565b6142dd60e0820184614271565b60609190911b6bffffffffffffffffffffffff19166101208201526101340195945050505050565b8381526143156020820184614271565b606081019190915260800192915050565b604081016121a18284614271565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b8181101561438557845183529383019391830191600101614369565b509098975050505050505050565b60006101c08201905061ffff8816825263ffffffff808816602084015280871660408401528086166060840152846080840152835481811660a08501526143e760c08501838360201c1663ffffffff169052565b6143fe60e08501838360401c1663ffffffff169052565b6144166101008501838360601c1663ffffffff169052565b61442e6101208501838360801c1663ffffffff169052565b62ffffff60a082901c811661014086015260b882901c811661016086015260d082901c1661018085015260e81c6101a090930192909252979650505050505050565b82815260608101612e0d6020830184614271565b6000604082018483526020604081850152818551808452606086019150828701935060005b818110156144c5578451835293830193918301916001016144a9565b5090979650505050505050565b6000608082016001600160601b038716835260206001600160401b0387168185015260018060a01b0380871660408601526080606086015282865180855260a087019150838801945060005b8181101561453c57855184168352948401949184019160010161451e565b50909a9950505050505050505050565b60405161012081016001600160401b038111828210171561456f5761456f6146fa565b60405290565b60008219821115614588576145886146a2565b500190565b60006001600160401b038083168185168083038211156145af576145af6146a2565b01949350505050565b60006001600160601b038083168185168083038211156145af576145af6146a2565b6000826145e9576145e96146b8565b500490565b6000816000190483118215151615614608576146086146a2565b500290565b60008282101561461f5761461f6146a2565b500390565b60006001600160601b0383811690831681811015614644576146446146a2565b039392505050565b6000600019821415614660576146606146a2565b5060010190565b60006001600160401b0380831681811415614684576146846146a2565b6001019392505050565b60008261469d5761469d6146b8565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122043e0b3b2c90ce3c052a9d07e2c608c212bc6ca9279857646d17397446965357b64736f6c63430008060033",
}

// CoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CoordinatorMetaData.ABI instead.
var CoordinatorABI = CoordinatorMetaData.ABI

// CoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoordinatorMetaData.Bin instead.
var CoordinatorBin = CoordinatorMetaData.Bin

// DeployCoordinator deploys a new Ethereum contract, binding an instance of Coordinator to it.
func DeployCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, linkEthFeed common.Address) (common.Address, *types.Transaction, *Coordinator, error) {
	parsed, err := CoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoordinatorBin), backend, link, linkEthFeed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coordinator{CoordinatorCaller: CoordinatorCaller{contract: contract}, CoordinatorTransactor: CoordinatorTransactor{contract: contract}, CoordinatorFilterer: CoordinatorFilterer{contract: contract}}, nil
}

// Coordinator is an auto generated Go binding around an Ethereum contract.
type Coordinator struct {
	CoordinatorCaller     // Read-only binding to the contract
	CoordinatorTransactor // Write-only binding to the contract
	CoordinatorFilterer   // Log filterer for contract events
}

// CoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoordinatorSession struct {
	Contract     *Coordinator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoordinatorCallerSession struct {
	Contract *CoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoordinatorTransactorSession struct {
	Contract     *CoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoordinatorRaw struct {
	Contract *Coordinator // Generic contract binding to access the raw methods on
}

// CoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoordinatorCallerRaw struct {
	Contract *CoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// CoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoordinatorTransactorRaw struct {
	Contract *CoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoordinator creates a new instance of Coordinator, bound to a specific deployed contract.
func NewCoordinator(address common.Address, backend bind.ContractBackend) (*Coordinator, error) {
	contract, err := bindCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coordinator{CoordinatorCaller: CoordinatorCaller{contract: contract}, CoordinatorTransactor: CoordinatorTransactor{contract: contract}, CoordinatorFilterer: CoordinatorFilterer{contract: contract}}, nil
}

// NewCoordinatorCaller creates a new read-only instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*CoordinatorCaller, error) {
	contract, err := bindCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinatorCaller{contract: contract}, nil
}

// NewCoordinatorTransactor creates a new write-only instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CoordinatorTransactor, error) {
	contract, err := bindCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinatorTransactor{contract: contract}, nil
}

// NewCoordinatorFilterer creates a new log filterer instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CoordinatorFilterer, error) {
	contract, err := bindCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoordinatorFilterer{contract: contract}, nil
}

// bindCoordinator binds a generic wrapper to an already deployed contract.
func bindCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinator *CoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coordinator.Contract.CoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinator *CoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.Contract.CoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinator *CoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinator.Contract.CoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinator *CoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinator *CoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinator *CoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinator.Contract.contract.Transact(opts, method, params...)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_Coordinator *CoordinatorCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_Coordinator *CoordinatorSession) LINK() (common.Address, error) {
	return _Coordinator.Contract.LINK(&_Coordinator.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_Coordinator *CoordinatorCallerSession) LINK() (common.Address, error) {
	return _Coordinator.Contract.LINK(&_Coordinator.CallOpts)
}

// LINKETHFEED is a free data retrieval call binding the contract method 0xad178361.
//
// Solidity: function LINK_ETH_FEED() view returns(address)
func (_Coordinator *CoordinatorCaller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINKETHFEED is a free data retrieval call binding the contract method 0xad178361.
//
// Solidity: function LINK_ETH_FEED() view returns(address)
func (_Coordinator *CoordinatorSession) LINKETHFEED() (common.Address, error) {
	return _Coordinator.Contract.LINKETHFEED(&_Coordinator.CallOpts)
}

// LINKETHFEED is a free data retrieval call binding the contract method 0xad178361.
//
// Solidity: function LINK_ETH_FEED() view returns(address)
func (_Coordinator *CoordinatorCallerSession) LINKETHFEED() (common.Address, error) {
	return _Coordinator.Contract.LINKETHFEED(&_Coordinator.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Coordinator *CoordinatorCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Coordinator *CoordinatorSession) MAXCONSUMERS() (uint16, error) {
	return _Coordinator.Contract.MAXCONSUMERS(&_Coordinator.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Coordinator *CoordinatorCallerSession) MAXCONSUMERS() (uint16, error) {
	return _Coordinator.Contract.MAXCONSUMERS(&_Coordinator.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Coordinator *CoordinatorCaller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Coordinator *CoordinatorSession) MAXNUMWORDS() (uint32, error) {
	return _Coordinator.Contract.MAXNUMWORDS(&_Coordinator.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Coordinator *CoordinatorCallerSession) MAXNUMWORDS() (uint32, error) {
	return _Coordinator.Contract.MAXNUMWORDS(&_Coordinator.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Coordinator *CoordinatorCaller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Coordinator *CoordinatorSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _Coordinator.Contract.MAXREQUESTCONFIRMATIONS(&_Coordinator.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Coordinator *CoordinatorCallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _Coordinator.Contract.MAXREQUESTCONFIRMATIONS(&_Coordinator.CallOpts)
}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Coordinator *CoordinatorCaller) GetCommitment(opts *bind.CallOpts, requestId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getCommitment", requestId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Coordinator *CoordinatorSession) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _Coordinator.Contract.GetCommitment(&_Coordinator.CallOpts, requestId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Coordinator *CoordinatorCallerSession) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _Coordinator.Contract.GetCommitment(&_Coordinator.CallOpts, requestId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation)
func (_Coordinator *CoordinatorCaller) GetConfig(opts *bind.CallOpts) (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getConfig")

	outstruct := new(struct {
		MinimumRequestConfirmations uint16
		MaxGasLimit                 uint32
		StalenessSeconds            uint32
		GasAfterPaymentCalculation  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation)
func (_Coordinator *CoordinatorSession) GetConfig() (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}, error) {
	return _Coordinator.Contract.GetConfig(&_Coordinator.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation)
func (_Coordinator *CoordinatorCallerSession) GetConfig() (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}, error) {
	return _Coordinator.Contract.GetConfig(&_Coordinator.CallOpts)
}

// GetCurrentSubId is a free data retrieval call binding the contract method 0x06bfa637.
//
// Solidity: function getCurrentSubId() view returns(uint64)
func (_Coordinator *CoordinatorCaller) GetCurrentSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getCurrentSubId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurrentSubId is a free data retrieval call binding the contract method 0x06bfa637.
//
// Solidity: function getCurrentSubId() view returns(uint64)
func (_Coordinator *CoordinatorSession) GetCurrentSubId() (uint64, error) {
	return _Coordinator.Contract.GetCurrentSubId(&_Coordinator.CallOpts)
}

// GetCurrentSubId is a free data retrieval call binding the contract method 0x06bfa637.
//
// Solidity: function getCurrentSubId() view returns(uint64)
func (_Coordinator *CoordinatorCallerSession) GetCurrentSubId() (uint64, error) {
	return _Coordinator.Contract.GetCurrentSubId(&_Coordinator.CallOpts)
}

// GetFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x356dac71.
//
// Solidity: function getFallbackWeiPerUnitLink() view returns(int256)
func (_Coordinator *CoordinatorCaller) GetFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getFallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x356dac71.
//
// Solidity: function getFallbackWeiPerUnitLink() view returns(int256)
func (_Coordinator *CoordinatorSession) GetFallbackWeiPerUnitLink() (*big.Int, error) {
	return _Coordinator.Contract.GetFallbackWeiPerUnitLink(&_Coordinator.CallOpts)
}

// GetFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x356dac71.
//
// Solidity: function getFallbackWeiPerUnitLink() view returns(int256)
func (_Coordinator *CoordinatorCallerSession) GetFallbackWeiPerUnitLink() (*big.Int, error) {
	return _Coordinator.Contract.GetFallbackWeiPerUnitLink(&_Coordinator.CallOpts)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns(uint32 fulfillmentFlatFeeLinkPPMTier1, uint32 fulfillmentFlatFeeLinkPPMTier2, uint32 fulfillmentFlatFeeLinkPPMTier3, uint32 fulfillmentFlatFeeLinkPPMTier4, uint32 fulfillmentFlatFeeLinkPPMTier5, uint24 reqsForTier2, uint24 reqsForTier3, uint24 reqsForTier4, uint24 reqsForTier5)
func (_Coordinator *CoordinatorCaller) GetFeeConfig(opts *bind.CallOpts) (struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getFeeConfig")

	outstruct := new(struct {
		FulfillmentFlatFeeLinkPPMTier1 uint32
		FulfillmentFlatFeeLinkPPMTier2 uint32
		FulfillmentFlatFeeLinkPPMTier3 uint32
		FulfillmentFlatFeeLinkPPMTier4 uint32
		FulfillmentFlatFeeLinkPPMTier5 uint32
		ReqsForTier2                   *big.Int
		ReqsForTier3                   *big.Int
		ReqsForTier4                   *big.Int
		ReqsForTier5                   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPMTier1 = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier2 = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier3 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier4 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier5 = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ReqsForTier2 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier3 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier4 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier5 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns(uint32 fulfillmentFlatFeeLinkPPMTier1, uint32 fulfillmentFlatFeeLinkPPMTier2, uint32 fulfillmentFlatFeeLinkPPMTier3, uint32 fulfillmentFlatFeeLinkPPMTier4, uint32 fulfillmentFlatFeeLinkPPMTier5, uint24 reqsForTier2, uint24 reqsForTier3, uint24 reqsForTier4, uint24 reqsForTier5)
func (_Coordinator *CoordinatorSession) GetFeeConfig() (struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}, error) {
	return _Coordinator.Contract.GetFeeConfig(&_Coordinator.CallOpts)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x5fbbc0d2.
//
// Solidity: function getFeeConfig() view returns(uint32 fulfillmentFlatFeeLinkPPMTier1, uint32 fulfillmentFlatFeeLinkPPMTier2, uint32 fulfillmentFlatFeeLinkPPMTier3, uint32 fulfillmentFlatFeeLinkPPMTier4, uint32 fulfillmentFlatFeeLinkPPMTier5, uint24 reqsForTier2, uint24 reqsForTier3, uint24 reqsForTier4, uint24 reqsForTier5)
func (_Coordinator *CoordinatorCallerSession) GetFeeConfig() (struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}, error) {
	return _Coordinator.Contract.GetFeeConfig(&_Coordinator.CallOpts)
}

// GetFeeTier is a free data retrieval call binding the contract method 0xd2f9f9a7.
//
// Solidity: function getFeeTier(uint64 reqCount) view returns(uint32)
func (_Coordinator *CoordinatorCaller) GetFeeTier(opts *bind.CallOpts, reqCount uint64) (uint32, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getFeeTier", reqCount)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetFeeTier is a free data retrieval call binding the contract method 0xd2f9f9a7.
//
// Solidity: function getFeeTier(uint64 reqCount) view returns(uint32)
func (_Coordinator *CoordinatorSession) GetFeeTier(reqCount uint64) (uint32, error) {
	return _Coordinator.Contract.GetFeeTier(&_Coordinator.CallOpts, reqCount)
}

// GetFeeTier is a free data retrieval call binding the contract method 0xd2f9f9a7.
//
// Solidity: function getFeeTier(uint64 reqCount) view returns(uint32)
func (_Coordinator *CoordinatorCallerSession) GetFeeTier(reqCount uint64) (uint32, error) {
	return _Coordinator.Contract.GetFeeTier(&_Coordinator.CallOpts, reqCount)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Coordinator *CoordinatorCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Coordinator *CoordinatorSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _Coordinator.Contract.GetRequestConfig(&_Coordinator.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Coordinator *CoordinatorCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _Coordinator.Contract.GetRequestConfig(&_Coordinator.CallOpts)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Coordinator *CoordinatorCaller) GetSubscription(opts *bind.CallOpts, subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance   *big.Int
		ReqCount  uint64
		Owner     common.Address
		Consumers []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Coordinator *CoordinatorSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _Coordinator.Contract.GetSubscription(&_Coordinator.CallOpts, subId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Coordinator *CoordinatorCallerSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _Coordinator.Contract.GetSubscription(&_Coordinator.CallOpts, subId)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_Coordinator *CoordinatorCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_Coordinator *CoordinatorSession) GetTotalBalance() (*big.Int, error) {
	return _Coordinator.Contract.GetTotalBalance(&_Coordinator.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_Coordinator *CoordinatorCallerSession) GetTotalBalance() (*big.Int, error) {
	return _Coordinator.Contract.GetTotalBalance(&_Coordinator.CallOpts)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Coordinator *CoordinatorCaller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Coordinator *CoordinatorSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _Coordinator.Contract.HashOfKey(&_Coordinator.CallOpts, publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Coordinator *CoordinatorCallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _Coordinator.Contract.HashOfKey(&_Coordinator.CallOpts, publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coordinator *CoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coordinator *CoordinatorSession) Owner() (common.Address, error) {
	return _Coordinator.Contract.Owner(&_Coordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coordinator *CoordinatorCallerSession) Owner() (common.Address, error) {
	return _Coordinator.Contract.Owner(&_Coordinator.CallOpts)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Coordinator *CoordinatorCaller) PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Coordinator *CoordinatorSession) PendingRequestExists(subId uint64) (bool, error) {
	return _Coordinator.Contract.PendingRequestExists(&_Coordinator.CallOpts, subId)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Coordinator *CoordinatorCallerSession) PendingRequestExists(subId uint64) (bool, error) {
	return _Coordinator.Contract.PendingRequestExists(&_Coordinator.CallOpts, subId)
}

// SSubscriptionConfigs is a free data retrieval call binding the contract method 0xd7ca6f9b.
//
// Solidity: function s_subscriptionConfigs(uint64 ) view returns(address owner, address requestedOwner)
func (_Coordinator *CoordinatorCaller) SSubscriptionConfigs(opts *bind.CallOpts, arg0 uint64) (struct {
	Owner          common.Address
	RequestedOwner common.Address
}, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "s_subscriptionConfigs", arg0)

	outstruct := new(struct {
		Owner          common.Address
		RequestedOwner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RequestedOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SSubscriptionConfigs is a free data retrieval call binding the contract method 0xd7ca6f9b.
//
// Solidity: function s_subscriptionConfigs(uint64 ) view returns(address owner, address requestedOwner)
func (_Coordinator *CoordinatorSession) SSubscriptionConfigs(arg0 uint64) (struct {
	Owner          common.Address
	RequestedOwner common.Address
}, error) {
	return _Coordinator.Contract.SSubscriptionConfigs(&_Coordinator.CallOpts, arg0)
}

// SSubscriptionConfigs is a free data retrieval call binding the contract method 0xd7ca6f9b.
//
// Solidity: function s_subscriptionConfigs(uint64 ) view returns(address owner, address requestedOwner)
func (_Coordinator *CoordinatorCallerSession) SSubscriptionConfigs(arg0 uint64) (struct {
	Owner          common.Address
	RequestedOwner common.Address
}, error) {
	return _Coordinator.Contract.SSubscriptionConfigs(&_Coordinator.CallOpts, arg0)
}

// SSubscriptions is a free data retrieval call binding the contract method 0x6840c05e.
//
// Solidity: function s_subscriptions(uint64 ) view returns(uint96 balance, uint64 reqCount)
func (_Coordinator *CoordinatorCaller) SSubscriptions(opts *bind.CallOpts, arg0 uint64) (struct {
	Balance  *big.Int
	ReqCount uint64
}, error) {
	var out []interface{}
	err := _Coordinator.contract.Call(opts, &out, "s_subscriptions", arg0)

	outstruct := new(struct {
		Balance  *big.Int
		ReqCount uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// SSubscriptions is a free data retrieval call binding the contract method 0x6840c05e.
//
// Solidity: function s_subscriptions(uint64 ) view returns(uint96 balance, uint64 reqCount)
func (_Coordinator *CoordinatorSession) SSubscriptions(arg0 uint64) (struct {
	Balance  *big.Int
	ReqCount uint64
}, error) {
	return _Coordinator.Contract.SSubscriptions(&_Coordinator.CallOpts, arg0)
}

// SSubscriptions is a free data retrieval call binding the contract method 0x6840c05e.
//
// Solidity: function s_subscriptions(uint64 ) view returns(uint96 balance, uint64 reqCount)
func (_Coordinator *CoordinatorCallerSession) SSubscriptions(arg0 uint64) (struct {
	Balance  *big.Int
	ReqCount uint64
}, error) {
	return _Coordinator.Contract.SSubscriptions(&_Coordinator.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coordinator *CoordinatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coordinator *CoordinatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coordinator.Contract.AcceptOwnership(&_Coordinator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coordinator *CoordinatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coordinator.Contract.AcceptOwnership(&_Coordinator.TransactOpts)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Coordinator *CoordinatorTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Coordinator *CoordinatorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _Coordinator.Contract.AcceptSubscriptionOwnerTransfer(&_Coordinator.TransactOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Coordinator *CoordinatorTransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _Coordinator.Contract.AcceptSubscriptionOwnerTransfer(&_Coordinator.TransactOpts, subId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorTransactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "addConsumer", subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.AddConsumer(&_Coordinator.TransactOpts, subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorTransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.AddConsumer(&_Coordinator.TransactOpts, subId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Coordinator *CoordinatorTransactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "cancelSubscription", subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Coordinator *CoordinatorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.CancelSubscription(&_Coordinator.TransactOpts, subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Coordinator *CoordinatorTransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.CancelSubscription(&_Coordinator.TransactOpts, subId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_Coordinator *CoordinatorTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_Coordinator *CoordinatorSession) CreateSubscription() (*types.Transaction, error) {
	return _Coordinator.Contract.CreateSubscription(&_Coordinator.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_Coordinator *CoordinatorTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _Coordinator.Contract.CreateSubscription(&_Coordinator.TransactOpts)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorTransactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.DeregisterProvingKey(&_Coordinator.TransactOpts, publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorTransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.DeregisterProvingKey(&_Coordinator.TransactOpts, publicProvingKey)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Coordinator *CoordinatorTransactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Coordinator *CoordinatorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _Coordinator.Contract.FulfillRandomWords(&_Coordinator.TransactOpts, proof, rc)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Coordinator *CoordinatorTransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _Coordinator.Contract.FulfillRandomWords(&_Coordinator.TransactOpts, proof, rc)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_Coordinator *CoordinatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_Coordinator *CoordinatorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Coordinator.Contract.OnTokenTransfer(&_Coordinator.TransactOpts, arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_Coordinator *CoordinatorTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Coordinator.Contract.OnTokenTransfer(&_Coordinator.TransactOpts, arg0, amount, data)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_Coordinator *CoordinatorTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_Coordinator *CoordinatorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.OracleWithdraw(&_Coordinator.TransactOpts, recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_Coordinator *CoordinatorTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.OracleWithdraw(&_Coordinator.TransactOpts, recipient, amount)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subId) returns()
func (_Coordinator *CoordinatorTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "ownerCancelSubscription", subId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subId) returns()
func (_Coordinator *CoordinatorSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _Coordinator.Contract.OwnerCancelSubscription(&_Coordinator.TransactOpts, subId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subId) returns()
func (_Coordinator *CoordinatorTransactorSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _Coordinator.Contract.OwnerCancelSubscription(&_Coordinator.TransactOpts, subId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_Coordinator *CoordinatorTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "recoverFunds", to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_Coordinator *CoordinatorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RecoverFunds(&_Coordinator.TransactOpts, to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_Coordinator *CoordinatorTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RecoverFunds(&_Coordinator.TransactOpts, to)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x6f64f03f.
//
// Solidity: function registerProvingKey(address oracle, uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorTransactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x6f64f03f.
//
// Solidity: function registerProvingKey(address oracle, uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.RegisterProvingKey(&_Coordinator.TransactOpts, oracle, publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x6f64f03f.
//
// Solidity: function registerProvingKey(address oracle, uint256[2] publicProvingKey) returns()
func (_Coordinator *CoordinatorTransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.RegisterProvingKey(&_Coordinator.TransactOpts, oracle, publicProvingKey)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorTransactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "removeConsumer", subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RemoveConsumer(&_Coordinator.TransactOpts, subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Coordinator *CoordinatorTransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RemoveConsumer(&_Coordinator.TransactOpts, subId, consumer)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Coordinator *CoordinatorTransactor) RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "requestRandomWords", keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Coordinator *CoordinatorSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Coordinator.Contract.RequestRandomWords(&_Coordinator.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Coordinator *CoordinatorTransactorSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Coordinator.Contract.RequestRandomWords(&_Coordinator.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Coordinator *CoordinatorTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Coordinator *CoordinatorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RequestSubscriptionOwnerTransfer(&_Coordinator.TransactOpts, subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Coordinator *CoordinatorTransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.RequestSubscriptionOwnerTransfer(&_Coordinator.TransactOpts, subId, newOwner)
}

// SetConfig is a paid mutator transaction binding the contract method 0x4cb48a54.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig) returns()
func (_Coordinator *CoordinatorTransactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0x4cb48a54.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig) returns()
func (_Coordinator *CoordinatorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _Coordinator.Contract.SetConfig(&_Coordinator.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0x4cb48a54.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig) returns()
func (_Coordinator *CoordinatorTransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _Coordinator.Contract.SetConfig(&_Coordinator.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Coordinator *CoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Coordinator *CoordinatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.TransferOwnership(&_Coordinator.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Coordinator *CoordinatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.TransferOwnership(&_Coordinator.TransactOpts, to)
}

// CoordinatorConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the Coordinator contract.
type CoordinatorConfigSetIterator struct {
	Event *CoordinatorConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorConfigSet represents a ConfigSet event raised by the Coordinator contract.
type CoordinatorConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV2FeeConfig
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0xc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb2.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig)
func (_Coordinator *CoordinatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CoordinatorConfigSetIterator, error) {

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CoordinatorConfigSetIterator{contract: _Coordinator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0xc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb2.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig)
func (_Coordinator *CoordinatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorConfigSet)
				if err := _Coordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigSet is a log parse operation binding the contract event 0xc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb2.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, (uint32,uint32,uint32,uint32,uint32,uint24,uint24,uint24,uint24) feeConfig)
func (_Coordinator *CoordinatorFilterer) ParseConfigSet(log types.Log) (*CoordinatorConfigSet, error) {
	event := new(CoordinatorConfigSet)
	if err := _Coordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorFundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the Coordinator contract.
type CoordinatorFundsRecoveredIterator struct {
	Event *CoordinatorFundsRecovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorFundsRecovered represents a FundsRecovered event raised by the Coordinator contract.
type CoordinatorFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*CoordinatorFundsRecoveredIterator, error) {

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &CoordinatorFundsRecoveredIterator{contract: _Coordinator.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *CoordinatorFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorFundsRecovered)
				if err := _Coordinator.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsRecovered is a log parse operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) ParseFundsRecovered(log types.Log) (*CoordinatorFundsRecovered, error) {
	event := new(CoordinatorFundsRecovered)
	if err := _Coordinator.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Coordinator contract.
type CoordinatorOwnershipTransferRequestedIterator struct {
	Event *CoordinatorOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Coordinator contract.
type CoordinatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CoordinatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorOwnershipTransferRequestedIterator{contract: _Coordinator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorOwnershipTransferRequested)
				if err := _Coordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*CoordinatorOwnershipTransferRequested, error) {
	event := new(CoordinatorOwnershipTransferRequested)
	if err := _Coordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Coordinator contract.
type CoordinatorOwnershipTransferredIterator struct {
	Event *CoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the Coordinator contract.
type CoordinatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CoordinatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorOwnershipTransferredIterator{contract: _Coordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorOwnershipTransferred)
				if err := _Coordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Coordinator *CoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*CoordinatorOwnershipTransferred, error) {
	event := new(CoordinatorOwnershipTransferred)
	if err := _Coordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorProvingKeyDeregisteredIterator is returned from FilterProvingKeyDeregistered and is used to iterate over the raw logs and unpacked data for ProvingKeyDeregistered events raised by the Coordinator contract.
type CoordinatorProvingKeyDeregisteredIterator struct {
	Event *CoordinatorProvingKeyDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorProvingKeyDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorProvingKeyDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorProvingKeyDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorProvingKeyDeregistered represents a ProvingKeyDeregistered event raised by the Coordinator contract.
type CoordinatorProvingKeyDeregistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyDeregistered is a free log retrieval operation binding the contract event 0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*CoordinatorProvingKeyDeregisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorProvingKeyDeregisteredIterator{contract: _Coordinator.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyDeregistered is a free log subscription operation binding the contract event 0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *CoordinatorProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorProvingKeyDeregistered)
				if err := _Coordinator.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingKeyDeregistered is a log parse operation binding the contract event 0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) ParseProvingKeyDeregistered(log types.Log) (*CoordinatorProvingKeyDeregistered, error) {
	event := new(CoordinatorProvingKeyDeregistered)
	if err := _Coordinator.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorProvingKeyRegisteredIterator is returned from FilterProvingKeyRegistered and is used to iterate over the raw logs and unpacked data for ProvingKeyRegistered events raised by the Coordinator contract.
type CoordinatorProvingKeyRegisteredIterator struct {
	Event *CoordinatorProvingKeyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorProvingKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorProvingKeyRegistered represents a ProvingKeyRegistered event raised by the Coordinator contract.
type CoordinatorProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyRegistered is a free log retrieval operation binding the contract event 0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*CoordinatorProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorProvingKeyRegisteredIterator{contract: _Coordinator.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyRegistered is a free log subscription operation binding the contract event 0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *CoordinatorProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorProvingKeyRegistered)
				if err := _Coordinator.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingKeyRegistered is a log parse operation binding the contract event 0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, address indexed oracle)
func (_Coordinator *CoordinatorFilterer) ParseProvingKeyRegistered(log types.Log) (*CoordinatorProvingKeyRegistered, error) {
	event := new(CoordinatorProvingKeyRegistered)
	if err := _Coordinator.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorRandomWordsFulfilledIterator is returned from FilterRandomWordsFulfilled and is used to iterate over the raw logs and unpacked data for RandomWordsFulfilled events raised by the Coordinator contract.
type CoordinatorRandomWordsFulfilledIterator struct {
	Event *CoordinatorRandomWordsFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorRandomWordsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorRandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorRandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorRandomWordsFulfilled represents a RandomWordsFulfilled event raised by the Coordinator contract.
type CoordinatorRandomWordsFulfilled struct {
	RequestId  *big.Int
	OutputSeed *big.Int
	Payment    *big.Int
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsFulfilled is a free log retrieval operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Coordinator *CoordinatorFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*CoordinatorRandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorRandomWordsFulfilledIterator{contract: _Coordinator.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomWordsFulfilled is a free log subscription operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Coordinator *CoordinatorFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *CoordinatorRandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorRandomWordsFulfilled)
				if err := _Coordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsFulfilled is a log parse operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Coordinator *CoordinatorFilterer) ParseRandomWordsFulfilled(log types.Log) (*CoordinatorRandomWordsFulfilled, error) {
	event := new(CoordinatorRandomWordsFulfilled)
	if err := _Coordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorRandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the Coordinator contract.
type CoordinatorRandomWordsRequestedIterator struct {
	Event *CoordinatorRandomWordsRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorRandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorRandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorRandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorRandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorRandomWordsRequested represents a RandomWordsRequested event raised by the Coordinator contract.
type CoordinatorRandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       uint64
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	Sender                      common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Coordinator *CoordinatorFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*CoordinatorRandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorRandomWordsRequestedIterator{contract: _Coordinator.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Coordinator *CoordinatorFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *CoordinatorRandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorRandomWordsRequested)
				if err := _Coordinator.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsRequested is a log parse operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Coordinator *CoordinatorFilterer) ParseRandomWordsRequested(log types.Log) (*CoordinatorRandomWordsRequested, error) {
	event := new(CoordinatorRandomWordsRequested)
	if err := _Coordinator.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionCanceledIterator is returned from FilterSubscriptionCanceled and is used to iterate over the raw logs and unpacked data for SubscriptionCanceled events raised by the Coordinator contract.
type CoordinatorSubscriptionCanceledIterator struct {
	Event *CoordinatorSubscriptionCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionCanceled represents a SubscriptionCanceled event raised by the Coordinator contract.
type CoordinatorSubscriptionCanceled struct {
	SubId  uint64
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCanceled is a free log retrieval operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionCanceledIterator{contract: _Coordinator.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCanceled is a free log subscription operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionCanceled, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionCanceled)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionCanceled is a log parse operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionCanceled(log types.Log) (*CoordinatorSubscriptionCanceled, error) {
	event := new(CoordinatorSubscriptionCanceled)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionConsumerAddedIterator is returned from FilterSubscriptionConsumerAdded and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerAdded events raised by the Coordinator contract.
type CoordinatorSubscriptionConsumerAddedIterator struct {
	Event *CoordinatorSubscriptionConsumerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionConsumerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionConsumerAdded represents a SubscriptionConsumerAdded event raised by the Coordinator contract.
type CoordinatorSubscriptionConsumerAdded struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerAdded is a free log retrieval operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionConsumerAddedIterator{contract: _Coordinator.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerAdded is a free log subscription operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionConsumerAdded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionConsumerAdded)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionConsumerAdded is a log parse operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*CoordinatorSubscriptionConsumerAdded, error) {
	event := new(CoordinatorSubscriptionConsumerAdded)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionConsumerRemovedIterator is returned from FilterSubscriptionConsumerRemoved and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerRemoved events raised by the Coordinator contract.
type CoordinatorSubscriptionConsumerRemovedIterator struct {
	Event *CoordinatorSubscriptionConsumerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionConsumerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionConsumerRemoved represents a SubscriptionConsumerRemoved event raised by the Coordinator contract.
type CoordinatorSubscriptionConsumerRemoved struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerRemoved is a free log retrieval operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionConsumerRemovedIterator{contract: _Coordinator.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerRemoved is a free log subscription operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionConsumerRemoved)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionConsumerRemoved is a log parse operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*CoordinatorSubscriptionConsumerRemoved, error) {
	event := new(CoordinatorSubscriptionConsumerRemoved)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionCreatedIterator is returned from FilterSubscriptionCreated and is used to iterate over the raw logs and unpacked data for SubscriptionCreated events raised by the Coordinator contract.
type CoordinatorSubscriptionCreatedIterator struct {
	Event *CoordinatorSubscriptionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionCreated represents a SubscriptionCreated event raised by the Coordinator contract.
type CoordinatorSubscriptionCreated struct {
	SubId uint64
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCreated is a free log retrieval operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionCreatedIterator{contract: _Coordinator.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCreated is a free log subscription operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionCreated, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionCreated)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionCreated is a log parse operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionCreated(log types.Log) (*CoordinatorSubscriptionCreated, error) {
	event := new(CoordinatorSubscriptionCreated)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionFundedIterator is returned from FilterSubscriptionFunded and is used to iterate over the raw logs and unpacked data for SubscriptionFunded events raised by the Coordinator contract.
type CoordinatorSubscriptionFundedIterator struct {
	Event *CoordinatorSubscriptionFunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionFunded represents a SubscriptionFunded event raised by the Coordinator contract.
type CoordinatorSubscriptionFunded struct {
	SubId      uint64
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionFunded is a free log retrieval operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionFundedIterator{contract: _Coordinator.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionFunded is a free log subscription operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionFunded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionFunded)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionFunded is a log parse operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionFunded(log types.Log) (*CoordinatorSubscriptionFunded, error) {
	event := new(CoordinatorSubscriptionFunded)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionOwnerTransferRequestedIterator is returned from FilterSubscriptionOwnerTransferRequested and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferRequested events raised by the Coordinator contract.
type CoordinatorSubscriptionOwnerTransferRequestedIterator struct {
	Event *CoordinatorSubscriptionOwnerTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionOwnerTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionOwnerTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionOwnerTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionOwnerTransferRequested represents a SubscriptionOwnerTransferRequested event raised by the Coordinator contract.
type CoordinatorSubscriptionOwnerTransferRequested struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferRequested is a free log retrieval operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionOwnerTransferRequestedIterator{contract: _Coordinator.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferRequested is a free log subscription operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionOwnerTransferRequested)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionOwnerTransferRequested is a log parse operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*CoordinatorSubscriptionOwnerTransferRequested, error) {
	event := new(CoordinatorSubscriptionOwnerTransferRequested)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoordinatorSubscriptionOwnerTransferredIterator is returned from FilterSubscriptionOwnerTransferred and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferred events raised by the Coordinator contract.
type CoordinatorSubscriptionOwnerTransferredIterator struct {
	Event *CoordinatorSubscriptionOwnerTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoordinatorSubscriptionOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoordinatorSubscriptionOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoordinatorSubscriptionOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoordinatorSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoordinatorSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoordinatorSubscriptionOwnerTransferred represents a SubscriptionOwnerTransferred event raised by the Coordinator contract.
type CoordinatorSubscriptionOwnerTransferred struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferred is a free log retrieval operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*CoordinatorSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &CoordinatorSubscriptionOwnerTransferredIterator{contract: _Coordinator.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferred is a free log subscription operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *CoordinatorSubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Coordinator.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoordinatorSubscriptionOwnerTransferred)
				if err := _Coordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionOwnerTransferred is a log parse operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subId, address from, address to)
func (_Coordinator *CoordinatorFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*CoordinatorSubscriptionOwnerTransferred, error) {
	event := new(CoordinatorSubscriptionOwnerTransferred)
	if err := _Coordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
