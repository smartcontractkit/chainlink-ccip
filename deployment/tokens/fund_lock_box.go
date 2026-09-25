package tokens

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// FundLockBoxConfig is the configuration for the FundLockBox changeset.
type FundLockBoxConfig struct {
	// Funding specifies the lockbox funding operations to perform.
	Funding []FundLockBoxPerLockBox
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// FundLockBoxPerLockBox groups deposits for a single lockbox on a single chain.
type FundLockBoxPerLockBox struct {
	// ChainSelector identifies the chain on which the lockbox lives.
	ChainSelector uint64
	// LockBoxRef references the ERC20LockBox to fund.
	LockBoxRef datastore.AddressRef
	// TokenRef references the token to deposit into the lockbox.
	TokenRef datastore.AddressRef
	// Deposits are the individual deposits to make. Each entry targets one bucket: a siloed bucket
	// (RemoteChainSelector set to the remote chain) or the unsiloed shared bucket
	// (RemoteChainSelector zero).
	Deposits []LockBoxDeposit
	// UsePlainTransfer, when true, transfers tokens directly to the lockbox via ERC20.transfer
	// instead of using the lockbox's deposit() function. This bypasses the Deposit event emission
	// and the per-bucket accounting. Use only as a break-glass option.
	UsePlainTransfer bool
}

// FundLockBox returns a changeset that funds a v2.0 ERC20LockBox directly, without a legacy pool to
// migrate from. It tops up a lockbox's siloed (per-chain-selector) and unsiloed (shared) buckets.
//
// All writes are proposal-only: the timelock must be an authorized caller on the lockbox to
// deposit, so the changeset authorizes it (idempotently) and then approves + deposits. The tokens
// are expected to be held by the timelock, since the deposit pulls from the caller's balance.
func FundLockBox(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[FundLockBoxConfig] {
	return cldf.CreateChangeSet(makeFundLockBoxApply(tokenRegistry, mcmsRegistry), makeFundLockBoxVerify())
}

func makeFundLockBoxVerify() func(cldf.Environment, FundLockBoxConfig) error {
	return func(_ cldf.Environment, cfg FundLockBoxConfig) error {
		if len(cfg.Funding) == 0 {
			return fmt.Errorf("at least one lockbox funding entry is required")
		}

		type lockBoxKey struct {
			selector uint64
			ref      string
		}
		seenLockBoxes := make(map[lockBoxKey]struct{})

		for i, entry := range cfg.Funding {
			if _, err := chain_selectors.GetSelectorFamily(entry.ChainSelector); err != nil {
				return fmt.Errorf("funding[%d]: invalid chain selector %d: %w", i, entry.ChainSelector, err)
			}
			if datastore_utils.IsAddressRefEmpty(entry.LockBoxRef) {
				return fmt.Errorf("funding[%d]: lockbox ref must not be empty", i)
			}
			if datastore_utils.IsAddressRefEmpty(entry.TokenRef) {
				return fmt.Errorf("funding[%d]: token ref must not be empty", i)
			}
			if len(entry.Deposits) == 0 {
				return fmt.Errorf("funding[%d]: at least one deposit is required", i)
			}

			seenBuckets := make(map[uint64]struct{}, len(entry.Deposits))
			for j, deposit := range entry.Deposits {
				if deposit.Amount == nil || deposit.Amount.Sign() <= 0 {
					return fmt.Errorf("funding[%d].deposits[%d]: amount must be positive", i, j)
				}
				if _, dup := seenBuckets[deposit.RemoteChainSelector]; dup {
					return fmt.Errorf("funding[%d]: duplicate remote chain selector %d in deposits", i, deposit.RemoteChainSelector)
				}
				seenBuckets[deposit.RemoteChainSelector] = struct{}{}
			}

			key := lockBoxKey{selector: entry.ChainSelector, ref: datastore_utils.SprintRef(entry.LockBoxRef)}
			if _, dup := seenLockBoxes[key]; dup {
				return fmt.Errorf("funding[%d]: duplicate lockbox entry for chain selector %d and ref %s", i, entry.ChainSelector, datastore_utils.SprintRef(entry.LockBoxRef))
			}
			seenLockBoxes[key] = struct{}{}
		}

		return nil
	}
}

func makeFundLockBoxApply(_ *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, FundLockBoxConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg FundLockBoxConfig) (cldf.ChangesetOutput, error) {
		tokenRegistry := GetTokenAdapterRegistry()
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for i, entry := range cfg.Funding {
			family, err := chain_selectors.GetSelectorFamily(entry.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: failed to get chain family for selector %d: %w", i, entry.ChainSelector, err)
			}

			lockBoxRef, err := ResolveTokenPoolRef(e, tokenRegistry, entry.ChainSelector, entry.LockBoxRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: failed to resolve lockbox ref: %w", i, err)
			}
			tokenRef, err := ResolveTokenRef(e, tokenRegistry, entry.ChainSelector, entry.TokenRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: failed to resolve token ref: %w", i, err)
			}

			adapter, ok := tokenRegistry.GetTokenAdapter(family, lockBoxRef.Version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: no token adapter registered for chain family '%s' and version '%v'", i, family, lockBoxRef.Version)
			}

			funder, ok := adapter.(LockBoxFunder)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: adapter for family '%s' version '%v' does not support lockbox funding", i, family, lockBoxRef.Version)
			}
			fundSeq := funder.FundLockBoxSequence()
			if fundSeq == nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: adapter for family '%s' version '%v' does not support lockbox funding", i, family, lockBoxRef.Version)
			}

			// Derive the timelock address from the MCMS config.
			mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: no MCMS reader registered for chain family '%s'", i, family)
			}
			timelockRef, err := mcmsReader.GetTimelockRef(e, entry.ChainSelector, cfg.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: failed to get timelock address from MCMS config: %w", i, err)
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, fundSeq, e.BlockChains, FundLockBoxSequenceInput{
				ChainSelector:    entry.ChainSelector,
				LockBoxAddress:   lockBoxRef.Address,
				TokenAddress:     tokenRef.Address,
				TimelockAddress:  timelockRef.Address,
				Deposits:         entry.Deposits,
				UsePlainTransfer: entry.UsePlainTransfer,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("funding[%d]: failed to fund lockbox %s on chain selector %d: %w", i, lockBoxRef.Address, entry.ChainSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
