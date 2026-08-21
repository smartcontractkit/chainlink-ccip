package tokens

import (
	"bytes"
	"fmt"
	"maps"
	"slices"

	"github.com/ethereum/go-ethereum/common"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// TokenTransferConfig specifies configuration for a token on one chain to enable transfers with other chains.
type TokenTransferConfig struct {
	// ChainSelector identifies the chain on which the token lives.
	ChainSelector uint64 `yaml:"chainSelector,string" json:"chainSelector,string"`
	// TokenPoolRef is a reference to the token pool in the datastore.
	// Populate the reference as needed to match the desired token pool.
	TokenPoolRef datastore.AddressRef `yaml:"tokenPoolRef" json:"tokenPoolRef"`
	// TokenRef is a reference to the token in the datastore. This is only needed if the token address cannot be derived from the pool reference.
	TokenRef datastore.AddressRef `yaml:"tokenRef" json:"tokenRef"`
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	// Leave empty to use internal administration.
	ExternalAdmin string `yaml:"externalAdmin" json:"externalAdmin"`
	// RegistryRef is a reference to the contract on which the token pool must be registered.
	// Populate the reference as needed to match the desired registry.
	RegistryRef datastore.AddressRef `yaml:"registryRef" json:"registryRef"`
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef] `yaml:"remoteChains" json:"remoteChains"`
	// AllowedFinalityConfig specifies the finality config to set on the token pool. If this is
	// the zero value, then the finality config will remain unchanged on-chain. Pre-v2 pools will
	// ignore this parameter as it is not supported on those versions.
	AllowedFinalityConfig finality.Config `yaml:"allowedFinalityConfig" json:"allowedFinalityConfig"`
	// AutoMigrateRemoteChains is only applicable when migrating a pre-V2 pool to V2. When true, the changeset
	// fetches the currently active pool from TAR, queries its supported remote chains, and populates RemoteChains
	// automatically with (token, pool, decimals, rate limits, and MigrationMetadata). Legacy lane fees are read
	// from the fee quoter (v1.6+) or onramp (v1.5) and merged with any user-provided tokenTransferFeeConfig on each
	// remote (set YAML fields win; unset fields are imported). Resolved connectivity, rate limits, and migration
	// metadata are passed to ConfigureTokenForTransfersSequence for on-chain apply. Forward migration runs only
	// for a genuine V2 upgrade (configured target pool is v2.0.0+) and requires the legacy (pre-V2 active pool)
	// adapter to implement the TokenPoolMigrator interface; discovery also requires the RateLimitReaderAdapter
	// interface. This knob has no effect if any of the following are true:
	//  (1) There is no active pool in TAR for the token
	//  (2) The active pool in TAR is already the target pool (extend mode)
	//  (3) The active pool in TAR is already v2.0.0 or higher
	//  (4) The configured target pool is not v2.0.0+ (i.e. there is no V2 pool to migrate to), so the knob is a no-op for that chain regardless of its adapter's interfaces
	//
	// When discovery is skipped, the changeset logs at info level and does not error. Remote chains,
	// connectivity, and fees are taken only from explicit RemoteChains YAML (same as autoMigrateRemoteChains: false).
	// Remove the flag after a one-time upgrade to avoid confusion.
	//
	// YAML precedence during upgrade (per remote chain):
	//  - Remote not listed: fully discovered from the active pool (token, pool, decimals, rate limits); fees are
	//    imported only when the legacy FeeQuoter/onRamp lane config is enabled for that token/lane.
	//  - Remote listed with empty remoteToken AND empty remotePool: backfill token, pool, and decimals
	//    from the active pool; YAML overrides fees and other fields when tokenTransferFeeConfig is set
	//    (isEnabled required).
	//  - Remote listed with only one of remoteToken or remotePool set: no connectivity backfill — the
	//    provided field is kept as-is and the missing field is not imported from the legacy pool.
	//    Provide both refs explicitly or leave both empty for full backfill.
	//  - Remote listed with both remoteToken and remotePool set: YAML wins (coordinated retarget);
	//    legacy active-pool refs are not overwritten.
	//  - Remote listed but not supported by the legacy active pool: not enriched by discovery; you must
	//    provide full connectivity in YAML.
	//
	// Fee discovery requires connected CCIP lanes (OnRamp/FeeQuoter resolvable per remote). Discovery failures
	// abort the entire changeset. If legacy lane fees are disabled and YAML omits tokenTransferFeeConfig,
	// no fee transactions are emitted on the new v2 pool.
	//
	// Connectivity is propagated in both directions across the existing web of pools. Given a fully connected
	// web A, B, C and migrating A to A_new:
	//  - Forward propagation: A_new's RemoteChains are populated with all remotes discovered from the
	//    active pool (A's legacy remotes B and C), so A_new is configured to reach B and C.
	//  - Reverse propagation: B and C are each told to add A_new as an additional remote pool, so the
	//    web stays reachable from B/C back to A_new. Same-batch peers that are being REPLACED in this
	//    changeset (their target pool differs from their active pool) are skipped for reverse
	//    propagation; forward config wires them instead. Same-batch peers that are not being replaced
	//    (e.g. non-V2 chains like Solana) are still reverse-propagated into so the web stays connected.
	//
	// Limitation: discovery calls getSupportedChains on the TAR-registered active pool. Pools that do not
	// implement that interface (e.g. USDCTokenPoolProxy) cause auto-migrate to fail; list remote chains
	// explicitly in that case.
	AutoMigrateRemoteChains bool `yaml:"autoMigrateRemoteChains" json:"autoMigrateRemoteChains"`
}

