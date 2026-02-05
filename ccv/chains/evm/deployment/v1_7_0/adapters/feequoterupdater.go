package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	sequence1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type FeeQuoterUpdater struct{}

func (fqu FeeQuoterUpdater) SequenceFeeQuoterInputCreation() *cldf_ops.Sequence[deploy.FeeQuoterUpdateInput, sequence1_7.FeeQuoterUpdate, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"fee-quoter-updater:input-creation",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdateInput for FeeQuoter update sequence",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.FeeQuoterUpdateInput) (output sequence1_7.FeeQuoterUpdate, err error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
			}
			switch input.ImportConfigFromVersion {
			case semver.MustParse("1.6.0"):
				report, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV160, chain, input)
				if err != nil {
					return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.6.0: %w", err)
				}
				output = report.Output
				output.ChainSelector = input.ChainSelector
				output.ExistingAddresses = input.ExistingAddresses
				return
			case semver.MustParse("1.5.0"):
				report, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV150, chain, input)
				if err != nil {
					return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.5.0: %w", err)
				}
				output = report.Output
				output.ChainSelector = input.ChainSelector
				output.ExistingAddresses = input.ExistingAddresses
				return
			default:
				return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("unsupported ImportFeeQuoterConfigFromVersion: %s", input.ImportConfigFromVersion.String())
			}
		},
	)
}

func (fqu FeeQuoterUpdater) SequenceDeployOrUpdateFeeQuoter() *cldf_ops.Sequence[sequence1_7.FeeQuoterUpdate, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return sequence1_7.SequenceFeeQuoterUpdate
}
