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

type InternalAny2EVMMessage struct {
	Header       InternalHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     uint32
	TokenAmounts []InternalTokenTransfer
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	ExtraData         []byte
	Amount            *big.Int
}

type SignatureQuorumVerifierSignatureConfigArgs struct {
	ConfigDigest [32]byte
	F            uint8
	Signers      []common.Address
}

var CommitOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveConfigDigests\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllActiveConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs[]\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeConfigDigests\",\"inputs\":[{\"name\":\"configDigests\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignatureConfigs\",\"inputs\":[{\"name\":\"signatureConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureQuorumVerifier.SignatureConfigArgs[]\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateReport\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originalState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigRevoked\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigDigestAlreadyExists\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfigDigest\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0346100b757601f61174538819003918201601f19168301916001600160401b038311848410176100bc578084926020946040528339810103126100b757516001600160a01b0381168082036100b75733156100a657600180546001600160a01b0319163317905546608052156100955760a05260405161167290816100d3823960805181610d93015260a05181610f020152f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806301e0fdea14610b7257806301ffc9a714610ab3578063181f5a77146109e35780635982e36914610967578063741e4ba81461085657806379ba50971461076d578063827535cf146105425780638da5cb5b146104f0578063f2fde38b146103fd5763f50a00941461008a57600080fd5b346103f85760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85760043567ffffffffffffffff81116103f8576100d9903690600401611222565b906100e26113ca565b6000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa181360301925b808310156103f6576000928060051b830135858112156103f2578301906060823603126103f25760405195610140876111c5565b8235875260208301359260ff841684036103ea576020880193845260408101359067ffffffffffffffff82116103ee57019336601f860112156103ea57843561018881611312565b9561019660405197886111e1565b81875260208088019260051b820101903682116103e657602001915b8183106103b557505050604088019480865260ff85511680159182156103aa575b505080156103a1575b61037957875187526004602052604087205461034c57948751875260026020526040872091879660018401975b875180518210156102c2576102338273ffffffffffffffffffffffffffffffffffffffff9261132a565b51161561029a5761026573ffffffffffffffffffffffffffffffffffffffff61025d838b5161132a565b51168a611610565b1561027257600101610209565b60048a7f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b60048a7fd6c62c9b000000000000000000000000000000000000000000000000000000008152fd5b50509861033a9298507f5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e09391975060ff6001979596511691827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905551935190604051928392604084526040840190611253565b9060208301520390a20191909261010c565b60248789517f95e5047d000000000000000000000000000000000000000000000000000000008252600452fd5b6004877f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b508751156101dc565b5110905038806101d3565b823573ffffffffffffffffffffffffffffffffffffffff811681036103e2578152602092830192016101b2565b8a80fd5b8980fd5b8680fd5b8780fd5b8480fd5b005b600080fd5b346103f85760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85760043573ffffffffffffffffffffffffffffffffffffffff81168091036103f8576104556113ca565b3381146104c657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346103f85760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346103f85760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f857610579611415565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06105bf6105a984611312565b936105b760405195866111e1565b808552611312565b0160005b81811061074157505060005b815181101561069f576105e2818361132a565b516105ed828461132a565b51600052600260205260ff60406000205416610609838561132a565b51600052600260205260016040600020019060405191805480845260208401916000526020600020906000905b80821061068757505050906106528360019695949303836111e1565b6040519261065f846111c5565b835260208301526040820152610675828661132a565b52610680818561132a565b50016105cf565b90919260016020819286548152019401920190610636565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b8282106106d857505050500390f35b91936020610731827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652606060408a518051845260ff8682015116868501520151918160408201520190611253565b96019201920185949391926106c9565b602090604051610750816111c5565b6000815260008382015260606040820152828287010152016105c3565b346103f85760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85760005473ffffffffffffffffffffffffffffffffffffffff8116330361082c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346103f85760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85760043567ffffffffffffffff81116103f8576108a5903690600401611222565b6108ad6113ca565b60005b818110156103f65760008160051b840135906108cb8261147a565b1561093c578181526002602052600160408220828155018054908281558161091e575b5050907ffdde4bfc1a9ef28a2e3dbe34a4ccc65b0ad588f6b0406e492637aeaa733421608260019493a2016108b0565b825260208220908101905b818110156108ee57828155600101610929565b6024917f2f01e576000000000000000000000000000000000000000000000000000000008252600452fd5b346103f85760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85761099e611415565b60405180916020820160208352815180915260206040840192019060005b8181106109ca575050500390f35b82518452859450602093840193909201916001016109bc565b346103f85760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f8576040805190610a2181836111e1565b601782527f436f6d6d69744f666652616d7020312e372e302d6465760000000000000000006020830152805180926020825280519081602084015260005b828110610a9c5750506000828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b602082820181015187830187015286945001610a5f565b346103f85760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f8576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036103f857807f01e0fdea0000000000000000000000000000000000000000000000000000000060209214908115610b48575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501482610b3d565b346103f85760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103f85760043567ffffffffffffffff81116103f8578036036101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126103f8576044359167ffffffffffffffff83116103f857366023840112156103f857826004013567ffffffffffffffff81116103f85783019160248301933685116103f8576064359360048510156103f857604090829003126103f857602481013567ffffffffffffffff81116103f857856024610c5d9284010161129d565b9460448201359167ffffffffffffffff83116103f857610c80920160240161129d565b93805194602082019560408382810103126103f857604087519301519667ffffffffffffffff88168098036103f8572060405160208101916024358352604082015260408152610cd16060826111e1565b51902090826000526002602052604060002091600183019384541561119857508151820160208101926020818303126103f857602081015167ffffffffffffffff81116103f8570190604090829003126103f857604051926040840184811067ffffffffffffffff82111761116957604052602082015167ffffffffffffffff81116103f857816020610d669285010161136d565b845260408201519167ffffffffffffffff83116103f857610d8a920160200161136d565b602083019081527f00000000000000000000000000000000000000000000000000000000000000004681036111385750600160ff8451519554160160ff81116111095760ff1684036110df5780515184036110b557600092835b858510610fb3578a8a8a8a83610df657005b60009215610e0057005b60248201359167ffffffffffffffff8316809303610faf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60848201359201821215610faf57019060048201359167ffffffffffffffff8311610faf57602401908236038213610faf579060846020927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8660405197889687957fe0e03cae00000000000000000000000000000000000000000000000000000000875260048701528b6024870152606060448701528160648701528686013788858286010152011681010301818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610fa4578291610f65575b5015610f3a57005b6024917f5c33785a000000000000000000000000000000000000000000000000000000008252600452fd5b90506020813d602011610f9c575b81610f80602093836111e1565b81010312610f9857518015158103610f985783610f32565b5080fd5b3d9150610f73565b6040513d84823e3d90fd5b8380fd5b602060006080610fc488865161132a565b51610fd089885161132a565b5160405191898352601b868401526040830152606082015282805260015afa156110a95773ffffffffffffffffffffffffffffffffffffffff6000511690611028828960019160005201602052604060002054151590565b1561107f5773ffffffffffffffffffffffffffffffffffffffff1681111561105557600190940193610de4565b7ff67bc7c40000000000000000000000000000000000000000000000000000000060005260046000fd5b7fca31867a0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7fa75d88af0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f2f01e5760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6060810190811067ffffffffffffffff82111761116957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761116957604052565b9181601f840112156103f85782359167ffffffffffffffff83116103f8576020808501948460051b0101116103f857565b906020808351928381520192019060005b8181106112715750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611264565b81601f820112156103f85780359067ffffffffffffffff821161116957604051926112f0601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856111e1565b828452602083830101116103f857816000926020809301838601378301015290565b67ffffffffffffffff81116111695760051b60200190565b805182101561133e5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f830112156103f857815161138481611312565b9261139260405194856111e1565b81845260208085019260051b8201019283116103f857602001905b8282106113ba5750505090565b81518152602091820191016113ad565b73ffffffffffffffffffffffffffffffffffffffff6001541633036113eb57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906003548083528260208101600360005260206000209260005b818110611449575050611447925003836111e1565b565b8454835260019485019487945060209093019201611432565b805482101561133e5760005260206000200190600090565b6000818152600460205260409020548015611609577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161110957600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111095781810361159a575b505050600354801561156b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611528816003611462565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6115f16115ab6115bc936003611462565b90549060031b1c9283926003611462565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260046020526040600020553880806114ef565b5050600090565b60008281526001820160205260409020546116095780549068010000000000000000821015611169578261164e6115bc846001809601855584611462565b90558054926000520160205260406000205560019056fea164736f6c634300081a000a",
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

