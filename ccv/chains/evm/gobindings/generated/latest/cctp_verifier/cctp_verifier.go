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

var CCTPVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CCTP_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIGNATURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalityConfig\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.FinalityConfig\",\"components\":[{\"name\":\"defaultCCTPFinalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultCCTPFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customCCIPFinalities\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"customCCTPFinalityBps\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomFinalitiesMustBeStrictlyIncreasing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomFinalityArraysMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDomainUpdate\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DomainUpdate\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MissingCustomFinalities\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101208060405234610b71576158d8803803809161001d8285610e8d565b83398101908082039060e08212610b71578051906001600160a01b038216808303610b71576020820151906001600160a01b038216808303610b71576040840151936001600160a01b03851692838603610b715760608201516001600160401b038111610b715782019289601f85011215610b715783519361009e85610eb0565b946100ac6040519687610e8d565b8086528b60208284010111610b71576100cb9160208088019101610ecb565b60808301516001600160401b038111610b715783019660a0888c0312610b71576040519860a08a016001600160401b0381118b8210176109135760405261011189610eee565b8a5261011f60208a01610eee565b60208b0190815260408a01519098906001600160401b038111610b71578d610148918c01610efd565b60408c0190815260608b0151909d906001600160401b038111610b715781610171918d01610efd565b60608d0190815260808c0151909b90916001600160401b038311610b715760409261019c9201610efd565b60808d019081529c609f190112610b715760408051969087016001600160401b03811188821017610913576101e59160c0916040526101dd60a08201610f6c565b895201610f6c565b96602087019788523315610e7c57600180546001600160a01b0319163317905560405160035490919060008361021a83611015565b80825291600184168015610e5e57600114610dfe575b61023c92500384610e8d565b8151906001600160401b0382116109135761025690611015565b601f8111610da4575b506020601f8211600114610d27576102b892826000805160206158b883398151915295936102c693600091610d1c575b508160011b916000199060031b1c1916176003555b60405193849360408552604085019061104f565b90838203602085015261104f565b0390a184158015610d14575b8015610d0c575b610ae457604051639cdbb18160e01b8152602081600481895afa8015610bc65763ffffffff91600091610ced575b501660018103610cd45750604051632c12192160e01b8152602081600481895afa908115610bc657600091610c9a575b5060405163054fd4d560e41b81526001600160a01b03919091169290602081600481875afa8015610bc65763ffffffff91600091610c7b575b501660018103610c6257506040516367e0ed8360e11b8152602081600481895afa908115610bc657600091610c19575b506001600160a01b0316838103610c01575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa908115610bc657600091610bd2575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529590602081604481855afa908115610bc657600091610b94575b506000198101809111610b7e576104da9060405190602082019863095ea7b360e01b8a526024830152604482015260448152610468606482610e8d565b6000806040998a519361047b8c86610e8d565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af13d15610b76573d906104bc82610eb0565b916104c98b519384610e8d565b82523d6000602084013e5b84611074565b805180610af5575b5050916080917fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c9363ffffffff61010051169188519384526020840152878301526060820152a180516001600160a01b031615610ae45751600580546001600160a01b03199081166001600160a01b03938416908117909255835160068054909216908416179055835190815291511660208201527f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b908290a185515115610ad357855151835151811490811591610ac6575b50610ab557600093845b875180518710156106115761ffff6105dc8863ffffffff93610f9c565b51169116101561060057600161ffff6105f6878a51610f9c565b51169501946105bf565b633fd3668360e21b60005260046000fd5b505085848861ffff8451166007549063ffff0000885160101b169163ffffffff19161717600755805180519060018060401b038211610913576801000000000000000082116109135760209060085483600855808410610a7f575b500190600860005260206000208160041c9160005b838110610a3e5750600f1981169003806109ef575b505083518051925090506001600160401b0382116109135768010000000000000000821161091357602090600954836009558084106109b9575b500190600960005260206000208160041c9160005b8381106109785750600f198116900380610929575b50508451805198925090506001600160401b0388116109135768010000000000000000881161091357602090600a5489600a55898181106108bf575b50500196600a60005260206000208160041c9160005b8381106108755750600f19811690038061081b575b887fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc02675461ffff8a6107e38b6107d08c6107bd8d878e8b51998a9960208b52511660208a015251168a8801525160a0606088015260c0870190610fdd565b9051858203601f19016080870152610fdd565b9051838203601f190160a0850152610fdd565b0390a1516147a69081611112823960805181505060a051816123c8015260c05181610f06015260e051816124d1015261010051815050f35b9860009960005b8181106108425750505001969096559394859490826107bd61ffff610761565b90919a602061086b60019261ffff8f51169085851b61ffff809160031b9316831b921b19161790565b9c01929101610822565b6000805b6010811061088e57508382015560010161074c565b9b9060206108b68e60019361ffff86511691851b61ffff809160031b9316831b921b19161790565b92019c01610879565b6108f091600a600052600f8560002091601e82850160041c84019460011b16806108f7575b500160041c0190610fc6565b8989610736565b600019850190815490600019908a0360031b1c1690558e6108e4565b634e487b7160e01b600052604160045260246000fd5b9260009360005b818110610945575050500155868080806106fa565b909194602061096e60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b9601929101610930565b6000805b601081106109915750838201556001016106e5565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161097c565b6109e990600960005283600020600f80870160041c820192601e8860011b16806108f757500160041c0190610fc6565b896106d0565b9260009360005b818110610a0b57505050015586808080610696565b9091946020610a3460019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b96019291016109f6565b6000805b60108110610a57575083820155600101610681565b865190969160019160209161ffff60048b901b81811b199092169216901b1792019601610a42565b610aaf90600860005283600020600f80870160041c820192601e8860011b16806108f757500160041c0190610fc6565b8961066c565b630cacb92960e41b60005260046000fd5b90508551511415386105b5565b636c33513960e11b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b8160209181010312610b715760200151801590811503610b7157610b1a5738806104e2565b855162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b6060906104d4565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610bbe575b81610baf60209383610e8d565b81010312610b7157513861042b565b3d9150610ba2565b6040513d6000823e3d90fd5b610bf4915060203d602011610bfa575b610bec8183610e8d565b810190610f80565b386103e1565b503d610be2565b836383395ca960e01b60005260045260245260446000fd5b6020813d602011610c5a575b81610c3260209383610e8d565b81010312610c565751906001600160a01b0382168203610c535750386103a0565b80fd5b5080fd5b3d9150610c25565b6331b6aa1b60e11b600052600160045260245260446000fd5b610c94915060203d602011610bfa57610bec8183610e8d565b38610370565b90506020813d602011610ccc575b81610cb560209383610e8d565b81010312610b7157610cc690610f6c565b38610337565b3d9150610ca8565b633785f8f160e01b600052600160045260245260446000fd5b610d06915060203d602011610bfa57610bec8183610e8d565b38610307565b5087156102d9565b5083156102d2565b90508201513861028f565b601f198216906003600052806000209160005b818110610d8c5750836102c6936102b896936000805160206158b8833981519152989660019410610d73575b5050811b016003556102a4565b84015160001960f88460031b161c191690553880610d66565b91926020600181928689015181550194019201610d3a565b6003600052610dee907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610df4575b601f0160051c0190610fc6565b3861025f565b9091508190610de1565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310610e4257505090602061023c92820101610230565b6020919350806001915483858a01015201910190918592610e2a565b5050602061023c9260ff19851682840152151560051b820101610230565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761091357604052565b6001600160401b03811161091357601f01601f191660200190565b60005b838110610ede5750506000910152565b8181015183820152602001610ece565b519061ffff82168203610b7157565b9080601f83011215610b71578151916001600160401b038311610913578260051b9060405193610f306020840186610e8d565b8452602080850192820101928311610b7157602001905b828210610f545750505090565b60208091610f6184610eee565b815201910190610f47565b51906001600160a01b0382168203610b7157565b90816020910312610b71575163ffffffff81168103610b715790565b8051821015610fb05760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b818110610fd1575050565b60008155600101610fc6565b906020808351928381520192019060005b818110610ffb5750505090565b825161ffff16845260209384019390920191600101610fee565b90600182811c92168015611045575b602083101461102f57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611024565b9060209161106881518092818552858086019101610ecb565b601f01601f1916010190565b919290156110d65750815115611088575090565b3b156110915790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156110e95750805190602001fd5b60405162461bcd60e51b81526020600482015290819061110d90602483019061104f565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146137fb57508063181f5a771461377c57806331b4ec5f14612fbd57806338a9eb2a146128f75780633bbbed4b1461217c5780634014d26514612141578063445b89d114611e67578063540bc5ea14611e2d5780635cb80c5d14611b4c5780636def4ce714611a775780637437ff9f146119b857806379ba5097146118d357806380485e2514611647578063869b7f62146114c55780638da5cb5b14611473578063b2bd751c1461112d578063bff0ec1d14610abb578063c9b146b314610687578063ceac5cee1461039d578063dfadfa35146102b8578063f2fde38b146101cb578063fe163eed146101725763fec888af1461011e57600080fd5b3461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5761016b61015761401c565b6040519182916020835260208301906139b7565b0390f35b80fd5b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5773ffffffffffffffffffffffffffffffffffffffff610218613bbc565b6102206140f3565b1633811461029057807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f57604060a09167ffffffffffffffff6102fe613c90565b826080855161030c816138b9565b8281528260208201528287820152826060820152015216815260046020522060405190610338826138b9565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff81116106835736602382011215610683576103fe903690602481600401359101613d06565b6104066140f3565b61040e61401c565b90805167ffffffffffffffff81116106565761042b600354613fc9565b601f81116105bf575b506020601f82116001146104da576104bb92827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd895936104c99388916104cf575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b6040519384936040855260408501906139b7565b9083820360208501526139b7565b0390a180f35b905082015138610475565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b8181106105a75750836104c9936104bb96937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610570575b5050811b016003556104a7565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610563565b91926020600181928689015181550194019201610525565b6106289060038652601f830160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01906020841061062e575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190614152565b38610434565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b91506105fa565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5080fd5b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff8111610683576106d7903690600401613c5f565b73ffffffffffffffffffffffffffffffffffffffff6001541633141580610a99575b610a7157919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b84811015610a6d578060051b82013583811215610a6957820191608083360312610a6957604051946080860186811067ffffffffffffffff821117610a3c5760405261077384613ca7565b865261078160208501613fa4565b9660208701978852604085013567ffffffffffffffff8111610a38576107aa9036908701614169565b9460408801958652606081013567ffffffffffffffff8111610a34576107d291369101614169565b946060880195865267ffffffffffffffff88511683526002602052604083209851151561084c818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b81515161090c575b5095976001019550815b8551805182101561089d579061089673ffffffffffffffffffffffffffffffffffffffff61088e8360019561413e565b51168961456a565b500161085e565b505095909694506001929193519081516108bd575b505001939293610728565b61090267ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613cbc565b0390a238806108b2565b989395929691909497986000146109fd57600184019591875b865180518210156109a25761094f8273ffffffffffffffffffffffffffffffffffffffff9261413e565b5116801561096b57906109646001928a6144d9565b5001610925565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816109f367ffffffffffffffff8a51169251604051918291602083526020830190613cbc565b0390a23880610854565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b8280fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8480fd5b8380f35b6004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff600654163314156106f9565b503461016f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff8111610683576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261068357604051906101c0820182811067ffffffffffffffff82111761065657604052610b5581600401613ca7565b8252610b6360248201613ca7565b6020830152610b7460448201613ca7565b6040830152610b8560648201613d58565b6060830152610b9660848201613d58565b6080830152610ba760a48201613a2c565b60a083015260c481013560c083015260e481013567ffffffffffffffff8111610a3457610bda9060043691840101613d3d565b60e083015261010481013567ffffffffffffffff8111610a3457610c049060043691840101613d3d565b61010083015261012481013567ffffffffffffffff8111610a3457610c2f9060043691840101613d3d565b61012083015261014481013567ffffffffffffffff8111610a3457610c5a9060043691840101613d3d565b61014083015261016481013567ffffffffffffffff8111610a3457610c859060043691840101613d3d565b61016083015261018481013567ffffffffffffffff8111610a3457810136602382011215610a3457600481013590610cbc82613a3b565b91610cca604051938461393c565b80835260051b810160240160208301368211610fcc5760248301905b8282106110f757505050506101808301526101a48101359067ffffffffffffffff8211610a34576004610d1c9236920101613d3d565b6101a082015260243560443567ffffffffffffffff8111610a3457610d45903690600401613c00565b6101e181106110cf5780600411610a69577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361108057508061018011610a695760049360009061017c8401357fffffffff00000000000000000000000000000000000000000000000000000000167f71e2e5630000000000000000000000000000000000000000000000000000000081016110325750826101a011610fcc576020946101808501359081810361100557505067ffffffffffffffff8151168752858552604087209060ff6003830154871c1615610fd057508261011c11610fcc5760fc8401359060010154808203610f9f57505061019c92610eec859383889450507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe606101a08401910160405196879586957f57ecfd28000000000000000000000000000000000000000000000000000000008752604082880152826044880152016064860137896102008501526102048401916102006024860152613f65565b03818773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1918215610f94578492610f67575b505015610f41575080f35b907fbc40f556000000000000000000000000000000000000000000000000000000008152fd5b610f869250803d10610f8d575b610f7e818361393c565b810190613fb1565b3880610f36565b503d610f74565b6040513d86823e3d90fd5b7f7d8b101a0000000000000000000000000000000000000000000000000000000088528652602452604486fd5b8680fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000885267ffffffffffffffff168652602487fd5b7f6c86fa3a0000000000000000000000000000000000000000000000000000000089528752602452604487fd5b7fadaf77390000000000000000000000000000000000000000000000000000000088527f8e1d1a9d000000000000000000000000000000000000000000000000000000008752602452604487fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b6004857fbba6473c000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff81116111295760209161111e839283600436928a010101613d69565b815201910190610ce6565b8880fd5b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff81116106835761117d903690600401613c2e565b6111856140f3565b61118e81613a3b565b9161119c604051938461393c565b81835260c0602084019202810190368211610a6957915b8183106113d6578480855b80518310156113d2576111d1838261413e565b519267ffffffffffffffff60206111e8838561413e565b510151169384156113a6578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff8151161561137a5773ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a20191906111be565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c083360312610a6957604051906113ed82613904565b83359073ffffffffffffffffffffffffffffffffffffffff82168203610fcc578260209260c09452611420838701613ca7565b8382015261143060408701613fa4565b604082015261144160608701613a2c565b606082015261145260808701613d58565b608082015261146360a08701613d58565b60a08201528152019201916111b3565b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461016f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760405161150181613920565b611509613bbc565b815260243573ffffffffffffffffffffffffffffffffffffffff81168103610a38576020820190815261153a6140f3565b73ffffffffffffffffffffffffffffffffffffffff8251161561161f578173ffffffffffffffffffffffffffffffffffffffff6104c992817f437ab4d204b228a666e183638a005a86a5813dba72e1d299ff89c505cc52ac0b9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b503461016f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5761167f613c90565b60243567ffffffffffffffff8111610a385760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610a3857604051906116ca826138b9565b806004013567ffffffffffffffff8111610a69576116ee9060043691840101613d3d565b8252602481013567ffffffffffffffff8111610a69576117149060043691840101613d3d565b6020830152604481013567ffffffffffffffff8111610a6957810136602382011215610a6957600481013561174881613a3b565b91611756604051938461393c565b818352602060048185019360061b83010101903682116118cf57602401915b81831061189757505050604083015261179060648201613bdf565b6060830152608481013567ffffffffffffffff8111610a695760809160046117bb9236920101613d3d565b9101526044359067ffffffffffffffff8211610a38576117e867ffffffffffffffff923690600401613d3d565b506117f1613a16565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff6040822054161561186c57816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b6040833603126118cf57602060409182516118b181613920565b6118ba86613bdf565b81528286013583820152815201920191611775565b8780fd5b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f57805473ffffffffffffffffffffffffffffffffffffffff81163303611990577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760206040516119f581613920565b828152015261016b604051611a0981613920565b73ffffffffffffffffffffffffffffffffffffffff60055416815273ffffffffffffffffffffffffffffffffffffffff60065416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5767ffffffffffffffff611ab8613c90565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b818110611b365773ffffffffffffffffffffffffffffffffffffffff8561016b88611b0e8189038261393c565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613cbc565b8254845260209093019260019283019201611ae1565b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff811161068357611b9c903690600401613c5f565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b83811015611e29578060051b8201359073ffffffffffffffffffffffffffffffffffffffff8216809203611e2557604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315611e1a578793611de7575b5082611c3d575b506001915001611bb9565b8460405193611cf860208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611c8860648261393c565b8a80604098895193611c9a8b8661393c565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d15611ddf573d90611cdb8261397d565b91611ce88a51938461393c565b82523d8d602084013e5b866146cd565b805180611d34575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338611c32565b611d4b929495969350602080918301019101613fb1565b15611d5c5792919085903880611d00565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090611cf2565b9092506020813d8211611e12575b81611e026020938361393c565b81010312610fcc57519138611c2b565b3d9150611df5565b6040513d89823e3d90fd5b8580fd5b8480f35b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f57602060405160418152f35b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff811161068357611eb7903690600401613c2e565b90611ec06140f3565b825b828110156120885760c08102820160c081360312610a695760405190611ee782613904565b80359182815260208201356020820191818352604081019060408501358252611f1260608601613d58565b9260608201938452611f2660808701613ca7565b611f3960a06080850198838a5201613fa4565b9760a084019889521591821561207f575b50811561206c575b50612006579263ffffffff9260039267ffffffffffffffff8560019a9998975194519251935116975115159660405194611f8b866138b9565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b1692161717905501611ec2565b8567ffffffffffffffff8663ffffffff8660c49689604051977f113b1fc2000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff9150161538611f52565b15915038611f4a565b506040519180602084016020855252604083019190845b8181106120d057857fcbd7889dd51c92285faabb0cffe5578ccef74f7579f6ceec9f155dd054e4621486860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015263ffffffff61210260608901613d58565b16606082015267ffffffffffffffff61211d60808901613ca7565b16608082015261212f60a08801613fa4565b151560a082015201940192910161209f565b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f57602060405161019c8152f35b503461016f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f576004359067ffffffffffffffff821161016f5781600401906101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc843603011261016f576121f9613b99565b5060843567ffffffffffffffff81116106835761221a903690600401613c00565b505061222a610124840183613e43565b90357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811691601481106128c2575b505060601c92602481019367ffffffffffffffff61227686613e94565b1680845260026020526040842090815490604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff86165afa80156128b7578690612854575b73ffffffffffffffffffffffffffffffffffffffff91501633036128285760f01c60ff166127e0575b505067ffffffffffffffff61231c85613e94565b1682526004602052604082209060038201549460ff8660201c16156127a257506101848101600161234d8287613ea9565b90500361276c5761237161236c6123648388613ea9565b369291613efd565b613d69565b906040820151602081519101517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612737575b505060601c9573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169687810361270c575060206124136124096124038585613ea9565b90613efd565b6080810190613e43565b9050036126bb5750506002830154801561269b5760a490915b519201359261ffff8416809403610a695754926040519261244c84613904565b8352602083019182526040830190602435825260608401948552612483608085019180835263ffffffff60a087019a168a526141ce565b9290919215612669575061ffff855191169081810291818304149015171561263c57612710900463ffffffff81116126115763ffffffff73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016955199511693519551925192604051937f8e1d1a9d00000000000000000000000000000000000000000000000000000000602086015260248501526024845261253b60448561393c565b853b156118cf579263ffffffff88999a93816125b8948b976040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c015260248b015260448a0152606489015260848801521660a48601521660c484015261010060e48401526101048301906139b7565b03925af1918215612604578161016b936125f4575b5050604051906125de60208361393c565b81526040519182916020835260208301906139b7565b6125fd9161393c565b38816125cd565b50604051903d90823e3d90fd5b7fb6f15d0f000000000000000000000000000000000000000000000000000000008752600452602486fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b517f33eae2c200000000000000000000000000000000000000000000000000000000885263ffffffff16600452602487fd5b506080810151602081805181010312610a6957602060a49101519161242c565b6124036126cb9261240992613ea9565b6127086040519283927fa3c8cf09000000000000000000000000000000000000000000000000000000008452602060048501526024840191613f65565b0390fd5b7f961c9a4f000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b161638806123ab565b8361277960249287613ea9565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8367ffffffffffffffff6127b7602493613e94565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b600082815260029091016020526040902054156127fd5780612308565b7fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d6020116128af575b8161286e6020938361393c565b81010312611e25575173ffffffffffffffffffffffffffffffffffffffff81168103611e255773ffffffffffffffffffffffffffffffffffffffff906122df565b3d9150612861565b6040513d88823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612259565b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760606080604051612936816138b9565b83815283602082015282604082015282808201520152604051612958816138b9565b61ffff600754818116835260101c166020820152604051600854808252816020810160088652602086209286905b80600f830110612ed557612a1c945491818110612cf6575b818110612cde575b818110612cc7575b818110612caf575b818110612c97575b818110612c7f575b818110612c67575b818110612c4f575b818110612c37575b818110612c1f575b818110612c07575b818110612bef575b818110612bd7575b818110612bbf575b818110612ba7575b10612b99575b50038261393c565b6040820152604051600954808252816020810160098652602086209286905b80600f830110612ded57612ac1945491818110612cf657818110612cde57818110612cc757818110612caf57818110612c9757818110612c7f57818110612c6757818110612c4f57818110612c3757818110612c1f57818110612c0757818110612bef57818110612bd757818110612bbf57818110612ba75710612b995750038261393c565b6060820152604051600a80548083529084526020820193907fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a85b81600f840110612d0b5794612b8892849261016b975491818110612cf657818110612cde57818110612cc757818110612caf57818110612c9757818110612c7f57818110612c6757818110612c4f57818110612c3757818110612c1f57818110612c0757818110612bef57818110612bd757818110612bbf57818110612ba75710612b995750038261393c565b608082015260405191829182613af0565b60f01c815260200138612a14565b92602060019161ffff8560e01c168152019301612a0e565b92602060019161ffff8560d01c168152019301612a06565b92602060019161ffff8560c01c1681520193016129fe565b92602060019161ffff8560b01c1681520193016129f6565b92602060019161ffff8560a01c1681520193016129ee565b92602060019161ffff8560901c1681520193016129e6565b92602060019161ffff8560801c1681520193016129de565b92602060019161ffff8560701c1681520193016129d6565b92602060019161ffff8560601c1681520193016129ce565b92602060019161ffff8560501c1681520193016129c6565b92602060019161ffff8560401c1681520193016129be565b92602060019161ffff8560301c1681520193016129b6565b92602060019161ffff85831c1681520193016129ae565b92602060019161ffff8560101c1681520193016129a6565b92602060019161ffff8516815201930161299e565b946001610200601092885461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019601920191612afb565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019401920184929391612a3b565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e0820152019401920184929391612986565b503461016f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f5760043567ffffffffffffffff81116106835760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126106835760405191613038836138b9565b61304482600401613a2c565b835261305260248301613a2c565b9260208101938452604483013567ffffffffffffffff8111610a385761307e9060043691860101613a53565b9360408201948552606484013567ffffffffffffffff8111610a34576130aa9060043691870101613a53565b9360608301948552608481013567ffffffffffffffff8111610a69576130d591369101600401613a53565b91608081019283526130e56140f3565b8551511561375457855151855151811490811591613747575b5061371f578392835b875180518610156131705761ffff6131248763ffffffff9361413e565b51169116101561314857600161ffff61313e868a5161413e565b5116940193613107565b6004857fff4d9a0c000000000000000000000000000000000000000000000000000000008152fd5b50508587869461ffff8551167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600754935160101b16921617176007555180519067ffffffffffffffff82116136f2576801000000000000000082116136f25760209060085483600855808410613667575b50019060088652602086208160041c91875b83811061362757507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff081169003806135da575b505050505180519067ffffffffffffffff82116135ad576801000000000000000082116135ad5760209060095483600955808410613522575b50019060098552602085208160041c91865b8381106134e257507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff08116900380613495575b505050505180519067ffffffffffffffff82116106565768010000000000000000821161065657602090600a5483600a55808410613409575b500190600a8452602084208160041c91855b8381106133c957507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0811690038061335a575b857fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc0267546104c98760405191829182613af0565b928593865b8181106133965750505001556104c97fbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc0267548480613328565b90919460206133bf60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b960192910161335f565b86875b601081106133e15750838201556001016132f5565b865190969160019160209161ffff60048b901b81811b199092169216901b17920196016133cc565b61343890600a8752838720600f80870160041c820192601e8860011b168061343e575b500160041c0190614152565b856132e3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558a61342c565b928693875b8181106134af575050500155838080806132aa565b90919460206134d860019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b960192910161349a565b87885b601081106134fa575083820155600101613277565b865190969160019160209161ffff60048b901b81811b199092169216901b17920196016134e5565b6135509060098852838820600f80870160041c820192601e8860011b168061355657500160041c0190614152565b86613265565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558b61342c565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b928793885b8181106135f45750505001558480808061322c565b909194602061361d60019261ffff8951169085851b61ffff809160031b9316831b921b19161790565b96019291016135df565b88895b6010811061363f5750838201556001016131f9565b865190969160019160209161ffff60048b901b81811b199092169216901b179201960161362a565b6136959060088952838920600f80870160041c820192601e8860011b168061369b57500160041c0190614152565b876131e7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8254918a0360031b1c1690558c61342c565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6004847fcacb9290000000000000000000000000000000000000000000000000000000008152fd5b90508351511415386130fe565b6004847fd866a272000000000000000000000000000000000000000000000000000000008152fd5b503461016f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016f575061016b6040516137bd60408261393c565b601681527f43435450566572696669657220312e372e302d6465760000000000000000000060208201526040519182916020835260208301906139b7565b9050346106835760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610683576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610a3857602092507ffacbd7dc00000000000000000000000000000000000000000000000000000000811490811561388f575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613888565b60a0810190811067ffffffffffffffff8211176138d557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176138d557604052565b6040810190811067ffffffffffffffff8211176138d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176138d557604052565b67ffffffffffffffff81116138d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613a015750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016139c2565b6064359061ffff82168203613a2757565b600080fd5b359061ffff82168203613a2757565b67ffffffffffffffff81116138d55760051b60200190565b9080601f83011215613a27578135613a6a81613a3b565b92613a78604051948561393c565b81845260208085019260051b820101928311613a2757602001905b828210613aa05750505090565b60208091613aad84613a2c565b815201910190613a93565b906020808351928381520192019060005b818110613ad65750505090565b825161ffff16845260209384019390920191600101613ac9565b90613b96916020815261ffff825116602082015261ffff60208301511660408201526080613b63613b30604085015160a0606086015260c0850190613ab8565b60608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152613ab8565b9201519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613ab8565b90565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203613a2757565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613a2757565b359073ffffffffffffffffffffffffffffffffffffffff82168203613a2757565b9181601f84011215613a275782359167ffffffffffffffff8311613a275760208381860195010111613a2757565b9181601f84011215613a275782359167ffffffffffffffff8311613a275760208085019460c08502010111613a2757565b9181601f84011215613a275782359167ffffffffffffffff8311613a27576020808501948460051b010111613a2757565b6004359067ffffffffffffffff82168203613a2757565b359067ffffffffffffffff82168203613a2757565b906020808351928381520192019060005b818110613cda5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613ccd565b929192613d128261397d565b91613d20604051938461393c565b829481845281830111613a27578281602093846000960137010152565b9080601f83011215613a2757816020613b9693359101613d06565b359063ffffffff82168203613a2757565b919060c083820312613a275760405190613d8282613904565b819380358352602081013567ffffffffffffffff8111613a275782613da8918301613d3d565b6020840152604081013567ffffffffffffffff8111613a275782613dcd918301613d3d565b6040840152606081013567ffffffffffffffff8111613a275782613df2918301613d3d565b6060840152608081013567ffffffffffffffff8111613a275782613e17918301613d3d565b608084015260a08101359167ffffffffffffffff8311613a275760a092613e3e9201613d3d565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613a27570180359067ffffffffffffffff8211613a2757602001918136038313613a2757565b3567ffffffffffffffff81168103613a275790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613a27570180359067ffffffffffffffff8211613a2757602001918160051b36038313613a2757565b9015613f36578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215613a27570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b35908115158203613a2757565b90816020910312613a2757518015158103613a275790565b90600182811c92168015614012575b6020831014613fe357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613fd8565b604051906000826003549161403083613fc9565b80835292600181169081156140b65750600114614056575b6140549250038361393c565b565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b81831061409a57505090602061405492820101614048565b6020919350806001915483858901015201910190918492614082565b602092506140549491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101614048565b73ffffffffffffffffffffffffffffffffffffffff60015416330361411457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015613f365760209160051b010190565b81811061415d575050565b60008155600101614152565b9080601f83011215613a2757813561418081613a3b565b9261418e604051948561393c565b81845260208085019260051b820101928311613a2757602001905b8282106141b65750505090565b602080916141c384613bdf565b8152019101906141a9565b63ffffffff16806141ef57506007549061ffff8083169260101c1690600190565b90604051600854808252816020810160086000526020600020926000905b80600f8301106143d957614293945491818110612cf657818110612cde57818110612cc757818110612caf57818110612c9757818110612c7f57818110612c6757818110612c4f57818110612c3757818110612c1f57818110612c0757818110612bef57818110612bd757818110612bbf57818110612ba75710612b995750038261393c565b60005b8151808210156143c9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161439a5781148015614383575b6142df57600101614296565b92505060006009548310156143565780600960209252208260041c809101601e8460011b1690549160f08560041b1694600090600a54111561435657600a90527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8015461ffff92851c8316941c9091169160019150565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b5061ffff614391828461413e565b511684106142d3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050509050600090600090600090565b916010919350610200600191865461ffff8116825261ffff81861c16602083015261ffff8160201c16604083015261ffff8160301c16606083015261ffff8160401c16608083015261ffff8160501c1660a083015261ffff8160601c1660c083015261ffff8160701c1660e083015261ffff8160801c1661010083015261ffff8160901c1661012083015261ffff8160a01c1661014083015261ffff8160b01c1661016083015261ffff8160c01c1661018083015261ffff8160d01c166101a083015261ffff8160e01c166101c083015260f01c6101e082015201940192018492939161420d565b8054821015613f365760005260206000200190600090565b600082815260018201602052604090205461456357805490680100000000000000008210156138d5578261454c6145178460018096018555846144c1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b90600182019181600052826020526040600020548015156000146146c4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161439a578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161439a5781810361468d575b5050508054801561465e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061461f82826144c1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6146ad61469d61451793866144c1565b90549060031b1c928392866144c1565b9055600052836020526040600020553880806145e7565b50505050600090565b9192901561474857508151156146e1575090565b3b156146ea5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561475b5750805190602001fd5b612708906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906139b756fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
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

func (_CCTPVerifier *CCTPVerifierCaller) CCTPMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "CCTP_MESSAGE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) CCTPMESSAGELENGTH() (*big.Int, error) {
	return _CCTPVerifier.Contract.CCTPMESSAGELENGTH(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) CCTPMESSAGELENGTH() (*big.Int, error) {
	return _CCTPVerifier.Contract.CCTPMESSAGELENGTH(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) SIGNATURELENGTH() (*big.Int, error) {
	return _CCTPVerifier.Contract.SIGNATURELENGTH(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _CCTPVerifier.Contract.SIGNATURELENGTH(&_CCTPVerifier.CallOpts)
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
	return common.HexToHash("0xbcf028e2a6415ef5958bbdc6235e93b9f77415141c54566dd764574fcc026754")
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
	CCTPMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error)

	SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (CCTPVerifierDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCTPVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetFinalityConfig(opts *bind.CallOpts) (CCTPVerifierFinalityConfig, error)

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
