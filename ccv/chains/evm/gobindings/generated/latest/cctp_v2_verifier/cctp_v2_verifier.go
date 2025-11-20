// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_v2_verifier

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

type BaseVerifierAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type BaseVerifierDestChainConfigArgs struct {
	Router             common.Address
	DestChainSelector  uint64
	AllowlistEnabled   bool
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}

type CCTPV2VerifierDynamicConfig struct {
	FeeAggregator  common.Address
	AllowlistAdmin common.Address
}

type CCTPV2VerifierFinalityConfig struct {
	DefaultCCTPFinalityThreshold uint16
	DefaultCCTPFinalityBps       uint16
	CustomCCIPFinalities         []uint16
	CustomCCTPFinalityThresholds []uint16
	CustomCCTPFinalityBps        []uint16
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

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	ExecutionGasLimit   uint32
	CcipReceiveGasLimit uint32
	Finality            uint16
	CcvAndExecutorHash  [32]byte
	OnRampAddress       []byte
	OffRampAddress      []byte
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []MessageV1CodecTokenTransferV1
	Data                []byte
}

type MessageV1CodecTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	TokenReceiver      []byte
	ExtraData          []byte
}

var CCTPV2VerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CCTP_V2_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIGNATURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalityConfig\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPV2Verifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPV2Verifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPV2Verifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CCVVersionMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomFinalitiesMustBeStrictlyIncreasing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomFinalityArraysMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageIdMismatch\",\"inputs\":[{\"name\":\"computed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"attested\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MissingCustomFinalities\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101208060405234610b715761543d803803809161001d8285610e8d565b83398101908082039060e08212610b71578051906001600160a01b038216808303610b71576020820151906001600160a01b038216808303610b71576040840151936001600160a01b03851692838603610b715760608201516001600160401b038111610b715782019289601f85011215610b715783519361009e85610eb0565b946100ac6040519687610e8d565b8086528b60208284010111610b71576100cb9160208088019101610ecb565b60808301516001600160401b038111610b715783019660a0888c0312610b71576040519860a08a016001600160401b0381118b8210176109135760405261011189610eee565b8a5261011f60208a01610eee565b60208b0190815260408a01519098906001600160401b038111610b71578d610148918c01610efd565b60408c0190815260608b0151909d906001600160401b038111610b715781610171918d01610efd565b60608d0190815260808c0151909b90916001600160401b038311610b715760409261019c9201610efd565b60808d019081529c609f190112610b715760408051969087016001600160401b03811188821017610913576101e59160c0916040526101dd60a08201610f6c565b895201610f6c565b96602087019788523315610e7c57600180546001600160a01b0319163317905560405160035490919060008361021a83611015565b80825291600184168015610e5e57600114610dfe575b61023c92500384610e8d565b8151906001600160401b0382116109135761025690611015565b601f8111610da4575b506020601f8211600114610d27576102b8928260008051602061541d83398151915295936102c693600091610d1c575b508160011b916000199060031b1c1916176003555b60405193849360408552604085019061104f565b90838203602085015261104f565b0390a184158015610d14575b8015610d0c575b610ae457604051639cdbb18160e01b8152602081600481895afa8015610bc65763ffffffff91600091610ced575b501660018103610cd45750604051632c12192160e01b8152602081600481895afa908115610bc657600091610c9a575b5060405163054fd4d560e41b81526001600160a01b03919091169290602081600481875afa8015610bc65763ffffffff91600091610c7b575b501660018103610c6257506040516367e0ed8360e11b8152602081600481895afa908115610bc657600091610c19575b506001600160a01b0316838103610c01575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa908115610bc657600091610bd2575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529590602081604481855afa908115610bc657600091610b94575b506000198101809111610b7e576104da9060405190602082019863095ea7b360e01b8a526024830152604482015260448152610468606482610e8d565b6000806040998a519361047b8c86610e8d565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af13d15610b76573d906104bc82610eb0565b916104c98b519384610e8d565b82523d6000602084013e5b84611074565b805180610af5575b5050916080917fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c9363ffffffff61010051169188519384526020840152878301526060820152a180516001600160a01b031615610ae45751600580546001600160a01b03199081166001600160a01b03938416908117909255835160068054909216908416179055835190815291511660208201527f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b908290a185515115610ad357855151835151811490811591610ac6575b50610ab557600093845b875180518710156106115761ffff6105dc8863ffffffff93610f9c565b51169116101561060057600161ffff6105f6878a51610f9c565b51169501946105bf565b633fd3668360e21b60005260046000fd5b505085848861ffff8451166007549063ffff0000885160101b169163ffffffff19161717600755805180519060018060401b038211610913576801000000000000000082116109135760209060085483600855808410610a7f575b500190600860005260206000208160041c9160005b838110610a3e5750600f1981169003806109ef575b505083518051925090506001600160401b0382116109135768010000000000000000821161091357602090600954836009558084106109b9575b500190600960005260206000208160041c9160005b8381106109785750600f198116900380610929575b50508451805198925090506001600160401b0388116109135768010000000000000000881161091357602090600a5489600a55898181106108bf575b50500196600a60005260206000208160041c9160005b8381106108755750600f19811690038061081b575b887fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc02675461ffff8a6107e38b6107d08c6107bd8d878e8b51998a9960208b52511660208a015251168a8801525160a0606088015260c0870190610fdd565b9051858203601f19016080870152610fdd565b9051838203601f190160a0850152610fdd565b0390a15161430b9081611112823960805181505060a05181611f23015260c05181610dc0015260e0518161202c015261010051815050f35b9860009960005b8181106108425750505001969096559394859490826107bd61ffff610761565b90919a602061086b60019261ffff8f51169085851b61ffff809160031b9316831b921b19161790565b9c01929101610822565b6000805b6010811061088e57508382015560010161074c565b9b9060206108b68e60019361ffff86511691851b61ffff809160031b9316831b921b19161790565b92019c01610879565b6108f091600a600052600f8560002091601e82850160041c84019460011b16806108f7575b500160041c0190610fc6565b8989610736565b600019850190815490600019908a0360031b1c1690558e6108e4565b634e487b7160e01b600052604160045260246000fd5b9260009360005b818110610945575050500155868080806106fa565b909194602061096e60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101610930565b6000805b601081106109915750838201556001016106e5565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161097c565b6109e990600960005283600020600f80870160041c820192601e8860011b16806108f757500160041c0190610fc6565b896106d0565b9260009360005b818110610a0b57505050015586808080610696565b9091946020610a3460019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b96019291016109f6565b6000805b60108110610a57575083820155600101610681565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601610a42565b610aaf90600860005283600020600f80870160041c820192601e8860011b16806108f757500160041c0190610fc6565b8961066c565b630cacb92960e41b60005260046000fd5b90508551511415386105b5565b636c33513960e11b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b8160209181010312610b715760200151801590811503610b7157610b1a5738806104e2565b855162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b6060906104d4565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610bbe575b81610baf60209383610e8d565b81010312610b7157513861042b565b3d9150610ba2565b6040513d6000823e3d90fd5b610bf4915060203d602011610bfa575b610bec8183610e8d565b810190610f80565b386103e1565b503d610be2565b836383395ca960e01b60005260045260245260446000fd5b6020813d602011610c5a575b81610c3260209383610e8d565b81010312610c565751906001600160a01b0382168203610c535750386103a0565b80fd5b5080fd5b3d9150610c25565b6331b6aa1b60e11b600052600160045260245260446000fd5b610c94915060203d602011610bfa57610bec8183610e8d565b38610370565b90506020813d602011610ccc575b81610cb560209383610e8d565b81010312610b7157610cc690610f6c565b38610337565b3d9150610ca8565b633785f8f160e01b600052600160045260245260446000fd5b610d06915060203d602011610bfa57610bec8183610e8d565b38610307565b5087156102d9565b5083156102d2565b90508201513861028f565b601f198216906003600052806000209160005b818110610d8c5750836102c6936102b8969360008051602061541d833981519152989660019410610d73575b5050811b016003556102a4565b84015160001960f88460031b161c191690553880610d66565b91926020600181928689015181550194019201610d3a565b6003600052610dee907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610df4575b601f0160051c0190610fc6565b3861025f565b9091508190610de1565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310610e4257505090602061023c92820101610230565b6020919350806001915483858a01015201910190918592610e2a565b5050602061023c9260ff19851682840152151560051b820101610230565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761091357604052565b6001600160401b03811161091357601f01601f191660200190565b60005b838110610ede5750506000910152565b8181015183820152602001610ece565b519061ffff82168203610b7157565b9080601f83011215610b71578151916001600160401b038311610913578260051b9060405193610f306020840186610e8d565b8452602080850192820101928311610b7157602001905b828210610f545750505090565b60208091610f6184610eee565b815201910190610f47565b51906001600160a01b0382168203610b7157565b90816020910312610b71575163ffffffff81168103610b715790565b8051821015610fb05760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b818110610fd1575050565b60008155600101610fc6565b906020808351928381520192019060005b818110610ffb5750505090565b825161ffff16845260209384019390920191600101610fee565b90600182811c92168015611045575b602083101461102f57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611024565b9060209161106881518092818552858086019101610ecb565b601f01601f1916010190565b919290156110d65750815115611088575090565b3b156110915790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156110e95750805190602001fd5b60405162461bcd60e51b81526020600482015290819061110d90602483019061104f565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461339157508063181f5a771461331257806331b4ec5f14612b5357806338a9eb2a1461248d5780633927ac2e146124525780633bbbed4b14611cd7578063540bc5ea14611c9d5780635cb80c5d146119bc5780636def4ce7146118e75780637437ff9f1461182857806379ba50971461174357806380485e25146114b7578063869b7f62146113355780638da5cb5b146112e3578063b2bd751c14610f86578063bff0ec1d146109c0578063c9b146b31461058c578063ceac5cee146102a2578063f2fde38b146101b5578063fe163eed1461015c5763fec888af1461010857600080fd5b3461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957610155610141613b81565b60405191829160208352602083019061354d565b0390f35b80fd5b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760206040517fb4161002000000000000000000000000000000000000000000000000000000008152f35b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595773ffffffffffffffffffffffffffffffffffffffff610202613752565b61020a613c58565b1633811461027a57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff811161058857366023820112156105885761030390369060248160040135910161386b565b61030b613c58565b610313613b81565b90805167ffffffffffffffff811161055b57610330600354613b2e565b601f81116104c4575b506020601f82116001146103df576103c092827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd895936103ce9388916103d4575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b60405193849360408552604085019061354d565b90838203602085015261354d565b0390a180f35b90508201513861037a565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b8181106104ac5750836103ce936103c096937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610475575b5050811b016003556103ac565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610468565b9192602060018192868901518155019401920161042a565b61052d9060038652601f830160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208410610533575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190613cb7565b38610339565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b91506104ff565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5080fd5b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff8111610588576105dc9036906004016137c4565b73ffffffffffffffffffffffffffffffffffffffff600154163314158061099e575b61097657919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b84811015610972578060051b8201358381121561096e5782019160808336031261096e57604051946080860186811067ffffffffffffffff821117610941576040526106788461380c565b865261068660208501613b09565b9660208701978852604085013567ffffffffffffffff811161093d576106af9036908701613cce565b9460408801958652606081013567ffffffffffffffff8111610939576106d791369101613cce565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610751818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610811575b5095976001019550815b855180518210156107a2579061079b73ffffffffffffffffffffffffffffffffffffffff61079383600195613ca3565b5116896140cf565b5001610763565b505095909694506001929193519081516107c2575b50500193929361062d565b61080767ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613821565b0390a238806107b7565b9893959296919094979860001461090257600184019591875b865180518210156108a7576108548273ffffffffffffffffffffffffffffffffffffffff92613ca3565b5116801561087057906108696001928a61403e565b500161082a565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816108f867ffffffffffffffff8a51169251604051918291602083526020830190613821565b0390a23880610759565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b8280fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8480fd5b8380f35b6004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff600654163314156105fe565b50346101595760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff8111610588576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261058857604051906101c0820182811067ffffffffffffffff82111761055b57604052610a5a8160040161380c565b8252610a686024820161380c565b6020830152610a796044820161380c565b6040830152610a8a606482016138bd565b6060830152610a9b608482016138bd565b6080830152610aac60a482016135c2565b60a083015260c481013560c083015260e481013567ffffffffffffffff811161093957610adf90600436918401016138a2565b60e083015261010481013567ffffffffffffffff811161093957610b0990600436918401016138a2565b61010083015261012481013567ffffffffffffffff811161093957610b3490600436918401016138a2565b61012083015261014481013567ffffffffffffffff811161093957610b5f90600436918401016138a2565b61014083015261016481013567ffffffffffffffff811161093957610b8a90600436918401016138a2565b61016083015261018481013567ffffffffffffffff81116109395781013660238201121561093957600481013590610bc1826135d1565b91610bcf60405193846134d2565b80835260051b810160240160208301368211610f825760248301905b828210610f4c57505050506101808301526101a481013567ffffffffffffffff8111610939576101a0916004610c2492369201016138a2565b91015260243560443567ffffffffffffffff811161093d57610c4a903690600401613796565b6101e18110610f245780600411610939577fffffffff000000000000000000000000000000000000000000000000000000008235167fb4161002000000000000000000000000000000000000000000000000000000008103610ed5575080610180116109395760049261017c8301357fffffffff00000000000000000000000000000000000000000000000000000000167f4be9effe000000000000000000000000000000000000000000000000000000008101610e875750816101a01161096e5761018083013590818103610e5a5750508261019c92610da66020937ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe606101a08401910160405196879586957f57ecfd28000000000000000000000000000000000000000000000000000000008752604082880152826044880152016064860137886102008501526102048401916102006024860152613aca565b03818673ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610e4f578391610e20575b5015610dfa575080f35b907fbc40f556000000000000000000000000000000000000000000000000000000008152fd5b610e42915060203d602011610e48575b610e3a81836134d2565b810190613b16565b38610df0565b503d610e30565b6040513d85823e3d90fd5b7ff1e2ee760000000000000000000000000000000000000000000000000000000086528452602452604484fd5b7ffc99ccbf0000000000000000000000000000000000000000000000000000000086527fb4161002000000000000000000000000000000000000000000000000000000008552602452604485fd5b7ffc99ccbf0000000000000000000000000000000000000000000000000000000085527fb416100200000000000000000000000000000000000000000000000000000000600452602452604484fd5b6004847fbba6473c000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610f7e57602091610f73839283600436928a0101016138ce565b815201910190610beb565b8880fd5b8680fd5b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff811161058857366023820112156105885780600401359067ffffffffffffffff821161093d57602460c0830282010136811161093957611000613c58565b611009836135d1565b9261101760405194856134d2565b83526024602084019201915b818310611246578480855b8051831015611242576110418382613ca3565b519267ffffffffffffffff60206110588385613ca3565b51015116938415611216578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156111ea5773ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a201919061102e565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261096e576040519061125d8261349a565b83359073ffffffffffffffffffffffffffffffffffffffff82168203610f82578260209260c0945261129083870161380c565b838201526112a060408701613b09565b60408201526112b1606087016135c2565b60608201526112c2608087016138bd565b60808201526112d360a087016138bd565b60a0820152815201920191611023565b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101595760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957604051611371816134b6565b611379613752565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361093d57602082019081526113aa613c58565b73ffffffffffffffffffffffffffffffffffffffff8251161561148f578173ffffffffffffffffffffffffffffffffffffffff6103ce92817f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346101595760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610159576114ef6137f5565b60243567ffffffffffffffff811161093d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261093d576040519061153a8261344f565b806004013567ffffffffffffffff811161096e5761155e90600436918401016138a2565b8252602481013567ffffffffffffffff811161096e5761158490600436918401016138a2565b6020830152604481013567ffffffffffffffff811161096e5781013660238201121561096e5760048101356115b8816135d1565b916115c660405193846134d2565b818352602060048185019360061b830101019036821161173f57602401915b81831061170757505050604083015261160060648201613775565b6060830152608481013567ffffffffffffffff811161096e57608091600461162b92369201016138a2565b9101526044359067ffffffffffffffff821161093d5761165867ffffffffffffffff9236906004016138a2565b506116616135ac565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff604082205416156116dc57816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b60408336031261173f5760206040918251611721816134b6565b61172a86613775565b815282860135838201528152019201916115e5565b8780fd5b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957805473ffffffffffffffffffffffffffffffffffffffff81163303611800577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610159576020604051611865816134b6565b8281520152610155604051611879816134b6565b73ffffffffffffffffffffffffffffffffffffffff60055416815273ffffffffffffffffffffffffffffffffffffffff60065416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595767ffffffffffffffff6119286137f5565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b8181106119a65773ffffffffffffffffffffffffffffffffffffffff856101558861197e818903826134d2565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613821565b8254845260209093019260019283019201611951565b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff811161058857611a0c9036906004016137c4565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b83811015611c99578060051b8201359073ffffffffffffffffffffffffffffffffffffffff8216809203611c9557604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315611c8a578793611c57575b5082611aad575b506001915001611a29565b8460405193611b6860208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611af86064826134d2565b8a80604098895193611b0a8b866134d2565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d15611c4f573d90611b4b82613513565b91611b588a5193846134d2565b82523d8d602084013e5b86614232565b805180611ba4575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338611aa2565b611bbb929495969350602080918301019101613b16565b15611bcc5792919085903880611b70565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090611b62565b9092506020813d8211611c82575b81611c72602093836134d2565b81010312610f8257519138611a9b565b3d9150611c65565b6040513d89823e3d90fd5b8580fd5b8480f35b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957602060405160418152f35b50346101595760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610159576004359067ffffffffffffffff82116101595781600401906101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc843603011261015957611d5461372f565b5060843567ffffffffffffffff811161058857611d75903690600401613796565b5050611d856101248401836139a8565b90357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116916014811061241d575b505060601c92602481019367ffffffffffffffff611dd1866139f9565b1680845260026020526040842090815490604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff86165afa80156124125786906123af575b73ffffffffffffffffffffffffffffffffffffffff91501633036123835760f01c60ff1661233b575b505067ffffffffffffffff611e77856139f9565b1682526004602052604082209060028201549460ff8660201c16156122fd575061018481016001611ea88287613a0e565b9050036122c757611ecc611ec7611ebf8388613a0e565b369291613a62565b6138ce565b906040820151602081519101517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612292575b505060601c9573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169687810361226757506020611f6e611f64611f5e8585613a0e565b90613a62565b60808101906139a8565b905003612216575050600183015480156121f65760a490915b519201359261ffff841680940361096e57549260405192611fa78461349a565b8352602083019182526040830190602435825260608401948552611fde608085019180835263ffffffff60a087019a168a52613d33565b92909192156121c4575061ffff855191169081810291818304149015171561219757612710900463ffffffff811161216c5763ffffffff73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016955199511693519551925192604051937fb41610020000000000000000000000000000000000000000000000000000000060208601526024850152602484526120966044856134d2565b853b1561173f579263ffffffff88999a9381612113948b976040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c015260248b015260448a0152606489015260848801521660a48601521660c484015261010060e484015261010483019061354d565b03925af191821561215f57816101559361214f575b5050604051906121396020836134d2565b815260405191829160208352602083019061354d565b612158916134d2565b3881612128565b50604051903d90823e3d90fd5b7fb6f15d0f000000000000000000000000000000000000000000000000000000008752600452602486fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b517f33eae2c200000000000000000000000000000000000000000000000000000000885263ffffffff16600452602487fd5b50608081015160208180518101031261096e57602060a491015191611f87565b611f5e61222692611f6492613a0e565b6122636040519283927fa3c8cf09000000000000000000000000000000000000000000000000000000008452602060048501526024840191613aca565b0390fd5b7f961c9a4f000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880611f06565b836122d460249287613a0e565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8367ffffffffffffffff6123126024936139f9565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b600082815260029091016020526040902054156123585780611e63565b7fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d60201161240a575b816123c9602093836134d2565b81010312611c95575173ffffffffffffffffffffffffffffffffffffffff81168103611c955773ffffffffffffffffffffffffffffffffffffffff90611e3a565b3d91506123bc565b6040513d88823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880611db4565b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957602060405161019c8152f35b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957606060806040516124cc8161344f565b838152836020820152826040820152828082015201526040516124ee8161344f565b61ffff600754818116835260101c166020820152604051600854808252816020810160088652602086209286905b80600f830110612a6b576125b294549181811061288c575b818110612874575b81811061285d575b818110612845575b81811061282d575b818110612815575b8181106127fd575b8181106127e5575b8181106127cd575b8181106127b5575b81811061279d575b818110612785575b81811061276d575b818110612755575b81811061273d575b1061272f575b5003826134d2565b6040820152604051600954808252816020810160098652602086209286905b80600f8301106129835761265794549181811061288c578181106128745781811061285d578181106128455781811061282d57818110612815578181106127fd578181106127e5578181106127cd578181106127b55781811061279d578181106127855781811061276d578181106127555781811061273d571061272f575003826134d2565b6060820152604051600a80548083529084526020820193907fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a85b81600f8401106128a1579461271e92849261015597549181811061288c578181106128745781811061285d578181106128455781811061282d57818110612815578181106127fd578181106127e5578181106127cd578181106127b55781811061279d578181106127855781811061276d578181106127555781811061273d571061272f575003826134d2565b608082015260405191829182613686565b60f01c8152602001386125aa565b92602060019161ffff8560e01c1681520193016125a4565b92602060019161ffff8560d01c16815201930161259c565b92602060019161ffff8560c01c168152019301612594565b92602060019161ffff8560b01c16815201930161258c565b92602060019161ffff8560a01c168152019301612584565b92602060019161ffff8560901c16815201930161257c565b92602060019161ffff8560801c168152019301612574565b92602060019161ffff8560701c16815201930161256c565b92602060019161ffff8560601c168152019301612564565b92602060019161ffff8560501c16815201930161255c565b92602060019161ffff8560401c168152019301612554565b92602060019161ffff8560301c16815201930161254c565b92602060019161ffff85831c168152019301612544565b92602060019161ffff8560101c16815201930161253c565b92602060019161ffff85168152019301612534565b946001610200601092885461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019601920191612691565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e08201520194019201849293916125d1565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e082015201940192018492939161251c565b50346101595760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101595760043567ffffffffffffffff81116105885760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105885760405191612bce8361344f565b612bda826004016135c2565b8352612be8602483016135c2565b9260208101938452604483013567ffffffffffffffff811161093d57612c1490600436918601016135e9565b9360408201948552606484013567ffffffffffffffff811161093957612c4090600436918701016135e9565b9360608301948552608481013567ffffffffffffffff811161096e57612c6b913691016004016135e9565b9160808101928352612c7b613c58565b855151156132ea578551518551518114908115916132dd575b506132b5578392835b87518051861015612d065761ffff612cba8763ffffffff93613ca3565b511691161015612cde57600161ffff612cd4868a51613ca3565b5116940193612c9d565b6004857fff4d9a0c000000000000000000000000000000000000000000000000000000008152fd5b50508587869461ffff8551167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600754935160101b16921617176007555180519067ffffffffffffffff82116132885768010000000000000000821161328857602090600854836008558084106131fd575b50019060088652602086208160041c91875b8381106131bd57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380613170575b505050505180519067ffffffffffffffff82116131435768010000000000000000821161314357602090600954836009558084106130b8575b50019060098552602085208160041c91865b83811061307857507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0811690038061302b575b505050505180519067ffffffffffffffff821161055b5768010000000000000000821161055b57602090600a5483600a55808410612f9f575b500190600a8452602084208160041c91855b838110612f5f57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380612ef0575b857fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc0267546103ce8760405191829182613686565b928593865b818110612f2c5750505001556103ce7fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc0267548480612ebe565b9091946020612f5560019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101612ef5565b86875b60108110612f77575083820155600101612e8b565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601612f62565b612fce90600a8752838720600f80870160041c820192601e8860011b1680612fd4575b500160041c0190613cb7565b85612e79565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558a612fc2565b928693875b81811061304557505050015583808080612e40565b909194602061306e60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101613030565b87885b60108110613090575083820155600101612e0d565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161307b565b6130e69060098852838820600f80870160041c820192601e8860011b16806130ec57500160041c0190613cb7565b86612dfb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558b612fc2565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b928793885b81811061318a57505050015584808080612dc2565b90919460206131b360019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101613175565b88895b601081106131d5575083820155600101612d8f565b865190969160019160209161ffff60048b901b81811b199092169216901b17920196016131c0565b61322b9060088952838920600f80870160041c820192601e8860011b168061323157500160041c0190613cb7565b87612d7d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558c612fc2565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6004847fcacb9290000000000000000000000000000000000000000000000000000000008152fd5b9050835151141538612c94565b6004847fd866a272000000000000000000000000000000000000000000000000000000008152fd5b503461015957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015957506101556040516133536040826134d2565b601881527f434354505632566572696669657220312e372e302d6465760000000000000000602082015260405191829160208352602083019061354d565b9050346105885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610588576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361093d57602092507ffacbd7dc000000000000000000000000000000000000000000000000000000008114908115613425575b5015158152f35b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861341e565b60a0810190811067ffffffffffffffff82111761346b57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff82111761346b57604052565b6040810190811067ffffffffffffffff82111761346b57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761346b57604052565b67ffffffffffffffff811161346b57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106135975750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613558565b6064359061ffff821682036135bd57565b600080fd5b359061ffff821682036135bd57565b67ffffffffffffffff811161346b5760051b60200190565b9080601f830112156135bd578135613600816135d1565b9261360e60405194856134d2565b81845260208085019260051b8201019283116135bd57602001905b8282106136365750505090565b60208091613643846135c2565b815201910190613629565b906020808351928381520192019060005b81811061366c5750505090565b825161ffff1684526020938401939092019160010161365f565b9061372c916020815261ffff825116602082015261ffff602083015116604082015260806136f96136c6604085015160a0606086015260c085019061364e565b60608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858303018486015261364e565b9201519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08285030191015261364e565b90565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036135bd57565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036135bd57565b359073ffffffffffffffffffffffffffffffffffffffff821682036135bd57565b9181601f840112156135bd5782359167ffffffffffffffff83116135bd57602083818601950101116135bd57565b9181601f840112156135bd5782359167ffffffffffffffff83116135bd576020808501948460051b0101116135bd57565b6004359067ffffffffffffffff821682036135bd57565b359067ffffffffffffffff821682036135bd57565b906020808351928381520192019060005b81811061383f5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613832565b92919261387782613513565b9161388560405193846134d2565b8294818452818301116135bd578281602093846000960137010152565b9080601f830112156135bd5781602061372c9335910161386b565b359063ffffffff821682036135bd57565b919060c0838203126135bd57604051906138e78261349a565b819380358352602081013567ffffffffffffffff81116135bd578261390d9183016138a2565b6020840152604081013567ffffffffffffffff81116135bd57826139329183016138a2565b6040840152606081013567ffffffffffffffff81116135bd57826139579183016138a2565b6060840152608081013567ffffffffffffffff81116135bd578261397c9183016138a2565b608084015260a08101359167ffffffffffffffff83116135bd5760a0926139a392016138a2565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156135bd570180359067ffffffffffffffff82116135bd576020019181360383136135bd57565b3567ffffffffffffffff811681036135bd5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156135bd570180359067ffffffffffffffff82116135bd57602001918160051b360383136135bd57565b9015613a9b578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156135bd570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b359081151582036135bd57565b908160209103126135bd575180151581036135bd5790565b90600182811c92168015613b77575b6020831014613b4857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613b3d565b6040519060008260035491613b9583613b2e565b8083529260018116908115613c1b5750600114613bbb575b613bb9925003836134d2565b565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310613bff575050906020613bb992820101613bad565b6020919350806001915483858901015201910190918492613be7565b60209250613bb99491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101613bad565b73ffffffffffffffffffffffffffffffffffffffff600154163303613c7957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015613a9b5760209160051b010190565b818110613cc2575050565b60008155600101613cb7565b9080601f830112156135bd578135613ce5816135d1565b92613cf360405194856134d2565b81845260208085019260051b8201019283116135bd57602001905b828210613d1b5750505090565b60208091613d2884613775565b815201910190613d0e565b63ffffffff1680613d5457506007549061ffff8083169260101c1690600190565b90604051600854808252816020810160086000526020600020926000905b80600f830110613f3e57613df894549181811061288c578181106128745781811061285d578181106128455781811061282d57818110612815578181106127fd578181106127e5578181106127cd578181106127b55781811061279d578181106127855781811061276d578181106127555781811061273d571061272f575003826134d2565b60005b815180821015613f2e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111613eff5781148015613ee8575b613e4457600101613dfb565b9250506000600954831015613ebb5780600960209252208260041c809101601e8460011b1690549160f08560041b1694600090600a541115613ebb57600a90527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8015461ffff92851c8316941c9091169160019150565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b5061ffff613ef68284613ca3565b51168410613e38565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050509050600090600090600090565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019401920184929391613d72565b8054821015613a9b5760005260206000200190600090565b60008281526001820160205260409020546140c8578054906801000000000000000082101561346b57826140b161407c846001809601855584614026565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014614229577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613eff578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613eff578181036141f2575b505050805480156141c3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906141848282614026565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61421261420261407c9386614026565b90549060031b1c92839286614026565b90556000528360205260406000205538808061414c565b50505050600090565b919290156142ad5750815115614246575090565b3b1561424f5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156142c05750805190602001fd5b612263906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061354d56fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
}

