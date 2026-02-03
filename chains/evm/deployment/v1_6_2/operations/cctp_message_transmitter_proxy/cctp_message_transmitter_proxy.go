package cctp_message_transmitter_proxy

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var ContractType cldf_deployment.ContractType = "CCTPMessageTransmitterProxy"
var Version = semver.MustParse("1.6.2")
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

const CCTPMessageTransmitterProxyABI = `[{"type":"constructor","inputs":[{"name":"tokenMessenger","type":"address","internalType":"contract ITokenMessenger"}],"stateMutability":"nonpayable"},{"type":"function","name":"acceptOwnership","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"configureAllowedCallers","inputs":[{"name":"configArgs","type":"tuple[]","internalType":"struct CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[]","components":[{"name":"caller","type":"address","internalType":"address"},{"name":"allowed","type":"bool","internalType":"bool"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"getAllowedCallers","inputs":[],"outputs":[{"name":"allowedCallers","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"i_cctpTransmitter","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IMessageTransmitter"}],"stateMutability":"view"},{"type":"function","name":"isAllowedCaller","inputs":[{"name":"caller","type":"address","internalType":"address"}],"outputs":[{"name":"allowed","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"receiveMessage","inputs":[{"name":"message","type":"bytes","internalType":"bytes"},{"name":"attestation","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"success","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"transferOwnership","inputs":[{"name":"to","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"typeAndVersion","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"event","name":"AllowedCallerAdded","inputs":[{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"AllowedCallerRemoved","inputs":[{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"OwnershipTransferRequested","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"error","name":"CannotTransferToSelf","inputs":[]},{"type":"error","name":"MustBeProposedOwner","inputs":[]},{"type":"error","name":"OnlyCallableByOwner","inputs":[]},{"type":"error","name":"OwnerCannotBeZero","inputs":[]},{"type":"error","name":"Unauthorized","inputs":[{"name":"caller","type":"address","internalType":"address"}]}]`
const CCTPMessageTransmitterProxyBin = "0x60a0806040523461010f57602081610e34803803809161001f8285610114565b83398101031261010f57516001600160a01b0381169081900361010f5733156100fe57600180546001600160a01b03191633179055604051632c12192160e01b815290602090829060049082905afa9081156100f2576000916100a9575b506001600160a01b0316608052604051610ce6908161014e82396080518181816101c501526106680152f35b6020813d6020116100ea575b816100c260209383610114565b810103126100e65751906001600160a01b03821682036100e357503861007d565b80fd5b5080fd5b3d91506100b5565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761013757604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816310807aa71461085857508063181f5a771461072057806357ecfd281461054957806379ba5097146104605780638da5cb5b1461040e578063a68012581461039b578063bd028e7c146101e9578063cfc1db061461017a5763f2fde38b1461008257600080fd5b346101755760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101755760043573ffffffffffffffffffffffffffffffffffffffff8116809103610175576100da610a51565b33811461014b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101755760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101755760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101755760043567ffffffffffffffff8111610175573660238201121561017557806004013567ffffffffffffffff8111610175576024820191602436918360061b01011161017557610265610a51565b60005b81811061027157005b602061027e8284866109f1565b0135908115158203610175576001911561031c576102c373ffffffffffffffffffffffffffffffffffffffff6102bd6102b88487896109f1565b610a30565b16610c79565b6102ce575b01610268565b73ffffffffffffffffffffffffffffffffffffffff6102f16102b88386886109f1565b167f663c7e9ed36d9138863ef4306bbfcf01f60e1e7ca69b370c53d3094369e2cb02600080a26102c8565b61034873ffffffffffffffffffffffffffffffffffffffff6103426102b88487896109f1565b16610ab4565b156102c85773ffffffffffffffffffffffffffffffffffffffff6103706102b88386886109f1565b167fbc0a6e072a312bde289d32bc84e5b758d7c617f734ecc0d69f995b2d7e69be36600080a26102c8565b346101755760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101755760043573ffffffffffffffffffffffffffffffffffffffff8116809103610175576104046020916000526003602052604060002054151590565b6040519015158152f35b346101755760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101755760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101755760005473ffffffffffffffffffffffffffffffffffffffff8116330361051f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101755760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101755760043567ffffffffffffffff811161017557610598903690600401610984565b60243567ffffffffffffffff8111610175576105b8903690600401610984565b9290916105d2336000526003602052604060002054151590565b156106f25761064d60209361061d9560405196879586957f57ecfd280000000000000000000000000000000000000000000000000000000087526040600488015260448701916109b2565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8584030160248601526109b2565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156106e6576000916106a6575b6020826040519015158152f35b6020813d6020116106de575b816106bf60209383610943565b810103126106da575180151581036106da5790506020610699565b5080fd5b3d91506106b2565b6040513d6000823e3d90fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101755760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610175576040516060810181811067ffffffffffffffff82111761082957604052602181527f434354504d6573736167655472616e736d697474657250726f787920312e362e60208201527f3200000000000000000000000000000000000000000000000000000000000000604082015260405190602082528181519182602083015260005b8381106108115750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b602082820181015160408784010152859350016107d1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101755760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610175576002549081815260208101809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061092d57505050816108d4910382610943565b6040519182916020830190602084525180915260408301919060005b8181106108fe575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016108f0565b82548452602090930192600192830192016108be565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761082957604052565b9181601f840112156101755782359167ffffffffffffffff8311610175576020838186019501011161017557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9190811015610a015760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036101755790565b73ffffffffffffffffffffffffffffffffffffffff600154163303610a7257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015610a015760005260206000200190600090565b6000818152600360205260409020548015610c72577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610c4357600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610c4357818103610bd4575b5050506002548015610ba5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610b62816002610a9c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610c2b610be5610bf6936002610a9c565b90549060031b1c9283926002610a9c565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610b29565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610cd3576002546801000000000000000081101561082957610cba610bf68260018594016002556002610a9c565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a"

