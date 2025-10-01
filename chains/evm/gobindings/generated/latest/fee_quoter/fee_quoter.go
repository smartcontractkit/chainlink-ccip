// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_quoter

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
	MaxNumberOfTokensPerMsg           uint16
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

var FeeQuoterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_BASE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e060405234610ed357616fe08038038061001981611089565b92833981018181036101008112610ed357606013610ed35760405191606083016001600160401b03811184821017610ed85760405280516001600160601b0381168103610ed357835261006e602082016110ae565b9160208401928352610082604083016110c2565b6040850190815260608301519094906001600160401b038111610ed357826100ab9185016110ea565b60808401519095906001600160401b038111610ed357836100cd9186016110ea565b60a08501519093906001600160401b038111610ed35785019481601f87011215610ed357855195610105610100886110d3565b611089565b9660208089838152019160051b83010191848311610ed35760208101915b838310610f2a575050505060c08101516001600160401b038111610ed35781019082601f83011215610ed357815161015d610100826110d3565b9260208085848152019260061b82010190858211610ed357602001915b818310610eee5750505060e0810151906001600160401b038211610ed3570182601f82011215610ed3578051906101b3610100836110d3565b936020610260818786815201940283010191818311610ed357602001925b828410610d2457505050503315610d1357600180546001600160a01b0319163317905560209561020087611089565b9560008752600036813761021261106a565b998a5286888b015260005b8751811015610284576001906001600160a01b0361023b828b611183565b51168a6102478261136f565b610254575b50500161021d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a61024c565b50879650888a519560005b8751811015610300576001600160a01b036102aa828a611183565b51169081156102ef577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8b836102e16001956112f7565b50604051908152a10161028f565b6342bcdf7f60e11b60005260046000fd5b50815189975088906001600160a01b0316158015610d01575b8015610cf2575b610ce15791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c05261035385611089565b9260008452600036813760005b84518110156103cf576001906103886001600160a01b036103818389611183565b5116611204565b610393575b01610360565b818060a01b036103a38288611183565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a261038d565b508593508460005b8351811015610446576001906103ff6001600160a01b036103f88388611183565b5116611336565b61040a575b016103d7565b818060a01b0361041a8287611183565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610404565b5091509160005b8251811015610a0a576104608184611183565b51856001600160401b036104748487611183565b51511691015190801580156109f7575b80156109d9575b8015610957575b61094357600081815260078852604090205460019392919060081b6001600160e01b03191661080957807faae4c0e87e0ed4dd3178367ccd926daebfe7fdb4d86913f3b555d9620622b961604051806105ed868291909161022063ffffffff8161024084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b6101608201511661016086015261ffff61018082015116610180860152826101a0820151166101a0860152826101c0820151166101c086015260018060401b036101e0820151166101e08601528261020082015116610200860152015116910152565b0390a25b600052600787528260406000208251151583825462ffff008c83015160081b1666ffffffff000000604084015160181b16936affffffff00000000000000606085015160381b166effffffff0000000000000000000000608086015160581b169560ff60781b60a087015160781b169260ff60801b60c088015160801b1661ffff60881b60e089015160881b169263ffffffff60981b6101008a015160981b169561ffff60b81b6101208b015160b81b169861ffff60c81b6101408c015160c81b169a63ffffffff60d81b9061016063ffffffff60d81b91015160081c16169b63ffffffff60d81b199961ffff60c81b199861ffff60b81b199763ffffffff60981b199661ffff60881b199560ff60801b199460ff60781b19936effffffff0000000000000000000000199260ff6affffffff000000000000001992169066ffffffffffffff19161716171617161716171617161716171617161717178155019061ffff610180820151169082549165ffffffff00006101a083015160101b169269ffffffff0000000000006101c084015160301b166a01000000000000000000008860901b036101e085015160501b169263ffffffff60901b61020086015160901b169461022063ffffffff60b01b91015160b01b169563ffffffff60b01b199363ffffffff60901b19926a01000000000000000000008c60901b0319918c8060501b03191617161716171617171790550161044d565b807f425825c9dae839e14fe8f6b91cc8aefe0131bfff74c5958aec031bad9566c4f36040518061093b868291909161022063ffffffff8161024084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b6101608201511661016086015261ffff61018082015116610180860152826101a0820151166101a0860152826101c0820151166101c086015260018060401b036101e0820151166101e08601528261020082015116610200860152015116910152565b0390a26105f1565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61016083015116630a04b54b60e21b81141590816109c7575b816109b5575b816109a3575b81610991575b50610492565b63647e2ba960e01b141590508861098b565b63c4e0595360e01b8114159150610985565b632b1dfffb60e21b811415915061097f565b6307842f7160e21b8114159150610979565b5063ffffffff6101c08301511663ffffffff6060840151161061048b565b5063ffffffff6101c08301511615610484565b84828560005b8151811015610a90576001906001600160a01b03610a2e8285611183565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610a648689611183565b510151168360005260068252604060002081878060401b0319825416179055604051908152a201610a10565b50506001610a9d83611089565b9160008352600091610cdc575b81925b8151841015610c2157610ac08483611183565b5180516001600160401b0316929086019190845b8784518051831015610c105782610aea91611183565b51015184516001600160a01b0390610b03908490611183565b51511690604081019063ffffffff8251168b8110610bf9575091877f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed160808d97969460019663ffffffff60408f87815260088d528181208b8060a01b038a1682528d522092818c8185511695805467ffffffff00000000838801938451901b169660606bffffffff0000000000000000875160401b169101976cff0000000000000000000000008951151560601b16928a6cff000000000000000000000000199160018060601b0319161716171717905560405195865251168c850152511660408301525115156060820152a301909150610ad4565b6312766e0160e11b8a526004849052602452604489fd5b505050925093600191500192610aad565b905083825b8251811015610ca4576001906001600160401b03610c448286611183565b515116828060a01b0384610c588488611183565b51015116908087526008855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610c26565b604051615bdc9081611404823960805181818161062f01526111af015260a0518181816106680152611146015260c051816106940152f35b610aaa565b63d794ef9560e01b60005260046000fd5b5063ffffffff82511615610320565b5080516001600160601b031615610319565b639b15e16f60e01b60005260046000fd5b8382036102608112610ed357610240610d3b61106a565b91610d4587611145565b8352601f190112610ed3576040519161024083016001600160401b03811184821017610ed857604052610d7a60208701611159565b8352610d8860408701611166565b6020840152610d99606087016110c2565b6040840152610daa608087016110c2565b6060840152610dbb60a087016110c2565b6080840152610dcc60c08701611175565b60a0840152610ddd60e08701611175565b60c0840152610def6101008701611166565b60e0840152610e0161012087016110c2565b610100840152610e146101408701611166565b610120840152610e276101608701611166565b610140840152610180860151916001600160e01b031983168303610ed35783602093610160610260960152610e5f6101a08901611166565b610180820152610e726101c089016110c2565b6101a0820152610e856101e089016110c2565b6101c0820152610e986102008901611145565b6101e0820152610eab61022089016110c2565b610200820152610ebe61024089016110c2565b610220820152838201528152019301926101d1565b600080fd5b634e487b7160e01b600052604160045260246000fd5b604083870312610ed3576020604091610f0561106a565b610f0e866110ae565b8152610f1b838701611145565b8382015281520192019161017a565b82516001600160401b038111610ed35782016040818803601f190112610ed357610f5261106a565b90610f5f60208201611145565b825260408101516001600160401b038111610ed357602091010187601f82011215610ed3578051610f92610100826110d3565b91602060a08185858152019302820101908a8211610ed357602001915b818310610fce5750505091816020938480940152815201920191610123565b828b0360a08112610ed3576080610fe361106a565b91610fed866110ae565b8352601f190112610ed3576040519160808301916001600160401b03831184841017610ed85760a0936020936040526110278488016110c2565b8152611035604088016110c2565b84820152611045606088016110c2565b604082015261105660808801611159565b606082015283820152815201920191610faf565b60408051919082016001600160401b03811183821017610ed857604052565b6040519190601f01601f191682016001600160401b03811183821017610ed857604052565b51906001600160a01b0382168203610ed357565b519063ffffffff82168203610ed357565b6001600160401b038111610ed85760051b60200190565b9080601f83011215610ed3578151611104610100826110d3565b9260208085848152019260051b820101928311610ed357602001905b82821061112d5750505090565b6020809161113a846110ae565b815201910190611120565b51906001600160401b0382168203610ed357565b51908115158203610ed357565b519061ffff82168203610ed357565b519060ff82168203610ed357565b80518210156111975760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156111975760005260206000200190600090565b805480156111ee5760001901906111dc82826111ad565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600a602052604090205480156112c55760001981018181116112af576009546000198101919082116112af5781810361125e575b50505061124a60096111c5565b600052600a60205260006040812055600190565b61129761126f6112809360096111ad565b90549060031b1c92839260096111ad565b819391549060031b91821b91600019901b19161790565b9055600052600a60205260406000205538808061123d565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80549068010000000000000000821015610ed857816112809160016112f3940181556111ad565b9055565b80600052600360205260406000205415600014611330576113198160026112cc565b600254906000526003602052604060002055600190565b50600090565b80600052600a60205260406000205415600014611330576113588160096112cc565b60095490600052600a602052604060002055600190565b60008181526003602052604090205480156112c55760001981018181116112af576002546000198101919082116112af578082036113c9575b5050506113b560026111c5565b600052600360205260006040812055600190565b6113eb6113da6112809360026111ad565b90549060031b1c92839260026111ad565b905560005260036020526040600020553880806113a856fe6080604052600436101561001257600080fd5b60003560e01c806241e5be146101d657806301447eaa146101d157806301ffc9a7146101cc578063061877e3146101c757806306285c69146101c257806315c34d5b146101bd578063181f5a77146101b85780632451a627146101b3578063325c868e146101ae5780633937306f146101a95780633a49bb49146101a457806345ac924d1461019f5780634ab35b0b1461019a578063514e8cff1461019557806359e7fdbe146101905780636def4ce71461018b57806379ba5097146101865780637afac3221461018157806382b49eb01461017c5780638da5cb5b1461017757806391a2749a14610172578063a69c64c01461016d578063cdc73d5114610168578063d02641a014610163578063d63d3af21461015e578063d8694ccd14610159578063f2fde38b146101545763ffdb4b371461014f57600080fd5b61264b565b612557565b6120a6565b61206c565b611ff0565b611f5b565b611e84565b611db4565b611d62565b611c34565b611ac3565b611973565b61188d565b611505565b611419565b6113ae565b6112a9565b6110b6565b610cc1565b610c87565b610be6565b610b19565b6108de565b6105c8565b610555565b610496565b6103ac565b6101fe565b73ffffffffffffffffffffffffffffffffffffffff8116036101f957565b600080fd5b346101f95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957602061025360043561023e816101db565b6024356044359161024e836101db565b612846565b604051908152f35b6004359067ffffffffffffffff821682036101f957565b6024359067ffffffffffffffff821682036101f957565b359067ffffffffffffffff821682036101f957565b9181601f840112156101f95782359167ffffffffffffffff83116101f9576020808501948460051b0101116101f957565b919082519283825260005b8481106103195750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016102da565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061036157505050505090565b909192939460208061039d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102cf565b97019301930191939290610352565b346101f95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576103e361025b565b60243567ffffffffffffffff81116101f95761040390369060040161029e565b6044929192359167ffffffffffffffff83116101f957366023840112156101f95782600401359167ffffffffffffffff83116101f9573660248460061b860101116101f957610465946024610459950192612a90565b6040519182918261032e565b0390f35b35907fffffffff00000000000000000000000000000000000000000000000000000000821682036101f957565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036101f957807f66792e80000000000000000000000000000000000000000000000000000000006020921490811561052b575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610520565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95773ffffffffffffffffffffffffffffffffffffffff6004356105a5816101db565b166000526006602052602067ffffffffffffffff60406000205416604051908152f35b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760006040805161060681610739565b828152826020820152015261046560405161062081610739565b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015260405191829182919091604063ffffffff8160608401956bffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761075557604052565b61070a565b6040810190811067ffffffffffffffff82111761075557604052565b6080810190811067ffffffffffffffff82111761075557604052565b60a0810190811067ffffffffffffffff82111761075557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761075557604052565b604051906107fe6040836107ae565b565b604051906107fe610240836107ae565b67ffffffffffffffff81116107555760051b60200190565b359063ffffffff821682036101f957565b801515036101f957565b35906107fe82610839565b81601f820112156101f95780359061086582610810565b9261087360405194856107ae565b82845260208085019360061b830101918183116101f957602001925b82841061089d575050505090565b6040848303126101f957602060409182516108b78161075a565b6108c087610289565b8152828701356108cf816101db565b8382015281520193019261088f565b346101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f957366023820112156101f957806004013561093881610810565b9161094660405193846107ae565b8183526024602084019260051b820101903682116101f95760248101925b828410610997576024358567ffffffffffffffff82116101f95761098f61099592369060040161084e565b90612cf9565b005b833567ffffffffffffffff81116101f957820160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc82360301126101f957604051906109e38261075a565b6109ef60248201610289565b8252604481013567ffffffffffffffff81116101f957602491010136601f820112156101f9578035610a2081610810565b91610a2e60405193846107ae565b818352602060a08185019302820101903682116101f957602001915b818310610a695750505091816020938480940152815201930192610964565b82360360a081126101f95760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192610aa48461075a565b8635610aaf816101db565b845201126101f95760a091602091604051610ac981610776565b610ad4848801610828565b8152610ae260408801610828565b84820152610af260608801610828565b60408201526080870135610b0581610839565b606082015283820152815201920191610a4a565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576104656040805190610b5a81836107ae565b601382527f46656551756f74657220312e362e332d646576000000000000000000000000006020830152519182916020835260208301906102cf565b602060408183019282815284518094520192019060005b818110610bba5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610bad565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610c715761046585610c65818703826107ae565b60405191829182610b96565b8254845260209093019260019283019201610c4e565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957602060405160248152f35b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f957806004019060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101f957610d3b613cd1565b610d458280613008565b4263ffffffff1692915060005b818110610f1f57505060240190610d698284613008565b92905060005b838110610d7857005b80610d97610d92600193610d8c868a613008565b90612922565b6130bc565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff610ee6610ec36020850194610eb5610df587517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610e24610e006107ef565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092168252565b63ffffffff8c166020820152610e5f610e45845167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b815160209092015160e01b7fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5167ffffffffffffffff1690565b93517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610d6f565b80610f38610f33600193610d8c8980613008565b613085565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a73ffffffffffffffffffffffffffffffffffffffff61101a610ec36020850194611000610fa287517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610fad610e006107ef565b63ffffffff8d166020820152610e5f610fda845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b5173ffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610d52565b9181601f840112156101f95782359167ffffffffffffffff83116101f957602083818601950101116101f957565b926110b394926110a5928552151560208501526080604085015260808401906102cf565b9160608184039101526102cf565b90565b346101f95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576110ed61025b565b602435906110fa826101db565b6044359160643567ffffffffffffffff81116101f95761111e903690600401611053565b93909160843567ffffffffffffffff81116101f957611141903690600401611053565b9290917f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff821614600014611227575050935b6bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168086116111f65750916111e793916104659693613d15565b90939160405194859485611081565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9161123192612846565b936111a0565b602060408183019282815284518094520192019060005b81811061125b5750505090565b909192602060408261129e600194885163ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b01940192910161124e565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f9576112f890369060040161029e565b61130181610810565b9161130f60405193846107ae565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061133c83610810565b0160005b81811061139757505060005b828110156113895760019061136d6113688260051b8501612937565b6137fb565b6113778287612a7c565b526113828186612a7c565b500161134c565b604051806104658682611237565b6020906113a26130e1565b82828801015201611340565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760206113f36004356113ee816101db565b6138f1565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95767ffffffffffffffff61145961025b565b6114616130e1565b5016600052600460205260406000206040519061147d8261075a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c6020820152604051809161046582604081019263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b359061ffff821682036101f957565b359060ff821682036101f957565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f957366023820112156101f95780600401359061156082610810565b9061156e60405192836107ae565b8282526024610260602084019402820101903682116101f957602401925b81841061159c57610995836130fa565b83360361026081126101f9576102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0604051926115d98461075a565b6115e288610289565b845201126101f957610260916020916115f9610800565b611604848901610843565b8152611612604089016114e8565b8482015261162260608901610828565b604082015261163360808901610828565b606082015261164460a08901610828565b608082015261165560c089016114f7565b60a082015261166660e089016114f7565b60c082015261167861010089016114e8565b60e082015261168a6101208901610828565b61010082015261169d61014089016114e8565b6101208201526116b061016089016114e8565b6101408201526116c36101808901610469565b6101608201526116d66101a089016114e8565b6101808201526116e96101c08901610828565b6101a08201526116fc6101e08901610828565b6101c082015261170f6102008901610289565b6101e08201526117226102208901610828565b6102008201526117356102408901610828565b6102208201528382015281520193019261158c565b6107fe909291926102208061024083019561176784825115159052565b60208181015161ffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015163ffffffff169085015260a08181015160ff169085015260c08181015160ff169085015260e08181015161ffff16908501526101008181015163ffffffff16908501526101208181015161ffff16908501526101408181015161ffff1690850152610160818101517fffffffff0000000000000000000000000000000000000000000000000000000016908501526101808181015161ffff16908501526101a08181015163ffffffff16908501526101c08181015163ffffffff16908501526101e08181015167ffffffffffffffff16908501526102008181015163ffffffff1690850152015163ffffffff16910152565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576104656119676119626118cd61025b565b60006102206118da610800565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e082015282610200820152015267ffffffffffffffff166000526007602052604060002090565b6133fc565b6040519182918261174a565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760005473ffffffffffffffffffffffffffffffffffffffff81163303611a32577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f830112156101f9578135611a7381610810565b92611a8160405194856107ae565b81845260208085019260051b8201019283116101f957602001905b828210611aa95750505090565b602080918335611ab8816101db565b815201910190611a9c565b346101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f957611b12903690600401611a5c565b60243567ffffffffffffffff81116101f957611b32903690600401611a5c565b90611b3b613c86565b60005b8151811015611bb75780611b5f611b5a61100060019486612a7c565b6156a3565b611b6a575b01611b3e565b73ffffffffffffffffffffffffffffffffffffffff611b8c6110008386612a7c565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2611b64565b8260005b81518110156109955780611bdc611bd761100060019486612a7c565b6156c4565b611be7575b01611bbb565b73ffffffffffffffffffffffffffffffffffffffff611c096110008386612a7c565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611be1565b346101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957610465611cdb611c7161025b565b67ffffffffffffffff60243591611c87836101db565b60006060604051611c9781610776565b828152826020820152826040820152015216600052600860205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b60ff60405191611cea83610776565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015260405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101f957604051611e2d8161075a565b816004013567ffffffffffffffff81116101f957611e519060043691850101611a5c565b8152602482013567ffffffffffffffff81116101f957610995926004611e7a9236920101611a5c565b602082015261357f565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760043567ffffffffffffffff81116101f957366023820112156101f957806004013590611edf82610810565b90611eed60405192836107ae565b8282526024602083019360061b820101903682116101f957602401925b818410611f1a576109958361371f565b6040843603126101f95760206040918251611f348161075a565b8635611f3f816101db565b8152611f4c838801610289565b83820152815201930192611f0a565b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95760405180602060095491828152019060096000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af9060005b818110611fda5761046585610c65818703826107ae565b8254845260209093019260019283019201611fc3565b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576040612030600435611368816101db565b61206a8251809263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565bf35b346101f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957602060405160128152f35b346101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f9576120dd61025b565b60243567ffffffffffffffff81116101f957806004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101f9576121416119628267ffffffffffffffff166000526007602052604060002090565b9161215361214f8451151590565b1590565b61251f57606481019261219761214f61216b86612937565b73ffffffffffffffffffffffffffffffffffffffff166000526001600901602052604060002054151590565b6124cd5760448201946121aa8682613008565b9490506121b8828483614664565b916121c56113ee88612937565b96846121e26121dc61020083015163ffffffff1690565b85614da2565b996000808a156124955750509161228a61227d61223d61ffff8561226f99612249999861223461222d6101a06122216101806122909d015161ffff1690565b95015163ffffffff1690565b918b613008565b94909316614ebd565b98919897909894612937565b73ffffffffffffffffffffffffffffffffffffffff166000526006602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b906127fa565b9560009761ffff6122a761014089015161ffff1690565b16612437575b509461228a61227d6101e061239d6104659d6dffffffffffffffffffffffffffff6123956123b59f9e9b6123907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9f9b9c6123ad9f6123909e63ffffffff61231b6123909f6024612325950190612981565b929050169061388c565b908b60a0810161234861234261233c835160ff1690565b60ff1690565b856127fa565b9360e083019161235a835161ffff1690565b9061ffff821683116123c5575b5050505060800151612390916123879163ffffffff166138d7565b6138d7565b63ffffffff1690565b61388c565b9116906127fa565b93015167ffffffffffffffff1690565b91169061280d565b6040519081529081906020820190565b612387949650612390959361ffff6124266124156123829661240f6124086123ff60809960ff6123f961242d9b5160ff1690565b16613899565b965161ffff1690565b61ffff1690565b906138b2565b61228a61233c60c08d015160ff1690565b911661388c565b9593839550612367565b9095949897508261245d8b989495986dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff169161247b6024890185612981565b90506124879388614fd9565b9697939438969392966122ad565b94939150945061228a61227d61226f6122496124c76124c261238761022061229099015163ffffffff1690565b61276d565b94612937565b61251b6124d985612937565b7f2502348c0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6000fd5b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b346101f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f95773ffffffffffffffffffffffffffffffffffffffff6004356125a7816101db565b6125af613c86565b1633811461262157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101f95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f957600435612686816101db565b61268e610272565b9067ffffffffffffffff82169182600052600760205260ff6040600020541615612710576126be6126df926138f1565b92600052600760205263ffffffff60016040600020015460901c1690614da2565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc10000149015171561278f57565b61273e565b908160051b918083046020149015171561278f57565b9061012c82029180830461012c149015171561278f57565b9061010c82029180830461010c149015171561278f57565b90655af3107a4000820291808304655af3107a4000149015171561278f57565b8181029291811591840414171561278f57565b8115612817570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b61288561287f6110b394937bffffffffffffffffffffffffffffffffffffffffffffffffffffffff61287881956138f1565b16906127fa565b926138f1565b169061280d565b9061289682610810565b6128a360405191826107ae565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06128d18294610810565b019060005b8281106128e257505050565b8060606020809385010152016128d6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156129325760061b0190565b6128f3565b356110b3816101db565b91908110156129325760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156101f9570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101f9570180359067ffffffffffffffff82116101f9576020019181360383136101f957565b92919267ffffffffffffffff82116107555760405191612a1a601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001846107ae565b8294818452818301116101f9578281602093846000960137010152565b90604051612a4481610776565b606060ff82945463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c166040850152821c161515910152565b80518210156129325760209160051b010190565b909291612add612ab48367ffffffffffffffff166000526007602052604060002090565b5460081b7fffffffff000000000000000000000000000000000000000000000000000000001690565b90612ae78161288c565b9560005b828110612afc575050505050505090565b612b0f612b0a828489612922565b612937565b8388612b29612b1f858484612941565b6040810190612981565b905060208111612c71575b508392612b6a612b64612b5d612b53600198612bb297612bad97612941565b6020810190612981565b36916129d2565b896139a0565b612b888967ffffffffffffffff166000526008602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b612a37565b606081015115612c3557612c19612bd36020612bed93015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107ae565b612c23828b612a7c565b52612c2e818a612a7c565b5001612aeb565b50612bed612c19612c6c84612c5e8a67ffffffffffffffff166000526007602052604060002090565b015460101c63ffffffff1690565b612bd3565b915050612ca9612387612c9c84612b888b67ffffffffffffffff166000526008602052604060002090565b5460401c63ffffffff1690565b10612cb657838838612b34565b7f36f536ca0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b90612d02613c86565b6000915b8051831015612f3a57612d198382612a7c565b5190612d2d825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b85518051821015612f2557612d5982602092612a7c565b510151612d85612d6a838951612a7c565b515173ffffffffffffffffffffffffffffffffffffffff1690565b604082015163ffffffff1660208110612ed7575090867f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed173ffffffffffffffffffffffffffffffffffffffff84612e84858f60019998612b88612dfc9267ffffffffffffffff166000526008602052604060002090565b815181546020808501516040808701516060978801517fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951663ffffffff96909616959095179190921b67ffffffff00000000161792901b6bffffffff0000000000000000169190911790151590921b6cff00000000000000000000000016919091179055565b612ece604051928392169582919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b0390a301612d42565b7f24ecdc020000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff90911660045263ffffffff1660245260446000fd5b50509550925092600191500191929092612d06565b50905060005b81518110156130045780612f68612f5960019385612a7c565b515167ffffffffffffffff1690565b67ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff612fb16020612f958689612a7c565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b6000612fd582612b888767ffffffffffffffff166000526008602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a301612f40565b5050565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101f9570180359067ffffffffffffffff82116101f957602001918160061b360383136101f957565b35907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821682036101f957565b6040813603126101f9576130b46020604051926130a18461075a565b80356130ac816101db565b84520161305c565b602082015290565b6040813603126101f9576130b46020604051926130d88461075a565b6130ac81610289565b604051906130ee8261075a565b60006020838281520152565b90613103613c86565b60005b82518110156133f7576131198184612a7c565b516020613129612f598487612a7c565b9101519067ffffffffffffffff8116801580156133d8575b80156133aa575b801561326a575b61323257916131f882600195946131a8613183612ab46131fd9767ffffffffffffffff166000526007602052604060002090565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b613203577faae4c0e87e0ed4dd3178367ccd926daebfe7fdb4d86913f3b555d9620622b961604051806131db878261174a565b0390a267ffffffffffffffff166000526007602052604060002090565b614052565b01613106565b7f425825c9dae839e14fe8f6b91cc8aefe0131bfff74c5958aec031bad9566c4f3604051806131db878261174a565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b507fffffffff000000000000000000000000000000000000000000000000000000006132ba6101608501517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c00000000000000000000000000000000000000000000000000000000811415908161337f575b81613354575b81613329575b816132fe575b5061314f565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501415386132f8565b7fc4e059530000000000000000000000000000000000000000000000000000000081141591506132f2565b7fac77ffec0000000000000000000000000000000000000000000000000000000081141591506132ec565b7f1e10bdc40000000000000000000000000000000000000000000000000000000081141591506132e6565b506101c083015163ffffffff1663ffffffff6133d0612387606087015163ffffffff1690565b911611613148565b5063ffffffff6133f06101c085015163ffffffff1690565b1615613141565b509050565b906107fe613571600161340d610800565b946135106134e682546134296134238260ff1690565b15158a52565b61ffff600882901c1660208a015263ffffffff601882901c1660408a015263ffffffff603882901c1660608a015263ffffffff605882901c1660808a015260ff607882901c1660a08a015260ff608082901c1660c08a015261ffff608882901c1660e08a015263ffffffff609882901c166101008a015261ffff60b882901c166101208a015261ffff60c882901c166101408a015260081b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000016610160880152565b015461ffff811661018086015263ffffffff601082901c166101a086015263ffffffff603082901c166101c086015267ffffffffffffffff605082901c166101e086015263ffffffff609082901c1661020086015260b01c63ffffffff1690565b63ffffffff16610220840152565b613587613c86565b60208101519160005b835181101561363b57806135a961100060019387612a7c565b6135e56135e073ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b615b08565b6135f1575b5001613590565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a1386135ea565b5091505160005b8151811015613004576136586110008284612a7c565b9073ffffffffffffffffffffffffffffffffffffffff8216156136f5577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6136ec836136c46136bf6135c760019773ffffffffffffffffffffffffffffffffffffffff1690565b615a8f565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101613642565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b613727613c86565b60005b8151811015613004578073ffffffffffffffffffffffffffffffffffffffff61375560019385612a7c565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d6137f267ffffffffffffffff60206137918689612a7c565b51015116836000526006602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a20161372a565b73ffffffffffffffffffffffffffffffffffffffff906138196130e1565b501660005260056020526040600020604051906138358261075a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c602082015290565b906002820180921161278f57565b906020820180921161278f57565b906001820180921161278f57565b9190820180921161278f57565b9061ffff8091169116029061ffff821691820361278f57565b9190820391821161278f57565b63ffffffff60209116019063ffffffff821161278f57565b9063ffffffff8091169116019063ffffffff821161278f57565b6138fa816137fb565b9063ffffffff602083015116158015613979575b6139355750517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8251161561390e565b7fffffffff000000000000000000000000000000000000000000000000000000001691907f2812d52c000000000000000000000000000000000000000000000000000000008314613ae7577f1e10bdc4000000000000000000000000000000000000000000000000000000008314613ad9577fac77ffec000000000000000000000000000000000000000000000000000000008314613ace577f647e2ba9000000000000000000000000000000000000000000000000000000008314613ac3577fc4e05953000000000000000000000000000000000000000000000000000000008314613ab557827f2ee820750000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6107fe919250600b90615180565b6107fe9192506151e1565b6107fe91925061511d565b6107fe919250600190615180565b6107fe91925061508d565b917fffffffff0000000000000000000000000000000000000000000000000000000083167f2812d52c000000000000000000000000000000000000000000000000000000008114613c7a577f1e10bdc4000000000000000000000000000000000000000000000000000000008114613c5a577fac77ffec000000000000000000000000000000000000000000000000000000008114613c4e577f647e2ba9000000000000000000000000000000000000000000000000000000008114613c42577fc4e059530000000000000000000000000000000000000000000000000000000014613c28577f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000831660045260246000fd5b6107fe925015613c3a57600b90615180565b600090615180565b50506107fe91506151e1565b50506107fe915061511d565b506107fe925015613c715760ff60015b1690615180565b60ff6000613c6a565b50506107fe915061508d565b73ffffffffffffffffffffffffffffffffffffffff600154163303613ca757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b33600052600360205260406000205415613ce757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b611962613d3c9196949395929667ffffffffffffffff166000526007602052604060002090565b946101608601947fffffffff00000000000000000000000000000000000000000000000000000000613d8e87517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c000000000000000000000000000000000000000000000000000000008114908115614028575b8115613ffe575b50613fb9577fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613e3288517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613f9b5750507f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613ea686517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613f275761251b613ed985517fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b613f689350612b5d6060613f52613f4b61238783613f94989a015163ffffffff1690565b8486615567565b0151604051958691602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018652856107ae565b9160019190565b945094506110b391613fae9136916129d2565b9360019336916129d2565b94509491613fdf91613fd96123876101c06110b396015163ffffffff1690565b9161530c565b93613ff66020613fee8761542d565b960151151590565b9336916129d2565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501438613dc1565b7fac77ffec0000000000000000000000000000000000000000000000000000000081149150613dba565b9061460b61022060016107fe9461409d61406c8651151590565b829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b6140e36140af602087015161ffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff1660089190911b62ffff0016178255565b61412f6140f7604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff1660189190911b66ffffffff00000016178255565b61417f614143606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffff1660389190911b6affffffff0000000000000016178255565b6141d3614193608087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff1660589190911b6effffffff000000000000000000000016178255565b6142256141e460a087015160ff1690565b82547fffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffff1660789190911b6fff00000000000000000000000000000016178255565b61427861423660c087015160ff1690565b82547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1660809190911b70ff0000000000000000000000000000000016178255565b6142ce61428a60e087015161ffff1690565b82547fffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffff1660889190911b72ffff000000000000000000000000000000000016178255565b61432b6142e361010087015163ffffffff1690565b82547fffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffff1660989190911b76ffffffff0000000000000000000000000000000000000016178255565b61438861433e61012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b6143e761439b61014087015161ffff1690565b82547fffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffff1660c89190911b7affff0000000000000000000000000000000000000000000000000016178255565b6144686144186101608701517fffffffff000000000000000000000000000000000000000000000000000000001690565b82547fff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff1660089190911c7effffffff00000000000000000000000000000000000000000000000000000016178255565b01926144ac61447d61018083015161ffff1690565b859061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6144f86144c16101a083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178555565b61454861450d6101c083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1660309190911b69ffffffff00000000000016178555565b6145a46145616101e083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff1660509190911b71ffffffffffffffff0000000000000000000016178555565b6146006145b961020083015163ffffffff1690565b85547fffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffff1660909190911b75ffffffff00000000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff0000000000000000000000000000000000000000000083549260b01b169116179055565b908160209103126101f9573590565b91906146736020830183612981565b939050604083016146848185613008565b9050604084019161469c612387845163ffffffff1690565b808811614d705750602085015161ffff16808311614d3a57506101608501966146e588517fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f2812d52c0000000000000000000000000000000000000000000000000000000081148015614d11575b8015614ce8575b156147c9575050505050505091816147c3612b5d6147926147bc9661476460806110b3980186612981565b61478c61238760606147806101c0879697015163ffffffff1690565b94015163ffffffff1690565b92615830565b51958694517fffffffff000000000000000000000000000000000000000000000000000000001690565b9280612981565b90613af2565b7fc4e0595300000000000000000000000000000000000000000000000000000000819b9a939495979996989b14600014614a6c57505061486c61483361485f999a606061482d61238761481f60808b018b612981565b939094015163ffffffff1690565b916156e5565b918251998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b6147c3612b5d8880612981565b60608101515190826148896148818780612981565b810190614655565b614a4f575081614a19575b8515159081614a0c575b506149e257604081116149b057506148c3906148bd85949397956127c2565b9061388c565b946000935b8385106149215750505050506123876148e5915163ffffffff1690565b8082116148f157505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614985612387612c9c6149508667ffffffffffffffff166000526008602052604060002090565b614961612b0a8d610d8c8b8d613008565b73ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b80156149a0576149949161388c565b965b01939291906148c8565b506149aa90613870565b96614996565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b604091500151153861489e565b61251b827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614a669193506148bd614a618461387e565b612794565b91614894565b7f1e10bdc40000000000000000000000000000000000000000000000000000000003614c995750614af4614ab961485f999a6060614ab361238761481f60808b018b612981565b91615567565b91614acb612387845163ffffffff1690565b998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b6080810151519082614b096148818780612981565b614c81575081614c4b575b85151580614c3f575b6149e25760408211614c0b576020015167ffffffffffffffff9081169081831c16614bd1575050614b55906148bd85949397956127aa565b946000935b838510614b775750505050506123876148e5915163ffffffff1690565b9091929395600190614ba6612387612c9c6149508667ffffffffffffffff166000526008602052604060002090565b8015614bc157614bb59161388c565b965b0193929190614b5a565b50614bcb90613870565b96614bb7565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614b1d565b61251b827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614c939193506148bd614a6184613862565b91614b14565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b507f647e2ba9000000000000000000000000000000000000000000000000000000008114614739565b507fac77ffec000000000000000000000000000000000000000000000000000000008114614732565b7fd88dddd600000000000000000000000000000000000000000000000000000000600052600483905261ffff1660245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff8116600052600460205260406000209160405192614dc88461075a565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116845260e01c9182602085015263ffffffff82169283614e2c575b505050506110b390517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b63ffffffff16429081039390841161278f578311614e4a5780614e02565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b6040813603126101f957602060405191614ea68361075a565b8035614eb1816101db565b83520135602082015290565b91939094926000926000966000966000945b808610614ee157505050505050929190565b909192939495614f3f612bad614f03614efe8a869d9e9d8b612922565b614e8d565b614961614f248767ffffffffffffffff166000526008602052604060002090565b915173ffffffffffffffffffffffffffffffffffffffff1690565b98614f5061214f60608c0151151590565b614fae5760208a015163ffffffff16614f68916138d7565b9960408a0151614f7b9063ffffffff1690565b614f84916138d7565b9851614f959063ffffffff1661276d565b614f9e9161388c565b956001905b019493929190614ecf565b969850614fd3614fcd85614fc76001946148bd8a61276d565b996138d7565b9a6138bf565b98614fa3565b91939093806101e00193846101e01161278f57610120810290808204610120149015171561278f576101e091010180931161278f5761240861014061506f6110b3966dffffffffffffffffffffffffffff61239561505a6150476150799a63ffffffff61228a9a169061388c565b61228a6124086101208c015161ffff1690565b6148bd6123876101008b015163ffffffff1690565b93015161ffff1690565b6127da565b908160209103126101f9575190565b60208151036150d0576150a9602082518301016020830161507e565b73ffffffffffffffffffffffffffffffffffffffff8111908115615111575b506150d05750565b61510d906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906102cf565b0390fd5b610400915010386150c8565b602081510361514357600b61513b602083518401016020840161507e565b106151435750565b61510d906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840181815201906102cf565b9060208251036151a65780615193575050565b61513b602083518401016020840161507e565b6040517fe0d7fb02000000000000000000000000000000000000000000000000000000008152602060048201528061510d60248201856102cf565b60248151036151f7576022810151156151f75750565b61510d906040519182917f373b0e4400000000000000000000000000000000000000000000000000000000835260206004840181815201906102cf565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110615268575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90929192836004116101f95783116101f957600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b906004116101f95790600490565b908160409103126101f9576020604051916152fd8361075a565b8051835201516130b481610839565b916153156130e1565b50811561540b5750615356612b5d82806153507fffffffff000000000000000000000000000000000000000000000000000000009587615234565b9561529a565b91167f181dcf100000000000000000000000000000000000000000000000000000000081036153935750806020806110b3935183010191016152e3565b7f97a657c900000000000000000000000000000000000000000000000000000000146153e3577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b806020806153f69351830101910161507e565b6153fe6107ef565b9081526000602082015290565b91505067ffffffffffffffff61541f6107ef565b911681526000602082015290565b6020604051917f181dcf10000000000000000000000000000000000000000000000000000000008284015280516024840152015115156044820152604481526110b36064826107ae565b9080601f830112156101f957813561548e81610810565b9261549c60405194856107ae565b81845260208085019260051b8201019283116101f957602001905b8282106154c45750505090565b81358152602091820191016154b7565b6020818303126101f95780359067ffffffffffffffff82116101f9570160a0818303126101f9576040519161550883610792565b61551182610828565b835261551f60208301610289565b6020840152604082013561553281610839565b604084015260608201356060840152608082013567ffffffffffffffff81116101f95761555f9201615477565b608082015290565b6060608060405161557781610792565b60008152600060208201526000604082015260008382015201528115615679577f1f3b3aba000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006155ec6155e685856152d5565b90615234565b160361564f5781615608926156009261529a565b8101906154d4565b9063ffffffff61561c835163ffffffff1690565b16116156255790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff6110b39116600961593c565b73ffffffffffffffffffffffffffffffffffffffff6110b391166009615aca565b906060806040516156f581610776565b60008152600060208201526000604082015201528015615679577f21ea4ca9000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061575e6155e684866152d5565b160361564f578061576e9261529a565b8192910160009260208183031261582c5780359067ffffffffffffffff821161582857019260808483031261582557604051936157aa85610776565b8035855260208101356157bc81610839565b60208601526040810135604086015260608101359167ffffffffffffffff831161582557506157ec929101615477565b60608301528151116157fb5790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b80fd5b8480fd5b8380fd5b9063ffffffff61584a936158426130e1565b50169161530c565b908151116157fb5790565b80548210156129325760005260206000200190600090565b916158a5918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8054801561590d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906158de8282615855565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615a2a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840184811161278f578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff850194851161278f5760009585836159db976159cc95036159e1575b5050506158a9565b90600052602052604060002090565b55600190565b615a11615a0b91615a026159f8615a219588615855565b90549060031b1c90565b92839187615855565b9061586d565b8590600052602052604060002090565b553880806159c4565b50505050600090565b805490680100000000000000008210156107555781615a5a9160016158a594018155615855565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b600081815260036020526040902054615ac457615aad816002615a33565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615b015780615aec83600193615a33565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615b01577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161278f57600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840193841161278f5783836159db9460009603615ba4575b505050615b9360026158a9565b600390600052602052604060002090565b615b93615a0b91615bbc6159f8615bc6956002615855565b9283916002615855565b55388080615b8656fea164736f6c634300081a000a",
}

