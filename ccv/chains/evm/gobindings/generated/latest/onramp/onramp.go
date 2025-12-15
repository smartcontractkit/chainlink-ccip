// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp

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

type OnRampDestChainConfig struct {
	Router               common.Address
	MessageNumber        uint64
	AddressBytesLength   uint8
	NetworkFeeUSDCents   uint16
	TokenReceiverAllowed bool
	BaseExecutionGasCost uint32
	DefaultExecutor      common.Address
	LaneMandatedCCVs     []common.Address
	DefaultCCVs          []common.Address
	OffRamp              []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector    uint64
	Router               common.Address
	AddressBytesLength   uint8
	NetworkFeeUSDCents   uint16
	TokenReceiverAllowed bool
	BaseExecutionGasCost uint32
	DefaultCCVs          []common.Address
	LaneMandatedCCVs     []common.Address
	DefaultExecutor      common.Address
	OffRamp              []byte
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type OnRampReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint32
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc5760405161668a38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a160405161633f908161034b82396080518181816109be015281816113ab0152611bb9015260a0518181816111400152611be5015260c051818181611c1401526126d90152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610117578063181f5a771461011257806320487ded1461010d5780633fc894981461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e557806390423fa2146100e0578063df0aa9e9146100db578063e8d80861146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611b34565b611a78565b611a09565b61109a565b610f0b565b610ec4565b610e13565b610d8c565b610cf4565b610b76565b610a39565b6109f6565b6105b3565b610376565b6102e8565b3461016a57600060031936011261016a576060610132611b99565b61016860405180926001600160a01b036040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610140810190811067ffffffffffffffff8211176101bb57604052565b61016f565b6060810190811067ffffffffffffffff8211176101bb57604052565b6040810190811067ffffffffffffffff8211176101bb57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101bb57604052565b6040519061022b6101c0836101f8565b565b6040519061022b610140836101f8565b6040519061022b60a0836101f8565b6040519061022b60c0836101f8565b67ffffffffffffffff81116101bb57601f01601f191660200190565b604051906102866020836101f8565b60008252565b60005b83811061029f5750506000910152565b818101518382015260200161028f565b90601f19601f6020936102cd8151809281875287808801910161028c565b0116010190565b9060206102e59281815201906102af565b90565b3461016a57600060031936011261016a57610347604080519061030b81836101f8565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102af565b0390f35b67ffffffffffffffff81160361016a57565b359061022b8261034b565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103938161034b565b60243567ffffffffffffffff811161016a576103b3903690600401610368565b6103d18267ffffffffffffffff166000526004602052604060002090565b9081546001600160a01b036103fb6103ef836001600160a01b031690565b6001600160a01b031690565b161561051857906103479361049461049a949361044461041e6080860186611c3c565b61042b6020880188611c3c565b90501591826104ff575b61043e89611e48565b87613100565b9461044d611f4d565b6040860161045b8188611c8d565b90506104ac575b506104866040880191825160608a0194610480600287519201611ce1565b9161389b565b9092525260e81c61ffff1690565b90613ea6565b60405190815292839250602083019150565b6104f9915060206104db6104c96104d46104cf6104c9868d611c8d565b90611fc9565b611fd7565b938a611c8d565b01356104ec60208a015161ffff1690565b9060e08a01519288613675565b38610462565b915061050e6040880188611c8d565b9050151591610435565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105af5781359367ffffffffffffffff85116105ac57506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105c136610554565b906105ca6147a2565b6000915b8083106105d757005b6105e2838284611fe1565b926105ec84612021565b9367ffffffffffffffff851690811580156109b2575b801561099c575b8015610983575b61094b57948561084861083e61067d60c061064b98999a019461066361065d610639888861204d565b61065560e08a949394019d8e8b61204d565b94909236916120a1565b9236916120a1565b906147e0565b67ffffffffffffffff166000526004602052604060002090565b966106c161068d60208601611fd7565b89906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61071e6106d06040860161202b565b89547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178955565b61077d61072d60608601612103565b89547fff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e89190911b7effff000000000000000000000000000000000000000000000000000000000016178955565b6107dc61078c6080860161210d565b89547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178955565b61083861082e6107ee60a08701612043565b9661082860018c0198899063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b8661204d565b9060038b01612202565b8361204d565b9060028801612202565b61010081019161085a6103ef84611fd7565b15610921576001956108db610907926108937fe24d61e75f1236506b973f38fe122980e9c3d9e27f5c6721a85b921f70c512c596611fd7565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6108f66108ec610120850185611c3c565b90600484016122c9565b5460a01c67ffffffffffffffff1690565b610916604051928392836124f1565b0390a20191906105ce565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5063ffffffff61099560a08301612043565b1615610610565b5060ff6109ab6040830161202b565b1615610609565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610602565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610a1260043561034b565b6020610a28602435610a23816109e5565b612694565b6001600160a01b0360405191168152f35b3461016a57610a4736610554565b90610a5a6003546001600160a01b031690565b9160005b818110610a6757005b610a786103ef6104cf838587612733565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa908115610b5b576001948891600093610b2b575b5082610ae0575b5050505001610a5e565b610aeb91839161493b565b6040519081526001600160a01b038716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610ad6565b610b4d91935060203d8111610b54575b610b4581836101f8565b810190612743565b9138610acf565b503d610b3b565b612688565b60ff81160361016a57565b359061022b82610b60565b3461016a57604060031936011261016a5760043567ffffffffffffffff811161016a573660238201121561016a57806004013567ffffffffffffffff811161016a57366024828401011161016a5761034791610be0916024803592610bda84610b60565b0161281a565b604051918291826102d4565b906020808351928381520192019060005b818110610c0a5750505090565b82516001600160a01b0316845260209384019390920191600101610bfd565b906102e59160208152610c486020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015161ffff1660808201526080820151151560a082015260a082015163ffffffff1660c082015260c08201516001600160a01b031660e0820152610120610cde610cc860e0850151610140610100860152610160850190610bec565b610100850151601f198583030184860152610bec565b92015190610140601f19828503019101526102af565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d1a8161034b565b6060610120604051610d2b8161019e565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e0820152826101008201520152166000526004602052610347610d806040600020611e48565b60405191829182610c29565b3461016a57600060031936011261016a57610da5611b7a565b50604051610db2816101c0565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405180916103478260608101926001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a576000546001600160a01b0381163303610e9a577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022b826109e5565b8015150361016a57565b359061022b82610ef6565b3461016a57606060031936011261016a576000604051610f2a816101c0565b600435610f36816109e5565b8152602435610f4481610ef6565b6020820190815260443590610f58826109e5565b60408301918252610f676147a2565b6001600160a01b0383511615918215611087575b50811561107c575b506110545780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39061103f611b99565b61104e60405192839283614a2e565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610f83565b516001600160a01b031615915038610f7b565b3461016a57608060031936011261016a576110b660043561034b565b60243567ffffffffffffffff811161016a576110d6903690600401610368565b604435906110e56064356109e5565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b5b576000916119da575b5061199d5760025460a01c60ff16611973576111c7740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6111e760043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b03606435161561194957805461120d6103ef6001600160a01b03831681565b330361191f57836112216080850185611c3c565b939060208601946112328688611c3c565b9050159081611906575b61124584611e48565b6112529390600435613100565b9260a01c67ffffffffffffffff1661126990612956565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617825590835163ffffffff16602085015161ffff166040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015296876034810103601f198101895261130b90896101f8565b6040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b166020820152601481529761134b60348a6101f8565b6113558a80611c3c565b86549a9161136b9160ff60e08e901c169161281a565b9060a08a0151928c604081016113819082611c8d565b61138b91506129b7565b9561139591611c3c565b9690976113a061021b565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a0152600060c08a015260e089015261142260048801611d88565b610100890152610120880152610140870152610160860152610180850152369061144b926127e3565b6101a0830152611459611f4d565b906114676040880188611c8d565b9050151591611490916114b5936118b6575b604087015190610480600260608a01519201611ce1565b60608601528060408601526114af60808601516001600160a01b031690565b90614abd565b60c08201526114dd83866114d56114ca612a06565b9760e81c61ffff1690565b600435613ea6565b63ffffffff909116606084015260208601939184521161188c57611502825186614b4b565b61150f6040860186611c8d565b90506117d8575b611522819592956155a4565b808552602081519101209061153b604085015151612a77565b9460408101958652606060009401935b6040860151805182101561172057602061157e6103ef6103ef611571866115c596612a63565b516001600160a01b031690565b61158c8460608b0151612a63565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612ac0565b03915afa8015610b5b576001600160a01b03916000916116f2575b5016801561169957906000878b9387838861164361160c8860608f61160490611fd7565b980151612a63565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612c0b565b03925af18015610b5b578161167191600194600091611678575b508a519061166b8383612a63565b52612a63565b500161154b565b611693913d8091833e61168b81836101f8565b810190612b23565b3861165d565b6105506116ad6115718460408b0151612a63565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611713915060203d8111611719575b61170b81836101f8565b810190612673565b386115e0565b503d611701565b505086907f6f70dc82f616ef8eb515e4e283343d4473bb3f0f4860db6938aa30aeaa17b74d67ffffffffffffffff610347966117988a61177561176f60408b9a015167ffffffffffffffff1690565b93611fd7565b95519651905190604051948594169767ffffffffffffffff600435169785612e8a565b0390a46117c87fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b60016117e76040870187611c8d565b9050036118625761184f6118396118046104c96040890189611c8d565b60c086015180511561185557905b602087015161ffff169060e08801519260643591611834600435913690612a25565b614e36565b6101808301519061184982612a56565b52612a56565b50611516565b5061014084015190611812565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119006118ce6104cf6104c960408c018c611c8d565b60206118e06104c960408d018d611c8d565b01356118f160208a015161ffff1690565b9060e08a015192600435613675565b90611479565b90506119156040880188611c8d565b905015159061123c565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105506004359067ffffffffffffffff60249216600452565b6119fc915060203d602011611a02575b6119f481836101f8565b810190612941565b38611171565b503d6119ea565b3461016a57602060031936011261016a5767ffffffffffffffff600435611a2f8161034b565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611a735760405167ffffffffffffffff9091168152602090f35b612117565b3461016a57602060031936011261016a576001600160a01b03600435611a9d816109e5565b611aa56147a2565b16338114611b0a57807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611b5060043561034b565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611b87826101c0565b60006040838281528260208201520152565b611ba1611b7a565b50604051611bae816101c0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660208201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016604082015290565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611d1357505061022b925003836101f8565b84546001600160a01b0316835260019485019487945060209093019201611cfe565b90600182811c92168015611d7e575b6020831014611d4f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611d44565b9060405191826000825492611d9c84611d35565b8084529360018116908115611e085750600114611dc1575b5061022b925003836101f8565b90506000929192526020600020906000915b818310611dec57505090602061022b9282010138611db4565b6020919350806001915483858901015201910190918492611dd3565b6020935061022b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611db4565b90611f2d6004611e5661022d565b93611ebd611eb48254611e7f611e72826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c16604089015261ffff60e882901c16606089015260f81c90565b15156080870152565b611f04611ef46001830154611ee5611ed88263ffffffff1690565b63ffffffff1660a08a0152565b60201c6001600160a01b031690565b6001600160a01b031660c0870152565b611f1060028201611ce1565b60e0860152611f2160038201611ce1565b61010086015201611d88565b610120830152565b67ffffffffffffffff81116101bb5760051b60200190565b60405190611f5c6020836101f8565b6000808352366020840137565b90611f7382611f35565b611f8060405191826101f8565b828152601f19611f908294611f35565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015611fd25790565b611f9a565b356102e5816109e5565b9190811015611fd25760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec18136030182121561016a570190565b356102e58161034b565b356102e581610b60565b63ffffffff81160361016a57565b356102e581612035565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b9291906120ad81611f35565b936120bb60405195866101f8565b602085838152019160051b810192831161016a57905b8282106120dd57505050565b6020809183356120ec816109e5565b8152019101906120d1565b61ffff81160361016a57565b356102e5816120f7565b356102e581610ef6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611a7357565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611a7357565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611a7357565b81810292918115918404141715611a7357565b8181106121f6575050565b600081556001016121eb565b9067ffffffffffffffff83116101bb576801000000000000000083116101bb578154838355808410612266575b5090600052602060002060005b8381106122495750505050565b600190602084359461225a866109e5565b0193818401550161223c565b61227e908360005284602060002091820191016121eb565b3861222f565b9190601f811161229357505050565b61022b926000526020600020906020601f840160051c830193106122bf575b601f0160051c01906121eb565b90915081906122b2565b90929167ffffffffffffffff81116101bb576122ef816122e98454611d35565b84612284565b6000601f821160011461234d57819061233e939495600092612342575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b01359050388061230c565b601f1982169461236284600052602060002090565b91805b8781106123bb575083600195969710612383575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612379565b90926020600181928686013581550194019101612365565b359061022b826120f7565b359061022b82612035565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124565750505090565b9091926020806001926001600160a01b038735612472816109e5565b168152019401929101612449565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b601f8260209493601f19938186528686013760008582860101520116010190565b67ffffffffffffffff6102e5939216815260406020820152612527604082016125198461035d565b67ffffffffffffffff169052565b61254661253660208401610eeb565b6001600160a01b03166060830152565b61255f61255560408401610b6b565b60ff166080830152565b61257961256e606084016123d3565b61ffff1660a0830152565b61259161258860808401610f00565b151560c0830152565b6125ad6125a060a084016123de565b63ffffffff1660e0830152565b6126426126156125d76125c360c08601866123e9565b61014061010087015261018086019161243c565b6125e460e08601866123e9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08684030161012087015261243c565b926126376126266101008301610eeb565b6001600160a01b0316610140850152565b610120810190612480565b916101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828603019101526124d0565b9081602091031261016a57516102e5816109e5565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610b5b576001600160a01b039160009161271657501690565b61272f915060203d6020116117195761170b81836101f8565b1690565b9190811015611fd25760051b0190565b9081602091031261016a575190565b9160206102e59381815201916124d0565b60ff166020039060ff8211611a7357565b90929192831161016a579190565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b3590602081106127b6575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9291926127ef8261025b565b916127fd60405193846101f8565b82948184528183011161016a578281602093846000960137010152565b9160ff81169060208210612876575b50810361283c57906102e59136916127e3565b906128726040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612752565b0390fd5b6020831161290c57602083036128295790506128aa6128a460ff61289c84969596612763565b168585612774565b906127a8565b6128d657916128cf91816128c96128c36102e596612763565b60ff1690565b91612790565b36916127e3565b506128726040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612752565b6040517f3aeba39000000000000000000000000000000000000000000000000000000000815280612872858760048401612752565b9081602091031261016a57516102e581610ef6565b67ffffffffffffffff1667ffffffffffffffff8114611a735760010190565b6040519060c0820182811067ffffffffffffffff8211176101bb57604052606060a0836000815282602082015282604082015282808201528260808201520152565b906129c182611f35565b6129ce60405191826101f8565b828152601f196129de8294611f35565b019060005b8281106129ef57505050565b6020906129fa612975565b828285010152016129e3565b60405190612a13826101c0565b60606040838281528260208201520152565b919082604091031261016a57604051612a3d816101dc565b60208082948035612a4d816109e5565b84520135910152565b805115611fd25760200190565b8051821015611fd25760209160051b010190565b90612a8182611f35565b612a8e60405191826101f8565b828152601f19612a9e8294611f35565b019060005b828110612aaf57505050565b806060602080938501015201612aa3565b60409067ffffffffffffffff6102e5949316815281602082015201906102af565b81601f8201121561016a578051612af78161025b565b92612b0560405194856101f8565b8184526020828401011161016a576102e5916020808501910161028c565b9060208282031261016a57815167ffffffffffffffff811161016a576102e59201612ae1565b9080602083519182815201916020808360051b8301019401926000915b838310612b7557505050505090565b9091929394602080612bfc83601f1986600196030187528951908151815260a0612beb612bd9612bc7612bb58887015160c08a88015260c08701906102af565b604087015186820360408801526102af565b606086015185820360608701526102af565b608085015184820360808601526102af565b9201519060a08184039101526102af565b97019301930191939290612b66565b9193906102e59593612e07612e1f9260a08652612c3560a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612dd4612d9e612d68612d32612cfc612cc88c61026060e08a0151916101c061018082015201906102af565b6101008801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001888f01526102af565b6101208701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101c08e01526102af565b6101408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101e08d01526102af565b6101608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016102008c01526102af565b6101808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016102208b0152612b49565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016102408801526102af565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102af565b9080602083519182815201916020808360051b8301019401926000915b838310612e5d57505050505090565b9091929394602080612e7b83601f19866001960301875289516102af565b97019301930191939290612e4e565b9493916001600160a01b03612ead921686526080602087015260808601906102af565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612eef57505050506102e59394506060818403910152612e31565b90919295602080612f5083601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102af565b980192019201909291612ed1565b60405190610100820182811067ffffffffffffffff8211176101bb57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612fe3575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a578261303d918501612ae1565b92602081015161304c81612035565b92604082015167ffffffffffffffff811161016a576102e59201612ae1565b60409067ffffffffffffffff6102e5959316815281602082015201916124d0565b9060ff6130a66020929594956040855260408501906102af565b9416910152565b9060018201809211611a7357565b9060028201809211611a7357565b6001019081600111611a7357565b9060148201809211611a7357565b90600c8201809211611a7357565b91908201809211611a7357565b939192909261310d612f5e565b60048310158061351b575b156133d0575090613128916158f4565b9260c08401908151516132fb575b5050604083019081515160005b81811061326657505081515115908161323c575b506131ab575b505b608082016001600160a01b0361317c82516001600160a01b031690565b161561318757505090565b61319e60c06102e59301516001600160a01b031690565b6001600160a01b03169052565b906101008194939401926131c0845151611f69565b83526131cd845151612a77565b946060810195865260005b8551805182101561322e579061320f6131f661157183600195612a63565b613201838951612a63565b906001600160a01b03169052565b61322781895161321d610277565b61166b8383612a63565b50016131d8565b50509350935090503861315d565b90508061324b575b1538613157565b5063ffffffff61325f845163ffffffff1690565b1615613244565b61326f816130ad565b82811061327f5750600101613143565b61328d611571838751612a63565b6001600160a01b036132a66103ef611571858a51612a63565b9116146132b55760010161326f565b6105506132c6611571848851612a63565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b61330f61330b6080860151151590565b1590565b613399575080600061335f925161332a604087015160ff1690565b9060405194859283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161308c565b0381305afa918215610b5b5760009261337c575b50523880613136565b6133929192503d806000833e61168b81836101f8565b9038613373565b7f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b9491613428926000926133f16103ef6103ef6002546001600160a01b031690565b91604051958694859384937f9cc199960000000000000000000000000000000000000000000000000000000085526004850161306b565b03915afa8015610b5b57600091829083926134f2575b5060a086019190915263ffffffff16845260c084015b81815281516134a4575b505080613489575b61315f5761347f61010082015180604085015251612a77565b606083015261315f565b5063ffffffff61349d835163ffffffff1690565b1615613466565b60006134b89261332a604087015160ff1690565b0381305afa918215610b5b576000926134d5575b5052388061345e565b6134eb9192503d806000833e61168b81836101f8565b90386134cc565b90506135149150613454923d8091833e61350c81836101f8565b810190613015565b919261343e565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061357161356b8686612782565b90612faf565b1614613118565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a5781516135ac81611f35565b926135ba60405194856101f8565b81845260208085019260051b82010192831161016a57602001905b8282106135e25750505090565b6020809183516135f1816109e5565b8152019101906135d5565b95949060009460a09467ffffffffffffffff613643956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102af565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611a735760010190565b9492949390936136a4600361369e8367ffffffffffffffff166000526004602052604060002090565b01611ce1565b946001600160a01b036136b8818316612694565b16926040517f01ffc9a70000000000000000000000000000000000000000000000000000000081526020818061371560048201907f3317103100000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610b5b5760009161387c575b5015613872579061376f600095949392604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016135fc565b03915afa928315610b5b5760009361384d575b508251156138485761379f61379a84518451906130f3565b611f69565b6000918293835b86518110156137f7576137bc6115718289612a63565b6001600160a01b038116156137eb57906137e56001926132016137de89613648565b9888612a63565b016137a6565b509450600180956137e5565b509193909450613808575b50815290565b60005b8151811015613840578061383a61382761157160019486612a63565b61320161383387613648565b9688612a63565b0161380b565b505038613802565b915090565b61386b9193503d806000833e61386381836101f8565b810190613578565b9138613782565b5050505050915090565b613895915060203d602011611a02576119f481836101f8565b38613729565b939192936138b76138af82518651906130f3565b8651906130f3565b906138ca6138c483611f69565b92612a77565b94600096875b835189101561393057886139266139196001936139016138f76115718e9f9d9e9d8b612a63565b613201838c612a63565b61391f61390e858c612a63565b519180938491613648565b9c612a63565b528b612a63565b50019796956138d0565b959250929350955060005b86518110156139bd576139516115718289612a63565b60006001600160a01b038216815b888110613991575b505090600192911561397b575b500161393b565b61398b906132016137de89613648565b38613974565b816139a26103ef611571848c612a63565b146139af5760010161395f565b506001915081905038613967565b509390945060005b8551811015613a4e576139db6115718288612a63565b60006001600160a01b038216815b878110613a22575b5050906001929115613a05575b50016139c5565b613a1c90613201613a1588613648565b9787612a63565b386139fe565b81613a336103ef611571848b612a63565b14613a40576001016139e9565b5060019150819050386139f1565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101bb5760405260606080836000815260006020820152600060408201526000838201520152565b90613aa382611f35565b613ab060405191826101f8565b828152601f19613ac08294611f35565b019060005b828110613ad157505050565b602090613adc613a5a565b82828501015201613ac5565b9081606091031261016a578051613afe816120f7565b9160406020830151613b0f81612035565b9201516102e581612035565b9160209082815201919060005b818110613b355750505090565b9091926040806001926001600160a01b038735613b51816109e5565b16815260208781013590820152019401929101613b28565b949391929067ffffffffffffffff16855260806020860152613be0613ba3613b918580612480565b60a060808a01526101208901916124d0565b613bb06020860186612480565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a01526124d0565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022b95613ca4613c7b83606097613ce3978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082613cd59a0301910152613b1b565b91613c9a613c8a888301610eeb565b6001600160a01b031660e08d0152565b6080810190612480565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c01526124d0565b9087820360408901526102af565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611a7357565b908160a091031261016a578051916020820151613d2481612035565b916040810151613d3381612035565b9160806060830151613d44816120f7565b9201516102e581610ef6565b9260c0946001600160a01b039167ffffffffffffffff61ffff95846102e59b9a9616885216602087015260408601521660608401521660808201528160a082015201906102af565b9081606091031261016a578051613afe81612035565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611a7357565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611a7357565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611a7357565b91908203918211611a7357565b919082608091031261016a578151613e5981612035565b916020810151916060604083015192015190565b8115613e77570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613ee0613edb613ed6885151613ece60408c01809c611c8d565b9190506130f3565b6130bb565b613a99565b9660009586955b88518051881015614144576103ef6103ef6115718a613f0594612a63565b613f5260206060880192613f1a8b8551612a63565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612ac0565b03915afa8015610b5b576001600160a01b0391600091614126575b501680156140d2579060608e9392613f868b8451612a63565b5190613f9760208b015161ffff1690565b958b613fd2604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613b69565b03915afa8015610b5b5760019361406f938b8f8f95600080958197614078575b509083929161ffff61401a85614013611571614063996140699d9e51612a63565b9451612a63565b519161403661402761023d565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261166b8383612a63565b50613cee565b99613cee565b96019596613ee7565b6140699750611571965084939291509361ffff61401a826140136140b56140639960603d81116140cb575b6140ad81836101f8565b810190613ae8565b9c9196909c9d5050505050505090919293613ff2565b503d6140a3565b610550886140e46115718c8f51612a63565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b61413e915060203d81116117195761170b81836101f8565b38613f6d565b50919a9496929395509897968a61415b8187611c8d565b90506143f0575b5050865161416f90613dae565b998561417e6020870187611c3c565b9261418a915087611c8d565b90506141969289615b9a565b6141a08b89612a63565b526141ab8a88612a63565b506141b68a88612a63565b516020015163ffffffff166141ca91613cee565b906141d58a88612a63565b516040015163ffffffff166141e991613cee565b916141f261023d565b33815290600060208301819052604083015261ffff166060820152614215610277565b6080820152865161422590613ddb565b906142308289612a63565b5261423b9087612a63565b506002546001600160a01b03169261425590606001611fd7565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610b5b576000966000976000946000926143b3575b506000965b86518810156143a65761432f6001916142f7886142f287612146565b613e6d565b61431060606143068d8d612a63565b51019182516121d8565b9052858a14614337575b60606143268b8b612a63565b510151906130f3565b9701966142d6565b8b8873eba517d2000000000000000000000000000000006001600160a01b0361436a60808c01516001600160a01b031690565b1603614378575b505061431a565b6142f261438492612176565b61439d60606143938d8d612a63565b51019182516130f3565b90528b88614371565b9598509694955050505050565b92985050506143db91925060803d6080116143e9575b6143d381836101f8565b810190613e42565b9197909392909190386142d1565b503d6143c9565b610a236103ef6104cf6104c9614409948a989698611c8d565b926001600160a01b03600091515194169060e0880190815161442961023d565b6001600160a01b0385168152908260208301528260408301528260608301526080820152614457878d612a63565b52614462868c612a63565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610b5b578f948c89968f96948d948f968891614783575b5061467c575b50505050505015614523575b61450261451461451b9561450e602061450261450e97604097612a63565b51015163ffffffff1690565b90613cee565b958b612a63565b90388a614162565b50506145a99160608c6145546104cf6104c961454d6103ef6103ef6002546001600160a01b031690565b938b611c8d565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610b5b5761450261451460409261450e60206145028f8b9061451b9b61450e9a60008060009261463e575b63ffffffff92935061462b9060606145f38888612a63565b5101946146208a6146048a8a612a63565b51019160406146138b8b612a63565b51019063ffffffff169052565b9063ffffffff169052565b16905297509750505050955050506144e4565b50505063ffffffff61466a61462b9260603d606011614675575b61466281836101f8565b810190613d98565b9093509150826145db565b503d614658565b8495985060a09697506146c660206146bc6060826146b36104c96146ac6104cf6104c98b6146ff9c9d9e9f611c8d565b998d611c8d565b01359901611fd7565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613d50565b03915afa8015610b5b578592828c939181908294614747575b5061473b9060606147298888612a63565b51019261462060206146048a8a612a63565b5288888f8c81386144d8565b91505061473b9250614771915060a03d60a01161477c575b61476981836101f8565b810190613d08565b949192919050614718565b503d61475f565b61479c915060203d602011611a02576119f481836101f8565b386144d2565b6001600160a01b036001541633036147b657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916147ee8151846130f3565b9283156149115760005b848110614806575050505050565b818110156148f65761481b6115718286612a63565b6001600160a01b03811680156148cc57614834836130ad565b878110614846575050506001016147f8565b848110156148a9576001600160a01b03614863611571838a612a63565b16821461487257600101614834565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036148c76115716148c18885613e35565b89612a63565b614863565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b61490c6115716149068484613e35565b85612a63565b61481b565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b906001600160a01b03614a009392604051938260208601947fa9059cbb0000000000000000000000000000000000000000000000000000000086521660248601526044850152604484526149906064856101f8565b166000806040938451956149a486886101f8565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614a25573d6149f16149e88261025b565b945194856101f8565b83523d6000602085013e61626e565b805180614a0b575050565b81602080614a209361022b9501019101612941565b615da7565b6060925061626e565b91606061022b929493614a6e8160c08101976001600160a01b036040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b90614aa08261025b565b614aad60405191826101f8565b828152601f19611f90829461025b565b918251601481029080820460141490151715611a7357614adf614ae4916130c9565b6130d7565b90614af6614af1836130e5565b614a96565b906014614b0283612a56565b5360009260215b8651851015614b34576014600191614b24611571888b612a63565b60601b8187015201940193614b09565b919550936020935060601b90820152828152012090565b90614b5b6103ef60608401611fd7565b614b8a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff936040810190611c8d565b9050614c02575b614b9b8251613ddb565b9260005b848110614bad575050505050565b808260019214614bfd576060614bc38287612a63565b5101518015614bf757614bf190614beb614bdd8489612a63565b51516001600160a01b031690565b8661493b565b01614b9f565b50614bf1565b614bf1565b9150614c0e8151613e08565b91614c1c614bdd8484612a63565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610b5b57600091614cb5575b50614c95575b50614b91565b614caf906060614ca58686612a63565b510151908361493b565b38614c8f565b614cce915060203d602011611a02576119f481836101f8565b38614c89565b60405190614ce1826101dc565b60606020838281520152565b919060408382031261016a5760405190614d06826101dc565b8193805167ffffffffffffffff811161016a5782614d25918301612ae1565b835260208101519167ffffffffffffffff831161016a57602092614d499201612ae1565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a576102e59201614ced565b9060806001600160a01b0381614d93855160a0865260a08601906102af565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206102e5928181520190614d74565b919060408382031261016a57825167ffffffffffffffff811161016a57602091614e03918501614ced565b92015190565b61ffff614e226102e59593606084526060840190614d74565b9316602082015260408184039101526102af565b90919293614e42612975565b50602082019081511561533f57614e696103ef610a236103ef86516001600160a01b031690565b956001600160a01b03871692831580156152b4575b61526b57614ee2815191614e90614cd4565b505186516001600160a01b031690614ecd614ea961023d565b8b815267ffffffffffffffff8b166020820152956001600160a01b03166040870152565b60608501526001600160a01b03166080840152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610b5b5760009161524c575b5015615156575091614f8e96979160008094604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614e09565b03925af18015610b5b5760009460009161510d575b509460006150ac9361504a61500461503c95611571614ff69a965b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b03601f1981018c528b6101f8565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b03601f1981018652856101f8565b61507761506d84519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161308c565b0381305afa928315610b5b576000936150ed575b5060200151936150ce61024c565b958652602086015260408501526060840152608083015260a082015290565b6020919350615106903d806000833e61168b81836101f8565b92906150c0565b614ff695506115719691506150ac9361504a61500461503c956151436000953d8088833e61513b81836101f8565b810190614dd8565b9b909b969b9a5050955050509350614fa3565b979161ffff91959793501661522257516151f85760006151a393604051809581927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614dc7565b038183855af1918215610b5b57614ff69561504a61500460009361157161503c976150ac9987916151d6575b5096614fbe565b6151f291503d8089833e6151ea81836101f8565b810190614d4e565b386151cf565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b615265915060203d602011611a02576119f481836101f8565b38614f46565b61055061527f86516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818c5afa908115610b5b57600091615320575b5015614e7e565b615339915060203d602011611a02576119f481836101f8565b38615319565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b959261543b947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906154576020928281519485920161028c565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b1681526154a3825180936020898501910161028c565b019160f81b16838201526154c182518093602060028501910161028c565b01019160f81b16838201526154e082518093602060028501910161028c565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006102e59f9e9c9860f81b168152615555825180936020898501910161028c565b019160f01b168382015261557382518093602060038501910161028c565b01019160f01b1683820152615591825180936020898501910161028c565b01019160f01b1660028201520190615444565b60e081019060ff825151116158c55761010081019060ff825151116158965761012081019260ff845151116158675761014082019060ff825151116158385761016083019461ffff86515111615809576101808401946001865151116157da576101a085019261ffff845151116157a957855167ffffffffffffffff16602087015167ffffffffffffffff169060408801516156479067ffffffffffffffff1690565b97606081015161565a9063ffffffff1690565b90608081015161566d9063ffffffff1690565b60a082015161ffff169160c00151926040519b8c96602088019661569097615369565b03601f19810187526156a290876101f8565b519081516156b09060ff1690565b9051805160ff1693519081516156c69060ff1690565b9060405195869560208701956156db9661545b565b03601f19810182526156ed90826101f8565b6060945190815115156157799761575b6102e59861577f97615779966157699561578d575b505196615720885160ff1690565b93519161572f835161ffff1690565b9161573c825161ffff1690565b90519361574b855161ffff1690565b936040519b8c9860208a016154e6565b03601f1981018552846101f8565b6040519687956020870190615444565b90615444565b03601f1981018352826101f8565b6157a291925061579c90612a56565b51615f6d565b9038615712565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105506024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b906158fd612f5e565b9160118210615b135780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603615aa05750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61597c81611f69565b6040860190815261598c82612a77565b906060870191825260005b838110615a5457505050506159ff83836159f56159e96159df6159d86159c5615a099887615a139c9b616103565b6001600160a01b0390911660808d015290565b85856161d9565b92919036916127e3565b60a08a01528383616241565b94919036916127e3565b60c08801526161d9565b93919036916127e3565b60e08401528103615a22575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b80600191615a99615a83615a7c615a6f615a939a8d8d616103565b9190613201868a51612a63565b8b8b6161d9565b9391889a919a51949a36916127e3565b92612a63565b5201615997565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b9081602091031261016a57516102e5816120f7565b9261ffff6102e5959367ffffffffffffffff615b8c94168652166020850152608060408501526080840190610bec565b9160608184039101526102af565b9091615ba4613a5a565b50615bc38267ffffffffffffffff166000526004602052604060002090565b93615bd3855460ff9060e01c1690565b90615c73615c65615c5c6080840194615c48615c22615c1a615c0d6001615c018b516001600160a01b031690565b9e015463ffffffff1690565b885163ffffffff1661450e565b9a60756130f3565b97615c56615c4e60ff615c3c60a08b019c8d5151906130f3565b951694615c48866121a8565b906130f3565b93604f6130f3565b906121d8565b63ffffffff1690565b92516001600160a01b031690565b6001600160a01b03811673eba517d20000000000000000000000000000000003615ced57505061ffff9250615cdf90615cd26000935b5195615cc5615cb661023d565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b90615d036103ef6020936001600160a01b031690565b6040615d138484015161ffff1690565b92015191855196615d53604051988995869485947fe962e69e00000000000000000000000000000000000000000000000000000000865260048601615b5c565b03915afa908115610b5b57615cd2615cdf9261ffff95600091615d78575b5093615ca9565b615d9a915060203d602011615da0575b615d9281836101f8565b810190615b47565b38615d71565b503d615d88565b15615dae57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b966002615f4197615f0e60226102e59f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615f0e9f82615f0e9c615f159c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615ebf825180936020898501910161028c565b019160f81b1683820152615edd82518093602060238501910161028c565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190615444565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116160d457604081019160ff835151116160a557606082019160ff8351511161607657608081019260ff845151116160475760a0820161ffff81515111616018576102e59461577f9351945191615fcf835160ff1690565b975191615fdd835160ff1690565b945190615feb825160ff1690565b905193615ff9855160ff1690565b935196616008885161ffff1690565b966040519c8d9b60208d01615e32565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b9291909260018201918483116161a75781013560001a82811561619c57506014810361616f57820193841161613b57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b919060028201918183116161a7578381013560f01c016002019281841161620d5791839161620693612790565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b919060018201918183116161a7578381013560001a016001019281841161620d5791839161620693612790565b919290156162e95750815115616282575090565b3b1561628b5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156162fc5750805190602001fd5b612872906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016102d456fea164736f6c634300081a000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, staticConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRamp{address: address, abi: *parsed, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

type OnRamp struct {
	address common.Address
	abi     abi.ABI
	OnRampCaller
	OnRampTransactor
	OnRampFilterer
}

type OnRampCaller struct {
	contract *bind.BoundContract
}

type OnRampTransactor struct {
	contract *bind.BoundContract
}

type OnRampFilterer struct {
	contract *bind.BoundContract
}

type OnRampSession struct {
	Contract     *OnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampCallerSession struct {
	Contract *OnRampCaller
	CallOpts bind.CallOpts
}

type OnRampTransactorSession struct {
	Contract     *OnRampTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRaw struct {
	Contract *OnRamp
}

type OnRampCallerRaw struct {
	Contract *OnRampCaller
}

type OnRampTransactorRaw struct {
	Contract *OnRampTransactor
}

func NewOnRamp(address common.Address, backend bind.ContractBackend) (*OnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRamp{address: address, abi: abi, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

func NewOnRampCaller(address common.Address, caller bind.ContractCaller) (*OnRampCaller, error) {
	contract, err := bindOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampCaller{contract: contract}, nil
}

func NewOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampTransactor, error) {
	contract, err := bindOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampTransactor{contract: contract}, nil
}

func NewOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampFilterer, error) {
	contract, err := bindOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampFilterer{contract: contract}, nil
}

func bindOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OnRamp *OnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.OnRampCaller.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transfer(opts)
}

func (_OnRamp *OnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transfer(opts)
}

func (_OnRamp *OnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (OnRampDestChainConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(OnRampDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDestChainConfig)).(*OnRampDestChainConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDestChainConfig(destChainSelector uint64) (OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDynamicConfig)).(*OnRampDynamicConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetExpectedNextMessageNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getExpectedNextMessageNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetExpectedNextMessageNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextMessageNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetExpectedNextMessageNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextMessageNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCaller) GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampStaticConfig)).(*OnRampStaticConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRamp *OnRampSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) ValidateDestChainAddress(opts *bind.CallOpts, rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "validateDestChainAddress", rawAddress, addressBytesLength)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_OnRamp *OnRampSession) ValidateDestChainAddress(rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	return _OnRamp.Contract.ValidateDestChainAddress(&_OnRamp.CallOpts, rawAddress, addressBytesLength)
}

func (_OnRamp *OnRampCallerSession) ValidateDestChainAddress(rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	return _OnRamp.Contract.ValidateDestChainAddress(&_OnRamp.CallOpts, rawAddress, addressBytesLength)
}

func (_OnRamp *OnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OnRamp *OnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_OnRamp *OnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OnRamp *OnRampSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactorSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRamp *OnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_OnRamp *OnRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts, feeTokens)
}

