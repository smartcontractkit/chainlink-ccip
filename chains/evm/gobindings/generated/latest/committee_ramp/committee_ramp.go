// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package committee_ramp

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

type BaseOnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type BaseOnRampDestChainConfigArgs struct {
	Router            common.Address
	DestChainSelector uint64
	AllowlistEnabled  bool
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

type CommitteeRampDynamicConfig struct {
	FeeQuoter      common.Address
	FeeAggregator  common.Address
	AllowlistAdmin common.Address
}

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	OnRampAddress       []byte
	OffRampAddress      []byte
	Finality            uint16
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []MessageV1CodecTokenTransferV1
	Data                []byte
}

type MessageV1CodecTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	ExtraData          []byte
}

var CommitteeRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitteeRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"originalCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitteeRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitteeRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignatureConfig\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitteeRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureConfigSet\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonOrderedOrNonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]}]",
	Bin: "0x60c0604052346101af57604051601f6126eb38819003918201601f19168301916001600160401b038311848410176101b4578084926060946040528339810103126101af57604051600090606081016001600160401b0381118282101761019b5760405261006c836101ca565b815261008d604061007f602086016101ca565b9460208401958652016101ca565b9160408201928352331561018c57600180546001600160a01b031916331790554660805281516001600160a01b031615801561017a575b61016b578151600680546001600160a01b03199081166001600160a01b039384169081179092558651600780548316918516919091179055855160088054909216908416179055604080519182528651831660208301528551909216918101919091527fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d90606090a160405161250c90816101df8239608051816116fa015260a051815050f35b6306b7c75960e31b8152600490fd5b5083516001600160a01b0316156100c4565b639b15e16f60e01b8152600490fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101af5756fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611bd757508063181f5a7714611b5a57806334d560e41461194a57806358bfa40a146116225780635cb80c5d146113355780636def4ce7146112575780636ed0e217146111ab57806371c5c2ba14610e8d5780637437ff9f14610d9957806379ba509714610cb05780638da5cb5b14610c5e578063a8dd2df214610a2e578063b2d6d66b14610808578063c9b146b3146103fc578063e75dc119146101c45763f2fde38b146100cf57600080fd5b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5773ffffffffffffffffffffffffffffffffffffffff61011b611d89565b610123612012565b1633811461019557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101bf5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf576101fb611d89565b50610204611e4f565b5060443567ffffffffffffffff81116101bf5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101bf576040519060a0820182811067ffffffffffffffff8211176103cd57604052806004013567ffffffffffffffff81116101bf576102859060043691840101611f63565b8252602481013567ffffffffffffffff81116101bf576102ab9060043691840101611f63565b6020830152604481013567ffffffffffffffff81116101bf578101366023820112156101bf5760048101356102df81611ee6565b916102ed6040519384611caf565b818352602060048185019360061b83010101903682116101bf57602401915b81831061038157505050604083015261032760648201611dcf565b6060830152608481013567ffffffffffffffff81116101bf5760809160046103529236920101611f63565b91015260643567ffffffffffffffff81116101bf57610375903690600401611f63565b50602060405160008152f35b6040833603126101bf5760405190604082019082821067ffffffffffffffff8311176103cd5760409260209284526103b886611dcf565b8152828601358382015281520192019161030c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf5761044b903690600401611e1e565b73ffffffffffffffffffffffffffffffffffffffff6001541633036107be575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b818410156107bc576000938060051b840135828112156107b8578401916080833603126107b857604051946080860186811067ffffffffffffffff82111761078b576040526104e384611e66565b86526104f160208501611fc2565b9660208701978852604085013567ffffffffffffffff81116107875761051a9036908701611efe565b9460408801958652606081013567ffffffffffffffff81116107835761054291369101611efe565b946060880195865267ffffffffffffffff88511682526005602052604082209851151561059a818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b815151610657575b50959760010195505b845180518210156105ea57906105e373ffffffffffffffffffffffffffffffffffffffff6105db83600195611fcf565b511688612246565b50016105ab565b5050959490935060019251908151610608575b505001929190610495565b61064d67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190611e7b565b0390a285806105fd565b989395929094979896919660001461074c57600184019591875b865180518210156106ee5761069b8273ffffffffffffffffffffffffffffffffffffffff92611fcf565b511680156106b757906106b06001928a6123da565b5001610671565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161074267ffffffffffffffff8b51169251604051918291602083526020830190611e7b565b0390a290896105a2565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff6008541633031561046b577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf57610857903690600401611efe565b6024359060ff8216918281036101bf5761086f612012565b82158015610a24575b61096f575b60025415610907576000600254156108da57600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace546108d49073ffffffffffffffffffffffffffffffffffffffff166120b0565b5061087d565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60005b82518110156109c35773ffffffffffffffffffffffffffffffffffffffff6109328285611fcf565b5116156109995761096273ffffffffffffffffffffffffffffffffffffffff61095b8386611fcf565b511661237a565b1561096f5760010161090a565b7f12823a5e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50907fc2e12b820aa2dc1a1673e9f59d1d809598d1041a90baccc742b7de5e5d2418a8927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006004541617600455610a1f60405192839283611ec5565b0390a1005b5081518311610878565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf57366023820112156101bf5780600401359067ffffffffffffffff82116101bf576024606083028201013681116101bf57610aa7612012565b610ab083611ee6565b92610abe6040519485611caf565b835260009160240190602084015b818310610bf257505050805b8251811015610bee57610aeb8184611fcf565b5167ffffffffffffffff6020610b018487611fcf565b51015116908115610bc257818452600560205260408085208251815493830151151560ff9081167fffffffffffffffffffffff000000000000000000000000000000000000000000909516600883901b74ffffffffffffffffffffffffffffffffffffffff00161794909417825560019594937f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9392541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a201610ad8565b602484837fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b606083360312610c5a57604051610c0881611c93565b833573ffffffffffffffffffffffffffffffffffffffff811681036107b8578152606091602091610c3a868401611e66565b83820152610c4a60408701611fc2565b6040820152815201920191610acc565b8380fd5b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760005473ffffffffffffffffffffffffffffffffffffffff81163303610d6f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600060408051610dd781611c93565b8281528260208201520152610e89604051610df181611c93565b73ffffffffffffffffffffffffffffffffffffffff60065416815273ffffffffffffffffffffffffffffffffffffffff60075416602082015273ffffffffffffffffffffffffffffffffffffffff60085416604082015260405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390f35b346101bf5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57610ec4611d89565b6024359067ffffffffffffffff82116101bf578136036101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101bf57610f0d611dac565b5060a43567ffffffffffffffff81116101bf57610f2e903690600401611df0565b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60c484013591018112156101bf57820160048101359067ffffffffffffffff82116101bf576024019080360382136101bf57602491357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110611176575b505060601c92013567ffffffffffffffff81168091036101bf57806000526005602052604060002091825491604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff8760081c165afa90811561116a57600091611104575b5073ffffffffffffffffffffffffffffffffffffffff8091169116036110d65760ff1661108c575b60405160206110748183611caf565b60008252610e89604051928284938452830190611d2a565b600082815260029091016020526040902054156110a95780611065565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011611162575b8161111d60209383611caf565b8101031261078757519073ffffffffffffffffffffffffffffffffffffffff8216820361115f575073ffffffffffffffffffffffffffffffffffffffff61103d565b80fd5b3d9150611110565b6040513d6000823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168480610fb4565b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760405180816020600254928381520160026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9260005b81811061123e57505061122892500382611caf565b60ff6004541690610e8960405192839283611ec5565b8454835260019485019486945060209093019201611213565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81168091036101bf57600052600560205260406000206001815491019060405191826020825491828152019160005260206000209060005b81811061131f5773ffffffffffffffffffffffffffffffffffffffff85610e89886112f781890382611caf565b604051938360ff8695161515855260081c166020840152606060408401526060830190611e7b565b82548452602090930192600192830192016112ca565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf57611384903690600401611e1e565b73ffffffffffffffffffffffffffffffffffffffff6007541660005b828110156107bc576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff8116809103610783576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156116175790859185916115e3575b508061142a575b50505060019150016113a0565b604051946114e460208701967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602482015283604482015260448152611474606482611caf565b82806040998a51936114868c86611caf565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d156115db573d906114c782611cf0565b916114d48b519384611caf565b82523d85602084013e5b8761242f565b805180611523575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385838161141d565b8192939596979450906020918101031261078757602001519081159182150361115f57506115585792919084908880806114ec565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060906114de565b9150506020813d821161160f575b816115fe60209383611caf565b81010312610c5a5784905188611416565b3d91506115f1565b6040513d86823e3d90fd5b346101bf5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57611659611d89565b5060243567ffffffffffffffff81116101bf577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc61016091360301126101bf5760443560643567ffffffffffffffff81116101bf576116bc903690600401611df0565b6002811061192057806002116101bf578135918260f01c9182600201806002116118965780821061192057116101bf57600201926002541561096f577f00000000000000000000000000000000000000000000000000000000000000004681036118ef575060ff60045416809360f61c106118c557600091825b84841061173f57005b8360061b848104604014851517156118965760208101908181116118965761177261176c8383878c611faa565b9061205d565b90600092604082018092116118695760209261179661176c86946080948a8f611faa565b60405191898352601b868401526040830152606082015282805260015afa1561185d5773ffffffffffffffffffffffffffffffffffffffff8151169182825260036020526040822054156118355773ffffffffffffffffffffffffffffffffffffffff1682111561180d5750600190930192611736565b807fb70ad94b0000000000000000000000000000000000000000000000000000000060049252fd5b6004827fca31867a000000000000000000000000000000000000000000000000000000008152fd5b604051903d90823e3d90fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7fbba6473c0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600060405161198781611c93565b61198f611d89565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361078357602082019081526044359073ffffffffffffffffffffffffffffffffffffffff82168203610c5a57604083019182526119e8612012565b73ffffffffffffffffffffffffffffffffffffffff835116158015611b3b575b611b13579173ffffffffffffffffffffffffffffffffffffffff611b0d92817fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d95818551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600854161760085560405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff81511615611a08565b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57610e896040805190611b9b8183611caf565b601782527f436f6d6d697474656552616d7020312e372e302d646576000000000000000000602083015251918291602083526020830190611d2a565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101bf57817fce27a7a90000000000000000000000000000000000000000000000000000000060209314908115611c69575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611c62565b6060810190811067ffffffffffffffff8211176103cd57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103cd57604052565b67ffffffffffffffff81116103cd57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110611d745750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611d35565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b9181601f840112156101bf5782359167ffffffffffffffff83116101bf57602083818601950101116101bf57565b9181601f840112156101bf5782359167ffffffffffffffff83116101bf576020808501948460051b0101116101bf57565b6024359067ffffffffffffffff821682036101bf57565b359067ffffffffffffffff821682036101bf57565b906020808351928381520192019060005b818110611e995750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611e8c565b9060ff611edf602092959495604085526040850190611e7b565b9416910152565b67ffffffffffffffff81116103cd5760051b60200190565b9080601f830112156101bf578135611f1581611ee6565b92611f236040519485611caf565b81845260208085019260051b8201019283116101bf57602001905b828210611f4b5750505090565b60208091611f5884611dcf565b815201910190611f3e565b81601f820112156101bf57803590611f7a82611cf0565b92611f886040519485611caf565b828452602083830101116101bf57816000926020809301838601378301015290565b909392938483116101bf5784116101bf578101920390565b359081151582036101bf57565b8051821015611fe35760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361203357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b35906020811061206b575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b8054821015611fe35760005260206000200190600090565b600081815260036020526040902054801561223f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161189657600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611896578181036121d0575b50505060025480156121a1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161215e816002612098565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6122276121e16121f2936002612098565b90549060031b1c9283926002612098565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080612125565b5050600090565b9060018201918160005282602052604060002054801515600014612371577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611896578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116118965781810361233a575b505050805480156121a1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906122fb8282612098565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b61235a61234a6121f29386612098565b90549060031b1c92839286612098565b9055600052836020526040600020553880806122c3565b50505050600090565b806000526003602052604060002054156000146123d457600254680100000000000000008110156103cd576123bb6121f28260018594016002556002612098565b9055600254906000526003602052604060002055600190565b50600090565b600082815260018201602052604090205461223f57805490680100000000000000008210156103cd57826124186121f2846001809601855584612098565b905580549260005201602052604060002055600190565b919290156124aa5750815115612443575090565b3b1561244c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156124bd5750805190602001fd5b6124fb906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611d2a565b0390fdfea164736f6c634300081a000a",
}

