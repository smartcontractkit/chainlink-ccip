package tokens

import (
	"errors"
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
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
	ChainSelector uint64 `yaml:"selector" json:"selector"`
	// Pools lists the pool configuration updates on this chain.
	Pools []PoolConfigUpdate `yaml:"pools" json:"pools"`
}

// PoolConfigUpdate describes a partial configuration update for a single token pool.
// Every field other than TokenPoolRef is optional: absent fields leave on-chain state
// untouched. To clear an admin value, provide it explicitly (e.g. the zero address).
// RouterRef cannot be cleared: token pools reject a zero router on every family.
type PoolConfigUpdate struct {
	// TokenPoolRef is a reference to the token pool in the datastore. For Solana, this must be
	// the pool's config PDA, not the pool program ID: a Solana pool address is a program ID
	// shared across many mints, so it cannot alone identify the token being configured.
	TokenPoolRef datastore.AddressRef `yaml:"tokenPoolRef" json:"tokenPoolRef"`
	// FinalityConfig, if set, is the allowed finality config to set on the pool (v2+ only).
	FinalityConfig *finality.Config `yaml:"finalityConfig,omitempty" json:"finalityConfig,omitempty"`
	// RateLimitAdmin, if set, is the desired rate limit admin address.
	RateLimitAdmin *string `yaml:"rateLimitAdmin,omitempty" json:"rateLimitAdmin,omitempty"`
	// FeeAdmin, if set, is the desired fee admin address (v2+ only).
	FeeAdmin *string `yaml:"feeAdmin,omitempty" json:"feeAdmin,omitempty"`
	// RouterRef, if set, selects the router the pool is wired to. An explicit non-zero Address
	// bypasses the datastore. Otherwise the ref is resolved in the datastore for this chain,
	// with Type defaulting to the production Router contract type when unset (set Type to the
	// TestRouter contract type to target the test router). Same semantics as
	// TokenExpansionInputPerChain.RouterRef. Like TokenPoolRef, a datastore lookup that finds
	// no match is reported at apply time, not by VerifyPreconditions.
	//
	// EVM only. Solana has a single router program per chain (it doubles as the OnRamp and the
	// token admin registry) that is upgraded in place, so repointing a pool at a different
	// router is not a meaningful operation there; the Solana adapter rejects a non-nil RouterRef.
	RouterRef *datastore.AddressRef `yaml:"routerRef,omitempty" json:"routerRef,omitempty"`
	// Remotes lists per-lane configuration updates.
	Remotes []RemoteConfigUpdate `yaml:"remotes,omitempty" json:"remotes,omitempty"`
}

