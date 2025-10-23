// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor

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

type ClientCCV struct {
	CcvAddress common.Address
	Args       []byte
}

type ExecutorDynamicConfig struct {
	FeeAggregator         common.Address
	MinBlockConfirmations uint16
	CcvAllowlistEnabled   bool
}

type ExecutorRemoteChainConfig struct {
	UsdCentsFee            uint16
	BaseExecGas            uint32
	DestAddressLengthBytes uint8
	Enabled                bool
}

type ExecutorRemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            ExecutorRemoteChainConfig
}

var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberOfTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccvs\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"execGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"execBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxCCVsPerMessage\",\"inputs\":[],\"outputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxCCVsPerMsgSet\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Executor__RequestedBlockDepthTooLow\",\"inputs\":[{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0604052346101c957604051611e8538819003601f8101601f191683016001600160401b038111848210176101ce5783928291604052833981010390608082126101c95780519060ff8216928383036101c957606090601f1901126101c95760405191606083016001600160401b038111848210176101ce5760405260208201516001600160a01b03811681036101c957835260408201519161ffff831683036101c95760208401928352606001519384151585036101c9576040840194855233156101b857600180546001600160a01b0319163317905580156101a4575060805281516001600160a01b03161561019357905160028054835185516001600160b81b03199092166001600160a01b039490941693841760a09190911b61ffff60a01b161790151560b01b60ff60b01b1617905560408051918252915161ffff16602082015291511515908201527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a1604051611ca090816101e5823960805181818161074a01526109db0152f35b6306b7c75960e31b60005260046000fd5b631f3f959360e01b60005260045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714611535578063240b96e9146112f5578063336e545a146110ad5780634c3281c514610f1e5780635cb80c5d14610bbd5780637437ff9f14610ae857806379ba5097146109ff57806384502414146109a357806384f369ce146106365780638da5cb5b146105e4578063a68c61a6146104f6578063a8eb211e146101ee578063b8d5005e146101ab5763f2fde38b146100b657600080fd5b346101a65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65773ffffffffffffffffffffffffffffffffffffffff6101026116cb565b61010a6117ed565b1633811461017c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602061ffff60025460a01c16604051908152f35b346101a65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a65761023d90369060040161168b565b6024359167ffffffffffffffff83116101a657366023840112156101a65782600401359167ffffffffffffffff83116101a65736602460a08502860101116101a6576102876117ed565b60005b8181106104865750505060005b818110156104845760a08102830160006024820167ffffffffffffffff6102bd826117d8565b1615610446576102de67ffffffffffffffff6102d8836117d8565b16611b71565b6102ee575b505050600101610297565b67ffffffffffffffff610300826117d8565b1682526007602052604082209160448401359261ffff84169485850361044257815493606482013563ffffffff81169081810361043e5760848401359260ff84169485850361043a5760a401359586151598898803610436579267ffffffffffffffff61041160809a979460019f9e9d9a97948d9a978f7f16ff55d9c2b936d957d46720f10be4133f749fc879cf8c30bdc15b07539c56ab9f907fffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000067ff000000000000007fffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffff9360381b169316171665ffffffff00008760101b161766ff0000000000008960301b16171790556117d8565b169a6040519850885250602087015250604085015250506060820152a29084806102e3565b8880fd5b8780fd5b8580fd5b8280fd5b9067ffffffffffffffff61045b6024936117d8565b7f020a07e500000000000000000000000000000000000000000000000000000000835216600452fd5b005b8067ffffffffffffffff6104a56104a0600194868861179a565b6117d8565b166104af816119e6565b6104bb575b500161028a565b806000526007602052600060408120557ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a2866104b4565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576040518060206003549283815201809260036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b8181106105ce57505050816105759103826115e2565b6040519182916020830190602084525180915260408301919060005b81811061059f575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610591565b825484526020909301926001928301920161055f565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a65760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81168091036101a6576106826116ee565b6044359063ffffffff82168092036101a6576064359160ff83168093036101a65760843567ffffffffffffffff81116101a6576106c390369060040161168b565b9060a4359367ffffffffffffffff85116101a657366023860112156101a65784600401359467ffffffffffffffff86116101a65760248636920101116101a657866000526007602052610719604060002061175a565b96606088015115610976575061ffff1680151580610964575b610929575060ff60025460b01c16610866575b5060ff7f00000000000000000000000000000000000000000000000000000000000000001690818111610836575050604d019081604d1161080757610789916117cb565b60408301805190919060ff916101fe6107a69260011b16906117cb565b915116604d019182604d11610807578281029281840414901517156108075763ffffffff916107d4916117cb565b1663ffffffff6020830151169063ffffffff82116108075761ffff60609351169160405192835260208301526040820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b93909491926000937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc186360301945b8781101561091b5760008160051b88013587811215610917576108ce73ffffffffffffffffffffffffffffffffffffffff918a016117aa565b1680825260046020526040822054156108eb575050600101610895565b602492507fa409d83e000000000000000000000000000000000000000000000000000000008252600452fd5b5080fd5b509295919450925085610745565b61ffff60025460a01c16907f2dba20cf0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5061ffff60025460a01c168110610732565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760005473ffffffffffffffffffffffffffffffffffffffff81163303610abe577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657600060408051610b26816115c6565b8281528260208201520152610bb9604051610b40816115c6565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835261ffff8160a01c16602084015260b01c161515604082015260405191829182919091604080606083019473ffffffffffffffffffffffffffffffffffffffff815116845261ffff602082015116602085015201511515910152565b0390f35b346101a65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a657610c0c90369060040161168b565b90610c156117ed565b73ffffffffffffffffffffffffffffffffffffffff600254169160005b818110610c3b57005b73ffffffffffffffffffffffffffffffffffffffff610c63610c5e83858761179a565b6117aa565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315610f1257600093610edc575b5082610cbc575b506001915001610c32565b6040519260208401937fa9059cbb00000000000000000000000000000000000000000000000000000000855287602482015281604482015260448152610d036064826115e2565b600080604096875193610d1689866115e2565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af1883d15610ece57503d90600067ffffffffffffffff8311610ea1575091818a949360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f610db19601160191610da08a5193846115e2565b82523d6000602084013e5b86611bcb565b805180610ded575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385610cb1565b8192949596935090602091810103126101a657602001518015908115036101a657610e1e5792919086908880610db9565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b9291610db191606090610dab565b90926020823d8211610f0a575b81610ef6602093836115e2565b81010312610f075750519186610caa565b80fd5b3d9150610ee9565b6040513d6000823e3d90fd5b346101a65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576000604051610f5b816115c6565b610f636116cb565b8152610f6d6116ee565b60208201908152610f7c6116bc565b9060408301918252610f8c6117ed565b73ffffffffffffffffffffffffffffffffffffffff8351161561108557825160028054925193517fffffffffffffffffff000000000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff90921691821760a09490941b75ffff0000000000000000000000000000000000000000169390931791151560b01b76ff0000000000000000000000000000000000000000000016919091179091556040805191825260208084015161ffff1690830152918201511515918101919091527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b346101a65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a6576110fc90369060040161168b565b9060243567ffffffffffffffff81116101a65761111d90369060040161168b565b90916111276116bc565b936111306117ed565b60005b81811061126d5750505060005b8181106111d25783600254901515908160ff8260b01c1615150361116057005b816020917fffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffff76ff000000000000000000000000000000000000000000007fd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed89560b01b16911617600255604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff6111f5610c5e83858761179a565b168015611240579081611209600193611b11565b611215575b5001611140565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a28561120e565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8061129c73ffffffffffffffffffffffffffffffffffffffff611296610c5e600195878961179a565b16611850565b6112a7575b01611133565b73ffffffffffffffffffffffffffffffffffffffff6112ca610c5e83868861179a565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a26112a1565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657600554611330816116ff565b9061133e60405192836115e2565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061136b826116ff565b0160005b8181106114ac5750506005549060005b81811061140f578360405180916020820160208352815180915260206040840192019060005b8181106113b3575050500390f35b91935091602060a0600192606083885167ffffffffffffffff8151168452015161ffff8151168584015263ffffffff8582015116604084015260ff604082015116828401520151151560808201520194019101918493926113a5565b60008382101561147f5790604081602084600560019652200167ffffffffffffffff60009154166114408489611717565b515267ffffffffffffffff6114558489611717565b5151168152600760205220611477602061146f8489611717565b51019161175a565b90520161137f565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60405190604082019180831067ffffffffffffffff84111761150657602092604052600081526040516114de816115aa565b600081526000848201526000604082015260006060820152838201528282870101520161136f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657610bb9604080519061157681836115e2565b601282527f4578656375746f7220312e372e302d646576000000000000000000000000000060208301525191829182611623565b6080810190811067ffffffffffffffff82111761150657604052565b6060810190811067ffffffffffffffff82111761150657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761150657604052565b9190916020815282519283602083015260005b8481106116755750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201611636565b9181601f840112156101a65782359167ffffffffffffffff83116101a6576020808501948460051b0101116101a657565b6044359081151582036101a657565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101a657565b6024359061ffff821682036101a657565b67ffffffffffffffff81116115065760051b60200190565b805182101561172b5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90604051611767816115aa565b606060ff82945461ffff8116845263ffffffff8160101c166020850152818160301c16604085015260381c161515910152565b919081101561172b5760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101a65790565b9190820180921161080757565b3567ffffffffffffffff811681036101a65790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361180e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805482101561172b5760005260206000200190600090565b60008181526004602052604090205480156119df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161080757600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161080757818103611970575b5050506003548015611941577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016118fe816003611838565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6119c7611981611992936003611838565b90549060031b1c9283926003611838565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260046020526040600020553880806118c5565b5050600090565b60008181526006602052604090205480156119df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161080757600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161080757818103611ad7575b5050506005548015611941577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611a94816005611838565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b611af9611ae8611992936005611838565b90549060031b1c9283926005611838565b90556000526006602052604060002055388080611a5b565b80600052600460205260406000205415600014611b6b576003546801000000000000000081101561150657611b526119928260018594016003556003611838565b9055600354906000526004602052604060002055600190565b50600090565b80600052600660205260406000205415600014611b6b576005546801000000000000000081101561150657611bb26119928260018594016005556005611838565b9055600554906000526006602052604060002055600190565b91929015611c465750815115611bdf575090565b3b15611be85790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015611c595750805190602001fd5b611c8f906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301611623565b0390fdfea164736f6c634300081a000a",
}

