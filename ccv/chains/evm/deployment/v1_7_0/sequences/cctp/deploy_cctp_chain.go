package cctp

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type Mechanism string

const (
	CCTPV1Mechanism        Mechanism = "CCTP_V1"
	CCTPV2Mechanism        Mechanism = "CCTP_V2"
	LockReleaseMechanism   Mechanism = "LOCK_RELEASE"
	CCTPV2WithCCVMechanism Mechanism = "CCTP_V2_WITH_CCV"
)

func (m Mechanism) ToUint8() (uint8, error) {
	switch m {
	case CCTPV1Mechanism:
		return 1, nil
	case CCTPV2Mechanism:
		return 2, nil
	case LockReleaseMechanism:
		return 3, nil
	case CCTPV2WithCCVMechanism:
		return 4, nil
	default:
		return 0, fmt.Errorf("invalid mechanism: %s", m)
	}
}

const (
	localTokenDecimals = 6
)

var cctpQualifier = "CCTP"

type DynamicConfig struct {
	FeeAggregator   string
	AllowlistAdmin  string
	FastFinalityBps uint16
}

type RemoteDomain struct {
	AllowedCallerOnDest   string
	AllowedCallerOnSource string
	MintRecipientOnDest   string
	DomainIdentifier      uint32
	Enabled               bool
}

type RemoteChainConfig struct {
	tokens_core.RemoteChainConfig[[]byte, string]
	AllowlistEnabled    bool
	FeeUSDCents         uint16
	GasForVerification  uint32
	PayloadSizeBytes    uint32
	LockOrBurnMechanism Mechanism
	LockReleasePool     string
	RemoteDomain        RemoteDomain
}

type TokenPools struct {
	LegacyCCTPV1Pool  string
	CCTPV1Pool        string
	CCTPV2Pool        string
	CCTPV2PoolWithCCV string
}

type DeployCCTPInput struct {
	ChainSelector                    uint64
	TokenPools                       TokenPools
	USDCTokenPoolProxy               string
	CCTPVerifier                     string
	MessageTransmitterProxy          string
	CCTPVerifierResolver             string
	TokenAdminRegistry               string
	AdvancedPoolHooks                string
	TokenMessenger                   string
	USDCToken                        string
	MinFinalityValue                 uint16
	StorageLocations                 []string
	DynamicConfig                    DynamicConfig
	RMN                              string
	Router                           string
	CREATE2Factory                   string
	Allowlist                        []string
	ThresholdAmountForAdditionalCCVs *big.Int
	RateLimitAdmin                   string
	RemoteChains                     map[uint64]RemoteChainConfig
}

