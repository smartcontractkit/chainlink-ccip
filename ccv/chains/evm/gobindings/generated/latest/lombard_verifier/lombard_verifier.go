// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lombard_verifier

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

type LombardVerifierDynamicConfig struct {
	FeeAggregator common.Address
}

type LombardVerifierPath struct {
	AllowedCaller [32]byte
	LChainId      [32]byte
}

type LombardVerifierRemoteAdapterArgs struct {
	RemoteChainSelector uint64
	Token               common.Address
	RemoteAdapter       [32]byte
}

type LombardVerifierSupportedTokenArgs struct {
	LocalToken   common.Address
	LocalAdapter common.Address
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

var LombardVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"},{\"name\":\"storageLocation\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteAdapter\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePaths\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRemoteAdapters\",\"inputs\":[{\"name\":\"remoteAdapterArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct LombardVerifier.RemoteAdapterArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSupportedTokens\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokensToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct LombardVerifier.SupportedTokenArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteAdapterSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenRemoved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"messageMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustTransferTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteTokenOrAdapterMismatch\",\"inputs\":[{\"name\":\"bridgeToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAllowedCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]}]",
	Bin: "0x60c0806040523461066957614b87803803809161001c828561066e565b8339810190808203608081126106695760201361066957604051602081016001600160401b038111828210176104765760405261005882610691565b81526020820151926001600160a01b038416928385036106695760408101516001600160401b03811161066957810182601f820112156106695780519061009e826106a5565b936100ac604051958661066e565b82855260208086019360051b830101918183116106695760208101935b8385106105f157505050505060606100e19101610691565b906001549080516100f1836106a5565b926100ff604051948561066e565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b83821061054c5750505060005b8181106104b857505060005b8181106103295750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161019d61018f9260405193849360408552604085019061074b565b90838203602085015261074b565b0390a16001600160a01b031680156103185760805233156103075760038054336001600160a01b0319918216179091559051600b80546001600160a01b039092169190921681179091556040519081527ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f6090602090a180156102f65760206004916040519283809263353c26b760e01b82525afa9081156102ea576000916102a6575b5060ff166002810361028d575060a0526040516143ca90816107bd823960805181613279015260a051818181610cdc01528181611a51015281816127240152818161285201526134fe0152f35b63398bbe0560e11b600052600260045260245260446000fd5b6020813d6020116102e2575b816102bf6020938361066e565b810103126102de57519060ff821682036102db575060ff610240565b80fd5b5080fd5b3d91506102b2565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b639b15e16f60e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b82518110156104a25760208160051b84010151600154680100000000000000008110156104765780600161036092016001556106df565b91909161048c578051906001600160401b0382116104765761038283546106fa565b601f8111610439575b50602090601f83116001146103ce57600194939291600091836103c3575b5050600019600383901b1c191690841b1790555b01610149565b0151905038806103a9565b90601f1983169184600052816000209260005b818110610421575091600196959492918388959310610408575b505050811b0190556103bd565b015160001960f88460031b161c191690553880806103fb565b929360206001819287860151815501950193016103e1565b61046690846000526020600020601f850160051c8101916020861061046c575b601f0160051c0190610734565b3861038b565b9091508190610459565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60015480156105365760001901906104cf826106df565b92909261048c57826104e3600194546106fa565b90816104f4575b505082550161013d565b81601f60009311861461050b5750555b38806104ea565b8183526020832061052691601f0160051c8101908701610734565b8082528160208120915555610504565b634e487b7160e01b600052603160045260246000fd5b6040516000845461055c816106fa565b80845290600181169081156105ce5750600114610596575b50600192826105888594602094038261066e565b815201930191019091610130565b6000868152602081209092505b8183106105b857505081016020016001610574565b60018160209254838688010152019201916105a3565b60ff191660208581019190915291151560051b8401909101915060019050610574565b84516001600160401b0381116106695782019083603f83011215610669576020820151906001600160401b0382116104765760405161063a601f8401601f19166020018261066e565b82815260408484010186106106695761065e602094938594604086850191016106bc565b8152019401936100c9565b600080fd5b601f909101601f19168101906001600160401b0382119082101761047657604052565b51906001600160a01b038216820361066957565b6001600160401b0381116104765760051b60200190565b60005b8381106106cf5750506000910152565b81810151838201526020016106bf565b6001548110156104a257600160005260206000200190600090565b90600182811c9216801561072a575b602083101461071457565b634e487b7160e01b600052602260045260246000fd5b91607f1691610709565b81811061073f575050565b60008155600101610734565b9080602083519182815201916020808360051b8301019401926000915b83831061077757505050505090565b909192939460208080600193601f1986820301875289516107a3815180928185528580860191016106bc565b601f01601f19160101970195949190910192019061076856fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146101b7578063181f5a77146101b2578063240028e8146101ad5780632e45ca68146101a8578063384ff3b7146101a357806338ff8c381461019e5780633bbbed4b146101995780635cb80c5d14610194578063708e1f791461018f578063737037e81461018a5780637437ff9f1461018557806379ba50971461018057806380485e251461017b57806387ae929214610176578063898068fc146101715780638da5cb5b1461016c5780638e0b8718146101675780638f2aaea4146101625780639ba439311461015d578063bcb6d4f714610158578063bff0ec1d14610153578063c4bffe2b1461014e578063c9b146b314610149578063d3c7c2c714610144578063f2fde38b1461013f5763fe163eed1461013a57600080fd5b6122f0565b6121fc565b612170565b611df1565b611d2e565b611817565b611767565b61163f565b611570565b6114f5565b6114a3565b611403565b611225565b611014565b610e18565b610d9e565b610d00565b610c91565b610b26565b610740565b61064e565b6105bd565b61052e565b6104bb565b610411565b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610276576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361027657807f83adcde1000000000000000000000000000000000000000000000000000000006020921490811561024c575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610241565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176102c657604052565b61027b565b6020810190811067ffffffffffffffff8211176102c657604052565b60c0810190811067ffffffffffffffff8211176102c657604052565b6080810190811067ffffffffffffffff8211176102c657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102c657604052565b6040519061036f60a08361031f565b565b67ffffffffffffffff81116102c657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106103be5750506000910152565b81810151838201526020016103ae565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361040a815180928187528780880191016103ab565b0116010190565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765761048e6040805190610452818361031f565b601982527f4c6f6d62617264566572696669657220322e302e302d646576000000000000006020830152519182916020835260208301906103ce565b0390f35b73ffffffffffffffffffffffffffffffffffffffff81160361027657565b359061036f82610492565b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027657602061052473ffffffffffffffffffffffffffffffffffffffff60043561051081610492565b166000526005602052604060002054151590565b6040519015158152f35b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff8111610276573660238201121561027657806004013567ffffffffffffffff81116102765736602460c08302840101116102765760246105a99201612349565b005b67ffffffffffffffff81160361027657565b346102765760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760206106456004356105fd816105ab565b67ffffffffffffffff6024359161061383610492565b16600052600a835260406000209073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b54604051908152f35b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765767ffffffffffffffff600435610692816105ab565b600060206040516106a2816102aa565b828152015216600052600960205261048e60406000206001604051916106c7836102aa565b8054835201546020820152604051918291829190916020806040830194805184520151910152565b90816101c09103126102765790565b9181601f840112156102765782359167ffffffffffffffff8311610276576020838186019501011161027657565b90602061073d9281815201906103ce565b90565b346102765760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff81116102765761078f9036906004016106ef565b6024359061079e604435610492565b60843567ffffffffffffffff8111610276576107be9036906004016106fe565b505060208101803560006107d1826105ab565b6107da82613227565b6101808401906107ea8286612460565b905015610a80576107fa836105ab565b61012085019273ffffffffffffffffffffffffffffffffffffffff61082a61082286896124b4565b810190612505565b169061084a8167ffffffffffffffff166000526000602052604060002090565b8054909161088273ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811615610a4a576040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152602090829060249082905afa8015610a455773ffffffffffffffffffffffffffffffffffffffff918691610a16575b501633036109ea5760f01c60ff1661095b575b61048e61094f8989896109478a61094161093b6109358d87612460565b90612549565b93612456565b936124b4565b929091613408565b6040519182918261072c565b61099b61099f9160016109846108698673ffffffffffffffffffffffffffffffffffffffff1690565b910160019160005201602052604060002054151590565b1590565b6109a95780610918565b7fd0d2597600000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff16600452602490fd5b7f728fe07b00000000000000000000000000000000000000000000000000000000845233600452602484fd5b610a38915060203d602011610a3e575b610a30818361031f565b810190612d8a565b38610905565b503d610a26565b612d9f565b7f4d1aff7e00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff8216600452602486fd5b807f4f73dc4d0000000000000000000000000000000000000000000000000000000060049252fd5b9181601f840112156102765782359167ffffffffffffffff8311610276576020808501948460051b01011161027657565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610276576004359067ffffffffffffffff821161027657610b2291600401610aa8565b9091565b3461027657610b3436610ad9565b9073ffffffffffffffffffffffffffffffffffffffff600b5416918215610c675760005b818110610b6157005b610b77610869610b72838587612893565b6128a3565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa8015610a45576001948892600092610c37575b5081610beb575b5050505001610b58565b81610c1b7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610c2b94613e96565b6040519081529081906020820190565b0390a338858180610be1565b610c5991925060203d8111610c60575b610c51818361031f565b81019061333e565b9038610bda565b503d610c47565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102765760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff811161027657610d4f903690600401610aa8565b6024359167ffffffffffffffff831161027657366023840112156102765782600401359167ffffffffffffffff8311610276573660248460061b860101116102765760246105a9940191612587565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610276576000604051610ddb816102cb565b5261048e604051610deb816102cb565b600b5473ffffffffffffffffffffffffffffffffffffffff16908190526040519081529081906020820190565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760025473ffffffffffffffffffffffffffffffffffffffff81163303610ed7577fffffffffffffffffffffffff00000000000000000000000000000000000000006003549133828416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610f0d82610371565b91610f1b604051938461031f565b829481845281830111610276578281602093846000960137010152565b9080601f830112156102765781602061073d93359101610f01565b67ffffffffffffffff81116102c65760051b60200190565b81601f8201121561027657803590610f8282610f53565b92610f90604051948561031f565b82845260208085019360061b8301019181831161027657602001925b828410610fba575050505090565b6040848303126102765760206040918251610fd4816102aa565b8635610fdf81610492565b81528287013583820152815201930192610fac565b6064359061ffff8216820361027657565b359061ffff8216820361027657565b346102765760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043561104f816105ab565b60243567ffffffffffffffff81116102765760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261027657611095610360565b90806004013567ffffffffffffffff8111610276576110ba9060043691840101610f38565b8252602481013567ffffffffffffffff8111610276576110e09060043691840101610f38565b6020830152604481013567ffffffffffffffff8111610276576111099060043691840101610f6b565b604083015261111a606482016104b0565b6060830152608481013567ffffffffffffffff81116102765760809160046111459236920101610f38565b91015260443567ffffffffffffffff81116102765761048e9161116f61117e923690600401610f38565b50611178610ff4565b506128f6565b6040805161ffff909416845263ffffffff92831660208501529116908201529081906060820190565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106111da57505050505090565b9091929394602080611216837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516103ce565b970193019301919392906111cb565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760015461126081610f53565b61126d604051918261031f565b818152602081019160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66000935b8285106112b4576040518061048e86826111a7565b604051600083548060011c90600181169081156113af575b60208310821461138257828552602085019190811561134d5750600114611311575b5050600192826113038594602094038261031f565b81520192019401939061129f565b90915061132385600052602060002090565b916000925b8184106113395750500182826112ee565b600181602092548686015201930192611328565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682525090151560051b01905082826112ee565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f16916112cc565b906020808351928381520192019060005b8181106113d75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016113ca565b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765767ffffffffffffffff600435611447816105ab565b16600052600060205273ffffffffffffffffffffffffffffffffffffffff604060002061048e61147b600183549301614001565b604051938360ff869560f01c16151585521660208401526060604084015260608301906113b9565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027657602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff8111610276573660238201121561027657806004013567ffffffffffffffff81116102765736602460608302840101116102765760246105a992016129c1565b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610276577ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f6061163a6040516115cf816102cb565b6004356115db81610492565b81526115e5612e87565b51600b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691821790556040519081529081906020820190565b0390a1005b346102765760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043561167a816105ab565b60243590604435611689612e87565b821561173d5780156117135761170e7f83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9916040516116c6816102aa565b818152600167ffffffffffffffff60208301968888521695866000526009602052604060002092518355519101556116fd8461407f565b506040519081529081906020820190565b0390a3005b7f55622b8a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5a39e3030000000000000000000000000000000000000000000000000000000060005260046000fd5b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff81116102765736602382011215610276578060040135906117c282610f53565b916117d0604051938461031f565b8083526024602084019160051b8301019136831161027657602401905b8282106117fd576105a984612ab4565b60208091833561180c816105ab565b8152019101906117ed565b346102765760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760043567ffffffffffffffff8111610276576118669036906004016106ef565b6024359060443567ffffffffffffffff81116102765761188a9036906004016106fe565b9061189c61189784612456565b613227565b6118ad6118a884612456565b6139fc565b6118c06118ba8383612c0e565b90612c6f565b7feba55588000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603611b905750600690818310611cc05761193b61193461192e611928858786612c1c565b90612d24565b60f01c90565b61ffff1690565b9261194e6119498585612d17565b612d04565b8110611cc05761196d6119646119c29585612d17565b80948385612c57565b949095610180810161198f6119856109358385612460565b60608101906124b4565b6119b96109356119b16119a761093587899799612460565b60808101906124b4565b959094612460565b3593898b613bcd565b6119e86119e261193461192e6119286119da88612d04565b888789612c57565b93612d04565b906119f38483612d17565b8110611cc057611a06611a0c9483612d17565b92612c57565b92604051927fd5438eae00000000000000000000000000000000000000000000000000000000845260208460048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215610a45576000948593611c87575b5073ffffffffffffffffffffffffffffffffffffffff8591611ad5604051988997889687947fa620850600000000000000000000000000000000000000000000000000000000865260048601612e60565b0393165af1908115610a4557600090600092611c60575b5015611c36576024815103611c035760246020820151910151907feba55588000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603611b905750818103611b5957005b7f6c86fa3a0000000000000000000000000000000000000000000000000000000060005260049190915260245260446000fd5b6000fd5b7fadaf7739000000000000000000000000000000000000000000000000000000006000527feba55588000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b517fc2fdac9800000000000000000000000000000000000000000000000000000000600052602460048190525260446000fd5b7f2532cf450000000000000000000000000000000000000000000000000000000060005260046000fd5b9050611c7f91503d806000833e611c77818361031f565b810190612dab565b915038611aec565b85919350611cb873ffffffffffffffffffffffffffffffffffffffff9160203d602011610a3e57610a30818361031f565b939150611a84565b7f1ede477b0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b818110611d0e5750505090565b825167ffffffffffffffff16845260209384019390920191600101611d01565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027657600754611d6981610f53565b90611d77604051928361031f565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611da482610f53565b0136602084013760005b818110611dc3576040518061048e8582611cea565b8067ffffffffffffffff611dd860019361404c565b90549060031b1c16611dea8286612bbe565b5201611dae565b3461027657611dff36610ad9565b611e07612e87565b6000905b808210611e1457005b611e27611e22838386613cc4565b613d6b565b92611e57611e3d855167ffffffffffffffff1690565b67ffffffffffffffff166000526000602052604060002090565b92611e67845460ff9060f01c1690565b916020860192611e778451151590565b90811515901515036120c6575b506060860194600101939060005b86518051821015611f6c5790611ec7611ead82600194612bbe565b5173ffffffffffffffffffffffffffffffffffffffff1690565b611eef611ee973ffffffffffffffffffffffffffffffffffffffff8316610869565b896142dc565b611efb575b5001611e92565b7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82611f6367ffffffffffffffff611f3a8d5167ffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff909516855216929081906020820190565b0390a238611ef4565b50509450949190926040830191825151611f8f575b505050506001019091611e0b565b5192959194909392156120b15760005b8551805182101561209e57611ead82611fb792612bbe565b73ffffffffffffffffffffffffffffffffffffffff8116156120535760019190611fff611ff973ffffffffffffffffffffffffffffffffffffffff8316610869565b88614110565b61200b575b5001611f9f565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da8061204a67ffffffffffffffff611f3a8c5167ffffffffffffffff1690565b0390a238612004565b611b8c612068895167ffffffffffffffff1690565b7f463258ff0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b5050935093506001915090388080611f81565b611b8c612068875167ffffffffffffffff1690565b85547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560f01b7eff000000000000000000000000000000000000000000000000000000000000161786557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f852349261216767ffffffffffffffff6121538a5167ffffffffffffffff1690565b604051941515855216929081906020820190565b0390a238611e84565b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610276576040516004548082526020820190600460005260206000209060005b8181106121e65761048e856121d28187038261031f565b6040519182916020835260208301906113b9565b82548452602090930192600192830192016121bb565b346102765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765773ffffffffffffffffffffffffffffffffffffffff60043561224c81610492565b612254612e87565b163381146122c657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102765760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102765760206040517feba55588000000000000000000000000000000000000000000000000000000008152f35b90612352612e87565b61235b81610f53565b91612369604051938461031f565b81835260c060208401920281019036821161027657915b8183106123935750505061036f90612ed2565b60c08336031261027657602060c0916040516123ae816102e7565b85356123b981610492565b8152828601356123c8816105ab565b8382015260408601356123da8161241d565b60408201526123eb60608701611005565b60608201526123fc60808701612427565b608082015261240d60a08701612427565b60a0820152815201920191612380565b8015150361027657565b359063ffffffff8216820361027657565b90604051612445816102aa565b602060018294805484520154910152565b3561073d816105ab565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610276570180359067ffffffffffffffff821161027657602001918160051b3603831361027657565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610276570180359067ffffffffffffffff82116102765760200191813603831361027657565b90816020910312610276573561073d81610492565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015612582578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610276570190565b61251a565b90929192612593612e87565b60005b81811061274f5750505060005b8181106125af57505050565b807f086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e36125e66125e160019486886128ad565b6128bd565b6126bb6126a061260a835173ffffffffffffffffffffffffffffffffffffffff1690565b926126376020820194612631865173ffffffffffffffffffffffffffffffffffffffff1690565b906139c8565b50612659610869855173ffffffffffffffffffffffffffffffffffffffff1690565b156126e957611ead612682610869835173ffffffffffffffffffffffffffffffffffffffff1690565b855173ffffffffffffffffffffffffffffffffffffffff1690613946565b915173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152a1016125a3565b61274a61270d610869835173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690613946565b611ead565b80612760610b726001938587612893565b61277f73ffffffffffffffffffffffffffffffffffffffff8216610869565b61279761279161086961086984613ef8565b91613f5b565b6127a4575b505001612596565b7fbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a3069161282d9173ffffffffffffffffffffffffffffffffffffffff811615612837576128069073ffffffffffffffffffffffffffffffffffffffff83166137ec565b60405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a1388061279c565b5061288e73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83166137ec565b612806565b91908110156125825760051b0190565b3561073d81610492565b91908110156125825760061b0190565b604081360312610276576020604051916128d6836102aa565b80356128e181610492565b835201356128ee81610492565b602082015290565b9067ffffffffffffffff821680600052600060205273ffffffffffffffffffffffffffffffffffffffff604060002054161561299457600052600060205261ffff60406000205460a01c169063ffffffff612989816129698667ffffffffffffffff166000526000602052604060002090565b5460b01c169467ffffffffffffffff166000526000602052604060002090565b5460d01c1691929190565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6129c9612e87565b60005b82811015612aaf5760019060006060820284017ff2bb53a7e6aae800a85fba961b2bc3124a23dd44d95fefe0b7b29bd90975a976604082013591803593612a12856105ab565b83612a696020612a368867ffffffffffffffff16600052600a602052604060002090565b9401358094612a4482610492565b9073ffffffffffffffffffffffffffffffffffffffff16600052602052604060002090565b55612a73856105ab565b612a7c82610492565b5060405192835273ffffffffffffffffffffffffffffffffffffffff169267ffffffffffffffff1691602090a3016129cc565b505050565b90612abd612e87565b6000915b8051831015612bb95767ffffffffffffffff612add8483612bbe565b511692612b06612b018567ffffffffffffffff166000526009602052604060002090565b612438565b612b0f856141ff565b15612b805784612b46612b3a600195969767ffffffffffffffff166000526009602052604060002090565b60016000918281550155565b60208281015192516040519081527f8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b9190a3019190612ac1565b847fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff602491166004526000fd5b509050565b80518210156125825760209160051b010190565b91612c0a918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b906004116102765790600490565b909291928360041161027657831161027657600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610276578411610276578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612ca3575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9060028201809211612d1257565b612cd5565b91908201809211612d1257565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110612d58575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b90816020910312610276575161073d81610492565b6040513d6000823e3d90fd5b9091606082840312610276578151926020830151612dc88161241d565b9260408101519067ffffffffffffffff8211610276570181601f82011215610276578051612df581610371565b92612e03604051948561031f565b818452602082840101116102765761073d91602080850191016103ab565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9290612e799061073d9593604086526040860191612e21565b926020818503910152612e21565b73ffffffffffffffffffffffffffffffffffffffff600354163303612ea857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060005b8251811015612bb957612ee98184612bbe565b51612f096020612ef98487612bbe565b51015167ffffffffffffffff1690565b67ffffffffffffffff81169182156131da57612f398267ffffffffffffffff166000526000602052604060002090565b91612f9c612f5b835173ffffffffffffffffffffffffffffffffffffffff1690565b849073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b612ffc612fac6040840151151590565b84547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f01b7eff00000000000000000000000000000000000000000000000000000000000016178455565b61305561300e606084015161ffff1690565b84547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff000000000000000000000000000000000000000016178455565b608082019061307461306b835163ffffffff1690565b63ffffffff1690565b156131a35750916131746131696108697f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946131066130bb60019a99985163ffffffff1690565b86547fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1660b09190911b79ffffffff0000000000000000000000000000000000000000000016178655565b611ead61311a60a083015163ffffffff1690565b86547fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff1660d09190911b7dffffffff000000000000000000000000000000000000000000000000000016178655565b915460f01c60ff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff939093168352901515602083015290a201612ed6565b7f9e7205510000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f97ccaab70000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b90816020910312610276575161073d8161241d565b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa908115610a45576000916132fe575b506132c75750565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b613320915060203d602011613326575b613318818361031f565b810190613212565b386132bf565b503d61330e565b91602061073d938181520191612e21565b90816020910312610276575190565b90604051917feba5558800000000000000000000000000000000000000000000000000000000602084015260248301526024825261036f60448361031f565b9190826040910312610276576020825192015190565b939061073d97969373ffffffffffffffffffffffffffffffffffffffff60e09794819388521660208701521660408501526060840152608083015260a08201528160c082015201906103ce565b906040519160208301526020825261036f60408361031f565b929194936080840192602061341d85876124b4565b9050116137a85761343761086961082260408801886124b4565b9161346861099b73ffffffffffffffffffffffffffffffffffffffff85166000526005602052604060002054151590565b6137645761348d612b018567ffffffffffffffff166000526009602052604060002090565b9485511561372c57613580959697986134ca6108696108696134c56108698973ffffffffffffffffffffffffffffffffffffffff1690565b613ef8565b73ffffffffffffffffffffffffffffffffffffffff8116613724575b5073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001695602089019060208783516040519b8c9283927f6e48b60d0000000000000000000000000000000000000000000000000000000084526004840190929173ffffffffffffffffffffffffffffffffffffffff6020916040840195845216910152565b03818b5afa908115610a45578b99600092613703575b506135b66135b16135aa60608d018d6124b4565b3691610f01565b613df1565b9081830361367a575b505050926135f86135ee6135b16135aa61362d966135e860409e9760009b9a519a810190612505565b9c6124b4565b9a3592519161334d565b9189519a8b998a9889977e8a1198000000000000000000000000000000000000000000000000000000008952600489016133a2565b03925af18015610a455761073d91600091613649575b506133ef565b61366b915060403d604011613673575b613663818361031f565b81019061338c565b905038613643565b503d613659565b889a5097612a446136a8929394959697989a9c9967ffffffffffffffff16600052600a602052604060002090565b5491821580156136f9575b6136c6578b99979a9896959493926135bf565b7fbce7b6cd0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50828114156136b3565b61371d91925060203d602011610c6057610c51818361031f565b9038613596565b9450386134e6565b7fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff851660045260246000fd5b7f06439c6b0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b6137b284866124b4565b906137e86040519283927fa3c8cf090000000000000000000000000000000000000000000000000000000084526004840161332d565b0390fd5b60405190602060008184017f095ea7b300000000000000000000000000000000000000000000000000000000815261387b8561384f8489602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810187528661031f565b84519082855af1600051903d8161390d575b501590505b61389b57505050565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9390931660248401526000604480850191909152835261036f926139089061390260648261031f565b82613f76565b613f76565b1515905061393a575061389273ffffffffffffffffffffffffffffffffffffffff82163b15155b3861388d565b60016138929114613934565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff851660248401527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60448401529192919060009061387b856064810161384f565b9073ffffffffffffffffffffffffffffffffffffffff8061073d931691826000526006602052166040600020556004614110565b613a3a610869613a208367ffffffffffffffff166000526000602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811615613b04576040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082905afa908115610a4557600091613ae5575b5015613ab757565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b613afe915060203d60201161332657613318818361031f565b38613aaf565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b9160c083830312610276578235926020810135926040820135926060830135613b6481610492565b926080810135613b7381610492565b9260a082013567ffffffffffffffff81116102765761073d9201610f38565b359060208110613ba0575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b95949392919581600411610276576004613bea9282019101613b3c565b95945050505050613c0a6021830151916081606185015194015197613b92565b808203613c94575050613c1d8383613b92565b03613c5f575050808203613c2f575050565b7f7c83fcf00000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b6137e86040519283927fa3c8cf090000000000000000000000000000000000000000000000000000000084526004840161332d565b7fd27ededb0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91908110156125825760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301821215610276570190565b9080601f83011215610276578135613d1b81610f53565b92613d29604051948561031f565b81845260208085019260051b82010192831161027657602001905b828210613d515750505090565b602080918335613d6081610492565b815201910190613d44565b6080813603126102765760405190613d8282610303565b8035613d8d816105ab565b82526020810135613d9d8161241d565b6020830152604081013567ffffffffffffffff811161027657613dc39036908301613d04565b604083015260608101359067ffffffffffffffff821161027657613de991369101613d04565b606082015290565b6020815111613e6057602081519101519060208110613e2f575b8060031b9080820460081490151715612d1257610100036101008111612d12571c90565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b1690613e0b565b6137e8906040519182917fe0d7fb020000000000000000000000000000000000000000000000000000000083526004830161072c565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff92909216602483015260448083019390935291815261036f9161390860648361031f565b80600052600660205260406000205490811580613f45575b613f18575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5060008181526005602052604090205415613f10565b61073d908060005260066020526000604081205560046142dc565b906000602091828151910182855af115612d9f576000513d613ff8575073ffffffffffffffffffffffffffffffffffffffff81163b155b613fb45750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415613fad565b906040519182815491828252602082019060005260206000209260005b81811061403357505061036f9250038361031f565b845483526001948501948794506020909301920161401e565b60075481101561258257600760005260206000200190600090565b80548210156125825760005260206000200190600090565b60008181526008602052604090205461410a57600754680100000000000000008110156102c6576140f16140bc8260018594016007556007614067565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600754906000526008602052604060002055600190565b50600090565b600082815260018201602052604090205461416557805490680100000000000000008210156102c6578261414e6140bc846001809601855584614067565b905580549260005201602052604060002055600190565b5050600090565b805480156141d0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906141a18282614067565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260086020526040902054908115614165577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211612d1257600754927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612d1257838360009561429b95036142a1575b50505061428a600761416c565b600890600052602052604060002090565b55600190565b61428a6142cd916142c36142b96142d3956007614067565b90549060031b1c90565b9283916007614067565b90612bd2565b5538808061427d565b60018101918060005282602052604060002054928315156000146143b4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612d12578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501948511612d1257600095858361429b9761436c950361437b575b50505061416c565b90600052602052604060002090565b61439b6142cd916143926142b96143ab9588614067565b92839187614067565b8590600052602052604060002090565b55388080614364565b5050505060009056fea164736f6c634300081a000a",
}