// RemoteConfigUpdate describes partial per-lane configuration for one remote chain.
type RemoteConfigUpdate struct {
	// RemoteChainSelector identifies the remote chain of the lane.
	RemoteChainSelector uint64 `yaml:"selector" json:"selector"`
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
	return func(_ cldf.Environment, cfg ConfigureTokenPoolInput) error {
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
				if pool.FinalityConfig == nil && pool.RateLimitAdmin == nil && pool.FeeAdmin == nil && pool.RouterRef == nil && len(pool.Remotes) == 0 {
					return fmt.Errorf("pool entry %s on chain selector %d has no fields to update", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector)
				}
				if pool.RouterRef != nil {
					if datastore_utils.IsAddressRefEmpty(*pool.RouterRef) {
						return fmt.Errorf("pool entry %s on chain selector %d has an empty routerRef", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector)
					}
					if pool.RouterRef.ChainSelector != 0 && pool.RouterRef.ChainSelector != chainCfg.ChainSelector {
						return fmt.Errorf("pool entry %s has routerRef.chainSelector %d that does not match the enclosing chain selector %d", datastore_utils.SprintRef(pool.TokenPoolRef), pool.RouterRef.ChainSelector, chainCfg.ChainSelector)
					}
					if pool.RouterRef.Address != "" {
						if _, err := normalizeRouterAddress(chainCfg.ChainSelector, pool.RouterRef.Address); err != nil {
							return fmt.Errorf("pool entry %s on chain selector %d has an invalid routerRef: %w", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector, err)
						}
					}
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
					if remote.TokenTransferFeeConfig != nil {
						if remote.TokenTransferFeeConfig.IsEmpty() {
							return fmt.Errorf("remote entry %d for pool on chain selector %d has nothing to update", remote.RemoteChainSelector, chainCfg.ChainSelector)
						}
						if !remote.TokenTransferFeeConfig.IsEnabled.IsPresent() {
							return fmt.Errorf("remote entry %d for pool on chain selector %d must specify isEnabled", remote.RemoteChainSelector, chainCfg.ChainSelector)
						}
					}
				}
				normalizedRef, err := deploy.TryNormalizeAddressRef(chainCfg.ChainSelector, pool.TokenPoolRef)
				if err != nil {
					return fmt.Errorf("pool entry %s on chain selector %d has an unnormalizable ref: %w", datastore_utils.SprintRef(pool.TokenPoolRef), chainCfg.ChainSelector, err)
				}
				key := chainConfigKey{poolRef: normalizedRef.Key(), selector: chainCfg.ChainSelector}
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

				if pool.RouterRef != nil || pool.RateLimitAdmin != nil || pool.FeeAdmin != nil {
					adminAdapter, ok := adapter.(TokenPoolAdminAdapter)
					if !ok {
						return cldf.ChangesetOutput{}, fmt.Errorf(
							"adapter for chain selector %d (family %s, version %s) does not support router or admin role updates",
							selector, family, fullPoolRef.Version,
						)
					}
					var router *string
					if pool.RouterRef != nil {
						resolved, err := resolveRouterRef(e.DataStore, selector, *pool.RouterRef)
						if err != nil {
							return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve router for pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.TokenPoolRef), selector, err)
						}
						router = &resolved
					}
					report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adminAdapter.SetTokenPoolAdmins(), e.BlockChains, SetTokenPoolAdminsSequenceInput{
						Selector:       selector,
						Router:         router,
						RateLimitAdmin: pool.RateLimitAdmin,
						FeeAdmin:       pool.FeeAdmin,
						TokenPoolRef:   fullPoolRef,
						TokenRef:       fullTokenRef,
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

// routerContractType is the datastore Type used for the production router on every family.
const routerContractType = datastore.ContractType("Router")

// resolveRouterRef resolves a RouterRef to a normalized, non-zero, family-specific address
// string for the given chain. Semantics match EVMTokenBase.ResolveRouterAddress (used by the
// deploy-token-pool sequences under the same `routerRef` YAML field):
//   - an explicit Address is normalized and used directly, with no datastore lookup;
//   - otherwise the ref is looked up in the datastore for this chain, with Type defaulting to
//     the production Router contract type when unset.
func resolveRouterRef(ds datastore.DataStore, selector uint64, routerRef datastore.AddressRef) (string, error) {
	if routerRef.Address != "" {
		return normalizeRouterAddress(selector, routerRef.Address)
	}

	filter := routerRef.Clone()
	if filter.Type == "" {
		filter.Type = routerContractType
	}
	found, err := datastore_utils.FindAndFormatRef(ds, filter, selector, datastore_utils.FullRef)
	if err != nil {
		return "", fmt.Errorf("failed to resolve router ref %s on chain selector %d: %w", datastore_utils.SprintRef(filter), selector, err)
	}
	return normalizeRouterAddress(selector, found.Address)
}

// normalizeRouterAddress canonicalizes a router address (explicit or datastore-resolved) for the
// given chain using the family AddressNormalizer and rejects the zero address. Token pools on
// every supported family reject a zero router on-chain, so this fails fast instead of producing
// a doomed transaction or proposal.
func normalizeRouterAddress(selector uint64, address string) (string, error) {
	if address == "" {
		return "", fmt.Errorf("router address on chain selector %d must not be empty", selector)
	}
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return "", fmt.Errorf("invalid chain selector %d: %w", selector, err)
	}
	normalizer, ok := deploy.GetAddressNormalizerRegistry().GetAddressNormalizer(family)
	if !ok {
		// No normalizer registered for this family; the adapter is responsible for validation.
		return address, nil
	}
	normalized, err := normalizer.NormalizeAddress(address)
	if err != nil {
		return "", fmt.Errorf("router address %s on chain selector %d is unnormalizable: %w", address, selector, err)
	}
	raw, err := normalizer.StringToBytes(normalized)
	if err != nil {
		return "", fmt.Errorf("router address %s on chain selector %d is unnormalizable: %w", address, selector, err)
	}
	allZero := true
	for _, b := range raw {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return "", fmt.Errorf("router address for chain selector %d must not be zero", selector)
	}
	return normalized, nil
}
