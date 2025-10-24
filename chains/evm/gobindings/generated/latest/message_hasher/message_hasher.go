// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_hasher

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

type ClientEVMExtraArgsV1 struct {
	GasLimit *big.Int
}

type ClientGenericExtraArgsV2 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
}

type ClientSVMExtraArgsV1 struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}

type ClientSuiExtraArgsV1 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	ReceiverObjectIds        [][32]byte
}

type InternalAny2EVMRampMessage struct {
	Header       InternalRampMessageHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     *big.Int
	TokenAmounts []InternalAny2EVMTokenTransfer
}

type InternalAny2EVMTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	DestGasAmount     uint32
	ExtraData         []byte
	Amount            *big.Int
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

var MessageHasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSVMExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSuiExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeAny2EVMTokenAmountsHashPreimage\",\"inputs\":[{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVM2AnyTokenAmountsHashPreimage\",\"inputs\":[{\"name\":\"tokenAmount\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeFinalHashPreimage\",\"inputs\":[{\"name\":\"leafDomainSeparator\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metaDataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fixedSizeFieldsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"senderHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAmountsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeFixedSizeFieldsHashPreimage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeMetadataHashPreimage\",\"inputs\":[{\"name\":\"any2EVMMessageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSUIExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Internal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"struct Internal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Internal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x608080604052346015576112ed908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806310ee47db14610b115780631914fbd214610a7b5780633ec7c377146109b457806355ad01df1461096b5780636fa473e41461090757806381c6b88b146101f25780638503839d146105e857806394b6624b1461035c578063ae5663d7146102c9578063b17df71414610270578063bf0619ad146101fc578063c63641bd146101f7578063c6ba5f28146101f7578063c7ca9a18146101f2578063e04767b81461016b5763e733d209146100cc57600080fd5b346101665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665761016260405161010a81610ba9565b6004358152604051907f97a657c90000000000000000000000000000000000000000000000000000000060208301525160248201526024815261014e604482610be1565b604051918291602083526020830190610d76565b0390f35b600080fd5b346101665760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665760243567ffffffffffffffff8116809103610166576101629067ffffffffffffffff6101c4610de6565b604051926004356020850152604084015216606082015260643560808201526080815261014e60a082610be1565b610f40565b61117f565b346101665760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016657610162604051600435602082015260243560408201526044356060820152606435608082015260843560a082015260a43560c082015260c0815261014e60e082610be1565b346101665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016657602060043560006040516102b281610ba9565b52806040516102c081610ba9565b52604051908152f35b346101665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665760043567ffffffffffffffff81116101665761033061014e61032161016293369060040161105b565b6040519283916020830161120a565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610be1565b346101665760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665760043567ffffffffffffffff81116101665736602382011215610166578060040135906103b782610c2f565b906103c56040519283610be1565b82825260208201906024829460051b820101903682116101665760248101925b8284106104f957858560405190604082019060208084015251809152606082019060608160051b84010193916000905b828210610451576101628561014e8189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610be1565b909192946020806104eb837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0896001960301865289519073ffffffffffffffffffffffffffffffffffffffff825116815260806104d06104be8685015160a08886015260a0850190610d76565b60408501518482036040860152610d76565b92606081015160608401520151906080818403910152610d76565b970192019201909291610415565b833567ffffffffffffffff811161016657820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8236030112610166576040519161054583610b8d565b61055160248301610eeb565b8352604482013567ffffffffffffffff8111610166576105779060243691850101610fe6565b6020840152606482013567ffffffffffffffff8111610166576105a09060243691850101610fe6565b60408401526084820135606084015260a48201359267ffffffffffffffff8411610166576105d8602094936024869536920101610fe6565b60808201528152019301926103e5565b346101665760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665760043567ffffffffffffffff8111610166577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc813603016101408112610166576040519060c082019082821067ffffffffffffffff8311176108d85760a091604052126101665760405161068a81610b8d565b8260040135815261069d60248401610dfd565b60208201526106ae60448401610dfd565b60408201526106bf60648401610dfd565b60608201526106d060848401610dfd565b6080820152815260a482013567ffffffffffffffff8111610166576106fb9060043691850101610fe6565b6020820190815260c483013567ffffffffffffffff8111610166576107269060043691860101610fe6565b6040830190815261073960e48501610eeb565b60608401908152608084019461010481013586526101248101359067ffffffffffffffff8211610166576004610772923692010161105b565b9060a085019182526024359567ffffffffffffffff87116101665761086c6107a06020983690600401610fe6565b87519067ffffffffffffffff6040818c8501511693015116908a8151910120604051918b8301937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f85526040840152606083015260808201526080815261080860a082610be1565b5190209651805193516060808301519451608093840151604080518e8101998a5273ffffffffffffffffffffffffffffffffffffffff9590951660208a015267ffffffffffffffff9788169089015291870152909316908401528160a08401610330565b51902092518581519101209151858151910120905160405161089581610330898201948561120a565b5190209160405193868501956000875260408601526060850152608084015260a083015260c082015260c081526108cd60e082610be1565b519020604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101665761091536610e12565b63ffffffff81511661016267ffffffffffffffff6020840151169260408101511515906080606082015191015191604051958695865260208601526040850152606084015260a0608084015260a0830190610f0c565b346101665761097936610ca4565b805161016260208301511515926060604082015191015190604051948594855260208501526040840152608060608401526080830190610f0c565b346101665760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101665760243573ffffffffffffffffffffffffffffffffffffffff8116810361016657610a0b610de6565b9060843567ffffffffffffffff811681036101665760408051600435602082015273ffffffffffffffffffffffffffffffffffffffff9093169083015267ffffffffffffffff928316606083015260643560808301529190911660a08201526101629061014e8160c08101610330565b3461016657610162608061014e610a9136610e12565b6103306040519384927f1f3b3aba0000000000000000000000000000000000000000000000000000000060208501526020602485015263ffffffff815116604485015267ffffffffffffffff6020820151166064850152604081015115156084850152606081015160a4850152015160a060c484015260e4830190610f0c565b3461016657610162606061014e610b2736610ca4565b6103306040519384927f21ea4ca90000000000000000000000000000000000000000000000000000000060208501526020602485015280516044850152602081015115156064850152604081015160848501520151608060a484015260c4830190610f0c565b60a0810190811067ffffffffffffffff8211176108d857604052565b6020810190811067ffffffffffffffff8211176108d857604052565b6040810190811067ffffffffffffffff8211176108d857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108d857604052565b3590811515820361016657565b67ffffffffffffffff81116108d85760051b60200190565b9080601f83011215610166578135610c5e81610c2f565b92610c6c6040519485610be1565b81845260208085019260051b82010192831161016657602001905b828210610c945750505090565b8135815260209182019101610c87565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101665760043567ffffffffffffffff81116101665760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828403011261016657604051916080830183811067ffffffffffffffff8211176108d85760405281600401358352610d3d60248301610c22565b60208401526044820135604084015260648201359167ffffffffffffffff831161016657610d6e9201600401610c47565b606082015290565b919082519283825260005b848110610dc05750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610d81565b359063ffffffff8216820361016657565b6044359067ffffffffffffffff8216820361016657565b359067ffffffffffffffff8216820361016657565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101665760043567ffffffffffffffff81116101665760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82840301126101665760405191610e8783610b8d565b610e9382600401610dd5565b8352610ea160248301610dfd565b6020840152610eb260448301610c22565b60408401526064820135606084015260848201359167ffffffffffffffff831161016657610ee39201600401610c47565b608082015290565b359073ffffffffffffffffffffffffffffffffffffffff8216820361016657565b906020808351928381520192019060005b818110610f2a5750505090565b8251845260209384019390920191600101610f1d565b3461016657600060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610fe35760405190610f7e82610bc5565b6004358252602435908115158203610fe35760208084018381526040517f181dcf1000000000000000000000000000000000000000000000000000000000928101929092528451602483015251151560448201526101629061014e8160648101610330565b80fd5b81601f820112156101665780359067ffffffffffffffff82116108d8576040519261103960207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601160185610be1565b8284526020838301011161016657816000926020809301838601378301015290565b81601f820112156101665780359061107282610c2f565b926110806040519485610be1565b82845260208085019360051b830101918183116101665760208101935b8385106110ac57505050505090565b843567ffffffffffffffff811161016657820160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261016657604051916110f883610b8d565b602082013567ffffffffffffffff81116101665785602061111b92850101610fe6565b835261112960408301610eeb565b602084015261113a60608301610dd5565b604084015260808201359267ffffffffffffffff84116101665760a083611168886020809881980101610fe6565b60608401520135608082015281520194019361109d565b346101665760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610166576024358015158091036101665761016290600060206040516111cf81610bc5565b8281520152604051906111e182610bc5565b600435825260208201526040519182918291909160208060408301948051845201511515910152565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061123d57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc085600195030186528851906080806112cb61128b855160a0865260a0860190610d76565b73ffffffffffffffffffffffffffffffffffffffff87870151168786015263ffffffff604087015116604086015260608601518582036060870152610d76565b9301519101529701930193019193929061122e56fea164736f6c634300081a000a",
}

