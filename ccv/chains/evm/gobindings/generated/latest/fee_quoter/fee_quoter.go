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
	NetworkFeeUSDCents          uint16
	LinkFeeMultiplierPercent    uint8
}

type FeeQuoterDestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"priceUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertTokenAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllDestChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.DestChainConfig[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllDestChainSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"destChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllTokenTransferFeeConfigs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig[]\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestinationChainGasPrice\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Internal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAndGasPrices\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokenPrice\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"gasPriceValue\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Internal.TimestampedPackedUint224\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrices\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.TimestampedPackedUint224[]\",\"components\":[{\"name\":\"value\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatedTokenPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint224\",\"internalType\":\"uint224\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessageArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"convertedExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPoolReturnData\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteGasForExec\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonCalldataGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"calldataSize\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasCostInUsdCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeTokenPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"premiumPercentMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFeeTokens\",\"inputs\":[{\"name\":\"feeTokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveLegacyArgs\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"struct Internal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.DestChainConfig\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByteBase\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"linkFeeMultiplierPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenAdded\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenRemoved\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeeQuoter.TokenTransferFeeConfig\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerUnitGasUpdated\",\"inputs\":[{\"name\":\"destChain\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeTokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidChainFamilySelector\",\"inputs\":[{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSVMExtraArgsWritableBitmap\",\"inputs\":[{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageComputeUnitLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageFeeTooHigh\",\"inputs\":[{\"name\":\"msgFeeJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageGasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageTooLarge\",\"inputs\":[{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoGasPriceAvailable\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooManySVMExtraArgsAccounts\",\"inputs\":[{\"name\":\"numAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAccounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManySuiExtraArgsReceiverObjectIds\",\"inputs\":[{\"name\":\"numReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxReceiverObjectIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedNumberOfTokens\",\"inputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c060405234610a3557616f618038038061001981610bc7565b92833981019080820360a08112610a3557604013610a3557610039610ba8565b81516001600160601b0381168103610a3557815261005960208301610bec565b6020820190815260408301519093906001600160401b038111610a355783019381601f86011215610a355784519461009861009387610c00565b610bc7565b9560208088838152019160051b83010191848311610a3557602001905b828210610b905750505060608401516001600160401b038111610a355784019382601f86011215610a35578451946100ef61009387610c00565b9560208088838152019160051b83010191858311610a355760208101915b838310610a5057505050506080810151906001600160401b038211610a35570182601f82011215610a355780519061014761009383610c00565b936020610180818786815201940283010191818311610a3557602001925b82841061090b575050505033156108fa57600180546001600160a01b0319163317905560209261019484610bc7565b926000845260003681376101a6610ba8565b968752838588015260005b8451811015610218576001906001600160a01b036101cf8288610c66565b5116876101db82610d62565b6101e8575b5050016101b1565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138876101e0565b508493508587519260005b8451811015610294576001600160a01b0361023e8287610c66565b5116908115610283577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8883610275600195610cea565b50604051908152a101610223565b6342bcdf7f60e11b60005260046000fd5b50845186919086906001600160a01b03161580156108e8575b6108d757516001600160a01b031660a052516001600160601b031660805260005b835181101561067a576102e18185610c66565b51826001600160401b036102f58488610c66565b51511691015181158015610667575b8015610649575b610634578161054e9160019493600052600a865263ffffffff60e01b60406000205460701b161560001461055557817f4efe320c85221c7c3684c54561bea5a9c4dcfad794c6ef9ff9e6b43fb307c0f860405180610410858291909161014060ff8161016084019580511515855263ffffffff602082015116602086015263ffffffff604082015116604086015263ffffffff606082015116606086015282608082015116608086015263ffffffff60e01b60a08201511660a086015261ffff60c08201511660c086015263ffffffff60e08201511660e086015263ffffffff6101008201511661010086015261ffff61012082015116610120860152015116910152565b0390a25b6000828152600a875260409081902082518154848a0151938501516060860151608087015160a08089015160c0808b015160e0808d01516101008e01516101208f0151610140909f01517fff00000000000000000000000000000000000000000000000000000000000000909b1660ff9c15159c909c169b909b1760089d909d1b64ffffffff00169c909c1760289890981b68ffffffff0000000000169790971760489690961b6cffffffff000000000000000000169590951760689490941b6dff00000000000000000000000000169390931760709190911c63ffffffff60701b161760909390931b61ffff60901b16929092179690911b63ffffffff60a01b16959095179290941b63ffffffff60c01b16919091179390921b61ffff60e01b169290921760f09190911b60ff60f01b16179055610d29565b50016102ce565b817f0c6380a4766d45f5d53ca170bf865bebfab44958dec379d5a90177264e6645b76040518061062c858291909161014060ff8161016084019580511515855263ffffffff602082015116602086015263ffffffff604082015116604086015263ffffffff606082015116606086015282608082015116608086015263ffffffff60e01b60a08201511660a086015261ffff60c08201511660c086015263ffffffff60e08201511660e086015263ffffffff6101008201511661010086015261ffff61012082015116610120860152015116910152565b0390a2610414565b5063c35aa79d60e01b60005260045260246000fd5b5063ffffffff6101008201511663ffffffff6040830151161061030b565b5063ffffffff6101008201511615610304565b50600161068682610bc7565b92600084526000916108d2575b81925b8151841015610812576106a98483610c66565b518051908201956001600160401b0390911692845b87518051821015610802576106d4828692610c66565b51015188516001600160a01b03906106ed908490610c66565b515116906040810163ffffffff8151168781106107eb57507f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed1608060019594938a9363ffffffff60408e600b8e898352528d828220908c8060a01b038b168352522091818d818651169480549567ffffffff00000000838901938451901b166bffffffff0000000000000000865160401b169060608a01986cff0000000000000000000000008a51151560601b16936cff000000000000000000000000199160018060601b031916171617171790556107c589610d29565b5081604051965116865251168d850152511660408301525115156060820152a3016106be565b6312766e0160e11b8a526004849052602452604489fd5b5050925094509260010192610696565b849150825b8251811015610895576001906001600160401b036108358286610c66565b515116828060a01b03846108498488610c66565b5101511690808752600b855260408720848060a01b038316885285528660408120557f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b8780a301610817565b6040516161179081610e4a82396080518181816105d90152611349015260a051818181610601015281816112e001528181611e5d01526128c90152f35b610693565b63d794ef9560e01b60005260046000fd5b5081516001600160601b0316156102ad565b639b15e16f60e01b60005260046000fd5b8382036101808112610a3557610160610922610ba8565b9161092c87610c17565b8352601f190112610a35576040519161016083016001600160401b03811184821017610a3a5760405261096160208701610c3c565b835261096f60408701610c2b565b602084015261098060608701610c2b565b604084015261099160808701610c2b565b60608401526109a260a08701610c49565b608084015260c0860151916001600160e01b031983168303610a35578360209360a06101809601526109d660e08901610c57565b60c08201526109e86101008901610c2b565b60e08201526109fa6101208901610c2b565b610100820152610a0d6101408901610c57565b610120820152610a206101608901610c49565b61014082015283820152815201930192610165565b600080fd5b634e487b7160e01b600052604160045260246000fd5b82516001600160401b038111610a355782016040818903601f190112610a3557610a78610ba8565b90610a8560208201610c17565b825260408101516001600160401b038111610a3557602091010188601f82011215610a35578051610ab861009382610c00565b91602060a08185858152019302820101908b8211610a3557602001915b818310610af4575050509181602093848094015281520192019161010d565b828c0360a08112610a35576080610b09610ba8565b91610b1386610bec565b8352601f190112610a35576040519160808301916001600160401b03831184841017610a3a5760a093602093604052610b4d848801610c2b565b8152610b5b60408801610c2b565b84820152610b6b60608801610c2b565b6040820152610b7c60808801610c3c565b606082015283820152815201920191610ad5565b60208091610b9d84610bec565b8152019101906100b5565b60408051919082016001600160401b03811183821017610a3a57604052565b6040519190601f01601f191682016001600160401b03811183821017610a3a57604052565b51906001600160a01b0382168203610a3557565b6001600160401b038111610a3a5760051b60200190565b51906001600160401b0382168203610a3557565b519063ffffffff82168203610a3557565b51908115158203610a3557565b519060ff82168203610a3557565b519061ffff82168203610a3557565b8051821015610c7a5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610c7a5760005260206000200190600090565b80549068010000000000000000821015610a3a5781610ccf916001610ce694018155610c90565b819391549060031b91821b91600019901b19161790565b9055565b80600052600360205260406000205415600014610d2357610d0c816002610ca8565b600254906000526003602052604060002055600190565b50600090565b80600052600960205260406000205415600014610d2357610d4b816008610ca8565b600854906000526009602052604060002055600190565b6000818152600360205260409020548015610e42576000198101818111610e2c57600254600019810191908211610e2c57808203610df2575b5050506002548015610ddc5760001901610db6816002610c90565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610e14610e03610ccf936002610c90565b90549060031b1c9283926002610c90565b90556000526003602052604060002055388080610d9b565b634e487b7160e01b600052601160045260246000fd5b505060009056fe6080604052600436101561001257600080fd5b60003560e01c806241e5be146101e657806301447eaa146101e1578063029c7290146101dc57806306285c69146101d7578063080d711a146101d257806315c34d5b146101cd578063181f5a77146101c85780632451a627146101c35780633937306f146101be5780633a49bb49146101b957806345ac924d146101b45780634ab35b0b146101af578063514e8cff146101aa5780636def4ce7146101a557806379ba5097146101a057806382b49eb01461019b57806389933a51146101965780638da5cb5b14610191578063910d8f591461018c57806391a2749a1461018757806391b2f60714610182578063947f82171461017d5780639cc1999614610178578063b2f6322214610173578063cdc73d511461016e578063d02641a014610169578063d8694ccd14610164578063f2fde38b1461015f5763ffdb4b371461015a57600080fd5b612c11565b612b1d565b61273f565b6126c3565b61262e565b612467565b612344565b6122d7565b612117565b611ffd565b611c96565b611c44565b611aff565b6119c8565b61189c565b611759565b6115b3565b611548565b611443565b611250565b610de9565b610d48565b610c6d565b6109e5565b6107d5565b61056e565b6104b7565b6103bc565b61020e565b73ffffffffffffffffffffffffffffffffffffffff81160361020957565b600080fd5b346102095760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602061026360043561024e816101eb565b6024356044359161025e836101eb565b612df6565b604051908152f35b6004359067ffffffffffffffff8216820361020957565b6024359067ffffffffffffffff8216820361020957565b359067ffffffffffffffff8216820361020957565b9181601f840112156102095782359167ffffffffffffffff8311610209576020808501948460051b01011161020957565b919082519283825260005b8481106103295750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016102ea565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061037157505050505090565b90919293946020806103ad837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102df565b97019301930191939290610362565b346102095760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576103f361026b565b60243567ffffffffffffffff8111610209576104139036906004016102ae565b6044929192359167ffffffffffffffff831161020957366023840112156102095782600401359167ffffffffffffffff8311610209573660248460061b860101116102095761047594602461046995019261300d565b6040519182918261033e565b0390f35b906020808351928381520192019060005b8181106104975750505090565b825167ffffffffffffffff1684526020938401939092019160010161048a565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576008546104f28161324a565b6008549160005b818110610516576040516020808252819061047590820186610479565b600090848110156105695760089091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee38101546001919067ffffffffffffffff166105628286612ff9565b52016104f9565b612ea3565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576105a5613299565b50604080516105b381610660565b73ffffffffffffffffffffffffffffffffffffffff60206bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283815201817f0000000000000000000000000000000000000000000000000000000000000000168152835192835251166020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761067c57604052565b610631565b6080810190811067ffffffffffffffff82111761067c57604052565b610160810190811067ffffffffffffffff82111761067c57604052565b60a0810190811067ffffffffffffffff82111761067c57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761067c57604052565b604051906107266040836106d6565b565b60405190610726610160836106d6565b604051906107266060836106d6565b604051906107266020836106d6565b67ffffffffffffffff811161067c5760051b60200190565b9080601f8301121561020957813561078581610756565b9261079360405194856106d6565b81845260208085019260051b82010192831161020957602001905b8282106107bb5750505090565b6020809183356107ca816101eb565b8152019101906107ae565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095761082490369060040161076e565b61082c614457565b60005b8151811015610907578061086473ffffffffffffffffffffffffffffffffffffffff61085d60019486612ff9565b5116615cc8565b6108bc575b60006108b573ffffffffffffffffffffffffffffffffffffffff61088d8487612ff9565b511673ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b550161082f565b73ffffffffffffffffffffffffffffffffffffffff6108db8285612ff9565b51167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91600080a2610869565b005b6024359063ffffffff8216820361020957565b6044359063ffffffff8216820361020957565b359063ffffffff8216820361020957565b8015150361020957565b359061072682610940565b81601f820112156102095780359061096c82610756565b9261097a60405194856106d6565b82845260208085019360061b8301019181831161020957602001925b8284106109a4575050505090565b60408483031261020957602060409182516109be81610660565b6109c787610299565b8152828701356109d6816101eb565b83820152815201930192610996565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095736602382011215610209578060040135610a3f81610756565b91610a4d60405193846106d6565b8183526024602084019260051b820101903682116102095760248101925b828410610a9c576024358567ffffffffffffffff821161020957610a96610907923690600401610955565b906132b2565b833567ffffffffffffffff811161020957820160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc82360301126102095760405190610ae882610660565b610af460248201610299565b8252604481013567ffffffffffffffff811161020957602491010136601f82011215610209578035610b2581610756565b91610b3360405193846106d6565b818352602060a081850193028201019036821161020957602001915b818310610b6e5750505091816020938480940152815201930192610a6b565b82360360a081126102095760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060405192610ba984610660565b8635610bb4816101eb565b845201126102095760a091602091604051610bce81610681565b610bd984880161092f565b8152610be76040880161092f565b84820152610bf76060880161092f565b60408201526080870135610c0a81610940565b606082015283820152815201920191610b4f565b67ffffffffffffffff811161067c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60405190610c676020836106d6565b60008252565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576104756040805190610cae81836106d6565b601382527f46656551756f74657220312e372e302d646576000000000000000000000000006020830152519182916020835260208301906102df565b906020808351928381520192019060005b818110610d085750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610cfb565b906020610d45928181520190610cea565b90565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610dd35761047585610dc7818703826106d6565b60405191829182610d34565b8254845260209093019260019283019201610db0565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff811161020957806004019060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261020957610e636144a2565b610e6d8280613592565b4263ffffffff1692915060005b81811061104757505060240190610e918284613592565b92905060005b838110610ea057005b80610ebf610eba600193610eb4868a613592565b90612ed2565b613646565b7fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e67ffffffffffffffff61100e610feb6020850194610fdd610f1d87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b610f4c610f28610717565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092168252565b63ffffffff8c166020820152610f87610f6d845167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b815160209092015160e01b7fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5167ffffffffffffffff1690565b93517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610e97565b8061106061105b600193610eb48980613592565b61360f565b7f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a73ffffffffffffffffffffffffffffffffffffffff61116d610feb60208501946111286110ca87517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b6110d5610f28610717565b63ffffffff8d166020820152610f87611102845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166000526005602052604060002090565b61114e611149825173ffffffffffffffffffffffffffffffffffffffff1690565b6144e6565b6111a6575b5173ffffffffffffffffffffffffffffffffffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9290921682524260208301529190931692a201610e7a565b836111c5825173ffffffffffffffffffffffffffffffffffffffff1690565b167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23600080a2611153565b9181601f840112156102095782359167ffffffffffffffff8311610209576020838186019501011161020957565b92610d459492611242928552151560208501526080604085015260808401906102df565b9160608184039101526102df565b346102095760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761128761026b565b60243590611294826101eb565b6044359160643567ffffffffffffffff8111610209576112b89036906004016111f0565b93909160843567ffffffffffffffff8111610209576112db9036906004016111f0565b9290917f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff8216146000146113c1575050935b6bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680861161139057509161138193916104759693614507565b9093916040519485948561121e565b857f6a92a4830000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b916113cb92612df6565b9361133a565b602060408183019282815284518094520192019060005b8181106113f55750505090565b9091926020604082611438600194885163ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b0194019291016113e8565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff8111610209576114929036906004016102ae565b61149b81610756565b916114a960405193846106d6565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06114d683610756565b0160005b81811061153157505060005b82811015611523576001906115076115028260051b8501612ee2565b614108565b6115118287612ff9565b5261151c8186612ff9565b50016114e6565b6040518061047586826113d1565b60209061153c613299565b828288010152016114da565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602061158d600435611588816101eb565b6136a5565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff60405191168152f35b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095767ffffffffffffffff6115f361026b565b6115fb613299565b5016600052600460205260406000206040519061161782610660565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c6020820152604051809161047582604081019263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565b90610140806107269361169784825115159052565b60208181015163ffffffff169085015260408181015163ffffffff169085015260608181015163ffffffff169085015260808181015160ff169085015260a0818101517fffffffff00000000000000000000000000000000000000000000000000000000169085015260c08181015161ffff169085015260e08181015163ffffffff16908501526101008181015163ffffffff16908501526101208181015161ffff1690850152015160ff16910152565b610160810192916107269190611682565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095767ffffffffffffffff61179961026b565b6117a16137a3565b5016600052600a6020526104756040600020611890611885604051926117c68461069d565b546117d560ff82165b15158552565b63ffffffff600882901c16602085015263ffffffff602882901c16604085015263ffffffff604882901c16606085015260ff606882901c1660808501527fffffffff00000000000000000000000000000000000000000000000000000000607082901b1660a085015261ffff609082901c1660c085015263ffffffff60a082901c1660e085015263ffffffff60c082901c1661010085015261ffff60e082901c1661012085015260f01c60ff1690565b60ff16610140830152565b60405191829182611748565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760005473ffffffffffffffffffffffffffffffffffffffff8116330361195b577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b6107269092919260808101936060809163ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957610475611a57611a0561026b565b67ffffffffffffffff60243591611a1b836101eb565b611a23613823565b5016600052600b60205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b60ff60405191611a6683610681565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015260405191829182611985565b90611ab890604083526040830190610479565b9060208183039101526020808351928381520192019060005b818110611ade5750505090565b909192602061016082611af46001948851611682565b019401929101611ad1565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957600854611b3a81610756565b90611b4860405192836106d6565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611b7582610756565b0160005b818110611c2d575050611b8b8161324a565b9060005b818110611ba757505061047560405192839283611aa5565b80611bdf611bc6611bb9600194615b8f565b67ffffffffffffffff1690565b611bd08387612ff9565b9067ffffffffffffffff169052565b611c11611c0c611bf2610fdd8488612ff9565b67ffffffffffffffff16600052600a602052604060002090565b6137f5565b611c1b8287612ff9565b52611c268186612ff9565b5001611b8f565b602090611c386137a3565b82828701015201611b79565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102095760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957611ccd61026b565b611cd5610909565b611cdd61091c565b9060643591611ceb836101eb565b611d0c611c0c8567ffffffffffffffff16600052600a602052604060002090565b91611d1e611d1a8451151590565b1590565b611fc557611d4c90611d46611d40611d3a608087015160ff1690565b60ff1690565b84613848565b90613865565b93611d6a611d61604085015163ffffffff1690565b63ffffffff1690565b9163ffffffff8616928311611f9b57602084015163ffffffff1663ffffffff811663ffffffff831611611f62575050611dbf611dba8267ffffffffffffffff166000526004602052604060002090565b61366b565b9063ffffffff611dd6602084015163ffffffff1690565b1615611f27575091611e45611e38611e3360ff9794611e2d611e1a611e1a61047599517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b6dffffffffffffffffffffffffffff1690565b90612daa565b61387f565b662386f26fc10000900490565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861614600014611efe57611ed1611ecb6101407bffffffffffffffffffffffffffffffffffffffffffffffffffffffff93015160ff1690565b956136a5565b16906040519586951692859094939260609263ffffffff6080840197168352602083015260408201520152565b507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff611ed16064956136a5565b7fa96740690000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f869337890000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f4c4fc93a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff851660045260246000fd5b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff81116102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126102095760405161207681610660565b816004013567ffffffffffffffff81116102095761209a906004369185010161076e565b8152602482013567ffffffffffffffff8111610209576109079260046120c3923692010161076e565b60208201526138ca565b359060ff8216820361020957565b35907fffffffff000000000000000000000000000000000000000000000000000000008216820361020957565b359061ffff8216820361020957565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760043567ffffffffffffffff811161020957366023820112156102095780600401359061217282610756565b9061218060405192836106d6565b82825260246101806020840194028201019036821161020957602401925b8184106121ae5761090783613a55565b8336036101808112610209576101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0604051926121eb84610660565b6121f488610299565b84520112610209576101809160209161220b610728565b61221684890161094a565b81526122246040890161092f565b848201526122346060890161092f565b60408201526122456080890161092f565b606082015261225660a089016120cd565b608082015261226760c089016120db565b60a082015261227860e08901612108565b60c082015261228a610100890161092f565b60e082015261229c610120890161092f565b6101008201526122af6101408901612108565b6101208201526122c261016089016120cd565b6101408201528382015281520193019261219e565b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957606063ffffffff8061232b61231961026b565b60243590612326826101eb565b613c22565b9193908160405195168552166020840152166040820152f35b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761237b61026b565b6024359067ffffffffffffffff8211610209576123ac610475916123a66123c89436906004016111f0565b91613d29565b63ffffffff6040949394519586956060875260608701906102df565b9216602085015283820360408501526102df565b906123ef90604083526040830190610cea565b9060208183039101526020808351928381520192019060005b8181106124155750505090565b909192602060808261245c60019488516060809163ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b019401929101612408565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761249e61026b565b6006546124aa81610756565b906124b860405192836106d6565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06124e582610756565b013660208401376124f581610756565b9261250360405194856106d6565b8184527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061253083610756565b0160005b81811061261757505060005b82811061255957505050610475604051928392836123dc565b806125fb6125f661258a612571612571600196615bc4565b73ffffffffffffffffffffffffffffffffffffffff1690565b6125b381612598868b612ff9565b9073ffffffffffffffffffffffffffffffffffffffff169052565b6125d18667ffffffffffffffff16600052600b602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b612fb4565b6126058288612ff9565b526126108187612ff9565b5001612540565b602090612622613823565b82828901015201612534565b346102095760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095760405180602060065491828152019060066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9060005b8181106126ad5761047585610dc7818703826106d6565b8254845260209093019260019283019201612696565b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610209576040612703600435611502816101eb565b61273d8251809263ffffffff602080927bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8151168552015116910152565bf35b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095761277661026b565b6024359067ffffffffffffffff821161020957816004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610209576127db611c0c8367ffffffffffffffff16600052600a602052604060002090565b916127e9611d1a8451151590565b612ae657606482019061282d611d1a61280184612ee2565b73ffffffffffffffffffffffffffffffffffffffff166000526001600601602052604060002054151590565b612a9857908161283f8686809561520c565b9561284c61158884612ee2565b946044810160008061285e8386613592565b159050612a6957505061ffff86856128a79361289e61289760e061288b60c06128b19d9e015161ffff1690565b95015163ffffffff1690565b9188613592565b94909316615920565b9096909590612ee2565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146000146129f8576101408801516129059060ff16612d1b565b61290e91612daa565b92611e2d611e1a611e1a6129b961299e6104759d6129996129e89e612999611d617bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9f6129e09f60606129999f6129db9f6129949363ffffffff61297961288b946024612983955b0190612f2c565b92905016906138bd565b611e2d611d3a60808a015160ff1690565b613865565b6138bd565b9467ffffffffffffffff166000526004602052604060002090565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b612d42565b911690612dbd565b6040519081529081906020820190565b612a0190612d42565b92611e2d611e1a611e1a6129b961299e6104759d6129996129e89e612999611d617bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9f6129e09f60606129999f6129db9f6129949363ffffffff61297961288b94602461298395612972565b95915095612a92612a8d612a866101206128b194015161ffff1690565b61ffff1690565b612d1b565b91612ee2565b611f5e612aa483612ee2565b7f2502348c0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b7f99ac52f20000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346102095760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102095773ffffffffffffffffffffffffffffffffffffffff600435612b6d816101eb565b612b75614457565b16338114612be757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102095760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020957600435612c4c816101eb565b67ffffffffffffffff612c5d610282565b169081600052600a60205260ff6040600020541615612cbe57612c7f906136a5565b6000918252600460209081526040928390205483517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152921690820152f35b507f99ac52f20000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90662386f26fc10000820291808304662386f26fc100001490151715612d3d57565b612cec565b90670de0b6b3a7640000820291808304670de0b6b3a76400001490151715612d3d57565b908160051b9180830460201490151715612d3d57565b9061012c82029180830461012c1490151715612d3d57565b90606c820291808304606c1490151715612d3d57565b81810292918115918404141715612d3d57565b8115612dc7570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b612e35612e2f610d4594937bffffffffffffffffffffffffffffffffffffffffffffffffffffffff612e2881956136a5565b1690612daa565b926136a5565b1690612dbd565b90612e4682610756565b612e5360405191826106d6565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612e818294610756565b019060005b828110612e9257505050565b806060602080938501015201612e86565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156105695760061b0190565b35610d45816101eb565b91908110156105695760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610209570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610209570180359067ffffffffffffffff82116102095760200191813603831361020957565b929192612f8982610c1e565b91612f9760405193846106d6565b829481845281830111610209578281602093846000960137010152565b90604051612fc181610681565b606060ff82945463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c166040850152821c161515910152565b80518210156105695760209160051b010190565b90929161305a6130318367ffffffffffffffff16600052600a602052604060002090565b5460701b7fffffffff000000000000000000000000000000000000000000000000000000001690565b9061306481612e3c565b9560005b828110613079575050505050505090565b61308c613087828489612ed2565b612ee2565b83886130a661309c858484612eec565b6040810190612f2c565b9050602081116131c2575b5083926130e76130e16130da6130d0600198613105976125f697612eec565b6020810190612f2c565b3691612f7d565b8961416f565b6125d18967ffffffffffffffff16600052600b602052604060002090565b6060810151156131885761316c613126602061314093015163ffffffff1690565b6040805163ffffffff909216602083015290928391820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826106d6565b613176828b612ff9565b52613181818a612ff9565b5001613068565b5061314061316c6131bd6131b08967ffffffffffffffff16600052600a602052604060002090565b5460a01c63ffffffff1690565b613126565b9150506131fa611d616131ed846125d18b67ffffffffffffffff16600052600b602052604060002090565b5460401c63ffffffff1690565b10613207578388386130b1565b7f36f536ca0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b9061325482610756565b61326160405191826106d6565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061328f8294610756565b0190602036910137565b604051906132a682610660565b60006020838281520152565b906132bb614457565b6000915b80518310156134c4576132d28382612ff9565b51906132e6825167ffffffffffffffff1690565b946020600093019367ffffffffffffffff8716935b855180518210156134af5761331282602092612ff9565b51015161333e613323838951612ff9565b515173ffffffffffffffffffffffffffffffffffffffff1690565b604082015163ffffffff1660208110613461575090867f5c55501634b3b87e45686082d77f017b6639b436c21cb423ba6313d843f66ed173ffffffffffffffffffffffffffffffffffffffff8461343d858f600199986125d16133b59267ffffffffffffffff16600052600b602052604060002090565b815181546020808501516040808701516060978801517fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951663ffffffff96909616959095179190921b67ffffffff00000000161792901b6bffffffff0000000000000000169190911790151590921b6cff00000000000000000000000016919091179055565b61344684615e08565b50613458604051928392169582611985565b0390a3016132fb565b7f24ecdc020000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff90911660045263ffffffff1660245260446000fd5b505095509250926001915001919290926132bf565b50905060005b815181101561358e57806134f26134e360019385612ff9565b515167ffffffffffffffff1690565b67ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff61353b602061351f8689612ff9565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b600061355f826125d18767ffffffffffffffff16600052600b602052604060002090565b551691167f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b600080a3016134ca565b5050565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610209570180359067ffffffffffffffff821161020957602001918160061b3603831361020957565b35907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216820361020957565b6040813603126102095761363e60206040519261362b84610660565b8035613636816101eb565b8452016135e6565b602082015290565b6040813603126102095761363e60206040519261366284610660565b61363681610299565b9060405161367881610660565b91547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116835260e01c6020830152565b73ffffffffffffffffffffffffffffffffffffffff81166000526005602052604060002090604051916136d783610660565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811680845260e09190911c602084018190521590811561377c575b506137385750517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff907f06439c6b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1615905038613710565b604051906137b08261069d565b6000610140838281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e082015282610100820152826101208201520152565b906107266040516138058161069d565b61014061381b8295546117d56117cf8260ff1690565b60ff16910152565b6040519061383082610681565b60006060838281528260208201528260408201520152565b9063ffffffff8091169116029063ffffffff8216918203612d3d57565b9063ffffffff8091169116019063ffffffff8211612d3d57565b90662386f26fc0ffff8201809211612d3d57565b9060028201809211612d3d57565b9060208201809211612d3d57565b9060018201809211612d3d57565b91908201809211612d3d57565b6138d2614457565b60208101519160005b835181101561397157806138f461115360019387612ff9565b61391b61391673ffffffffffffffffffffffffffffffffffffffff8316612571565b616043565b613927575b50016138db565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613920565b5091505160005b815181101561358e5761398e6111538284612ff9565b9073ffffffffffffffffffffffffffffffffffffffff821615613a2b577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613a22836139fa6139f561257160019773ffffffffffffffffffffffffffffffffffffffff1690565b615e43565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101613978565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b90613a5e614457565b60005b8251811015613c1d57613a748184612ff9565b516020613a846134e38487612ff9565b91015167ffffffffffffffff82169182158015613bfe575b8015613bd0575b613b995791613b55613b5a92613b5085613aff613ada61303160019a9967ffffffffffffffff16600052600a602052604060002090565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b613b6157847f4efe320c85221c7c3684c54561bea5a9c4dcfad794c6ef9ff9e6b43fb307c0f860405180613b338782611748565b0390a267ffffffffffffffff16600052600a602052604060002090565b614803565b615e08565b5001613a61565b847f0c6380a4766d45f5d53ca170bf865bebfab44958dec379d5a90177264e6645b760405180613b918782611748565b0390a2611bf2565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5061010082015163ffffffff1663ffffffff613bf6611d61604086015163ffffffff1690565b911611613aa3565b5063ffffffff613c1661010084015163ffffffff1690565b1615613a9c565b509050565b9190613c659067ffffffffffffffff8416600052600b60205260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b9160405192613c7384610681565b549263ffffffff8416808252613cb7606063ffffffff8760201c169384602082015260ff63ffffffff8960401c1698896040840152831c1615159182910152151590565b613d02575050613cdd91925067ffffffffffffffff16600052600a602052604060002090565b549061ffff613cf9609084901c82169360a01c63ffffffff1690565b92169190602090565b613d2391939250613d19613d199163ffffffff1690565b9363ffffffff1690565b91929190565b611c0c613d4d9194929467ffffffffffffffff16600052600a602052604060002090565b9260a08401927fffffffff00000000000000000000000000000000000000000000000000000000613d9e85517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c0000000000000000000000000000000000000000000000000000000081149081156140de575b81156140b4575b50614060577f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613e4286517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613fc7577fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613eb486517fffffffff000000000000000000000000000000000000000000000000000000001690565b1614613f3557611f5e613ee785517fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b613f759350613f52611d616040613f589597015163ffffffff1690565b91615033565b91613fa16040840151604051938491602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836106d6565b613d236060613fb4855163ffffffff1690565b940151613fbf610747565b908152615154565b613f759350613fe4611d616040613fea9597015163ffffffff1690565b91614d8b565b916140076060840151604051938491602083019190602083019252565b613d23614018845163ffffffff1690565b936140336020608083015192015167ffffffffffffffff1690565b9061405661403f610738565b6000815267ffffffffffffffff9093166020840152565b6040820152614f87565b6140a19350936140979294614091611d61604061408561010086015163ffffffff1690565b94015163ffffffff1690565b92614bc7565b5163ffffffff1690565b906140aa610c58565b9190610d45610c58565b7f647e2ba90000000000000000000000000000000000000000000000000000000091501438613dd1565b7fac77ffec0000000000000000000000000000000000000000000000000000000081149150613dca565b73ffffffffffffffffffffffffffffffffffffffff90614126613299565b5016600052600560205260406000206040519061414282610660565b547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825260e01c602082015290565b7fffffffff000000000000000000000000000000000000000000000000000000001691907f2812d52c0000000000000000000000000000000000000000000000000000000083146142b7577f1e10bdc40000000000000000000000000000000000000000000000000000000083146142a9577fac77ffec00000000000000000000000000000000000000000000000000000000831461429e577f647e2ba9000000000000000000000000000000000000000000000000000000008314614293577fc4e0595300000000000000000000000000000000000000000000000000000000831461428457827f2ee820750000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61072691925061dee990615ac3565b610726919250615b24565b610726919250615a60565b610726919250600190615ac3565b6107269192506159d0565b917fffffffff0000000000000000000000000000000000000000000000000000000083167f2812d52c00000000000000000000000000000000000000000000000000000000811461444b577f1e10bdc400000000000000000000000000000000000000000000000000000000811461442b577fac77ffec00000000000000000000000000000000000000000000000000000000811461441f577f647e2ba9000000000000000000000000000000000000000000000000000000008114614413577fc4e0595300000000000000000000000000000000000000000000000000000000146143f8577f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000831660045260246000fd5b61072692501561440b5761dee990615ac3565b600090615ac3565b50506107269150615b24565b50506107269150615a60565b506107269250156144425760ff60015b1690615ac3565b60ff600061443b565b505061072691506159d0565b73ffffffffffffffffffffffffffffffffffffffff60015416330361447857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b336000526003602052604060002054156144b857565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff610d4591166006615e78565b611c0c61452e9196949395929667ffffffffffffffff16600052600a602052604060002090565b9460a08601947fffffffff0000000000000000000000000000000000000000000000000000000061457f87517fffffffff000000000000000000000000000000000000000000000000000000001690565b167f2812d52c0000000000000000000000000000000000000000000000000000000081149081156147d9575b81156147af575b5061476a5750507fc4e05953000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061462586517fffffffff000000000000000000000000000000000000000000000000000000001690565b161461473f577f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061469786517fffffffff000000000000000000000000000000000000000000000000000000001690565b16146146ca57611f5e613ee785517fffffffff000000000000000000000000000000000000000000000000000000001690565b61470c93506130da60606146f66146ef611d616040614738989a015163ffffffff1690565b8486614d8b565b0151604051958691602083019190602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018652856106d6565b9160019190565b61470c93506130da60406146f6614763611d6183614738989a015163ffffffff1690565b8486615033565b945094916147909161478a611d61610100610d4596015163ffffffff1690565b91615ed8565b936147a7602061479f87615ff9565b960151151590565b933691612f7d565b7f647e2ba900000000000000000000000000000000000000000000000000000000915014386145b2565b7fac77ffec00000000000000000000000000000000000000000000000000000000811491506145ab565b614b786101406107269361484b61481a8251151590565b859060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b61489561485f602083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ff1660089190911b64ffffffff0016178555565b6148e36148a9604083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1660289190911b68ffffffff000000000016178555565b6149356148f7606083015163ffffffff1690565b85547fffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffff1660489190911b6cffffffff00000000000000000016178555565b614985614946608083015160ff1690565b85547fffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffff1660689190911b6dff0000000000000000000000000016178555565b6149f86149b560a08301517fffffffff000000000000000000000000000000000000000000000000000000001690565b85547fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1660709190911c71ffffffff000000000000000000000000000016178555565b614a4f614a0a60c083015161ffff1690565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b614aac614a6360e083015163ffffffff1690565b85547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b77ffffffff000000000000000000000000000000000000000016178555565b614b0e614ac161010083015163ffffffff1690565b85547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff1660c09190911b7bffffffff00000000000000000000000000000000000000000000000016178555565b614b70614b2161012083015161ffff1690565b85547fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7dffff0000000000000000000000000000000000000000000000000000000016178555565b015160ff1690565b7fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff00000000000000000000000000000000000000000000000000000000000083549260f01b169116179055565b9063ffffffff614be193614bd9613299565b501691615ed8565b90815111611f9b5790565b906004116102095790600490565b909291928360041161020957831161020957600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110614c69575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9080601f83011215610209578135614cb281610756565b92614cc060405194856106d6565b81845260208085019260051b82010192831161020957602001905b828210614ce85750505090565b8135815260209182019101614cdb565b6020818303126102095780359067ffffffffffffffff8211610209570160a0818303126102095760405191614d2c836106ba565b614d358261092f565b8352614d4360208301610299565b60208401526040820135614d5681610940565b604084015260608201356060840152608082013567ffffffffffffffff811161020957614d839201614c9b565b608082015290565b60606080604051614d9b816106ba565b60008152600060208201526000604082015260008382015201528115614e9d577f1f3b3aba000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000614e10614e0a8585614bec565b90614c35565b1603614e735781614e2c92614e2492614bfa565b810190614cf8565b9063ffffffff614e40835163ffffffff1690565b1611614e495790565b7f2e2b0c290000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fb00b53dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b805160209091019060005b818110614edf5750505090565b8251845260209384019390920191600101614ed2565b90610d4594937fffffffffffffffff000000000000000000000000000000000000000000000000600e947fff0000000000000000000000000000000000000000000000000000000000000080947f1a2b3c4d00000000000000000000000000000000000000000000000000000000875260f81b16600486015260c01b16600584015260f81b16600d8201520190614ec7565b604081019081515160ff8111614fff578151916003831015614fd0576020015160ff93610d45936131409267ffffffffffffffff16915191604051968795169160208601614ef5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600d600452600060245260446000fd5b9060608060405161504381610681565b60008152600060208201526000604082015201528015614e9d577f21ea4ca9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006150ac614e0a8486614bec565b1603614e7357806150bc92614bfa565b819291016000926020818303126151505780359067ffffffffffffffff821161514c57019260808483031261514957604051936150f885610681565b80358552602081013561510a81610940565b60208601526040810135604086015260608101359167ffffffffffffffff8311615149575061513a929101614c9b565b6060830152815111611f9b5790565b80fd5b8480fd5b8380fd5b8051519060ff82116151c9577fff0000000000000000000000000000000000000000000000000000000000000091613140610d4592516040519485937f5e6f7a8b00000000000000000000000000000000000000000000000000000000602086015260f81b1660248401526025830190614ec7565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600e600452600060245260446000fd5b90816020910312610209573590565b919061521b6020830183612f2c565b9390506040830161522c8185613592565b90506020840191615244611d61845163ffffffff1690565b8088116158ee5750600182116158ba5760a085019661528388517fffffffff000000000000000000000000000000000000000000000000000000001690565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f2812d52c0000000000000000000000000000000000000000000000000000000081148015615891575b8015615868575b156153555750505050505050918161534f6130da61531e615348966153026080610d45980186612f2c565b614091611d616040614085610100879697015163ffffffff1690565b51958694517fffffffff000000000000000000000000000000000000000000000000000000001690565b9280612f2c565b906142c2565b7fc4e0595300000000000000000000000000000000000000000000000000000000819b9a939495979996989b146000146155f25750506153f26153b96153e5999a6040613f52611d616153ab60808b018b612f2c565b939094015163ffffffff1690565b918251998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b61534f6130da8880612f2c565b606081015151908261540f6154078780612f2c565b8101906151fd565b6155d557508161559f575b8515159081615592575b5061556857604081116155365750615449906154438594939795612d94565b906138bd565b946000935b8385106154a7575050505050611d6161546b915163ffffffff1690565b80821161547757505090565b7f869337890000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b909192939560019061550b611d616131ed6154d68667ffffffffffffffff16600052600b602052604060002090565b6154e76130878d610eb48b8d613592565b73ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b80156155265761551a916138bd565b965b019392919061544e565b50615530906138a1565b9661551c565b7fc327a56c00000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f5bed51920000000000000000000000000000000000000000000000000000000060005260046000fd5b6040915001511538615424565b611f5e827fc327a56c00000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b6155ec9193506154436155e7846138af565b612d66565b9161541a565b7f1e10bdc4000000000000000000000000000000000000000000000000000000000361581957506156746156396153e5999a6040613fe4611d616153ab60808b018b612f2c565b9161564b611d61845163ffffffff1690565b998a91517fffffffff000000000000000000000000000000000000000000000000000000001690565b60808101515190826156896154078780612f2c565b6158015750816157cb575b851515806157bf575b615568576040821161578b576020015167ffffffffffffffff9081169081831c166157515750506156d5906154438594939795612d7c565b946000935b8385106156f7575050505050611d6161546b915163ffffffff1690565b9091929395600190615726611d616131ed6154d68667ffffffffffffffff16600052600b602052604060002090565b801561574157615735916138bd565b965b01939291906156da565b5061574b906138a1565b96615737565b7fafa933080000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260245260446000fd5b7f8a0d71f7000000000000000000000000000000000000000000000000000000006000526004829052604060245260446000fd5b5060608101511561569d565b611f5e827f8a0d71f700000000000000000000000000000000000000000000000000000000600052906044916004526000602452565b6158139193506154436155e784613893565b91615694565b7f2ee82075000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b507f647e2ba90000000000000000000000000000000000000000000000000000000081146152d7565b507fac77ffec0000000000000000000000000000000000000000000000000000000081146152d0565b7fd88dddd6000000000000000000000000000000000000000000000000000000006000526004829052600160245260446000fd5b7f8693378900000000000000000000000000000000000000000000000000000000600052600452602487905260446000fd5b94939192909282156159b05767ffffffffffffffff16600052600b602052604060002091156105695761595c916125f69135906125d1826101eb565b9261596d611d1a6060860151151590565b61599e575050615987612a8d611d61845163ffffffff1690565b90613d236040614085602086015163ffffffff1690565b6159a9919350612d1b565b9190602090565b505050509050600090600090600090565b90816020910312610209575190565b6020815103615a13576159ec60208251830101602083016159c1565b73ffffffffffffffffffffffffffffffffffffffff8111908115615a54575b50615a135750565b615a50906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b0390fd5b61040091501038615a0b565b6020815103615a8657600b615a7e60208351840101602084016159c1565b10615a865750565b615a50906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b906020825103615ae95780615ad6575050565b615a7e60208351840101602084016159c1565b6040517fe0d7fb020000000000000000000000000000000000000000000000000000000081526020600482015280615a5060248201856102df565b6024815103615b3a57602281015115615b3a5750565b615a50906040519182917f373b0e4400000000000000000000000000000000000000000000000000000000835260206004840181815201906102df565b80548210156105695760005260206000200190600090565b6008548110156105695760086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3015490565b6006548110156105695760066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015490565b91615c31918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b80548015615c99577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615c6a8282615b77565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260076020526040902054908115615da5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211612d3d57600654927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612d3d578383600095615d649503615d6a575b505050615d536006615c35565b600790600052602052604060002090565b55600190565b615d53615d9691615d8c615d82615d9c956006615b77565b90549060031b1c90565b9283916006615b77565b90615bf9565b55388080615d46565b5050600090565b8054906801000000000000000082101561067c5781615dd3916001615c3194018155615b77565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b600081815260096020526040902054615e3d57615e26816008615dac565b600854906000526009602052604060002055600190565b50600090565b600081815260036020526040902054615e3d57615e61816002615dac565b600254906000526003602052604060002055600190565b6000828152600182016020526040902054615da55780615e9a83600193615dac565b80549260005201602052604060002055600190565b9081604091031261020957602060405191615ec983610660565b80518352015161363e81610940565b91615ee1613299565b508115615fd75750615f226130da8280615f1c7fffffffff000000000000000000000000000000000000000000000000000000009587614c35565b95614bfa565b91167f181dcf10000000000000000000000000000000000000000000000000000000008103615f5f575080602080610d4593518301019101615eaf565b7f97a657c90000000000000000000000000000000000000000000000000000000014615faf577f5247fdce0000000000000000000000000000000000000000000000000000000060005260046000fd5b80602080615fc2935183010191016159c1565b615fca610717565b9081526000602082015290565b91505067ffffffffffffffff615feb610717565b911681526000602082015290565b6020604051917f181dcf1000000000000000000000000000000000000000000000000000000000828401528051602484015201511515604482015260448152610d456064826106d6565b600081815260036020526040902054908115615da5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211612d3d57600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612d3d578383615d6494600096036160df575b5050506160ce6002615c35565b600390600052602052604060002090565b6160ce615d96916160f7615d82616101956002615b77565b9283916002615b77565b553880806160c156fea164736f6c634300081a000a",
}

