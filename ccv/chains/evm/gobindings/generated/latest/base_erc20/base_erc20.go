// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base_erc20

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

type BaseERC20ConstructorParams struct {
	Name      string
	Symbol    string
	MaxSupply *big.Int
	PreMint   *big.Int
	Decimals  uint8
	CcipAdmin common.Address
}

var BaseERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct BaseERC20.ConstructorParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyCCIPAdmin\",\"inputs\":[]}]",
	Bin: "0x60c060405234610531576112ba8038038061001981610536565b928339810190602081830312610531578051906001600160401b03821161053157019060c082820312610531576040519060c082016001600160401b0381118382101761042e5760405282516001600160401b038111610531578161007f91850161055b565b82526020830151906001600160401b038211610531576100a091840161055b565b90816020820152604083015192604082019384526060810151926060830193845260808201519160ff83168303610531576080840192835260a00151926001600160a01b03841684036105315760a08101938452518051906001600160401b03821161042e5760035490600182811c92168015610527575b602083101461040e5781601f8493116104b7575b50602090601f831160011461044f57600092610444575b50508160011b916000199060031b1c1916176003555b8051906001600160401b03821161042e57600454600181811c91168015610424575b602082101461040e57601f81116103a9575b50602090601f831160011461033d5760ff93929160009183610332575b50508160011b916000199060031b1c1916176004555b5116608052825160a052516001600160a01b03168061032c575033915b8151908161024e575b600580546001600160a01b038681166001600160a01b0319831681179093556040519291167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3610cf390816105c782396080518161052c015260a051816101840152f35b518015159081610322575b5061030e5750516001600160a01b0382169081156102f8573082146102e3576002548181018091116102cd576002557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef602060009284845283825260408420818154019055604051908152a33880806101e6565b634e487b7160e01b600052601160045260246000fd5b50630bc2c5df60e11b60005260045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b63cbbf111360e01b60005260045260246000fd5b9050811138610259565b916101dd565b0151905038806101aa565b90601f198316916004600052816000209260005b818110610391575091600193918560ff97969410610378575b505050811b016004556101c0565b015160001960f88460031b161c1916905538808061036a565b92936020600181928786015181550195019301610351565b60046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610404575b601f0160051c01905b8181106103f8575061018d565b600081556001016103eb565b90915081906103e2565b634e487b7160e01b600052602260045260246000fd5b90607f169061017b565b634e487b7160e01b600052604160045260246000fd5b015190503880610143565b600360009081528281209350601f198516905b81811061049f5750908460019594939210610486575b505050811b01600355610159565b015160001960f88460031b161c19169055388080610478565b92936020600181928786015181550195019301610462565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c8101916020851061051d575b90601f859493920160051c01905b81811061050e575061012c565b60008155849350600101610501565b90915081906104f3565b91607f1691610118565b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761042e57604052565b81601f82011215610531578051906001600160401b03821161042e5761058a601f8301601f1916602001610536565b92828452602083830101116105315760005b8281106105b157505060206000918301015290565b8060208092840101518282870101520161059c56fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146109dc5750806306fdde03146108ff578063095ea7b31461081b57806318160ddd146107df578063181f5a771461072957806323b872dd14610550578063313ce567146104f457806370a082311461048f5780638fd6a6ac1461043d57806395d89b41146102e0578063a8fa343c146101f6578063a9059cbb146101a7578063d5abeb011461014e5763dd62ed3e146100b957600080fd5b346101495760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610149576100f0610b00565b73ffffffffffffffffffffffffffffffffffffffff61010d610b23565b9116600052600160205273ffffffffffffffffffffffffffffffffffffffff604060002091166000526020526020604060002054604051908152f35b600080fd5b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101495760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101495760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610149576101eb6101e1610b00565b6024359033610b87565b602060405160018152f35b346101495760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101495761022d610b00565b60055473ffffffffffffffffffffffffffffffffffffffff8116908133036102b65773ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffff0000000000000000000000000000000000000000931692839116176005557f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3005b7f69aec5370000000000000000000000000000000000000000000000000000000060005260046000fd5b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101495760405160006004548060011c90600181168015610433575b602083108114610406578285529081156103c45750600114610364575b6103608361035481850382610b46565b60405191829182610a98565b0390f35b91905060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b916000905b8082106103aa57509091508101602001610354610344565b919260018160209254838588010152019101909291610392565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506103549050610344565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f1691610327565b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957602073ffffffffffffffffffffffffffffffffffffffff60055416604051908152f35b346101495760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101495773ffffffffffffffffffffffffffffffffffffffff6104db610b00565b1660005260006020526020604060002054604051908152f35b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101495760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957610587610b00565b61058f610b23565b6044359073ffffffffffffffffffffffffffffffffffffffff831692836000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff33166000526020526040600020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811061060f575b506101eb9350610b87565b8381106106f3573033146106c5578415610696573315610667576101eb946000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff33166000526020528360406000209103905584610604565b7f94280d6200000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7fe602df0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7f17858bbe000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b83907ffb8f41b2000000000000000000000000000000000000000000000000000000006000523360045260245260445260646000fd5b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957604051604081019080821067ffffffffffffffff8311176107b05761036091604052601381527f42617365455243323020322e302e302d64657600000000000000000000000000602082015260405191829182610a98565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610149576020600254604051908152f35b346101495760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957610852610b00565b73ffffffffffffffffffffffffffffffffffffffff166024353082146108d157331561069657811561066757336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b507f17858bbe0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101495760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101495760405160006003548060011c906001811680156109d2575b602083108114610406578285529081156103c45750600114610972576103608361035481850382610b46565b91905060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b916000905b8082106109b857509091508101602001610354610344565b9192600181602092548385880101520191019092916109a0565b91607f1691610946565b346101495760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014957600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361014957817f36372b070000000000000000000000000000000000000000000000000000000060209314908115610a6e575b5015158152f35b7f8fd6a6ac0000000000000000000000000000000000000000000000000000000091501483610a67565b9190916020815282519283602083015260005b848110610aea5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201610aab565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361014957565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361014957565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176107b057604052565b73ffffffffffffffffffffffffffffffffffffffff16908115610cb75773ffffffffffffffffffffffffffffffffffffffff16918215610c8857308314610c5a576000828152806020526040812054828110610c275791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b6064937fe450d38c0000000000000000000000000000000000000000000000000000000083949352600452602452604452fd5b827f17858bbe0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fec442f0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7f96c6fd1e00000000000000000000000000000000000000000000000000000000600052600060045260246000fdfea164736f6c634300081a000a",
}

