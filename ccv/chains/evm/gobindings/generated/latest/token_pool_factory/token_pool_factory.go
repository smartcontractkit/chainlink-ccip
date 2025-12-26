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
	RemoteLockBox       common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contract RegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101203461015c57601f6117f438819003918201601f19168301916001600160401b038311848410176101615780849260a09460405283398101031261015c5780516001600160a01b038116919082810361015c5760208201516001600160a01b0381169182820361015c5761007760408501610177565b92610090608061008960608801610177565b9601610177565b9515908115610153575b508015610142575b8015610131575b8015610120575b61010f5760805260a05260c05260e05261010052604051611668908161018c823960805181610234015260a0518161019a015260c0518181816110c001526114db015260e0518181816110e80152611503015261010051816111100152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b038516156100b0565b506001600160a01b038416156100a9565b506001600160a01b038316156100a2565b9050153861009a565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361015c5756fe60c080604052600436101561001357600080fd5b600090813560e01c908163111233601461059b57508063181f5a771461051c57632e1ab66c1461004257600080fd5b346104055760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104055760043567ffffffffffffffff8111610413576100919036906004016106f2565b9161009a6106e2565b916044359360028510156104375760643567ffffffffffffffff8111610518576100c890369060040161087c565b9160843567ffffffffffffffff811161051457906100ef61017c9493923690600401610723565b92909188610175604051602081019061016b8161013f3360a43586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107d5565b519020968761158b565b9889610995565b9215610485575b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610413576040517f96ea2f7a00000000000000000000000000000000000000000000000000000000815282816024818373ffffffffffffffffffffffffffffffffffffffff8916968760048401525af1801561042c57908391610470575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610437576040517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af180156104505790849161045b575b5050803b15610437576040517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015286166024820152838160448183865af180156104505790849161043b575b5050803b15610437576040517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201523360248201529083908290604490829084905af1801561042c57908391610417575b5050803b15610413578180916024604051809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af18015610408576103f0575b50506040805173ffffffffffffffffffffffffffffffffffffffff928316815292909116602083015290f35b0390f35b6103fb8280926107d5565b61040557806103c0565b80fd5b6040513d84823e3d90fd5b5080fd5b81610421916107d5565b610413578138610375565b6040513d85823e3d90fd5b8280fd5b81610445916107d5565b610437578238610308565b6040513d86823e3d90fd5b81610465916107d5565b61043757823861029e565b8161047a916107d5565b61041357813861021b565b73ffffffffffffffffffffffffffffffffffffffff8216803b15610413578180916024604051809481937fc630948d00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff8a1660048401525af18015610408578290610504575b5050610183565b61050d916107d5565b38816104fd565b8480fd5b8380fd5b503461040557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261040557506103ec60405161055d6040826107d5565b601681527f546f6b656e506f6f6c466163746f727920312e352e31000000000000000000006020820152604051918291602083526020830190610839565b82346104055760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610405576004359173ffffffffffffffffffffffffffffffffffffffff83168303610413576105f46106e2565b916044359060028210156104055760643567ffffffffffffffff8111610413576106229036906004016106f2565b9290916084359067ffffffffffffffff821161040557602061069e8989898989896106503660048c01610723565b60a4358a87019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015290969195610695816034840161013f565b51902096610995565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b359073ffffffffffffffffffffffffffffffffffffffff821682036106dd57565b600080fd5b6024359060ff821682036106dd57565b9181601f840112156106dd5782359167ffffffffffffffff83116106dd576020808501948460051b0101116106dd57565b9181601f840112156106dd5782359167ffffffffffffffff83116106dd57602083818601950101116106dd57565b6060810190811067ffffffffffffffff82111761076d57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761076d57604052565b610100810190811067ffffffffffffffff82111761076d57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761076d57604052565b60005b8381106108295750506000910152565b8181015183820152602001610819565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361087581518092818752878088019101610816565b0116010190565b81601f820112156106dd5780359067ffffffffffffffff821161076d57604051926108cf601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856107d5565b828452602083830101116106dd57816000926020809301838601378301015290565b67ffffffffffffffff811161076d5760051b60200190565b6040519061091682610751565b60006040838281528260208201520152565b35906fffffffffffffffffffffffffffffffff821682036106dd57565b8051156109525760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156109525760209160051b010190565b95909693979297608052600060a0526109af6080516108f1565b956109bd60405197886107d5565b60805187527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109ee6080516108f1565b0160a0515b81811061153e575050604051610a08816107b8565b60a05181526060602082015260606040820152604051610a278161079c565b60a051815260a051602082015260a051604082015260a051606082015260a0516080820152606082015260a0516080820152606060a0820152606060c0820152610a6f610909565b60e0919091015260a051604096905b60805181101561106157600581901b8b0135368c90037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4101811215610ff5576101c0818d01360312610ff557885190610ad6826107b8565b808d013567ffffffffffffffff81168103610ff55782528c81016020013567ffffffffffffffff8111610ff557818e610b158d9336908484010161087c565b602086015201013567ffffffffffffffff8111610ff5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0828f8d610b6160a09536908585010161087c565b908701520136030112610ff557895160e0828f610b7d8461079c565b610b8b6060838301016106bc565b8452610b9b6080838301016106bc565b60208501528d610baf60a0848401016106bc565b90850152610bc160c0838301016106bc565b60608501520101359060ff82168203610ff5578e816101009360808694015260608601520101356002811015610ff55760808301528c8101610120013567ffffffffffffffff8111610ff557818e610c216101409336908484010161087c565b60a08601520101359067ffffffffffffffff8211610ff55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea08f93610c6e849136908388010161087c565b9460c087019586520136030112610ff557806101608f928d5193610c9185610751565b0101358015158103610ff5578f91610cc4926101a0928552610cb861018083830101610928565b60208601520101610928565b8b82015260e083015260a08201515115610ffb575b5060208101515115610da9575b8851610cf28a826107d5565b6001815260a0515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08b018110610d98575090600192916020820151610d3782610945565b52610d4181610945565b5067ffffffffffffffff8251169160e060a0820151910151918c5193610d668561079c565b845260208401528b8301528060608301526080820152610d86828c610981565b52610d91818b610981565b5001610a7e565b806060602080938501015201610cfa565b8881015160608201519060a0830151602081805181010312610ff5576020015173ffffffffffffffffffffffffffffffffffffffff81168103610ff5576080840151906002821015610fc45773ffffffffffffffffffffffffffffffffffffffff8d610efc92610f2296600160209660a0515060a0515014600014610f5857610ea1928160ff608061013f9401511686838301511690876060818c86015116940151169351978896168a87019373ffffffffffffffffffffffffffffffffffffffff92908360a09560ff82949a99959a8360c08b019c168a521660208901526000604089015216606087015216608085015216910152565b8d5192839181610eba8185019788815193849201610816565b8301610ece82518093858085019101610816565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107d5565b51902073ffffffffffffffffffffffffffffffffffffffff60608401515116908a611625565b73ffffffffffffffffffffffffffffffffffffffff8a519116602082015260208152610f4e8a826107d5565b6020820152610ce6565b610fbf928160ff608061013f94015116868981858501511693015116925196879516898601929373ffffffffffffffffffffffffffffffffffffffff8092969560ff6080958360a089019a1688521660208701526000604087015216606085015216910152565b610ea1565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526021600452602460a051fd5b60a05180fd5b61102a90516020815191012073ffffffffffffffffffffffffffffffffffffffff60608401515116908a611625565b73ffffffffffffffffffffffffffffffffffffffff8a5191166020820152602081526110568a826107d5565b60a082015238610cd9565b50919493985091949660a051506002861015610fc4576020611197938192600161119c99146000146114a8578a5173ffffffffffffffffffffffffffffffffffffffff92831684820190815260ff9092166020830152600060408301527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f0000000000000000000000000000000000000000000000000000000000000000831660808301527f00000000000000000000000000000000000000000000000000000000000000009290921660a082015261114690829060c00161013f565b611169878b519889968588013785019183830160a0518152815194859201610816565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836107d5565b61158b565b8251909273ffffffffffffffffffffffffffffffffffffffff8416929091602091906111c883826107d5565b60a08051825251368137843b15610ff5578351927fe8a1da17000000000000000000000000000000000000000000000000000000008452604484018560048601528251809152816064860193019060a0515b81811061148a575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8483030160248501528251908183528083019281808460051b83010195019360a051915b84831061132057505050505050818060a05192038160a051875af1801561131457611302575b50813b15610ff5578051917ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401528260248160a0519360a051905af19081156112f757506112e1575090565b60a0516112ed916107d5565b60a051610ff55790565b513d60a051823e3d90fd5b60a05161130e916107d5565b3861128f565b82513d60a051823e3d90fd5b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301835284875191610120810167ffffffffffffffff84511682528b83850151916101208585015282518091526101408401856101408360051b87010194019160a0515b81811061143b5750505050839260c060806113ba88958561142b9660019b01519086830390870152610839565b946113f7606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193019194939290611269565b92959680919450611477867ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08a60019699030188528851610839565b960194019101908e928b9695949261138d565b825167ffffffffffffffff168552938301939183019160010161121a565b8a5173ffffffffffffffffffffffffffffffffffffffff92831684820190815260ff9092166020830152600060408301527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f000000000000000000000000000000000000000000000000000000000000000092909216608082015261153990829060a00161013f565b611146565b60209060409a989a516115508161079c565b60a05181526060838201526060604082015261156a610909565b6060820152611577610909565b608082015282828c010152019896986109f3565b908051156115fb576020815191016000f5903d15198215166115ef5773ffffffffffffffffffffffffffffffffffffffff8216156115c557565b7fb06ebf3d0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b60559173ffffffffffffffffffffffffffffffffffffffff93600b92604051926040840152602083015281520160ff815320169056fea164736f6c634300081a000a",
}

var TokenPoolFactoryABI = TokenPoolFactoryMetaData.ABI

var TokenPoolFactoryBin = TokenPoolFactoryMetaData.Bin

func DeployTokenPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address, tokenAdminModule common.Address, rmnProxy common.Address, ccipRouter common.Address, lockBox common.Address) (common.Address, *types.Transaction, *TokenPoolFactory, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPoolFactoryBin), backend, tokenAdminRegistry, tokenAdminModule, rmnProxy, ccipRouter, lockBox)
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

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenAndTokenPool", remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenPoolWithExistingToken", token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactory) Address() common.Address {
	return _TokenPoolFactory.address
}

type TokenPoolFactoryInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	Address() common.Address
}
