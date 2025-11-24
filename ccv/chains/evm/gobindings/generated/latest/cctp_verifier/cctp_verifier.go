// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_verifier

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

type CCTPVerifierDomain struct {
	AllowedCallerOnDest   [32]byte
	AllowedCallerOnSource [32]byte
	MintRecipientOnDest   [32]byte
	DomainIdentifier      uint32
	Enabled               bool
}

type CCTPVerifierDynamicConfig struct {
	FeeAggregator  common.Address
	AllowlistAdmin common.Address
}

type CCTPVerifierFinalityConfig struct {
	DefaultCCTPFinalityThreshold uint32
	DefaultCCTPFinalityBps       uint16
	CustomCCIPFinalities         []uint16
	CustomCCTPFinalityThresholds []uint32
	CustomCCTPFinalityBps        []uint16
}

type CCTPVerifierSetDomainArgs struct {
	AllowedCallerOnDest   [32]byte
	AllowedCallerOnSource [32]byte
	MintRecipientOnDest   [32]byte
	ChainSelector         uint64
	DomainIdentifier      uint32
	Enabled               bool
}

type CCTPVerifierStaticConfig struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	UsdcToken               common.Address
	LocalDomainIdentifier   uint32
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

var CCTPVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.StaticConfig\",\"components\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalityConfig\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomFinalitiesMustBeStrictlyIncreasing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomFinalityArraysMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MissingCustomFinalities\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101208060405234610c3a57615c64803803809161001d8285610fb5565b83398101908082039060e08212610c3a578051906001600160a01b03821690818303610c3a576020810151926001600160a01b038416808503610c3a5760408301516001600160a01b0381169390848103610c3a5760608201516001600160401b038111610c3a5782019289601f85011215610c3a5783519361009f85610fd8565b946100ad6040519687610fb5565b8086528b60208284010111610c3a576100cc9160208088019101610ff3565b60808301516001600160401b038111610c3a5783019760a0898c0312610c3a576040519560a087016001600160401b038111888210176109b1576040526101128a611016565b875261012060208b01611027565b6020880190815260408b01519099906001600160401b038111610c3a578d610149918d0161104d565b6040890190815260608c0151909d909b6001600160401b038d11610c3a5781601f9d82019d8e011215610c3a578c519c6101828e611036565b9d6040519e8f906101939082610fb5565b8181526020019060051b820160200191848311610c3a57602001905b828210610f9d5750505060608a019c8d526080810151916001600160401b038311610c3a576040926101e1920161104d565b60808a019081529c609f190112610c3a5760408051969087016001600160401b038111888210176109b15761022a9160c09160405261022260a082016110b2565b8952016110b2565b96602087019788523315610f8c57600180546001600160a01b0319163317905560405160035490919060008361025f8361113f565b80825291600184168015610f6e57600114610f0e575b61028192500384610fb5565b8151906001600160401b0382116109b15761029b9061113f565b601f8111610eb4575b506020601f8211600114610e37576102fd9282600080516020615c44833981519152959361030b93600091610e2c575b508160011b916000199060031b1c1916176003555b604051938493604085526040850190611179565b908382036020850152611179565b0390a180158015610e24575b8015610e1c575b610bad57604051639cdbb18160e01b8152602081600481855afa8015610c8f57600090610ddf575b63ffffffff91501660018103610dc65750602060049160405192838092632c12192160e01b82525afa908115610c8f57600091610d8c575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa8015610c8f57600090610d4f575b63ffffffff91501660018103610d3657506020600491604051928380926367e0ed8360e11b82525afa908115610c8f57600091610ced575b506001600160a01b0316838103610cd5575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa908115610c8f57600091610c9b575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529490602081604481855afa908115610c8f57600091610c5d575b506000198101809111610c47576105219160405191602083019763095ea7b360e01b895260248401526044830152604482526104b1606483610fb5565b6000806040988951946104c48b87610fb5565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15610c3f573d9161050583610fd8565b926105128a519485610fb5565b83523d6000602085013e61119e565b805180610bbe575b505060e05160c05160a0516101005187516001600160a01b039485168152928416602084015292168187015263ffffffff90911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c90608090a180516001600160a01b031615610bad5751600580546001600160a01b03199081166001600160a01b03938416908117909255835160068054909216908416179055845190815291511660208201527f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b908390a185515115610b9c57855151845151811490811591610b8f575b50610b7e57600093845b8751805187101561066c5761ffff6106378882936110c6565b51169116101561065b57600161ffff610651878a516110c6565b511695019461061e565b633fd3668360e21b60005260046000fd5b505085878563ffffffff8551166007549065ffff00000000835160201b169165ffffffffffff19161717600755815180519060018060401b0382116109b1576801000000000000000082116109b15760209060085483600855808410610b48575b500190600860005260206000208160041c9160005b838110610b075750600f198116900380610ab8575b505085518051925090506001600160401b0382116109b1576801000000000000000082116109b15760209060095483600955808410610a5e575b500190600960005260206000208160031c9160005b838110610a1b575060071981169003806109c7575b505084518051925090506001600160401b0382116109b1576801000000000000000082116109b157602090600a5483600a55808410610957575b500190600a60005260206000208160041c9160005b8381106109165750600f1981169003806108c3575b505050509061ffff6107f69263ffffffff88519760208952511660208801525116868601525160a0606086015260c0850190611107565b915191601f198482030160808501526020808451928381520193019060005b8181106108a757867f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd48780610857898951601f198483030160a0850152611107565b0390a151614a08908161123c823960805181505060a051818181612c100152613989015260c051818181611a7d015261395c015260e051818181612d0d015261390b015261010051816139b00152f35b825163ffffffff16855260209485019490920191600101610815565b9260009360005b8181106108e357505050015561ffff6107f688806107bf565b909194602061090c60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b96019291016108ca565b6000805b6010811061092f5750838201556001016107aa565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161091a565b61098890600a60005283600020600f80870160041c820192601e8860011b168061098e575b500160041c01906110f0565b89610795565b6109ab906000198601908154906000199060200360031b1c169055565b8e61097c565b634e487b7160e01b600052604160045260246000fd5b9260009360005b8181106109e35750505001558680808061075b565b9091946020610a1160019263ffffffff895116908560021b63ffffffff809160031b9316831b921b19161790565b96019291016109ce565b6000805b60088110610a34575083820155600101610746565b865190969160019160209163ffffffff60058b901b81811b199092169216901b1792019601610a1f565b610a8f90600960005283600020600780870160031c820192601c8860021b1680610a95575b500160031c01906110f0565b89610731565b610ab2906000198601908154906000199060200360031b1c169055565b8e610a83565b9260009360005b818110610ad4575050500155868080806106f7565b9091946020610afd60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101610abf565b6000805b60108110610b205750838201556001016106e2565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601610b0b565b610b7890600860005283600020600f80870160041c820192601e8860011b168061098e57500160041c01906110f0565b896106cd565b630cacb92960e41b60005260046000fd5b9050855151141538610614565b636c33513960e11b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b8160209181010312610c3a5760200151801590811503610c3a57610be3573880610529565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b60609161119e565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610c87575b81610c7860209383610fb5565b81010312610c3a575138610474565b3d9150610c6b565b6040513d6000823e3d90fd5b90506020813d602011610ccd575b81610cb660209383610fb5565b81010312610c3a57610cc790611016565b3861042a565b3d9150610ca9565b836383395ca960e01b60005260045260245260446000fd5b6020813d602011610d2e575b81610d0660209383610fb5565b81010312610d2a5751906001600160a01b0382168203610d275750386103e9565b80fd5b5080fd5b3d9150610cf9565b6331b6aa1b60e11b600052600160045260245260446000fd5b506020813d602011610d84575b81610d6960209383610fb5565b81010312610c3a57610d7f63ffffffff91611016565b6103b1565b3d9150610d5c565b90506020813d602011610dbe575b81610da760209383610fb5565b81010312610c3a57610db8906110b2565b3861037e565b3d9150610d9a565b633785f8f160e01b600052600160045260245260446000fd5b506020813d602011610e14575b81610df960209383610fb5565b81010312610c3a57610e0f63ffffffff91611016565b610346565b3d9150610dec565b50881561031e565b508315610317565b9050820151386102d4565b601f198216906003600052806000209160005b818110610e9c57508361030b936102fd9693600080516020615c44833981519152989660019410610e83575b5050811b016003556102e9565b84015160001960f88460031b161c191690553880610e76565b91926020600181928689015181550194019201610e4a565b6003600052610efe907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610f04575b601f0160051c01906110f0565b386102a4565b9091508190610ef1565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310610f5257505090602061028192820101610275565b6020919350806001915483858a01015201910190918592610f3a565b505060206102819260ff19851682840152151560051b820101610275565b639b15e16f60e01b60005260046000fd5b60208091610faa84611016565b8152019101906101af565b601f909101601f19168101906001600160401b038211908210176109b157604052565b6001600160401b0381116109b157601f01601f191660200190565b60005b8381106110065750506000910152565b8181015183820152602001610ff6565b519063ffffffff82168203610c3a57565b519061ffff82168203610c3a57565b6001600160401b0381116109b15760051b60200190565b9080601f83011215610c3a57815161106481611036565b926110726040519485610fb5565b81845260208085019260051b820101928311610c3a57602001905b82821061109a5750505090565b602080916110a784611027565b81520191019061108d565b51906001600160a01b0382168203610c3a57565b80518210156110da5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8181106110fb575050565b600081556001016110f0565b906020808351928381520192019060005b8181106111255750505090565b825161ffff16845260209384019390920191600101611118565b90600182811c9216801561116f575b602083101461115957565b634e487b7160e01b600052602260045260246000fd5b91607f169161114e565b9060209161119281518092818552858086019101610ff3565b601f01601f1916010190565b9192901561120057508151156111b2575090565b3b156111bb5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156112135750805190602001fd5b60405162461bcd60e51b815260206004820152908190611237906024830190611179565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146139ef5750806306285c69146138a4578063181f5a771461382557806338a9eb2a1461312c5780633bbbed4b146129bb5780635cb80c5d146126da5780636def4ce7146126055780637437ff9f1461254657806379ba50971461246157806380485e25146121d5578063869b7f62146120535780638da5cb5b14612001578063b2bd751c14611cbb578063bff0ec1d146115d5578063c757d73f14610d72578063c9b146b314610988578063ceac5cee146106a2578063d52e545a14610392578063dfadfa35146102ad578063f2fde38b146101c0578063fe163eed146101675763fec888af1461011357600080fd5b3461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645761016061014c61429c565b604051918291602083526020830190613bc7565b0390f35b80fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645773ffffffffffffffffffffffffffffffffffffffff61020d613d64565b610215614373565b1633811461028557807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457604060a09167ffffffffffffffff6102f3613e07565b826080855161030181613af8565b828152826020820152828782015282606082015201521681526004602052206040519061032d82613af8565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e576103e2903690600401613f07565b6103ed929192614373565b81925b818410156105e65760c0840281019360c0853603126105e2576040519461041686613b14565b803590818752602081013590602088019782895260408101906040830135825261044260608401613e1e565b93606082019480865261046a60a061045c60808801613f38565b9660808601978852016141e9565b9660a08401978852159182156105d9575b5081156105c6575b50610560579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d51945192519351169751151596604051946104bc86613af8565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b169116179055019291906103f0565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff9150161538610483565b1591503861047b565b8380fd5b6040519180602084016020855252604083019190845b81811061062d57857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff61066360608901613e1e565b16606082015263ffffffff61067a60808901613f38565b16608082015261068c60a088016141e9565b151560a08201520194019291016105fc565b5080fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e573660238201121561069e57610703903690602481600401359101613e7d565b61070b614373565b61071361429c565b90805167ffffffffffffffff811161095b57610730600354614249565b601f81116108c4575b506020601f82116001146107df576107c092827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd895936107ce9388916107d4575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b604051938493604085526040850190613bc7565b908382036020850152613bc7565b0390a180f35b90508201513861077a565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b8181106108ac5750836107ce936107c096937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610875575b5050811b016003556107ac565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610868565b9192602060018192868901518155019401920161082a565b61092d9060038652601f830160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208410610933575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01906143d2565b38610739565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b91506108ff565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e576109d8903690600401613dd6565b73ffffffffffffffffffffffffffffffffffffffff600154163303610d2a575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b84811015610d26578060051b82013583811215610d2257820191608083360312610d225760405194610a5486613aad565b610a5d84613e1e565b8652610a6b602085016141e9565b9660208701978852604085013567ffffffffffffffff8111610d1e57610a9490369087016143e9565b9460408801958652606081013567ffffffffffffffff81116105e257610abc913691016143e9565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610b36818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610bf6575b5095976001019550815b85518051821015610b875790610b8073ffffffffffffffffffffffffffffffffffffffff610b78836001956143be565b5116896147cc565b5001610b48565b50509590969450600192919351908151610ba7575b505001939293610a23565b610bec67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613e33565b0390a23880610b9c565b98939592969190949798600014610ce757600184019591875b86518051821015610c8c57610c398273ffffffffffffffffffffffffffffffffffffffff926143be565b51168015610c555790610c4e6001928a61473b565b5001610c0f565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281610cdd67ffffffffffffffff8a51169251604051918291602083526020830190613e33565b0390a23880610b3e565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b8480fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff600654163303156109f8576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261069e5760405191610ded83613af8565b610df982600401613f38565b8352610e0760248301613ef8565b9260208101938452604483013567ffffffffffffffff8111610d1e57610e339060043691860101614023565b9360408201948552606484013567ffffffffffffffff81116105e257840193366023860112156105e257600485013594610e6c86613ecf565b95610e7a6040519788613b4c565b808752602060048189019260051b84010101913683116115d157602401905b8282106115b95750505060608301948552608481013567ffffffffffffffff8111610d2257610ecd91369101600401614023565b9160808101928352610edd614373565b8551511561159157855151855151811490811591611584575b5061155c578392835b87518051861015610f645761ffff610f188782936143be565b511691161015610f3c57600161ffff610f32868a516143be565b5116940193610eff565b6004857fff4d9a0c000000000000000000000000000000000000000000000000000000008152fd5b50508587869463ffffffff8551167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006007541617600755517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff65ffff000000006007549260201b169116176007555180519067ffffffffffffffff821161152f5768010000000000000000821161152f576020906008548360085580841061149d575b50019060088652602086208160041c91875b83811061145d57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380611410575b505050505180519067ffffffffffffffff82116113e3576801000000000000000082116113e35760209060095483600955808410611350575b50019060098552602085208160031c91865b83811061130e57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff881169003806112bc575b505050505180519067ffffffffffffffff821161095b5768010000000000000000821161095b57602090600a5483600a55808410611229575b500190600a8452602084208160041c91855b8381106111e957507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0811690038061117a575b857f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd46107ce8760405191829182613c5e565b928593865b8181106111b65750505001556107ce7f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd48480611148565b90919460206111df60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b960192910161117f565b86875b60108110611201575083820155600101611115565b865190969160019160209161ffff60048b901b81811b199092169216901b17920196016111ec565b61125890600a8752838720600f80870160041c820192601e8860011b168061125e575b500160041c01906143d2565b85611103565b6112b6907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8a61124c565b928693875b8181106112d6575050500155838080806110ca565b909194602061130460019263ffffffff895116908560021b63ffffffff809160031b9316831b921b19161790565b96019291016112c1565b87885b60088110611326575083820155600101611097565b865190969160019160209163ffffffff60058b901b81811b199092169216901b1792019601611311565b61137f9060098852838820600780870160031c820192601c8860021b1680611385575b500160031c01906143d2565b86611085565b6113dd907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8b611373565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b928793885b81811061142a5750505001558480808061104c565b909194602061145360019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101611415565b88895b60108110611475575083820155600101611019565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601611460565b6114cb9060088952838920600f80870160041c820192601e8860011b16806114d157500160041c01906143d2565b87611007565b611529907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8c61124c565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6004847fcacb9290000000000000000000000000000000000000000000000000000000008152fd5b9050835151141538610ef6565b6004847fd866a272000000000000000000000000000000000000000000000000000000008152fd5b602080916115c684613f38565b815201910190610e99565b8680fd5b50346101645760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261069e57604051906101c0820182811067ffffffffffffffff82111761095b5760405261166f81600401613e1e565b825261167d60248201613e1e565b602083015261168e60448201613e1e565b604083015261169f60648201613f38565b60608301526116b060848201613f38565b60808301526116c160a48201613ef8565b60a083015260c481013560c083015260e481013567ffffffffffffffff81116105e2576116f49060043691840101613eb4565b60e083015261010481013567ffffffffffffffff81116105e25761171e9060043691840101613eb4565b61010083015261012481013567ffffffffffffffff81116105e2576117499060043691840101613eb4565b61012083015261014481013567ffffffffffffffff81116105e2576117749060043691840101613eb4565b61014083015261016481013567ffffffffffffffff81116105e25761179f9060043691840101613eb4565b61016083015261018481013567ffffffffffffffff81116105e2578101366023820112156105e2576004810135906117d682613ecf565b916117e46040519384613b4c565b80835260051b8101602401602083013682116115d15760248301905b828210611c8557505050506101808301526101a48101359067ffffffffffffffff82116105e25760046118369236920101613eb4565b6101a082015260243560443567ffffffffffffffff81116105e25761185f903690600401613da8565b90926101e18210611c5d5781600411610d22577fffffffff000000000000000000000000000000000000000000000000000000008435167f8e1d1a9d000000000000000000000000000000000000000000000000000000008103611c0e57506101809261017c600084861161016457507fffffffff00000000000000000000000000000000000000000000000000000000908681013580831691870360048110611bf9575b5050167f8e1d1a9d000000000000000000000000000000000000000000000000000000008103611baa57506000906101a094858111610d1e57848611610d1e57611953908087039088016141f6565b90818103611b7c57505067ffffffffffffffff82511686526004602052604086209160ff600384015460201c1615611b46575061011c9050600083821161016457506001906119c7907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc87016141f6565b910154808203611b185750508082116105e25761019c92611a636004927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe606020958401910160405196879586957f57ecfd280000000000000000000000000000000000000000000000000000000087526040828801528260448801520160648601378761020085015261020484019161020060248601526141aa565b03818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611b0d578291611ade575b5015611ab65780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b611b00915060203d602011611b06575b611af88183613b4c565b810190614231565b38611aad565b503d611aee565b6040513d84823e3d90fd5b7f7d8b101a000000000000000000000000000000000000000000000000000000008652600452602452604484fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000875267ffffffffffffffff16600452602486fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008852600452602452604486fd5b7fadaf77390000000000000000000000000000000000000000000000000000000087527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604486fd5b839250829060040360031b1b16163880611904565b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b6004857fbba6473c000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111611cb757602091611cac839283600436928a010101613f49565b815201910190611800565b8880fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e57611d0b903690600401613f07565b611d13614373565b611d1c81613ecf565b91611d2a6040519384613b4c565b81835260c0602084019202810190368211610d2257915b818310611f64578480855b8051831015611f6057611d5f83826143be565b519267ffffffffffffffff6020611d7683856143be565b51015116938415611f34578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff81511615611f085773ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a2019190611d4c565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c083360312610d225760405190611f7b82613b14565b83359073ffffffffffffffffffffffffffffffffffffffff821682036115d1578260209260c09452611fae838701613e1e565b83820152611fbe604087016141e9565b6040820152611fcf60608701613ef8565b6060820152611fe060808701613f38565b6080820152611ff160a08701613f38565b60a0820152815201920191611d41565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101645760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760405161208f81613b30565b612097613d64565b815260243573ffffffffffffffffffffffffffffffffffffffff81168103610d1e57602082019081526120c8614373565b73ffffffffffffffffffffffffffffffffffffffff825116156121ad578173ffffffffffffffffffffffffffffffffffffffff6107ce92817f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346101645760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645761220d613e07565b60243567ffffffffffffffff8111610d1e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610d1e576040519061225882613af8565b806004013567ffffffffffffffff8111610d225761227c9060043691840101613eb4565b8252602481013567ffffffffffffffff8111610d22576122a29060043691840101613eb4565b6020830152604481013567ffffffffffffffff8111610d2257810136602382011215610d225760048101356122d681613ecf565b916122e46040519384613b4c565b818352602060048185019360061b830101019036821161245d57602401915b81831061242557505050604083015261231e60648201613d87565b6060830152608481013567ffffffffffffffff8111610d225760809160046123499236920101613eb4565b9101526044359067ffffffffffffffff8211610d1e5761237667ffffffffffffffff923690600401613eb4565b5061237f613ee7565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff604082205416156123fa57816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b60408336031261245d576020604091825161243f81613b30565b61244886613d87565b81528286013583820152815201920191612303565b8780fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457805473ffffffffffffffffffffffffffffffffffffffff8116330361251e577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457602060405161258381613b30565b828152015261016060405161259781613b30565b73ffffffffffffffffffffffffffffffffffffffff60055416815273ffffffffffffffffffffffffffffffffffffffff60065416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645767ffffffffffffffff612646613e07565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b8181106126c45773ffffffffffffffffffffffffffffffffffffffff856101608861269c81890382613b4c565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613e33565b825484526020909301926001928301920161266f565b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161069e5761272a903690600401613dd6565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b838110156129b7578060051b8201359073ffffffffffffffffffffffffffffffffffffffff82168092036129b357604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa9283156129a8578793612975575b50826127cb575b506001915001612747565b846040519361288660208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152612816606482613b4c565b8a806040988951936128288b86613b4c565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d1561296d573d9061286982613b8d565b916128768a519384613b4c565b82523d8d602084013e5b8661492f565b8051806128c2575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3386127c0565b6128d9929495969350602080918301019101614231565b156128ea579291908590388061288e565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090612880565b9092506020813d82116129a0575b8161299060209383613b4c565b810103126115d1575191386127b9565b3d9150612983565b6040513d89823e3d90fd5b8580fd5b8480f35b50346101645760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610164576004359067ffffffffffffffff82116101645781600401906101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc843603011261016457612a38613d3c565b5060843567ffffffffffffffff811161069e57612a59903690600401613da8565b5050612a69610124840183614088565b90357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811691601481106130f7575b505060601c92602481019367ffffffffffffffff612ab5866140d9565b168084526002602052604084209081549073ffffffffffffffffffffffffffffffffffffffff82169081156130cc576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156130c157869061305e575b73ffffffffffffffffffffffffffffffffffffffff91501633036130325760f01c60ff16612fea575b505067ffffffffffffffff612b64856140d9565b1682526004602052604082209060038201549460ff8660201c1615612fac575061018481016001612b9582876140ee565b905003612f7657612bb9612bb4612bac83886140ee565b369291614142565b613f49565b906040820151602081519101517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612f41575b505060601c9573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001696878103612f1657506020612c5b612c51612c4b85856140ee565b90614142565b6080810190614088565b905003612ec557505060028301548015612ea55760a490915b519201359261ffff8416809403610d2257549260405192612c9484613b14565b835260208301918252612cc9604084019160243583526060850195865280608086015263ffffffff60a086019916895261444e565b61ffff8551911690818102918183041490151715612e7857612710900463ffffffff8111612e4d5763ffffffff73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016955199511693519551925192604051937f8e1d1a9d000000000000000000000000000000000000000000000000000000006020860152602485015260248452612d77604485613b4c565b853b1561245d579263ffffffff88999a9381612df4948b976040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c015260248b015260448a0152606489015260848801521660a48601521660c484015261010060e4840152610104830190613bc7565b03925af1918215612e40578161016093612e30575b505060405190612e1a602083613b4c565b8152604051918291602083526020830190613bc7565b612e3991613b4c565b3881612e09565b50604051903d90823e3d90fd5b7fb6f15d0f000000000000000000000000000000000000000000000000000000008752600452602486fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b506080810151602081805181010312610d2257602060a491015191612c74565b612c4b612ed592612c51926140ee565b612f126040519283927fa3c8cf090000000000000000000000000000000000000000000000000000000084526020600485015260248401916141aa565b0390fd5b7f961c9a4f000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612bf3565b83612f83602492876140ee565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8367ffffffffffffffff612fc16024936140d9565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b600082815260029091016020526040902054156130075780612b50565b7fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d6020116130b9575b8161307860209383613b4c565b810103126129b3575173ffffffffffffffffffffffffffffffffffffffff811681036129b35773ffffffffffffffffffffffffffffffffffffffff90612b27565b3d915061306b565b6040513d88823e3d90fd5b7f8a4e93c9000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612a98565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610164576060608060405161316b81613af8565b8381528360208201528260408201528280820152015260405161318d81613af8565b61ffff60075463ffffffff8116835260201c166020820152604051600854808252816020810160088652602086209286905b80600f83011061373d576132559454918181106134ff575b8181106134e7575b8181106134d0575b8181106134b8575b8181106134a0575b818110613488575b818110613470575b818110613458575b818110613440575b818110613428575b818110613410575b8181106133f8575b8181106133e0575b8181106133c8575b8181106133b0575b106133a2575b500382613b4c565b604082015260405180816020600954928381520160098652602086209286905b8060078301106136b6576132ca94549181811061369f575b818110613686575b81811061366c575b818110613652575b818110613638575b81811061361e575b818110613604575b106135f657500382613b4c565b6060820152604051600a80548083529084526020820193907fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a85b81600f84011061351457946133919284926101609754918181106134ff578181106134e7578181106134d0578181106134b8578181106134a0578181106134885781811061347057818110613458578181106134405781811061342857818110613410578181106133f8578181106133e0578181106133c8578181106133b057106133a257500382613b4c565b608082015260405191829182613c5e565b60f01c81526020013861324d565b92602060019161ffff8560e01c168152019301613247565b92602060019161ffff8560d01c16815201930161323f565b92602060019161ffff8560c01c168152019301613237565b92602060019161ffff8560b01c16815201930161322f565b92602060019161ffff8560a01c168152019301613227565b92602060019161ffff8560901c16815201930161321f565b92602060019161ffff8560801c168152019301613217565b92602060019161ffff8560701c16815201930161320f565b92602060019161ffff8560601c168152019301613207565b92602060019161ffff8560501c1681520193016131ff565b92602060019161ffff8560401c1681520193016131f7565b92602060019161ffff8560301c1681520193016131ef565b92602060019161ffff85831c1681520193016131e7565b92602060019161ffff8560101c1681520193016131df565b92602060019161ffff851681520193016131d7565b946001610200601092885461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019601920191613304565b60e01c81526020013861324d565b92602060019163ffffffff8560c01c1681520193016132bd565b92602060019163ffffffff8560a01c1681520193016132b5565b92602060019163ffffffff8560801c1681520193016132ad565b92602060019163ffffffff8560601c1681520193016132a5565b92602060019163ffffffff8560401c16815201930161329d565b92602060019163ffffffff85831c168152019301613295565b92602060019163ffffffff8516815201930161328d565b916008919350610100600191865463ffffffff8116825263ffffffff8160201c16602083015263ffffffff8160401c16604083015263ffffffff8160601c16606083015263ffffffff8160801c16608083015263ffffffff8160a01c1660a083015263ffffffff8160c01c1660c083015260e01c60e0820152019401920184929391613275565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e08201520194019201849293916131bf565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645750610160604051613866604082613b4c565b601681527f43435450566572696669657220312e372e302d646576000000000000000000006020820152604051918291602083526020830190613bc7565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760606040516138e181613aad565b8281528260208201528260408201520152608073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff60405161393c81613aad565b82815273ffffffffffffffffffffffffffffffffffffffff6020820191817f00000000000000000000000000000000000000000000000000000000000000001683528160606040830192827f00000000000000000000000000000000000000000000000000000000000000001684520193857f0000000000000000000000000000000000000000000000000000000000000000168552604051968752511660208601525116604084015251166060820152f35b90503461069e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261069e576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610d1e57602092507ffacbd7dc000000000000000000000000000000000000000000000000000000008114908115613a83575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613a7c565b6080810190811067ffffffffffffffff821117613ac957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff821117613ac957604052565b60c0810190811067ffffffffffffffff821117613ac957604052565b6040810190811067ffffffffffffffff821117613ac957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613ac957604052565b67ffffffffffffffff8111613ac957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613c115750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613bd2565b906020808351928381520192019060005b818110613c445750505090565b825161ffff16845260209384019390920191600101613c37565b9190916020815263ffffffff835116602082015261ffff6020840151166040820152613c99604084015160a0606084015260c0830190613c26565b906060840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030160808301526020808451928381520193019060005b818110613d20575050506080613d1d939401519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613c26565b90565b825163ffffffff16855260209485019490920191600101613cda565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203613d5f57565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613d5f57565b359073ffffffffffffffffffffffffffffffffffffffff82168203613d5f57565b9181601f84011215613d5f5782359167ffffffffffffffff8311613d5f5760208381860195010111613d5f57565b9181601f84011215613d5f5782359167ffffffffffffffff8311613d5f576020808501948460051b010111613d5f57565b6004359067ffffffffffffffff82168203613d5f57565b359067ffffffffffffffff82168203613d5f57565b906020808351928381520192019060005b818110613e515750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613e44565b929192613e8982613b8d565b91613e976040519384613b4c565b829481845281830111613d5f578281602093846000960137010152565b9080601f83011215613d5f57816020613d1d93359101613e7d565b67ffffffffffffffff8111613ac95760051b60200190565b6064359061ffff82168203613d5f57565b359061ffff82168203613d5f57565b9181601f84011215613d5f5782359167ffffffffffffffff8311613d5f5760208085019460c08502010111613d5f57565b359063ffffffff82168203613d5f57565b919060c083820312613d5f5760405190613f6282613b14565b819380358352602081013567ffffffffffffffff8111613d5f5782613f88918301613eb4565b6020840152604081013567ffffffffffffffff8111613d5f5782613fad918301613eb4565b6040840152606081013567ffffffffffffffff8111613d5f5782613fd2918301613eb4565b6060840152608081013567ffffffffffffffff8111613d5f5782613ff7918301613eb4565b608084015260a08101359167ffffffffffffffff8311613d5f5760a09261401e9201613eb4565b910152565b9080601f83011215613d5f57813561403a81613ecf565b926140486040519485613b4c565b81845260208085019260051b820101928311613d5f57602001905b8282106140705750505090565b6020809161407d84613ef8565b815201910190614063565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613d5f570180359067ffffffffffffffff8211613d5f57602001918136038313613d5f57565b3567ffffffffffffffff81168103613d5f5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613d5f570180359067ffffffffffffffff8211613d5f57602001918160051b36038313613d5f57565b901561417b578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215613d5f570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b35908115158203613d5f57565b359060208110614204575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b90816020910312613d5f57518015158103613d5f5790565b90600182811c92168015614292575b602083101461426357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614258565b60405190600082600354916142b083614249565b808352926001811690811561433657506001146142d6575b6142d492500383613b4c565b565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b81831061431a5750509060206142d4928201016142c8565b6020919350806001915483858901015201910190918492614302565b602092506142d49491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016142c8565b73ffffffffffffffffffffffffffffffffffffffff60015416330361439457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805182101561417b5760209160051b010190565b8181106143dd575050565b600081556001016143d2565b9080601f83011215613d5f57813561440081613ecf565b9261440e6040519485613b4c565b81845260208085019260051b820101928311613d5f57602001905b8282106144365750505090565b6020809161444384613d87565b815201910190614429565b63ffffffff1680614471575b506007549061ffff63ffffffff83169260201c1690565b90604051600854808252816020810160086000526020600020926000905b80600f83011061463b576145159454918181106134ff578181106134e7578181106134d0578181106134b8578181106134a0578181106134885781811061347057818110613458578181106134405781811061342857818110613410578181106133f8578181106133e0578181106133c8578181106133b057106133a257500382613b4c565b60005b815180821015614630577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161460157811480156145ea575b61456157600101614518565b92505060095482101561417b578160031c7f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0154600a5483101561417b5761ffff90600a60005260f063ffffffff8560041c7fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801549260e08760051b161c169460041b161c1690565b5061ffff6145f882846143be565b51168410614555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505090503861445a565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e082015201940192018492939161448f565b805482101561417b5760005260206000200190600090565b60008281526001820160205260409020546147c55780549068010000000000000000821015613ac957826147ae614779846001809601855584614723565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014614926577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614601578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614601578181036148ef575b505050805480156148c0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906148818282614723565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61490f6148ff6147799386614723565b90549060031b1c92839286614723565b905560005283602052604060002055388080614849565b50505050600090565b919290156149aa5750815115614943575090565b3b1561494c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156149bd5750805190602001fd5b612f12906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613bc756fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
}

