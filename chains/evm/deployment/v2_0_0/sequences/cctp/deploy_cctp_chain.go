package cctp

import (
	"fmt"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	cmtp162_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/cctp_message_transmitter_proxy"
	cmtp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_message_transmitter_proxy"
	ctctp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_through_ccv_token_pool"
	cv_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_verifier"
	sutp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_usdc_token_pool"
	utpp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/usdc_token_pool_proxy"
	vvr_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	cctp_message_transmitter_proxy_v1_6_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool_cctp_v2"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/usdc_token_pool_proxy"
	v2_0_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/verifier_tags"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
)

const (
	cctpV1SupportedUSDCVersion uint32 = 0
)

var (
	// CCTP V1 uses the USDCTokenPool v1.6.5 implementation.
	cctpV1Version      = semver.MustParse("1.6.5")
	cctpV1ContractType = deployment.ContractType("USDCTokenPool")
)

var DeployCCTPChain = cldf_ops.NewSequence(
	"deploy-cctp-chain",
	semver.MustParse("2.0.0"),
	"Deploys & configures the CCTP contracts on a chain",
	func(b cldf_ops.Bundle, dep adapters.DeployCCTPChainDeps, input adapters.DeployCCTPInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		// Resolve chain and existing addresses
		existingAddresses := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
		)
		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		tokenMessengerV2Address := common.HexToAddress(input.TokenMessengerV2)
		usdcTokenAddress := common.HexToAddress(input.USDCToken)
		feeAggregatorAddress := common.HexToAddress(input.FeeAggregator)
		create2FactoryAddress := common.HexToAddress(input.DeployerContract)
		enableCCTPV1 := input.TokenMessengerV1 != ""

		rmnAddress, routerAddress, err := resolveBaseChainRefs(dep.DataStore, chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Deploy core CCTP contracts
		var cctpV1MessageTransmitterProxyAddress common.Address
		if enableCCTPV1 {
			// CCTP V1 must use existing CCTPMessageTransmitterProxy v1.6.2.
			_, cctpV1ProxyAddressRef, err := resolveExistingCCTPV1MessageTransmitterProxy(
				b,
				chain,
				existingAddresses,
			)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			cctpV1MessageTransmitterProxyAddress = common.HexToAddress(cctpV1ProxyAddressRef.Address)
			addresses = append(addresses, cctpV1ProxyAddressRef)
		}

		// Deploy/resolve CCTPMessageTransmitterProxy (v2.0.0 op) for CCTPVerifier + CCTP V2 pool wiring.
		cctpV2MessageTransmitterProxyRef, err := evmops.MaybeDeployContract(b, cctp_message_transmitter_proxy.Deploy, chain, contract.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_message_transmitter_proxy.ContractType, *cctp_message_transmitter_proxy.Version),
			Args: cctp_message_transmitter_proxy.ConstructorArgs{
				TokenMessenger: tokenMessengerV2Address,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPMessageTransmitterProxy v2.0.0: %w", err)
		}
		addresses = append(addresses, cctpV2MessageTransmitterProxyRef)
		cctpV2MessageTransmitterProxyAddress := common.HexToAddress(cctpV2MessageTransmitterProxyRef.Address)

		// Deploy CCTPVerifier if needed
		cctpVerifierRef, err := evmops.MaybeDeployContract(b, cctp_verifier.Deploy, chain, contract.DeployInput[cctp_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version),
			Args: cctp_verifier.ConstructorArgs{
				TokenMessenger:          tokenMessengerV2Address,
				MessageTransmitterProxy: cctpV2MessageTransmitterProxyAddress,
				UsdcToken:               usdcTokenAddress,
				DynamicConfig: cv_bindings.CCTPVerifierDynamicConfig{
					FeeAggregator:   feeAggregatorAddress,
					FastFinalityBps: input.FastFinalityBps,
				},
				BaseVerifierArgs: cv_bindings.CCTPVerifierBaseVerifierArgs{
					StorageLocations: input.StorageLocations,
					Rmn:              rmnAddress,
					VersionTag:       verifier_tags.CCTPVerifierV2(),
				},
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPVerifier: %w", err)
		}
		addresses = append(addresses, cctpVerifierRef)
		cctpVerifierAddress := common.HexToAddress(cctpVerifierRef.Address)

		cctpVerifierResolverRef, err := deployOrResolveCCTPVerifierResolver(b, dep.DataStore, chain, input.ChainSelector, create2FactoryAddress, &addresses, &writes)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		cctpVerifierResolverAddress := common.HexToAddress(cctpVerifierResolverRef.Address)

		// Deploy token pools
		// Deploy CCTPThroughCCVTokenPool if needed
		cctpV2WithCCVsTokenPoolRef, err := evmops.MaybeDeployContract(b, cctp_through_ccv_token_pool.Deploy, chain, contract.DeployInput[cctp_through_ccv_token_pool.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_through_ccv_token_pool.ContractType, *cctp_through_ccv_token_pool.Version),
			Args: cctp_through_ccv_token_pool.ConstructorArgs{
				Token:              usdcTokenAddress,
				LocalTokenDecimals: input.TokenDecimals,
				RmnProxy:           rmnAddress,
				Router:             routerAddress,
				CctpVerifier:       cctpVerifierResolverAddress,
				AllowedCallers:     []common.Address{},
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPTokenPool: %w", err)
		}
		addresses = append(addresses, cctpV2WithCCVsTokenPoolRef)
		cctpV2WithCCVsTokenPoolAddress := common.HexToAddress(cctpV2WithCCVsTokenPoolRef.Address)

		isHomeChain := chain.Selector == chain_selectors.ETHEREUM_MAINNET.Selector || chain.Selector == chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
		var siloedLockReleaseTokenPoolRef datastore.AddressRef
		if isHomeChain {
			// Check if SiloedUSDCTokenPool is already present in the datastore,
			// otherwise the sequence will deploy a fresh pool. Lockboxes are deployed per
			// lane during ConfigureCCTPChainForLanes, not here.
			var existingSiloedPoolAddr string
			existingSiloedRefs := dep.DataStore.Addresses().Filter(
				datastore.AddressRefByChainSelector(input.ChainSelector),
				datastore.AddressRefByType(datastore.ContractType(siloed_usdc_token_pool.ContractType)),
				datastore.AddressRefByVersion(siloed_usdc_token_pool.Version),
			)
			if len(existingSiloedRefs) > 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected at most one SiloedUSDCTokenPool v%s on chain %d, found %d", siloed_usdc_token_pool.Version.String(), input.ChainSelector, len(existingSiloedRefs))
			}
			if len(existingSiloedRefs) == 1 {
				existingSiloedPoolAddr = existingSiloedRefs[0].Address
			}
			siloedLockReleaseReport, err := cldf_ops.ExecuteSequence(b, DeploySiloedUSDCLockRelease, dep.BlockChains, DeploySiloedUSDCLockReleaseInput{
				ChainSelector:       input.ChainSelector,
				TokenDecimals:       input.TokenDecimals,
				USDCToken:           input.USDCToken,
				RMN:                 rmnAddress.Hex(),
				Router:              routerAddress.Hex(),
				SiloedUSDCTokenPool: existingSiloedPoolAddr,
				ExistingAddresses:   existingAddresses,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy siloed USDC lock release stack: %w", err)
			}
			siloedLockReleaseTokenPoolRef.Address = siloedLockReleaseReport.Output.SiloedPoolAddress
			addresses = append(addresses, siloedLockReleaseReport.Output.Addresses...)
			batchOps = append(batchOps, siloedLockReleaseReport.Output.BatchOps...)
		}
		siloedLockReleaseTokenPoolAddress := common.HexToAddress(siloedLockReleaseTokenPoolRef.Address)

		// Deploy USDCTokenPool (CCTP V1 pool) if needed.
		var cctpV1PoolAddress common.Address
		if enableCCTPV1 {
			tokenMessengerV1Address := common.HexToAddress(input.TokenMessengerV1)
			cctpV1PoolAddressRef, err := evmops.MaybeDeployContract(b, usdc_token_pool.DeployV2, chain, contract.DeployInput[usdc_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(
					usdc_token_pool.ContractType,
					*usdc_token_pool.Version,
				),
				Args: usdc_token_pool.ConstructorArgs{
					TokenMessenger:              tokenMessengerV1Address,
					CctpMessageTransmitterProxy: cctpV1MessageTransmitterProxyAddress,
					Token:                       usdcTokenAddress,
					Allowlist:                   []common.Address{},
					RmnProxy:                    rmnAddress,
					Router:                      routerAddress,
					SupportedUSDCVersion:        cctpV1SupportedUSDCVersion,
				},
			}, existingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPool on %s: %w", chain, err)
			}
			addresses = append(addresses, cctpV1PoolAddressRef)
			cctpV1PoolAddress = common.HexToAddress(cctpV1PoolAddressRef.Address)
		}

		// Deploy USDCTokenPoolCCTPV2 (CCTP V2 pool)
		cctpV2TokenPoolAddressRef, err := evmops.MaybeDeployContract(b, usdc_token_pool_cctp_v2.DeployV2, chain, contract.DeployInput[usdc_token_pool_cctp_v2.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(
				usdc_token_pool_cctp_v2.ContractType,
				*usdc_token_pool_cctp_v2.Version,
			),
			Args: usdc_token_pool_cctp_v2.ConstructorArgs{
				TokenMessenger:              tokenMessengerV2Address,
				CctpMessageTransmitterProxy: cctpV2MessageTransmitterProxyAddress,
				Token:                       usdcTokenAddress,
				Allowlist:                   []common.Address{},
				RmnProxy:                    rmnAddress,
				Router:                      routerAddress,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolCCTPV2 on %s: %w", chain, err)
		}
		addresses = append(addresses, cctpV2TokenPoolAddressRef)
		cctpV2TokenPoolAddress := common.HexToAddress(cctpV2TokenPoolAddressRef.Address)

		// Deploy USDCTokenPoolProxy
		usdcTokenPoolProxyRef, err := evmops.MaybeDeployContract(b, usdc_token_pool_proxy.Deploy, chain, contract.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token: usdcTokenAddress,
				Pools: utpp_bindings.USDCTokenPoolProxyPoolAddresses{
					CctpV1Pool:            cctpV1PoolAddress,
					CctpV2Pool:            cctpV2TokenPoolAddress,
					CctpV2PoolWithCCV:     cctpV2WithCCVsTokenPoolAddress,
					SiloedLockReleasePool: siloedLockReleaseTokenPoolAddress,
				},
				Router:       routerAddress,
				CctpVerifier: cctpVerifierResolverAddress,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy: %w", err)
		}
		addresses = append(addresses, usdcTokenPoolProxyRef)
		usdcTokenPoolProxyAddress := common.HexToAddress(usdcTokenPoolProxyRef.Address)

		// Configure proxy and authorized callers
		if isHomeChain {
			siloedPoolWrites, err := configureSiloedPoolProxyWiring(b, chain, input.ChainSelector, usdcTokenPoolProxyAddress, siloedLockReleaseTokenPoolAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure siloed pool proxy wiring: %w", err)
			}
			writes = append(writes, siloedPoolWrites...)
		}

		// Set the fee aggregator on the USDCTokenPoolProxy, if not already set to the desired address.
		currentFeeAggregatorReport, err := evmops.ExecuteRead(b, chain, usdcTokenPoolProxyAddress, evmops.BindAs[utpp_bindings.USDCTokenPoolProxyInterface](utpp_bindings.NewUSDCTokenPoolProxy), usdc_token_pool_proxy.NewReadGetFeeAggregator, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee aggregator from USDCTokenPoolProxy: %w", err)
		}
		if currentFeeAggregatorReport.Output != feeAggregatorAddress {
			setFeeAggregatorReport, err := evmops.ExecuteWrite(b, chain, usdcTokenPoolProxyAddress, evmops.BindAs[utpp_bindings.USDCTokenPoolProxyInterface](utpp_bindings.NewUSDCTokenPoolProxy), usdc_token_pool_proxy.NewWriteSetFeeAggregator, feeAggregatorAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set fee aggregator on USDCTokenPoolProxy: %w", err)
			}
			writes = append(writes, setFeeAggregatorReport.Output)
		}

		authorizedCallerWrites, err := applyCCTPAuthorizedCallerWrites(
			b,
			chain,
			enableCCTPV1,
			cctpV1MessageTransmitterProxyAddress,
			cctpV2MessageTransmitterProxyAddress,
			cctpVerifierAddress,
			cctpV1PoolAddress,
			cctpV2TokenPoolAddress,
			cctpV2WithCCVsTokenPoolAddress,
			usdcTokenPoolProxyAddress,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, authorizedCallerWrites...)

		inboundImplWrites, err := setCCTPVerifierResolverInbound(b, chain, cctpVerifierAddress, cctpVerifierResolverAddress)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, inboundImplWrites...)

		chainBatchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build batch operation from writes: %w", err)
		}
		if len(chainBatchOp.Transactions) > 0 {
			batchOps = append(batchOps, chainBatchOp)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)

func configureSiloedPoolProxyWiring(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	proxyAddr common.Address,
	siloedPoolAddr common.Address,
) ([]contract.WriteOutput, error) {
	writes := make([]contract.WriteOutput, 0)
	callersReport, err := evmops.ExecuteRead(b, chain, siloedPoolAddr, evmops.BindAs[sutp_bindings.SiloedUSDCTokenPoolInterface](sutp_bindings.NewSiloedUSDCTokenPool), siloed_usdc_token_pool.NewReadGetAllAuthorizedCallers, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from siloed pool: %w", err)
	}
	if !slices.Contains(callersReport.Output, proxyAddr) {
		poolAuthReport, err := evmops.ExecuteWrite(b, chain, siloedPoolAddr, evmops.BindAs[sutp_bindings.SiloedUSDCTokenPoolInterface](sutp_bindings.NewSiloedUSDCTokenPool), siloed_usdc_token_pool.NewWriteApplyAuthorizedCallerUpdates, sutp_bindings.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to authorize proxy on siloed pool: %w", err)
		}
		writes = append(writes, poolAuthReport.Output)
	}
	poolsReport, err := evmops.ExecuteRead(b, chain, proxyAddr, evmops.BindAs[utpp_bindings.USDCTokenPoolProxyInterface](utpp_bindings.NewUSDCTokenPoolProxy), usdc_token_pool_proxy.NewReadGetPools, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get existing proxy pools: %w", err)
	}
	currentPools := poolsReport.Output
	if currentPools.SiloedLockReleasePool != siloedPoolAddr {
		updatePoolsReport, err := evmops.ExecuteWrite(b, chain, proxyAddr, evmops.BindAs[utpp_bindings.USDCTokenPoolProxyInterface](utpp_bindings.NewUSDCTokenPoolProxy), usdc_token_pool_proxy.NewWriteUpdatePoolAddresses, utpp_bindings.USDCTokenPoolProxyPoolAddresses{
			CctpV1Pool:            currentPools.CctpV1Pool,
			CctpV2Pool:            currentPools.CctpV2Pool,
			CctpV2PoolWithCCV:     currentPools.CctpV2PoolWithCCV,
			SiloedLockReleasePool: siloedPoolAddr,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update proxy pool addresses: %w", err)
		}
		writes = append(writes, updatePoolsReport.Output)
	}

	return writes, nil
}

// resolveBaseChainRefs resolves RMN and Router address refs for the given chain.
func resolveBaseChainRefs(ds datastore.DataStore, chainSelector uint64) (rmnAddress, routerAddress common.Address, err error) {
	rmnRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmn_proxy.ContractType),
		Version: rmn_proxy.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to find RMN proxy ref on chain %d: %w", chainSelector, err)
	}
	routerRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(router.ContractType),
		Version: router.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to find router ref on chain %d: %w", chainSelector, err)
	}
	return common.HexToAddress(rmnRef.Address), common.HexToAddress(routerRef.Address), nil
}

// deployOrResolveCCTPVerifierResolver deploys the CCTPVerifierResolver via CREATE2 if not present, or reuses the existing one. Appends to addresses and writes.
func deployOrResolveCCTPVerifierResolver(
	b cldf_ops.Bundle,
	ds datastore.DataStore,
	chain evm.Chain,
	chainSelector uint64,
	create2FactoryAddress common.Address,
	addresses *[]datastore.AddressRef,
	writes *[]contract.WriteOutput,
) (datastore.AddressRef, error) {
	refs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chain.Selector),
		datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.CCTPVerifierResolverType)),
		datastore.AddressRefByVersion(cctp_verifier.Version),
	)
	if len(refs) == 0 {
		if create2FactoryAddress == (common.Address{}) {
			return datastore.AddressRef{}, fmt.Errorf("deployer contract is required")
		}
		report, err := cldf_ops.ExecuteSequence(b, v2_0_0_sequences.DeployVerifierResolverViaCREATE2, chain, v2_0_0_sequences.DeployVerifierResolverViaCREATE2Input{
			ChainSelector:  chainSelector,
			Type:           datastore.ContractType(versioned_verifier_resolver.CCTPVerifierResolverType),
			Version:        cctp_verifier.Version,
			CREATE2Factory: create2FactoryAddress,
		})
		if err != nil {
			return datastore.AddressRef{}, fmt.Errorf("failed to deploy CCTPVerifierResolver: %w", err)
		}
		if len(report.Output.Addresses) != 1 {
			return datastore.AddressRef{}, fmt.Errorf("expected 1 CCTPVerifierResolver address, got %d", len(report.Output.Addresses))
		}
		*addresses = append(*addresses, report.Output.Addresses...)
		*writes = append(*writes, report.Output.Writes...)
		return report.Output.Addresses[0], nil
	}
	if len(refs) > 1 {
		return datastore.AddressRef{}, fmt.Errorf("expected 0 or 1 CCTPVerifierResolver addresses, got %d", len(refs))
	}
	*addresses = append(*addresses, refs[0])
	return refs[0], nil
}

