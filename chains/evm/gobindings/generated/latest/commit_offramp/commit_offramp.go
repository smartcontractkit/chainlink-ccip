// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_offramp

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

type InternalMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	OnRampAddress       []byte
	OffRampAddress      []byte
	Finality            uint16
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []InternalTokenTransferV1
	Data                []byte
}

type InternalTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	ExtraData          []byte
}

type SignatureQuorumVerifierSignatureConfigArgs struct {
	F       uint8
	Signers []common.Address
}

var CommitOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs\",\"components\":[{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSignatureConfig\",\"inputs\":[{\"name\":\"signatureConfig\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs\",\"components\":[{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateReport\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originalState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0346100b757601f6113ce38819003918201601f19168301916001600160401b038311848410176100bc578084926020946040528339810103126100b757516001600160a01b0381168082036100b75733156100a657600180546001600160a01b0319163317905546608052156100955760a0526040516112fb90816100d3823960805181610373015260a051816105070152f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714610d775750806310803bd214610aad578063181f5a77146109fb5780636ed0e217146108c757806379ba5097146107de5780638da5cb5b1461078c578063a0c3b7891461016f5763f2fde38b1461007757600080fd5b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043573ffffffffffffffffffffffffffffffffffffffff811680910361016a576100cf611019565b33811461014057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461016a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a57803603906101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83011261016a576044359067ffffffffffffffff821161016a573660238301121561016a57816004013567ffffffffffffffff811161016a57820191602483019036821161016a5760643593600485101561016a576040908290031261016a57602481013567ffffffffffffffff811161016a5782602461025b92840101610fa4565b9160448201359167ffffffffffffffff831161016a5761027e9201602401610fa4565b90604051916020830160243581528251936102d660408260208701976102a7818484018b610ebf565b810103017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610e7e565b519020600354156107625781518201602081019260208183031261016a57602081015167ffffffffffffffff811161016a5701906040908290031261016a576040519261032284610e33565b602082015167ffffffffffffffff811161016a5781602061034592850101611064565b845260408201519167ffffffffffffffff831161016a576103699201602001611064565b90602083019182527f0000000000000000000000000000000000000000000000000000000000000000468103610731575082515191600160ff600254160160ff81116107025760ff1683036106d85780515183036106ae57600093845b8486106105b85750505050505080602091518101031261016a57519267ffffffffffffffff841680940361016a57836103fb57005b6000921561040557005b81600401359167ffffffffffffffff83168093036105b4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60c482013592018212156105b457019060048201359167ffffffffffffffff83116105b4576024019082360382136105b4579060846020927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8660405197889687957fe0e03cae00000000000000000000000000000000000000000000000000000000875260048701528b6024870152606060448701528160648701528686013788858286010152011681010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156105a957829161056a575b501561053f57005b6024917f5c33785a000000000000000000000000000000000000000000000000000000008252600452fd5b90506020813d6020116105a1575b8161058560209383610e7e565b8101031261059d5751801515810361059d5783610537565b5080fd5b3d9150610578565b6040513d84823e3d90fd5b8380fd5b6020600060806105c98986516110cc565b516105d58a88516110cc565b5160405191898352601b868401526040830152606082015282805260015afa156106a25773ffffffffffffffffffffffffffffffffffffffff60005116906040600083815260046020522054156106785773ffffffffffffffffffffffffffffffffffffffff1681111561064e576001909501946103c6565b7ff67bc7c40000000000000000000000000000000000000000000000000000000060005260046000fd5b7fca31867a0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7fa75d88af0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760005473ffffffffffffffffffffffffffffffffffffffff8116330361089d577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a576060602060405161090681610e33565b60008152015260ff6002541660405180602060035491828152019060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b8181106109e55750505081610961910382610e7e565b6040519161096e83610e33565b8252602082019081526040519182916020835260ff6060840192511660208401525190604080840152815180915260206080840192019060005b8181106109b6575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016109a8565b825484526020909301926001928301920161094b565b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57610aa460408051610a3b8282610e7e565b601781527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602083017f436f6d6d69744f666652616d7020312e372e302d646576000000000000000000815284519586946020865251809281602088015287870190610ebf565b01168101030190f35b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a578060040160407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261016a57610b26611019565b60ff610b3182610ee2565b16158015610d54575b610762575b60035415610bc957600060035415610b9c57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b54610b969073ffffffffffffffffffffffffffffffffffffffff166110f8565b50610b3f565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b9060240160005b610bda8284610ef0565b9050811015610c805773ffffffffffffffffffffffffffffffffffffffff610c14610c0f83610c098688610ef0565b90610f44565b610f83565b1615610c5657610c4973ffffffffffffffffffffffffffffffffffffffff610c43610c0f84610c098789610ef0565b1661128e565b1561076257600101610bd0565b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50610cc3610ccb9160ff610c9385610ee2565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00600254161760025583610ef0565b919092610ee2565b9160405191806040840160408552526060830191906000905b808210610d1c5760ff861660208601527f65dcea71616ef4f8e7b901bfb6c7f2cfd07ae7925b4c245aefca27a911424cb385850386a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff821680920361016a57602081600193829352019401920190610ce4565b50610d5e81610ee2565b60ff610d6d6024850184610ef0565b9290501611610b3a565b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361016a57817fa0c3b7890000000000000000000000000000000000000000000000000000000060209314908115610e09575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483610e02565b6040810190811067ffffffffffffffff821117610e4f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610e4f57604052565b60005b838110610ed25750506000910152565b8181015183820152602001610ec2565b3560ff8116810361016a5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b9190811015610f545760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff8116810361016a5790565b81601f8201121561016a5780359067ffffffffffffffff8211610e4f5760405192610ff7601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185610e7e565b8284526020838301011161016a57816000926020809301838601378301015290565b73ffffffffffffffffffffffffffffffffffffffff60015416330361103a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561016a5781519167ffffffffffffffff8311610e4f578260051b90604051936110986020840186610e7e565b845260208085019282010192831161016a57602001905b8282106110bc5750505090565b81518152602091820191016110af565b8051821015610f545760209160051b010190565b8054821015610f545760005260206000200190600090565b6000818152600460205260409020548015611287577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161070257600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161070257818103611218575b50505060035480156111e9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016111a68160036110e0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61126f61122961123a9360036110e0565b90549060031b1c92839260036110e0565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600460205260406000205538808061116d565b5050600090565b806000526004602052604060002054156000146112e85760035468010000000000000000811015610e4f576112cf61123a82600185940160035560036110e0565b9055600354906000526004602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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

