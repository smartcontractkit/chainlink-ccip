// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_onramp

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

type CommitOnRampDynamicConfig struct {
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

var CommitOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifierArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c0604052346101a857604051601f6121f038819003918201601f19168301916001600160401b038311848410176101ad578084926060946040528339810103126101a857604051600090606081016001600160401b038111828210176101945760405261006c836101c3565b815261008d604061007f602086016101c3565b9460208401958652016101c3565b9160408201928352331561018557600180546001600160a01b0319163317905581516001600160a01b0316158015610173575b610164578151600380546001600160a01b03199081166001600160a01b039384169081179092558651600480548316918516919091179055855160058054909216908416179055604080519182528651831660208301528551909216918101919091527fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d90606090a160405161201890816101d8823960805181505060a051815050f35b6306b7c75960e31b8152600490fd5b5083516001600160a01b0316156100c0565b639b15e16f60e01b8152600490fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101a85756fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461182257806334d560e41461162a5780634a7597b51461122b5780635cb80c5d14610f3f5780636def4ce714610e645780637437ff9f14610d7057806379ba509714610c875780638da5cb5b14610c35578063a8dd2df214610a05578063c527f200146105a1578063c9b146b3146101955763f2fde38b146100a057600080fd5b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905773ffffffffffffffffffffffffffffffffffffffff6100ec61199c565b6100f4611c0d565b1633811461016657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff8111610190576101e4903690600401611a8e565b73ffffffffffffffffffffffffffffffffffffffff600154163303610557575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610555576000938060051b840135828112156105515784019160808336031261055157604051946080860186811067ffffffffffffffff8211176105245760405261027c84611a1a565b865261028a60208501611b09565b9660208701978852604085013567ffffffffffffffff8111610520576102b39036908701611c9b565b9460408801958652606081013567ffffffffffffffff811161051c576102db91369101611c9b565b946060880195865267ffffffffffffffff885116825260026020526040822098511515610333818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b8151516103f0575b50959760010195505b84518051821015610383579061037c73ffffffffffffffffffffffffffffffffffffffff61037483600195611c58565b511688611da9565b5001610344565b50509594909350600192519081516103a1575b50500192919061022e565b6103e667ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190611abf565b0390a28580610396565b98939592909497989691966000146104e557600184019591875b86518051821015610487576104348273ffffffffffffffffffffffffffffffffffffffff92611c58565b5116801561045057906104496001928a611d18565b500161040a565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816104db67ffffffffffffffff8b51169251604051918291602083526020830190611abf565b0390a2908961033b565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60055416330315610204577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101905760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff811161019057806004016101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101905761061b6119bf565b9160843567ffffffffffffffff811161019057366023820112156101905780600401359067ffffffffffffffff82116101905736602483830101116101905761066760c4840185611b16565b90357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811691601481106109d0575b505060601c93602484019467ffffffffffffffff6106b387611b67565b16806000526002602052604060002090815490604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff8660081c165afa9081156108e65760009161096a575b5073ffffffffffffffffffffffffffffffffffffffff16330361093c5760ff166108f2575b5050926108196000959367ffffffffffffffff936107e961079173ffffffffffffffffffffffffffffffffffffffff9a9860e46107898d600354169a611b67565b970190611b16565b9290936040519b8c9a8b998a997f3a49bb49000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152606435604488015260a06064880152602460a488019201611bce565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc858403016084860152611bce565b03915afa80156108e657610877575b60405160208152602081817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f606051808486015261086d81604087016080611936565b0116820101030190f35b3d806000833e61088781836118bb565b81016080828203126101905761089f60208301611b7c565b50604082015167ffffffffffffffff811161019057816108c0918401611b89565b50606082015167ffffffffffffffff8111610190576108df9201611b89565b5080610828565b6040513d6000823e3d90fd5b6000828152600290910160205260409020541561090f5780610748565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d6020116109c8575b81610983602093836118bb565b8101031261052057519073ffffffffffffffffffffffffffffffffffffffff821682036109c5575073ffffffffffffffffffffffffffffffffffffffff610723565b80fd5b3d9150610976565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168680610696565b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff811161019057366023820112156101905780600401359067ffffffffffffffff82116101905760246060830282010136811161019057610a7e611c0d565b610a8783611a76565b92610a9560405194856118bb565b835260009160240190602084015b818310610bc957505050805b8251811015610bc557610ac28184611c58565b5167ffffffffffffffff6020610ad88487611c58565b51015116908115610b9957818452600260205260408085208251815493830151151560ff9081167fffffffffffffffffffffff000000000000000000000000000000000000000000909516600883901b74ffffffffffffffffffffffffffffffffffffffff00161794909417825560019594937f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9392541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a201610aaf565b602484837fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b606083360312610c3157604051610bdf8161189f565b833573ffffffffffffffffffffffffffffffffffffffff81168103610551578152606091602091610c11868401611a1a565b83820152610c2160408701611b09565b6040820152815201920191610aa3565b8380fd5b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760005473ffffffffffffffffffffffffffffffffffffffff81163303610d46577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057600060408051610dae8161189f565b8281528260208201520152610e60604051610dc88161189f565b73ffffffffffffffffffffffffffffffffffffffff60035416815273ffffffffffffffffffffffffffffffffffffffff60045416602082015273ffffffffffffffffffffffffffffffffffffffff60055416604082015260405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390f35b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905767ffffffffffffffff610ea4611a03565b16600052600260205260406000206001815491019060405191826020825491828152019160005260206000209060005b818110610f295773ffffffffffffffffffffffffffffffffffffffff85610e6088610f01818903826118bb565b604051938360ff8695161515855260081c166020840152606060408401526060830190611abf565b8254845260209093019260019283019201610ed4565b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff811161019057610f8e903690600401611a8e565b73ffffffffffffffffffffffffffffffffffffffff6004541660005b82811015610555576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff811680910361051c576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156112205790859185916111ec575b5080611034575b5050506001915001610faa565b604051946110ee60208701967fa9059cbb0000000000000000000000000000000000000000000000000000000088528460248201528360448201526044815261107e6064826118bb565b82806040998a51936110908c866118bb565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d156111e4573d906110d1826118fc565b916110de8b5193846118bb565b82523d85602084013e5b87611f3b565b9081518061112e575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381611027565b829395969794509160209192810103126109c55750602061114f9101611b7c565b156111615792919084908880806110f7565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060906110e8565b9150506020813d8211611218575b81611207602093836118bb565b81010312610c315784905188611020565b3d91506111fa565b6040513d86823e3d90fd5b346101905760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057611262611a03565b60243567ffffffffffffffff81116101905760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610190576040519060a0820182811067ffffffffffffffff8211176115fb57604052806004013567ffffffffffffffff8111610190576112e29060043691840101611a2f565b8252602481013567ffffffffffffffff8111610190576113089060043691840101611a2f565b60208301908152604482013567ffffffffffffffff8111610190578201913660238401121561019057600483013561133f81611a76565b9361134d60405195866118bb565b818552602060048187019360061b830101019036821161019057602401915b8183106115af5750505060408401928352611389606482016119e2565b906060850191825260848101359067ffffffffffffffff82116101905760046113b59236920101611a2f565b91608085019283526044359067ffffffffffffffff821161019057611486906113ef67ffffffffffffffff98969594933690600401611a2f565b5061145573ffffffffffffffffffffffffffffffffffffffff6003541697604051998a987fd8694ccd000000000000000000000000000000000000000000000000000000008a52166004890152604060248901525160a0604489015260e4880190611959565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc878303016064880152611959565b9251927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8582030160848601526020808551928381520194019060005b8181106115745750505061152460209593859373ffffffffffffffffffffffffffffffffffffffff8594511660a4850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8483030160c4850152611959565b03915afa80156108e657600090611541575b602090604051908152f35b506020813d60201161156c575b8161155b602093836118bb565b810103126101905760209051611536565b3d915061154e565b8251805173ffffffffffffffffffffffffffffffffffffffff16875260209081015181880152899750604090960195909201916001016114c3565b6040833603126101905760405190604082019082821067ffffffffffffffff8311176115fb5760409260209284526115e6866119e2565b8152828601358382015281520192019161136c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101905760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760006040516116678161189f565b61166f61199c565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361051c57602082019081526116a06119bf565b90604083019182526116b0611c0d565b73ffffffffffffffffffffffffffffffffffffffff835116158015611803575b6117db579173ffffffffffffffffffffffffffffffffffffffff6117d592817fe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d95818551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055560405191829182919091604073ffffffffffffffffffffffffffffffffffffffff816060840195828151168552826020820151166020860152015116910152565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff815116156116d0565b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057610e60604080519061186381836118bb565b601682527f436f6d6d69744f6e52616d7020312e372e302d64657600000000000000000000602083015251918291602083526020830190611959565b6060810190811067ffffffffffffffff8211176115fb57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176115fb57604052565b67ffffffffffffffff81116115fb57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106119495750506000910152565b8181015183820152602001611939565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361199581518092818752878088019101611936565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019057565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361019057565b359073ffffffffffffffffffffffffffffffffffffffff8216820361019057565b6004359067ffffffffffffffff8216820361019057565b359067ffffffffffffffff8216820361019057565b81601f8201121561019057803590611a46826118fc565b92611a5460405194856118bb565b8284526020838301011161019057816000926020809301838601378301015290565b67ffffffffffffffff81116115fb5760051b60200190565b9181601f840112156101905782359167ffffffffffffffff8311610190576020808501948460051b01011161019057565b906020808351928381520192019060005b818110611add5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611ad0565b3590811515820361019057565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610190570180359067ffffffffffffffff82116101905760200191813603831361019057565b3567ffffffffffffffff811681036101905790565b5190811515820361019057565b81601f82011215610190578051611b9f816118fc565b92611bad60405194856118bb565b8184526020828401011161019057611bcb9160208085019101611936565b90565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff600154163303611c2e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015611c6c5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080601f83011215610190578135611cb281611a76565b92611cc060405194856118bb565b81845260208085019260051b82010192831161019057602001905b828210611ce85750505090565b60208091611cf5846119e2565b815201910190611cdb565b8054821015611c6c5760005260206000200190600090565b6000828152600182016020526040902054611da257805490680100000000000000008210156115fb5782611d8b611d56846001809601855584611d00565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014611f32577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611f03578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611f0357818103611ecc575b50505080548015611e9d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190611e5e8282611d00565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b611eec611edc611d569386611d00565b90549060031b1c92839286611d00565b905560005283602052604060002055388080611e26565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b91929015611fb65750815115611f4f575090565b3b15611f585790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015611fc95750805190602001fd5b612007906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611959565b0390fdfea164736f6c634300081a000a",
}

