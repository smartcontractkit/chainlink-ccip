package cctp

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	cctp_message_transmitter_proxy_v1_6_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool_cctp_v2"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
)

const (
	localTokenDecimals                = 6
	cctpV1SupportedUSDCVersion uint32 = 0
)

var (
	cctpQualifier = "CCTP"

	// CCTP V1 uses the USDCTokenPool v1.6.5 implementation.
	cctpV1Version      = semver.MustParse("1.6.5")
	cctpV1ContractType = deployment.ContractType("USDCTokenPool")
)

var DeployCCTPChain = cldf_ops.NewSequence(
	"deploy-cctp-chain",
	semver.MustParse("1.7.0"),
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

		// Deploy/resolve CCTPMessageTransmitterProxy (v1.7.0 op) for CCTPVerifier + CCTP V2 pool wiring.
		cctpV2MessageTransmitterProxyRef, err := contract_utils.MaybeDeployContract(b, cctp_message_transmitter_proxy.Deploy, chain, contract_utils.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_message_transmitter_proxy.ContractType, *cctp_message_transmitter_proxy.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: cctp_message_transmitter_proxy.ConstructorArgs{
				TokenMessenger: tokenMessengerV2Address,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPMessageTransmitterProxy v1.7.0: %w", err)
		}
		addresses = append(addresses, cctpV2MessageTransmitterProxyRef)
		cctpV2MessageTransmitterProxyAddress := common.HexToAddress(cctpV2MessageTransmitterProxyRef.Address)

		// Deploy CCTPVerifier if needed
		cctpVerifierRef, err := contract_utils.MaybeDeployContract(b, cctp_verifier.Deploy, chain, contract_utils.DeployInput[cctp_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: cctp_verifier.ConstructorArgs{
				TokenMessenger:          tokenMessengerV2Address,
				MessageTransmitterProxy: cctpV2MessageTransmitterProxyAddress,
				USDCToken:               usdcTokenAddress,
				StorageLocations:        input.StorageLocations,
				DynamicConfig: cctp_verifier.DynamicConfig{
					FeeAggregator:   feeAggregatorAddress,
					FastFinalityBps: input.FastFinalityBps,
				},
				RMN: rmnAddress,
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
		cctpV2WithCCVsTokenPoolRef, err := contract_utils.MaybeDeployContract(b, cctp_through_ccv_token_pool.Deploy, chain, contract_utils.DeployInput[cctp_through_ccv_token_pool.ConstructorArgs]{
			ChainSelector:  chain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_through_ccv_token_pool.ContractType, *cctp_through_ccv_token_pool.Version),
			Qualifier:      &cctpQualifier,
			Args: cctp_through_ccv_token_pool.ConstructorArgs{
				Token:              usdcTokenAddress,
				LocalTokenDecimals: localTokenDecimals,
				RMNProxy:           rmnAddress,
				Router:             routerAddress,
				CCTPVerifier:       cctpVerifierAddress,
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
			// Only the siloed USDC lock release pool will be deployed here.
			// Lockboxes are deployed per lane and are therefore deployed during ConfigureCCTPChainForLanes.
			siloedLockReleaseReport, err := cldf_ops.ExecuteSequence(b, DeploySiloedUSDCLockRelease, dep.BlockChains, DeploySiloedUSDCLockReleaseInput{
				ChainSelector: input.ChainSelector,
				USDCToken:     input.USDCToken,
				RMN:           rmnAddress.Hex(),
				Router:        routerAddress.Hex(),
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
			cctpV1PoolAddressRef, err := contract_utils.MaybeDeployContract(b, usdc_token_pool.Deploy, chain, contract_utils.DeployInput[usdc_token_pool.ConstructorArgs]{
				ChainSelector: input.ChainSelector,
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
		cctpV2TokenPoolAddressRef, err := contract_utils.MaybeDeployContract(b, usdc_token_pool_cctp_v2.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_cctp_v2.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
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
		usdcTokenPoolProxyRef, err := contract_utils.MaybeDeployContract(b, usdc_token_pool_proxy.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token: usdcTokenAddress,
				Pools: usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{
					CctpV1Pool:            cctpV1PoolAddress,
					CctpV2Pool:            cctpV2TokenPoolAddress,
					CctpV2PoolWithCCV:     cctpV2WithCCVsTokenPoolAddress,
					SiloedLockReleasePool: siloedLockReleaseTokenPoolAddress,
				},
				Router:       routerAddress,
				CCTPVerifier: cctpVerifierResolverAddress,
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
		setFeeAggregatorReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.SetFeeAggregator, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.SetFeeAggregatorArgs]{
			ChainSelector: chain.Selector,
			Address:       usdcTokenPoolProxyAddress,
			Args: usdc_token_pool_proxy.SetFeeAggregatorArgs{
				FeeAggregator: feeAggregatorAddress,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set fee aggregator on USDCTokenPoolProxy: %w", err)
		}
		writes = append(writes, setFeeAggregatorReport.Output)

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
	// Get authorized callers on siloed pool.
	callersReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       siloedPoolAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers from siloed pool: %w", err)
	}
	// Authorize proxy if not already authorized.
	if !slices.Contains(callersReport.Output, proxyAddr) {
		poolAuthReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[siloed_usdc_token_pool.AuthorizedCallerArgs]{
			ChainSelector: chainSelector,
			Address:       siloedPoolAddr,
			Args: siloed_usdc_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{proxyAddr},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to authorize proxy on siloed pool: %w", err)
		}
		writes = append(writes, poolAuthReport.Output)
	}
	// Get current pools from proxy.
	poolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.GetPools, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       proxyAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get existing proxy pools: %w", err)
	}
	currentPools := poolsReport.Output
	// If siloed pool address is not set correctly, update it.
	if currentPools.SiloedLockReleasePool != siloedPoolAddr {
		updatePoolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdatePoolAddresses, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.PoolAddresses]{
			ChainSelector: chainSelector,
			Address:       proxyAddr,
			Args: usdc_token_pool_proxy.PoolAddresses{
				CctpV1Pool:            currentPools.CctpV1Pool,
				CctpV2Pool:            currentPools.CctpV2Pool,
				CctpV2PoolWithCCV:     currentPools.CctpV2PoolWithCCV,
				SiloedLockReleasePool: siloedPoolAddr,
			},
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
	writes *[]contract_utils.WriteOutput,
) (datastore.AddressRef, error) {
	refs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chain.Selector),
		datastore.AddressRefByType(datastore.ContractType(cctp_verifier.ResolverType)),
		datastore.AddressRefByVersion(cctp_verifier.Version),
		datastore.AddressRefByQualifier(cctpQualifier),
	)
	if len(refs) == 0 {
		if create2FactoryAddress == (common.Address{}) {
			return datastore.AddressRef{}, fmt.Errorf("deployer contract is required")
		}
		report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
			ChainSelector:  chainSelector,
			Qualifier:      cctpQualifier,
			Type:           datastore.ContractType(cctp_verifier.ResolverType),
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
) ([]contract_utils.WriteOutput, error) {
	writes := make([]contract_utils.WriteOutput, 0)
	if enableCCTPV1 {
		v1MsgTxReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy_v1_6_2.ConfigureAllowedCallers, chain, contract_utils.FunctionInput[[]cctp_message_transmitter_proxy_v1_6_2.AllowedCallerConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpV1MessageTransmitterProxyAddr,
			Args: []cctp_message_transmitter_proxy_v1_6_2.AllowedCallerConfigArgs{
				{Caller: cctpVerifierAddr, Allowed: true},
				{Caller: cctpV1PoolAddr, Allowed: true},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to configure allowed callers on CCTPMessageTransmitterProxy v1.6.2: %w", err)
		}
		writes = append(writes, v1MsgTxReport.Output)
	}

	v2MsgTxReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_message_transmitter_proxy.AuthorizedCallerArgs]{
		ChainSelector: chain.Selector,
		Address:       cctpV2MessageTransmitterProxyAddr,
		Args: cctp_message_transmitter_proxy.AuthorizedCallerArgs{
			AddedCallers: []common.Address{cctpVerifierAddr, cctpV2PoolAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPMessageTransmitterProxy v1.7.0: %w", err)
	}
	writes = append(writes, v2MsgTxReport.Output)

	if enableCCTPV1 {
		cctpV1TokenPoolReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[usdc_token_pool.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpV1PoolAddr,
			Args: usdc_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{proxyAddr},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPool: %w", err)
		}
		writes = append(writes, cctpV1TokenPoolReport.Output)
	}

	cctpV2ThroughCCVTokenPoolReport, err := cldf_ops.ExecuteOperation(b, cctp_through_ccv_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_through_ccv_token_pool.AuthorizedCallerArgs]{
		ChainSelector: chain.Selector,
		Address:       cctpV2WithCCVsPoolAddr,
		Args: cctp_through_ccv_token_pool.AuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply authorized caller updates to CCTPThroughCCVTokenPool: %w", err)
	}
	writes = append(writes, cctpV2ThroughCCVTokenPoolReport.Output)

	cctpV2TokenPoolReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_cctp_v2.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[usdc_token_pool_cctp_v2.AuthorizedCallerArgs]{
		ChainSelector: chain.Selector,
		Address:       cctpV2PoolAddr,
		Args: usdc_token_pool_cctp_v2.AuthorizedCallerArgs{
			AddedCallers: []common.Address{proxyAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply authorized caller updates to USDCTokenPoolCCTPV2: %w", err)
	}
	writes = append(writes, cctpV2TokenPoolReport.Output)
	return writes, nil
}

// setCCTPVerifierResolverInbound sets the CCTPVerifier as the inbound implementation on the resolver. Returns writes to append.
func setCCTPVerifierResolverInbound(
	b cldf_ops.Bundle,
	chain evm.Chain,
	cctpVerifierAddr, resolverAddr common.Address,
) ([]contract_utils.WriteOutput, error) {
	versionTagReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.GetVersionTag, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       cctpVerifierAddr,
	})
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
