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

// TokenTransferConfig specifies configuration for a token on one chain to enable transfers with other chains.
type TokenTransferConfig struct {
	// ChainSelector identifies the chain on which the token lives.
	ChainSelector uint64
	// TokenPool is a reference to the token pool in the datastore.
	// Populate the reference as needed to match the desired token pool.
	TokenPool datastore.AddressRef
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	// Leave empty to use internal administration.
	ExternalAdmin string
	// Registry is a reference to the contract on which the token pool must be registered.
	// Populate the reference as needed to match the desired registry.
	Registry datastore.AddressRef
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]
	// MinFinalityValue is the minimum finality value required by the token pool.
	// This can be interpreted as # of block confirmations, an ID, or otherwise.
	// Interpretation is left to each chain family.
	MinFinalityValue uint16
}

// ConfigureTokensForTransfersConfig is the configuration for the ConfigureTokensForTransfers changeset.
type ConfigureTokensForTransfersConfig struct {
	// Tokens specifies the tokens to configure for cross-chain transfers.
	Tokens []TokenTransferConfig
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

// ConfigureTokensForTransfers returns a changeset that configures tokens on multiple chains for transfers with other chains.
func ConfigureTokensForTransfers(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureTokensForTransfersConfig] {
	return cldf.CreateChangeSet(makeApply(tokenRegistry, mcmsRegistry), makeVerify(tokenRegistry, mcmsRegistry))
}

func makeVerify(_ *TokenAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) error {
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) error {
		if cfg.MCMS != nil {
			err := cfg.MCMS.Validate()
			if err != nil {
				return fmt.Errorf("failed to validate MCMS input: %w", err)
			}
		}

		for _, token := range cfg.Tokens {
			if _, err := chain_selectors.GetSelectorFamily(token.ChainSelector); err != nil {
				return err
			}
			if datastore_utils.IsAddressRefEmpty(token.TokenPool) {
				return fmt.Errorf("token pool ref is empty for token on chain with selector %d", token.ChainSelector)
			}
			for remoteChainSelector := range token.RemoteChains {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return fmt.Errorf("remote chain %d has unknown chain selector %d: %w", remoteChainSelector, remoteChainSelector, err)
				}
			}
		}

		return nil
	}
}

func makeApply(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, token := range cfg.Tokens {
			var input ConfigureTokenForTransfersInput

			if !datastore_utils.IsAddressRefEmpty(token.Registry) {
				registry, err := datastore_utils.FindAndFormatRef(e.DataStore, token.Registry, token.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve registry ref on chain with selector %d: %w", token.ChainSelector, err)
				}
				input.RegistryAddress = registry.Address
			}

			tokenPool, err := datastore_utils.FindAndFormatRef(e.DataStore, token.TokenPool, token.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool ref on chain with selector %d: %w", token.ChainSelector, err)
			}
			input.TokenPoolAddress = tokenPool.Address

			family, err := chain_selectors.GetSelectorFamily(token.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", token.ChainSelector, err)
			}
			adapter, ok := tokenRegistry.GetTokenAdapter(family, tokenPool.Version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no token adapter registered for chain family '%s' and token pool version '%s'", family, tokenPool.Version)
			}

			remoteChains := make(map[uint64]RemoteChainConfig[[]byte, string], len(token.RemoteChains))
			for remoteChainSelector, inCfg := range token.RemoteChains {
				remoteFamily, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
				}
				remoteAdapter, ok := tokenRegistry.GetTokenAdapter(remoteFamily, inCfg.RemotePool.Version)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no token adapter registered for chain family '%s' and token pool version '%s'", remoteFamily, inCfg.RemotePool.Version)
				}
				remoteChains[remoteChainSelector], err = convertRemoteChainConfig(e, token.ChainSelector, remoteAdapter, remoteChainSelector, inCfg)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to process remote chain config for remote chain selector %d: %w", remoteChainSelector, err)
				}
			}

			input.RemoteChains = remoteChains
			input.ChainSelector = token.ChainSelector
			input.ExternalAdmin = token.ExternalAdmin
			input.MinFinalityValue = token.MinFinalityValue

			configureTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureTokenForTransfersSequence(), e.BlockChains, input)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", token.ChainSelector, err)
			}

			batchOps = append(batchOps, configureTokenReport.Output.BatchOps...)
			reports = append(reports, configureTokenReport.ExecutionReports...)
		}

		var mcmsInput mcms.Input
		if cfg.MCMS != nil {
			mcmsInput = *cfg.MCMS
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func convertRemoteChainConfig(
	e cldf.Environment,
	chainSelector uint64,
	remoteAdapter TokenAdapter,
	remoteChainSelector uint64,
	inCfg RemoteChainConfig[datastore.AddressRef, datastore.AddressRef],
) (RemoteChainConfig[[]byte, string], error) {
	outCfg := RemoteChainConfig[[]byte, string]{
		DefaultFinalityInboundRateLimiterConfig:  inCfg.DefaultFinalityInboundRateLimiterConfig,
		DefaultFinalityOutboundRateLimiterConfig: inCfg.DefaultFinalityOutboundRateLimiterConfig,
		CustomFinalityInboundRateLimiterConfig:   inCfg.CustomFinalityInboundRateLimiterConfig,
		CustomFinalityOutboundRateLimiterConfig:  inCfg.CustomFinalityOutboundRateLimiterConfig,
		TokenTransferFeeConfig:                   inCfg.TokenTransferFeeConfig,
	}
	if !datastore_utils.IsAddressRefEmpty(inCfg.RemotePool) {
		fullRemotePoolRef, err := datastore_utils.FindAndFormatRef(e.DataStore, inCfg.RemotePool, remoteChainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote pool ref %s: %w", datastore_utils.SprintRef(inCfg.RemotePool), err)
		}
		outCfg.RemotePool, err = remoteAdapter.AddressRefToBytes(fullRemotePoolRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to convert remote pool ref %s to bytes: %w", datastore_utils.SprintRef(inCfg.RemotePool), err)
		}
		// Can either provide the token reference directly or derive it from the pool reference.
		if !datastore_utils.IsAddressRefEmpty(inCfg.RemoteToken) {
			outCfg.RemoteToken, err = datastore_utils.FindAndFormatRef(e.DataStore, inCfg.RemoteToken, remoteChainSelector, remoteAdapter.AddressRefToBytes)
			if err != nil {
				return outCfg, fmt.Errorf("failed to resolve remote token ref %s: %w", datastore_utils.SprintRef(inCfg.RemoteToken), err)
			}
		} else {
			outCfg.RemoteToken, err = remoteAdapter.DeriveTokenAddress(e, remoteChainSelector, fullRemotePoolRef)
			if err != nil {
				return outCfg, fmt.Errorf("failed to get remote token address via pool ref (%s) for remote chain selector %d: %w", datastore_utils.SprintRef(inCfg.RemotePool), remoteChainSelector, err)
			}
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
	for _, ccvRef := range inCfg.OutboundCCVsToAddAboveThreshold {
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve outbound CCV to add above threshold ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.OutboundCCVsToAddAboveThreshold = append(outCfg.OutboundCCVsToAddAboveThreshold, fullCCVRef.Address)
	}
	for _, ccvRef := range inCfg.InboundCCVsToAddAboveThreshold {
		fullCCVRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve inbound CCV to add above threshold ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.InboundCCVsToAddAboveThreshold = append(outCfg.InboundCCVsToAddAboveThreshold, fullCCVRef.Address)
	}
	return outCfg, nil
}
