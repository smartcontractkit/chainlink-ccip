// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type InternalEVM2AnyVerifierMessage struct {
	Header         InternalHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenTransfer  InternalEVMTokenTransfer
	Receipts       []InternalReceipt
}

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	SourcePoolAddress  common.Address
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
	DestExecData       []byte
	RequiredVerifierId [32]byte
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalReceipt struct {
	ReceiptType       uint8
	Issuer            common.Address
	FeeTokenAmount    *big.Int
	DestGasLimit      uint64
	DestBytesOverhead uint32
	ExtraArgs         []byte
}

type VerifierProxyDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	DefaultVerifier   common.Address
	RequiredVerifier  common.Address
}

type VerifierProxyDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type VerifierProxyStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
}

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVMTokenTransfer\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Receipt[]\",\"components\":[{\"name\":\"receiptType\",\"type\":\"uint8\",\"internalType\":\"enumInternal.ReceiptType\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierId\",\"inputs\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MalformedVerifierExtraArgs\",\"inputs\":[{\"name\":\"verifierExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60e060405234610488576131be80380380610019816104c2565b92833981019080820360e0811261048857606081126104885761003a6104a3565b90610044836104e7565b82526020830151906001600160a01b038216820361048857602083019182526060610071604086016104fb565b6040850190815291605f1901126104885761008a6104a3565b91610097606086016104fb565b8352608085015193841515850361048857602084019485526100bb60a087016104fb565b6040850190815260c087015190966001600160401b038211610488570187601f82011215610488578051906001600160401b03821161048d5761010360208360051b016104c2565b986020808b858152019360071b8301019181831161048857602001925b828410610409575050505033156103f857600180546001600160a01b0319163317905580516001600160401b03161580156103e6575b80156103d4575b6103a757516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156103c2575b80156103b8575b6103a75780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406102116104a3565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b03908116911660406102496104a3565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a16000905b8051821015610368576102ae828261050f565b51916001600160401b036102c2828461050f565b51511692831561035357600084815260046020908152604091829020928101518354600160401b600160e01b0319811682851b600160401b600160e01b03161790945582516001600160401b0390941684526001600160a01b031690830152929360019390917ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed09190a2019061029b565b8363c35aa79d60e01b60005260045260246000fd5b604051612c84908161053a8239608051818181610679015261235c015260a051818181611d270152612395015260c0518181816123d101526124ff0152f35b6306b7c75960e31b60005260046000fd5b5081511515610196565b5082516001600160a01b03161561018f565b5082516001600160a01b03161561015d565b5081516001600160a01b031615610156565b639b15e16f60e01b60005260046000fd5b6080848303126104885760405190608082016001600160401b0381118382101761048d57604052610439856104e7565b82526020850151906001600160a01b03821682036104885782602092836080950152610467604088016104fb565b6040820152610478606088016104fb565b6060820152815201930192610120565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761048d57604052565b6040519190601f01601f191682016001600160401b0381118382101761048d57604052565b51906001600160401b038216820361048857565b51906001600160a01b038216820361048857565b80518210156105235760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c6914612076578063181f5a7714611ff557806320487ded14611c3a57806348a98aa414611bb95780635cb80c5d146118b85780636def4ce7146118395780637437ff9f1461176f57806379ba5097146116865780638da5cb5b146116345780639041be3d146115b457806390423fa214611348578063caf1bd0e14611111578063df0aa9e914610218578063f2fde38b146101285763fbca3b74146100c157600080fd5b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576100f86122aa565b507f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235773ffffffffffffffffffffffffffffffffffffffff6101746122c1565b61017c6129f1565b163381146101ee57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235761024f6122aa565b6024359067ffffffffffffffff8211610123578160040160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8436030112610123576064359073ffffffffffffffffffffffffffffffffffffffff8216809203610123576002549260ff8460a01c166110e75767ffffffffffffffff90740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff861617600255168060005260046020526040600020926040519361032785612159565b80549567ffffffffffffffff8716865273ffffffffffffffffffffffffffffffffffffffff602087019760401c16875273ffffffffffffffffffffffffffffffffffffffff60038160018501541693604089019485528260028201541660608a0152015416966080870197885283156110bd575173ffffffffffffffffffffffffffffffffffffffff163303611093576103c460848901866125d2565b6040519192919060006103d683612175565b8083526020830192606084526040810191606083526060808301526080820152606060a08201529082600411610123578535937fffffffff0000000000000000000000000000000000000000000000000000000085167f302326cb000000000000000000000000000000000000000000000000000000000361102b575050508301926020818503126101235767ffffffffffffffff8211610123570160c081840312610123576040519261048984612175565b610492826122e4565b8452602082013567ffffffffffffffff811161012357816104b4918401612aa9565b6020850152604082013567ffffffffffffffff811161012357816104d9918401612ac4565b6040850152606082013567ffffffffffffffff811161012357816104fe918401612ac4565b936060810194855260808301359260ff84168403610123576080820193845260a08101359267ffffffffffffffff84116101235760ff9361053f9201612aa9565b60a082015293515191511690818111156101235715159081611022575b50610123575b6040820191825151906060810151518201809211610b585773ffffffffffffffffffffffffffffffffffffffff996000956105ec9560209415610fd4575b50508a82511615610fc9575b5001516040519889809481937f9b1115e4000000000000000000000000000000000000000000000000000000008352602060048401526024830190612267565b0392165afa948515610e0e57600095610fac575b50845115610f7d575b61061161268e565b946044870161062081866126c7565b9050610b87575b505067ffffffffffffffff8451169567ffffffffffffffff8714610b585760646106c49667ffffffffffffffff6001809a01168097526040519661066a8861213d565b6000885267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908160208a01528660408a015260608901526106cd6106bb60248601896125d2565b9a9098806125d2565b949095016125b1565b936020976040519a6106df8a8d6121c9565b60008c5260009c610b53575b6040519a6101208c018c811067ffffffffffffffff821117610b26576040528b52898b01978852610731929161072291369161271b565b9460408b01958652369161271b565b936060890194855273ffffffffffffffffffffffffffffffffffffffff60808a019116815273ffffffffffffffffffffffffffffffffffffffff60a08a019260443584528c60c08c015260e08b019485526101008b019b8c526040518a8101917f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f732183526040820152896060820152306080820152608081526107d460a0826121c9565b5190209651169173ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60608c51015116925116905190604051928a84019485526040840152606083015260808201526080815261082e60a0826121c9565b51902092518681519101209151868151910120905160405161088a8161085e8a8201948b86526040830190612752565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826121c9565b5190209160405193878501958b875260408601526060850152608084015260a083015260c082015260c081526108c160e0826121c9565b51902083515260405192828085015261090d846108e160408201846127e6565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018652856121c9565b8451517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061095361093d83612305565b9261094b60405194856121c9565b808452612305565b0184885b828110610b1657505050865b86518051821015610a21578561098e8373ffffffffffffffffffffffffffffffffffffffff936129dd565b51015116908860405180937f1234eab0000000000000000000000000000000000000000000000000000000008252604060048301528183816109d48d6044830190612267565b87602483015203925af1918215610a16576001926109f4575b5001610963565b610a0f903d808c833e610a0781836121c9565b810190612668565b50896109ed565b6040513d8b823e3d90fd5b505090918667ffffffffffffffff60608551015116926040519160408352610a4c60408401876127e6565b8381038885015282519081815288810189808460051b840101950193915b8a848410610acd578a8a8a7fe26fecb9952e6111d8ef6051a82bf31440597d92d3d60725e81c07c3a2d7c1df8b8b038ca37fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff600254166002555151604051908152f35b80610b07887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001969798999a9b030187528951612267565b97019301930191939290610a6a565b6060828286010152018590610957565b60248f7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6106eb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b909195506001610b9782866126c7565b905003610f5357610ba890846126c7565b15610f245760408136031261012357604051610bc381612191565b610bcc826122e4565b815260208082019201358252610be061268e565b50815115610efa5773ffffffffffffffffffffffffffffffffffffffff610c09818351166124a0565b169283158015610e5f575b610e1a57906000610cb884938a73ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff989751818751169060405194610c5486612159565b85528260208601948d86526040870190815260608701928352608087019384526040519c8d9889987f9a4575b9000000000000000000000000000000000000000000000000000000008a52602060048b01525160a060248b015260c48a0190612267565b965116604488015251166064860152516084850152511660a4830152038183875af1938415610e0e57600094610d56575b5073ffffffffffffffffffffffffffffffffffffffff905116926020815191015191519260405194610d1a866121ad565b85526020850152604084015260608301526080820152604051610d3e6020826121c9565b6000815260a0820152600060c0820152938680610627565b9390933d8083833e610d6881836121c9565b810190602081830312610e065780519067ffffffffffffffff8211610e0a570191604083830312610e035760405192610da084612191565b805167ffffffffffffffff8111610e065783610dbd918301612623565b845260208101519167ffffffffffffffff8311610e03575091610df79173ffffffffffffffffffffffffffffffffffffffff949301612623565b60208201529390610ce9565b80fd5b8280fd5b8380fd5b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff8251167fbf16aab60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610e0e57600091610ecb575b5015610c14565b610eed915060203d602011610ef3575b610ee581836121c9565b8101906123f9565b8a610ec4565b503d610edb565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b9350610fa6610f8c83806125d2565b604080516020808201529392849261085e92840191612461565b93610609565b610fc29195503d806000833e610a0781836121c9565b9386610600565b518a1681528b6105ac565b8c61101a92610fe1612a3c565b835251169060405191610ff383612191565b825260405161100287826121c9565b898152868301525190611014826129d0565b526129d0565b508c806105a0565b9050158a61055c565b61108d94509561107d91929661104236868461271b565b905261104c612a3c565b835273ffffffffffffffffffffffffffffffffffffffff885116936040519461107486612191565b8552369161271b565b60208301525190611014826129d0565b50610562565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760043567ffffffffffffffff8111610123573660238201121561012357806004013561116b81612305565b9161117960405193846121c9565b8183526024602084019260071b8201019036821161012357602401915b8183106112c257836111a66129f1565b6000905b80518210156112c0576111bd82826129dd565b519167ffffffffffffffff6111d282846129dd565b51511692831561129257927ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed060406001949583600052600460205273ffffffffffffffffffffffffffffffffffffffff6020836000209201518254927bffffffffffffffffffffffffffffffffffffffff000000000000000082861b167fffffffff0000000000000000000000000000000000000000ffffffffffffffff851617905567ffffffffffffffff845193168352166020820152a201906111aa565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b005b60808336031261012357604051906112d98261213d565b833567ffffffffffffffff8116810361012357825260208401359073ffffffffffffffffffffffffffffffffffffffff821682036101235782602092836080950152611327604087016122e4565b6040820152611338606087016122e4565b6060820152815201920191611196565b346101235760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576000604051611385816120f2565b61138d6122c1565b81526024358015158103610e0657602082019081526044359073ffffffffffffffffffffffffffffffffffffffff82168203610e0a57604083019182526113d26129f1565b73ffffffffffffffffffffffffffffffffffffffff835116158015611595575b801561158b575b611563579173ffffffffffffffffffffffffffffffffffffffff60c0927f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a394828451167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035561155f6114e361233c565b91611527604051809473ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b606083019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565ba180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b50805115156113f9565b5073ffffffffffffffffffffffffffffffffffffffff825116156113f2565b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235767ffffffffffffffff6115f46122aa565b166000526004602052600167ffffffffffffffff604060002054160167ffffffffffffffff8111610b585760209067ffffffffffffffff60405191168152f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760005473ffffffffffffffffffffffffffffffffffffffff81163303611745577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576117a661231d565b5060606040516117b5816120f2565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152611837604051809273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565bf35b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235767ffffffffffffffff6118796122aa565b1660005260046020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519167ffffffffffffffff81168352831c166020820152f35b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760043567ffffffffffffffff8111610123573660238201121561012357806004013567ffffffffffffffff8111610123573660248260051b840101116101235773ffffffffffffffffffffffffffffffffffffffff6003541660005b828110156112c05760009073ffffffffffffffffffffffffffffffffffffffff61197460248360051b8801016125b1565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115611bae579085918591611b7a575b50806119cf575b5050506001915001611943565b611a8860405160208101967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602483015283604483015260448252611a186064836121c9565b80806040998a5194611a2a8c876121c9565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208701525190828a5af13d15611b71573d611a6a8161220a565b90611a778b5192836121c9565b8152809260203d92013e5b86612ba7565b805180611ac6575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a38583816119c2565b611add9294959693506020809183010191016123f9565b15611aee5792919084908880611a90565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609150611a82565b9150506020813d8211611ba6575b81611b95602093836121c9565b81010312610e0a57849051886119bb565b3d9150611b88565b6040513d86823e3d90fd5b346101235760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357611bf06122aa565b5060243573ffffffffffffffffffffffffffffffffffffffff8116810361012357611c1c6020916124a0565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101235760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357611c716122aa565b60243567ffffffffffffffff811161012357806004018136039060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101235767ffffffffffffffff84169377ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e0e57600091611fd6575b50611fa85773ffffffffffffffffffffffffffffffffffffffff6002541692604051947fd8694ccd000000000000000000000000000000000000000000000000000000008652600486015260406024860152611e08611dcb611dba8480612411565b60a060448a015260e4890191612461565b611dd86024840185612411565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8984030160648a0152612461565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd604483013591018112156101235781016024600482013591019367ffffffffffffffff8211610123578160061b36038513610123578681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160848801528181528694602090910193929160005b818110611f6157505050611f1160209593611ee1869460848573ffffffffffffffffffffffffffffffffffffffff611ed460648a99016122e4565b1660a48801520190612411565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8584030160c4860152612461565b03915afa8015610e0e57600090611f2e575b602090604051908152f35b506020813d602011611f59575b81611f48602093836121c9565b810103126101235760209051611f23565b3d9150611f3b565b91955091929360408060019273ffffffffffffffffffffffffffffffffffffffff611f8b8a6122e4565b168152602089013560208201520196019101918795949392611e99565b837ffdbd6a720000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611fef915060203d602011610ef357610ee581836121c9565b85611d58565b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357612072604080519061203681836121c9565b601782527f566572696669657250726f787920312e372e302d646576000000000000000000602083015251918291602083526020830190612267565b0390f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760606120af61233c565b611837604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b6060810190811067ffffffffffffffff82111761210e57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761210e57604052565b60a0810190811067ffffffffffffffff82111761210e57604052565b60c0810190811067ffffffffffffffff82111761210e57604052565b6040810190811067ffffffffffffffff82111761210e57604052565b60e0810190811067ffffffffffffffff82111761210e57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761210e57604052565b67ffffffffffffffff811161210e57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106122575750506000910152565b8181015183820152602001612247565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936122a381518092818752878088019101612244565b0116010190565b6004359067ffffffffffffffff8216820361012357565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361012357565b359073ffffffffffffffffffffffffffffffffffffffff8216820361012357565b67ffffffffffffffff811161210e5760051b60200190565b6040519061232a826120f2565b60006040838281528260208201520152565b61234461231d565b50604051612351816120f2565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610123575180151581036101235790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561012357016020813591019167ffffffffffffffff821161012357813603831361012357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e0e5760009161254a575b5073ffffffffffffffffffffffffffffffffffffffff1690565b6020813d6020116125a9575b81612563602093836121c9565b810103126125a557519073ffffffffffffffffffffffffffffffffffffffff82168203610e03575073ffffffffffffffffffffffffffffffffffffffff612530565b5080fd5b3d9150612556565b3573ffffffffffffffffffffffffffffffffffffffff811681036101235790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610123570180359067ffffffffffffffff82116101235760200191813603831361012357565b81601f820112156101235780516126398161220a565b9261264760405194856121c9565b81845260208284010111610123576126659160208085019101612244565b90565b9060208282031261012357815167ffffffffffffffff8111610123576126659201612623565b6040519061269b826121ad565b600060c08382815282602082015260606040820152606080820152826080820152606060a08201520152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610123570180359067ffffffffffffffff821161012357602001918160061b3603831361012357565b9291926127278261220a565b9161273560405193846121c9565b829481845281830111610123578281602093846000960137010152565b9073ffffffffffffffffffffffffffffffffffffffff825116815273ffffffffffffffffffffffffffffffffffffffff602083015116602082015260c0806127dd6127c16127af604087015160e0604088015260e0870190612267565b60608701518682036060880152612267565b6080860151608086015260a086015185820360a0870152612267565b93015191015290565b67ffffffffffffffff6060825180518552826020820151166020860152826040820151166040860152015116606083015273ffffffffffffffffffffffffffffffffffffffff60208201511660808301526101006128b461286d61285b604085015161018060a0880152610180870190612267565b606085015186820360c0880152612267565b73ffffffffffffffffffffffffffffffffffffffff60808501511660e086015260a08401518386015260c084015161012086015260e0840151858203610140870152612752565b91015191610160818303910152815180825260208201906020808260051b8501019401926000905b8282106128eb57505050505090565b90919293947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301855285519081519160028310156129a1576129948260c060a06020959460019787965273ffffffffffffffffffffffffffffffffffffffff8682015116868501526040810151604085015267ffffffffffffffff606082015116606085015263ffffffff60808201511660808501520151918160a08201520190612267565b97019501939201906128dc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b805115610f245760200190565b8051821015610f245760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff600154163303612a1257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60408051909190612a4d83826121c9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b828110612a8457505050565b602090604051612a9381612191565b6000815260608382015282828501015201612a78565b9080601f83011215610123578160206126659335910161271b565b9080601f8301121561012357813591612adc83612305565b92612aea60405194856121c9565b80845260208085019160051b830101918383116101235760208101915b838310612b1657505050505090565b823567ffffffffffffffff81116101235782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101235760405190612b6382612191565b612b6f602084016122e4565b825260408301359167ffffffffffffffff831161012357612b9888602080969581960101612aa9565b83820152815201920191612b07565b91929015612c225750815115612bbb575090565b3b15612bc45790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015612c355750805190602001fd5b612c73906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612267565b0390fdfea164736f6c634300081a000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig VerifierProxyStaticConfig, dynamicConfig VerifierProxyDynamicConfig, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{address: address, abi: *parsed, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_VerifierProxy *VerifierProxySession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(VerifierProxyDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyDynamicConfig)).(*VerifierProxyDynamicConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCaller) GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(VerifierProxyStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyStaticConfig)).(*VerifierProxyStaticConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxySession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxySession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_VerifierProxy *VerifierProxySession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_VerifierProxy *VerifierProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

