// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_proxy

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

type CCVProxyDestChainConfig struct {
	Router           common.Address
	SequenceNumber   uint64
	DefaultExecutor  common.Address
	LaneMandatedCCVs []common.Address
	DefaultCCVs      []common.Address
	CcvAggregator    []byte
}

type CCVProxyDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	DefaultCCVs       []common.Address
	LaneMandatedCCVs  []common.Address
	DefaultExecutor   common.Address
	CcvAggregator     []byte
}

type CCVProxyDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type CCVProxyReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint64
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

type CCVProxyStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
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

var CCVProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"verifierReceipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCCVProxy.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executorReceipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVInUserInput\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enumMessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc57604051614feb38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051614ca0908161034b8239608051818181610aa9015281816114420152611c5d015260a0518181816104c60152611c96015260c051818181611cd20152611d7a0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635cb80c5d146100f357806366c3a5c7146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611bba565b611ac6565b61110d565b610f39565b610e97565b610e45565b610d5c565b610c90565b610bf8565b6107fc565b610695565b6105aa565b610418565b610377565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576060610140611c3d565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176101d557604052565b61018a565b6060810190811067ffffffffffffffff8211176101d557604052565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b6040519061027f6101608361022e565b565b6040519061027f60a08361022e565b6040519061027f60e08361022e565b6040519061027f60408361022e565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102f760208361022e565b60008252565b60005b8381106103105750506000910152565b8181015183820152602001610300565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361035c815180928187528780880191016102fd565b0116010190565b906020610374928181520190610320565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103f460408051906103b8818361022e565b601282527f43435650726f787920312e372e302d6465760000000000000000000000000000602083015251918291602083526020830190610320565b0390f35b67ffffffffffffffff81160361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435610453816103f8565b60243567ffffffffffffffff81116101855761047390369060040161040a565b506040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561058757600091610558575b5061051d5760405160008152602090f35b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b61057a915060203d602011610580575b610572818361022e565b810190611cfa565b3861050c565b503d610568565b611d0f565b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576105e46004356103f8565b60206105fa6024356105f58161058c565b611d1b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101855760043567ffffffffffffffff81116101855760040160009280601f830112156106915781359367ffffffffffffffff851161068e57506020808301928560051b010111610185579190565b80fd5b8380fd5b34610185576106a336610618565b906106c360035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b8181106106d057005b6106ff6106e66106e1838587611e43565b611e58565b73ffffffffffffffffffffffffffffffffffffffff1690565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156105875760019488916000936107cc575b5082610774575b50505050016106c7565b61077f918391612e80565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a33880868161076a565b6107ee91935060203d81116107f5575b6107e6818361022e565b810190611e62565b9138610763565b503d6107dc565b346101855761080a36610618565b90610813612f80565b6000915b80831061082057005b61082b838284611e71565b9261083584611eb1565b9367ffffffffffffffff85169081158015610a9d575b610a65576108969394956108b060408301916108678385611ebb565b9790610890606087019961088861087e8c8a611ebb565b9490923691611f27565b923691611f27565b90613005565b67ffffffffffffffff166000526004602052604060002090565b60208301906109026108c183611e58565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61091961090f8486611ebb565b9060038401611fc3565b6109306109268886611ebb565b9060028401611fc3565b608084016109406106e682611e58565b15610a3b576001977f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae95610a1b83610a21610a30966109c7610984610a1398611e58565b8f83019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610a0a610a036109fd60a08801936109ec6109e2868b612045565b906004840161212e565b5460a01c67ffffffffffffffff1690565b9a611e58565b9a86611ebb565b97909686611ebb565b949093611e58565b94612045565b959094604051998a998a6122e6565b0390a2019190610817565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016821461084b565b906020808351928381520192019060005b818110610aee5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610ae1565b90610374916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015273ffffffffffffffffffffffffffffffffffffffff604083015116606082015260a0610bc5610b92606085015160c0608086015260e0850190610ad0565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610ad0565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610320565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610c3c816103f8565b606060a0604051610c4c816101b9565b600081526000602082015260006040820152828082015282608082015201521660005260046020526103f4610c846040600020612475565b60405191829182610b1a565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610cc7611c1e565b50604051610cd4816101da565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103f482606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff81163303610e1b577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610edb816103f8565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610f1f5760405167ffffffffffffffff9091168152602090f35b611f7d565b359061027f8261058c565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051610f76816101da565b600435610f828161058c565b8152602435610f9081610f2f565b6020820190815260443590610fa48261058c565b60408301918252610fb3612f80565b73ffffffffffffffffffffffffffffffffffffffff835116159182156110ed575b5081156110e2575b506110ba5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3906110a5611c3d565b6110b4604051928392836131ae565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fdc565b5173ffffffffffffffffffffffffffffffffffffffff1615915038610fd4565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435611148816103f8565b60243567ffffffffffffffff81116101855761116890369060040161040a565b6044356064356111778161058c565b60025460a01c60ff16611a9c576111c8740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6111e68467ffffffffffffffff166000526004602052604060002090565b9373ffffffffffffffffffffffffffffffffffffffff821615611a72578454926112266106e673ffffffffffffffffffffffffffffffffffffffff861681565b3303611a48576112b061124e61123f6080880188612045565b906112498a612475565b61354f565b9260006112766106e66106e660025473ffffffffffffffffffffffffffffffffffffffff1690565b60a08601519060405180809681947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610363565b03915afa928315610587578792600094611a2b575b508351156119d6575b6112d6612581565b9685518a60208801998a5160408a019384516112f29060ff1690565b926112fc94613996565b60ff16909252908952865260a01c67ffffffffffffffff1661131d9061259d565b89547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178a556040513060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526014815290926113a660348361022e565b606087810151604051918a901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602083015260148252909b9061ffff166113f060348e61022e565b6113fa8780612045565b80611404926125bc565b604089019e8f611414908b612640565b61141e91506126bf565b996020810161142c91612045565b94909561143761026f565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529867ffffffffffffffff891660208b015260408a019b61148e908d9067ffffffffffffffff169052565b60608a015260040161149f906123b5565b608089015261ffff1660a088015260c087015236906114bd9261272c565b60e08501526114ca6102e8565b610100850152610120840196875236906114e39261272c565b6101408301528651518951516114f891612771565b6115019061277e565b9860005b8851805182101561159f57908b61158d826115228160019661281b565b51926020611544855173ffffffffffffffffffffffffffffffffffffffff1690565b94015161156e611552610281565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b600060208601526000604086015260006060860152608085015261281b565b52611598818d61281b565b5001611505565b50508a8c9a98999a60005b8351805182101561165057906116498c8e61164261163b856115ce8160019961281b565b519360206115f0865173ffffffffffffffffffffffffffffffffffffffff1690565b95015161161a6115fe610281565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b60006020870152600060408701526000606087015260808601525151612771565b809361281b565b528d61281b565b50016115aa565b50508b8b9697999a989a608088015161167c9073ffffffffffffffffffffffffffffffffffffffff1690565b9b60a089015161168a610281565b73ffffffffffffffffffffffffffffffffffffffff909e168e52600060208f0181905260408f0181905260608f015260808e0152866116c936876128e3565b906116d4918b613d7e565b506116df8486612640565b905061194b575b5050505060606116f69101611e58565b9361170084614623565b97885160208a01209361172061171b84515186515190612771565b6129a1565b9960005b845180518210156117fa579060008a61173e838e9561281b565b518a838d6117a5602061176b6106e66106e6885173ffffffffffffffffffffffffffffffffffffffff1690565b950151604051998a97889687957f71c5c2ba0000000000000000000000000000000000000000000000000000000087523060048801612ad3565b03925af191821561058757818e6001946117d2936000916117d9575b506117cc838361281b565b5261281b565b5001611724565b6117f4913d8091833e6117ec818361022e565b81019061254a565b386117c1565b5050919399979a9890929860005b8b5180518210156118af579060008a8f93836118239161281b565b518a838d611850602061176b6106e66106e6885173ffffffffffffffffffffffffffffffffffffffff1690565b03925af18015610587578b61188391838f60019660009361188a575b5061187c9161163b915151612771565b528c61281b565b5001611808565b916118a661187c9361163b93953d8091833e6117ec818361022e565b9391509161186c565b50506103f496507fdc37a122d47e708a593d43fba77d7a22899a573dfbd0cd4423c7d41a6291e0ff9161190b8a67ffffffffffffffff936118f98a995167ffffffffffffffff1690565b93856040519687961699169785612dcb565b0390a461193b7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b600161195a8587959697612640565b9050036119ac576119916060946119a3938961198c61198761197f6116f69b8a612640565b369291612998565b61284a565b614064565b90519061199d8261280e565b5261280e565b50918b806116e6565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b9250611a256119e58380612045565b91906119f960405193849260208401612570565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261022e565b926112ce565b611a419194503d806000833e6117ec818361022e565b92386112c5565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff600435611b168161058c565b611b1e612f80565b16338114611b9057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557611bf46004356103f8565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611c2b826101da565b60006040838281528260208201520152565b611c45611c1e565b50604051611c52816101da565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610185575161037481610f2f565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa801561058757600090611dc5575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011611e0c575b81611ddf6020938361022e565b810103126101855773ffffffffffffffffffffffffffffffffffffffff9051611e078161058c565b611daa565b3d9150611dd2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611e535760051b0190565b611e14565b356103748161058c565b90816020910312610185575190565b9190811015611e535760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610185570190565b35610374816103f8565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160051b3603831361018557565b67ffffffffffffffff81116101d55760051b60200190565b929190611f3381611f0f565b93611f41604051958661022e565b602085838152019160051b810192831161018557905b828210611f6357505050565b602080918335611f728161058c565b815201910190611f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110611fb7575050565b60008155600101611fac565b9067ffffffffffffffff83116101d5576801000000000000000083116101d5578154838355808410612027575b5090600052602060002060005b83811061200a5750505050565b600190602084359461201b8661058c565b01938184015501611ffd565b61203f90836000528460206000209182019101611fac565b38611ff0565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b90600182811c921680156120df575b60208310146120b057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916120a5565b9190601f81116120f857505050565b61027f926000526020600020906020601f840160051c83019310612124575b601f0160051c0190611fac565b9091508190612117565b90929167ffffffffffffffff81116101d5576121548161214e8454612096565b846120e9565b6000601f82116001146121b25781906121a39394956000926121a7575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612171565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946121e584600052602060002090565b91805b87811061223e575083600195969710612206575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806121fc565b909260206001819286860135815501940191016121e8565b9160209082815201919060005b8181106122705750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356122998161058c565b168152019401929101612263565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916123409373ffffffffffffffffffffffffffffffffffffffff95866123329367ffffffffffffffff6103749e9c9d9b96168a5216602089015260c0604089015260c0880191612256565b918583036060870152612256565b9416608082015260a08185039101526122a7565b906040519182815491828252602082019060005260206000209260005b81811061238657505061027f9250038361022e565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612371565b90604051918260008254926123c984612096565b808452936001811690811561243557506001146123ee575b5061027f9250038361022e565b90506000929192526020600020906000915b81831061241957505090602061027f92820101386123e1565b6020919350806001915483858901015201910190918492612400565b6020935061027f9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386123e1565b90604051612482816101b9565b60a06125036004839567ffffffffffffffff80825473ffffffffffffffffffffffffffffffffffffffff81168852861c1616602086015273ffffffffffffffffffffffffffffffffffffffff806001830154161660408601526124e760028201612354565b60608601526124f860038201612354565b6080860152016123b5565b910152565b81601f8201121561018557805161251e816102ae565b9261252c604051948561022e565b818452602082840101116101855761037491602080850191016102fd565b9060208282031261018557815167ffffffffffffffff8111610185576103749201612508565b9160206103749381815201916122a7565b6040519061259060208361022e565b6000808352366020840137565b67ffffffffffffffff1667ffffffffffffffff8114610f1f5760010190565b9092919283600c1161018557831161018557600c01917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40190565b906004116101855790600490565b909291928360041161018557831161018557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b604051906126a1826101f6565b60606080836000815282602082015282604082015282808201520152565b906126c982611f0f565b6126d6604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06127048294611f0f565b019060005b82811061271557505050565b602090612720612694565b82828501015201612709565b929192612738826102ae565b91612746604051938461022e565b829481845281830111610185578281602093846000960137010152565b9060018201809211610f1f57565b91908201809211610f1f57565b9061278882611f0f565b612795604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06127c38294611f0f565b019060005b8281106127d457505050565b6020906040516127e3816101f6565b60008152600083820152600060408201526000606082015260606080820152828285010152016127c8565b805115611e535760200190565b8051821015611e535760209160051b010190565b9080601f83011215610185578160206103749335910161272c565b91908260409103126101855760405161286281610212565b602080829480356128728161058c565b84520135910152565b81601f8201121561018557803561289181611f0f565b9261289f604051948561022e565b81845260208085019260061b8401019281841161018557602001915b8383106128c9575050505090565b60206040916128d8848661284a565b8152019201916128bb565b91909160a081840312610185576128f8610281565b92813567ffffffffffffffff8111610185578161291691840161282f565b8452602082013567ffffffffffffffff8111610185578161293891840161282f565b6020850152604082013567ffffffffffffffff8111610185578161295d91840161287b565b604085015261296e60608301610f24565b6060850152608082013567ffffffffffffffff811161018557612991920161282f565b6080830152565b9015611e535790565b906129ab82611f0f565b6129b8604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129e68294611f0f565b019060005b8281106129f757505050565b8060606020809385010152016129eb565b9080602083519182815201916020808360051b8301019401926000915b838310612a3457505050505090565b9091929394602080612ac4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080612ab3612aa1612a8f8786015160a08987015260a0860190610320565b60408601518582036040870152610320565b60608501518482036060860152610320565b920151906080818403910152610320565b97019301930191939290612a25565b92949193612cbc610374979573ffffffffffffffffffffffffffffffffffffffff612ce19416865260c06020870152612b1960c08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660e0870152604081015167ffffffffffffffff16610100870152610140612c89612c53612c1d612be8612ba4612b718c61022060608a0151916101606101208201520190610320565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4001888f0152610320565b60a087015161ffff166101608d015260c08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40016101808e0152610320565b60e08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408c8303016101a08d0152610320565b6101008501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408b8303016101c08c0152610320565b6101208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a8303016101e08b0152612a08565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4087830301610200880152610320565b956040850152606084019073ffffffffffffffffffffffffffffffffffffffff169052565b608082015260a0818403910152610320565b9060a060806103749373ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff602082015116602085015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610320565b9080602083519182815201916020808360051b8301019401926000915b838310612d8057505050505090565b9091929394602080612dbc837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610320565b97019301930191939290612d71565b9294939190612de290608085526080850190610320565b948386036020850152815180875260208701906020808260051b8a01019401916000905b828210612e37575050505061037494955090612e29918482036040860152612cf3565b916060818403910152612d54565b90919294602080612e72837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08e600196030186528951612cf3565b970192019201909291612e06565b9073ffffffffffffffffffffffffffffffffffffffff612f529392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452612ee260648561022e565b16600080604093845195612ef6868861022e565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15612f77573d612f43612f3a826102ae565b9451948561022e565b83523d6000602085013e614bcb565b805180612f5d575050565b81602080612f729361027f9501019101611cfa565b614942565b60609250614bcb565b73ffffffffffffffffffffffffffffffffffffffff600154163303612fa157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610f1f57565b91908203918211610f1f57565b805191613013815184612771565b9283156131845760005b84811061302b575050505050565b818110156131695761305a613040828661281b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8116801561313f5761308083612763565b8781106130925750505060010161301d565b8481101561310f5773ffffffffffffffffffffffffffffffffffffffff6130bc613040838a61281b565b1682146130cb57600101613080565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61313a6130406131348885612ff8565b8961281b565b6130bc565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b61317f6130406131798484612ff8565b8561281b565b61305a565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91606061027f9294936131fb8160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6040519060e0820182811067ffffffffffffffff8211176101d557604052606060c08382815282602082015260006040820152600083820152600060808201528260a08201520152565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106132ae575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b906132ea82611f0f565b6132f7604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06133258294611f0f565b019060005b82811061333657505050565b60209060405161334581610212565b600081526060838201528282850101520161332a565b9080601f830112156101855781359161337383611f0f565b92613381604051948561022e565b80845260208085019160051b830101918383116101855760208101915b8383106133ad57505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261018557604051906133fa82610212565b60208301356134088161058c565b825260408301359167ffffffffffffffff8311610185576134318860208096958196010161282f565b8382015281520192019161339e565b359060ff8216820361018557565b359061ffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160e0818303126101855761348c610290565b91813567ffffffffffffffff811161018557816134aa91840161335b565b8352602082013567ffffffffffffffff811161018557816134cc91840161335b565b60208401526134dd60408301613440565b60408401526134ee6060830161344e565b60608401526134ff60808301610f24565b608084015260a082013567ffffffffffffffff8111610185578161352491840161282f565b60a084015260c082013567ffffffffffffffff811161018557613547920161282f565b60c082015290565b9190613559613230565b6004831015806138c6575b156138435750816135809261357892612605565b81019061345d565b9060208201805151806137ee575b5082515161359e82515182612771565b60005b8181106136c6575050506135bb9083515190515190612771565b15613641575b6080820173ffffffffffffffffffffffffffffffffffffffff6135f8825173ffffffffffffffffffffffffffffffffffffffff1690565b161561360357505090565b613627604061037493015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b60808101926136518451516132e0565b835260005b845180518210156136bd57906136b6816136756130408260019661281b565b61369c61368061029f565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b6136a46102e8565b60208201528751906117cc838361281b565b5001613656565b505092506135c1565b828110156137d7576136f76136dc82885161281b565b515173ffffffffffffffffffffffffffffffffffffffff1690565b61370082612763565b8381106137115750506001016135a1565b848110156137a55773ffffffffffffffffffffffffffffffffffffffff61373c6136dc838b5161281b565b1673ffffffffffffffffffffffffffffffffffffffff83161461376157600101613700565b7fd757e5e80000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6137d26136dc88516137cc8986612ff8565b9061281b565b61373c565b6137e96136dc85516137cc8685612ff8565b6136f7565b604084015160ff1690811080159190613837575b5061380d573861358e565b7fb273a0e40000000000000000000000000000000000000000000000000000000060005260046000fd5b60ff9150161538613802565b91939261385136868461272c565b60a084015260808101936138668551516132e0565b845260005b855180518210156138b957906138b28161388a6130408260019661281b565b61389561368061029f565b6138a0368c8a61272c565b60208201528851906117cc838361281b565b500161386b565b50509450919250506135c1565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061391c61391686866125f7565b9061327a565b1614613564565b8054821015611e535760005260206000200190600090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610f1f5760010190565b60ff168015610f1f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b6002909593950180546139aa835182612771565b926139b4846132e0565b9260009485905b808210613a6c575050505050816139d3575050929190565b91949092936139eb6139e6858851612771565b6132e0565b9260005b858110613a4257505060005b8651811015613a385780613a31613a146001938a61281b565b51613a1f8389612771565b90613a2a828961281b565b528661281b565b50016139fb565b5091945092909150565b80613a4f6001928461281b565b51613a5a828861281b565b52613a65818761281b565b50016139ef565b93999298919790969195929492938a881015613cc457613aaf613a8f898b613923565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b985b6000805b878110613c85575b50613c725760009a8b5b8551811015613c5f57613ae06106e66136dc838961281b565b73ffffffffffffffffffffffffffffffffffffffff8d1614613b0457600101613ac7565b5093989195979a5093959a919860015b15613b2a575b506001905b0190939291926139bb565b999792613b688b613b4c613b458b9d98959d9996999b61393b565b9a8a61281b565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60005b8c51811015613c4d578c8c73ffffffffffffffffffffffffffffffffffffffff613b9b6106e66136dc868661281b565b911614613bab5750600101613b6b565b613bff92959a9c50613bc6826020929d95989d99969961281b565b5101516020613bdd613bd78c612fcb565b8b61281b565b5101528c61187c82613bf8613bf28451612fcb565b8461281b565b519261281b565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8b51018b5260ff8a16613c39575b6001905b90613b1a565b98613c45600191613968565b999050613c2f565b50929799509297600190949194613c33565b5093989195979a9b92999094969b613b14565b93959a9198509395989196600190613b1f565b8b73ffffffffffffffffffffffffffffffffffffffff80613ca96136dc858c61281b565b9216911614613cba57600101613ab5565b5050600138613abd565b613cda613040613cd48d8b612ff8565b8c61281b565b98613ab1565b9080602083519182815201916020808360051b8301019401926000915b838310613d0c57505050505090565b9091929394602080613d6f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610320565b97019301930191939290613cfd565b9073ffffffffffffffffffffffffffffffffffffffff60808301511682519360a060208501519401516040519586947fa32845bd00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff600487019416845260a06020850152613e35613e01825160a080880152610140870190610320565b60208301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608783030160c0880152610320565b906040810151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608682030160e08701526020808451928381520193019060005b818110613f4b5750505092613f0e602098613f00613f1c95613ef28b999660808a613ec060608e9d01516101008c019073ffffffffffffffffffffffffffffffffffffffff169052565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60898303016101208a0152610320565b908682036040880152613ce0565b908482036060860152613ce0565b916080818403910152610320565b03915afa90811561058757600091613f32575090565b610374915060203d6020116107f5576107e6818361022e565b8251805173ffffffffffffffffffffffffffffffffffffffff168652602090810151818701528c9a5060409095019490920191600101613e76565b6020818303126101855780519067ffffffffffffffff821161018557016040818303126101855760405191613fba83610212565b815167ffffffffffffffff81116101855781613fd7918401612508565b8352602082015167ffffffffffffffff811161018557613ff79201612508565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080614031855184602087015260c0860190610320565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b93929161406f612694565b5060208501918251156143a0576140a36106e66105f56106e6895173ffffffffffffffffffffffffffffffffffffffff1690565b9373ffffffffffffffffffffffffffffffffffffffff8516158015614315575b6142b2579161415860009261418b969798946141368751916141196140fc895173ffffffffffffffffffffffffffffffffffffffff1690565b94614105610281565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809581927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301613fff565b038183885af1918215610587576141ed9360009361427b575b506142196119f992613040614251935197604051978891602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810188528761022e565b604051928391602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b6020825192015192614261610281565b948552602085015260408401526060830152608082015290565b6142519193506119f9926130406142a7614219933d806000833e61429f818361022e565b810190613f86565b9593505092506141a4565b6105546142d3885173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481895afa90811561058757600091614381575b50156140c3565b61439a915060203d60201161058057610572818361022e565b3861437a565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b906143dd602092828151948592016102fd565b0190565b6144a395614494601a6020937fff0000000000000000000000000000000000000000000000000000000000000060029d9c997fffffffffffffffff00000000000000000000000000000000000000000000000060019a816143dd9f9b81869c7f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b1660198201520191828151948592016102fd565b019160f81b16815201906143ca565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b6145b9966144a3967fffff000000000000000000000000000000000000000000000000000000000000600160029c986103749f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a8498826145b99b60f81b16815261454a82518093602089850191016102fd565b019160f81b1683820152614569825160029360208295850191016102fd565b01019160f01b16838201526145888251809360206003850191016102fd565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b01906143ca565b61027f90929192602060405194826145e187945180928580880191016102fd565b83016145f5825180938580850191016102fd565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361022e565b606081019060ff8251511161491357608081019060ff825151116148e45760c081019160ff835151116148b55760e082019060ff825151116148865761010083019361ffff85515111614857576101208401936001855151116148285761014081019261ffff845151116147f9576060955180516147dd575b50815167ffffffffffffffff169060208301516146c09067ffffffffffffffff1690565b9260408101516146d79067ffffffffffffffff1690565b99519081516146e69060ff1690565b92519182516146f59060ff1690565b60a09092015161ffff16936040519c8d976020890197614714986143e1565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752614744908761022e565b519283516147529060ff1690565b92519081516147619060ff1690565b95519283516147719061ffff1690565b93825161477f9061ffff1690565b915194855161478f9061ffff1690565b94604051998a9960208b01996147a49a6144cf565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526147d4908261022e565b610374916145c0565b6147f29196506147ec9061280e565b51614a7e565b943861469c565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b1561494957565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60226145b9966144a3967fff0000000000000000000000000000000000000000000000000000000000000060029b9781956103749f9e9c978e9984917f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152614a4d82518093602089850191016102fd565b019160f81b1683820152614a6b8251809360206023850191016102fd565b01019160f81b16600182015201906143ca565b602081019060ff82515111614b9c57604081019160ff83515111614b6d57606082019160ff83515111614b3e57608081019261ffff84515111614b0d57610374936119f99251935190614ad2825160ff1690565b965190614ae0825160ff1690565b935191614aee835160ff1690565b915194614afd865161ffff1690565b946040519a8b9960208b016149cd565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105546024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91929015614c465750815115614bdf575090565b3b15614be85790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614c595750805190602001fd5b614c8f906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610363565b0390fdfea164736f6c634300081a000a",
}

