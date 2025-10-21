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
	Router           common.Address
	SequenceNumber   uint64
	DefaultExecutor  common.Address
	LaneMandatedCCVs []common.Address
	DefaultCCVs      []common.Address
	OffRamp          []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	DefaultCCVs       []common.Address
	LaneMandatedCCVs  []common.Address
	DefaultExecutor   common.Address
	OffRamp           []byte
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type OnRampReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102f55760405161560638819003601f8101601f191683016001600160401b038111848210176102fa578392829160405283398101039060c082126102f557606082126102f557610054610310565b81519092906001600160401b03811681036102f55783526020820151906001600160a01b03821682036102f5576020840191825260606100966040850161032f565b6040860190815291605f1901126102f5576100af610310565b916100bc6060850161032f565b835260808401519384151585036102f55760a06100e091602086019687520161032f565b946040840195865233156102e457600180546001600160a01b0319163317905580516001600160401b03161580156102d2575b80156102c0575b61029357516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102ae575b80156102a4575b6102935780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610310565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610310565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a16040516152c290816103448239608051818181610b79015281816114670152611b01015260a05181611b3a015260c051818181611b760152611fea0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b038111838210176102fa57604052565b51906001600160a01b03821682036102f55756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635cb80c5d146100f357806366c3a5c7146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611a5e565b61196a565b6111dd565b611009565b610f67565b610f15565b610e2c565b610d60565b610cc8565b6108cc565b61077e565b610693565b610409565b610368565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576060610140611ae1565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176101d557604052565b61018a565b6060810190811067ffffffffffffffff8211176101d557604052565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b6040519061027f6101608361022e565b565b6040519061027f60a08361022e565b6040519061027f60408361022e565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102e860208361022e565b60008252565b60005b8381106103015750506000910152565b81810151838201526020016102f1565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361034d815180928187528780880191016102ee565b0116010190565b906020610365928181520190610311565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103e560408051906103a9818361022e565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610311565b0390f35b67ffffffffffffffff81160361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435610444816103e9565b60243567ffffffffffffffff8111610185576104649036906004016103fb565b6000906104858367ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6104d56104bc845473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b16156106395761053d92936105366105026104f36080850185611b9e565b906104fd87611d63565b612e6c565b9361050b611e0e565b906040850161051a8187611e79565b90506105e4575b50610530600287519201611bef565b906134ee565b8352613bb2565b6000916000916000915b81518310156105d9576105cf6105a761057160019360606105688888611f21565b51015190611f72565b966105a161059460206105848989611f21565b51015167ffffffffffffffff1690565b67ffffffffffffffff1690565b90611f72565b946105a16105c660406105ba8888611f21565b51015163ffffffff1690565b63ffffffff1690565b9201919293610547565b604051908152602090f35b610632919250602061061461060261060d610608610602868c611e79565b90611efc565b611f0a565b9389611e79565b0135610625602089015161ffff1690565b9060808901519287613283565b9038610521565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576106cd6004356103e9565b60206106e36024356106de81610675565b611f8b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101855760043567ffffffffffffffff81116101855760040160009280601f8301121561077a5781359367ffffffffffffffff851161077757506020808301928560051b010111610185579190565b80fd5b8380fd5b346101855761078c36610701565b906107ac60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b8181106107b957005b6107ca6104bc610608838587612084565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156108c7576001948891600093610897575b508261083f575b50505050016107b0565b61084a918391613ea3565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610835565b6108b991935060203d81116108c0575b6108b1818361022e565b810190612094565b913861082e565b503d6108a7565b611f7f565b34610185576108da36610701565b906108e3613fa3565b6000915b8083106108f057005b6108fb8382846120a3565b92610905846120e3565b9367ffffffffffffffff85169081158015610b6d575b610b3557610966939495610980604083019161093783856120ed565b9790610960606087019961095861094e8c8a6120ed565b9490923691612141565b923691612141565b90613fee565b67ffffffffffffffff166000526004602052604060002090565b60208301906109d261099183611f0a565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6109e96109df84866120ed565b90600384016121ae565b610a006109f688866120ed565b90600284016121ae565b60808401610a106104bc82611f0a565b15610b0b576001977f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae95610aeb83610af1610b0096610a97610a54610ae398611f0a565b8f83019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610ada610ad3610acd60a0880193610abc610ab2868b611b9e565b9060048401612275565b5460a01c67ffffffffffffffff1690565b9a611f0a565b9a866120ed565b979096866120ed565b949093611f0a565b94611b9e565b959094604051998a998a61242d565b0390a20191906108e7565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016821461091b565b906020808351928381520192019060005b818110610bbe5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610bb1565b90610365916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015273ffffffffffffffffffffffffffffffffffffffff604083015116606082015260a0610c95610c62606085015160c0608086015260e0850190610ba0565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610ba0565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610311565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610d0c816103e9565b606060a0604051610d1c816101b9565b600081526000602082015260006040820152828082015282608082015201521660005260046020526103e5610d546040600020611d63565b60405191829182610bea565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610d97611ac2565b50604051610da4816101da565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103e582606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff81163303610eeb577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610fab816103e9565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610fef5760405167ffffffffffffffff9091168152602090f35b611f35565b359061027f82610675565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051611046816101da565b60043561105281610675565b815260243561106081610fff565b602082019081526044359061107482610675565b60408301918252611083613fa3565b73ffffffffffffffffffffffffffffffffffffffff835116159182156111bd575b5081156111b2575b5061118a5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611175611ae1565b6111846040519283928361417d565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386110ac565b5173ffffffffffffffffffffffffffffffffffffffff16159150386110a4565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435611218816103e9565b60243567ffffffffffffffff8111610185576112389036906004016103fb565b604435916064359061124982610675565b60025460a01c60ff166119405761129a740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6112b88167ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff8316156119165781546112f76104bc73ffffffffffffffffffffffffffffffffffffffff831681565b33036118ec57849361130c6080860186611b9e565b61131586611d63565b9161131f92612e6c565b9160a01c67ffffffffffffffff166113369061249b565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff0000000000000000000000000000000000000000161785556040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260148152909390916113c160348461022e565b866020850180516113d39061ffff1690565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b16602082015260148152989061141360348b61022e565b61141d8480611b9e565b80611427926124ba565b9a90604086019b6114398d8098611e79565b6114439150612569565b9d6020810161145191611b9e565b93909461145c61026f565b67ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682528c1660208201529b67ffffffffffffffff909e1660408d019081529d60608d01526114b760048801611ca3565b60808d015261ffff1660a08c015260c08b015236906114d5926125d6565b60e08901526114e26102d9565b61010089015261012088019b8c5236906114fb926125d6565b6101408701528a61150a611e0e565b926115158583611e79565b6115319691506118a3575b505050610530600288519201611bef565b845261153b61260d565b96611547858a85613bb2565b9660208901978852611559818b611e79565b905061171e575b50505061156f82979497614971565b8087526020815191012092519461158686516126d6565b9660408101978852606060009601955b875181101561167c578a906000898989838a61162660206115ec896115e66115e06104bc806115c5858e611f21565b515173ffffffffffffffffffffffffffffffffffffffff1690565b98611f0a565b98611f21565b510151604051998a97889687957f71c5c2ba0000000000000000000000000000000000000000000000000000000087523060048801612808565b03925af180156108c757816116549160019460009161165b575b508c519061164e8383611f21565b52611f21565b5001611596565b611676913d8091833e61166e818361022e565b81019061266e565b38611640565b506103e5955084935067ffffffffffffffff7f9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f916116de8a6116c68d5167ffffffffffffffff1690565b92519551905190846040519586951698169684612a9f565b0390a461170e7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b600161172a828c611e79565b9050036118795761179692600061175c6104bc6104bc60025473ffffffffffffffffffffffffffffffffffffffff1690565b60608901519060405180809881947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610354565b03915afa80156108c757856117dd936117ef9660009361185c575b508d8351156117f8575b6117cb6117d8926117d392611e79565b369291611efc565b6126a5565b6142fe565b9051906117e982611f14565b52611f14565b50388080611560565b92506117d36117cb6117d8926118516118118780611b9e565b919061182560405193849260208401612694565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261022e565b9592509250506117bb565b6118729193503d806000833e61166e818361022e565b91386117b1565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b6118e393945060206118cb610602836118c56106086106026118d6978a611e79565b96611e79565b0135915161ffff1690565b9060808a01519288613283565b908a388a611520565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff6004356119ba81610675565b6119c2613fa3565b16338114611a3457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557611a986004356103e9565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611acf826101da565b60006040838281528260208201520152565b611ae9611ac2565b50604051611af6816101da565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b906040519182815491828252602082019060005260206000209260005b818110611c2157505061027f9250038361022e565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201611c0c565b90600182811c92168015611c99575b6020831014611c6a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611c5f565b9060405191826000825492611cb784611c50565b8084529360018116908115611d235750600114611cdc575b5061027f9250038361022e565b90506000929192526020600020906000915b818310611d0757505090602061027f9282010138611ccf565b6020919350806001915483858901015201910190918492611cee565b6020935061027f9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611ccf565b90604051611d70816101b9565b60a0611df16004839567ffffffffffffffff80825473ffffffffffffffffffffffffffffffffffffffff81168852861c1616602086015273ffffffffffffffffffffffffffffffffffffffff80600183015416166040860152611dd560028201611bef565b6060860152611de660038201611bef565b608086015201611ca3565b910152565b67ffffffffffffffff81116101d55760051b60200190565b60405190611e1d60208361022e565b6000808352366020840137565b90611e3482611df6565b611e41604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611e6f8294611df6565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015611f055790565b611ecd565b3561036581610675565b805115611f055760200190565b8051821015611f055760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9060018201809211610fef57565b91908201809211610fef57565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156108c757600090612035575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d60201161207c575b8161204f6020938361022e565b810103126101855773ffffffffffffffffffffffffffffffffffffffff905161207781610675565b61201a565b3d9150612042565b9190811015611f055760051b0190565b90816020910312610185575190565b9190811015611f055760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610185570190565b35610365816103e9565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160051b3603831361018557565b92919061214d81611df6565b9361215b604051958661022e565b602085838152019160051b810192831161018557905b82821061217d57505050565b60208091833561218c81610675565b815201910190612171565b8181106121a2575050565b60008155600101612197565b9067ffffffffffffffff83116101d5576801000000000000000083116101d5578154838355808410612212575b5090600052602060002060005b8381106121f55750505050565b600190602084359461220686610675565b019381840155016121e8565b61222a90836000528460206000209182019101612197565b386121db565b9190601f811161223f57505050565b61027f926000526020600020906020601f840160051c8301931061226b575b601f0160051c0190612197565b909150819061225e565b90929167ffffffffffffffff81116101d55761229b816122958454611c50565b84612230565b6000601f82116001146122f95781906122ea9394956000926122ee575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0135905038806122b8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461232c84600052602060002090565b91805b87811061238557508360019596971061234d575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612343565b9092602060018192868601358155019401910161232f565b9160209082815201919060005b8181106123b75750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356123e081610675565b1681520194019291016123aa565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916124879373ffffffffffffffffffffffffffffffffffffffff95866124799367ffffffffffffffff6103659e9c9d9b96168a5216602089015260c0604089015260c088019161239d565b91858303606087015261239d565b9416608082015260a08185039101526123ee565b67ffffffffffffffff1667ffffffffffffffff8114610fef5760010190565b9092919283600c1161018557831161018557600c01917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40190565b906004116101855790600490565b909291928360041161018557831161018557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b6040519061254b826101f6565b60606080836000815282602082015282604082015282808201520152565b9061257382611df6565b612580604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125ae8294611df6565b019060005b8281106125bf57505050565b6020906125ca61253e565b828285010152016125b3565b9291926125e28261029f565b916125f0604051938461022e565b829481845281830111610185578281602093846000960137010152565b6040519061261a826101da565b60606040838281528260208201520152565b81601f820112156101855780516126428161029f565b92612650604051948561022e565b818452602082840101116101855761036591602080850191016102ee565b9060208282031261018557815167ffffffffffffffff811161018557610365920161262c565b9160206103659381815201916123ee565b9190826040910312610185576040516126bd81610212565b602080829480356126cd81610675565b84520135910152565b906126e082611df6565b6126ed604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061271b8294611df6565b019060005b82811061272c57505050565b806060602080938501015201612720565b9080602083519182815201916020808360051b8301019401926000915b83831061276957505050505090565b90919293946020806127f9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260806127e86127d66127c48786015160a08987015260a0860190610311565b60408601518582036040870152610311565b60608501518482036060860152610311565b920151906080818403910152610311565b9701930193019193929061275a565b929491936129f1610365979573ffffffffffffffffffffffffffffffffffffffff612a169416865260c0602087015261284e60c08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660e0870152604081015167ffffffffffffffff166101008701526101406129be61298861295261291d6128d96128a68c61022060608a0151916101606101208201520190610311565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4001888f0152610311565b60a087015161ffff166101608d015260c08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40016101808e0152610311565b60e08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408c8303016101a08d0152610311565b6101008501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408b8303016101c08c0152610311565b6101208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a8303016101e08b015261273d565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4087830301610200880152610311565b956040850152606084019073ffffffffffffffffffffffffffffffffffffffff169052565b608082015260a0818403910152610311565b9080602083519182815201916020808360051b8301019401926000915b838310612a5457505050505090565b9091929394602080612a90837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610311565b97019301930191939290612a45565b939290612ab490606086526060860190610311565b938085036020820152825180865260208601906020808260051b8901019501916000905b828210612af657505050506103659394506040818403910152612a28565b90919295602080612b86837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610311565b980192019201909291612ad8565b60405190612ba1826101f6565b6060608083828152600060208201526000604082015282808201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612bf4575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90612c3082611df6565b612c3d604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612c6b8294611df6565b019060005b828110612c7c57505050565b602090604051612c8b81610212565b6000815260608382015282828501015201612c70565b9080601f8301121561018557816020610365933591016125d6565b9080601f8301121561018557813591612cd483611df6565b92612ce2604051948561022e565b80845260208085019160051b830101918383116101855760208101915b838310612d0e57505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101855760405190612d5b82610212565b6020830135612d6981610675565b825260408301359167ffffffffffffffff831161018557612d9288602080969581960101612ca1565b83820152815201920191612cff565b359061ffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160a08183031261018557612ddf610281565b91813567ffffffffffffffff81116101855781612dfd918401612cbc565b8352612e0b60208301612da1565b6020840152612e1c60408301610ff4565b6040840152606082013567ffffffffffffffff81116101855781612e41918401612ca1565b6060840152608082013567ffffffffffffffff811161018557612e649201612ca1565b608082015290565b9190612e76612b94565b600483101580613107575b15613084575081612e9d92612e9592612503565b810190612db0565b9081515160005b818110612fd557505081515115612f36575b6040820173ffffffffffffffffffffffffffffffffffffffff612eed825173ffffffffffffffffffffffffffffffffffffffff1690565b1615612ef857505090565b612f1c604061036593015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b6080810192612f46845151612c26565b835260005b84518051821015612fcc5790612fc581612f84612f6a82600196611f21565b5173ffffffffffffffffffffffffffffffffffffffff1690565b612fab612f8f610290565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b612fb36102d9565b602082015287519061164e8383611f21565b5001612f4b565b50509250612eb6565b612fde81611f64565b828110612fee5750600101612ea4565b612ffc6115c5838751611f21565b73ffffffffffffffffffffffffffffffffffffffff6130226104bc6115c5858a51611f21565b91161461303157600101612fde565b6106716130426115c5848851611f21565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b9193926130923686846125d6565b606084015260808101936130a7855151612c26565b845260005b855180518210156130fa57906130f3816130cb612f6a82600196611f21565b6130d6612f8f610290565b6130e1368c8a6125d6565b602082015288519061164e8383611f21565b50016130ac565b5050945091925050612eb6565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061315d61315786866124f5565b90612bc0565b1614612e81565b90816020910312610185575161036581610fff565b6020818303126101855780519067ffffffffffffffff821161018557019080601f830112156101855781516131ad81611df6565b926131bb604051948561022e565b81845260208085019260051b82010192831161018557602001905b8282106131e35750505090565b6020809183516131f281610675565b8152019101906131d6565b95949060009460a09467ffffffffffffffff6132519573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610311565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610fef5760010190565b9492949390936132b260036132ac8367ffffffffffffffff166000526004602052604060002090565b01611bef565b9473ffffffffffffffffffffffffffffffffffffffff6132d3818316611f8b565b16926040517f01ffc9a70000000000000000000000000000000000000000000000000000000081526020818061333060048201907f479eecb200000000000000000000000000000000000000000000000000000000602083019252565b0381885afa9081156108c7576000916134bf575b50156134b5579061338a600095949392604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016131fd565b03915afa9283156108c757600093613490575b5082511561348b576133ba6133b58451845190611f72565b611e2a565b6000918293835b865181101561343a576133d7612f6a8289611f21565b73ffffffffffffffffffffffffffffffffffffffff81161561342e579061342860019261340d61340689613256565b9888611f21565b9073ffffffffffffffffffffffffffffffffffffffff169052565b016133c1565b50945060018095613428565b50919390945061344b575b50815290565b60005b8151811015613483578061347d61346a612f6a60019486611f21565b61340d61347687613256565b9688611f21565b0161344e565b505038613445565b915090565b6134ae9193503d806000833e6134a6818361022e565b810190613179565b913861339d565b5050505050915090565b6134e1915060203d6020116134e7575b6134d9818361022e565b810190613164565b38613344565b503d6134cf565b61350e6135096135018351855190611f72565b855190611f72565b612c26565b91600094855b8351811015613551578061354a61352d60019387611f21565b519861353881613256565b99613543828a611f21565b5287611f21565b5001613514565b50939094915060005b855181101561362c57613570612f6a8288611f21565b600073ffffffffffffffffffffffffffffffffffffffff8216815b868110613602575b5050156135a4575b5060010161355a565b92906135fb6001926135d36135b7610290565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b6135db6102d9565b60208701526135e981613256565b956135f48289611f21565b5286611f21565b509061359b565b816136136104bc6115c5848c611f21565b146136205760010161358b565b50505060013880613593565b5093506000905b845182101561370757613649612f6a8387611f21565b91600073ffffffffffffffffffffffffffffffffffffffff8416815b8481106136dd575b50501561367f575b6001019150613633565b6136d46001926136ac613690610290565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b6136b46102d9565b60208601526136c281613256565b946136cd8288611f21565b5285611f21565b50829150613675565b816136ee6104bc6115c5848b611f21565b146136fb57600101613665565b5050506001388061366d565b8252509150565b9061371882611df6565b613725604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06137538294611df6565b019060005b82811061376457505050565b602090604051613773816101f6565b6000815260008382015260006040820152600060608201526060608082015282828501015201613758565b519063ffffffff8216820361018557565b908160609103126101855780519161036560406137ce6020850161379e565b930161379e565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b9160209082815201919060005b81811061383f5750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff873561386881610675565b16815260208781013590820152019401929101613832565b959492939173ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff9216875216602086015260a060408601526139126138d56138c485806137d5565b60a0808a01526101408901916123ee565b6138e260208601866137d5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608984030160c08a01526123ee565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610185578401916020833593019167ffffffffffffffff8411610185578360061b360383136101855761027f956139e46139ad83608097613a23978d60e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6082613a159a0301910152613825565b916139db6139bd60608301610ff4565b73ffffffffffffffffffffffffffffffffffffffff166101008d0152565b868101906137d5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8403016101208c01526123ee565b908782036060890152610311565b94019061ffff169052565b81601f82011215610185578035613a4481611df6565b92613a52604051948561022e565b81845260208085019260061b8401019281841161018557602001915b838310613a7c575050505090565b6020604091613a8b84866126a5565b815201920191613a6e565b91909160a08184031261018557613aab610281565b92813567ffffffffffffffff81116101855781613ac9918401612ca1565b8452602082013567ffffffffffffffff81116101855781613aeb918401612ca1565b6020850152604082013567ffffffffffffffff81116101855781613b10918401613a2e565b6040850152613b2160608301610ff4565b6060850152608082013567ffffffffffffffff811161018557613b449201612ca1565b6080830152565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610fef57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211610fef57565b91908203918211610fef57565b929190815151613be2613bdd613bd86040850193613bd08587611e79565b919050611f72565b611f64565b61370e565b9060005b84518051821015613d725781613bfb91611f21565b5190613c216104bc6104bc845173ffffffffffffffffffffffffffffffffffffffff1690565b602083019060608251613c3960208b015161ffff1690565b92898d613c76604051968795869485947f2b0560010000000000000000000000000000000000000000000000000000000086523060048701613880565b03915afa9384156108c75760019460009182938392613d2a575b5090613cf763ffffffff613cbb613d04945173ffffffffffffffffffffffffffffffffffffffff1690565b965195613ce5613cc9610281565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b1667ffffffffffffffff166020870152565b63ffffffff166040850152565b60608301526080820152613d188286611f21565b52613d238185611f21565b5001613be6565b63ffffffff9450613cbb9350613d049250613d5e613cf79160603d8111613d6b575b613d56818361022e565b8101906137af565b9096509094509250613c90565b503d613d4c565b5050909394613e0d90939293613db3613da2604086015173ffffffffffffffffffffffffffffffffffffffff1690565b91613dad3688613a96565b86614d2e565b606085015190613de0613dc4610281565b73ffffffffffffffffffffffffffffffffffffffff9094168452565b600060208401526000604084015260608301526080820152613e028651613b4b565b906136cd8288611f21565b50613e188184611e79565b9050613e25575b50505090565b613e3a610608610602608093613e9a96611e79565b910151613e64613e48610281565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b6000602083015260006040830152600060608301526080820152613e888351613b78565b90613e938285611f21565b5282611f21565b50388080613e1f565b9073ffffffffffffffffffffffffffffffffffffffff613f759392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452613f0560648561022e565b16600080604093845195613f19868861022e565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15613f9a573d613f66613f5d8261029f565b9451948561022e565b83523d6000602085013e6151ed565b805180613f80575050565b81602080613f959361027f9501019101613164565b614f64565b606092506151ed565b73ffffffffffffffffffffffffffffffffffffffff600154163303613fc457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191613ffc815184611f72565b9283156141535760005b848110614014575050505050565b8181101561413857614029612f6a8286611f21565b73ffffffffffffffffffffffffffffffffffffffff8116801561410e5761404f83611f64565b87811061406157505050600101614006565b848110156140de5773ffffffffffffffffffffffffffffffffffffffff61408b612f6a838a611f21565b16821461409a5760010161404f565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614109612f6a6141038885613ba5565b89611f21565b61408b565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b61414e612f6a6141488484613ba5565b85611f21565b614029565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91606061027f9294936141ca8160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6020818303126101855780519067ffffffffffffffff82116101855701604081830312610185576040519161423383610212565b815167ffffffffffffffff8111610185578161425091840161262c565b8352602082015167ffffffffffffffff811161018557614270920161262c565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff60806142aa855184602087015260c0860190610311565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b60409067ffffffffffffffff61036594931681528160208201520190610311565b909161430861253e565b5060208201908151156146ee5761433c6104bc6106de6104bc865173ffffffffffffffffffffffffffffffffffffffff1690565b9473ffffffffffffffffffffffffffffffffffffffff8616158015614663575b6146005782916143e960009261441c95516143c761438e895173ffffffffffffffffffffffffffffffffffffffff1690565b92614397610281565b95865267ffffffffffffffff8b16602087015273ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614278565b038183895af180156108c75761448594614572946144e9946000936145cd575b506144b161451591612f6a60009596519a6040519a8b91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018b528a61022e565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561022e565b61453a6104bc6104bc60025473ffffffffffffffffffffffffffffffffffffffff1690565b8351916040518097819482937f1e0185c3000000000000000000000000000000000000000000000000000000008452600484016142dd565b03915afa9283156108c7576000936145ad575b506020015192614593610281565b948552602085015260408401526060830152608082015290565b60209193506145c6903d806000833e61166e818361022e565b9290614585565b6000935061451591612f6a6145f66144b1933d8089833e6145ee818361022e565b8101906141ff565b955050915061443c565b610671614621855173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818a5afa9081156108c7576000916146cf575b501561435c565b6146e8915060203d6020116134e7576134d9818361022e565b386146c8565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9061472b602092828151948592016102ee565b0190565b6147f1956147e2601a6020937fff0000000000000000000000000000000000000000000000000000000000000060029d9c997fffffffffffffffff00000000000000000000000000000000000000000000000060019a8161472b9f9b81869c7f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b1660198201520191828151948592016102ee565b019160f81b1681520190614718565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b614907966147f1967fffff000000000000000000000000000000000000000000000000000000000000600160029c986103659f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a8498826149079b60f81b16815261489882518093602089850191016102ee565b019160f81b16838201526148b7825160029360208295850191016102ee565b01019160f01b16838201526148d68251809360206003850191016102ee565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b0190614718565b61027f909291926020604051948261492f87945180928580880191016102ee565b8301614943825180938580850191016102ee565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361022e565b606081019060ff82515111614c6157608081019060ff82515111614c325760c081019160ff83515111614c035760e082019060ff82515111614bd45761010083019361ffff85515111614ba557610120840193600185515111614b765761014081019261ffff84515111614b4757606095518051614b2b575b50815167ffffffffffffffff16906020830151614a0e9067ffffffffffffffff1690565b926040810151614a259067ffffffffffffffff1690565b9951908151614a349060ff1690565b9251918251614a439060ff1690565b60a09092015161ffff16936040519c8d976020890197614a629861472f565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752614a92908761022e565b51928351614aa09060ff1690565b9251908151614aaf9060ff1690565b9551928351614abf9061ffff1690565b938251614acd9061ffff1690565b9151948551614add9061ffff1690565b94604051998a9960208b0199614af29a61481d565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252614b22908261022e565b6103659161490e565b614b40919650614b3a90611f14565b516150a0565b94386149ea565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b9080602083519182815201916020808360051b8301019401926000915b838310614cbc57505050505090565b9091929394602080614d1f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610311565b97019301930191939290614cad565b909167ffffffffffffffff9273ffffffffffffffffffffffffffffffffffffffff6040840151169260608151910151916040519586947f9f01e16400000000000000000000000000000000000000000000000000000000865216600485015260806024850152614de2614dae825160a06084880152610124870190610311565b60208301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c8783030160a4880152610311565b906040810151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c8682030160c48701526020808451928381520193019060005b818110614f295750505060209593614eca8694614e9a8695608086614e686060614efa99015160e48b019073ffffffffffffffffffffffffffffffffffffffff169052565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c88830301610104890152610311565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016044870152614c90565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152610311565b03915afa9081156108c757600091614f10575090565b610365915060203d6020116108c0576108b1818361022e565b8251805173ffffffffffffffffffffffffffffffffffffffff168652602090810151818701528a985060409095019490920191600101614e23565b15614f6b57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6022614907966147f1967fff0000000000000000000000000000000000000000000000000000000000000060029b9781956103659f9e9c978e9984917f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b16602182015261506f82518093602089850191016102ee565b019160f81b168382015261508d8251809360206023850191016102ee565b01019160f81b1660018201520190614718565b602081019060ff825151116151be57604081019160ff8351511161518f57606082019160ff8351511161516057608081019261ffff8451511161512f576103659361182592519351906150f4825160ff1690565b965190615102825160ff1690565b935191615110835160ff1690565b91519461511f865161ffff1690565b946040519a8b9960208b01614fef565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526106716024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919290156152685750815115615201575090565b3b1561520a5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561527b5750805190602001fd5b6152b1906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610354565b0390fdfea164736f6c634300081a000a",
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
	return common.HexToHash("0x9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f")
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
