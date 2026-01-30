package cctp

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

const (
	localTokenDecimals = 6
)

var (
	cctpQualifier = "CCTP"

	// This sequence assumes that CCTP V2 pools are on version 1.6.4 and CCTP V1 pools on 1.6.2
	cctpV2PrevVersion  = semver.MustParse("1.6.4")
	cctpV1PrevVersion  = semver.MustParse("1.6.2")
	cctpV2ContractType = deployment.ContractType("USDCTokenPoolCCTPV2")
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

		existingAddresses, err := dep.DataStore.Addresses().Fetch()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to fetch all addresses: %w", err)
		}

		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		tokenMessengerAddress := common.HexToAddress(input.TokenMessenger)
		usdcTokenAddress := common.HexToAddress(input.USDCToken)
		feeAggregatorAddress := common.HexToAddress(input.FeeAggregator)
		create2FactoryAddress := common.HexToAddress(input.DeployerContract)

		rmnRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMN proxy ref on chain %d: %w", chain.Selector, err)
		}
		rmnAddress := common.HexToAddress(rmnRef.Address)

		routerRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: router.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find router ref on chain %d: %w", chain.Selector, err)
		}
		routerAddress := common.HexToAddress(routerRef.Address)

		// Deploy CCTPMessageTransmitterProxy if needed
		cctpMessageTransmitterProxyRef, err := contract_utils.MaybeDeployContract(b, cctp_message_transmitter_proxy.Deploy, chain, contract_utils.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_message_transmitter_proxy.ContractType, *cctp_message_transmitter_proxy.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: cctp_message_transmitter_proxy.ConstructorArgs{
				TokenMessenger: tokenMessengerAddress,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPMessageTransmitterProxy: %w", err)
		}
		addresses = append(addresses, cctpMessageTransmitterProxyRef)
		cctpMessageTransmitterProxyAddress := common.HexToAddress(cctpMessageTransmitterProxyRef.Address)

		// Deploy CCTPVerifier if needed
		cctpVerifierRef, err := contract_utils.MaybeDeployContract(b, cctp_verifier.Deploy, chain, contract_utils.DeployInput[cctp_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: cctp_verifier.ConstructorArgs{
				TokenMessenger:          tokenMessengerAddress,
				MessageTransmitterProxy: cctpMessageTransmitterProxyAddress,
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

		// Deploy CCTPVerifierResolver if needed
		cctpVerifierResolverRefs := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByType(datastore.ContractType(cctp_verifier.ResolverType)),
			datastore.AddressRefByVersion(cctp_verifier.Version),
			datastore.AddressRefByQualifier(cctpQualifier),
		)
		var cctpVerifierResolverRef datastore.AddressRef
		if len(cctpVerifierResolverRefs) == 0 {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.ChainSelector,
				Qualifier:      cctpQualifier,
				Type:           datastore.ContractType(cctp_verifier.ResolverType),
				Version:        cctp_verifier.Version,
				CREATE2Factory: create2FactoryAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
			}
			if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected 1 CCTPVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
			}
			cctpVerifierResolverRef = deployVerifierResolverViaCREATE2Report.Output.Addresses[0]
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
		} else if len(cctpVerifierResolverRefs) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected 0 or 1 CCTPVerifierResolver addresses, got %d", len(cctpVerifierResolverRefs))
		} else {
			cctpVerifierResolverRef = cctpVerifierResolverRefs[0]
			addresses = append(addresses, cctpVerifierResolverRef)
		}
		cctpVerifierResolverAddress := common.HexToAddress(cctpVerifierResolverRef.Address)

		// Deploy CCTPThroughCCVTokenPool if needed
		cctpTokenPoolRef, err := contract_utils.MaybeDeployContract(b, cctp_through_ccv_token_pool.Deploy, chain, contract_utils.DeployInput[cctp_through_ccv_token_pool.ConstructorArgs]{
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
		addresses = append(addresses, cctpTokenPoolRef)
		cctpTokenPoolAddress := common.HexToAddress(cctpTokenPoolRef.Address)

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: cctpTokenPoolAddress,
			RouterAddress:    routerAddress,
			FeeAggregator:    feeAggregatorAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		isHomeChain := chain.Selector == chain_selectors.ETHEREUM_MAINNET.Selector || chain.Selector == chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
		var siloedLockReleaseTokenPoolRef datastore.AddressRef
		if isHomeChain {
			ref, err := contract_utils.MaybeDeployContract(b, lock_release_token_pool.Deploy, chain, contract_utils.DeployInput[lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(lock_release_token_pool.SiloedUSDCTokenPoolContractType, *lock_release_token_pool.Version),
				ChainSelector:  chain.Selector,
				Args: lock_release_token_pool.ConstructorArgs{
					Token:              usdcTokenAddress,
					LocalTokenDecimals: localTokenDecimals,
					AdvancedPoolHooks:  common.Address{},
					RMNProxy:           rmnAddress,
					Router:             routerAddress,
				},
			}, existingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy SiloedUSDCTokenPool: %w", err)
			}
			siloedLockReleaseTokenPoolRef = ref
			addresses = append(addresses, ref)
		}
		siloedLockReleaseTokenPoolAddress := common.HexToAddress(siloedLockReleaseTokenPoolRef.Address)

		cctpV1PoolAddresses := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByType(datastore.ContractType(cctpV1ContractType)),
			datastore.AddressRefByVersion(cctpV1PrevVersion),
		)
		if len(cctpV1PoolAddresses) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected 0 or 1 CCTP V1 pool addresses, got %d", len(cctpV1PoolAddresses))
		} else if len(cctpV1PoolAddresses) == 0 {
			cctpV1PoolAddresses = append(cctpV1PoolAddresses, datastore.AddressRef{})
		}

		cctpV2PoolAddresses := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByType(datastore.ContractType(cctpV2ContractType)),
			datastore.AddressRefByVersion(cctpV2PrevVersion),
		)
		if len(cctpV2PoolAddresses) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected 0 or 1 CCTP V2 pool addresses, got %d", len(cctpV2PoolAddresses))
		} else if len(cctpV2PoolAddresses) == 0 {
			cctpV2PoolAddresses = append(cctpV2PoolAddresses, datastore.AddressRef{})
		}

		cctpV1PoolAddress := common.HexToAddress(cctpV1PoolAddresses[0].Address)
		cctpV2PoolAddress := common.HexToAddress(cctpV2PoolAddresses[0].Address)

		// Deploy USDCTokenPoolProxy
		usdcTokenPoolProxyRef, err := contract_utils.MaybeDeployContract(b, usdc_token_pool_proxy.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &cctpQualifier,
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token: usdcTokenAddress,
				Pools: usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{
					CctpV1Pool:            cctpV1PoolAddress,
					CctpV2Pool:            cctpV2PoolAddress,
					CctpV2PoolWithCCV:     cctpTokenPoolAddress,
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

		// Add CCTPVerifier as an authorized caller on the CCTPMessageTransmitterProxy
		verifierApplyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_message_transmitter_proxy.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpMessageTransmitterProxyAddress,
			Args: cctp_message_transmitter_proxy.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					cctpVerifierAddress,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to message transmitter proxy: %w", err)
		}
		writes = append(writes, verifierApplyAuthorizedCallerUpdatesReport.Output)

		// Add USDCTokenPoolProxy as an authorized caller on the CCTPTokenPool
		poolApplyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_through_ccv_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_through_ccv_token_pool.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpTokenPoolAddress,
			Args: cctp_through_ccv_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					usdcTokenPoolProxyAddress,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to CCTPTokenPool: %w", err)
		}
		writes = append(writes, poolApplyAuthorizedCallerUpdatesReport.Output)

		// Set inbound implementation on the CCTPVerifierResolver
		committeeVerifierVersionTagReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.GetVersionTag, chain, contract_utils.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       cctpVerifierAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CCTPVerifier: %w", err)
		}
		setInboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       cctpVerifierResolverAddress,
			Args: []versioned_verifier_resolver.InboundImplementationArgs{
				{
					Version:  committeeVerifierVersionTagReport.Output,
					Verifier: cctpVerifierAddress,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set inbound implementation on CCTPVerifierResolver: %w", err)
		}
		writes = append(writes, setInboundImplementationReport.Output)

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
	if !containsAddress(callersReport.Output, proxyAddr) {
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