var LombardVerifierABI = LombardVerifierMetaData.ABI

var LombardVerifierBin = LombardVerifierMetaData.Bin

func DeployLombardVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig LombardVerifierDynamicConfig, bridge common.Address, storageLocation []string, rmn common.Address) (common.Address, *types.Transaction, *LombardVerifier, error) {
	parsed, err := LombardVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardVerifierBin), backend, dynamicConfig, bridge, storageLocation, rmn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LombardVerifier{address: address, abi: *parsed, LombardVerifierCaller: LombardVerifierCaller{contract: contract}, LombardVerifierTransactor: LombardVerifierTransactor{contract: contract}, LombardVerifierFilterer: LombardVerifierFilterer{contract: contract}}, nil
}

type LombardVerifier struct {
	address common.Address
	abi     abi.ABI
	LombardVerifierCaller
	LombardVerifierTransactor
	LombardVerifierFilterer
}

type LombardVerifierCaller struct {
	contract *bind.BoundContract
}

type LombardVerifierTransactor struct {
	contract *bind.BoundContract
}

type LombardVerifierFilterer struct {
	contract *bind.BoundContract
}

type LombardVerifierSession struct {
	Contract     *LombardVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LombardVerifierCallerSession struct {
	Contract *LombardVerifierCaller
	CallOpts bind.CallOpts
}

type LombardVerifierTransactorSession struct {
	Contract     *LombardVerifierTransactor
	TransactOpts bind.TransactOpts
}

type LombardVerifierRaw struct {
	Contract *LombardVerifier
}

type LombardVerifierCallerRaw struct {
	Contract *LombardVerifierCaller
}

type LombardVerifierTransactorRaw struct {
	Contract *LombardVerifierTransactor
}

func NewLombardVerifier(address common.Address, backend bind.ContractBackend) (*LombardVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(LombardVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLombardVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LombardVerifier{address: address, abi: abi, LombardVerifierCaller: LombardVerifierCaller{contract: contract}, LombardVerifierTransactor: LombardVerifierTransactor{contract: contract}, LombardVerifierFilterer: LombardVerifierFilterer{contract: contract}}, nil
}

func NewLombardVerifierCaller(address common.Address, caller bind.ContractCaller) (*LombardVerifierCaller, error) {
	contract, err := bindLombardVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierCaller{contract: contract}, nil
}

func NewLombardVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*LombardVerifierTransactor, error) {
	contract, err := bindLombardVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierTransactor{contract: contract}, nil
}

func NewLombardVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*LombardVerifierFilterer, error) {
	contract, err := bindLombardVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierFilterer{contract: contract}, nil
}

func bindLombardVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LombardVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LombardVerifier *LombardVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardVerifier.Contract.LombardVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_LombardVerifier *LombardVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.Contract.LombardVerifierTransactor.contract.Transfer(opts)
}

func (_LombardVerifier *LombardVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardVerifier.Contract.LombardVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_LombardVerifier *LombardVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_LombardVerifier *LombardVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.Contract.contract.Transfer(opts)
}

func (_LombardVerifier *LombardVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_LombardVerifier *LombardVerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (LombardVerifierDynamicConfig, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(LombardVerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(LombardVerifierDynamicConfig)).(*LombardVerifierDynamicConfig)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetDynamicConfig() (LombardVerifierDynamicConfig, error) {
	return _LombardVerifier.Contract.GetDynamicConfig(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetDynamicConfig() (LombardVerifierDynamicConfig, error) {
	return _LombardVerifier.Contract.GetDynamicConfig(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_LombardVerifier *LombardVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _LombardVerifier.Contract.GetFee(&_LombardVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _LombardVerifier.Contract.GetFee(&_LombardVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_LombardVerifier *LombardVerifierCaller) GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardVerifierPath, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getPath", remoteChainSelector)

	if err != nil {
		return *new(LombardVerifierPath), err
	}

	out0 := *abi.ConvertType(out[0], new(LombardVerifierPath)).(*LombardVerifierPath)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetPath(remoteChainSelector uint64) (LombardVerifierPath, error) {
	return _LombardVerifier.Contract.GetPath(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetPath(remoteChainSelector uint64) (LombardVerifierPath, error) {
	return _LombardVerifier.Contract.GetPath(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCaller) GetRemoteAdapter(opts *bind.CallOpts, remoteChainSelector uint64, token common.Address) ([32]byte, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getRemoteAdapter", remoteChainSelector, token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetRemoteAdapter(remoteChainSelector uint64, token common.Address) ([32]byte, error) {
	return _LombardVerifier.Contract.GetRemoteAdapter(&_LombardVerifier.CallOpts, remoteChainSelector, token)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetRemoteAdapter(remoteChainSelector uint64, token common.Address) ([32]byte, error) {
	return _LombardVerifier.Contract.GetRemoteAdapter(&_LombardVerifier.CallOpts, remoteChainSelector, token)
}

func (_LombardVerifier *LombardVerifierCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_LombardVerifier *LombardVerifierSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _LombardVerifier.Contract.GetRemoteChainConfig(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _LombardVerifier.Contract.GetRemoteChainConfig(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCaller) GetStorageLocations(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getStorageLocations")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetStorageLocations() ([]string, error) {
	return _LombardVerifier.Contract.GetStorageLocations(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetStorageLocations() ([]string, error) {
	return _LombardVerifier.Contract.GetStorageLocations(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetSupportedChains() ([]uint64, error) {
	return _LombardVerifier.Contract.GetSupportedChains(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LombardVerifier.Contract.GetSupportedChains(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetSupportedTokens() ([]common.Address, error) {
	return _LombardVerifier.Contract.GetSupportedTokens(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _LombardVerifier.Contract.GetSupportedTokens(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) IBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "i_bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) IBridge() (common.Address, error) {
	return _LombardVerifier.Contract.IBridge(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) IBridge() (common.Address, error) {
	return _LombardVerifier.Contract.IBridge(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardVerifier.Contract.IsSupportedToken(&_LombardVerifier.CallOpts, token)
}

func (_LombardVerifier *LombardVerifierCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardVerifier.Contract.IsSupportedToken(&_LombardVerifier.CallOpts, token)
}

func (_LombardVerifier *LombardVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) Owner() (common.Address, error) {
	return _LombardVerifier.Contract.Owner(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) Owner() (common.Address, error) {
	return _LombardVerifier.Contract.Owner(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardVerifier.Contract.SupportsInterface(&_LombardVerifier.CallOpts, interfaceId)
}

func (_LombardVerifier *LombardVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardVerifier.Contract.SupportsInterface(&_LombardVerifier.CallOpts, interfaceId)
}

func (_LombardVerifier *LombardVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) TypeAndVersion() (string, error) {
	return _LombardVerifier.Contract.TypeAndVersion(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) TypeAndVersion() (string, error) {
	return _LombardVerifier.Contract.TypeAndVersion(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) VersionTag() ([4]byte, error) {
	return _LombardVerifier.Contract.VersionTag(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) VersionTag() ([4]byte, error) {
	return _LombardVerifier.Contract.VersionTag(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_LombardVerifier *LombardVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardVerifier.Contract.AcceptOwnership(&_LombardVerifier.TransactOpts)
}

func (_LombardVerifier *LombardVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardVerifier.Contract.AcceptOwnership(&_LombardVerifier.TransactOpts)
}

func (_LombardVerifier *LombardVerifierTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_LombardVerifier *LombardVerifierSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyAllowlistUpdates(&_LombardVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_LombardVerifier *LombardVerifierTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyAllowlistUpdates(&_LombardVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_LombardVerifier *LombardVerifierTransactor) ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "applyRemoteChainConfigUpdates", remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyRemoteChainConfigUpdates(&_LombardVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierTransactorSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyRemoteChainConfigUpdates(&_LombardVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "forwardToVerifier", message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ForwardToVerifier(&_LombardVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ForwardToVerifier(&_LombardVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierTransactor) RemovePaths(opts *bind.TransactOpts, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "removePaths", remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierSession) RemovePaths(remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.Contract.RemovePaths(&_LombardVerifier.TransactOpts, remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierTransactorSession) RemovePaths(remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.Contract.RemovePaths(&_LombardVerifier.TransactOpts, remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_LombardVerifier *LombardVerifierSession) SetDynamicConfig(dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetDynamicConfig(&_LombardVerifier.TransactOpts, dynamicConfig)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetDynamicConfig(dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetDynamicConfig(&_LombardVerifier.TransactOpts, dynamicConfig)
}

func (_LombardVerifier *LombardVerifierTransactor) SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setPath", remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierTransactor) SetRemoteAdapters(opts *bind.TransactOpts, remoteAdapterArgs []LombardVerifierRemoteAdapterArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setRemoteAdapters", remoteAdapterArgs)
}

func (_LombardVerifier *LombardVerifierSession) SetRemoteAdapters(remoteAdapterArgs []LombardVerifierRemoteAdapterArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetRemoteAdapters(&_LombardVerifier.TransactOpts, remoteAdapterArgs)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetRemoteAdapters(remoteAdapterArgs []LombardVerifierRemoteAdapterArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetRemoteAdapters(&_LombardVerifier.TransactOpts, remoteAdapterArgs)
}

func (_LombardVerifier *LombardVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_LombardVerifier *LombardVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.TransferOwnership(&_LombardVerifier.TransactOpts, to)
}

func (_LombardVerifier *LombardVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.TransferOwnership(&_LombardVerifier.TransactOpts, to)
}

func (_LombardVerifier *LombardVerifierTransactor) UpdateSupportedTokens(opts *bind.TransactOpts, tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "updateSupportedTokens", tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierSession) UpdateSupportedTokens(tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.UpdateSupportedTokens(&_LombardVerifier.TransactOpts, tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierTransactorSession) UpdateSupportedTokens(tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.UpdateSupportedTokens(&_LombardVerifier.TransactOpts, tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierTransactor) VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "verifyMessage", message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.VerifyMessage(&_LombardVerifier.TransactOpts, message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierTransactorSession) VerifyMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.VerifyMessage(&_LombardVerifier.TransactOpts, message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_LombardVerifier *LombardVerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.WithdrawFeeTokens(&_LombardVerifier.TransactOpts, feeTokens)
}

func (_LombardVerifier *LombardVerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.WithdrawFeeTokens(&_LombardVerifier.TransactOpts, feeTokens)
}

type LombardVerifierAllowListSendersAddedIterator struct {
	Event *LombardVerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierAllowListSendersAdded)
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
		it.Event = new(LombardVerifierAllowListSendersAdded)
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

func (it *LombardVerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierAllowListSendersAddedIterator{contract: _LombardVerifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierAllowListSendersAdded)
				if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*LombardVerifierAllowListSendersAdded, error) {
	event := new(LombardVerifierAllowListSendersAdded)
	if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierAllowListSendersRemovedIterator struct {
	Event *LombardVerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierAllowListSendersRemoved)
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
		it.Event = new(LombardVerifierAllowListSendersRemoved)
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

func (it *LombardVerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierAllowListSendersRemovedIterator{contract: _LombardVerifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierAllowListSendersRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*LombardVerifierAllowListSendersRemoved, error) {
	event := new(LombardVerifierAllowListSendersRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierAllowListStateChangedIterator struct {
	Event *LombardVerifierAllowListStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierAllowListStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierAllowListStateChanged)
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
		it.Event = new(LombardVerifierAllowListStateChanged)
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

func (it *LombardVerifierAllowListStateChangedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierAllowListStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierAllowListStateChanged struct {
	DestChainSelector uint64
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListStateChangedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierAllowListStateChangedIterator{contract: _LombardVerifier.contract, event: "AllowListStateChanged", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierAllowListStateChanged)
				if err := _LombardVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseAllowListStateChanged(log types.Log) (*LombardVerifierAllowListStateChanged, error) {
	event := new(LombardVerifierAllowListStateChanged)
	if err := _LombardVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierDynamicConfigSetIterator struct {
	Event *LombardVerifierDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierDynamicConfigSet)
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
		it.Event = new(LombardVerifierDynamicConfigSet)
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

func (it *LombardVerifierDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierDynamicConfigSet struct {
	DynamicConfig LombardVerifierDynamicConfig
	Raw           types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardVerifierDynamicConfigSetIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierDynamicConfigSetIterator{contract: _LombardVerifier.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierDynamicConfigSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseDynamicConfigSet(log types.Log) (*LombardVerifierDynamicConfigSet, error) {
	event := new(LombardVerifierDynamicConfigSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierFeeTokenWithdrawnIterator struct {
	Event *LombardVerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierFeeTokenWithdrawn)
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
		it.Event = new(LombardVerifierFeeTokenWithdrawn)
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

func (it *LombardVerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardVerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierFeeTokenWithdrawnIterator{contract: _LombardVerifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierFeeTokenWithdrawn)
				if err := _LombardVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*LombardVerifierFeeTokenWithdrawn, error) {
	event := new(LombardVerifierFeeTokenWithdrawn)
	if err := _LombardVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierOwnershipTransferRequestedIterator struct {
	Event *LombardVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierOwnershipTransferRequested)
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
		it.Event = new(LombardVerifierOwnershipTransferRequested)
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

func (it *LombardVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierOwnershipTransferRequestedIterator{contract: _LombardVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierOwnershipTransferRequested)
				if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*LombardVerifierOwnershipTransferRequested, error) {
	event := new(LombardVerifierOwnershipTransferRequested)
	if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierOwnershipTransferredIterator struct {
	Event *LombardVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierOwnershipTransferred)
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
		it.Event = new(LombardVerifierOwnershipTransferred)
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

func (it *LombardVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierOwnershipTransferredIterator{contract: _LombardVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierOwnershipTransferred)
				if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*LombardVerifierOwnershipTransferred, error) {
	event := new(LombardVerifierOwnershipTransferred)
	if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierPathRemovedIterator struct {
	Event *LombardVerifierPathRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierPathRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierPathRemoved)
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
		it.Event = new(LombardVerifierPathRemoved)
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

func (it *LombardVerifierPathRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierPathRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierPathRemoved struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierPathRemovedIterator{contract: _LombardVerifier.contract, event: "PathRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierPathRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "PathRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParsePathRemoved(log types.Log) (*LombardVerifierPathRemoved, error) {
	event := new(LombardVerifierPathRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "PathRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierPathSetIterator struct {
	Event *LombardVerifierPathSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierPathSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierPathSet)
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
		it.Event = new(LombardVerifierPathSet)
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

func (it *LombardVerifierPathSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierPathSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierPathSet struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierPathSetIterator{contract: _LombardVerifier.contract, event: "PathSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierPathSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "PathSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParsePathSet(log types.Log) (*LombardVerifierPathSet, error) {
	event := new(LombardVerifierPathSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "PathSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierRemoteAdapterSetIterator struct {
	Event *LombardVerifierRemoteAdapterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierRemoteAdapterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierRemoteAdapterSet)
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
		it.Event = new(LombardVerifierRemoteAdapterSet)
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

func (it *LombardVerifierRemoteAdapterSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierRemoteAdapterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierRemoteAdapterSet struct {
	RemoteChainSelector uint64
	Token               common.Address
	RemoteAdapter       [32]byte
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterRemoteAdapterSet(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*LombardVerifierRemoteAdapterSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "RemoteAdapterSet", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierRemoteAdapterSetIterator{contract: _LombardVerifier.contract, event: "RemoteAdapterSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchRemoteAdapterSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteAdapterSet, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "RemoteAdapterSet", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierRemoteAdapterSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "RemoteAdapterSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseRemoteAdapterSet(log types.Log) (*LombardVerifierRemoteAdapterSet, error) {
	event := new(LombardVerifierRemoteAdapterSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "RemoteAdapterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierRemoteChainConfigSetIterator struct {
	Event *LombardVerifierRemoteChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierRemoteChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierRemoteChainConfigSet)
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
		it.Event = new(LombardVerifierRemoteChainConfigSet)
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

func (it *LombardVerifierRemoteChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierRemoteChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierRemoteChainConfigSet struct {
	RemoteChainSelector uint64
	Router              common.Address
	AllowlistEnabled    bool
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardVerifierRemoteChainConfigSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierRemoteChainConfigSetIterator{contract: _LombardVerifier.contract, event: "RemoteChainConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierRemoteChainConfigSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseRemoteChainConfigSet(log types.Log) (*LombardVerifierRemoteChainConfigSet, error) {
	event := new(LombardVerifierRemoteChainConfigSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierStorageLocationsUpdatedIterator struct {
	Event *LombardVerifierStorageLocationsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierStorageLocationsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierStorageLocationsUpdated)
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
		it.Event = new(LombardVerifierStorageLocationsUpdated)
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

func (it *LombardVerifierStorageLocationsUpdatedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierStorageLocationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierStorageLocationsUpdated struct {
	OldLocations []string
	NewLocations []string
	Raw          types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*LombardVerifierStorageLocationsUpdatedIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierStorageLocationsUpdatedIterator{contract: _LombardVerifier.contract, event: "StorageLocationsUpdated", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *LombardVerifierStorageLocationsUpdated) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierStorageLocationsUpdated)
				if err := _LombardVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseStorageLocationsUpdated(log types.Log) (*LombardVerifierStorageLocationsUpdated, error) {
	event := new(LombardVerifierStorageLocationsUpdated)
	if err := _LombardVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierSupportedTokenRemovedIterator struct {
	Event *LombardVerifierSupportedTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierSupportedTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierSupportedTokenRemoved)
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
		it.Event = new(LombardVerifierSupportedTokenRemoved)
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

func (it *LombardVerifierSupportedTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierSupportedTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierSupportedTokenRemoved struct {
	Token common.Address
	Raw   types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterSupportedTokenRemoved(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenRemovedIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "SupportedTokenRemoved")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierSupportedTokenRemovedIterator{contract: _LombardVerifier.contract, event: "SupportedTokenRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchSupportedTokenRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "SupportedTokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierSupportedTokenRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseSupportedTokenRemoved(log types.Log) (*LombardVerifierSupportedTokenRemoved, error) {
	event := new(LombardVerifierSupportedTokenRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierSupportedTokenSetIterator struct {
	Event *LombardVerifierSupportedTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierSupportedTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierSupportedTokenSet)
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
		it.Event = new(LombardVerifierSupportedTokenSet)
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

func (it *LombardVerifierSupportedTokenSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierSupportedTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierSupportedTokenSet struct {
	LocalToken   common.Address
	LocalAdapter common.Address
	Raw          types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterSupportedTokenSet(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenSetIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "SupportedTokenSet")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierSupportedTokenSetIterator{contract: _LombardVerifier.contract, event: "SupportedTokenSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchSupportedTokenSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenSet) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "SupportedTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierSupportedTokenSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseSupportedTokenSet(log types.Log) (*LombardVerifierSupportedTokenSet, error) {
	event := new(LombardVerifierSupportedTokenSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenSet", log); err != nil {
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

func (LombardVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80")
}

func (LombardVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0x9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82")
}

func (LombardVerifierAllowListStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492")
}

func (LombardVerifierDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f60")
}

func (LombardVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (LombardVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LombardVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LombardVerifierPathRemoved) Topic() common.Hash {
	return common.HexToHash("0x8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b")
}

func (LombardVerifierPathSet) Topic() common.Hash {
	return common.HexToHash("0x83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9")
}

func (LombardVerifierRemoteAdapterSet) Topic() common.Hash {
	return common.HexToHash("0xf2bb53a7e6aae800a85fba961b2bc3124a23dd44d95fefe0b7b29bd90975a976")
}

func (LombardVerifierRemoteChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9")
}

func (LombardVerifierStorageLocationsUpdated) Topic() common.Hash {
	return common.HexToHash("0xec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586")
}

func (LombardVerifierSupportedTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0xbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a306")
}

func (LombardVerifierSupportedTokenSet) Topic() common.Hash {
	return common.HexToHash("0x086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e3")
}

func (_LombardVerifier *LombardVerifier) Address() common.Address {
	return _LombardVerifier.address
}

type LombardVerifierInterface interface {
	GetDynamicConfig(opts *bind.CallOpts) (LombardVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardVerifierPath, error)

	GetRemoteAdapter(opts *bind.CallOpts, remoteChainSelector uint64, token common.Address) ([32]byte, error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetStorageLocations(opts *bind.CallOpts) ([]string, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	IBridge(opts *bind.CallOpts) (common.Address, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error)

	RemovePaths(opts *bind.TransactOpts, remoteChainSelectors []uint64) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error)

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error)

	SetRemoteAdapters(opts *bind.TransactOpts, remoteAdapterArgs []LombardVerifierRemoteAdapterArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSupportedTokens(opts *bind.TransactOpts, tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*LombardVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*LombardVerifierAllowListSendersRemoved, error)

	FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListStateChangedIterator, error)

	WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListStateChanged(log types.Log) (*LombardVerifierAllowListStateChanged, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardVerifierDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LombardVerifierDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LombardVerifierFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LombardVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LombardVerifierOwnershipTransferred, error)

	FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathRemovedIterator, error)

	WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathRemoved(log types.Log) (*LombardVerifierPathRemoved, error)

	FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathSetIterator, error)

	WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathSet(log types.Log) (*LombardVerifierPathSet, error)

	FilterRemoteAdapterSet(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*LombardVerifierRemoteAdapterSetIterator, error)

	WatchRemoteAdapterSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteAdapterSet, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseRemoteAdapterSet(log types.Log) (*LombardVerifierRemoteAdapterSet, error)

	FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardVerifierRemoteChainConfigSetIterator, error)

	WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigSet(log types.Log) (*LombardVerifierRemoteChainConfigSet, error)

	FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*LombardVerifierStorageLocationsUpdatedIterator, error)

	WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *LombardVerifierStorageLocationsUpdated) (event.Subscription, error)

	ParseStorageLocationsUpdated(log types.Log) (*LombardVerifierStorageLocationsUpdated, error)

	FilterSupportedTokenRemoved(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenRemovedIterator, error)

	WatchSupportedTokenRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenRemoved) (event.Subscription, error)

	ParseSupportedTokenRemoved(log types.Log) (*LombardVerifierSupportedTokenRemoved, error)

	FilterSupportedTokenSet(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenSetIterator, error)

	WatchSupportedTokenSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenSet) (event.Subscription, error)

	ParseSupportedTokenSet(log types.Log) (*LombardVerifierSupportedTokenSet, error)

	Address() common.Address
}
