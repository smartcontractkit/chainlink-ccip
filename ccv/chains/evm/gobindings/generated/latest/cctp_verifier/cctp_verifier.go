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

type BaseVerifierRemoteChainConfigArgs struct {
	Router              common.Address
	RemoteChainSelector uint64
	AllowlistEnabled    bool
	FeeUSDCents         uint16
	GasForVerification  uint32
	PayloadSizeBytes    uint32
}

type CCTPVerifierDomain struct {
	AllowedCallerOnDest   [32]byte
	AllowedCallerOnSource [32]byte
	MintRecipientOnDest   [32]byte
	DomainIdentifier      uint32
	Enabled               bool
}

type CCTPVerifierDynamicConfig struct {
	FeeAggregator   common.Address
	AllowlistAdmin  common.Address
	FastFinalityBps uint16
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
	MessageNumber       uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifierArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.StaticConfig\",\"components\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocations\",\"inputs\":[{\"name\":\"newLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidFastFinalityBps\",\"inputs\":[{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierArgsLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461065257614ab1803803809161001d8285610b40565b833981018181039161010083126106525780516001600160a01b038116908181036106525760208301516001600160a01b038116908181036106525760408501516001600160a01b038116979093908885036106525760608701516001600160401b03811161065257870188601f820112156106525780519061009f82610b63565b996100ad6040519b8c610b40565b828b526020808c019360051b830101918183116106525760208101935b838510610ad7575050505050606090607f1901126106525760405194606086016001600160401b0381118782101761094a5760405261010b60808801610bb8565b865261011960a08801610bb8565b976020870198895260c08801519761ffff891689036106525760e06101459160408a019a8b5201610bb8565b903315610ac657600180546001600160a01b03191633179055600354815190919061016f83610b63565b9261017d6040519485610b40565b808452600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90602086015b838210610a215750505060005b81811061098c57505060005b8181106107fd5750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161021b61020d92604051938493604085526040850190610c79565b908382036020850152610c79565b0390a16001600160a01b031680156105c557608052801580156107f5575b80156107ed575b6105c557604051639cdbb18160e01b8152602081600481855afa80156106a75763ffffffff916000916107ce575b5016600181036107b55750602060049160405192838092632c12192160e01b82525afa9081156106a75760009161077b575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa80156106a75763ffffffff9160009161075c575b50166001810361074357506020600491604051928380926367e0ed8360e11b82525afa9081156106a7576000916106fa575b506001600160a01b03168381036106e2575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa9081156106a7576000916106b3575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529490602081604481855afa9081156106a757600091610675575b50600019810180911161065f576104439160405191602083019763095ea7b360e01b895260248401526044830152604482526103d3606483610b40565b6000806040988951946103e68b87610b40565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15610657573d9161042783610b7a565b926104348a519485610b40565b83523d6000602085013e610cd2565b8051806105d6575b505060e05160c05160a0516101005187516001600160a01b039485168152928416602084015292168187015263ffffffff90911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c90608090a180516001600160a01b0316156105c55761ffff825116801580156105ba575b6105a6575051600580546001600160a01b039283166001600160a01b03199091168117909155835160068054855161ffff60a01b60a09190911b169285166001600160b01b031990911617919091179055845190815292511660208301525161ffff16818301527f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197290606090a151613d419081610d708239608051816138fa015260a0518181816125250152612fcf015260c05181818161120e0152612fa2015260e0518181816125c40152612f5101526101005181612ff60152f35b630c74dcaf60e41b60005260045260246000fd5b5061271081116104c9565b6342bcdf7f60e11b60005260046000fd5b81602091810103126106525760200151801590811503610652576105fb57388061044b565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b606091610cd2565b634e487b7160e01b600052601160045260246000fd5b90506020813d60201161069f575b8161069060209383610b40565b81010312610652575138610396565b3d9150610683565b6040513d6000823e3d90fd5b6106d5915060203d6020116106db575b6106cd8183610b40565b810190610bcc565b3861034c565b503d6106c3565b836383395ca960e01b60005260045260245260446000fd5b6020813d60201161073b575b8161071360209383610b40565b810103126107375751906001600160a01b038216820361073457503861030b565b80fd5b5080fd5b3d9150610706565b6331b6aa1b60e11b600052600160045260245260446000fd5b610775915060203d6020116106db576106cd8183610b40565b386102d9565b90506020813d6020116107ad575b8161079660209383610b40565b81010312610652576107a790610bb8565b386102a0565b3d9150610789565b633785f8f160e01b600052600160045260245260446000fd5b6107e7915060203d6020116106db576106cd8183610b40565b3861026e565b508815610240565b508315610239565b82518110156109765760208160051b840101516003546801000000000000000081101561094a578060016108349201600355610be8565b919091610960578051906001600160401b03821161094a576108568354610c03565b601f811161090d575b50602090601f83116001146108a25760019493929160009183610897575b5050600019600383901b1c191690841b1790555b016101c7565b01519050388061087d565b90601f1983169184600052816000209260005b8181106108f55750916001969594929183889593106108dc575b505050811b019055610891565b015160001960f88460031b161c191690553880806108cf565b929360206001819287860151815501950193016108b5565b61093a90846000526020600020601f850160051c81019160208610610940575b601f0160051c0190610c3d565b3861085f565b909150819061092d565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6003548015610a0b5760001901906109a382610be8565b92909261096057826109b760019454610c03565b90816109c9575b5050600355016101bb565b81601f6000931186146109e05750555b38806109be565b818352602083206109fb91601f0160051c8101908701610c3d565b80825281602081209155556109d9565b634e487b7160e01b600052603160045260246000fd5b60405160008454610a3181610c03565b8084529060018116908115610aa35750600114610a6b575b5060019282610a5d85946020940382610b40565b8152019301910190916101ae565b6000868152602081209092505b818310610a8d57505081016020016001610a49565b6001816020925483868801015201920191610a78565b60ff191660208581019190915291151560051b8401909101915060019050610a49565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116106525782019083603f8301121561065257602082015190610b0482610b7a565b610b116040519182610b40565b828152604084840101861061065257610b3560209493859460408685019101610b95565b8152019401936100ca565b601f909101601f19168101906001600160401b0382119082101761094a57604052565b6001600160401b03811161094a5760051b60200190565b6001600160401b03811161094a57601f01601f191660200190565b60005b838110610ba85750506000910152565b8181015183820152602001610b98565b51906001600160a01b038216820361065257565b90816020910312610652575163ffffffff811681036106525790565b60035481101561097657600360005260206000200190600090565b90600182811c92168015610c33575b6020831014610c1d57565b634e487b7160e01b600052602260045260246000fd5b91607f1691610c12565b818110610c48575050565b60008155600101610c3d565b90602091610c6d81518092818552858086019101610b95565b601f01601f1916010190565b9080602083519182815201916020808360051b8301019401926000915b838310610ca557505050505090565b9091929394602080610cc3600193601f198682030187528951610c54565b97019301930191939290610c96565b91929015610d345750815115610ce6575090565b3b15610cef5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d475750805190602001fd5b60405162461bcd60e51b815260206004820152908190610d6b906024830190610c54565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146130355750806306285c6914612eea578063181f5a7714612e6b5780632e45ca6814612aff5780633bbbed4b146122585780635cb80c5d14611f7b5780635ef2c64b14611b1e5780637437ff9f14611a2b57806379ba50971461194657806380485e251461167e57806387ae929214611630578063898068fc146115575780638da5cb5b14611505578063bff0ec1d14610c52578063c9b146b31461086c578063d52e545a14610586578063dfadfa35146104a1578063e023ddb114610245578063f2fde38b146101585763fe163eed146100fd57600080fd5b3461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b80fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555773ffffffffffffffffffffffffffffffffffffffff6101a56132c5565b6101ad613806565b1633811461021d57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101555760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576040516102818161315a565b6102896132c5565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361049d57602082019081526044359061ffff8216820361049957604083019182526102d0613806565b73ffffffffffffffffffffffffffffffffffffffff835116156104715761ffff82511680158015610466575b61043b5750916104359173ffffffffffffffffffffffffffffffffffffffff7f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197294818451167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655517fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000806006549360a01b161691161760065560405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390a180f35b7fc74dcaf0000000000000000000000000000000000000000000000000000000008552600452602484fd5b5061271081116102fc565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b8280fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557604060a09167ffffffffffffffff6104e76133b7565b82608085516104f581613176565b828152826020820152828782015282606082015201521681526004602052206040519061052182613176565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff8111610868576105d690369060040161326c565b6105e1929192613806565b81925b818410156107b05760c0840281019360c085360312610499576040519461060a8661313e565b8035908187526020810135906020880197828952604081019060408301358252610636606084016133ce565b93606082019480865261065e60a0610650608088016134e2565b9660808601978852016135cd565b9660a08401978852159182156107a7575b508115610794575b5061072e579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d51945192519351169751151596604051946106b086613176565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b16921617179055019291906105e4565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff9150161538610677565b1591503861066f565b6040519180602084016020855252604083019190845b8181106107f757857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff61082d606089016133ce565b16606082015263ffffffff610844608089016134e2565b16608082015261085660a088016135cd565b151560a08201520194019291016107c6565b5080fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff8111610868576108bc903690600401613337565b73ffffffffffffffffffffffffffffffffffffffff600154163303610c0a575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b84811015610c06578060051b82013583811215610c0257820191608083360312610c025760405194610938866130f3565b610941846133ce565b865261094f602085016135cd565b9660208701978852604085013567ffffffffffffffff811161049d5761097890369087016139e0565b9460408801958652606081013567ffffffffffffffff8111610499576109a0913691016139e0565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610a1a818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610ada575b5095976001019550815b85518051821015610a6b5790610a6473ffffffffffffffffffffffffffffffffffffffff610a5c83600195613851565b511689613ad6565b5001610a2c565b50509590969450600192919351908151610a8b575b505001939293610907565b610ad067ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613498565b0390a23880610a80565b98939592969190949798600014610bcb57600184019591875b86518051821015610b7057610b1d8273ffffffffffffffffffffffffffffffffffffffff92613851565b51168015610b395790610b326001928a613a45565b5001610af3565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281610bc167ffffffffffffffff8a51169251604051918291602083526020830190613498565b0390a23880610a22565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8480fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff600654163303156108dc576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b50346101555760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576004359067ffffffffffffffff8211610155576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261015557604051916101c0830183811067ffffffffffffffff8211176114d857604052610ced816004016133ce565b8352610cfb602482016133ce565b6020840152610d0c604482016133ce565b6040840152610d1d606482016134e2565b6060840152610d2e608482016134e2565b6080840152610d3f60a48201613412565b60a084015260c481013560c084015260e481013567ffffffffffffffff811161049d57610d7290600436918401016133e3565b60e084015261010481013567ffffffffffffffff811161049d57610d9c90600436918401016133e3565b61010084015261012481013567ffffffffffffffff811161049d57610dc790600436918401016133e3565b61012084015261014481013567ffffffffffffffff811161049d57610df290600436918401016133e3565b61014084015261016481013567ffffffffffffffff811161049d57610e1d90600436918401016133e3565b61016084015261018481013567ffffffffffffffff811161049d5781013660238201121561049d57600481013590610e5482613368565b91610e626040519384613192565b80835260051b8101602401602083013682116114d45760248301905b82821061149e57505050506101808401526101a481013567ffffffffffffffff811161049d57610eb3913691016004016133e3565b6101a083015260243560443567ffffffffffffffff811161049d57610edc903690600401613309565b610ef067ffffffffffffffff865116613894565b67ffffffffffffffff855116808552600260205273ffffffffffffffffffffffffffffffffffffffff604086205416908115611473576020906044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611468578591611449575b501561141d576101e181106113f55780600411610499577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d0000000000000000000000000000000000000000000000000000000081036113a657506101809261017c600083861161015557507fffffffff00000000000000000000000000000000000000000000000000000000908481013580831691870360048110611391575b5050167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361134257506000906101a09485811161049d5783861161049d57611068908087039086016137b3565b9081810361131457505067ffffffffffffffff86511685526004602052604085209560ff600388015460201c16156112de575061011c9050600082821161015557506001906110dc907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc85016137b3565b950154948581036112ae57508394508083116112a957836102047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f6020969487956004995087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6082019061019c6040519b8c9a8b997f57ecfd28000000000000000000000000000000000000000000000000000000008b526040838c0152508260448b01520160648901378b6102008801528186880161020060248a0152526102248701378986857ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608489010101015201168201010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561129e57829161126f575b50156112475780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b611291915060203d602011611297575b6112898183613192565b8101906137ee565b3861123e565b503d61127f565b6040513d84823e3d90fd5b505050fd5b84604491877f7d8b101a000000000000000000000000000000000000000000000000000000008352600452602452fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff16600452602485fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b839250829060040360031b1b16163880611019565b7fadaf77390000000000000000000000000000000000000000000000000000000085527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604484fd5b6004847f1ede477b000000000000000000000000000000000000000000000000000000008152fd5b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b611462915060203d602011611297576112898183613192565b38610f70565b6040513d87823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b813567ffffffffffffffff81116114d0576020916114c5839283600436928a0101016134f3565b815201910190610e7e565b8780fd5b8580fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555767ffffffffffffffff6115986133b7565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b81811061161a5773ffffffffffffffffffffffffffffffffffffffff85611616886115ee81890382613192565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613498565b0390f35b82548452602090930192600192830192016115c1565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555761161661166a613696565b604051918291602083526020830190613421565b50346101555760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576116b66133b7565b60243567ffffffffffffffff811161049d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261049d576040519061170182613176565b806004013567ffffffffffffffff8111610c025761172590600436918401016133e3565b8252602481013567ffffffffffffffff8111610c025761174b90600436918401016133e3565b6020830152604481013567ffffffffffffffff8111610c0257810136602382011215610c0257600481013561177f81613368565b9161178d6040519384613192565b818352602060048185019360061b83010101903682116114d057602401915b8183106118ce5750505060408301526117c7606482016132e8565b6060830152608481013567ffffffffffffffff8111610c025760809160046117f292369201016133e3565b9101526044359067ffffffffffffffff821161049d5761181f67ffffffffffffffff9236906004016133e3565b50611828613401565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff604082205416156118a357816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f4d1aff7e000000000000000000000000000000000000000000000000000000008252600452fd5b6040833603126114d0576040516040810181811067ffffffffffffffff82111761191957916020916040938452611904866132e8565b815282860135838201528152019201916117ac565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557805473ffffffffffffffffffffffffffffffffffffffff81163303611a03577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760408051611a678161315a565b8281528260208201520152611616604051611a818161315a565b73ffffffffffffffffffffffffffffffffffffffff60055416815261ffff60065473ffffffffffffffffffffffffffffffffffffffff8116602084015260a01c16604082015260405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff81116108685736602382011215610868578060040135611b7981613368565b91611b876040519384613192565b8183526024602084019260051b82010190368211610c025760248101925b828410611f36578585611bb6613806565b600354908051611bc4613696565b92845b818110611e43575050835b818110611c2557847fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586611c178661043587604051938493604085526040850190613421565b908382036020850152613421565b611c2f8184613851565b5160035468010000000000000000811015611e1657806001611c549201600355613996565b919091611dea5780519067ffffffffffffffff8211611dbd57611c778354613643565b601f8111611d82575b50602090601f8311600114611cdf57600194939291899183611cd4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82861b9260031b1c19161790555b01611bd2565b015190508980611c9d565b83895281892091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168a5b818110611d6a575091600196959492918388959310611d33575b505050811b019055611cce565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611d26565b92936020600181928786015181550195019301611d0c565b611dad90848a5260208a20601f850160051c81019160208610611db3575b601f0160051c01906139c9565b88611c80565b9091508190611da0565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6024877f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6003548015611f09577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611e7781613996565b611edd57908782611e8b6001959454613643565b80611e9d575b50505060035501611bc7565b601f81118614611eb25750555b878980611e91565b81835260208320611ecd91601f0160051c81019087016139c9565b8082528160208120915555611eaa565b6024887f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b833567ffffffffffffffff8111611f7757820136604382011215611f7757602091611f6c83923690604460248201359101613380565b815201930192611ba5565b8680fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff811161086857611fcb903690600401613337565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b83811015612254578060051b8201359073ffffffffffffffffffffffffffffffffffffffff82168092036114d457604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315612249578793612216575b508261206c575b506001915001611fe8565b846040519361212760208601957fa9059cbb000000000000000000000000000000000000000000000000000000008752836024820152826044820152604481526120b7606482613192565b8a806040988951936120c98b86613192565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d1561220e573d9061210a826131d3565b916121178a519384613192565b82523d8d602084013e5b86613c68565b805180612163575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338612061565b61217a9294959693506020809183010191016137ee565b1561218b579291908590388061212f565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090612121565b9092506020813d8211612241575b8161223160209383613192565b81010312611f775751913861205a565b3d9150612224565b6040513d89823e3d90fd5b8480f35b50346101555760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576004359067ffffffffffffffff821161015557816004018236036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261049d576122d56132a2565b5060843567ffffffffffffffff8111610499576122f6903690600401613309565b602486949294019261230f61230a856135da565b613894565b612318846135da565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd6101248901359101811215611f7757870160048101359067ffffffffffffffff82116114d057602481019180360383136127055760209082019190910312611f7757359073ffffffffffffffffffffffffffffffffffffffff8216809203611f775767ffffffffffffffff168087526002602052604087209081549073ffffffffffffffffffffffffffffffffffffffff8216908115612ad4576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612ac9578990612a66575b73ffffffffffffffffffffffffffffffffffffffff9150163303612a3a5760f01c60ff166129f2575b505067ffffffffffffffff612454846135da565b1685526004602052604085209160038301549360ff8560201c16156129b45750610184870190600161248683836135ef565b90500361297e5790612497916135ef565b156129515780357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4182360301811215611f77576124d6913691016134f3565b93604085015191602083805181010312611f775760208301519273ffffffffffffffffffffffffffffffffffffffff84168094036114d05773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016809403612913575060808601805160208151116128d55750600285015480156127ef5760a49150985b01359061ffff82168092036114d0576040519261258a8461315a565b602435845282602085015260408401916107d083528993612709575b505063ffffffff73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001697519554915116925193604051947f8e1d1a9d00000000000000000000000000000000000000000000000000000000602087015260248601526024855261262c604486613192565b873b156127055793889795929363ffffffff89956126ac94829a986040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c01521660248a01526044890152606488015260848701521660a485015260c484015261010060e484015261010483019061320d565b03925af19182156126f85781611616936126e8575b5050604051906126d2602083613192565b815260405191829160208352602083019061320d565b6126f191613192565b38816126c1565b50604051903d90823e3d90fd5b8880fd5b6103e88352919250908015612798576020810361276d5781602091810103126114d05735905b63ffffffff82116127415738806125a6565b602488837fb6f15d0f000000000000000000000000000000000000000000000000000000008252600452fd5b7f4a895e59000000000000000000000000000000000000000000000000000000008952600452602488fd5b5050855161ffff60065460a01c16908181029181830414901517156127c25761271090049061272f565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5051602081511161289357602081519101519060208110612862575b8060031b908082046008149015171561283557610100036101008111612835571c9760a49061256e565b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b169061280b565b6128d1906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840152602483019061320d565b0390fd5b6128d1906040519182917fa3c8cf0900000000000000000000000000000000000000000000000000000000835260206004840152602483019061320d565b6128d1906040519182917f22d4cfe200000000000000000000000000000000000000000000000000000000835260206004840152602483019061320d565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b61298b60249288926135ef565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8667ffffffffffffffff6129c96024936135da565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612a0f5780612440565b7fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6024887f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612ac1575b81612a8060209383613192565b81010312612705575173ffffffffffffffffffffffffffffffffffffffff811681036127055773ffffffffffffffffffffffffffffffffffffffff90612417565b3d9150612a73565b6040513d8b823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008a52600452602489fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff811161086857612b4f90369060040161326c565b612b57613806565b612b6081613368565b91612b6e6040519384613192565b81835260c0602084019202810190368211610c0257915b818310612dce578480855b8051831015612dca57612ba38382613851565b519267ffffffffffffffff6020612bba8385613851565b51015116938415612d9e578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff00000000000000000000000000000000000000001617825560808101805163ffffffff1615612d725760019495969260ff73ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040945184547fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff000000000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a086015160d01b169360b01b1691161717809455511691835192835260f01c1615156020820152a2019190612b90565b602486887f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c083360312610c025760405190612de58261313e565b83359073ffffffffffffffffffffffffffffffffffffffff82168203611f77578260209260c09452612e188387016133ce565b83820152612e28604087016135cd565b6040820152612e3960608701613412565b6060820152612e4a608087016134e2565b6080820152612e5b60a087016134e2565b60a0820152815201920191612b85565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555750611616604051612eac604082613192565b601681527f43435450566572696669657220312e372e302d64657600000000000000000000602082015260405191829160208352602083019061320d565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576060604051612f27816130f3565b8281528260208201528260408201520152608073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff604051612f82816130f3565b82815273ffffffffffffffffffffffffffffffffffffffff6020820191817f00000000000000000000000000000000000000000000000000000000000000001683528160606040830192827f00000000000000000000000000000000000000000000000000000000000000001684520193857f0000000000000000000000000000000000000000000000000000000000000000168552604051968752511660208601525116604084015251166060820152f35b9050346108685760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610868576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361049d57602092507f83adcde10000000000000000000000000000000000000000000000000000000081149081156130c9575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386130c2565b6080810190811067ffffffffffffffff82111761310f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff82111761310f57604052565b6060810190811067ffffffffffffffff82111761310f57604052565b60a0810190811067ffffffffffffffff82111761310f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761310f57604052565b67ffffffffffffffff811161310f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106132575750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613218565b9181601f8401121561329d5782359167ffffffffffffffff831161329d5760208085019460c0850201011161329d57565b600080fd5b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361329d57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361329d57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361329d57565b9181601f8401121561329d5782359167ffffffffffffffff831161329d576020838186019501011161329d57565b9181601f8401121561329d5782359167ffffffffffffffff831161329d576020808501948460051b01011161329d57565b67ffffffffffffffff811161310f5760051b60200190565b92919261338c826131d3565b9161339a6040519384613192565b82948184528183011161329d578281602093846000960137010152565b6004359067ffffffffffffffff8216820361329d57565b359067ffffffffffffffff8216820361329d57565b9080601f8301121561329d578160206133fe93359101613380565b90565b6064359061ffff8216820361329d57565b359061ffff8216820361329d57565b9080602083519182815201916020808360051b8301019401926000915b83831061344d57505050505090565b9091929394602080613489837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161320d565b9701930193019193929061343e565b906020808351928381520192019060005b8181106134b65750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016134a9565b359063ffffffff8216820361329d57565b919060c08382031261329d576040519061350c8261313e565b819380358352602081013567ffffffffffffffff811161329d57826135329183016133e3565b6020840152604081013567ffffffffffffffff811161329d57826135579183016133e3565b6040840152606081013567ffffffffffffffff811161329d578261357c9183016133e3565b6060840152608081013567ffffffffffffffff811161329d57826135a19183016133e3565b608084015260a08101359167ffffffffffffffff831161329d5760a0926135c892016133e3565b910152565b3590811515820361329d57565b3567ffffffffffffffff8116810361329d5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561329d570180359067ffffffffffffffff821161329d57602001918160051b3603831361329d57565b90600182811c9216801561368c575b602083101461365d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613652565b600354906136a382613368565b916136b16040519384613192565b808352600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9190602085015b8282106136f05750505050565b6040516000855461370081613643565b8084529060018116908115613772575060011461373a575b506001928261372c85946020940382613192565b8152019401910190926136e3565b6000878152602081209092505b81831061375c57505081016020016001613718565b6001816020925483868801015201920191613747565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050613718565b3590602081106137c1575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9081602091031261329d5751801515810361329d5790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361382757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156138655760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561398a5760009161396b575b506139335750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b613984915060203d602011611297576112898183613192565b3861392b565b6040513d6000823e3d90fd5b60035481101561386557600360005260206000200190600090565b80548210156138655760005260206000200190600090565b8181106139d4575050565b600081556001016139c9565b9080601f8301121561329d5781356139f781613368565b92613a056040519485613192565b81845260208085019260051b82010192831161329d57602001905b828210613a2d5750505090565b60208091613a3a846132e8565b815201910190613a20565b6000828152600182016020526040902054613acf578054906801000000000000000082101561310f5782613ab8613a838460018096018555846139b1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014613c5f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613c30578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613c3057818103613bf9575b50505080548015613bca577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190613b8b82826139b1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b613c19613c09613a8393866139b1565b90549060031b1c928392866139b1565b905560005283602052604060002055388080613b53565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b91929015613ce35750815115613c7c575090565b3b15613c855790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015613cf65750805190602001fd5b6128d1906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061320d56fea164736f6c634300081a000a",
}

