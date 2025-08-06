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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"requiredVerifierId\",\"type\":\"bytes32\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVMTokenTransfer\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumCCIPMessageSentEmitter.ReceiptType\",\"name\":\"receiptType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPMessageSentEmitter.Receipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structCCIPMessageSentEmitter.EVM2AnyVerifierMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPMessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofTokenTransfer[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"proofs\",\"type\":\"bytes[]\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"requiredVerifierId\",\"type\":\"bytes32\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVMTokenTransfer\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumCCIPMessageSentEmitter.ReceiptType\",\"name\":\"receiptType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPMessageSentEmitter.Receipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCCIPMessageSentEmitter.EVM2AnyVerifierMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"emitCCIPMessageSent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIPMessageSentEmitter.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofTokenTransfer[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\"}],\"internalType\":\"structCCIPMessageSentEmitter.Any2EVMMultiProofMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"proofs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"proofs\",\"type\":\"bytes[]\"}],\"name\":\"executeRaw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"isExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_nonces\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"setExecuted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611fdc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637f512829116100665780637f51282914610130578063964ce3a314610160578063c4e0519b14610190578063c8d61748146101ac578063ee0117c2146101c857610093565b806302399390146100985780630d7496ee146100b457806322bd4080146100d057806374065f3314610100575b600080fd5b6100b260048036038101906100ad91906105a6565b6101e4565b005b6100ce60048036038101906100c99190610c65565b61024d565b005b6100ea60048036038101906100e59190610cdd565b61028a565b6040516100f79190610d2c565b60405180910390f35b61011a60048036038101906101159190610d47565b6102ee565b6040516101279190610d96565b60405180910390f35b61014a60048036038101906101459190610d47565b610324565b6040516101579190610d96565b60405180910390f35b61017a60048036038101906101759190610cdd565b6103a6565b6040516101879190610d2c565b60405180910390f35b6101aa60048036038101906101a591906111ce565b6103d5565b005b6101c660048036038101906101c19190611217565b610435565b005b6101e260048036038101906101dd919061126a565b6104c4565b005b80600160008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b7fb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a828260405161027e929190611729565b60405180910390a15050565b6000600160008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006020528160005260406000206020528060005260406000206000915091509054906101000a900467ffffffffffffffff1681565b60008060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16905092915050565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b80600001516060015167ffffffffffffffff1681600001516040015167ffffffffffffffff167fb96f6686b286eb9b9f7114d0d8d35921e425c4fbf84733ac3ea67912d8320ce18360405161042a9190611aa9565b60405180910390a350565b806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b6000828060200190518101906104da9190611f5d565b90507fb856cfccf99d6742dfb3d807d9e2ebc655e00ead78c4c352651eb645cb48151a818360405161050d929190611729565b60405180910390a1505050565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b61054b8161052e565b811461055657600080fd5b50565b60008135905061056881610542565b92915050565b60008115159050919050565b6105838161056e565b811461058e57600080fd5b50565b6000813590506105a08161057a565b92915050565b6000806000606084860312156105bf576105be610524565b5b60006105cd86828701610559565b93505060206105de86828701610559565b92505060406105ef86828701610591565b9150509250925092565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610647826105fe565b810181811067ffffffffffffffff821117156106665761066561060f565b5b80604052505050565b600061067961051a565b9050610685828261063e565b919050565b600080fd5b6000819050919050565b6106a28161068f565b81146106ad57600080fd5b50565b6000813590506106bf81610699565b92915050565b6000608082840312156106db576106da6105f9565b5b6106e5608061066f565b905060006106f5848285016106b0565b600083015250602061070984828501610559565b602083015250604061071d84828501610559565b604083015250606061073184828501610559565b60608301525092915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156107625761076161060f565b5b61076b826105fe565b9050602081019050919050565b82818337600083830152505050565b600061079a61079584610747565b61066f565b9050828152602081018484840111156107b6576107b5610742565b5b6107c1848285610778565b509392505050565b600082601f8301126107de576107dd61073d565b5b81356107ee848260208601610787565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610822826107f7565b9050919050565b61083281610817565b811461083d57600080fd5b50565b60008135905061084f81610829565b92915050565b600063ffffffff82169050919050565b61086e81610855565b811461087957600080fd5b50565b60008135905061088b81610865565b92915050565b600067ffffffffffffffff8211156108ac576108ab61060f565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b6108d5816108c2565b81146108e057600080fd5b50565b6000813590506108f2816108cc565b92915050565b60006080828403121561090e5761090d6105f9565b5b610918608061066f565b9050600082013567ffffffffffffffff8111156109385761093761068a565b5b610944848285016107c9565b600083015250602061095884828501610840565b602083015250604082013567ffffffffffffffff81111561097c5761097b61068a565b5b610988848285016107c9565b604083015250606061099c848285016108e3565b60608301525092915050565b60006109bb6109b684610891565b61066f565b905080838252602082019050602084028301858111156109de576109dd6108bd565b5b835b81811015610a2557803567ffffffffffffffff811115610a0357610a0261073d565b5b808601610a1089826108f8565b855260208501945050506020810190506109e0565b5050509392505050565b600082601f830112610a4457610a4361073d565b5b8135610a548482602086016109a8565b91505092915050565b600067ffffffffffffffff821115610a7857610a7761060f565b5b602082029050602081019050919050565b6000610a9c610a9784610a5d565b61066f565b90508083825260208201905060208402830185811115610abf57610abe6108bd565b5b835b81811015610b0657803567ffffffffffffffff811115610ae457610ae361073d565b5b808601610af189826107c9565b85526020850194505050602081019050610ac1565b5050509392505050565b600082601f830112610b2557610b2461073d565b5b8135610b35848260208601610a89565b91505092915050565b60006101408284031215610b5557610b546105f9565b5b610b5f60e061066f565b90506000610b6f848285016106c5565b600083015250608082013567ffffffffffffffff811115610b9357610b9261068a565b5b610b9f848285016107c9565b60208301525060a082013567ffffffffffffffff811115610bc357610bc261068a565b5b610bcf848285016107c9565b60408301525060c0610be384828501610840565b60608301525060e0610bf78482850161087c565b60808301525061010082013567ffffffffffffffff811115610c1c57610c1b61068a565b5b610c2884828501610a2f565b60a08301525061012082013567ffffffffffffffff811115610c4d57610c4c61068a565b5b610c5984828501610b10565b60c08301525092915050565b60008060408385031215610c7c57610c7b610524565b5b600083013567ffffffffffffffff811115610c9a57610c99610529565b5b610ca685828601610b3e565b925050602083013567ffffffffffffffff811115610cc757610cc6610529565b5b610cd385828601610b10565b9150509250929050565b60008060408385031215610cf457610cf3610524565b5b6000610d0285828601610559565b9250506020610d1385828601610559565b9150509250929050565b610d268161056e565b82525050565b6000602082019050610d416000830184610d1d565b92915050565b60008060408385031215610d5e57610d5d610524565b5b6000610d6c85828601610559565b9250506020610d7d85828601610840565b9150509250929050565b610d908161052e565b82525050565b6000602082019050610dab6000830184610d87565b92915050565b600060e08284031215610dc757610dc66105f9565b5b610dd160e061066f565b90506000610de184828501610840565b6000830152506020610df584828501610840565b602083015250604082013567ffffffffffffffff811115610e1957610e1861068a565b5b610e25848285016107c9565b604083015250606082013567ffffffffffffffff811115610e4957610e4861068a565b5b610e55848285016107c9565b6060830152506080610e69848285016108e3565b60808301525060a082013567ffffffffffffffff811115610e8d57610e8c61068a565b5b610e99848285016107c9565b60a08301525060c0610ead848285016106b0565b60c08301525092915050565b600067ffffffffffffffff821115610ed457610ed361060f565b5b602082029050602081019050919050565b60028110610ef257600080fd5b50565b600081359050610f0481610ee5565b92915050565b600060c08284031215610f2057610f1f6105f9565b5b610f2a60c061066f565b90506000610f3a84828501610ef5565b6000830152506020610f4e84828501610840565b6020830152506040610f62848285016108e3565b6040830152506060610f7684828501610559565b6060830152506080610f8a8482850161087c565b60808301525060a082013567ffffffffffffffff811115610fae57610fad61068a565b5b610fba848285016107c9565b60a08301525092915050565b6000610fd9610fd484610eb9565b61066f565b90508083825260208201905060208402830185811115610ffc57610ffb6108bd565b5b835b8181101561104357803567ffffffffffffffff8111156110215761102061073d565b5b80860161102e8982610f0a565b85526020850194505050602081019050610ffe565b5050509392505050565b600082601f8301126110625761106161073d565b5b8135611072848260208601610fc6565b91505092915050565b60006101808284031215611092576110916105f9565b5b61109d61012061066f565b905060006110ad848285016106c5565b60008301525060806110c184828501610840565b60208301525060a082013567ffffffffffffffff8111156110e5576110e461068a565b5b6110f1848285016107c9565b60408301525060c082013567ffffffffffffffff8111156111155761111461068a565b5b611121848285016107c9565b60608301525060e061113584828501610840565b60808301525061010061114a848285016108e3565b60a08301525061012061115f848285016108e3565b60c08301525061014082013567ffffffffffffffff8111156111845761118361068a565b5b61119084828501610db1565b60e08301525061016082013567ffffffffffffffff8111156111b5576111b461068a565b5b6111c18482850161104d565b6101008301525092915050565b6000602082840312156111e4576111e3610524565b5b600082013567ffffffffffffffff81111561120257611201610529565b5b61120e8482850161107b565b91505092915050565b6000806000606084860312156112305761122f610524565b5b600061123e86828701610559565b935050602061124f86828701610840565b925050604061126086828701610559565b9150509250925092565b6000806040838503121561128157611280610524565b5b600083013567ffffffffffffffff81111561129f5761129e610529565b5b6112ab858286016107c9565b925050602083013567ffffffffffffffff8111156112cc576112cb610529565b5b6112d885828601610b10565b9150509250929050565b6112eb8161068f565b82525050565b6112fa8161052e565b82525050565b60808201600082015161131660008501826112e2565b50602082015161132960208501826112f1565b50604082015161133c60408501826112f1565b50606082015161134f60608501826112f1565b50505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561138f578082015181840152602081019050611374565b60008484015250505050565b60006113a682611355565b6113b08185611360565b93506113c0818560208601611371565b6113c9816105fe565b840191505092915050565b6113dd81610817565b82525050565b6113ec81610855565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611427816108c2565b82525050565b6000608083016000830151848203600086015261144a828261139b565b915050602083015161145f60208601826113d4565b5060408301518482036040860152611477828261139b565b915050606083015161148c606086018261141e565b508091505092915050565b60006114a3838361142d565b905092915050565b6000602082019050919050565b60006114c3826113f2565b6114cd81856113fd565b9350836020820285016114df8561140e565b8060005b8581101561151b57848403895281516114fc8582611497565b9450611507836114ab565b925060208a019950506001810190506114e3565b50829750879550505050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000611565838361139b565b905092915050565b6000602082019050919050565b60006115858261152d565b61158f8185611538565b9350836020820285016115a185611549565b8060005b858110156115dd57848403895281516115be8582611559565b94506115c98361156d565b925060208a019950506001810190506115a5565b50829750879550505050505092915050565b6000610140830160008301516116086000860182611300565b5060208301518482036080860152611620828261139b565b915050604083015184820360a086015261163a828261139b565b915050606083015161164f60c08601826113d4565b50608083015161166260e08601826113e3565b5060a083015184820361010086015261167b82826114b8565b91505060c0830151848203610120860152611696828261157a565b9150508091505092915050565b600082825260208201905092915050565b60006116bf8261152d565b6116c981856116a3565b9350836020820285016116db85611549565b8060005b8581101561171757848403895281516116f88582611559565b94506117038361156d565b925060208a019950506001810190506116df565b50829750879550505050505092915050565b6000604082019050818103600083015261174381856115ef565b9050818103602083015261175781846116b4565b90509392505050565b600060e08301600083015161177860008601826113d4565b50602083015161178b60208601826113d4565b50604083015184820360408601526117a3828261139b565b915050606083015184820360608601526117bd828261139b565b91505060808301516117d2608086018261141e565b5060a083015184820360a08601526117ea828261139b565b91505060c08301516117ff60c08601826112e2565b508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061187657611875611836565b5b50565b600081905061188782611865565b919050565b600061189782611879565b9050919050565b6118a78161188c565b82525050565b600060c0830160008301516118c5600086018261189e565b5060208301516118d860208601826113d4565b5060408301516118eb604086018261141e565b5060608301516118fe60608601826112f1565b50608083015161191160808601826113e3565b5060a083015184820360a0860152611929828261139b565b9150508091505092915050565b600061194283836118ad565b905092915050565b6000602082019050919050565b60006119628261180a565b61196c8185611815565b93508360208202850161197e85611826565b8060005b858110156119ba578484038952815161199b8582611936565b94506119a68361194a565b925060208a01995050600181019050611982565b50829750879550505050505092915050565b6000610180830160008301516119e56000860182611300565b5060208301516119f860808601826113d4565b50604083015184820360a0860152611a10828261139b565b915050606083015184820360c0860152611a2a828261139b565b9150506080830151611a3f60e08601826113d4565b5060a0830151611a5361010086018261141e565b5060c0830151611a6761012086018261141e565b5060e0830151848203610140860152611a808282611760565b915050610100830151848203610160860152611a9c8282611957565b9150508091505092915050565b60006020820190508181036000830152611ac381846119cc565b905092915050565b600081519050611ada81610699565b92915050565b600081519050611aef81610542565b92915050565b600060808284031215611b0b57611b0a6105f9565b5b611b15608061066f565b90506000611b2584828501611acb565b6000830152506020611b3984828501611ae0565b6020830152506040611b4d84828501611ae0565b6040830152506060611b6184828501611ae0565b60608301525092915050565b6000611b80611b7b84610747565b61066f565b905082815260208101848484011115611b9c57611b9b610742565b5b611ba7848285611371565b509392505050565b600082601f830112611bc457611bc361073d565b5b8151611bd4848260208601611b6d565b91505092915050565b600081519050611bec81610829565b92915050565b600081519050611c0181610865565b92915050565b600081519050611c16816108cc565b92915050565b600060808284031215611c3257611c316105f9565b5b611c3c608061066f565b9050600082015167ffffffffffffffff811115611c5c57611c5b61068a565b5b611c6884828501611baf565b6000830152506020611c7c84828501611bdd565b602083015250604082015167ffffffffffffffff811115611ca057611c9f61068a565b5b611cac84828501611baf565b6040830152506060611cc084828501611c07565b60608301525092915050565b6000611cdf611cda84610891565b61066f565b90508083825260208201905060208402830185811115611d0257611d016108bd565b5b835b81811015611d4957805167ffffffffffffffff811115611d2757611d2661073d565b5b808601611d348982611c1c565b85526020850194505050602081019050611d04565b5050509392505050565b600082601f830112611d6857611d6761073d565b5b8151611d78848260208601611ccc565b91505092915050565b6000611d94611d8f84610a5d565b61066f565b90508083825260208201905060208402830185811115611db757611db66108bd565b5b835b81811015611dfe57805167ffffffffffffffff811115611ddc57611ddb61073d565b5b808601611de98982611baf565b85526020850194505050602081019050611db9565b5050509392505050565b600082601f830112611e1d57611e1c61073d565b5b8151611e2d848260208601611d81565b91505092915050565b60006101408284031215611e4d57611e4c6105f9565b5b611e5760e061066f565b90506000611e6784828501611af5565b600083015250608082015167ffffffffffffffff811115611e8b57611e8a61068a565b5b611e9784828501611baf565b60208301525060a082015167ffffffffffffffff811115611ebb57611eba61068a565b5b611ec784828501611baf565b60408301525060c0611edb84828501611bdd565b60608301525060e0611eef84828501611bf2565b60808301525061010082015167ffffffffffffffff811115611f1457611f1361068a565b5b611f2084828501611d53565b60a08301525061012082015167ffffffffffffffff811115611f4557611f4461068a565b5b611f5184828501611e08565b60c08301525092915050565b600060208284031215611f7357611f72610524565b5b600082015167ffffffffffffffff811115611f9157611f90610529565b5b611f9d84828501611e36565b9150509291505056fea2646970667358221220f26f42c211ec720677eac847dfbd66691bf1d9fe5a7299d28e9400d799dbb8d464736f6c63430008130033",
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

// ExecuteRaw is a paid mutator transaction binding the contract method 0xee0117c2.
//
// Solidity: function executeRaw(bytes encodedMessage, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactor) ExecuteRaw(opts *bind.TransactOpts, encodedMessage []byte, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.contract.Transact(opts, "executeRaw", encodedMessage, proofs)
}

// ExecuteRaw is a paid mutator transaction binding the contract method 0xee0117c2.
//
// Solidity: function executeRaw(bytes encodedMessage, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterSession) ExecuteRaw(encodedMessage []byte, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.ExecuteRaw(&_CCIPMessageSentEmitter.TransactOpts, encodedMessage, proofs)
}

// ExecuteRaw is a paid mutator transaction binding the contract method 0xee0117c2.
//
// Solidity: function executeRaw(bytes encodedMessage, bytes[] proofs) returns()
func (_CCIPMessageSentEmitter *CCIPMessageSentEmitterTransactorSession) ExecuteRaw(encodedMessage []byte, proofs [][]byte) (*types.Transaction, error) {
	return _CCIPMessageSentEmitter.Contract.ExecuteRaw(&_CCIPMessageSentEmitter.TransactOpts, encodedMessage, proofs)
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
