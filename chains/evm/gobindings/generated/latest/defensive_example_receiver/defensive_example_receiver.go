// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defensive_example_receiver

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

type CCIPClientExampleRemoteChainConfig struct {
	ExtraArgsBytes    []byte
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}

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

var DefensiveExampleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouterClient\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCIPClientExample.RemoteChainConfig\",\"components\":[{\"name\":\"extraArgsBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideFeeToken\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideNativeToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"retryFailedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_feeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_feeTokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_messageContents\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_nativeTokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendData\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumCCIPClientExample.PaymentMethod\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendDataAndTokens\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumCCIPClientExample.PaymentMethod\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTokens\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumCCIPClientExample.PaymentMethod\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSimRevert\",\"inputs\":[{\"name\":\"simRevert\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeToken\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNativeToken\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenDeposited\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageReceived\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageRecovered\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSucceeded\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeTokenDeposited\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeTokenWithdrawn\",\"inputs\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTokenTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientNativeTokenBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidPaymentMethod\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumCCIPClientExample.PaymentMethod\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MessageNotFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NativeTokenTransferFailed\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461012b576040816136a1803803809161001f828561016b565b83398101031261012b5780516001600160a01b038116919082900361012b57602001516001600160a01b038116919082900361012b5780156101555780608052331561014457600180546001600160a01b03199081163317909155600280549091168317905560405163095ea7b360e01b81526004810191909152600019602482015290602090829060449082906000905af18015610138576100fb575b60ff19600a5416600a556040516134fc90816101a58239608051818181610c3c01528181610f4501528181611351015281816119ef015261306b0152f35b6020813d602011610130575b816101146020938361016b565b8101031261012b57518015150361012b57386100bd565b600080fd5b3d9150610107565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b6335fdcccd60e21b600052600060045260246000fd5b601f909101601f19168101906001600160401b0382119082101761018e57604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611e30575080633185842b14611dcb578063369f7f6614611b1b5780633ebf92051461189c5780634c42071c146117d357806352f813c31461175d578063536c6bfa146115a7578063691b93011461121f578063723f6381146111ba5780637909b549146110ec57806379ba50971461100357806385572ffb14610f20578063898068fc14610dce5780638da5cb5b14610d7c578063ae81083714610c60578063b0f479a114610bf1578063b2ef14e314610a60578063cf6730f814610902578063d0d1fc3d14610862578063db4eac8614610520578063f084fba114610320578063f2fde38b14610230578063fc79e8751461017e5763ff2deec31461012757600080fd5b346101795760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b600080fd5b60007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795734156102025733600052600460205260406000206101c7348254612e36565b9055604080513381523460208201527f62c2c8e34665db7c56b2cabd7f5fb9702ccd352ffa8150147e450797e9f8e8f391819081015b0390a1005b7f3728b83d000000000000000000000000000000000000000000000000000000006000523460045260246000fd5b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795773ffffffffffffffffffffffffffffffffffffffff61027c611f1d565b610284612e53565b163381146102f657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101795760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957600435600281101561017957610363611f61565b60443567ffffffffffffffff81116101795761038390369060040161205e565b60643567ffffffffffffffff8111610179576103a390369060040161205e565b9167ffffffffffffffff8116928360005260036020526103c760406000205461225c565b156104f2576103d5856123d6565b841515806104de575b6104a757916104919161048760019673ffffffffffffffffffffffffffffffffffffffff7f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a9796602097604051906104368a83611fe3565b6000825260009b6104a2575b8b5260038952600160408c2094610458816123d6565b0361049b578260025416915b6040519761047189611f8f565b88528988015260408701521660608501526122af565b6080830152613041565b604051908152a180f35b8a91610464565b610442565b847f051e75ad000000000000000000000000000000000000000000000000000000006000526104d5816123d6565b60045260246000fd5b506104e8856123d6565b60018514156103de565b837fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101795760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957610557611f78565b60243567ffffffffffffffff81116101795761057790369060040161205e565b60443567ffffffffffffffff811161017957610597903690600401612371565b60643567ffffffffffffffff8111610179576105b7903690600401612371565b906084359360ff85168095036101795767ffffffffffffffff906105d9612e53565b604051946105e686611fc7565b85526020850192835260408501938452606085019586521660005260036020526040600020925180519067ffffffffffffffff821161075b576106338261062d875461225c565b876126a0565b602090601f83116001146107bf576106819291600091836107b4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b83555b51805190600184019067ffffffffffffffff831161075b576020906106a984846124a5565b0190600052602060002060005b83811061078a57505050506002820190519081519167ffffffffffffffff831161075b576020906106e784846124a5565b0190600052602060002060005b838110610731578560ff600387019151167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055600080f35b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016106f4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016106b6565b01519050878061064f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169186600052816000209260005b81811061084a5750908460019594939210610813575b505050811b018355610684565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055868080610806565b929360206001819287860151815501950193016107f0565b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610179576004356000526006602052604060002080546108fe67ffffffffffffffff600184015416926108f06108d160036108ca600285016122af565b93016122af565b91604051958695865260208601526080604086015260808501906121fd565b9083820360608501526121fd565b0390f35b34610179576109103661218e565b303303610a365767ffffffffffffffff61092c602083016125ab565b1680600052600360205261094460406000205461225c565b15610a09575060ff600a54166109df57608081019060005b61096683836126e5565b90508110156109dd57806109d773ffffffffffffffffffffffffffffffffffffffff6109a66109a160019561099b89896126e5565b90612e43565b612739565b1673ffffffffffffffffffffffffffffffffffffffff84541660206109cf8561099b8a8a6126e5565b013591612e9e565b0161095c565b005b7f79f79e0b0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f14d4a4e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101795760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957610a97611f1d565b73ffffffffffffffffffffffffffffffffffffffff81166024358115610bc3578015610b965733600052600560205260406000205492818410610b685790610b358160809493610b08827ff1f4a289503d669d2f79056086538264607226a3896bc870ad418089ee7104d1986124de565b33600052600560205260406000205573ffffffffffffffffffffffffffffffffffffffff60025416612e9e565b73ffffffffffffffffffffffffffffffffffffffff600254169160405192835233602084015260408301526060820152a1005b837f5f7b27fb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f3728b83d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f8e4c8aa60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101795760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610179577f32938ef74f35947fa82c8a4cf8722f037f2d4bfe7c2f16293db21336828d71776101fd6004353360005260056020526040600020610ccf828254612e36565b90556002546040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015233602482015230604482015260648101839052610d649173ffffffffffffffffffffffffffffffffffffffff16610d5f82608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283611fe3565b613289565b60408051338152602081019290925290918291820190565b346101795760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795767ffffffffffffffff610e0e611f78565b60006060604051610e1e81611fc7565b8181528160208201528160408201520152166000526003602052610eb2604060002060405190610e4d82611fc7565b610e56816122af565b825260ff610f14610e696001840161254a565b9260208501938452610ee3836003610e836002850161254a565b936040890194855201541694606087019586526040519788976020895251608060208a015260a08901906121fd565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0888303016040890152612144565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303016060870152612144565b91511660808301520390f35b3461017957610f2e3661218e565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163303610fd55767ffffffffffffffff610f80602083016125ab565b1690816000526003602052610f9960406000205461225c565b15610fa7576109dd9061275a565b507fc9ff038f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101795760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795760005473ffffffffffffffffffffffffffffffffffffffff811633036110c2577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795767ffffffffffffffff61112c611f78565b1660005260036020526111a2604060002060606111b060405161114e81611fc7565b611157846122af565b81526111656001850161254a565b906020810191825260ff600361117d6002880161254a565b9687604085015201541693849101525192604051948594606086526060860190612144565b908482036020860152612144565b9060408301520390f35b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795773ffffffffffffffffffffffffffffffffffffffff611206611f1d565b1660005260056020526020604060002054604051908152f35b346101795760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957600435600281101561017957611262611f61565b9060443567ffffffffffffffff81116101795761128390369060040161205e565b9060643567ffffffffffffffff8111610179576112a49036906004016120bd565b9067ffffffffffffffff84168060005260036020526112c760406000205461225c565b15610a09576112d5826123d6565b81151580611593575b611565579061134a9160005260036020526001604060002091611300816123d6565b036115465773ffffffffffffffffffffffffffffffffffffffff80600254169594955b6040519661133088611f8f565b8752606060208801528460408801521660608601526122af565b60808401527f00000000000000000000000000000000000000000000000000000000000000009260005b8251811015611510578060206113fe73ffffffffffffffffffffffffffffffffffffffff6113a46000958861240f565b515116826113b2858961240f565b5101516040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915294859283919082906064820190565b03925af19182156114e8576000926114f4575b5060206114a073ffffffffffffffffffffffffffffffffffffffff611436848861240f565b51511682611444858961240f565b5101516040519586809481937f095ea7b30000000000000000000000000000000000000000000000000000000083528d600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af19182156114e8576001926114ba575b5001611374565b6114da9060203d81116114e1575b6114d28183611fe3565b810190612452565b50866114b3565b503d6114c8565b6040513d6000823e3d90fd5b61150b9060203d81116114e1576114d28183611fe3565b611411565b7f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a602061153d8487613041565b604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff6000959495611323565b507f051e75ad000000000000000000000000000000000000000000000000000000006000526104d5816123d6565b5061159d826123d6565b60018214156112de565b346101795760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610179576115de611f1d565b60243573ffffffffffffffffffffffffffffffffffffffff821691821561172f578115611701573360005260046020526040600020548281106116d45782611625916124de565b336000526004602052604060002055600080808085855af161164561251a565b50156116a2576040805133815273ffffffffffffffffffffffffffffffffffffffff90921660208301528101919091527fc31c07d3d0aa96dfe35beac3846c2a6f1e00aa573fc7ebb1748fdb20bb4a657c915080606081016101fd565b50907fe3c188040000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd1d3ee180000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3728b83d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b827f8e4c8aa60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957600435801515809103610179576117a1612e53565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00600a5416911617600a55600080f35b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795767ffffffffffffffff611813611f78565b61181b612e53565b1660005260036020526000600360408220611836815461225c565b80611859575b5061184960018201612481565b61185560028201612481565b0155005b601f811160011461186f57508281555b8361183c565b8184526020842061188b91601f0160051c81019060010161246a565b808352826020812081835555611869565b346101795760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610179576004356002811015610179576118df611f61565b9060443567ffffffffffffffff81116101795761190090369060040161205e565b9060643567ffffffffffffffff81116101795761192190369060040161205e565b9060843567ffffffffffffffff8111610179576119429036906004016120bd565b9167ffffffffffffffff851680600052600360205261196560406000205461225c565b15610a0957611973836123d6565b82151580611b07575b611ad9576119e8929173ffffffffffffffffffffffffffffffffffffffff91600052600360205260016040600020936119b4816123d6565b03611ace578160025416909695965b604051976119d089611f8f565b885260208801528460408801521660608601526122af565b60808401527f00000000000000000000000000000000000000000000000000000000000000009260005b825181101561151057806020611a4273ffffffffffffffffffffffffffffffffffffffff6113a46000958861240f565b03925af19182156114e857600092611ab2575b506020611a7a73ffffffffffffffffffffffffffffffffffffffff611436848861240f565b03925af19182156114e857600192611a94575b5001611a12565b611aab9060203d81116114e1576114d28183611fe3565b5086611a8d565b611ac99060203d81116114e1576114d28183611fe3565b611a55565b6000909695966119c3565b827f051e75ad000000000000000000000000000000000000000000000000000000006000526104d5816123d6565b50611b11836123d6565b600183141561197c565b346101795760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795760243560043573ffffffffffffffffffffffffffffffffffffffff8216820361017957611b75612e53565b8060005260096020526040600020548015801590611db4575b15611d56577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611d295780600052600960205260006040812055611bd2816133a2565b50806000526006602052604060002091600460405193611bf185611f8f565b8054855267ffffffffffffffff6001820154166020860152611c15600282016122af565b6040860152611c26600382016122af565b606086015201928354611c38816120a5565b94611c466040519687611fe3565b818652602086019060005260206000206000915b838310611ce6575050505060800192835260005b83518051821015611cbe5790611cb873ffffffffffffffffffffffffffffffffffffffff611c9e8360019561240f565b515116846020611caf858a5161240f565b51015191612e9e565b01611c6e565b837fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad600080a2005b60026020600192604051611cf981611fab565b73ffffffffffffffffffffffffffffffffffffffff86541681528486015483820152815201920192019190611c5a565b7fb6e782600000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152fd5b508160005260086020526040600020541515611b8e565b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101795773ffffffffffffffffffffffffffffffffffffffff611e17611f1d565b1660005260046020526020604060002054604051908152f35b346101795760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017957600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361017957817f85572ffb0000000000000000000000000000000000000000000000000000000060209314908115611ef3575b8115611ec9575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611ec2565b7f7909b5490000000000000000000000000000000000000000000000000000000081149150611ebb565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361017957565b359073ffffffffffffffffffffffffffffffffffffffff8216820361017957565b6024359067ffffffffffffffff8216820361017957565b6004359067ffffffffffffffff8216820361017957565b60a0810190811067ffffffffffffffff82111761075b57604052565b6040810190811067ffffffffffffffff82111761075b57604052565b6080810190811067ffffffffffffffff82111761075b57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761075b57604052565b67ffffffffffffffff811161075b57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f820112156101795780359061207582612024565b926120836040519485611fe3565b8284526020838301011161017957816000926020809301838601378301015290565b67ffffffffffffffff811161075b5760051b60200190565b81601f82011215610179578035906120d4826120a5565b926120e26040519485611fe3565b82845260208085019360061b8301019181831161017957602001925b82841061210c575050505090565b604084830312610179576020604091825161212681611fab565b61212f87611f40565b815282870135838201528152019301926120fe565b906020808351928381520192019060005b8181106121625750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612155565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610179576004359067ffffffffffffffff8211610179577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126101795760040190565b919082519283825260005b8481106122475750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612208565b90600182811c921680156122a5575b602083101461227657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161226b565b90604051918260008254926122c38461225c565b808452936001811690811561233157506001146122ea575b506122e892500383611fe3565b565b90506000929192526020600020906000915b8183106123155750509060206122e892820101386122db565b60209193508060019154838589010152019101909184926122fc565b602093506122e89592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386122db565b9080601f83011215610179578135612388816120a5565b926123966040519485611fe3565b81845260208085019260051b82010192831161017957602001905b8282106123be5750505090565b602080916123cb84611f40565b8152019101906123b1565b600211156123e057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b80518210156124235760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90816020910312610179575180151581036101795790565b818110612475575050565b6000815560010161246a565b80546000825580612490575050565b6122e89160005260206000209081019061246a565b9068010000000000000000811161075b578154908083558181106124c857505050565b6122e8926000526020600020918201910161246a565b919082039182116124eb57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b3d15612545573d9061252b82612024565b916125396040519384611fe3565b82523d6000602084013e565b606090565b906040519182815491828252602082019060005260206000209260005b81811061257c5750506122e892500383611fe3565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612567565b3567ffffffffffffffff811681036101795790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561017957016020813591019167ffffffffffffffff821161017957813603831361017957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610179570180359067ffffffffffffffff82116101795760200191813603831361017957565b9190601f81116126af57505050565b6122e8926000526020600020906020601f840160051c830193106126db575b601f0160051c019061246a565b90915081906126ce565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610179570180359067ffffffffffffffff821161017957602001918160061b3603831361017957565b3573ffffffffffffffffffffffffffffffffffffffff811681036101795790565b303b15610179576040517fcf6730f8000000000000000000000000000000000000000000000000000000008152600091602060048301528035928360248401526020820192833567ffffffffffffffff8116809103612e3257604482015260408301936127db6127ca86866125c0565b60a0606486015260c4850191612610565b61281d60608601916127ed83886125c0565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc878403016084880152612610565b92608086019384357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301811215612e2e5787016020813591019167ffffffffffffffff8211612e2a578160061b36038313612e2a578381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160a48501528181528791849160200190835b818110612de557505081929350038183305af19081612dd1575b50612da5576128d361251a565b958785526009602052600160408620556128ec886133a2565b508785526006602052604085209288845567ffffffffffffffff61291360018601926125ab565b167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082541617905561294960028401918761264f565b9067ffffffffffffffff8211612d785761296d82612967855461225c565b856126a0565b8690601f8311600114612cd8576129b89291889183612bf85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6129c960038301918661264f565b9067ffffffffffffffff8211612cab576129e782612967855461225c565b8590601f8311600114612c035782612a4596959360049593612a3b938a92612bf85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b01936126e5565b9290680100000000000000008411612bcb578154848355808510612b19575b5090825260208220905b838310612ab75750505050612ab27f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f916040519182916020835260208301906121fd565b0390a2565b600260408273ffffffffffffffffffffffffffffffffffffffff612adc600195612739565b167fffffffffffffffffffffffff000000000000000000000000000000000000000086541617855560208101358486015501920192019190612a6e565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168103612b9e577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85168503612b9e57828452602084209060011b8101908560011b015b818110612b8c5750612a64565b80856002925585600182015501612b7f565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b01359050388061064f565b83875260208720917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416885b818110612c935750926001928592612a459998966004989610612c5b575b505050811b019055612a3e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612c4e565b91936020600181928787013581550195019201612c30565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b83885260208820917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416895b818110612d605750908460019594939210612d28575b505050811b0190556129bb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612d1b565b91936020600181928787013581550195019201612d05565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50505090507fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f915080a2565b85612dde91969296611fe3565b93386128c6565b9250925060408060019273ffffffffffffffffffffffffffffffffffffffff612e0d88611f40565b1681526020870135602082015201940191019288928592946128ac565b8780fd5b8680fd5b8280fd5b919082018092116124eb57565b91908110156124235760061b0190565b73ffffffffffffffffffffffffffffffffffffffff600154163303612e7457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff909216602483015260448201929092526122e891610d5f8260648101610d33565b9067ffffffffffffffff9093929316815260406020820152612f62612f2e845160a0604085015260e08401906121fd565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030160608501526121fd565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b8181106130095750505060808473ffffffffffffffffffffffffffffffffffffffff6060613006969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828503019101526121fd565b90565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101612fa3565b6040517f20487ded00000000000000000000000000000000000000000000000000000000815290917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16602083806130b2858860048401612efd565b0381845afa9283156114e857600093613252575b50600073ffffffffffffffffffffffffffffffffffffffff606084015116156000146131da5750336000526004602052826040600020541061319e5761315993602093336000526004855260406000206131218282546124de565b90555b6040518096819582947f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401612efd565b03925af19081156114e85760009161316f575090565b90506020813d602011613196575b8161318a60209383611fe3565b81010312610179575190565b3d915061317d565b3360005260046020526040600020547fd1d3ee180000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9333600052600560205283604060002054106132165761315994602094336000526005865261320f60406000209182546124de565b9055613124565b3360005260056020526040600020547f5f7b27fb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90926020823d602011613281575b8161326d60209383611fe3565b8101031261327e57505191386130c6565b80fd5b3d9150613260565b73ffffffffffffffffffffffffffffffffffffffff6132f99116916040926000808551936132b78786611fe3565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af16132f361251a565b9161341f565b8051908161330657505050565b602080613317938301019101612452565b1561331f5750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80600052600860205260406000205415600014613419576007546801000000000000000081101561075b576001810180600755811015612423577fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880181905560075460009182526008602052604090912055600190565b50600090565b9192901561349a5750815115613433575090565b3b1561343c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156134ad5750805190602001fd5b6134eb906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906121fd565b0390fdfea164736f6c634300081a000a",
}

