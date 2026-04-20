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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCurseAdminUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurseAdmins\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCursedSubjects\",\"inputs\":[],\"outputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocalChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReportDigestHeader\",\"inputs\":[],\"outputs\":[{\"name\":\"digestHeader\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getVersionedConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCurseAdmin\",\"inputs\":[{\"name\":\"curseAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"newConfig\",\"type\":\"tuple\",\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"offRampAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"merkleRoots\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"struct IRMNRemote.Signature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RMNRemote.Config\",\"components\":[{\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"struct RMNRemote.Signer[]\",\"components\":[{\"name\":\"onchainPublicKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"fSign\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CurseAdminAdded\",\"inputs\":[{\"name\":\"curseAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CurseAdminRemoved\",\"inputs\":[{\"name\":\"curseAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Uncursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOnchainPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignerOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"NotEnoughSigners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyOwnerOrCurseAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OutOfOrderSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroValueNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a03460ab57601f61254938819003918201601f19168301916001600160401b0383118484101760b05780849260209460405283398101031260ab57516001600160401b03811680820360ab573315609a57600180546001600160a01b031916331790551560895760805260405161248290816100c7823960805181818161037e01526108cc0152f35b63273e150360e21b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714611b80578063198f0f77146114965780631add205f146112cb5780632cbc26bb1461128a578063397796f7146112515780636271638f146110a957806362eed41514610f4d5780636509a95414610ef45780636d2d399314610dd457806370a9089e146107cd57806379ba5097146106e45780638da5cb5b146106925780639a19b329146105a4578063b260f165146104b6578063d881e09214610411578063dd25e7f7146103a2578063eaa83ddd1461033f578063f2fde38b1461024f5763f8bb876e146100ed57600080fd5b3461024a576100fb36611d54565b73ffffffffffffffffffffffffffffffffffffffff600154163314158061022c575b6101fe5760005b81518110156101c9576101627fffffffffffffffffffffffffffffffff0000000000000000000000000000000061015b83856120a1565b511661241b565b1561016f57600101610124565b61019a907fffffffffffffffffffffffffffffffff00000000000000000000000000000000926120a1565b51167f19d5c79b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040517f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f7490806101f98582611e1b565b0390a1005b7fbdc8a0a7000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b50610244336000526009602052604060002054151590565b1561011d565b600080fd5b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5773ffffffffffffffffffffffffffffffffffffffff61029b611d19565b6102a36120b5565b1633811461031557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57602060405167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57602061040773ffffffffffffffffffffffffffffffffffffffff6103f3611d19565b166000526009602052604060002054151590565b6040519015158152f35b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760405180602060065491828152019060066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9060005b8181106104a05761049c8561049081870382611c19565b60405191829182611e1b565b0390f35b8254845260209093019260019283019201610479565b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a576040518060206008549283815201809260086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee39060005b81811061058e5750505081610535910382611c19565b6040519182916020830190602084525180915260408301919060005b81811061055f575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610551565b825484526020909301926001928301920161051f565b3461024a576105b236611d54565b6105ba6120b5565b60005b8151811015610662576105fb7fffffffffffffffffffffffffffffffff000000000000000000000000000000006105f483856120a1565b5116612296565b15610608576001016105bd565b610633907fffffffffffffffffffffffffffffffff00000000000000000000000000000000926120a1565b51167f73281fa10000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806101f98582611e1b565b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760005473ffffffffffffffffffffffffffffffffffffffff811633036107a3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57610804611d19565b60243567ffffffffffffffff811161024a57610824903690600401611ce8565b916044359267ffffffffffffffff841161024a573660238501121561024a5783600401359267ffffffffffffffff841161024a576024850194602436918660061b01011161024a5763ffffffff6005541615610daa5767ffffffffffffffff6108908160045416611f1f565b168410610d8057600254906040519160c0830183811067ffffffffffffffff821117610d5157604052468352602083019467ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168652604084019430865273ffffffffffffffffffffffffffffffffffffffff60608601921682526080850192835261092381611d3c565b936109316040519586611c19565b818552602085019160051b81019036821161024a5780925b828410610c1357505050509073ffffffffffffffffffffffffffffffffffffffff8095939260a0860193845260405196879567ffffffffffffffff602088019a7f9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf538c526040808a01526101208901995160608a015251166080880152511660a0860152511660c08401525160e08301525160c0610100830152805180935261014082019260206101408260051b8501019201936000905b828210610b7a57505050610a3b9250037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282611c19565b519020600092835b838510610a4c57005b602060006080610a5d888887611ecb565b3583610a6a8a8a89611ecb565b013560405191888352601b868401526040830152606082015282805260015afa15610b6e576000519073ffffffffffffffffffffffffffffffffffffffff8216908115610b445773ffffffffffffffffffffffffffffffffffffffff8291161015610b1a57600052600a60205260ff6040600020541615610af057600190940193610a43565b7faaaa91410000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbbe15e7f0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f8baa579f0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b91936020847ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08293600195970301855287519067ffffffffffffffff8251168152608080610bd58585015160a08786015260a0850190611c5a565b9367ffffffffffffffff604082015116604085015267ffffffffffffffff606082015116606085015201519101529601920192018593919492610a00565b833567ffffffffffffffff811161024a5782019060a08236031261024a576040519060a0820182811067ffffffffffffffff821117610d5157604052610c5883611f78565b8252602083013567ffffffffffffffff811161024a5783019136601f8401121561024a57823592600067ffffffffffffffff8511610d24575060405194610cc7601f86017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200187611c19565b848652366020868401011161024a576020956000878781988260809701838601378301015285840152610cfc60408201611f78565b6040840152610d0d60608201611f78565b606084015201356080820152815201930192610949565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f59fa4a930000000000000000000000000000000000000000000000000000000060005260046000fd5b7face124bc0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57610e0b611cb9565b604090815190610e1b8383611c19565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083013660208401377fffffffffffffffffffffffffffffffff00000000000000000000000000000000610e7283612094565b91169052610e7e6120b5565b60005b8151811015610ec557610eb87fffffffffffffffffffffffffffffffff000000000000000000000000000000006105f483856120a1565b1561060857600101610e81565b82517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806101f98582611e1b565b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760206040517f9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf538152f35b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a57610f84611cb9565b604090815190610f948383611c19565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083013660208401377fffffffffffffffffffffffffffffffff00000000000000000000000000000000610feb83612094565b9116905273ffffffffffffffffffffffffffffffffffffffff600154163314158061108b575b6101fe5760005b815181101561105c5761104f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000061015b83856120a1565b1561016f57600101611018565b82517f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f7490806101f98582611e1b565b506110a3336000526009602052604060002054151590565b15611011565b3461024a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760043567ffffffffffffffff811161024a576110f8903690600401611ce8565b60243567ffffffffffffffff811161024a57611118903690600401611ce8565b9290916111236120b5565b60005b8181106111c95750505060005b82811061113c57005b8061117073ffffffffffffffffffffffffffffffffffffffff61116a6111656001958888612084565b611f57565b166123c1565b61117b575b01611133565b73ffffffffffffffffffffffffffffffffffffffff61119e611165838787612084565b167f4a9e3347c9e1e3758466519e66ece7a69006ac2d00be8bf1ae9249c24e9e2319600080a2611175565b806111f873ffffffffffffffffffffffffffffffffffffffff6111f26111656001958789612084565b16612100565b611203575b01611126565b73ffffffffffffffffffffffffffffffffffffffff611226611165838688612084565b167f70b6d1d3991617d244ed75113a2d36f5a0f03e3985cbf6bd4db78c5b7b5bdf76600080a26111fd565b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a576020610407612027565b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760206104076112c6611cb9565b611f8d565b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760006040805161130981611bfd565b82815260606020820152015263ffffffff600554166040519061132b82611bfd565b60025482526003549161133d83611d3c565b9261134b6040519485611c19565b808452600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9190602086015b82821061143757505050506020810192835267ffffffffffffffff6004541692604082019384526040519283526040602084015260a083019151604084015251906060808401528151809152602060c0840192019060005b8181106113f55750505067ffffffffffffffff8293511660808301520390f35b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015167ffffffffffffffff1681860152604090940193909201916001016113d5565b6040516040810181811067ffffffffffffffff821117610d5157600192839260209260405267ffffffffffffffff885473ffffffffffffffffffffffffffffffffffffffff8116835260a01c168382015281520194019101909261137d565b3461024a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5760043567ffffffffffffffff811161024a57806004018136039160607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc84011261024a576115116120b5565b8135908115611b5657909260248201919060015b61152f8486611e77565b9050811015611610576115428486611e77565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301918383116115e1576115839260209261157d92611ecb565b01611f0a565b67ffffffffffffffff806115a6602061157d866115a08b8d611e77565b90611ecb565b16911610156115b757600101611525565b7f448515170000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5083908561161e8584611e77565b60448601915061162d82611f0a565b60011b6801fffffffffffffffe67fffffffffffffffe8216911681036115e15761165f67ffffffffffffffff91611f1f565b1611611b2c57600354805b611a29575060005b61167c8786611e77565b90508110156117505773ffffffffffffffffffffffffffffffffffffffff6116ab611165836115a08b8a611e77565b16600052600a60205260ff60406000205416611726578073ffffffffffffffffffffffffffffffffffffffff6116ea6111656001946115a08c8b611e77565b16600052600a6020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905501611672565b7f28cae27d0000000000000000000000000000000000000000000000000000000060005260046000fd5b50846117628780959685600255611e77565b90680100000000000000008211610d5157600354826003558083106119e3575b50600360005260206000206000915b838310611935575050505067ffffffffffffffff6117ae83611f0a565b167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000060045416176004556005549463ffffffff86169563ffffffff87146115e15763ffffffff60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000098011696879116176005557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60405195602087526080870195602088015235910181121561024a57016024600482013591019267ffffffffffffffff821161024a578160061b3603841361024a578190606060408701525260a0840192906000905b8082106118e257867f7f22bf988149dbe8de8fb879c6b97a4e56e68b2bd57421ce1a4e79d4ef6b496c87808867ffffffffffffffff6118d78a611f78565b1660608301520390a2005b90919384359073ffffffffffffffffffffffffffffffffffffffff821680920361024a5760408160019382935267ffffffffffffffff61192460208a01611f78565b166020820152019501920190611899565b600160408273ffffffffffffffffffffffffffffffffffffffff6119598495611f57565b167fffffffffffffffffffffffff000000000000000000000000000000000000000086541617855561198d60208201611f0a565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000087549260a01b16911617855501920192019190611791565b60036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9081019083015b818110611a1d5750611782565b60008155600101611a10565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116115e1576000906003541115611aff57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85a81015473ffffffffffffffffffffffffffffffffffffffff166000908152600a6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905580156115e1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018061166a565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b7f014c50200000000000000000000000000000000000000000000000000000000060005260046000fd5b7f9cf8540c0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024a5761049c6040805190611bc18183611c19565b600f82527f524d4e52656d6f746520312e362e300000000000000000000000000000000000602083015251918291602083526020830190611c5a565b6060810190811067ffffffffffffffff821117610d5157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d5157604052565b919082519283825260005b848110611ca45750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611c65565b600435907fffffffffffffffffffffffffffffffff000000000000000000000000000000008216820361024a57565b9181601f8401121561024a5782359167ffffffffffffffff831161024a576020808501948460051b01011161024a57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361024a57565b67ffffffffffffffff8111610d515760051b60200190565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261024a576004359067ffffffffffffffff821161024a578060238301121561024a57816004013590611dab82611d3c565b92611db96040519485611c19565b8284526024602085019360051b82010191821161024a57602401915b818310611de25750505090565b82357fffffffffffffffffffffffffffffffff000000000000000000000000000000008116810361024a57815260209283019201611dd5565b602060408183019282815284518094520192019060005b818110611e3f5750505090565b82517fffffffffffffffffffffffffffffffff0000000000000000000000000000000016845260209384019390920191600101611e32565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561024a570180359067ffffffffffffffff821161024a57602001918160061b3603831361024a57565b9190811015611edb5760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff8116810361024a5790565b67ffffffffffffffff60019116019067ffffffffffffffff82116115e157565b8054821015611edb5760005260206000200190600090565b3573ffffffffffffffffffffffffffffffffffffffff8116810361024a5790565b359067ffffffffffffffff8216820361024a57565b60065415612021577fffffffffffffffffffffffffffffffff0000000000000000000000000000000016600052600760205260406000205415801590611fd05790565b507f010000000000000000000000000000010000000000000000000000000000000060005260076020527f70b766b11586b6b505ed3893938b0cc6c6c98bd6f65e969ac311168d34e4f9e254151590565b50600090565b6006541561207f577f010000000000000000000000000000010000000000000000000000000000000060005260076020527f70b766b11586b6b505ed3893938b0cc6c6c98bd6f65e969ac311168d34e4f9e254151590565b600090565b9190811015611edb5760051b0190565b805115611edb5760200190565b8051821015611edb5760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff6001541633036120d657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b600081815260096020526040902054801561228f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116115e157600854907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116115e157818103612220575b50505060085480156121f1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016121ae816008611f3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600855600052600960205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b612277612231612242936008611f3f565b90549060031b1c9283926008611f3f565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526009602052604060002055388080612175565b5050600090565b600081815260076020526040902054801561228f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116115e157600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116115e157818103612387575b50505060065480156121f1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612344816006611f3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b6123a9612398612242936006611f3f565b90549060031b1c9283926006611f3f565b9055600052600760205260406000205538808061230b565b806000526009602052604060002054156000146120215760085468010000000000000000811015610d51576124026122428260018594016008556008611f3f565b9055600854906000526009602052604060002055600190565b806000526007602052604060002054156000146120215760065468010000000000000000811015610d515761245c6122428260018594016006556006611f3f565b905560065490600052600760205260406000205560019056fea164736f6c634300081a000a",
}

