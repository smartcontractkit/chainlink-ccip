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
	Bin: "0x61010060405234610367576040516163a338819003601f8101601f191683016001600160401b0381118482101761036c578392829160405283398101039060e08212610367576080821261036757610055610382565b81519092906001600160401b03811681036103675783526020820151926001600160a01b03841684036103675760208101938452604083015163ffffffff81168103610367576040820190815260606100af8186016103a1565b83820190815293607f1901126103675760405192606084016001600160401b0381118582101761036c576040526100e8608086016103a1565b845260a08501519485151586036103675760c061010c9160208701978852016103a1565b9560408501968752331561035657600180546001600160a01b0319163317905583516001600160401b0316158015610344575b8015610332575b8015610323575b6103085792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e052815116158015610319575b6103085780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e09260006060610206610382565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff92831691166060610251610382565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615fed90816103b68239608051818181610ac4015281816114830152611d06015260a0518181816112410152611d32015260c051818181611d8d01526127ee015260e051818181611d5e0152613ffb0152f35b6306b7c75960e31b60005260046000fd5b508151151561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761036c57604052565b51906001600160a01b03821682036103675756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611c88565b611bcc565b611b5d565b61119b565b610fea565b610fa3565b610ef2565b610e7f565b610dad565b610b41565b610afc565b61060a565b61037e565b6102f0565b3461016a57600060031936011261016a576080610122611cce565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101bb57604052565b61016f565b6060810190811067ffffffffffffffff8211176101bb57604052565b6080810190811067ffffffffffffffff8211176101bb57604052565b6040810190811067ffffffffffffffff8211176101bb57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101bb57604052565b604051906102476101c083610214565b565b6040519061024761016083610214565b6040519061024760a083610214565b6040519061024760c083610214565b67ffffffffffffffff81116101bb57601f01601f191660200190565b604051906102a2602083610214565b60008252565b60005b8381106102bb5750506000910152565b81810151838201526020016102ab565b90601f19601f6020936102e9815180928187528780880191016102a8565b0116010190565b3461016a57600060031936011261016a5761034f60408051906103138183610214565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102cb565b0390f35b67ffffffffffffffff81160361016a57565b359061024782610353565b908160a091031261016a5790565b3461016a57604060031936011261016a5760043561039b81610353565b60243567ffffffffffffffff811161016a576103bb903690600401610370565b6103d98267ffffffffffffffff166000526004602052604060002090565b918254916001600160a01b036104046103f8856001600160a01b031690565b6001600160a01b031690565b161561056f576040810191600161041b8484611db5565b9050116105455761034f946104be946104aa61046361043d6080870187611deb565b61044a6020890189611deb565b905015918261052f575b61045d87611f85565b88612f70565b9561046c61209f565b6104768288611db5565b90506104de575b604088016104a081519260608b0193845161049a60028b01611e1e565b91613566565b9092525285611db5565b151590506104d05760f01c90505b90613ae2565b60405190815292839250602083019150565b506001015461ffff166104b8565b5061052a6104fd6104f86104f2848a611db5565b9061211b565b612129565b602061050c6104f2858b611db5565b013561051d60208b015161ffff1690565b9060e08b015192896132ec565b61047d565b915061053b8989611db5565b9050151591610454565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156106065781359367ffffffffffffffff851161060357506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a57610618366105ab565b90610621614451565b6000915b80831061062e57005b610639838284612133565b9261064384612173565b67ffffffffffffffff81169081158015610ab8575b8015610aa2575b8015610a89575b610a5257856108c8916108e26108d8836108d26106d860e08301956106be6106b861069189876121aa565b6106b06106a66101008a95949501809a6121aa565b94909236916121e0565b9236916121e0565b9061448f565b67ffffffffffffffff166000526004602052604060002090565b96879561071e6106ea60208a01612129565b88906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6108c261088b60c060408b019a6107856107378d612188565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b6107e361079460808301612242565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b61082660016107f460a08401612242565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108858d6108366060840161224c565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016121a0565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6121aa565b9060038901612357565b8a6121aa565b9060028601612357565b6101208801906108f46103f883612129565b15610a285761090561094f92612129565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061097361096d610965848b611deb565b939050612188565b60ff1690565b036109e457956109ca7f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26926109b96109af600198999a85611deb565b9060048401612450565b5460a01c67ffffffffffffffff1690565b6109d9604051928392836125ea565b0390a2019190610625565b6109ee9087611deb565b90610a246040519283927f3aeba390000000000000000000000000000000000000000000000000000000008452600484016123fa565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a9b60c088016121a0565b1615610666565b5060ff610ab160408801612188565b161561065f565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610658565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610b18600435610353565b610b2c602435610b2781610aeb565b6127a9565b6040516001600160a01b039091168152602090f35b3461016a57610b4f366105ab565b906001600160a01b0360035416918215610c685760005b818110610b6f57005b610b806103f86104f88385876145c0565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c63576001948892600092610c33575b5081610be7575b5050505001610b66565b81610c177f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610c2794615b57565b6040519081529081906020820190565b0390a338858180610bdd565b610c5591925060203d8111610c5c575b610c4d8183610214565b8101906145d0565b9038610bd6565b503d610c43565b61279d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610cb05750505090565b82516001600160a01b0316845260209384019390920191600101610ca3565b90610daa9160208152610cee6020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d94610d7e610100850151610160610120860152610180850190610c92565b610120850151601f198583030184860152610c92565b92015190610160601f19828503019101526102cb565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610dd381610353565b6060610140604051610de48161019e565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e08201528261010082015282610120820152015216600052600460205261034f610e416040600020611f85565b60405191829182610ccf565b6102479092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e9f816101c0565b828152826020820152015261034f604051610eb9816101c0565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610e4d565b3461016a57600060031936011261016a576000546001600160a01b0381163303610f79577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061024782610aeb565b8015150361016a57565b359061024782610fd5565b3461016a57606060031936011261016a576000604051611009816101c0565b60043561101581610aeb565b815260243561102381610fd5565b602082019081526044359061103782610aeb565b60408301918252611046614451565b6001600160a01b03835116158015611191575b611169576001600160a01b03839261111661114c936110ca847f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9851166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006002541617600255565b5115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000060025492151560a01b16911617600255565b51166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006003541617600355565b611154611cce565b611163604051928392836145df565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515611059565b3461016a57608060031936011261016a576111b7600435610353565b60243567ffffffffffffffff811161016a576111d7903690600401610370565b604435906111e6606435610aeb565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c6357600091611b2e575b50611af15760025460a01c60ff16611ac7576112c8740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6112e860043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b036064351615611a9d5780549061130f6103f86001600160a01b03841681565b3303611a73578383916113256080840184611deb565b949060208501956113368787611deb565b9050159081611a5a575b61134985611f85565b6113569390600435612f70565b93849160a01c67ffffffffffffffff1661136f9061285d565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b526113ed908b610214565b604080516064356001600160a01b031660208083019190915281529a90611414908c610214565b61141e8680611deb565b86549c9160e08e901c60ff169136906114369261287c565b9060ff166114439161468f565b9060a08901519260408901611458908a611db5565b61146291506128f5565b9461146d908a611deb565b969097611478610237565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b01526114ec60048801611ec5565b6101008b01526101208a015261014089015261016088015261018087015236906115159261287c565b6101a085015261152361209f565b6115306040840184611db5565b90506119de575b9061157f61155a85949360406115b0970151606087015161049a60028701611e1e565b606086015280604086015261157960808601516001600160a01b031690565b90614719565b60c086015261158c612944565b9761159a6040840184611db5565b151590506119d05760f01c90505b600435613ae2565b63ffffffff90911660608401526020860193918452116119a6576115d58251866147a7565b6115e26040860186611db5565b9050611886575b6115f581959295615181565b808552602081519101209061160e6040850151516129e9565b9460408101958652606060009401935b604086015180518210156117f35760206116516103f86103f861164486611698966129a1565b516001600160a01b031690565b61165f8460608b01516129a1565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612a32565b03915afa8015610c63576001600160a01b03916000916117c5575b5016801561176c57906000878b938783886117166116df8860608f6116d790612129565b9801516129a1565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612b7d565b03925af18015610c6357816117449160019460009161174b575b508a519061173e83836129a1565b526129a1565b500161161e565b611766913d8091833e61175e8183610214565b810190612a95565b38611730565b6105a76117806116448460408b01516129a1565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b6117e6915060203d81116117ec575b6117de8183610214565b810190612788565b386116b3565b503d6117d4565b61034f85808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6118258d612129565b925193519051906118566040519283926001600160a01b03606435169767ffffffffffffffff600435169785612d48565b0390a4610c177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6118e86118d261189c6104f26040890189611db5565b60c086018051511561198b5751905b602087015161ffff169060e088015192606435916118cd600435913690612963565b614a74565b610180830151906118e282612994565b52612994565b5061190a60406118fe84518287015151906129a1565b51015163ffffffff1690565b60a061191a610180840151612994565b5101515163ffffffff821681116119325750506115e9565b6105a7925061194a6104f86104f260408a018a611db5565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506119a06119998980611deb565b369161287c565b906118ab565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166115a8565b5093506001915060406119f2910187611db5565b905003610545576115b08386889461157f61155a611a4e611a1c6104f86104f26040880188611db5565b6020611a2e6104f26040890189611db5565b0135611a3f602089015161ffff1690565b9060e0890151926004356132ec565b92939495505050611537565b9050611a696040870187611db5565b9050151590611340565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105a76004359067ffffffffffffffff60249216600452565b611b50915060203d602011611b56575b611b488183610214565b810190612848565b38611272565b503d611b3e565b3461016a57602060031936011261016a5767ffffffffffffffff600435611b8381610353565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611bc75760405167ffffffffffffffff9091168152602090f35b612256565b3461016a57602060031936011261016a576001600160a01b03600435611bf181610aeb565b611bf9614451565b16338114611c5e57807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611ca4600435610353565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611cde816101dc565b8281528260208201528260408201520152604051611cfb816101dc565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611e5057505061024792500383610214565b84546001600160a01b0316835260019485019487945060209093019201611e3b565b90600182811c92168015611ebb575b6020831014611e8c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611e81565b9060405191826000825492611ed984611e72565b8084529360018116908115611f455750600114611efe575b5061024792500383610214565b90506000929192526020600020906000915b818310611f295750509060206102479282010138611ef1565b6020919350806001915483858901015201910190918492611f10565b602093506102479592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611ef1565b9061207f6004611f93610249565b93612002611ff78254611fbc611faf826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611ff160e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b612055612045600183015461202661201b8261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b61206160028201611e1e565b61010086015261207360038201611e1e565b61012086015201611ec5565b610140830152565b67ffffffffffffffff81116101bb5760051b60200190565b604051906120ae602083610214565b6000808352366020840137565b906120c582612087565b6120d26040519182610214565b828152601f196120e28294612087565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156121245790565b6120ec565b35610daa81610aeb565b91908110156121245760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610daa81610353565b60ff81160361016a57565b35610daa8161217d565b63ffffffff81160361016a57565b35610daa81612192565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b9291906121ec81612087565b936121fa6040519586610214565b602085838152019160051b810192831161016a57905b82821061221c57505050565b60208091833561222b81610aeb565b815201910190612210565b61ffff81160361016a57565b35610daa81612236565b35610daa81610fd5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611bc757565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611bc757565b908160031b9180830460081490151715611bc757565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611bc757565b81810292918115918404141715611bc757565b81811061234b575050565b60008155600101612340565b9067ffffffffffffffff83116101bb576801000000000000000083116101bb5781548383558084106123bb575b5090600052602060002060005b83811061239e5750505050565b60019060208435946123af86610aeb565b01938184015501612391565b6123d390836000528460206000209182019101612340565b38612384565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610daa9381815201916123d9565b9190601f811161241a57505050565b610247926000526020600020906020601f840160051c83019310612446575b601f0160051c0190612340565b9091508190612439565b90929167ffffffffffffffff81116101bb57612476816124708454611e72565b8461240b565b6000601f82116001146124b65781906124a79394956000926124ab575b50506000198260011b9260031b1c19161790565b9055565b013590503880612493565b601f198216946124cb84600052602060002090565b91805b8781106125065750836001959697106124ec575b505050811b019055565b60001960f88560031b161c199101351690553880806124e2565b909260206001819286860135815501940191016124ce565b35906102478261217d565b359061024782612236565b359061024782612192565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b81811061258e5750505090565b9091926020806001926001600160a01b0387356125aa81610aeb565b168152019401929101612581565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610daa9392168152604060208201526126206040820161261284610365565b67ffffffffffffffff169052565b61263f61262f60208401610fca565b6001600160a01b03166060830152565b61265861264e6040840161251e565b60ff166080830152565b61267061266760608401610fdf565b151560a0830152565b61268a61267f60808401612529565b61ffff1660c0830152565b6126a461269960a08401612529565b61ffff1660e0830152565b6126c16126b360c08401612534565b63ffffffff16610100830152565b61275761272a6126eb6126d760e086018661253f565b6101606101208701526101a0860191612574565b6126f961010086018661253f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086840301610140870152612574565b9261274c61273b6101208301610fca565b6001600160a01b0316610160850152565b6101408101906125b8565b916101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828603019101526123d9565b9081602091031261016a5751610daa81610aeb565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c63576001600160a01b039160009161282b57501690565b612844915060203d6020116117ec576117de8183610214565b1690565b9081602091031261016a5751610daa81610fd5565b67ffffffffffffffff1667ffffffffffffffff8114611bc75760010190565b92919261288882610277565b916128966040519384610214565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101bb57604052606060a0836000815282602082015282604082015282808201528260808201520152565b906128ff82612087565b61290c6040519182610214565b828152601f1961291c8294612087565b019060005b82811061292d57505050565b6020906129386128b3565b82828501015201612921565b60405190612951826101c0565b60606040838281528260208201520152565b919082604091031261016a5760405161297b816101f8565b6020808294803561298b81610aeb565b84520135910152565b8051156121245760200190565b80518210156121245760209160051b010190565b604051906129c4602083610214565b600080835282815b8281106129d857505050565b8060606020809385010152016129cc565b906129f382612087565b612a006040519182610214565b828152601f19612a108294612087565b019060005b828110612a2157505050565b806060602080938501015201612a15565b60409067ffffffffffffffff610daa949316815281602082015201906102cb565b81601f8201121561016a578051612a6981610277565b92612a776040519485610214565b8184526020828401011161016a57610daa91602080850191016102a8565b9060208282031261016a57815167ffffffffffffffff811161016a57610daa9201612a53565b9080602083519182815201916020808360051b8301019401926000915b838310612ae757505050505090565b9091929394602080612b6e83601f1986600196030187528951908151815260a0612b5d612b4b612b39612b278887015160c08a88015260c08701906102cb565b604087015186820360408801526102cb565b606086015185820360608701526102cb565b608085015184820360808601526102cb565b9201519060a08184039101526102cb565b97019301930191939290612ad8565b919390610daa9593612cc5612cdd9260a08652612ba760a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612cb0612c98612c80612c68612c50612c3a8c61026060e08a0151916101c061018082015201906102cb565b6101008801518d8203609f1901888f01526102cb565b6101208701518c8203609f19016101c08e01526102cb565b610140860151609f198c8303016101e08d01526102cb565b610160850151609f198b8303016102008c01526102cb565b610180840151609f198a8303016102208b0152612abb565b910151609f19878303016102408801526102cb565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102cb565b9080602083519182815201916020808360051b8301019401926000915b838310612d1b57505050505090565b9091929394602080612d3983601f19866001960301875289516102cb565b97019301930191939290612d0c565b9493916001600160a01b03612d6b921686526080602087015260808601906102cb565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612dad5750505050610daa9394506060818403910152612cef565b90919295602080612e0e83601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102cb565b980192019201909291612d8f565b60405190610100820182811067ffffffffffffffff8211176101bb57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612ec7575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612f21918501612a53565b926020810151612f3081612192565b92604082015167ffffffffffffffff811161016a57610daa9201612a53565b60409067ffffffffffffffff610daa959316815281602082015201916123d9565b91929092612f7c612e1c565b600483101580613145575b1561308f575090612f979161543e565b92612fa56040850151615691565b80613074575b60408401612fc88151926060870193845161012088015191615735565b9092525260c083015151613024575b50608082016001600160a01b03612ff582516001600160a01b031690565b161561300057505090565b61301760e0610daa9301516001600160a01b031690565b6001600160a01b03169052565b6130386130346060840151151590565b1590565b15612fd7577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff613088845163ffffffff1690565b1615612fab565b94916000906130e7926130b06103f86103f86002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612f4f565b03915afa8015610c6357600090600090600090613119575b60a088015263ffffffff16865290505b60c0850152612fa5565b50505061313b61310f913d806000833e6131338183610214565b810190612ef9565b91925082916130ff565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061319b6131958686612e6d565b90612e93565b1614612f87565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a5781516131d681612087565b926131e46040519485610214565b81845260208085019260051b82010192831161016a57602001905b82821061320c5750505090565b60208091835161321b81610aeb565b8152019101906131ff565b95949060009460a09467ffffffffffffffff61326d956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102cb565b930152565b9060028201809211611bc757565b9060018201809211611bc757565b6001019081600111611bc757565b9060148201809211611bc757565b90600c8201809211611bc757565b91908201809211611bc757565b6000198114611bc75760010190565b80548210156121245760005260206000200190600090565b929394919060036133118567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036133268184166127a9565b16918215613530576040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6357600091613511575b5015613501576133d7600095969798604051998a96879586957f89720a6200000000000000000000000000000000000000000000000000000000875260048701613226565b03915afa928315610c63576000936134dc575b508251156134d157825190613409613404845480946132b8565b6120bb565b906000928394845b875181101561347057613427611644828a6129a1565b6001600160a01b03811615613464579061345e6001926134506134498a6132c5565b99896129a1565b906001600160a01b03169052565b01613411565b5095506001809661345e565b509195509193613482575b5050815290565b60005b828110613492575061347b565b806134cb6134b86134a5600194866132d4565b90546001600160a01b039160031b1c1690565b6134506134c4886132c5565b97896129a1565b01613485565b9150610daa90611e1e565b6134fa9193503d806000833e6134f28183610214565b8101906131a2565b91386133ea565b505050509250610daa9150611e1e565b61352a915060203d602011611b5657611b488183610214565b38613392565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b9391929361358261357a82518651906132b8565b8651906132b8565b9061359561358f836120bb565b926129e9565b94600096875b83518910156135fb57886135f16135e46001936135cc6135c26116448e9f9d9e9d8b6129a1565b613450838c6129a1565b6135ea6135d9858c6129a1565b5191809384916132c5565b9c6129a1565b528b6129a1565b500197969561359b565b959250929350955060005b865181101561368f5761361c61164482896129a1565b60006001600160a01b038216815b888110613663575b5050906001929115613646575b5001613606565b61365d90613450613656896132c5565b98886129a1565b3861363f565b816136746103f8611644848c6129a1565b146136815760010161362a565b506001915081905038613632565b509390945060005b8551811015613720576136ad61164482886129a1565b60006001600160a01b038216815b8781106136f4575b50509060019291156136d7575b5001613697565b6136ee906134506136e7886132c5565b97876129a1565b386136d0565b816137056103f8611644848b6129a1565b14613712576001016136bb565b5060019150819050386136c3565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101bb5760405260606080836000815260006020820152600060408201526000838201520152565b9061377582612087565b6137826040519182610214565b828152601f196137928294612087565b019060005b8281106137a357505050565b6020906137ae61372c565b82828501015201613797565b9081606091031261016a5780516137d081612236565b91604060208301516137e181612192565b920151610daa81612192565b9160209082815201919060005b8181106138075750505090565b9091926040806001926001600160a01b03873561382381610aeb565b168152602087810135908201520194019291016137fa565b949391929067ffffffffffffffff1685526080602086015261389461387561386385806125b8565b60a060808a01526101208901916123d9565b61388260208601866125b8565b90607f198984030160a08a01526123d9565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a576102479561391c6138f38360609761393d978d60c0607f198261392f9a03019101526137ed565b91613912613902888301610fca565b6001600160a01b031660e08d0152565b60808101906125b8565b90607f198b8403016101008c01526123d9565b9087820360408901526102cb565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611bc757565b908160a091031261016a57805191602082015161397e81612192565b91604081015161398d81612192565b916080606083015161399e81612236565b920151610daa81610fd5565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610daa9b9a9616885216602087015260408601521660608401521660808201528160a082015201906102cb565b9081606091031261016a5780516137d081612192565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611bc757565b906000198201918211611bc757565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611bc757565b91908203918211611bc757565b919082608091031261016a578151613a9581612192565b916020810151916060604083015192015190565b8115613ab3570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613b1c613b17613b12885151613b0a60408c01809c611db5565b9190506132b8565b613272565b61376b565b9660009586955b88518051881015613d80576103f86103f86116448a613b41946129a1565b613b8e60206060880192613b568b85516129a1565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612a32565b03915afa8015610c63576001600160a01b0391600091613d62575b50168015613d0e579060608e9392613bc28b84516129a1565b5190613bd360208b015161ffff1690565b958b613c0e604051988995869485947f80485e250000000000000000000000000000000000000000000000000000000086526004860161383b565b03915afa8015610c6357600193613cab938b8f8f95600080958197613cb4575b509083929161ffff613c5685613c4f611644613c9f99613ca59d9e516129a1565b94516129a1565b5191613c72613c63610259565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261173e83836129a1565b50613948565b99613948565b96019596613b23565b613ca59750611644965084939291509361ffff613c5682613c4f613cf1613c9f9960603d8111613d07575b613ce98183610214565b8101906137ba565b9c9196909c9d5050505050505090919293613c2e565b503d613cdf565b6105a788613d206116448c8f516129a1565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613d7a915060203d81116117ec576117de8183610214565b38613ba9565b50919a9496929395509897968a613d978187611db5565b90506140ab575b50508651613dab90613a08565b99613db96020860186611deb565b91613dc5915086611db5565b9560609150019486613dd687612129565b91613de1938a61594e565b613deb8b896129a1565b52613df68a886129a1565b50613e018a886129a1565b516020015163ffffffff16613e1591613948565b90613e208a886129a1565b516040015163ffffffff16613e3491613948565b91613e3d610259565b33815290600060208301819052604083015261ffff166060820152613e60610293565b60808201528651613e7090613a35565b90613e7b82896129a1565b52613e8690876129a1565b506002546001600160a01b031692613e9d90612129565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c635760009660009760009460009261406e575b506000965b8651881015613fee57613f77600191613f3f88613f3a87612285565b613aa9565b613f586060613f4e8d8d6129a1565b510191825161232d565b9052858a14613f7f575b6060613f6e8b8b6129a1565b510151906132b8565b970196613f1e565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613fb260808c01516001600160a01b031690565b1603613fc0575b5050613f62565b613f3a613fcc926122b5565b613fe56060613fdb8d8d6129a1565b51019182516132b8565b90528b88613fb9565b979650975050505061402a7f000000000000000000000000000000000000000000000000000000000000000091613f3a63ffffffff84166122b5565b84116140365750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b929850505061409691925060803d6080116140a4575b61408e8183610214565b810190613a7e565b919790939290919038613f19565b503d614084565b610b276103f86104f86104f26140c4948a989698611db5565b926001600160a01b03600091515194169060e088019081516140e4610259565b6001600160a01b0385168152908260208301528260408301528260608301526080820152614112878d6129a1565b5261411d868c6129a1565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610c63578f948c89968f96948d948f968891614432575b5061432b575b505050505050156141d2575b6118fe6141c36141ca956141bd60206118fe6141bd976040976129a1565b90613948565b958b6129a1565b90388a613d9e565b50506142589160608c6142036104f86104f26141fc6103f86103f86002546001600160a01b031690565b938b611db5565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c63576118fe6141c36040926141bd60206118fe8f8b906141ca9b6141bd9a6000806000926142ed575b63ffffffff9293506142da9060606142a288886129a1565b5101946142cf8a6142b38a8a6129a1565b51019160406142c28b8b6129a1565b51019063ffffffff169052565b9063ffffffff169052565b169052975097505050509550505061419f565b50505063ffffffff6143196142da9260603d606011614324575b6143118183610214565b8101906139f2565b90935091508261428a565b503d614307565b8495985060a0969750614375602061436b6060826143626104f261435b6104f86104f28b6143ae9c9d9e9f611db5565b998d611db5565b01359901612129565b9a015161ffff1690565b905190604051998a97889687967f2c063404000000000000000000000000000000000000000000000000000000008852600488016139aa565b03915afa8015610c63578592828c9391819082946143f6575b506143ea9060606143d888886129a1565b5101926142cf60206142b38a8a6129a1565b5288888f8c8138614193565b9150506143ea9250614420915060a03d60a01161442b575b6144188183610214565b810190613962565b9491929190506143c7565b503d61440e565b61444b915060203d602011611b5657611b488183610214565b3861418d565b6001600160a01b0360015416330361446557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161449d8151846132b8565b9283156145965760005b8481106144b5575050505050565b8181101561457b576144ca61164482866129a1565b6001600160a01b0381168015610c68576144e383613280565b8781106144f5575050506001016144a7565b84811015614558576001600160a01b03614512611644838a6129a1565b168214614521576001016144e3565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036145766116446145708885613a71565b896129a1565b614512565b61459161164461458b8484613a71565b856129a1565b6144ca565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156121245760051b0190565b9081602091031261016a575190565b91608061024792949361462f8160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610daa9281815201906102cb565b9061467282610277565b61467f6040519182610214565b828152601f196120e28294610277565b908151602082108061470f575b6146de57036146a85790565b610a24906040519182917f3aeba39000000000000000000000000000000000000000000000000000000000835260048301614657565b5090602081015190816146f0846122e7565b1c6146a857506146ff82614668565b9160200360031b1b602082015290565b506020811461469c565b918251601481029080820460141490151715611bc75761473b6147409161328e565b61329c565b9061475261474d836132aa565b614668565b90601461475e83612994565b5360009260215b8651851015614790576014600191614780611644888b6129a1565b60601b8187015201940193614765565b919550936020935060601b90820152828152012090565b906147b76103f860608401612129565b6147c8600019936040810190611db5565b9050614840575b6147d98251613a35565b9260005b8481106147eb575050505050565b80826001921461483b57606061480182876129a1565b51015180156148355761482f9061482961481b84896129a1565b51516001600160a01b031690565b86615b57565b016147dd565b5061482f565b61482f565b915061484c8151613a44565b9161485a61481b84846129a1565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610c63576000916148f3575b506148d3575b506147cf565b6148ed9060606148e386866129a1565b5101519083615b57565b386148cd565b61490c915060203d602011611b5657611b488183610214565b386148c7565b6040519061491f826101f8565b60606020838281520152565b919060408382031261016a5760405190614944826101f8565b8193805167ffffffffffffffff811161016a5782614963918301612a53565b835260208101519167ffffffffffffffff831161016a576020926149879201612a53565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610daa920161492b565b9060806001600160a01b03816149d1855160a0865260a08601906102cb565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610daa9281815201906149b2565b919060408382031261016a57825167ffffffffffffffff811161016a57602091614a4191850161492b565b92015190565b61ffff614a60610daa95936060845260608401906149b2565b9316602082015260408184039101526102cb565b909193929594614a826128b3565b5060208201805115614f1c57614aa5610b276103f885516001600160a01b031690565b946001600160a01b038616916040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180614b0c60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c6357600091614efd575b5015614eb457614b9888999a825192614b37614912565b5051614b83614b4d89516001600160a01b031690565b926040614b58610259565b9e8f908152614b748d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6357600091614e95575b5015614d9d5750906000929183614c449899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614a47565b03925af18015610c6357600094600091614d55575b50614d2a614c939596614ccd614ca1614cbf956116446020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b610214565b6040519586918683019190916001600160a01b036020820193169052565b03601f198101865285610214565b614d0961096d614cff614d0f8951614d0961096d614cff8c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9061468f565b9767ffffffffffffffff166000526004602052604060002090565b93015193614d36610268565b958652602086015260408501526060840152608083015260a082015290565b614c9395506020915061164496614ccd614ca1614cbf95614d8b614d2a953d806000833e614d838183610214565b810190614a16565b9b909b96505095505050969550614c59565b9793929061ffff16614e6b5751614e4157614dec6000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614a05565b03925af1908115610c6357614c9395614ccd614ca1614d2a93611644602096614cbf98600091614e1e575b5099614c75565b614e3b91503d806000833e614e338183610214565b81019061498c565b38614e17565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614eae915060203d602011611b5657611b488183610214565b38614bfc565b6105a7614ec886516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614f16915060203d602011611b5657611b488183610214565b38614b20565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592615018947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90615034602092828151948592016102a8565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b16815261508082518093602089850191016102a8565b019160f81b168382015261509e8251809360206002850191016102a8565b01019160f81b16838201526150bd8251809360206002850191016102a8565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610daa9f9e9c9860f81b16815261513282518093602089850191016102a8565b019160f01b16838201526151508251809360206003850191016102a8565b01019160f01b168382015261516e82518093602089850191016102a8565b01019160f01b1660028201520190615021565b60e081019060ff8251511161542857610100810160ff815151116154125761012082019260ff845151116153fc57610140830160ff815151116153e65761016084019061ffff825151116153d0576101808501936001855151116153b8576101a086019361ffff855151116153a257606095518051615366575b50865167ffffffffffffffff16602088015167ffffffffffffffff1690604089015161522e9067ffffffffffffffff1690565b9860608101516152419063ffffffff1690565b9060808101516152549063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661527797614f46565b03601f19810188526152899088610214565b519081516152979060ff1690565b9051805160ff1698519081516152ad9060ff1690565b906040519a8b9560208701956152c296615038565b03601f19810187526152d49087610214565b519182516152e29060ff1690565b9151805161ffff169480516152f89061ffff1690565b92519283516153089061ffff1690565b92604051978897602089019761531d986150c3565b03601f198101825261532f9082610214565b6040519283926020840161534291615021565b61534b91615021565b61535491615021565b03601f1981018252610daa9082610214565b61537b91965061537590612994565b51615d4e565b9461ffff86511161538c57386151fb565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b6000526105a76024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b90615447612e1c565b916011821061565d5780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216036155ea5750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a6154c6816120bb565b604086019081526154d6826129e9565b906060870191825260005b83811061559e5750505050615549838361553f61553361552961552261550f615553988761555d9c9b615e75565b6001600160a01b0390911660808d015290565b8585615f4b565b929190369161287c565b60a08a01528383615fb3565b949190369161287c565b60c0880152615f4b565b939190369161287c565b60e0840152810361556c575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b806001916155e36155cd6155c66155b96155dd9a8d8d615e75565b9190613450868a516129a1565b8b8b615f4b565b9391889a919a51949a369161287c565b926129a1565b52016154e1565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b80519060005b8281106156a357505050565b60018101808211611bc7575b8381106156bf5750600101615697565b6001600160a01b036156d183856129a1565b51166156e36103f861164484876129a1565b146156f0576001016156af565b6105a761570061164484866129a1565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b909391928151156158c1575061574a81615691565b80519260005b84811061575f57509093925050565b61576f6103f861164483866129a1565b1561577c57600101615750565b93919461579661340461578e85613a35565b8451906132b8565b926157b36157ae6157a683613a35565b8551906132b8565b6129e9565b968792600097885b8481106158635750505050505060005b8151811015615856576000805b868110615817575b50906001916158125761580c6157f961164483866129a1565b613450615805896132c5565b98876129a1565b016157cb565b61580c565b61582461164482876129a1565b6001600160a01b0361583c6103f861164487896129a1565b91161461584b576001016157d8565b5060019050806157e0565b5050909180825283529190565b9091929394988282146158b757906158a961589c8361588f8b6134506001976155dd611644898e6129a1565b6158a26135d985896129a1565b9e6129a1565b528c6129a1565b505b019089949392916157bb565b98506001906158ab565b91935050156158dc57506158d361209f565b90610daa6129b5565b90610daa82516129e9565b9081602091031261016a5751610daa81612236565b9361593960809461ffff6001600160a01b039567ffffffffffffffff615947969b9a9b16895216602088015260a0604088015260a0870190610c92565b9085820360608701526102cb565b9416910152565b9291909261595a61372c565b506159798167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff16916080850192835161599c906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff166159bc91613948565b966159c890608d6132b8565b9460a08701958651516159da916132b8565b9160ff16916159e8836122fd565b6159f1916132b8565b916159fd9060676132b8565b615a069161232d565b615a0f916132b8565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b03871603615a9c5750505061ffff9250615a8e90615a816000935b5195615a74615a65610259565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b615ab26103f8602094976001600160a01b031690565b906040615ac38583015161ffff1690565b91015192615b03875198604051998a96879586957ff2388958000000000000000000000000000000000000000000000000000000008752600487016158fc565b03915afa908115610c6357615a81615a8e9261ffff95600091615b28575b5093615a58565b615b4a915060203d602011615b50575b615b428183610214565b8101906158e7565b38615b21565b503d615b38565b91602091600091604051906001600160a01b03858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152615bab606482610214565b519082855af11561279d576000513d615c0a57506001600160a01b0381163b155b615bd35750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615bcc565b966002615d2297615cef6022610daa9f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615cef9f82615cef9c615cf69c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615ca082518093602089850191016102a8565b019160f81b1683820152615cbe8251809360206023850191016102a8565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190615021565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615e5f57604081019160ff83515111615e4957606082019160ff83515111615e3357608081019260ff84515111615e1d5760a0820161ffff81515111615e0757610daa94615df99351945191615db0835160ff1690565b975191615dbe835160ff1690565b945190615dcc825160ff1690565b905193615dda855160ff1690565b935196615de9885161ffff1690565b966040519c8d9b60208d01615c13565b03601f198101835282610214565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b929190926001820191848311615f195781013560001a828115615f0e575060148103615ee1578201938411615ead57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b91906002820191818311615f19578381013560f01c0160020192818411615f7f57918391615f7893612e7b565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b91906001820191818311615f19578381013560001a0160010192818411615f7f57918391615f7893612e7b56fea164736f6c634300081a000a",
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
