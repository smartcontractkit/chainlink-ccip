// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package advanced_pool_hooks

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

type AdvancedPoolHooksCCVConfigArg struct {
	RemoteChainSelector   uint64
	OutboundCCVs          []common.Address
	ThresholdOutboundCCVs []common.Address
	InboundCCVs           []common.Address
	ThresholdInboundCCVs  []common.Address
}

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender          []byte
	RemoteChainSelector     uint64
	Receiver                common.Address
	SourceDenominatedAmount *big.Int
	LocalToken              common.Address
	SourcePoolAddress       []byte
	SourcePoolData          []byte
	OffchainTokenData       []byte
}

var AdvancedPoolHooksMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct AdvancedPoolHooks.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkAllowList\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThresholdAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postFlightCheck\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"localAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preflightCheck\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngine\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThresholdAmount\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholdOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholdInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineSet\",\"inputs\":[{\"name\":\"oldPolicyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdAmountSet\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyUnderThresholdCCVsForThresholdCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a0806040523461024957612698803803809161001c8285610264565b833981016060828203126102495781516001600160401b0381116102495782019080601f83011215610249578151916001600160401b03831161024e578260051b9060208201936100706040519586610264565b845260208085019282010192831161024957602001905b828210610231575050506100a2604060208401519301610287565b90331561022057600180546001600160a01b031916331790558051151560808190526100fc575b506100d6916004556102c5565b604051612169908161052f823960805181818161022801528181610d38015261194f0152f35b906020906040519061010e8383610264565b6000825260003681376080511561020f5760005b8251811015610189576001906001600160a01b03610140828661029b565b51168561014c826103d0565b610159575b505001610122565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13885610151565b5092905060005b8151811015610204576001906001600160a01b036101ae828561029b565b511680156101fe57846101c0826104ce565b6101ce575b50505b01610190565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138846101c5565b506101c8565b5050506100d66100c9565b6335f4a7b360e01b60005260046000fd5b639b15e16f60e01b60005260046000fd5b6020809161023e84610287565b815201910190610087565b600080fd5b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761024e57604052565b51906001600160a01b038216820361024957565b80518210156102af5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6005546001600160a01b03918216911660008282146103b35781610373575b600580546001600160a01b0319168417905582610322575b807ffb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e91a3565b823b1561036557604051631100482d60e01b8152818160048183885af1801561036857610350575b506102fc565b61035b828092610264565b610365573861034a565b80fd5b6040513d84823e3d90fd5b813b1561036557604051628950d760e61b8152818160048183875af180156103685782906103a3575b50506102e4565b6103ac91610264565b388161039c565b505050565b80548210156102af5760005260206000200190600090565b60008181526003602052604090205480156104c75760001981018181116104b1576002546000198101919082116104b157818103610460575b505050600254801561044a57600019016104248160026103b8565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6104996104716104829360026103b8565b90549060031b1c92839260026103b8565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610409565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610528576002546801000000000000000081101561024e5761050f61048282600185940160025560026103b8565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806354c8a4f314610ca65780635c3af7ca14610bfc5780635eff3bf714610b705780636135b08514610b295780636831731214610ad757806379ba5097146109ee57806389720a62146109595780638da5cb5b14610907578063961e2e7c1461089d578063a7cd63b7146107f8578063ce07c7c8146107b7578063d966866b1461024d578063e0351e13146101f2578063f2fde38b146101025763f72c071b146100c157600080fd5b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd576020600454604051908152f35b600080fd5b346100fd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5773ffffffffffffffffffffffffffffffffffffffff61014e610ef0565b610156611ae1565b163381146101c857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346100fd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043567ffffffffffffffff81116100fd5761029c903690600401610e80565b6102a4611ae1565b6000905b8082106102b157005b6102c46102bf8383866119cf565b6110bf565b906102dd6102d38483876119cf565b6020810190611a3e565b94906102f76102ed8685856119cf565b6040810190611a3e565b96906103116103078887876119cf565b6060810190611a3e565b9061032a6103208a89896119cf565b6080810190611a3e565b92909361034061033b36888a610ff9565b611deb565b61034e61033b368486610ff9565b8b610783575b83610721575b6040519b6103678d610f84565b61037236888a610ff9565b8d528c610380368385610ff9565b9b602082019c8d5267ffffffffffffffff61039c368789610ff9565b91604084019283526103af368a8c610ff9565b6060850152169c8d60005260066020526040600020925180519067ffffffffffffffff82116105f1576801000000000000000082116105f15784548286558083106106f8575b5060200184600052602060002060005b8381106106ce57505050506001839e9c9d9e0190519081519167ffffffffffffffff83116105f1576801000000000000000083116105f157815483835580841061069d575b5060200190600052602060002060005b83811061067357505050506002820190519081519167ffffffffffffffff83116105f1576801000000000000000083116105f157815483835580841061064a575b509e9f939495969798999a9b9c9d9e60200190600052602060002060005b838110610620575050505060036060919e9c9d9e019101519081519167ffffffffffffffff83116105f1576801000000000000000083116105f15781548383558084106105c8575b5060200190600052602060002060005b83811061059e57505050506105939360019a999896936105777fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d999794610585946105696040519b8c9b60808d5260808d0191611a92565b918a830360208c0152611a92565b918783036040890152611a92565b918483036060860152611a92565b0390a20190916102a8565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610511565b8260005283602060002091820191015b8181106105e55750610501565b600081556001016105d8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016104b9565b8260005283602060002091820191015b818110610667575061049b565b6000815560010161065a565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161045a565b9d9f9e9d8260005283602060002091820191015b8181106106c257509f9d9e9f61044a565b600081556001016106b1565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610405565b8560005282602060002091820191015b81811061071557506103f5565b60008155600101610708565b81156107595761073561033b368688610ff9565b610754610743368486610ff9565b61074e368789610ff9565b90611eb5565b61035a565b7f1d56c21d0000000000000000000000000000000000000000000000000000000060005260046000fd5b85156107595761079761033b368e84610ff9565b6107b28c61074e6107a9368a8c610ff9565b91369085610ff9565b610354565b346100fd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd576107f66107f1610ef0565b61194d565b005b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610887576108838561087781870382610fa0565b60405191829182610f34565b0390f35b8254845260209093019260019283019201610860565b346100fd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd577f80dc2a1a49dda9f8bd85c1c376266e011db6448050b7bfd5c2f423e162c1114560206004356108fa611ae1565b80600455604051908152a1005b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100fd5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57610990610ef0565b5060243567ffffffffffffffff811681036100fd576109ad610eb1565b5060843567ffffffffffffffff81116100fd576109ce903690600401610ec2565b505060a4359060028210156100fd576108839161087791604435906118b5565b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760005473ffffffffffffffffffffffffffffffffffffffff81163303610aad577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100fd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602073ffffffffffffffffffffffffffffffffffffffff60055416604051908152f35b346100fd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd576107f6610b63610ef0565b610b6b611ae1565b611b40565b346100fd5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043567ffffffffffffffff81116100fd576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100fd5760443561ffff811681036100fd576107f6916024359060040161150d565b346100fd5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043567ffffffffffffffff81116100fd5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100fd5760243561ffff811681036100fd576044359167ffffffffffffffff83116100fd57610c9b6107f6933690600401610ec2565b92909160040161123a565b346100fd5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043567ffffffffffffffff81116100fd57610cf5903690600401610e80565b6024359067ffffffffffffffff82116100fd57610d2e610d1c610d36933690600401610e80565b949092610d27611ae1565b3691610ff9565b923691610ff9565b7f000000000000000000000000000000000000000000000000000000000000000015610e565760005b8251811015610dd2578073ffffffffffffffffffffffffffffffffffffffff610d8a60019386611b2c565b5116610d9581611f66565b610da1575b5001610d5f565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184610d9a565b5060005b81518110156107f6578073ffffffffffffffffffffffffffffffffffffffff610e0160019385611b2c565b51168015610e5057610e12816120fc565b610e1f575b505b01610dd6565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183610e17565b50610e19565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156100fd5782359167ffffffffffffffff83116100fd576020808501948460051b0101116100fd57565b6064359061ffff821682036100fd57565b9181601f840112156100fd5782359167ffffffffffffffff83116100fd57602083818601950101116100fd57565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100fd57565b359073ffffffffffffffffffffffffffffffffffffffff821682036100fd57565b602060408183019282815284518094520192019060005b818110610f585750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610f4b565b6080810190811067ffffffffffffffff8211176105f157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176105f157604052565b67ffffffffffffffff81116105f15760051b60200190565b92919061100581610fe1565b936110136040519586610fa0565b602085838152019160051b81019283116100fd57905b82821061103557505050565b6020809161104284610f13565b815201910190611029565b3573ffffffffffffffffffffffffffffffffffffffff811681036100fd5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100fd570180359067ffffffffffffffff82116100fd576020019181360383136100fd57565b3567ffffffffffffffff811681036100fd5790565b92919267ffffffffffffffff82116105f1576040519161111c601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184610fa0565b8294818452818301116100fd578281602093846000960137010152565b919082519283825260005b8481106111835750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611144565b9061123791602081527fffffffff00000000000000000000000000000000000000000000000000000000825116602082015273ffffffffffffffffffffffffffffffffffffffff60208301511660408201526060611204604084015160808385015260a0840190611139565b9201519060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152611139565b90565b9190604083019061124d6107f18361104d565b6005549473ffffffffffffffffffffffffffffffffffffffff600096169485156115045761127b818061106e565b95909261129361128d602085016110bf565b9661104d565b956112a06080850161104d565b946040519860e08a018a811067ffffffffffffffff8211176114d7579573ffffffffffffffffffffffffffffffffffffffff61133c6114159b9761ffff97888f60a0906101249f9a6113e99f9967ffffffffffffffff9f9b611308918a9d60405236916110d4565b8252602082019e8f9116905286604082019d168d5260608082019801358852866080820199168952019816885236916110d4565b9660c08d0197885267ffffffffffffffff6113956040519e8f9d8e917f73bb902c00000000000000000000000000000000000000000000000000000000602084015260206024840152519160e060448201520190611139565b99511660648c0152511660848a01525160a4890152511660c4870152511660e4850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc84830301610104850152611139565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610fa0565b6040519061142282610f84565b7f5c3af7ca000000000000000000000000000000000000000000000000000000008252336020830152604082015260405161145e602082610fa0565b8381526060820152813b156114d3576114a9839283926040519485809481937fc2098e0800000000000000000000000000000000000000000000000000000000835260048301611198565b03925af180156114c8576114bb575050565b816114c591610fa0565b50565b6040513d84823e3d90fd5b8280fd5b60248d7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50505050505050565b6005549273ffffffffffffffffffffffffffffffffffffffff6000941692831561185e5761153b838061106e565b92909361154a602082016110bf565b946115576040830161104d565b956115646080840161104d565b9161157260a085018561106e565b93909161158260c087018761106e565b9361159060e089018961106e565b9790936040519c8d610140810190811067ffffffffffffffff82111761182e5760405236906115be926110d4565b8c5260208c019667ffffffffffffffff16875260408c019c73ffffffffffffffffffffffffffffffffffffffff168d5260608c019860600135895260808c019473ffffffffffffffffffffffffffffffffffffffff1685523690611621926110d4565b9360a08b019485523690611634926110d4565b9460c08a019586523690611647926110d4565b9460e08901958652610100890197885261012089019661ffff168752604051998a9960208b017fe8deab7900000000000000000000000000000000000000000000000000000000905260248b01602090525160448b0161014090526101848b016116b091611139565b945167ffffffffffffffff1660648b01525173ffffffffffffffffffffffffffffffffffffffff1660848a01525160a48901525173ffffffffffffffffffffffffffffffffffffffff1660c488015251908681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160e488015261173491611139565b9051908581037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0161010487015261176b91611139565b9051908481037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc016101248601526117a291611139565b91516101448401525161ffff16610164830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526117e59082610fa0565b604051906117f282610f84565b7f5eff3bf7000000000000000000000000000000000000000000000000000000008252336020830152604082015260405161145e602082610fa0565b50505060248f7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5050505050565b906020825491828152019160005260206000209060005b8181106118895750505090565b825473ffffffffffffffffffffffffffffffffffffffff1684526020909301926001928301920161187c565b67ffffffffffffffff166000526006602052604060002091600281101561191e5760011461190357611237916001604051916118fc836118f58184611865565b0384610fa0565b0190611cbc565b611237916003604051916118fc836118f58160028501611865565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f00000000000000000000000000000000000000000000000000000000000000006119755750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156119a25750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190811015611a0f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100fd570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100fd570180359067ffffffffffffffff82116100fd57602001918160051b360383136100fd57565b9160209082815201919060005b818110611aac5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff611ad388610f13565b168152019401929101611a9f565b73ffffffffffffffffffffffffffffffffffffffff600154163303611b0257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015611a0f5760209160051b010190565b60055473ffffffffffffffffffffffffffffffffffffffff91821691166000828214611c7b5781611c21575b827fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055582611bc2575b807ffb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e91a3565b823b15611c1e576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183885af180156114c857611c09575b50611b9c565b611c14828092610fa0565b611c1e5738611c03565b80fd5b813b15611c1e576040517f225435c0000000000000000000000000000000000000000000000000000000008152818160048183875af180156114c8578290611c6b575b5050611b6c565b611c7491610fa0565b3881611c64565b505050565b91908201809211611c8d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b916004548015159182611de0575b5050611cd4575090565b90611cef611ce89260405193848092611865565b0383610fa0565b815180611cfc5750905090565b611d07908251611c80565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d4b611d3586610fe1565b95611d436040519788610fa0565b808752610fe1565b0136602086013760005b8251811015611d93578073ffffffffffffffffffffffffffffffffffffffff611d8060019386611b2c565b5116611d8c8288611b2c565b5201611d55565b509160005b8151811015611c7b578073ffffffffffffffffffffffffffffffffffffffff611dc360019385611b2c565b5116611dd9611dd3838751611c80565b88611b2c565b5201611d98565b101590503880611cca565b805160005b818110611dfc57505050565b60018101808211611c8d575b828110611e185750600101611df0565b73ffffffffffffffffffffffffffffffffffffffff611e378386611b2c565b511673ffffffffffffffffffffffffffffffffffffffff611e588387611b2c565b511614611e6757600101611e08565b73ffffffffffffffffffffffffffffffffffffffff611e868386611b2c565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9081519080519060005b838110611ecd575050505050565b60005b838110611ee05750600101611ebf565b73ffffffffffffffffffffffffffffffffffffffff611eff8388611b2c565b511673ffffffffffffffffffffffffffffffffffffffff611f208386611b2c565b511614611f2f57600101611ed0565b73ffffffffffffffffffffffffffffffffffffffff611e868388611b2c565b8054821015611a0f5760005260206000200190600090565b60008181526003602052604090205480156120f5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611c8d57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611c8d57818103612086575b5050506002548015612057577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612014816002611f4e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6120dd6120976120a8936002611f4e565b90549060031b1c9283926002611f4e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080611fdb565b5050600090565b8060005260036020526040600020541560001461215657600254680100000000000000008110156105f15761213d6120a88260018594016002556002611f4e565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var AdvancedPoolHooksABI = AdvancedPoolHooksMetaData.ABI