var MessageHasherABI = MessageHasherMetaData.ABI

var MessageHasherBin = MessageHasherMetaData.Bin

func DeployMessageHasher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHasher, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageHasherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHasher{address: address, abi: *parsed, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

type MessageHasher struct {
	address common.Address
	abi     abi.ABI
	MessageHasherCaller
	MessageHasherTransactor
	MessageHasherFilterer
}

type MessageHasherCaller struct {
	contract *bind.BoundContract
}

type MessageHasherTransactor struct {
	contract *bind.BoundContract
}

type MessageHasherFilterer struct {
	contract *bind.BoundContract
}

type MessageHasherSession struct {
	Contract     *MessageHasher
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageHasherCallerSession struct {
	Contract *MessageHasherCaller
	CallOpts bind.CallOpts
}

type MessageHasherTransactorSession struct {
	Contract     *MessageHasherTransactor
	TransactOpts bind.TransactOpts
}

type MessageHasherRaw struct {
	Contract *MessageHasher
}

type MessageHasherCallerRaw struct {
	Contract *MessageHasherCaller
}

type MessageHasherTransactorRaw struct {
	Contract *MessageHasherTransactor
}

func NewMessageHasher(address common.Address, backend bind.ContractBackend) (*MessageHasher, error) {
	abi, err := abi.JSON(strings.NewReader(MessageHasherABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageHasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHasher{address: address, abi: abi, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

func NewMessageHasherCaller(address common.Address, caller bind.ContractCaller) (*MessageHasherCaller, error) {
	contract, err := bindMessageHasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherCaller{contract: contract}, nil
}

func NewMessageHasherTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHasherTransactor, error) {
	contract, err := bindMessageHasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherTransactor{contract: contract}, nil
}

func NewMessageHasherFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageHasherFilterer, error) {
	contract, err := bindMessageHasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageHasherFilterer{contract: contract}, nil
}

func bindMessageHasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MessageHasher *MessageHasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.MessageHasherCaller.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCaller) DecodeEVMExtraArgsV1(opts *bind.CallOpts, gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeEVMExtraArgsV1", gasLimit)

	if err != nil {
		return *new(ClientEVMExtraArgsV1), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientEVMExtraArgsV1)).(*ClientEVMExtraArgsV1)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeEVMExtraArgsV1(gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV1(&_MessageHasher.CallOpts, gasLimit)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeEVMExtraArgsV1(gasLimit *big.Int) (ClientEVMExtraArgsV1, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV1(&_MessageHasher.CallOpts, gasLimit)
}

func (_MessageHasher *MessageHasherCaller) DecodeEVMExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeEVMExtraArgsV2", gasLimit, allowOutOfOrderExecution)

	if err != nil {
		return *new(ClientGenericExtraArgsV2), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientGenericExtraArgsV2)).(*ClientGenericExtraArgsV2)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeEVMExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeEVMExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeEVMExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCaller) DecodeGenericExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeGenericExtraArgsV2", gasLimit, allowOutOfOrderExecution)

	if err != nil {
		return *new(ClientGenericExtraArgsV2), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientGenericExtraArgsV2)).(*ClientGenericExtraArgsV2)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) DecodeGenericExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeGenericExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeGenericExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error) {
	return _MessageHasher.Contract.DecodeGenericExtraArgsV2(&_MessageHasher.CallOpts, gasLimit, allowOutOfOrderExecution)
}

