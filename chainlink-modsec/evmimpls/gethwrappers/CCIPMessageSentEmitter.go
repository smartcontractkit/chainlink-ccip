// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gethwrappers

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

// Reference imports to suppress errors if they are not otherwise used.
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

// CCIPMessageSentEmitterAny2EVMMultiProofMessage is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterAny2EVMMultiProofMessage struct {
	Header            CCIPMessageSentEmitterHeader
	Sender            []byte
	Data              []byte
	Receiver          common.Address
	GasLimit          uint32
	TokenAmounts      []CCIPMessageSentEmitterAny2EVMMultiProofTokenTransfer
	RequiredVerifiers [][]byte
}

// CCIPMessageSentEmitterAny2EVMMultiProofTokenTransfer is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterAny2EVMMultiProofTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	ExtraData         []byte
	Amount            *big.Int
}

// CCIPMessageSentEmitterEVM2AnyVerifierMessage is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterEVM2AnyVerifierMessage struct {
	Header         CCIPMessageSentEmitterHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenTransfer  CCIPMessageSentEmitterEVMTokenTransfer
	Receipts       []CCIPMessageSentEmitterReceipt
}

// CCIPMessageSentEmitterEVMTokenTransfer is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	SourcePoolAddress  common.Address
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
	DestExecData       []byte
	RequiredVerifierId [32]byte
}

// CCIPMessageSentEmitterHeader is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

// CCIPMessageSentEmitterReceipt is an auto generated low-level Go binding around an user-defined struct.
type CCIPMessageSentEmitterReceipt struct {
	ReceiptType       uint8
	Issuer            common.Address
	FeeTokenAmount    *big.Int
	DestGasLimit      uint64
	DestBytesOverhead uint32
	ExtraArgs         []byte
}