// ConfigureTokensForTransfersConfig is the configuration for the ConfigureTokensForTransfers changeset.
type ConfigureTokensForTransfersConfig struct {
	// Tokens specifies the tokens to configure for cross-chain transfers.
	Tokens []TokenTransferConfig
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// ConfigureTokensForTransfers returns a changeset that configures tokens on multiple chains for transfers with other chains.
func ConfigureTokensForTransfers(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureTokensForTransfersConfig] {
	return cldf.CreateChangeSet(makeApply(tokenRegistry, mcmsRegistry), makeVerify(tokenRegistry, mcmsRegistry))
}

func makeVerify(_ *TokenAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) error {
	return func(_ cldf.Environment, cfg ConfigureTokensForTransfersConfig) error {
		return nil
	}
}

func makeApply(_ *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
		configs := make(map[uint64]TokenTransferConfig, len(cfg.Tokens))
		for _, config := range cfg.Tokens {
			configs[config.ChainSelector] = config
		}
		batchOps, reports, ds, err := processTokenConfigForChain(e, configs)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to process token configs for chains: %w", err)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}

func processTokenConfigForChain(e cldf.Environment, cfg map[uint64]TokenTransferConfig) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], *datastore.MemoryDataStore, error) {
	normalizerRegistry := deploy.GetAddressNormalizerRegistry()
	tokenRegistry := GetTokenAdapterRegistry()
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	ds := datastore.NewMemoryDataStore()

	// Process chains in deterministic (sorted) selector order (Go map iteration is randomized)
	var err error
	for _, selector := range slices.Sorted(maps.Keys(cfg)) {
		token := cfg[selector]

		token.RegistryRef, err = deploy.TryNormalizeAddressRef(selector, token.RegistryRef)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to normalize registry ref address for chain selector %d: %w", selector, err)
		}
		cfg[selector] = token

		var registryAddr string
		if datastore_utils.IsAddressRefEmpty(token.RegistryRef) {
			e.Logger.Warnf("Registry ref is empty for chain selector %d. We will rely on the underlying adapter to resolve this field.", selector)
		} else {
			if registry, err := datastore_utils.FindAndFormatRef(e.DataStore, token.RegistryRef, selector, datastore_utils.FullRef); err != nil {
				return nil, nil, nil, fmt.Errorf("failed to resolve registry ref on chain with selector %d: %w", selector, err)
			} else {
				registryAddr = registry.Address
			}
		}

		adapter, family, tokenPool, fullTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, selector, token.TokenPoolRef, token.TokenRef)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to resolve adapter and refs for chain selector %d: %w", selector, err)
		}

		remoteChains := make(map[uint64]RemoteChainConfig[[]byte, string], len(token.RemoteChains))
		for remoteChainSelector, inCfg := range token.RemoteChains {
			counterpart, ok := cfg[remoteChainSelector]
			if !ok {
				return nil, nil, nil, fmt.Errorf("missing token transfer config for remote chain selector %d", remoteChainSelector)
			}
			counterpartRemoteChainCfg, ok := counterpart.RemoteChains[selector]
			if !ok {
				return nil, nil, nil, fmt.Errorf("missing remote chain config for chain selector %d in token transfer config for remote chain selector %d", selector, remoteChainSelector)
			}
			remoteChains[remoteChainSelector], err = convertRemoteChainConfig(
				e,
				selector,
				tokenRegistry,
				remoteChainSelector,
				inCfg,
				counterpartRemoteChainCfg,
			)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to process remote chain config for remote chain selector %d: %w", remoteChainSelector, err)
			}
		}

		type DiscoveredRemoteChain struct {
			remoteSelector uint64
			remoteTokenRef datastore.AddressRef
			remotePoolRef  datastore.AddressRef
			remoteAdapter  TokenAdapter
		}

		// When AutoMigrateRemoteChains = true, we read the legacy pool's remote chains and import them into
		// the new pool. Example: if pools A, B, and C form a fully connected web and we migrate B to B_new,
		// then the code below will discover that remote pools A and C need to be added to B_new to maintain
		// connectivity from B_new to A and C. The reverse propagation is handled later in the code.
		var discoveredRemotes []DiscoveredRemoteChain
		if token.AutoMigrateRemoteChains {
			tarReader, ok := tokenRegistry.GetTokenAdminRegistryReader(family)
			if !ok {
				return nil, nil, nil, fmt.Errorf("no token admin registry reader for chain family %s", family)
			}
			activePool, err := tarReader.GetActivePool(e, selector, fullTokenRef, token.RegistryRef)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to get active pool for token pool on chain selector %d: %w", selector, err)
			}
			var (
				legacyRateLimitReader RateLimitReaderAdapter
				legacyPoolMigrator    TokenPoolMigrator
				allRemoteSelectors    []uint64
				activePoolRef         datastore.AddressRef
				localDecimals         uint8
			)
			if len(activePool) > 0 {
				targetPoolBytes, err := adapter.AddressRefToBytes(tokenPool)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to convert target pool ref to bytes on chain selector %d: %w", selector, err)
				}
				if !bytes.Equal(activePool, targetPoolBytes) {
					localNormalizer, ok := normalizerRegistry.GetAddressNormalizer(family)
					if !ok {
						return nil, nil, nil, fmt.Errorf("no address normalizer found for chain family %s on chain selector %d", family, selector)
					}
					activePoolAddr, err := localNormalizer.BytesToString(activePool)
					if err != nil {
						return nil, nil, nil, fmt.Errorf("failed to normalize active pool address on chain selector %d: %w", selector, err)
					}
					activePoolRef, err = ResolveTokenPoolRef(e, tokenRegistry, selector, datastore.AddressRef{Address: activePoolAddr})
					if err != nil {
						return nil, nil, nil, fmt.Errorf("failed to resolve active pool ref on chain selector %d: %w", selector, err)
					}
					if activePoolRef.Version == nil {
						return nil, nil, nil, fmt.Errorf("active pool version is required for auto-migrate on chain selector %d", selector)
					}
					if activePoolRef.Version.LessThan(utils.Version_2_0_0) {
						legacyAdapter, _, err := ResolveAdapter(tokenRegistry, selector, activePoolRef.Version)
						if err != nil {
							return nil, nil, nil, fmt.Errorf("failed to resolve adapter for active pool on chain selector %d: %w", selector, err)
						}
						// Auto-migration is only meaningful for a genuine V2 upgrade, so require the configured
						// target pool to be v2.0.0+ - keying this on the target's version (rather than on which
						// interfaces the active-pool adapter happens to implement) keeps the gate robust across
						// chain families including adapters which implement TokenPoolMigrator while their chain
						// is still not on V2. The active-pool adapter must also implement the TokenPoolMigrator
						// interface for the discovery body to run; if either condition fails, auto-migrate is a
						// no-op for that chain.
						if tokenPool.Version != nil && tokenPool.Version.GreaterThanEqual(utils.Version_2_0_0) {
							if legacyPoolMigrator, ok = legacyAdapter.(TokenPoolMigrator); ok {
								if legacyRateLimitReader, ok = legacyAdapter.(RateLimitReaderAdapter); !ok {
									return nil, nil, nil, fmt.Errorf(
										"adapter for active pool version %s on chain selector %d does not implement RateLimitReaderAdapter",
										activePoolRef.Version, selector,
									)
								}
								tokenBytes, err := legacyAdapter.AddressRefToBytes(fullTokenRef)
								if err != nil {
									return nil, nil, nil, fmt.Errorf("failed to convert token ref to bytes on chain selector %d: %w", selector, err)
								}
								localDecimals, err = legacyAdapter.DeriveTokenDecimals(e, selector, activePoolRef, tokenBytes)
								if err != nil {
									return nil, nil, nil, fmt.Errorf("failed to derive local token decimals on chain selector %d: %w", selector, err)
								}
								if supported, err := legacyPoolMigrator.GetSupportedChains(e, selector, activePool); err != nil {
									return nil, nil, nil, fmt.Errorf("failed to get supported remote chains for token pool on chain selector %d: %w", selector, err)
								} else {
									allRemoteSelectors = supported
								}
							} else {
								e.Logger.Infof("adapter for active pool version %s on chain selector %d does not support token pool migration, skipping auto-migration of remote chains", activePoolRef.Version, selector)
							}
						} else {
							e.Logger.Infof("Active pool version %s on chain selector %d has no v2.0.0+ migration target, skipping auto-migration of remote chains", activePoolRef.Version, selector)
						}
					} else {
						e.Logger.Infof("Active pool on chain selector %d is already v2.0.0 or higher, skipping auto-migration of remote chains", selector)
					}
				} else {
					e.Logger.Infof("Active pool on chain selector %d is already the target pool, skipping auto-migration of remote chains", selector)
				}
			}
			for _, remoteSelector := range allRemoteSelectors {
				remoteFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
				}
				remoteNormalizer, ok := normalizerRegistry.GetAddressNormalizer(remoteFamily)
				if !ok {
					return nil, nil, nil, fmt.Errorf("no address normalizer found for chain family %s of remote chain selector %d", remoteFamily, remoteSelector)
				}
				remoteTokenBytes, err := legacyPoolMigrator.GetRemoteToken(e, selector, activePool, remoteSelector)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to get remote token for remote chain selector %d: %w", remoteSelector, err)
				}
				remoteTokenAddr, err := remoteNormalizer.BytesToString(remoteTokenBytes)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to normalize remote token address for remote chain selector %d: %w", remoteSelector, err)
				}
				var remotePoolBytes []byte
				if counterpartCfg, alsoMigrating := cfg[remoteSelector]; alsoMigrating {
					fullRemotePoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, remoteSelector, counterpartCfg.TokenPoolRef)
					if err != nil {
						return nil, nil, nil, fmt.Errorf("failed to resolve counterpart pool ref for remote chain selector %d: %w", remoteSelector, err)
					}
					remotePoolBytes, err = remoteNormalizer.StringToBytes(fullRemotePoolRef.Address)
					if err != nil {
						return nil, nil, nil, fmt.Errorf("failed to convert counterpart pool ref to bytes for chain selector %d: %w", remoteSelector, err)
					}
				} else {
					remoteRegReader, ok := tokenRegistry.GetTokenAdminRegistryReader(remoteFamily)
					if !ok {
						return nil, nil, nil, fmt.Errorf("no admin registry reader for remote chain family %s", remoteFamily)
					}
					remotePoolBytes, err = remoteRegReader.GetActivePool(e, remoteSelector, datastore.AddressRef{Address: remoteTokenAddr})
					if err != nil {
						return nil, nil, nil, fmt.Errorf("failed to get active pool for remote chain selector %d: %w", remoteSelector, err)
					}
				}
				remotePoolAddr, err := remoteNormalizer.BytesToString(remotePoolBytes)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to normalize remote pool address for remote chain selector %d: %w", remoteSelector, err)
				}
				remoteAdapter, _, remotePoolRef, remoteTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, remoteSelector, datastore.AddressRef{Address: remotePoolAddr}, datastore.AddressRef{Address: remoteTokenAddr})
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to resolve adapter and refs for remote chain selector %d: %w", remoteSelector, err)
				}
				remotePools, err := legacyPoolMigrator.GetRemotePools(e, selector, activePool, remoteSelector)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to get remote pools for remote chain selector %d: %w", remoteSelector, err)
				}
				discoveredRemotes = append(discoveredRemotes, DiscoveredRemoteChain{
					remoteSelector: remoteSelector,
					remoteTokenRef: remoteTokenRef,
					remotePoolRef:  remotePoolRef,
					remoteAdapter:  remoteAdapter,
				})
				remoteTokenDecimals, err := remoteAdapter.DeriveTokenDecimals(e, remoteSelector, remotePoolRef, remoteTokenBytes)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to derive remote token decimals for remote chain selector %d: %w", remoteSelector, err)
				}
				var rc RemoteChainConfig[[]byte, string]
				if rc, ok = remoteChains[remoteSelector]; ok {
					if len(rc.RemoteToken) == 0 && len(rc.RemotePool) == 0 {
						rc.RemoteDecimals = remoteTokenDecimals
						rc.RemoteToken = remoteTokenBytes
						rc.RemotePool = remotePoolBytes
					}
				} else {
					rc = RemoteChainConfig[[]byte, string]{
						RemoteDecimals: remoteTokenDecimals,
						RemoteToken:    remoteTokenBytes,
						RemotePool:     remotePoolBytes,
					}
				}
				legacyRL, err := LegacyRateLimitsForAutoMigrate(e, legacyRateLimitReader, selector, remoteSelector, activePoolRef, fullTokenRef, localDecimals, rc)
				if err != nil {
					return nil, nil, nil, fmt.Errorf(
						"failed to resolve auto-migrated rate limits for chain selector %d and remote chain selector %d: %w",
						selector, remoteSelector, err,
					)
				}
				legacyFC, err := LegacyFeesForAutoMigrate(e, selector, remoteSelector, fullTokenRef.Address, rc)
				if err != nil {
					return nil, nil, nil, fmt.Errorf(
						"failed to resolve auto-migrated fee config for chain selector %d and remote chain selector %d: %w",
						selector, remoteSelector, err,
					)
				}
				rc.MigrationMetadata = MigrationMetadata{
					LegacyPoolVersion:            activePoolRef.Version,
					LegacyPoolType:               activePoolRef.Type.String(),
					LegacyRemotePools:            remotePools,
					LegacyTokenTransferFeeConfig: legacyFC,
					LegacyRateLimits:             legacyRL,
				}
				remoteChains[remoteSelector] = rc
			}
		}

		// NOTE: this changeset is already responsible for applying the fee configs
		// so we set `TokenTransferFeeConfig` to nil BEFORE calling the lower-level
		// ConfigureTokenForTransfersSequence - this prevents any type of duplicate
		// work from being done at a lower-level.
		remoteChainsWithoutFeeConfigs := make(map[uint64]RemoteChainConfig[[]byte, string], len(remoteChains))
		for sel, rc := range remoteChains {
			rc.TokenTransferFeeConfig = nil
			remoteChainsWithoutFeeConfigs[sel] = rc
		}

		// Configure pool remotes (fee configs are excluded as they require
		// special handling see comment below)
		configureTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureTokenForTransfersSequence(), e.BlockChains, ConfigureTokenForTransfersInput{
			ChainSelector:         selector,
			TokenPoolAddress:      tokenPool.Address,
			RemoteChains:          remoteChainsWithoutFeeConfigs,
			ExternalAdmin:         token.ExternalAdmin,
			RegistryAddress:       registryAddr,
			TokenRef:              fullTokenRef,
			PoolType:              tokenPool.Type.String(),
			ExistingDataStore:     e.DataStore,
			AllowedFinalityConfig: token.AllowedFinalityConfig,
		})
		if err != nil {
			return batchOps, reports, nil, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", selector, err)
		}
		batchOps = append(batchOps, configureTokenReport.Output.BatchOps...)
		reports = append(reports, configureTokenReport.ExecutionReports...)
		for _, r := range configureTokenReport.Output.Addresses {
			if err := ds.Addresses().Add(r); err != nil {
				return nil, nil, nil, fmt.Errorf("failed to add address %s to datastore: %w", r.Address, err)
			}
		}

		// Fee configs require special handling - determining the correct
		// fee contract to configure is a multi-step process that usually
		// involves invoking the bindings for multiple contract versions.
		// Technically speaking this could be done inside the lower level
		// sequence, but it defeats the whole purpose of having versioned
		// adapters (i.e. the v2 adapter would need to know how to do pre
		// v2 operations). To avoid this anti-pattern, it would be better
		// to do this type of resolution here, in the top level changeset
		// as it's already responsible for multi-version orchestration.
		for remoteSelector, inCfg := range remoteChains {
			var feeCfg *PartialTokenTransferFeeConfig
			if inCfg.MigrationMetadata.IsPopulated() {
				feeCfg = inCfg.MigrationMetadata.LegacyTokenTransferFeeConfig
			} else {
				feeCfg = inCfg.TokenTransferFeeConfig
			}
			if feeCfg != nil {
				feeBatchOps, feeReports, err := applyTokenTransferFeeConfig(
					e,
					selector,
					remoteSelector,
					tokenPool,
					fullTokenRef,
					*feeCfg,
				)
				if err != nil {
					return nil, nil, nil, fmt.Errorf("failed to apply token transfer fee config for remote chain selector %d: %w", remoteSelector, err)
				}
				batchOps = append(batchOps, feeBatchOps...)
				reports = append(reports, feeReports...)
			}
		}

		// Reverse-propagate the new pool address to every counterpart discovered by autoMigrateRemoteChains.
		// Example: if pools A, B, and C form a fully connected web and we migrate B to B_new, then this step
		// tells A and C to add B_new as an additional remote pool, so that web remains fully connected after
		// the migration is performed. Without this, the forward direction (A → B_new) works, but the reverse
		// (B_new → A) would fail because A's pool still only knows about B_old.
		//
		// This loop doesn't need to import rate limits, fee configs, or any other per-chain configs onto the
		// counterpart pools (A, C). If a counterpart is itself migrated later then `autoMigrateRemoteChains`
		// on that upgrade will discover B_new as the active pool on chain B and handle those imports at that
		// time through the normal forward flow.
		if len(discoveredRemotes) > 0 {
			migratedTokenBytes, err := adapter.AddressRefToBytes(fullTokenRef)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to convert token ref to bytes for reverse propagation on chain selector %d: %w", selector, err)
			}

			migratedPoolBytes, err := adapter.AddressRefToBytes(tokenPool)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to convert new pool ref to bytes for reverse propagation on chain selector %d: %w", selector, err)
			}
			for _, ru := range discoveredRemotes {
				// After a pool migrates reverse-propagation tells every peer it was connected to about the new
				// pool, keeping the web reachable in both directions. We skip one case: a same-batch peer that
				// is itself being REPLACED — its new pool isn't live yet, so configuring it now would activate
				// it too early and break that peer's own migration. Its own forward config wires it instead.
				//
				// A peer is being replaced exactly when its configured TARGET pool is V2 (>= 2.0.0): such a
				// peer is skipped. Any peer with a pre-V2 target is not being replaced and must still learn
				// about the new pool. Example: in one batch say we migrate A and C to A2/C2 while B remains
				// on its v1 pool. Processing A tells B about A2 but leaves C alone (C is getting C2 in this
				// same batch); processing C tells B about C2 and leaves A alone. So B learns both A2 and C2,
				// C2 pairs with A2 via C's own forward config, A2 pairse with C2 via A's own forward config,
				// and the web stays fully connected.
				//
				// Deciding by the target version (not by which interfaces the peer's adapter implements) stays
				// correct even if a non-V2 family later adds a migrator.
				if _, alsoMigrating := cfg[ru.remoteSelector]; alsoMigrating {
					if ru.remotePoolRef.Version == nil {
						return nil, nil, nil, fmt.Errorf("remote pool version is required for reverse propagation to counterpart chain selector %d for chain selector %d", ru.remoteSelector, selector)
					}
					if ru.remotePoolRef.Version.GreaterThanEqual(utils.Version_2_0_0) {
						continue
					}
				}
				reverseInput := ConfigureTokenForTransfersInput{
					// Reverse propagation only *ADDS* this migration's new pool as an additional remote to an
					// existing active pool including pools that may already be migrated and currently support
					// many chains. The upgrade-safety check that requires remoteChains to cover all supported
					// chains is therefore not applicable here — we are extending, not replacing the pool.
					SkipActivePoolSupportedChainsCheck: true,
					ExistingDataStore:                  e.DataStore,
					TokenPoolAddress:                   ru.remotePoolRef.Address,
					ChainSelector:                      ru.remoteSelector,
					TokenRef:                           ru.remoteTokenRef,
					PoolType:                           ru.remotePoolRef.Type.String(),
					RemoteChains: map[uint64]RemoteChainConfig[[]byte, string]{
						selector: {
							// Pad to match on-chain storage format (32-byte left-padded). The address
							// comparisons in `ConfigureTokenPoolForRemoteChain()` compare against on-
							// chain bytes without padding, so a 20-byte input would mismatch and fire
							// a remove+re-add instead of addRemotePool
							RemoteToken: common.LeftPadBytes(migratedTokenBytes, 32),
							RemotePool:  common.LeftPadBytes(migratedPoolBytes, 32),
						},
					},
				}
				reverseReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ru.remoteAdapter.ConfigureTokenForTransfersSequence(), e.BlockChains, reverseInput)
				if err != nil {
					return nil, nil, nil, fmt.Errorf(
						"failed to propagate new pool to counterpart chain selector %d for chain selector %d: %w",
						ru.remoteSelector, selector, err,
					)
				}
				batchOps = append(batchOps, reverseReport.Output.BatchOps...)
				reports = append(reports, reverseReport.ExecutionReports...)
				for _, r := range reverseReport.Output.Addresses {
					if err := ds.Addresses().Add(r); err != nil {
						return nil, nil, nil, fmt.Errorf("failed to add address %s to datastore: %w", r.Address, err)
					}
				}
			}
		}
	}

	return batchOps, reports, ds, nil
}

