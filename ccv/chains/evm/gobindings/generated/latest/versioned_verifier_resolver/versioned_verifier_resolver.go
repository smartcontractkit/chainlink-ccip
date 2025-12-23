// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package versioned_verifier_resolver

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

type VersionedVerifierResolverInboundImplementationArgs struct {
	Version  [4]byte
	Verifier common.Address
}

type VersionedVerifierResolverOutboundImplementationArgs struct {
	DestChainSelector uint64
	Verifier          common.Address
}

var VersionedVerifierResolverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInboundImplementationUpdates\",\"inputs\":[{\"name\":\"implementations\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.InboundImplementationArgs[]\",\"components\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyOutboundImplementationUpdates\",\"inputs\":[{\"name\":\"implementations\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.OutboundImplementationArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllInboundImplementations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.InboundImplementationArgs[]\",\"components\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllOutboundImplementations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.OutboundImplementationArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInboundImplementation\",\"inputs\":[{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutboundImplementation\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeAggregatorUpdated\",\"inputs\":[{\"name\":\"oldFeeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newFeeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundImplementationRemoved\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundImplementationUpdated\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"prevImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundImplementationRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundImplementationUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"prevImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainSelector\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561198290816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806315b358e0146111cf578063181f5a771461112f5780635cb80c5d14610dde57806379ba509714610cf55780637a9c2ef914610a9b5780638da5cb5b14610a49578063958021a71461097d5780639cb406c91461092b578063b5cbfb6814610786578063c3a7ded61461068c578063c3eba22214610472578063e7076918146101a05763f2fde38b146100ab57600080fd5b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5773ffffffffffffffffffffffffffffffffffffffff6100f7611292565b6100ff6114a0565b1633811461017157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760043567ffffffffffffffff811161019b576101ef903690600401611404565b6101f76114a0565b60005b81811061020357005b61020e818385611435565b60408136031261019b57604051610224816112d6565b8135917fffffffff0000000000000000000000000000000000000000000000000000000083169283810361019b578252610260906020016112b5565b9173ffffffffffffffffffffffffffffffffffffffff602083019380855216156103da57507fffffffff0000000000000000000000000000000000000000000000000000000081511680156103ad5750606060019392827fffffffff000000000000000000000000000000000000000000000000000000007f240744c957da89d5c44d43838bbc5553c6ec57314f9e62435f9158c45b4e3413945116600052600260205273ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000008160406000205416928285511682825116600052600260205283604060002091167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905561038e82825116611853565b5051169251169060405192835260208301526040820152a15b016101fa565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60019392507fffffffff000000000000000000000000000000000000000000000000000000007f5dd8185b50a7df2c96bed0b91303df2507335646714c0d7896403165e4a58013926020926000526002835260406000207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055610463828251166116c8565b505116604051908152a16103a7565b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b576003546104ad81611474565b906104bb6040519283611321565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06104e882611474565b0160005b8181106106675750506003549060005b81811061058d578360405180916020820160208352815180915260206040840192019060005b818110610530575050500390f35b825180517fffffffff0000000000000000000000000000000000000000000000000000000016855260209081015173ffffffffffffffffffffffffffffffffffffffff168186015286955060409094019390920191600101610522565b60008382101561063a579073ffffffffffffffffffffffffffffffffffffffff60408260208560036001975220017fffffffff0000000000000000000000000000000000000000000000000000000060009154166105eb858a61148c565b51527fffffffff00000000000000000000000000000000000000000000000000000000610618858a61148c565b515116815260026020522054166020610631838861148c565b510152016104fc565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b602090604051610676816112d6565b60008152600083820152828287010152016104ec565b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760043567ffffffffffffffff811161019b573660238201121561019b57806004013567ffffffffffffffff811161019b57366024828401011161019b576004811061075c5760041161019b5760247fffffffff00000000000000000000000000000000000000000000000000000000910135166000526002602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b7f535e7c6d0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b576006546107c181611474565b906107cf6040519283611321565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06107fc82611474565b0160005b8181106109065750506006549060005b818110610889578360405180916020820160208352815180915260206040840192019060005b818110610844575050500390f35b8251805167ffffffffffffffff16855260209081015173ffffffffffffffffffffffffffffffffffffffff168186015286955060409094019390920191600101610836565b60008382101561063a579073ffffffffffffffffffffffffffffffffffffffff604082602085600660019752200167ffffffffffffffff60009154166108cf858a61148c565b515267ffffffffffffffff6108e4858a61148c565b5151168152600560205220541660206108fd838861148c565b51015201610810565b602090604051610915816112d6565b6000815260008382015282828701015201610800565b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b57602073ffffffffffffffffffffffffffffffffffffffff60085416604051908152f35b3461019b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760043567ffffffffffffffff811680910361019b5760243567ffffffffffffffff811161019b573660238201121561019b578060040135906109ed82611362565b916109fb6040519384611321565b808352366024828401011161019b5760009281602460209401848301370101526000526005602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760043567ffffffffffffffff811161019b57610aea903690600401611404565b610af26114a0565b60005b818110610afe57005b610b09818385611435565b60408136031261019b57604051610b1f816112d6565b81359167ffffffffffffffff83169283810361019b578252610b43906020016112b5565b9173ffffffffffffffffffffffffffffffffffffffff60208301938085521615610c75575067ffffffffffffffff8151168015610c4857506060600193928267ffffffffffffffff7fc12b226506536cd62f34841a87d2333621e547ff4af0f3b13f3ac204bfb47ab1945116600052600560205273ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff8160406000205416928285511682825116600052600560205283604060002091167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055610c29828251166117f3565b5051169251169060405192835260208301526040820152a15b01610af5565b7fef75b4cf0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600193925067ffffffffffffffff7f243416eecc562f47eb105155ee12ae26bb6e8dcbfce4c10e3ee69273e167214a926020926000526005835260406000207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055610ce682825116611503565b505116604051908152a1610c42565b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760005473ffffffffffffffffffffffffffffffffffffffff81163303610db4577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5760043567ffffffffffffffff811161019b573660238201121561019b57806004013567ffffffffffffffff811161019b573660248260051b8401011161019b5773ffffffffffffffffffffffffffffffffffffffff6008541680156111055760005b828110156111035760009060248160051b8601013573ffffffffffffffffffffffffffffffffffffffff81168091036110ff576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156110f45790859185916110bc575b5080610efc575b5050506001915001610e6f565b60405194610fb660208701967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602482015283604482015260448152610f46606482611321565b82806040998a5193610f588c86611321565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d156110b4573d90610f9982611362565b91610fa68b519384611321565b82523d85602084013e5b876118ad565b805180610ff5575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381610eef565b819293959697945090602091810103126110b05760200151908115918215036110ad575061102a579291908490888080610fbe565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b5080fd5b606090610fb0565b9150506020813d82116110ec575b816110d760209383611321565b810103126110e85784905188610ee8565b8380fd5b3d91506110ca565b6040513d86823e3d90fd5b8280fd5b005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461019b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b576111cb60405161116f606082611321565b602381527f56657273696f6e656456657269666965725265736f6c76657220312e372e302d60208201527f646576000000000000000000000000000000000000000000000000000000000060408201526040519182918261139c565b0390f35b3461019b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019b5773ffffffffffffffffffffffffffffffffffffffff61121b611292565b6112236114a0565b1680156111055773ffffffffffffffffffffffffffffffffffffffff600854827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600855167f5f93cfaedcfeead9f6922f03a6557cc9c40dd65f320e80dd4aa68fce736bf723600080a3005b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019b57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361019b57565b6040810190811067ffffffffffffffff8211176112f257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176112f257604052565b67ffffffffffffffff81116112f257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9190916020815282519283602083015260005b8481106113ee5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b80602080928401015160408286010152016113af565b9181601f8401121561019b5782359167ffffffffffffffff831161019b576020808501948460061b01011161019b57565b91908110156114455760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b67ffffffffffffffff81116112f25760051b60200190565b80518210156114455760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff6001541633036114c157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156114455760005260206000200190600090565b60008181526007602052604090205480156116c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161169257600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161169257818103611623575b50505060065480156115f4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016115b18160066114eb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61167a6116346116459360066114eb565b90549060031b1c92839260066114eb565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526007602052604060002055388080611578565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b60008181526004602052604090205480156116c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161169257600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611692578181036117b9575b50505060035480156115f4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016117768160036114eb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b6117db6117ca6116459360036114eb565b90549060031b1c92839260036114eb565b9055600052600460205260406000205538808061173d565b8060005260076020526040600020541560001461184d57600654680100000000000000008110156112f25761183461164582600185940160065560066114eb565b9055600654906000526007602052604060002055600190565b50600090565b8060005260046020526040600020541560001461184d57600354680100000000000000008110156112f25761189461164582600185940160035560036114eb565b9055600354906000526004602052604060002055600190565b9192901561192857508151156118c1575090565b3b156118ca5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561193b5750805190602001fd5b611971906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161139c565b0390fdfea164736f6c634300081a000a",
}