var CommitteeRampABI = CommitteeRampMetaData.ABI

var CommitteeRampBin = CommitteeRampMetaData.Bin

func DeployCommitteeRamp(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig CommitteeRampDynamicConfig) (common.Address, *types.Transaction, *CommitteeRamp, error) {
	parsed, err := CommitteeRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeRampBin), backend, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitteeRamp{address: address, abi: *parsed, CommitteeRampCaller: CommitteeRampCaller{contract: contract}, CommitteeRampTransactor: CommitteeRampTransactor{contract: contract}, CommitteeRampFilterer: CommitteeRampFilterer{contract: contract}}, nil
}

type CommitteeRamp struct {
	address common.Address
	abi     abi.ABI
	CommitteeRampCaller
	CommitteeRampTransactor
	CommitteeRampFilterer
}

type CommitteeRampCaller struct {
	contract *bind.BoundContract
}

type CommitteeRampTransactor struct {
	contract *bind.BoundContract
}

type CommitteeRampFilterer struct {
	contract *bind.BoundContract
}

type CommitteeRampSession struct {
	Contract     *CommitteeRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitteeRampCallerSession struct {
	Contract *CommitteeRampCaller
	CallOpts bind.CallOpts
}

type CommitteeRampTransactorSession struct {
	Contract     *CommitteeRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitteeRampRaw struct {
	Contract *CommitteeRamp
}

type CommitteeRampCallerRaw struct {
	Contract *CommitteeRampCaller
}

type CommitteeRampTransactorRaw struct {
	Contract *CommitteeRampTransactor
}

func NewCommitteeRamp(address common.Address, backend bind.ContractBackend) (*CommitteeRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitteeRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitteeRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitteeRamp{address: address, abi: abi, CommitteeRampCaller: CommitteeRampCaller{contract: contract}, CommitteeRampTransactor: CommitteeRampTransactor{contract: contract}, CommitteeRampFilterer: CommitteeRampFilterer{contract: contract}}, nil
}

func NewCommitteeRampCaller(address common.Address, caller bind.ContractCaller) (*CommitteeRampCaller, error) {
	contract, err := bindCommitteeRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampCaller{contract: contract}, nil
}

func NewCommitteeRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeRampTransactor, error) {
	contract, err := bindCommitteeRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampTransactor{contract: contract}, nil
}

func NewCommitteeRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeRampFilterer, error) {
	contract, err := bindCommitteeRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampFilterer{contract: contract}, nil
}

func bindCommitteeRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitteeRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitteeRamp *CommitteeRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeRamp.Contract.CommitteeRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitteeRamp *CommitteeRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.CommitteeRampTransactor.contract.Transfer(opts)
}

func (_CommitteeRamp *CommitteeRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.CommitteeRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitteeRamp *CommitteeRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitteeRamp *CommitteeRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.contract.Transfer(opts)
}

func (_CommitteeRamp *CommitteeRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitteeRamp *CommitteeRampCaller) ForwardToVerifier(opts *bind.CallOpts, originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "forwardToVerifier", originalCaller, message, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) ForwardToVerifier(originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	return _CommitteeRamp.Contract.ForwardToVerifier(&_CommitteeRamp.CallOpts, originalCaller, message, arg2, arg3, arg4, arg5)
}

func (_CommitteeRamp *CommitteeRampCallerSession) ForwardToVerifier(originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	return _CommitteeRamp.Contract.ForwardToVerifier(&_CommitteeRamp.CallOpts, originalCaller, message, arg2, arg3, arg4, arg5)
}

func (_CommitteeRamp *CommitteeRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitteeRamp *CommitteeRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitteeRamp.Contract.GetDestChainConfig(&_CommitteeRamp.CallOpts, destChainSelector)
}

func (_CommitteeRamp *CommitteeRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitteeRamp.Contract.GetDestChainConfig(&_CommitteeRamp.CallOpts, destChainSelector)
}

