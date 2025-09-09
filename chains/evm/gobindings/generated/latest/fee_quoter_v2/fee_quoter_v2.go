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
	Bin: "0x60e0604052346110b257617677803803806100198161131e565b92833981019080820361012081126110b2576060136110b25761003a6112e0565b81516001600160601b03811681036110b257815261005a60208301611343565b906020810191825261006e60408401611357565b6040820190815260608401516001600160401b0381116110b2578561009491860161137f565b60808501519094906001600160401b0381116110b257866100b691830161137f565b60a08201519096906001600160401b0381116110b25782019080601f830112156110b25781516100ed6100e882611368565b61131e565b9260208085848152019260071b820101908382116110b257602001915b81831061126b5750505060c08301516001600160401b0381116110b25783019781601f8a0112156110b2578851986101446100e88b611368565b996020808c838152019160051b830101918483116110b25760208101915b838310611109575050505060e08401516001600160401b0381116110b25784019382601f860112156110b257845161019c6100e882611368565b9560208088848152019260061b820101908582116110b257602001915b8183106110cd57505050610100810151906001600160401b0382116110b2570182601f820112156110b2578051906101f36100e883611368565b9360206102808187868152019402830101918183116110b257602001925b828410610ef057505050503315610edf57600180546001600160a01b031916331790556020986102408a61131e565b976000895260003681376102526112ff565b998a52888b8b015260005b89518110156102c4576001906001600160a01b0361027b828d611418565b51168d61028782611604565b610294575b50500161025d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388d61028c565b508a985089519660005b885181101561033f576001600160a01b036102e9828b611418565b511690811561032e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8c8361032060019561158c565b50604051908152a1016102ce565b6342bcdf7f60e11b60005260046000fd5b5081518a985089906001600160a01b0316158015610ecd575b8015610ebe575b610ead5791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c0526103928661131e565b9360008552600036813760005b855181101561040e576001906103c76001600160a01b036103c0838a611418565b5116611499565b6103d2575b0161039f565b818060a01b036103e28289611418565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26103cc565b508694508560005b84518110156104855760019061043e6001600160a01b036104378389611418565b51166115cb565b610449575b01610416565b818060a01b036104598288611418565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610443565b508593508460005b835181101561054757806104a360019286611418565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf606089858060a01b038451169301518360005260078b5260406000209060ff878060a01b038251169283898060a01b03198254161781558d8301908151604082549501948460a81b8651151560a81b16918560a01b9060a01b169061ffff60a01b19161717905560405193845251168c8301525115156040820152a20161048d565b5091509160005b8251811015610b45576105618184611418565b51856001600160401b036105758487611418565b5151169101519080158015610b32575b8015610b14575b8015610a92575b610a7e57600081815260098852604090205460019392919060081b6001600160e01b03191661093657807f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806106fc868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a25b60005260098752826040600020825115158382549162ffff008c83015160081b169066ffffffff000000604084015160181b166affffffff00000000000000606085015160381b16926effffffff0000000000000000000000608086015160581b169260ff60781b60a087015160781b169460ff60801b60c088015160801b169161ffff60881b60e089015160881b169063ffffffff60981b6101008a015160981b169361ffff60b81b6101208b015160b81b169661ffff60c81b6101408c015160c81b169963ffffffff60d81b6101608d015160081c169b61018060ff60f81b910151151560f81b169c8f8060f81b039a63ffffffff60d81b199961ffff60c81b199861ffff60b81b199763ffffffff60981b199661ffff60881b199560ff60801b199460ff60781b19936effffffff0000000000000000000000199260ff6affffffff000000000000001992169066ffffffffffffff19161716171617161716171617161716171617161716179063ffffffff60d81b1617178155019061ffff6101a0820151169082549165ffffffff00006101c083015160101b169269ffffffff0000000000006101e084015160301b166a01000000000000000000008860901b0361020085015160501b169263ffffffff60901b61022086015160901b169461024063ffffffff60b01b91015160b01b169563ffffffff60b01b199363ffffffff60901b19926a01000000000000000000008c60901b0319918c8060501b03191617161716171617171790550161054e565b807f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e360405180610a76868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a2610700565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61016083015116630a04b54b60e21b8114159081610b02575b81610af0575b81610ade575b81610acc575b50610593565b63647e2ba960e01b1415905088610ac6565b63c4e0595360e01b8114159150610ac0565b632b1dfffb60e21b8114159150610aba565b6307842f7160e21b8114159150610ab4565b5063ffffffff6101e08301511663ffffffff6060840151161061058c565b5063ffffffff6101e08301511615610585565b84828560005b8151811015610bcb576001906001600160a01b03610b698285611418565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610b9f8689611418565b510151168360005260088252604060002081878060401b0319825416179055604051908152a201610b4b565b83600184610bd88361131e565b9060008252600092610ea8575b909282935b8251851015610de757610bfd8584611418565b5180516001600160401b0316939083019190855b83518051821015610dd657610c27828792611418565b51015184516001600160a01b0390610c40908490611418565b5151169063ffffffff815116908781019163ffffffff8351169081811015610dc15750506080810163ffffffff815116898110610daa575090899291838c52600a8a5260408c20600160a01b6001900386168d528a5260408c2092825163ffffffff169380549180518d1b67ffffffff0000000016916040860192835160401b69ffff000000000000000016966060810195865160501b6dffffffff00000000000000000000169063ffffffff60701b895160701b169260a001998b60ff60901b8c51151560901b169560ff60901b199363ffffffff60701b19926dffffffff000000000000000000001991600160501b60019003191617161716171617171790556040519586525163ffffffff168c8601525161ffff1660408501525163ffffffff1660608401525163ffffffff16608083015251151560a082015260c07f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b591a3600101610c11565b6312766e0160e11b8c52600485905260245260448bfd5b6305a7b3d160e11b8c5260045260245260448afd5b505060019096019593509050610bea565b9150825b8251811015610e69576001906001600160401b03610e098286611418565b515116828060a01b0384610e1d8488611418565b5101511690808752600a855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610deb565b604051615fde908161169982396080518181816105c20152610c1c015260a0518181816105f80152610bcd015260c05181818161061f0152613aca0152f35b610be5565b63d794ef9560e01b60005260046000fd5b5063ffffffff8251161561035f565b5080516001600160601b031615610358565b639b15e16f60e01b60005260046000fd5b83820361028081126110b257610260610f076112ff565b91610f11876113f5565b8352601f1901126110b2576040519161026083016001600160401b038111848210176110b757604052610f46602087016113e8565b8352610f5460408701611409565b6020840152610f6560608701611357565b6040840152610f7660808701611357565b6060840152610f8760a08701611357565b6080840152610f9860c087016113da565b60a0840152610fa960e087016113da565b60c0840152610fbb6101008701611409565b60e0840152610fcd6101208701611357565b610100840152610fe06101408701611409565b610120840152610ff36101608701611409565b610140840152610180860151916001600160e01b0319831683036110b2578360209361016061028096015261102b6101a089016113e8565b61018082015261103e6101c08901611409565b6101a08201526110516101e08901611357565b6101c08201526110646102008901611357565b6101e082015261107761022089016113f5565b61020082015261108a6102408901611357565b61022082015261109d6102608901611357565b61024082015283820152815201930192610211565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040838703126110b25760206040916110e46112ff565b6110ed86611343565b81526110fa8387016113f5565b838201528152019201916101b9565b82516001600160401b0381116110b25782016040818803601f1901126110b2576111316112ff565b9061113e602082016113f5565b825260408101516001600160401b0381116110b257602091010187601f820112156110b25780516111716100e882611368565b91602060e08185858152019302820101908a82116110b257602001915b8183106111ad5750505091816020938480940152815201920191610162565b828b0360e081126110b25760c06111c26112ff565b916111cc86611343565b8352601f1901126110b2576040519160c08301916001600160401b038311848410176110b75760e093602093604052611206848801611357565b815261121460408801611357565b8482015261122460608801611409565b604082015261123560808801611357565b606082015261124660a08801611357565b608082015261125760c088016113e8565b60a08201528382015281520192019161118e565b828403608081126110b25760606112806112ff565b9161128a86611343565b8352601f1901126110b2576080916020916112a36112e0565b6112ae848801611343565b81526112bc604088016113da565b848201526112cc606088016113e8565b60408201528382015281520192019161010a565b60405190606082016001600160401b038111838210176110b757604052565b60408051919082016001600160401b038111838210176110b757604052565b6040519190601f01601f191682016001600160401b038111838210176110b757604052565b51906001600160a01b03821682036110b257565b519063ffffffff821682036110b257565b6001600160401b0381116110b75760051b60200190565b9080601f830112156110b25781516113996100e882611368565b9260208085848152019260051b8201019283116110b257602001905b8282106113c25750505090565b602080916113cf84611343565b8152019101906113b5565b519060ff821682036110b257565b519081151582036110b257565b51906001600160401b03821682036110b257565b519061ffff821682036110b257565b805182101561142c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561142c5760005260206000200190600090565b805480156114835760001901906114718282611442565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600c6020526040902054801561155a57600019810181811161154457600b54600019810191908211611544578181036114f3575b5050506114df600b61145a565b600052600c60205260006040812055600190565b61152c61150461151593600b611442565b90549060031b1c928392600b611442565b819391549060031b91821b91600019901b19161790565b9055600052600c6020526040600020553880806114d2565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156110b7578161151591600161158894018155611442565b9055565b806000526003602052604060002054156000146115c5576115ae816002611561565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146115c5576115ed81600b611561565b600b5490600052600c602052604060002055600190565b600081815260036020526040902054801561155a576000198101818111611544576002546000198101919082116115445780820361165e575b50505061164a600261145a565b600052600360205260006040812055600190565b61168061166f611515936002611442565b90549060031b1c9283926002611442565b9055600052600360205260406000205538808061163d56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be1461023657806301447eaa1461023157806301ffc9a71461022c578063061877e31461022757806306285c6914610222578063181f5a771461021d5780632451a62714610218578063325c868e146102135780633937306f1461020e5780633a49bb491461020957806341ed29e71461020457806345ac924d146101ff5780634ab35b0b146101fa578063514e8cff146101f5578063614a8adc146101f05780636def4ce7146101eb578063770e2dc4146101e657806379ba5097146101e15780637afac322146101dc578063805f2132146101d757806382b49eb0146101d257806387b8d879146101cd5780638da5cb5b146101c857806391a2749a146101c35780639b1115e4146101be578063a69c64c0146101b9578063bf78e03f146101b4578063cdc73d51146101af578063d02641a0146101aa578063d63d3af2146101a5578063d8694ccd146101a0578063f2fde38b1461019b578063fbe3f778146101965763ffdb4b371461019157600080fd5b6127bd565b6126c0565b612604565b612202565b6121e6565b61219d565b612126565b612080565b611fc7565b611f76565b611ee2565b611ebb565b611c9f565b611b22565b611887565b61174e565b611636565b611435565b6112b6565b611044565b610f67565b610f2f565b610e66565b610cd1565b610b5b565b610881565b610865565b6107e2565b610740565b610586565b61053e565b610484565b6103d0565b61025e565b6001600160a01b0381160361024c57565b600080fd5b359061025c8261023b565b565b3461024c57606060031936011261024c5760206102956004356102808161023b565b602435604435916102908361023b565b612953565b604051908152f35b6004359067ffffffffffffffff8216820361024c57565b6024359067ffffffffffffffff8216820361024c57565b359067ffffffffffffffff8216820361024c57565b9181601f8401121561024c5782359167ffffffffffffffff831161024c576020808501948460051b01011161024c57565b919082519283825260005b84811061033d575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161031c565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061038557505050505090565b90919293946020806103c1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610311565b97019301930191939290610376565b3461024c57606060031936011261024c576103e961029d565b60243567ffffffffffffffff811161024c576104099036906004016102e0565b6044929192359167ffffffffffffffff831161024c573660238401121561024c5782600401359167ffffffffffffffff831161024c573660248460061b8601011161024c5761046b94602461045f950192612b2a565b60405191829182610352565b0390f35b35906001600160e01b03198216820361024c57565b3461024c57602060031936011261024c576004356001600160e01b0319811680910361024c57807f805f21320000000000000000000000000000000000000000000000000000000060209214908115610514575b81156104ea575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386104df565b7f66792e8000000000000000000000000000000000000000000000000000000000811491506104d8565b3461024c57602060031936011261024c576001600160a01b036004356105638161023b565b166000526008602052602067ffffffffffffffff60406000205416604051908152f35b3461024c57600060031936011261024c5761059f612cfc565b5060606040516105ae8161066d565b63ffffffff6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918281526001600160a01b0360406020830192827f00000000000000000000000000000000000000000000000000000000000000001684520191837f00000000000000000000000000000000000000000000000000000000000000001683526040519485525116602084015251166040820152f35b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761068957604052565b610657565b60a0810190811067ffffffffffffffff82111761068957604052565b6040810190811067ffffffffffffffff82111761068957604052565b60c0810190811067ffffffffffffffff82111761068957604052565b6080810190811067ffffffffffffffff82111761068957604052565b90601f601f19910116810190811067ffffffffffffffff82111761068957604052565b6040519061025c6040836106fe565b6040519061025c610260836106fe565b3461024c57600060031936011261024c5761046b604080519061076381836106fe565b601382527f46656551756f74657220312e362e332d64657600000000000000000000000000602083015251918291602083526020830190610311565b602060408183019282815284518094520192019060005b8181106107c35750505090565b82516001600160a01b03168452602093840193909201916001016107b6565b3461024c57600060031936011261024c5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061084f5761046b85610843818703826106fe565b6040519182918261079f565b825484526020909301926001928301920161082c565b3461024c57600060031936011261024c57602060405160248152f35b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c5780600401906040600319823603011261024c576108bf613f4e565b6108c98280612d1b565b4263ffffffff1692915060005b818110610a22575050602401906108ed8284612d1b565b92905060005b8381106108fc57005b8061091b610916600193610910868a612d1b565b906129e3565b612d9c565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff6109fe6109f060208501946109e261096487516001600160e01b031690565b61097e61096f610721565b6001600160e01b039092168252565b63ffffffff8c1660208201526109b961099f845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b815160209092015160e01b6001600160e01b0319166001600160e01b0392909216919091179055565b5167ffffffffffffffff1690565b93516001600160e01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108f3565b80610a3b610a366001936109108980612d1b565b612d65565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03610ad46109f06020850194610ac7610a8387516001600160e01b031690565b610a8e61096f610721565b63ffffffff8d1660208201526109b9610aae84516001600160a01b031690565b6001600160a01b03166000526006602052604060002090565b516001600160a01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108d6565b9181601f8401121561024c5782359167ffffffffffffffff831161024c576020838186019501011161024c57565b92610b589492610b4a92855215156020850152608060408501526080840190610311565b916060818403910152610311565b90565b3461024c5760a060031936011261024c57610b7461029d565b60243590610b818261023b565b6044359160643567ffffffffffffffff811161024c57610ba5903690600401610af8565b93909160843567ffffffffffffffff811161024c57610bc8903690600401610af8565b9290917f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b0382166001600160a01b03821614600014610c94575050935b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016808611610c63575091610c54939161046b9693613f92565b90939160405194859485610b26565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91610c9e92612953565b93610c0d565b67ffffffffffffffff81116106895760051b60200190565b8015150361024c57565b359061025c82610cbc565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013590610d0e82610ca4565b90610d1c60405192836106fe565b828252602460a06020840194028201019036821161024c57602401925b818410610d4b57610d4983612dc1565b005b60a08436031261024c5760405190610d628261068e565b8435610d6d8161023b565b825260208501357fffffffffffffffffffff000000000000000000000000000000000000000000008116810361024c5760208301526040850135907fffff0000000000000000000000000000000000000000000000000000000000008216820361024c5782602092604060a0950152610de860608801610251565b6060820152610df960808801610cc6565b6080820152815201930192610d39565b602060408183019282815284518094520192019060005b818110610e2d5750505090565b9091926020604082610e5b600194885163ffffffff602080926001600160e01b038151168552015116910152565b019401929101610e20565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c57610e979036906004016102e0565b610ea081610ca4565b91610eae60405193846106fe565b818352601f19610ebd83610ca4565b0160005b818110610f1857505060005b82811015610f0a57600190610eee610ee98260051b85016129f8565b613a76565b610ef88287612b16565b52610f038186612b16565b5001610ecd565b6040518061046b8682610e09565b602090610f23612f01565b82828801015201610ec1565b3461024c57602060031936011261024c576020610f56600435610f518161023b565b613d64565b6001600160e01b0360405191168152f35b3461024c57602060031936011261024c5767ffffffffffffffff610f8961029d565b610f91612f01565b50166000526005602052604060002060405190610fad826106aa565b546001600160e01b038116825260e01c6020820152604051809161046b82604081019263ffffffff602080926001600160e01b038151168552015116910152565b90604060031983011261024c5760043567ffffffffffffffff8116810361024c579160243567ffffffffffffffff811161024c5760a0906004018092031261024c5790565b906020610b58928181520190610311565b3461024c5761046b61112c61113a6110ee6110f361106136610fee565b67ffffffffffffffff8294921660005260096020526110b06110936040600020546001600160e01b03199060081b1690565b6110aa6110a36020850185612a42565b3691612a75565b90613ddc565b6110d76110d18567ffffffffffffffff16600052600a602052604060002090565b916129f8565b6001600160a01b0316600052602052604060002090565b612abc565b9061110160a0830151151590565b1561114657506060015163ffffffff165b6040805163ffffffff909216602083015290928391820190565b03601f1981018352826106fe565b60405191829182611033565b611179915061116b60019167ffffffffffffffff166000526009602052604060002090565b015460101c63ffffffff1690565b611112565b61025c909291926102408061026083019561119b84825115159052565b60208181015161ffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015163ffffffff169085015260a08181015160ff169085015260c08181015160ff169085015260e08181015161ffff16908501526101008181015163ffffffff16908501526101208181015161ffff16908501526101408181015161ffff1690850152610160818101516001600160e01b03191690850152610180818101511515908501526101a08181015161ffff16908501526101c08181015163ffffffff16908501526101e08181015163ffffffff16908501526102008181015167ffffffffffffffff16908501526102208181015163ffffffff1690850152015163ffffffff16910152565b3461024c57602060031936011261024c5761046b6113796113746112d861029d565b60006102406112e5610730565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e08201528261020082015282610220820152015267ffffffffffffffff166000526009602052604060002090565b612f3f565b6040519182918261117e565b359063ffffffff8216820361024c57565b359061ffff8216820361024c57565b81601f8201121561024c578035906113bc82610ca4565b926113ca60405194856106fe565b82845260208085019360061b8301019181831161024c57602001925b8284106113f4575050505090565b60408483031261024c576020604091825161140e816106aa565b611417876102cb565b8152828701356114268161023b565b838201528152019301926113e6565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013561147181610ca4565b9161147f60405193846106fe565b8183526024602084019260051b8201019036821161024c5760248101925b8284106114ce576024358567ffffffffffffffff821161024c576114c8610d499236906004016113a5565b90613095565b833567ffffffffffffffff811161024c5782016040602319823603011261024c57604051906114fc826106aa565b611508602482016102cb565b8252604481013567ffffffffffffffff811161024c57602491010136601f8201121561024c57803561153981610ca4565b9161154760405193846106fe565b818352602060e081850193028201019036821161024c57602001915b818310611582575050509181602093848094015281520193019261149d565b82360360e0811261024c5760c0601f196040519261159f846106aa565b86356115aa8161023b565b8452011261024c5760e0916020916040516115c4816106c6565b6115cf848801611385565b81526115dd60408801611385565b848201526115ed60608801611396565b60408201526115fe60808801611385565b606082015261160f60a08801611385565b608082015260c087013561162281610cbc565b60a082015283820152815201920191611563565b3461024c57600060031936011261024c576000546001600160a01b03811633036116bd577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561024c5781356116fe81610ca4565b9261170c60405194856106fe565b81845260208085019260051b82010192831161024c57602001905b8282106117345750505090565b6020809183356117438161023b565b815201910190611727565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c5761177f9036906004016116e7565b60243567ffffffffffffffff811161024c5761179f9036906004016116e7565b906117a861416e565b60005b815181101561181757806117cc6117c7610ac760019486612b16565b615b81565b6117d7575b016117ab565b6001600160a01b036117ec610ac78386612b16565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26117d1565b8260005b8151811015610d49578061183c611837610ac760019486612b16565b615b95565b611847575b0161181b565b6001600160a01b0361185c610ac78386612b16565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611841565b3461024c57604060031936011261024c5760043567ffffffffffffffff811161024c576118b8903690600401610af8565b6024359167ffffffffffffffff831161024c576119116119096118ef6118e5611919963690600401610af8565b9490953691612a75565b90604082015190605e604a84015160601c93015191929190565b919033614300565b810190613344565b60005b8151811015610d495761196461195f6119466119388486612b16565b51516001600160a01b031690565b6001600160a01b03166000526007602052604060002090565b613403565b6119786119746040830151151590565b1590565b611ad957906119c36119906020600194015160ff1690565b6119bd6119b160206119a28689612b16565b5101516001600160e01b031690565b6001600160e01b031690565b906143d2565b6119de60406119d28487612b16565b51015163ffffffff1690565b63ffffffff611a09611a006119f9610aae611938888b612b16565b5460e01c90565b63ffffffff1690565b911610611ad357611a57611a2260406119d28588612b16565b611a47611a2d610721565b6001600160e01b03851681529163ffffffff166020830152565b6109b9610aae6119388689612b16565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03611a8d6119388588612b16565b611ac9611a9f60406119d2888b612b16565b60405193849316958390929163ffffffff6020916001600160e01b03604085019616845216910152565b0390a25b0161191c565b50611acd565b611b1e611ae96119388486612b16565b7f06439c6b000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6000fd5b3461024c57604060031936011261024c5761046b611baa611b4161029d565b67ffffffffffffffff60243591611b578361023b565b600060a0604051611b67816106c6565b828152826020820152826040820152826060820152826080820152015216600052600a6020526040600020906001600160a01b0316600052602052604060002090565b611c26611c1d60405192611bbd846106c6565b5463ffffffff8116845263ffffffff8160201c16602085015261ffff8160401c166040850152611c04611bf78263ffffffff9060501c1690565b63ffffffff166060860152565b63ffffffff607082901c16608085015260901c60ff1690565b151560a0830152565b6040519182918291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b60ff81160361024c57565b359061025c82611c89565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c57806004013590611cdc82610ca4565b90611cea60405192836106fe565b82825260246102806020840194028201019036821161024c57602401925b818410611d1857610d4983613439565b833603610280811261024c57610260601f1960405192611d37846106aa565b611d40886102cb565b8452011261024c5761028091602091611d57610730565b611d62848901610cc6565b8152611d7060408901611396565b84820152611d8060608901611385565b6040820152611d9160808901611385565b6060820152611da260a08901611385565b6080820152611db360c08901611c94565b60a0820152611dc460e08901611c94565b60c0820152611dd66101008901611396565b60e0820152611de86101208901611385565b610100820152611dfb6101408901611396565b610120820152611e0e6101608901611396565b610140820152611e21610180890161046f565b610160820152611e346101a08901610cc6565b610180820152611e476101c08901611396565b6101a0820152611e5a6101e08901611385565b6101c0820152611e6d6102008901611385565b6101e0820152611e8061022089016102cb565b610200820152611e936102408901611385565b610220820152611ea66102608901611385565b61024082015283820152815201930192611d08565b3461024c57600060031936011261024c5760206001600160a01b0360015416604051908152f35b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c576040600319823603011261024c57604051611f1f816106aa565b816004013567ffffffffffffffff811161024c57611f4390600436918501016116e7565b8152602482013567ffffffffffffffff811161024c57610d49926004611f6c92369201016116e7565b6020820152613671565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c57611fb3611fad61046b923690600401610af8565b90613914565b604051918291602083526020830190610311565b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c5780600401359061200482610ca4565b9061201260405192836106fe565b8282526024602083019360061b8201019036821161024c57602401925b81841061203f57610d498361399a565b60408436031261024c5760206040918251612059816106aa565b86356120648161023b565b81526120718388016102cb565b8382015281520193019261202f565b3461024c57602060031936011261024c576001600160a01b036004356120a58161023b565b6120ad612cfc565b5016600052600760205261046b604060002060ff604051916120ce8361066d565b546001600160a01b0381168352818160a01c16602084015260a81c16151560408201526040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b3461024c57600060031936011261024c57604051806020600b54918281520190600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db99060005b8181106121875761046b85610843818703826106fe565b8254845260209093019260019283019201612170565b3461024c57602060031936011261024c5760406121bf600435610ee98161023b565b6121e48251809263ffffffff602080926001600160e01b038151168552015116910152565bf35b3461024c57600060031936011261024c57602060405160128152f35b3461024c5761221036610fee565b6122316113748367ffffffffffffffff166000526009602052604060002090565b9161223f6119748451151590565b6125cd5760608201612275611974612256836129f8565b6001600160a01b03166000526001600b01602052604060002054151590565b61258c578290604082019061228a8284612d1b565b949050612298848883614cb1565b91876122a6610f51836129f8565b9788936122c46122be61022085015163ffffffff1690565b826152ce565b966000808b1561255357505061232961ffff8561234e9861233598966123699661235c966123206123106101c06123046101a061236f9f015161ffff1690565b97015163ffffffff1690565b9161231a8c6129f8565b94612d1b565b969095166153bf565b979197969097946129f8565b6001600160a01b03166000526008602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b90612920565b9460009661ffff6123866101408c015161ffff1690565b166124f7575b5061046b9861236961235c61020061245d61246d996dffffffffffffffffffffffffffff6124556124759f9e9b6124506001600160e01b039f9b9c6123ee6124509e63ffffffff6123e46124509f6020810190612a42565b9290501690613bbc565b908b60a0810161241161240b612405835160ff1690565b60ff1690565b85612920565b9360e0830191612423835161ffff1690565b9061ffff82168311612485575b505050506080015161245091611a009163ffffffff16613bfa565b613bfa565b613bbc565b911690612920565b93015167ffffffffffffffff1690565b911690612933565b6040519081529081906020820190565b611a00949650612450959361ffff6124e66124d561244b966124cf6124c86124bf60809960ff6124b96124ed9b5160ff1690565b16613bc9565b965161ffff1690565b61ffff1690565b90613a69565b61236961240560c08d015160ff1690565b9116613bbc565b9593839550612430565b909395918598975061251c8194966dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff169161253a6020870187612a42565b9050612546938c6155c4565b959693919094923861238c565b9593509550505061236961235c61234e612335612586612581611a0061024061236f99015163ffffffff1690565b612893565b946129f8565b612598611b1e916129f8565b7f2502348c000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b3461024c57602060031936011261024c576001600160a01b036004356126298161023b565b61263161416e565b1633811461269657807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024c57602060031936011261024c5760043567ffffffffffffffff811161024c573660238201121561024c578060040135906126fd82610ca4565b9061270b60405192836106fe565b8282526024602083019360071b8201019036821161024c57602401925b81841061273857610d4983613c14565b8336036080811261024c576060601f1960405192612755846106aa565b87356127608161023b565b8452011261024c5760809160209160405161277a8161066d565b838801356127878161023b565b8152604088013561279781611c89565b8482015260608801356127a981610cbc565b604082015283820152815201930192612728565b3461024c57604060031936011261024c576004356127da8161023b565b6127e26102b4565b9067ffffffffffffffff82169182600052600960205260ff604060002054161561284f5761281261283392613d64565b92600052600960205263ffffffff60016040600020015460901c16906152ce565b604080516001600160e01b039384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc1000014901517156128b557565b61287d565b908160051b91808304602014901517156128b557565b9061012c82029180830461012c14901517156128b557565b9061010c82029180830461010c14901517156128b557565b90655af3107a4000820291808304655af3107a400014901517156128b557565b818102929181159184041417156128b557565b811561293d570490565b634e487b7160e01b600052601260045260246000fd5b61297d612977610b5894936001600160e01b036129708195613d64565b1690612920565b92613d64565b1690612933565b9061298e82610ca4565b61299b60405191826106fe565b828152601f196129ab8294610ca4565b019060005b8281106129bc57505050565b8060606020809385010152016129b0565b634e487b7160e01b600052603260045260246000fd5b91908110156129f35760061b0190565b6129cd565b35610b588161023b565b91908110156129f35760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561024c570190565b903590601e198136030182121561024c570180359067ffffffffffffffff821161024c5760200191813603831361024c57565b92919267ffffffffffffffff82116106895760405191612a9f601f8201601f1916602001846106fe565b82948184528183011161024c578281602093846000960137010152565b9061025c604051612acc816106c6565b925463ffffffff8082168552602082811c821690860152604082811c61ffff1690860152605082901c81166060860152607082901c16608085015260901c60ff16151560a0840152565b80518210156129f35760209160051b010190565b909291612b5f612b4e8367ffffffffffffffff166000526009602052604060002090565b5460081b6001600160e01b03191690565b90612b6981612984565b9560005b828110612b7e575050505050505090565b612b91612b8c8284896129e3565b6129f8565b8388612bab612ba1858484612a02565b6040810190612a42565b905060208111612c81575b508392612be5612bdf6110a3612bd5600198612c1b976110ee97612a02565b6020810190612a42565b89613ddc565b612c038967ffffffffffffffff16600052600a602052604060002090565b906001600160a01b0316600052602052604060002090565b60a081015115612c5857612c3c611112606061112c93015163ffffffff1690565b612c46828b612b16565b52612c51818a612b16565b5001612b6d565b5061112c612c3c6111798461116b8a67ffffffffffffffff166000526009602052604060002090565b915050612cb9611a00612cac84612c038b67ffffffffffffffff16600052600a602052604060002090565b5460701c63ffffffff1690565b10612cc657838838612bb6565b7f36f536ca000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b60405190612d098261066d565b60006040838281528260208201520152565b903590601e198136030182121561024c570180359067ffffffffffffffff821161024c57602001918160061b3603831361024c57565b35906001600160e01b038216820361024c57565b60408136031261024c57612d94602060405192612d81846106aa565b8035612d8c8161023b565b845201612d51565b602082015290565b60408136031261024c57612d94602060405192612db8846106aa565b612d8c816102cb565b90612dca61416e565b60005b8251811015612efc5780612de360019285612b16565b517f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a360a06001600160a01b038351169260608101936001600160a01b0380865116957fffff000000000000000000000000000000000000000000000000000000000000612e8160208601947fffffffffffffffffffff00000000000000000000000000000000000000000000865116604088019a848c511692615afd565b977fffffffffffffffffffff000000000000000000000000000000000000000000006080870195612ecf875115158c600052600460205260406000209060ff60ff1983541691151516179055565b8560405198511688525116602087015251166040850152511660608301525115156080820152a201612dcd565b509050565b60405190612f0e826106aa565b60006020838281520152565b90604051612f27816106aa565b91546001600160e01b038116835260e01c6020830152565b9061025c6130876001612f50610730565b9461302661301c8254612f6c612f668260ff1690565b15158a52565b61ffff600882901c1660208a015263ffffffff601882901c1660408a015263ffffffff603882901c1660608a015263ffffffff605882901c1660808a015260ff607882901c1660a08a015260ff608082901c1660c08a015261ffff608882901c1660e08a015263ffffffff609882901c166101008a015261ffff60b882901c166101208a015261ffff60c882901c166101408a01526001600160e01b0319600882901b166101608a015260f81c90565b1515610180880152565b015461ffff81166101a086015263ffffffff601082901c166101c086015263ffffffff603082901c166101e086015267ffffffffffffffff605082901c1661020086015263ffffffff609082901c1661022086015260b01c63ffffffff1690565b63ffffffff16610240840152565b9061309e61416e565b6000915b8051831015613290576130b58382612b16565b51906130c9825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b8551805182101561327b576130f582602092612b16565b510151613106611938838951612b16565b8151602083015163ffffffff908116911681811015613242575050608082015163ffffffff1660208110613201575090867f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b56001600160a01b0384613190858f60019998612c0361318b9267ffffffffffffffff16600052600a602052604060002090565b6141ac565b6131f860405192839216958291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b0390a3016130de565b7f24ecdc02000000000000000000000000000000000000000000000000000000006000526001600160a01b0390911660045263ffffffff1660245260446000fd5b7f0b4f67a20000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b505095509250926001915001919290926130a2565b50905060005b815181101561334057806132be6132af60019385612b16565b515167ffffffffffffffff1690565b67ffffffffffffffff6001600160a01b036132ed60206132de8689612b16565b5101516001600160a01b031690565b600061331182612c038767ffffffffffffffff16600052600a602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a301613296565b5050565b60208183031261024c5780359067ffffffffffffffff821161024c570181601f8201121561024c5780359061337882610ca4565b9261338660405194856106fe565b8284526020606081860194028301019181831161024c57602001925b8284106133b0575050505090565b60608483031261024c5760206060916040516133cb8161066d565b86356133d68161023b565b81526133e3838801612d51565b838201526133f360408801611385565b60408201528152019301926133a2565b906040516134108161066d565b604060ff8294546001600160a01b0381168452818160a01c16602085015260a81c161515910152565b9061344261416e565b60005b8251811015612efc576134588184612b16565b5160206134686132af8487612b16565b9101519067ffffffffffffffff811680158015613652575b8015613624575b8015613591575b613559579161351f82600195946134cf6134c2612b4e6135249767ffffffffffffffff166000526009602052604060002090565b6001600160e01b03191690565b61352a577f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad60405180613502878261117e565b0390a267ffffffffffffffff166000526009602052604060002090565b6144a3565b01613445565b7f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e360405180613502878261117e565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b506001600160e01b03196135b16101608501516001600160e01b03191690565b16630a04b54b60e21b8114159081613612575b81613600575b816135ee575b816135dc575b5061348e565b63647e2ba960e01b91501415386135d6565b63c4e0595360e01b81141591506135d0565b632b1dfffb60e21b81141591506135ca565b6307842f7160e21b81141591506135c4565b506101e083015163ffffffff1663ffffffff61364a611a00606087015163ffffffff1690565b911611613487565b5063ffffffff61366a6101e085015163ffffffff1690565b1615613480565b61367961416e565b60208101519160005b8351811015613706578061369b610ac760019387612b16565b6136bd6136b86001600160a01b0383165b6001600160a01b031690565b615f46565b6136c9575b5001613682565b6040516001600160a01b039190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a1386136c2565b5091505160005b815181101561334057613723610ac78284612b16565b906001600160a01b03821615613799577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613790836137756137706136ac6001976001600160a01b031690565b615ecd565b506040516001600160a01b0390911681529081906020820190565b0390a10161370d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b909291928360041161024c57831161024c57600401916003190190565b9060041161024c5790600490565b919091356001600160e01b03198116926004811061380a575050565b6001600160e01b0319929350829060040360031b1b161690565b9080601f8301121561024c57813561383b81610ca4565b9261384960405194856106fe565b81845260208085019260051b82010192831161024c57602001905b8282106138715750505090565b8135815260209182019101613864565b60208183031261024c5780359067ffffffffffffffff821161024c570160a08183031261024c57604051916138b58361068e565b6138be82611385565b83526138cc602083016102cb565b602084015260408201356138df81610cbc565b604084015260608201356060840152608082013567ffffffffffffffff811161024c5761390c9201613824565b608082015290565b908060041161024c577fe0c4c5460000000000000000000000000000000000000000000000000000000082356001600160e01b031916016139835761395f81613967926060946137c3565b810190613881565b015160408051602080820193909352918252610b5890826106fe565b50506040516139936020826106fe565b6000815290565b6139a261416e565b60005b815181101561334057806001600160a01b036139c360019385612b16565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d613a6067ffffffffffffffff60206139ff8689612b16565b51015116836000526008602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a2016139a5565b919082039182116128b557565b613a7e612f01565b50613aa4613a9f826001600160a01b03166000526006602052604060002090565b612f1a565b6020810191613ac3613abd611a00855163ffffffff1690565b42613a69565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001611613b6b5761195f613b0f916001600160a01b03166000526007602052604060002090565b613b1f6119746040830151151590565b8015613b71575b613b6b57613b3390614b3b565b9163ffffffff613b5b611a00613b50602087015163ffffffff1690565b935163ffffffff1690565b911610613b66575090565b905090565b50905090565b506001600160a01b03613b8b82516001600160a01b031690565b1615613b26565b90600282018092116128b557565b90602082018092116128b557565b90600182018092116128b557565b919082018092116128b557565b9061ffff8091169116029061ffff82169182036128b557565b63ffffffff60209116019063ffffffff82116128b557565b9063ffffffff8091169116019063ffffffff82116128b557565b90613c1d61416e565b60005b8251811015612efc5780613c3660019285612b16565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf613d5b60206001600160a01b038451169301518360005260076020526040600020613cbb6001600160a01b0383511682906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b602082015181547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000075ff0000000000000000000000000000000000000000006040870151151560a81b169360a01b169116171790556040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b0390a201613c20565b613d6d81613a76565b9063ffffffff602083015116158015613dca575b613d935750516001600160e01b031690565b6001600160a01b03907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b506001600160e01b0382511615613d81565b6001600160e01b0319169190630a04b54b60e21b8314613e75576307842f7160e21b8314613e6757632b1dfffb60e21b8314613e5c5763647e2ba960e01b8314613e515763c4e0595360e01b8314613e435782632ee8207560e01b60005260045260246000fd5b61025c919250600b9061575a565b61025c9192506157bb565b61025c9192506156f7565b61025c91925060019061575a565b61025c919250615678565b916001600160e01b03198316630a04b54b60e21b8114613f42576307842f7160e21b8114613f2257632b1dfffb60e21b8114613f165763647e2ba960e01b8114613f0a5763c4e0595360e01b14613ef057632ee8207560e01b6000526001600160e01b0319831660045260246000fd5b61025c925015613f0257600b9061575a565b60009061575a565b505061025c91506157bb565b505061025c91506156f7565b5061025c925015613f395760ff60015b169061575a565b60ff6000613f32565b505061025c9150615678565b33600052600360205260406000205415613f6457565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b611374613fb99196949395929667ffffffffffffffff166000526009602052604060002090565b946101608601946001600160e01b0319613fdb87516001600160e01b03191690565b16630a04b54b60e21b811490811561415d575b811561414c575b506141075763c4e0595360e01b6001600160e01b031961401d88516001600160e01b03191690565b16146140e95750506307842f7160e21b6001600160e01b031961404886516001600160e01b03191690565b161461408057611b1e61406385516001600160e01b03191690565b632ee8207560e01b6000526001600160e01b031916600452602490565b6140d493506110a360606140be8763ffffffff6140b56101806140ad866140e29b9d015163ffffffff1690565b930151151590565b911685876159b8565b0151604051958691602083019190602083019252565b03601f1981018652856106fe565b9160019190565b94509450610b58916140fc913691612a75565b936001933691612a75565b9450949161412d91614127611a006101e0610b5896015163ffffffff1690565b91615837565b93614144602061413c87615940565b960151151590565b933691612a75565b63647e2ba960e01b91501438613ff5565b632b1dfffb60e21b81149150613fee565b6001600160a01b0360015416330361418257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b815181546020808501516040808701517fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000090941663ffffffff958616179190921b67ffffffff00000000161791901b69ffff000000000000000016178255606083015161025c936142bc9260a09261425e911685547fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff1660509190911b6dffffffff0000000000000000000016178555565b6142b5614272608083015163ffffffff1690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911b71ffffffff000000000000000000000000000016178555565b0151151590565b81547fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1690151560901b72ff00000000000000000000000000000000000016179055565b9192909261431082828686615afd565b600052600460205260ff604060002054161561432c5750505050565b6040517f097e17ff0000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529390921660248401527fffffffffffffffffffff0000000000000000000000000000000000000000000090911660448301527fffff000000000000000000000000000000000000000000000000000000000000166064820152608490fd5b0390fd5b604d81116128b557600a0a90565b60ff1660120160ff81116128b55760ff1690602482111561444f5760231982019182116128b55761440561440b926143c4565b90612933565b6001600160e01b038111614425576001600160e01b031690565b7f10cb51d10000000000000000000000000000000000000000000000000000000060005260046000fd5b9060240390602482116128b557612369614468926143c4565b61440b565b9060ff80911691160160ff81116128b55760ff1690602482111561444f5760231982019182116128b55761440561440b926143c4565b90614a87610240600161025c946144d06144bd8651151590565b829060ff60ff1983541691151516179055565b6145166144e2602087015161ffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff1660089190911b62ffff0016178255565b61456261452a604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff1660189190911b66ffffffff00000016178255565b6145b2614576606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffff1660389190911b6affffffff0000000000000016178255565b6146066145c6608087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff1660589190911b6effffffff000000000000000000000016178255565b61465861461760a087015160ff1690565b82547fffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffff1660789190911b6fff00000000000000000000000000000016178255565b6146ab61466960c087015160ff1690565b82547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1660809190911b70ff0000000000000000000000000000000016178255565b6147016146bd60e087015161ffff1690565b82547fffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffff1660889190911b72ffff000000000000000000000000000000000016178255565b61475e61471661010087015163ffffffff1690565b82547fffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffff1660989190911b76ffffffff0000000000000000000000000000000000000016178255565b6147bb61477161012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b61481a6147ce61014087015161ffff1690565b82547fffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffff1660c89190911b7affff0000000000000000000000000000000000000000000000000016178255565b6148836148336101608701516001600160e01b03191690565b82547fff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff1660089190911c7effffffff00000000000000000000000000000000000000000000000000000016178255565b6148e4614894610180870151151590565b82547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178255565b01926149286148f96101a083015161ffff1690565b859061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b61497461493d6101c083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178555565b6149c46149896101e083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1660309190911b69ffffffff00000000000016178555565b614a206149dd61020083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff1660509190911b71ffffffffffffffff0000000000000000000016178555565b614a7c614a3561022083015163ffffffff1690565b85547fffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffff1660909190911b75ffffffff00000000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff0000000000000000000000000000000000000000000083549260b01b169116179055565b519069ffffffffffffffffffff8216820361024c57565b908160a091031261024c57614afc81614ad1565b91602082015191604081015191610b58608060608401519301614ad1565b6040513d6000823e3d90fd5b9081602091031261024c5751610b5881611c89565b614b43612f01565b50614b5b6136ac6136ac83516001600160a01b031690565b90604051907ffeaf968c00000000000000000000000000000000000000000000000000000000825260a082600481865afa928315614c6357600092600094614c68575b5060008312614425576020600491604051928380927f313ce5670000000000000000000000000000000000000000000000000000000082525afa928315614c6357610b589363ffffffff93614c0493600092614c2d575b506020015160ff165b9061446d565b92614c1f614c10610721565b6001600160e01b039095168552565b1663ffffffff166020830152565b614bfe919250614c54602091823d8411614c5c575b614c4c81836106fe565b810190614b26565b929150614bf5565b503d614c42565b614b1a565b909350614c8e91925060a03d60a011614c9b575b614c8681836106fe565b810190614ae8565b5093925050919238614b9e565b503d614c7c565b9081602091031261024c573590565b9190614cc06020830183612a42565b93905060408301614cd18185612d1b565b90506040840191614ce9611a00845163ffffffff1690565b80881161529c5750602085015161ffff168083116152665750610160850196614d1a88516001600160e01b03191690565b6001600160e01b03198116630a04b54b60e21b81148015615256575b8015615246575b15614db857505050505050509181614db26110a3614d99614dab96614d686080610b58980186612a42565b6101e083015163ffffffff169063ffffffff614d9161018061413c606088015163ffffffff1690565b941692615d02565b51958694516001600160e01b03191690565b9280612a42565b90613e80565b63c4e0595360e01b819b9a939495979996989b14600014615016575050614e3a614e19614e2d999a614ded6080880188612a42565b63ffffffff614e11610180614e09606087015163ffffffff1690565b950151151590565b931691615c4b565b918251998a91516001600160e01b03191690565b614db26110a38880612a42565b6060810151519082614e57614e4f8780612a42565b810190614ca2565b614ff9575081614fc3575b8515159081614fb6575b50614f8c5760408111614f5a5750614e9190614e8b85949397956128e8565b90613bbc565b946000935b838510614eef575050505050611a00614eb3915163ffffffff1690565b808211614ebf57505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614f2f611a00612cac614f1e8667ffffffffffffffff16600052600a602052604060002090565b6110d7612b8c8d6109108b8d612d1b565b8015614f4a57614f3e91613bbc565b965b0193929190614e96565b50614f5490613ba0565b96614f40565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614e6c565b611b1e827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b615010919350614e8b61500b84613bae565b6128ba565b91614e62565b6307842f7160e21b036152285750615083615060614e2d999a61503c6080880188612a42565b63ffffffff615058610180614e09606087015163ffffffff1690565b9316916159b8565b91615072611a00845163ffffffff1690565b998a91516001600160e01b03191690565b6080810151519082615098614e4f8780612a42565b6152105750816151da575b851515806151ce575b614f8c576040821161519a576020015167ffffffffffffffff9081169081831c166151605750506150e490614e8b85949397956128d0565b946000935b838510615106575050505050611a00614eb3915163ffffffff1690565b9091929395600190615135611a00612cac614f1e8667ffffffffffffffff16600052600a602052604060002090565b80156151505761514491613bbc565b965b01939291906150e9565b5061515a90613ba0565b96615146565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b506060810151156150ac565b611b1e827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b615222919350614e8b61500b84613b92565b916150a3565b632ee8207560e01b6000526001600160e01b03191660045260246000fd5b5063647e2ba960e01b8114614d3d565b50632b1dfffb60e21b8114614d36565b7fd88dddd600000000000000000000000000000000000000000000000000000000600052600483905261ffff1660245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff81166000526005602052604060002091604051926152f4846106aa565b546001600160e01b038116845260e01c9182602085015263ffffffff8216928361532e575b50505050610b5890516001600160e01b031690565b63ffffffff1642908103939084116128b557831161534c5780615319565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b60408136031261024c576020604051916153a8836106aa565b80356153b38161023b565b83520135602082015290565b9694919695929390956000946000986000986000965b8088106153e9575050505050505050929190565b9091929394959697999a6154066154018a848b6129e3565b61538f565b9a61543d6110ee8d6110d761542f8967ffffffffffffffff16600052600a602052604060002090565b91516001600160a01b031690565b9161544e61197460a0850151151590565b6155975760009c60408401906154696124c8835161ffff1690565b61551f575b5050606083015163ffffffff1661548491613bfa565b9c60808301516154979063ffffffff1690565b6154a091613bfa565b9b82516154b09063ffffffff1690565b63ffffffff166154bf90612893565b600193908083106155135750612581611a0060206154e293015163ffffffff1690565b80821161550257506154f391613bbc565b985b01969594939291906153d5565b905061550d91613bbc565b986154f5565b91505061550d91613bbc565b90612369615588939f61557661557f9460208f8e6124c895506001600160a01b0361555185516001600160a01b031690565b91166001600160a01b038216146155905761556c9150613d64565b915b015190615d40565b925161ffff1690565b620186a0900490565b9b388061546e565b509161556e565b999b50600191506155b8846155b26155be93614e8b8b612893565b9b613bfa565b9c613be2565b9a6154f5565b91939093806101e00193846101e0116128b55761012081029080820461012014901517156128b5576101e09101018093116128b5576124c861014061565a610b58966dffffffffffffffffffffffffffff6124556156456156326156649a63ffffffff6123699a1690613bbc565b6123696124c86101208c015161ffff1690565b614e8b611a006101008b015163ffffffff1690565b93015161ffff1690565b612900565b9081602091031261024c575190565b60208151036156ae576156946020825183010160208301615669565b6001600160a01b0381119081156156eb575b506156ae5750565b6143c0906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b610400915010386156a6565b602081510361571d57600b6157156020835184010160208401615669565b1061571d5750565b6143c0906040519182917fe0d7fb020000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b906020825103615780578061576d575050565b6157156020835184010160208401615669565b6040517fe0d7fb0200000000000000000000000000000000000000000000000000000000815260206004820152806143c06024820185610311565b60248151036157d1576022810151156157d15750565b6143c0906040519182917f373b0e440000000000000000000000000000000000000000000000000000000083526020600484018181520190610311565b9081604091031261024c57602060405191615828836106aa565b805183520151612d9481610cbc565b91615840612f01565b50811561591e57506158696110a382806158636001600160e01b031995876137ee565b956137c3565b91167f181dcf100000000000000000000000000000000000000000000000000000000081036158a6575080602080610b589351830101910161580e565b7f97a657c900000000000000000000000000000000000000000000000000000000146158f6577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060208061590993518301019101615669565b615911610721565b9081526000602082015290565b91505067ffffffffffffffff615932610721565b911681526000602082015290565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610b586064826106fe565b604051906159978261068e565b60606080836000815260006020820152600060408201526000838201520152565b6159c061598a565b508115615ad3577f1f3b3aba000000000000000000000000000000000000000000000000000000006001600160e01b0319615a046159fe85856137e0565b906137ee565b1603615aa95781615a189261395f926137c3565b9180615a93575b615a695763ffffffff615a36835163ffffffff1690565b1611615a3f5790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fee433e990000000000000000000000000000000000000000000000000000000060005260046000fd5b50615aa46119746040840151151590565b615a1f565b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b604080516001600160a01b039283166020820190815292909316908301527fffffffffffffffffffff0000000000000000000000000000000000000000000090921660608201527fffff000000000000000000000000000000000000000000000000000000000000909216608083015290615b7b8160a0810161112c565b51902090565b6001600160a01b03610b589116600b615dd4565b6001600160a01b03610b589116600b615f08565b60405190615bb6826106e2565b606080836000815260006020820152600060408201520152565b60208183031261024c5780359067ffffffffffffffff821161024c570160808183031261024c5760405191615c04836106e2565b813583526020820135615c1681610cbc565b602084015260408201356040840152606082013567ffffffffffffffff811161024c57615c439201613824565b606082015290565b615c53615ba9565b508115615ad3577f21ea4ca9000000000000000000000000000000000000000000000000000000006001600160e01b0319615c916159fe85856137e0565b1603615aa95781615cad92615ca5926137c3565b810190615bd0565b9180615cec575b615a6957815111615cc25790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50615cfd6119746020840151151590565b615cb4565b9063ffffffff615d1f93959495615d17612f01565b501691615837565b91825111615cc25780615d34575b615a695790565b50602081015115615d2d565b670de0b6b3a7640000916001600160e01b03615d5c9216612920565b0490565b80548210156129f35760005260206000200190600090565b91615d92918354906000199060031b92831b921b19161790565b9055565b80548015615dbe576000190190615dad8282615d60565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615e865760001984018481116128b55783549360001985019485116128b5576000958583615e3797615e289503615e3d575b505050615d96565b90600052602052604060002090565b55600190565b615e6d615e6791615e5e615e54615e7d9588615d60565b90549060031b1c90565b92839187615d60565b90615d78565b8590600052602052604060002090565b55388080615e20565b50505050600090565b805490680100000000000000008210156106895781615eb6916001615d9294018155615d60565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054615f0257615eeb816002615e8f565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615f3f5780615f2a83600193615e8f565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615f3f576000198201908282116128b5576002549260001984019384116128b5578383615e379460009603615fa6575b505050615f956002615d96565b600390600052602052604060002090565b615f95615e6791615fbe615e54615fc8956002615d60565b9283916002615d60565b55388080615f8856fea164736f6c634300081a000a",
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