func applyTokenTransferFeeConfig(
	e cldf.Environment,
	src, dst uint64,
	fullSrcPoolRef datastore.AddressRef,
	fullSrcTokenRef datastore.AddressRef,
	srcToDstFeeCfg PartialTokenTransferFeeConfig,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	if fullSrcPoolRef.Version == nil {
		return nil, nil, fmt.Errorf("token pool version is required to apply token transfer fee config for chain selector %d and remote chain selector %d", src, dst)
	}

	// NOTE: fee configs are applied differently based on the pool version:
	//   Pre-V2 pools: apply the fee config on the fee quoter / onRamp (legacy lane).
	//   On V2+ pools: apply the fee config on the token pool via TokenFeeAdapter.
	if fullSrcPoolRef.Version.LessThan(utils.Version_2_0_0) {
		return applyTokenTransferFeeConfigOnFeeQuoter(e, src, dst, fullSrcTokenRef, srcToDstFeeCfg)
	} else {
		return applyTokenTransferFeeConfigOnTokenPool(e, src, dst, fullSrcPoolRef, srcToDstFeeCfg)
	}
}

func applyTokenTransferFeeConfigOnTokenPool(
	e cldf.Environment,
	src, dst uint64,
	fullSrcPoolRef datastore.AddressRef,
	partial PartialTokenTransferFeeConfig,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	feeAdapter, err := ResolveTokenFeeAdapter(e, src, fullSrcPoolRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to resolve token fee adapter for chain selector %d and token pool address %s: %w", src, fullSrcPoolRef.Address, err)
	}
	poolAddress := fullSrcPoolRef.Address
	if poolAddress == "" {
		return nil, nil, fmt.Errorf("token pool address is required to apply token transfer fee config for chain selector %d and remote chain selector %d", src, dst)
	}

	onChainConfig, err := feeAdapter.GetOnchainTokenTransferFeeConfig(e, poolAddress, src, dst)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get on-chain token transfer fee config for pool %s on chain selector %d and remote chain selector %d: %w", poolAddress, src, dst, err)
	}
	defaultConfig := GetDefaultChainAgnosticTokenTransferFeeConfig(
		src,
		dst,
	)

	// Resolution strategy:
	// (1) If on-chain config is enabled, merge it with the user's provided config (giving precedence to user's config)
	// (2) Fall back to sensible defaults merged with user's provided config (giving precedence to user's config)
	var requestedConfig TokenTransferFeeConfig
	if onChainConfig.IsEnabled {
		requestedConfig = partial.MergeWith(onChainConfig)
	} else {
		requestedConfig = partial.MergeWith(defaultConfig)
	}

	if !requestedConfig.IsEnabled && !onChainConfig.IsEnabled {
		e.Logger.Infof("Skipping token transfer fee config for chain selector %d and remote chain selector %d since pool fee override is already disabled", src, dst)
		return nil, nil, nil
	}

	if requestedConfig == onChainConfig {
		e.Logger.Infof("Skipping token transfer fee config for chain selector %d and remote chain selector %d since the desired config is the same as the current on-chain config", src, dst)
		return nil, nil, nil
	}

	result, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle,
		feeAdapter.SetTokenTransferFee(&e),
		e.BlockChains,
		SetTokenTransferFeeSequenceInput{
			Selector: src,
			Settings: map[string]map[uint64]*TokenTransferFeeConfig{
				poolAddress: {
					dst: &requestedConfig,
				},
			},
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute set token transfer fee sequence for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	return result.Output.BatchOps, result.ExecutionReports, nil
}

