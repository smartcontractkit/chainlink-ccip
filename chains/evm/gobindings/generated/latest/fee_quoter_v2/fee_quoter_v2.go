// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_quoter_v2

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type FeeQuoterDestChainConfig struct {
	IsEnabled                   bool
	MaxDataBytes                uint32
	MaxPerMsgGasLimit           uint32
	DestGasOverhead             uint32
	DestGasPerPayloadByteBase   uint8
	ChainFamilySelector         [4]byte
	DefaultTokenFeeUSDCents     uint16
	DefaultTokenDestGasOverhead uint32
	DefaultTxGasLimit           uint32
	NetworkFeeUSDCents          uint32
}

type FeeQuoterDestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
}

type FeeQuoterPremiumMultiplierWeiPerEthArgs struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
}

type FeeQuoterStaticConfig struct {
	MaxFeeJuelsPerMsg *big.Int
	LinkToken         common.Address
}

type FeeQuoterTokenTransferFeeConfig struct {
	FeeUSDCents       uint32
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	IsEnabled         bool
}

type FeeQuoterTokenTransferFeeConfigArgs struct {
	DestChainSelector       uint64
	TokenTransferFeeConfigs []FeeQuoterTokenTransferFeeConfigSingleTokenArgs
}

type FeeQuoterTokenTransferFeeConfigRemoveArgs struct {
	DestChainSelector uint64
	Token             common.Address
}

