package lombard

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_token_pool"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
)

var ConfigureLombardChainForLanes = cldf_ops.NewSequence(
	"configure-lombard-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures the Lombard chain to support CCIP lanes",
	func(b cldf_ops.Bundle, dep adapters.ConfigureLombardChainForLanesDeps, input adapters.ConfigureLombardChainForLanesInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		lombardVerifierAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(lombard_verifier.ContractType),
			Version: lombard_verifier.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTPVerifier ref on chain %d: %w", chain.Selector, err)
		}

		lombardVerifierResolverAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(lombard_verifier.ResolverType),
			Version: lombard_verifier.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find CCTPVerifierResolver ref on chain %d: %w", chain.Selector, err)
		}

		routerRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: router.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find router ref on chain %d: %w", chain.Selector, err)
		}

		advancedPoolHooksRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:      datastore.ContractType(advanced_pool_hooks.ContractType),
			Version:   advanced_pool_hooks.Version,
			Qualifier: *tokenPoolQualifier(input.TokenQualifier),
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find AdvancedPoolHooks ref on chain %d: %w", chain.Selector, err)
		}

		tokenPoolRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:      datastore.ContractType(lombard_token_pool.ContractType),
			Version:   lombard_token_pool.Version,
			Qualifier: *tokenPoolQualifier(input.TokenQualifier),
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find AdvancedPoolHooks ref on chain %d: %w", chain.Selector, err)
		}

		tokenAdminRegistryAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find TokenAdminRegistry ref on chain %d: %w", chain.Selector, err)
		}

		advancedPoolHooksAddress := common.HexToAddress(advancedPoolHooksRef.Address)
		routerAddress := common.HexToAddress(routerRef.Address)
		lombardVerifierAddress := common.HexToAddress(lombardVerifierAddressRef.Address)
		lombardVerifierResolverAddress := common.HexToAddress(lombardVerifierResolverAddressRef.Address)

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainSelectors := make([]uint64, 0)
		remoteChainConfigArgs := make([]lombard_verifier.RemoteChainConfigArgs, 0)
		advancedPoolHooks := make([]advanced_pool_hooks.CCVConfigArg, 0)
		advancedPoolHooks = append(advancedPoolHooks, advanced_pool_hooks.CCVConfigArg{
			RemoteChainSelector: chain.Selector,
			InboundCCVs: []common.Address{
				lombardVerifierAddress,
				{}, // This means "require the default CCV(s) for this lane".
			},
		})

		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remotePoolAddress, err := dep.RemoteChains[remoteChainSelector].RemoteTokenPoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector, *tokenPoolQualifier(input.TokenQualifier))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pool address: %w", err)
			}
			remoteTokenAddress, err := dep.RemoteChains[remoteChainSelector].RemoteTokenAddress(b, dep.DataStore, dep.BlockChains, remoteChainSelector, *tokenPoolQualifier(input.TokenQualifier))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token address: %w", err)
			}

			remoteChainConfigs[remoteChainSelector] = tokens_core.RemoteChainConfig[[]byte, string]{
				RemotePool:             common.LeftPadBytes(remotePoolAddress, 32),
				RemoteToken:            common.LeftPadBytes(remoteTokenAddress, 32),
				TokenTransferFeeConfig: remoteChain.TokenTransferFeeConfig,
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
				Verifier:          lombardVerifierAddress,
			})

			remoteChainConfigArgs = append(remoteChainConfigArgs, lombard_verifier.RemoteChainConfigArgs{
				Router:              routerAddress,
				RemoteChainSelector: remoteChainSelector,
				FeeUSDCents:         remoteChain.FeeUSDCents,
				GasForVerification:  remoteChain.GasForVerification,
				PayloadSizeBytes:    remoteChain.PayloadSizeBytes,
			})

			advancedPoolHooks = append(advancedPoolHooks, advanced_pool_hooks.CCVConfigArg{
				RemoteChainSelector: remoteChainSelector,
				OutboundCCVs: []common.Address{
					lombardVerifierAddress,
					{}, // This means "require the default CCV(s) for this lane".
				},
			})

			remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)

			remoteCallerOnDest, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed caller on dest for remote chain %d: %w", remoteChainSelector, err)
			}

			allowedCaller, err := toBytes32LeftPad(remoteCallerOnDest)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller to bytes32: %w", err)
			}

			lchainID, err := paddedLombardChainID(remoteChain.LombardChainId)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lombardChainID to bytes32: %w", err)
			}

			setRemotePathReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.SetRemotePath, chain, contract_utils.FunctionInput[lombard_verifier.RemotePathArgs]{
				ChainSelector: chain.Selector,
				Address:       lombardVerifierAddress,
				Args: lombard_verifier.RemotePathArgs{
					RemoteChainSelector: remoteChainSelector,
					LChainId:            lchainID,
					AllowedCaller:       allowedCaller,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set remote path on LombardVerifier for remote chain %d: %w", remoteChainSelector, err)
			}
			writes = append(writes, setRemotePathReport.Output)
		}

		// Set outbound implementation on the LombardVerifierResolver for each remote chain
		setOutboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       lombardVerifierResolverAddress,
			Args:          outboundImplementations,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set outbound implementation on LombardVerifierResolver: %w", err)
		}
		writes = append(writes, setOutboundImplementationReport.Output)

		// Apply remote chain config updates on the LombardVerifier
		applyRemoteChainConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]lombard_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       lombardVerifierAddress,
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on LombardVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		// Apply advanced pool hooks CCV config updates
		advancedPoolHooksApplyReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyCCVConfigUpdates, chain, contract_utils.FunctionInput[[]advanced_pool_hooks.CCVConfigArg]{
			Address:       advancedPoolHooksAddress,
			ChainSelector: chain.Selector,
			Args:          advancedPoolHooks,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply advanced pool hooks CCV config updates: %w", err)
		}
		writes = append(writes, advancedPoolHooksApplyReport.Output)

		// Call into configure token for transfers sequence
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, dep.BlockChains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:    input.ChainSelector,
			TokenAddress:     input.Token,
			TokenPoolAddress: tokenPoolRef.Address,
			RegistryAddress:  tokenAdminRegistryAddressRef.Address,
			MinFinalityValue: 0,
			RemoteChains:     remoteChainConfigs,
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

func toBytes32LeftPad(b []byte) ([32]byte, error) {
	if len(b) > 32 {
		return [32]byte{}, errors.New("byte slice is too long")
	}
	var result [32]byte
	copy(result[32-len(b):], b)
	return result, nil
}

func paddedLombardChainID(lombardChainID uint32) ([32]byte, error) {
	lchainIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lchainIDBytes, lombardChainID)
	return toBytes32LeftPad(lchainIDBytes)
}
