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

type FeeQuoterFeeTokenArgs struct {
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

var FeeQuoterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokens\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.FeeTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeeTokensUpdates\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeTokensToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.FeeTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Internal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPremiumMultiplierWeiPerEth\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Internal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveLegacyArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"struct Internal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAddedOrFeeUpdated\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeMultiplierWeiPerEth\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timePassed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c060405234610b96576166c38038038061001981610d67565b92833981019080820360c08112610b9657604013610b9657610039610d48565b81519091906001600160601b0381168103610b9657825261005c60208201610d8c565b6020830190815260408201516001600160401b038111610b965782019284601f85011215610b965783519361009861009386610da0565b610d67565b9460208087838152019160051b83010191878311610b9657602001905b828210610d305750505060608301516001600160401b038111610b965783019285601f85011215610b965783516100ee61009382610da0565b9460208087848152019260061b82010190888211610b9657602001915b818310610cf45750505060808101516001600160401b038111610b965781019086601f83011215610b965781519161014561009384610da0565b9260208085838152019160051b83010191898311610b965760208101915b838310610bb1575050505060a0810151906001600160401b038211610b96570186601f82011215610b965780519061019d61009383610da0565b976020610160818b86815201940283010191818311610b9657602001925b828410610a7857505050503315610a6757600180546001600160a01b031916331790556020956101ea87610d67565b946000865260003681376101fc610d48565b968752858888015260005b865181101561026e576001906001600160a01b03610225828a610de9565b51168a61023182610e6a565b61023e575b505001610207565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a610236565b5087955086519360005b85518110156102e9576001600160a01b036102938288610de9565b51169081156102d8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef89836102ca600195610f5d565b50604051908152a101610278565b6342bcdf7f60e11b60005260046000fd5b5085518792919087906001600160a01b0316158015610a55575b610a4457516001600160a01b031660a052516001600160601b031660805261032a82610d67565b9360008552600036813760005b85518110156103b55760019061036e6001600160a01b03610358838a610de9565b511680600052600a875260006040812055610fd5565b610379575b01610337565b818060a01b036103898289610de9565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2610373565b50828460005b845181101561044257806103d160019287610de9565b517f37b04c5bbd78ea1fe719591b7877f1fcf5748b289d069074c078439acbfd1ef085848060a01b038351169261042182820194878060401b0386511681600052600a8552604060002055610f9c565b50858060a01b0390511692858060401b03905116604051908152a2016103bb565b50925060005b83518110156107fc5761045b8185610de9565b51826001600160401b0361046f8488610de9565b51511691015190801580156107e9575b80156107cb575b801561074a575b61073657600081815260068552604090205460019392919060701b6001600160e01b03191661067b57807f6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b60405180610569868291909161012063ffffffff8161014084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff60808201511660808601528260e01b60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e08601528261010082015116610100860152015116910152565b0390a25b6000908152600685526040908190208251815484880151938501516060860151608087015160a08089015160c0808b015160e0808d01516101008e0151610120909e01516001600160e01b03199a8b1660ff9c15159c909c169b909b1760089d909d1b64ffffffff00169c909c1760289890981b68ffffffff0000000000169790971760489690961b6cffffffff000000000000000000169590951760689490941b6dff00000000000000000000000000169390931760709190911c63ffffffff60701b161760909390931b61ffff60901b16929092178a841b8b9003169690911b63ffffffff60a01b16959095179590941b63ffffffff60c01b1694909417921b90921617905501610448565b807f8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad6040518061072e868291909161012063ffffffff8161014084019580511515855282602082015116602086015282604082015116604086015282606082015116606086015260ff60808201511660808601528260e01b60a08201511660a086015261ffff60c08201511660c08601528260e08201511660e08601528261010082015116610100860152015116910152565b0390a261056d565b63c35aa79d60e01b60005260045260246000fd5b5063ffffffff60e01b60a083015116630a04b54b60e21b81141590816107b9575b816107a7575b81610795575b81610783575b5061048d565b63647e2ba960e01b141590508761077d565b63c4e0595360e01b8114159150610777565b632b1dfffb60e21b8114159150610771565b6307842f7160e21b811415915061076b565b5063ffffffff6101008301511663ffffffff60408401511610610486565b5063ffffffff610100830151161561047f565b5090600161080983610d67565b9160008352600091610a3f575b81925b815184101561098d5761082c8483610de9565b5180516001600160401b0316929086019190845b878451805183101561097c578261085691610de9565b51015184516001600160a01b039061086f908490610de9565b51511690604081019063ffffffff8251168b8110610965575091877f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed160808d97969460019663ffffffff60408f87815260078d528181208b8060a01b038a1682528d522092818c8185511695805467ffffffff00000000838801938451901b169660606bffffffff0000000000000000875160401b169101976cff0000000000000000000000008951151560601b16928a6cff000000000000000000000000199160018060601b0319161716171717905560405195865251168c850152511660408301525115156060820152a301909150610840565b6312766e0160e11b8a526004849052602452604489fd5b505050925093600191500192610819565b905083825b8251811015610a10576001906001600160401b036109b08286610de9565b515116828060a01b03846109c48488610de9565b51015116908087526007855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610992565b604051615659908161106a823960805181818161056601526111bc015260a05181818161058e01526111530152f35b610816565b63d794ef9560e01b60005260046000fd5b5081516001600160601b031615610303565b639b15e16f60e01b60005260046000fd5b8382036101608112610b9657610140610a8f610d48565b91610a9987610db7565b8352601f190112610b96576040519161014083016001600160401b03811184821017610b9b57604052610ace60208701610ddc565b8352610adc60408701610dcb565b6020840152610aed60608701610dcb565b6040840152610afe60808701610dcb565b606084015260a086015160ff81168103610b9657608084015260c08601516001600160e01b031981168103610b965760a084015260e08601519161ffff83168303610b96578360209360c0610160960152610b5c6101008901610dcb565b60e0820152610b6e6101208901610dcb565b610100820152610b816101408901610dcb565b610120820152838201528152019301926101bb565b600080fd5b634e487b7160e01b600052604160045260246000fd5b82516001600160401b038111610b965782016040818d03601f190112610b9657610bd9610d48565b90610be660208201610db7565b825260408101516001600160401b038111610b965760209101018c601f82011215610b965780518d610c1a61009383610da0565b92602060a08186868152019402820101918211610b9657602001918f5b828410610c57575050505091816020938480940152815201920191610163565b83900360a08112610b96576080610c6c610d48565b91610c7686610d8c565b8352601f190112610b96576040519160808301916001600160401b03831184841017610b9b5760a093602093604052610cb0848801610dcb565b8152610cbe60408801610dcb565b84820152610cce60608801610dcb565b6040820152610cdf60808801610ddc565b6060820152838201528152019201918f610c37565b6040838a0312610b96576020604091610d0b610d48565b610d1486610d8c565b8152610d21838701610db7565b8382015281520192019161010b565b60208091610d3d84610d8c565b8152019101906100b5565b60408051919082016001600160401b03811183821017610b9b57604052565b6040519190601f01601f191682016001600160401b03811183821017610b9b57604052565b51906001600160a01b0382168203610b9657565b6001600160401b038111610b9b5760051b60200190565b51906001600160401b0382168203610b9657565b519063ffffffff82168203610b9657565b51908115158203610b9657565b8051821015610dfd5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610dfd5760005260206000200190600090565b80548015610e54576000190190610e428282610e13565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600360205260409020548015610f2b576000198101818111610f1557600254600019810191908211610f1557808203610ec4575b505050610eb06002610e2b565b600052600360205260006040812055600190565b610efd610ed5610ee6936002610e13565b90549060031b1c9283926002610e13565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610ea3565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80549068010000000000000000821015610b9b5781610ee6916001610f5994018155610e13565b9055565b80600052600360205260406000205415600014610f9657610f7f816002610f32565b600254906000526003602052604060002055600190565b50600090565b80600052600960205260406000205415600014610f9657610fbe816008610f32565b600854906000526009602052604060002055600190565b6000818152600960205260409020548015610f2b576000198101818111610f1557600854600019810191908211610f155781810361102f575b50505061101b6008610e2b565b600052600960205260006040812055600190565b611051611040610ee6936008610e13565b90549060031b1c9283926008610e13565b9055600052600960205260406000205538808061100e56fe6080604052600436101561001257600080fd5b60003560e01c806241e5be146101a657806301447eaa146101a1578063061877e31461019c57806306285c691461019757806315c34d5b14610192578063181f5a771461018d5780632451a62714610188578063277df5e3146101835780633937306f1461017e5780633a49bb491461017957806345ac924d146101745780634ab35b0b1461016f578063514e8cff1461016a5780636def4ce71461016557806379ba50971461016057806382b49eb01461015b5780638da5cb5b1461015657806391a2749a146101515780639cc199961461014c578063bce21f2214610147578063cdc73d5114610142578063d02641a01461013d578063d8694ccd14610138578063f2fde38b146101335763ffdb4b371461012e57600080fd5b612208565b612114565b611deb565b611d6f565b611cf7565b611b4a565b611a68565b611998565b611946565b611818565b61172f565b6115b7565b611426565b6113bb565b6112b6565b6110c3565b610ccd565b610bd5565b610aea565b610a1d565b610793565b6104fb565b610439565b61037c565b6101ce565b73ffffffffffffffffffffffffffffffffffffffff8116036101c957565b600080fd5b346101c95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957602061022360043561020e816101ab565b6024356044359161021e836101ab565b6123ed565b604051908152f35b6004359067ffffffffffffffff821682036101c957565b6024359067ffffffffffffffff821682036101c957565b359067ffffffffffffffff821682036101c957565b9181601f840112156101c95782359167ffffffffffffffff83116101c9576020808501948460051b0101116101c957565b919082519283825260005b8481106102e95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016102aa565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061033157505050505090565b909192939460208061036d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161029f565b97019301930191939290610322565b346101c95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576103b361022b565b60243567ffffffffffffffff81116101c9576103d390369060040161026e565b6044929192359167ffffffffffffffff83116101c957366023840112156101c95782600401359167ffffffffffffffff83116101c9573660248460061b860101116101c957610435946024610429950192612609565b604051918291826102fe565b0390f35b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95773ffffffffffffffffffffffffffffffffffffffff600435610489816101ab565b166000818152600a6020526040812054918215806104e7575b6104bb5760208367ffffffffffffffff60405191168152f35b602492507f02b56686000000000000000000000000000000000000000000000000000000008252600452fd5b5080825260096020526040822054156104a2565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957610532612870565b5060408051610540816105ed565b73ffffffffffffffffffffffffffffffffffffffff60206bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283815201817f0000000000000000000000000000000000000000000000000000000000000000168152835192835251166020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761060957604052565b6105be565b6080810190811067ffffffffffffffff82111761060957604052565b610140810190811067ffffffffffffffff82111761060957604052565b60a0810190811067ffffffffffffffff82111761060957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761060957604052565b604051906106b3604083610663565b565b604051906106b361014083610663565b67ffffffffffffffff81116106095760051b60200190565b359063ffffffff821682036101c957565b801515036101c957565b35906106b3826106ee565b81601f820112156101c95780359061071a826106c5565b926107286040519485610663565b82845260208085019360061b830101918183116101c957602001925b828410610752575050505090565b6040848303126101c9576020604091825161076c816105ed565b61077587610259565b815282870135610784816101ab565b83820152815201930192610744565b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c957366023820112156101c95780600401356107ed816106c5565b916107fb6040519384610663565b8183526024602084019260051b820101903682116101c95760248101925b82841061084c576024358567ffffffffffffffff82116101c95761084461084a923690600401610703565b90612889565b005b833567ffffffffffffffff81116101c957820160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc82360301126101c95760405190610898826105ed565b6108a460248201610259565b8252604481013567ffffffffffffffff81116101c957602491010136601f820112156101c95780356108d5816106c5565b916108e36040519384610663565b818352602060a08185019302820101903682116101c957602001915b81831061091e5750505091816020938480940152815201930192610819565b82360360a081126101c95760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192610959846105ed565b8635610964816101ab565b845201126101c95760a09160209160405161097e8161060e565b6109898488016106dd565b8152610997604088016106dd565b848201526109a7606088016106dd565b604082015260808701356109ba816106ee565b6060820152838201528152019201916108ff565b67ffffffffffffffff811161060957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60405190610a17602083610663565b60008252565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576104356040805190610a5e8183610663565b601382527f46656551756f74657220312e372e302d6465760000000000000000000000000060208301525191829160208352602083019061029f565b602060408183019282815284518094520192019060005b818110610abe5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610ab1565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576040516002548082526020820190600260005260206000209060005b818110610b585761043585610b4c81870382610663565b60405191829182610a9a565b8254845260209093019260019283019201610b35565b9080601f830112156101c9578135610b85816106c5565b92610b936040519485610663565b81845260208085019260051b8201019283116101c957602001905b828210610bbb5750505090565b602080918335610bca816101ab565b815201910190610bae565b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c957610c24903690600401610b6e565b60243567ffffffffffffffff81116101c957366023820112156101c957806004013591610c50836106c5565b91610c5e6040519384610663565b8383526024602084019460061b820101903682116101c957602401935b818510610c8c5761084a8484612b98565b6040853603126101c95760206040918251610ca6816105ed565b8735610cb1816101ab565b8152610cbe838901610259565b83820152815201940193610c7b565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c957806004019060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c957610d47613d56565b610d518280612d14565b4263ffffffff1692915060005b818110610f2c57505060240190610d758284612d14565b92905060005b838110610d8457005b80610da3610d9e600193610d98868a612d14565b906124c9565b612dc8565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff610ef3610ed06020850194610ec2610e0187517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610e30610e0c6106a4565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092168252565b63ffffffff8c166020820152610e6b610e51845167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b81516020929092015160e01b7fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5167ffffffffffffffff1690565b93517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610d7b565b80610f45610f40600193610d988980612d14565b612d91565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a73ffffffffffffffffffffffffffffffffffffffff611027610ed0602085019461100d610faf87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610fba610e0c6106a4565b63ffffffff8d166020820152610e6b610fe7845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b5173ffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610d5e565b9181601f840112156101c95782359167ffffffffffffffff83116101c957602083818601950101116101c957565b926110c094926110b29285521515602085015260806040850152608084019061029f565b91606081840391015261029f565b90565b346101c95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576110fa61022b565b60243590611107826101ab565b6044359160643567ffffffffffffffff81116101c95761112b903690600401611060565b93909160843567ffffffffffffffff81116101c95761114e903690600401611060565b9290917f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff821614600014611234575050935b6bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168086116112035750916111f493916104359693613d9a565b9093916040519485948561108e565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9161123e926123ed565b936111ad565b602060408183019282815284518094520192019060005b8181106112685750505090565b90919260206040826112ab600194885163ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b01940192910161125b565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c95761130590369060040161026e565b61130e816106c5565b9161131c6040519384610663565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611349836106c5565b0160005b8181106113a457505060005b828110156113965760019061137a6113758260051b85016124de565b61393f565b61138482876125f5565b5261138f81866125f5565b5001611359565b604051806104358682611244565b6020906113af612870565b8282880101520161134d565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760206114006004356113fb816101ab565b612ded565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95767ffffffffffffffff61146661022b565b61146e612870565b5016600052600460205260406000206040519061148a826105ed565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c6020820152604051809161043582604081019263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b6106b3909291926101208061014083019561151284825115159052565b60208181015163ffffffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015160ff169085015260a0818101517fffffffff00000000000000000000000000000000000000000000000000000000169085015260c08181015161ffff169085015260e08181015163ffffffff16908501526101008181015163ffffffff1690850152015163ffffffff16910152565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95767ffffffffffffffff6115f761022b565b60006101206040516116088161062a565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015201521660005260066020526104356040600020611723611715604051926116628461062a565b5460ff811615158452611687600882901c63ffffffff165b63ffffffff166020860152565b63ffffffff602882901c16604085015263ffffffff604882901c16606085015260ff606882901c1660808501527fffffffff00000000000000000000000000000000000000000000000000000000607082901b1660a085015261ffff609082901c1660c085015263ffffffff60a082901c1660e085015263ffffffff60c082901c1661010085015260e01c90565b63ffffffff16610120830152565b604051918291826114f5565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760005473ffffffffffffffffffffffffffffffffffffffff811633036117ee577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576104356118bf61185561022b565b67ffffffffffffffff6024359161186b836101ab565b6000606060405161187b8161060e565b828152826020820152826040820152015216600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b60ff604051916118ce8361060e565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015260405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c957604051611a11816105ed565b816004013567ffffffffffffffff81116101c957611a359060043691850101610b6e565b8152602482013567ffffffffffffffff81116101c95761084a926004611a5e9236920101610b6e565b6020820152612f2b565b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957611a9f61022b565b6024359067ffffffffffffffff82116101c957611ad061043591611aca611aec943690600401611060565b916132a6565b63ffffffff60409493945195869560608752606087019061029f565b92166020850152838203604085015261029f565b359060ff821682036101c957565b35907fffffffff00000000000000000000000000000000000000000000000000000000821682036101c957565b359061ffff821682036101c957565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95760043567ffffffffffffffff81116101c957366023820112156101c957806004013590611ba5826106c5565b90611bb36040519283610663565b8282526024610160602084019402820101903682116101c957602401925b818410611be15761084a83613643565b83360361016081126101c9576101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192611c1e846105ed565b611c2788610259565b845201126101c95761016091602091611c3e6106b5565b611c498489016106f8565b8152611c57604089016106dd565b84820152611c67606089016106dd565b6040820152611c78608089016106dd565b6060820152611c8960a08901611b00565b6080820152611c9a60c08901611b0e565b60a0820152611cab60e08901611b3b565b60c0820152611cbd61010089016106dd565b60e0820152611ccf61012089016106dd565b610100820152611ce261014089016106dd565b61012082015283820152815201930192611bd1565b346101c95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576040516008548082526020820190600860005260206000209060005b818110611d595761043585610b4c81870382610663565b8254845260209093019260019283019201611d42565b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c9576040611daf600435611375816101ab565b611de98251809263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565bf35b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957611e2261022b565b60243567ffffffffffffffff81116101c957806004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101c957611e8b611e868267ffffffffffffffff166000526006602052604060002090565b612eeb565b611e9c611e988251151590565b1590565b6120dc576064830190611ee0611e98611eb4846124de565b73ffffffffffffffffffffffffffffffffffffffff166000526001600801602052604060002054151590565b61208a57611eef858285614637565b91611efc6113fb826124de565b9560448601600080611f0e8385612d14565b15905061205757505090611f5761ffff611f619388611f3260c089015161ffff1690565b91611f4e611f4760e08b015163ffffffff1690565b9187612d14565b94909316614d51565b90949093906124de565b611f6a90613cc0565b611f73916123a1565b96602401611f8091612528565b611f919263ffffffff16915061324a565b608083015160ff1660ff16611fa5916123a1565b60609092015163ffffffff1690611fbb916139be565b63ffffffff1690611fcb9161324a565b90611fd59161324a565b90611ff49067ffffffffffffffff166000526004602052604060002090565b54612011916dffffffffffffffffffffffffffff909116906123a1565b61201a90612339565b906120249161324a565b907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661204c916123b4565b604051908152602090f35b92915092611f6161208461207f61207661012089015163ffffffff1690565b63ffffffff1690565b612312565b916124de565b6120d8612096836124de565b7f2502348c0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6000fd5b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b346101c95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c95773ffffffffffffffffffffffffffffffffffffffff600435612164816101ab565b61216c613d0b565b163381146121de57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c957600435612243816101ab565b67ffffffffffffffff612254610242565b169081600052600660205260ff60406000205416156122b55761227690612ded565b6000918252600460209081526040928390205483517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152921690820152f35b507f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc10000149015171561233457565b6122e3565b90670de0b6b3a7640000820291808304670de0b6b3a7640000149015171561233457565b908160051b918083046020149015171561233457565b9061012c82029180830461012c149015171561233457565b90606c820291808304606c149015171561233457565b8181029291811591840414171561233457565b81156123be570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b61242c6124266110c094937bffffffffffffffffffffffffffffffffffffffffffffffffffffffff61241f8195612ded565b16906123a1565b92612ded565b16906123b4565b9061243d826106c5565b61244a6040519182610663565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061247882946106c5565b019060005b82811061248957505050565b80606060208093850101520161247d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156124d95760061b0190565b61249a565b356110c0816101ab565b91908110156124d95760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156101c9570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c9570180359067ffffffffffffffff82116101c9576020019181360383136101c957565b929192612585826109ce565b916125936040519384610663565b8294818452818301116101c9578281602093846000960137010152565b906040516125bd8161060e565b606060ff82945463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c166040850152821c161515910152565b80518210156124d95760209160051b010190565b90929161265661262d8367ffffffffffffffff166000526006602052604060002090565b5460701b7fffffffff000000000000000000000000000000000000000000000000000000001690565b9061266081612433565b9560005b828110612675575050505050505090565b6126886126838284896124c9565b6124de565b83886126a26126988584846124e8565b6040810190612528565b9050602081116127e8575b5083926126e36126dd6126d66126cc60019861272b97612726976124e8565b6020810190612528565b3691612579565b896139d8565b6127018967ffffffffffffffff166000526007602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b6125b0565b6060810151156127ae5761279261274c602061276693015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610663565b61279c828b6125f5565b526127a7818a6125f5565b5001612664565b506127666127926127e36127d68967ffffffffffffffff166000526006602052604060002090565b5460a01c63ffffffff1690565b61274c565b915050612820612076612813846127018b67ffffffffffffffff166000526007602052604060002090565b5460401c63ffffffff1690565b1061282d578388386126ad565b7f36f536ca0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6040519061287d826105ed565b60006020838281520152565b90612892613d0b565b6000915b8051831015612aca576128a983826125f5565b51906128bd825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b85518051821015612ab5576128e9826020926125f5565b5101516129156128fa8389516125f5565b515173ffffffffffffffffffffffffffffffffffffffff1690565b604082015163ffffffff1660208110612a67575090867f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed173ffffffffffffffffffffffffffffffffffffffff84612a14858f6001999861270161298c9267ffffffffffffffff166000526007602052604060002090565b815181546020808501516040808701516060978801517fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951663ffffffff96909616959095179190921b67ffffffff00000000161792901b6bffffffff0000000000000000169190911790151590921b6cff00000000000000000000000016919091179055565b612a5e604051928392169582919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b0390a3016128d2565b7f24ecdc020000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff90911660045263ffffffff1660245260446000fd5b50509550925092600191500191929092612896565b50905060005b8151811015612b945780612af8612ae9600193856125f5565b515167ffffffffffffffff1690565b67ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff612b416020612b2586896125f5565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b6000612b65826127018767ffffffffffffffff166000526007602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a301612ad0565b5050565b919091612ba3613d0b565b60005b8151811015612c1f5780612bc7612bc261100d600194866125f5565b615001565b612bd2575b01612ba6565b73ffffffffffffffffffffffffffffffffffffffff612bf461100d83866125f5565b167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2612bcc565b505060005b8251811015612d0f5780612c3a600192856125f5565b517f37b04c5bbd78ea1fe719591b7877f1fcf5748b289d069074c078439acbfd1ef073ffffffffffffffffffffffffffffffffffffffff612cf0612ce1612c95855173ffffffffffffffffffffffffffffffffffffffff1690565b94612cc66020820196612cc0612cb3895167ffffffffffffffff1690565b67ffffffffffffffff1690565b90615032565b505173ffffffffffffffffffffffffffffffffffffffff1690565b935167ffffffffffffffff1690565b60405167ffffffffffffffff919091168152921691602090a201612c24565b509050565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c9570180359067ffffffffffffffff82116101c957602001918160061b360383136101c957565b35907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821682036101c957565b6040813603126101c957612dc0602060405192612dad846105ed565b8035612db8816101ab565b845201612d68565b602082015290565b6040813603126101c957612dc0602060405192612de4846105ed565b612db881610259565b73ffffffffffffffffffffffffffffffffffffffff8116600052600560205260406000209060405191612e1f836105ed565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811680845260e09190911c6020840181905215908115612ec4575b50612e805750517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1615905038612e58565b906106b3604051612efb8161062a565b610120612f2082955460ff81161515845261168761167a8263ffffffff9060081c1690565b63ffffffff16910152565b612f33613d0b565b60208101519160005b8351811015612fe75780612f5561100d600193876125f5565b612f91612f8c73ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b6153b9565b612f9d575b5001612f3c565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138612f96565b5091505160005b8151811015612b945761300461100d82846125f5565b9073ffffffffffffffffffffffffffffffffffffffff8216156130a1577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6130988361307061306b612f7360019773ffffffffffffffffffffffffffffffffffffffff1690565b6154f9565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101612fee565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b90929192836004116101c95783116101c957600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b906004116101c95790600490565b9080601f830112156101c957813561312b816106c5565b926131396040519485610663565b81845260208085019260051b8201019283116101c957602001905b8282106131615750505090565b8135815260209182019101613154565b6020818303126101c95780359067ffffffffffffffff82116101c9570160a0818303126101c957604051916131a583610647565b6131ae826106dd565b83526131bc60208301610259565b602084015260408201356131cf816106ee565b604084015260608201356060840152608082013567ffffffffffffffff81116101c9576131fc9201613114565b608082015290565b600a019081600a1161233457565b600201908160021161233457565b906002820180921161233457565b906020820180921161233457565b906001820180921161233457565b9190820180921161233457565b90613261826109ce565b61326e6040519182610663565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061329c82946109ce565b0190602036910137565b611e866132ca9194939467ffffffffffffffff166000526006602052604060002090565b9260a08401937fffffffff0000000000000000000000000000000000000000000000000000000061331b86517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c000000000000000000000000000000000000000000000000000000008114908115613619575b81156135ef575b506135a5577f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006133bf87517fffffffff000000000000000000000000000000000000000000000000000000001690565b161461354f577fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061343187517fffffffff000000000000000000000000000000000000000000000000000000001690565b16146134b2576120d861346486517fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b6135159450906134d061207660406134d69594015163ffffffff1690565b9161414b565b9061354b6134f86134f36134ee60608601515161235d565b613212565b613257565b926135416040820151604051968791602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101875286610663565b5163ffffffff1690565b9190565b506135159350806135639161356b936130cb565b810190613171565b9061354b6135886134f361358360808601515161235d565b613204565b926135416060820151604051968791602083019190602083019252565b6135dc93945061354192916135d661207660406135ca61010086015163ffffffff1690565b94015163ffffffff1690565b92614096565b906135e5610a08565b91906110c0610a08565b7f647e2ba9000000000000000000000000000000000000000000000000000000009150143861334e565b7fac77ffec0000000000000000000000000000000000000000000000000000000081149150613347565b9061364c613d0b565b60005b8251811015612d0f5761366281846125f5565b516020613672612ae984876125f5565b9101519067ffffffffffffffff811680158015613920575b80156138f2575b80156137b3575b61377b579161374182600195946136f16136cc61262d6137469767ffffffffffffffff166000526006602052604060002090565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b61374c577f6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b6040518061372487826114f5565b0390a267ffffffffffffffff166000526006602052604060002090565b6142c6565b0161364f565b7f8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad6040518061372487826114f5565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b507fffffffff0000000000000000000000000000000000000000000000000000000061380260a08501517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c0000000000000000000000000000000000000000000000000000000081141590816138c7575b8161389c575b81613871575b81613846575b50613698565b7f647e2ba9000000000000000000000000000000000000000000000000000000009150141538613840565b7fc4e0595300000000000000000000000000000000000000000000000000000000811415915061383a565b7fac77ffec000000000000000000000000000000000000000000000000000000008114159150613834565b7f1e10bdc400000000000000000000000000000000000000000000000000000000811415915061382e565b5061010083015163ffffffff1663ffffffff613918612076604087015163ffffffff1690565b911611613691565b5063ffffffff61393861010085015163ffffffff1690565b161561368a565b73ffffffffffffffffffffffffffffffffffffffff9061395d612870565b50166000526005602052604060002060405190613979826105ed565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c602082015290565b63ffffffff60209116019063ffffffff821161233457565b9063ffffffff8091169116019063ffffffff821161233457565b7fffffffff000000000000000000000000000000000000000000000000000000001691907f2812d52c000000000000000000000000000000000000000000000000000000008314613b20577f1e10bdc4000000000000000000000000000000000000000000000000000000008314613b12577fac77ffec000000000000000000000000000000000000000000000000000000008314613b07577f647e2ba9000000000000000000000000000000000000000000000000000000008314613afc577fc4e05953000000000000000000000000000000000000000000000000000000008314613aed57827f2ee820750000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6106b391925061dee990614f4d565b6106b3919250614fae565b6106b3919250614eea565b6106b3919250600190614f4d565b6106b3919250614e5a565b917fffffffff0000000000000000000000000000000000000000000000000000000083167f2812d52c000000000000000000000000000000000000000000000000000000008114613cb4577f1e10bdc4000000000000000000000000000000000000000000000000000000008114613c94577fac77ffec000000000000000000000000000000000000000000000000000000008114613c88577f647e2ba9000000000000000000000000000000000000000000000000000000008114613c7c577fc4e059530000000000000000000000000000000000000000000000000000000014613c61577f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000831660045260246000fd5b6106b3925015613c745761dee990614f4d565b600090614f4d565b50506106b39150614fae565b50506106b39150614eea565b506106b3925015613cab5760ff60015b1690614f4d565b60ff6000613ca4565b50506106b39150614e5a565b73ffffffffffffffffffffffffffffffffffffffff166000818152600a602052604081205491821580613cf7575b6104bb57505090565b508082526009602052604082205415613cee565b73ffffffffffffffffffffffffffffffffffffffff600154163303613d2c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b33600052600360205260406000205415613d6c57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b611e86613dc19196949395929667ffffffffffffffff166000526006602052604060002090565b9460a08601947fffffffff00000000000000000000000000000000000000000000000000000000613e1287517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c00000000000000000000000000000000000000000000000000000000811490811561406c575b8115614042575b50613ffd5750507fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613eb886517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613fd2577f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613f2a86517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613f5d576120d861346485517fffffffff000000000000000000000000000000000000000000000000000000001690565b613f9f93506126d66060613f89613f826120766040613fcb989a015163ffffffff1690565b84866151f8565b0151604051958691602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610663565b9160019190565b613f9f93506126d66040613f89613ff661207683613fcb989a015163ffffffff1690565b848661414b565b945094916140239161401d6120766101006110c096015163ffffffff1690565b9161508d565b9361403a6020614032876151ae565b960151151590565b933691612579565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501438613e45565b7fac77ffec0000000000000000000000000000000000000000000000000000000081149150613e3e565b9063ffffffff6140b0936140a8612870565b50169161508d565b908151116140bb5790565b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110614119575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9060608060405161415b8161060e565b6000815260006020820152600060408201520152801561429c577f21ea4ca9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006141ca6141c48486613106565b906140e5565b160361427257806141da926130cb565b8192910160009260208183031261426e5780359067ffffffffffffffff821161426a57019260808483031261426757604051936142168561060e565b803585526020810135614228816106ee565b60208601526040810135604086015260608101359167ffffffffffffffff83116142675750614258929101613114565b60608301528151116140bb5790565b80fd5b8480fd5b8380fd5b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b6145dc6101206106b39361430e6142dd8251151590565b859060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b614358614322602083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ff1660089190911b64ffffffff0016178555565b6143a661436c604083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1660289190911b68ffffffff000000000016178555565b6143f86143ba606083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffff1660489190911b6cffffffff00000000000000000016178555565b614448614409608083015160ff1690565b85547fffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffff1660689190911b6dff0000000000000000000000000016178555565b6144bb61447860a08301517fffffffff000000000000000000000000000000000000000000000000000000001690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911c71ffffffff000000000000000000000000000016178555565b6145126144cd60c083015161ffff1690565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b61456f61452660e083015163ffffffff1690565b85547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b77ffffffff000000000000000000000000000000000000000016178555565b6145d161458461010083015163ffffffff1690565b85547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff1660c09190911b7bffffffff00000000000000000000000000000000000000000000000016178555565b015163ffffffff1690565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000083549260e01b169116179055565b908160209103126101c9573590565b91906146466020830183612528565b939050604083016146578185612d14565b9050602084019161466f612076845163ffffffff1690565b808811614d1f575060018211614ceb5760a08501966146ae88517fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f2812d52c0000000000000000000000000000000000000000000000000000000081148015614cc2575b8015614c99575b156147805750505050505050918161477a6126d66147496147739661472d60806110c0980186612528565b6135d661207660406135ca610100879697015163ffffffff1690565b51958694517fffffffff000000000000000000000000000000000000000000000000000000001690565b9280612528565b90613b2b565b7fc4e0595300000000000000000000000000000000000000000000000000000000819b9a939495979996989b14600014614a1d57505061481d6147e4614810999a60406134d06120766147d660808b018b612528565b939094015163ffffffff1690565b918251998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b61477a6126d68880612528565b606081015151908261483a6148328780612528565b810190614628565b614a005750816149ca575b85151590816149bd575b50614993576040811161496157506148749061486e859493979561238b565b9061324a565b946000935b8385106148d2575050505050612076614896915163ffffffff1690565b8082116148a257505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b90919293956001906149366120766128136149018667ffffffffffffffff166000526007602052604060002090565b6149126126838d610d988b8d612d14565b73ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b8015614951576149459161324a565b965b0193929190614879565b5061495b9061322e565b96614947565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b604091500151153861484f565b6120d8827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614a1791935061486e614a128461323c565b61235d565b91614845565b7f1e10bdc40000000000000000000000000000000000000000000000000000000003614c4a5750614aa5614a6a614810999a6040614a646120766147d660808b018b612528565b916151f8565b91614a7c612076845163ffffffff1690565b998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b6080810151519082614aba6148328780612528565b614c32575081614bfc575b85151580614bf0575b6149935760408211614bbc576020015167ffffffffffffffff9081169081831c16614b82575050614b069061486e8594939795612373565b946000935b838510614b28575050505050612076614896915163ffffffff1690565b9091929395600190614b576120766128136149018667ffffffffffffffff166000526007602052604060002090565b8015614b7257614b669161324a565b965b0193929190614b0b565b50614b7c9061322e565b96614b68565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b50606081015115614ace565b6120d8827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b614c4491935061486e614a1284613220565b91614ac5565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b507f647e2ba9000000000000000000000000000000000000000000000000000000008114614702565b507fac77ffec0000000000000000000000000000000000000000000000000000000081146146fb565b7fd88dddd6000000000000000000000000000000000000000000000000000000006000526004829052600160245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b9194929390946000926000966000966000945b808610614d7657505050505050929190565b909192939495614db1612726614da3889c9b9c67ffffffffffffffff166000526007602052604060002090565b6149126126838b87896124c9565b98614dc2611e9860608c0151151590565b614e205760208a015163ffffffff16614dda916139be565b9960408a0151614ded9063ffffffff1690565b614df6916139be565b9851614e079063ffffffff16612312565b614e109161324a565b956001905b019493929190614d64565b969850614e45614e3f85614e3960019461486e8a612312565b996139be565b9a6139a6565b98614e15565b908160209103126101c9575190565b6020815103614e9d57614e766020825183010160208301614e4b565b73ffffffffffffffffffffffffffffffffffffffff8111908115614ede575b50614e9d5750565b614eda906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352602060048401818152019061029f565b0390fd5b61040091501038614e95565b6020815103614f1057600b614f086020835184010160208401614e4b565b10614f105750565b614eda906040519182917fe0d7fb02000000000000000000000000000000000000000000000000000000008352602060048401818152019061029f565b906020825103614f735780614f60575050565b614f086020835184010160208401614e4b565b6040517fe0d7fb020000000000000000000000000000000000000000000000000000000081526020600482015280614eda602482018561029f565b6024815103614fc457602281015115614fc45750565b614eda906040519182917f373b0e44000000000000000000000000000000000000000000000000000000008352602060048401818152019061029f565b73ffffffffffffffffffffffffffffffffffffffff6110c0911680600052600a60205260006040812055600861556b565b9073ffffffffffffffffffffffffffffffffffffffff6110c092169081600052600a6020526040600020556008615534565b908160409103126101c95760206040519161507e836105ed565b805183520151612dc0816106ee565b91615096612870565b50811561518c57506150d76126d682806150d17fffffffff0000000000000000000000000000000000000000000000000000000095876140e5565b956130cb565b91167f181dcf100000000000000000000000000000000000000000000000000000000081036151145750806020806110c093518301019101615064565b7f97a657c90000000000000000000000000000000000000000000000000000000014615164577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060208061517793518301019101614e4b565b61517f6106a4565b9081526000602082015290565b91505067ffffffffffffffff6151a06106a4565b911681526000602082015290565b6020604051917f181dcf10000000000000000000000000000000000000000000000000000000008284015280516024840152015115156044820152604481526110c0606482610663565b6060608060405161520881610647565b6000815260006020820152600060408201526000838201520152811561429c577f1f3b3aba000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006152776141c48585613106565b1603614272578161528b92613563926130cb565b9063ffffffff61529f835163ffffffff1690565b16116152a85790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b9161530a918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b80548210156124d95760005260206000200190600090565b8054801561538a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061535b828261530e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115615496577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161233457600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612334578383615455946000960361545b575b5050506154446002615326565b600390600052602052604060002090565b55600190565b6154446154879161547d61547361548d95600261530e565b90549060031b1c90565b928391600261530e565b906152d2565b55388080615437565b5050600090565b8054906801000000000000000082101561060957816154c491600161530a9401815561530e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b60008181526003602052604090205461552e5761551781600261549d565b600254906000526003602052604060002055600190565b50600090565b600082815260018201602052604090205461549657806155568360019361549d565b80549260005201602052604060002055600190565b6001810191806000528260205260406000205492831515600014615643577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612334578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501948511612334576000958583615455976155fb950361560a575b505050615326565b90600052602052604060002090565b61562a6154879161562161547361563a958861530e565b9283918761530e565b8590600052602052604060002090565b553880806155f3565b5050505060009056fea164736f6c634300081a000a",
}

