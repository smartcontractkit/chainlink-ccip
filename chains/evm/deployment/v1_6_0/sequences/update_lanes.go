package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type UpdateLanesSequenceInput struct {
	FeeQuoterApplyDestChainConfigUpdatesSequenceInput FeeQuoterApplyDestChainConfigUpdatesSequenceInput
	FeeQuoterUpdatePricesSequenceInput                FeeQuoterUpdatePricesSequenceInput
	OffRampApplySourceChainConfigUpdatesSequenceInput OffRampApplySourceChainConfigUpdatesSequenceInput
	OnRampApplyDestChainConfigUpdatesSequenceInput    OnRampApplyDestChainConfigUpdatesSequenceInput
	RouterApplyRampUpdatesSequenceInput               RouterApplyRampUpdatesSequenceInput
}

var ConfigureLaneLegAsDest = operations.NewSequence(
	"ConfigureLaneLegAsDest",
	semver.MustParse("1.0.0"),
	"Updates lanes on CCIP 1.6.0",
	func(b operations.Bundle, chains map[uint64]cldf_evm.Chain, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput

		result, err := runAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, input.FeeQuoterApplyDestChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on FeeQuoters")

		result, err = runAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, input.FeeQuoterUpdatePricesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Gas prices updated on FeeQuoters")

		result, err = runAndMergeSequence(b, chains, OffRampApplySourceChainConfigUpdatesSequence, input.OffRampApplySourceChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on OnRamps")

		result, err = runAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, input.OnRampApplyDestChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Source configs updated on OffRamps")

		result, err = runAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, input.RouterApplyRampUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Ramps updated on Routers")

		return result, nil
	},
)

var ConfigureLaneLegAsSource = operations.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.6.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b operations.Bundle, chains map[uint64]cldf_evm.Chain, input UpdateLanesSequenceInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput

		result, err := runAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, input.FeeQuoterApplyDestChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on FeeQuoters")

		result, err = runAndMergeSequence(b, chains, FeeQuoterUpdatePricesSequence, input.FeeQuoterUpdatePricesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Gas prices updated on FeeQuoters")

		result, err = runAndMergeSequence(b, chains, OffRampApplySourceChainConfigUpdatesSequence, input.OffRampApplySourceChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Destination configs updated on OnRamps")

		result, err = runAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, input.OnRampApplyDestChainConfigUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Source configs updated on OffRamps")

		result, err = runAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, input.RouterApplyRampUpdatesSequenceInput, result)
		if err != nil {
			return result, err
		}
		b.Logger.Info("Ramps updated on Routers")

		return result, nil
	},
)

func runAndMergeSequence[IN any](
	b operations.Bundle,
	chains map[uint64]cldf_evm.Chain,
	seq *operations.Sequence[IN, sequences.OnChainOutput, map[uint64]cldf_evm.Chain],
	input IN,
	agg sequences.OnChainOutput,
) (sequences.OnChainOutput, error) {
	report, err := operations.ExecuteSequence(b, seq, chains, input)
	if err != nil {
		return agg, fmt.Errorf("failed to execute %s: %w", seq.ID(), err)
	}
	agg.BatchOps = append(agg.BatchOps, report.Output.BatchOps...)
	agg.Addresses = append(agg.Addresses, report.Output.Addresses...)
	return agg, nil
}
