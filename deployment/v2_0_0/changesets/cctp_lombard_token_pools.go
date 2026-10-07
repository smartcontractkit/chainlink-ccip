package changesets

import (
	"errors"
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// CCTP and Lombard pools are managed by their own adapters instead of the common token pool
// changesets. The changesets in this file give them the same token pool operations. Each entry is
// the sequence input the common changesets build, already resolved, and is run on the CCTP or
// Lombard adapter of the entry's chain family. The adapter must implement the optional interface
// of the operation (see CCTPChain and LombardChain).

// CCTPTokenPoolsConfig is the input for the CCTP token pool changesets.
type CCTPTokenPoolsConfig[In any] struct {
	// USDCType selects the CCTP adapter of each chain family.
	USDCType adapters.USDCType
	// Entries are the sequence inputs, one per pool (and per lane for rate limits).
	Entries []In
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

// LombardTokenPoolsConfig is the input for the Lombard token pool changesets.
type LombardTokenPoolsConfig[In any] struct {
	// Entries are the sequence inputs, one per pool (and per lane for rate limits).
	Entries []In
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

// RemoveCCTPRemotePools removes remote pool entries from CCTP token pools.
func RemoveCCTPRemotePools(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[CCTPTokenPoolsConfig[tokens.RemoveRemotePoolsSequenceInput]] {
	return newCCTPTokenPoolsChangeset(cctpChainRegistry, mcmsRegistry, "remove remote pools", removeRemotePoolsSelector, tokens.RemotePoolRemover.RemoveRemotePools, nil)
}

// SetCCTPTokenPoolRateLimits sets rate limits on CCTP token pools.
func SetCCTPTokenPoolRateLimits(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[CCTPTokenPoolsConfig[tokens.TPRLRemotes]] {
	return newCCTPTokenPoolsChangeset(cctpChainRegistry, mcmsRegistry, "set rate limits", rateLimitsSelector, tokens.TokenPoolRateLimitSetter.SetTokenPoolRateLimits, withExistingDataStore)
}

// SetCCTPTokenPoolDynamicConfig sets the router and admin roles of CCTP token pools.
func SetCCTPTokenPoolDynamicConfig(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[CCTPTokenPoolsConfig[tokens.SetTokenPoolDynamicConfigSequenceInput]] {
	return newCCTPTokenPoolsChangeset(cctpChainRegistry, mcmsRegistry, "set dynamic config", dynamicConfigSelector, tokens.TokenPoolDynamicConfigAdapter.SetTokenPoolDynamicConfig, nil)
}

// RemoveLombardRemotePools removes remote pool entries from Lombard token pools.
func RemoveLombardRemotePools(lombardChainRegistry *adapters.LombardChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[LombardTokenPoolsConfig[tokens.RemoveRemotePoolsSequenceInput]] {
	return newLombardTokenPoolsChangeset(lombardChainRegistry, mcmsRegistry, "remove remote pools", removeRemotePoolsSelector, tokens.RemotePoolRemover.RemoveRemotePools, nil)
}

// SetLombardTokenPoolRateLimits sets rate limits on Lombard token pools.
func SetLombardTokenPoolRateLimits(lombardChainRegistry *adapters.LombardChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[LombardTokenPoolsConfig[tokens.TPRLRemotes]] {
	return newLombardTokenPoolsChangeset(lombardChainRegistry, mcmsRegistry, "set rate limits", rateLimitsSelector, tokens.TokenPoolRateLimitSetter.SetTokenPoolRateLimits, withExistingDataStore)
}

// SetLombardTokenPoolDynamicConfig sets the router and admin roles of Lombard token pools.
func SetLombardTokenPoolDynamicConfig(lombardChainRegistry *adapters.LombardChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[LombardTokenPoolsConfig[tokens.SetTokenPoolDynamicConfigSequenceInput]] {
	return newLombardTokenPoolsChangeset(lombardChainRegistry, mcmsRegistry, "set dynamic config", dynamicConfigSelector, tokens.TokenPoolDynamicConfigAdapter.SetTokenPoolDynamicConfig, nil)
}

func removeRemotePoolsSelector(in tokens.RemoveRemotePoolsSequenceInput) uint64 { return in.Selector }

func rateLimitsSelector(in tokens.TPRLRemotes) uint64 { return in.ChainSelector }

func dynamicConfigSelector(in tokens.SetTokenPoolDynamicConfigSequenceInput) uint64 {
	return in.Selector
}

// withExistingDataStore fills the datastore the rate limit sequences read refs from.
func withExistingDataStore(e cldf.Environment, in tokens.TPRLRemotes) tokens.TPRLRemotes {
	if in.ExistingDataStore == nil {
		in.ExistingDataStore = e.DataStore
	}
	return in
}

// poolSequence returns the sequence of an operation from an adapter that supports it.
type poolSequence[Cap any, In any] func(Cap) *cldf_ops.Sequence[In, sequences.OnChainOutput, cldf_chain.BlockChains]

// poolAdapterResolver returns the CCTP or Lombard adapter registered for a chain family.
type poolAdapterResolver func(family string) (any, bool)

func newCCTPTokenPoolsChangeset[Cap any, In any](
	cctpChainRegistry *adapters.CCTPChainRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	operation string,
	selectorOf func(In) uint64,
	sequenceOf poolSequence[Cap, In],
	prepare func(cldf.Environment, In) In,
) cldf.ChangeSetV2[CCTPTokenPoolsConfig[In]] {
	apply := func(e cldf.Environment, cfg CCTPTokenPoolsConfig[In]) (cldf.ChangesetOutput, error) {
		resolve := func(family string) (any, bool) { return cctpChainRegistry.GetCCTPChain(family, cfg.USDCType) }
		return applyTokenPoolSequences(e, mcmsRegistry, cfg.MCMS, "CCTP", operation, resolve, cfg.Entries, selectorOf, sequenceOf, prepare)
	}
	verify := func(_ cldf.Environment, cfg CCTPTokenPoolsConfig[In]) error {
		if !cfg.USDCType.IsValid() {
			return fmt.Errorf("invalid USDC type: %q", cfg.USDCType)
		}
		return verifyTokenPoolEntries(cfg.MCMS, cfg.Entries, selectorOf)
	}
	return cldf.CreateChangeSet(apply, verify)
}

func newLombardTokenPoolsChangeset[Cap any, In any](
	lombardChainRegistry *adapters.LombardChainRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	operation string,
	selectorOf func(In) uint64,
	sequenceOf poolSequence[Cap, In],
	prepare func(cldf.Environment, In) In,
) cldf.ChangeSetV2[LombardTokenPoolsConfig[In]] {
	apply := func(e cldf.Environment, cfg LombardTokenPoolsConfig[In]) (cldf.ChangesetOutput, error) {
		resolve := func(family string) (any, bool) { return lombardChainRegistry.GetLombardChain(family) }
		return applyTokenPoolSequences(e, mcmsRegistry, cfg.MCMS, "Lombard", operation, resolve, cfg.Entries, selectorOf, sequenceOf, prepare)
	}
	verify := func(_ cldf.Environment, cfg LombardTokenPoolsConfig[In]) error {
		return verifyTokenPoolEntries(cfg.MCMS, cfg.Entries, selectorOf)
	}
	return cldf.CreateChangeSet(apply, verify)
}

func verifyTokenPoolEntries[In any](mcmsInput *mcms.Input, entries []In, selectorOf func(In) uint64) error {
	if mcmsInput != nil {
		if err := mcmsInput.Validate(); err != nil {
			return fmt.Errorf("failed to validate MCMS input: %w", err)
		}
	}
	if len(entries) == 0 {
		return errors.New("input must contain at least one entry")
	}
	for _, in := range entries {
		if _, err := chain_selectors.GetSelectorFamily(selectorOf(in)); err != nil {
			return err
		}
	}
	return nil
}

// applyTokenPoolSequences runs the operation's sequence for each entry on the adapter of the
// entry's chain family, and builds one output from all of them.
func applyTokenPoolSequences[Cap any, In any](
	e cldf.Environment,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	mcmsInput *mcms.Input,
	poolKind string,
	operation string,
	resolve poolAdapterResolver,
	entries []In,
	selectorOf func(In) uint64,
	sequenceOf poolSequence[Cap, In],
	prepare func(cldf.Environment, In) In,
) (cldf.ChangesetOutput, error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	for _, in := range entries {
		chainSel := selectorOf(in)
		family, err := chain_selectors.GetSelectorFamily(chainSel)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainSel, err)
		}
		adapter, ok := resolve(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no %s adapter registered for chain family '%s'", poolKind, family)
		}
		capable, ok := adapter.(Cap)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("%s adapter for chain family '%s' does not support %s", poolKind, family, operation)
		}
		if prepare != nil {
			in = prepare(e, in)
		}
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequenceOf(capable), e.BlockChains, in)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to %s on %s token pool on chain with selector %d: %w", operation, poolKind, chainSel, err)
		}
		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)
	}

	var input mcms.Input
	if mcmsInput != nil {
		input = *mcmsInput
	}
	return changesets.NewOutputBuilder(e, mcmsRegistry).
		WithReports(reports).
		WithBatchOps(batchOps).
		Build(input)
}
