// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp_over_superchain_interop

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type InternalAny2EVMRampMessage struct {
	Header       InternalRampMessageHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     *big.Int
	TokenAmounts []InternalAny2EVMTokenTransfer
}

type InternalAny2EVMTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	DestGasAmount     uint32
	ExtraData         []byte
	Amount            *big.Int
}

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenAmounts   []InternalEVM2AnyTokenTransfer
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type OnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	MessageInterceptor     common.Address
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

var OnRampOverSuperchainInteropMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extractGasLimit\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMessageId\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashInteropMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reemitInteropMessage\",\"inputs\":[{\"name\":\"interopMessage\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPSuperchainMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgsTooShort\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceChainSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageDoesNotExist\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MessageIdUnexpectedlySet\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6101006040523461056d576153ff8038038061001a816105a7565b92833981019080820390610140821261056d576080821261056d5761003d610588565b90610047816105cc565b82526020810151906001600160a01b038216820361056d5760208301918252610072604082016105e0565b926040810193845260a0610088606084016105e0565b6060830190815295607f19011261056d5760405160a081016001600160401b03811182821017610572576040526100c1608084016105e0565b81526100cf60a084016105f4565b602082019081526100e260c085016105e0565b91604081019283526100f660e086016105e0565b936060820194855261010b61010087016105e0565b6080830190815261012087015190966001600160401b03821161056d57018a601f8201121561056d578051906001600160401b0382116105725760209b8c6060610159828660051b016105a7565b9e8f8681520194028301019181831161056d57602001925b8284106104ff575050505033156104ee57600180546001600160a01b0319163317905580516001600160401b03161580156104dc575b80156104ca575b80156104b8575b61048b57516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e081905282519091161580156104a6575b801561049c575b61048b57815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19861012098606061029f610588565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a16000905b805182101561040b5761032b8282610601565b51916001600160401b0361033f8284610601565b5151169283156103f65760008481526006602090815260409182902081840151815494840151600160401b600160e81b03198616604883901b600160481b600160e81b031617901515851b68ff000000000000000016179182905583516001600160401b0390951685526001600160a01b031691840191909152811c60ff1615159082015291926001927fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590606090a20190610318565b8363c35aa79d60e01b60005260045260246000fd5b604051614dd3908161062c823960805181818161023b0152818161038401528181610be801528181611ce7015281816131060152613a68015260a051818181610274015281816106e90152610c21015260c05181818161029c01528181610c5d01526120fc015260e0518181816102c401528181610c990152612f600152f35b6306b7c75960e31b60005260046000fd5b5082511515610200565b5084516001600160a01b0316156101f9565b5088516001600160a01b0316156101b5565b5087516001600160a01b0316156101ae565b5086516001600160a01b0316156101a7565b639b15e16f60e01b60005260046000fd5b60608483031261056d5760405190606082016001600160401b038111838210176105725760405261052f856105cc565b82526020850151906001600160a01b038216820361056d578260209283606095015261055d604088016105f4565b6040820152815201930192610171565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761057257604052565b6040519190601f01601f191682016001600160401b0381118382101761057257604052565b51906001600160401b038216820361056d57565b51906001600160a01b038216820361056d57565b5190811515820361056d57565b80518210156106155760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146101675780630be49dbf14610162578063181f5a771461015d57806320487ded146101585780632716072b1461015357806327e936f11461014e57806328938e371461014957806348a98aa4146101445780635cb80c5d1461013f5780636def4ce71461013a5780637437ff9f1461013557806379ba5097146101305780638da5cb5b1461012b5780639041be3d14610126578063972b461214610121578063a8cfe8141461011c578063c9b146b314610117578063df0aa9e914610112578063f2890a211461010d578063f2fde38b146101085763fbca3b741461010357600080fd5b6126f8565b6125f3565b612471565b611b09565b6117d3565b61177f565b611462565b611376565b611324565b61123b565b611134565b611049565b610f0b565b610e20565b610dbf565b610acc565b6109d2565b610639565b610575565b6102f9565b6101d1565b600091031261017757565b600080fd5b6101cf90929192608081019373ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b565b346101775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017757600060606040516102108161089f565b82815282602082015282604082015201526102f56040516102308161089f565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660208301527f0000000000000000000000000000000000000000000000000000000000000000811660408301527f00000000000000000000000000000000000000000000000000000000000000001660608201526040519182918261017c565b0390f35b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043567ffffffffffffffff811161017757806004016101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261017757602482016103788161275c565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036104b957506103be6103b936836116d0565b6130a2565b604483019060646103ce8361275c565b9401936103da8561275c565b9082610417836103fe8467ffffffffffffffff166000526007602052604060002090565b9067ffffffffffffffff16600052602052604060002090565b54036104755750505067ffffffffffffffff61045c6104567fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a9361275c565b9461275c565b610470826040519384931696169482612969565b0390a3005b7fbf1f53100000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260445260646000fd5b6000fd5b6104c56104b59161275c565b7ff094888d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b60005b83811061050e5750506000910152565b81810151838201526020016104fe565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361055a815180928187528780880191016104fb565b0116010190565b90602061057292818152019061051e565b90565b346101775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610177576102f56040516105b5606082610914565b602581527f4f6e52616d704f7665725375706572636861696e496e7465726f7020312e362e60208201527f312d646576000000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061051e565b67ffffffffffffffff81160361017757565b908160a09103126101775790565b346101775760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043561067481610619565b60243567ffffffffffffffff81116101775761069490369060040161062b565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561080557600091610841575b5061080a576107af9160209161077961076061076060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401612b69565b03915afa8015610805576102f5916000916107d6575b506040519081529081906020820190565b6107f8915060203d6020116107fe575b6107f08183610914565b810190612aff565b386107c5565b503d6107e6565b612af3565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610863915060203d602011610869575b61085b8183610914565b810190612ade565b3861072f565b503d610851565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176108bb57604052565b610870565b6060810190811067ffffffffffffffff8211176108bb57604052565b60a0810190811067ffffffffffffffff8211176108bb57604052565b6040810190811067ffffffffffffffff8211176108bb57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108bb57604052565b604051906101cf60c083610914565b604051906101cf60a083610914565b604051906101cf61012083610914565b604051906101cf608083610914565b67ffffffffffffffff81116108bb5760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361017757565b8015150361017757565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043567ffffffffffffffff8111610177573660238201121561017757806004013590610a2d82610992565b90610a3b6040519283610914565b828252602460606020840194028201019036821161017757602401925b818410610a6a57610a6883612cd3565b005b606084360312610177576020606091604051610a85816108c0565b8635610a9081610619565b815282870135610a9f816109aa565b838201526040870135610ab1816109c8565b6040820152815201930192610a58565b35906101cf826109aa565b346101775760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610177576000604051610b09816108dc565b600435610b15816109aa565b8152602435610b23816109c8565b6020820152604435610b34816109aa565b6040820152606435610b45816109aa565b6060820152608435610b56816109aa565b6080820152610b63613c2f565b73ffffffffffffffffffffffffffffffffffffffff610b96825173ffffffffffffffffffffffffffffffffffffffff1690565b16158015610d08575b8015610cfb575b610cd35780610bd57fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f192613c7a565b610bdd610983565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152610ccd60405192839283613dae565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5060208101511515610ba6565b50610d2d610760606083015173ffffffffffffffffffffffffffffffffffffffff1690565b15610b9f565b67ffffffffffffffff81116108bb57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192610d7982610d33565b91610d876040519384610914565b829481845281830111610177578281602093846000960137010152565b9080601f830112156101775781602061057293359101610d6d565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043567ffffffffffffffff811161017757610e18610e136020923690600401610da4565b612ec2565b604051908152f35b346101775760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017757610e5a600435610619565b6020610e70602435610e6b816109aa565b612f01565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101775760043567ffffffffffffffff81116101775760040160009280601f83011215610f075781359367ffffffffffffffff8511610f0457506020808301928560051b010111610177579190565b80fd5b8380fd5b3461017757610f1936610e8e565b90610f3960045473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610f4657005b610f5c610760610f57838587613029565b61303e565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610805576001948891600093611029575b5082610fd1575b5050505001610f3d565b610fdc918391613e55565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610fc7565b61104291935060203d81116107fe576107f08183610914565b9138610fc0565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775767ffffffffffffffff60043561108d81610619565b166000526006602052606060406000205473ffffffffffffffffffffffffffffffffffffffff6040519167ffffffffffffffff8116835260ff8160401c161515602084015260481c166040820152f35b6101cf9092919260a081019373ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b346101775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775761116b613048565b506102f560405161117b816108dc565b60025473ffffffffffffffffffffffffffffffffffffffff808216835260a09190911c60ff16151560208301526003541660408201526111f06111d360045473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166060830152565b61122f61121260055473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080830152565b604051918291826110dd565b346101775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760005473ffffffffffffffffffffffffffffffffffffffff811633036112fa577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775767ffffffffffffffff6004356113ba81610619565b166000526006602052600167ffffffffffffffff604060002054160167ffffffffffffffff81116113fa5760209067ffffffffffffffff60405191168152f35b613073565b906020808351928381520192019060005b81811061141d5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611410565b60409061057293921515815281602082015201906113ff565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775767ffffffffffffffff6004356114a681610619565b1680600052600660205260ff60406000205460401c169060005260066020526001604060002001906040518083602082955493848152019060005260206000209260005b81811061150e5750506114ff92500383610914565b6102f560405192839283611449565b84548352600194850194879450602090930192016114ea565b91908260a09103126101775760405161153f816108dc565b608080829480358452602081013561155681610619565b6020850152604081013561156981610619565b6040850152606081013561157c81610619565b606085015201359161158d83610619565b0152565b63ffffffff81160361017757565b35906101cf82611591565b81601f82011215610177578035906115c182610992565b926115cf6040519485610914565b82845260208085019360051b830101918183116101775760208101935b8385106115fb57505050505090565b843567ffffffffffffffff811161017757820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101775760405191611647836108dc565b602082013567ffffffffffffffff81116101775785602061166a92850101610da4565b8352604082013561167a816109aa565b602084015261168b6060830161159f565b604084015260808201359267ffffffffffffffff84116101775760a0836116b9886020809881980101610da4565b6060840152013560808201528152019401936115ec565b91909161014081840312610177576116e6610955565b926116f18183611527565b845260a082013567ffffffffffffffff81116101775781611713918401610da4565b602085015260c082013567ffffffffffffffff81116101775781611738918401610da4565b604085015261174960e08301610ac1565b6060850152610100820135608085015261012082013567ffffffffffffffff81116101775761177892016115aa565b60a0830152565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043567ffffffffffffffff811161017757610e186103b960209236906004016116d0565b34610177576117e136610e8e565b9061180461076060015473ffffffffffffffffffffffffffffffffffffffff1690565b3303611ab7575b906000905b80821061181957005b61182c61182783838661329f565b613346565b9161185c611842845167ffffffffffffffff1690565b67ffffffffffffffff166000526006602052604060002090565b9061186a6020850151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1681151560401b68ff0000000000000000161783556040850190815151611979575b50506000949294926001606086019301935b8351805182101561190657906118ff6118f96118df836001956133cc565b5173ffffffffffffffffffffffffffffffffffffffff1690565b87614035565b50016118c1565b50509390916001935051908151611922575b5050019091611810565b61196f67ffffffffffffffff6119617fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586935167ffffffffffffffff1690565b1692604051918291826126e7565b0390a23880611918565b959390929495600014611aa25760018501939060005b84518051821015611a41576118df826119a7926133cc565b73ffffffffffffffffffffffffffffffffffffffff8116156119f657906119ef6119e961076060019473ffffffffffffffffffffffffffffffffffffffff1690565b88614ad7565b500161198f565b6104b5611a0b8a5167ffffffffffffffff1690565b7f463258ff0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b505093509493917f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328167ffffffffffffffff611a84875167ffffffffffffffff1690565b925192611a986040519283921694826126e7565b0390a238806118af565b6104b5611a0b875167ffffffffffffffff1690565b611ad961076060055473ffffffffffffffffffffffffffffffffffffffff1690565b331461180b577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101775760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017757600435611b4481610619565b60243567ffffffffffffffff811161017757611b6490369060040161062b565b60643591604435611b74846109aa565b60025460a01c60ff1661230d57611bc5740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b611be38267ffffffffffffffff166000526006602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff8516156122e3578254604081901c60ff1661225b575b611c33906107609060481c73ffffffffffffffffffffffffffffffffffffffff1681565b330361223157849273ffffffffffffffffffffffffffffffffffffffff611c6f60035473ffffffffffffffffffffffffffffffffffffffff1690565b16806121b5575b5080611ccd611c98611c93611d79945467ffffffffffffffff1690565b6133e0565b82547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff821617909255565b611d2e611cd8610964565b6000815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201529167ffffffffffffffff8516604084015267ffffffffffffffff166060830152565b60006080820152600086611d4560208201826133ff565b929096611dbc611d5584806133ff565b611db0606087019b611d668d61303e565b98611da9611d8060408b019d8e8c613450565b90506134d0565b9f611d89610973565b9c8d5273ffffffffffffffffffffffffffffffffffffffff1660208d0152565b3691610d6d565b60408901523691610d6d565b6060860152611df2611dcc612ac9565b6080870190815273ffffffffffffffffffffffffffffffffffffffff90951660a0870152565b8060c086015260e08501978289526101008601998a52611e36611e3061076061076060025473ffffffffffffffffffffffffffffffffffffffff1690565b9161303e565b88611e8a611e53611e4a60808901896133ff565b919098806133ff565b91604051998a98899788977f3a49bb49000000000000000000000000000000000000000000000000000000008952600489016135db565b03915afa9586156108055760009182938391849961218a575b505252611eba611eb38489613450565b3691613634565b9260005b611ec8828a613450565b9050811015611f0f5790600182611f0581611ef28e8c8c611eec611ec89a8e6133cc565b51614133565b8c5190611eff83836133cc565b526133cc565b5001909150611ebe565b838689611f778d8760008b611f3f61076061076060025473ffffffffffffffffffffffffffffffffffffffff1690565b86516040518097819482937f01447eaa0000000000000000000000000000000000000000000000000000000084528c6004850161381e565b03915afa92831561080557600093612165575b50156120865750611fab60005b60808651019067ffffffffffffffff169052565b60005b825151811015611fdc5780611fc5600192846133cc565b516080611fd38387516133cc565b51015201611fae565b6102f584611fe987614618565b90611ff382613a19565b82515281516060015167ffffffffffffffff167f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f3267ffffffffffffffff8060405193169316918061204486826138b3565b0390a36120747fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b6040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015260208180604481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1801561080557611fab91600091612136575b50611f97565b612158915060203d60201161215e575b6121508183610914565b81019061389e565b86612130565b503d612146565b6121839193503d806000833e61217b8183610914565b8101906136ac565b9186611f8a565b9294509750506121ab913d8091833e6121a38183610914565b81019061357f565b9791939138611ea3565b8094503b1561017757600060405180957fe0a0e5060000000000000000000000000000000000000000000000000000000082528183816121f98b8960048401612b69565b03925af1908115610805578694611d7992612216575b5090611c76565b80612225600061222b93610914565b8061016c565b3861220f565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b61229a61229661228073ffffffffffffffffffffffffffffffffffffffff8916610760565b6000908152600287016020526040902054151590565b1590565b15611c0f577fd0d259760000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff861660045260246000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f830112156101775781359161234f83610992565b9261235d6040519485610914565b80845260208085019160051b830101918383116101775760208101915b83831061238957505050505090565b823567ffffffffffffffff81116101775782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610177576123d1610964565b906123de60208401610ac1565b8252604083013567ffffffffffffffff81116101775787602061240392860101610da4565b6020830152606083013567ffffffffffffffff81116101775787602061242b92860101610da4565b60408301526080830135606083015260a08301359167ffffffffffffffff83116101775761246188602080969581960101610da4565b608082015281520192019161237a565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775760043567ffffffffffffffff8111610177576101a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610177576124e7610973565b906124f53682600401611527565b825261250360a48201610ac1565b602083015260c481013567ffffffffffffffff81116101775761252c9060043691840101610da4565b604083015260e481013567ffffffffffffffff8111610177576125559060043691840101610da4565b606083015261010481013567ffffffffffffffff81116101775761257f9060043691840101610da4565b60808301526125916101248201610ac1565b60a083015261014481013560c083015261016481013560e08301526101848101359167ffffffffffffffff8311610177576125d86125e39260046102f59536920101612337565b610100820152613a19565b6040519081529081906020820190565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101775773ffffffffffffffffffffffffffffffffffffffff600435612643816109aa565b61264b613c2f565b163381146126bd57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206105729281815201906113ff565b346101775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017757612732600435610619565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b3561057281610619565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561017757016020813591019167ffffffffffffffff821161017757813603831361017757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561017757016020813591019167ffffffffffffffff8211610177578160051b3603831361017757565b90602083828152019160208260051b8501019381936000915b8483106128715750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618636030182121561017757602080918760019401906080806129536128f86128ea8680612766565b60a0875260a08701916127b6565b73ffffffffffffffffffffffffffffffffffffffff8787013561291a816109aa565b168786015263ffffffff604087013561293281611591565b1660408601526129456060870187612766565b9086830360608801526127b6565b9301359101529801930193019194939290612861565b9061057291602081528135602082015267ffffffffffffffff602083013561299081610619565b16604082015267ffffffffffffffff60408301356129ad81610619565b16606082015267ffffffffffffffff60608301356129ca81610619565b16608082015267ffffffffffffffff60808301356129e781610619565b1660a0820152612a98612a53612a16612a0360a0860186612766565b61014060c08701526101608601916127b6565b612a2360c0860186612766565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e08701526127b6565b92612a81612a6360e08301610ac1565b73ffffffffffffffffffffffffffffffffffffffff16610100850152565b6101008101356101208401526101208101906127f5565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152612848565b60405190612ad8602083610914565b60008252565b908160209103126101775751610572816109c8565b6040513d6000823e3d90fd5b90816020910312610177575190565b9160209082815201919060005b818110612b285750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735612b51816109aa565b16815260208781013590820152019401929101612b1b565b919067ffffffffffffffff16825260406020830152612bdc612b9f612b8e8380612766565b60a0604087015260e08601916127b6565b612bac6020840184612766565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08684030160608701526127b6565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561017757820160208135910167ffffffffffffffff8211610177578160061b360381136101775784612ca392612c6c927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866105729903016080870152612b0e565b92612c99612c7c60608301610ac1565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b6080810190612766565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828603019101526127b6565b612cdb613c2f565b6000915b8151831015612ebd57612cf283836133cc565b5192612d10612d0182856133cc565b515167ffffffffffffffff1690565b9367ffffffffffffffff8516908115612e85577fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590612d67600195969767ffffffffffffffff166000526006602052604060002090565b612e1d612de36040612d90602086015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffff0000000000000000000000000000000000000000ffffffffffffffffff16604882901b7cffffffffffffffffffffffffffffffffffffffff00000000000000000016178555940151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1690151560401b68ff000000000000000016178255565b54612e7a612e3967ffffffffffffffff83169260401c60ff1690565b60405193849384919273ffffffffffffffffffffffffffffffffffffffff604092959467ffffffffffffffff60608601971685521660208401521515910152565b0390a2019190612cdf565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b915050565b805160248110612ed457506024015190565b7f7946a7cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa801561080557600090612fab575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011612ff2575b81612fc560209383610914565b810103126101775773ffffffffffffffffffffffffffffffffffffffff9051612fed816109aa565b612f90565b3d9150612fb8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156130395760051b0190565b612ffa565b35610572816109aa565b60405190613055826108dc565b60006080838281528260208201528260408201528260608201520152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff604082510151166040516020810190308252602081526130cc604082610914565b5190206040519060208201927f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f845267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166040840152606083015260808201526080815261314460a082610914565b519020906132998151805190613225613174606086015173ffffffffffffffffffffffffffffffffffffffff1690565b6131f961318c606085015167ffffffffffffffff1690565b936131a66080808a015192015167ffffffffffffffff1690565b906040519586946020860198899367ffffffffffffffff60809473ffffffffffffffffffffffffffffffffffffffff82959998949960a089019a8952166020880152166040860152606085015216910152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610914565b5190206131f96020840151602081519101209360a060408201516020815191012091015160405161325e816131f9602082019485614024565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b51902090565b91908110156130395760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301821215610177570190565b9080601f830112156101775781356132f681610992565b926133046040519485610914565b81845260208085019260051b82010192831161017757602001905b82821061332c5750505090565b60208091833561333b816109aa565b81520191019061331f565b608081360312610177576040519061335d8261089f565b803561336881610619565b82526020810135613378816109c8565b6020830152604081013567ffffffffffffffff81116101775761339e90369083016132df565b604083015260608101359067ffffffffffffffff8211610177576133c4913691016132df565b606082015290565b80518210156130395760209160051b010190565b67ffffffffffffffff1667ffffffffffffffff81146113fa5760010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610177570180359067ffffffffffffffff82116101775760200191813603831361017757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610177570180359067ffffffffffffffff821161017757602001918160061b3603831361017757565b604051906134b1826108dc565b6060608083600081528260208201528260408201526000838201520152565b906134da82610992565b6134e76040519182610914565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06135158294610992565b019060005b82811061352657505050565b6020906135316134a4565b8282850101520161351a565b81601f8201121561017757805161355381610d33565b926135616040519485610914565b818452602082840101116101775761057291602080850191016104fb565b9060808282031261017757815192602083015161359b816109c8565b92604081015167ffffffffffffffff811161017757836135bc91830161353d565b92606082015167ffffffffffffffff811161017757610572920161353d565b9593919273ffffffffffffffffffffffffffffffffffffffff6136269467ffffffffffffffff6105729a9894168952166020880152604087015260a0606087015260a08601916127b6565b9260808185039101526127b6565b92919261364082610992565b9361364e6040519586610914565b602085848152019260061b82019181831161017757925b8284106136725750505050565b604084830312610177576020604091825161368c816108f8565b8635613697816109aa565b81528287013583820152815201930192613665565b6020818303126101775780519067ffffffffffffffff821161017757019080601f830112156101775781516136e081610992565b926136ee6040519485610914565b81845260208085019260051b820101918383116101775760208201905b83821061371a57505050505090565b815167ffffffffffffffff81116101775760209161373d8784809488010161353d565b81520191019061370b565b9080602083519182815201916020808360051b8301019401926000915b83831061377457505050505090565b909192939460208061380f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806137f46137e28685015160a08886015260a085019061051e565b6040850151848203604086015261051e565b9260608101516060840152015190608081840391015261051e565b97019301930191939290613765565b9167ffffffffffffffff61384092168352606060208401526060830190613748565b9060408183039101526020808351928381520192019060005b8181106138665750505090565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101613859565b90816020910312610177575161057281610619565b90610572916020815261390260208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660c08201526101006139ae61397961394660408601516101a060e08701526101c086019061051e565b60608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303018587015261051e565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030161012086015261051e565b60a084015173ffffffffffffffffffffffffffffffffffffffff166101408401529260c081015161016084015260e08101516101808401520151906101a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613748565b80515180613c02575067ffffffffffffffff6040825101511660405160208101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321835267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166040830152606082015230608082015260808152613aa760a082610914565b51902090613299613acf602083015173ffffffffffffffffffffffffffffffffffffffff1690565b8251613b7c613afe6080613aee606085015167ffffffffffffffff1690565b93015167ffffffffffffffff1690565b916131f9613b2360a088015173ffffffffffffffffffffffffffffffffffffffff1690565b60c088015190604051958694602086019889919367ffffffffffffffff6080948173ffffffffffffffffffffffffffffffffffffffff949998978560a088019b1687521660208601521660408401521660608201520152565b5190206131f960608401516020815191012093604081015160208151910120906080610100820151604051613bb9816131f96020820194856149e7565b519020910151602081519101209160405196879560208701998a9260c094919796959260e085019860008652602086015260408501526060840152608083015260a08201520152565b7f4c8ebcc00000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303613c5057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169183169190911790556101cf91608090613d698360608301511660049073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b01511660059073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b9160806101cf929493613e088161012081019773ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b9073ffffffffffffffffffffffffffffffffffffffff613f279392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452613eb7606485610914565b16600080604093845195613ecb8688610914565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15613f4c573d613f18613f0f82610d33565b94519485610914565b83523d6000602085013e614cfe565b805180613f32575050565b81602080613f47936101cf9501019101612ade565b6149f8565b60609250614cfe565b9080602083519182815201916020808360051b8301019401926000915b838310613f8157505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519060808061400f613fcf855160a0865260a086019061051e565b73ffffffffffffffffffffffffffffffffffffffff87870151168786015263ffffffff60408701511660408601526060860151858203606087015261051e565b93015191015297019301930191939290613f72565b906020610572928181520190613f55565b73ffffffffffffffffffffffffffffffffffffffff610572921690614bfb565b6020818303126101775780519067ffffffffffffffff821161017757016040818303126101775760405191614089836108f8565b815167ffffffffffffffff811161017757816140a691840161353d565b8352602082015167ffffffffffffffff8111610177576140c6920161353d565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080614100855184602087015260c086019061051e565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b93929161413e6134a4565b5060208501918251156143d257614172610760610e6b610760895173ffffffffffffffffffffffffffffffffffffffff1690565b9373ffffffffffffffffffffffffffffffffffffffff8516158015614347575b6142e45760009291614259959697614204614226936141e76141ca8951945173ffffffffffffffffffffffffffffffffffffffff1690565b946141d3610964565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016140ce565b038183875af1918215610805576000926142bf575b50602082519201519051916142a0614284610964565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b6020840152604083015260608201526142b7612ac9565b608082015290565b6142dd9192503d806000833e6142d58183610914565b810190614055565b903861426e565b6104b5614305885173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481895afa908115610805576000916143b3575b5015614192565b6143cc915060203d6020116108695761085b8183610914565b386143ac565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190610120820182811067ffffffffffffffff8211176108bb57604052606061010083614429613048565b8152600060208201528260408201528280820152826080820152600060a0820152600060c0820152600060e08201520152565b9061446682610992565b6144736040519182610914565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06144a18294610992565b019060005b8281106144b257505050565b6020906040516144c1816108dc565b606081526000838201526000604082015260608082015260006080820152828285010152016144a6565b908160209103126101775751610572816109aa565b90816020910312610177575161057281611591565b90610572916020815261456460208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b60a06145b8614584602085015161014060c086015261016085019061051e565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160e086015261051e565b9273ffffffffffffffffffffffffffffffffffffffff60608201511661010084015260808101516101208401520151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613f55565b61462a906146246143fc565b50614cf2565b906146386080830151612ec2565b61010083019161464983515161445c565b9360005b8451805182101561479157906146b4614668826001946133cc565b516146e061468a825173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290938491820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610914565b6146f76020820151602080825183010191016144eb565b9061476b614712608083015160208082518301019101614500565b61475e73ffffffffffffffffffffffffffffffffffffffff606060408601519501519561473d610964565b9788521673ffffffffffffffffffffffffffffffffffffffff166020870152565b63ffffffff166040850152565b6060830152608082015261477f82896133cc565b5261478a81886133cc565b500161464d565b505092509282519161483e6147b1602085015167ffffffffffffffff1690565b9361482d6147ca604083015167ffffffffffffffff1690565b9161481c6147e86080613aee606085015167ffffffffffffffff1690565b9361480b6147f4610964565b6000815267ffffffffffffffff909a1660208b0152565b67ffffffffffffffff166040890152565b67ffffffffffffffff166060870152565b67ffffffffffffffff166080850152565b6148f773ffffffffffffffffffffffffffffffffffffffff6131f96148a761487d602089015173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290928391820190565b60408701516148c36060890151602080825183010191016144eb565b916148cc610955565b978852602088015260408701521673ffffffffffffffffffffffffffffffffffffffff166060850152565b608083015260a082015261490a82613a19565b815152614916816130a2565b61497561494d614933604085510167ffffffffffffffff90511690565b67ffffffffffffffff166000526007602052604060002090565b83516060015167ffffffffffffffff1667ffffffffffffffff16600052602052604060002090565b558051907fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a67ffffffffffffffff6149cd60606149bd604087015167ffffffffffffffff1690565b95015167ffffffffffffffff1690565b6149e1826040519384931696169482614515565b0390a390565b906020610572928181520190613748565b156149ff57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80548210156130395760005260206000200190600090565b91614ad3918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b6000828152600182016020526040902054614b6157805490680100000000000000008210156108bb5782614b4a614b15846001809601855584614a83565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b80548015614bcc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614b9d8282614a83565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014614ce9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116113fa578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116113fa576000958583614c9a97614c8b9503614ca0575b505050614b68565b90600052602052604060002090565b55600190565b614cd0614cca91614cc1614cb7614ce09588614a83565b90549060031b1c90565b92839187614a83565b90614a9b565b8590600052602052604060002090565b55388080614c83565b50505050600090565b614cfa6143fc565b5090565b91929015614d795750815115614d12575090565b3b15614d1b5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614d8c5750805190602001fd5b614dc2906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610561565b0390fdfea164736f6c634300081a000a",
}