type FeeQuoterTokenTransferFeeConfigSingleTokenArgs struct {
	Token                  common.Address
	TokenTransferFeeConfig FeeQuoterTokenTransferFeeConfig
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalTimestampedPackedUint224 struct {
	Value     *big.Int
	Timestamp uint32
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

var FeeQuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveTokenReceiver\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c060405234610be45761643f8038038061001981610d9a565b928339810181810360e08112610be457604013610be457610038610d7b565b82519091906001600160601b0381168103610be457825261005b60208401610dbf565b6020830190815260408401516001600160401b038111610be45782610081918601610dea565b60608501519093906001600160401b038111610be457836100a3918701610dea565b60808601519093906001600160401b038111610be45786019581601f88011215610be4578651966100db6100d689610dd3565b610d9a565b976020808a838152019160051b83010191848311610be45760208101915b838310610c3b575050505060a08101516001600160401b038111610be45781019082601f83011215610be45781516101336100d682610dd3565b9260208085848152019260061b82010190858211610be457602001915b818310610bff5750505060c0810151906001600160401b038211610be4570182601f82011215610be4578051906101896100d683610dd3565b936020610160818786815201940283010191818311610be457602001925b828410610ac657505050503315610ab557600180546001600160a01b031916331790556020966101d688610d9a565b956000875260003681376101e8610d7b565b978852868989015260005b875181101561025a576001906001600160a01b03610211828b610e77565b51168b61021d82611063565b61022a575b5050016101f3565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388b610222565b5088965087519460005b86518110156102d5576001600160a01b0361027f8289610e77565b51169081156102c4577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8a836102b6600195610feb565b50604051908152a101610264565b6342bcdf7f60e11b60005260046000fd5b508651949550879487906001600160a01b0316158015610aa3575b610a9257516001600160a01b031660a052516001600160601b031660805261031784610d9a565b9260008452600036813760005b84518110156103935760019061034c6001600160a01b036103458389610e77565b5116610ef8565b610357575b01610324565b818060a01b036103678288610e77565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2610351565b508593508460005b835181101561040a576001906103c36001600160a01b036103bc8388610e77565b511661102a565b6103ce575b0161039b565b818060a01b036103de8287610e77565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a26103c8565b5091509160005b84518110156107c5576104248186610e77565b51836001600160401b036104388489610e77565b51511691015190801580156107b2575b8015610794575b8015610713575b6106ff57600081815260078652604090205460019392919060701b6001600160e01b03191661064457807f6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b60405180610532868291909161012063ffffffff8161014084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff60808201511660808601528260e01b60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e08601528261010082015116610100860152015116910152565b0390a25b6000908152600786526040908190208251815484890151938501516060860151608087015160a08089015160c0808b015160e0808d01516101008e0151610120909e01516001600160e01b03199a8b1660ff9c15159c909c169b909b1760089d909d1b64ffffffff00169c909c1760289890981b68ffffffff0000000000169790971760489690961b6cffffffff000000000000000000169590951760689490941b6dff00000000000000000000000000169390931760709190911c63ffffffff60701b161760909390931b61ffff60901b16929092178a841b8b9003169690911b63ffffffff60a01b16959095179590941b63ffffffff60c01b1694909417921b90921617905501610411565b807f8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad604051806106f7868291909161012063ffffffff8161014084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff60808201511660808601528260e01b60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e08601528261010082015116610100860152015116910152565b0390a2610536565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b60a083015116630a04b54b60e21b8114159081610782575b81610770575b8161075e575b8161074c575b50610456565b63647e2ba960e01b1415905088610746565b63c4e0595360e01b8114159150610740565b632b1dfffb60e21b811415915061073a565b6307842f7160e21b8114159150610734565b5063ffffffff6101008301511663ffffffff6040840151161061044f565b5063ffffffff6101008301511615610448565b508260005b815181101561084a576001906001600160a01b036107e88285610e77565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b038161081e8689610e77565b510151168360005260068252604060002081878060401b0319825416179055604051908152a2016107ca565b5050600161085783610d9a565b9160008352600091610a8d575b81925b81518410156109db5761087a8483610e77565b5180516001600160401b0316929086019190845b87845180518310156109ca57826108a491610e77565b51015184516001600160a01b03906108bd908490610e77565b51511690604081019063ffffffff8251168b81106109b3575091877f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed160808d97969460019663ffffffff60408f87815260088d528181208b8060a01b038a1682528d522092818c8185511695805467ffffffff00000000838801938451901b169660606bffffffff0000000000000000875160401b169101976cff0000000000000000000000008951151560601b16928a6cff000000000000000000000000199160018060601b0319161716171717905560405195865251168c850152511660408301525115156060820152a30190915061088e565b6312766e0160e11b8a526004849052602452604489fd5b505050925093600191500192610867565b905083825b8251811015610a5e576001906001600160401b036109fe8286610e77565b515116828060a01b0384610a128488610e77565b51015116908087526008855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a3016109e0565b60405161534790816110f882396080518181816105270152611026015260a05181818161054f0152610fbd0152f35b610864565b63d794ef9560e01b60005260046000fd5b5081516001600160601b0316156102f0565b639b15e16f60e01b60005260046000fd5b8382036101608112610be457610140610add610d7b565b91610ae787610e45565b8352601f190112610be4576040519161014083016001600160401b03811184821017610be957604052610b1c60208701610e6a565b8352610b2a60408701610e59565b6020840152610b3b60608701610e59565b6040840152610b4c60808701610e59565b606084015260a086015160ff81168103610be457608084015260c08601516001600160e01b031981168103610be45760a084015260e08601519161ffff83168303610be4578360209360c0610160960152610baa6101008901610e59565b60e0820152610bbc6101208901610e59565b610100820152610bcf6101408901610e59565b610120820152838201528152019301926101a7565b600080fd5b634e487b7160e01b600052604160045260246000fd5b604083870312610be4576020604091610c16610d7b565b610c1f86610dbf565b8152610c2c838701610e45565b83820152815201920191610150565b82516001600160401b038111610be45782016040818803601f190112610be457610c63610d7b565b90610c7060208201610e45565b825260408101516001600160401b038111610be457602091010187601f82011215610be4578051610ca36100d682610dd3565b91602060a08185858152019302820101908a8211610be457602001915b818310610cdf57505050918160209384809401528152019201916100f9565b828b0360a08112610be4576080610cf4610d7b565b91610cfe86610dbf565b8352601f190112610be4576040519160808301916001600160401b03831184841017610be95760a093602093604052610d38848801610e59565b8152610d4660408801610e59565b84820152610d5660608801610e59565b6040820152610d6760808801610e6a565b606082015283820152815201920191610cc0565b60408051919082016001600160401b03811183821017610be957604052565b6040519190601f01601f191682016001600160401b03811183821017610be957604052565b51906001600160a01b0382168203610be457565b6001600160401b038111610be95760051b60200190565b9080601f83011215610be4578151610e046100d682610dd3565b9260208085848152019260051b820101928311610be457602001905b828210610e2d5750505090565b60208091610e3a84610dbf565b815201910190610e20565b51906001600160401b0382168203610be457565b519063ffffffff82168203610be457565b51908115158203610be457565b8051821015610e8b5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610e8b5760005260206000200190600090565b80548015610ee2576000190190610ed08282610ea1565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600a60205260409020548015610fb9576000198101818111610fa357600954600019810191908211610fa357818103610f52575b505050610f3e6009610eb9565b600052600a60205260006040812055600190565b610f8b610f63610f74936009610ea1565b90549060031b1c9283926009610ea1565b819391549060031b91821b91600019901b19161790565b9055600052600a602052604060002055388080610f31565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80549068010000000000000000821015610be95781610f74916001610fe794018155610ea1565b9055565b806000526003602052604060002054156000146110245761100d816002610fc0565b600254906000526003602052604060002055600190565b50600090565b80600052600a602052604060002054156000146110245761104c816009610fc0565b60095490600052600a602052604060002055600190565b6000818152600360205260409020548015610fb9576000198101818111610fa357600254600019810191908211610fa3578082036110bd575b5050506110a96002610eb9565b600052600360205260006040812055600190565b6110df6110ce610f74936002610ea1565b90549060031b1c9283926002610ea1565b9055600052600360205260406000205538808061109c56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be146101b657806301447eaa146101b1578063061877e3146101ac57806306285c69146101a757806315c34d5b146101a2578063181f5a771461019d5780632451a627146101985780633937306f146101935780633a49bb491461018e57806345ac924d146101895780634ab35b0b14610184578063514e8cff1461017f5780636def4ce71461017a57806379ba5097146101755780637afac3221461017057806382b49eb01461016b5780638da5cb5b1461016657806391a2749a146101615780639b1115e41461015c578063a69c64c014610157578063bce21f2214610152578063cdc73d511461014d578063d02641a014610148578063d8694ccd14610143578063f2fde38b1461013e5763ffdb4b371461013957600080fd5b61234b565b612257565b611ef8565b611e7c565b611de7565b611c3a565b611b19565b611aaa565b6119da565b611988565b61185a565b6116e9565b611599565b611421565b611290565b611225565b611120565b610f2d565b610b37565b610a96565b6109c9565b610754565b6104bc565b610449565b61038c565b6101de565b73ffffffffffffffffffffffffffffffffffffffff8116036101d957565b600080fd5b346101d95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957602061023360043561021e816101bb565b6024356044359161022e836101bb565b612532565b604051908152f35b6004359067ffffffffffffffff821682036101d957565b6024359067ffffffffffffffff821682036101d957565b359067ffffffffffffffff821682036101d957565b9181601f840112156101d95782359167ffffffffffffffff83116101d9576020808501948460051b0101116101d957565b919082519283825260005b8481106102f95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016102ba565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061034157505050505090565b909192939460208061037d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102af565b97019301930191939290610332565b346101d95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d9576103c361023b565b60243567ffffffffffffffff81116101d9576103e390369060040161027e565b6044929192359167ffffffffffffffff83116101d957366023840112156101d95782600401359167ffffffffffffffff83116101d9573660248460061b860101116101d95761044594602461043995019261274e565b6040519182918261030e565b0390f35b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95773ffffffffffffffffffffffffffffffffffffffff600435610499816101bb565b166000526006602052602067ffffffffffffffff60406000205416604051908152f35b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d9576104f36129b5565b5060408051610501816105ae565b73ffffffffffffffffffffffffffffffffffffffff60206bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283815201817f0000000000000000000000000000000000000000000000000000000000000000168152835192835251166020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176105ca57604052565b61057f565b6080810190811067ffffffffffffffff8211176105ca57604052565b610140810190811067ffffffffffffffff8211176105ca57604052565b60a0810190811067ffffffffffffffff8211176105ca57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176105ca57604052565b60405190610674604083610624565b565b6040519061067461014083610624565b67ffffffffffffffff81116105ca5760051b60200190565b359063ffffffff821682036101d957565b801515036101d957565b3590610674826106af565b81601f820112156101d9578035906106db82610686565b926106e96040519485610624565b82845260208085019360061b830101918183116101d957602001925b828410610713575050505090565b6040848303126101d9576020604091825161072d816105ae565b61073687610269565b815282870135610745816101bb565b83820152815201930192610705565b346101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957366023820112156101d95780600401356107ae81610686565b916107bc6040519384610624565b8183526024602084019260051b820101903682116101d95760248101925b82841061080d576024358567ffffffffffffffff82116101d95761080561080b9236906004016106c4565b906129ce565b005b833567ffffffffffffffff81116101d957820160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc82360301126101d95760405190610859826105ae565b61086560248201610269565b8252604481013567ffffffffffffffff81116101d957602491010136601f820112156101d957803561089681610686565b916108a46040519384610624565b818352602060a08185019302820101903682116101d957602001915b8183106108df57505050918160209384809401528152019301926107da565b82360360a081126101d95760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06040519261091a846105ae565b8635610925816101bb565b845201126101d95760a09160209160405161093f816105cf565b61094a84880161069e565b81526109586040880161069e565b848201526109686060880161069e565b6040820152608087013561097b816106af565b6060820152838201528152019201916108c0565b67ffffffffffffffff81116105ca57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d9576104456040805190610a0a8183610624565b601382527f46656551756f74657220312e362e332d646576000000000000000000000000006020830152519182916020835260208301906102af565b602060408183019282815284518094520192019060005b818110610a6a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610a5d565b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610b215761044585610b1581870382610624565b60405191829182610a46565b8254845260209093019260019283019201610afe565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957806004019060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101d957610bb1613a72565b610bbb8280612cdd565b4263ffffffff1692915060005b818110610d9657505060240190610bdf8284612cdd565b92905060005b838110610bee57005b80610c0d610c08600193610c02868a612cdd565b9061260e565b612d91565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff610d5d610d3a6020850194610d2c610c6b87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610c9a610c76610665565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092168252565b63ffffffff8c166020820152610cd5610cbb845167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b81516020929092015160e01b7fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5167ffffffffffffffff1690565b93517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610be5565b80610daf610daa600193610c028980612cdd565b612d5a565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a73ffffffffffffffffffffffffffffffffffffffff610e91610d3a6020850194610e77610e1987517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610e24610c76610665565b63ffffffff8d166020820152610cd5610e51845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b5173ffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610bc8565b9181601f840112156101d95782359167ffffffffffffffff83116101d957602083818601950101116101d957565b92610f2a9492610f1c928552151560208501526080604085015260808401906102af565b9160608184039101526102af565b90565b346101d95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957610f6461023b565b60243590610f71826101bb565b6044359160643567ffffffffffffffff81116101d957610f95903690600401610eca565b93909160843567ffffffffffffffff81116101d957610fb8903690600401610eca565b9290917f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff82161460001461109e575050935b6bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680861161106d57509161105e93916104459693613ab6565b90939160405194859485610ef8565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b916110a892612532565b93611017565b602060408183019282815284518094520192019060005b8181106110d25750505090565b9091926020604082611115600194885163ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b0194019291016110c5565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d95761116f90369060040161027e565b61117881610686565b916111866040519384610624565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06111b383610686565b0160005b81811061120e57505060005b82811015611200576001906111e46111df8260051b8501612623565b6135c2565b6111ee828761273a565b526111f9818661273a565b50016111c3565b6040518061044586826110ae565b6020906112196129b5565b828288010152016111b7565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957602061126a600435611265816101bb565b613692565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95767ffffffffffffffff6112d061023b565b6112d86129b5565b501660005260046020526040600020604051906112f4826105ae565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c6020820152604051809161044582604081019263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b610674909291926101208061014083019561137c84825115159052565b60208181015163ffffffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015160ff169085015260a0818101517fffffffff00000000000000000000000000000000000000000000000000000000169085015260c08181015161ffff169085015260e08181015163ffffffff16908501526101008181015163ffffffff1690850152015163ffffffff16910152565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95767ffffffffffffffff61146161023b565b6000610120604051611472816105eb565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201520152166000526007602052610445604060002061158d61157f604051926114cc846105eb565b5460ff8116151584526114f1600882901c63ffffffff165b63ffffffff166020860152565b63ffffffff602882901c16604085015263ffffffff604882901c16606085015260ff606882901c1660808501527fffffffff00000000000000000000000000000000000000000000000000000000607082901b1660a085015261ffff609082901c1660c085015263ffffffff60a082901c1660e085015263ffffffff60c082901c1661010085015260e01c90565b63ffffffff16610120830152565b6040519182918261135f565b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760005473ffffffffffffffffffffffffffffffffffffffff81163303611658577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f830112156101d957813561169981610686565b926116a76040519485610624565b81845260208085019260051b8201019283116101d957602001905b8282106116cf5750505090565b6020809183356116de816101bb565b8152019101906116c2565b346101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957611738903690600401611682565b60243567ffffffffffffffff81116101d957611758903690600401611682565b90611761613a27565b60005b81518110156117dd5780611785611780610e776001948661273a565b614e0e565b611790575b01611764565b73ffffffffffffffffffffffffffffffffffffffff6117b2610e77838661273a565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a261178a565b8260005b815181101561080b57806118026117fd610e776001948661273a565b614e2f565b61180d575b016117e1565b73ffffffffffffffffffffffffffffffffffffffff61182f610e77838661273a565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611807565b346101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95761044561190161189761023b565b67ffffffffffffffff602435916118ad836101bb565b600060606040516118bd816105cf565b828152826020820152826040820152015216600052600860205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b60ff60405191611910836105cf565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015260405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101d957604051611a53816105ae565b816004013567ffffffffffffffff81116101d957611a779060043691850101611682565b8152602482013567ffffffffffffffff81116101d95761080b926004611aa09236920101611682565b6020820152612df6565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957611b05611aff610445923690600401610eca565b90613135565b6040519182916020835260208301906102af565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957366023820112156101d957806004013590611b7482610686565b90611b826040519283610624565b8282526024602083019360061b820101903682116101d957602401925b818410611baf5761080b836131e5565b6040843603126101d95760206040918251611bc9816105ae565b8635611bd4816101bb565b8152611be1838801610269565b83820152815201930192611b9f565b359060ff821682036101d957565b35907fffffffff00000000000000000000000000000000000000000000000000000000821682036101d957565b359061ffff821682036101d957565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760043567ffffffffffffffff81116101d957366023820112156101d957806004013590611c9582610686565b90611ca36040519283610624565b8282526024610160602084019402820101903682116101d957602401925b818410611cd15761080b836132c1565b83360361016081126101d9576101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192611d0e846105ae565b611d1788610269565b845201126101d95761016091602091611d2e610676565b611d398489016106b9565b8152611d476040890161069e565b84820152611d576060890161069e565b6040820152611d686080890161069e565b6060820152611d7960a08901611bf0565b6080820152611d8a60c08901611bfe565b60a0820152611d9b60e08901611c2b565b60c0820152611dad610100890161069e565b60e0820152611dbf610120890161069e565b610100820152611dd2610140890161069e565b61012082015283820152815201930192611cc1565b346101d95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95760405180602060095491828152019060096000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af9060005b818110611e665761044585610b1581870382610624565b8254845260209093019260019283019201611e4f565b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d9576040611ebc6004356111df816101bb565b611ef68251809263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565bf35b346101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957611f2f61023b565b60243567ffffffffffffffff81116101d957806004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101d957611f98611f938267ffffffffffffffff166000526007602052604060002090565b612db6565b611fa9611fa58251151590565b1590565b61221f576064830190611fed611fa5611fc184612623565b73ffffffffffffffffffffffffffffffffffffffff166000526001600901602052604060002054151590565b6121cd57611ffc858285614164565b9161200961126582612623565b956044860160008061201b8385612cdd565b15905061219a5750509061206461ffff61206e938861203f60c089015161ffff1690565b9161205b61205460e08b015163ffffffff1690565b9187612cdd565b94909316614896565b9094909390612623565b6120989073ffffffffffffffffffffffffffffffffffffffff166000526006602052604060002090565b5467ffffffffffffffff1667ffffffffffffffff166120b6916124e6565b966024016120c39161266d565b6120d49263ffffffff169150613653565b608083015160ff1660ff166120e8916124e6565b60609092015163ffffffff16906120fe91613678565b63ffffffff169061210e91613653565b9061211891613653565b906121379067ffffffffffffffff166000526004602052604060002090565b54612154916dffffffffffffffffffffffffffff909116906124e6565b61215d9061247c565b9061216791613653565b907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661218f916124f9565b604051908152602090f35b9291509261206e6121c76121c26121b961012089015163ffffffff1690565b63ffffffff1690565b612455565b91612623565b61221b6121d983612623565b7f2502348c0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6000fd5b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b346101d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d95773ffffffffffffffffffffffffffffffffffffffff6004356122a7816101bb565b6122af613a27565b1633811461232157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d957600435612386816101bb565b67ffffffffffffffff612397610252565b169081600052600760205260ff60406000205416156123f8576123b990613692565b6000918252600460209081526040928390205483517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152921690820152f35b507f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc10000149015171561247757565b612426565b90670de0b6b3a7640000820291808304670de0b6b3a7640000149015171561247757565b908160051b918083046020149015171561247757565b9061012c82029180830461012c149015171561247757565b9061010c82029180830461010c149015171561247757565b8181029291811591840414171561247757565b8115612503570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b61257161256b610f2a94937bffffffffffffffffffffffffffffffffffffffffffffffffffffffff6125648195613692565b16906124e6565b92613692565b16906124f9565b9061258282610686565b61258f6040519182610624565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125bd8294610686565b019060005b8281106125ce57505050565b8060606020809385010152016125c2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561261e5760061b0190565b6125df565b35610f2a816101bb565b919081101561261e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156101d9570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101d9570180359067ffffffffffffffff82116101d9576020019181360383136101d957565b9291926126ca8261098f565b916126d86040519384610624565b8294818452818301116101d9578281602093846000960137010152565b90604051612702816105cf565b606060ff82945463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c166040850152821c161515910152565b805182101561261e5760209160051b010190565b90929161279b6127728367ffffffffffffffff166000526007602052604060002090565b5460701b7fffffffff000000000000000000000000000000000000000000000000000000001690565b906127a581612578565b9560005b8281106127ba575050505050505090565b6127cd6127c882848961260e565b612623565b83886127e76127dd85848461262d565b604081019061266d565b90506020811161292d575b50839261282861282261281b6128116001986128709761286b9761262d565b602081019061266d565b36916126be565b89613741565b6128468967ffffffffffffffff166000526008602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b6126f5565b6060810151156128f3576128d761289160206128ab93015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610624565b6128e1828b61273a565b526128ec818a61273a565b50016127a9565b506128ab6128d761292861291b8967ffffffffffffffff166000526007602052604060002090565b5460a01c63ffffffff1690565b612891565b9150506129656121b9612958846128468b67ffffffffffffffff166000526008602052604060002090565b5460401c63ffffffff1690565b10612972578388386127f2565b7f36f536ca0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b604051906129c2826105ae565b60006020838281520152565b906129d7613a27565b6000915b8051831015612c0f576129ee838261273a565b5190612a02825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b85518051821015612bfa57612a2e8260209261273a565b510151612a5a612a3f83895161273a565b515173ffffffffffffffffffffffffffffffffffffffff1690565b604082015163ffffffff1660208110612bac575090867f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed173ffffffffffffffffffffffffffffffffffffffff84612b59858f60019998612846612ad19267ffffffffffffffff166000526008602052604060002090565b815181546020808501516040808701516060978801517fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951663ffffffff96909616959095179190921b67ffffffff00000000161792901b6bffffffff0000000000000000169190911790151590921b6cff00000000000000000000000016919091179055565b612ba3604051928392169582919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b0390a301612a17565b7f24ecdc020000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff90911660045263ffffffff1660245260446000fd5b505095509250926001915001919290926129db565b50905060005b8151811015612cd95780612c3d612c2e6001938561273a565b515167ffffffffffffffff1690565b67ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff612c866020612c6a868961273a565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b6000612caa826128468767ffffffffffffffff166000526008602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a301612c15565b5050565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101d9570180359067ffffffffffffffff82116101d957602001918160061b360383136101d957565b35907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821682036101d957565b6040813603126101d957612d89602060405192612d76846105ae565b8035612d81816101bb565b845201612d31565b602082015290565b6040813603126101d957612d89602060405192612dad846105ae565b612d8181610269565b90610674604051612dc6816105eb565b610120612deb82955460ff8116151584526114f16114e48263ffffffff9060081c1690565b63ffffffff16910152565b612dfe613a27565b60208101519160005b8351811015612eb25780612e20610e776001938761273a565b612e5c612e5773ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b615273565b612e68575b5001612e07565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138612e61565b5091505160005b8151811015612cd957612ecf610e77828461273a565b9073ffffffffffffffffffffffffffffffffffffffff821615612f6c577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef612f6383612f3b612f36612e3e60019773ffffffffffffffffffffffffffffffffffffffff1690565b6151fa565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101612eb9565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b90929192836004116101d95783116101d957600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b906004116101d95790600490565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613013575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9080601f830112156101d957813561305c81610686565b9261306a6040519485610624565b81845260208085019260051b8201019283116101d957602001905b8282106130925750505090565b8135815260209182019101613085565b6020818303126101d95780359067ffffffffffffffff82116101d9570160a0818303126101d957604051916130d683610608565b6130df8261069e565b83526130ed60208301610269565b60208401526040820135613100816106af565b604084015260608201356060840152608082013567ffffffffffffffff81116101d95761312d9201613045565b608082015290565b90600481108015613190575b61317957606061316361315b83610f2a946128ab96612f96565b8101906130a2565b0151604051928391602083019190602083019252565b5050604051613189602082610624565b6000815290565b50806004116101d95781357fffffffff00000000000000000000000000000000000000000000000000000000167f1f3b3aba000000000000000000000000000000000000000000000000000000001415613141565b6131ed613a27565b60005b8151811015612cd9578073ffffffffffffffffffffffffffffffffffffffff61321b6001938561273a565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d6132b867ffffffffffffffff6020613257868961273a565b51015116836000526006602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a2016131f0565b906132ca613a27565b60005b82518110156135bd576132e0818461273a565b5160206132f0612c2e848761273a565b9101519067ffffffffffffffff81168015801561359e575b8015613570575b8015613431575b6133f957916133bf826001959461336f61334a6127726133c49767ffffffffffffffff166000526007602052604060002090565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b6133ca577f6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b604051806133a2878261135f565b0390a267ffffffffffffffff166000526007602052604060002090565b613df3565b016132cd565b7f8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad604051806133a2878261135f565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b507fffffffff0000000000000000000000000000000000000000000000000000000061348060a08501517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c000000000000000000000000000000000000000000000000000000008114159081613545575b8161351a575b816134ef575b816134c4575b50613316565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501415386134be565b7fc4e059530000000000000000000000000000000000000000000000000000000081141591506134b8565b7fac77ffec0000000000000000000000000000000000000000000000000000000081141591506134b2565b7f1e10bdc40000000000000000000000000000000000000000000000000000000081141591506134ac565b5061010083015163ffffffff1663ffffffff6135966121b9604087015163ffffffff1690565b91161161330f565b5063ffffffff6135b661010085015163ffffffff1690565b1615613308565b509050565b73ffffffffffffffffffffffffffffffffffffffff906135e06129b5565b501660005260056020526040600020604051906135fc826105ae565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c602082015290565b906002820180921161247757565b906020820180921161247757565b906001820180921161247757565b9190820180921161247757565b63ffffffff60209116019063ffffffff821161247757565b9063ffffffff8091169116019063ffffffff821161247757565b61369b816135c2565b9063ffffffff60208301511615801561371a575b6136d65750517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff825116156136af565b7fffffffff000000000000000000000000000000000000000000000000000000001691907f2812d52c000000000000000000000000000000000000000000000000000000008314613888577f1e10bdc400000000000000000000000000000000000000000000000000000000831461387a577fac77ffec00000000000000000000000000000000000000000000000000000000831461386f577f647e2ba9000000000000000000000000000000000000000000000000000000008314613864577fc4e0595300000000000000000000000000000000000000000000000000000000831461385657827f2ee820750000000000000000000000000000000000000000000000000000000060005260045260246000fd5b610674919250600b90614a92565b610674919250614af3565b610674919250614a2f565b610674919250600190614a92565b61067491925061499f565b917fffffffff0000000000000000000000000000000000000000000000000000000083167f2812d52c000000000000000000000000000000000000000000000000000000008114613a1b577f1e10bdc40000000000000000000000000000000000000000000000000000000081146139fb577fac77ffec0000000000000000000000000000000000000000000000000000000081146139ef577f647e2ba90000000000000000000000000000000000000000000000000000000081146139e3577fc4e0595300000000000000000000000000000000000000000000000000000000146139c9577f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000831660045260246000fd5b6106749250156139db57600b90614a92565b600090614a92565b50506106749150614af3565b50506106749150614a2f565b50610674925015613a125760ff60015b1690614a92565b60ff6000613a0b565b5050610674915061499f565b73ffffffffffffffffffffffffffffffffffffffff600154163303613a4857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b33600052600360205260406000205415613a8857565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b611f93613add9196949395929667ffffffffffffffff166000526007602052604060002090565b9460a08601947fffffffff00000000000000000000000000000000000000000000000000000000613b2e87517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c000000000000000000000000000000000000000000000000000000008114908115613dc9575b8115613d9f575b50613d5a577fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613bd288517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613d3c5750507f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613c4686517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613cc75761221b613c7985517fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b613d09935061281b6060613cf3613cec6121b96040613d35989a015163ffffffff1690565b8486614cda565b0151604051958691602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610624565b9160019190565b94509450610f2a91613d4f9136916126be565b9360019336916126be565b94509491613d8091613d7a6121b9610100610f2a96015163ffffffff1690565b91614b6f565b93613d976020613d8f87614c90565b960151151590565b9336916126be565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501438613b61565b7fac77ffec0000000000000000000000000000000000000000000000000000000081149150613b5a565b61410961012061067493613e3b613e0a8251151590565b859060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b613e85613e4f602083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ff1660089190911b64ffffffff0016178555565b613ed3613e99604083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1660289190911b68ffffffff000000000016178555565b613f25613ee7606083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffff1660489190911b6cffffffff00000000000000000016178555565b613f75613f36608083015160ff1690565b85547fffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffff1660689190911b6dff0000000000000000000000000016178555565b613fe8613fa560a08301517fffffffff000000000000000000000000000000000000000000000000000000001690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911c71ffffffff000000000000000000000000000016178555565b61403f613ffa60c083015161ffff1690565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b61409c61405360e083015163ffffffff1690565b85547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b77ffffffff000000000000000000000000000000000000000016178555565b6140fe6140b161010083015163ffffffff1690565b85547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff1660c09190911b7bffffffff00000000000000000000000000000000000000000000000016178555565b015163ffffffff1690565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000083549260e01b169116179055565b908160209103126101d9573590565b9190614173602083018361266d565b939050604083016141848185612cdd565b9050602084019161419c6121b9845163ffffffff1690565b8088116148645750600182116148305760a08501966141db88517fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f2812d52c0000000000000000000000000000000000000000000000000000000081148015614807575b80156147de575b156142bf575050505050505091816142b961281b6142886142b29661425a6080610f2a98018661266d565b6142826121b96040614276610100879697015163ffffffff1690565b94015163ffffffff1690565b92614f9b565b51958694517fffffffff000000000000000000000000000000000000000000000000000000001690565b928061266d565b90613893565b7fc4e0595300000000000000000000000000000000000000000000000000000000819b9a939495979996989b14600014614562575050614362614329614355999a60406143236121b961431560808b018b61266d565b939094015163ffffffff1690565b91614e50565b918251998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b6142b961281b888061266d565b606081015151908261437f614377878061266d565b810190614155565b61454557508161450f575b8515159081614502575b506144d857604081116144a657506143b9906143b385949397956124ce565b90613653565b946000935b8385106144175750505050506121b96143db915163ffffffff1690565b8082116143e757505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b909192939560019061447b6121b96129586144468667ffffffffffffffff166000526008602052604060002090565b6144576127c88d610c028b8d612cdd565b73ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b80156144965761448a91613653565b965b01939291906143be565b506144a090613637565b9661448c565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614394565b61221b827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b61455c9193506143b361455784613645565b6124a0565b9161438a565b7f1e10bdc4000000000000000000000000000000000000000000000000000000000361478f57506145ea6145af614355999a60406145a96121b961431560808b018b61266d565b91614cda565b916145c16121b9845163ffffffff1690565b998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b60808101515190826145ff614377878061266d565b614777575081614741575b85151580614735575b6144d85760408211614701576020015167ffffffffffffffff9081169081831c166146c757505061464b906143b385949397956124b6565b946000935b83851061466d5750505050506121b96143db915163ffffffff1690565b909192939560019061469c6121b96129586144468667ffffffffffffffff166000526008602052604060002090565b80156146b7576146ab91613653565b965b0193929190614650565b506146c190613637565b966146ad565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614613565b61221b827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b6147899193506143b361455784613629565b9161460a565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b507f647e2ba900000000000000000000000000000000000000000000000000000000811461422f565b507fac77ffec000000000000000000000000000000000000000000000000000000008114614228565b7fd88dddd6000000000000000000000000000000000000000000000000000000006000526004829052600160245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b9194929390946000926000966000966000945b8086106148bb57505050505050929190565b9091929394956148f661286b6148e8889c9b9c67ffffffffffffffff166000526008602052604060002090565b6144576127c88b878961260e565b98614907611fa560608c0151151590565b6149655760208a015163ffffffff1661491f91613678565b9960408a01516149329063ffffffff1690565b61493b91613678565b985161494c9063ffffffff16612455565b61495591613653565b956001905b0194939291906148a9565b96985061498a6149848561497e6001946143b38a612455565b99613678565b9a613660565b9861495a565b908160209103126101d9575190565b60208151036149e2576149bb6020825183010160208301614990565b73ffffffffffffffffffffffffffffffffffffffff8111908115614a23575b506149e25750565b614a1f906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906102af565b0390fd5b610400915010386149da565b6020815103614a5557600b614a4d6020835184010160208401614990565b10614a555750565b614a1f906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840181815201906102af565b906020825103614ab85780614aa5575050565b614a4d6020835184010160208401614990565b6040517fe0d7fb020000000000000000000000000000000000000000000000000000000081526020600482015280614a1f60248201856102af565b6024815103614b0957602281015115614b095750565b614a1f906040519182917f373b0e4400000000000000000000000000000000000000000000000000000000835260206004840181815201906102af565b908160409103126101d957602060405191614b60836105ae565b805183520151612d89816106af565b91614b786129b5565b508115614c6e5750614bb961281b8280614bb37fffffffff000000000000000000000000000000000000000000000000000000009587612fdf565b95612f96565b91167f181dcf10000000000000000000000000000000000000000000000000000000008103614bf6575080602080610f2a93518301019101614b46565b7f97a657c90000000000000000000000000000000000000000000000000000000014614c46577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b80602080614c5993518301019101614990565b614c61610665565b9081526000602082015290565b91505067ffffffffffffffff614c82610665565b911681526000602082015290565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610f2a606482610624565b60606080604051614cea81610608565b60008152600060208201526000604082015260008382015201528115614de4577f1f3b3aba000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000614d5f614d598585612fd1565b90612fdf565b1603614dba5781614d739261315b92612f96565b9063ffffffff614d87835163ffffffff1690565b1611614d905790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff610f2a911660096150a7565b73ffffffffffffffffffffffffffffffffffffffff610f2a91166009615235565b90606080604051614e60816105cf565b60008152600060208201526000604082015201528015614de4577f21ea4ca9000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000614ec9614d598486612fd1565b1603614dba5780614ed992612f96565b81929101600092602081830312614f975780359067ffffffffffffffff8211614f93570192608084830312614f905760405193614f15856105cf565b803585526020810135614f27816106af565b60208601526040810135604086015260608101359167ffffffffffffffff8311614f905750614f57929101613045565b6060830152815111614f665790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b80fd5b8480fd5b8380fd5b9063ffffffff614fb593614fad6129b5565b501691614b6f565b90815111614f665790565b805482101561261e5760005260206000200190600090565b91615010918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b80548015615078577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906150498282614fc0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615195577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612477578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff850194851161247757600095858361514697615137950361514c575b505050615014565b90600052602052604060002090565b55600190565b61517c6151769161516d61516361518c9588614fc0565b90549060031b1c90565b92839187614fc0565b90614fd8565b8590600052602052604060002090565b5538808061512f565b50505050600090565b805490680100000000000000008210156105ca57816151c591600161501094018155614fc0565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b60008181526003602052604090205461522f5761521881600261519e565b600254906000526003602052604060002055600190565b50600090565b600082815260018201602052604090205461526c57806152578360019361519e565b80549260005201602052604060002055600190565b5050600090565b60008181526003602052604090205490811561526c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161247757600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612477578383615146946000960361530f575b5050506152fe6002615014565b600390600052602052604060002090565b6152fe61517691615327615163615331956002614fc0565b9283916002614fc0565b553880806152f156fea164736f6c634300081a000a",
}

var FeeQuoterV2ABI = FeeQuoterV2MetaData.ABI

var FeeQuoterV2Bin = FeeQuoterV2MetaData.Bin

func DeployFeeQuoterV2(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoterV2, error) {
	parsed, err := FeeQuoterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterV2Bin), backend, staticConfig, priceUpdaters, feeTokens, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeQuoterV2{address: address, abi: *parsed, FeeQuoterV2Caller: FeeQuoterV2Caller{contract: contract}, FeeQuoterV2Transactor: FeeQuoterV2Transactor{contract: contract}, FeeQuoterV2Filterer: FeeQuoterV2Filterer{contract: contract}}, nil
}

