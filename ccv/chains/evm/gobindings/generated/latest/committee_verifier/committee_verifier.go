// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package committee_verifier

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

type CommitteeVerifierDynamicConfig struct {
	FeeAggregator  common.Address
	AllowlistAdmin common.Address
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

type SignatureQuorumValidatorSignatureConfig struct {
	SourceChainSelector uint64
	Threshold           uint8
	Signers             []common.Address
}

var CommitteeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySignatureConfigs\",\"inputs\":[{\"name\":\"sourceChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"signatureConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SignatureQuorumValidator.SignatureConfig[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSignatureConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"struct SignatureQuorumValidator.SignatureConfig[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"pendingStorageLocationsAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"storageLocationsAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllowedFinalityConfig\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferStorageLocationsAdmin\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocations\",\"inputs\":[{\"name\":\"newLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"tag\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsAdminTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsAdminTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"verifierVersion\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedStorageLocationsAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonOrderedOrNonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByStorageLocationsAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignerCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceNotConfigured\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c080604052346105ce576140cb803803809161001c82856105d3565b8339810190808203608081126105ce576040136105ce5760408051929083016001600160401b038111848210176103c957604052610059826105f6565b8352610067602083016105f6565b6020840190815260408301519092906001600160401b0381116105ce57810182601f820112156105ce5780519061009d8261060a565b936100ab60405195866105d3565b82855260208086019360051b830101918183116105ce5760208101935b838510610556578888886100de60608a016105f6565b90331561054557600180546001600160a01b0319163317905546608052600654815190919061010c8361060a565b9261011a60405194856105d3565b808452600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f90602086015b8382106104a05750505060005b81811061040b57505060005b81811061027c5750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586916101b86101aa926040519384936040855260408501906106b0565b9083820360208501526106b0565b0390a16001600160a01b031690811561026b5760a0919091529051600880546001600160a01b03199081166001600160a01b0393841690811790925583516009805490921690841617905560408051918252925190911660208201527f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd9190a1600a80546001600160a01b031916331790556040516139a99081610722823960805181611011015260a0518161330e0152f35b6342bcdf7f60e11b60005260046000fd5b82518110156103f55760208160051b84010151600654680100000000000000008110156103c9578060016102b39201600655610644565b9190916103df578051906001600160401b0382116103c9576102d5835461065f565b601f811161038c575b50602090601f83116001146103215760019493929160009183610316575b5050600019600383901b1c191690841b1790555b01610164565b015190508b806102fc565b90601f1983169184600052816000209260005b81811061037457509160019695949291838895931061035b575b505050811b019055610310565b015160001960f88460031b161c191690558b808061034e565b92936020600181928786015181550195019301610334565b6103b990846000526020600020601f850160051c810191602086106103bf575b601f0160051c0190610699565b8a6102de565b90915081906103ac565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600654801561048a57600019019061042282610644565b9290926103df57826104366001945461065f565b9081610448575b505060065501610158565b81601f60009311861461045f5750555b8a8061043d565b8183526020832061047a91601f0160051c8101908701610699565b8082528160208120915555610458565b634e487b7160e01b600052603160045260246000fd5b604051600084546104b08161065f565b808452906001811690811561052257506001146104ea575b50600192826104dc859460209403826105d3565b81520193019101909161014b565b6000868152602081209092505b81831061050c575050810160200160016104c8565b60018160209254838688010152019201916104f7565b60ff191660208581019190915291151560051b84019091019150600190506104c8565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116105ce5782019083603f830112156105ce576020820151906001600160401b0382116103c95760405161059f601f8401601f1916602001826105d3565b82815260408484010186106105ce576105c360209493859460408685019101610621565b8152019401936100c8565b600080fd5b601f909101601f19168101906001600160401b038211908210176103c957604052565b51906001600160a01b03821682036105ce57565b6001600160401b0381116103c95760051b60200190565b60005b8381106106345750506000910152565b8181015183820152602001610624565b6006548110156103f557600660005260206000200190600090565b90600182811c9216801561068f575b602083101461067957565b634e487b7160e01b600052602260045260246000fd5b91607f169161066e565b8181106106a4575050565b60008155600101610699565b9080602083519182815201916020808360051b8301019401926000915b8383106106dc57505050505090565b909192939460208080600193601f19868203018752895161070881518092818552858086019101610621565b601f01601f1916010197019594919091019201906106cd56fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714612bdf57508063181f5a7714612b62578063296947061461285b5780632b7ae733146127945780633a3d72b5146126e3578063449e6a971461225b578063597b95c314611efc5780635cb80c5d14611cc75780635ef2c64b146118985780637437ff9f146117d757806379ba5097146116ee5780638282cbfe14611605578063869b7f62146114c157806387ae929214611473578063898068fc146112cb57806389e364c714610e535780638da5cb5b14610e01578063a4422ad814610daf578063aa8dac9814610b32578063b6cfa3b714610a78578063c9b146b314610619578063ec6ae7a7146105b8578063f2fde38b146104c8578063f4cdd89e146101ed578063f96c05ca1461019b5763fe163eed1461013d57600080fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760206040517fe9a05a20000000000000000000000000000000000000000000000000000000008152f35b600080fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b346101965760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657610224612e4f565b60243567ffffffffffffffff81116101965760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610196576040519060a0820182811067ffffffffffffffff82111761049957604052806004013567ffffffffffffffff8111610196576102a4906004369184010161300b565b8252602481013567ffffffffffffffff8111610196576102ca906004369184010161300b565b6020830152604481013567ffffffffffffffff8111610196578101366023820112156101965760048101356102fe81612f17565b9161030c6040519384612d1a565b818352602060048185019360061b830101019036821161019657602401915b81831061046157505050604083015261034660648201612e00565b6060830152608481013567ffffffffffffffff8111610196576080916004610371923692010161300b565b91015260443567ffffffffffffffff81116101965761039490369060040161300b565b50606435907fffffffff00000000000000000000000000000000000000000000000000000000821682036101965767ffffffffffffffff168060005260056020526040600020549073ffffffffffffffffffffffffffffffffffffffff821615610434575061040b60609260075460e01b9061348a565b61ffff60405191818160a01c16835263ffffffff8160b01c16602084015260d01c166040820152f35b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b604083360312610196576020604091825161047b81612cfe565b61048486612e00565b8152828601358382015281520192019161032b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965773ffffffffffffffffffffffffffffffffffffffff610514612ddd565b61051c6133ba565b1633811461058e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602060075460e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff811161019657610668903690600401612ee6565b9073ffffffffffffffffffffffffffffffffffffffff600154163303610a2e575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b80841015610a2c576000938060051b83013584811215610a28578301608081360312610a2857604051956080870187811067ffffffffffffffff8211176109fb5760405261070182612e66565b875261070f602083016130f6565b9360208801948552604083013567ffffffffffffffff81116109f757610738903690850161307d565b9260408901938452606081013567ffffffffffffffff81116109f3576107609136910161307d565b966060890197885267ffffffffffffffff89511683526005602052604083209160ff835460e01c1687511515809115150361096a575b50600184989301975b89518051821015610824579073ffffffffffffffffffffffffffffffffffffffff6107cc826001946130e2565b51168c6107d9828d6136d8565b6107e6575b50500161079f565b602067ffffffffffffffff7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d8292511692604051908152a28c8c6107de565b5050975094959095835151610848575b5050506001919091019450909290506106b4565b9690949195929651151560001461093357855b8751805182101561091b576108858273ffffffffffffffffffffffffffffffffffffffff926130e2565b511680156108e457908161089b6001938961386c565b6108a7575b500161085b565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80602067ffffffffffffffff8d511692604051908152a28a6108a0565b60248867ffffffffffffffff8c51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509650935093506001915084939291868080610834565b60248667ffffffffffffffff8a51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b83547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560e01b7cff00000000000000000000000000000000000000000000000000000000161784557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492602067ffffffffffffffff8d511692604051908152a28a610796565b8380fd5b8280fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60095416330315610689577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196577f307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a6020610ad2612c97565b610ada6133ba565b8060e01c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000060075416176007557fffffffff0000000000000000000000000000000000000000000000000000000060405191168152a1005b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760405180816020600354928381520160036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9260005b818110610d96575050610baf92500382612d1a565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610bf5610bdf84612f17565b93610bed6040519586612d1a565b808552612f17565b0160005b818110610d6a57505060005b8151811015610cbe5767ffffffffffffffff610c2182846130e2565b5116806000526002602052604060002060ff6002820154166040519182602082549182815201916000526020600020906000905b808210610ca65750505090610c71836001969594930383612d1a565b60405192610c7e84612cc6565b835260208301526040820152610c9482866130e2565b52610c9f81856130e2565b5001610c05565b90919260016020819286548152019401920190610c55565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610cf757505050500390f35b91936020610d5a827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652606060408a5167ffffffffffffffff815116845260ff8682015116868501520151918160408201520190612e7b565b9601920192018594939192610ce8565b602090604051610d7981612cc6565b600081526000838201526060604082015282828701015201610bf9565b8454835260019485019486945060209093019201610b9a565b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602073ffffffffffffffffffffffffffffffffffffffff600b5416604051908152f35b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101965760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff8111610196576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260040192360301126101965760443567ffffffffffffffff811161019657610ee5903690600401612e21565b91610ef7610ef282613029565b6132a8565b600683106112735782600411610196577fffffffff00000000000000000000000000000000000000000000000000000000823516917fe9a05a2000000000000000000000000000000000000000000000000000000000830361129d578360061161019657600481013560f01c91610f6d83613282565b851061127357610f7f610faa91613029565b9360405160208101918252602435602482015260248152610fa1604482612d1a565b51902092613282565b936000908560061161126f57851161126c57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa600667ffffffffffffffff9201940192169182600052600260205260406000209160ff60028401541693841561123f57507f000000000000000000000000000000000000000000000000000000000000000046810361120e5750838260061c106111e457600091825b85841061105057005b8360061b848104604014851517156111b55760208101908181116111b55761108361107d8383878d613290565b9061344f565b9060009260408201809211611188576020926110a861107d86946080948f8b90613290565b60405191898352601b868401526040830152606082015282805260015afa1561117c5773ffffffffffffffffffffffffffffffffffffffff815116916110fe838860019160005201602052604060002054151590565b156111545773ffffffffffffffffffffffffffffffffffffffff1682111561112c5750600190930192611047565b807fb70ad94b0000000000000000000000000000000000000000000000000000000060049252fd5b6004827fca31867a000000000000000000000000000000000000000000000000000000008152fd5b604051903d90823e3d90fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f320951440000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80fd5b5080fd5b7f1ede477b0000000000000000000000000000000000000000000000000000000060005260046000fd5b827fef8a07ee0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965767ffffffffffffffff61130b612e4f565b600060a060405161131b81612ce2565b8281528260208201528260408201528260608201528260808201520152168060005260056020526040600020908154906040519161135883612ce2565b73ffffffffffffffffffffffffffffffffffffffff8116835260208301918252604083019360ff8260e01c16151585526060840161ffff8360a01c1681526001608086019263ffffffff8560b01c16845261ffff60a088019560d01c168552019560405194856020895491828152019860005260206000209060005b81811061145d575050509261ffff86959363ffffffff9367ffffffffffffffff839761141a6114599b73ffffffffffffffffffffffffffffffffffffffff9e038b612d1a565b6040519c8d9c51168c52511660208b015251151560408a01525116606088015251166080860152511660a084015260e060c084015260e0830190612e7b565b0390f35b82548b526020909a0199600192830192016113d4565b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576114596114ad613165565b604051918291602083526020830190612f94565b346101965760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760006040516114fe81612cfe565b611506612ddd565b815260243573ffffffffffffffffffffffffffffffffffffffff811681036109f7578173ffffffffffffffffffffffffffffffffffffffff6115ff9260207f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd95019081526115726133ba565b818351167fffffffffffffffffffffffff0000000000000000000000000000000000000000600854161760085551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600954161760095560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b0390a180f35b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657600b5473ffffffffffffffffffffffffffffffffffffffff811633036116c4577fffffffffffffffffffffffff0000000000000000000000000000000000000000600a54913382841617600a5516600b5573ffffffffffffffffffffffffffffffffffffffff3391167fa3ddd2c19634c07b63b5c8b2685e01ac8be465118ec23afa866803f1f0b9bc4a600080a3005b7f2798c9ea0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760005473ffffffffffffffffffffffffffffffffffffffff811633036117ad577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576000602060405161181681612cfe565b828152015261145960405161182a81612cfe565b73ffffffffffffffffffffffffffffffffffffffff60085416815273ffffffffffffffffffffffffffffffffffffffff60095416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff811161019657366023820112156101965780600401356118f281612f17565b916119006040519384612d1a565b8183526024602084019260051b820101903682116101965760248101925b828410611c86578473ffffffffffffffffffffffffffffffffffffffff600a54163303611c5c57600654815191611953613165565b9160005b818110611b885750506000925b8084106119bb577fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075866119a8846119b685604051938493604085526040850190612f94565b908382036020850152612f94565b0390a1005b6119c584836130e2565b519360065468010000000000000000811015610499578060016119eb9201600655613405565b611b5957855167ffffffffffffffff811161049957611a0a8254613112565b601f8111611b1c575b506020601f8211600114611a76578190600195969798600092611a6b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82861b9260031b1c19161790555b01929190611964565b015190508880611a31565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169783600052816000209860005b818110611b04575091600196979899918488959410611acd575b505050811b019055611a62565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080611ac0565b838301518b556001909a019960209384019301611aa6565b611b4990836000526020600020601f840160051c81019160208510611b4f575b601f0160051c0190613438565b87611a13565b9091508190611b3c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b600694929394548015611c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190611bc182613405565b929092611b595782611bd560019454613112565b9081611beb575b50506006550193929193611957565b81601f600093118614611c025750555b8780611bdc565b81835260208320611c1d91601f0160051c8101908701613438565b8082528160208120915555611bfb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f5c8f80f00000000000000000000000000000000000000000000000000000000060005260046000fd5b833567ffffffffffffffff81116101965782013660438201121561019657602091611cbc83923690604460248201359101612f2f565b81520193019261191e565b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff811161019657611d16903690600401612ee6565b9073ffffffffffffffffffffffffffffffffffffffff60085416918215611ed25760005b818110611d4357005b611d4e81838561303e565b3573ffffffffffffffffffffffffffffffffffffffff8116809103610196576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115611e9557600091611ea1575b5080611dc2575b5050600101611d3a565b60206000604051828101907fa9059cbb00000000000000000000000000000000000000000000000000000000825289602482015284604482015260448152611e0b606482612d1a565b519082865af115611e95576000513d611e8c5750813b155b611e5e5790857f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39085611db8565b507f5274afe70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60011415611e23565b6040513d6000823e3d90fd5b906020823d8211611eca575b81611eba60209383612d1a565b8101031261126c57505186611db1565b3d9150611ead565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff8111610196573660238201121561019657806004013567ffffffffffffffff811161019657602460c0820283010136811161019657611f746133ba565b611f7d82612f17565b91611f8b6040519384612d1a565b825260009260240190602083015b8183106121ae578480855b80518310156121aa57611fb783826130e2565b519267ffffffffffffffff60208501511693841561217e578484526005602052604080852082518154928401517fffffff00ffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560e01b7cff000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156121525773ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007bffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260e01c1615156020820152a2019190611fa4565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261225757604051906121c582612ce2565b833573ffffffffffffffffffffffffffffffffffffffff811681036122535782526121f260208501612e66565b6020830152612203604085016130f6565b604083015261221460608501613103565b606083015260808401359063ffffffff821682036122535782602092608060c095015261224360a08701613103565b60a0820152815201920191611f99565b8680fd5b8480fd5b346101965760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff8111610196576122aa903690600401612ee6565b60243567ffffffffffffffff8111610196576122ca903690600401612ee6565b9290916122d56133ba565b60005b81811061257e575050506000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1823603015b81841015610a2c576000928460051b8101358281121561225757810194606086360312612257576040519161233f83612cc6565b61234887612e66565b835260208701359660ff881688036122535760208401978852604081013567ffffffffffffffff811161257a576123819136910161307d565b93604084019480865260ff895116801591821561256f575b50506125475767ffffffffffffffff84511687526002602052604087209586548860018901905b82811061252157505050878755875b86518051821015612488576123f98273ffffffffffffffffffffffffffffffffffffffff926130e2565b5116156124605761242b73ffffffffffffffffffffffffffffffffffffffff612423838a516130e2565b51168961386c565b15612438576001016123cf565b6004897f12823a5e000000000000000000000000000000000000000000000000000000008152fd5b6004897fcfb6108a000000000000000000000000000000000000000000000000000000008152fd5b50509790957f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b8392975067ffffffffffffffff600196949560ff926002848651169101907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556124fe8282511661380c565b50511693519151169061251660405192839283612ec5565b0390a201929061230b565b8061252e6001928c613420565b90549060031b1c8c52826020528b6040812055016123c0565b6004877f12823a5e000000000000000000000000000000000000000000000000000000008152fd5b511090508980612399565b8780fd5b604067ffffffffffffffff61259c61259784868861303e565b613029565b166000908152600460205220546125b6575b6001016122d8565b67ffffffffffffffff6125cd61259783858761303e565b16600052600260205260406000208054600060018301905b8281106126bb575050600082555060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560019061263f67ffffffffffffffff61263961259784878961303e565b16613571565b5061264e61259782858761303e565b7f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b836126ab60209267ffffffffffffffff6040519161268c8684612d1a565b6000835260003681376000604051948594604086526040860190612e7b565b9684015216930390a290506125ae565b806126c860019286613420565b90549060031b1c6000528260205260006040812055016125e5565b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965767ffffffffffffffff612723612e4f565b166000526002602052604060002060405190818092602083549182815201908360005260206000209060005b81811061277b578460ff60028861276884890385612d1a565b0154169061145960405192839283612ec5565b825484528694506020909301926001928301920161274f565b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576127cb612ddd565b73ffffffffffffffffffffffffffffffffffffffff600a541690813303611c5c5773ffffffffffffffffffffffffffffffffffffffff169033821461058e57817fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b557fdfee8caf308a35b723489c72952cf11683462281c34aa62e8af474dcd012f41a600080a3005b346101965760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff8111610196578036036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610196576128d2612dba565b5060843567ffffffffffffffff8111610196576128f3903690600401612e21565b50506024820191612906610ef284613029565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd6101248201359201821215610196570160048101359067ffffffffffffffff82116101965760248101918036038313610196576020908201919091031261019657359073ffffffffffffffffffffffffffffffffffffffff82168092036101965761299a67ffffffffffffffff91613029565b1680600052600560205260406000209081549073ffffffffffffffffffffffffffffffffffffffff8216908115610434576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa908115611e9557600091612aff575b5073ffffffffffffffffffffffffffffffffffffffff163303612ad15760e01c60ff16612a87575b6114596040517fe9a05a2000000000000000000000000000000000000000000000000000000000602082015260048152612a73602482612d1a565b604051918291602083526020830190612d5b565b60008281526002909101602052604090205415612aa45780612a38565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011612b5a575b81612b1860209383612d1a565b8101031261126f57519073ffffffffffffffffffffffffffffffffffffffff8216820361126c575073ffffffffffffffffffffffffffffffffffffffff612a10565b3d9150612b0b565b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576114596040805190612ba38183612d1a565b601782527f436f6d6d6974746565566572696669657220322e302e30000000000000000000602083015251918291602083526020830190612d5b565b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576020907fffffffff00000000000000000000000000000000000000000000000000000000612c3a612c97565b167fd3e969cd000000000000000000000000000000000000000000000000000000008114908115612c6d575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612c66565b600435907fffffffff000000000000000000000000000000000000000000000000000000008216820361019657565b6060810190811067ffffffffffffffff82111761049957604052565b60c0810190811067ffffffffffffffff82111761049957604052565b6040810190811067ffffffffffffffff82111761049957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761049957604052565b919082519283825260005b848110612da55750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612d66565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361019657565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019657565b359073ffffffffffffffffffffffffffffffffffffffff8216820361019657565b9181601f840112156101965782359167ffffffffffffffff8311610196576020838186019501011161019657565b6004359067ffffffffffffffff8216820361019657565b359067ffffffffffffffff8216820361019657565b906020808351928381520192019060005b818110612e995750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612e8c565b9060ff612edf602092959495604085526040850190612e7b565b9416910152565b9181601f840112156101965782359167ffffffffffffffff8311610196576020808501948460051b01011161019657565b67ffffffffffffffff81116104995760051b60200190565b92919267ffffffffffffffff82116104995760405191612f77601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184612d1a565b829481845281830111610196578281602093846000960137010152565b9080602083519182815201916020808360051b8301019401926000915b838310612fc057505050505090565b9091929394602080612ffc837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951612d5b565b97019301930191939290612fb1565b9080601f830112156101965781602061302693359101612f2f565b90565b3567ffffffffffffffff811681036101965790565b919081101561304e5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561019657813561309481612f17565b926130a26040519485612d1a565b81845260208085019260051b82010192831161019657602001905b8282106130ca5750505090565b602080916130d784612e00565b8152019101906130bd565b805182101561304e5760209160051b010190565b3590811515820361019657565b359061ffff8216820361019657565b90600182811c9216801561315b575b602083101461312c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613121565b6006549061317282612f17565b916131806040519384612d1a565b808352600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9190602085015b8282106131bf5750505050565b604051600085546131cf81613112565b80845290600181169081156132415750600114613209575b50600192826131fb85946020940382612d1a565b8152019401910190926131b2565b6000878152602081209092505b81831061322b575050810160200160016131e7565b6001816020925483868801015201920191613216565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506131e7565b60060190816006116111b557565b90939293848311610196578411610196578101920390565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e955760009161337f575b506133475750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6020813d6020116133b2575b8161339860209383612d1a565b8101031261126f575190811515820361126c57503861333f565b3d915061338b565b73ffffffffffffffffffffffffffffffffffffffff6001541633036133db57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60065481101561304e57600660005260206000200190600090565b805482101561304e5760005260206000200190600090565b818110613443575050565b60008155600101613438565b35906020811061345d575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b7fffffffff00000000000000000000000000000000000000000000000000000000811690811561356c576134bd816138c1565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c161661356c5761ffff8360e01c16801591821561355b575b5050613507575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff1610905038806134fd565b505050565b60008181526004602052604090205480156136d1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116111b557600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111b557818103613662575b5050506003548015611c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161361f816003613420565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b6136b9613673613684936003613420565b90549060031b1c9283926003613420565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260046020526040600020553880806135e6565b5050600090565b9060018201918160005282602052604060002054801515600014613803577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116111b5578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111b5578181036137cc575b50505080548015611c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061378d8282613420565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6137ec6137dc6136849386613420565b90549060031b1c92839286613420565b905560005283602052604060002055388080613755565b50505050600090565b8060005260046020526040600020541560001461386657600354680100000000000000008110156104995761384d6136848260018594016003556003613420565b9055600354906000526004602052604060002055600190565b50600090565b60008281526001820160205260409020546136d1578054906801000000000000000082101561049957826138aa613684846001809601855584613420565b905580549260005201602052604060002055600190565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115613998577dffff0000000000000000000000000000000000000000000000000000000081161561398f5760ff60015b169060f01c80613959575b5060010361392c5750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b6010811061396a5750613921565b6001811b821661397d575b60010161395c565b91600181018091116111b55791613975565b60ff6000613916565b505056fea164736f6c634300081a000a",
}

var CommitteeVerifierABI = CommitteeVerifierMetaData.ABI

var CommitteeVerifierBin = CommitteeVerifierMetaData.Bin

func DeployCommitteeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig CommitteeVerifierDynamicConfig, storageLocations []string, rmn common.Address) (common.Address, *types.Transaction, *CommitteeVerifier, error) {
	parsed, err := CommitteeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeVerifierBin), backend, dynamicConfig, storageLocations, rmn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitteeVerifier{address: address, abi: *parsed, CommitteeVerifierCaller: CommitteeVerifierCaller{contract: contract}, CommitteeVerifierTransactor: CommitteeVerifierTransactor{contract: contract}, CommitteeVerifierFilterer: CommitteeVerifierFilterer{contract: contract}}, nil
}