var BaseERC20ABI = BaseERC20MetaData.ABI

var BaseERC20Bin = BaseERC20MetaData.Bin

func DeployBaseERC20(auth *bind.TransactOpts, backend bind.ContractBackend, args BaseERC20ConstructorParams) (common.Address, *types.Transaction, *BaseERC20, error) {
	parsed, err := BaseERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseERC20Bin), backend, args)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseERC20{address: address, abi: *parsed, BaseERC20Caller: BaseERC20Caller{contract: contract}, BaseERC20Transactor: BaseERC20Transactor{contract: contract}, BaseERC20Filterer: BaseERC20Filterer{contract: contract}}, nil
}

type BaseERC20 struct {
	address common.Address
	abi     abi.ABI
	BaseERC20Caller
	BaseERC20Transactor
	BaseERC20Filterer
}

type BaseERC20Caller struct {
	contract *bind.BoundContract
}

type BaseERC20Transactor struct {
	contract *bind.BoundContract
}

type BaseERC20Filterer struct {
	contract *bind.BoundContract
}

type BaseERC20Session struct {
	Contract     *BaseERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BaseERC20CallerSession struct {
	Contract *BaseERC20Caller
	CallOpts bind.CallOpts
}

type BaseERC20TransactorSession struct {
	Contract     *BaseERC20Transactor
	TransactOpts bind.TransactOpts
}

type BaseERC20Raw struct {
	Contract *BaseERC20
}

type BaseERC20CallerRaw struct {
	Contract *BaseERC20Caller
}

type BaseERC20TransactorRaw struct {
	Contract *BaseERC20Transactor
}

func NewBaseERC20(address common.Address, backend bind.ContractBackend) (*BaseERC20, error) {
	abi, err := abi.JSON(strings.NewReader(BaseERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBaseERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseERC20{address: address, abi: abi, BaseERC20Caller: BaseERC20Caller{contract: contract}, BaseERC20Transactor: BaseERC20Transactor{contract: contract}, BaseERC20Filterer: BaseERC20Filterer{contract: contract}}, nil
}

func NewBaseERC20Caller(address common.Address, caller bind.ContractCaller) (*BaseERC20Caller, error) {
	contract, err := bindBaseERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseERC20Caller{contract: contract}, nil
}

func NewBaseERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BaseERC20Transactor, error) {
	contract, err := bindBaseERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseERC20Transactor{contract: contract}, nil
}

func NewBaseERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BaseERC20Filterer, error) {
	contract, err := bindBaseERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseERC20Filterer{contract: contract}, nil
}

func bindBaseERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BaseERC20 *BaseERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseERC20.Contract.BaseERC20Caller.contract.Call(opts, result, method, params...)
}

