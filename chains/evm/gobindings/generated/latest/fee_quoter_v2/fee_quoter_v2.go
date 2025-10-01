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

type KeystoneFeedsPermissionHandlerPermission struct {
	Forwarder     common.Address
	WorkflowName  [10]byte
	ReportName    [2]byte
	WorkflowOwner common.Address
	IsAllowed     bool
}

var FeeQuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_BASE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPriceFeedConfig\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveTokenReceiver\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setReportPermissions\",\"inputs\":[{\"name\":\"permissions\",\"type\":\"tuple[]\",\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenPriceFeeds\",\"inputs\":[{\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"priceFeedConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportPermissionSet\",\"inputs\":[{\"name\":\"reportId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"permission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataFeedValueOutOfUint224Range\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeRange\",\"inputs\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportForwarderUnauthorized\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346110b257617584803803806100198161131e565b92833981019080820361012081126110b2576060136110b25761003a6112e0565b81516001600160601b03811681036110b257815261005a60208301611343565b906020810191825261006e60408401611357565b6040820190815260608401516001600160401b0381116110b2578561009491860161137f565b60808501519094906001600160401b0381116110b257866100b691830161137f565b60a08201519096906001600160401b0381116110b25782019080601f830112156110b25781516100ed6100e882611368565b61131e565b9260208085848152019260071b820101908382116110b257602001915b81831061126b5750505060c08301516001600160401b0381116110b25783019781601f8a0112156110b2578851986101446100e88b611368565b996020808c838152019160051b830101918483116110b25760208101915b838310611109575050505060e08401516001600160401b0381116110b25784019382601f860112156110b257845161019c6100e882611368565b9560208088848152019260061b820101908582116110b257602001915b8183106110cd57505050610100810151906001600160401b0382116110b2570182601f820112156110b2578051906101f36100e883611368565b9360206102808187868152019402830101918183116110b257602001925b828410610ef057505050503315610edf57600180546001600160a01b031916331790556020986102408a61131e565b976000895260003681376102526112ff565b998a52888b8b015260005b89518110156102c4576001906001600160a01b0361027b828d611418565b51168d61028782611604565b610294575b50500161025d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388d61028c565b508a985089519660005b885181101561033f576001600160a01b036102e9828b611418565b511690811561032e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8c8361032060019561158c565b50604051908152a1016102ce565b6342bcdf7f60e11b60005260046000fd5b5081518a985089906001600160a01b0316158015610ecd575b8015610ebe575b610ead5791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c0526103928661131e565b9360008552600036813760005b855181101561040e576001906103c76001600160a01b036103c0838a611418565b5116611499565b6103d2575b0161039f565b818060a01b036103e28289611418565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26103cc565b508694508560005b84518110156104855760019061043e6001600160a01b036104378389611418565b51166115cb565b610449575b01610416565b818060a01b036104598288611418565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610443565b508593508460005b835181101561054757806104a360019286611418565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf606089858060a01b038451169301518360005260078b5260406000209060ff878060a01b038251169283898060a01b03198254161781558d8301908151604082549501948460a81b8651151560a81b16918560a01b9060a01b169061ffff60a01b19161717905560405193845251168c8301525115156040820152a20161048d565b5091509160005b8251811015610b45576105618184611418565b51856001600160401b036105758487611418565b5151169101519080158015610b32575b8015610b14575b8015610a92575b610a7e57600081815260098852604090205460019392919060081b6001600160e01b03191661093657807f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806106fc868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a25b60005260098752826040600020825115158382549162ffff008c83015160081b169066ffffffff000000604084015160181b166affffffff00000000000000606085015160381b16926effffffff0000000000000000000000608086015160581b169260ff60781b60a087015160781b169460ff60801b60c088015160801b169161ffff60881b60e089015160881b169063ffffffff60981b6101008a015160981b169361ffff60b81b6101208b015160b81b169661ffff60c81b6101408c015160c81b169963ffffffff60d81b6101608d015160081c169b61018060ff60f81b910151151560f81b169c8f8060f81b039a63ffffffff60d81b199961ffff60c81b199861ffff60b81b199763ffffffff60981b199661ffff60881b199560ff60801b199460ff60781b19936effffffff0000000000000000000000199260ff6affffffff000000000000001992169066ffffffffffffff19161716171617161716171617161716171617161716179063ffffffff60d81b1617178155019061ffff6101a0820151169082549165ffffffff00006101c083015160101b169269ffffffff0000000000006101e084015160301b166a01000000000000000000008860901b0361020085015160501b169263ffffffff60901b61022086015160901b169461024063ffffffff60b01b91015160b01b169563ffffffff60b01b199363ffffffff60901b19926a01000000000000000000008c60901b0319918c8060501b03191617161716171617171790550161054e565b807f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e360405180610a76868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a2610700565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61016083015116630a04b54b60e21b8114159081610b02575b81610af0575b81610ade575b81610acc575b50610593565b63647e2ba960e01b1415905088610ac6565b63c4e0595360e01b8114159150610ac0565b632b1dfffb60e21b8114159150610aba565b6307842f7160e21b8114159150610ab4565b5063ffffffff6101e08301511663ffffffff6060840151161061058c565b5063ffffffff6101e08301511615610585565b84828560005b8151811015610bcb576001906001600160a01b03610b698285611418565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610b9f8689611418565b510151168360005260088252604060002081878060401b0319825416179055604051908152a201610b4b565b83600184610bd88361131e565b9060008252600092610ea8575b909282935b8251851015610de757610bfd8584611418565b5180516001600160401b0316939083019190855b83518051821015610dd657610c27828792611418565b51015184516001600160a01b0390610c40908490611418565b5151169063ffffffff815116908781019163ffffffff8351169081811015610dc15750506080810163ffffffff815116898110610daa575090899291838c52600a8a5260408c20600160a01b6001900386168d528a5260408c2092825163ffffffff169380549180518d1b67ffffffff0000000016916040860192835160401b69ffff000000000000000016966060810195865160501b6dffffffff00000000000000000000169063ffffffff60701b895160701b169260a001998b60ff60901b8c51151560901b169560ff60901b199363ffffffff60701b19926dffffffff000000000000000000001991600160501b60019003191617161716171617171790556040519586525163ffffffff168c8601525161ffff1660408501525163ffffffff1660608401525163ffffffff16608083015251151560a082015260c07f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b591a3600101610c11565b6312766e0160e11b8c52600485905260245260448bfd5b6305a7b3d160e11b8c5260045260245260448afd5b505060019096019593509050610bea565b9150825b8251811015610e69576001906001600160401b03610e098286611418565b515116828060a01b0384610e1d8488611418565b5101511690808752600a855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610deb565b604051615eeb908161169982396080518181816105b20152610c28015260a0518181816105e80152610bd9015260c05181818161060f01526139c00152f35b610be5565b63d794ef9560e01b60005260046000fd5b5063ffffffff8251161561035f565b5080516001600160601b031615610358565b639b15e16f60e01b60005260046000fd5b83820361028081126110b257610260610f076112ff565b91610f11876113f5565b8352601f1901126110b2576040519161026083016001600160401b038111848210176110b757604052610f46602087016113e8565b8352610f5460408701611409565b6020840152610f6560608701611357565b6040840152610f7660808701611357565b6060840152610f8760a08701611357565b6080840152610f9860c087016113da565b60a0840152610fa960e087016113da565b60c0840152610fbb6101008701611409565b60e0840152610fcd6101208701611357565b610100840152610fe06101408701611409565b610120840152610ff36101608701611409565b610140840152610180860151916001600160e01b0319831683036110b2578360209361016061028096015261102b6101a089016113e8565b61018082015261103e6101c08901611409565b6101a08201526110516101e08901611357565b6101c08201526110646102008901611357565b6101e082015261107761022089016113f5565b61020082015261108a6102408901611357565b61022082015261109d6102608901611357565b61024082015283820152815201930192610211565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040838703126110b25760206040916110e46112ff565b6110ed86611343565b81526110fa8387016113f5565b838201528152019201916101b9565b82516001600160401b0381116110b25782016040818803601f1901126110b2576111316112ff565b9061113e602082016113f5565b825260408101516001600160401b0381116110b257602091010187601f820112156110b25780516111716100e882611368565b91602060e08185858152019302820101908a82116110b257602001915b8183106111ad5750505091816020938480940152815201920191610162565b828b0360e081126110b25760c06111c26112ff565b916111cc86611343565b8352601f1901126110b2576040519160c08301916001600160401b038311848410176110b75760e093602093604052611206848801611357565b815261121460408801611357565b8482015261122460608801611409565b604082015261123560808801611357565b606082015261124660a08801611357565b608082015261125760c088016113e8565b60a08201528382015281520192019161118e565b828403608081126110b25760606112806112ff565b9161128a86611343565b8352601f1901126110b2576080916020916112a36112e0565b6112ae848801611343565b81526112bc604088016113da565b848201526112cc606088016113e8565b60408201528382015281520192019161010a565b60405190606082016001600160401b038111838210176110b757604052565b60408051919082016001600160401b038111838210176110b757604052565b6040519190601f01601f191682016001600160401b038111838210176110b757604052565b51906001600160a01b03821682036110b257565b519063ffffffff821682036110b257565b6001600160401b0381116110b75760051b60200190565b9080601f830112156110b25781516113996100e882611368565b9260208085848152019260051b8201019283116110b257602001905b8282106113c25750505090565b602080916113cf84611343565b8152019101906113b5565b519060ff821682036110b257565b519081151582036110b257565b51906001600160401b03821682036110b257565b519061ffff821682036110b257565b805182101561142c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561142c5760005260206000200190600090565b805480156114835760001901906114718282611442565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600c6020526040902054801561155a57600019810181811161154457600b54600019810191908211611544578181036114f3575b5050506114df600b61145a565b600052600c60205260006040812055600190565b61152c61150461151593600b611442565b90549060031b1c928392600b611442565b819391549060031b91821b91600019901b19161790565b9055600052600c6020526040600020553880806114d2565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156110b7578161151591600161158894018155611442565b9055565b806000526003602052604060002054156000146115c5576115ae816002611561565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146115c5576115ed81600b611561565b600b5490600052600c602052604060002055600190565b600081815260036020526040902054801561155a576000198101818111611544576002546000198101919082116115445780820361165e575b50505061164a600261145a565b600052600360205260006040812055600190565b61168061166f611515936002611442565b90549060031b1c9283926002611442565b9055600052600360205260406000205538808061163d56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be1461022657806301447eaa1461022157806301ffc9a71461021c578063061877e31461021757806306285c6914610212578063181f5a771461020d5780632451a62714610208578063325c868e146102035780633937306f146101fe5780633a49bb49146101f957806341ed29e7146101f457806345ac924d146101ef5780634ab35b0b146101ea578063514e8cff146101e55780636def4ce7146101e0578063770e2dc4146101db57806379ba5097146101d65780637afac322146101d1578063805f2132146101cc57806382b49eb0146101c757806387b8d879146101c25780638da5cb5b146101bd57806391a2749a146101b85780639b1115e4146101b3578063a69c64c0146101ae578063bf78e03f146101a9578063cdc73d51146101a4578063d02641a01461019f578063d63d3af21461019a578063d8694ccd14610195578063f2fde38b14610190578063fbe3f7781461018b5763ffdb4b371461018657600080fd5b61266a565b61256d565b6124b1565b61207e565b612062565b612019565b611fa2565b611efc565b611e43565b611df2565b611d5e565b611d37565b611b1b565b61199e565b611703565b6115ca565b6114b2565b6112b1565b611132565b610f73565b610f3b565b610e72565b610cdd565b610b67565b61088d565b610871565b6107ee565b61074c565b610576565b61052e565b610474565b6103c0565b61024e565b6001600160a01b0381160361023c57565b600080fd5b359061024c8261022b565b565b3461023c57606060031936011261023c5760206102856004356102708161022b565b602435604435916102808361022b565b612800565b604051908152f35b6004359067ffffffffffffffff8216820361023c57565b6024359067ffffffffffffffff8216820361023c57565b359067ffffffffffffffff8216820361023c57565b9181601f8401121561023c5782359167ffffffffffffffff831161023c576020808501948460051b01011161023c57565b919082519283825260005b84811061032d575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161030c565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061037557505050505090565b90919293946020806103b1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610301565b97019301930191939290610366565b3461023c57606060031936011261023c576103d961028d565b60243567ffffffffffffffff811161023c576103f99036906004016102d0565b6044929192359167ffffffffffffffff831161023c573660238401121561023c5782600401359167ffffffffffffffff831161023c573660248460061b8601011161023c5761045b94602461044f9501926129c7565b60405191829182610342565b0390f35b35906001600160e01b03198216820361023c57565b3461023c57602060031936011261023c576004356001600160e01b0319811680910361023c57807f805f21320000000000000000000000000000000000000000000000000000000060209214908115610504575b81156104da575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386104cf565b7f66792e8000000000000000000000000000000000000000000000000000000000811491506104c8565b3461023c57602060031936011261023c576001600160a01b036004356105538161022b565b166000526008602052602067ffffffffffffffff60406000205416604051908152f35b3461023c57600060031936011261023c5761058f612be0565b50606060405161059e8161065d565b63ffffffff6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918281526001600160a01b0360406020830192827f00000000000000000000000000000000000000000000000000000000000000001684520191837f00000000000000000000000000000000000000000000000000000000000000001683526040519485525116602084015251166040820152f35b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761067957604052565b610647565b60a0810190811067ffffffffffffffff82111761067957604052565b6040810190811067ffffffffffffffff82111761067957604052565b60c0810190811067ffffffffffffffff82111761067957604052565b6080810190811067ffffffffffffffff82111761067957604052565b90601f601f19910116810190811067ffffffffffffffff82111761067957604052565b6040519061024c6040836106ee565b6040519061024c610260836106ee565b67ffffffffffffffff811161067957601f01601f191660200190565b3461023c57600060031936011261023c5761045b604080519061076f81836106ee565b601382527f46656551756f74657220312e362e332d64657600000000000000000000000000602083015251918291602083526020830190610301565b602060408183019282815284518094520192019060005b8181106107cf5750505090565b82516001600160a01b03168452602093840193909201916001016107c2565b3461023c57600060031936011261023c5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061085b5761045b8561084f818703826106ee565b604051918291826107ab565b8254845260209093019260019283019201610838565b3461023c57600060031936011261023c57602060405160248152f35b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c5780600401906040600319823603011261023c576108cb613e44565b6108d58280612bff565b4263ffffffff1692915060005b818110610a2e575050602401906108f98284612bff565b92905060005b83811061090857005b8061092761092260019361091c868a612bff565b90612890565b612c80565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff610a0a6109fc60208501946109ee61097087516001600160e01b031690565b61098a61097b610711565b6001600160e01b039092168252565b63ffffffff8c1660208201526109c56109ab845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b815160209092015160e01b6001600160e01b0319166001600160e01b0392909216919091179055565b5167ffffffffffffffff1690565b93516001600160e01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108ff565b80610a47610a4260019361091c8980612bff565b612c49565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03610ae06109fc6020850194610ad3610a8f87516001600160e01b031690565b610a9a61097b610711565b63ffffffff8d1660208201526109c5610aba84516001600160a01b031690565b6001600160a01b03166000526006602052604060002090565b516001600160a01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108e2565b9181601f8401121561023c5782359167ffffffffffffffff831161023c576020838186019501011161023c57565b92610b649492610b5692855215156020850152608060408501526080840190610301565b916060818403910152610301565b90565b3461023c5760a060031936011261023c57610b8061028d565b60243590610b8d8261022b565b6044359160643567ffffffffffffffff811161023c57610bb1903690600401610b04565b93909160843567ffffffffffffffff811161023c57610bd4903690600401610b04565b9290917f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b0382166001600160a01b03821614600014610ca0575050935b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016808611610c6f575091610c60939161045b9693613e88565b90939160405194859485610b32565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91610caa92612800565b93610c19565b67ffffffffffffffff81116106795760051b60200190565b8015150361023c57565b359061024c82610cc8565b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c573660238201121561023c57806004013590610d1a82610cb0565b90610d2860405192836106ee565b828252602460a06020840194028201019036821161023c57602401925b818410610d5757610d5583612ca5565b005b60a08436031261023c5760405190610d6e8261067e565b8435610d798161022b565b825260208501357fffffffffffffffffffff000000000000000000000000000000000000000000008116810361023c5760208301526040850135907fffff0000000000000000000000000000000000000000000000000000000000008216820361023c5782602092604060a0950152610df460608801610241565b6060820152610e0560808801610cd2565b6080820152815201930192610d45565b602060408183019282815284518094520192019060005b818110610e395750505090565b9091926020604082610e67600194885163ffffffff602080926001600160e01b038151168552015116910152565b019401929101610e2c565b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c57610ea39036906004016102d0565b610eac81610cb0565b91610eba60405193846106ee565b818352601f19610ec983610cb0565b0160005b818110610f2457505060005b82811015610f1657600190610efa610ef58260051b85016128a5565b61396c565b610f0482876129b3565b52610f0f81866129b3565b5001610ed9565b6040518061045b8682610e15565b602090610f2f612de5565b82828801015201610ecd565b3461023c57602060031936011261023c576020610f62600435610f5d8161022b565b613c5a565b6001600160e01b0360405191168152f35b3461023c57602060031936011261023c5767ffffffffffffffff610f9561028d565b610f9d612de5565b50166000526005602052604060002060405190610fb98261069a565b546001600160e01b038116825260e01c6020820152604051809161045b82604081019263ffffffff602080926001600160e01b038151168552015116910152565b61024c909291926102408061026083019561101784825115159052565b60208181015161ffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015163ffffffff169085015260a08181015160ff169085015260c08181015160ff169085015260e08181015161ffff16908501526101008181015163ffffffff16908501526101208181015161ffff16908501526101408181015161ffff1690850152610160818101516001600160e01b03191690850152610180818101511515908501526101a08181015161ffff16908501526101c08181015163ffffffff16908501526101e08181015163ffffffff16908501526102008181015167ffffffffffffffff16908501526102208181015163ffffffff1690850152015163ffffffff16910152565b3461023c57602060031936011261023c5761045b6111f56111f061115461028d565b6000610240611161610720565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e08201528261020082015282610220820152015267ffffffffffffffff166000526009602052604060002090565b612e23565b60405191829182610ffa565b359063ffffffff8216820361023c57565b359061ffff8216820361023c57565b81601f8201121561023c5780359061123882610cb0565b9261124660405194856106ee565b82845260208085019360061b8301019181831161023c57602001925b828410611270575050505090565b60408483031261023c576020604091825161128a8161069a565b611293876102bb565b8152828701356112a28161022b565b83820152815201930192611262565b3461023c57604060031936011261023c5760043567ffffffffffffffff811161023c573660238201121561023c5780600401356112ed81610cb0565b916112fb60405193846106ee565b8183526024602084019260051b8201019036821161023c5760248101925b82841061134a576024358567ffffffffffffffff821161023c57611344610d55923690600401611221565b90612f79565b833567ffffffffffffffff811161023c5782016040602319823603011261023c57604051906113788261069a565b611384602482016102bb565b8252604481013567ffffffffffffffff811161023c57602491010136601f8201121561023c5780356113b581610cb0565b916113c360405193846106ee565b818352602060e081850193028201019036821161023c57602001915b8183106113fe5750505091816020938480940152815201930192611319565b82360360e0811261023c5760c0601f196040519261141b8461069a565b86356114268161022b565b8452011261023c5760e091602091604051611440816106b6565b61144b848801611201565b815261145960408801611201565b8482015261146960608801611212565b604082015261147a60808801611201565b606082015261148b60a08801611201565b608082015260c087013561149e81610cc8565b60a0820152838201528152019201916113df565b3461023c57600060031936011261023c576000546001600160a01b0381163303611539577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561023c57813561157a81610cb0565b9261158860405194856106ee565b81845260208085019260051b82010192831161023c57602001905b8282106115b05750505090565b6020809183356115bf8161022b565b8152019101906115a3565b3461023c57604060031936011261023c5760043567ffffffffffffffff811161023c576115fb903690600401611563565b60243567ffffffffffffffff811161023c5761161b903690600401611563565b90611624614064565b60005b81518110156116935780611648611643610ad3600194866129b3565b615a8e565b611653575b01611627565b6001600160a01b03611668610ad383866129b3565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a261164d565b8260005b8151811015610d5557806116b86116b3610ad3600194866129b3565b615aa2565b6116c3575b01611697565b6001600160a01b036116d8610ad383866129b3565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a26116bd565b3461023c57604060031936011261023c5760043567ffffffffffffffff811161023c57611734903690600401610b04565b6024359167ffffffffffffffff831161023c5761178d61178561176b611761611795963690600401610b04565b9490953691612922565b90604082015190605e604a84015160601c93015191929190565b9190336141f6565b810190613228565b60005b8151811015610d55576117e06117db6117c26117b484866129b3565b51516001600160a01b031690565b6001600160a01b03166000526007602052604060002090565b6132e7565b6117f46117f06040830151151590565b1590565b611955579061183f61180c6020600194015160ff1690565b61183961182d602061181e86896129b3565b5101516001600160e01b031690565b6001600160e01b031690565b906142c8565b61185a604061184e84876129b3565b51015163ffffffff1690565b63ffffffff61188561187c611875610aba6117b4888b6129b3565b5460e01c90565b63ffffffff1690565b91161061194f576118d361189e604061184e85886129b3565b6118c36118a9610711565b6001600160e01b03851681529163ffffffff166020830152565b6109c5610aba6117b486896129b3565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b036119096117b485886129b3565b61194561191b604061184e888b6129b3565b60405193849316958390929163ffffffff6020916001600160e01b03604085019616845216910152565b0390a25b01611798565b50611949565b61199a6119656117b484866129b3565b7f06439c6b000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6000fd5b3461023c57604060031936011261023c5761045b611a266119bd61028d565b67ffffffffffffffff602435916119d38361022b565b600060a06040516119e3816106b6565b828152826020820152826040820152826060820152826080820152015216600052600a6020526040600020906001600160a01b0316600052602052604060002090565b611aa2611a9960405192611a39846106b6565b5463ffffffff8116845263ffffffff8160201c16602085015261ffff8160401c166040850152611a80611a738263ffffffff9060501c1690565b63ffffffff166060860152565b63ffffffff607082901c16608085015260901c60ff1690565b151560a0830152565b6040519182918291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b60ff81160361023c57565b359061024c82611b05565b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c573660238201121561023c57806004013590611b5882610cb0565b90611b6660405192836106ee565b82825260246102806020840194028201019036821161023c57602401925b818410611b9457610d558361331d565b833603610280811261023c57610260601f1960405192611bb38461069a565b611bbc886102bb565b8452011261023c5761028091602091611bd3610720565b611bde848901610cd2565b8152611bec60408901611212565b84820152611bfc60608901611201565b6040820152611c0d60808901611201565b6060820152611c1e60a08901611201565b6080820152611c2f60c08901611b10565b60a0820152611c4060e08901611b10565b60c0820152611c526101008901611212565b60e0820152611c646101208901611201565b610100820152611c776101408901611212565b610120820152611c8a6101608901611212565b610140820152611c9d610180890161045f565b610160820152611cb06101a08901610cd2565b610180820152611cc36101c08901611212565b6101a0820152611cd66101e08901611201565b6101c0820152611ce96102008901611201565b6101e0820152611cfc61022089016102bb565b610200820152611d0f6102408901611201565b610220820152611d226102608901611201565b61024082015283820152815201930192611b84565b3461023c57600060031936011261023c5760206001600160a01b0360015416604051908152f35b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c576040600319823603011261023c57604051611d9b8161069a565b816004013567ffffffffffffffff811161023c57611dbf9060043691850101611563565b8152602482013567ffffffffffffffff811161023c57610d55926004611de89236920101611563565b6020820152613555565b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c57611e2f611e2961045b923690600401610b04565b906137f8565b604051918291602083526020830190610301565b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c573660238201121561023c57806004013590611e8082610cb0565b90611e8e60405192836106ee565b8282526024602083019360061b8201019036821161023c57602401925b818410611ebb57610d5583613890565b60408436031261023c5760206040918251611ed58161069a565b8635611ee08161022b565b8152611eed8388016102bb565b83820152815201930192611eab565b3461023c57602060031936011261023c576001600160a01b03600435611f218161022b565b611f29612be0565b5016600052600760205261045b604060002060ff60405191611f4a8361065d565b546001600160a01b0381168352818160a01c16602084015260a81c16151560408201526040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b3461023c57600060031936011261023c57604051806020600b54918281520190600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db99060005b8181106120035761045b8561084f818703826106ee565b8254845260209093019260019283019201611fec565b3461023c57602060031936011261023c57604061203b600435610ef58161022b565b6120608251809263ffffffff602080926001600160e01b038151168552015116910152565bf35b3461023c57600060031936011261023c57602060405160128152f35b3461023c57604060031936011261023c5761209761028d565b60243567ffffffffffffffff811161023c578060040160a0600319833603011261023c576120dc6111f08467ffffffffffffffff166000526009602052604060002090565b916120ea6117f08451151590565b61247957606481016121206117f0612101836128a5565b6001600160a01b03166000526001600b01602052604060002054151590565b612438578390604483016121348186612bff565b95905061214281858a614ba7565b9261214f610f5d826128a5565b97889361216d61216761022084015163ffffffff1690565b8c6151db565b9a6000808b156123ff5750506122056121d261ffff856121f7996121de999896612212966121b66101c06121aa6101a06122189d015161ffff1690565b95015163ffffffff1690565b6121c96121c28b6128a5565b938d612bff565b969095166152cc565b989198979098946128a5565b6001600160a01b03166000526008602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b906127cd565b9560009761ffff61222f61014089015161ffff1690565b166123a1575b509461221261220561020061230761045b9d6dffffffffffffffffffffffffffff6122ff61231f9f9e9b6122fa6001600160e01b039f9b9c6123179f6122fa9e63ffffffff61228e6122fa9f60246122989501906128ef565b9290501690613ab2565b908b60a081016122bb6122b56122af835160ff1690565b60ff1690565b856127cd565b9360e08301916122cd835161ffff1690565b9061ffff8216831161232f575b50505050608001516122fa9161187c9163ffffffff16613af0565b613af0565b613ab2565b9116906127cd565b93015167ffffffffffffffff1690565b9116906127e0565b6040519081529081906020820190565b61187c9496506122fa959361ffff61239061237f6122f59661237961237261236960809960ff6123636123979b5160ff1690565b16613abf565b965161ffff1690565b61ffff1690565b9061395f565b6122126122af60c08d015160ff1690565b9116613ab2565b95938395506122da565b909594989750826123c78b989495986dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff16916123e560248901856128ef565b90506123f193886154d1565b969793943896939296612235565b959492509550506122126122056121f76121de61243261242d61187c61024061221899015163ffffffff1690565b612740565b946128a5565b61244461199a916128a5565b7f2502348c000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b3461023c57602060031936011261023c576001600160a01b036004356124d68161022b565b6124de614064565b1633811461254357807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461023c57602060031936011261023c5760043567ffffffffffffffff811161023c573660238201121561023c578060040135906125aa82610cb0565b906125b860405192836106ee565b8282526024602083019360071b8201019036821161023c57602401925b8184106125e557610d5583613b0a565b8336036080811261023c576060601f19604051926126028461069a565b873561260d8161022b565b8452011261023c576080916020916040516126278161065d565b838801356126348161022b565b8152604088013561264481611b05565b84820152606088013561265681610cc8565b6040820152838201528152019301926125d5565b3461023c57604060031936011261023c576004356126878161022b565b61268f6102a4565b9067ffffffffffffffff82169182600052600960205260ff60406000205416156126fc576126bf6126e092613c5a565b92600052600960205263ffffffff60016040600020015460901c16906151db565b604080516001600160e01b039384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc10000149015171561276257565b61272a565b908160051b918083046020149015171561276257565b9061012c82029180830461012c149015171561276257565b9061010c82029180830461010c149015171561276257565b90655af3107a4000820291808304655af3107a4000149015171561276257565b8181029291811591840414171561276257565b81156127ea570490565b634e487b7160e01b600052601260045260246000fd5b61282a612824610b6494936001600160e01b0361281d8195613c5a565b16906127cd565b92613c5a565b16906127e0565b9061283b82610cb0565b61284860405191826106ee565b828152601f196128588294610cb0565b019060005b82811061286957505050565b80606060208093850101520161285d565b634e487b7160e01b600052603260045260246000fd5b91908110156128a05760061b0190565b61287a565b35610b648161022b565b91908110156128a05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561023c570190565b903590601e198136030182121561023c570180359067ffffffffffffffff821161023c5760200191813603831361023c57565b92919261292e82610730565b9161293c60405193846106ee565b82948184528183011161023c578281602093846000960137010152565b9061024c604051612969816106b6565b925463ffffffff8082168552602082811c821690860152604082811c61ffff1690860152605082901c81166060860152607082901c16608085015260901c60ff16151560a0840152565b80518210156128a05760209160051b010190565b9092916129fc6129eb8367ffffffffffffffff166000526009602052604060002090565b5460081b6001600160e01b03191690565b90612a0681612831565b9560005b828110612a1b575050505050505090565b612a2e612a29828489612890565b6128a5565b8388612a48612a3e8584846128af565b60408101906128ef565b905060208111612b65575b508392612a89612a83612a7c612a72600198612ac497612abf976128af565b60208101906128ef565b3691612922565b89613cd2565b612aa78967ffffffffffffffff16600052600a602052604060002090565b906001600160a01b0316600052602052604060002090565b612959565b60a081015115612b2957612b0d612ae56060612aff93015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b03601f1981018352826106ee565b612b17828b6129b3565b52612b22818a6129b3565b5001612a0a565b50612aff612b0d612b6084612b528a67ffffffffffffffff166000526009602052604060002090565b015460101c63ffffffff1690565b612ae5565b915050612b9d61187c612b9084612aa78b67ffffffffffffffff16600052600a602052604060002090565b5460701c63ffffffff1690565b10612baa57838838612a53565b7f36f536ca000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b60405190612bed8261065d565b60006040838281528260208201520152565b903590601e198136030182121561023c570180359067ffffffffffffffff821161023c57602001918160061b3603831361023c57565b35906001600160e01b038216820361023c57565b60408136031261023c57612c78602060405192612c658461069a565b8035612c708161022b565b845201612c35565b602082015290565b60408136031261023c57612c78602060405192612c9c8461069a565b612c70816102bb565b90612cae614064565b60005b8251811015612de05780612cc7600192856129b3565b517f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a360a06001600160a01b038351169260608101936001600160a01b0380865116957fffff000000000000000000000000000000000000000000000000000000000000612d6560208601947fffffffffffffffffffff00000000000000000000000000000000000000000000865116604088019a848c511692615a0a565b977fffffffffffffffffffff000000000000000000000000000000000000000000006080870195612db3875115158c600052600460205260406000209060ff60ff1983541691151516179055565b8560405198511688525116602087015251166040850152511660608301525115156080820152a201612cb1565b509050565b60405190612df28261069a565b60006020838281520152565b90604051612e0b8161069a565b91546001600160e01b038116835260e01c6020830152565b9061024c612f6b6001612e34610720565b94612f0a612f008254612e50612e4a8260ff1690565b15158a52565b61ffff600882901c1660208a015263ffffffff601882901c1660408a015263ffffffff603882901c1660608a015263ffffffff605882901c1660808a015260ff607882901c1660a08a015260ff608082901c1660c08a015261ffff608882901c1660e08a015263ffffffff609882901c166101008a015261ffff60b882901c166101208a015261ffff60c882901c166101408a01526001600160e01b0319600882901b166101608a015260f81c90565b1515610180880152565b015461ffff81166101a086015263ffffffff601082901c166101c086015263ffffffff603082901c166101e086015267ffffffffffffffff605082901c1661020086015263ffffffff609082901c1661022086015260b01c63ffffffff1690565b63ffffffff16610240840152565b90612f82614064565b6000915b805183101561317457612f9983826129b3565b5190612fad825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b8551805182101561315f57612fd9826020926129b3565b510151612fea6117b48389516129b3565b8151602083015163ffffffff908116911681811015613126575050608082015163ffffffff16602081106130e5575090867f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b56001600160a01b0384613074858f60019998612aa761306f9267ffffffffffffffff16600052600a602052604060002090565b6140a2565b6130dc60405192839216958291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b0390a301612fc2565b7f24ecdc02000000000000000000000000000000000000000000000000000000006000526001600160a01b0390911660045263ffffffff1660245260446000fd5b7f0b4f67a20000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b50509550925092600191500191929092612f86565b50905060005b815181101561322457806131a2613193600193856129b3565b515167ffffffffffffffff1690565b67ffffffffffffffff6001600160a01b036131d160206131c286896129b3565b5101516001600160a01b031690565b60006131f582612aa78767ffffffffffffffff16600052600a602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a30161317a565b5050565b60208183031261023c5780359067ffffffffffffffff821161023c570181601f8201121561023c5780359061325c82610cb0565b9261326a60405194856106ee565b8284526020606081860194028301019181831161023c57602001925b828410613294575050505090565b60608483031261023c5760206060916040516132af8161065d565b86356132ba8161022b565b81526132c7838801612c35565b838201526132d760408801611201565b6040820152815201930192613286565b906040516132f48161065d565b604060ff8294546001600160a01b0381168452818160a01c16602085015260a81c161515910152565b90613326614064565b60005b8251811015612de05761333c81846129b3565b51602061334c61319384876129b3565b9101519067ffffffffffffffff811680158015613536575b8015613508575b8015613475575b61343d579161340382600195946133b36133a66129eb6134089767ffffffffffffffff166000526009602052604060002090565b6001600160e01b03191690565b61340e577f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806133e68782610ffa565b0390a267ffffffffffffffff166000526009602052604060002090565b614399565b01613329565b7f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e3604051806133e68782610ffa565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b506001600160e01b03196134956101608501516001600160e01b03191690565b16630a04b54b60e21b81141590816134f6575b816134e4575b816134d2575b816134c0575b50613372565b63647e2ba960e01b91501415386134ba565b63c4e0595360e01b81141591506134b4565b632b1dfffb60e21b81141591506134ae565b6307842f7160e21b81141591506134a8565b506101e083015163ffffffff1663ffffffff61352e61187c606087015163ffffffff1690565b91161161336b565b5063ffffffff61354e6101e085015163ffffffff1690565b1615613364565b61355d614064565b60208101519160005b83518110156135ea578061357f610ad3600193876129b3565b6135a161359c6001600160a01b0383165b6001600160a01b031690565b615e53565b6135ad575b5001613566565b6040516001600160a01b039190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a1386135a6565b5091505160005b815181101561322457613607610ad382846129b3565b906001600160a01b0382161561367d577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613674836136596136546135906001976001600160a01b031690565b615dda565b506040516001600160a01b0390911681529081906020820190565b0390a1016135f1565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b909291928360041161023c57831161023c57600401916003190190565b9060041161023c5790600490565b919091356001600160e01b0319811692600481106136ee575050565b6001600160e01b0319929350829060040360031b1b161690565b9080601f8301121561023c57813561371f81610cb0565b9261372d60405194856106ee565b81845260208085019260051b82010192831161023c57602001905b8282106137555750505090565b8135815260209182019101613748565b60208183031261023c5780359067ffffffffffffffff821161023c570160a08183031261023c57604051916137998361067e565b6137a282611201565b83526137b0602083016102bb565b602084015260408201356137c381610cc8565b604084015260608201356060840152608082013567ffffffffffffffff811161023c576137f09201613708565b608082015290565b90600481108015613853575b61383c57606061382661381e83610b6494612aff966136a7565b810190613765565b0151604051928391602083019190602083019252565b505060405161384c6020826106ee565b6000815290565b508060041161023c5781356001600160e01b0319167f1f3b3aba000000000000000000000000000000000000000000000000000000001415613804565b613898614064565b60005b815181101561322457806001600160a01b036138b9600193856129b3565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d61395667ffffffffffffffff60206138f586896129b3565b51015116836000526008602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a20161389b565b9190820391821161276257565b613974612de5565b5061399a613995826001600160a01b03166000526006602052604060002090565b612dfe565b60208101916139b96139b361187c855163ffffffff1690565b4261395f565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001611613a61576117db613a05916001600160a01b03166000526007602052604060002090565b613a156117f06040830151151590565b8015613a67575b613a6157613a2990614a31565b9163ffffffff613a5161187c613a46602087015163ffffffff1690565b935163ffffffff1690565b911610613a5c575090565b905090565b50905090565b506001600160a01b03613a8182516001600160a01b031690565b1615613a1c565b906002820180921161276257565b906020820180921161276257565b906001820180921161276257565b9190820180921161276257565b9061ffff8091169116029061ffff821691820361276257565b63ffffffff60209116019063ffffffff821161276257565b9063ffffffff8091169116019063ffffffff821161276257565b90613b13614064565b60005b8251811015612de05780613b2c600192856129b3565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf613c5160206001600160a01b038451169301518360005260076020526040600020613bb16001600160a01b0383511682906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b602082015181547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000075ff0000000000000000000000000000000000000000006040870151151560a81b169360a01b169116171790556040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b0390a201613b16565b613c638161396c565b9063ffffffff602083015116158015613cc0575b613c895750516001600160e01b031690565b6001600160a01b03907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b506001600160e01b0382511615613c77565b6001600160e01b0319169190630a04b54b60e21b8314613d6b576307842f7160e21b8314613d5d57632b1dfffb60e21b8314613d525763647e2ba960e01b8314613d475763c4e0595360e01b8314613d395782632ee8207560e01b60005260045260246000fd5b61024c919250600b90615667565b61024c9192506156c8565b61024c919250615604565b61024c919250600190615667565b61024c919250615585565b916001600160e01b03198316630a04b54b60e21b8114613e38576307842f7160e21b8114613e1857632b1dfffb60e21b8114613e0c5763647e2ba960e01b8114613e005763c4e0595360e01b14613de657632ee8207560e01b6000526001600160e01b0319831660045260246000fd5b61024c925015613df857600b90615667565b600090615667565b505061024c91506156c8565b505061024c9150615604565b5061024c925015613e2f5760ff60015b1690615667565b60ff6000613e28565b505061024c9150615585565b33600052600360205260406000205415613e5a57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6111f0613eaf9196949395929667ffffffffffffffff166000526009602052604060002090565b946101608601946001600160e01b0319613ed187516001600160e01b03191690565b16630a04b54b60e21b8114908115614053575b8115614042575b50613ffd5763c4e0595360e01b6001600160e01b0319613f1388516001600160e01b03191690565b1614613fdf5750506307842f7160e21b6001600160e01b0319613f3e86516001600160e01b03191690565b1614613f765761199a613f5985516001600160e01b03191690565b632ee8207560e01b6000526001600160e01b031916600452602490565b613fca9350612a7c6060613fb48763ffffffff613fab610180613fa386613fd89b9d015163ffffffff1690565b930151151590565b911685876158c5565b0151604051958691602083019190602083019252565b03601f1981018652856106ee565b9160019190565b94509450610b6491613ff2913691612922565b936001933691612922565b945094916140239161401d61187c6101e0610b6496015163ffffffff1690565b91615744565b9361403a60206140328761584d565b960151151590565b933691612922565b63647e2ba960e01b91501438613eeb565b632b1dfffb60e21b81149150613ee4565b6001600160a01b0360015416330361407857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b815181546020808501516040808701517fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000090941663ffffffff958616179190921b67ffffffff00000000161791901b69ffff000000000000000016178255606083015161024c936141b29260a092614154911685547fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff1660509190911b6dffffffff0000000000000000000016178555565b6141ab614168608083015163ffffffff1690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911b71ffffffff000000000000000000000000000016178555565b0151151590565b81547fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1690151560901b72ff00000000000000000000000000000000000016179055565b9192909261420682828686615a0a565b600052600460205260ff60406000205416156142225750505050565b6040517f097e17ff0000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529390921660248401527fffffffffffffffffffff0000000000000000000000000000000000000000000090911660448301527fffff000000000000000000000000000000000000000000000000000000000000166064820152608490fd5b0390fd5b604d811161276257600a0a90565b60ff1660120160ff81116127625760ff16906024821115614345576023198201918211612762576142fb614301926142ba565b906127e0565b6001600160e01b03811161431b576001600160e01b031690565b7f10cb51d10000000000000000000000000000000000000000000000000000000060005260046000fd5b9060240390602482116127625761221261435e926142ba565b614301565b9060ff80911691160160ff81116127625760ff16906024821115614345576023198201918211612762576142fb614301926142ba565b9061497d610240600161024c946143c66143b38651151590565b829060ff60ff1983541691151516179055565b61440c6143d8602087015161ffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff1660089190911b62ffff0016178255565b614458614420604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff1660189190911b66ffffffff00000016178255565b6144a861446c606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffff1660389190911b6affffffff0000000000000016178255565b6144fc6144bc608087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff1660589190911b6effffffff000000000000000000000016178255565b61454e61450d60a087015160ff1690565b82547fffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffff1660789190911b6fff00000000000000000000000000000016178255565b6145a161455f60c087015160ff1690565b82547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1660809190911b70ff0000000000000000000000000000000016178255565b6145f76145b360e087015161ffff1690565b82547fffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffff1660889190911b72ffff000000000000000000000000000000000016178255565b61465461460c61010087015163ffffffff1690565b82547fffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffff1660989190911b76ffffffff0000000000000000000000000000000000000016178255565b6146b161466761012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b6147106146c461014087015161ffff1690565b82547fffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffff1660c89190911b7affff0000000000000000000000000000000000000000000000000016178255565b6147796147296101608701516001600160e01b03191690565b82547fff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff1660089190911c7effffffff00000000000000000000000000000000000000000000000000000016178255565b6147da61478a610180870151151590565b82547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178255565b019261481e6147ef6101a083015161ffff1690565b859061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b61486a6148336101c083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178555565b6148ba61487f6101e083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1660309190911b69ffffffff00000000000016178555565b6149166148d361020083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff1660509190911b71ffffffffffffffff0000000000000000000016178555565b61497261492b61022083015163ffffffff1690565b85547fffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffff1660909190911b75ffffffff00000000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff0000000000000000000000000000000000000000000083549260b01b169116179055565b519069ffffffffffffffffffff8216820361023c57565b908160a091031261023c576149f2816149c7565b91602082015191604081015191610b646080606084015193016149c7565b6040513d6000823e3d90fd5b9081602091031261023c5751610b6481611b05565b614a39612de5565b50614a5161359061359083516001600160a01b031690565b90604051907ffeaf968c00000000000000000000000000000000000000000000000000000000825260a082600481865afa928315614b5957600092600094614b5e575b506000831261431b576020600491604051928380927f313ce5670000000000000000000000000000000000000000000000000000000082525afa928315614b5957610b649363ffffffff93614afa93600092614b23575b506020015160ff165b90614363565b92614b15614b06610711565b6001600160e01b039095168552565b1663ffffffff166020830152565b614af4919250614b4a602091823d8411614b52575b614b4281836106ee565b810190614a1c565b929150614aeb565b503d614b38565b614a10565b909350614b8491925060a03d60a011614b91575b614b7c81836106ee565b8101906149de565b5093925050919238614a94565b503d614b72565b9081602091031261023c573590565b9190614bb660208301836128ef565b93905060408301614bc78185612bff565b90506040840191614bdf61187c845163ffffffff1690565b8088116151a95750602085015161ffff168083116151735750610160850196614c1088516001600160e01b03191690565b6001600160e01b03198116630a04b54b60e21b81148015615163575b8015615153575b15614cae57505050505050509181614ca8612a7c614c8f614ca196614c5e6080610b649801866128ef565b6101e083015163ffffffff169063ffffffff614c87610180614032606088015163ffffffff1690565b941692615c0f565b51958694516001600160e01b03191690565b92806128ef565b90613d76565b63c4e0595360e01b819b9a939495979996989b14600014614f23575050614d30614d0f614d23999a614ce360808801886128ef565b63ffffffff614d07610180614cff606087015163ffffffff1690565b950151151590565b931691615b58565b918251998a91516001600160e01b03191690565b614ca8612a7c88806128ef565b6060810151519082614d4d614d4587806128ef565b810190614b98565b614f06575081614ed0575b8515159081614ec3575b50614e995760408111614e675750614d8790614d818594939795612795565b90613ab2565b946000935b838510614de557505050505061187c614da9915163ffffffff1690565b808211614db557505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614e3c61187c612b90614e148667ffffffffffffffff16600052600a602052604060002090565b614e25612a298d61091c8b8d612bff565b6001600160a01b0316600052602052604060002090565b8015614e5757614e4b91613ab2565b965b0193929190614d8c565b50614e6190613a96565b96614e4d565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614d62565b61199a827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614f1d919350614d81614f1884613aa4565b612767565b91614d58565b6307842f7160e21b036151355750614f90614f6d614d23999a614f4960808801886128ef565b63ffffffff614f65610180614cff606087015163ffffffff1690565b9316916158c5565b91614f7f61187c845163ffffffff1690565b998a91516001600160e01b03191690565b6080810151519082614fa5614d4587806128ef565b61511d5750816150e7575b851515806150db575b614e9957604082116150a7576020015167ffffffffffffffff9081169081831c1661506d575050614ff190614d81859493979561277d565b946000935b83851061501357505050505061187c614da9915163ffffffff1690565b909192939560019061504261187c612b90614e148667ffffffffffffffff16600052600a602052604060002090565b801561505d5761505191613ab2565b965b0193929190614ff6565b5061506790613a96565b96615053565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614fb9565b61199a827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b61512f919350614d81614f1884613a88565b91614fb0565b632ee8207560e01b6000526001600160e01b03191660045260246000fd5b5063647e2ba960e01b8114614c33565b50632b1dfffb60e21b8114614c2c565b7fd88dddd600000000000000000000000000000000000000000000000000000000600052600483905261ffff1660245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff81166000526005602052604060002091604051926152018461069a565b546001600160e01b038116845260e01c9182602085015263ffffffff8216928361523b575b50505050610b6490516001600160e01b031690565b63ffffffff1642908103939084116127625783116152595780615226565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b60408136031261023c576020604051916152b58361069a565b80356152c08161022b565b83520135602082015290565b9694919695929390956000946000986000986000965b8088106152f6575050505050505050929190565b9091929394959697999a61531361530e8a848b612890565b61529c565b9a61534a612abf8d614e2561533c8967ffffffffffffffff16600052600a602052604060002090565b91516001600160a01b031690565b9161535b6117f060a0850151151590565b6154a45760009c6040840190615376612372835161ffff1690565b61542c575b5050606083015163ffffffff1661539191613af0565b9c60808301516153a49063ffffffff1690565b6153ad91613af0565b9b82516153bd9063ffffffff1690565b63ffffffff166153cc90612740565b60019390808310615420575061242d61187c60206153ef93015163ffffffff1690565b80821161540f575061540091613ab2565b985b01969594939291906152e2565b905061541a91613ab2565b98615402565b91505061541a91613ab2565b90612212615495939f61548361548c9460208f8e61237295506001600160a01b0361545e85516001600160a01b031690565b91166001600160a01b0382161461549d576154799150613c5a565b915b015190615c4d565b925161ffff1690565b620186a0900490565b9b388061537b565b509161547b565b999b50600191506154c5846154bf6154cb93614d818b612740565b9b613af0565b9c613ad8565b9a615402565b91939093806101e00193846101e011612762576101208102908082046101201490151715612762576101e091010180931161276257612372610140615567610b64966dffffffffffffffffffffffffffff6122ff61555261553f6155719a63ffffffff6122129a1690613ab2565b6122126123726101208c015161ffff1690565b614d8161187c6101008b015163ffffffff1690565b93015161ffff1690565b6127ad565b9081602091031261023c575190565b60208151036155bb576155a16020825183010160208301615576565b6001600160a01b0381119081156155f8575b506155bb5750565b6142b6906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526020600484018181520190610301565b610400915010386155b3565b602081510361562a57600b6156226020835184010160208401615576565b1061562a5750565b6142b6906040519182917fe0d7fb020000000000000000000000000000000000000000000000000000000083526020600484018181520190610301565b90602082510361568d578061567a575050565b6156226020835184010160208401615576565b6040517fe0d7fb0200000000000000000000000000000000000000000000000000000000815260206004820152806142b66024820185610301565b60248151036156de576022810151156156de5750565b6142b6906040519182917f373b0e440000000000000000000000000000000000000000000000000000000083526020600484018181520190610301565b9081604091031261023c576020604051916157358361069a565b805183520151612c7881610cc8565b9161574d612de5565b50811561582b5750615776612a7c82806157706001600160e01b031995876136d2565b956136a7565b91167f181dcf100000000000000000000000000000000000000000000000000000000081036157b3575080602080610b649351830101910161571b565b7f97a657c90000000000000000000000000000000000000000000000000000000014615803577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060208061581693518301019101615576565b61581e610711565b9081526000602082015290565b91505067ffffffffffffffff61583f610711565b911681526000602082015290565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610b646064826106ee565b604051906158a48261067e565b60606080836000815260006020820152600060408201526000838201520152565b6158cd615897565b5081156159e0577f1f3b3aba000000000000000000000000000000000000000000000000000000006001600160e01b031961591161590b85856136c4565b906136d2565b16036159b657816159259261381e926136a7565b91806159a0575b6159765763ffffffff615943835163ffffffff1690565b161161594c5790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fee433e990000000000000000000000000000000000000000000000000000000060005260046000fd5b506159b16117f06040840151151590565b61592c565b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b604080516001600160a01b039283166020820190815292909316908301527fffffffffffffffffffff0000000000000000000000000000000000000000000090921660608201527fffff000000000000000000000000000000000000000000000000000000000000909216608083015290615a888160a08101612aff565b51902090565b6001600160a01b03610b649116600b615ce1565b6001600160a01b03610b649116600b615e15565b60405190615ac3826106d2565b606080836000815260006020820152600060408201520152565b60208183031261023c5780359067ffffffffffffffff821161023c570160808183031261023c5760405191615b11836106d2565b813583526020820135615b2381610cc8565b602084015260408201356040840152606082013567ffffffffffffffff811161023c57615b509201613708565b606082015290565b615b60615ab6565b5081156159e0577f21ea4ca9000000000000000000000000000000000000000000000000000000006001600160e01b0319615b9e61590b85856136c4565b16036159b65781615bba92615bb2926136a7565b810190615add565b9180615bf9575b61597657815111615bcf5790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50615c0a6117f06020840151151590565b615bc1565b9063ffffffff615c2c93959495615c24612de5565b501691615744565b91825111615bcf5780615c41575b6159765790565b50602081015115615c3a565b670de0b6b3a7640000916001600160e01b03615c6992166127cd565b0490565b80548210156128a05760005260206000200190600090565b91615c9f918354906000199060031b92831b921b19161790565b9055565b80548015615ccb576000190190615cba8282615c6d565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615d93576000198401848111612762578354936000198501948511612762576000958583615d4497615d359503615d4a575b505050615ca3565b90600052602052604060002090565b55600190565b615d7a615d7491615d6b615d61615d8a9588615c6d565b90549060031b1c90565b92839187615c6d565b90615c85565b8590600052602052604060002090565b55388080615d2d565b50505050600090565b805490680100000000000000008210156106795781615dc3916001615c9f94018155615c6d565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054615e0f57615df8816002615d9c565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615e4c5780615e3783600193615d9c565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615e4c5760001982019082821161276257600254926000198401938411612762578383615d449460009603615eb3575b505050615ea26002615ca3565b600390600052602052604060002090565b615ea2615d7491615ecb615d61615ed5956002615c6d565b9283916002615c6d565b55388080615e9556fea164736f6c634300081a000a",
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