func (_OnRamp *OnRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts, feeTokens)
}

type OnRampCCIPMessageSentIterator struct {
	Event *OnRampCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCCIPMessageSent)
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
		it.Event = new(OnRampCCIPMessageSent)
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

func (it *OnRampCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCCIPMessageSent struct {
	DestChainSelector uint64
	MessageNumber     uint64
	MessageId         [32]byte
	FeeToken          common.Address
	EncodedMessage    []byte
	Receipts          []OnRampReceipt
	VerifierBlobs     [][]byte
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, messageNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, messageNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCCIPMessageSent)
				if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error) {
	event := new(OnRampCCIPMessageSent)
	if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampConfigSetIterator struct {
	Event *OnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampConfigSet)
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
		it.Event = new(OnRampConfigSet)
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

func (it *OnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampConfigSet struct {
	StaticConfig  OnRampStaticConfig
	DynamicConfig OnRampDynamicConfig
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampConfigSetIterator{contract: _OnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseConfigSet(log types.Log) (*OnRampConfigSet, error) {
	event := new(OnRampConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampDestChainConfigSetIterator struct {
	Event *OnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampDestChainConfigSet)
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
		it.Event = new(OnRampDestChainConfigSet)
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

func (it *OnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampDestChainConfigSet struct {
	DestChainSelector uint64
	MessageNumber     uint64
	Config            OnRampDestChainConfigArgs
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampDestChainConfigSetIterator{contract: _OnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampDestChainConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error) {
	event := new(OnRampDestChainConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeeTokenWithdrawnIterator struct {
	Event *OnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeeTokenWithdrawn)
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
		it.Event = new(OnRampFeeTokenWithdrawn)
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

func (it *OnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeeTokenWithdrawnIterator{contract: _OnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeeTokenWithdrawn)
				if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error) {
	event := new(OnRampFeeTokenWithdrawn)
	if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferRequestedIterator struct {
	Event *OnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferRequested)
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
		it.Event = new(OnRampOwnershipTransferRequested)
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

func (it *OnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferRequestedIterator{contract: _OnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferRequested)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error) {
	event := new(OnRampOwnershipTransferRequested)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferredIterator struct {
	Event *OnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferred)
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
		it.Event = new(OnRampOwnershipTransferred)
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

func (it *OnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferredIterator{contract: _OnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferred)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error) {
	event := new(OnRampOwnershipTransferred)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OnRampCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x6f70dc82f616ef8eb515e4e283343d4473bb3f0f4860db6938aa30aeaa17b74d")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe24d61e75f1236506b973f38fe122980e9c3d9e27f5c6721a85b921f70c512c5")
}

func (OnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (OnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRamp *OnRamp) Address() common.Address {
	return _OnRamp.address
}

type OnRampInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (OnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextMessageNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ValidateDestChainAddress(opts *bind.CallOpts, rawAddress []byte, addressBytesLength uint8) ([]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	Address() common.Address
}
