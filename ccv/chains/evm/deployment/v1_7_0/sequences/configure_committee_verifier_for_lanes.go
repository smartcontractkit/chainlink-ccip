package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureCommitteeVerifierForLanesInput struct {
	ChainSelector uint64
	Router        string
	adapters.CommitteeVerifierConfig[datastore.AddressRef]
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
		for _, addr := range input.CommitteeVerifier {
			switch addr.Type {
			case datastore.ContractType(committee_verifier.ContractType):
				committeeVerifier = addr.Address
			case datastore.ContractType(committee_verifier.ResolverType):
				committeeVerifierResolver = addr.Address
			}
		}
		if committeeVerifier == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("committee verifier contract not found on chain %d", input.ChainSelector)
		}
		if committeeVerifierResolver == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("committee verifier resolver contract not found on chain %d", input.ChainSelector)
		}

		remoteChainConfigArgs := make([]committee_verifier.RemoteChainConfigArgs, 0, len(input.RemoteChains))
		allowlistArgs := make([]committee_verifier.AllowlistConfigArgs, 0, len(input.RemoteChains))
		signatureConfigs := make([]committee_verifier.SignatureConfig, 0, len(input.RemoteChains))

		for remoteSelector, remoteConfig := range input.RemoteChains {
			desiredRemoteChainArg := committee_verifier.RemoteChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				RemoteChainSelector: remoteSelector,
				AllowlistEnabled:    remoteConfig.AllowlistEnabled,
				FeeUSDCents:         remoteConfig.FeeUSDCents,
				GasForVerification:  remoteConfig.GasForVerification,
				PayloadSizeBytes:    remoteConfig.PayloadSizeBytes,
			}
			// Only append when remote chain config or allowlist differs from on-chain.
			currentRemoteReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote chain config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
			}
			currRemote := currentRemoteReport.Output

			if currRemote.Router != desiredRemoteChainArg.Router ||
				currRemote.AllowlistEnabled != desiredRemoteChainArg.AllowlistEnabled {
				remoteChainConfigArgs = append(remoteChainConfigArgs, desiredRemoteChainArg)
			} else {
				// We have to check these parameters separately because if we call GetFee before the router is set, it will revert.
				getFeeReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetFee, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(committeeVerifier),
					Args:          remoteSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
				}
				currFee := getFeeReport.Output

				if currFee.FeeUSDCents != desiredRemoteChainArg.FeeUSDCents ||
					currFee.GasForVerification != desiredRemoteChainArg.GasForVerification ||
					currFee.PayloadSizeBytes != desiredRemoteChainArg.PayloadSizeBytes {
					remoteChainConfigArgs = append(remoteChainConfigArgs, desiredRemoteChainArg)
				}
			}

			toAdd, toRemove := makeAllowlistUpdates(currRemote.AllowedSendersList, remoteConfig.AddedAllowlistedSenders, remoteConfig.RemovedAllowlistedSenders)
			allowlistDiffers := len(toAdd) > 0 || len(toRemove) > 0
			if allowlistDiffers {
				allowlistArgs = append(allowlistArgs, committee_verifier.AllowlistConfigArgs{
					AllowlistEnabled:          remoteConfig.AllowlistEnabled,
					AddedAllowlistedSenders:   toAdd,
					RemovedAllowlistedSenders: toRemove,
					DestChainSelector:         remoteSelector,
				})
			}

			signers := make([]common.Address, 0, len(remoteConfig.SignatureConfig.Signers))
			for _, signer := range remoteConfig.SignatureConfig.Signers {
				signers = append(signers, common.HexToAddress(signer))
			}
			desiredSig := committee_verifier.SignatureConfig{
				SourceChainSelector: remoteSelector,
				Threshold:           remoteConfig.SignatureConfig.Threshold,
				Signers:             signers,
			}
			currentSigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetSignatureConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get signature config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
			}
			curSig := currentSigReport.Output
			if curSig.Threshold != desiredSig.Threshold || !UnorderedSliceEqual(curSig.Signers, desiredSig.Signers, func(x, y common.Address) bool { return x == y }) {
				signatureConfigs = append(signatureConfigs, desiredSig)
			}
		}

		// Build outbound implementation args only for remotes where on-chain verifier differs (idempotent).
		currentOutboundReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.GetAllOutboundImplementations, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierResolver),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound implementations from CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		currentOutbound := make(map[uint64]common.Address)
		for _, o := range currentOutboundReport.Output {
			currentOutbound[o.DestChainSelector] = o.Verifier
		}
		outboundImplementationArgs := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0, len(input.RemoteChains))
		committeeVerifierAddr := common.HexToAddress(committeeVerifier)
		for remoteSelector := range input.RemoteChains {
			if currentOutbound[remoteSelector] != committeeVerifierAddr {
				outboundImplementationArgs = append(outboundImplementationArgs, versioned_verifier_resolver.OutboundImplementationArgs{
					DestChainSelector: remoteSelector,
					Verifier:          committeeVerifierAddr,
				})
			}
		}

		// Apply remote chain config updates only when there are changes.
		if len(remoteChainConfigArgs) > 0 {
			committeeVerifierReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyRemoteChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.RemoteChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteChainConfigArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierReport.Output)
		}

		// Apply allowlist updates only when there are changes.
		if len(allowlistArgs) > 0 {
			committeeVerifierAllowlistReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]committee_verifier.AllowlistConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          allowlistArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierAllowlistReport.Output)
		}

		// Apply signature configs only when there are changes.
		if len(signatureConfigs) > 0 {
			committeeVerifierSignatureConfigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplySignatureConfigs, chain, contract.FunctionInput[committee_verifier.SignatureConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args: committee_verifier.SignatureConfigArgs{
					SignatureConfigUpdates: signatureConfigs,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply signature configs to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierSignatureConfigReport.Output)
		}

		// Apply inbound implementation updates on CommitteeVerifierResolver only when not already set (idempotent).
		committeeVerifierVersionTagReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetVersionTag, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CommitteeVerifier on chain %s: %w", chain, err)
		}
		desiredInbound := versioned_verifier_resolver.InboundImplementationArgs{
			Version:  committeeVerifierVersionTagReport.Output,
			Verifier: common.HexToAddress(committeeVerifier),
		}
		currentInboundReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.GetAllInboundImplementations, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierResolver),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound implementations from CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		inboundAlreadySet := false
		for _, cur := range currentInboundReport.Output {
			if cur.Version == desiredInbound.Version && cur.Verifier == desiredInbound.Verifier {
				inboundAlreadySet = true
				break
			}
		}
		if !inboundAlreadySet {
			committeeVerifierResolverInboundImplementationUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifierResolver),
				Args:          []versioned_verifier_resolver.InboundImplementationArgs{desiredInbound},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply inbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierResolverInboundImplementationUpdatesReport.Output)
		}

		// Apply outbound implementation updates on CommitteeVerifierResolver only when there are changes (idempotent).
		if len(outboundImplementationArgs) > 0 {
			committeeVerifierResolverOutboundImplementationUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifierResolver),
				Args:          outboundImplementationArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply outbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierResolverOutboundImplementationUpdatesReport.Output)
		}

		batchOps, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOps},
		}, nil
	})

