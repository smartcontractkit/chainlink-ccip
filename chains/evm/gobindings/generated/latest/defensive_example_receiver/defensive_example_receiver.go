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

type CCIPClientExampleChainConfig struct {
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouterClient\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableChain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableChain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainConfig\",\"inputs\":[{\"name\":\"selector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCIPClientExample.ChainConfig\",\"components\":[{\"name\":\"extraArgsBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retryFailedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_feeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_messageContents\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendDataAndTokens\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendDataPayFeeToken\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendDataPayNative\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTokens\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSimRevert\",\"inputs\":[{\"name\":\"simRevert\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageReceived\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageRecovered\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSucceeded\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MessageNotFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461013257604081612f3d803803809161001f8285610172565b8339810103126101325780516001600160a01b038116919082900361013257602001516001600160a01b038116919082900361013257801561015c5780608052331561014b57600180546001600160a01b03199081163317909155600280549091168317905560405163095ea7b360e01b81526004810191909152600019602482015290602090829060449082906000905af1801561013f57610102575b60ff1960085416600855604051612d9190816101ac82396080518181816104ca015281816107cb0152818161086601528181610ba101528181610ede01526111750152f35b6020813d602011610137575b8161011b60209383610172565b8101031261013257518015150361013257386100bd565b600080fd5b3d915061010e565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b0382119082101761019557604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461191257508063369f7f661461168157806341eade46146115b85780634fa1d6161461127657806352f813c3146112005780635d8e7f181461108c578063686fab7914610e3c5780637383c8fb14610adb5780637909b54914610a0d57806379ba50971461092457806385572ffb146108415780638da5cb5b146107ef578063b0f479a114610780578063b149092b1461062e578063c0cb2d0814610438578063cf6730f8146102da578063d0d1fc3d1461023a578063f2fde38b146101475763ff2deec3146100f057600080fd5b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b600080fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760043573ffffffffffffffffffffffffffffffffffffffff81168091036101425761019f612a8c565b33811461021057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610142576004356000526004602052604060002080546102d667ffffffffffffffff600184015416926102c86102a960036102a260028501611e3d565b9301611e3d565b9160405195869586526020860152608060408601526080850190611d8b565b908382036060850152611d8b565b0390f35b34610142576102e836611d1c565b30330361040e5767ffffffffffffffff610304602083016121fd565b1680600052600360205261031c604060002054611dea565b156103e1575060ff600854166103b757608081019060005b61033e8383612322565b90508110156103b557806103af73ffffffffffffffffffffffffffffffffffffffff61037e6103796001956103738989612322565b90612a7c565b612376565b1673ffffffffffffffffffffffffffffffffffffffff84541660206103a7856103738a8a612322565b013591612ad7565b01610334565b005b7f79f79e0b0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fd79f2ea40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f14d4a4e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425761044636611bca565b67ffffffffffffffff8316806000526003602052610468604060002054611dea565b156103e157906104ae9161047a611ffb565b9060005260036020526040600020916040519461049686611a37565b85526020850152604084015260006060840152611e3d565b608082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166040517f20487ded00000000000000000000000000000000000000000000000000000000815260208180610522868860048401612043565b0381855afa9081156105e8576000916105f4575b509061057593602093926040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401612043565b03925af19081156105e8576000916105b5575b7f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b90506020813d6020116105e0575b816105d060209383611a8b565b8101031261014257516020610588565b3d91506105c3565b6040513d6000823e3d90fd5b929190506020833d602011610626575b8161061160209383611a8b565b81010312610142579151909190610575610536565b3d9150610604565b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff61066e611a20565b6000606060405161067e81611a6f565b81815281602082015281604082015201521660005260036020526107126040600020604051906106ad82611a6f565b6106b681611e3d565b825260ff6107746106c96001840161219c565b92602085019384526107438360036106e36002850161219c565b936040890194855201541694606087019586526040519788976020895251608060208a015260a0890190611d8b565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0888303016040890152611cd2565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303016060870152611cd2565b91511660808301520390f35b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101425761084f36611d1c565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036108f65767ffffffffffffffff6108a1602083016121fd565b16908160005260036020526108ba604060002054611dea565b156108c8576103b590612397565b507fd79f2ea40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101425760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760005473ffffffffffffffffffffffffffffffffffffffff811633036109e3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff610a4d611a20565b166000526003602052610ac360406000206060610ad1604051610a6f81611a6f565b610a7884611e3d565b8152610a866001850161219c565b906020810191825260ff6003610a9e6002880161219c565b9687604085015201541693849101525192604051948594606086526060860190611cd2565b908482036020860152611cd2565b9060408301520390f35b346101425760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257610b12611a20565b60243567ffffffffffffffff811161014257610b32903690600401611b06565b9060443567ffffffffffffffff811161014257610b53903690600401611b06565b9060643567ffffffffffffffff811161014257610b74903690600401611c4b565b9067ffffffffffffffff811690816000526003602052610b98604060002054611dea565b156108c85790927f0000000000000000000000000000000000000000000000000000000000000000919060005b8451811015610d5557806020610c4f73ffffffffffffffffffffffffffffffffffffffff610bf56000958a611eff565b51511682610c03858b611eff565b5101516040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915294859283919082906064820190565b03925af19182156105e857600092610d39575b506020610cf173ffffffffffffffffffffffffffffffffffffffff610c87848a611eff565b51511682610c95858b611eff565b5101516040519586809481937f095ea7b30000000000000000000000000000000000000000000000000000000083528c600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af19182156105e857600192610d0b575b5001610bc5565b610d2b9060203d8111610d32575b610d238183611a8b565b810190612184565b5088610d04565b503d610d19565b610d509060203d8111610d3257610d238183611a8b565b610c62565b50610dfc9593610dab916020966000526003875260406000209173ffffffffffffffffffffffffffffffffffffffff600254169160405197610d9689611a37565b88528888015260408701526060860152611e3d565b6080840152600073ffffffffffffffffffffffffffffffffffffffff6040518097819682957f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401612043565b0393165af19081156105e8576000916105b5577f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b346101425760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257610e73611a20565b60243567ffffffffffffffff811161014257610e93903690600401611b06565b9060443567ffffffffffffffff811161014257610eb4903690600401611c4b565b67ffffffffffffffff8216806000526003602052610ed6604060002054611dea565b156103e157917f00000000000000000000000000000000000000000000000000000000000000009060005b835181101561103557806020610f3f73ffffffffffffffffffffffffffffffffffffffff610f3160009589611eff565b51511682610c03858a611eff565b03925af19182156105e857600092611019575b506020610fe173ffffffffffffffffffffffffffffffffffffffff610f778489611eff565b51511682610f85858a611eff565b5101516040519586809481937f095ea7b30000000000000000000000000000000000000000000000000000000083528b600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af19182156105e857600192610ffb575b5001610f01565b6110129060203d8111610d3257610d238183611a8b565b5087610ff4565b6110309060203d8111610d3257610d238183611a8b565b610f52565b50610dab610dfc95936020956000526003865260406000209073ffffffffffffffffffffffffffffffffffffffff60025416906040519661107588611a37565b875260608888015260408701526060860152611e3d565b346101425761109a36611bca565b9167ffffffffffffffff8116928360005260036020526110be604060002054611dea565b156111d2579061112161115a94602094936110d7611ffb565b916000526003865260406000209173ffffffffffffffffffffffffffffffffffffffff60025416916040519661110c88611a37565b87528787015260408601526060850152611e3d565b608083015260405193849283927f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401612043565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156105e8576000916105b5577f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602083604051908152a1005b837fd79f2ea40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760043580151580910361014257611244612a8c565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060085416911617600855600080f35b346101425760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610142576112ad611a20565b60243567ffffffffffffffff8111610142576112cd903690600401611b06565b60443567ffffffffffffffff8111610142576112ed903690600401611b65565b60643567ffffffffffffffff81116101425761130d903690600401611b65565b906084359360ff85168095036101425767ffffffffffffffff9061132f612a8c565b6040519461133c86611a6f565b85526020850192835260408501938452606085019586521660005260036020526040600020925180519067ffffffffffffffff82116114b157611389826113838754611dea565b87611fb6565b602090601f8311600114611515576113d792916000918361150a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b83555b51805190600184019067ffffffffffffffff83116114b1576020906113ff8484611f7d565b0190600052602060002060005b8381106114e057505050506002820190519081519167ffffffffffffffff83116114b15760209061143d8484611f7d565b0190600052602060002060005b838110611487578560ff600387019151167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055600080f35b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161144a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161140c565b0151905087806113a5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169186600052816000209260005b8181106115a05750908460019594939210611569575b505050811b0183556113da565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905586808061155c565b92936020600181928786015181550195019301611546565b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425767ffffffffffffffff6115f8611a20565b611600612a8c565b166000526003602052600060036040822061161b8154611dea565b8061163e575b5061162e60018201611f59565b61163a60028201611f59565b0155005b601f811160011461165457508281555b83611621565b8184526020842061167091601f0160051c810190600101611f42565b80835282602081208183555561164e565b346101425760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101425760243560043573ffffffffffffffffffffffffffffffffffffffff82168203610142576116db612a8c565b80600052600760205260406000205480158015906118fb575b1561189d57600103611870578060005260076020526000604081205561171981612c37565b5080600052600460205260406000209160046040519361173885611a37565b8054855267ffffffffffffffff600182015416602086015261175c60028201611e3d565b604086015261176d60038201611e3d565b60608601520192835461177f81611b4d565b9461178d6040519687611a8b565b818652602086019060005260206000206000915b83831061182d575050505060800192835260005b8351805182101561180557906117ff73ffffffffffffffffffffffffffffffffffffffff6117e583600195611eff565b5151168460206117f6858a51611eff565b51015191612ad7565b016117b5565b837fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad600080a2005b6002602060019260405161184081611a53565b73ffffffffffffffffffffffffffffffffffffffff865416815284860154838201528152019201920191906117a1565b7fb6e782600000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152fd5b5081600052600660205260406000205415156116f4565b346101425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014257600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361014257817f85572ffb00000000000000000000000000000000000000000000000000000000602093149081156119d5575b81156119ab575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836119a4565b7f7909b549000000000000000000000000000000000000000000000000000000008114915061199d565b359073ffffffffffffffffffffffffffffffffffffffff8216820361014257565b6004359067ffffffffffffffff8216820361014257565b60a0810190811067ffffffffffffffff8211176114b157604052565b6040810190811067ffffffffffffffff8211176114b157604052565b6080810190811067ffffffffffffffff8211176114b157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176114b157604052565b67ffffffffffffffff81116114b157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561014257803590611b1d82611acc565b92611b2b6040519485611a8b565b8284526020838301011161014257816000926020809301838601378301015290565b67ffffffffffffffff81116114b15760051b60200190565b9080601f83011215610142578135611b7c81611b4d565b92611b8a6040519485611a8b565b81845260208085019260051b82010192831161014257602001905b828210611bb25750505090565b60208091611bbf846119ff565b815201910190611ba5565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101425760043567ffffffffffffffff81168103610142579160243567ffffffffffffffff81116101425782611c2891600401611b06565b916044359067ffffffffffffffff821161014257611c4891600401611b06565b90565b81601f8201121561014257803590611c6282611b4d565b92611c706040519485611a8b565b82845260208085019360061b8301019181831161014257602001925b828410611c9a575050505090565b6040848303126101425760206040918251611cb481611a53565b611cbd876119ff565b81528287013583820152815201930192611c8c565b906020808351928381520192019060005b818110611cf05750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611ce3565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610142576004359067ffffffffffffffff8211610142577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126101425760040190565b919082519283825260005b848110611dd55750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611d96565b90600182811c92168015611e33575b6020831014611e0457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611df9565b9060405191826000825492611e5184611dea565b8084529360018116908115611ebf5750600114611e78575b50611e7692500383611a8b565b565b90506000929192526020600020906000915b818310611ea3575050906020611e769282010138611e69565b6020919350806001915483858901015201910190918492611e8a565b60209350611e769592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e69565b8051821015611f135760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b818110611f4d575050565b60008155600101611f42565b80546000825580611f68575050565b611e7691600052602060002090810190611f42565b906801000000000000000081116114b157815490808355818110611fa057505050565b611e769260005260206000209182019101611f42565b9190601f8111611fc557505050565b611e76926000526020600020906020601f840160051c83019310611ff1575b601f0160051c0190611f42565b9091508190611fe4565b6040519061200a602083611a8b565b600080835282815b82811061201e57505050565b60209060405161202d81611a53565b6000815260008382015282828501015201612012565b9067ffffffffffffffff90939293168152604060208201526120a8612074845160a0604085015260e0840190611d8b565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152611d8b565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b81811061214c5750505060808473ffffffffffffffffffffffffffffffffffffffff6060611c48969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152611d8b565b8251805173ffffffffffffffffffffffffffffffffffffffff16865260209081015181870152604090950194909201916001016120e9565b90816020910312610142575180151581036101425790565b906040519182815491828252602082019060005260206000209260005b8181106121ce575050611e7692500383611a8b565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016121b9565b3567ffffffffffffffff811681036101425790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561014257016020813591019167ffffffffffffffff821161014257813603831361014257565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b3d156122cc573d906122b282611acc565b916122c06040519384611a8b565b82523d6000602084013e565b606090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610142570180359067ffffffffffffffff82116101425760200191813603831361014257565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610142570180359067ffffffffffffffff821161014257602001918160061b3603831361014257565b3573ffffffffffffffffffffffffffffffffffffffff811681036101425790565b303b15610142576040517fcf6730f8000000000000000000000000000000000000000000000000000000008152600091602060048301528035928360248401526020820192833567ffffffffffffffff81168103612a785767ffffffffffffffff16604482015260408301936124216124108686612212565b60a0606486015260c4850191612262565b61246360608601916124338388612212565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc878403016084880152612262565b92608086019384357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301811215612a745787016020813591019167ffffffffffffffff8211612a70578160061b36038313612a70578381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160a48501528181528791849160200190835b818110612a2b57505081929350038183305af19081612a17575b506129eb576125196122a1565b9587855260076020526001604086205561253288612c37565b508785526004602052604085209288845567ffffffffffffffff61255960018601926121fd565b167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082541617905561258f6002840191876122d1565b9067ffffffffffffffff82116129be576125b3826125ad8554611dea565b85611fb6565b8690601f831160011461291e576125fe929188918361283e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b61260f6003830191866122d1565b9067ffffffffffffffff82116128f15761262d826125ad8554611dea565b8590601f8311600114612849578261268b96959360049593612681938a9261283e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b0193612322565b929068010000000000000000841161281157815484835580851061275f575b5090825260208220905b8383106126fd57505050506126f87f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f91604051918291602083526020830190611d8b565b0390a2565b600260408273ffffffffffffffffffffffffffffffffffffffff612722600195612376565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000865416178555602081013584860155019201920191906126b4565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681036127e4577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff851685036127e457828452602084209060011b8101908560011b015b8181106127d257506126aa565b808560029255856001820155016127c5565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b0135905038806113a5565b83875260208720917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416885b8181106128d9575092600192859261268b99989660049896106128a1575b505050811b019055612684565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612894565b91936020600181928787013581550195019201612876565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b83885260208820917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416895b8181106129a6575090846001959493921061296e575b505050811b019055612601565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612961565b9193602060018192878701358155019501920161294b565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50505090507fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f915080a2565b85612a2491969296611a8b565b933861250c565b9250925060408060019273ffffffffffffffffffffffffffffffffffffffff612a53886119ff565b1681526020870135602082015201940191019288928592946124f2565b8780fd5b8680fd5b8280fd5b9190811015611f135760061b0190565b73ffffffffffffffffffffffffffffffffffffffff600154163303612aad57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff9384166024830152604480830195909552938152612b8e929091612b3c606484611a8b565b16600080604095865194612b508887611a8b565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af1612b886122a1565b91612cb4565b80519081612b9b57505050565b602080612bac938301019101612184565b15612bb45750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80600052600660205260406000205415600014612cae57600554680100000000000000008110156114b1576001810180600555811015611f13577f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00181905560055460009182526006602052604090912055600190565b50600090565b91929015612d2f5750815115612cc8575090565b3b15612cd15790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015612d425750805190602001fd5b612d80906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611d8b565b0390fdfea164736f6c634300081a000a",
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

func (_DefensiveExample *DefensiveExampleCaller) GetChainConfig(opts *bind.CallOpts, selector uint64) (CCIPClientExampleChainConfig, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getChainConfig", selector)

	if err != nil {
		return *new(CCIPClientExampleChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPClientExampleChainConfig)).(*CCIPClientExampleChainConfig)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) GetChainConfig(selector uint64) (CCIPClientExampleChainConfig, error) {
	return _DefensiveExample.Contract.GetChainConfig(&_DefensiveExample.CallOpts, selector)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetChainConfig(selector uint64) (CCIPClientExampleChainConfig, error) {
	return _DefensiveExample.Contract.GetChainConfig(&_DefensiveExample.CallOpts, selector)
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

func (_DefensiveExample *DefensiveExampleTransactor) DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "disableChain", chainSelector)
}

func (_DefensiveExample *DefensiveExampleSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableChain(&_DefensiveExample.TransactOpts, chainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableChain(&_DefensiveExample.TransactOpts, chainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactor) EnableChain(opts *bind.TransactOpts, chainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "enableChain", chainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleSession) EnableChain(chainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableChain(&_DefensiveExample.TransactOpts, chainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) EnableChain(chainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableChain(&_DefensiveExample.TransactOpts, chainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
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

	GetChainConfig(opts *bind.CallOpts, selector uint64) (CCIPClientExampleChainConfig, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SFeeToken(opts *bind.CallOpts) (common.Address, error)

	SMessageContents(opts *bind.CallOpts, messageId [32]byte) (SMessageContents,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error)

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
