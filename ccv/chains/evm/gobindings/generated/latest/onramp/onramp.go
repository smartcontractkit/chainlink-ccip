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
	SequenceNumber       uint64
	AddressBytesLength   uint8
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc57604051616bab38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051616860908161034b82396080518181816105bf0152818161164e0152611e57015260a0518181816113cc0152611e90015260c051818181611ecc01526128180152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461011757806314a8cfa314610112578063181f5a771461010d57806320487ded1461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611db4565b611cc0565b6112fb565b611127565b611085565b611033565b610f4a565b610e7e565b610dd8565b610bf6565b610a9d565b610a2f565b61084a565b6107a9565b610217565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611e37565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f830112156102135781359367ffffffffffffffff851161021057506020808301928560051b010111610195579190565b80fd5b8380fd5b34610195576102253661019a565b61022d61314c565b60005b81811061023957005b610244818385611f23565b9061024e82611f68565b67ffffffffffffffff811690811580156105b3575b801561059d575b8015610584575b610549576102ba906102d4608086019161028b8388611f94565b94906102b460a08a01966102ac6102a2898d611f94565b9490923691612000565b923691612000565b90613243565b67ffffffffffffffff166000526004602052604060002090565b9160208601906103276102e683612056565b859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61038461033660408901611f72565b85547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178555565b61039060608801611f8a565b6103c96001860191829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6103e06103d6858a611f94565b906003880161211b565b6103f76103ed838a611f94565b906002880161211b565b60c088019161042161040884612056565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561051f57600198610501846105077f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae996104ad6104f9976104656105169a612056565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6104f06104e96104e360e08801936104d26104c8868b61219d565b9060048401612286565b5460a01c67ffffffffffffffff1690565b9a612056565b9a86611f94565b97909686611f94565b949093612056565b9461219d565b959094604051998a998a61243e565b0390a201610230565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b5063ffffffff61059660608601611f8a565b1615610271565b5060ff6105ac60408601611f72565b161561026a565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610263565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610100810190811067ffffffffffffffff82111761063257604052565b6105e6565b6060810190811067ffffffffffffffff82111761063257604052565b6040810190811067ffffffffffffffff82111761063257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761063257604052565b604051906106c06101c08361066f565b565b604051906106c060a08361066f565b604051906106c060c08361066f565b67ffffffffffffffff811161063257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061072960208361066f565b60008252565b60005b8381106107425750506000910152565b8181015183820152602001610732565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361078e8151809281875287808801910161072f565b0116010190565b9060206107a6928181520190610752565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761082660408051906107ea818361066f565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610752565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576004356108858161082a565b60243567ffffffffffffffff8111610195576108a590369060040161083c565b6108c38267ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6108fa610408845473ffffffffffffffffffffffffffffffffffffffff1690565b16156109d957908161082693610929610919608061097496018461219d565b90610923866125cd565b8461350f565b926109326126d0565b60408401610940818661273b565b9050610986575b5061096b6040860191825160608801946109656002875192016124ac565b91613c84565b909252526142dd565b60405190815292839250602083019150565b6109d3915060206109b56109a36109ae6109a96109a3868b61273b565b9061278f565b612056565b938861273b565b01356109c6602088015161ffff1690565b9060e08801519286613a44565b38610947565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610a6960043561082a565b6020610a7f602435610a7a81610a11565b6127b9565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461019557610aab3661019a565b90610acb60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610ad857005b610ae96104086109a983858761287f565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610be6576001948891600093610bb6575b5082610b5e575b5050505001610acf565b610b69918391614c91565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610b54565b610bd891935060203d8111610bdf575b610bd0818361066f565b81019061288f565b9138610b4d565b503d610bc6565b6127ad565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff81116101955736602482840101116101955761082691610c7e916024803592610c7884610beb565b01612966565b60405191829182610795565b906020808351928381520192019060005b818110610ca85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610c9b565b906107a69160208152610d0060208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015163ffffffff166080820152608082015173ffffffffffffffffffffffffffffffffffffffff1660a082015260e0610da4610d7160a085015161010060c0860152610120850190610c8a565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610c8a565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610752565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610e1c8161082a565b606060e0604051610e2c81610615565b600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152166000526004602052610826610e7260406000206125cd565b60405191829182610cd4565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610eb5611e18565b50604051610ec281610637565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152604051809161082682606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff81163303611009577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff6004356110c98161082a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff811161110d5760405167ffffffffffffffff9091168152602090f35b612060565b35906106c082610a11565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161116481610637565b60043561117081610a11565b815260243561117e8161111d565b602082019081526044359061119282610a11565b604083019182526111a161314c565b73ffffffffffffffffffffffffffffffffffffffff835116159182156112db575b5081156112d0575b506112a85780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611293611e37565b6112a260405192839283614d91565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386111ca565b5173ffffffffffffffffffffffffffffffffffffffff16159150386111c2565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761133560043561082a565b60243567ffffffffffffffff81116101955761135590369060040161083c565b60443590611364606435610a11565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff00000000000000000000000000000000169082015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610be657600091611c91575b50611c545760025460a01c60ff16611c2a57611453740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61147360043567ffffffffffffffff166000526004602052604060002090565b73ffffffffffffffffffffffffffffffffffffffff6064351615611c005780546114b361040873ffffffffffffffffffffffffffffffffffffffff831681565b3303611bd6576114c6608084018461219d565b6114cf846125cd565b6114dc929060043561350f565b9060a01c67ffffffffffffffff166114f390612aa2565b82547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617835591815163ffffffff16602083015161ffff166040513060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201528060348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526115b2908261066f565b6040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b166020820152601481526115f160348261066f565b6115fb888061219d565b865460e01c60ff169061160d92612966565b60a08701519161162060408b018b61273b565b61162a9150612b03565b9361163860208c018c61219d565b9690976116436106b0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529b67ffffffffffffffff6004351660208e015267ffffffffffffffff1660408d0152600060608d015263ffffffff1660808c015261ffff1660a08b0152600060c08b015260e08a01526116c56004880161250d565b6101008a015261012089015261014088015261016087015261018086015236906116ee9261292f565b6101a08401526116fc6126d0565b9061170a604086018661273b565b90501515916117339161176593611b86575b6040850151906109656002606088015192016124ac565b606084015280604084015261175f608084015173ffffffffffffffffffffffffffffffffffffffff1690565b90614e58565b60c0830152611772612b70565b61179261178283866004356142dd565b63ffffffff909116606087015291565b6020830152808610611b5457506117b0602082015185600435614ee6565b6117bd604085018561273b565b9050611aa0575b926117ce83615a00565b808552602081519101206117e6604084015151612be1565b9360408601948552606060009301925b604085015180518210156119f257602061183661040861040861181c8661187d96612bcd565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6118448460608a0151612bcd565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612c48565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff916000916119c4575b5016801561195e57906000611908926118bd87612056565b908b836118ce8660608d0151612bcd565b5193604051978895869485937f3bbbed4b0000000000000000000000000000000000000000000000000000000085528d8d60048701612db1565b03925af18015610be657816119369160019460009161193d575b508951906119308383612bcd565b52612bcd565b50016117f6565b611958913d8091833e611950818361066f565b810190612cab565b38611922565b61058061197261181c8460408a0151612bcd565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660049081523567ffffffffffffffff16602452604490565b6119e5915060203d81116119eb575b6119dd818361066f565b810190612798565b386118a5565b503d6119d3565b6108268480898b7f276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c9667ffffffffffffffff611a3860408b015167ffffffffffffffff1690565b611a60602085519501519551604051938493169667ffffffffffffffff60043516968461305b565b0390a4611a907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b6001611aaf604086018661273b565b905003611b2a57611b17611b01611acc6109a3604088018861273b565b60c0850151805115611b1d57905b602086015161ffff169060e08701519260643591611afc600435913690612b8f565b6151fb565b61018085015190611b1182612bc0565b52612bc0565b506117c4565b5061014086015190611ada565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f543f439800000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b9050611bd0611b9e6109a96109a360408a018a61273b565b6020611bb06109a360408b018b61273b565b0135611bc1602088015161ffff1690565b9060e088015192600435613a44565b9061171c565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105806004359067ffffffffffffffff60249216600452565b611cb3915060203d602011611cb9575b611cab818361066f565b810190612a8d565b386113fd565b503d611ca1565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611d1081610a11565b611d1861314c565b16338114611d8a57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611dee60043561082a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611e2582610637565b60006040838281528260208201520152565b611e3f611e18565b50604051611e4c81610637565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611f635760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215610195570190565b611ef4565b356107a68161082a565b356107a681610beb565b63ffffffff81160361019557565b356107a681611f7c565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b67ffffffffffffffff81116106325760051b60200190565b92919061200c81611fe8565b9361201a604051958661066f565b602085838152019160051b810192831161019557905b82821061203c57505050565b60208091833561204b81610a11565b815201910190612030565b356107a681610a11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e6400000000149015171561110d57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361110d57565b8181029291811591840414171561110d57565b81811061210f575050565b60008155600101612104565b9067ffffffffffffffff83116106325768010000000000000000831161063257815483835580841061217f575b5090600052602060002060005b8381106121625750505050565b600190602084359461217386610a11565b01938184015501612155565b61219790836000528460206000209182019101612104565b38612148565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b90600182811c92168015612237575b602083101461220857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916121fd565b9190601f811161225057505050565b6106c0926000526020600020906020601f840160051c8301931061227c575b601f0160051c0190612104565b909150819061226f565b90929167ffffffffffffffff8111610632576122ac816122a684546121ee565b84612241565b6000601f821160011461230a5781906122fb9394956000926122ff575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0135905038806122c9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461233d84600052602060002090565b91805b87811061239657508360019596971061235e575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612354565b90926020600181928686013581550194019101612340565b9160209082815201919060005b8181106123c85750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356123f181610a11565b1681520194019291016123bb565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916124989373ffffffffffffffffffffffffffffffffffffffff958661248a9367ffffffffffffffff6107a69e9c9d9b96168a5216602089015260c0604089015260c08801916123ae565b9185830360608701526123ae565b9416608082015260a08185039101526123ff565b906040519182815491828252602082019060005260206000209260005b8181106124de5750506106c09250038361066f565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016124c9565b9060405191826000825492612521846121ee565b808452936001811690811561258d5750600114612546575b506106c09250038361066f565b90506000929192526020600020906000915b8183106125715750509060206106c09282010138612539565b6020919350806001915483858901015201910190918492612558565b602093506106c09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612539565b906040516125da81610615565b60e06126cb60048395612642612638825473ffffffffffffffffffffffffffffffffffffffff80821616885261262f61261e8267ffffffffffffffff9060a01c1690565b67ffffffffffffffff1660208a0152565b60e01c60ff1690565b60ff166040870152565b6126a3612686600183015461266a61265d8263ffffffff1690565b63ffffffff1660608a0152565b60201c73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b6126af600282016124ac565b60a08601526126c0600382016124ac565b60c08601520161250d565b910152565b604051906126df60208361066f565b6000808352366020840137565b906126f682611fe8565b612703604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06127318294611fe8565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b9015611f635790565b9081602091031261019557516107a681610a11565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161286257501690565b61287b915060203d6020116119eb576119dd818361066f565b1690565b9190811015611f635760051b0190565b90816020910312610195575190565b9160206107a69381815201916123ff565b60ff166020039060ff821161110d57565b909291928311610195579190565b906004116101955790600490565b90939293848311610195578411610195578101920390565b359060208110612902575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b92919261293b826106e0565b91612949604051938461066f565b829481845281830111610195578281602093846000960137010152565b9160ff811690602082106129c2575b50810361298857906107a691369161292f565b906129be6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161289e565b0390fd5b60208311612a5857602083036129755790506129f66129f060ff6129e8849695966128af565b1685856128c0565b906128f4565b612a225791612a1b9181612a15612a0f6107a6966128af565b60ff1690565b916128dc565b369161292f565b506129be6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161289e565b6040517f3aeba390000000000000000000000000000000000000000000000000000000008152806129be85876004840161289e565b9081602091031261019557516107a68161111d565b67ffffffffffffffff1667ffffffffffffffff811461110d5760010190565b6040519060c0820182811067ffffffffffffffff82111761063257604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612b0d82611fe8565b612b1a604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612b488294611fe8565b019060005b828110612b5957505050565b602090612b64612ac1565b82828501015201612b4d565b60405190612b7d82610637565b60606040838281528260208201520152565b919082604091031261019557604051612ba781610653565b60208082948035612bb781610a11565b84520135910152565b805115611f635760200190565b8051821015611f635760209160051b010190565b90612beb82611fe8565b612bf8604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612c268294611fe8565b019060005b828110612c3757505050565b806060602080938501015201612c2b565b60409067ffffffffffffffff6107a694931681528160208201520190610752565b81601f82011215610195578051612c7f816106e0565b92612c8d604051948561066f565b81845260208284010111610195576107a6916020808501910161072f565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201612c69565b9080602083519182815201916020808360051b8301019401926000915b838310612cfd57505050505090565b9091929394602080612da2837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612d91612d7f612d6d612d5b8887015160c08a88015260c0870190610752565b60408701518682036040880152610752565b60608601518582036060870152610752565b60808501518482036080860152610752565b9201519060a0818403910152610752565b97019301930191939290612cee565b9193906107a69593612fad612fd29260a08652612ddb60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612f7a612f44612f0e612ed8612ea2612e6e8c61026060e08a0151916101c06101808201520190610752565b6101008801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001888f0152610752565b6101208701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101c08e0152610752565b6101408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101e08d0152610752565b6101608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016102008c0152610752565b6101808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016102208b0152612cd1565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610240880152610752565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610752565b9080602083519182815201916020808360051b8301019401926000915b83831061301057505050505090565b909192939460208061304c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610752565b97019301930191939290613001565b93929061307090606086526060860190610752565b938085036020820152825180865260208601906020808260051b8901019501916000905b8282106130b257505050506107a69394506040818403910152612fe4565b9091929560208061313e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845263ffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610752565b980192019201909291613094565b73ffffffffffffffffffffffffffffffffffffffff60015416330361316d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b906001820180921161110d57565b600101908160011161110d57565b906014820180921161110d57565b90600c820180921161110d57565b9190820180921161110d57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161110d57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820191821161110d57565b9190820391821161110d57565b8051916132518151846131cf565b9283156133a85760005b848110613269575050505050565b8181101561338d5761327e61181c8286612bcd565b73ffffffffffffffffffffffffffffffffffffffff81168015613363576132a483613197565b8781106132b65750505060010161325b565b848110156133335773ffffffffffffffffffffffffffffffffffffffff6132e061181c838a612bcd565b1682146132ef576001016132a4565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61335e61181c6133588885613236565b89612bcd565b6132e0565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6133a361181c61339d8484613236565b85612bcd565b61327e565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906133df82610615565b606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613445575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261019557825167ffffffffffffffff8111610195578261349f918501612c69565b9260208101516134ae81611f7c565b92604082015167ffffffffffffffff8111610195576107a69201612c69565b60409067ffffffffffffffff6107a6959316815281602082015201916123ff565b9060ff613508602092959495604085526040850190610752565b9416910152565b9291909261351b6133d2565b6004841015806138dd575b156137a35750509061353791615dc8565b9060c0820151805161372f575b506040820180515160005b818110613680575050805151156135e3575b505b6080820173ffffffffffffffffffffffffffffffffffffffff61359a825173ffffffffffffffffffffffffffffffffffffffff1690565b16156135a557505090565b6135c960806107a693015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b9060c08194939401926135f78451516126ec565b8352613604845151612be1565b946060810195865260005b85518051821015613672579061365361362d61181c83600195612bcd565b613638838951612bcd565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61366b81895161366161071a565b6119308383612bcd565b500161360f565b505093509350905038613561565b61368981613197565b828110613699575060010161354f565b6136a761181c838651612bcd565b73ffffffffffffffffffffffffffffffffffffffff6136cd61040861181c858951612bcd565b9116146136dc57600101613689565b6105806136ed61181c848751612bcd565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b600061377891613743604085015160ff1690565b9060405193849283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134ee565b0381305afa8015610be657156135445761379c903d806000833e611950818361066f565b5038613544565b60c08596929601946137b68651516126ec565b94604083019586526137c9875151612be1565b976060840198895260005b8851805182101561381257906137fd6137f261181c83600195612bcd565b613638838c51612bcd565b61380b818c5161366161071a565b50016137d4565b5050919550919550600092965061387f9361384861040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc19996000000000000000000000000000000000000000000000000000000008552600485016134cd565b03915afa8015610be6576000906000906000906138b1575b60a086015263ffffffff16845290505b60c0830152613563565b5050506138d36138a7913d806000833e6138cb818361066f565b810190613477565b9192508291613897565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061393361392d87876128ce565b90613411565b1614613526565b6020818303126101955780519067ffffffffffffffff821161019557019080601f8301121561019557815161396e81611fe8565b9261397c604051948561066f565b81845260208085019260051b82010192831161019557602001905b8282106139a45750505090565b6020809183516139b381610a11565b815201910190613997565b95949060009460a09467ffffffffffffffff613a129573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610752565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461110d5760010190565b949294939093613a736003613a6d8367ffffffffffffffff166000526004602052604060002090565b016124ac565b9473ffffffffffffffffffffffffffffffffffffffff613a948183166127b9565b16926040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180613af160048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610be657600091613c65575b5015613c5b5790613b4b600095949392604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016139be565b03915afa928315610be657600093613c36575b50825115613c3157613b7b613b7684518451906131cf565b6126ec565b6000918293835b8651811015613be057613b9861181c8289612bcd565b73ffffffffffffffffffffffffffffffffffffffff811615613bd45790613bce600192613638613bc789613a17565b9888612bcd565b01613b82565b50945060018095613bce565b509193909450613bf1575b50815290565b60005b8151811015613c295780613c23613c1061181c60019486612bcd565b613638613c1c87613a17565b9688612bcd565b01613bf4565b505038613beb565b915090565b613c549193503d806000833e613c4c818361066f565b81019061393a565b9138613b5e565b5050505050915090565b613c7e915060203d602011611cb957611cab818361066f565b38613b05565b93919293613ca0613c9882518651906131cf565b8651906131cf565b90613cb3613cad836126ec565b92612be1565b94600096875b8351891015613d195788613d0f613d02600193613cea613ce061181c8e9f9d9e9d8b612bcd565b613638838c612bcd565b613d08613cf7858c612bcd565b519180938491613a17565b9c612bcd565b528b612bcd565b5001979695613cb9565b959250929350955060005b8651811015613db357613d3a61181c8289612bcd565b600073ffffffffffffffffffffffffffffffffffffffff8216815b888110613d87575b5050906001929115613d71575b5001613d24565b613d8190613638613bc789613a17565b38613d6a565b81613d9861040861181c848c612bcd565b14613da557600101613d55565b506001915081905038613d5d565b509390945060005b8551811015613e5157613dd161181c8288612bcd565b600073ffffffffffffffffffffffffffffffffffffffff8216815b878110613e25575b5050906001929115613e08575b5001613dbb565b613e1f90613638613e1888613a17565b9787612bcd565b38613e01565b81613e3661040861181c848b612bcd565b14613e4357600101613dec565b506001915081905038613df4565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176106325760405260606080836000815260006020820152600060408201526000838201520152565b90613ea682611fe8565b613eb3604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613ee18294611fe8565b019060005b828110613ef257505050565b602090613efd613e5d565b82828501015201613ee6565b519061ffff8216820361019557565b9081606091031261019557613f2c81613f09565b9160406020830151613f3d81611f7c565b9201516107a681611f7c565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613fb35750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613fdc81610a11565b16815260208781013590820152019401929101613fa6565b949391929067ffffffffffffffff1685526080602086015261406b61402e61401c8580613f49565b60a060808a01526101208901916123ff565b61403b6020860186613f49565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a01526123ff565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b36038313610195576106c09561413c6141068360609761417b978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808261416d9a0301910152613f99565b91614132614115888301611112565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613f49565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c01526123ff565b908782036040890152610752565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff821161110d57565b908160a09103126101955780519160208201516141bc81611f7c565b9160408101516141cb81611f7c565b9160806141da60608401613f09565b9201516107a68161111d565b9260c09473ffffffffffffffffffffffffffffffffffffffff9167ffffffffffffffff61ffff95846107a69b9a9616885216602087015260408601521660608401521660808201528160a08201520190610752565b90816060910312610195578051613f2c81611f7c565b9190826040910312610195576020825161426a81611f7c565b92015190565b9081602091031261019557517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681036101955790565b81156142ae570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9291909260009260009360408401928351519461431a61431561431060408b01986143088a8d61273b565b9190506131cf565b613197565b613e9c565b956000946000945b875180518710156145a55761040861040861181c8961434094612bcd565b61438c602060608701926143558a8551612bcd565b519060405180809581947f958021a70000000000000000000000000000000000000000000000000000000083528d60048401612c48565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff91600091614587575b50168015614526579060608d93926143cd8a8451612bcd565b51906143de60208a015161ffff1690565b958a614419604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613ff4565b03915afa8015610be6576001936144c3938a8e8e956000809581976144cc575b509083929161ffff6144618561445a61181c6144b7996144bd9d9e51612bcd565b9451612bcd565b519161448a61446e6106c2565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b63ffffffff8916602085015263ffffffff8b16604085015216606083015260808201526119308383612bcd565b50614186565b98614186565b95019495614322565b6144bd975061181c965084939291509361ffff6144618261445a6145096144b79960603d811161451f575b614501818361066f565b810190613f18565b9c9196909c9d5050505050505090919293614439565b503d6144f7565b6105808761453861181c8b8e51612bcd565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045267ffffffffffffffff16602452604490565b61459f915060203d81116119eb576119dd818361066f565b386143b4565b50919293959994509796956145ba838661273b565b90506148f2575b506146426146336040926145fc866145d98b516131dc565b966145f36145ea60208c018c61219d565b9290508b61273b565b9190508b61607a565b614606868b612bcd565b52614611858a612bcd565b5061462d6020614621878c612bcd565b51015163ffffffff1690565b90614186565b9961462d83614621868b612bcd565b9861466861040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b82517f4aa67c0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff97909716600488015263ffffffff91821660248801529916604486015293979388606481875afa918215610be657600098600093614875575b5060209361478a95936147389373eba517d20000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff6147296080606097015173ffffffffffffffffffffffffffffffffffffffff1690565b1603614854575b505001612056565b9060405180809581947f4ab35b0b0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03915afa8015610be6577bffffffffffffffffffffffffffffffffffffffffffffffffffffffff91600091614825575b5016936000925b825184101561481c576148146001916147f0886147eb60606147e38a8a612bcd565b51015161208f565b6142a4565b60606147fc8888612bcd565b510152606061480b8787612bcd565b510151906131cf565b9301926147c1565b91945092909150565b614847915060203d60201161484d575b61483f818361066f565b810190614270565b386147ba565b503d614835565b8361486261486c928b612bcd565b51019182516131cf565b90523880614730565b6080995061478a959350936060916147389373eba517d20000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff6147296148d760209a60403d6040116148eb575b6148cf818361066f565b810190614251565b9f909f999b505050505093509150936146cf565b503d6148c5565b614908610a7a6104086109a96109a3878a61273b565b9060009051519173ffffffffffffffffffffffffffffffffffffffff6149346109a96109a3888b61273b565b9160e088019283516149636149476106c2565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b8560208301528560408301528560608301526080820152614984868d612bcd565b5261498f858c612bcd565b5016906040517f01ffc9a7000000000000000000000000000000000000000000000000000000008152602081806149ed60048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381865afa928315610be6578a888a958c948891614c72575b50614b6f575b5050505050506145c15790614ab4906060614a4261040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b614a526109a96109a3888b61273b565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015293849190829081906044820190565b03915afa908115610be65760409363ffffffff61463393614642958c600092600092600091614b32575b5090614b1c614b2793928b614b0f6060614af88b87612bcd565b5101996020614b078288612bcd565b510195612bcd565b51019063ffffffff169052565b9063ffffffff169052565b1690529250506145c1565b614b1c9450614b279350614b5e915060603d606011614b68575b614b56818361066f565b81019061423b565b9194909350614ade565b503d614b4c565b614bea60a095614bb16020614ba7606082614b9e6109a3614b976109a96109a38f8d9061273b565b998d61273b565b01359901612056565b99015161ffff1690565b905190604051988997889687967f2c063404000000000000000000000000000000000000000000000000000000008852600488016141e6565b03915afa918215610be657809181908294614c36575b50614c2a908b614b1c6060614c158984612bcd565b5101946040614b0f8a6020614b078288612bcd565b52843887818a88614a0c565b915050614c2a9250614c60915060a03d60a011614c6b575b614c58818361066f565b8101906141a0565b949192919050614c00565b503d614c4e565b614c8b915060203d602011611cb957611cab818361066f565b38614a06565b9073ffffffffffffffffffffffffffffffffffffffff614d639392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452614cf360648561066f565b16600080604093845195614d07868861066f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614d88573d614d54614d4b826106e0565b9451948561066f565b83523d6000602085013e61678f565b805180614d6e575050565b81602080614d83936106c09501019101612a8d565b6162c8565b6060925061678f565b9160606106c0929493614dde8160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b90614e1d826106e0565b614e2a604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061273182946106e0565b91825160148102908082046014149015171561110d57614e7a614e7f916131a5565b6131b3565b90614e91614e8c836131c1565b614e13565b906014614e9d83612bc0565b5360009260215b8651851015614ecf576014600191614ebf61181c888b612bcd565b60601b8187015201940193614ea4565b919550936020935060601b90820152828152012090565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929160009160408101614f1b818361273b565b9050615062575b506104086060614f329201612056565b9160005b825181101561505a576060614f4b8285612bcd565b51015190811561505157868114614f9657614f90600192614f8a614f6f8488612bcd565b515173ffffffffffffffffffffffffffffffffffffffff1690565b87614c91565b01614f36565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd360000000000000000000000000000000000000000000000000000000060048201529173ffffffffffffffffffffffffffffffffffffffff841690602084602481855afa8015610be657600194600091615033575b50615022575b5050614f90565b61502c9187614c91565b388061501b565b61504b915060203d8111611cb957611cab818361066f565b38615015565b60019150614f90565b505050509050565b94509150614f32610408606061508e610a7a6104086109a96109a36150878951613209565b9b8a61273b565b94925050614f22565b604051906150a482610653565b60606020838281520152565b919060408382031261019557604051906150c982610653565b8193805167ffffffffffffffff811161019557826150e8918301612c69565b835260208101519167ffffffffffffffff8311610195576020926126cb9201612c69565b9060208282031261019557815167ffffffffffffffff8111610195576107a692016150b0565b90608073ffffffffffffffffffffffffffffffffffffffff8161515e855160a0865260a0860190610752565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206107a6928181520190615132565b919060408382031261019557825167ffffffffffffffff81116101955760209161426a9185016150b0565b61ffff6151e76107a69593606084526060840190615132565b931660208201526040818403910152610752565b90919293615207612ac1565b50602082019081511561579b5761523b610408610a7a610408865173ffffffffffffffffffffffffffffffffffffffff1690565b9573ffffffffffffffffffffffffffffffffffffffff87169283158015615710575b6156ad576152e881519161526f615097565b5051865173ffffffffffffffffffffffffffffffffffffffff16906152c66152956106c2565b8b815267ffffffffffffffff8b1660208201529573ffffffffffffffffffffffffffffffffffffffff166040870152565b606085015273ffffffffffffffffffffffffffffffffffffffff166080840152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd36000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610be65760009161568e575b501561559857509161539496979160008094604051998a95869485937fb1c71c65000000000000000000000000000000000000000000000000000000008552600485016151ce565b03925af18015610be65760009460009161554f575b509460006154ee9361548c6154286154609561181c6153fc9a965b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b61066f565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561066f565b6154b96154af84519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134ee565b0381305afa928315610be65760009361552f575b5060200151936155106106d1565b958652602086015260408501526060840152608083015260a082015290565b6020919350615548903d806000833e611950818361066f565b9290615502565b6153fc955061181c9691506154ee9361548c615428615460956155856000953d8088833e61557d818361066f565b8101906151a3565b9b909b969b9a50509550505093506153a9565b979161ffff919597935016615664575161563a5760006155e593604051809581927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301615192565b038183855af1918215610be6576153fc9561548c61542860009361181c615460976154ee998791615618575b50966153c4565b61563491503d8089833e61562c818361066f565b81019061510c565b38615611565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b6156a7915060203d602011611cb957611cab818361066f565b3861534c565b6105806156ce865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818c5afa908115610be65760009161577c575b501561525d565b615795915060203d602011611cb957611cab818361066f565b38615775565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592615897947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906158b36020928281519485920161072f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b1681526158ff825180936020898501910161072f565b019160f81b168382015261591d82518093602060028501910161072f565b01019160f81b168382015261593c82518093602060028501910161072f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006107a69f9e9c9860f81b1681526159b1825180936020898501910161072f565b019160f01b16838201526159cf82518093602060038501910161072f565b01019160f01b16838201526159ed825180936020898501910161072f565b01019160f01b16600282015201906158a0565b60e081019060ff82515111615d995761010081019060ff82515111615d6a5761012081019260ff84515111615d3b5761014082019060ff82515111615d0c5761016083019461ffff86515111615cdd57610180840194600186515111615cae576101a085019261ffff84515111615c7d57855167ffffffffffffffff16602087015167ffffffffffffffff16906040880151615aa39067ffffffffffffffff1690565b976060810151615ab69063ffffffff1690565b906080810151615ac99063ffffffff1690565b60a082015161ffff169160c00151926040519b8c966020880196615aec976157c5565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752615b1c908761066f565b51908151615b2a9060ff1690565b9051805160ff169351908151615b409060ff1690565b906040519586956020870195615b55966158b7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252615b85908261066f565b606094519081511515615c2f97615bf36107a698615c3597615c2f96615c1f95615c61575b505196615bb8885160ff1690565b935191615bc7835161ffff1690565b91615bd4825161ffff1690565b905193615be3855161ffff1690565b936040519b8c9860208a01615942565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810185528461066f565b60405196879560208701906158a0565b906158a0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261066f565b615c76919250615c7090612bc0565b5161648e565b9038615baa565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105806024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b90615dd16133d2565b9160118210615ff45780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603615f815750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a615e50816126ec565b60408601908152615e6082612be1565b906060870191825260005b838110615f355750505050615ee08383615ed6615eca615ec0615eb9615e99615eea9887615ef49c9b616624565b73ffffffffffffffffffffffffffffffffffffffff90911660808d015290565b85856166fa565b929190369161292f565b60a08a01528383616762565b949190369161292f565b60c08801526166fa565b939190369161292f565b60e08401528103615f03575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b80600191615f7a615f64615f5d615f50615f749a8d8d616624565b9190613638868a51612bcd565b8b8b6166fa565b9391889a919a51949a369161292f565b92612bcd565b5201615e6b565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b90816020910312610195576107a690613f09565b9261ffff6107a6959367ffffffffffffffff61606c94168652166020850152608060408501526080840190610c8a565b916060818403910152610752565b9091616084613e5d565b506160a38267ffffffffffffffff166000526004602052604060002090565b936160b3855460ff9060e01c1690565b9061616d616152616149608084019461613561610f6161076160fa60016160ee8b5173ffffffffffffffffffffffffffffffffffffffff1690565b9e015463ffffffff1690565b885163ffffffff1661462d565b9a60756131cf565b9761614361613b60ff61612960a08b019c8d5151906131cf565b951694616135866120c1565b906131cf565b93604f6131cf565b906120f1565b63ffffffff1690565b925173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811673eba517d2000000000000000000000000000000000361620157505061ffff92506161f3906161e66000935b51956161d96161bd6106c2565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b9061622461040860209373ffffffffffffffffffffffffffffffffffffffff1690565b60406162348484015161ffff1690565b92015191855196616274604051988995869485947fe962e69e0000000000000000000000000000000000000000000000000000000086526004860161603c565b03915afa908115610be6576161e66161f39261ffff95600091616299575b50936161b0565b6162bb915060203d6020116162c1575b6162b3818361066f565b810190616028565b38616292565b503d6162a9565b156162cf57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9660026164629761642f60226107a69f9e9c9799600199859f9b7fff000000000000000000000000000000000000000000000000000000000000009061642f9f8261642f9c6164369c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b1660218201526163e0825180936020898501910161072f565b019160f81b16838201526163fe82518093602060238501910161072f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b01906158a0565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116165f557604081019160ff835151116165c657606082019160ff8351511161659757608081019260ff845151116165685760a0820161ffff81515111616539576107a694615c3593519451916164f0835160ff1690565b9751916164fe835160ff1690565b94519061650c825160ff1690565b90519361651a855160ff1690565b935196616529885161ffff1690565b966040519c8d9b60208d01616353565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b9291909260018201918483116166c85781013560001a8281156166bd57506014810361669057820193841161665c57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b919060028201918183116166c8578381013560f01c016002019281841161672e57918391616727936128dc565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b919060018201918183116166c8578381013560001a016001019281841161672e57918391616727936128dc565b9192901561680a57508151156167a3575090565b3b156167ac5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561681d5750805190602001fd5b6129be906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161079556fea164736f6c634300081a000a",
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

func (_OnRamp *OnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
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
	SequenceNumber    uint64
	MessageId         [32]byte
	EncodedMessage    []byte
	Receipts          []OnRampReceipt
	VerifierBlobs     [][]byte
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
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
	SequenceNumber    uint64
	Router            common.Address
	DefaultCCVs       []common.Address
	LaneMandatedCCVs  []common.Address
	DefaultExecutor   common.Address
	OffRamp           []byte
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
	return common.HexToHash("0x276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c96")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae")
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

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

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

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

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
