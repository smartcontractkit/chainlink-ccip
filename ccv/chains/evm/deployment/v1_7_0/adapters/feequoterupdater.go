package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
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
			// get the FeeQuoterUpdateOutput from both v1.6.0 and v1.5.0 sequences and combine them to create the input for the fee quoter update sequence
			report, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV163, chain, input)
			if err != nil {
				return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.6.0: %w", err)
			}
			output16 := report.Output

			report15, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV150, chain, input)
			if err != nil {
				return sequence1_7.FeeQuoterUpdate{}, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.5.0: %w", err)
			}
			output15 := report15.Output
			// combine the outputs from both sequences to create the input for the fee quoter update sequence
			output.DestChainConfigs = append(output16.DestChainConfigs, output15.DestChainConfigs...)
			output.TokenTransferFeeConfigUpdates = fqops.ApplyTokenTransferFeeConfigUpdatesArgs{
				TokenTransferFeeConfigArgs: append(
					output16.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs,
					output15.TokenTransferFeeConfigUpdates.TokenTransferFeeConfigArgs...),
				TokensToUseDefaultFeeConfigs: append(
					output16.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs,
					output15.TokenTransferFeeConfigUpdates.TokensToUseDefaultFeeConfigs...),
			}
			output.AuthorizedCallerUpdates.AddedCallers = append(
				output16.AuthorizedCallerUpdates.AddedCallers,
				output15.AuthorizedCallerUpdates.AddedCallers...)

			// constructor args from the v1.6.0 sequence takes precedence over the constructor args from the v1.5.0 sequence,
			// as FeeQuoter 1.6.3 is more closely related to FeeQuoter 1.7.0
			// if the constructor args from the v1.6.0 sequence is empty, it means there is no FeeQuoter 1.6 exists for that chain
			// use the constructor args from the v1.5.0 sequence in that case
			if output16.ConstructorArgs.IsEmpty() {
				output.ConstructorArgs = output15.ConstructorArgs
			}

			// check if output is empty, if so, return an error
			if empty, err := output.IsEmpty(); err != nil || empty {
				return sequence1_7.FeeQuoterUpdate{},
					fmt.Errorf("could not create input for fee quoter 1.7 update sequence: %w", err)
			}
			output.ChainSelector = input.ChainSelector
			output.ExistingAddresses = input.ExistingAddresses
			return output, nil
		},
	)
}

func (fqu FeeQuoterUpdater) SequenceDeployOrUpdateFeeQuoter() *cldf_ops.Sequence[sequence1_7.FeeQuoterUpdate, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return sequence1_7.SequenceFeeQuoterUpdate
}