type FeeQuoterV2 struct {
	address common.Address
	abi     abi.ABI
	FeeQuoterV2Caller
	FeeQuoterV2Transactor
	FeeQuoterV2Filterer
}

type FeeQuoterV2Caller struct {
	contract *bind.BoundContract
}

type FeeQuoterV2Transactor struct {
	contract *bind.BoundContract
}

type FeeQuoterV2Filterer struct {
	contract *bind.BoundContract
}

type FeeQuoterV2Session struct {
	Contract     *FeeQuoterV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeeQuoterV2CallerSession struct {
	Contract *FeeQuoterV2Caller
	CallOpts bind.CallOpts
}

type FeeQuoterV2TransactorSession struct {
	Contract     *FeeQuoterV2Transactor
	TransactOpts bind.TransactOpts
}

type FeeQuoterV2Raw struct {
	Contract *FeeQuoterV2
}

type FeeQuoterV2CallerRaw struct {
	Contract *FeeQuoterV2Caller
}

type FeeQuoterV2TransactorRaw struct {
	Contract *FeeQuoterV2Transactor
}

func NewFeeQuoterV2(address common.Address, backend bind.ContractBackend) (*FeeQuoterV2, error) {
	abi, err := abi.JSON(strings.NewReader(FeeQuoterV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeeQuoterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2{address: address, abi: abi, FeeQuoterV2Caller: FeeQuoterV2Caller{contract: contract}, FeeQuoterV2Transactor: FeeQuoterV2Transactor{contract: contract}, FeeQuoterV2Filterer: FeeQuoterV2Filterer{contract: contract}}, nil
}

func NewFeeQuoterV2Caller(address common.Address, caller bind.ContractCaller) (*FeeQuoterV2Caller, error) {
	contract, err := bindFeeQuoterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2Caller{contract: contract}, nil
}

func NewFeeQuoterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FeeQuoterV2Transactor, error) {
	contract, err := bindFeeQuoterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2Transactor{contract: contract}, nil
}

func NewFeeQuoterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FeeQuoterV2Filterer, error) {
	contract, err := bindFeeQuoterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2Filterer{contract: contract}, nil
}

func bindFeeQuoterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeQuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FeeQuoterV2 *FeeQuoterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoterV2.Contract.FeeQuoterV2Caller.contract.Call(opts, result, method, params...)
}

