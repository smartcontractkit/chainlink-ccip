// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_helper

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

var VerifierTestHelperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"testRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionTag\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTestRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"tag\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MessageCannotHaveSideEffects\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustUseAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustUseTestRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"VersionTagCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346104e257612c6a803803809161001c828561055f565b833981016080828203126104e25761003382610582565b60208301519092906001600160401b0381116104e257810182601f820112156104e25780519061006282610596565b93610070604051958661055f565b82855260208086019360051b830101918183116104e25760208101935b8385106104e75787878760606100a560408301610582565b9101519063ffffffff60e01b8216928383036104e2576001549080516100ca83610596565b926100d8604051948561055f565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b83821061043d5750505060005b8181106103a957505060005b81811061021a5750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586916101766101689260405193849360408552604085019061063c565b90838203602085015261063c565b0390a16001600160a01b031691821561020957156101f85760a05260805233156101e757600380546001600160a01b0319163317905560c0526040516125bc90816106ae82396080518161202f015260a0518181816101260152611a73015260c0518181816105fc01526114c60152f35b639b15e16f60e01b60005260046000fd5b631027401f60e21b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b82518110156103935760208160051b84010151600154680100000000000000008110156103675780600161025192016001556105d0565b91909161037d578051906001600160401b0382116103675761027383546105eb565b601f811161032a575b50602090601f83116001146102bf57600194939291600091836102b4575b5050600019600383901b1c191690841b1790555b01610122565b015190508c8061029a565b90601f1983169184600052816000209260005b8181106103125750916001969594929183889593106102f9575b505050811b0190556102ae565b015160001960f88460031b161c191690558c80806102ec565b929360206001819287860151815501950193016102d2565b61035790846000526020600020601f850160051c8101916020861061035d575b601f0160051c0190610625565b8b61027c565b909150819061034a565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60015480156104275760001901906103c0826105d0565b92909261037d57826103d4600194546105eb565b90816103e5575b5050825501610116565b81601f6000931186146103fc5750555b8b806103db565b8183526020832061041791601f0160051c8101908701610625565b80825281602081209155556103f5565b634e487b7160e01b600052603160045260246000fd5b6040516000845461044d816105eb565b80845290600181169081156104bf5750600114610487575b50600192826104798594602094038261055f565b815201930191019091610109565b6000868152602081209092505b8183106104a957505081016020016001610465565b6001816020925483868801015201920191610494565b60ff191660208581019190915291151560051b8401909101915060019050610465565b600080fd5b84516001600160401b0381116104e25782019083603f830112156104e2576020820151906001600160401b03821161036757604051610530601f8401601f19166020018261055f565b82815260408484010186106104e257610554602094938594604086850191016105ad565b81520194019361008d565b601f909101601f19168101906001600160401b0382119082101761036757604052565b51906001600160a01b03821682036104e257565b6001600160401b0381116103675760051b60200190565b60005b8381106105c05750506000910152565b81810151838201526020016105b0565b60015481101561039357600160005260206000200190600090565b90600182811c9216801561061b575b602083101461060557565b634e487b7160e01b600052602260045260246000fd5b91607f16916105fa565b818110610630575050565b60008155600101610625565b9080602083519182815201916020808360051b8301019401926000915b83831061066857505050505090565b909192939460208080600193601f198682030187528951610694815180928185528580860191016105ad565b601f01601f19160101970195949190910192019061065956fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611bbc57508063181f5a7714611b3f578063296947061461188d578063597b95c31461144a5780635cb80c5d1461120557806379ba5097146110f957806387ae929214610eec578063898068fc14610d0157806389e364c714610a6c5780638da5cb5b14610a1a578063c9b146b314610620578063e0d9ef59146105b1578063ec6ae7a714610550578063f2fde38b1461043e578063f4cdd89e1461014f5763fe163eed146100cf57600080fd5b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760206040517fffffffff000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000168152f35b600080fd5b3461014a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57610186611d4c565b60243567ffffffffffffffff811161014a5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261014a576040519060a0820182811067ffffffffffffffff82111761040f57604052806004013567ffffffffffffffff811161014a576102069060043691840101611dd5565b8252602481013567ffffffffffffffff811161014a5761022c9060043691840101611dd5565b6020830152604481013567ffffffffffffffff811161014a5781013660238201121561014a57600481013561026081611e4a565b9161026e6040519384611d94565b818352602060048185019360061b830101019036821161014a57602401915b8183106103c35750505060408301526102a860648201611cfa565b6060830152608481013567ffffffffffffffff811161014a5760809160046102d39236920101611dd5565b91015260443567ffffffffffffffff811161014a576102f6903690600401611dd5565b50606435907fffffffff000000000000000000000000000000000000000000000000000000008216820361014a5767ffffffffffffffff168060005260006020526040600020549073ffffffffffffffffffffffffffffffffffffffff821615610396575061036d60609260025460e01b906121b2565b61ffff60405191818160a01c16835263ffffffff8160b01c16602084015260d01c166040820152f35b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60408336031261014a5760405190604082019082821067ffffffffffffffff83111761040f5760409260209284526103fa86611cfa565b8152828601358382015281520192019161028d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043573ffffffffffffffffffffffffffffffffffffffff81169081810361014a576104976120bf565b338214610526577fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff000000006002549260201b1691161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57602060025460e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043567ffffffffffffffff811161014a5761066f903690600401611d1b565b6106776120bf565b6000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b80841015610a18576000938060051b83013584811215610a14578301608081360312610a1457604051956080870187811067ffffffffffffffff8211176109e7576040526106ee82611d63565b87526106fc60208301611ec8565b9360208801948552604083013567ffffffffffffffff81116109e357610725903690850161214d565b9260408901938452606081013567ffffffffffffffff81116109df5761074d9136910161214d565b966060890197885267ffffffffffffffff895116835282602052604083209160ff835460e01c16875115158091151503610956575b50600184989301975b89518051821015610810579073ffffffffffffffffffffffffffffffffffffffff6107b88260019461210a565b51168c6107c5828d6123bb565b6107d2575b50500161078b565b602067ffffffffffffffff7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d8292511692604051908152a28c8c6107ca565b5050975094959095835151610834575b5050506001919091019450909290506106a1565b9690949195929651151560001461091f57855b87518051821015610907576108718273ffffffffffffffffffffffffffffffffffffffff9261210a565b511680156108d057908161088760019389612553565b610893575b5001610847565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80602067ffffffffffffffff8d511692604051908152a28a61088c565b60248867ffffffffffffffff8c51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509650935093506001915084939291868080610820565b60248667ffffffffffffffff8a51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b83547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560e01b7cff00000000000000000000000000000000000000000000000000000000161784557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492602067ffffffffffffffff8d511692604051908152a28a610782565b8380fd5b8280fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b3461014a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043567ffffffffffffffff811161014a57806004016101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261014a5760443567ffffffffffffffff811161014a57610afe903690600401611dd5565b50610b10610b0b82611e62565b611fc9565b67ffffffffffffffff610b2282611e62565b1680600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000205416908115610396576020906044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115610cf557600091610cc6575b5015610c985780610bad610bda92611ee4565b6101446040610bbe60248601611e62565b67ffffffffffffffff6000911681528060205220930190611e77565b90357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110610c63575b505060601c9060ff815460e01c16610c1d57005b60008281526002909101602052604090205415610c3657005b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168280610c09565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b610ce8915060203d602011610cee575b610ce08183611d94565b810190611fb1565b83610b9a565b503d610cd6565b6040513d6000823e3d90fd5b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5767ffffffffffffffff610d41611d4c565b600060a0604051610d5181611d78565b828152826020820152826040820152826060820152826080820152015216806000526000602052604060002080549160405190610d8d82611d78565b73ffffffffffffffffffffffffffffffffffffffff8416825260208201908152604082019360ff8160e01c1615158552606083019361ffff8260a01c1685526001608085019163ffffffff8460b01c16835261ffff60a087019460d01c168452019460405196879460208854998a81520198899860005260206000209060005b818110610ed6575073ffffffffffffffffffffffffffffffffffffffff60e08b8b8b8f8c63ffffffff8d61ffff8e8e67ffffffffffffffff8f610e538a869a038b611d94565b6040519d8e9d8e019b51168d52511660208c015251151560408b01525116606089015251166080870152511660a085015260e060c0850152518091526101008301919060005b818110610ea7575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610e99565b82548c526020909b019a60019283019201610e0d565b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57600154610f2781611e4a565b90610f356040519283611d94565b80825260208201809160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6916000905b828210610ff357848660405191829160208301906020845251809152604083019060408160051b85010192916000905b828210610fa857505050500390f35b91936020610fe3827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851611c78565b9601920192018594939192610f99565b604051600085548060011c906001811680156110ef575b6020831081146110c2578285529081156110815750600114611049575b506001928261103b85946020940382611d94565b815201940191019092610f69565b6000878152602081209092505b81831061106b57505081016020016001611027565b6001816020925483868801015201920191611056565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611027565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f169161100a565b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760025473ffffffffffffffffffffffffffffffffffffffff8160201c1633036111db577fffffffffffffffff0000000000000000000000000000000000000000ffffffff60035491337fffffffffffffffffffffffff00000000000000000000000000000000000000008416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043567ffffffffffffffff811161014a57611254903690600401611d1b565b73ffffffffffffffffffffffffffffffffffffffff6003541680156114205760005b82811015610a185760008160051b8501359073ffffffffffffffffffffffffffffffffffffffff821680920361141d57604051907f70a08231000000000000000000000000000000000000000000000000000000008252306004830152602082602481865afa9182156114105781926113d9575b50816112fc575b505050600101611276565b602081604051828101907fa9059cbb00000000000000000000000000000000000000000000000000000000825288602482015285604482015260448152611344606482611d94565b519082875af1156113cd5780513d6113c45750823b155b611398575090837f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39085806112f1565b80837f5274afe70000000000000000000000000000000000000000000000000000000060249352600452fd5b6001141561135b565b604051903d90823e3d90fd5b9091506020813d8211611408575b816113f460209383611d94565b81010312611404575190876112ea565b5080fd5b3d91506113e7565b50604051903d90823e3d90fd5b80fd5b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043567ffffffffffffffff811161014a573660238201121561014a5780600401359067ffffffffffffffff821161014a57602460c0830282010136811161014a57916114c46120bf565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169060005b818110156115a357600060c08202850160648101358015908115036109e35761157b576024013573ffffffffffffffffffffffffffffffffffffffff811680910361140457840361155357506001016114ff565b807f1b16420e0000000000000000000000000000000000000000000000000000000060049252fd5b6004827ffc0a3f38000000000000000000000000000000000000000000000000000000008152fd5b8382866115af82611e4a565b916115bd6040519384611d94565b825260009260240190602083015b8183106117e0578480855b80518310156117dc576115e9838261210a565b519267ffffffffffffffff6020850151169384156117b0578484526020849052604080852082518154928401517fffffff00ffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560e01b7cff000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156117845773ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007bffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260e01c1615156020820152a20191906115d6565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261188957604051906117f782611d78565b833573ffffffffffffffffffffffffffffffffffffffff8116810361188557825261182460208501611d63565b602083015261183560408501611ec8565b604083015261184660608501611ed5565b606083015260808401359063ffffffff821682036118855782602092608060c095015261187560a08701611ed5565b60a08201528152019201916115cb565b8680fd5b8480fd5b3461014a5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a5760043567ffffffffffffffff811161014a57806004016101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261014a57611907611cd7565b5060843567ffffffffffffffff811161014a573660238201121561014a57806004013567ffffffffffffffff811161014a573691016024011161014a5761196e81611953602093611ee4565b6101246024850194611967610b0b87611e62565b0190611e77565b908092918101031261014a57359073ffffffffffffffffffffffffffffffffffffffff821680920361014a576119ac67ffffffffffffffff91611e62565b1680600052600060205260406000209081549073ffffffffffffffffffffffffffffffffffffffff8216908115610396576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa908115610cf557600091611adc575b5073ffffffffffffffffffffffffffffffffffffffff163303610c985760e01c60ff16611abf575b611abb6040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000016602082015260048152611aa7602482611d94565b604051918291602083526020830190611c78565b0390f35b60008281526002909101602052604090205415610c365780611a4a565b6020813d602011611b37575b81611af560209383611d94565b8101031261140457519073ffffffffffffffffffffffffffffffffffffffff8216820361141d575073ffffffffffffffffffffffffffffffffffffffff611a22565b3d9150611ae8565b3461014a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57611abb6040805190611b808183611d94565b601882527f56657269666965725465737448656c70657220322e302e300000000000000000602083015251918291602083526020830190611c78565b3461014a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014a57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361014a57817fd3e969cd0000000000000000000000000000000000000000000000000000000060209314908115611c4e575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611c47565b919082519283825260005b848110611cc25750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611c83565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361014a57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361014a57565b9181601f8401121561014a5782359167ffffffffffffffff831161014a576020808501948460051b01011161014a57565b6004359067ffffffffffffffff8216820361014a57565b359067ffffffffffffffff8216820361014a57565b60c0810190811067ffffffffffffffff82111761040f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761040f57604052565b81601f8201121561014a5780359067ffffffffffffffff821161040f5760405192611e28601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185611d94565b8284526020838301011161014a57816000926020809301838601378301015290565b67ffffffffffffffff811161040f5760051b60200190565b3567ffffffffffffffff8116810361014a5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561014a570180359067ffffffffffffffff821161014a5760200191813603831361014a57565b3590811515820361014a57565b359061ffff8216820361014a57565b6101808101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561014a57810180359067ffffffffffffffff821161014a5760208260051b360391011361014a571590811591611f97575b8115611f7b575b50611f5157565b7fec14a34d0000000000000000000000000000000000000000000000000000000060005260046000fd5b60809150013563ffffffff811680910361014a57151538611f4a565b9050611fa76101a0820182611e77565b9050151590611f43565b9081602091031261014a5751801515810361014a5790565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610cf5576000916120a0575b506120685750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6120b9915060203d602011610cee57610ce08183611d94565b38612060565b73ffffffffffffffffffffffffffffffffffffffff6003541633036120e057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805182101561211e5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561014a57813561216481611e4a565b926121726040519485611d94565b81845260208085019260051b82010192831161014a57602001905b82821061219a5750505090565b602080916121a784611cfa565b81520191019061218d565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115612294576121e581612299565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c16166122945761ffff8360e01c168015918215612283575b505061222f575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff161090503880612225565b505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811690811561239f577dffff000000000000000000000000000000000000000000000000000000008116156123965760ff60015b169060f01c80612331575b506001036123045750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b6010811061234257506122f9565b6001811b8216612355575b600101612334565b9160018101809111612367579161234d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff60006122ee565b5050565b805482101561211e5760005260206000200190600090565b906001820191816000528260205260406000205480151560001461254a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111612367578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211612367578181036124de575b505050805480156124af577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061247082826123a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6125336124ee6124fe93866123a3565b90549060031b1c928392866123a3565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005283602052604060002055388080612438565b50505050600090565b60008281526001820160205260409020546125a8578054906801000000000000000082101561040f57826125916124fe8460018096018555846123a3565b905580549260005201602052604060002055600190565b505060009056fea164736f6c634300081a000a",
}

