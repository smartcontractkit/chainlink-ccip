package sequences

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	vvr_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	cv_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
)

type ConfigureCommitteeVerifierAsSourceInput struct {
	ChainSelector uint64
	Router        string
	adapters.CommitteeVerifierConfig[datastore.AddressRef]
}

type ConfigureCommitteeVerifierAsDestInput struct {
	ChainSelector uint64
	adapters.CommitteeVerifierConfig[datastore.AddressRef]
}

func extractCommitteeVerifierAddresses(refs []datastore.AddressRef, chainSelector uint64) (verifier string, resolver string, err error) {
	for _, addr := range refs {
		switch addr.Type {
		case datastore.ContractType(committee_verifier.ContractType):
			if verifier != "" {
				return "", "", fmt.Errorf("duplicate committee verifier contract on chain %d", chainSelector)
			}
			verifier = addr.Address
		case datastore.ContractType(CommitteeVerifierResolverType):
			if resolver != "" {
				return "", "", fmt.Errorf("duplicate committee verifier resolver contract on chain %d", chainSelector)
			}
			resolver = addr.Address
		}
	}
	if verifier == "" {
		return "", "", fmt.Errorf("committee verifier contract not found on chain %d", chainSelector)
	}
	if resolver == "" {
		return "", "", fmt.Errorf("committee verifier resolver contract not found on chain %d", chainSelector)
	}
	return verifier, resolver, nil
}

// ConfigureCommitteeVerifierAsSource configures outbound CommitteeVerifier settings for remote chains:
// remote chain config (router, fees, allowlist flag), allowlist updates, and outbound resolver implementation.
// Deprecated: Use ConfigureChainForLanes from the ChainFamilyAdapter instead this is the canonical way to configure lanes for 2.0. This will be unexported in the future to be used internally only.
var ConfigureCommitteeVerifierAsSource = cldf_ops.NewSequence(
	"configure-committee-verifier-as-source",
	semver.MustParse("2.0.0"),
	"Configures outbound CommitteeVerifier settings for remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input ConfigureCommitteeVerifierAsSourceInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		committeeVerifier, committeeVerifierResolver, err := extractCommitteeVerifierAddresses(input.CommitteeVerifier, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		remoteChainConfigArgs := make([]cv_bindings.BaseVerifierRemoteChainConfigArgs, 0, len(input.RemoteChains))
		allowlistArgs := make([]cv_bindings.BaseVerifierAllowlistConfigArgs, 0, len(input.RemoteChains))

		for remoteSelector, remoteConfig := range input.RemoteChains {
			desiredRemoteChainArg := cv_bindings.BaseVerifierRemoteChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				RemoteChainSelector: remoteSelector,
				AllowlistEnabled:    remoteConfig.AllowlistEnabled,
				FeeUSDCents:         remoteConfig.FeeUSDCents,
				GasForVerification:  remoteConfig.GasForVerification,
				// TODO : standardise payload size
				PayloadSizeBytes: remoteConfig.PayloadSizeBytes,
			}
			currentRemoteReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadGetRemoteChainConfig, remoteSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote chain config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
			}
			currRemote := currentRemoteReport.Output

			if currRemote.RemoteChainConfig.Router != desiredRemoteChainArg.Router ||
				currRemote.RemoteChainConfig.AllowlistEnabled != desiredRemoteChainArg.AllowlistEnabled {
				remoteChainConfigArgs = append(remoteChainConfigArgs, desiredRemoteChainArg)
			} else {
				getFeeReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadGetFee, committee_verifier.GetFeeArgs{
					DestChainSelector: remoteSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
				}
				currFee := getFeeReport.Output

				if currFee.FeeUSDCents != desiredRemoteChainArg.FeeUSDCents ||
					currFee.GasForVerification != desiredRemoteChainArg.GasForVerification ||
					uint16(currFee.PayloadSizeBytes) != desiredRemoteChainArg.PayloadSizeBytes {
					remoteChainConfigArgs = append(remoteChainConfigArgs, desiredRemoteChainArg)
				}
			}

			toAdd, toRemove, err := makeAllowlistUpdates(currRemote.AllowedSendersList, remoteConfig.AddedAllowlistedSenders, remoteConfig.RemovedAllowlistedSenders)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid allowlist addresses for remote chain %d: %w", remoteSelector, err)
			}
			if len(toAdd) > 0 || len(toRemove) > 0 {
				allowlistArgs = append(allowlistArgs, cv_bindings.BaseVerifierAllowlistConfigArgs{
					AllowlistEnabled:          remoteConfig.AllowlistEnabled,
					AddedAllowlistedSenders:   toAdd,
					RemovedAllowlistedSenders: toRemove,
					DestChainSelector:         remoteSelector,
				})
			}
		}

		// Build outbound implementation args only for remotes where on-chain verifier differs (idempotent).
		currentOutboundReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifierResolver), vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewReadGetAllOutboundImplementations, nil)
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

		if len(remoteChainConfigArgs) > 0 {
			committeeVerifierReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewWriteApplyRemoteChainConfigUpdates, remoteChainConfigArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply remote chain config updates to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierReport.Output)
		}

		if len(allowlistArgs) > 0 {
			committeeVerifierAllowlistReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewWriteApplyAllowlistUpdates, allowlistArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, committeeVerifierAllowlistReport.Output)
		}

		if len(outboundImplementationArgs) > 0 {
			outboundReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifierResolver), vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewWriteApplyOutboundImplementationUpdates, outboundImplementationArgs)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply outbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
			}
			writes = append(writes, outboundReport.Output)
		}

		if !input.AllowedFinalityConfig.IsZero() {
			desiredFinality := input.AllowedFinalityConfig.Raw()
			currentFinalityReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadGetAllowedFinalityConfig, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed finality config from CommitteeVerifier on chain %s: %w", chain, err)
			}
			if currentFinalityReport.Output != desiredFinality {
				setFinalityReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewWriteSetAllowedFinalityConfig, desiredFinality)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set allowed finality config on CommitteeVerifier on chain %s: %w", chain, err)
				}
				writes = append(writes, setFinalityReport.Output)
			}
		}

		batchOps, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOps},
		}, nil
	})