var OnRampOverSuperchainInteropABI = OnRampOverSuperchainInteropMetaData.ABI

var OnRampOverSuperchainInteropBin = OnRampOverSuperchainInteropMetaData.Bin

func DeployOnRampOverSuperchainInterop(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig, destChainConfigs []OnRampDestChainConfigArgs) (common.Address, *types.Transaction, *OnRampOverSuperchainInterop, error) {
	parsed, err := OnRampOverSuperchainInteropMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampOverSuperchainInteropBin), backend, staticConfig, dynamicConfig, destChainConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRampOverSuperchainInterop{address: address, abi: *parsed, OnRampOverSuperchainInteropCaller: OnRampOverSuperchainInteropCaller{contract: contract}, OnRampOverSuperchainInteropTransactor: OnRampOverSuperchainInteropTransactor{contract: contract}, OnRampOverSuperchainInteropFilterer: OnRampOverSuperchainInteropFilterer{contract: contract}}, nil
}

type OnRampOverSuperchainInterop struct {
	address common.Address
	abi     abi.ABI
	OnRampOverSuperchainInteropCaller
	OnRampOverSuperchainInteropTransactor
	OnRampOverSuperchainInteropFilterer
}

type OnRampOverSuperchainInteropCaller struct {
	contract *bind.BoundContract
}