var VerifierTestHelperABI = VerifierTestHelperMetaData.ABI

var VerifierTestHelperBin = VerifierTestHelperMetaData.Bin

func DeployVerifierTestHelper(auth *bind.TransactOpts, backend bind.ContractBackend, testRouter common.Address, storageLocations []string, rmn common.Address, versionTag [4]byte) (common.Address, *types.Transaction, *VerifierTestHelper, error) {
	parsed, err := VerifierTestHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierTestHelperBin), backend, testRouter, storageLocations, rmn, versionTag)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierTestHelper{address: address, abi: *parsed, VerifierTestHelperCaller: VerifierTestHelperCaller{contract: contract}, VerifierTestHelperTransactor: VerifierTestHelperTransactor{contract: contract}, VerifierTestHelperFilterer: VerifierTestHelperFilterer{contract: contract}}, nil
}

type VerifierTestHelper struct {
	address common.Address
	abi     abi.ABI
	VerifierTestHelperCaller
	VerifierTestHelperTransactor
	VerifierTestHelperFilterer
}

type VerifierTestHelperCaller struct {
	contract *bind.BoundContract
}

type VerifierTestHelperTransactor struct {
	contract *bind.BoundContract
}

type VerifierTestHelperFilterer struct {
	contract *bind.BoundContract
}

