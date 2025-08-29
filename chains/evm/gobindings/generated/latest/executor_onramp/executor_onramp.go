// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor_onramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type ExecutorOnRampDynamicConfig struct {
	FeeQuoter             common.Address
	FeeAggregator         common.Address
	MaxPossibleCCVsPerMsg uint8
	MaxRequiredCCVsPerMsg uint8
}

var ExecutorOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxRequiredCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxRequiredCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_allowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxRequiredCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxRequiredCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxPossibleCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExceedsMaxRequiredCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsVersion\",\"inputs\":[{\"name\":\"provided\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x6080604052346101f357604051601f611c8738819003918201601f19168301916001600160401b038311848410176101f8578084926080946040528339810103126101f35760405190600090608083016001600160401b038111848210176101df5760405261006d8161020e565b835261007b6020820161020e565b92602081019384526100a2606061009460408501610222565b936040840194855201610222565b926060820193845233156101d057600180546001600160a01b0319163317905581516001600160a01b03161580156101be575b80156101ad575b61019e578151600680546001600160a01b0319166001600160a01b03928316908117909155865160078054875189516001600160b01b03199092169386169390931760a09390931b60ff60a01b169290921760a89290921b60ff60a81b169190911790556040805191825287519092166020820152845160ff90811692820192909252855190911660608201527f30c2daad3a22daff505703d0b198d15e894cb8a8db323db4f002e2123b238c5e90608090a1604051611a5690816102318239f35b6306b7c75960e31b8152600490fd5b5060ff84511660ff845116106100dc565b5084516001600160a01b0316156100d5565b639b15e16f60e01b8152600490fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101f357565b519060ff821682036101f35756fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714611282578063240b96e91461115e578063336e545a14610f7d5780634a7597b51461099e57806363115c3a146107125780637437ff9f1461060557806379ba50971461051c5780638da5cb5b146104ca578063a422fdb514610347578063a68c61a6146101d9578063d5d91083146101955763f2fde38b146100a057600080fd5b346101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905773ffffffffffffffffffffffffffffffffffffffff6100ec61148d565b6100f461166b565b1633811461016657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057602060ff60015460a01c166040519015158152f35b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760025461021c61021782611475565b61138f565b9080825261022981611475565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376002549060005b8181106102bd5783856040519182916020830190602084525180915260408301919060005b81811061028e575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610280565b60008382101561031a57600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101546001919073ffffffffffffffffffffffffffffffffffffffff1661031382886114d1565b520161025b565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b346101905760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff8111610190576103969036906004016113d3565b9060243567ffffffffffffffff8111610190576103b79036906004016113d3565b9190926103c261166b565b60005b81811061043b5750505060005b8181106103db57005b8067ffffffffffffffff6103fa6103f56001948688611514565b611656565b166104048161191e565b610410575b50016103d2565b7ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a284610409565b67ffffffffffffffff6104526103f5838587611514565b16801561049d579081610466600193611763565b610472575b50016103c5565b7f6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f600080a28661046b565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760005473ffffffffffffffffffffffffffffffffffffffff811633036105db577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610190576000606061064061136f565b828152826020820152826040820152015261070e61065c61136f565b73ffffffffffffffffffffffffffffffffffffffff60065416815260ff60075473ffffffffffffffffffffffffffffffffffffffff81166020840152818160a01c16604084015260a81c16606082015260405191829182919091606060ff81608084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152826040820151166040860152015116910152565b0390f35b346101905760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019057600061074b61136f565b61075361148d565b815260243573ffffffffffffffffffffffffffffffffffffffff8116810361099a576020820190815260443560ff8116810361099657604083019081526064359160ff8316830361099257606084019283526107ad61166b565b73ffffffffffffffffffffffffffffffffffffffff845116158015610973575b8015610962575b61093a5791839173ffffffffffffffffffffffffffffffffffffffff61093494817f30c2daad3a22daff505703d0b198d15e894cb8a8db323db4f002e2123b238c5e9751167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075551907fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000075ff000000000000000000000000000000000000000000600754935160a81b169360a01b169116171760075560405191829182919091606060ff81608084019573ffffffffffffffffffffffffffffffffffffffff815116855273ffffffffffffffffffffffffffffffffffffffff6020820151166020860152826040820151166040860152015116910152565b0390a180f35b6004857f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5060ff83511660ff835116106107d4565b5073ffffffffffffffffffffffffffffffffffffffff815116156107cd565b8480fd5b8380fd5b8280fd5b346101905760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff81168091036101905760243567ffffffffffffffff81116101905760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610190576040519060a0820182811067ffffffffffffffff821117610eec57604052806004013567ffffffffffffffff811161019057610a629060043691840101611404565b8252602481013567ffffffffffffffff811161019057610a889060043691840101611404565b6020830152604481013567ffffffffffffffff811161019057810136602382011215610190576004810135610abf61021782611475565b91602060048185858152019360061b830101019036821161019057602401915b818310610f48575050506040830152610afa606482016114b0565b6060830152608481013567ffffffffffffffff8111610190576080916004610b259236920101611404565b9101526044359067ffffffffffffffff821161019057366023830112156101905781600401359167ffffffffffffffff831161019057828101926024840192368411610190578060005260056020526040600020541561049d575060041180610190577fffffffff000000000000000000000000000000000000000000000000000000006024830135167f302326cb000000000000000000000000000000000000000000000000000000008103610f1b57506101905760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828503011261019057602881013567ffffffffffffffff8111610190577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd860249160e09301809503010112610190576040519160e0830183811067ffffffffffffffff821117610eec57604052602881013567ffffffffffffffff811161019057826028610c8e92840101611545565b8352604881013567ffffffffffffffff811161019057826028610cb392840101611545565b906020840191825260688101359260ff841684036101905760408501938452608882013563ffffffff81168103610190576060860152610cf560a883016114b0565b608086015260c882013567ffffffffffffffff811161019057816028610d1d92850101611404565b60a086015260e88201359167ffffffffffffffff831161019057610d449201602801611404565b60c084015260ff60015460a01c16610e06575b825151905151610d669161161a565b916007549260ff8460a01c1690818111610dd6578460ff610d8e86828751519151169061161a565b9160a81c1690818111610da657602060405160008152f35b7f6e39d84c0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fab57cee00000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b919060005b82518051821015610e8a57610e358273ffffffffffffffffffffffffffffffffffffffff926114d1565b515116610e4f816000526003602052604060002054151590565b15610e5d5750600101610e0b565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5050919060005b82518051821015610ee357610ebb8273ffffffffffffffffffffffffffffffffffffffff926114d1565b515116610ed5816000526003602052604060002054151590565b15610e5d5750600101610e91565b50509190610d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7ff08625a80000000000000000000000000000000000000000000000000000000060005260045260246000fd5b604083360312610190576020604091610f5f61134f565b610f68866114b0565b81528286013583820152815201920191610adf565b346101905760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760043567ffffffffffffffff811161019057610fcc9036906004016113d3565b9060243567ffffffffffffffff811161019057610fed9036906004016113d3565b9091604435938415158095036101905761100561166b565b60005b8181106110f05750505060005b81811061106357600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1660a086901b74ff000000000000000000000000000000000000000016179055005b8061109773ffffffffffffffffffffffffffffffffffffffff61109161108c6001958789611514565b611524565b166117bd565b6110a2575b01611015565b73ffffffffffffffffffffffffffffffffffffffff6110c561108c838688611514565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a261109c565b73ffffffffffffffffffffffffffffffffffffffff61111361108c838587611514565b168015610e5d5790816111276001936116ce565b611133575b5001611008565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a28761112c565b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760045461119c61021782611475565b908082526111a981611475565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376004549060005b8181106112315783856040519182916020830190602084525180915260408301919060005b81811061120e575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611200565b60008382101561031a57600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b8101546001919067ffffffffffffffff1661127b82886114d1565b52016111db565b346101905760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101905760406112bc8161138f565b90601882527f4578656375746f724f6e52616d7020312e372e302d64657600000000000000006020830152805180926020825280519081602084015260005b8281106113385750506000828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b6020828201810151878301870152869450016112fb565b604051906040820182811067ffffffffffffffff821117610eec57604052565b604051906080820182811067ffffffffffffffff821117610eec57604052565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff821117610eec57604052565b9181601f840112156101905782359167ffffffffffffffff8311610190576020808501948460051b01011161019057565b81601f820112156101905780359067ffffffffffffffff8211610eec5761145260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160161138f565b928284526020838301011161019057816000926020809301838601378301015290565b67ffffffffffffffff8111610eec5760051b60200190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361019057565b359073ffffffffffffffffffffffffffffffffffffffff8216820361019057565b80518210156114e55760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156114e55760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036101905790565b9080601f830112156101905781359161156061021784611475565b9260208085838152019160051b830101918383116101905760208101915b83831061158d57505050505090565b823567ffffffffffffffff81116101905782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610190576115d561134f565b906115e2602084016114b0565b825260408301359167ffffffffffffffff83116101905761160b88602080969581960101611404565b8382015281520192019161157e565b9190820180921161162757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b3567ffffffffffffffff811681036101905790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361168c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156114e55760005260206000200190600090565b8060005260036020526040600020541560001461175d5760025468010000000000000000811015610eec5761174461170f82600185940160025560026116b6565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b8060005260056020526040600020541560001461175d5760045468010000000000000000811015610eec576117a461170f82600185940160045560046116b6565b9055600454906000526005602052604060002055600190565b6000818152600360205260409020548015611917577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161162757600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611627578181036118dd575b50505060025480156118ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161186b8160026116b6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6118ff6118ee61170f9360026116b6565b90549060031b1c92839260026116b6565b90556000526003602052604060002055388080611832565b5050600090565b6000818152600560205260409020548015611917577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161162757600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161162757818103611a0f575b50505060045480156118ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016119cc8160046116b6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b611a31611a2061170f9360046116b6565b90549060031b1c92839260046116b6565b9055600052600560205260406000205538808061199356fea164736f6c634300081a000a",
}

