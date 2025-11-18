package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type DeployUSDCTokenPoolProxyInput struct {
	ChainInputs []DeployUSDCTokenPoolProxyPerChainInput
	MCMS        mcms.Input
}

type DeployUSDCTokenPoolProxyPerChainInput struct {
	ChainSelector uint64
	Token         common.Address
	PoolAddresses usdc_token_pool_proxy_ops.PoolAddresses
	Router        common.Address
}

func DeployUSDCTokenPoolProxyChangeset() cldf.ChangeSetV2[DeployUSDCTokenPoolProxyInput] {
	return cldf.CreateChangeSet(deployUSDCTokenPoolProxyApply(), deployUSDCTokenPoolProxyVerify())
}

func deployUSDCTokenPoolProxyApply() func(cldf.Environment, DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		// Execute the sequence for each chain input
		for _, perChainInput := range input.ChainInputs {
			sequenceInput := sequences.DeployUSDCTokenPoolProxySequenceInput{
				ChainSelector: perChainInput.ChainSelector,
				Token:         perChainInput.Token,
				PoolAddresses: perChainInput.PoolAddresses,
				Router:        perChainInput.Router,
			}

			// Note: Other example deployment changesets have a deployer Registry check.
			// This does not. It may be necessary at a later date.

			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployUSDCTokenPoolProxySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy on chain %d: %w", perChainInput.ChainSelector, err)
			}

			for _, r := range deployReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}

			reports = append(reports, deployReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(input.MCMS)
	}
}

func deployUSDCTokenPoolProxyVerify() func(cldf.Environment, DeployUSDCTokenPoolProxyInput) error {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
