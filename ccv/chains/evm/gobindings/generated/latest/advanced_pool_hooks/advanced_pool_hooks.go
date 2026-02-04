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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct AdvancedPoolHooks.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkAllowList\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuthorizedCallersEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThresholdAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postFlightCheck\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preflightCheck\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngine\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngineAllowFailedDetach\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThresholdAmount\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateCaller\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholdOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholdInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineSet\",\"inputs\":[{\"name\":\"oldPolicyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdAmountSet\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuthorizedCallersNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyUnderThresholdCCVsForThresholdCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyEngineDetachFailed\",\"inputs\":[{\"name\":\"oldPolicyEngine\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c0806040523461036357612d49803803809161001c8285610368565b833981016080828203126103635781516001600160401b038111610363578161004691840161039f565b6020830151926100586040820161038b565b60608201519093906001600160401b03811161036357610078920161039f565b331561035257600180546001600160a01b0319163317905560405160209290916100a28484610368565b60008352600036813760408051969087016001600160401b0381118882101761033c57604052818752838588015260005b845181101561013a576001906001600160a01b036100f1828861040f565b5116876100fd826107b2565b61010a575b5050016100d3565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13887610102565b508493508587519260005b84518110156101b6576001600160a01b03610160828761040f565b51169081156101a5577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef888361019760019561073a565b50604051908152a101610145565b6342bcdf7f60e11b60005260046000fd5b5085858051151580608052610217575b50506101da9260065551151560a052610439565b6040516125029081610847823960805181818161026a01528181610ffc01526119f0015260a051818181610a780152818161125601526119850152f35b9091604051916102278484610368565b6000835260003681376080511561032b5760005b83518110156102a2576001906001600160a01b03610259828761040f565b51168661026582610647565b610272575b50500161023b565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1888661026a565b5091509260005b825181101561031d576001906001600160a01b036102c7828661040f565b5116801561031757856102d982610779565b6102e7575b50505b016102a9565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a187856102de565b506102e1565b50929150506101da846101c6565b6335f4a7b360e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761033c57604052565b51906001600160a01b038216820361036357565b9080601f83011215610363578151916001600160401b03831161033c578260051b9060208201936103d36040519586610368565b845260208085019282010192831161036357602001905b8282106103f75750505090565b602080916104048461038b565b8152019101906103ea565b80518210156104235760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6007546001600160a01b03918216911660008282146105eb57816104e7575b600780546001600160a01b0319168417905582610496575b807ffb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e91a3565b823b156104d957604051631100482d60e01b8152818160048183885af180156104dc576104c4575b50610470565b6104cf828092610368565b6104d957386104be565b80fd5b6040513d84823e3d90fd5b918193913b156105e757604051628950d760e61b8152838160048183895af190816105d3575b506105ca5750503d156105c1573d906001600160401b0382116105ad5760405191610542601f8201601f191660200184610368565b82523d81602084013e915b604051928391635c3a3f6360e01b8352600483015260406024830152825192836044840152815b84811061059557505091606492838284010152601f80199101168101030190fd5b60208282018101516064888401015286945001610574565b634e487b7160e01b81526041600452602490fd5b6060909161054d565b91929092610458565b846105e091959295610368565b923861050d565b8280fd5b505050565b80548210156104235760005260206000200190600090565b8054801561063157600019019061061f82826105f0565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b60008181526005602052604090205480156107085760001981018181116106f2576004546000198101919082116106f2578181036106a1575b50505061068d6004610608565b600052600560205260006040812055600190565b6106da6106b26106c39360046105f0565b90549060031b1c92839260046105f0565b819391549060031b91821b91600019901b19161790565b90556000526005602052604060002055388080610680565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8054906801000000000000000082101561033c57816106c3916001610736940181556105f0565b9055565b806000526003602052604060002054156000146107735761075c81600261070f565b600254906000526003602052604060002055600190565b50600090565b806000526005602052604060002054156000146107735761079b81600461070f565b600454906000526005602052604060002055600190565b60008181526003602052604090205480156107085760001981018181116106f2576002546000198101919082116106f25780820361080c575b5050506107f86002610608565b600052600360205260006040812055600190565b61082e61081d6106c39360026105f0565b90549060031b1c92839260026105f0565b905560005260036020526040600020553880806107eb56fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461127b578063201b52c3146112205780632451a6271461118b5780634ef34bc01461114457806354c8a4f314610f6a5780635c3af7ca14610ec75780635eff3bf714610e445780636135b08514610dfd5780636831731214610dab57806379ba509714610cc257806389720a6214610c2d5780638da5cb5b14610bdb57806391a2749a14610992578063961e2e7c14610928578063a7cd63b714610883578063cc488ac91461084c578063ce07c7c81461080b578063d966866b1461028f578063e0351e1314610234578063f2fde38b146101445763f72c071b1461010357600080fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f576020600654604051908152f35b600080fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5773ffffffffffffffffffffffffffffffffffffffff61019061143e565b610198611b82565b1633811461020a57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f576102de903690600401611482565b6102e6611b82565b6000905b8082106102f357005b6102fe828285611a70565b359067ffffffffffffffff8216820361013f5761032961031f848387611a70565b6020810190611adf565b6103436103398685899599611a70565b6040810190611adf565b949061035d610353888787611a70565b6060810190611adf565b61037561036b8a8989611a70565b6080810190611adf565b93909861038c6103878d89369161152c565b612068565b61039a61038736858761152c565b806107d9575b84610777575b6040519b6103b38d6112f8565b6103be36898361152c565b8d528c8b6103cd36858761152c565b602083019081526103f26103e236898b61152c565b92604085019384528a369161152c565b606084015267ffffffffffffffff8a1660005260086020526040600020925180519067ffffffffffffffff82116106475768010000000000000000821161064757845482865580831061074e575b5060200184600052602060002060005b83811061072457505050506001839e9c9d9e0190519081519167ffffffffffffffff8311610647576801000000000000000083116106475781548383558084106106f3575b5060200190600052602060002060005b8381106106c957505050506002820190519081519167ffffffffffffffff8311610647576801000000000000000083116106475781548383558084106106a0575b509e9f939495969798999a9b9c9d9e60200190600052602060002060005b838110610676575050505060036060919e9c9d9e019101519081519167ffffffffffffffff83116106475768010000000000000000831161064757815483835580841061061e575b5060200190600052602060002060005b8381106105f457505050506105d96080956105e9956105cb7fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a966105bd67ffffffffffffffff976040519d8d8f9e8f9081520191611b33565b918b830360208d0152611b33565b9188830360408a0152611b33565b9285840360608701521696611b33565b0390a20190916102ea565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161055c565b8260005283602060002091820191015b81811061063b575061054c565b6000815560010161062e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610504565b8260005283602060002091820191015b8181106106bd57506104e6565b600081556001016106b0565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016104a5565b9d9f9e9d8260005283602060002091820191015b81811061071857509f9d9e9f610495565b60008155600101610707565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610450565b8560005282602060002091820191015b81811061076b5750610440565b6000815560010161075e565b82156107af5761078b61038736878d61152c565b6107aa61079936858761152c565b6107a436888e61152c565b90612132565b6103a6565b7f1d56c21d0000000000000000000000000000000000000000000000000000000060005260046000fd5b86156107af576107ed61038736838561152c565b6108066107fb36898f61152c565b6107a436848661152c565b6103a0565b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761084a61084561143e565b6119ee565b005b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761084a611983565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760405180602060045491828152019060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9060005b8181106109125761090e8561090281870382611314565b604051918291826113ee565b0390f35b82548452602090930192600192830192016108eb565b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f577f80dc2a1a49dda9f8bd85c1c376266e011db6448050b7bfd5c2f423e162c111456020600435610985611b82565b80600655604051908152a1005b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261013f57604051906040820182811067ffffffffffffffff82111761064757604052806004013567ffffffffffffffff811161013f57610a419060043691840101611580565b825260248101359067ffffffffffffffff821161013f576004610a679236920101611580565b60208201908152610a76611b82565b7f000000000000000000000000000000000000000000000000000000000000000015610bb157519060005b8251811015610b14578073ffffffffffffffffffffffffffffffffffffffff610acc60019386611ee9565b5116610ad781612425565b610ae3575b5001610aa1565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610adc565b505160005b815181101561084a5773ffffffffffffffffffffffffffffffffffffffff610b418284611ee9565b5116908115610b87577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083610b796001956123ec565b50604051908152a101610b19565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fec802d6f0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461013f5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57610c6461143e565b5060243567ffffffffffffffff8116810361013f57610c816114d5565b5060843567ffffffffffffffff811161013f57610ca29036906004016114e6565b505060a43590600282101561013f5761090e9161090291604435906118eb565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760005473ffffffffffffffffffffffffffffffffffffffff81163303610d81577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761084a610e3761143e565b610e3f611b82565b611d32565b3461013f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261013f5761084a90610ebe6114c4565b50600401611800565b3461013f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261013f57610f3c6114b3565b506044359067ffffffffffffffff821161013f57610f6161084a9236906004016114e6565b916004016116ad565b3461013f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043567ffffffffffffffff811161013f57610fb9903690600401611482565b6024359067ffffffffffffffff821161013f57610ff2610fe0610ffa933690600401611482565b949092610feb611b82565b369161152c565b92369161152c565b7f00000000000000000000000000000000000000000000000000000000000000001561111a5760005b8251811015611096578073ffffffffffffffffffffffffffffffffffffffff61104e60019386611ee9565b511661105981612276565b611065575b5001611023565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a18461105e565b5060005b815181101561084a578073ffffffffffffffffffffffffffffffffffffffff6110c560019385611ee9565b51168015611114576110d6816123ad565b6110e3575b505b0161109a565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1836110db565b506110dd565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761084a61117e61143e565b611186611b82565b611bcd565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061120a5761090e8561090281870382611314565b82548452602090930192600192830192016111f3565b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461013f5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5761090e60408051906112bc8183611314565b601b82527f416476616e636564506f6f6c486f6f6b7320312e372e302d646576000000000060208301525191829160208352602083019061138f565b6080810190811067ffffffffffffffff82111761064757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761064757604052565b67ffffffffffffffff811161064757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106113d95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161139a565b602060408183019282815284518094520192019060005b8181106114125750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611405565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361013f57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361013f57565b9181601f8401121561013f5782359167ffffffffffffffff831161013f576020808501948460051b01011161013f57565b6024359061ffff8216820361013f57565b6044359061ffff8216820361013f57565b6064359061ffff8216820361013f57565b9181601f8401121561013f5782359167ffffffffffffffff831161013f576020838186019501011161013f57565b67ffffffffffffffff81116106475760051b60200190565b92919061153881611514565b936115466040519586611314565b602085838152019160051b810192831161013f57905b82821061156857505050565b6020809161157584611461565b81520191019061155c565b9080601f8301121561013f5781602061159b9335910161152c565b90565b9190916115aa81611355565b6115b76040519182611314565b8093828252826004011161013f5781816000936004602080950137010152565b9291926115e382611355565b916115f16040519384611314565b82948184528183011161013f578281602093846000960137010152565b9061159b91602081527fffffffff00000000000000000000000000000000000000000000000000000000825116602082015273ffffffffffffffffffffffffffffffffffffffff6020830151166040820152606061167a604084015160808385015260a084019061138f565b9201519060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08285030191015261138f565b916116b6611983565b6040600093013573ffffffffffffffffffffffffffffffffffffffff811681036117f6576116e3906119ee565b73ffffffffffffffffffffffffffffffffffffffff600754169182156117fa57366004116117f657611780906040519261171c846112f8565b7fffffffff000000000000000000000000000000000000000000000000000000008635168452336020850152611774367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360161159e565b604085015236916115d7565b6060820152813b156117f2576117c8839283926040519485809481937fc2098e080000000000000000000000000000000000000000000000000000000083526004830161160e565b03925af180156117e7576117da575050565b816117e491611314565b50565b6040513d84823e3d90fd5b8280fd5b8380fd5b50505050565b611808611983565b6007549073ffffffffffffffffffffffffffffffffffffffff6000921690811561189657366004116117f25760e0810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156117f6570180359067ffffffffffffffff82116117f65760200181360381136117f657611780906040519261171c846112f8565b505050565b906020825491828152019160005260206000209060005b8181106118bf5750505090565b825473ffffffffffffffffffffffffffffffffffffffff168452602090930192600192830192016118b2565b67ffffffffffffffff1660005260086020526040600020916002811015611954576001146119395761159b916001604051916119328361192b818461189b565b0384611314565b0190611f39565b61159b916003604051916119328361192b816002850161189b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f00000000000000000000000000000000000000000000000000000000000000006119aa57565b336000526003602052604060002054156119c057565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f0000000000000000000000000000000000000000000000000000000000000000611a165750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600560205260406000205415611a435750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190811015611ab05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561013f570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561013f570180359067ffffffffffffffff821161013f57602001918160051b3603831361013f57565b9160209082815201919060005b818110611b4d5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff611b7488611461565b168152019401929101611b40565b73ffffffffffffffffffffffffffffffffffffffff600154163303611ba357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60075473ffffffffffffffffffffffffffffffffffffffff918216911660008282146118965781611cb1575b827fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075582611c4f575b807ffb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e91a3565b823b15611cae576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183885af180156117e757611c96575b50611c29565b81611ca091611314565b80600012611cae5738611c90565b80fd5b813b15611cae57604051907f225435c0000000000000000000000000000000000000000000000000000000008252808260048183875af19182611d22575b5090611d1d573d15611d1d573d611d0581611355565b90611d136040519283611314565b81528160203d92013e5b611bf9565b81611d2c91611314565b38611cef565b60075473ffffffffffffffffffffffffffffffffffffffff918216911660008282146118965781611e0e575b827fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075582611db357807ffb3c698262b8ff219e7285565d54621a2e73556110f0249aeb7b5de1b0b9d32e91a3565b823b15611cae576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183885af180156117e757611df95750611c29565b611e04828092611314565b611cae5738611c90565b929091823b156117f6576040517f225435c0000000000000000000000000000000000000000000000000000000008152848160048183885af19081611ed5575b50611ecd5750503d15611ec4573d611e6581611355565b90611e736040519283611314565b8152809260203d92013e5b611ec06040519283927f5c3a3f63000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061138f565b0390fd5b60609150611e7e565b919092611d5e565b85611ee291969296611314565b9338611e4e565b8051821015611ab05760209160051b010190565b91908201809211611f0a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91600654801515918261205d575b5050611f51575090565b90611f6c611f65926040519384809261189b565b0383611314565b815180611f795750905090565b611f84908251611efd565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611fc8611fb286611514565b95611fc06040519788611314565b808752611514565b0136602086013760005b8251811015612010578073ffffffffffffffffffffffffffffffffffffffff611ffd60019386611ee9565b51166120098288611ee9565b5201611fd2565b509160005b8151811015611896578073ffffffffffffffffffffffffffffffffffffffff61204060019385611ee9565b5116612056612050838751611efd565b88611ee9565b5201612015565b101590503880611f47565b805160005b81811061207957505050565b60018101808211611f0a575b828110612095575060010161206d565b73ffffffffffffffffffffffffffffffffffffffff6120b48386611ee9565b511673ffffffffffffffffffffffffffffffffffffffff6120d58387611ee9565b5116146120e457600101612085565b73ffffffffffffffffffffffffffffffffffffffff6121038386611ee9565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9081519080519060005b83811061214a575050505050565b60005b83811061215d575060010161213c565b73ffffffffffffffffffffffffffffffffffffffff61217c8388611ee9565b511673ffffffffffffffffffffffffffffffffffffffff61219d8386611ee9565b5116146121ac5760010161214d565b73ffffffffffffffffffffffffffffffffffffffff6121038388611ee9565b8054821015611ab05760005260206000200190600090565b80548015612247577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061221882826121cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260056020526040902054801561237b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611f0a57600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611f0a5781810361230c575b5050506122f860046121e3565b600052600560205260006040812055600190565b61236361231d61232e9360046121cb565b90549060031b1c92839260046121cb565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260056020526040600020553880806122eb565b5050600090565b80549068010000000000000000821015610647578161232e9160016123a9940181556121cb565b9055565b806000526005602052604060002054156000146123e6576123cf816004612382565b600454906000526005602052604060002055600190565b50600090565b806000526003602052604060002054156000146123e65761240e816002612382565b600254906000526003602052604060002055600190565b600081815260036020526040902054801561237b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611f0a57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611f0a578082036124bb575b5050506124a760026121e3565b600052600360205260006040812055600190565b6124dd6124cc61232e9360026121cb565b90549060031b1c92839260026121cb565b9055600052600360205260406000205538808061249a56fea164736f6c634300081a000a",
}

