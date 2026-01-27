package lombard

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/type_and_version"
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

var indexAddressesByTypeAndVersion = type_and_version.IndexAddressesByTypeAndVersion

var (
	lombardQualifier = "Lombard"
)

var DeployLombardChain = cldf_ops.NewSequence(
	"deploy-lombard-chain",
	semver.MustParse("1.0.0"),
	"Deploys the Lombard chain with all required contracts and configurations",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input adapters.DeployLombardInput[string, []byte]) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		verifierTypeAndVersionToAddr, err := indexAddressesByTypeAndVersion(b, chain, input.LombardVerifier)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to index addresses by type and version: %w", err)
		}

		lombardVerifierAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(lombard_verifier.ContractType, *cctp_verifier.Version).String()]
		lombardVerifierResolverAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(versioned_verifier_resolver.ContractType, *versioned_verifier_resolver.Version).String()]

		if lombardVerifierAddress == "" {
			lombardVerifierReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.Deploy, chain, contract_utils.DeployInput[lombard_verifier.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(lombard_verifier.ContractType, *lombard_verifier.Version),
				ChainSelector:  input.ChainSelector,
				Qualifier:      &lombardQualifier,
				Args: lombard_verifier.ConstructorArgs{
					Bridge:           common.HexToAddress(input.Bridge),
					StorageLocations: input.StorageLocations,
					DynamicConfig: lombard_verifier.DynamicConfig{
						FeeAggregator: common.HexToAddress(input.FeeAggregator),
					},
					RMN: common.HexToAddress(input.RMN),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifier: %w", err)
			}
			addresses = append(addresses, lombardVerifierReport.Output)
			lombardVerifierAddress = lombardVerifierReport.Output.Address
		}

		if lombardVerifierResolverAddress == "" {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.ChainSelector,
				Qualifier:      lombardQualifier,
				Type:           datastore.ContractType(lombard_verifier.ResolverType),
				Version:        lombard_verifier.Version,
				CREATE2Factory: common.HexToAddress(input.DeployerContract),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifierResolver: %w", err)
			}
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
			if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected 1 LombardVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
			}
			lombardVerifierResolverAddress = deployVerifierResolverViaCREATE2Report.Output.Addresses[0].Address
		}

		lombardTokenPool := input.TokenPool
		if lombardTokenPool == "" {
			lombardTokenPoolReport, err := cldf_ops.ExecuteOperation(b, lombard_token_pool.Deploy, chain, contract_utils.DeployInput[lombard_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(lombard_token_pool.ContractType, *lombard_token_pool.Version),
				ChainSelector:  input.ChainSelector,
				Qualifier:      &lombardQualifier,
				Args: lombard_token_pool.ConstructorArgs{
					Token:             common.HexToAddress(input.Token),
					LombardVerifier:   common.HexToAddress(lombardVerifierAddress),
					BridgeV2:          common.HexToAddress(input.Bridge),
					AdvancedPoolHooks: common.Address{},
					RMNProxy:          common.HexToAddress(input.RMN),
					Router:            common.HexToAddress(input.Router),
					FallbackDecimals:  18,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardTokenPool: %w", err)
			}
			addresses = append(addresses, lombardTokenPoolReport.Output)
			lombardTokenPool = lombardTokenPoolReport.Output.Address
		}

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: common.HexToAddress(lombardTokenPool),
			RouterAddress:    common.HexToAddress(input.Router),
			RateLimitAdmin:   common.HexToAddress(input.RateLimitAdmin),
			FeeAggregator:    common.HexToAddress(input.FeeAggregator),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainConfigArgs := make([]lombard_verifier.RemoteChainConfigArgs, 0)
		remoteChainSelectors := make([]uint64, 0)

		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remoteChainConfigs[remoteChainSelector] = remoteChain.TokenPoolConfig
			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          common.HexToAddress(lombardVerifierAddress),
			})
			remoteChainSelectors = append(remoteChainSelectors, remoteChainSelector)
			allowedCaller, err := toBytes32LeftPad(remoteChain.RemoteDomain.AllowedCaller)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert allowed caller to bytes32: %w", err)
			}

			lchainIDBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(lchainIDBytes, remoteChain.RemoteDomain.LChainId)
			lchainID, err := toBytes32LeftPad(lchainIDBytes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to convert lombardChainID to bytes32: %w", err)
			}

			setRemotePathReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.SetRemotePath, chain, contract_utils.FunctionInput[lombard_verifier.RemotePathArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(lombardVerifierAddress),
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
			Address:       common.HexToAddress(lombardVerifierResolverAddress),
			Args:          outboundImplementations,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set outbound implementation on LombardVerifierResolver: %w", err)
		}
		writes = append(writes, setOutboundImplementationReport.Output)

		// Apply remote chain config updates on the LombardVerifier
		applyRemoteChainConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.ApplyRemoteChainConfigUpdates, chain, contract_utils.FunctionInput[[]lombard_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(lombardVerifierAddress),
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates on LombardVerifier: %w", err)
		}
		writes = append(writes, applyRemoteChainConfigUpdatesReport.Output)

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	})

func toBytes32LeftPad(b []byte) ([32]byte, error) {
	if len(b) > 32 {
		return [32]byte{}, errors.New("byte slice is too long")
	}
	var result [32]byte
	copy(result[32-len(b):], b)
	return result, nil
}
