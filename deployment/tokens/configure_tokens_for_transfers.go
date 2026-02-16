package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TokenTransferConfig specifies configuration for a token on one chain to enable transfers with other chains.
type TokenTransferConfig struct {
	// ChainSelector identifies the chain on which the token lives.
	ChainSelector uint64
	// TokenPoolRef is a reference to the token pool in the datastore.
	// Populate the reference as needed to match the desired token pool.
	TokenPoolRef datastore.AddressRef
	// TokenRef is a reference to the token in the datastore. This is only needed if the token address cannot be derived from the pool reference.
	TokenRef datastore.AddressRef
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	// Leave empty to use internal administration.
	ExternalAdmin string
	// RegistryRef is a reference to the contract on which the token pool must be registered.
	// Populate the reference as needed to match the desired registry.
	RegistryRef datastore.AddressRef
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]
}

// ConfigureTokensForTransfersConfig is the configuration for the ConfigureTokensForTransfers changeset.
type ConfigureTokensForTransfersConfig struct {
	// ChainAdapterVersion specifies the version of the chain adapter to use for this chain.
	ChainAdapterVersion *semver.Version `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
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

func makeApply(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
		batchOps, reports, err := processTokenConfigForChain(e, cfg.Tokens, cfg.ChainAdapterVersion)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to process token configs for chains: %w", err)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func processTokenConfigForChain(e deployment.Environment, cfg []TokenTransferConfig, chainAdapterVersion *semver.Version) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	tokenRegistry := GetTokenAdapterRegistry()
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	for _, token := range cfg {
		tokenPool, err := datastore_utils.FindAndFormatRef(e.DataStore, token.TokenPoolRef, token.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve token pool ref on chain with selector %d: %w", token.ChainSelector, err)
		}
		registry, err := datastore_utils.FindAndFormatRef(e.DataStore, token.RegistryRef, token.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve registry ref on chain with selector %d: %w", token.ChainSelector, err)
		}

		remoteChains := make(map[uint64]RemoteChainConfig[[]byte, string], len(token.RemoteChains))
		for remoteChainSelector, inCfg := range token.RemoteChains {
			family, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
			}
			adapter, ok := tokenRegistry.GetTokenAdapter(family, chainAdapterVersion)
			if !ok {
				return nil, nil, fmt.Errorf("no token adapter registered for chain family '%s' and chain adapter version '%s'", family, chainAdapterVersion)
			}
			remoteChains[remoteChainSelector], err = convertRemoteChainConfig(e, adapter, token.ChainSelector, remoteChainSelector, inCfg)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to process remote chain config for remote chain selector %d: %w", remoteChainSelector, err)
			}
		}

		family, err := chain_selectors.GetSelectorFamily(token.ChainSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", token.ChainSelector, err)
		}
		adapter, ok := tokenRegistry.GetTokenAdapter(family, chainAdapterVersion)
		if !ok {
			return nil, nil, fmt.Errorf("no token adapter registered for chain family '%s' and chain adapter version '%s'", family, chainAdapterVersion)
		}
		configureTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureTokenForTransfersSequence(), e.BlockChains, ConfigureTokenForTransfersInput{
			ChainSelector:     token.ChainSelector,
			TokenPoolAddress:  tokenPool.Address,
			RemoteChains:      remoteChains,
			ExternalAdmin:     token.ExternalAdmin,
			RegistryAddress:   registry.Address,
			TokenRef:          token.TokenRef,
			PoolType:          tokenPool.Type.String(),
			ExistingDataStore: e.DataStore,
		})
		if err != nil {
			return batchOps, reports, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", token.ChainSelector, err)
		}

		batchOps = append(batchOps, configureTokenReport.Output.BatchOps...)
		reports = append(reports, configureTokenReport.ExecutionReports...)
	}
	return batchOps, reports, nil
}

func convertRemoteChainConfig(
	e deployment.Environment,
	adapter TokenAdapter,
	chainSelector uint64,
	remoteChainSelector uint64,
	inCfg RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef],
) (RemoteChainConfig[[]byte, string], error) {
	outCfg := RemoteChainConfig[[]byte, string]{
		InboundRateLimiterConfig:  inCfg.InboundRateLimiterConfig,
		OutboundRateLimiterConfig: inCfg.OutboundRateLimiterConfig,
	}
	if inCfg.RemotePool != nil {
		fullRemotePoolRef, err := datastore_utils.FindAndFormatRef(e.DataStore, *inCfg.RemotePool, remoteChainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote pool ref %s: %w", datastore_utils.SprintRef(*inCfg.RemotePool), err)
		}
		outCfg.RemotePool, err = adapter.AddressRefToBytes(fullRemotePoolRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to convert remote pool ref %s to bytes: %w", datastore_utils.SprintRef(*inCfg.RemotePool), err)
		}
		// Can either provide the token reference directly or derive it from the pool reference.
		if inCfg.RemoteToken != nil {
			outCfg.RemoteToken, err = datastore_utils.FindAndFormatRef(e.DataStore, *inCfg.RemoteToken, remoteChainSelector, adapter.AddressRefToBytes)
			if err != nil {
				return outCfg, fmt.Errorf("failed to resolve remote token ref %s: %w", datastore_utils.SprintRef(*inCfg.RemoteToken), err)
			}
		} else {
			outCfg.RemoteToken, err = adapter.DeriveTokenAddress(e, remoteChainSelector, fullRemotePoolRef)
			if err != nil {
				return outCfg, fmt.Errorf("failed to get remote token address via pool ref (%s) for remote chain selector %d: %w", datastore_utils.SprintRef(*inCfg.RemotePool), remoteChainSelector, err)
			}
		}
		outCfg.RemoteToken = common.LeftPadBytes(outCfg.RemoteToken, 32)
		outCfg.RemoteDecimals, err = adapter.DeriveTokenDecimals(e, remoteChainSelector, fullRemotePoolRef, outCfg.RemoteToken)
		if err != nil {
			return outCfg, fmt.Errorf("failed to get remote token decimals for remote chain selector %d: %w", remoteChainSelector, err)
		}
		outCfg.RemotePool, err = adapter.DeriveTokenPoolCounterpart(e, remoteChainSelector, outCfg.RemotePool, outCfg.RemoteToken)
		if err != nil {
			return outCfg, fmt.Errorf("failed to derive remote pool counterpart for remote chain selector %d: %w", remoteChainSelector, err)
		}
	}
	for _, ccvRef := range inCfg.OutboundCCVs {
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve outbound CCV ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.OutboundCCVs = append(outCfg.OutboundCCVs, fullCCVRef.Address)
	}
	for _, ccvRef := range inCfg.InboundCCVs {
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve inbound CCV ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.InboundCCVs = append(outCfg.InboundCCVs, fullCCVRef.Address)
	}
	return outCfg, nil
}