var CCTPV2VerifierABI = CCTPV2VerifierMetaData.ABI

var CCTPV2VerifierBin = CCTPV2VerifierMetaData.Bin

func DeployCCTPV2Verifier(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, messageTransmitterProxy common.Address, usdcToken common.Address, storageLocation string, finalityConfig CCTPV2VerifierFinalityConfig, dynamicConfig CCTPV2VerifierDynamicConfig) (common.Address, *types.Transaction, *CCTPV2Verifier, error) {
	parsed, err := CCTPV2VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPV2VerifierBin), backend, tokenMessenger, messageTransmitterProxy, usdcToken, storageLocation, finalityConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPV2Verifier{address: address, abi: *parsed, CCTPV2VerifierCaller: CCTPV2VerifierCaller{contract: contract}, CCTPV2VerifierTransactor: CCTPV2VerifierTransactor{contract: contract}, CCTPV2VerifierFilterer: CCTPV2VerifierFilterer{contract: contract}}, nil
}

type CCTPV2Verifier struct {
	address common.Address
	abi     abi.ABI
	CCTPV2VerifierCaller
	CCTPV2VerifierTransactor
	CCTPV2VerifierFilterer
}

type CCTPV2VerifierCaller struct {
	contract *bind.BoundContract
}

type CCTPV2VerifierTransactor struct {
	contract *bind.BoundContract
}