var DefensiveExampleABI = DefensiveExampleMetaData.ABI

var DefensiveExampleBin = DefensiveExampleMetaData.Bin

func DeployDefensiveExample(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address) (common.Address, *types.Transaction, *DefensiveExample, error) {
	parsed, err := DefensiveExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefensiveExampleBin), backend, router, feeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefensiveExample{address: address, abi: *parsed, DefensiveExampleCaller: DefensiveExampleCaller{contract: contract}, DefensiveExampleTransactor: DefensiveExampleTransactor{contract: contract}, DefensiveExampleFilterer: DefensiveExampleFilterer{contract: contract}}, nil
}

type DefensiveExample struct {
	address common.Address
	abi     abi.ABI
	DefensiveExampleCaller
	DefensiveExampleTransactor
	DefensiveExampleFilterer
}

type DefensiveExampleCaller struct {
	contract *bind.BoundContract
}

type DefensiveExampleTransactor struct {
	contract *bind.BoundContract
}

type DefensiveExampleFilterer struct {
	contract *bind.BoundContract
}

type DefensiveExampleSession struct {
	Contract     *DefensiveExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DefensiveExampleCallerSession struct {
	Contract *DefensiveExampleCaller
	CallOpts bind.CallOpts
}

type DefensiveExampleTransactorSession struct {
	Contract     *DefensiveExampleTransactor
	TransactOpts bind.TransactOpts
}

type DefensiveExampleRaw struct {
	Contract *DefensiveExample
}

type DefensiveExampleCallerRaw struct {
	Contract *DefensiveExampleCaller
}

type DefensiveExampleTransactorRaw struct {
	Contract *DefensiveExampleTransactor
}

func NewDefensiveExample(address common.Address, backend bind.ContractBackend) (*DefensiveExample, error) {
	abi, err := abi.JSON(strings.NewReader(DefensiveExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindDefensiveExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefensiveExample{address: address, abi: abi, DefensiveExampleCaller: DefensiveExampleCaller{contract: contract}, DefensiveExampleTransactor: DefensiveExampleTransactor{contract: contract}, DefensiveExampleFilterer: DefensiveExampleFilterer{contract: contract}}, nil
}

func NewDefensiveExampleCaller(address common.Address, caller bind.ContractCaller) (*DefensiveExampleCaller, error) {
	contract, err := bindDefensiveExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleCaller{contract: contract}, nil
}

func NewDefensiveExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*DefensiveExampleTransactor, error) {
	contract, err := bindDefensiveExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleTransactor{contract: contract}, nil
}

func NewDefensiveExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*DefensiveExampleFilterer, error) {
	contract, err := bindDefensiveExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleFilterer{contract: contract}, nil
}

func bindDefensiveExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefensiveExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_DefensiveExample *DefensiveExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefensiveExample.Contract.DefensiveExampleCaller.contract.Call(opts, result, method, params...)
}

