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
	EnforceOutOfOrder                 bool
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

type FeeQuoterTokenPriceFeedConfig struct {
	DataFeedAddress common.Address
	TokenDecimals   uint8
	IsEnabled       bool
}

type FeeQuoterTokenPriceFeedUpdate struct {
	SourceToken common.Address
	FeedConfig  FeeQuoterTokenPriceFeedConfig
}

type FeeQuoterTokenTransferFeeConfig struct {
	MinFeeUSDCents    uint32
	MaxFeeUSDCents    uint32
	DeciBps           uint16
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

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	DestTokenAddress   []byte
	Amount             *big.Int
	ExtraData          []byte
	Receipt            InternalReceipt
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint64
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

type InternalTimestampedPackedUint224 struct {
	Value     *big.Int
	Timestamp uint32
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

type KeystoneFeedsPermissionHandlerPermission struct {
	Forwarder     common.Address
	WorkflowName  [10]byte
	ReportName    [2]byte
	WorkflowOwner common.Address
	IsAllowed     bool
}

var FeeQuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_BASE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPriceFeedConfig\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnDataNew\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVMTokenTransfer\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveTokenReceiver\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setReportPermissions\",\"inputs\":[{\"name\":\"permissions\",\"type\":\"tuple[]\",\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenPriceFeeds\",\"inputs\":[{\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"priceFeedConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportPermissionSet\",\"inputs\":[{\"name\":\"reportId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"permission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataFeedValueOutOfUint224Range\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeRange\",\"inputs\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportForwarderUnauthorized\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346110b257617685803803806100198161131e565b92833981019080820361012081126110b2576060136110b25761003a6112e0565b81516001600160601b03811681036110b257815261005a60208301611343565b906020810191825261006e60408401611357565b6040820190815260608401516001600160401b0381116110b2578561009491860161137f565b60808501519094906001600160401b0381116110b257866100b691830161137f565b60a08201519096906001600160401b0381116110b25782019080601f830112156110b25781516100ed6100e882611368565b61131e565b9260208085848152019260071b820101908382116110b257602001915b81831061126b5750505060c08301516001600160401b0381116110b25783019781601f8a0112156110b2578851986101446100e88b611368565b996020808c838152019160051b830101918483116110b25760208101915b838310611109575050505060e08401516001600160401b0381116110b25784019382601f860112156110b257845161019c6100e882611368565b9560208088848152019260061b820101908582116110b257602001915b8183106110cd57505050610100810151906001600160401b0382116110b2570182601f820112156110b2578051906101f36100e883611368565b9360206102808187868152019402830101918183116110b257602001925b828410610ef057505050503315610edf57600180546001600160a01b031916331790556020986102408a61131e565b976000895260003681376102526112ff565b998a52888b8b015260005b89518110156102c4576001906001600160a01b0361027b828d611418565b51168d61028782611604565b610294575b50500161025d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388d61028c565b508a985089519660005b885181101561033f576001600160a01b036102e9828b611418565b511690811561032e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8c8361032060019561158c565b50604051908152a1016102ce565b6342bcdf7f60e11b60005260046000fd5b5081518a985089906001600160a01b0316158015610ecd575b8015610ebe575b610ead5791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c0526103928661131e565b9360008552600036813760005b855181101561040e576001906103c76001600160a01b036103c0838a611418565b5116611499565b6103d2575b0161039f565b818060a01b036103e28289611418565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26103cc565b508694508560005b84518110156104855760019061043e6001600160a01b036104378389611418565b51166115cb565b610449575b01610416565b818060a01b036104598288611418565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610443565b508593508460005b835181101561054757806104a360019286611418565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf606089858060a01b038451169301518360005260078b5260406000209060ff878060a01b038251169283898060a01b03198254161781558d8301908151604082549501948460a81b8651151560a81b16918560a01b9060a01b169061ffff60a01b19161717905560405193845251168c8301525115156040820152a20161048d565b5091509160005b8251811015610b45576105618184611418565b51856001600160401b036105758487611418565b5151169101519080158015610b32575b8015610b14575b8015610a92575b610a7e57600081815260098852604090205460019392919060081b6001600160e01b03191661093657807f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806106fc868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a25b60005260098752826040600020825115158382549162ffff008c83015160081b169066ffffffff000000604084015160181b166affffffff00000000000000606085015160381b16926effffffff0000000000000000000000608086015160581b169260ff60781b60a087015160781b169460ff60801b60c088015160801b169161ffff60881b60e089015160881b169063ffffffff60981b6101008a015160981b169361ffff60b81b6101208b015160b81b169661ffff60c81b6101408c015160c81b169963ffffffff60d81b6101608d015160081c169b61018060ff60f81b910151151560f81b169c8f8060f81b039a63ffffffff60d81b199961ffff60c81b199861ffff60b81b199763ffffffff60981b199661ffff60881b199560ff60801b199460ff60781b19936effffffff0000000000000000000000199260ff6affffffff000000000000001992169066ffffffffffffff19161716171617161716171617161716171617161716179063ffffffff60d81b1617178155019061ffff6101a0820151169082549165ffffffff00006101c083015160101b169269ffffffff0000000000006101e084015160301b166a01000000000000000000008860901b0361020085015160501b169263ffffffff60901b61022086015160901b169461024063ffffffff60b01b91015160b01b169563ffffffff60b01b199363ffffffff60901b19926a01000000000000000000008c60901b0319918c8060501b03191617161716171617171790550161054e565b807f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e360405180610a76868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a2610700565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61016083015116630a04b54b60e21b8114159081610b02575b81610af0575b81610ade575b81610acc575b50610593565b63647e2ba960e01b1415905088610ac6565b63c4e0595360e01b8114159150610ac0565b632b1dfffb60e21b8114159150610aba565b6307842f7160e21b8114159150610ab4565b5063ffffffff6101e08301511663ffffffff6060840151161061058c565b5063ffffffff6101e08301511615610585565b84828560005b8151811015610bcb576001906001600160a01b03610b698285611418565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610b9f8689611418565b510151168360005260088252604060002081878060401b0319825416179055604051908152a201610b4b565b83600184610bd88361131e565b9060008252600092610ea8575b909282935b8251851015610de757610bfd8584611418565b5180516001600160401b0316939083019190855b83518051821015610dd657610c27828792611418565b51015184516001600160a01b0390610c40908490611418565b5151169063ffffffff815116908781019163ffffffff8351169081811015610dc15750506080810163ffffffff815116898110610daa575090899291838c52600a8a5260408c20600160a01b6001900386168d528a5260408c2092825163ffffffff169380549180518d1b67ffffffff0000000016916040860192835160401b69ffff000000000000000016966060810195865160501b6dffffffff00000000000000000000169063ffffffff60701b895160701b169260a001998b60ff60901b8c51151560901b169560ff60901b199363ffffffff60701b19926dffffffff000000000000000000001991600160501b60019003191617161716171617171790556040519586525163ffffffff168c8601525161ffff1660408501525163ffffffff1660608401525163ffffffff16608083015251151560a082015260c07f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b591a3600101610c11565b6312766e0160e11b8c52600485905260245260448bfd5b6305a7b3d160e11b8c5260045260245260448afd5b505060019096019593509050610bea565b9150825b8251811015610e69576001906001600160401b03610e098286611418565b515116828060a01b0384610e1d8488611418565b5101511690808752600a855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610deb565b604051615fec908161169982396080518181816105c20152610c1c015260a0518181816105f80152610bcd015260c05181818161061f0152613aaf0152f35b610be5565b63d794ef9560e01b60005260046000fd5b5063ffffffff8251161561035f565b5080516001600160601b031615610358565b639b15e16f60e01b60005260046000fd5b83820361028081126110b257610260610f076112ff565b91610f11876113f5565b8352601f1901126110b2576040519161026083016001600160401b038111848210176110b757604052610f46602087016113e8565b8352610f5460408701611409565b6020840152610f6560608701611357565b6040840152610f7660808701611357565b6060840152610f8760a08701611357565b6080840152610f9860c087016113da565b60a0840152610fa960e087016113da565b60c0840152610fbb6101008701611409565b60e0840152610fcd6101208701611357565b610100840152610fe06101408701611409565b610120840152610ff36101608701611409565b610140840152610180860151916001600160e01b0319831683036110b2578360209361016061028096015261102b6101a089016113e8565b61018082015261103e6101c08901611409565b6101a08201526110516101e08901611357565b6101c08201526110646102008901611357565b6101e082015261107761022089016113f5565b61020082015261108a6102408901611357565b61022082015261109d6102608901611357565b61024082015283820152815201930192610211565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040838703126110b25760206040916110e46112ff565b6110ed86611343565b81526110fa8387016113f5565b838201528152019201916101b9565b82516001600160401b0381116110b25782016040818803601f1901126110b2576111316112ff565b9061113e602082016113f5565b825260408101516001600160401b0381116110b257602091010187601f820112156110b25780516111716100e882611368565b91602060e08185858152019302820101908a82116110b257602001915b8183106111ad5750505091816020938480940152815201920191610162565b828b0360e081126110b25760c06111c26112ff565b916111cc86611343565b8352601f1901126110b2576040519160c08301916001600160401b038311848410176110b75760e093602093604052611206848801611357565b815261121460408801611357565b8482015261122460608801611409565b604082015261123560808801611357565b606082015261124660a08801611357565b608082015261125760c088016113e8565b60a08201528382015281520192019161118e565b828403608081126110b25760606112806112ff565b9161128a86611343565b8352601f1901126110b2576080916020916112a36112e0565b6112ae848801611343565b81526112bc604088016113da565b848201526112cc606088016113e8565b60408201528382015281520192019161010a565b60405190606082016001600160401b038111838210176110b757604052565b60408051919082016001600160401b038111838210176110b757604052565b6040519190601f01601f191682016001600160401b038111838210176110b757604052565b51906001600160a01b03821682036110b257565b519063ffffffff821682036110b257565b6001600160401b0381116110b75760051b60200190565b9080601f830112156110b25781516113996100e882611368565b9260208085848152019260051b8201019283116110b257602001905b8282106113c25750505090565b602080916113cf84611343565b8152019101906113b5565b519060ff821682036110b257565b519081151582036110b257565b51906001600160401b03821682036110b257565b519061ffff821682036110b257565b805182101561142c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561142c5760005260206000200190600090565b805480156114835760001901906114718282611442565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600c6020526040902054801561155a57600019810181811161154457600b54600019810191908211611544578181036114f3575b5050506114df600b61145a565b600052600c60205260006040812055600190565b61152c61150461151593600b611442565b90549060031b1c928392600b611442565b819391549060031b91821b91600019901b19161790565b9055600052600c6020526040600020553880806114d2565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156110b7578161151591600161158894018155611442565b9055565b806000526003602052604060002054156000146115c5576115ae816002611561565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146115c5576115ed81600b611561565b600b5490600052600c602052604060002055600190565b600081815260036020526040902054801561155a576000198101818111611544576002546000198101919082116115445780820361165e575b50505061164a600261145a565b600052600360205260006040812055600190565b61168061166f611515936002611442565b90549060031b1c9283926002611442565b9055600052600360205260406000205538808061163d56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be1461023657806301447eaa1461023157806301ffc9a71461022c578063061877e31461022757806306285c6914610222578063181f5a771461021d5780632451a62714610218578063325c868e146102135780633937306f1461020e5780633a49bb491461020957806341ed29e71461020457806345ac924d146101ff5780634ab35b0b146101fa578063514e8cff146101f5578063614a8adc146101f05780636def4ce7146101eb578063770e2dc4146101e657806379ba5097146101e15780637afac322146101dc578063805f2132146101d757806382b49eb0146101d257806387b8d879146101cd5780638da5cb5b146101c857806391a2749a146101c35780639b1115e4146101be578063a69c64c0146101b9578063bf78e03f146101b4578063cdc73d51146101af578063d02641a0146101aa578063d63d3af2146101a5578063d8694ccd146101a0578063f2fde38b1461019b578063fbe3f778146101965763ffdb4b371461019157600080fd5b6127b4565b6126b7565b6125fb565b6121f9565b6121dd565b612194565b61211d565b612077565b611fbe565b611f6d565b611ed9565b611eb2565b611c96565b611b19565b61187e565b611745565b61162d565b61142c565b6112ad565b61103b565b610f5e565b610f26565b610e5d565b610cd1565b610b5b565b610881565b610865565b6107e2565b610740565b610586565b61053e565b610484565b6103d0565b61025e565b6001600160a01b0381160361024c57565b600080fd5b359061025c8261023b565b565b3461024c57606060031936011261024c5760206102956004356102808161023b565b602435604435916102908361023b565b61294a565b604051908152f35b6004359067ffffffffffffffff8216820361024c57565b6024359067ffffffffffffffff8216820361024c57565b359067ffffffffffffffff8216820361024c57565b9181601f8401121561024c5782359167ffffffffffffffff831161024c576020808501948460051b01011161024c57565b919082519283825260005b84811061033d575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161031c565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061038557505050505090565b90919293946020806103c1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610311565b97019301930191939290610376565b3461024c57606060031936011261024c576103e961029d565b60243567ffffffffffffffff811161024c576104099036906004016102e0565b6044929192359167ffffffffffffffff831161024c573660238401121561024c5782600401359167ffffffffffffffff831161024c573660248460061b8601011161024c5761046b94602461045f950192612b21565b60405191829182610352565b0390f35b35906001600160e01b03198216820361024c57565b3461024c57602060031936011261024c576004356001600160e01b0319811680910361024c57807f805f21320000000000000000000000000000000000000000000000000000000060209214908115610514575b81156104ea575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386104df565b7f66792e8000000000000000000000000000000000000000000000000000000000811491506104d8565b3461024c57602060031936011261024c576001600160a01b036004356105638161023b565b166000526008602052602067ffffffffffffffff60406000205416604051908152f35b3461024c57600060031936011261024c5761059f612cf3565b5060606040516105ae8161066d565b63ffffffff6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918281526001600160a01b0360406020830192827f00000000000000000000000000000000000000000000000000000000000000001684520191837f00000000000000000000000000000000000000000000000000000000000000001683526040519485525116602084015251166040820152f35b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761068957604052565b610657565b60a0810190811067ffffffffffffffff82111761068957604052565b6040810190811067ffffffffffffffff82111761068957604052565b60c0810190811067ffffffffffffffff82111761068957604052565b6080810190811067ffffffffffffffff82111761068957604052565b90601f601f19910116810190811067ffffffffffffffff82111761068957604052565b6040519061025c6040836106fe565b6040519061025c610260836106fe565b3461024c57600060031936011261024c5761046b604080519061076381836106fe565b601382527f46656551756f74657220312e362e332d64657600000000000000000000000000602083015251918291602083526020830190610311565b602060408183019282815284518094520192019060005b8181106107c35750505090565b82516001600160a01b03168452602093840193909201916001016107b6565b3461024c57600060031936011261024c5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061084f5761046b85610843818703826106fe565b6040519182918261079f565b825484526020909301926001928301920161082c565b3461024c57600060031936011261024c57602060405160248152f35b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c5780600401906040600319823603011261024c576108bf613f33565b6108c98280612d12565b4263ffffffff1692915060005b818110610a22575050602401906108ed8284612d12565b92905060005b8381106108fc57005b8061091b610916600193610910868a612d12565b906129da565b612d93565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff6109fe6109f060208501946109e261096487516001600160e01b031690565b61097e61096f610721565b6001600160e01b039092168252565b63ffffffff8c1660208201526109b961099f845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b815160209092015160e01b6001600160e01b0319166001600160e01b0392909216919091179055565b5167ffffffffffffffff1690565b93516001600160e01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108f3565b80610a3b610a366001936109108980612d12565b612d5c565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03610ad46109f06020850194610ac7610a8387516001600160e01b031690565b610a8e61096f610721565b63ffffffff8d1660208201526109b9610aae84516001600160a01b031690565b6001600160a01b03166000526006602052604060002090565b516001600160a01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108d6565b9181601f8401121561024c5782359167ffffffffffffffff831161024c576020838186019501011161024c57565b92610b589492610b4a92855215156020850152608060408501526080840190610311565b916060818403910152610311565b90565b3461024c5760a060031936011261024c57610b7461029d565b60243590610b818261023b565b6044359160643567ffffffffffffffff811161024c57610ba5903690600401610af8565b93909160843567ffffffffffffffff811161024c57610bc8903690600401610af8565b9290917f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b0382166001600160a01b03821614600014610c94575050935b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016808611610c63575091610c54939161046b9693613f77565b90939160405194859485610b26565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91610c9e9261294a565b93610c0d565b67ffffffffffffffff81116106895760051b60200190565b8015150361024c57565b359061025c82610cbc565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013590610d0e82610ca4565b90610d1c60405192836106fe565b828252602460a06020840194028201019036821161024c57602401925b818410610d4b57610d4983612db8565b005b60a08436031261024c5760405190610d628261068e565b8435610d6d8161023b565b8252602085013575ffffffffffffffffffffffffffffffffffffffffffff198116810361024c5760208301526040850135907fffff0000000000000000000000000000000000000000000000000000000000008216820361024c5782602092604060a0950152610ddf60608801610251565b6060820152610df060808801610cc6565b6080820152815201930192610d39565b602060408183019282815284518094520192019060005b818110610e245750505090565b9091926020604082610e52600194885163ffffffff602080926001600160e01b038151168552015116910152565b019401929101610e17565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c57610e8e9036906004016102e0565b610e9781610ca4565b91610ea560405193846106fe565b818352601f19610eb483610ca4565b0160005b818110610f0f57505060005b82811015610f0157600190610ee5610ee08260051b85016129ef565b613a5b565b610eef8287612b0d565b52610efa8186612b0d565b5001610ec4565b6040518061046b8682610e00565b602090610f1a612ee6565b82828801015201610eb8565b3461024c57602060031936011261024c576020610f4d600435610f488161023b565b613d49565b6001600160e01b0360405191168152f35b3461024c57602060031936011261024c5767ffffffffffffffff610f8061029d565b610f88612ee6565b50166000526005602052604060002060405190610fa4826106aa565b546001600160e01b038116825260e01c6020820152604051809161046b82604081019263ffffffff602080926001600160e01b038151168552015116910152565b90604060031983011261024c5760043567ffffffffffffffff8116810361024c579160243567ffffffffffffffff811161024c5760a0906004018092031261024c5790565b906020610b58928181520190610311565b3461024c5761046b6111236111316110e56110ea61105836610fe5565b67ffffffffffffffff8294921660005260096020526110a761108a6040600020546001600160e01b03199060081b1690565b6110a161109a6020850185612a39565b3691612a6c565b90613dc1565b6110ce6110c88567ffffffffffffffff16600052600a602052604060002090565b916129ef565b6001600160a01b0316600052602052604060002090565b612ab3565b906110f860a0830151151590565b1561113d57506060015163ffffffff165b6040805163ffffffff909216602083015290928391820190565b03601f1981018352826106fe565b6040519182918261102a565b611170915061116260019167ffffffffffffffff166000526009602052604060002090565b015460101c63ffffffff1690565b611109565b61025c909291926102408061026083019561119284825115159052565b60208181015161ffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015163ffffffff169085015260a08181015160ff169085015260c08181015160ff169085015260e08181015161ffff16908501526101008181015163ffffffff16908501526101208181015161ffff16908501526101408181015161ffff1690850152610160818101516001600160e01b03191690850152610180818101511515908501526101a08181015161ffff16908501526101c08181015163ffffffff16908501526101e08181015163ffffffff16908501526102008181015167ffffffffffffffff16908501526102208181015163ffffffff1690850152015163ffffffff16910152565b3461024c57602060031936011261024c5761046b61137061136b6112cf61029d565b60006102406112dc610730565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e08201528261020082015282610220820152015267ffffffffffffffff166000526009602052604060002090565b612f24565b60405191829182611175565b359063ffffffff8216820361024c57565b359061ffff8216820361024c57565b81601f8201121561024c578035906113b382610ca4565b926113c160405194856106fe565b82845260208085019360061b8301019181831161024c57602001925b8284106113eb575050505090565b60408483031261024c5760206040918251611405816106aa565b61140e876102cb565b81528287013561141d8161023b565b838201528152019301926113dd565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013561146881610ca4565b9161147660405193846106fe565b8183526024602084019260051b8201019036821161024c5760248101925b8284106114c5576024358567ffffffffffffffff821161024c576114bf610d4992369060040161139c565b9061307a565b833567ffffffffffffffff811161024c5782016040602319823603011261024c57604051906114f3826106aa565b6114ff602482016102cb565b8252604481013567ffffffffffffffff811161024c57602491010136601f8201121561024c57803561153081610ca4565b9161153e60405193846106fe565b818352602060e081850193028201019036821161024c57602001915b8183106115795750505091816020938480940152815201930192611494565b82360360e0811261024c5760c0601f1960405192611596846106aa565b86356115a18161023b565b8452011261024c5760e0916020916040516115bb816106c6565b6115c684880161137c565b81526115d46040880161137c565b848201526115e46060880161138d565b60408201526115f56080880161137c565b606082015261160660a0880161137c565b608082015260c087013561161981610cbc565b60a08201528382015281520192019161155a565b3461024c57600060031936011261024c576000546001600160a01b03811633036116b4577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561024c5781356116f581610ca4565b9261170360405194856106fe565b81845260208085019260051b82010192831161024c57602001905b82821061172b5750505090565b60208091833561173a8161023b565b81520191019061171e565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c576117769036906004016116de565b60243567ffffffffffffffff811161024c576117969036906004016116de565b9061179f614153565b60005b815181101561180e57806117c36117be610ac760019486612b0d565b615b8f565b6117ce575b016117a2565b6001600160a01b036117e3610ac78386612b0d565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26117c8565b8260005b8151811015610d49578061183361182e610ac760019486612b0d565b615ba3565b61183e575b01611812565b6001600160a01b03611853610ac78386612b0d565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611838565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c576118af903690600401610af8565b6024359167ffffffffffffffff831161024c576119086119006118e66118dc611910963690600401610af8565b9490953691612a6c565b90604082015190605e604a84015160601c93015191929190565b9190336142e5565b810190613329565b60005b8151811015610d495761195b61195661193d61192f8486612b0d565b51516001600160a01b031690565b6001600160a01b03166000526007602052604060002090565b6133e8565b61196f61196b6040830151151590565b1590565b611ad057906119ba6119876020600194015160ff1690565b6119b46119a860206119998689612b0d565b5101516001600160e01b031690565b6001600160e01b031690565b906143ae565b6119d560406119c98487612b0d565b51015163ffffffff1690565b63ffffffff611a006119f76119f0610aae61192f888b612b0d565b5460e01c90565b63ffffffff1690565b911610611aca57611a4e611a1960406119c98588612b0d565b611a3e611a24610721565b6001600160e01b03851681529163ffffffff166020830152565b6109b9610aae61192f8689612b0d565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03611a8461192f8588612b0d565b611ac0611a9660406119c9888b612b0d565b60405193849316958390929163ffffffff6020916001600160e01b03604085019616845216910152565b0390a25b01611913565b50611ac4565b611b15611ae061192f8486612b0d565b7f06439c6b000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6000fd5b3461024c57604060031936011261024c5761046b611ba1611b3861029d565b67ffffffffffffffff60243591611b4e8361023b565b600060a0604051611b5e816106c6565b828152826020820152826040820152826060820152826080820152015216600052600a6020526040600020906001600160a01b0316600052602052604060002090565b611c1d611c1460405192611bb4846106c6565b5463ffffffff8116845263ffffffff8160201c16602085015261ffff8160401c166040850152611bfb611bee8263ffffffff9060501c1690565b63ffffffff166060860152565b63ffffffff607082901c16608085015260901c60ff1690565b151560a0830152565b6040519182918291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b60ff81160361024c57565b359061025c82611c80565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013590611cd382610ca4565b90611ce160405192836106fe565b82825260246102806020840194028201019036821161024c57602401925b818410611d0f57610d498361341e565b833603610280811261024c57610260601f1960405192611d2e846106aa565b611d37886102cb565b8452011261024c5761028091602091611d4e610730565b611d59848901610cc6565b8152611d676040890161138d565b84820152611d776060890161137c565b6040820152611d886080890161137c565b6060820152611d9960a0890161137c565b6080820152611daa60c08901611c8b565b60a0820152611dbb60e08901611c8b565b60c0820152611dcd610100890161138d565b60e0820152611ddf610120890161137c565b610100820152611df2610140890161138d565b610120820152611e05610160890161138d565b610140820152611e18610180890161046f565b610160820152611e2b6101a08901610cc6565b610180820152611e3e6101c0890161138d565b6101a0820152611e516101e0890161137c565b6101c0820152611e64610200890161137c565b6101e0820152611e7761022089016102cb565b610200820152611e8a610240890161137c565b610220820152611e9d610260890161137c565b61024082015283820152815201930192611cff565b3461024c57600060031936011261024c5760206001600160a01b0360015416604051908152f35b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c576040600319823603011261024c57604051611f16816106aa565b816004013567ffffffffffffffff811161024c57611f3a90600436918501016116de565b8152602482013567ffffffffffffffff811161024c57610d49926004611f6392369201016116de565b6020820152613656565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c57611faa611fa461046b923690600401610af8565b906138f9565b604051918291602083526020830190610311565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013590611ffb82610ca4565b9061200960405192836106fe565b8282526024602083019360061b8201019036821161024c57602401925b81841061203657610d498361397f565b60408436031261024c5760206040918251612050816106aa565b863561205b8161023b565b81526120688388016102cb565b83820152815201930192612026565b3461024c57602060031936011261024c576001600160a01b0360043561209c8161023b565b6120a4612cf3565b5016600052600760205261046b604060002060ff604051916120c58361066d565b546001600160a01b0381168352818160a01c16602084015260a81c16151560408201526040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b3461024c57600060031936011261024c57604051806020600b54918281520190600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db99060005b81811061217e5761046b85610843818703826106fe565b8254845260209093019260019283019201612167565b3461024c57602060031936011261024c5760406121b6600435610ee08161023b565b6121db8251809263ffffffff602080926001600160e01b038151168552015116910152565bf35b3461024c57600060031936011261024c57602060405160128152f35b3461024c5761220736610fe5565b61222861136b8367ffffffffffffffff166000526009602052604060002090565b9161223661196b8451151590565b6125c4576060820161226c61196b61224d836129ef565b6001600160a01b03166000526001600b01602052604060002054151590565b61258357829060408201906122818284612d12565b94905061228f848883614c8d565b918761229d610f48836129ef565b9788936122bb6122b561022085015163ffffffff1690565b826152aa565b966000808b1561254a57505061232061ffff856123459861232c989661236096612353966123176123076101c06122fb6101a06123669f015161ffff1690565b97015163ffffffff1690565b916123118c6129ef565b94612d12565b9690951661539b565b979197969097946129ef565b6001600160a01b03166000526008602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b90612917565b9460009661ffff61237d6101408c015161ffff1690565b166124ee575b5061046b98612360612353610200612454612464996dffffffffffffffffffffffffffff61244c61246c9f9e9b6124476001600160e01b039f9b9c6123e56124479e63ffffffff6123db6124479f6020810190612a39565b9290501690613ba1565b908b60a081016124086124026123fc835160ff1690565b60ff1690565b85612917565b9360e083019161241a835161ffff1690565b9061ffff8216831161247c575b5050505060800151612447916119f79163ffffffff16613bdf565b613bdf565b613ba1565b911690612917565b93015167ffffffffffffffff1690565b91169061292a565b6040519081529081906020820190565b6119f7949650612447959361ffff6124dd6124cc612442966124c66124bf6124b660809960ff6124b06124e49b5160ff1690565b16613bae565b965161ffff1690565b61ffff1690565b90613a4e565b6123606123fc60c08d015160ff1690565b9116613ba1565b9593839550612427565b90939591859897506125138194966dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff16916125316020870187612a39565b905061253d938c6155a0565b9596939190949238612383565b9593509550505061236061235361234561232c61257d6125786119f761024061236699015163ffffffff1690565b61288a565b946129ef565b61258f611b15916129ef565b7f2502348c000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b3461024c57602060031936011261024c576001600160a01b036004356126208161023b565b612628614153565b1633811461268d57807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c578060040135906126f482610ca4565b9061270260405192836106fe565b8282526024602083019360071b8201019036821161024c57602401925b81841061272f57610d4983613bf9565b8336036080811261024c576060601f196040519261274c846106aa565b87356127578161023b565b8452011261024c576080916020916040516127718161066d565b8388013561277e8161023b565b8152604088013561278e81611c80565b8482015260608801356127a081610cbc565b60408201528382015281520193019261271f565b3461024c57604060031936011261024c576004356127d18161023b565b6127d96102b4565b9067ffffffffffffffff82169182600052600960205260ff60406000205416156128465761280961282a92613d49565b92600052600960205263ffffffff60016040600020015460901c16906152aa565b604080516001600160e01b039384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc1000014901517156128ac57565b612874565b908160051b91808304602014901517156128ac57565b9061012c82029180830461012c14901517156128ac57565b9061010c82029180830461010c14901517156128ac57565b90655af3107a4000820291808304655af3107a400014901517156128ac57565b818102929181159184041417156128ac57565b8115612934570490565b634e487b7160e01b600052601260045260246000fd5b61297461296e610b5894936001600160e01b036129678195613d49565b1690612917565b92613d49565b169061292a565b9061298582610ca4565b61299260405191826106fe565b828152601f196129a28294610ca4565b019060005b8281106129b357505050565b8060606020809385010152016129a7565b634e487b7160e01b600052603260045260246000fd5b91908110156129ea5760061b0190565b6129c4565b35610b588161023b565b91908110156129ea5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561024c570190565b903590601e198136030182121561024c570180359067ffffffffffffffff821161024c5760200191813603831361024c57565b92919267ffffffffffffffff82116106895760405191612a96601f8201601f1916602001846106fe565b82948184528183011161024c578281602093846000960137010152565b9061025c604051612ac3816106c6565b925463ffffffff8082168552602082811c821690860152604082811c61ffff1690860152605082901c81166060860152607082901c16608085015260901c60ff16151560a0840152565b80518210156129ea5760209160051b010190565b909291612b56612b458367ffffffffffffffff166000526009602052604060002090565b5460081b6001600160e01b03191690565b90612b608161297b565b9560005b828110612b75575050505050505090565b612b88612b838284896129da565b6129ef565b8388612ba2612b988584846129f9565b6040810190612a39565b905060208111612c78575b508392612bdc612bd661109a612bcc600198612c12976110e5976129f9565b6020810190612a39565b89613dc1565b612bfa8967ffffffffffffffff16600052600a602052604060002090565b906001600160a01b0316600052602052604060002090565b60a081015115612c4f57612c33611109606061112393015163ffffffff1690565b612c3d828b612b0d565b52612c48818a612b0d565b5001612b64565b50611123612c33611170846111628a67ffffffffffffffff166000526009602052604060002090565b915050612cb06119f7612ca384612bfa8b67ffffffffffffffff16600052600a602052604060002090565b5460701c63ffffffff1690565b10612cbd57838838612bad565b7f36f536ca000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b60405190612d008261066d565b60006040838281528260208201520152565b903590601e198136030182121561024c570180359067ffffffffffffffff821161024c57602001918160061b3603831361024c57565b35906001600160e01b038216820361024c57565b60408136031261024c57612d8b602060405192612d78846106aa565b8035612d838161023b565b845201612d48565b602082015290565b60408136031261024c57612d8b602060405192612daf846106aa565b612d83816102cb565b90612dc1614153565b60005b8251811015612ee15780612dda60019285612b0d565b517f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a360a06001600160a01b038351169260608101936001600160a01b0380865116957fffff000000000000000000000000000000000000000000000000000000000000612e6f602086019475ffffffffffffffffffffffffffffffffffffffffffff19865116604088019a848c511692615b14565b9775ffffffffffffffffffffffffffffffffffffffffffff196080870195612eb4875115158c600052600460205260406000209060ff60ff1983541691151516179055565b8560405198511688525116602087015251166040850152511660608301525115156080820152a201612dc4565b509050565b60405190612ef3826106aa565b60006020838281520152565b90604051612f0c816106aa565b91546001600160e01b038116835260e01c6020830152565b9061025c61306c6001612f35610730565b9461300b6130018254612f51612f4b8260ff1690565b15158a52565b61ffff600882901c1660208a015263ffffffff601882901c1660408a015263ffffffff603882901c1660608a015263ffffffff605882901c1660808a015260ff607882901c1660a08a015260ff608082901c1660c08a015261ffff608882901c1660e08a015263ffffffff609882901c166101008a015261ffff60b882901c166101208a015261ffff60c882901c166101408a01526001600160e01b0319600882901b166101608a015260f81c90565b1515610180880152565b015461ffff81166101a086015263ffffffff601082901c166101c086015263ffffffff603082901c166101e086015267ffffffffffffffff605082901c1661020086015263ffffffff609082901c1661022086015260b01c63ffffffff1690565b63ffffffff16610240840152565b90613083614153565b6000915b80518310156132755761309a8382612b0d565b51906130ae825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b85518051821015613260576130da82602092612b0d565b5101516130eb61192f838951612b0d565b8151602083015163ffffffff908116911681811015613227575050608082015163ffffffff16602081106131e6575090867f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b56001600160a01b0384613175858f60019998612bfa6131709267ffffffffffffffff16600052600a602052604060002090565b614191565b6131dd60405192839216958291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b0390a3016130c3565b7f24ecdc02000000000000000000000000000000000000000000000000000000006000526001600160a01b0390911660045263ffffffff1660245260446000fd5b7f0b4f67a20000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b50509550925092600191500191929092613087565b50905060005b815181101561332557806132a361329460019385612b0d565b515167ffffffffffffffff1690565b67ffffffffffffffff6001600160a01b036132d260206132c38689612b0d565b5101516001600160a01b031690565b60006132f682612bfa8767ffffffffffffffff16600052600a602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a30161327b565b5050565b60208183031261024c5780359067ffffffffffffffff821161024c570181601f8201121561024c5780359061335d82610ca4565b9261336b60405194856106fe565b8284526020606081860194028301019181831161024c57602001925b828410613395575050505090565b60608483031261024c5760206060916040516133b08161066d565b86356133bb8161023b565b81526133c8838801612d48565b838201526133d86040880161137c565b6040820152815201930192613387565b906040516133f58161066d565b604060ff8294546001600160a01b0381168452818160a01c16602085015260a81c161515910152565b90613427614153565b60005b8251811015612ee15761343d8184612b0d565b51602061344d6132948487612b0d565b9101519067ffffffffffffffff811680158015613637575b8015613609575b8015613576575b61353e579161350482600195946134b46134a7612b456135099767ffffffffffffffff166000526009602052604060002090565b6001600160e01b03191690565b61350f577f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806134e78782611175565b0390a267ffffffffffffffff166000526009602052604060002090565b61447f565b0161342a565b7f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e3604051806134e78782611175565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b506001600160e01b03196135966101608501516001600160e01b03191690565b16630a04b54b60e21b81141590816135f7575b816135e5575b816135d3575b816135c1575b50613473565b63647e2ba960e01b91501415386135bb565b63c4e0595360e01b81141591506135b5565b632b1dfffb60e21b81141591506135af565b6307842f7160e21b81141591506135a9565b506101e083015163ffffffff1663ffffffff61362f6119f7606087015163ffffffff1690565b91161161346c565b5063ffffffff61364f6101e085015163ffffffff1690565b1615613465565b61365e614153565b60208101519160005b83518110156136eb5780613680610ac760019387612b0d565b6136a261369d6001600160a01b0383165b6001600160a01b031690565b615f54565b6136ae575b5001613667565b6040516001600160a01b039190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a1386136a7565b5091505160005b815181101561332557613708610ac78284612b0d565b906001600160a01b0382161561377e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6137758361375a6137556136916001976001600160a01b031690565b615edb565b506040516001600160a01b0390911681529081906020820190565b0390a1016136f2565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b909291928360041161024c57831161024c57600401916003190190565b9060041161024c5790600490565b919091356001600160e01b0319811692600481106137ef575050565b6001600160e01b0319929350829060040360031b1b161690565b9080601f8301121561024c57813561382081610ca4565b9261382e60405194856106fe565b81845260208085019260051b82010192831161024c57602001905b8282106138565750505090565b8135815260209182019101613849565b60208183031261024c5780359067ffffffffffffffff821161024c570160a08183031261024c576040519161389a8361068e565b6138a38261137c565b83526138b1602083016102cb565b602084015260408201356138c481610cbc565b604084015260608201356060840152608082013567ffffffffffffffff811161024c576138f19201613809565b608082015290565b908060041161024c577fe0c4c5460000000000000000000000000000000000000000000000000000000082356001600160e01b03191601613968576139448161394c926060946137a8565b810190613866565b015160408051602080820193909352918252610b5890826106fe565b50506040516139786020826106fe565b6000815290565b613987614153565b60005b815181101561332557806001600160a01b036139a860019385612b0d565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d613a4567ffffffffffffffff60206139e48689612b0d565b51015116836000526008602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a20161398a565b919082039182116128ac57565b613a63612ee6565b50613a89613a84826001600160a01b03166000526006602052604060002090565b612eff565b6020810191613aa8613aa26119f7855163ffffffff1690565b42613a4e565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001611613b5057611956613af4916001600160a01b03166000526007602052604060002090565b613b0461196b6040830151151590565b8015613b56575b613b5057613b1890614b17565b9163ffffffff613b406119f7613b35602087015163ffffffff1690565b935163ffffffff1690565b911610613b4b575090565b905090565b50905090565b506001600160a01b03613b7082516001600160a01b031690565b1615613b0b565b90600282018092116128ac57565b90602082018092116128ac57565b90600182018092116128ac57565b919082018092116128ac57565b9061ffff8091169116029061ffff82169182036128ac57565b63ffffffff60209116019063ffffffff82116128ac57565b9063ffffffff8091169116019063ffffffff82116128ac57565b90613c02614153565b60005b8251811015612ee15780613c1b60019285612b0d565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf613d4060206001600160a01b038451169301518360005260076020526040600020613ca06001600160a01b0383511682906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b602082015181547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000075ff0000000000000000000000000000000000000000006040870151151560a81b169360a01b169116171790556040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b0390a201613c05565b613d5281613a5b565b9063ffffffff602083015116158015613daf575b613d785750516001600160e01b031690565b6001600160a01b03907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b506001600160e01b0382511615613d66565b6001600160e01b0319169190630a04b54b60e21b8314613e5a576307842f7160e21b8314613e4c57632b1dfffb60e21b8314613e415763647e2ba960e01b8314613e365763c4e0595360e01b8314613e285782632ee8207560e01b60005260045260246000fd5b61025c919250600b90615736565b61025c919250615797565b61025c9192506156d3565b61025c919250600190615736565b61025c919250615654565b916001600160e01b03198316630a04b54b60e21b8114613f27576307842f7160e21b8114613f0757632b1dfffb60e21b8114613efb5763647e2ba960e01b8114613eef5763c4e0595360e01b14613ed557632ee8207560e01b6000526001600160e01b0319831660045260246000fd5b61025c925015613ee757600b90615736565b600090615736565b505061025c9150615797565b505061025c91506156d3565b5061025c925015613f1e5760ff60015b1690615736565b60ff6000613f17565b505061025c9150615654565b33600052600360205260406000205415613f4957565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b61136b613f9e9196949395929667ffffffffffffffff166000526009602052604060002090565b946101608601946001600160e01b0319613fc087516001600160e01b03191690565b16630a04b54b60e21b8114908115614142575b8115614131575b506140ec5763c4e0595360e01b6001600160e01b031961400288516001600160e01b03191690565b16146140ce5750506307842f7160e21b6001600160e01b031961402d86516001600160e01b03191690565b161461406557611b1561404885516001600160e01b03191690565b632ee8207560e01b6000526001600160e01b031916600452602490565b6140b9935061109a60606140a38763ffffffff61409a610180614092866140c79b9d015163ffffffff1690565b930151151590565b911685876159cf565b0151604051958691602083019190602083019252565b03601f1981018652856106fe565b9160019190565b94509450610b58916140e1913691612a6c565b936001933691612a6c565b945094916141129161410c6119f76101e0610b5896015163ffffffff1690565b91615813565b93614129602061412187615957565b960151151590565b933691612a6c565b63647e2ba960e01b91501438613fda565b632b1dfffb60e21b81149150613fd3565b6001600160a01b0360015416330361416757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b815181546020808501516040808701517fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000090941663ffffffff958616179190921b67ffffffff00000000161791901b69ffff000000000000000016178255606083015161025c936142a19260a092614243911685547fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff1660509190911b6dffffffff0000000000000000000016178555565b61429a614257608083015163ffffffff1690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911b71ffffffff000000000000000000000000000016178555565b0151151590565b81547fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1690151560901b72ff00000000000000000000000000000000000016179055565b919290926142f582828686615b14565b600052600460205260ff60406000205416156143115750505050565b6040517f097e17ff0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015293909216602484015275ffffffffffffffffffffffffffffffffffffffffffff1990911660448301527fffff000000000000000000000000000000000000000000000000000000000000166064820152608490fd5b0390fd5b604d81116128ac57600a0a90565b60ff1660120160ff81116128ac5760ff1690602482111561442b5760231982019182116128ac576143e16143e7926143a0565b9061292a565b6001600160e01b038111614401576001600160e01b031690565b7f10cb51d10000000000000000000000000000000000000000000000000000000060005260046000fd5b9060240390602482116128ac57612360614444926143a0565b6143e7565b9060ff80911691160160ff81116128ac5760ff1690602482111561442b5760231982019182116128ac576143e16143e7926143a0565b90614a63610240600161025c946144ac6144998651151590565b829060ff60ff1983541691151516179055565b6144f26144be602087015161ffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff1660089190911b62ffff0016178255565b61453e614506604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff1660189190911b66ffffffff00000016178255565b61458e614552606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffff1660389190911b6affffffff0000000000000016178255565b6145e26145a2608087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff1660589190911b6effffffff000000000000000000000016178255565b6146346145f360a087015160ff1690565b82547fffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffff1660789190911b6fff00000000000000000000000000000016178255565b61468761464560c087015160ff1690565b82547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1660809190911b70ff0000000000000000000000000000000016178255565b6146dd61469960e087015161ffff1690565b82547fffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffff1660889190911b72ffff000000000000000000000000000000000016178255565b61473a6146f261010087015163ffffffff1690565b82547fffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffff1660989190911b76ffffffff0000000000000000000000000000000000000016178255565b61479761474d61012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b6147f66147aa61014087015161ffff1690565b82547fffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffff1660c89190911b7affff0000000000000000000000000000000000000000000000000016178255565b61485f61480f6101608701516001600160e01b03191690565b82547fff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff1660089190911c7effffffff00000000000000000000000000000000000000000000000000000016178255565b6148c0614870610180870151151590565b82547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178255565b01926149046148d56101a083015161ffff1690565b859061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6149506149196101c083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178555565b6149a06149656101e083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1660309190911b69ffffffff00000000000016178555565b6149fc6149b961020083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff1660509190911b71ffffffffffffffff0000000000000000000016178555565b614a58614a1161022083015163ffffffff1690565b85547fffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffff1660909190911b75ffffffff00000000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff0000000000000000000000000000000000000000000083549260b01b169116179055565b519069ffffffffffffffffffff8216820361024c57565b908160a091031261024c57614ad881614aad565b91602082015191604081015191610b58608060608401519301614aad565b6040513d6000823e3d90fd5b9081602091031261024c5751610b5881611c80565b614b1f612ee6565b50614b3761369161369183516001600160a01b031690565b90604051907ffeaf968c00000000000000000000000000000000000000000000000000000000825260a082600481865afa928315614c3f57600092600094614c44575b5060008312614401576020600491604051928380927f313ce5670000000000000000000000000000000000000000000000000000000082525afa928315614c3f57610b589363ffffffff93614be093600092614c09575b506020015160ff165b90614449565b92614bfb614bec610721565b6001600160e01b039095168552565b1663ffffffff166020830152565b614bda919250614c30602091823d8411614c38575b614c2881836106fe565b810190614b02565b929150614bd1565b503d614c1e565b614af6565b909350614c6a91925060a03d60a011614c77575b614c6281836106fe565b810190614ac4565b5093925050919238614b7a565b503d614c58565b9081602091031261024c573590565b9190614c9c6020830183612a39565b93905060408301614cad8185612d12565b90506040840191614cc56119f7845163ffffffff1690565b8088116152785750602085015161ffff168083116152425750610160850196614cf688516001600160e01b03191690565b6001600160e01b03198116630a04b54b60e21b81148015615232575b8015615222575b15614d9457505050505050509181614d8e61109a614d75614d8796614d446080610b58980186612a39565b6101e083015163ffffffff169063ffffffff614d6d610180614121606088015163ffffffff1690565b941692615d10565b51958694516001600160e01b03191690565b9280612a39565b90613e65565b63c4e0595360e01b819b9a939495979996989b14600014614ff2575050614e16614df5614e09999a614dc96080880188612a39565b63ffffffff614ded610180614de5606087015163ffffffff1690565b950151151590565b931691615c59565b918251998a91516001600160e01b03191690565b614d8e61109a8880612a39565b6060810151519082614e33614e2b8780612a39565b810190614c7e565b614fd5575081614f9f575b8515159081614f92575b50614f685760408111614f365750614e6d90614e6785949397956128df565b90613ba1565b946000935b838510614ecb5750505050506119f7614e8f915163ffffffff1690565b808211614e9b57505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614f0b6119f7612ca3614efa8667ffffffffffffffff16600052600a602052604060002090565b6110ce612b838d6109108b8d612d12565b8015614f2657614f1a91613ba1565b965b0193929190614e72565b50614f3090613b85565b96614f1c565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614e48565b611b15827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614fec919350614e67614fe784613b93565b6128b1565b91614e3e565b6307842f7160e21b03615204575061505f61503c614e09999a6150186080880188612a39565b63ffffffff615034610180614de5606087015163ffffffff1690565b9316916159cf565b9161504e6119f7845163ffffffff1690565b998a91516001600160e01b03191690565b6080810151519082615074614e2b8780612a39565b6151ec5750816151b6575b851515806151aa575b614f685760408211615176576020015167ffffffffffffffff9081169081831c1661513c5750506150c090614e6785949397956128c7565b946000935b8385106150e25750505050506119f7614e8f915163ffffffff1690565b90919293956001906151116119f7612ca3614efa8667ffffffffffffffff16600052600a602052604060002090565b801561512c5761512091613ba1565b965b01939291906150c5565b5061513690613b85565b96615122565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115615088565b611b15827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b6151fe919350614e67614fe784613b77565b9161507f565b632ee8207560e01b6000526001600160e01b03191660045260246000fd5b5063647e2ba960e01b8114614d19565b50632b1dfffb60e21b8114614d12565b7fd88dddd600000000000000000000000000000000000000000000000000000000600052600483905261ffff1660245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff81166000526005602052604060002091604051926152d0846106aa565b546001600160e01b038116845260e01c9182602085015263ffffffff8216928361530a575b50505050610b5890516001600160e01b031690565b63ffffffff1642908103939084116128ac57831161532857806152f5565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b60408136031261024c57602060405191615384836106aa565b803561538f8161023b565b83520135602082015290565b9694919695929390956000946000986000986000965b8088106153c5575050505050505050929190565b9091929394959697999a6153e26153dd8a848b6129da565b61536b565b9a6154196110e58d6110ce61540b8967ffffffffffffffff16600052600a602052604060002090565b91516001600160a01b031690565b9161542a61196b60a0850151151590565b6155735760009c60408401906154456124bf835161ffff1690565b6154fb575b5050606083015163ffffffff1661546091613bdf565b9c60808301516154739063ffffffff1690565b61547c91613bdf565b9b825161548c9063ffffffff1690565b63ffffffff1661549b9061288a565b600193908083106154ef57506125786119f760206154be93015163ffffffff1690565b8082116154de57506154cf91613ba1565b985b01969594939291906153b1565b90506154e991613ba1565b986154d1565b9150506154e991613ba1565b90612360615564939f61555261555b9460208f8e6124bf95506001600160a01b0361552d85516001600160a01b031690565b91166001600160a01b0382161461556c576155489150613d49565b915b015190615d4e565b925161ffff1690565b620186a0900490565b9b388061544a565b509161554a565b999b50600191506155948461558e61559a93614e678b61288a565b9b613bdf565b9c613bc7565b9a6154d1565b91939093806101e00193846101e0116128ac5761012081029080820461012014901517156128ac576101e09101018093116128ac576124bf610140615636610b58966dffffffffffffffffffffffffffff61244c61562161560e6156409a63ffffffff6123609a1690613ba1565b6123606124bf6101208c015161ffff1690565b614e676119f76101008b015163ffffffff1690565b93015161ffff1690565b6128f7565b9081602091031261024c575190565b602081510361568a576156706020825183010160208301615645565b6001600160a01b0381119081156156c7575b5061568a5750565b61439c906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b61040091501038615682565b60208151036156f957600b6156f16020835184010160208401615645565b106156f95750565b61439c906040519182917fe0d7fb020000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b90602082510361575c5780615749575050565b6156f16020835184010160208401615645565b6040517fe0d7fb02000000000000000000000000000000000000000000000000000000008152602060048201528061439c6024820185610311565b60248151036157ad576022810151156157ad5750565b61439c906040519182917f373b0e440000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b9081604091031261024c57602060405191615804836106aa565b805183520151612d8b81610cbc565b9061581c612ee6565b5080156159445761584461109a828061583e6001600160e01b031995876137d3565b956137a8565b91167f302326cb00000000000000000000000000000000000000000000000000000000810361588e57505067ffffffffffffffff615880610721565b911681526000602082015290565b9091507f181dcf100000000000000000000000000000000000000000000000000000000081036158cc575080602080610b58935183010191016157ea565b7f97a657c9000000000000000000000000000000000000000000000000000000001461591c577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060208061592f93518301019101615645565b615937610721565b9081526000602082015290565b505067ffffffffffffffff615880610721565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610b586064826106fe565b604051906159ae8261068e565b60606080836000815260006020820152600060408201526000838201520152565b6159d76159a1565b508115615aea577f1f3b3aba000000000000000000000000000000000000000000000000000000006001600160e01b0319615a1b615a1585856137c5565b906137d3565b1603615ac05781615a2f92613944926137a8565b9180615aaa575b615a805763ffffffff615a4d835163ffffffff1690565b1611615a565790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fee433e990000000000000000000000000000000000000000000000000000000060005260046000fd5b50615abb61196b6040840151151590565b615a36565b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b604080516001600160a01b0392831660208201908152929093169083015275ffffffffffffffffffffffffffffffffffffffffffff1990921660608201527fffff000000000000000000000000000000000000000000000000000000000000909216608083015290615b898160a08101611123565b51902090565b6001600160a01b03610b589116600b615de2565b6001600160a01b03610b589116600b615f16565b60405190615bc4826106e2565b606080836000815260006020820152600060408201520152565b60208183031261024c5780359067ffffffffffffffff821161024c570160808183031261024c5760405191615c12836106e2565b813583526020820135615c2481610cbc565b602084015260408201356040840152606082013567ffffffffffffffff811161024c57615c519201613809565b606082015290565b615c61615bb7565b508115615aea577f21ea4ca9000000000000000000000000000000000000000000000000000000006001600160e01b0319615c9f615a1585856137c5565b1603615ac05781615cbb92615cb3926137a8565b810190615bde565b9180615cfa575b615a8057815111615cd05790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50615d0b61196b6020840151151590565b615cc2565b9063ffffffff615d2d93959495615d25612ee6565b501691615813565b91825111615cd05780615d42575b615a805790565b50602081015115615d3b565b670de0b6b3a7640000916001600160e01b03615d6a9216612917565b0490565b80548210156129ea5760005260206000200190600090565b91615da0918354906000199060031b92831b921b19161790565b9055565b80548015615dcc576000190190615dbb8282615d6e565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615e945760001984018481116128ac5783549360001985019485116128ac576000958583615e4597615e369503615e4b575b505050615da4565b90600052602052604060002090565b55600190565b615e7b615e7591615e6c615e62615e8b9588615d6e565b90549060031b1c90565b92839187615d6e565b90615d86565b8590600052602052604060002090565b55388080615e2e565b50505050600090565b805490680100000000000000008210156106895781615ec4916001615da094018155615d6e565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054615f1057615ef9816002615e9d565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615f4d5780615f3883600193615e9d565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615f4d576000198201908282116128ac576002549260001984019384116128ac578383615e459460009603615fb4575b505050615fa36002615da4565b600390600052602052604060002090565b615fa3615e7591615fcc615e62615fd6956002615d6e565b9283916002615d6e565b55388080615f9656fea164736f6c634300081a000a",
}

var FeeQuoterV2ABI = FeeQuoterV2MetaData.ABI

var FeeQuoterV2Bin = FeeQuoterV2MetaData.Bin

func DeployFeeQuoterV2(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []FeeQuoterTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoterV2, error) {
	parsed, err := FeeQuoterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterV2Bin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
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

func (_FeeQuoterV2 *FeeQuoterV2Caller) GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "getTokenPriceFeedConfig", token)

	if err != nil {
		return *new(FeeQuoterTokenPriceFeedConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenPriceFeedConfig)).(*FeeQuoterTokenPriceFeedConfig)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoterV2.Contract.GetTokenPriceFeedConfig(&_FeeQuoterV2.CallOpts, token)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoterV2.Contract.GetTokenPriceFeedConfig(&_FeeQuoterV2.CallOpts, token)
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

func (_FeeQuoterV2 *FeeQuoterV2Caller) ProcessPoolReturnDataNew(opts *bind.CallOpts, destChainSelector uint64, tokenTransfer InternalEVMTokenTransfer) ([]byte, error) {
	var out []interface{}
	err := _FeeQuoterV2.contract.Call(opts, &out, "processPoolReturnDataNew", destChainSelector, tokenTransfer)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FeeQuoterV2 *FeeQuoterV2Session) ProcessPoolReturnDataNew(destChainSelector uint64, tokenTransfer InternalEVMTokenTransfer) ([]byte, error) {
	return _FeeQuoterV2.Contract.ProcessPoolReturnDataNew(&_FeeQuoterV2.CallOpts, destChainSelector, tokenTransfer)
}

func (_FeeQuoterV2 *FeeQuoterV2CallerSession) ProcessPoolReturnDataNew(destChainSelector uint64, tokenTransfer InternalEVMTokenTransfer) ([]byte, error) {
	return _FeeQuoterV2.Contract.ProcessPoolReturnDataNew(&_FeeQuoterV2.CallOpts, destChainSelector, tokenTransfer)
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

func (_FeeQuoterV2 *FeeQuoterV2Transactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "onReport", metadata, report)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.OnReport(&_FeeQuoterV2.TransactOpts, metadata, report)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.OnReport(&_FeeQuoterV2.TransactOpts, metadata, report)
}

func (_FeeQuoterV2 *FeeQuoterV2Transactor) SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "setReportPermissions", permissions)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.SetReportPermissions(&_FeeQuoterV2.TransactOpts, permissions)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.SetReportPermissions(&_FeeQuoterV2.TransactOpts, permissions)
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