func (_FeeQuoterV2 *FeeQuoterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.FeeQuoterV2Transactor.contract.Transfer(opts)
}

func (_FeeQuoterV2 *FeeQuoterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.FeeQuoterV2Transactor.contract.Transact(opts, method, params...)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoterV2.Contract.contract.Call(opts, result, method, params...)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.contract.Transfer(opts)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.contract.Transact(opts, method, params...)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "convertTokenAmount", fromToken, fromTokenAmount, toToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoterV2.Contract.ConvertTokenAmount(&_FeeQuoterV2.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoterV2.Contract.ConvertTokenAmount(&_FeeQuoterV2.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoterV2.Contract.GetAllAuthorizedCallers(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoterV2.Contract.GetAllAuthorizedCallers(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(FeeQuoterDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterDestChainConfig)).(*FeeQuoterDestChainConfig)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoterV2.Contract.GetDestChainConfig(&_FeeQuoterV2.CallOpts, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoterV2.Contract.GetDestChainConfig(&_FeeQuoterV2.CallOpts, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getDestinationChainGasPrice", destChainSelector)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetDestinationChainGasPrice(&_FeeQuoterV2.CallOpts, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetDestinationChainGasPrice(&_FeeQuoterV2.CallOpts, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getFeeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoterV2.Contract.GetFeeTokens(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoterV2.Contract.GetFeeTokens(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getPremiumMultiplierWeiPerEth", token)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoterV2.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoterV2.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FeeQuoterStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterStaticConfig)).(*FeeQuoterStaticConfig)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoterV2.Contract.GetStaticConfig(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoterV2.Contract.GetStaticConfig(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getTokenAndGasPrices", token, destChainSelector)

	outstruct := new(GetTokenAndGasPrices)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPriceValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoterV2.Contract.GetTokenAndGasPrices(&_FeeQuoterV2.CallOpts, token, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoterV2.Contract.GetTokenAndGasPrices(&_FeeQuoterV2.CallOpts, token, destChainSelector)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getTokenPrice", token)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetTokenPrice(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetTokenPrice(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getTokenPrices", tokens)

	if err != nil {
		return *new([]InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new([]InternalTimestampedPackedUint224)).(*[]InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetTokenPrices(&_FeeQuoterV2.CallOpts, tokens)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoterV2.Contract.GetTokenPrices(&_FeeQuoterV2.CallOpts, tokens)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getTokenTransferFeeConfig", destChainSelector, token)

	if err != nil {
		return *new(FeeQuoterTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenTransferFeeConfig)).(*FeeQuoterTokenTransferFeeConfig)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoterV2.Contract.GetTokenTransferFeeConfig(&_FeeQuoterV2.CallOpts, destChainSelector, token)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoterV2.Contract.GetTokenTransferFeeConfig(&_FeeQuoterV2.CallOpts, destChainSelector, token)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getValidatedFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoterV2.Contract.GetValidatedFee(&_FeeQuoterV2.CallOpts, destChainSelector, message)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoterV2.Contract.GetValidatedFee(&_FeeQuoterV2.CallOpts, destChainSelector, message)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getValidatedTokenPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoterV2.Contract.GetValidatedTokenPrice(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoterV2.Contract.GetValidatedTokenPrice(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) Owner() (common.Address, error) {
	return _FeeQuoterV2.Contract.Owner(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) Owner() (common.Address, error) {
	return _FeeQuoterV2.Contract.Owner(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "processMessageArgs", destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)

	outstruct := new(ProcessMessageArgs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MsgFeeJuels = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConvertedExtraArgs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.TokenReceiver = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	return _FeeQuoterV2.Contract.ProcessMessageArgs(&_FeeQuoterV2.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	return _FeeQuoterV2.Contract.ProcessMessageArgs(&_FeeQuoterV2.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) ProcessPoolReturnData(opts *bind.CallOpts, destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "processPoolReturnData", destChainSelector, onRampTokenTransfers, sourceTokenAmounts)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) ProcessPoolReturnData(destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	return _FeeQuoterV2.Contract.ProcessPoolReturnData(&_FeeQuoterV2.CallOpts, destChainSelector, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) ProcessPoolReturnData(destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	return _FeeQuoterV2.Contract.ProcessPoolReturnData(&_FeeQuoterV2.CallOpts, destChainSelector, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) ResolveTokenReceiver(opts *bind.CallOpts, extraArgs []byte) ([]byte, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "resolveTokenReceiver", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) ResolveTokenReceiver(extraArgs []byte) ([]byte, error) {
	return _FeeQuoterV2.Contract.ResolveTokenReceiver(&_FeeQuoterV2.CallOpts, extraArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) ResolveTokenReceiver(extraArgs []byte) ([]byte, error) {
	return _FeeQuoterV2.Contract.ResolveTokenReceiver(&_FeeQuoterV2.CallOpts, extraArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) TypeAndVersion() (string, error) {
	return _FeeQuoterV2.Contract.TypeAndVersion(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) TypeAndVersion() (string, error) {
	return _FeeQuoterV2.Contract.TypeAndVersion(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "acceptOwnership")
}

func (_FeeQuoterV2 *FeeQuoterV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.AcceptOwnership(&_FeeQuoterV2.TransactOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.AcceptOwnership(&_FeeQuoterV2.TransactOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoterV2.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoterV2.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyDestChainConfigUpdates(&_FeeQuoterV2.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyDestChainConfigUpdates(&_FeeQuoterV2.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "applyFeeTokensUpdates", feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyFeeTokensUpdates(&_FeeQuoterV2.TransactOpts, feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyFeeTokensUpdates(&_FeeQuoterV2.TransactOpts, feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "applyPremiumMultiplierWeiPerEthUpdates", premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoterV2.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoterV2.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoterV2.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoterV2.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "transferOwnership", to)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.TransferOwnership(&_FeeQuoterV2.TransactOpts, to)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.TransferOwnership(&_FeeQuoterV2.TransactOpts, to)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "updatePrices", priceUpdates)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.UpdatePrices(&_FeeQuoterV2.TransactOpts, priceUpdates)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.UpdatePrices(&_FeeQuoterV2.TransactOpts, priceUpdates)
}

type FeeQuoterV2AuthorizedCallerAddedIterator struct {
	Event *FeeQuoterV2AuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2AuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2AuthorizedCallerAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2AuthorizedCallerAdded)
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

func (it *FeeQuoterV2AuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2AuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2AuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterV2AuthorizedCallerAddedIterator, error) {

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2AuthorizedCallerAddedIterator{contract: _FeeQuoterV2.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2AuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2AuthorizedCallerAdded)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterV2AuthorizedCallerAdded, error) {
	event := new(FeeQuoterV2AuthorizedCallerAdded)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2AuthorizedCallerRemovedIterator struct {
	Event *FeeQuoterV2AuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2AuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2AuthorizedCallerRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2AuthorizedCallerRemoved)
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

func (it *FeeQuoterV2AuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2AuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2AuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterV2AuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2AuthorizedCallerRemovedIterator{contract: _FeeQuoterV2.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2AuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2AuthorizedCallerRemoved)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterV2AuthorizedCallerRemoved, error) {
	event := new(FeeQuoterV2AuthorizedCallerRemoved)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2DestChainAddedIterator struct {
	Event *FeeQuoterV2DestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2DestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2DestChainAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2DestChainAdded)
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

func (it *FeeQuoterV2DestChainAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2DestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2DestChainAdded struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterV2DestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2DestChainAddedIterator{contract: _FeeQuoterV2.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2DestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2DestChainAdded)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseDestChainAdded(log types.Log) (*FeeQuoterV2DestChainAdded, error) {
	event := new(FeeQuoterV2DestChainAdded)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2DestChainConfigUpdatedIterator struct {
	Event *FeeQuoterV2DestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2DestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2DestChainConfigUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2DestChainConfigUpdated)
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

func (it *FeeQuoterV2DestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2DestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2DestChainConfigUpdated struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterV2DestChainConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2DestChainConfigUpdatedIterator{contract: _FeeQuoterV2.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2DestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2DestChainConfigUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterV2DestChainConfigUpdated, error) {
	event := new(FeeQuoterV2DestChainConfigUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2FeeTokenAddedIterator struct {
	Event *FeeQuoterV2FeeTokenAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2FeeTokenAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2FeeTokenAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2FeeTokenAdded)
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

func (it *FeeQuoterV2FeeTokenAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2FeeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2FeeTokenAdded struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterV2FeeTokenAddedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2FeeTokenAddedIterator{contract: _FeeQuoterV2.contract, event: "FeeTokenAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2FeeTokenAdded, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2FeeTokenAdded)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseFeeTokenAdded(log types.Log) (*FeeQuoterV2FeeTokenAdded, error) {
	event := new(FeeQuoterV2FeeTokenAdded)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2FeeTokenRemovedIterator struct {
	Event *FeeQuoterV2FeeTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2FeeTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2FeeTokenRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2FeeTokenRemoved)
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

func (it *FeeQuoterV2FeeTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2FeeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2FeeTokenRemoved struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterV2FeeTokenRemovedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2FeeTokenRemovedIterator{contract: _FeeQuoterV2.contract, event: "FeeTokenRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2FeeTokenRemoved, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2FeeTokenRemoved)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseFeeTokenRemoved(log types.Log) (*FeeQuoterV2FeeTokenRemoved, error) {
	event := new(FeeQuoterV2FeeTokenRemoved)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2OwnershipTransferRequestedIterator struct {
	Event *FeeQuoterV2OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2OwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2OwnershipTransferRequested)
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

func (it *FeeQuoterV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2OwnershipTransferRequestedIterator{contract: _FeeQuoterV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2OwnershipTransferRequested)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterV2OwnershipTransferRequested, error) {
	event := new(FeeQuoterV2OwnershipTransferRequested)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2OwnershipTransferredIterator struct {
	Event *FeeQuoterV2OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2OwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2OwnershipTransferred)
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

func (it *FeeQuoterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2OwnershipTransferredIterator{contract: _FeeQuoterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2OwnershipTransferred)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseOwnershipTransferred(log types.Log) (*FeeQuoterV2OwnershipTransferred, error) {
	event := new(FeeQuoterV2OwnershipTransferred)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator struct {
	Event *FeeQuoterV2PremiumMultiplierWeiPerEthUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2PremiumMultiplierWeiPerEthUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2PremiumMultiplierWeiPerEthUpdated)
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

func (it *FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2PremiumMultiplierWeiPerEthUpdated struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
	Raw                        types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator{contract: _FeeQuoterV2.contract, event: "PremiumMultiplierWeiPerEthUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2PremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2PremiumMultiplierWeiPerEthUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterV2PremiumMultiplierWeiPerEthUpdated, error) {
	event := new(FeeQuoterV2PremiumMultiplierWeiPerEthUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2TokenTransferFeeConfigDeletedIterator struct {
	Event *FeeQuoterV2TokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2TokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2TokenTransferFeeConfigDeleted)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2TokenTransferFeeConfigDeleted)
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

func (it *FeeQuoterV2TokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2TokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2TokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Token             common.Address
	Raw               types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterV2TokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2TokenTransferFeeConfigDeletedIterator{contract: _FeeQuoterV2.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2TokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2TokenTransferFeeConfigDeleted)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterV2TokenTransferFeeConfigDeleted, error) {
	event := new(FeeQuoterV2TokenTransferFeeConfigDeleted)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2TokenTransferFeeConfigUpdatedIterator struct {
	Event *FeeQuoterV2TokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2TokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2TokenTransferFeeConfigUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2TokenTransferFeeConfigUpdated)
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

func (it *FeeQuoterV2TokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2TokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2TokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	Token                  common.Address
	TokenTransferFeeConfig FeeQuoterTokenTransferFeeConfig
	Raw                    types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterV2TokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2TokenTransferFeeConfigUpdatedIterator{contract: _FeeQuoterV2.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2TokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2TokenTransferFeeConfigUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterV2TokenTransferFeeConfigUpdated, error) {
	event := new(FeeQuoterV2TokenTransferFeeConfigUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2UsdPerTokenUpdatedIterator struct {
	Event *FeeQuoterV2UsdPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2UsdPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2UsdPerTokenUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2UsdPerTokenUpdated)
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

func (it *FeeQuoterV2UsdPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2UsdPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2UsdPerTokenUpdated struct {
	Token     common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2UsdPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2UsdPerTokenUpdatedIterator{contract: _FeeQuoterV2.contract, event: "UsdPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2UsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2UsdPerTokenUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterV2UsdPerTokenUpdated, error) {
	event := new(FeeQuoterV2UsdPerTokenUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2UsdPerUnitGasUpdatedIterator struct {
	Event *FeeQuoterV2UsdPerUnitGasUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2UsdPerUnitGasUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2UsdPerUnitGasUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(FeeQuoterV2UsdPerUnitGasUpdated)
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

func (it *FeeQuoterV2UsdPerUnitGasUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2UsdPerUnitGasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2UsdPerUnitGasUpdated struct {
	DestChain uint64
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterV2UsdPerUnitGasUpdatedIterator, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2UsdPerUnitGasUpdatedIterator{contract: _FeeQuoterV2.contract, event: "UsdPerUnitGasUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2UsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2UsdPerUnitGasUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterV2UsdPerUnitGasUpdated, error) {
	event := new(FeeQuoterV2UsdPerUnitGasUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetTokenAndGasPrices struct {
	TokenPrice    *big.Int
	GasPriceValue *big.Int
}
type ProcessMessageArgs struct {
	MsgFeeJuels           *big.Int
	IsOutOfOrderExecution bool
	ConvertedExtraArgs    []byte
	TokenReceiver         []byte
}

func (FeeQuoterV2AuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (FeeQuoterV2AuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (FeeQuoterV2DestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b")
}

func (FeeQuoterV2DestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad")
}

func (FeeQuoterV2FeeTokenAdded) Topic() common.Hash {
	return common.HexToHash("0xdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23")
}

func (FeeQuoterV2FeeTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91")
}

func (FeeQuoterV2OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeeQuoterV2OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FeeQuoterV2PremiumMultiplierWeiPerEthUpdated) Topic() common.Hash {
	return common.HexToHash("0xbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d")
}

func (FeeQuoterV2TokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (FeeQuoterV2TokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed1")
}

func (FeeQuoterV2UsdPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a")
}

func (FeeQuoterV2UsdPerUnitGasUpdated) Topic() common.Hash {
	return common.HexToHash("0xdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e")
}

func (_FeeQuoterV2 *FeeQuoterV2) Address() common.Address {
	return _FeeQuoterV2.address
}

type FeeQuoterV2Interface interface {
	ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error)

	GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error)

	GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error)

	GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

		error)

	GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error)

	GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error)

	GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

		error)

	ProcessPoolReturnData(opts *bind.CallOpts, destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error)

	ResolveTokenReceiver(opts *bind.CallOpts, extraArgs []byte) ([]byte, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterV2AuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2AuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterV2AuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterV2AuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2AuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterV2AuthorizedCallerRemoved, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterV2DestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2DestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*FeeQuoterV2DestChainAdded, error)

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterV2DestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2DestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterV2DestChainConfigUpdated, error)

	FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterV2FeeTokenAddedIterator, error)

	WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2FeeTokenAdded, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenAdded(log types.Log) (*FeeQuoterV2FeeTokenAdded, error)

	FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterV2FeeTokenRemovedIterator, error)

	WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2FeeTokenRemoved, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenRemoved(log types.Log) (*FeeQuoterV2FeeTokenRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeQuoterV2OwnershipTransferred, error)

	FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2PremiumMultiplierWeiPerEthUpdatedIterator, error)

	WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2PremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error)

	ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterV2PremiumMultiplierWeiPerEthUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterV2TokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2TokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterV2TokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterV2TokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2TokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterV2TokenTransferFeeConfigUpdated, error)

	FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2UsdPerTokenUpdatedIterator, error)

	WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2UsdPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterV2UsdPerTokenUpdated, error)

	FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterV2UsdPerUnitGasUpdatedIterator, error)

	WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2UsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error)

	ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterV2UsdPerUnitGasUpdated, error)

	Address() common.Address
}