var VersionedVerifierResolverABI = VersionedVerifierResolverMetaData.ABI

var VersionedVerifierResolverBin = VersionedVerifierResolverMetaData.Bin

func DeployVersionedVerifierResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VersionedVerifierResolver, error) {
	parsed, err := VersionedVerifierResolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionedVerifierResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VersionedVerifierResolver{address: address, abi: *parsed, VersionedVerifierResolverCaller: VersionedVerifierResolverCaller{contract: contract}, VersionedVerifierResolverTransactor: VersionedVerifierResolverTransactor{contract: contract}, VersionedVerifierResolverFilterer: VersionedVerifierResolverFilterer{contract: contract}}, nil
}

type VersionedVerifierResolver struct {
	address common.Address
	abi     abi.ABI
	VersionedVerifierResolverCaller
	VersionedVerifierResolverTransactor
	VersionedVerifierResolverFilterer
}

type VersionedVerifierResolverCaller struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverTransactor struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverFilterer struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverSession struct {
	Contract     *VersionedVerifierResolver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VersionedVerifierResolverCallerSession struct {
	Contract *VersionedVerifierResolverCaller
	CallOpts bind.CallOpts
}

type VersionedVerifierResolverTransactorSession struct {
	Contract     *VersionedVerifierResolverTransactor
	TransactOpts bind.TransactOpts
}

type VersionedVerifierResolverRaw struct {
	Contract *VersionedVerifierResolver
}

type VersionedVerifierResolverCallerRaw struct {
	Contract *VersionedVerifierResolverCaller
}

type VersionedVerifierResolverTransactorRaw struct {
	Contract *VersionedVerifierResolverTransactor
}

func NewVersionedVerifierResolver(address common.Address, backend bind.ContractBackend) (*VersionedVerifierResolver, error) {
	abi, err := abi.JSON(strings.NewReader(VersionedVerifierResolverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVersionedVerifierResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolver{address: address, abi: abi, VersionedVerifierResolverCaller: VersionedVerifierResolverCaller{contract: contract}, VersionedVerifierResolverTransactor: VersionedVerifierResolverTransactor{contract: contract}, VersionedVerifierResolverFilterer: VersionedVerifierResolverFilterer{contract: contract}}, nil
}

func NewVersionedVerifierResolverCaller(address common.Address, caller bind.ContractCaller) (*VersionedVerifierResolverCaller, error) {
	contract, err := bindVersionedVerifierResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverCaller{contract: contract}, nil
}

func NewVersionedVerifierResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedVerifierResolverTransactor, error) {
	contract, err := bindVersionedVerifierResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverTransactor{contract: contract}, nil
}

func NewVersionedVerifierResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedVerifierResolverFilterer, error) {
	contract, err := bindVersionedVerifierResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverFilterer{contract: contract}, nil
}

func bindVersionedVerifierResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionedVerifierResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverCaller.contract.Call(opts, result, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverTransactor.contract.Transfer(opts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverTransactor.contract.Transact(opts, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedVerifierResolver.Contract.contract.Call(opts, result, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.contract.Transfer(opts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.contract.Transact(opts, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetAllInboundImplementations(opts *bind.CallOpts) ([]VersionedVerifierResolverInboundImplementationArgs, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getAllInboundImplementations")

	if err != nil {
		return *new([]VersionedVerifierResolverInboundImplementationArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]VersionedVerifierResolverInboundImplementationArgs)).(*[]VersionedVerifierResolverInboundImplementationArgs)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetAllInboundImplementations() ([]VersionedVerifierResolverInboundImplementationArgs, error) {
	return _VersionedVerifierResolver.Contract.GetAllInboundImplementations(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetAllInboundImplementations() ([]VersionedVerifierResolverInboundImplementationArgs, error) {
	return _VersionedVerifierResolver.Contract.GetAllInboundImplementations(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetAllOutboundImplementations(opts *bind.CallOpts) ([]VersionedVerifierResolverOutboundImplementationArgs, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getAllOutboundImplementations")

	if err != nil {
		return *new([]VersionedVerifierResolverOutboundImplementationArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]VersionedVerifierResolverOutboundImplementationArgs)).(*[]VersionedVerifierResolverOutboundImplementationArgs)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetAllOutboundImplementations() ([]VersionedVerifierResolverOutboundImplementationArgs, error) {
	return _VersionedVerifierResolver.Contract.GetAllOutboundImplementations(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetAllOutboundImplementations() ([]VersionedVerifierResolverOutboundImplementationArgs, error) {
	return _VersionedVerifierResolver.Contract.GetAllOutboundImplementations(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetFeeAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getFeeAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetFeeAggregator() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetFeeAggregator(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetFeeAggregator() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetFeeAggregator(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetInboundImplementation(opts *bind.CallOpts, verifierResults []byte) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getInboundImplementation", verifierResults)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetInboundImplementation(verifierResults []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementation(&_VersionedVerifierResolver.CallOpts, verifierResults)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetInboundImplementation(verifierResults []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementation(&_VersionedVerifierResolver.CallOpts, verifierResults)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetOutboundImplementation(opts *bind.CallOpts, destChainSelector uint64, arg1 []byte) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getOutboundImplementation", destChainSelector, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetOutboundImplementation(destChainSelector uint64, arg1 []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetOutboundImplementation(&_VersionedVerifierResolver.CallOpts, destChainSelector, arg1)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetOutboundImplementation(destChainSelector uint64, arg1 []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetOutboundImplementation(&_VersionedVerifierResolver.CallOpts, destChainSelector, arg1)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) Owner() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.Owner(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) Owner() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.Owner(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) TypeAndVersion() (string, error) {
	return _VersionedVerifierResolver.Contract.TypeAndVersion(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) TypeAndVersion() (string, error) {
	return _VersionedVerifierResolver.Contract.TypeAndVersion(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "acceptOwnership")
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) AcceptOwnership() (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.AcceptOwnership(&_VersionedVerifierResolver.TransactOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.AcceptOwnership(&_VersionedVerifierResolver.TransactOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) ApplyInboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "applyInboundImplementationUpdates", implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) ApplyInboundImplementationUpdates(implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyInboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) ApplyInboundImplementationUpdates(implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyInboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) ApplyOutboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "applyOutboundImplementationUpdates", implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) ApplyOutboundImplementationUpdates(implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyOutboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) ApplyOutboundImplementationUpdates(implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyOutboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "setFeeAggregator", feeAggregator)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.SetFeeAggregator(&_VersionedVerifierResolver.TransactOpts, feeAggregator)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.SetFeeAggregator(&_VersionedVerifierResolver.TransactOpts, feeAggregator)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "transferOwnership", to)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.TransferOwnership(&_VersionedVerifierResolver.TransactOpts, to)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.TransferOwnership(&_VersionedVerifierResolver.TransactOpts, to)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.WithdrawFeeTokens(&_VersionedVerifierResolver.TransactOpts, feeTokens)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.WithdrawFeeTokens(&_VersionedVerifierResolver.TransactOpts, feeTokens)
}

type VersionedVerifierResolverFeeAggregatorUpdatedIterator struct {
	Event *VersionedVerifierResolverFeeAggregatorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverFeeAggregatorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverFeeAggregatorUpdated)
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
		it.Event = new(VersionedVerifierResolverFeeAggregatorUpdated)
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

func (it *VersionedVerifierResolverFeeAggregatorUpdatedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverFeeAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverFeeAggregatorUpdated struct {
	OldFeeAggregator common.Address
	NewFeeAggregator common.Address
	Raw              types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterFeeAggregatorUpdated(opts *bind.FilterOpts, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (*VersionedVerifierResolverFeeAggregatorUpdatedIterator, error) {

	var oldFeeAggregatorRule []interface{}
	for _, oldFeeAggregatorItem := range oldFeeAggregator {
		oldFeeAggregatorRule = append(oldFeeAggregatorRule, oldFeeAggregatorItem)
	}
	var newFeeAggregatorRule []interface{}
	for _, newFeeAggregatorItem := range newFeeAggregator {
		newFeeAggregatorRule = append(newFeeAggregatorRule, newFeeAggregatorItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "FeeAggregatorUpdated", oldFeeAggregatorRule, newFeeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverFeeAggregatorUpdatedIterator{contract: _VersionedVerifierResolver.contract, event: "FeeAggregatorUpdated", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchFeeAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverFeeAggregatorUpdated, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (event.Subscription, error) {

	var oldFeeAggregatorRule []interface{}
	for _, oldFeeAggregatorItem := range oldFeeAggregator {
		oldFeeAggregatorRule = append(oldFeeAggregatorRule, oldFeeAggregatorItem)
	}
	var newFeeAggregatorRule []interface{}
	for _, newFeeAggregatorItem := range newFeeAggregator {
		newFeeAggregatorRule = append(newFeeAggregatorRule, newFeeAggregatorItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "FeeAggregatorUpdated", oldFeeAggregatorRule, newFeeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverFeeAggregatorUpdated)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "FeeAggregatorUpdated", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseFeeAggregatorUpdated(log types.Log) (*VersionedVerifierResolverFeeAggregatorUpdated, error) {
	event := new(VersionedVerifierResolverFeeAggregatorUpdated)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "FeeAggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverFeeTokenWithdrawnIterator struct {
	Event *VersionedVerifierResolverFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverFeeTokenWithdrawn)
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
		it.Event = new(VersionedVerifierResolverFeeTokenWithdrawn)
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

func (it *VersionedVerifierResolverFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*VersionedVerifierResolverFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverFeeTokenWithdrawnIterator{contract: _VersionedVerifierResolver.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverFeeTokenWithdrawn)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseFeeTokenWithdrawn(log types.Log) (*VersionedVerifierResolverFeeTokenWithdrawn, error) {
	event := new(VersionedVerifierResolverFeeTokenWithdrawn)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverInboundImplementationRemovedIterator struct {
	Event *VersionedVerifierResolverInboundImplementationRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverInboundImplementationRemoved)
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
		it.Event = new(VersionedVerifierResolverInboundImplementationRemoved)
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

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverInboundImplementationRemoved struct {
	Version [4]byte
	Raw     types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterInboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationRemovedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "InboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverInboundImplementationRemovedIterator{contract: _VersionedVerifierResolver.contract, event: "InboundImplementationRemoved", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchInboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationRemoved) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "InboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverInboundImplementationRemoved)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationRemoved", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseInboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverInboundImplementationRemoved, error) {
	event := new(VersionedVerifierResolverInboundImplementationRemoved)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverInboundImplementationUpdatedIterator struct {
	Event *VersionedVerifierResolverInboundImplementationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverInboundImplementationUpdated)
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
		it.Event = new(VersionedVerifierResolverInboundImplementationUpdated)
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

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverInboundImplementationUpdated struct {
	Version  [4]byte
	PrevImpl common.Address
	NewImpl  common.Address
	Raw      types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterInboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationUpdatedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "InboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverInboundImplementationUpdatedIterator{contract: _VersionedVerifierResolver.contract, event: "InboundImplementationUpdated", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchInboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "InboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverInboundImplementationUpdated)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationUpdated", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseInboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverInboundImplementationUpdated, error) {
	event := new(VersionedVerifierResolverInboundImplementationUpdated)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOutboundImplementationRemovedIterator struct {
	Event *VersionedVerifierResolverOutboundImplementationRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOutboundImplementationRemoved)
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
		it.Event = new(VersionedVerifierResolverOutboundImplementationRemoved)
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

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOutboundImplementationRemoved struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOutboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationRemovedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OutboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOutboundImplementationRemovedIterator{contract: _VersionedVerifierResolver.contract, event: "OutboundImplementationRemoved", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOutboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationRemoved) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OutboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOutboundImplementationRemoved)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationRemoved", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOutboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverOutboundImplementationRemoved, error) {
	event := new(VersionedVerifierResolverOutboundImplementationRemoved)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOutboundImplementationUpdatedIterator struct {
	Event *VersionedVerifierResolverOutboundImplementationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOutboundImplementationUpdated)
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
		it.Event = new(VersionedVerifierResolverOutboundImplementationUpdated)
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

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOutboundImplementationUpdated struct {
	DestChainSelector uint64
	PrevImpl          common.Address
	NewImpl           common.Address
	Raw               types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOutboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationUpdatedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OutboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOutboundImplementationUpdatedIterator{contract: _VersionedVerifierResolver.contract, event: "OutboundImplementationUpdated", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOutboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OutboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOutboundImplementationUpdated)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationUpdated", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOutboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverOutboundImplementationUpdated, error) {
	event := new(VersionedVerifierResolverOutboundImplementationUpdated)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOwnershipTransferRequestedIterator struct {
	Event *VersionedVerifierResolverOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOwnershipTransferRequested)
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
		it.Event = new(VersionedVerifierResolverOwnershipTransferRequested)
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

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOwnershipTransferRequestedIterator{contract: _VersionedVerifierResolver.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOwnershipTransferRequested)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOwnershipTransferRequested(log types.Log) (*VersionedVerifierResolverOwnershipTransferRequested, error) {
	event := new(VersionedVerifierResolverOwnershipTransferRequested)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOwnershipTransferredIterator struct {
	Event *VersionedVerifierResolverOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOwnershipTransferred)
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
		it.Event = new(VersionedVerifierResolverOwnershipTransferred)
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

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOwnershipTransferredIterator{contract: _VersionedVerifierResolver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOwnershipTransferred)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOwnershipTransferred(log types.Log) (*VersionedVerifierResolverOwnershipTransferred, error) {
	event := new(VersionedVerifierResolverOwnershipTransferred)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (VersionedVerifierResolverFeeAggregatorUpdated) Topic() common.Hash {
	return common.HexToHash("0x5f93cfaedcfeead9f6922f03a6557cc9c40dd65f320e80dd4aa68fce736bf723")
}

func (VersionedVerifierResolverFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (VersionedVerifierResolverInboundImplementationRemoved) Topic() common.Hash {
	return common.HexToHash("0x5dd8185b50a7df2c96bed0b91303df2507335646714c0d7896403165e4a58013")
}

func (VersionedVerifierResolverInboundImplementationUpdated) Topic() common.Hash {
	return common.HexToHash("0x240744c957da89d5c44d43838bbc5553c6ec57314f9e62435f9158c45b4e3413")
}

func (VersionedVerifierResolverOutboundImplementationRemoved) Topic() common.Hash {
	return common.HexToHash("0x243416eecc562f47eb105155ee12ae26bb6e8dcbfce4c10e3ee69273e167214a")
}

func (VersionedVerifierResolverOutboundImplementationUpdated) Topic() common.Hash {
	return common.HexToHash("0xc12b226506536cd62f34841a87d2333621e547ff4af0f3b13f3ac204bfb47ab1")
}

func (VersionedVerifierResolverOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VersionedVerifierResolverOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VersionedVerifierResolver *VersionedVerifierResolver) Address() common.Address {
	return _VersionedVerifierResolver.address
}

type VersionedVerifierResolverInterface interface {
	GetAllInboundImplementations(opts *bind.CallOpts) ([]VersionedVerifierResolverInboundImplementationArgs, error)

	GetAllOutboundImplementations(opts *bind.CallOpts) ([]VersionedVerifierResolverOutboundImplementationArgs, error)

	GetFeeAggregator(opts *bind.CallOpts) (common.Address, error)

	GetInboundImplementation(opts *bind.CallOpts, verifierResults []byte) (common.Address, error)

	GetOutboundImplementation(opts *bind.CallOpts, destChainSelector uint64, arg1 []byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyInboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error)

	ApplyOutboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error)

	SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterFeeAggregatorUpdated(opts *bind.FilterOpts, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (*VersionedVerifierResolverFeeAggregatorUpdatedIterator, error)

	WatchFeeAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverFeeAggregatorUpdated, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (event.Subscription, error)

	ParseFeeAggregatorUpdated(log types.Log) (*VersionedVerifierResolverFeeAggregatorUpdated, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*VersionedVerifierResolverFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*VersionedVerifierResolverFeeTokenWithdrawn, error)

	FilterInboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationRemovedIterator, error)

	WatchInboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationRemoved) (event.Subscription, error)

	ParseInboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverInboundImplementationRemoved, error)

	FilterInboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationUpdatedIterator, error)

	WatchInboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationUpdated) (event.Subscription, error)

	ParseInboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverInboundImplementationUpdated, error)

	FilterOutboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationRemovedIterator, error)

	WatchOutboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationRemoved) (event.Subscription, error)

	ParseOutboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverOutboundImplementationRemoved, error)

	FilterOutboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationUpdatedIterator, error)

	WatchOutboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationUpdated) (event.Subscription, error)

	ParseOutboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverOutboundImplementationUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VersionedVerifierResolverOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VersionedVerifierResolverOwnershipTransferred, error)

	Address() common.Address
}