// makeAllowlistUpdates returns the adds and removes to apply so the allowlist becomes (current ∪ added) \ removed.
// It takes the current on-chain allowlist and the config's added/removed sender lists (hex addresses).
func makeAllowlistUpdates(current []common.Address, added, removed []string) (toAdd, toRemove []common.Address) {
	curSet := make(map[common.Address]struct{}, len(current))
	for _, a := range current {
		curSet[a] = struct{}{}
	}
	addedSet := make(map[common.Address]struct{}, len(added))
	for _, s := range added {
		addedSet[common.HexToAddress(s)] = struct{}{}
	}
	removedSet := make(map[common.Address]struct{}, len(removed))
	for _, s := range removed {
		removedSet[common.HexToAddress(s)] = struct{}{}
	}
	desiredSet := make(map[common.Address]struct{})
	for a := range curSet {
		if _, remove := removedSet[a]; !remove {
			desiredSet[a] = struct{}{}
		}
	}
	for a := range addedSet {
		desiredSet[a] = struct{}{}
	}
	desired := make([]common.Address, 0, len(desiredSet))
	for a := range desiredSet {
		desired = append(desired, a)
	}
	toAdd = AddressesNotIn(desired, current)
	toRemove = AddressesNotIn(current, desired)
	return toAdd, toRemove
}