func applyTokenTransferFeeConfigOnFeeQuoter(
	e cldf.Environment,
	src, dst uint64,
	fullSrcTokenRef datastore.AddressRef,
	partial PartialTokenTransferFeeConfig,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(e.OperationsBundle, e.BlockChains, e.DataStore, src, dst)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to resolve fee adapter for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}
	tokAddress := fullSrcTokenRef.Address
	if tokAddress == "" {
		return nil, nil, fmt.Errorf("source token address is required to apply token transfer fee config for remote chain selector %d", dst)
	}

	// NOTE: the TokenTransferFeeConfig for token pools is V2-focused and
	// does NOT have MaxFeeUSDCents fields. As a result we will reuse the
	// existing value from the chain or fallback to a sensible default if
	// it isn't set on chain. It can't be configured directly by the user
	// at the moment, but realistically speaking this should not an issue
	// since we've never had the need to modify it after we initially set
	// it to MaxUint32.
	onChainConfig, err := feeAdapter.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, fqRef, src, dst, fullSrcTokenRef.Address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get current on-chain token transfer fee config for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}
	defaultConfig := fees.GetDefaultChainAgnosticTokenTransferFeeConfig(
		src,
		dst,
	)

	// Resolution strategy:
	// (1) If on-chain config is enabled, merge it with the user's provided config (giving precedence to user's config)
	// (2) Fall back to sensible defaults merged with user's provided config (giving precedence to user's config)
	var requestedConfig fees.TokenTransferFeeArgs
	if onChainConfig.IsEnabled {
		requestedConfig = fees.TokenTransferFeeArgs{
			MinFeeUSDCents:    partial.DefaultFinalityFeeUSDCents.GetOrDefault(onChainConfig.MinFeeUSDCents),
			DeciBps:           partial.DefaultFinalityTransferFeeBps.GetOrDefault(onChainConfig.DeciBps),
			DestBytesOverhead: partial.DestBytesOverhead.GetOrDefault(onChainConfig.DestBytesOverhead),
			DestGasOverhead:   partial.DestGasOverhead.GetOrDefault(onChainConfig.DestGasOverhead),
			IsEnabled:         partial.IsEnabled.GetOrDefault(onChainConfig.IsEnabled),
			MaxFeeUSDCents:    onChainConfig.MaxFeeUSDCents,
		}
	} else {
		requestedConfig = fees.TokenTransferFeeArgs{
			MinFeeUSDCents:    partial.DefaultFinalityFeeUSDCents.GetOrDefault(defaultConfig.MinFeeUSDCents),
			DeciBps:           partial.DefaultFinalityTransferFeeBps.GetOrDefault(defaultConfig.DeciBps),
			DestBytesOverhead: partial.DestBytesOverhead.GetOrDefault(defaultConfig.DestBytesOverhead),
			DestGasOverhead:   partial.DestGasOverhead.GetOrDefault(defaultConfig.DestGasOverhead),
			IsEnabled:         partial.IsEnabled.GetOrDefault(defaultConfig.IsEnabled),
			MaxFeeUSDCents:    defaultConfig.MaxFeeUSDCents,
		}
	}

	if !requestedConfig.IsEnabled && !onChainConfig.IsEnabled {
		e.Logger.Infof("Skipping token transfer fee config for chain selector %d and remote chain selector %d since legacy lane fee config is already disabled", src, dst)
		return nil, nil, nil
	}

	if requestedConfig == onChainConfig {
		e.Logger.Infof("Skipping token transfer fee config for chain selector %d and remote chain selector %d since the desired config is the same as the current on-chain config", src, dst)
		return nil, nil, nil
	}

	result, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle,
		feeAdapter.SetTokenTransferFee(e.DataStore, fqRef),
		e.BlockChains,
		fees.SetTokenTransferFeeSequenceInput{
			Selector: src,
			Settings: map[uint64]map[string]*fees.TokenTransferFeeArgs{
				dst: {
					fullSrcTokenRef.Address: &requestedConfig,
				},
			},
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute set token transfer fee sequence for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	return result.Output.BatchOps, result.ExecutionReports, nil
}

