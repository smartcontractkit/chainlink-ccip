// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_verifier_offramp

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

type OCRVerifierOCRConfig struct {
	ConfigDigest [32]byte
	F            uint8
	N            uint8
}

type OCRVerifierOCRConfigArgs struct {
	ConfigDigest [32]byte
	F            uint8
	Signers      []common.Address
}

var CommitVerifierOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"ocrConfig\",\"type\":\"tuple\",\"internalType\":\"structOCRVerifier.OCRConfig\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"n\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOCR3Config\",\"inputs\":[{\"name\":\"ocrConfigArgs\",\"type\":\"tuple\",\"internalType\":\"structOCRVerifier.OCRConfigArgs\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateReport\",\"inputs\":[{\"name\":\"rawReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ocrProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transmitted\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigDigestMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[{\"name\":\"errorType\",\"type\":\"uint8\",\"internalType\":\"enumOCRVerifier.InvalidConfigErrorType\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0346100b757601f61164e38819003918201601f19168301916001600160401b038311848410176100bc578084926020946040528339810103126100b757516001600160a01b0381168082036100b75733156100a657600180546001600160a01b0319163317905546608052156100955760a05260405161157b90816100d38239608051816106c2015260a05181610a480152f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461101d57806379ba509714610f3457806381ff704814610ea45780638da5cb5b14610e52578063cba4c71a14610595578063f2fde38b146104a25763f300ce6d1461006957600080fd5b3461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760043567ffffffffffffffff811161049d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261049d57604051906100e38261109e565b8060040135825260248101359060ff8216820361049d576020830191825260448101359067ffffffffffffffff821161049d57013660238201121561049d57600481013590610131826111c5565b9161013f60405193846110d6565b808352602060048185019260051b840101019136831161049d57602401905b8282106104855750505060408301908152610177611315565b60ff825116156104565751604081511161042757805160ff8351166003029060ff82169182036103f85711156103c9576040519283846020600454928381520160046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9260005b8181106103b05750506101f6925003856110d6565b60005b8451811015610235578061022e73ffffffffffffffffffffffffffffffffffffffff61022760019489611275565b511661140d565b50016101f9565b508260005b83518110156102c65773ffffffffffffffffffffffffffffffffffffffff6102628286611275565b51161561029c578061029573ffffffffffffffffffffffffffffffffffffffff61028e60019488611275565b5116611378565b500161023a565b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b509060ff908051828451168551916040516102e08161109e565b81815282602082015260408685169101526002557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000061ff006003549360081b169216171760035551915116604051916060830190835260606020840152835180915260206080840194019060005b818110610384577f5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e08580888760408301520390a1005b825173ffffffffffffffffffffffffffffffffffffffff1686526020958601959092019160010161034e565b84548352600194850194899450602090930192016101e1565b7f367f56a200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f367f56a200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f367f56a200000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b60208091610492846111a4565b81520191019061015e565b600080fd5b3461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760043573ffffffffffffffffffffffffffffffffffffffff811680910361049d576104fa611315565b33811461056b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461049d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760043567ffffffffffffffff811161049d576105e4903690600401611176565b60243567ffffffffffffffff811161049d57610604903690600401611176565b919060643592600484101561049d5781019060208183031261049d5780359067ffffffffffffffff821161049d57019060808282031261049d576040519061064b826110ba565b8235825261065b602084016111dd565b9260208301938452604081013567ffffffffffffffff811161049d57826106839183016112b8565b9160408401928352606082013567ffffffffffffffff811161049d576106a992016112b8565b9060608301918252600254835190818103610e225750507f0000000000000000000000000000000000000000000000000000000000000000468103610df15750805151600160ff600354160160ff81116103f85760ff1603610dc75780515182515103610d9d5761071b3686896111f2565b60208151910120835167ffffffffffffffff865116604051916020830193845260408301526060820152606081526107546080826110d6565b5190209051915182519260005b848110610cc5578989897fe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea260408b67ffffffffffffffff8c5191511682519182526020820152a182019160208184031261049d5780359067ffffffffffffffff821161049d570180830392610140841261049d576040519360e085019085821067ffffffffffffffff831117610bd4576080916040521261049d57604051610808816110ba565b82358152610818602084016111dd565b6020820152610829604084016111dd565b604082015261083a606084016111dd565b60608201528452608082013567ffffffffffffffff811161049d5781610861918401611257565b916020850192835260a081013567ffffffffffffffff811161049d5782610889918301611257565b604086015261089a60c082016111a4565b606086015260e081013563ffffffff8116810361049d57608086015261010081013567ffffffffffffffff811161049d57810182601f8201121561049d5780356108e3816111c5565b916108f160405193846110d6565b81835260208084019260051b8201019085821161049d5760208101925b828410610c03575050505060a08601526101208101359067ffffffffffffffff821161049d57019080601f8301121561049d5781359161094d836111c5565b9261095b60405194856110d6565b80845260208085019160051b8301019183831161049d5760208101915b838310610af55788888860206109978a8060c087015260443590611275565b51015160208180518101031261049d57602001519267ffffffffffffffff841680940361049d57836109c557005b600092156109cf57005b610a2e9167ffffffffffffffff602080935101511690519060405193849283927fe0e03cae0000000000000000000000000000000000000000000000000000000084526004840152876024840152606060448401526064830190611117565b03818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610aea578291610aab575b5015610a8057005b6024917f5c33785a000000000000000000000000000000000000000000000000000000008252600452fd5b90506020813d602011610ae2575b81610ac6602093836110d6565b81010312610ade57518015158103610ade5783610a78565b5080fd5b3d9150610ab9565b6040513d84823e3d90fd5b823567ffffffffffffffff811161049d5782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261049d576040519060a0820182811067ffffffffffffffff821117610bd45760405260208301358252604083013567ffffffffffffffff811161049d57876020610b7d92860101611257565b602083015260608301356040830152610b98608084016111dd565b606083015260a08301359167ffffffffffffffff831161049d57610bc488602080969581960101611257565b6080820152815201920191610978565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b833567ffffffffffffffff811161049d57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828a03011261049d5760405191610c4f836110ba565b602082013567ffffffffffffffff811161049d57896020610c7292850101611257565b8352610c80604083016111a4565b602084015260608201359267ffffffffffffffff841161049d57608083610cae8c6020809881980101611257565b60408401520135606082015281520193019261090e565b602060006080610cd58486611275565b51610ce08588611275565b5160405191898352601b868401526040830152606082015282805260015afa15610d915773ffffffffffffffffffffffffffffffffffffffff60005116604060008281526005602052205415610d675715610d3d57600101610761565b7ff67bc7c40000000000000000000000000000000000000000000000000000000060005260046000fd5b7fca31867a0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7fa75d88af0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f93df584c0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b3461049d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461049d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57600060408051610ee28161109e565b82815282602082015201526060604051610efb8161109e565b60ff60025491828152816003549181604060208301928286168452019360081c1683526040519485525116602084015251166040820152f35b3461049d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760005473ffffffffffffffffffffffffffffffffffffffff81163303610ff3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461049d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5761109a604080519061105e81836110d6565b601582527f4f4352566572696669657220312e372e302d6465760000000000000000000000602083015251918291602083526020830190611117565b0390f35b6060810190811067ffffffffffffffff821117610bd457604052565b6080810190811067ffffffffffffffff821117610bd457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610bd457604052565b919082519283825260005b8481106111615750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611122565b9181601f8401121561049d5782359167ffffffffffffffff831161049d576020838186019501011161049d57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361049d57565b67ffffffffffffffff8111610bd45760051b60200190565b359067ffffffffffffffff8216820361049d57565b92919267ffffffffffffffff8211610bd4576040519161123a601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001846110d6565b82948184528183011161049d578281602093846000960137010152565b9080601f8301121561049d57816020611272933591016111f2565b90565b80518210156112895760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f8301121561049d5781356112cf816111c5565b926112dd60405194856110d6565b81845260208085019260051b82010192831161049d57602001905b8282106113055750505090565b81358152602091820191016112f8565b73ffffffffffffffffffffffffffffffffffffffff60015416330361133657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156112895760005260206000200190600090565b806000526005602052604060002054156000146114075760045468010000000000000000811015610bd4576113ee6113b98260018594016004556004611360565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600454906000526005602052604060002055600190565b50600090565b6000818152600560205260409020548015611567577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116103f857600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116103f85781810361152d575b50505060045480156114fe577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114bb816004611360565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61154f61153e6113b9936004611360565b90549060031b1c9283926004611360565b90556000526005602052604060002055388080611482565b505060009056fea164736f6c634300081a000a",
}

