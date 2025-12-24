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

type ExecutorDynamicConfig struct {
	FeeAggregator         common.Address
	MinBlockConfirmations uint16
	CcvAllowlistEnabled   bool
}

type ExecutorRemoteChainConfig struct {
	UsdCentsFee uint16
	Enabled     bool
}

type ExecutorRemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            ExecutorRemoteChainConfig
}

var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxCCVsPerMessage\",\"inputs\":[],\"outputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Executor__RequestedBlockDepthTooLow\",\"inputs\":[{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a0604052346101c957604051611ac138819003601f8101601f191683016001600160401b038111848210176101ce5783928291604052833981010390608082126101c95780519060ff8216928383036101c957606090601f1901126101c95760405191606083016001600160401b038111848210176101ce5760405260208201516001600160a01b03811681036101c957835260408201519161ffff831683036101c95760208401928352606001519384151585036101c9576040840194855233156101b857600180546001600160a01b0319163317905580156101a4575060805281516001600160a01b03161561019357905160028054835185516001600160b81b03199092166001600160a01b039490941693841760a09190911b61ffff60a01b161790151560b01b60ff60b01b1617905560408051918252915161ffff16602082015291511515908201527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a16040516118dc90816101e582396080518181816102a701526108600152f35b6306b7c75960e31b60005260046000fd5b631f3f959360e01b60005260045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a77146111ee57508063240b96e91461101b578063336e545a14610dd35780634c3281c514610c445780635cb80c5d14610a425780637437ff9f1461096d57806379ba50971461088457806384502414146108285780638da5cb5b146107d6578063913682e01461055f578063a68c61a614610471578063b8d5005e1461042e578063f2388958146101ae5763f2fde38b146100b957600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95773ffffffffffffffffffffffffffffffffffffffff6101056113a5565b61010d6114c2565b1633811461017f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81168091036101a9576101fa6113eb565b60443567ffffffffffffffff81116101a95761021a903690600401611365565b9160643567ffffffffffffffff81116101a957366023820112156101a957806004013567ffffffffffffffff81116101a957369101602401116101a95761025f6113c8565b508360005260076020526102766040600020611457565b93602085015115610401575061ffff16801515806103ef575b6103b4575060ff60025460b01c1661030f575b5060ff7f000000000000000000000000000000000000000000000000000000000000000016908181116102df57602061ffff845116604051908152f35b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b60005b82811061031f57506102a2565b604073ffffffffffffffffffffffffffffffffffffffff61034961034484878761147c565b61148c565b166000908152600460205220541561036357600101610312565b6103449073ffffffffffffffffffffffffffffffffffffffff936103869361147c565b7fa409d83e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b61ffff60025460a01c16907f2dba20cf0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5061ffff60025460a01c16811061028f565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602061ffff60025460a01c16604051908152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576040518060206003549283815201809260036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b81811061054957505050816104f0910382611324565b6040519182916020830190602084525180915260408301919060005b81811061051a575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161050c565b82548452602090930192600192830192016104da565b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a9576105ae903690600401611365565b6024359167ffffffffffffffff83116101a957366023840112156101a95782600401359167ffffffffffffffff83116101a95736602460608502860101116101a9576105f86114c2565b60005b81811061076657600085855b8083101561076457600092606081028301602481019067ffffffffffffffff61062f836114ad565b16156107265761065067ffffffffffffffff61064a846114ad565b16611875565b5067ffffffffffffffff610663836114ad565b1686526007602052604086209060448101359161ffff8316938484036107225760648254930135918215159384840361071e579267ffffffffffffffff610706869460019b9c9d947f57ecbe7fefba319b9178ff7edc65aa2cfc028720fa679055210bf987a037eaf699978b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000062ff000060409c60101b169216171790556114ad565b169685519450845250506020820152a2019190610607565b8a80fd5b8880fd5b60248667ffffffffffffffff61073b856114ad565b7f020a07e500000000000000000000000000000000000000000000000000000000835216600452fd5b005b8067ffffffffffffffff610785610780600194868861147c565b6114ad565b1661078f816116ea565b61079b575b50016105fb565b806000526007602052600060408120557ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a286610794565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760005473ffffffffffffffffffffffffffffffffffffffff81163303610943577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576000604080516109ab81611308565b8281528260208201520152610a3e6040516109c581611308565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835261ffff8160a01c16602084015260b01c161515604082015260405191829182919091604080606083019473ffffffffffffffffffffffffffffffffffffffff815116845261ffff602082015116602085015201511515910152565b0390f35b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957610a91903690600401611365565b9073ffffffffffffffffffffffffffffffffffffffff600254169160005b818110610ab857005b73ffffffffffffffffffffffffffffffffffffffff610adb61034483858761147c565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115610c0457600091610c10575b5080610b31575b5050600101610aaf565b60206000604051828101907fa9059cbb00000000000000000000000000000000000000000000000000000000825289602482015284604482015260448152610b7a606482611324565b519082865af115610c04576000513d610bfb5750813b155b610bcd5790857f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39085610b27565b507f5274afe70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60011415610b92565b6040513d6000823e3d90fd5b906020823d8211610c3c575b81610c2960209383611324565b81010312610c3957505186610b20565b80fd5b3d9150610c1c565b346101a95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576000604051610c8181611308565b610c896113a5565b8152610c936113eb565b60208201908152610ca2611396565b9060408301918252610cb26114c2565b73ffffffffffffffffffffffffffffffffffffffff83511615610dab57825160028054925193517fffffffffffffffffff000000000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff90921691821760a09490941b75ffff0000000000000000000000000000000000000000169390931791151560b01b76ff0000000000000000000000000000000000000000000016919091179091556040805191825260208084015161ffff1690830152918201511515918101919091527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b346101a95760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957610e22903690600401611365565b9060243567ffffffffffffffff81116101a957610e43903690600401611365565b9091610e4d611396565b93610e566114c2565b60005b818110610f935750505060005b818110610ef85783600254901515908160ff8260b01c16151503610e8657005b816020917fffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffff76ff000000000000000000000000000000000000000000007fd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed89560b01b16911617600255604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff610f1b61034483858761147c565b168015610f66579081610f2f600193611815565b610f3b575b5001610e66565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a285610f34565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80610fc273ffffffffffffffffffffffffffffffffffffffff610fbc610344600195878961147c565b16611525565b610fcd575b01610e59565b73ffffffffffffffffffffffffffffffffffffffff610ff061034483868861147c565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a2610fc7565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957600554611056816113fc565b906110646040519283611324565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611091826113fc565b0160005b8181106111b55750506005549060005b818110611118578360405180916020820160208352815180915260206040840192019060005b8181106110d9575050500390f35b8251805167ffffffffffffffff168552602090810151805161ffff168287015281015115156040860152869550606090940193909201916001016110cb565b6000838210156111885790604081602084600560019652200167ffffffffffffffff60009154166111498489611414565b515267ffffffffffffffff61115e8489611414565b515116815260076020522061118060206111788489611414565b510191611457565b9052016110a5565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b6020906040516111c4816112bd565b600081526040516111d4816112bd565b600081526000848201528382015282828701015201611095565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957611226816112bd565b601281527f4578656375746f7220312e372e302d6465760000000000000000000000000000602082015260405190602082528181519182602083015260005b8381106112a55750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001611265565b6040810190811067ffffffffffffffff8211176112d957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176112d957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176112d957604052565b9181601f840112156101a95782359167ffffffffffffffff83116101a9576020808501948460051b0101116101a957565b6044359081151582036101a957565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b6084359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b6024359061ffff821682036101a957565b67ffffffffffffffff81116112d95760051b60200190565b80518210156114285760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90604051611464816112bd565b915461ffff8116835260101c60ff1615156020830152565b91908110156114285760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101a95790565b3567ffffffffffffffff811681036101a95790565b73ffffffffffffffffffffffffffffffffffffffff6001541633036114e357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156114285760005260206000200190600090565b60008181526004602052604090205480156116e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116116b457600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116116b457818103611645575b5050506003548015611616577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016115d381600361150d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61169c61165661166793600361150d565b90549060031b1c928392600361150d565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600460205260406000205538808061159a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b60008181526006602052604090205480156116e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116116b457600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116116b4578181036117db575b5050506005548015611616577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161179881600561150d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6117fd6117ec61166793600561150d565b90549060031b1c928392600561150d565b9055600052600660205260406000205538808061175f565b8060005260046020526040600020541560001461186f57600354680100000000000000008110156112d957611856611667826001859401600355600361150d565b9055600354906000526004602052604060002055600190565b50600090565b8060005260066020526040600020541560001461186f57600554680100000000000000008110156112d9576118b6611667826001859401600555600561150d565b905560055490600052600660205260406000205560019056fea164736f6c634300081a000a",
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

func (_Executor *ExecutorCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, ccvs []common.Address, arg3 []byte, arg4 common.Address) (uint16, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getFee", destChainSelector, requestedBlockDepth, ccvs, arg3, arg4)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Executor *ExecutorSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, ccvs []common.Address, arg3 []byte, arg4 common.Address) (uint16, error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, ccvs, arg3, arg4)
}

func (_Executor *ExecutorCallerSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, ccvs []common.Address, arg3 []byte, arg4 common.Address) (uint16, error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, ccvs, arg3, arg4)
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
	return common.HexToHash("0x57ecbe7fefba319b9178ff7edc65aa2cfc028720fa679055210bf987a037eaf6")
}

func (ExecutorDestChainRemoved) Topic() common.Hash {
	return common.HexToHash("0xf74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1")
}

func (ExecutorFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, ccvs []common.Address, arg3 []byte, arg4 common.Address) (uint16, error)

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

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOwnershipTransferred, error)

	Address() common.Address
}
