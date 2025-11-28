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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contract RegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RemoteChainConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101203461015c57601f6117d138819003918201601f19168301916001600160401b038311848410176101615780849260a09460405283398101031261015c5780516001600160a01b038116919082810361015c5760208201516001600160a01b0381169182820361015c5761007760408501610177565b92610090608061008960608801610177565b9601610177565b9515908115610153575b508015610142575b8015610131575b8015610120575b61010f5760805260a05260c05260e05261010052604051611645908161018c82396080518161028c015260a05181610207015260c05181818161109d01526114b8015260e0518181816110c501526114e0015261010051816110ed0152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b038516156100b0565b506001600160a01b038416156100a9565b506001600160a01b038316156100a2565b9050153861009a565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361015c5756fe60c080604052600436101561001357600080fd5b600090813560e01c908163111233601461057857508063181f5a77146104f957632e1ab66c1461004257600080fd5b3461045d5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261045d5760043567ffffffffffffffff811161046b576100919036906004016106cf565b6100996106bf565b916044359060028210156104f55760643567ffffffffffffffff81116104f1576100c7903690600401610859565b916084359367ffffffffffffffff85116104ed576100ec610178953690600401610700565b93909261017160405160208101906101678161013b3360a43586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107b2565b5190209687611568565b9788610972565b9173ffffffffffffffffffffffffffffffffffffffff821690813b1561045d576040517fc630948d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152818160248183875af18015610460576104dd575b509073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561048f578280916024604051809481937f96ea2f7a0000000000000000000000000000000000000000000000000000000083528760048401525af18015610484579083916104c8575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561048f576040517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af180156104a8579084916104b3575b5050803b1561048f576040517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015286166024820152838160448183865af180156104a857908491610493575b5050803b1561048f576040517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201523360248201529083908290604490829084905af180156104845790839161046f575b5050803b1561046b578180916024604051809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af1801561046057610448575b50506040805173ffffffffffffffffffffffffffffffffffffffff928316815292909116602083015290f35b0390f35b6104538280926107b2565b61045d5780610418565b80fd5b6040513d84823e3d90fd5b5080fd5b81610479916107b2565b61046b5781386103cd565b6040513d85823e3d90fd5b8280fd5b8161049d916107b2565b61048f578238610360565b6040513d86823e3d90fd5b816104bd916107b2565b61048f5782386102f6565b816104d2916107b2565b61046b578138610273565b816104e7916107b2565b386101ee565b8680fd5b8580fd5b8480fd5b503461045d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261045d575061044460405161053a6040826107b2565b601681527f546f6b656e506f6f6c466163746f727920312e352e31000000000000000000006020820152604051918291602083526020830190610816565b823461045d5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261045d576004359173ffffffffffffffffffffffffffffffffffffffff8316830361046b576105d16106bf565b9160443590600282101561045d5760643567ffffffffffffffff811161046b576105ff9036906004016106cf565b9290916084359067ffffffffffffffff821161045d57602061067b89898989898961062d3660048c01610700565b60a4358a87019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015290969195610672816034840161013b565b51902096610972565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b359073ffffffffffffffffffffffffffffffffffffffff821682036106ba57565b600080fd5b6024359060ff821682036106ba57565b9181601f840112156106ba5782359167ffffffffffffffff83116106ba576020808501948460051b0101116106ba57565b9181601f840112156106ba5782359167ffffffffffffffff83116106ba57602083818601950101116106ba57565b6060810190811067ffffffffffffffff82111761074a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761074a57604052565b610100810190811067ffffffffffffffff82111761074a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761074a57604052565b60005b8381106108065750506000910152565b81810151838201526020016107f6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610852815180928187528780880191016107f3565b0116010190565b81601f820112156106ba5780359067ffffffffffffffff821161074a57604051926108ac601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856107b2565b828452602083830101116106ba57816000926020809301838601378301015290565b67ffffffffffffffff811161074a5760051b60200190565b604051906108f38261072e565b60006040838281528260208201520152565b35906fffffffffffffffffffffffffffffffff821682036106ba57565b80511561092f5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561092f5760209160051b010190565b95909693979297608052600060a05261098c6080516108ce565b9561099a60405197886107b2565b60805187527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109cb6080516108ce565b0160a0515b81811061151b5750506040516109e581610795565b60a05181526060602082015260606040820152604051610a0481610779565b60a051815260a051602082015260a051604082015260a051606082015260a0516080820152606082015260a0516080820152606060a0820152606060c0820152610a4c6108e6565b60e0919091015260a051604096905b60805181101561103e57600581901b8b0135368c90037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4101811215610fd2576101c0818d01360312610fd257885190610ab382610795565b808d013567ffffffffffffffff81168103610fd25782528c81016020013567ffffffffffffffff8111610fd257818e610af28d93369084840101610859565b602086015201013567ffffffffffffffff8111610fd2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0828f8d610b3e60a095369085850101610859565b908701520136030112610fd257895160e0828f610b5a84610779565b610b68606083830101610699565b8452610b78608083830101610699565b60208501528d610b8c60a084840101610699565b90850152610b9e60c083830101610699565b60608501520101359060ff82168203610fd2578e816101009360808694015260608601520101356002811015610fd25760808301528c8101610120013567ffffffffffffffff8111610fd257818e610bfe61014093369084840101610859565b60a08601520101359067ffffffffffffffff8211610fd25760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea08f93610c4b8491369083880101610859565b9460c087019586520136030112610fd257806101608f928d5193610c6e8561072e565b0101358015158103610fd2578f91610ca1926101a0928552610c9561018083830101610905565b60208601520101610905565b8b82015260e083015260a08201515115610fd8575b5060208101515115610d86575b8851610ccf8a826107b2565b6001815260a0515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08b018110610d75575090600192916020820151610d1482610922565b52610d1e81610922565b5067ffffffffffffffff8251169160e060a0820151910151918c5193610d4385610779565b845260208401528b8301528060608301526080820152610d63828c61095e565b52610d6e818b61095e565b5001610a5b565b806060602080938501015201610cd7565b8881015160608201519060a0830151602081805181010312610fd2576020015173ffffffffffffffffffffffffffffffffffffffff81168103610fd2576080840151906002821015610fa15773ffffffffffffffffffffffffffffffffffffffff8d610ed992610eff96600160209660a0515060a0515014600014610f3557610e7e928160ff608061013b9401511686838301511690876060818c86015116940151169351978896168a87019373ffffffffffffffffffffffffffffffffffffffff92908360a09560ff82949a99959a8360c08b019c168a521660208901526000604089015216606087015216608085015216910152565b8d5192839181610e9781850197888151938492016107f3565b8301610eab825180938580850191016107f3565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107b2565b51902073ffffffffffffffffffffffffffffffffffffffff60608401515116908a611602565b73ffffffffffffffffffffffffffffffffffffffff8a519116602082015260208152610f2b8a826107b2565b6020820152610cc3565b610f9c928160ff608061013b94015116868981858501511693015116925196879516898601929373ffffffffffffffffffffffffffffffffffffffff8092969560ff6080958360a089019a1688521660208701526000604087015216606085015216910152565b610e7e565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526021600452602460a051fd5b60a05180fd5b61100790516020815191012073ffffffffffffffffffffffffffffffffffffffff60608401515116908a611602565b73ffffffffffffffffffffffffffffffffffffffff8a5191166020820152602081526110338a826107b2565b60a082015238610cb6565b50919493985091949660a051506002861015610fa157602061117493819260016111799914600014611485578a5173ffffffffffffffffffffffffffffffffffffffff92831684820190815260ff9092166020830152600060408301527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f0000000000000000000000000000000000000000000000000000000000000000831660808301527f00000000000000000000000000000000000000000000000000000000000000009290921660a082015261112390829060c00161013b565b611146878b519889968588013785019183830160a05181528151948592016107f3565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836107b2565b611568565b8251909273ffffffffffffffffffffffffffffffffffffffff8416929091602091906111a583826107b2565b60a08051825251368137843b15610fd2578351927fe8a1da17000000000000000000000000000000000000000000000000000000008452604484018560048601528251809152816064860193019060a0515b818110611467575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8483030160248501528251908183528083019281808460051b83010195019360a051915b8483106112fd57505050505050818060a05192038160a051875af180156112f1576112df575b50813b15610fd2578051917ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401528260248160a0519360a051905af19081156112d457506112be575090565b60a0516112ca916107b2565b60a051610fd25790565b513d60a051823e3d90fd5b60a0516112eb916107b2565b3861126c565b82513d60a051823e3d90fd5b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301835284875191610120810167ffffffffffffffff84511682528b83850151916101208585015282518091526101408401856101408360051b87010194019160a0515b8181106114185750505050839260c060806113978895856114089660019b01519086830390870152610816565b946113d4606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193019194939290611246565b92959680919450611454867ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08a60019699030188528851610816565b960194019101908e928b9695949261136a565b825167ffffffffffffffff16855293830193918301916001016111f7565b8a5173ffffffffffffffffffffffffffffffffffffffff92831684820190815260ff9092166020830152600060408301527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f000000000000000000000000000000000000000000000000000000000000000092909216608082015261151690829060a00161013b565b611123565b60209060409a989a5161152d81610779565b60a0518152606083820152606060408201526115476108e6565b60608201526115546108e6565b608082015282828c010152019896986109d0565b908051156115d8576020815191016000f5903d15198215166115cc5773ffffffffffffffffffffffffffffffffffffffff8216156115a257565b7fb06ebf3d0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b60559173ffffffffffffffffffffffffffffffffffffffff93600b92604051926040840152602083015281520160ff815320169056fea164736f6c634300081a000a",
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
	return common.HexToHash("0xcf2e104173e7782dc2782d45728a7c097f4abfd93ed53dbf6c39da81c1a8f33c")
}

func (_TokenPoolFactory *TokenPoolFactory) Address() common.Address {
	return _TokenPoolFactory.address
}

type TokenPoolFactoryInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error)

	WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error)

	Address() common.Address
}