var CCVProxyABI = CCVProxyMetaData.ABI

var CCVProxyBin = CCVProxyMetaData.Bin

func DeployCCVProxy(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CCVProxyStaticConfig, dynamicConfig CCVProxyDynamicConfig) (common.Address, *types.Transaction, *CCVProxy, error) {
	parsed, err := CCVProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCVProxyBin), backend, staticConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCVProxy{address: address, abi: *parsed, CCVProxyCaller: CCVProxyCaller{contract: contract}, CCVProxyTransactor: CCVProxyTransactor{contract: contract}, CCVProxyFilterer: CCVProxyFilterer{contract: contract}}, nil
}

type CCVProxy struct {
	address common.Address
	abi     abi.ABI
	CCVProxyCaller
	CCVProxyTransactor
	CCVProxyFilterer
}

type CCVProxyCaller struct {
	contract *bind.BoundContract
}

type CCVProxyTransactor struct {
	contract *bind.BoundContract
}

type CCVProxyFilterer struct {
	contract *bind.BoundContract
}

type CCVProxySession struct {
	Contract     *CCVProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCVProxyCallerSession struct {
	Contract *CCVProxyCaller
	CallOpts bind.CallOpts
}

type CCVProxyTransactorSession struct {
	Contract     *CCVProxyTransactor
	TransactOpts bind.TransactOpts
}

type CCVProxyRaw struct {
	Contract *CCVProxy
}

type CCVProxyCallerRaw struct {
	Contract *CCVProxyCaller
}

type CCVProxyTransactorRaw struct {
	Contract *CCVProxyTransactor
}

func NewCCVProxy(address common.Address, backend bind.ContractBackend) (*CCVProxy, error) {
	abi, err := abi.JSON(strings.NewReader(CCVProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCVProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCVProxy{address: address, abi: abi, CCVProxyCaller: CCVProxyCaller{contract: contract}, CCVProxyTransactor: CCVProxyTransactor{contract: contract}, CCVProxyFilterer: CCVProxyFilterer{contract: contract}}, nil
}

func NewCCVProxyCaller(address common.Address, caller bind.ContractCaller) (*CCVProxyCaller, error) {
	contract, err := bindCCVProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCVProxyCaller{contract: contract}, nil
}

func NewCCVProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CCVProxyTransactor, error) {
	contract, err := bindCCVProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCVProxyTransactor{contract: contract}, nil
}

func NewCCVProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CCVProxyFilterer, error) {
	contract, err := bindCCVProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCVProxyFilterer{contract: contract}, nil
}

func bindCCVProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCVProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCVProxy *CCVProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVProxy.Contract.CCVProxyCaller.contract.Call(opts, result, method, params...)
}