type VerifierTestHelperSession struct {
	Contract     *VerifierTestHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierTestHelperCallerSession struct {
	Contract *VerifierTestHelperCaller
	CallOpts bind.CallOpts
}

type VerifierTestHelperTransactorSession struct {
	Contract     *VerifierTestHelperTransactor
	TransactOpts bind.TransactOpts
}

type VerifierTestHelperRaw struct {
	Contract *VerifierTestHelper
}

type VerifierTestHelperCallerRaw struct {
	Contract *VerifierTestHelperCaller
}

type VerifierTestHelperTransactorRaw struct {
	Contract *VerifierTestHelperTransactor
}

func NewVerifierTestHelper(address common.Address, backend bind.ContractBackend) (*VerifierTestHelper, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierTestHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierTestHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelper{address: address, abi: abi, VerifierTestHelperCaller: VerifierTestHelperCaller{contract: contract}, VerifierTestHelperTransactor: VerifierTestHelperTransactor{contract: contract}, VerifierTestHelperFilterer: VerifierTestHelperFilterer{contract: contract}}, nil
}

func NewVerifierTestHelperCaller(address common.Address, caller bind.ContractCaller) (*VerifierTestHelperCaller, error) {
	contract, err := bindVerifierTestHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperCaller{contract: contract}, nil
}

func NewVerifierTestHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTestHelperTransactor, error) {
	contract, err := bindVerifierTestHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperTransactor{contract: contract}, nil
}

func NewVerifierTestHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierTestHelperFilterer, error) {
	contract, err := bindVerifierTestHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperFilterer{contract: contract}, nil
}

func bindVerifierTestHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierTestHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierTestHelper *VerifierTestHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierTestHelper.Contract.VerifierTestHelperCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierTestHelper *VerifierTestHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.VerifierTestHelperTransactor.contract.Transfer(opts)
}

func (_VerifierTestHelper *VerifierTestHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.VerifierTestHelperTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierTestHelper *VerifierTestHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierTestHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.contract.Transfer(opts)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "forwardToVerifier", message, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _VerifierTestHelper.Contract.ForwardToVerifier(&_VerifierTestHelper.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _VerifierTestHelper.Contract.ForwardToVerifier(&_VerifierTestHelper.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _VerifierTestHelper.Contract.GetAllowedFinalityConfig(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _VerifierTestHelper.Contract.GetAllowedFinalityConfig(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, requestedFinality)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _VerifierTestHelper.Contract.GetFee(&_VerifierTestHelper.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

	error) {
	return _VerifierTestHelper.Contract.GetFee(&_VerifierTestHelper.CallOpts, destChainSelector, arg1, arg2, requestedFinality)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemoteChainConfig = *abi.ConvertType(out[0], new(BaseVerifierRemoteChainConfigArgs)).(*BaseVerifierRemoteChainConfigArgs)
	outstruct.AllowedSendersList = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _VerifierTestHelper.Contract.GetRemoteChainConfig(&_VerifierTestHelper.CallOpts, remoteChainSelector)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _VerifierTestHelper.Contract.GetRemoteChainConfig(&_VerifierTestHelper.CallOpts, remoteChainSelector)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) GetStorageLocations(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "getStorageLocations")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) GetStorageLocations() ([]string, error) {
	return _VerifierTestHelper.Contract.GetStorageLocations(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) GetStorageLocations() ([]string, error) {
	return _VerifierTestHelper.Contract.GetStorageLocations(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) GetTestRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "getTestRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) GetTestRouter() (common.Address, error) {
	return _VerifierTestHelper.Contract.GetTestRouter(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) GetTestRouter() (common.Address, error) {
	return _VerifierTestHelper.Contract.GetTestRouter(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) Owner() (common.Address, error) {
	return _VerifierTestHelper.Contract.Owner(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) Owner() (common.Address, error) {
	return _VerifierTestHelper.Contract.Owner(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerifierTestHelper.Contract.SupportsInterface(&_VerifierTestHelper.CallOpts, interfaceId)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerifierTestHelper.Contract.SupportsInterface(&_VerifierTestHelper.CallOpts, interfaceId)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) TypeAndVersion() (string, error) {
	return _VerifierTestHelper.Contract.TypeAndVersion(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) TypeAndVersion() (string, error) {
	return _VerifierTestHelper.Contract.TypeAndVersion(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 []byte) error {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "verifyMessage", message, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

func (_VerifierTestHelper *VerifierTestHelperSession) VerifyMessage(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 []byte) error {
	return _VerifierTestHelper.Contract.VerifyMessage(&_VerifierTestHelper.CallOpts, message, arg1, arg2)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) VerifyMessage(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 []byte) error {
	return _VerifierTestHelper.Contract.VerifyMessage(&_VerifierTestHelper.CallOpts, message, arg1, arg2)
}

func (_VerifierTestHelper *VerifierTestHelperCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _VerifierTestHelper.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_VerifierTestHelper *VerifierTestHelperSession) VersionTag() ([4]byte, error) {
	return _VerifierTestHelper.Contract.VersionTag(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperCallerSession) VersionTag() ([4]byte, error) {
	return _VerifierTestHelper.Contract.VersionTag(&_VerifierTestHelper.CallOpts)
}

func (_VerifierTestHelper *VerifierTestHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierTestHelper.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierTestHelper *VerifierTestHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.AcceptOwnership(&_VerifierTestHelper.TransactOpts)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.AcceptOwnership(&_VerifierTestHelper.TransactOpts)
}

func (_VerifierTestHelper *VerifierTestHelperTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_VerifierTestHelper *VerifierTestHelperSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.ApplyAllowlistUpdates(&_VerifierTestHelper.TransactOpts, allowlistConfigArgsItems)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.ApplyAllowlistUpdates(&_VerifierTestHelper.TransactOpts, allowlistConfigArgsItems)
}

func (_VerifierTestHelper *VerifierTestHelperTransactor) ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.contract.Transact(opts, "applyRemoteChainConfigUpdates", remoteChainConfigArgs)
}

func (_VerifierTestHelper *VerifierTestHelperSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.ApplyRemoteChainConfigUpdates(&_VerifierTestHelper.TransactOpts, remoteChainConfigArgs)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.ApplyRemoteChainConfigUpdates(&_VerifierTestHelper.TransactOpts, remoteChainConfigArgs)
}

func (_VerifierTestHelper *VerifierTestHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierTestHelper *VerifierTestHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.TransferOwnership(&_VerifierTestHelper.TransactOpts, to)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.TransferOwnership(&_VerifierTestHelper.TransactOpts, to)
}

func (_VerifierTestHelper *VerifierTestHelperTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_VerifierTestHelper *VerifierTestHelperSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.WithdrawFeeTokens(&_VerifierTestHelper.TransactOpts, feeTokens)
}

func (_VerifierTestHelper *VerifierTestHelperTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierTestHelper.Contract.WithdrawFeeTokens(&_VerifierTestHelper.TransactOpts, feeTokens)
}

type VerifierTestHelperAllowListSendersAddedIterator struct {
	Event *VerifierTestHelperAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperAllowListSendersAdded)
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
		it.Event = new(VerifierTestHelperAllowListSendersAdded)
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

func (it *VerifierTestHelperAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperAllowListSendersAddedIterator{contract: _VerifierTestHelper.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperAllowListSendersAdded)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseAllowListSendersAdded(log types.Log) (*VerifierTestHelperAllowListSendersAdded, error) {
	event := new(VerifierTestHelperAllowListSendersAdded)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperAllowListSendersRemovedIterator struct {
	Event *VerifierTestHelperAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperAllowListSendersRemoved)
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
		it.Event = new(VerifierTestHelperAllowListSendersRemoved)
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

func (it *VerifierTestHelperAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           common.Address
	Raw               types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperAllowListSendersRemovedIterator{contract: _VerifierTestHelper.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperAllowListSendersRemoved)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseAllowListSendersRemoved(log types.Log) (*VerifierTestHelperAllowListSendersRemoved, error) {
	event := new(VerifierTestHelperAllowListSendersRemoved)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperAllowListStateChangedIterator struct {
	Event *VerifierTestHelperAllowListStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperAllowListStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperAllowListStateChanged)
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
		it.Event = new(VerifierTestHelperAllowListStateChanged)
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

func (it *VerifierTestHelperAllowListStateChangedIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperAllowListStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperAllowListStateChanged struct {
	DestChainSelector uint64
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListStateChangedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperAllowListStateChangedIterator{contract: _VerifierTestHelper.contract, event: "AllowListStateChanged", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "AllowListStateChanged", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperAllowListStateChanged)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseAllowListStateChanged(log types.Log) (*VerifierTestHelperAllowListStateChanged, error) {
	event := new(VerifierTestHelperAllowListStateChanged)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "AllowListStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperFeeTokenWithdrawnIterator struct {
	Event *VerifierTestHelperFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperFeeTokenWithdrawn)
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
		it.Event = new(VerifierTestHelperFeeTokenWithdrawn)
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

func (it *VerifierTestHelperFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*VerifierTestHelperFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperFeeTokenWithdrawnIterator{contract: _VerifierTestHelper.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperFeeTokenWithdrawn)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseFeeTokenWithdrawn(log types.Log) (*VerifierTestHelperFeeTokenWithdrawn, error) {
	event := new(VerifierTestHelperFeeTokenWithdrawn)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperFinalityConfigSetIterator struct {
	Event *VerifierTestHelperFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperFinalityConfigSet)
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
		it.Event = new(VerifierTestHelperFinalityConfigSet)
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

func (it *VerifierTestHelperFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperFinalityConfigSet struct {
	AllowedFinality [4]byte
	Raw             types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*VerifierTestHelperFinalityConfigSetIterator, error) {

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperFinalityConfigSetIterator{contract: _VerifierTestHelper.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperFinalityConfigSet)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseFinalityConfigSet(log types.Log) (*VerifierTestHelperFinalityConfigSet, error) {
	event := new(VerifierTestHelperFinalityConfigSet)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperOwnershipTransferRequestedIterator struct {
	Event *VerifierTestHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperOwnershipTransferRequested)
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
		it.Event = new(VerifierTestHelperOwnershipTransferRequested)
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

func (it *VerifierTestHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierTestHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperOwnershipTransferRequestedIterator{contract: _VerifierTestHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperOwnershipTransferRequested)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierTestHelperOwnershipTransferRequested, error) {
	event := new(VerifierTestHelperOwnershipTransferRequested)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperOwnershipTransferredIterator struct {
	Event *VerifierTestHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperOwnershipTransferred)
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
		it.Event = new(VerifierTestHelperOwnershipTransferred)
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

func (it *VerifierTestHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierTestHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperOwnershipTransferredIterator{contract: _VerifierTestHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperOwnershipTransferred)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierTestHelperOwnershipTransferred, error) {
	event := new(VerifierTestHelperOwnershipTransferred)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperRemoteChainConfigSetIterator struct {
	Event *VerifierTestHelperRemoteChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperRemoteChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperRemoteChainConfigSet)
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
		it.Event = new(VerifierTestHelperRemoteChainConfigSet)
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

func (it *VerifierTestHelperRemoteChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperRemoteChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperRemoteChainConfigSet struct {
	RemoteChainSelector uint64
	Router              common.Address
	AllowlistEnabled    bool
	Raw                 types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*VerifierTestHelperRemoteChainConfigSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperRemoteChainConfigSetIterator{contract: _VerifierTestHelper.contract, event: "RemoteChainConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperRemoteChainConfigSet)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseRemoteChainConfigSet(log types.Log) (*VerifierTestHelperRemoteChainConfigSet, error) {
	event := new(VerifierTestHelperRemoteChainConfigSet)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierTestHelperStorageLocationsUpdatedIterator struct {
	Event *VerifierTestHelperStorageLocationsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierTestHelperStorageLocationsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierTestHelperStorageLocationsUpdated)
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
		it.Event = new(VerifierTestHelperStorageLocationsUpdated)
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

func (it *VerifierTestHelperStorageLocationsUpdatedIterator) Error() error {
	return it.fail
}

func (it *VerifierTestHelperStorageLocationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierTestHelperStorageLocationsUpdated struct {
	OldLocations []string
	NewLocations []string
	Raw          types.Log
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*VerifierTestHelperStorageLocationsUpdatedIterator, error) {

	logs, sub, err := _VerifierTestHelper.contract.FilterLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return &VerifierTestHelperStorageLocationsUpdatedIterator{contract: _VerifierTestHelper.contract, event: "StorageLocationsUpdated", logs: logs, sub: sub}, nil
}

func (_VerifierTestHelper *VerifierTestHelperFilterer) WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperStorageLocationsUpdated) (event.Subscription, error) {

	logs, sub, err := _VerifierTestHelper.contract.WatchLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierTestHelperStorageLocationsUpdated)
				if err := _VerifierTestHelper.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (_VerifierTestHelper *VerifierTestHelperFilterer) ParseStorageLocationsUpdated(log types.Log) (*VerifierTestHelperStorageLocationsUpdated, error) {
	event := new(VerifierTestHelperStorageLocationsUpdated)
	if err := _VerifierTestHelper.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (VerifierTestHelperAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80")
}

func (VerifierTestHelperAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0x9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d82")
}

func (VerifierTestHelperAllowListStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492")
}

func (VerifierTestHelperFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (VerifierTestHelperFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
}

func (VerifierTestHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierTestHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierTestHelperRemoteChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9")
}

func (VerifierTestHelperStorageLocationsUpdated) Topic() common.Hash {
	return common.HexToHash("0xec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586")
}

func (_VerifierTestHelper *VerifierTestHelper) Address() common.Address {
	return _VerifierTestHelper.address
}

type VerifierTestHelperInterface interface {
	ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error)

	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, requestedFinality [4]byte) (GetFee,

		error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetStorageLocations(opts *bind.CallOpts) ([]string, error)

	GetTestRouter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 []byte) error

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*VerifierTestHelperAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*VerifierTestHelperAllowListSendersRemoved, error)

	FilterAllowListStateChanged(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierTestHelperAllowListStateChangedIterator, error)

	WatchAllowListStateChanged(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperAllowListStateChanged, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListStateChanged(log types.Log) (*VerifierTestHelperAllowListStateChanged, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*VerifierTestHelperFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*VerifierTestHelperFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*VerifierTestHelperFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*VerifierTestHelperFinalityConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierTestHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierTestHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierTestHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierTestHelperOwnershipTransferred, error)

	FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*VerifierTestHelperRemoteChainConfigSetIterator, error)

	WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigSet(log types.Log) (*VerifierTestHelperRemoteChainConfigSet, error)

	FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*VerifierTestHelperStorageLocationsUpdatedIterator, error)

	WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *VerifierTestHelperStorageLocationsUpdated) (event.Subscription, error)

	ParseStorageLocationsUpdated(log types.Log) (*VerifierTestHelperStorageLocationsUpdated, error)

	Address() common.Address
}