type CCTPV2VerifierFilterer struct {
	contract *bind.BoundContract
}

type CCTPV2VerifierSession struct {
	Contract     *CCTPV2Verifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPV2VerifierCallerSession struct {
	Contract *CCTPV2VerifierCaller
	CallOpts bind.CallOpts
}

type CCTPV2VerifierTransactorSession struct {
	Contract     *CCTPV2VerifierTransactor
	TransactOpts bind.TransactOpts
}

type CCTPV2VerifierRaw struct {
	Contract *CCTPV2Verifier
}

type CCTPV2VerifierCallerRaw struct {
	Contract *CCTPV2VerifierCaller
}

type CCTPV2VerifierTransactorRaw struct {
	Contract *CCTPV2VerifierTransactor
}

func NewCCTPV2Verifier(address common.Address, backend bind.ContractBackend) (*CCTPV2Verifier, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPV2VerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPV2Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPV2Verifier{address: address, abi: abi, CCTPV2VerifierCaller: CCTPV2VerifierCaller{contract: contract}, CCTPV2VerifierTransactor: CCTPV2VerifierTransactor{contract: contract}, CCTPV2VerifierFilterer: CCTPV2VerifierFilterer{contract: contract}}, nil
}

func NewCCTPV2VerifierCaller(address common.Address, caller bind.ContractCaller) (*CCTPV2VerifierCaller, error) {
	contract, err := bindCCTPV2Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierCaller{contract: contract}, nil
}

func NewCCTPV2VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPV2VerifierTransactor, error) {
	contract, err := bindCCTPV2Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierTransactor{contract: contract}, nil
}

