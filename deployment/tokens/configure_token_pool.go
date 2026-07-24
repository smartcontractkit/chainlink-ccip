package tokens

import (
	"errors"
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// ConfigureTokenPoolInput is the input for the ConfigureTokenPool changeset. It applies small,
// targeted configuration changes to existing token pools. Unlike TokenExpansion and
// SetTokenPoolRateLimits, it has no bidirectionality constraints: each entry configures one
// pool's local view of a lane and never requires a counterpart section. Counterpart entries
// MAY still be provided (to configure both sides in one changeset); no symmetry is enforced.
type ConfigureTokenPoolInput struct {
	// Chains lists the per-chain pool configuration updates.
	Chains []ConfigureTokenPoolPerChain `yaml:"chains" json:"chains"`
	// MCMS configures the resulting proposal.
	MCMS mcms.Input `yaml:"mcms,omitempty" json:"mcms"`
}

// ConfigureTokenPoolPerChain groups pool updates for a single chain.
type ConfigureTokenPoolPerChain struct {
	// ChainSelector identifies the chain on which the pools live.
	ChainSelector uint64 `yaml:"selector,string" json:"selector,string"`
	// Pools lists the pool configuration updates on this chain.
	Pools []PoolConfigUpdate `yaml:"pools" json:"pools"`
}

// PoolConfigUpdate describes a partial configuration update for a single token pool.
// Every field other than TokenPoolRef is optional: absent fields leave on-chain state
// untouched. To clear a value, provide it explicitly (e.g. the zero address).
type PoolConfigUpdate struct {
	// TokenPoolRef is a reference to the token pool in the datastore.
	TokenPoolRef datastore.AddressRef `yaml:"tokenPoolRef" json:"tokenPoolRef"`
	// FinalityConfig, if set, is the allowed finality config to set on the pool (v2+ only).
	FinalityConfig *finality.Config `yaml:"finalityConfig,omitempty" json:"finalityConfig,omitempty"`
	// RateLimitAdmin, if set, is the desired rate limit admin address.
	RateLimitAdmin *string `yaml:"rateLimitAdmin,omitempty" json:"rateLimitAdmin,omitempty"`
	// FeeAdmin, if set, is the desired fee admin address (v2+ only).
	FeeAdmin *string `yaml:"feeAdmin,omitempty" json:"feeAdmin,omitempty"`
	// Remotes lists per-lane configuration updates.
	Remotes []RemoteConfigUpdate `yaml:"remotes,omitempty" json:"remotes,omitempty"`
}

// RemoteConfigUpdate describes partial per-lane configuration for one remote chain.
type RemoteConfigUpdate struct {
	// RemoteChainSelector identifies the remote chain of the lane.
	RemoteChainSelector uint64 `yaml:"selector,string" json:"selector,string"`
	// TokenTransferFeeConfig, if set, is merged with the current on-chain fee config
	// (user-set fields win; unset fields keep their on-chain values).
	TokenTransferFeeConfig *PartialTokenTransferFeeConfig `yaml:"tokenTransferFeeConfig,omitempty" json:"tokenTransferFeeConfig,omitempty"`
}

// ConfigureTokenPool returns a changeset that applies partial configuration updates to
// existing token pools.
func ConfigureTokenPool() cldf.ChangeSetV2[ConfigureTokenPoolInput] {
	return cldf.CreateChangeSet(configureTokenPoolApply(), configureTokenPoolVerify())
}

func configureTokenPoolVerify() func(cldf.Environment, ConfigureTokenPoolInput) error {
	return func(e cldf.Environment, cfg ConfigureTokenPoolInput) error {
		if len(cfg.Chains) == 0 {
			return errors.New("input must contain at least one chain entry")
		}
		// Structural checks only — no datastore/on-chain resolution. Obvious input mistakes
		// (bad selectors, empty updates, duplicate entries) surface before apply runs.
		type chainConfigKey struct {
			poolRef  datastore.AddressRefKey
			selector uint64
		}
		seenPools := make(map[chainConfigKey]struct{})
		for _, chainCfg := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chainCfg.ChainSelector); err != nil {
				return fmt.Errorf("invalid chain selector %d: %w", chainCfg.ChainSelector, err)
			}
			if len(chainCfg.Pools) == 0 {
				return fmt.Errorf("no pools provided for chain selector %d", chainCfg.ChainSelector)
			}
			for _, pool := range chainCfg.Pools {
				if datastore_utils.IsAddressRefEmpty(pool.TokenPoolRef) {
					return fmt.Errorf("pool entry on chain selector %d has an empty tokenPoolRef", chainCfg.ChainSelector)
				}
				if pool.TokenPoolRef.ChainSelector != 0 && pool.TokenPoolRef.ChainSelector != chainCfg.ChainSelector {
					return fmt.Errorf("pool entry %s has tokenPoolRef.chainSelector %d that does not match the enclosing chain selector %d", datastore_utils.SprintRef(pool.TokenPoolRef), pool.TokenPoolRef.ChainSelector, chainCfg.ChainSelector)
				}
				if pool.FinalityConfig == nil && pool.RateLimitAdmin == nil && pool.FeeAdmin == nil && len(pool.Remotes) == 0 {
					return fmt.Errorf("pool entry %s on chain selector %d has no fields to update", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector)
				}
				if pool.FinalityConfig != nil {
					if err := pool.FinalityConfig.Validate(); err != nil {
						return fmt.Errorf("finality config for pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector, err)
					}
				}
				seenRemotes := make(map[uint64]struct{})
				for _, remote := range pool.Remotes {
					if remote.RemoteChainSelector == chainCfg.ChainSelector {
						return fmt.Errorf("remote chain selector %d must not equal the pool's own chain selector", remote.RemoteChainSelector)
					}
					if _, err := chain_selectors.GetSelectorFamily(remote.RemoteChainSelector); err != nil {
						return fmt.Errorf("invalid remote chain selector %d: %w", remote.RemoteChainSelector, err)
					}
					if _, dup := seenRemotes[remote.RemoteChainSelector]; dup {
						return fmt.Errorf("duplicate remote chain selector %d for pool on chain selector %d", remote.RemoteChainSelector, chainCfg.ChainSelector)
					}
					seenRemotes[remote.RemoteChainSelector] = struct{}{}
					if remote.TokenTransferFeeConfig == nil {
						return fmt.Errorf("remote entry %d for pool on chain selector %d has nothing to update", remote.RemoteChainSelector, chainCfg.ChainSelector)
					}
					if v, ok := remote.TokenTransferFeeConfig.DestBytesOverhead.Get(); ok && v < 32 {
						return fmt.Errorf("destBytesOverhead must be at least 32 for remote %d on chain selector %d, got %d", remote.RemoteChainSelector, chainCfg.ChainSelector, v)
					}
				}
				key := chainConfigKey{poolRef: pool.TokenPoolRef.Key(), selector: chainCfg.ChainSelector}
				if _, dup := seenPools[key]; dup {
					return fmt.Errorf("duplicate pool entry for chain selector %d and ref %s", chainCfg.ChainSelector, datastore_utils.SprintRef(pool.TokenPoolRef))
				}
				seenPools[key] = struct{}{}
			}
		}
		return nil
	}
}