func (_FeeQuoterV2 *FeeQuoterV2Transactor) UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoterV2.contract.Transact(opts, "updateTokenPriceFeeds", tokenPriceFeedUpdates)
}

func (_FeeQuoterV2 *FeeQuoterV2Session) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.UpdateTokenPriceFeeds(&_FeeQuoterV2.TransactOpts, tokenPriceFeedUpdates)
}

func (_FeeQuoterV2 *FeeQuoterV2TransactorSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoterV2.Contract.UpdateTokenPriceFeeds(&_FeeQuoterV2.TransactOpts, tokenPriceFeedUpdates)
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

type FeeQuoterV2PriceFeedPerTokenUpdatedIterator struct {
	Event *FeeQuoterV2PriceFeedPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2PriceFeedPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2PriceFeedPerTokenUpdated)
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
		it.Event = new(FeeQuoterV2PriceFeedPerTokenUpdated)
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

func (it *FeeQuoterV2PriceFeedPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2PriceFeedPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2PriceFeedPerTokenUpdated struct {
	Token           common.Address
	PriceFeedConfig FeeQuoterTokenPriceFeedConfig
	Raw             types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2PriceFeedPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2PriceFeedPerTokenUpdatedIterator{contract: _FeeQuoterV2.contract, event: "PriceFeedPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2PriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2PriceFeedPerTokenUpdated)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterV2PriceFeedPerTokenUpdated, error) {
	event := new(FeeQuoterV2PriceFeedPerTokenUpdated)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterV2ReportPermissionSetIterator struct {
	Event *FeeQuoterV2ReportPermissionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterV2ReportPermissionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterV2ReportPermissionSet)
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
		it.Event = new(FeeQuoterV2ReportPermissionSet)
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

func (it *FeeQuoterV2ReportPermissionSetIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterV2ReportPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterV2ReportPermissionSet struct {
	ReportId   [32]byte
	Permission KeystoneFeedsPermissionHandlerPermission
	Raw        types.Log
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterV2ReportPermissionSetIterator, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.FilterLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterV2ReportPermissionSetIterator{contract: _FeeQuoterV2.contract, event: "ReportPermissionSet", logs: logs, sub: sub}, nil
}

func (_FeeQuoterV2 *FeeQuoterV2Filterer) WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2ReportPermissionSet, reportId [][32]byte) (event.Subscription, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoterV2.contract.WatchLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterV2ReportPermissionSet)
				if err := _FeeQuoterV2.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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

func (_FeeQuoterV2 *FeeQuoterV2Filterer) ParseReportPermissionSet(log types.Log) (*FeeQuoterV2ReportPermissionSet, error) {
	event := new(FeeQuoterV2ReportPermissionSet)
	if err := _FeeQuoterV2.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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
	return common.HexToHash("0x71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad")
}

func (FeeQuoterV2DestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e3")
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

func (FeeQuoterV2PriceFeedPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0xe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf")
}

func (FeeQuoterV2ReportPermissionSet) Topic() common.Hash {
	return common.HexToHash("0x32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3")
}

func (FeeQuoterV2TokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (FeeQuoterV2TokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5")
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

	GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error)

	GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error)

	GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

		error)

	ProcessPoolReturnData(opts *bind.CallOpts, destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error)

	ProcessPoolReturnDataNew(opts *bind.CallOpts, destChainSelector uint64, tokenTransfer InternalEVMTokenTransfer) ([]byte, error)

	ResolveTokenReceiver(opts *bind.CallOpts, extraArgs []byte) ([]byte, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []common.Address) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error)

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

	FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterV2PriceFeedPerTokenUpdatedIterator, error)

	WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2PriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterV2PriceFeedPerTokenUpdated, error)

	FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterV2ReportPermissionSetIterator, error)

	WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterV2ReportPermissionSet, reportId [][32]byte) (event.Subscription, error)

	ParseReportPermissionSet(log types.Log) (*FeeQuoterV2ReportPermissionSet, error)

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
