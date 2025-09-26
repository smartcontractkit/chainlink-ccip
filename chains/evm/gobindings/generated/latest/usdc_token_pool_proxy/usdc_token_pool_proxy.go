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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockReleasePoolAddress\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateLockReleasePoolAddresses\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"lockReleasePools\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleasePoolUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestinationPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c0604052346101d75760405161284b38819003601f8101601f191683016001600160401b038111848210176101c1578392829160405283398101039060a082126101d75780516001600160a01b038116928382036101d757606090601f1901126101d757604051606081016001600160401b038111828210176101c15760405261008c602084016101dc565b815261009a604084016101dc565b91602082019283526100c160806100b3606087016101dc565b9560408501968752016101dc565b9433156101b057600180546001600160a01b0319163317905515801561019e575b801561018c575b801561017b575b61016a5760805251600480546001600160a01b03199081166001600160a01b039384161790915591516005805484169183169190911790559151600680549092169083161790551660a05260405161265a90816101f1823960805181818161051f0152610ccd015260a051818181610bd101526117290152f35b6303988b8160e61b60005260046000fd5b506001600160a01b038516156100f0565b5083516001600160a01b0316156100e9565b5082516001600160a01b0316156100e2565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b51906001600160a01b03821682036101d75756fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610107578063181f5a77146101025780631c255a9d146100fd578063240028e8146100f857806339077537146100f3578063673a2a1f146100ee57806379ba5097146100e95780638302d09f146100e45780638926f54f146100df5780638da5cb5b146100da5780639a4575b9146100d5578063aa86a754146100d0578063c527a6c2146100cb578063db4c2aef146100c65763f2fde38b146100c157600080fd5b611171565b61109c565b610f3f565b610eda565b610aed565b610a48565b6109ad565b6107dc565b610652565b6105c9565b610544565b6104b1565b610415565b610377565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c6576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036101c657807f0e64dd29000000000000000000000000000000000000000000000000000000006020921490811561019c575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610191565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761021657604052565b6101cb565b6020810190811067ffffffffffffffff82111761021657604052565b6040810190811067ffffffffffffffff82111761021657604052565b610100810190811067ffffffffffffffff82111761021657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761021657604052565b604051906102c161010083610270565b565b67ffffffffffffffff811161021657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106103105750506000910152565b8181015183820152602001610300565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361035c815180928187528780880191016102fd565b0116010190565b906020610374928181520190610320565b90565b346101c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c6576103f460408051906103b88183610270565b601c82527f55534443546f6b656e506f6f6c50726f787920312e362e332d64657600000000602083015251918291602083526020830190610320565b0390f35b67ffffffffffffffff8116036101c657565b35906102c1826103f8565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65767ffffffffffffffff600435610459816103f8565b166000526003602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b73ffffffffffffffffffffffffffffffffffffffff8116036101c657565b35906102c182610488565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65760206004356104ee81610488565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65760043567ffffffffffffffff81116101c6576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c6576105c06020916004016116ab565b60405190518152f35b346101c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c657600060408051610607816101fa565b8281528260208201520152606061061c611a89565b73ffffffffffffffffffffffffffffffffffffffff60408051928281511684528260208201511660208501520151166040820152f35b346101c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65760005473ffffffffffffffffffffffffffffffffffffffff81163303610711577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101c65782359167ffffffffffffffff83116101c6576020808501948460051b0101116101c657565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c65760043567ffffffffffffffff81116101c657816107b59160040161073b565b929092916024359067ffffffffffffffff82116101c6576107d89160040161073b565b9091565b346101c6576107ea3661076c565b926107f3612108565b8383036109835760005b83811061080657005b61083561081c610817838887611af0565b611b53565b73ffffffffffffffffffffffffffffffffffffffff1690565b151580610960575b61093657806108c56108556108176001948988611af0565b61088561086b610866858a89611af0565b611278565b67ffffffffffffffff166000526003602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6108d3610866828786611af0565b7f2a4ec2a96b51064b74fa8f2157f98cf8fb9fd4dcef4e9fdf3c44c0d74d0e826467ffffffffffffffff61090b610817858b8a611af0565b60405173ffffffffffffffffffffffffffffffffffffffff919091168152921691602090a2016107fd565b7f8c909bc20000000000000000000000000000000000000000000000000000000060005260046000fd5b5061097e61097a610975610817848988611af0565b612153565b1590565b61083d565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65767ffffffffffffffff6004356109f1816103f8565b16600052600260205260ff604060002054166004811015610a19576040519015158152602090f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346101c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9061037491602081526020610aba83516040838501526060840190610320565b9201519060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610320565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65760043567ffffffffffffffff81116101c65760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c657610b62611b5d565b5060248101610bb86020610b7583611278565b6040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610d6f57600091610e7c575b5073ffffffffffffffffffffffffffffffffffffffff33911603610e5257610c50610c49610c2f83611278565b67ffffffffffffffff166000526002602052604060002090565b5460ff1690565b610c5981610eab565b8015610e2157610c67611a89565b91600091610c7481610eab565b60028103610d9e575050506040015173ffffffffffffffffffffffffffffffffffffffff16905b73ffffffffffffffffffffffffffffffffffffffff8216918215610d74576000928392610cf1610d29936064830135907f000000000000000000000000000000000000000000000000000000000000000061226a565b6040519485809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260040160048301611c46565b03925af18015610d6f576103f491600091610d4c575b5060405191829182610a9a565b610d6991503d806000833e610d618183610270565b810190611bcd565b38610d3f565b61129a565b7fb348dbbe0000000000000000000000000000000000000000000000000000000060005260046000fd5b610da781610eab565b60018103610dd2575050506020015173ffffffffffffffffffffffffffffffffffffffff1690610c9b565b60039192949350610de281610eab565b14610dee575b50610c9b565b610e1a91925061086b610e0091611278565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9038610de8565b610e4e907f31603b1200000000000000000000000000000000000000000000000000000000600052610eb5565b6000fd5b7f82b429000000000000000000000000000000000000000000000000000000000060005260046000fd5b610e9e915060203d602011610ea4575b610e968183610270565b810190611b76565b38610c02565b503d610e8c565b60041115610a1957565b906024916004811015610a1957600452565b919060208301926004821015610a195752565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65767ffffffffffffffff600435610f1e816103f8565b1660005260026020526103f460ff6040600020541660405191829182610ec7565b346101c65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c6576000610f78612108565b73ffffffffffffffffffffffffffffffffffffffff610f95611b2f565b1615801561108a575b61106257610fb061097a610975611b2f565b801561104e575b8015611026575b610ffe57610fca611cce565b7fa33a66dd32843dc04c77f6aa048e4557f085fd2d793a098e06d6e9f89c715aec60405180610ff881611db1565b0390a180f35b807f8c909bc20000000000000000000000000000000000000000000000000000000060049252fd5b5061103261081c611b47565b15158015610fbe575061104961097a610975611b47565b610fbe565b5061105d61097a610975611b3b565b610fb7565b807fe622e0400000000000000000000000000000000000000000000000000000000060049252fd5b5061109661081c611b3b565b15610f9e565b346101c6576110aa3661076c565b926110b3612108565b8383036109835760005b8381106110c657005b806111106110df6110da6001948988611af0565b611e2b565b67ffffffffffffffff6110f3848988611af0565b356110fd816103f8565b1660005260026020526040600020611e38565b61111e610866828786611af0565b7f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc9447167ffffffffffffffff6111566110da858b8a611af0565b92611168604051928392169482610ec7565b0390a2016110bd565b346101c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c65773ffffffffffffffffffffffffffffffffffffffff6004356111c181610488565b6111c9612108565b1633811461123b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906112728261021b565b60008252565b35610374816103f8565b908160209103126101c6575180151581036101c65790565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c6570180359067ffffffffffffffff82116101c6576020019181360383136101c657565b908160209103126101c6576040519061130f8261021b565b51815290565b90610374916020815260e061140a6113d761133e85516101006020870152610120860190610320565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526113a3608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c0870152610320565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610320565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610320565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101c657016020813591019167ffffffffffffffff82116101c65781360383136101c657565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9061037491602081526116066115fb6115be6114ff6114ec868061143e565b610100602088015261012087019161148e565b61151f61150e6020880161040a565b67ffffffffffffffff166040870152565b61154b61152e604088016104a6565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b60608601356080860152611581611564608088016104a6565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b61158e60a087018761143e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c088015261148e565b6115cb60c086018661143e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e087015261148e565b9260e081019061143e565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015261148e565b906004116101c65790600490565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110611679575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b6116b3611265565b50602081019061171060206116c784611278565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015233602482015291829081906044820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610d6f57600091611a5a575b5015610e525760c08101604061177082846112a6565b9050146119a65761178d61178761179392846112a6565b90611637565b90611645565b917fffffffff000000000000000000000000000000000000000000000000000000008316907ffa7c07de00000000000000000000000000000000000000000000000000000000821461198557507ff3567d18000000000000000000000000000000000000000000000000000000008114611955577fb148ea5f00000000000000000000000000000000000000000000000000000000811490811561192b575b50611887577fcacdaf2b000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000821660045260246000fd5b600091506118ec6020916118b661081c61081c60065473ffffffffffffffffffffffffffffffffffffffff1690565b906040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016114cd565b03925af1908115610d6f57600091611902575090565b610374915060203d602011611924575b61191c8183610270565b8101906112f7565b503d611912565b7f0458016c0000000000000000000000000000000000000000000000000000000091501438611832565b50600091506118ec6020916118b661081c61081c60055473ffffffffffffffffffffffffffffffffffffffff1690565b600093506118ec9150916118b661081c61081c610e0061086b602097611278565b5060009150806119c46119be60e060209401836112a6565b90611eb6565b83146119f3576118ec906118b661081c61081c60045473ffffffffffffffffffffffffffffffffffffffff1690565b6118ec90611a25611a1f61081c61081c60055473ffffffffffffffffffffffffffffffffffffffff1690565b91611fc3565b6040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301611315565b611a7c915060203d602011611a82575b611a748183610270565b810190611282565b3861175a565b503d611a6a565b60405190611a96826101fa565b8173ffffffffffffffffffffffffffffffffffffffff60045416815273ffffffffffffffffffffffffffffffffffffffff600554166020820152604073ffffffffffffffffffffffffffffffffffffffff60065416910152565b9190811015611b005760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60243561037481610488565b60443561037481610488565b60043561037481610488565b3561037481610488565b60405190611b6a82610237565b60606020838281520152565b908160209103126101c6575161037481610488565b81601f820112156101c6578051611ba1816102c3565b92611baf6040519485610270565b818452602082840101116101c65761037491602080850191016102fd565b6020818303126101c65780519067ffffffffffffffff82116101c657016040818303126101c65760405191611c0183610237565b815167ffffffffffffffff81116101c65781611c1e918401611b8b565b8352602082015167ffffffffffffffff81116101c657611c3e9201611b8b565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080611c80611c70868061143e565b85602088015260c087019161148e565b9467ffffffffffffffff6020820135611c98816103f8565b166040860152826040820135611cad81610488565b1660608601526060810135828601520135611cc781610488565b1691015290565b73ffffffffffffffffffffffffffffffffffffffff600435611cef81610488565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045573ffffffffffffffffffffffffffffffffffffffff602435611d3a81610488565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055573ffffffffffffffffffffffffffffffffffffffff604435611d8581610488565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655565b90606082019173ffffffffffffffffffffffffffffffffffffffff600435611dd881610488565b16815273ffffffffffffffffffffffffffffffffffffffff602435611dfc81610488565b166020820152604073ffffffffffffffffffffffffffffffffffffffff604435611e2581610488565b16910152565b3560048110156101c65790565b906004811015610a195760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b81601f820112156101c657803590611e86826102c3565b92611e946040519485610270565b828452602083830101116101c657816000926020809301838601378301015290565b73ffffffffffffffffffffffffffffffffffffffff60045416918215611f7d578101906020818303126101c65780359067ffffffffffffffff82116101c657016040818303126101c65760405191611f0d83610237565b813567ffffffffffffffff81116101c65781611f2a918401611e6f565b8352602082013567ffffffffffffffff81116101c657611f7993611f5861081c9360749361081c9601611e6f565b602082015251015173ffffffffffffffffffffffffffffffffffffffff1690565b1490565b505050600090565b908160409103126101c657602060405191611f9f83610237565b8035611faa816103f8565b8352013563ffffffff811681036101c657602082015290565b606060e0604051611fd381610253565b8281526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152610100813603126101c6576120126102b1565b90803567ffffffffffffffff81116101c6576120319036908301611e6f565b825261203f6020820161040a565b6020830152612050604082016104a6565b60408301526060810135606083015261206b608082016104a6565b608083015260a081013567ffffffffffffffff81116101c6576120919036908301611e6f565b60a083015260c08101803567ffffffffffffffff81116101c6576120b89036908401611e6f565b9160c0840192835260e081013567ffffffffffffffff81116101c657612103926120fe926120ec6120f69336908301611e6f565b60e08801526112a6565b810190611f85565b61236a565b905290565b73ffffffffffffffffffffffffffffffffffffffff60015416330361212957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000007f0e64dd2900000000000000000000000000000000000000000000000000000000166024820152602481526121d5604482610270565b6179185a1061224057602091600091519084617530fa903d6000519083612234575b508261222a575b5081612218575b8161220e575090565b610374915061249a565b9050612223816123fe565b1590612205565b15159150386121fe565b602011159250386121f7565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b9073ffffffffffffffffffffffffffffffffffffffff61233c9392604051938260208601947fa9059cbb0000000000000000000000000000000000000000000000000000000086521660248601526044850152604484526122cc606485610270565b166000806040938451956122e08688610270565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15612361573d61232d612324826102c3565b94519485610270565b83523d6000602085013e612585565b805180612347575050565b8160208061235c936102c19501019101611282565b6124fa565b60609250612585565b7fffffffff00000000000000000000000000000000000000000000000000000000602082519201517fffffffffffffffff000000000000000000000000000000000000000000000000604051937ff3567d1800000000000000000000000000000000000000000000000000000000602086015260c01b16602484015260e01b16602c82015260108152610374603082610270565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff0000000000000000000000000000000000000000000000000000000060248301526024825261245e604483610270565b6179185a10612240576020926000925191617530fa6000513d8261248e575b5081612487575090565b9050151590565b6020111591503861247d565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a70000000000000000000000000000000000000000000000000000000060248301526024825261245e604483610270565b1561250157565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b919290156126005750815115612599575090565b3b156125a25790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156126135750805190602001fd5b612649906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610363565b0390fdfea164736f6c634300081a000a",
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedChain(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedChain(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedToken(&_USDCTokenPoolProxy.CallOpts, token)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedToken(&_USDCTokenPoolProxy.CallOpts, token)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.SupportsInterface(&_USDCTokenPoolProxy.CallOpts, interfaceId)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.SupportsInterface(&_USDCTokenPoolProxy.CallOpts, interfaceId)
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

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

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
