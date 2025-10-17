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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSend\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IWrappedNative\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"gotToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expectedToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAmounts\",\"inputs\":[{\"name\":\"gotAmounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenAmountNotEqualToMsgValue\",\"inputs\":[{\"name\":\"gotAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c080604052346101245761002a9061174780380380916100208285610180565b83398101906101b9565b6001600160a01b03811690811561016a5760805260405163e861e90760e01b815290602082600481845afa908115610131576044602092600094859161013d575b506001600160a01b031660a081905260405163095ea7b360e01b8152600481019390935284196024840152919384928391905af18015610131576100f4575b60405161156e90816101d9823960805181818160c801528181610190015281816106b80152610e78015260a0518181816102da0152818161062701528181610e0001526114080152f35b6020813d602011610129575b8161010d60209383610180565b8101031261012457518015150361012457386100aa565b600080fd5b3d9150610100565b6040513d6000823e3d90fd5b61015d9150843d8611610163575b6101558183610180565b8101906101b9565b3861006b565b503d61014b565b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b038211908210176101a357604052565b634e487b7160e01b600052604160045260246000fd5b9081602091031261012457516001600160a01b0381168103610124579056fe608080604052600436101561001d575b50361561001b57600080fd5b005b600090813560e01c90816301ffc9a7146107ac57508063181f5a771461072957806320487ded1461064b5780634dbe7e92146105dc5780637909b5491461054657806385572ffb1461010b57806396f4e9f9146100ef5763b0f479a10361000f57346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b80fd5b60206101036100fd366109e0565b90610d57565b604051908152f35b50346100ec5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec5760043567ffffffffffffffff81116105425760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126105425773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163303610516576040516101c28161089f565b81600401358152602482013567ffffffffffffffff81168103610473576020820152604482013567ffffffffffffffff8111610473576102089060043691850101610bf3565b6040820152606482013567ffffffffffffffff8111610473576102319060043691850101610bf3565b9160608201928352608481013567ffffffffffffffff811161051257608091600461025f9236920101610c5b565b9101918183525160208180518101031261047357602001519073ffffffffffffffffffffffffffffffffffffffff82168092036104735751600181036104e7575073ffffffffffffffffffffffffffffffffffffffff6102bf8351610d03565b5151169173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168093036104965761030d60209151610d03565b51015190823b15610473576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152826004820152848160248183885af1801561048b57610477575b508380808085855af1610369611461565b5015610373578380f35b823b1561047357836040517fd0e30db0000000000000000000000000000000000000000000000000000000008152818160048187895af1801561044f5761045a575b506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482019290925291602091839160449183915af1801561044f57610420575b80808380f35b6104419060203d602011610448575b61043981836108ea565b810190610d3f565b503861041a565b503d61042f565b6040513d84823e3d90fd5b81610467919493946108ea565b610473579083386103b5565b8380fd5b84610484919592956108ea565b9238610358565b6040513d87823e3d90fd5b838373ffffffffffffffffffffffffffffffffffffffff6104b960449451610d03565b5151167f0fc746a1000000000000000000000000000000000000000000000000000000008352600452602452fd5b7f83b9f0ae000000000000000000000000000000000000000000000000000000008452600452602483fd5b8480fd5b6024827fd7f7333400000000000000000000000000000000000000000000000000000000815233600452fd5b5080fd5b50346100ec5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec576105c5906105826109c4565b50604051906105d2602061059681856108ea565b828452823681376040516105aa82826108ea565b83815283368137604051958695606087526060870190610a65565b9185830390860152610a65565b9060408301520390f35b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100ec5761069f602061066a610662366109e0565b9190916112e2565b9060405193849283927f20487ded00000000000000000000000000000000000000000000000000000000845260048401610aaf565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561044f5782916106f3575b602082604051908152f35b90506020813d602011610721575b8161070e602093836108ea565b81010312610542576020915051386106e8565b3d9150610701565b50346100ec57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ec57506107a860405161076a6040826108ea565b601981527f457468657253656e646572526563656976657220312e352e30000000000000006020820152604051918291602083526020830190610965565b0390f35b9050346105425760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610542576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361089b57602092507f85572ffb000000000000000000000000000000000000000000000000000000008114908115610871575b8115610847575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610840565b7f7909b5490000000000000000000000000000000000000000000000000000000081149150610839565b8280fd5b60a0810190811067ffffffffffffffff8211176108bb57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108bb57604052565b67ffffffffffffffff81116108bb57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106109af5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610970565b6004359067ffffffffffffffff821682036109db57565b600080fd5b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126109db5760043567ffffffffffffffff811681036109db57916024359067ffffffffffffffff82116109db577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126109db5760040190565b906020808351928381520192019060005b818110610a835750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610a76565b9067ffffffffffffffff9093929316815260406020820152610b14610ae0845160a0604085015260e0840190610965565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152610965565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b818110610bbb5750505060808473ffffffffffffffffffffffffffffffffffffffff6060610bb8969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152610965565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101610b55565b81601f820112156109db57803590610c0a8261092b565b92610c1860405194856108ea565b828452602083830101116109db57816000926020809301838601378301015290565b359073ffffffffffffffffffffffffffffffffffffffff821682036109db57565b81601f820112156109db5780359067ffffffffffffffff82116108bb5760208260051b0192610c8d60405194856108ea565b82845260208085019360061b830101918183116109db57602001925b828410610cb7575050505090565b6040848303126109db5760405190604082019082821067ffffffffffffffff8311176108bb576040926020928452610cee87610c3a565b81528287013583820152815201930192610ca9565b805115610d105760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b908160209103126109db575180151581036109db5790565b9060009060408101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561089b578101803567ffffffffffffffff8111610473578060061b3603602083011361047357156112b55760400135606082013573ffffffffffffffffffffffffffffffffffffffff81168091036104735761127f575b50610de9906112e2565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166020610e2f6040840151610d03565b510151813b156104735783600491604051928380927fd0e30db0000000000000000000000000000000000000000000000000000000008252865af1801561127457611260575b507f00000000000000000000000000000000000000000000000000000000000000009373ffffffffffffffffffffffffffffffffffffffff851691604051907f20487ded00000000000000000000000000000000000000000000000000000000825260208280610ee9888760048401610aaf565b0381875afa918215611255578692611221575b50606085019673ffffffffffffffffffffffffffffffffffffffff8851169788610fb45750505050610f659394509060209291476040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610aaf565b03925af1918215610fa85791610f79575090565b90506020813d602011610fa0575b81610f94602093836108ea565b810103126109db575190565b3d9150610f87565b604051903d90823e3d90fd5b611054604051602081019a7f23b872dd000000000000000000000000000000000000000000000000000000008c52336024830152306044830152866064830152606482526110036084836108ea565b8a8060409d8e94611016865196876108ea565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af161104e611461565b91611491565b805180611180575b50505173ffffffffffffffffffffffffffffffffffffffff169182036110dd575b5050509160209184936110be9587518097819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610aaf565b03925af19283156110d3575091610f79575090565b51903d90823e3d90fd5b87517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101929092526020908290604490829089905af180156111765791602093916110be969593611159575b819395965082945061107d565b61116f90853d87116104485761043981836108ea565b503861114c565b86513d87823e3d90fd5b90602080611192938301019101610d3f565b1561119e57388061105c565b608489517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9091506020813d60201161124d575b8161123d602093836108ea565b810103126109db57519038610efc565b3d9150611230565b6040513d88823e3d90fd5b8361126d919492946108ea565b9138610e75565b6040513d86823e3d90fd5b348114610ddf577fba2f746700000000000000000000000000000000000000000000000000000000835260045234602452604482fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b606060806040516112f28161089f565b828152826020820152826040820152600083820152015260a0813603126109db57604051906113208261089f565b803567ffffffffffffffff81116109db5761133e9036908301610bf3565b8252602081013567ffffffffffffffff81116109db576113619036908301610bf3565b60208301908152604082013567ffffffffffffffff81116109db576113899036908401610c5b565b916040840192835261139d60608201610c3a565b606085015260808101359067ffffffffffffffff82116109db576113c391369101610bf3565b608084015281515160018103611434575060405190336020830152602082526113ed6040836108ea565b5261142f73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169151610d03565b515290565b7f83b9f0ae0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3d1561148c573d906114728261092b565b9161148060405193846108ea565b82523d6000602084013e565b606090565b9192901561150c57508151156114a5575090565b3b156114ae5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561151f5750805190602001fd5b61155d906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190610965565b0390fdfea164736f6c634300081a000a",
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