var CCTPVerifierABI = CCTPVerifierMetaData.ABI

var CCTPVerifierBin = CCTPVerifierMetaData.Bin

func DeployCCTPVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, messageTransmitterProxy common.Address, usdcToken common.Address, storageLocations []string, dynamicConfig CCTPVerifierDynamicConfig, rmn common.Address) (common.Address, *types.Transaction, *CCTPVerifier, error) {
	parsed, err := CCTPVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPVerifierBin), backend, tokenMessenger, messageTransmitterProxy, usdcToken, storageLocations, dynamicConfig, rmn)
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

func (_CCTPVerifier *CCTPVerifierCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _CCTPVerifier.Contract.GetRemoteChainConfig(&_CCTPVerifier.CallOpts, remoteChainSelector)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _CCTPVerifier.Contract.GetRemoteChainConfig(&_CCTPVerifier.CallOpts, remoteChainSelector)
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

func (_CCTPVerifier *CCTPVerifierCaller) GetStorageLocations(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getStorageLocations")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetStorageLocations() ([]string, error) {
	return _CCTPVerifier.Contract.GetStorageLocations(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetStorageLocations() ([]string, error) {
	return _CCTPVerifier.Contract.GetStorageLocations(&_CCTPVerifier.CallOpts)
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

func (_CCTPVerifier *CCTPVerifierTransactor) ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "applyRemoteChainConfigUpdates", remoteChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyRemoteChainConfigUpdates(&_CCTPVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ApplyRemoteChainConfigUpdates(&_CCTPVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_CCTPVerifier *CCTPVerifierTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "forwardToVerifier", message, messageId, arg2, arg3, verifierArgs)
}

func (_CCTPVerifier *CCTPVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ForwardToVerifier(&_CCTPVerifier.TransactOpts, message, messageId, arg2, arg3, verifierArgs)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.ForwardToVerifier(&_CCTPVerifier.TransactOpts, message, messageId, arg2, arg3, verifierArgs)
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

func (_CCTPVerifier *CCTPVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPVerifier *CCTPVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.TransferOwnership(&_CCTPVerifier.TransactOpts, to)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.TransferOwnership(&_CCTPVerifier.TransactOpts, to)
}

func (_CCTPVerifier *CCTPVerifierTransactor) UpdateStorageLocations(opts *bind.TransactOpts, newLocations []string) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "updateStorageLocations", newLocations)
}

func (_CCTPVerifier *CCTPVerifierSession) UpdateStorageLocations(newLocations []string) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.UpdateStorageLocations(&_CCTPVerifier.TransactOpts, newLocations)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) UpdateStorageLocations(newLocations []string) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.UpdateStorageLocations(&_CCTPVerifier.TransactOpts, newLocations)
}

func (_CCTPVerifier *CCTPVerifierTransactor) VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "verifyMessage", message, messageHash, verifierResults)
}

