package lombard

import (
	"encoding/binary"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
)

var ConfigureLombardChainForLanes = cldf_ops.NewSequence(
	"configure-lombard-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures the Lombard chain to support CCIP lanes",
	func(b cldf_ops.Bundle, dep adapters.DeployLombardChainDeps, input adapters.ConfigureLombardChainForLanesInput) (output sequences.OnChainOutput, err error) {
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

		lombardVerifierAddress := common.HexToAddress(lombardVerifierAddressRef.Address)
		lombardVerifierResolverAddress := common.HexToAddress(lombardVerifierResolverAddressRef.Address)

		remoteChainConfigs := make(map[uint64]tokens_core.RemoteChainConfig[[]byte, string])
		outboundImplementations := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0)
		remoteChainSelectors := make([]uint64, 0)
		remoteChainConfigArgs := make([]cctp_verifier.RemoteChainConfigArgs, 0)

		for remoteChainSelector, remoteChain := range input.RemoteChains {
			remoteChainConfigs[remoteChainSelector] = remoteChain.TokenPoolConfig
			outboundImplementations = append(outboundImplementations, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteChainSelector,
				Verifier:          lombardVerifierAddress,
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

		return 1, nil
	},
)
