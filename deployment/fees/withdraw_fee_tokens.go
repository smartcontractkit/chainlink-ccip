package fees

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// FeeTokenWithdrawal names a single fee token to withdraw and, optionally, how much of it.
//
// A nil Amount withdraws the whole accumulated balance of the token. A non-nil Amount withdraws
// exactly that many base units, and is only honoured by chain families whose withdrawal primitive
// takes an amount. EVM's withdrawFeeTokens always sweeps the full balance, so the EVM adapters
// reject a non-nil Amount rather than silently withdrawing more than was asked for.
type FeeTokenWithdrawal struct {
	Token  string   `json:"token" yaml:"token"`
	Amount *big.Int `json:"amount,omitempty" yaml:"amount,omitempty"`
}

// WithdrawFeeTokensForChain is a fee token withdrawal for a single chain. One entry may carry many
// fee tokens; the adapter withdraws each of them in turn.
//
// When Contracts is empty the adapter uses its default contract(s). When Contracts is provided the
// adapter withdraws only from the specified contracts, using the full ref (Type + Version +
// Qualifier) for unambiguous datastore resolution.
type WithdrawFeeTokensForChain struct {
	ChainSelector uint64                 `json:"chainSelector" yaml:"chainSelector"`
	FeeTokens     []FeeTokenWithdrawal   `json:"feeTokens" yaml:"feeTokens"`
	Contracts     []datastore.AddressRef `json:"contracts,omitempty" yaml:"contracts,omitempty"`
}

// WithdrawFeeTokensInput is the top-level input for the WithdrawFeeTokens changeset. Args may span
// as many chains as needed; each chain is resolved to its own family adapter.
type WithdrawFeeTokensInput struct {
	Version *semver.Version             `json:"version" yaml:"version"`
	Args    []WithdrawFeeTokensForChain `json:"args" yaml:"args"`
	MCMS    mcms.Input                  `json:"mcms" yaml:"mcms"`
}

// WithdrawFeeTokens creates a changeset that withdraws accumulated fee token balances to the
// configured fee aggregator on one or more chains.
func WithdrawFeeTokens() cldf.ChangeSetV2[WithdrawFeeTokensInput] {
	registry := GetFeeAggregatorRegistry()
	mcmsRegistry := changesets.GetRegistry()
	return cldf.CreateChangeSet(makeWithdrawFeeTokensApply(registry, mcmsRegistry), makeWithdrawFeeTokensVerify(registry))
}

func makeWithdrawFeeTokensVerify(registry *FeeAggregatorAdapterRegistry) func(cldf.Environment, WithdrawFeeTokensInput) error {
	return func(_ cldf.Environment, cfg WithdrawFeeTokensInput) error {
		if cfg.Version == nil {
			return fmt.Errorf("version is required")
		}
		if len(cfg.Args) == 0 {
			return fmt.Errorf("at least one chain must be specified")
		}

		seen := map[uint64]bool{}
		for i, arg := range cfg.Args {
			if seen[arg.ChainSelector] {
				return fmt.Errorf("duplicate chain selector at args[%d]: %d", i, arg.ChainSelector)
			}
			seen[arg.ChainSelector] = true

			if len(arg.FeeTokens) == 0 {
				return fmt.Errorf("at least one fee token is required at args[%d] (chain=%d)", i, arg.ChainSelector)
			}

			// Token address formats are family specific, so only emptiness, duplicates and amount
			// sanity are checked here; each adapter parses the addresses it understands.
			seenTokens := map[string]bool{}
			for j, tok := range arg.FeeTokens {
				if tok.Token == "" {
					return fmt.Errorf("fee token address is required at args[%d].feeTokens[%d] (chain=%d)", i, j, arg.ChainSelector)
				}
				if seenTokens[tok.Token] {
					return fmt.Errorf("duplicate fee token %q at args[%d].feeTokens[%d] (chain=%d)", tok.Token, i, j, arg.ChainSelector)
				}
				seenTokens[tok.Token] = true

				if tok.Amount != nil && tok.Amount.Sign() <= 0 {
					return fmt.Errorf("amount must be positive at args[%d].feeTokens[%d] (chain=%d, token=%s); omit it to withdraw the full balance",
						i, j, arg.ChainSelector, tok.Token)
				}
			}

			family, err := chain_selectors.GetSelectorFamily(arg.ChainSelector)
			if err != nil {
				return fmt.Errorf("failed to get chain family for selector %d: %w", arg.ChainSelector, err)
			}

			if _, exists := registry.GetFeeAggregatorAdapter(family, cfg.Version); !exists {
				return fmt.Errorf("no fee aggregator adapter found for chain family %s and version %s", family, cfg.Version.String())
			}
		}

		return nil
	}
}

func makeWithdrawFeeTokensApply(registry *FeeAggregatorAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, WithdrawFeeTokensInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg WithdrawFeeTokensInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, arg := range cfg.Args {
			family, err := chain_selectors.GetSelectorFamily(arg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", arg.ChainSelector, err)
			}

			adapter, exists := registry.GetFeeAggregatorAdapter(family, cfg.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no fee aggregator adapter found for chain family %s and version %s", family, cfg.Version.String())
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.WithdrawFeeTokens(e),
				e.BlockChains,
				WithdrawFeeTokensForChain{
					ChainSelector: arg.ChainSelector,
					FeeTokens:     arg.FeeTokens,
					Contracts:     arg.Contracts,
				},
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to withdraw fee tokens for chain %d: %w", arg.ChainSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			WithReports(reports).
			Build(cfg.MCMS)
	}
}