var CCTPVerifierABI = CCTPVerifierMetaData.ABI

var CCTPVerifierBin = CCTPVerifierMetaData.Bin

func DeployCCTPVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, messageTransmitterProxy common.Address, usdcToken common.Address, storageLocation string, finalityConfig CCTPVerifierFinalityConfig, dynamicConfig CCTPVerifierDynamicConfig) (common.Address, *types.Transaction, *CCTPVerifier, error) {
	parsed, err := CCTPVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPVerifierBin), backend, tokenMessenger, messageTransmitterProxy, usdcToken, storageLocation, finalityConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPVerifier{address: address, abi: *parsed, CCTPVerifierCaller: CCTPVerifierCaller{contract: contract}, CCTPVerifierTransactor: CCTPVerifierTransactor{contract: contract}, CCTPVerifierFilterer: CCTPVerifierFilterer{contract: contract}}, nil
}

type CCTPVerifier struct {
	address common.Address
	abi     abi.ABI
	CCTPVerifierCaller
	CCTPVerifierTransactor
	CCTPVerifierFilterer
}

type CCTPVerifierCaller struct {
	contract *bind.BoundContract
}

type CCTPVerifierTransactor struct {
	contract *bind.BoundContract
}

type CCTPVerifierFilterer struct {
	contract *bind.BoundContract
}

