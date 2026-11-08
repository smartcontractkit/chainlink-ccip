// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multi_token_pool_hooks

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

var MultiTokenPoolHooksMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"},{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkAllowList\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postflightCheck\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preflightCheck\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngine\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPolicyEngineAllowFailedDetach\",\"inputs\":[{\"name\":\"newPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineAttached\",\"inputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IPolicyEngine\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyEngineDetachFailed\",\"inputs\":[{\"name\":\"policyEngine\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contract IPolicyEngine\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyEngineDetachReverted\",\"inputs\":[{\"name\":\"oldPolicyEngine\",\"type\":\"address\",\"internalType\":\"contract IPolicyEngine\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461033c57612040803803809161001c8285610341565b8339810160608282031261033c5781516001600160401b03811161033c5781610046918401610364565b6020830151909290916001600160a01b038316830361033c5760408201516001600160401b03811161033c5761007c9201610364565b90331561032b57600180546001600160a01b0319163317905560405160209390916100a78584610341565b60008352600036813760408051949085016001600160401b03811186821017610315576040528452828585015260005b835181101561013e576001906001600160a01b036100f582876103db565b511687610101826106db565b61010e575b5050016100d7565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13887610106565b50915091519060005b82518110156101b8576001600160a01b0361016282856103db565b51169081156101a7577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8683610199600195610663565b50604051908152a101610147565b6342bcdf7f60e11b60005260046000fd5b50828481511515806080526101f7575b6101d183610405565b6040516118d0908161077082396080518181816101ec015281816108e601526111230152f35b6040516102048282610341565b600081526000368137608051156103045760005b815181101561027f576001906001600160a01b0361023682856103db565b51168461024282610570565b61024f575b505001610218565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a18684610247565b505060005b82518110156102f8576001906001600160a01b036102a282866103db565b511680156102f257836102b4826106a2565b6102c2575b50505b01610284565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a185836102b9565b506102bc565b5050506101d1826101c8565b6335f4a7b360e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761031557604052565b81601f8201121561033c578051916001600160401b038311610315578260051b9160208301936103976040519586610341565b845260208085019382010191821161033c57602001915b8183106103bb5750505090565b82516001600160a01b038116810361033c578152602092830192016103ae565b80518210156103ef5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6006546001600160a01b03918216916000911682811461051457806104b5575b50600680546001600160a01b0319168317905581610464575b807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b156104a757604051631100482d60e01b8152818160048183875af180156104aa57610492575b5061043e565b61049d828092610341565b6104a7573861048c565b80fd5b6040513d84823e3d90fd5b9091813b1561051057604051628950d760e61b8152838160048183875af190816104fc575b506104f35763e047565560e01b83526004829052602483fd5b91905038610425565b8461050991959295610341565b92386104da565b8280fd5b505050565b80548210156103ef5760005260206000200190600090565b8054801561055a5760001901906105488282610519565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b600081815260056020526040902054801561063157600019810181811161061b5760045460001981019190821161061b578181036105ca575b5050506105b66004610531565b600052600560205260006040812055600190565b6106036105db6105ec936004610519565b90549060031b1c9283926004610519565b819391549060031b91821b91600019901b19161790565b905560005260056020526040600020553880806105a9565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8054906801000000000000000082101561031557816105ec91600161065f94018155610519565b9055565b8060005260036020526040600020541560001461069c57610685816002610638565b600254906000526003602052604060002055600190565b50600090565b8060005260056020526040600020541560001461069c576106c4816004610638565b600454906000526005602052604060002055600190565b600081815260036020526040902054801561063157600019810181811161061b5760025460001981019190821161061b57808203610735575b5050506107216002610531565b600052600360205260006040812055600190565b6107576107466105ec936002610519565b90549060031b1c9283926002610519565b9055600052600360205260406000205538808061071456fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714610b235780632451a62714610a8e5780634ef34bc014610a2e57806354c8a4f3146108545780636135b085146107f45780636371157414610771578063683173121461071f57806379ba5097146106365780638da5cb5b146105e457806391a2749a146103bc578063a7cd63b714610317578063a8027c0f14610252578063ce07c7c814610211578063e0351e13146101b65763f2fde38b146100c157600080fd5b346101b15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15773ffffffffffffffffffffffffffffffffffffffff61010d610d37565b6101156111bf565b1633811461018757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346101b15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15761025061024b610d37565b611121565b005b346101b15760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043567ffffffffffffffff81116101b15760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b1576102c7610cab565b506044359067ffffffffffffffff82116101b157366023830112156101b15781600401359067ffffffffffffffff82116101b15736602483850101116101b15760246102509301906004016110ac565b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760405180602060045491828152019060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9060005b8181106103a6576103a28561039681870382610cf6565b60405191829182610bfb565b0390f35b825484526020909301926001928301920161037f565b346101b15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043567ffffffffffffffff81116101b15760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b157604051906040820182811067ffffffffffffffff8211176105b557604052806004013567ffffffffffffffff81116101b15761046b9060043691840101610dcc565b825260248101359067ffffffffffffffff82116101b15760046104919236920101610dcc565b602082019081526104a06111bf565b519060005b8251811015610518578073ffffffffffffffffffffffffffffffffffffffff6104d0600193866114e3565b51166104db816117f3565b6104e7575b50016104a5565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1846104e0565b505160005b81518110156102505773ffffffffffffffffffffffffffffffffffffffff61054582846114e3565b511690811561058b577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361057d6001956117ba565b50604051908152a10161051d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b157602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760005473ffffffffffffffffffffffffffffffffffffffff811633036106f5577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b157602073ffffffffffffffffffffffffffffffffffffffff60065416604051908152f35b346101b15760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043567ffffffffffffffff81116101b1576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b157610250906107eb610c7c565b50600401610f33565b346101b15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043573ffffffffffffffffffffffffffffffffffffffff811681036101b1576102509061084f6111bf565b611379565b346101b15760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043567ffffffffffffffff81116101b1576108a3903690600401610c4b565b6024359067ffffffffffffffff82116101b1576108dc6108ca6108e4933690600401610c4b565b9490926108d56111bf565b3691610d5a565b923691610d5a565b7f000000000000000000000000000000000000000000000000000000000000000015610a045760005b8251811015610980578073ffffffffffffffffffffffffffffffffffffffff610938600193866114e3565b511661094381611615565b61094f575b500161090d565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184610948565b5060005b8151811015610250578073ffffffffffffffffffffffffffffffffffffffff6109af600193856114e3565b511680156109fe576109c08161177b565b6109cd575b505b01610984565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1836109c5565b506109c7565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b15760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760043573ffffffffffffffffffffffffffffffffffffffff811681036101b15761025090610a896111bf565b61120a565b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b15760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610b0d576103a28561039681870382610cf6565b8254845260209093019260019283019201610af6565b346101b15760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b1576103a26040805190610b648183610cf6565b601d82527f4d756c7469546f6b656e506f6f6c486f6f6b7320322e312e302d6465760000006020830152519182916020835260208301905b919082519283825260005b848110610be65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610ba7565b602060408183019282815284518094520192019060005b818110610c1f5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610c12565b9181601f840112156101b15782359167ffffffffffffffff83116101b1576020808501948460051b0101116101b157565b604435907fffffffff00000000000000000000000000000000000000000000000000000000821682036101b157565b602435907fffffffff00000000000000000000000000000000000000000000000000000000821682036101b157565b6080810190811067ffffffffffffffff8211176105b557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176105b557604052565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101b157565b90929167ffffffffffffffff84116105b5578360051b916020604051610d8282860182610cf6565b80968152019281019182116101b157915b818310610d9f57505050565b823573ffffffffffffffffffffffffffffffffffffffff811681036101b157815260209283019201610d93565b9080601f830112156101b157816020610de793359101610d5a565b90565b67ffffffffffffffff81116105b557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919091610e3081610dea565b610e3d6040519182610cf6565b809382825282600401116101b15781816000936004602080950137010152565b929192610e6982610dea565b91610e776040519384610cf6565b8294818452818301116101b1578281602093846000960137010152565b90610de791602081527fffffffff00000000000000000000000000000000000000000000000000000000825116602082015273ffffffffffffffffffffffffffffffffffffffff60208301511660408201526060610f00604084015160808385015260a0840190610b9c565b9201519060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610b9c565b610f3b611526565b6006549073ffffffffffffffffffffffffffffffffffffffff600092169081156110a7573660041161109f5760e0810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156110a3570180359067ffffffffffffffff82116110a35760200181360381136110a35761102d9060405192610fc984610cda565b7fffffffff000000000000000000000000000000000000000000000000000000008635168452336020850152611021367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601610e24565b60408501523691610e5d565b6060820152813b1561109f57611075839283926040519485809481937fc2098e0800000000000000000000000000000000000000000000000000000000835260048301610e94565b03925af1801561109457611087575050565b8161109191610cf6565b50565b6040513d84823e3d90fd5b8280fd5b8380fd5b505050565b916110b5611526565b6040600093013573ffffffffffffffffffffffffffffffffffffffff811681036110a3576110e290611121565b73ffffffffffffffffffffffffffffffffffffffff6006541691821561111b57366004116110a35761102d9060405192610fc984610cda565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000080611193575b61114f5750565b73ffffffffffffffffffffffffffffffffffffffff907fd0d25976000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff8116600052600560205260406000205415611148565b73ffffffffffffffffffffffffffffffffffffffff6001541633036111e057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60065473ffffffffffffffffffffffffffffffffffffffff80600092169216918083146110a757806112f0575b50817fffffffffffffffffffffffff000000000000000000000000000000000000000060065416176006558161128e575b807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b156112ed576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183875af18015611094576112d5575b50611268565b816112df91610cf6565b806000126112ed57386112cf565b80fd5b90813b156112ed57604051907f225435c0000000000000000000000000000000000000000000000000000000008252808260048183875af19182611369575b509061136257907fdbc729cc8d59cb8abc2a107be310a639f87b7673410abc028f2a3db03ff7b5b2600080a25b38611237565b905061135c565b8161137391610cf6565b3861132f565b600073ffffffffffffffffffffffffffffffffffffffff80600654169216918083146110a75780611456575b50817fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655816113fb57807f57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac91a2565b813b156112ed576040517f1100482d000000000000000000000000000000000000000000000000000000008152818160048183875af18015611094576114415750611268565b61144c828092610cf6565b6112ed57386112cf565b9091813b1561109f576040517f225435c0000000000000000000000000000000000000000000000000000000008152838160048183875af190816114cf575b506114c657602483837fe0475655000000000000000000000000000000000000000000000000000000008252600452fd5b919050386113a5565b846114dc91959295610cf6565b9238611495565b80518210156114f75760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3360005260036020526040600020541561153c57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b80548210156114f75760005260206000200190600090565b805480156115e6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906115b7828261156a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600560205260409020548015611749577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161171a57600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161171a578181036116ab575b5050506116976004611582565b600052600560205260006040812055600190565b6117026116bc6116cd93600461156a565b90549060031b1c928392600461156a565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600560205260406000205538808061168a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b805490680100000000000000008210156105b557816116cd9160016117779401815561156a565b9055565b806000526005602052604060002054156000146117b45761179d816004611750565b600454906000526005602052604060002055600190565b50600090565b806000526003602052604060002054156000146117b4576117dc816002611750565b600254906000526003602052604060002055600190565b6000818152600360205260409020548015611749577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161171a57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161171a57808203611889575b5050506118756002611582565b600052600360205260006040812055600190565b6118ab61189a6116cd93600261156a565b90549060031b1c928392600261156a565b9055600052600360205260406000205538808061186856fea164736f6c634300081a000a",
}

