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
	IsEnabled                         bool
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DestGasOverhead                   uint32
	DestGasPerPayloadByteBase         uint8
	DestGasPerPayloadByteHigh         uint8
	DestGasPerPayloadByteThreshold    uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	ChainFamilySelector               [4]byte
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTxGasLimit                 uint32
	GasMultiplierWeiPerEth            uint64
	GasPriceStalenessThreshold        uint32
	NetworkFeeUSDCents                uint32
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
	MaxFeeJuelsPerMsg            *big.Int
	LinkToken                    common.Address
	TokenPriceStalenessThreshold uint32
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_BASE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveTokenReceiver\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e060405234610e855761700b803803806100198161103b565b92833981018181036101008112610e8557606013610e855760405191606083016001600160401b03811184821017610e8a5760405280516001600160601b0381168103610e8557835261006e60208201611060565b916020840192835261008260408301611074565b6040850190815260608301519094906001600160401b038111610e8557826100ab91850161109c565b60808401519095906001600160401b038111610e8557836100cd91860161109c565b60a08501519093906001600160401b038111610e855785019481601f87011215610e855785519561010561010088611085565b61103b565b9660208089838152019160051b83010191848311610e855760208101915b838310610edc575050505060c08101516001600160401b038111610e855781019082601f83011215610e8557815161015d61010082611085565b9260208085848152019260061b82010190858211610e8557602001915b818310610ea05750505060e0810151906001600160401b038211610e85570182601f82011215610e85578051906101b361010083611085565b936020610240818786815201940283010191818311610e8557602001925b828410610ce957505050503315610cd857600180546001600160a01b031916331790556020956102008761103b565b9560008752600036813761021261101c565b998a5286888b015260005b8751811015610284576001906001600160a01b0361023b828b611135565b51168a61024782611321565b610254575b50500161021d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a61024c565b50879650888a519560005b8751811015610300576001600160a01b036102aa828a611135565b51169081156102ef577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8b836102e16001956112a9565b50604051908152a10161028f565b6342bcdf7f60e11b60005260046000fd5b50815189975088906001600160a01b0316158015610cc6575b8015610cb7575b610ca65791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c0526103538561103b565b9260008452600036813760005b84518110156103cf576001906103886001600160a01b036103818389611135565b51166111b6565b610393575b01610360565b818060a01b036103a38288611135565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a261038d565b508593508460005b8351811015610446576001906103ff6001600160a01b036103f88388611135565b51166112e8565b61040a575b016103d7565b818060a01b0361041a8287611135565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610404565b5091509160005b82518110156109b6576104608184611135565b51856001600160401b036104748487611135565b51511691015190801580156109a3575b8015610985575b8015610903575b6108ef57600081815260078852604090205460019392919060181b6001600160e01b0319166107c557807fe21664081fc677b92cffc05bf1c3609b90b261aafb6b647b8dae2a01d1c76eb7604051806105dd868291909161020063ffffffff8161022084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff608082015116608086015260ff60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e086015261ffff6101008201511661010086015261ffff610120820151166101208601528260e01b6101408201511661014086015261ffff610160820151166101608601528261018082015116610180860152826101a0820151166101a086015260018060401b036101c0820151166101c0860152826101e0820151166101e0860152015116910152565b0390a25b60009081526007885260409081902082518154848b0151938501516060860151608087015160a088015160c089015160e08a01516101008b01516101208c01516101408d01516101608e01517fff00000000000000000000000000000000000000000000000000000000000000909a1660ff9b15159b909b169a909a1760089c909c1b64ffffffff00169b909b1760289790971b68ffffffff0000000000169690961760489590951b6cffffffff000000000000000000169490941760689390931b6dff00000000000000000000000000169290921760709190911b60ff60701b161760789190911b61ffff60781b161760889190911b63ffffffff60881b161760a89190911b61ffff60a81b161760b89490941b61ffff60b81b169390931760189190911c63ffffffff60c81b161760e89290921b61ffff60e81b1691909117815561018082015190840180546101a08401516001600160401b031990911663ffffffff939093169290921760209290921b67ffffffff0000000016919091178155906101c081015190825463ffffffff60801b6101e083015160801b169161020063ffffffff60a01b91015160a01b1692680100000000000000008760801b039060401b1690680100000000000000008760c01b03191617171790550161044d565b807f0b24b226620af755f455598afdfb191eaf1236f264c327646ac8644608d7d31a604051806108e7868291909161020063ffffffff8161022084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff608082015116608086015260ff60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e086015261ffff6101008201511661010086015261ffff610120820151166101208601528260e01b6101408201511661014086015261ffff610160820151166101608601528261018082015116610180860152826101a0820151166101a086015260018060401b036101c0820151166101c0860152826101e0820151166101e0860152015116910152565b0390a26105e1565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61014083015116630a04b54b60e21b8114159081610973575b81610961575b8161094f575b8161093d575b50610492565b63647e2ba960e01b1415905088610937565b63c4e0595360e01b8114159150610931565b632b1dfffb60e21b811415915061092b565b6307842f7160e21b8114159150610925565b5063ffffffff6101a08301511663ffffffff6040840151161061048b565b5063ffffffff6101a08301511615610484565b84828560005b8151811015610a3c576001906001600160a01b036109da8285611135565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610a108689611135565b510151168360005260068252604060002081878060401b0319825416179055604051908152a2016109bc565b50506001610a498361103b565b9160008352600091610ca1575b81925b8151841015610be657610a6c8483611135565b5180516001600160401b0316929086019190845b8784518051831015610bd55782610a9691611135565b51015184516001600160a01b0390610aaf908490611135565b51511690604081019063ffffffff8251168b8110610bbe575091877f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed160808d97969460019663ffffffff60408f87815260088d528181208b8060a01b038a1682528d52209281808451169481861682198254161781558d850190610b5383835116829067ffffffff0000000082549160201b169067ffffffff000000001916179055565b8351606082549701966bffffffff00000000000000006cff0000000000000000000000008951151560601b169260401b16906cffffffffff000000000000000019161717905560405195865251168c850152511660408301525115156060820152a301909150610a80565b6312766e0160e11b8a526004849052602452604489fd5b505050925093600191500192610a59565b905083825b8251811015610c69576001906001600160401b03610c098286611135565b515116828060a01b0384610c1d8488611135565b51015116908087526008855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610beb565b604051615c5590816113b6823960805181818161063f0152611448015260a05181818161067801526113df015260c051816106a40152f35b610a56565b63d794ef9560e01b60005260046000fd5b5063ffffffff82511615610320565b5080516001600160601b031615610319565b639b15e16f60e01b60005260046000fd5b8382036102408112610e8557610220610d0061101c565b91610d0a876110f7565b8352601f190112610e85576040519161022083016001600160401b03811184821017610e8a57604052610d3f6020870161110b565b8352610d4d60408701611074565b6020840152610d5e60608701611074565b6040840152610d6f60808701611074565b6060840152610d8060a08701611118565b6080840152610d9160c08701611118565b60a0840152610da260e08701611126565b60c0840152610db46101008701611074565b60e0840152610dc66101208701611126565b610100840152610dd96101408701611126565b610120840152610160860151916001600160e01b031983168303610e855783602093610140610240960152610e116101808901611126565b610160820152610e246101a08901611074565b610180820152610e376101c08901611074565b6101a0820152610e4a6101e089016110f7565b6101c0820152610e5d6102008901611074565b6101e0820152610e706102208901611074565b610200820152838201528152019301926101d1565b600080fd5b634e487b7160e01b600052604160045260246000fd5b604083870312610e85576020604091610eb761101c565b610ec086611060565b8152610ecd8387016110f7565b8382015281520192019161017a565b82516001600160401b038111610e855782016040818803601f190112610e8557610f0461101c565b90610f11602082016110f7565b825260408101516001600160401b038111610e8557602091010187601f82011215610e85578051610f4461010082611085565b91602060a08185858152019302820101908a8211610e8557602001915b818310610f805750505091816020938480940152815201920191610123565b828b0360a08112610e85576080610f9561101c565b91610f9f86611060565b8352601f190112610e85576040519160808301916001600160401b03831184841017610e8a5760a093602093604052610fd9848801611074565b8152610fe760408801611074565b84820152610ff760608801611074565b60408201526110086080880161110b565b606082015283820152815201920191610f61565b60408051919082016001600160401b03811183821017610e8a57604052565b6040519190601f01601f191682016001600160401b03811183821017610e8a57604052565b51906001600160a01b0382168203610e8557565b519063ffffffff82168203610e8557565b6001600160401b038111610e8a5760051b60200190565b9080601f83011215610e855781516110b661010082611085565b9260208085848152019260051b820101928311610e8557602001905b8282106110df5750505090565b602080916110ec84611060565b8152019101906110d2565b51906001600160401b0382168203610e8557565b51908115158203610e8557565b519060ff82168203610e8557565b519061ffff82168203610e8557565b80518210156111495760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156111495760005260206000200190600090565b805480156111a057600019019061118e828261115f565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600a602052604090205480156112775760001981018181116112615760095460001981019190821161126157818103611210575b5050506111fc6009611177565b600052600a60205260006040812055600190565b61124961122161123293600961115f565b90549060031b1c928392600961115f565b819391549060031b91821b91600019901b19161790565b9055600052600a6020526040600020553880806111ef565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80549068010000000000000000821015610e8a57816112329160016112a59401815561115f565b9055565b806000526003602052604060002054156000146112e2576112cb81600261127e565b600254906000526003602052604060002055600190565b50600090565b80600052600a602052604060002054156000146112e25761130a81600961127e565b60095490600052600a602052604060002055600190565b6000818152600360205260409020548015611277576000198101818111611261576002546000198101919082116112615780820361137b575b5050506113676002611177565b600052600360205260006040812055600190565b61139d61138c61123293600261115f565b90549060031b1c928392600261115f565b9055600052600360205260406000205538808061135a56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be146101e657806301447eaa146101e157806301ffc9a7146101dc578063061877e3146101d757806306285c69146101d257806315158878146101cd57806315c34d5b146101c8578063181f5a77146101c35780632451a627146101be578063325c868e146101b95780633937306f146101b45780633a49bb49146101af57806345ac924d146101aa5780634ab35b0b146101a5578063514e8cff146101a05780636def4ce71461019b57806379ba5097146101965780637afac3221461019157806382b49eb01461018c5780638da5cb5b1461018757806391a2749a146101825780639b1115e41461017d578063a69c64c014610178578063cdc73d5114610173578063d02641a01461016e578063d63d3af214610169578063d8694ccd14610164578063f2fde38b1461015f5763ffdb4b371461015a57600080fd5b6126db565b6125e7565b612136565b6120fc565b612080565b611feb565b611f14565b611ea5565b611dd5565b611d83565b611c55565b611ae4565b611994565b6118b5565b6116b2565b611647565b611542565b61134f565b610f5a565b610f20565b610e7f565b610db2565b610b3f565b61087b565b6105d8565b610565565b6104a6565b6103bc565b61020e565b73ffffffffffffffffffffffffffffffffffffffff81160361020957565b600080fd5b346102095760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602061026360043561024e816101eb565b6024356044359161025e836101eb565b6128d6565b604051908152f35b6004359067ffffffffffffffff8216820361020957565b6024359067ffffffffffffffff8216820361020957565b359067ffffffffffffffff8216820361020957565b9181601f840112156102095782359167ffffffffffffffff8311610209576020808501948460051b01011161020957565b919082519283825260005b8481106103295750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016102ea565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061037157505050505090565b90919293946020806103ad837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102df565b97019301930191939290610362565b346102095760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576103f361026b565b60243567ffffffffffffffff8111610209576104139036906004016102ae565b6044929192359167ffffffffffffffff831161020957366023840112156102095782600401359167ffffffffffffffff8311610209573660248460061b8601011161020957610475946024610469950192612af2565b6040519182918261033e565b0390f35b35907fffffffff000000000000000000000000000000000000000000000000000000008216820361020957565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361020957807f66792e80000000000000000000000000000000000000000000000000000000006020921490811561053b575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610530565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095773ffffffffffffffffffffffffffffffffffffffff6004356105b5816101eb565b166000526006602052602067ffffffffffffffff60406000205416604051908152f35b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760006040805161061681610749565b828152826020820152015261047560405161063081610749565b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015260405191829182919091604063ffffffff8160608401956bffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761076557604052565b61071a565b6040810190811067ffffffffffffffff82111761076557604052565b6080810190811067ffffffffffffffff82111761076557604052565b60a0810190811067ffffffffffffffff82111761076557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761076557604052565b6040519061080f610220836107be565b565b6040519061080f6040836107be565b67ffffffffffffffff81116107655760051b60200190565b8015150361020957565b359061080f82610838565b359063ffffffff8216820361020957565b359060ff8216820361020957565b359061ffff8216820361020957565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095736602382011215610209578060040135906108d682610820565b906108e460405192836107be565b82825260246102406020840194028201019036821161020957602401925b8184106109145761091283612d58565b005b8336036102408112610209576102207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0604051926109518461076a565b61095a88610299565b8452011261020957610240916020916109716107ff565b61097c848901610842565b815261098a6040890161084d565b8482015261099a6060890161084d565b60408201526109ab6080890161084d565b60608201526109bc60a0890161085e565b60808201526109cd60c0890161085e565b60a08201526109de60e0890161086c565b60c08201526109f0610100890161084d565b60e0820152610a02610120890161086c565b610100820152610a15610140890161086c565b610120820152610a286101608901610479565b610140820152610a3b610180890161086c565b610160820152610a4e6101a0890161084d565b610180820152610a616101c0890161084d565b6101a0820152610a746101e08901610299565b6101c0820152610a87610200890161084d565b6101e0820152610a9a610220890161084d565b61020082015283820152815201930192610902565b81601f8201121561020957803590610ac682610820565b92610ad460405194856107be565b82845260208085019360061b8301019181831161020957602001925b828410610afe575050505090565b6040848303126102095760206040918251610b188161076a565b610b2187610299565b815282870135610b30816101eb565b83820152815201930192610af0565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095736602382011215610209578060040135610b9981610820565b91610ba760405193846107be565b8183526024602084019260051b820101903682116102095760248101925b828410610bf6576024358567ffffffffffffffff821161020957610bf0610912923690600401610aaf565b90613069565b833567ffffffffffffffff811161020957820160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc82360301126102095760405190610c428261076a565b610c4e60248201610299565b8252604481013567ffffffffffffffff811161020957602491010136601f82011215610209578035610c7f81610820565b91610c8d60405193846107be565b818352602060a081850193028201019036821161020957602001915b818310610cc85750505091816020938480940152815201930192610bc5565b82360360a081126102095760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192610d038461076a565b8635610d0e816101eb565b845201126102095760a091602091604051610d2881610786565b610d3384880161084d565b8152610d416040880161084d565b84820152610d516060880161084d565b60408201526080870135610d6481610838565b606082015283820152815201920191610ca9565b67ffffffffffffffff811161076557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576104756040805190610df381836107be565b601382527f46656551756f74657220312e362e332d646576000000000000000000000000006020830152519182916020835260208301906102df565b602060408183019282815284518094520192019060005b818110610e535750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610e46565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610f0a5761047585610efe818703826107be565b60405191829182610e2f565b8254845260209093019260019283019201610ee7565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602060405160248152f35b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff811161020957806004019060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261020957610fd4614551565b610fde82806133ad565b4263ffffffff1692915060005b8181106111b85750506024019061100282846133ad565b92905060005b83811061101157005b8061103061102b600193611025868a6133ad565b906129b2565b613461565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff61117f61115c602085019461114e61108e87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b6110bd611099610811565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092168252565b63ffffffff8c1660208201526110f86110de845167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b815160209092015160e01b7fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5167ffffffffffffffff1690565b93517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201611008565b806111d16111cc60019361102589806133ad565b61342a565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a73ffffffffffffffffffffffffffffffffffffffff6112b361115c602085019461129961123b87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b611246611099610811565b63ffffffff8d1660208201526110f8611273845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b5173ffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610feb565b9181601f840112156102095782359167ffffffffffffffff8311610209576020838186019501011161020957565b9261134c949261133e928552151560208501526080604085015260808401906102df565b9160608184039101526102df565b90565b346102095760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761138661026b565b60243590611393826101eb565b6044359160643567ffffffffffffffff8111610209576113b79036906004016112ec565b93909160843567ffffffffffffffff8111610209576113da9036906004016112ec565b9290917f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff8216146000146114c0575050935b6bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680861161148f57509161148093916104759693614595565b9093916040519485948561131a565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b916114ca926128d6565b93611439565b602060408183019282815284518094520192019060005b8181106114f45750505090565b9091926020604082611537600194885163ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b0194019291016114e7565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff8111610209576115919036906004016102ae565b61159a81610820565b916115a860405193846107be565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06115d583610820565b0160005b81811061163057505060005b82811015611622576001906116066116018260051b85016129c7565b613ac0565b6116108287612ade565b5261161b8186612ade565b50016115e5565b6040518061047586826114d0565b60209061163b613486565b828288010152016115d9565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602061168c600435611687816101eb565b613bb6565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff60405191168152f35b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095767ffffffffffffffff6116f261026b565b6116fa613486565b501660005260046020526040600020604051906117168261076a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c6020820152604051809161047582604081019263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b61080f909291926102008061022083019561179e84825115159052565b60208181015163ffffffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015160ff169085015260a08181015160ff169085015260c08181015161ffff169085015260e08181015163ffffffff16908501526101008181015161ffff16908501526101208181015161ffff1690850152610140818101517fffffffff0000000000000000000000000000000000000000000000000000000016908501526101608181015161ffff16908501526101808181015163ffffffff16908501526101a08181015163ffffffff16908501526101c08181015167ffffffffffffffff16908501526101e08181015163ffffffff1690850152015163ffffffff16910152565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576104756119886119836118f561026b565b60006102006119026107ff565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e0820152015267ffffffffffffffff166000526007602052604060002090565b61349f565b60405191829182611781565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760005473ffffffffffffffffffffffffffffffffffffffff81163303611a53577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f83011215610209578135611a9481610820565b92611aa260405194856107be565b81845260208085019260051b82010192831161020957602001905b828210611aca5750505090565b602080918335611ad9816101eb565b815201910190611abd565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff811161020957611b33903690600401611a7d565b60243567ffffffffffffffff811161020957611b53903690600401611a7d565b90611b5c613f4b565b60005b8151811015611bd85780611b80611b7b61129960019486612ade565b61571c565b611b8b575b01611b5f565b73ffffffffffffffffffffffffffffffffffffffff611bad6112998386612ade565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2611b85565b8260005b81518110156109125780611bfd611bf861129960019486612ade565b61573d565b611c08575b01611bdc565b73ffffffffffffffffffffffffffffffffffffffff611c2a6112998386612ade565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611c02565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957610475611cfc611c9261026b565b67ffffffffffffffff60243591611ca8836101eb565b60006060604051611cb881610786565b828152826020820152826040820152015216600052600860205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b60ff60405191611d0b83610786565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015260405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261020957604051611e4e8161076a565b816004013567ffffffffffffffff811161020957611e729060043691850101611a7d565b8152602482013567ffffffffffffffff811161020957610912926004611e9b9236920101611a7d565b60208201526135f5565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff811161020957611f00611efa6104759236906004016112ec565b90613934565b6040519182916020835260208301906102df565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff8111610209573660238201121561020957806004013590611f6f82610820565b90611f7d60405192836107be565b8282526024602083019360061b8201019036821161020957602401925b818410611faa57610912836139e4565b6040843603126102095760206040918251611fc48161076a565b8635611fcf816101eb565b8152611fdc838801610299565b83820152815201930192611f9a565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760405180602060095491828152019060096000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af9060005b81811061206a5761047585610efe818703826107be565b8254845260209093019260019283019201612053565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760406120c0600435611601816101eb565b6120fa8251809263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565bf35b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602060405160128152f35b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761216d61026b565b60243567ffffffffffffffff811161020957806004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8336030112610209576121d16119838267ffffffffffffffff166000526007602052604060002090565b916121e36121df8451151590565b1590565b6125af5760648101926122276121df6121fb866129c7565b73ffffffffffffffffffffffffffffffffffffffff166000526001600901602052604060002054151590565b61255d57604482019461223a86826133ad565b9490506122488284836148e2565b91612255611687886129c7565b968461227261226c6101e083015163ffffffff1690565b85615015565b996000808a156125255750509161231a61230d6122cd61ffff856122ff996122d999986122c46122bd6101806122b16101606123209d015161ffff1690565b95015163ffffffff1690565b918b6133ad565b94909316615100565b989198979098946129c7565b73ffffffffffffffffffffffffffffffffffffffff166000526006602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b9061288a565b9560009761ffff61233761012089015161ffff1690565b166124c7575b509461231a61230d6101c061242d6104759d6dffffffffffffffffffffffffffff6124256124459f9e9b6124207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9f9b9c61243d9f6124209e63ffffffff6123ab6124209f60246123b5950190612a11565b9290501690613b51565b908b608081016123d86123d26123cc835160ff1690565b60ff1690565b8561288a565b9360c08301916123ea835161ffff1690565b9061ffff82168311612455575b5050505060600151612420916124179163ffffffff16613b9c565b613b9c565b63ffffffff1690565b613b51565b91169061288a565b93015167ffffffffffffffff1690565b91169061289d565b6040519081529081906020820190565b612417949650612420959361ffff6124b66124a56124129661249f61249861248f60609960ff6124896124bd9b5160ff1690565b16613b5e565b965161ffff1690565b61ffff1690565b90613b77565b61231a6123cc60a08d015160ff1690565b9116613b51565b95938395506123f7565b909594989750826124ed8b989495986dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff169161250b6024890185612a11565b905061251793886151fa565b96979394389693929661233d565b94939150945061231a61230d6122ff6122d961255761255261241761020061232099015163ffffffff1690565b6127fd565b946129c7565b6125ab612569856129c7565b7f2502348c0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6000fd5b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095773ffffffffffffffffffffffffffffffffffffffff600435612637816101eb565b61263f613f4b565b163381146126b157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957600435612716816101eb565b61271e610282565b9067ffffffffffffffff82169182600052600760205260ff60406000205416156127a05761274e61276f92613bb6565b92600052600760205263ffffffff60016040600020015460801c1690615015565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc10000149015171561281f57565b6127ce565b908160051b918083046020149015171561281f57565b9061012c82029180830461012c149015171561281f57565b9061010c82029180830461010c149015171561281f57565b90655af3107a4000820291808304655af3107a4000149015171561281f57565b8181029291811591840414171561281f57565b81156128a7570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b61291561290f61134c94937bffffffffffffffffffffffffffffffffffffffffffffffffffffffff6129088195613bb6565b169061288a565b92613bb6565b169061289d565b9061292682610820565b61293360405191826107be565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129618294610820565b019060005b82811061297257505050565b806060602080938501015201612966565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156129c25760061b0190565b612983565b3561134c816101eb565b91908110156129c25760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610209570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610209570180359067ffffffffffffffff82116102095760200191813603831361020957565b929192612a6e82610d78565b91612a7c60405193846107be565b829481845281830111610209578281602093846000960137010152565b90604051612aa681610786565b606060ff82945463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c166040850152821c161515910152565b80518210156129c25760209160051b010190565b909291612b3f612b168367ffffffffffffffff166000526007602052604060002090565b5460181b7fffffffff000000000000000000000000000000000000000000000000000000001690565b90612b498161291c565b9560005b828110612b5e575050505050505090565b612b71612b6c8284896129b2565b6129c7565b8388612b8b612b818584846129d1565b6040810190612a11565b905060208111612cd0575b508392612bcc612bc6612bbf612bb5600198612c1497612c0f976129d1565b6020810190612a11565b3691612a62565b89613c65565b612bea8967ffffffffffffffff166000526008602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b612a99565b606081015115612c9757612c7b612c356020612c4f93015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107be565b612c85828b612ade565b52612c90818a612ade565b5001612b4d565b50612c4f612c7b612ccb84612cc08a67ffffffffffffffff166000526007602052604060002090565b015463ffffffff1690565b612c35565b915050612d08612417612cfb84612bea8b67ffffffffffffffff166000526008602052604060002090565b5460401c63ffffffff1690565b10612d1557838838612b96565b7f36f536ca0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b90612d61613f4b565b60005b825181101561306457612d778184612ade565b516020612d96612d878487612ade565b515167ffffffffffffffff1690565b9101519067ffffffffffffffff811680158015613045575b8015613017575b8015612ed7575b612e9f5791612e658260019594612e15612df0612b16612e6a9767ffffffffffffffff166000526007602052604060002090565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b612e70577fe21664081fc677b92cffc05bf1c3609b90b261aafb6b647b8dae2a01d1c76eb760405180612e488782611781565b0390a267ffffffffffffffff166000526007602052604060002090565b613f96565b01612d64565b7f0b24b226620af755f455598afdfb191eaf1236f264c327646ac8644608d7d31a60405180612e488782611781565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b507fffffffff00000000000000000000000000000000000000000000000000000000612f276101408501517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c000000000000000000000000000000000000000000000000000000008114159081612fec575b81612fc1575b81612f96575b81612f6b575b50612dbc565b7f647e2ba9000000000000000000000000000000000000000000000000000000009150141538612f65565b7fc4e05953000000000000000000000000000000000000000000000000000000008114159150612f5f565b7fac77ffec000000000000000000000000000000000000000000000000000000008114159150612f59565b7f1e10bdc4000000000000000000000000000000000000000000000000000000008114159150612f53565b506101a083015163ffffffff1663ffffffff61303d612417604087015163ffffffff1690565b911611612db5565b5063ffffffff61305d6101a085015163ffffffff1690565b1615612dae565b509050565b90613072613f4b565b6000915b80518310156132ee576130898382612ade565b519061309d825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b855180518210156132d9576130c982602092612ade565b5101516130f56130da838951612ade565b515173ffffffffffffffffffffffffffffffffffffffff1690565b604082015163ffffffff166020811061328b575090867f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed173ffffffffffffffffffffffffffffffffffffffff84613238858f60019998612bea61316c9267ffffffffffffffff166000526008602052604060002090565b815181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff919091161781559060208181015183547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16911b67ffffffff000000001617825560408181015183546060938401517fffffffffffffffffffffffffffffffffffffff0000000000ffffffffffffffff919091169190921b6bffffffff0000000000000000161790151590911b6cff00000000000000000000000016179055565b613282604051928392169582919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b0390a3016130b2565b7f24ecdc020000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff90911660045263ffffffff1660245260446000fd5b50509550925092600191500191929092613076565b50905060005b81518110156133a9578061330d612d8760019385612ade565b67ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff613356602061333a8689612ade565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b600061337a82612bea8767ffffffffffffffff166000526008602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a3016132f4565b5050565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610209570180359067ffffffffffffffff821161020957602001918160061b3603831361020957565b35907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216820361020957565b604081360312610209576134596020604051926134468461076a565b8035613451816101eb565b845201613401565b602082015290565b6040813603126102095761345960206040519261347d8461076a565b61345181610299565b604051906134938261076a565b60006020838281520152565b9061080f6135e760016134b06107ff565b9461359561358982546134cc6134c68260ff1690565b15158a52565b63ffffffff600882901c1660208a015263ffffffff602882901c1660408a015263ffffffff604882901c1660608a015260ff606882901c1660808a015260ff607082901c1660a08a015261ffff607882901c1660c08a015263ffffffff608882901c1660e08a015261ffff60a882901c166101008a015261ffff60b882901c166101208a01527fffffffff00000000000000000000000000000000000000000000000000000000601882901b166101408a015260e81c61ffff1690565b61ffff16610160880152565b015463ffffffff811661018086015263ffffffff602082901c166101a086015267ffffffffffffffff604082901c166101c086015263ffffffff608082901c166101e086015260a01c63ffffffff1690565b63ffffffff16610200840152565b6135fd613f4b565b60208101519160005b83518110156136b1578061361f61129960019387612ade565b61365b61365673ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b615b81565b613667575b5001613606565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613660565b5091505160005b81518110156133a9576136ce6112998284612ade565b9073ffffffffffffffffffffffffffffffffffffffff82161561376b577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6137628361373a61373561363d60019773ffffffffffffffffffffffffffffffffffffffff1690565b615b08565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a1016136b8565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b909291928360041161020957831161020957600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b906004116102095790600490565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613812575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9080601f8301121561020957813561385b81610820565b9261386960405194856107be565b81845260208085019260051b82010192831161020957602001905b8282106138915750505090565b8135815260209182019101613884565b6020818303126102095780359067ffffffffffffffff8211610209570160a08183031261020957604051916138d5836107a2565b6138de8261084d565b83526138ec60208301610299565b602084015260408201356138ff81610838565b604084015260608201356060840152608082013567ffffffffffffffff81116102095761392c9201613844565b608082015290565b9060048110801561398f575b61397857606061396261395a8361134c94612c4f96613795565b8101906138a1565b0151604051928391602083019190602083019252565b50506040516139886020826107be565b6000815290565b50806004116102095781357fffffffff00000000000000000000000000000000000000000000000000000000167f1f3b3aba000000000000000000000000000000000000000000000000000000001415613940565b6139ec613f4b565b60005b81518110156133a9578073ffffffffffffffffffffffffffffffffffffffff613a1a60019385612ade565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d613ab767ffffffffffffffff6020613a568689612ade565b51015116836000526006602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a2016139ef565b73ffffffffffffffffffffffffffffffffffffffff90613ade613486565b50166000526005602052604060002060405190613afa8261076a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c602082015290565b906002820180921161281f57565b906020820180921161281f57565b906001820180921161281f57565b9190820180921161281f57565b9061ffff8091169116029061ffff821691820361281f57565b9190820391821161281f57565b63ffffffff60209116019063ffffffff821161281f57565b9063ffffffff8091169116019063ffffffff821161281f57565b613bbf81613ac0565b9063ffffffff602083015116158015613c3e575b613bfa5750517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff82511615613bd3565b7fffffffff000000000000000000000000000000000000000000000000000000001691907f2812d52c000000000000000000000000000000000000000000000000000000008314613dac577f1e10bdc4000000000000000000000000000000000000000000000000000000008314613d9e577fac77ffec000000000000000000000000000000000000000000000000000000008314613d93577f647e2ba9000000000000000000000000000000000000000000000000000000008314613d88577fc4e05953000000000000000000000000000000000000000000000000000000008314613d7a57827f2ee820750000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61080f919250600b906153a0565b61080f919250615401565b61080f91925061533d565b61080f9192506001906153a0565b61080f9192506152ad565b917fffffffff0000000000000000000000000000000000000000000000000000000083167f2812d52c000000000000000000000000000000000000000000000000000000008114613f3f577f1e10bdc4000000000000000000000000000000000000000000000000000000008114613f1f577fac77ffec000000000000000000000000000000000000000000000000000000008114613f13577f647e2ba9000000000000000000000000000000000000000000000000000000008114613f07577fc4e059530000000000000000000000000000000000000000000000000000000014613eed577f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000831660045260246000fd5b61080f925015613eff57600b906153a0565b6000906153a0565b505061080f9150615401565b505061080f915061533d565b5061080f925015613f365760ff60015b16906153a0565b60ff6000613f2f565b505061080f91506152ad565b73ffffffffffffffffffffffffffffffffffffffff600154163303613f6c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90614509610200600161080f94613fe1613fb08651151590565b829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b61402b613ff5602087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ff1660089190911b64ffffffff0016178255565b61407961403f604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1660289190911b68ffffffff000000000016178255565b6140cb61408d606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffff1660489190911b6cffffffff00000000000000000016178255565b61411b6140dc608087015160ff1690565b82547fffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffff1660689190911b6dff0000000000000000000000000016178255565b61416c61412c60a087015160ff1690565b82547fffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffff1660709190911b6eff000000000000000000000000000016178255565b6141c061417e60c087015161ffff1690565b82547fffffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffff1660789190911b70ffff00000000000000000000000000000016178255565b61421a6141d460e087015163ffffffff1690565b82547fffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffff1660889190911b74ffffffff000000000000000000000000000000000016178255565b61427561422d61010087015161ffff1690565b82547fffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffff1660a89190911b76ffff00000000000000000000000000000000000000000016178255565b6142d261428861012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b6143516143036101408701517fffffffff000000000000000000000000000000000000000000000000000000001690565b82547fffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffff1660189190911c7cffffffff0000000000000000000000000000000000000000000000000016178255565b6143b461436461016087015161ffff1690565b82547fff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e89190911b7effff000000000000000000000000000000000000000000000000000000000016178255565b01926143fc6143cb61018083015163ffffffff1690565b859063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b61444a6144116101a083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178555565b6144a46144636101c083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff1660409190911b6fffffffffffffffff000000000000000016178555565b6144fe6144b96101e083015163ffffffff1690565b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1660809190911b73ffffffff0000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff000000000000000000000000000000000000000083549260a01b169116179055565b3360005260036020526040600020541561456757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6119836145bc9196949395929667ffffffffffffffff166000526007602052604060002090565b946101408601947fffffffff0000000000000000000000000000000000000000000000000000000061460e87517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c0000000000000000000000000000000000000000000000000000000081149081156148a9575b811561487f575b5061483a577fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006146b288517fffffffff000000000000000000000000000000000000000000000000000000001690565b161461481c5750507f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061472686517fffffffff000000000000000000000000000000000000000000000000000000001690565b16146147a7576125ab61475985517fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b6147e99350612bbf60606147d36147cc6124176040614815989a015163ffffffff1690565b84866155e8565b0151604051958691602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018652856107be565b9160019190565b9450945061134c9161482f913691612a62565b936001933691612a62565b945094916148609161485a6124176101a061134c96015163ffffffff1690565b9161547d565b93614877602061486f8761559e565b960151151590565b933691612a62565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501438614641565b7fac77ffec000000000000000000000000000000000000000000000000000000008114915061463a565b90816020910312610209573590565b91906148f16020830183612a11565b9390506040830161490281856133ad565b9050602084019161491a612417845163ffffffff1690565b808811614fe3575060018211614faf5761014085019661495a88517fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f2812d52c0000000000000000000000000000000000000000000000000000000081148015614f86575b8015614f5d575b15614a3e57505050505050509181614a38612bbf614a07614a31966149d9608061134c980186612a11565b614a0161241760406149f56101a0879697015163ffffffff1690565b94015163ffffffff1690565b926158a9565b51958694517fffffffff000000000000000000000000000000000000000000000000000000001690565b9280612a11565b90613db7565b7fc4e0595300000000000000000000000000000000000000000000000000000000819b9a939495979996989b14600014614ce1575050614ae1614aa8614ad4999a6040614aa2612417614a9460808b018b612a11565b939094015163ffffffff1690565b9161575e565b918251998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b614a38612bbf8880612a11565b6060810151519082614afe614af68780612a11565b8101906148d3565b614cc4575081614c8e575b8515159081614c81575b50614c575760408111614c255750614b3890614b328594939795612852565b90613b51565b946000935b838510614b96575050505050612417614b5a915163ffffffff1690565b808211614b6657505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614bfa612417612cfb614bc58667ffffffffffffffff166000526008602052604060002090565b614bd6612b6c8d6110258b8d6133ad565b73ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b8015614c1557614c0991613b51565b965b0193929190614b3d565b50614c1f90613b35565b96614c0b565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614b13565b6125ab827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614cdb919350614b32614cd684613b43565b612824565b91614b09565b7f1e10bdc40000000000000000000000000000000000000000000000000000000003614f0e5750614d69614d2e614ad4999a6040614d28612417614a9460808b018b612a11565b916155e8565b91614d40612417845163ffffffff1690565b998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b6080810151519082614d7e614af68780612a11565b614ef6575081614ec0575b85151580614eb4575b614c575760408211614e80576020015167ffffffffffffffff9081169081831c16614e46575050614dca90614b32859493979561283a565b946000935b838510614dec575050505050612417614b5a915163ffffffff1690565b9091929395600190614e1b612417612cfb614bc58667ffffffffffffffff166000526008602052604060002090565b8015614e3657614e2a91613b51565b965b0193929190614dcf565b50614e4090613b35565b96614e2c565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614d92565b6125ab827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614f08919350614b32614cd684613b27565b91614d89565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b507f647e2ba90000000000000000000000000000000000000000000000000000000081146149ae565b507fac77ffec0000000000000000000000000000000000000000000000000000000081146149a7565b7fd88dddd6000000000000000000000000000000000000000000000000000000006000526004829052600160245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff811660005260046020526040600020916040519261503b8461076a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116845260e01c9182602085015263ffffffff8216928361509f575b5050505061134c90517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b63ffffffff16429081039390841161281f5783116150bd5780615075565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b9194929390946000926000966000966000945b80861061512557505050505050929190565b909192939495615160612c0f615152889c9b9c67ffffffffffffffff166000526008602052604060002090565b614bd6612b6c8b87896129b2565b986151716121df60608c0151151590565b6151cf5760208a015163ffffffff1661518991613b9c565b9960408a015161519c9063ffffffff1690565b6151a591613b9c565b98516151b69063ffffffff166127fd565b6151bf91613b51565b956001905b019493929190615113565b9698506151f46151ee856151e8600194614b328a6127fd565b99613b9c565b9a613b84565b986151c4565b91939093806101e00193846101e01161281f57610120810290808204610120149015171561281f576101e091010180931161281f5761249861012061528f61134c966dffffffffffffffffffffffffffff61242561527b6152686152999a63ffffffff61231a9a1690613b51565b61231a6124986101008c015161ffff1690565b614b3261241760e08b015163ffffffff1690565b93015161ffff1690565b61286a565b90816020910312610209575190565b60208151036152f0576152c9602082518301016020830161529e565b73ffffffffffffffffffffffffffffffffffffffff8111908115615331575b506152f05750565b61532d906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b0390fd5b610400915010386152e8565b602081510361536357600b61535b602083518401016020840161529e565b106153635750565b61532d906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b9060208251036153c657806153b3575050565b61535b602083518401016020840161529e565b6040517fe0d7fb02000000000000000000000000000000000000000000000000000000008152602060048201528061532d60248201856102df565b6024815103615417576022810151156154175750565b61532d906040519182917f373b0e4400000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b908160409103126102095760206040519161546e8361076a565b80518352015161345981610838565b91615486613486565b50811561557c57506154c7612bbf82806154c17fffffffff0000000000000000000000000000000000000000000000000000000095876137de565b95613795565b91167f181dcf1000000000000000000000000000000000000000000000000000000000810361550457508060208061134c93518301019101615454565b7f97a657c90000000000000000000000000000000000000000000000000000000014615554577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b806020806155679351830101910161529e565b61556f610811565b9081526000602082015290565b91505067ffffffffffffffff615590610811565b911681526000602082015290565b6020604051917f181dcf100000000000000000000000000000000000000000000000000000000082840152805160248401520151151560448201526044815261134c6064826107be565b606060806040516155f8816107a2565b600081526000602082015260006040820152600083820152015281156156f2577f1f3b3aba000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061566d61566785856137d0565b906137de565b16036156c857816156819261395a92613795565b9063ffffffff615695835163ffffffff1690565b161161569e5790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff61134c911660096159b5565b73ffffffffffffffffffffffffffffffffffffffff61134c91166009615b43565b9060608060405161576e81610786565b600081526000602082015260006040820152015280156156f2577f21ea4ca9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006157d761566784866137d0565b16036156c857806157e792613795565b819291016000926020818303126158a55780359067ffffffffffffffff82116158a157019260808483031261589e576040519361582385610786565b80358552602081013561583581610838565b60208601526040810135604086015260608101359167ffffffffffffffff831161589e5750615865929101613844565b60608301528151116158745790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b80fd5b8480fd5b8380fd5b9063ffffffff6158c3936158bb613486565b50169161547d565b908151116158745790565b80548210156129c25760005260206000200190600090565b9161591e918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b80548015615986577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061595782826158ce565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615aa3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840184811161281f578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff850194851161281f576000958583615a5497615a459503615a5a575b505050615922565b90600052602052604060002090565b55600190565b615a8a615a8491615a7b615a71615a9a95886158ce565b90549060031b1c90565b928391876158ce565b906158e6565b8590600052602052604060002090565b55388080615a3d565b50505050600090565b805490680100000000000000008210156107655781615ad391600161591e940181556158ce565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b600081815260036020526040902054615b3d57615b26816002615aac565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615b7a5780615b6583600193615aac565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615b7a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161281f57600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840193841161281f578383615a549460009603615c1d575b505050615c0c6002615922565b600390600052602052604060002090565b615c0c615a8491615c35615a71615c3f9560026158ce565b92839160026158ce565b55388080615bff56fea164736f6c634300081a000a",
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

func (_FeeQuoterV2 *FeeQuoterV2Caller) FEEBASEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "FEE_BASE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoterV2.Contract.FEEBASEDECIMALS(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoterV2.Contract.FEEBASEDECIMALS(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2Caller) KEYSTONEPRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "KEYSTONE_PRICE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoterV2.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoterV2.CallOpts)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoterV2.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoterV2.CallOpts)
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

func (_FeeQuoterV2 *FeeQuoterV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeQuoterV2.Contract.SupportsInterface(&_FeeQuoterV2.CallOpts, interfaceId)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeQuoterV2.Contract.SupportsInterface(&_FeeQuoterV2.CallOpts, interfaceId)
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
	return common.HexToHash("0xe21664081fc677b92cffc05bf1c3609b90b261aafb6b647b8dae2a01d1c76eb7")
}

func (FeeQuoterV2DestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x0b24b226620af755f455598afdfb191eaf1236f264c327646ac8644608d7d31a")
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
	FEEBASEDECIMALS(opts *bind.CallOpts) (*big.Int, error)

	KEYSTONEPRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error)

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

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

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
