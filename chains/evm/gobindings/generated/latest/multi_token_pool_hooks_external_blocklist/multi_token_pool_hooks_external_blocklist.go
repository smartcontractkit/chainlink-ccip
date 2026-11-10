// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multi_token_pool_hooks_external_blocklist

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

var MultiTokenPoolHooksExternalBlocklistMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"},{\"name\":\"blocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"},{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkAllowList\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlocklist\",\"inputs\":[],\"outputs\":[{\"name\":\"blocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postflightCheck\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"localAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preflightCheck\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amountPostFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlocklist\",\"inputs\":[{\"name\":\"newBlocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngine\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngineAllowFailedDetach\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlocklistSet\",\"inputs\":[{\"name\":\"oldBlocklist\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IAddressBlocklist\"},{\"name\":\"newBlocklist\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IAddressBlocklist\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineAttached\",\"inputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IPolicyEngine\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineDetachFailed\",\"inputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IPolicyEngine\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlocklist\",\"inputs\":[{\"name\":\"blocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"}]},{\"type\":\"error\",\"name\":\"LocalAddressBlocked\",\"inputs\":[{\"name\":\"blocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyEngineDetachReverted\",\"inputs\":[{\"name\":\"oldPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}]},{\"type\":\"error\",\"name\":\"RemoteAddressBlocked\",\"inputs\":[{\"name\":\"blocklist\",\"type\":\"address\",\"internalType\":\"contract IAddressBlocklist\"},{\"name\":\"account\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a08060405234610362576125a9803803809161001c8285610367565b83398101906080818303126103625780516001600160401b038111610362578261004791830161038a565b602082015190916001600160a01b0382168203610362576040810151936001600160a01b03851685036103625760608201516001600160401b03811161036257610091920161038a565b91331561035157600180546001600160a01b0319163317905560405160209490926100bc8685610367565b60008452600036813760408051959086016001600160401b0381118782101761033b576040528552838686015260005b8451811015610153576001906001600160a01b0361010a8288610484565b5116886101168261077f565b610123575b5050016100ec565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388861011b565b5085935084519160005b83518110156101ce576001600160a01b036101788286610484565b51169081156101bd577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef87836101af600195610707565b50604051908152a10161015d565b6342bcdf7f60e11b60005260046000fd5b5083858151151580608052610216575b6101f0836101eb866104ae565b610401565b604051611d95908161081482396080518181816102020152818161098a0152610fb50152f35b9091604051906102268383610367565b6000825260003681376080511561032a5760005b82518110156102a1576001906001600160a01b036102588286610484565b51168561026482610614565b610271575b50500161023a565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a18785610269565b5092905060005b815181101561031c576001906001600160a01b036102c68285610484565b5116801561031657846102d882610746565b6102e6575b50505b016102a8565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a186846102dd565b506102e0565b5050506101eb6101f06101de565b6335f4a7b360e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761033b57604052565b81601f82011215610362578051916001600160401b03831161033b578260051b9160208301936103bd6040519586610367565b845260208085019382010191821161036257602001915b8183106103e15750505090565b82516001600160a01b0381168103610362578152602092830192016103d4565b6001600160a01b03168015158061047b575b610467576007546001600160a01b03811690828214610462576001600160a01b03191682176007557f7e053cdc9069fe4f629b6b3fa2a01bb53a9a8305ec11830b259d95e9e75b7304600080a3565b505050565b6357129f5360e01b60005260045260246000fd5b50803b15610413565b80518210156104985760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6006546001600160a01b039182169160009116828114610462578061055e575b50600680546001600160a01b031916831790558161050d575b807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b1561055057604051631100482d60e01b8152818160048183875af180156105535761053b575b506104e7565b610546828092610367565b6105505738610535565b80fd5b6040513d84823e3d90fd5b9091813b156105b957604051628950d760e61b8152838160048183875af190816105a5575b5061059c5763e047565560e01b83526004829052602483fd5b919050386104ce565b846105b291959295610367565b9238610583565b8280fd5b80548210156104985760005260206000200190600090565b805480156105fe5760001901906105ec82826105bd565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b60008181526005602052604090205480156106d55760001981018181116106bf576004546000198101919082116106bf5781810361066e575b50505061065a60046105d5565b600052600560205260006040812055600190565b6106a761067f6106909360046105bd565b90549060031b1c92839260046105bd565b819391549060031b91821b91600019901b19161790565b9055600052600560205260406000205538808061064d565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8054906801000000000000000082101561033b5781610690916001610703940181556105bd565b9055565b80600052600360205260406000205415600014610740576107298160026106dc565b600254906000526003602052604060002055600190565b50600090565b80600052600560205260406000205415600014610740576107688160046106dc565b600454906000526005602052604060002055600190565b60008181526003602052604090205480156106d55760001981018181116106bf576002546000198101919082116106bf578082036107d9575b5050506107c560026105d5565b600052600360205260006040812055600190565b6107fb6107ea6106909360026105bd565b90549060031b1c92839260026105bd565b905560005260036020526040600020553880806107b856fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714610c195780631bcfb97c14610bc75780632451a62714610b325780634ef34bc014610ad257806354c8a4f3146108f85780636135b08514610898578063637115741461080857806368317312146107b657806379ba5097146106cd5780638da5cb5b1461067b57806391a2749a14610453578063a7cd63b7146103ae578063a8027c0f146102c8578063aef18ae714610268578063ce07c7c814610227578063e0351e13146101cc5763f2fde38b146100d757600080fd5b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75773ffffffffffffffffffffffffffffffffffffffff610123610e54565b61012b611051565b1633811461019d57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c757610266610261610e54565b610fb3565b005b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043573ffffffffffffffffffffffffffffffffffffffff811681036101c757610266906102c3611051565b611922565b346101c75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81116101c7578060040160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101c757610341610dc8565b506044359167ffffffffffffffff83116101c757366023840112156101c75782600401359267ffffffffffffffff84116101c75736602485830101116101c75760446103a89261039b6102669660246103a19501876118d0565b01610f41565b9180610f62565b91611612565b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760405180602060045491828152019060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9060005b81811061043d576104398561042d81870382610e13565b60405191829182610d18565b0390f35b8254845260209093019260019283019201610416565b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81116101c75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c757604051906040820182811067ffffffffffffffff82111761064c57604052806004013567ffffffffffffffff81116101c7576105029060043691840101610ee9565b825260248101359067ffffffffffffffff82116101c75760046105289236920101610ee9565b60208201908152610537611051565b519060005b82518110156105af578073ffffffffffffffffffffffffffffffffffffffff61056760019386611389565b511661057281611cb8565b61057e575b500161053c565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610577565b505160005b81518110156102665773ffffffffffffffffffffffffffffffffffffffff6105dc8284611389565b5116908115610622577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083610614600195611c7f565b50604051908152a1016105b4565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760005473ffffffffffffffffffffffffffffffffffffffff8116330361078c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c757602073ffffffffffffffffffffffffffffffffffffffff60065416604051908152f35b346101c75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81116101c757806004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126101c7576103a86103a160446102669461088e610d99565b5061039b856114db565b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043573ffffffffffffffffffffffffffffffffffffffff811681036101c757610266906108f3611051565b61121b565b346101c75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81116101c757610947903690600401610d68565b6024359067ffffffffffffffff82116101c75761098061096e610988933690600401610d68565b949092610979611051565b3691610e77565b923691610e77565b7f000000000000000000000000000000000000000000000000000000000000000015610aa85760005b8251811015610a24578073ffffffffffffffffffffffffffffffffffffffff6109dc60019386611389565b51166109e781611ada565b6109f3575b50016109b1565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1846109ec565b5060005b8151811015610266578073ffffffffffffffffffffffffffffffffffffffff610a5360019385611389565b51168015610aa257610a6481611c40565b610a71575b505b01610a28565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183610a69565b50610a6b565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043573ffffffffffffffffffffffffffffffffffffffff811681036101c75761026690610b2d611051565b61109c565b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610bb1576104398561042d81870382610e13565b8254845260209093019260019283019201610b9a565b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c757602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346101c75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c757610439604051610c59606082610e13565b602e81527f4d756c7469546f6b656e506f6f6c486f6f6b7345787465726e616c426c6f636b60208201527f6c69737420322e312e302d64657600000000000000000000000000000000000060408201526040519182916020835260208301905b919082519283825260005b848110610d035750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610cc4565b602060408183019282815284518094520192019060005b818110610d3c5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610d2f565b9181601f840112156101c75782359167ffffffffffffffff83116101c7576020808501948460051b0101116101c757565b604435907fffffffff00000000000000000000000000000000000000000000000000000000821682036101c757565b602435907fffffffff00000000000000000000000000000000000000000000000000000000821682036101c757565b6080810190811067ffffffffffffffff82111761064c57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761064c57604052565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c757565b90929167ffffffffffffffff841161064c578360051b916020604051610e9f82860182610e13565b80968152019281019182116101c757915b818310610ebc57505050565b823573ffffffffffffffffffffffffffffffffffffffff811681036101c757815260209283019201610eb0565b9080601f830112156101c757816020610f0493359101610e77565b90565b67ffffffffffffffff811161064c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101c75790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c7570180359067ffffffffffffffff82116101c7576020019181360383136101c757565b7f000000000000000000000000000000000000000000000000000000000000000080611025575b610fe15750565b73ffffffffffffffffffffffffffffffffffffffff907fd0d25976000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff8116600052600560205260406000205415610fda565b73ffffffffffffffffffffffffffffffffffffffff60015416330361107257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60065473ffffffffffffffffffffffffffffffffffffffff8060009216921691808314611216578061118d575b50817fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065581611120575b807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b1561117f576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183875af1801561118257611167575b506110fa565b8161117191610e13565b8060001261117f5738611161565b80fd5b6040513d84823e3d90fd5b90813b1561117f57604051907f225435c0000000000000000000000000000000000000000000000000000000008252808260048183875af19182611206575b50906111ff57907fdbc729cc8d59cb8abc2a107be310a639f87b7673410abc028f2a3db03ff7b5b2600080a25b386110c9565b90506111f9565b8161121091610e13565b386111cc565b505050565b600073ffffffffffffffffffffffffffffffffffffffff806006541692169180831461121657806112f8575b50817fffffffffffffffffffffffff000000000000000000000000000000000000000060065416176006558161129d57807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b1561117f576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183875af18015611182576112e357506110fa565b6112ee828092610e13565b61117f5738611161565b9091813b15611385576040517f225435c0000000000000000000000000000000000000000000000000000000008152838160048183875af19081611371575b5061136857602483837fe0475655000000000000000000000000000000000000000000000000000000008252600452fd5b91905038611247565b8461137e91959295610e13565b9238611337565b8280fd5b805182101561139d5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190916113d881610f07565b6113e56040519182610e13565b809382825282600401116101c75781816000936004602080950137010152565b92919261141182610f07565b9161141f6040519384610e13565b8294818452818301116101c7578281602093846000960137010152565b90610f0491602081527fffffffff00000000000000000000000000000000000000000000000000000000825116602082015273ffffffffffffffffffffffffffffffffffffffff602083015116604082015260606114a8604084015160808385015260a0840190610cb9565b9201519060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610cb9565b6114e36119eb565b6006549073ffffffffffffffffffffffffffffffffffffffff600092169081156112165736600411611385576115208160e0611593930190610f62565b91906040519261152f84610df7565b7fffffffff000000000000000000000000000000000000000000000000000000008635168452336020850152611587367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36016113cc565b60408501523691611405565b6060820152813b15611385576115db839283926040519485809481937fc2098e080000000000000000000000000000000000000000000000000000000083526004830161143c565b03925af18015611182576115ed575050565b816115f791610e13565b50565b908160209103126101c7575180151581036101c75790565b909173ffffffffffffffffffffffffffffffffffffffff600754169182156118ca5773ffffffffffffffffffffffffffffffffffffffff604051917ffbac39510000000000000000000000000000000000000000000000000000000083521690816004820152602081602481875afa9081156117dd576000916118ab575b5061187a57506014810361181d5782357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811690601483106117e9575b5060601c5b73ffffffffffffffffffffffffffffffffffffffff604051917ffbac3951000000000000000000000000000000000000000000000000000000008352166004820152602081602481865afa9081156117dd576000916117ae575b5061173657505050565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836064946040519687957fb6ed5ab600000000000000000000000000000000000000000000000000000000875260048701526040602487015281604487015286860137600085828601015201168101030190fd5b6117d0915060203d6020116117d6575b6117c88183610e13565b8101906115fa565b3861172c565b503d6117be565b6040513d6000823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008092508360140360031b1b1616386116cd565b602081036112165773ffffffffffffffffffffffffffffffffffffffff83356020831061184b575b166116d2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360031b1b16611845565b827f4b38d24b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b6118c4915060203d6020116117d6576117c88183610e13565b38611690565b50505050565b61026160406118e19261039b6119eb565b6006549173ffffffffffffffffffffffffffffffffffffffff600093169182156118ca573660041161191e57611593906040519261152f84610df7565b8380fd5b73ffffffffffffffffffffffffffffffffffffffff16801515806119e2575b6119b55760075473ffffffffffffffffffffffffffffffffffffffff811690818314611216577fffffffffffffffffffffffff00000000000000000000000000000000000000001682176007557f7e053cdc9069fe4f629b6b3fa2a01bb53a9a8305ec11830b259d95e9e75b7304600080a3565b7f57129f530000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50803b15611941565b33600052600360205260406000205415611a0157565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b805482101561139d5760005260206000200190600090565b80548015611aab577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190611a7c8282611a2f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600560205260409020548015611c0e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611bdf57600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611bdf57818103611b70575b505050611b5c6004611a47565b600052600560205260006040812055600190565b611bc7611b81611b92936004611a2f565b90549060031b1c9283926004611a2f565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526005602052604060002055388080611b4f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b8054906801000000000000000082101561064c5781611b92916001611c3c94018155611a2f565b9055565b80600052600560205260406000205415600014611c7957611c62816004611c15565b600454906000526005602052604060002055600190565b50600090565b80600052600360205260406000205415600014611c7957611ca1816002611c15565b600254906000526003602052604060002055600190565b6000818152600360205260409020548015611c0e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611bdf57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611bdf57808203611d4e575b505050611d3a6002611a47565b600052600360205260006040812055600190565b611d70611d5f611b92936002611a2f565b90549060031b1c9283926002611a2f565b90556000526003602052604060002055388080611d2d56fea164736f6c634300081a000a",
}

