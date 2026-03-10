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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSend\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getCCVsAndMinBlockDepth\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IWrappedNative\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"gotToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expectedToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAmounts\",\"inputs\":[{\"name\":\"gotAmounts\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenAmountNotEqualToMsgValue\",\"inputs\":[{\"name\":\"gotAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c080604052346101255761002a906115d480380380916100208285610181565b83398101906101ba565b6001600160a01b03811690811561016b5760805260405163e861e90760e01b815290602082600481845afa908115610132576044602092600094859161013e575b506001600160a01b031660a081905260405163095ea7b360e01b8152600481019390935284196024840152919384928391905af18015610132576100f5575b6040516113fa90816101da82396080518181816101ac01528181610271015281816107230152610ee0015260a0518181816103bb0152818161069201528181610e6701526113940152f35b6020813d60201161012a575b8161010e60209383610181565b8101031261012557518015150361012557386100aa565b600080fd5b3d9150610101565b6040513d6000823e3d90fd5b61015e9150843d8611610164575b6101568183610181565b8101906101ba565b3861006b565b503d61014c565b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b038211908210176101a457604052565b634e487b7160e01b600052604160045260246000fd5b9081602091031261012557516001600160a01b0381168103610125579056fe608080604052600436101561001d575b50361561001b57600080fd5b005b600090813560e01c90816301ffc9a71461081757508063181f5a771461079457806320487ded146106b65780634dbe7e921461064757806385572ffb146101ec57806396f4e9f9146101d0578063b0f479a1146101615763bcd7c1110361000f57346101565760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610156576100b5610a2b565b5060243567ffffffffffffffff811161015d573660238201121561015d57806004013567ffffffffffffffff81116101595736910160240111610156576040516101399161014660206101088185610951565b828452600036813760405161011d8282610951565b8381526000368137604051958695608087526080870190610acc565b9185830390860152610acc565b9080604084015260608301520390f35b80fd5b8280fd5b5080fd5b503461015657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60206101e46101de36610a47565b90610dbe565b604051908152f35b50346101565760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101565760043567ffffffffffffffff811161015d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261015d5773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361061b576040516102a381610906565b81600401358152602482013567ffffffffffffffff81168103610573576020820152604482013567ffffffffffffffff8111610573576102e99060043691850101610c5a565b6040820152606482013567ffffffffffffffff8111610573576103129060043691850101610c5a565b9160608201928352608481013567ffffffffffffffff81116106175760809160046103409236920101610cc2565b9101918183525160208180518101031261057357602001519073ffffffffffffffffffffffffffffffffffffffff82168092036105735751600181036105ec575073ffffffffffffffffffffffffffffffffffffffff6103a08351610d6a565b5151169173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680930361059b576103ee60209151610d6a565b51015190823b15610573576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152826004820152848160248183885af180156105905761057c575b508380808085855af13d15610577573d61045281610992565b906104606040519283610951565b81528560203d92013e5b15610473578380f35b823b1561057357836040517fd0e30db0000000000000000000000000000000000000000000000000000000008152818160048187895af1801561054f5761055a575b506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482019290925291602091839160449183915af1801561054f57610520575b80808380f35b6105419060203d602011610548575b6105398183610951565b810190610da6565b503861051a565b503d61052f565b6040513d84823e3d90fd5b8161056791949394610951565b610573579083386104b5565b8380fd5b61046a565b8461058991959295610951565b9238610439565b6040513d87823e3d90fd5b838373ffffffffffffffffffffffffffffffffffffffff6105be60449451610d6a565b5151167f0fc746a1000000000000000000000000000000000000000000000000000000008352600452602452fd5b7f83b9f0ae000000000000000000000000000000000000000000000000000000008452600452602483fd5b8480fd5b6024827fd7f7333400000000000000000000000000000000000000000000000000000000815233600452fd5b503461015657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101565761070a60206106d56106cd36610a47565b91909161126e565b9060405193849283927f20487ded00000000000000000000000000000000000000000000000000000000845260048401610b16565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561054f57829161075e575b602082604051908152f35b90506020813d60201161078c575b8161077960209383610951565b8101031261015d57602091505138610753565b3d915061076c565b503461015657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015657506108136040516107d5604082610951565b601981527f457468657253656e646572526563656976657220312e352e300000000000000060208201526040519182916020835260208301906109cc565b0390f35b90503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361015957602092507f85572ffb0000000000000000000000000000000000000000000000000000000081149081156108dc575b81156108b2575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386108ab565b7fbcd7c11100000000000000000000000000000000000000000000000000000000811491506108a4565b60a0810190811067ffffffffffffffff82111761092257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761092257604052565b67ffffffffffffffff811161092257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110610a165750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016109d7565b6004359067ffffffffffffffff82168203610a4257565b600080fd5b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112610a425760043567ffffffffffffffff81168103610a4257916024359067ffffffffffffffff8211610a42577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a092030112610a425760040190565b906020808351928381520192019060005b818110610aea5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610add565b9067ffffffffffffffff9093929316815260406020820152610b7b610b47845160a0604085015260e08401906109cc565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030160608501526109cc565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b818110610c225750505060808473ffffffffffffffffffffffffffffffffffffffff6060610c1f969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828503019101526109cc565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101610bbc565b81601f82011215610a4257803590610c7182610992565b92610c7f6040519485610951565b82845260208383010111610a4257816000926020809301838601378301015290565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a4257565b81601f82011215610a425780359067ffffffffffffffff82116109225760208260051b0192610cf46040519485610951565b82845260208085019360061b83010191818311610a4257602001925b828410610d1e575050505090565b604084830312610a425760405190604082019082821067ffffffffffffffff831117610922576040926020928452610d5587610ca1565b81528287013583820152815201930192610d10565b805115610d775760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90816020910312610a4257518015158103610a425790565b60009160408101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610573578101803567ffffffffffffffff8111610617578060061b3603602083011361061757156112415760400135606082013573ffffffffffffffffffffffffffffffffffffffff81168091036106175761120b575b50610e4f9061126e565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906020610e976040850151610d6a565b510151823b156106175784600491604051928380927fd0e30db0000000000000000000000000000000000000000000000000000000008252875af18015610590576111f7575b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff811692604051917f20487ded00000000000000000000000000000000000000000000000000000000835260208380610f50898860048401610b16565b0381885afa9283156111ec5787936111b8575b506060860173ffffffffffffffffffffffffffffffffffffffff8151168061101757505050505091602091610fc893476040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610b16565b03925af191821561100b5791610fdc575090565b90506020813d602011611003575b81610ff760209383610951565b81010312610a42575190565b3d9150610fea565b604051903d90823e3d90fd5b602089604051828101907f23b872dd00000000000000000000000000000000000000000000000000000000825233602482015230604482015288606482015260648152611065608482610951565b519082855af1156111ad5788513d6111a45750803b155b61117957505173ffffffffffffffffffffffffffffffffffffffff169182036110e1575b50505091602091610fc893856040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401610b16565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101929092526020908290604490829089905af18015610590579160209391610fc8959361115c575b819395508294506110a0565b61117290853d8711610548576105398183610951565b5038611150565b7f5274afe7000000000000000000000000000000000000000000000000000000008952600452602488fd5b6001141561107c565b6040513d8a823e3d90fd5b9092506020813d6020116111e4575b816111d460209383610951565b81010312610a4257519138610f63565b3d91506111c7565b6040513d89823e3d90fd5b8461120491959295610951565b9238610edd565b348114610e45577fba2f746700000000000000000000000000000000000000000000000000000000845260045234602452604483fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b6060608060405161127e81610906565b828152826020820152826040820152600083820152015260a081360312610a4257604051906112ac82610906565b803567ffffffffffffffff8111610a42576112ca9036908301610c5a565b8252602081013567ffffffffffffffff8111610a42576112ed9036908301610c5a565b60208301908152604082013567ffffffffffffffff8111610a42576113159036908401610cc2565b916040840192835261132960608201610ca1565b606085015260808101359067ffffffffffffffff8211610a425761134f91369101610c5a565b6080840152815151600181036113c057506040519033602083015260208252611379604083610951565b526113bb73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169151610d6a565b515290565b7f83b9f0ae0000000000000000000000000000000000000000000000000000000060005260045260246000fdfea164736f6c634300081a000a",
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

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetCCVsAndMinBlockDepth(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockDepth,

	error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getCCVsAndMinBlockDepth", arg0, arg1)

	outstruct := new(GetCCVsAndMinBlockDepth)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MinBlockDepth = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetCCVsAndMinBlockDepth(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockDepth,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVsAndMinBlockDepth(&_EtherSenderReceiver.CallOpts, arg0, arg1)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetCCVsAndMinBlockDepth(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockDepth,

	error) {
	return _EtherSenderReceiver.Contract.GetCCVsAndMinBlockDepth(&_EtherSenderReceiver.CallOpts, arg0, arg1)
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

type GetCCVsAndMinBlockDepth struct {
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
	MinBlockDepth     uint16
}

func (_EtherSenderReceiver *EtherSenderReceiver) Address() common.Address {
	return _EtherSenderReceiver.address
}

type EtherSenderReceiverInterface interface {
	GetCCVsAndMinBlockDepth(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockDepth,

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
