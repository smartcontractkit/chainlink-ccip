// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_lombard_bridge

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

var MockLombardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MSG_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"optionalMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getAllowedDestinationToken\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_lastPayloadHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_mailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllowedDestinationToken\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMailbox\",\"inputs\":[{\"name\":\"mailbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60808060405234608257610bca8181016001600160401b03811183821017606c57829161078e833903906000f08015606057600080546001600160a01b0319166001600160a01b039290921691909117905560405161070690816100888239f35b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c9081628a1198146103df57508063353c26b7146103a55780636e48b60d14610332578063793ea55b146102015780639da0ed1e14610191578063a936a63f14610140578063d5438eae14610140578063ea845bfa146101045763f3c61d6b1461008357600080fd5b346101015760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101015760043573ffffffffffffffffffffffffffffffffffffffff81168091036100fd577fffffffffffffffffffffffff000000000000000000000000000000000000000082541617815580f35b5080fd5b80fd5b503461010157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610101576020600154604051908152f35b503461010157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101015773ffffffffffffffffffffffffffffffffffffffff6020915416604051908152f35b50346101015760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610101576101c96105e7565b6004358252600260205273ffffffffffffffffffffffffffffffffffffffff6040832091166000526020526040600020604435905580f35b5060c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610101576102346105e7565b61023c61060f565b506040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260843560448201526020816064818673ffffffffffffffffffffffffffffffffffffffff87165af1801561032757604093506102fa575b5081517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000602082019242845260601b1683820152603481526102e9605482610632565b519020815190600182526020820152f35b61031b9060203d602011610320575b6103138183610632565b8101906106a2565b6102a6565b503d610309565b6040513d85823e3d90fd5b50346101015760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101015773ffffffffffffffffffffffffffffffffffffffff60406103816105e7565b92600435815260026020522091166000526020526020604060002054604051908152f35b503461010157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261010157602060405160018152f35b8260e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610101576104126105e7565b61041a61060f565b5060c4359267ffffffffffffffff84116105c257366023850112156105c25783600401359167ffffffffffffffff83116105e357602485019460248436920101116105e3577f23b872dd000000000000000000000000000000000000000000000000000000008252336004830152306024830152608435604483015260209082906064908290879073ffffffffffffffffffffffffffffffffffffffff165af18015610327576105c6575b506040516020810190428252604080820152610515816104e96060820186896106ba565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610632565b51902060015573ffffffffffffffffffffffffffffffffffffffff82541690813b156105c25792829161058094836040518097819582947ffb81d98d0000000000000000000000000000000000000000000000000000000084526020600485015260248401916106ba565b03925af19182156105b557816040936105a5575b505060015482519182526020820152f35b6105ae91610632565b8281610594565b50604051903d90823e3d90fd5b8280fd5b6105de9060203d602011610320576103138183610632565b6104c5565b8380fd5b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361060a57565b600080fd5b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361060a57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761067357604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9081602091031261060a5751801515810361060a5790565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093818652868601376000858286010152011601019056fea164736f6c634300081a000a60808060405234602157600160ff1981541617600155610ba390816100278239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806336e75928146104d65780639e31ddb614610468578063a6208506146103c7578063fb81d98d146101795763ff2f1e461461005357600080fd5b346101745760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101745760405160006003546100938161077f565b808452906001811690811561013257506001146100d3575b6100cf836100bb818503826107d2565b604051918291602083526020830190610720565b0390f35b600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b808210610118575090915081016020016100bb6100ab565b919260018160209254838588010152019101909291610100565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506100bb90506100ab565b600080fd5b346101745760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101745760043567ffffffffffffffff8111610174576101c89036906004016106f2565b67ffffffffffffffff8111610398576101e260025461077f565b601f81116102fa575b506000601f821160011461024557819260009261023a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c191617600255600080f35b013590508280610203565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216927f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91805b8581106102e2575083600195106102aa575b505050811b01600255005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905582808061029f565b9092602060018192868601358155019401910161028d565b601f820160051c7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace019060208310610370575b601f0160051c7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905b81811061036457506101eb565b60008155600101610357565b7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915061032d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101745760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101745760043567ffffffffffffffff8111610174576104169036906004016106f2565b6024359067ffffffffffffffff8211610174576100cf9261043e6104459336906004016106f2565b5050610878565b604093919351938493845215156020840152606060408401526060830190610720565b346101745760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610174576004358015158091036101745760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060015416911617600155600080f35b346101745760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610174576024358015158091036101745760443567ffffffffffffffff8111610174576105329036906004016106f2565b909160043560005560ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006001541691161760015567ffffffffffffffff81116103985761058160025461077f565b601f8111610654575b506000601f82116001146105d857819260009261023a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c191617600255600080f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216927f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91805b85811061063c575083600195106102aa57505050811b01600255005b90926020600181928686013581550194019101610620565b601f820160051c7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0190602083106106ca575b601f0160051c7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905b8181106106be575061058a565b600081556001016106b1565b7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9150610687565b9181601f840112156101745782359167ffffffffffffffff8311610174576020838186019501011161017457565b919082519283825260005b84811061076a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161072b565b90600182811c921680156107c8575b602083101461079957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161078e565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761039857604052565b92919267ffffffffffffffff8211610398576040519161085b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601846107d2565b829481845281830111610174578281602093846000960137010152565b67ffffffffffffffff82116103985761089260035461077f565b601f8111610af3575b50816000601f8211600114610a3657600091610a2b575b508260011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8460031b1c1916176003555b6000549182158080610a19575b6109f5576109dc575050905b60ff6001541690604051600081600254916109188361077f565b808352926001811690811561099f575060011461093f575b61093c925003826107d2565b90565b506002600090815290917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b81831061098357505090602061093c92820101610930565b602091935080600191548385880101520191019091839261096b565b6020925061093c9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101610930565b6109e892503691610813565b60208151910120906108fe565b5091509161093c610a07368486610813565b60208151910120936001933691610813565b50610a2560025461077f565b156108f2565b9050810135386108b2565b600381527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9184907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016825b818110610ad8575010610aa0575b5050600182811b016003556108e5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19908301351690553880610a90565b85840135855560019094019360209384019387935001610a82565b6003600052601f830160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208410610b6e575b601f0160051c7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01905b818110610b62575061089b565b60008155600101610b55565b7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9150610b2b56fea164736f6c634300081a000a",
}