// ConfigureCommitteeVerifierAsDest configures inbound CommitteeVerifier settings for remote chains:
// signature config (signers, threshold) and inbound resolver implementation.
// Deprecated: Use ConfigureChainForLanes from the ChainFamilyAdapter instead this is the canonical way to configure lanes for 2.0. This will be unexported in the future to be used internally only.
var ConfigureCommitteeVerifierAsDest = cldf_ops.NewSequence(
	"configure-committee-verifier-as-dest",
	semver.MustParse("2.0.0"),
	"Configures inbound CommitteeVerifier settings for remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input ConfigureCommitteeVerifierAsDestInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		committeeVerifier, committeeVerifierResolver, err := extractCommitteeVerifierAddresses(input.CommitteeVerifier, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		signatureConfigs := make([]cv_bindings.SignatureQuorumValidatorSignatureConfig, 0, len(input.RemoteChains))

		for remoteSelector, remoteConfig := range input.RemoteChains {
			signers := make([]common.Address, 0, len(remoteConfig.SignatureConfig.Signers))
			for _, signer := range remoteConfig.SignatureConfig.Signers {
				signers = append(signers, common.HexToAddress(signer))
			}
			desiredSig := cv_bindings.SignatureQuorumValidatorSignatureConfig{
				SourceChainSelector: remoteSelector,
				Threshold:           remoteConfig.SignatureConfig.Threshold,
				Signers:             signers,
			}
			currentSigReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadGetSignatureConfig, remoteSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get signature config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
			}
			curSig := currentSigReport.Output
			if curSig.Threshold != desiredSig.Threshold || !UnorderedSliceEqual(curSig.Signers, desiredSig.Signers, func(x, y common.Address) bool { return x == y }) {
				signatureConfigs = append(signatureConfigs, desiredSig)
			}
		}

		if len(signatureConfigs) > 0 {
			signatureConfigReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewWriteApplySignatureConfigs, committee_verifier.ApplySignatureConfigsArgs{
				SignatureConfigs: signatureConfigs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply signature configs to CommitteeVerifier on chain %s: %w", chain, err)
			}
			writes = append(writes, signatureConfigReport.Output)
		}

		// Apply inbound implementation updates on CommitteeVerifierResolver only when not already set (idempotent).
		committeeVerifierVersionTagReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifier), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadVersionTag, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CommitteeVerifier on chain %s: %w", chain, err)
		}
		desiredInbound := versioned_verifier_resolver.InboundImplementationArgs{
			Version:  committeeVerifierVersionTagReport.Output,
			Verifier: common.HexToAddress(committeeVerifier),
		}
		currentInboundReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifierResolver), vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewReadGetAllInboundImplementations, nil)
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
			inboundReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifierResolver), vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewWriteApplyInboundImplementationUpdates, []versioned_verifier_resolver.InboundImplementationArgs{desiredInbound})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply inbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
			}
			writes = append(writes, inboundReport.Output)
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
func makeAllowlistUpdates(current []common.Address, added, removed []string) (toAdd, toRemove []common.Address, err error) {
	curSet := make(map[common.Address]struct{}, len(current))
	for _, a := range current {
		curSet[a] = struct{}{}
	}
	addedSet := make(map[common.Address]struct{}, len(added))
	for _, s := range added {
		if !common.IsHexAddress(s) {
			return nil, nil, fmt.Errorf("invalid hex address in added allowlist: %q", s)
		}
		addedSet[common.HexToAddress(s)] = struct{}{}
	}
	removedSet := make(map[common.Address]struct{}, len(removed))
	for _, s := range removed {
		if !common.IsHexAddress(s) {
			return nil, nil, fmt.Errorf("invalid hex address in removed allowlist: %q", s)
		}
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
	return toAdd, toRemove, nil
}