var ExecutorOnRampABI = ExecutorOnRampMetaData.ABI

var ExecutorOnRampBin = ExecutorOnRampMetaData.Bin

func DeployExecutorOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig ExecutorOnRampDynamicConfig) (common.Address, *types.Transaction, *ExecutorOnRamp, error) {
	parsed, err := ExecutorOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorOnRampBin), backend, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutorOnRamp{address: address, abi: *parsed, ExecutorOnRampCaller: ExecutorOnRampCaller{contract: contract}, ExecutorOnRampTransactor: ExecutorOnRampTransactor{contract: contract}, ExecutorOnRampFilterer: ExecutorOnRampFilterer{contract: contract}}, nil
}

type ExecutorOnRamp struct {
	address common.Address
	abi     abi.ABI
	ExecutorOnRampCaller
	ExecutorOnRampTransactor
	ExecutorOnRampFilterer
}

type ExecutorOnRampCaller struct {
	contract *bind.BoundContract
}

type ExecutorOnRampTransactor struct {
	contract *bind.BoundContract
}

type ExecutorOnRampFilterer struct {
	contract *bind.BoundContract
}

type ExecutorOnRampSession struct {
	Contract     *ExecutorOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ExecutorOnRampCallerSession struct {
	Contract *ExecutorOnRampCaller
	CallOpts bind.CallOpts
}

type ExecutorOnRampTransactorSession struct {
	Contract     *ExecutorOnRampTransactor
	TransactOpts bind.TransactOpts
}

type ExecutorOnRampRaw struct {
	Contract *ExecutorOnRamp
}

type ExecutorOnRampCallerRaw struct {
	Contract *ExecutorOnRampCaller
}

type ExecutorOnRampTransactorRaw struct {
	Contract *ExecutorOnRampTransactor
}

func NewExecutorOnRamp(address common.Address, backend bind.ContractBackend) (*ExecutorOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(ExecutorOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindExecutorOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRamp{address: address, abi: abi, ExecutorOnRampCaller: ExecutorOnRampCaller{contract: contract}, ExecutorOnRampTransactor: ExecutorOnRampTransactor{contract: contract}, ExecutorOnRampFilterer: ExecutorOnRampFilterer{contract: contract}}, nil
}

func NewExecutorOnRampCaller(address common.Address, caller bind.ContractCaller) (*ExecutorOnRampCaller, error) {
	contract, err := bindExecutorOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCaller{contract: contract}, nil
}

func NewExecutorOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorOnRampTransactor, error) {
	contract, err := bindExecutorOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampTransactor{contract: contract}, nil
}

func NewExecutorOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorOnRampFilterer, error) {
	contract, err := bindExecutorOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampFilterer{contract: contract}, nil
}

func bindExecutorOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutorOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorOnRamp.Contract.ExecutorOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ExecutorOnRampTransactor.contract.Transfer(opts)
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ExecutorOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.contract.Transfer(opts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getAllowedCCVs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetAllowedCCVs() ([]common.Address, error) {
	return _ExecutorOnRamp.Contract.GetAllowedCCVs(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetAllowedCCVs() ([]common.Address, error) {
	return _ExecutorOnRamp.Contract.GetAllowedCCVs(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetDestChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getDestChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetDestChains() ([]uint64, error) {
	return _ExecutorOnRamp.Contract.GetDestChains(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetDestChains() ([]uint64, error) {
	return _ExecutorOnRamp.Contract.GetDestChains(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (ExecutorOnRampDynamicConfig, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(ExecutorOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ExecutorOnRampDynamicConfig)).(*ExecutorOnRampDynamicConfig)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetDynamicConfig() (ExecutorOnRampDynamicConfig, error) {
	return _ExecutorOnRamp.Contract.GetDynamicConfig(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetDynamicConfig() (ExecutorOnRampDynamicConfig, error) {
	return _ExecutorOnRamp.Contract.GetDynamicConfig(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, extraArgs []byte) (*big.Int, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, arg1, extraArgs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, extraArgs []byte) (*big.Int, error) {
	return _ExecutorOnRamp.Contract.GetFee(&_ExecutorOnRamp.CallOpts, destChainSelector, arg1, extraArgs)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, extraArgs []byte) (*big.Int, error) {
	return _ExecutorOnRamp.Contract.GetFee(&_ExecutorOnRamp.CallOpts, destChainSelector, arg1, extraArgs)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) Owner() (common.Address, error) {
	return _ExecutorOnRamp.Contract.Owner(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) Owner() (common.Address, error) {
	return _ExecutorOnRamp.Contract.Owner(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) SAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "s_allowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) SAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.SAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) SAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.SAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) TypeAndVersion() (string, error) {
	return _ExecutorOnRamp.Contract.TypeAndVersion(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) TypeAndVersion() (string, error) {
	return _ExecutorOnRamp.Contract.TypeAndVersion(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_ExecutorOnRamp *ExecutorOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.AcceptOwnership(&_ExecutorOnRamp.TransactOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.AcceptOwnership(&_ExecutorOnRamp.TransactOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyAllowedCCVUpdates", ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyAllowedCCVUpdates(ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyAllowedCCVUpdates(ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyDestChainUpdates", destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyDestChainUpdates(destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyDestChainUpdates(destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) SetDynamicConfig(dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetDynamicConfig(&_ExecutorOnRamp.TransactOpts, dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) SetDynamicConfig(dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetDynamicConfig(&_ExecutorOnRamp.TransactOpts, dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.TransferOwnership(&_ExecutorOnRamp.TransactOpts, to)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.TransferOwnership(&_ExecutorOnRamp.TransactOpts, to)
}

type ExecutorOnRampCCVAddedIterator struct {
	Event *ExecutorOnRampCCVAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampCCVAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampCCVAdded)
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
		it.Event = new(ExecutorOnRampCCVAdded)
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

func (it *ExecutorOnRampCCVAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampCCVAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampCCVAdded struct {
	Ccv common.Address
	Raw types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVAddedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCCVAddedIterator{contract: _ExecutorOnRamp.contract, event: "CCVAdded", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAdded, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampCCVAdded)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAdded", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseCCVAdded(log types.Log) (*ExecutorOnRampCCVAdded, error) {
	event := new(ExecutorOnRampCCVAdded)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampCCVRemovedIterator struct {
	Event *ExecutorOnRampCCVRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampCCVRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampCCVRemoved)
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
		it.Event = new(ExecutorOnRampCCVRemoved)
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

func (it *ExecutorOnRampCCVRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampCCVRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampCCVRemoved struct {
	Ccv common.Address
	Raw types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVRemovedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCCVRemovedIterator{contract: _ExecutorOnRamp.contract, event: "CCVRemoved", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVRemoved, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampCCVRemoved)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseCCVRemoved(log types.Log) (*ExecutorOnRampCCVRemoved, error) {
	event := new(ExecutorOnRampCCVRemoved)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampConfigSetIterator struct {
	Event *ExecutorOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampConfigSet)
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
		it.Event = new(ExecutorOnRampConfigSet)
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

func (it *ExecutorOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampConfigSet struct {
	DynamicConfig ExecutorOnRampDynamicConfig
	Raw           types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ExecutorOnRampConfigSetIterator, error) {

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampConfigSetIterator{contract: _ExecutorOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampConfigSet)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseConfigSet(log types.Log) (*ExecutorOnRampConfigSet, error) {
	event := new(ExecutorOnRampConfigSet)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampDestChainAddedIterator struct {
	Event *ExecutorOnRampDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampDestChainAdded)
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
		it.Event = new(ExecutorOnRampDestChainAdded)
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

func (it *ExecutorOnRampDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampDestChainAdded struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampDestChainAddedIterator{contract: _ExecutorOnRamp.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampDestChainAdded)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseDestChainAdded(log types.Log) (*ExecutorOnRampDestChainAdded, error) {
	event := new(ExecutorOnRampDestChainAdded)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampDestChainRemovedIterator struct {
	Event *ExecutorOnRampDestChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampDestChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampDestChainRemoved)
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
		it.Event = new(ExecutorOnRampDestChainRemoved)
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

func (it *ExecutorOnRampDestChainRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampDestChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampDestChainRemoved struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampDestChainRemovedIterator{contract: _ExecutorOnRamp.contract, event: "DestChainRemoved", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampDestChainRemoved)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseDestChainRemoved(log types.Log) (*ExecutorOnRampDestChainRemoved, error) {
	event := new(ExecutorOnRampDestChainRemoved)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampOwnershipTransferRequestedIterator struct {
	Event *ExecutorOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampOwnershipTransferRequested)
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
		it.Event = new(ExecutorOnRampOwnershipTransferRequested)
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

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampOwnershipTransferRequestedIterator{contract: _ExecutorOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampOwnershipTransferRequested)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*ExecutorOnRampOwnershipTransferRequested, error) {
	event := new(ExecutorOnRampOwnershipTransferRequested)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampOwnershipTransferredIterator struct {
	Event *ExecutorOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampOwnershipTransferred)
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
		it.Event = new(ExecutorOnRampOwnershipTransferred)
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

func (it *ExecutorOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampOwnershipTransferredIterator{contract: _ExecutorOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampOwnershipTransferred)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutorOnRampOwnershipTransferred, error) {
	event := new(ExecutorOnRampOwnershipTransferred)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ExecutorOnRamp *ExecutorOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ExecutorOnRamp.abi.Events["CCVAdded"].ID:
		return _ExecutorOnRamp.ParseCCVAdded(log)
	case _ExecutorOnRamp.abi.Events["CCVRemoved"].ID:
		return _ExecutorOnRamp.ParseCCVRemoved(log)
	case _ExecutorOnRamp.abi.Events["ConfigSet"].ID:
		return _ExecutorOnRamp.ParseConfigSet(log)
	case _ExecutorOnRamp.abi.Events["DestChainAdded"].ID:
		return _ExecutorOnRamp.ParseDestChainAdded(log)
	case _ExecutorOnRamp.abi.Events["DestChainRemoved"].ID:
		return _ExecutorOnRamp.ParseDestChainRemoved(log)
	case _ExecutorOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _ExecutorOnRamp.ParseOwnershipTransferRequested(log)
	case _ExecutorOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _ExecutorOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ExecutorOnRampCCVAdded) Topic() common.Hash {
	return common.HexToHash("0xba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e")
}

func (ExecutorOnRampCCVRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e")
}

func (ExecutorOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x30c2daad3a22daff505703d0b198d15e894cb8a8db323db4f002e2123b238c5e")
}

func (ExecutorOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f")
}

func (ExecutorOnRampDestChainRemoved) Topic() common.Hash {
	return common.HexToHash("0xf74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1")
}

func (ExecutorOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ExecutorOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ExecutorOnRamp *ExecutorOnRamp) Address() common.Address {
	return _ExecutorOnRamp.address
}

type ExecutorOnRampInterface interface {
	GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChains(opts *bind.CallOpts) ([]uint64, error)

	GetDynamicConfig(opts *bind.CallOpts) (ExecutorOnRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, extraArgs []byte) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error)

	ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVAddedIterator, error)

	WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAdded, ccv []common.Address) (event.Subscription, error)

	ParseCCVAdded(log types.Log) (*ExecutorOnRampCCVAdded, error)

	FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVRemovedIterator, error)

	WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVRemoved, ccv []common.Address) (event.Subscription, error)

	ParseCCVRemoved(log types.Log) (*ExecutorOnRampCCVRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*ExecutorOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*ExecutorOnRampConfigSet, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*ExecutorOnRampDestChainAdded, error)

	FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainRemovedIterator, error)

	WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainRemoved(log types.Log) (*ExecutorOnRampDestChainRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
