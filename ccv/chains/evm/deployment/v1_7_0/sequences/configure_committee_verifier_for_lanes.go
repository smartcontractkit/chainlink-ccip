package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureCommitteeVerifierForLanesInput struct {
	ChainSelector            uint64
	RemoteChainsToDisconnect []uint64
	Router                   string
	adapters.CommitteeVerifierConfig[string]
}

var ConfigureCommitteeVerifierForLanes = cldf_ops.NewSequence(
	"configure-committee-verifier-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures a CommitteeVerifier contract for multiple remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input ConfigureCommitteeVerifierForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		var committeeVerifier string
		var committeeVerifierResolver string
		typeAndVersionToAddress, err := type_and_version.IndexAddressesByTypeAndVersion(b, chain, input.CommitteeVerifier)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to index addresses by type and version for committee verifier: %w", err)
		}
		committeeVerifier = typeAndVersionToAddress[deployment.NewTypeAndVersion(committee_verifier.ContractType, *committee_verifier.Version).String()]
		committeeVerifierResolver = typeAndVersionToAddress[deployment.NewTypeAndVersion(versioned_verifier_resolver.ContractType, *versioned_verifier_resolver.Version).String()]

		if committeeVerifier == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("committee verifier contract not found on chain %d", input.ChainSelector)
		}
		if committeeVerifierResolver == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("committee verifier resolver contract not found on chain %d", input.ChainSelector)
		}

		remoteChainConfigArgs := make([]committee_verifier.RemoteChainConfigArgs, 0, len(input.RemoteChains))
		signatureConfigs := make([]committee_verifier.SignatureConfig, 0, len(input.RemoteChains))
		outboundImplementationArgs := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0, len(input.RemoteChains))

		for remoteSelector, remoteConfig := range input.RemoteChains {
			remoteChainConfigArgs = append(remoteChainConfigArgs, committee_verifier.RemoteChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				RemoteChainSelector: remoteSelector,
				AllowlistEnabled:    remoteConfig.AllowlistEnabled,
				FeeUSDCents:         remoteConfig.FeeUSDCents,
				GasForVerification:  remoteConfig.GasForVerification,
				PayloadSizeBytes:    remoteConfig.PayloadSizeBytes,
			})
			addedAllowlistedSenders := make([]common.Address, 0, len(remoteConfig.AddedAllowlistedSenders))
			for _, sender := range remoteConfig.AddedAllowlistedSenders {
				addedAllowlistedSenders = append(addedAllowlistedSenders, common.HexToAddress(sender))
			}
			removedAllowlistedSenders := make([]common.Address, 0, len(remoteConfig.RemovedAllowlistedSenders))
			for _, sender := range remoteConfig.RemovedAllowlistedSenders {
				removedAllowlistedSenders = append(removedAllowlistedSenders, common.HexToAddress(sender))
			}
			signers := make([]common.Address, 0, len(remoteConfig.SignatureConfig.Signers))
			for _, signer := range remoteConfig.SignatureConfig.Signers {
				signers = append(signers, common.HexToAddress(signer))
			}
			signatureConfigs = append(signatureConfigs, committee_verifier.SignatureConfig{
				SourceChainSelector: remoteSelector,
				Threshold:           remoteConfig.SignatureConfig.Threshold,
				Signers:             signers,
			})
			outboundImplementationArgs = append(outboundImplementationArgs, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteSelector,
				Verifier:          common.HexToAddress(committeeVerifier),
			})
		}

		// ApplyRemoteChainConfigUpdates on CommitteeVerifier
		committeeVerifierReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyRemoteChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates to CommitteeVerifier on chain %s: %w", chain, err)
		}
		writes = append(writes, committeeVerifierReport.Output)

		// ApplySignatureConfigs on CommitteeVerifier
		committeeVerifierSignatureConfigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplySignatureConfigs, chain, contract.FunctionInput[committee_verifier.SignatureConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args: committee_verifier.SignatureConfigArgs{
				SourceChainSelectorsToRemove: input.RemoteChainsToDisconnect,
				SignatureConfigUpdates:       signatureConfigs,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply signature configs to CommitteeVerifier on chain %s: %w", chain, err)
		}
		writes = append(writes, committeeVerifierSignatureConfigReport.Output)

		// Apply inbound implementation updates on CommitteeVerifierResolver
		// Get the version tag from the CommitteeVerifier
		committeeVerifierVersionTagReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetVersionTag, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CommitteeVerifier on chain %s: %w", chain, err)
		}
		committeeVerifierResolverInboundImplementationUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierResolver),
			Args: []versioned_verifier_resolver.InboundImplementationArgs{
				{
					Version:  committeeVerifierVersionTagReport.Output,
					Verifier: common.HexToAddress(committeeVerifier),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply inbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		writes = append(writes, committeeVerifierResolverInboundImplementationUpdatesReport.Output)

		// Apply outbound implementation updates on CommitteeVerifierResolver
		committeeVerifierResolverOutboundImplementationUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierResolver),
			Args:          outboundImplementationArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply outbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		writes = append(writes, committeeVerifierResolverOutboundImplementationUpdatesReport.Output)

		batchOps, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOps},
		}, nil
	})
