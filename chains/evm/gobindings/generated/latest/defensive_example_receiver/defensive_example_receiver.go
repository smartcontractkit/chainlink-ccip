// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defensive_example_receiver

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

type CCIPClientExampleRemoteChainConfig struct {
	ExtraArgsBytes    []byte
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var DefensiveExampleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouterClient\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCIPClientExample.RemoteChainConfig\",\"components\":[{\"name\":\"extraArgsBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retryFailedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_feeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_messageContents\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendDataAndTokens\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendDataPayFeeToken\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendDataPayNative\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTokens\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSimRevert\",\"inputs\":[{\"name\":\"simRevert\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageReceived\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageRecovered\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSucceeded\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MessageNotFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461013257604081612f3d803803809161001f8285610172565b8339810103126101325780516001600160a01b038116919082900361013257602001516001600160a01b038116919082900361013257801561015c5780608052331561014b57600180546001600160a01b03199081163317909155600280549091168317905560405163095ea7b360e01b81526004810191909152600019602482015290602090829060449082906000905af1801561013f57610102575b60ff1960085416600855604051612d9190816101ac823960805181818161080c015281816109bb01528181610ba801528181610ee30152818161122001526114b70152f35b6020813d602011610137575b8161011b60209383610172565b8101031261013257518015150361013257386100bd565b600080fd5b3d915061010e565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b0382119082101761019557604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461191257508063369f7f66146116815780634c42071c146115b857806352f813c3146115425780635d8e7f18146113ce578063686fab791461117e5780637383c8fb14610e1d5780637909b54914610d4f57806379ba509714610c6657806385572ffb14610b83578063898068fc14610a315780638da5cb5b146109df578063b0f479a114610970578063c0cb2d081461077a578063cf6730f81461061c578063d0d1fc3d1461057c578063db4eac861461023a578063f2fde38b146101475763ff2deec3146100f057600080fd5b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b600080fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760043573ffffffffffffffffffffffffffffffffffffffff81168091036101425761019f612a8c565b33811461021057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257610271611a20565b60243567ffffffffffffffff811161014257610291903690600401611b06565b60443567ffffffffffffffff8111610142576102b1903690600401611e9a565b60643567ffffffffffffffff8111610142576102d1903690600401611e9a565b906084359360ff85168095036101425767ffffffffffffffff906102f3612a8c565b6040519461030086611a6f565b85526020850192835260408501938452606085019586521660005260036020526040600020925180519067ffffffffffffffff82116104755761034d826103478754611d85565b876122dd565b602090601f83116001146104d95761039b9291600091836104ce575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b83555b51805190600184019067ffffffffffffffff8311610475576020906103c38484611f7d565b0190600052602060002060005b8381106104a457505050506002820190519081519167ffffffffffffffff8311610475576020906104018484611f7d565b0190600052602060002060005b83811061044b578560ff600387019151167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055600080f35b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161040e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016103d0565b015190508780610369565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169186600052816000209260005b818110610564575090846001959493921061052d575b505050811b01835561039e565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055868080610520565b9293602060018192878601518155019501930161050a565b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760043560005260046020526040600020805461061867ffffffffffffffff6001840154169261060a6105eb60036105e460028501611dd8565b9301611dd8565b9160405195869586526020860152608060408601526080850190611d26565b908382036060850152611d26565b0390f35b346101425761062a36611cb7565b3033036107505767ffffffffffffffff610646602083016121b8565b1680600052600360205261065e604060002054611d85565b15610723575060ff600854166106f957608081019060005b6106808383612322565b90508110156106f757806106f173ffffffffffffffffffffffffffffffffffffffff6106c06106bb6001956106b58989612322565b90612a7c565b612376565b1673ffffffffffffffffffffffffffffffffffffffff84541660206106e9856106b58a8a612322565b013591612ad7565b01610676565b005b7f79f79e0b0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f14d4a4e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425761078836611b4d565b67ffffffffffffffff83168060005260036020526107aa604060002054611d85565b1561072357906107f0916107bc611fb6565b906000526003602052604060002091604051946107d886611a37565b85526020850152604084015260006060840152611dd8565b608082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166040517f20487ded00000000000000000000000000000000000000000000000000000000815260208180610864868860048401611ffe565b0381855afa90811561092a57600091610936575b50906108b793602093926040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401611ffe565b03925af190811561092a576000916108f7575b7f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b90506020813d602011610922575b8161091260209383611a8b565b81010312610142575160206108ca565b3d9150610905565b6040513d6000823e3d90fd5b929190506020833d602011610968575b8161095360209383611a8b565b810103126101425791519091906108b7610878565b3d9150610946565b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff610a71611a20565b60006060604051610a8181611a6f565b8181528160208201528160408201520152166000526003602052610b15604060002060405190610ab082611a6f565b610ab981611dd8565b825260ff610b77610acc60018401612157565b9260208501938452610b46836003610ae660028501612157565b936040890194855201541694606087019586526040519788976020895251608060208a015260a0890190611d26565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0888303016040890152611c6d565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303016060870152611c6d565b91511660808301520390f35b3461014257610b9136611cb7565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163303610c385767ffffffffffffffff610be3602083016121b8565b1690816000526003602052610bfc604060002054611d85565b15610c0a576106f790612397565b507fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760005473ffffffffffffffffffffffffffffffffffffffff81163303610d25577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff610d8f611a20565b166000526003602052610e0560406000206060610e13604051610db181611a6f565b610dba84611dd8565b8152610dc860018501612157565b906020810191825260ff6003610de060028801612157565b9687604085015201541693849101525192604051948594606086526060860190611c6d565b908482036020860152611c6d565b9060408301520390f35b346101425760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257610e54611a20565b60243567ffffffffffffffff811161014257610e74903690600401611b06565b9060443567ffffffffffffffff811161014257610e95903690600401611b06565b9060643567ffffffffffffffff811161014257610eb6903690600401611be6565b9067ffffffffffffffff811690816000526003602052610eda604060002054611d85565b15610c0a5790927f0000000000000000000000000000000000000000000000000000000000000000919060005b845181101561109757806020610f9173ffffffffffffffffffffffffffffffffffffffff610f376000958a611eff565b51511682610f45858b611eff565b5101516040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915294859283919082906064820190565b03925af191821561092a5760009261107b575b50602061103373ffffffffffffffffffffffffffffffffffffffff610fc9848a611eff565b51511682610fd7858b611eff565b5101516040519586809481937f095ea7b30000000000000000000000000000000000000000000000000000000083528c600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af191821561092a5760019261104d575b5001610f07565b61106d9060203d8111611074575b6110658183611a8b565b81019061213f565b5088611046565b503d61105b565b6110929060203d8111611074576110658183611a8b565b610fa4565b5061113e95936110ed916020966000526003875260406000209173ffffffffffffffffffffffffffffffffffffffff6002541691604051976110d889611a37565b88528888015260408701526060860152611dd8565b6080840152600073ffffffffffffffffffffffffffffffffffffffff6040518097819682957f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401611ffe565b0393165af190811561092a576000916108f7577f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b346101425760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610142576111b5611a20565b60243567ffffffffffffffff8111610142576111d5903690600401611b06565b9060443567ffffffffffffffff8111610142576111f6903690600401611be6565b67ffffffffffffffff8216806000526003602052611218604060002054611d85565b1561072357917f00000000000000000000000000000000000000000000000000000000000000009060005b83518110156113775780602061128173ffffffffffffffffffffffffffffffffffffffff61127360009589611eff565b51511682610f45858a611eff565b03925af191821561092a5760009261135b575b50602061132373ffffffffffffffffffffffffffffffffffffffff6112b98489611eff565b515116826112c7858a611eff565b5101516040519586809481937f095ea7b30000000000000000000000000000000000000000000000000000000083528b600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af191821561092a5760019261133d575b5001611243565b6113549060203d8111611074576110658183611a8b565b5087611336565b6113729060203d8111611074576110658183611a8b565b611294565b506110ed61113e95936020956000526003865260406000209073ffffffffffffffffffffffffffffffffffffffff6002541690604051966113b788611a37565b875260608888015260408701526060860152611dd8565b34610142576113dc36611b4d565b9167ffffffffffffffff811692836000526003602052611400604060002054611d85565b15611514579061146361149c9460209493611419611fb6565b916000526003865260406000209173ffffffffffffffffffffffffffffffffffffffff60025416916040519661144e88611a37565b87528787015260408601526060850152611dd8565b608083015260405193849283927f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401611ffe565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561092a576000916108f7577f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b837fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760043580151580910361014257611586612a8c565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060085416911617600855600080f35b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff6115f8611a20565b611600612a8c565b166000526003602052600060036040822061161b8154611d85565b8061163e575b5061162e60018201611f59565b61163a60028201611f59565b0155005b601f811160011461165457508281555b83611621565b8184526020842061167091601f0160051c810190600101611f42565b80835282602081208183555561164e565b346101425760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760243560043573ffffffffffffffffffffffffffffffffffffffff82168203610142576116db612a8c565b80600052600760205260406000205480158015906118fb575b1561189d57600103611870578060005260076020526000604081205561171981612c37565b5080600052600460205260406000209160046040519361173885611a37565b8054855267ffffffffffffffff600182015416602086015261175c60028201611dd8565b604086015261176d60038201611dd8565b60608601520192835461177f81611bce565b9461178d6040519687611a8b565b818652602086019060005260206000206000915b83831061182d575050505060800192835260005b8351805182101561180557906117ff73ffffffffffffffffffffffffffffffffffffffff6117e583600195611eff565b5151168460206117f6858a51611eff565b51015191612ad7565b016117b5565b837fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad600080a2005b6002602060019260405161184081611a53565b73ffffffffffffffffffffffffffffffffffffffff865416815284860154838201528152019201920191906117a1565b7fb6e782600000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152fd5b5081600052600660205260406000205415156116f4565b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361014257817f85572ffb00000000000000000000000000000000000000000000000000000000602093149081156119d5575b81156119ab575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836119a4565b7f7909b549000000000000000000000000000000000000000000000000000000008114915061199d565b359073ffffffffffffffffffffffffffffffffffffffff8216820361014257565b6004359067ffffffffffffffff8216820361014257565b60a0810190811067ffffffffffffffff82111761047557604052565b6040810190811067ffffffffffffffff82111761047557604052565b6080810190811067ffffffffffffffff82111761047557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761047557604052565b67ffffffffffffffff811161047557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561014257803590611b1d82611acc565b92611b2b6040519485611a8b565b8284526020838301011161014257816000926020809301838601378301015290565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101425760043567ffffffffffffffff81168103610142579160243567ffffffffffffffff81116101425782611bab91600401611b06565b916044359067ffffffffffffffff821161014257611bcb91600401611b06565b90565b67ffffffffffffffff81116104755760051b60200190565b81601f8201121561014257803590611bfd82611bce565b92611c0b6040519485611a8b565b82845260208085019360061b8301019181831161014257602001925b828410611c35575050505090565b6040848303126101425760206040918251611c4f81611a53565b611c58876119ff565b81528287013583820152815201930192611c27565b906020808351928381520192019060005b818110611c8b5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611c7e565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610142576004359067ffffffffffffffff8211610142577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126101425760040190565b919082519283825260005b848110611d705750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611d31565b90600182811c92168015611dce575b6020831014611d9f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611d94565b9060405191826000825492611dec84611d85565b8084529360018116908115611e5a5750600114611e13575b50611e1192500383611a8b565b565b90506000929192526020600020906000915b818310611e3e575050906020611e119282010138611e04565b6020919350806001915483858901015201910190918492611e25565b60209350611e119592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e04565b9080601f83011215610142578135611eb181611bce565b92611ebf6040519485611a8b565b81845260208085019260051b82010192831161014257602001905b828210611ee75750505090565b60208091611ef4846119ff565b815201910190611eda565b8051821015611f135760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b818110611f4d575050565b60008155600101611f42565b80546000825580611f68575050565b611e1191600052602060002090810190611f42565b9068010000000000000000811161047557815490808355818110611fa057505050565b611e119260005260206000209182019101611f42565b60405190611fc5602083611a8b565b600080835282815b828110611fd957505050565b602090604051611fe881611a53565b6000815260008382015282828501015201611fcd565b9067ffffffffffffffff909392931681526040602082015261206361202f845160a0604085015260e0840190611d26565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152611d26565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b8181106121075750505060808473ffffffffffffffffffffffffffffffffffffffff6060611bcb969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152611d26565b8251805173ffffffffffffffffffffffffffffffffffffffff16865260209081015181870152604090950194909201916001016120a4565b90816020910312610142575180151581036101425790565b906040519182815491828252602082019060005260206000209260005b818110612189575050611e1192500383611a8b565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612174565b3567ffffffffffffffff811681036101425790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561014257016020813591019167ffffffffffffffff821161014257813603831361014257565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b3d15612287573d9061226d82611acc565b9161227b6040519384611a8b565b82523d6000602084013e565b606090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610142570180359067ffffffffffffffff82116101425760200191813603831361014257565b9190601f81116122ec57505050565b611e11926000526020600020906020601f840160051c83019310612318575b601f0160051c0190611f42565b909150819061230b565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610142570180359067ffffffffffffffff821161014257602001918160061b3603831361014257565b3573ffffffffffffffffffffffffffffffffffffffff811681036101425790565b303b15610142576040517fcf6730f8000000000000000000000000000000000000000000000000000000008152600091602060048301528035928360248401526020820192833567ffffffffffffffff81168103612a785767ffffffffffffffff166044820152604083019361242161241086866121cd565b60a0606486015260c485019161221d565b612463606086019161243383886121cd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc87840301608488015261221d565b92608086019384357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301811215612a745787016020813591019167ffffffffffffffff8211612a70578160061b36038313612a70578381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160a48501528181528791849160200190835b818110612a2b57505081929350038183305af19081612a17575b506129eb5761251961225c565b9587855260076020526001604086205561253288612c37565b508785526004602052604085209288845567ffffffffffffffff61255960018601926121b8565b167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082541617905561258f60028401918761228c565b9067ffffffffffffffff82116129be576125b3826125ad8554611d85565b856122dd565b8690601f831160011461291e576125fe929188918361283e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b61260f60038301918661228c565b9067ffffffffffffffff82116128f15761262d826125ad8554611d85565b8590601f8311600114612849578261268b96959360049593612681938a9261283e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b0193612322565b929068010000000000000000841161281157815484835580851061275f575b5090825260208220905b8383106126fd57505050506126f87f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f91604051918291602083526020830190611d26565b0390a2565b600260408273ffffffffffffffffffffffffffffffffffffffff612722600195612376565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000865416178555602081013584860155019201920191906126b4565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681036127e4577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff851685036127e457828452602084209060011b8101908560011b015b8181106127d257506126aa565b808560029255856001820155016127c5565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b013590503880610369565b83875260208720917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416885b8181106128d9575092600192859261268b99989660049896106128a1575b505050811b019055612684565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612894565b91936020600181928787013581550195019201612876565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b83885260208820917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416895b8181106129a6575090846001959493921061296e575b505050811b019055612601565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612961565b9193602060018192878701358155019501920161294b565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50505090507fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f915080a2565b85612a2491969296611a8b565b933861250c565b9250925060408060019273ffffffffffffffffffffffffffffffffffffffff612a53886119ff565b1681526020870135602082015201940191019288928592946124f2565b8780fd5b8680fd5b8280fd5b9190811015611f135760061b0190565b73ffffffffffffffffffffffffffffffffffffffff600154163303612aad57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff9384166024830152604480830195909552938152612b8e929091612b3c606484611a8b565b16600080604095865194612b508887611a8b565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af1612b8861225c565b91612cb4565b80519081612b9b57505050565b602080612bac93830101910161213f565b15612bb45750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80600052600660205260406000205415600014612cae5760055468010000000000000000811015610475576001810180600555811015611f13577f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00181905560055460009182526006602052604090912055600190565b50600090565b91929015612d2f5750815115612cc8575090565b3b15612cd15790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015612d425750805190602001fd5b612d80906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611d26565b0390fdfea164736f6c634300081a000a",
}

