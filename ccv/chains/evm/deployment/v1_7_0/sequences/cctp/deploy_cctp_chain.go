package cctp

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
)

var indexAddressesByTypeAndVersion = type_and_version.IndexAddressesByTypeAndVersion

const (
	localTokenDecimals = 6
)

var (
	cctpQualifier = "CCTP"

	// This sequence assumes that CCTP V2 and CCTP V1 are all on pool version 1.6.4
	prevVersion        = semver.MustParse("1.6.4")
	cctpV2ContractType = deployment.ContractType("USDCTokenPoolCCTPV2")
	cctpV1ContractType = deployment.ContractType("USDCTokenPool")
)

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

		poolTypeAndVersionToAddr, err := indexAddressesByTypeAndVersion(b, chain, input.TokenPool)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to index addresses by type and version: %w", err)
		}

		// Deploy CCTPMessageTransmitterProxy if needed
		if input.MessageTransmitterProxy == "" {
			cctpMessageTransmitterProxyReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.Deploy, chain, contract_utils.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_message_transmitter_proxy.ContractType, *cctp_message_transmitter_proxy.Version),
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

		verifierTypeAndVersionToAddr, err := indexAddressesByTypeAndVersion(b, chain, input.CCTPVerifier)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to index addresses by type and version: %w", err)
		}
		cctpVerifierAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version).String()]
		cctpVerifierResolverAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(versioned_verifier_resolver.ContractType, *versioned_verifier_resolver.Version).String()]

		// Deploy CCTPVerifier if needed
		if cctpVerifierAddress == "" {
			cctpVerifierReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.Deploy, chain, contract_utils.DeployInput[cctp_verifier.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_verifier.ContractType, *cctp_verifier.Version),
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
			cctpVerifierAddress = cctpVerifierReport.Output.Address
		}

		// Deploy CCTPVerifierResolver if needed
		if cctpVerifierResolverAddress == "" {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

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
			cctpVerifierResolverAddress = deployVerifierResolverViaCREATE2Report.Output.Addresses[0].Address
		}

		// Deploy CCTPTokenPool if needed
		cctpTokenPoolAddress := poolTypeAndVersionToAddr[deployment.NewTypeAndVersion(cctp_through_ccv_token_pool.ContractType, *cctp_through_ccv_token_pool.Version).String()]
		if cctpTokenPoolAddress == "" {
			cctpTokenPoolReport, err := cldf_ops.ExecuteOperation(b, cctp_through_ccv_token_pool.Deploy, chain, contract_utils.DeployInput[cctp_through_ccv_token_pool.ConstructorArgs]{
				ChainSelector:  chain.Selector,
				TypeAndVersion: deployment.NewTypeAndVersion(cctp_through_ccv_token_pool.ContractType, *cctp_through_ccv_token_pool.Version),
				Qualifier:      &cctpQualifier,
				Args: cctp_through_ccv_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: localTokenDecimals,
					RMNProxy:           common.HexToAddress(input.RMN),
					Router:             common.HexToAddress(input.Router),
					CCTPVerifier:       common.HexToAddress(cctpVerifierAddress),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPTokenPool: %w", err)
			}
			addresses = append(addresses, cctpTokenPoolReport.Output)
			cctpTokenPoolAddress = cctpTokenPoolReport.Output.Address
		}

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: common.HexToAddress(cctpTokenPoolAddress),
			RouterAddress:    common.HexToAddress(input.Router),
			RateLimitAdmin:   common.HexToAddress(input.RateLimitAdmin),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		// Deploy USDCTokenPoolProxy if needed
		usdcTokenPoolProxyAddress := poolTypeAndVersionToAddr[deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version).String()]
		if usdcTokenPoolProxyAddress == "" {
			cctpV1PoolAddress := poolTypeAndVersionToAddr[deployment.NewTypeAndVersion(cctpV1ContractType, *prevVersion).String()]
			if cctpV1PoolAddress == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("cctp v1 pool with type and version %s not found", deployment.NewTypeAndVersion(cctpV1ContractType, *prevVersion).String())
			}
			cctpV2PoolAddress := poolTypeAndVersionToAddr[deployment.NewTypeAndVersion(cctpV2ContractType, *prevVersion).String()]
			if cctpV2PoolAddress == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("cctp v2 pool with type and version %s not found", deployment.NewTypeAndVersion(cctpV2ContractType, *prevVersion).String())
			}
			// Siloed lock release pool is required on Ethereum mainnet and Ethereum testnet Sepolia
			siloedLockReleasePoolAddress := poolTypeAndVersionToAddr[deployment.NewTypeAndVersion(siloed_usdc_token_pool.ContractType, *siloed_usdc_token_pool.Version).String()]
			if siloedLockReleasePoolAddress == "" && (input.ChainSelector == chain_selectors.ETHEREUM_MAINNET.Selector || input.ChainSelector == chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector) {
				return sequences.OnChainOutput{}, fmt.Errorf("siloed lock release pool with type and version %s not found", deployment.NewTypeAndVersion(siloed_usdc_token_pool.ContractType, *siloed_usdc_token_pool.Version).String())
			}

			usdcTokenPoolProxyReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.Deploy, chain, contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
				ChainSelector:  chain.Selector,
				Qualifier:      &cctpQualifier,
				Args: usdc_token_pool_proxy.ConstructorArgs{
					Token: common.HexToAddress(input.USDCToken),
					Pools: usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{
						CctpV1Pool:            common.HexToAddress(cctpV1PoolAddress),
						CctpV2Pool:            common.HexToAddress(cctpV2PoolAddress),
						CctpV2PoolWithCCV:     common.HexToAddress(cctpTokenPoolAddress),
						SiloedLockReleasePool: common.HexToAddress(siloedLockReleasePoolAddress),
					},
					Router:       common.HexToAddress(input.Router),
					CCTPVerifier: common.HexToAddress(cctpVerifierResolverAddress),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy: %w", err)
			}
			addresses = append(addresses, usdcTokenPoolProxyReport.Output)
			usdcTokenPoolProxyAddress = usdcTokenPoolProxyReport.Output.Address
		}

		// Add CCTPVerifier as an authorized caller on the CCTPMessageTransmitterProxy
		verifierApplyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_message_transmitter_proxy.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[cctp_message_transmitter_proxy.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.MessageTransmitterProxy),
			Args: cctp_message_transmitter_proxy.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(cctpVerifierAddress),
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
			Address:       common.HexToAddress(cctpTokenPoolAddress),
			Args: cctp_through_ccv_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(usdcTokenPoolProxyAddress),
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
			Address:       common.HexToAddress(cctpVerifierAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CCTPVerifier: %w", err)
		}
		setInboundImplementationReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cctpVerifierResolverAddress),
			Args: []versioned_verifier_resolver.InboundImplementationArgs{
				{
					Version:  committeeVerifierVersionTagReport.Output,
					Verifier: common.HexToAddress(cctpVerifierAddress),
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
		setDomainArgs := make([]cctp_verifier.SetDomainArgs, 0)
		remoteChainConfigArgs := make([]cctp_verifier.RemoteChainConfigArgs, 0)
		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remoteChainConfigs[remoteChainSelector] = remoteChain.TokenPoolConfig
			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          common.HexToAddress(cctpVerifierAddress),
			})
			remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)
			mechanism, err := convertMechanismToUint8(remoteChain.LockOrBurnMechanism)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lock or burn mechanism to uint8: %w", err)
			}
			mechanisms = append(mechanisms, mechanism)
			allowedCallerOnDest, err := toBytes32LeftPad(remoteChain.RemoteDomain.AllowedCallerOnDest)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller on dest to bytes32: %w", err)
			}
			allowedCallerOnSource, err := toBytes32LeftPad(remoteChain.RemoteDomain.AllowedCallerOnSource)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller on source to bytes32: %w", err)
			}
			mintRecipientOnDest, err := toBytes32LeftPad(remoteChain.RemoteDomain.MintRecipientOnDest)
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
			Address:       common.HexToAddress(cctpVerifierResolverAddress),
			Args:          outboundImplementations,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set outbound implementation on CCTPVerifierResolver: %w", err)
		}
		writes = append(writes, setOutboundImplementationReport.Output)

		// Set lock or burn mechanism for each remote chain
		updateLockOrBurnMechanismsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(usdcTokenPoolProxyAddress),
			Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
				RemoteChainSelectors: remoteChainSelectors,
				Mechanisms:           mechanisms,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update lock or burn mechanisms on USDCTokenPoolProxy: %w", err)
		}
		writes = append(writes, updateLockOrBurnMechanismsReport.Output)

		// Apply remote chain config updates on the CCTPVerifier
		applyRemoteChainConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]cctp_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cctpVerifierAddress),
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on CCTPVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		// Set each remote domain on the CCTPVerifier
		setDomainsReport, err := cldf_ops.ExecuteOperation(b, cctp_verifier.SetDomains, chain, contract_utils.FunctionInput[[]cctp_verifier.SetDomainArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cctpVerifierAddress),
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
			TokenPoolAddress: cctpTokenPoolAddress,
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

func toBytes32LeftPad(b []byte) ([32]byte, error) {
	if len(b) > 32 {
		return [32]byte{}, errors.New("byte slice is too long")
	}
	var result [32]byte
	copy(result[32-len(b):], b)
	return result, nil
}

func convertMechanismToUint8(mechanism string) (uint8, error) {
	switch mechanism {
	case "CCTP_V1":
		return 1, nil
	case "CCTP_V2":
		return 2, nil
	case "LOCK_RELEASE":
		return 3, nil
	case "CCTP_V2_WITH_CCV":
		return 4, nil
	default:
		return 0, fmt.Errorf("invalid mechanism, must be CCTP_V1, CCTP_V2, LOCK_RELEASE, or CCTP_V2_WITH_CCV: %s", mechanism)
	}
}