var MultiTokenPoolHooksABI = MultiTokenPoolHooksMetaData.ABI

var MultiTokenPoolHooksBin = MultiTokenPoolHooksMetaData.Bin

func DeployMultiTokenPoolHooks(auth *bind.TransactOpts, backend bind.ContractBackend, allowlist []common.Address, policyEngine common.Address, authorizedCallers []common.Address) (common.Address, *types.Transaction, *MultiTokenPoolHooks, error) {
	parsed, err := MultiTokenPoolHooksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiTokenPoolHooksBin), backend, allowlist, policyEngine, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiTokenPoolHooks{address: address, abi: *parsed, MultiTokenPoolHooksCaller: MultiTokenPoolHooksCaller{contract: contract}, MultiTokenPoolHooksTransactor: MultiTokenPoolHooksTransactor{contract: contract}, MultiTokenPoolHooksFilterer: MultiTokenPoolHooksFilterer{contract: contract}}, nil
}

type MultiTokenPoolHooks struct {
	address common.Address
	abi     abi.ABI
	MultiTokenPoolHooksCaller
	MultiTokenPoolHooksTransactor
	MultiTokenPoolHooksFilterer
}

type MultiTokenPoolHooksCaller struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksTransactor struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksFilterer struct {
	contract *bind.BoundContract
}