func configureTokenPoolApply() func(cldf.Environment, ConfigureTokenPoolInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokenPoolInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		for _, chainCfg := range cfg.Chains {
			selector := chainCfg.ChainSelector
			for _, pool := range chainCfg.Pools {
				adapter, family, fullPoolRef, fullTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, selector, pool.TokenPoolRef, datastore.AddressRef{})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.TokenPoolRef), selector, err)
				}

				if pool.FinalityConfig != nil {
					feeAdapter, ok := adapter.(TokenFeeAdapter)
					if !ok {
						return cldf.ChangesetOutput{}, fmt.Errorf(
							"adapter for chain selector %d (family %s, version %s) does not support finality config updates",
							selector, family, fullPoolRef.Version,
						)
					}
					report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, feeAdapter.SetAllowedFinalityConfig(&e), e.BlockChains, SetAllowedFinalityConfigSequenceInput{
						Selector: selector,
						Settings: map[string]finality.Config{fullPoolRef.Address: *pool.FinalityConfig},
					})
					if err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to set finality config on pool %s: %w", fullPoolRef.Address, err)
					}
					batchOps = append(batchOps, report.Output.BatchOps...)
					reports = append(reports, report.ExecutionReports...)
				}

				if pool.RateLimitAdmin != nil || pool.FeeAdmin != nil {
					adminAdapter, ok := adapter.(TokenPoolAdminAdapter)
					if !ok {
						return cldf.ChangesetOutput{}, fmt.Errorf(
							"adapter for chain selector %d (family %s, version %s) does not support admin role updates",
							selector, family, fullPoolRef.Version,
						)
					}
					report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adminAdapter.SetTokenPoolAdmins(), e.BlockChains, SetTokenPoolAdminsSequenceInput{
						Selector:       selector,
						PoolAddress:    fullPoolRef.Address,
						RateLimitAdmin: pool.RateLimitAdmin,
						FeeAdmin:       pool.FeeAdmin,
					})
					if err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to set admin roles on pool %s: %w", fullPoolRef.Address, err)
					}
					batchOps = append(batchOps, report.Output.BatchOps...)
					reports = append(reports, report.ExecutionReports...)
				}

				for _, remote := range pool.Remotes {
					if remote.TokenTransferFeeConfig != nil {
						feeBatchOps, feeReports, err := applyTokenTransferFeeConfig(e, selector, remote.RemoteChainSelector, fullPoolRef, fullTokenRef, *remote.TokenTransferFeeConfig)
						if err != nil {
							return cldf.ChangesetOutput{}, fmt.Errorf("failed to apply fee config for remote chain selector %d: %w", remote.RemoteChainSelector, err)
						}
						batchOps = append(batchOps, feeBatchOps...)
						reports = append(reports, feeReports...)
					}
				}
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
