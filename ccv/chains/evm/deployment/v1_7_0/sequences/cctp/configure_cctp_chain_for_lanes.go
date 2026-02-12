package cctp

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_6_1_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
)

const (
	mechanismCCTPV1        = "CCTP_V1"
	mechanismCCTPV2        = "CCTP_V2"
	mechanismLockRelease   = "LOCK_RELEASE"
	mechanismCCTPV2WithCCV = "CCTP_V2_WITH_CCV"
)

// configureCCTPChainRefs holds resolved address refs for ConfigureCCTPChainForLanes.
// CCTPV1TokenPool is optional; nil when the chain has no CCTP V1 pool deployed.
type configureCCTPChainRefs struct {
	USDCTokenPoolProxy   datastore.AddressRef
	Router               datastore.AddressRef
	CCTPVerifier         datastore.AddressRef
	CCTPVerifierResolver datastore.AddressRef
	CCTPV2WithCCVsPool   datastore.AddressRef
	TokenAdminRegistry   datastore.AddressRef
	CCTPV2TokenPool      datastore.AddressRef
	RegisteredPool       datastore.AddressRef
	CCTPV1TokenPool      *datastore.AddressRef
}

var ConfigureCCTPChainForLanes = cldf_ops.NewSequence(
	"configure-cctp-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures the CCTP contracts on a chain for multiple remote chains",
	func(b cldf_ops.Bundle, dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		// Resolve chain and validate
		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		lockReleaseSelectors := make([]uint64, 0)
		for sel, cfg := range input.RemoteChains {
			if cfg.LockOrBurnMechanism == mechanismLockRelease {
				lockReleaseSelectors = append(lockReleaseSelectors, sel)
			}
		}

		isHomeChain := chain.Selector == chain_selectors.ETHEREUM_MAINNET.Selector || chain.Selector == chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
		if !isHomeChain && len(lockReleaseSelectors) > 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("lock-release configuration is only supported on home chains")
		}
		isHomeChainAndConfigureSiloedPool := isHomeChain && len(lockReleaseSelectors) > 0

		// Resolve address refs
		refs, siloedUSDCRef, err := resolveConfigureCCTPChainRefs(dep.DataStore, chain.Selector, isHomeChainAndConfigureSiloedPool, input.RegisteredPoolRef)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Build remote chain configs (used by siloed deploy, token pools, and configure token for transfers)
		remoteChainConfigs, err := buildRemoteChainConfigs(dep, input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Siloed USDC pool (home chain + lock-release only)
		if isHomeChainAndConfigureSiloedPool {
			siloedRemoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string], len(lockReleaseSelectors))
			for _, sel := range lockReleaseSelectors {
				siloedRemoteChainConfigs[sel] = remoteChainConfigs[sel]
			}
			// Siloed USDC lock release will not be deployed here as it already exists.
			// One lockbox will be deployed per lock-release selector.
			siloedLockReleaseReport, err := cldf_ops.ExecuteSequence(b, DeploySiloedUSDCLockRelease, dep.BlockChains, DeploySiloedUSDCLockReleaseInput{
				ChainSelector:             input.ChainSelector,
				USDCToken:                 input.USDCToken,
				SiloedUSDCTokenPool:       siloedUSDCRef.Address,
				LockReleaseChainSelectors: lockReleaseSelectors,
				RemoteChainConfigs:        siloedRemoteChainConfigs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy siloed USDC lock release stack: %w", err)
			}
			addresses = append(addresses, siloedLockReleaseReport.Output.Addresses...)
			batchOps = append(batchOps, siloedLockReleaseReport.Output.BatchOps...)
		}

		cctpVerifierAddress := common.HexToAddress(refs.CCTPVerifier.Address)
		routerAddress := common.HexToAddress(refs.Router.Address)

		// CCTPVerifierResolver: set outbound implementation per remote chain
		outboundImpls := buildVerifierResolverOutboundArgs(input, cctpVerifierAddress)
		w, err := applyVerifierResolverOutboundWrites(b, chain, common.HexToAddress(refs.CCTPVerifierResolver.Address), outboundImpls)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, w...)

		// USDCTokenPoolProxy: lock/burn mechanism per remote chain
		remoteSelectors, mechanisms, err := buildUSDCTokenPoolProxyMechanismArgs(input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if len(remoteSelectors) > 0 {
			w, err = applyUSDCTokenPoolProxyMechanismWrites(b, chain, common.HexToAddress(refs.USDCTokenPoolProxy.Address), remoteSelectors, mechanisms)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			writes = append(writes, w...)
		}

		// CCTPVerifier: remote chain config (fee, gas, payload) and domain args
		verifierSetDomainArgs, verifierRemoteChainConfigArgs, err := buildCCTPVerifierArgs(dep, input, routerAddress, cctpVerifierAddress)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		w, err = applyCCTPVerifierWrites(b, chain, cctpVerifierAddress, verifierSetDomainArgs, verifierRemoteChainConfigArgs)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, w...)

		// CCTP V2 token pool: set domains
		cctpV2DomainUpdates, err := buildCCTPV2PoolDomainUpdates(dep, input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if len(cctpV2DomainUpdates) > 0 {
			w, err = applyCCTPV2PoolSetDomainsWrites(b, chain, common.HexToAddress(refs.CCTPV2TokenPool.Address), cctpV2DomainUpdates)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			writes = append(writes, w...)
		}

		// Create batch operation from writes
		if len(writes) > 0 {
			batchOpFromWrites, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOpFromWrites)
		}

		// Configure remote chains on CCTP V2 token pool (1.6.1 sequence)
		cctpV2TokenPoolAddress := common.HexToAddress(refs.CCTPV2TokenPool.Address)
		for remoteChainSelector, remoteChainConfig := range remoteChainConfigs {
			report, err := cldf_ops.ExecuteSequence(b, v1_6_1_tokens.ConfigureTokenPoolForRemoteChain, chain, v1_6_1_tokens.ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    cctpV2TokenPoolAddress,
				RemoteChainSelector: remoteChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CCTP V2 token pool for remote chain %d: %w", remoteChainSelector, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
		}

		// Configure remote chains on CCTP V1 token pool (1.6.1 sequence), if present
		if refs.CCTPV1TokenPool != nil {
			cctpV1TokenPoolAddress := common.HexToAddress(refs.CCTPV1TokenPool.Address)
			for remoteChainSelector, remoteChainConfig := range remoteChainConfigs {
				report, err := cldf_ops.ExecuteSequence(b, v1_6_1_tokens.ConfigureTokenPoolForRemoteChain, chain, v1_6_1_tokens.ConfigureTokenPoolForRemoteChainInput{
					ChainSelector:       input.ChainSelector,
					TokenPoolAddress:    cctpV1TokenPoolAddress,
					RemoteChainSelector: remoteChainSelector,
					RemoteChainConfig:   remoteChainConfig,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CCTP V1 token pool for remote chain %d: %w", remoteChainSelector, err)
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
			}
		}

		// Configure token for transfers (CCTP-through-CCV pool; registration is done once)
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, dep.BlockChains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:            input.ChainSelector,
			TokenAddress:             input.USDCToken,
			TokenPoolAddress:         refs.CCTPV2WithCCVsPool.Address,
			RegistryTokenPoolAddress: refs.RegisteredPool.Address,
			RegistryAddress:          refs.TokenAdminRegistry.Address,
			MinFinalityValue:         1,
			RemoteChains:             remoteChainConfigs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token for transfers: %w", err)
		}
		batchOps = append(batchOps, configureTokenForTransfersReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)

// resolveConfigureCCTPChainRefs resolves all address refs needed for ConfigureCCTPChainForLanes.
func resolveConfigureCCTPChainRefs(
	ds datastore.DataStore,
	chainSelector uint64,
	needSiloedUSDC bool,
	registeredPoolRef datastore.AddressRef,
) (configureCCTPChainRefs, *datastore.AddressRef, error) {
	refs := configureCCTPChainRefs{}
	var err error
	refs.USDCTokenPoolProxy, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
		Version: usdc_token_pool_proxy.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find USDCTokenPoolProxy ref on chain %d: %w", chainSelector, err)
	}
	refs.Router, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(router.ContractType),
		Version: router.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find Router ref on chain %d: %w", chainSelector, err)
	}
	refs.CCTPVerifier, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_verifier.ContractType),
		Version: cctp_verifier.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find CCTPVerifier ref on chain %d: %w", chainSelector, err)
	}
	refs.CCTPVerifierResolver, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_verifier.ResolverType),
		Version: cctp_verifier.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find CCTPVerifierResolver ref on chain %d: %w", chainSelector, err)
	}
	refs.CCTPV2WithCCVsPool, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_through_ccv_token_pool.ContractType),
		Version: cctp_through_ccv_token_pool.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find CCTPTokenPool ref on chain %d: %w", chainSelector, err)
	}
	refs.TokenAdminRegistry, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(token_admin_registry.ContractType),
		Version: token_admin_registry.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find TokenAdminRegistry ref on chain %d: %w", chainSelector, err)
	}
	refs.CCTPV2TokenPool, err = datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(usdc_token_pool_cctp_v2.ContractType),
		Version: usdc_token_pool_cctp_v2.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find CCTP V2 token pool ref on chain %d: %w", chainSelector, err)
	}
	refs.RegisteredPool, err = datastore_utils.FindAndFormatRef(ds, registeredPoolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return refs, nil, fmt.Errorf("failed to find RegisteredPool ref on chain %d: %w", chainSelector, err)
	}
	cctpV1PoolRefs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(cctpV1ContractType)),
		datastore.AddressRefByVersion(cctpV1PrevVersion),
	)
	if len(cctpV1PoolRefs) > 1 {
		return refs, nil, fmt.Errorf("expected 0 or 1 CCTP V1 token pool refs on chain %d, found %d", chainSelector, len(cctpV1PoolRefs))
	}
	if len(cctpV1PoolRefs) == 1 {
		refs.CCTPV1TokenPool = &cctpV1PoolRefs[0]
	} else {
		refs.CCTPV1TokenPool = nil
	}
	var siloedRef *datastore.AddressRef
	if needSiloedUSDC {
		siloed, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
			Type:    datastore.ContractType(siloed_usdc_token_pool.ContractType),
			Version: siloed_usdc_token_pool.Version,
		}, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return refs, nil, fmt.Errorf("failed to find siloed USDC token pool ref on chain %d: %w", chainSelector, err)
		}
		siloedRef = &siloed
	}
	return refs, siloedRef, nil
}

