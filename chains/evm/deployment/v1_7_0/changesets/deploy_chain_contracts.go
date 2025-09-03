package changesets

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var DeployChainContracts = cldf_deployment.CreateChangeSet(applyDeployChainContracts, verifyDeployChainContracts)

type DeployChainContractsCfg struct {
	ChainSelector uint64
	Params        sequences.ContractParams
}

func verifyDeployChainContracts(e cldf_deployment.Environment, cfg DeployChainContractsCfg) error {
	// TODO: Verify inputs, environment state, etc.
	evmChains := e.BlockChains.EVMChains()
	if _, exists := evmChains[cfg.ChainSelector]; !exists {
		return fmt.Errorf("no EVM chain with selector %d found in environment", cfg.ChainSelector)
	}
	return nil
}

func applyDeployChainContracts(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (cldf_deployment.ChangesetOutput, error) {
	addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSelector))
	chain := e.BlockChains.EVMChains()[cfg.ChainSelector]

	report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, chain, sequences.DeployChainContractsInput{
		ChainSelector:     cfg.ChainSelector,
		ExistingAddresses: addresses,
		ContractParams:    cfg.Params,
	})
	if err != nil {
		return cldf_deployment.ChangesetOutput{Reports: report.ExecutionReports}, fmt.Errorf("failed to execute DeployChain sequence: %w", err)
	}

	ds := datastore.NewMemoryDataStore()
	for _, r := range report.Output.Addresses {
		if err := ds.Addresses().Add(r); err != nil {
			return cldf_deployment.ChangesetOutput{Reports: report.ExecutionReports}, fmt.Errorf("failed to add %s %s to datastore: %w", r.Type, r.Version, err)
		}
	}

	return changesets.NewOutputBuilder().
		WithReports(report.ExecutionReports).
		WithDataStore(ds).
		WithWriteOutputs(report.Output.Writes).
		Build(changesets.MCMSParams{
			Description: fmt.Sprintf("Initial confiuration of 1.7.0 contracts on %s", chain),
			// TODO: Populate these with correct values later
			OverridePreviousRoot: false,
			ValidUntil:           2756219818,
			TimelockDelay:        mcms_types.NewDuration(3 * time.Hour),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			TimelockAddresses:    nil,
			ChainMetadata:        nil,
		})
}
