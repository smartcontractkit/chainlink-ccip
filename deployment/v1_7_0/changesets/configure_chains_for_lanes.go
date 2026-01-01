package changesets

import (
	"fmt"
	"slices"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// ChainConfig specifies configuration required for a chain to connect with other chains.
type ChainConfig struct {
	// The selector of the chain being configured.
	ChainSelector uint64
	// The Router on the chain being configured.
	// We assume that all connections defined will use the same router, either test or production.
	Router datastore.AddressRef
	// The OnRamp on the chain being configured.
	// Similarly, we assume that all connections will use the same OnRamp.
	OnRamp datastore.AddressRef
	// The CommitteeVerifiers on the chain being configured.
	// There can be multiple committee verifiers on a chain, each controlled by a different entity.
	CommitteeVerifiers []adapters.CommitteeVerifierConfig[datastore.AddressRef]
	// The FeeQuoter on the chain being configured.
	FeeQuoter datastore.AddressRef
	// The OffRamp on the chain being configured
	OffRamp datastore.AddressRef
	// The configuration for each remote chain that we want to connect to.
	RemoteChains map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]
	// The remote chains that we wish to disconnect from.
	RemoteChainsToDisconnect []uint64
}

// ConfigureChainsForLanesConfig is the configuration for the ConfigureChainsForLanes changeset.
type ConfigureChainsForLanesConfig struct {
	// Chains specifies the chains to configure.
	Chains []ChainConfig
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

// ConfigureChainsForLanes returns a changeset that configures chains for messaging with other chains.
func ConfigureChainsForLanes(chainFamilyRegistry *adapters.ChainFamilyRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureChainsForLanesConfig] {
	return cldf.CreateChangeSet(makeApply(chainFamilyRegistry, mcmsRegistry), makeVerify(chainFamilyRegistry, mcmsRegistry))
}

func makeVerify(_ *adapters.ChainFamilyRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureChainsForLanesConfig) error {
	return func(e cldf.Environment, cfg ConfigureChainsForLanesConfig) error {
		if cfg.MCMS != nil {
			err := cfg.MCMS.Validate()
			if err != nil {
				return fmt.Errorf("failed to validate MCMS input: %w", err)
			}
		}

		for _, chain := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chain.ChainSelector); err != nil {
				return err
			}
			if datastore_utils.IsAddressRefEmpty(chain.Router) {
				return fmt.Errorf("router ref is empty for chain with selector %d", chain.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(chain.OnRamp) {
				return fmt.Errorf("onRamp ref is empty for chain with selector %d", chain.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(chain.FeeQuoter) {
				return fmt.Errorf("feeQuoter ref is empty for chain with selector %d", chain.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(chain.OffRamp) {
				return fmt.Errorf("offRamp ref is empty for chain with selector %d", chain.ChainSelector)
			}

			for _, ccv := range chain.CommitteeVerifiers {
				if len(ccv.CommitteeVerifier) == 0 {
					return fmt.Errorf("committee verifier on chain with selector %d has no contracts", chain.ChainSelector)
				}
				for remoteChainSelector, remoteChain := range ccv.RemoteChains {
					if slices.Contains(chain.RemoteChainsToDisconnect, remoteChainSelector) {
						return fmt.Errorf("committee verifier on chain %d with remote chain %d has remote chain %d in both RemoteChains and RemoteChainsToDisconnect", chain.ChainSelector, remoteChainSelector, remoteChainSelector)
					}
					if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
						return err
					}
					if int(remoteChain.SignatureConfig.Threshold) > len(remoteChain.SignatureConfig.Signers) {
						return fmt.Errorf("committee verifier on chain %d with remote chain %d has threshold greater than the number of signers", chain.ChainSelector, remoteChainSelector)
					}
				}
			}

			for remoteChainSelector, remoteChain := range chain.RemoteChains {
				if slices.Contains(chain.RemoteChainsToDisconnect, remoteChainSelector) {
					return fmt.Errorf("chain %d has remote chain %d in both RemoteChains and RemoteChainsToDisconnect", chain.ChainSelector, remoteChainSelector)
				}
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return err
				}
				if datastore_utils.IsAddressRefEmpty(remoteChain.OffRamp) {
					return fmt.Errorf("chain %d has empty offRamp ref for remote chain %d", chain.ChainSelector, remoteChainSelector)
				}
				if len(remoteChain.OnRamps) == 0 {
					return fmt.Errorf("chain %d has no onRamps for remote chain %d", chain.ChainSelector, remoteChainSelector)
				}
				if datastore_utils.IsAddressRefEmpty(remoteChain.DefaultExecutor) {
					return fmt.Errorf("chain %d has empty default executor ref for remote chain %d", chain.ChainSelector, remoteChainSelector)
				}
			}

			for _, remoteChainToDisconnect := range chain.RemoteChainsToDisconnect {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainToDisconnect); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

func makeApply(chainFamilyRegistry *adapters.ChainFamilyRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureChainsForLanesConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConfigureChainsForLanesConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, chain := range cfg.Chains {
			router, err := datastore_utils.FindAndFormatRef(e.DataStore, chain.Router, chain.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve router ref on chain with selector %d: %w", chain.ChainSelector, err)
			}
			onRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, chain.OnRamp, chain.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve onRamp ref on chain with selector %d: %w", chain.ChainSelector, err)
			}
			feeQuoter, err := datastore_utils.FindAndFormatRef(e.DataStore, chain.FeeQuoter, chain.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve feeQuoter ref on chain with selector %d: %w", chain.ChainSelector, err)
			}
			offRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, chain.OffRamp, chain.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve offRamp ref on chain with selector %d: %w", chain.ChainSelector, err)
			}

			committeeVerifiers := make([]adapters.CommitteeVerifierConfig[string], len(chain.CommitteeVerifiers))
			for i, verifier := range chain.CommitteeVerifiers {
				contracts := make([]string, 0, len(verifier.CommitteeVerifier))
				for _, contract := range verifier.CommitteeVerifier {
					contract, err := datastore_utils.FindAndFormatRef(e.DataStore, contract, chain.ChainSelector, datastore_utils.FullRef)
					if err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve CommitteeVerifier contract ref on chain with selector %d: %w", chain.ChainSelector, err)
					}
					contracts = append(contracts, contract.Address)
				}
				committeeVerifiers[i] = adapters.CommitteeVerifierConfig[string]{
					CommitteeVerifier: contracts,
					RemoteChains:      verifier.RemoteChains,
				}
			}

			family, err := chain_selectors.GetSelectorFamily(chain.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chain.ChainSelector, err)
			}
			adapter, ok := chainFamilyRegistry.GetChainFamily(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no adapter registered for chain family '%s'", family)
			}

			remoteChains := make(map[uint64]adapters.RemoteChainConfig[[]byte, string], len(chain.RemoteChains))
			for remoteChainSelector, inCfg := range chain.RemoteChains {
				remoteFamily, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
				}
				remoteAdapter, ok := chainFamilyRegistry.GetChainFamily(remoteFamily)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no remote adapter registered for remote chain family '%s'", remoteFamily)
				}

				remoteChains[remoteChainSelector], err = convertRemoteChainConfig(e, chain.ChainSelector, remoteChainSelector, remoteAdapter, inCfg)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to process remote chain config for selector %d: %w", remoteChainSelector, err)
				}
			}
			configureChainForLanesReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureChainForLanes(), e.BlockChains, adapters.ConfigureChainForLanesInput{
				ChainSelector:            chain.ChainSelector,
				Router:                   router.Address,
				OnRamp:                   onRamp.Address,
				CommitteeVerifiers:       committeeVerifiers,
				FeeQuoter:                feeQuoter.Address,
				OffRamp:                  offRamp.Address,
				RemoteChains:             remoteChains,
				RemoteChainsToDisconnect: chain.RemoteChainsToDisconnect,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure chain with selector %d: %w", chain.ChainSelector, err)
			}

			batchOps = append(batchOps, configureChainForLanesReport.Output.BatchOps...)
			reports = append(reports, configureChainForLanesReport.ExecutionReports...)
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
	e deployment.Environment,
	chainSelector uint64,
	remoteChainSelector uint64,
	remoteChainAdapter adapters.ChainFamily,
	inCfg adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef],
) (adapters.RemoteChainConfig[[]byte, string], error) {
	outCfg := adapters.RemoteChainConfig[[]byte, string]{
		DisableTrafficFrom:        inCfg.DisableTrafficFrom,
		FeeQuoterDestChainConfig:  inCfg.FeeQuoterDestChainConfig,
		ExecutorDestChainConfig:   inCfg.ExecutorDestChainConfig,
		AddressBytesLength:        inCfg.AddressBytesLength,
		BaseExecutionGasCost:      inCfg.BaseExecutionGasCost,
		MessageNetworkFeeUSDCents: inCfg.MessageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   inCfg.TokenNetworkFeeUSDCents,
	}

	onRamps := make([][]byte, 0, len(inCfg.OnRamps))
	for _, onRamp := range inCfg.OnRamps {
		onRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, onRamp, remoteChainSelector, remoteChainAdapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve onRamp ref on remote chain with selector %d: %w", remoteChainSelector, err)
		}
		onRamps = append(onRamps, onRamp)
	}
	outCfg.OnRamps = onRamps

	offRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, inCfg.OffRamp, remoteChainSelector, remoteChainAdapter.AddressRefToBytes)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve offRamp ref on remote chain with selector %d: %w", remoteChainSelector, err)
	}
	outCfg.OffRamp = offRamp

	executor, err := datastore_utils.FindAndFormatRef(e.DataStore, inCfg.DefaultExecutor, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve executor ref on chain with selector %d: %w", remoteChainSelector, err)
	}
	outCfg.DefaultExecutor = executor.Address

	laneMandatedInboundCCVs := make([]string, 0, len(inCfg.LaneMandatedInboundCCVs))
	for _, ccv := range inCfg.LaneMandatedInboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(e.DataStore, ccv, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainSelector, err)
		}
		laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, resolvedCCV.Address)
	}
	outCfg.LaneMandatedInboundCCVs = laneMandatedInboundCCVs

	laneMandatedOutboundCCVs := make([]string, 0, len(inCfg.LaneMandatedOutboundCCVs))
	for _, ccv := range inCfg.LaneMandatedOutboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(e.DataStore, ccv, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainSelector, err)
		}
		laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, resolvedCCV.Address)
	}
	outCfg.LaneMandatedOutboundCCVs = laneMandatedOutboundCCVs

	defaultInboundCCVs := make([]string, 0, len(inCfg.DefaultInboundCCVs))
	for _, ccv := range inCfg.DefaultInboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(e.DataStore, ccv, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainSelector, err)
		}
		defaultInboundCCVs = append(defaultInboundCCVs, resolvedCCV.Address)
	}
	outCfg.DefaultInboundCCVs = defaultInboundCCVs

	defaultOutboundCCVs := make([]string, 0, len(inCfg.DefaultOutboundCCVs))
	for _, ccv := range inCfg.DefaultOutboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(e.DataStore, ccv, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainSelector, err)
		}
		defaultOutboundCCVs = append(defaultOutboundCCVs, resolvedCCV.Address)
	}
	outCfg.DefaultOutboundCCVs = defaultOutboundCCVs

	return outCfg, nil
}
