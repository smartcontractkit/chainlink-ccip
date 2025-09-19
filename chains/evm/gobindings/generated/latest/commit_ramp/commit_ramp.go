// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_ramp

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

type CommitRampDynamicConfig struct {
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

var CommitRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"originalCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignatureConfig\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureConfigSet\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonOrderedOrNonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]}]",
	Bin: "0x60c0604052346101af57604051601f6126cb38819003918201601f19168301916001600160401b038311848410176101b4578084926060946040528339810103126101af57604051600090606081016001600160401b0381118282101761019b5760405261006c836101ca565b815261008d604061007f602086016101ca565b9460208401958652016101ca565b9160408201928352331561018c57600180546001600160a01b031916331790554660805281516001600160a01b031615801561017a575b61016b578151600680546001600160a01b03199081166001600160a01b039384169081179092558651600780548316918516919091179055855160088054909216908416179055604080519182528651831660208301528551909216918101919091527fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d90606090a16040516124ec90816101df8239608051816114c2015260a051815050f35b6306b7c75960e31b8152600490fd5b5083516001600160a01b0316156100c4565b639b15e16f60e01b8152600490fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101af5756fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611bce57508063181f5a7714611b515780631a60a71a1461192257806334d560e41461171257806358bfa40a146113ea5780635cb80c5d146110fd5780636def4ce71461101f5780636ed0e21714610f7357806371c5c2ba14610c555780637437ff9f14610b6157806379ba509714610a785780638da5cb5b14610a26578063a8dd2df2146107f6578063b2d6d66b146105d0578063c9b146b3146101c45763f2fde38b146100cf57600080fd5b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5773ffffffffffffffffffffffffffffffffffffffff61011b611d80565b610123611ff2565b1633811461019557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf57610213903690600401611e74565b73ffffffffffffffffffffffffffffffffffffffff600154163303610586575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610584576000938060051b840135828112156105805784019160808336031261058057604051946080860186811067ffffffffffffffff821117610553576040526102ab84611ea5565b86526102b960208501611fa2565b9660208701978852604085013567ffffffffffffffff811161054f576102e29036908701611f25565b9460408801958652606081013567ffffffffffffffff811161054b5761030a91369101611f25565b946060880195865267ffffffffffffffff885116825260056020526040822098511515610362818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b81515161041f575b50959760010195505b845180518210156103b257906103ab73ffffffffffffffffffffffffffffffffffffffff6103a383600195611faf565b511688612226565b5001610373565b50509594909350600192519081516103d0575b50500192919061025d565b61041567ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190611eba565b0390a285806103c5565b989395929094979896919660001461051457600184019591875b865180518210156104b6576104638273ffffffffffffffffffffffffffffffffffffffff92611faf565b5116801561047f57906104786001928a6123ba565b5001610439565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc328161050a67ffffffffffffffff8b51169251604051918291602083526020830190611eba565b0390a2908961036a565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60085416330315610233577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf5761061f903690600401611f25565b6024359060ff8216918281036101bf57610637611ff2565b821580156107ec575b610737575b600254156106cf576000600254156106a257600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5461069c9073ffffffffffffffffffffffffffffffffffffffff16612090565b50610645565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60005b825181101561078b5773ffffffffffffffffffffffffffffffffffffffff6106fa8285611faf565b5116156107615761072a73ffffffffffffffffffffffffffffffffffffffff6107238386611faf565b511661235a565b15610737576001016106d2565b7f12823a5e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50907fc2e12b820aa2dc1a1673e9f59d1d809598d1041a90baccc742b7de5e5d2418a8927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060045416176004556107e760405192839283611f04565b0390a1005b5081518311610640565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf57366023820112156101bf5780600401359067ffffffffffffffff82116101bf576024606083028201013681116101bf5761086f611ff2565b61087883611e2e565b926108866040519485611ca6565b835260009160240190602084015b8183106109ba57505050805b82518110156109b6576108b38184611faf565b5167ffffffffffffffff60206108c98487611faf565b5101511690811561098a57818452600560205260408085208251815493830151151560ff9081167fffffffffffffffffffffff000000000000000000000000000000000000000000909516600883901b74ffffffffffffffffffffffffffffffffffffffff00161794909417825560019594937f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9392541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a2016108a0565b602484837fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b606083360312610a22576040516109d081611c8a565b833573ffffffffffffffffffffffffffffffffffffffff81168103610580578152606091602091610a02868401611ea5565b83820152610a1260408701611fa2565b6040820152815201920191610894565b8380fd5b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760005473ffffffffffffffffffffffffffffffffffffffff81163303610b37577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600060408051610b9f81611c8a565b8281528260208201520152610c51604051610bb981611c8a565b73ffffffffffffffffffffffffffffffffffffffff60065416815273ffffffffffffffffffffffffffffffffffffffff60075416602082015273ffffffffffffffffffffffffffffffffffffffff60085416604082015260405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390f35b346101bf5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57610c8c611d80565b6024359067ffffffffffffffff82116101bf578136036101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101bf57610cd5611da3565b5060a43567ffffffffffffffff81116101bf57610cf6903690600401611e46565b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60c484013591018112156101bf57820160048101359067ffffffffffffffff82116101bf576024019080360382136101bf57602491357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169160148110610f3e575b505060601c92013567ffffffffffffffff81168091036101bf57806000526005602052604060002091825491604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff8760081c165afa908115610f3257600091610ecc575b5073ffffffffffffffffffffffffffffffffffffffff809116911603610e9e5760ff16610e54575b6040516020610e3c8183611ca6565b60008252610c51604051928284938452830190611d21565b60008281526002909101602052604090205415610e715780610e2d565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011610f2a575b81610ee560209383611ca6565b8101031261054f57519073ffffffffffffffffffffffffffffffffffffffff82168203610f27575073ffffffffffffffffffffffffffffffffffffffff610e05565b80fd5b3d9150610ed8565b6040513d6000823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168480610d7c565b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760405180816020600254928381520160026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9260005b818110611006575050610ff092500382611ca6565b60ff6004541690610c5160405192839283611f04565b8454835260019485019486945060209093019201610fdb565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81168091036101bf57600052600560205260406000206001815491019060405191826020825491828152019160005260206000209060005b8181106110e75773ffffffffffffffffffffffffffffffffffffffff85610c51886110bf81890382611ca6565b604051938360ff8695161515855260081c166020840152606060408401526060830190611eba565b8254845260209093019260019283019201611092565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5760043567ffffffffffffffff81116101bf5761114c903690600401611e74565b73ffffffffffffffffffffffffffffffffffffffff6007541660005b82811015610584576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff811680910361054b576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156113df5790859185916113ab575b50806111f2575b5050506001915001611168565b604051946112ac60208701967fa9059cbb0000000000000000000000000000000000000000000000000000000088528460248201528360448201526044815261123c606482611ca6565b82806040998a519361124e8c86611ca6565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d156113a3573d9061128f82611ce7565b9161129c8b519384611ca6565b82523d85602084013e5b8761240f565b8051806112eb575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a38583816111e5565b8192939596979450906020918101031261054f576020015190811591821503610f2757506113205792919084908880806112b4565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060906112a6565b9150506020813d82116113d7575b816113c660209383611ca6565b81010312610a2257849051886111de565b3d91506113b9565b6040513d86823e3d90fd5b346101bf5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57611421611d80565b5060243567ffffffffffffffff81116101bf577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc61016091360301126101bf5760443560643567ffffffffffffffff81116101bf57611484903690600401611e46565b600281106116e857806002116101bf578135918260f01c91826002018060021161165e578082106116e857116101bf576002019260025415610737577f00000000000000000000000000000000000000000000000000000000000000004681036116b7575060ff60045416809360f61c1061168d57600091825b84841061150757005b8360061b8481046040148515171561165e57602081019081811161165e5761153a6115348383878c611f8a565b9061203d565b90600092604082018092116116315760209261155e61153486946080948a8f611f8a565b60405191898352601b868401526040830152606082015282805260015afa156116255773ffffffffffffffffffffffffffffffffffffffff8151169182825260036020526040822054156115fd5773ffffffffffffffffffffffffffffffffffffffff168211156115d557506001909301926114fe565b807fb70ad94b0000000000000000000000000000000000000000000000000000000060049252fd5b6004827fca31867a000000000000000000000000000000000000000000000000000000008152fd5b604051903d90823e3d90fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7fbba6473c0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101bf5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600060405161174f81611c8a565b611757611d80565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361054b57602082019081526044359073ffffffffffffffffffffffffffffffffffffffff82168203610a2257604083019182526117b0611ff2565b73ffffffffffffffffffffffffffffffffffffffff835116158015611903575b6118db579173ffffffffffffffffffffffffffffffffffffffff6118d592817fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d95818551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600854161760085560405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff815116156117d0565b346101bf5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57611959611d80565b5060243567ffffffffffffffff81116101bf5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101bf576040519060a0820182811067ffffffffffffffff821117611b2257604052806004013567ffffffffffffffff81116101bf576119da9060043691840101611de7565b8252602481013567ffffffffffffffff81116101bf57611a009060043691840101611de7565b6020830152604481013567ffffffffffffffff81116101bf578101366023820112156101bf576004810135611a3481611e2e565b91611a426040519384611ca6565b818352602060048185019360061b83010101903682116101bf57602401915b818310611ad6575050506040830152611a7c60648201611dc6565b6060830152608481013567ffffffffffffffff81116101bf576080916004611aa79236920101611de7565b91015260443567ffffffffffffffff81116101bf57611aca903690600401611de7565b50602060405160008152f35b6040833603126101bf5760405190604082019082821067ffffffffffffffff831117611b22576040926020928452611b0d86611dc6565b81528286013583820152815201920191611a61565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101bf5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57610c516040805190611b928183611ca6565b601482527f436f6d6d697452616d7020312e372e302d646576000000000000000000000000602083015251918291602083526020830190611d21565b346101bf5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101bf57817f58bfa40a0000000000000000000000000000000000000000000000000000000060209314908115611c60575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611c59565b6060810190811067ffffffffffffffff821117611b2257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611b2257604052565b67ffffffffffffffff8111611b2257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110611d6b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611d2c565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b359073ffffffffffffffffffffffffffffffffffffffff821682036101bf57565b81601f820112156101bf57803590611dfe82611ce7565b92611e0c6040519485611ca6565b828452602083830101116101bf57816000926020809301838601378301015290565b67ffffffffffffffff8111611b225760051b60200190565b9181601f840112156101bf5782359167ffffffffffffffff83116101bf57602083818601950101116101bf57565b9181601f840112156101bf5782359167ffffffffffffffff83116101bf576020808501948460051b0101116101bf57565b359067ffffffffffffffff821682036101bf57565b906020808351928381520192019060005b818110611ed85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611ecb565b9060ff611f1e602092959495604085526040850190611eba565b9416910152565b9080601f830112156101bf578135611f3c81611e2e565b92611f4a6040519485611ca6565b81845260208085019260051b8201019283116101bf57602001905b828210611f725750505090565b60208091611f7f84611dc6565b815201910190611f65565b909392938483116101bf5784116101bf578101920390565b359081151582036101bf57565b8051821015611fc35760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361201357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b35906020811061204b575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b8054821015611fc35760005260206000200190600090565b600081815260036020526040902054801561221f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161165e57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161165e578181036121b0575b5050506002548015612181577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161213e816002612078565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6122076121c16121d2936002612078565b90549060031b1c9283926002612078565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080612105565b5050600090565b9060018201918160005282602052604060002054801515600014612351577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161165e578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161165e5781810361231a575b50505080548015612181577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906122db8282612078565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b61233a61232a6121d29386612078565b90549060031b1c92839286612078565b9055600052836020526040600020553880806122a3565b50505050600090565b806000526003602052604060002054156000146123b45760025468010000000000000000811015611b225761239b6121d28260018594016002556002612078565b9055600254906000526003602052604060002055600190565b50600090565b600082815260018201602052604090205461221f5780549068010000000000000000821015611b2257826123f86121d2846001809601855584612078565b905580549260005201602052604060002055600190565b9192901561248a5750815115612423575090565b3b1561242c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561249d5750805190602001fd5b6124db906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611d21565b0390fdfea164736f6c634300081a000a",
}