type OnRampOverSuperchainInteropTransactor struct {
	contract *bind.BoundContract
}

type OnRampOverSuperchainInteropFilterer struct {
	contract *bind.BoundContract
}

type OnRampOverSuperchainInteropSession struct {
	Contract     *OnRampOverSuperchainInterop
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampOverSuperchainInteropCallerSession struct {
	Contract *OnRampOverSuperchainInteropCaller
	CallOpts bind.CallOpts
}

type OnRampOverSuperchainInteropTransactorSession struct {
	Contract     *OnRampOverSuperchainInteropTransactor
	TransactOpts bind.TransactOpts
}

type OnRampOverSuperchainInteropRaw struct {
	Contract *OnRampOverSuperchainInterop
}

type OnRampOverSuperchainInteropCallerRaw struct {
	Contract *OnRampOverSuperchainInteropCaller
}

type OnRampOverSuperchainInteropTransactorRaw struct {
	Contract *OnRampOverSuperchainInteropTransactor
}

func NewOnRampOverSuperchainInterop(address common.Address, backend bind.ContractBackend) (*OnRampOverSuperchainInterop, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampOverSuperchainInteropABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRampOverSuperchainInterop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInterop{address: address, abi: abi, OnRampOverSuperchainInteropCaller: OnRampOverSuperchainInteropCaller{contract: contract}, OnRampOverSuperchainInteropTransactor: OnRampOverSuperchainInteropTransactor{contract: contract}, OnRampOverSuperchainInteropFilterer: OnRampOverSuperchainInteropFilterer{contract: contract}}, nil
}

func NewOnRampOverSuperchainInteropCaller(address common.Address, caller bind.ContractCaller) (*OnRampOverSuperchainInteropCaller, error) {
	contract, err := bindOnRampOverSuperchainInterop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropCaller{contract: contract}, nil
}

func NewOnRampOverSuperchainInteropTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampOverSuperchainInteropTransactor, error) {
	contract, err := bindOnRampOverSuperchainInterop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropTransactor{contract: contract}, nil
}

func NewOnRampOverSuperchainInteropFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampOverSuperchainInteropFilterer, error) {
	contract, err := bindOnRampOverSuperchainInterop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropFilterer{contract: contract}, nil
}

func bindOnRampOverSuperchainInterop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnRampOverSuperchainInteropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampOverSuperchainInterop.Contract.OnRampOverSuperchainInteropCaller.contract.Call(opts, result, method, params...)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.OnRampOverSuperchainInteropTransactor.contract.Transfer(opts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.OnRampOverSuperchainInteropTransactor.contract.Transact(opts, method, params...)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampOverSuperchainInterop.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.contract.Transfer(opts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.contract.Transact(opts, method, params...)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) ExtractGasLimit(opts *bind.CallOpts, extraArgs []byte) (*big.Int, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "extractGasLimit", extraArgs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) ExtractGasLimit(extraArgs []byte) (*big.Int, error) {
	return _OnRampOverSuperchainInterop.Contract.ExtractGasLimit(&_OnRampOverSuperchainInterop.CallOpts, extraArgs)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) ExtractGasLimit(extraArgs []byte) (*big.Int, error) {
	return _OnRampOverSuperchainInterop.Contract.ExtractGasLimit(&_OnRampOverSuperchainInterop.CallOpts, extraArgs)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "generateMessageId", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRampOverSuperchainInterop.Contract.GenerateMessageId(&_OnRampOverSuperchainInterop.CallOpts, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRampOverSuperchainInterop.Contract.GenerateMessageId(&_OnRampOverSuperchainInterop.CallOpts, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

	error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	outstruct := new(GetAllowedSendersList)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfiguredAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRampOverSuperchainInterop.Contract.GetAllowedSendersList(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRampOverSuperchainInterop.Contract.GetAllowedSendersList(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AllowlistEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRampOverSuperchainInterop.Contract.GetDestChainConfig(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRampOverSuperchainInterop.Contract.GetDestChainConfig(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDynamicConfig)).(*OnRampDynamicConfig)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRampOverSuperchainInterop.Contract.GetDynamicConfig(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRampOverSuperchainInterop.Contract.GetDynamicConfig(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRampOverSuperchainInterop.Contract.GetExpectedNextSequenceNumber(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRampOverSuperchainInterop.Contract.GetExpectedNextSequenceNumber(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRampOverSuperchainInterop.Contract.GetFee(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRampOverSuperchainInterop.Contract.GetFee(&_OnRampOverSuperchainInterop.CallOpts, destChainSelector, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.GetPoolBySourceToken(&_OnRampOverSuperchainInterop.CallOpts, arg0, sourceToken)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.GetPoolBySourceToken(&_OnRampOverSuperchainInterop.CallOpts, arg0, sourceToken)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampStaticConfig)).(*OnRampStaticConfig)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRampOverSuperchainInterop.Contract.GetStaticConfig(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRampOverSuperchainInterop.Contract.GetStaticConfig(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.GetSupportedTokens(&_OnRampOverSuperchainInterop.CallOpts, arg0)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.GetSupportedTokens(&_OnRampOverSuperchainInterop.CallOpts, arg0)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) HashInteropMessage(opts *bind.CallOpts, message InternalAny2EVMRampMessage) ([32]byte, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "hashInteropMessage", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) HashInteropMessage(message InternalAny2EVMRampMessage) ([32]byte, error) {
	return _OnRampOverSuperchainInterop.Contract.HashInteropMessage(&_OnRampOverSuperchainInterop.CallOpts, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) HashInteropMessage(message InternalAny2EVMRampMessage) ([32]byte, error) {
	return _OnRampOverSuperchainInterop.Contract.HashInteropMessage(&_OnRampOverSuperchainInterop.CallOpts, message)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) Owner() (common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.Owner(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) Owner() (common.Address, error) {
	return _OnRampOverSuperchainInterop.Contract.Owner(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRampOverSuperchainInterop.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) TypeAndVersion() (string, error) {
	return _OnRampOverSuperchainInterop.Contract.TypeAndVersion(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropCallerSession) TypeAndVersion() (string, error) {
	return _OnRampOverSuperchainInterop.Contract.TypeAndVersion(&_OnRampOverSuperchainInterop.CallOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "acceptOwnership")
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.AcceptOwnership(&_OnRampOverSuperchainInterop.TransactOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.AcceptOwnership(&_OnRampOverSuperchainInterop.TransactOpts)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ApplyAllowlistUpdates(&_OnRampOverSuperchainInterop.TransactOpts, allowlistConfigArgsItems)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ApplyAllowlistUpdates(&_OnRampOverSuperchainInterop.TransactOpts, allowlistConfigArgsItems)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ApplyDestChainConfigUpdates(&_OnRampOverSuperchainInterop.TransactOpts, destChainConfigArgs)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ApplyDestChainConfigUpdates(&_OnRampOverSuperchainInterop.TransactOpts, destChainConfigArgs)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ForwardFromRouter(&_OnRampOverSuperchainInterop.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ForwardFromRouter(&_OnRampOverSuperchainInterop.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) ReemitInteropMessage(opts *bind.TransactOpts, interopMessage InternalAny2EVMRampMessage) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "reemitInteropMessage", interopMessage)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) ReemitInteropMessage(interopMessage InternalAny2EVMRampMessage) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ReemitInteropMessage(&_OnRampOverSuperchainInterop.TransactOpts, interopMessage)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) ReemitInteropMessage(interopMessage InternalAny2EVMRampMessage) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.ReemitInteropMessage(&_OnRampOverSuperchainInterop.TransactOpts, interopMessage)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.SetDynamicConfig(&_OnRampOverSuperchainInterop.TransactOpts, dynamicConfig)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.SetDynamicConfig(&_OnRampOverSuperchainInterop.TransactOpts, dynamicConfig)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.TransferOwnership(&_OnRampOverSuperchainInterop.TransactOpts, to)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.TransferOwnership(&_OnRampOverSuperchainInterop.TransactOpts, to)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.WithdrawFeeTokens(&_OnRampOverSuperchainInterop.TransactOpts, feeTokens)
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampOverSuperchainInterop.Contract.WithdrawFeeTokens(&_OnRampOverSuperchainInterop.TransactOpts, feeTokens)
}

type OnRampOverSuperchainInteropAllowListAdminSetIterator struct {
	Event *OnRampOverSuperchainInteropAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropAllowListAdminSet)
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
		it.Event = new(OnRampOverSuperchainInteropAllowListAdminSet)
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

func (it *OnRampOverSuperchainInteropAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropAllowListAdminSet struct {
	AllowlistAdmin common.Address
	Raw            types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampOverSuperchainInteropAllowListAdminSetIterator, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropAllowListAdminSetIterator{contract: _OnRampOverSuperchainInterop.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropAllowListAdminSet)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseAllowListAdminSet(log types.Log) (*OnRampOverSuperchainInteropAllowListAdminSet, error) {
	event := new(OnRampOverSuperchainInteropAllowListAdminSet)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropAllowListSendersAddedIterator struct {
	Event *OnRampOverSuperchainInteropAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropAllowListSendersAdded)
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
		it.Event = new(OnRampOverSuperchainInteropAllowListSendersAdded)
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

func (it *OnRampOverSuperchainInteropAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropAllowListSendersAddedIterator{contract: _OnRampOverSuperchainInterop.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropAllowListSendersAdded)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseAllowListSendersAdded(log types.Log) (*OnRampOverSuperchainInteropAllowListSendersAdded, error) {
	event := new(OnRampOverSuperchainInteropAllowListSendersAdded)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropAllowListSendersRemovedIterator struct {
	Event *OnRampOverSuperchainInteropAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropAllowListSendersRemoved)
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
		it.Event = new(OnRampOverSuperchainInteropAllowListSendersRemoved)
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

func (it *OnRampOverSuperchainInteropAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropAllowListSendersRemovedIterator{contract: _OnRampOverSuperchainInterop.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropAllowListSendersRemoved)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseAllowListSendersRemoved(log types.Log) (*OnRampOverSuperchainInteropAllowListSendersRemoved, error) {
	event := new(OnRampOverSuperchainInteropAllowListSendersRemoved)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropCCIPMessageSentIterator struct {
	Event *OnRampOverSuperchainInteropCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropCCIPMessageSent)
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
		it.Event = new(OnRampOverSuperchainInteropCCIPMessageSent)
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

func (it *OnRampOverSuperchainInteropCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampOverSuperchainInteropCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropCCIPMessageSentIterator{contract: _OnRampOverSuperchainInterop.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropCCIPMessageSent)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseCCIPMessageSent(log types.Log) (*OnRampOverSuperchainInteropCCIPMessageSent, error) {
	event := new(OnRampOverSuperchainInteropCCIPMessageSent)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator struct {
	Event *OnRampOverSuperchainInteropCCIPSuperchainMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropCCIPSuperchainMessageSent)
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
		it.Event = new(OnRampOverSuperchainInteropCCIPSuperchainMessageSent)
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

func (it *OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropCCIPSuperchainMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalAny2EVMRampMessage
	Raw               types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterCCIPSuperchainMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "CCIPSuperchainMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator{contract: _OnRampOverSuperchainInterop.contract, event: "CCIPSuperchainMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchCCIPSuperchainMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropCCIPSuperchainMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "CCIPSuperchainMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropCCIPSuperchainMessageSent)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "CCIPSuperchainMessageSent", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseCCIPSuperchainMessageSent(log types.Log) (*OnRampOverSuperchainInteropCCIPSuperchainMessageSent, error) {
	event := new(OnRampOverSuperchainInteropCCIPSuperchainMessageSent)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "CCIPSuperchainMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropConfigSetIterator struct {
	Event *OnRampOverSuperchainInteropConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropConfigSet)
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
		it.Event = new(OnRampOverSuperchainInteropConfigSet)
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

func (it *OnRampOverSuperchainInteropConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropConfigSet struct {
	StaticConfig  OnRampStaticConfig
	DynamicConfig OnRampDynamicConfig
	Raw           types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OnRampOverSuperchainInteropConfigSetIterator, error) {

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropConfigSetIterator{contract: _OnRampOverSuperchainInterop.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropConfigSet)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseConfigSet(log types.Log) (*OnRampOverSuperchainInteropConfigSet, error) {
	event := new(OnRampOverSuperchainInteropConfigSet)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropDestChainConfigSetIterator struct {
	Event *OnRampOverSuperchainInteropDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropDestChainConfigSet)
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
		it.Event = new(OnRampOverSuperchainInteropDestChainConfigSet)
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

func (it *OnRampOverSuperchainInteropDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropDestChainConfigSetIterator{contract: _OnRampOverSuperchainInterop.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropDestChainConfigSet)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseDestChainConfigSet(log types.Log) (*OnRampOverSuperchainInteropDestChainConfigSet, error) {
	event := new(OnRampOverSuperchainInteropDestChainConfigSet)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropFeeTokenWithdrawnIterator struct {
	Event *OnRampOverSuperchainInteropFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropFeeTokenWithdrawn)
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
		it.Event = new(OnRampOverSuperchainInteropFeeTokenWithdrawn)
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

func (it *OnRampOverSuperchainInteropFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampOverSuperchainInteropFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropFeeTokenWithdrawnIterator{contract: _OnRampOverSuperchainInterop.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropFeeTokenWithdrawn)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseFeeTokenWithdrawn(log types.Log) (*OnRampOverSuperchainInteropFeeTokenWithdrawn, error) {
	event := new(OnRampOverSuperchainInteropFeeTokenWithdrawn)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropOwnershipTransferRequestedIterator struct {
	Event *OnRampOverSuperchainInteropOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropOwnershipTransferRequested)
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
		it.Event = new(OnRampOverSuperchainInteropOwnershipTransferRequested)
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

func (it *OnRampOverSuperchainInteropOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOverSuperchainInteropOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropOwnershipTransferRequestedIterator{contract: _OnRampOverSuperchainInterop.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropOwnershipTransferRequested)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampOverSuperchainInteropOwnershipTransferRequested, error) {
	event := new(OnRampOverSuperchainInteropOwnershipTransferRequested)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOverSuperchainInteropOwnershipTransferredIterator struct {
	Event *OnRampOverSuperchainInteropOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOverSuperchainInteropOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOverSuperchainInteropOwnershipTransferred)
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
		it.Event = new(OnRampOverSuperchainInteropOwnershipTransferred)
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

func (it *OnRampOverSuperchainInteropOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampOverSuperchainInteropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOverSuperchainInteropOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOverSuperchainInteropOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOverSuperchainInteropOwnershipTransferredIterator{contract: _OnRampOverSuperchainInterop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampOverSuperchainInterop.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOverSuperchainInteropOwnershipTransferred)
				if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInteropFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampOverSuperchainInteropOwnershipTransferred, error) {
	event := new(OnRampOverSuperchainInteropOwnershipTransferred)
	if err := _OnRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetAllowedSendersList struct {
	IsEnabled           bool
	ConfiguredAddresses []common.Address
}
type GetDestChainConfig struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInterop) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRampOverSuperchainInterop.abi.Events["AllowListAdminSet"].ID:
		return _OnRampOverSuperchainInterop.ParseAllowListAdminSet(log)
	case _OnRampOverSuperchainInterop.abi.Events["AllowListSendersAdded"].ID:
		return _OnRampOverSuperchainInterop.ParseAllowListSendersAdded(log)
	case _OnRampOverSuperchainInterop.abi.Events["AllowListSendersRemoved"].ID:
		return _OnRampOverSuperchainInterop.ParseAllowListSendersRemoved(log)
	case _OnRampOverSuperchainInterop.abi.Events["CCIPMessageSent"].ID:
		return _OnRampOverSuperchainInterop.ParseCCIPMessageSent(log)
	case _OnRampOverSuperchainInterop.abi.Events["CCIPSuperchainMessageSent"].ID:
		return _OnRampOverSuperchainInterop.ParseCCIPSuperchainMessageSent(log)
	case _OnRampOverSuperchainInterop.abi.Events["ConfigSet"].ID:
		return _OnRampOverSuperchainInterop.ParseConfigSet(log)
	case _OnRampOverSuperchainInterop.abi.Events["DestChainConfigSet"].ID:
		return _OnRampOverSuperchainInterop.ParseDestChainConfigSet(log)
	case _OnRampOverSuperchainInterop.abi.Events["FeeTokenWithdrawn"].ID:
		return _OnRampOverSuperchainInterop.ParseFeeTokenWithdrawn(log)
	case _OnRampOverSuperchainInterop.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRampOverSuperchainInterop.ParseOwnershipTransferRequested(log)
	case _OnRampOverSuperchainInterop.abi.Events["OwnershipTransferred"].ID:
		return _OnRampOverSuperchainInterop.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampOverSuperchainInteropAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (OnRampOverSuperchainInteropAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (OnRampOverSuperchainInteropAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (OnRampOverSuperchainInteropCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32")
}

func (OnRampOverSuperchainInteropCCIPSuperchainMessageSent) Topic() common.Hash {
	return common.HexToHash("0xb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a")
}

func (OnRampOverSuperchainInteropConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1")
}

func (OnRampOverSuperchainInteropDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
}

func (OnRampOverSuperchainInteropFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (OnRampOverSuperchainInteropOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOverSuperchainInteropOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRampOverSuperchainInterop *OnRampOverSuperchainInterop) Address() common.Address {
	return _OnRampOverSuperchainInterop.address
}

type OnRampOverSuperchainInteropInterface interface {
	ExtractGasLimit(opts *bind.CallOpts, extraArgs []byte) (*big.Int, error)

	GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error)

	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	HashInteropMessage(opts *bind.CallOpts, message InternalAny2EVMRampMessage) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	ReemitInteropMessage(opts *bind.TransactOpts, interopMessage InternalAny2EVMRampMessage) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampOverSuperchainInteropAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*OnRampOverSuperchainInteropAllowListAdminSet, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*OnRampOverSuperchainInteropAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*OnRampOverSuperchainInteropAllowListSendersRemoved, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampOverSuperchainInteropCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampOverSuperchainInteropCCIPMessageSent, error)

	FilterCCIPSuperchainMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampOverSuperchainInteropCCIPSuperchainMessageSentIterator, error)

	WatchCCIPSuperchainMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropCCIPSuperchainMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPSuperchainMessageSent(log types.Log) (*OnRampOverSuperchainInteropCCIPSuperchainMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampOverSuperchainInteropConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampOverSuperchainInteropConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampOverSuperchainInteropDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampOverSuperchainInteropDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampOverSuperchainInteropFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampOverSuperchainInteropFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOverSuperchainInteropOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOverSuperchainInteropOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOverSuperchainInteropOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOverSuperchainInteropOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOverSuperchainInteropOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