var MockLombardBridgeABI = MockLombardBridgeMetaData.ABI

var MockLombardBridgeBin = MockLombardBridgeMetaData.Bin

func DeployMockLombardBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockLombardBridge, error) {
	parsed, err := MockLombardBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockLombardBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockLombardBridge{address: address, abi: *parsed, MockLombardBridgeCaller: MockLombardBridgeCaller{contract: contract}, MockLombardBridgeTransactor: MockLombardBridgeTransactor{contract: contract}, MockLombardBridgeFilterer: MockLombardBridgeFilterer{contract: contract}}, nil
}

type MockLombardBridge struct {
	address common.Address
	abi     abi.ABI
	MockLombardBridgeCaller
	MockLombardBridgeTransactor
	MockLombardBridgeFilterer
}

type MockLombardBridgeCaller struct {
	contract *bind.BoundContract
}

type MockLombardBridgeTransactor struct {
	contract *bind.BoundContract
}

type MockLombardBridgeFilterer struct {
	contract *bind.BoundContract
}

type MockLombardBridgeSession struct {
	Contract     *MockLombardBridge
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockLombardBridgeCallerSession struct {
	Contract *MockLombardBridgeCaller
	CallOpts bind.CallOpts
}

type MockLombardBridgeTransactorSession struct {
	Contract     *MockLombardBridgeTransactor
	TransactOpts bind.TransactOpts
}

type MockLombardBridgeRaw struct {
	Contract *MockLombardBridge
}

type MockLombardBridgeCallerRaw struct {
	Contract *MockLombardBridgeCaller
}

type MockLombardBridgeTransactorRaw struct {
	Contract *MockLombardBridgeTransactor
}

func NewMockLombardBridge(address common.Address, backend bind.ContractBackend) (*MockLombardBridge, error) {
	abi, err := abi.JSON(strings.NewReader(MockLombardBridgeABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockLombardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockLombardBridge{address: address, abi: abi, MockLombardBridgeCaller: MockLombardBridgeCaller{contract: contract}, MockLombardBridgeTransactor: MockLombardBridgeTransactor{contract: contract}, MockLombardBridgeFilterer: MockLombardBridgeFilterer{contract: contract}}, nil
}

func NewMockLombardBridgeCaller(address common.Address, caller bind.ContractCaller) (*MockLombardBridgeCaller, error) {
	contract, err := bindMockLombardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockLombardBridgeCaller{contract: contract}, nil
}

func NewMockLombardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*MockLombardBridgeTransactor, error) {
	contract, err := bindMockLombardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockLombardBridgeTransactor{contract: contract}, nil
}

func NewMockLombardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*MockLombardBridgeFilterer, error) {
	contract, err := bindMockLombardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockLombardBridgeFilterer{contract: contract}, nil
}

func bindMockLombardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockLombardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockLombardBridge *MockLombardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockLombardBridge.Contract.MockLombardBridgeCaller.contract.Call(opts, result, method, params...)
}

