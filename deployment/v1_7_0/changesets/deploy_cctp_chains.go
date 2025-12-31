package changesets

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
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

// DeployCCTPChainsConfig is the configuration for the DeployCCTPChains changeset.
type DeployCCTPChainsConfig struct {
	// Chains specifies the chains to deploy CCTP on.
	Chains []adapters.DeployCCTPInput[datastore.AddressRef, datastore.AddressRef]
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// DeployCCTPChains returns a changeset that deploys CCTP contracts on chains.
func DeployCCTPChains(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[DeployCCTPChainsConfig] {
	return cldf.CreateChangeSet(makeApplyDeployCCTPChains(cctpChainRegistry, mcmsRegistry), makeVerifyDeployCCTPChains(cctpChainRegistry, mcmsRegistry))
}

func makeVerifyDeployCCTPChains(_ *adapters.CCTPChainRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployCCTPChainsConfig) error {
	return func(e cldf.Environment, cfg DeployCCTPChainsConfig) error {
		// TODO: implement
		return nil
	}
}

func makeApplyDeployCCTPChains(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployCCTPChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployCCTPChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, chainCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainCfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainCfg.ChainSelector, err)
			}
			adapter, ok := cctpChainRegistry.GetCCTPChain(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
			}

			// Resolve AddressRefs in the adapter input
			resolvedInput, err := resolveDeployCCTPChainInput(e, chainCfg.ChainSelector, chainCfg, adapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve DeployCCTPInput for chain selector %d: %w", chainCfg.ChainSelector, err)
			}

			// Call into DeployCCTPChain sequence
			deployCCTPChainReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.DeployCCTPChain(), e.BlockChains, resolvedInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy CCTP on chain with selector %d: %w", chainCfg.ChainSelector, err)
			}

			batchOps = append(batchOps, deployCCTPChainReport.Output.BatchOps...)
			reports = append(reports, deployCCTPChainReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

// resolveDeployCCTPChainInput resolves AddressRefs in the adapter input from the datastore.
// The adapter will handle conversion from adapter input to sequence input internally.
func resolveDeployCCTPChainInput(
	e deployment.Environment,
	chainSelector uint64,
	adapterInput adapters.DeployCCTPInput[datastore.AddressRef, datastore.AddressRef],
	adapter adapters.CCTPChain,
) (adapters.DeployCCTPInput[string, []byte], error) {
	out := adapters.DeployCCTPInput[string, []byte]{
		ChainSelector:                    chainSelector,
		TokenMessenger:                   adapterInput.TokenMessenger,
		USDCToken:                        adapterInput.USDCToken,
		DeployerContract:                 adapterInput.DeployerContract,
		MinFinalityValue:                 adapterInput.MinFinalityValue,
		ThresholdAmountForAdditionalCCVs: adapterInput.ThresholdAmountForAdditionalCCVs,
		FastFinalityBps:                  adapterInput.FastFinalityBps,
		StorageLocations:                 adapterInput.StorageLocations,
		Allowlist:                        adapterInput.Allowlist,
		RateLimitAdmin:                   adapterInput.RateLimitAdmin,
		FeeAggregator:                    adapterInput.FeeAggregator,
		AllowlistAdmin:                   adapterInput.AllowlistAdmin,
	}

	if addressRefExists(adapterInput.TokenPools.LegacyCCTPV1Pool) {
		legacyCCTPV1Pool, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenPools.LegacyCCTPV1Pool, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve LegacyCCTPV1Pool for chain selector %d: %w", chainSelector, err)
		}
		out.TokenPools.LegacyCCTPV1Pool = legacyCCTPV1Pool.Address
	}

	if addressRefExists(adapterInput.TokenPools.CCTPV1Pool) {
		cctpv1Pool, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenPools.CCTPV1Pool, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve CCTPV1Pool for chain selector %d: %w", chainSelector, err)
		}
		out.TokenPools.CCTPV1Pool = cctpv1Pool.Address
	}

	if addressRefExists(adapterInput.TokenPools.CCTPV2Pool) {
		cctpv2Pool, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenPools.CCTPV2Pool, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve CCTPV2Pool for chain selector %d: %w", chainSelector, err)
		}
		out.TokenPools.CCTPV2Pool = cctpv2Pool.Address
	}

	if addressRefExists(adapterInput.TokenPools.CCTPV2PoolWithCCV) {
		cctpv2PoolWithCCV, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenPools.CCTPV2PoolWithCCV, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve CCTPV2PoolWithCCV for chain selector %d: %w", chainSelector, err)
		}
		out.TokenPools.CCTPV2PoolWithCCV = cctpv2PoolWithCCV.Address
	}

	if addressRefExists(adapterInput.TokenPools.LockReleasePool) {
		lockReleasePool, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenPools.LockReleasePool, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve LockReleasePool for chain selector %d: %w", chainSelector, err)
		}
		out.TokenPools.LockReleasePool = lockReleasePool.Address
	}

	if addressRefExists(adapterInput.USDCTokenPoolProxy) {
		usdctokenPoolProxy, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.USDCTokenPoolProxy, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve USDCTokenPoolProxy for chain selector %d: %w", chainSelector, err)
		}
		out.USDCTokenPoolProxy = usdctokenPoolProxy.Address
	}

	cctpVerifier := make([]datastore.AddressRef, 0, len(adapterInput.CCTPVerifier))
	for _, ref := range adapterInput.CCTPVerifier {
		fullRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve CCTPVerifier for chain selector %d: %w", chainSelector, err)
		}
		cctpVerifier = append(cctpVerifier, fullRef)
	}
	out.CCTPVerifier = cctpVerifier

	if addressRefExists(adapterInput.MessageTransmitterProxy) {
		messageTransmitterProxy, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.MessageTransmitterProxy, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve MessageTransmitterProxy for chain selector %d: %w", chainSelector, err)
		}
		out.MessageTransmitterProxy = messageTransmitterProxy.Address
	}

	if addressRefExists(adapterInput.TokenAdminRegistry) {
		tokenAdminRegistry, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.TokenAdminRegistry, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve TokenAdminRegistry for chain selector %d: %w", chainSelector, err)
		}
		out.TokenAdminRegistry = tokenAdminRegistry.Address
	}

	if addressRefExists(adapterInput.RMN) {
		rmn, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.RMN, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve RMN for chain selector %d: %w", chainSelector, err)
		}
		out.RMN = rmn.Address
	}

	if addressRefExists(adapterInput.Router) {
		router, err := datastore_utils.FindAndFormatRef(e.DataStore, adapterInput.Router, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve Router for chain selector %d: %w", chainSelector, err)
		}
		out.Router = router.Address
	}

	remoteChains := make(map[uint64]adapters.RemoteCCTPChainConfig[string, []byte])
	for remoteChainSelector, remoteChainCfg := range adapterInput.RemoteChains {
		remoteChainCfg, err := resolveRemoteCCTPChainConfig(e, chainSelector, remoteChainSelector, remoteChainCfg, adapter)
		if err != nil {
			return adapters.DeployCCTPInput[string, []byte]{}, fmt.Errorf("failed to resolve RemoteCCTPChainConfig for remote chain selector %d: %w", remoteChainSelector, err)
		}
		remoteChains[remoteChainSelector] = remoteChainCfg
	}
	out.RemoteChains = remoteChains

	return out, nil
}

