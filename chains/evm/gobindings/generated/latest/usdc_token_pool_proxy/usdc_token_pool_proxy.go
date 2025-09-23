// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool_proxy

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

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolLockOrBurnOutV1 struct {
	DestTokenAddress []byte
	DestPoolData     []byte
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

type PoolReleaseOrMintOutV1 struct {
	DestinationAmount *big.Int
}

type USDCTokenPoolProxyPoolAddresses struct {
	LegacyCctpV1Pool common.Address
	CctpV1Pool       common.Address
	CctpV2Pool       common.Address
}

var USDCTokenPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockReleasePoolAddress\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateLockReleasePoolAddresses\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"lockReleasePools\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleasePoolUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestinationPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c0604052346101d05760405161287938819003601f8101601f191683016001600160401b038111848210176101ba578392829160405283398101039060a082126101d05780516001600160a01b038116928382036101d057606090601f1901126101d057604051606081016001600160401b038111828210176101ba5760405261008c602084016101d5565b815261009a604084016101d5565b91602082019283526100c160806100b3606087016101d5565b9560408501968752016101d5565b9433156101a957600180546001600160a01b03191633179055158015610197575b8015610185575b8015610174575b6101635760805251600480546001600160a01b03199081166001600160a01b039384161790915591516005805484169183169190911790559151600680549092169083161790551660a05260405161268f90816101ea823960805181610a7f015260a05181818161098301526115370152f35b6303988b8160e61b60005260046000fd5b506001600160a01b038516156100f0565b5083516001600160a01b0316156100e9565b5082516001600160a01b0316156100e2565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b51906001600160a01b03821682036101d05756fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a77146100d75780631c255a9d146100d257806339077537146100cd578063673a2a1f146100c857806379ba5097146100c35780638302d09f146100be5780638da5cb5b146100b95780639a4575b9146100b4578063aa86a754146100af578063c527a6c2146100aa578063db4c2aef146100a55763f2fde38b146100a057600080fd5b610f80565b610e82565b610d25565b610cc0565b61089f565b6107fa565b610629565b61049f565b610416565b610391565b61031e565b61027b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761012757604052565b6100dc565b6020810190811067ffffffffffffffff82111761012757604052565b6040810190811067ffffffffffffffff82111761012757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761012757604052565b604051906101b561010083610164565b565b604051906101b561014083610164565b67ffffffffffffffff811161012757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106102145750506000910152565b8181015183820152602001610204565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361026081518092818752878088019101610201565b0116010190565b906020610278928181520190610224565b90565b346102fc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc576102f860408051906102bc8183610164565b601c82527f55534443546f6b656e506f6f6c50726f787920312e362e332d64657600000000602083015251918291602083526020830190610224565b0390f35b600080fd5b67ffffffffffffffff8116036102fc57565b35906101b582610301565b346102fc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5767ffffffffffffffff60043561036281610301565b166000526003602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346102fc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5760043567ffffffffffffffff81116102fc576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126102fc5761040d6020916004016114ba565b60405190518152f35b346102fc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc576000604080516104548161010b565b82815282602082015201526060610469611854565b73ffffffffffffffffffffffffffffffffffffffff60408051928281511684528260208201511660208501520151166040820152f35b346102fc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5760005473ffffffffffffffffffffffffffffffffffffffff8116330361055e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156102fc5782359167ffffffffffffffff83116102fc576020808501948460051b0101116102fc57565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102fc5760043567ffffffffffffffff81116102fc578161060291600401610588565b929092916024359067ffffffffffffffff82116102fc5761062591600401610588565b9091565b346102fc57610637366105b9565b926106406121d1565b8383036107d05760005b83811061065357005b6106826106696106648388876118bb565b61191e565b73ffffffffffffffffffffffffffffffffffffffff1690565b1515806107ad575b61078357806107126106a261066460019489886118bb565b6106d26106b86106b3858a896118bb565b611087565b67ffffffffffffffff166000526003602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6107206106b38287866118bb565b7f2a4ec2a96b51064b74fa8f2157f98cf8fb9fd4dcef4e9fdf3c44c0d74d0e826467ffffffffffffffff610758610664858b8a6118bb565b60405173ffffffffffffffffffffffffffffffffffffffff919091168152921691602090a20161064a565b7f8c909bc20000000000000000000000000000000000000000000000000000000060005260046000fd5b506107cb6107c76107c26106648489886118bb565b61221c565b1590565b61068a565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346102fc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b906102789160208152602061086c83516040838501526060840190610224565b9201519060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610224565b346102fc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5760043567ffffffffffffffff81116102fc5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126102fc57610914611928565b506024810161096a602061092783611087565b6040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b2157600091610c2e575b5073ffffffffffffffffffffffffffffffffffffffff33911603610c0457610a026109fb6109e183611087565b67ffffffffffffffff166000526002602052604060002090565b5460ff1690565b610a0b81610c8c565b8015610bd357610a19611854565b91600091610a2681610c8c565b60028103610b50575050506040015173ffffffffffffffffffffffffffffffffffffffff16905b73ffffffffffffffffffffffffffffffffffffffff8216918215610b26576000928392610aa3610adb936064830135907f0000000000000000000000000000000000000000000000000000000000000000612333565b6040519485809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260040160048301611a11565b03925af18015610b21576102f891600091610afe575b506040519182918261084c565b610b1b91503d806000833e610b138183610164565b810190611998565b38610af1565b6110a9565b7fb348dbbe0000000000000000000000000000000000000000000000000000000060005260046000fd5b610b5981610c8c565b60018103610b84575050506020015173ffffffffffffffffffffffffffffffffffffffff1690610a4d565b60039192949350610b9481610c8c565b14610ba0575b50610a4d565b610bcc9192506106b8610bb291611087565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9038610b9a565b610c00907f31603b1200000000000000000000000000000000000000000000000000000000600052610c9b565b6000fd5b7f82b429000000000000000000000000000000000000000000000000000000000060005260046000fd5b610c50915060203d602011610c56575b610c488183610164565b810190611941565b386109b4565b503d610c3e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60041115610c9657565b610c5d565b906024916004811015610c9657600452565b919060208301926004821015610c965752565b346102fc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5767ffffffffffffffff600435610d0481610301565b1660005260026020526102f860ff6040600020541660405191829182610cad565b346102fc5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc576000610d5e6121d1565b73ffffffffffffffffffffffffffffffffffffffff610d7b6118fa565b16158015610e70575b610e4857610d966107c76107c26118fa565b8015610e34575b8015610e0c575b610de457610db0611a99565b7fa33a66dd32843dc04c77f6aa048e4557f085fd2d793a098e06d6e9f89c715aec60405180610dde81611b7c565b0390a180f35b807f8c909bc20000000000000000000000000000000000000000000000000000000060049252fd5b50610e18610669611912565b15158015610da45750610e2f6107c76107c2611912565b610da4565b50610e436107c76107c2611906565b610d9d565b807fe622e0400000000000000000000000000000000000000000000000000000000060049252fd5b50610e7c610669611906565b15610d84565b346102fc57610e90366105b9565b92610e996121d1565b8383036107d05760005b838110610eac57005b80610ef6610ec5610ec060019489886118bb565b611bf6565b67ffffffffffffffff610ed98489886118bb565b35610ee381610301565b1660005260026020526040600020611c03565b610f046106b38287866118bb565b7f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc9447167ffffffffffffffff610f3c610ec0858b8a6118bb565b92610f4e604051928392169482610cad565b0390a201610ea3565b73ffffffffffffffffffffffffffffffffffffffff8116036102fc57565b35906101b582610f57565b346102fc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102fc5773ffffffffffffffffffffffffffffffffffffffff600435610fd081610f57565b610fd86121d1565b1633811461104a57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906110818261012c565b60008252565b3561027881610301565b908160209103126102fc575180151581036102fc5790565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156102fc570180359067ffffffffffffffff82116102fc576020019181360383136102fc57565b919091357fffffffff000000000000000000000000000000000000000000000000000000008116926004811061113a575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b908160209103126102fc57604051906111848261012c565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156102fc57016020813591019167ffffffffffffffff82116102fc5781360383136102fc57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90610278916020815261135261134761130a61124b611238868061118a565b61010060208801526101208701916111da565b61126b61125a60208801610313565b67ffffffffffffffff166040870152565b61129761127a60408801610f75565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b606086013560808601526112cd6112b060808801610f75565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b6112da60a087018761118a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c08801526111da565b61131760c086018661118a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e08701526111da565b9260e081019061118a565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526111da565b90610278916020815260e06114786114456113ac85516101006020870152610120860190610224565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152611411608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c0870152610224565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610224565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610224565b906004116102fc5790600490565b6114c2611074565b506020810161151e60206114d583611087565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015233602482015291829081906044820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b2157600091611825575b5015610c045760c08201907ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006115c86115c285876110b5565b90611106565b161461180657506115db604091836110b5565b905014611755576115ef60e08201826110b5565b9060048210611726576115c261160e92611608926114ac565b60e01c90565b63ffffffff8116806116be575050602061167f60009261164961066961066960055473ffffffffffffffffffffffffffffffffffffffff1690565b906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301611219565b03925af1908115610b2157600091611695575090565b610278915060203d6020116116b7575b6116af8183610164565b81019061116c565b503d6116a5565b6001036116f35750602061167f60009261164961066961066960065473ffffffffffffffffffffffffffffffffffffffff1690565b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b7f758b22cc00000000000000000000000000000000000000000000000000000000600052600482905260246000fd5b60208161177061176a60e060009501836110b5565b90611c81565b831461179f5761167f9061164961066961066960045473ffffffffffffffffffffffffffffffffffffffff1690565b61167f906117d16117cb61066961066960055473ffffffffffffffffffffffffffffffffffffffff1690565b91611f97565b6040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301611383565b60209150600092611649610669610669610bb26106b861167f96611087565b611847915060203d60201161184d575b61183f8183610164565b810190611091565b38611568565b503d611835565b604051906118618261010b565b8173ffffffffffffffffffffffffffffffffffffffff60045416815273ffffffffffffffffffffffffffffffffffffffff600554166020820152604073ffffffffffffffffffffffffffffffffffffffff60065416910152565b91908110156118cb5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60243561027881610f57565b60443561027881610f57565b60043561027881610f57565b3561027881610f57565b6040519061193582610148565b60606020838281520152565b908160209103126102fc575161027881610f57565b81601f820112156102fc57805161196c816101c7565b9261197a6040519485610164565b818452602082840101116102fc576102789160208085019101610201565b6020818303126102fc5780519067ffffffffffffffff82116102fc57016040818303126102fc57604051916119cc83610148565b815167ffffffffffffffff81116102fc57816119e9918401611956565b8352602082015167ffffffffffffffff81116102fc57611a099201611956565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080611a4b611a3b868061118a565b85602088015260c08701916111da565b9467ffffffffffffffff6020820135611a6381610301565b166040860152826040820135611a7881610f57565b1660608601526060810135828601520135611a9281610f57565b1691015290565b73ffffffffffffffffffffffffffffffffffffffff600435611aba81610f57565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045573ffffffffffffffffffffffffffffffffffffffff602435611b0581610f57565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055573ffffffffffffffffffffffffffffffffffffffff604435611b5081610f57565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655565b90606082019173ffffffffffffffffffffffffffffffffffffffff600435611ba381610f57565b16815273ffffffffffffffffffffffffffffffffffffffff602435611bc781610f57565b166020820152604073ffffffffffffffffffffffffffffffffffffffff604435611bf081610f57565b16910152565b3560048110156102fc5790565b906004811015610c965760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b81601f820112156102fc57803590611c51826101c7565b92611c5f6040519485610164565b828452602083830101116102fc57816000926020809301838601378301015290565b73ffffffffffffffffffffffffffffffffffffffff60045416918215611d48578101906020818303126102fc5780359067ffffffffffffffff82116102fc57016040818303126102fc5760405191611cd883610148565b813567ffffffffffffffff81116102fc5781611cf5918401611c3a565b8352602082013567ffffffffffffffff81116102fc57611d4493611d23610669936074936106699601611c3a565b602082015251015173ffffffffffffffffffffffffffffffffffffffff1690565b1490565b505050600090565b60405190610100820182811067ffffffffffffffff82111761012757604052606060e0838281526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152565b919091610100818403126102fc57611db86101a5565b92813567ffffffffffffffff81116102fc5781611dd6918401611c3a565b8452611de460208301610313565b6020850152611df560408301610f75565b604085015260608201356060850152611e1060808301610f75565b608085015260a082013567ffffffffffffffff81116102fc5781611e35918401611c3a565b60a085015260c082013567ffffffffffffffff81116102fc5781611e5a918401611c3a565b60c085015260e082013567ffffffffffffffff81116102fc57611e7d9201611c3a565b60e0830152565b63ffffffff8116036102fc57565b908160209103126102fc575161027881611e84565b91908260409103126102fc5760208235611ec081610301565b92013561027881611e84565b906003821015610c965752565b6101b59092919261012080610140830195611efe84825167ffffffffffffffff169052565b60208181015163ffffffff1690850152611f2060408201516040860190611ecc565b60608101516060850152611f416080820151608086019063ffffffff169052565b60a081015160a0850152611f7260c082015160c086019073ffffffffffffffffffffffffffffffffffffffff169052565b60e081015160e0850152610100810151610100850152015191019063ffffffff169052565b611f9f611d50565b50611faa3682611da2565b90611fd061066961066960055473ffffffffffffffffffffffffffffffffffffffff1690565b90604051917f6b716b0d000000000000000000000000000000000000000000000000000000008352602083600481845afa908115610b215760049360009261219f575b50602090604051948580927f98db96430000000000000000000000000000000000000000000000000000000082525afa928315610b21576121429361216e9373ffffffffffffffffffffffffffffffffffffffff92600092612176575b506120fb8161209061208860c061211d9501836110b5565b810190611ea7565b92909660608301356120e26120b960806120b261066961066960408a0161191e565b960161191e565b956120d56120c56101b7565b67ffffffffffffffff909c168c52565b63ffffffff1660208b0152565b600160408a0152606089015263ffffffff166080880152565b60a086015273ffffffffffffffffffffffffffffffffffffffff1660c0850152565b1660e08201526000610100820152600061012082015260405192839160208301611ed9565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610164565b60c082015290565b61211d9192506121976120fb9160203d602011610c5657610c488183610164565b929150612070565b60209192506121c390823d84116121ca575b6121bb8183610164565b810190611e92565b9190612013565b503d6121b1565b73ffffffffffffffffffffffffffffffffffffffff6001541633036121f257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000007f0e64dd29000000000000000000000000000000000000000000000000000000001660248201526024815261229e604482610164565b6179185a1061230957602091600091519084617530fa903d60005190836122fd575b50826122f3575b50816122e1575b816122d7575090565b61027891506124cf565b90506122ec81612433565b15906122ce565b15159150386122c7565b602011159250386122c0565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b9073ffffffffffffffffffffffffffffffffffffffff6124059392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452612395606485610164565b166000806040938451956123a98688610164565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d1561242a573d6123f66123ed826101c7565b94519485610164565b83523d6000602085013e6125ba565b805180612410575050565b81602080612425936101b59501019101611091565b61252f565b606092506125ba565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252612493604483610164565b6179185a10612309576020926000925191617530fa6000513d826124c3575b50816124bc575090565b9050151590565b602011159150386124b2565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252612493604483610164565b1561253657565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9192901561263557508151156125ce575090565b3b156125d75790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156126485750805190602001fd5b61267e906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610267565b0390fdfea164736f6c634300081a000a",
}