func (_BaseERC20 *BaseERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseERC20.Contract.BaseERC20Transactor.contract.Transfer(opts)
}

func (_BaseERC20 *BaseERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseERC20.Contract.BaseERC20Transactor.contract.Transact(opts, method, params...)
}

func (_BaseERC20 *BaseERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_BaseERC20 *BaseERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseERC20.Contract.contract.Transfer(opts)
}

func (_BaseERC20 *BaseERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseERC20.Contract.contract.Transact(opts, method, params...)
}

func (_BaseERC20 *BaseERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseERC20.Contract.Allowance(&_BaseERC20.CallOpts, owner, spender)
}

func (_BaseERC20 *BaseERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseERC20.Contract.Allowance(&_BaseERC20.CallOpts, owner, spender)
}

func (_BaseERC20 *BaseERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseERC20.Contract.BalanceOf(&_BaseERC20.CallOpts, account)
}

func (_BaseERC20 *BaseERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseERC20.Contract.BalanceOf(&_BaseERC20.CallOpts, account)
}

func (_BaseERC20 *BaseERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) Decimals() (uint8, error) {
	return _BaseERC20.Contract.Decimals(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) Decimals() (uint8, error) {
	return _BaseERC20.Contract.Decimals(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) GetCCIPAdmin() (common.Address, error) {
	return _BaseERC20.Contract.GetCCIPAdmin(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BaseERC20.Contract.GetCCIPAdmin(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) MaxSupply() (*big.Int, error) {
	return _BaseERC20.Contract.MaxSupply(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) MaxSupply() (*big.Int, error) {
	return _BaseERC20.Contract.MaxSupply(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) Name() (string, error) {
	return _BaseERC20.Contract.Name(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) Name() (string, error) {
	return _BaseERC20.Contract.Name(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseERC20.Contract.SupportsInterface(&_BaseERC20.CallOpts, interfaceId)
}

func (_BaseERC20 *BaseERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseERC20.Contract.SupportsInterface(&_BaseERC20.CallOpts, interfaceId)
}

func (_BaseERC20 *BaseERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) Symbol() (string, error) {
	return _BaseERC20.Contract.Symbol(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) Symbol() (string, error) {
	return _BaseERC20.Contract.Symbol(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) TotalSupply() (*big.Int, error) {
	return _BaseERC20.Contract.TotalSupply(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BaseERC20.Contract.TotalSupply(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseERC20.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BaseERC20 *BaseERC20Session) TypeAndVersion() (string, error) {
	return _BaseERC20.Contract.TypeAndVersion(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20CallerSession) TypeAndVersion() (string, error) {
	return _BaseERC20.Contract.TypeAndVersion(&_BaseERC20.CallOpts)
}

func (_BaseERC20 *BaseERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.contract.Transact(opts, "approve", spender, value)
}

func (_BaseERC20 *BaseERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.Approve(&_BaseERC20.TransactOpts, spender, value)
}

func (_BaseERC20 *BaseERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.Approve(&_BaseERC20.TransactOpts, spender, value)
}

func (_BaseERC20 *BaseERC20Transactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BaseERC20.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BaseERC20 *BaseERC20Session) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BaseERC20.Contract.SetCCIPAdmin(&_BaseERC20.TransactOpts, newAdmin)
}

func (_BaseERC20 *BaseERC20TransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BaseERC20.Contract.SetCCIPAdmin(&_BaseERC20.TransactOpts, newAdmin)
}

func (_BaseERC20 *BaseERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.contract.Transact(opts, "transfer", to, value)
}

func (_BaseERC20 *BaseERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.Transfer(&_BaseERC20.TransactOpts, to, value)
}

func (_BaseERC20 *BaseERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.Transfer(&_BaseERC20.TransactOpts, to, value)
}

func (_BaseERC20 *BaseERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BaseERC20 *BaseERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.TransferFrom(&_BaseERC20.TransactOpts, from, to, value)
}

func (_BaseERC20 *BaseERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseERC20.Contract.TransferFrom(&_BaseERC20.TransactOpts, from, to, value)
}

type BaseERC20ApprovalIterator struct {
	Event *BaseERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BaseERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseERC20Approval)
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
		it.Event = new(BaseERC20Approval)
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

func (it *BaseERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *BaseERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BaseERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BaseERC20 *BaseERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BaseERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BaseERC20ApprovalIterator{contract: _BaseERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BaseERC20 *BaseERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BaseERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BaseERC20Approval)
				if err := _BaseERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BaseERC20 *BaseERC20Filterer) ParseApproval(log types.Log) (*BaseERC20Approval, error) {
	event := new(BaseERC20Approval)
	if err := _BaseERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BaseERC20CCIPAdminTransferredIterator struct {
	Event *BaseERC20CCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BaseERC20CCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseERC20CCIPAdminTransferred)
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
		it.Event = new(BaseERC20CCIPAdminTransferred)
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

func (it *BaseERC20CCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BaseERC20CCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BaseERC20CCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BaseERC20 *BaseERC20Filterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BaseERC20CCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BaseERC20.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BaseERC20CCIPAdminTransferredIterator{contract: _BaseERC20.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BaseERC20 *BaseERC20Filterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BaseERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BaseERC20.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BaseERC20CCIPAdminTransferred)
				if err := _BaseERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BaseERC20 *BaseERC20Filterer) ParseCCIPAdminTransferred(log types.Log) (*BaseERC20CCIPAdminTransferred, error) {
	event := new(BaseERC20CCIPAdminTransferred)
	if err := _BaseERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BaseERC20TransferIterator struct {
	Event *BaseERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BaseERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseERC20Transfer)
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
		it.Event = new(BaseERC20Transfer)
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

func (it *BaseERC20TransferIterator) Error() error {
	return it.fail
}

func (it *BaseERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BaseERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BaseERC20 *BaseERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseERC20TransferIterator{contract: _BaseERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BaseERC20 *BaseERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaseERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BaseERC20Transfer)
				if err := _BaseERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BaseERC20 *BaseERC20Filterer) ParseTransfer(log types.Log) (*BaseERC20Transfer, error) {
	event := new(BaseERC20Transfer)
	if err := _BaseERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (BaseERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BaseERC20CCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BaseERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BaseERC20 *BaseERC20) Address() common.Address {
	return _BaseERC20.address
}

type BaseERC20Interface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BaseERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BaseERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BaseERC20Approval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BaseERC20CCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BaseERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BaseERC20CCIPAdminTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaseERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BaseERC20Transfer, error)

	Address() common.Address
}
