package common

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-ccip/devenv/sequences"
)

var DeployHomeChain = deployment.CreateChangeSet(applyDeployHomeChain, validateDeployHomeChain)

func validateDeployHomeChain(e deployment.Environment, cfg sequences.DeployHomeChainConfig) error {
	return nil
}

func applyDeployHomeChain(e deployment.Environment, cfg sequences.DeployHomeChainConfig) (deployment.ChangesetOutput, error) {
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployHomeChain,
		e.BlockChains,
		cfg,
	)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	ds := datastore.NewMemoryDataStore()
	for _, addr := range report.Output.Addresses {
		ds.AddressRefStore.Add(addr)
	}
	return deployment.ChangesetOutput{
		DataStore: ds,
	}, nil
}