func (_MessageHasher *MessageHasherCaller) DecodeSVMExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeSVMExtraArgsStruct", extraArgs)

	outstruct := new(DecodeSVMExtraArgsStruct)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ComputeUnits = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AccountIsWritableBitmap = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AllowOutOfOrderExecution = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.TokenReceiver = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Accounts = *abi.ConvertType(out[4], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_MessageHasher *MessageHasherSession) DecodeSVMExtraArgsStruct(extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSVMExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeSVMExtraArgsStruct(extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSVMExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) DecodeSuiExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "decodeSuiExtraArgsStruct", extraArgs)

	outstruct := new(DecodeSuiExtraArgsStruct)
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AllowOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TokenReceiver = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.ReceiverObjectIds = *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_MessageHasher *MessageHasherSession) DecodeSuiExtraArgsStruct(extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSuiExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) DecodeSuiExtraArgsStruct(extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

	error) {
	return _MessageHasher.Contract.DecodeSuiExtraArgsStruct(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeAny2EVMTokenAmountsHashPreimage(opts *bind.CallOpts, tokenAmounts []InternalAny2EVMTokenTransfer) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeAny2EVMTokenAmountsHashPreimage", tokenAmounts)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeAny2EVMTokenAmountsHashPreimage(tokenAmounts []InternalAny2EVMTokenTransfer) ([]byte, error) {
	return _MessageHasher.Contract.EncodeAny2EVMTokenAmountsHashPreimage(&_MessageHasher.CallOpts, tokenAmounts)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeAny2EVMTokenAmountsHashPreimage(tokenAmounts []InternalAny2EVMTokenTransfer) ([]byte, error) {
	return _MessageHasher.Contract.EncodeAny2EVMTokenAmountsHashPreimage(&_MessageHasher.CallOpts, tokenAmounts)
}

func (_MessageHasher *MessageHasherCaller) EncodeEVM2AnyTokenAmountsHashPreimage(opts *bind.CallOpts, tokenAmount []InternalEVM2AnyTokenTransfer) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeEVM2AnyTokenAmountsHashPreimage", tokenAmount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeEVM2AnyTokenAmountsHashPreimage(tokenAmount []InternalEVM2AnyTokenTransfer) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVM2AnyTokenAmountsHashPreimage(&_MessageHasher.CallOpts, tokenAmount)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeEVM2AnyTokenAmountsHashPreimage(tokenAmount []InternalEVM2AnyTokenTransfer) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVM2AnyTokenAmountsHashPreimage(&_MessageHasher.CallOpts, tokenAmount)
}

func (_MessageHasher *MessageHasherCaller) EncodeEVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeEVMExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeEVMExtraArgsV1(extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeEVMExtraArgsV1(extraArgs ClientEVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeEVMExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeEVMExtraArgsV2", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeEVMExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeEVMExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeEVMExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeFinalHashPreimage(opts *bind.CallOpts, leafDomainSeparator [32]byte, metaDataHash [32]byte, fixedSizeFieldsHash [32]byte, senderHash [32]byte, dataHash [32]byte, tokenAmountsHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeFinalHashPreimage", leafDomainSeparator, metaDataHash, fixedSizeFieldsHash, senderHash, dataHash, tokenAmountsHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeFinalHashPreimage(leafDomainSeparator [32]byte, metaDataHash [32]byte, fixedSizeFieldsHash [32]byte, senderHash [32]byte, dataHash [32]byte, tokenAmountsHash [32]byte) ([]byte, error) {
	return _MessageHasher.Contract.EncodeFinalHashPreimage(&_MessageHasher.CallOpts, leafDomainSeparator, metaDataHash, fixedSizeFieldsHash, senderHash, dataHash, tokenAmountsHash)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeFinalHashPreimage(leafDomainSeparator [32]byte, metaDataHash [32]byte, fixedSizeFieldsHash [32]byte, senderHash [32]byte, dataHash [32]byte, tokenAmountsHash [32]byte) ([]byte, error) {
	return _MessageHasher.Contract.EncodeFinalHashPreimage(&_MessageHasher.CallOpts, leafDomainSeparator, metaDataHash, fixedSizeFieldsHash, senderHash, dataHash, tokenAmountsHash)
}

func (_MessageHasher *MessageHasherCaller) EncodeFixedSizeFieldsHashPreimage(opts *bind.CallOpts, messageId [32]byte, receiver common.Address, sequenceNumber uint64, gasLimit *big.Int, nonce uint64) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeFixedSizeFieldsHashPreimage", messageId, receiver, sequenceNumber, gasLimit, nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeFixedSizeFieldsHashPreimage(messageId [32]byte, receiver common.Address, sequenceNumber uint64, gasLimit *big.Int, nonce uint64) ([]byte, error) {
	return _MessageHasher.Contract.EncodeFixedSizeFieldsHashPreimage(&_MessageHasher.CallOpts, messageId, receiver, sequenceNumber, gasLimit, nonce)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeFixedSizeFieldsHashPreimage(messageId [32]byte, receiver common.Address, sequenceNumber uint64, gasLimit *big.Int, nonce uint64) ([]byte, error) {
	return _MessageHasher.Contract.EncodeFixedSizeFieldsHashPreimage(&_MessageHasher.CallOpts, messageId, receiver, sequenceNumber, gasLimit, nonce)
}

func (_MessageHasher *MessageHasherCaller) EncodeGenericExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeGenericExtraArgsV2", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeGenericExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeGenericExtraArgsV2(extraArgs ClientGenericExtraArgsV2) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV2(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeMetadataHashPreimage(opts *bind.CallOpts, any2EVMMessageHash [32]byte, sourceChainSelector uint64, destChainSelector uint64, onRampHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeMetadataHashPreimage", any2EVMMessageHash, sourceChainSelector, destChainSelector, onRampHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeMetadataHashPreimage(any2EVMMessageHash [32]byte, sourceChainSelector uint64, destChainSelector uint64, onRampHash [32]byte) ([]byte, error) {
	return _MessageHasher.Contract.EncodeMetadataHashPreimage(&_MessageHasher.CallOpts, any2EVMMessageHash, sourceChainSelector, destChainSelector, onRampHash)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeMetadataHashPreimage(any2EVMMessageHash [32]byte, sourceChainSelector uint64, destChainSelector uint64, onRampHash [32]byte) ([]byte, error) {
	return _MessageHasher.Contract.EncodeMetadataHashPreimage(&_MessageHasher.CallOpts, any2EVMMessageHash, sourceChainSelector, destChainSelector, onRampHash)
}

func (_MessageHasher *MessageHasherCaller) EncodeSUIExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeSUIExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeSUIExtraArgsV1(extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSUIExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeSUIExtraArgsV1(extraArgs ClientSuiExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSUIExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) EncodeSVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeSVMExtraArgsV1", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeSVMExtraArgsV1(extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeSVMExtraArgsV1(extraArgs ClientSVMExtraArgsV1) ([]byte, error) {
	return _MessageHasher.Contract.EncodeSVMExtraArgsV1(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCaller) Hash(opts *bind.CallOpts, message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "hash", message, onRamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) Hash(message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash(&_MessageHasher.CallOpts, message, onRamp)
}

func (_MessageHasher *MessageHasherCallerSession) Hash(message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash(&_MessageHasher.CallOpts, message, onRamp)
}

type DecodeSVMExtraArgsStruct struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}
type DecodeSuiExtraArgsStruct struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	ReceiverObjectIds        [][32]byte
}

func (_MessageHasher *MessageHasher) Address() common.Address {
	return _MessageHasher.address
}

type MessageHasherInterface interface {
	DecodeEVMExtraArgsV1(opts *bind.CallOpts, gasLimit *big.Int) (ClientEVMExtraArgsV1, error)

	DecodeEVMExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error)

	DecodeGenericExtraArgsV2(opts *bind.CallOpts, gasLimit *big.Int, allowOutOfOrderExecution bool) (ClientGenericExtraArgsV2, error)

	DecodeSVMExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) (DecodeSVMExtraArgsStruct,

		error)

	DecodeSuiExtraArgsStruct(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) (DecodeSuiExtraArgsStruct,

		error)

	EncodeAny2EVMTokenAmountsHashPreimage(opts *bind.CallOpts, tokenAmounts []InternalAny2EVMTokenTransfer) ([]byte, error)

	EncodeEVM2AnyTokenAmountsHashPreimage(opts *bind.CallOpts, tokenAmount []InternalEVM2AnyTokenTransfer) ([]byte, error)

	EncodeEVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV1) ([]byte, error)

	EncodeEVMExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeFinalHashPreimage(opts *bind.CallOpts, leafDomainSeparator [32]byte, metaDataHash [32]byte, fixedSizeFieldsHash [32]byte, senderHash [32]byte, dataHash [32]byte, tokenAmountsHash [32]byte) ([]byte, error)

	EncodeFixedSizeFieldsHashPreimage(opts *bind.CallOpts, messageId [32]byte, receiver common.Address, sequenceNumber uint64, gasLimit *big.Int, nonce uint64) ([]byte, error)

	EncodeGenericExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeMetadataHashPreimage(opts *bind.CallOpts, any2EVMMessageHash [32]byte, sourceChainSelector uint64, destChainSelector uint64, onRampHash [32]byte) ([]byte, error)

	EncodeSUIExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) ([]byte, error)

	EncodeSVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) ([]byte, error)

	Hash(opts *bind.CallOpts, message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error)

	Address() common.Address
}