var MultiTokenPoolHooksExternalBlocklistABI = MultiTokenPoolHooksExternalBlocklistMetaData.ABI

var MultiTokenPoolHooksExternalBlocklistBin = MultiTokenPoolHooksExternalBlocklistMetaData.Bin

func DeployMultiTokenPoolHooksExternalBlocklist(auth *bind.TransactOpts, backend bind.ContractBackend, allowlist []common.Address, policyEngine common.Address, blocklist common.Address, authorizedCallers []common.Address) (common.Address, *types.Transaction, *MultiTokenPoolHooksExternalBlocklist, error) {
	parsed, err := MultiTokenPoolHooksExternalBlocklistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiTokenPoolHooksExternalBlocklistBin), backend, allowlist, policyEngine, blocklist, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiTokenPoolHooksExternalBlocklist{address: address, abi: *parsed, MultiTokenPoolHooksExternalBlocklistCaller: MultiTokenPoolHooksExternalBlocklistCaller{contract: contract}, MultiTokenPoolHooksExternalBlocklistTransactor: MultiTokenPoolHooksExternalBlocklistTransactor{contract: contract}, MultiTokenPoolHooksExternalBlocklistFilterer: MultiTokenPoolHooksExternalBlocklistFilterer{contract: contract}}, nil
}

