package cctp

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	cmtp162ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool_cctp_v2"
	cmtp162bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/cctp_message_transmitter_proxy"
	cmtpv2bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_message_transmitter_proxy"
	ccvtpbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_through_ccv_token_pool"
	cvbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_verifier"
	sutpbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_usdc_token_pool"
	utppbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/usdc_token_pool_proxy"
	utpbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_5/usdc_token_pool"
	utpv2bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_5/usdc_token_pool_cctp_v2"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/usdc_token_pool_proxy"
	v2_0_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
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
		writes := make([]contract_utils.WriteOutput, 0)
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
		cctpV2MessageTransmitterProxyRef, err := ops2contract.MaybeDeployContract(b, cctp_message_transmitter_proxy.Deploy, chain, ops2contract.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
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
		cctpVerifierRef, err := ops2contract.MaybeDeployContract(b, cctp_verifier.Deploy, chain, ops2contract.DeployInput[cctp_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version),
			Args: cctp_verifier.ConstructorArgs{
				TokenMessenger:          tokenMessengerV2Address,
				MessageTransmitterProxy: cctpV2MessageTransmitterProxyAddress,
				UsdcToken:               usdcTokenAddress,
				DynamicConfig: cvbind.CCTPVerifierDynamicConfig{
					FeeAggregator:   feeAggregatorAddress,
					FastFinalityBps: input.FastFinalityBps,
				},
				BaseVerifierArgs: cvbind.CCTPVerifierBaseVerifierArgs{
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
		cctpV2WithCCVsTokenPoolRef, err := ops2contract.MaybeDeployContract(b, cctp_through_ccv_token_pool.Deploy, chain, ops2contract.DeployInput[cctp_through_ccv_token_pool.ConstructorArgs]{
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
			cctpV1PoolAddressRef, err := ops2contract.MaybeDeployContract(b, usdc_token_pool.Deploy, chain, ops2contract.DeployInput[usdc_token_pool.ConstructorArgs]{
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
		cctpV2TokenPoolAddressRef, err := ops2contract.MaybeDeployContract(b, usdc_token_pool_cctp_v2.Deploy, chain, ops2contract.DeployInput[usdc_token_pool_cctp_v2.ConstructorArgs]{
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
		usdcTokenPoolProxyRef, err := ops2contract.MaybeDeployContract(b, usdc_token_pool_proxy.Deploy, chain, ops2contract.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token: usdcTokenAddress,
				Pools: utppbind.USDCTokenPoolProxyPoolAddresses{
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

		// Set the fee aggregator on the USDCTokenPoolProxy
		usdcProxy, err := utppbind.NewUSDCTokenPoolProxy(usdcTokenPoolProxyAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to bind USDCTokenPoolProxy at %s: %w", usdcTokenPoolProxyAddress.Hex(), err)
		}
		setFeeAggregatorReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.NewWriteSetFeeAggregator(usdcProxy), chain, ops2contract.FunctionInput[common.Address]{
			Args: feeAggregatorAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set fee aggregator on USDCTokenPoolProxy: %w", err)
		}
		writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(setFeeAggregatorReport.Output))

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

		chainBatchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
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
) ([]contract_utils.WriteOutput, error) {
	writes := make([]contract_utils.WriteOutput, 0)
	siloedPool, err := sutpbind.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind SiloedUSDCTokenPool at %s: %w", siloedPoolAddr.Hex(), err)
	}
	callersReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.NewReadGetAllAuthorizedCallers(siloedPool), chain, ops2contract.FunctionInput[struct{}]{})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from siloed pool: %w", err)
	}
	if !slices.Contains(callersReport.Output, proxyAddr) {
		poolAuthReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.NewWriteApplyAuthorizedCallerUpdates(siloedPool), chain, ops2contract.FunctionInput[sutpbind.AuthorizedCallersAuthorizedCallerArgs]{
			Args: sutpbind.AuthorizedCallersAuthorizedCallerArgs{
				AddedCallers: []common.Address{proxyAddr},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to authorize proxy on siloed pool: %w", err)
		}
		writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(poolAuthReport.Output))
	}
	usdcProxy, err := utppbind.NewUSDCTokenPoolProxy(proxyAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind USDCTokenPoolProxy at %s: %w", proxyAddr.Hex(), err)
	}
	poolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.NewReadGetPools(usdcProxy), chain, ops2contract.FunctionInput[struct{}]{})
	if err != nil {
		return nil, fmt.Errorf("failed to get existing proxy pools: %w", err)
	}
	currentPools := poolsReport.Output
	if currentPools.SiloedLockReleasePool != siloedPoolAddr {
		updatePoolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.NewWriteUpdatePoolAddresses(usdcProxy), chain, ops2contract.FunctionInput[utppbind.USDCTokenPoolProxyPoolAddresses]{
			Args: utppbind.USDCTokenPoolProxyPoolAddresses{
				CctpV1Pool:            currentPools.CctpV1Pool,
				CctpV2Pool:            currentPools.CctpV2Pool,
				CctpV2PoolWithCCV:     currentPools.CctpV2PoolWithCCV,
				SiloedLockReleasePool: siloedPoolAddr,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update proxy pool addresses: %w", err)
		}
		writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(updatePoolsReport.Output))
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
	writes *[]contract_utils.WriteOutput,
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
			ref.Type == datastore.ContractType(cmtp162ops.ContractType) &&
			ref.Version != nil &&
			ref.Version.Equal(cmtp162ops.Version) {
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
) ([]contract_utils.WriteOutput, error) {
	writes := make([]contract_utils.WriteOutput, 0)
	if enableCCTPV1 {
		cmtpV1, err := cmtp162bind.NewCCTPMessageTransmitterProxy(cctpV1MessageTransmitterProxyAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to bind CCTPMessageTransmitterProxy v1.6.2 at %s: %w", cctpV1MessageTransmitterProxyAddr.Hex(), err)
		}
		v1CurrentReport, err := cldf_ops.ExecuteOperation(b, cmtp162ops.NewReadGetAllowedCallers(cmtpV1), chain, ops2contract.FunctionInput[struct{}]{})
		if err != nil {
			return nil, fmt.Errorf("failed to get allowed callers from CCTPMessageTransmitterProxy v1.6.2: %w", err)
		}
		v1Current := v1CurrentReport.Output
		toAddV1 := make([]cmtp162bind.CCTPMessageTransmitterProxyAllowedCallerConfigArgs, 0)
		for _, caller := range []common.Address{cctpVerifierAddr, cctpV1PoolAddr} {
			if !slices.Contains(v1Current, caller) {
				toAddV1 = append(toAddV1, cmtp162bind.CCTPMessageTransmitterProxyAllowedCallerConfigArgs{Caller: caller, Allowed: true})
			}
		}
		if len(toAddV1) > 0 {
			v1MsgTxReport, err := cldf_ops.ExecuteOperation(b, cmtp162ops.NewWriteConfigureAllowedCallers(cmtpV1), chain, ops2contract.FunctionInput[[]cmtp162bind.CCTPMessageTransmitterProxyAllowedCallerConfigArgs]{
				Args: toAddV1,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to configure allowed callers on CCTPMessageTransmitterProxy v1.6.2: %w", err)
			}
			writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(v1MsgTxReport.Output))
		}
	}

	cmtp, err := cmtpv2bind.NewCCTPMessageTransmitterProxy(cctpV2MessageTransmitterProxyAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CCTPMessageTransmitterProxy v2.0.0 at %s: %w", cctpV2MessageTransmitterProxyAddr.Hex(), err)
	}
	v2CurrentReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.NewReadGetAllAuthorizedCallers(cmtp), chain, ops2contract.FunctionInput[struct{}]{})
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
		v2MsgTxReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.NewWriteApplyAuthorizedCallerUpdates(cmtp), chain, ops2contract.FunctionInput[cmtpv2bind.AuthorizedCallersAuthorizedCallerArgs]{
			Args: cmtpv2bind.AuthorizedCallersAuthorizedCallerArgs{
				AddedCallers: toAddV2,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPMessageTransmitterProxy v2.0.0: %w", err)
		}
		writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(v2MsgTxReport.Output))
	}

	if enableCCTPV1 {
		cctpV1Pool, err := utpbind.NewUSDCTokenPool(cctpV1PoolAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to bind USDCTokenPool at %s: %w", cctpV1PoolAddr.Hex(), err)
		}
		cctpV1TokenPoolReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool.NewWriteApplyAuthorizedCallerUpdates(cctpV1Pool), chain, ops2contract.FunctionInput[utpbind.AuthorizedCallersAuthorizedCallerArgs]{
			Args: utpbind.AuthorizedCallersAuthorizedCallerArgs{
				AddedCallers: []common.Address{proxyAddr},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPool: %w", err)
		}
		writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(cctpV1TokenPoolReport.Output))
	}

	ccvPool, err := ccvtpbind.NewCCTPThroughCCVTokenPool(cctpV2WithCCVsPoolAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CCTPThroughCCVTokenPool at %s: %w", cctpV2WithCCVsPoolAddr.Hex(), err)
	}
	cctpV2ThroughCCVTokenPoolReport, err := cldf_ops.ExecuteOperation(b, cctp_through_ccv_token_pool.NewWriteApplyAuthorizedCallerUpdates(ccvPool), chain, ops2contract.FunctionInput[ccvtpbind.AuthorizedCallersAuthorizedCallerArgs]{
		Args: ccvtpbind.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPThroughCCVTokenPool: %w", err)
	}
	writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(cctpV2ThroughCCVTokenPoolReport.Output))

	cctpV2Pool, err := utpv2bind.NewUSDCTokenPoolCCTPV2(cctpV2PoolAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind USDCTokenPoolCCTPV2 at %s: %w", cctpV2PoolAddr.Hex(), err)
	}
	cctpV2TokenPoolReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_cctp_v2.NewWriteApplyAuthorizedCallerUpdates(cctpV2Pool), chain, ops2contract.FunctionInput[utpv2bind.AuthorizedCallersAuthorizedCallerArgs]{
		Args: utpv2bind.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPoolCCTPV2: %w", err)
	}
	writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(cctpV2TokenPoolReport.Output))
	return writes, nil
}

// setCCTPVerifierResolverInbound sets the CCTPVerifier as the inbound implementation on the resolver. Returns writes to append.
func setCCTPVerifierResolverInbound(
	b cldf_ops.Bundle,
	chain evm.Chain,
	cctpVerifierAddr, resolverAddr common.Address,
) ([]contract_utils.WriteOutput, error) {
	cv, err := cvbind.NewCCTPVerifier(cctpVerifierAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CCTPVerifier at %s: %w", cctpVerifierAddr.Hex(), err)
	}
	versionTagReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.NewReadVersionTag(cv), chain, ops2contract.FunctionInput[struct{}]{})
	if err != nil {
		return nil, fmt.Errorf("failed to get version tag from CCTPVerifier: %w", err)
	}
	report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
		ChainSelector: chain.Selector,
		Address:       resolverAddr,
		Args: []versioned_verifier_resolver.InboundImplementationArgs{
			{Version: versionTagReport.Output, Verifier: cctpVerifierAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set inbound implementation on CCTPVerifierResolver: %w", err)
	}
	return []contract_utils.WriteOutput{report.Output}, nil
}
