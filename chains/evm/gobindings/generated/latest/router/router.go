// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

type RouterOffRamp struct {
	SourceChainSelector uint64
	OffRamp             common.Address
}

type RouterOnRamp struct {
	DestChainSelector uint64
	OnRamp            common.Address
}

var RouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"wrappedNative\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"armProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_RET_BYTES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRampUpdates\",\"inputs\":[{\"name\":\"onRampUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct Router.OnRamp[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRamp\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offRampRemoves\",\"type\":\"tuple[]\",\"internalType\":\"struct Router.OffRamp[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offRampAdds\",\"type\":\"tuple[]\",\"internalType\":\"struct Router.OffRamp[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSend\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getArmProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOffRamps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Router.OffRamp[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOnRamp\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWrappedNative\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isChainSupported\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOffRamp\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routeMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"retData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWrappedNative\",\"inputs\":[{\"name\":\"wrappedNative\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessageExecuted\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"calldataHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OffRampAdded\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OffRampRemoved\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnRampSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"onRamp\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadARMSignal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToSendValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRecipientAddress\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OffRampMismatch\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyOffRamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedDestinationChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]",
	Bin: "0x60a0346100f557601f6129bc38819003918201601f19168301916001600160401b038311848410176100fa5780849260409485528339810103126100f557610052602061004b83610110565b9201610110565b9033156100b057600080546001600160a01b03199081163317909155600280549091166001600160a01b03929092169190911790556080526040516128979081610125823960805181818161083a01528181610be70152611d180152f35b60405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036100f55756fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461013757806320487ded146101325780633cf979831461012d5780635246492f1461012857806352cb60ca146101235780635f3e849f1461011e578063787350e31461011957806379ba50971461011457806383826b2b1461010f5780638da5cb5b1461010a57806396f4e9f914610105578063a40e69c714610100578063a48a9058146100fb578063a8d87a3b146100f6578063da5fcac8146100f1578063e861e907146100ec578063f2fde38b146100e75763fbca3b74146100e257600080fd5b61190e565b6117ca565b611778565b6113ff565b61135b565b6112e8565b611197565b610b9b565b610b49565b610ad7565b61097c565b610942565b6108e5565b61085e565b6107ef565b610744565b610569565b610296565b600091031261014757565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761019757604052565b61014c565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761019757604052565b604051906101ec60a08361019c565b565b604051906101ec60408361019c565b67ffffffffffffffff811161019757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106102815750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610242565b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475761031360408051906102d7818361019c565b600c82527f526f7574657220312e322e300000000000000000000000000000000000000000602083015251918291602083526020830190610237565b0390f35b67ffffffffffffffff81160361014757565b81601f8201121561014757803590610340826101fd565b9261034e604051948561019c565b8284526020838301011161014757816000926020809301838601378301015290565b67ffffffffffffffff81116101975760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361014757565b35906101ec82610388565b81601f82011215610147578035906103c882610370565b926103d6604051948561019c565b82845260208085019360061b8301019181831161014757602001925b828410610400575050505090565b604084830312610147576020604091825161041a8161017b565b863561042581610388565b815282870135838201528152019301926103f2565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101475760043561047181610317565b9160243567ffffffffffffffff81116101475760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8284030112610147576104b86101dd565b91816004013567ffffffffffffffff8111610147578160046104dc92850101610329565b8352602482013567ffffffffffffffff81116101475781600461050192850101610329565b6020840152604482013567ffffffffffffffff811161014757816004610529928501016103b1565b604084015261053a606483016103a6565b606084015260848201359167ffffffffffffffff8311610147576105619201600401610329565b608082015290565b34610147576105773661043a565b6060810173ffffffffffffffffffffffffffffffffffffffff6105ae825173ffffffffffffffffffffffffffffffffffffffff1690565b16156106dd575b5073ffffffffffffffffffffffffffffffffffffffff6106066105ec8467ffffffffffffffff166000526003602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b1680156106a5579060209161064a936040518095819482937f20487ded00000000000000000000000000000000000000000000000000000000845260048401611a87565b03915afa80156106a05761031391600091610671575b506040519081529081906020820190565b610693915060203d602011610699575b61068b818361019c565b8101906119ac565b38610660565b503d610681565b611aa8565b7fae236d9c0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b61071a9061070060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b386105b5565b93929161073f9060409215158652606060208701526060860190610237565b930152565b346101475760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475760043567ffffffffffffffff81116101475760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101475760243561ffff8116810361014757610313916107e09160443590606435926107d884610388565b600401611cd1565b60409391935193849384610720565b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475773ffffffffffffffffffffffffffffffffffffffff6004356108ae81610388565b6108b66123e2565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006002541617600255600080f35b346101475760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475761094060043561092381610388565b60243561092f81610388565b6044359161093b6123e2565b611eb3565b005b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757602060405160848152f35b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475773ffffffffffffffffffffffffffffffffffffffff600154163303610a795760005473ffffffffffffffffffffffffffffffffffffffff16600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055610a3a7fffffffffffffffffffffffff000000000000000000000000000000000000000060015416600155565b73ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b346101475760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147576020610b3f610b2c600435610b1a81610317565b60243590610b2782610388565b6124f2565b6000526005602052604060002054151590565b6040519015158152f35b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b610ba43661043a565b6040517f397796f700000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106a057600091611102575b506110d857610c3e6105ec8367ffffffffffffffff166000526003602052604060002090565b73ffffffffffffffffffffffffffffffffffffffff811680156110a0576060830191610c9a610c81845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b610fcd57610cda610cc060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168452565b6040517f20487ded00000000000000000000000000000000000000000000000000000000815260208180610d12888a60048401611a87565b0381865afa9081156106a057600091610fae575b503410610f84573492610d53610c81610c81835173ffffffffffffffffffffffffffffffffffffffff1690565b91823b15610147576000600493604051948580927fd0e30db000000000000000000000000000000000000000000000000000000000825234905af19283156106a057610dc693610f69575b50610dc1610c8134935173ffffffffffffffffffffffffffffffffffffffff1690565b612461565b9190915b604082019160005b83518051821015610ee157610c81610ded83610e0893611fc0565b515173ffffffffffffffffffffffffffffffffffffffff1690565b6040517f48a98aa400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8816600482015273ffffffffffffffffffffffffffffffffffffffff82166024820152909190602081604481885afa80156106a057600193610eab92600092610eb1575b5073ffffffffffffffffffffffffffffffffffffffff6020610e9e868b51611fc0565b510151921690339061253c565b01610dd2565b610ed391925060203d8111610eda575b610ecb818361019c565b810190611fd9565b9038610e7b565b503d610ec1565b610f246020888689600088604051968795869485937fdf0aa9e9000000000000000000000000000000000000000000000000000000008552339260048601611fee565b03925af180156106a05761031391600091610f4a57506040519081529081906020820190565b610f63915060203d6020116106995761068b818361019c565b82610660565b80610f786000610f7e9361019c565b8061013c565b38610d9e565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b610fc7915060203d6020116106995761068b818361019c565b38610d26565b3461107657604051907f20487ded0000000000000000000000000000000000000000000000000000000082526020828061100b888a60048401611a87565b0381865afa9081156106a05761104d92600092611055575b50611046610c8183965173ffffffffffffffffffffffffffffffffffffffff1690565b339061253c565b919091610dca565b61106f91925060203d6020116106995761068b818361019c565b9038611023565b7f1841b4e10000000000000000000000000000000000000000000000000000000060005260046000fd5b7fae236d9c0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b7fc14837150000000000000000000000000000000000000000000000000000000060005260046000fd5b611124915060203d60201161112a575b61111c818361019c565b810190611ab4565b38610c18565b503d611112565b602060408183019282815284518094520192019060005b8181106111555750505090565b8251805167ffffffffffffffff16855260209081015173ffffffffffffffffffffffffffffffffffffffff168186015260409094019390920191600101611148565b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475760405180816020600454928381520160046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9260005b8181106112cf5750506112149250038261019c565b61121e8151612036565b9060005b81518110156112c1578061123860019284611fc0565b516112a573ffffffffffffffffffffffffffffffffffffffff61126b61125e8460a01c90565b67ffffffffffffffff1690565b926112876112776101ee565b67ffffffffffffffff9095168552565b1673ffffffffffffffffffffffffffffffffffffffff166020830152565b6112af8286611fc0565b526112ba8185611fc0565b5001611222565b604051806103138582611131565b84548352600194850194869450602090930192016111ff565b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147576020610b3f60043561132881610317565b67ffffffffffffffff16600052600360205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475767ffffffffffffffff60043561139f81610317565b166000526003602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b9181601f840112156101475782359167ffffffffffffffff8311610147576020808501948460061b01011161014757565b346101475760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475760043567ffffffffffffffff81116101475761144e9036906004016113ce565b60243567ffffffffffffffff81116101475761146e9036906004016113ce565b60443567ffffffffffffffff81116101475761148e9036906004016113ce565b9490936114996123e2565b60005b8181106116495750505060005b8181106115615750505060005b8281106114bf57005b806114d56114d060019386866120b1565b611acc565b6114eb60206114e58488886120b1565b016120fa565b906114fe6114f983836124f2565b6127f9565b61150b575b5050016114b6565b60405173ffffffffffffffffffffffffffffffffffffffff92909216825267ffffffffffffffff16907fa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b9490602090a23880611503565b61156f6114d08284866120b1565b61157f60206114e58486886120b1565b9061159961159561159084846124f2565b612715565b1590565b6115f95760405173ffffffffffffffffffffffffffffffffffffffff9290921682526001929167ffffffffffffffff91909116907fa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb90602090a2016114a9565b7f496477900000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045273ffffffffffffffffffffffffffffffffffffffff1660245260446000fd5b8061165f61165a60019385876120b1565b6120c1565b7f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2367ffffffffffffffff61174d61173260208501946117246116b5875173ffffffffffffffffffffffffffffffffffffffff1690565b6116e46116ca845167ffffffffffffffff1690565b67ffffffffffffffff166000526003602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5167ffffffffffffffff1690565b935173ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff919091168152921691602090a20161149c565b346101475760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475773ffffffffffffffffffffffffffffffffffffffff60043561181a81610388565b6118226123e2565b163381146118b057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff61188860005473ffffffffffffffffffffffffffffffffffffffff1690565b167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475761195160043561194c81610317565b612188565b60405180916020820160208352815180915260206040840192019060005b81811061197d575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161196f565b90816020910312610147575190565b91906119e56119d3845160a0845260a0840190610237565b60208501518382036020850152610237565b9060408401519181810360408301526020808451928381520193019060005b818110611a4f57505050608084611a3c6060611a4c96970151606085019073ffffffffffffffffffffffffffffffffffffffff169052565b0151906080818403910152610237565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101611a04565b60409067ffffffffffffffff611a4c949316815281602082015201906119bb565b6040513d6000823e3d90fd5b90816020910312610147575180151581036101475790565b35611a4c81610317565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561014757016020813591019167ffffffffffffffff821161014757813603831361014757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160209082815201919060005b818110611b7f5750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735611ba881610388565b16815260208781013590820152019401929101611b72565b90602082528035602083015267ffffffffffffffff6020820135611be381610317565b166040830152611c4a611c0d611bfc6040840184611ad6565b60a0606087015260c0860191611b26565b611c1a6060840184611ad6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403016080870152611b26565b9060808101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610147570160208135910167ffffffffffffffff8211610147578160061b36038113610147578360a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611a4c96860301910152611b65565b939190926040517f397796f700000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106a057600091611e94575b506110d85760208501611d6a610b2c8235611d6381610317565b33906124f2565b15611e6a57611dda611e62611e0c7f85572ffb0000000000000000000000000000000000000000000000000000000097611e18967f9b877de93ea9895756e337442c657f95a34fc68e7eb988bdfa693d5be83016b696611e068c604051978891602083019e8f5260248301611bc0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810188528761019c565b856122bc565b98919690993594611acc565b925190206040519384933391859094939273ffffffffffffffffffffffffffffffffffffffff9067ffffffffffffffff606094608085019885521660208401521660408201520152565b0390a1929190565b7fd2316ede0000000000000000000000000000000000000000000000000000000060005260046000fd5b611ead915060203d60201161112a5761111c818361019c565b38611d49565b91909173ffffffffffffffffffffffffffffffffffffffff83168015611f64575073ffffffffffffffffffffffffffffffffffffffff16918215611efa576101ec92612461565b6000809350809281925af13d15611f5f573d611f15816101fd565b90611f23604051928361019c565b8152600060203d92013e5b15611f3557565b7fe417b80b0000000000000000000000000000000000000000000000000000000060005260046000fd5b611f2e565b7f26a78f8f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015611fd45760209160051b010190565b611f91565b908160209103126101475751611a4c81610388565b92949361202a60609367ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff941686526080602087015260808601906119bb565b95604085015216910152565b9061204082610370565b61204d604051918261019c565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061207b8294610370565b019060005b82811061208c57505050565b60209060405161209b8161017b565b6000815260008382015282828501015201612080565b9190811015611fd45760061b0190565b604081360312610147576020604051916120da8361017b565b80356120e581610317565b835201356120f281610388565b602082015290565b35611a4c81610388565b6020818303126101475780519067ffffffffffffffff821161014757019080601f8301121561014757815161213881610370565b92612146604051948561019c565b81845260208085019260051b82010192831161014757602001905b82821061216e5750505090565b60208091835161217d81610388565b815201910190612161565b6121bf8167ffffffffffffffff16600052600360205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b1561227c5760006122499167ffffffffffffffff811682526003602052612204610c81610c816040852073ffffffffffffffffffffffffffffffffffffffff90541690565b60405180809581947ffbca3b740000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa9081156106a05760009161225f575090565b611a4c91503d806000833e612274818361019c565b810190612104565b50604051602061228c818361019c565b600082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810190369083013790565b9391936122c960846101fd565b946122d7604051968761019c565b608486526122e560846101fd565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b156123b8575a9080821061238e578291038060061c90031115612364576000918291825a9560208451940192f1905a9003923d906084821161235b575b6000908287523e929190565b6084915061234f565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff60005416330361240357565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fd5b6101ec9273ffffffffffffffffffffffffffffffffffffffff604051937fa9059cbb0000000000000000000000000000000000000000000000000000000060208601521660248401526044830152604482526124be60648361019c565b6125a3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7bffffffffffffffff000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9160a01b16911681018091116125375790565b6124c3565b90919273ffffffffffffffffffffffffffffffffffffffff6101ec9481604051957f23b872dd0000000000000000000000000000000000000000000000000000000060208801521660248601521660448401526064830152606482526124be60848361019c565b906000602091828151910182855af115611aa8576000513d612625575073ffffffffffffffffffffffffffffffffffffffff81163b155b6125e15750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b600114156125da565b8054821015611fd45760005260206000200190600090565b9161267e918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b805480156126e6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906126b7828261262e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600560205260409020549081156127f2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161253757600454927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116125375783836000956127b195036127b7575b5050506127a06004612682565b600590600052602052604060002090565b55600190565b6127a06127e3916127d96127cf6127e995600461262e565b90549060031b1c90565b928391600461262e565b90612646565b55388080612793565b5050600090565b60008181526005602052604090205461288457600454680100000000000000008110156101975761286b612836826001859401600455600461262e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600454906000526005602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var RouterABI = RouterMetaData.ABI

