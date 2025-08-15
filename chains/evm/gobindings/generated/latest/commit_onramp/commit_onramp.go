// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_onramp

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

type BaseOnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type BaseOnRampDestChainConfigArgs struct {
	VerifierAggregator common.Address
	DestChainSelector  uint64
	AllowlistEnabled   bool
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

type CommitOnRampStaticConfig struct {
	RmnRemote    common.Address
	NonceManager common.Address
}

var CommitOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structBaseOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByVerifierAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c0604052346103d95761297a80380380610019816103fd565b9283398101908082039060c082126103d95761003481610422565b606061004260208401610422565b93603f1901126103d9576100546103de565b9161006160408201610422565b835261006f60608201610422565b916020840192835261008360808301610422565b6040850190815260a083015190926001600160401b0382116103d9570186601f820112156103d9578051906001600160401b038211610321576100cb60208360051b016103fd565b9760206060818b868152019402830101918183116103d957602001925b82841061037c5750505050331561036b57600180546001600160a01b031916331790556001600160a01b03166080819052918215801561035a575b6103375760a085905283516001600160a01b0316158015610348575b610337578351600380546001600160a01b039283166001600160a01b03199182161790915582516004805491841691831691909117905583516005805491909316911617905560408051908101959094908587106001600160401b038811176103215760409687528486526001600160a01b0390811660209687019081528751958652518116958501959095525184169483019490945292518216606082015291511660808201527fb3ed693dbc8a60220245fab67b6fa0323da19566c2a33c5f0e0c04b89ec880729060a090a16000905b80518210156102e4576102248282610436565b51916001600160401b03602061023a8385610436565b510151169283156102cf5760008481526002602090815260409182902083518154948401516001600160a81b0319909516600882901b610100600160a81b03161794151560ff1694851790915582516001600160a01b03909116815292151590830152929360019390917f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9190a20190610211565b8363c35aa79d60e01b60005260045260246000fd5b6040516125199081610461823960805181818161130001528181611beb015261204a015260a051818181611339015281816119b10152611c240152f35b634e487b7160e01b600052604160045260246000fd5b6306b7c75960e31b60005260046000fd5b5080516001600160a01b03161561013f565b506001600160a01b03851615610123565b639b15e16f60e01b60005260046000fd5b6060848303126103d95761038e6103de565b9061039885610422565b825260208501516001600160401b03811681036103d957602083015260408501519081151582036103d95782602092604060609501528152019301926100e8565b600080fd5b60405190606082016001600160401b0381118382101761032157604052565b6040519190601f01601f191682016001600160401b0381118382101761032157604052565b51906001600160a01b03821682036103d957565b805182101561044a5760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816306285c6914611b85575080631234eab01461148d578063181f5a771461141057806334d560e4146111585780634a7597b514610d985780635cb80c5d14610aa55780636def4ce7146109c65780637437ff9f146108db57806379ba5097146107f25780638da5cb5b146107a0578063a8dd2df214610571578063c9b146b3146101a35763f2fde38b146100ae57600080fd5b3461019e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5773ffffffffffffffffffffffffffffffffffffffff6100fa611dc8565b610102612151565b1633811461017457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461019e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760043567ffffffffffffffff811161019e576101f2903690600401611e97565b73ffffffffffffffffffffffffffffffffffffffff600154163303610527575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610525576000938060051b8401358281121561052157840191608083360312610521576040519461027086611cf9565b61027984611e23565b865261028760208501611f12565b9660208701978852604085013567ffffffffffffffff811161051d576102b0903690870161219c565b9460408801958652606081013567ffffffffffffffff8111610519576102d89136910161219c565b946060880195865267ffffffffffffffff885116825260026020526040822098511515610330818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b8151516103ed575b50959760010195505b84518051821015610380579061037973ffffffffffffffffffffffffffffffffffffffff61037183600195611f9e565b5116886122aa565b5001610341565b505095949093506001925190815161039e575b50500192919061023c565b6103e367ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190611ec8565b0390a28580610393565b98939592909497989691966000146104e257600184019591875b86518051821015610484576104318273ffffffffffffffffffffffffffffffffffffffff92611f9e565b5116801561044d57906104466001928a612219565b5001610407565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816104d867ffffffffffffffff8b51169251604051918291602083526020830190611ec8565b0390a29089610338565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60055416330315610212577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461019e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760043567ffffffffffffffff811161019e573660238201121561019e5780600401356105cb81611e7f565b916105d96040519384611d4d565b818352602460606020850193028201019036821161019e57602401915b81831061074f5783610606612151565b6000905b80518210156105255761061d8282611f9e565b519167ffffffffffffffff60206106348385611f9e565b5101511692831561072157927f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c60406001949583600052600260205260ff826000206106f58484519483547fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff008860081b16911617845501511515829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a2019061060a565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60608336031261019e57602060609160405161076a81611d15565b61077386611deb565b8152610780838701611e23565b8382015261079060408701611f12565b60408201528152019201916105f6565b3461019e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461019e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760005473ffffffffffffffffffffffffffffffffffffffff811633036108b1577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461019e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760006040805161091981611d15565b8281528260208201520152606060405161093281611d15565b73ffffffffffffffffffffffffffffffffffffffff60035416815273ffffffffffffffffffffffffffffffffffffffff60045416602082015273ffffffffffffffffffffffffffffffffffffffff6005541660408201526109c4604051809273ffffffffffffffffffffffffffffffffffffffff60408092828151168552826020820151166020860152015116910152565bf35b3461019e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5767ffffffffffffffff610a06611e0c565b16600052600260205260406000206001815491019060405191826020825491828152019160005260206000209060005b818110610a8f5773ffffffffffffffffffffffffffffffffffffffff85610a8b88610a6381890382611d4d565b604051938360ff8695161515855260081c166020840152606060408401526060830190611ec8565b0390f35b8254845260209093019260019283019201610a36565b3461019e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760043567ffffffffffffffff811161019e57610af4903690600401611e97565b73ffffffffffffffffffffffffffffffffffffffff6004541660005b82811015610525576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff8116809103610519576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115610d8d579085918591610d55575b5080610b9a575b5050506001915001610b10565b60405194610c5460208701967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602482015283604482015260448152610be4606482611d4d565b82806040998a5193610bf68c86611d4d565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d15610d4d573d90610c3782611d8e565b91610c448b519384611d4d565b82523d85602084013e5b8761243c565b90815180610c94575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381610b8d565b82939596979450916020919281010312610d4a57506020610cb59101611fe1565b15610cc7579291908490888080610c5d565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b606090610c4e565b9150506020813d8211610d85575b81610d7060209383611d4d565b81010312610d815784905188610b86565b8380fd5b3d9150610d63565b6040513d86823e3d90fd5b3461019e5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e57610dcf611e0c565b60243567ffffffffffffffff811161019e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261019e5760405190610e1a82611d31565b806004013567ffffffffffffffff811161019e57610e3e9060043691840101611e38565b8252602481013567ffffffffffffffff811161019e57610e649060043691840101611e38565b60208301908152604482013567ffffffffffffffff811161019e578201913660238401121561019e576004830135610e9b81611e7f565b93610ea96040519586611d4d565b818552602060048187019360061b830101019036821161019e57602401915b8183106111205750505060408401928352610ee560648201611deb565b906060850191825260848101359067ffffffffffffffff821161019e576004610f119236920101611e38565b91608085019283526044359067ffffffffffffffff821161019e57610feb90610f4b67ffffffffffffffff98969594933690600401611e38565b50610f5586612033565b610fba73ffffffffffffffffffffffffffffffffffffffff6003541697604051998a987fd8694ccd000000000000000000000000000000000000000000000000000000008a52166004890152604060248901525160a0604489015260e4880190611c9a565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc878303016064880152611c9a565b9251927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8582030160848601526020808551928381520194019060005b8181106110e55750505061108960209593859373ffffffffffffffffffffffffffffffffffffffff8594511660a4850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8483030160c4850152611c9a565b03915afa80156110d9576000906110a6575b602090604051908152f35b506020813d6020116110d1575b816110c060209383611d4d565b8101031261019e576020905161109b565b3d91506110b3565b6040513d6000823e3d90fd5b8251805173ffffffffffffffffffffffffffffffffffffffff1687526020908101518188015289975060409096019590920191600101611028565b60408336031261019e576020604091825161113a81611cdd565b61114386611deb565b81528286013583820152815201920191610ec8565b3461019e5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e57600060405161119581611d15565b61119d611dc8565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361051957602082019081526044359073ffffffffffffffffffffffffffffffffffffffff82168203610d8157604083019182526111f6612151565b73ffffffffffffffffffffffffffffffffffffffff8351161580156113f1575b6113c9579173ffffffffffffffffffffffffffffffffffffffff60a092817fb3ed693dbc8a60220245fab67b6fa0323da19566c2a33c5f0e0c04b89ec8807295818551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff000000000000000000000000000000000000000060055416176005556113c5604051916112e983611cdd565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602084015261138d604051809473ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b604083019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552826020820151166020860152015116910152565ba180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff81511615611216565b3461019e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e57610a8b60408051906114518183611d4d565b601682527f436f6d6d69744f6e52616d7020312e372e302d64657600000000000000000000602083015251918291602083526020830190611c9a565b3461019e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e5760043567ffffffffffffffff811161019e573660238201121561019e57806004013567ffffffffffffffff811161019e57810190602482019036821161019e5760208184031261019e5760248101359067ffffffffffffffff821161019e570190818303906101c0821261019e576040519161016083019083821067ffffffffffffffff831117611b56576080916040521261019e5760405161155e81611cf9565b6024840135815261157160448501611e23565b602082015261158260648501611e23565b604082015261159360848501611e23565b606082015282526115a660a48401611deb565b926020830193845260c481013567ffffffffffffffff811161019e578260246115d192840101611e38565b604084015260e481013567ffffffffffffffff811161019e578260246115f992840101611e38565b946060840195865261160e6101048301611deb565b916080850192835260a0850191610124820135835261014482013560c087015261016482013567ffffffffffffffff811161019e578201602481019160809190031261019e576040519061166182611cf9565b61166a81611deb565b8252602081013567ffffffffffffffff811161019e578661168c918301611e38565b602083015260408101359067ffffffffffffffff821161019e576116b4876060938301611e38565b60408401520135606082015260e086015261018481013567ffffffffffffffff811161019e576024908201019380601f8601121561019e5784356116f781611e7f565b956117056040519788611d4d565b81875260208088019260051b8201019183831161019e5760208201905b838210611b28575050505061010086019485526101a482013567ffffffffffffffff811161019e5781602461175992850101611f1f565b6101208701526101c48201359167ffffffffffffffff831161019e576117829201602401611f1f565b6101408501526117a067ffffffffffffffff60408651015116612033565b67ffffffffffffffff6040855101511673ffffffffffffffffffffffffffffffffffffffff8651169060005260026020526040600020805473ffffffffffffffffffffffffffffffffffffffff8160081c163303611aff5760ff16611ab5575b5050906000929173ffffffffffffffffffffffffffffffffffffffff60035416916118e3608061185e73ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60408b510151169451169451966024359051611f9e565b51015198516118b36040519a8b97889687967f3a49bb4900000000000000000000000000000000000000000000000000000000885260048801526024870152604486015260a0606486015260a4850190611c9a565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016084850152611c9a565b03915afa9283156110d957600093611a40575b506000921561193a575b610a8b8367ffffffffffffffff6040519116602082015260208152611926604082611d4d565b604051918291602083526020830190611c9a565b67ffffffffffffffff604073ffffffffffffffffffffffffffffffffffffffff9251015116915116604051917fea458c0c000000000000000000000000000000000000000000000000000000008352600483015260248201526020816044818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611a355782916119ed575b509050610a8b82611900565b90506020813d602011611a2d575b81611a0860209383611d4d565b8101031261051d575167ffffffffffffffff8116810361051d57610a8b9150826119e1565b3d91506119fb565b6040513d84823e3d90fd5b90923d8082843e611a518184611d4d565b820190608083830312610d4a57611a6a60208401611fe1565b92604081015167ffffffffffffffff81116105195783611a8b918301611fee565b5060608101519167ffffffffffffffff8311610d4a5750611aad929101611fee565b5091836118f6565b60008281526002909101602052604090205415611ad25780611800565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7ed135dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b813567ffffffffffffffff811161019e57602091611b4b87848094880101611f1f565b815201910190611722565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461019e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019e57602081611bc2600093611cdd565b828152015260408051611bd481611cdd565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201526109c48251809273ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b60005b838110611c8a5750506000910152565b8181015183820152602001611c7a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611cd681518092818752878088019101611c77565b0116010190565b6040810190811067ffffffffffffffff821117611b5657604052565b6080810190811067ffffffffffffffff821117611b5657604052565b6060810190811067ffffffffffffffff821117611b5657604052565b60a0810190811067ffffffffffffffff821117611b5657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611b5657604052565b67ffffffffffffffff8111611b5657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019e57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361019e57565b6004359067ffffffffffffffff8216820361019e57565b359067ffffffffffffffff8216820361019e57565b81601f8201121561019e57803590611e4f82611d8e565b92611e5d6040519485611d4d565b8284526020838301011161019e57816000926020809301838601378301015290565b67ffffffffffffffff8111611b565760051b60200190565b9181601f8401121561019e5782359167ffffffffffffffff831161019e576020808501948460051b01011161019e57565b906020808351928381520192019060005b818110611ee65750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611ed9565b3590811515820361019e57565b919060a08382031261019e5760405190611f3882611d31565b8193611f4381611deb565b8352611f5160208201611e23565b6020840152604081013563ffffffff8116810361019e5760408401526060810135606084015260808101359167ffffffffffffffff831161019e57608092611f999201611e38565b910152565b8051821015611fb25760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b5190811515820361019e57565b81601f8201121561019e57805161200481611d8e565b926120126040519485611d4d565b8184526020828401011161019e576120309160208085019101611c77565b90565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561214d576020602491604051928380927f2cbc26bb00000000000000000000000000000000000000000000000000000000825277ffffffffffffffff000000000000000000000000000000008760801b1660048301525afa9081156110d957600091612113575b506120db5750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b90506020813d602011612145575b8161212e60209383611d4d565b8101031261019e5761213f90611fe1565b386120d3565b3d9150612121565b5050565b73ffffffffffffffffffffffffffffffffffffffff60015416330361217257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080601f8301121561019e5781356121b381611e7f565b926121c16040519485611d4d565b81845260208085019260051b82010192831161019e57602001905b8282106121e95750505090565b602080916121f684611deb565b8152019101906121dc565b8054821015611fb25760005260206000200190600090565b60008281526001820160205260409020546122a35780549068010000000000000000821015611b56578261228c612257846001809601855584612201565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201918160005282602052604060002054801515600014612433577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111612404578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211612404578181036123cd575b5050508054801561239e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061235f8282612201565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6123ed6123dd6122579386612201565b90549060031b1c92839286612201565b905560005283602052604060002055388080612327565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b919290156124b75750815115612450575090565b3b156124595790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156124ca5750805190602001fd5b612508906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611c9a565b0390fdfea164736f6c634300081a000a",
}

