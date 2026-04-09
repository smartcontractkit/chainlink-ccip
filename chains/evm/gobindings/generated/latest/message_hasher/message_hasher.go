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

type ExtraArgsCodecGenericExtraArgsV3 struct {
	GasLimit                uint32
	RequestedFinalityConfig [4]byte
	Ccvs                    []common.Address
	CcvArgs                 [][]byte
	Executor                common.Address
	ExecutorArgs            []byte
	TokenReceiver           []byte
	TokenArgs               []byte
}

var MessageHasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSVMExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decodeSuiExtraArgsStruct\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeBlockDepth\",\"inputs\":[{\"name\":\"blockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeBlockDepthAndSafeFlag\",\"inputs\":[{\"name\":\"blockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVMExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeEVMExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeGenericExtraArgsV2\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.GenericExtraArgsV2\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeGenericExtraArgsV3\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct ExtraArgsCodec.GenericExtraArgsV3\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSUIExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SuiExtraArgsV1\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiverObjectIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeSVMExtraArgsV1\",\"inputs\":[{\"name\":\"extraArgs\",\"type\":\"tuple\",\"internalType\":\"struct Client.SVMExtraArgsV1\",\"components\":[{\"name\":\"computeUnits\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountIsWritableBitmap\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowOutOfOrderExecution\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"ensureRequestedFinalityAllowed\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validateRequestedFinality\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"CCVArrayLengthMismatch\",\"inputs\":[{\"name\":\"ccvsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ccvArgsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]}]",
	Bin: "0x60808060405234601557611761908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806310ee47db146100fd5780631914fbd2146100f857806355ad01df146100f35780636fa473e4146100ee57806381c6b88b146100d057806392515ab0146100e9578063a1b581a1146100e4578063af60bc1d146100df578063b17df714146100da578063c63641bd146100d5578063c6ba5f28146100d5578063c7ca9a18146100d0578063e1ff618c146100cb578063e733d209146100c65763fea2e8f6146100c157600080fd5b610c86565b610a7e565b610a12565b6106c8565b610986565b61092d565b6108ca565b610846565b6107ca565b610664565b61061b565b610555565b6103a9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761014d57604052565b610102565b60a0810190811067ffffffffffffffff82111761014d57604052565b6040810190811067ffffffffffffffff82111761014d57604052565b6020810190811067ffffffffffffffff82111761014d57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761014d57604052565b604051906101f7610100836101a6565b565b3590811515820361020657565b600080fd5b67ffffffffffffffff811161014d5760051b60200190565b9080601f8301121561020657813561023a8161020b565b9261024860405194856101a6565b81845260208085019260051b82010192831161020657602001905b8282106102705750505090565b8135815260209182019101610263565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102065760043567ffffffffffffffff81116102065760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828403011261020657604051916102f583610131565b81600401358352610308602483016101f9565b60208401526044820135604084015260648201359167ffffffffffffffff8311610206576103399201600401610223565b606082015290565b9190916020815282519283602083015260005b8481106103935750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201610354565b346102065761045d60606104516103bf36610280565b6104256040519384927f21ea4ca90000000000000000000000000000000000000000000000000000000060208501526020602485015280516044850152602081015115156064850152604081015160848501520151608060a484015260c48301906105e7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826101a6565b60405191829182610341565b0390f35b359063ffffffff8216820361020657565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102065760043567ffffffffffffffff81116102065760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc828403011261020657604051916104e783610152565b6104f382600401610461565b8352602482013567ffffffffffffffff8116810361020657602084015261051c604483016101f9565b60408401526064820135606084015260848201359167ffffffffffffffff83116102065761054d9201600401610223565b608082015290565b346102065761045d608061045161056b36610472565b6104256040519384927f1f3b3aba0000000000000000000000000000000000000000000000000000000060208501526020602485015263ffffffff815116604485015267ffffffffffffffff6020820151166064850152604081015115156084850152606081015160a4850152015160a060c484015260e48301905b906020808351928381520192019060005b8181106106055750505090565b82518452602093840193909201916001016105f8565b346102065761062936610280565b805161045d602083015115159260606040820151910151906040519485948552602085015260408401526080606084015260808301906105e7565b346102065761067236610472565b63ffffffff81511661045d67ffffffffffffffff6020840151169260408101511515906080606082015191015191604051958695865260208601526040850152606084015260a0608084015260a08301906105e7565b3461020657600060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261076b57604051906107068261016e565b600435825260243590811515820361076b5760208084018381526040517f181dcf10000000000000000000000000000000000000000000000000000000009281019290925284516024830152511515604482015261045d906104518160648101610425565b80fd5b600435907fffffffff000000000000000000000000000000000000000000000000000000008216820361020657565b35907fffffffff000000000000000000000000000000000000000000000000000000008216820361020657565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065761080961080461076e565b610e8d565b005b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60209101126102065760043561ffff811681036102065790565b346102065760207fffffffff0000000000000000000000000000000000000000000000000000000061ffff61087a3661080b565b1660e01b167fffffffff000000000000000000000000000000000000000000000000000000007e010000000000000000000000000000000000000000000000000000000000006040519217168152f35b346102065760207fffffffff0000000000000000000000000000000000000000000000000000000061ffff6108fe3661080b565b1660e01b167fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576020600435600060405161096f8161018a565b528060405161097d8161018a565b52604051908152f35b346102065760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657600435602435908115158092036102065761045d91600060206040516109da8161016e565b8281520152604051916109ec8361016e565b825260208201526040519182918291909160208060408301948051845201511515910152565b346102065760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657610a4961076e565b6024357fffffffff00000000000000000000000000000000000000000000000000000000811681036102065761080991610f94565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065761045d604051610abc8161018a565b6004358152604051907f97a657c9000000000000000000000000000000000000000000000000000000006020830152516024820152602481526104516044826101a6565b359073ffffffffffffffffffffffffffffffffffffffff8216820361020657565b9080601f83011215610206578135610b388161020b565b92610b4660405194856101a6565b81845260208085019260051b82010192831161020657602001905b828210610b6e5750505090565b60208091610b7b84610b00565b815201910190610b61565b67ffffffffffffffff811161014d57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561020657803590610bd782610b86565b92610be560405194856101a6565b8284526020838301011161020657816000926020809301838601378301015290565b9080601f83011215610206578135610c1e8161020b565b92610c2c60405194856101a6565b81845260208085019260051b820101918383116102065760208201905b838210610c5857505050505090565b813567ffffffffffffffff811161020657602091610c7b87848094880101610bc0565b815201910190610c49565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff8111610206576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261020657610cfc6101e7565b90610d0981600401610461565b8252610d176024820161079d565b6020830152604481013567ffffffffffffffff811161020657610d409060043691840101610b21565b6040830152606481013567ffffffffffffffff811161020657610d699060043691840101610c07565b6060830152610d7a60848201610b00565b608083015260a481013567ffffffffffffffff811161020657610da39060043691840101610bc0565b60a083015260c481013567ffffffffffffffff811161020657610dcc9060043691840101610bc0565b60c083015260e48101359167ffffffffffffffff831161020657610dfc61045192600461045d9536920101610bc0565b60e08201526111ba565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9060018201809211610e4357565b610e06565b6001019081600111610e4357565b9060028201809211610e4357565b6013019081601311610e4357565b9060208201809211610e4357565b91908201809211610e4357565b7fffffffff00000000000000000000000000000000000000000000000000000000811615610f91577dffff00000000000000000000000000000000000000000000000000000000811615610f885760ff60015b1660f082901c80610f4a575b50600103610ef75750565b7fc512f96c000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b6000fd5b60005b60108110610f5b5750610eec565b63ffffffff6001821b831616610f74575b600101610f4d565b91610f80600191610e35565b929050610f6c565b60ff6000610ee0565b50565b7fffffffff0000000000000000000000000000000000000000000000000000000081161561107457610fc581610e8d565b601082811c9082901c167dffff00000000000000000000000000000000000000000000000000000000166110745761ffff8260e01c168015908115611063575b5061100e575050565b7fdf63778f000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000009081166004521660245260446000fd5b905061ffff8260e01c161038611005565b5050565b90600f8210156110855752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6044906110c360046008611078565b6000602452565b6044906110c36004600a611078565b6044906110c36004600b611078565b6044906110c36004600c611078565b6044906110c360046009611078565b805182101561111a5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9061115382610b86565b61116060405191826101a6565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061118e8294610b86565b0190602036910137565b91908203918211610e4357565b906044916111b560046003611078565b602452565b90604082019182515160608201805151808303611626575060ff82116115fa5760a083018051519161ffff83116115ce5760c08501928351519460ff86116115a25760e08701958651519061ffff821161157657608089019373ffffffffffffffffffffffffffffffffffffffff611246865173ffffffffffffffffffffffffffffffffffffffff1690565b1661156c5760009593955b600090818e5b8982106114bb57505061128d9361128361128894611283611292989561128360ff611283971691610e64565b610e80565b610e72565b611149565b9661139861136c6112d460206112ac855163ffffffff1690565b9401517fffffffff000000000000000000000000000000000000000000000000000000001690565b6040517fa69dd4aa000000000000000000000000000000000000000000000000000000006020820190815260e09590951b7fffffffff000000000000000000000000000000000000000000000000000000009081166024830152909116602882015260f887901b7fff0000000000000000000000000000000000000000000000000000000000000016602c820152908190602d820190565b03907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0820181526101a6565b516020880152602d8701926000915b81831061146b575050506113f895969750916113e86113f0926113e26113e896955173ffffffffffffffffffffffffffffffffffffffff1690565b90611658565b90519061167c565b9051906116ee565b61142b6112887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084510180855284610e80565b8103611435575090565b610f469161144291611198565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526111a5565b9091936114b26114a06001926113e26114868f8a9051611106565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6114ab878551611106565b519061167c565b940191906113a7565b6114d2611486836114eb939c9a9c96949651611106565b73ffffffffffffffffffffffffffffffffffffffff1690565b611564576000905b6114fe848b51611106565b51519061ffff82116115385761152761152d9261128361152260ff60019716610e48565b610e56565b90610e80565b92018e989698611257565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052610f466110f7565b6014906114f3565b6014959395611251565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052610f466110e8565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052610f466110d9565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052610f466110ca565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052610f466110b4565b7f21dbdf3d00000000000000000000000000000000000000000000000000000000600052600483905260245260446000fd5b60148215150291828253600182019261167057505090565b60601b90915260150190565b60028251918260081c815360ff8316600182015301918161169c57505090565b6020828183019201015b8082106116db5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f82011603900390565b90926020809185518152019301906116a6565b600182519182815301918161170257505090565b6020828183019201015b8082106117415750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f82011603900390565b909260208091855181520193019061170c56fea164736f6c634300081a000a",
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