type MultiTokenPoolHooksSession struct {
	Contract     *MultiTokenPoolHooks
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolHooksCallerSession struct {
	Contract *MultiTokenPoolHooksCaller
	CallOpts bind.CallOpts
}

type MultiTokenPoolHooksTransactorSession struct {
	Contract     *MultiTokenPoolHooksTransactor
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolHooksRaw struct {
	Contract *MultiTokenPoolHooks
}

type MultiTokenPoolHooksCallerRaw struct {
	Contract *MultiTokenPoolHooksCaller
}

type MultiTokenPoolHooksTransactorRaw struct {
	Contract *MultiTokenPoolHooksTransactor
}

func NewMultiTokenPoolHooks(address common.Address, backend bind.ContractBackend) (*MultiTokenPoolHooks, error) {
	abi, err := abi.JSON(strings.NewReader(MultiTokenPoolHooksABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMultiTokenPoolHooks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooks{address: address, abi: abi, MultiTokenPoolHooksCaller: MultiTokenPoolHooksCaller{contract: contract}, MultiTokenPoolHooksTransactor: MultiTokenPoolHooksTransactor{contract: contract}, MultiTokenPoolHooksFilterer: MultiTokenPoolHooksFilterer{contract: contract}}, nil
}

func NewMultiTokenPoolHooksCaller(address common.Address, caller bind.ContractCaller) (*MultiTokenPoolHooksCaller, error) {
	contract, err := bindMultiTokenPoolHooks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksCaller{contract: contract}, nil
}

func NewMultiTokenPoolHooksTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTokenPoolHooksTransactor, error) {
	contract, err := bindMultiTokenPoolHooks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksTransactor{contract: contract}, nil
}

func NewMultiTokenPoolHooksFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiTokenPoolHooksFilterer, error) {
	contract, err := bindMultiTokenPoolHooks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksFilterer{contract: contract}, nil
}

func bindMultiTokenPoolHooks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiTokenPoolHooksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPoolHooks.Contract.MultiTokenPoolHooksCaller.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.MultiTokenPoolHooksTransactor.contract.Transfer(opts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.MultiTokenPoolHooksTransactor.contract.Transact(opts, method, params...)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPoolHooks.Contract.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.contract.Transfer(opts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.contract.Transact(opts, method, params...)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) CheckAllowList(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "checkAllowList", sender)

	if err != nil {
		return err
	}

	return err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) CheckAllowList(sender common.Address) error {
	return _MultiTokenPoolHooks.Contract.CheckAllowList(&_MultiTokenPoolHooks.CallOpts, sender)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) CheckAllowList(sender common.Address) error {
	return _MultiTokenPoolHooks.Contract.CheckAllowList(&_MultiTokenPoolHooks.CallOpts, sender)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetAllAuthorizedCallers(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetAllAuthorizedCallers(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) GetAllowList() ([]common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetAllowList(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) GetAllowList() ([]common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetAllowList(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) GetAllowListEnabled() (bool, error) {
	return _MultiTokenPoolHooks.Contract.GetAllowListEnabled(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) GetAllowListEnabled() (bool, error) {
	return _MultiTokenPoolHooks.Contract.GetAllowListEnabled(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) GetPolicyEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "getPolicyEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) GetPolicyEngine() (common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetPolicyEngine(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) GetPolicyEngine() (common.Address, error) {
	return _MultiTokenPoolHooks.Contract.GetPolicyEngine(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) Owner() (common.Address, error) {
	return _MultiTokenPoolHooks.Contract.Owner(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) Owner() (common.Address, error) {
	return _MultiTokenPoolHooks.Contract.Owner(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MultiTokenPoolHooks.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) TypeAndVersion() (string, error) {
	return _MultiTokenPoolHooks.Contract.TypeAndVersion(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksCallerSession) TypeAndVersion() (string, error) {
	return _MultiTokenPoolHooks.Contract.TypeAndVersion(&_MultiTokenPoolHooks.CallOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "acceptOwnership")
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.AcceptOwnership(&_MultiTokenPoolHooks.TransactOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.AcceptOwnership(&_MultiTokenPoolHooks.TransactOpts)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.ApplyAllowListUpdates(&_MultiTokenPoolHooks.TransactOpts, removes, adds)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.ApplyAllowListUpdates(&_MultiTokenPoolHooks.TransactOpts, removes, adds)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.ApplyAuthorizedCallerUpdates(&_MultiTokenPoolHooks.TransactOpts, authorizedCallerArgs)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.ApplyAuthorizedCallerUpdates(&_MultiTokenPoolHooks.TransactOpts, authorizedCallerArgs)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) PostflightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "postflightCheck", releaseOrMintIn, arg1, arg2)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) PostflightCheck(releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.PostflightCheck(&_MultiTokenPoolHooks.TransactOpts, releaseOrMintIn, arg1, arg2)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) PostflightCheck(releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 [4]byte) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.PostflightCheck(&_MultiTokenPoolHooks.TransactOpts, releaseOrMintIn, arg1, arg2)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 [4]byte, tokenArgs []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "preflightCheck", lockOrBurnIn, arg1, tokenArgs, arg3)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 [4]byte, tokenArgs []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.PreflightCheck(&_MultiTokenPoolHooks.TransactOpts, lockOrBurnIn, arg1, tokenArgs, arg3)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 [4]byte, tokenArgs []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.PreflightCheck(&_MultiTokenPoolHooks.TransactOpts, lockOrBurnIn, arg1, tokenArgs, arg3)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "setPolicyEngine", newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.SetPolicyEngine(&_MultiTokenPoolHooks.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) SetPolicyEngine(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.SetPolicyEngine(&_MultiTokenPoolHooks.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "setPolicyEngineAllowFailedDetach", newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.SetPolicyEngineAllowFailedDetach(&_MultiTokenPoolHooks.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) SetPolicyEngineAllowFailedDetach(newPolicyEngine common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.SetPolicyEngineAllowFailedDetach(&_MultiTokenPoolHooks.TransactOpts, newPolicyEngine)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.contract.Transact(opts, "transferOwnership", to)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.TransferOwnership(&_MultiTokenPoolHooks.TransactOpts, to)
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPoolHooks.Contract.TransferOwnership(&_MultiTokenPoolHooks.TransactOpts, to)
}

type MultiTokenPoolHooksAllowListAddIterator struct {
	Event *MultiTokenPoolHooksAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksAllowListAdd)
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
		it.Event = new(MultiTokenPoolHooksAllowListAdd)
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

func (it *MultiTokenPoolHooksAllowListAddIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*MultiTokenPoolHooksAllowListAddIterator, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksAllowListAddIterator{contract: _MultiTokenPoolHooks.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksAllowListAdd)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseAllowListAdd(log types.Log) (*MultiTokenPoolHooksAllowListAdd, error) {
	event := new(MultiTokenPoolHooksAllowListAdd)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksAllowListRemoveIterator struct {
	Event *MultiTokenPoolHooksAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksAllowListRemove)
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
		it.Event = new(MultiTokenPoolHooksAllowListRemove)
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

func (it *MultiTokenPoolHooksAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*MultiTokenPoolHooksAllowListRemoveIterator, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksAllowListRemoveIterator{contract: _MultiTokenPoolHooks.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksAllowListRemove)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseAllowListRemove(log types.Log) (*MultiTokenPoolHooksAllowListRemove, error) {
	event := new(MultiTokenPoolHooksAllowListRemove)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksAuthorizedCallerAddedIterator struct {
	Event *MultiTokenPoolHooksAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksAuthorizedCallerAdded)
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
		it.Event = new(MultiTokenPoolHooksAuthorizedCallerAdded)
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

func (it *MultiTokenPoolHooksAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*MultiTokenPoolHooksAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksAuthorizedCallerAddedIterator{contract: _MultiTokenPoolHooks.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksAuthorizedCallerAdded)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseAuthorizedCallerAdded(log types.Log) (*MultiTokenPoolHooksAuthorizedCallerAdded, error) {
	event := new(MultiTokenPoolHooksAuthorizedCallerAdded)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksAuthorizedCallerRemovedIterator struct {
	Event *MultiTokenPoolHooksAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksAuthorizedCallerRemoved)
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
		it.Event = new(MultiTokenPoolHooksAuthorizedCallerRemoved)
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

func (it *MultiTokenPoolHooksAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*MultiTokenPoolHooksAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksAuthorizedCallerRemovedIterator{contract: _MultiTokenPoolHooks.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksAuthorizedCallerRemoved)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*MultiTokenPoolHooksAuthorizedCallerRemoved, error) {
	event := new(MultiTokenPoolHooksAuthorizedCallerRemoved)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksOwnershipTransferRequestedIterator struct {
	Event *MultiTokenPoolHooksOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksOwnershipTransferRequested)
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
		it.Event = new(MultiTokenPoolHooksOwnershipTransferRequested)
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

func (it *MultiTokenPoolHooksOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksOwnershipTransferRequestedIterator{contract: _MultiTokenPoolHooks.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksOwnershipTransferRequested)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolHooksOwnershipTransferRequested, error) {
	event := new(MultiTokenPoolHooksOwnershipTransferRequested)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksOwnershipTransferredIterator struct {
	Event *MultiTokenPoolHooksOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksOwnershipTransferred)
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
		it.Event = new(MultiTokenPoolHooksOwnershipTransferred)
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

func (it *MultiTokenPoolHooksOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksOwnershipTransferredIterator{contract: _MultiTokenPoolHooks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksOwnershipTransferred)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolHooksOwnershipTransferred, error) {
	event := new(MultiTokenPoolHooksOwnershipTransferred)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksPolicyEngineAttachedIterator struct {
	Event *MultiTokenPoolHooksPolicyEngineAttached

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksPolicyEngineAttachedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksPolicyEngineAttached)
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
		it.Event = new(MultiTokenPoolHooksPolicyEngineAttached)
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

func (it *MultiTokenPoolHooksPolicyEngineAttachedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksPolicyEngineAttachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksPolicyEngineAttached struct {
	PolicyEngine common.Address
	Raw          types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterPolicyEngineAttached(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksPolicyEngineAttachedIterator, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "PolicyEngineAttached", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksPolicyEngineAttachedIterator{contract: _MultiTokenPoolHooks.contract, event: "PolicyEngineAttached", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchPolicyEngineAttached(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksPolicyEngineAttached, policyEngine []common.Address) (event.Subscription, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "PolicyEngineAttached", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksPolicyEngineAttached)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "PolicyEngineAttached", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParsePolicyEngineAttached(log types.Log) (*MultiTokenPoolHooksPolicyEngineAttached, error) {
	event := new(MultiTokenPoolHooksPolicyEngineAttached)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "PolicyEngineAttached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolHooksPolicyEngineDetachFailedIterator struct {
	Event *MultiTokenPoolHooksPolicyEngineDetachFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolHooksPolicyEngineDetachFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolHooksPolicyEngineDetachFailed)
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
		it.Event = new(MultiTokenPoolHooksPolicyEngineDetachFailed)
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

func (it *MultiTokenPoolHooksPolicyEngineDetachFailedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolHooksPolicyEngineDetachFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolHooksPolicyEngineDetachFailed struct {
	PolicyEngine common.Address
	Raw          types.Log
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) FilterPolicyEngineDetachFailed(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksPolicyEngineDetachFailedIterator, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.FilterLogs(opts, "PolicyEngineDetachFailed", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolHooksPolicyEngineDetachFailedIterator{contract: _MultiTokenPoolHooks.contract, event: "PolicyEngineDetachFailed", logs: logs, sub: sub}, nil
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) WatchPolicyEngineDetachFailed(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksPolicyEngineDetachFailed, policyEngine []common.Address) (event.Subscription, error) {

	var policyEngineRule []interface{}
	for _, policyEngineItem := range policyEngine {
		policyEngineRule = append(policyEngineRule, policyEngineItem)
	}

	logs, sub, err := _MultiTokenPoolHooks.contract.WatchLogs(opts, "PolicyEngineDetachFailed", policyEngineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolHooksPolicyEngineDetachFailed)
				if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "PolicyEngineDetachFailed", log); err != nil {
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

func (_MultiTokenPoolHooks *MultiTokenPoolHooksFilterer) ParsePolicyEngineDetachFailed(log types.Log) (*MultiTokenPoolHooksPolicyEngineDetachFailed, error) {
	event := new(MultiTokenPoolHooksPolicyEngineDetachFailed)
	if err := _MultiTokenPoolHooks.contract.UnpackLog(event, "PolicyEngineDetachFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (MultiTokenPoolHooksAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MultiTokenPoolHooksAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MultiTokenPoolHooksAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (MultiTokenPoolHooksAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (MultiTokenPoolHooksOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MultiTokenPoolHooksOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MultiTokenPoolHooksPolicyEngineAttached) Topic() common.Hash {
	return common.HexToHash("0x57d241970863a27bedbf58b705b45a0b267f76f9a3a7fd432e217a37e4173fac")
}

func (MultiTokenPoolHooksPolicyEngineDetachFailed) Topic() common.Hash {
	return common.HexToHash("0xdbc729cc8d59cb8abc2a107be310a639f87b7673410abc028f2a3db03ff7b5b2")
}

func (_MultiTokenPoolHooks *MultiTokenPoolHooks) Address() common.Address {
	return _MultiTokenPoolHooks.address
}

type MultiTokenPoolHooksInterface interface {
	CheckAllowList(opts *bind.CallOpts, sender common.Address) error

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetPolicyEngine(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	PostflightCheck(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, arg1 *big.Int, arg2 [4]byte) (*types.Transaction, error)

	PreflightCheck(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 [4]byte, tokenArgs []byte, arg3 *big.Int) (*types.Transaction, error)

	SetPolicyEngine(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	SetPolicyEngineAllowFailedDetach(opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*MultiTokenPoolHooksAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*MultiTokenPoolHooksAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*MultiTokenPoolHooksAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*MultiTokenPoolHooksAllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*MultiTokenPoolHooksAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*MultiTokenPoolHooksAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*MultiTokenPoolHooksAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*MultiTokenPoolHooksAuthorizedCallerRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolHooksOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolHooksOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolHooksOwnershipTransferred, error)

	FilterPolicyEngineAttached(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksPolicyEngineAttachedIterator, error)

	WatchPolicyEngineAttached(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksPolicyEngineAttached, policyEngine []common.Address) (event.Subscription, error)

	ParsePolicyEngineAttached(log types.Log) (*MultiTokenPoolHooksPolicyEngineAttached, error)

	FilterPolicyEngineDetachFailed(opts *bind.FilterOpts, policyEngine []common.Address) (*MultiTokenPoolHooksPolicyEngineDetachFailedIterator, error)

	WatchPolicyEngineDetachFailed(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolHooksPolicyEngineDetachFailed, policyEngine []common.Address) (event.Subscription, error)

	ParsePolicyEngineDetachFailed(log types.Log) (*MultiTokenPoolHooksPolicyEngineDetachFailed, error)

	Address() common.Address
}
