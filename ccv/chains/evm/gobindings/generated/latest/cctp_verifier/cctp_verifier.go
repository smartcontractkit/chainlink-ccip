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
	TokenMessengerProxy     common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessengerProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPTokenMessengerProxy\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.StaticConfig\",\"components\":[{\"name\":\"tokenMessengerProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessengerProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidFastFinalityBps\",\"inputs\":[{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461050a57614283803803809161001d82856107f5565b8339810181810360c0811261050a578251906001600160a01b0382169081830361050a5760208501516001600160a01b0381169182820361050a5760408701516001600160401b03811161050a5787019086601f8301121561050a578151916001600160401b03831161075057604051976100a2601f8501601f19166020018a6107f5565b8389526020848301011161050a576060926100c3916020808b019101610818565b605f19011261050a5760405193606085016001600160401b03811186821017610750576040526100f56060880161083b565b855260a06101056080890161083b565b976020870198895201519561ffff8716870361050a576040860196875233156107e457600180546001600160a01b0319163317905560405160035490919060008361014f8361086b565b808252916001841680156107c657600114610766575b610171925003846107f5565b8151906001600160401b0382116107505761018b9061086b565b601f81116106eb575b506020601f821160011461066e576101ed928260008051602061424383398151915295936101fb93600091610663575b508160011b916000199060031b1c1916176003555b6040519384936040855260408501906108a5565b9083820360208501526108a5565b0390a18315801561065b575b6104b557604051639cdbb18160e01b8152602081600481885afa80156105175763ffffffff9160009161063c575b5016600181036106235750604051632c12192160e01b8152602081600481885afa908115610517576000916105e9575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa80156105175763ffffffff916000916105ca575b5016600181036105b157506020600491604051928380926367e0ed8360e11b82525afa90811561051757600091610568575b506001600160a01b0316838103610550575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa91821561051757600492602092600091610523575b506101005260405163f175240960e01b815292839182905afa908115610517576000916104c6575b506001600160a01b0390811660a081905260e05160c0516101005160408051938616845291909416602083015281019190915263ffffffff909116606082015260008051602061426383398151915290608090a180516001600160a01b0316156104b55761ffff825116801580156104aa575b610496578151600580546001600160a01b039283166001600160a01b03199091168117909155855160068054875161ffff60a01b60a09190911b169285166001600160b01b0319909116179190911790556040805191825286519092166020820152845161ffff16918101919091527f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197290606090a160405161397890816108cb823960805181505060a05181818161261f0152612d4a015260c05181818161164a0152612d1d015260e0518181816126da0152612ccc01526101005181612d710152f35b630c74dcaf60e41b60005260045260246000fd5b5061271081116103b2565b6342bcdf7f60e11b60005260046000fd5b90506020813d60201161050f575b816104e1602093836107f5565b8101031261050a5760806105036000805160206142638339815191529261083b565b915061033f565b600080fd5b3d91506104d4565b6040513d6000823e3d90fd5b6105439150833d8511610549575b61053b81836107f5565b81019061084f565b38610317565b503d610531565b836383395ca960e01b60005260045260245260446000fd5b6020813d6020116105a9575b81610581602093836107f5565b810103126105a55751906001600160a01b03821682036105a25750386102d0565b80fd5b5080fd5b3d9150610574565b6331b6aa1b60e11b600052600160045260245260446000fd5b6105e3915060203d6020116105495761053b81836107f5565b3861029e565b90506020813d60201161061b575b81610604602093836107f5565b8101031261050a576106159061083b565b38610265565b3d91506105f7565b633785f8f160e01b600052600160045260245260446000fd5b610655915060203d6020116105495761053b81836107f5565b38610235565b508215610207565b9050820151386101c4565b601f198216906003600052806000209160005b8181106106d35750836101fb936101ed96936000805160206142438339815191529896600194106106ba575b5050811b016003556101d9565b84015160001960f88460031b161c1916905538806106ad565b91926020600181928689015181550194019201610681565b60036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c81019160208410610746575b601f0160051c01905b81811061073a5750610194565b6000815560010161072d565b9091508190610724565b634e487b7160e01b600052604160045260246000fd5b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8183106107aa57505090602061017192820101610165565b6020919350806001915483858a01015201910190918592610792565b505060206101719260ff19851682840152151560051b820101610165565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761075057604052565b60005b83811061082b5750506000910152565b818101518382015260200161081b565b51906001600160a01b038216820361050a57565b9081602091031261050a575163ffffffff8116810361050a5790565b90600182811c9216801561089b575b602083101461088557565b634e487b7160e01b600052602260045260246000fd5b91607f169161087a565b906020916108be81518092818552858086019101610818565b601f01601f191601019056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612db05750806306285c6914612c65578063181f5a7714612be65780633bbbed4b1461236b5780635cb80c5d146121f75780636def4ce7146121225780637437ff9f1461202f57806379ba509714611f4a57806380485e2514611c825780638da5cb5b14611c30578063b2bd751c146118c0578063bff0ec1d14611128578063c9b146b314610d42578063ceac5cee1461098f578063d52e545a146106a9578063dfadfa35146105c4578063e023ddb114610368578063f2fde38b1461027b578063fe163eed146102225763fec888af146100fd57600080fd5b3461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760405190806003549061013d8261339a565b80855291600181169081156101da575060011461017d575b6101798461016581860382612f0d565b604051918291602083526020830190612f88565b0390f35b600381527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b8082106101c05750909150810160200161016582610155565b9192600181602092548385880101520191019092916101a7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b850190920192506101659150839050610155565b80fd5b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5773ffffffffffffffffffffffffffffffffffffffff6102c861300f565b6102d0613571565b1633811461034057807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461021f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f576040516103a481612eb9565b6103ac61300f565b815260243573ffffffffffffffffffffffffffffffffffffffff811681036105c057602082019081526044359061ffff821682036105bc57604083019182526103f3613571565b73ffffffffffffffffffffffffffffffffffffffff835116156105945761ffff82511680158015610589575b61055e5750916105589173ffffffffffffffffffffffffffffffffffffffff7f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197294818451167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655517fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000806006549360a01b161691161760065560405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390a180f35b7fc74dcaf0000000000000000000000000000000000000000000000000000000008552600452602484fd5b50612710811161041f565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b8280fd5b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f57604060a09167ffffffffffffffff61060a6130b2565b826080855161061881612ed5565b828152826020820152828782015282606082015201521681526004602052206040519061064482612ed5565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760043567ffffffffffffffff811161098b576106f99036906004016131b5565b610704929192613571565b81925b818410156108d35760c0840281019360c0853603126105bc576040519461072d86612ef1565b8035908187526020810135906020880197828952604081019060408301358252610759606084016130c9565b93606082019480865261078160a0610773608088016131e6565b96608086019788520161333a565b9660a08401978852159182156108ca575b5081156108b7575b50610851579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d51945192519351169751151596604051946107d386612ed5565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b1692161717905501929190610707565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff915016153861079a565b15915038610792565b6040519180602084016020855252604083019190845b81811061091a57857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff610950606089016130c9565b16606082015263ffffffff610967608089016131e6565b16608082015261097960a0880161333a565b151560a08201520194019291016108e9565b5080fd5b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760043567ffffffffffffffff811161098b573660238201121561098b576109f0903690602481600401359101613128565b6109f8613571565b60405190828260035491610a0b8361339a565b8083529260018116908115610d055750600114610ca7575b610a2f92500383612f0d565b805167ffffffffffffffff8111610c7a57610a4b60035461339a565b601f8111610bd9575b506020601f8211600114610af457610adb92827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd89593610558938891610ae9575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b604051938493604085526040850190612f88565b908382036020850152612f88565b905082015138610a95565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b818110610bc157508361055893610adb96937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610b8a575b5050811b01600355610ac7565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610b7d565b91926020600181928689015181550194019201610b3f565b60038552601f820160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208310610c52575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01905b818110610c475750610a54565b858155600101610c3a565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9150610c10565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50600385529084907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b818310610ce9575050906020610a2f92820101610a23565b6020919350806001915483858901015201910190918492610cd1565b60209250610a2f9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101610a23565b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760043567ffffffffffffffff811161098b57610d92903690600401613081565b73ffffffffffffffffffffffffffffffffffffffff6001541633036110e0575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b848110156110dc578060051b820135838112156110d8578201916080833603126110d85760405194610e0e86612e6e565b610e17846130c9565b8652610e256020850161333a565b9660208701978852604085013567ffffffffffffffff81116105c057610e4e90369087016135ff565b9460408801958652606081013567ffffffffffffffff81116105bc57610e76913691016135ff565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610ef0818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610fb0575b5095976001019550815b85518051821015610f415790610f3a73ffffffffffffffffffffffffffffffffffffffff610f32836001956135bc565b51168961370d565b5001610f02565b50509590969450600192919351908151610f61575b505001939293610ddd565b610fa667ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586925116926040519182916020835260208301906130de565b0390a23880610f56565b989395929691909497986000146110a157600184019591875b8651805182101561104657610ff38273ffffffffffffffffffffffffffffffffffffffff926135bc565b5116801561100f57906110086001928a61367c565b5001610fc9565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161109767ffffffffffffffff8a511692516040519182916020835260208301906130de565b0390a23880610ef8565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8480fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff60065416330315610db2576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b503461021f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f576004359067ffffffffffffffff821161021f576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261021f57604051916101c0830183811067ffffffffffffffff821117611893576040526111c3816004016130c9565b83526111d1602482016130c9565b60208401526111e2604482016130c9565b60408401526111f3606482016131e6565b6060840152611204608482016131e6565b608084015261121560a482016131a6565b60a084015260c481013560c084015260e481013567ffffffffffffffff81116105c057611248906004369184010161315f565b60e084015261010481013567ffffffffffffffff81116105c057611272906004369184010161315f565b61010084015261012481013567ffffffffffffffff81116105c05761129d906004369184010161315f565b61012084015261014481013567ffffffffffffffff81116105c0576112c8906004369184010161315f565b61014084015261016481013567ffffffffffffffff81116105c0576112f3906004369184010161315f565b61016084015261018481013567ffffffffffffffff81116105c0578101366023820112156105c05760048101359061132a8261317d565b916113386040519384612f0d565b80835260051b81016024016020830136821161188f5760248301905b82821061185957505050506101808401526101a481013567ffffffffffffffff81116105c0576113899136910160040161315f565b6101a083015260243560443567ffffffffffffffff81116105c0576113b2903690600401613053565b6101e1811061183157806004116105bc577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d0000000000000000000000000000000000000000000000000000000081036117e257506101809261017c600083861161021f57507fffffffff000000000000000000000000000000000000000000000000000000009084810135808316918703600481106117cd575b5050167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361177e57506000906101a0948581116105c0578386116105c0576114a490808703908601613347565b9081810361175057505067ffffffffffffffff86511685526004602052604085209560ff600388015460201c161561171a575061011c9050600082821161021f5750600190611518907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc8501613347565b950154948581036116ea57508394508083116116e557836102047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f6020969487956004995087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6082019061019c6040519b8c9a8b997f57ecfd28000000000000000000000000000000000000000000000000000000008b526040838c0152508260448b01520160648901378b6102008801528186880161020060248a0152526102248701378986857ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608489010101015201168201010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156116da5782916116ab575b50156116835780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b6116cd915060203d6020116116d3575b6116c58183612f0d565b810190613382565b3861167a565b503d6116bb565b6040513d84823e3d90fd5b505050fd5b84604491877f7d8b101a000000000000000000000000000000000000000000000000000000008352600452602452fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff16600452602485fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b839250829060040360031b1b16163880611455565b7fadaf77390000000000000000000000000000000000000000000000000000000085527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604484fd5b6004847f1ede477b000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff811161188b57602091611880839283600436928a0101016131f7565b815201910190611354565b8780fd5b8580fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760043567ffffffffffffffff811161098b576119109036906004016131b5565b611918613571565b6119218161317d565b9161192f6040519384612f0d565b81835260c06020840192028101903682116110d857915b818310611b8f578480855b8051831015611b8b5761196483826135bc565b519267ffffffffffffffff602061197b83856135bc565b51015116938415611b5f578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff00000000000000000000000000000000000000001617825560808101805163ffffffff1615611b335760019495969260ff73ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040945184547fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff000000000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a086015160d01b169360b01b1691161717809455511691835192835260f01c1615156020820152a2019190611951565b602486887f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c0833603126110d85760405190611ba682612ef1565b83359073ffffffffffffffffffffffffffffffffffffffff82168203611c2c578260209260c09452611bd98387016130c9565b83820152611be96040870161333a565b6040820152611bfa606087016131a6565b6060820152611c0b608087016131e6565b6080820152611c1c60a087016131e6565b60a0820152815201920191611946565b8680fd5b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461021f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f57611cba6130b2565b60243567ffffffffffffffff81116105c05760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105c05760405190611d0582612ed5565b806004013567ffffffffffffffff81116110d857611d29906004369184010161315f565b8252602481013567ffffffffffffffff81116110d857611d4f906004369184010161315f565b6020830152604481013567ffffffffffffffff81116110d8578101366023820112156110d8576004810135611d838161317d565b91611d916040519384612f0d565b818352602060048185019360061b830101019036821161188b57602401915b818310611ed2575050506040830152611dcb60648201613032565b6060830152608481013567ffffffffffffffff81116110d8576080916004611df6923692010161315f565b9101526044359067ffffffffffffffff82116105c057611e2367ffffffffffffffff92369060040161315f565b50611e2c613195565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff60408220541615611ea757816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b60408336031261188b576040516040810181811067ffffffffffffffff821117611f1d57916020916040938452611f0886613032565b81528286013583820152815201920191611db0565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f57805473ffffffffffffffffffffffffffffffffffffffff81163303612007577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f576040805161206b81612eb9565b828152826020820152015261017960405161208581612eb9565b73ffffffffffffffffffffffffffffffffffffffff60055416815261ffff60065473ffffffffffffffffffffffffffffffffffffffff8116602084015260a01c16604082015260405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5767ffffffffffffffff6121636130b2565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b8181106121e15773ffffffffffffffffffffffffffffffffffffffff85610179886121b981890382612f0d565b604051938360ff869560f01c16151585521660208401526060604084015260608301906130de565b825484526020909301926001928301920161218c565b503461021f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5760043567ffffffffffffffff811161098b57612247903690600401613081565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b83811015612367578060051b8201359073ffffffffffffffffffffffffffffffffffffffff821680920361188f57604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa801561235c5785908890612328575b60019450806122ea575b50505001612264565b6020816123197f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385876133ed565b604051908152a33884816122e1565b50506020833d8211612354575b8161234260209383612f0d565b81010312611c2c5784600193516122d7565b3d9150612335565b6040513d89823e3d90fd5b8480f35b503461021f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f576004359067ffffffffffffffff821161021f578160040190823603926101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc85011261098b576123ea612fe7565b5060843567ffffffffffffffff81116105c05761240b903690600401613053565b5050602481019361241b856132d1565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd61012484013591018112156105bc57820160048101359067ffffffffffffffff82116110d85760240181360381136110d8579067ffffffffffffffff91357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612bb1575b505060601c91168084526002602052604084209081549073ffffffffffffffffffffffffffffffffffffffff8216908115612b86576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612b7b578690612b18575b73ffffffffffffffffffffffffffffffffffffffff9150163303612aec5760f01c60ff16612aa4575b505067ffffffffffffffff61255d856132d1565b16825260046020526040822092600384019460ff865460201c1615612a665750610184820190600161258f83836132e6565b905003612a3057906125a0916132e6565b15612a035780357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41823603018112156105bc576125df913691016131f7565b9260408401519081519460208301517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811696601481106129cc575b50507f00000000000000000000000000000000000000000000000000000000000000009273ffffffffffffffffffffffffffffffffffffffff8416809760601c0361298e5750608081018051602081511161294c5750600283015480156128b15760a49150945b013561ffff811680910361188f576040519061269d82612eb9565b602435825280602083015260408201926107d084528791612822575b63ffffffff809161270473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001698898351916133ed565b519a54169454935116915192604051937f8e1d1a9d00000000000000000000000000000000000000000000000000000000602086015260248501526024845261274e604485612f0d565b853b1561188b57928798999263ffffffff6127c9938a966040519d8e9b8c9a8b997fabbce439000000000000000000000000000000000000000000000000000000008b5260048b015260248a01526044890152606488015260848701521660a485015260c484015261010060e4840152610104830190612f88565b03925af1918215612815578161017993612805575b5050604051906127ef602083612f0d565b8152604051918291602083526020830190612f88565b61280e91612f0d565b38816127de565b50604051903d90823e3d90fd5b90506103e88352805161ffff60065460a01c16908181029181830414901517156128845761271090049063ffffffff8211156126b957602488837fb6f15d0f000000000000000000000000000000000000000000000000000000008252600452fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b505160208151910151906020811061291b575b8060031b90808204600814901517156128ee576101000361010081116128ee571c9360a490612682565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b16906128c4565b61298a906040519182917fa3c8cf09000000000000000000000000000000000000000000000000000000008352602060048401526024830190612f88565b0390fd5b61298a906040519182917f22d4cfe2000000000000000000000000000000000000000000000000000000008352602060048401526024830190612f88565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929750829060140360031b1b161694388061261b565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b612a3d60249285926132e6565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8367ffffffffffffffff612a7b6024936132d1565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612ac15780612549565b7fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612b73575b81612b3260209383612f0d565b8101031261188f575173ffffffffffffffffffffffffffffffffffffffff8116810361188f5773ffffffffffffffffffffffffffffffffffffffff90612520565b3d9150612b25565b6040513d88823e3d90fd5b7f8a4e93c9000000000000000000000000000000000000000000000000000000008752600452602486fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b161638806124a8565b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f5750610179604051612c27604082612f0d565b601681527f43435450566572696669657220312e372e302d646576000000000000000000006020820152604051918291602083526020830190612f88565b503461021f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021f576060604051612ca281612e6e565b8281528260208201528260408201520152608073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff604051612cfd81612e6e565b82815273ffffffffffffffffffffffffffffffffffffffff6020820191817f00000000000000000000000000000000000000000000000000000000000000001683528160606040830192827f00000000000000000000000000000000000000000000000000000000000000001684520193857f0000000000000000000000000000000000000000000000000000000000000000168552604051968752511660208601525116604084015251166060820152f35b90503461098b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261098b576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036105c057602092507ffacbd7dc000000000000000000000000000000000000000000000000000000008114908115612e44575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612e3d565b6080810190811067ffffffffffffffff821117612e8a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff821117612e8a57604052565b60a0810190811067ffffffffffffffff821117612e8a57604052565b60c0810190811067ffffffffffffffff821117612e8a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612e8a57604052565b67ffffffffffffffff8111612e8a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110612fd25750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612f93565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361300a57565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361300a57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361300a57565b9181601f8401121561300a5782359167ffffffffffffffff831161300a576020838186019501011161300a57565b9181601f8401121561300a5782359167ffffffffffffffff831161300a576020808501948460051b01011161300a57565b6004359067ffffffffffffffff8216820361300a57565b359067ffffffffffffffff8216820361300a57565b906020808351928381520192019060005b8181106130fc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016130ef565b92919261313482612f4e565b916131426040519384612f0d565b82948184528183011161300a578281602093846000960137010152565b9080601f8301121561300a5781602061317a93359101613128565b90565b67ffffffffffffffff8111612e8a5760051b60200190565b6064359061ffff8216820361300a57565b359061ffff8216820361300a57565b9181601f8401121561300a5782359167ffffffffffffffff831161300a5760208085019460c0850201011161300a57565b359063ffffffff8216820361300a57565b919060c08382031261300a576040519061321082612ef1565b819380358352602081013567ffffffffffffffff811161300a578261323691830161315f565b6020840152604081013567ffffffffffffffff811161300a578261325b91830161315f565b6040840152606081013567ffffffffffffffff811161300a578261328091830161315f565b6060840152608081013567ffffffffffffffff811161300a57826132a591830161315f565b608084015260a08101359167ffffffffffffffff831161300a5760a0926132cc920161315f565b910152565b3567ffffffffffffffff8116810361300a5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561300a570180359067ffffffffffffffff821161300a57602001918160051b3603831361300a57565b3590811515820361300a57565b359060208110613355575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9081602091031261300a5751801515810361300a5790565b90600182811c921680156133e3575b60208310146133b457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916133a9565b9173ffffffffffffffffffffffffffffffffffffffff6134c09293604051928260208501967fa9059cbb00000000000000000000000000000000000000000000000000000000885216602485015260448401526044835261344f606484612f0d565b166000806040958651946134638887612f0d565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15613569573d916134a483612f4e565b926134b187519485612f0d565b83523d6000602085013e61389f565b805190816134cd57505050565b6020806134de938301019101613382565b156134e65750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609161389f565b73ffffffffffffffffffffffffffffffffffffffff60015416330361359257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156135d05760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561300a5781356136168161317d565b926136246040519485612f0d565b81845260208085019260051b82010192831161300a57602001905b82821061364c5750505090565b6020809161365984613032565b81520191019061363f565b80548210156135d05760005260206000200190600090565b60008281526001820160205260409020546137065780549068010000000000000000821015612e8a57826136ef6136ba846001809601855584613664565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014613896577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613867578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161386757818103613830575b50505080548015613801577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906137c28282613664565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6138506138406136ba9386613664565b90549060031b1c92839286613664565b90556000528360205260406000205538808061378a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b9192901561391a57508151156138b3575090565b3b156138bc5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561392d5750805190602001fd5b61298a906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612f8856fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8a87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c",
}

var CCTPVerifierABI = CCTPVerifierMetaData.ABI

var CCTPVerifierBin = CCTPVerifierMetaData.Bin

func DeployCCTPVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessengerProxy common.Address, messageTransmitterProxy common.Address, storageLocation string, dynamicConfig CCTPVerifierDynamicConfig) (common.Address, *types.Transaction, *CCTPVerifier, error) {
	parsed, err := CCTPVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPVerifierBin), backend, tokenMessengerProxy, messageTransmitterProxy, storageLocation, dynamicConfig)
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
	TokenMessengerProxy     common.Address
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

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) (*types.Transaction, error)

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