func (_MockLombardBridge *MockLombardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.MockLombardBridgeTransactor.contract.Transfer(opts)
}

func (_MockLombardBridge *MockLombardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.MockLombardBridgeTransactor.contract.Transact(opts, method, params...)
}

func (_MockLombardBridge *MockLombardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockLombardBridge.Contract.contract.Call(opts, result, method, params...)
}

func (_MockLombardBridge *MockLombardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.contract.Transfer(opts)
}

func (_MockLombardBridge *MockLombardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.contract.Transact(opts, method, params...)
}

func (_MockLombardBridge *MockLombardBridgeCaller) MSGVERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockLombardBridge.contract.Call(opts, &out, "MSG_VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MockLombardBridge *MockLombardBridgeSession) MSGVERSION() (uint8, error) {
	return _MockLombardBridge.Contract.MSGVERSION(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCallerSession) MSGVERSION() (uint8, error) {
	return _MockLombardBridge.Contract.MSGVERSION(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCaller) GetAllowedDestinationToken(opts *bind.CallOpts, destinationChain [32]byte, sourceToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _MockLombardBridge.contract.Call(opts, &out, "getAllowedDestinationToken", destinationChain, sourceToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockLombardBridge *MockLombardBridgeSession) GetAllowedDestinationToken(destinationChain [32]byte, sourceToken common.Address) ([32]byte, error) {
	return _MockLombardBridge.Contract.GetAllowedDestinationToken(&_MockLombardBridge.CallOpts, destinationChain, sourceToken)
}

func (_MockLombardBridge *MockLombardBridgeCallerSession) GetAllowedDestinationToken(destinationChain [32]byte, sourceToken common.Address) ([32]byte, error) {
	return _MockLombardBridge.Contract.GetAllowedDestinationToken(&_MockLombardBridge.CallOpts, destinationChain, sourceToken)
}

func (_MockLombardBridge *MockLombardBridgeCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLombardBridge.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLombardBridge *MockLombardBridgeSession) Mailbox() (common.Address, error) {
	return _MockLombardBridge.Contract.Mailbox(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCallerSession) Mailbox() (common.Address, error) {
	return _MockLombardBridge.Contract.Mailbox(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCaller) SLastPayloadHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockLombardBridge.contract.Call(opts, &out, "s_lastPayloadHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockLombardBridge *MockLombardBridgeSession) SLastPayloadHash() ([32]byte, error) {
	return _MockLombardBridge.Contract.SLastPayloadHash(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCallerSession) SLastPayloadHash() ([32]byte, error) {
	return _MockLombardBridge.Contract.SLastPayloadHash(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCaller) SMailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLombardBridge.contract.Call(opts, &out, "s_mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLombardBridge *MockLombardBridgeSession) SMailbox() (common.Address, error) {
	return _MockLombardBridge.Contract.SMailbox(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeCallerSession) SMailbox() (common.Address, error) {
	return _MockLombardBridge.Contract.SMailbox(&_MockLombardBridge.CallOpts)
}

func (_MockLombardBridge *MockLombardBridgeTransactor) Deposit(opts *bind.TransactOpts, arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte, optionalMessage []byte) (*types.Transaction, error) {
	return _MockLombardBridge.contract.Transact(opts, "deposit", arg0, token, arg2, arg3, amount, arg5, optionalMessage)
}

func (_MockLombardBridge *MockLombardBridgeSession) Deposit(arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte, optionalMessage []byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.Deposit(&_MockLombardBridge.TransactOpts, arg0, token, arg2, arg3, amount, arg5, optionalMessage)
}

func (_MockLombardBridge *MockLombardBridgeTransactorSession) Deposit(arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte, optionalMessage []byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.Deposit(&_MockLombardBridge.TransactOpts, arg0, token, arg2, arg3, amount, arg5, optionalMessage)
}

func (_MockLombardBridge *MockLombardBridgeTransactor) Deposit0(opts *bind.TransactOpts, arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.contract.Transact(opts, "deposit0", arg0, token, arg2, arg3, amount, arg5)
}

func (_MockLombardBridge *MockLombardBridgeSession) Deposit0(arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.Deposit0(&_MockLombardBridge.TransactOpts, arg0, token, arg2, arg3, amount, arg5)
}

func (_MockLombardBridge *MockLombardBridgeTransactorSession) Deposit0(arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.Deposit0(&_MockLombardBridge.TransactOpts, arg0, token, arg2, arg3, amount, arg5)
}

func (_MockLombardBridge *MockLombardBridgeTransactor) SetAllowedDestinationToken(opts *bind.TransactOpts, destinationChain [32]byte, sourceToken common.Address, destinationToken [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.contract.Transact(opts, "setAllowedDestinationToken", destinationChain, sourceToken, destinationToken)
}

func (_MockLombardBridge *MockLombardBridgeSession) SetAllowedDestinationToken(destinationChain [32]byte, sourceToken common.Address, destinationToken [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.SetAllowedDestinationToken(&_MockLombardBridge.TransactOpts, destinationChain, sourceToken, destinationToken)
}

func (_MockLombardBridge *MockLombardBridgeTransactorSession) SetAllowedDestinationToken(destinationChain [32]byte, sourceToken common.Address, destinationToken [32]byte) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.SetAllowedDestinationToken(&_MockLombardBridge.TransactOpts, destinationChain, sourceToken, destinationToken)
}

func (_MockLombardBridge *MockLombardBridgeTransactor) SetMailbox(opts *bind.TransactOpts, mailbox_ common.Address) (*types.Transaction, error) {
	return _MockLombardBridge.contract.Transact(opts, "setMailbox", mailbox_)
}

func (_MockLombardBridge *MockLombardBridgeSession) SetMailbox(mailbox_ common.Address) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.SetMailbox(&_MockLombardBridge.TransactOpts, mailbox_)
}

func (_MockLombardBridge *MockLombardBridgeTransactorSession) SetMailbox(mailbox_ common.Address) (*types.Transaction, error) {
	return _MockLombardBridge.Contract.SetMailbox(&_MockLombardBridge.TransactOpts, mailbox_)
}

func (_MockLombardBridge *MockLombardBridge) Address() common.Address {
	return _MockLombardBridge.address
}

type MockLombardBridgeInterface interface {
	MSGVERSION(opts *bind.CallOpts) (uint8, error)

	GetAllowedDestinationToken(opts *bind.CallOpts, destinationChain [32]byte, sourceToken common.Address) ([32]byte, error)

	Mailbox(opts *bind.CallOpts) (common.Address, error)

	SLastPayloadHash(opts *bind.CallOpts) ([32]byte, error)

	SMailbox(opts *bind.CallOpts) (common.Address, error)

	Deposit(opts *bind.TransactOpts, arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte, optionalMessage []byte) (*types.Transaction, error)

	Deposit0(opts *bind.TransactOpts, arg0 [32]byte, token common.Address, arg2 common.Address, arg3 [32]byte, amount *big.Int, arg5 [32]byte) (*types.Transaction, error)

	SetAllowedDestinationToken(opts *bind.TransactOpts, destinationChain [32]byte, sourceToken common.Address, destinationToken [32]byte) (*types.Transaction, error)

	SetMailbox(opts *bind.TransactOpts, mailbox_ common.Address) (*types.Transaction, error)

	Address() common.Address
}
