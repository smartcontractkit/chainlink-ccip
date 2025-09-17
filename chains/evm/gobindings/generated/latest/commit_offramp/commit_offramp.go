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

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	OnRampAddress       []byte
	OffRampAddress      []byte
	Finality            uint16
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []MessageV1CodecTokenTransferV1
	Data                []byte
}

type MessageV1CodecTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	ExtraData          []byte
}

type SignatureQuorumVerifierSignatureConfigArgs struct {
	Threshold uint8
	Signers   []common.Address
}

var CommitOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSignatureConfig\",\"inputs\":[{\"name\":\"signatureConfig\",\"type\":\"tuple\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0346100b457601f61115138819003918201601f19168301916001600160401b038311848410176100b9578084926020946040528339810103126100b457516001600160a01b0381168082036100b45733156100a357600180546001600160a01b0319163317905546608052156100925760a05260405161108190816100d0823960805181610369015260a051815050f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714610b725750806310803bd2146108a8578063181f5a77146107f65780636ed0e217146106c257806379ba5097146105d95780638da5cb5b14610587578063bc5931a41461016f5763f2fde38b1461007757600080fd5b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043573ffffffffffffffffffffffffffffffffffffffff811680910361016a576100cf610dd8565b33811461014057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461016a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc610160913603011261016a5760443567ffffffffffffffff811161016a573660238201121561016a57806004013567ffffffffffffffff811161016a57810190602482019136831161016a576004606435101561016a576060908290031261016a57602481013567ffffffffffffffff811161016a578101908260438301121561016a5760248201359067ffffffffffffffff821161055857604051916102a0601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184610c4a565b8083526020830193604481830101861061016a5781600092604460209301873784010152604481013567ffffffffffffffff811161016a578460246102e792840101610d70565b9360648201359167ffffffffffffffff831161016a5761030a9201602401610d70565b9161035c60408051809361032d6020830196602435885251809285850190610c8b565b810103017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610c4a565b5190206003541561052e577f00000000000000000000000000000000000000000000000000000000000000004681036104fd575082519160ff6002541683036104d357805183036104a957600093845b8486106103b557005b6020600060806103c58986610e23565b516103d08a88610e23565b5160405191898352601b868401526040830152606082015282805260015afa1561049d5773ffffffffffffffffffffffffffffffffffffffff60005116906040600083815260046020522054156104735773ffffffffffffffffffffffffffffffffffffffff16811115610449576001909501946103ac565b7ff67bc7c40000000000000000000000000000000000000000000000000000000060005260046000fd5b7fca31867a0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7fa75d88af0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760005473ffffffffffffffffffffffffffffffffffffffff81163303610698577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a576060602060405161070181610c2e565b60008152015260ff6002541660405180602060035491828152019060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b8181106107e0575050508161075c910382610c4a565b6040519161076983610c2e565b8252602082019081526040519182916020835260ff6060840192511660208401525190604080840152815180915260206080840192019060005b8181106107b1575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016107a3565b8254845260209093019260019283019201610746565b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5761089f604080516108368282610c4a565b601781527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602083017f436f6d6d69744f666652616d7020312e372e302d646576000000000000000000815284519586946020865251809281602088015287870190610c8b565b01168101030190f35b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a578060040160407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261016a57610921610dd8565b60ff61092c82610cae565b16158015610b4f575b61052e575b600354156109c45760006003541561099757600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b546109919073ffffffffffffffffffffffffffffffffffffffff16610e4f565b5061093a565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b9060240160005b6109d58284610cbc565b9050811015610a7b5773ffffffffffffffffffffffffffffffffffffffff610a0f610a0a83610a048688610cbc565b90610d10565b610d4f565b1615610a5157610a4473ffffffffffffffffffffffffffffffffffffffff610a3e610a0a84610a048789610cbc565b16611014565b1561052e576001016109cb565b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50610abe610ac69160ff610a8e85610cae565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00600254161760025583610cbc565b919092610cae565b9160405191806040840160408552526060830191906000905b808210610b175760ff861660208601527f65dcea71616ef4f8e7b901bfb6c7f2cfd07ae7925b4c245aefca27a911424cb385850386a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff821680920361016a57602081600193829352019401920190610adf565b50610b5981610cae565b60ff610b686024850184610cbc565b9290501611610935565b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361016a57817fbc5931a40000000000000000000000000000000000000000000000000000000060209314908115610c04575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483610bfd565b6040810190811067ffffffffffffffff82111761055857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761055857604052565b60005b838110610c9e5750506000910152565b8181015183820152602001610c8e565b3560ff8116810361016a5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b9190811015610d205760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff8116810361016a5790565b9080601f8301121561016a5781359167ffffffffffffffff8311610558578260051b9060405193610da46020840186610c4a565b845260208085019282010192831161016a57602001905b828210610dc85750505090565b8135815260209182019101610dbb565b73ffffffffffffffffffffffffffffffffffffffff600154163303610df957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015610d205760209160051b010190565b8054821015610d205760005260206000200190600090565b600081815260046020526040902054801561100d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610fde57600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610fde57818103610f6f575b5050506003548015610f40577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610efd816003610e37565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610fc6610f80610f91936003610e37565b90549060031b1c9283926003610e37565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080610ec4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b8060005260046020526040600020541560001461106e576003546801000000000000000081101561055857611055610f918260018594016003556003610e37565b9055600354906000526004602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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

func (_CommitOffRamp *CommitOffRampTransactor) VerifyMessage(opts *bind.TransactOpts, arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte, arg3 uint8) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "verifyMessage", arg0, messageHash, ccvData, arg3)
}

func (_CommitOffRamp *CommitOffRampSession) VerifyMessage(arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte, arg3 uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.VerifyMessage(&_CommitOffRamp.TransactOpts, arg0, messageHash, ccvData, arg3)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) VerifyMessage(arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte, arg3 uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.VerifyMessage(&_CommitOffRamp.TransactOpts, arg0, messageHash, ccvData, arg3)
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

	VerifyMessage(opts *bind.TransactOpts, arg0 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte, arg3 uint8) (*types.Transaction, error)

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