// buildRemoteChainConfigs builds the remote chain config map used by token pools and configure-token-for-transfers.
func buildRemoteChainConfigs(dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (map[uint64]tokens_core.RemoteChainConfig[[]byte, string], error) {
	configs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string], len(input.RemoteChains))
	for remoteChainSelector, remoteChain := range input.RemoteChains {
		remotePoolAddress, err := dep.RemoteChains[remoteChainSelector].PoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector, input.RemoteRegisteredPoolRefs[remoteChainSelector])
		if err != nil {
			return nil, fmt.Errorf("failed to get remote pool address: %w", err)
		}
		remoteTokenAddress, err := dep.RemoteChains[remoteChainSelector].TokenAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get remote token address: %w", err)
		}
		configs[remoteChainSelector] = tokens_core.RemoteChainConfig[[]byte, string]{
			RemotePool:                               common.LeftPadBytes(remotePoolAddress, 32),
			RemoteToken:                              common.LeftPadBytes(remoteTokenAddress, 32),
			TokenTransferFeeConfig:                   remoteChain.TokenTransferFeeConfig,
			DefaultFinalityOutboundRateLimiterConfig: tokens_core.RateLimiterConfig{Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			CustomFinalityOutboundRateLimiterConfig:  tokens_core.RateLimiterConfig{Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			DefaultFinalityInboundRateLimiterConfig:  tokens_core.RateLimiterConfig{Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			CustomFinalityInboundRateLimiterConfig:   tokens_core.RateLimiterConfig{Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		}
	}
	return configs, nil
}

// buildVerifierResolverOutboundArgs builds outbound implementation args for the CCTPVerifierResolver (one per remote chain).
func buildVerifierResolverOutboundArgs(input adapters.ConfigureCCTPChainForLanesInput, cctpVerifierAddress common.Address) []versioned_verifier_resolver.OutboundImplementationArgs {
	out := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0, len(input.RemoteChains))
	for remoteChainSelector := range input.RemoteChains {
		out = append(out, versioned_verifier_resolver.OutboundImplementationArgs{
			DestChainSelector: remoteChainSelector,
			Verifier:          cctpVerifierAddress,
		})
	}
	return out
}

// buildUSDCTokenPoolProxyMechanismArgs builds remote chain selectors and lock/burn mechanisms for the USDCTokenPoolProxy.
func buildUSDCTokenPoolProxyMechanismArgs(input adapters.ConfigureCCTPChainForLanesInput) (remoteChainSelectors []uint64, mechanisms []uint8, err error) {
	remoteChainSelectors = make([]uint64, 0, len(input.RemoteChains))
	mechanisms = make([]uint8, 0, len(input.RemoteChains))
	for remoteChainSelector, remoteChain := range input.RemoteChains {
		mechanism, err := convertMechanismToUint8(remoteChain.LockOrBurnMechanism)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert lock or burn mechanism to uint8: %w", err)
		}
		remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)
		mechanisms = append(mechanisms, mechanism)
	}
	return remoteChainSelectors, mechanisms, nil
}