var ExecutorABI = ExecutorMetaData.ABI

var ExecutorBin = ExecutorMetaData.Bin

func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, maxCCVsPerMsg uint8, dynamicConfig ExecutorDynamicConfig) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend, maxCCVsPerMsg, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executor{address: address, abi: *parsed, ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

type Executor struct {
	address common.Address
	abi     abi.ABI
	ExecutorCaller
	ExecutorTransactor
	ExecutorFilterer
}

type ExecutorCaller struct {
	contract *bind.BoundContract
}

type ExecutorTransactor struct {
	contract *bind.BoundContract
}

type ExecutorFilterer struct {
	contract *bind.BoundContract
}

type ExecutorSession struct {
	Contract     *Executor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ExecutorCallerSession struct {
	Contract *ExecutorCaller
	CallOpts bind.CallOpts
}

type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor
	TransactOpts bind.TransactOpts
}

type ExecutorRaw struct {
	Contract *Executor
}

type ExecutorCallerRaw struct {
	Contract *ExecutorCaller
}

type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor
}

func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	abi, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{address: address, abi: abi, ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

func (_Executor *ExecutorCaller) GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getAllowedCCVs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Executor *ExecutorSession) GetAllowedCCVs() ([]common.Address, error) {
	return _Executor.Contract.GetAllowedCCVs(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) GetAllowedCCVs() ([]common.Address, error) {
	return _Executor.Contract.GetAllowedCCVs(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) GetDestChains(opts *bind.CallOpts) ([]ExecutorRemoteChainConfigArgs, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getDestChains")

	if err != nil {
		return *new([]ExecutorRemoteChainConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]ExecutorRemoteChainConfigArgs)).(*[]ExecutorRemoteChainConfigArgs)

	return out0, err

}

func (_Executor *ExecutorSession) GetDestChains() ([]ExecutorRemoteChainConfigArgs, error) {
	return _Executor.Contract.GetDestChains(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) GetDestChains() ([]ExecutorRemoteChainConfigArgs, error) {
	return _Executor.Contract.GetDestChains(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) GetDynamicConfig(opts *bind.CallOpts) (ExecutorDynamicConfig, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(ExecutorDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ExecutorDynamicConfig)).(*ExecutorDynamicConfig)

	return out0, err

}

func (_Executor *ExecutorSession) GetDynamicConfig() (ExecutorDynamicConfig, error) {
	return _Executor.Contract.GetDynamicConfig(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) GetDynamicConfig() (ExecutorDynamicConfig, error) {
	return _Executor.Contract.GetDynamicConfig(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, dataLength uint32, numberOfTokens uint8, ccvs []ClientCCV, extraArgs []byte) (GetFee,

	error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getFee", destChainSelector, requestedBlockDepth, dataLength, numberOfTokens, ccvs, extraArgs)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.UsdCentsFee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.ExecGasCost = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ExecBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_Executor *ExecutorSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, dataLength uint32, numberOfTokens uint8, ccvs []ClientCCV, extraArgs []byte) (GetFee,

	error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, dataLength, numberOfTokens, ccvs, extraArgs)
}

func (_Executor *ExecutorCallerSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, dataLength uint32, numberOfTokens uint8, ccvs []ClientCCV, extraArgs []byte) (GetFee,

	error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, dataLength, numberOfTokens, ccvs, extraArgs)
}

func (_Executor *ExecutorCaller) GetMaxCCVsPerMessage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getMaxCCVsPerMessage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Executor *ExecutorSession) GetMaxCCVsPerMessage() (uint8, error) {
	return _Executor.Contract.GetMaxCCVsPerMessage(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) GetMaxCCVsPerMessage() (uint8, error) {
	return _Executor.Contract.GetMaxCCVsPerMessage(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getMinBlockConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Executor *ExecutorSession) GetMinBlockConfirmations() (uint16, error) {
	return _Executor.Contract.GetMinBlockConfirmations(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) GetMinBlockConfirmations() (uint16, error) {
	return _Executor.Contract.GetMinBlockConfirmations(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Executor *ExecutorSession) Owner() (common.Address, error) {
	return _Executor.Contract.Owner(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) Owner() (common.Address, error) {
	return _Executor.Contract.Owner(&_Executor.CallOpts)
}

func (_Executor *ExecutorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Executor *ExecutorSession) TypeAndVersion() (string, error) {
	return _Executor.Contract.TypeAndVersion(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) TypeAndVersion() (string, error) {
	return _Executor.Contract.TypeAndVersion(&_Executor.CallOpts)
}

func (_Executor *ExecutorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "acceptOwnership")
}

func (_Executor *ExecutorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Executor.Contract.AcceptOwnership(&_Executor.TransactOpts)
}

func (_Executor *ExecutorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Executor.Contract.AcceptOwnership(&_Executor.TransactOpts)
}

func (_Executor *ExecutorTransactor) ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "applyAllowedCCVUpdates", ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_Executor *ExecutorSession) ApplyAllowedCCVUpdates(ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _Executor.Contract.ApplyAllowedCCVUpdates(&_Executor.TransactOpts, ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_Executor *ExecutorTransactorSession) ApplyAllowedCCVUpdates(ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _Executor.Contract.ApplyAllowedCCVUpdates(&_Executor.TransactOpts, ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_Executor *ExecutorTransactor) ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []ExecutorRemoteChainConfigArgs) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "applyDestChainUpdates", destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_Executor *ExecutorSession) ApplyDestChainUpdates(destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []ExecutorRemoteChainConfigArgs) (*types.Transaction, error) {
	return _Executor.Contract.ApplyDestChainUpdates(&_Executor.TransactOpts, destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_Executor *ExecutorTransactorSession) ApplyDestChainUpdates(destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []ExecutorRemoteChainConfigArgs) (*types.Transaction, error) {
	return _Executor.Contract.ApplyDestChainUpdates(&_Executor.TransactOpts, destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_Executor *ExecutorTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorDynamicConfig) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_Executor *ExecutorSession) SetDynamicConfig(dynamicConfig ExecutorDynamicConfig) (*types.Transaction, error) {
	return _Executor.Contract.SetDynamicConfig(&_Executor.TransactOpts, dynamicConfig)
}

func (_Executor *ExecutorTransactorSession) SetDynamicConfig(dynamicConfig ExecutorDynamicConfig) (*types.Transaction, error) {
	return _Executor.Contract.SetDynamicConfig(&_Executor.TransactOpts, dynamicConfig)
}

func (_Executor *ExecutorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "transferOwnership", to)
}

func (_Executor *ExecutorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Executor.Contract.TransferOwnership(&_Executor.TransactOpts, to)
}

func (_Executor *ExecutorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Executor.Contract.TransferOwnership(&_Executor.TransactOpts, to)
}

func (_Executor *ExecutorTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_Executor *ExecutorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _Executor.Contract.WithdrawFeeTokens(&_Executor.TransactOpts, feeTokens)
}

func (_Executor *ExecutorTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _Executor.Contract.WithdrawFeeTokens(&_Executor.TransactOpts, feeTokens)
}

type ExecutorCCVAddedIterator struct {
	Event *ExecutorCCVAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorCCVAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorCCVAdded)
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
		it.Event = new(ExecutorCCVAdded)
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

func (it *ExecutorCCVAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorCCVAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorCCVAdded struct {
	Ccv common.Address
	Raw types.Log
}

func (_Executor *ExecutorFilterer) FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVAddedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorCCVAddedIterator{contract: _Executor.contract, event: "CCVAdded", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAdded, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorCCVAdded)
				if err := _Executor.contract.UnpackLog(event, "CCVAdded", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseCCVAdded(log types.Log) (*ExecutorCCVAdded, error) {
	event := new(ExecutorCCVAdded)
	if err := _Executor.contract.UnpackLog(event, "CCVAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorCCVAllowlistUpdatedIterator struct {
	Event *ExecutorCCVAllowlistUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorCCVAllowlistUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorCCVAllowlistUpdated)
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
		it.Event = new(ExecutorCCVAllowlistUpdated)
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

func (it *ExecutorCCVAllowlistUpdatedIterator) Error() error {
	return it.fail
}

func (it *ExecutorCCVAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorCCVAllowlistUpdated struct {
	Enabled bool
	Raw     types.Log
}

func (_Executor *ExecutorFilterer) FilterCCVAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorCCVAllowlistUpdatedIterator, error) {

	logs, sub, err := _Executor.contract.FilterLogs(opts, "CCVAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &ExecutorCCVAllowlistUpdatedIterator{contract: _Executor.contract, event: "CCVAllowlistUpdated", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchCCVAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _Executor.contract.WatchLogs(opts, "CCVAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorCCVAllowlistUpdated)
				if err := _Executor.contract.UnpackLog(event, "CCVAllowlistUpdated", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseCCVAllowlistUpdated(log types.Log) (*ExecutorCCVAllowlistUpdated, error) {
	event := new(ExecutorCCVAllowlistUpdated)
	if err := _Executor.contract.UnpackLog(event, "CCVAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorCCVRemovedIterator struct {
	Event *ExecutorCCVRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorCCVRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorCCVRemoved)
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
		it.Event = new(ExecutorCCVRemoved)
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

func (it *ExecutorCCVRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorCCVRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorCCVRemoved struct {
	Ccv common.Address
	Raw types.Log
}

func (_Executor *ExecutorFilterer) FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVRemovedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorCCVRemovedIterator{contract: _Executor.contract, event: "CCVRemoved", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorCCVRemoved, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorCCVRemoved)
				if err := _Executor.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseCCVRemoved(log types.Log) (*ExecutorCCVRemoved, error) {
	event := new(ExecutorCCVRemoved)
	if err := _Executor.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorConfigSetIterator struct {
	Event *ExecutorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorConfigSet)
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
		it.Event = new(ExecutorConfigSet)
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

func (it *ExecutorConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorConfigSet struct {
	DynamicConfig ExecutorDynamicConfig
	Raw           types.Log
}

func (_Executor *ExecutorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ExecutorConfigSetIterator, error) {

	logs, sub, err := _Executor.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorConfigSetIterator{contract: _Executor.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorConfigSet) (event.Subscription, error) {

	logs, sub, err := _Executor.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorConfigSet)
				if err := _Executor.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseConfigSet(log types.Log) (*ExecutorConfigSet, error) {
	event := new(ExecutorConfigSet)
	if err := _Executor.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorDestChainAddedIterator struct {
	Event *ExecutorDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorDestChainAdded)
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
		it.Event = new(ExecutorDestChainAdded)
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

func (it *ExecutorDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorDestChainAdded struct {
	DestChainSelector uint64
	Config            ExecutorRemoteChainConfig
	Raw               types.Log
}

func (_Executor *ExecutorFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorDestChainAddedIterator{contract: _Executor.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorDestChainAdded)
				if err := _Executor.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseDestChainAdded(log types.Log) (*ExecutorDestChainAdded, error) {
	event := new(ExecutorDestChainAdded)
	if err := _Executor.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorDestChainRemovedIterator struct {
	Event *ExecutorDestChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorDestChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorDestChainRemoved)
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
		it.Event = new(ExecutorDestChainRemoved)
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

func (it *ExecutorDestChainRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorDestChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorDestChainRemoved struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_Executor *ExecutorFilterer) FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorDestChainRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorDestChainRemovedIterator{contract: _Executor.contract, event: "DestChainRemoved", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorDestChainRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorDestChainRemoved)
				if err := _Executor.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseDestChainRemoved(log types.Log) (*ExecutorDestChainRemoved, error) {
	event := new(ExecutorDestChainRemoved)
	if err := _Executor.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorFeeTokenWithdrawnIterator struct {
	Event *ExecutorFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorFeeTokenWithdrawn)
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
		it.Event = new(ExecutorFeeTokenWithdrawn)
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

func (it *ExecutorFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *ExecutorFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_Executor *ExecutorFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*ExecutorFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorFeeTokenWithdrawnIterator{contract: _Executor.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *ExecutorFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorFeeTokenWithdrawn)
				if err := _Executor.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseFeeTokenWithdrawn(log types.Log) (*ExecutorFeeTokenWithdrawn, error) {
	event := new(ExecutorFeeTokenWithdrawn)
	if err := _Executor.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorMaxCCVsPerMsgSetIterator struct {
	Event *ExecutorMaxCCVsPerMsgSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorMaxCCVsPerMsgSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorMaxCCVsPerMsgSet)
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
		it.Event = new(ExecutorMaxCCVsPerMsgSet)
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

func (it *ExecutorMaxCCVsPerMsgSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorMaxCCVsPerMsgSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorMaxCCVsPerMsgSet struct {
	MaxCCVsPerMsg uint8
	Raw           types.Log
}

func (_Executor *ExecutorFilterer) FilterMaxCCVsPerMsgSet(opts *bind.FilterOpts) (*ExecutorMaxCCVsPerMsgSetIterator, error) {

	logs, sub, err := _Executor.contract.FilterLogs(opts, "MaxCCVsPerMsgSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorMaxCCVsPerMsgSetIterator{contract: _Executor.contract, event: "MaxCCVsPerMsgSet", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchMaxCCVsPerMsgSet(opts *bind.WatchOpts, sink chan<- *ExecutorMaxCCVsPerMsgSet) (event.Subscription, error) {

	logs, sub, err := _Executor.contract.WatchLogs(opts, "MaxCCVsPerMsgSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorMaxCCVsPerMsgSet)
				if err := _Executor.contract.UnpackLog(event, "MaxCCVsPerMsgSet", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseMaxCCVsPerMsgSet(log types.Log) (*ExecutorMaxCCVsPerMsgSet, error) {
	event := new(ExecutorMaxCCVsPerMsgSet)
	if err := _Executor.contract.UnpackLog(event, "MaxCCVsPerMsgSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOwnershipTransferRequestedIterator struct {
	Event *ExecutorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOwnershipTransferRequested)
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
		it.Event = new(ExecutorOwnershipTransferRequested)
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

func (it *ExecutorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Executor *ExecutorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOwnershipTransferRequestedIterator{contract: _Executor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOwnershipTransferRequested)
				if err := _Executor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ExecutorOwnershipTransferRequested, error) {
	event := new(ExecutorOwnershipTransferRequested)
	if err := _Executor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOwnershipTransferredIterator struct {
	Event *ExecutorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOwnershipTransferred)
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
		it.Event = new(ExecutorOwnershipTransferred)
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

func (it *ExecutorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExecutorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Executor *ExecutorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOwnershipTransferredIterator{contract: _Executor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOwnershipTransferred)
				if err := _Executor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutorOwnershipTransferred, error) {
	event := new(ExecutorOwnershipTransferred)
	if err := _Executor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetFee struct {
	UsdCentsFee uint16
	ExecGasCost uint32
	ExecBytes   uint32
}

func (ExecutorCCVAdded) Topic() common.Hash {
	return common.HexToHash("0xba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e")
}

func (ExecutorCCVAllowlistUpdated) Topic() common.Hash {
	return common.HexToHash("0xd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed8")
}

func (ExecutorCCVRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e")
}

func (ExecutorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b65892")
}

func (ExecutorDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x16ff55d9c2b936d957d46720f10be4133f749fc879cf8c30bdc15b07539c56ab")
}

func (ExecutorDestChainRemoved) Topic() common.Hash {
	return common.HexToHash("0xf74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1")
}

func (ExecutorFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (ExecutorMaxCCVsPerMsgSet) Topic() common.Hash {
	return common.HexToHash("0xcd39dd44d856487a5d3ff100b17da01d09fd38f56a6bc6c1430458ec9cd31bd8")
}

func (ExecutorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ExecutorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_Executor *Executor) Address() common.Address {
	return _Executor.address
}

type ExecutorInterface interface {
	GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChains(opts *bind.CallOpts) ([]ExecutorRemoteChainConfigArgs, error)

	GetDynamicConfig(opts *bind.CallOpts) (ExecutorDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, dataLength uint32, numberOfTokens uint8, ccvs []ClientCCV, extraArgs []byte) (GetFee,

		error)

	GetMaxCCVsPerMessage(opts *bind.CallOpts) (uint8, error)

	GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error)

	ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []ExecutorRemoteChainConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVAddedIterator, error)

	WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAdded, ccv []common.Address) (event.Subscription, error)

	ParseCCVAdded(log types.Log) (*ExecutorCCVAdded, error)

	FilterCCVAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorCCVAllowlistUpdatedIterator, error)

	WatchCCVAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAllowlistUpdated) (event.Subscription, error)

	ParseCCVAllowlistUpdated(log types.Log) (*ExecutorCCVAllowlistUpdated, error)

	FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVRemovedIterator, error)

	WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorCCVRemoved, ccv []common.Address) (event.Subscription, error)

	ParseCCVRemoved(log types.Log) (*ExecutorCCVRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*ExecutorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*ExecutorConfigSet, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*ExecutorDestChainAdded, error)

	FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorDestChainRemovedIterator, error)

	WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorDestChainRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainRemoved(log types.Log) (*ExecutorDestChainRemoved, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*ExecutorFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *ExecutorFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*ExecutorFeeTokenWithdrawn, error)

	FilterMaxCCVsPerMsgSet(opts *bind.FilterOpts) (*ExecutorMaxCCVsPerMsgSetIterator, error)

	WatchMaxCCVsPerMsgSet(opts *bind.WatchOpts, sink chan<- *ExecutorMaxCCVsPerMsgSet) (event.Subscription, error)

	ParseMaxCCVsPerMsgSet(log types.Log) (*ExecutorMaxCCVsPerMsgSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOwnershipTransferred, error)

	Address() common.Address
}
