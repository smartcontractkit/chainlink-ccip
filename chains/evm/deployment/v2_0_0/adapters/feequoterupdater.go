package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	fqseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
)

// FeeQuoterUpdater uses FeeQUpdateArgs any so it implements deploy.FeeQuoterUpdater[any] and can be registered directly.
// The implementation is for v2.0.0 and uses fqseq.FeeQuoterUpdate internally.
type FeeQuoterUpdater[FeeQUpdateArgs any] struct{}

func (fqu FeeQuoterUpdater[FeeQUpdateArgs]) SequenceFeeQuoterInputCreation() *cldf_ops.Sequence[deploy.FeeQuoterUpdateInput, FeeQUpdateArgs, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"fee-quoter-updater:input-creation",
		semver.MustParse("2.0.0"),
		"Creates FeeQuoterUpdateInput for FeeQuoter update sequence",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.FeeQuoterUpdateInput) (output FeeQUpdateArgs, err error) {
			var zero FeeQUpdateArgs
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return zero, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
			}
			var output16, output15 fqseq.FeeQuoterUpdate
			for _, version := range input.PreviousVersions {
				switch version.String() {
				case "1.6.0":
					report, err := cldf_ops.ExecuteSequence(b, fqseq.CreateFeeQuoterUpdateInputFromV16x, chain, input)
					if err != nil {
						return zero, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.6.0: %w", err)
					}
					output16 = report.Output
				case "1.5.0":
					report15, err := cldf_ops.ExecuteSequence(b, fqseq.CreateFeeQuoterUpdateInputFromV150, chain, input)
					if err != nil {
						return zero, fmt.Errorf("failed to create FeeQuoterUpdateInput from v1.5.0: %w", err)
					}
					output15 = report15.Output
				}
			}
			// combine the outputs from both sequences to create the input for the fee quoter update sequence
			out, err := fqseq.MergeFeeQuoterUpdateOutputs(output16, output15)
			if err != nil {
				return zero, fmt.Errorf("failed to merge FeeQuoterUpdateInput from v1.6.0 and v1.5.0: %w", err)
			}
			// check if output is empty, if so, return an error
			empty, err := out.IsEmpty()
			if err != nil {
				return zero, fmt.Errorf("could not create input for fee quoter 2.0.0 update sequence: %w", err)
			}
			if empty {
				return zero, fmt.Errorf("could not create input for fee quoter 2.0.0 update sequence: output is empty")
			}

			out.ChainSelector = input.ChainSelector
			out.ExistingAddresses = input.ExistingAddresses

			if input.TimelockAddress != "" {
				timelockAddr := common.HexToAddress(input.TimelockAddress)
				if !fqseq.IsConstructorArgsEmpty(out.ConstructorArgs) {
					out.ConstructorArgs.PriceUpdaters = append(out.ConstructorArgs.PriceUpdaters, timelockAddr)
				} else {
					out.AuthorizedCallerUpdates.AddedCallers = append(out.AuthorizedCallerUpdates.AddedCallers, timelockAddr)
				}
			}

			return any(out).(FeeQUpdateArgs), nil
		},
	)
}

func (fqu FeeQuoterUpdater[FeeQUpdateArgs]) SequenceDeployOrUpdateFeeQuoter() *cldf_ops.Sequence[FeeQUpdateArgs, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"fee-quoter-v2.0.0:update-sequence",
		semver.MustParse("2.0.0"),
		"Deploys or fetches existing FeeQuoter contract and applies config updates",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input FeeQUpdateArgs) (output sequences.OnChainOutput, err error) {
			fqInput, ok := any(input).(fqseq.FeeQuoterUpdate)
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("expected fqseq.FeeQuoterUpdate, got %T", input)
			}
			report, err := cldf_ops.ExecuteSequence(b, fqseq.SequenceFeeQuoterUpdate, chains, fqInput)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			return report.Output, nil
		},
	)
}
