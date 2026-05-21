package tokens

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
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
	// AllowedFinalityConfig, if set, specifies the allowed finality configurations to set on the token pool. If this is unspecified, then one of
	// two things will happen. If this is a new pool, then the onchain code will use a sensible default (e.g. WAIT_FOR_FINALITY). Otherwise, this
	// config will be left as-is, meaning that the existing allowed finality config on the pool remains in place.
	AllowedFinalityConfig finality.Config `yaml:"allowedFinalityConfig" json:"allowedFinalityConfig"`
	// LiquidityMigrationAmount, if set, specifies an exact token amount to migrate from the old pool (read from the
	// TokenAdminRegistry) to the new pool's lockbox. Mutually exclusive with LiquidityMigrationBasisPoints.
	// When either LiquidityMigrationAmount or LiquidityMigrationBasisPoints is set, a liquidity migration is triggered.
	// The old pool address is derived from the TokenAdminRegistry, and the timelock address from the MCMS config.
	LiquidityMigrationAmount *big.Int `yaml:"liquidityMigrationAmount" json:"liquidityMigrationAmount"`
	// LiquidityMigrationBasisPoints specifies a percentage of the old pool's balance to migrate (1-10000, where 10000 = 100%).
	// Mutually exclusive with LiquidityMigrationAmount.
	LiquidityMigrationBasisPoints *uint16 `yaml:"liquidityMigrationBasisPoints,string" json:"liquidityMigrationBasisPoints,string"`
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
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) error {
		// TODO: implement
		return nil
	}
}

func makeApply(_ *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
		configs := make(map[uint64]TokenTransferConfig, len(cfg.Tokens))
		for _, config := range cfg.Tokens {
			configs[config.ChainSelector] = config
		}
		batchOps, reports, ds, err := processTokenConfigForChain(e, mcmsRegistry, cfg.MCMS, configs)
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

func processTokenConfigForChain(e deployment.Environment, mcmsRegistry *changesets.MCMSReaderRegistry, mcmsInput mcms.Input, cfg map[uint64]TokenTransferConfig) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], *datastore.MemoryDataStore, error) {
	tokenRegistry := GetTokenAdapterRegistry()
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	ds := datastore.NewMemoryDataStore()

	var err error
	for selector, token := range cfg {
		token.RegistryRef, err = TryNormalizeAddressRef(selector, token.RegistryRef)
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

		// Resolve the timelock address if a liquidity migration is requested.
		var timelockAddress string
		if token.LiquidityMigrationAmount != nil || token.LiquidityMigrationBasisPoints != nil {
			mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
			if !ok {
				return nil, nil, nil, fmt.Errorf("no MCMS reader registered for chain family '%s' on chain %d", family, selector)
			}
			timelockRef, err := mcmsReader.GetTimelockRef(e, selector, mcmsInput)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to get timelock address from MCMS config on chain %d: %w", selector, err)
			}
			timelockAddress = timelockRef.Address
		}

		configureTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureTokenForTransfersSequence(), e.BlockChains, ConfigureTokenForTransfersInput{
			ChainSelector:                 selector,
			TokenPoolAddress:              tokenPool.Address,
			RemoteChains:                  remoteChains,
			ExternalAdmin:                 token.ExternalAdmin,
			RegistryAddress:               registryAddr,
			TokenRef:                      fullTokenRef,
			PoolType:                      tokenPool.Type.String(),
			ExistingDataStore:             e.DataStore,
			AllowedFinalityConfig:         token.AllowedFinalityConfig,
			LiquidityMigrationAmount:      token.LiquidityMigrationAmount,
			LiquidityMigrationBasisPoints: token.LiquidityMigrationBasisPoints,
			TimelockAddress:               timelockAddress,
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
	}
	return batchOps, reports, ds, nil
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
			e.Logger.Infof("Successfully derived remote token address %s for remote chain selector %d from remote pool ref %s", derivedTokenAddr, remoteChainSelector, datastore_utils.SprintRef(*inCfg.RemotePool))
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
		ref, err := TryNormalizeAddressRef(chainSelector, ccvRef)
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
		ref, err := TryNormalizeAddressRef(chainSelector, ccvRef)
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
		ref, err := TryNormalizeAddressRef(chainSelector, ccvRef)
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
		ref, err := TryNormalizeAddressRef(chainSelector, ccvRef)
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