var RouterBin = RouterMetaData.Bin

func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, wrappedNative common.Address, armProxy common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, wrappedNative, armProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{address: address, abi: *parsed, RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

type Router struct {
	address common.Address
	abi     abi.ABI
	RouterCaller
	RouterTransactor
	RouterFilterer
}

type RouterCaller struct {
	contract *bind.BoundContract
}

type RouterTransactor struct {
	contract *bind.BoundContract
}

type RouterFilterer struct {
	contract *bind.BoundContract
}

type RouterSession struct {
	Contract     *Router
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RouterCallerSession struct {
	Contract *RouterCaller
	CallOpts bind.CallOpts
}

type RouterTransactorSession struct {
	Contract     *RouterTransactor
	TransactOpts bind.TransactOpts
}

type RouterRaw struct {
	Contract *Router
}

type RouterCallerRaw struct {
	Contract *RouterCaller
}

type RouterTransactorRaw struct {
	Contract *RouterTransactor
}

func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	abi, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{address: address, abi: abi, RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

func (_Router *RouterCaller) MAXRETBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "MAX_RET_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Router *RouterSession) MAXRETBYTES() (uint16, error) {
	return _Router.Contract.MAXRETBYTES(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) MAXRETBYTES() (uint16, error) {
	return _Router.Contract.MAXRETBYTES(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetArmProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getArmProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetArmProxy() (common.Address, error) {
	return _Router.Contract.GetArmProxy(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetArmProxy() (common.Address, error) {
	return _Router.Contract.GetArmProxy(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getFee", destinationChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Router *RouterSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainSelector, message)
}

func (_Router *RouterCallerSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainSelector, message)
}

func (_Router *RouterCaller) GetOffRamps(opts *bind.CallOpts) ([]RouterOffRamp, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]RouterOffRamp), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterOffRamp)).(*[]RouterOffRamp)

	return out0, err

}

func (_Router *RouterSession) GetOffRamps() ([]RouterOffRamp, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetOffRamps() ([]RouterOffRamp, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetOnRamp(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOnRamp", destChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOnRamp(destChainSelector uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainSelector)
}

func (_Router *RouterCallerSession) GetOnRamp(destChainSelector uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainSelector)
}

func (_Router *RouterCaller) GetSupportedTokens(opts *bind.CallOpts, chainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSupportedTokens", chainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetSupportedTokens(chainSelector uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCallerSession) GetSupportedTokens(chainSelector uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCaller) GetWrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getWrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetWrappedNative() (common.Address, error) {
	return _Router.Contract.GetWrappedNative(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetWrappedNative() (common.Address, error) {
	return _Router.Contract.GetWrappedNative(&_Router.CallOpts)
}

func (_Router *RouterCaller) IsChainSupported(opts *bind.CallOpts, chainSelector uint64) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isChainSupported", chainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsChainSupported(chainSelector uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCallerSession) IsChainSupported(chainSelector uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCaller) IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isOffRamp", sourceChainSelector, offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, sourceChainSelector, offRamp)
}

func (_Router *RouterCallerSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, sourceChainSelector, offRamp)
}

func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

func (_Router *RouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Router *RouterSession) TypeAndVersion() (string, error) {
	return _Router.Contract.TypeAndVersion(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) TypeAndVersion() (string, error) {
	return _Router.Contract.TypeAndVersion(&_Router.CallOpts)
}

func (_Router *RouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "acceptOwnership")
}

func (_Router *RouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

func (_Router *RouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

func (_Router *RouterTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "applyRampUpdates", onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterSession) ApplyRampUpdates(onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterTransactorSession) ApplyRampUpdates(onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "ccipSend", destinationChainSelector, message)
}

func (_Router *RouterSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainSelector, message)
}

func (_Router *RouterTransactorSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainSelector, message)
}

func (_Router *RouterTransactor) RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "recoverTokens", tokenAddress, to, amount)
}