// CCIPMessageSentEmitterMetaData contains all meta data concerning the CCIPMessageSentEmitter contract.
var CCIPMessageSentEmitterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"requiredVerifierId\",\"type\":\"bytes32\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVMTokenTransfer\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumCCIPMessageSentEmitter.ReceiptType\",\"name\":\"receiptType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPMessageSentEmitter.Receipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structCCIPMessageSentEmitter.EVM2AnyVerifierMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPMessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofTokenTransfer[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"proofs\",\"type\":\"bytes[]\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"requiredVerifierId\",\"type\":\"bytes32\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVMTokenTransfer\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumCCIPMessageSentEmitter.ReceiptType\",\"name\":\"receiptType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPMessageSentEmitter.Receipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVM2AnyVerifierMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"emitCCIPMessageSent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofTokenTransfer[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"proofs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"isExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_nonces\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"setExecuted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611a0c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637f5128291161005b5780637f51282914610125578063964ce3a314610155578063c4e0519b14610185578063c8d61748146101a157610088565b8063023993901461008d5780630d7496ee146100a957806322bd4080146100c557806374065f33146100f5575b600080fd5b6100a760048036038101906100a29190610529565b6101bd565b005b6100c360048036038101906100be9190610be8565b610226565b005b6100df60048036038101906100da9190610c60565b610263565b6040516100ec9190610caf565b60405180910390f35b61010f600480360381019061010a9190610cca565b6102c7565b60405161011c9190610d19565b60405180910390f35b61013f600480360381019061013a9190610cca565b6102fd565b60405161014c9190610d19565b60405180910390f35b61016f600480360381019061016a9190610c60565b61037f565b60405161017c9190610caf565b60405180910390f35b61019f600480360381019061019a9190611151565b6103ae565b005b6101bb60048036038101906101b6919061119a565b61040e565b005b80600160008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b7fb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a8282604051610257929190611634565b60405180910390a15050565b6000600160008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006020528160005260406000206020528060005260406000206000915091509054906101000a900467ffffffffffffffff1681565b60008060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16905092915050565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b80600001516060015167ffffffffffffffff1681600001516040015167ffffffffffffffff167fb96f6686b286eb9b9f7114d0d8d35921e425c4fbf84733ac3ea67912d8320ce18360405161040391906119b4565b60405180910390a350565b806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b6104ce816104b1565b81146104d957600080fd5b50565b6000813590506104eb816104c5565b92915050565b60008115159050919050565b610506816104f1565b811461051157600080fd5b50565b600081359050610523816104fd565b92915050565b600080600060608486031215610542576105416104a7565b5b6000610550868287016104dc565b9350506020610561868287016104dc565b925050604061057286828701610514565b9150509250925092565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105ca82610581565b810181811067ffffffffffffffff821117156105e9576105e8610592565b5b80604052505050565b60006105fc61049d565b905061060882826105c1565b919050565b600080fd5b6000819050919050565b61062581610612565b811461063057600080fd5b50565b6000813590506106428161061c565b92915050565b60006080828403121561065e5761065d61057c565b5b61066860806105f2565b9050600061067884828501610633565b600083015250602061068c848285016104dc565b60208301525060406106a0848285016104dc565b60408301525060606106b4848285016104dc565b60608301525092915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156106e5576106e4610592565b5b6106ee82610581565b9050602081019050919050565b82818337600083830152505050565b600061071d610718846106ca565b6105f2565b905082815260208101848484011115610739576107386106c5565b5b6107448482856106fb565b509392505050565b600082601f830112610761576107606106c0565b5b813561077184826020860161070a565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107a58261077a565b9050919050565b6107b58161079a565b81146107c057600080fd5b50565b6000813590506107d2816107ac565b92915050565b600063ffffffff82169050919050565b6107f1816107d8565b81146107fc57600080fd5b50565b60008135905061080e816107e8565b92915050565b600067ffffffffffffffff82111561082f5761082e610592565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b61085881610845565b811461086357600080fd5b50565b6000813590506108758161084f565b92915050565b6000608082840312156108915761089061057c565b5b61089b60806105f2565b9050600082013567ffffffffffffffff8111156108bb576108ba61060d565b5b6108c78482850161074c565b60008301525060206108db848285016107c3565b602083015250604082013567ffffffffffffffff8111156108ff576108fe61060d565b5b61090b8482850161074c565b604083015250606061091f84828501610866565b60608301525092915050565b600061093e61093984610814565b6105f2565b9050808382526020820190506020840283018581111561096157610960610840565b5b835b818110156109a857803567ffffffffffffffff811115610986576109856106c0565b5b808601610993898261087b565b85526020850194505050602081019050610963565b5050509392505050565b600082601f8301126109c7576109c66106c0565b5b81356109d784826020860161092b565b91505092915050565b600067ffffffffffffffff8211156109fb576109fa610592565b5b602082029050602081019050919050565b6000610a1f610a1a846109e0565b6105f2565b90508083825260208201905060208402830185811115610a4257610a41610840565b5b835b81811015610a8957803567ffffffffffffffff811115610a6757610a666106c0565b5b808601610a74898261074c565b85526020850194505050602081019050610a44565b5050509392505050565b600082601f830112610aa857610aa76106c0565b5b8135610ab8848260208601610a0c565b91505092915050565b60006101408284031215610ad857610ad761057c565b5b610ae260e06105f2565b90506000610af284828501610648565b600083015250608082013567ffffffffffffffff811115610b1657610b1561060d565b5b610b228482850161074c565b60208301525060a082013567ffffffffffffffff811115610b4657610b4561060d565b5b610b528482850161074c565b60408301525060c0610b66848285016107c3565b60608301525060e0610b7a848285016107ff565b60808301525061010082013567ffffffffffffffff811115610b9f57610b9e61060d565b5b610bab848285016109b2565b60a08301525061012082013567ffffffffffffffff811115610bd057610bcf61060d565b5b610bdc84828501610a93565b60c08301525092915050565b60008060408385031215610bff57610bfe6104a7565b5b600083013567ffffffffffffffff811115610c1d57610c1c6104ac565b5b610c2985828601610ac1565b925050602083013567ffffffffffffffff811115610c4a57610c496104ac565b5b610c5685828601610a93565b9150509250929050565b60008060408385031215610c7757610c766104a7565b5b6000610c85858286016104dc565b9250506020610c96858286016104dc565b9150509250929050565b610ca9816104f1565b82525050565b6000602082019050610cc46000830184610ca0565b92915050565b60008060408385031215610ce157610ce06104a7565b5b6000610cef858286016104dc565b9250506020610d00858286016107c3565b9150509250929050565b610d13816104b1565b82525050565b6000602082019050610d2e6000830184610d0a565b92915050565b600060e08284031215610d4a57610d4961057c565b5b610d5460e06105f2565b90506000610d64848285016107c3565b6000830152506020610d78848285016107c3565b602083015250604082013567ffffffffffffffff811115610d9c57610d9b61060d565b5b610da88482850161074c565b604083015250606082013567ffffffffffffffff811115610dcc57610dcb61060d565b5b610dd88482850161074c565b6060830152506080610dec84828501610866565b60808301525060a082013567ffffffffffffffff811115610e1057610e0f61060d565b5b610e1c8482850161074c565b60a08301525060c0610e3084828501610633565b60c08301525092915050565b600067ffffffffffffffff821115610e5757610e56610592565b5b602082029050602081019050919050565b60028110610e7557600080fd5b50565b600081359050610e8781610e68565b92915050565b600060c08284031215610ea357610ea261057c565b5b610ead60c06105f2565b90506000610ebd84828501610e78565b6000830152506020610ed1848285016107c3565b6020830152506040610ee584828501610866565b6040830152506060610ef9848285016104dc565b6060830152506080610f0d848285016107ff565b60808301525060a082013567ffffffffffffffff811115610f3157610f3061060d565b5b610f3d8482850161074c565b60a08301525092915050565b6000610f5c610f5784610e3c565b6105f2565b90508083825260208201905060208402830185811115610f7f57610f7e610840565b5b835b81811015610fc657803567ffffffffffffffff811115610fa457610fa36106c0565b5b808601610fb18982610e8d565b85526020850194505050602081019050610f81565b5050509392505050565b600082601f830112610fe557610fe46106c0565b5b8135610ff5848260208601610f49565b91505092915050565b600061018082840312156110155761101461057c565b5b6110206101206105f2565b9050600061103084828501610648565b6000830152506080611044848285016107c3565b60208301525060a082013567ffffffffffffffff8111156110685761106761060d565b5b6110748482850161074c565b60408301525060c082013567ffffffffffffffff8111156110985761109761060d565b5b6110a48482850161074c565b60608301525060e06110b8848285016107c3565b6080830152506101006110cd84828501610866565b60a0830152506101206110e284828501610866565b60c08301525061014082013567ffffffffffffffff8111156111075761110661060d565b5b61111384828501610d34565b60e08301525061016082013567ffffffffffffffff8111156111385761113761060d565b5b61114484828501610fd0565b6101008301525092915050565b600060208284031215611167576111666104a7565b5b600082013567ffffffffffffffff811115611185576111846104ac565b5b61119184828501610ffe565b91505092915050565b6000806000606084860312156111b3576111b26104a7565b5b60006111c1868287016104dc565b93505060206111d2868287016107c3565b92505060406111e3868287016104dc565b9150509250925092565b6111f681610612565b82525050565b611205816104b1565b82525050565b60808201600082015161122160008501826111ed565b50602082015161123460208501826111fc565b50604082015161124760408501826111fc565b50606082015161125a60608501826111fc565b50505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561129a57808201518184015260208101905061127f565b60008484015250505050565b60006112b182611260565b6112bb818561126b565b93506112cb81856020860161127c565b6112d481610581565b840191505092915050565b6112e88161079a565b82525050565b6112f7816107d8565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61133281610845565b82525050565b6000608083016000830151848203600086015261135582826112a6565b915050602083015161136a60208601826112df565b506040830151848203604086015261138282826112a6565b91505060608301516113976060860182611329565b508091505092915050565b60006113ae8383611338565b905092915050565b6000602082019050919050565b60006113ce826112fd565b6113d88185611308565b9350836020820285016113ea85611319565b8060005b85811015611426578484038952815161140785826113a2565b9450611412836113b6565b925060208a019950506001810190506113ee565b50829750879550505050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061147083836112a6565b905092915050565b6000602082019050919050565b600061149082611438565b61149a8185611443565b9350836020820285016114ac85611454565b8060005b858110156114e857848403895281516114c98582611464565b94506114d483611478565b925060208a019950506001810190506114b0565b50829750879550505050505092915050565b600061014083016000830151611513600086018261120b565b506020830151848203608086015261152b82826112a6565b915050604083015184820360a086015261154582826112a6565b915050606083015161155a60c08601826112df565b50608083015161156d60e08601826112ee565b5060a083015184820361010086015261158682826113c3565b91505060c08301518482036101208601526115a18282611485565b9150508091505092915050565b600082825260208201905092915050565b60006115ca82611438565b6115d481856115ae565b9350836020820285016115e685611454565b8060005b8581101561162257848403895281516116038582611464565b945061160e83611478565b925060208a019950506001810190506115ea565b50829750879550505050505092915050565b6000604082019050818103600083015261164e81856114fa565b9050818103602083015261166281846115bf565b90509392505050565b600060e08301600083015161168360008601826112df565b50602083015161169660208601826112df565b50604083015184820360408601526116ae82826112a6565b915050606083015184820360608601526116c882826112a6565b91505060808301516116dd6080860182611329565b5060a083015184820360a08601526116f582826112a6565b91505060c083015161170a60c08601826111ed565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061178157611780611741565b5b50565b600081905061179282611770565b919050565b60006117a282611784565b9050919050565b6117b281611797565b82525050565b600060c0830160008301516117d060008601826117a9565b5060208301516117e360208601826112df565b5060408301516117f66040860182611329565b50606083015161180960608601826111fc565b50608083015161181c60808601826112ee565b5060a083015184820360a086015261183482826112a6565b9150508091505092915050565b600061184d83836117b8565b905092915050565b6000602082019050919050565b600061186d82611715565b6118778185611720565b93508360208202850161188985611731565b8060005b858110156118c557848403895281516118a68582611841565b94506118b183611855565b925060208a0199505060018101905061188d565b50829750879550505050505092915050565b6000610180830160008301516118f0600086018261120b565b50602083015161190360808601826112df565b50604083015184820360a086015261191b82826112a6565b915050606083015184820360c086015261193582826112a6565b915050608083015161194a60e08601826112df565b5060a083015161195e610100860182611329565b5060c0830151611972610120860182611329565b5060e083015184820361014086015261198b828261166b565b9150506101008301518482036101608601526119a78282611862565b9150508091505092915050565b600060208201905081810360008301526119ce81846118d7565b90509291505056fea264697066735822122022062246cb1f45aa5a2395f57c7fce074a3b2e01cc728dbba507fd5c806b2ecd64736f6c63430008130033",
}

// CCIPMessageSentEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use CCIPMessageSentEmitterMetaData.ABI instead.
var CCIPMessageSentEmitterABI = CCIPMessageSentEmitterMetaData.ABI

// CCIPMessageSentEmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CCIPMessageSentEmitterMetaData.Bin instead.
var CCIPMessageSentEmitterBin = CCIPMessageSentEmitterMetaData.Bin

// DeployCCIPMessageSentEmitter deploys a new Ethereum contract, binding an instance of CCIPMessageSentEmitter to it.
func DeployCCIPMessageSentEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CCIPMessageSentEmitter, error) {
	parsed, err := CCIPMessageSentEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCIPMessageSentEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCIPMessageSentEmitter{CCIPMessageSentEmitterCaller: CCIPMessageSentEmitterCaller{contract: contract}, CCIPMessageSentEmitterTransactor: CCIPMessageSentEmitterTransactor{contract: contract}, CCIPMessageSentEmitterFilterer: CCIPMessageSentEmitterFilterer{contract: contract}}, nil
}

// CCIPMessageSentEmitter is an auto generated Go binding around an Ethereum contract.
type CCIPMessageSentEmitter struct {
	CCIPMessageSentEmitterCaller     // Read-only binding to the contract
	CCIPMessageSentEmitterTransactor // Write-only binding to the contract
	CCIPMessageSentEmitterFilterer   // Log filterer for contract events
}

// CCIPMessageSentEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CCIPMessageSentEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCIPMessageSentEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CCIPMessageSentEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCIPMessageSentEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CCIPMessageSentEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCIPMessageSentEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CCIPMessageSentEmitterSession struct {
	Contract     *CCIPMessageSentEmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CCIPMessageSentEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CCIPMessageSentEmitterCallerSession struct {
	Contract *CCIPMessageSentEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CCIPMessageSentEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CCIPMessageSentEmitterTransactorSession struct {
	Contract     *CCIPMessageSentEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CCIPMessageSentEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CCIPMessageSentEmitterRaw struct {
	Contract *CCIPMessageSentEmitter // Generic contract binding to access the raw methods on
}

// CCIPMessageSentEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CCIPMessageSentEmitterCallerRaw struct {
	Contract *CCIPMessageSentEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// CCIPMessageSentEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CCIPMessageSentEmitterTransactorRaw struct {
	Contract *CCIPMessageSentEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCCIPMessageSentEmitter creates a new instance of CCIPMessageSentEmitter, bound to a specific deployed contract.
func NewCCIPMessageSentEmitter(address common.Address, backend bind.ContractBackend) (*CCIPMessageSentEmitter, error) {
	contract, err := bindCCIPMessageSentEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitter{CCIPMessageSentEmitterCaller: CCIPMessageSentEmitterCaller{contract: contract}, CCIPMessageSentEmitterTransactor: CCIPMessageSentEmitterTransactor{contract: contract}, CCIPMessageSentEmitterFilterer: CCIPMessageSentEmitterFilterer{contract: contract}}, nil
}

// NewCCIPMessageSentEmitterCaller creates a new read-only instance of CCIPMessageSentEmitter, bound to a specific deployed contract.
func NewCCIPMessageSentEmitterCaller(address common.Address, caller bind.ContractCaller) (*CCIPMessageSentEmitterCaller, error) {
	contract, err := bindCCIPMessageSentEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitterCaller{contract: contract}, nil
}

// NewCCIPMessageSentEmitterTransactor creates a new write-only instance of CCIPMessageSentEmitter, bound to a specific deployed contract.
func NewCCIPMessageSentEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*CCIPMessageSentEmitterTransactor, error) {
	contract, err := bindCCIPMessageSentEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitterTransactor{contract: contract}, nil
}

// NewCCIPMessageSentEmitterFilterer creates a new log filterer instance of CCIPMessageSentEmitter, bound to a specific deployed contract.
func NewCCIPMessageSentEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*CCIPMessageSentEmitterFilterer, error) {
	contract, err := bindCCIPMessageSentEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitterFilterer{contract: contract}, nil
}