func (_DefensiveExample *DefensiveExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DefensiveExampleTransactor.contract.Transfer(opts)
}

func (_DefensiveExample *DefensiveExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DefensiveExampleTransactor.contract.Transact(opts, method, params...)
}

func (_DefensiveExample *DefensiveExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefensiveExample.Contract.contract.Call(opts, result, method, params...)
}

func (_DefensiveExample *DefensiveExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.Contract.contract.Transfer(opts)
}

func (_DefensiveExample *DefensiveExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefensiveExample.Contract.contract.Transact(opts, method, params...)
}

func (_DefensiveExample *DefensiveExampleCaller) GetCCVs(opts *bind.CallOpts, sourceChainSelector uint64) (GetCCVs,

	error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getCCVs", sourceChainSelector)

	outstruct := new(GetCCVs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_DefensiveExample *DefensiveExampleSession) GetCCVs(sourceChainSelector uint64) (GetCCVs,

	error) {
	return _DefensiveExample.Contract.GetCCVs(&_DefensiveExample.CallOpts, sourceChainSelector)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetCCVs(sourceChainSelector uint64) (GetCCVs,

	error) {
	return _DefensiveExample.Contract.GetCCVs(&_DefensiveExample.CallOpts, sourceChainSelector)
}

func (_DefensiveExample *DefensiveExampleCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	if err != nil {
		return *new(CCIPClientExampleRemoteChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPClientExampleRemoteChainConfig)).(*CCIPClientExampleRemoteChainConfig)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) GetRemoteChainConfig(remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	return _DefensiveExample.Contract.GetRemoteChainConfig(&_DefensiveExample.CallOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error) {
	return _DefensiveExample.Contract.GetRemoteChainConfig(&_DefensiveExample.CallOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) GetRouter() (common.Address, error) {
	return _DefensiveExample.Contract.GetRouter(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) GetRouter() (common.Address, error) {
	return _DefensiveExample.Contract.GetRouter(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) Owner() (common.Address, error) {
	return _DefensiveExample.Contract.Owner(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) Owner() (common.Address, error) {
	return _DefensiveExample.Contract.Owner(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) SFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SFeeToken() (common.Address, error) {
	return _DefensiveExample.Contract.SFeeToken(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SFeeToken() (common.Address, error) {
	return _DefensiveExample.Contract.SFeeToken(&_DefensiveExample.CallOpts)
}

func (_DefensiveExample *DefensiveExampleCaller) SFeeTokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_feeTokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SFeeTokenBalances(arg0 common.Address) (*big.Int, error) {
	return _DefensiveExample.Contract.SFeeTokenBalances(&_DefensiveExample.CallOpts, arg0)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SFeeTokenBalances(arg0 common.Address) (*big.Int, error) {
	return _DefensiveExample.Contract.SFeeTokenBalances(&_DefensiveExample.CallOpts, arg0)
}

func (_DefensiveExample *DefensiveExampleCaller) SMessageContents(opts *bind.CallOpts, messageId [32]byte) (SMessageContents,

	error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_messageContents", messageId)

	outstruct := new(SMessageContents)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SourceChainSelector = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_DefensiveExample *DefensiveExampleSession) SMessageContents(messageId [32]byte) (SMessageContents,

	error) {
	return _DefensiveExample.Contract.SMessageContents(&_DefensiveExample.CallOpts, messageId)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SMessageContents(messageId [32]byte) (SMessageContents,

	error) {
	return _DefensiveExample.Contract.SMessageContents(&_DefensiveExample.CallOpts, messageId)
}

func (_DefensiveExample *DefensiveExampleCaller) SNativeTokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "s_nativeTokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SNativeTokenBalances(arg0 common.Address) (*big.Int, error) {
	return _DefensiveExample.Contract.SNativeTokenBalances(&_DefensiveExample.CallOpts, arg0)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SNativeTokenBalances(arg0 common.Address) (*big.Int, error) {
	return _DefensiveExample.Contract.SNativeTokenBalances(&_DefensiveExample.CallOpts, arg0)
}

func (_DefensiveExample *DefensiveExampleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DefensiveExample.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_DefensiveExample *DefensiveExampleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DefensiveExample.Contract.SupportsInterface(&_DefensiveExample.CallOpts, interfaceId)
}

func (_DefensiveExample *DefensiveExampleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DefensiveExample.Contract.SupportsInterface(&_DefensiveExample.CallOpts, interfaceId)
}

func (_DefensiveExample *DefensiveExampleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "acceptOwnership")
}

func (_DefensiveExample *DefensiveExampleSession) AcceptOwnership() (*types.Transaction, error) {
	return _DefensiveExample.Contract.AcceptOwnership(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DefensiveExample.Contract.AcceptOwnership(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "ccipReceive", message)
}

func (_DefensiveExample *DefensiveExampleSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.CcipReceive(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.CcipReceive(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactor) DisableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "disableRemoteChain", remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleSession) DisableRemoteChain(remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) DisableRemoteChain(remoteChainSelector uint64) (*types.Transaction, error) {
	return _DefensiveExample.Contract.DisableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector)
}

func (_DefensiveExample *DefensiveExampleTransactor) EnableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "enableRemoteChain", remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleSession) EnableRemoteChain(remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) EnableRemoteChain(remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error) {
	return _DefensiveExample.Contract.EnableRemoteChain(&_DefensiveExample.TransactOpts, remoteChainSelector, extraArgs, requiredCCVs, optionalCCVs, optionalThreshold)
}

func (_DefensiveExample *DefensiveExampleTransactor) ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "processMessage", message)
}

func (_DefensiveExample *DefensiveExampleSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProcessMessage(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProcessMessage(&_DefensiveExample.TransactOpts, message)
}

func (_DefensiveExample *DefensiveExampleTransactor) ProvideFeeToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "provideFeeToken", amount)
}

func (_DefensiveExample *DefensiveExampleSession) ProvideFeeToken(amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProvideFeeToken(&_DefensiveExample.TransactOpts, amount)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) ProvideFeeToken(amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProvideFeeToken(&_DefensiveExample.TransactOpts, amount)
}

func (_DefensiveExample *DefensiveExampleTransactor) ProvideNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "provideNativeToken")
}

func (_DefensiveExample *DefensiveExampleSession) ProvideNativeToken() (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProvideNativeToken(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) ProvideNativeToken() (*types.Transaction, error) {
	return _DefensiveExample.Contract.ProvideNativeToken(&_DefensiveExample.TransactOpts)
}

func (_DefensiveExample *DefensiveExampleTransactor) RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "retryFailedMessage", messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleSession) RetryFailedMessage(messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.RetryFailedMessage(&_DefensiveExample.TransactOpts, messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) RetryFailedMessage(messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.RetryFailedMessage(&_DefensiveExample.TransactOpts, messageId, tokenReceiver)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendData(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendData", method, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleSession) SendData(method uint8, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendData(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendData(method uint8, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendData(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, data)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendDataAndTokens(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendDataAndTokens", method, destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleSession) SendDataAndTokens(method uint8, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataAndTokens(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendDataAndTokens(method uint8, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendDataAndTokens(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, data, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactor) SendTokens(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "sendTokens", method, destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleSession) SendTokens(method uint8, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendTokens(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SendTokens(method uint8, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SendTokens(&_DefensiveExample.TransactOpts, method, destChainSelector, receiver, tokenAmounts)
}

func (_DefensiveExample *DefensiveExampleTransactor) SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "setSimRevert", simRevert)
}

func (_DefensiveExample *DefensiveExampleSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SetSimRevert(&_DefensiveExample.TransactOpts, simRevert)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _DefensiveExample.Contract.SetSimRevert(&_DefensiveExample.TransactOpts, simRevert)
}

func (_DefensiveExample *DefensiveExampleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "transferOwnership", to)
}

func (_DefensiveExample *DefensiveExampleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.TransferOwnership(&_DefensiveExample.TransactOpts, to)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DefensiveExample.Contract.TransferOwnership(&_DefensiveExample.TransactOpts, to)
}

func (_DefensiveExample *DefensiveExampleTransactor) WithdrawFeeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "withdrawFeeToken", to, amount)
}

func (_DefensiveExample *DefensiveExampleSession) WithdrawFeeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.WithdrawFeeToken(&_DefensiveExample.TransactOpts, to, amount)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) WithdrawFeeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.WithdrawFeeToken(&_DefensiveExample.TransactOpts, to, amount)
}

func (_DefensiveExample *DefensiveExampleTransactor) WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.contract.Transact(opts, "withdrawNativeToken", to, amount)
}

func (_DefensiveExample *DefensiveExampleSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.WithdrawNativeToken(&_DefensiveExample.TransactOpts, to, amount)
}

func (_DefensiveExample *DefensiveExampleTransactorSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefensiveExample.Contract.WithdrawNativeToken(&_DefensiveExample.TransactOpts, to, amount)
}

type DefensiveExampleFeeTokenDepositedIterator struct {
	Event *DefensiveExampleFeeTokenDeposited

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleFeeTokenDepositedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleFeeTokenDeposited)
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
		it.Event = new(DefensiveExampleFeeTokenDeposited)
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

func (it *DefensiveExampleFeeTokenDepositedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleFeeTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleFeeTokenDeposited struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterFeeTokenDeposited(opts *bind.FilterOpts) (*DefensiveExampleFeeTokenDepositedIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "FeeTokenDeposited")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleFeeTokenDepositedIterator{contract: _DefensiveExample.contract, event: "FeeTokenDeposited", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchFeeTokenDeposited(opts *bind.WatchOpts, sink chan<- *DefensiveExampleFeeTokenDeposited) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "FeeTokenDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleFeeTokenDeposited)
				if err := _DefensiveExample.contract.UnpackLog(event, "FeeTokenDeposited", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseFeeTokenDeposited(log types.Log) (*DefensiveExampleFeeTokenDeposited, error) {
	event := new(DefensiveExampleFeeTokenDeposited)
	if err := _DefensiveExample.contract.UnpackLog(event, "FeeTokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleFeeTokenWithdrawnIterator struct {
	Event *DefensiveExampleFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleFeeTokenWithdrawn)
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
		it.Event = new(DefensiveExampleFeeTokenWithdrawn)
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

func (it *DefensiveExampleFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleFeeTokenWithdrawn struct {
	Token      common.Address
	Withdrawer common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts) (*DefensiveExampleFeeTokenWithdrawnIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "FeeTokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleFeeTokenWithdrawnIterator{contract: _DefensiveExample.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DefensiveExampleFeeTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "FeeTokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleFeeTokenWithdrawn)
				if err := _DefensiveExample.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseFeeTokenWithdrawn(log types.Log) (*DefensiveExampleFeeTokenWithdrawn, error) {
	event := new(DefensiveExampleFeeTokenWithdrawn)
	if err := _DefensiveExample.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageFailedIterator struct {
	Event *DefensiveExampleMessageFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageFailed)
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
		it.Event = new(DefensiveExampleMessageFailed)
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

func (it *DefensiveExampleMessageFailedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageFailedIterator{contract: _DefensiveExample.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageFailed)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageFailed(log types.Log) (*DefensiveExampleMessageFailed, error) {
	event := new(DefensiveExampleMessageFailed)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageReceivedIterator struct {
	Event *DefensiveExampleMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageReceived)
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
		it.Event = new(DefensiveExampleMessageReceived)
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

func (it *DefensiveExampleMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageReceived struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*DefensiveExampleMessageReceivedIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageReceivedIterator{contract: _DefensiveExample.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageReceived) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageReceived)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageReceived(log types.Log) (*DefensiveExampleMessageReceived, error) {
	event := new(DefensiveExampleMessageReceived)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageRecoveredIterator struct {
	Event *DefensiveExampleMessageRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageRecovered)
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
		it.Event = new(DefensiveExampleMessageRecovered)
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

func (it *DefensiveExampleMessageRecoveredIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageRecovered struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageRecoveredIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageRecoveredIterator{contract: _DefensiveExample.contract, event: "MessageRecovered", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageRecovered, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageRecovered)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageRecovered(log types.Log) (*DefensiveExampleMessageRecovered, error) {
	event := new(DefensiveExampleMessageRecovered)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageSentIterator struct {
	Event *DefensiveExampleMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageSent)
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
		it.Event = new(DefensiveExampleMessageSent)
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

func (it *DefensiveExampleMessageSentIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageSent struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageSent(opts *bind.FilterOpts) (*DefensiveExampleMessageSentIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageSentIterator{contract: _DefensiveExample.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSent) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageSent)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageSent(log types.Log) (*DefensiveExampleMessageSent, error) {
	event := new(DefensiveExampleMessageSent)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleMessageSucceededIterator struct {
	Event *DefensiveExampleMessageSucceeded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleMessageSucceededIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleMessageSucceeded)
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
		it.Event = new(DefensiveExampleMessageSucceeded)
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

func (it *DefensiveExampleMessageSucceededIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleMessageSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleMessageSucceeded struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageSucceededIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleMessageSucceededIterator{contract: _DefensiveExample.contract, event: "MessageSucceeded", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSucceeded, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleMessageSucceeded)
				if err := _DefensiveExample.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseMessageSucceeded(log types.Log) (*DefensiveExampleMessageSucceeded, error) {
	event := new(DefensiveExampleMessageSucceeded)
	if err := _DefensiveExample.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleNativeTokenDepositedIterator struct {
	Event *DefensiveExampleNativeTokenDeposited

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleNativeTokenDepositedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleNativeTokenDeposited)
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
		it.Event = new(DefensiveExampleNativeTokenDeposited)
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

func (it *DefensiveExampleNativeTokenDepositedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleNativeTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleNativeTokenDeposited struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterNativeTokenDeposited(opts *bind.FilterOpts) (*DefensiveExampleNativeTokenDepositedIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "NativeTokenDeposited")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleNativeTokenDepositedIterator{contract: _DefensiveExample.contract, event: "NativeTokenDeposited", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchNativeTokenDeposited(opts *bind.WatchOpts, sink chan<- *DefensiveExampleNativeTokenDeposited) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "NativeTokenDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleNativeTokenDeposited)
				if err := _DefensiveExample.contract.UnpackLog(event, "NativeTokenDeposited", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseNativeTokenDeposited(log types.Log) (*DefensiveExampleNativeTokenDeposited, error) {
	event := new(DefensiveExampleNativeTokenDeposited)
	if err := _DefensiveExample.contract.UnpackLog(event, "NativeTokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleNativeTokenWithdrawnIterator struct {
	Event *DefensiveExampleNativeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleNativeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleNativeTokenWithdrawn)
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
		it.Event = new(DefensiveExampleNativeTokenWithdrawn)
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

func (it *DefensiveExampleNativeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleNativeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleNativeTokenWithdrawn struct {
	Withdrawer common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterNativeTokenWithdrawn(opts *bind.FilterOpts) (*DefensiveExampleNativeTokenWithdrawnIterator, error) {

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "NativeTokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleNativeTokenWithdrawnIterator{contract: _DefensiveExample.contract, event: "NativeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchNativeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DefensiveExampleNativeTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "NativeTokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleNativeTokenWithdrawn)
				if err := _DefensiveExample.contract.UnpackLog(event, "NativeTokenWithdrawn", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseNativeTokenWithdrawn(log types.Log) (*DefensiveExampleNativeTokenWithdrawn, error) {
	event := new(DefensiveExampleNativeTokenWithdrawn)
	if err := _DefensiveExample.contract.UnpackLog(event, "NativeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleOwnershipTransferRequestedIterator struct {
	Event *DefensiveExampleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleOwnershipTransferRequested)
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
		it.Event = new(DefensiveExampleOwnershipTransferRequested)
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

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleOwnershipTransferRequestedIterator{contract: _DefensiveExample.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleOwnershipTransferRequested)
				if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseOwnershipTransferRequested(log types.Log) (*DefensiveExampleOwnershipTransferRequested, error) {
	event := new(DefensiveExampleOwnershipTransferRequested)
	if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DefensiveExampleOwnershipTransferredIterator struct {
	Event *DefensiveExampleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DefensiveExampleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefensiveExampleOwnershipTransferred)
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
		it.Event = new(DefensiveExampleOwnershipTransferred)
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

func (it *DefensiveExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *DefensiveExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DefensiveExampleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DefensiveExample *DefensiveExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DefensiveExampleOwnershipTransferredIterator{contract: _DefensiveExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_DefensiveExample *DefensiveExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DefensiveExample.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DefensiveExampleOwnershipTransferred)
				if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_DefensiveExample *DefensiveExampleFilterer) ParseOwnershipTransferred(log types.Log) (*DefensiveExampleOwnershipTransferred, error) {
	event := new(DefensiveExampleOwnershipTransferred)
	if err := _DefensiveExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
type SMessageContents struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
}

func (DefensiveExampleFeeTokenDeposited) Topic() common.Hash {
	return common.HexToHash("0x32938ef74f35947fa82c8a4cf8722f037f2d4bfe7c2f16293db21336828d7177")
}

func (DefensiveExampleFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0xf1f4a289503d669d2f79056086538264607226a3896bc870ad418089ee7104d1")
}

func (DefensiveExampleMessageFailed) Topic() common.Hash {
	return common.HexToHash("0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f")
}

func (DefensiveExampleMessageReceived) Topic() common.Hash {
	return common.HexToHash("0xe29dc34207c78fc0f6048a32f159139c33339c6d6df8b07dcd33f6d699ff2327")
}

func (DefensiveExampleMessageRecovered) Topic() common.Hash {
	return common.HexToHash("0xef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad")
}

func (DefensiveExampleMessageSent) Topic() common.Hash {
	return common.HexToHash("0x54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a")
}

func (DefensiveExampleMessageSucceeded) Topic() common.Hash {
	return common.HexToHash("0xdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f")
}

func (DefensiveExampleNativeTokenDeposited) Topic() common.Hash {
	return common.HexToHash("0x62c2c8e34665db7c56b2cabd7f5fb9702ccd352ffa8150147e450797e9f8e8f3")
}

func (DefensiveExampleNativeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0xc31c07d3d0aa96dfe35beac3846c2a6f1e00aa573fc7ebb1748fdb20bb4a657c")
}

func (DefensiveExampleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (DefensiveExampleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_DefensiveExample *DefensiveExample) Address() common.Address {
	return _DefensiveExample.address
}

type DefensiveExampleInterface interface {
	GetCCVs(opts *bind.CallOpts, sourceChainSelector uint64) (GetCCVs,

		error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (CCIPClientExampleRemoteChainConfig, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SFeeToken(opts *bind.CallOpts) (common.Address, error)

	SFeeTokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	SMessageContents(opts *bind.CallOpts, messageId [32]byte) (SMessageContents,

		error)

	SNativeTokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	EnableRemoteChain(opts *bind.TransactOpts, remoteChainSelector uint64, extraArgs []byte, requiredCCVs []common.Address, optionalCCVs []common.Address, optionalThreshold uint8) (*types.Transaction, error)

	ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	ProvideFeeToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideNativeToken(opts *bind.TransactOpts) (*types.Transaction, error)

	RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, tokenReceiver common.Address) (*types.Transaction, error)

	SendData(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, data []byte) (*types.Transaction, error)

	SendDataAndTokens(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, data []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error)

	SendTokens(opts *bind.TransactOpts, method uint8, destChainSelector uint64, receiver []byte, tokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error)

	SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeTokenDeposited(opts *bind.FilterOpts) (*DefensiveExampleFeeTokenDepositedIterator, error)

	WatchFeeTokenDeposited(opts *bind.WatchOpts, sink chan<- *DefensiveExampleFeeTokenDeposited) (event.Subscription, error)

	ParseFeeTokenDeposited(log types.Log) (*DefensiveExampleFeeTokenDeposited, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts) (*DefensiveExampleFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DefensiveExampleFeeTokenWithdrawn) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*DefensiveExampleFeeTokenWithdrawn, error)

	FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageFailedIterator, error)

	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageFailed, messageId [][32]byte) (event.Subscription, error)

	ParseMessageFailed(log types.Log) (*DefensiveExampleMessageFailed, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*DefensiveExampleMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*DefensiveExampleMessageReceived, error)

	FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageRecoveredIterator, error)

	WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageRecovered, messageId [][32]byte) (event.Subscription, error)

	ParseMessageRecovered(log types.Log) (*DefensiveExampleMessageRecovered, error)

	FilterMessageSent(opts *bind.FilterOpts) (*DefensiveExampleMessageSentIterator, error)

	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSent) (event.Subscription, error)

	ParseMessageSent(log types.Log) (*DefensiveExampleMessageSent, error)

	FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*DefensiveExampleMessageSucceededIterator, error)

	WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *DefensiveExampleMessageSucceeded, messageId [][32]byte) (event.Subscription, error)

	ParseMessageSucceeded(log types.Log) (*DefensiveExampleMessageSucceeded, error)

	FilterNativeTokenDeposited(opts *bind.FilterOpts) (*DefensiveExampleNativeTokenDepositedIterator, error)

	WatchNativeTokenDeposited(opts *bind.WatchOpts, sink chan<- *DefensiveExampleNativeTokenDeposited) (event.Subscription, error)

	ParseNativeTokenDeposited(log types.Log) (*DefensiveExampleNativeTokenDeposited, error)

	FilterNativeTokenWithdrawn(opts *bind.FilterOpts) (*DefensiveExampleNativeTokenWithdrawnIterator, error)

	WatchNativeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DefensiveExampleNativeTokenWithdrawn) (event.Subscription, error)

	ParseNativeTokenWithdrawn(log types.Log) (*DefensiveExampleNativeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*DefensiveExampleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefensiveExampleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefensiveExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*DefensiveExampleOwnershipTransferred, error)

	Address() common.Address
}
