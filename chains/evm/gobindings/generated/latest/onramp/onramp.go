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

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMessageId\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageIdUnexpectedlySet\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6101006040523461055f576140688038038061001a81610599565b92833981019080820390610140821261055f576080821261055f5761003d61057a565b90610047816105be565b82526020810151906001600160a01b038216820361055f5760208301918252610072604082016105d2565b926040810193845260a0610088606084016105d2565b6060830190815295607f19011261055f5760405160a081016001600160401b03811182821017610564576040526100c1608084016105d2565b81526100cf60a084016105e6565b602082019081526100e260c085016105d2565b91604081019283526100f660e086016105d2565b936060820194855261010b61010087016105d2565b6080830190815261012087015190966001600160401b03821161055f57018a601f8201121561055f578051906001600160401b0382116105645760209b8c6060610159828660051b01610599565b9e8f8681520194028301019181831161055f57602001925b8284106104f1575050505033156104e057600180546001600160a01b0319163317905580516001600160401b03161580156104ce575b80156104bc575b80156104aa575b61047d57516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e08190528251909116158015610498575b801561048e575b61047d57815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19861012098606061029f61057a565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a16000905b805182101561040b5761032b82826105f3565b51916001600160401b0361033f82846105f3565b5151169283156103f65760008481526006602090815260409182902081840151815494840151600160401b600160e81b03198616604883901b600160481b600160e81b031617901515851b68ff000000000000000016179182905583516001600160401b0390951685526001600160a01b031691840191909152811c60ff1615159082015291926001927fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef590606090a20190610318565b8363c35aa79d60e01b60005260045260246000fd5b604051613a4a908161061e8239608051818181610704015281816122ce01528181612a710152613538015260a051818181612307015281816128240152612aaa015260c051818181611113015281816123430152612ae6015260e05181818161237f01528181612b2201526131520152f35b6306b7c75960e31b60005260046000fd5b5082511515610200565b5084516001600160a01b0316156101f9565b5088516001600160a01b0316156101b5565b5087516001600160a01b0316156101ae565b5086516001600160a01b0316156101a7565b639b15e16f60e01b60005260046000fd5b60608483031261055f5760405190606082016001600160401b0381118382101761056457604052610521856105be565b82526020850151906001600160a01b038216820361055f578260209283606095015261054f604088016105e6565b6040820152815201930192610171565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761056457604052565b6040519190601f01601f191682016001600160401b0381118382101761056457604052565b51906001600160401b038216820361055f57565b51906001600160a01b038216820361055f57565b5190811515820361055f57565b80518210156106075760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c90816306285c6914612a0b57508063181f5a771461298c57806320487ded146127485780632716072b1461249857806327e936f11461209257806348a98aa41461200f5780635cb80c5d14611d275780636def4ce714611c985780637437ff9f14611b7b57806379ba509714611a965780638da5cb5b14611a445780639041be3d14611997578063972b4612146118c9578063c9b146b314611504578063df0aa9e91461054d578063f2890a2114610237578063f2fde38b1461014a5763fbca3b74146100e757600080fd5b346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757600490610121612bfc565b507f9e7177c8000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475773ffffffffffffffffffffffffffffffffffffffff610197612d33565b61019f6136e7565b1633811461020f57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147576004359067ffffffffffffffff8211610147577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301916101a083126105455760a0604051936102b885612c94565b12610545576040516102c981612c78565b816004013581526102dc60248301612c18565b60208201526102ed60448301612c18565b60408201526102fe60648301612c18565b606082015261030f60848301612c18565b6080820152835261032260a48201612d79565b602084015260c481013567ffffffffffffffff81116105495761034b9060043691840101612e86565b604084015260e481013567ffffffffffffffff8111610549576103749060043691840101612e86565b606084015261010481013567ffffffffffffffff81116105495761039e9060043691840101612e86565b60808401526103b06101248201612d79565b60a084015261014481013560c084015261016481013560e08401526101848101359067ffffffffffffffff8211610549570136602382011215610545576004810135906103fc82612d0e565b9261040a6040519485612ccd565b828452602060048186019460051b84010101913683116105455760248101935b83851061044b57602061044388886101008201526134e9565b604051908152f35b843567ffffffffffffffff81116105415760049083010160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08236030112610541576040519061049b82612c78565b6104a760208201612d79565b8252604081013567ffffffffffffffff811161053d576104cd9060203691840101612e86565b6020830152606081013567ffffffffffffffff811161053d576104f69060203691840101612e86565b60408301526080810135606083015260a081013567ffffffffffffffff811161053d579161052d6020949285809536920101612e86565b608082015281520194019361042a565b8580fd5b8380fd5b5080fd5b8280fd5b50346101475760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757610585612bfc565b67ffffffffffffffff602435116105455760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60243536030112610545576105cc612d56565b60025460ff8160a01c166114dc577fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001760025567ffffffffffffffff8216835260066020526040832073ffffffffffffffffffffffffffffffffffffffff8216156114b457805460ff8160401c16611446575b60481c73ffffffffffffffffffffffffffffffffffffffff16330361141e5773ffffffffffffffffffffffffffffffffffffffff60035416806113b5575b50805467ffffffffffffffff811667ffffffffffffffff8114611388579067ffffffffffffffff60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009493011692839116179055604051906106f682612c78565b84825267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083015267ffffffffffffffff8416604083015260608201528360808201526107576024803501602435600401613300565b61076660046024350180613300565b61077460646024350161320c565b93610789604460243501602435600401613351565b9490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06107cf6107b987612d0e565b966107c76040519889612ccd565b808852612d0e565b018a5b818110611371575050604051966107e888612c94565b875273ffffffffffffffffffffffffffffffffffffffff88166020880152369061081192612e4f565b6040860152369061082192612e4f565b60608401526020916040516108368482612ccd565b878152608085015273ffffffffffffffffffffffffffffffffffffffff1660a084015260443560c08401528560e084015261010083015260025473ffffffffffffffffffffffffffffffffffffffff1693856024356064016108979061320c565b956108ab6024356084810190600401613300565b9091906108bd60046024350180613300565b99906040519a8b95869485947f3a49bb4900000000000000000000000000000000000000000000000000000000865267ffffffffffffffff8b16600487015273ffffffffffffffffffffffffffffffffffffffff16602486015260443560448601526064850160a0905260a485019061093592612f01565b908382037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc01608485015261096992612f01565b03915afa801561126f57868796889089936112f0575b50608086015260e085015261099e604460243501602435600401613351565b90966109a982612d0e565b916109b76040519384612ccd565b8083528583018099368360061b8201116112ec5780915b8360061b820183106112b55750505050885b6109f4604460243501602435600401613351565b9050811015610d5c57610a0781846132bd565b5190610a116133a5565b508682015115610d345773ffffffffffffffffffffffffffffffffffffffff610a3c818451166130f3565b169182158015610c91575b610c4f57808c878a8a73ffffffffffffffffffffffffffffffffffffffff8f83610af498015182808951169260405197610a8089612c78565b885267ffffffffffffffff87890196168652816040890191168152606088019283526080880193845267ffffffffffffffff6040519b8c998a997f9a4575b9000000000000000000000000000000000000000000000000000000008b5260048b01525160a060248b015260c48a0190612bb9565b965116604488015251166064860152516084850152511660a4830152038183885af1918215610c44578d8093610b8c575b50506001938284928b80610b8596519301519101519160405193610b4885612c78565b84528c840152604083015260608201528d60405190610b678c83612ccd565b815260808201526101008b015190610b7f83836132bd565b526132bd565b50016109e0565b9250933d8093863e610b9e8386612ccd565b89858481010312610c405784519167ffffffffffffffff8311610c3c5760408387018588010312610c3c57604051610bd581612cb1565b8387015167ffffffffffffffff8111610c3757610bf99086890190868a01016133d1565b81528b84880101519467ffffffffffffffff8611610c375787610b85968896610c289360019b019201016133d1565b8c82015293509150938d610b25565b508f80fd5b8e80fd5b8d80fd5b6040513d8f823e3d90fd5b517fbf16aab6000000000000000000000000000000000000000000000000000000008c5273ffffffffffffffffffffffffffffffffffffffff1660045260248bfd5b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201528881602481875afa908115610c44578d91610cfb575b5015610a47565b90508881813d8311610d2d575b610d128183612ccd565b81010312610d2957610d2390612ea4565b38610cf4565b8c80fd5b503d610d08565b60048b7f5cf04449000000000000000000000000000000000000000000000000000000008152fd5b5090889692508690610dd49873ffffffffffffffffffffffffffffffffffffffff600254169061010089015192886040519c8d957f01447eaa00000000000000000000000000000000000000000000000000000000875267ffffffffffffffff8b166004880152606060248801526064870190613413565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8684030160448701525191828152019190855b8a82821061127a575050505082809103915afa96871561126f578697611196575b50156110a45750835b67ffffffffffffffff608085510191169052835b61010084015151811015610e7d5780610e62600192886132bd565b516080610e74836101008901516132bd565b51015201610e47565b50909261010060609160405190610e9382612c94565b610e9b61322d565b8252808783015283604083015283808301528360808301528060a08301528060c083015260e08201520152610ecf826134e9565b82515267ffffffffffffffff60608351015116907f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f3260405185815267ffffffffffffffff608086518051898501528289820151166040850152826040820151166060850152826060820151168285015201511660a082015273ffffffffffffffffffffffffffffffffffffffff60208601511660c08201528061106f610ff6610fc1610f8c60408a01516101a060e08701526101c0860190612bb9565b60608a01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086830301610100870152612bb9565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301610120860152612bb9565b73ffffffffffffffffffffffffffffffffffffffff60a08901511661014084015260c088015161016084015260e088015161018084015267ffffffffffffffff610100890151967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403016101a08601521695613413565b0390a37fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff600254166002555151604051908152f35b73ffffffffffffffffffffffffffffffffffffffff604051917fea458c0c00000000000000000000000000000000000000000000000000000000835267ffffffffffffffff8416600484015216602482015282816044818873ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561118b578591611149575b50610e33565b90508281813d8311611184575b6111608183612ccd565b81010312611180575167ffffffffffffffff811681036111805786611143565b8480fd5b503d611156565b6040513d87823e3d90fd5b9096503d908187823e6111a98282612ccd565b8481838101031261126b5780519167ffffffffffffffff831161126757808201601f8484010112156112675782820151916111e383612d0e565b936111f16040519586612ccd565b83855287850192808301898660051b848601010111611263578882840101935b898660051b8486010101851061122e575050505050509587610e2a565b845167ffffffffffffffff8111610d29578a8092611256829383878a0191898b0101016133d1565b8152019501949050611211565b8a80fd5b8780fd5b8680fd5b6040513d88823e3d90fd5b8351805173ffffffffffffffffffffffffffffffffffffffff168652810151818601528d97508e965060409094019390920191600101610e09565b604083360312610d29578860409182516112ce81612cb1565b6112d786612d79565b815282860135838201528152019201916109ce565b8b80fd5b97505050503d8087873e6113048187612ccd565b608086828101031261126b5785519061131e848801612ea4565b90604088015167ffffffffffffffff811161136d5761134290828a01908a016133d1565b97606081015167ffffffffffffffff81116112635761136492820191016133d1565b9190963861097f565b8980fd5b60209061137c6133a5565b82828a010152016107d2565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b803b15611180578460405180927fe0a0e5060000000000000000000000000000000000000000000000000000000082528183816113fb6024356004018b60048401612f40565b03925af1801561118b5715610694578461141791959295612ccd565b9238610694565b6004847f1c0a3529000000000000000000000000000000000000000000000000000000008152fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526002830160205260409020546106565760248573ffffffffffffffffffffffffffffffffffffffff857fd0d2597600000000000000000000000000000000000000000000000000000000835216600452fd5b6004847fa4ec7479000000000000000000000000000000000000000000000000000000008152fd5b6004847f3ee5aeb5000000000000000000000000000000000000000000000000000000008152fd5b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475760043567ffffffffffffffff811161054557611554903690600401612d9a565b73ffffffffffffffffffffffffffffffffffffffff600154163303611881575b919081907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8181360301915b8481101561187d578060051b820135838112156111805782019160808336031261118057604051946115d086612c2d565b6115d984612c18565b86526115e760208501612d26565b9660208701978852604085013567ffffffffffffffff8111610549576116109036908701613258565b9460408801958652606081013567ffffffffffffffff81116105415761163891369101613258565b60608801908152875167ffffffffffffffff1683526006602052604080842099518a547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff169015159182901b68ff000000000000000016178a55909590815151611755575b5095976001019550815b855180518210156116e657906116df73ffffffffffffffffffffffffffffffffffffffff6116d7836001956132bd565b5116896137db565b50016116a7565b50509590969450600192919351908151611706575b50500193929361159f565b61174b67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190612dcb565b0390a238806116fb565b9893959296919094979860001461184657600184019591875b865180518210156117eb576117988273ffffffffffffffffffffffffffffffffffffffff926132bd565b511680156117b457906117ad6001928a61374a565b500161176e565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509692955090929796937f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161183c67ffffffffffffffff8a51169251604051918291602083526020830190612dcb565b0390a2388061169d565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8380f35b73ffffffffffffffffffffffffffffffffffffffff60055416330315611574576004837f905d7d9b000000000000000000000000000000000000000000000000000000008152fd5b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475767ffffffffffffffff61190a612bfc565b16808252600660205260ff604083205460401c16908252600660205260016040832001916040518093849160208254918281520191845260208420935b81811061197e57505061195c92500383612ccd565b61197a60405192839215158352604060208401526040830190612dcb565b0390f35b8454835260019485019487945060209093019201611947565b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475767ffffffffffffffff6119d8612bfc565b1681526006602052600167ffffffffffffffff604083205416019067ffffffffffffffff8211611a175760208267ffffffffffffffff60405191168152f35b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b503461014757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461014757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757805473ffffffffffffffffffffffffffffffffffffffff81163303611b53577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461014757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757611bb261322d565b5060a0604051611bc181612c78565b60ff60025473ffffffffffffffffffffffffffffffffffffffff81168352831c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015273ffffffffffffffffffffffffffffffffffffffff60045416606082015273ffffffffffffffffffffffffffffffffffffffff600554166080820152611c96604051809273ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565bf35b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757604060609167ffffffffffffffff611cde612bfc565b1681526006602052205473ffffffffffffffffffffffffffffffffffffffff6040519167ffffffffffffffff8116835260ff8160401c161515602084015260481c166040820152f35b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101475760043567ffffffffffffffff811161054557611d77903690600401612d9a565b9073ffffffffffffffffffffffffffffffffffffffff6004541690835b8381101561200b5773ffffffffffffffffffffffffffffffffffffffff611dbf8260051b840161320c565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315612000578793611fcd575b5082611e17575b506001915001611d94565b8460405193611ed260208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611e62606482612ccd565b8a80604098895193611e748b86612ccd565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d15611fc5573d90611eb582612e15565b91611ec28a519384612ccd565b82523d8d602084013e5b8661396d565b805180611f0e575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a338611e0c565b819294959693509060209181010312611fc1576020611f2d9101612ea4565b15611f3e5792919085903880611eda565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b8880fd5b606090611ecc565b9092506020813d8211611ff8575b81611fe860209383612ccd565b8101031261126b57519138611e05565b3d9150611fdb565b6040513d89823e3d90fd5b8480f35b50346101475760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757612047612bfc565b506024359073ffffffffffffffffffffffffffffffffffffffff82168203610147576020612074836130f3565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b50346101475760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147576040516120ce81612c78565b6120d6612d33565b81526024358015158103610549576020820190815260443573ffffffffffffffffffffffffffffffffffffffff81168103610541576040830190815261211a612d56565b90606084019182526084359273ffffffffffffffffffffffffffffffffffffffff8416840361053d57608085019384526121526136e7565b73ffffffffffffffffffffffffffffffffffffffff855116158015612479575b801561246f575b612447579273ffffffffffffffffffffffffffffffffffffffff859381809461012097827fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19a51167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff00000000000000000000000000000000000000006005541617600555612443604051916122c383612c2d565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602084015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604084015273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660608401526123f3604051809473ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b608083019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565ba180f35b6004867f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515612179565b5073ffffffffffffffffffffffffffffffffffffffff83511615612172565b50346101475760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147576004359067ffffffffffffffff821161014757366023830112156101475781600401356124f481612d0e565b926125026040519485612ccd565b818452602460606020860193028201019036821161054157602401915b8183106126a0575050506125316136e7565b805b825181101561269c5761254681846132bd565b5167ffffffffffffffff61255a83866132bd565b51511690811561267057907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5606060019493838752600660205260ff60408820612631604060208501519483547fffffff0000000000000000000000000000000000000000ffffffffffffffffff7cffffffffffffffffffffffffffffffffffffffff0000000000000000008860481b1691161784550151151582907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff0000000000000000835492151560401b169116179055565b5473ffffffffffffffffffffffffffffffffffffffff6040519367ffffffffffffffff8316855216602084015260401c1615156040820152a201612533565b602484837fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b606083360312610541576040516060810181811067ffffffffffffffff82111761271b576040526126d084612c18565b8152602084013573ffffffffffffffffffffffffffffffffffffffff8116810361053d57918160609360208094015261270b60408701612d26565b604082015281520192019161251f565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346101475760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261014757612780612bfc565b60243567ffffffffffffffff81116105495760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610549576040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008360801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612981578491612947575b50612911576128b39160209173ffffffffffffffffffffffffffffffffffffffff60025416906040518095819482937fd8694ccd0000000000000000000000000000000000000000000000000000000084526004019060048401612f40565b03915afa9081156129065782916128d0575b602082604051908152f35b90506020813d6020116128fe575b816128eb60209383612ccd565b81010312610545576020915051386128c5565b3d91506128de565b6040513d84823e3d90fd5b60248367ffffffffffffffff847ffdbd6a7200000000000000000000000000000000000000000000000000000000835216600452fd5b90506020813d602011612979575b8161296260209383612ccd565b810103126105415761297390612ea4565b38612854565b3d9150612955565b6040513d86823e3d90fd5b503461014757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610147575061197a6040516129cd604082612ccd565b601081527f4f6e52616d7020312e362e322d646576000000000000000000000000000000006020820152604051918291602083526020830190612bb9565b90503461054557817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105455780612a47606092612c2d565b82815282602082015282604082015201526080604051612a6681612c2d565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060820152611c96604051809273ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b60005b838110612ba95750506000910152565b8181015183820152602001612b99565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093612bf581518092818752878088019101612b96565b0116010190565b6004359067ffffffffffffffff82168203612c1357565b600080fd5b359067ffffffffffffffff82168203612c1357565b6080810190811067ffffffffffffffff821117612c4957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff821117612c4957604052565b610120810190811067ffffffffffffffff821117612c4957604052565b6040810190811067ffffffffffffffff821117612c4957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612c4957604052565b67ffffffffffffffff8111612c495760051b60200190565b35908115158203612c1357565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203612c1357565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203612c1357565b359073ffffffffffffffffffffffffffffffffffffffff82168203612c1357565b9181601f84011215612c135782359167ffffffffffffffff8311612c13576020808501948460051b010111612c1357565b906020808351928381520192019060005b818110612de95750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612ddc565b67ffffffffffffffff8111612c4957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612e5b82612e15565b91612e696040519384612ccd565b829481845281830111612c13578281602093846000960137010152565b9080601f83011215612c1357816020612ea193359101612e4f565b90565b51908115158203612c1357565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215612c1357016020813591019167ffffffffffffffff8211612c13578136038313612c1357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9067ffffffffffffffff9093929316815260406020820152612fb6612f79612f688580612eb1565b60a0604086015260e0850191612f01565b612f866020860186612eb1565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0858403016060860152612f01565b9060408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215612c135784016020813591019267ffffffffffffffff8211612c13578160061b36038413612c13578281037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0016080840152818152602001929060005b8181106130b4575050506130848473ffffffffffffffffffffffffffffffffffffffff6130746060612ea1979801612d79565b1660a08401526080810190612eb1565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152612f01565b90919360408060019273ffffffffffffffffffffffffffffffffffffffff6130db89612d79565b16815260208881013590820152019501929101613041565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156132005760009161319d575b5073ffffffffffffffffffffffffffffffffffffffff1690565b6020813d6020116131f8575b816131b660209383612ccd565b8101031261054557519073ffffffffffffffffffffffffffffffffffffffff82168203610147575073ffffffffffffffffffffffffffffffffffffffff613183565b3d91506131a9565b6040513d6000823e3d90fd5b3573ffffffffffffffffffffffffffffffffffffffff81168103612c135790565b6040519061323a82612c78565b60006080838281528260208201528260408201528260608201520152565b9080601f83011215612c1357813561326f81612d0e565b9261327d6040519485612ccd565b81845260208085019260051b820101928311612c1357602001905b8282106132a55750505090565b602080916132b284612d79565b815201910190613298565b80518210156132d15760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612c13570180359067ffffffffffffffff8211612c1357602001918136038313612c1357565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612c13570180359067ffffffffffffffff8211612c1357602001918160061b36038313612c1357565b604051906133b282612c78565b6060608083600081528260208201528260408201526000838201520152565b81601f82011215612c135780516133e781612e15565b926133f56040519485612ccd565b81845260208284010111612c1357612ea19160208085019101612b96565b9080602083519182815201916020808360051b8301019401926000915b83831061343f57505050505090565b90919293946020806134da837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806134bf6134ad8685015160a08886015260a0850190612bb9565b60408501518482036040860152612bb9565b92606081015160608401520151906080818403910152612bb9565b97019301930191939290613430565b805151806136ba575067ffffffffffffffff6040825101511660405160208101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321835267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604083015260608201523060808201526080815261357760a082612ccd565b5190209073ffffffffffffffffffffffffffffffffffffffff602082015116815167ffffffffffffffff6080816060840151169201511673ffffffffffffffffffffffffffffffffffffffff60a08501511660c085015191604051936020850195865260408501526060840152608083015260a082015260a081526135fd60c082612ccd565b5190209060608101516020815191012090604081015160208151910120608061010083015160405161366b8161363f6020820194602086526040830190613413565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612ccd565b51902092015160208151910120926040519460208601966000885260408701526060860152608085015260a084015260c083015260e082015260e081526136b461010082612ccd565b51902090565b7f4c8ebcc00000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361370857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156132d15760005260206000200190600090565b60008281526001820160205260409020546137d45780549068010000000000000000821015612c4957826137bd613788846001809601855584613732565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014613964577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613935578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613935578181036138fe575b505050805480156138cf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906138908282613732565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61391e61390e6137889386613732565b90549060031b1c92839286613732565b905560005283602052604060002055388080613858565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b919290156139e85750815115613981575090565b3b1561398a5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156139fb5750805190602001fd5b613a39906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612bb9565b0390fdfea164736f6c634300081a000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig, destChainConfigArgs []OnRampDestChainConfigArgs) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
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