func convertRemoteChainConfig(
	e cldf.Environment,
	chainSelector uint64,
	tokenAdapterRegistry *TokenAdapterRegistry,
	remoteChainSelector uint64,
	inCfg RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef],
	cpCfg RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef],
) (RemoteChainConfig[[]byte, string], error) {
	if err := inCfg.Validate(); err != nil {
		return RemoteChainConfig[[]byte, string]{}, fmt.Errorf("invalid remote chain config (chain %d → %d): %w", chainSelector, remoteChainSelector, err)
	}
	if err := cpCfg.Validate(); err != nil {
		return RemoteChainConfig[[]byte, string]{}, fmt.Errorf("invalid counterpart remote chain config (chain %d → %d): %w", remoteChainSelector, chainSelector, err)
	}

	var outbound, inbound *RateLimiterConfigFloatInput
	if ob, inOk := inCfg.GetOutboundRateLimitBuckets().DefaultBucket(); inOk {
		outbound = &ob.RateLimit
	}
	if ib, cpOk := cpCfg.GetOutboundRateLimitBuckets().DefaultBucket(); cpOk {
		inbound = &ib.RateLimit
	}

	// a chain's inbound rate limiter config should be based on the remote chain's outbound rate limiter config
	// to ensure that the remote chain is configured to allow the desired traffic from this chain.
	// The values here should NOT be passed in decimal adjusted but rather the adapters should be responsible for performing
	// any necessary decimal adjustments based on the token decimals on each chain.
	outCfg := RemoteChainConfig[[]byte, string]{
		InboundRateLimiterConfig:  inbound,
		OutboundRateLimiterConfig: outbound,
		InboundRateLimits:         cpCfg.OutboundRateLimits,
		OutboundRateLimits:        inCfg.OutboundRateLimits,
		TokenTransferFeeConfig:    inCfg.TokenTransferFeeConfig,
	}

	if inCfg.RemotePool != nil {
		fullRemotePoolRef, err := ResolveTokenPoolRef(e, tokenAdapterRegistry, remoteChainSelector, *inCfg.RemotePool)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote pool ref %s: %w", datastore_utils.SprintRef(*inCfg.RemotePool), err)
		}
		remoteAdapter, _, err := ResolveAdapter(tokenAdapterRegistry, remoteChainSelector, fullRemotePoolRef.Version)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote adapter for remote chain selector %d: %w", remoteChainSelector, err)
		}
		outCfg.RemotePool, err = remoteAdapter.AddressRefToBytes(fullRemotePoolRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to convert remote pool ref %s to bytes: %w", datastore_utils.SprintRef(*inCfg.RemotePool), err)
		}

		// If DeriveTokenAddress succeeds, then this has higher precedence than the token ref provided in the input since it is
		// derived from on chain data (and hence more reliable). If it fails, then we fall back to using the token ref provided
		// in the input and try to resolve it from the datastore first (to avoid RPC calls) then fall back to on chain data.
		derivedTokenAddr, deriveErr := remoteAdapter.DeriveTokenAddress(e, remoteChainSelector, fullRemotePoolRef)
		switch {
		case deriveErr == nil:
			e.Logger.Infof("Successfully derived remote token address %s for remote chain selector %d from remote pool ref %s", derivedTokenAddr, remoteChainSelector, datastore_utils.SprintRef(fullRemotePoolRef))
			resolvedRef, err := ResolveTokenRef(e, tokenAdapterRegistry, remoteChainSelector, datastore.AddressRef{ChainSelector: remoteChainSelector, Address: derivedTokenAddr})
			if err != nil {
				return outCfg, fmt.Errorf("failed to resolve remote token after derivation %s: %w", derivedTokenAddr, err)
			}
			outCfg.RemoteToken, err = remoteAdapter.AddressRefToBytes(resolvedRef)
			if err != nil {
				return outCfg, fmt.Errorf("failed to convert resolved remote token to bytes %s: %w", derivedTokenAddr, err)
			}
		case inCfg.RemoteToken != nil:
			e.Logger.Infof("Derivation of remote token address failed for remote chain selector %d (%s). Falling back to resolving remote token from provided token ref %s", remoteChainSelector, deriveErr.Error(), datastore_utils.SprintRef(*inCfg.RemoteToken))
			resolvedRef, err := ResolveTokenRef(e, tokenAdapterRegistry, remoteChainSelector, *inCfg.RemoteToken)
			if err != nil {
				return outCfg, fmt.Errorf("failed to resolve remote token ref %s: %w", datastore_utils.SprintRef(*inCfg.RemoteToken), err)
			}
			outCfg.RemoteToken, err = remoteAdapter.AddressRefToBytes(resolvedRef)
			if err != nil {
				return outCfg, fmt.Errorf("failed to convert remote token ref %s to bytes: %w", datastore_utils.SprintRef(*inCfg.RemoteToken), err)
			}
		default:
			return outCfg, fmt.Errorf("failed to derive remote token address and no remote token ref provided for remote chain selector %d: %w", remoteChainSelector, deriveErr)
		}

		outCfg.RemoteToken = common.LeftPadBytes(outCfg.RemoteToken, 32)
		outCfg.RemoteDecimals, err = remoteAdapter.DeriveTokenDecimals(e, remoteChainSelector, fullRemotePoolRef, outCfg.RemoteToken)
		if err != nil {
			return outCfg, fmt.Errorf("failed to get remote token decimals for remote chain selector %d: %w", remoteChainSelector, err)
		}
		outCfg.RemotePool, err = remoteAdapter.DeriveTokenPoolCounterpart(e, remoteChainSelector, outCfg.RemotePool, outCfg.RemoteToken)
		if err != nil {
			return outCfg, fmt.Errorf("failed to derive remote pool counterpart for remote chain selector %d: %w", remoteChainSelector, err)
		}
	}
	for _, ccvRef := range inCfg.OutboundCCVs {
		ref, err := deploy.TryNormalizeAddressRef(chainSelector, ccvRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to normalize outbound CCV ref address for chain selector %d: %w", chainSelector, err)
		}
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve outbound CCV ref %s: %w", datastore_utils.SprintRef(ref), err)
		}
		outCfg.OutboundCCVs = append(outCfg.OutboundCCVs, fullCCVRef.Address)
	}
	for _, ccvRef := range inCfg.InboundCCVs {
		ref, err := deploy.TryNormalizeAddressRef(chainSelector, ccvRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to normalize inbound CCV ref address for chain selector %d: %w", chainSelector, err)
		}
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve inbound CCV ref %s: %w", datastore_utils.SprintRef(ref), err)
		}
		outCfg.InboundCCVs = append(outCfg.InboundCCVs, fullCCVRef.Address)
	}
	for _, ccvRef := range inCfg.OutboundCCVsToAddAboveThreshold {
		ref, err := deploy.TryNormalizeAddressRef(chainSelector, ccvRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to normalize outbound CCV-above-threshold ref address for chain selector %d: %w", chainSelector, err)
		}
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve outbound CCV to add above threshold ref %s: %w", datastore_utils.SprintRef(ref), err)
		}
		outCfg.OutboundCCVsToAddAboveThreshold = append(outCfg.OutboundCCVsToAddAboveThreshold, fullCCVRef.Address)
	}
	for _, ccvRef := range inCfg.InboundCCVsToAddAboveThreshold {
		ref, err := deploy.TryNormalizeAddressRef(chainSelector, ccvRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to normalize inbound CCV-above-threshold ref address for chain selector %d: %w", chainSelector, err)
		}
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve inbound CCV to add above threshold ref %s: %w", datastore_utils.SprintRef(ref), err)
		}
		outCfg.InboundCCVsToAddAboveThreshold = append(outCfg.InboundCCVsToAddAboveThreshold, fullCCVRef.Address)
	}
	return outCfg, nil
}