// buildCCTPVerifierArgs builds set-domain args and remote-chain-config args for the CCTPVerifier.
// allowedCallerOnSource is the current chain's verifier (source chain when sending to remote).
func buildCCTPVerifierArgs(dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput, routerAddress common.Address, allowedCallerOnSource common.Address) ([]cctp_verifier.SetDomainArgs, []cctp_verifier.RemoteChainConfigArgs, error) {
	setDomainArgs := make([]cctp_verifier.SetDomainArgs, 0, len(input.RemoteChains))
	remoteChainConfigArgs := make([]cctp_verifier.RemoteChainConfigArgs, 0, len(input.RemoteChains))
	for remoteChainSelector, remoteChain := range input.RemoteChains {
		allowedCallerOnDest, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
		if err != nil {
			return nil, nil, err
		}
		mintRecipientOnDest, err := dep.RemoteChains[remoteChainSelector].MintRecipientOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
		if err != nil {
			return nil, nil, err
		}
		allowedCallerOnDest = common.LeftPadBytes(allowedCallerOnDest, 32)
		allowedCallerOnSourceBytes := common.LeftPadBytes(allowedCallerOnSource.Bytes(), 32)
		mintRecipientOnDest = common.LeftPadBytes(mintRecipientOnDest, 32)
		var allowedCallerOnDestBytes32, allowedCallerOnSourceBytes32, mintRecipientOnDestBytes32 [32]byte
		copy(allowedCallerOnDestBytes32[32-len(allowedCallerOnDest):], allowedCallerOnDest)
		copy(allowedCallerOnSourceBytes32[32-len(allowedCallerOnSourceBytes):], allowedCallerOnSourceBytes)
		copy(mintRecipientOnDestBytes32[32-len(mintRecipientOnDest):], mintRecipientOnDest)
		setDomainArgs = append(setDomainArgs, cctp_verifier.SetDomainArgs{
			AllowedCallerOnDest:   allowedCallerOnDestBytes32,
			AllowedCallerOnSource: allowedCallerOnSourceBytes32,
			MintRecipientOnDest:   mintRecipientOnDestBytes32,
			DomainIdentifier:      remoteChain.DomainIdentifier,
			Enabled:               true,
			ChainSelector:         remoteChainSelector,
		})
		remoteChainConfigArgs = append(remoteChainConfigArgs, cctp_verifier.RemoteChainConfigArgs{
			Router:              routerAddress,
			RemoteChainSelector: remoteChainSelector,
			FeeUSDCents:         remoteChain.FeeUSDCents,
			GasForVerification:  remoteChain.GasForVerification,
			PayloadSizeBytes:    remoteChain.PayloadSizeBytes,
		})
	}
	return setDomainArgs, remoteChainConfigArgs, nil
}