func (_OnRamp *OnRampCaller) GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "generateMessageId", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OnRamp *OnRampSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRamp.Contract.GenerateMessageId(&_OnRamp.CallOpts, message)
}

func (_OnRamp *OnRampCallerSession) GenerateMessageId(message InternalEVM2AnyRampMessage) ([32]byte, error) {
	return _OnRamp.Contract.GenerateMessageId(&_OnRamp.CallOpts, message)
}

func (_OnRamp *OnRampCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

	error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	outstruct := new(GetAllowedSendersList)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfiguredAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_OnRamp *OnRampSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRamp.Contract.GetAllowedSendersList(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _OnRamp.Contract.GetAllowedSendersList(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AllowlistEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_OnRamp *OnRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
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

func (_OnRamp *OnRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_OnRamp *OnRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowlistUpdates(&_OnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_OnRamp *OnRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []OnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowlistUpdates(&_OnRamp.TransactOpts, allowlistConfigArgsItems)
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

type OnRampAllowListAdminSetIterator struct {
	Event *OnRampAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListAdminSet)
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
		it.Event = new(OnRampAllowListAdminSet)
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

func (it *OnRampAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListAdminSet struct {
	AllowlistAdmin common.Address
	Raw            types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampAllowListAdminSetIterator, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListAdminSetIterator{contract: _OnRamp.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListAdminSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListAdminSet(log types.Log) (*OnRampAllowListAdminSet, error) {
	event := new(OnRampAllowListAdminSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListSendersAddedIterator struct {
	Event *OnRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListSendersAdded)
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
		it.Event = new(OnRampAllowListSendersAdded)
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

func (it *OnRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListSendersAddedIterator{contract: _OnRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListSendersAdded)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListSendersAdded(log types.Log) (*OnRampAllowListSendersAdded, error) {
	event := new(OnRampAllowListSendersAdded)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListSendersRemovedIterator struct {
	Event *OnRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListSendersRemoved)
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
		it.Event = new(OnRampAllowListSendersRemoved)
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

func (it *OnRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListSendersRemovedIterator{contract: _OnRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListSendersRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*OnRampAllowListSendersRemoved, error) {
	event := new(OnRampAllowListSendersRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
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
	AllowlistEnabled  bool
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

type GetAllowedSendersList struct {
	IsEnabled           bool
	ConfiguredAddresses []common.Address
}
type GetDestChainConfig struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}

func (OnRampAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (OnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (OnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (OnRampCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
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
	GenerateMessageId(opts *bind.CallOpts, message InternalEVM2AnyRampMessage) ([32]byte, error)

	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

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

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*OnRampAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*OnRampAllowListAdminSet, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*OnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*OnRampAllowListSendersRemoved, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

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