// LegacyFeesForAutoMigrate reads legacy lane fees from the fee quoter or onramp and resolves them
// against rc's YAML input. When YAML specifies tokenTransferFeeConfig, fields merge with legacy;
// when YAML omits fees, legacy is imported when enabled. Returns nil when legacy is disabled and
// YAML omits fees. Partial YAML without isEnabled is rejected.
func LegacyFeesForAutoMigrate[R any, CCV any](
	e cldf.Environment,
	localSelector uint64,
	remoteSelector uint64,
	tokenAddress string,
	rc RemoteChainConfig[R, CCV],
) (*PartialTokenTransferFeeConfig, error) {
	input := rc.TokenTransferFeeConfig
	if input != nil {
		if enabled, ok := input.IsEnabled.Get(); ok && !enabled {
			return input, nil
		}
	}

	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(e.OperationsBundle, e.BlockChains, e.DataStore, localSelector, remoteSelector)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to resolve fee adapter for chain selector %d and remote chain selector %d: %w",
			localSelector, remoteSelector, err,
		)
	}

	legacyTTFC, err := feeAdapter.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, fqRef, localSelector, remoteSelector, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to discover token transfer fee config for remote chain selector %d on chain selector %d: %w",
			remoteSelector, localSelector, err,
		)
	}

	legacy := TokenTransferFeeConfig{
		DestGasOverhead:               legacyTTFC.DestGasOverhead,
		DestBytesOverhead:             legacyTTFC.DestBytesOverhead,
		DefaultFinalityFeeUSDCents:    legacyTTFC.MinFeeUSDCents,
		CustomFinalityFeeUSDCents:     0,
		DefaultFinalityTransferFeeBps: legacyTTFC.DeciBps,
		CustomFinalityTransferFeeBps:  0,
		IsEnabled:                     legacyTTFC.IsEnabled,
	}

	if legacy.IsEnabled {
		var partial PartialTokenTransferFeeConfig
		if input == nil {
			partial = partial.Populate(legacy)
			return &partial, nil
		}
		if _, ok := input.IsEnabled.Get(); ok {
			partial = partial.Populate(input.MergeWith(legacy))
			return &partial, nil
		}
		return nil, fmt.Errorf("tokenTransferFeeConfig must set isEnabled")
	}
	if input == nil {
		return nil, nil
	}
	if _, ok := input.IsEnabled.Get(); !ok {
		return nil, fmt.Errorf("tokenTransferFeeConfig must set isEnabled")
	}

	return input, nil
}