// bindCCIPMessageSentEmitter binds a generic wrapper to an already deployed contract.
func bindCCIPMessageSentEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCIPMessageSentEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPMessageSentEmitter.Contract.CCIPMessageSentEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.CCIPMessageSentEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.CCIPMessageSentEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPMessageSentEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x7f512829.
//
// Solidity: function getNonce(uint64 sourceChainSelector, address account) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCaller) GetNonce(opts *bind.CallOpts, sourceChainSelector uint64, account common.Address) (uint64, error) {
	var out []interface{}
	err := _CCIPMessageSentEmitter.contract.Call(opts, &out, "getNonce", sourceChainSelector, account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x7f512829.
//
// Solidity: function getNonce(uint64 sourceChainSelector, address account) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) GetNonce(sourceChainSelector uint64, account common.Address) (uint64, error) {
	return _CCIPMessageSentEmitter.Contract.GetNonce(&_CCIPMessageSentEmitter.CallOpts, sourceChainSelector, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x7f512829.
//
// Solidity: function getNonce(uint64 sourceChainSelector, address account) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCallerSession) GetNonce(sourceChainSelector uint64, account common.Address) (uint64, error) {
	return _CCIPMessageSentEmitter.Contract.GetNonce(&_CCIPMessageSentEmitter.CallOpts, sourceChainSelector, account)
}

// IsExecuted is a free data retrieval call binding the contract method 0x22bd4080.
//
// Solidity: function isExecuted(uint64 sourceChainSelector, uint64 sequenceNumber) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCaller) IsExecuted(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (bool, error) {
	var out []interface{}
	err := _CCIPMessageSentEmitter.contract.Call(opts, &out, "isExecuted", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecuted is a free data retrieval call binding the contract method 0x22bd4080.
//
// Solidity: function isExecuted(uint64 sourceChainSelector, uint64 sequenceNumber) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) IsExecuted(sourceChainSelector uint64, sequenceNumber uint64) (bool, error) {
	return _CCIPMessageSentEmitter.Contract.IsExecuted(&_CCIPMessageSentEmitter.CallOpts, sourceChainSelector, sequenceNumber)
}

// IsExecuted is a free data retrieval call binding the contract method 0x22bd4080.
//
// Solidity: function isExecuted(uint64 sourceChainSelector, uint64 sequenceNumber) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCallerSession) IsExecuted(sourceChainSelector uint64, sequenceNumber uint64) (bool, error) {
	return _CCIPMessageSentEmitter.Contract.IsExecuted(&_CCIPMessageSentEmitter.CallOpts, sourceChainSelector, sequenceNumber)
}

