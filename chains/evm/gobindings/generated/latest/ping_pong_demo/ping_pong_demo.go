// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ping_pong_demo

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var PingPongDemoMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCounterpartAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCounterpartChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutOfOrderExecution\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCounterpart\",\"inputs\":[{\"name\":\"counterpartChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"counterpartAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCounterpartAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCounterpartChainSelector\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOutOfOrderExecution\",\"inputs\":[{\"name\":\"outOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaused\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startPingPong\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"OutOfOrderExecutionChange\",\"inputs\":[{\"name\":\"isOutOfOrder\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Ping\",\"inputs\":[{\"name\":\"pingPongCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Pong\",\"inputs\":[{\"name\":\"pingPongCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461012c57604081611a8a803803809161001f828561016c565b83398101031261012c5780516001600160a01b0381169182820361012c5760200151906001600160a01b038216820361012c57821561015657608052331561014557600180546001600160a01b03191633179055600380546001600160a81b031916600892831b610100600160a81b0316179081905560405163095ea7b360e01b815260048101939093526000196024840152602091839160449183916000911c6001600160a01b03165af18015610139576100fc575b6040516118e490816101a682396080518181816103120152818161045001526110c60152f35b6020813d602011610131575b816101156020938361016c565b8101031261012c57518015150361012c57386100d6565b600080fd5b3d9150610108565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b0382119082101761018f57604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461147457508063124c2cad146112b857806316c38b3c14611242578063181f5a77146111c55780631892b906146111355780632874d8bf14610e0e5780632b6e5d6314610dbc57806358ec273914610bbe578063665ed53714610b075780637909b54914610a6e57806379ba50971461098557806385572ffb146103cc5780638da5cb5b1461037a578063ae90de5514610336578063b0f479a1146102c7578063b187bd2614610286578063bee518a41461023d578063ca709a25146101e85763f2fde38b146100f057600080fd5b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e35760043573ffffffffffffffffffffffffffffffffffffffff81168091036101e35761014861188c565b3381146101b957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602073ffffffffffffffffffffffffffffffffffffffff60035460081c16604051908152f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602067ffffffffffffffff60015460a01c16604051908152f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602060ff600354166040519015158152f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602060ff60035460a81c166040519015158152f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e35760043567ffffffffffffffff81116101e35760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101e35773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168033036109575760009160405161048681611561565b81600401358152602482013567ffffffffffffffff8116810361094f576020820152604482013567ffffffffffffffff811161094f576104cc90600436918501016115da565b6040820152606482013567ffffffffffffffff811161094f576104f590600436918501016115da565b916060820192835260848101359067ffffffffffffffff82116109535701903660238301121561094f57600482013567ffffffffffffffff81116109225760208160051b01926105486040519485611599565b818452602060048186019360061b830101019036821161091e57602401915b8183106108c85750505060800152516020818051810103126108c4576020015160ff6003541615610596578280f35b6001810180911161089757600181811603610868577f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f6020604051838152a15b604051906020820152602081526105ee604082611599565b6020916040516105fe8482611599565b84815291839185600354946040516106158161157d565b62030d40815285810160ff8860a81c1615158152604051917f181dcf10000000000000000000000000000000000000000000000000000000008884015251602483015251151560448201526044815261066f606482611599565b6040519661067c88611561565b6106846117b5565b88528688019586526040880192835273ffffffffffffffffffffffffffffffffffffffff606089019160081c1681526080880191825261074967ffffffffffffffff60015460a01c16966107186040519a8b997f96f4e9f9000000000000000000000000000000000000000000000000000000008b5260048b0152604060248b01525160a060448b015260e48a019061164f565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8983030160648a015261164f565b9251927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc878203016084880152878085519283815201940190855b81811061082b57505050859392849273ffffffffffffffffffffffffffffffffffffffff6107e393511660a4850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8483030160c485015261164f565b03925af18015610820576107f5578280f35b813d8311610819575b6108088183611599565b810103126108165781808280f35b80fd5b503d6107fe565b6040513d85823e3d90fd5b8251805173ffffffffffffffffffffffffffffffffffffffff1687528a01518a8701528b998b99508d975060409096019590920191600101610784565b7f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15256020604051838152a16105d6565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b8280fd5b60408336031261091e57604051906108df8261157d565b83359073ffffffffffffffffffffffffffffffffffffffff8216820361091a5782602092604094528286013583820152815201920191610567565b8980fd5b8780fd5b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8480fd5b8580fd5b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e35760005473ffffffffffffffffffffffffffffffffffffffff81163303610a44577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357610aa56116ae565b50610aef6020610afc60405191610abc8184611599565b600083526000368137604051610ad28282611599565b6000815260003681376040519485946060865260608601906116c5565b91848303908501526116c5565b600060408301520390f35b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e3576004358015158091036101e35760207f05a3fef9935c9013a24c6193df2240d34fcf6b0ebf8786b85efe8401d696cdd991610b6f61188c565b6003547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff75ff0000000000000000000000000000000000000000008360a81b16911617600355604051908152a1005b346101e35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357610bf56116ae565b6024359067ffffffffffffffff82116101e357366023830112156101e35781600401359067ffffffffffffffff82116101e35736602483850101116101e357610c3c61188c565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000006001549260a01b16911617600155600090610c9f81610c9a60025461170f565b611762565b81601f8211600114610cff5781908394610ceb9492610cf1575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b60025580f35b602492500101358480610cb9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91845b868110610da15750836001959610610d66575b505050811b0160025580f35b01602401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600384901b60f8161c19169055838080610d5a565b90926020600181926024878701013581550194019101610d47565b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357610e0a610df66117b5565b60405191829160208352602083019061164f565b0390f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357610e4561188c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00600354166003557f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f602060405160018152a1600160405181602082015260208152610eb2604082611599565b602090604051610ec28382611599565b60008152600093611130575b8260035492604051610edf8161157d565b62030d40815282810160ff8660a81c1615158152604051917f181dcf100000000000000000000000000000000000000000000000000000000085840152516024830152511515604482015260448152610f39606482611599565b60405194610f4686611561565b610f4e6117b5565b86528386019283526040860194855273ffffffffffffffffffffffffffffffffffffffff606087019160081c1681526080860191825261101367ffffffffffffffff60015460a01c1693610fe26040519889967f96f4e9f90000000000000000000000000000000000000000000000000000000088526004880152604060248801525160a0604488015260e487019061164f565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc86830301606487015261164f565b9451947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc848203016084850152848087519283815201960190895b8181106110f657505050936110ac9173ffffffffffffffffffffffffffffffffffffffff849596511660a4850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8483030160c485015261164f565b03818673ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af18015610820576107f5578280f35b8251805173ffffffffffffffffffffffffffffffffffffffff1689528701518789015260409097019689968996509092019160010161104e565b610ece565b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e35761116c6116ae565b61117461188c565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000006001549260a01b16911617600155600080f35b346101e35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357610e0a60408051906112068183611599565b601282527f50696e67506f6e6744656d6f20312e352e30000000000000000000000000000060208301525191829160208352602083019061164f565b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e3576004358015158091036101e35761128661188c565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060035416911617600355600080f35b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e35760043567ffffffffffffffff81116101e3576113079036906004016115da565b61130f61188c565b805167ffffffffffffffff81116114455761132f81610c9a60025461170f565b602091601f821160011461138b5761137b92600091836113805750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b600255005b015190508380610cb9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9160005b85811061142d575083600195106113f6575b505050811b01600255005b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558280806113eb565b919260206001819286850151815501940192016113d9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101e35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101e357600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101e357817f85572ffb0000000000000000000000000000000000000000000000000000000060209314908115611537575b811561150d575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611506565b7f7909b54900000000000000000000000000000000000000000000000000000000811491506114ff565b60a0810190811067ffffffffffffffff82111761144557604052565b6040810190811067ffffffffffffffff82111761144557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761144557604052565b81601f820112156101e35780359067ffffffffffffffff8211611445576040519261162d601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185611599565b828452602083830101116101e357816000926020809301838601378301015290565b919082519283825260005b8481106116995750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161165a565b6004359067ffffffffffffffff821682036101e357565b906020808351928381520192019060005b8181106116e35750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016116d6565b90600182811c92168015611758575b602083101461172957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161171e565b601f811161176e575050565b60026000526020600020906020601f840160051c830193106117ab575b601f0160051c01905b81811061179f575050565b60008155600101611794565b909150819061178b565b60405190600082600254916117c98361170f565b808352926001811690811561184f57506001146117ef575b6117ed92500383611599565b565b506002600090815290917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b8183106118335750509060206117ed928201016117e1565b602091935080600191548385890101520191019091849261181b565b602092506117ed9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016117e1565b73ffffffffffffffffffffffffffffffffffffffff6001541633036118ad57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
}