type CommitteeVerifier struct {
	address common.Address
	abi     abi.ABI
	CommitteeVerifierCaller
	CommitteeVerifierTransactor
	CommitteeVerifierFilterer
}

type CommitteeVerifierCaller struct {
	contract *bind.BoundContract
}

type CommitteeVerifierTransactor struct {
	contract *bind.BoundContract
}

type CommitteeVerifierFilterer struct {
	contract *bind.BoundContract
}

type CommitteeVerifierSession struct {
	Contract     *CommitteeVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitteeVerifierCallerSession struct {
	Contract *CommitteeVerifierCaller
	CallOpts bind.CallOpts
}

type CommitteeVerifierTransactorSession struct {
	Contract     *CommitteeVerifierTransactor
	TransactOpts bind.TransactOpts
}

type CommitteeVerifierRaw struct {
	Contract *CommitteeVerifier
}

type CommitteeVerifierCallerRaw struct {
	Contract *CommitteeVerifierCaller
}

type CommitteeVerifierTransactorRaw struct {
	Contract *CommitteeVerifierTransactor
}

func NewCommitteeVerifier(address common.Address, backend bind.ContractBackend) (*CommitteeVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(CommitteeVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitteeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifier{address: address, abi: abi, CommitteeVerifierCaller: CommitteeVerifierCaller{contract: contract}, CommitteeVerifierTransactor: CommitteeVerifierTransactor{contract: contract}, CommitteeVerifierFilterer: CommitteeVerifierFilterer{contract: contract}}, nil
}

func NewCommitteeVerifierCaller(address common.Address, caller bind.ContractCaller) (*CommitteeVerifierCaller, error) {
	contract, err := bindCommitteeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierCaller{contract: contract}, nil
}

func NewCommitteeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeVerifierTransactor, error) {
	contract, err := bindCommitteeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierTransactor{contract: contract}, nil
}

func NewCommitteeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeVerifierFilterer, error) {
	contract, err := bindCommitteeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierFilterer{contract: contract}, nil
}

func bindCommitteeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitteeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeVerifier.Contract.CommitteeVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.CommitteeVerifierTransactor.contract.Transfer(opts)
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.CommitteeVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.contract.Transfer(opts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "forwardToVerifier", message, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _CommitteeVerifier.Contract.ForwardToVerifier(&_CommitteeVerifier.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _CommitteeVerifier.Contract.ForwardToVerifier(&_CommitteeVerifier.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetAllSignatureConfigs(opts *bind.CallOpts) ([]SignatureQuorumValidatorSignatureConfig, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getAllSignatureConfigs")

	if err != nil {
		return *new([]SignatureQuorumValidatorSignatureConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]SignatureQuorumValidatorSignatureConfig)).(*[]SignatureQuorumValidatorSignatureConfig)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetAllSignatureConfigs() ([]SignatureQuorumValidatorSignatureConfig, error) {
	return _CommitteeVerifier.Contract.GetAllSignatureConfigs(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetAllSignatureConfigs() ([]SignatureQuorumValidatorSignatureConfig, error) {
	return _CommitteeVerifier.Contract.GetAllSignatureConfigs(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CommitteeVerifier.Contract.GetAllowedFinalityConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CommitteeVerifier.Contract.GetAllowedFinalityConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitteeVerifierDynamicConfig, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitteeVerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitteeVerifierDynamicConfig)).(*CommitteeVerifierDynamicConfig)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetDynamicConfig() (CommitteeVerifierDynamicConfig, error) {
	return _CommitteeVerifier.Contract.GetDynamicConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetDynamicConfig() (CommitteeVerifierDynamicConfig, error) {
	return _CommitteeVerifier.Contract.GetDynamicConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, requestedFinality)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetPendingStorageLocationsAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getPendingStorageLocationsAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetPendingStorageLocationsAdmin() (common.Address, error) {
	return _CommitteeVerifier.Contract.GetPendingStorageLocationsAdmin(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetPendingStorageLocationsAdmin() (common.Address, error) {
	return _CommitteeVerifier.Contract.GetPendingStorageLocationsAdmin(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemoteChainConfig = *abi.ConvertType(out[0], new(BaseVerifierRemoteChainConfigArgs)).(*BaseVerifierRemoteChainConfigArgs)
	outstruct.AllowedSendersList = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _CommitteeVerifier.Contract.GetRemoteChainConfig(&_CommitteeVerifier.CallOpts, remoteChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _CommitteeVerifier.Contract.GetRemoteChainConfig(&_CommitteeVerifier.CallOpts, remoteChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) (GetSignatureConfig,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getSignatureConfig", sourceChainSelector)

	outstruct := new(GetSignatureConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetSignatureConfig(sourceChainSelector uint64) (GetSignatureConfig,

	error) {
	return _CommitteeVerifier.Contract.GetSignatureConfig(&_CommitteeVerifier.CallOpts, sourceChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetSignatureConfig(sourceChainSelector uint64) (GetSignatureConfig,

	error) {
	return _CommitteeVerifier.Contract.GetSignatureConfig(&_CommitteeVerifier.CallOpts, sourceChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetStorageLocations(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getStorageLocations")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetStorageLocations() ([]string, error) {
	return _CommitteeVerifier.Contract.GetStorageLocations(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetStorageLocations() ([]string, error) {
	return _CommitteeVerifier.Contract.GetStorageLocations(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetStorageLocationsAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getStorageLocationsAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetStorageLocationsAdmin() (common.Address, error) {
	return _CommitteeVerifier.Contract.GetStorageLocationsAdmin(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetStorageLocationsAdmin() (common.Address, error) {
	return _CommitteeVerifier.Contract.GetStorageLocationsAdmin(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) Owner() (common.Address, error) {
	return _CommitteeVerifier.Contract.Owner(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) Owner() (common.Address, error) {
	return _CommitteeVerifier.Contract.Owner(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeVerifier.Contract.SupportsInterface(&_CommitteeVerifier.CallOpts, interfaceId)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeVerifier.Contract.SupportsInterface(&_CommitteeVerifier.CallOpts, interfaceId)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) TypeAndVersion() (string, error) {
	return _CommitteeVerifier.Contract.TypeAndVersion(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) TypeAndVersion() (string, error) {
	return _CommitteeVerifier.Contract.TypeAndVersion(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) error {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "verifyMessage", message, messageHash, verifierResults)

	if err != nil {
		return err
	}

	return err

}

func (_CommitteeVerifier *CommitteeVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) error {
	return _CommitteeVerifier.Contract.VerifyMessage(&_CommitteeVerifier.CallOpts, message, messageHash, verifierResults)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) error {
	return _CommitteeVerifier.Contract.VerifyMessage(&_CommitteeVerifier.CallOpts, message, messageHash, verifierResults)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) VersionTag() ([4]byte, error) {
	return _CommitteeVerifier.Contract.VersionTag(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) VersionTag() ([4]byte, error) {
	return _CommitteeVerifier.Contract.VersionTag(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_CommitteeVerifier *CommitteeVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptOwnership(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptOwnership(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) AcceptStorageLocationsAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "acceptStorageLocationsAdmin")
}

func (_CommitteeVerifier *CommitteeVerifierSession) AcceptStorageLocationsAdmin() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptStorageLocationsAdmin(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) AcceptStorageLocationsAdmin() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptStorageLocationsAdmin(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyAllowlistUpdates(&_CommitteeVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyAllowlistUpdates(&_CommitteeVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applyRemoteChainConfigUpdates", remoteChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyRemoteChainConfigUpdates(&_CommitteeVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyRemoteChainConfigUpdates(&_CommitteeVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplySignatureConfigs(opts *bind.TransactOpts, sourceChainSelectorsToRemove []uint64, signatureConfigs []SignatureQuorumValidatorSignatureConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applySignatureConfigs", sourceChainSelectorsToRemove, signatureConfigs)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplySignatureConfigs(sourceChainSelectorsToRemove []uint64, signatureConfigs []SignatureQuorumValidatorSignatureConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplySignatureConfigs(&_CommitteeVerifier.TransactOpts, sourceChainSelectorsToRemove, signatureConfigs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplySignatureConfigs(sourceChainSelectorsToRemove []uint64, signatureConfigs []SignatureQuorumValidatorSignatureConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplySignatureConfigs(&_CommitteeVerifier.TransactOpts, sourceChainSelectorsToRemove, signatureConfigs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "setAllowedFinalityConfig", allowedFinality)
}

func (_CommitteeVerifier *CommitteeVerifierSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetAllowedFinalityConfig(&_CommitteeVerifier.TransactOpts, allowedFinality)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetAllowedFinalityConfig(&_CommitteeVerifier.TransactOpts, allowedFinality)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierSession) SetDynamicConfig(dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetDynamicConfig(&_CommitteeVerifier.TransactOpts, dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) SetDynamicConfig(dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetDynamicConfig(&_CommitteeVerifier.TransactOpts, dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitteeVerifier *CommitteeVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferOwnership(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferOwnership(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) TransferStorageLocationsAdmin(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "transferStorageLocationsAdmin", to)
}

func (_CommitteeVerifier *CommitteeVerifierSession) TransferStorageLocationsAdmin(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferStorageLocationsAdmin(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) TransferStorageLocationsAdmin(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferStorageLocationsAdmin(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) UpdateStorageLocations(opts *bind.TransactOpts, newLocations []string) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "updateStorageLocations", newLocations)
}

func (_CommitteeVerifier *CommitteeVerifierSession) UpdateStorageLocations(newLocations []string) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.UpdateStorageLocations(&_CommitteeVerifier.TransactOpts, newLocations)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) UpdateStorageLocations(newLocations []string) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.UpdateStorageLocations(&_CommitteeVerifier.TransactOpts, newLocations)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitteeVerifier *CommitteeVerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.WithdrawFeeTokens(&_CommitteeVerifier.TransactOpts, feeTokens)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.WithdrawFeeTokens(&_CommitteeVerifier.TransactOpts, feeTokens)
}

type CommitteeVerifierAllowListSendersAddedIterator struct {
	Event *CommitteeVerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierAllowListSendersAdded)
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
		it.Event = new(CommitteeVerifierAllowListSendersAdded)
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

func (it *CommitteeVerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierAllowListSendersAddedIterator{contract: _CommitteeVerifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierAllowListSendersAdded)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitteeVerifierAllowListSendersAdded, error) {
	event := new(CommitteeVerifierAllowListSendersAdded)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierAllowListSendersRemovedIterator struct {
	Event *CommitteeVerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierAllowListSendersRemoved)
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
		it.Event = new(CommitteeVerifierAllowListSendersRemoved)
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

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierAllowListSendersRemovedIterator{contract: _CommitteeVerifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierAllowListSendersRemoved)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitteeVerifierAllowListSendersRemoved, error) {
	event := new(CommitteeVerifierAllowListSendersRemoved)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierAllowListStateChangedIterator struct {
	Event *CommitteeVerifierAllowListStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierAllowListStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierAllowListStateChanged)
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
		it.Event = new(CommitteeVerifierAllowListStateChanged)
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

func (it *CommitteeVerifierAllowListStateChangedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierAllowListStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierAllowListStateChanged struct {
	DestChainSelector uint64
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListStateChangedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierAllowListStateChangedIterator{contract: _CommitteeVerifier.contract, event: "AllowListStateChanged", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierAllowListStateChanged)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseAllowListStateChanged(log types.Log) (*CommitteeVerifierAllowListStateChanged, error) {
	event := new(CommitteeVerifierAllowListStateChanged)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierConfigSetIterator struct {
	Event *CommitteeVerifierConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierConfigSet)
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
		it.Event = new(CommitteeVerifierConfigSet)
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

func (it *CommitteeVerifierConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierConfigSet struct {
	DynamicConfig CommitteeVerifierDynamicConfig
	Raw           types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierConfigSetIterator, error) {

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierConfigSetIterator{contract: _CommitteeVerifier.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseConfigSet(log types.Log) (*CommitteeVerifierConfigSet, error) {
	event := new(CommitteeVerifierConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierFeeTokenWithdrawnIterator struct {
	Event *CommitteeVerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierFeeTokenWithdrawn)
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
		it.Event = new(CommitteeVerifierFeeTokenWithdrawn)
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

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeVerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierFeeTokenWithdrawnIterator{contract: _CommitteeVerifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierFeeTokenWithdrawn)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitteeVerifierFeeTokenWithdrawn, error) {
	event := new(CommitteeVerifierFeeTokenWithdrawn)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierFinalityConfigSetIterator struct {
	Event *CommitteeVerifierFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierFinalityConfigSet)
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
		it.Event = new(CommitteeVerifierFinalityConfigSet)
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

func (it *CommitteeVerifierFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierFinalityConfigSet struct {
	AllowedFinality [4]byte
	Raw             types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierFinalityConfigSetIterator, error) {

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierFinalityConfigSetIterator{contract: _CommitteeVerifier.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierFinalityConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseFinalityConfigSet(log types.Log) (*CommitteeVerifierFinalityConfigSet, error) {
	event := new(CommitteeVerifierFinalityConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierOwnershipTransferRequestedIterator struct {
	Event *CommitteeVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierOwnershipTransferRequested)
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
		it.Event = new(CommitteeVerifierOwnershipTransferRequested)
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

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierOwnershipTransferRequestedIterator{contract: _CommitteeVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierOwnershipTransferRequested)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitteeVerifierOwnershipTransferRequested, error) {
	event := new(CommitteeVerifierOwnershipTransferRequested)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierOwnershipTransferredIterator struct {
	Event *CommitteeVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierOwnershipTransferred)
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
		it.Event = new(CommitteeVerifierOwnershipTransferred)
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

func (it *CommitteeVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierOwnershipTransferredIterator{contract: _CommitteeVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierOwnershipTransferred)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*CommitteeVerifierOwnershipTransferred, error) {
	event := new(CommitteeVerifierOwnershipTransferred)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierRemoteChainConfigSetIterator struct {
	Event *CommitteeVerifierRemoteChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierRemoteChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierRemoteChainConfigSet)
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
		it.Event = new(CommitteeVerifierRemoteChainConfigSet)
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

func (it *CommitteeVerifierRemoteChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierRemoteChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierRemoteChainConfigSet struct {
	RemoteChainSelector uint64
	Router              common.Address
	AllowlistEnabled    bool
	Raw                 types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CommitteeVerifierRemoteChainConfigSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierRemoteChainConfigSetIterator{contract: _CommitteeVerifier.contract, event: "RemoteChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierRemoteChainConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseRemoteChainConfigSet(log types.Log) (*CommitteeVerifierRemoteChainConfigSet, error) {
	event := new(CommitteeVerifierRemoteChainConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierSignatureConfigSetIterator struct {
	Event *CommitteeVerifierSignatureConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierSignatureConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierSignatureConfigSet)
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
		it.Event = new(CommitteeVerifierSignatureConfigSet)
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

func (it *CommitteeVerifierSignatureConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierSignatureConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierSignatureConfigSet struct {
	SourceChainSelector uint64
	Signers             []common.Address
	Threshold           uint8
	Raw                 types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterSignatureConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CommitteeVerifierSignatureConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "SignatureConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierSignatureConfigSetIterator{contract: _CommitteeVerifier.contract, event: "SignatureConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierSignatureConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "SignatureConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierSignatureConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseSignatureConfigSet(log types.Log) (*CommitteeVerifierSignatureConfigSet, error) {
	event := new(CommitteeVerifierSignatureConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierStorageLocationsAdminTransferRequestedIterator struct {
	Event *CommitteeVerifierStorageLocationsAdminTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierStorageLocationsAdminTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierStorageLocationsAdminTransferRequested)
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
		it.Event = new(CommitteeVerifierStorageLocationsAdminTransferRequested)
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

func (it *CommitteeVerifierStorageLocationsAdminTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierStorageLocationsAdminTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierStorageLocationsAdminTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterStorageLocationsAdminTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierStorageLocationsAdminTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "StorageLocationsAdminTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierStorageLocationsAdminTransferRequestedIterator{contract: _CommitteeVerifier.contract, event: "StorageLocationsAdminTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchStorageLocationsAdminTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsAdminTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "StorageLocationsAdminTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierStorageLocationsAdminTransferRequested)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsAdminTransferRequested", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseStorageLocationsAdminTransferRequested(log types.Log) (*CommitteeVerifierStorageLocationsAdminTransferRequested, error) {
	event := new(CommitteeVerifierStorageLocationsAdminTransferRequested)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsAdminTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierStorageLocationsAdminTransferredIterator struct {
	Event *CommitteeVerifierStorageLocationsAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierStorageLocationsAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierStorageLocationsAdminTransferred)
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
		it.Event = new(CommitteeVerifierStorageLocationsAdminTransferred)
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

func (it *CommitteeVerifierStorageLocationsAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierStorageLocationsAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierStorageLocationsAdminTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterStorageLocationsAdminTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierStorageLocationsAdminTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "StorageLocationsAdminTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierStorageLocationsAdminTransferredIterator{contract: _CommitteeVerifier.contract, event: "StorageLocationsAdminTransferred", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchStorageLocationsAdminTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsAdminTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "StorageLocationsAdminTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierStorageLocationsAdminTransferred)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsAdminTransferred", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseStorageLocationsAdminTransferred(log types.Log) (*CommitteeVerifierStorageLocationsAdminTransferred, error) {
	event := new(CommitteeVerifierStorageLocationsAdminTransferred)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierStorageLocationsUpdatedIterator struct {
	Event *CommitteeVerifierStorageLocationsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierStorageLocationsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierStorageLocationsUpdated)
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
		it.Event = new(CommitteeVerifierStorageLocationsUpdated)
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

func (it *CommitteeVerifierStorageLocationsUpdatedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierStorageLocationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierStorageLocationsUpdated struct {
	OldLocations []string
	NewLocations []string
	Raw          types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*CommitteeVerifierStorageLocationsUpdatedIterator, error) {

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierStorageLocationsUpdatedIterator{contract: _CommitteeVerifier.contract, event: "StorageLocationsUpdated", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsUpdated) (event.Subscription, error) {

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierStorageLocationsUpdated)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseStorageLocationsUpdated(log types.Log) (*CommitteeVerifierStorageLocationsUpdated, error) {
	event := new(CommitteeVerifierStorageLocationsUpdated)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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
type GetSignatureConfig struct {
	Signers   []common.Address
	Threshold uint8
}

func (CommitteeVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80")
}

func (CommitteeVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0x9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82")
}

func (CommitteeVerifierAllowListStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492")
}

func (CommitteeVerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0x781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd")
}

func (CommitteeVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitteeVerifierFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
}

func (CommitteeVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitteeVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CommitteeVerifierRemoteChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9")
}

func (CommitteeVerifierSignatureConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b83")
}

func (CommitteeVerifierStorageLocationsAdminTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xdfee8caf308a35b723489c72952cf11683462281c34aa62e8af474dcd012f41a")
}

func (CommitteeVerifierStorageLocationsAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0xa3ddd2c19634c07b63b5c8b2685e01ac8be465118ec23afa866803f1f0b9bc4a")
}

func (CommitteeVerifierStorageLocationsUpdated) Topic() common.Hash {
	return common.HexToHash("0xec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586")
}

func (_CommitteeVerifier *CommitteeVerifier) Address() common.Address {
	return _CommitteeVerifier.address
}

type CommitteeVerifierInterface interface {
	ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error)

	GetAllSignatureConfigs(opts *bind.CallOpts) ([]SignatureQuorumValidatorSignatureConfig, error)

	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitteeVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

		error)

	GetPendingStorageLocationsAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) (GetSignatureConfig,

		error)

	GetStorageLocations(opts *bind.CallOpts) ([]string, error)

	GetStorageLocationsAdmin(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, messageHash [32]byte, verifierResults []byte) error

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptStorageLocationsAdmin(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error)

	ApplySignatureConfigs(opts *bind.TransactOpts, sourceChainSelectorsToRemove []uint64, signatureConfigs []SignatureQuorumValidatorSignatureConfig) (*types.Transaction, error)

	SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferStorageLocationsAdmin(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocations(opts *bind.TransactOpts, newLocations []string) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitteeVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitteeVerifierAllowListSendersRemoved, error)

	FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListStateChangedIterator, error)

	WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListStateChanged(log types.Log) (*CommitteeVerifierAllowListStateChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitteeVerifierConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitteeVerifierFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CommitteeVerifierFinalityConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitteeVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitteeVerifierOwnershipTransferred, error)

	FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CommitteeVerifierRemoteChainConfigSetIterator, error)

	WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigSet(log types.Log) (*CommitteeVerifierRemoteChainConfigSet, error)

	FilterSignatureConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CommitteeVerifierSignatureConfigSetIterator, error)

	WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierSignatureConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSignatureConfigSet(log types.Log) (*CommitteeVerifierSignatureConfigSet, error)

	FilterStorageLocationsAdminTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierStorageLocationsAdminTransferRequestedIterator, error)

	WatchStorageLocationsAdminTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsAdminTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseStorageLocationsAdminTransferRequested(log types.Log) (*CommitteeVerifierStorageLocationsAdminTransferRequested, error)

	FilterStorageLocationsAdminTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierStorageLocationsAdminTransferredIterator, error)

	WatchStorageLocationsAdminTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsAdminTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseStorageLocationsAdminTransferred(log types.Log) (*CommitteeVerifierStorageLocationsAdminTransferred, error)

	FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*CommitteeVerifierStorageLocationsUpdatedIterator, error)

	WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationsUpdated) (event.Subscription, error)

	ParseStorageLocationsUpdated(log types.Log) (*CommitteeVerifierStorageLocationsUpdated, error)

	Address() common.Address
}
