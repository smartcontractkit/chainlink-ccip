// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_token_messenger_proxy

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

var CCTPTokenMessengerProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"usdcToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithCaller\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithHook\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUSDCToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMessageTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageBodyVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60c080604052346103b4576119d9803803809161001c828561045c565b83398101906060818303126103b45780516001600160a01b0381168082036103b45760208301516001600160a01b038116939092908484036103b4576040810151906001600160401b0382116103b4570185601f820112156103b4578051956001600160401b0387116103b9578660051b9160208301976100a0604051998a61045c565b88526020808901938201019182116103b457602001915b81831061043c57505050331561042b57600180546001600160a01b0319163317905560405160209390926100eb858561045c565b60008452600036813760408051979088016001600160401b038111898210176103b9576040528752838588015260005b8451811015610182576001906001600160a01b03610139828861047f565b511687610145826104c1565b610152575b50500161011b565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388761014a565b508493508587519260005b84518110156101fe576001600160a01b036101a8828761047f565b51169081156101ed577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef88836101df6001956105a9565b50604051908152a10161018d565b6342bcdf7f60e11b60005260046000fd5b509085918560805260a052604051636eb1769f60e11b81523060048201528360248201528281604481855afa90811561041f576000916103f2575b5060001981018091116103dc57604051908382019463095ea7b360e01b8652602483015260448201526044815261027160648261045c565b600080604095865193610284888661045c565b8685527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656487860152519082865af13d156103cf573d906001600160401b0382116103b95785516102f19490926102e3601f8201601f191688018561045c565b83523d60008785013e610609565b80518061033c575b83516112ff90816106da823960805181818161017a015281816103c201528181610459015281816105d4015281816107750152610bce015260a051816103530152f35b818391810103126103b4578101518015908115036103b45761035f5780806102f9565b608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916102f192606091610609565b634e487b7160e01b600052601160045260246000fd5b90508281813d8311610418575b610409818361045c565b810103126103b4575184610239565b503d6103ff565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b03811681036103b4578152602092830192016100b7565b601f909101601f19168101906001600160401b038211908210176103b957604052565b80518210156104935760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104935760005260206000200190600090565b60008181526003602052604090205480156105a25760001981018181116103dc576002546000198101919082116103dc57808203610551575b505050600254801561053b57600019016105158160026104a9565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61058a6105626105739360026104a9565b90549060031b1c92839260026104a9565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806104fa565b5050600090565b8060005260036020526040600020541560001461060357600254680100000000000000008110156103b9576105ea61057382600185940160025560026104a9565b9055600254906000526003602052604060002055600190565b50600090565b9192901561066b575081511561061d575090565b3b156106265790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561067e5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106106c15750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161069f56fe608080604052600436101561001357600080fd5b600090813560e01c908163181f5a7714610d5c575080632451a62714610c725780632c12192114610b5c57806379ba509714610a775780638da5cb5b14610a2557806391a2749a146107f95780639cdbb18114610703578063abbce43914610523578063d04857b0146103e6578063e6236af814610377578063f175240914610308578063f2fde38b146102185763f856ddb6146100b057600080fd5b346102155760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610215576100e7610faa565b73ffffffffffffffffffffffffffffffffffffffff610104610efe565b61010c611071565b63ffffffff604051937ff856ddb600000000000000000000000000000000000000000000000000000000855260043560048601521660248401526044356044840152166064820152608435608482015260208160a4818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561020a5782916101bf575b60208267ffffffffffffffff60405191168152f35b90506020813d602011610202575b816101da60209383610e8e565b810103126101fe575167ffffffffffffffff811681036101fe5760209150386101aa565b5080fd5b3d91506101cd565b6040513d84823e3d90fd5b80fd5b50346102155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102155760043573ffffffffffffffffffffffffffffffffffffffff81168091036101fe57610271610fe3565b3381146102e057807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102155760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610215578061041f610faa565b610427610efe565b90610430610fbd565b90610439610fd0565b91610442611071565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b1561051f578563ffffffff938482968160e49673ffffffffffffffffffffffffffffffffffffffff6040519b8c9a8b997fd04857b0000000000000000000000000000000000000000000000000000000008b5260043560048c01521660248a015260443560448a015216606488015260843560848801521660a48601521660c48401525af1801561020a5761050e5750f35b8161051891610e8e565b6102155780f35b8580fd5b5034610215576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102155761055c610faa565b81610565610efe565b61056d610fbd565b610575610fd0565b9360e4359267ffffffffffffffff84116106ff57366023850112156106ff5783600401359367ffffffffffffffff851161051f57366024868301011161051f576105bd611071565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b156106fb5785879687936040519a8b98899788967fabbce439000000000000000000000000000000000000000000000000000000008852600435600489015263ffffffff166024880152604435604488015273ffffffffffffffffffffffffffffffffffffffff166064870152608435608487015263ffffffff1660a486015263ffffffff1660c485015260e48401610100905281610104850152602401610124840137828183016101240152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01681010361012401925af180156106ee576106e05780f35b6106e991610e8e565b388180f35b50604051903d90823e3d90fd5b8680fd5b8480fd5b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610215576040517f9cdbb18100000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561020a5782916107b6575b60208263ffffffff60405191168152f35b90506020813d6020116107f1575b816107d160209383610e8e565b810103126101fe575163ffffffff811681036101fe5760209150386107a5565b3d91506107c4565b50346102155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102155760043567ffffffffffffffff81116101fe5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101fe57604051906040820182811067ffffffffffffffff8211176109f857604052806004013567ffffffffffffffff81116109f4576108a99060043691840101610f26565b825260248101359067ffffffffffffffff82116109f45760046108cf9236920101610f26565b602082019081526108de610fe3565b5191805b8351811015610955578073ffffffffffffffffffffffffffffffffffffffff61090d6001938761102e565b5116610918816110cd565b610924575b50016108e2565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13861091d565b509051815b81518110156109f05773ffffffffffffffffffffffffffffffffffffffff610982828461102e565b511680156109c857907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836109ba600195611292565b50604051908152a10161095a565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b8380fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021557805473ffffffffffffffffffffffffffffffffffffffff81163303610b34577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610215576040517f2c12192100000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561020a578291610c1f575b60208273ffffffffffffffffffffffffffffffffffffffff60405191168152f35b90506020813d602011610c6a575b81610c3a60209383610e8e565b810103126101fe575173ffffffffffffffffffffffffffffffffffffffff811681036101fe576020915038610bfe565b3d9150610c2d565b503461021557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102155760405180602060025491828152018091600285527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90855b818110610d465750505082610cef910383610e8e565b604051928392602084019060208552518091526040840192915b818110610d17575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610d09565b8254845260209093019260019283019201610cd9565b9050346101fe57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101fe576060810181811067ffffffffffffffff821117610e6157604052602181527f43435450546f6b656e4d657373656e67657250726f787920312e372e302d646560208201527f76000000000000000000000000000000000000000000000000000000000000006040820152604051809260208252825192836020840152815b848110610e49575050601f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe092604080968601015201168101030190f35b60208282018101516040888401015286945001610e0a565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610ecf57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203610f2157565b600080fd5b81601f82011215610f215780359167ffffffffffffffff8311610ecf578260051b9160405193610f596020850186610e8e565b8452602080850193820101918211610f2157602001915b818310610f7d5750505090565b823573ffffffffffffffffffffffffffffffffffffffff81168103610f2157815260209283019201610f70565b6024359063ffffffff82168203610f2157565b60a4359063ffffffff82168203610f2157565b60c4359063ffffffff82168203610f2157565b73ffffffffffffffffffffffffffffffffffffffff60015416330361100457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156110425760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3360005260036020526040600020541561108757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b80548210156110425760005260206000200190600090565b600081815260036020526040902054801561128b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161125c57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161125c578082036111ed575b50505060025480156111be577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161117b8160026110b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6112446111fe61120f9360026110b5565b90549060031b1c92839260026110b5565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080611142565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146112ec5760025468010000000000000000811015610ecf576112d361120f82600185940160025560026110b5565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var CCTPTokenMessengerProxyABI = CCTPTokenMessengerProxyMetaData.ABI

