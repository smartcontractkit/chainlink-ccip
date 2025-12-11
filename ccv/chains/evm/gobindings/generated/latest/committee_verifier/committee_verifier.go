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

type SignatureQuorumValidatorSignatureConfig struct {
	SourceChainSelector uint64
	Threshold           uint8
	Signers             []common.Address
}

var CommitteeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySignatureConfigs\",\"inputs\":[{\"name\":\"sourceChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"signatureConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SignatureQuorumValidator.SignatureConfig[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSignatureConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"struct SignatureQuorumValidator.SignatureConfig[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocationsAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferStorageLocationsAdmin\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocations\",\"inputs\":[{\"name\":\"newLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsAdminTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsAdminTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"verifierVersion\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierResults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedStorageLocationsAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonOrderedOrNonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByStorageLocationsAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignerCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceNotConfigured\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c080604052346105e657613e15803803809161001c82856105eb565b83398101818103608081126105e6576040136105e65760408051919082016001600160401b038111838210176103e1576040526100588361060e565b82526100666020840161060e565b6020830190815260408401519093906001600160401b0381116105e657810182601f820112156105e65780519061009c82610622565b936100aa60405195866105eb565b82855260208086019360051b830101918183116105e65760208101935b83851061056e578888886100dd60608a0161060e565b90331561055d57600180546001600160a01b0319163317905546608052600654815190919061010b83610622565b9261011960405194856105eb565b808452600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f90602086015b8382106104b85750505060005b81811061042357505060005b8181106102945750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586916101b76101a9926040519384936040855260408501906106c8565b9083820360208501526106c8565b0390a16001600160a01b031680156102835760a05280516001600160a01b0316156102725751600780546001600160a01b03199081166001600160a01b0393841690811790925583516008805490921690841617905560408051918252925190911660208201527f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd9190a1600980546001600160a01b031916331790556040516136db908161073a8239608051816108a1015260a051815050f35b6306b7c75960e31b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b825181101561040d5760208160051b84010151600654680100000000000000008110156103e1578060016102cb920160065561065c565b9190916103f7578051906001600160401b0382116103e1576102ed8354610677565b601f81116103a4575b50602090601f8311600114610339576001949392916000918361032e575b5050600019600383901b1c191690841b1790555b01610163565b015190508b80610314565b90601f1983169184600052816000209260005b81811061038c575091600196959492918388959310610373575b505050811b019055610328565b015160001960f88460031b161c191690558b8080610366565b9293602060018192878601518155019501930161034c565b6103d190846000526020600020601f850160051c810191602086106103d7575b601f0160051c01906106b1565b8a6102f6565b90915081906103c4565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60065480156104a257600019019061043a8261065c565b9290926103f7578261044e60019454610677565b9081610460575b505060065501610157565b81601f6000931186146104775750555b8a80610455565b8183526020832061049291601f0160051c81019087016106b1565b8082528160208120915555610470565b634e487b7160e01b600052603160045260246000fd5b604051600084546104c881610677565b808452906001811690811561053a5750600114610502575b50600192826104f4859460209403826105eb565b81520193019101909161014a565b6000868152602081209092505b818310610524575050810160200160016104e0565b600181602092548386880101520192019161050f565b60ff191660208581019190915291151560051b84019091019150600190506104e0565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116105e65782019083603f830112156105e6576020820151906001600160401b0382116103e1576040516105b7601f8401601f1916602001826105eb565b82815260408484010186106105e6576105db60209493859460408685019101610639565b8152019401936100c7565b600080fd5b601f909101601f19168101906001600160401b038211908210176103e157604052565b51906001600160a01b03821682036105e657565b6001600160401b0381116103e15760051b60200190565b60005b83811061064c5750506000910152565b818101518382015260200161063c565b60065481101561040d57600660005260206000200190600090565b90600182811c921680156106a7575b602083101461069157565b634e487b7160e01b600052602260045260246000fd5b91607f1691610686565b8181106106bc575050565b600081556001016106b1565b9080602083519182815201916020808360051b8301019401926000915b8383106106f457505050505090565b909192939460208080600193601f19868203018752895161072081518092818552858086019101610639565b601f01601f1916010197019594919091019201906106e556fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714612b3d57508063181f5a7714612ac05780632b7ae733146129f95780632e45ca68146126585780633a3d72b5146125a75780633bbbed4b14612272578063449e6a9714611de25780635cb80c5d14611af65780635ef2c64b146116c75780637437ff9f1461160657806379ba50971461151d57806380485e251461125b5780638282cbfe14611172578063869b7f6214610fe957806387ae929214610f9b578063898068fc14610ebc5780638da5cb5b14610e6a578063a4422ad814610e18578063aa8dac9814610b57578063bff0ec1d146106f5578063c9b146b3146102c7578063f2fde38b146101d7578063f96c05ca146101855763fe163eed1461012757600080fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760206040517f49ff34ed000000000000000000000000000000000000000000000000000000008152f35b600080fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805773ffffffffffffffffffffffffffffffffffffffff610223612d0b565b61022b6131de565b1633811461029d57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff811161018057610316903690600401612e37565b73ffffffffffffffffffffffffffffffffffffffff6001541633036106ab575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b818410156106a9576000938060051b840135828112156106a5578401916080833603126106a557604051946080860186811067ffffffffffffffff821117610678576040526103ae84612d89565b86526103bc60208501612f5d565b9660208701978852604085013567ffffffffffffffff8111610674576103e59036908701612fcf565b9460408801958652606081013567ffffffffffffffff81116106705761040d91369101612fcf565b946060880195865267ffffffffffffffff885116825260056020526040822098511515610487818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610544575b50959760010195505b845180518210156104d757906104d073ffffffffffffffffffffffffffffffffffffffff6104c883600195613034565b511688613415565b5001610498565b50509594909350600192519081516104f5575b505001929190610360565b61053a67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190612d9e565b0390a285806104ea565b989395929094979896919660001461063957600184019591875b865180518210156105db576105888273ffffffffffffffffffffffffffffffffffffffff92613034565b511680156105a4579061059d6001928a6135a9565b500161055e565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161062f67ffffffffffffffff8b51169251604051918291602083526020830190612d9e565b0390a2908961048f565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60085416330315610336577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101805760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff8111610180576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101805760443567ffffffffffffffff811161018057610783903690600401612e09565b9160068310610aff5782600411610180577fffffffff00000000000000000000000000000000000000000000000000000000823516917f49ff34ed000000000000000000000000000000000000000000000000000000008303610b29578360061161018057600481013560f01c916107fa836131b8565b8510610aff5761080f61083a91600401612f7b565b9360405160208101918252602435602482015260248152610831604482612c31565b519020926131b8565b9360009085600611610674578511610afc57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa600667ffffffffffffffff9201940192169182600052600260205260406000209160ff600284015416938415610acf57507f0000000000000000000000000000000000000000000000000000000000000000468103610a9e5750838260061c10610a7457600091825b8584106108e057005b8360061b84810460401485151715610a45576020810190818111610a455761091361090d8383878d6131c6565b90613273565b9060009260408201809211610a185760209261093861090d86946080948f8b906131c6565b60405191898352601b868401526040830152606082015282805260015afa15610a0c5773ffffffffffffffffffffffffffffffffffffffff8151169161098e838860019160005201602052604060002054151590565b156109e45773ffffffffffffffffffffffffffffffffffffffff168211156109bc57506001909301926108d7565b807fb70ad94b0000000000000000000000000000000000000000000000000000000060049252fd5b6004827fca31867a000000000000000000000000000000000000000000000000000000008152fd5b604051903d90823e3d90fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f320951440000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80fd5b7f1ede477b0000000000000000000000000000000000000000000000000000000060005260046000fd5b827fef8a07ee0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760405180816020600354928381520160036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9260005b818110610dff575050610bd492500382612c31565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610c1a610c0484612e68565b93610c126040519586612c31565b808552612e68565b0160005b818110610dd357505060005b8151811015610ce35767ffffffffffffffff610c468284613034565b5116806000526002602052604060002060ff6002820154166040519182602082549182815201916000526020600020906000905b808210610ccb5750505090610c96836001969594930383612c31565b60405192610ca384612bf9565b835260208301526040820152610cb98286613034565b52610cc48185613034565b5001610c2a565b90919260016020819286548152019401920190610c7a565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610d1c57505050500390f35b91939092947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc09082030182528451602060806040606085019367ffffffffffffffff815116865260ff848201511684870152015193606060408201528451809452019201906000905b808210610da5575050506020806001929601920192018594939192610d0d565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8751168152019401920190610d85565b602090604051610de281612bf9565b600081526000838201526060604082015282828701015201610c1e565b8454835260019485019486945060209093019201610bbf565b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805767ffffffffffffffff610efc612d72565b16600052600560205260406000206001815491019060405191826020825491828152019160005260206000209060005b818110610f855773ffffffffffffffffffffffffffffffffffffffff85610f8188610f5981890382612c31565b604051938360ff869560f01c1615158552166020840152606060408401526060830190612d9e565b0390f35b8254845260209093019260019283019201610f2c565b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057610f81610fd561309b565b604051918291602083526020830190612ee6565b346101805760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057600060405161102681612c15565b61102e612d0b565b815260243573ffffffffffffffffffffffffffffffffffffffff81168103610670576020820190815261105f6131de565b73ffffffffffffffffffffffffffffffffffffffff8251161561114a578173ffffffffffffffffffffffffffffffffffffffff61114492817f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600854161760085560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b0390a180f35b6004837f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057600a5473ffffffffffffffffffffffffffffffffffffffff81163303611231577fffffffffffffffffffffffff000000000000000000000000000000000000000060095491338284161760095516600a5573ffffffffffffffffffffffffffffffffffffffff3391167fa3ddd2c19634c07b63b5c8b2685e01ac8be465118ec23afa866803f1f0b9bc4a600080a3005b7f2798c9ea0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101805760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057611292612d72565b60243567ffffffffffffffff81116101805760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610180576040519060a0820182811067ffffffffffffffff8211176114ee57604052806004013567ffffffffffffffff8111610180576113129060043691840101612eb7565b8252602481013567ffffffffffffffff8111610180576113389060043691840101612eb7565b6020830152604481013567ffffffffffffffff81116101805781013660238201121561018057600481013561136c81612e68565b9161137a6040519384612c31565b818352602060048185019360061b830101019036821161018057602401915b8183106114b65750505060408301526113b460648201612d51565b6060830152608481013567ffffffffffffffff81116101805760809160046113df9236920101612eb7565b9101526044359067ffffffffffffffff82116101805761140c67ffffffffffffffff923690600401612eb7565b50611415612ed5565b501680600052600560205273ffffffffffffffffffffffffffffffffffffffff60406000205416156114895760009081526005602090815260409182902054825161ffff60a083901c16815263ffffffff60b083901c81169382019390935260d09190911c90911691810191909152606090f35b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60408336031261018057602060409182516114d081612c15565b6114d986612d51565b81528286013583820152815201920191611399565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760005473ffffffffffffffffffffffffffffffffffffffff811633036115dc577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610180576000602060405161164581612c15565b8281520152610f8160405161165981612c15565b73ffffffffffffffffffffffffffffffffffffffff60075416815273ffffffffffffffffffffffffffffffffffffffff60085416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff8111610180573660238201121561018057806004013561172181612e68565b9161172f6040519384612c31565b8183526024602084019260051b820101903682116101805760248101925b828410611ab5578473ffffffffffffffffffffffffffffffffffffffff600954163303611a8b5760065481519161178261309b565b9160005b8181106119b75750506000925b8084106117ea577fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075866117d7846117e585604051938493604085526040850190612ee6565b908382036020850152612ee6565b0390a1005b6117f48483613034565b5193600654680100000000000000008110156114ee5780600161181a9201600655613229565b61198857855167ffffffffffffffff81116114ee576118398254613048565b601f811161194b575b506020601f82116001146118a557819060019596979860009261189a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82861b9260031b1c19161790555b01929190611793565b015190508880611860565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169783600052816000209860005b8181106119335750916001969798999184889594106118fc575b505050811b019055611891565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558880806118ef565b838301518b556001909a0199602093840193016118d5565b61197890836000526020600020601f840160051c8101916020851061197e575b601f0160051c019061325c565b87611842565b909150819061196b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b600694929394548015611a5c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906119f082613229565b9290926119885782611a0460019454613048565b9081611a1a575b50506006550193929193611786565b81601f600093118614611a315750555b8780611a0b565b81835260208320611a4c91601f0160051c810190870161325c565b8082528160208120915555611a2a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f5c8f80f00000000000000000000000000000000000000000000000000000000060005260046000fd5b833567ffffffffffffffff81116101805782013660438201121561018057602091611aeb83923690604460248201359101612e80565b81520193019261174d565b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff811161018057611b45903690600401612e37565b9073ffffffffffffffffffffffffffffffffffffffff600754169160005b818110611b6c57005b611b77818385612f90565b359073ffffffffffffffffffffffffffffffffffffffff821680920361018057604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315611dd657600093611da3575b5082611bee575b506001915001611b63565b8560405193611cab60208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611c39606482612c31565b600080604098895193611c4c8b86612c31565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d15611d9b573d90611c8d82612c72565b91611c9a8a519384612c31565b82523d6000602084013e5b866135fe565b805180611ce7575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385611be3565b819294959693509060209181010312610180576020015180159081150361018057611d185792919086908880611cb3565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090611ca5565b90926020823d8211611dce575b81611dbd60209383612c31565b81010312610afc5750519186611bdc565b3d9150611db0565b6040513d6000823e3d90fd5b346101805760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff811161018057611e31903690600401612e37565b60243567ffffffffffffffff811161018057611e51903690600401612e37565b929091611e5c6131de565b60005b81811061210d575050506000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1823603015b818410156106a9576000928460051b81013582811215612109578101946060863603126121095760405191611ec683612bf9565b611ecf87612d89565b835260208701359660ff881688036121055760208401978852604081013567ffffffffffffffff811161210157611f0891369101612fcf565b93604084019480865260ff89511680159182156120f6575b50506120ce5767ffffffffffffffff84511687526002602052604087209586548860018901905b8281106120a857505050878755875b8651805182101561200f57611f808273ffffffffffffffffffffffffffffffffffffffff92613034565b511615611fe757611fb273ffffffffffffffffffffffffffffffffffffffff611faa838a51613034565b5116896135a9565b15611fbf57600101611f56565b6004897f12823a5e000000000000000000000000000000000000000000000000000000008152fd5b6004897fcfb6108a000000000000000000000000000000000000000000000000000000008152fd5b50509790957f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b8392975067ffffffffffffffff600196949560ff926002848651169101907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905561208582825116613549565b50511693519151169061209d60405192839283612de8565b0390a2019290611e92565b806120b56001928c613244565b90549060031b1c8c52826020528b604081205501611f47565b6004877f12823a5e000000000000000000000000000000000000000000000000000000008152fd5b511090508980611f20565b8780fd5b8680fd5b8480fd5b604067ffffffffffffffff61212b612126848688612f90565b612f7b565b16600090815260046020522054612145575b600101611e5f565b67ffffffffffffffff61215c612126838587612f90565b16600052600260205260406000208054600060018301905b82811061224a575050600082555060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556001906121ce67ffffffffffffffff6121c8612126848789612f90565b166132ae565b506121dd612126828587612f90565b7f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b8361223a60209267ffffffffffffffff6040519161221b8684612c31565b6000835260003681376000604051948594604086526040860190612d9e565b9684015216930390a2905061213d565b8061225760019286613244565b90549060031b1c600052826020526000604081205501612174565b346101805760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff8111610180578036036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610180576122e9612d2e565b5060843567ffffffffffffffff81116101805761230a903690600401612e09565b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd610124830135910181121561018057810160048101359067ffffffffffffffff8211610180576024019080360382136101805760246123aa9167ffffffffffffffff93357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110612572575b505060601c9301612f7b565b1680600052600560205260406000209081549073ffffffffffffffffffffffffffffffffffffffff8216908115611489576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa908115611dd65760009161250f575b5073ffffffffffffffffffffffffffffffffffffffff1633036124e15760f01c60ff16612497575b610f816040517f49ff34ed00000000000000000000000000000000000000000000000000000000602082015260048152612483602482612c31565b604051918291602083526020830190612cac565b600082815260029091016020526040902054156124b45780612448565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d60201161256a575b8161252860209383612c31565b8101031261067457519073ffffffffffffffffffffffffffffffffffffffff82168203610afc575073ffffffffffffffffffffffffffffffffffffffff612420565b3d915061251b565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b1616858061239e565b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805767ffffffffffffffff6125e7612d72565b166000526002602052604060002060405190818092602083549182815201908360005260206000209060005b81811061263f578460ff60028861262c84890385612c31565b01541690610f8160405192839283612de8565b8254845286945060209093019260019283019201612613565b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101805760043567ffffffffffffffff8111610180573660238201121561018057806004013567ffffffffffffffff811161018057602460c08202830101368111610180576126d06131de565b6126d982612e68565b916126e76040519384612c31565b825260009260240190602083015b818310612918578480855b8051831015612914576127138382613034565b519267ffffffffffffffff602061272a8385613034565b510151169384156128e8578484526005602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156128bc5773ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a2019190612700565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c083360312612109576040519060c0820182811067ffffffffffffffff8211176129cc57604052833573ffffffffffffffffffffffffffffffffffffffff8116810361210557825261296d60208501612d89565b602083015261297e60408501612f5d565b604083015260608401359061ffff821682036121055782602092606060c09501526129ab60808701612f6a565b60808201526129bc60a08701612f6a565b60a08201528152019201916126f5565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057612a30612d0b565b73ffffffffffffffffffffffffffffffffffffffff6009541690813303611a8b5773ffffffffffffffffffffffffffffffffffffffff169033821461029d57817fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fdfee8caf308a35b723489c72952cf11683462281c34aa62e8af474dcd012f41a600080a3005b346101805760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057610f816040805190612b018183612c31565b601b82527f436f6d6d6974746565566572696669657220312e372e302d6465760000000000602083015251918291602083526020830190612cac565b346101805760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018057600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361018057817f83adcde10000000000000000000000000000000000000000000000000000000060209314908115612bcf575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612bc8565b6060810190811067ffffffffffffffff8211176114ee57604052565b6040810190811067ffffffffffffffff8211176114ee57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176114ee57604052565b67ffffffffffffffff81116114ee57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110612cf65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612cb7565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361018057565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361018057565b359073ffffffffffffffffffffffffffffffffffffffff8216820361018057565b6004359067ffffffffffffffff8216820361018057565b359067ffffffffffffffff8216820361018057565b906020808351928381520192019060005b818110612dbc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612daf565b9060ff612e02602092959495604085526040850190612d9e565b9416910152565b9181601f840112156101805782359167ffffffffffffffff8311610180576020838186019501011161018057565b9181601f840112156101805782359167ffffffffffffffff8311610180576020808501948460051b01011161018057565b67ffffffffffffffff81116114ee5760051b60200190565b929192612e8c82612c72565b91612e9a6040519384612c31565b829481845281830111610180578281602093846000960137010152565b9080601f8301121561018057816020612ed293359101612e80565b90565b6064359061ffff8216820361018057565b9080602083519182815201916020808360051b8301019401926000915b838310612f1257505050505090565b9091929394602080612f4e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951612cac565b97019301930191939290612f03565b3590811515820361018057565b359063ffffffff8216820361018057565b3567ffffffffffffffff811681036101805790565b9190811015612fa05760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f83011215610180578135612fe681612e68565b92612ff46040519485612c31565b81845260208085019260051b82010192831161018057602001905b82821061301c5750505090565b6020809161302984612d51565b81520191019061300f565b8051821015612fa05760209160051b010190565b90600182811c92168015613091575b602083101461306257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613057565b600654906130a882612e68565b916130b66040519384612c31565b808352600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9190602085015b8282106130f55750505050565b6040516000855461310581613048565b8084529060018116908115613177575060011461313f575b506001928261313185946020940382612c31565b8152019401910190926130e8565b6000878152602081209092505b8183106131615750508101602001600161311d565b600181602092548386880101520192019161314c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b840190910191506001905061311d565b6006019081600611610a4557565b90939293848311610180578411610180578101920390565b73ffffffffffffffffffffffffffffffffffffffff6001541633036131ff57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b600654811015612fa057600660005260206000200190600090565b8054821015612fa05760005260206000200190600090565b818110613267575050565b6000815560010161325c565b359060208110613281575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b600081815260046020526040902054801561340e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610a4557600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610a455781810361339f575b5050506003548015611a5c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161335c816003613244565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b6133f66133b06133c1936003613244565b90549060031b1c9283926003613244565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080613323565b5050600090565b9060018201918160005282602052604060002054801515600014613540577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610a45578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610a4557818103613509575b50505080548015611a5c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906134ca8282613244565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6135296135196133c19386613244565b90549060031b1c92839286613244565b905560005283602052604060002055388080613492565b50505050600090565b806000526004602052604060002054156000146135a357600354680100000000000000008110156114ee5761358a6133c18260018594016003556003613244565b9055600354906000526004602052604060002055600190565b50600090565b600082815260018201602052604090205461340e57805490680100000000000000008210156114ee57826135e76133c1846001809601855584613244565b905580549260005201602052604060002055600190565b919290156136795750815115613612575090565b3b1561361b5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561368c5750805190602001fd5b6136ca906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612cac565b0390fdfea164736f6c634300081a000a",
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

func (_CommitteeVerifier *CommitteeVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
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

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

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

func (_CommitteeVerifier *CommitteeVerifierCaller) GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getSignatureConfig", sourceChainSelector)

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetSignatureConfig(sourceChainSelector uint64) ([]common.Address, uint8, error) {
	return _CommitteeVerifier.Contract.GetSignatureConfig(&_CommitteeVerifier.CallOpts, sourceChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetSignatureConfig(sourceChainSelector uint64) ([]common.Address, uint8, error) {
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
	Senders           []common.Address
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
	Senders           []common.Address
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
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}

func (CommitteeVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitteeVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitteeVerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0x781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd")
}

func (CommitteeVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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

	GetDynamicConfig(opts *bind.CallOpts) (CommitteeVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetPendingStorageLocationsAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) ([]common.Address, uint8, error)

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

	FilterConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitteeVerifierConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitteeVerifierFeeTokenWithdrawn, error)

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