var USDCTokenPoolProxyABI = USDCTokenPoolProxyMetaData.ABI

var USDCTokenPoolProxyBin = USDCTokenPoolProxyMetaData.Bin

func DeployUSDCTokenPoolProxy(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, pools USDCTokenPoolProxyPoolAddresses, router common.Address) (common.Address, *types.Transaction, *USDCTokenPoolProxy, error) {
	parsed, err := USDCTokenPoolProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolProxyBin), backend, token, pools, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDCTokenPoolProxy{address: address, abi: *parsed, USDCTokenPoolProxyCaller: USDCTokenPoolProxyCaller{contract: contract}, USDCTokenPoolProxyTransactor: USDCTokenPoolProxyTransactor{contract: contract}, USDCTokenPoolProxyFilterer: USDCTokenPoolProxyFilterer{contract: contract}}, nil
}

type USDCTokenPoolProxy struct {
	address common.Address
	abi     abi.ABI
	USDCTokenPoolProxyCaller
	USDCTokenPoolProxyTransactor
	USDCTokenPoolProxyFilterer
}

type USDCTokenPoolProxyCaller struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxyTransactor struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxyFilterer struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxySession struct {
	Contract     *USDCTokenPoolProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolProxyCallerSession struct {
	Contract *USDCTokenPoolProxyCaller
	CallOpts bind.CallOpts
}

type USDCTokenPoolProxyTransactorSession struct {
	Contract     *USDCTokenPoolProxyTransactor
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolProxyRaw struct {
	Contract *USDCTokenPoolProxy
}

type USDCTokenPoolProxyCallerRaw struct {
	Contract *USDCTokenPoolProxyCaller
}

type USDCTokenPoolProxyTransactorRaw struct {
	Contract *USDCTokenPoolProxyTransactor
}

func NewUSDCTokenPoolProxy(address common.Address, backend bind.ContractBackend) (*USDCTokenPoolProxy, error) {
	abi, err := abi.JSON(strings.NewReader(USDCTokenPoolProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindUSDCTokenPoolProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxy{address: address, abi: abi, USDCTokenPoolProxyCaller: USDCTokenPoolProxyCaller{contract: contract}, USDCTokenPoolProxyTransactor: USDCTokenPoolProxyTransactor{contract: contract}, USDCTokenPoolProxyFilterer: USDCTokenPoolProxyFilterer{contract: contract}}, nil
}

func NewUSDCTokenPoolProxyCaller(address common.Address, caller bind.ContractCaller) (*USDCTokenPoolProxyCaller, error) {
	contract, err := bindUSDCTokenPoolProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyCaller{contract: contract}, nil
}

func NewUSDCTokenPoolProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCTokenPoolProxyTransactor, error) {
	contract, err := bindUSDCTokenPoolProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyTransactor{contract: contract}, nil
}

func NewUSDCTokenPoolProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCTokenPoolProxyFilterer, error) {
	contract, err := bindUSDCTokenPoolProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyFilterer{contract: contract}, nil
}

func bindUSDCTokenPoolProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCTokenPoolProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyCaller.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyTransactor.contract.Transfer(opts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyTransactor.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.contract.Transfer(opts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getLockOrBurnMechanism", remoteChainSelector)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetLockOrBurnMechanism(remoteChainSelector uint64) (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetLockOrBurnMechanism(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetLockOrBurnMechanism(remoteChainSelector uint64) (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetLockOrBurnMechanism(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetLockReleasePoolAddress(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getLockReleasePoolAddress", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetLockReleasePoolAddress(remoteChainSelector uint64) (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetLockReleasePoolAddress(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetLockReleasePoolAddress(remoteChainSelector uint64) (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetLockReleasePoolAddress(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getPools")

	if err != nil {
		return *new(USDCTokenPoolProxyPoolAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolProxyPoolAddresses)).(*USDCTokenPoolProxyPoolAddresses)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetPools() (USDCTokenPoolProxyPoolAddresses, error) {
	return _USDCTokenPoolProxy.Contract.GetPools(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetPools() (USDCTokenPoolProxyPoolAddresses, error) {
	return _USDCTokenPoolProxy.Contract.GetPools(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) Owner() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.Owner(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) Owner() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.Owner(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) TypeAndVersion() (string, error) {
	return _USDCTokenPoolProxy.Contract.TypeAndVersion(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) TypeAndVersion() (string, error) {
	return _USDCTokenPoolProxy.Contract.TypeAndVersion(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "acceptOwnership")
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AcceptOwnership(&_USDCTokenPoolProxy.TransactOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AcceptOwnership(&_USDCTokenPoolProxy.TransactOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.TransferOwnership(&_USDCTokenPoolProxy.TransactOpts, to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.TransferOwnership(&_USDCTokenPoolProxy.TransactOpts, to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updateLockOrBurnMechanisms", remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdateLockOrBurnMechanisms(remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockOrBurnMechanisms(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdateLockOrBurnMechanisms(remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockOrBurnMechanisms(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdateLockReleasePoolAddresses(opts *bind.TransactOpts, remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updateLockReleasePoolAddresses", remoteChainSelectors, lockReleasePools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdateLockReleasePoolAddresses(remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockReleasePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, lockReleasePools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdateLockReleasePoolAddresses(remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockReleasePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, lockReleasePools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updatePoolAddresses", pools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdatePoolAddresses(pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdatePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, pools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdatePoolAddresses(pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdatePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, pools)
}

type USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator struct {
	Event *USDCTokenPoolProxyLockOrBurnMechanismUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
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
		it.Event = new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
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

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyLockOrBurnMechanismUpdated struct {
	RemoteChainSelector uint64
	Mechanism           uint8
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "LockOrBurnMechanismUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "LockOrBurnMechanismUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "LockOrBurnMechanismUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockOrBurnMechanismUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error) {
	event := new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockOrBurnMechanismUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyLockReleasePoolUpdatedIterator struct {
	Event *USDCTokenPoolProxyLockReleasePoolUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyLockReleasePoolUpdated)
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
		it.Event = new(USDCTokenPoolProxyLockReleasePoolUpdated)
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

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyLockReleasePoolUpdated struct {
	RemoteChainSelector uint64
	LockReleasePool     common.Address
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterLockReleasePoolUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockReleasePoolUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "LockReleasePoolUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyLockReleasePoolUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "LockReleasePoolUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchLockReleasePoolUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockReleasePoolUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "LockReleasePoolUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyLockReleasePoolUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockReleasePoolUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseLockReleasePoolUpdated(log types.Log) (*USDCTokenPoolProxyLockReleasePoolUpdated, error) {
	event := new(USDCTokenPoolProxyLockReleasePoolUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockReleasePoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyOwnershipTransferRequestedIterator struct {
	Event *USDCTokenPoolProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyOwnershipTransferRequested)
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
		it.Event = new(USDCTokenPoolProxyOwnershipTransferRequested)
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

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyOwnershipTransferRequestedIterator{contract: _USDCTokenPoolProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyOwnershipTransferRequested)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolProxyOwnershipTransferRequested, error) {
	event := new(USDCTokenPoolProxyOwnershipTransferRequested)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyOwnershipTransferredIterator struct {
	Event *USDCTokenPoolProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyOwnershipTransferred)
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
		it.Event = new(USDCTokenPoolProxyOwnershipTransferred)
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

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyOwnershipTransferredIterator{contract: _USDCTokenPoolProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyOwnershipTransferred)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolProxyOwnershipTransferred, error) {
	event := new(USDCTokenPoolProxyOwnershipTransferred)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyPoolAddressesUpdatedIterator struct {
	Event *USDCTokenPoolProxyPoolAddressesUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyPoolAddressesUpdated)
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
		it.Event = new(USDCTokenPoolProxyPoolAddressesUpdated)
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

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyPoolAddressesUpdated struct {
	Pools USDCTokenPoolProxyPoolAddresses
	Raw   types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterPoolAddressesUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyPoolAddressesUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "PoolAddressesUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyPoolAddressesUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "PoolAddressesUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchPoolAddressesUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyPoolAddressesUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "PoolAddressesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyPoolAddressesUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "PoolAddressesUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParsePoolAddressesUpdated(log types.Log) (*USDCTokenPoolProxyPoolAddressesUpdated, error) {
	event := new(USDCTokenPoolProxyPoolAddressesUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "PoolAddressesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (USDCTokenPoolProxyLockOrBurnMechanismUpdated) Topic() common.Hash {
	return common.HexToHash("0x2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471")
}

func (USDCTokenPoolProxyLockReleasePoolUpdated) Topic() common.Hash {
	return common.HexToHash("0x2a4ec2a96b51064b74fa8f2157f98cf8fb9fd4dcef4e9fdf3c44c0d74d0e8264")
}

func (USDCTokenPoolProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (USDCTokenPoolProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (USDCTokenPoolProxyPoolAddressesUpdated) Topic() common.Hash {
	return common.HexToHash("0xa33a66dd32843dc04c77f6aa048e4557f085fd2d793a098e06d6e9f89c715aec")
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxy) Address() common.Address {
	return _USDCTokenPoolProxy.address
}

type USDCTokenPoolProxyInterface interface {
	GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error)

	GetLockReleasePoolAddress(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error)

	UpdateLockReleasePoolAddresses(opts *bind.TransactOpts, remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error)

	UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error)

	FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error)

	WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error)

	FilterLockReleasePoolUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockReleasePoolUpdatedIterator, error)

	WatchLockReleasePoolUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockReleasePoolUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockReleasePoolUpdated(log types.Log) (*USDCTokenPoolProxyLockReleasePoolUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolProxyOwnershipTransferred, error)

	FilterPoolAddressesUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyPoolAddressesUpdatedIterator, error)

	WatchPoolAddressesUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyPoolAddressesUpdated) (event.Subscription, error)

	ParsePoolAddressesUpdated(log types.Log) (*USDCTokenPoolProxyPoolAddressesUpdated, error)

	Address() common.Address
}