var CCTPTokenMessengerProxyBin = CCTPTokenMessengerProxyMetaData.Bin

func DeployCCTPTokenMessengerProxy(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, usdcToken common.Address, authorizedCallers []common.Address) (common.Address, *types.Transaction, *CCTPTokenMessengerProxy, error) {
	parsed, err := CCTPTokenMessengerProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPTokenMessengerProxyBin), backend, tokenMessenger, usdcToken, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPTokenMessengerProxy{address: address, abi: *parsed, CCTPTokenMessengerProxyCaller: CCTPTokenMessengerProxyCaller{contract: contract}, CCTPTokenMessengerProxyTransactor: CCTPTokenMessengerProxyTransactor{contract: contract}, CCTPTokenMessengerProxyFilterer: CCTPTokenMessengerProxyFilterer{contract: contract}}, nil
}

type CCTPTokenMessengerProxy struct {
	address common.Address
	abi     abi.ABI
	CCTPTokenMessengerProxyCaller
	CCTPTokenMessengerProxyTransactor
	CCTPTokenMessengerProxyFilterer
}

type CCTPTokenMessengerProxyCaller struct {
	contract *bind.BoundContract
}

type CCTPTokenMessengerProxyTransactor struct {
	contract *bind.BoundContract
}

type CCTPTokenMessengerProxyFilterer struct {
	contract *bind.BoundContract
}

