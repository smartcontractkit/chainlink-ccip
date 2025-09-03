// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp_over_superchain_interop

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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extractGasLimit\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMessageId\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reemitInteropMessage\",\"inputs\":[{\"name\":\"interopMessage\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPSuperchainMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgsTooShort\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceChainSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageDoesNotExist\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MessageIdUnexpectedlySet\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x61010060405234610566576153f68038038061001a816105a0565b92833981019080820390610140821261056657608082126105665761003d610581565b90610047816105c5565b82526020810151906001600160a01b03821682036105665760208301918252610072604082016105d9565b926040810193845260a0610088606084016105d9565b6060830190815295607f1901126105665760405160a081016001600160401b0381118282101761056b576040526100c1608084016105d9565b81526100cf60a084016105ed565b602082019081526100e260c085016105d9565b91604081019283526100f660e086016105d9565b936060820194855261010b61010087016105d9565b6080830190815261012087015190966001600160401b03821161056657018a601f82011215610566578051906001600160401b03821161056b5760209b8c6060610159828660051b016105a0565b9e8f8681520194028301019181831161056657602001925b8284106104f8575050505033156104e757600180546001600160a01b0319163317905580516001600160401b03161580156104d5575b80156104c3575b80156104b1575b61048457516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e0819052825190911615801561049f575b8015610495575b61048457815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19861012098606061029f610581565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a16000905b805182101561040b5761032b82826105fa565b51916001600160401b0361033f82846105fa565b5151169283156103f65760008481526006602090815260409182902081840151815494840151600160401b600160e81b03198616604883901b600160481b600160e81b031617901515851b68ff000000000000000016179182905583516001600160401b0390951685526001600160a01b031691840191909152811c60ff1615159082015291926001927fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590606090a20190610318565b8363c35aa79d60e01b60005260045260246000fd5b604051614dd19081610625823960805181818161022b0152818161037401528181610c3701528181611a500152613817015260a0518181816102640152818161085a0152610c70015260c05181818161028c01528181610cac0152611e65015260e0518181816102b401528181610ce80152612f0c0152f35b6306b7c75960e31b60005260046000fd5b5082511515610200565b5084516001600160a01b0316156101f9565b5088516001600160a01b0316156101b5565b5087516001600160a01b0316156101ae565b5086516001600160a01b0316156101a7565b639b15e16f60e01b60005260046000fd5b6060848303126105665760405190606082016001600160401b0381118382101761056b57604052610528856105c5565b82526020850151906001600160a01b03821682036105665782602092836060950152610556604088016105ed565b6040820152815201930192610171565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761056b57604052565b6040519190601f01601f191682016001600160401b0381118382101761056b57604052565b51906001600160401b038216820361056657565b51906001600160a01b038216820361056657565b5190811515820361056657565b805182101561060e5760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146101575780630be49dbf14610152578063181f5a771461014d57806320487ded146101485780632716072b1461014357806327e936f11461013e57806328938e371461013957806348a98aa4146101345780635cb80c5d1461012f5780636def4ce71461012a5780637437ff9f1461012557806379ba5097146101205780638da5cb5b1461011b5780639041be3d14610116578063972b461214610111578063c9b146b31461010c578063df0aa9e914610107578063f2890a2114610102578063f2fde38b146100fd5763fbca3b74146100f857600080fd5b6124cb565b6123c6565b612244565b611872565b61153c565b611477565b61138b565b611339565b611250565b611149565b61105e565b610f20565b610e35565b610dd4565b610b1b565b610a21565b6107aa565b6106e6565b6102e9565b6101c1565b600091031261016757565b600080fd5b6101bf90929192608081019373ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757600060606040516102008161052a565b82815282602082015282604082015201526102e56040516102208161052a565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660208301527f0000000000000000000000000000000000000000000000000000000000000000811660408301527f00000000000000000000000000000000000000000000000000000000000000001660608201526040519182918261016c565b0390f35b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757806004016101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261016757602482016103688161252f565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036104b957506103af306103aa3684612678565b613a10565b60448301908061040460646103e06103c68661252f565b67ffffffffffffffff166000526007602052604060002090565b9601956103ec8761252f565b67ffffffffffffffff16600052602052604060002090565b5403610460575067ffffffffffffffff6104476104417fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a9361252f565b9461252f565b61045b82604051938493169616948261292a565b0390a3005b836104766104706104b59461252f565b9161252f565b7fbf1f53100000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604452606490565b6000fd5b6104c56104b59161252f565b7ff094888d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761054657604052565b6104fb565b6060810190811067ffffffffffffffff82111761054657604052565b60a0810190811067ffffffffffffffff82111761054657604052565b6040810190811067ffffffffffffffff82111761054657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761054657604052565b604051906101bf60a08361059f565b604051906101bf6101208361059f565b604051906101bf60c08361059f565b604051906101bf60808361059f565b67ffffffffffffffff811161054657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061066660208361059f565b60008252565b60005b83811061067f5750506000910152565b818101518382015260200161066f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936106cb8151809281875287808801910161066c565b0116010190565b9060206106e392818152019061068f565b90565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576102e560405161072660608261059f565b602581527f4f6e52616d704f7665725375706572636861696e496e7465726f7020312e362e60208201527f322d646576000000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061068f565b67ffffffffffffffff81160361016757565b908160a09103126101675790565b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576004356107e58161078a565b60243567ffffffffffffffff81116101675761080590369060040161079c565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa908115610976576000916109b2575b5061097b57610920916020916108ea6108d16108d160025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401612b15565b03915afa8015610976576102e591600091610947575b506040519081529081906020820190565b610969915060203d60201161096f575b610961818361059f565b810190612aab565b38610936565b503d610957565b612a9f565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6109d4915060203d6020116109da575b6109cc818361059f565b810190612a8a565b386108a0565b503d6109c2565b67ffffffffffffffff81116105465760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361016757565b8015150361016757565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff8111610167573660238201121561016757806004013590610a7c826109e1565b90610a8a604051928361059f565b828252602460606020840194028201019036821161016757602401925b818410610ab957610ab783612c7f565b005b606084360312610167576020606091604051610ad48161054b565b8635610adf8161078a565b815282870135610aee816109f9565b838201526040870135610b0081610a17565b6040820152815201930192610aa7565b35906101bf826109f9565b346101675760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576000604051610b5881610567565b600435610b64816109f9565b8152602435610b7281610a17565b6020820152604435610b83816109f9565b6040820152606435610b94816109f9565b6060820152608435610ba5816109f9565b6080820152610bb2613ab3565b73ffffffffffffffffffffffffffffffffffffffff610be5825173ffffffffffffffffffffffffffffffffffffffff1690565b16158015610d57575b8015610d4a575b610d225780610c247fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f192613afe565b610c2c61060e565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152610d1c60405192839283613c32565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5060208101511515610bf5565b50610d7c6108d1606083015173ffffffffffffffffffffffffffffffffffffffff1690565b15610bee565b929192610d8e8261061d565b91610d9c604051938461059f565b829481845281830111610167578281602093846000960137010152565b9080601f83011215610167578160206106e393359101610d82565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757610e2d610e286020923690600401610db9565b612e6e565b604051908152f35b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757610e6f60043561078a565b6020610e85602435610e80816109f9565b612ead565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101675760043567ffffffffffffffff81116101675760040160009280601f83011215610f1c5781359367ffffffffffffffff8511610f1957506020808301928560051b010111610167579190565b80fd5b8380fd5b3461016757610f2e36610ea3565b90610f4e60045473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610f5b57005b610f716108d1610f6c838587612fd5565b612fea565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa90811561097657600194889160009361103e575b5082610fe6575b5050505001610f52565b610ff1918391613cd9565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610fdc565b61105791935060203d811161096f57610961818361059f565b9138610fd5565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff6004356110a28161078a565b166000526006602052606060406000205473ffffffffffffffffffffffffffffffffffffffff6040519167ffffffffffffffff8116835260ff8160401c161515602084015260481c166040820152f35b6101bf9092919260a081019373ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757611180612ff4565b506102e560405161119081610567565b60025473ffffffffffffffffffffffffffffffffffffffff808216835260a09190911c60ff16151560208301526003541660408201526112056111e860045473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166060830152565b61124461122760055473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080830152565b604051918291826110f2565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760005473ffffffffffffffffffffffffffffffffffffffff8116330361130f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff6004356113cf8161078a565b166000526006602052600167ffffffffffffffff604060002054160167ffffffffffffffff811161140f5760209067ffffffffffffffff60405191168152f35b61301f565b906020808351928381520192019060005b8181106114325750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611425565b6040906106e39392151581528160208201520190611414565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff6004356114bb8161078a565b1680600052600660205260ff60406000205460401c169060005260066020526001604060002001906040518083602082955493848152019060005260206000209260005b8181106115235750506115149250038361059f565b6102e56040519283928361145e565b84548352600194850194879450602090930192016114ff565b346101675761154a36610ea3565b9061156d6108d160015473ffffffffffffffffffffffffffffffffffffffff1690565b3303611820575b906000905b80821061158257005b61159561159083838661304e565b6130f5565b916115c56115ab845167ffffffffffffffff1690565b67ffffffffffffffff166000526006602052604060002090565b906115d36020850151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1681151560401b68ff00000000000000001617835560408501908151516116e2575b50506000949294926001606086019301935b8351805182101561166f57906116686116626116488360019561317b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b87613dd9565b500161162a565b5050939091600193505190815161168b575b5050019091611579565b6116d867ffffffffffffffff6116ca7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586935167ffffffffffffffff1690565b1692604051918291826124ba565b0390a23880611681565b95939092949560001461180b5760018501939060005b845180518210156117aa57611648826117109261317b565b73ffffffffffffffffffffffffffffffffffffffff81161561175f57906117586117526108d160019473ffffffffffffffffffffffffffffffffffffffff1690565b88614a55565b50016116f8565b6104b56117748a5167ffffffffffffffff1690565b7f463258ff0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b505093509493917f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328167ffffffffffffffff6117ed875167ffffffffffffffff1690565b9251926118016040519283921694826124ba565b0390a23880611618565b6104b5611774875167ffffffffffffffff1690565b6118426108d160055473ffffffffffffffffffffffffffffffffffffffff1690565b3314611574577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101675760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576004356118ad8161078a565b60243567ffffffffffffffff8111610167576118cd90369060040161079c565b606435916044356118dd846109f9565b60025460a01c60ff166120765761192e740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61194c8267ffffffffffffffff166000526006602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff85161561204c578254604081901c60ff16611fc4575b61199c906108d19060481c73ffffffffffffffffffffffffffffffffffffffff1681565b3303611f9a57849273ffffffffffffffffffffffffffffffffffffffff6119d860035473ffffffffffffffffffffffffffffffffffffffff1690565b1680611f1e575b5080611a36611a016119fc611ae2945467ffffffffffffffff1690565b61318f565b82547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff821617909255565b611a97611a416105e0565b6000815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201529167ffffffffffffffff8516604084015267ffffffffffffffff166060830152565b60006080820152600086611aae60208201826131ae565b929096611b25611abe84806131ae565b611b19606087019b611acf8d612fea565b98611b12611ae960408b019d8e8c6131ff565b905061327f565b9f611af26105ef565b9c8d5273ffffffffffffffffffffffffffffffffffffffff1660208d0152565b3691610d82565b60408901523691610d82565b6060860152611b5b611b35610657565b6080870190815273ffffffffffffffffffffffffffffffffffffffff90951660a0870152565b8060c086015260e08501978289526101008601998a52611b9f611b996108d16108d160025473ffffffffffffffffffffffffffffffffffffffff1690565b91612fea565b88611bf3611bbc611bb360808901896131ae565b919098806131ae565b91604051998a98899788977f3a49bb490000000000000000000000000000000000000000000000000000000089526004890161338a565b03915afa95861561097657600091829383918499611ef3575b505252611c23611c1c84896131ff565b36916133e3565b9260005b611c31828a6131ff565b9050811015611c785790600182611c6e81611c5b8e8c8c611c55611c319a8e61317b565b51613ed7565b8c5190611c68838361317b565b5261317b565b5001909150611c27565b838689611ce08d8760008b611ca86108d16108d160025473ffffffffffffffffffffffffffffffffffffffff1690565b86516040518097819482937f01447eaa0000000000000000000000000000000000000000000000000000000084528c600485016135cd565b03915afa92831561097657600093611ece575b5015611def5750611d1460005b60808651019067ffffffffffffffff169052565b60005b825151811015611d455780611d2e6001928461317b565b516080611d3c83875161317b565b51015201611d17565b6102e584611d528761448b565b90611d5c826137c8565b82515281516060015167ffffffffffffffff167f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f3267ffffffffffffffff80604051931693169180611dad8682613662565b0390a3611ddd7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b6040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015260208180604481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1801561097657611d1491600091611e9f575b50611d00565b611ec1915060203d602011611ec7575b611eb9818361059f565b81019061364d565b86611e99565b503d611eaf565b611eec9193503d806000833e611ee4818361059f565b81019061345b565b9186611cf3565b929450975050611f14913d8091833e611f0c818361059f565b81019061332e565b9791939138611c0c565b8094503b1561016757600060405180957fe0a0e506000000000000000000000000000000000000000000000000000000008252818381611f628b8960048401612b15565b03925af1908115610976578694611ae292611f7f575b50906119df565b80611f8e6000611f949361059f565b8061015c565b38611f78565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b612003611fff611fe973ffffffffffffffffffffffffffffffffffffffff89166108d1565b6000908152600287016020526040902054151590565b1590565b15611978577fd0d259760000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff861660045260246000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b91908260a0910312610167576040516120b881610567565b60808082948035845260208101356120cf8161078a565b602085015260408101356120e28161078a565b604085015260608101356120f58161078a565b60608501520135916121068361078a565b0152565b9080601f8301121561016757813591612122836109e1565b92612130604051948561059f565b80845260208085019160051b830101918383116101675760208101915b83831061215c57505050505090565b823567ffffffffffffffff81116101675782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610167576121a46105e0565b906121b160208401610b10565b8252604083013567ffffffffffffffff8111610167578760206121d692860101610db9565b6020830152606083013567ffffffffffffffff8111610167578760206121fe92860101610db9565b60408301526080830135606083015260a08301359167ffffffffffffffff83116101675761223488602080969581960101610db9565b608082015281520192019161214d565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff8111610167576101a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610167576122ba6105ef565b906122c836826004016120a0565b82526122d660a48201610b10565b602083015260c481013567ffffffffffffffff8111610167576122ff9060043691840101610db9565b604083015260e481013567ffffffffffffffff8111610167576123289060043691840101610db9565b606083015261010481013567ffffffffffffffff8111610167576123529060043691840101610db9565b60808301526123646101248201610b10565b60a083015261014481013560c083015261016481013560e08301526101848101359167ffffffffffffffff8311610167576123ab6123b69260046102e5953692010161210a565b6101008201526137c8565b6040519081529081906020820190565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675773ffffffffffffffffffffffffffffffffffffffff600435612416816109f9565b61241e613ab3565b1633811461249057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206106e3928181520190611414565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675761250560043561078a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b356106e38161078a565b63ffffffff81160361016757565b35906101bf82612539565b81601f8201121561016757803590612569826109e1565b92612577604051948561059f565b82845260208085019360051b830101918183116101675760208101935b8385106125a357505050505090565b843567ffffffffffffffff811161016757820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261016757604051916125ef83610567565b602082013567ffffffffffffffff81116101675785602061261292850101610db9565b83526040820135612622816109f9565b602084015261263360608301612547565b604084015260808201359267ffffffffffffffff84116101675760a083612661886020809881980101610db9565b606084015201356080820152815201940193612594565b919091610140818403126101675761268e6105ff565b9261269981836120a0565b845260a082013567ffffffffffffffff811161016757816126bb918401610db9565b602085015260c082013567ffffffffffffffff811161016757816126e0918401610db9565b60408501526126f160e08301610b10565b6060850152610100820135608085015261012082013567ffffffffffffffff8111610167576127209201612552565b60a0830152565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016757016020813591019167ffffffffffffffff821161016757813603831361016757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016757016020813591019167ffffffffffffffff8211610167578160051b3603831361016757565b90602083828152019160208260051b8501019381936000915b8483106128325750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618636030182121561016757602080918760019401906080806129146128b96128ab8680612727565b60a0875260a0870191612777565b73ffffffffffffffffffffffffffffffffffffffff878701356128db816109f9565b168786015263ffffffff60408701356128f381612539565b1660408601526129066060870187612727565b908683036060880152612777565b9301359101529801930193019194939290612822565b906106e391602081528135602082015267ffffffffffffffff60208301356129518161078a565b16604082015267ffffffffffffffff604083013561296e8161078a565b16606082015267ffffffffffffffff606083013561298b8161078a565b16608082015267ffffffffffffffff60808301356129a88161078a565b1660a0820152612a59612a146129d76129c460a0860186612727565b61014060c0870152610160860191612777565b6129e460c0860186612727565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e0870152612777565b92612a42612a2460e08301610b10565b73ffffffffffffffffffffffffffffffffffffffff16610100850152565b6101008101356101208401526101208101906127b6565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152612809565b9081602091031261016757516106e381610a17565b6040513d6000823e3d90fd5b90816020910312610167575190565b9160209082815201919060005b818110612ad45750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735612afd816109f9565b16815260208781013590820152019401929101612ac7565b919067ffffffffffffffff16825260406020830152612b88612b4b612b3a8380612727565b60a0604087015260e0860191612777565b612b586020840184612727565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016060870152612777565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561016757820160208135910167ffffffffffffffff8211610167578160061b360381136101675784612c4f92612c18927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866106e39903016080870152612aba565b92612c45612c2860608301610b10565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b6080810190612727565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152612777565b612c87613ab3565b6000915b8151831015612e6957612c9e838361317b565b5192612cbc612cad828561317b565b515167ffffffffffffffff1690565b9367ffffffffffffffff8516908115612e31577fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590612d13600195969767ffffffffffffffff166000526006602052604060002090565b612dc9612d8f6040612d3c602086015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffff0000000000000000000000000000000000000000ffffffffffffffffff16604882901b7cffffffffffffffffffffffffffffffffffffffff00000000000000000016178555940151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1690151560401b68ff000000000000000016178255565b54612e26612de567ffffffffffffffff83169260401c60ff1690565b60405193849384919273ffffffffffffffffffffffffffffffffffffffff604092959467ffffffffffffffff60608601971685521660208401521515910152565b0390a2019190612c8b565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b915050565b805160248110612e8057506024015190565b7f7946a7cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa801561097657600090612f57575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011612f9e575b81612f716020938361059f565b810103126101675773ffffffffffffffffffffffffffffffffffffffff9051612f99816109f9565b612f3c565b3d9150612f64565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015612fe55760051b0190565b612fa6565b356106e3816109f9565b6040519061300182610567565b60006080838281528260208201528260408201528260608201520152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190811015612fe55760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301821215610167570190565b9080601f830112156101675781356130a5816109e1565b926130b3604051948561059f565b81845260208085019260051b82010192831161016757602001905b8282106130db5750505090565b6020809183356130ea816109f9565b8152019101906130ce565b608081360312610167576040519061310c8261052a565b80356131178161078a565b8252602081013561312781610a17565b6020830152604081013567ffffffffffffffff81116101675761314d903690830161308e565b604083015260608101359067ffffffffffffffff8211610167576131739136910161308e565b606082015290565b8051821015612fe55760209160051b010190565b67ffffffffffffffff1667ffffffffffffffff811461140f5760010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610167570180359067ffffffffffffffff82116101675760200191813603831361016757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610167570180359067ffffffffffffffff821161016757602001918160061b3603831361016757565b6040519061326082610567565b6060608083600081528260208201528260408201526000838201520152565b90613289826109e1565b613296604051918261059f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06132c482946109e1565b019060005b8281106132d557505050565b6020906132e0613253565b828285010152016132c9565b81601f820112156101675780516133028161061d565b92613310604051948561059f565b81845260208284010111610167576106e3916020808501910161066c565b9060808282031261016757815192602083015161334a81610a17565b92604081015167ffffffffffffffff8111610167578361336b9183016132ec565b92606082015167ffffffffffffffff8111610167576106e392016132ec565b9593919273ffffffffffffffffffffffffffffffffffffffff6133d59467ffffffffffffffff6106e39a9894168952166020880152604087015260a0606087015260a0860191612777565b926080818503910152612777565b9291926133ef826109e1565b936133fd604051958661059f565b602085848152019260061b82019181831161016757925b8284106134215750505050565b604084830312610167576020604091825161343b81610583565b8635613446816109f9565b81528287013583820152815201930192613414565b6020818303126101675780519067ffffffffffffffff821161016757019080601f8301121561016757815161348f816109e1565b9261349d604051948561059f565b81845260208085019260051b820101918383116101675760208201905b8382106134c957505050505090565b815167ffffffffffffffff8111610167576020916134ec878480948801016132ec565b8152019101906134ba565b9080602083519182815201916020808360051b8301019401926000915b83831061352357505050505090565b90919293946020806135be837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806135a36135918685015160a08886015260a085019061068f565b6040850151848203604086015261068f565b9260608101516060840152015190608081840391015261068f565b97019301930191939290613514565b9167ffffffffffffffff6135ef921683526060602084015260608301906134f7565b9060408183039101526020808351928381520192019060005b8181106136155750505090565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101613608565b9081602091031261016757516106e38161078a565b906106e391602081526136b160208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660c082015261010061375d6137286136f560408601516101a060e08701526101c086019061068f565b60608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303018587015261068f565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030161012086015261068f565b60a084015173ffffffffffffffffffffffffffffffffffffffff166101408401529260c081015161016084015260e08101516101808401520151906101a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526134f7565b805151806139e3575067ffffffffffffffff6040825101511660405160208101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321835267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604083015260608201523060808201526080815261385660a08261059f565b519020906139dd61387e602083015173ffffffffffffffffffffffffffffffffffffffff1690565b82516139576138ad608061389d606085015167ffffffffffffffff1690565b93015167ffffffffffffffff1690565b9161392b6138d260a088015173ffffffffffffffffffffffffffffffffffffffff1690565b60c088015190604051958694602086019889919367ffffffffffffffff6080948173ffffffffffffffffffffffffffffffffffffffff949998978560a088019b1687521660208601521660408401521660608201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261059f565b51902061392b606084015160208151910120936040810151602081519101209060806101008201516040516139948161392b60208201948561482f565b519020910151602081519101209160405196879560208701998a9260c094919796959260e085019860008652602086015260408501526060840152608083015260a08201520152565b51902090565b7f4c8ebcc00000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6106e39181519067ffffffffffffffff604081602085015116930151169060405173ffffffffffffffffffffffffffffffffffffffff602082019216825260208152613a5d60408261059f565b5190206040519160208301937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f855260408401526060830152608082015260808152613aaa60a08261059f565b51902090614851565b73ffffffffffffffffffffffffffffffffffffffff600154163303613ad457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169183169190911790556101bf91608090613bed8360608301511660049073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b01511660059073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b9160806101bf929493613c8c8161012081019773ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b9073ffffffffffffffffffffffffffffffffffffffff613dab9392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452613d3b60648561059f565b16600080604093845195613d4f868861059f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15613dd0573d613d9c613d938261061d565b9451948561059f565b83523d6000602085013e614d00565b805180613db6575050565b81602080613dcb936101bf9501019101612a8a565b614976565b60609250614d00565b73ffffffffffffffffffffffffffffffffffffffff6106e3921690614b79565b6020818303126101675780519067ffffffffffffffff821161016757016040818303126101675760405191613e2d83610583565b815167ffffffffffffffff81116101675781613e4a9184016132ec565b8352602082015167ffffffffffffffff811161016757613e6a92016132ec565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080613ea4855184602087015260c086019061068f565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b939291613ee2613253565b50602085019182511561417657613f166108d1610e806108d1895173ffffffffffffffffffffffffffffffffffffffff1690565b9373ffffffffffffffffffffffffffffffffffffffff85161580156140eb575b6140885760009291613ffd959697613fa8613fca93613f8b613f6e8951945173ffffffffffffffffffffffffffffffffffffffff1690565b94613f776105e0565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301613e72565b038183875af191821561097657600092614063575b50602082519201519051916140446140286105e0565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b60208401526040830152606082015261405b610657565b608082015290565b6140819192503d806000833e614079818361059f565b810190613df9565b9038614012565b6104b56140a9885173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481895afa90811561097657600091614157575b5015613f36565b614170915060203d6020116109da576109cc818361059f565b38614150565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190610120820182811067ffffffffffffffff821117610546576040526060610100836141cd612ff4565b8152600060208201528260408201528280820152826080820152600060a0820152600060c0820152600060e08201520152565b9061420a826109e1565b614217604051918261059f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061424582946109e1565b019060005b82811061425657505050565b60209060405161426581610567565b6060815260008382015260006040820152606080820152600060808201528282850101520161424a565b9081602091031261016757516106e3816109f9565b9081602091031261016757516106e381612539565b9080602083519182815201916020808360051b8301019401926000915b8383106142e557505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190608080614373614333855160a0865260a086019061068f565b73ffffffffffffffffffffffffffffffffffffffff87870151168786015263ffffffff60408701511660408601526060860151858203606087015261068f565b930151910152970193019301919392906142d6565b906106e391602081526143d760208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b60a061442b6143f7602085015161014060c086015261016085019061068f565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160e086015261068f565b9273ffffffffffffffffffffffffffffffffffffffff60608201511661010084015260808101516101208401520151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526142b9565b906144946141a0565b506144a26080830151612e6e565b6101008301916144b3835151614200565b9360005b84518051821015614605579061452f6144d28260019461317b565b5161456b602082016144e48151614c70565b61455b614505845173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290958691820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561059f565b516020808251830101910161428f565b906145df6145866080830151602080825183010191016142a4565b6145d273ffffffffffffffffffffffffffffffffffffffff60606040860151950151956145b16105e0565b9788521673ffffffffffffffffffffffffffffffffffffffff166020870152565b63ffffffff166040850152565b606083015260808201526145f3828961317b565b526145fe818861317b565b50016144b7565b50509250928251916146b2614625602085015167ffffffffffffffff1690565b936146a161463e604083015167ffffffffffffffff1690565b9161469061465c608061389d606085015167ffffffffffffffff1690565b9361467f6146686105e0565b6000815267ffffffffffffffff909a1660208b0152565b67ffffffffffffffff166040890152565b67ffffffffffffffff166060870152565b67ffffffffffffffff166080850152565b61476b73ffffffffffffffffffffffffffffffffffffffff61392b61471b6146f1602089015173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290928391820190565b604087015161473760608901516020808251830101910161428f565b916147406105ff565b978852602088015260408701521673ffffffffffffffffffffffffffffffffffffffff166060850152565b608083015260a082015261477e826137c8565b81515261478b3082613a10565b6147bd6147a86103c6604085510167ffffffffffffffff90511690565b83516060015167ffffffffffffffff166103ec565b558051907fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a67ffffffffffffffff6148156060614805604087015167ffffffffffffffff1690565b95015167ffffffffffffffff1690565b614829826040519384931696169482614388565b0390a390565b9060206106e39281815201906134f7565b9060206106e39281815201906142b9565b6139dd815180519061490261487d606086015173ffffffffffffffffffffffffffffffffffffffff1690565b61392b614895606085015167ffffffffffffffff1690565b936148af6080808a015192015167ffffffffffffffff1690565b906040519586946020860198899367ffffffffffffffff60809473ffffffffffffffffffffffffffffffffffffffff82959998949960a089019a8952166020880152166040860152606085015216910152565b51902061392b6020840151602081519101209360a060408201516020815191012091015160405161493b8161392b602082019485614840565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b1561497d57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b8054821015612fe55760005260206000200190600090565b91614a51918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b6000828152600182016020526040902054614adf57805490680100000000000000008210156105465782614ac8614a93846001809601855584614a01565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b80548015614b4a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614b1b8282614a01565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014614c67577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840184811161140f578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff850194851161140f576000958583614c1897614c099503614c1e575b505050614ae6565b90600052602052604060002090565b55600190565b614c4e614c4891614c3f614c35614c5e9588614a01565b90549060031b1c90565b92839187614a01565b90614a19565b8590600052602052604060002090565b55388080614c01565b50505050600090565b6020815103614cb357614c8c6020825183010160208301612aab565b73ffffffffffffffffffffffffffffffffffffffff8111908115614cf4575b50614cb35750565b614cf0906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352602060048401818152019061068f565b0390fd5b61040091501038614cab565b91929015614d7b5750815115614d14575090565b3b15614d1d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614d8e5750805190602001fd5b614cf0906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016106d256fea164736f6c634300081a000a",
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

	Address() common.Address
}
