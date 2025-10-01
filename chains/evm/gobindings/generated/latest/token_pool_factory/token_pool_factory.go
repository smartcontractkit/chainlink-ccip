// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_pool_factory

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

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type TokenPoolFactoryRemoteChainConfig struct {
	RemotePoolFactory   common.Address
	RemoteRouter        common.Address
	RemoteRMNProxy      common.Address
	RemoteTokenDecimals uint8
}

type TokenPoolFactoryRemoteTokenPoolInfo struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	RemotePoolInitCode  []byte
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	PoolType            uint8
	RemoteTokenAddress  []byte
	RemoteTokenInitCode []byte
	RateLimiterConfig   RateLimiterConfig
}

var TokenPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contractITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contractRegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enumTokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enumTokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RemoteChainConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101003461011b57601f6115bd38819003918201601f19168301916001600160401b038311848410176101205780849260809460405283398101031261011b5780516001600160a01b038116919082810361011b576020820151906001600160a01b03821680830361011b57610083606061007c60408701610136565b9501610136565b9415908115610112575b508015610101575b80156100f0575b6100df5760805260a05260c05260e052604051611472908161014b823960805181610280015260a051816101fb015260c05181610fff015260e05181610fdd0152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b0384161561009c565b506001600160a01b03831615610095565b9050153861008d565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361011b5756fe60c080604052600436101561001357600080fd5b600090813560e01c908163068d792e1461056857508063181f5a77146104e95763eb03cac11461004257600080fd5b346104515760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104515760043567ffffffffffffffff811161045f576100919036906004016106af565b9061009a61069f565b9160443567ffffffffffffffff81116104e5576100bb903690600401610855565b906064359267ffffffffffffffff84116104e1576100e061016c9436906004016106e0565b929091610165604051602081019061015b8161012f3360843586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107ae565b51902095866113c1565b9687610a0e565b9173ffffffffffffffffffffffffffffffffffffffff821690813b15610451576040517fc630948d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152818160248183875af18015610454576104d1575b509073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610483578280916024604051809481937f96ea2f7a0000000000000000000000000000000000000000000000000000000083528760048401525af18015610478579083916104bc575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610483576040517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af1801561049c579084916104a7575b5050803b15610483576040517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015286166024820152838160448183865af1801561049c57908491610487575b5050803b15610483576040517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201523360248201529083908290604490829084905af1801561047857908391610463575b5050803b1561045f578180916024604051809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af180156104545761043c575b50506040805173ffffffffffffffffffffffffffffffffffffffff928316815292909116602083015290f35b0390f35b6104478280926107ae565b610451578061040c565b80fd5b6040513d84823e3d90fd5b5080fd5b8161046d916107ae565b61045f5781386103c1565b6040513d85823e3d90fd5b8280fd5b81610491916107ae565b610483578238610354565b6040513d86823e3d90fd5b816104b1916107ae565b6104835782386102ea565b816104c6916107ae565b61045f578138610267565b816104db916107ae565b386101e2565b8580fd5b8480fd5b503461045157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610451575061043860405161052a6040826107ae565b601681527f546f6b656e506f6f6c466163746f727920312e352e31000000000000000000006020820152604051918291602083526020830190610812565b82346104515760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610451576004359173ffffffffffffffffffffffffffffffffffffffff8316830361045f576105c161069f565b9160443567ffffffffffffffff811161045f576105e29036906004016106af565b916064359067ffffffffffffffff821161045157602061065b888888888861060d3660048b016106e0565b6084358986019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015290959194610652816034840161012f565b51902095610a0e565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b359073ffffffffffffffffffffffffffffffffffffffff8216820361069a57565b600080fd5b6024359060ff8216820361069a57565b9181601f8401121561069a5782359167ffffffffffffffff831161069a576020808501948460051b01011161069a57565b9181601f8401121561069a5782359167ffffffffffffffff831161069a576020838186019501011161069a57565b6060810190811067ffffffffffffffff82111761072a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761072a57604052565b610100810190811067ffffffffffffffff82111761072a57604052565b6080810190811067ffffffffffffffff82111761072a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761072a57604052565b60005b8381106108025750506000910152565b81810151838201526020016107f2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361084e815180928187528780880191016107ef565b0116010190565b81601f8201121561069a5780359067ffffffffffffffff821161072a57604051926108a8601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856107ae565b8284526020838301011161069a57816000926020809301838601378301015290565b67ffffffffffffffff811161072a5760051b60200190565b604051906108ef8261070e565b60006040838281528260208201520152565b35906fffffffffffffffffffffffffffffffff8216820361069a57565b80511561092b5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561092b5760209160051b010190565b939594929160ff9073ffffffffffffffffffffffffffffffffffffffff60a087019316865216602085015260a060408501528151809152602060c0850192019060005b8181106109e25750505060809173ffffffffffffffffffffffffffffffffffffffff80929616606085015216910152565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016109b1565b9692959490919394608052600094610a25876108ca565b97610a33604051998a6107ae565b8789527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610a60896108ca565b01875b81811061137b575050604051610a7881610775565b8781526060602082015260606040820152604051610a9581610792565b8881528860208201528860408201528860608201526060820152876080820152606060a0820152606060c0820152610acb6108e2565b60e09190910152604060a052865b88811015610f9f578060051b8701357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6188360301811215610f9b578701803603906101a08212610f975760a0515191610b3183610775565b813567ffffffffffffffff81168103610f93578352602082013567ffffffffffffffff8111610f9357610b679036908401610855565b602084015260a05182013567ffffffffffffffff8111610f9357610b8e9036908401610855565b9060a051840191825260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0820112610f935760a0515192610bcf84610792565b610bdb60608201610679565b8452610be960808201610679565b6020850152610bfa60a08201610679565b60a05185015260c081013560ff81168103610f8f5760608501526060850193845260e08101356002811015610f8f57608086015261010081013567ffffffffffffffff8111610f8f57610c509036908301610855565b60a086015261012081013567ffffffffffffffff8111610f8f577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec0610c9a60609236908501610855565b9360c088019485520112610f1e5760a0515190610cb68261070e565b6101408101358015158103610f8b57610ce791610180918452610cdc6101608201610901565b602085015201610901565b60a05182015260e085015260a08401515115610f22575b5060208301515115610dd8575b505060a05151610d1d60a051826107ae565b60018152895b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060a051018110610dc7575090600192916020820151610d628261091e565b52610d6c8161091e565b5067ffffffffffffffff8251169160e060a08201519101519160a0515193610d9385610759565b8452602084015260a0518301528060608301526080820152610db5828d61095a565b52610dc0818c61095a565b5001610ad9565b806060602080938501015201610d23565b51815160a0840151602081805181010312610f1e57602001519073ffffffffffffffffffffffffffffffffffffffff8216809203610f1e57610ee4928d9492610ed483610e7860ff606073ffffffffffffffffffffffffffffffffffffffff970151169161012f60209a8b9260a0515191610e5385846107ae565b80835236813789848160a05184015116920151169160a051519687958601998a61096e565b8760a05151938492610ea683610e97818701998a8151938492016107ef565b850191518093858401906107ef565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107ae565b5190209151511690608051611445565b9073ffffffffffffffffffffffffffffffffffffffff60a051519216818301528152610f1260a051826107ae565b60208201523880610d0b565b8c80fd5b610f5090516020815191012073ffffffffffffffffffffffffffffffffffffffff8451511690608051611445565b73ffffffffffffffffffffffffffffffffffffffff60a051519116602082015260208152610f8060a051826107ae565b60a084015238610cfe565b8e80fd5b8d80fd5b8b80fd5b8980fd5b8880fd5b506110739293975061107b9498969550602097889161105160a0515191610fc685846107ae565b8983528936813761102560a05151948592878401957f0000000000000000000000000000000000000000000000000000000000000000927f0000000000000000000000000000000000000000000000000000000000000000928861096e565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836107ae565b610ea68660a0515197889686880137850191848301938a8552519384916107ef565b6080516113c1565b9273ffffffffffffffffffffffffffffffffffffffff84169260a051516110a283826107ae565b83815283368137843b15611377579091839060a051519384917fe8a1da170000000000000000000000000000000000000000000000000000000083526044830160a051600485015285518091528160648501960190855b818110611353575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8386030160248401528151908186528086019581808460051b83010194019686915b8483106111e657505050505081929350038183875af180156111ca576111d6575b5090803b1561045f57818091602460a05151809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af180156111ca576111b757505090565b6111c28280926107ae565b610451575090565b60a051513d84823e3d90fd5b816111e0916107ae565b38611167565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe091949750809396508592950301835284875191610120810167ffffffffffffffff84511682528b8385015191610120858501528251809152610140840190856101408260051b8701019401925b8181106113075750505050926112f4839260c0608061128360019860a05187015185820360a051870152610812565b946112c0606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193018895929693889592611146565b91939580611342887ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec089600196989a9b030188528851610812565b960194019101918a95949392611254565b825167ffffffffffffffff16885296830196899650889550918301916001016110f9565b8380fd5b808b602080936040519261138e84610759565b8d8452606083850152606060408501526113a66108e2565b60608501526113b36108e2565b608085015201015201610a63565b9080511561141b576020815191016000f59073ffffffffffffffffffffffffffffffffffffffff8216156113f157565b7f741752c20000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b90605592600b92604051926040840152602083015281520160ff8153209056fea164736f6c634300081a000a",
}

