package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	sequence1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
)

// FeeQuoterUpdater uses FeeQUpdateArgs any so it implements deploy.FeeQuoterUpdater[any] and can be registered directly.
// The implementation is for v1.7.0 and uses sequence1_7.FeeQuoterUpdate internally.
type FeeQuoterUpdater[FeeQUpdateArgs any] struct{}

func (fqu FeeQuoterUpdater[FeeQUpdateArgs]) SequenceFeeQuoterInputCreation() *cldf_ops.Sequence[deploy.FeeQuoterUpdateInput, FeeQUpdateArgs, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"fee-quoter-updater:input-creation",
		semver.MustParse("1.7.0"),
		"Creates FeeQuoterUpdateInput for FeeQuoter update sequence",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.FeeQuoterUpdateInput) (output FeeQUpdateArgs, err error) {
			var zero FeeQUpdateArgs
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return zero, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
			}
			// get the FeeQuoterUpdateOutput from both v1.6.0 and v1.5.0 sequences and combine them to create the input for the fee quoter update sequence
			report, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV163, chain, input)
			if err != nil {
				return zero, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.6.0: %w", err)
			}
			output16 := report.Output

			report15, err := cldf_ops.ExecuteSequence(b, sequence1_7.CreateFeeQuoterUpdateInputFromV150, chain, input)
			if err != nil {
				return zero, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.5.0: %w", err)
			}
			output15 := report15.Output
			// combine the outputs from both sequences to create the input for the fee quoter update sequence
			out, err := sequence1_7.MergeFeeQuoterUpdateOutputs(output16, output15)
			if err != nil {
				return zero, fmt.Errorf("failed to merge FeeQuoterUpdateInput from v1.6.0 and v1.5.0: %w", err)
			}
			// check if output is empty, if so, return an error
			if empty, err := out.IsEmpty(); err != nil || empty {
				return zero, fmt.Errorf("could not create input for fee quoter 1.7 update sequence: %w", err)
			}
			out.ChainSelector = input.ChainSelector
			out.ExistingAddresses = input.ExistingAddresses
			return any(out).(FeeQUpdateArgs), nil
		},
	)
}

func (fqu FeeQuoterUpdater[FeeQUpdateArgs]) SequenceDeployOrUpdateFeeQuoter() *cldf_ops.Sequence[FeeQUpdateArgs, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"fee-quoter-v1.7.0:update-sequence",
		semver.MustParse("1.7.0"),
		"Deploys or fetches existing FeeQuoter contract and applies config updates",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input FeeQUpdateArgs) (output sequences.OnChainOutput, err error) {
			fqInput, ok := any(input).(sequence1_7.FeeQuoterUpdate)
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("expected sequence1_7.FeeQuoterUpdate, got %T", input)
			}
			report, err := cldf_ops.ExecuteSequence(b, sequence1_7.SequenceFeeQuoterUpdate, chains, fqInput)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			return report.Output, nil
		},
	)
}
