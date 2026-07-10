package lombard

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	ltp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lombard_token_pool"
	lombard_verifier_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lombard_verifier"
	aph_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	vvr_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_verifier"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

var ConfigureLombardChainForLanes = cldf_ops.NewSequence(
	"configure-lombard-chain-for-lanes",
	semver.MustParse("2.0.0"),
	"Configures the Lombard chain to support CCIP lanes",
	func(b cldf_ops.Bundle, dep adapters.ConfigureLombardChainForLanesDeps, input adapters.ConfigureLombardChainForLanesInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
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
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find LombardVerifier ref on chain %d: %w", chain.Selector, err)
		}

		lombardVerifierResolverAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType),
			Version: lombard_verifier.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find LombardVerifierResolver ref on chain %d: %w", chain.Selector, err)
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
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find LombardTokenPool ref on chain %d: %w", chain.Selector, err)
		}

		tokenAdminRegistryAddressRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find TokenAdminRegistry ref on chain %d: %w", chain.Selector, err)
		}

		advancedPoolHooksAddress := common.HexToAddress(advancedPoolHooksRef.Address)
		tokenPoolAddress := common.HexToAddress(tokenPoolRef.Address)
		routerAddress := common.HexToAddress(routerRef.Address)
		lombardVerifierAddress := common.HexToAddress(lombardVerifierAddressRef.Address)
		lombardVerifierResolverAddress := common.HexToAddress(lombardVerifierResolverAddressRef.Address)
		sourceTokenOrAdapterForRemoteAdapterLookup := common.HexToAddress(input.Token)
		if input.LocalAdapter != "" {
			sourceTokenOrAdapterForRemoteAdapterLookup = common.HexToAddress(input.LocalAdapter)
		}

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainConfigArgs := make([]lombard_verifier_bindings.BaseVerifierRemoteChainConfigArgs, 0)
		remoteAdapterArgs := make([]lombard_verifier_bindings.LombardVerifierRemoteAdapterArgs, 0)
		tokenPoolPathArgs := make([]lombard_token_pool.SetPathArgs, 0)
		advancedPoolHooks := make([]aph_bindings.AdvancedPoolHooksCCVConfigArg, 0)

		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remotePoolAddress, err := dep.RemoteChains[remoteChainSelector].RemoteTokenPoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector, *tokenPoolQualifier(input.TokenQualifier))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pool address: %w", err)
			}
			remoteTokenAddress, err := dep.RemoteChains[remoteChainSelector].RemoteTokenAddress(b, dep.DataStore, dep.BlockChains, remoteChainSelector, *tokenPoolQualifier(input.TokenQualifier))
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token address: %w", err)
			}

			feeCfg := (tokens_core.PartialTokenTransferFeeConfig{}).Populate(remoteChain.TokenTransferFeeConfig)
			remoteChainConfigs[remoteChainSelector] = tokens_core.RemoteChainConfig[[]byte, string]{
				RemotePool:             common.LeftPadBytes(remotePoolAddress, 32),
				RemoteToken:            common.LeftPadBytes(remoteTokenAddress, 32),
				TokenTransferFeeConfig: &feeCfg,
				// Lombard does not use rate limiters
				InboundRateLimiterConfig: &tokens_core.RateLimiterConfigFloatInput{
					Capacity: 0,
					Rate:     0,
				},
				OutboundRateLimiterConfig: &tokens_core.RateLimiterConfigFloatInput{
					Capacity: 0,
					Rate:     0,
				},
			}

			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          lombardVerifierAddress,
			})

			remoteChainConfigArgs = append(remoteChainConfigArgs, lombard_verifier_bindings.BaseVerifierRemoteChainConfigArgs{
				Router:              routerAddress,
				RemoteChainSelector: remoteChainSelector,
				FeeUSDCents:         remoteChain.FeeUSDCents,
				GasForVerification:  remoteChain.GasForVerification,
				PayloadSizeBytes:    remoteChain.PayloadSizeBytes,
			})

			advancedPoolHooks = append(advancedPoolHooks, aph_bindings.AdvancedPoolHooksCCVConfigArg{
				RemoteChainSelector: remoteChainSelector,
				OutboundCCVs: []common.Address{
					lombardVerifierResolverAddress,
					{}, // This means "require the default CCV(s) for this lane".
				},
				InboundCCVs: []common.Address{
					lombardVerifierResolverAddress,
					{}, // This means "require the default CCV(s) for this lane".
				},
			})

			remoteAdapter := [32]byte{}
			if remoteChain.RemoteAdapter != "" {
				remoteAdapter, err = parseRemoteAdapter(remoteChain.RemoteAdapter)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to parse remote adapter for remote chain %d: %w", remoteChainSelector, err)
				}
			}

			existingRemoteAdapterReport, err := evmops.ExecuteRead(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewReadGetRemoteAdapter, lombard_verifier.GetRemoteAdapterArgs{
				RemoteChainSelector: remoteChainSelector,
				Token:               sourceTokenOrAdapterForRemoteAdapterLookup,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote adapter on LombardVerifier for remote chain %d: %w", remoteChainSelector, err)
			}
			if existingRemoteAdapterReport.Output != remoteAdapter {
				remoteAdapterArgs = append(remoteAdapterArgs, lombard_verifier_bindings.LombardVerifierRemoteAdapterArgs{
					RemoteChainSelector: remoteChainSelector,
					Token:               sourceTokenOrAdapterForRemoteAdapterLookup,
					RemoteAdapter:       remoteAdapter,
				})
			}

			remoteCallerOnDest, err := dep.RemoteChains[remoteChainSelector].AllowedCallerOnDest(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed caller on dest for remote chain %d: %w", remoteChainSelector, err)
			}
			verifierAllowedCaller, err := toBytes32LeftPad(remoteCallerOnDest)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller to bytes32 for verifier path on chain %d: %w", remoteChainSelector, err)
			}
			lchainID, err := paddedLombardChainID(remoteChain.LombardChainId)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lombardChainID to bytes32: %w", err)
			}

			tokenPoolPathArgs = append(tokenPoolPathArgs, lombard_token_pool.SetPathArgs{
				RemoteChainSelector: remoteChainSelector,
				LChainId:            lchainID,
				// TokenPool remote pools are configured as ABI-style padded bytes for EVM addresses.
				AllowedCaller: common.LeftPadBytes(remotePoolAddress, 32),
				RemoteAdapter: remoteAdapter,
			})

			existingVerifierPathReport, err := evmops.ExecuteRead(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewReadGetPath, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote path on LombardVerifier for remote chain %d: %w", remoteChainSelector, err)
			}

			if existingVerifierPathReport.Output.LChainId != lchainID || existingVerifierPathReport.Output.AllowedCaller != verifierAllowedCaller {
				setRemotePathReport, err := evmops.ExecuteWrite(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewWriteSetPath, lombard_verifier.SetPathArgs{
					RemoteChainSelector: remoteChainSelector,
						LChainId:            lchainID,
						AllowedCaller:       remoteCallerOnDest,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set remote path on LombardVerifier for remote chain %d: %w", remoteChainSelector, err)
				}
				writes = append(writes, setRemotePathReport.Output)
			}
		}

		// Set outbound implementation on the LombardVerifierResolver for each remote chain
		setOutboundImplementationReport, err := evmops.ExecuteWrite(b, chain, lombardVerifierResolverAddress, vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewWriteApplyOutboundImplementationUpdates, outboundImplementations)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set outbound implementation on LombardVerifierResolver: %w", err)
		}
		writes = append(writes, setOutboundImplementationReport.Output)

		applyRemoteChainConfigUpdatesReport, err := evmops.ExecuteWrite(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewWriteApplyRemoteChainConfigUpdates, remoteChainConfigArgs)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on LombardVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		if len(remoteAdapterArgs) > 0 {
			setRemoteAdaptersReport, err := evmops.ExecuteWrite(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewWriteSetRemoteAdapters, remoteAdapterArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set remote adapters on LombardVerifier: %w", err)
			}
			writes = append(writes, setRemoteAdaptersReport.Output)
		}

		advancedPoolHooksApplyReport, err := evmops.ExecuteWrite(b, chain, advancedPoolHooksAddress, evmops.BindAs[aph_bindings.AdvancedPoolHooksInterface](aph_bindings.NewAdvancedPoolHooks), advanced_pool_hooks.NewWriteApplyCCVConfigUpdates, advancedPoolHooks)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply advanced pool hooks CCV config updates: %w", err)
		}
		writes = append(writes, advancedPoolHooksApplyReport.Output)

		// Call into configure token for transfers sequence
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, dep.BlockChains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:         input.ChainSelector,
			TokenAddress:          input.Token,
			TokenPoolAddress:      tokenPoolRef.Address,
			RegistryAddress:       tokenAdminRegistryAddressRef.Address,
			AllowedFinalityConfig: finality.Config{WaitForFinality: true},
			RemoteChains:          remoteChainConfigs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token for transfers: %w", err)
		}
		batchOps = append(batchOps, configureTokenForTransfersReport.Output.BatchOps...)

		for _, pathArgs := range tokenPoolPathArgs {
			existingPathReport, err := evmops.ExecuteRead(b, chain, tokenPoolAddress, evmops.BindAs[ltp_bindings.LombardTokenPoolInterface](ltp_bindings.NewLombardTokenPool), lombard_token_pool.NewReadGetPath, pathArgs.RemoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get LombardTokenPool path for remote chain %d: %w", pathArgs.RemoteChainSelector, err)
			}

			if existingPathReport.Output.LChainId == pathArgs.LChainId &&
				existingPathReport.Output.RemoteAdapter == pathArgs.RemoteAdapter &&
				bytes.Equal(existingPathReport.Output.AllowedCaller[:], pathArgs.AllowedCaller) {
				continue
			}

			setPathReport, err := evmops.ExecuteWrite(b, chain, tokenPoolAddress, evmops.BindAs[ltp_bindings.LombardTokenPoolInterface](ltp_bindings.NewLombardTokenPool), lombard_token_pool.NewWriteSetPath, pathArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set LombardTokenPool path for remote chain %d: %w", pathArgs.RemoteChainSelector, err)
			}
			writes = append(writes, setPathReport.Output)
		}

		if len(writes) > 0 {
			batchOpFromWrites, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOpFromWrites)
		}

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

func parseRemoteAdapter(raw string) ([32]byte, error) {
	if common.IsHexAddress(raw) {
		return toBytes32LeftPad(common.HexToAddress(raw).Bytes())
	}

	decoded, err := hexutil.Decode(raw)
	if err != nil {
		return [32]byte{}, err
	}
	if len(decoded) == 0 {
		return [32]byte{}, errors.New("remote adapter cannot be empty")
	}

	return toBytes32LeftPad(decoded)
}