var PingPongDemoABI = PingPongDemoMetaData.ABI

var PingPongDemoBin = PingPongDemoMetaData.Bin

func DeployPingPongDemo(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address) (common.Address, *types.Transaction, *PingPongDemo, error) {
	parsed, err := PingPongDemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongDemoBin), backend, router, feeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPongDemo{address: address, abi: *parsed, PingPongDemoCaller: PingPongDemoCaller{contract: contract}, PingPongDemoTransactor: PingPongDemoTransactor{contract: contract}, PingPongDemoFilterer: PingPongDemoFilterer{contract: contract}}, nil
}

type PingPongDemo struct {
	address common.Address
	abi     abi.ABI
	PingPongDemoCaller
	PingPongDemoTransactor
	PingPongDemoFilterer
}

type PingPongDemoCaller struct {
	contract *bind.BoundContract
}

type PingPongDemoTransactor struct {
	contract *bind.BoundContract
}

type PingPongDemoFilterer struct {
	contract *bind.BoundContract
}

type PingPongDemoSession struct {
	Contract     *PingPongDemo
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type PingPongDemoCallerSession struct {
	Contract *PingPongDemoCaller
	CallOpts bind.CallOpts
}

type PingPongDemoTransactorSession struct {
	Contract     *PingPongDemoTransactor
	TransactOpts bind.TransactOpts
}

type PingPongDemoRaw struct {
	Contract *PingPongDemo
}

type PingPongDemoCallerRaw struct {
	Contract *PingPongDemoCaller
}

type PingPongDemoTransactorRaw struct {
	Contract *PingPongDemoTransactor
}

func NewPingPongDemo(address common.Address, backend bind.ContractBackend) (*PingPongDemo, error) {
	abi, err := abi.JSON(strings.NewReader(PingPongDemoABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindPingPongDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPongDemo{address: address, abi: abi, PingPongDemoCaller: PingPongDemoCaller{contract: contract}, PingPongDemoTransactor: PingPongDemoTransactor{contract: contract}, PingPongDemoFilterer: PingPongDemoFilterer{contract: contract}}, nil
}

func NewPingPongDemoCaller(address common.Address, caller bind.ContractCaller) (*PingPongDemoCaller, error) {
	contract, err := bindPingPongDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoCaller{contract: contract}, nil
}

func NewPingPongDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongDemoTransactor, error) {
	contract, err := bindPingPongDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoTransactor{contract: contract}, nil
}

func NewPingPongDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongDemoFilterer, error) {
	contract, err := bindPingPongDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoFilterer{contract: contract}, nil
}

func bindPingPongDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PingPongDemoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_PingPongDemo *PingPongDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongDemo.Contract.PingPongDemoCaller.contract.Call(opts, result, method, params...)
}

