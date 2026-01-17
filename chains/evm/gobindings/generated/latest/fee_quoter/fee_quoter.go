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

var FeeQuoterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_BASE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"inputs\":[{\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPriceFeedConfig\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setReportPermissions\",\"inputs\":[{\"name\":\"permissions\",\"type\":\"tuple[]\",\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenPriceFeeds\",\"inputs\":[{\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedConfig\",\"type\":\"tuple\",\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteHigh\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"destGasPerPayloadByteThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"priceFeedConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"components\":[{\"name\":\"dataFeedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportPermissionSet\",\"inputs\":[{\"name\":\"reportId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"permission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"components\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"deciBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataFeedValueOutOfUint224Range\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeRange\",\"inputs\":[{\"name\":\"minFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportForwarderUnauthorized\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"reportName\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346110b257617487803803806100198161131e565b92833981019080820361012081126110b2576060136110b25761003a6112e0565b81516001600160601b03811681036110b257815261005a60208301611343565b906020810191825261006e60408401611357565b6040820190815260608401516001600160401b0381116110b2578561009491860161137f565b60808501519094906001600160401b0381116110b257866100b691830161137f565b60a08201519096906001600160401b0381116110b25782019080601f830112156110b25781516100ed6100e882611368565b61131e565b9260208085848152019260071b820101908382116110b257602001915b81831061126b5750505060c08301516001600160401b0381116110b25783019781601f8a0112156110b2578851986101446100e88b611368565b996020808c838152019160051b830101918483116110b25760208101915b838310611109575050505060e08401516001600160401b0381116110b25784019382601f860112156110b257845161019c6100e882611368565b9560208088848152019260061b820101908582116110b257602001915b8183106110cd57505050610100810151906001600160401b0382116110b2570182601f820112156110b2578051906101f36100e883611368565b9360206102808187868152019402830101918183116110b257602001925b828410610ef057505050503315610edf57600180546001600160a01b031916331790556020986102408a61131e565b976000895260003681376102526112ff565b998a52888b8b015260005b89518110156102c4576001906001600160a01b0361027b828d611418565b51168d61028782611604565b610294575b50500161025d565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388d61028c565b508a985089519660005b885181101561033f576001600160a01b036102e9828b611418565b511690811561032e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8c8361032060019561158c565b50604051908152a1016102ce565b6342bcdf7f60e11b60005260046000fd5b5081518a985089906001600160a01b0316158015610ecd575b8015610ebe575b610ead5791516001600160a01b031660a05290516001600160601b03166080525163ffffffff1660c0526103928661131e565b9360008552600036813760005b855181101561040e576001906103c76001600160a01b036103c0838a611418565b5116611499565b6103d2575b0161039f565b818060a01b036103e28289611418565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a26103cc565b508694508560005b84518110156104855760019061043e6001600160a01b036104378389611418565b51166115cb565b610449575b01610416565b818060a01b036104598288611418565b51167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2610443565b508593508460005b835181101561054757806104a360019286611418565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf606089858060a01b038451169301518360005260078b5260406000209060ff878060a01b038251169283898060a01b03198254161781558d8301908151604082549501948460a81b8651151560a81b16918560a01b9060a01b169061ffff60a01b19161717905560405193845251168c8301525115156040820152a20161048d565b5091509160005b8251811015610b45576105618184611418565b51856001600160401b036105758487611418565b5151169101519080158015610b32575b8015610b14575b8015610a92575b610a7e57600081815260098852604090205460019392919060081b6001600160e01b03191661093657807f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806106fc868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a25b60005260098752826040600020825115158382549162ffff008c83015160081b169066ffffffff000000604084015160181b166affffffff00000000000000606085015160381b16926effffffff0000000000000000000000608086015160581b169260ff60781b60a087015160781b169460ff60801b60c088015160801b169161ffff60881b60e089015160881b169063ffffffff60981b6101008a015160981b169361ffff60b81b6101208b015160b81b169661ffff60c81b6101408c015160c81b169963ffffffff60d81b6101608d015160081c169b61018060ff60f81b910151151560f81b169c8f8060f81b039a63ffffffff60d81b199961ffff60c81b199861ffff60b81b199763ffffffff60981b199661ffff60881b199560ff60801b199460ff60781b19936effffffff0000000000000000000000199260ff6affffffff000000000000001992169066ffffffffffffff19161716171617161716171617161716171617161716179063ffffffff60d81b1617178155019061ffff6101a0820151169082549165ffffffff00006101c083015160101b169269ffffffff0000000000006101e084015160301b166a01000000000000000000008860901b0361020085015160501b169263ffffffff60901b61022086015160901b169461024063ffffffff60b01b91015160b01b169563ffffffff60b01b199363ffffffff60901b19926a01000000000000000000008c60901b0319918c8060501b03191617161716171617171790550161054e565b807f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e360405180610a76868291909161024063ffffffff8161026084019580511515855261ffff602082015116602086015282604082015116604086015282606082015116606086015282608082015116608086015260ff60a08201511660a086015260ff60c08201511660c086015261ffff60e08201511660e0860152826101008201511661010086015261ffff6101208201511661012086015261ffff610140820151166101408601528260e01b61016082015116610160860152610180810151151561018086015261ffff6101a0820151166101a0860152826101c0820151166101c0860152826101e0820151166101e086015260018060401b03610200820151166102008601528261022082015116610220860152015116910152565b0390a2610700565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b61016083015116630a04b54b60e21b8114159081610b02575b81610af0575b81610ade575b81610acc575b50610593565b63647e2ba960e01b1415905088610ac6565b63c4e0595360e01b8114159150610ac0565b632b1dfffb60e21b8114159150610aba565b6307842f7160e21b8114159150610ab4565b5063ffffffff6101e08301511663ffffffff6060840151161061058c565b5063ffffffff6101e08301511615610585565b84828560005b8151811015610bcb576001906001600160a01b03610b698285611418565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d86848060401b0381610b9f8689611418565b510151168360005260088252604060002081878060401b0319825416179055604051908152a201610b4b565b83600184610bd88361131e565b9060008252600092610ea8575b909282935b8251851015610de757610bfd8584611418565b5180516001600160401b0316939083019190855b83518051821015610dd657610c27828792611418565b51015184516001600160a01b0390610c40908490611418565b5151169063ffffffff815116908781019163ffffffff8351169081811015610dc15750506080810163ffffffff815116898110610daa575090899291838c52600a8a5260408c20600160a01b6001900386168d528a5260408c2092825163ffffffff169380549180518d1b67ffffffff0000000016916040860192835160401b69ffff000000000000000016966060810195865160501b6dffffffff00000000000000000000169063ffffffff60701b895160701b169260a001998b60ff60901b8c51151560901b169560ff60901b199363ffffffff60701b19926dffffffff000000000000000000001991600160501b60019003191617161716171617171790556040519586525163ffffffff168c8601525161ffff1660408501525163ffffffff1660608401525163ffffffff16608083015251151560a082015260c07f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b591a3600101610c11565b6312766e0160e11b8c52600485905260245260448bfd5b6305a7b3d160e11b8c5260045260245260448afd5b505060019096019593509050610bea565b9150825b8251811015610e69576001906001600160401b03610e098286611418565b515116828060a01b0384610e1d8488611418565b5101511690808752600a855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610deb565b604051615dee908161169982396080518181816105a20152610bfc015260a0518181816105d80152610bad015260c0518181816105ff015261376a0152f35b610be5565b63d794ef9560e01b60005260046000fd5b5063ffffffff8251161561035f565b5080516001600160601b031615610358565b639b15e16f60e01b60005260046000fd5b83820361028081126110b257610260610f076112ff565b91610f11876113f5565b8352601f1901126110b2576040519161026083016001600160401b038111848210176110b757604052610f46602087016113e8565b8352610f5460408701611409565b6020840152610f6560608701611357565b6040840152610f7660808701611357565b6060840152610f8760a08701611357565b6080840152610f9860c087016113da565b60a0840152610fa960e087016113da565b60c0840152610fbb6101008701611409565b60e0840152610fcd6101208701611357565b610100840152610fe06101408701611409565b610120840152610ff36101608701611409565b610140840152610180860151916001600160e01b0319831683036110b2578360209361016061028096015261102b6101a089016113e8565b61018082015261103e6101c08901611409565b6101a08201526110516101e08901611357565b6101c08201526110646102008901611357565b6101e082015261107761022089016113f5565b61020082015261108a6102408901611357565b61022082015261109d6102608901611357565b61024082015283820152815201930192610211565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040838703126110b25760206040916110e46112ff565b6110ed86611343565b81526110fa8387016113f5565b838201528152019201916101b9565b82516001600160401b0381116110b25782016040818803601f1901126110b2576111316112ff565b9061113e602082016113f5565b825260408101516001600160401b0381116110b257602091010187601f820112156110b25780516111716100e882611368565b91602060e08185858152019302820101908a82116110b257602001915b8183106111ad5750505091816020938480940152815201920191610162565b828b0360e081126110b25760c06111c26112ff565b916111cc86611343565b8352601f1901126110b2576040519160c08301916001600160401b038311848410176110b75760e093602093604052611206848801611357565b815261121460408801611357565b8482015261122460608801611409565b604082015261123560808801611357565b606082015261124660a08801611357565b608082015261125760c088016113e8565b60a08201528382015281520192019161118e565b828403608081126110b25760606112806112ff565b9161128a86611343565b8352601f1901126110b2576080916020916112a36112e0565b6112ae848801611343565b81526112bc604088016113da565b848201526112cc606088016113e8565b60408201528382015281520192019161010a565b60405190606082016001600160401b038111838210176110b757604052565b60408051919082016001600160401b038111838210176110b757604052565b6040519190601f01601f191682016001600160401b038111838210176110b757604052565b51906001600160a01b03821682036110b257565b519063ffffffff821682036110b257565b6001600160401b0381116110b75760051b60200190565b9080601f830112156110b25781516113996100e882611368565b9260208085848152019260051b8201019283116110b257602001905b8282106113c25750505090565b602080916113cf84611343565b8152019101906113b5565b519060ff821682036110b257565b519081151582036110b257565b51906001600160401b03821682036110b257565b519061ffff821682036110b257565b805182101561142c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561142c5760005260206000200190600090565b805480156114835760001901906114718282611442565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600c6020526040902054801561155a57600019810181811161154457600b54600019810191908211611544578181036114f3575b5050506114df600b61145a565b600052600c60205260006040812055600190565b61152c61150461151593600b611442565b90549060031b1c928392600b611442565b819391549060031b91821b91600019901b19161790565b9055600052600c6020526040600020553880806114d2565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156110b7578161151591600161158894018155611442565b9055565b806000526003602052604060002054156000146115c5576115ae816002611561565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146115c5576115ed81600b611561565b600b5490600052600c602052604060002055600190565b600081815260036020526040902054801561155a576000198101818111611544576002546000198101919082116115445780820361165e575b50505061164a600261145a565b600052600360205260006040812055600190565b61168061166f611515936002611442565b90549060031b1c9283926002611442565b9055600052600360205260406000205538808061163d56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be1461021657806301447eaa1461021157806301ffc9a71461020c578063061877e31461020757806306285c6914610202578063181f5a77146101fd5780632451a627146101f8578063325c868e146101f35780633937306f146101ee5780633a49bb49146101e957806341ed29e7146101e457806345ac924d146101df5780634ab35b0b146101da578063514e8cff146101d55780636def4ce7146101d0578063770e2dc4146101cb57806379ba5097146101c65780637afac322146101c1578063805f2132146101bc57806382b49eb0146101b757806387b8d879146101b25780638da5cb5b146101ad57806391a2749a146101a8578063a69c64c0146101a3578063bf78e03f1461019e578063cdc73d5114610199578063d02641a014610194578063d63d3af21461018f578063d8694ccd1461018a578063f2fde38b14610185578063fbe3f778146101805763ffdb4b371461017b57600080fd5b6125ed565b6124f0565b612434565b612001565b611fe5565b611f9c565b611f25565b611e7f565b611dc6565b611d32565b611d0b565b611aef565b611972565b6116d7565b61159e565b611486565b611285565b611106565b610f47565b610f0f565b610e46565b610cb1565b610b3b565b610861565b610845565b6107c2565b610720565b610566565b61051e565b610464565b6103b0565b61023e565b6001600160a01b0381160361022c57565b600080fd5b359061023c8261021b565b565b3461022c57606060031936011261022c5760206102756004356102608161021b565b602435604435916102708361021b565b612783565b604051908152f35b6004359067ffffffffffffffff8216820361022c57565b6024359067ffffffffffffffff8216820361022c57565b359067ffffffffffffffff8216820361022c57565b9181601f8401121561022c5782359167ffffffffffffffff831161022c576020808501948460051b01011161022c57565b919082519283825260005b84811061031d575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016102fc565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061036557505050505090565b90919293946020806103a1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102f1565b97019301930191939290610356565b3461022c57606060031936011261022c576103c961027d565b60243567ffffffffffffffff811161022c576103e99036906004016102c0565b6044929192359167ffffffffffffffff831161022c573660238401121561022c5782600401359167ffffffffffffffff831161022c573660248460061b8601011161022c5761044b94602461043f95019261295a565b60405191829182610332565b0390f35b35906001600160e01b03198216820361022c57565b3461022c57602060031936011261022c576004356001600160e01b0319811680910361022c57807f805f213200000000000000000000000000000000000000000000000000000000602092149081156104f4575b81156104ca575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386104bf565b7f66792e8000000000000000000000000000000000000000000000000000000000811491506104b8565b3461022c57602060031936011261022c576001600160a01b036004356105438161021b565b166000526008602052602067ffffffffffffffff60406000205416604051908152f35b3461022c57600060031936011261022c5761057f612b73565b50606060405161058e8161064d565b63ffffffff6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918281526001600160a01b0360406020830192827f00000000000000000000000000000000000000000000000000000000000000001684520191837f00000000000000000000000000000000000000000000000000000000000000001683526040519485525116602084015251166040820152f35b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761066957604052565b610637565b60a0810190811067ffffffffffffffff82111761066957604052565b6040810190811067ffffffffffffffff82111761066957604052565b60c0810190811067ffffffffffffffff82111761066957604052565b6080810190811067ffffffffffffffff82111761066957604052565b90601f601f19910116810190811067ffffffffffffffff82111761066957604052565b6040519061023c6040836106de565b6040519061023c610260836106de565b3461022c57600060031936011261022c5761044b604080519061074381836106de565b601382527f46656551756f74657220312e362e332d646576000000000000000000000000006020830152519182916020835260208301906102f1565b602060408183019282815284518094520192019060005b8181106107a35750505090565b82516001600160a01b0316845260209384019390920191600101610796565b3461022c57600060031936011261022c5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061082f5761044b85610823818703826106de565b6040519182918261077f565b825484526020909301926001928301920161080c565b3461022c57600060031936011261022c57602060405160248152f35b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c5780600401906040600319823603011261022c5761089f613bee565b6108a98280612b92565b4263ffffffff1692915060005b818110610a02575050602401906108cd8284612b92565b92905060005b8381106108dc57005b806108fb6108f66001936108f0868a612b92565b90612813565b612c13565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff6109de6109d060208501946109c261094487516001600160e01b031690565b61095e61094f610701565b6001600160e01b039092168252565b63ffffffff8c16602082015261099961097f845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b815160209092015160e01b6001600160e01b0319166001600160e01b0392909216919091179055565b5167ffffffffffffffff1690565b93516001600160e01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108d3565b80610a1b610a166001936108f08980612b92565b612bdc565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b03610ab46109d06020850194610aa7610a6387516001600160e01b031690565b610a6e61094f610701565b63ffffffff8d166020820152610999610a8e84516001600160a01b031690565b6001600160a01b03166000526006602052604060002090565b516001600160a01b031690565b604080516001600160e01b039290921682524260208301529190931692a2016108b6565b9181601f8401121561022c5782359167ffffffffffffffff831161022c576020838186019501011161022c57565b92610b389492610b2a928552151560208501526080604085015260808401906102f1565b9160608184039101526102f1565b90565b3461022c5760a060031936011261022c57610b5461027d565b60243590610b618261021b565b6044359160643567ffffffffffffffff811161022c57610b85903690600401610ad8565b93909160843567ffffffffffffffff811161022c57610ba8903690600401610ad8565b9290917f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b0382166001600160a01b03821614600014610c74575050935b6bffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016808611610c43575091610c34939161044b9693613c32565b90939160405194859485610b06565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91610c7e92612783565b93610bed565b67ffffffffffffffff81116106695760051b60200190565b8015150361022c57565b359061023c82610c9c565b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c573660238201121561022c57806004013590610cee82610c84565b90610cfc60405192836106de565b828252602460a06020840194028201019036821161022c57602401925b818410610d2b57610d2983612c38565b005b60a08436031261022c5760405190610d428261066e565b8435610d4d8161021b565b825260208501357fffffffffffffffffffff000000000000000000000000000000000000000000008116810361022c5760208301526040850135907fffff0000000000000000000000000000000000000000000000000000000000008216820361022c5782602092604060a0950152610dc860608801610231565b6060820152610dd960808801610ca6565b6080820152815201930192610d19565b602060408183019282815284518094520192019060005b818110610e0d5750505090565b9091926020604082610e3b600194885163ffffffff602080926001600160e01b038151168552015116910152565b019401929101610e00565b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c57610e779036906004016102c0565b610e8081610c84565b91610e8e60405193846106de565b818352601f19610e9d83610c84565b0160005b818110610ef857505060005b82811015610eea57600190610ece610ec98260051b8501612828565b613716565b610ed88287612946565b52610ee38186612946565b5001610ead565b6040518061044b8682610de9565b602090610f03612d78565b82828801015201610ea1565b3461022c57602060031936011261022c576020610f36600435610f318161021b565b613a04565b6001600160e01b0360405191168152f35b3461022c57602060031936011261022c5767ffffffffffffffff610f6961027d565b610f71612d78565b50166000526005602052604060002060405190610f8d8261068a565b546001600160e01b038116825260e01c6020820152604051809161044b82604081019263ffffffff602080926001600160e01b038151168552015116910152565b61023c9092919261024080610260830195610feb84825115159052565b60208181015161ffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015163ffffffff169085015260a08181015160ff169085015260c08181015160ff169085015260e08181015161ffff16908501526101008181015163ffffffff16908501526101208181015161ffff16908501526101408181015161ffff1690850152610160818101516001600160e01b03191690850152610180818101511515908501526101a08181015161ffff16908501526101c08181015163ffffffff16908501526101e08181015163ffffffff16908501526102008181015167ffffffffffffffff16908501526102208181015163ffffffff1690850152015163ffffffff16910152565b3461022c57602060031936011261022c5761044b6111c96111c461112861027d565b6000610240611135610710565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015282610180820152826101a0820152826101c0820152826101e08201528261020082015282610220820152015267ffffffffffffffff166000526009602052604060002090565b612db6565b60405191829182610fce565b359063ffffffff8216820361022c57565b359061ffff8216820361022c57565b81601f8201121561022c5780359061120c82610c84565b9261121a60405194856106de565b82845260208085019360061b8301019181831161022c57602001925b828410611244575050505090565b60408483031261022c576020604091825161125e8161068a565b611267876102ab565b8152828701356112768161021b565b83820152815201930192611236565b3461022c57604060031936011261022c5760043567ffffffffffffffff811161022c573660238201121561022c5780600401356112c181610c84565b916112cf60405193846106de565b8183526024602084019260051b8201019036821161022c5760248101925b82841061131e576024358567ffffffffffffffff821161022c57611318610d299236906004016111f5565b90612f0c565b833567ffffffffffffffff811161022c5782016040602319823603011261022c576040519061134c8261068a565b611358602482016102ab565b8252604481013567ffffffffffffffff811161022c57602491010136601f8201121561022c57803561138981610c84565b9161139760405193846106de565b818352602060e081850193028201019036821161022c57602001915b8183106113d257505050918160209384809401528152019301926112ed565b82360360e0811261022c5760c0601f19604051926113ef8461068a565b86356113fa8161021b565b8452011261022c5760e091602091604051611414816106a6565b61141f8488016111d5565b815261142d604088016111d5565b8482015261143d606088016111e6565b604082015261144e608088016111d5565b606082015261145f60a088016111d5565b608082015260c087013561147281610c9c565b60a0820152838201528152019201916113b3565b3461022c57600060031936011261022c576000546001600160a01b038116330361150d577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561022c57813561154e81610c84565b9261155c60405194856106de565b81845260208085019260051b82010192831161022c57602001905b8282106115845750505090565b6020809183356115938161021b565b815201910190611577565b3461022c57604060031936011261022c5760043567ffffffffffffffff811161022c576115cf903690600401611537565b60243567ffffffffffffffff811161022c576115ef903690600401611537565b906115f8613e0e565b60005b8151811015611667578061161c611617610aa760019486612946565b615991565b611627575b016115fb565b6001600160a01b0361163c610aa78386612946565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2611621565b8260005b8151811015610d29578061168c611687610aa760019486612946565b6159a5565b611697575b0161166b565b6001600160a01b036116ac610aa78386612946565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611691565b3461022c57604060031936011261022c5760043567ffffffffffffffff811161022c57611708903690600401610ad8565b6024359167ffffffffffffffff831161022c5761176161175961173f611735611769963690600401610ad8565b94909536916128a5565b90604082015190605e604a84015160601c93015191929190565b919033613fa0565b8101906131bb565b60005b8151811015610d29576117b46117af6117966117888486612946565b51516001600160a01b031690565b6001600160a01b03166000526007602052604060002090565b61327a565b6117c86117c46040830151151590565b1590565b61192957906118136117e06020600194015160ff1690565b61180d61180160206117f28689612946565b5101516001600160e01b031690565b6001600160e01b031690565b90614072565b61182e60406118228487612946565b51015163ffffffff1690565b63ffffffff611859611850611849610a8e611788888b612946565b5460e01c90565b63ffffffff1690565b911610611923576118a761187260406118228588612946565b61189761187d610701565b6001600160e01b03851681529163ffffffff166020830152565b610999610a8e6117888689612946565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a6001600160a01b036118dd6117888588612946565b6119196118ef6040611822888b612946565b60405193849316958390929163ffffffff6020916001600160e01b03604085019616845216910152565b0390a25b0161176c565b5061191d565b61196e6119396117888486612946565b7f06439c6b000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6000fd5b3461022c57604060031936011261022c5761044b6119fa61199161027d565b67ffffffffffffffff602435916119a78361021b565b600060a06040516119b7816106a6565b828152826020820152826040820152826060820152826080820152015216600052600a6020526040600020906001600160a01b0316600052602052604060002090565b611a76611a6d60405192611a0d846106a6565b5463ffffffff8116845263ffffffff8160201c16602085015261ffff8160401c166040850152611a54611a478263ffffffff9060501c1690565b63ffffffff166060860152565b63ffffffff607082901c16608085015260901c60ff1690565b151560a0830152565b6040519182918291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b60ff81160361022c57565b359061023c82611ad9565b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c573660238201121561022c57806004013590611b2c82610c84565b90611b3a60405192836106de565b82825260246102806020840194028201019036821161022c57602401925b818410611b6857610d29836132b0565b833603610280811261022c57610260601f1960405192611b878461068a565b611b90886102ab565b8452011261022c5761028091602091611ba7610710565b611bb2848901610ca6565b8152611bc0604089016111e6565b84820152611bd0606089016111d5565b6040820152611be1608089016111d5565b6060820152611bf260a089016111d5565b6080820152611c0360c08901611ae4565b60a0820152611c1460e08901611ae4565b60c0820152611c2661010089016111e6565b60e0820152611c3861012089016111d5565b610100820152611c4b61014089016111e6565b610120820152611c5e61016089016111e6565b610140820152611c71610180890161044f565b610160820152611c846101a08901610ca6565b610180820152611c976101c089016111e6565b6101a0820152611caa6101e089016111d5565b6101c0820152611cbd61020089016111d5565b6101e0820152611cd061022089016102ab565b610200820152611ce361024089016111d5565b610220820152611cf661026089016111d5565b61024082015283820152815201930192611b58565b3461022c57600060031936011261022c5760206001600160a01b0360015416604051908152f35b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c576040600319823603011261022c57604051611d6f8161068a565b816004013567ffffffffffffffff811161022c57611d939060043691850101611537565b8152602482013567ffffffffffffffff811161022c57610d29926004611dbc9236920101611537565b60208201526134e8565b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c573660238201121561022c57806004013590611e0382610c84565b90611e1160405192836106de565b8282526024602083019360061b8201019036821161022c57602401925b818410611e3e57610d298361363a565b60408436031261022c5760206040918251611e588161068a565b8635611e638161021b565b8152611e708388016102ab565b83820152815201930192611e2e565b3461022c57602060031936011261022c576001600160a01b03600435611ea48161021b565b611eac612b73565b5016600052600760205261044b604060002060ff60405191611ecd8361064d565b546001600160a01b0381168352818160a01c16602084015260a81c16151560408201526040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b3461022c57600060031936011261022c57604051806020600b54918281520190600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db99060005b818110611f865761044b85610823818703826106de565b8254845260209093019260019283019201611f6f565b3461022c57602060031936011261022c576040611fbe600435610ec98161021b565b611fe38251809263ffffffff602080926001600160e01b038151168552015116910152565bf35b3461022c57600060031936011261022c57602060405160128152f35b3461022c57604060031936011261022c5761201a61027d565b60243567ffffffffffffffff811161022c578060040160a0600319833603011261022c5761205f6111c48467ffffffffffffffff166000526009602052604060002090565b9161206d6117c48451151590565b6123fc57606481016120a36117c461208483612828565b6001600160a01b03166000526001600b01602052604060002054151590565b6123bb578390604483016120b78186612b92565b9590506120c581858a614951565b926120d2610f3182612828565b9788936120f06120ea61022084015163ffffffff1690565b8c614f85565b9a6000808b1561238257505061218861215561ffff8561217a99612161999896612195966121396101c061212d6101a061219b9d015161ffff1690565b95015163ffffffff1690565b61214c6121458b612828565b938d612b92565b96909516615076565b98919897909894612828565b6001600160a01b03166000526008602052604060002090565b5467ffffffffffffffff1690565b67ffffffffffffffff1690565b90612750565b9560009761ffff6121b261014089015161ffff1690565b16612324575b509461219561218861020061228a61044b9d6dffffffffffffffffffffffffffff6122826122a29f9e9b61227d6001600160e01b039f9b9c61229a9f61227d9e63ffffffff61221161227d9f602461221b950190612872565b929050169061385c565b908b60a0810161223e612238612232835160ff1690565b60ff1690565b85612750565b9360e0830191612250835161ffff1690565b9061ffff821683116122b2575b505050506080015161227d916118509163ffffffff1661389a565b61389a565b61385c565b911690612750565b93015167ffffffffffffffff1690565b911690612763565b6040519081529081906020820190565b61185094965061227d959361ffff612313612302612278966122fc6122f56122ec60809960ff6122e661231a9b5160ff1690565b16613869565b965161ffff1690565b61ffff1690565b90613709565b61219561223260c08d015160ff1690565b911661385c565b959383955061225d565b9095949897508261234a8b989495986dffffffffffffffffffffffffffff9060701c1690565b6dffffffffffffffffffffffffffff16916123686024890185612872565b9050612374938861527b565b9697939438969392966121b8565b9594925095505061219561218861217a6121616123b56123b061185061024061219b99015163ffffffff1690565b6126c3565b94612828565b6123c761196e91612828565b7f2502348c000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b3461022c57602060031936011261022c576001600160a01b036004356124598161021b565b612461613e0e565b163381146124c657807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461022c57602060031936011261022c5760043567ffffffffffffffff811161022c573660238201121561022c5780600401359061252d82610c84565b9061253b60405192836106de565b8282526024602083019360071b8201019036821161022c57602401925b81841061256857610d29836138b4565b8336036080811261022c576060601f19604051926125858461068a565b87356125908161021b565b8452011261022c576080916020916040516125aa8161064d565b838801356125b78161021b565b815260408801356125c781611ad9565b8482015260608801356125d981610c9c565b604082015283820152815201930192612558565b3461022c57604060031936011261022c5760043561260a8161021b565b612612610294565b9067ffffffffffffffff82169182600052600960205260ff604060002054161561267f5761264261266392613a04565b92600052600960205263ffffffff60016040600020015460901c1690614f85565b604080516001600160e01b039384168152919092166020820152f35b827f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc1000014901517156126e557565b6126ad565b908160051b91808304602014901517156126e557565b9061012c82029180830461012c14901517156126e557565b9061010c82029180830461010c14901517156126e557565b90655af3107a4000820291808304655af3107a400014901517156126e557565b818102929181159184041417156126e557565b811561276d570490565b634e487b7160e01b600052601260045260246000fd5b6127ad6127a7610b3894936001600160e01b036127a08195613a04565b1690612750565b92613a04565b1690612763565b906127be82610c84565b6127cb60405191826106de565b828152601f196127db8294610c84565b019060005b8281106127ec57505050565b8060606020809385010152016127e0565b634e487b7160e01b600052603260045260246000fd5b91908110156128235760061b0190565b6127fd565b35610b388161021b565b91908110156128235760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561022c570190565b903590601e198136030182121561022c570180359067ffffffffffffffff821161022c5760200191813603831361022c57565b92919267ffffffffffffffff821161066957604051916128cf601f8201601f1916602001846106de565b82948184528183011161022c578281602093846000960137010152565b9061023c6040516128fc816106a6565b925463ffffffff8082168552602082811c821690860152604082811c61ffff1690860152605082901c81166060860152607082901c16608085015260901c60ff16151560a0840152565b80518210156128235760209160051b010190565b90929161298f61297e8367ffffffffffffffff166000526009602052604060002090565b5460081b6001600160e01b03191690565b90612999816127b4565b9560005b8281106129ae575050505050505090565b6129c16129bc828489612813565b612828565b83886129db6129d1858484612832565b6040810190612872565b905060208111612af8575b508392612a1c612a16612a0f612a05600198612a5797612a5297612832565b6020810190612872565b36916128a5565b89613a7c565b612a3a8967ffffffffffffffff16600052600a602052604060002090565b906001600160a01b0316600052602052604060002090565b6128ec565b60a081015115612abc57612aa0612a786060612a9293015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b03601f1981018352826106de565b612aaa828b612946565b52612ab5818a612946565b500161299d565b50612a92612aa0612af384612ae58a67ffffffffffffffff166000526009602052604060002090565b015460101c63ffffffff1690565b612a78565b915050612b30611850612b2384612a3a8b67ffffffffffffffff16600052600a602052604060002090565b5460701c63ffffffff1690565b10612b3d578388386129e6565b7f36f536ca000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b60405190612b808261064d565b60006040838281528260208201520152565b903590601e198136030182121561022c570180359067ffffffffffffffff821161022c57602001918160061b3603831361022c57565b35906001600160e01b038216820361022c57565b60408136031261022c57612c0b602060405192612bf88461068a565b8035612c038161021b565b845201612bc8565b602082015290565b60408136031261022c57612c0b602060405192612c2f8461068a565b612c03816102ab565b90612c41613e0e565b60005b8251811015612d735780612c5a60019285612946565b517f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a360a06001600160a01b038351169260608101936001600160a01b0380865116957fffff000000000000000000000000000000000000000000000000000000000000612cf860208601947fffffffffffffffffffff00000000000000000000000000000000000000000000865116604088019a848c51169261590d565b977fffffffffffffffffffff000000000000000000000000000000000000000000006080870195612d46875115158c600052600460205260406000209060ff60ff1983541691151516179055565b8560405198511688525116602087015251166040850152511660608301525115156080820152a201612c44565b509050565b60405190612d858261068a565b60006020838281520152565b90604051612d9e8161068a565b91546001600160e01b038116835260e01c6020830152565b9061023c612efe6001612dc7610710565b94612e9d612e938254612de3612ddd8260ff1690565b15158a52565b61ffff600882901c1660208a015263ffffffff601882901c1660408a015263ffffffff603882901c1660608a015263ffffffff605882901c1660808a015260ff607882901c1660a08a015260ff608082901c1660c08a015261ffff608882901c1660e08a015263ffffffff609882901c166101008a015261ffff60b882901c166101208a015261ffff60c882901c166101408a01526001600160e01b0319600882901b166101608a015260f81c90565b1515610180880152565b015461ffff81166101a086015263ffffffff601082901c166101c086015263ffffffff603082901c166101e086015267ffffffffffffffff605082901c1661020086015263ffffffff609082901c1661022086015260b01c63ffffffff1690565b63ffffffff16610240840152565b90612f15613e0e565b6000915b805183101561310757612f2c8382612946565b5190612f40825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b855180518210156130f257612f6c82602092612946565b510151612f7d611788838951612946565b8151602083015163ffffffff9081169116818110156130b9575050608082015163ffffffff1660208110613078575090867f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b56001600160a01b0384613007858f60019998612a3a6130029267ffffffffffffffff16600052600a602052604060002090565b613e4c565b61306f60405192839216958291909160a08060c083019463ffffffff815116845263ffffffff602082015116602085015261ffff604082015116604085015263ffffffff606082015116606085015263ffffffff608082015116608085015201511515910152565b0390a301612f55565b7f24ecdc02000000000000000000000000000000000000000000000000000000006000526001600160a01b0390911660045263ffffffff1660245260446000fd5b7f0b4f67a20000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b50509550925092600191500191929092612f19565b50905060005b81518110156131b7578061313561312660019385612946565b515167ffffffffffffffff1690565b67ffffffffffffffff6001600160a01b0361316460206131558689612946565b5101516001600160a01b031690565b600061318882612a3a8767ffffffffffffffff16600052600a602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a30161310d565b5050565b60208183031261022c5780359067ffffffffffffffff821161022c570181601f8201121561022c578035906131ef82610c84565b926131fd60405194856106de565b8284526020606081860194028301019181831161022c57602001925b828410613227575050505090565b60608483031261022c5760206060916040516132428161064d565b863561324d8161021b565b815261325a838801612bc8565b8382015261326a604088016111d5565b6040820152815201930192613219565b906040516132878161064d565b604060ff8294546001600160a01b0381168452818160a01c16602085015260a81c161515910152565b906132b9613e0e565b60005b8251811015612d73576132cf8184612946565b5160206132df6131268487612946565b9101519067ffffffffffffffff8116801580156134c9575b801561349b575b8015613408575b6133d05791613396826001959461334661333961297e61339b9767ffffffffffffffff166000526009602052604060002090565b6001600160e01b03191690565b6133a1577f71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad604051806133798782610fce565b0390a267ffffffffffffffff166000526009602052604060002090565b614143565b016132bc565b7f2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e3604051806133798782610fce565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b506001600160e01b03196134286101608501516001600160e01b03191690565b16630a04b54b60e21b8114159081613489575b81613477575b81613465575b81613453575b50613305565b63647e2ba960e01b915014153861344d565b63c4e0595360e01b8114159150613447565b632b1dfffb60e21b8114159150613441565b6307842f7160e21b811415915061343b565b506101e083015163ffffffff1663ffffffff6134c1611850606087015163ffffffff1690565b9116116132fe565b5063ffffffff6134e16101e085015163ffffffff1690565b16156132f7565b6134f0613e0e565b60208101519160005b835181101561357d5780613512610aa760019387612946565b61353461352f6001600160a01b0383165b6001600160a01b031690565b615d56565b613540575b50016134f9565b6040516001600160a01b039190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613539565b5091505160005b81518110156131b75761359a610aa78284612946565b906001600160a01b03821615613610577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613607836135ec6135e76135236001976001600160a01b031690565b615cdd565b506040516001600160a01b0390911681529081906020820190565b0390a101613584565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b613642613e0e565b60005b81518110156131b757806001600160a01b0361366360019385612946565b5151167fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d61370067ffffffffffffffff602061369f8689612946565b51015116836000526008602052604060002067ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556040519182918291909167ffffffffffffffff6020820193169052565b0390a201613645565b919082039182116126e557565b61371e612d78565b5061374461373f826001600160a01b03166000526006602052604060002090565b612d91565b602081019161376361375d611850855163ffffffff1690565b42613709565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000161161380b576117af6137af916001600160a01b03166000526007602052604060002090565b6137bf6117c46040830151151590565b8015613811575b61380b576137d3906147db565b9163ffffffff6137fb6118506137f0602087015163ffffffff1690565b935163ffffffff1690565b911610613806575090565b905090565b50905090565b506001600160a01b0361382b82516001600160a01b031690565b16156137c6565b90600282018092116126e557565b90602082018092116126e557565b90600182018092116126e557565b919082018092116126e557565b9061ffff8091169116029061ffff82169182036126e557565b63ffffffff60209116019063ffffffff82116126e557565b9063ffffffff8091169116019063ffffffff82116126e557565b906138bd613e0e565b60005b8251811015612d7357806138d660019285612946565b517fe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf6139fb60206001600160a01b03845116930151836000526007602052604060002061395b6001600160a01b0383511682906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b602082015181547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000075ff0000000000000000000000000000000000000000006040870151151560a81b169360a01b169116171790556040519182918291909160408060608301946001600160a01b03815116845260ff602082015116602085015201511515910152565b0390a2016138c0565b613a0d81613716565b9063ffffffff602083015116158015613a6a575b613a335750516001600160e01b031690565b6001600160a01b03907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b506001600160e01b0382511615613a21565b6001600160e01b0319169190630a04b54b60e21b8314613b15576307842f7160e21b8314613b0757632b1dfffb60e21b8314613afc5763647e2ba960e01b8314613af15763c4e0595360e01b8314613ae35782632ee8207560e01b60005260045260246000fd5b61023c919250600b90615411565b61023c919250615472565b61023c9192506153ae565b61023c919250600190615411565b61023c91925061532f565b916001600160e01b03198316630a04b54b60e21b8114613be2576307842f7160e21b8114613bc257632b1dfffb60e21b8114613bb65763647e2ba960e01b8114613baa5763c4e0595360e01b14613b9057632ee8207560e01b6000526001600160e01b0319831660045260246000fd5b61023c925015613ba257600b90615411565b600090615411565b505061023c9150615472565b505061023c91506153ae565b5061023c925015613bd95760ff60015b1690615411565b60ff6000613bd2565b505061023c915061532f565b33600052600360205260406000205415613c0457565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6111c4613c599196949395929667ffffffffffffffff166000526009602052604060002090565b946101608601946001600160e01b0319613c7b87516001600160e01b03191690565b16630a04b54b60e21b8114908115613dfd575b8115613dec575b50613da75763c4e0595360e01b6001600160e01b0319613cbd88516001600160e01b03191690565b1614613d895750506307842f7160e21b6001600160e01b0319613ce886516001600160e01b03191690565b1614613d205761196e613d0385516001600160e01b03191690565b632ee8207560e01b6000526001600160e01b031916600452602490565b613d749350612a0f6060613d5e8763ffffffff613d55610180613d4d86613d829b9d015163ffffffff1690565b930151151590565b911685876157c0565b0151604051958691602083019190602083019252565b03601f1981018652856106de565b9160019190565b94509450610b3891613d9c9136916128a5565b9360019336916128a5565b94509491613dcd91613dc76118506101e0610b3896015163ffffffff1690565b9161554f565b93613de46020613ddc87615658565b960151151590565b9336916128a5565b63647e2ba960e01b91501438613c95565b632b1dfffb60e21b81149150613c8e565b6001600160a01b03600154163303613e2257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b815181546020808501516040808701517fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000090941663ffffffff958616179190921b67ffffffff00000000161791901b69ffff000000000000000016178255606083015161023c93613f5c9260a092613efe911685547fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff1660509190911b6dffffffff0000000000000000000016178555565b613f55613f12608083015163ffffffff1690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911b71ffffffff000000000000000000000000000016178555565b0151151590565b81547fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1690151560901b72ff00000000000000000000000000000000000016179055565b91929092613fb08282868661590d565b600052600460205260ff6040600020541615613fcc5750505050565b6040517f097e17ff0000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529390921660248401527fffffffffffffffffffff0000000000000000000000000000000000000000000090911660448301527fffff000000000000000000000000000000000000000000000000000000000000166064820152608490fd5b0390fd5b604d81116126e557600a0a90565b60ff1660120160ff81116126e55760ff169060248211156140ef5760231982019182116126e5576140a56140ab92614064565b90612763565b6001600160e01b0381116140c5576001600160e01b031690565b7f10cb51d10000000000000000000000000000000000000000000000000000000060005260046000fd5b9060240390602482116126e55761219561410892614064565b6140ab565b9060ff80911691160160ff81116126e55760ff169060248211156140ef5760231982019182116126e5576140a56140ab92614064565b90614727610240600161023c9461417061415d8651151590565b829060ff60ff1983541691151516179055565b6141b6614182602087015161ffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff1660089190911b62ffff0016178255565b6142026141ca604087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff1660189190911b66ffffffff00000016178255565b614252614216606087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffff1660389190911b6affffffff0000000000000016178255565b6142a6614266608087015163ffffffff1690565b82547fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff1660589190911b6effffffff000000000000000000000016178255565b6142f86142b760a087015160ff1690565b82547fffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffff1660789190911b6fff00000000000000000000000000000016178255565b61434b61430960c087015160ff1690565b82547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff1660809190911b70ff0000000000000000000000000000000016178255565b6143a161435d60e087015161ffff1690565b82547fffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffff1660889190911b72ffff000000000000000000000000000000000016178255565b6143fe6143b661010087015163ffffffff1690565b82547fffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffff1660989190911b76ffffffff0000000000000000000000000000000000000016178255565b61445b61441161012087015161ffff1690565b82547fffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffff1660b89190911b78ffff000000000000000000000000000000000000000000000016178255565b6144ba61446e61014087015161ffff1690565b82547fffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffff1660c89190911b7affff0000000000000000000000000000000000000000000000000016178255565b6145236144d36101608701516001600160e01b03191690565b82547fff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff1660089190911c7effffffff00000000000000000000000000000000000000000000000000000016178255565b614584614534610180870151151590565b82547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178255565b01926145c86145996101a083015161ffff1690565b859061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6146146145dd6101c083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178555565b6146646146296101e083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1660309190911b69ffffffff00000000000016178555565b6146c061467d61020083015167ffffffffffffffff1690565b85547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff1660509190911b71ffffffffffffffff0000000000000000000016178555565b61471c6146d561022083015163ffffffff1690565b85547fffffffffffffffffffff00000000ffffffffffffffffffffffffffffffffffff1660909190911b75ffffffff00000000000000000000000000000000000016178555565b015163ffffffff1690565b7fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff0000000000000000000000000000000000000000000083549260b01b169116179055565b519069ffffffffffffffffffff8216820361022c57565b908160a091031261022c5761479c81614771565b91602082015191604081015191610b38608060608401519301614771565b6040513d6000823e3d90fd5b9081602091031261022c5751610b3881611ad9565b6147e3612d78565b506147fb61352361352383516001600160a01b031690565b90604051907ffeaf968c00000000000000000000000000000000000000000000000000000000825260a082600481865afa92831561490357600092600094614908575b50600083126140c5576020600491604051928380927f313ce5670000000000000000000000000000000000000000000000000000000082525afa92831561490357610b389363ffffffff936148a4936000926148cd575b506020015160ff165b9061410d565b926148bf6148b0610701565b6001600160e01b039095168552565b1663ffffffff166020830152565b61489e9192506148f4602091823d84116148fc575b6148ec81836106de565b8101906147c6565b929150614895565b503d6148e2565b6147ba565b90935061492e91925060a03d60a01161493b575b61492681836106de565b810190614788565b509392505091923861483e565b503d61491c565b9081602091031261022c573590565b91906149606020830183612872565b939050604083016149718185612b92565b90506040840191614989611850845163ffffffff1690565b808811614f535750602085015161ffff16808311614f1d57506101608501966149ba88516001600160e01b03191690565b6001600160e01b03198116630a04b54b60e21b81148015614f0d575b8015614efd575b15614a5857505050505050509181614a52612a0f614a39614a4b96614a086080610b38980186612872565b6101e083015163ffffffff169063ffffffff614a31610180613ddc606088015163ffffffff1690565b941692615b12565b51958694516001600160e01b03191690565b9280612872565b90613b20565b63c4e0595360e01b819b9a939495979996989b14600014614ccd575050614ada614ab9614acd999a614a8d6080880188612872565b63ffffffff614ab1610180614aa9606087015163ffffffff1690565b950151151590565b931691615a5b565b918251998a91516001600160e01b03191690565b614a52612a0f8880612872565b6060810151519082614af7614aef8780612872565b810190614942565b614cb0575081614c7a575b8515159081614c6d575b50614c435760408111614c115750614b3190614b2b8594939795612718565b9061385c565b946000935b838510614b8f575050505050611850614b53915163ffffffff1690565b808211614b5f57505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9091929395600190614be6611850612b23614bbe8667ffffffffffffffff16600052600a602052604060002090565b614bcf6129bc8d6108f08b8d612b92565b6001600160a01b0316600052602052604060002090565b8015614c0157614bf59161385c565b965b0193929190614b36565b50614c0b90613840565b96614bf7565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538614b0c565b61196e827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614cc7919350614b2b614cc28461384e565b6126ea565b91614b02565b6307842f7160e21b03614edf5750614d3a614d17614acd999a614cf36080880188612872565b63ffffffff614d0f610180614aa9606087015163ffffffff1690565b9316916157c0565b91614d29611850845163ffffffff1690565b998a91516001600160e01b03191690565b6080810151519082614d4f614aef8780612872565b614ec7575081614e91575b85151580614e85575b614c435760408211614e51576020015167ffffffffffffffff9081169081831c16614e17575050614d9b90614b2b8594939795612700565b946000935b838510614dbd575050505050611850614b53915163ffffffff1690565b9091929395600190614dec611850612b23614bbe8667ffffffffffffffff16600052600a602052604060002090565b8015614e0757614dfb9161385c565b965b0193929190614da0565b50614e1190613840565b96614dfd565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614d63565b61196e827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614ed9919350614b2b614cc284613832565b91614d5a565b632ee8207560e01b6000526001600160e01b03191660045260246000fd5b5063647e2ba960e01b81146149dd565b50632b1dfffb60e21b81146149d6565b7fd88dddd600000000000000000000000000000000000000000000000000000000600052600483905261ffff1660245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b67ffffffffffffffff8116600052600560205260406000209160405192614fab8461068a565b546001600160e01b038116845260e01c9182602085015263ffffffff82169283614fe5575b50505050610b3890516001600160e01b031690565b63ffffffff1642908103939084116126e55783116150035780614fd0565b7ff08bcb3e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045263ffffffff1660245260445260646000fd5b60408136031261022c5760206040519161505f8361068a565b803561506a8161021b565b83520135602082015290565b9694919695929390956000946000986000986000965b8088106150a0575050505050505050929190565b9091929394959697999a6150bd6150b88a848b612813565b615046565b9a6150f4612a528d614bcf6150e68967ffffffffffffffff16600052600a602052604060002090565b91516001600160a01b031690565b916151056117c460a0850151151590565b61524e5760009c60408401906151206122f5835161ffff1690565b6151d6575b5050606083015163ffffffff1661513b9161389a565b9c608083015161514e9063ffffffff1690565b6151579161389a565b9b82516151679063ffffffff1690565b63ffffffff16615176906126c3565b600193908083106151ca57506123b0611850602061519993015163ffffffff1690565b8082116151b957506151aa9161385c565b985b019695949392919061508c565b90506151c49161385c565b986151ac565b9150506151c49161385c565b9061219561523f939f61522d6152369460208f8e6122f595506001600160a01b0361520885516001600160a01b031690565b91166001600160a01b03821614615247576152239150613a04565b915b015190615b50565b925161ffff1690565b620186a0900490565b9b3880615125565b5091615225565b999b506001915061526f8461526961527593614b2b8b6126c3565b9b61389a565b9c613882565b9a6151ac565b91939093806101e00193846101e0116126e55761012081029080820461012014901517156126e5576101e09101018093116126e5576122f5610140615311610b38966dffffffffffffffffffffffffffff6122826152fc6152e961531b9a63ffffffff6121959a169061385c565b6121956122f56101208c015161ffff1690565b614b2b6118506101008b015163ffffffff1690565b93015161ffff1690565b612730565b9081602091031261022c575190565b60208151036153655761534b6020825183010160208301615320565b6001600160a01b0381119081156153a2575b506153655750565b614060906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906102f1565b6104009150103861535d565b60208151036153d457600b6153cc6020835184010160208401615320565b106153d45750565b614060906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840181815201906102f1565b9060208251036154375780615424575050565b6153cc6020835184010160208401615320565b6040517fe0d7fb02000000000000000000000000000000000000000000000000000000008152602060048201528061406060248201856102f1565b6024815103615488576022810151156154885750565b614060906040519182917f373b0e4400000000000000000000000000000000000000000000000000000000835260206004840181815201906102f1565b919091356001600160e01b0319811692600481106154e1575050565b6001600160e01b0319929350829060040360031b1b161690565b909291928360041161022c57831161022c57600401916003190190565b9060041161022c5790600490565b9081604091031261022c576020604051916155408361068a565b805183520151612c0b81610c9c565b91615558612d78565b5081156156365750615581612a0f828061557b6001600160e01b031995876154c5565b956154fb565b91167f181dcf100000000000000000000000000000000000000000000000000000000081036155be575080602080610b3893518301019101615526565b7f97a657c9000000000000000000000000000000000000000000000000000000001461560e577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060208061562193518301019101615320565b615629610701565b9081526000602082015290565b91505067ffffffffffffffff61564a610701565b911681526000602082015290565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610b386064826106de565b604051906156af8261066e565b60606080836000815260006020820152600060408201526000838201520152565b9080601f8301121561022c5781356156e781610c84565b926156f560405194856106de565b81845260208085019260051b82010192831161022c57602001905b82821061571d5750505090565b8135815260209182019101615710565b60208183031261022c5780359067ffffffffffffffff821161022c570160a08183031261022c57604051916157618361066e565b61576a826111d5565b8352615778602083016102ab565b6020840152604082013561578b81610c9c565b604084015260608201356060840152608082013567ffffffffffffffff811161022c576157b892016156d0565b608082015290565b6157c86156a2565b5081156158e3577f1f3b3aba000000000000000000000000000000000000000000000000000000006001600160e01b031961580c6158068585615518565b906154c5565b16036158b9578161582892615820926154fb565b81019061572d565b91806158a3575b6158795763ffffffff615846835163ffffffff1690565b161161584f5790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fee433e990000000000000000000000000000000000000000000000000000000060005260046000fd5b506158b46117c46040840151151590565b61582f565b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b604080516001600160a01b039283166020820190815292909316908301527fffffffffffffffffffff0000000000000000000000000000000000000000000090921660608201527fffff00000000000000000000000000000000000000000000000000000000000090921660808301529061598b8160a08101612a92565b51902090565b6001600160a01b03610b389116600b615be4565b6001600160a01b03610b389116600b615d18565b604051906159c6826106c2565b606080836000815260006020820152600060408201520152565b60208183031261022c5780359067ffffffffffffffff821161022c570160808183031261022c5760405191615a14836106c2565b813583526020820135615a2681610c9c565b602084015260408201356040840152606082013567ffffffffffffffff811161022c57615a5392016156d0565b606082015290565b615a636159b9565b5081156158e3577f21ea4ca9000000000000000000000000000000000000000000000000000000006001600160e01b0319615aa16158068585615518565b16036158b95781615abd92615ab5926154fb565b8101906159e0565b9180615afc575b61587957815111615ad25790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50615b0d6117c46020840151151590565b615ac4565b9063ffffffff615b2f93959495615b27612d78565b50169161554f565b91825111615ad25780615b44575b6158795790565b50602081015115615b3d565b670de0b6b3a7640000916001600160e01b03615b6c9216612750565b0490565b80548210156128235760005260206000200190600090565b91615ba2918354906000199060031b92831b921b19161790565b9055565b80548015615bce576000190190615bbd8282615b70565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014615c965760001984018481116126e55783549360001985019485116126e5576000958583615c4797615c389503615c4d575b505050615ba6565b90600052602052604060002090565b55600190565b615c7d615c7791615c6e615c64615c8d9588615b70565b90549060031b1c90565b92839187615b70565b90615b88565b8590600052602052604060002090565b55388080615c30565b50505050600090565b805490680100000000000000008210156106695781615cc6916001615ba294018155615b70565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054615d1257615cfb816002615c9f565b600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615d4f5780615d3a83600193615c9f565b80549260005201602052604060002055600190565b5050600090565b600081815260036020526040902054908115615d4f576000198201908282116126e5576002549260001984019384116126e5578383615c479460009603615db6575b505050615da56002615ba6565b600390600052602052604060002090565b615da5615c7791615dce615c64615dd8956002615b70565b9283916002615b70565b55388080615d9856fea164736f6c634300081a000a",
}

