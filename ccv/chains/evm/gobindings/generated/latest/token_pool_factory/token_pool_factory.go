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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contract RegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RemoteChainConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101203461016357601f61183c38819003918201601f19168301916001600160401b038311848410176101685780849260a0946040528339810103126101635780516001600160a01b03811691908281036101635760208201516001600160a01b03811691828203610163576100776040850161017e565b9261009060806100896060880161017e565b960161017e565b951590811561015a575b508015610149575b8015610138575b8015610127575b6101165760805260a05260c05260e052610100526040516116a990816101938239608051816103ad015260a05181610328015260c05181818161122e0152611602015260e05181818161120c01526115e0015261010051818181610fbd01526111ea0152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b038516156100b0565b506001600160a01b038416156100a9565b506001600160a01b038316156100a2565b9050153861009a565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101635756fe60c0604052600436101561001257600080fd5b6000803560e01c8063181f5a7714610613578063a32fd8971461018e5763ac9883451461003e57600080fd5b346101875760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610187576004359073ffffffffffffffffffffffffffffffffffffffff821682036101875761009661080f565b6044359160028310156101875760643567ffffffffffffffff811161018a576100c39036906004016107d9565b916084359067ffffffffffffffff821161018757602061016988888888886100ee3660048b01610894565b60405160a4358a82019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015291969295919061016081603484015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610732565b51902096610af4565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b80fd5b5080fd5b50346101875760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101875760043567ffffffffffffffff811161018a576101de9036906004016107d9565b6101e661080f565b9160443590600282101561060f5760643567ffffffffffffffff811161060b5761021490369060040161081f565b916084359367ffffffffffffffff851161060757610239610299953690600401610894565b9390926102926040516020810190610288816101343360a43586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b51902096876108e3565b9788610af4565b9173ffffffffffffffffffffffffffffffffffffffff821690813b15610187576040517fc630948d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152818160248183875af1801561057e576105f7575b509073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156105a9578280916024604051809481937f96ea2f7a0000000000000000000000000000000000000000000000000000000083528760048401525af1801561059e579083916105e2575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156105a9576040517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af180156105c2579084916105cd575b5050803b156105a9576040517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015286166024820152838160448183865af180156105c2579084916105ad575b5050803b156105a9576040517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201523360248201529083908290604490829084905af1801561059e57908391610589575b5050803b1561018a578180916024604051809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af1801561057e57610569575b50506040805173ffffffffffffffffffffffffffffffffffffffff928316815292909116602083015290f35b0390f35b610574828092610732565b6101875780610539565b6040513d84823e3d90fd5b8161059391610732565b61018a5781386104ee565b6040513d85823e3d90fd5b8280fd5b816105b791610732565b6105a9578238610481565b6040513d86823e3d90fd5b816105d791610732565b6105a9578238610417565b816105ec91610732565b61018a578138610394565b8161060191610732565b3861030f565b8680fd5b8580fd5b8480fd5b503461018757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101875750610565604051610654604082610732565b601681527f546f6b656e506f6f6c466163746f727920312e352e31000000000000000000006020820152604051918291602083526020830190610796565b6060810190811067ffffffffffffffff8211176106ae57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176106ae57604052565b610100810190811067ffffffffffffffff8211176106ae57604052565b6080810190811067ffffffffffffffff8211176106ae57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176106ae57604052565b60005b8381106107865750506000910152565b8181015183820152602001610776565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936107d281518092818752878088019101610773565b0116010190565b9181601f8401121561080a5782359167ffffffffffffffff831161080a576020808501948460051b01011161080a57565b600080fd5b6024359060ff8216820361080a57565b81601f8201121561080a5780359067ffffffffffffffff82116106ae5760405192610872601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185610732565b8284526020838301011161080a57816000926020809301838601378301015290565b9181601f8401121561080a5782359167ffffffffffffffff831161080a576020838186019501011161080a57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361080a57565b9080511561093d576020815191016000f59073ffffffffffffffffffffffffffffffffffffffff82161561091357565b7f741752c20000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116106ae5760051b60200190565b6040519061098c82610692565b60006040838281528260208201520152565b35906fffffffffffffffffffffffffffffffff8216820361080a57565b8051156109c85760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156109c85760209160051b010190565b906020808351928381520192019060005b818110610a295750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610a1c565b93959490610a9460809460ff73ffffffffffffffffffffffffffffffffffffffff9586809516895216602088015260a0604088015260a0870190610a0b565b9616606085015216910152565b9473ffffffffffffffffffffffffffffffffffffffff610ae1819560ff60a0989b9a96848097168b521660208a015260c060408a015260c0890190610a0b565b9816606087015216608085015216910152565b9693979297959095608052600060a052610b0f608051610967565b95610b1d6040519788610732565b60805187527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610b4e608051610967565b0160a0515b81811061162f575050604051610b68816106f9565b60a05181526060602082015260606040820152604051610b8781610716565b60a051815260a051602082015260a051604082015260a0516060820152606082015260a0516080820152606060a0820152606060c0820152610bc761097f565b60e0919091015260a051604096905b60805181101561119857600581901b8b0135368c90037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe610181121561112c576101a0818d0136031261112c57885190610c2e826106f9565b808d013567ffffffffffffffff8116810361112c5782528c81016020013567ffffffffffffffff811161112c57818e610c6d8d9336908484010161081f565b602086015201013567ffffffffffffffff811161112c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0828f8d610cb960809536908585010161081f565b90870152013603011261112c57895160c0828f610cd584610716565b610ce36060838301016108c2565b8452610cf36080838301016108c2565b60208501528d610d0760a0848401016108c2565b908501520101359060ff8216820361112c578e8160e0936060869401526060860152010135600281101561112c5760808301528c8101610100013567ffffffffffffffff811161112c57818e610d656101209336908484010161081f565b60a08601520101359067ffffffffffffffff821161112c5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08f93610db2849136908388010161081f565b9460c08701958652013603011261112c57806101408f928d5193610dd585610692565b010135801515810361112c578f91610e0892610180928552610dfc6101608383010161099e565b6020860152010161099e565b8b82015260e083015260a08201515115611132575b5060208101515115610eed575b8851610e368a82610732565b6001815260a0515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08b018110610edc575090600192916020820151610e7b826109bb565b52610e85816109bb565b5067ffffffffffffffff8251169160e060a0820151910151918c5193610eaa856106dd565b845260208401528b8301528060608301526080820152610eca828c6109f7565b52610ed5818b6109f7565b5001610bd6565b806060602080938501015201610e3e565b8881015160608201519060a083015160208180518101031261112c57602001519173ffffffffffffffffffffffffffffffffffffffff8316830361112c5760808401519260028410156110fb5761106793602092611041928f9260609060010361109d57506060820151835160ff90911693610fe69387936101349390929190610f778682610732565b60a0805182525136813773ffffffffffffffffffffffffffffffffffffffff8681858501511693015116925197889673ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000009616908801610aa1565b8d5192839181610fff8185019788815193849201610773565b830161101382518093858085019101610773565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610732565b51902073ffffffffffffffffffffffffffffffffffffffff60608401515116908a61167c565b73ffffffffffffffffffffffffffffffffffffffff8a5191166020820152602081526110938a82610732565b6020820152610e2a565b820151835173ffffffffffffffffffffffffffffffffffffffff946110f694610134939192909160ff16896110d28184610732565b60a08051845251368137888181878701511695015116945198899716908701610a55565b610fe6565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526021600452602460a051fd5b60a05180fd5b61116190516020815191012073ffffffffffffffffffffffffffffffffffffffff60608401515116908a61167c565b73ffffffffffffffffffffffffffffffffffffffff8a51911660208201526020815261118d8a82610732565b60a082015238610e1d565b50919493985091949660a0515060028610156110fb5760206112a793819260016112ac99146000146115b8578a516112569190610134906111d98682610732565b60a080518252513681378d519485937f0000000000000000000000000000000000000000000000000000000000000000927f0000000000000000000000000000000000000000000000000000000000000000927f0000000000000000000000000000000000000000000000000000000000000000928a8801610aa1565b611279878b519889968588013785019183830160a0518152815194859201610773565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610732565b6108e3565b8251909273ffffffffffffffffffffffffffffffffffffffff8416929091602091906112d88382610732565b60a08051825251368137843b1561112c578351927fe8a1da17000000000000000000000000000000000000000000000000000000008452604484018560048601528251809152816064860193019060a0515b81811061159a575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8483030160248501528251908183528083019281808460051b83010195019360a051915b84831061143057505050505050818060a05192038160a051875af1801561142457611412575b50813b1561112c578051917ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401528260248160a0519360a051905af190811561140757506113f1575090565b60a0516113fd91610732565b60a05161112c5790565b513d60a051823e3d90fd5b60a05161141e91610732565b3861139f565b82513d60a051823e3d90fd5b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301835284875191610120810167ffffffffffffffff84511682528b83850151916101208585015282518091526101408401856101408360051b87010194019160a0515b81811061154b5750505050839260c060806114ca88958561153b9660019b01519086830390870152610796565b94611507606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193019194939290611379565b92959680919450611587867ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08a60019699030188528851610796565b960194019101908e928b9695949261149d565b825167ffffffffffffffff168552938301939183019160010161132a565b8a51909161162a9190610134906115cf8684610732565b60a080518452513681378d519485937f0000000000000000000000000000000000000000000000000000000000000000927f000000000000000000000000000000000000000000000000000000000000000092898701610a55565b611256565b60209060409a989a51611641816106dd565b60a05181526060838201526060604082015261165b61097f565b606082015261166861097f565b608082015282828c01015201989698610b53565b90605592600b92604051926040840152602083015281520160ff8153209056fea164736f6c634300081a000a",
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
	return common.HexToHash("0xe3606343290c8e3853a7f92686979dfe24a0f29594186b61177236cb145126f0")
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
