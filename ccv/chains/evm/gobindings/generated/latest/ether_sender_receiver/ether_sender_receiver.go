// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ether_sender_receiver

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var EtherSenderReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSend\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getCCVsAndMinBlockConfirmations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IWrappedNative\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"gotToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expectedToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAmounts\",\"inputs\":[{\"name\":\"gotAmounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenAmountNotEqualToMsgValue\",\"inputs\":[{\"name\":\"gotAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c080604052346101245761002a906115d180380380916100208285610180565b83398101906101b9565b6001600160a01b03811690811561016a5760805260405163e861e90760e01b815290602082600481845afa908115610131576044602092600094859161013d575b506001600160a01b031660a081905260405163095ea7b360e01b8152600481019390935284196024840152919384928391905af18015610131576100f4575b6040516113f890816101d9823960805181818160c801528181610190015281816107210152610ede015260a0518181816102da0152818161069001528181610e6501526113920152f35b6020813d602011610129575b8161010d60209383610180565b8101031261012457518015150361012457386100aa565b600080fd5b3d9150610100565b6040513d6000823e3d90fd5b61015d9150843d8611610163575b6101558183610180565b8101906101b9565b3861006b565b503d61014b565b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b038211908210176101a357604052565b634e487b7160e01b600052604160045260246000fd5b9081602091031261012457516001600160a01b0381168103610124579056fe608080604052600436101561001d575b50361561001b57600080fd5b005b600090813560e01c90816301ffc9a71461081557508063181f5a771461079257806320487ded146106b45780634dbe7e921461064557806370709a4d1461056a57806385572ffb1461010b57806396f4e9f9146100ef5763b0f479a10361000f57346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b80fd5b60206101036100fd36610a45565b90610dbc565b604051908152f35b50346100ec5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec5760043567ffffffffffffffff81116105665760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105665773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361053a576040516101c281610904565b81600401358152602482013567ffffffffffffffff81168103610492576020820152604482013567ffffffffffffffff8111610492576102089060043691850101610c58565b6040820152606482013567ffffffffffffffff8111610492576102319060043691850101610c58565b9160608201928352608481013567ffffffffffffffff811161053657608091600461025f9236920101610cc0565b9101918183525160208180518101031261049257602001519073ffffffffffffffffffffffffffffffffffffffff821680920361049257516001810361050b575073ffffffffffffffffffffffffffffffffffffffff6102bf8351610d68565b5151169173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168093036104ba5761030d60209151610d68565b51015190823b15610492576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152826004820152848160248183885af180156104af5761049b575b508380808085855af13d15610496573d61037181610990565b9061037f604051928361094f565b81528560203d92013e5b15610392578380f35b823b1561049257836040517fd0e30db0000000000000000000000000000000000000000000000000000000008152818160048187895af1801561046e57610479575b506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482019290925291602091839160449183915af1801561046e5761043f575b80808380f35b6104609060203d602011610467575b610458818361094f565b810190610da4565b5038610439565b503d61044e565b6040513d84823e3d90fd5b816104869194939461094f565b610492579083386103d4565b8380fd5b610389565b846104a89195929561094f565b9238610358565b6040513d87823e3d90fd5b838373ffffffffffffffffffffffffffffffffffffffff6104dd60449451610d68565b5151167f0fc746a1000000000000000000000000000000000000000000000000000000008352600452602452fd5b7f83b9f0ae000000000000000000000000000000000000000000000000000000008452600452602483fd5b8480fd5b6024827fd7f7333400000000000000000000000000000000000000000000000000000000815233600452fd5b5080fd5b50346100ec5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec576105a2610a29565b5060243567ffffffffffffffff8111610566573660238201121561056657806004013567ffffffffffffffff811161064157369101602401116100ec576040516106249161063160206105f5818561094f565b82845282368137604051610609828261094f565b83815283368137604051958695608087526080870190610aca565b9185830390860152610aca565b9080604084015260608301520390f35b8280fd5b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100ec5761070860206106d36106cb36610a45565b91909161126c565b9060405193849283927f20487ded00000000000000000000000000000000000000000000000000000000845260048401610b14565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561046e57829161075c575b602082604051908152f35b90506020813d60201161078a575b816107776020938361094f565b8101031261056657602091505138610751565b3d915061076a565b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57506108116040516107d360408261094f565b601981527f457468657253656e646572526563656976657220312e352e300000000000000060208201526040519182916020835260208301906109ca565b0390f35b9050346105665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610566576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361064157602092507f85572ffb0000000000000000000000000000000000000000000000000000000081149081156108da575b81156108b0575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386108a9565b7f70709a4d00000000000000000000000000000000000000000000000000000000811491506108a2565b60a0810190811067ffffffffffffffff82111761092057604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761092057604052565b67ffffffffffffffff811161092057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110610a145750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016109d5565b6004359067ffffffffffffffff82168203610a4057565b600080fd5b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112610a405760043567ffffffffffffffff81168103610a4057916024359067ffffffffffffffff8211610a40577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a092030112610a405760040190565b906020808351928381520192019060005b818110610ae85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610adb565b9067ffffffffffffffff9093929316815260406020820152610b79610b45845160a0604085015260e08401906109ca565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030160608501526109ca565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b818110610c205750505060808473ffffffffffffffffffffffffffffffffffffffff6060610c1d969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828503019101526109ca565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101610bba565b81601f82011215610a4057803590610c6f82610990565b92610c7d604051948561094f565b82845260208383010111610a4057816000926020809301838601378301015290565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a4057565b81601f82011215610a405780359067ffffffffffffffff82116109205760208260051b0192610cf2604051948561094f565b82845260208085019360061b83010191818311610a4057602001925b828410610d1c575050505090565b604084830312610a405760405190604082019082821067ffffffffffffffff831117610920576040926020928452610d5387610c9f565b81528287013583820152815201930192610d0e565b805115610d755760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90816020910312610a4057518015158103610a405790565b60009160408101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610492578101803567ffffffffffffffff8111610536578060061b36036020830113610536571561123f5760400135606082013573ffffffffffffffffffffffffffffffffffffffff811680910361053657611209575b50610e4d9061126c565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906020610e956040850151610d68565b510151823b156105365784600491604051928380927fd0e30db0000000000000000000000000000000000000000000000000000000008252875af180156104af576111f5575b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff811692604051917f20487ded00000000000000000000000000000000000000000000000000000000835260208380610f4e898860048401610b14565b0381885afa9283156111ea5787936111b6575b506060860173ffffffffffffffffffffffffffffffffffffffff8151168061101557505050505091602091610fc693476040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610b14565b03925af19182156110095791610fda575090565b90506020813d602011611001575b81610ff56020938361094f565b81010312610a40575190565b3d9150610fe8565b604051903d90823e3d90fd5b602089604051828101907f23b872dd0000000000000000000000000000000000000000000000000000000082523360248201523060448201528860648201526064815261106360848261094f565b519082855af1156111ab5788513d6111a25750803b155b61117757505173ffffffffffffffffffffffffffffffffffffffff169182036110df575b50505091602091610fc693856040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610b14565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101929092526020908290604490829089905af180156104af579160209391610fc6959361115a575b8193955082945061109e565b61117090853d871161046757610458818361094f565b503861114e565b7f5274afe7000000000000000000000000000000000000000000000000000000008952600452602488fd5b6001141561107a565b6040513d8a823e3d90fd5b9092506020813d6020116111e2575b816111d26020938361094f565b81010312610a4057519138610f61565b3d91506111c5565b6040513d89823e3d90fd5b846112029195929561094f565b9238610edb565b348114610e43577fba2f746700000000000000000000000000000000000000000000000000000000845260045234602452604483fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b6060608060405161127c81610904565b828152826020820152826040820152600083820152015260a081360312610a4057604051906112aa82610904565b803567ffffffffffffffff8111610a40576112c89036908301610c58565b8252602081013567ffffffffffffffff8111610a40576112eb9036908301610c58565b60208301908152604082013567ffffffffffffffff8111610a40576113139036908401610cc0565b916040840192835261132760608201610c9f565b606085015260808101359067ffffffffffffffff8211610a405761134d91369101610c58565b6080840152815151600181036113be5750604051903360208301526020825261137760408361094f565b526113b973ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169151610d68565b515290565b7f83b9f0ae0000000000000000000000000000000000000000000000000000000060005260045260246000fdfea164736f6c634300081a000a",
}

