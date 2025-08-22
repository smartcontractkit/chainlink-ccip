// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_offramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type SignatureQuorumVerifierSignatureConfigArgs struct {
	ConfigDigest [32]byte
	F            uint8
	Signers      []common.Address
}

type SignatureQuorumVerifierSignatureConfigConfig struct {
	ConfigDigest [32]byte
	F            uint8
	N            uint8
}

var CommitOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"ocrConfig\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigConfig\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"n\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSignatureConfig\",\"inputs\":[{\"name\":\"ocrConfigArgs\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateReport\",\"inputs\":[{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originalState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigDigestMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0346100b757601f61142e38819003918201601f19168301916001600160401b038311848410176100bc578084926020946040528339810103126100b757516001600160a01b0381168082036100b75733156100a657600180546001600160a01b0319163317905546608052156100955760a05260405161135b90816100d3823960805181610850015260a05181610b090152f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714610de457806339f2a396146106a957806340ee29f01461032c57806379ba50971461024357806381ff7048146101b35780638da5cb5b146101615763f2fde38b1461006957600080fd5b3461015c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c5760043573ffffffffffffffffffffffffffffffffffffffff811680910361015c576100c16110f5565b33811461013257807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461015c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461015c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c576000604080516101f181610ea8565b8281528260208201520152606060405161020a81610ea8565b60ff60025491828152816003549181604060208301928286168452019360081c1683526040519485525116602084015251166040820152f35b3461015c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c5760005473ffffffffffffffffffffffffffffffffffffffff81163303610302577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461015c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c5760043567ffffffffffffffff811161015c5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261015c57604051906103a682610ea8565b8060040135825260248101359060ff8216820361015c576020830191825260448101359067ffffffffffffffff821161015c57013660238201121561015c576004810135906103f482610f92565b916104026040519384610ec4565b808352602060048185019260051b840101019136831161015c57602401905b828210610691575050506040830190815261043a6110f5565b60ff8251161561066757516040519283846020600454928381520160046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9260005b81811061064e57505061049492500385610ec4565b60005b84518110156104d357806104cc73ffffffffffffffffffffffffffffffffffffffff6104c560019489611055565b51166111ed565b5001610497565b508260005b83518110156105645773ffffffffffffffffffffffffffffffffffffffff6105008286611055565b51161561053a578061053373ffffffffffffffffffffffffffffffffffffffff61052c60019488611055565b5116611158565b50016104d8565b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b509060ff9080518284511685519160405161057e81610ea8565b81815282602082015260408685169101526002557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000061ff006003549360081b169216171760035551915116604051916060830190835260606020840152835180915260206080840194019060005b818110610622577f5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e08580888760408301520390a1005b825173ffffffffffffffffffffffffffffffffffffffff168652602095860195909201916001016105ec565b845483526001948501948994506020909301920161047f565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b6020809161069e84610faa565b815201910190610421565b3461015c5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c5760043567ffffffffffffffff811161015c576106f8903690600401610f64565b9060443567ffffffffffffffff811161015c57610719903690600401610f64565b91909260643592600484101561015c5784019360408186031261015c57803567ffffffffffffffff811161015c5785610753918301610fcb565b94602082013567ffffffffffffffff811161015c576107729201610fcb565b938051946020820195604083828101031261015c57604087519301519667ffffffffffffffff881680980361015c5720600254838103610db2575081518201602081019260208183031261015c57602081015167ffffffffffffffff811161015c5701906040908290031261015c57604051926040840184811067ffffffffffffffff821117610d8357604052602082015167ffffffffffffffff811161015c5781602061082292850101611098565b845260408201519167ffffffffffffffff831161015c576108469201602001611098565b92602083019384527f0000000000000000000000000000000000000000000000000000000000000000468103610d525750825151600160ff600354160160ff8111610d235760ff1603610cf95782515184515103610ccf57604051906020820192602435845260408301526060820152606081526108c5608082610ec4565b5190209051915182519260005b848110610bf7575050505050836108e557005b600092156108ef57005b810190602081830312610bf35780359067ffffffffffffffff8211610bc2570190818103906101208212610bc2576040519160c0830183811067ffffffffffffffff821117610bc657604052608013610bc25760405161094e81610e8c565b8335815261095e60208501611040565b602082015261096f60408501611040565b604082015261098060608501611040565b60608201528252608083013567ffffffffffffffff8111610bba57816109a7918501610fcb565b926020830193845260a081013567ffffffffffffffff8111610bbe57826109cf918301610fcb565b60408401526109e060c08201610faa565b606084015260e081013563ffffffff81168103610bbe5760808401526101008101359067ffffffffffffffff8211610bbe5701608081830312610bba5760405191610a2a83610e8c565b813567ffffffffffffffff8111610bb65781610a47918401610fcb565b8352610a5560208301610faa565b6020840152604082013567ffffffffffffffff8111610bb65793610aef9593606084610a9060209895899667ffffffffffffffff9801610fcb565b60408401520135606082015260a08201525101511690519060405193849283927fe0e03cae0000000000000000000000000000000000000000000000000000000084526004840152876024840152606060448401526064830190610f05565b03818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610bab578291610b6c575b5015610b4157005b6024917f5c33785a000000000000000000000000000000000000000000000000000000008252600452fd5b90506020813d602011610ba3575b81610b8760209383610ec4565b81010312610b9f57518015158103610b9f5783610b39565b5080fd5b3d9150610b7a565b6040513d84823e3d90fd5b8680fd5b8480fd5b8580fd5b8380fd5b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8280fd5b602060006080610c078486611055565b51610c128588611055565b5160405191898352601b868401526040830152606082015282805260015afa15610cc35773ffffffffffffffffffffffffffffffffffffffff60005116604060008281526005602052205415610c995715610c6f576001016108d2565b7ff67bc7c40000000000000000000000000000000000000000000000000000000060005260046000fd5b7fca31867a0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7fa75d88af0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b83907f93df584c0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b3461015c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015c57610e88604051610e24606082610ec4565b602181527f5369676e617475726551756f72756d566572696669657220312e372e302d646560208201527f76000000000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190610f05565b0390f35b6080810190811067ffffffffffffffff821117610d8357604052565b6060810190811067ffffffffffffffff821117610d8357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d8357604052565b919082519283825260005b848110610f4f5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610f10565b9181601f8401121561015c5782359167ffffffffffffffff831161015c576020838186019501011161015c57565b67ffffffffffffffff8111610d835760051b60200190565b359073ffffffffffffffffffffffffffffffffffffffff8216820361015c57565b81601f8201121561015c5780359067ffffffffffffffff8211610d83576040519261101e601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185610ec4565b8284526020838301011161015c57816000926020809301838601378301015290565b359067ffffffffffffffff8216820361015c57565b80518210156110695760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561015c5781516110af81610f92565b926110bd6040519485610ec4565b81845260208085019260051b82010192831161015c57602001905b8282106110e55750505090565b81518152602091820191016110d8565b73ffffffffffffffffffffffffffffffffffffffff60015416330361111657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156110695760005260206000200190600090565b806000526005602052604060002054156000146111e75760045468010000000000000000811015610d83576111ce6111998260018594016004556004611140565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600454906000526005602052604060002055600190565b50600090565b6000818152600560205260409020548015611347577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610d2357600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610d235781810361130d575b50505060045480156112de577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161129b816004611140565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61132f61131e611199936004611140565b90549060031b1c9283926004611140565b90556000526005602052604060002055388080611262565b505060009056fea164736f6c634300081a000a",
}