func resolveExistingCCTPV1MessageTransmitterProxy(
	_ cldf_ops.Bundle,
	chain evm.Chain,
	existingAddresses []datastore.AddressRef,
) (common.Address, datastore.AddressRef, error) {
	var legacyRef *datastore.AddressRef
	for i := range existingAddresses {
		ref := existingAddresses[i]
		if ref.ChainSelector == chain.Selector &&
			ref.Type == datastore.ContractType(cctp_message_transmitter_proxy_v1_6_2.ContractType) &&
			ref.Version != nil &&
			ref.Version.Equal(cctp_message_transmitter_proxy_v1_6_2.Version) {
			if legacyRef != nil {
				return common.Address{}, datastore.AddressRef{}, fmt.Errorf("expected exactly 1 CCTPMessageTransmitterProxy v1.6.2 ref on chain %d, found multiple", chain.Selector)
			}
			copy := ref
			legacyRef = &copy
		}
	}
	if legacyRef != nil {
		return common.HexToAddress(legacyRef.Address), *legacyRef, nil
	}
	return common.Address{}, datastore.AddressRef{}, fmt.Errorf("missing required CCTPMessageTransmitterProxy v1.6.2 on chain %d for CCTP V1 pool deployment", chain.Selector)
}

// applyCCTPAuthorizedCallerWrites adds CCTPVerifier + CCTP pools to message transmitter proxy and authorizes proxy on CCTP pools. Returns writes to append.
func applyCCTPAuthorizedCallerWrites(
	b cldf_ops.Bundle,
	chain evm.Chain,
	enableCCTPV1 bool,
	cctpV1MessageTransmitterProxyAddr, cctpV2MessageTransmitterProxyAddr, cctpVerifierAddr, cctpV1PoolAddr, cctpV2PoolAddr, cctpV2WithCCVsPoolAddr, proxyAddr common.Address,
) ([]contract.WriteOutput, error) {
	writes := make([]contract.WriteOutput, 0)
	if enableCCTPV1 {
		v1CurrentReport, err := evmops.ExecuteRead(b, chain, cctpV1MessageTransmitterProxyAddr, evmops.BindAs[cmtp162_bindings.CCTPMessageTransmitterProxyInterface](cmtp162_bindings.NewCCTPMessageTransmitterProxy), cctp_message_transmitter_proxy_v1_6_2.NewReadGetAllowedCallers, struct{}{})
		if err != nil {
			return nil, fmt.Errorf("failed to get allowed callers from CCTPMessageTransmitterProxy v1.6.2: %w", err)
		}
		v1Current := v1CurrentReport.Output
		toAddV1 := make([]cmtp162_bindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs, 0)
		for _, caller := range []common.Address{cctpVerifierAddr, cctpV1PoolAddr} {
			if !slices.Contains(v1Current, caller) {
				toAddV1 = append(toAddV1, cmtp162_bindings.CCTPMessageTransmitterProxyAllowedCallerConfigArgs{Caller: caller, Allowed: true})
			}
		}
		if len(toAddV1) > 0 {
			v1MsgTxReport, err := evmops.ExecuteWrite(b, chain, cctpV1MessageTransmitterProxyAddr, evmops.BindAs[cmtp162_bindings.CCTPMessageTransmitterProxyInterface](cmtp162_bindings.NewCCTPMessageTransmitterProxy), cctp_message_transmitter_proxy_v1_6_2.NewWriteConfigureAllowedCallers, toAddV1)
			if err != nil {
				return nil, fmt.Errorf("failed to configure allowed callers on CCTPMessageTransmitterProxy v1.6.2: %w", err)
			}
			writes = append(writes, v1MsgTxReport.Output)
		}
	}

	v2CurrentReport, err := evmops.ExecuteRead(b, chain, cctpV2MessageTransmitterProxyAddr, evmops.BindAs[cmtp_bindings.CCTPMessageTransmitterProxyInterface](cmtp_bindings.NewCCTPMessageTransmitterProxy), cctp_message_transmitter_proxy.NewReadGetAllAuthorizedCallers, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from CCTPMessageTransmitterProxy v2.0.0: %w", err)
	}
	v2Current := v2CurrentReport.Output
	toAddV2 := make([]common.Address, 0)
	for _, caller := range []common.Address{cctpVerifierAddr, cctpV2PoolAddr} {
		if !slices.Contains(v2Current, caller) {
			toAddV2 = append(toAddV2, caller)
		}
	}
	if len(toAddV2) > 0 {
		v2MsgTxReport, err := evmops.ExecuteWrite(b, chain, cctpV2MessageTransmitterProxyAddr, evmops.BindAs[cmtp_bindings.CCTPMessageTransmitterProxyInterface](cmtp_bindings.NewCCTPMessageTransmitterProxy), cctp_message_transmitter_proxy.NewWriteApplyAuthorizedCallerUpdates, cmtp_bindings.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: toAddV2,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPMessageTransmitterProxy v2.0.0: %w", err)
		}
		writes = append(writes, v2MsgTxReport.Output)
	}

	if enableCCTPV1 {
		cctpV1PoolCurrentReport, err := evmops.ExecuteRead(b, chain, cctpV1PoolAddr, usdc_token_pool.NewUSDCTokenPoolContract, usdc_token_pool.NewReadGetAllAuthorizedCallers, struct{}{})
		if err != nil {
			return nil, fmt.Errorf("failed to get authorized callers from USDCTokenPool: %w", err)
		}
		if !slices.Contains(cctpV1PoolCurrentReport.Output, proxyAddr) {
			cctpV1TokenPoolReport, err := evmops.ExecuteWrite(b, chain, cctpV1PoolAddr, usdc_token_pool.NewUSDCTokenPoolContract, usdc_token_pool.NewWriteApplyAuthorizedCallerUpdates, usdc_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{proxyAddr},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPool: %w", err)
			}
			writes = append(writes, cctpV1TokenPoolReport.Output)
		}
	}

	cctpV2ThroughCCVCurrentReport, err := evmops.ExecuteRead(b, chain, cctpV2WithCCVsPoolAddr, evmops.BindAs[ctctp_bindings.CCTPThroughCCVTokenPoolInterface](ctctp_bindings.NewCCTPThroughCCVTokenPool), cctp_through_ccv_token_pool.NewReadGetAllAuthorizedCallers, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from CCTPThroughCCVTokenPool: %w", err)
	}
	if !slices.Contains(cctpV2ThroughCCVCurrentReport.Output, proxyAddr) {
		cctpV2ThroughCCVTokenPoolReport, err := evmops.ExecuteWrite(b, chain, cctpV2WithCCVsPoolAddr, evmops.BindAs[ctctp_bindings.CCTPThroughCCVTokenPoolInterface](ctctp_bindings.NewCCTPThroughCCVTokenPool), cctp_through_ccv_token_pool.NewWriteApplyAuthorizedCallerUpdates, ctctp_bindings.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPThroughCCVTokenPool: %w", err)
		}
		writes = append(writes, cctpV2ThroughCCVTokenPoolReport.Output)
	}

	cctpV2PoolCurrentReport, err := evmops.ExecuteRead(b, chain, cctpV2PoolAddr, usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2Contract, usdc_token_pool_cctp_v2.NewReadGetAllAuthorizedCallers, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from USDCTokenPoolCCTPV2: %w", err)
	}
	if !slices.Contains(cctpV2PoolCurrentReport.Output, proxyAddr) {
		cctpV2TokenPoolReport, err := evmops.ExecuteWrite(b, chain, cctpV2PoolAddr, usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2Contract, usdc_token_pool_cctp_v2.NewWriteApplyAuthorizedCallerUpdates, usdc_token_pool_cctp_v2.AuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPoolCCTPV2: %w", err)
		}
		writes = append(writes, cctpV2TokenPoolReport.Output)
	}
	return writes, nil
}

