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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461038057604051615f0938819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615b3a90816103cf8239608051818181610a5c015281816113e30152611c85015260a0518181816111a10152611cb1015260c051818181611d0c01526126e6015260e051818181611cdd0152613e780152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611c07565b611b56565b611ae7565b6110fb565b610f77565b610f30565b610e8a565b610e17565b610d45565b610ad9565b610a94565b6105ad565b610365565b6102d7565b3461016a57600060031936011261016a576080610122611c4d565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101a257604052565b61016f565b6060810190811067ffffffffffffffff8211176101a257604052565b6080810190811067ffffffffffffffff8211176101a257604052565b6040810190811067ffffffffffffffff8211176101a257604052565b90601f601f19910116810190811067ffffffffffffffff8211176101a257604052565b6040519061022e6101c0836101fb565b565b6040519061022e610160836101fb565b6040519061022e60a0836101fb565b6040519061022e60c0836101fb565b67ffffffffffffffff81116101a257601f01601f191660200190565b604051906102896020836101fb565b60008252565b60005b8381106102a25750506000910152565b8181015183820152602001610292565b90601f19601f6020936102d08151809281875287808801910161028f565b0116010190565b3461016a57600060031936011261016a5761033660408051906102fa81836101fb565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102b2565b0390f35b67ffffffffffffffff81160361016a57565b359061022e8261033a565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103828161033a565b60243567ffffffffffffffff811161016a576103a2903690600401610357565b6103c08267ffffffffffffffff166000526004602052604060002090565b805491906001600160a01b036103df8185165b6001600160a01b031690565b161561051257906103369361048893926104256103ff6080850185611d34565b61040c6020870187611d34565b90501591826104f9575b61041f85611eeb565b86612e34565b93610474610431612005565b60408601906104408288611d67565b90506104a8575b6040880161046a81519260608b0193845161046460028b01611d9d565b916133fc565b9092525285611d67565b1515905061049a5760f01c90505b9061395f565b60405190815292839250602083019150565b506001015461ffff16610482565b506104f46104c76104c26104bc848a611d67565b90612068565b612076565b60206104d66104bc858b611d67565b01356104e760208b015161ffff1690565b9060e08b015192896131b4565b610447565b91506105086040870187611d67565b9050151591610416565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105a95781359367ffffffffffffffff85116105a657506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105bb3661054e565b906105c461429c565b6000915b8083106105d157005b6105dc838284612080565b926105e6846120c0565b67ffffffffffffffff81169081158015610a50575b8015610a3a575b8015610a21575b6109ea57856108609161087a6108708361086a61067b60e083019561066161065b61063489876120f7565b6106536106496101008a95949501809a6120f7565b949092369161212d565b92369161212d565b906142da565b67ffffffffffffffff166000526004602052604060002090565b9687956106b661068d60208a01612076565b88906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61085a61082360c060408b019a61071d6106cf8d6120d5565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61077b61072c6080830161218f565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107be600161078c60a0840161218f565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b61081d8d6107ce60608401612199565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016120ed565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6120f7565b906003890161228b565b8a6120f7565b906002860161228b565b61012088019061088c6103d383612076565b156109c05761089d6108e792612076565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061090b6109056108fd848b611d34565b9390506120d5565b60ff1690565b0361097c57956109627f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb2692610951610947600198999a85611d34565b9060048401612384565b5460a01c67ffffffffffffffff1690565b6109716040519283928361251e565b0390a20191906105c8565b6109869087611d34565b906109bc6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161232e565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a3360c088016120ed565b1615610609565b5060ff610a49604088016120d5565b1615610602565b5067ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682146105fb565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610ab060043561033a565b610ac4602435610abf81610a83565b6126a1565b6040516001600160a01b039091168152602090f35b3461016a57610ae73661054e565b906001600160a01b0360035416918215610c005760005b818110610b0757005b610b186103d36104c283858761440b565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610bfb576001948892600092610bcb575b5081610b7f575b5050505001610afe565b81610baf7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610bbf946156ef565b6040519081529081906020820190565b0390a338858180610b75565b610bed91925060203d8111610bf4575b610be581836101fb565b81019061441b565b9038610b6e565b503d610bdb565b612695565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610c485750505090565b82516001600160a01b0316845260209384019390920191600101610c3b565b90610d429160208152610c866020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d2c610d16610100850151610160610120860152610180850190610c2a565b610120850151601f198583030184860152610c2a565b92015190610160601f19828503019101526102b2565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d6b8161033a565b6060610140604051610d7c81610185565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152166000526004602052610336610dd96040600020611eeb565b60405191829182610c67565b61022e9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e37816101a7565b8281528260208201520152610336604051610e51816101a7565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610de5565b3461016a57600060031936011261016a576000546001600160a01b0381163303610f065773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022e82610a83565b8015150361016a57565b359061022e82610f62565b3461016a57606060031936011261016a576000604051610f96816101a7565b600435610fa281610a83565b8152602435610fb081610f62565b6020820190815260443590610fc482610a83565b60408301918252610fd361429c565b6001600160a01b03835116159182156110e8575b5081156110dd575b506110b55780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff00000000000000000000000000000000000000001617905560408201516003805473ffffffffffffffffffffffffffffffffffffffff1916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c906110a0611c4d565b6110af6040519283928361442a565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fef565b516001600160a01b031615915038610fe7565b3461016a57608060031936011261016a5761111760043561033a565b60243567ffffffffffffffff811161016a57611137903690600401610357565b60443590611146606435610a83565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610bfb57600091611ab8575b50611a7b5760025460a01c60ff16611a5157611228740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61124860043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b036064351615611a275780549061126f6103d36001600160a01b03841681565b33036119fd578383916112856080840184611d34565b949060208501956112968787611d34565b90501590816119e4575b6112a985611eeb565b6112b69390600435612e34565b93849160a01c67ffffffffffffffff166112cf90612755565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b5261134d908b6101fb565b604080516064356001600160a01b031660208083019190915281529a90611374908c6101fb565b61137e8680611d34565b86549c9160e08e901c60ff1691369061139692612774565b9060ff166113a3916144da565b9060a089015192604089016113b8908a611d67565b6113c291506127ed565b946113cd908a611d34565b9690976113d861021e565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261144c60048801611e2b565b6101008b01526101208a0152610140890152610160880152610180870152369061147592612774565b6101a0850152611483612005565b6114906040840184611d67565b905061193e575b906114df6114ba8594936040611510970151606087015161046460028701611d9d565b60608601528060408601526114d960808601516001600160a01b031690565b90614564565b60c08601526114ec61283c565b976114fa6040840184611d67565b151590506119305760f01c90505b60043561395f565b63ffffffff9091166060840152602086019391845211611906576115358251866145f2565b6115426040860186611d67565b90506117e6575b61155581959295614f4f565b808552602081519101209061156e6040850151516128ad565b9460408101958652606060009401935b604086015180518210156117535760206115b16103d36103d36115a4866115f896612899565b516001600160a01b031690565b6115bf8460608b0151612899565b519060405180809581947f958021a7000000000000000000000000000000000000000000000000000000008352600435600484016128f6565b03915afa8015610bfb576001600160a01b0391600091611725575b501680156116cc57906000878b9387838861167661163f8860608f61163790612076565b980151612899565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612a41565b03925af18015610bfb57816116a4916001946000916116ab575b508a519061169e8383612899565b52612899565b500161157e565b6116c6913d8091833e6116be81836101fb565b810190612959565b38611690565b61054a6116e06115a48460408b0151612899565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611746915060203d811161174c575b61173e81836101fb565b810190612680565b38611613565b503d611734565b61033685808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6117858d612076565b925193519051906117b66040519283926001600160a01b03606435169767ffffffffffffffff600435169785612c0c565b0390a4610baf7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6118486118326117fc6104bc6040890189611d67565b60c08601805151156118eb5751905b602087015161ffff169060e0880151926064359161182d60043591369061285b565b61488d565b610180830151906118428261288c565b5261288c565b5061186a604061185e8451828701515190612899565b51015163ffffffff1690565b60a061187a61018084015161288c565b5101515163ffffffff82168111611892575050611549565b61054a92506118aa6104c26104bc60408a018a611d67565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506119006118f98980611d34565b3691612774565b9061180b565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff16611508565b509350600191506040611952910187611d67565b9050036119ba57611510838688946114df6114ba6119ae61197c6104c26104bc6040880188611d67565b602061198e6104bc6040890189611d67565b013561199f602089015161ffff1690565b9060e0890151926004356131b4565b92939495505050611497565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119f36040870187611d67565b90501515906112a0565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261054a6004359067ffffffffffffffff60249216600452565b611ada915060203d602011611ae0575b611ad281836101fb565b810190612740565b386111d2565b503d611ac8565b3461016a57602060031936011261016a5767ffffffffffffffff600435611b0d8161033a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611b515760405167ffffffffffffffff9091168152602090f35b6121a3565b3461016a57602060031936011261016a576001600160a01b03600435611b7b81610a83565b611b8361429c565b16338114611bdd578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611c2360043561033a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611c5d816101c3565b8281528260208201528260408201520152604051611c7a816101c3565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611dcf57505061022e925003836101fb565b84546001600160a01b0316835260019485019487945060209093019201611dba565b90600182811c92168015611e21575b6020831014611e0b57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611e00565b9060405191826000825492611e3f84611df1565b8084529360018116908115611eab5750600114611e64575b5061022e925003836101fb565b90506000929192526020600020906000915b818310611e8f57505090602061022e9282010138611e57565b6020919350806001915483858901015201910190918492611e76565b6020935061022e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e57565b90611fe56004611ef9610230565b93611f68611f5d8254611f22611f15826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611f5760e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b611fbb611fab6001830154611f8c611f818261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b611fc760028201611d9d565b610100860152611fd960038201611d9d565b61012086015201611e2b565b610140830152565b67ffffffffffffffff81116101a25760051b60200190565b604051906120146020836101fb565b6000808352366020840137565b9061202b82611fed565b61203860405191826101fb565b828152601f196120488294611fed565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b90156120715790565b612052565b35610d4281610a83565b91908110156120715760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610d428161033a565b60ff81160361016a57565b35610d42816120ca565b63ffffffff81160361016a57565b35610d42816120df565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061213981611fed565b9361214760405195866101fb565b602085838152019160051b810192831161016a57905b82821061216957505050565b60208091833561217881610a83565b81520191019061215d565b61ffff81160361016a57565b35610d4281612183565b35610d4281610f62565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611b5157565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611b5157565b908160031b9180830460081490151715611b5157565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611b5157565b81810292918115918404141715611b5157565b81811061227f575050565b60008155600101612274565b9067ffffffffffffffff83116101a2576801000000000000000083116101a25781548383558084106122ef575b5090600052602060002060005b8381106122d25750505050565b60019060208435946122e386610a83565b019381840155016122c5565b61230790836000528460206000209182019101612274565b386122b8565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d4293818152019161230d565b9190601f811161234e57505050565b61022e926000526020600020906020601f840160051c8301931061237a575b601f0160051c0190612274565b909150819061236d565b90929167ffffffffffffffff81116101a2576123aa816123a48454611df1565b8461233f565b6000601f82116001146123ea5781906123db9394956000926123df575b50506000198260011b9260031b1c19161790565b9055565b0135905038806123c7565b601f198216946123ff84600052602060002090565b91805b87811061243a575083600195969710612420575b505050811b019055565b60001960f88560031b161c19910135169055388080612416565b90926020600181928686013581550194019101612402565b359061022e826120ca565b359061022e82612183565b359061022e826120df565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124c25750505090565b9091926020806001926001600160a01b0387356124de81610a83565b1681520194019291016124b5565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610d42939216815260406020820152612554604082016125468461034c565b67ffffffffffffffff169052565b61257361256360208401610f57565b6001600160a01b03166060830152565b61258c61258260408401612452565b60ff166080830152565b6125a461259b60608401610f6c565b151560a0830152565b6125be6125b36080840161245d565b61ffff1660c0830152565b6125d86125cd60a0840161245d565b61ffff1660e0830152565b6125f56125e760c08401612468565b63ffffffff16610100830152565b61266d61264061261f61260b60e0860186612473565b6101606101208701526101a08601916124a8565b61262d610100860186612473565b90603f19868403016101408701526124a8565b926126626126516101208301610f57565b6001600160a01b0316610160850152565b6101408101906124ec565b91610180603f198286030191015261230d565b9081602091031261016a5751610d4281610a83565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610bfb576001600160a01b039160009161272357501690565b61273c915060203d60201161174c5761173e81836101fb565b1690565b9081602091031261016a5751610d4281610f62565b67ffffffffffffffff1667ffffffffffffffff8114611b515760010190565b9291926127808261025e565b9161278e60405193846101fb565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101a257604052606060a0836000815282602082015282604082015282808201528260808201520152565b906127f782611fed565b61280460405191826101fb565b828152601f196128148294611fed565b019060005b82811061282557505050565b6020906128306127ab565b82828501015201612819565b60405190612849826101a7565b60606040838281528260208201520152565b919082604091031261016a57604051612873816101df565b6020808294803561288381610a83565b84520135910152565b8051156120715760200190565b80518210156120715760209160051b010190565b906128b782611fed565b6128c460405191826101fb565b828152601f196128d48294611fed565b019060005b8281106128e557505050565b8060606020809385010152016128d9565b60409067ffffffffffffffff610d42949316815281602082015201906102b2565b81601f8201121561016a57805161292d8161025e565b9261293b60405194856101fb565b8184526020828401011161016a57610d42916020808501910161028f565b9060208282031261016a57815167ffffffffffffffff811161016a57610d429201612917565b9080602083519182815201916020808360051b8301019401926000915b8383106129ab57505050505090565b9091929394602080612a3283601f1986600196030187528951908151815260a0612a21612a0f6129fd6129eb8887015160c08a88015260c08701906102b2565b604087015186820360408801526102b2565b606086015185820360608701526102b2565b608085015184820360808601526102b2565b9201519060a08184039101526102b2565b9701930193019193929061299c565b919390610d429593612b89612ba19260a08652612a6b60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612b74612b5c612b44612b2c612b14612afe8c61026060e08a0151916101c061018082015201906102b2565b6101008801518d8203609f1901888f01526102b2565b6101208701518c8203609f19016101c08e01526102b2565b610140860151609f198c8303016101e08d01526102b2565b610160850151609f198b8303016102008c01526102b2565b610180840151609f198a8303016102208b015261297f565b910151609f19878303016102408801526102b2565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102b2565b9080602083519182815201916020808360051b8301019401926000915b838310612bdf57505050505090565b9091929394602080612bfd83601f19866001960301875289516102b2565b97019301930191939290612bd0565b9493916001600160a01b03612c2f921686526080602087015260808601906102b2565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612c715750505050610d429394506060818403910152612bb3565b90919295602080612cd283601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102b2565b980192019201909291612c53565b60405190610100820182811067ffffffffffffffff8211176101a257604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612d8b575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612de5918501612917565b926020810151612df4816120df565b92604082015167ffffffffffffffff811161016a57610d429201612917565b60409067ffffffffffffffff610d429593168152816020820152019161230d565b91929092612e40612ce0565b600483101580613026575b15612f70575090612e5b916151ec565b92612e6960408501516153db565b6040840190815151159081612f46575b50612f28575b5060c083015151612ed8575b50608082016001600160a01b03612ea982516001600160a01b031690565b1615612eb457505090565b612ecb60e0610d429301516001600160a01b031690565b6001600160a01b03169052565b612eec612ee86060840151151590565b1590565b15612e8b577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612f3b90610120840151809152516128ad565b606084015238612e7f565b905080612f55575b1538612e79565b5063ffffffff612f69855163ffffffff1690565b1615612f4e565b9491600090612fc892612f916103d36103d36002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612e13565b03915afa8015610bfb57600090600090600090612ffa575b60a088015263ffffffff16865290505b60c0850152612e69565b50505061301c612ff0913d806000833e61301481836101fb565b810190612dbd565b9192508291612fe0565b5063302326cb60e01b7fffffffff0000000000000000000000000000000000000000000000000000000061306361305d8686612d31565b90612d57565b1614612e4b565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a57815161309e81611fed565b926130ac60405194856101fb565b81845260208085019260051b82010192831161016a57602001905b8282106130d45750505090565b6020809183516130e381610a83565b8152019101906130c7565b95949060009460a09467ffffffffffffffff613135956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102b2565b930152565b9060028201809211611b5157565b9060018201809211611b5157565b6001019081600111611b5157565b9060148201809211611b5157565b90600c8201809211611b5157565b91908201809211611b5157565b6000198114611b515760010190565b80548210156120715760005260206000200190600090565b929394919060036131d98567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036131ee8184166126a1565b169182156133c6576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bfb576000916133a7575b50156133975761326d600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016130ee565b03915afa928315610bfb57600093613372575b508251156133675782519061329f61329a84548094613180565b612021565b906000928394845b8751811015613306576132bd6115a4828a612899565b6001600160a01b038116156132fa57906132f46001926132e66132df8a61318d565b9989612899565b906001600160a01b03169052565b016132a7565b509550600180966132f4565b509195509193613318575b5050815290565b60005b8281106133285750613311565b8061336161334e61333b6001948661319c565b90546001600160a01b039160031b1c1690565b6132e661335a8861318d565b9789612899565b0161331b565b9150610d4290611d9d565b6133909193503d806000833e61338881836101fb565b81019061306a565b9138613280565b505050509250610d429150611d9d565b6133c0915060203d602011611ae057611ad281836101fb565b38613228565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936134186134108251865190613180565b865190613180565b9061342b61342583612021565b926128ad565b94600096875b8351891015613491578861348761347a6001936134626134586115a48e9f9d9e9d8b612899565b6132e6838c612899565b61348061346f858c612899565b51918093849161318d565b9c612899565b528b612899565b5001979695613431565b959250929350955060005b8651811015613525576134b26115a48289612899565b60006001600160a01b038216815b8881106134f9575b50509060019291156134dc575b500161349c565b6134f3906132e66134ec8961318d565b9888612899565b386134d5565b8161350a6103d36115a4848c612899565b14613517576001016134c0565b5060019150819050386134c8565b509390945060005b85518110156135b6576135436115a48288612899565b60006001600160a01b038216815b87811061358a575b505090600192911561356d575b500161352d565b613584906132e661357d8861318d565b9787612899565b38613566565b8161359b6103d36115a4848b612899565b146135a857600101613551565b506001915081905038613559565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101a25760405260606080836000815260006020820152600060408201526000838201520152565b9061360b82611fed565b61361860405191826101fb565b828152601f196136288294611fed565b019060005b82811061363957505050565b6020906136446135c2565b8282850101520161362d565b9081606091031261016a57805161366681612183565b9160406020830151613677816120df565b920151610d42816120df565b9160209082815201919060005b81811061369d5750505090565b9091926040806001926001600160a01b0387356136b981610a83565b16815260208781013590820152019401929101613690565b949391929067ffffffffffffffff1685526080602086015261372a61370b6136f985806124ec565b60a060808a015261012089019161230d565b61371860208601866124ec565b90607f198984030160a08a015261230d565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022e956137b2613789836060976137d3978d60c0607f19826137c59a0301910152613683565b916137a8613798888301610f57565b6001600160a01b031660e08d0152565b60808101906124ec565b90607f198b8403016101008c015261230d565b9087820360408901526102b2565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611b5157565b908160a091031261016a578051916020820151613814816120df565b916040810151613823816120df565b916080606083015161383481612183565b920151610d4281610f62565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d429b9a9616885216602087015260408601521660608401521660808201528160a082015201906102b2565b9081606091031261016a578051613666816120df565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611b5157565b906000198201918211611b5157565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611b5157565b91908203918211611b5157565b919082608091031261016a57815161392b816120df565b916020810151916060604083015192015190565b8115613949570490565b634e487b7160e01b600052601260045260246000fd5b938294600090600095604081019461399961399461398f88515161398760408c01809c611d67565b919050613180565b61313a565b613601565b9660009586955b88518051881015613bfd576103d36103d36115a48a6139be94612899565b613a0b602060608801926139d38b8551612899565b51908a6040518095819482937f958021a7000000000000000000000000000000000000000000000000000000008452600484016128f6565b03915afa8015610bfb576001600160a01b0391600091613bdf575b50168015613b8b579060608e9392613a3f8b8451612899565b5190613a5060208b015161ffff1690565b958b613a8b604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016136d1565b03915afa8015610bfb57600193613b28938b8f8f95600080958197613b31575b509083929161ffff613ad385613acc6115a4613b1c99613b229d9e51612899565b9451612899565b5191613aef613ae0610240565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261169e8383612899565b506137de565b996137de565b960195966139a0565b613b2297506115a4965084939291509361ffff613ad382613acc613b6e613b1c9960603d8111613b84575b613b6681836101fb565b810190613650565b9c9196909c9d5050505050505090919293613aab565b503d613b5c565b61054a88613b9d6115a48c8f51612899565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613bf7915060203d811161174c5761173e81836101fb565b38613a26565b50919a9496929395509897968a613c148187611d67565b9050613f28575b50508651613c289061389e565b99613c366020860186611d34565b91613c42915086611d67565b9560609150019486613c5387612076565b91613c5e938a6154e6565b613c688b89612899565b52613c738a88612899565b50613c7e8a88612899565b516020015163ffffffff16613c92916137de565b90613c9d8a88612899565b516040015163ffffffff16613cb1916137de565b91613cba610240565b33815290600060208301819052604083015261ffff166060820152613cdd61027a565b60808201528651613ced906138cb565b90613cf88289612899565b52613d039087612899565b506002546001600160a01b031692613d1a90612076565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610bfb57600096600097600094600092613eeb575b506000965b8651881015613e6b57613df4600191613dbc88613db7876121b9565b61393f565b613dd56060613dcb8d8d612899565b5101918251612261565b9052858a14613dfc575b6060613deb8b8b612899565b51015190613180565b970196613d9b565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613e2f60808c01516001600160a01b031690565b1603613e3d575b5050613ddf565b613db7613e49926121e9565b613e626060613e588d8d612899565b5101918251613180565b90528b88613e36565b9796509750505050613ea77f000000000000000000000000000000000000000000000000000000000000000091613db763ffffffff84166121e9565b8411613eb35750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613f1391925060803d608011613f21575b613f0b81836101fb565b810190613914565b919790939290919038613d96565b503d613f01565b610abf6103d36104c26104bc613f41948a989698611d67565b926001600160a01b03600091515194169060e08801908151613f61610240565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613f8f878d612899565b52613f9a868c612899565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610bfb578f948c89968f96948d948f96889161427d575b50614176575b5050505050501561401d575b61185e61400e61401595614008602061185e61400897604097612899565b906137de565b958b612899565b90388a613c1b565b50506140a39160608c61404e6104c26104bc6140476103d36103d36002546001600160a01b031690565b938b611d67565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610bfb5761185e61400e604092614008602061185e8f8b906140159b6140089a600080600092614138575b63ffffffff9293506141259060606140ed8888612899565b51019461411a8a6140fe8a8a612899565b510191604061410d8b8b612899565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050613fea565b50505063ffffffff6141646141259260603d60601161416f575b61415c81836101fb565b810190613888565b9093509150826140d5565b503d614152565b8495985060a09697506141c060206141b66060826141ad6104bc6141a66104c26104bc8b6141f99c9d9e9f611d67565b998d611d67565b01359901612076565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613840565b03915afa8015610bfb578592828c939181908294614241575b506142359060606142238888612899565b51019261411a60206140fe8a8a612899565b5288888f8c8138613fde565b915050614235925061426b915060a03d60a011614276575b61426381836101fb565b8101906137f8565b949192919050614212565b503d614259565b614296915060203d602011611ae057611ad281836101fb565b38613fd8565b6001600160a01b036001541633036142b057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916142e8815184613180565b9283156143e15760005b848110614300575050505050565b818110156143c6576143156115a48286612899565b6001600160a01b0381168015610c005761432e83613148565b878110614340575050506001016142f2565b848110156143a3576001600160a01b0361435d6115a4838a612899565b16821461436c5760010161432e565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036143c16115a46143bb8885613907565b89612899565b61435d565b6143dc6115a46143d68484613907565b85612899565b614315565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156120715760051b0190565b9081602091031261016a575190565b91608061022e92949361447a8160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d429281815201906102b2565b906144bd8261025e565b6144ca60405191826101fb565b828152601f19612048829461025e565b908151602082108061455a575b61452957036144f35790565b6109bc906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016144a2565b50906020810151908161453b8461221b565b1c6144f3575061454a826144b3565b9160200360031b1b602082015290565b50602081146144e7565b918251601481029080820460141490151715611b515761458661458b91613156565b613164565b9061459d61459883613172565b6144b3565b9060146145a98361288c565b5360009260215b86518510156145db5760146001916145cb6115a4888b612899565b60601b81870152019401936145b0565b919550936020935060601b90820152828152012090565b906146026103d360608401612076565b614613600019936040810190611d67565b905061468b575b61462482516138cb565b9260005b848110614636575050505050565b80826001921461468657606061464c8287612899565b51015180156146805761467a906146746146668489612899565b51516001600160a01b031690565b866156ef565b01614628565b5061467a565b61467a565b915061469781516138da565b916146a56146668484612899565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610bfb5760009161470c575b506146ec575b5061461a565b6147069060606146fc8686612899565b51015190836156ef565b386146e6565b614725915060203d602011611ae057611ad281836101fb565b386146e0565b60405190614738826101df565b60606020838281520152565b919060408382031261016a576040519061475d826101df565b8193805167ffffffffffffffff811161016a578261477c918301612917565b835260208101519167ffffffffffffffff831161016a576020926147a09201612917565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610d429201614744565b9060806001600160a01b03816147ea855160a0865260a08601906102b2565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d429281815201906147cb565b919060408382031261016a57825167ffffffffffffffff811161016a5760209161485a918501614744565b92015190565b61ffff614879610d4295936060845260608401906147cb565b9316602082015260408184039101526102b2565b90919392959461489b6127ab565b5060208201805115614cea576148be610abf6103d385516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061490c60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610bfb57600091614ccb575b5015614c825761499888999a82519261493761472b565b505161498361494d89516001600160a01b031690565b926040614958610240565b9e8f9081526149748d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bfb57600091614c63575b5015614b6b5750906000929183614a129899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614860565b03925af18015610bfb57600094600091614b23575b50614af8614a619596614a9b614a6f614a8d956115a46020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101fb565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101fb565b614ad7610905614acd614add8951614ad7610905614acd8c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906144da565b9767ffffffffffffffff166000526004602052604060002090565b93015193614b0461024f565b958652602086015260408501526060840152608083015260a082015290565b614a619550602091506115a496614a9b614a6f614a8d95614b59614af8953d806000833e614b5181836101fb565b81019061482f565b9b909b96505095505050969550614a27565b9793929061ffff16614c395751614c0f57614bba6000939184926040519586809481937f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161481e565b03925af1908115610bfb57614a6195614a9b614a6f614af8936115a4602096614a8d98600091614bec575b5099614a43565b614c0991503d806000833e614c0181836101fb565b8101906147a5565b38614be5565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614c7c915060203d602011611ae057611ad281836101fb565b386149ca565b61054a614c9686516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614ce4915060203d602011611ae057611ad281836101fb565b38614920565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614de6947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614e026020928281519485920161028f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152614e4e825180936020898501910161028f565b019160f81b1683820152614e6c82518093602060028501910161028f565b01019160f81b1683820152614e8b82518093602060028501910161028f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d429f9e9c9860f81b168152614f00825180936020898501910161028f565b019160f01b1683820152614f1e82518093602060038501910161028f565b01019160f01b1683820152614f3c825180936020898501910161028f565b01019160f01b1660028201520190614def565b60e081019060ff825151116151d657610100810160ff815151116151c05761012082019260ff845151116151aa57610140830160ff815151116151945761016084019061ffff8251511161517e57610180850193600185515111615166576101a086019361ffff8551511161515057606095518051615134575b50865167ffffffffffffffff16602088015167ffffffffffffffff16906040890151614ffc9067ffffffffffffffff1690565b98606081015161500f9063ffffffff1690565b9060808101516150229063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661504597614d14565b03601f198101885261505790886101fb565b519081516150659060ff1690565b9051805160ff16985190815161507b9060ff1690565b906040519a8b95602087019561509096614e06565b03601f19810187526150a290876101fb565b519182516150b09060ff1690565b9151805161ffff169480516150c69061ffff1690565b92519283516150d69061ffff1690565b9260405197889760208901976150eb98614e91565b03601f19810182526150fd90826101fb565b6040519283926020840161511091614def565b61511991614def565b61512291614def565b03601f1981018252610d4290826101fb565b6151499196506151439061288c565b516158e6565b9438614fc9565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b60005261054a6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b906151f5612ce0565b91601182106153c057803563302326cb60e01b7fffffffff000000000000000000000000000000000000000000000000000000008216036153665750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61525b81612021565b6040860190815261526b826128ad565b906060870191825260005b83811061531a57505050506152de83836152d46152c86152be6152b76152a46152e898876152f29c9b615a0d565b6001600160a01b0390911660808d015290565b8585615ab1565b9291903691612774565b60a08a01528383615b00565b9491903691612774565b60c0880152615ab1565b9391903691612774565b60e08401528103615301575090565b63d9437f9d60e01b600052600360045260245260446000fd5b8060019161535f6153496153426153356153599a8d8d615a0d565b91906132e6868a51612899565b8b8b615ab1565b9391889a919a51949a3691612774565b92612899565b5201615276565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263302326cb60e01b6004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b8281106153ed57505050565b60018101808211611b51575b83811061540957506001016153e1565b6001600160a01b0361541b8385612899565b511661542d6103d36115a48487612899565b1461543a576001016153f9565b61054a61544a6115a48486612899565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9081602091031261016a5751610d4281612183565b936154d160809461ffff6001600160a01b039567ffffffffffffffff6154df969b9a9b16895216602088015260a0604088015260a0870190610c2a565b9085820360608701526102b2565b9416910152565b929190926154f26135c2565b506155118167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff169160808501928351615534906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615554916137de565b9661556090608d613180565b9460a087019586515161557291613180565b9160ff169161558083612231565b61558991613180565b91615595906067613180565b61559e91612261565b6155a791613180565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b038716036156345750505061ffff9250615626906156196000935b519561560c6155fd610240565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b61564a6103d3602094976001600160a01b031690565b90604061565b8583015161ffff1690565b9101519261569b875198604051998a96879586957ff238895800000000000000000000000000000000000000000000000000000000875260048701615494565b03915afa908115610bfb576156196156269261ffff956000916156c0575b50936155f0565b6156e2915060203d6020116156e8575b6156da81836101fb565b81019061547f565b386156b9565b503d6156d0565b91602091600091604051906001600160a01b03858301937fa9059cbb0000000000000000000000000000000000000000000000000000000085521660248301526044820152604481526157436064826101fb565b519082855af115612695576000513d6157a257506001600160a01b0381163b155b61576b5750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615764565b9660026158ba976158876022610d429f9e9c9799600199859f9b7fff00000000000000000000000000000000000000000000000000000000000000906158879f826158879c61588e9c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615838825180936020898501910161028f565b019160f81b168382015261585682518093602060238501910161028f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614def565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116159f757604081019160ff835151116159e157606082019160ff835151116159cb57608081019260ff845151116159b55760a0820161ffff8151511161599f57610d42946159919351945191615948835160ff1690565b975191615956835160ff1690565b945190615964825160ff1690565b905193615972855160ff1690565b935196615981885161ffff1690565b966040519c8d9b60208d016157ab565b03601f1981018352826101fb565b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b635a102da160e11b600052602660045260246000fd5b929190926001820191848311615a985781013560001a828115615a8d575060148103615a60578201938411615a4557013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b91906002820191818311615a98578381013560f01c0160020192818411615ae557918391615ade93612d3f565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b91906001820191818311615a98578381013560001a0160010192818411615ae557918391615ade93612d3f56fea164736f6c634300081a000a",
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