func (_CCTPVerifier *CCTPVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.VerifyMessage(&_CCTPVerifier.TransactOpts, message, messageHash, verifierResults)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.VerifyMessage(&_CCTPVerifier.TransactOpts, message, messageHash, verifierResults)
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

type CCTPVerifierRemoteChainConfigSetIterator struct {
	Event *CCTPVerifierRemoteChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierRemoteChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierRemoteChainConfigSet)
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
		it.Event = new(CCTPVerifierRemoteChainConfigSet)
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

func (it *CCTPVerifierRemoteChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierRemoteChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierRemoteChainConfigSet struct {
	RemoteChainSelector uint64
	Router              common.Address
	AllowlistEnabled    bool
	Raw                 types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPVerifierRemoteChainConfigSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierRemoteChainConfigSetIterator{contract: _CCTPVerifier.contract, event: "RemoteChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierRemoteChainConfigSet)
				if err := _CCTPVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseRemoteChainConfigSet(log types.Log) (*CCTPVerifierRemoteChainConfigSet, error) {
	event := new(CCTPVerifierRemoteChainConfigSet)
	if err := _CCTPVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

type CCTPVerifierStorageLocationsUpdatedIterator struct {
	Event *CCTPVerifierStorageLocationsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierStorageLocationsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierStorageLocationsUpdated)
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
		it.Event = new(CCTPVerifierStorageLocationsUpdated)
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

func (it *CCTPVerifierStorageLocationsUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierStorageLocationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierStorageLocationsUpdated struct {
	OldLocations []string
	NewLocations []string
	Raw          types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*CCTPVerifierStorageLocationsUpdatedIterator, error) {

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierStorageLocationsUpdatedIterator{contract: _CCTPVerifier.contract, event: "StorageLocationsUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStorageLocationsUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierStorageLocationsUpdated)
				if err := _CCTPVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseStorageLocationsUpdated(log types.Log) (*CCTPVerifierStorageLocationsUpdated, error) {
	event := new(CCTPVerifierStorageLocationsUpdated)
	if err := _CCTPVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetFee struct {
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}
type GetRemoteChainConfig struct {
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}

func (CCTPVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CCTPVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CCTPVerifierDomainsSet) Topic() common.Hash {
	return common.HexToHash("0x4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a30046")
}

func (CCTPVerifierDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c0947341972")
}

func (CCTPVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCTPVerifierRemoteChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9")
}

func (CCTPVerifierStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c")
}

func (CCTPVerifierStorageLocationsUpdated) Topic() common.Hash {
	return common.HexToHash("0xec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586")
}

func (_CCTPVerifier *CCTPVerifier) Address() common.Address {
	return _CCTPVerifier.address
}

type CCTPVerifierInterface interface {
	GetDomain(opts *bind.CallOpts, chainSelector uint64) (CCTPVerifierDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCTPVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetStaticConfig(opts *bind.CallOpts) (CCTPVerifierStaticConfig, error)

	GetStorageLocations(opts *bind.CallOpts) ([]string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, verifierArgs []byte) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []CCTPVerifierSetDomainArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCTPVerifierDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocations(opts *bind.TransactOpts, newLocations []string) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CCTPVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CCTPVerifierAllowListSendersRemoved, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*CCTPVerifierDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*CCTPVerifierDomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPVerifierDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPVerifierDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPVerifierFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPVerifierOwnershipTransferred, error)

	FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPVerifierRemoteChainConfigSetIterator, error)

	WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigSet(log types.Log) (*CCTPVerifierRemoteChainConfigSet, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*CCTPVerifierStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*CCTPVerifierStaticConfigSet, error)

	FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*CCTPVerifierStorageLocationsUpdatedIterator, error)

	WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *CCTPVerifierStorageLocationsUpdated) (event.Subscription, error)

	ParseStorageLocationsUpdated(log types.Log) (*CCTPVerifierStorageLocationsUpdated, error)

	Address() common.Address
}