var CommitRampABI = CommitRampMetaData.ABI

var CommitRampBin = CommitRampMetaData.Bin

func DeployCommitRamp(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig CommitRampDynamicConfig) (common.Address, *types.Transaction, *CommitRamp, error) {
	parsed, err := CommitRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitRampBin), backend, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitRamp{address: address, abi: *parsed, CommitRampCaller: CommitRampCaller{contract: contract}, CommitRampTransactor: CommitRampTransactor{contract: contract}, CommitRampFilterer: CommitRampFilterer{contract: contract}}, nil
}

type CommitRamp struct {
	address common.Address
	abi     abi.ABI
	CommitRampCaller
	CommitRampTransactor
	CommitRampFilterer
}

type CommitRampCaller struct {
	contract *bind.BoundContract
}

type CommitRampTransactor struct {
	contract *bind.BoundContract
}

type CommitRampFilterer struct {
	contract *bind.BoundContract
}

type CommitRampSession struct {
	Contract     *CommitRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitRampCallerSession struct {
	Contract *CommitRampCaller
	CallOpts bind.CallOpts
}

type CommitRampTransactorSession struct {
	Contract     *CommitRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitRampRaw struct {
	Contract *CommitRamp
}

type CommitRampCallerRaw struct {
	Contract *CommitRampCaller
}

type CommitRampTransactorRaw struct {
	Contract *CommitRampTransactor
}

func NewCommitRamp(address common.Address, backend bind.ContractBackend) (*CommitRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitRamp{address: address, abi: abi, CommitRampCaller: CommitRampCaller{contract: contract}, CommitRampTransactor: CommitRampTransactor{contract: contract}, CommitRampFilterer: CommitRampFilterer{contract: contract}}, nil
}

func NewCommitRampCaller(address common.Address, caller bind.ContractCaller) (*CommitRampCaller, error) {
	contract, err := bindCommitRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRampCaller{contract: contract}, nil
}

func NewCommitRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitRampTransactor, error) {
	contract, err := bindCommitRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRampTransactor{contract: contract}, nil
}

func NewCommitRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitRampFilterer, error) {
	contract, err := bindCommitRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitRampFilterer{contract: contract}, nil
}

func bindCommitRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitRamp *CommitRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitRamp.Contract.CommitRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitRamp *CommitRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitRamp.Contract.CommitRampTransactor.contract.Transfer(opts)
}

func (_CommitRamp *CommitRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitRamp.Contract.CommitRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitRamp *CommitRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitRamp *CommitRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitRamp.Contract.contract.Transfer(opts)
}

func (_CommitRamp *CommitRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitRamp *CommitRampCaller) ForwardToVerifier(opts *bind.CallOpts, originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "forwardToVerifier", originalCaller, message, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CommitRamp *CommitRampSession) ForwardToVerifier(originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	return _CommitRamp.Contract.ForwardToVerifier(&_CommitRamp.CallOpts, originalCaller, message, arg2, arg3, arg4, arg5)
}

func (_CommitRamp *CommitRampCallerSession) ForwardToVerifier(originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error) {
	return _CommitRamp.Contract.ForwardToVerifier(&_CommitRamp.CallOpts, originalCaller, message, arg2, arg3, arg4, arg5)
}

func (_CommitRamp *CommitRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitRamp *CommitRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitRamp.Contract.GetDestChainConfig(&_CommitRamp.CallOpts, destChainSelector)
}

func (_CommitRamp *CommitRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitRamp.Contract.GetDestChainConfig(&_CommitRamp.CallOpts, destChainSelector)
}

