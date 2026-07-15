package tokens

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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
	Chains []ConfigureTokenPoolPerChain `yaml:"input" json:"input"`
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
	// Version overrides the pool version resolved from the datastore ref. Only needed when
	// the datastore entry is missing or ambiguous.
	Version *semver.Version `yaml:"version,omitempty" json:"version,omitempty"`
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
	// RateLimits lists the rate limit buckets to set on this lane (at most one per
	// fastFinality flag). Values are human-readable token amounts; inbound gets +10%.
	RateLimits []RateLimitBucketInput `yaml:"rateLimits,omitempty" json:"rateLimits,omitempty"`
}

// RateLimitBucketInput is one outbound/inbound rate limit pair for a finality bucket.
type RateLimitBucketInput struct {
	FastFinality bool                        `yaml:"fastFinality" json:"fastFinality"`
	Outbound     RateLimiterConfigFloatInput `yaml:"outbound" json:"outbound"`
	Inbound      RateLimiterConfigFloatInput `yaml:"inbound" json:"inbound"`
}

// rateLimitBuckets converts the bucket list into two RemoteOutbounds (one per direction)
// so the structural validation rules from SetTokenPoolRateLimits can be reused verbatim.
func (r RemoteConfigUpdate) rateLimitBuckets() (outbounds, inbounds RemoteOutbounds) {
	for _, b := range r.RateLimits {
		outbounds.Outbounds = append(outbounds.Outbounds, RateLimitConfig{RateLimit: b.Outbound, FastFinality: b.FastFinality})
		inbounds.Outbounds = append(inbounds.Outbounds, RateLimitConfig{RateLimit: b.Inbound, FastFinality: b.FastFinality})
	}
	return outbounds, inbounds
}

// isEmpty reports whether the pool entry has nothing to apply. Empty entries are rejected in
// verify as probable YAML mistakes (e.g. indentation errors silently dropping fields).
func (p PoolConfigUpdate) isEmpty() bool {
	return p.FinalityConfig == nil && p.RateLimitAdmin == nil && p.FeeAdmin == nil && len(p.Remotes) == 0
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
		registry := GetTokenAdapterRegistry()

		// First pass: purely structural checks (no datastore/on-chain resolution). Obvious
		// input mistakes — bad selectors, empty updates, duplicate entries, malformed rate
		// limits — must surface before any resolution error from a later pool would mask them.
		seenPools := make(map[string]struct{})
		for _, chainCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainCfg.ChainSelector)
			if err != nil {
				return fmt.Errorf("invalid chain selector %d: %w", chainCfg.ChainSelector, err)
			}
			if len(chainCfg.Pools) == 0 {
				return fmt.Errorf("no pools provided for chain selector %d", chainCfg.ChainSelector)
			}
			for _, pool := range chainCfg.Pools {
				if err := validatePoolConfigUpdate(chainCfg.ChainSelector, family, pool); err != nil {
					return err
				}
				key := fmt.Sprintf("%d|%s", chainCfg.ChainSelector, datastore_utils.SprintRef(pool.TokenPoolRef))
				if _, dup := seenPools[key]; dup {
					return fmt.Errorf("duplicate pool entry for chain selector %d and ref %s", chainCfg.ChainSelector, datastore_utils.SprintRef(pool.TokenPoolRef))
				}
				seenPools[key] = struct{}{}
			}
		}

		// Second pass: resolution + version gating. The gate runs BEFORE adapter resolution so a
		// pre-v2 pool yields the scope error even when no pre-v2 adapter is registered.
		for _, chainCfg := range cfg.Chains {
			for _, pool := range chainCfg.Pools {
				fullPoolRef, err := ResolveTokenPoolRef(e, registry, chainCfg.ChainSelector, pool.TokenPoolRef)
				if err != nil {
					return fmt.Errorf("chain selector %d: failed to resolve token pool ref: %w", chainCfg.ChainSelector, err)
				}
				if pool.Version != nil {
					fullPoolRef.Version = pool.Version
				}
				// PR#1 scope: EVM v2.0.0+ pools only. PR#2 adds EVM 1.5.x/1.6.x, PR#3 adds Solana.
				if fullPoolRef.Version != nil && fullPoolRef.Version.LessThan(utils.Version_2_0_0) {
					return fmt.Errorf(
						"pool %s on chain selector %d has version %s: only v2.0.0+ pools are currently supported by this changeset",
						fullPoolRef.Address, chainCfg.ChainSelector, fullPoolRef.Version,
					)
				}
				// A nil version falls through to ResolveAdapter, which reports it clearly.
				if _, _, err := ResolveAdapter(registry, chainCfg.ChainSelector, fullPoolRef.Version); err != nil {
					return fmt.Errorf("chain selector %d: %w", chainCfg.ChainSelector, err)
				}
			}
		}
		return nil
	}
}