var DefensiveExampleABI = DefensiveExampleMetaData.ABI

var DefensiveExampleBin = DefensiveExampleMetaData.Bin

func DeployDefensiveExample(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address) (common.Address, *types.Transaction, *DefensiveExample, error) {
	parsed, err := DefensiveExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefensiveExampleBin), backend, router, feeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefensiveExample{address: address, abi: *parsed, DefensiveExampleCaller: DefensiveExampleCaller{contract: contract}, DefensiveExampleTransactor: DefensiveExampleTransactor{contract: contract}, DefensiveExampleFilterer: DefensiveExampleFilterer{contract: contract}}, nil
}

type DefensiveExample struct {
	address common.Address
	abi     abi.ABI
	DefensiveExampleCaller
	DefensiveExampleTransactor
	DefensiveExampleFilterer
}

type DefensiveExampleCaller struct {
	contract *bind.BoundContract
}

type DefensiveExampleTransactor struct {
	contract *bind.BoundContract
}

type DefensiveExampleFilterer struct {
	contract *bind.BoundContract
}

type DefensiveExampleSession struct {
	Contract     *DefensiveExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DefensiveExampleCallerSession struct {
	Contract *DefensiveExampleCaller
	CallOpts bind.CallOpts
}

type DefensiveExampleTransactorSession struct {
	Contract     *DefensiveExampleTransactor
	TransactOpts bind.TransactOpts
}

type DefensiveExampleRaw struct {
	Contract *DefensiveExample
}

type DefensiveExampleCallerRaw struct {
	Contract *DefensiveExampleCaller
}

type DefensiveExampleTransactorRaw struct {
	Contract *DefensiveExampleTransactor
}

func NewDefensiveExample(address common.Address, backend bind.ContractBackend) (*DefensiveExample, error) {
	abi, err := abi.JSON(strings.NewReader(DefensiveExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindDefensiveExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefensiveExample{address: address, abi: abi, DefensiveExampleCaller: DefensiveExampleCaller{contract: contract}, DefensiveExampleTransactor: DefensiveExampleTransactor{contract: contract}, DefensiveExampleFilterer: DefensiveExampleFilterer{contract: contract}}, nil
}

func NewDefensiveExampleCaller(address common.Address, caller bind.ContractCaller) (*DefensiveExampleCaller, error) {
	contract, err := bindDefensiveExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleCaller{contract: contract}, nil
}

func NewDefensiveExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*DefensiveExampleTransactor, error) {
	contract, err := bindDefensiveExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleTransactor{contract: contract}, nil
}

func NewDefensiveExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*DefensiveExampleFilterer, error) {
	contract, err := bindDefensiveExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleFilterer{contract: contract}, nil
}

func bindDefensiveExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefensiveExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_DefensiveExample *DefensiveExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefensiveExample.Contract.DefensiveExampleCaller.contract.Call(opts, result, method, params...)
}