// LegacyRateLimitsForAutoMigrate reads default-bucket rate limits from the legacy active pool and
// resolves them against rc's YAML input. When YAML specifies both directions, limits are validated
// and nil is returned. When YAML omits both, on-chain limits are returned for MigrationMetadata
// (bigint passthrough, including disabled or mixed enablement). Partial YAML is rejected.
// legacyPoolRef supplies version/type for inbound decimal-basis normalization; rc supplies remote decimals.
func LegacyRateLimitsForAutoMigrate[R any, CCV any](
	e cldf.Environment,
	reader RateLimitReaderAdapter,
	localSelector uint64,
	remoteSelector uint64,
	legacyPoolRef datastore.AddressRef,
	tokenRef datastore.AddressRef,
	localDecimals uint8,
	rc RemoteChainConfig[R, CCV],
) (*OnchainRateLimits, error) {
	defaultOutbound, defaultOutboundOk := rc.GetOutboundRateLimitBuckets().DefaultBucket()
	defaultInbound, defaultInboundOk := rc.GetInboundRateLimitBuckets().DefaultBucket()
	if defaultOutboundOk != defaultInboundOk {
		return nil, fmt.Errorf(
			"default outbound and inbound rate limits must both be specified together in deployment input or fully omitted",
		)
	}

	if defaultOutboundOk && defaultInboundOk {
		if err := defaultOutbound.RateLimit.Validate(); err != nil {
			return nil, fmt.Errorf("outbound rate limiter config: %w", err)
		}
		if err := defaultInbound.RateLimit.Validate(); err != nil {
			return nil, fmt.Errorf("inbound rate limiter config: %w", err)
		}
		return nil, nil
	}

	legacy, err := reader.GetOnchainRateLimits(
		e.OperationsBundle,
		e.BlockChains,
		e.DataStore,
		localSelector,
		legacyPoolRef,
		tokenRef,
		remoteSelector,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read on-chain rate limits from legacy pool for remote chain %d: %w",
			remoteSelector, err,
		)
	}

	chainFamily, err := chain_selectors.GetSelectorFamily(localSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", localSelector, err)
	}

	inboundLegacy := legacy.Inbound
	if !DoesPoolUseLocalDecimals(chainFamily, legacyPoolRef.Version, legacyPoolRef.Type.String()) && rc.RemoteDecimals != 0 {
		inboundLegacy = RebaseRateLimiterConfig(legacy.Inbound, rc.RemoteDecimals, localDecimals)
	}

	return &OnchainRateLimits{
		Outbound: legacy.Outbound,
		Inbound:  inboundLegacy,
	}, nil
}

