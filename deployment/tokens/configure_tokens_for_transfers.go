package tokens

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var ConfigureTokensForTransfers = cldf.CreateChangeSet(apply, verify)

type TokenTransferConfig struct {
	// ChainSelector identifies the chain on which the token lives.
	ChainSelector uint64
	// TokenPoolRef is a reference to the token pool in the datastore.
	// Populate the reference as needed to match the desired token pool.
	TokenPoolRef datastore.AddressRef
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	// Leave empty to use internal administration.
	ExternalAdmin string
	// RegistryRef is a reference to the contract on which the token pool must be registered.
	// Populate the reference as needed to match the desired registry.
	RegistryAddress datastore.AddressRef
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]
}

type ConfigureTokensForTransfersConfig struct {
	Tokens []TokenTransferConfig
	MCMS   mcms.Input
}

func verify(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) error {
	// TODO: implement
	return nil
}

func apply(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (cldf.ChangesetOutput, error) {
	writes := make([]contract.WriteOutput, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	for _, token := range cfg.Tokens {
		refs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{
			token.TokenPoolRef,
			token.RegistryAddress,
		}, datastore_utils.FullRef)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool and registry refs on chain with selector %d: %w", token.ChainSelector, err)
		}
		tokenPool := refs[0]
		registry := refs[1]

		family, err := chain_selectors.GetSelectorFamily(token.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", token.ChainSelector, err)
		}
		adapterID := newTokenAdapterID(family, tokenPool.Version)
		adapter, ok := registeredTokenAdapters[adapterID]
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no token adapter registered for chain family '%s' and token pool version '%s'", family, tokenPool.Version)
		}

		remoteChains := make(map[uint64]RemoteChainConfig[[]byte, string], len(token.RemoteChains))
		for remoteChainSelector, inCfg := range token.RemoteChains {
			remoteChains[remoteChainSelector], err = convertRemoteChainConfig(e, adapter, token.ChainSelector, remoteChainSelector, inCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to process remote chain config for remote chain selector %d: %w", remoteChainSelector, err)
			}
		}
		configureTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureTokenForTransfersSequence(), e.BlockChains, ConfigureTokenForTransfersInput{
			ChainSelector:    token.ChainSelector,
			TokenPoolAddress: tokenPool.Address,
			RemoteChains:     remoteChains,
			ExternalAdmin:    token.ExternalAdmin,
			RegistryAddress:  registry.Address,
		})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", token.ChainSelector, err)
		}

		writes = append(writes, configureTokenReport.Output.Writes...)
		reports = append(reports, configureTokenReport.ExecutionReports...)
	}

	timelockAddresses, err := getTimelockAddresses(e, cfg)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to get timelock addresses: %w", err)
	}
	chainMetadata, err := getChainMetadata(e, cfg)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain metadata: %w", err)
	}

	return changesets.NewOutputBuilder().
		WithReports(reports).
		WithWriteOutputs(writes).
		Build(changesets.MCMSBuildParams{
			Input:             cfg.MCMS,
			Description:       "Configure tokens for transfers",
			TimelockAddresses: timelockAddresses,
			ChainMetadata:     chainMetadata,
		})
}

func getTimelockAddresses(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (map[mcms_types.ChainSelector]string, error) {
	timelocks := make(map[mcms_types.ChainSelector]string)
	for _, token := range cfg.Tokens {
		chainSel := mcms_types.ChainSelector(token.ChainSelector)
		if _, exists := timelocks[chainSel]; exists {
			continue // Already have timelock for this chain
		}
		refs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{cfg.MCMS.TimelockAddressRef}, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve timelock ref on chain with selector %d: %w", token.ChainSelector, err)
		}
		timelocks[chainSel] = refs[0].Address
	}

	return timelocks, nil
}

func getChainMetadata(e cldf.Environment, cfg ConfigureTokensForTransfersConfig) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error) {
	metadata := make(map[mcms_types.ChainSelector]mcms_types.ChainMetadata)
	for _, token := range cfg.Tokens {
		chainSel := mcms_types.ChainSelector(token.ChainSelector)
		if _, exists := metadata[chainSel]; exists {
			continue // Already have metadata for this chain
		}
		refs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{cfg.MCMS.MCMSAddressRef}, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve mcm ref on chain with selector %d: %w", token.ChainSelector, err)
		}
		mcmAddress := refs[0].Address

		family, err := chain_selectors.GetSelectorFamily(token.ChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", token.ChainSelector, err)
		}
		adapterID := newTokenAdapterID(family, token.TokenPoolRef.Version)
		adapter, ok := registeredTokenAdapters[adapterID]
		if !ok {
			return nil, fmt.Errorf("no token adapter registered for chain family '%s' and token pool version '%s'", family, token.TokenPoolRef.Version)
		}

		opCount, err := adapter.GetCurrentOpCount(e, token.ChainSelector, mcmAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get current op count from MCMS at address %s on chain with selector %d: %w", mcmAddress, token.ChainSelector, err)
		}
		metadata[chainSel] = mcms_types.ChainMetadata{
			MCMAddress:      mcmAddress,
			StartingOpCount: opCount,
		}
	}

	return metadata, nil
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
	if inCfg.RemoteToken != nil {
		inCfg.RemoteToken.ChainSelector = remoteChainSelector
		tokenAddr, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{*inCfg.RemoteToken}, adapter.ConvertRefToBytes)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote token ref %s: %w", datastore_utils.SprintRef(*inCfg.RemoteToken), err)
		}
		outCfg.RemoteToken = tokenAddr[0]
	}
	if inCfg.RemotePool != nil {
		inCfg.RemotePool.ChainSelector = remoteChainSelector
		poolAddr, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{*inCfg.RemotePool}, adapter.ConvertRefToBytes)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve remote pool ref %s: %w", datastore_utils.SprintRef(*inCfg.RemotePool), err)
		}
		outCfg.RemotePool = poolAddr[0]
	}
	for _, ccvRef := range inCfg.OutboundCCVs {
		ccvRef.ChainSelector = chainSelector
		fullCCVRef, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{ccvRef}, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve outbound CCV ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.OutboundCCVs = append(outCfg.OutboundCCVs, fullCCVRef[0].Address)
	}
	for _, ccvRef := range inCfg.InboundCCVs {
		ccvRef.ChainSelector = remoteChainSelector
		fullCCVRef, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{ccvRef}, datastore_utils.FullRef)
		if err != nil {
			return outCfg, fmt.Errorf("failed to resolve inbound CCV ref %s: %w", datastore_utils.SprintRef(ccvRef), err)
		}
		outCfg.InboundCCVs = append(outCfg.InboundCCVs, fullCCVRef[0].Address)
	}
	return outCfg, nil
}