// SExecuted is a free data retrieval call binding the contract method 0x964ce3a3.
//
// Solidity: function s_executed(uint64 , uint64 ) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCaller) SExecuted(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (bool, error) {
	var out []interface{}
	err := _CCIPMessageSentEmitter.contract.Call(opts, &out, "s_executed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SExecuted is a free data retrieval call binding the contract method 0x964ce3a3.
//
// Solidity: function s_executed(uint64 , uint64 ) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) SExecuted(arg0 uint64, arg1 uint64) (bool, error) {
	return _CCIPMessageSentEmitter.Contract.SExecuted(&_CCIPMessageSentEmitter.CallOpts, arg0, arg1)
}

// SExecuted is a free data retrieval call binding the contract method 0x964ce3a3.
//
// Solidity: function s_executed(uint64 , uint64 ) view returns(bool)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCallerSession) SExecuted(arg0 uint64, arg1 uint64) (bool, error) {
	return _CCIPMessageSentEmitter.Contract.SExecuted(&_CCIPMessageSentEmitter.CallOpts, arg0, arg1)
}

// SNonces is a free data retrieval call binding the contract method 0x74065f33.
//
// Solidity: function s_nonces(uint64 , address ) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCaller) SNonces(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _CCIPMessageSentEmitter.contract.Call(opts, &out, "s_nonces", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SNonces is a free data retrieval call binding the contract method 0x74065f33.
//
// Solidity: function s_nonces(uint64 , address ) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) SNonces(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _CCIPMessageSentEmitter.Contract.SNonces(&_CCIPMessageSentEmitter.CallOpts, arg0, arg1)
}

