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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"storageLocations\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionTag\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"tag\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListStateChanged\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MessageCannotHaveSideEffects\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"VersionTagCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c080604052346104c257612ad8803803809161001c828561053f565b833981016060828203126104c25781516001600160401b0381116104c257820181601f820112156104c25780519061005382610562565b92610061604051948561053f565b82845260208085019360051b830101918183116104c25760208101935b8385106104c75760208701516001600160a01b03811690879089908390036104c25760400151906001600160e01b03198216908183036104c2576001549080516100c783610562565b926100d5604051948561053f565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b83821061041d5750505060005b81811061038957505060005b8181106101fa5750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161017361016592604051938493604085526040850190610608565b908382036020850152610608565b0390a182156101e957156101d85760a05260805233156101c757600380546001600160a01b0319163317905560405161245e908161067a823960805181611ed1015260a05181818161011b01526119150152f35b639b15e16f60e01b60005260046000fd5b631027401f60e21b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b82518110156103735760208160051b840101516001546801000000000000000081101561034757806001610231920160015561059c565b91909161035d578051906001600160401b0382116103475761025383546105b7565b601f811161030a575b50602090601f831160011461029f5760019493929160009183610294575b5050600019600383901b1c191690841b1790555b0161011f565b015190508b8061027a565b90601f1983169184600052816000209260005b8181106102f25750916001969594929183889593106102d9575b505050811b01905561028e565b015160001960f88460031b161c191690558b80806102cc565b929360206001819287860151815501950193016102b2565b61033790846000526020600020601f850160051c8101916020861061033d575b601f0160051c01906105f1565b8a61025c565b909150819061032a565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60015480156104075760001901906103a08261059c565b92909261035d57826103b4600194546105b7565b90816103c5575b5050825501610113565b81601f6000931186146103dc5750555b8a806103bb565b818352602083206103f791601f0160051c81019087016105f1565b80825281602081209155556103d5565b634e487b7160e01b600052603160045260246000fd5b6040516000845461042d816105b7565b808452906001811690811561049f5750600114610467575b50600192826104598594602094038261053f565b815201930191019091610106565b6000868152602081209092505b81831061048957505081016020016001610445565b6001816020925483868801015201920191610474565b60ff191660208581019190915291151560051b8401909101915060019050610445565b600080fd5b84516001600160401b0381116104c25782019083603f830112156104c2576020820151906001600160401b03821161034757604051610510601f8401601f19166020018261053f565b82815260408484010186106104c25761053460209493859460408685019101610579565b81520194019361007e565b601f909101601f19168101906001600160401b0382119082101761034757604052565b6001600160401b0381116103475760051b60200190565b60005b83811061058c5750506000910152565b818101518382015260200161057c565b60015481101561037357600160005260206000200190600090565b90600182811c921680156105e7575b60208310146105d157565b634e487b7160e01b600052602260045260246000fd5b91607f16916105c6565b8181106105fc575050565b600081556001016105f1565b9080602083519182815201916020808360051b8301019401926000915b83831061063457505050505090565b909192939460208080600193601f19868203018752895161066081518092818552858086019101610579565b601f01601f19160101970195949190910192019061062556fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611a5e57508063181f5a77146119e1578063296947061461172f578063597b95c3146113d05780635cb80c5d1461118b57806379ba50971461107f57806387ae929214610e72578063898068fc14610c8757806389e364c7146109f25780638da5cb5b146109a0578063c9b146b3146105a6578063ec6ae7a714610545578063f2fde38b14610433578063f4cdd89e146101445763fe163eed146100c457600080fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760206040517fffffffff000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000168152f35b600080fd5b3461013f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761017b611bee565b60243567ffffffffffffffff811161013f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261013f576040519060a0820182811067ffffffffffffffff82111761040457604052806004013567ffffffffffffffff811161013f576101fb9060043691840101611c77565b8252602481013567ffffffffffffffff811161013f576102219060043691840101611c77565b6020830152604481013567ffffffffffffffff811161013f5781013660238201121561013f57600481013561025581611cec565b916102636040519384611c36565b818352602060048185019360061b830101019036821161013f57602401915b8183106103b857505050604083015261029d60648201611b9c565b6060830152608481013567ffffffffffffffff811161013f5760809160046102c89236920101611c77565b91015260443567ffffffffffffffff811161013f576102eb903690600401611c77565b50606435907fffffffff000000000000000000000000000000000000000000000000000000008216820361013f5767ffffffffffffffff168060005260006020526040600020549073ffffffffffffffffffffffffffffffffffffffff82161561038b575061036260609260025460e01b90612054565b61ffff60405191818160a01c16835263ffffffff8160b01c16602084015260d01c166040820152f35b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60408336031261013f5760405190604082019082821067ffffffffffffffff8311176104045760409260209284526103ef86611b9c565b81528286013583820152815201920191610282565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043573ffffffffffffffffffffffffffffffffffffffff81169081810361013f5761048c611f61565b33821461051b577fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff000000006002549260201b1691161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602060025460e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f576105f5903690600401611bbd565b6105fd611f61565b6000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b8084101561099e576000938060051b8301358481121561099a57830160808136031261099a57604051956080870187811067ffffffffffffffff82111761096d5760405261067482611c05565b875261068260208301611d6a565b9360208801948552604083013567ffffffffffffffff8111610969576106ab9036908501611fef565b9260408901938452606081013567ffffffffffffffff8111610965576106d391369101611fef565b966060890197885267ffffffffffffffff895116835282602052604083209160ff835460e01c168751151580911515036108dc575b50600184989301975b89518051821015610796579073ffffffffffffffffffffffffffffffffffffffff61073e82600194611fac565b51168c61074b828d61225d565b610758575b505001610711565b602067ffffffffffffffff7f9ac16e02c9a455144d35e2f0d80817a608340dee3c104f547ceb4433df418d8292511692604051908152a28c8c610750565b50509750949590958351516107ba575b505050600191909101945090929050610627565b969094919592965115156000146108a557855b8751805182101561088d576107f78273ffffffffffffffffffffffffffffffffffffffff92611fac565b5116801561085657908161080d600193896123f5565b610819575b50016107cd565b7f85682793ee26ba7d2d073ce790a50b388a1791aab25fc368bcce99d3b1d4da80602067ffffffffffffffff8d511692604051908152a28a610812565b60248867ffffffffffffffff8c51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b505096509350935060019150849392918680806107a6565b60248667ffffffffffffffff8a51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b83547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681151560e01b7cff00000000000000000000000000000000000000000000000000000000161784557f8504171b9fc8a6c38617bdd508715ec759043b69df1608d7b0db90c0f8523492602067ffffffffffffffff8d511692604051908152a28a610708565b8380fd5b8280fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b3461013f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f57806004016101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261013f5760443567ffffffffffffffff811161013f57610a84903690600401611c77565b50610a96610a9182611d04565b611e6b565b67ffffffffffffffff610aa882611d04565b1680600052600060205273ffffffffffffffffffffffffffffffffffffffff6040600020541690811561038b576020906044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115610c7b57600091610c4c575b5015610c1e5780610b33610b6092611d86565b6101446040610b4460248601611d04565b67ffffffffffffffff6000911681528060205220930190611d19565b90357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110610be9575b505060601c9060ff815460e01c16610ba357005b60008281526002909101602052604090205415610bbc57005b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168280610b8f565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b610c6e915060203d602011610c74575b610c668183611c36565b810190611e53565b83610b20565b503d610c5c565b6040513d6000823e3d90fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5767ffffffffffffffff610cc7611bee565b600060a0604051610cd781611c1a565b828152826020820152826040820152826060820152826080820152015216806000526000602052604060002080549160405190610d1382611c1a565b73ffffffffffffffffffffffffffffffffffffffff8416825260208201908152604082019360ff8160e01c1615158552606083019361ffff8260a01c1685526001608085019163ffffffff8460b01c16835261ffff60a087019460d01c168452019460405196879460208854998a81520198899860005260206000209060005b818110610e5c575073ffffffffffffffffffffffffffffffffffffffff60e08b8b8b8f8c63ffffffff8d61ffff8e8e67ffffffffffffffff8f610dd98a869a038b611c36565b6040519d8e9d8e019b51168d52511660208c015251151560408b01525116606089015251166080870152511660a085015260e060c0850152518091526101008301919060005b818110610e2d575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610e1f565b82548c526020909b019a60019283019201610d93565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57600154610ead81611cec565b90610ebb6040519283611c36565b80825260208201809160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6916000905b828210610f7957848660405191829160208301906020845251809152604083019060408160051b85010192916000905b828210610f2e57505050500390f35b91936020610f69827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851611b1a565b9601920192018594939192610f1f565b604051600085548060011c90600181168015611075575b602083108114611048578285529081156110075750600114610fcf575b5060019282610fc185946020940382611c36565b815201940191019092610eef565b6000878152602081209092505b818310610ff157505081016020016001610fad565b6001816020925483868801015201920191610fdc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050610fad565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f1691610f90565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760025473ffffffffffffffffffffffffffffffffffffffff8160201c163303611161577fffffffffffffffff0000000000000000000000000000000000000000ffffffff60035491337fffffffffffffffffffffffff00000000000000000000000000000000000000008416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f576111da903690600401611bbd565b73ffffffffffffffffffffffffffffffffffffffff6003541680156113a65760005b8281101561099e5760008160051b8501359073ffffffffffffffffffffffffffffffffffffffff82168092036113a357604051907f70a08231000000000000000000000000000000000000000000000000000000008252306004830152602082602481865afa91821561139657819261135f575b5081611282575b5050506001016111fc565b602081604051828101907fa9059cbb000000000000000000000000000000000000000000000000000000008252886024820152856044820152604481526112ca606482611c36565b519082875af1156113535780513d61134a5750823b155b61131e575090837f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a3908580611277565b80837f5274afe70000000000000000000000000000000000000000000000000000000060249352600452fd5b600114156112e1565b604051903d90823e3d90fd5b9091506020813d821161138e575b8161137a60209383611c36565b8101031261138a57519087611270565b5080fd5b3d915061136d565b50604051903d90823e3d90fd5b80fd5b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f573660238201121561013f57806004013567ffffffffffffffff811161013f57602460c0820283010136811161013f57611448611f61565b61145182611cec565b9161145f6040519384611c36565b825260009260240190602083015b818310611682578480855b805183101561167e5761148b8382611fac565b519267ffffffffffffffff602085015116938415611652578484526020849052604080852082518154928401517fffffff00ffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560e01b7cff000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156116265773ffffffffffffffffffffffffffffffffffffffff7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007bffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260e01c1615156020820152a2019190611478565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867f97ccaab7000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261172b576040519061169982611c1a565b833573ffffffffffffffffffffffffffffffffffffffff811681036117275782526116c660208501611c05565b60208301526116d760408501611d6a565b60408301526116e860608501611d77565b606083015260808401359063ffffffff821682036117275782602092608060c095015261171760a08701611d77565b60a082015281520192019161146d565b8680fd5b8480fd5b3461013f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f57806004016101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261013f576117a9611b79565b5060843567ffffffffffffffff811161013f573660238201121561013f57806004013567ffffffffffffffff811161013f573691016024011161013f57611810816117f5602093611d86565b6101246024850194611809610a9187611d04565b0190611d19565b908092918101031261013f57359073ffffffffffffffffffffffffffffffffffffffff821680920361013f5761184e67ffffffffffffffff91611d04565b1680600052600060205260406000209081549073ffffffffffffffffffffffffffffffffffffffff821690811561038b576020906024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa908115610c7b5760009161197e575b5073ffffffffffffffffffffffffffffffffffffffff163303610c1e5760e01c60ff16611961575b61195d6040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000016602082015260048152611949602482611c36565b604051918291602083526020830190611b1a565b0390f35b60008281526002909101602052604090205415610bbc57806118ec565b6020813d6020116119d9575b8161199760209383611c36565b8101031261138a57519073ffffffffffffffffffffffffffffffffffffffff821682036113a3575073ffffffffffffffffffffffffffffffffffffffff6118c4565b3d915061198a565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761195d6040805190611a228183611c36565b601882527f56657269666965725465737448656c70657220322e302e300000000000000000602083015251918291602083526020830190611b1a565b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361013f57817fd3e969cd0000000000000000000000000000000000000000000000000000000060209314908115611af0575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611ae9565b919082519283825260005b848110611b645750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611b25565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361013f57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361013f57565b9181601f8401121561013f5782359167ffffffffffffffff831161013f576020808501948460051b01011161013f57565b6004359067ffffffffffffffff8216820361013f57565b359067ffffffffffffffff8216820361013f57565b60c0810190811067ffffffffffffffff82111761040457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761040457604052565b81601f8201121561013f5780359067ffffffffffffffff82116104045760405192611cca601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185611c36565b8284526020838301011161013f57816000926020809301838601378301015290565b67ffffffffffffffff81116104045760051b60200190565b3567ffffffffffffffff8116810361013f5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561013f570180359067ffffffffffffffff821161013f5760200191813603831361013f57565b3590811515820361013f57565b359061ffff8216820361013f57565b6101808101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561013f57810180359067ffffffffffffffff821161013f5760208260051b360391011361013f571590811591611e39575b8115611e1d575b50611df357565b7fec14a34d0000000000000000000000000000000000000000000000000000000060005260046000fd5b60809150013563ffffffff811680910361013f57151538611dec565b9050611e496101a0820182611d19565b9050151590611de5565b9081602091031261013f5751801515810361013f5790565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c7b57600091611f42575b50611f0a5750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b611f5b915060203d602011610c7457610c668183611c36565b38611f02565b73ffffffffffffffffffffffffffffffffffffffff600354163303611f8257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015611fc05760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561013f57813561200681611cec565b926120146040519485611c36565b81845260208085019260051b82010192831161013f57602001905b82821061203c5750505090565b6020809161204984611b9c565b81520191019061202f565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115612136576120878161213b565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c16166121365761ffff8360e01c168015918215612125575b50506120d1575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff1610905038806120c7565b505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115612241577dffff000000000000000000000000000000000000000000000000000000008116156122385760ff60015b169060f01c806121d3575b506001036121a65750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b601081106121e4575061219b565b6001811b82166121f7575b6001016121d6565b916001810180911161220957916121ef565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff6000612190565b5050565b8054821015611fc05760005260206000200190600090565b90600182019181600052826020526040600020548015156000146123ec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111612209578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161220957818103612380575b50505080548015612351577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906123128282612245565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6123d56123906123a09386612245565b90549060031b1c92839286612245565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052836020526040600020553880806122da565b50505050600090565b600082815260018201602052604090205461244a578054906801000000000000000082101561040457826124336123a0846001809601855584612245565b905580549260005201602052604060002055600190565b505060009056fea164736f6c634300081a000a",
}

var VerifierTestHelperABI = VerifierTestHelperMetaData.ABI

var VerifierTestHelperBin = VerifierTestHelperMetaData.Bin

func DeployVerifierTestHelper(auth *bind.TransactOpts, backend bind.ContractBackend, storageLocations []string, rmn common.Address, versionTag [4]byte) (common.Address, *types.Transaction, *VerifierTestHelper, error) {
	parsed, err := VerifierTestHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierTestHelperBin), backend, storageLocations, rmn, versionTag)
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
