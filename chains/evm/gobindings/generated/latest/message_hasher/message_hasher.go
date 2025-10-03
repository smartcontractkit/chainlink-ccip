// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_hasher

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

type ClientCCV struct {
	CcvAddress common.Address
	Args       []byte
}

type ClientEVMExtraArgsV1 struct {
	GasLimit *big.Int
}

type ClientEVMExtraArgsV3 struct {
	RequiredCCV       []ClientCCV
	OptionalCCV       []ClientCCV
	OptionalThreshold uint8
	FinalityConfig    uint16
	Executor          common.Address
	ExecutorArgs      []byte
	TokenArgs         []byte
}

type ClientGenericExtraArgsV2 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
}

type ClientSVMExtraArgsV1 struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}

type ClientSuiExtraArgsV1 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	ReceiverObjectIds        [][32]byte
}

var MessageHasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSVMExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSuiExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeGenericExtraArgsV3\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.EVMExtraArgsV3\",\"components\":[{\"name\":\"requiredCCV\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"optionalCCV\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSUIExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"structClient.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x60808060405234601557610d39908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806310ee47db1461060e5780631914fbd21461054c57806355ad01df146105035780635e5ad095146101fb5780636fa473e41461019757806381c6b88b14610134578063b17df7141461013e578063c63641bd14610139578063c6ba5f2814610139578063c7ca9a18146101345763e733d2091461009557600080fd5b3461012f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012f5761012b6040516100d38161068a565b6004358152604051907f97a657c9000000000000000000000000000000000000000000000000000000006020830152516024820152602481526101176044826106c2565b604051918291602083526020830190610857565b0390f35b600080fd5b610b5d565b610c03565b3461012f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012f57602060043560006040516101808161068a565b528060405161018e8161068a565b52604051908152f35b3461012f576101a5366108b6565b63ffffffff81511661012b67ffffffffffffffff6020840151169260408101511515906080606082015191015191604051958695865260208601526040850152606084015260a0608084015260a08301906109b0565b3461012f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012f5760043567ffffffffffffffff811161012f5760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261012f5760405160e0810181811067ffffffffffffffff8211176104d457604052816004013567ffffffffffffffff811161012f576102a99060043691850101610a7a565b8152602482013567ffffffffffffffff811161012f576102cf9060043691850101610a7a565b9160208201928352604481013560ff8116810361012f576040830190815260648201359161ffff8316830361012f5760608401928352610311608482016109e4565b6080850190815260a482013567ffffffffffffffff811161012f5761033c9060043691850101610a05565b9160a0860192835260c481013567ffffffffffffffff811161012f573691016004019061036891610a05565b9360c08601948552604051968796602088017f302326cb0000000000000000000000000000000000000000000000000000000090526024880160209052516044880160e0905261012488016103bc91610c8e565b9051908781037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160648901526103f291610c8e565b935160ff1660848701525161ffff1660a48601525173ffffffffffffffffffffffffffffffffffffffff1660c485015251908381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160e485015261045791610857565b9051908281037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0161010484015261048e91610857565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526104be90826106c2565b6040518091602082526020820161012b91610857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461012f5761051136610785565b805161012b602083015115159260606040820151910151906040519485948552602085015260408401526080606084015260808301906109b0565b3461012f5761012b6080610117610562366108b6565b6105e26040519384927f1f3b3aba0000000000000000000000000000000000000000000000000000000060208501526020602485015263ffffffff815116604485015267ffffffffffffffff6020820151166064850152604081015115156084850152606081015160a4850152015160a060c484015260e48301906109b0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826106c2565b3461012f5761012b606061011761062436610785565b6105e26040519384927f21ea4ca90000000000000000000000000000000000000000000000000000000060208501526020602485015280516044850152602081015115156064850152604081015160848501520151608060a484015260c48301906109b0565b6020810190811067ffffffffffffffff8211176104d457604052565b6040810190811067ffffffffffffffff8211176104d457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176104d457604052565b3590811515820361012f57565b67ffffffffffffffff81116104d45760051b60200190565b9080601f8301121561012f57813561073f81610710565b9261074d60405194856106c2565b81845260208085019260051b82010192831161012f57602001905b8282106107755750505090565b8135815260209182019101610768565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261012f5760043567ffffffffffffffff811161012f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828403011261012f57604051916080830183811067ffffffffffffffff8211176104d4576040528160040135835261081e60248301610703565b60208401526044820135604084015260648201359167ffffffffffffffff831161012f5761084f9201600401610728565b606082015290565b919082519283825260005b8481106108a15750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610862565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261012f5760043567ffffffffffffffff811161012f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828403011261012f576040519160a0830183811067ffffffffffffffff8211176104d457604052816004013563ffffffff8116810361012f578352602482013567ffffffffffffffff8116810361012f57602084015261097760448301610703565b60408401526064820135606084015260848201359167ffffffffffffffff831161012f576109a89201600401610728565b608082015290565b906020808351928381520192019060005b8181106109ce5750505090565b82518452602093840193909201916001016109c1565b359073ffffffffffffffffffffffffffffffffffffffff8216820361012f57565b81601f8201121561012f5780359067ffffffffffffffff82116104d45760405192610a5860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f86011601856106c2565b8284526020838301011161012f57816000926020809301838601378301015290565b9080601f8301121561012f57813591610a9283610710565b92610aa060405194856106c2565b80845260208085019160051b8301019183831161012f5760208101915b838310610acc57505050505090565b823567ffffffffffffffff811161012f5782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261012f5760405190610b19826106a6565b610b25602084016109e4565b825260408301359167ffffffffffffffff831161012f57610b4e88602080969581960101610a05565b83820152815201920191610abd565b3461012f57600060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c005760405190610b9b826106a6565b6004358252602435908115158203610c005760208084018381526040517f181dcf10000000000000000000000000000000000000000000000000000000009281019290925284516024830152511515604482015261012b9061011781606481016105e2565b80fd5b3461012f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012f5760243580151580910361012f5761012b9060006020604051610c53816106a6565b828152015260405190610c65826106a6565b600435825260208201526040519182918291909160208060408301948051845201511515910152565b9080602083519182815201916020808360051b8301019401926000915b838310610cba57505050505090565b9091929394602080610d1d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610857565b97019301930191939290610cab56fea164736f6c634300081a000a",
}