func (_CommitOffRamp *CommitOffRampCaller) GetSignatureConfig(opts *bind.CallOpts) (SignatureQuorumVerifierSignatureConfigArgs, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "getSignatureConfig")

	if err != nil {
		return *new(SignatureQuorumVerifierSignatureConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(SignatureQuorumVerifierSignatureConfigArgs)).(*SignatureQuorumVerifierSignatureConfigArgs)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) GetSignatureConfig() (SignatureQuorumVerifierSignatureConfigArgs, error) {
	return _CommitOffRamp.Contract.GetSignatureConfig(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) GetSignatureConfig() (SignatureQuorumVerifierSignatureConfigArgs, error) {
	return _CommitOffRamp.Contract.GetSignatureConfig(&_CommitOffRamp.CallOpts)
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

func (_CommitOffRamp *CommitOffRampCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitOffRamp.Contract.SupportsInterface(&_CommitOffRamp.CallOpts, interfaceId)
}

func (_CommitOffRamp *CommitOffRampCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitOffRamp.Contract.SupportsInterface(&_CommitOffRamp.CallOpts, interfaceId)
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

func (_CommitOffRamp *CommitOffRampTransactor) SetSignatureConfig(opts *bind.TransactOpts, signatureConfig SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "setSignatureConfig", signatureConfig)
}

func (_CommitOffRamp *CommitOffRampSession) SetSignatureConfig(signatureConfig SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfig(&_CommitOffRamp.TransactOpts, signatureConfig)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) SetSignatureConfig(signatureConfig SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfig(&_CommitOffRamp.TransactOpts, signatureConfig)
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

func (_CommitOffRamp *CommitOffRampTransactor) ValidateReport(opts *bind.TransactOpts, message InternalMessageV1, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "validateReport", message, messageHash, ccvData, originalState)
}

func (_CommitOffRamp *CommitOffRampSession) ValidateReport(message InternalMessageV1, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, message, messageHash, ccvData, originalState)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) ValidateReport(message InternalMessageV1, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, message, messageHash, ccvData, originalState)
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
	Signers []common.Address
	F       uint8
	Raw     types.Log
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

func (CommitOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x65dcea71616ef4f8e7b901bfb6c7f2cfd07ae7925b4c245aefca27a911424cb3")
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
	GetSignatureConfig(opts *bind.CallOpts) (SignatureQuorumVerifierSignatureConfigArgs, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetSignatureConfig(opts *bind.TransactOpts, signatureConfig SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	ValidateReport(opts *bind.TransactOpts, message InternalMessageV1, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitOffRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitOffRampOwnershipTransferred, error)

	Address() common.Address
}