func NewCCTPV2VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPV2VerifierFilterer, error) {
	contract, err := bindCCTPV2Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierFilterer{contract: contract}, nil
}

func bindCCTPV2Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPV2VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPV2Verifier *CCTPV2VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPV2Verifier.Contract.CCTPV2VerifierCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPV2Verifier *CCTPV2VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.CCTPV2VerifierTransactor.contract.Transfer(opts)
}

func (_CCTPV2Verifier *CCTPV2VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.CCTPV2VerifierTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPV2Verifier.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.contract.Transfer(opts)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) CCTPV2MESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "CCTP_V2_MESSAGE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) CCTPV2MESSAGELENGTH() (*big.Int, error) {
	return _CCTPV2Verifier.Contract.CCTPV2MESSAGELENGTH(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) CCTPV2MESSAGELENGTH() (*big.Int, error) {
	return _CCTPV2Verifier.Contract.CCTPV2MESSAGELENGTH(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) SIGNATURELENGTH() (*big.Int, error) {
	return _CCTPV2Verifier.Contract.SIGNATURELENGTH(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _CCTPV2Verifier.Contract.SIGNATURELENGTH(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCTPV2Verifier.Contract.GetDestChainConfig(&_CCTPV2Verifier.CallOpts, destChainSelector)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCTPV2Verifier.Contract.GetDestChainConfig(&_CCTPV2Verifier.CallOpts, destChainSelector)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (CCTPV2VerifierDynamicConfig, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CCTPV2VerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPV2VerifierDynamicConfig)).(*CCTPV2VerifierDynamicConfig)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) GetDynamicConfig() (CCTPV2VerifierDynamicConfig, error) {
	return _CCTPV2Verifier.Contract.GetDynamicConfig(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) GetDynamicConfig() (CCTPV2VerifierDynamicConfig, error) {
	return _CCTPV2Verifier.Contract.GetDynamicConfig(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CCTPV2Verifier.Contract.GetFee(&_CCTPV2Verifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CCTPV2Verifier.Contract.GetFee(&_CCTPV2Verifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) GetFinalityConfig(opts *bind.CallOpts) (CCTPV2VerifierFinalityConfig, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "getFinalityConfig")

	if err != nil {
		return *new(CCTPV2VerifierFinalityConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPV2VerifierFinalityConfig)).(*CCTPV2VerifierFinalityConfig)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) GetFinalityConfig() (CCTPV2VerifierFinalityConfig, error) {
	return _CCTPV2Verifier.Contract.GetFinalityConfig(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) GetFinalityConfig() (CCTPV2VerifierFinalityConfig, error) {
	return _CCTPV2Verifier.Contract.GetFinalityConfig(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) GetStorageLocation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "getStorageLocation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) GetStorageLocation() (string, error) {
	return _CCTPV2Verifier.Contract.GetStorageLocation(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) GetStorageLocation() (string, error) {
	return _CCTPV2Verifier.Contract.GetStorageLocation(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) Owner() (common.Address, error) {
	return _CCTPV2Verifier.Contract.Owner(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) Owner() (common.Address, error) {
	return _CCTPV2Verifier.Contract.Owner(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPV2Verifier.Contract.SupportsInterface(&_CCTPV2Verifier.CallOpts, interfaceId)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPV2Verifier.Contract.SupportsInterface(&_CCTPV2Verifier.CallOpts, interfaceId)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) TypeAndVersion() (string, error) {
	return _CCTPV2Verifier.Contract.TypeAndVersion(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) TypeAndVersion() (string, error) {
	return _CCTPV2Verifier.Contract.TypeAndVersion(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CCTPV2Verifier.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CCTPV2Verifier *CCTPV2VerifierSession) VersionTag() ([4]byte, error) {
	return _CCTPV2Verifier.Contract.VersionTag(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierCallerSession) VersionTag() ([4]byte, error) {
	return _CCTPV2Verifier.Contract.VersionTag(&_CCTPV2Verifier.CallOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.AcceptOwnership(&_CCTPV2Verifier.TransactOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.AcceptOwnership(&_CCTPV2Verifier.TransactOpts)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ApplyAllowlistUpdates(&_CCTPV2Verifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ApplyAllowlistUpdates(&_CCTPV2Verifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ApplyDestChainConfigUpdates(&_CCTPV2Verifier.TransactOpts, destChainConfigArgs)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ApplyDestChainConfigUpdates(&_CCTPV2Verifier.TransactOpts, destChainConfigArgs)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "forwardToVerifier", message, messageId, arg2, arg3, arg4)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ForwardToVerifier(&_CCTPV2Verifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.ForwardToVerifier(&_CCTPV2Verifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCTPV2VerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) SetDynamicConfig(dynamicConfig CCTPV2VerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.SetDynamicConfig(&_CCTPV2Verifier.TransactOpts, dynamicConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) SetDynamicConfig(dynamicConfig CCTPV2VerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.SetDynamicConfig(&_CCTPV2Verifier.TransactOpts, dynamicConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) SetFinalityConfig(opts *bind.TransactOpts, finalityConfig CCTPV2VerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "setFinalityConfig", finalityConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) SetFinalityConfig(finalityConfig CCTPV2VerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.SetFinalityConfig(&_CCTPV2Verifier.TransactOpts, finalityConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) SetFinalityConfig(finalityConfig CCTPV2VerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.SetFinalityConfig(&_CCTPV2Verifier.TransactOpts, finalityConfig)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.TransferOwnership(&_CCTPV2Verifier.TransactOpts, to)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.TransferOwnership(&_CCTPV2Verifier.TransactOpts, to)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "updateStorageLocation", newLocation)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.UpdateStorageLocation(&_CCTPV2Verifier.TransactOpts, newLocation)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.UpdateStorageLocation(&_CCTPV2Verifier.TransactOpts, newLocation)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) VerifyMessage(opts *bind.TransactOpts, arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "verifyMessage", arg0, messageHash, ccvData)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) VerifyMessage(arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.VerifyMessage(&_CCTPV2Verifier.TransactOpts, arg0, messageHash, ccvData)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) VerifyMessage(arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.VerifyMessage(&_CCTPV2Verifier.TransactOpts, arg0, messageHash, ccvData)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CCTPV2Verifier *CCTPV2VerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.WithdrawFeeTokens(&_CCTPV2Verifier.TransactOpts, feeTokens)
}

func (_CCTPV2Verifier *CCTPV2VerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPV2Verifier.Contract.WithdrawFeeTokens(&_CCTPV2Verifier.TransactOpts, feeTokens)
}

type CCTPV2VerifierAllowListSendersAddedIterator struct {
	Event *CCTPV2VerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierAllowListSendersAdded)
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
		it.Event = new(CCTPV2VerifierAllowListSendersAdded)
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

func (it *CCTPV2VerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierAllowListSendersAddedIterator{contract: _CCTPV2Verifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierAllowListSendersAdded)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*CCTPV2VerifierAllowListSendersAdded, error) {
	event := new(CCTPV2VerifierAllowListSendersAdded)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierAllowListSendersRemovedIterator struct {
	Event *CCTPV2VerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierAllowListSendersRemoved)
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
		it.Event = new(CCTPV2VerifierAllowListSendersRemoved)
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

func (it *CCTPV2VerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierAllowListSendersRemovedIterator{contract: _CCTPV2Verifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierAllowListSendersRemoved)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*CCTPV2VerifierAllowListSendersRemoved, error) {
	event := new(CCTPV2VerifierAllowListSendersRemoved)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierDestChainConfigSetIterator struct {
	Event *CCTPV2VerifierDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierDestChainConfigSet)
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
		it.Event = new(CCTPV2VerifierDestChainConfigSet)
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

func (it *CCTPV2VerifierDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierDestChainConfigSetIterator{contract: _CCTPV2Verifier.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierDestChainConfigSet)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseDestChainConfigSet(log types.Log) (*CCTPV2VerifierDestChainConfigSet, error) {
	event := new(CCTPV2VerifierDestChainConfigSet)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierDynamicConfigSetIterator struct {
	Event *CCTPV2VerifierDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierDynamicConfigSet)
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
		it.Event = new(CCTPV2VerifierDynamicConfigSet)
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

func (it *CCTPV2VerifierDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierDynamicConfigSet struct {
	DynamicConfig CCTPV2VerifierDynamicConfig
	Raw           types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierDynamicConfigSetIterator, error) {

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierDynamicConfigSetIterator{contract: _CCTPV2Verifier.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierDynamicConfigSet)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseDynamicConfigSet(log types.Log) (*CCTPV2VerifierDynamicConfigSet, error) {
	event := new(CCTPV2VerifierDynamicConfigSet)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierFeeTokenWithdrawnIterator struct {
	Event *CCTPV2VerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierFeeTokenWithdrawn)
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
		it.Event = new(CCTPV2VerifierFeeTokenWithdrawn)
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

func (it *CCTPV2VerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPV2VerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierFeeTokenWithdrawnIterator{contract: _CCTPV2Verifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierFeeTokenWithdrawn)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCTPV2VerifierFeeTokenWithdrawn, error) {
	event := new(CCTPV2VerifierFeeTokenWithdrawn)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierFinalityConfigSetIterator struct {
	Event *CCTPV2VerifierFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierFinalityConfigSet)
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
		it.Event = new(CCTPV2VerifierFinalityConfigSet)
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

func (it *CCTPV2VerifierFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierFinalityConfigSet struct {
	FinalityConfig CCTPV2VerifierFinalityConfig
	Raw            types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierFinalityConfigSetIterator, error) {

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierFinalityConfigSetIterator{contract: _CCTPV2Verifier.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierFinalityConfigSet)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseFinalityConfigSet(log types.Log) (*CCTPV2VerifierFinalityConfigSet, error) {
	event := new(CCTPV2VerifierFinalityConfigSet)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierOwnershipTransferRequestedIterator struct {
	Event *CCTPV2VerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierOwnershipTransferRequested)
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
		it.Event = new(CCTPV2VerifierOwnershipTransferRequested)
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

func (it *CCTPV2VerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPV2VerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierOwnershipTransferRequestedIterator{contract: _CCTPV2Verifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierOwnershipTransferRequested)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPV2VerifierOwnershipTransferRequested, error) {
	event := new(CCTPV2VerifierOwnershipTransferRequested)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierOwnershipTransferredIterator struct {
	Event *CCTPV2VerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierOwnershipTransferred)
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
		it.Event = new(CCTPV2VerifierOwnershipTransferred)
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

func (it *CCTPV2VerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPV2VerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierOwnershipTransferredIterator{contract: _CCTPV2Verifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierOwnershipTransferred)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPV2VerifierOwnershipTransferred, error) {
	event := new(CCTPV2VerifierOwnershipTransferred)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierStaticConfigSetIterator struct {
	Event *CCTPV2VerifierStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierStaticConfigSet)
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
		it.Event = new(CCTPV2VerifierStaticConfigSet)
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

func (it *CCTPV2VerifierStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierStaticConfigSet struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	UsdcToken               common.Address
	LocalDomainIdentifier   uint32
	Raw                     types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierStaticConfigSetIterator, error) {

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierStaticConfigSetIterator{contract: _CCTPV2Verifier.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierStaticConfigSet)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseStaticConfigSet(log types.Log) (*CCTPV2VerifierStaticConfigSet, error) {
	event := new(CCTPV2VerifierStaticConfigSet)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPV2VerifierStorageLocationUpdatedIterator struct {
	Event *CCTPV2VerifierStorageLocationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPV2VerifierStorageLocationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPV2VerifierStorageLocationUpdated)
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
		it.Event = new(CCTPV2VerifierStorageLocationUpdated)
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

func (it *CCTPV2VerifierStorageLocationUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPV2VerifierStorageLocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPV2VerifierStorageLocationUpdated struct {
	OldLocation string
	NewLocation string
	Raw         types.Log
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CCTPV2VerifierStorageLocationUpdatedIterator, error) {

	logs, sub, err := _CCTPV2Verifier.contract.FilterLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPV2VerifierStorageLocationUpdatedIterator{contract: _CCTPV2Verifier.contract, event: "StorageLocationUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierStorageLocationUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPV2Verifier.contract.WatchLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPV2VerifierStorageLocationUpdated)
				if err := _CCTPV2Verifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
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

func (_CCTPV2Verifier *CCTPV2VerifierFilterer) ParseStorageLocationUpdated(log types.Log) (*CCTPV2VerifierStorageLocationUpdated, error) {
	event := new(CCTPV2VerifierStorageLocationUpdated)
	if err := _CCTPV2Verifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}
type GetFee struct {
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}

func (CCTPV2VerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CCTPV2VerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CCTPV2VerifierDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CCTPV2VerifierDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b")
}

func (CCTPV2VerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPV2VerifierFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0xbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc026754")
}

func (CCTPV2VerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPV2VerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCTPV2VerifierStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c")
}

func (CCTPV2VerifierStorageLocationUpdated) Topic() common.Hash {
	return common.HexToHash("0xbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8")
}

func (_CCTPV2Verifier *CCTPV2Verifier) Address() common.Address {
	return _CCTPV2Verifier.address
}

type CCTPV2VerifierInterface interface {
	CCTPV2MESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error)

	SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CCTPV2VerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetFinalityConfig(opts *bind.CallOpts) (CCTPV2VerifierFinalityConfig, error)

	GetStorageLocation(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCTPV2VerifierDynamicConfig) (*types.Transaction, error)

	SetFinalityConfig(opts *bind.TransactOpts, finalityConfig CCTPV2VerifierFinalityConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CCTPV2VerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CCTPV2VerifierAllowListSendersRemoved, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPV2VerifierDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CCTPV2VerifierDestChainConfigSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPV2VerifierDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPV2VerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPV2VerifierFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CCTPV2VerifierFinalityConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPV2VerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPV2VerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPV2VerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPV2VerifierOwnershipTransferred, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*CCTPV2VerifierStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*CCTPV2VerifierStaticConfigSet, error)

	FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CCTPV2VerifierStorageLocationUpdatedIterator, error)

	WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPV2VerifierStorageLocationUpdated) (event.Subscription, error)

	ParseStorageLocationUpdated(log types.Log) (*CCTPV2VerifierStorageLocationUpdated, error)

	Address() common.Address
}