type MultiTokenPoolHooksExternalBlocklist struct {
	address common.Address
	abi     abi.ABI
	MultiTokenPoolHooksExternalBlocklistCaller
	MultiTokenPoolHooksExternalBlocklistTransactor
	MultiTokenPoolHooksExternalBlocklistFilterer
}

type MultiTokenPoolHooksExternalBlocklistCaller struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksExternalBlocklistTransactor struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksExternalBlocklistFilterer struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksExternalBlocklistSession struct {
	Contract     *MultiTokenPoolHooksExternalBlocklist
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolHooksExternalBlocklistCallerSession struct {
	Contract *MultiTokenPoolHooksExternalBlocklistCaller
	CallOpts bind.CallOpts
}

type MultiTokenPoolHooksExternalBlocklistTransactorSession struct {
	Contract     *MultiTokenPoolHooksExternalBlocklistTransactor
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolHooksExternalBlocklistRaw struct {
	Contract *MultiTokenPoolHooksExternalBlocklist
}

type MultiTokenPoolHooksExternalBlocklistCallerRaw struct {
	Contract *MultiTokenPoolHooksExternalBlocklistCaller
}

type MultiTokenPoolHooksExternalBlocklistTransactorRaw struct {
	Contract *MultiTokenPoolHooksExternalBlocklistTransactor
}

func NewMultiTokenPoolHooksExternalBlocklist(address common.Address, backend bind.ContractBackend) (*MultiTokenPoolHooksExternalBlocklist, error) {
	abi, err := abi.JSON(strings.NewReader(MultiTokenPoolHooksExternalBlocklistABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMultiTokenPoolHooksExternalBlocklist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklist{address: address, abi: abi, MultiTokenPoolHooksExternalBlocklistCaller: MultiTokenPoolHooksExternalBlocklistCaller{contract: contract}, MultiTokenPoolHooksExternalBlocklistTransactor: MultiTokenPoolHooksExternalBlocklistTransactor{contract: contract}, MultiTokenPoolHooksExternalBlocklistFilterer: MultiTokenPoolHooksExternalBlocklistFilterer{contract: contract}}, nil
}

func NewMultiTokenPoolHooksExternalBlocklistCaller(address common.Address, caller bind.ContractCaller) (*MultiTokenPoolHooksExternalBlocklistCaller, error) {
	contract, err := bindMultiTokenPoolHooksExternalBlocklist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistCaller{contract: contract}, nil
}

func NewMultiTokenPoolHooksExternalBlocklistTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTokenPoolHooksExternalBlocklistTransactor, error) {
	contract, err := bindMultiTokenPoolHooksExternalBlocklist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistTransactor{contract: contract}, nil
}

func NewMultiTokenPoolHooksExternalBlocklistFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiTokenPoolHooksExternalBlocklistFilterer, error) {
	contract, err := bindMultiTokenPoolHooksExternalBlocklist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistFilterer{contract: contract}, nil
}

func bindMultiTokenPoolHooksExternalBlocklist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiTokenPoolHooksExternalBlocklistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.MultiTokenPoolHooksExternalBlocklistCaller.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.MultiTokenPoolHooksExternalBlocklistTransactor.contract.Transfer(opts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.MultiTokenPoolHooksExternalBlocklistTransactor.contract.Transact(opts, method, params...)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.contract.Transfer(opts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.contract.Transact(opts, method, params...)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) CheckAllowList(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "checkAllowList", sender)

	if err != nil {
		return err
	}

	return err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) CheckAllowList(sender common.Address) error {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.CheckAllowList(&_MultiTokenPoolHooksExternalBlocklist.CallOpts, sender)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) CheckAllowList(sender common.Address) error {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.CheckAllowList(&_MultiTokenPoolHooksExternalBlocklist.CallOpts, sender)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllAuthorizedCallers(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllAuthorizedCallers(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) GetAllowList() ([]common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllowList(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) GetAllowList() ([]common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllowList(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) GetAllowListEnabled() (bool, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllowListEnabled(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) GetAllowListEnabled() (bool, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetAllowListEnabled(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) GetBlocklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "getBlocklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) GetBlocklist() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetBlocklist(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) GetBlocklist() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetBlocklist(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) GetPolicyEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "getPolicyEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) GetPolicyEngine() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetPolicyEngine(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) GetPolicyEngine() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.GetPolicyEngine(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) Owner() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.Owner(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) Owner() (common.Address, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.Owner(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MultiTokenPoolHooksExternalBlocklist.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) TypeAndVersion() (string, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.TypeAndVersion(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistCallerSession) TypeAndVersion() (string, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.TypeAndVersion(&_MultiTokenPoolHooksExternalBlocklist.CallOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "acceptOwnership")
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.AcceptOwnership(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.AcceptOwnership(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.ApplyAllowListUpdates(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, removes, adds)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.ApplyAllowListUpdates(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, removes, adds)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.ApplyAuthorizedCallerUpdates(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, authorizedCallerArgs)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.ApplyAuthorizedCallerUpdates(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, authorizedCallerArgs)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) PostflightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "postflightCheck", releaseOrMintIn, localAmount, requestedFinalityConfig)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) PostflightCheck(releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.PostflightCheck(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, releaseOrMintIn, localAmount, requestedFinalityConfig)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) PostflightCheck(releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.PostflightCheck(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, releaseOrMintIn, localAmount, requestedFinalityConfig)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte, amountPostFee *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "preflightCheck", lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte, amountPostFee *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.PreflightCheck(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte, amountPostFee *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.PreflightCheck(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) SetBlocklist(opts *bind.TransactOpts, newBlocklist common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "setBlocklist", newBlocklist)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) SetBlocklist(newBlocklist common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetBlocklist(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newBlocklist)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) SetBlocklist(newBlocklist common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetBlocklist(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newBlocklist)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "setPolicyEngine", newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetPolicyEngine(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetPolicyEngine(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "setPolicyEngineAllowFailedDetach", newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetPolicyEngineAllowFailedDetach(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.SetPolicyEngineAllowFailedDetach(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.contract.Transact(opts, "transferOwnership", to)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.TransferOwnership(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, to)
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooksExternalBlocklist.Contract.TransferOwnership(&_MultiTokenPoolHooksExternalBlocklist.TransactOpts, to)
}

type MultiTokenPoolHooksExternalBlocklistAllowListAddIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistAllowListAdd)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistAllowListAdd)
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

func (it *MultiTokenPoolHooksExternalBlocklistAllowListAddIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAllowListAddIterator, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistAllowListAddIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistAllowListAdd)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseAllowListAdd(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAllowListAdd, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistAllowListAdd)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistAllowListRemove)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistAllowListRemove)
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

func (it *MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistAllowListRemove)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseAllowListRemove(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAllowListRemove, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistAllowListRemove)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded)
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

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseAuthorizedCallerAdded(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved)
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

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistBlocklistSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistBlocklistSet)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistBlocklistSet)
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

func (it *MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistBlocklistSet struct {
	OldBlocklist common.Address
	NewBlocklist common.Address
	Raw          types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterBlocklistSet(opts *bind.FilterOpts, oldBlocklist []common.Address, newBlocklist []common.Address) (*MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator, error) {

	var oldBlocklistRule []interface{}
	for _, oldBlocklistItem := range oldBlocklist {
		oldBlocklistRule = append(oldBlocklistRule, oldBlocklistItem)
	}
	var newBlocklistRule []interface{}
	for _, newBlocklistItem := range newBlocklist {
		newBlocklistRule = append(newBlocklistRule, newBlocklistItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "BlocklistSet", oldBlocklistRule, newBlocklistRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "BlocklistSet", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchBlocklistSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistBlocklistSet, oldBlocklist []common.Address, newBlocklist []common.Address) (event.Subscription, error) {

	var oldBlocklistRule []interface{}
	for _, oldBlocklistItem := range oldBlocklist {
		oldBlocklistRule = append(oldBlocklistRule, oldBlocklistItem)
	}
	var newBlocklistRule []interface{}
	for _, newBlocklistItem := range newBlocklist {
		newBlocklistRule = append(newBlocklistRule, newBlocklistItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "BlocklistSet", oldBlocklistRule, newBlocklistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistBlocklistSet)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "BlocklistSet", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseBlocklistSet(log types.Log) (*MultiTokenPoolHooksExternalBlocklistBlocklistSet, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistBlocklistSet)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "BlocklistSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested)
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

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferred)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferred)
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

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferred)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferred, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistOwnershipTransferred)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached)
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

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached struct {
	PolicyEngine common.Address
	Raw          types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterPolicyEngineAttached(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "PolicyEngineAttached", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "PolicyEngineAttached", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchPolicyEngineAttached(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached, policyEngine []common.Address) (event.Subscription, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "PolicyEngineAttached", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "PolicyEngineAttached", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParsePolicyEngineAttached(log types.Log) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "PolicyEngineAttached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator struct {
	Event *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed)
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
		it.Event = new(MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed)
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

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed struct {
	PolicyEngine common.Address
	Raw          types.Log
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) FilterPolicyEngineDetachFailed(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.FilterLogs(opts, "PolicyEngineDetachFailed", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator{contract: _MultiTokenPoolHooksExternalBlocklist.contract, event: "PolicyEngineDetachFailed", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) WatchPolicyEngineDetachFailed(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed, policyEngine []common.Address) (event.Subscription, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooksExternalBlocklist.contract.WatchLogs(opts, "PolicyEngineDetachFailed", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed)
				if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "PolicyEngineDetachFailed", log); err != nil {
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

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklistFilterer) ParsePolicyEngineDetachFailed(log types.Log) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed, error) {
	event := new(MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed)
	if err := _MultiTokenPoolHooksExternalBlocklist.contract.UnpackLog(event, "PolicyEngineDetachFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (MultiTokenPoolHooksExternalBlocklistAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MultiTokenPoolHooksExternalBlocklistAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (MultiTokenPoolHooksExternalBlocklistBlocklistSet) Topic() common.Hash {
	return common.HexToHash("0x7e053cdc9069fe4f629b6b3fa2a01bb53a9a8305ec11830b259d95e9e75b7304")
}

func (MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MultiTokenPoolHooksExternalBlocklistOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached) Topic() common.Hash {
	return common.HexToHash("0x57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac")
}

func (MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed) Topic() common.Hash {
	return common.HexToHash("0xdbc729cc8d59cb8abc2a107be310a639f87b7673410abc028f2a3db03ff7b5b2")
}

func (_MultiTokenPoolHooksExternalBlocklist *MultiTokenPoolHooksExternalBlocklist) Address() common.Address {
	return _MultiTokenPoolHooksExternalBlocklist.address
}

type MultiTokenPoolHooksExternalBlocklistInterface interface {
	CheckAllowList(opts *bind.CallOpts, sender common.Address) error

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetBlocklist(opts *bind.CallOpts) (common.Address, error)

	GetPolicyEngine(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	PostflightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte, amountPostFee *big.Int) (*types.Transaction, error)

	SetBlocklist(opts *bind.TransactOpts, newBlocklist common.Address) (*types.Transaction, error)

	SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*MultiTokenPoolHooksExternalBlocklistAuthorizedCallerRemoved, error)

	FilterBlocklistSet(opts *bind.FilterOpts, oldBlocklist []common.Address, newBlocklist []common.Address) (*MultiTokenPoolHooksExternalBlocklistBlocklistSetIterator, error)

	WatchBlocklistSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistBlocklistSet, oldBlocklist []common.Address, newBlocklist []common.Address) (event.Subscription, error)

	ParseBlocklistSet(log types.Log) (*MultiTokenPoolHooksExternalBlocklistBlocklistSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolHooksExternalBlocklistOwnershipTransferred, error)

	FilterPolicyEngineAttached(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineAttachedIterator, error)

	WatchPolicyEngineAttached(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached, policyEngine []common.Address) (event.Subscription, error)

	ParsePolicyEngineAttached(log types.Log) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineAttached, error)

	FilterPolicyEngineDetachFailed(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailedIterator, error)

	WatchPolicyEngineDetachFailed(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed, policyEngine []common.Address) (event.Subscription, error)

	ParsePolicyEngineDetachFailed(log types.Log) (*MultiTokenPoolHooksExternalBlocklistPolicyEngineDetachFailed, error)

	Address() common.Address
}