func resolveRemoteCCTPChainConfig(
	e deployment.Environment,
	localChainSelector,
	remoteChainSelector uint64,
	remoteChainCfg adapters.RemoteCCTPChainConfig[datastore.AddressRef, datastore.AddressRef],
	adapter adapters.CCTPChain,
) (adapters.RemoteCCTPChainConfig[string, []byte], error) {
	out := adapters.RemoteCCTPChainConfig[string, []byte]{
		TokenPoolConfig: tokens.RemoteChainConfig[[]byte, string]{
			TokenTransferFeeConfig:                   remoteChainCfg.TokenPoolConfig.TokenTransferFeeConfig,
			DefaultFinalityInboundRateLimiterConfig:  remoteChainCfg.TokenPoolConfig.DefaultFinalityInboundRateLimiterConfig,
			DefaultFinalityOutboundRateLimiterConfig: remoteChainCfg.TokenPoolConfig.DefaultFinalityOutboundRateLimiterConfig,
			CustomFinalityInboundRateLimiterConfig:   remoteChainCfg.TokenPoolConfig.CustomFinalityInboundRateLimiterConfig,
			CustomFinalityOutboundRateLimiterConfig:  remoteChainCfg.TokenPoolConfig.CustomFinalityOutboundRateLimiterConfig,
		},
		FeeUSDCents:         remoteChainCfg.FeeUSDCents,
		GasForVerification:  remoteChainCfg.GasForVerification,
		PayloadSizeBytes:    remoteChainCfg.PayloadSizeBytes,
		LockOrBurnMechanism: remoteChainCfg.LockOrBurnMechanism,
		RemoteDomain: adapters.RemoteDomain[[]byte]{
			DomainIdentifier: remoteChainCfg.RemoteDomain.DomainIdentifier,
		},
	}

	if addressRefExists(remoteChainCfg.RemoteDomain.AllowedCallerOnDest) {
		allowedCallerOnDest, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteChainCfg.RemoteDomain.AllowedCallerOnDest, remoteChainSelector, adapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve AllowedCallerOnDest for remote chain selector %d: %w", remoteChainSelector, err)
		}
		out.RemoteDomain.AllowedCallerOnDest = allowedCallerOnDest
	}

	if addressRefExists(remoteChainCfg.RemoteDomain.AllowedCallerOnSource) {
		allowedCallerOnSource, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteChainCfg.RemoteDomain.AllowedCallerOnSource, remoteChainSelector, adapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve AllowedCallerOnSource for remote chain selector %d: %w", remoteChainSelector, err)
		}
		out.RemoteDomain.AllowedCallerOnSource = allowedCallerOnSource
	}

	if addressRefExists(remoteChainCfg.RemoteDomain.MintRecipientOnDest) {
		mintRecipientOnDest, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteChainCfg.RemoteDomain.MintRecipientOnDest, remoteChainSelector, adapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve MintRecipientOnDest for remote chain selector %d: %w", remoteChainSelector, err)
		}
		out.RemoteDomain.MintRecipientOnDest = mintRecipientOnDest
	}

	if addressRefExists(remoteChainCfg.TokenPoolConfig.RemotePool) {
		remotePool, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteChainCfg.TokenPoolConfig.RemotePool, remoteChainSelector, adapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve RemotePool for remote chain selector %d: %w", remoteChainSelector, err)
		}
		out.TokenPoolConfig.RemotePool = remotePool
	}

	if addressRefExists(remoteChainCfg.TokenPoolConfig.RemoteToken) {
		remoteToken, err := datastore_utils.FindAndFormatRef(e.DataStore, remoteChainCfg.TokenPoolConfig.RemoteToken, remoteChainSelector, adapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve RemoteToken for remote chain selector %d: %w", remoteChainSelector, err)
		}
		out.TokenPoolConfig.RemoteToken = remoteToken
	}

	for _, ccvRef := range remoteChainCfg.TokenPoolConfig.OutboundCCVs {
		ccv, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, localChainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve OutboundCCV for remote chain selector %d: %w", localChainSelector, err)
		}
		out.TokenPoolConfig.OutboundCCVs = append(out.TokenPoolConfig.OutboundCCVs, ccv.Address)
	}

	for _, ccvRef := range remoteChainCfg.TokenPoolConfig.InboundCCVs {
		ccv, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, localChainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve InboundCCV for remote chain selector %d: %w", localChainSelector, err)
		}
		out.TokenPoolConfig.InboundCCVs = append(out.TokenPoolConfig.InboundCCVs, ccv.Address)
	}

	for _, ccvRef := range remoteChainCfg.TokenPoolConfig.OutboundCCVsToAddAboveThreshold {
		ccv, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, localChainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve OutboundCCVToAddAboveThreshold for remote chain selector %d: %w", localChainSelector, err)
		}
		out.TokenPoolConfig.OutboundCCVsToAddAboveThreshold = append(out.TokenPoolConfig.OutboundCCVsToAddAboveThreshold, ccv.Address)
	}

	for _, ccvRef := range remoteChainCfg.TokenPoolConfig.InboundCCVsToAddAboveThreshold {
		ccv, err := datastore_utils.FindAndFormatRef(e.DataStore, ccvRef, localChainSelector, datastore_utils.FullRef)
		if err != nil {
			return adapters.RemoteCCTPChainConfig[string, []byte]{}, fmt.Errorf("failed to resolve InboundCCVToAddAboveThreshold for remote chain selector %d: %w", localChainSelector, err)
		}
		out.TokenPoolConfig.InboundCCVsToAddAboveThreshold = append(out.TokenPoolConfig.InboundCCVsToAddAboveThreshold, ccv.Address)
	}

	return out, nil
}

// addressRefExists checks if an AddressRef is populated.
func addressRefExists(ref datastore.AddressRef) bool {
	return ref.Address != "" ||
		ref.Type != "" ||
		ref.Version != nil ||
		ref.Qualifier != "" ||
		ref.ChainSelector != 0 ||
		ref.Labels.Length() != 0
}
