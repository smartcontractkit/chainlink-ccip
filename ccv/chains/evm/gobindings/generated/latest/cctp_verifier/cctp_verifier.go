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
	MintRecipient         [32]byte
	DomainIdentifier      uint32
	Enabled               bool
}

type CCTPVerifierDomainUpdate struct {
	AllowedCallerOnDest   [32]byte
	AllowedCallerOnSource [32]byte
	MintRecipient         [32]byte
	DomainIdentifier      uint32
	DestChainSelector     uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.StaticConfig\",\"components\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalityConfig\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomFinalitiesMustBeStrictlyIncreasing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomFinalityArraysMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDomainUpdate\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DomainUpdate\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MissingCustomFinalities\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101208060405234610c3a57615bbc803803809161001d8285610fb5565b83398101908082039060e08212610c3a578051906001600160a01b03821690818303610c3a576020810151926001600160a01b038416808503610c3a5760408301516001600160a01b0381169390848103610c3a5760608201516001600160401b038111610c3a5782019289601f85011215610c3a5783519361009f85610fd8565b946100ad6040519687610fb5565b8086528b60208284010111610c3a576100cc9160208088019101610ff3565b60808301516001600160401b038111610c3a5783019760a0898c0312610c3a576040519560a087016001600160401b038111888210176109b1576040526101128a611016565b875261012060208b01611027565b6020880190815260408b01519099906001600160401b038111610c3a578d610149918d0161104d565b6040890190815260608c0151909d909b6001600160401b038d11610c3a5781601f9d82019d8e011215610c3a578c519c6101828e611036565b9d6040519e8f906101939082610fb5565b8181526020019060051b820160200191848311610c3a57602001905b828210610f9d5750505060608a019c8d526080810151916001600160401b038311610c3a576040926101e1920161104d565b60808a019081529c609f190112610c3a5760408051969087016001600160401b038111888210176109b15761022a9160c09160405261022260a082016110b2565b8952016110b2565b96602087019788523315610f8c57600180546001600160a01b0319163317905560405160035490919060008361025f8361113f565b80825291600184168015610f6e57600114610f0e575b61028192500384610fb5565b8151906001600160401b0382116109b15761029b9061113f565b601f8111610eb4575b506020601f8211600114610e37576102fd9282600080516020615b9c833981519152959361030b93600091610e2c575b508160011b916000199060031b1c1916176003555b604051938493604085526040850190611179565b908382036020850152611179565b0390a180158015610e24575b8015610e1c575b610bad57604051639cdbb18160e01b8152602081600481855afa8015610c8f57600090610ddf575b63ffffffff91501660018103610dc65750602060049160405192838092632c12192160e01b82525afa908115610c8f57600091610d8c575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa8015610c8f57600090610d4f575b63ffffffff91501660018103610d3657506020600491604051928380926367e0ed8360e11b82525afa908115610c8f57600091610ced575b506001600160a01b0316838103610cd5575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa908115610c8f57600091610c9b575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529490602081604481855afa908115610c8f57600091610c5d575b506000198101809111610c47576105219160405191602083019763095ea7b360e01b895260248401526044830152604482526104b1606483610fb5565b6000806040988951946104c48b87610fb5565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15610c3f573d9161050583610fd8565b926105128a519485610fb5565b83523d6000602085013e61119e565b805180610bbe575b505060e05160c05160a0516101005187516001600160a01b039485168152928416602084015292168187015263ffffffff90911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c90608090a180516001600160a01b031615610bad5751600580546001600160a01b03199081166001600160a01b03938416908117909255835160068054909216908416179055845190815291511660208201527f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b908390a185515115610b9c57855151845151811490811591610b8f575b50610b7e57600093845b8751805187101561066c5761ffff6106378882936110c6565b51169116101561065b57600161ffff610651878a516110c6565b511695019461061e565b633fd3668360e21b60005260046000fd5b505085878563ffffffff8551166007549065ffff00000000835160201b169165ffffffffffff19161717600755815180519060018060401b0382116109b1576801000000000000000082116109b15760209060085483600855808410610b48575b500190600860005260206000208160041c9160005b838110610b075750600f198116900380610ab8575b505085518051925090506001600160401b0382116109b1576801000000000000000082116109b15760209060095483600955808410610a5e575b500190600960005260206000208160031c9160005b838110610a1b575060071981169003806109c7575b505084518051925090506001600160401b0382116109b1576801000000000000000082116109b157602090600a5483600a55808410610957575b500190600a60005260206000208160041c9160005b8381106109165750600f1981169003806108c3575b505050509061ffff6107f69263ffffffff88519760208952511660208801525116868601525160a0606086015260c0850190611107565b915191601f198482030160808501526020808451928381520193019060005b8181106108a757867f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd48780610857898951601f198483030160a0850152611107565b0390a151614960908161123c823960805181505060a051818181612b980152613791015260c0518181816117200152613764015260e051818181612c950152613713015261010051816137b80152f35b825163ffffffff16855260209485019490920191600101610815565b9260009360005b8181106108e357505050015561ffff6107f688806107bf565b909194602061090c60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b96019291016108ca565b6000805b6010811061092f5750838201556001016107aa565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161091a565b61098890600a60005283600020600f80870160041c820192601e8860011b168061098e575b500160041c01906110f0565b89610795565b6109ab906000198601908154906000199060200360031b1c169055565b8e61097c565b634e487b7160e01b600052604160045260246000fd5b9260009360005b8181106109e35750505001558680808061075b565b9091946020610a1160019263ffffffff895116908560021b63ffffffff809160031b9316831b921b19161790565b96019291016109ce565b6000805b60088110610a34575083820155600101610746565b865190969160019160209163ffffffff60058b901b81811b199092169216901b1792019601610a1f565b610a8f90600960005283600020600780870160031c820192601c8860021b1680610a95575b500160031c01906110f0565b89610731565b610ab2906000198601908154906000199060200360031b1c169055565b8e610a83565b9260009360005b818110610ad4575050500155868080806106f7565b9091946020610afd60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101610abf565b6000805b60108110610b205750838201556001016106e2565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601610b0b565b610b7890600860005283600020600f80870160041c820192601e8860011b168061098e57500160041c01906110f0565b896106cd565b630cacb92960e41b60005260046000fd5b9050855151141538610614565b636c33513960e11b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b8160209181010312610c3a5760200151801590811503610c3a57610be3573880610529565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b60609161119e565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610c87575b81610c7860209383610fb5565b81010312610c3a575138610474565b3d9150610c6b565b6040513d6000823e3d90fd5b90506020813d602011610ccd575b81610cb660209383610fb5565b81010312610c3a57610cc790611016565b3861042a565b3d9150610ca9565b836383395ca960e01b60005260045260245260446000fd5b6020813d602011610d2e575b81610d0660209383610fb5565b81010312610d2a5751906001600160a01b0382168203610d275750386103e9565b80fd5b5080fd5b3d9150610cf9565b6331b6aa1b60e11b600052600160045260245260446000fd5b506020813d602011610d84575b81610d6960209383610fb5565b81010312610c3a57610d7f63ffffffff91611016565b6103b1565b3d9150610d5c565b90506020813d602011610dbe575b81610da760209383610fb5565b81010312610c3a57610db8906110b2565b3861037e565b3d9150610d9a565b633785f8f160e01b600052600160045260245260446000fd5b506020813d602011610e14575b81610df960209383610fb5565b81010312610c3a57610e0f63ffffffff91611016565b610346565b3d9150610dec565b50881561031e565b508315610317565b9050820151386102d4565b601f198216906003600052806000209160005b818110610e9c57508361030b936102fd9693600080516020615b9c833981519152989660019410610e83575b5050811b016003556102e9565b84015160001960f88460031b161c191690553880610e76565b91926020600181928689015181550194019201610e4a565b6003600052610efe907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610f04575b601f0160051c01906110f0565b386102a4565b9091508190610ef1565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310610f5257505090602061028192820101610275565b6020919350806001915483858a01015201910190918592610f3a565b505060206102819260ff19851682840152151560051b820101610275565b639b15e16f60e01b60005260046000fd5b60208091610faa84611016565b8152019101906101af565b601f909101601f19168101906001600160401b038211908210176109b157604052565b6001600160401b0381116109b157601f01601f191660200190565b60005b8381106110065750506000910152565b8181015183820152602001610ff6565b519063ffffffff82168203610c3a57565b519061ffff82168203610c3a57565b6001600160401b0381116109b15760051b60200190565b9080601f83011215610c3a57815161106481611036565b926110726040519485610fb5565b81845260208085019260051b820101928311610c3a57602001905b82821061109a5750505090565b602080916110a784611027565b81520191019061108d565b51906001600160a01b0382168203610c3a57565b80518210156110da5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8181106110fb575050565b600081556001016110f0565b906020808351928381520192019060005b8181106111255750505090565b825161ffff16845260209384019390920191600101611118565b90600182811c9216801561116f575b602083101461115957565b634e487b7160e01b600052602260045260246000fd5b91607f169161114e565b9060209161119281518092818552858086019101610ff3565b601f01601f1916010190565b9192901561120057508151156111b2575090565b3b156111bb5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156112135750805190602001fd5b60405162461bcd60e51b815260206004820152908190611237906024830190611179565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146137f75750806306285c69146136ac578063181f5a771461362d57806338a9eb2a146130b45780633bbbed4b14612943578063445b89d1146126435780635cb80c5d146123625780636def4ce71461228d5780637437ff9f146121ce57806379ba5097146120e957806380485e2514611e5d578063869b7f6214611cdb5780638da5cb5b14611c89578063b2bd751c14611943578063bff0ec1d146112d5578063c757d73f14610a72578063c9b146b31461067c578063ceac5cee14610392578063dfadfa35146102ad578063f2fde38b146101c0578063fe163eed146101675763fec888af1461011357600080fd5b3461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645761016061014c614379565b6040519182916020835260208301906139cf565b0390f35b80fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645773ffffffffffffffffffffffffffffffffffffffff61020d613b6c565b61021561444e565b1633811461028557807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457604060a09167ffffffffffffffff6102f3613c40565b826080855161030181613900565b828152826020820152828782015282606082015201521681526004602052206040519061032d82613900565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff81116106785736602382011215610678576103f3903690602481600401359101613cb6565b6103fb61444e565b610403614379565b90805167ffffffffffffffff811161064b57610420600354614326565b601f81116105b4575b506020601f82116001146104cf576104b092827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd895936104be9388916104c4575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b6040519384936040855260408501906139cf565b9083820360208501526139cf565b0390a180f35b90508201513861046a565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b81811061059c5750836104be936104b096937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610565575b5050811b0160035561049c565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610558565b9192602060018192868901518155019401920161051a565b61061d9060038652601f830160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208410610623575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01906144ad565b38610429565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b91506105ef565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5080fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff8111610678576106cc903690600401613c0f565b73ffffffffffffffffffffffffffffffffffffffff6001541633141580610a50575b610a2857919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b84811015610a24578060051b82013583811215610a2057820191608083360312610a20576040519461074e866138b5565b61075784613c57565b865261076560208501614301565b9660208701978852604085013567ffffffffffffffff8111610a1c5761078e90369087016144c4565b9460408801958652606081013567ffffffffffffffff8111610a18576107b6913691016144c4565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610830818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b8151516108f0575b5095976001019550815b85518051821015610881579061087a73ffffffffffffffffffffffffffffffffffffffff61087283600195614499565b511689614724565b5001610842565b505095909694506001929193519081516108a1575b50500193929361071d565b6108e667ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613c6c565b0390a23880610896565b989395929691909497986000146109e157600184019591875b86518051821015610986576109338273ffffffffffffffffffffffffffffffffffffffff92614499565b5116801561094f57906109486001928a614693565b5001610909565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816109d767ffffffffffffffff8a51169251604051918291602083526020830190613c6c565b0390a23880610838565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b8280fd5b8480fd5b8380f35b6004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff600654163314156106ee565b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff81116106785760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126106785760405191610aed83613900565b610af982600401613d40565b8352610b0760248301613d31565b9260208101938452604483013567ffffffffffffffff8111610a1c57610b339060043691860101613e2b565b9360408201948552606484013567ffffffffffffffff8111610a185784019336602386011215610a1857600485013594610b6c86613d08565b95610b7a6040519788613954565b808752602060048189019260051b84010101913683116112d157602401905b8282106112b95750505060608301948552608481013567ffffffffffffffff8111610a2057610bcd91369101600401613e2b565b9160808101928352610bdd61444e565b8551511561129157855151855151811490811591611284575b5061125c578392835b87518051861015610c645761ffff610c18878293614499565b511691161015610c3c57600161ffff610c32868a51614499565b5116940193610bff565b6004857fff4d9a0c000000000000000000000000000000000000000000000000000000008152fd5b50508587869463ffffffff8551167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006007541617600755517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff65ffff000000006007549260201b169116176007555180519067ffffffffffffffff821161122f5768010000000000000000821161122f576020906008548360085580841061119d575b50019060088652602086208160041c91875b83811061115d57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380611110575b505050505180519067ffffffffffffffff82116110e3576801000000000000000082116110e35760209060095483600955808410611050575b50019060098552602085208160031c91865b83811061100e57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff88116900380610fbc575b505050505180519067ffffffffffffffff821161064b5768010000000000000000821161064b57602090600a5483600a55808410610f29575b500190600a8452602084208160041c91855b838110610ee957507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380610e7a575b857f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd46104be8760405191829182613a66565b928593865b818110610eb65750505001556104be7f493a888e94d423db3eb836e046ad4a5466a0184af32683f7b3b705087c49ebd48480610e48565b9091946020610edf60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101610e7f565b86875b60108110610f01575083820155600101610e15565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601610eec565b610f5890600a8752838720600f80870160041c820192601e8860011b1680610f5e575b500160041c01906144ad565b85610e03565b610fb6907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8a610f4c565b928693875b818110610fd657505050015583808080610dca565b909194602061100460019263ffffffff895116908560021b63ffffffff809160031b9316831b921b19161790565b9601929101610fc1565b87885b60088110611026575083820155600101610d97565b865190969160019160209163ffffffff60058b901b81811b199092169216901b1792019601611011565b61107f9060098852838820600780870160031c820192601c8860021b1680611085575b500160031c01906144ad565b86610d85565b6110dd907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8b611073565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b928793885b81811061112a57505050015584808080610d4c565b909194602061115360019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101611115565b88895b60108110611175575083820155600101610d19565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601611160565b6111cb9060088952838920600f80870160041c820192601e8860011b16806111d157500160041c01906144ad565b87610d07565b611229907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8601907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160200360031b1c169055565b8c610f4c565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6004847fcacb9290000000000000000000000000000000000000000000000000000000008152fd5b9050835151141538610bf6565b6004847fd866a272000000000000000000000000000000000000000000000000000000008152fd5b602080916112c684613d40565b815201910190610b99565b8680fd5b50346101645760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff8111610678576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261067857604051906101c0820182811067ffffffffffffffff82111761064b5760405261136f81600401613c57565b825261137d60248201613c57565b602083015261138e60448201613c57565b604083015261139f60648201613d40565b60608301526113b060848201613d40565b60808301526113c160a48201613d31565b60a083015260c481013560c083015260e481013567ffffffffffffffff8111610a18576113f49060043691840101613ced565b60e083015261010481013567ffffffffffffffff8111610a185761141e9060043691840101613ced565b61010083015261012481013567ffffffffffffffff8111610a18576114499060043691840101613ced565b61012083015261014481013567ffffffffffffffff8111610a18576114749060043691840101613ced565b61014083015261016481013567ffffffffffffffff8111610a185761149f9060043691840101613ced565b61016083015261018481013567ffffffffffffffff8111610a1857810136602382011215610a18576004810135906114d682613d08565b916114e46040519384613954565b80835260051b8101602401602083013682116112d15760248301905b82821061190d57505050506101808301526101a48101359067ffffffffffffffff8211610a185760046115369236920101613ced565b6101a082015260243560443567ffffffffffffffff8111610a185761155f903690600401613bb0565b6101e181106118e55780600411610a20577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361189657508061018011610a205760049360009061017c8401357fffffffff00000000000000000000000000000000000000000000000000000000167f71e2e5630000000000000000000000000000000000000000000000000000000081016118485750826101a0116112d1576020946101808501359081810361181b57505067ffffffffffffffff8151168752858552604087209060ff6003830154871c16156117e657508261011c116112d15760fc84013590600101548082036117b957505061019c92611706859383889450507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe606101a08401910160405196879586957f57ecfd280000000000000000000000000000000000000000000000000000000087526040828801528260448801520160648601378961020085015261020484019161020060248601526142c2565b03818773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19182156117ae578492611781575b50501561175b575080f35b907fbc40f556000000000000000000000000000000000000000000000000000000008152fd5b6117a09250803d106117a7575b6117988183613954565b81019061430e565b3880611750565b503d61178e565b6040513d86823e3d90fd5b7f7d8b101a0000000000000000000000000000000000000000000000000000000088528652602452604486fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000885267ffffffffffffffff168652602487fd5b7f6c86fa3a0000000000000000000000000000000000000000000000000000000089528752602452604487fd5b7fadaf77390000000000000000000000000000000000000000000000000000000088527f8e1d1a9d000000000000000000000000000000000000000000000000000000008752602452604487fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b6004857fbba6473c000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff811161193f57602091611934839283600436928a010101613d51565b815201910190611500565b8880fd5b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161067857611993903690600401613bde565b61199b61444e565b6119a481613d08565b916119b26040519384613954565b81835260c0602084019202810190368211610a2057915b818310611bec578480855b8051831015611be8576119e78382614499565b519267ffffffffffffffff60206119fe8385614499565b51015116938415611bbc578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff81511615611b905773ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a20191906119d4565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c083360312610a205760405190611c038261391c565b83359073ffffffffffffffffffffffffffffffffffffffff821682036112d1578260209260c09452611c36838701613c57565b83820152611c4660408701614301565b6040820152611c5760608701613d31565b6060820152611c6860808701613d40565b6080820152611c7960a08701613d40565b60a08201528152019201916119c9565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101645760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457604051611d1781613938565b611d1f613b6c565b815260243573ffffffffffffffffffffffffffffffffffffffff81168103610a1c5760208201908152611d5061444e565b73ffffffffffffffffffffffffffffffffffffffff82511615611e35578173ffffffffffffffffffffffffffffffffffffffff6104be92817f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346101645760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457611e95613c40565b60243567ffffffffffffffff8111610a1c5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610a1c5760405190611ee082613900565b806004013567ffffffffffffffff8111610a2057611f049060043691840101613ced565b8252602481013567ffffffffffffffff8111610a2057611f2a9060043691840101613ced565b6020830152604481013567ffffffffffffffff8111610a2057810136602382011215610a20576004810135611f5e81613d08565b91611f6c6040519384613954565b818352602060048185019360061b83010101903682116120e557602401915b8183106120ad575050506040830152611fa660648201613b8f565b6060830152608481013567ffffffffffffffff8111610a20576080916004611fd19236920101613ced565b9101526044359067ffffffffffffffff8211610a1c57611ffe67ffffffffffffffff923690600401613ced565b50612007613d20565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff6040822054161561208257816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b6040833603126120e557602060409182516120c781613938565b6120d086613b8f565b81528286013583820152815201920191611f8b565b8780fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457805473ffffffffffffffffffffffffffffffffffffffff811633036121a6577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457602060405161220b81613938565b828152015261016060405161221f81613938565b73ffffffffffffffffffffffffffffffffffffffff60055416815273ffffffffffffffffffffffffffffffffffffffff60065416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645767ffffffffffffffff6122ce613c40565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b81811061234c5773ffffffffffffffffffffffffffffffffffffffff856101608861232481890382613954565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613c6c565b82548452602090930192600192830192016122f7565b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff8111610678576123b2903690600401613c0f565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b8381101561263f578060051b8201359073ffffffffffffffffffffffffffffffffffffffff821680920361263b57604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa9283156126305787936125fd575b5082612453575b5060019150016123cf565b846040519361250e60208601957fa9059cbb0000000000000000000000000000000000000000000000000000000087528360248201528260448201526044815261249e606482613954565b8a806040988951936124b08b86613954565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d156125f5573d906124f182613995565b916124fe8a519384613954565b82523d8d602084013e5b86614887565b80518061254a575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338612448565b61256192949596935060208091830101910161430e565b156125725792919085903880612516565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090612508565b9092506020813d8211612628575b8161261860209383613954565b810103126112d157519138612441565b3d915061260b565b6040513d89823e3d90fd5b8580fd5b8480f35b50346101645760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760043567ffffffffffffffff811161067857612693903690600401613bde565b9061269c61444e565b825b8281101561288a5760c08102820160c081360312610a2057604051906126c38261391c565b803591828152602082013560208201918183526040810190604085013582526126ee60608601613d40565b926060820193845261270260808701613c57565b61271560a06080850198838a5201614301565b9760a0840198895215918215612881575b50811561286e575b50612808579263ffffffff9260039267ffffffffffffffff8560019a999897519451925193511697511515966040519461276786613900565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b1691161790550161269e565b8567ffffffffffffffff8663ffffffff8660c49689604051977f113b1fc2000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff915016153861272e565b15915038612726565b506040519180602084016020855252604083019190845b8181106128d257857fcbd7889dd51c92285faabb0cffe5578ccef74f7579f6ceec9f155dd054e4621486860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015263ffffffff61290460608901613d40565b16606082015267ffffffffffffffff61291f60808901613c57565b16608082015261293160a08801614301565b151560a08201520194019291016128a1565b50346101645760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610164576004359067ffffffffffffffff82116101645781600401906101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8436030112610164576129c0613b44565b5060843567ffffffffffffffff8111610678576129e1903690600401613bb0565b50506129f16101248401836141a0565b90357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116916014811061307f575b505060601c92602481019367ffffffffffffffff612a3d866141f1565b168084526002602052604084209081549073ffffffffffffffffffffffffffffffffffffffff8216908115613054576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015613049578690612fe6575b73ffffffffffffffffffffffffffffffffffffffff9150163303612fba5760f01c60ff16612f72575b505067ffffffffffffffff612aec856141f1565b1682526004602052604082209060038201549460ff8660201c1615612f34575061018481016001612b1d8287614206565b905003612efe57612b41612b3c612b348388614206565b36929161425a565b613d51565b906040820151602081519101517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612ec9575b505060601c9573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001696878103612e9e57506020612be3612bd9612bd38585614206565b9061425a565b60808101906141a0565b905003612e4d57505060028301548015612e2d5760a490915b519201359261ffff8416809403610a2057549260405192612c1c8461391c565b835260208301918252612c51604084019160243583526060850195865280608086015263ffffffff60a0860199168952614529565b61ffff8551911690818102918183041490151715612e0057612710900463ffffffff8111612dd55763ffffffff73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016955199511693519551925192604051937f8e1d1a9d000000000000000000000000000000000000000000000000000000006020860152602485015260248452612cff604485613954565b853b156120e5579263ffffffff88999a9381612d7c948b976040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c015260248b015260448a0152606489015260848801521660a48601521660c484015261010060e48401526101048301906139cf565b03925af1918215612dc8578161016093612db8575b505060405190612da2602083613954565b81526040519182916020835260208301906139cf565b612dc191613954565b3881612d91565b50604051903d90823e3d90fd5b7fb6f15d0f000000000000000000000000000000000000000000000000000000008752600452602486fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b506080810151602081805181010312610a2057602060a491015191612bfc565b612bd3612e5d92612bd992614206565b612e9a6040519283927fa3c8cf090000000000000000000000000000000000000000000000000000000084526020600485015260248401916142c2565b0390fd5b7f961c9a4f000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612b7b565b83612f0b60249287614206565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8367ffffffffffffffff612f496024936141f1565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612f8f5780612ad8565b7fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011613041575b8161300060209383613954565b8101031261263b575173ffffffffffffffffffffffffffffffffffffffff8116810361263b5773ffffffffffffffffffffffffffffffffffffffff90612aaf565b3d9150612ff3565b6040513d88823e3d90fd5b7f8a4e93c9000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612a20565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016457606060806040516130f381613900565b8381528360208201528260408201528280820152015260405161311581613900565b61ffff60075463ffffffff8116835260201c166020820152613135613e90565b604082015260405180816020600954928381520160098652602086209286905b8060078301106135a6576131ab94549181811061358f575b818110613576575b81811061355c575b818110613542575b818110613528575b81811061350e575b8181106134f4575b106134e6575b500382613954565b6060820152604051600a80548083529084526020820193907fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a85b81600f84011061340457946132819284926101609754918181106133ef575b8181106133d7575b8181106133c0575b8181106133a8575b818110613390575b818110613378575b818110613360575b818110613348575b818110613330575b818110613318575b818110613300575b8181106132e8575b8181106132d0575b8181106132b8575b8181106132a0575b1061329257500382613954565b608082015260405191829182613a66565b60f01c8152602001386131a3565b92602060019161ffff8560e01c168152019301613274565b92602060019161ffff8560d01c16815201930161326c565b92602060019161ffff8560c01c168152019301613264565b92602060019161ffff8560b01c16815201930161325c565b92602060019161ffff8560a01c168152019301613254565b92602060019161ffff8560901c16815201930161324c565b92602060019161ffff8560801c168152019301613244565b92602060019161ffff8560701c16815201930161323c565b92602060019161ffff8560601c168152019301613234565b92602060019161ffff8560501c16815201930161322c565b92602060019161ffff8560401c168152019301613224565b92602060019161ffff8560301c16815201930161321c565b92602060019161ffff85831c168152019301613214565b92602060019161ffff8560101c16815201930161320c565b92602060019161ffff85168152019301613204565b946001610200601092885461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e08201520196019201916131e5565b60e01c8152602001386131a3565b92602060019163ffffffff8560c01c16815201930161319d565b92602060019163ffffffff8560a01c168152019301613195565b92602060019163ffffffff8560801c16815201930161318d565b92602060019163ffffffff8560601c168152019301613185565b92602060019163ffffffff8560401c16815201930161317d565b92602060019163ffffffff85831c168152019301613175565b92602060019163ffffffff8516815201930161316d565b916008919350610100600191865463ffffffff8116825263ffffffff8160201c16602083015263ffffffff8160401c16604083015263ffffffff8160601c16606083015263ffffffff8160801c16608083015263ffffffff8160a01c1660a083015263ffffffff8160c01c1660c083015260e01c60e0820152019401920184929391613155565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610164575061016060405161366e604082613954565b601681527f43435450566572696669657220312e372e302d6465760000000000000000000060208201526040519182916020835260208301906139cf565b503461016457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645760606040516136e9816138b5565b8281528260208201528260408201520152608073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff604051613744816138b5565b82815273ffffffffffffffffffffffffffffffffffffffff6020820191817f00000000000000000000000000000000000000000000000000000000000000001683528160606040830192827f00000000000000000000000000000000000000000000000000000000000000001684520193857f0000000000000000000000000000000000000000000000000000000000000000168552604051968752511660208601525116604084015251166060820152f35b9050346106785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610678576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610a1c57602092507ffacbd7dc00000000000000000000000000000000000000000000000000000000811490811561388b575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613884565b6080810190811067ffffffffffffffff8211176138d157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176138d157604052565b60c0810190811067ffffffffffffffff8211176138d157604052565b6040810190811067ffffffffffffffff8211176138d157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176138d157604052565b67ffffffffffffffff81116138d157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613a195750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016139da565b906020808351928381520192019060005b818110613a4c5750505090565b825161ffff16845260209384019390920191600101613a3f565b9190916020815263ffffffff835116602082015261ffff6020840151166040820152613aa1604084015160a0606084015260c0830190613a2e565b906060840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030160808301526020808451928381520193019060005b818110613b28575050506080613b25939401519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613a2e565b90565b825163ffffffff16855260209485019490920191600101613ae2565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203613b6757565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613b6757565b359073ffffffffffffffffffffffffffffffffffffffff82168203613b6757565b9181601f84011215613b675782359167ffffffffffffffff8311613b675760208381860195010111613b6757565b9181601f84011215613b675782359167ffffffffffffffff8311613b675760208085019460c08502010111613b6757565b9181601f84011215613b675782359167ffffffffffffffff8311613b67576020808501948460051b010111613b6757565b6004359067ffffffffffffffff82168203613b6757565b359067ffffffffffffffff82168203613b6757565b906020808351928381520192019060005b818110613c8a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613c7d565b929192613cc282613995565b91613cd06040519384613954565b829481845281830111613b67578281602093846000960137010152565b9080601f83011215613b6757816020613b2593359101613cb6565b67ffffffffffffffff81116138d15760051b60200190565b6064359061ffff82168203613b6757565b359061ffff82168203613b6757565b359063ffffffff82168203613b6757565b919060c083820312613b675760405190613d6a8261391c565b819380358352602081013567ffffffffffffffff8111613b675782613d90918301613ced565b6020840152604081013567ffffffffffffffff8111613b675782613db5918301613ced565b6040840152606081013567ffffffffffffffff8111613b675782613dda918301613ced565b6060840152608081013567ffffffffffffffff8111613b675782613dff918301613ced565b608084015260a08101359167ffffffffffffffff8311613b675760a092613e269201613ced565b910152565b9080601f83011215613b67578135613e4281613d08565b92613e506040519485613954565b81845260208085019260051b820101928311613b6757602001905b828210613e785750505090565b60208091613e8584613d31565b815201910190613e6b565b60405190600854808352826020810160086000526020600020926000905b80600f8301106140b857613f449454918181106140a3575b81811061408b575b818110614074575b81811061405c575b818110614044575b81811061402c575b818110614014575b818110613ffc575b818110613fe4575b818110613fcc575b818110613fb4575b818110613f9c575b818110613f84575b818110613f6c575b818110613f54575b10613f46575b500383613954565b565b60f01c815260200138613f3c565b92602060019161ffff8560e01c168152019301613f36565b92602060019161ffff8560d01c168152019301613f2e565b92602060019161ffff8560c01c168152019301613f26565b92602060019161ffff8560b01c168152019301613f1e565b92602060019161ffff8560a01c168152019301613f16565b92602060019161ffff8560901c168152019301613f0e565b92602060019161ffff8560801c168152019301613f06565b92602060019161ffff8560701c168152019301613efe565b92602060019161ffff8560601c168152019301613ef6565b92602060019161ffff8560501c168152019301613eee565b92602060019161ffff8560401c168152019301613ee6565b92602060019161ffff8560301c168152019301613ede565b92602060019161ffff85831c168152019301613ed6565b92602060019161ffff8560101c168152019301613ece565b92602060019161ffff85168152019301613ec6565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019401920185929391613eae565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613b67570180359067ffffffffffffffff8211613b6757602001918136038313613b6757565b3567ffffffffffffffff81168103613b675790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613b67570180359067ffffffffffffffff8211613b6757602001918160051b36038313613b6757565b9015614293578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215613b67570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b35908115158203613b6757565b90816020910312613b6757518015158103613b675790565b90600182811c9216801561436f575b602083101461434057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614335565b604051906000826003549161438d83614326565b808352926001811690811561441157506001146143b1575b613f4492500383613954565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8183106143f5575050906020613f44928201016143a5565b60209193508060019154838589010152019101909184926143dd565b60209250613f449491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016143a5565b73ffffffffffffffffffffffffffffffffffffffff60015416330361446f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156142935760209160051b010190565b8181106144b8575050565b600081556001016144ad565b9080601f83011215613b675781356144db81613d08565b926144e96040519485613954565b81845260208085019260051b820101928311613b6757602001905b8282106145115750505090565b6020809161451e84613b8f565b815201910190614504565b63ffffffff168061454c575b506007549061ffff63ffffffff83169260201c1690565b90614555613e90565b60005b815180821015614670577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111614641578114801561462a575b6145a157600101614558565b925050600954821015614293578160031c7f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0154600a548310156142935761ffff90600a60005260f063ffffffff8560041c7fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801549260e08760051b161c169460041b161c1690565b5061ffff6146388284614499565b51168410614595565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b505050905038614535565b80548210156142935760005260206000200190600090565b600082815260018201602052604090205461471d57805490680100000000000000008210156138d157826147066146d184600180960185558461467b565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b906001820191816000528260205260406000205480151560001461487e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614641578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161464157818103614847575b50505080548015614818577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906147d9828261467b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6148676148576146d1938661467b565b90549060031b1c9283928661467b565b9055600052836020526040600020553880806147a1565b50505050600090565b91929015614902575081511561489b575090565b3b156148a45790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156149155750805190602001fd5b612e9a906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906139cf56fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
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

func (_CCTPVerifier *CCTPVerifierTransactor) SetDomains(opts *bind.TransactOpts, domains []CCTPVerifierDomainUpdate) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "setDomains", domains)
}

func (_CCTPVerifier *CCTPVerifierSession) SetDomains(domains []CCTPVerifierDomainUpdate) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetDomains(&_CCTPVerifier.TransactOpts, domains)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) SetDomains(domains []CCTPVerifierDomainUpdate) (*types.Transaction, error) {
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
	Domains []CCTPVerifierDomainUpdate
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
	return common.HexToHash("0xcbd7889dd51c92285faabb0cffe5578ccef74f7579f6ceec9f155dd054e46214")
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

	SetDomains(opts *bind.TransactOpts, domains []CCTPVerifierDomainUpdate) (*types.Transaction, error)

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