func (_VerifierProxy *VerifierProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

type VerifierProxyCCIPMessageSentIterator struct {
	Event *VerifierProxyCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyCCIPMessageSent)
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
		it.Event = new(VerifierProxyCCIPMessageSent)
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

func (it *VerifierProxyCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyVerifierMessage
	ReceiptBlobs      [][]byte
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCCIPMessageSentIterator{contract: _VerifierProxy.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyCCIPMessageSent)
				if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error) {
	event := new(VerifierProxyCCIPMessageSent)
	if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyConfigSetIterator struct {
	Event *VerifierProxyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyConfigSet)
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
		it.Event = new(VerifierProxyConfigSet)
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

func (it *VerifierProxyConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyConfigSet struct {
	StaticConfig  VerifierProxyStaticConfig
	DynamicConfig VerifierProxyDynamicConfig
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyConfigSetIterator{contract: _VerifierProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error) {
	event := new(VerifierProxyConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyDestChainConfigSetIterator struct {
	Event *VerifierProxyDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyDestChainConfigSet)
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
		it.Event = new(VerifierProxyDestChainConfigSet)
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

func (it *VerifierProxyDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyDestChainConfigSetIterator{contract: _VerifierProxy.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyDestChainConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error) {
	event := new(VerifierProxyDestChainConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyFeeTokenWithdrawnIterator struct {
	Event *VerifierProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyFeeTokenWithdrawn)
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
		it.Event = new(VerifierProxyFeeTokenWithdrawn)
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

func (it *VerifierProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFeeTokenWithdrawnIterator{contract: _VerifierProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyFeeTokenWithdrawn)
				if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error) {
	event := new(VerifierProxyFeeTokenWithdrawn)
	if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
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
		it.Event = new(VerifierProxyOwnershipTransferRequested)
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

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
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
		it.Event = new(VerifierProxyOwnershipTransferred)
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

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	SequenceNumber uint64
	Router         common.Address
}

func (_VerifierProxy *VerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierProxy.abi.Events["CCIPMessageSent"].ID:
		return _VerifierProxy.ParseCCIPMessageSent(log)
	case _VerifierProxy.abi.Events["ConfigSet"].ID:
		return _VerifierProxy.ParseConfigSet(log)
	case _VerifierProxy.abi.Events["DestChainConfigSet"].ID:
		return _VerifierProxy.ParseDestChainConfigSet(log)
	case _VerifierProxy.abi.Events["FeeTokenWithdrawn"].ID:
		return _VerifierProxy.ParseFeeTokenWithdrawn(log)
	case _VerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierProxy.ParseOwnershipTransferRequested(log)
	case _VerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _VerifierProxy.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierProxyCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0xe26fecb9952e6111d8ef6051a82bf31440597d92d3d60725e81c07c3a2d7c1df")
}

func (VerifierProxyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (VerifierProxyDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed0")
}

func (VerifierProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
