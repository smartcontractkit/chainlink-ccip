package cctp

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
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
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

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

var ConfigureCCTPChainForLanes = cldf_ops.NewSequence(
	"configure-cctp-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures the CCTP contracts on a chain for multiple remote chains",
	func(b cldf_ops.Bundle, dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

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

		usdcTokenPoolProxyAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
			Version: usdc_token_pool_proxy.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find USDCTokenPoolProxy ref on chain %d: %w", chain.Selector, err)
		}

		routerAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: router.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find Router ref on chain %d: %w", chain.Selector, err)
		}

		cctpVerifierAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(cctp_verifier.ContractType),
			Version: cctp_verifier.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTPVerifier ref on chain %d: %w", chain.Selector, err)
		}

		cctpVerifierResolverAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(cctp_verifier.ResolverType),
			Version: cctp_verifier.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTPVerifierResolver ref on chain %d: %w", chain.Selector, err)
		}

		cctpV2WithCCVsTokenPoolAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(cctp_through_ccv_token_pool.ContractType),
			Version: cctp_through_ccv_token_pool.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTPTokenPool ref on chain %d: %w", chain.Selector, err)
		}

		tokenAdminRegistryAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find TokenAdminRegistry ref on chain %d: %w", chain.Selector, err)
		}

		cctpV2TokenPoolAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(usdc_token_pool_cctp_v2.ContractType),
			Version: usdc_token_pool_cctp_v2.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTP V2 token pool ref on chain %d: %w", chain.Selector, err)
		}

		cctpV1TokenPoolAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(cctpV1ContractType),
			Version: cctpV1PrevVersion,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTP V1 token pool ref on chain %d: %w", chain.Selector, err)
		}

		// Configure each remote chain of the siloed USDC pool, if required.
		// This involves deploying a lockbox for each remote chain and configuring it on the pool.
		if isHomeChainAndConfigureSiloedPool {
			siloedUSDCAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(siloed_usdc_token_pool.ContractType),
				Version: siloed_usdc_token_pool.Version,
			}, chain.Selector, datastore_utils.FullRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find siloed USDC token pool ref on chain %d: %w", chain.Selector, err)
			}

			siloedLockReleaseReport, err := cldf_ops.ExecuteSequence(b, DeploySiloedUSDCLockRelease, dep.BlockChains, DeploySiloedUSDCLockReleaseInput{
				ChainSelector:             input.ChainSelector,
				USDCToken:                 input.USDCToken,
				SiloedUSDCTokenPool:       siloedUSDCAddressRef.Address,
				LockReleaseChainSelectors: lockReleaseSelectors,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy siloed USDC lock release stack: %w", err)
			}
			addresses = append(addresses, siloedLockReleaseReport.Output.Addresses...)
			batchOps = append(batchOps, siloedLockReleaseReport.Output.BatchOps...)
		}

		cctpVerifierAddress := common.HexToAddress(cctpVerifierAddressRef.Address)
		cctpVerifierResolverAddress := common.HexToAddress(cctpVerifierResolverAddressRef.Address)
		usdcTokenPoolProxyAddress := common.HexToAddress(usdcTokenPoolProxyAddressRef.Address)
		routerAddress := common.HexToAddress(routerAddressRef.Address)

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainSelectors := make([]uint64, 0)
		mechanisms := make([]uint8, 0)
		setDomainArgs := make([]cctp_verifier.SetDomainArgs, 0)
		cctpV2DomainUpdates := make([]usdc_token_pool_cctp_v2.DomainUpdate, 0)
		remoteChainConfigArgs := make([]cctp_verifier.RemoteChainConfigArgs, 0)
		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remotePoolAddress, err := dep.RemoteChains[remoteChainSelector].PoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pool address: %w", err)
			}
			remoteTokenAddress, err := dep.RemoteChains[remoteChainSelector].TokenAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token address: %w", err)
			}
			remoteChainConfigs[remoteChainSelector] = tokens_core.RemoteChainConfig[[]byte, string]{
				RemotePool:             common.LeftPadBytes(remotePoolAddress, 32),
				RemoteToken:            common.LeftPadBytes(remoteTokenAddress, 32),
				TokenTransferFeeConfig: remoteChain.TokenTransferFeeConfig,
				// CCTP does not use rate limiters, so we set them to 0 here.
				DefaultFinalityOutboundRateLimiterConfig: tokens_core.RateLimiterConfig{
					Capacity: big.NewInt(0),
					Rate:     big.NewInt(0),
				},
				CustomFinalityOutboundRateLimiterConfig: tokens_core.RateLimiterConfig{
					Capacity: big.NewInt(0),
					Rate:     big.NewInt(0),
				},
				DefaultFinalityInboundRateLimiterConfig: tokens_core.RateLimiterConfig{
					Capacity: big.NewInt(0),
					Rate:     big.NewInt(0),
				},
				CustomFinalityInboundRateLimiterConfig: tokens_core.RateLimiterConfig{
					Capacity: big.NewInt(0),
					Rate:     big.NewInt(0),
				},
			}
			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          cctpVerifierAddress,
			})
			remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)
			mechanism, err := convertMechanismToUint8(remoteChain.LockOrBurnMechanism)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lock or burn mechanism to uint8: %w", err)
			}
			mechanisms = append(mechanisms, mechanism)

			unpaddedAllowedCallerOnDest, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed caller on dest: %w", err)
			}
			allowedCallerOnDest := common.LeftPadBytes(unpaddedAllowedCallerOnDest, 32)

			unpaddedAllowedCallerOnSource, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnSource(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed caller on source: %w", err)
			}
			allowedCallerOnSource := common.LeftPadBytes(unpaddedAllowedCallerOnSource, 32)

			unpaddedMintRecipientOnDest, err := dep.RemoteChains[remoteChainSelector].MintRecipientOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get mint recipient on dest: %w", err)
			}
			mintRecipientOnDest := common.LeftPadBytes(unpaddedMintRecipientOnDest, 32)

			var allowedCallerOnDestBytes32 [32]byte
			copy(allowedCallerOnDestBytes32[32-len(allowedCallerOnDest):], allowedCallerOnDest)
			var allowedCallerOnSourceBytes32 [32]byte
			copy(allowedCallerOnSourceBytes32[32-len(allowedCallerOnSource):], allowedCallerOnSource)
			var mintRecipientOnDestBytes32 [32]byte
			copy(mintRecipientOnDestBytes32[32-len(mintRecipientOnDest):], mintRecipientOnDest)
			setDomainArgs = append(setDomainArgs, cctp_verifier.SetDomainArgs{
				AllowedCallerOnDest:   allowedCallerOnDestBytes32,
				AllowedCallerOnSource: allowedCallerOnSourceBytes32,
				MintRecipientOnDest:   mintRecipientOnDestBytes32,
				DomainIdentifier:      remoteChain.DomainIdentifier,
				Enabled:               true,
				ChainSelector:         remoteChainSelector,
			})
			cctpV2DomainUpdates = append(cctpV2DomainUpdates, usdc_token_pool_cctp_v2.DomainUpdate{
				AllowedCaller:                 allowedCallerOnDestBytes32,
				MintRecipient:                 mintRecipientOnDestBytes32,
				DomainIdentifier:              remoteChain.DomainIdentifier,
				DestChainSelector:             remoteChainSelector,
				Enabled:                       true,
				UseLegacySourcePoolDataFormat: false,
			})
			remoteChainConfigArgs = append(remoteChainConfigArgs, cctp_verifier.RemoteChainConfigArgs{
				Router:              routerAddress,
				RemoteChainSelector: remoteChainSelector,
				FeeUSDCents:         remoteChain.FeeUSDCents,
				GasForVerification:  remoteChain.GasForVerification,
				PayloadSizeBytes:    remoteChain.PayloadSizeBytes,
			})
		}

		// Set outbound implementation on the CCTPVerifierResolver for each remote chain
		setOutboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpVerifierResolverAddress,
			Args:          outboundImplementations,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set outbound implementation on CCTPVerifierResolver: %w", err)
		}
		writes = append(writes, setOutboundImplementationReport.Output)

		// Set lock or burn mechanism for each remote chain.
		if len(remoteChainSelectors) > 0 {
			updateLockOrBurnMechanismsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
				ChainSelector: chain.Selector,
				Address:       usdcTokenPoolProxyAddress,
				Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
					RemoteChainSelectors: remoteChainSelectors,
					Mechanisms:           mechanisms,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update lock or burn mechanisms on USDCTokenPoolProxy: %w", err)
			}
			writes = append(writes, updateLockOrBurnMechanismsReport.Output)
		}

		// Apply remote chain config updates on the CCTPVerifier
		applyRemoteChainConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]cctp_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpVerifierAddress,
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on CCTPVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		// Set each remote domain on the CCTPVerifier
		setDomainsReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.SetDomains, chain, contract_utils.FunctionInput[[]cctp_verifier.SetDomainArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpVerifierAddress,
			Args:          setDomainArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set domains on CCTPVerifier: %w", err)
		}
		writes = append(writes, setDomainsReport.Output)

		// Create batch operation from writes
		if len(writes) > 0 {
			batchOpFromWrites, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOpFromWrites)
		}

		// Set domains on the CCTP V2 token pool
		setCCTPV2DomainsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_cctp_v2.USDCTokenPoolSetDomains, chain, contract_utils.FunctionInput[[]usdc_token_pool_cctp_v2.DomainUpdate]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cctpV2TokenPoolAddressRef.Address),
			Args:          cctpV2DomainUpdates,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set domains on CCTP V2 token pool: %w", err)
		}
		writes = append(writes, setCCTPV2DomainsReport.Output)

		// Configure remote chains on the CCTP V2 token pool (1.6.x sequence)
		cctpV2TokenPoolAddress := common.HexToAddress(cctpV2TokenPoolAddressRef.Address)
		for remoteChainSelector, remoteChainConfig := range remoteChainConfigs {
			configureCCTPV2PoolReport, err := cldf_ops.ExecuteSequence(b, v1_6_1_tokens.ConfigureTokenPoolForRemoteChain, chain, v1_6_1_tokens.ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    cctpV2TokenPoolAddress,
				RemoteChainSelector: remoteChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CCTP V2 token pool for remote chain %d: %w", remoteChainSelector, err)
			}
			batchOps = append(batchOps, configureCCTPV2PoolReport.Output.BatchOps...)
		}

		// Configure remote chains on the CCTP V1 token pool (1.6.x sequence)
		cctpV1TokenPoolAddress := common.HexToAddress(cctpV1TokenPoolAddressRef.Address)
		for remoteChainSelector, remoteChainConfig := range remoteChainConfigs {
			configureCCTPV1PoolReport, err := cldf_ops.ExecuteSequence(b, v1_6_1_tokens.ConfigureTokenPoolForRemoteChain, chain, v1_6_1_tokens.ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    cctpV1TokenPoolAddress,
				RemoteChainSelector: remoteChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure CCTP V1 token pool for remote chain %d: %w", remoteChainSelector, err)
			}
			batchOps = append(batchOps, configureCCTPV1PoolReport.Output.BatchOps...)
		}

		// Call into configure token for transfers sequence for the CCTP token pool with CCVs
		// Configure token for transfers performs full registration, which we only need to do once.
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, dep.BlockChains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:            input.ChainSelector,
			TokenAddress:             input.USDCToken,
			TokenPoolAddress:         cctpV2WithCCVsTokenPoolAddressRef.Address,
			RegistryTokenPoolAddress: usdcTokenPoolProxyAddressRef.Address,
			RegistryAddress:          tokenAdminRegistryAddressRef.Address,
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