type CCTPVerifierSession struct {
	Contract     *CCTPVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPVerifierCallerSession struct {
	Contract *CCTPVerifierCaller
	CallOpts bind.CallOpts
}

type CCTPVerifierTransactorSession struct {
	Contract     *CCTPVerifierTransactor
	TransactOpts bind.TransactOpts
}

type CCTPVerifierRaw struct {
	Contract *CCTPVerifier
}

type CCTPVerifierCallerRaw struct {
	Contract *CCTPVerifierCaller
}

type CCTPVerifierTransactorRaw struct {
	Contract *CCTPVerifierTransactor
}

func NewCCTPVerifier(address common.Address, backend bind.ContractBackend) (*CCTPVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifier{address: address, abi: abi, CCTPVerifierCaller: CCTPVerifierCaller{contract: contract}, CCTPVerifierTransactor: CCTPVerifierTransactor{contract: contract}, CCTPVerifierFilterer: CCTPVerifierFilterer{contract: contract}}, nil
}

func NewCCTPVerifierCaller(address common.Address, caller bind.ContractCaller) (*CCTPVerifierCaller, error) {
	contract, err := bindCCTPVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierCaller{contract: contract}, nil
}

func NewCCTPVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPVerifierTransactor, error) {
	contract, err := bindCCTPVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierTransactor{contract: contract}, nil
}

func NewCCTPVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPVerifierFilterer, error) {
	contract, err := bindCCTPVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierFilterer{contract: contract}, nil
}

func bindCCTPVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPVerifier *CCTPVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPVerifier.Contract.CCTPVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPVerifier *CCTPVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.CCTPVerifierTransactor.contract.Transfer(opts)
}

func (_CCTPVerifier *CCTPVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.CCTPVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPVerifier *CCTPVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPVerifier *CCTPVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.contract.Transfer(opts)
}

func (_CCTPVerifier *CCTPVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCTPVerifier.Contract.GetDestChainConfig(&_CCTPVerifier.CallOpts, destChainSelector)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCTPVerifier.Contract.GetDestChainConfig(&_CCTPVerifier.CallOpts, destChainSelector)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetDomain(opts *bind.CallOpts, chainSelector uint64) (CCTPVerifierDomain, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getDomain", chainSelector)

	if err != nil {
		return *new(CCTPVerifierDomain), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPVerifierDomain)).(*CCTPVerifierDomain)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetDomain(chainSelector uint64) (CCTPVerifierDomain, error) {
	return _CCTPVerifier.Contract.GetDomain(&_CCTPVerifier.CallOpts, chainSelector)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetDomain(chainSelector uint64) (CCTPVerifierDomain, error) {
	return _CCTPVerifier.Contract.GetDomain(&_CCTPVerifier.CallOpts, chainSelector)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (CCTPVerifierDynamicConfig, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CCTPVerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPVerifierDynamicConfig)).(*CCTPVerifierDynamicConfig)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetDynamicConfig() (CCTPVerifierDynamicConfig, error) {
	return _CCTPVerifier.Contract.GetDynamicConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetDynamicConfig() (CCTPVerifierDynamicConfig, error) {
	return _CCTPVerifier.Contract.GetDynamicConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CCTPVerifier.Contract.GetFee(&_CCTPVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CCTPVerifier.Contract.GetFee(&_CCTPVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetFinalityConfig(opts *bind.CallOpts) (CCTPVerifierFinalityConfig, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getFinalityConfig")

	if err != nil {
		return *new(CCTPVerifierFinalityConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPVerifierFinalityConfig)).(*CCTPVerifierFinalityConfig)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetFinalityConfig() (CCTPVerifierFinalityConfig, error) {
	return _CCTPVerifier.Contract.GetFinalityConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetFinalityConfig() (CCTPVerifierFinalityConfig, error) {
	return _CCTPVerifier.Contract.GetFinalityConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetStaticConfig(opts *bind.CallOpts) (CCTPVerifierStaticConfig, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CCTPVerifierStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCTPVerifierStaticConfig)).(*CCTPVerifierStaticConfig)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetStaticConfig() (CCTPVerifierStaticConfig, error) {
	return _CCTPVerifier.Contract.GetStaticConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetStaticConfig() (CCTPVerifierStaticConfig, error) {
	return _CCTPVerifier.Contract.GetStaticConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetStorageLocation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getStorageLocation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetStorageLocation() (string, error) {
	return _CCTPVerifier.Contract.GetStorageLocation(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetStorageLocation() (string, error) {
	return _CCTPVerifier.Contract.GetStorageLocation(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) Owner() (common.Address, error) {
	return _CCTPVerifier.Contract.Owner(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) Owner() (common.Address, error) {
	return _CCTPVerifier.Contract.Owner(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPVerifier.Contract.SupportsInterface(&_CCTPVerifier.CallOpts, interfaceId)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPVerifier.Contract.SupportsInterface(&_CCTPVerifier.CallOpts, interfaceId)
}

func (_CCTPVerifier *CCTPVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) TypeAndVersion() (string, error) {
	return _CCTPVerifier.Contract.TypeAndVersion(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) TypeAndVersion() (string, error) {
	return _CCTPVerifier.Contract.TypeAndVersion(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) VersionTag() ([4]byte, error) {
	return _CCTPVerifier.Contract.VersionTag(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) VersionTag() ([4]byte, error) {
	return _CCTPVerifier.Contract.VersionTag(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPVerifier *CCTPVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPVerifier.Contract.AcceptOwnership(&_CCTPVerifier.TransactOpts)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPVerifier.Contract.AcceptOwnership(&_CCTPVerifier.TransactOpts)
}

func (_CCTPVerifier *CCTPVerifierTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CCTPVerifier *CCTPVerifierSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyAllowlistUpdates(&_CCTPVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyAllowlistUpdates(&_CCTPVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CCTPVerifier *CCTPVerifierTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyDestChainConfigUpdates(&_CCTPVerifier.TransactOpts, destChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyDestChainConfigUpdates(&_CCTPVerifier.TransactOpts, destChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "forwardToVerifier", message, messageId, arg2, arg3, arg4)
}

func (_CCTPVerifier *CCTPVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ForwardToVerifier(&_CCTPVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ForwardToVerifier(&_CCTPVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_CCTPVerifier *CCTPVerifierTransactor) SetDomains(opts *bind.TransactOpts, domains []CCTPVerifierSetDomainArgs) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "setDomains", domains)
}

func (_CCTPVerifier *CCTPVerifierSession) SetDomains(domains []CCTPVerifierSetDomainArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetDomains(&_CCTPVerifier.TransactOpts, domains)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) SetDomains(domains []CCTPVerifierSetDomainArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetDomains(&_CCTPVerifier.TransactOpts, domains)
}

func (_CCTPVerifier *CCTPVerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCTPVerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CCTPVerifier *CCTPVerifierSession) SetDynamicConfig(dynamicConfig CCTPVerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetDynamicConfig(&_CCTPVerifier.TransactOpts, dynamicConfig)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) SetDynamicConfig(dynamicConfig CCTPVerifierDynamicConfig) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetDynamicConfig(&_CCTPVerifier.TransactOpts, dynamicConfig)
}

func (_CCTPVerifier *CCTPVerifierTransactor) SetFinalityConfig(opts *bind.TransactOpts, finalityConfig CCTPVerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "setFinalityConfig", finalityConfig)
}

func (_CCTPVerifier *CCTPVerifierSession) SetFinalityConfig(finalityConfig CCTPVerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetFinalityConfig(&_CCTPVerifier.TransactOpts, finalityConfig)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) SetFinalityConfig(finalityConfig CCTPVerifierFinalityConfig) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetFinalityConfig(&_CCTPVerifier.TransactOpts, finalityConfig)
}

func (_CCTPVerifier *CCTPVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPVerifier *CCTPVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.TransferOwnership(&_CCTPVerifier.TransactOpts, to)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.TransferOwnership(&_CCTPVerifier.TransactOpts, to)
}

func (_CCTPVerifier *CCTPVerifierTransactor) UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "updateStorageLocation", newLocation)
}

func (_CCTPVerifier *CCTPVerifierSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.UpdateStorageLocation(&_CCTPVerifier.TransactOpts, newLocation)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.UpdateStorageLocation(&_CCTPVerifier.TransactOpts, newLocation)
}

func (_CCTPVerifier *CCTPVerifierTransactor) VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "verifyMessage", message, messageHash, ccvData)
}

func (_CCTPVerifier *CCTPVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.VerifyMessage(&_CCTPVerifier.TransactOpts, message, messageHash, ccvData)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.VerifyMessage(&_CCTPVerifier.TransactOpts, message, messageHash, ccvData)
}

func (_CCTPVerifier *CCTPVerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CCTPVerifier *CCTPVerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.WithdrawFeeTokens(&_CCTPVerifier.TransactOpts, feeTokens)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.WithdrawFeeTokens(&_CCTPVerifier.TransactOpts, feeTokens)
}

type CCTPVerifierAllowListSendersAddedIterator struct {
	Event *CCTPVerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierAllowListSendersAdded)
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
		it.Event = new(CCTPVerifierAllowListSendersAdded)
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

func (it *CCTPVerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierAllowListSendersAddedIterator{contract: _CCTPVerifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierAllowListSendersAdded)
				if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*CCTPVerifierAllowListSendersAdded, error) {
	event := new(CCTPVerifierAllowListSendersAdded)
	if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierAllowListSendersRemovedIterator struct {
	Event *CCTPVerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierAllowListSendersRemoved)
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
		it.Event = new(CCTPVerifierAllowListSendersRemoved)
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

func (it *CCTPVerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierAllowListSendersRemovedIterator{contract: _CCTPVerifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierAllowListSendersRemoved)
				if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*CCTPVerifierAllowListSendersRemoved, error) {
	event := new(CCTPVerifierAllowListSendersRemoved)
	if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierDestChainConfigSetIterator struct {
	Event *CCTPVerifierDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierDestChainConfigSet)
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
		it.Event = new(CCTPVerifierDestChainConfigSet)
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

func (it *CCTPVerifierDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierDestChainConfigSetIterator{contract: _CCTPVerifier.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierDestChainConfigSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseDestChainConfigSet(log types.Log) (*CCTPVerifierDestChainConfigSet, error) {
	event := new(CCTPVerifierDestChainConfigSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierDomainsSetIterator struct {
	Event *CCTPVerifierDomainsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierDomainsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierDomainsSet)
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
		it.Event = new(CCTPVerifierDomainsSet)
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

func (it *CCTPVerifierDomainsSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierDomainsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierDomainsSet struct {
	Domains []CCTPVerifierSetDomainArgs
	Raw     types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterDomainsSet(opts *bind.FilterOpts) (*CCTPVerifierDomainsSetIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierDomainsSetIterator{contract: _CCTPVerifier.contract, event: "DomainsSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDomainsSet) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierDomainsSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "DomainsSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseDomainsSet(log types.Log) (*CCTPVerifierDomainsSet, error) {
	event := new(CCTPVerifierDomainsSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "DomainsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierDynamicConfigSetIterator struct {
	Event *CCTPVerifierDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierDynamicConfigSet)
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
		it.Event = new(CCTPVerifierDynamicConfigSet)
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

func (it *CCTPVerifierDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierDynamicConfigSet struct {
	DynamicConfig CCTPVerifierDynamicConfig
	Raw           types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPVerifierDynamicConfigSetIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierDynamicConfigSetIterator{contract: _CCTPVerifier.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierDynamicConfigSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseDynamicConfigSet(log types.Log) (*CCTPVerifierDynamicConfigSet, error) {
	event := new(CCTPVerifierDynamicConfigSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierFeeTokenWithdrawnIterator struct {
	Event *CCTPVerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierFeeTokenWithdrawn)
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
		it.Event = new(CCTPVerifierFeeTokenWithdrawn)
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

func (it *CCTPVerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPVerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierFeeTokenWithdrawnIterator{contract: _CCTPVerifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierFeeTokenWithdrawn)
				if err := _CCTPVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCTPVerifierFeeTokenWithdrawn, error) {
	event := new(CCTPVerifierFeeTokenWithdrawn)
	if err := _CCTPVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierFinalityConfigSetIterator struct {
	Event *CCTPVerifierFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierFinalityConfigSet)
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
		it.Event = new(CCTPVerifierFinalityConfigSet)
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

func (it *CCTPVerifierFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierFinalityConfigSet struct {
	FinalityConfig CCTPVerifierFinalityConfig
	Raw            types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPVerifierFinalityConfigSetIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierFinalityConfigSetIterator{contract: _CCTPVerifier.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierFinalityConfigSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseFinalityConfigSet(log types.Log) (*CCTPVerifierFinalityConfigSet, error) {
	event := new(CCTPVerifierFinalityConfigSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierOwnershipTransferRequestedIterator struct {
	Event *CCTPVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierOwnershipTransferRequested)
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
		it.Event = new(CCTPVerifierOwnershipTransferRequested)
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

func (it *CCTPVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierOwnershipTransferRequestedIterator{contract: _CCTPVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierOwnershipTransferRequested)
				if err := _CCTPVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPVerifierOwnershipTransferRequested, error) {
	event := new(CCTPVerifierOwnershipTransferRequested)
	if err := _CCTPVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierOwnershipTransferredIterator struct {
	Event *CCTPVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierOwnershipTransferred)
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
		it.Event = new(CCTPVerifierOwnershipTransferred)
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

func (it *CCTPVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierOwnershipTransferredIterator{contract: _CCTPVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierOwnershipTransferred)
				if err := _CCTPVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPVerifierOwnershipTransferred, error) {
	event := new(CCTPVerifierOwnershipTransferred)
	if err := _CCTPVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierStaticConfigSetIterator struct {
	Event *CCTPVerifierStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierStaticConfigSet)
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
		it.Event = new(CCTPVerifierStaticConfigSet)
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

func (it *CCTPVerifierStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierStaticConfigSet struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	UsdcToken               common.Address
	LocalDomainIdentifier   uint32
	Raw                     types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*CCTPVerifierStaticConfigSetIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierStaticConfigSetIterator{contract: _CCTPVerifier.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierStaticConfigSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseStaticConfigSet(log types.Log) (*CCTPVerifierStaticConfigSet, error) {
	event := new(CCTPVerifierStaticConfigSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPVerifierStorageLocationUpdatedIterator struct {
	Event *CCTPVerifierStorageLocationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierStorageLocationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierStorageLocationUpdated)
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
		it.Event = new(CCTPVerifierStorageLocationUpdated)
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

func (it *CCTPVerifierStorageLocationUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierStorageLocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierStorageLocationUpdated struct {
	OldLocation string
	NewLocation string
	Raw         types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CCTPVerifierStorageLocationUpdatedIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierStorageLocationUpdatedIterator{contract: _CCTPVerifier.contract, event: "StorageLocationUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStorageLocationUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierStorageLocationUpdated)
				if err := _CCTPVerifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseStorageLocationUpdated(log types.Log) (*CCTPVerifierStorageLocationUpdated, error) {
	event := new(CCTPVerifierStorageLocationUpdated)
	if err := _CCTPVerifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
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

func (CCTPVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CCTPVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CCTPVerifierDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CCTPVerifierDomainsSet) Topic() common.Hash {
	return common.HexToHash("0x4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a30046")
}

func (CCTPVerifierDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b")
}

func (CCTPVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPVerifierFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd4")
}

func (CCTPVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCTPVerifierStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c")
}

func (CCTPVerifierStorageLocationUpdated) Topic() common.Hash {
	return common.HexToHash("0xbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8")
}

func (_CCTPVerifier *CCTPVerifier) Address() common.Address {
	return _CCTPVerifier.address
}

type CCTPVerifierInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (CCTPVerifierDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCTPVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetFinalityConfig(opts *bind.CallOpts) (CCTPVerifierFinalityConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (CCTPVerifierStaticConfig, error)

	GetStorageLocation(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []CCTPVerifierSetDomainArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCTPVerifierDynamicConfig) (*types.Transaction, error)

	SetFinalityConfig(opts *bind.TransactOpts, finalityConfig CCTPVerifierFinalityConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CCTPVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CCTPVerifierAllowListSendersRemoved, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CCTPVerifierDestChainConfigSet, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*CCTPVerifierDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*CCTPVerifierDomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPVerifierDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPVerifierDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPVerifierFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPVerifierFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CCTPVerifierFinalityConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPVerifierOwnershipTransferred, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*CCTPVerifierStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*CCTPVerifierStaticConfigSet, error)

	FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CCTPVerifierStorageLocationUpdatedIterator, error)

	WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStorageLocationUpdated) (event.Subscription, error)

	ParseStorageLocationUpdated(log types.Log) (*CCTPVerifierStorageLocationUpdated, error)

	Address() common.Address
}