func (_CommitteeRamp *CommitteeRampCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitteeRampDynamicConfig, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitteeRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitteeRampDynamicConfig)).(*CommitteeRampDynamicConfig)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) GetDynamicConfig() (CommitteeRampDynamicConfig, error) {
	return _CommitteeRamp.Contract.GetDynamicConfig(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCallerSession) GetDynamicConfig() (CommitteeRampDynamicConfig, error) {
	return _CommitteeRamp.Contract.GetDynamicConfig(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, arg1 uint64, arg2 ClientEVM2AnyMessage, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "getFee", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) GetFee(arg0 common.Address, arg1 uint64, arg2 ClientEVM2AnyMessage, arg3 []byte) (*big.Int, error) {
	return _CommitteeRamp.Contract.GetFee(&_CommitteeRamp.CallOpts, arg0, arg1, arg2, arg3)
}

func (_CommitteeRamp *CommitteeRampCallerSession) GetFee(arg0 common.Address, arg1 uint64, arg2 ClientEVM2AnyMessage, arg3 []byte) (*big.Int, error) {
	return _CommitteeRamp.Contract.GetFee(&_CommitteeRamp.CallOpts, arg0, arg1, arg2, arg3)
}

func (_CommitteeRamp *CommitteeRampCaller) GetSignatureConfig(opts *bind.CallOpts) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "getSignatureConfig")

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

func (_CommitteeRamp *CommitteeRampSession) GetSignatureConfig() ([]common.Address, uint8, error) {
	return _CommitteeRamp.Contract.GetSignatureConfig(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCallerSession) GetSignatureConfig() ([]common.Address, uint8, error) {
	return _CommitteeRamp.Contract.GetSignatureConfig(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) Owner() (common.Address, error) {
	return _CommitteeRamp.Contract.Owner(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCallerSession) Owner() (common.Address, error) {
	return _CommitteeRamp.Contract.Owner(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeRamp.Contract.SupportsInterface(&_CommitteeRamp.CallOpts, interfaceId)
}

func (_CommitteeRamp *CommitteeRampCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeRamp.Contract.SupportsInterface(&_CommitteeRamp.CallOpts, interfaceId)
}

func (_CommitteeRamp *CommitteeRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitteeRamp *CommitteeRampSession) TypeAndVersion() (string, error) {
	return _CommitteeRamp.Contract.TypeAndVersion(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitteeRamp.Contract.TypeAndVersion(&_CommitteeRamp.CallOpts)
}

func (_CommitteeRamp *CommitteeRampCaller) VerifyMessage(opts *bind.CallOpts, arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	var out []interface{}
	err := _CommitteeRamp.contract.Call(opts, &out, "verifyMessage", arg0, arg1, messageHash, ccvData)

	if err != nil {
		return err
	}

	return err

}

func (_CommitteeRamp *CommitteeRampSession) VerifyMessage(arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitteeRamp.Contract.VerifyMessage(&_CommitteeRamp.CallOpts, arg0, arg1, messageHash, ccvData)
}

func (_CommitteeRamp *CommitteeRampCallerSession) VerifyMessage(arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitteeRamp.Contract.VerifyMessage(&_CommitteeRamp.CallOpts, arg0, arg1, messageHash, ccvData)
}

func (_CommitteeRamp *CommitteeRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitteeRamp *CommitteeRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeRamp.Contract.AcceptOwnership(&_CommitteeRamp.TransactOpts)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeRamp.Contract.AcceptOwnership(&_CommitteeRamp.TransactOpts)
}

func (_CommitteeRamp *CommitteeRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitteeRamp *CommitteeRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.ApplyAllowlistUpdates(&_CommitteeRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.ApplyAllowlistUpdates(&_CommitteeRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeRamp *CommitteeRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitteeRamp *CommitteeRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.ApplyDestChainConfigUpdates(&_CommitteeRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.ApplyDestChainConfigUpdates(&_CommitteeRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitteeRamp *CommitteeRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeRampDynamicConfig) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitteeRamp *CommitteeRampSession) SetDynamicConfig(dynamicConfig CommitteeRampDynamicConfig) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.SetDynamicConfig(&_CommitteeRamp.TransactOpts, dynamicConfig)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) SetDynamicConfig(dynamicConfig CommitteeRampDynamicConfig) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.SetDynamicConfig(&_CommitteeRamp.TransactOpts, dynamicConfig)
}

func (_CommitteeRamp *CommitteeRampTransactor) SetSignatureConfig(opts *bind.TransactOpts, signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "setSignatureConfig", signers, threshold)
}

func (_CommitteeRamp *CommitteeRampSession) SetSignatureConfig(signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.SetSignatureConfig(&_CommitteeRamp.TransactOpts, signers, threshold)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) SetSignatureConfig(signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.SetSignatureConfig(&_CommitteeRamp.TransactOpts, signers, threshold)
}

func (_CommitteeRamp *CommitteeRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitteeRamp *CommitteeRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.TransferOwnership(&_CommitteeRamp.TransactOpts, to)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.TransferOwnership(&_CommitteeRamp.TransactOpts, to)
}

func (_CommitteeRamp *CommitteeRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitteeRamp *CommitteeRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.WithdrawFeeTokens(&_CommitteeRamp.TransactOpts, feeTokens)
}

func (_CommitteeRamp *CommitteeRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeRamp.Contract.WithdrawFeeTokens(&_CommitteeRamp.TransactOpts, feeTokens)
}

type CommitteeRampAllowListSendersAddedIterator struct {
	Event *CommitteeRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampAllowListSendersAdded)
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
		it.Event = new(CommitteeRampAllowListSendersAdded)
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

func (it *CommitteeRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampAllowListSendersAddedIterator{contract: _CommitteeRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampAllowListSendersAdded)
				if err := _CommitteeRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitteeRampAllowListSendersAdded, error) {
	event := new(CommitteeRampAllowListSendersAdded)
	if err := _CommitteeRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampAllowListSendersRemovedIterator struct {
	Event *CommitteeRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampAllowListSendersRemoved)
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
		it.Event = new(CommitteeRampAllowListSendersRemoved)
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

func (it *CommitteeRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampAllowListSendersRemovedIterator{contract: _CommitteeRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampAllowListSendersRemoved)
				if err := _CommitteeRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitteeRampAllowListSendersRemoved, error) {
	event := new(CommitteeRampAllowListSendersRemoved)
	if err := _CommitteeRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampConfigSetIterator struct {
	Event *CommitteeRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampConfigSet)
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
		it.Event = new(CommitteeRampConfigSet)
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

func (it *CommitteeRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampConfigSet struct {
	DynamicConfig CommitteeRampDynamicConfig
	Raw           types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitteeRampConfigSetIterator, error) {

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitteeRampConfigSetIterator{contract: _CommitteeRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampConfigSet)
				if err := _CommitteeRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseConfigSet(log types.Log) (*CommitteeRampConfigSet, error) {
	event := new(CommitteeRampConfigSet)
	if err := _CommitteeRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampDestChainConfigSetIterator struct {
	Event *CommitteeRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampDestChainConfigSet)
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
		it.Event = new(CommitteeRampDestChainConfigSet)
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

func (it *CommitteeRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampDestChainConfigSetIterator{contract: _CommitteeRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampDestChainConfigSet)
				if err := _CommitteeRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseDestChainConfigSet(log types.Log) (*CommitteeRampDestChainConfigSet, error) {
	event := new(CommitteeRampDestChainConfigSet)
	if err := _CommitteeRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampFeeTokenWithdrawnIterator struct {
	Event *CommitteeRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampFeeTokenWithdrawn)
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
		it.Event = new(CommitteeRampFeeTokenWithdrawn)
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

func (it *CommitteeRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeRampFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampFeeTokenWithdrawnIterator{contract: _CommitteeRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampFeeTokenWithdrawn)
				if err := _CommitteeRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitteeRampFeeTokenWithdrawn, error) {
	event := new(CommitteeRampFeeTokenWithdrawn)
	if err := _CommitteeRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampOwnershipTransferRequestedIterator struct {
	Event *CommitteeRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampOwnershipTransferRequested)
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
		it.Event = new(CommitteeRampOwnershipTransferRequested)
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

func (it *CommitteeRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampOwnershipTransferRequestedIterator{contract: _CommitteeRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampOwnershipTransferRequested)
				if err := _CommitteeRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitteeRampOwnershipTransferRequested, error) {
	event := new(CommitteeRampOwnershipTransferRequested)
	if err := _CommitteeRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampOwnershipTransferredIterator struct {
	Event *CommitteeRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampOwnershipTransferred)
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
		it.Event = new(CommitteeRampOwnershipTransferred)
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

func (it *CommitteeRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeRampOwnershipTransferredIterator{contract: _CommitteeRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampOwnershipTransferred)
				if err := _CommitteeRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitteeRampOwnershipTransferred, error) {
	event := new(CommitteeRampOwnershipTransferred)
	if err := _CommitteeRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeRampSignatureConfigSetIterator struct {
	Event *CommitteeRampSignatureConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeRampSignatureConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeRampSignatureConfigSet)
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
		it.Event = new(CommitteeRampSignatureConfigSet)
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

func (it *CommitteeRampSignatureConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeRampSignatureConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeRampSignatureConfigSet struct {
	Signers   []common.Address
	Threshold uint8
	Raw       types.Log
}

func (_CommitteeRamp *CommitteeRampFilterer) FilterSignatureConfigSet(opts *bind.FilterOpts) (*CommitteeRampSignatureConfigSetIterator, error) {

	logs, sub, err := _CommitteeRamp.contract.FilterLogs(opts, "SignatureConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitteeRampSignatureConfigSetIterator{contract: _CommitteeRamp.contract, event: "SignatureConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeRamp *CommitteeRampFilterer) WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampSignatureConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitteeRamp.contract.WatchLogs(opts, "SignatureConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeRampSignatureConfigSet)
				if err := _CommitteeRamp.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
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

func (_CommitteeRamp *CommitteeRampFilterer) ParseSignatureConfigSet(log types.Log) (*CommitteeRampSignatureConfigSet, error) {
	event := new(CommitteeRampSignatureConfigSet)
	if err := _CommitteeRamp.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}

func (CommitteeRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitteeRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitteeRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d")
}

func (CommitteeRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitteeRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitteeRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitteeRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CommitteeRampSignatureConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc2e12b820aa2dc1a1673e9f59d1d809598d1041a90baccc742b7de5e5d2418a8")
}

func (_CommitteeRamp *CommitteeRamp) Address() common.Address {
	return _CommitteeRamp.address
}

type CommitteeRampInterface interface {
	ForwardToVerifier(opts *bind.CallOpts, originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitteeRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, arg1 uint64, arg2 ClientEVM2AnyMessage, arg3 []byte) (*big.Int, error)

	GetSignatureConfig(opts *bind.CallOpts) ([]common.Address, uint8, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyMessage(opts *bind.CallOpts, arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeRampDynamicConfig) (*types.Transaction, error)

	SetSignatureConfig(opts *bind.TransactOpts, signers []common.Address, threshold uint8) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitteeRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitteeRampAllowListSendersRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitteeRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitteeRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitteeRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitteeRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitteeRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitteeRampOwnershipTransferred, error)

	FilterSignatureConfigSet(opts *bind.FilterOpts) (*CommitteeRampSignatureConfigSetIterator, error)

	WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeRampSignatureConfigSet) (event.Subscription, error)

	ParseSignatureConfigSet(log types.Log) (*CommitteeRampSignatureConfigSet, error)

	Address() common.Address
}