var FeeQuoterABI = FeeQuoterMetaData.ABI

var FeeQuoterBin = FeeQuoterMetaData.Bin

func DeployFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []FeeQuoterFeeTokenArgs, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoter, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterBin), backend, staticConfig, priceUpdaters, feeTokens, tokenTransferFeeConfigArgs, destChainConfigArgs)
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

func (_FeeQuoter *FeeQuoterCaller) ResolveLegacyArgs(opts *bind.CallOpts, destChainSelector uint64, extraArgs []byte) (ResolveLegacyArgs,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "resolveLegacyArgs", destChainSelector, extraArgs)

	outstruct := new(ResolveLegacyArgs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenReceiver = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ExecutorArgs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) ResolveLegacyArgs(destChainSelector uint64, extraArgs []byte) (ResolveLegacyArgs,

	error) {
	return _FeeQuoter.Contract.ResolveLegacyArgs(&_FeeQuoter.CallOpts, destChainSelector, extraArgs)
}

func (_FeeQuoter *FeeQuoterCallerSession) ResolveLegacyArgs(destChainSelector uint64, extraArgs []byte) (ResolveLegacyArgs,

	error) {
	return _FeeQuoter.Contract.ResolveLegacyArgs(&_FeeQuoter.CallOpts, destChainSelector, extraArgs)
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

func (_FeeQuoter *FeeQuoterTransactor) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []FeeQuoterFeeTokenArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyFeeTokensUpdates", feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoter *FeeQuoterSession) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []FeeQuoterFeeTokenArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToRemove, feeTokensToAdd)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyFeeTokensUpdates(feeTokensToRemove []common.Address, feeTokensToAdd []FeeQuoterFeeTokenArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToRemove, feeTokensToAdd)
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

type FeeQuoterFeeTokenAddedOrFeeUpdatedIterator struct {
	Event *FeeQuoterFeeTokenAddedOrFeeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterFeeTokenAddedOrFeeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterFeeTokenAddedOrFeeUpdated)
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
		it.Event = new(FeeQuoterFeeTokenAddedOrFeeUpdated)
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

func (it *FeeQuoterFeeTokenAddedOrFeeUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterFeeTokenAddedOrFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterFeeTokenAddedOrFeeUpdated struct {
	FeeToken               common.Address
	FeeMultiplierWeiPerEth *big.Int
	Raw                    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterFeeTokenAddedOrFeeUpdated(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedOrFeeUpdatedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "FeeTokenAddedOrFeeUpdated", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFeeTokenAddedOrFeeUpdatedIterator{contract: _FeeQuoter.contract, event: "FeeTokenAddedOrFeeUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchFeeTokenAddedOrFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAddedOrFeeUpdated, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "FeeTokenAddedOrFeeUpdated", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterFeeTokenAddedOrFeeUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAddedOrFeeUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseFeeTokenAddedOrFeeUpdated(log types.Log) (*FeeQuoterFeeTokenAddedOrFeeUpdated, error) {
	event := new(FeeQuoterFeeTokenAddedOrFeeUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAddedOrFeeUpdated", log); err != nil {
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
type ResolveLegacyArgs struct {
	TokenReceiver []byte
	GasLimit      uint32
	ExecutorArgs  []byte
}

func (FeeQuoterAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (FeeQuoterAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (FeeQuoterDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x6b11e6862fc0e0450589b9745e371904c955b26cdd7170afeb5351202f5a567b")
}

func (FeeQuoterDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x8cea13ba9137e529428513c2987788feb1cb8a27380c7b9432bc18359799e5ad")
}

func (FeeQuoterFeeTokenAddedOrFeeUpdated) Topic() common.Hash {
	return common.HexToHash("0x37b04c5bbd78ea1fe719591b7877f1fcf5748b289d069074c078439acbfd1ef0")
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

	ResolveLegacyArgs(opts *bind.CallOpts, destChainSelector uint64, extraArgs []byte) (ResolveLegacyArgs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToRemove []common.Address, feeTokensToAdd []FeeQuoterFeeTokenArgs) (*types.Transaction, error)

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

	FilterFeeTokenAddedOrFeeUpdated(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedOrFeeUpdatedIterator, error)

	WatchFeeTokenAddedOrFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAddedOrFeeUpdated, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenAddedOrFeeUpdated(log types.Log) (*FeeQuoterFeeTokenAddedOrFeeUpdated, error)

	FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenRemovedIterator, error)

	WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenRemoved(log types.Log) (*FeeQuoterFeeTokenRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeQuoterOwnershipTransferred, error)

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