// SNonces is a free data retrieval call binding the contract method 0x74065f33.
//
// Solidity: function s_nonces(uint64 , address ) view returns(uint64)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterCallerSession) SNonces(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _CCIPMessageSentEmitter.Contract.SNonces(&_CCIPMessageSentEmitter.CallOpts, arg0, arg1)
}

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xc4e0519b.
//
// Solidity: function emitCCIPMessageSent(((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactor) EmitCCIPMessageSent(opts *bind.TransactOpts, message CCIPMessageSentEmitterEVM2AnyVerifierMessage) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.contract.Transact(opts, "emitCCIPMessageSent", message)
}

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xc4e0519b.
//
// Solidity: function emitCCIPMessageSent(((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) EmitCCIPMessageSent(message CCIPMessageSentEmitterEVM2AnyVerifierMessage) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.EmitCCIPMessageSent(&_CCIPMessageSentEmitter.TransactOpts, message)
}

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xc4e0519b.
//
// Solidity: function emitCCIPMessageSent(((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorSession) EmitCCIPMessageSent(message CCIPMessageSentEmitterEVM2AnyVerifierMessage) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.EmitCCIPMessageSent(&_CCIPMessageSentEmitter.TransactOpts, message)
}

// Execute is a paid mutator transaction binding the contract method 0x0d7496ee.
//
// Solidity: function execute(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactor) Execute(opts *bind.TransactOpts, message CCIPMessageSentEmitterAny2EVMMultiProofMessage, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.contract.Transact(opts, "execute", message, proofs)
}

// Execute is a paid mutator transaction binding the contract method 0x0d7496ee.
//
// Solidity: function execute(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) Execute(message CCIPMessageSentEmitterAny2EVMMultiProofMessage, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.Execute(&_CCIPMessageSentEmitter.TransactOpts, message, proofs)
}

// Execute is a paid mutator transaction binding the contract method 0x0d7496ee.
//
// Solidity: function execute(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorSession) Execute(message CCIPMessageSentEmitterAny2EVMMultiProofMessage, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.Execute(&_CCIPMessageSentEmitter.TransactOpts, message, proofs)
}

// SetExecuted is a paid mutator transaction binding the contract method 0x02399390.
//
// Solidity: function setExecuted(uint64 sourceChainSelector, uint64 sequenceNumber, bool executed) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactor) SetExecuted(opts *bind.TransactOpts, sourceChainSelector uint64, sequenceNumber uint64, executed bool) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.contract.Transact(opts, "setExecuted", sourceChainSelector, sequenceNumber, executed)
}

// SetExecuted is a paid mutator transaction binding the contract method 0x02399390.
//
// Solidity: function setExecuted(uint64 sourceChainSelector, uint64 sequenceNumber, bool executed) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) SetExecuted(sourceChainSelector uint64, sequenceNumber uint64, executed bool) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.SetExecuted(&_CCIPMessageSentEmitter.TransactOpts, sourceChainSelector, sequenceNumber, executed)
}

// SetExecuted is a paid mutator transaction binding the contract method 0x02399390.
//
// Solidity: function setExecuted(uint64 sourceChainSelector, uint64 sequenceNumber, bool executed) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorSession) SetExecuted(sourceChainSelector uint64, sequenceNumber uint64, executed bool) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.SetExecuted(&_CCIPMessageSentEmitter.TransactOpts, sourceChainSelector, sequenceNumber, executed)
}

// SetNonce is a paid mutator transaction binding the contract method 0xc8d61748.
//
// Solidity: function setNonce(uint64 sourceChainSelector, address account, uint64 nonce) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactor) SetNonce(opts *bind.TransactOpts, sourceChainSelector uint64, account common.Address, nonce uint64) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.contract.Transact(opts, "setNonce", sourceChainSelector, account, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xc8d61748.
//
// Solidity: function setNonce(uint64 sourceChainSelector, address account, uint64 nonce) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) SetNonce(sourceChainSelector uint64, account common.Address, nonce uint64) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.SetNonce(&_CCIPMessageSentEmitter.TransactOpts, sourceChainSelector, account, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xc8d61748.
//
// Solidity: function setNonce(uint64 sourceChainSelector, address account, uint64 nonce) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorSession) SetNonce(sourceChainSelector uint64, account common.Address, nonce uint64) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.SetNonce(&_CCIPMessageSentEmitter.TransactOpts, sourceChainSelector, account, nonce)
}