// buildCCTPV2PoolDomainUpdates builds domain updates for the CCTP V2 token pool.
func buildCCTPV2PoolDomainUpdates(dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) ([]usdc_token_pool_cctp_v2.DomainUpdate, error) {
	out := make([]usdc_token_pool_cctp_v2.DomainUpdate, 0, len(input.RemoteChains))
	for remoteChainSelector, remoteChain := range input.RemoteChains {
		allowedCallerOnDest, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get allowed caller on dest: %w", err)
		}
		mintRecipientOnDest, err := dep.RemoteChains[remoteChainSelector].MintRecipientOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get mint recipient on dest: %w", err)
		}
		allowedCallerOnDest = common.LeftPadBytes(allowedCallerOnDest, 32)
		mintRecipientOnDest = common.LeftPadBytes(mintRecipientOnDest, 32)
		var allowedCallerBytes32, mintRecipientBytes32 [32]byte
		copy(allowedCallerBytes32[32-len(allowedCallerOnDest):], allowedCallerOnDest)
		copy(mintRecipientBytes32[32-len(mintRecipientOnDest):], mintRecipientOnDest)
		out = append(out, usdc_token_pool_cctp_v2.DomainUpdate{
			AllowedCaller:                 allowedCallerBytes32,
			MintRecipient:                 mintRecipientBytes32,
			DomainIdentifier:              remoteChain.DomainIdentifier,
			DestChainSelector:             remoteChainSelector,
			Enabled:                       true,
			UseLegacySourcePoolDataFormat: false,
		})
	}
	return out, nil
}

