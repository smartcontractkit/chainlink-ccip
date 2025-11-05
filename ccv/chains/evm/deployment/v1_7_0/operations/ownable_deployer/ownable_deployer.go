package ownable_deployer

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/ownable_deployer"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OwnableDeployer"

type ConstructorArgs struct{}

type ComputeAddressArgs struct {
	Sender   common.Address
	InitCode []byte
	Salt     [32]byte
}

type DeployAndTransferOwnershipArgs struct {
	InitCode []byte
	Salt     [32]byte
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "ownable-deployer:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the OwnableDeployer contract",
	ContractMetadata: ownable_deployer.OwnableDeployerMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(ownable_deployer.OwnableDeployerBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var DeployAndTransferOwnership = contract.NewWrite(contract.WriteParams[DeployAndTransferOwnershipArgs, *ownable_deployer.OwnableDeployer]{
	Name:            "ownable-deployer:deploy-and-transfer-ownership",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Deploys and transfers ownership of a contract with the given init code and salt",
	ContractType:    ContractType,
	ContractABI:     ownable_deployer.OwnableDeployerABI,
	NewContract:     ownable_deployer.NewOwnableDeployer,
	IsAllowedCaller: contract.AllCallersAllowed[*ownable_deployer.OwnableDeployer, DeployAndTransferOwnershipArgs],
	Validate:        func(DeployAndTransferOwnershipArgs) error { return nil },
	CallContract: func(contract *ownable_deployer.OwnableDeployer, opts *bind.TransactOpts, input DeployAndTransferOwnershipArgs) (*types.Transaction, error) {
		return contract.DeployAndTransferOwnership(opts, input.InitCode, input.Salt)
	},
})

var ComputeAddress = contract.NewRead(contract.ReadParams[ComputeAddressArgs, common.Address, *ownable_deployer.OwnableDeployer]{
	Name:         "ownable-deployer:compute-address",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Computes the address of a contract that will be deployed with the given init code and salt by the given sender",
	ContractType: ContractType,
	NewContract:  ownable_deployer.NewOwnableDeployer,
	CallContract: func(contract *ownable_deployer.OwnableDeployer, opts *bind.CallOpts, input ComputeAddressArgs) (common.Address, error) {
		return contract.ComputeAddress(opts, input.Sender, input.InitCode, input.Salt)
	},
})
