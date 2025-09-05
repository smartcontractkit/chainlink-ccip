// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp_with_message_transformer

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

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenAmounts   []InternalEVM2AnyTokenTransfer
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type OnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	MessageInterceptor     common.Address
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

var OnRampWithMessageTransformerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"messageTransformerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMessageId\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageTransformer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMessageTransformer\",\"inputs\":[{\"name\":\"messageTransformerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageIdUnexpectedlySet\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MessageTransformError\",\"inputs\":[{\"name\":\"errorReason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346105a1576145088038038061001a816105db565b92833981019080820361016081126105a157608081126105a15761003c6105bc565b9161004681610600565b835260208101516001600160a01b03811681036105a1576020840190815261007060408301610614565b916040850192835260a061008660608301610614565b6060870190815294607f1901126105a15760405160a081016001600160401b038111828210176105a6576040526100bf60808301610614565b81526100cd60a08301610628565b602082019081526100e060c08401610614565b90604083019182526100f460e08501610614565b92606081019384526101096101008601610614565b608082019081526101208601519095906001600160401b0381116105a15781018b601f820112156105a1578051906001600160401b0382116105a6578160051b602001610155906105db565b9c8d838152602001926060028201602001918183116105a157602001925b828410610533575050505061014061018b9101610614565b98331561052257600180546001600160a01b0319163317905580516001600160401b0316158015610510575b80156104fe575b80156104ec575b6104bf57516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e081905282519091161580156104da575b80156104d0575b6104bf57815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1986101209860606102af6105bc565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a160005b82518110156104185761033a8184610635565b516001600160401b0361034d8386610635565b5151169081156104035760008281526006602090815260409182902081840151815494840151600160401b600160e81b03198616604883901b600160481b600160e81b031617901515851b68ff000000000000000016179182905583516001600160401b0390951685526001600160a01b031691840191909152811c60ff1615159082015260019291907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590606090a201610327565b5063c35aa79d60e01b60005260045260246000fd5b506001600160a01b031680156104ae57600780546001600160a01b031916919091179055604051613ea89081610660823960805181818161071a0152818161218e0152818161297801526135be015260a0518181816121c7015281816126e401526129b1015260c051818181610f680152818161220301526129ed015260e05181818161223f01528181612a29015261307a0152f35b6342bcdf7f60e11b60005260046000fd5b6306b7c75960e31b60005260046000fd5b5082511515610210565b5084516001600160a01b031615610209565b5088516001600160a01b0316156101c5565b5087516001600160a01b0316156101be565b5086516001600160a01b0316156101b7565b639b15e16f60e01b60005260046000fd5b6060848303126105a15760405190606082016001600160401b038111838210176105a65760405261056385610600565b82526020850151906001600160a01b03821682036105a1578260209283606095015261059160408801610628565b6040820152815201930192610173565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b038111838210176105a657604052565b6040519190601f01601f191682016001600160401b038111838210176105a657604052565b51906001600160401b03821682036105a157565b51906001600160a01b03821682036105a157565b519081151582036105a157565b80518210156106495760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c90816306285c69146129125750806315777ab2146128c0578063181f5a771461284157806320487ded146126085780632716072b1461235857806327e936f114611f5257806348a98aa414611ecf5780635cb80c5d14611c1257806365b81aab14611b625780636def4ce714611ad35780637437ff9f146119b657806379ba5097146118d15780638da5cb5b1461187f5780639041be3d146117d2578063972b461214611704578063c9b146b31461133f578063df0aa9e914610563578063f2890a211461024d578063f2fde38b146101605763fbca3b74146100fd57600080fd5b3461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57600490610137612b03565b507f9e7177c8000000000000000000000000000000000000000000000000000000008152fd5b80fd5b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5773ffffffffffffffffffffffffffffffffffffffff6101ad612c3a565b6101b561376d565b1633811461022557807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d576004359067ffffffffffffffff821161015d577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301916101a0831261055b5760a0604051936102ce85612b9b565b1261055b576040516102df81612b7f565b816004013581526102f260248301612b1f565b602082015261030360448301612b1f565b604082015261031460648301612b1f565b606082015261032560848301612b1f565b6080820152835261033860a48201612c80565b602084015260c481013567ffffffffffffffff811161055f576103619060043691840101612d8d565b604084015260e481013567ffffffffffffffff811161055f5761038a9060043691840101612d8d565b606084015261010481013567ffffffffffffffff811161055f576103b49060043691840101612d8d565b60808401526103c66101248201612c80565b60a084015261014481013560c084015261016481013560e08401526101848101359067ffffffffffffffff821161055f57013660238201121561055b5760048101359061041282612c15565b926104206040519485612bd4565b828452602060048186019460051b840101019136831161055b5760248101935b838510610461576020610459888861010082015261356f565b604051908152f35b843567ffffffffffffffff81116105575760049083010160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0823603011261055757604051906104b182612b7f565b6104bd60208201612c80565b8252604081013567ffffffffffffffff8111610553576104e39060203691840101612d8d565b6020830152606081013567ffffffffffffffff81116105535761050c9060203691840101612d8d565b60408301526080810135606083015260a081013567ffffffffffffffff811161055357916105436020949285809536920101612d8d565b6080820152815201940193610440565b8580fd5b8380fd5b5080fd5b8280fd5b503461015d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5761059b612b03565b67ffffffffffffffff6024351161055b5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6024353603011261055b576105e2612c5d565b60025460ff8160a01c16611317577fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001760025567ffffffffffffffff8216835260066020526040832073ffffffffffffffffffffffffffffffffffffffff8216156112ef57805460ff8160401c16611281575b60481c73ffffffffffffffffffffffffffffffffffffffff1633036112595773ffffffffffffffffffffffffffffffffffffffff60035416806111e5575b50805467ffffffffffffffff811667ffffffffffffffff81146111b8579067ffffffffffffffff60017fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000094930116928391161790556040519061070c82612b7f565b84825267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083015267ffffffffffffffff84166040830152606082015283608082015261076d6024803501602435600401613212565b909361077e60046024350180613212565b9290606460243501936107908561311e565b976107a5604460243501602435600401613263565b9490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06107eb6107d587612c15565b966107e36040519889612bd4565b808852612c15565b018b5b8181106111a157505061083f939291610833916040519961080e8b612b9b565b8a5273ffffffffffffffffffffffffffffffffffffffff8b1660208b01523691612d56565b60408801523691612d56565b606085015273ffffffffffffffffffffffffffffffffffffffff87602094826040519961086c888c612bd4565b838b52608089019a8b521660a088015260443560c088015260e08701938285526101008801526108a083600254169161311e565b9067ffffffffffffffff8661095a6108c2608460243501602435600401613212565b9061092a6108d560046024350180613212565b9290936040519b8c9a8b998a997f3a49bb49000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604435604488015260a0606488015260a4870191612e08565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc858403016084860152612e08565b03915afa91821561119657889189988a918b9561111b575b505252610989604460243501602435600401613263565b9661099388612c15565b916109a16040519384612bd4565b888352858301809960061b82019136831161111757905b8282106110e057505050885b6109d8604460243501602435600401613263565b9050811015610d40576109eb81846131cf565b51906109f56132b7565b508682015115610d185773ffffffffffffffffffffffffffffffffffffffff610a208184511661301b565b169182158015610c75575b610c3357808c878a8a73ffffffffffffffffffffffffffffffffffffffff8f83610ad898015182808951169260405197610a6489612b7f565b885267ffffffffffffffff87890196168652816040890191168152606088019283526080880193845267ffffffffffffffff6040519b8c998a997f9a4575b9000000000000000000000000000000000000000000000000000000008b5260048b01525160a060248b015260c48a0190612ac0565b965116604488015251166064860152516084850152511660a4830152038183885af1918215610c28578d8093610b70575b50506001938284928b80610b6996519301519101519160405193610b2c85612b7f565b84528c840152604083015260608201528d60405190610b4b8c83612bd4565b815260808201526101008b015190610b6383836131cf565b526131cf565b50016109c4565b9250933d8093863e610b828386612bd4565b89858481010312610c245784519167ffffffffffffffff8311610c205760408387018588010312610c2057604051610bb981612bb8565b8387015167ffffffffffffffff8111610c1b57610bdd9086890190868a01016132e3565b81528b84880101519467ffffffffffffffff8611610c1b5787610b69968896610c0c9360019b019201016132e3565b8c82015293509150938d610b09565b508f80fd5b8e80fd5b8d80fd5b6040513d8f823e3d90fd5b517fbf16aab6000000000000000000000000000000000000000000000000000000008c5273ffffffffffffffffffffffffffffffffffffffff1660045260248bfd5b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201528881602481875afa908115610c28578d91610cdf575b5015610a2b565b90508881813d8311610d11575b610cf68183612bd4565b81010312610d0d57610d0790612dab565b38610cd8565b8c80fd5b503d610cec565b60048b7f5cf04449000000000000000000000000000000000000000000000000000000008152fd5b5090838994979350849167ffffffffffffffff9973ffffffffffffffffffffffffffffffffffffffff600254169188610db86101008c01516040519e8f977f01447eaa000000000000000000000000000000000000000000000000000000008952166004880152606060248801526064870190613325565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8684030160448701525191828152019190855b8a8282106110a5575050505082809103915afa96871561109a578397610fde575b5015610ef95750805b67ffffffffffffffff6080855101911690525b61010083015151811015610e605780610e45600192876131cf565b516080610e57836101008801516131cf565b51015201610e2a565b5082610e6b836137e8565b90610e758261356f565b82515267ffffffffffffffff60608351015116907f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f3267ffffffffffffffff60405192169180610ec48682613410565b0390a37fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff600254166002555151604051908152f35b73ffffffffffffffffffffffffffffffffffffffff604051917fea458c0c00000000000000000000000000000000000000000000000000000000835267ffffffffffffffff8716600484015216602482015282816044818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610fd3578291610f9e575b50610e17565b90508281813d8311610fcc575b610fb58183612bd4565b8101031261055b57610fc6906133fb565b86610f98565b503d610fab565b6040513d84823e3d90fd5b9096503d8084833e610ff08183612bd4565b81019084818303126105575780519067ffffffffffffffff8211611096570181601f8201121561055757805161102581612c15565b926110336040519485612bd4565b818452868085019260051b8401019281841161109257878101925b8484106110615750505050509587610e0e565b835167ffffffffffffffff811161108e578991611083858480948701016132e3565b81520193019261104e565b8880fd5b8680fd5b8480fd5b6040513d85823e3d90fd5b8351805173ffffffffffffffffffffffffffffffffffffffff168652810151818601528a97508e965060409094019390920191600101610ded565b604082360312611117578760409182516110f981612bb8565b61110285612c80565b815282850135838201528152019101906109b8565b8b80fd5b935093505096503d8089833e6111318183612bd4565b81019660808289031261108e57815161114b868401612dab565b92604081015167ffffffffffffffff8111611117578a61116c9183016132e3565b99606082015167ffffffffffffffff8111610d0d5761118b92016132e3565b909298909338610972565b6040513d8a823e3d90fd5b6020906111ac6132b7565b82828a010152016107ee565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b803b15611096578460405180927fe0a0e50600000000000000000000000000000000000000000000000000000000825281838161122b6024356004018b60048401612e47565b03925af1801561124e57156106aa578461124791959295612bd4565b92386106aa565b6040513d87823e3d90fd5b6004847f1c0a3529000000000000000000000000000000000000000000000000000000008152fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260028301602052604090205461066c5760248573ffffffffffffffffffffffffffffffffffffffff857fd0d2597600000000000000000000000000000000000000000000000000000000835216600452fd5b6004847fa4ec7479000000000000000000000000000000000000000000000000000000008152fd5b6004847f3ee5aeb5000000000000000000000000000000000000000000000000000000008152fd5b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5760043567ffffffffffffffff811161055b5761138f903690600401612ca1565b73ffffffffffffffffffffffffffffffffffffffff6001541633036116bc575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b848110156116b8578060051b8201358381121561109657820191608083360312611096576040519461140b86612b34565b61141484612b1f565b865261142260208501612c2d565b9660208701978852604085013567ffffffffffffffff811161055f5761144b903690870161316a565b9460408801958652606081013567ffffffffffffffff8111610557576114739136910161316a565b60608801908152875167ffffffffffffffff1683526006602052604080842099518a547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff169015159182901b68ff000000000000000016178a55909590815151611590575b5095976001019550815b85518051821015611521579061151a73ffffffffffffffffffffffffffffffffffffffff611512836001956131cf565b511689613c3d565b50016114e2565b50509590969450600192919351908151611541575b5050019392936113da565b61158667ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190612cd2565b0390a23880611536565b9893959296919094979860001461168157600184019591875b86518051821015611626576115d38273ffffffffffffffffffffffffffffffffffffffff926131cf565b511680156115ef57906115e86001928a613bac565b50016115a9565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161167767ffffffffffffffff8a51169251604051918291602083526020830190612cd2565b0390a238806114d8565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff600554163303156113af576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5767ffffffffffffffff611745612b03565b16808252600660205260ff604083205460401c16908252600660205260016040832001916040518093849160208254918281520191845260208420935b8181106117b957505061179792500383612bd4565b6117b560405192839215158352604060208401526040830190612cd2565b0390f35b8454835260019485019487945060209093019201611782565b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5767ffffffffffffffff611813612b03565b1681526006602052600167ffffffffffffffff604083205416019067ffffffffffffffff82116118525760208267ffffffffffffffff60405191168152f35b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b503461015d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461015d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57805473ffffffffffffffffffffffffffffffffffffffff8116330361198e577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461015d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d576119ed61313f565b5060a06040516119fc81612b7f565b60ff60025473ffffffffffffffffffffffffffffffffffffffff81168352831c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015273ffffffffffffffffffffffffffffffffffffffff60045416606082015273ffffffffffffffffffffffffffffffffffffffff600554166080820152611ad1604051809273ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565bf35b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57604060609167ffffffffffffffff611b19612b03565b1681526006602052205473ffffffffffffffffffffffffffffffffffffffff6040519167ffffffffffffffff8116835260ff8160401c161515602084015260481c166040820152f35b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5773ffffffffffffffffffffffffffffffffffffffff611baf612c3a565b611bb761376d565b168015611bea577fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075580f35b6004827f8579befe000000000000000000000000000000000000000000000000000000008152fd5b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d5760043567ffffffffffffffff811161055b57611c62903690600401612ca1565b9073ffffffffffffffffffffffffffffffffffffffff6004541690835b83811015611ecb5773ffffffffffffffffffffffffffffffffffffffff611caa8260051b840161311e565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315611ec0578793611e8d575b5082611d02575b506001915001611c7f565b8460405193611d9e60208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611d4d606482612bd4565b8a80604098895193611d5f8b86612bd4565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af1611d976137b8565b9086613dcf565b805180611dda575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338611cf7565b81929495969350906020918101031261108e576020611df99101612dab565b15611e0a5792919085903880611da6565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9092506020813d8211611eb8575b81611ea860209383612bd4565b8101031261109257519138611cf0565b3d9150611e9b565b6040513d89823e3d90fd5b8480f35b503461015d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57611f07612b03565b506024359073ffffffffffffffffffffffffffffffffffffffff8216820361015d576020611f348361301b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b503461015d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57604051611f8e81612b7f565b611f96612c3a565b8152602435801515810361055f576020820190815260443573ffffffffffffffffffffffffffffffffffffffff811681036105575760408301908152611fda612c5d565b90606084019182526084359273ffffffffffffffffffffffffffffffffffffffff84168403610553576080850193845261201261376d565b73ffffffffffffffffffffffffffffffffffffffff855116158015612339575b801561232f575b612307579273ffffffffffffffffffffffffffffffffffffffff859381809461012097827fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19a51167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff000000000000000000000000000000000000000060055416176005556123036040519161218383612b34565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602084015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604084015273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660608401526122b3604051809473ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b608083019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565ba180f35b6004867f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515612039565b5073ffffffffffffffffffffffffffffffffffffffff83511615612032565b503461015d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d576004359067ffffffffffffffff821161015d573660238301121561015d5781600401356123b481612c15565b926123c26040519485612bd4565b818452602460606020860193028201019036821161055757602401915b818310612560575050506123f161376d565b805b825181101561255c5761240681846131cf565b5167ffffffffffffffff61241a83866131cf565b51511690811561253057907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5606060019493838752600660205260ff604088206124f1604060208501519483547fffffff0000000000000000000000000000000000000000ffffffffffffffffff7cffffffffffffffffffffffffffffffffffffffff0000000000000000008860481b1691161784550151151582907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff0000000000000000835492151560401b169116179055565b5473ffffffffffffffffffffffffffffffffffffffff6040519367ffffffffffffffff8316855216602084015260401c1615156040820152a2016123f3565b602484837fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b606083360312610557576040516060810181811067ffffffffffffffff8211176125db5760405261259084612b1f565b8152602084013573ffffffffffffffffffffffffffffffffffffffff811681036105535791816060936020809401526125cb60408701612c2d565b60408201528152019201916123df565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461015d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57612640612b03565b60243567ffffffffffffffff811161055f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261055f576040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008360801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156128365784916127fc575b506127c6576127739160209173ffffffffffffffffffffffffffffffffffffffff60025416906040518095819482937fd8694ccd0000000000000000000000000000000000000000000000000000000084526004019060048401612e47565b03915afa908115610fd3578291612790575b602082604051908152f35b90506020813d6020116127be575b816127ab60209383612bd4565b8101031261055b57602091505138612785565b3d915061279e565b60248367ffffffffffffffff847ffdbd6a7200000000000000000000000000000000000000000000000000000000835216600452fd5b90506020813d60201161282e575b8161281760209383612bd4565b810103126105575761282890612dab565b38612714565b3d915061280a565b6040513d86823e3d90fd5b503461015d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57506117b5604051612882604082612bd4565b601081527f4f6e52616d7020312e362e322d646576000000000000000000000000000000006020820152604051918291602083526020830190612ac0565b503461015d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015d57602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b90503461055b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261055b578061294e606092612b34565b8281528260208201528260408201520152608060405161296d81612b34565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152611ad1604051809273ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b60005b838110612ab05750506000910152565b8181015183820152602001612aa0565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093612afc81518092818752878088019101612a9d565b0116010190565b6004359067ffffffffffffffff82168203612b1a57565b600080fd5b359067ffffffffffffffff82168203612b1a57565b6080810190811067ffffffffffffffff821117612b5057604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff821117612b5057604052565b610120810190811067ffffffffffffffff821117612b5057604052565b6040810190811067ffffffffffffffff821117612b5057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612b5057604052565b67ffffffffffffffff8111612b505760051b60200190565b35908115158203612b1a57565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203612b1a57565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203612b1a57565b359073ffffffffffffffffffffffffffffffffffffffff82168203612b1a57565b9181601f84011215612b1a5782359167ffffffffffffffff8311612b1a576020808501948460051b010111612b1a57565b906020808351928381520192019060005b818110612cf05750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612ce3565b67ffffffffffffffff8111612b5057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612d6282612d1c565b91612d706040519384612bd4565b829481845281830111612b1a578281602093846000960137010152565b9080601f83011215612b1a57816020612da893359101612d56565b90565b51908115158203612b1a57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215612b1a57016020813591019167ffffffffffffffff8211612b1a578136038313612b1a57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9067ffffffffffffffff9093929316815260406020820152612ebd612e80612e6f8580612db8565b60a0604086015260e0850191612e08565b612e8d6020860186612db8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0858403016060860152612e08565b9060408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215612b1a5784016020813591019267ffffffffffffffff8211612b1a578160061b36038413612b1a578281037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0016080840152818152602001929060005b818110612fbb57505050612f8b8473ffffffffffffffffffffffffffffffffffffffff612f7b6060612da8979801612c80565b1660a08401526080810190612db8565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152612e08565b90919360408060019273ffffffffffffffffffffffffffffffffffffffff612fe289612c80565b16815260208881013590820152019501929101612f48565b519073ffffffffffffffffffffffffffffffffffffffff82168203612b1a57565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015613112576000906130c5575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d60201161310a575b816130df60209383612bd4565b81010312612b1a5761310573ffffffffffffffffffffffffffffffffffffffff91612ffa565b6130aa565b3d91506130d2565b6040513d6000823e3d90fd5b3573ffffffffffffffffffffffffffffffffffffffff81168103612b1a5790565b6040519061314c82612b7f565b60006080838281528260208201528260408201528260608201520152565b9080601f83011215612b1a57813561318181612c15565b9261318f6040519485612bd4565b81845260208085019260051b820101928311612b1a57602001905b8282106131b75750505090565b602080916131c484612c80565b8152019101906131aa565b80518210156131e35760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612b1a570180359067ffffffffffffffff8211612b1a57602001918136038313612b1a57565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612b1a570180359067ffffffffffffffff8211612b1a57602001918160061b36038313612b1a57565b604051906132c482612b7f565b6060608083600081528260208201528260408201526000838201520152565b81601f82011215612b1a5780516132f981612d1c565b926133076040519485612bd4565b81845260208284010111612b1a57612da89160208085019101612a9d565b9080602083519182815201916020808360051b8301019401926000915b83831061335157505050505090565b90919293946020806133ec837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806133d16133bf8685015160a08886015260a0850190612ac0565b60408501518482036040860152612ac0565b92606081015160608401520151906080818403910152612ac0565b97019301930191939290613342565b519067ffffffffffffffff82168203612b1a57565b90612da8916020815267ffffffffffffffff6080835180516020850152826020820151166040850152826040820151166060850152826060820151168285015201511660a082015273ffffffffffffffffffffffffffffffffffffffff60208301511660c08201526101006135046134cf61349c60408601516101a060e08701526101c0860190612ac0565b60608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030185870152612ac0565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301610120860152612ac0565b9273ffffffffffffffffffffffffffffffffffffffff60a08201511661014084015260c081015161016084015260e08101516101808401520151906101a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613325565b80515180613740575067ffffffffffffffff6040825101511660405160208101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321835267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408301526060820152306080820152608081526135fd60a082612bd4565b5190209073ffffffffffffffffffffffffffffffffffffffff602082015116815167ffffffffffffffff6080816060840151169201511673ffffffffffffffffffffffffffffffffffffffff60a08501511660c085015191604051936020850195865260408501526060840152608083015260a082015260a0815261368360c082612bd4565b519020906060810151602081519101209060408101516020815191012060806101008301516040516136f1816136c56020820194602086526040830190613325565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612bd4565b51902092015160208151910120926040519460208601966000885260408701526060860152608085015260a084015260c083015260e082015260e0815261373a61010082612bd4565b51902090565b7f4c8ebcc00000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361378e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b3d156137e3573d906137c982612d1c565b916137d76040519384612bd4565b82523d6000602084013e565b606090565b600061388981926040516137fb81612b9b565b61380361313f565b815283602082015260606040820152606080820152606060808201528360a08201528360c08201528360e082015260606101008201525073ffffffffffffffffffffffffffffffffffffffff60075416906040519485809481937f8a06fadb00000000000000000000000000000000000000000000000000000000835260048301613410565b03925af180916000916138e6575b5090612da8576138e26138a86137b8565b6040519182917f828ebdfb000000000000000000000000000000000000000000000000000000008352602060048401526024830190612ac0565b0390fd5b3d8083833e6138f58183612bd4565b81019060208183031261055f5780519067ffffffffffffffff8211610557570191828203926101a0841261055b5760a06040519461393286612b9b565b1261055b5760405161394381612b7f565b81518152613953602083016133fb565b6020820152613964604083016133fb565b6040820152613975606083016133fb565b6060820152613986608083016133fb565b6080820152845261399960a08201612ffa565b602085015260c081015167ffffffffffffffff811161055f57836139be9183016132e3565b604085015260e081015167ffffffffffffffff811161055f57836139e39183016132e3565b606085015261010081015167ffffffffffffffff811161055f5783613a099183016132e3565b6080850152613a1b6101208201612ffa565b60a085015261014081015160c085015261016081015160e08501526101808101519067ffffffffffffffff821161055f570182601f8201121561055b57805191613a6483612c15565b93613a726040519586612bd4565b83855260208086019460051b8401019281841161055f5760208101945b848610613aa85750505050505061010082015238613897565b855167ffffffffffffffff811161109657820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126110965760405190613af482612b7f565b613b0060208201612ffa565b8252604081015167ffffffffffffffff811161109257856020613b25928401016132e3565b6020830152606081015167ffffffffffffffff811161109257856020613b4d928401016132e3565b60408301526080810151606083015260a08101519067ffffffffffffffff82116110925791613b84866020809694819601016132e3565b6080820152815201950194613a8f565b80548210156131e35760005260206000200190600090565b6000828152600182016020526040902054613c365780549068010000000000000000821015612b505782613c1f613bea846001809601855584613b94565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014613dc6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613d97578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613d9757818103613d60575b50505080548015613d31577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190613cf28282613b94565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b613d80613d70613bea9386613b94565b90549060031b1c92839286613b94565b905560005283602052604060002055388080613cba565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b91929015613e4a5750815115613de3575090565b3b15613dec5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015613e5d5750805190602001fd5b6138e2906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612ac056fea164736f6c634300081a000a",
}