var FeeQuoterABI = FeeQuoterMetaData.ABI

var FeeQuoterBin = FeeQuoterMetaData.Bin

func DeployFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoter, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterBin), backend, staticConfig, priceUpdaters, feeTokens, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeQuoter{address: address, abi: *parsed, FeeQuoterCaller: FeeQuoterCaller{contract: contract}, FeeQuoterTransactor: FeeQuoterTransactor{contract: contract}, FeeQuoterFilterer: FeeQuoterFilterer{contract: contract}}, nil
}

type FeeQuoter struct {
	address common.Address
	abi     abi.ABI
	FeeQuoterCaller
	FeeQuoterTransactor
	FeeQuoterFilterer
}

type FeeQuoterCaller struct {
	contract *bind.BoundContract
}

type FeeQuoterTransactor struct {
	contract *bind.BoundContract
}

type FeeQuoterFilterer struct {
	contract *bind.BoundContract
}

type FeeQuoterSession struct {
	Contract     *FeeQuoter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeeQuoterCallerSession struct {
	Contract *FeeQuoterCaller
	CallOpts bind.CallOpts
}

type FeeQuoterTransactorSession struct {
	Contract     *FeeQuoterTransactor
	TransactOpts bind.TransactOpts
}

type FeeQuoterRaw struct {
	Contract *FeeQuoter
}

type FeeQuoterCallerRaw struct {
	Contract *FeeQuoterCaller
}

type FeeQuoterTransactorRaw struct {
	Contract *FeeQuoterTransactor
}

func NewFeeQuoter(address common.Address, backend bind.ContractBackend) (*FeeQuoter, error) {
	abi, err := abi.JSON(strings.NewReader(FeeQuoterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeeQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeQuoter{address: address, abi: abi, FeeQuoterCaller: FeeQuoterCaller{contract: contract}, FeeQuoterTransactor: FeeQuoterTransactor{contract: contract}, FeeQuoterFilterer: FeeQuoterFilterer{contract: contract}}, nil
}

func NewFeeQuoterCaller(address common.Address, caller bind.ContractCaller) (*FeeQuoterCaller, error) {
	contract, err := bindFeeQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterCaller{contract: contract}, nil
}

func NewFeeQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeQuoterTransactor, error) {
	contract, err := bindFeeQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTransactor{contract: contract}, nil
}

func NewFeeQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeQuoterFilterer, error) {
	contract, err := bindFeeQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFilterer{contract: contract}, nil
}

func bindFeeQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FeeQuoter *FeeQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoter.Contract.FeeQuoterCaller.contract.Call(opts, result, method, params...)
}

