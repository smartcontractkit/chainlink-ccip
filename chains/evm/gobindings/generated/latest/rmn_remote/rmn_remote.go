// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rmn_remote

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

type IRMNRemoteSignature struct {
	R [32]byte
	S [32]byte
}

type InternalMerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
}

type RMNRemoteConfig struct {
	RmnHomeContractConfigDigest [32]byte
	Signers                     []RMNRemoteSigner
	FSign                       uint64
}

type RMNRemoteSigner struct {
	OnchainPublicKey common.Address
	NodeIndex        uint64
}

var RMNRemoteMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"curseAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCursedSubjects\",\"inputs\":[],\"outputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocalChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReportDigestHeader\",\"inputs\":[],\"outputs\":[{\"name\":\"digestHeader\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getVersionedConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"newConfig\",\"type\":\"tuple\",\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"offRampAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"merkleRoots\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"struct IRMNRemote.Signature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Uncursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOnchainPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignerOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"NotEnoughSigners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfOrderSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroValueNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0604052346102565761291a803803806100198161025b565b92833981016040828203126102565781516001600160401b03811692838203610256576020810151906001600160401b038211610256570182601f82011215610256578051926001600160401b03841161020f578360051b9160208061008081860161025b565b80978152019382010191821161025657602001915b81831061023657505050331561022557600180546001600160a01b031916331790556020926100c38461025b565b9160008352600036813760408051949085016001600160401b0381118682101761020f576040528452828585015260005b835181101561015b576001906001600160a01b036101128287610280565b51168761011e826102c2565b61012b575b5050016100f4565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13887610123565b50915091519060005b82518110156101d5576001600160a01b0361017f8285610280565b51169081156101c4577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef86836101b66001956103c0565b50604051908152a101610164565b6342bcdf7f60e11b60005260046000fd5b5082156101fe576080526040516124f99081610421823960805181818161036a01526109680152f35b63273e150360e21b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b038116810361025657815260209283019201610095565b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761020f57604052565b80518210156102945760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156102945760005260206000200190600090565b60008181526003602052604090205480156103b95760001981018181116103a3576002546000198101919082116103a357808203610352575b505050600254801561033c57600019016103168160026102aa565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61038b6103636103749360026102aa565b90549060031b1c92839260026102aa565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102fb565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461041a576002546801000000000000000081101561020f5761040161037482600185940160025560026102aa565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714611bff578063198f0f77146115145780631add205f1461135e5780632451a627146112705780632cbc26bb1461122f578063397796f7146111ec57806362eed4151461106b5780636509a954146110125780636d2d399314610ef257806370a9089e1461084657806379ba50971461075d5780638da5cb5b1461070b57806391a2749a146105215780639a19b32914610433578063d881e0921461038e578063eaa83ddd1461032b578063f2fde38b1461023b5763f8bb876e146100e257600080fd5b34610236576100f036611e44565b73ffffffffffffffffffffffffffffffffffffffff6001541633036101ed575b60005b81518110156101b8576101517fffffffffffffffffffffffffffffffff0000000000000000000000000000000061014a8385612181565b51166123aa565b1561015e57600101610113565b610189907fffffffffffffffffffffffffffffffff0000000000000000000000000000000092612181565b51167f19d5c79b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040517f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f7490806101e88582611f0b565b0390a1005b610204336000526003602052604060002054151590565b610110577fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600080fd5b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365773ffffffffffffffffffffffffffffffffffffffff610287611d83565b61028f612195565b1633811461030157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657602060405167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760405180602060085491828152019060086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee39060005b81811061041d576104198561040d81870382611cb4565b60405191829182611f0b565b0390f35b82548452602090930192600192830192016103f6565b346102365761044136611e44565b610449612195565b60005b81518110156104f15761048a7fffffffffffffffffffffffffffffffff000000000000000000000000000000006104838385612181565b5116612273565b156104975760010161044c565b6104c2907fffffffffffffffffffffffffffffffff0000000000000000000000000000000092612181565b51167f73281fa10000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806101e88582611f0b565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff81116102365760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610236576040519061059b82611c98565b806004013567ffffffffffffffff8111610236576105bf9060043691840101611ddf565b825260248101359067ffffffffffffffff82116102365760046105e59236920101611ddf565b602082019081526105f4612195565b519060005b825181101561066c578073ffffffffffffffffffffffffffffffffffffffff61062460019386612181565b511661062f8161241c565b61063b575b50016105f9565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610634565b505160005b81518110156107095773ffffffffffffffffffffffffffffffffffffffff6106998284612181565b51169081156106df577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836106d16001956123e3565b50604051908152a101610671565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b005b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760005473ffffffffffffffffffffffffffffffffffffffff8116330361081c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102365760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365761087d611d83565b67ffffffffffffffff602435116102365736602360243501121561023657602435600401359067ffffffffffffffff8211610236573660248360051b8135010111610236576044359067ffffffffffffffff821161023657366023830112156102365767ffffffffffffffff82600401351161023657366024836004013560061b840101116102365763ffffffff6007541615610ec85767ffffffffffffffff61092a816006541661200f565b16826004013510610e9e5760045460405160c0810181811067ffffffffffffffff821117610e6f57604052468152602081019267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168452604082019230845273ffffffffffffffffffffffffffffffffffffffff6060840192168252608083019081526109bf87611dc7565b916109cd6040519384611cb4565b8783526000976024803501602085015b60248360051b813501018210610cb7575050509073ffffffffffffffffffffffffffffffffffffffff8095939260a0860193845260405196879567ffffffffffffffff602088019a7f9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf538c526040808a01526101208901995160608a015251166080880152511660a0860152511660c08401525160e08301525160c0610100830152805180935261014082019260206101408260051b85010192019388905b828210610c1e57505050610ad69250037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282611cb4565b5190208290815b83600401358310610aec578480f35b6020856080610b0386886004013560248a01611fbb565b3583610b17888a6004013560248c01611fbb565b013560405191878352601b868401526040830152606082015282805260015afa15610c135784519073ffffffffffffffffffffffffffffffffffffffff8216908115610beb5773ffffffffffffffffffffffffffffffffffffffff8291161015610bc3578552600a60205260ff60408620541615610b9b5760019290920191610add565b6004857faaaa9141000000000000000000000000000000000000000000000000000000008152fd5b6004867fbbe15e7f000000000000000000000000000000000000000000000000000000008152fd5b6004877f8baa579f000000000000000000000000000000000000000000000000000000008152fd5b6040513d86823e3d90fd5b91936020847ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08293600195970301855287519067ffffffffffffffff8251168152608080610c798585015160a08786015260a0850190611cf5565b9367ffffffffffffffff604082015116604085015267ffffffffffffffff606082015116606085015201519101529601920192018593919492610a9b565b81359067ffffffffffffffff8211610e6b5760a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc836024350136030112610e6b576040519060a0820182811067ffffffffffffffff821117610e3a57604052610d2660248481350101612068565b82526044836024350101359167ffffffffffffffff8311610e675736604360243586018501011215610e675767ffffffffffffffff6024848682350101013511610e3a57908d9160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6024878982350101013501160193610dad6040519586611cb4565b60248035870182019081013580875236910160440111610e3657602495602095869560a49387908a90813586018101808301359060440186850137858235010101358301015285840152610e0660648289350101612068565b6040840152610e1a60848289350101612068565b60608401528635010135608082015281520192019190506109dd565b8380fd5b60248e7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8d80fd5b8b80fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f59fa4a930000000000000000000000000000000000000000000000000000000060005260046000fd5b7face124bc0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657610f29611d54565b604090815190610f398383611cb4565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083013660208401377fffffffffffffffffffffffffffffffff00000000000000000000000000000000610f9083612174565b91169052610f9c612195565b60005b8151811015610fe357610fd67fffffffffffffffffffffffffffffffff000000000000000000000000000000006104838385612181565b1561049757600101610f9f565b82517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806101e88582611f0b565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760206040517f9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf538152f35b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576110a2611d54565b6040908151906110b28383611cb4565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083013660208401377fffffffffffffffffffffffffffffffff0000000000000000000000000000000061110983612174565b9116905273ffffffffffffffffffffffffffffffffffffffff6001541633036111a3575b60005b8151811015611174576111677fffffffffffffffffffffffffffffffff0000000000000000000000000000000061014a8385612181565b1561015e57600101611130565b82517f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f7490806101e88582611f0b565b6111ba336000526003602052604060002054151590565b61112d577fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576020611225612117565b6040519015158152f35b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657602061122561126b611d54565b61207d565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061134857505050816112ef910382611cb4565b6040519182916020830190602084525180915260408301919060005b818110611319575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161130b565b82548452602090930192600192830192016112d9565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760006040805161139c81611c7c565b82815260606020820152015263ffffffff60075416604051906113be82611c7c565b6004548252600554916113d083611dc7565b926113de6040519485611cb4565b808452600560009081527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0602086015b8383106114c857505050506020810192835267ffffffffffffffff6006541692604082019384526040519283526040602084015260a083019151604084015251906060808401528151809152602060c0840192019060005b8181106114865750505067ffffffffffffffff8293511660808301520390f35b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015167ffffffffffffffff168186015260409094019390920191600101611466565b6001602081926040516114da81611c98565b67ffffffffffffffff865473ffffffffffffffffffffffffffffffffffffffff8116835260a01c168382015281520192019201919061140e565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff811161023657806004018136039160607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8401126102365761158f612195565b8135908115611bd557909260248201919060015b6115ad8486611f67565b905081101561168e576115c08486611f67565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83019183831161165f57611601926020926115fb92611fbb565b01611ffa565b67ffffffffffffffff8061162460206115fb8661161e8b8d611f67565b90611fbb565b1691161015611635576001016115a3565b7f448515170000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5083908561169c8584611f67565b6044860191506116ab82611ffa565b60011b6801fffffffffffffffe67fffffffffffffffe82169116810361165f576116dd67ffffffffffffffff9161200f565b1611611bab57600554805b611aa8575060005b6116fa8786611f67565b90508110156117d35773ffffffffffffffffffffffffffffffffffffffff61172e6117298361161e8b8a611f67565b612047565b16600052600a60205260ff604060002054166117a9578073ffffffffffffffffffffffffffffffffffffffff61176d61172960019461161e8c8b611f67565b16600052600a6020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055016116f0565b7f28cae27d0000000000000000000000000000000000000000000000000000000060005260046000fd5b50846117e58780959685600455611f67565b90680100000000000000008211610e6f5760055482600555808310611a62575b50600560005260206000206000915b8383106119b4575050505067ffffffffffffffff61183183611ffa565b167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000060065416176006556007549463ffffffff86169563ffffffff871461165f5763ffffffff60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000098011696879116176007557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60405195602087526080870195602088015235910181121561023657016024600482013591019267ffffffffffffffff8211610236578160061b36038413610236578190606060408701525260a08401929060005b81811061196457867f7f22bf988149dbe8de8fb879c6b97a4e56e68b2bd57421ce1a4e79d4ef6b496c87808867ffffffffffffffff6119598a612068565b1660608301520390a2005b90919360408060019273ffffffffffffffffffffffffffffffffffffffff61198b89611da6565b16815267ffffffffffffffff6119a360208a01612068565b16602082015201950192910161191b565b600160408273ffffffffffffffffffffffffffffffffffffffff6119d88495612047565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000865416178555611a0c60208201611ffa565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000087549260a01b16911617855501920192019190611814565b60056000527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db09081019083015b818110611a9c5750611805565b60008155600101611a8f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161165f576000906005541115611b7e57600590527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3daf81015473ffffffffffffffffffffffffffffffffffffffff166000908152600a6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055801561165f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01806116e8565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b7f014c50200000000000000000000000000000000000000000000000000000000060005260046000fd5b7f9cf8540c0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576104196040805190611c408183611cb4565b600f82527f524d4e52656d6f746520322e312e300000000000000000000000000000000000602083015251918291602083526020830190611cf5565b6060810190811067ffffffffffffffff821117610e6f57604052565b6040810190811067ffffffffffffffff821117610e6f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610e6f57604052565b919082519283825260005b848110611d3f5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611d00565b600435907fffffffffffffffffffffffffffffffff000000000000000000000000000000008216820361023657565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361023657565b359073ffffffffffffffffffffffffffffffffffffffff8216820361023657565b67ffffffffffffffff8111610e6f5760051b60200190565b9080601f83011215610236578135611df681611dc7565b92611e046040519485611cb4565b81845260208085019260051b82010192831161023657602001905b828210611e2c5750505090565b60208091611e3984611da6565b815201910190611e1f565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112610236576004359067ffffffffffffffff8211610236578060238301121561023657816004013590611e9b82611dc7565b92611ea96040519485611cb4565b8284526024602085019360051b82010191821161023657602401915b818310611ed25750505090565b82357fffffffffffffffffffffffffffffffff000000000000000000000000000000008116810361023657815260209283019201611ec5565b602060408183019282815284518094520192019060005b818110611f2f5750505090565b82517fffffffffffffffffffffffffffffffff0000000000000000000000000000000016845260209384019390920191600101611f22565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610236570180359067ffffffffffffffff821161023657602001918160061b3603831361023657565b9190811015611fcb5760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036102365790565b67ffffffffffffffff60019116019067ffffffffffffffff821161165f57565b8054821015611fcb5760005260206000200190600090565b3573ffffffffffffffffffffffffffffffffffffffff811681036102365790565b359067ffffffffffffffff8216820361023657565b60085415612111577fffffffffffffffffffffffffffffffff00000000000000000000000000000000166000526009602052604060002054158015906120c05790565b507f010000000000000000000000000000010000000000000000000000000000000060005260096020527f9ed150dd87eadff6494e247a6a2f4527ecd448746381dba90d27370069b832f054151590565b50600090565b6008541561216f577f010000000000000000000000000000010000000000000000000000000000000060005260096020527f9ed150dd87eadff6494e247a6a2f4527ecd448746381dba90d27370069b832f054151590565b600090565b805115611fcb5760200190565b8051821015611fcb5760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff6001541633036121b657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548015612244577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190612215828261202f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600960205260409020548015612378577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161165f57600854907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161165f57818103612309575b5050506122f560086121e0565b600052600960205260006040812055600190565b61236061231a61232b93600861202f565b90549060031b1c928392600861202f565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260096020526040600020553880806122e8565b5050600090565b80549068010000000000000000821015610e6f578161232b9160016123a69401815561202f565b9055565b80600052600960205260406000205415600014612111576123cc81600861237f565b600854906000526009602052604060002055600190565b806000526003602052604060002054156000146121115761240581600261237f565b600254906000526003602052604060002055600190565b6000818152600360205260409020548015612378577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161165f57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161165f578082036124b2575b50505061249e60026121e0565b600052600360205260006040812055600190565b6124d46124c361232b93600261202f565b90549060031b1c928392600261202f565b9055600052600360205260406000205538808061249156fea164736f6c634300081a000a",
}