var DeployCCTPChain = cldf_ops.NewSequence(
	"deploy-cctp-chain",
	semver.MustParse("1.7.0"),
	"Deploys & configures the CCTP contracts on a chain",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input DeployCCTPInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Deploy AdvancedPoolHooks if needed (for CCTPTokenPool)
		if input.AdvancedPoolHooks == "" {
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
			input.AdvancedPoolHooks = advancedPoolHooksReport.Output.Address
		}

		// Deploy CCTPTokenPool if needed
		if input.TokenPools.CCTPV2PoolWithCCV == "" {
			cctpTokenPoolReport, err := cldf_ops.ExecuteOperation(b, cctp_token_pool.Deploy, chain, contract_utils.DeployInput[cctp_token_pool.ConstructorArgs]{
				ChainSelector: chain.Selector,
				Qualifier:     &cctpQualifier,
				Args: cctp_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: localTokenDecimals,
					AdvancedPoolHooks:  common.HexToAddress(input.AdvancedPoolHooks),
					RMNProxy:           common.HexToAddress(input.RMN),
					Router:             common.HexToAddress(input.Router),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPTokenPool: %w", err)
			}
			addresses = append(addresses, cctpTokenPoolReport.Output)
			input.TokenPools.CCTPV2PoolWithCCV = cctpTokenPoolReport.Output.Address
		}

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSelector,
			TokenPoolAddress:                 common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
			AdvancedPoolHooks:                common.HexToAddress(input.AdvancedPoolHooks),
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

		// Deploy CCTPVerifier if needed
		if input.CCTPVerifier == "" {
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
						FeeAggregator:   common.HexToAddress(input.DynamicConfig.FeeAggregator),
						AllowlistAdmin:  common.HexToAddress(input.DynamicConfig.AllowlistAdmin),
						FastFinalityBps: input.DynamicConfig.FastFinalityBps,
					},
					RMN: common.HexToAddress(input.RMN),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPVerifier: %w", err)
			}
			addresses = append(addresses, cctpVerifierReport.Output)
			input.CCTPVerifier = cctpVerifierReport.Output.Address
		}

		// Deploy USDCTokenPoolProxy if needed
		if input.USDCTokenPoolProxy == "" {
			usdcTokenPoolProxyReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Args: usdc_token_pool_proxy.ConstructorArgs{
					Token: common.HexToAddress(input.USDCToken),
					Pools: usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{
						LegacyCctpV1Pool:  common.HexToAddress(input.TokenPools.LegacyCCTPV1Pool),
						CctpV1Pool:        common.HexToAddress(input.TokenPools.CCTPV1Pool),
						CctpV2Pool:        common.HexToAddress(input.TokenPools.CCTPV2Pool),
						CctpV2PoolWithCCV: common.HexToAddress(input.TokenPools.CCTPV2PoolWithCCV),
					},
					Router:       common.HexToAddress(input.Router),
					CCTPVerifier: common.HexToAddress(input.CCTPVerifier),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy: %w", err)
			}
			addresses = append(addresses, usdcTokenPoolProxyReport.Output)
			input.USDCTokenPoolProxy = usdcTokenPoolProxyReport.Output.Address
		}

		// Deploy CCTPVerifierResolver if needed
		if input.CCTPVerifierResolver == "" {
			if input.CREATE2Factory != "" {
				deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
					ChainSelector:  input.ChainSelector,
					Qualifier:      cctpQualifier,
					Type:           datastore.ContractType(cctp_verifier.ResolverType),
					Version:        cctp_verifier.Version,
					CREATE2Factory: common.HexToAddress(input.CREATE2Factory),
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
				}
				addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
				writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
				if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
					return sequences.OnChainOutput{}, fmt.Errorf("expected 1 CCTPVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
				}
				input.CCTPVerifierResolver = deployVerifierResolverViaCREATE2Report.Output.Addresses[0].Address
			} else {
				cctpVerifierResolverReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.Deploy, chain, contract_utils.DeployInput[versioned_verifier_resolver.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ResolverType, *semver.MustParse("1.7.0")),
					ChainSelector:  chain.Selector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPVerifierResolver: %w", err)
				}
				addresses = append(addresses, cctpVerifierResolverReport.Output)
				input.CCTPVerifierResolver = cctpVerifierResolverReport.Output.Address
			}
		}

		// Add CCTPVerifier as an authorized caller on the CCTPMessageTransmitterProxy
		configureAllowedCallersReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.ConfigureAllowedCallers, chain, contract_utils.FunctionInput[[]cctp_message_transmitter_proxy.AllowedCallerConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.MessageTransmitterProxy),
			Args: []cctp_message_transmitter_proxy.AllowedCallerConfigArgs{
				{
					Caller:  common.HexToAddress(input.CCTPVerifier),
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
			Address:       common.HexToAddress(input.CCTPVerifier),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CCTPVerifier: %w", err)
		}
		setInboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.CCTPVerifierResolver),
			Args: []versioned_verifier_resolver.InboundImplementationArgs{
				{
					Version:  committeeVerifierVersionTagReport.Output,
					Verifier: common.HexToAddress(input.CCTPVerifier),
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
			remoteChainConfigs[remoteChainSelector] = remoteChain.RemoteChainConfig
			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          common.HexToAddress(input.CCTPVerifier),
			})
			remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)
			mechanism, err := remoteChain.LockOrBurnMechanism.ToUint8()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lock or burn mechanism to uint8: %w", err)
			}
			mechanisms = append(mechanisms, mechanism)
			lockReleasePools = append(lockReleasePools, common.HexToAddress(remoteChain.LockReleasePool))
			setDomainArgs = append(setDomainArgs, cctp_verifier.SetDomainArgs{
				AllowedCallerOnDest:   convertAddressToBytes32(common.HexToAddress(remoteChain.RemoteDomain.AllowedCallerOnDest)),
				AllowedCallerOnSource: convertAddressToBytes32(common.HexToAddress(remoteChain.RemoteDomain.AllowedCallerOnSource)),
				MintRecipientOnDest:   convertAddressToBytes32(common.HexToAddress(remoteChain.RemoteDomain.MintRecipientOnDest)),
				DomainIdentifier:      remoteChain.RemoteDomain.DomainIdentifier,
				Enabled:               remoteChain.RemoteDomain.Enabled,
			})
			remoteChainConfigArgs = append(remoteChainConfigArgs, cctp_verifier.RemoteChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				RemoteChainSelector: remoteChainSelector,
				AllowlistEnabled:    remoteChain.AllowlistEnabled,
				FeeUSDCents:         remoteChain.FeeUSDCents,
				GasForVerification:  remoteChain.GasForVerification,
				PayloadSizeBytes:    remoteChain.PayloadSizeBytes,
			})
		}

		// Set outbound implementation on the CCTPVerifierResolver for each remote chain
		setOutboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.CCTPVerifierResolver),
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
			Address:       common.HexToAddress(input.CCTPVerifier),
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on CCTPVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		// Set each remote domain on the CCTPVerifier
		setDomainsReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.SetDomains, chain, contract_utils.FunctionInput[[]cctp_verifier.SetDomainArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.CCTPVerifier),
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
			remoteChains[remoteChainSelector] = remoteChain.RemoteChainConfig
		}
		configureTokenForTransfersReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenForTransfers, chains, tokens_core.ConfigureTokenForTransfersInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: input.TokenPools.CCTPV2PoolWithCCV,
			RegistryAddress:  input.TokenAdminRegistry,
			FinalityValue:    input.MinFinalityValue,
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

func convertAddressToBytes32(addr common.Address) [32]byte {
	leftPaddedBytes := common.LeftPadBytes(addr.Bytes(), 32)
	var result [32]byte
	copy(result[:], leftPaddedBytes)
	return result
}
