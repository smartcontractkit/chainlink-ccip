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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.Domain\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.StaticConfig\",\"components\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.SetDomainArgs[]\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CCTPVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messageTransmitterProxy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localDomainIdentifier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"got\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidFastFinalityBps\",\"inputs\":[{\"name\":\"fastFinalityBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageSender\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterOnProxy\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageTransmitterVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSetDomainArgs\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct CCTPVerifier.SetDomainArgs\",\"components\":[{\"name\":\"allowedCallerOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCallerOnSource\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipientOnDest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxFeeExceedsUint32\",\"inputs\":[{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiveMessageCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFinality\",\"inputs\":[{\"name\":\"finality\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461062957614471803803809161001d8285610966565b83398101908082039160e083126106295781516001600160a01b038116908181036106295760208401516001600160a01b038116908181036106295760408601516001600160a01b038116979093908885036106295760608801516001600160401b0381116106295788019087601f83011215610629578151916100a083610989565b986100ae6040519a8b610966565b838a5260208483010111610629576060926100cf916020808c0191016109a4565b607f1901126106295760405194606086016001600160401b038111878210176108c157604052610101608089016109c7565b865260c061011160a08a016109c7565b9860208801998a5201519661ffff881688036106295760408701978852331561095557600180546001600160a01b0319163317905560405160035490919060008361015b836109f7565b80825291600184168015610937576001146108d7575b61017d92500384610966565b8151906001600160401b0382116108c157610197906109f7565b601f811161085c575b506020601f82116001146107df576101f992826000805160206144518339815191529593610207936000916107d4575b508160011b916000199060031b1c1916176003555b604051938493604085526040850190610a31565b908382036020850152610a31565b0390a1801580156107cc575b80156107c4575b61059c57604051639cdbb18160e01b8152602081600481855afa801561067e5763ffffffff916000916107a5575b50166001810361078c5750602060049160405192838092632c12192160e01b82525afa90811561067e57600091610752575b5060405163054fd4d560e41b81526001600160a01b03919091169390602081600481885afa801561067e5763ffffffff91600091610733575b50166001810361071a57506020600491604051928380926367e0ed8360e11b82525afa90811561067e576000916106d1575b506001600160a01b03168381036106b9575060e05260c05260405163234d8e3d60e21b815290602090829060049082905afa90811561067e5760009161068a575b506101005260a05260e051604051636eb1769f60e11b81523060048201526001600160a01b03909116602482018190529490602081604481855afa90811561067e5760009161064c575b5060001981018091116106365761041d9160405191602083019763095ea7b360e01b895260248401526044830152604482526103ad606483610966565b6000806040988951946103c08b87610966565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561062e573d9161040183610989565b9261040e8a519485610966565b83523d6000602085013e610a56565b8051806105ad575b505060e05160c05160a0516101005187516001600160a01b039485168152928416602084015292168187015263ffffffff90911660608201527fa87b8269841db42720476aa066520b7a92a17a6182da92012f7c52b7dd9ba25c90608090a180516001600160a01b03161561059c5761ffff82511680158015610591575b61057d575051600580546001600160a01b039283166001600160a01b03199091168117909155835160068054855161ffff60a01b60a09190911b169285166001600160b01b031990911617919091179055845190815292511660208301525161ffff16818301527f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197290606090a15161395d9081610af4823960805181505060a0518181816126cf0152612dcf015260c05181818161157b0152612da2015260e05181818161276d0152612d5101526101005181612df60152f35b630c74dcaf60e41b60005260045260246000fd5b5061271081116104a3565b6342bcdf7f60e11b60005260046000fd5b81602091810103126106295760200151801590811503610629576105d2573880610425565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b606091610a56565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610676575b8161066760209383610966565b81010312610629575138610370565b3d915061065a565b6040513d6000823e3d90fd5b6106ac915060203d6020116106b2575b6106a48183610966565b8101906109db565b38610326565b503d61069a565b836383395ca960e01b60005260045260245260446000fd5b6020813d602011610712575b816106ea60209383610966565b8101031261070e5751906001600160a01b038216820361070b5750386102e5565b80fd5b5080fd5b3d91506106dd565b6331b6aa1b60e11b600052600160045260245260446000fd5b61074c915060203d6020116106b2576106a48183610966565b386102b3565b90506020813d602011610784575b8161076d60209383610966565b810103126106295761077e906109c7565b3861027a565b3d9150610760565b633785f8f160e01b600052600160045260245260446000fd5b6107be915060203d6020116106b2576106a48183610966565b38610248565b50881561021a565b508315610213565b9050820151386101d0565b601f198216906003600052806000209160005b818110610844575083610207936101f9969360008051602061445183398151915298966001941061082b575b5050811b016003556101e5565b84015160001960f88460031b161c19169055388061081e565b919260206001819286890151815501940192016107f2565b60036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c810191602084106108b7575b601f0160051c01905b8181106108ab57506101a0565b6000815560010161089e565b9091508190610895565b634e487b7160e01b600052604160045260246000fd5b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b81831061091b57505090602061017d92820101610171565b6020919350806001915483858a01015201910190918592610903565b5050602061017d9260ff19851682840152151560051b820101610171565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176108c157604052565b6001600160401b0381116108c157601f01601f191660200190565b60005b8381106109b75750506000910152565b81810151838201526020016109a7565b51906001600160a01b038216820361062957565b90816020910312610629575163ffffffff811681036106295790565b90600182811c92168015610a27575b6020831014610a1157565b634e487b7160e01b600052602260045260246000fd5b91607f1691610a06565b90602091610a4a815180928185528580860191016109a4565b601f01601f1916010190565b91929015610ab85750815115610a6a575090565b3b15610a735790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610acb5750805190602001fd5b60405162461bcd60e51b815260206004820152908190610aef906024830190610a31565b0390fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612e355750806306285c6914612cea578063181f5a7714612c6b5780633bbbed4b146124095780635cb80c5d1461212c5780636def4ce7146120575780637437ff9f14611f6057806379ba509714611e7b57806380485e2514611bb35780638da5cb5b14611b61578063b2bd751c146117f1578063bff0ec1d14611059578063c9b146b314610c73578063ceac5cee14610989578063d52e545a146106a3578063dfadfa35146105be578063e023ddb114610362578063f2fde38b14610275578063fe163eed1461021c5763fec888af146100fd57600080fd5b3461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957604080516101398282612f92565b600181527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201835b8181106102065750506101736134c2565b61017c8261341f565b526101868161341f565b50815192839260208401906020855283518092526020818601918360051b870101940192905b8282106101bb57505050500390f35b919360206101f6827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652885161300d565b96019201920185949391926101ac565b6060602082850181019190915201610162565b80fd5b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760206040517f8e1d1a9d000000000000000000000000000000000000000000000000000000008152f35b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195773ffffffffffffffffffffffffffffffffffffffff6102c2613094565b6102ca613599565b1633811461033a57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346102195760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760405161039e81612f3e565b6103a6613094565b815260243573ffffffffffffffffffffffffffffffffffffffff811681036105ba57602082019081526044359061ffff821682036105b657604083019182526103ed613599565b73ffffffffffffffffffffffffffffffffffffffff8351161561058e5761ffff82511680158015610583575b6105585750916105529173ffffffffffffffffffffffffffffffffffffffff7f9cb852253eef65e3ff88bbfd63deac06075b7dea8d990e9ae4a51c094734197294818451167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655517fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000806006549360a01b161691161760065560405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390a180f35b7fc74dcaf0000000000000000000000000000000000000000000000000000000008552600452602484fd5b506127108111610419565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b8280fd5b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957604060a09167ffffffffffffffff610604613137565b826080855161061281612f5a565b828152826020820152828782015282606082015201521681526004602052206040519061063e82612f5a565b63ffffffff8154928381526001830154926020820193845260036002820154916040840192835201549360ff608060608501948688168652019560201c1615158552604051958652516020860152516040850152511660608301525115156080820152f35b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760043567ffffffffffffffff8111610985576106f390369060040161323a565b6106fe929192613599565b81925b818410156108cd5760c0840281019360c0853603126105b6576040519461072786612f76565b80359081875260208101359060208801978289526040810190604083013582526107536060840161314e565b93606082019480865261077b60a061076d6080880161326b565b9660808601978852016133bf565b9660a08401978852159182156108c4575b5081156108b1575b5061084b579863ffffffff9260039267ffffffffffffffff8560019a9b9c9d51945192519351169751151596604051946107cd86612f5a565b85526020850192835260408501938452606085019889526080850197885251168c52600460205260408c20925183555188830155516002820155019251167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000084549351151560201b1692161717905501929190610701565b8463ffffffff8467ffffffffffffffff8760c4968f604051977f3c0f3232000000000000000000000000000000000000000000000000000000008952516004890152516024880152516044870152511660648501525116608483015251151560a4820152fd5b67ffffffffffffffff9150161538610794565b1591503861078c565b6040519180602084016020855252604083019190845b81811061091457857f4b8db73ca99bc17f5741d6b1dcae3396bfcaad3ecf3742785f459ef163a3004686860387a180f35b90919260c08060019286358152602087013560208201526040870135604082015267ffffffffffffffff61094a6060890161314e565b16606082015263ffffffff6109616080890161326b565b16608082015261097360a088016133bf565b151560a08201520194019291016108e3565b5080fd5b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760043567ffffffffffffffff81116109855736602382011215610985576109ea9036906024816004013591016131ad565b6109f2613599565b6109fa6134c2565b90805167ffffffffffffffff8111610c4657610a1760035461346f565b601f8111610ba5575b506020601f8211600114610ac057610aa792827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd89593610552938891610ab5575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176003555b60405193849360408552604085019061300d565b90838203602085015261300d565b905082015138610a61565b600385527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316865b818110610b8d57508361055293610aa796937fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8989660019410610b56575b5050811b01600355610a93565b8401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880610b49565b91926020600181928689015181550194019201610b0b565b60038552601f820160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208310610c1e575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01905b818110610c135750610a20565b858155600101610c06565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9150610bdc565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760043567ffffffffffffffff811161098557610cc3903690600401613106565b73ffffffffffffffffffffffffffffffffffffffff600154163303611011575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b8481101561100d578060051b82013583811215611009578201916080833603126110095760405194610d3f86612ef3565b610d488461314e565b8652610d56602085016133bf565b9660208701978852604085013567ffffffffffffffff81116105ba57610d7f90369087016135e4565b9460408801958652606081013567ffffffffffffffff81116105b657610da7913691016135e4565b946060880195865267ffffffffffffffff885116835260026020526040832098511515610e21818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610ee1575b5095976001019550815b85518051821015610e725790610e6b73ffffffffffffffffffffffffffffffffffffffff610e638360019561345b565b5116896136f2565b5001610e33565b50509590969450600192919351908151610e92575b505001939293610d0e565b610ed767ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190613163565b0390a23880610e87565b98939592969190949798600014610fd257600184019591875b86518051821015610f7757610f248273ffffffffffffffffffffffffffffffffffffffff9261345b565b51168015610f405790610f396001928a613661565b5001610efa565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281610fc867ffffffffffffffff8a51169251604051918291602083526020830190613163565b0390a23880610e29565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8480fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff60065416330315610ce3576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b50346102195760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219576004359067ffffffffffffffff8211610219576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261021957604051916101c0830183811067ffffffffffffffff8211176117c4576040526110f48160040161314e565b83526111026024820161314e565b60208401526111136044820161314e565b60408401526111246064820161326b565b60608401526111356084820161326b565b608084015261114660a4820161322b565b60a084015260c481013560c084015260e481013567ffffffffffffffff81116105ba5761117990600436918401016131e4565b60e084015261010481013567ffffffffffffffff81116105ba576111a390600436918401016131e4565b61010084015261012481013567ffffffffffffffff81116105ba576111ce90600436918401016131e4565b61012084015261014481013567ffffffffffffffff81116105ba576111f990600436918401016131e4565b61014084015261016481013567ffffffffffffffff81116105ba5761122490600436918401016131e4565b61016084015261018481013567ffffffffffffffff81116105ba578101366023820112156105ba5760048101359061125b82613202565b916112696040519384612f92565b80835260051b8101602401602083013682116117c05760248301905b82821061178a57505050506101808401526101a481013567ffffffffffffffff81116105ba576112ba913691016004016131e4565b6101a083015260243560443567ffffffffffffffff81116105ba576112e39036906004016130d8565b6101e1811061176257806004116105b6577fffffffff000000000000000000000000000000000000000000000000000000008235167f8e1d1a9d00000000000000000000000000000000000000000000000000000000810361171357506101809261017c600083861161021957507fffffffff000000000000000000000000000000000000000000000000000000009084810135808316918703600481106116fe575b5050167f8e1d1a9d0000000000000000000000000000000000000000000000000000000081036116af57506000906101a0948581116105ba578386116105ba576113d5908087039086016133cc565b9081810361168157505067ffffffffffffffff86511685526004602052604085209560ff600388015460201c161561164b575061011c905060008282116102195750600190611449907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff040160fc85016133cc565b9501549485810361161b575083945080831161161657836102047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f6020969487956004995087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6082019061019c6040519b8c9a8b997f57ecfd28000000000000000000000000000000000000000000000000000000008b526040838c0152508260448b01520160648901378b6102008801528186880161020060248a0152526102248701378986857ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608489010101015201168201010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561160b5782916115dc575b50156115b45780f35b807fbc40f5560000000000000000000000000000000000000000000000000000000060049252fd5b6115fe915060203d602011611604575b6115f68183612f92565b810190613407565b386115ab565b503d6115ec565b6040513d84823e3d90fd5b505050fd5b84604491877f7d8b101a000000000000000000000000000000000000000000000000000000008352600452602452fd5b517fd201c48a00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff16600452602485fd5b7f6c86fa3a000000000000000000000000000000000000000000000000000000008752600452602452604485fd5b7fadaf77390000000000000000000000000000000000000000000000000000000086527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604485fd5b839250829060040360031b1b16163880611386565b7fadaf77390000000000000000000000000000000000000000000000000000000085527f8e1d1a9d00000000000000000000000000000000000000000000000000000000600452602452604484fd5b6004847f1ede477b000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff81116117bc576020916117b1839283600436928a01010161327c565b815201910190611285565b8780fd5b8580fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760043567ffffffffffffffff81116109855761184190369060040161323a565b611849613599565b61185281613202565b916118606040519384612f92565b81835260c060208401920281019036821161100957915b818310611ac0578480855b8051831015611abc57611895838261345b565b519267ffffffffffffffff60206118ac838561345b565b51015116938415611a90578484526002602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff00000000000000000000000000000000000000001617825560808101805163ffffffff1615611a645760019495969260ff73ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040945184547fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff79ffffffff000000000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a086015160d01b169360b01b1691161717809455511691835192835260f01c1615156020820152a2019190611882565b602486887f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c0833603126110095760405190611ad782612f76565b83359073ffffffffffffffffffffffffffffffffffffffff82168203611b5d578260209260c09452611b0a83870161314e565b83820152611b1a604087016133bf565b6040820152611b2b6060870161322b565b6060820152611b3c6080870161326b565b6080820152611b4d60a0870161326b565b60a0820152815201920191611877565b8680fd5b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102195760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957611beb613137565b60243567ffffffffffffffff81116105ba5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105ba5760405190611c3682612f5a565b806004013567ffffffffffffffff811161100957611c5a90600436918401016131e4565b8252602481013567ffffffffffffffff811161100957611c8090600436918401016131e4565b6020830152604481013567ffffffffffffffff811161100957810136602382011215611009576004810135611cb481613202565b91611cc26040519384612f92565b818352602060048185019360061b83010101903682116117bc57602401915b818310611e03575050506040830152611cfc606482016130b7565b6060830152608481013567ffffffffffffffff8111611009576080916004611d2792369201016131e4565b9101526044359067ffffffffffffffff82116105ba57611d5467ffffffffffffffff9236906004016131e4565b50611d5d61321a565b501690818152600260205273ffffffffffffffffffffffffffffffffffffffff60408220541615611dd857816060928252600260205263ffffffff604061ffff8185205460a01c16938381526002602052828282205460b01c169381526002602052205460d01c169060405192835260208301526040820152f35b6024917f8a4e93c9000000000000000000000000000000000000000000000000000000008252600452fd5b6040833603126117bc576040516040810181811067ffffffffffffffff821117611e4e57916020916040938452611e39866130b7565b81528286013583820152815201920191611ce1565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021957805473ffffffffffffffffffffffffffffffffffffffff81163303611f38577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760408051611f9c81612f3e565b8281528260208201520152612053604051611fb681612f3e565b73ffffffffffffffffffffffffffffffffffffffff60055416815261ffff60065473ffffffffffffffffffffffffffffffffffffffff8116602084015260a01c16604082015260405191829182919091604061ffff81606084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152015116910152565b0390f35b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195767ffffffffffffffff612098613137565b168152600260205260408120600181549101604051928360208354918281520192825260208220915b8181106121165773ffffffffffffffffffffffffffffffffffffffff85612053886120ee81890382612f92565b604051938360ff869560f01c1615158552166020840152606060408401526060830190613163565b82548452602090930192600192830192016120c1565b50346102195760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195760043567ffffffffffffffff81116109855761217c903690600401613106565b9073ffffffffffffffffffffffffffffffffffffffff6005541690835b83811015612405578060051b8201359073ffffffffffffffffffffffffffffffffffffffff82168092036117c057604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa9283156123fa5787936123c7575b508261221d575b506001915001612199565b84604051936122d860208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152612268606482612f92565b8a8060409889519361227a8b86612f92565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d156123bf573d906122bb82612fd3565b916122c88a519384612f92565b82523d8d602084013e5b86613884565b805180612314575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338612212565b61232b929495969350602080918301019101613407565b1561233c57929190859038806122e0565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060906122d2565b9092506020813d82116123f2575b816123e260209383612f92565b81010312611b5d5751913861220b565b3d91506123d5565b6040513d89823e3d90fd5b8480f35b50346102195760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219576004359067ffffffffffffffff821161021957816004018236036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126105ba5761248661306c565b5060843567ffffffffffffffff81116105b6576124a79036906004016130d8565b505060248401906124b782613356565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd610124870135910181121561100957850160048101359067ffffffffffffffff82116117c05760240181360381136117c0579067ffffffffffffffff91357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612c36575b505060601c91168085526002602052604085209081549073ffffffffffffffffffffffffffffffffffffffff8216908115612c0b576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156123fa578790612ba8575b73ffffffffffffffffffffffffffffffffffffffff9150163303612b7c5760f01c60ff16612b34575b505067ffffffffffffffff6125f982613356565b1683526004602052604083209160038301549160ff8360201c1615612af65750610184850190600161262b838361336b565b905003612ac0579061263c9161336b565b15612a935780357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41823603018112156110095761267b9136910161327c565b91604083015180519060208101517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110612a5c575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016809260601c03612a1e575060808401805160208151116129dc5750600283015480156129415760a49150965b013561ffff81168091036117c0576040519061273682612f3e565b6024358252806020830152604082016107d0815287916128b2575b63ffffffff73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001697519554915116925193604051947f8e1d1a9d0000000000000000000000000000000000000000000000000000000060208701526024860152602485526127d5604486612f92565b873b156128ae5793889795929363ffffffff899561285594829a986040519e8f9c8d9b8c9a7fabbce439000000000000000000000000000000000000000000000000000000008c5260048c01521660248a01526044890152606488015260848701521660a485015260c484015261010060e484015261010483019061300d565b03925af19182156128a1578161205393612891575b50506040519061287b602083612f92565b815260405191829160208352602083019061300d565b61289a91612f92565b388161286a565b50604051903d90823e3d90fd5b8880fd5b90506103e88152855161ffff60065460a01c16908181029181830414901517156129145761271090049063ffffffff82111561275157602488837fb6f15d0f000000000000000000000000000000000000000000000000000000008252600452fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b50516020815191015190602081106129ab575b8060031b908082046008149015171561297e5761010003610100811161297e571c9560a49061271b565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b1690612954565b612a1a906040519182917fa3c8cf0900000000000000000000000000000000000000000000000000000000835260206004840152602483019061300d565b0390fd5b612a1a906040519182917f22d4cfe200000000000000000000000000000000000000000000000000000000835260206004840152602483019061300d565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b16169038806126b6565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b612acd602492869261336b565b7f40de710500000000000000000000000000000000000000000000000000000000835260045250fd5b8467ffffffffffffffff612b0b602493613356565b7fd201c48a00000000000000000000000000000000000000000000000000000000835216600452fd5b60008281526002909101602052604090205415612b5157806125e5565b7fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612c03575b81612bc260209383612f92565b81010312611b5d575173ffffffffffffffffffffffffffffffffffffffff81168103611b5d5773ffffffffffffffffffffffffffffffffffffffff906125bc565b3d9150612bb5565b7f8a4e93c9000000000000000000000000000000000000000000000000000000008852600452602487fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16163880612544565b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102195750612053604051612cac604082612f92565b601681527f43435450566572696669657220312e372e302d64657600000000000000000000602082015260405191829160208352602083019061300d565b503461021957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610219576060604051612d2781612ef3565b8281528260208201528260408201520152608073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff604051612d8281612ef3565b82815273ffffffffffffffffffffffffffffffffffffffff6020820191817f00000000000000000000000000000000000000000000000000000000000000001683528160606040830192827f00000000000000000000000000000000000000000000000000000000000000001684520193857f0000000000000000000000000000000000000000000000000000000000000000168552604051968752511660208601525116604084015251166060820152f35b9050346109855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610985576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036105ba57602092507ffacbd7dc000000000000000000000000000000000000000000000000000000008114908115612ec9575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612ec2565b6080810190811067ffffffffffffffff821117612f0f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff821117612f0f57604052565b60a0810190811067ffffffffffffffff821117612f0f57604052565b60c0810190811067ffffffffffffffff821117612f0f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612f0f57604052565b67ffffffffffffffff8111612f0f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106130575750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613018565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361308f57565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361308f57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361308f57565b9181601f8401121561308f5782359167ffffffffffffffff831161308f576020838186019501011161308f57565b9181601f8401121561308f5782359167ffffffffffffffff831161308f576020808501948460051b01011161308f57565b6004359067ffffffffffffffff8216820361308f57565b359067ffffffffffffffff8216820361308f57565b906020808351928381520192019060005b8181106131815750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613174565b9291926131b982612fd3565b916131c76040519384612f92565b82948184528183011161308f578281602093846000960137010152565b9080601f8301121561308f578160206131ff933591016131ad565b90565b67ffffffffffffffff8111612f0f5760051b60200190565b6064359061ffff8216820361308f57565b359061ffff8216820361308f57565b9181601f8401121561308f5782359167ffffffffffffffff831161308f5760208085019460c0850201011161308f57565b359063ffffffff8216820361308f57565b919060c08382031261308f576040519061329582612f76565b819380358352602081013567ffffffffffffffff811161308f57826132bb9183016131e4565b6020840152604081013567ffffffffffffffff811161308f57826132e09183016131e4565b6040840152606081013567ffffffffffffffff811161308f57826133059183016131e4565b6060840152608081013567ffffffffffffffff811161308f578261332a9183016131e4565b608084015260a08101359167ffffffffffffffff831161308f5760a09261335192016131e4565b910152565b3567ffffffffffffffff8116810361308f5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561308f570180359067ffffffffffffffff821161308f57602001918160051b3603831361308f57565b3590811515820361308f57565b3590602081106133da575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9081602091031261308f5751801515810361308f5790565b80511561342c5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561342c5760209160051b010190565b90600182811c921680156134b8575b602083101461348957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161347e565b60405190600082600354916134d68361346f565b808352926001811690811561355c57506001146134fc575b6134fa92500383612f92565b565b506003600090815290917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8183106135405750509060206134fa928201016134ee565b6020919350806001915483858901015201910190918492613528565b602092506134fa9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016134ee565b73ffffffffffffffffffffffffffffffffffffffff6001541633036135ba57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561308f5781356135fb81613202565b926136096040519485612f92565b81845260208085019260051b82010192831161308f57602001905b8282106136315750505090565b6020809161363e846130b7565b815201910190613624565b805482101561342c5760005260206000200190600090565b60008281526001820160205260409020546136eb5780549068010000000000000000821015612f0f57826136d461369f846001809601855584613649565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b906001820191816000528260205260406000205480151560001461387b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161384c578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161384c57818103613815575b505050805480156137e6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906137a78282613649565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61383561382561369f9386613649565b90549060031b1c92839286613649565b90556000528360205260406000205538808061376f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b919290156138ff5750815115613898575090565b3b156138a15790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156139125750805190602001fd5b612a1a906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061300d56fea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
}

var CCTPVerifierABI = CCTPVerifierMetaData.ABI

var CCTPVerifierBin = CCTPVerifierMetaData.Bin

func DeployCCTPVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, messageTransmitterProxy common.Address, usdcToken common.Address, storageLocation string, dynamicConfig CCTPVerifierDynamicConfig) (common.Address, *types.Transaction, *CCTPVerifier, error) {
	parsed, err := CCTPVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPVerifierBin), backend, tokenMessenger, messageTransmitterProxy, usdcToken, storageLocation, dynamicConfig)
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

func (_CCTPVerifier *CCTPVerifierCaller) GetStorageLocation(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CCTPVerifier.contract.Call(opts, &out, "getStorageLocation")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_CCTPVerifier *CCTPVerifierSession) GetStorageLocation() ([]string, error) {
	return _CCTPVerifier.Contract.GetStorageLocation(&_CCTPVerifier.CallOpts)
}

func (_CCTPVerifier *CCTPVerifierCallerSession) GetStorageLocation() ([]string, error) {
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

	GetStorageLocation(opts *bind.CallOpts) ([]string, error)

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
