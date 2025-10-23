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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberOfTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccvs\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"execGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"execBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxCCVsPerMessage\",\"inputs\":[],\"outputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCCVAllowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destAddressLengthBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxCCVsPerMsgSet\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationsSet\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Executor__RequestedBlockDepthTooLow\",\"inputs\":[{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0346100dc57601f611be638819003918201601f19168301916001600160401b038311848410176100e15780849260409485528339810103126100dc57805160ff8116918282036100dc57602001519061ffff821682036100dc5733156100cb576001549280156100b75750608052600162ffff0160a01b03199091163361ffff60a81b19161760a89190911b61ffff60a81b1617600155604051611aee90816100f88239608051818181610bb90152610e4a0152f35b631f3f959360e01b60005260045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a77146113de578063240b96e91461119e578063336e545a14610f5757806379ba509714610e6e5780638450241414610e1257806384f369ce14610aa25780638da5cb5b14610a505780639dd5072314610a0c578063a68c61a61461091e578063a8eb211e14610616578063b8d5005e146105d3578063c7230a601461026a578063d4d6de23146101ae5763f2fde38b146100b657600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043573ffffffffffffffffffffffffffffffffffffffff81168091036101a95761010e61163b565b33811461017f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043561ffff8116908181036101a9577f46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a019160209161021a61163b565b7fffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffff76ffff0000000000000000000000000000000000000000006001549260a81b16911617600155604051908152a1005b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a9576102b990369060040161151c565b906024359173ffffffffffffffffffffffffffffffffffffffff83168093036101a9576102e461163b565b60005b8181106102f057005b73ffffffffffffffffffffffffffffffffffffffff6103186103138385876115e8565b6115f8565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa9283156105c757600093610591575b5082610371575b5060019150016102e7565b6040519260208401937fa9059cbb000000000000000000000000000000000000000000000000000000008552876024820152816044820152604481526103b8606482611473565b6000806040968751936103cb8986611473565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af1883d1561058357503d90600067ffffffffffffffff8311610556575091818a949360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f61046696011601916104558a519384611473565b82523d6000602084013e5b86611a19565b8051806104a2575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385610366565b8192949596935090602091810103126101a957602001518015908115036101a9576104d3579291908690888061046e565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b929161046691606090610460565b90926020823d82116105bf575b816105ab60209383611473565b810103126105bc575051918661035f565b80fd5b3d915061059e565b6040513d6000823e3d90fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602061ffff60015460a81c16604051908152f35b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a95761066590369060040161151c565b6024359167ffffffffffffffff83116101a957366023840112156101a95782600401359167ffffffffffffffff83116101a95736602460a08502860101116101a9576106af61163b565b60005b8181106108ae5750505060005b818110156108ac5760a08102830160006024820167ffffffffffffffff6106e582611626565b161561086e5761070667ffffffffffffffff61070083611626565b166119bf565b610716575b5050506001016106bf565b67ffffffffffffffff61072882611626565b1682526006602052604082209160448401359261ffff84169485850361086a57815493606482013563ffffffff8116908181036108665760848401359260ff8416948585036108625760a40135958615159889880361085e579267ffffffffffffffff61083960809a979460019f9e9d9a97948d9a978f7f16ff55d9c2b936d957d46720f10be4133f749fc879cf8c30bdc15b07539c56ab9f907fffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000067ff000000000000007fffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffff9360381b169316171665ffffffff00008760101b161766ff0000000000008960301b1617179055611626565b169a6040519850885250602087015250604085015250506060820152a290848061070b565b8880fd5b8780fd5b8580fd5b8280fd5b9067ffffffffffffffff610883602493611626565b7f020a07e500000000000000000000000000000000000000000000000000000000835216600452fd5b005b8067ffffffffffffffff6108cd6108c860019486886115e8565b611626565b166108d781611834565b6108e3575b50016106b2565b806000526006602052600060408120557ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a2866108dc565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106109f6575050508161099d910382611473565b6040519182916020830190602084525180915260408301919060005b8181106109c7575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016109b9565b8254845260209093019260019283019201610987565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602060ff60015460a01c166040519015158152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a95760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81168091036101a95760243561ffff81168091036101a9576044359063ffffffff82168092036101a9576064359160ff83168093036101a95760843567ffffffffffffffff81116101a957610b3690369060040161151c565b9060a4359367ffffffffffffffff85116101a957366023860112156101a95784600401359467ffffffffffffffff86116101a95760248636920101116101a957866000526006602052610b8c60406000206115a8565b96606088015115610de5575080151580610dd3575b610d98575060ff60015460a01c16610cd5575b5060ff7f00000000000000000000000000000000000000000000000000000000000000001690818111610ca5575050604d019081604d11610c7657610bf891611619565b60408301805190919060ff916101fe610c159260011b1690611619565b915116604d019182604d11610c7657828102928184041490151715610c765763ffffffff91610c4391611619565b1663ffffffff6020830151169063ffffffff8211610c765761ffff60609351169160405192835260208301526040820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b93909491926000937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc186360301945b87811015610d8a5760008160051b88013587811215610d8657610d3d73ffffffffffffffffffffffffffffffffffffffff918a016115f8565b168082526003602052604082205415610d5a575050600101610d04565b602492507fa409d83e000000000000000000000000000000000000000000000000000000008252600452fd5b5080fd5b509295919450925085610bb4565b61ffff60015460a81c16907f2dba20cf0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5061ffff60015460a81c168110610ba1565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760005473ffffffffffffffffffffffffffffffffffffffff81163303610f2d577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957610fa690369060040161151c565b9060243567ffffffffffffffff81116101a957610fc790369060040161151c565b9091604435938415158095036101a957610fdf61163b565b60005b8181106111165750505060005b81811061107b57836001548160ff8260a01c1615150361100b57005b816020917fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000007fd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed89560a01b16911617600155604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff61109e6103138385876115e8565b1680156110e95790816110b260019361195f565b6110be575b5001610fef565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a2856110b7565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8061114573ffffffffffffffffffffffffffffffffffffffff61113f61031360019587896115e8565b1661169e565b611150575b01610fe2565b73ffffffffffffffffffffffffffffffffffffffff6111736103138386886115e8565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a261114a565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576004546111d98161154d565b906111e76040519283611473565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06112148261154d565b0160005b8181106113555750506004549060005b8181106112b8578360405180916020820160208352815180915260206040840192019060005b81811061125c575050500390f35b91935091602060a0600192606083885167ffffffffffffffff8151168452015161ffff8151168584015263ffffffff8582015116604084015260ff6040820151168284015201511515608082015201940191019184939261124e565b6000838210156113285790604081602084600460019652200167ffffffffffffffff60009154166112e98489611565565b515267ffffffffffffffff6112fe8489611565565b515116815260066020522061132060206113188489611565565b5101916115a8565b905201611228565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60405190604082019180831067ffffffffffffffff8411176113af576020926040526000815260405161138781611457565b6000815260008482015260006040820152600060608201528382015282828701015201611218565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957611453604080519061141f8183611473565b601282527f4578656375746f7220312e372e302d6465760000000000000000000000000000602083015251918291826114b4565b0390f35b6080810190811067ffffffffffffffff8211176113af57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176113af57604052565b9190916020815282519283602083015260005b8481106115065750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b80602080928401015160408286010152016114c7565b9181601f840112156101a95782359167ffffffffffffffff83116101a9576020808501948460051b0101116101a957565b67ffffffffffffffff81116113af5760051b60200190565b80518210156115795760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906040516115b581611457565b606060ff82945461ffff8116845263ffffffff8160101c166020850152818160301c16604085015260381c161515910152565b91908110156115795760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101a95790565b91908201809211610c7657565b3567ffffffffffffffff811681036101a95790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361165c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156115795760005260206000200190600090565b600081815260036020526040902054801561182d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610c7657600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610c76578181036117be575b505050600254801561178f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161174c816002611686565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6118156117cf6117e0936002611686565b90549060031b1c9283926002611686565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080611713565b5050600090565b600081815260056020526040902054801561182d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610c7657600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610c7657818103611925575b505050600454801561178f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016118e2816004611686565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b6119476119366117e0936004611686565b90549060031b1c9283926004611686565b905560005260056020526040600020553880806118a9565b806000526003602052604060002054156000146119b957600254680100000000000000008110156113af576119a06117e08260018594016002556002611686565b9055600254906000526003602052604060002055600190565b50600090565b806000526005602052604060002054156000146119b957600454680100000000000000008110156113af57611a006117e08260018594016004556004611686565b9055600454906000526005602052604060002055600190565b91929015611a945750815115611a2d575090565b3b15611a365790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015611aa75750805190602001fd5b611add906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016114b4565b0390fdfea164736f6c634300081a000a",
}

