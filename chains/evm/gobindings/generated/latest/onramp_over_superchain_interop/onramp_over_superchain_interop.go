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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extractGasLimit\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMessageId\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reemitInteropMessage\",\"inputs\":[{\"name\":\"interopMessage\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPSuperchainMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExtraArgsTooShort\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceChainSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageDoesNotExist\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MessageIdUnexpectedlySet\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x61010060405234610566576153358038038061001a816105a0565b92833981019080820390610140821261056657608082126105665761003d610581565b90610047816105c5565b82526020810151906001600160a01b03821682036105665760208301918252610072604082016105d9565b926040810193845260a0610088606084016105d9565b6060830190815295607f1901126105665760405160a081016001600160401b0381118282101761056b576040526100c1608084016105d9565b81526100cf60a084016105ed565b602082019081526100e260c085016105d9565b91604081019283526100f660e086016105d9565b936060820194855261010b61010087016105d9565b6080830190815261012087015190966001600160401b03821161056657018a601f82011215610566578051906001600160401b03821161056b5760209b8c6060610159828660051b016105a0565b9e8f8681520194028301019181831161056657602001925b8284106104f8575050505033156104e757600180546001600160a01b0319163317905580516001600160401b03161580156104d5575b80156104c3575b80156104b1575b61048457516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e0819052825190911615801561049f575b8015610495575b61048457815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19861012098606061029f610581565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a16000905b805182101561040b5761032b82826105fa565b51916001600160401b0361033f82846105fa565b5151169283156103f65760008481526006602090815260409182902081840151815494840151600160401b600160e81b03198616604883901b600160481b600160e81b031617901515851b68ff000000000000000016179182905583516001600160401b0390951685526001600160a01b031691840191909152811c60ff1615159082015291926001927fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590606090a20190610318565b8363c35aa79d60e01b60005260045260246000fd5b604051614d109081610625823960805181818161022b0152818161037401528181610bbc01528181611a0f01526137eb015260a051818181610264015281816106bd0152610bf5015260c05181818161028c01528181610c310152611e24015260e0518181816102b401528181610c6d0152612ee00152f35b6306b7c75960e31b60005260046000fd5b5082511515610200565b5084516001600160a01b0316156101f9565b5088516001600160a01b0316156101b5565b5087516001600160a01b0316156101ae565b5086516001600160a01b0316156101a7565b639b15e16f60e01b60005260046000fd5b6060848303126105665760405190606082016001600160401b0381118382101761056b57604052610528856105c5565b82526020850151906001600160a01b03821682036105665782602092836060950152610556604088016105ed565b6040820152815201930192610171565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761056b57604052565b6040519190601f01601f191682016001600160401b0381118382101761056b57604052565b51906001600160401b038216820361056657565b51906001600160a01b038216820361056657565b5190811515820361056657565b805182101561060e5760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146101575780630be49dbf14610152578063181f5a771461014d57806320487ded146101485780632716072b1461014357806327e936f11461013e57806328938e371461013957806348a98aa4146101345780635cb80c5d1461012f5780636def4ce71461012a5780637437ff9f1461012557806379ba5097146101205780638da5cb5b1461011b5780639041be3d14610116578063972b461214610111578063c9b146b31461010c578063df0aa9e914610107578063f2890a2114610102578063f2fde38b146100fd5763fbca3b74146100f857600080fd5b61248a565b612385565b612203565b611831565b6114fb565b611436565b61134a565b6112f8565b61120f565b611108565b61101d565b610edf565b610df4565b610d93565b610aa0565b6109a6565b61060d565b610549565b6102e9565b6101c1565b600091031261016757565b600080fd5b6101bf90929192608081019373ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576000606060405161020081610873565b82815282602082015282604082015201526102e560405161022081610873565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660208301527f0000000000000000000000000000000000000000000000000000000000000000811660408301527f00000000000000000000000000000000000000000000000000000000000000001660608201526040519182918261016c565b0390f35b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757806004016101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101675760248201610368816124ee565b67ffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361048d57506103af306103aa3684612637565b6139e4565b604483019060646103bf836124ee565b9401936103cb856124ee565b90826103eb8367ffffffffffffffff166000526007602052604060002090565b54036104495750505067ffffffffffffffff61043061042a7fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a936124ee565b946124ee565b6104448260405193849316961694826128e9565b0390a3005b7fbf1f53100000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260445260646000fd5b6000fd5b610499610489916124ee565b7ff094888d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b60005b8381106104e25750506000910152565b81810151838201526020016104d2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361052e815180928187528780880191016104cf565b0116010190565b9060206105469281815201906104f2565b90565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576102e56040516105896060826108e8565b602581527f4f6e52616d704f7665725375706572636861696e496e7465726f7020312e362e60208201527f312d64657600000000000000000000000000000000000000000000000000000060408201526040519182916020835260208301906104f2565b67ffffffffffffffff81160361016757565b908160a09103126101675790565b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757600435610648816105ed565b60243567ffffffffffffffff8111610167576106689036906004016105ff565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156107d957600091610815575b506107de576107839160209161074d61073461073460025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401612ae9565b03915afa80156107d9576102e5916000916107aa575b506040519081529081906020820190565b6107cc915060203d6020116107d2575b6107c481836108e8565b810190612a7f565b38610799565b503d6107ba565b612a73565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610837915060203d60201161083d575b61082f81836108e8565b810190612a5e565b38610703565b503d610825565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761088f57604052565b610844565b6060810190811067ffffffffffffffff82111761088f57604052565b60a0810190811067ffffffffffffffff82111761088f57604052565b6040810190811067ffffffffffffffff82111761088f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761088f57604052565b604051906101bf60a0836108e8565b604051906101bf610120836108e8565b604051906101bf60c0836108e8565b604051906101bf6080836108e8565b67ffffffffffffffff811161088f5760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361016757565b8015150361016757565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff8111610167573660238201121561016757806004013590610a0182610966565b90610a0f60405192836108e8565b828252602460606020840194028201019036821161016757602401925b818410610a3e57610a3c83612c53565b005b606084360312610167576020606091604051610a5981610894565b8635610a64816105ed565b815282870135610a738161097e565b838201526040870135610a858161099c565b6040820152815201930192610a2c565b35906101bf8261097e565b346101675760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576000604051610add816108b0565b600435610ae98161097e565b8152602435610af78161099c565b6020820152604435610b088161097e565b6040820152606435610b198161097e565b6060820152608435610b2a8161097e565b6080820152610b37613a87565b73ffffffffffffffffffffffffffffffffffffffff610b6a825173ffffffffffffffffffffffffffffffffffffffff1690565b16158015610cdc575b8015610ccf575b610ca75780610ba97fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f192613ad2565b610bb1610957565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152610ca160405192839283613c06565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5060208101511515610b7a565b50610d01610734606083015173ffffffffffffffffffffffffffffffffffffffff1690565b15610b73565b67ffffffffffffffff811161088f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192610d4d82610d07565b91610d5b60405193846108e8565b829481845281830111610167578281602093846000960137010152565b9080601f830112156101675781602061054693359101610d41565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757610dec610de76020923690600401610d78565b612e42565b604051908152f35b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757610e2e6004356105ed565b6020610e44602435610e3f8161097e565b612e81565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101675760043567ffffffffffffffff81116101675760040160009280601f83011215610edb5781359367ffffffffffffffff8511610ed857506020808301928560051b010111610167579190565b80fd5b8380fd5b3461016757610eed36610e62565b90610f0d60045473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610f1a57005b610f30610734610f2b838587612fa9565b612fbe565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156107d9576001948891600093610ffd575b5082610fa5575b5050505001610f11565b610fb0918391613cad565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610f9b565b61101691935060203d81116107d2576107c481836108e8565b9138610f94565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff600435611061816105ed565b166000526006602052606060406000205473ffffffffffffffffffffffffffffffffffffffff6040519167ffffffffffffffff8116835260ff8160401c161515602084015260481c166040820152f35b6101bf9092919260a081019373ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675761113f612fc8565b506102e560405161114f816108b0565b60025473ffffffffffffffffffffffffffffffffffffffff808216835260a09190911c60ff16151560208301526003541660408201526111c46111a760045473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166060830152565b6112036111e660055473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080830152565b604051918291826110b1565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760005473ffffffffffffffffffffffffffffffffffffffff811633036112ce577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff60043561138e816105ed565b166000526006602052600167ffffffffffffffff604060002054160167ffffffffffffffff81116113ce5760209067ffffffffffffffff60405191168152f35b612ff3565b906020808351928381520192019060005b8181106113f15750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016113e4565b60409061054693921515815281602082015201906113d3565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675767ffffffffffffffff60043561147a816105ed565b1680600052600660205260ff60406000205460401c169060005260066020526001604060002001906040518083602082955493848152019060005260206000209260005b8181106114e25750506114d3925003836108e8565b6102e56040519283928361141d565b84548352600194850194879450602090930192016114be565b346101675761150936610e62565b9061152c61073460015473ffffffffffffffffffffffffffffffffffffffff1690565b33036117df575b906000905b80821061154157005b61155461154f838386613022565b6130c9565b9161158461156a845167ffffffffffffffff1690565b67ffffffffffffffff166000526006602052604060002090565b906115926020850151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1681151560401b68ff00000000000000001617835560408501908151516116a1575b50506000949294926001606086019301935b8351805182101561162e57906116276116216116078360019561314f565b5173ffffffffffffffffffffffffffffffffffffffff1690565b87613dad565b50016115e9565b5050939091600193505190815161164a575b5050019091611538565b61169767ffffffffffffffff6116897fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586935167ffffffffffffffff1690565b169260405191829182612479565b0390a23880611640565b9593909294956000146117ca5760018501939060005b8451805182101561176957611607826116cf9261314f565b73ffffffffffffffffffffffffffffffffffffffff81161561171e579061171761171161073460019473ffffffffffffffffffffffffffffffffffffffff1690565b88614a14565b50016116b7565b6104896117338a5167ffffffffffffffff1690565b7f463258ff0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b505093509493917f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328167ffffffffffffffff6117ac875167ffffffffffffffff1690565b9251926117c0604051928392169482612479565b0390a238806115d7565b610489611733875167ffffffffffffffff1690565b61180161073460055473ffffffffffffffffffffffffffffffffffffffff1690565b3314611533577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101675760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043561186c816105ed565b60243567ffffffffffffffff81116101675761188c9036906004016105ff565b6064359160443561189c8461097e565b60025460a01c60ff16612035576118ed740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61190b8267ffffffffffffffff166000526006602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff85161561200b578254604081901c60ff16611f83575b61195b906107349060481c73ffffffffffffffffffffffffffffffffffffffff1681565b3303611f5957849273ffffffffffffffffffffffffffffffffffffffff61199760035473ffffffffffffffffffffffffffffffffffffffff1690565b1680611edd575b50806119f56119c06119bb611aa1945467ffffffffffffffff1690565b613163565b82547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff821617909255565b611a56611a00610929565b6000815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201529167ffffffffffffffff8516604084015267ffffffffffffffff166060830152565b60006080820152600086611a6d6020820182613182565b929096611ae4611a7d8480613182565b611ad8606087019b611a8e8d612fbe565b98611ad1611aa860408b019d8e8c6131d3565b9050613253565b9f611ab1610938565b9c8d5273ffffffffffffffffffffffffffffffffffffffff1660208d0152565b3691610d41565b60408901523691610d41565b6060860152611b1a611af4612a49565b6080870190815273ffffffffffffffffffffffffffffffffffffffff90951660a0870152565b8060c086015260e08501978289526101008601998a52611b5e611b5861073461073460025473ffffffffffffffffffffffffffffffffffffffff1690565b91612fbe565b88611bb2611b7b611b726080890189613182565b91909880613182565b91604051998a98899788977f3a49bb490000000000000000000000000000000000000000000000000000000089526004890161335e565b03915afa9586156107d957600091829383918499611eb2575b505252611be2611bdb84896131d3565b36916133b7565b9260005b611bf0828a6131d3565b9050811015611c375790600182611c2d81611c1a8e8c8c611c14611bf09a8e61314f565b51613eab565b8c5190611c27838361314f565b5261314f565b5001909150611be6565b838689611c9f8d8760008b611c6761073461073460025473ffffffffffffffffffffffffffffffffffffffff1690565b86516040518097819482937f01447eaa0000000000000000000000000000000000000000000000000000000084528c600485016135a1565b03915afa9283156107d957600093611e8d575b5015611dae5750611cd360005b60808651019067ffffffffffffffff169052565b60005b825151811015611d045780611ced6001928461314f565b516080611cfb83875161314f565b51015201611cd6565b6102e584611d118761445f565b90611d1b8261379c565b82515281516060015167ffffffffffffffff167f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f3267ffffffffffffffff80604051931693169180611d6c8682613636565b0390a3611d9c7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b6040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015260208180604481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af180156107d957611cd391600091611e5e575b50611cbf565b611e80915060203d602011611e86575b611e7881836108e8565b810190613621565b86611e58565b503d611e6e565b611eab9193503d806000833e611ea381836108e8565b81019061342f565b9186611cb2565b929450975050611ed3913d8091833e611ecb81836108e8565b810190613302565b9791939138611bcb565b8094503b1561016757600060405180957fe0a0e506000000000000000000000000000000000000000000000000000000008252818381611f218b8960048401612ae9565b03925af19081156107d9578694611aa192611f3e575b509061199e565b80611f4d6000611f53936108e8565b8061015c565b38611f37565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b611fc2611fbe611fa873ffffffffffffffffffffffffffffffffffffffff8916610734565b6000908152600287016020526040902054151590565b1590565b15611937577fd0d259760000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff861660045260246000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b91908260a091031261016757604051612077816108b0565b608080829480358452602081013561208e816105ed565b602085015260408101356120a1816105ed565b604085015260608101356120b4816105ed565b60608501520135916120c5836105ed565b0152565b9080601f83011215610167578135916120e183610966565b926120ef60405194856108e8565b80845260208085019160051b830101918383116101675760208101915b83831061211b57505050505090565b823567ffffffffffffffff81116101675782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261016757612163610929565b9061217060208401610a95565b8252604083013567ffffffffffffffff81116101675787602061219592860101610d78565b6020830152606083013567ffffffffffffffff8111610167578760206121bd92860101610d78565b60408301526080830135606083015260a08301359167ffffffffffffffff8311610167576121f388602080969581960101610d78565b608082015281520192019161210c565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff8111610167576101a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261016757612279610938565b90612287368260040161205f565b825261229560a48201610a95565b602083015260c481013567ffffffffffffffff8111610167576122be9060043691840101610d78565b604083015260e481013567ffffffffffffffff8111610167576122e79060043691840101610d78565b606083015261010481013567ffffffffffffffff8111610167576123119060043691840101610d78565b60808301526123236101248201610a95565b60a083015261014481013560c083015261016481013560e08301526101848101359167ffffffffffffffff83116101675761236a6123759260046102e595369201016120c9565b61010082015261379c565b6040519081529081906020820190565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675773ffffffffffffffffffffffffffffffffffffffff6004356123d58161097e565b6123dd613a87565b1633811461244f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206105469281815201906113d3565b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576124c46004356105ed565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b35610546816105ed565b63ffffffff81160361016757565b35906101bf826124f8565b81601f820112156101675780359061252882610966565b9261253660405194856108e8565b82845260208085019360051b830101918183116101675760208101935b83851061256257505050505090565b843567ffffffffffffffff811161016757820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261016757604051916125ae836108b0565b602082013567ffffffffffffffff8111610167578560206125d192850101610d78565b835260408201356125e18161097e565b60208401526125f260608301612506565b604084015260808201359267ffffffffffffffff84116101675760a083612620886020809881980101610d78565b606084015201356080820152815201940193612553565b919091610140818403126101675761264d610948565b92612658818361205f565b845260a082013567ffffffffffffffff8111610167578161267a918401610d78565b602085015260c082013567ffffffffffffffff8111610167578161269f918401610d78565b60408501526126b060e08301610a95565b6060850152610100820135608085015261012082013567ffffffffffffffff8111610167576126df9201612511565b60a0830152565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016757016020813591019167ffffffffffffffff821161016757813603831361016757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016757016020813591019167ffffffffffffffff8211610167578160051b3603831361016757565b90602083828152019160208260051b8501019381936000915b8483106127f15750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618636030182121561016757602080918760019401906080806128d361287861286a86806126e6565b60a0875260a0870191612736565b73ffffffffffffffffffffffffffffffffffffffff8787013561289a8161097e565b168786015263ffffffff60408701356128b2816124f8565b1660408601526128c560608701876126e6565b908683036060880152612736565b93013591015298019301930191949392906127e1565b9061054691602081528135602082015267ffffffffffffffff6020830135612910816105ed565b16604082015267ffffffffffffffff604083013561292d816105ed565b16606082015267ffffffffffffffff606083013561294a816105ed565b16608082015267ffffffffffffffff6080830135612967816105ed565b1660a0820152612a186129d361299661298360a08601866126e6565b61014060c0870152610160860191612736565b6129a360c08601866126e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e0870152612736565b92612a016129e360e08301610a95565b73ffffffffffffffffffffffffffffffffffffffff16610100850152565b610100810135610120840152610120810190612775565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526127c8565b60405190612a586020836108e8565b60008252565b9081602091031261016757516105468161099c565b6040513d6000823e3d90fd5b90816020910312610167575190565b9160209082815201919060005b818110612aa85750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735612ad18161097e565b16815260208781013590820152019401929101612a9b565b919067ffffffffffffffff16825260406020830152612b5c612b1f612b0e83806126e6565b60a0604087015260e0860191612736565b612b2c60208401846126e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016060870152612736565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561016757820160208135910167ffffffffffffffff8211610167578160061b360381136101675784612c2392612bec927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866105469903016080870152612a8e565b92612c19612bfc60608301610a95565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b60808101906126e6565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152612736565b612c5b613a87565b6000915b8151831015612e3d57612c72838361314f565b5192612c90612c81828561314f565b515167ffffffffffffffff1690565b9367ffffffffffffffff8516908115612e05577fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590612ce7600195969767ffffffffffffffff166000526006602052604060002090565b612d9d612d636040612d10602086015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffff0000000000000000000000000000000000000000ffffffffffffffffff16604882901b7cffffffffffffffffffffffffffffffffffffffff00000000000000000016178555940151151590565b82547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1690151560401b68ff000000000000000016178255565b54612dfa612db967ffffffffffffffff83169260401c60ff1690565b60405193849384919273ffffffffffffffffffffffffffffffffffffffff604092959467ffffffffffffffff60608601971685521660208401521515910152565b0390a2019190612c5f565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b915050565b805160248110612e5457506024015190565b7f7946a7cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156107d957600090612f2b575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011612f72575b81612f45602093836108e8565b810103126101675773ffffffffffffffffffffffffffffffffffffffff9051612f6d8161097e565b612f10565b3d9150612f38565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015612fb95760051b0190565b612f7a565b356105468161097e565b60405190612fd5826108b0565b60006080838281528260208201528260408201528260608201520152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190811015612fb95760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301821215610167570190565b9080601f8301121561016757813561307981610966565b9261308760405194856108e8565b81845260208085019260051b82010192831161016757602001905b8282106130af5750505090565b6020809183356130be8161097e565b8152019101906130a2565b60808136031261016757604051906130e082610873565b80356130eb816105ed565b825260208101356130fb8161099c565b6020830152604081013567ffffffffffffffff8111610167576131219036908301613062565b604083015260608101359067ffffffffffffffff82116101675761314791369101613062565b606082015290565b8051821015612fb95760209160051b010190565b67ffffffffffffffff1667ffffffffffffffff81146113ce5760010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610167570180359067ffffffffffffffff82116101675760200191813603831361016757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610167570180359067ffffffffffffffff821161016757602001918160061b3603831361016757565b60405190613234826108b0565b6060608083600081528260208201528260408201526000838201520152565b9061325d82610966565b61326a60405191826108e8565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06132988294610966565b019060005b8281106132a957505050565b6020906132b4613227565b8282850101520161329d565b81601f820112156101675780516132d681610d07565b926132e460405194856108e8565b818452602082840101116101675761054691602080850191016104cf565b9060808282031261016757815192602083015161331e8161099c565b92604081015167ffffffffffffffff8111610167578361333f9183016132c0565b92606082015167ffffffffffffffff81116101675761054692016132c0565b9593919273ffffffffffffffffffffffffffffffffffffffff6133a99467ffffffffffffffff6105469a9894168952166020880152604087015260a0606087015260a0860191612736565b926080818503910152612736565b9291926133c382610966565b936133d160405195866108e8565b602085848152019260061b82019181831161016757925b8284106133f55750505050565b604084830312610167576020604091825161340f816108cc565b863561341a8161097e565b815282870135838201528152019301926133e8565b6020818303126101675780519067ffffffffffffffff821161016757019080601f8301121561016757815161346381610966565b9261347160405194856108e8565b81845260208085019260051b820101918383116101675760208201905b83821061349d57505050505090565b815167ffffffffffffffff8111610167576020916134c0878480948801016132c0565b81520191019061348e565b9080602083519182815201916020808360051b8301019401926000915b8383106134f757505050505090565b9091929394602080613592837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806135776135658685015160a08886015260a08501906104f2565b604085015184820360408601526104f2565b926060810151606084015201519060808184039101526104f2565b970193019301919392906134e8565b9167ffffffffffffffff6135c3921683526060602084015260608301906134cb565b9060408183039101526020808351928381520192019060005b8181106135e95750505090565b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015181860152604090940193909201916001016135dc565b908160209103126101675751610546816105ed565b90610546916020815261368560208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660c08201526101006137316136fc6136c960408601516101a060e08701526101c08601906104f2565b60608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086830301858701526104f2565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858303016101208601526104f2565b60a084015173ffffffffffffffffffffffffffffffffffffffff166101408401529260c081015161016084015260e08101516101808401520151906101a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526134cb565b805151806139b7575067ffffffffffffffff6040825101511660405160208101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321835267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604083015260608201523060808201526080815261382a60a0826108e8565b519020906139b1613852602083015173ffffffffffffffffffffffffffffffffffffffff1690565b825161392b6138816080613871606085015167ffffffffffffffff1690565b93015167ffffffffffffffff1690565b916138ff6138a660a088015173ffffffffffffffffffffffffffffffffffffffff1690565b60c088015190604051958694602086019889919367ffffffffffffffff6080948173ffffffffffffffffffffffffffffffffffffffff949998978560a088019b1687521660208601521660408401521660608201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826108e8565b5190206138ff60608401516020815191012093604081015160208151910120906080610100820151604051613968816138ff6020820194856147ee565b519020910151602081519101209160405196879560208701998a9260c094919796959260e085019860008652602086015260408501526060840152608083015260a08201520152565b51902090565b7f4c8ebcc00000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6105469181519067ffffffffffffffff604081602085015116930151169060405173ffffffffffffffffffffffffffffffffffffffff602082019216825260208152613a316040826108e8565b5190206040519160208301937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f855260408401526060830152608082015260808152613a7e60a0826108e8565b51902090614810565b73ffffffffffffffffffffffffffffffffffffffff600154163303613aa857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169183169190911790556101bf91608090613bc18360608301511660049073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b01511660059073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b9160806101bf929493613c608161012081019773ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565b9073ffffffffffffffffffffffffffffffffffffffff613d7f9392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452613d0f6064856108e8565b16600080604093845195613d2386886108e8565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15613da4573d613d70613d6782610d07565b945194856108e8565b83523d6000602085013e614c3b565b805180613d8a575050565b81602080613d9f936101bf9501019101612a5e565b614935565b60609250614c3b565b73ffffffffffffffffffffffffffffffffffffffff610546921690614b38565b6020818303126101675780519067ffffffffffffffff821161016757016040818303126101675760405191613e01836108cc565b815167ffffffffffffffff81116101675781613e1e9184016132c0565b8352602082015167ffffffffffffffff811161016757613e3e92016132c0565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080613e78855184602087015260c08601906104f2565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b939291613eb6613227565b50602085019182511561414a57613eea610734610e3f610734895173ffffffffffffffffffffffffffffffffffffffff1690565b9373ffffffffffffffffffffffffffffffffffffffff85161580156140bf575b61405c5760009291613fd1959697613f7c613f9e93613f5f613f428951945173ffffffffffffffffffffffffffffffffffffffff1690565b94613f4b610929565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301613e46565b038183875af19182156107d957600092614037575b5060208251920151905191614018613ffc610929565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b60208401526040830152606082015261402f612a49565b608082015290565b6140559192503d806000833e61404d81836108e8565b810190613dcd565b9038613fe6565b61048961407d885173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481895afa9081156107d95760009161412b575b5015613f0a565b614144915060203d60201161083d5761082f81836108e8565b38614124565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190610120820182811067ffffffffffffffff82111761088f576040526060610100836141a1612fc8565b8152600060208201528260408201528280820152826080820152600060a0820152600060c0820152600060e08201520152565b906141de82610966565b6141eb60405191826108e8565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06142198294610966565b019060005b82811061422a57505050565b602090604051614239816108b0565b6060815260008382015260006040820152606080820152600060808201528282850101520161421e565b9081602091031261016757516105468161097e565b908160209103126101675751610546816124f8565b9080602083519182815201916020808360051b8301019401926000915b8383106142b957505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190608080614347614307855160a0865260a08601906104f2565b73ffffffffffffffffffffffffffffffffffffffff87870151168786015263ffffffff6040870151166040860152606086015185820360608701526104f2565b930151910152970193019301919392906142aa565b9061054691602081526143ab60208201835167ffffffffffffffff6080809280518552826020820151166020860152826040820151166040860152826060820151166060860152015116910152565b60a06143ff6143cb602085015161014060c08601526101608501906104f2565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160e08601526104f2565b9273ffffffffffffffffffffffffffffffffffffffff60608201511661010084015260808101516101208401520151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08285030191015261428d565b6144719061446b614174565b50614c2f565b9061447f6080830151612e42565b6101008301916144908351516141d4565b9360005b845180518210156145d857906144fb6144af8260019461314f565b516145276144d1825173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290938491820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836108e8565b61453e602082015160208082518301019101614263565b906145b2614559608083015160208082518301019101614278565b6145a573ffffffffffffffffffffffffffffffffffffffff6060604086015195015195614584610929565b9788521673ffffffffffffffffffffffffffffffffffffffff166020870152565b63ffffffff166040850152565b606083015260808201526145c6828961314f565b526145d1818861314f565b5001614494565b50509250928251916146856145f8602085015167ffffffffffffffff1690565b93614674614611604083015167ffffffffffffffff1690565b9161466361462f6080613871606085015167ffffffffffffffff1690565b9361465261463b610929565b6000815267ffffffffffffffff909a1660208b0152565b67ffffffffffffffff166040890152565b67ffffffffffffffff166060870152565b67ffffffffffffffff166080850152565b61473e73ffffffffffffffffffffffffffffffffffffffff6138ff6146ee6146c4602089015173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290928391820190565b604087015161470a606089015160208082518301019101614263565b91614713610948565b978852602088015260408701521673ffffffffffffffffffffffffffffffffffffffff166060850152565b608083015260a08201526147518261379c565b81515261475e30826139e4565b81516060015167ffffffffffffffff166000908152600760205260409020558051907fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a67ffffffffffffffff6147d460606147c4604087015167ffffffffffffffff1690565b95015167ffffffffffffffff1690565b6147e882604051938493169616948261435c565b0390a390565b9060206105469281815201906134cb565b90602061054692818152019061428d565b6139b181518051906148c161483c606086015173ffffffffffffffffffffffffffffffffffffffff1690565b6138ff614854606085015167ffffffffffffffff1690565b9361486e6080808a015192015167ffffffffffffffff1690565b906040519586946020860198899367ffffffffffffffff60809473ffffffffffffffffffffffffffffffffffffffff82959998949960a089019a8952166020880152166040860152606085015216910152565b5190206138ff6020840151602081519101209360a06040820151602081519101209101516040516148fa816138ff6020820194856147ff565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b1561493c57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b8054821015612fb95760005260206000200190600090565b91614a10918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b6000828152600182016020526040902054614a9e578054906801000000000000000082101561088f5782614a87614a528460018096018555846149c0565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b80548015614b09577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614ada82826149c0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001810191806000528260205260406000205492831515600014614c26577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116113ce578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116113ce576000958583614bd797614bc89503614bdd575b505050614aa5565b90600052602052604060002090565b55600190565b614c0d614c0791614bfe614bf4614c1d95886149c0565b90549060031b1c90565b928391876149c0565b906149d8565b8590600052602052604060002090565b55388080614bc0565b50505050600090565b614c37614174565b5090565b91929015614cb65750815115614c4f575090565b3b15614c585790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614cc95750805190602001fd5b614cff906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610535565b0390fdfea164736f6c634300081a000a",
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