// validatePoolConfigUpdate performs structural (non-resolving) validation of one pool entry.
func validatePoolConfigUpdate(chainSelector uint64, family string, pool PoolConfigUpdate) error {
	if pool.isEmpty() {
		return fmt.Errorf("pool entry %s on chain selector %d has no fields to update", datastore_utils.SprintRef(pool.TokenPoolRef), chainSelector)
	}
	if pool.FinalityConfig != nil {
		if err := pool.FinalityConfig.Validate(); err != nil {
			return fmt.Errorf("finality config for pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.TokenPoolRef), chainSelector, err)
		}
	}
	if family == chain_selectors.FamilyEVM {
		if pool.RateLimitAdmin != nil && !common.IsHexAddress(*pool.RateLimitAdmin) {
			return fmt.Errorf("invalid rateLimitAdmin address %q for pool on chain selector %d", *pool.RateLimitAdmin, chainSelector)
		}
		if pool.FeeAdmin != nil && !common.IsHexAddress(*pool.FeeAdmin) {
			return fmt.Errorf("invalid feeAdmin address %q for pool on chain selector %d", *pool.FeeAdmin, chainSelector)
		}
	}
	seenRemotes := make(map[uint64]struct{})
	for _, remote := range pool.Remotes {
		if remote.RemoteChainSelector == chainSelector {
			return fmt.Errorf("remote chain selector %d must not equal the pool's own chain selector", remote.RemoteChainSelector)
		}
		if _, err := chain_selectors.GetSelectorFamily(remote.RemoteChainSelector); err != nil {
			return fmt.Errorf("invalid remote chain selector %d: %w", remote.RemoteChainSelector, err)
		}
		if _, dup := seenRemotes[remote.RemoteChainSelector]; dup {
			return fmt.Errorf("duplicate remote chain selector %d for pool on chain selector %d", remote.RemoteChainSelector, chainSelector)
		}
		seenRemotes[remote.RemoteChainSelector] = struct{}{}
		if remote.TokenTransferFeeConfig == nil && len(remote.RateLimits) == 0 {
			return fmt.Errorf("remote entry %d for pool on chain selector %d has nothing to update", remote.RemoteChainSelector, chainSelector)
		}
		if remote.TokenTransferFeeConfig != nil {
			if v, ok := remote.TokenTransferFeeConfig.DestBytesOverhead.Get(); ok && v < 32 {
				return fmt.Errorf("destBytesOverhead must be at least 32 for remote %d on chain selector %d, got %d", remote.RemoteChainSelector, chainSelector, v)
			}
		}
		outbounds, inbounds := remote.rateLimitBuckets()
		if err := outbounds.Validate(); err != nil {
			return fmt.Errorf("outbound rate limits for remote %d on chain selector %d: %w", remote.RemoteChainSelector, chainSelector, err)
		}
		if err := inbounds.Validate(); err != nil {
			return fmt.Errorf("inbound rate limits for remote %d on chain selector %d: %w", remote.RemoteChainSelector, chainSelector, err)
		}
	}
	return nil
}

// resolvePoolForConfigure resolves the pool ref, applies the optional explicit version
// override, and resolves the adapter and token ref. It mirrors ResolveAdapterAndRefs but
// honors PoolConfigUpdate.Version between pool-ref resolution and adapter lookup.
func resolvePoolForConfigure(
	e cldf.Environment,
	reg *TokenAdapterRegistry,
	sel uint64,
	pool PoolConfigUpdate,
) (TokenAdapter, string, datastore.AddressRef, datastore.AddressRef, error) {
	fullPoolRef, err := ResolveTokenPoolRef(e, reg, sel, pool.TokenPoolRef)
	if err != nil {
		return nil, "", datastore.AddressRef{}, datastore.AddressRef{}, fmt.Errorf("failed to resolve token pool ref: %w", err)
	}
	if pool.Version != nil {
		fullPoolRef.Version = pool.Version
	}
	adapter, family, err := ResolveAdapter(reg, sel, fullPoolRef.Version)
	if err != nil {
		return nil, "", datastore.AddressRef{}, datastore.AddressRef{}, fmt.Errorf("failed to resolve adapter: %w", err)
	}
	derivedTokenAddr, err := adapter.DeriveTokenAddress(e, sel, fullPoolRef)
	if err != nil {
		return nil, "", datastore.AddressRef{}, datastore.AddressRef{}, fmt.Errorf("failed to derive token address from pool %s: %w", fullPoolRef.Address, err)
	}
	fullTokenRef, err := ResolveTokenRef(e, reg, sel, datastore.AddressRef{ChainSelector: sel, Address: derivedTokenAddr})
	if err != nil {
		return nil, "", datastore.AddressRef{}, datastore.AddressRef{}, fmt.Errorf("failed to resolve token ref for chain selector %d: %w", sel, err)
	}
	return adapter, family, fullPoolRef, fullTokenRef, nil
}

func configureTokenPoolApply() func(cldf.Environment, ConfigureTokenPoolInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokenPoolInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		registry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		for _, chainCfg := range cfg.Chains {
			for _, pool := range chainCfg.Pools {
				poolBatchOps, poolReports, err := applyPoolConfigUpdate(e, registry, chainCfg.ChainSelector, pool)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector, err)
				}
				batchOps = append(batchOps, poolBatchOps...)
				reports = append(reports, poolReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

// applyPoolConfigUpdate applies one pool entry: admin/finality updates first, then
// per-remote configs in input order, so proposal contents are deterministic across runs.
// Each step reads current on-chain state and emits no operations when it already matches.
func applyPoolConfigUpdate(
	e cldf.Environment,
	registry *TokenAdapterRegistry,
	selector uint64,
	pool PoolConfigUpdate,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	adapter, family, fullPoolRef, fullTokenRef, err := resolvePoolForConfigure(e, registry, selector, pool)
	if err != nil {
		return nil, nil, err
	}
	_ = family       // used in Task 5 (rate limit scaling)
	_ = fullTokenRef // used in Task 5 (decimals + on-chain reads)

	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	if pool.FinalityConfig != nil {
		feeAdapter, ok := adapter.(TokenFeeAdapter)
		if !ok {
			return nil, nil, fmt.Errorf(
				"adapter for chain selector %d (family %s, version %s) does not support finality config updates",
				selector, family, fullPoolRef.Version,
			)
		}
		// The sequence reads the current on-chain finality config and emits no writes
		// when it already matches the desired value.
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, feeAdapter.SetAllowedFinalityConfig(&e), e.BlockChains, SetAllowedFinalityConfigSequenceInput{
			Selector: selector,
			Settings: map[string]finality.Config{fullPoolRef.Address: *pool.FinalityConfig},
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to set finality config on pool %s: %w", fullPoolRef.Address, err)
		}
		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)
	}

	// Task 3 adds: rate limit admin / fee admin
	// Task 4 adds: per-remote fee configs
	// Task 5 adds: per-remote rate limits

	return batchOps, reports, nil
}
