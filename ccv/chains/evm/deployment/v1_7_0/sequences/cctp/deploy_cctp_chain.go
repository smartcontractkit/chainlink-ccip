package cctp

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

const (
	localTokenDecimals = 6
)

var cctpQualifier = "CCTP"

var DeployCCTPChain = cldf_ops.NewSequence(
	"deploy-cctp-chain",
	semver.MustParse("1.7.0"),
	"Deploys & configures the CCTP contracts on a chain",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input adapters.DeployCCTPInput[string, []byte]) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Deploy CCTPTokenPool & advanced pool hooks if needed
		var advancedPoolHooksAddress common.Address
		if input.TokenPools.CCTPV2PoolWithCCV == "" {
			advancedPoolHooksReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.Deploy, chain, contract_utils.DeployInput[advanced_pool_hooks.ConstructorArgs]{
				ChainSelector:  chain.Selector,
				TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *semver.MustParse("1.7.0")),
				Qualifier:      &cctpQualifier,
				Args: advanced_pool_hooks.ConstructorArgs{
					Allowlist:                        convertStringsToAddresses(input.Allowlist),
					ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy AdvancedPoolHooks: %w", err)
			}
			addresses = append(addresses, advancedPoolHooksReport.Output)
			advancedPoolHooksAddress = common.HexToAddress(advancedPoolHooksReport.Output.Address)

			cctpTokenPoolReport, err := cldf_ops.ExecuteOperation(b, cctp_token_pool.Deploy, chain, contract_utils.DeployInput[cctp_token_pool.ConstructorArgs]{
				ChainSelector:  chain.Selector,
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_token_pool.ContractType, *semver.MustParse("1.7.0")),
				Qualifier:      &cctpQualifier,
				Args: cctp_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: localTokenDecimals,
					AdvancedPoolHooks:  advancedPoolHooksAddress,
					RMNProxy:           common.HexToAddress(input.RMN),
					Router:             common.HexToAddress(input.Router),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPTokenPool: %w", err)
			}
			addresses = append(addresses, cctpTokenPoolReport.Output)
			input.TokenPools.CCTPV2PoolWithCCV = cctpTokenPoolReport.Output.Address
		} else {
			// Fetch the advanced pool hooks address from the token pool
			advancedPoolHooksAddressReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks address from token pool with address %s on %s: %w", input.TokenPools.CCTPV2PoolWithCCV, chain, err)
			}
			advancedPoolHooksAddress = advancedPoolHooksAddressReport.Output
		}

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSelector,
			TokenPoolAddress:                 common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
			AdvancedPoolHooks:                advancedPoolHooksAddress,
			AllowList:                        convertStringsToAddresses(input.Allowlist),
			RouterAddress:                    common.HexToAddress(input.Router),
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
			RateLimitAdmin:                   common.HexToAddress(input.RateLimitAdmin),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		// Deploy CCTPMessageTransmitterProxy if needed
		if input.MessageTransmitterProxy == "" {
			cctpMessageTransmitterProxyReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.Deploy, chain, contract_utils.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_message_transmitter_proxy.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Qualifier:      &cctpQualifier,
				Args: cctp_message_transmitter_proxy.ConstructorArgs{
					TokenMessenger: common.HexToAddress(input.TokenMessenger),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPMessageTransmitterProxy: %w", err)
			}
			addresses = append(addresses, cctpMessageTransmitterProxyReport.Output)
			input.MessageTransmitterProxy = cctpMessageTransmitterProxyReport.Output.Address
		}

		var cctpVerifierAddress common.Address
		var cctpVerifierResolverAddress common.Address
		for _, cctpVerifier := range input.CCTPVerifier {
			if cctpVerifier.Type == datastore.ContractType(cctp_verifier.ContractType) {
				cctpVerifierAddress = common.HexToAddress(cctpVerifier.Address)
			} else if cctpVerifier.Type == datastore.ContractType(cctp_verifier.ResolverType) {
				cctpVerifierResolverAddress = common.HexToAddress(cctpVerifier.Address)
			}
		}

		// Deploy CCTPVerifier if needed
		if cctpVerifierAddress == (common.Address{}) {
			cctpVerifierReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.Deploy, chain, contract_utils.DeployInput[cctp_verifier.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Qualifier:      &cctpQualifier,
				Args: cctp_verifier.ConstructorArgs{
					TokenMessenger:          common.HexToAddress(input.TokenMessenger),
					MessageTransmitterProxy: common.HexToAddress(input.MessageTransmitterProxy),
					USDCToken:               common.HexToAddress(input.USDCToken),
					StorageLocations:        input.StorageLocations,
					DynamicConfig: cctp_verifier.DynamicConfig{
						FeeAggregator:   common.HexToAddress(input.FeeAggregator),
						AllowlistAdmin:  common.HexToAddress(input.AllowlistAdmin),
						FastFinalityBps: input.FastFinalityBps,
					},
					RMN: common.HexToAddress(input.RMN),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPVerifier: %w", err)
			}
			addresses = append(addresses, cctpVerifierReport.Output)
			cctpVerifierAddress = common.HexToAddress(cctpVerifierReport.Output.Address)
		}

		// Deploy CCTPVerifierResolver if needed
		if cctpVerifierResolverAddress == (common.Address{}) {
			if input.DeployerContract != "" {
				deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
					ChainSelector:  input.ChainSelector,
					Qualifier:      cctpQualifier,
					Type:           datastore.ContractType(cctp_verifier.ResolverType),
					Version:        cctp_verifier.Version,
					CREATE2Factory: common.HexToAddress(input.DeployerContract),
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
				}
				addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
				writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
				if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
					return sequences.OnChainOutput{}, fmt.Errorf("expected 1 CCTPVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
				}
				cctpVerifierResolverAddress = common.HexToAddress(deployVerifierResolverViaCREATE2Report.Output.Addresses[0].Address)
			} else {
				cctpVerifierResolverReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.Deploy, chain, contract_utils.DeployInput[versioned_verifier_resolver.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ResolverType, *semver.MustParse("1.7.0")),
					Qualifier:      &cctpQualifier,
					ChainSelector:  chain.Selector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPVerifierResolver: %w", err)
				}
				addresses = append(addresses, cctpVerifierResolverReport.Output)
				cctpVerifierResolverAddress = common.HexToAddress(cctpVerifierResolverReport.Output.Address)
			}
		}

		// Deploy USDCTokenPoolProxy if needed
		if input.USDCTokenPoolProxy == "" {
			usdcTokenPoolProxyReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Qualifier:      &cctpQualifier,
				Args: usdc_token_pool_proxy.ConstructorArgs{
					Token: common.HexToAddress(input.USDCToken),
					Pools: usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{
						LegacyCctpV1Pool:  common.HexToAddress(input.TokenPools.LegacyCCTPV1Pool),
						CctpV1Pool:        common.HexToAddress(input.TokenPools.CCTPV1Pool),
						CctpV2Pool:        common.HexToAddress(input.TokenPools.CCTPV2Pool),
						CctpV2PoolWithCCV: common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
					},
					Router:       common.HexToAddress(input.Router),
					CCTPVerifier: cctpVerifierResolverAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy: %w", err)
			}
			addresses = append(addresses, usdcTokenPoolProxyReport.Output)
			input.USDCTokenPoolProxy = usdcTokenPoolProxyReport.Output.Address
		}

		// Add CCTPVerifier as an authorized caller on the CCTPMessageTransmitterProxy
		configureAllowedCallersReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.ConfigureAllowedCallers, chain, contract_utils.FunctionInput[[]cctp_message_transmitter_proxy.AllowedCallerConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.MessageTransmitterProxy),
			Args: []cctp_message_transmitter_proxy.AllowedCallerConfigArgs{
				{
					Caller:  cctpVerifierAddress,
					Allowed: true,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to message transmitter proxy: %w", err)
		}
		writes = append(writes, configureAllowedCallersReport.Output)

		// Add USDCTokenPoolProxy as an authorized caller on the CCTPTokenPool
		applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_token_pool.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
			Args: cctp_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(input.USDCTokenPoolProxy),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to CCTPTokenPool: %w", err)
		}
		writes = append(writes, applyAuthorizedCallerUpdatesReport.Output)

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

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainSelectors := make([]uint64, 0)
		mechanisms := make([]uint8, 0)
		lockReleasePools := make([]common.Address, 0)
		setDomainArgs := make([]cctp_verifier.SetDomainArgs, 0)
		remoteChainConfigArgs := make([]cctp_verifier.RemoteChainConfigArgs, 0)
		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remoteChainConfigs[remoteChainSelector] = remoteChain.TokenPoolConfig
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
			lockReleasePools = append(lockReleasePools, common.HexToAddress(remoteChain.LockReleasePool))
			allowedCallerOnDest, err := toBytes32(remoteChain.RemoteDomain.AllowedCallerOnDest)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller on dest to bytes32: %w", err)
			}
			allowedCallerOnSource, err := toBytes32(remoteChain.RemoteDomain.AllowedCallerOnSource)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller on source to bytes32: %w", err)
			}
			mintRecipientOnDest, err := toBytes32(remoteChain.RemoteDomain.MintRecipientOnDest)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert mint recipient on dest to bytes32: %w", err)
			}
			setDomainArgs = append(setDomainArgs, cctp_verifier.SetDomainArgs{
				AllowedCallerOnDest:   allowedCallerOnDest,
				AllowedCallerOnSource: allowedCallerOnSource,
				MintRecipientOnDest:   mintRecipientOnDest,
				DomainIdentifier:      remoteChain.RemoteDomain.DomainIdentifier,
				Enabled:               true,
				ChainSelector:         remoteChainSelector,
			})
			remoteChainConfigArgs = append(remoteChainConfigArgs, cctp_verifier.RemoteChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
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

		// Set lock or burn mechanism for each remote chain
		updateLockOrBurnMechanismsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.USDCTokenPoolProxy),
			Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
				RemoteChainSelectors: remoteChainSelectors,
				Mechanisms:           mechanisms,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update lock or burn mechanisms on USDCTokenPoolProxy: %w", err)
		}
		writes = append(writes, updateLockOrBurnMechanismsReport.Output)

		// Update lock release pools on the USDCTokenPoolProxy
		updateLockReleasePoolAddressesReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockReleasePoolAddresses, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockReleasePoolAddressesArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.USDCTokenPoolProxy),
			Args: usdc_token_pool_proxy.UpdateLockReleasePoolAddressesArgs{
				RemoteChainSelectors: remoteChainSelectors,
				LockReleasePools:     lockReleasePools,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update lock release pools on USDCTokenPoolProxy: %w", err)
		}
		writes = append(writes, updateLockReleasePoolAddressesReport.Output)

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

		// Call into configure token for transfers sequence
		remoteChains := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remoteChains[remoteChainSelector] = remoteChain.TokenPoolConfig
		}
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, chains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: input.TokenPools.CCTPV2PoolWithCCV,
			RegistryAddress:  input.TokenAdminRegistry,
			MinFinalityValue: input.MinFinalityValue,
			RemoteChains:     remoteChains,
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

func convertStringsToAddresses(strs []string) []common.Address {
	addrs := make([]common.Address, len(strs))
	for i, str := range strs {
		addrs[i] = common.HexToAddress(str)
	}
	return addrs
}

func toBytes32(b []byte) ([32]byte, error) {
	if len(b) > 32 {
		return [32]byte{}, errors.New("byte slice is too long")
	}
	var result [32]byte
	copy(result[:], b)
	return result, nil
}

func convertMechanismToUint8(mechanism adapters.Mechanism) (uint8, error) {
	switch mechanism {
	case adapters.CCTPV1Mechanism:
		return 1, nil
	case adapters.CCTPV2Mechanism:
		return 2, nil
	case adapters.LockReleaseMechanism:
		return 3, nil
	case adapters.CCTPV2WithCCVMechanism:
		return 4, nil
	default:
		return 0, fmt.Errorf("invalid mechanism: %s", mechanism)
	}
}