var CommitVerifierOffRampABI = CommitVerifierOffRampMetaData.ABI

var CommitVerifierOffRampBin = CommitVerifierOffRampMetaData.Bin

func DeployCommitVerifierOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, nonceManager common.Address) (common.Address, *types.Transaction, *CommitVerifierOffRamp, error) {
	parsed, err := CommitVerifierOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitVerifierOffRampBin), backend, nonceManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitVerifierOffRamp{address: address, abi: *parsed, CommitVerifierOffRampCaller: CommitVerifierOffRampCaller{contract: contract}, CommitVerifierOffRampTransactor: CommitVerifierOffRampTransactor{contract: contract}, CommitVerifierOffRampFilterer: CommitVerifierOffRampFilterer{contract: contract}}, nil
}

type CommitVerifierOffRamp struct {
	address common.Address
	abi     abi.ABI
	CommitVerifierOffRampCaller
	CommitVerifierOffRampTransactor
	CommitVerifierOffRampFilterer
}

type CommitVerifierOffRampCaller struct {
	contract *bind.BoundContract
}

type CommitVerifierOffRampTransactor struct {
	contract *bind.BoundContract
}

type CommitVerifierOffRampFilterer struct {
	contract *bind.BoundContract
}

type CommitVerifierOffRampSession struct {
	Contract     *CommitVerifierOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitVerifierOffRampCallerSession struct {
	Contract *CommitVerifierOffRampCaller
	CallOpts bind.CallOpts
}

type CommitVerifierOffRampTransactorSession struct {
	Contract     *CommitVerifierOffRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitVerifierOffRampRaw struct {
	Contract *CommitVerifierOffRamp
}

type CommitVerifierOffRampCallerRaw struct {
	Contract *CommitVerifierOffRampCaller
}

type CommitVerifierOffRampTransactorRaw struct {
	Contract *CommitVerifierOffRampTransactor
}

func NewCommitVerifierOffRamp(address common.Address, backend bind.ContractBackend) (*CommitVerifierOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitVerifierOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitVerifierOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRamp{address: address, abi: abi, CommitVerifierOffRampCaller: CommitVerifierOffRampCaller{contract: contract}, CommitVerifierOffRampTransactor: CommitVerifierOffRampTransactor{contract: contract}, CommitVerifierOffRampFilterer: CommitVerifierOffRampFilterer{contract: contract}}, nil
}

func NewCommitVerifierOffRampCaller(address common.Address, caller bind.ContractCaller) (*CommitVerifierOffRampCaller, error) {
	contract, err := bindCommitVerifierOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampCaller{contract: contract}, nil
}

func NewCommitVerifierOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitVerifierOffRampTransactor, error) {
	contract, err := bindCommitVerifierOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampTransactor{contract: contract}, nil
}

func NewCommitVerifierOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitVerifierOffRampFilterer, error) {
	contract, err := bindCommitVerifierOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampFilterer{contract: contract}, nil
}

func bindCommitVerifierOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitVerifierOffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitVerifierOffRamp.Contract.CommitVerifierOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.CommitVerifierOffRampTransactor.contract.Transfer(opts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.CommitVerifierOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitVerifierOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.contract.Transfer(opts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (OCRVerifierOCRConfig, error) {
	var out []interface{}
	err := _CommitVerifierOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	if err != nil {
		return *new(OCRVerifierOCRConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OCRVerifierOCRConfig)).(*OCRVerifierOCRConfig)

	return out0, err

}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) LatestConfigDetails() (OCRVerifierOCRConfig, error) {
	return _CommitVerifierOffRamp.Contract.LatestConfigDetails(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCallerSession) LatestConfigDetails() (OCRVerifierOCRConfig, error) {
	return _CommitVerifierOffRamp.Contract.LatestConfigDetails(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitVerifierOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) Owner() (common.Address, error) {
	return _CommitVerifierOffRamp.Contract.Owner(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCallerSession) Owner() (common.Address, error) {
	return _CommitVerifierOffRamp.Contract.Owner(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitVerifierOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) TypeAndVersion() (string, error) {
	return _CommitVerifierOffRamp.Contract.TypeAndVersion(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitVerifierOffRamp.Contract.TypeAndVersion(&_CommitVerifierOffRamp.CallOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.AcceptOwnership(&_CommitVerifierOffRamp.TransactOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.AcceptOwnership(&_CommitVerifierOffRamp.TransactOpts)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactor) SetOCR3Config(opts *bind.TransactOpts, ocrConfigArgs OCRVerifierOCRConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.contract.Transact(opts, "setOCR3Config", ocrConfigArgs)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) SetOCR3Config(ocrConfigArgs OCRVerifierOCRConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.SetOCR3Config(&_CommitVerifierOffRamp.TransactOpts, ocrConfigArgs)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorSession) SetOCR3Config(ocrConfigArgs OCRVerifierOCRConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.SetOCR3Config(&_CommitVerifierOffRamp.TransactOpts, ocrConfigArgs)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.TransferOwnership(&_CommitVerifierOffRamp.TransactOpts, to)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.TransferOwnership(&_CommitVerifierOffRamp.TransactOpts, to)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactor) ValidateReport(opts *bind.TransactOpts, rawReport []byte, ocrProof []byte, verifierIndex *big.Int, originalState uint8) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.contract.Transact(opts, "validateReport", rawReport, ocrProof, verifierIndex, originalState)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampSession) ValidateReport(rawReport []byte, ocrProof []byte, verifierIndex *big.Int, originalState uint8) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.ValidateReport(&_CommitVerifierOffRamp.TransactOpts, rawReport, ocrProof, verifierIndex, originalState)
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampTransactorSession) ValidateReport(rawReport []byte, ocrProof []byte, verifierIndex *big.Int, originalState uint8) (*types.Transaction, error) {
	return _CommitVerifierOffRamp.Contract.ValidateReport(&_CommitVerifierOffRamp.TransactOpts, rawReport, ocrProof, verifierIndex, originalState)
}

type CommitVerifierOffRampConfigSetIterator struct {
	Event *CommitVerifierOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOffRampConfigSet)
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
		it.Event = new(CommitVerifierOffRampConfigSet)
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

func (it *CommitVerifierOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOffRampConfigSet struct {
	ConfigDigest [32]byte
	Signers      []common.Address
	F            uint8
	Raw          types.Log
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitVerifierOffRampConfigSetIterator, error) {

	logs, sub, err := _CommitVerifierOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampConfigSetIterator{contract: _CommitVerifierOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitVerifierOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOffRampConfigSet)
				if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) ParseConfigSet(log types.Log) (*CommitVerifierOffRampConfigSet, error) {
	event := new(CommitVerifierOffRampConfigSet)
	if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOffRampOwnershipTransferRequestedIterator struct {
	Event *CommitVerifierOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOffRampOwnershipTransferRequested)
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
		it.Event = new(CommitVerifierOffRampOwnershipTransferRequested)
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

func (it *CommitVerifierOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampOwnershipTransferRequestedIterator{contract: _CommitVerifierOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOffRampOwnershipTransferRequested)
				if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitVerifierOffRampOwnershipTransferRequested, error) {
	event := new(CommitVerifierOffRampOwnershipTransferRequested)
	if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOffRampOwnershipTransferredIterator struct {
	Event *CommitVerifierOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOffRampOwnershipTransferred)
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
		it.Event = new(CommitVerifierOffRampOwnershipTransferred)
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

func (it *CommitVerifierOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampOwnershipTransferredIterator{contract: _CommitVerifierOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOffRampOwnershipTransferred)
				if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitVerifierOffRampOwnershipTransferred, error) {
	event := new(CommitVerifierOffRampOwnershipTransferred)
	if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOffRampTransmittedIterator struct {
	Event *CommitVerifierOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOffRampTransmitted)
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
		it.Event = new(CommitVerifierOffRampTransmitted)
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

func (it *CommitVerifierOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOffRampTransmitted struct {
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*CommitVerifierOffRampTransmittedIterator, error) {

	logs, sub, err := _CommitVerifierOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOffRampTransmittedIterator{contract: _CommitVerifierOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _CommitVerifierOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOffRampTransmitted)
				if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_CommitVerifierOffRamp *CommitVerifierOffRampFilterer) ParseTransmitted(log types.Log) (*CommitVerifierOffRampTransmitted, error) {
	event := new(CommitVerifierOffRampTransmitted)
	if err := _CommitVerifierOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CommitVerifierOffRamp *CommitVerifierOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CommitVerifierOffRamp.abi.Events["ConfigSet"].ID:
		return _CommitVerifierOffRamp.ParseConfigSet(log)
	case _CommitVerifierOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _CommitVerifierOffRamp.ParseOwnershipTransferRequested(log)
	case _CommitVerifierOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _CommitVerifierOffRamp.ParseOwnershipTransferred(log)
	case _CommitVerifierOffRamp.abi.Events["Transmitted"].ID:
		return _CommitVerifierOffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CommitVerifierOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e0")
}

func (CommitVerifierOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitVerifierOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CommitVerifierOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2")
}

func (_CommitVerifierOffRamp *CommitVerifierOffRamp) Address() common.Address {
	return _CommitVerifierOffRamp.address
}

type CommitVerifierOffRampInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (OCRVerifierOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOCR3Config(opts *bind.TransactOpts, ocrConfigArgs OCRVerifierOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	ValidateReport(opts *bind.TransactOpts, rawReport []byte, ocrProof []byte, verifierIndex *big.Int, originalState uint8) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitVerifierOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitVerifierOffRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitVerifierOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitVerifierOffRampOwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*CommitVerifierOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *CommitVerifierOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*CommitVerifierOffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
