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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifierArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocations\",\"inputs\":[{\"name\":\"newLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidFastFinalityBps\",\"inputs\":[{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierArgsLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461068d5761487f803803809161001d8285610a45565b8339810190808203610100811261068d5781516001600160a01b038116919082810361068d5760208401516001600160a01b0381169081810361068d5760408601516001600160a01b0381169490939085850361068d5760608801516001600160401b03811161068d57880189601f8201121561068d578051906100a082610a68565b9a6100ae6040519c8d610a45565b828c526020808d019360051b8301019181831161068d5760208101935b8385106109cd575050505050606090607f19011261068d5760405195606087016001600160401b038111888210176108405760405261010c60808901610aa2565b875261011a60a08901610aa2565b9860208801998a5260c08901519861ffff8a168a0361068d5760e06101469160408b019b8c5201610aa2565b9033156109bc57600180546001600160a01b03191633179055600354815190919061017083610a68565b9261017e6040519485610a45565b808452600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90602086015b8382106109175750505060005b81811061088257505060005b8181106106f35750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161021c61020e92604051938493604085526040850190610b3e565b908382036020850152610b3e565b0390a16001600160a01b031680156106d257608052801580156106eb575b80156106e3575b6106d257604051639cdbb18160e01b8152602081600481855afa80156105b65763ffffffff916000916106b3575b50166001810361069a5750602060049160405192838092632c12192160e01b82525afa9081156105b65760009161065b575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa80156105b65763ffffffff9160009161063c575b50166001810361062357506020600491604051928380926367e0ed8360e11b82525afa9081156105b6576000916105da575b506001600160a01b03168381036105c2575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa9081156105b657600091610587575b506101005260a05260018060a01b0360e0511690604051906020600081840163095ea7b360e01b81528560248601528119604486015260448552610392606486610a45565b84519082855af16000513d8261056b575b505015610525575b505060e05160c05160a05161010051604080516001600160a01b0395861681529385166020850152919093169082015263ffffffff9190911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c9150608090a161ffff8251168015801561051a575b610506578151600580546001600160a01b039283166001600160a01b03199091168117909155855160068054875161ffff60a01b60a09190911b169285166001600160b01b0319909116179190911790556040805191825286519092166020820152845161ffff16918101919091527f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197290606090a1604051613c749081610c0b8239608051816138f9015260a0518181816126a50152613126015260c05181818161124f01526130ea015260e05181818161274401526130b1015261010051816131520152f35b630c74dcaf60e41b60005260045260246000fd5b50612710811161041f565b61055e610563936040519063095ea7b360e01b602083015260248201526000604482015260448152610558606482610a45565b82610baf565b610baf565b3880806103ab565b90915061057f5750803b15155b38806103a3565b600114610578565b6105a9915060203d6020116105af575b6105a18183610a45565b810190610ab6565b3861034d565b503d610597565b6040513d6000823e3d90fd5b836383395ca960e01b60005260045260245260446000fd5b6020813d60201161061b575b816105f360209383610a45565b810103126106175751906001600160a01b038216820361061457503861030c565b80fd5b5080fd5b3d91506105e6565b6331b6aa1b60e11b600052600160045260245260446000fd5b610655915060203d6020116105af576105a18183610a45565b386102da565b90506020813d602011610692575b8161067660209383610a45565b8101031261068d5761068790610aa2565b386102a1565b600080fd5b3d9150610669565b633785f8f160e01b600052600160045260245260446000fd5b6106cc915060203d6020116105af576105a18183610a45565b3861026f565b6342bcdf7f60e11b60005260046000fd5b508515610241565b50831561023a565b825181101561086c5760208160051b84010151600354680100000000000000008110156108405780600161072a9201600355610ad2565b919091610856578051906001600160401b0382116108405761074c8354610aed565b601f8111610803575b50602090601f8311600114610798576001949392916000918361078d575b5050600019600383901b1c191690841b1790555b016101c8565b015190503880610773565b90601f1983169184600052816000209260005b8181106107eb5750916001969594929183889593106107d2575b505050811b019055610787565b015160001960f88460031b161c191690553880806107c5565b929360206001819287860151815501950193016107ab565b61083090846000526020600020601f850160051c81019160208610610836575b601f0160051c0190610b27565b38610755565b9091508190610823565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600354801561090157600019019061089982610ad2565b92909261085657826108ad60019454610aed565b90816108bf575b5050600355016101bc565b81601f6000931186146108d65750555b38806108b4565b818352602083206108f191601f0160051c8101908701610b27565b80825281602081209155556108cf565b634e487b7160e01b600052603160045260246000fd5b6040516000845461092781610aed565b80845290600181169081156109995750600114610961575b506001928261095385946020940382610a45565b8152019301910190916101af565b6000868152602081209092505b8183106109835750508101602001600161093f565b600181602092548386880101520192019161096e565b60ff191660208581019190915291151560051b840190910191506001905061093f565b639b15e16f60e01b60005260046000fd5b84516001600160401b03811161068d5782019083603f8301121561068d576020820151906001600160401b03821161084057604051610a16601f8401601f191660200182610a45565b828152604084840101861061068d57610a3a60209493859460408685019101610a7f565b8152019401936100cb565b601f909101601f19168101906001600160401b0382119082101761084057604052565b6001600160401b0381116108405760051b60200190565b60005b838110610a925750506000910152565b8181015183820152602001610a82565b51906001600160a01b038216820361068d57565b9081602091031261068d575163ffffffff8116810361068d5790565b60035481101561086c57600360005260206000200190600090565b90600182811c92168015610b1d575b6020831014610b0757565b634e487b7160e01b600052602260045260246000fd5b91607f1691610afc565b818110610b32575050565b60008155600101610b27565b9080602083519182815201916020808360051b8301019401926000915b838310610b6a57505050505090565b909192939460208080600193601f198682030187528951610b9681518092818552858086019101610a7f565b601f01601f191601019701959491909101920190610b5b565b906000602091828151910182855af1156105b6576000513d610c0157506001600160a01b0381163b155b610be05750565b635274afe760e01b60009081526001600160a01b0391909116600452602490fd5b60011415610bd956fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146131795750806306285c6914613066578063181f5a7714612fe75780632e45ca6814612c7f5780633bbbed4b146123d85780635cb80c5d146121a25780635ef2c64b14611c395780637437ff9f14611b4657806379ba509714611a6157806380485e251461179957806387ae929214611644578063898068fc1461156b5780638da5cb5b14611519578063bff0ec1d14610c93578063c9b146b314610827578063d52e545a14610541578063dfadfa351461045c578063e023ddb114610245578063f2fde38b146101585763fe163eed146100fd57600080fd5b3461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b80fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555773ffffffffffffffffffffffffffffffffffffffff6101a56133b3565b6101ad613805565b1633811461021d57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101555760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760405161028181613282565b6102896133b3565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361045857602082019081526044359061ffff8216820361045457604083019182526102d0613805565b61ffff82511680158015610449575b61041e5750916104189173ffffffffffffffffffffffffffffffffffffffff7f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197294818451167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655517fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000806006549360a01b161691161760065560405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390a180f35b7fc74dcaf0000000000000000000000000000000000000000000000000000000008552600452602484fd5b5061271081116102df565b8380fd5b8280fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557604060a09167ffffffffffffffff6104a26134d3565b82608085516104b08161329e565b82815282602082015282878201528260608201520152168152600460205220604051906104dc8261329e565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff81116108235761059190369060040161335a565b61059c929192613805565b81925b8184101561076b5760c0840281019360c08536031261045457604051946105c586613237565b80359081875260208101359060208801978289526040810190604083013582526105f1606084016134ea565b93606082019480865261061960a061060b608088016135fe565b9660808601978852016136e9565b9660a0840197885215918215610762575b50811561074f575b506106e9579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d519451925193511697511515966040519461066b8661329e565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b169216171790550192919061059f565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff9150161538610632565b1591503861062a565b6040519180602084016020855252604083019190845b8181106107b257857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff6107e8606089016134ea565b16606082015263ffffffff6107ff608089016135fe565b16608082015261081160a088016136e9565b151560a0820152019401929101610781565b5080fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff811161082357610877903690600401613425565b73ffffffffffffffffffffffffffffffffffffffff600154163303610c4b575b9091807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8184360301915b83811015610c47578060051b850135838112156104585785019160808336031261015557604051926080840184811067ffffffffffffffff821117610c1a5760405261090c816134ea565b845261091a602082016136e9565b9560208501968752604082013567ffffffffffffffff81116104545761094390369084016139df565b9160408601928352606081013567ffffffffffffffff8111610c165761096b913691016139df565b976060860198895267ffffffffffffffff86511684526002602052604084209160ff835460f01c16895115158091151503610b8b575b50600185989301975b8a5180518210156109e957906109e273ffffffffffffffffffffffffffffffffffffffff6109da83600195613850565b51168b613a44565b50016109aa565b5050979095939896929196518051610b3d575b50815151610a14575b505050506001019390936108c1565b9795929796909396949194511515600014610b0657855b84518051821015610aa857610a558273ffffffffffffffffffffffffffffffffffffffff92613850565b51168015610a715790610a6a60019288613c0b565b5001610a2b565b60248867ffffffffffffffff8c51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b505095919496600193919450610afa67ffffffffffffffff7f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328192511692516040519182916020835260208301906135b4565b0390a290388080610a05565b60248667ffffffffffffffff8a51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586610b8267ffffffffffffffff875116926040519182916020835260208301906135b4565b0390a2386109fc565b83547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560f01b7eff000000000000000000000000000000000000000000000000000000000000161784557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492602067ffffffffffffffff8a511692604051908152a2386109a1565b8480fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5080f35b73ffffffffffffffffffffffffffffffffffffffff60065416330315610897576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b50346101555760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576004359067ffffffffffffffff8211610155576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261015557604051916101c0830183811067ffffffffffffffff821117610c1a57604052610d2e816004016134ea565b8352610d3c602482016134ea565b6020840152610d4d604482016134ea565b6040840152610d5e606482016135fe565b6060840152610d6f608482016135fe565b6080840152610d8060a4820161352e565b60a084015260c481013560c084015260e481013567ffffffffffffffff811161045857610db390600436918401016134ff565b60e084015261010481013567ffffffffffffffff811161045857610ddd90600436918401016134ff565b61010084015261012481013567ffffffffffffffff811161045857610e0890600436918401016134ff565b61012084015261014481013567ffffffffffffffff811161045857610e3390600436918401016134ff565b61014084015261016481013567ffffffffffffffff811161045857610e5e90600436918401016134ff565b61016084015261018481013567ffffffffffffffff81116104585781013660238201121561045857600481013590610e9582613456565b91610ea360405193846132ba565b80835260051b8101602401602083013682116115155760248301905b8282106114df57505050506101808401526101a481013567ffffffffffffffff811161045857610ef4913691016004016134ff565b6101a083015260243560443567ffffffffffffffff811161045857610f1d9036906004016133f7565b610f3167ffffffffffffffff865116613893565b67ffffffffffffffff855116808552600260205273ffffffffffffffffffffffffffffffffffffffff6040862054169081156114b4576020906044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156114a957859161148a575b501561145e576101e181106114365780600411610454577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d0000000000000000000000000000000000000000000000000000000081036113e757506101809261017c600083861161015557507fffffffff000000000000000000000000000000000000000000000000000000009084810135808316918703600481106113d2575b5050167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361138357506000906101a09485811161045857838611610458576110a9908087039086016137b2565b9081810361135557505067ffffffffffffffff86511685526004602052604085209560ff600388015460201c161561131f575061011c90506000828211610155575060019061111d907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc85016137b2565b950154948581036112ef57508394508083116112ea57836102047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f6020969487956004995087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6082019061019c6040519b8c9a8b997f57ecfd28000000000000000000000000000000000000000000000000000000008b526040838c0152508260448b01520160648901378b6102008801528186880161020060248a0152526102248701378986857ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608489010101015201168201010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156112df5782916112b0575b50156112885780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b6112d2915060203d6020116112d8575b6112ca81836132ba565b8101906137ed565b3861127f565b503d6112c0565b6040513d84823e3d90fd5b505050fd5b84604491877f7d8b101a000000000000000000000000000000000000000000000000000000008352600452602452fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff16600452602485fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b839250829060040360031b1b1616388061105a565b7fadaf77390000000000000000000000000000000000000000000000000000000085527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604484fd5b6004847f1ede477b000000000000000000000000000000000000000000000000000000008152fd5b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6114a3915060203d6020116112d8576112ca81836132ba565b38610fb1565b6040513d87823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b813567ffffffffffffffff811161151157602091611506839283600436928a01010161360f565b815201910190610ebf565b8780fd5b8580fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555767ffffffffffffffff6115ac6134d3565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b81811061162e5773ffffffffffffffffffffffffffffffffffffffff8561162a88611602818903826132ba565b604051938360ff869560f01c16151585521660208401526060604084015260608301906135b4565b0390f35b82548452602090930192600192830192016115d5565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760035461167f81613456565b9061168d60405192836132ba565b808252600383527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9280602084015b8383106116d9576040516020808252819061162a9082018861353d565b6040518287546116e88161375f565b80845290600181169081156117585750600114611722575b5060019282611714859460209403826132ba565b8152019601920191946116bc565b888552602085208592505b81831061174257505081016020016001611700565b600181602092548386880101520192019161172d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611700565b50346101555760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576117d16134d3565b60243567ffffffffffffffff81116104585760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610458576040519061181c8261329e565b806004013567ffffffffffffffff8111610c165761184090600436918401016134ff565b8252602481013567ffffffffffffffff8111610c165761186690600436918401016134ff565b6020830152604481013567ffffffffffffffff8111610c1657810136602382011215610c1657600481013561189a81613456565b916118a860405193846132ba565b818352602060048185019360061b830101019036821161151157602401915b8183106119e95750505060408301526118e2606482016133d6565b6060830152608481013567ffffffffffffffff8111610c1657608091600461190d92369201016134ff565b9101526044359067ffffffffffffffff82116104585761193a67ffffffffffffffff9236906004016134ff565b5061194361351d565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff604082205416156119be57816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f4d1aff7e000000000000000000000000000000000000000000000000000000008252600452fd5b604083360312611511576040516040810181811067ffffffffffffffff821117611a3457916020916040938452611a1f866133d6565b815282860135838201528152019201916118c7565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557805473ffffffffffffffffffffffffffffffffffffffff81163303611b1e577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760408051611b8281613282565b828152826020820152015261162a604051611b9c81613282565b73ffffffffffffffffffffffffffffffffffffffff60055416815261ffff60065473ffffffffffffffffffffffffffffffffffffffff8116602084015260a01c16604082015260405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff81116108235736602382011215610823578060040135611c9481613456565b91611ca260405193846132ba565b8183526024602084019260051b82010190368211610c165760248101925b82841061215d578585611cd1613805565b60035490805160035492611ce484613456565b93611cf260405195866132ba565b808552600386527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9086602087015b82821061209d5750505050845b818110611faa575050835b818110611d8c57847fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586611d7e866104188760405193849360408552604085019061353d565b90838203602085015261353d565b611d968184613850565b5160035468010000000000000000811015611f7d57806001611dbb9201600355613995565b919091611f515780519067ffffffffffffffff8211611f2457611dde835461375f565b601f8111611ee9575b50602090601f8311600114611e4657600194939291899183611e3b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82861b9260031b1c19161790555b01611d39565b015190508980611e04565b83895281892091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168a5b818110611ed1575091600196959492918388959310611e9a575b505050811b019055611e35565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611e8d565b92936020600181928786015181550195019301611e73565b611f1490848a5260208a20601f850160051c81019160208610611f1a575b601f0160051c01906139c8565b88611de7565b9091508190611f07565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6024877f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6003548015612070577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611fde81613995565b61204457908782611ff2600195945461375f565b80612004575b50505060035501611d2e565b601f811186146120195750555b878980611ff8565b8183526020832061203491601f0160051c81019087016139c8565b8082528160208120915555612011565b6024887f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b6040518985546120ac8161375f565b808452906001811690811561211c57506001146120e6575b50600192826120d8859460209403826132ba565b815201940191019092611d21565b868c5260208c208c92505b818310612106575050810160200160016120c4565b60018160209254838688010152019201916120f1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506120c4565b833567ffffffffffffffff811161219e5782013660438201121561219e576020916121938392369060446024820135910161346e565b815201930192611cc0565b8680fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff8111610823576121f2903690600401613425565b9073ffffffffffffffffffffffffffffffffffffffff600554169081156123b057835b838110156123ac578060051b82013573ffffffffffffffffffffffffffffffffffffffff8116809103611515576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156123a1578791612370575b5080612296575b5050600101612215565b602087604051828101907fa9059cbb000000000000000000000000000000000000000000000000000000008252886024820152846044820152604481526122de6064826132ba565b519082865af1156123655786513d61235c5750813b155b6123305790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a3903861228c565b602487837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b600114156122f5565b6040513d88823e3d90fd5b90506020813d8211612399575b8161238a602093836132ba565b8101031261219e575138612285565b3d915061237d565b6040513d89823e3d90fd5b8480f35b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346101555760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155576004359067ffffffffffffffff821161015557816004018236036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261045857612455613390565b5060843567ffffffffffffffff8111610454576124769036906004016133f7565b602486949294019261248f61248a856136f6565b613893565b612498846136f6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd610124890135910181121561219e57870160048101359067ffffffffffffffff82116115115760248101918036038313612885576020908201919091031261219e57359073ffffffffffffffffffffffffffffffffffffffff821680920361219e5767ffffffffffffffff168087526002602052604087209081549073ffffffffffffffffffffffffffffffffffffffff8216908115612c54576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612c49578990612be6575b73ffffffffffffffffffffffffffffffffffffffff9150163303612bba5760f01c60ff16612b72575b505067ffffffffffffffff6125d4846136f6565b1685526004602052604085209160038301549360ff8560201c1615612b3457506101848701906001612606838361370b565b905003612afe57906126179161370b565b15612ad15780357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff418236030181121561219e576126569136910161360f565b9360408501519160208380518101031261219e5760208301519273ffffffffffffffffffffffffffffffffffffffff84168094036115115773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016809403612a9357506080860180516020815111612a5557506002850154801561296f5760a49150985b01359061ffff8216809203611511576040519261270a84613282565b602435845282602085015260408401916107d083528993612889575b505063ffffffff73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001697519554915116925193604051947f8e1d1a9d0000000000000000000000000000000000000000000000000000000060208701526024860152602485526127ac6044866132ba565b873b156128855793889795929363ffffffff899561282c94829a986040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c01521660248a01526044890152606488015260848701521660a485015260c484015261010060e48401526101048301906132fb565b03925af1918215612878578161162a93612868575b5050604051906128526020836132ba565b81526040519182916020835260208301906132fb565b612871916132ba565b3881612841565b50604051903d90823e3d90fd5b8880fd5b6103e8835291925090801561291857602081036128ed5781602091810103126115115735905b63ffffffff82116128c1573880612726565b602488837fb6f15d0f000000000000000000000000000000000000000000000000000000008252600452fd5b7f4a895e59000000000000000000000000000000000000000000000000000000008952600452602488fd5b5050855161ffff60065460a01c1690818102918183041490151715612942576127109004906128af565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b50516020815111612a13576020815191015190602081106129e2575b8060031b90808204600814901517156129b5576101000361010081116129b5571c9760a4906126ee565b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b169061298b565b612a51906040519182917fe0d7fb020000000000000000000000000000000000000000000000000000000083526020600484015260248301906132fb565b0390fd5b612a51906040519182917fa3c8cf090000000000000000000000000000000000000000000000000000000083526020600484015260248301906132fb565b612a51906040519182917f22d4cfe20000000000000000000000000000000000000000000000000000000083526020600484015260248301906132fb565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b612b0b602492889261370b565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8667ffffffffffffffff612b496024936136f6565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612b8f57806125c0565b7fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6024887f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612c41575b81612c00602093836132ba565b81010312612885575173ffffffffffffffffffffffffffffffffffffffff811681036128855773ffffffffffffffffffffffffffffffffffffffff90612597565b3d9150612bf3565b6040513d8b823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008a52600452602489fd5b50346101555760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101555760043567ffffffffffffffff811161082357612ccf90369060040161335a565b612cd7613805565b612ce081613456565b91612cee60405193846132ba565b81835260c0602084019202810190368211610c1657915b818310612f4a578480855b8051831015610c4757612d238382613850565b519267ffffffffffffffff6020612d3a8385613850565b51015116938415612f1e578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff00000000000000000000000000000000000000001617825560808101805163ffffffff1615612ef25760019495969260ff73ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040945184547fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff000000000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a086015160d01b169360b01b1691161717809455511691835192835260f01c1615156020820152a2019190612d10565b602486887f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b60c083360312610c165760405190612f6182613237565b83359073ffffffffffffffffffffffffffffffffffffffff8216820361219e578260209260c09452612f948387016134ea565b83820152612fa4604087016136e9565b6040820152612fb56060870161352e565b6060820152612fc6608087016135fe565b6080820152612fd760a087016135fe565b60a0820152815201920191612d05565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610155575061162a6040516130286040826132ba565b601681527f43435450566572696669657220322e302e302d6465760000000000000000000060208201526040519182916020835260208301906132fb565b503461015557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015557608060405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152f35b9050346108235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610823576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361045857602092507f83adcde100000000000000000000000000000000000000000000000000000000811490811561320d575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613206565b60c0810190811067ffffffffffffffff82111761325357604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761325357604052565b60a0810190811067ffffffffffffffff82111761325357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761325357604052565b919082519283825260005b8481106133455750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613306565b9181601f8401121561338b5782359167ffffffffffffffff831161338b5760208085019460c0850201011161338b57565b600080fd5b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361338b57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361338b57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361338b57565b9181601f8401121561338b5782359167ffffffffffffffff831161338b576020838186019501011161338b57565b9181601f8401121561338b5782359167ffffffffffffffff831161338b576020808501948460051b01011161338b57565b67ffffffffffffffff81116132535760051b60200190565b92919267ffffffffffffffff821161325357604051916134b6601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001846132ba565b82948184528183011161338b578281602093846000960137010152565b6004359067ffffffffffffffff8216820361338b57565b359067ffffffffffffffff8216820361338b57565b9080601f8301121561338b5781602061351a9335910161346e565b90565b6064359061ffff8216820361338b57565b359061ffff8216820361338b57565b9080602083519182815201916020808360051b8301019401926000915b83831061356957505050505090565b90919293946020806135a5837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516132fb565b9701930193019193929061355a565b906020808351928381520192019060005b8181106135d25750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016135c5565b359063ffffffff8216820361338b57565b919060c08382031261338b576040519061362882613237565b819380358352602081013567ffffffffffffffff811161338b578261364e9183016134ff565b6020840152604081013567ffffffffffffffff811161338b57826136739183016134ff565b6040840152606081013567ffffffffffffffff811161338b57826136989183016134ff565b6060840152608081013567ffffffffffffffff811161338b57826136bd9183016134ff565b608084015260a08101359167ffffffffffffffff831161338b5760a0926136e492016134ff565b910152565b3590811515820361338b57565b3567ffffffffffffffff8116810361338b5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561338b570180359067ffffffffffffffff821161338b57602001918160051b3603831361338b57565b90600182811c921680156137a8575b602083101461377957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161376e565b3590602081106137c0575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9081602091031261338b5751801515810361338b5790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361382657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156138645760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156139895760009161396a575b506139325750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b613983915060203d6020116112d8576112ca81836132ba565b3861392a565b6040513d6000823e3d90fd5b60035481101561386457600360005260206000200190600090565b80548210156138645760005260206000200190600090565b8181106139d3575050565b600081556001016139c8565b9080601f8301121561338b5781356139f681613456565b92613a0460405194856132ba565b81845260208085019260051b82010192831161338b57602001905b828210613a2c5750505090565b60208091613a39846133d6565b815201910190613a1f565b9060018201918160005282602052604060002054801515600014613c02577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613bd3578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613bd357818103613b67575b50505080548015613b38577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190613af982826139b0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b613bbc613b77613b8793866139b0565b90549060031b1c928392866139b0565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005283602052604060002055388080613ac1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b6000828152600182016020526040902054613c6057805490680100000000000000008210156132535782613c49613b878460018096018555846139b0565b905580549260005201602052604060002055600190565b505060009056fea164736f6c634300081a000a",
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