func (_CommitRamp *CommitRampCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitRampDynamicConfig, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitRampDynamicConfig)).(*CommitRampDynamicConfig)

	return out0, err

}

func (_CommitRamp *CommitRampSession) GetDynamicConfig() (CommitRampDynamicConfig, error) {
	return _CommitRamp.Contract.GetDynamicConfig(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCallerSession) GetDynamicConfig() (CommitRampDynamicConfig, error) {
	return _CommitRamp.Contract.GetDynamicConfig(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "getFee", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitRamp *CommitRampSession) GetFee(arg0 common.Address, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitRamp.Contract.GetFee(&_CommitRamp.CallOpts, arg0, arg1, arg2)
}

func (_CommitRamp *CommitRampCallerSession) GetFee(arg0 common.Address, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitRamp.Contract.GetFee(&_CommitRamp.CallOpts, arg0, arg1, arg2)
}

func (_CommitRamp *CommitRampCaller) GetSignatureConfig(opts *bind.CallOpts) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "getSignatureConfig")

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

func (_CommitRamp *CommitRampSession) GetSignatureConfig() ([]common.Address, uint8, error) {
	return _CommitRamp.Contract.GetSignatureConfig(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCallerSession) GetSignatureConfig() ([]common.Address, uint8, error) {
	return _CommitRamp.Contract.GetSignatureConfig(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitRamp *CommitRampSession) Owner() (common.Address, error) {
	return _CommitRamp.Contract.Owner(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCallerSession) Owner() (common.Address, error) {
	return _CommitRamp.Contract.Owner(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CommitRamp *CommitRampSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitRamp.Contract.SupportsInterface(&_CommitRamp.CallOpts, interfaceId)
}

func (_CommitRamp *CommitRampCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitRamp.Contract.SupportsInterface(&_CommitRamp.CallOpts, interfaceId)
}

func (_CommitRamp *CommitRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitRamp *CommitRampSession) TypeAndVersion() (string, error) {
	return _CommitRamp.Contract.TypeAndVersion(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitRamp.Contract.TypeAndVersion(&_CommitRamp.CallOpts)
}

func (_CommitRamp *CommitRampCaller) VerifyMessage(opts *bind.CallOpts, arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	var out []interface{}
	err := _CommitRamp.contract.Call(opts, &out, "verifyMessage", arg0, arg1, messageHash, ccvData)

	if err != nil {
		return err
	}

	return err

}

func (_CommitRamp *CommitRampSession) VerifyMessage(arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitRamp.Contract.VerifyMessage(&_CommitRamp.CallOpts, arg0, arg1, messageHash, ccvData)
}

func (_CommitRamp *CommitRampCallerSession) VerifyMessage(arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitRamp.Contract.VerifyMessage(&_CommitRamp.CallOpts, arg0, arg1, messageHash, ccvData)
}

func (_CommitRamp *CommitRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitRamp *CommitRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitRamp.Contract.AcceptOwnership(&_CommitRamp.TransactOpts)
}

func (_CommitRamp *CommitRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitRamp.Contract.AcceptOwnership(&_CommitRamp.TransactOpts)
}

func (_CommitRamp *CommitRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitRamp *CommitRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.Contract.ApplyAllowlistUpdates(&_CommitRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitRamp *CommitRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.Contract.ApplyAllowlistUpdates(&_CommitRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitRamp *CommitRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitRamp *CommitRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.Contract.ApplyDestChainConfigUpdates(&_CommitRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitRamp *CommitRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitRamp.Contract.ApplyDestChainConfigUpdates(&_CommitRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitRamp *CommitRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitRampDynamicConfig) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitRamp *CommitRampSession) SetDynamicConfig(dynamicConfig CommitRampDynamicConfig) (*types.Transaction, error) {
	return _CommitRamp.Contract.SetDynamicConfig(&_CommitRamp.TransactOpts, dynamicConfig)
}

func (_CommitRamp *CommitRampTransactorSession) SetDynamicConfig(dynamicConfig CommitRampDynamicConfig) (*types.Transaction, error) {
	return _CommitRamp.Contract.SetDynamicConfig(&_CommitRamp.TransactOpts, dynamicConfig)
}

func (_CommitRamp *CommitRampTransactor) SetSignatureConfig(opts *bind.TransactOpts, signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "setSignatureConfig", signers, threshold)
}

func (_CommitRamp *CommitRampSession) SetSignatureConfig(signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitRamp.Contract.SetSignatureConfig(&_CommitRamp.TransactOpts, signers, threshold)
}

func (_CommitRamp *CommitRampTransactorSession) SetSignatureConfig(signers []common.Address, threshold uint8) (*types.Transaction, error) {
	return _CommitRamp.Contract.SetSignatureConfig(&_CommitRamp.TransactOpts, signers, threshold)
}

func (_CommitRamp *CommitRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitRamp *CommitRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitRamp.Contract.TransferOwnership(&_CommitRamp.TransactOpts, to)
}

func (_CommitRamp *CommitRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitRamp.Contract.TransferOwnership(&_CommitRamp.TransactOpts, to)
}

func (_CommitRamp *CommitRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitRamp *CommitRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitRamp.Contract.WithdrawFeeTokens(&_CommitRamp.TransactOpts, feeTokens)
}

func (_CommitRamp *CommitRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitRamp.Contract.WithdrawFeeTokens(&_CommitRamp.TransactOpts, feeTokens)
}

type CommitRampAllowListSendersAddedIterator struct {
	Event *CommitRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampAllowListSendersAdded)
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
		it.Event = new(CommitRampAllowListSendersAdded)
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

func (it *CommitRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampAllowListSendersAddedIterator{contract: _CommitRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampAllowListSendersAdded)
				if err := _CommitRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitRampAllowListSendersAdded, error) {
	event := new(CommitRampAllowListSendersAdded)
	if err := _CommitRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampAllowListSendersRemovedIterator struct {
	Event *CommitRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampAllowListSendersRemoved)
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
		it.Event = new(CommitRampAllowListSendersRemoved)
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

func (it *CommitRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampAllowListSendersRemovedIterator{contract: _CommitRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampAllowListSendersRemoved)
				if err := _CommitRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitRampAllowListSendersRemoved, error) {
	event := new(CommitRampAllowListSendersRemoved)
	if err := _CommitRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampConfigSetIterator struct {
	Event *CommitRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampConfigSet)
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
		it.Event = new(CommitRampConfigSet)
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

func (it *CommitRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampConfigSet struct {
	DynamicConfig CommitRampDynamicConfig
	Raw           types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitRampConfigSetIterator, error) {

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitRampConfigSetIterator{contract: _CommitRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampConfigSet)
				if err := _CommitRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseConfigSet(log types.Log) (*CommitRampConfigSet, error) {
	event := new(CommitRampConfigSet)
	if err := _CommitRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampDestChainConfigSetIterator struct {
	Event *CommitRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampDestChainConfigSet)
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
		it.Event = new(CommitRampDestChainConfigSet)
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

func (it *CommitRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampDestChainConfigSetIterator{contract: _CommitRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampDestChainConfigSet)
				if err := _CommitRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseDestChainConfigSet(log types.Log) (*CommitRampDestChainConfigSet, error) {
	event := new(CommitRampDestChainConfigSet)
	if err := _CommitRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampFeeTokenWithdrawnIterator struct {
	Event *CommitRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampFeeTokenWithdrawn)
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
		it.Event = new(CommitRampFeeTokenWithdrawn)
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

func (it *CommitRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitRampFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampFeeTokenWithdrawnIterator{contract: _CommitRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampFeeTokenWithdrawn)
				if err := _CommitRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitRampFeeTokenWithdrawn, error) {
	event := new(CommitRampFeeTokenWithdrawn)
	if err := _CommitRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampOwnershipTransferRequestedIterator struct {
	Event *CommitRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampOwnershipTransferRequested)
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
		it.Event = new(CommitRampOwnershipTransferRequested)
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

func (it *CommitRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampOwnershipTransferRequestedIterator{contract: _CommitRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampOwnershipTransferRequested)
				if err := _CommitRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitRampOwnershipTransferRequested, error) {
	event := new(CommitRampOwnershipTransferRequested)
	if err := _CommitRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampOwnershipTransferredIterator struct {
	Event *CommitRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampOwnershipTransferred)
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
		it.Event = new(CommitRampOwnershipTransferred)
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

func (it *CommitRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitRampOwnershipTransferredIterator{contract: _CommitRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampOwnershipTransferred)
				if err := _CommitRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitRampOwnershipTransferred, error) {
	event := new(CommitRampOwnershipTransferred)
	if err := _CommitRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitRampSignatureConfigSetIterator struct {
	Event *CommitRampSignatureConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitRampSignatureConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRampSignatureConfigSet)
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
		it.Event = new(CommitRampSignatureConfigSet)
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

func (it *CommitRampSignatureConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitRampSignatureConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitRampSignatureConfigSet struct {
	Signers   []common.Address
	Threshold uint8
	Raw       types.Log
}

func (_CommitRamp *CommitRampFilterer) FilterSignatureConfigSet(opts *bind.FilterOpts) (*CommitRampSignatureConfigSetIterator, error) {

	logs, sub, err := _CommitRamp.contract.FilterLogs(opts, "SignatureConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitRampSignatureConfigSetIterator{contract: _CommitRamp.contract, event: "SignatureConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitRamp *CommitRampFilterer) WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampSignatureConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitRamp.contract.WatchLogs(opts, "SignatureConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitRampSignatureConfigSet)
				if err := _CommitRamp.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
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

func (_CommitRamp *CommitRampFilterer) ParseSignatureConfigSet(log types.Log) (*CommitRampSignatureConfigSet, error) {
	event := new(CommitRampSignatureConfigSet)
	if err := _CommitRamp.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
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

func (CommitRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d")
}

func (CommitRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CommitRampSignatureConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc2e12b820aa2dc1a1673e9f59d1d809598d1041a90baccc742b7de5e5d2418a8")
}

func (_CommitRamp *CommitRamp) Address() common.Address {
	return _CommitRamp.address
}

type CommitRampInterface interface {
	ForwardToVerifier(opts *bind.CallOpts, originalCaller common.Address, message MessageV1CodecMessageV1, arg2 [32]byte, arg3 common.Address, arg4 *big.Int, arg5 []byte) ([]byte, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error)

	GetSignatureConfig(opts *bind.CallOpts) ([]common.Address, uint8, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyMessage(opts *bind.CallOpts, arg0 common.Address, arg1 MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitRampDynamicConfig) (*types.Transaction, error)

	SetSignatureConfig(opts *bind.TransactOpts, signers []common.Address, threshold uint8) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitRampAllowListSendersRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitRampOwnershipTransferred, error)

	FilterSignatureConfigSet(opts *bind.FilterOpts) (*CommitRampSignatureConfigSetIterator, error)

	WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitRampSignatureConfigSet) (event.Subscription, error)

	ParseSignatureConfigSet(log types.Log) (*CommitRampSignatureConfigSet, error)

	Address() common.Address
}