func (_MessageHasher *MessageHasherCaller) EncodeBlockDepth(opts *bind.CallOpts, blockDepth uint16) ([4]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeBlockDepth", blockDepth)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeBlockDepth(blockDepth uint16) ([4]byte, error) {
	return _MessageHasher.Contract.EncodeBlockDepth(&_MessageHasher.CallOpts, blockDepth)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeBlockDepth(blockDepth uint16) ([4]byte, error) {
	return _MessageHasher.Contract.EncodeBlockDepth(&_MessageHasher.CallOpts, blockDepth)
}

func (_MessageHasher *MessageHasherCaller) EncodeBlockDepthAndSafeFlag(opts *bind.CallOpts, blockDepth uint16) ([4]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeBlockDepthAndSafeFlag", blockDepth)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeBlockDepthAndSafeFlag(blockDepth uint16) ([4]byte, error) {
	return _MessageHasher.Contract.EncodeBlockDepthAndSafeFlag(&_MessageHasher.CallOpts, blockDepth)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeBlockDepthAndSafeFlag(blockDepth uint16) ([4]byte, error) {
	return _MessageHasher.Contract.EncodeBlockDepthAndSafeFlag(&_MessageHasher.CallOpts, blockDepth)
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

func (_MessageHasher *MessageHasherCaller) EncodeGenericExtraArgsV3(opts *bind.CallOpts, extraArgs ExtraArgsCodecGenericExtraArgsV3) ([]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "encodeGenericExtraArgsV3", extraArgs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) EncodeGenericExtraArgsV3(extraArgs ExtraArgsCodecGenericExtraArgsV3) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV3(&_MessageHasher.CallOpts, extraArgs)
}

func (_MessageHasher *MessageHasherCallerSession) EncodeGenericExtraArgsV3(extraArgs ExtraArgsCodecGenericExtraArgsV3) ([]byte, error) {
	return _MessageHasher.Contract.EncodeGenericExtraArgsV3(&_MessageHasher.CallOpts, extraArgs)
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

func (_MessageHasher *MessageHasherCaller) EnsureRequestedFinalityAllowed(opts *bind.CallOpts, requestedFinality [4]byte, allowedFinality [4]byte) error {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "ensureRequestedFinalityAllowed", requestedFinality, allowedFinality)

	if err != nil {
		return err
	}

	return err

}

func (_MessageHasher *MessageHasherSession) EnsureRequestedFinalityAllowed(requestedFinality [4]byte, allowedFinality [4]byte) error {
	return _MessageHasher.Contract.EnsureRequestedFinalityAllowed(&_MessageHasher.CallOpts, requestedFinality, allowedFinality)
}

func (_MessageHasher *MessageHasherCallerSession) EnsureRequestedFinalityAllowed(requestedFinality [4]byte, allowedFinality [4]byte) error {
	return _MessageHasher.Contract.EnsureRequestedFinalityAllowed(&_MessageHasher.CallOpts, requestedFinality, allowedFinality)
}

func (_MessageHasher *MessageHasherCaller) ValidateRequestedFinality(opts *bind.CallOpts, encodedFinality [4]byte) error {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "validateRequestedFinality", encodedFinality)

	if err != nil {
		return err
	}

	return err

}

func (_MessageHasher *MessageHasherSession) ValidateRequestedFinality(encodedFinality [4]byte) error {
	return _MessageHasher.Contract.ValidateRequestedFinality(&_MessageHasher.CallOpts, encodedFinality)
}

func (_MessageHasher *MessageHasherCallerSession) ValidateRequestedFinality(encodedFinality [4]byte) error {
	return _MessageHasher.Contract.ValidateRequestedFinality(&_MessageHasher.CallOpts, encodedFinality)
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

	EncodeBlockDepth(opts *bind.CallOpts, blockDepth uint16) ([4]byte, error)

	EncodeBlockDepthAndSafeFlag(opts *bind.CallOpts, blockDepth uint16) ([4]byte, error)

	EncodeEVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientEVMExtraArgsV1) ([]byte, error)

	EncodeEVMExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeGenericExtraArgsV2(opts *bind.CallOpts, extraArgs ClientGenericExtraArgsV2) ([]byte, error)

	EncodeGenericExtraArgsV3(opts *bind.CallOpts, extraArgs ExtraArgsCodecGenericExtraArgsV3) ([]byte, error)

	EncodeSUIExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSuiExtraArgsV1) ([]byte, error)

	EncodeSVMExtraArgsV1(opts *bind.CallOpts, extraArgs ClientSVMExtraArgsV1) ([]byte, error)

	EnsureRequestedFinalityAllowed(opts *bind.CallOpts, requestedFinality [4]byte, allowedFinality [4]byte) error

	ValidateRequestedFinality(opts *bind.CallOpts, encodedFinality [4]byte) error

	Address() common.Address
}
