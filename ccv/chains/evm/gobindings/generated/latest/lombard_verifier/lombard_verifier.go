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
	RemoteAdapter [32]byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"},{\"name\":\"storageLocation\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePaths\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSupportedTokens\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokensToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct LombardVerifier.SupportedTokenArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenRemoved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"messageMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustTransferTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteTokenOrAdapterMismatch\",\"inputs\":[{\"name\":\"bridgeToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAllowedCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]}]",
	Bin: "0x60c080604052346106705761497f803803809161001c8285610675565b8339810190808203608081126106705760201361067057604051602081016001600160401b0381118282101761047d5760405261005882610698565b81526020820151926001600160a01b038416928385036106705760408101516001600160401b03811161067057810182601f820112156106705780519061009e826106ac565b936100ac6040519586610675565b82855260208086019360051b830101918183116106705760208101935b8385106105f857505050505060606100e19101610698565b906001549080516100f1836106ac565b926100ff6040519485610675565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b8382106105535750505060005b8181106104bf57505060005b8181106103305750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161019d61018f92604051938493604085526040850190610752565b908382036020850152610752565b0390a16001600160a01b0316801561031f57608052331561030e5760038054336001600160a01b0319918216179091559051600a80546001600160a01b039092169190921681179091556040519081527ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f6090602090a180156102fd5760206004916040519283809263353c26b760e01b82525afa9081156102f1576000916102ad575b5060ff1660028103610294575060a0526040516141bb90816107c48239608051816130d3015260a051818181610d9e01528181611970015281816124e10152818161271001528181612782015261338a0152f35b63398bbe0560e11b600052600260045260245260446000fd5b6020813d6020116102e9575b816102c660209383610675565b810103126102e557519060ff821682036102e2575060ff610240565b80fd5b5080fd5b3d91506102b9565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b639b15e16f60e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b82518110156104a95760208160051b840101516001546801000000000000000081101561047d5780600161036792016001556106e6565b919091610493578051906001600160401b03821161047d576103898354610701565b601f8111610440575b50602090601f83116001146103d557600194939291600091836103ca575b5050600019600383901b1c191690841b1790555b01610149565b0151905038806103b0565b90601f1983169184600052816000209260005b81811061042857509160019695949291838895931061040f575b505050811b0190556103c4565b015160001960f88460031b161c19169055388080610402565b929360206001819287860151815501950193016103e8565b61046d90846000526020600020601f850160051c81019160208610610473575b601f0160051c019061073b565b38610392565b9091508190610460565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600154801561053d5760001901906104d6826106e6565b92909261049357826104ea60019454610701565b90816104fb575b505082550161013d565b81601f6000931186146105125750555b38806104f1565b8183526020832061052d91601f0160051c810190870161073b565b808252816020812091555561050b565b634e487b7160e01b600052603160045260246000fd5b6040516000845461056381610701565b80845290600181169081156105d5575060011461059d575b506001928261058f85946020940382610675565b815201930191019091610130565b6000868152602081209092505b8183106105bf5750508101602001600161057b565b60018160209254838688010152019201916105aa565b60ff191660208581019190915291151560051b840190910191506001905061057b565b84516001600160401b0381116106705782019083603f83011215610670576020820151906001600160401b03821161047d57604051610641601f8401601f191660200182610675565b828152604084840101861061067057610665602094938594604086850191016106c3565b8152019401936100c9565b600080fd5b601f909101601f19168101906001600160401b0382119082101761047d57604052565b51906001600160a01b038216820361067057565b6001600160401b03811161047d5760051b60200190565b60005b8381106106d65750506000910152565b81810151838201526020016106c6565b6001548110156104a957600160005260206000200190600090565b90600182811c92168015610731575b602083101461071b57565b634e487b7160e01b600052602260045260246000fd5b91607f1691610710565b818110610746575050565b6000815560010161073b565b9080602083519182815201916020808360051b8301019401926000915b83831061077e57505050505090565b909192939460208080600193601f1986820301875289516107aa815180928185528580860191016106c3565b601f01601f19160101970195949190910192019061076f56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610197578063082522cb14610192578063181f5a771461018d578063240028e8146101885780632e45ca681461018357806338ff8c381461017e5780633bbbed4b146101795780635cb80c5d14610174578063708e1f791461016f578063737037e81461016a5780637437ff9f1461016557806379ba50971461016057806380485e251461015b57806387ae929214610156578063898068fc146101515780638da5cb5b1461014c5780638f2aaea414610147578063bcb6d4f714610142578063bff0ec1d1461013d578063c4bffe2b14610138578063c9b146b314610133578063d3c7c2c71461012e578063f2fde38b146101295763fe163eed1461012457600080fd5b61220f565b61211b565b61208f565b611d10565b611c4d565b611736565b611686565b6115b7565b611565565b6114c5565b6112e7565b6110d6565b610eda565b610e60565b610dc2565b610d53565b610be8565b610802565b6106f7565b61067a565b610607565b61055d565b61026d565b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361025657807f83adcde1000000000000000000000000000000000000000000000000000000006020921490811561022c575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610221565b600080fd5b67ffffffffffffffff81160361025657565b346102565760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576004356102a88161025b565b602435906044356064356102ba612ce1565b8315610381578115610357577fbe237be2cca72f95760d5f21feb5f0cf6579119971f023d6ccc49c749ddc92639161033b67ffffffffffffffff60405195610301876103da565b8387528760208801528460408801521694856000526009602052604060002090604060029180518455602081015160018501550151910155565b61034484613c15565b50604080519182526020820192909252a3005b7f55622b8a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f5a39e3030000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176103f657604052565b6103ab565b6020810190811067ffffffffffffffff8211176103f657604052565b6040810190811067ffffffffffffffff8211176103f657604052565b60c0810190811067ffffffffffffffff8211176103f657604052565b6080810190811067ffffffffffffffff8211176103f657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103f657604052565b604051906104bb60a08361046b565b565b67ffffffffffffffff81116103f657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b83811061050a5750506000910152565b81810151838201526020016104fa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610556815180928187528780880191016104f7565b0116010190565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576105da604080519061059e818361046b565b601982527f4c6f6d62617264566572696669657220322e302e302d6465760000000000000060208301525191829160208352602083019061051a565b0390f35b73ffffffffffffffffffffffffffffffffffffffff81160361025657565b35906104bb826105de565b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025657602061067073ffffffffffffffffffffffffffffffffffffffff60043561065c816105de565b166000526005602052604060002054151590565b6040519015158152f35b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760043567ffffffffffffffff8111610256573660238201121561025657806004013567ffffffffffffffff81116102565736602460c08302840101116102565760246106f59201612268565b005b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565767ffffffffffffffff60043561073b8161025b565b60006040805161074a816103da565b82815282602082015201521660005260096020526105da6040600020600260405191610775836103da565b80548352600181015460208401520154604082015260405191829182919091604080606083019480518452602081015160208501520151910152565b90816101c09103126102565790565b9181601f840112156102565782359167ffffffffffffffff8311610256576020838186019501011161025657565b9060206107ff92818152019061051a565b90565b346102565760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760043567ffffffffffffffff8111610256576108519036906004016107b1565b602435906108606044356105de565b60843567ffffffffffffffff8111610256576108809036906004016107c0565b505060208101803560006108938261025b565b61089c82613081565b6101808401906108ac8286612389565b905015610b42576108bc8361025b565b61012085019273ffffffffffffffffffffffffffffffffffffffff6108ec6108e486896123dd565b81019061242e565b169061090c8167ffffffffffffffff166000526000602052604060002090565b8054909161094473ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811615610b0c576040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152602090829060249082905afa8015610b075773ffffffffffffffffffffffffffffffffffffffff918691610ad8575b50163303610aac5760f01c60ff16610a1d575b6105da610a11898989610a098a610a036109fd6109f78d87612389565b90612472565b9361237f565b936123dd565b929091613262565b604051918291826107ee565b610a5d610a61916001610a4661092b8673ffffffffffffffffffffffffffffffffffffffff1690565b910160019160005201602052604060002054151590565b1590565b610a6b57806109da565b7fd0d2597600000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff16600452602490fd5b7f728fe07b00000000000000000000000000000000000000000000000000000000845233600452602484fd5b610afa915060203d602011610b00575b610af2818361046b565b810190612be4565b386109c7565b503d610ae8565b612bf9565b7f4d1aff7e00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff8216600452602486fd5b807f4f73dc4d0000000000000000000000000000000000000000000000000000000060049252fd5b9181601f840112156102565782359167ffffffffffffffff8311610256576020808501948460051b01011161025657565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610256576004359067ffffffffffffffff821161025657610be491600401610b6a565b9091565b3461025657610bf636610b9b565b9073ffffffffffffffffffffffffffffffffffffffff600a5416918215610d295760005b818110610c2357005b610c3961092b610c348385876127c3565b6127d3565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa8015610b07576001948892600092610cf9575b5081610cad575b5050505001610c1a565b81610cdd7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610ced94613da7565b6040519081529081906020820190565b0390a338858180610ca3565b610d1b91925060203d8111610d22575b610d13818361046b565b810190613198565b9038610c9c565b503d610d09565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102565760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760043567ffffffffffffffff811161025657610e11903690600401610b6a565b6024359167ffffffffffffffff831161025657366023840112156102565782600401359167ffffffffffffffff8311610256573660248460061b860101116102565760246106f59401916124b0565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576000604051610e9d816103fb565b526105da604051610ead816103fb565b600a5473ffffffffffffffffffffffffffffffffffffffff16908190526040519081529081906020820190565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760025473ffffffffffffffffffffffffffffffffffffffff81163303610f99577fffffffffffffffffffffffff00000000000000000000000000000000000000006003549133828416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610fcf826104bd565b91610fdd604051938461046b565b829481845281830111610256578281602093846000960137010152565b9080601f83011215610256578160206107ff93359101610fc3565b67ffffffffffffffff81116103f65760051b60200190565b81601f820112156102565780359061104482611015565b92611052604051948561046b565b82845260208085019360061b8301019181831161025657602001925b82841061107c575050505090565b604084830312610256576020604091825161109681610417565b86356110a1816105de565b8152828701358382015281520193019261106e565b6064359061ffff8216820361025657565b359061ffff8216820361025657565b346102565760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576004356111118161025b565b60243567ffffffffffffffff81116102565760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610256576111576104ac565b90806004013567ffffffffffffffff81116102565761117c9060043691840101610ffa565b8252602481013567ffffffffffffffff8111610256576111a29060043691840101610ffa565b6020830152604481013567ffffffffffffffff8111610256576111cb906004369184010161102d565b60408301526111dc606482016105fc565b6060830152608481013567ffffffffffffffff81116102565760809160046112079236920101610ffa565b91015260443567ffffffffffffffff8111610256576105da91611231611240923690600401610ffa565b5061123a6110b6565b50612826565b6040805161ffff909416845263ffffffff92831660208501529116908201529081906060820190565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061129c57505050505090565b90919293946020806112d8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161051a565b9701930193019193929061128d565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760015461132281611015565b61132f604051918261046b565b818152602081019160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66000935b82851061137657604051806105da8682611269565b604051600083548060011c9060018116908115611471575b60208310821461144457828552602085019190811561140f57506001146113d3575b5050600192826113c58594602094038261046b565b815201920194019390611361565b9091506113e585600052602060002090565b916000925b8184106113fb5750500182826113b0565b6001816020925486860152019301926113ea565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682525090151560051b01905082826113b0565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f169161138e565b906020808351928381520192019060005b8181106114995750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161148c565b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565767ffffffffffffffff6004356115098161025b565b16600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000206105da61153d600183549301613f12565b604051938360ff869560f01c161515855216602084015260606040840152606083019061147b565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025657602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256577ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f60611681604051611616816103fb565b600435611622816105de565b815261162c612ce1565b51600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691821790556040519081529081906020820190565b0390a1005b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760043567ffffffffffffffff81116102565736602382011215610256578060040135906116e182611015565b916116ef604051938461046b565b8083526024602084019160051b8301019136831161025657602401905b82821061171c576106f5846128f1565b60208091833561172b8161025b565b81520191019061170c565b346102565760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760043567ffffffffffffffff8111610256576117859036906004016107b1565b6024359060443567ffffffffffffffff8111610256576117a99036906004016107c0565b906117bb6117b68461237f565b613081565b6117cc6117c78461237f565b6137ed565b6117df6117d98383612a68565b90612ac9565b7ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603611aaf5750600690818310611bdf5761185a61185361184d611847858786612a76565b90612b7e565b60f01c90565b61ffff1690565b9261186d6118688585612b71565b612b5e565b8110611bdf5761188c6118836118e19585612b71565b80948385612ab1565b94909561018081016118ae6118a46109f78385612389565b60608101906123dd565b6118d86109f76118d06118c66109f787899799612389565b60808101906123dd565b959094612389565b3593898b6139be565b61190761190161185361184d6118476118f988612b5e565b888789612ab1565b93612b5e565b906119128483612b71565b8110611bdf5761192561192b9483612b71565b92612ab1565b92604051927fd5438eae00000000000000000000000000000000000000000000000000000000845260208460048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215610b07576000948593611ba6575b5073ffffffffffffffffffffffffffffffffffffffff85916119f4604051988997889687947fa620850600000000000000000000000000000000000000000000000000000000865260048601612cba565b0393165af1908115610b0757600090600092611b7f575b5015611b55576024815103611b225760246020820151910151907ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603611aaf5750818103611a7857005b7f6c86fa3a0000000000000000000000000000000000000000000000000000000060005260049190915260245260446000fd5b6000fd5b7fadaf7739000000000000000000000000000000000000000000000000000000006000527ff0f3a135000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b517fc2fdac9800000000000000000000000000000000000000000000000000000000600052602460048190525260446000fd5b7f2532cf450000000000000000000000000000000000000000000000000000000060005260046000fd5b9050611b9e91503d806000833e611b96818361046b565b810190612c05565b915038611a0b565b85919350611bd773ffffffffffffffffffffffffffffffffffffffff9160203d602011610b0057610af2818361046b565b9391506119a3565b7f1ede477b0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b818110611c2d5750505090565b825167ffffffffffffffff16845260209384019390920191600101611c20565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025657600754611c8881611015565b90611c96604051928361046b565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611cc382611015565b0136602084013760005b818110611ce257604051806105da8582611c09565b8067ffffffffffffffff611cf7600193613be2565b90549060031b1c16611d098286612a18565b5201611ccd565b3461025657611d1e36610b9b565b611d26612ce1565b6000905b808210611d3357005b611d46611d41838386613ab5565b613b5c565b92611d76611d5c855167ffffffffffffffff1690565b67ffffffffffffffff166000526000602052604060002090565b92611d86845460ff9060f01c1690565b916020860192611d968451151590565b9081151590151503611fe5575b506060860194600101939060005b86518051821015611e8b5790611de6611dcc82600194612a18565b5173ffffffffffffffffffffffffffffffffffffffff1690565b611e0e611e0873ffffffffffffffffffffffffffffffffffffffff831661092b565b896140cd565b611e1a575b5001611db1565b7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82611e8267ffffffffffffffff611e598d5167ffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff909516855216929081906020820190565b0390a238611e13565b50509450949190926040830191825151611eae575b505050506001019091611d2a565b519295919490939215611fd05760005b85518051821015611fbd57611dcc82611ed692612a18565b73ffffffffffffffffffffffffffffffffffffffff811615611f725760019190611f1e611f1873ffffffffffffffffffffffffffffffffffffffff831661092b565b88613ca6565b611f2a575b5001611ebe565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80611f6967ffffffffffffffff611e598c5167ffffffffffffffff1690565b0390a238611f23565b611aab611f87895167ffffffffffffffff1690565b7f463258ff0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b5050935093506001915090388080611ea0565b611aab611f87875167ffffffffffffffff1690565b85547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560f01b7eff000000000000000000000000000000000000000000000000000000000000161786557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f852349261208667ffffffffffffffff6120728a5167ffffffffffffffff1690565b604051941515855216929081906020820190565b0390a238611da3565b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610256576040516004548082526020820190600460005260206000209060005b818110612105576105da856120f18187038261046b565b60405191829160208352602083019061147b565b82548452602090930192600192830192016120da565b346102565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565773ffffffffffffffffffffffffffffffffffffffff60043561216b816105de565b612173612ce1565b163381146121e557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102565760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102565760206040517ff0f3a135000000000000000000000000000000000000000000000000000000008152f35b90612271612ce1565b61227a81611015565b91612288604051938461046b565b81835260c060208401920281019036821161025657915b8183106122b2575050506104bb90612d2c565b60c08336031261025657602060c0916040516122cd81610433565b85356122d8816105de565b8152828601356122e78161025b565b8382015260408601356122f98161233c565b604082015261230a606087016110c7565b606082015261231b60808701612346565b608082015261232c60a08701612346565b60a082015281520192019161229f565b8015150361025657565b359063ffffffff8216820361025657565b90604051612364816103da565b60406002829480548452600181015460208501520154910152565b356107ff8161025b565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610256570180359067ffffffffffffffff821161025657602001918160051b3603831361025657565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610256570180359067ffffffffffffffff82116102565760200191813603831361025657565b9081602091031261025657356107ff816105de565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156124ab578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610256570190565b612443565b906124b9612ce1565b60005b81811061265e575050509060009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915b8181106125105750505050565b807f086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e361254761254260019486896127dd565b6127ed565b6125f36125d861256b835173ffffffffffffffffffffffffffffffffffffffff1690565b926125986020820194612592865173ffffffffffffffffffffffffffffffffffffffff1690565b906137b9565b50835173ffffffffffffffffffffffffffffffffffffffff168981156126215773ffffffffffffffffffffffffffffffffffffffff611dcc925b16613737565b915173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152a101612503565b611dcc915073ffffffffffffffffffffffffffffffffffffffff612659845173ffffffffffffffffffffffffffffffffffffffff1690565b6125d2565b8061266f610c3460019385876127c3565b61268e73ffffffffffffffffffffffffffffffffffffffff821661092b565b6126a66126a061092b61092b84613e09565b91613e6c565b6126b3575b5050016124bc565b7fbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a3069161275d9173ffffffffffffffffffffffffffffffffffffffff168015612767576127369073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906135dd565b60405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a138806126ab565b506127be73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83166135dd565b612736565b91908110156124ab5760051b0190565b356107ff816105de565b91908110156124ab5760061b0190565b6040813603126102565760206040519161280683610417565b8035612811816105de565b8352013561281e816105de565b602082015290565b9067ffffffffffffffff821680600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000205416156128c457600052600060205261ffff60406000205460a01c169063ffffffff6128b9816128998667ffffffffffffffff166000526000602052604060002090565b5460b01c169467ffffffffffffffff166000526000602052604060002090565b5460d01c1691929190565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b906128fa612ce1565b6000915b8051831015612a135767ffffffffffffffff61291a8483612a18565b51169261294361293e8567ffffffffffffffff166000526009602052604060002090565b612357565b67ffffffffffffffff851661295a610a5d82613ff0565b6129db57612995612983600195969767ffffffffffffffff166000526009602052604060002090565b60026000918281558260018201550155565b6020828101518351604094850151855191825292810192909252927f465d9b27e0af9978f975c48406a226aab254b237e8027798ee924ef96ee9bb0491a30191906128fe565b7fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b509050565b80518210156124ab5760209160051b010190565b91612a64918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b906004116102565790600490565b909291928360041161025657831161025657600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610256578411610256578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612afd575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9060028201809211612b6c57565b612b2f565b91908201809211612b6c57565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110612bb2575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b9081602091031261025657516107ff816105de565b6040513d6000823e3d90fd5b9091606082840312610256578151926020830151612c228161233c565b9260408101519067ffffffffffffffff8211610256570181601f82011215610256578051612c4f816104bd565b92612c5d604051948561046b565b81845260208284010111610256576107ff91602080850191016104f7565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9290612cd3906107ff9593604086526040860191612c7b565b926020818503910152612c7b565b73ffffffffffffffffffffffffffffffffffffffff600354163303612d0257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060005b8251811015612a1357612d438184612a18565b51612d636020612d538487612a18565b51015167ffffffffffffffff1690565b67ffffffffffffffff811691821561303457612d938267ffffffffffffffff166000526000602052604060002090565b91612df6612db5835173ffffffffffffffffffffffffffffffffffffffff1690565b849073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b612e56612e066040840151151590565b84547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f01b7eff00000000000000000000000000000000000000000000000000000000000016178455565b612eaf612e68606084015161ffff1690565b84547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff000000000000000000000000000000000000000016178455565b6080820190612ece612ec5835163ffffffff1690565b63ffffffff1690565b15612ffd575091612fce612fc361092b7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc994612f60612f1560019a99985163ffffffff1690565b86547fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1660b09190911b79ffffffff0000000000000000000000000000000000000000000016178655565b611dcc612f7460a083015163ffffffff1690565b86547fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff1660d09190911b7dffffffff000000000000000000000000000000000000000000000000000016178655565b915460f01c60ff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff939093168352901515602083015290a201612d30565b7f9e7205510000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f97ccaab70000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b9081602091031261025657516107ff8161233c565b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa908115610b0757600091613158575b506131215750565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b61317a915060203d602011613180575b613172818361046b565b81019061306c565b38613119565b503d613168565b9160206107ff938181520191612c7b565b90816020910312610256575190565b90604051917ff0f3a1350000000000000000000000000000000000000000000000000000000060208401526024830152602482526104bb60448361046b565b9190826040910312610256576020825192015190565b93906107ff97969373ffffffffffffffffffffffffffffffffffffffff60e09794819388521660208701521660408501526060840152608083015260a08201528160c0820152019061051a565b90604051916020830152602082526104bb60408361046b565b939192906080850192602061327785886123dd565b9050116135995761329161092b6108e460408901896123dd565b936132c2610a5d73ffffffffffffffffffffffffffffffffffffffff87166000526005602052604060002054151590565b613555576132e761293e8467ffffffffffffffff166000526009602052604060002090565b9283511561351e575061331961092b8061331473ffffffffffffffffffffffffffffffffffffffff891682565b613e09565b73ffffffffffffffffffffffffffffffffffffffff8116613516575b50602083810180516040517f6e48b60d000000000000000000000000000000000000000000000000000000008152600481019190915273ffffffffffffffffffffffffffffffffffffffff88811660248301527f0000000000000000000000000000000000000000000000000000000000000000169892816044818c5afa908115610b07576000916134f7575b506133e26133dd6133d660608d018d6123dd565b3691610fc3565b613d02565b80821415806134e9575b6134ae5750509360009361342c6134226133dd6133d660409c9a9761341c8f9d9b9861346199519a81019061242e565b9c6123dd565b9a359251916131a7565b9189519a8b998a9889977e8a1198000000000000000000000000000000000000000000000000000000008952600489016131fc565b03925af18015610b07576107ff9160009161347d575b50613249565b61349f915060403d6040116134a7575b613497818361046b565b8101906131e6565b905038613477565b503d61348d565b60408701517fbce7b6cd0000000000000000000000000000000000000000000000000000000060005260049290925260245260445260646000fd5b5060408701518214156133ec565b613510915060203d602011610d2257610d13818361046b565b386133c2565b945038613335565b7fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f06439c6b0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff851660045260246000fd5b6135a384876123dd565b906135d96040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613187565b0390fd5b60405190602060008184017f095ea7b300000000000000000000000000000000000000000000000000000000815261366c856136408489602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810187528661046b565b84519082855af1600051903d816136fe575b501590505b61368c57505050565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff939093166024840152600060448085019190915283526104bb926136f9906136f360648261046b565b82613e87565b613e87565b1515905061372b575061368373ffffffffffffffffffffffffffffffffffffffff82163b15155b3861367e565b60016136839114613725565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff851660248401527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60448401529192919060009061366c8560648101613640565b9073ffffffffffffffffffffffffffffffffffffffff806107ff931691826000526006602052166040600020556004613ca6565b61382b61092b6138118367ffffffffffffffff166000526000602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8116156138f5576040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082905afa908115610b07576000916138d6575b50156138a857565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6138ef915060203d60201161318057613172818361046b565b386138a0565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b9160c083830312610256578235926020810135926040820135926060830135613955816105de565b926080810135613964816105de565b9260a082013567ffffffffffffffff8111610256576107ff9201610ffa565b359060208110613991575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b959493929195816004116102565760046139db928201910161392d565b959450505050506139fb6021830151916081606185015194015197613983565b808203613a85575050613a0e8383613983565b03613a50575050808203613a20575050565b7f7c83fcf00000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b6135d96040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613187565b7fd27ededb0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b91908110156124ab5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301821215610256570190565b9080601f83011215610256578135613b0c81611015565b92613b1a604051948561046b565b81845260208085019260051b82010192831161025657602001905b828210613b425750505090565b602080918335613b51816105de565b815201910190613b35565b6080813603126102565760405190613b738261044f565b8035613b7e8161025b565b82526020810135613b8e8161233c565b6020830152604081013567ffffffffffffffff811161025657613bb49036908301613af5565b604083015260608101359067ffffffffffffffff821161025657613bda91369101613af5565b606082015290565b6007548110156124ab57600760005260206000200190600090565b80548210156124ab5760005260206000200190600090565b600081815260086020526040902054613ca057600754680100000000000000008110156103f657613c87613c528260018594016007556007613bfd565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600754906000526008602052604060002055600190565b50600090565b6000828152600182016020526040902054613cfb57805490680100000000000000008210156103f65782613ce4613c52846001809601855584613bfd565b905580549260005201602052604060002055600190565b5050600090565b6020815111613d7157602081519101519060208110613d40575b8060031b9080820460081490151715612b6c57610100036101008111612b6c571c90565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b1690613d1c565b6135d9906040519182917fe0d7fb02000000000000000000000000000000000000000000000000000000008352600483016107ee565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9290921660248301526044808301939093529181526104bb916136f960648361046b565b80600052600660205260406000205490811580613e56575b613e29575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5060008181526005602052604090205415613e21565b6107ff908060005260066020526000604081205560046140cd565b906000602091828151910182855af115612bf9576000513d613f09575073ffffffffffffffffffffffffffffffffffffffff81163b155b613ec55750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415613ebe565b906040519182815491828252602082019060005260206000209260005b818110613f445750506104bb9250038361046b565b8454835260019485019487945060209093019201613f2f565b80548015613fc1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190613f928282613bfd565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260086020526040902054908115613cfb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211612b6c57600754927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612b6c57838360009561408c9503614092575b50505061407b6007613f5d565b600890600052602052604060002090565b55600190565b61407b6140be916140b46140aa6140c4956007613bfd565b90549060031b1c90565b9283916007613bfd565b90612a2c565b5538808061406e565b60018101918060005282602052604060002054928315156000146141a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612b6c578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501948511612b6c57600095858361408c9761415d950361416c575b505050613f5d565b90600052602052604060002090565b61418c6140be916141836140aa61419c9588613bfd565b92839187613bfd565b8590600052602052604060002090565b55388080614155565b5050505060009056fea164736f6c634300081a000a",
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

func (_LombardVerifier *LombardVerifierTransactor) SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setPath", remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
}

func (_LombardVerifier *LombardVerifierSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
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
	RemoteAdapter       [32]byte
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
	RemoteAdapter       [32]byte
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
	return common.HexToHash("0x465d9b27e0af9978f975c48406a226aab254b237e8027798ee924ef96ee9bb04")
}

func (LombardVerifierPathSet) Topic() common.Hash {
	return common.HexToHash("0xbe237be2cca72f95760d5f21feb5f0cf6579119971f023d6ccc49c749ddc9263")
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

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte, remoteAdapter [32]byte) (*types.Transaction, error)

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