func ResolveTokenFeeAdapter(e cldf.Environment, sel uint64, poolRef datastore.AddressRef) (TokenFeeAdapter, error) {
	registry := GetTokenAdapterRegistry()

	fam, err := chain_selectors.GetSelectorFamily(sel)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain selector family for selector %d: %w", sel, err)
	}

	ref, err := ResolveTokenPoolRef(e, registry, sel, poolRef)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token pool ref: %w", err)
	}

	tokAdapter, ok := registry.GetTokenAdapter(fam, ref.Version)
	if !ok {
		return nil, fmt.Errorf("no token adapter found for chain family %s and pool ref %s", fam, datastore_utils.SprintRef(poolRef))
	}

	feeAdapter, ok := tokAdapter.(TokenFeeAdapter)
	if !ok {
		return nil, fmt.Errorf("token adapter for chain family %s and pool ref %s does not implement TokenFeeAdapter", fam, datastore_utils.SprintRef(poolRef))
	}

	return feeAdapter, nil
}

// GetDefaultChainAgnosticTokenTransferFeeConfig returns sensible default token transfer fee configuration
// for the given source and destination chain selectors. It may be overridden via the optional overrides
// argument, allowing the caller to customize specific fields of the returned config.
func GetDefaultChainAgnosticTokenTransferFeeConfig(src uint64, dst uint64, overrides ...func(*TokenTransferFeeConfig)) TokenTransferFeeConfig {
	var minFeeUSDCents uint32
	switch {
	case src == chain_selectors.ETHEREUM_MAINNET.Selector:
		minFeeUSDCents = 50
	case dst == chain_selectors.ETHEREUM_MAINNET.Selector:
		minFeeUSDCents = 150
	default:
		minFeeUSDCents = 25
	}

	cfg := TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: 0,
		CustomFinalityTransferFeeBps:  0,
		DefaultFinalityFeeUSDCents:    minFeeUSDCents,
		CustomFinalityFeeUSDCents:     minFeeUSDCents,
		DestBytesOverhead:             32,
		DestGasOverhead:               90_000,
		IsEnabled:                     true,
	}

	for _, override := range overrides {
		override(&cfg)
	}

	return cfg
}