type CCTPMessageTransmitterProxyContract struct {
	address  common.Address
	abi      abi.ABI
	backend  bind.ContractBackend
	contract *bind.BoundContract
}

func NewCCTPMessageTransmitterProxyContract(
	address common.Address,
	backend bind.ContractBackend,
) (*CCTPMessageTransmitterProxyContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CCTPMessageTransmitterProxyABI))
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyContract{
		address:  address,
		abi:      parsed,
		backend:  backend,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (c *CCTPMessageTransmitterProxyContract) Address() common.Address {
	return c.address
}

func (c *CCTPMessageTransmitterProxyContract) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := c.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

func (c *CCTPMessageTransmitterProxyContract) ConfigureAllowedCallers(opts *bind.TransactOpts, args []AllowedCallerConfigArgs) (*types.Transaction, error) {
	return c.contract.Transact(opts, "configureAllowedCallers", args)
}

type AllowedCallerConfigArgs struct {
	Caller  common.Address
	Allowed bool
}

type ConstructorArgs struct {
	TokenMessenger common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:        "cctp-message-transmitter-proxy:deploy",
	Version:     Version,
	Description: "Deploys the CCTPMessageTransmitterProxy contract",
	ContractMetadata: &bind.MetaData{
		ABI: CCTPMessageTransmitterProxyABI,
		Bin: CCTPMessageTransmitterProxyBin,
	},
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(CCTPMessageTransmitterProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ConfigureAllowedCallers = contract.NewWrite(contract.WriteParams[[]AllowedCallerConfigArgs, *CCTPMessageTransmitterProxyContract]{
	Name:            "cctp-message-transmitter-proxy:configure-allowed-callers",
	Version:         Version,
	Description:     "Calls configureAllowedCallers on the contract",
	ContractType:    ContractType,
	ContractABI:     CCTPMessageTransmitterProxyABI,
	NewContract:     NewCCTPMessageTransmitterProxyContract,
	IsAllowedCaller: contract.OnlyOwner[*CCTPMessageTransmitterProxyContract, []AllowedCallerConfigArgs],
	Validate:        func([]AllowedCallerConfigArgs) error { return nil },
	CallContract: func(
		c *CCTPMessageTransmitterProxyContract,
		opts *bind.TransactOpts,
		args []AllowedCallerConfigArgs,
	) (*types.Transaction, error) {
		return c.ConfigureAllowedCallers(opts, args)
	},
})
