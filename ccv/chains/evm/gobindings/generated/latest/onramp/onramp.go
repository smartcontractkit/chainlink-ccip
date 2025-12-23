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
	ChainSelector         uint64
	RmnRemote             common.Address
	MaxUSDCentsPerMessage uint32
	TokenAdminRegistry    common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"networkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461038057604051615f6038819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615b9190816103cf8239608051818181610a1d015281816113860152611be9015260a0518181816111460152611c15015260c051818181611c700152612619015260e051818181611c410152613dab0152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780633fc89498146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611b6b565b611aba565b611a4b565b6110a1565b610f1d565b610ed6565b610e30565b610dbd565b610cf3565b610a9a565b610a55565b6105b6565b610379565b6102eb565b3461016a57600060031936011261016a576080610122611bb1565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b610140810190811067ffffffffffffffff8211176101a257604052565b61016f565b6060810190811067ffffffffffffffff8211176101a257604052565b6080810190811067ffffffffffffffff8211176101a257604052565b6040810190811067ffffffffffffffff8211176101a257604052565b90601f601f19910116810190811067ffffffffffffffff8211176101a257604052565b6040519061022e6101c0836101fb565b565b6040519061022e610140836101fb565b6040519061022e60a0836101fb565b6040519061022e60c0836101fb565b67ffffffffffffffff81116101a257601f01601f191660200190565b604051906102896020836101fb565b60008252565b60005b8381106102a25750506000910152565b8181015183820152602001610292565b90601f19601f6020936102d08151809281875287808801910161028f565b0116010190565b9060206102e89281815201906102b2565b90565b3461016a57600060031936011261016a5761034a604080519061030e81836101fb565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102b2565b0390f35b67ffffffffffffffff81160361016a57565b359061022e8261034e565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103968161034e565b60243567ffffffffffffffff811161016a576103b690369060040161036b565b6103d48267ffffffffffffffff166000526004602052604060002090565b9081546001600160a01b036103fe6103f2836001600160a01b031690565b6001600160a01b031690565b161561051b579061034a9361049761049d94936104476104216080860186611c98565b61042e6020880188611c98565b9050159182610502575b61044189611e4f565b87612d67565b94610450611f54565b6040860161045e8188611ccb565b90506104af575b506104896040880191825160608a0194610483600287519201611d01565b9161332f565b9092525260e81c61ffff1690565b90613892565b60405190815292839250602083019150565b6104fc915060206104de6104cc6104d76104d26104cc868d611ccb565b90611fb7565b611fc5565b938a611ccb565b01356104ef60208a015161ffff1690565b9060e08a015192886130e7565b38610465565b91506105116040880188611ccb565b9050151591610438565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105b25781359367ffffffffffffffff85116105af57506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105c436610557565b906105cd6141cf565b6000915b8083106105da57005b6105e5838284611fcf565b926105ef8461200f565b67ffffffffffffffff81169081158015610a11575b80156109fb575b80156109e2575b6109ab57856106629161083d61083361067c60c08501936106338587612046565b979061065c60e089019961065461064a8c8c612046565b949092369161207c565b92369161207c565b9061420d565b67ffffffffffffffff166000526004602052604060002090565b946106b561068c60208701611fc5565b87906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61082d6108236107e360a0604089019861071f6106d18b612024565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61077e61072e606083016120de565b8c547fff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e89190911b7effff000000000000000000000000000000000000000000000000000000000016178c55565b6107dd61078d608083016120e8565b8c547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f81b7fff0000000000000000000000000000000000000000000000000000000000000016178c55565b0161203c565b9561081d60018a0197889063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b8d612046565b90600389016121da565b8a612046565b90600286016121da565b61010088019061084f6103f283611fc5565b15610981576108606108a892611fc5565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6101208701906108cc6108c66108be848b611c98565b939050612024565b60ff1690565b0361093d57956109237fe24d61e75f1236506b973f38fe122980e9c3d9e27f5c6721a85b921f70c512c592610912610908600198999a85611c98565b90600484016122d3565b5460a01c67ffffffffffffffff1690565b6109326040519283928361246d565b0390a20191906105d1565b6109479087611c98565b9061097d6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161227d565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff6109f460a0880161203c565b1615610612565b5060ff610a0a60408801612024565b161561060b565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610604565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610a7160043561034e565b610a85602435610a8081610a44565b6125d4565b6040516001600160a01b039091168152602090f35b3461016a57610aa836610557565b906001600160a01b0360035416918215610bc15760005b818110610ac857005b610ad96103f26104d283858761433e565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610bbc576001948892600092610b8c575b5081610b40575b5050505001610abf565b81610b707f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610b809461560e565b6040519081529081906020820190565b0390a338858180610b36565b610bae91925060203d8111610bb5575b610ba681836101fb565b81019061434e565b9038610b2f565b503d610b9c565b6125c8565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610c095750505090565b82516001600160a01b0316845260209384019390920191600101610bfc565b906102e89160208152610c476020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015161ffff1660808201526080820151151560a082015260a082015163ffffffff1660c082015260c08201516001600160a01b031660e0820152610120610cdd610cc760e0850151610140610100860152610160850190610beb565b610100850151601f198583030184860152610beb565b92015190610140601f19828503019101526102b2565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d198161034e565b6060610120604051610d2a81610185565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152015216600052600460205261034a610d7f6040600020611e4f565b60405191829182610c28565b61022e9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610ddd816101a7565b828152826020820152015261034a604051610df7816101a7565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610d8b565b3461016a57600060031936011261016a576000546001600160a01b0381163303610eac5773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022e82610a44565b8015150361016a57565b359061022e82610f08565b3461016a57606060031936011261016a576000604051610f3c816101a7565b600435610f4881610a44565b8152602435610f5681610f08565b6020820190815260443590610f6a82610a44565b60408301918252610f796141cf565b6001600160a01b038351161591821561108e575b508115611083575b5061105b5780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff00000000000000000000000000000000000000001617905560408201516003805473ffffffffffffffffffffffffffffffffffffffff1916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c90611046611bb1565b6110556040519283928361435d565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610f95565b516001600160a01b031615915038610f8d565b3461016a57608060031936011261016a576110bd60043561034e565b60243567ffffffffffffffff811161016a576110dd90369060040161036b565b6044356110eb606435610a44565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610bbc57600091611a1c575b506119df5760025460a01c60ff166119b5576111cd740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6111ed60043567ffffffffffffffff166000526004602052604060002090565b916001600160a01b0360643516801561198b578354936112166103f26001600160a01b03871681565b3303611961576112296080840184611c98565b9060208501916112398387611c98565b9050159081611948575b61124c85611e4f565b6112599390600435612d67565b9560a01c67ffffffffffffffff1661127090612688565b82547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178355865163ffffffff169160208801516112d39061ffff1690565b604080513060208083019190915281529491906112f090866101fb565b604080516064356001600160a01b03166020808301919091528152959061131790876101fb565b6113218980611c98565b8854979160e089901c60ff16913690611339926126a7565b9060ff16611346916143fc565b9060a08d01519260408c0161135b908d611ccb565b6113659150612720565b96611370908d611c98565b95909661137b61021e565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529967ffffffffffffffff6004351660208c015267ffffffffffffffff1660408b0152600060608b015263ffffffff1660808a015261ffff1660a0890152600060c089015260e08801526113fd60048a01611d8f565b61010088015261012087015261014086015261016085015261018084019283523690611428926126a7565b6101a0830152611436611f54565b976114446040880188611ccb565b90506118b1575b908161148861146f60408b9695019b8c51606085019961048360028c519201611d01565b8852808c5260808301516001600160a01b031690614486565b60c08401526114b081896114a861149d61276f565b9860e81c61ffff1690565b600435613892565b63ffffffff9091166060860152602087019591865211611887576114d5845189614514565b6114e26040890189611ccb565b905061176d575b50506114f781969396614e71565b808752602081519101209161150d8951516127e0565b9560408801968752606060009501945b8a5180518210156116e657602061154d6103f26103f261154086611591966127cc565b516001600160a01b031690565b611558848b516127cc565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612829565b03915afa8015610bbc576001600160a01b03916000916116b8575b5016801561166257906000888c9388838961160c6115d5888f6115ce90611fc5565b97516127cc565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612974565b03925af18015610bbc578161163a91600194600091611641575b508b519061163483836127cc565b526127cc565b500161151d565b61165c913d8091833e61165481836101fb565b81019061288c565b38611626565b610553611673611540848f516127cc565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b6116d9915060203d81116116df575b6116d181836101fb565b8101906125b3565b386115ac565b503d6116c7565b5061034a8580848c7fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d8961173d61171d8f611fc5565b94519151925160405193849367ffffffffffffffff600435169785612b3f565b0390a4610b707fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6117c26117d4916117846104cc60408c018c611ccb565b60c082018051519192911561186c5751915b60e06117a7602084015161ffff1690565b92015192606435916117bd60043591369061278e565b6147af565b8251906117ce826127bf565b526127bf565b5060a06118016117fa60406117ee8d8851905151906127cc565b51015163ffffffff1690565b92516127bf565b5101515163ffffffff82168111156114e957610553925061182b6104d26104cc60408b018b611ccb565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b5061188161187a8c80611c98565b36916126a7565b91611796565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b97509060016118c36040880188611ccb565b90500361191e5786916119146118e26104d26104cc60408b018b611ccb565b60206118f46104cc60408c018c611ccb565b013561190560208d015161ffff1690565b9060e08d0151926004356130e7565b989091925061144b565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119576040870187611ccb565b9050151590611243565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105536004359067ffffffffffffffff60249216600452565b611a3e915060203d602011611a44575b611a3681836101fb565b810190612673565b38611177565b503d611a2c565b3461016a57602060031936011261016a5767ffffffffffffffff600435611a718161034e565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611ab55760405167ffffffffffffffff9091168152602090f35b6120f2565b3461016a57602060031936011261016a576001600160a01b03600435611adf81610a44565b611ae76141cf565b16338114611b41578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611b8760043561034e565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611bc1816101c3565b8281528260208201528260408201520152604051611bde816101c3565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611d3357505061022e925003836101fb565b84546001600160a01b0316835260019485019487945060209093019201611d1e565b90600182811c92168015611d85575b6020831014611d6f57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611d64565b9060405191826000825492611da384611d55565b8084529360018116908115611e0f5750600114611dc8575b5061022e925003836101fb565b90506000929192526020600020906000915b818310611df357505090602061022e9282010138611dbb565b6020919350806001915483858901015201910190918492611dda565b6020935061022e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611dbb565b90611f346004611e5d610230565b93611ec4611ebb8254611e86611e79826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c16604089015261ffff60e882901c16606089015260f81c90565b15156080870152565b611f0b611efb6001830154611eec611edf8263ffffffff1690565b63ffffffff1660a08a0152565b60201c6001600160a01b031690565b6001600160a01b031660c0870152565b611f1760028201611d01565b60e0860152611f2860038201611d01565b61010086015201611d8f565b610120830152565b67ffffffffffffffff81116101a25760051b60200190565b60405190611f636020836101fb565b6000808352366020840137565b90611f7a82611f3c565b611f8760405191826101fb565b828152601f19611f978294611f3c565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b9015611fc05790565b611fa1565b356102e881610a44565b9190811015611fc05760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec18136030182121561016a570190565b356102e88161034e565b60ff81160361016a57565b356102e881612019565b63ffffffff81160361016a57565b356102e88161202e565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061208881611f3c565b9361209660405195866101fb565b602085838152019160051b810192831161016a57905b8282106120b857505050565b6020809183356120c781610a44565b8152019101906120ac565b61ffff81160361016a57565b356102e8816120d2565b356102e881610f08565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611ab557565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611ab557565b908160031b9180830460081490151715611ab557565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611ab557565b81810292918115918404141715611ab557565b8181106121ce575050565b600081556001016121c3565b9067ffffffffffffffff83116101a2576801000000000000000083116101a257815483835580841061223e575b5090600052602060002060005b8381106122215750505050565b600190602084359461223286610a44565b01938184015501612214565b612256908360005284602060002091820191016121c3565b38612207565b601f8260209493601f19938186528686013760008582860101520116010190565b9160206102e893818152019161225c565b9190601f811161229d57505050565b61022e926000526020600020906020601f840160051c830193106122c9575b601f0160051c01906121c3565b90915081906122bc565b90929167ffffffffffffffff81116101a2576122f9816122f38454611d55565b8461228e565b6000601f821160011461233957819061232a93949560009261232e575b50506000198260011b9260031b1c19161790565b9055565b013590503880612316565b601f1982169461234e84600052602060002090565b91805b87811061238957508360019596971061236f575b505050811b019055565b60001960f88560031b161c19910135169055388080612365565b90926020600181928686013581550194019101612351565b359061022e82612019565b359061022e826120d2565b359061022e8261202e565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124115750505090565b9091926020806001926001600160a01b03873561242d81610a44565b168152019401929101612404565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff6102e89392168152604060208201526124a36040820161249584610360565b67ffffffffffffffff169052565b6124c26124b260208401610efd565b6001600160a01b03166060830152565b6124db6124d1604084016123a1565b60ff166080830152565b6124f56124ea606084016123ac565b61ffff1660a0830152565b61250d61250460808401610f12565b151560c0830152565b61252961251c60a084016123b7565b63ffffffff1660e0830152565b6125a061257361255361253f60c08601866123c2565b6101406101008701526101808601916123f7565b61256060e08601866123c2565b90603f19868403016101208701526123f7565b926125956125846101008301610efd565b6001600160a01b0316610140850152565b61012081019061243b565b91610160603f198286030191015261225c565b9081602091031261016a57516102e881610a44565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610bbc576001600160a01b039160009161265657501690565b61266f915060203d6020116116df576116d181836101fb565b1690565b9081602091031261016a57516102e881610f08565b67ffffffffffffffff1667ffffffffffffffff8114611ab55760010190565b9291926126b38261025e565b916126c160405193846101fb565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101a257604052606060a0836000815282602082015282604082015282808201528260808201520152565b9061272a82611f3c565b61273760405191826101fb565b828152601f196127478294611f3c565b019060005b82811061275857505050565b6020906127636126de565b8282850101520161274c565b6040519061277c826101a7565b60606040838281528260208201520152565b919082604091031261016a576040516127a6816101df565b602080829480356127b681610a44565b84520135910152565b805115611fc05760200190565b8051821015611fc05760209160051b010190565b906127ea82611f3c565b6127f760405191826101fb565b828152601f196128078294611f3c565b019060005b82811061281857505050565b80606060208093850101520161280c565b60409067ffffffffffffffff6102e8949316815281602082015201906102b2565b81601f8201121561016a5780516128608161025e565b9261286e60405194856101fb565b8184526020828401011161016a576102e8916020808501910161028f565b9060208282031261016a57815167ffffffffffffffff811161016a576102e8920161284a565b9080602083519182815201916020808360051b8301019401926000915b8383106128de57505050505090565b909192939460208061296583601f1986600196030187528951908151815260a061295461294261293061291e8887015160c08a88015260c08701906102b2565b604087015186820360408801526102b2565b606086015185820360608701526102b2565b608085015184820360808601526102b2565b9201519060a08184039101526102b2565b970193019301919392906128cf565b9193906102e89593612abc612ad49260a0865261299e60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612aa7612a8f612a77612a5f612a47612a318c61026060e08a0151916101c061018082015201906102b2565b6101008801518d8203609f1901888f01526102b2565b6101208701518c8203609f19016101c08e01526102b2565b610140860151609f198c8303016101e08d01526102b2565b610160850151609f198b8303016102008c01526102b2565b610180840151609f198a8303016102208b01526128b2565b910151609f19878303016102408801526102b2565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102b2565b9080602083519182815201916020808360051b8301019401926000915b838310612b1257505050505090565b9091929394602080612b3083601f19866001960301875289516102b2565b97019301930191939290612b03565b9493916001600160a01b03612b62921686526080602087015260808601906102b2565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612ba457505050506102e89394506060818403910152612ae6565b90919295602080612c0583601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102b2565b980192019201909291612b86565b60405190610100820182811067ffffffffffffffff8211176101a257604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612cbe575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612d1891850161284a565b926020810151612d278161202e565b92604082015167ffffffffffffffff811161016a576102e8920161284a565b60409067ffffffffffffffff6102e89593168152816020820152019161225c565b91929092612d73612c13565b600483101580612f59575b15612ea3575090612d8e9161510e565b92612d9c60408501516152fd565b6040840190815151159081612e79575b50612e5b575b5060c083015151612e0b575b50608082016001600160a01b03612ddc82516001600160a01b031690565b1615612de757505090565b612dfe60c06102e89301516001600160a01b031690565b6001600160a01b03169052565b612e1f612e1b6080840151151590565b1590565b15612dbe577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612e6e90610100840151809152516127e0565b606084015238612db2565b905080612e88575b1538612dac565b5063ffffffff612e9c855163ffffffff1690565b1615612e81565b9491600090612efb92612ec46103f26103f26002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612d46565b03915afa8015610bbc57600090600090600090612f2d575b60a088015263ffffffff16865290505b60c0850152612d9c565b505050612f4f612f23913d806000833e612f4781836101fb565b810190612cf0565b9192508291612f13565b5063302326cb60e01b7fffffffff00000000000000000000000000000000000000000000000000000000612f96612f908686612c64565b90612c8a565b1614612d7e565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a578151612fd181611f3c565b92612fdf60405194856101fb565b81845260208085019260051b82010192831161016a57602001905b8282106130075750505090565b60208091835161301681610a44565b815201910190612ffa565b95949060009460a09467ffffffffffffffff613068956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102b2565b930152565b9060028201809211611ab557565b9060018201809211611ab557565b6001019081600111611ab557565b9060148201809211611ab557565b90600c8201809211611ab557565b91908201809211611ab557565b6000198114611ab55760010190565b8054821015611fc05760005260206000200190600090565b9293949190600361310c8567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036131218184166125d4565b169182156132f9576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bbc576000916132da575b50156132ca576131a0600095969798604051998a96879586957f89720a6200000000000000000000000000000000000000000000000000000000875260048701613021565b03915afa928315610bbc576000936132a5575b5082511561329a578251906131d26131cd845480946130b3565b611f70565b906000928394845b8751811015613239576131f0611540828a6127cc565b6001600160a01b0381161561322d57906132276001926132196132128a6130c0565b99896127cc565b906001600160a01b03169052565b016131da565b50955060018096613227565b50919550919361324b575b5050815290565b60005b82811061325b5750613244565b8061329461328161326e600194866130cf565b90546001600160a01b039160031b1c1690565b61321961328d886130c0565b97896127cc565b0161324e565b91506102e890611d01565b6132c39193503d806000833e6132bb81836101fb565b810190612f9d565b91386131b3565b5050505092506102e89150611d01565b6132f3915060203d602011611a4457611a3681836101fb565b3861315b565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b9391929361334b61334382518651906130b3565b8651906130b3565b9061335e61335883611f70565b926127e0565b94600096875b83518910156133c457886133ba6133ad60019361339561338b6115408e9f9d9e9d8b6127cc565b613219838c6127cc565b6133b36133a2858c6127cc565b5191809384916130c0565b9c6127cc565b528b6127cc565b5001979695613364565b959250929350955060005b8651811015613458576133e561154082896127cc565b60006001600160a01b038216815b88811061342c575b505090600192911561340f575b50016133cf565b6134269061321961341f896130c0565b98886127cc565b38613408565b8161343d6103f2611540848c6127cc565b1461344a576001016133f3565b5060019150819050386133fb565b509390945060005b85518110156134e95761347661154082886127cc565b60006001600160a01b038216815b8781106134bd575b50509060019291156134a0575b5001613460565b6134b7906132196134b0886130c0565b97876127cc565b38613499565b816134ce6103f2611540848b6127cc565b146134db57600101613484565b50600191508190503861348c565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101a25760405260606080836000815260006020820152600060408201526000838201520152565b9061353e82611f3c565b61354b60405191826101fb565b828152601f1961355b8294611f3c565b019060005b82811061356c57505050565b6020906135776134f5565b82828501015201613560565b9081606091031261016a578051613599816120d2565b91604060208301516135aa8161202e565b9201516102e88161202e565b9160209082815201919060005b8181106135d05750505090565b9091926040806001926001600160a01b0387356135ec81610a44565b168152602087810135908201520194019291016135c3565b949391929067ffffffffffffffff1685526080602086015261365d61363e61362c858061243b565b60a060808a015261012089019161225c565b61364b602086018661243b565b90607f198984030160a08a015261225c565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022e956136e56136bc83606097613706978d60c0607f19826136f89a03019101526135b6565b916136db6136cb888301610efd565b6001600160a01b031660e08d0152565b608081019061243b565b90607f198b8403016101008c015261225c565b9087820360408901526102b2565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611ab557565b908160a091031261016a5780519160208201516137478161202e565b9160408101516137568161202e565b9160806060830151613767816120d2565b9201516102e881610f08565b9260c0946001600160a01b039167ffffffffffffffff61ffff95846102e89b9a9616885216602087015260408601521660608401521660808201528160a082015201906102b2565b9081606091031261016a5780516135998161202e565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611ab557565b906000198201918211611ab557565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611ab557565b91908203918211611ab557565b919082608091031261016a57815161385e8161202e565b916020810151916060604083015192015190565b811561387c570490565b634e487b7160e01b600052601260045260246000fd5b93829460009060009560408101946138cc6138c76138c28851516138ba60408c01809c611ccb565b9190506130b3565b61306d565b613534565b9660009586955b88518051881015613b30576103f26103f26115408a6138f1946127cc565b61393e602060608801926139068b85516127cc565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612829565b03915afa8015610bbc576001600160a01b0391600091613b12575b50168015613abe579060608e93926139728b84516127cc565b519061398360208b015161ffff1690565b958b6139be604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613604565b03915afa8015610bbc57600193613a5b938b8f8f95600080958197613a64575b509083929161ffff613a06856139ff611540613a4f99613a559d9e516127cc565b94516127cc565b5191613a22613a13610240565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261163483836127cc565b50613711565b99613711565b960195966138d3565b613a559750611540965084939291509361ffff613a06826139ff613aa1613a4f9960603d8111613ab7575b613a9981836101fb565b810190613583565b9c9196909c9d50505050505050909192936139de565b503d613a8f565b61055388613ad06115408c8f516127cc565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613b2a915060203d81116116df576116d181836101fb565b38613959565b50919a9496929395509897968a613b478187611ccb565b9050613e5b575b50508651613b5b906137d1565b99613b696020860186611c98565b91613b75915086611ccb565b9560609150019486613b8687611fc5565b91613b91938a615408565b613b9b8b896127cc565b52613ba68a886127cc565b50613bb18a886127cc565b516020015163ffffffff16613bc591613711565b90613bd08a886127cc565b516040015163ffffffff16613be491613711565b91613bed610240565b33815290600060208301819052604083015261ffff166060820152613c1061027a565b60808201528651613c20906137fe565b90613c2b82896127cc565b52613c3690876127cc565b506002546001600160a01b031692613c4d90611fc5565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610bbc57600096600097600094600092613e1e575b506000965b8651881015613d9e57613d27600191613cef88613cea87612108565b613872565b613d086060613cfe8d8d6127cc565b51019182516121b0565b9052858a14613d2f575b6060613d1e8b8b6127cc565b510151906130b3565b970196613cce565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613d6260808c01516001600160a01b031690565b1603613d70575b5050613d12565b613cea613d7c92612138565b613d956060613d8b8d8d6127cc565b51019182516130b3565b90528b88613d69565b9796509750505050613dda7f000000000000000000000000000000000000000000000000000000000000000091613cea63ffffffff8416612138565b8411613de65750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613e4691925060803d608011613e54575b613e3e81836101fb565b810190613847565b919790939290919038613cc9565b503d613e34565b610a806103f26104d26104cc613e74948a989698611ccb565b926001600160a01b03600091515194169060e08801908151613e94610240565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613ec2878d6127cc565b52613ecd868c6127cc565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610bbc578f948c89968f96948d948f9688916141b0575b506140a9575b50505050505015613f50575b6117ee613f41613f4895613f3b60206117ee613f3b976040976127cc565b90613711565b958b6127cc565b90388a613b4e565b5050613fd69160608c613f816104d26104cc613f7a6103f26103f26002546001600160a01b031690565b938b611ccb565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610bbc576117ee613f41604092613f3b60206117ee8f8b90613f489b613f3b9a60008060009261406b575b63ffffffff92935061405890606061402088886127cc565b51019461404d8a6140318a8a6127cc565b51019160406140408b8b6127cc565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050613f1d565b50505063ffffffff6140976140589260603d6060116140a2575b61408f81836101fb565b8101906137bb565b909350915082614008565b503d614085565b8495985060a09697506140f360206140e96060826140e06104cc6140d96104d26104cc8b61412c9c9d9e9f611ccb565b998d611ccb565b01359901611fc5565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613773565b03915afa8015610bbc578592828c939181908294614174575b5061416890606061415688886127cc565b51019261404d60206140318a8a6127cc565b5288888f8c8138613f11565b915050614168925061419e915060a03d60a0116141a9575b61419681836101fb565b81019061372b565b949192919050614145565b503d61418c565b6141c9915060203d602011611a4457611a3681836101fb565b38613f0b565b6001600160a01b036001541633036141e357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161421b8151846130b3565b9283156143145760005b848110614233575050505050565b818110156142f95761424861154082866127cc565b6001600160a01b0381168015610bc1576142618361307b565b87811061427357505050600101614225565b848110156142d6576001600160a01b03614290611540838a6127cc565b16821461429f57600101614261565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036142f46115406142ee888561383a565b896127cc565b614290565b61430f611540614309848461383a565b856127cc565b614248565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b9190811015611fc05760051b0190565b9081602091031261016a575190565b91608061022e9294936143ad8160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906143df8261025e565b6143ec60405191826101fb565b828152601f19611f97829461025e565b908151602082108061447c575b61444b57036144155790565b61097d906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016102d7565b50906020810151908161445d8461216a565b1c614415575061446c826143d5565b9160200360031b1b602082015290565b5060208114614409565b918251601481029080820460141490151715611ab5576144a86144ad91613089565b613097565b906144bf6144ba836130a5565b6143d5565b9060146144cb836127bf565b5360009260215b86518510156144fd5760146001916144ed611540888b6127cc565b60601b81870152019401936144d2565b919550936020935060601b90820152828152012090565b906145246103f260608401611fc5565b614535600019936040810190611ccb565b90506145ad575b61454682516137fe565b9260005b848110614558575050505050565b8082600192146145a857606061456e82876127cc565b51015180156145a25761459c9061459661458884896127cc565b51516001600160a01b031690565b8661560e565b0161454a565b5061459c565b61459c565b91506145b9815161380d565b916145c761458884846127cc565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610bbc5760009161462e575b5061460e575b5061453c565b61462890606061461e86866127cc565b510151908361560e565b38614608565b614647915060203d602011611a4457611a3681836101fb565b38614602565b6040519061465a826101df565b60606020838281520152565b919060408382031261016a576040519061467f826101df565b8193805167ffffffffffffffff811161016a578261469e91830161284a565b835260208101519167ffffffffffffffff831161016a576020926146c2920161284a565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a576102e89201614666565b9060806001600160a01b038161470c855160a0865260a08601906102b2565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206102e89281815201906146ed565b919060408382031261016a57825167ffffffffffffffff811161016a5760209161477c918501614666565b92015190565b61ffff61479b6102e895936060845260608401906146ed565b9316602082015260408184039101526102b2565b9091939295946147bd6126de565b5060208201805115614c0c576147e0610a806103f285516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061482e60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610bbc57600091614bed575b5015614ba4576148ba88999a82519261485961464d565b50516148a561486f89516001600160a01b031690565b92604061487a610240565b9e8f9081526148968d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bbc57600091614b85575b5015614a8d57509060009291836149349899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614782565b03925af18015610bbc57600094600091614a45575b50614a1a61498395966149bd6149916149af956115406020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101fb565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101fb565b6149f96108c66149ef6149ff89516149f96108c66149ef8c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906143fc565b9767ffffffffffffffff166000526004602052604060002090565b93015193614a2661024f565b958652602086015260408501526060840152608083015260a082015290565b614983955060209150611540966149bd6149916149af95614a7b614a1a953d806000833e614a7381836101fb565b810190614751565b9b909b96505095505050969550614949565b9793929061ffff16614b5b5751614b3157614adc6000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614740565b03925af1908115610bbc57614983956149bd614991614a1a936115406020966149af98600091614b0e575b5099614965565b614b2b91503d806000833e614b2381836101fb565b8101906146c7565b38614b07565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614b9e915060203d602011611a4457611a3681836101fb565b386148ec565b610553614bb886516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614c06915060203d602011611a4457611a3681836101fb565b38614842565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614d08947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614d246020928281519485920161028f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152614d70825180936020898501910161028f565b019160f81b1683820152614d8e82518093602060028501910161028f565b01019160f81b1683820152614dad82518093602060028501910161028f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006102e89f9e9c9860f81b168152614e22825180936020898501910161028f565b019160f01b1683820152614e4082518093602060038501910161028f565b01019160f01b1683820152614e5e825180936020898501910161028f565b01019160f01b1660028201520190614d11565b60e081019060ff825151116150f857610100810160ff815151116150e25761012082019260ff845151116150cc57610140830160ff815151116150b65761016084019061ffff825151116150a057610180850193600185515111615088576101a086019361ffff8551511161507257606095518051615056575b50865167ffffffffffffffff16602088015167ffffffffffffffff16906040890151614f1e9067ffffffffffffffff1690565b986060810151614f319063ffffffff1690565b906080810151614f449063ffffffff1690565b60a082015161ffff169160c00151926040519c8d966020880196614f6797614c36565b03601f1981018852614f7990886101fb565b51908151614f879060ff1690565b9051805160ff169851908151614f9d9060ff1690565b906040519a8b956020870195614fb296614d28565b03601f1981018752614fc490876101fb565b51918251614fd29060ff1690565b9151805161ffff16948051614fe89061ffff1690565b9251928351614ff89061ffff1690565b92604051978897602089019761500d98614db3565b03601f198101825261501f90826101fb565b6040519283926020840161503291614d11565b61503b91614d11565b61504491614d11565b03601f19810182526102e890826101fb565b61506b919650615065906127bf565b5161583c565b9438614eeb565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b6000526105536024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b90615117612c13565b91601182106152e257803563302326cb60e01b7fffffffff000000000000000000000000000000000000000000000000000000008216036152885750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61517d81611f70565b6040860190815261518d826127e0565b906060870191825260005b83811061523c575050505061520083836151f66151ea6151e06151d96151c661520a98876152149c9b615963565b6001600160a01b0390911660808d015290565b8585615a07565b92919036916126a7565b60a08a01528383615a56565b94919036916126a7565b60c0880152615a07565b93919036916126a7565b60e08401528103615223575090565b63d9437f9d60e01b600052600360045260245260446000fd5b8060019161528161526b61526461525761527b9a8d8d615963565b9190613219868a516127cc565b8b8b615a07565b9391889a919a51949a36916126a7565b926127cc565b5201615198565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263302326cb60e01b6004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b82811061530f57505050565b60018101808211611ab5575b83811061532b5750600101615303565b6001600160a01b0361533d83856127cc565b511661534f6103f261154084876127cc565b1461535c5760010161531b565b61055361536c61154084866127cc565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9081602091031261016a57516102e8816120d2565b936153f360809461ffff6001600160a01b039567ffffffffffffffff615401969b9a9b16895216602088015260a0604088015260a0870190610beb565b9085820360608701526102b2565b9416910152565b929190926154146134f5565b506154338167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff169160808501928351615456906001600160a01b031690565b60019098015463ffffffff16865163ffffffff1661547391613711565b9661547f90608d6130b3565b9460a0870195865151615491916130b3565b9160ff169161549f83612180565b6154a8916130b3565b916154b49060676130b3565b6154bd916121b0565b6154c6916130b3565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b038716036155535750505061ffff9250615545906155386000935b519561552b61551c610240565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b6155696103f2602094976001600160a01b031690565b90604061557a8583015161ffff1690565b910151926155ba875198604051998a96879586957ff2388958000000000000000000000000000000000000000000000000000000008752600487016153b6565b03915afa908115610bbc576155386155459261ffff956000916155df575b509361550f565b615601915060203d602011615607575b6155f981836101fb565b8101906153a1565b386155d8565b503d6155ef565b906001600160a01b036156d39392604051938260208601947fa9059cbb0000000000000000000000000000000000000000000000000000000086521660248601526044850152604484526156636064856101fb565b1660008060409384519561567786886101fb565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d156156f8573d6156c46156bb8261025e565b945194856101fb565b83523d6000602085013e615af4565b8051806156de575050565b816020806156f39361022e9501019101612673565b615a83565b60609250615af4565b966002615810976157dd60226102e89f9e9c9799600199859f9b7fff00000000000000000000000000000000000000000000000000000000000000906157dd9f826157dd9c6157e49c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b16602182015261578e825180936020898501910161028f565b019160f81b16838201526157ac82518093602060238501910161028f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614d11565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff8251511161594d57604081019160ff8351511161593757606082019160ff8351511161592157608081019260ff8451511161590b5760a0820161ffff815151116158f5576102e8946158e7935194519161589e835160ff1690565b9751916158ac835160ff1690565b9451906158ba825160ff1690565b9051936158c8855160ff1690565b9351966158d7885161ffff1690565b966040519c8d9b60208d01615701565b03601f1981018352826101fb565b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b635a102da160e11b600052602660045260246000fd5b9291909260018201918483116159ee5781013560001a8281156159e35750601481036159b657820193841161599b57013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b919060028201918183116159ee578381013560f01c0160020192818411615a3b57918391615a3493612c72565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b919060018201918183116159ee578381013560001a0160010192818411615a3b57918391615a3493612c72565b15615a8a57565b608460405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91929015615b555750815115615b08575090565b3b15615b115790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615b685750805190602001fd5b61097d9060405191829162461bcd60e51b8352600483016102d756fea164736f6c634300081a000a",
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