var CommitOnRampABI = CommitOnRampMetaData.ABI

var CommitOnRampBin = CommitOnRampMetaData.Bin

func DeployCommitOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, rmnRemote common.Address, nonceManager common.Address, dynamicConfig CommitOnRampDynamicConfig, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (common.Address, *types.Transaction, *CommitOnRamp, error) {
	parsed, err := CommitOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitOnRampBin), backend, rmnRemote, nonceManager, dynamicConfig, destChainConfigArgs)
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
	outstruct.VerifierAggregator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
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

func (_CommitOnRamp *CommitOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (CommitOnRampStaticConfig, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CommitOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitOnRampStaticConfig)).(*CommitOnRampStaticConfig)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetStaticConfig() (CommitOnRampStaticConfig, error) {
	return _CommitOnRamp.Contract.GetStaticConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetStaticConfig() (CommitOnRampStaticConfig, error) {
	return _CommitOnRamp.Contract.GetStaticConfig(&_CommitOnRamp.CallOpts)
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

func (_CommitOnRamp *CommitOnRampTransactor) ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "forwardToVerifier", rawMessage, verifierIndex)
}

func (_CommitOnRamp *CommitOnRampSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, rawMessage, verifierIndex)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, rawMessage, verifierIndex)
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
	StaticConfig  CommitOnRampStaticConfig
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
	VerifierAggregator common.Address
	AllowedSendersList []common.Address
}

func (_CommitOnRamp *CommitOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CommitOnRamp.abi.Events["AllowListSendersAdded"].ID:
		return _CommitOnRamp.ParseAllowListSendersAdded(log)
	case _CommitOnRamp.abi.Events["AllowListSendersRemoved"].ID:
		return _CommitOnRamp.ParseAllowListSendersRemoved(log)
	case _CommitOnRamp.abi.Events["ConfigSet"].ID:
		return _CommitOnRamp.ParseConfigSet(log)
	case _CommitOnRamp.abi.Events["DestChainConfigSet"].ID:
		return _CommitOnRamp.ParseDestChainConfigSet(log)
	case _CommitOnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _CommitOnRamp.ParseFeeTokenWithdrawn(log)
	case _CommitOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _CommitOnRamp.ParseOwnershipTransferRequested(log)
	case _CommitOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _CommitOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CommitOnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitOnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb3ed693dbc8a60220245fab67b6fa0323da19566c2a33c5f0e0c04b89ec88072")
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

	GetStaticConfig(opts *bind.CallOpts) (CommitOnRampStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseOnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
