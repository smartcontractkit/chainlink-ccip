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
	Router                    common.Address
	MessageNumber             uint64
	AddressBytesLength        uint8
	TokenReceiverAllowed      bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
	BaseExecutionGasCost      uint32
	DefaultExecutor           common.Address
	LaneMandatedCCVs          []common.Address
	DefaultCCVs               []common.Address
	OffRamp                   []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector         uint64
	Router                    common.Address
	AddressBytesLength        uint8
	TokenReceiverAllowed      bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
	BaseExecutionGasCost      uint32
	DefaultCCVs               []common.Address
	LaneMandatedCCVs          []common.Address
	DefaultExecutor           common.Address
	OffRamp                   []byte
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
	ChainSelector         uint64
	RmnRemote             common.Address
	MaxUSDCentsPerMessage uint32
	TokenAdminRegistry    common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346103805760405161619138819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615dc290816103cf8239608051818181610a80015281816113ed0152611c9a015260a0518181816111ab0152611cc6015260c051818181611d210152612782015260e051818181611cf20152613f780152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611c1c565b611b60565b611af1565b611105565b610f76565b610f2f565b610e7e565b610e0b565b610d39565b610afd565b610ab8565b6105c6565b61037e565b6102f0565b3461016a57600060031936011261016a576080610122611c62565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101bb57604052565b61016f565b6060810190811067ffffffffffffffff8211176101bb57604052565b6080810190811067ffffffffffffffff8211176101bb57604052565b6040810190811067ffffffffffffffff8211176101bb57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101bb57604052565b604051906102476101c083610214565b565b6040519061024761016083610214565b6040519061024760a083610214565b6040519061024760c083610214565b67ffffffffffffffff81116101bb57601f01601f191660200190565b604051906102a2602083610214565b60008252565b60005b8381106102bb5750506000910152565b81810151838201526020016102ab565b90601f19601f6020936102e9815180928187528780880191016102a8565b0116010190565b3461016a57600060031936011261016a5761034f60408051906103138183610214565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102cb565b0390f35b67ffffffffffffffff81160361016a57565b359061024782610353565b908160a091031261016a5790565b3461016a57604060031936011261016a5760043561039b81610353565b60243567ffffffffffffffff811161016a576103bb903690600401610370565b6103d98267ffffffffffffffff166000526004602052604060002090565b805491906001600160a01b036103f88185165b6001600160a01b031690565b161561052b579061034f936104a1939261043e6104186080850185611d49565b6104256020870187611d49565b9050159182610512575b61043885611f19565b86612ed0565b9361048d61044a612033565b60408601906104598288611d7c565b90506104c1575b6040880161048381519260608b0193845161047d60028b01611db2565b916134e3565b9092525285611d7c565b151590506104b35760f01c90505b90613a5f565b60405190815292839250602083019150565b506001015461ffff1661049b565b5061050d6104e06104db6104d5848a611d7c565b906120af565b6120bd565b60206104ef6104d5858b611d7c565b013561050060208b015161ffff1690565b9060e08b01519289613269565b610460565b91506105216040870187611d7c565b905015159161042f565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105c25781359367ffffffffffffffff85116105bf57506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105d436610567565b906105dd6143ce565b6000915b8083106105ea57005b6105f58382846120c7565b926105ff84612107565b67ffffffffffffffff81169081158015610a74575b8015610a5e575b8015610a45575b610a0e57856108849161089e6108948361088e61069460e083019561067a61067461064d898761213e565b61066c6106626101008a95949501809a61213e565b9490923691612174565b923691612174565b9061440c565b67ffffffffffffffff166000526004602052604060002090565b9687956106da6106a660208a016120bd565b88906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61087e61084760c060408b019a6107416106f38d61211c565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61079f610750608083016121d6565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107e260016107b060a084016121d6565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108418d6107f2606084016121e0565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b01612134565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d61213e565b90600389016122eb565b8a61213e565b90600286016122eb565b6101208801906108b06103ec836120bd565b156109e4576108c161090b926120bd565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061092f610929610921848b611d49565b93905061211c565b60ff1690565b036109a057956109867f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb269261097561096b600198999a85611d49565b90600484016123e4565b5460a01c67ffffffffffffffff1690565b6109956040519283928361257e565b0390a20191906105e1565b6109aa9087611d49565b906109e06040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161238e565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a5760c08801612134565b1615610622565b5060ff610a6d6040880161211c565b161561061b565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610614565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610ad4600435610353565b610ae8602435610ae381610aa7565b61273d565b6040516001600160a01b039091168152602090f35b3461016a57610b0b36610567565b906001600160a01b03600354169160005b818110610b2557005b610b366103ec6104db838587614567565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c19576001948892600092610be9575b5081610b9d575b5050505001610b1c565b81610bcd7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610bdd9461592c565b6040519081529081906020820190565b0390a338858180610b93565b610c0b91925060203d8111610c12575b610c038183610214565b810190614577565b9038610b8c565b503d610bf9565b612731565b906020808351928381520192019060005b818110610c3c5750505090565b82516001600160a01b0316845260209384019390920191600101610c2f565b90610d369160208152610c7a6020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d20610d0a610100850151610160610120860152610180850190610c1e565b610120850151601f198583030184860152610c1e565b92015190610160601f19828503019101526102cb565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d5f81610353565b6060610140604051610d708161019e565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e08201528261010082015282610120820152015216600052600460205261034f610dcd6040600020611f19565b60405191829182610c5b565b6102479092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e2b816101c0565b828152826020820152015261034f604051610e45816101c0565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610dd9565b3461016a57600060031936011261016a576000546001600160a01b0381163303610f05577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061024782610aa7565b8015150361016a57565b359061024782610f61565b3461016a57606060031936011261016a576000604051610f95816101c0565b600435610fa181610aa7565b8152602435610faf81610f61565b6020820190815260443590610fc382610aa7565b60408301918252610fd26143ce565b6001600160a01b03835116159182156110f2575b5081156110e7575b506110bf5780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c906110aa611c62565b6110b960405192839283614586565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fee565b516001600160a01b031615915038610fe6565b3461016a57608060031936011261016a57611121600435610353565b60243567ffffffffffffffff811161016a57611141903690600401610370565b60443590611150606435610aa7565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c1957600091611ac2575b50611a855760025460a01c60ff16611a5b57611232740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61125260043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b036064351615611a31578054906112796103ec6001600160a01b03841681565b3303611a075783839161128f6080840184611d49565b949060208501956112a08787611d49565b90501590816119ee575b6112b385611f19565b6112c09390600435612ed0565b93849160a01c67ffffffffffffffff166112d9906127f1565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b52611357908b610214565b604080516064356001600160a01b031660208083019190915281529a9061137e908c610214565b6113888680611d49565b86549c9160e08e901c60ff169136906113a092612810565b9060ff166113ad91614636565b9060a089015192604089016113c2908a611d7c565b6113cc9150612889565b946113d7908a611d49565b9690976113e2610237565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261145660048801611e59565b6101008b01526101208a0152610140890152610160880152610180870152369061147f92612810565b6101a085015261148d612033565b61149a6040840184611d7c565b9050611948575b906114e96114c4859493604061151a970151606087015161047d60028701611db2565b60608601528060408601526114e360808601516001600160a01b031690565b906146c0565b60c08601526114f66128d8565b976115046040840184611d7c565b1515905061193a5760f01c90505b600435613a5f565b63ffffffff90911660608401526020860193918452116119105761153f82518661474e565b61154c6040860186611d7c565b90506117f0575b61155f81959295615128565b8085526020815191012090611578604085015151612949565b9460408101958652606060009401935b6040860151805182101561175d5760206115bb6103ec6103ec6115ae8661160296612935565b516001600160a01b031690565b6115c98460608b0151612935565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612992565b03915afa8015610c19576001600160a01b039160009161172f575b501680156116d657906000878b938783886116806116498860608f611641906120bd565b980151612935565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612add565b03925af18015610c1957816116ae916001946000916116b5575b508a51906116a88383612935565b52612935565b5001611588565b6116d0913d8091833e6116c88183610214565b8101906129f5565b3861169a565b6105636116ea6115ae8460408b0151612935565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611750915060203d8111611756575b6117488183610214565b81019061271c565b3861161d565b503d61173e565b61034f85808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d61178f8d6120bd565b925193519051906117c06040519283926001600160a01b03606435169767ffffffffffffffff600435169785612ca8565b0390a4610bcd7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b61185261183c6118066104d56040890189611d7c565b60c08601805151156118f55751905b602087015161ffff169060e088015192606435916118376004359136906128f7565b614a1b565b6101808301519061184c82612928565b52612928565b5061187460406118688451828701515190612935565b51015163ffffffff1690565b60a0611884610180840151612928565b5101515163ffffffff8216811161189c575050611553565b61056392506118b46104db6104d560408a018a611d7c565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b5061190a6119038980611d49565b3691612810565b90611815565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff16611512565b50935060019150604061195c910187611d7c565b9050036119c45761151a838688946114e96114c46119b86119866104db6104d56040880188611d7c565b60206119986104d56040890189611d7c565b01356119a9602089015161ffff1690565b9060e089015192600435613269565b929394955050506114a1565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119fd6040870187611d7c565b90501515906112aa565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105636004359067ffffffffffffffff60249216600452565b611ae4915060203d602011611aea575b611adc8183610214565b8101906127dc565b386111dc565b503d611ad2565b3461016a57602060031936011261016a5767ffffffffffffffff600435611b1781610353565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611b5b5760405167ffffffffffffffff9091168152602090f35b6121ea565b3461016a57602060031936011261016a576001600160a01b03600435611b8581610aa7565b611b8d6143ce565b16338114611bf257807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611c38600435610353565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611c72816101dc565b8281528260208201528260408201520152604051611c8f816101dc565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611de457505061024792500383610214565b84546001600160a01b0316835260019485019487945060209093019201611dcf565b90600182811c92168015611e4f575b6020831014611e2057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611e15565b9060405191826000825492611e6d84611e06565b8084529360018116908115611ed95750600114611e92575b5061024792500383610214565b90506000929192526020600020906000915b818310611ebd5750509060206102479282010138611e85565b6020919350806001915483858901015201910190918492611ea4565b602093506102479592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e85565b906120136004611f27610249565b93611f96611f8b8254611f50611f43826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611f8560e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b611fe9611fd96001830154611fba611faf8261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b611ff560028201611db2565b61010086015261200760038201611db2565b61012086015201611e59565b610140830152565b67ffffffffffffffff81116101bb5760051b60200190565b60405190612042602083610214565b6000808352366020840137565b906120598261201b565b6120666040519182610214565b828152601f19612076829461201b565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156120b85790565b612080565b35610d3681610aa7565b91908110156120b85760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610d3681610353565b60ff81160361016a57565b35610d3681612111565b63ffffffff81160361016a57565b35610d3681612126565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b9291906121808161201b565b9361218e6040519586610214565b602085838152019160051b810192831161016a57905b8282106121b057505050565b6020809183356121bf81610aa7565b8152019101906121a4565b61ffff81160361016a57565b35610d36816121ca565b35610d3681610f61565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611b5b57565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611b5b57565b908160031b9180830460081490151715611b5b57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611b5b57565b81810292918115918404141715611b5b57565b8181106122df575050565b600081556001016122d4565b9067ffffffffffffffff83116101bb576801000000000000000083116101bb57815483835580841061234f575b5090600052602060002060005b8381106123325750505050565b600190602084359461234386610aa7565b01938184015501612325565b612367908360005284602060002091820191016122d4565b38612318565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d3693818152019161236d565b9190601f81116123ae57505050565b610247926000526020600020906020601f840160051c830193106123da575b601f0160051c01906122d4565b90915081906123cd565b90929167ffffffffffffffff81116101bb5761240a816124048454611e06565b8461239f565b6000601f821160011461244a57819061243b93949560009261243f575b50506000198260011b9260031b1c19161790565b9055565b013590503880612427565b601f1982169461245f84600052602060002090565b91805b87811061249a575083600195969710612480575b505050811b019055565b60001960f88560031b161c19910135169055388080612476565b90926020600181928686013581550194019101612462565b359061024782612111565b3590610247826121ca565b359061024782612126565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106125225750505090565b9091926020806001926001600160a01b03873561253e81610aa7565b168152019401929101612515565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610d369392168152604060208201526125b4604082016125a684610365565b67ffffffffffffffff169052565b6125d36125c360208401610f56565b6001600160a01b03166060830152565b6125ec6125e2604084016124b2565b60ff166080830152565b6126046125fb60608401610f6b565b151560a0830152565b61261e612613608084016124bd565b61ffff1660c0830152565b61263861262d60a084016124bd565b61ffff1660e0830152565b61265561264760c084016124c8565b63ffffffff16610100830152565b6126eb6126be61267f61266b60e08601866124d3565b6101606101208701526101a0860191612508565b61268d6101008601866124d3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086840301610140870152612508565b926126e06126cf6101208301610f56565b6001600160a01b0316610160850152565b61014081019061254c565b916101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08286030191015261236d565b9081602091031261016a5751610d3681610aa7565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c19576001600160a01b03916000916127bf57501690565b6127d8915060203d602011611756576117488183610214565b1690565b9081602091031261016a5751610d3681610f61565b67ffffffffffffffff1667ffffffffffffffff8114611b5b5760010190565b92919261281c82610277565b9161282a6040519384610214565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101bb57604052606060a0836000815282602082015282604082015282808201528260808201520152565b906128938261201b565b6128a06040519182610214565b828152601f196128b0829461201b565b019060005b8281106128c157505050565b6020906128cc612847565b828285010152016128b5565b604051906128e5826101c0565b60606040838281528260208201520152565b919082604091031261016a5760405161290f816101f8565b6020808294803561291f81610aa7565b84520135910152565b8051156120b85760200190565b80518210156120b85760209160051b010190565b906129538261201b565b6129606040519182610214565b828152601f19612970829461201b565b019060005b82811061298157505050565b806060602080938501015201612975565b60409067ffffffffffffffff610d36949316815281602082015201906102cb565b81601f8201121561016a5780516129c981610277565b926129d76040519485610214565b8184526020828401011161016a57610d3691602080850191016102a8565b9060208282031261016a57815167ffffffffffffffff811161016a57610d3692016129b3565b9080602083519182815201916020808360051b8301019401926000915b838310612a4757505050505090565b9091929394602080612ace83601f1986600196030187528951908151815260a0612abd612aab612a99612a878887015160c08a88015260c08701906102cb565b604087015186820360408801526102cb565b606086015185820360608701526102cb565b608085015184820360808601526102cb565b9201519060a08184039101526102cb565b97019301930191939290612a38565b919390610d369593612c25612c3d9260a08652612b0760a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612c10612bf8612be0612bc8612bb0612b9a8c61026060e08a0151916101c061018082015201906102cb565b6101008801518d8203609f1901888f01526102cb565b6101208701518c8203609f19016101c08e01526102cb565b610140860151609f198c8303016101e08d01526102cb565b610160850151609f198b8303016102008c01526102cb565b610180840151609f198a8303016102208b0152612a1b565b910151609f19878303016102408801526102cb565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102cb565b9080602083519182815201916020808360051b8301019401926000915b838310612c7b57505050505090565b9091929394602080612c9983601f19866001960301875289516102cb565b97019301930191939290612c6c565b9493916001600160a01b03612ccb921686526080602087015260808601906102cb565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612d0d5750505050610d369394506060818403910152612c4f565b90919295602080612d6e83601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102cb565b980192019201909291612cef565b60405190610100820182811067ffffffffffffffff8211176101bb57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612e27575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612e819185016129b3565b926020810151612e9081612126565b92604082015167ffffffffffffffff811161016a57610d3692016129b3565b60409067ffffffffffffffff610d369593168152816020820152019161236d565b91929092612edc612d7c565b6004831015806130c2575b1561300c575090612ef7916153c5565b92612f056040850151615618565b6040840190815151159081612fe2575b50612fc4575b5060c083015151612f74575b50608082016001600160a01b03612f4582516001600160a01b031690565b1615612f5057505090565b612f6760e0610d369301516001600160a01b031690565b6001600160a01b03169052565b612f88612f846060840151151590565b1590565b15612f27577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612fd79061012084015180915251612949565b606084015238612f1b565b905080612ff1575b1538612f15565b5063ffffffff613005855163ffffffff1690565b1615612fea565b94916000906130649261302d6103ec6103ec6002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612eaf565b03915afa8015610c1957600090600090600090613096575b60a088015263ffffffff16865290505b60c0850152612f05565b5050506130b861308c913d806000833e6130b08183610214565b810190612e59565b919250829161307c565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006131186131128686612dcd565b90612df3565b1614612ee7565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a5781516131538161201b565b926131616040519485610214565b81845260208085019260051b82010192831161016a57602001905b8282106131895750505090565b60208091835161319881610aa7565b81520191019061317c565b95949060009460a09467ffffffffffffffff6131ea956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102cb565b930152565b9060028201809211611b5b57565b9060018201809211611b5b57565b6001019081600111611b5b57565b9060148201809211611b5b57565b90600c8201809211611b5b57565b91908201809211611b5b57565b6000198114611b5b5760010190565b80548210156120b85760005260206000200190600090565b9293949190600361328e8567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036132a381841661273d565b169182156134ad576040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c195760009161348e575b501561347e57613354600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016131a3565b03915afa928315610c1957600093613459575b5082511561344e5782519061338661338184548094613235565b61204f565b906000928394845b87518110156133ed576133a46115ae828a612935565b6001600160a01b038116156133e157906133db6001926133cd6133c68a613242565b9989612935565b906001600160a01b03169052565b0161338e565b509550600180966133db565b5091955091936133ff575b5050815290565b60005b82811061340f57506133f8565b8061344861343561342260019486613251565b90546001600160a01b039160031b1c1690565b6133cd61344188613242565b9789612935565b01613402565b9150610d3690611db2565b6134779193503d806000833e61346f8183610214565b81019061311f565b9138613367565b505050509250610d369150611db2565b6134a7915060203d602011611aea57611adc8183610214565b3861330f565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936134ff6134f78251865190613235565b865190613235565b9061351261350c8361204f565b92612949565b94600096875b8351891015613578578861356e61356160019361354961353f6115ae8e9f9d9e9d8b612935565b6133cd838c612935565b613567613556858c612935565b519180938491613242565b9c612935565b528b612935565b5001979695613518565b959250929350955060005b865181101561360c576135996115ae8289612935565b60006001600160a01b038216815b8881106135e0575b50509060019291156135c3575b5001613583565b6135da906133cd6135d389613242565b9888612935565b386135bc565b816135f16103ec6115ae848c612935565b146135fe576001016135a7565b5060019150819050386135af565b509390945060005b855181101561369d5761362a6115ae8288612935565b60006001600160a01b038216815b878110613671575b5050906001929115613654575b5001613614565b61366b906133cd61366488613242565b9787612935565b3861364d565b816136826103ec6115ae848b612935565b1461368f57600101613638565b506001915081905038613640565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101bb5760405260606080836000815260006020820152600060408201526000838201520152565b906136f28261201b565b6136ff6040519182610214565b828152601f1961370f829461201b565b019060005b82811061372057505050565b60209061372b6136a9565b82828501015201613714565b9081606091031261016a57805161374d816121ca565b916040602083015161375e81612126565b920151610d3681612126565b9160209082815201919060005b8181106137845750505090565b9091926040806001926001600160a01b0387356137a081610aa7565b16815260208781013590820152019401929101613777565b949391929067ffffffffffffffff168552608060208601526138116137f26137e0858061254c565b60a060808a015261012089019161236d565b6137ff602086018661254c565b90607f198984030160a08a015261236d565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761024795613899613870836060976138ba978d60c0607f19826138ac9a030191015261376a565b9161388f61387f888301610f56565b6001600160a01b031660e08d0152565b608081019061254c565b90607f198b8403016101008c015261236d565b9087820360408901526102cb565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611b5b57565b908160a091031261016a5780519160208201516138fb81612126565b91604081015161390a81612126565b916080606083015161391b816121ca565b920151610d3681610f61565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d369b9a9616885216602087015260408601521660608401521660808201528160a082015201906102cb565b9081606091031261016a57805161374d81612126565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611b5b57565b906000198201918211611b5b57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611b5b57565b91908203918211611b5b57565b919082608091031261016a578151613a1281612126565b916020810151916060604083015192015190565b8115613a30570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613a99613a94613a8f885151613a8760408c01809c611d7c565b919050613235565b6131ef565b6136e8565b9660009586955b88518051881015613cfd576103ec6103ec6115ae8a613abe94612935565b613b0b60206060880192613ad38b8551612935565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612992565b03915afa8015610c19576001600160a01b0391600091613cdf575b50168015613c8b579060608e9392613b3f8b8451612935565b5190613b5060208b015161ffff1690565b958b613b8b604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016137b8565b03915afa8015610c1957600193613c28938b8f8f95600080958197613c31575b509083929161ffff613bd385613bcc6115ae613c1c99613c229d9e51612935565b9451612935565b5191613bef613be0610259565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b16604085015216606083015260808201526116a88383612935565b506138c5565b996138c5565b96019596613aa0565b613c2297506115ae965084939291509361ffff613bd382613bcc613c6e613c1c9960603d8111613c84575b613c668183610214565b810190613737565b9c9196909c9d5050505050505090919293613bab565b503d613c5c565b61056388613c9d6115ae8c8f51612935565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613cf7915060203d8111611756576117488183610214565b38613b26565b50919a9496929395509897968a613d148187611d7c565b9050614028575b50508651613d2890613985565b99613d366020860186611d49565b91613d42915086611d7c565b9560609150019486613d53876120bd565b91613d5e938a615723565b613d688b89612935565b52613d738a88612935565b50613d7e8a88612935565b516020015163ffffffff16613d92916138c5565b90613d9d8a88612935565b516040015163ffffffff16613db1916138c5565b91613dba610259565b33815290600060208301819052604083015261ffff166060820152613ddd610293565b60808201528651613ded906139b2565b90613df88289612935565b52613e039087612935565b506002546001600160a01b031692613e1a906120bd565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c1957600096600097600094600092613feb575b506000965b8651881015613f6b57613ef4600191613ebc88613eb787612219565b613a26565b613ed56060613ecb8d8d612935565b51019182516122c1565b9052858a14613efc575b6060613eeb8b8b612935565b51015190613235565b970196613e9b565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613f2f60808c01516001600160a01b031690565b1603613f3d575b5050613edf565b613eb7613f4992612249565b613f626060613f588d8d612935565b5101918251613235565b90528b88613f36565b9796509750505050613fa77f000000000000000000000000000000000000000000000000000000000000000091613eb763ffffffff8416612249565b8411613fb35750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b929850505061401391925060803d608011614021575b61400b8183610214565b8101906139fb565b919790939290919038613e96565b503d614001565b610ae36103ec6104db6104d5614041948a989698611d7c565b926001600160a01b03600091515194169060e08801908151614061610259565b6001600160a01b038516815290826020830152826040830152826060830152608082015261408f878d612935565b5261409a868c612935565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610c19578f948c89968f96948d948f9688916143af575b506142a8575b5050505050501561414f575b6118686141406141479561413a602061186861413a97604097612935565b906138c5565b958b612935565b90388a613d1b565b50506141d59160608c6141806104db6104d56141796103ec6103ec6002546001600160a01b031690565b938b611d7c565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c195761186861414060409261413a60206118688f8b906141479b61413a9a60008060009261426a575b63ffffffff92935061425790606061421f8888612935565b51019461424c8a6142308a8a612935565b510191604061423f8b8b612935565b51019063ffffffff169052565b9063ffffffff169052565b169052975097505050509550505061411c565b50505063ffffffff6142966142579260603d6060116142a1575b61428e8183610214565b81019061396f565b909350915082614207565b503d614284565b8495985060a09697506142f260206142e86060826142df6104d56142d86104db6104d58b61432b9c9d9e9f611d7c565b998d611d7c565b013599016120bd565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613927565b03915afa8015610c19578592828c939181908294614373575b506143679060606143558888612935565b51019261424c60206142308a8a612935565b5288888f8c8138614110565b915050614367925061439d915060a03d60a0116143a8575b6143958183610214565b8101906138df565b949192919050614344565b503d61438b565b6143c8915060203d602011611aea57611adc8183610214565b3861410a565b6001600160a01b036001541633036143e257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161441a815184613235565b92831561453d5760005b848110614432575050505050565b81811015614522576144476115ae8286612935565b6001600160a01b03811680156144f857614460836131fd565b87811061447257505050600101614424565b848110156144d5576001600160a01b0361448f6115ae838a612935565b16821461449e57600101614460565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036144f36115ae6144ed88856139ee565b89612935565b61448f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6145386115ae61453284846139ee565b85612935565b614447565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156120b85760051b0190565b9081602091031261016a575190565b9160806102479294936145d68160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d369281815201906102cb565b9061461982610277565b6146266040519182610214565b828152601f196120768294610277565b90815160208210806146b6575b614685570361464f5790565b6109e0906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016145fe565b5090602081015190816146978461227b565b1c61464f57506146a68261460f565b9160200360031b1b602082015290565b5060208114614643565b918251601481029080820460141490151715611b5b576146e26146e79161320b565b613219565b906146f96146f483613227565b61460f565b90601461470583612928565b5360009260215b86518510156147375760146001916147276115ae888b612935565b60601b818701520194019361470c565b919550936020935060601b90820152828152012090565b9061475e6103ec606084016120bd565b61476f600019936040810190611d7c565b90506147e7575b61478082516139b2565b9260005b848110614792575050505050565b8082600192146147e25760606147a88287612935565b51015180156147dc576147d6906147d06147c28489612935565b51516001600160a01b031690565b8661592c565b01614784565b506147d6565b6147d6565b91506147f381516139c1565b916148016147c28484612935565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610c195760009161489a575b5061487a575b50614776565b61489490606061488a8686612935565b510151908361592c565b38614874565b6148b3915060203d602011611aea57611adc8183610214565b3861486e565b604051906148c6826101f8565b60606020838281520152565b919060408382031261016a57604051906148eb826101f8565b8193805167ffffffffffffffff811161016a578261490a9183016129b3565b835260208101519167ffffffffffffffff831161016a5760209261492e92016129b3565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610d3692016148d2565b9060806001600160a01b0381614978855160a0865260a08601906102cb565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d36928181520190614959565b919060408382031261016a57825167ffffffffffffffff811161016a576020916149e89185016148d2565b92015190565b61ffff614a07610d369593606084526060840190614959565b9316602082015260408184039101526102cb565b909193929594614a29612847565b5060208201805115614ec357614a4c610ae36103ec85516001600160a01b031690565b946001600160a01b038616916040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180614ab360048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c1957600091614ea4575b5015614e5b57614b3f88999a825192614ade6148b9565b5051614b2a614af489516001600160a01b031690565b926040614aff610259565b9e8f908152614b1b8d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c1957600091614e3c575b5015614d445750906000929183614beb9899604051998a95869485937fb1c71c65000000000000000000000000000000000000000000000000000000008552600485016149ee565b03925af18015610c1957600094600091614cfc575b50614cd1614c3a9596614c74614c48614c66956115ae6020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b610214565b6040519586918683019190916001600160a01b036020820193169052565b03601f198101865285610214565b614cb0610929614ca6614cb68951614cb0610929614ca68c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b90614636565b9767ffffffffffffffff166000526004602052604060002090565b93015193614cdd610268565b958652602086015260408501526060840152608083015260a082015290565b614c3a9550602091506115ae96614c74614c48614c6695614d32614cd1953d806000833e614d2a8183610214565b8101906149bd565b9b909b96505095505050969550614c00565b9793929061ffff16614e125751614de857614d936000939184926040519586809481937f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016149ac565b03925af1908115610c1957614c3a95614c74614c48614cd1936115ae602096614c6698600091614dc5575b5099614c1c565b614de291503d806000833e614dda8183610214565b810190614933565b38614dbe565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614e55915060203d602011611aea57611adc8183610214565b38614ba3565b610563614e6f86516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614ebd915060203d602011611aea57611adc8183610214565b38614ac7565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614fbf947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614fdb602092828151948592016102a8565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b16815261502782518093602089850191016102a8565b019160f81b16838201526150458251809360206002850191016102a8565b01019160f81b16838201526150648251809360206002850191016102a8565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d369f9e9c9860f81b1681526150d982518093602089850191016102a8565b019160f01b16838201526150f78251809360206003850191016102a8565b01019160f01b168382015261511582518093602089850191016102a8565b01019160f01b1660028201520190614fc8565b60e081019060ff825151116153af57610100810160ff815151116153995761012082019260ff8451511161538357610140830160ff8151511161536d5761016084019061ffff825151116153575761018085019360018551511161533f576101a086019361ffff855151116153295760609551805161530d575b50865167ffffffffffffffff16602088015167ffffffffffffffff169060408901516151d59067ffffffffffffffff1690565b9860608101516151e89063ffffffff1690565b9060808101516151fb9063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661521e97614eed565b03601f19810188526152309088610214565b5190815161523e9060ff1690565b9051805160ff1698519081516152549060ff1690565b906040519a8b95602087019561526996614fdf565b03601f198101875261527b9087610214565b519182516152899060ff1690565b9151805161ffff1694805161529f9061ffff1690565b92519283516152af9061ffff1690565b9260405197889760208901976152c49861506a565b03601f19810182526152d69082610214565b604051928392602084016152e991614fc8565b6152f291614fc8565b6152fb91614fc8565b03601f1981018252610d369082610214565b61532291965061531c90612928565b51615b23565b94386151a2565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b6000526105636024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b906153ce612d7c565b91601182106155e45780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216036155715750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61544d8161204f565b6040860190815261545d82612949565b906060870191825260005b83811061552557505050506154d083836154c66154ba6154b06154a96154966154da98876154e49c9b615c4a565b6001600160a01b0390911660808d015290565b8585615d20565b9291903691612810565b60a08a01528383615d88565b9491903691612810565b60c0880152615d20565b9391903691612810565b60e084015281036154f3575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b8060019161556a61555461554d6155406155649a8d8d615c4a565b91906133cd868a51612935565b8b8b615d20565b9391889a919a51949a3691612810565b92612935565b5201615468565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b80519060005b82811061562a57505050565b60018101808211611b5b575b838110615646575060010161561e565b6001600160a01b036156588385612935565b511661566a6103ec6115ae8487612935565b1461567757600101615636565b6105636156876115ae8486612935565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9081602091031261016a5751610d36816121ca565b9361570e60809461ffff6001600160a01b039567ffffffffffffffff61571c969b9a9b16895216602088015260a0604088015260a0870190610c1e565b9085820360608701526102cb565b9416910152565b9291909261572f6136a9565b5061574e8167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff169160808501928351615771906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615791916138c5565b9661579d90608d613235565b9460a08701958651516157af91613235565b9160ff16916157bd83612291565b6157c691613235565b916157d2906067613235565b6157db916122c1565b6157e491613235565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b038716036158715750505061ffff9250615863906158566000935b519561584961583a610259565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b6158876103ec602094976001600160a01b031690565b9060406158988583015161ffff1690565b910151926158d8875198604051998a96879586957ff2388958000000000000000000000000000000000000000000000000000000008752600487016156d1565b03915afa908115610c19576158566158639261ffff956000916158fd575b509361582d565b61591f915060203d602011615925575b6159178183610214565b8101906156bc565b386158f6565b503d61590d565b91602091600091604051906001600160a01b03858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152615980606482610214565b519082855af115612731576000513d6159df57506001600160a01b0381163b155b6159a85750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b600114156159a1565b966002615af797615ac46022610d369f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615ac49f82615ac49c615acb9c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615a7582518093602089850191016102a8565b019160f81b1683820152615a938251809360206023850191016102a8565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614fc8565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615c3457604081019160ff83515111615c1e57606082019160ff83515111615c0857608081019260ff84515111615bf25760a0820161ffff81515111615bdc57610d3694615bce9351945191615b85835160ff1690565b975191615b93835160ff1690565b945190615ba1825160ff1690565b905193615baf855160ff1690565b935196615bbe885161ffff1690565b966040519c8d9b60208d016159e8565b03601f198101835282610214565b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b635a102da160e11b600052602660045260246000fd5b929190926001820191848311615cee5781013560001a828115615ce3575060148103615cb6578201938411615c8257013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b91906002820191818311615cee578381013560f01c0160020192818411615d5457918391615d4d93612ddb565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b91906001820191818311615cee578381013560001a0160010192818411615d5457918391615d4d93612ddb56fea164736f6c634300081a000a",
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
	Sender            common.Address
	MessageId         [32]byte
	FeeToken          common.Address
	EncodedMessage    []byte
	Receipts          []OnRampReceipt
	VerifierBlobs     [][]byte
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, senderRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, senderRule, messageIdRule)
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
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeeTokenWithdrawnIterator{contract: _OnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
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
	return common.HexToHash("0xb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26")
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

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	Address() common.Address
}