type CCTPTokenMessengerProxySession struct {
	Contract     *CCTPTokenMessengerProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPTokenMessengerProxyCallerSession struct {
	Contract *CCTPTokenMessengerProxyCaller
	CallOpts bind.CallOpts
}

type CCTPTokenMessengerProxyTransactorSession struct {
	Contract     *CCTPTokenMessengerProxyTransactor
	TransactOpts bind.TransactOpts
}

type CCTPTokenMessengerProxyRaw struct {
	Contract *CCTPTokenMessengerProxy
}

type CCTPTokenMessengerProxyCallerRaw struct {
	Contract *CCTPTokenMessengerProxyCaller
}

type CCTPTokenMessengerProxyTransactorRaw struct {
	Contract *CCTPTokenMessengerProxyTransactor
}

func NewCCTPTokenMessengerProxy(address common.Address, backend bind.ContractBackend) (*CCTPTokenMessengerProxy, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPTokenMessengerProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPTokenMessengerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxy{address: address, abi: abi, CCTPTokenMessengerProxyCaller: CCTPTokenMessengerProxyCaller{contract: contract}, CCTPTokenMessengerProxyTransactor: CCTPTokenMessengerProxyTransactor{contract: contract}, CCTPTokenMessengerProxyFilterer: CCTPTokenMessengerProxyFilterer{contract: contract}}, nil
}

func NewCCTPTokenMessengerProxyCaller(address common.Address, caller bind.ContractCaller) (*CCTPTokenMessengerProxyCaller, error) {
	contract, err := bindCCTPTokenMessengerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyCaller{contract: contract}, nil
}

func NewCCTPTokenMessengerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPTokenMessengerProxyTransactor, error) {
	contract, err := bindCCTPTokenMessengerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyTransactor{contract: contract}, nil
}

func NewCCTPTokenMessengerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPTokenMessengerProxyFilterer, error) {
	contract, err := bindCCTPTokenMessengerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyFilterer{contract: contract}, nil
}

func bindCCTPTokenMessengerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPTokenMessengerProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPTokenMessengerProxy.Contract.CCTPTokenMessengerProxyCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.CCTPTokenMessengerProxyTransactor.contract.Transfer(opts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.CCTPTokenMessengerProxyTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPTokenMessengerProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.contract.Transfer(opts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetAllAuthorizedCallers(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetAllAuthorizedCallers(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) GetTokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "getTokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) GetTokenMessenger() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetTokenMessenger(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) GetTokenMessenger() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetTokenMessenger(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) GetUSDCToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "getUSDCToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) GetUSDCToken() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetUSDCToken(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) GetUSDCToken() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.GetUSDCToken(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) LocalMessageTransmitter() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.LocalMessageTransmitter(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.LocalMessageTransmitter(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) MessageBodyVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "messageBodyVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) MessageBodyVersion() (uint32, error) {
	return _CCTPTokenMessengerProxy.Contract.MessageBodyVersion(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) MessageBodyVersion() (uint32, error) {
	return _CCTPTokenMessengerProxy.Contract.MessageBodyVersion(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) Owner() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.Owner(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) Owner() (common.Address, error) {
	return _CCTPTokenMessengerProxy.Contract.Owner(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPTokenMessengerProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) TypeAndVersion() (string, error) {
	return _CCTPTokenMessengerProxy.Contract.TypeAndVersion(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyCallerSession) TypeAndVersion() (string, error) {
	return _CCTPTokenMessengerProxy.Contract.TypeAndVersion(&_CCTPTokenMessengerProxy.CallOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.AcceptOwnership(&_CCTPTokenMessengerProxy.TransactOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.AcceptOwnership(&_CCTPTokenMessengerProxy.TransactOpts)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.ApplyAuthorizedCallerUpdates(&_CCTPTokenMessengerProxy.TransactOpts, authorizedCallerArgs)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.ApplyAuthorizedCallerUpdates(&_CCTPTokenMessengerProxy.TransactOpts, authorizedCallerArgs)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "depositForBurn", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurn(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurn(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurnWithCaller(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurnWithCaller(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) DepositForBurnWithHook(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "depositForBurnWithHook", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) DepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurnWithHook(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) DepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.DepositForBurnWithHook(&_CCTPTokenMessengerProxy.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.TransferOwnership(&_CCTPTokenMessengerProxy.TransactOpts, to)
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPTokenMessengerProxy.Contract.TransferOwnership(&_CCTPTokenMessengerProxy.TransactOpts, to)
}

type CCTPTokenMessengerProxyAuthorizedCallerAddedIterator struct {
	Event *CCTPTokenMessengerProxyAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyAuthorizedCallerAdded)
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
		it.Event = new(CCTPTokenMessengerProxyAuthorizedCallerAdded)
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

func (it *CCTPTokenMessengerProxyAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPTokenMessengerProxyAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyAuthorizedCallerAddedIterator{contract: _CCTPTokenMessengerProxy.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyAuthorizedCallerAdded)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseAuthorizedCallerAdded(log types.Log) (*CCTPTokenMessengerProxyAuthorizedCallerAdded, error) {
	event := new(CCTPTokenMessengerProxyAuthorizedCallerAdded)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator struct {
	Event *CCTPTokenMessengerProxyAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyAuthorizedCallerRemoved)
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
		it.Event = new(CCTPTokenMessengerProxyAuthorizedCallerRemoved)
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

func (it *CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator{contract: _CCTPTokenMessengerProxy.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyAuthorizedCallerRemoved)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*CCTPTokenMessengerProxyAuthorizedCallerRemoved, error) {
	event := new(CCTPTokenMessengerProxyAuthorizedCallerRemoved)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenMessengerProxyDepositForBurnIterator struct {
	Event *CCTPTokenMessengerProxyDepositForBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyDepositForBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyDepositForBurn)
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
		it.Event = new(CCTPTokenMessengerProxyDepositForBurn)
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

func (it *CCTPTokenMessengerProxyDepositForBurnIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyDepositForBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyDepositForBurn struct {
	Nonce                     uint64
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	Raw                       types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*CCTPTokenMessengerProxyDepositForBurnIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyDepositForBurnIterator{contract: _CCTPTokenMessengerProxy.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyDepositForBurn)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseDepositForBurn(log types.Log) (*CCTPTokenMessengerProxyDepositForBurn, error) {
	event := new(CCTPTokenMessengerProxyDepositForBurn)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenMessengerProxyDepositForBurn0Iterator struct {
	Event *CCTPTokenMessengerProxyDepositForBurn0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyDepositForBurn0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyDepositForBurn0)
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
		it.Event = new(CCTPTokenMessengerProxyDepositForBurn0)
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

func (it *CCTPTokenMessengerProxyDepositForBurn0Iterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyDepositForBurn0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyDepositForBurn0 struct {
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	MaxFee                    uint32
	MinFinalityThreshold      uint32
	HookData                  []byte
	Raw                       types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterDepositForBurn0(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*CCTPTokenMessengerProxyDepositForBurn0Iterator, error) {

	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var minFinalityThresholdRule []interface{}
	for _, minFinalityThresholdItem := range minFinalityThreshold {
		minFinalityThresholdRule = append(minFinalityThresholdRule, minFinalityThresholdItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "DepositForBurn0", burnTokenRule, depositorRule, minFinalityThresholdRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyDepositForBurn0Iterator{contract: _CCTPTokenMessengerProxy.contract, event: "DepositForBurn0", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchDepositForBurn0(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyDepositForBurn0, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error) {

	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var minFinalityThresholdRule []interface{}
	for _, minFinalityThresholdItem := range minFinalityThreshold {
		minFinalityThresholdRule = append(minFinalityThresholdRule, minFinalityThresholdItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "DepositForBurn0", burnTokenRule, depositorRule, minFinalityThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyDepositForBurn0)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "DepositForBurn0", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseDepositForBurn0(log types.Log) (*CCTPTokenMessengerProxyDepositForBurn0, error) {
	event := new(CCTPTokenMessengerProxyDepositForBurn0)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "DepositForBurn0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenMessengerProxyOwnershipTransferRequestedIterator struct {
	Event *CCTPTokenMessengerProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyOwnershipTransferRequested)
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
		it.Event = new(CCTPTokenMessengerProxyOwnershipTransferRequested)
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

func (it *CCTPTokenMessengerProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenMessengerProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyOwnershipTransferRequestedIterator{contract: _CCTPTokenMessengerProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyOwnershipTransferRequested)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPTokenMessengerProxyOwnershipTransferRequested, error) {
	event := new(CCTPTokenMessengerProxyOwnershipTransferRequested)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenMessengerProxyOwnershipTransferredIterator struct {
	Event *CCTPTokenMessengerProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenMessengerProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenMessengerProxyOwnershipTransferred)
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
		it.Event = new(CCTPTokenMessengerProxyOwnershipTransferred)
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

func (it *CCTPTokenMessengerProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenMessengerProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenMessengerProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenMessengerProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenMessengerProxyOwnershipTransferredIterator{contract: _CCTPTokenMessengerProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenMessengerProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenMessengerProxyOwnershipTransferred)
				if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxyFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPTokenMessengerProxyOwnershipTransferred, error) {
	event := new(CCTPTokenMessengerProxyOwnershipTransferred)
	if err := _CCTPTokenMessengerProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (CCTPTokenMessengerProxyAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (CCTPTokenMessengerProxyAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (CCTPTokenMessengerProxyDepositForBurn) Topic() common.Hash {
	return common.HexToHash("0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0")
}

func (CCTPTokenMessengerProxyDepositForBurn0) Topic() common.Hash {
	return common.HexToHash("0x6a4c152b4ad8c08f204453d58ef2ac1c0bb69627dd545cf47507d32d036e67d5")
}

func (CCTPTokenMessengerProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPTokenMessengerProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCTPTokenMessengerProxy *CCTPTokenMessengerProxy) Address() common.Address {
	return _CCTPTokenMessengerProxy.address
}

type CCTPTokenMessengerProxyInterface interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenMessenger(opts *bind.CallOpts) (common.Address, error)

	GetUSDCToken(opts *bind.CallOpts) (common.Address, error)

	LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error)

	MessageBodyVersion(opts *bind.CallOpts) (uint32, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error)

	DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error)

	DepositForBurnWithHook(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPTokenMessengerProxyAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*CCTPTokenMessengerProxyAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPTokenMessengerProxyAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*CCTPTokenMessengerProxyAuthorizedCallerRemoved, error)

	FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*CCTPTokenMessengerProxyDepositForBurnIterator, error)

	WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseDepositForBurn(log types.Log) (*CCTPTokenMessengerProxyDepositForBurn, error)

	FilterDepositForBurn0(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*CCTPTokenMessengerProxyDepositForBurn0Iterator, error)

	WatchDepositForBurn0(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyDepositForBurn0, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error)

	ParseDepositForBurn0(log types.Log) (*CCTPTokenMessengerProxyDepositForBurn0, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenMessengerProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPTokenMessengerProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenMessengerProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPTokenMessengerProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPTokenMessengerProxyOwnershipTransferred, error)

	Address() common.Address
}