var FeeQuoterABI = FeeQuoterMetaData.ABI

var FeeQuoterBin = FeeQuoterMetaData.Bin

func DeployFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []FeeQuoterTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoter, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterBin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
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

func (_FeeQuoter *FeeQuoterCaller) GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPriceFeedConfig", token)

	if err != nil {
		return *new(FeeQuoterTokenPriceFeedConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenPriceFeedConfig)).(*FeeQuoterTokenPriceFeedConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoter.Contract.GetTokenPriceFeedConfig(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoter.Contract.GetTokenPriceFeedConfig(&_FeeQuoter.CallOpts, token)
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

func (_FeeQuoter *FeeQuoterTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "onReport", metadata, report)
}

func (_FeeQuoter *FeeQuoterSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.Contract.OnReport(&_FeeQuoter.TransactOpts, metadata, report)
}

func (_FeeQuoter *FeeQuoterTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.Contract.OnReport(&_FeeQuoter.TransactOpts, metadata, report)
}

func (_FeeQuoter *FeeQuoterTransactor) SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "setReportPermissions", permissions)
}

func (_FeeQuoter *FeeQuoterSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.Contract.SetReportPermissions(&_FeeQuoter.TransactOpts, permissions)
}

func (_FeeQuoter *FeeQuoterTransactorSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.Contract.SetReportPermissions(&_FeeQuoter.TransactOpts, permissions)
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

func (_FeeQuoter *FeeQuoterTransactor) UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "updateTokenPriceFeeds", tokenPriceFeedUpdates)
}

