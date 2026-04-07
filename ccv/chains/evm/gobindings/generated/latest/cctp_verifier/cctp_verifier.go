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
	PayloadSizeBytes    uint16
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
	Finality            [4]byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifierArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllowedFinalityConfig\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocations\",\"inputs\":[{\"name\":\"newLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidFastFinalityBps\",\"inputs\":[{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierArgsLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461068d57614be7803803809161001d8285610a45565b8339810190808203610100811261068d5781516001600160a01b038116919082810361068d5760208401516001600160a01b0381169081810361068d5760408601516001600160a01b0381169490939085850361068d5760608801516001600160401b03811161068d57880189601f8201121561068d578051906100a082610a68565b9a6100ae6040519c8d610a45565b828c526020808d019360051b8301019181831161068d5760208101935b8385106109cd575050505050606090607f19011261068d5760405195606087016001600160401b038111888210176108405760405261010c60808901610aa2565b875261011a60a08901610aa2565b9860208801998a5260c08901519861ffff8a168a0361068d5760e06101469160408b019b8c5201610aa2565b9033156109bc57600180546001600160a01b03191633179055600354815190919061017083610a68565b9261017e6040519485610a45565b808452600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90602086015b8382106109175750505060005b81811061088257505060005b8181106106f35750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161021c61020e92604051938493604085526040850190610b3e565b908382036020850152610b3e565b0390a16001600160a01b031680156106d257608052801580156106eb575b80156106e3575b6106d257604051639cdbb18160e01b8152602081600481855afa80156105b65763ffffffff916000916106b3575b50166001810361069a5750602060049160405192838092632c12192160e01b82525afa9081156105b65760009161065b575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa80156105b65763ffffffff9160009161063c575b50166001810361062357506020600491604051928380926367e0ed8360e11b82525afa9081156105b6576000916105da575b506001600160a01b03168381036105c2575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa9081156105b657600091610587575b506101005260a05260018060a01b0360e0511690604051906020600081840163095ea7b360e01b81528560248601528119604486015260448552610392606486610a45565b84519082855af16000513d8261056b575b505015610525575b505060e05160c05160a05161010051604080516001600160a01b0395861681529385166020850152919093169082015263ffffffff9190911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c9150608090a161ffff8251168015801561051a575b610506578151600680546001600160a01b039283166001600160a01b03199091168117909155855160078054875161ffff60a01b60a09190911b169285166001600160b01b0319909116179190911790556040805191825286519092166020820152845161ffff16918101919091527f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197290606090a1604051613fdc9081610c0b823960805181613a11015260a051818181612a64015261315d015260c0518181816116a20152613121015260e051818181612b1a01526130e8015261010051816131890152f35b630c74dcaf60e41b60005260045260246000fd5b50612710811161041f565b61055e610563936040519063095ea7b360e01b602083015260248201526000604482015260448152610558606482610a45565b82610baf565b610baf565b3880806103ab565b90915061057f5750803b15155b38806103a3565b600114610578565b6105a9915060203d6020116105af575b6105a18183610a45565b810190610ab6565b3861034d565b503d610597565b6040513d6000823e3d90fd5b836383395ca960e01b60005260045260245260446000fd5b6020813d60201161061b575b816105f360209383610a45565b810103126106175751906001600160a01b038216820361061457503861030c565b80fd5b5080fd5b3d91506105e6565b6331b6aa1b60e11b600052600160045260245260446000fd5b610655915060203d6020116105af576105a18183610a45565b386102da565b90506020813d602011610692575b8161067660209383610a45565b8101031261068d5761068790610aa2565b386102a1565b600080fd5b3d9150610669565b633785f8f160e01b600052600160045260245260446000fd5b6106cc915060203d6020116105af576105a18183610a45565b3861026f565b6342bcdf7f60e11b60005260046000fd5b508515610241565b50831561023a565b825181101561086c5760208160051b84010151600354680100000000000000008110156108405780600161072a9201600355610ad2565b919091610856578051906001600160401b0382116108405761074c8354610aed565b601f8111610803575b50602090601f8311600114610798576001949392916000918361078d575b5050600019600383901b1c191690841b1790555b016101c8565b015190503880610773565b90601f1983169184600052816000209260005b8181106107eb5750916001969594929183889593106107d2575b505050811b019055610787565b015160001960f88460031b161c191690553880806107c5565b929360206001819287860151815501950193016107ab565b61083090846000526020600020601f850160051c81019160208610610836575b601f0160051c0190610b27565b38610755565b9091508190610823565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600354801561090157600019019061089982610ad2565b92909261085657826108ad60019454610aed565b90816108bf575b5050600355016101bc565b81601f6000931186146108d65750555b38806108b4565b818352602083206108f191601f0160051c8101908701610b27565b80825281602081209155556108cf565b634e487b7160e01b600052603160045260246000fd5b6040516000845461092781610aed565b80845290600181169081156109995750600114610961575b506001928261095385946020940382610a45565b8152019301910190916101af565b6000868152602081209092505b8183106109835750508101602001600161093f565b600181602092548386880101520192019161096e565b60ff191660208581019190915291151560051b840190910191506001905061093f565b639b15e16f60e01b60005260046000fd5b84516001600160401b03811161068d5782019083603f8301121561068d576020820151906001600160401b03821161084057604051610a16601f8401601f191660200182610a45565b828152604084840101861061068d57610a3a60209493859460408685019101610a7f565b8152019401936100cb565b601f909101601f19168101906001600160401b0382119082101761084057604052565b6001600160401b0381116108405760051b60200190565b60005b838110610a925750506000910152565b8181015183820152602001610a82565b51906001600160a01b038216820361068d57565b9081602091031261068d575163ffffffff8116810361068d5790565b60035481101561086c57600360005260206000200190600090565b90600182811c92168015610b1d575b6020831014610b0757565b634e487b7160e01b600052602260045260246000fd5b91607f1691610afc565b818110610b32575050565b60008155600101610b27565b9080602083519182815201916020808360051b8301019401926000915b838310610b6a57505050505090565b909192939460208080600193601f198682030187528951610b9681518092818552858086019101610a7f565b601f01601f191601019701959491909101920190610b5b565b906000602091828151910182855af1156105b6576000513d610c0157506001600160a01b0381163b155b610be05750565b635274afe760e01b60009081526001600160a01b0391909116600452602490fd5b60011415610bd956fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146131b05750806306285c691461309d578063181f5a771461301e5780632969470614612796578063597b95c3146124385780635cb80c5d146122115780635ef2c64b14611db85780637437ff9f14611cc557806379ba509714611be057806387ae929214611b8e578063898068fc146119ab57806389e364c71461109d5780638da5cb5b1461104b578063b6cfa3b714610f8f578063c9b146b314610b3d578063d52e545a1461085b578063dfadfa3514610776578063e023ddb114610563578063ec6ae7a714610502578063f2fde38b14610415578063f4cdd89e1461016e5763fe163eed1461011357600080fd5b3461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760206040517f35a25838000000000000000000000000000000000000000000000000000000008152f35b80fd5b503461016b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b576101a66135c8565b9060243567ffffffffffffffff81116103d15760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126103d157604051906101f282613321565b806004013567ffffffffffffffff81116103d5576102169060043691840101613605565b8252602481013567ffffffffffffffff81116103d55761023c9060043691840101613605565b6020830152604481013567ffffffffffffffff81116103d5578101366023820112156103d5576004810135610270816134d4565b9161027e604051938461333d565b818352602060048185019360061b830101019036821161041157602401915b8183106103d95750505060408301526102b860648201613423565b6060830152608481013567ffffffffffffffff81116103d55760809160046102e39236920101613605565b91015260443567ffffffffffffffff81116103d157610306903690600401613605565b50606435917fffffffff00000000000000000000000000000000000000000000000000000000831683036103d15767ffffffffffffffff1690818152600260205260408120549173ffffffffffffffffffffffffffffffffffffffff8316156103a55760608361037c8660045460e01b90613bea565b61ffff60405191818160a01c16835263ffffffff8160b01c16602084015260d01c166040820152f35b602492507f4d1aff7e000000000000000000000000000000000000000000000000000000008252600452fd5b5080fd5b8380fd5b60408336031261041157602060409182516103f38161329e565b6103fc86613423565b8152828601358382015281520192019161029d565b8680fd5b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5773ffffffffffffffffffffffffffffffffffffffff610462613400565b61046a613aad565b163381146104da57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57602060045460e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b503461016b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760405161059f81613305565b6105a7613400565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361077257602082019081526044359061ffff821682036103d557604083019182526105ee613aad565b61ffff82511680158015610767575b61073c5750916107369173ffffffffffffffffffffffffffffffffffffffff7f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197294818451167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065551167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755517fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000806007549360a01b161691161760075560405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390a180f35b7fc74dcaf0000000000000000000000000000000000000000000000000000000008552600452602484fd5b5061271081116105fd565b8280fd5b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57604060a09167ffffffffffffffff6107bc6135c8565b82608085516107ca81613321565b82815282602082015282878201528260608201520152168152600560205220604051906107f682613321565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d1576108ab903690600401613472565b6108b6929192613aad565b81925b81841015610a855760c0840281019360c0853603126103d557604051946108df866132e9565b803590818752602081013590602088019782895260408101906040830135825261090b606084016135df565b93606082019480865261093360a0610925608088016135f4565b966080860197885201613775565b9660a0840197885215918215610a7c575b508115610a69575b50610a03579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d519451925193511697511515966040519461098586613321565b85526020850192835260408501938452606085019889526080850197885251168c52600560205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b16921617179055019291906108b9565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff915016153861094c565b15915038610944565b6040519180602084016020855252604083019190845b818110610acc57857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff610b02606089016135df565b16606082015263ffffffff610b19608089016135f4565b166080820152610b2b60a08801613775565b151560a0820152019401929101610a9b565b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d157610b8d9036906004016134a3565b73ffffffffffffffffffffffffffffffffffffffff600193929354163303610f47575b819291907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301915b81811015610f43578060051b84013583811215610f3f578401608081360312610f3f57604051956080870187811067ffffffffffffffff821117610f1257604052610c25826135df565b8752610c3360208301613775565b9360208801948552604083013567ffffffffffffffff811161077257610c5c9036908501613b85565b9260408901938452606081013567ffffffffffffffff81116103d557610c8491369101613b85565b966060890197885267ffffffffffffffff89511683526002602052604083209160ff835460e01c16875115158091151503610e89575b50600184989301975b89518051821015610d48579073ffffffffffffffffffffffffffffffffffffffff610cf082600194613af8565b51168c610cfd828d613ddb565b610d0a575b505001610cc3565b602067ffffffffffffffff7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d8292511692604051908152a2388c610d02565b5050989392975094959095825151610d6b575b5050505060010193909293610bdb565b9790949196959297511515600014610e5257855b87518051821015610e3f57610da98273ffffffffffffffffffffffffffffffffffffffff92613af8565b51168015610e08579081610dbf60019389613f73565b610dcb575b5001610d7f565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80602067ffffffffffffffff8d511692604051908152a238610dc4565b60248867ffffffffffffffff8c51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b5050965092509293506001388080610d5b565b60248667ffffffffffffffff8a51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b83547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560e01b7cff00000000000000000000000000000000000000000000000000000000161784557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492602067ffffffffffffffff8d511692604051908152a238610cba565b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b8480f35b73ffffffffffffffffffffffffffffffffffffffff60075416330315610bb0576004827f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b577f307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a6020610fea61326a565b610ff2613aad565b8060e01c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000060045416176004557fffffffff0000000000000000000000000000000000000000000000000000000060405191168152a180f35b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461016b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d1576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126103d157604051906101c0820182811067ffffffffffffffff82111761197e57604052611137816004016135df565b8252611145602482016135df565b6020830152611156604482016135df565b6040830152611167606482016135f4565b6060830152611178608482016135f4565b608083015260a48101357fffffffff00000000000000000000000000000000000000000000000000000000811681036103d55760a083015260c481013560c083015260e481013567ffffffffffffffff81116103d5576111de9060043691840101613605565b60e083015261010481013567ffffffffffffffff81116103d5576112089060043691840101613605565b61010083015261012481013567ffffffffffffffff81116103d5576112339060043691840101613605565b61012083015261014481013567ffffffffffffffff81116103d55761125e9060043691840101613605565b61014083015261016481013567ffffffffffffffff81116103d5576112899060043691840101613605565b61016083015261018481013567ffffffffffffffff81116103d5578101366023820112156103d5576004810135906112c0826134d4565b916112ce604051938461333d565b80835260051b8101602401602083013682116104115760248301905b82821061194857505050506101808301526101a48101359067ffffffffffffffff82116103d55760046113209236920101613605565b6101a08201526024359060443567ffffffffffffffff81116103d55761134a903690600401613444565b9261135f67ffffffffffffffff8451166139ab565b67ffffffffffffffff835116808652600260205273ffffffffffffffffffffffffffffffffffffffff60408720541690811561191d576020906044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119125786916118f3575b50156118c7576101e1841061189f578360041161189b577fffffffff000000000000000000000000000000000000000000000000000000008235167f35a2583800000000000000000000000000000000000000000000000000000000810361184c57506101809261017c600086861161016b5750611484817fffffffff00000000000000000000000000000000000000000000000000000000928601908703906138f2565b167f35a258380000000000000000000000000000000000000000000000000000000081036117fd57506000916101a0948581116103d5578686116103d5576114d190808703908601613958565b908181036117cf57505067ffffffffffffffff81511686526005602052604086209160038301549160ff8360201c161561179957506000905085600c1161016b575063ffffffff6115266004600886016138f2565b60e01c911680820361176b5750600091905061011c85811161077257600190611574907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc8601613958565b91015480820361173d5750508484841161016b576102047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f60049760209788965087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6082019061019c6040519b8c9a8b997f57ecfd28000000000000000000000000000000000000000000000000000000008b526040838c0152508260448b01520160648901378b6102008801528186880161020060248a0152526102248701378986857ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608489010101015201168201010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611732578291611703575b50156116db5780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b611725915060203d60201161172b575b61171d818361333d565b810190613993565b386116d2565b503d611713565b6040513d84823e3d90fd5b7f7d8b101a000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b7fe366a117000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000885267ffffffffffffffff16600452602487fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008852600452602452604486fd5b7fadaf77390000000000000000000000000000000000000000000000000000000087527f35a2583800000000000000000000000000000000000000000000000000000000600452602452604486fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f35a2583800000000000000000000000000000000000000000000000000000000600452602452604485fd5b8480fd5b6004857f1ede477b000000000000000000000000000000000000000000000000000000008152fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61190c915060203d60201161172b5761171d818361333d565b386113df565b6040513d88823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b813567ffffffffffffffff811161197a5760209161196f839283600436928a010101613623565b8152019101906112ea565b8880fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5767ffffffffffffffff6119ec6135c8565b8260a06040516119fb816132e9565b82815282602082015282604082015282606082015282608082015201521680825260026020526040822080549260405193611a35856132e9565b73ffffffffffffffffffffffffffffffffffffffff8116855260208501938452604085019060ff8160e01c1615158252606086019361ffff8260a01c1685526001608088019163ffffffff8460b01c16835261ffff60a08a019460d01c1684520194604051938460208854918281520190819888526020882090885b818110611b78575050509261ffff86959363ffffffff93611adf67ffffffffffffffff9d998498038961333d565b6040519c8d9c8d73ffffffffffffffffffffffffffffffffffffffff60e082019c51169052511660208d015251151560408c0152511660608a015251166080880152511660a086015260e060c086015251809152610100840192915b818110611b49575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611b3b565b8254845260209093019260019283019201611ab1565b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57611bdc611bc86137d5565b604051918291602083526020830190613551565b0390f35b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57805473ffffffffffffffffffffffffffffffffffffffff81163303611c9d577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760408051611d0181613305565b8281528260208201520152611bdc604051611d1b81613305565b73ffffffffffffffffffffffffffffffffffffffff60065416815261ffff60075473ffffffffffffffffffffffffffffffffffffffff8116602084015260a01c16604082015260405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d157366023820112156103d1578060040135611e13816134d4565b91611e21604051938461333d565b8183526024602084019260051b8201019036821161189b5760248101925b8284106121d0578585611e50613aad565b600354908051611e5e6137d5565b92845b8181106120dd575050835b818110611ebf57847fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586611eb18661073687604051938493604085526040850190613551565b908382036020850152613551565b611ec98184613af8565b51600354680100000000000000008110156120b057806001611eee9201600355613b3b565b9190916120845780519067ffffffffffffffff821161205757611f118354613782565b601f811161201c575b50602090601f8311600114611f7957600194939291899183611f6e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82861b9260031b1c19161790555b01611e6c565b015190508980611f37565b83895281892091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168a5b818110612004575091600196959492918388959310611fcd575b505050811b019055611f68565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611fc0565b92936020600181928786015181550195019301611fa6565b61204790848a5260208a20601f850160051c8101916020861061204d575b601f0160051c0190613b6e565b88611f1a565b909150819061203a565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6024877f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60035480156121a3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161211181613b3b565b612177579087826121256001959454613782565b80612137575b50505060035501611e61565b601f8111861461214c5750555b87898061212b565b8183526020832061216791601f0160051c8101908701613b6e565b8082528160208120915555612144565b6024887f4e487b7100000000000000000000000000000000000000000000000000000000815280600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b833567ffffffffffffffff81116104115782013660438201121561041157602091612206839236906044602482013591016134ec565b815201930192611e3f565b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d1576122619036906004016134a3565b9073ffffffffffffffffffffffffffffffffffffffff6006541690811561241057835b83811015610f43578060051b82013573ffffffffffffffffffffffffffffffffffffffff8116809103610f3f576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156124055787916123d4575b5080612305575b5050600101612284565b602087604051828101907fa9059cbb0000000000000000000000000000000000000000000000000000000082528860248201528460448201526044815261234d60648261333d565b519082865af1156119125786513d6123cb5750813b155b61239f5790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a390386122fb565b602487837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b60011415612364565b90506020813d82116123fd575b816123ee6020938361333d565b810103126104115751386122f4565b3d91506123e1565b6040513d89823e3d90fd5b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b503461016b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5760043567ffffffffffffffff81116103d157612488903690600401613472565b612490613aad565b612499816134d4565b916124a7604051938461333d565b81835260c060208401920281019036821161189b57915b8183106126f9578480855b80518310156126f5576124dc8382613af8565b519267ffffffffffffffff6020850151169384156126c9578484526002602052604080852082518154928401517fffffff00ffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560e01b7cff000000000000000000000000000000000000000000000000000000001691909117815590606081015182547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff00000000000000000000000000000000000000001617825560808101805163ffffffff161561269d5760019495969260ff73ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040945184547fffffffff000000000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff000000000000000000000000000000000000000000007bffff000000000000000000000000000000000000000000000000000060a086015160d01b169360b01b1691161717809455511691835192835260e01c1615156020820152a20191906124c9565b602486887f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261189b5760405190612710826132e9565b83359073ffffffffffffffffffffffffffffffffffffffff82168203610411578260209260c094526127438387016135df565b8382015261275360408701613775565b6040820152612764606087016136fd565b6060820152612775608087016135f4565b608082015261278660a087016136fd565b60a08201528152019201916124be565b503461016b5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b576004359067ffffffffffffffff821161016b5781600401918036036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610772576128146133dd565b5060843567ffffffffffffffff81116103d557612835903690600401613444565b949092602481019261284e6128498561370c565b6139ab565b6128578461370c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd610124840135910181121561041157820160048101359067ffffffffffffffff8211612c86576024810191803603831361197a576020908201919091031261041157359073ffffffffffffffffffffffffffffffffffffffff82168092036104115767ffffffffffffffff168087526002602052604087209081549073ffffffffffffffffffffffffffffffffffffffff8216908115612ff3576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612fe8578990612f85575b73ffffffffffffffffffffffffffffffffffffffff9150163303612f595760e01c60ff16612f11575b505067ffffffffffffffff6129938461370c565b1685526005602052604085209160038301549360ff8560201c1615612ed3575061018482019060016129c58383613721565b905003612e9d57906129d691613721565b15612e705780357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff418236030181121561041157612a1591369101613623565b936040850151916020838051810103126104115760208301519273ffffffffffffffffffffffffffffffffffffffff8416809403612c865773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016809403612e3257506080860180516020815111612df457506002850154908115612d105750975b60405192612ab88461329e565b602435845260208401916107d0835260a48a9401357fffffffff000000000000000000000000000000000000000000000000000000008116809103612d0c57612c59575b505063ffffffff73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001697519554915116925193604051947f35a25838000000000000000000000000000000000000000000000000000000006020870152602486015260248552612b8260448661333d565b873b1561197a579388979592938894612c009363ffffffff99976040519d8e9b8c9a8b997f779b432d000000000000000000000000000000000000000000000000000000008b5260048b015216602489015260448801526064870152608486015260a485015260c484015261010060e484015261010483019061337e565b03925af1918215612c4c5781611bdc93612c3c575b505060405190612c2660208361333d565b815260405191829160208352602083019061337e565b612c459161333d565b3881612c15565b50604051903d90823e3d90fd5b6103e88352919250908015612cb55760208103612c8a578160209181010312612c865735905b3880612afc565b8780fd5b7f4a895e59000000000000000000000000000000000000000000000000000000008952600452602488fd5b5050855161ffff60075460a01c1690818102918183041490151715612cdf57612710900490612c7f565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b8a80fd5b9050516020815111612db257602081519101519060208110612d81575b8060031b9080820460081490151715612d5457610100036101008111612d54571c97612aab565b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b1690612d2d565b612df0906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260206004840152602483019061337e565b0390fd5b612df0906040519182917fa3c8cf0900000000000000000000000000000000000000000000000000000000835260206004840152602483019061337e565b612df0906040519182917f22d4cfe200000000000000000000000000000000000000000000000000000000835260206004840152602483019061337e565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b612eaa6024928892613721565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8667ffffffffffffffff612ee860249361370c565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612f2e578061297f565b7fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6024887f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612fe0575b81612f9f6020938361333d565b8101031261197a575173ffffffffffffffffffffffffffffffffffffffff8116810361197a5773ffffffffffffffffffffffffffffffffffffffff90612956565b3d9150612f92565b6040513d8b823e3d90fd5b7f4d1aff7e000000000000000000000000000000000000000000000000000000008a52600452602489fd5b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b5750611bdc60405161305f60408261333d565b601281527f43435450566572696669657220322e302e300000000000000000000000000000602082015260405191829160208352602083019061337e565b503461016b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016b57608060405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152f35b9050346103d15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103d1576020907fffffffff0000000000000000000000000000000000000000000000000000000061320d61326a565b167fd3e969cd000000000000000000000000000000000000000000000000000000008114908115613240575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613239565b600435907fffffffff000000000000000000000000000000000000000000000000000000008216820361329957565b600080fd5b6040810190811067ffffffffffffffff8211176132ba57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176132ba57604052565b6060810190811067ffffffffffffffff8211176132ba57604052565b60a0810190811067ffffffffffffffff8211176132ba57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176132ba57604052565b919082519283825260005b8481106133c85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613389565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361329957565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361329957565b359073ffffffffffffffffffffffffffffffffffffffff8216820361329957565b9181601f840112156132995782359167ffffffffffffffff8311613299576020838186019501011161329957565b9181601f840112156132995782359167ffffffffffffffff83116132995760208085019460c0850201011161329957565b9181601f840112156132995782359167ffffffffffffffff8311613299576020808501948460051b01011161329957565b67ffffffffffffffff81116132ba5760051b60200190565b92919267ffffffffffffffff82116132ba5760405191613534601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461333d565b829481845281830111613299578281602093846000960137010152565b9080602083519182815201916020808360051b8301019401926000915b83831061357d57505050505090565b90919293946020806135b9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161337e565b9701930193019193929061356e565b6004359067ffffffffffffffff8216820361329957565b359067ffffffffffffffff8216820361329957565b359063ffffffff8216820361329957565b9080601f8301121561329957816020613620933591016134ec565b90565b919060c083820312613299576040519061363c826132e9565b819380358352602081013567ffffffffffffffff81116132995782613662918301613605565b6020840152604081013567ffffffffffffffff81116132995782613687918301613605565b6040840152606081013567ffffffffffffffff811161329957826136ac918301613605565b6060840152608081013567ffffffffffffffff811161329957826136d1918301613605565b608084015260a08101359167ffffffffffffffff83116132995760a0926136f89201613605565b910152565b359061ffff8216820361329957565b3567ffffffffffffffff811681036132995790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613299570180359067ffffffffffffffff821161329957602001918160051b3603831361329957565b3590811515820361329957565b90600182811c921680156137cb575b602083101461379c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613791565b600354906137e2826134d4565b916137f0604051938461333d565b808352600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9190602085015b82821061382f5750505050565b6040516000855461383f81613782565b80845290600181169081156138b15750600114613879575b506001928261386b8594602094038261333d565b815201940191019092613822565b6000878152602081209092505b81831061389b57505081016020016001613857565b6001816020925483868801015201920191613886565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050613857565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613926575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b359060208110613966575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b90816020910312613299575180151581036132995790565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613aa157600091613a82575b50613a4a5750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b613a9b915060203d60201161172b5761171d818361333d565b38613a42565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303613ace57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015613b0c5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600354811015613b0c57600360005260206000200190600090565b8054821015613b0c5760005260206000200190600090565b818110613b79575050565b60008155600101613b6e565b9080601f83011215613299578135613b9c816134d4565b92613baa604051948561333d565b81845260208085019260051b82010192831161329957602001905b828210613bd25750505090565b60208091613bdf84613423565b815201910190613bc5565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115613ccc57613c1d81613cd1565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c1616613ccc5761ffff8360e01c168015918215613cbb575b5050613c67575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff161090503880613c5d565b505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115613dd7577dffff00000000000000000000000000000000000000000000000000000000811615613dce5760ff60015b169060f01c80613d69575b50600103613d3c5750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b60108110613d7a5750613d31565b6001811b8216613d8d575b600101613d6c565b9160018101809111613d9f5791613d85565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff6000613d26565b5050565b9060018201918160005282602052604060002054801515600014613f6a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613d9f578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613d9f57818103613efe575b50505080548015613ecf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190613e908282613b56565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b613f53613f0e613f1e9386613b56565b90549060031b1c92839286613b56565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005283602052604060002055388080613e58565b50505050600090565b6000828152600182016020526040902054613fc857805490680100000000000000008210156132ba5782613fb1613f1e846001809601855584613b56565b905580549260005201602052604060002055600190565b505060009056fea164736f6c634300081a000a",
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