func (_Router *RouterSession) RecoverTokens(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RecoverTokens(&_Router.TransactOpts, tokenAddress, to, amount)
}

func (_Router *RouterTransactorSession) RecoverTokens(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RecoverTokens(&_Router.TransactOpts, tokenAddress, to, amount)
}

func (_Router *RouterTransactor) RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "routeMessage", message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterSession) RouteMessage(message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterTransactorSession) RouteMessage(message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterTransactor) SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setWrappedNative", wrappedNative)
}

func (_Router *RouterSession) SetWrappedNative(wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetWrappedNative(&_Router.TransactOpts, wrappedNative)
}

func (_Router *RouterTransactorSession) SetWrappedNative(wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetWrappedNative(&_Router.TransactOpts, wrappedNative)
}

func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", to)
}

func (_Router *RouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, to)
}

func (_Router *RouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, to)
}

type RouterMessageExecutedIterator struct {
	Event *RouterMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterMessageExecuted)
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
		it.Event = new(RouterMessageExecuted)
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

func (it *RouterMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *RouterMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterMessageExecuted struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	OffRamp             common.Address
	CalldataHash        [32]byte
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterMessageExecuted(opts *bind.FilterOpts) (*RouterMessageExecutedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return &RouterMessageExecutedIterator{contract: _Router.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *RouterMessageExecuted) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterMessageExecuted)
				if err := _Router.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

func (_Router *RouterFilterer) ParseMessageExecuted(log types.Log) (*RouterMessageExecuted, error) {
	event := new(RouterMessageExecuted)
	if err := _Router.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOffRampAddedIterator struct {
	Event *RouterOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampAdded)
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
		it.Event = new(RouterOffRampAdded)
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

func (it *RouterOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampAdded struct {
	SourceChainSelector uint64
	OffRamp             common.Address
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampAddedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampAdded", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampAddedIterator{contract: _Router.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampAdded", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampAdded)
				if err := _Router.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error) {
	event := new(RouterOffRampAdded)
	if err := _Router.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOffRampRemovedIterator struct {
	Event *RouterOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampRemoved)
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
		it.Event = new(RouterOffRampRemoved)
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

func (it *RouterOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampRemoved struct {
	SourceChainSelector uint64
	OffRamp             common.Address
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampRemovedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampRemoved", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampRemovedIterator{contract: _Router.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampRemoved", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampRemoved)
				if err := _Router.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error) {
	event := new(RouterOffRampRemoved)
	if err := _Router.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOnRampSetIterator struct {
	Event *RouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOnRampSet)
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
		it.Event = new(RouterOnRampSet)
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

func (it *RouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *RouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOnRampSet struct {
	DestChainSelector uint64
	OnRamp            common.Address
	Raw               types.Log
}

func (_Router *RouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, destChainSelector []uint64) (*RouterOnRampSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OnRampSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOnRampSetIterator{contract: _Router.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OnRampSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOnRampSet)
				if err := _Router.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_Router *RouterFilterer) ParseOnRampSet(log types.Log) (*RouterOnRampSet, error) {
	event := new(RouterOnRampSet)
	if err := _Router.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOwnershipTransferRequestedIterator struct {
	Event *RouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferRequested)
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
		it.Event = new(RouterOwnershipTransferRequested)
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

func (it *RouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Router *RouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferRequestedIterator{contract: _Router.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOwnershipTransferRequested)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Router *RouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*RouterOwnershipTransferRequested, error) {
	event := new(RouterOwnershipTransferRequested)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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

func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (RouterMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x9b877de93ea9895756e337442c657f95a34fc68e7eb988bdfa693d5be83016b6")
}

func (RouterOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0xa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b94")
}

func (RouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb")
}

func (RouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23")
}

func (RouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_Router *Router) Address() common.Address {
	return _Router.address
}

type RouterInterface interface {
	MAXRETBYTES(opts *bind.CallOpts) (uint16, error)

	GetArmProxy(opts *bind.CallOpts) (common.Address, error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]RouterOffRamp, error)

	GetOnRamp(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts, chainSelector uint64) ([]common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainSelector uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterMessageExecuted(opts *bind.FilterOpts) (*RouterMessageExecutedIterator, error)

	WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *RouterMessageExecuted) (event.Subscription, error)

	ParseMessageExecuted(log types.Log) (*RouterMessageExecuted, error)

	FilterOffRampAdded(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainSelector []uint64) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainSelector []uint64) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, destChainSelector []uint64) (*RouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainSelector []uint64) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*RouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error)

	Address() common.Address
}
