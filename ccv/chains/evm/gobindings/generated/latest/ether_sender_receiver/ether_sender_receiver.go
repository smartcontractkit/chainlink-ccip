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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSend\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IWrappedNative\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"gotToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expectedToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAmounts\",\"inputs\":[{\"name\":\"gotAmounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenAmountNotEqualToMsgValue\",\"inputs\":[{\"name\":\"gotAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c080604052346101245761002a9061159080380380916100208285610180565b83398101906101b9565b6001600160a01b03811690811561016a5760805260405163e861e90760e01b815290602082600481845afa908115610131576044602092600094859161013d575b506001600160a01b031660a081905260405163095ea7b360e01b8152600481019390935284196024840152919384928391905af18015610131576100f4575b6040516113b790816101d9823960805181818160c801528181610190015281816106dc0152610e9d015260a0518181816102da0152818161064b01528181610e2401526113510152f35b6020813d602011610129575b8161010d60209383610180565b8101031261012457518015150361012457386100aa565b600080fd5b3d9150610100565b6040513d6000823e3d90fd5b61015d9150843d8611610163575b6101558183610180565b8101906101b9565b3861006b565b503d61014b565b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b038211908210176101a357604052565b634e487b7160e01b600052604160045260246000fd5b9081602091031261012457516001600160a01b0381168103610124579056fe608080604052600436101561001d575b50361561001b57600080fd5b005b600090813560e01c90816301ffc9a7146107d057508063181f5a771461074d57806320487ded1461066f5780634dbe7e92146106005780637909b5491461056a57806385572ffb1461010b57806396f4e9f9146100ef5763b0f479a10361000f57346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b80fd5b60206101036100fd36610a04565b90610d7b565b604051908152f35b50346100ec5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec5760043567ffffffffffffffff81116105665760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105665773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361053a576040516101c2816108c3565b81600401358152602482013567ffffffffffffffff81168103610492576020820152604482013567ffffffffffffffff8111610492576102089060043691850101610c17565b6040820152606482013567ffffffffffffffff8111610492576102319060043691850101610c17565b9160608201928352608481013567ffffffffffffffff811161053657608091600461025f9236920101610c7f565b9101918183525160208180518101031261049257602001519073ffffffffffffffffffffffffffffffffffffffff821680920361049257516001810361050b575073ffffffffffffffffffffffffffffffffffffffff6102bf8351610d27565b5151169173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168093036104ba5761030d60209151610d27565b51015190823b15610492576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152826004820152848160248183885af180156104af5761049b575b508380808085855af13d15610496573d6103718161094f565b9061037f604051928361090e565b81528560203d92013e5b15610392578380f35b823b1561049257836040517fd0e30db0000000000000000000000000000000000000000000000000000000008152818160048187895af1801561046e57610479575b506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482019290925291602091839160449183915af1801561046e5761043f575b80808380f35b6104609060203d602011610467575b610458818361090e565b810190610d63565b5038610439565b503d61044e565b6040513d84823e3d90fd5b816104869194939461090e565b610492579083386103d4565b8380fd5b610389565b846104a89195929561090e565b9238610358565b6040513d87823e3d90fd5b838373ffffffffffffffffffffffffffffffffffffffff6104dd60449451610d27565b5151167f0fc746a1000000000000000000000000000000000000000000000000000000008352600452602452fd5b7f83b9f0ae000000000000000000000000000000000000000000000000000000008452600452602483fd5b8480fd5b6024827fd7f7333400000000000000000000000000000000000000000000000000000000815233600452fd5b5080fd5b50346100ec5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec576105e9906105a66109e8565b50604051906105f660206105ba818561090e565b828452823681376040516105ce828261090e565b83815283368137604051958695606087526060870190610a89565b9185830390860152610a89565b9060408301520390f35b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100ec576106c3602061068e61068636610a04565b91909161122b565b9060405193849283927f20487ded00000000000000000000000000000000000000000000000000000000845260048401610ad3565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561046e578291610717575b602082604051908152f35b90506020813d602011610745575b816107326020938361090e565b810103126105665760209150513861070c565b3d9150610725565b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57506107cc60405161078e60408261090e565b601981527f457468657253656e646572526563656976657220312e352e30000000000000006020820152604051918291602083526020830190610989565b0390f35b9050346105665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610566576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036108bf57602092507f85572ffb000000000000000000000000000000000000000000000000000000008114908115610895575b811561086b575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610864565b7f7909b549000000000000000000000000000000000000000000000000000000008114915061085d565b8280fd5b60a0810190811067ffffffffffffffff8211176108df57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108df57604052565b67ffffffffffffffff81116108df57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106109d35750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610994565b6004359067ffffffffffffffff821682036109ff57565b600080fd5b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126109ff5760043567ffffffffffffffff811681036109ff57916024359067ffffffffffffffff82116109ff577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126109ff5760040190565b906020808351928381520192019060005b818110610aa75750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610a9a565b9067ffffffffffffffff9093929316815260406020820152610b38610b04845160a0604085015260e0840190610989565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152610989565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b818110610bdf5750505060808473ffffffffffffffffffffffffffffffffffffffff6060610bdc969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152610989565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101610b79565b81601f820112156109ff57803590610c2e8261094f565b92610c3c604051948561090e565b828452602083830101116109ff57816000926020809301838601378301015290565b359073ffffffffffffffffffffffffffffffffffffffff821682036109ff57565b81601f820112156109ff5780359067ffffffffffffffff82116108df5760208260051b0192610cb1604051948561090e565b82845260208085019360061b830101918183116109ff57602001925b828410610cdb575050505090565b6040848303126109ff5760405190604082019082821067ffffffffffffffff8311176108df576040926020928452610d1287610c5e565b81528287013583820152815201930192610ccd565b805115610d345760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b908160209103126109ff575180151581036109ff5790565b60009160408101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610492578101803567ffffffffffffffff8111610536578060061b3603602083011361053657156111fe5760400135606082013573ffffffffffffffffffffffffffffffffffffffff8116809103610536576111c8575b50610e0c9061122b565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906020610e546040850151610d27565b510151823b156105365784600491604051928380927fd0e30db0000000000000000000000000000000000000000000000000000000008252875af180156104af576111b4575b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff811692604051917f20487ded00000000000000000000000000000000000000000000000000000000835260208380610f0d898860048401610ad3565b0381885afa9283156111a9578793611175575b506060860173ffffffffffffffffffffffffffffffffffffffff81511680610fd457505050505091602091610f8593476040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610ad3565b03925af1918215610fc85791610f99575090565b90506020813d602011610fc0575b81610fb46020938361090e565b810103126109ff575190565b3d9150610fa7565b604051903d90823e3d90fd5b602089604051828101907f23b872dd0000000000000000000000000000000000000000000000000000000082523360248201523060448201528860648201526064815261102260848261090e565b519082855af11561116a5788513d6111615750803b155b61113657505173ffffffffffffffffffffffffffffffffffffffff1691820361109e575b50505091602091610f8593856040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610ad3565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101929092526020908290604490829089905af180156104af579160209391610f859593611119575b8193955082945061105d565b61112f90853d871161046757610458818361090e565b503861110d565b7f5274afe7000000000000000000000000000000000000000000000000000000008952600452602488fd5b60011415611039565b6040513d8a823e3d90fd5b9092506020813d6020116111a1575b816111916020938361090e565b810103126109ff57519138610f20565b3d9150611184565b6040513d89823e3d90fd5b846111c19195929561090e565b9238610e9a565b348114610e02577fba2f746700000000000000000000000000000000000000000000000000000000845260045234602452604483fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b6060608060405161123b816108c3565b828152826020820152826040820152600083820152015260a0813603126109ff5760405190611269826108c3565b803567ffffffffffffffff81116109ff576112879036908301610c17565b8252602081013567ffffffffffffffff81116109ff576112aa9036908301610c17565b60208301908152604082013567ffffffffffffffff81116109ff576112d29036908401610c7f565b91604084019283526112e660608201610c5e565b606085015260808101359067ffffffffffffffff82116109ff5761130c91369101610c17565b60808401528151516001810361137d5750604051903360208301526020825261133660408361090e565b5261137873ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169151610d27565b515290565b7f83b9f0ae0000000000000000000000000000000000000000000000000000000060005260045260246000fdfea164736f6c634300081a000a",
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

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

	error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getCCVs", arg0)

	outstruct := new(GetCCVs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVs(&_EtherSenderReceiver.CallOpts, arg0)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVs(&_EtherSenderReceiver.CallOpts, arg0)
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

type GetCCVs struct {
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}

func (_EtherSenderReceiver *EtherSenderReceiver) Address() common.Address {
	return _EtherSenderReceiver.address
}

type EtherSenderReceiverInterface interface {
	GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

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