var FeeQuoterABI = FeeQuoterMetaData.ABI

var FeeQuoterBin = FeeQuoterMetaData.Bin

func DeployFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *types.Transaction, *FeeQuoter, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterBin), backend, staticConfig, priceUpdaters, tokenTransferFeeConfigArgs, destChainConfigArgs)
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

func (_FeeQuoter *FeeQuoterCaller) GetAllDestChainConfigs(opts *bind.CallOpts) ([]uint64, []FeeQuoterDestChainConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getAllDestChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]FeeQuoterDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]FeeQuoterDestChainConfig)).(*[]FeeQuoterDestChainConfig)

	return out0, out1, err

}

func (_FeeQuoter *FeeQuoterSession) GetAllDestChainConfigs() ([]uint64, []FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetAllDestChainConfigs(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetAllDestChainConfigs() ([]uint64, []FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetAllDestChainConfigs(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetAllDestChainSelectors(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getAllDestChainSelectors")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetAllDestChainSelectors() ([]uint64, error) {
	return _FeeQuoter.Contract.GetAllDestChainSelectors(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetAllDestChainSelectors() ([]uint64, error) {
	return _FeeQuoter.Contract.GetAllDestChainSelectors(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetAllTokenTransferFeeConfigs(opts *bind.CallOpts, destChainSelector uint64) (GetAllTokenTransferFeeConfigs,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getAllTokenTransferFeeConfigs", destChainSelector)

	outstruct := new(GetAllTokenTransferFeeConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TokenTransferFeeConfigs = *abi.ConvertType(out[1], new([]FeeQuoterTokenTransferFeeConfig)).(*[]FeeQuoterTokenTransferFeeConfig)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) GetAllTokenTransferFeeConfigs(destChainSelector uint64) (GetAllTokenTransferFeeConfigs,

	error) {
	return _FeeQuoter.Contract.GetAllTokenTransferFeeConfigs(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetAllTokenTransferFeeConfigs(destChainSelector uint64) (GetAllTokenTransferFeeConfigs,

	error) {
	return _FeeQuoter.Contract.GetAllTokenTransferFeeConfigs(&_FeeQuoter.CallOpts, destChainSelector)
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

func (_FeeQuoter *FeeQuoterCaller) GetTokenTransferFee(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (GetTokenTransferFee,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenTransferFee", destChainSelector, token)

	outstruct := new(GetTokenTransferFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenTransferFee(destChainSelector uint64, token common.Address) (GetTokenTransferFee,

	error) {
	return _FeeQuoter.Contract.GetTokenTransferFee(&_FeeQuoter.CallOpts, destChainSelector, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenTransferFee(destChainSelector uint64, token common.Address) (GetTokenTransferFee,

	error) {
	return _FeeQuoter.Contract.GetTokenTransferFee(&_FeeQuoter.CallOpts, destChainSelector, token)
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

func (_FeeQuoter *FeeQuoterCaller) QuoteGasForExec(opts *bind.CallOpts, destChainSelector uint64, nonCalldataGas uint32, calldataSize uint32, feeToken common.Address) (QuoteGasForExec,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "quoteGasForExec", destChainSelector, nonCalldataGas, calldataSize, feeToken)

	outstruct := new(QuoteGasForExec)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalGas = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.GasCostInUsdCents = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeTokenPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PremiumPercentMultiplier = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) QuoteGasForExec(destChainSelector uint64, nonCalldataGas uint32, calldataSize uint32, feeToken common.Address) (QuoteGasForExec,

	error) {
	return _FeeQuoter.Contract.QuoteGasForExec(&_FeeQuoter.CallOpts, destChainSelector, nonCalldataGas, calldataSize, feeToken)
}

func (_FeeQuoter *FeeQuoterCallerSession) QuoteGasForExec(destChainSelector uint64, nonCalldataGas uint32, calldataSize uint32, feeToken common.Address) (QuoteGasForExec,

	error) {
	return _FeeQuoter.Contract.QuoteGasForExec(&_FeeQuoter.CallOpts, destChainSelector, nonCalldataGas, calldataSize, feeToken)
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

func (_FeeQuoter *FeeQuoterTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactor) RemoveFeeTokens(opts *bind.TransactOpts, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "removeFeeTokens", feeTokensToRemove)
}

func (_FeeQuoter *FeeQuoterSession) RemoveFeeTokens(feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.RemoveFeeTokens(&_FeeQuoter.TransactOpts, feeTokensToRemove)
}

func (_FeeQuoter *FeeQuoterTransactorSession) RemoveFeeTokens(feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.RemoveFeeTokens(&_FeeQuoter.TransactOpts, feeTokensToRemove)
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

type GetAllTokenTransferFeeConfigs struct {
	Tokens                  []common.Address
	TokenTransferFeeConfigs []FeeQuoterTokenTransferFeeConfig
}
type GetTokenAndGasPrices struct {
	TokenPrice    *big.Int
	GasPriceValue *big.Int
}
type GetTokenTransferFee struct {
	FeeUSDCents       uint32
	DestGasOverhead   uint32
	DestBytesOverhead uint32
}
type ProcessMessageArgs struct {
	MsgFeeJuels           *big.Int
	IsOutOfOrderExecution bool
	ConvertedExtraArgs    []byte
	TokenReceiver         []byte
}
type QuoteGasForExec struct {
	TotalGas                 uint32
	GasCostInUsdCents        *big.Int
	FeeTokenPrice            *big.Int
	PremiumPercentMultiplier *big.Int
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
	return common.HexToHash("0x4efe320c85221c7c3684c54561bea5a9c4dcfad794c6ef9ff9e6b43fb307c0f8")
}

func (FeeQuoterDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x0c6380a4766d45f5d53ca170bf865bebfab44958dec379d5a90177264e6645b7")
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

	GetAllDestChainConfigs(opts *bind.CallOpts) ([]uint64, []FeeQuoterDestChainConfig, error)

	GetAllDestChainSelectors(opts *bind.CallOpts) ([]uint64, error)

	GetAllTokenTransferFeeConfigs(opts *bind.CallOpts, destChainSelector uint64) (GetAllTokenTransferFeeConfigs,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error)

	GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error)

	GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error)

	GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

		error)

	GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error)

	GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error)

	GetTokenTransferFee(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (GetTokenTransferFee,

		error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error)

	GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, messageReceiver []byte) (ProcessMessageArgs,

		error)

	ProcessPoolReturnData(opts *bind.CallOpts, destChainSelector uint64, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) ([][]byte, error)

	QuoteGasForExec(opts *bind.CallOpts, destChainSelector uint64, nonCalldataGas uint32, calldataSize uint32, feeToken common.Address) (QuoteGasForExec,

		error)

	ResolveLegacyArgs(opts *bind.CallOpts, destChainSelector uint64, extraArgs []byte) (ResolveLegacyArgs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	RemoveFeeTokens(opts *bind.TransactOpts, feeTokensToRemove []common.Address) (*types.Transaction, error)

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