func (_DefensiveExample *DefensiveExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DefensiveExampleTransactor.contract.Transfer(opts)
}

func (_DefensiveExample *DefensiveExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DefensiveExampleTransactor.contract.Transact(opts, method, params...)
}

func (_DefensiveExample *DefensiveExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefensiveExample.Contract.contract.Call(opts, result, method, params...)
}

func (_DefensiveExample *DefensiveExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.Contract.contract.Transfer(opts)
}

func (_DefensiveExample *DefensiveExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefensiveExample.Contract.contract.Transact(opts, method, params...)
}

func (_DefensiveExample *DefensiveExampleCaller) GetCCVs(opts *bind.CallOpts, sourceChainSelector uint64) (GetCCVs,

	error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getCCVs", sourceChainSelector)

	outstruct := new(GetCCVs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_DefensiveExample *DefensiveExampleSession) GetCCVs(sourceChainSelector uint64) (GetCCVs,

	error) {
	return _DefensiveExample.Contract.GetCCVs(&_DefensiveExample.CallOpts, sourceChainSelector)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetCCVs(sourceChainSelector uint64) (GetCCVs,

	error) {
	return _DefensiveExample.Contract.GetCCVs(&_DefensiveExample.CallOpts, sourceChainSelector)
}

func (_DefensiveExample *DefensiveExampleCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	if err != nil {
		return *new(CCIPClientExampleRemoteChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPClientExampleRemoteChainConfig)).(*CCIPClientExampleRemoteChainConfig)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) GetRemoteChainConfig(remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	return _DefensiveExample.Contract.GetRemoteChainConfig(&_DefensiveExample.CallOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	return _DefensiveExample.Contract.GetRemoteChainConfig(&_DefensiveExample.CallOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) GetRouter() (common.Address, error) {
	return _DefensiveExample.Contract.GetRouter(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetRouter() (common.Address, error) {
	return _DefensiveExample.Contract.GetRouter(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) Owner() (common.Address, error) {
	return _DefensiveExample.Contract.Owner(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) Owner() (common.Address, error) {
	return _DefensiveExample.Contract.Owner(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) SFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SFeeToken() (common.Address, error) {
	return _DefensiveExample.Contract.SFeeToken(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SFeeToken() (common.Address, error) {
	return _DefensiveExample.Contract.SFeeToken(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) SMessageContents(opts *bind.CallOpts, messageId [32]byte) (SMessageContents,

	error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_messageContents", messageId)

	outstruct := new(SMessageContents)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SourceChainSelector = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_DefensiveExample *DefensiveExampleSession) SMessageContents(messageId [32]byte) (SMessageContents,

	error) {
	return _DefensiveExample.Contract.SMessageContents(&_DefensiveExample.CallOpts, messageId)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SMessageContents(messageId [32]byte) (SMessageContents,

	error) {
	return _DefensiveExample.Contract.SMessageContents(&_DefensiveExample.CallOpts, messageId)
}

func (_DefensiveExample *DefensiveExampleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DefensiveExample.Contract.SupportsInterface(&_DefensiveExample.CallOpts, interfaceId)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DefensiveExample.Contract.SupportsInterface(&_DefensiveExample.CallOpts, interfaceId)
}

func (_DefensiveExample *DefensiveExampleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "acceptOwnership")
}

func (_DefensiveExample *DefensiveExampleSession) AcceptOwnership() (*types.Transaction, error) {
	return _DefensiveExample.Contract.AcceptOwnership(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DefensiveExample.Contract.AcceptOwnership(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "ccipReceive", message)
}

func (_DefensiveExample *DefensiveExampleSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.CcipReceive(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.CcipReceive(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactor) DisableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "disableRemoteChain", remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleSession) DisableRemoteChain(remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) DisableRemoteChain(remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactor) EnableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "enableRemoteChain", remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleSession) EnableRemoteChain(remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) EnableRemoteChain(remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleTransactor) ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "processMessage", message)
}

func (_DefensiveExample *DefensiveExampleSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProcessMessage(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProcessMessage(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactor) RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "retryFailedMessage", messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleSession) RetryFailedMessage(messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.RetryFailedMessage(&_DefensiveExample.TransactOpts, messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) RetryFailedMessage(messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.RetryFailedMessage(&_DefensiveExample.TransactOpts, messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendDataAndTokens(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendDataAndTokens", destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleSession) SendDataAndTokens(destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataAndTokens(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendDataAndTokens(destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataAndTokens(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendDataPayFeeToken(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendDataPayFeeToken", destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleSession) SendDataPayFeeToken(destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataPayFeeToken(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendDataPayFeeToken(destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataPayFeeToken(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendDataPayNative(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendDataPayNative", destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleSession) SendDataPayNative(destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataPayNative(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendDataPayNative(destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataPayNative(&_DefensiveExample.TransactOpts, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendTokens(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendTokens", destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleSession) SendTokens(destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendTokens(&_DefensiveExample.TransactOpts, destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendTokens(destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendTokens(&_DefensiveExample.TransactOpts, destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactor) SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "setSimRevert", simRevert)
}

func (_DefensiveExample *DefensiveExampleSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SetSimRevert(&_DefensiveExample.TransactOpts, simRevert)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SetSimRevert(&_DefensiveExample.TransactOpts, simRevert)
}

func (_DefensiveExample *DefensiveExampleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "transferOwnership", to)
}

func (_DefensiveExample *DefensiveExampleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.TransferOwnership(&_DefensiveExample.TransactOpts, to)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.TransferOwnership(&_DefensiveExample.TransactOpts, to)
}

type DefensiveExampleMessageFailedIterator struct {
	Event *DefensiveExampleMessageFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageFailed)
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
		it.Event = new(DefensiveExampleMessageFailed)
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

func (it *DefensiveExampleMessageFailedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageFailedIterator{contract: _DefensiveExample.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageFailed)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageFailed(log types.Log) (*DefensiveExampleMessageFailed, error) {
	event := new(DefensiveExampleMessageFailed)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageReceivedIterator struct {
	Event *DefensiveExampleMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageReceived)
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
		it.Event = new(DefensiveExampleMessageReceived)
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

func (it *DefensiveExampleMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageReceived struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*DefensiveExampleMessageReceivedIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageReceivedIterator{contract: _DefensiveExample.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageReceived) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageReceived)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageReceived(log types.Log) (*DefensiveExampleMessageReceived, error) {
	event := new(DefensiveExampleMessageReceived)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageRecoveredIterator struct {
	Event *DefensiveExampleMessageRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageRecovered)
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
		it.Event = new(DefensiveExampleMessageRecovered)
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

func (it *DefensiveExampleMessageRecoveredIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageRecovered struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageRecoveredIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageRecoveredIterator{contract: _DefensiveExample.contract, event: "MessageRecovered", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageRecovered, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageRecovered)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageRecovered(log types.Log) (*DefensiveExampleMessageRecovered, error) {
	event := new(DefensiveExampleMessageRecovered)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageSentIterator struct {
	Event *DefensiveExampleMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageSent)
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
		it.Event = new(DefensiveExampleMessageSent)
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

func (it *DefensiveExampleMessageSentIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageSent struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageSent(opts *bind.FilterOpts) (*DefensiveExampleMessageSentIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageSentIterator{contract: _DefensiveExample.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSent) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageSent)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageSent(log types.Log) (*DefensiveExampleMessageSent, error) {
	event := new(DefensiveExampleMessageSent)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageSucceededIterator struct {
	Event *DefensiveExampleMessageSucceeded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageSucceededIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageSucceeded)
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
		it.Event = new(DefensiveExampleMessageSucceeded)
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

func (it *DefensiveExampleMessageSucceededIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageSucceeded struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageSucceededIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageSucceededIterator{contract: _DefensiveExample.contract, event: "MessageSucceeded", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSucceeded, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageSucceeded)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageSucceeded(log types.Log) (*DefensiveExampleMessageSucceeded, error) {
	event := new(DefensiveExampleMessageSucceeded)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleOwnershipTransferRequestedIterator struct {
	Event *DefensiveExampleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleOwnershipTransferRequested)
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
		it.Event = new(DefensiveExampleOwnershipTransferRequested)
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

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleOwnershipTransferRequestedIterator{contract: _DefensiveExample.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleOwnershipTransferRequested)
				if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseOwnershipTransferRequested(log types.Log) (*DefensiveExampleOwnershipTransferRequested, error) {
	event := new(DefensiveExampleOwnershipTransferRequested)
	if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleOwnershipTransferredIterator struct {
	Event *DefensiveExampleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleOwnershipTransferred)
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
		it.Event = new(DefensiveExampleOwnershipTransferred)
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

func (it *DefensiveExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleOwnershipTransferredIterator{contract: _DefensiveExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleOwnershipTransferred)
				if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseOwnershipTransferred(log types.Log) (*DefensiveExampleOwnershipTransferred, error) {
	event := new(DefensiveExampleOwnershipTransferred)
	if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCCVs struct {
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}
type SMessageContents struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
}

func (DefensiveExampleMessageFailed) Topic() common.Hash {
	return common.HexToHash("0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f")
}

func (DefensiveExampleMessageReceived) Topic() common.Hash {
	return common.HexToHash("0xe29dc34207c78fc0f6048a32f159139c33339c6d6df8b07dcd33f6d699ff2327")
}

func (DefensiveExampleMessageRecovered) Topic() common.Hash {
	return common.HexToHash("0xef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad")
}

func (DefensiveExampleMessageSent) Topic() common.Hash {
	return common.HexToHash("0x54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a")
}

func (DefensiveExampleMessageSucceeded) Topic() common.Hash {
	return common.HexToHash("0xdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f")
}

func (DefensiveExampleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (DefensiveExampleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_DefensiveExample *DefensiveExample) Address() common.Address {
	return _DefensiveExample.address
}

type DefensiveExampleInterface interface {
	GetCCVs(opts *bind.CallOpts, sourceChainSelector uint64) (GetCCVs,

		error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SFeeToken(opts *bind.CallOpts) (common.Address, error)

	SMessageContents(opts *bind.CallOpts, messageId [32]byte) (SMessageContents,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	EnableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error)

	ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error)

	SendDataAndTokens(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error)

	SendDataPayFeeToken(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error)

	SendDataPayNative(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error)

	SendTokens(opts *bind.TransactOpts, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error)

	SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageFailedIterator, error)

	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageFailed, messageId [][32]byte) (event.Subscription, error)

	ParseMessageFailed(log types.Log) (*DefensiveExampleMessageFailed, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*DefensiveExampleMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*DefensiveExampleMessageReceived, error)

	FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageRecoveredIterator, error)

	WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageRecovered, messageId [][32]byte) (event.Subscription, error)

	ParseMessageRecovered(log types.Log) (*DefensiveExampleMessageRecovered, error)

	FilterMessageSent(opts *bind.FilterOpts) (*DefensiveExampleMessageSentIterator, error)

	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSent) (event.Subscription, error)

	ParseMessageSent(log types.Log) (*DefensiveExampleMessageSent, error)

	FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageSucceededIterator, error)

	WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSucceeded, messageId [][32]byte) (event.Subscription, error)

	ParseMessageSucceeded(log types.Log) (*DefensiveExampleMessageSucceeded, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*DefensiveExampleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*DefensiveExampleOwnershipTransferred, error)

	Address() common.Address
}