func (_CCTPVerifier *CCTPVerifierCaller) GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getStaticConfig")

	outstruct := new(GetStaticConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenMessenger = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MessageTransmitterProxy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.UsdcToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LocalDomainIdentifier = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetStaticConfig() (GetStaticConfig,

	error) {
	return _CCTPVerifier.Contract.GetStaticConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetStaticConfig() (GetStaticConfig,

	error) {
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

type CCTPVerifierAllowListStateChangedIterator struct {
	Event *CCTPVerifierAllowListStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPVerifierAllowListStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPVerifierAllowListStateChanged)
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
		it.Event = new(CCTPVerifierAllowListStateChanged)
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

func (it *CCTPVerifierAllowListStateChangedIterator) Error() error {
	return it.fail
}

func (it *CCTPVerifierAllowListStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPVerifierAllowListStateChanged struct {
	DestChainSelector uint64
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CCTPVerifier *CCTPVerifierFilterer) FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListStateChangedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.FilterLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPVerifierAllowListStateChangedIterator{contract: _CCTPVerifier.contract, event: "AllowListStateChanged", logs: logs, sub: sub}, nil
}

func (_CCTPVerifier *CCTPVerifierFilterer) WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPVerifier.contract.WatchLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPVerifierAllowListStateChanged)
				if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
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

func (_CCTPVerifier *CCTPVerifierFilterer) ParseAllowListStateChanged(log types.Log) (*CCTPVerifierAllowListStateChanged, error) {
	event := new(CCTPVerifierAllowListStateChanged)
	if err := _CCTPVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
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
type GetStaticConfig struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	UsdcToken               common.Address
	LocalDomainIdentifier   uint32
}

func (CCTPVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CCTPVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CCTPVerifierAllowListStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492")
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

	GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

		error)

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

	FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPVerifierAllowListStateChangedIterator, error)

	WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *CCTPVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListStateChanged(log types.Log) (*CCTPVerifierAllowListStateChanged, error)

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