func (_FeeQuoter *FeeQuoterSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdateTokenPriceFeeds(&_FeeQuoter.TransactOpts, tokenPriceFeedUpdates)
}

func (_FeeQuoter *FeeQuoterTransactorSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdateTokenPriceFeeds(&_FeeQuoter.TransactOpts, tokenPriceFeedUpdates)
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

type FeeQuoterPriceFeedPerTokenUpdatedIterator struct {
	Event *FeeQuoterPriceFeedPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterPriceFeedPerTokenUpdated)
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
		it.Event = new(FeeQuoterPriceFeedPerTokenUpdated)
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

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterPriceFeedPerTokenUpdated struct {
	Token           common.Address
	PriceFeedConfig FeeQuoterTokenPriceFeedConfig
	Raw             types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPriceFeedPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterPriceFeedPerTokenUpdatedIterator{contract: _FeeQuoter.contract, event: "PriceFeedPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterPriceFeedPerTokenUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterPriceFeedPerTokenUpdated, error) {
	event := new(FeeQuoterPriceFeedPerTokenUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterReportPermissionSetIterator struct {
	Event *FeeQuoterReportPermissionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterReportPermissionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterReportPermissionSet)
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
		it.Event = new(FeeQuoterReportPermissionSet)
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

func (it *FeeQuoterReportPermissionSetIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterReportPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterReportPermissionSet struct {
	ReportId   [32]byte
	Permission KeystoneFeedsPermissionHandlerPermission
	Raw        types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterReportPermissionSetIterator, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterReportPermissionSetIterator{contract: _FeeQuoter.contract, event: "ReportPermissionSet", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterReportPermissionSet, reportId [][32]byte) (event.Subscription, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterReportPermissionSet)
				if err := _FeeQuoter.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseReportPermissionSet(log types.Log) (*FeeQuoterReportPermissionSet, error) {
	event := new(FeeQuoterReportPermissionSet)
	if err := _FeeQuoter.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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
	return common.HexToHash("0x71e9302ab4e912a9678ae7f5a8542856706806f2817e1bf2a20b171e265cb4ad")
}

func (FeeQuoterDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x2431cc0363f2f66b21782c7e3d54dd9085927981a21bd0cc6be45a51b19689e3")
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

func (FeeQuoterPriceFeedPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0xe6a7a17d710bf0b2cd05e5397dc6f97a5da4ee79e31e234bf5f965ee2bd9a5bf")
}

func (FeeQuoterReportPermissionSet) Topic() common.Hash {
	return common.HexToHash("0x32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3")
}

func (FeeQuoterTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (FeeQuoterTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5")
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

	GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error)

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

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error)

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

	FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPriceFeedPerTokenUpdatedIterator, error)

	WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterPriceFeedPerTokenUpdated, error)

	FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterReportPermissionSetIterator, error)

	WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterReportPermissionSet, reportId [][32]byte) (event.Subscription, error)

	ParseReportPermissionSet(log types.Log) (*FeeQuoterReportPermissionSet, error)

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