var RMNRemoteABI = RMNRemoteMetaData.ABI

var RMNRemoteBin = RMNRemoteMetaData.Bin

func DeployRMNRemote(auth *bind.TransactOpts, backend bind.ContractBackend, localChainSelector uint64) (common.Address, *types.Transaction, *RMNRemote, error) {
	parsed, err := RMNRemoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RMNRemoteBin), backend, localChainSelector)
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

func (_RMNRemote *RMNRemoteCaller) GetCurseAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getCurseAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetCurseAdmins() ([]common.Address, error) {
	return _RMNRemote.Contract.GetCurseAdmins(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetCurseAdmins() ([]common.Address, error) {
	return _RMNRemote.Contract.GetCurseAdmins(&_RMNRemote.CallOpts)
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

func (_RMNRemote *RMNRemoteCaller) IsCurseAdmin(opts *bind.CallOpts, curseAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "isCurseAdmin", curseAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) IsCurseAdmin(curseAdmin common.Address) (bool, error) {
	return _RMNRemote.Contract.IsCurseAdmin(&_RMNRemote.CallOpts, curseAdmin)
}

func (_RMNRemote *RMNRemoteCallerSession) IsCurseAdmin(curseAdmin common.Address) (bool, error) {
	return _RMNRemote.Contract.IsCurseAdmin(&_RMNRemote.CallOpts, curseAdmin)
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

func (_RMNRemote *RMNRemoteTransactor) ApplyCurseAdminUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "applyCurseAdminUpdates", removes, adds)
}

func (_RMNRemote *RMNRemoteSession) ApplyCurseAdminUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.ApplyCurseAdminUpdates(&_RMNRemote.TransactOpts, removes, adds)
}

func (_RMNRemote *RMNRemoteTransactorSession) ApplyCurseAdminUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.ApplyCurseAdminUpdates(&_RMNRemote.TransactOpts, removes, adds)
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

type RMNRemoteCurseAdminAddedIterator struct {
	Event *RMNRemoteCurseAdminAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteCurseAdminAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteCurseAdminAdded)
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
		it.Event = new(RMNRemoteCurseAdminAdded)
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

func (it *RMNRemoteCurseAdminAddedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteCurseAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteCurseAdminAdded struct {
	CurseAdmin common.Address
	Raw        types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterCurseAdminAdded(opts *bind.FilterOpts, curseAdmin []common.Address) (*RMNRemoteCurseAdminAddedIterator, error) {

	var curseAdminRule []interface{}
	for _, curseAdminItem := range curseAdmin {
		curseAdminRule = append(curseAdminRule, curseAdminItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "CurseAdminAdded", curseAdminRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCurseAdminAddedIterator{contract: _RMNRemote.contract, event: "CurseAdminAdded", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchCurseAdminAdded(opts *bind.WatchOpts, sink chan<- *RMNRemoteCurseAdminAdded, curseAdmin []common.Address) (event.Subscription, error) {

	var curseAdminRule []interface{}
	for _, curseAdminItem := range curseAdmin {
		curseAdminRule = append(curseAdminRule, curseAdminItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "CurseAdminAdded", curseAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteCurseAdminAdded)
				if err := _RMNRemote.contract.UnpackLog(event, "CurseAdminAdded", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseCurseAdminAdded(log types.Log) (*RMNRemoteCurseAdminAdded, error) {
	event := new(RMNRemoteCurseAdminAdded)
	if err := _RMNRemote.contract.UnpackLog(event, "CurseAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteCurseAdminRemovedIterator struct {
	Event *RMNRemoteCurseAdminRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteCurseAdminRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteCurseAdminRemoved)
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
		it.Event = new(RMNRemoteCurseAdminRemoved)
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

func (it *RMNRemoteCurseAdminRemovedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteCurseAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteCurseAdminRemoved struct {
	CurseAdmin common.Address
	Raw        types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterCurseAdminRemoved(opts *bind.FilterOpts, curseAdmin []common.Address) (*RMNRemoteCurseAdminRemovedIterator, error) {

	var curseAdminRule []interface{}
	for _, curseAdminItem := range curseAdmin {
		curseAdminRule = append(curseAdminRule, curseAdminItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "CurseAdminRemoved", curseAdminRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCurseAdminRemovedIterator{contract: _RMNRemote.contract, event: "CurseAdminRemoved", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchCurseAdminRemoved(opts *bind.WatchOpts, sink chan<- *RMNRemoteCurseAdminRemoved, curseAdmin []common.Address) (event.Subscription, error) {

	var curseAdminRule []interface{}
	for _, curseAdminItem := range curseAdmin {
		curseAdminRule = append(curseAdminRule, curseAdminItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "CurseAdminRemoved", curseAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteCurseAdminRemoved)
				if err := _RMNRemote.contract.UnpackLog(event, "CurseAdminRemoved", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseCurseAdminRemoved(log types.Log) (*RMNRemoteCurseAdminRemoved, error) {
	event := new(RMNRemoteCurseAdminRemoved)
	if err := _RMNRemote.contract.UnpackLog(event, "CurseAdminRemoved", log); err != nil {
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

func (RMNRemoteConfigSet) Topic() common.Hash {
	return common.HexToHash("0x7f22bf988149dbe8de8fb879c6b97a4e56e68b2bd57421ce1a4e79d4ef6b496c")
}

func (RMNRemoteCurseAdminAdded) Topic() common.Hash {
	return common.HexToHash("0x4a9e3347c9e1e3758466519e66ece7a69006ac2d00be8bf1ae9249c24e9e2319")
}

func (RMNRemoteCurseAdminRemoved) Topic() common.Hash {
	return common.HexToHash("0x70b6d1d3991617d244ed75113a2d36f5a0f03e3985cbf6bd4db78c5b7b5bdf76")
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
	GetCurseAdmins(opts *bind.CallOpts) ([]common.Address, error)

	GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error)

	GetLocalChainSelector(opts *bind.CallOpts) (uint64, error)

	GetReportDigestHeader(opts *bind.CallOpts) ([32]byte, error)

	GetVersionedConfig(opts *bind.CallOpts) (GetVersionedConfig,

		error)

	IsCurseAdmin(opts *bind.CallOpts, curseAdmin common.Address) (bool, error)

	IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error)

	IsCursed0(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Verify(opts *bind.CallOpts, offRampAddress common.Address, merkleRoots []InternalMerkleRoot, signatures []IRMNRemoteSignature) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyCurseAdminUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts, version []uint32) (*RMNRemoteConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet, version []uint32) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error)

	FilterCurseAdminAdded(opts *bind.FilterOpts, curseAdmin []common.Address) (*RMNRemoteCurseAdminAddedIterator, error)

	WatchCurseAdminAdded(opts *bind.WatchOpts, sink chan<- *RMNRemoteCurseAdminAdded, curseAdmin []common.Address) (event.Subscription, error)

	ParseCurseAdminAdded(log types.Log) (*RMNRemoteCurseAdminAdded, error)

	FilterCurseAdminRemoved(opts *bind.FilterOpts, curseAdmin []common.Address) (*RMNRemoteCurseAdminRemovedIterator, error)

	WatchCurseAdminRemoved(opts *bind.WatchOpts, sink chan<- *RMNRemoteCurseAdminRemoved, curseAdmin []common.Address) (event.Subscription, error)

	ParseCurseAdminRemoved(log types.Log) (*RMNRemoteCurseAdminRemoved, error)

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