func (_CommitOffRamp *CommitOffRampCaller) GetActiveConfigDigests(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "getActiveConfigDigests")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) GetActiveConfigDigests() ([][32]byte, error) {
	return _CommitOffRamp.Contract.GetActiveConfigDigests(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) GetActiveConfigDigests() ([][32]byte, error) {
	return _CommitOffRamp.Contract.GetActiveConfigDigests(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCaller) GetAllActiveConfigs(opts *bind.CallOpts) ([]SignatureQuorumVerifierSignatureConfigArgs, error) {
	var out []interface{}
	err := _CommitOffRamp.contract.Call(opts, &out, "getAllActiveConfigs")

	if err != nil {
		return *new([]SignatureQuorumVerifierSignatureConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]SignatureQuorumVerifierSignatureConfigArgs)).(*[]SignatureQuorumVerifierSignatureConfigArgs)

	return out0, err

}

func (_CommitOffRamp *CommitOffRampSession) GetAllActiveConfigs() ([]SignatureQuorumVerifierSignatureConfigArgs, error) {
	return _CommitOffRamp.Contract.GetAllActiveConfigs(&_CommitOffRamp.CallOpts)
}

func (_CommitOffRamp *CommitOffRampCallerSession) GetAllActiveConfigs() ([]SignatureQuorumVerifierSignatureConfigArgs, error) {
	return _CommitOffRamp.Contract.GetAllActiveConfigs(&_CommitOffRamp.CallOpts)
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

func (_CommitOffRamp *CommitOffRampTransactor) RevokeConfigDigests(opts *bind.TransactOpts, configDigests [][32]byte) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "revokeConfigDigests", configDigests)
}