// CCIPMessageSentEmitterCCIPMessageSentIterator is returned from FilterCCIPMessageSent and is used to iterate over the raw logs and unpacked data for CCIPMessageSent events raised by the CCIPMessageSentEmitter contract.
type CCIPMessageSentEmitterCCIPMessageSentIterator struct {
	Event *CCIPMessageSentEmitterCCIPMessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CCIPMessageSentEmitterCCIPMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPMessageSentEmitterCCIPMessageSent)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CCIPMessageSentEmitterCCIPMessageSent)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CCIPMessageSentEmitterCCIPMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCIPMessageSentEmitterCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCIPMessageSentEmitterCCIPMessageSent represents a CCIPMessageSent event raised by the CCIPMessageSentEmitter contract.
type CCIPMessageSentEmitterCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           CCIPMessageSentEmitterEVM2AnyVerifierMessage
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCCIPMessageSent is a free log retrieval operation binding the contract event 0xb96f6686b286eb9b9f7114d0d8d35921e425c4fbf84733ac3ea67912d8320ce1.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*CCIPMessageSentEmitterCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _CCIPMessageSentEmitter.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitterCCIPMessageSentIterator{contract: _CCIPMessageSentEmitter.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

// WatchCCIPMessageSent is a free log subscription operation binding the contract event 0xb96f6686b286eb9b9f7114d0d8d35921e425c4fbf84733ac3ea67912d8320ce1.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CCIPMessageSentEmitterCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _CCIPMessageSentEmitter.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCIPMessageSentEmitterCCIPMessageSent)
				if err := _CCIPMessageSentEmitter.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

// ParseCCIPMessageSent is a log parse operation binding the contract event 0xb96f6686b286eb9b9f7114d0d8d35921e425c4fbf84733ac3ea67912d8320ce1.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64),address,bytes,bytes,address,uint256,uint256,(address,address,bytes,bytes,uint256,bytes,bytes32),(uint8,address,uint256,uint64,uint32,bytes)[]) message)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) ParseCCIPMessageSent(log types.Log) (*CCIPMessageSentEmitterCCIPMessageSent, error) {
	event := new(CCIPMessageSentEmitterCCIPMessageSent)
	if err := _CCIPMessageSentEmitter.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCIPMessageSentEmitterExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the CCIPMessageSentEmitter contract.
type CCIPMessageSentEmitterExecutedIterator struct {
	Event *CCIPMessageSentEmitterExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CCIPMessageSentEmitterExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPMessageSentEmitterExecuted)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CCIPMessageSentEmitterExecuted)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CCIPMessageSentEmitterExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCIPMessageSentEmitterExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCIPMessageSentEmitterExecuted represents a Executed event raised by the CCIPMessageSentEmitter contract.
type CCIPMessageSentEmitterExecuted struct {
	Message CCIPMessageSentEmitterAny2EVMMultiProofMessage
	Proofs  [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a.
//
// Solidity: event Executed(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) FilterExecuted(opts *bind.FilterOpts) (*CCIPMessageSentEmitterExecutedIterator, error) {

	logs, sub, err := _CCIPMessageSentEmitter.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &CCIPMessageSentEmitterExecutedIterator{contract: _CCIPMessageSentEmitter.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a.
//
// Solidity: event Executed(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *CCIPMessageSentEmitterExecuted) (event.Subscription, error) {

	logs, sub, err := _CCIPMessageSentEmitter.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCIPMessageSentEmitterExecuted)
				if err := _CCIPMessageSentEmitter.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a.
//
// Solidity: event Executed(((bytes32,uint64,uint64,uint64),bytes,bytes,address,uint32,(bytes,address,bytes,uint256)[],bytes[]) message, bytes[] proofs)
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterFilterer) ParseExecuted(log types.Log) (*CCIPMessageSentEmitterExecuted, error) {
	event := new(CCIPMessageSentEmitterExecuted)
	if err := _CCIPMessageSentEmitter.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
