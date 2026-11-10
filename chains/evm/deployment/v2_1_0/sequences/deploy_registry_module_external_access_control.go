package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/registry_module_external_access_control"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// DeployRegistryModuleExternalAccessControlInput is the input for the DeployRegistryModuleExternalAccessControl sequence.
type DeployRegistryModuleExternalAccessControlInput struct {
	ChainSelector uint64
	// TokenAdminRegistry is the address of the TokenAdminRegistry the module will register admins in.
	TokenAdminRegistry common.Address
	// ExistingAddresses are the addresses already present in the datastore for the target chain.
	ExistingAddresses []datastore.AddressRef
}

// DeployRegistryModuleExternalAccessControl deploys the RegistryModuleExternalAccessControl contract and registers it as
// a registry module on the chain's TokenAdminRegistry. Both steps are idempotent: the module is only deployed if it is
// not already present in the datastore, and it is only added to the TokenAdminRegistry if it is not already registered.
var DeployRegistryModuleExternalAccessControl = cldf_ops.NewSequence(
	"deploy-registry-module-external-access-control",
	semver.MustParse("2.1.0"),
	"Deploys the RegistryModuleExternalAccessControl contract and registers it on the TokenAdminRegistry",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployRegistryModuleExternalAccessControlInput) (sequences.OnChainOutput, error) {
		if input.TokenAdminRegistry == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("token admin registry address is required")
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		// Deploy the module, or reuse the existing deployment if one is already recorded.
		moduleRef, err := contract_utils.MaybeDeployContract(
			b,
			registry_module_external_access_control.Deploy,
			chain,
			contract_utils.DeployInput[registry_module_external_access_control.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(
					registry_module_external_access_control.ContractType,
					*registry_module_external_access_control.Version,
				),
				ChainSelector: chain.Selector,
				Args: registry_module_external_access_control.ConstructorArgs{
					TokenAdminRegistry: input.TokenAdminRegistry,
				},
			},
			input.ExistingAddresses,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RegistryModuleExternalAccessControl: %w", err)
		}
		addresses = append(addresses, moduleRef)

		moduleAddress := common.HexToAddress(moduleRef.Address)

		// Register the module on the TokenAdminRegistry if it is not already registered. The TAR is owned by the
		// timelock on live chains, so this write is prepared as an MCMS transaction when the deployer key is not the
		// owner.
		//
		// WithForceExecute ensures the idempotency read is never served from a cached report: a re-run after an
		// out-of-band change (or after a previous run in the same bundle) has to see the current on-chain state,
		// otherwise a redundant write would be emitted.
		isRegisteredReport, err := cldf_ops.ExecuteOperation(
			b,
			tarops.IsRegistryModule,
			chain,
			contract_utils.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       input.TokenAdminRegistry,
				Args:          moduleAddress,
			},
			cldf_ops.WithForceExecute[contract_utils.FunctionInput[common.Address], evm.Chain](),
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if RegistryModuleExternalAccessControl is registered on TokenAdminRegistry: %w", err)
		}

		// Only emit the write when there is an on-chain diff. addRegistryModule is itself idempotent on-chain, but
		// skipping it keeps re-runs from emitting no-op transactions in MCMS proposals.
		if !isRegisteredReport.Output {
			addRegistryModuleReport, err := cldf_ops.ExecuteOperation(b, tarops.AddRegistryModule, chain, contract_utils.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       input.TokenAdminRegistry,
				Args:          moduleAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to add RegistryModuleExternalAccessControl to TokenAdminRegistry: %w", err)
			}
			writes = append(writes, addRegistryModuleReport.Output)
		}

		// Build the batch op only from unexecuted writes. When the deployer key owns the TAR the write executes
		// directly and there is nothing to propose, so the output carries no BatchOps rather than an empty one.
		batchOps := make([]mcms_types.BatchOperation, 0, 1)
		if len(writes) > 0 {
			batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			if len(batchOp.Transactions) > 0 {
				batchOps = append(batchOps, batchOp)
			}
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	})