func (_FeeQuoter *FeeQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.Contract.FeeQuoterTransactor.contract.Transfer(opts)
}

func (_FeeQuoter *FeeQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoter.Contract.FeeQuoterTransactor.contract.Transact(opts, method, params...)
}

func (_FeeQuoter *FeeQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoter.Contract.contract.Call(opts, result, method, params...)
}

func (_FeeQuoter *FeeQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.Contract.contract.Transfer(opts)
}

func (_FeeQuoter *FeeQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoter.Contract.contract.Transact(opts, method, params...)
}

func (_FeeQuoter *FeeQuoterCaller) FEEBASEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "FEE_BASE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.FEEBASEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.FEEBASEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) KEYSTONEPRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "KEYSTONE_PRICE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "convertTokenAmount", fromToken, fromTokenAmount, toToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.ConvertTokenAmount(&_FeeQuoter.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoter *FeeQuoterCallerSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.ConvertTokenAmount(&_FeeQuoter.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoter *FeeQuoterCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetAllAuthorizedCallers(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetAllAuthorizedCallers(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(FeeQuoterDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterDestChainConfig)).(*FeeQuoterDestChainConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetDestChainConfig(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetDestChainConfig(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getDestinationChainGasPrice", destChainSelector)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetDestinationChainGasPrice(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetDestinationChainGasPrice(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getFeeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetFeeTokens(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetFeeTokens(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getPremiumMultiplierWeiPerEth", token)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoter.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoter.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FeeQuoterStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterStaticConfig)).(*FeeQuoterStaticConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoter.Contract.GetStaticConfig(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoter.Contract.GetStaticConfig(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenAndGasPrices", token, destChainSelector)

	outstruct := new(GetTokenAndGasPrices)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPriceValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoter.Contract.GetTokenAndGasPrices(&_FeeQuoter.CallOpts, token, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoter.Contract.GetTokenAndGasPrices(&_FeeQuoter.CallOpts, token, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPrice", token)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPrices", tokens)

	if err != nil {
		return *new([]InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new([]InternalTimestampedPackedUint224)).(*[]InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrices(&_FeeQuoter.CallOpts, tokens)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrices(&_FeeQuoter.CallOpts, tokens)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenTransferFeeConfig", destChainSelector, token)

	if err != nil {
		return *new(FeeQuoterTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenTransferFeeConfig)).(*FeeQuoterTokenTransferFeeConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoter.Contract.GetTokenTransferFeeConfig(&_FeeQuoter.CallOpts, destChainSelector, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoter.Contract.GetTokenTransferFeeConfig(&_FeeQuoter.CallOpts, destChainSelector, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getValidatedFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedFee(&_FeeQuoter.CallOpts, destChainSelector, message)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedFee(&_FeeQuoter.CallOpts, destChainSelector, message)
}

func (_FeeQuoter *FeeQuoterCaller) GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getValidatedTokenPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) Owner() (common.Address, error) {
	return _FeeQuoter.Contract.Owner(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) Owner() (common.Address, error) {
	return _FeeQuoter.Contract.Owner(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "processMessageArgs", destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)

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

func (_FeeQuoter *FeeQuoterSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	return _FeeQuoter.Contract.ProcessMessageArgs(&_FeeQuoter.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)
}

func (_FeeQuoter *FeeQuoterCallerSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

	error) {
	return _FeeQuoter.Contract.ProcessMessageArgs(&_FeeQuoter.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, messageReceiver)
}

func (_FeeQuoter *FeeQuoterCaller) ProcessPoolReturnData(opts *bind.CallOpts, destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "processPoolReturnData", destChainSelector, onRampTokenTransfers, sourceTokenAmounts)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) ProcessPoolReturnData(destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	return _FeeQuoter.Contract.ProcessPoolReturnData(&_FeeQuoter.CallOpts, destChainSelector, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoter *FeeQuoterCallerSession) ProcessPoolReturnData(destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error) {
	return _FeeQuoter.Contract.ProcessPoolReturnData(&_FeeQuoter.CallOpts, destChainSelector, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoter *FeeQuoterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeQuoter.Contract.SupportsInterface(&_FeeQuoter.CallOpts, interfaceId)
}

func (_FeeQuoter *FeeQuoterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeQuoter.Contract.SupportsInterface(&_FeeQuoter.CallOpts, interfaceId)
}

func (_FeeQuoter *FeeQuoterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) TypeAndVersion() (string, error) {
	return _FeeQuoter.Contract.TypeAndVersion(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) TypeAndVersion() (string, error) {
	return _FeeQuoter.Contract.TypeAndVersion(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "acceptOwnership")
}

func (_FeeQuoter *FeeQuoterSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoter.Contract.AcceptOwnership(&_FeeQuoter.TransactOpts)
}

func (_FeeQuoter *FeeQuoterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoter.Contract.AcceptOwnership(&_FeeQuoter.TransactOpts)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoter.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoter.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyDestChainConfigUpdates(&_FeeQuoter.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyDestChainConfigUpdates(&_FeeQuoter.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyFeeTokensUpdates", feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoter *FeeQuoterSession) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyPremiumMultiplierWeiPerEthUpdates", premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoter.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoter.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "transferOwnership", to)
}

func (_FeeQuoter *FeeQuoterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.TransferOwnership(&_FeeQuoter.TransactOpts, to)
}

func (_FeeQuoter *FeeQuoterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.TransferOwnership(&_FeeQuoter.TransactOpts, to)
}

func (_FeeQuoter *FeeQuoterTransactor) UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "updatePrices", priceUpdates)
}

func (_FeeQuoter *FeeQuoterSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdatePrices(&_FeeQuoter.TransactOpts, priceUpdates)
}

func (_FeeQuoter *FeeQuoterTransactorSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdatePrices(&_FeeQuoter.TransactOpts, priceUpdates)
}

type FeeQuoterAuthorizedCallerAddedIterator struct {
	Event *FeeQuoterAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterAuthorizedCallerAdded)
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
		it.Event = new(FeeQuoterAuthorizedCallerAdded)
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

func (it *FeeQuoterAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterAuthorizedCallerAddedIterator{contract: _FeeQuoter.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterAuthorizedCallerAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterAuthorizedCallerAdded, error) {
	event := new(FeeQuoterAuthorizedCallerAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterAuthorizedCallerRemovedIterator struct {
	Event *FeeQuoterAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterAuthorizedCallerRemoved)
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
		it.Event = new(FeeQuoterAuthorizedCallerRemoved)
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

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterAuthorizedCallerRemovedIterator{contract: _FeeQuoter.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterAuthorizedCallerRemoved)
				if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterAuthorizedCallerRemoved, error) {
	event := new(FeeQuoterAuthorizedCallerRemoved)
	if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterDestChainAddedIterator struct {
	Event *FeeQuoterDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterDestChainAdded)
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
		it.Event = new(FeeQuoterDestChainAdded)
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

func (it *FeeQuoterDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterDestChainAdded struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterDestChainAddedIterator{contract: _FeeQuoter.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterDestChainAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseDestChainAdded(log types.Log) (*FeeQuoterDestChainAdded, error) {
	event := new(FeeQuoterDestChainAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterDestChainConfigUpdatedIterator struct {
	Event *FeeQuoterDestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterDestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterDestChainConfigUpdated)
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
		it.Event = new(FeeQuoterDestChainConfigUpdated)
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

func (it *FeeQuoterDestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterDestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterDestChainConfigUpdated struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterDestChainConfigUpdatedIterator{contract: _FeeQuoter.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterDestChainConfigUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterDestChainConfigUpdated, error) {
	event := new(FeeQuoterDestChainConfigUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterFeeTokenAddedIterator struct {
	Event *FeeQuoterFeeTokenAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterFeeTokenAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterFeeTokenAdded)
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
		it.Event = new(FeeQuoterFeeTokenAdded)
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

func (it *FeeQuoterFeeTokenAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterFeeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterFeeTokenAdded struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFeeTokenAddedIterator{contract: _FeeQuoter.contract, event: "FeeTokenAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAdded, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterFeeTokenAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseFeeTokenAdded(log types.Log) (*FeeQuoterFeeTokenAdded, error) {
	event := new(FeeQuoterFeeTokenAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterFeeTokenRemovedIterator struct {
	Event *FeeQuoterFeeTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterFeeTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterFeeTokenRemoved)
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
		it.Event = new(FeeQuoterFeeTokenRemoved)
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

func (it *FeeQuoterFeeTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterFeeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterFeeTokenRemoved struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenRemovedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFeeTokenRemovedIterator{contract: _FeeQuoter.contract, event: "FeeTokenRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterFeeTokenRemoved)
				if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseFeeTokenRemoved(log types.Log) (*FeeQuoterFeeTokenRemoved, error) {
	event := new(FeeQuoterFeeTokenRemoved)
	if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterOwnershipTransferRequestedIterator struct {
	Event *FeeQuoterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterOwnershipTransferRequested)
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
		it.Event = new(FeeQuoterOwnershipTransferRequested)
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

func (it *FeeQuoterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterOwnershipTransferRequestedIterator{contract: _FeeQuoter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterOwnershipTransferRequested)
				if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterOwnershipTransferRequested, error) {
	event := new(FeeQuoterOwnershipTransferRequested)
	if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterOwnershipTransferredIterator struct {
	Event *FeeQuoterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterOwnershipTransferred)
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
		it.Event = new(FeeQuoterOwnershipTransferred)
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

func (it *FeeQuoterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterOwnershipTransferredIterator{contract: _FeeQuoter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterOwnershipTransferred)
				if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseOwnershipTransferred(log types.Log) (*FeeQuoterOwnershipTransferred, error) {
	event := new(FeeQuoterOwnershipTransferred)
	if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator struct {
	Event *FeeQuoterPremiumMultiplierWeiPerEthUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
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
		it.Event = new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
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

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterPremiumMultiplierWeiPerEthUpdated struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
	Raw                        types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator{contract: _FeeQuoter.contract, event: "PremiumMultiplierWeiPerEthUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterPremiumMultiplierWeiPerEthUpdated, error) {
	event := new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterTokenTransferFeeConfigDeletedIterator struct {
	Event *FeeQuoterTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterTokenTransferFeeConfigDeleted)
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
		it.Event = new(FeeQuoterTokenTransferFeeConfigDeleted)
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

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Token             common.Address
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTokenTransferFeeConfigDeletedIterator{contract: _FeeQuoter.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterTokenTransferFeeConfigDeleted)
				if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterTokenTransferFeeConfigDeleted, error) {
	event := new(FeeQuoterTokenTransferFeeConfigDeleted)
	if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterTokenTransferFeeConfigUpdatedIterator struct {
	Event *FeeQuoterTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterTokenTransferFeeConfigUpdated)
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
		it.Event = new(FeeQuoterTokenTransferFeeConfigUpdated)
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

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	Token                  common.Address
	TokenTransferFeeConfig FeeQuoterTokenTransferFeeConfig
	Raw                    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTokenTransferFeeConfigUpdatedIterator{contract: _FeeQuoter.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterTokenTransferFeeConfigUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterTokenTransferFeeConfigUpdated, error) {
	event := new(FeeQuoterTokenTransferFeeConfigUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterUsdPerTokenUpdatedIterator struct {
	Event *FeeQuoterUsdPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterUsdPerTokenUpdated)
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
		it.Event = new(FeeQuoterUsdPerTokenUpdated)
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

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterUsdPerTokenUpdated struct {
	Token     common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterUsdPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterUsdPerTokenUpdatedIterator{contract: _FeeQuoter.contract, event: "UsdPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterUsdPerTokenUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterUsdPerTokenUpdated, error) {
	event := new(FeeQuoterUsdPerTokenUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterUsdPerUnitGasUpdatedIterator struct {
	Event *FeeQuoterUsdPerUnitGasUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterUsdPerUnitGasUpdated)
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
		it.Event = new(FeeQuoterUsdPerUnitGasUpdated)
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

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterUsdPerUnitGasUpdated struct {
	DestChain uint64
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterUsdPerUnitGasUpdatedIterator, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterUsdPerUnitGasUpdatedIterator{contract: _FeeQuoter.contract, event: "UsdPerUnitGasUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterUsdPerUnitGasUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterUsdPerUnitGasUpdated, error) {
	event := new(FeeQuoterUsdPerUnitGasUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (FeeQuoterAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (FeeQuoterAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (FeeQuoterDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0xaae4c0e87e0ed4dd3178367ccd926daebfe7fdb4d86913f3b555d9620622b961")
}

func (FeeQuoterDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x425825c9dae839e14fe8f6b91cc8aefe0131bfff74c5958aec031bad9566c4f3")
}

func (FeeQuoterFeeTokenAdded) Topic() common.Hash {
	return common.HexToHash("0xdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23")
}

func (FeeQuoterFeeTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91")
}

func (FeeQuoterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeeQuoterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FeeQuoterPremiumMultiplierWeiPerEthUpdated) Topic() common.Hash {
	return common.HexToHash("0xbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d")
}

func (FeeQuoterTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (FeeQuoterTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed1")
}

func (FeeQuoterUsdPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a")
}

func (FeeQuoterUsdPerUnitGasUpdated) Topic() common.Hash {
	return common.HexToHash("0xdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e")
}

func (_FeeQuoter *FeeQuoter) Address() common.Address {
	return _FeeQuoter.address
}

type FeeQuoterInterface interface {
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

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterAuthorizedCallerRemoved, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*FeeQuoterDestChainAdded, error)

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterDestChainConfigUpdated, error)

	FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedIterator, error)

	WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAdded, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenAdded(log types.Log) (*FeeQuoterFeeTokenAdded, error)

	FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenRemovedIterator, error)

	WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenRemoved(log types.Log) (*FeeQuoterFeeTokenRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeQuoterOwnershipTransferred, error)

	FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator, error)

	WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error)

	ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterPremiumMultiplierWeiPerEthUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterTokenTransferFeeConfigUpdated, error)

	FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterUsdPerTokenUpdatedIterator, error)

	WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterUsdPerTokenUpdated, error)

	FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterUsdPerUnitGasUpdatedIterator, error)

	WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error)

	ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterUsdPerUnitGasUpdated, error)

	Address() common.Address
}
