package create2_factory

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CREATE2Factory"

type ConstructorArgs struct {
	AllowList []common.Address
}

type ComputeAddressArgs struct {
	ABI             string
	Bin             string
	ConstructorArgs []any
	Salt            string
}

type CreateAndTransferOwnershipArgs struct {
	ComputeAddressArgs
	To common.Address
}

type ApplyAllowListUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "create2-factory:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CREATE2Factory contract",
	ContractMetadata: create2_factory.CREATE2FactoryMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(create2_factory.CREATE2FactoryBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var CreateAndTransferOwnership = contract.NewWrite(contract.WriteParams[CreateAndTransferOwnershipArgs, *create2_factory.CREATE2Factory]{
	Name:         "create2-factory:deploy-and-transfer-ownership",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Deploys a contract with the given creation code + salt and transfers ownership to the given address",
	ContractType: ContractType,
	ContractABI:  create2_factory.CREATE2FactoryABI,
	NewContract:  create2_factory.NewCREATE2Factory,
	IsAllowedCaller: func(contract *create2_factory.CREATE2Factory, opts *bind.CallOpts, caller common.Address, input CreateAndTransferOwnershipArgs) (bool, error) {
		allowList, err := contract.GetAllowList(opts)
		if err != nil {
			return false, err
		}
		return slices.Contains(allowList, caller), nil
	},
	Validate: func(CreateAndTransferOwnershipArgs) error { return nil },
	CallContract: func(contract *create2_factory.CREATE2Factory, opts *bind.TransactOpts, input CreateAndTransferOwnershipArgs) (*types.Transaction, error) {
		creationCode, err := makeCreationCode(input.ABI, input.Bin, input.ConstructorArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to make creation code: %w", err)
		}
		return contract.CreateAndTransferOwnership(opts, creationCode, hashSalt(input.Salt), input.To)
	},
})

var ComputeAddress = contract.NewRead(contract.ReadParams[ComputeAddressArgs, common.Address, *create2_factory.CREATE2Factory]{
	Name:         "create2-factory:compute-address",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Computes the address of a contract that will be deployed with the given creation code and salt",
	ContractType: ContractType,
	NewContract:  create2_factory.NewCREATE2Factory,
	CallContract: func(contract *create2_factory.CREATE2Factory, opts *bind.CallOpts, input ComputeAddressArgs) (common.Address, error) {
		creationCode, err := makeCreationCode(input.ABI, input.Bin, input.ConstructorArgs...)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to make creation code: %w", err)
		}
		return contract.ComputeAddress(opts, creationCode, hashSalt(input.Salt))
	},
})

var ApplyAllowListUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowListUpdatesArgs, *create2_factory.CREATE2Factory]{
	Name:            "create2-factory:apply-allow-list-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies the allow list updates to the CREATE2Factory contract",
	ContractType:    ContractType,
	ContractABI:     create2_factory.CREATE2FactoryABI,
	NewContract:     create2_factory.NewCREATE2Factory,
	IsAllowedCaller: contract.OnlyOwner[*create2_factory.CREATE2Factory, ApplyAllowListUpdatesArgs],
	Validate:        func(ApplyAllowListUpdatesArgs) error { return nil },
	CallContract: func(contract *create2_factory.CREATE2Factory, opts *bind.TransactOpts, input ApplyAllowListUpdatesArgs) (*types.Transaction, error) {
		return contract.ApplyAllowListUpdates(opts, input.Removes, input.Adds)
	},
})

func makeCreationCode(abi string, bin string, constructorArgs ...any) ([]byte, error) {
	parsedABI, err := ethabi.JSON(strings.NewReader(abi))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}
	packedConstructorArgs, err := parsedABI.Pack("", constructorArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack constructor arguments: %w", err)
	}
	return append(common.FromHex(bin), packedConstructorArgs...), nil
}

func hashSalt(salt string) [32]byte {
	hasher := hashutil.NewKeccak()
	return hasher.Hash(common.LeftPadBytes([]byte(salt), 32))
}