var AdvancedPoolHooksABI = AdvancedPoolHooksMetaData.ABI

var AdvancedPoolHooksBin = AdvancedPoolHooksMetaData.Bin

func DeployAdvancedPoolHooks(auth *bind.TransactOpts, backend bind.ContractBackend, allowlist []common.Address, thresholdAmountForAdditionalCCVs *big.Int, policyEngine common.Address, authorizedCallers []common.Address) (common.Address, *types.Transaction, *AdvancedPoolHooks, error) {
	parsed, err := AdvancedPoolHooksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdvancedPoolHooksBin), backend, allowlist, thresholdAmountForAdditionalCCVs, policyEngine, authorizedCallers)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllAuthorizedCallers(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllAuthorizedCallers(&_AdvancedPoolHooks.CallOpts)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAuthorizedCallersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAuthorizedCallersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAuthorizedCallersEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAuthorizedCallersEnabled(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAuthorizedCallersEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAuthorizedCallersEnabled(&_AdvancedPoolHooks.CallOpts)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) TypeAndVersion() (string, error) {
	return _AdvancedPoolHooks.Contract.TypeAndVersion(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) TypeAndVersion() (string, error) {
	return _AdvancedPoolHooks.Contract.TypeAndVersion(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) ValidateCaller(opts *bind.CallOpts) error {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "validateCaller")

	if err != nil {
		return err
	}

	return err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ValidateCaller() error {
	return _AdvancedPoolHooks.Contract.ValidateCaller(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) ValidateCaller() error {
	return _AdvancedPoolHooks.Contract.ValidateCaller(&_AdvancedPoolHooks.CallOpts)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAuthorizedCallerUpdates(&_AdvancedPoolHooks.TransactOpts, authorizedCallerArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAuthorizedCallerUpdates(&_AdvancedPoolHooks.TransactOpts, authorizedCallerArgs)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) PostFlightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "postFlightCheck", releaseOrMintIn, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PostFlightCheck(releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.TransactOpts, releaseOrMintIn, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) PostFlightCheck(releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.TransactOpts, releaseOrMintIn, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "preflightCheck", lockOrBurnIn, arg1, tokenArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.TransactOpts, lockOrBurnIn, arg1, tokenArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.TransactOpts, lockOrBurnIn, arg1, tokenArgs)
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

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "setPolicyEngineAllowFailedDetach", newPolicyEngine)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetPolicyEngineAllowFailedDetach(&_AdvancedPoolHooks.TransactOpts, newPolicyEngine)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetPolicyEngineAllowFailedDetach(&_AdvancedPoolHooks.TransactOpts, newPolicyEngine)
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

type AdvancedPoolHooksAuthorizedCallerAddedIterator struct {
	Event *AdvancedPoolHooksAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAuthorizedCallerAdded)
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
		it.Event = new(AdvancedPoolHooksAuthorizedCallerAdded)
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

func (it *AdvancedPoolHooksAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*AdvancedPoolHooksAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAuthorizedCallerAddedIterator{contract: _AdvancedPoolHooks.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAuthorizedCallerAdded)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAuthorizedCallerAdded(log types.Log) (*AdvancedPoolHooksAuthorizedCallerAdded, error) {
	event := new(AdvancedPoolHooksAuthorizedCallerAdded)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksAuthorizedCallerRemovedIterator struct {
	Event *AdvancedPoolHooksAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAuthorizedCallerRemoved)
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
		it.Event = new(AdvancedPoolHooksAuthorizedCallerRemoved)
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

func (it *AdvancedPoolHooksAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*AdvancedPoolHooksAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAuthorizedCallerRemovedIterator{contract: _AdvancedPoolHooks.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAuthorizedCallerRemoved)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*AdvancedPoolHooksAuthorizedCallerRemoved, error) {
	event := new(AdvancedPoolHooksAuthorizedCallerRemoved)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (AdvancedPoolHooksAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (AdvancedPoolHooksAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
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

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAuthorizedCallersEnabled(opts *bind.CallOpts) (bool, error)

	GetPolicyEngine(opts *bind.CallOpts) (common.Address, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetThresholdAmount(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ValidateCaller(opts *bind.CallOpts) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error)

	PostFlightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) (*types.Transaction, error)

	PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, tokenArgs []byte) (*types.Transaction, error)

	SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	SetThresholdAmount(opts *bind.TransactOpts, thresholdAmount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*AdvancedPoolHooksAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*AdvancedPoolHooksAllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*AdvancedPoolHooksAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*AdvancedPoolHooksAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*AdvancedPoolHooksAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*AdvancedPoolHooksAuthorizedCallerRemoved, error)

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