var EtherSenderReceiverABI = EtherSenderReceiverMetaData.ABI

var EtherSenderReceiverBin = EtherSenderReceiverMetaData.Bin

func DeployEtherSenderReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *EtherSenderReceiver, error) {
	parsed, err := EtherSenderReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherSenderReceiverBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherSenderReceiver{address: address, abi: *parsed, EtherSenderReceiverCaller: EtherSenderReceiverCaller{contract: contract}, EtherSenderReceiverTransactor: EtherSenderReceiverTransactor{contract: contract}, EtherSenderReceiverFilterer: EtherSenderReceiverFilterer{contract: contract}}, nil
}

type EtherSenderReceiver struct {
	address common.Address
	abi     abi.ABI
	EtherSenderReceiverCaller
	EtherSenderReceiverTransactor
	EtherSenderReceiverFilterer
}

type EtherSenderReceiverCaller struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverTransactor struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverFilterer struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverSession struct {
	Contract     *EtherSenderReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EtherSenderReceiverCallerSession struct {
	Contract *EtherSenderReceiverCaller
	CallOpts bind.CallOpts
}

type EtherSenderReceiverTransactorSession struct {
	Contract     *EtherSenderReceiverTransactor
	TransactOpts bind.TransactOpts
}

type EtherSenderReceiverRaw struct {
	Contract *EtherSenderReceiver
}

type EtherSenderReceiverCallerRaw struct {
	Contract *EtherSenderReceiverCaller
}

type EtherSenderReceiverTransactorRaw struct {
	Contract *EtherSenderReceiverTransactor
}

func NewEtherSenderReceiver(address common.Address, backend bind.ContractBackend) (*EtherSenderReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(EtherSenderReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEtherSenderReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiver{address: address, abi: abi, EtherSenderReceiverCaller: EtherSenderReceiverCaller{contract: contract}, EtherSenderReceiverTransactor: EtherSenderReceiverTransactor{contract: contract}, EtherSenderReceiverFilterer: EtherSenderReceiverFilterer{contract: contract}}, nil
}

func NewEtherSenderReceiverCaller(address common.Address, caller bind.ContractCaller) (*EtherSenderReceiverCaller, error) {
	contract, err := bindEtherSenderReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverCaller{contract: contract}, nil
}

func NewEtherSenderReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherSenderReceiverTransactor, error) {
	contract, err := bindEtherSenderReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverTransactor{contract: contract}, nil
}

func NewEtherSenderReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherSenderReceiverFilterer, error) {
	contract, err := bindEtherSenderReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverFilterer{contract: contract}, nil
}

func bindEtherSenderReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherSenderReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverTransactor.contract.Transfer(opts)
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherSenderReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.contract.Transfer(opts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetCCVsAndMinBlockConfirmations(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getCCVsAndMinBlockConfirmations", arg0, arg1)

	outstruct := new(GetCCVsAndMinBlockConfirmations)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetCCVsAndMinBlockConfirmations(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVsAndMinBlockConfirmations(&_EtherSenderReceiver.CallOpts, arg0, arg1)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetCCVsAndMinBlockConfirmations(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVsAndMinBlockConfirmations(&_EtherSenderReceiver.CallOpts, arg0, arg1)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getFee", destinationChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EtherSenderReceiver.Contract.GetFee(&_EtherSenderReceiver.CallOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EtherSenderReceiver.Contract.GetFee(&_EtherSenderReceiver.CallOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetRouter() (common.Address, error) {
	return _EtherSenderReceiver.Contract.GetRouter(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetRouter() (common.Address, error) {
	return _EtherSenderReceiver.Contract.GetRouter(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) IWeth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "i_weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) IWeth() (common.Address, error) {
	return _EtherSenderReceiver.Contract.IWeth(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) IWeth() (common.Address, error) {
	return _EtherSenderReceiver.Contract.IWeth(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherSenderReceiver.Contract.SupportsInterface(&_EtherSenderReceiver.CallOpts, interfaceId)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherSenderReceiver.Contract.SupportsInterface(&_EtherSenderReceiver.CallOpts, interfaceId)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) TypeAndVersion() (string, error) {
	return _EtherSenderReceiver.Contract.TypeAndVersion(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) TypeAndVersion() (string, error) {
	return _EtherSenderReceiver.Contract.TypeAndVersion(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.Transact(opts, "ccipReceive", message)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipReceive(&_EtherSenderReceiver.TransactOpts, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipReceive(&_EtherSenderReceiver.TransactOpts, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.Transact(opts, "ccipSend", destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipSend(&_EtherSenderReceiver.TransactOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipSend(&_EtherSenderReceiver.TransactOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.RawTransact(opts, nil)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) Receive() (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.Receive(&_EtherSenderReceiver.TransactOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.Receive(&_EtherSenderReceiver.TransactOpts)
}

type GetCCVsAndMinBlockConfirmations struct {
	RequiredCCVs          []common.Address
	OptionalCCVs          []common.Address
	OptionalThreshold     uint8
	MinBlockConfirmations uint16
}

func (_EtherSenderReceiver *EtherSenderReceiver) Address() common.Address {
	return _EtherSenderReceiver.address
}

type EtherSenderReceiverInterface interface {
	GetCCVsAndMinBlockConfirmations(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

		error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IWeth(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	Address() common.Address
}