func (_CCTPVerifier *CCTPVerifierCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CCTPVerifier.Contract.GetAllowedFinalityConfig(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CCTPVerifier.Contract.GetAllowedFinalityConfig(&_CCTPVerifier.CallOpts)
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

func (_CCTPVerifier *CCTPVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, requestedFinality)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _CCTPVerifier.Contract.GetFee(&_CCTPVerifier.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _CCTPVerifier.Contract.GetFee(&_CCTPVerifier.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_CCTPVerifier *CCTPVerifierCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemoteChainConfig = *abi.ConvertType(out[0], new(BaseVerifierRemoteChainConfigArgs)).(*BaseVerifierRemoteChainConfigArgs)
	outstruct.AllowedSendersList = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

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

func (_CCTPVerifier *CCTPVerifierTransactor) SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPVerifier.contract.Transact(opts, "setAllowedFinalityConfig", allowedFinality)
}

func (_CCTPVerifier *CCTPVerifierSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetAllowedFinalityConfig(&_CCTPVerifier.TransactOpts, allowedFinality)
}

func (_CCTPVerifier *CCTPVerifierTransactorSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPVerifier.Contract.SetAllowedFinalityConfig(&_CCTPVerifier.TransactOpts, allowedFinality)
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
	Senders           common.Address
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
	Senders           common.Address
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
	AllowedFinality [4]byte
	Raw             types.Log
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
	RemoteChainConfig  BaseVerifierRemoteChainConfigArgs
	AllowedSendersList []common.Address
}
type GetStaticConfig struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	UsdcToken               common.Address
	LocalDomainIdentifier   uint32
}

func (CCTPVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80")
}

func (CCTPVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0x9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82")
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

func (CCTPVerifierFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
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
	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (CCTPVerifierDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCTPVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

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

	SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error)

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

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPVerifierFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPVerifierFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CCTPVerifierFinalityConfigSet, error)

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