var AdvancedPoolHooksBin = AdvancedPoolHooksMetaData.Bin

func DeployAdvancedPoolHooks(auth *bind.TransactOpts, backend bind.ContractBackend, allowlist []common.Address, thresholdAmountForAdditionalCCVs *big.Int, policyEngine common.Address) (common.Address, *types.Transaction, *AdvancedPoolHooks, error) {
	parsed, err := AdvancedPoolHooksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdvancedPoolHooksBin), backend, allowlist, thresholdAmountForAdditionalCCVs, policyEngine)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdvancedPoolHooks{address: address, abi: *parsed, AdvancedPoolHooksCaller: AdvancedPoolHooksCaller{contract: contract}, AdvancedPoolHooksTransactor: AdvancedPoolHooksTransactor{contract: contract}, AdvancedPoolHooksFilterer: AdvancedPoolHooksFilterer{contract: contract}}, nil
}

type AdvancedPoolHooks struct {
	address common.Address
	abi     abi.ABI
	AdvancedPoolHooksCaller
	AdvancedPoolHooksTransactor
	AdvancedPoolHooksFilterer
}

type AdvancedPoolHooksCaller struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksTransactor struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksFilterer struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksSession struct {
	Contract     *AdvancedPoolHooks
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksCallerSession struct {
	Contract *AdvancedPoolHooksCaller
	CallOpts bind.CallOpts
}

type AdvancedPoolHooksTransactorSession struct {
	Contract     *AdvancedPoolHooksTransactor
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksRaw struct {
	Contract *AdvancedPoolHooks
}

type AdvancedPoolHooksCallerRaw struct {
	Contract *AdvancedPoolHooksCaller
}

type AdvancedPoolHooksTransactorRaw struct {
	Contract *AdvancedPoolHooksTransactor
}

func NewAdvancedPoolHooks(address common.Address, backend bind.ContractBackend) (*AdvancedPoolHooks, error) {
	abi, err := abi.JSON(strings.NewReader(AdvancedPoolHooksABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAdvancedPoolHooks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooks{address: address, abi: abi, AdvancedPoolHooksCaller: AdvancedPoolHooksCaller{contract: contract}, AdvancedPoolHooksTransactor: AdvancedPoolHooksTransactor{contract: contract}, AdvancedPoolHooksFilterer: AdvancedPoolHooksFilterer{contract: contract}}, nil
}

func NewAdvancedPoolHooksCaller(address common.Address, caller bind.ContractCaller) (*AdvancedPoolHooksCaller, error) {
	contract, err := bindAdvancedPoolHooks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksCaller{contract: contract}, nil
}

func NewAdvancedPoolHooksTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedPoolHooksTransactor, error) {
	contract, err := bindAdvancedPoolHooks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksTransactor{contract: contract}, nil
}

func NewAdvancedPoolHooksFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedPoolHooksFilterer, error) {
	contract, err := bindAdvancedPoolHooks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksFilterer{contract: contract}, nil
}

func bindAdvancedPoolHooks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdvancedPoolHooksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksCaller.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksTransactor.contract.Transfer(opts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksTransactor.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooks.Contract.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.contract.Transfer(opts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) CheckAllowList(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "checkAllowList", sender)

	if err != nil {
		return err
	}

	return err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) CheckAllowList(sender common.Address) error {
	return _AdvancedPoolHooks.Contract.CheckAllowList(&_AdvancedPoolHooks.CallOpts, sender)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) CheckAllowList(sender common.Address) error {
	return _AdvancedPoolHooks.Contract.CheckAllowList(&_AdvancedPoolHooks.CallOpts, sender)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAllowList() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllowList(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAllowList() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllowList(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAllowListEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAllowListEnabled(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAllowListEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAllowListEnabled(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetPolicyEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getPolicyEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetPolicyEngine() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetPolicyEngine(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetPolicyEngine() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetPolicyEngine(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetRequiredCCVs(&_AdvancedPoolHooks.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetRequiredCCVs(&_AdvancedPoolHooks.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetThresholdAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getThresholdAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetThresholdAmount() (*big.Int, error) {
	return _AdvancedPoolHooks.Contract.GetThresholdAmount(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetThresholdAmount() (*big.Int, error) {
	return _AdvancedPoolHooks.Contract.GetThresholdAmount(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) Owner() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.Owner(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) Owner() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.Owner(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "acceptOwnership")
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) AcceptOwnership() (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AcceptOwnership(&_AdvancedPoolHooks.TransactOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AcceptOwnership(&_AdvancedPoolHooks.TransactOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAllowListUpdates(&_AdvancedPoolHooks.TransactOpts, removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAllowListUpdates(&_AdvancedPoolHooks.TransactOpts, removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ApplyCCVConfigUpdates(ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyCCVConfigUpdates(&_AdvancedPoolHooks.TransactOpts, ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyCCVConfigUpdates(&_AdvancedPoolHooks.TransactOpts, ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) PostFlightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "postFlightCheck", releaseOrMintIn, localAmount, blockConfirmationRequested)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PostFlightCheck(releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.TransactOpts, releaseOrMintIn, localAmount, blockConfirmationRequested)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) PostFlightCheck(releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.TransactOpts, releaseOrMintIn, localAmount, blockConfirmationRequested)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "preflightCheck", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "setPolicyEngine", newPolicyEngine)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetPolicyEngine(&_AdvancedPoolHooks.TransactOpts, newPolicyEngine)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetPolicyEngine(&_AdvancedPoolHooks.TransactOpts, newPolicyEngine)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) SetThresholdAmount(opts *bind.TransactOpts, thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "setThresholdAmount", thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) SetThresholdAmount(thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetThresholdAmount(&_AdvancedPoolHooks.TransactOpts, thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) SetThresholdAmount(thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetThresholdAmount(&_AdvancedPoolHooks.TransactOpts, thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "transferOwnership", to)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.TransferOwnership(&_AdvancedPoolHooks.TransactOpts, to)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.TransferOwnership(&_AdvancedPoolHooks.TransactOpts, to)
}

type AdvancedPoolHooksAllowListAddIterator struct {
	Event *AdvancedPoolHooksAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAllowListAdd)
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
		it.Event = new(AdvancedPoolHooksAllowListAdd)
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

func (it *AdvancedPoolHooksAllowListAddIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListAddIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAllowListAddIterator{contract: _AdvancedPoolHooks.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAllowListAdd)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAllowListAdd(log types.Log) (*AdvancedPoolHooksAllowListAdd, error) {
	event := new(AdvancedPoolHooksAllowListAdd)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksAllowListRemoveIterator struct {
	Event *AdvancedPoolHooksAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAllowListRemove)
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
		it.Event = new(AdvancedPoolHooksAllowListRemove)
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

func (it *AdvancedPoolHooksAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListRemoveIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAllowListRemoveIterator{contract: _AdvancedPoolHooks.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAllowListRemove)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAllowListRemove(log types.Log) (*AdvancedPoolHooksAllowListRemove, error) {
	event := new(AdvancedPoolHooksAllowListRemove)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksCCVConfigUpdatedIterator struct {
	Event *AdvancedPoolHooksCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksCCVConfigUpdated)
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
		it.Event = new(AdvancedPoolHooksCCVConfigUpdated)
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

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksCCVConfigUpdated struct {
	RemoteChainSelector   uint64
	OutboundCCVs          []common.Address
	ThresholdOutboundCCVs []common.Address
	InboundCCVs           []common.Address
	ThresholdInboundCCVs  []common.Address
	Raw                   types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*AdvancedPoolHooksCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksCCVConfigUpdatedIterator{contract: _AdvancedPoolHooks.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksCCVConfigUpdated)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseCCVConfigUpdated(log types.Log) (*AdvancedPoolHooksCCVConfigUpdated, error) {
	event := new(AdvancedPoolHooksCCVConfigUpdated)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksOwnershipTransferRequestedIterator struct {
	Event *AdvancedPoolHooksOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksOwnershipTransferRequested)
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
		it.Event = new(AdvancedPoolHooksOwnershipTransferRequested)
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

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksOwnershipTransferRequestedIterator{contract: _AdvancedPoolHooks.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksOwnershipTransferRequested)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseOwnershipTransferRequested(log types.Log) (*AdvancedPoolHooksOwnershipTransferRequested, error) {
	event := new(AdvancedPoolHooksOwnershipTransferRequested)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksOwnershipTransferredIterator struct {
	Event *AdvancedPoolHooksOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksOwnershipTransferred)
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
		it.Event = new(AdvancedPoolHooksOwnershipTransferred)
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

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksOwnershipTransferredIterator{contract: _AdvancedPoolHooks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksOwnershipTransferred)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseOwnershipTransferred(log types.Log) (*AdvancedPoolHooksOwnershipTransferred, error) {
	event := new(AdvancedPoolHooksOwnershipTransferred)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksPolicyEngineSetIterator struct {
	Event *AdvancedPoolHooksPolicyEngineSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksPolicyEngineSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksPolicyEngineSet)
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
		it.Event = new(AdvancedPoolHooksPolicyEngineSet)
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

func (it *AdvancedPoolHooksPolicyEngineSetIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksPolicyEngineSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksPolicyEngineSet struct {
	OldPolicyEngine common.Address
	NewPolicyEngine common.Address
	Raw             types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterPolicyEngineSet(opts *bind.FilterOpts, oldPolicyEngine []common.Address, newPolicyEngine []common.Address) (*AdvancedPoolHooksPolicyEngineSetIterator, error) {

	var oldPolicyEngineRule []interface{}
	for _, oldPolicyEngineItem := range oldPolicyEngine {
		oldPolicyEngineRule = append(oldPolicyEngineRule, oldPolicyEngineItem)
	}
	var newPolicyEngineRule []interface{}
	for _, newPolicyEngineItem := range newPolicyEngine {
		newPolicyEngineRule = append(newPolicyEngineRule, newPolicyEngineItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "PolicyEngineSet", oldPolicyEngineRule, newPolicyEngineRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksPolicyEngineSetIterator{contract: _AdvancedPoolHooks.contract, event: "PolicyEngineSet", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchPolicyEngineSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksPolicyEngineSet, oldPolicyEngine []common.Address, newPolicyEngine []common.Address) (event.Subscription, error) {

	var oldPolicyEngineRule []interface{}
	for _, oldPolicyEngineItem := range oldPolicyEngine {
		oldPolicyEngineRule = append(oldPolicyEngineRule, oldPolicyEngineItem)
	}
	var newPolicyEngineRule []interface{}
	for _, newPolicyEngineItem := range newPolicyEngine {
		newPolicyEngineRule = append(newPolicyEngineRule, newPolicyEngineItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "PolicyEngineSet", oldPolicyEngineRule, newPolicyEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksPolicyEngineSet)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "PolicyEngineSet", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParsePolicyEngineSet(log types.Log) (*AdvancedPoolHooksPolicyEngineSet, error) {
	event := new(AdvancedPoolHooksPolicyEngineSet)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "PolicyEngineSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksThresholdAmountSetIterator struct {
	Event *AdvancedPoolHooksThresholdAmountSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksThresholdAmountSet)
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
		it.Event = new(AdvancedPoolHooksThresholdAmountSet)
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

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksThresholdAmountSet struct {
	ThresholdAmount *big.Int
	Raw             types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterThresholdAmountSet(opts *bind.FilterOpts) (*AdvancedPoolHooksThresholdAmountSetIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "ThresholdAmountSet")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksThresholdAmountSetIterator{contract: _AdvancedPoolHooks.contract, event: "ThresholdAmountSet", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchThresholdAmountSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksThresholdAmountSet) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "ThresholdAmountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksThresholdAmountSet)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "ThresholdAmountSet", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseThresholdAmountSet(log types.Log) (*AdvancedPoolHooksThresholdAmountSet, error) {
	event := new(AdvancedPoolHooksThresholdAmountSet)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "ThresholdAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (AdvancedPoolHooksAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (AdvancedPoolHooksAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (AdvancedPoolHooksCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (AdvancedPoolHooksOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AdvancedPoolHooksOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AdvancedPoolHooksPolicyEngineSet) Topic() common.Hash {
	return common.HexToHash("0xfb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e")
}

func (AdvancedPoolHooksThresholdAmountSet) Topic() common.Hash {
	return common.HexToHash("0x80dc2a1a49dda9f8bd85c1c376266e011db6448050b7bfd5c2f423e162c11145")
}

func (_AdvancedPoolHooks *AdvancedPoolHooks) Address() common.Address {
	return _AdvancedPoolHooks.address
}

type AdvancedPoolHooksInterface interface {
	CheckAllowList(opts *bind.CallOpts, sender common.Address) error

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetPolicyEngine(opts *bind.CallOpts) (common.Address, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetThresholdAmount(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error)

	PostFlightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, blockConfirmationRequested uint16) (*types.Transaction, error)

	PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	SetThresholdAmount(opts *bind.TransactOpts, thresholdAmount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*AdvancedPoolHooksAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*AdvancedPoolHooksAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*AdvancedPoolHooksCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*AdvancedPoolHooksCCVConfigUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AdvancedPoolHooksOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AdvancedPoolHooksOwnershipTransferred, error)

	FilterPolicyEngineSet(opts *bind.FilterOpts, oldPolicyEngine []common.Address, newPolicyEngine []common.Address) (*AdvancedPoolHooksPolicyEngineSetIterator, error)

	WatchPolicyEngineSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksPolicyEngineSet, oldPolicyEngine []common.Address, newPolicyEngine []common.Address) (event.Subscription, error)

	ParsePolicyEngineSet(log types.Log) (*AdvancedPoolHooksPolicyEngineSet, error)

	FilterThresholdAmountSet(opts *bind.FilterOpts) (*AdvancedPoolHooksThresholdAmountSetIterator, error)

	WatchThresholdAmountSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksThresholdAmountSet) (event.Subscription, error)

	ParseThresholdAmountSet(log types.Log) (*AdvancedPoolHooksThresholdAmountSet, error)

	Address() common.Address
}
