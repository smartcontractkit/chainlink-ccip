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
	UsdCentsFee uint16
	Enabled     bool
}

type ExecutorRemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            ExecutorRemoteChainConfig
}

var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct Executor.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvs\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxCCVsPerMessage\",\"inputs\":[],\"outputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct Executor.RemoteChainConfig\",\"components\":[{\"name\":\"usdCentsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Executor__RequestedBlockDepthTooLow\",\"inputs\":[{\"name\":\"requestedBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0604052346101c957604051611cdf38819003601f8101601f191683016001600160401b038111848210176101ce5783928291604052833981010390608082126101c95780519060ff8216928383036101c957606090601f1901126101c95760405191606083016001600160401b038111848210176101ce5760405260208201516001600160a01b03811681036101c957835260408201519161ffff831683036101c95760208401928352606001519384151585036101c9576040840194855233156101b857600180546001600160a01b0319163317905580156101a4575060805281516001600160a01b03161561019357905160028054835185516001600160b81b03199092166001600160a01b039490941693841760a09190911b61ffff60a01b161790151560b01b60ff60b01b1617905560408051918252915161ffff16602082015291511515908201527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a1604051611afa90816101e582396080518181816105dd01526106f10152f35b6306b7c75960e31b60005260046000fd5b631f3f959360e01b60005260045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714611359578063240b96e914611186578063336e545a14610f3e5780634c3281c514610daf5780635cb80c5d14610a4e5780637437ff9f1461097957806379ba5097146108905780637cb4e8f21461060157806384502414146105a55780638da5cb5b14610553578063913682e0146102dc578063a68c61a6146101ee578063b8d5005e146101ab5763f2fde38b146100b657600080fd5b346101a65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65773ffffffffffffffffffffffffffffffffffffffff61010261151e565b61010a611618565b1633811461017c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602061ffff60025460a01c16604051908152f35b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576040518060206003549283815201809260036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b8181106102c6575050508161026d910382611435565b6040519182916020830190602084525180915260408301919060005b818110610297575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610289565b8254845260209093019260019283019201610257565b346101a65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a65761032b9036906004016114de565b6024359167ffffffffffffffff83116101a657366023840112156101a65782600401359167ffffffffffffffff83116101a65736602460608502860101116101a657610375611618565b60005b8181106104e357600085855b808310156104e157600092606081028301602481019067ffffffffffffffff6103ac83611603565b16156104a3576103cd67ffffffffffffffff6103c784611603565b166119cb565b5067ffffffffffffffff6103e083611603565b1686526007602052604086209060448101359161ffff83169384840361049f5760648254930135918215159384840361049b579267ffffffffffffffff610483869460019b9c9d947f57ecbe7fefba319b9178ff7edc65aa2cfc028720fa679055210bf987a037eaf699978b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000062ff000060409c60101b16921617179055611603565b169685519450845250506020820152a2019190610384565b8a80fd5b8880fd5b60248667ffffffffffffffff6104b885611603565b7f020a07e500000000000000000000000000000000000000000000000000000000835216600452fd5b005b8067ffffffffffffffff6105026104fd60019486886115d2565b611603565b1661050c81611840565b610518575b5001610378565b806000526007602052600060408120557ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a286610511565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101a65760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81168091036101a65761064d611541565b60443567ffffffffffffffff81116101a65761066d9036906004016114de565b9160643567ffffffffffffffff81116101a657366023820112156101a657806004013567ffffffffffffffff81116101a657369101602401116101a6578360005260076020526106c060406000206115ad565b93602085015115610863575061ffff1680151580610851575b610816575060ff60025460b01c16610759575b5060ff7f0000000000000000000000000000000000000000000000000000000000000000169081811161072957602061ffff845116604051908152f35b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b90916000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc183360301915b8481101561080b5760008160051b85013584811215610807576107be73ffffffffffffffffffffffffffffffffffffffff9187016115e2565b1680825260046020526040822054156107db575050600101610785565b602492507fa409d83e000000000000000000000000000000000000000000000000000000008252600452fd5b5080fd5b5092915050826106ec565b61ffff60025460a01c16907f2dba20cf0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5061ffff60025460a01c1681106106d9565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760005473ffffffffffffffffffffffffffffffffffffffff8116330361094f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576000604080516109b781611419565b8281528260208201520152610a4a6040516109d181611419565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835261ffff8160a01c16602084015260b01c161515604082015260405191829182919091604080606083019473ffffffffffffffffffffffffffffffffffffffff815116845261ffff602082015116602085015201511515910152565b0390f35b346101a65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a657610a9d9036906004016114de565b90610aa6611618565b73ffffffffffffffffffffffffffffffffffffffff600254169160005b818110610acc57005b73ffffffffffffffffffffffffffffffffffffffff610af4610aef8385876115d2565b6115e2565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa928315610da357600093610d6d575b5082610b4d575b506001915001610ac3565b6040519260208401937fa9059cbb00000000000000000000000000000000000000000000000000000000855287602482015281604482015260448152610b94606482611435565b600080604096875193610ba78986611435565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af1883d15610d5f57503d90600067ffffffffffffffff8311610d32575091818a949360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f610c429601160191610c318a519384611435565b82523d6000602084013e5b86611a25565b805180610c7e575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385610b42565b8192949596935090602091810103126101a657602001518015908115036101a657610caf5792919086908880610c4a565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b9291610c4291606090610c3c565b90926020823d8211610d9b575b81610d8760209383611435565b81010312610d985750519186610b3b565b80fd5b3d9150610d7a565b6040513d6000823e3d90fd5b346101a65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576000604051610dec81611419565b610df461151e565b8152610dfe611541565b60208201908152610e0d61150f565b9060408301918252610e1d611618565b73ffffffffffffffffffffffffffffffffffffffff83511615610f1657825160028054925193517fffffffffffffffffff000000000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff90921691821760a09490941b75ffff0000000000000000000000000000000000000000169390931791151560b01b76ff0000000000000000000000000000000000000000000016919091179091556040805191825260208084015161ffff1690830152918201511515918101919091527f4c475597c445491197895d935b9c8eaf2c59a575d8a4577ed31a8bbb48b6589290606090a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b346101a65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a65760043567ffffffffffffffff81116101a657610f8d9036906004016114de565b9060243567ffffffffffffffff81116101a657610fae9036906004016114de565b9091610fb861150f565b93610fc1611618565b60005b8181106110fe5750505060005b8181106110635783600254901515908160ff8260b01c16151503610ff157005b816020917fffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffff76ff000000000000000000000000000000000000000000007fd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed89560b01b16911617600255604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff611086610aef8385876115d2565b1680156110d157908161109a60019361196b565b6110a6575b5001610fd1565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a28561109f565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8061112d73ffffffffffffffffffffffffffffffffffffffff611127610aef60019587896115d2565b1661167b565b611138575b01610fc4565b73ffffffffffffffffffffffffffffffffffffffff61115b610aef8386886115d2565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a2611132565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a6576005546111c181611552565b906111cf6040519283611435565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06111fc82611552565b0160005b8181106113205750506005549060005b818110611283578360405180916020820160208352815180915260206040840192019060005b818110611244575050500390f35b8251805167ffffffffffffffff168552602090810151805161ffff16828701528101511515604086015286955060609094019390920191600101611236565b6000838210156112f35790604081602084600560019652200167ffffffffffffffff60009154166112b4848961156a565b515267ffffffffffffffff6112c9848961156a565b51511681526007602052206112eb60206112e3848961156a565b5101916115ad565b905201611210565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60209060405161132f816113ce565b6000815260405161133f816113ce565b600081526000848201528382015282828701015201611200565b346101a65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a657610a4a604080519061139a8183611435565b601282527f4578656375746f7220312e372e302d646576000000000000000000000000000060208301525191829182611476565b6040810190811067ffffffffffffffff8211176113ea57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176113ea57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176113ea57604052565b9190916020815282519283602083015260005b8481106114c85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201611489565b9181601f840112156101a65782359167ffffffffffffffff83116101a6576020808501948460051b0101116101a657565b6044359081151582036101a657565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101a657565b6024359061ffff821682036101a657565b67ffffffffffffffff81116113ea5760051b60200190565b805182101561157e5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906040516115ba816113ce565b915461ffff8116835260101c60ff1615156020830152565b919081101561157e5760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101a65790565b3567ffffffffffffffff811681036101a65790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361163957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805482101561157e5760005260206000200190600090565b6000818152600460205260409020548015611839577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161180a57600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161180a5781810361179b575b505050600354801561176c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611729816003611663565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6117f26117ac6117bd936003611663565b90549060031b1c9283926003611663565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260046020526040600020553880806116f0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b6000818152600660205260409020548015611839577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161180a57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161180a57818103611931575b505050600554801561176c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016118ee816005611663565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6119536119426117bd936005611663565b90549060031b1c9283926005611663565b905560005260066020526040600020553880806118b5565b806000526004602052604060002054156000146119c557600354680100000000000000008110156113ea576119ac6117bd8260018594016003556003611663565b9055600354906000526004602052604060002055600190565b50600090565b806000526006602052604060002054156000146119c557600554680100000000000000008110156113ea57611a0c6117bd8260018594016005556005611663565b9055600554906000526006602052604060002055600190565b91929015611aa05750815115611a39575090565b3b15611a425790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015611ab35750805190602001fd5b611ae9906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301611476565b0390fdfea164736f6c634300081a000a",
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

func (_Executor *ExecutorCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, ccvs []ClientCCV, arg3 []byte) (uint16, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getFee", destChainSelector, requestedBlockDepth, ccvs, arg3)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Executor *ExecutorSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, ccvs []ClientCCV, arg3 []byte) (uint16, error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, ccvs, arg3)
}

func (_Executor *ExecutorCallerSession) GetFee(destChainSelector uint64, requestedBlockDepth uint16, ccvs []ClientCCV, arg3 []byte) (uint16, error) {
	return _Executor.Contract.GetFee(&_Executor.CallOpts, destChainSelector, requestedBlockDepth, ccvs, arg3)
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, requestedBlockDepth uint16, ccvs []ClientCCV, arg3 []byte) (uint16, error)

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