var TokenPoolFactoryABI = TokenPoolFactoryMetaData.ABI

var TokenPoolFactoryBin = TokenPoolFactoryMetaData.Bin

func DeployTokenPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address, tokenAdminModule common.Address, rmnProxy common.Address, ccipRouter common.Address) (common.Address, *types.Transaction, *TokenPoolFactory, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPoolFactoryBin), backend, tokenAdminRegistry, tokenAdminModule, rmnProxy, ccipRouter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPoolFactory{address: address, abi: *parsed, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

type TokenPoolFactory struct {
	address common.Address
	abi     abi.ABI
	TokenPoolFactoryCaller
	TokenPoolFactoryTransactor
	TokenPoolFactoryFilterer
}

type TokenPoolFactoryCaller struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryTransactor struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryFilterer struct {
	contract *bind.BoundContract
}

type TokenPoolFactorySession struct {
	Contract     *TokenPoolFactory
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryCallerSession struct {
	Contract *TokenPoolFactoryCaller
	CallOpts bind.CallOpts
}

type TokenPoolFactoryTransactorSession struct {
	Contract     *TokenPoolFactoryTransactor
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryRaw struct {
	Contract *TokenPoolFactory
}

type TokenPoolFactoryCallerRaw struct {
	Contract *TokenPoolFactoryCaller
}

type TokenPoolFactoryTransactorRaw struct {
	Contract *TokenPoolFactoryTransactor
}

func NewTokenPoolFactory(address common.Address, backend bind.ContractBackend) (*TokenPoolFactory, error) {
	abi, err := abi.JSON(strings.NewReader(TokenPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactory{address: address, abi: abi, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

func NewTokenPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenPoolFactoryCaller, error) {
	contract, err := bindTokenPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryCaller{contract: contract}, nil
}

func NewTokenPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPoolFactoryTransactor, error) {
	contract, err := bindTokenPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryTransactor{contract: contract}, nil
}

func NewTokenPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPoolFactoryFilterer, error) {
	contract, err := bindTokenPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryFilterer{contract: contract}, nil
}

func bindTokenPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.TokenPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenPoolFactory.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TokenPoolFactory *TokenPoolFactorySession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerSession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenAndTokenPool", remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenPoolWithExistingToken", token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt)
}

type TokenPoolFactoryRemoteChainConfigUpdatedIterator struct {
	Event *TokenPoolFactoryRemoteChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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
		it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolFactoryRemoteChainConfigUpdated struct {
	RemoteChainSelector uint64
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	Raw                 types.Log
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.FilterLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryRemoteChainConfigUpdatedIterator{contract: _TokenPoolFactory.contract, event: "RemoteChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.WatchLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolFactoryRemoteChainConfigUpdated)
				if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
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

func (_TokenPoolFactory *TokenPoolFactoryFilterer) ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error) {
	event := new(TokenPoolFactoryRemoteChainConfigUpdated)
	if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (TokenPoolFactoryRemoteChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xe3606343290c8e3853a7f92686979dfe24a0f29594186b61177236cb145126f0")
}

func (_TokenPoolFactory *TokenPoolFactory) Address() common.Address {
	return _TokenPoolFactory.address
}

type TokenPoolFactoryInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error)

	WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error)

	Address() common.Address
}