func (_CCVProxy *CCVProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.Contract.CCVProxyTransactor.contract.Transfer(opts)
}

func (_CCVProxy *CCVProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVProxy.Contract.CCVProxyTransactor.contract.Transact(opts, method, params...)
}

func (_CCVProxy *CCVProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_CCVProxy *CCVProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.Contract.contract.Transfer(opts)
}

func (_CCVProxy *CCVProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVProxy.Contract.contract.Transact(opts, method, params...)
}

func (_CCVProxy *CCVProxyCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (CCVProxyDestChainConfig, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(CCVProxyDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVProxyDestChainConfig)).(*CCVProxyDestChainConfig)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetDestChainConfig(destChainSelector uint64) (CCVProxyDestChainConfig, error) {
	return _CCVProxy.Contract.GetDestChainConfig(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCallerSession) GetDestChainConfig(destChainSelector uint64) (CCVProxyDestChainConfig, error) {
	return _CCVProxy.Contract.GetDestChainConfig(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCaller) GetDynamicConfig(opts *bind.CallOpts) (CCVProxyDynamicConfig, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CCVProxyDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVProxyDynamicConfig)).(*CCVProxyDynamicConfig)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetDynamicConfig() (CCVProxyDynamicConfig, error) {
	return _CCVProxy.Contract.GetDynamicConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) GetDynamicConfig() (CCVProxyDynamicConfig, error) {
	return _CCVProxy.Contract.GetDynamicConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _CCVProxy.Contract.GetExpectedNextSequenceNumber(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _CCVProxy.Contract.GetExpectedNextSequenceNumber(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getFee", destChainSelector, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, arg1)
}

func (_CCVProxy *CCVProxyCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, arg1)
}

func (_CCVProxy *CCVProxyCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _CCVProxy.Contract.GetPoolBySourceToken(&_CCVProxy.CallOpts, arg0, sourceToken)
}

func (_CCVProxy *CCVProxyCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _CCVProxy.Contract.GetPoolBySourceToken(&_CCVProxy.CallOpts, arg0, sourceToken)
}

func (_CCVProxy *CCVProxyCaller) GetStaticConfig(opts *bind.CallOpts) (CCVProxyStaticConfig, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CCVProxyStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVProxyStaticConfig)).(*CCVProxyStaticConfig)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetStaticConfig() (CCVProxyStaticConfig, error) {
	return _CCVProxy.Contract.GetStaticConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) GetStaticConfig() (CCVProxyStaticConfig, error) {
	return _CCVProxy.Contract.GetStaticConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _CCVProxy.Contract.GetSupportedTokens(&_CCVProxy.CallOpts, arg0)
}

func (_CCVProxy *CCVProxyCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _CCVProxy.Contract.GetSupportedTokens(&_CCVProxy.CallOpts, arg0)
}

func (_CCVProxy *CCVProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) Owner() (common.Address, error) {
	return _CCVProxy.Contract.Owner(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) Owner() (common.Address, error) {
	return _CCVProxy.Contract.Owner(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCVProxy *CCVProxySession) TypeAndVersion() (string, error) {
	return _CCVProxy.Contract.TypeAndVersion(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) TypeAndVersion() (string, error) {
	return _CCVProxy.Contract.TypeAndVersion(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "acceptOwnership")
}

func (_CCVProxy *CCVProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVProxy.Contract.AcceptOwnership(&_CCVProxy.TransactOpts)
}

func (_CCVProxy *CCVProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVProxy.Contract.AcceptOwnership(&_CCVProxy.TransactOpts)
}

func (_CCVProxy *CCVProxyTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CCVProxy *CCVProxySession) ApplyDestChainConfigUpdates(destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.Contract.ApplyDestChainConfigUpdates(&_CCVProxy.TransactOpts, destChainConfigArgs)
}

func (_CCVProxy *CCVProxyTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.Contract.ApplyDestChainConfigUpdates(&_CCVProxy.TransactOpts, destChainConfigArgs)
}

func (_CCVProxy *CCVProxyTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxySession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.ForwardFromRouter(&_CCVProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxyTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.ForwardFromRouter(&_CCVProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxyTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CCVProxy *CCVProxySession) SetDynamicConfig(dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.Contract.SetDynamicConfig(&_CCVProxy.TransactOpts, dynamicConfig)
}

func (_CCVProxy *CCVProxyTransactorSession) SetDynamicConfig(dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.Contract.SetDynamicConfig(&_CCVProxy.TransactOpts, dynamicConfig)
}

func (_CCVProxy *CCVProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_CCVProxy *CCVProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.TransferOwnership(&_CCVProxy.TransactOpts, to)
}

func (_CCVProxy *CCVProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.TransferOwnership(&_CCVProxy.TransactOpts, to)
}

func (_CCVProxy *CCVProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CCVProxy *CCVProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.WithdrawFeeTokens(&_CCVProxy.TransactOpts, feeTokens)
}

func (_CCVProxy *CCVProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.WithdrawFeeTokens(&_CCVProxy.TransactOpts, feeTokens)
}

type CCVProxyCCIPMessageSentIterator struct {
	Event *CCVProxyCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyCCIPMessageSent)
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
		it.Event = new(CCVProxyCCIPMessageSent)
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

func (it *CCVProxyCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *CCVProxyCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	MessageId         [32]byte
	EncodedMessage    []byte
	VerifierReceipts  []CCVProxyReceipt
	ExecutorReceipt   CCVProxyReceipt
	ReceiptBlobs      [][]byte
	Raw               types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*CCVProxyCCIPMessageSentIterator, error) {

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

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyCCIPMessageSentIterator{contract: _CCVProxy.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CCVProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyCCIPMessageSent)
				if err := _CCVProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseCCIPMessageSent(log types.Log) (*CCVProxyCCIPMessageSent, error) {
	event := new(CCVProxyCCIPMessageSent)
	if err := _CCVProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyConfigSetIterator struct {
	Event *CCVProxyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyConfigSet)
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
		it.Event = new(CCVProxyConfigSet)
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

func (it *CCVProxyConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyConfigSet struct {
	StaticConfig  CCVProxyStaticConfig
	DynamicConfig CCVProxyDynamicConfig
	Raw           types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CCVProxyConfigSetIterator, error) {

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCVProxyConfigSetIterator{contract: _CCVProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyConfigSet)
				if err := _CCVProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseConfigSet(log types.Log) (*CCVProxyConfigSet, error) {
	event := new(CCVProxyConfigSet)
	if err := _CCVProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyDestChainConfigSetIterator struct {
	Event *CCVProxyDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyDestChainConfigSet)
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
		it.Event = new(CCVProxyDestChainConfigSet)
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

func (it *CCVProxyDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVProxyDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	DefaultCCVs       []common.Address
	LaneMandatedCCVs  []common.Address
	DefaultExecutor   common.Address
	CcvAggregator     []byte
	Raw               types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCVProxyDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyDestChainConfigSetIterator{contract: _CCVProxy.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyDestChainConfigSet)
				if err := _CCVProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseDestChainConfigSet(log types.Log) (*CCVProxyDestChainConfigSet, error) {
	event := new(CCVProxyDestChainConfigSet)
	if err := _CCVProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyFeeTokenWithdrawnIterator struct {
	Event *CCVProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyFeeTokenWithdrawn)
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
		it.Event = new(CCVProxyFeeTokenWithdrawn)
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

func (it *CCVProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCVProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CCVProxyFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyFeeTokenWithdrawnIterator{contract: _CCVProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCVProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyFeeTokenWithdrawn)
				if err := _CCVProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCVProxyFeeTokenWithdrawn, error) {
	event := new(CCVProxyFeeTokenWithdrawn)
	if err := _CCVProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyOwnershipTransferRequestedIterator struct {
	Event *CCVProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyOwnershipTransferRequested)
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
		it.Event = new(CCVProxyOwnershipTransferRequested)
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

func (it *CCVProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCVProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyOwnershipTransferRequestedIterator{contract: _CCVProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyOwnershipTransferRequested)
				if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCVProxyOwnershipTransferRequested, error) {
	event := new(CCVProxyOwnershipTransferRequested)
	if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyOwnershipTransferredIterator struct {
	Event *CCVProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyOwnershipTransferred)
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
		it.Event = new(CCVProxyOwnershipTransferred)
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

func (it *CCVProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCVProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyOwnershipTransferredIterator{contract: _CCVProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyOwnershipTransferred)
				if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseOwnershipTransferred(log types.Log) (*CCVProxyOwnershipTransferred, error) {
	event := new(CCVProxyOwnershipTransferred)
	if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (CCVProxyCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0xdc37a122d47e708a593d43fba77d7a22899a573dfbd0cd4423c7d41a6291e0ff")
}

func (CCVProxyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (CCVProxyDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae")
}

func (CCVProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCVProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCVProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCVProxy *CCVProxy) Address() common.Address {
	return _CCVProxy.address
}

type CCVProxyInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (CCVProxyDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCVProxyDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (CCVProxyStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*CCVProxyCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CCVProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*CCVProxyCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CCVProxyConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CCVProxyConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCVProxyDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CCVProxyDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CCVProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCVProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCVProxyFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCVProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCVProxyOwnershipTransferred, error)

	Address() common.Address
}
