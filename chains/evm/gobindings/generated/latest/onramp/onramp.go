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
	Bin: "0x60e0604052346102f55760405161524438819003601f8101601f191683016001600160401b038111848210176102fa578392829160405283398101039060c082126102f557606082126102f557610054610310565b81519092906001600160401b03811681036102f55783526020820151906001600160a01b03821682036102f5576020840191825260606100966040850161032f565b6040860190815291605f1901126102f5576100af610310565b916100bc6060850161032f565b835260808401519384151585036102f55760a06100e091602086019687520161032f565b946040840195865233156102e457600180546001600160a01b0319163317905580516001600160401b03161580156102d2575b80156102c0575b61029357516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102ae575b80156102a4575b6102935780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610310565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610310565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051614f0090816103448239608051818181610b14015281816114730152611a49015260a05181611a82015260c051818181611abe0152611e480152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b038111838210176102fa57604052565b51906001600160a01b03821682036102f55756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635cb80c5d146100f357806366c3a5c7146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b6119a6565b6118b2565b611178565b610fa4565b610f02565b610eb0565b610dc7565b610cfb565b610c63565b610867565b610714565b610629565b610409565b610368565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576060610140611a29565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176101d557604052565b61018a565b6060810190811067ffffffffffffffff8211176101d557604052565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b6040519061027f6101608361022e565b565b6040519061027f60a08361022e565b6040519061027f60408361022e565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102e860208361022e565b60008252565b60005b8381106103015750506000910152565b81810151838201526020016102f1565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361034d815180928187528780880191016102ee565b0116010190565b906020610365928181520190610311565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103e560408051906103a9818361022e565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610311565b0390f35b67ffffffffffffffff81160361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435610444816103e9565b60243567ffffffffffffffff8111610185576104649036906004016103fb565b6000906104858367ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6104d56104bc845473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b16156105cf5761052892936105216105026104f36080850185611ae6565b906104fd87611cab565b612d49565b93845161051b6002610514878761305c565b9301611b37565b90613101565b83526137c5565b6000916000916000915b81518310156105c4576105ba61059261055c60019360606105538888611d7f565b51015190611dd0565b9661058c61057f602061056f8989611d7f565b51015167ffffffffffffffff1690565b67ffffffffffffffff1690565b90611dd0565b9461058c6105b160406105a58888611d7f565b51015163ffffffff1690565b63ffffffff1690565b9201919293610532565b604051908152602090f35b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576106636004356103e9565b60206106796024356106748161060b565b611de9565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101855760043567ffffffffffffffff81116101855760040160009280601f830112156107105781359367ffffffffffffffff851161070d57506020808301928560051b010111610185579190565b80fd5b8380fd5b346101855761072236610697565b9061074260035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b81811061074f57005b6107656104bc610760838587611ee2565b611ef2565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610862576001948891600093610832575b50826107da575b5050505001610746565b6107e5918391613abc565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a3388086816107d0565b61085491935060203d811161085b575b61084c818361022e565b810190611efc565b91386107c9565b503d610842565b611ddd565b346101855761087536610697565b9061087e613bbc565b6000915b80831061088b57005b610896838284611f0b565b926108a084611f4b565b9367ffffffffffffffff85169081158015610b08575b610ad05761090193949561091b60408301916108d28385611f55565b97906108fb60608701996108f36108e98c8a611f55565b9490923691611fc1565b923691611fc1565b90613c07565b67ffffffffffffffff166000526004602052604060002090565b602083019061096d61092c83611ef2565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61098461097a8486611f55565b906003840161202e565b61099b6109918886611f55565b906002840161202e565b608084016109ab6104bc82611ef2565b15610aa6576001977f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae95610a8683610a8c610a9b96610a326109ef610a7e98611ef2565b8f83019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610a75610a6e610a6860a0880193610a57610a4d868b611ae6565b90600484016120f5565b5460a01c67ffffffffffffffff1690565b9a611ef2565b9a86611f55565b97909686611f55565b949093611ef2565b94611ae6565b959094604051998a998a6122ad565b0390a2019190610882565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5067ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682146108b6565b906020808351928381520192019060005b818110610b595750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610b4c565b90610365916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015273ffffffffffffffffffffffffffffffffffffffff604083015116606082015260a0610c30610bfd606085015160c0608086015260e0850190610b3b565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610b3b565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610311565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610ca7816103e9565b606060a0604051610cb7816101b9565b600081526000602082015260006040820152828082015282608082015201521660005260046020526103e5610cef6040600020611cab565b60405191829182610b85565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610d32611a0a565b50604051610d3f816101da565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103e582606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff81163303610e86577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610f46816103e9565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610f8a5760405167ffffffffffffffff9091168152602090f35b611d93565b359061027f8261060b565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051610fe1816101da565b600435610fed8161060b565b8152602435610ffb81610f9a565b602082019081526044359061100f8261060b565b6040830191825261101e613bbc565b73ffffffffffffffffffffffffffffffffffffffff83511615918215611158575b50811561114d575b506111255780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611110611a29565b61111f60405192839283613d96565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538611047565b5173ffffffffffffffffffffffffffffffffffffffff161591503861103f565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576111b26004356103e9565b60243567ffffffffffffffff8111610185576111d29036906004016103fb565b604435906111e160643561060b565b60025460a01c60ff1661188857611232740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61125260043567ffffffffffffffff166000526004602052604060002090565b73ffffffffffffffffffffffffffffffffffffffff606435161561185e5780546112926104bc73ffffffffffffffffffffffffffffffffffffffff831681565b330361183457611318906112b66112ac6080860186611ae6565b906104fd86611cab565b9060006112de6104bc6104bc60025473ffffffffffffffffffffffffffffffffffffffff1690565b60608401519060405180809781947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610354565b03915afa92831561086257600093611817575b508251156117bd575b60a01c67ffffffffffffffff165b61134b90612394565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff0000000000000000000000000000000000000000161784556040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260148152939092906113d660348661022e565b6020838101516040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b16928101929092526014825290959061ffff1661142460348861022e565b61142e8880611ae6565b80611438926123b3565b61144560408b018b612437565b61144f91506124b6565b9861145d60208c018c611ae6565b94909561146861026f565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c015260608b01526114c860048801611beb565b60808b015261ffff1660a08a015260c089015236906114e692612523565b60e08701526114f36102d9565b6101008701526101208601968752369061150c92612523565b6101408501528251906115218760043561305c565b9060020161152e90611b37565b61153792613101565b825261154161255a565b9061154f83876004356137c5565b60208301908152946115646040880188612437565b905061172d575b5050611579839592956145af565b808252602081519101209061158f8651516125b3565b9460408201958652606060009401935b8751805182101561167f579060006115b8828c94611d7f565b518787838b61162960206115ef6115e96104bc6104bc8a5173ffffffffffffffffffffffffffffffffffffffff1690565b96611ef2565b960151604051998a97889687957f71c5c2ba00000000000000000000000000000000000000000000000000000000875230600488016126e5565b03925af1801561086257816116579160019460009161165e575b508a51906116518383611d7f565b52611d7f565b500161159f565b611679913d8091833e611671818361022e565b81019061235d565b38611643565b50508293507f9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f67ffffffffffffffff87936116ed6116cb60406103e59a015167ffffffffffffffff1690565b915194519551604051938493169667ffffffffffffffff60043516968461297c565b0390a461171d7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b600161173c6040890189612437565b9050036117935761177961178b9261175760408a018a612437565b6117746064939293359361176f600435933692612579565b612582565b613f2c565b90519061178582611d6d565b52611d6d565b50388061156b565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b915061134261180f6117cf8680611ae6565b91906117e360405193849260208401612383565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261022e565b929050611334565b61182d9193503d806000833e611671818361022e565b913861132b565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff6004356119028161060b565b61190a613bbc565b1633811461197c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576119e06004356103e9565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611a17826101da565b60006040838281528260208201520152565b611a31611a0a565b50604051611a3e816101da565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b906040519182815491828252602082019060005260206000209260005b818110611b6957505061027f9250038361022e565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201611b54565b90600182811c92168015611be1575b6020831014611bb257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611ba7565b9060405191826000825492611bff84611b98565b8084529360018116908115611c6b5750600114611c24575b5061027f9250038361022e565b90506000929192526020600020906000915b818310611c4f57505090602061027f9282010138611c17565b6020919350806001915483858901015201910190918492611c36565b6020935061027f9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611c17565b90604051611cb8816101b9565b60a0611d396004839567ffffffffffffffff80825473ffffffffffffffffffffffffffffffffffffffff81168852861c1616602086015273ffffffffffffffffffffffffffffffffffffffff80600183015416166040860152611d1d60028201611b37565b6060860152611d2e60038201611b37565b608086015201611beb565b910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805115611d7a5760200190565b611d3e565b8051821015611d7a5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9060018201809211610f8a57565b91908201809211610f8a57565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa801561086257600090611e93575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011611eda575b81611ead6020938361022e565b810103126101855773ffffffffffffffffffffffffffffffffffffffff9051611ed58161060b565b611e78565b3d9150611ea0565b9190811015611d7a5760051b0190565b356103658161060b565b90816020910312610185575190565b9190811015611d7a5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610185570190565b35610365816103e9565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160051b3603831361018557565b67ffffffffffffffff81116101d55760051b60200190565b929190611fcd81611fa9565b93611fdb604051958661022e565b602085838152019160051b810192831161018557905b828210611ffd57505050565b60208091833561200c8161060b565b815201910190611ff1565b818110612022575050565b60008155600101612017565b9067ffffffffffffffff83116101d5576801000000000000000083116101d5578154838355808410612092575b5090600052602060002060005b8381106120755750505050565b60019060208435946120868661060b565b01938184015501612068565b6120aa90836000528460206000209182019101612017565b3861205b565b9190601f81116120bf57505050565b61027f926000526020600020906020601f840160051c830193106120eb575b601f0160051c0190612017565b90915081906120de565b90929167ffffffffffffffff81116101d55761211b816121158454611b98565b846120b0565b6000601f821160011461217957819061216a93949560009261216e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612138565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946121ac84600052602060002090565b91805b8781106122055750836001959697106121cd575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806121c3565b909260206001819286860135815501940191016121af565b9160209082815201919060005b8181106122375750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356122608161060b565b16815201940192910161222a565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916123079373ffffffffffffffffffffffffffffffffffffffff95866122f99367ffffffffffffffff6103659e9c9d9b96168a5216602089015260c0604089015260c088019161221d565b91858303606087015261221d565b9416608082015260a081850391015261226e565b81601f820112156101855780516123318161029f565b9261233f604051948561022e565b818452602082840101116101855761036591602080850191016102ee565b9060208282031261018557815167ffffffffffffffff811161018557610365920161231b565b91602061036593818152019161226e565b67ffffffffffffffff1667ffffffffffffffff8114610f8a5760010190565b9092919283600c1161018557831161018557600c01917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40190565b906004116101855790600490565b909291928360041161018557831161018557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b60405190612498826101f6565b60606080836000815282602082015282604082015282808201520152565b906124c082611fa9565b6124cd604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06124fb8294611fa9565b019060005b82811061250c57505050565b60209061251761248b565b82828501015201612500565b92919261252f8261029f565b9161253d604051938461022e565b829481845281830111610185578281602093846000960137010152565b60405190612567826101da565b60606040838281528260208201520152565b9015611d7a5790565b91908260409103126101855760405161259a81610212565b602080829480356125aa8161060b565b84520135910152565b906125bd82611fa9565b6125ca604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125f88294611fa9565b019060005b82811061260957505050565b8060606020809385010152016125fd565b9080602083519182815201916020808360051b8301019401926000915b83831061264657505050505090565b90919293946020806126d6837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260806126c56126b36126a18786015160a08987015260a0860190610311565b60408601518582036040870152610311565b60608501518482036060860152610311565b920151906080818403910152610311565b97019301930191939290612637565b929491936128ce610365979573ffffffffffffffffffffffffffffffffffffffff6128f39416865260c0602087015261272b60c08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660e0870152604081015167ffffffffffffffff1661010087015261014061289b61286561282f6127fa6127b66127838c61022060608a0151916101606101208201520190610311565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4001888f0152610311565b60a087015161ffff166101608d015260c08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40016101808e0152610311565b60e08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408c8303016101a08d0152610311565b6101008501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408b8303016101c08c0152610311565b6101208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a8303016101e08b015261261a565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4087830301610200880152610311565b956040850152606084019073ffffffffffffffffffffffffffffffffffffffff169052565b608082015260a0818403910152610311565b9080602083519182815201916020808360051b8301019401926000915b83831061293157505050505090565b909192939460208061296d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610311565b97019301930191939290612922565b93929061299190606086526060860190610311565b938085036020820152825180865260208601906020808260051b8901019501916000905b8282106129d357505050506103659394506040818403910152612905565b90919295602080612a63837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610311565b9801920192019092916129b5565b60405190612a7e826101f6565b6060608083828152600060208201526000604082015282808201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612ad1575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90612b0d82611fa9565b612b1a604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612b488294611fa9565b019060005b828110612b5957505050565b602090604051612b6881610212565b6000815260608382015282828501015201612b4d565b9080601f830112156101855781602061036593359101612523565b9080601f8301121561018557813591612bb183611fa9565b92612bbf604051948561022e565b80845260208085019160051b830101918383116101855760208101915b838310612beb57505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101855760405190612c3882610212565b6020830135612c468161060b565b825260408301359167ffffffffffffffff831161018557612c6f88602080969581960101612b7e565b83820152815201920191612bdc565b359061ffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160a08183031261018557612cbc610281565b91813567ffffffffffffffff81116101855781612cda918401612b99565b8352612ce860208301612c7e565b6020840152612cf960408301610f8f565b6040840152606082013567ffffffffffffffff81116101855781612d1e918401612b7e565b6060840152608082013567ffffffffffffffff811161018557612d419201612b7e565b608082015290565b9190612d53612a71565b600483101580612fff575b15612f7c575081612d7a92612d72926123fc565b810190612c8d565b9081515160005b818110612eb257505081515115612e13575b6040820173ffffffffffffffffffffffffffffffffffffffff612dca825173ffffffffffffffffffffffffffffffffffffffff1690565b1615612dd557505090565b612df9604061036593015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b6080810192612e23845151612b03565b835260005b84518051821015612ea95790612ea281612e61612e4782600196611d7f565b5173ffffffffffffffffffffffffffffffffffffffff1690565b612e88612e6c610290565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b612e906102d9565b60208201528751906116518383611d7f565b5001612e28565b50509250612d93565b612ebb81611dc2565b828110612ecb5750600101612d81565b612ef4612ed9838751611d7f565b515173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff612f1a6104bc612ed9858a51611d7f565b911614612f2957600101612ebb565b610607612f3a612ed9848851611d7f565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b919392612f8a368684612523565b60608401526080810193612f9f855151612b03565b845260005b85518051821015612ff25790612feb81612fc3612e4782600196611d7f565b612fce612e6c610290565b612fd9368c8a612523565b60208201528851906116518383611d7f565b5001612fa4565b5050945091925050612d93565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061305561304f86866123ee565b90612a9d565b1614612d5e565b9061306b906040810190612437565b9050156130945767ffffffffffffffff1660005260046020526103656003604060002001611b37565b5060405160206130a4818361022e565b600082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810190369083013790565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610f8a5760010190565b61312161311c6131148351855190611dd0565b855190611dd0565b612b03565b91600094855b8351811015613164578061315d61314060019387611d7f565b519861314b816130d4565b99613156828a611d7f565b5287611d7f565b5001613127565b50939094915060005b855181101561323f57613183612e478288611d7f565b600073ffffffffffffffffffffffffffffffffffffffff8216815b868110613215575b5050156131b7575b5060010161316d565b929061320e6001926131e66131ca610290565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b6131ee6102d9565b60208701526131fc816130d4565b956132078289611d7f565b5286611d7f565b50906131ae565b816132266104bc612ed9848c611d7f565b146132335760010161319e565b505050600138806131a6565b5093506000905b845182101561331a5761325c612e478387611d7f565b91600073ffffffffffffffffffffffffffffffffffffffff8416815b8481106132f0575b505015613292575b6001019150613246565b6132e76001926132bf6132a3610290565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b6132c76102d9565b60208601526132d5816130d4565b946132e08288611d7f565b5285611d7f565b50829150613288565b816133016104bc612ed9848b611d7f565b1461330e57600101613278565b50505060013880613280565b8252509150565b9061332b82611fa9565b613338604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06133668294611fa9565b019060005b82811061337757505050565b602090604051613386816101f6565b600081526000838201526000604082015260006060820152606060808201528282850101520161336b565b519063ffffffff8216820361018557565b908160609103126101855780519161036560406133e1602085016133b1565b93016133b1565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b9160209082815201919060005b8181106134525750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff873561347b8161060b565b16815260208781013590820152019401929101613445565b959492939173ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff9216875216602086015260a060408601526135256134e86134d785806133e8565b60a0808a015261014089019161226e565b6134f560208601866133e8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608984030160c08a015261226e565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610185578401916020833593019167ffffffffffffffff8411610185578360061b360383136101855761027f956135f76135c083608097613636978d60e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60826136289a0301910152613438565b916135ee6135d060608301610f8f565b73ffffffffffffffffffffffffffffffffffffffff166101008d0152565b868101906133e8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8403016101208c015261226e565b908782036060890152610311565b94019061ffff169052565b81601f8201121561018557803561365781611fa9565b92613665604051948561022e565b81845260208085019260061b8401019281841161018557602001915b83831061368f575050505090565b602060409161369e8486612582565b815201920191613681565b91909160a081840312610185576136be610281565b92813567ffffffffffffffff811161018557816136dc918401612b7e565b8452602082013567ffffffffffffffff811161018557816136fe918401612b7e565b6020850152604082013567ffffffffffffffff81116101855781613723918401613641565b604085015261373460608301610f8f565b6060850152608082013567ffffffffffffffff8111610185576137579201612b7e565b6080830152565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610f8a57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211610f8a57565b91908203918211610f8a57565b9291908151516137f56137f06137eb60408501936137e38587612437565b919050611dd0565b611dc2565b613321565b9060005b84518051821015613985578161380e91611d7f565b51906138346104bc6104bc845173ffffffffffffffffffffffffffffffffffffffff1690565b60208301906060825161384c60208b015161ffff1690565b92898d613889604051968795869485947f2b0560010000000000000000000000000000000000000000000000000000000086523060048701613493565b03915afa938415610862576001946000918293839261393d575b509061390a63ffffffff6138ce613917945173ffffffffffffffffffffffffffffffffffffffff1690565b9651956138f86138dc610281565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b1667ffffffffffffffff166020870152565b63ffffffff166040850152565b6060830152608082015261392b8286611d7f565b526139368185611d7f565b50016137f9565b63ffffffff94506138ce9350613917925061397161390a9160603d811161397e575b613969818361022e565b8101906133c2565b90965090945092506138a3565b503d61395f565b5050909394613a20909392936139c66139b5604086015173ffffffffffffffffffffffffffffffffffffffff1690565b916139c036886136a9565b8661496c565b6060850151906139f36139d7610281565b73ffffffffffffffffffffffffffffffffffffffff9094168452565b600060208401526000604084015260608301526080820152613a15865161375e565b906132e08288611d7f565b50613a2b8184612437565b9050613a38575b50505090565b613a53610760613a4d608093613ab396612437565b90612579565b910151613a7d613a61610281565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b6000602083015260006040830152600060608301526080820152613aa1835161378b565b90613aac8285611d7f565b5282611d7f565b50388080613a32565b9073ffffffffffffffffffffffffffffffffffffffff613b8e9392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452613b1e60648561022e565b16600080604093845195613b32868861022e565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15613bb3573d613b7f613b768261029f565b9451948561022e565b83523d6000602085013e614e2b565b805180613b99575050565b81602080613bae9361027f9501019101613e18565b614ba2565b60609250614e2b565b73ffffffffffffffffffffffffffffffffffffffff600154163303613bdd57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191613c15815184611dd0565b928315613d6c5760005b848110613c2d575050505050565b81811015613d5157613c42612e478286611d7f565b73ffffffffffffffffffffffffffffffffffffffff81168015613d2757613c6883611dc2565b878110613c7a57505050600101613c1f565b84811015613cf75773ffffffffffffffffffffffffffffffffffffffff613ca4612e47838a611d7f565b168214613cb357600101613c68565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff613d22612e47613d1c88856137b8565b89611d7f565b613ca4565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b613d67612e47613d6184846137b8565b85611d7f565b613c42565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91606061027f929493613de38160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b90816020910312610185575161036581610f9a565b6020818303126101855780519067ffffffffffffffff821161018557016040818303126101855760405191613e6183610212565b815167ffffffffffffffff81116101855781613e7e91840161231b565b8352602082015167ffffffffffffffff811161018557613e9e920161231b565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080613ed8855184602087015260c0860190610311565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b60409067ffffffffffffffff61036594931681528160208201520190610311565b9091613f3661248b565b50602082019081511561432c57613f6a6104bc6106746104bc865173ffffffffffffffffffffffffffffffffffffffff1690565b9473ffffffffffffffffffffffffffffffffffffffff8616158015614291575b61422e57829161401760009261404a9551613ff5613fbc895173ffffffffffffffffffffffffffffffffffffffff1690565b92613fc5610281565b95865267ffffffffffffffff8b16602087015273ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301613ea6565b038183895af18015610862576140b3946141a094614117946000936141fb575b506140df61414391612e4760009596519a6040519a8b91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018b528a61022e565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561022e565b6141686104bc6104bc60025473ffffffffffffffffffffffffffffffffffffffff1690565b8351916040518097819482937f1e0185c300000000000000000000000000000000000000000000000000000000845260048401613f0b565b03915afa928315610862576000936141db575b5060200151926141c1610281565b948552602085015260408401526060830152608082015290565b60209193506141f4903d806000833e611671818361022e565b92906141b3565b6000935061414391612e476142246140df933d8089833e61421c818361022e565b810190613e2d565b955050915061406a565b61060761424f855173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818a5afa908115610862576000916142fd575b5015613f8a565b61431f915060203d602011614325575b614317818361022e565b810190613e18565b386142f6565b503d61430d565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b90614369602092828151948592016102ee565b0190565b61442f95614420601a6020937fff0000000000000000000000000000000000000000000000000000000000000060029d9c997fffffffffffffffff00000000000000000000000000000000000000000000000060019a816143699f9b81869c7f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b1660198201520191828151948592016102ee565b019160f81b1681520190614356565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b6145459661442f967fffff000000000000000000000000000000000000000000000000000000000000600160029c986103659f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a8498826145459b60f81b1681526144d682518093602089850191016102ee565b019160f81b16838201526144f5825160029360208295850191016102ee565b01019160f01b16838201526145148251809360206003850191016102ee565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b0190614356565b61027f909291926020604051948261456d87945180928580880191016102ee565b8301614581825180938580850191016102ee565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361022e565b606081019060ff8251511161489f57608081019060ff825151116148705760c081019160ff835151116148415760e082019060ff825151116148125761010083019361ffff855151116147e3576101208401936001855151116147b45761014081019261ffff8451511161478557606095518051614769575b50815167ffffffffffffffff1690602083015161464c9067ffffffffffffffff1690565b9260408101516146639067ffffffffffffffff1690565b99519081516146729060ff1690565b92519182516146819060ff1690565b60a09092015161ffff16936040519c8d9760208901976146a09861436d565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810187526146d0908761022e565b519283516146de9060ff1690565b92519081516146ed9060ff1690565b95519283516146fd9061ffff1690565b93825161470b9061ffff1690565b915194855161471b9061ffff1690565b94604051998a9960208b01996147309a61445b565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252614760908261022e565b6103659161454c565b61477e91965061477890611d6d565b51614cde565b9438614628565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b9080602083519182815201916020808360051b8301019401926000915b8383106148fa57505050505090565b909192939460208061495d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610311565b970193019301919392906148eb565b909167ffffffffffffffff9273ffffffffffffffffffffffffffffffffffffffff6040840151169260608151910151916040519586947f9f01e16400000000000000000000000000000000000000000000000000000000865216600485015260806024850152614a206149ec825160a06084880152610124870190610311565b60208301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c8783030160a4880152610311565b906040810151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c8682030160c48701526020808451928381520193019060005b818110614b675750505060209593614b088694614ad88695608086614aa66060614b3899015160e48b019073ffffffffffffffffffffffffffffffffffffffff169052565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c88830301610104890152610311565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8683030160448701526148ce565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152610311565b03915afa90811561086257600091614b4e575090565b610365915060203d60201161085b5761084c818361022e565b8251805173ffffffffffffffffffffffffffffffffffffffff168652602090810151818701528a985060409095019490920191600101614a61565b15614ba957565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60226145459661442f967fff0000000000000000000000000000000000000000000000000000000000000060029b9781956103659f9e9c978e9984917f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152614cad82518093602089850191016102ee565b019160f81b1683820152614ccb8251809360206023850191016102ee565b01019160f81b1660018201520190614356565b602081019060ff82515111614dfc57604081019160ff83515111614dcd57606082019160ff83515111614d9e57608081019261ffff84515111614d6d57610365936117e39251935190614d32825160ff1690565b965190614d40825160ff1690565b935191614d4e835160ff1690565b915194614d5d865161ffff1690565b946040519a8b9960208b01614c2d565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526106076024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91929015614ea65750815115614e3f575090565b3b15614e485790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614eb95750805190602001fd5b614eef906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610354565b0390fdfea164736f6c634300081a000a",
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