// applyVerifierResolverOutboundWrites sets the outbound implementation on the CCTPVerifierResolver.
func applyVerifierResolverOutboundWrites(b cldf_ops.Bundle, chain evm.Chain, resolverAddress common.Address, args []versioned_verifier_resolver.OutboundImplementationArgs) ([]contract_utils.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
		ChainSelector: chain.Selector,
		Address:       resolverAddress,
		Args:          args,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set outbound implementation on CCTPVerifierResolver: %w", err)
	}
	return []contract_utils.WriteOutput{report.Output}, nil
}

// applyUSDCTokenPoolProxyMechanismWrites updates lock/burn mechanisms on the USDCTokenPoolProxy.
func applyUSDCTokenPoolProxyMechanismWrites(b cldf_ops.Bundle, chain evm.Chain, proxyAddress common.Address, remoteChainSelectors []uint64, mechanisms []uint8) ([]contract_utils.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
		ChainSelector: chain.Selector,
		Address:       proxyAddress,
		Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
			RemoteChainSelectors: remoteChainSelectors,
			Mechanisms:           mechanisms,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update lock or burn mechanisms on USDCTokenPoolProxy: %w", err)
	}
	return []contract_utils.WriteOutput{report.Output}, nil
}

// applyCCTPVerifierWrites applies remote chain config and set-domains on the CCTPVerifier.
func applyCCTPVerifierWrites(b cldf_ops.Bundle, chain evm.Chain, verifierAddress common.Address, setDomainArgs []cctp_verifier.SetDomainArgs, remoteChainConfigArgs []cctp_verifier.RemoteChainConfigArgs) ([]contract_utils.WriteOutput, error) {
	writes := make([]contract_utils.WriteOutput, 0)
	remoteConfigReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]cctp_verifier.RemoteChainConfigArgs]{
		ChainSelector: chain.Selector,
		Address:       verifierAddress,
		Args:          remoteChainConfigArgs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply remote chain config updates on CCTPVerifier: %w", err)
	}
	writes = append(writes, remoteConfigReport.Output)
	domainsReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.SetDomains, chain, contract_utils.FunctionInput[[]cctp_verifier.SetDomainArgs]{
		ChainSelector: chain.Selector,
		Address:       verifierAddress,
		Args:          setDomainArgs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set domains on CCTPVerifier: %w", err)
	}
	writes = append(writes, domainsReport.Output)
	return writes, nil
}

// applyCCTPV2PoolSetDomainsWrites sets domains on the CCTP V2 token pool.
func applyCCTPV2PoolSetDomainsWrites(b cldf_ops.Bundle, chain evm.Chain, poolAddress common.Address, domainUpdates []usdc_token_pool_cctp_v2.DomainUpdate) ([]contract_utils.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_cctp_v2.SetDomains, chain, contract_utils.FunctionInput[[]usdc_token_pool_cctp_v2.DomainUpdate]{
		ChainSelector: chain.Selector,
		Address:       poolAddress,
		Args:          domainUpdates,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set domains on CCTP V2 token pool: %w", err)
	}
	return []contract_utils.WriteOutput{report.Output}, nil
}

func convertMechanismToUint8(mechanism string) (uint8, error) {
	switch mechanism {
	case mechanismCCTPV1:
		return 1, nil
	case mechanismCCTPV2:
		return 2, nil
	case mechanismLockRelease:
		return 3, nil
	case mechanismCCTPV2WithCCV:
		return 4, nil
	default:
		return 0, fmt.Errorf("invalid mechanism, must be %s, %s, %s, or %s: %s", mechanismCCTPV1, mechanismCCTPV2, mechanismLockRelease, mechanismCCTPV2WithCCV, mechanism)
	}
}