var MessageHasherABI = MessageHasherMetaData.ABI

var MessageHasherBin = MessageHasherMetaData.Bin

func DeployMessageHasher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHasher, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageHasherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHasher{address: address, abi: *parsed, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

type MessageHasher struct {
	address common.Address
	abi     abi.ABI
	MessageHasherCaller
	MessageHasherTransactor
	MessageHasherFilterer
}

type MessageHasherCaller struct {
	contract *bind.BoundContract
}

type MessageHasherTransactor struct {
	contract *bind.BoundContract
}

type MessageHasherFilterer struct {
	contract *bind.BoundContract
}

type MessageHasherSession struct {
	Contract     *MessageHasher
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageHasherCallerSession struct {
	Contract *MessageHasherCaller
	CallOpts bind.CallOpts
}

type MessageHasherTransactorSession struct {
	Contract     *MessageHasherTransactor
	TransactOpts bind.TransactOpts
}

type MessageHasherRaw struct {
	Contract *MessageHasher
}

type MessageHasherCallerRaw struct {
	Contract *MessageHasherCaller
}

type MessageHasherTransactorRaw struct {
	Contract *MessageHasherTransactor
}

func NewMessageHasher(address common.Address, backend bind.ContractBackend) (*MessageHasher, error) {
	abi, err := abi.JSON(strings.NewReader(MessageHasherABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageHasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHasher{address: address, abi: abi, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

func NewMessageHasherCaller(address common.Address, caller bind.ContractCaller) (*MessageHasherCaller, error) {
	contract, err := bindMessageHasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherCaller{contract: contract}, nil
}

func NewMessageHasherTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHasherTransactor, error) {
	contract, err := bindMessageHasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherTransactor{contract: contract}, nil
}

func NewMessageHasherFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageHasherFilterer, error) {
	contract, err := bindMessageHasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageHasherFilterer{contract: contract}, nil
}

func bindMessageHasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MessageHasher *MessageHasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.MessageHasherCaller.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCaller) DecodeEVMExtraArgsV1(opts *bind.CallOpts, gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeEVMExtraArgsV1", gasLimit)

	if err != nil {
		return *new(ClientEVMExtraArgsV1), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientEVMExtraArgsV1)).(*ClientEVMExtraArgsV1)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeEVMExtraArgsV1(gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV1(&_MessageHasher.CallOpts, gasLimit)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeEVMExtraArgsV1(gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV1(&_MessageHasher.CallOpts, gasLimit)
}

func (_MessageHasher *MessageHasherCaller) DecodeEVMExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeEVMExtraArgsV2", gasLimit, allowOutOfOrderExecution)

	if err != nil {
		return *new(ClientGenericExtraArgsV2), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientGenericExtraArgsV2)).(*ClientGenericExtraArgsV2)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeEVMExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeEVMExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCaller) DecodeGenericExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeGenericExtraArgsV2", gasLimit, allowOutOfOrderExecution)

	if err != nil {
		return *new(ClientGenericExtraArgsV2), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientGenericExtraArgsV2)).(*ClientGenericExtraArgsV2)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeGenericExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeGenericExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeGenericExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeGenericExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCaller) DecodeSVMExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeSVMExtraArgsStruct", extraArgs)

	outstruct := new(DecodeSVMExtraArgsStruct)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ComputeUnits = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AccountIsWritableBitmap = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AllowOutOfOrderExecution = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.TokenReceiver = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Accounts = *abi.ConvertType(out[4], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_MessageHasher *MessageHasherSession) DecodeSVMExtraArgsStruct(extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSVMExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeSVMExtraArgsStruct(extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSVMExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) DecodeSuiExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeSuiExtraArgsStruct", extraArgs)

	outstruct := new(DecodeSuiExtraArgsStruct)
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AllowOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TokenReceiver = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.ReceiverObjectIds = *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_MessageHasher *MessageHasherSession) DecodeSuiExtraArgsStruct(extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSuiExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeSuiExtraArgsStruct(extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSuiExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeEVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeEVMExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeEVMExtraArgsV1(extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeEVMExtraArgsV1(extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeEVMExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeEVMExtraArgsV2", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeEVMExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeEVMExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeGenericExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeGenericExtraArgsV2", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeGenericExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeGenericExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeGenericExtraArgsV3(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV3) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeGenericExtraArgsV3", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeGenericExtraArgsV3(extraArgs ClientEVMExtraArgsV3) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV3(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeGenericExtraArgsV3(extraArgs ClientEVMExtraArgsV3) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV3(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeSUIExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeSUIExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeSUIExtraArgsV1(extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSUIExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeSUIExtraArgsV1(extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSUIExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeSVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeSVMExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeSVMExtraArgsV1(extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeSVMExtraArgsV1(extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

type DecodeSVMExtraArgsStruct struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}
type DecodeSuiExtraArgsStruct struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	ReceiverObjectIds        [][32]byte
}

func (_MessageHasher *MessageHasher) Address() common.Address {
	return _MessageHasher.address
}

type MessageHasherInterface interface {
	DecodeEVMExtraArgsV1(opts *bind.CallOpts, gasLimit *big.Int) (ClientEVMExtraArgsV1, error)

	DecodeEVMExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error)

	DecodeGenericExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error)

	DecodeSVMExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

		error)

	DecodeSuiExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

		error)

	EncodeEVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV1) ([]byte, error)

	EncodeEVMExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeGenericExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeGenericExtraArgsV3(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV3) ([]byte, error)

	EncodeSUIExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) ([]byte, error)

	EncodeSVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) ([]byte, error)

	Address() common.Address
}