var OnRampWithMessageTransformerABI = OnRampWithMessageTransformerMetaData.ABI

var OnRampWithMessageTransformerBin = OnRampWithMessageTransformerMetaData.Bin

func DeployOnRampWithMessageTransformer(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig, destChainConfigs []OnRampDestChainConfigArgs, messageTransformerAddr common.Address) (common.Address, *types.Transaction, *OnRampWithMessageTransformer, error) {
	parsed, err := OnRampWithMessageTransformerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampWithMessageTransformerBin), backend, staticConfig, dynamicConfig, destChainConfigs, messageTransformerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRampWithMessageTransformer{address: address, abi: *parsed, OnRampWithMessageTransformerCaller: OnRampWithMessageTransformerCaller{contract: contract}, OnRampWithMessageTransformerTransactor: OnRampWithMessageTransformerTransactor{contract: contract}, OnRampWithMessageTransformerFilterer: OnRampWithMessageTransformerFilterer{contract: contract}}, nil
}

type OnRampWithMessageTransformer struct {
	address common.Address
	abi     abi.ABI
	OnRampWithMessageTransformerCaller
	OnRampWithMessageTransformerTransactor
	OnRampWithMessageTransformerFilterer
}

type OnRampWithMessageTransformerCaller struct {
	contract *bind.BoundContract
}