func (_PingPongDemo *PingPongDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.Contract.PingPongDemoTransactor.contract.Transfer(opts)
}

func (_PingPongDemo *PingPongDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongDemo.Contract.PingPongDemoTransactor.contract.Transact(opts, method, params...)
}

func (_PingPongDemo *PingPongDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongDemo.Contract.contract.Call(opts, result, method, params...)
}

func (_PingPongDemo *PingPongDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.Contract.contract.Transfer(opts)
}

func (_PingPongDemo *PingPongDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongDemo.Contract.contract.Transact(opts, method, params...)
}

func (_PingPongDemo *PingPongDemoCaller) GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

	error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCCVs", arg0)

	outstruct := new(GetCCVs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_PingPongDemo *PingPongDemoSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _PingPongDemo.Contract.GetCCVs(&_PingPongDemo.CallOpts, arg0)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _PingPongDemo.Contract.GetCCVs(&_PingPongDemo.CallOpts, arg0)
}

func (_PingPongDemo *PingPongDemoCaller) GetCounterpartAddress(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCounterpartAddress")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetCounterpartAddress() ([]byte, error) {
	return _PingPongDemo.Contract.GetCounterpartAddress(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCounterpartAddress() ([]byte, error) {
	return _PingPongDemo.Contract.GetCounterpartAddress(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCounterpartChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetCounterpartChainSelector() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainSelector(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCounterpartChainSelector() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainSelector(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetFeeToken() (common.Address, error) {
	return _PingPongDemo.Contract.GetFeeToken(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetFeeToken() (common.Address, error) {
	return _PingPongDemo.Contract.GetFeeToken(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetOutOfOrderExecution(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getOutOfOrderExecution")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetOutOfOrderExecution() (bool, error) {
	return _PingPongDemo.Contract.GetOutOfOrderExecution(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetOutOfOrderExecution() (bool, error) {
	return _PingPongDemo.Contract.GetOutOfOrderExecution(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetRouter() (common.Address, error) {
	return _PingPongDemo.Contract.GetRouter(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetRouter() (common.Address, error) {
	return _PingPongDemo.Contract.GetRouter(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) IsPaused() (bool, error) {
	return _PingPongDemo.Contract.IsPaused(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) IsPaused() (bool, error) {
	return _PingPongDemo.Contract.IsPaused(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) Owner() (common.Address, error) {
	return _PingPongDemo.Contract.Owner(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) Owner() (common.Address, error) {
	return _PingPongDemo.Contract.Owner(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PingPongDemo.Contract.SupportsInterface(&_PingPongDemo.CallOpts, interfaceId)
}

func (_PingPongDemo *PingPongDemoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PingPongDemo.Contract.SupportsInterface(&_PingPongDemo.CallOpts, interfaceId)
}

func (_PingPongDemo *PingPongDemoCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) TypeAndVersion() (string, error) {
	return _PingPongDemo.Contract.TypeAndVersion(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) TypeAndVersion() (string, error) {
	return _PingPongDemo.Contract.TypeAndVersion(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "acceptOwnership")
}

func (_PingPongDemo *PingPongDemoSession) AcceptOwnership() (*types.Transaction, error) {
	return _PingPongDemo.Contract.AcceptOwnership(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PingPongDemo.Contract.AcceptOwnership(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "ccipReceive", message)
}

func (_PingPongDemo *PingPongDemoSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.Contract.CcipReceive(&_PingPongDemo.TransactOpts, message)
}

func (_PingPongDemo *PingPongDemoTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.Contract.CcipReceive(&_PingPongDemo.TransactOpts, message)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress []byte) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpart", counterpartChainSelector, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress []byte) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress []byte) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpartAddress(opts *bind.TransactOpts, addr []byte) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpartAddress", addr)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpartAddress(addr []byte) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartAddress(&_PingPongDemo.TransactOpts, addr)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpartAddress(addr []byte) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartAddress(&_PingPongDemo.TransactOpts, addr)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpartChainSelector", chainSelector)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainSelector(&_PingPongDemo.TransactOpts, chainSelector)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainSelector(&_PingPongDemo.TransactOpts, chainSelector)
}

func (_PingPongDemo *PingPongDemoTransactor) SetOutOfOrderExecution(opts *bind.TransactOpts, outOfOrderExecution bool) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setOutOfOrderExecution", outOfOrderExecution)
}

func (_PingPongDemo *PingPongDemoSession) SetOutOfOrderExecution(outOfOrderExecution bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetOutOfOrderExecution(&_PingPongDemo.TransactOpts, outOfOrderExecution)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetOutOfOrderExecution(outOfOrderExecution bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetOutOfOrderExecution(&_PingPongDemo.TransactOpts, outOfOrderExecution)
}

func (_PingPongDemo *PingPongDemoTransactor) SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setPaused", pause)
}

func (_PingPongDemo *PingPongDemoSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetPaused(&_PingPongDemo.TransactOpts, pause)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetPaused(&_PingPongDemo.TransactOpts, pause)
}

func (_PingPongDemo *PingPongDemoTransactor) StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "startPingPong")
}

func (_PingPongDemo *PingPongDemoSession) StartPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.StartPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactorSession) StartPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.StartPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "transferOwnership", to)
}

func (_PingPongDemo *PingPongDemoSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.TransferOwnership(&_PingPongDemo.TransactOpts, to)
}

func (_PingPongDemo *PingPongDemoTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.TransferOwnership(&_PingPongDemo.TransactOpts, to)
}

type PingPongDemoOutOfOrderExecutionChangeIterator struct {
	Event *PingPongDemoOutOfOrderExecutionChange

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoOutOfOrderExecutionChangeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoOutOfOrderExecutionChange)
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
		it.Event = new(PingPongDemoOutOfOrderExecutionChange)
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

func (it *PingPongDemoOutOfOrderExecutionChangeIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoOutOfOrderExecutionChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoOutOfOrderExecutionChange struct {
	IsOutOfOrder bool
	Raw          types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterOutOfOrderExecutionChange(opts *bind.FilterOpts) (*PingPongDemoOutOfOrderExecutionChangeIterator, error) {

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "OutOfOrderExecutionChange")
	if err != nil {
		return nil, err
	}
	return &PingPongDemoOutOfOrderExecutionChangeIterator{contract: _PingPongDemo.contract, event: "OutOfOrderExecutionChange", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchOutOfOrderExecutionChange(opts *bind.WatchOpts, sink chan<- *PingPongDemoOutOfOrderExecutionChange) (event.Subscription, error) {

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "OutOfOrderExecutionChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoOutOfOrderExecutionChange)
				if err := _PingPongDemo.contract.UnpackLog(event, "OutOfOrderExecutionChange", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseOutOfOrderExecutionChange(log types.Log) (*PingPongDemoOutOfOrderExecutionChange, error) {
	event := new(PingPongDemoOutOfOrderExecutionChange)
	if err := _PingPongDemo.contract.UnpackLog(event, "OutOfOrderExecutionChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoOwnershipTransferRequestedIterator struct {
	Event *PingPongDemoOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoOwnershipTransferRequested)
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
		it.Event = new(PingPongDemoOwnershipTransferRequested)
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

func (it *PingPongDemoOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoOwnershipTransferRequestedIterator{contract: _PingPongDemo.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoOwnershipTransferRequested)
				if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseOwnershipTransferRequested(log types.Log) (*PingPongDemoOwnershipTransferRequested, error) {
	event := new(PingPongDemoOwnershipTransferRequested)
	if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoOwnershipTransferredIterator struct {
	Event *PingPongDemoOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoOwnershipTransferred)
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
		it.Event = new(PingPongDemoOwnershipTransferred)
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

func (it *PingPongDemoOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoOwnershipTransferredIterator{contract: _PingPongDemo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoOwnershipTransferred)
				if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseOwnershipTransferred(log types.Log) (*PingPongDemoOwnershipTransferred, error) {
	event := new(PingPongDemoOwnershipTransferred)
	if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoPingIterator struct {
	Event *PingPongDemoPing

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoPingIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoPing)
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
		it.Event = new(PingPongDemoPing)
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

func (it *PingPongDemoPingIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoPing struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterPing(opts *bind.FilterOpts) (*PingPongDemoPingIterator, error) {

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &PingPongDemoPingIterator{contract: _PingPongDemo.contract, event: "Ping", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongDemoPing) (event.Subscription, error) {

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoPing)
				if err := _PingPongDemo.contract.UnpackLog(event, "Ping", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParsePing(log types.Log) (*PingPongDemoPing, error) {
	event := new(PingPongDemoPing)
	if err := _PingPongDemo.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoPongIterator struct {
	Event *PingPongDemoPong

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoPongIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoPong)
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
		it.Event = new(PingPongDemoPong)
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

func (it *PingPongDemoPongIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoPong struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterPong(opts *bind.FilterOpts) (*PingPongDemoPongIterator, error) {

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &PingPongDemoPongIterator{contract: _PingPongDemo.contract, event: "Pong", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *PingPongDemoPong) (event.Subscription, error) {

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoPong)
				if err := _PingPongDemo.contract.UnpackLog(event, "Pong", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParsePong(log types.Log) (*PingPongDemoPong, error) {
	event := new(PingPongDemoPong)
	if err := _PingPongDemo.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCCVs struct {
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}

func (_PingPongDemo *PingPongDemo) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _PingPongDemo.abi.Events["OutOfOrderExecutionChange"].ID:
		return _PingPongDemo.ParseOutOfOrderExecutionChange(log)
	case _PingPongDemo.abi.Events["OwnershipTransferRequested"].ID:
		return _PingPongDemo.ParseOwnershipTransferRequested(log)
	case _PingPongDemo.abi.Events["OwnershipTransferred"].ID:
		return _PingPongDemo.ParseOwnershipTransferred(log)
	case _PingPongDemo.abi.Events["Ping"].ID:
		return _PingPongDemo.ParsePing(log)
	case _PingPongDemo.abi.Events["Pong"].ID:
		return _PingPongDemo.ParsePong(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (PingPongDemoOutOfOrderExecutionChange) Topic() common.Hash {
	return common.HexToHash("0x05a3fef9935c9013a24c6193df2240d34fcf6b0ebf8786b85efe8401d696cdd9")
}

func (PingPongDemoOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (PingPongDemoOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (PingPongDemoPing) Topic() common.Hash {
	return common.HexToHash("0x48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f")
}

func (PingPongDemoPong) Topic() common.Hash {
	return common.HexToHash("0x58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b1525")
}

func (_PingPongDemo *PingPongDemo) Address() common.Address {
	return _PingPongDemo.address
}

type PingPongDemoInterface interface {
	GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

		error)

	GetCounterpartAddress(opts *bind.CallOpts) ([]byte, error)

	GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error)

	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetOutOfOrderExecution(opts *bind.CallOpts) (bool, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress []byte) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr []byte) (*types.Transaction, error)

	SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	SetOutOfOrderExecution(opts *bind.TransactOpts, outOfOrderExecution bool) (*types.Transaction, error)

	SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error)

	StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOutOfOrderExecutionChange(opts *bind.FilterOpts) (*PingPongDemoOutOfOrderExecutionChangeIterator, error)

	WatchOutOfOrderExecutionChange(opts *bind.WatchOpts, sink chan<- *PingPongDemoOutOfOrderExecutionChange) (event.Subscription, error)

	ParseOutOfOrderExecutionChange(log types.Log) (*PingPongDemoOutOfOrderExecutionChange, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*PingPongDemoOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*PingPongDemoOwnershipTransferred, error)

	FilterPing(opts *bind.FilterOpts) (*PingPongDemoPingIterator, error)

	WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongDemoPing) (event.Subscription, error)

	ParsePing(log types.Log) (*PingPongDemoPing, error)

	FilterPong(opts *bind.FilterOpts) (*PingPongDemoPongIterator, error)

	WatchPong(opts *bind.WatchOpts, sink chan<- *PingPongDemoPong) (event.Subscription, error)

	ParsePong(log types.Log) (*PingPongDemoPong, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