var CommitOffRampABI = CommitOffRampMetaData.ABI

var CommitOffRampBin = CommitOffRampMetaData.Bin

func DeployCommitOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, nonceManager common.Address) (common.Address, *types.Transaction, *CommitOffRamp, error) {
	parsed, err := CommitOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitOffRampBin), backend, nonceManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitOffRamp{address: address, abi: *parsed, CommitOffRampCaller: CommitOffRampCaller{contract: contract}, CommitOffRampTransactor: CommitOffRampTransactor{contract: contract}, CommitOffRampFilterer: CommitOffRampFilterer{contract: contract}}, nil
}

type CommitOffRamp struct {
	address common.Address
	abi     abi.ABI
	CommitOffRampCaller
	CommitOffRampTransactor
	CommitOffRampFilterer
}

type CommitOffRampCaller struct {
	contract *bind.BoundContract
}

type CommitOffRampTransactor struct {
	contract *bind.BoundContract
}

type CommitOffRampFilterer struct {
	contract *bind.BoundContract
}

type CommitOffRampSession struct {
	Contract     *CommitOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitOffRampCallerSession struct {
	Contract *CommitOffRampCaller
	CallOpts bind.CallOpts
}

type CommitOffRampTransactorSession struct {
	Contract     *CommitOffRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitOffRampRaw struct {
	Contract *CommitOffRamp
}

type CommitOffRampCallerRaw struct {
	Contract *CommitOffRampCaller
}

type CommitOffRampTransactorRaw struct {
	Contract *CommitOffRampTransactor
}

func NewCommitOffRamp(address common.Address, backend bind.ContractBackend) (*CommitOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitOffRamp{address: address, abi: abi, CommitOffRampCaller: CommitOffRampCaller{contract: contract}, CommitOffRampTransactor: CommitOffRampTransactor{contract: contract}, CommitOffRampFilterer: CommitOffRampFilterer{contract: contract}}, nil
}

func NewCommitOffRampCaller(address common.Address, caller bind.ContractCaller) (*CommitOffRampCaller, error) {
	contract, err := bindCommitOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampCaller{contract: contract}, nil
}

func NewCommitOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitOffRampTransactor, error) {
	contract, err := bindCommitOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampTransactor{contract: contract}, nil
}

func NewCommitOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitOffRampFilterer, error) {
	contract, err := bindCommitOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampFilterer{contract: contract}, nil
}

func bindCommitOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitOffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitOffRamp *CommitOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOffRamp.Contract.CommitOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitOffRamp *CommitOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.CommitOffRampTransactor.contract.Transfer(opts)
}

func (_CommitOffRamp *CommitOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.CommitOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitOffRamp *CommitOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitOffRamp *CommitOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.contract.Transfer(opts)
}

func (_CommitOffRamp *CommitOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitOffRamp *CommitOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (SignatureQuorumVerifierSignatureConfigConfig, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	if err != nil {
		return *new(SignatureQuorumVerifierSignatureConfigConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(SignatureQuorumVerifierSignatureConfigConfig)).(*SignatureQuorumVerifierSignatureConfigConfig)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) LatestConfigDetails() (SignatureQuorumVerifierSignatureConfigConfig, error) {
	return _CommitOffRamp.Contract.LatestConfigDetails(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) LatestConfigDetails() (SignatureQuorumVerifierSignatureConfigConfig, error) {
	return _CommitOffRamp.Contract.LatestConfigDetails(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) Owner() (common.Address, error) {
	return _CommitOffRamp.Contract.Owner(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) Owner() (common.Address, error) {
	return _CommitOffRamp.Contract.Owner(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) TypeAndVersion() (string, error) {
	return _CommitOffRamp.Contract.TypeAndVersion(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitOffRamp.Contract.TypeAndVersion(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitOffRamp *CommitOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOffRamp.Contract.AcceptOwnership(&_CommitOffRamp.TransactOpts)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOffRamp.Contract.AcceptOwnership(&_CommitOffRamp.TransactOpts)
}

func (_CommitOffRamp *CommitOffRampTransactor) SetSignatureConfig(opts *bind.TransactOpts, ocrConfigArgs SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "setSignatureConfig", ocrConfigArgs)
}

func (_CommitOffRamp *CommitOffRampSession) SetSignatureConfig(ocrConfigArgs SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfig(&_CommitOffRamp.TransactOpts, ocrConfigArgs)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) SetSignatureConfig(ocrConfigArgs SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfig(&_CommitOffRamp.TransactOpts, ocrConfigArgs)
}

func (_CommitOffRamp *CommitOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitOffRamp *CommitOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.TransferOwnership(&_CommitOffRamp.TransactOpts, to)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.TransferOwnership(&_CommitOffRamp.TransactOpts, to)
}

func (_CommitOffRamp *CommitOffRampTransactor) ValidateReport(opts *bind.TransactOpts, rawMessage []byte, messageHash [32]byte, proof []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "validateReport", rawMessage, messageHash, proof, originalState)
}

func (_CommitOffRamp *CommitOffRampSession) ValidateReport(rawMessage []byte, messageHash [32]byte, proof []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, rawMessage, messageHash, proof, originalState)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) ValidateReport(rawMessage []byte, messageHash [32]byte, proof []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, rawMessage, messageHash, proof, originalState)
}

type CommitOffRampConfigSetIterator struct {
	Event *CommitOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOffRampConfigSet)
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
		it.Event = new(CommitOffRampConfigSet)
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

func (it *CommitOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOffRampConfigSet struct {
	ConfigDigest [32]byte
	Signers      []common.Address
	F            uint8
	Raw          types.Log
}

func (_CommitOffRamp *CommitOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitOffRampConfigSetIterator, error) {

	logs, sub, err := _CommitOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitOffRampConfigSetIterator{contract: _CommitOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOffRamp *CommitOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOffRampConfigSet)
				if err := _CommitOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitOffRamp *CommitOffRampFilterer) ParseConfigSet(log types.Log) (*CommitOffRampConfigSet, error) {
	event := new(CommitOffRampConfigSet)
	if err := _CommitOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOffRampOwnershipTransferRequestedIterator struct {
	Event *CommitOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOffRampOwnershipTransferRequested)
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
		it.Event = new(CommitOffRampOwnershipTransferRequested)
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

func (it *CommitOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOffRamp *CommitOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampOwnershipTransferRequestedIterator{contract: _CommitOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitOffRamp *CommitOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOffRampOwnershipTransferRequested)
				if err := _CommitOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitOffRamp *CommitOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitOffRampOwnershipTransferRequested, error) {
	event := new(CommitOffRampOwnershipTransferRequested)
	if err := _CommitOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOffRampOwnershipTransferredIterator struct {
	Event *CommitOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOffRampOwnershipTransferred)
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
		it.Event = new(CommitOffRampOwnershipTransferred)
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

func (it *CommitOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOffRamp *CommitOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampOwnershipTransferredIterator{contract: _CommitOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitOffRamp *CommitOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOffRampOwnershipTransferred)
				if err := _CommitOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitOffRamp *CommitOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitOffRampOwnershipTransferred, error) {
	event := new(CommitOffRampOwnershipTransferred)
	if err := _CommitOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CommitOffRamp *CommitOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CommitOffRamp.abi.Events["ConfigSet"].ID:
		return _CommitOffRamp.ParseConfigSet(log)
	case _CommitOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _CommitOffRamp.ParseOwnershipTransferRequested(log)
	case _CommitOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _CommitOffRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CommitOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e0")
}

func (CommitOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CommitOffRamp *CommitOffRamp) Address() common.Address {
	return _CommitOffRamp.address
}

type CommitOffRampInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (SignatureQuorumVerifierSignatureConfigConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetSignatureConfig(opts *bind.TransactOpts, ocrConfigArgs SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	ValidateReport(opts *bind.TransactOpts, rawMessage []byte, messageHash [32]byte, proof []byte, originalState uint8) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitOffRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitOffRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