type OnRampWithMessageTransformerTransactor struct {
	contract *bind.BoundContract
}

type OnRampWithMessageTransformerFilterer struct {
	contract *bind.BoundContract
}

type OnRampWithMessageTransformerSession struct {
	Contract     *OnRampWithMessageTransformer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampWithMessageTransformerCallerSession struct {
	Contract *OnRampWithMessageTransformerCaller
	CallOpts bind.CallOpts
}

type OnRampWithMessageTransformerTransactorSession struct {
	Contract     *OnRampWithMessageTransformerTransactor
	TransactOpts bind.TransactOpts
}

type OnRampWithMessageTransformerRaw struct {
	Contract *OnRampWithMessageTransformer
}

type OnRampWithMessageTransformerCallerRaw struct {
	Contract *OnRampWithMessageTransformerCaller
}

type OnRampWithMessageTransformerTransactorRaw struct {
	Contract *OnRampWithMessageTransformerTransactor
}

func NewOnRampWithMessageTransformer(address common.Address, backend bind.ContractBackend) (*OnRampWithMessageTransformer, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampWithMessageTransformerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRampWithMessageTransformer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformer{address: address, abi: abi, OnRampWithMessageTransformerCaller: OnRampWithMessageTransformerCaller{contract: contract}, OnRampWithMessageTransformerTransactor: OnRampWithMessageTransformerTransactor{contract: contract}, OnRampWithMessageTransformerFilterer: OnRampWithMessageTransformerFilterer{contract: contract}}, nil
}

func NewOnRampWithMessageTransformerCaller(address common.Address, caller bind.ContractCaller) (*OnRampWithMessageTransformerCaller, error) {
	contract, err := bindOnRampWithMessageTransformer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerCaller{contract: contract}, nil
}

func NewOnRampWithMessageTransformerTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampWithMessageTransformerTransactor, error) {
	contract, err := bindOnRampWithMessageTransformer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerTransactor{contract: contract}, nil
}

func NewOnRampWithMessageTransformerFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampWithMessageTransformerFilterer, error) {
	contract, err := bindOnRampWithMessageTransformer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerFilterer{contract: contract}, nil
}

func bindOnRampWithMessageTransformer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnRampWithMessageTransformerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampWithMessageTransformer.Contract.OnRampWithMessageTransformerCaller.contract.Call(opts, result, method, params...)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.OnRampWithMessageTransformerTransactor.contract.Transfer(opts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.OnRampWithMessageTransformerTransactor.contract.Transact(opts, method, params...)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampWithMessageTransformer.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.contract.Transfer(opts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.contract.Transact(opts, method, params...)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "generateMessageId", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRampWithMessageTransformer.Contract.GenerateMessageId(&_OnRampWithMessageTransformer.CallOpts, message)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRampWithMessageTransformer.Contract.GenerateMessageId(&_OnRampWithMessageTransformer.CallOpts, message)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

	error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	outstruct := new(GetAllowedSendersList)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfiguredAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRampWithMessageTransformer.Contract.GetAllowedSendersList(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRampWithMessageTransformer.Contract.GetAllowedSendersList(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AllowlistEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRampWithMessageTransformer.Contract.GetDestChainConfig(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRampWithMessageTransformer.Contract.GetDestChainConfig(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDynamicConfig)).(*OnRampDynamicConfig)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRampWithMessageTransformer.Contract.GetDynamicConfig(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRampWithMessageTransformer.Contract.GetDynamicConfig(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRampWithMessageTransformer.Contract.GetExpectedNextSequenceNumber(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRampWithMessageTransformer.Contract.GetExpectedNextSequenceNumber(&_OnRampWithMessageTransformer.CallOpts, destChainSelector)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRampWithMessageTransformer.Contract.GetFee(&_OnRampWithMessageTransformer.CallOpts, destChainSelector, message)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRampWithMessageTransformer.Contract.GetFee(&_OnRampWithMessageTransformer.CallOpts, destChainSelector, message)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetMessageTransformer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getMessageTransformer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetMessageTransformer() (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetMessageTransformer(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetMessageTransformer() (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetMessageTransformer(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetPoolBySourceToken(&_OnRampWithMessageTransformer.CallOpts, arg0, sourceToken)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetPoolBySourceToken(&_OnRampWithMessageTransformer.CallOpts, arg0, sourceToken)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampStaticConfig)).(*OnRampStaticConfig)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRampWithMessageTransformer.Contract.GetStaticConfig(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRampWithMessageTransformer.Contract.GetStaticConfig(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetSupportedTokens(&_OnRampWithMessageTransformer.CallOpts, arg0)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.GetSupportedTokens(&_OnRampWithMessageTransformer.CallOpts, arg0)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) Owner() (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.Owner(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) Owner() (common.Address, error) {
	return _OnRampWithMessageTransformer.Contract.Owner(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRampWithMessageTransformer.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) TypeAndVersion() (string, error) {
	return _OnRampWithMessageTransformer.Contract.TypeAndVersion(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerCallerSession) TypeAndVersion() (string, error) {
	return _OnRampWithMessageTransformer.Contract.TypeAndVersion(&_OnRampWithMessageTransformer.CallOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "acceptOwnership")
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.AcceptOwnership(&_OnRampWithMessageTransformer.TransactOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.AcceptOwnership(&_OnRampWithMessageTransformer.TransactOpts)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ApplyAllowlistUpdates(&_OnRampWithMessageTransformer.TransactOpts, allowlistConfigArgsItems)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ApplyAllowlistUpdates(&_OnRampWithMessageTransformer.TransactOpts, allowlistConfigArgsItems)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ApplyDestChainConfigUpdates(&_OnRampWithMessageTransformer.TransactOpts, destChainConfigArgs)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ApplyDestChainConfigUpdates(&_OnRampWithMessageTransformer.TransactOpts, destChainConfigArgs)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ForwardFromRouter(&_OnRampWithMessageTransformer.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.ForwardFromRouter(&_OnRampWithMessageTransformer.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.SetDynamicConfig(&_OnRampWithMessageTransformer.TransactOpts, dynamicConfig)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.SetDynamicConfig(&_OnRampWithMessageTransformer.TransactOpts, dynamicConfig)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) SetMessageTransformer(opts *bind.TransactOpts, messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "setMessageTransformer", messageTransformerAddr)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) SetMessageTransformer(messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.SetMessageTransformer(&_OnRampWithMessageTransformer.TransactOpts, messageTransformerAddr)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) SetMessageTransformer(messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.SetMessageTransformer(&_OnRampWithMessageTransformer.TransactOpts, messageTransformerAddr)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.TransferOwnership(&_OnRampWithMessageTransformer.TransactOpts, to)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.TransferOwnership(&_OnRampWithMessageTransformer.TransactOpts, to)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.WithdrawFeeTokens(&_OnRampWithMessageTransformer.TransactOpts, feeTokens)
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRampWithMessageTransformer.Contract.WithdrawFeeTokens(&_OnRampWithMessageTransformer.TransactOpts, feeTokens)
}

type OnRampWithMessageTransformerAllowListAdminSetIterator struct {
	Event *OnRampWithMessageTransformerAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerAllowListAdminSet)
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
		it.Event = new(OnRampWithMessageTransformerAllowListAdminSet)
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

func (it *OnRampWithMessageTransformerAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerAllowListAdminSet struct {
	AllowlistAdmin common.Address
	Raw            types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampWithMessageTransformerAllowListAdminSetIterator, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerAllowListAdminSetIterator{contract: _OnRampWithMessageTransformer.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerAllowListAdminSet)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseAllowListAdminSet(log types.Log) (*OnRampWithMessageTransformerAllowListAdminSet, error) {
	event := new(OnRampWithMessageTransformerAllowListAdminSet)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerAllowListSendersAddedIterator struct {
	Event *OnRampWithMessageTransformerAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerAllowListSendersAdded)
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
		it.Event = new(OnRampWithMessageTransformerAllowListSendersAdded)
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

func (it *OnRampWithMessageTransformerAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerAllowListSendersAddedIterator{contract: _OnRampWithMessageTransformer.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerAllowListSendersAdded)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseAllowListSendersAdded(log types.Log) (*OnRampWithMessageTransformerAllowListSendersAdded, error) {
	event := new(OnRampWithMessageTransformerAllowListSendersAdded)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerAllowListSendersRemovedIterator struct {
	Event *OnRampWithMessageTransformerAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerAllowListSendersRemoved)
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
		it.Event = new(OnRampWithMessageTransformerAllowListSendersRemoved)
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

func (it *OnRampWithMessageTransformerAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerAllowListSendersRemovedIterator{contract: _OnRampWithMessageTransformer.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerAllowListSendersRemoved)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseAllowListSendersRemoved(log types.Log) (*OnRampWithMessageTransformerAllowListSendersRemoved, error) {
	event := new(OnRampWithMessageTransformerAllowListSendersRemoved)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerCCIPMessageSentIterator struct {
	Event *OnRampWithMessageTransformerCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerCCIPMessageSent)
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
		it.Event = new(OnRampWithMessageTransformerCCIPMessageSent)
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

func (it *OnRampWithMessageTransformerCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampWithMessageTransformerCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerCCIPMessageSentIterator{contract: _OnRampWithMessageTransformer.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerCCIPMessageSent)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseCCIPMessageSent(log types.Log) (*OnRampWithMessageTransformerCCIPMessageSent, error) {
	event := new(OnRampWithMessageTransformerCCIPMessageSent)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerConfigSetIterator struct {
	Event *OnRampWithMessageTransformerConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerConfigSet)
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
		it.Event = new(OnRampWithMessageTransformerConfigSet)
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

func (it *OnRampWithMessageTransformerConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerConfigSet struct {
	StaticConfig  OnRampStaticConfig
	DynamicConfig OnRampDynamicConfig
	Raw           types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OnRampWithMessageTransformerConfigSetIterator, error) {

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerConfigSetIterator{contract: _OnRampWithMessageTransformer.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerConfigSet)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseConfigSet(log types.Log) (*OnRampWithMessageTransformerConfigSet, error) {
	event := new(OnRampWithMessageTransformerConfigSet)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerDestChainConfigSetIterator struct {
	Event *OnRampWithMessageTransformerDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerDestChainConfigSet)
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
		it.Event = new(OnRampWithMessageTransformerDestChainConfigSet)
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

func (it *OnRampWithMessageTransformerDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerDestChainConfigSetIterator{contract: _OnRampWithMessageTransformer.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerDestChainConfigSet)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseDestChainConfigSet(log types.Log) (*OnRampWithMessageTransformerDestChainConfigSet, error) {
	event := new(OnRampWithMessageTransformerDestChainConfigSet)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerFeeTokenWithdrawnIterator struct {
	Event *OnRampWithMessageTransformerFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerFeeTokenWithdrawn)
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
		it.Event = new(OnRampWithMessageTransformerFeeTokenWithdrawn)
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

func (it *OnRampWithMessageTransformerFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampWithMessageTransformerFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerFeeTokenWithdrawnIterator{contract: _OnRampWithMessageTransformer.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerFeeTokenWithdrawn)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseFeeTokenWithdrawn(log types.Log) (*OnRampWithMessageTransformerFeeTokenWithdrawn, error) {
	event := new(OnRampWithMessageTransformerFeeTokenWithdrawn)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerOwnershipTransferRequestedIterator struct {
	Event *OnRampWithMessageTransformerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerOwnershipTransferRequested)
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
		it.Event = new(OnRampWithMessageTransformerOwnershipTransferRequested)
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

func (it *OnRampWithMessageTransformerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampWithMessageTransformerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerOwnershipTransferRequestedIterator{contract: _OnRampWithMessageTransformer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerOwnershipTransferRequested)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampWithMessageTransformerOwnershipTransferRequested, error) {
	event := new(OnRampWithMessageTransformerOwnershipTransferRequested)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampWithMessageTransformerOwnershipTransferredIterator struct {
	Event *OnRampWithMessageTransformerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampWithMessageTransformerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampWithMessageTransformerOwnershipTransferred)
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
		it.Event = new(OnRampWithMessageTransformerOwnershipTransferred)
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

func (it *OnRampWithMessageTransformerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampWithMessageTransformerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampWithMessageTransformerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampWithMessageTransformerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampWithMessageTransformerOwnershipTransferredIterator{contract: _OnRampWithMessageTransformer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampWithMessageTransformer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampWithMessageTransformerOwnershipTransferred)
				if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformerFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampWithMessageTransformerOwnershipTransferred, error) {
	event := new(OnRampWithMessageTransformerOwnershipTransferred)
	if err := _OnRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetAllowedSendersList struct {
	IsEnabled           bool
	ConfiguredAddresses []common.Address
}
type GetDestChainConfig struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}

func (OnRampWithMessageTransformerAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (OnRampWithMessageTransformerAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (OnRampWithMessageTransformerAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (OnRampWithMessageTransformerCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32")
}

func (OnRampWithMessageTransformerConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1")
}

func (OnRampWithMessageTransformerDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
}

func (OnRampWithMessageTransformerFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (OnRampWithMessageTransformerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampWithMessageTransformerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRampWithMessageTransformer *OnRampWithMessageTransformer) Address() common.Address {
	return _OnRampWithMessageTransformer.address
}

type OnRampWithMessageTransformerInterface interface {
	GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error)

	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetMessageTransformer(opts *bind.CallOpts) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	SetMessageTransformer(opts *bind.TransactOpts, messageTransformerAddr common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampWithMessageTransformerAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*OnRampWithMessageTransformerAllowListAdminSet, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*OnRampWithMessageTransformerAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*OnRampWithMessageTransformerAllowListSendersRemoved, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampWithMessageTransformerCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampWithMessageTransformerCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampWithMessageTransformerConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampWithMessageTransformerConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampWithMessageTransformerDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampWithMessageTransformerDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampWithMessageTransformerFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampWithMessageTransformerFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampWithMessageTransformerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampWithMessageTransformerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampWithMessageTransformerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampWithMessageTransformerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampWithMessageTransformerOwnershipTransferred, error)

	Address() common.Address
}
