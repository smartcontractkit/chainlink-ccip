package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// ActivateRMNCfg configures the ActivateRMN changeset.
type ActivateRMNCfg struct {
	ChainSels []uint64
}

var ActivateRMN = func(mcmsRegistry *changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[ActivateRMNCfg]] {
	return cldf_deployment.CreateChangeSet(
		func(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) (cldf_deployment.ChangesetOutput, error) {
			return applyDeployAndActivateRMN(e, mcmsRegistry, input)
		},
		validateActivateRMN,
	)
}

func validateActivateRMN(e cldf_deployment.Environment, input changesets.WithMCMS[ActivateRMNCfg]) error {
	if len(input.Cfg.ChainSels) == 0 {
		return fmt.Errorf("at least one chain selector is required")
	}
	evmChains := e.BlockChains.EVMChains()
	for _, sel := range input.Cfg.ChainSels {
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))
		if err := validateActivateRMNAddresses(addresses, sel); err != nil {
			return err
		}
	}
	return nil
}

func applyDeployAndActivateRMN(
	e cldf_deployment.Environment,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	input changesets.WithMCMS[ActivateRMNCfg],
) (cldf_deployment.ChangesetOutput, error) {
	evmChains := e.BlockChains.EVMChains()
	outputDS := datastore.NewMemoryDataStore()
	builder := changesets.NewOutputBuilder(e, mcmsRegistry).WithDataStore(outputDS)

	for _, sel := range input.Cfg.ChainSels {
		chain := evmChains[sel]
		addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployAndActivateRMN, chain, sequences.ActivateRMNInput{
			ChainSelector:     sel,
			ExistingAddresses: addresses,
		})
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to activate RMN on chain %d: %w", sel, err)
		}

		for _, ref := range report.Output.Addresses {
			if addErr := outputDS.Addresses().Add(ref); addErr != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", addErr)
			}
		}
		builder = builder.WithBatchOps(report.Output.BatchOps)
	}

	return builder.Build(input.MCMS)
}

func validateActivateRMNAddresses(addresses []datastore.AddressRef, chainSelector uint64) error {
	if ref := datastore_utils.GetAddressRef(
		addresses,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		common_utils.RMNTimelockQualifier,
	); ref.Address == "" {
		return fmt.Errorf(
			"ownership transfer requires RMNMCMS RBACTimelock (qualifier %q) in datastore for chain %d",
			common_utils.RMNTimelockQualifier, chainSelector,
		)
	}
	if ref := datastore_utils.GetAddressRef(
		addresses,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		common_utils.UltraFastCurseMCMSQualifier,
	); ref.Address == "" {
		return fmt.Errorf(
			"Ultra Fast Curse RBACTimelock (qualifier %q) not found in datastore for chain %d",
			common_utils.UltraFastCurseMCMSQualifier, chainSelector,
		)
	}
	return nil
}
