package contract_factory

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ContractFactory"

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
	Name:             "contract-factory:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the ContractFactory contract",
	ContractMetadata: contract_factory.ContractFactoryMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(contract_factory.ContractFactoryBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var CreateAndTransferOwnership = contract.NewWrite(contract.WriteParams[CreateAndTransferOwnershipArgs, *contract_factory.ContractFactory]{
	Name:         "contract-factory:deploy-and-transfer-ownership",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Deploys a contract with the given creation code + salt and transfers ownership to the given address",
	ContractType: ContractType,
	ContractABI:  contract_factory.ContractFactoryABI,
	NewContract:  contract_factory.NewContractFactory,
	IsAllowedCaller: func(contract *contract_factory.ContractFactory, opts *bind.CallOpts, caller common.Address, input CreateAndTransferOwnershipArgs) (bool, error) {
		allowList, err := contract.GetAllowList(opts)
		if err != nil {
			return false, err
		}
		return slices.Contains(allowList, caller), nil
	},
	Validate: func(CreateAndTransferOwnershipArgs) error { return nil },
	CallContract: func(contract *contract_factory.ContractFactory, opts *bind.TransactOpts, input CreateAndTransferOwnershipArgs) (*types.Transaction, error) {
		creationCode, err := makeCreationCode(input.ABI, input.Bin, input.ConstructorArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to make creation code: %w", err)
		}
		calls := make([][]byte, 0)
		// We use the ContractFactory ABI to construct the transferOwnership call because we conveniently have access to it here.
		// Any other contract that implements IOwnable will be compatible with this operation.
		parsedABI, err := contract_factory.ContractFactoryMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to parse ABI: %w", err)
		}
		transferOwnershipCall, err := parsedABI.Pack("transferOwnership", input.To)
		if err != nil {
			return nil, fmt.Errorf("failed to pack transferOwnership call: %w", err)
		}
		calls = append(calls, transferOwnershipCall)

		return contract.CreateAndCall(opts, creationCode, hashSalt(input.Salt), calls)
	},
})

var ComputeAddress = contract.NewRead(contract.ReadParams[ComputeAddressArgs, common.Address, *contract_factory.ContractFactory]{
	Name:         "contract-factory:compute-address",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Computes the address of a contract that will be deployed with the given creation code and salt",
	ContractType: ContractType,
	NewContract:  contract_factory.NewContractFactory,
	CallContract: func(contract *contract_factory.ContractFactory, opts *bind.CallOpts, input ComputeAddressArgs) (common.Address, error) {
		creationCode, err := makeCreationCode(input.ABI, input.Bin, input.ConstructorArgs...)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to make creation code: %w", err)
		}
		return contract.ComputeAddress(opts, creationCode, hashSalt(input.Salt))
	},
})

var ApplyAllowListUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowListUpdatesArgs, *contract_factory.ContractFactory]{
	Name:            "contract-factory:apply-allow-list-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies the allow list updates to the ContractFactory contract",
	ContractType:    ContractType,
	ContractABI:     contract_factory.ContractFactoryABI,
	NewContract:     contract_factory.NewContractFactory,
	IsAllowedCaller: contract.OnlyOwner[*contract_factory.ContractFactory, ApplyAllowListUpdatesArgs],
	Validate:        func(ApplyAllowListUpdatesArgs) error { return nil },
	CallContract: func(contract *contract_factory.ContractFactory, opts *bind.TransactOpts, input ApplyAllowListUpdatesArgs) (*types.Transaction, error) {
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