// setCCTPVerifierResolverInbound sets the CCTPVerifier as the inbound implementation on the resolver,
// skipping the write if the on-chain state already matches.
func setCCTPVerifierResolverInbound(
	b cldf_ops.Bundle,
	chain evm.Chain,
	cctpVerifierAddr, resolverAddr common.Address,
) ([]contract.WriteOutput, error) {
	versionTagReport, err := evmops.ExecuteRead(b, chain, cctpVerifierAddr, evmops.BindAs[cv_bindings.CCTPVerifierInterface](cv_bindings.NewCCTPVerifier), cctp_verifier.NewReadVersionTag, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get version tag from CCTPVerifier: %w", err)
	}
	desired := versioned_verifier_resolver.InboundImplementationArgs{
		Version:  versionTagReport.Output,
		Verifier: cctpVerifierAddr,
	}
	currentReport, err := evmops.ExecuteRead(b, chain, resolverAddr, vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewReadGetAllInboundImplementations, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get inbound implementations from CCTPVerifierResolver: %w", err)
	}
	for _, cur := range currentReport.Output {
		if cur.Version == desired.Version && cur.Verifier == desired.Verifier {
			return nil, nil
		}
	}
	report, err := evmops.ExecuteWrite(b, chain, resolverAddr, vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewWriteApplyInboundImplementationUpdates, []versioned_verifier_resolver.InboundImplementationArgs{desired})
	if err != nil {
		return nil, fmt.Errorf("failed to set inbound implementation on CCTPVerifierResolver: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