var CommitOnRampABI = CommitOnRampMetaData.ABI

var CommitOnRampBin = CommitOnRampMetaData.Bin

func DeployCommitOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig CommitOnRampDynamicConfig) (common.Address, *types.Transaction, *CommitOnRamp, error) {
	parsed, err := CommitOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitOnRampBin), backend, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitOnRamp{address: address, abi: *parsed, CommitOnRampCaller: CommitOnRampCaller{contract: contract}, CommitOnRampTransactor: CommitOnRampTransactor{contract: contract}, CommitOnRampFilterer: CommitOnRampFilterer{contract: contract}}, nil
}

type CommitOnRamp struct {
	address common.Address
	abi     abi.ABI
	CommitOnRampCaller
	CommitOnRampTransactor
	CommitOnRampFilterer
}

type CommitOnRampCaller struct {
	contract *bind.BoundContract
}

type CommitOnRampTransactor struct {
	contract *bind.BoundContract
}

type CommitOnRampFilterer struct {
	contract *bind.BoundContract
}

type CommitOnRampSession struct {
	Contract     *CommitOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitOnRampCallerSession struct {
	Contract *CommitOnRampCaller
	CallOpts bind.CallOpts
}

type CommitOnRampTransactorSession struct {
	Contract     *CommitOnRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitOnRampRaw struct {
	Contract *CommitOnRamp
}

type CommitOnRampCallerRaw struct {
	Contract *CommitOnRampCaller
}

type CommitOnRampTransactorRaw struct {
	Contract *CommitOnRampTransactor
}

func NewCommitOnRamp(address common.Address, backend bind.ContractBackend) (*CommitOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitOnRamp{address: address, abi: abi, CommitOnRampCaller: CommitOnRampCaller{contract: contract}, CommitOnRampTransactor: CommitOnRampTransactor{contract: contract}, CommitOnRampFilterer: CommitOnRampFilterer{contract: contract}}, nil
}

func NewCommitOnRampCaller(address common.Address, caller bind.ContractCaller) (*CommitOnRampCaller, error) {
	contract, err := bindCommitOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampCaller{contract: contract}, nil
}

func NewCommitOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitOnRampTransactor, error) {
	contract, err := bindCommitOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampTransactor{contract: contract}, nil
}

func NewCommitOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitOnRampFilterer, error) {
	contract, err := bindCommitOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampFilterer{contract: contract}, nil
}

func bindCommitOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitOnRamp *CommitOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOnRamp.Contract.CommitOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitOnRamp *CommitOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.CommitOnRampTransactor.contract.Transfer(opts)
}

func (_CommitOnRamp *CommitOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.CommitOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitOnRamp *CommitOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitOnRamp *CommitOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.contract.Transfer(opts)
}

func (_CommitOnRamp *CommitOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitOnRamp *CommitOnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitOnRamp *CommitOnRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitOnRamp.Contract.GetDestChainConfig(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitOnRamp.Contract.GetDestChainConfig(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitOnRampDynamicConfig, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitOnRampDynamicConfig)).(*CommitOnRampDynamicConfig)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetDynamicConfig() (CommitOnRampDynamicConfig, error) {
	return _CommitOnRamp.Contract.GetDynamicConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetDynamicConfig() (CommitOnRampDynamicConfig, error) {
	return _CommitOnRamp.Contract.GetDynamicConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee(&_CommitOnRamp.CallOpts, destChainSelector, message, arg2)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee(&_CommitOnRamp.CallOpts, destChainSelector, message, arg2)
}

func (_CommitOnRamp *CommitOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) Owner() (common.Address, error) {
	return _CommitOnRamp.Contract.Owner(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) Owner() (common.Address, error) {
	return _CommitOnRamp.Contract.Owner(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) TypeAndVersion() (string, error) {
	return _CommitOnRamp.Contract.TypeAndVersion(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitOnRamp.Contract.TypeAndVersion(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitOnRamp *CommitOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOnRamp.Contract.AcceptOwnership(&_CommitOnRamp.TransactOpts)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOnRamp.Contract.AcceptOwnership(&_CommitOnRamp.TransactOpts)
}

func (_CommitOnRamp *CommitOnRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyAllowlistUpdates(&_CommitOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyAllowlistUpdates(&_CommitOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, feeToken common.Address, feeTokenAmount *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "forwardToVerifier", message, messageId, feeToken, feeTokenAmount, verifierArgs)
}

func (_CommitOnRamp *CommitOnRampSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, feeToken common.Address, feeTokenAmount *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, message, messageId, feeToken, feeTokenAmount, verifierArgs)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, feeToken common.Address, feeTokenAmount *big.Int, verifierArgs []byte) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, message, messageId, feeToken, feeTokenAmount, verifierArgs)
}

func (_CommitOnRamp *CommitOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampSession) SetDynamicConfig(dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.SetDynamicConfig(&_CommitOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) SetDynamicConfig(dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.SetDynamicConfig(&_CommitOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitOnRamp *CommitOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.TransferOwnership(&_CommitOnRamp.TransactOpts, to)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.TransferOwnership(&_CommitOnRamp.TransactOpts, to)
}

func (_CommitOnRamp *CommitOnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitOnRamp *CommitOnRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.WithdrawFeeTokens(&_CommitOnRamp.TransactOpts, feeTokens)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.WithdrawFeeTokens(&_CommitOnRamp.TransactOpts, feeTokens)
}

type CommitOnRampAllowListSendersAddedIterator struct {
	Event *CommitOnRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampAllowListSendersAdded)
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
		it.Event = new(CommitOnRampAllowListSendersAdded)
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

func (it *CommitOnRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampAllowListSendersAddedIterator{contract: _CommitOnRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampAllowListSendersAdded)
				if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitOnRampAllowListSendersAdded, error) {
	event := new(CommitOnRampAllowListSendersAdded)
	if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampAllowListSendersRemovedIterator struct {
	Event *CommitOnRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampAllowListSendersRemoved)
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
		it.Event = new(CommitOnRampAllowListSendersRemoved)
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

func (it *CommitOnRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampAllowListSendersRemovedIterator{contract: _CommitOnRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampAllowListSendersRemoved)
				if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitOnRampAllowListSendersRemoved, error) {
	event := new(CommitOnRampAllowListSendersRemoved)
	if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampConfigSetIterator struct {
	Event *CommitOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampConfigSet)
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
		it.Event = new(CommitOnRampConfigSet)
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

func (it *CommitOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampConfigSet struct {
	DynamicConfig CommitOnRampDynamicConfig
	Raw           types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitOnRampConfigSetIterator, error) {

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitOnRampConfigSetIterator{contract: _CommitOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampConfigSet)
				if err := _CommitOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseConfigSet(log types.Log) (*CommitOnRampConfigSet, error) {
	event := new(CommitOnRampConfigSet)
	if err := _CommitOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampDestChainConfigSetIterator struct {
	Event *CommitOnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampDestChainConfigSet)
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
		it.Event = new(CommitOnRampDestChainConfigSet)
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

func (it *CommitOnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampDestChainConfigSetIterator{contract: _CommitOnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampDestChainConfigSet)
				if err := _CommitOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseDestChainConfigSet(log types.Log) (*CommitOnRampDestChainConfigSet, error) {
	event := new(CommitOnRampDestChainConfigSet)
	if err := _CommitOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampFeeTokenWithdrawnIterator struct {
	Event *CommitOnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampFeeTokenWithdrawn)
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
		it.Event = new(CommitOnRampFeeTokenWithdrawn)
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

func (it *CommitOnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitOnRampFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampFeeTokenWithdrawnIterator{contract: _CommitOnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitOnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampFeeTokenWithdrawn)
				if err := _CommitOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitOnRampFeeTokenWithdrawn, error) {
	event := new(CommitOnRampFeeTokenWithdrawn)
	if err := _CommitOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampOwnershipTransferRequestedIterator struct {
	Event *CommitOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampOwnershipTransferRequested)
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
		it.Event = new(CommitOnRampOwnershipTransferRequested)
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

func (it *CommitOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampOwnershipTransferRequestedIterator{contract: _CommitOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampOwnershipTransferRequested)
				if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitOnRampOwnershipTransferRequested, error) {
	event := new(CommitOnRampOwnershipTransferRequested)
	if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampOwnershipTransferredIterator struct {
	Event *CommitOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampOwnershipTransferred)
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
		it.Event = new(CommitOnRampOwnershipTransferred)
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

func (it *CommitOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampOwnershipTransferredIterator{contract: _CommitOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampOwnershipTransferred)
				if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitOnRampOwnershipTransferred, error) {
	event := new(CommitOnRampOwnershipTransferred)
	if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (CommitOnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitOnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe00542b2f9aa6cec740b3c4f8dcfbb444bac8a2cf03f7827f62bbdf74def306d")
}

func (CommitOnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitOnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CommitOnRamp *CommitOnRamp) Address() common.Address {
	return _CommitOnRamp.address
}

type CommitOnRampInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitOnRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, feeToken common.Address, feeTokenAmount *big.Int, verifierArgs []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitOnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitOnRampAllowListSendersRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitOnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitOnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitOnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitOnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitOnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitOnRampOwnershipTransferred, error)

	Address() common.Address
}