var ExecutorABI = ExecutorMetaData.ABI

var ExecutorBin = ExecutorMetaData.Bin

func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, maxCCVsPerMsg uint8, minBlockConfirmations uint16) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend, maxCCVsPerMsg, minBlockConfirmations)
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

func (_Executor *ExecutorCaller) IsCCVAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "isCCVAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Executor *ExecutorSession) IsCCVAllowlistEnabled() (bool, error) {
	return _Executor.Contract.IsCCVAllowlistEnabled(&_Executor.CallOpts)
}

func (_Executor *ExecutorCallerSession) IsCCVAllowlistEnabled() (bool, error) {
	return _Executor.Contract.IsCCVAllowlistEnabled(&_Executor.CallOpts)
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

func (_Executor *ExecutorTransactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_Executor *ExecutorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _Executor.Contract.SetMinBlockConfirmations(&_Executor.TransactOpts, minBlockConfirmations)
}

func (_Executor *ExecutorTransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _Executor.Contract.SetMinBlockConfirmations(&_Executor.TransactOpts, minBlockConfirmations)
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

func (_Executor *ExecutorTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "withdrawFeeTokens", feeTokens, feeAggregator)
}

func (_Executor *ExecutorSession) WithdrawFeeTokens(feeTokens []common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _Executor.Contract.WithdrawFeeTokens(&_Executor.TransactOpts, feeTokens, feeAggregator)
}

func (_Executor *ExecutorTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _Executor.Contract.WithdrawFeeTokens(&_Executor.TransactOpts, feeTokens, feeAggregator)
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

type ExecutorMinBlockConfirmationsSetIterator struct {
	Event *ExecutorMinBlockConfirmationsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorMinBlockConfirmationsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorMinBlockConfirmationsSet)
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
		it.Event = new(ExecutorMinBlockConfirmationsSet)
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

func (it *ExecutorMinBlockConfirmationsSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorMinBlockConfirmationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorMinBlockConfirmationsSet struct {
	MinBlockConfirmations uint16
	Raw                   types.Log
}

func (_Executor *ExecutorFilterer) FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*ExecutorMinBlockConfirmationsSetIterator, error) {

	logs, sub, err := _Executor.contract.FilterLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorMinBlockConfirmationsSetIterator{contract: _Executor.contract, event: "MinBlockConfirmationsSet", logs: logs, sub: sub}, nil
}

func (_Executor *ExecutorFilterer) WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *ExecutorMinBlockConfirmationsSet) (event.Subscription, error) {

	logs, sub, err := _Executor.contract.WatchLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorMinBlockConfirmationsSet)
				if err := _Executor.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (_Executor *ExecutorFilterer) ParseMinBlockConfirmationsSet(log types.Log) (*ExecutorMinBlockConfirmationsSet, error) {
	event := new(ExecutorMinBlockConfirmationsSet)
	if err := _Executor.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (ExecutorMinBlockConfirmationsSet) Topic() common.Hash {
	return common.HexToHash("0x46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01")
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, dataLength uint32, numberOfTokens uint8, ccvs []ClientCCV, extraArgs []byte) (GetFee,

		error)

	GetMaxCCVsPerMessage(opts *bind.CallOpts) (uint8, error)

	GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error)

	IsCCVAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error)

	ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []ExecutorRemoteChainConfigArgs) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, feeAggregator common.Address) (*types.Transaction, error)

	FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVAddedIterator, error)

	WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAdded, ccv []common.Address) (event.Subscription, error)

	ParseCCVAdded(log types.Log) (*ExecutorCCVAdded, error)

	FilterCCVAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorCCVAllowlistUpdatedIterator, error)

	WatchCCVAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorCCVAllowlistUpdated) (event.Subscription, error)

	ParseCCVAllowlistUpdated(log types.Log) (*ExecutorCCVAllowlistUpdated, error)

	FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorCCVRemovedIterator, error)

	WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorCCVRemoved, ccv []common.Address) (event.Subscription, error)

	ParseCCVRemoved(log types.Log) (*ExecutorCCVRemoved, error)

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

	FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*ExecutorMinBlockConfirmationsSetIterator, error)

	WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *ExecutorMinBlockConfirmationsSet) (event.Subscription, error)

	ParseMinBlockConfirmationsSet(log types.Log) (*ExecutorMinBlockConfirmationsSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOwnershipTransferred, error)

	Address() common.Address
}