var RMNRemoteABI = RMNRemoteMetaData.ABI

var RMNRemoteBin = RMNRemoteMetaData.Bin

func DeployRMNRemote(auth *bind.TransactOpts, backend bind.ContractBackend, localChainSelector uint64, curseAdmins []common.Address) (common.Address, *types.Transaction, *RMNRemote, error) {
	parsed, err := RMNRemoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RMNRemoteBin), backend, localChainSelector, curseAdmins)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RMNRemote{address: address, abi: *parsed, RMNRemoteCaller: RMNRemoteCaller{contract: contract}, RMNRemoteTransactor: RMNRemoteTransactor{contract: contract}, RMNRemoteFilterer: RMNRemoteFilterer{contract: contract}}, nil
}

type RMNRemote struct {
	address common.Address
	abi     abi.ABI
	RMNRemoteCaller
	RMNRemoteTransactor
	RMNRemoteFilterer
}

type RMNRemoteCaller struct {
	contract *bind.BoundContract
}

type RMNRemoteTransactor struct {
	contract *bind.BoundContract
}

type RMNRemoteFilterer struct {
	contract *bind.BoundContract
}

type RMNRemoteSession struct {
	Contract     *RMNRemote
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RMNRemoteCallerSession struct {
	Contract *RMNRemoteCaller
	CallOpts bind.CallOpts
}

type RMNRemoteTransactorSession struct {
	Contract     *RMNRemoteTransactor
	TransactOpts bind.TransactOpts
}

type RMNRemoteRaw struct {
	Contract *RMNRemote
}

type RMNRemoteCallerRaw struct {
	Contract *RMNRemoteCaller
}

type RMNRemoteTransactorRaw struct {
	Contract *RMNRemoteTransactor
}

func NewRMNRemote(address common.Address, backend bind.ContractBackend) (*RMNRemote, error) {
	abi, err := abi.JSON(strings.NewReader(RMNRemoteABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRMNRemote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RMNRemote{address: address, abi: abi, RMNRemoteCaller: RMNRemoteCaller{contract: contract}, RMNRemoteTransactor: RMNRemoteTransactor{contract: contract}, RMNRemoteFilterer: RMNRemoteFilterer{contract: contract}}, nil
}

func NewRMNRemoteCaller(address common.Address, caller bind.ContractCaller) (*RMNRemoteCaller, error) {
	contract, err := bindRMNRemote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCaller{contract: contract}, nil
}

func NewRMNRemoteTransactor(address common.Address, transactor bind.ContractTransactor) (*RMNRemoteTransactor, error) {
	contract, err := bindRMNRemote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteTransactor{contract: contract}, nil
}

func NewRMNRemoteFilterer(address common.Address, filterer bind.ContractFilterer) (*RMNRemoteFilterer, error) {
	contract, err := bindRMNRemote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteFilterer{contract: contract}, nil
}

func bindRMNRemote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RMNRemoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RMNRemote *RMNRemoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNRemote.Contract.RMNRemoteCaller.contract.Call(opts, result, method, params...)
}

func (_RMNRemote *RMNRemoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.Contract.RMNRemoteTransactor.contract.Transfer(opts)
}

func (_RMNRemote *RMNRemoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNRemote.Contract.RMNRemoteTransactor.contract.Transact(opts, method, params...)
}

func (_RMNRemote *RMNRemoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNRemote.Contract.contract.Call(opts, result, method, params...)
}

func (_RMNRemote *RMNRemoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.Contract.contract.Transfer(opts)
}

func (_RMNRemote *RMNRemoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNRemote.Contract.contract.Transact(opts, method, params...)
}

func (_RMNRemote *RMNRemoteCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _RMNRemote.Contract.GetAllAuthorizedCallers(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _RMNRemote.Contract.GetAllAuthorizedCallers(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getCursedSubjects")

	if err != nil {
		return *new([][16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][16]byte)).(*[][16]byte)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMNRemote.Contract.GetCursedSubjects(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMNRemote.Contract.GetCursedSubjects(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) GetLocalChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getLocalChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetLocalChainSelector() (uint64, error) {
	return _RMNRemote.Contract.GetLocalChainSelector(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetLocalChainSelector() (uint64, error) {
	return _RMNRemote.Contract.GetLocalChainSelector(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) GetReportDigestHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getReportDigestHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetReportDigestHeader() ([32]byte, error) {
	return _RMNRemote.Contract.GetReportDigestHeader(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetReportDigestHeader() ([32]byte, error) {
	return _RMNRemote.Contract.GetReportDigestHeader(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) GetVersionedConfig(opts *bind.CallOpts) (GetVersionedConfig,

	error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getVersionedConfig")

	outstruct := new(GetVersionedConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Config = *abi.ConvertType(out[1], new(RMNRemoteConfig)).(*RMNRemoteConfig)

	return *outstruct, err

}

func (_RMNRemote *RMNRemoteSession) GetVersionedConfig() (GetVersionedConfig,

	error) {
	return _RMNRemote.Contract.GetVersionedConfig(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetVersionedConfig() (GetVersionedConfig,

	error) {
	return _RMNRemote.Contract.GetVersionedConfig(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "isCursed", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMNRemote.Contract.IsCursed(&_RMNRemote.CallOpts, subject)
}

func (_RMNRemote *RMNRemoteCallerSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMNRemote.Contract.IsCursed(&_RMNRemote.CallOpts, subject)
}

func (_RMNRemote *RMNRemoteCaller) IsCursed0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "isCursed0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) IsCursed0() (bool, error) {
	return _RMNRemote.Contract.IsCursed0(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) IsCursed0() (bool, error) {
	return _RMNRemote.Contract.IsCursed0(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) Owner() (common.Address, error) {
	return _RMNRemote.Contract.Owner(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) Owner() (common.Address, error) {
	return _RMNRemote.Contract.Owner(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) TypeAndVersion() (string, error) {
	return _RMNRemote.Contract.TypeAndVersion(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) TypeAndVersion() (string, error) {
	return _RMNRemote.Contract.TypeAndVersion(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) Verify(opts *bind.CallOpts, offRampAddress common.Address, merkleRoots []InternalMerkleRoot, signatures []IRMNRemoteSignature) error {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "verify", offRampAddress, merkleRoots, signatures)

	if err != nil {
		return err
	}

	return err

}

func (_RMNRemote *RMNRemoteSession) Verify(offRampAddress common.Address, merkleRoots []InternalMerkleRoot, signatures []IRMNRemoteSignature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, offRampAddress, merkleRoots, signatures)
}

func (_RMNRemote *RMNRemoteCallerSession) Verify(offRampAddress common.Address, merkleRoots []InternalMerkleRoot, signatures []IRMNRemoteSignature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, offRampAddress, merkleRoots, signatures)
}

func (_RMNRemote *RMNRemoteTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "acceptOwnership")
}

func (_RMNRemote *RMNRemoteSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNRemote.Contract.AcceptOwnership(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNRemote.Contract.AcceptOwnership(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_RMNRemote *RMNRemoteSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMNRemote.Contract.ApplyAuthorizedCallerUpdates(&_RMNRemote.TransactOpts, authorizedCallerArgs)
}

func (_RMNRemote *RMNRemoteTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMNRemote.Contract.ApplyAuthorizedCallerUpdates(&_RMNRemote.TransactOpts, authorizedCallerArgs)
}

func (_RMNRemote *RMNRemoteTransactor) Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "curse", subject)
}

func (_RMNRemote *RMNRemoteSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactorSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactor) Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "curse0", subjects)
}

func (_RMNRemote *RMNRemoteSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse0(&_RMNRemote.TransactOpts, subjects)
}

func (_RMNRemote *RMNRemoteTransactorSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse0(&_RMNRemote.TransactOpts, subjects)
}

func (_RMNRemote *RMNRemoteTransactor) SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "setConfig", newConfig)
}

func (_RMNRemote *RMNRemoteSession) SetConfig(newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.Contract.SetConfig(&_RMNRemote.TransactOpts, newConfig)
}

func (_RMNRemote *RMNRemoteTransactorSession) SetConfig(newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.Contract.SetConfig(&_RMNRemote.TransactOpts, newConfig)
}

func (_RMNRemote *RMNRemoteTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "transferOwnership", to)
}

func (_RMNRemote *RMNRemoteSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

func (_RMNRemote *RMNRemoteTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

func (_RMNRemote *RMNRemoteTransactor) Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "uncurse", subject)
}

func (_RMNRemote *RMNRemoteSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactorSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactor) Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "uncurse0", subjects)
}

func (_RMNRemote *RMNRemoteSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse0(&_RMNRemote.TransactOpts, subjects)
}

func (_RMNRemote *RMNRemoteTransactorSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse0(&_RMNRemote.TransactOpts, subjects)
}

type RMNRemoteAuthorizedCallerAddedIterator struct {
	Event *RMNRemoteAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteAuthorizedCallerAdded)
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
		it.Event = new(RMNRemoteAuthorizedCallerAdded)
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

func (it *RMNRemoteAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*RMNRemoteAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteAuthorizedCallerAddedIterator{contract: _RMNRemote.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *RMNRemoteAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteAuthorizedCallerAdded)
				if err := _RMNRemote.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseAuthorizedCallerAdded(log types.Log) (*RMNRemoteAuthorizedCallerAdded, error) {
	event := new(RMNRemoteAuthorizedCallerAdded)
	if err := _RMNRemote.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteAuthorizedCallerRemovedIterator struct {
	Event *RMNRemoteAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteAuthorizedCallerRemoved)
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
		it.Event = new(RMNRemoteAuthorizedCallerRemoved)
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

func (it *RMNRemoteAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*RMNRemoteAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteAuthorizedCallerRemovedIterator{contract: _RMNRemote.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *RMNRemoteAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteAuthorizedCallerRemoved)
				if err := _RMNRemote.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*RMNRemoteAuthorizedCallerRemoved, error) {
	event := new(RMNRemoteAuthorizedCallerRemoved)
	if err := _RMNRemote.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteConfigSetIterator struct {
	Event *RMNRemoteConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteConfigSet)
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
		it.Event = new(RMNRemoteConfigSet)
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

func (it *RMNRemoteConfigSetIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteConfigSet struct {
	Version uint32
	Config  RMNRemoteConfig
	Raw     types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterConfigSet(opts *bind.FilterOpts, version []uint32) (*RMNRemoteConfigSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "ConfigSet", versionRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteConfigSetIterator{contract: _RMNRemote.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet, version []uint32) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "ConfigSet", versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteConfigSet)
				if err := _RMNRemote.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error) {
	event := new(RMNRemoteConfigSet)
	if err := _RMNRemote.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteCursedIterator struct {
	Event *RMNRemoteCursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteCursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteCursed)
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
		it.Event = new(RMNRemoteCursed)
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

func (it *RMNRemoteCursedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteCursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteCursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterCursed(opts *bind.FilterOpts) (*RMNRemoteCursedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCursedIterator{contract: _RMNRemote.contract, event: "Cursed", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteCursed) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteCursed)
				if err := _RMNRemote.contract.UnpackLog(event, "Cursed", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseCursed(log types.Log) (*RMNRemoteCursed, error) {
	event := new(RMNRemoteCursed)
	if err := _RMNRemote.contract.UnpackLog(event, "Cursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteOwnershipTransferRequestedIterator struct {
	Event *RMNRemoteOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteOwnershipTransferRequested)
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
		it.Event = new(RMNRemoteOwnershipTransferRequested)
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

func (it *RMNRemoteOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteOwnershipTransferRequestedIterator{contract: _RMNRemote.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteOwnershipTransferRequested)
				if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseOwnershipTransferRequested(log types.Log) (*RMNRemoteOwnershipTransferRequested, error) {
	event := new(RMNRemoteOwnershipTransferRequested)
	if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteOwnershipTransferredIterator struct {
	Event *RMNRemoteOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteOwnershipTransferred)
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
		it.Event = new(RMNRemoteOwnershipTransferred)
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

func (it *RMNRemoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteOwnershipTransferredIterator{contract: _RMNRemote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteOwnershipTransferred)
				if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseOwnershipTransferred(log types.Log) (*RMNRemoteOwnershipTransferred, error) {
	event := new(RMNRemoteOwnershipTransferred)
	if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteUncursedIterator struct {
	Event *RMNRemoteUncursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteUncursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteUncursed)
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
		it.Event = new(RMNRemoteUncursed)
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

func (it *RMNRemoteUncursedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteUncursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteUncursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterUncursed(opts *bind.FilterOpts) (*RMNRemoteUncursedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteUncursedIterator{contract: _RMNRemote.contract, event: "Uncursed", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteUncursed) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteUncursed)
				if err := _RMNRemote.contract.UnpackLog(event, "Uncursed", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseUncursed(log types.Log) (*RMNRemoteUncursed, error) {
	event := new(RMNRemoteUncursed)
	if err := _RMNRemote.contract.UnpackLog(event, "Uncursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetVersionedConfig struct {
	Version uint32
	Config  RMNRemoteConfig
}

func (RMNRemoteAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (RMNRemoteAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (RMNRemoteConfigSet) Topic() common.Hash {
	return common.HexToHash("0x7f22bf988149dbe8de8fb879c6b97a4e56e68b2bd57421ce1a4e79d4ef6b496c")
}

func (RMNRemoteCursed) Topic() common.Hash {
	return common.HexToHash("0x1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f74")
}

func (RMNRemoteOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RMNRemoteOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (RMNRemoteUncursed) Topic() common.Hash {
	return common.HexToHash("0x0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba1")
}

func (_RMNRemote *RMNRemote) Address() common.Address {
	return _RMNRemote.address
}

type RMNRemoteInterface interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error)

	GetLocalChainSelector(opts *bind.CallOpts) (uint64, error)

	GetReportDigestHeader(opts *bind.CallOpts) ([32]byte, error)

	GetVersionedConfig(opts *bind.CallOpts) (GetVersionedConfig,

		error)

	IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error)

	IsCursed0(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Verify(opts *bind.CallOpts, offRampAddress common.Address, merkleRoots []InternalMerkleRoot, signatures []IRMNRemoteSignature) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*RMNRemoteAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *RMNRemoteAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*RMNRemoteAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*RMNRemoteAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *RMNRemoteAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*RMNRemoteAuthorizedCallerRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts, version []uint32) (*RMNRemoteConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet, version []uint32) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error)

	FilterCursed(opts *bind.FilterOpts) (*RMNRemoteCursedIterator, error)

	WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteCursed) (event.Subscription, error)

	ParseCursed(log types.Log) (*RMNRemoteCursed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RMNRemoteOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RMNRemoteOwnershipTransferred, error)

	FilterUncursed(opts *bind.FilterOpts) (*RMNRemoteUncursedIterator, error)

	WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteUncursed) (event.Subscription, error)

	ParseUncursed(log types.Log) (*RMNRemoteUncursed, error)

	Address() common.Address
}
