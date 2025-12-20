package common

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployHomeChain = deployment.CreateChangeSet(applyDeployHomeChain, validateDeployHomeChain)

func validateDeployHomeChain(e deployment.Environment, cfg DeployHomeChainConfig) error {
	return nil
}

func applyDeployHomeChain(e deployment.Environment, cfg DeployHomeChainConfig) (deployment.ChangesetOutput, error) {
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		SeqDeployHomeChain,
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
