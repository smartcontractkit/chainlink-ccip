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

type InternalReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint64
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

var CCVProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"verifierReceipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executorReceipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ccvAggregator\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVInUserInput\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enumMessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc5760405161528038819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051614f35908161034b8239608051818181610b47015281816114e00152611cfa015260a0518181816104c80152611d33015260c051818181611d6f015261207a0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635cb80c5d146100f357806366c3a5c7146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611c57565b611b63565b6111ab565b610fd7565b610f35565b610ee3565b610dfa565b610d2e565b610c96565b61089a565b61075c565b610671565b610418565b610377565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576060610140611cda565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60c0810190811067ffffffffffffffff8211176101d557604052565b61018a565b6060810190811067ffffffffffffffff8211176101d557604052565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b6040519061027f6101608361022e565b565b6040519061027f60a08361022e565b6040519061027f60e08361022e565b6040519061027f60408361022e565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102f760208361022e565b60008252565b60005b8381106103105750506000910152565b8181015183820152602001610300565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361035c815180928187528780880191016102fd565b0116010190565b906020610374928181520190610320565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103f460408051906103b8818361022e565b601282527f43435650726f787920312e372e302d6465760000000000000000000000000000602083015251918291602083526020830190610320565b0390f35b67ffffffffffffffff81160361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557600435610453816103f8565b60243567ffffffffffffffff81116101855761047390369060040161040a565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156105e457600091610624575b506105e95761058e9160209161055861053f61053f60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401611eb1565b03915afa80156105e4576103f4916000916105b5575b506040519081529081906020820190565b6105d7915060203d6020116105dd575b6105cf818361022e565b810190611db8565b386105a4565b503d6105c5565b611dac565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b610646915060203d60201161064c575b61063e818361022e565b810190611d97565b3861050e565b503d610634565b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576106ab6004356103f8565b60206106c16024356106bc81610653565b61201b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101855760043567ffffffffffffffff81116101855760040160009280601f830112156107585781359367ffffffffffffffff851161075557506020808301928560051b010111610185579190565b80fd5b8380fd5b346101855761076a366106df565b9061078a60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b81811061079757005b6107ad61053f6107a8838587612143565b612158565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156105e457600194889160009361087a575b5082610822575b505050500161078e565b61082d918391613115565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610818565b61089391935060203d81116105dd576105cf818361022e565b9138610811565b34610185576108a8366106df565b906108b1613215565b6000915b8083106108be57005b6108c9838284612162565b926108d3846121a2565b9367ffffffffffffffff85169081158015610b3b575b610b035761093493949561094e604083019161090583856121ac565b979061092e606087019961092661091c8c8a6121ac565b9490923691612218565b923691612218565b9061329a565b67ffffffffffffffff166000526004602052604060002090565b60208301906109a061095f83612158565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6109b76109ad84866121ac565b90600384016122b4565b6109ce6109c488866121ac565b90600284016122b4565b608084016109de61053f82612158565b15610ad9576001977f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae95610ab983610abf610ace96610a65610a22610ab198612158565b8f83019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610aa8610aa1610a9b60a0880193610a8a610a80868b612336565b906004840161241f565b5460a01c67ffffffffffffffff1690565b9a612158565b9a866121ac565b979096866121ac565b949093612158565b94612336565b959094604051998a998a612598565b0390a20191906108b5565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5067ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682146108e9565b906020808351928381520192019060005b818110610b8c5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610b7f565b90610374916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015273ffffffffffffffffffffffffffffffffffffffff604083015116606082015260a0610c63610c30606085015160c0608086015260e0850190610b6e565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610b6e565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610320565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610cda816103f8565b606060a0604051610cea816101b9565b600081526000602082015260006040820152828082015282608082015201521660005260046020526103f4610d226040600020612727565b60405191829182610bb8565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610d65611cbb565b50604051610d72816101da565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103f482606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff81163303610eb9577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff600435610f79816103f8565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610fbd5760405167ffffffffffffffff9091168152602090f35b61226e565b359061027f82610653565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051611014816101da565b60043561102081610653565b815260243561102e81610fcd565b602082019081526044359061104282610653565b60408301918252611051613215565b73ffffffffffffffffffffffffffffffffffffffff8351161591821561118b575b508115611180575b506111585780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611143611cda565b61115260405192839283613443565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b90505115153861107a565b5173ffffffffffffffffffffffffffffffffffffffff1615915038611072565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576004356111e6816103f8565b60243567ffffffffffffffff81116101855761120690369060040161040a565b60443560643561121581610653565b60025460a01c60ff16611b3957611266740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6112848467ffffffffffffffff166000526004602052604060002090565b9373ffffffffffffffffffffffffffffffffffffffff821615611b0f578454926112c461053f73ffffffffffffffffffffffffffffffffffffffff861681565b3303611ae55761134e6112ec6112dd6080880188612336565b906112e78a612727565b6137e4565b92600061131461053f61053f60025473ffffffffffffffffffffffffffffffffffffffff1690565b60a08601519060405180809681947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610363565b03915afa9283156105e4578792600094611ac8575b50835115611a73575b611374612833565b9685518a60208801998a5160408a019384516113909060ff1690565b9261139a94613c2b565b60ff16909252908952865260a01c67ffffffffffffffff166113bb9061284f565b89547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178a556040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260148152909261144460348361022e565b606087810151604051918a901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602083015260148252909b9061ffff1661148e60348e61022e565b6114988780612336565b806114a29261286e565b604089019e8f6114b2908b6128f2565b6114bc9150612971565b99602081016114ca91612336565b9490956114d561026f565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529867ffffffffffffffff891660208b015260408a019b61152c908d9067ffffffffffffffff169052565b60608a015260040161153d90612667565b608089015261ffff1660a088015260c0870152369061155b926129de565b60e08501526115686102e8565b61010085015261012084019687523690611581926129de565b61014083015286515189515161159691612a23565b61159f90612a30565b9860005b8851805182101561163d57908b61162b826115c081600196612acd565b519260206115e2855173ffffffffffffffffffffffffffffffffffffffff1690565b94015161160c6115f0610281565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b6000602086015260006040860152600060608601526080850152612acd565b52611636818d612acd565b50016115a3565b50508a8c9a98999a60005b835180518210156116ee57906116e78c8e6116e06116d98561166c81600199612acd565b5193602061168e865173ffffffffffffffffffffffffffffffffffffffff1690565b9501516116b861169c610281565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b60006020870152600060408701526000606087015260808601525151612a23565b8093612acd565b528d612acd565b5001611648565b50508b8b9697999a989a608088015161171a9073ffffffffffffffffffffffffffffffffffffffff1690565b9b60a0890151611728610281565b73ffffffffffffffffffffffffffffffffffffffff909e168e52600060208f0181905260408f0181905260608f015260808e0152866117673687612b95565b90611772918b614013565b5061177d84866128f2565b90506119e8575b5050505060606117949101612158565b9361179e846148b8565b97885160208a0120936117be6117b984515186515190612a23565b612c53565b9960005b84518051821015611897579060008a6117dc838e95612acd565b518a838d611842602061180961053f61053f885173ffffffffffffffffffffffffffffffffffffffff1690565b950151604051998a97889687957fc527f20000000000000000000000000000000000000000000000000000000000875260048701612d85565b03925af19182156105e457818e60019461186f93600091611876575b506118698383612acd565b52612acd565b50016117c2565b611891913d8091833e611889818361022e565b8101906127fc565b3861185e565b5050919399979a9890929860005b8b51805182101561194c579060008a8f93836118c091612acd565b518a838d6118ed602061180961053f61053f885173ffffffffffffffffffffffffffffffffffffffff1690565b03925af180156105e4578b61192091838f600196600093611927575b50611919916116d9915151612a23565b528c612acd565b50016118a5565b91611943611919936116d993953d8091833e611889818361022e565b93915091611909565b50506103f496507fdc37a122d47e708a593d43fba77d7a22899a573dfbd0cd4423c7d41a6291e0ff916119a88a67ffffffffffffffff936119968a995167ffffffffffffffff1690565b93856040519687961699169785613060565b0390a46119d87fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b60016119f785879596976128f2565b905003611a4957611a2e606094611a409389611a29611a24611a1c6117949b8a6128f2565b369291612c4a565b612afc565b6142f9565b905190611a3a82612ac0565b52612ac0565b50918b80611784565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b9250611ac2611a828380612336565b9190611a9660405193849260208401612822565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261022e565b9261136c565b611ade9194503d806000833e611889818361022e565b9238611363565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff600435611bb381610653565b611bbb613215565b16338114611c2d57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557611c916004356103f8565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611cc8826101da565b60006040838281528260208201520152565b611ce2611cbb565b50604051611cef816101da565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610185575161037481610fcd565b6040513d6000823e3d90fd5b90816020910312610185575190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160209082815201919060005b818110611e705750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735611e9981610653565b16815260208781013590820152019401929101611e63565b919067ffffffffffffffff16825260406020830152611f24611ee7611ed68380611dc7565b60a0604087015260e0860191611e17565b611ef46020840184611dc7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016060870152611e17565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561018557820160208135910167ffffffffffffffff8211610185578160061b360381136101855784611feb92611fb4927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866103749903016080870152611e56565b92611fe1611fc460608301610fc2565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b6080810190611dc7565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152611e17565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156105e4576000906120c5575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d60201161210c575b816120df6020938361022e565b810103126101855773ffffffffffffffffffffffffffffffffffffffff905161210781610653565b6120aa565b3d91506120d2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156121535760051b0190565b612114565b3561037481610653565b91908110156121535760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610185570190565b35610374816103f8565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160051b3603831361018557565b67ffffffffffffffff81116101d55760051b60200190565b92919061222481612200565b93612232604051958661022e565b602085838152019160051b810192831161018557905b82821061225457505050565b60208091833561226381610653565b815201910190612248565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106122a8575050565b6000815560010161229d565b9067ffffffffffffffff83116101d5576801000000000000000083116101d5578154838355808410612318575b5090600052602060002060005b8381106122fb5750505050565b600190602084359461230c86610653565b019381840155016122ee565b6123309083600052846020600020918201910161229d565b386122e1565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b90600182811c921680156123d0575b60208310146123a157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612396565b9190601f81116123e957505050565b61027f926000526020600020906020601f840160051c83019310612415575b601f0160051c019061229d565b9091508190612408565b90929167ffffffffffffffff81116101d5576124458161243f8454612387565b846123da565b6000601f82116001146124a3578190612494939495600092612498575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612462565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946124d684600052602060002090565b91805b87811061252f5750836001959697106124f7575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806124ed565b909260206001819286860135815501940191016124d9565b9160209082815201919060005b8181106125615750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561258a81610653565b168152019401929101612554565b94916125f29373ffffffffffffffffffffffffffffffffffffffff95866125e49367ffffffffffffffff6103749e9c9d9b96168a5216602089015260c0604089015260c0880191612547565b918583036060870152612547565b9416608082015260a0818503910152611e17565b906040519182815491828252602082019060005260206000209260005b81811061263857505061027f9250038361022e565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612623565b906040519182600082549261267b84612387565b80845293600181169081156126e757506001146126a0575b5061027f9250038361022e565b90506000929192526020600020906000915b8183106126cb57505090602061027f9282010138612693565b60209193508060019154838589010152019101909184926126b2565b6020935061027f9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612693565b90604051612734816101b9565b60a06127b56004839567ffffffffffffffff80825473ffffffffffffffffffffffffffffffffffffffff81168852861c1616602086015273ffffffffffffffffffffffffffffffffffffffff8060018301541616604086015261279960028201612606565b60608601526127aa60038201612606565b608086015201612667565b910152565b81601f820112156101855780516127d0816102ae565b926127de604051948561022e565b818452602082840101116101855761037491602080850191016102fd565b9060208282031261018557815167ffffffffffffffff81116101855761037492016127ba565b916020610374938181520191611e17565b6040519061284260208361022e565b6000808352366020840137565b67ffffffffffffffff1667ffffffffffffffff8114610fbd5760010190565b9092919283600c1161018557831161018557600c01917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40190565b906004116101855790600490565b909291928360041161018557831161018557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b60405190612953826101f6565b60606080836000815282602082015282604082015282808201520152565b9061297b82612200565b612988604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129b68294612200565b019060005b8281106129c757505050565b6020906129d2612946565b828285010152016129bb565b9291926129ea826102ae565b916129f8604051938461022e565b829481845281830111610185578281602093846000960137010152565b9060018201809211610fbd57565b91908201809211610fbd57565b90612a3a82612200565b612a47604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612a758294612200565b019060005b828110612a8657505050565b602090604051612a95816101f6565b6000815260008382015260006040820152600060608201526060608082015282828501015201612a7a565b8051156121535760200190565b80518210156121535760209160051b010190565b9080601f8301121561018557816020610374933591016129de565b919082604091031261018557604051612b1481610212565b60208082948035612b2481610653565b84520135910152565b81601f82011215610185578035612b4381612200565b92612b51604051948561022e565b81845260208085019260061b8401019281841161018557602001915b838310612b7b575050505090565b6020604091612b8a8486612afc565b815201920191612b6d565b91909160a08184031261018557612baa610281565b92813567ffffffffffffffff81116101855781612bc8918401612ae1565b8452602082013567ffffffffffffffff81116101855781612bea918401612ae1565b6020850152604082013567ffffffffffffffff81116101855781612c0f918401612b2d565b6040850152612c2060608301610fc2565b6060850152608082013567ffffffffffffffff811161018557612c439201612ae1565b6080830152565b90156121535790565b90612c5d82612200565b612c6a604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612c988294612200565b019060005b828110612ca957505050565b806060602080938501015201612c9d565b9080602083519182815201916020808360051b8301019401926000915b838310612ce657505050505090565b9091929394602080612d76837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080612d65612d53612d418786015160a08987015260a0860190610320565b60408601518582036040870152610320565b60608501518482036060860152610320565b920151906080818403910152610320565b97019301930191939290612cd7565b9193906103749593612f51612f769260a08652612daf60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152610140612f1e612ee8612eb2612e7d612e3b612e068c61020060608a0151916101606101008201520190610320565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101208f0152610320565b60a087015161ffff16868d015260c08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101608e0152610320565b60e08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101808d0152610320565b6101008501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016101a08c0152610320565b6101208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016101c08b0152612cba565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e0880152610320565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610320565b9060a060806103749373ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff602082015116602085015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610320565b9080602083519182815201916020808360051b8301019401926000915b83831061301557505050505090565b9091929394602080613051837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610320565b97019301930191939290613006565b929493919061307790608085526080850190610320565b948386036020850152815180875260208701906020808260051b8a01019401916000905b8282106130cc5750505050610374949550906130be918482036040860152612f88565b916060818403910152612fe9565b90919294602080613107837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08e600196030186528951612f88565b97019201920190929161309b565b9073ffffffffffffffffffffffffffffffffffffffff6131e79392604051938260208601947fa9059cbb00000000000000000000000000000000000000000000000000000000865216602486015260448501526044845261317760648561022e565b1660008060409384519561318b868861022e565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d1561320c573d6131d86131cf826102ae565b9451948561022e565b83523d6000602085013e614e60565b8051806131f2575050565b816020806132079361027f9501019101611d97565b614bd7565b60609250614e60565b73ffffffffffffffffffffffffffffffffffffffff60015416330361323657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610fbd57565b91908203918211610fbd57565b8051916132a8815184612a23565b9283156134195760005b8481106132c0575050505050565b818110156133fe576132ef6132d58286612acd565b5173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811680156133d45761331583612a15565b878110613327575050506001016132b2565b848110156133a45773ffffffffffffffffffffffffffffffffffffffff6133516132d5838a612acd565b16821461336057600101613315565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6133cf6132d56133c9888561328d565b89612acd565b613351565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6134146132d561340e848461328d565b85612acd565b6132ef565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91606061027f9294936134908160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6040519060e0820182811067ffffffffffffffff8211176101d557604052606060c08382815282602082015260006040820152600083820152600060808201528260a08201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613543575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9061357f82612200565b61358c604051918261022e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06135ba8294612200565b019060005b8281106135cb57505050565b6020906040516135da81610212565b60008152606083820152828285010152016135bf565b9080601f830112156101855781359161360883612200565b92613616604051948561022e565b80845260208085019160051b830101918383116101855760208101915b83831061364257505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610185576040519061368f82610212565b602083013561369d81610653565b825260408301359167ffffffffffffffff8311610185576136c688602080969581960101612ae1565b83820152815201920191613633565b359060ff8216820361018557565b359061ffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160e08183031261018557613721610290565b91813567ffffffffffffffff8111610185578161373f9184016135f0565b8352602082013567ffffffffffffffff811161018557816137619184016135f0565b6020840152613772604083016136d5565b6040840152613783606083016136e3565b606084015261379460808301610fc2565b608084015260a082013567ffffffffffffffff811161018557816137b9918401612ae1565b60a084015260c082013567ffffffffffffffff8111610185576137dc9201612ae1565b60c082015290565b91906137ee6134c5565b600483101580613b5b575b15613ad85750816138159261380d926128b7565b8101906136f2565b906020820180515180613a83575b5082515161383382515182612a23565b60005b81811061395b575050506138509083515190515190612a23565b156138d6575b6080820173ffffffffffffffffffffffffffffffffffffffff61388d825173ffffffffffffffffffffffffffffffffffffffff1690565b161561389857505090565b6138bc604061037493015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b60808101926138e6845151613575565b835260005b84518051821015613952579061394b8161390a6132d582600196612acd565b61393161391561029f565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b6139396102e8565b60208201528751906118698383612acd565b50016138eb565b50509250613856565b82811015613a6c5761398c613971828851612acd565b515173ffffffffffffffffffffffffffffffffffffffff1690565b61399582612a15565b8381106139a6575050600101613836565b84811015613a3a5773ffffffffffffffffffffffffffffffffffffffff6139d1613971838b51612acd565b1673ffffffffffffffffffffffffffffffffffffffff8316146139f657600101613995565b7fd757e5e80000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff613a676139718851613a61898661328d565b90612acd565b6139d1565b613a7e6139718551613a61868561328d565b61398c565b604084015160ff1690811080159190613acc575b50613aa25738613823565b7fb273a0e40000000000000000000000000000000000000000000000000000000060005260046000fd5b60ff9150161538613a97565b919392613ae63686846129de565b60a08401526080810193613afb855151613575565b845260005b85518051821015613b4e5790613b4781613b1f6132d582600196612acd565b613b2a61391561029f565b613b35368c8a6129de565b60208201528851906118698383612acd565b5001613b00565b5050945091925050613856565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613bb1613bab86866128a9565b9061350f565b16146137f9565b80548210156121535760005260206000200190600090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610fbd5760010190565b60ff168015610fbd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b600290959395018054613c3f835182612a23565b92613c4984613575565b9260009485905b808210613d0157505050505081613c68575050929190565b9194909293613c80613c7b858851612a23565b613575565b9260005b858110613cd757505060005b8651811015613ccd5780613cc6613ca96001938a612acd565b51613cb48389612a23565b90613cbf8289612acd565b5286612acd565b5001613c90565b5091945092909150565b80613ce460019284612acd565b51613cef8288612acd565b52613cfa8187612acd565b5001613c84565b93999298919790969195929492938a881015613f5957613d44613d24898b613bb8565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b985b6000805b878110613f1a575b50613f075760009a8b5b8551811015613ef457613d7561053f6139718389612acd565b73ffffffffffffffffffffffffffffffffffffffff8d1614613d9957600101613d5c565b5093989195979a5093959a919860015b15613dbf575b506001905b019093929192613c50565b999792613dfd8b613de1613dda8b9d98959d9996999b613bd0565b9a8a612acd565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60005b8c51811015613ee2578c8c73ffffffffffffffffffffffffffffffffffffffff613e3061053f6139718686612acd565b911614613e405750600101613e00565b613e9492959a9c50613e5b826020929d95989d999699612acd565b5101516020613e72613e6c8c613260565b8b612acd565b5101528c61191982613e8d613e878451613260565b84612acd565b5192612acd565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8b51018b5260ff8a16613ece575b6001905b90613daf565b98613eda600191613bfd565b999050613ec4565b50929799509297600190949194613ec8565b5093989195979a9b92999094969b613da9565b93959a9198509395989196600190613db4565b8b73ffffffffffffffffffffffffffffffffffffffff80613f3e613971858c612acd565b9216911614613f4f57600101613d4a565b5050600138613d52565b613f6f6132d5613f698d8b61328d565b8c612acd565b98613d46565b9080602083519182815201916020808360051b8301019401926000915b838310613fa157505050505090565b9091929394602080614004837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610320565b97019301930191939290613f92565b9073ffffffffffffffffffffffffffffffffffffffff60808301511682519360a060208501519401516040519586947fa32845bd00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff600487019416845260a060208501526140ca614096825160a080880152610140870190610320565b60208301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608783030160c0880152610320565b906040810151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608682030160e08701526020808451928381520193019060005b8181106141e057505050926141a36020986141956141b1956141878b999660808a61415560608e9d01516101008c019073ffffffffffffffffffffffffffffffffffffffff169052565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60898303016101208a0152610320565b908682036040880152613f75565b908482036060860152613f75565b916080818403910152610320565b03915afa9081156105e4576000916141c7575090565b610374915060203d6020116105dd576105cf818361022e565b8251805173ffffffffffffffffffffffffffffffffffffffff168652602090810151818701528c9a506040909501949092019160010161410b565b6020818303126101855780519067ffffffffffffffff82116101855701604081830312610185576040519161424f83610212565b815167ffffffffffffffff8111610185578161426c9184016127ba565b8352602082015167ffffffffffffffff81116101855761428c92016127ba565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff60806142c6855184602087015260c0860190610320565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b939291614304612946565b5060208501918251156146355761433861053f6106bc61053f895173ffffffffffffffffffffffffffffffffffffffff1690565b9373ffffffffffffffffffffffffffffffffffffffff85161580156145aa575b61454757916143ed600092614420969798946143cb8751916143ae614391895173ffffffffffffffffffffffffffffffffffffffff1690565b9461439a610281565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809581927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614294565b038183885af19182156105e45761448293600093614510575b506144ae611a96926132d56144e6935197604051978891602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810188528761022e565b604051928391602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b60208251920151926144f6610281565b948552602085015260408401526060830152608082015290565b6144e6919350611a96926132d561453c6144ae933d806000833e614534818361022e565b81019061421b565b959350509250614439565b610620614568885173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481895afa9081156105e457600091614616575b5015614358565b61462f915060203d60201161064c5761063e818361022e565b3861460f565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b90614672602092828151948592016102fd565b0190565b61473895614729601a6020937fff0000000000000000000000000000000000000000000000000000000000000060029d9c997fffffffffffffffff00000000000000000000000000000000000000000000000060019a816146729f9b81869c7f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b1660198201520191828151948592016102fd565b019160f81b168152019061465f565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b61484e96614738967fffff000000000000000000000000000000000000000000000000000000000000600160029c986103749f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a84988261484e9b60f81b1681526147df82518093602089850191016102fd565b019160f81b16838201526147fe825160029360208295850191016102fd565b01019160f01b168382015261481d8251809360206003850191016102fd565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b019061465f565b61027f909291926020604051948261487687945180928580880191016102fd565b830161488a825180938580850191016102fd565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361022e565b606081019060ff82515111614ba857608081019060ff82515111614b795760c081019160ff83515111614b4a5760e082019060ff82515111614b1b5761010083019361ffff85515111614aec57610120840193600185515111614abd5761014081019261ffff84515111614a8e57606095518051614a72575b50815167ffffffffffffffff169060208301516149559067ffffffffffffffff1690565b92604081015161496c9067ffffffffffffffff1690565b995190815161497b9060ff1690565b925191825161498a9060ff1690565b60a09092015161ffff16936040519c8d9760208901976149a998614676565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810187526149d9908761022e565b519283516149e79060ff1690565b92519081516149f69060ff1690565b9551928351614a069061ffff1690565b938251614a149061ffff1690565b9151948551614a249061ffff1690565b94604051998a9960208b0199614a399a614764565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252614a69908261022e565b61037491614855565b614a87919650614a8190612ac0565b51614d13565b9438614931565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b15614bde57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b602261484e96614738967fff0000000000000000000000000000000000000000000000000000000000000060029b9781956103749f9e9c978e9984917f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152614ce282518093602089850191016102fd565b019160f81b1683820152614d008251809360206023850191016102fd565b01019160f81b166001820152019061465f565b602081019060ff82515111614e3157604081019160ff83515111614e0257606082019160ff83515111614dd357608081019261ffff84515111614da25761037493611a969251935190614d67825160ff1690565b965190614d75825160ff1690565b935191614d83835160ff1690565b915194614d92865161ffff1690565b946040519a8b9960208b01614c62565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526106206024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91929015614edb5750815115614e74575090565b3b15614e7d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015614eee5750805190602001fd5b614f24906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610363565b0390fdfea164736f6c634300081a000a",
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

func (_CCVProxy *CCVProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, message)
}

func (_CCVProxy *CCVProxyCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, message)
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
	VerifierReceipts  []InternalReceipt
	ExecutorReceipt   InternalReceipt
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

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