func (_CommitOffRamp *CommitOffRampSession) RevokeConfigDigests(configDigests [][32]byte) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.RevokeConfigDigests(&_CommitOffRamp.TransactOpts, configDigests)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) RevokeConfigDigests(configDigests [][32]byte) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.RevokeConfigDigests(&_CommitOffRamp.TransactOpts, configDigests)
}

func (_CommitOffRamp *CommitOffRampTransactor) SetSignatureConfigs(opts *bind.TransactOpts, signatureConfigs []SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "setSignatureConfigs", signatureConfigs)
}

func (_CommitOffRamp *CommitOffRampSession) SetSignatureConfigs(signatureConfigs []SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfigs(&_CommitOffRamp.TransactOpts, signatureConfigs)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) SetSignatureConfigs(signatureConfigs []SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.SetSignatureConfigs(&_CommitOffRamp.TransactOpts, signatureConfigs)
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

func (_CommitOffRamp *CommitOffRampTransactor) ValidateReport(opts *bind.TransactOpts, message InternalAny2EVMMessage, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.contract.Transact(opts, "validateReport", message, messageHash, ccvData, originalState)
}

func (_CommitOffRamp *CommitOffRampSession) ValidateReport(message InternalAny2EVMMessage, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, message, messageHash, ccvData, originalState)
}

func (_CommitOffRamp *CommitOffRampTransactorSession) ValidateReport(message InternalAny2EVMMessage, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error) {
	return _CommitOffRamp.Contract.ValidateReport(&_CommitOffRamp.TransactOpts, message, messageHash, ccvData, originalState)
}

type CommitOffRampConfigRevokedIterator struct {
	Event *CommitOffRampConfigRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOffRampConfigRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOffRampConfigRevoked)
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
		it.Event = new(CommitOffRampConfigRevoked)
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

func (it *CommitOffRampConfigRevokedIterator) Error() error {
	return it.fail
}

func (it *CommitOffRampConfigRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOffRampConfigRevoked struct {
	ConfigDigest [32]byte
	Raw          types.Log
}

func (_CommitOffRamp *CommitOffRampFilterer) FilterConfigRevoked(opts *bind.FilterOpts, configDigest [][32]byte) (*CommitOffRampConfigRevokedIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _CommitOffRamp.contract.FilterLogs(opts, "ConfigRevoked", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampConfigRevokedIterator{contract: _CommitOffRamp.contract, event: "ConfigRevoked", logs: logs, sub: sub}, nil
}

func (_CommitOffRamp *CommitOffRampFilterer) WatchConfigRevoked(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigRevoked, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _CommitOffRamp.contract.WatchLogs(opts, "ConfigRevoked", configDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOffRampConfigRevoked)
				if err := _CommitOffRamp.contract.UnpackLog(event, "ConfigRevoked", log); err != nil {
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

func (_CommitOffRamp *CommitOffRampFilterer) ParseConfigRevoked(log types.Log) (*CommitOffRampConfigRevoked, error) {
	event := new(CommitOffRampConfigRevoked)
	if err := _CommitOffRamp.contract.UnpackLog(event, "ConfigRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

func (_CommitOffRamp *CommitOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts, configDigest [][32]byte) (*CommitOffRampConfigSetIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _CommitOffRamp.contract.FilterLogs(opts, "ConfigSet", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &CommitOffRampConfigSetIterator{contract: _CommitOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOffRamp *CommitOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigSet, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _CommitOffRamp.contract.WatchLogs(opts, "ConfigSet", configDigestRule)
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

func (CommitOffRampConfigRevoked) Topic() common.Hash {
	return common.HexToHash("0xfdde4bfc1a9ef28a2e3dbe34a4ccc65b0ad588f6b0406e492637aeaa73342160")
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
	GetActiveConfigDigests(opts *bind.CallOpts) ([][32]byte, error)

	GetAllActiveConfigs(opts *bind.CallOpts) ([]SignatureQuorumVerifierSignatureConfigArgs, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RevokeConfigDigests(opts *bind.TransactOpts, configDigests [][32]byte) (*types.Transaction, error)

	SetSignatureConfigs(opts *bind.TransactOpts, signatureConfigs []SignatureQuorumVerifierSignatureConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	ValidateReport(opts *bind.TransactOpts, message InternalAny2EVMMessage, messageHash [32]byte, ccvData []byte, originalState uint8) (*types.Transaction, error)

	FilterConfigRevoked(opts *bind.FilterOpts, configDigest [][32]byte) (*CommitOffRampConfigRevokedIterator, error)

	WatchConfigRevoked(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigRevoked, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigRevoked(log types.Log) (*CommitOffRampConfigRevoked, error)

	FilterConfigSet(opts *bind.FilterOpts, configDigest [][32]byte) (*CommitOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOffRampConfigSet, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitOffRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitOffRampOwnershipTransferred, error)

	Address() common.Address
}
