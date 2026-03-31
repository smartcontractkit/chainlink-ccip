package changesets

import (
	"fmt"
	"slices"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// CommitteeVerifierRemoteChainConfig is the user-facing input for a single remote chain on a
// CommitteeVerifier. It intentionally omits SignatureConfig because signers and threshold are
// derived automatically from the topology (NOPs × committee × chain) — the caller should never
// have to specify them manually.
type CommitteeVerifierRemoteChainConfig struct {
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               uint16
	GasForVerification        uint32
	PayloadSizeBytes          uint32
}

type CommitteeVerifierInputConfig struct {
	CommitteeQualifier string
	RemoteChains       map[uint64]CommitteeVerifierRemoteChainConfig
}

// PartialChainConfig describes the desired state for a single local chain. "Partial" because
// the user supplies datastore.AddressRef pointers, not resolved addresses — the changeset
// resolves them at apply time. Each chain can reference a different set of remote chains, so
// per-source/dest configuration (e.g. different fee quoter parameters per destination) is
// supported naturally.
type PartialChainConfig struct {
	ChainSelector      uint64
	Router             datastore.AddressRef
	OnRamp             datastore.AddressRef
	CommitteeVerifiers []CommitteeVerifierInputConfig
	FeeQuoter          datastore.AddressRef
	OffRamp            datastore.AddressRef
	RemoteChains       map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]
}

type ConfigureChainsForLanesFromTopologyConfig struct {
	Topology *offchain.EnvironmentTopology
	Chains   []PartialChainConfig
	MCMS     mcms.Input
}

// enrichedChainConfig is the internal representation of a chain config after the
// topology-level enrichment (signing keys, contract resolution). It mirrors PartialChainConfig
// but carries fully populated CommitteeVerifierConfig (with signers/threshold filled in).
type enrichedChainConfig struct {
	ChainSelector      uint64
	Router             datastore.AddressRef
	OnRamp             datastore.AddressRef
	CommitteeVerifiers []adapters.CommitteeVerifierConfig[datastore.AddressRef]
	FeeQuoter          datastore.AddressRef
	OffRamp            datastore.AddressRef
	RemoteChains       map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]
}

// ConfigureChainsForLanesFromTopology is the canonical changeset for configuring CCIP 2.0
// lanes. It is chain-centric: each entry in cfg.Chains describes one local chain and the set
// of remote chains it should communicate with.
//
// The changeset has three phases:
//  1. Enrichment — for each chain, fetch signing keys from JD (if missing in topology),
//     look up CommitteeVerifier contract addresses from the registry, and derive the
//     signature quorum config (signers + threshold) from the topology's NOP membership.
//  2. Resolution — convert datastore.AddressRef pointers into resolved addresses. Remote
//     contract addresses (OnRamp, OffRamp on the remote chain) are converted to []byte
//     via the remote chain family adapter's AddressRefToBytes, which handles cross-family
//     encoding (e.g. EVM 20-byte vs Solana 32-byte). Local contracts are resolved to
//     string addresses.
//  3. Dispatch — delegate to the chain family adapter's ConfigureChainForLanes sequence,
//     which performs the actual on-chain writes (idempotent, ordered, router-last).
//
// The three registries are injected at construction time rather than looked up from globals
// so the changeset is testable with mocks.
func ConfigureChainsForLanesFromTopology(
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
) deployment.ChangeSetV2[ConfigureChainsForLanesFromTopologyConfig] {
	validate := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) error {
		if committeeVerifierContractRegistry == nil {
			return fmt.Errorf("committee verifier contract registry is required")
		}
		if chainFamilyRegistry == nil {
			return fmt.Errorf("chain family registry is required")
		}
		if mcmsRegistry == nil {
			return fmt.Errorf("mcms registry is required")
		}
		if cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}
		if cfg.Topology.NOPTopology == nil || len(cfg.Topology.NOPTopology.Committees) == 0 {
			return fmt.Errorf("no committees defined in topology")
		}
		if len(cfg.Chains) == 0 {
			return fmt.Errorf("at least one chain must be specified")
		}
		for _, chainCfg := range cfg.Chains {
			if !slices.Contains(e.BlockChains.ListChainSelectors(), chainCfg.ChainSelector) {
				return fmt.Errorf("chain selector %d is not available in environment", chainCfg.ChainSelector)
			}
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) (deployment.ChangesetOutput, error) {
		// ── Phase 1: Enrichment ──────────────────────────────────────────────────
		// Pre-fetch signing keys only for NOPs that are actually referenced as
		// committee signers. Executor-only NOPs are excluded to avoid unnecessary
		// JD calls that could fail if those NOPs don't exist in the node registry.
		targetFamilies := deriveFamiliesFromChains(cfg.Chains)
		committeeNOPs := filterNOPsToCommitteeMembers(cfg.Topology.NOPTopology, cfg.Chains)
		signingKeysByNOP, err := fetchSigningKeysForNOPsByFamilies(e, committeeNOPs, targetFamilies)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to fetch signing keys: %w", err)
		}

		// For each chain, populate the CommitteeVerifier config with signers/threshold
		// derived from the topology and resolve the verifier contract addresses.
		chains := make([]enrichedChainConfig, 0, len(cfg.Chains))
		for _, chainCfg := range cfg.Chains {
			committeeVerifiers := make([]adapters.CommitteeVerifierConfig[datastore.AddressRef], 0, len(chainCfg.CommitteeVerifiers))
			for _, verifierInput := range chainCfg.CommitteeVerifiers {
				remoteChains := make(map[uint64]adapters.CommitteeVerifierRemoteChainConfig, len(verifierInput.RemoteChains))
				for remoteSelector, remoteConfig := range verifierInput.RemoteChains {
					// The signature config shaped in the local chain's context, however the set of signers and threshold are derived from the remote (source) chain
					signatureConfig, err := getSignatureConfigForLane(
						e,
						cfg.Topology,
						verifierInput.CommitteeQualifier,
						chainCfg.ChainSelector,
						remoteSelector,
						signingKeysByNOP,
					)
					if err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to get signature config for lane local chain %d -> remote chain %d: %w", chainCfg.ChainSelector, remoteSelector, err)
					}
					remoteChains[remoteSelector] = adapters.CommitteeVerifierRemoteChainConfig{
						AllowlistEnabled:          remoteConfig.AllowlistEnabled,
						AddedAllowlistedSenders:   remoteConfig.AddedAllowlistedSenders,
						RemovedAllowlistedSenders: remoteConfig.RemovedAllowlistedSenders,
						FeeUSDCents:               remoteConfig.FeeUSDCents,
						GasForVerification:        remoteConfig.GasForVerification,
						PayloadSizeBytes:          remoteConfig.PayloadSizeBytes,
						SignatureConfig:           *signatureConfig,
					}
				}

				contractAdapter, err := committeeVerifierContractRegistry.GetByChain(chainCfg.ChainSelector)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("no committee verifier contract adapter for chain %d: %w", chainCfg.ChainSelector, err)
				}
				// Resolve the committee verifier contracts for the local chain, using the committee qualifier from the verifier input
				contracts, err := contractAdapter.ResolveCommitteeVerifierContracts(e.DataStore, chainCfg.ChainSelector, verifierInput.CommitteeQualifier)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve committee verifier contracts for chain %d qualifier %q: %w", chainCfg.ChainSelector, verifierInput.CommitteeQualifier, err)
				}

				committeeVerifiers = append(committeeVerifiers, adapters.CommitteeVerifierConfig[datastore.AddressRef]{
					CommitteeVerifier: contracts,
					RemoteChains:      remoteChains,
				})
			}

			chains = append(chains, enrichedChainConfig{
				ChainSelector:      chainCfg.ChainSelector,
				Router:             chainCfg.Router,
				OnRamp:             chainCfg.OnRamp,
				CommitteeVerifiers: committeeVerifiers,
				FeeQuoter:          chainCfg.FeeQuoter,
				OffRamp:            chainCfg.OffRamp,
				RemoteChains:       chainCfg.RemoteChains,
			})
		}

		return applyConfigureChains(e, chainFamilyRegistry, mcmsRegistry, chains, cfg.MCMS)
	}

	return deployment.CreateChangeSet(apply, validate)
}

// applyConfigureChains performs Phase 2 (resolution) and Phase 3 (dispatch).
//
// Resolution: all datastore.AddressRef inputs are resolved to concrete addresses. Remote
// contracts are resolved via the remote family's AddressRefToBytes for cross-family encoding.
// Local contracts are resolved to their string address.
//
// Dispatch: for each chain, the resolved config is passed to the family adapter's
// ConfigureChainForLanes sequence. The sequence produces BatchOps (MCMS proposal
// transactions) or executes immediately in deployer-key-owned mode. BatchOps from all
// chains are aggregated and returned via the OutputBuilder which handles MCMS wrapping.
func applyConfigureChains(
	e deployment.Environment,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
	chains []enrichedChainConfig,
	mcmsInput mcms.Input,
) (deployment.ChangesetOutput, error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	ds := datastore.NewMemoryDataStore()

	for _, chainCfg := range chains {
		// ── Phase 2: Resolution ──────────────────────────────────────────────
		// Resolve local contract refs to their on-chain addresses via the datastore.
		router, err := datastore_utils.FindAndFormatRef(e.DataStore, chainCfg.Router, chainCfg.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve router ref on chain with selector %d: %w", chainCfg.ChainSelector, err)
		}
		onRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, chainCfg.OnRamp, chainCfg.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve onRamp ref on chain with selector %d: %w", chainCfg.ChainSelector, err)
		}
		feeQuoter, err := datastore_utils.FindAndFormatRef(e.DataStore, chainCfg.FeeQuoter, chainCfg.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve feeQuoter ref on chain with selector %d: %w", chainCfg.ChainSelector, err)
		}
		offRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, chainCfg.OffRamp, chainCfg.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve offRamp ref on chain with selector %d: %w", chainCfg.ChainSelector, err)
		}

		committeeVerifiers := make([]adapters.CommitteeVerifierConfig[datastore.AddressRef], len(chainCfg.CommitteeVerifiers))
		for i, verifier := range chainCfg.CommitteeVerifiers {
			contracts := make([]datastore.AddressRef, 0, len(verifier.CommitteeVerifier))
			for _, contractRef := range verifier.CommitteeVerifier {
				resolvedContract, err := datastore_utils.FindAndFormatRef(e.DataStore, contractRef, chainCfg.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve committee verifier contract ref on chain with selector %d: %w", chainCfg.ChainSelector, err)
				}
				contracts = append(contracts, resolvedContract)
			}
			committeeVerifiers[i] = adapters.CommitteeVerifierConfig[datastore.AddressRef]{
				CommitteeVerifier: contracts,
				RemoteChains:      verifier.RemoteChains,
			}
		}

		family, err := chainsel.GetSelectorFamily(chainCfg.ChainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainCfg.ChainSelector, err)
		}
		adapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return deployment.ChangesetOutput{}, fmt.Errorf("no adapter registered for chain family %q", family)
		}

		// Remote contract addresses (OnRamps, OffRamp) are resolved to []byte via the
		// remote chain family's AddressRefToBytes, which handles cross-family encoding
		// (e.g. EVM 20-byte vs Solana 32-byte). Local contracts are resolved to their
		// string address via the datastore.
		remoteChains := make(map[uint64]adapters.RemoteChainConfig[[]byte, string], len(chainCfg.RemoteChains))
		for remoteSelector, remoteChainCfg := range chainCfg.RemoteChains {
			remoteFamily, err := chainsel.GetSelectorFamily(remoteSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
			}
			remoteAdapter, ok := chainFamilyRegistry.GetChainFamily(remoteFamily)
			if !ok {
				return deployment.ChangesetOutput{}, fmt.Errorf("no adapter registered for remote chain family %q", remoteFamily)
			}

			convertedRemoteConfig, err := convertRemoteChainConfigForTopologyChangeset(e, chainCfg.ChainSelector, remoteSelector, remoteAdapter, remoteChainCfg)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to process remote chain config for selector %d: %w", remoteSelector, err)
			}
			remoteChains[remoteSelector] = convertedRemoteConfig
		}

		// ── Phase 3: Dispatch ──────────────────────────────────────────────
		// Delegate to the family-specific sequence which owns all on-chain writes.
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureChainForLanes(), e.BlockChains, adapters.ConfigureChainForLanesInput{
			ChainSelector:      chainCfg.ChainSelector,
			Router:             router.Address,
			OnRamp:             onRamp.Address,
			CommitteeVerifiers: committeeVerifiers,
			FeeQuoter:          feeQuoter.Address,
			OffRamp:            offRamp.Address,
			RemoteChains:       remoteChains,
		})
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to configure chain with selector %d: %w", chainCfg.ChainSelector, err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		for _, addr := range report.Output.Addresses {
			if err := ds.Addresses().Add(addr); err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", addr.Type, addr.Version, addr.Address, addr.ChainSelector, err)
			}
		}
		if err := sequences.WriteMetadataToDatastore(ds, report.Output.Metadata); err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to write metadata to datastore: %w", err)
		}
	}

	return changesetscore.NewOutputBuilder(e, mcmsRegistry).
		WithReports(reports).
		WithDataStore(ds).
		WithBatchOps(batchOps).
		Build(mcmsInput)
}

// convertRemoteChainConfigForTopologyChangeset transforms a remote chain config from
// datastore.AddressRef form (user input) to the resolved form expected by the sequence:
//   - Remote contracts (OnRamp, OffRamp) → []byte via the remote family adapter's
//     AddressRefToBytes, so cross-family encoding is handled transparently.
//   - Local contracts (Executor, CCVs) → string address via the environment datastore.
//
// This separation exists because remote contract addresses are stored on-chain as raw bytes
// (the local chain has no knowledge of the remote chain's address format), while local
// contracts are always addressed by their native string representation.
func convertRemoteChainConfigForTopologyChangeset(
	e deployment.Environment,
	chainSelector uint64,
	remoteChainSelector uint64,
	remoteChainAdapter adapters.ChainFamily,
	inCfg adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef],
) (adapters.RemoteChainConfig[[]byte, string], error) {
	outCfg := adapters.RemoteChainConfig[[]byte, string]{
		AllowTrafficFrom:          inCfg.AllowTrafficFrom,
		FeeQuoterDestChainConfig:  inCfg.FeeQuoterDestChainConfig,
		ExecutorDestChainConfig:   inCfg.ExecutorDestChainConfig,
		AddressBytesLength:        inCfg.AddressBytesLength,
		BaseExecutionGasCost:      inCfg.BaseExecutionGasCost,
		TokenReceiverAllowed:      inCfg.TokenReceiverAllowed,
		MessageNetworkFeeUSDCents: inCfg.MessageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   inCfg.TokenNetworkFeeUSDCents,
	}

	remoteOnRamps := make([][]byte, 0, len(inCfg.OnRamps))
	for _, onRampRef := range inCfg.OnRamps {
		resolvedOnRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, onRampRef, remoteChainSelector, remoteChainAdapter.AddressRefToBytes)
		if err != nil {
			return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve onRamp ref on remote chain with selector %d: %w", remoteChainSelector, err)
		}
		remoteOnRamps = append(remoteOnRamps, resolvedOnRamp)
	}
	outCfg.OnRamps = remoteOnRamps

	offRamp, err := datastore_utils.FindAndFormatRef(e.DataStore, inCfg.OffRamp, remoteChainSelector, remoteChainAdapter.AddressRefToBytes)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve offRamp ref on remote chain with selector %d: %w", remoteChainSelector, err)
	}
	outCfg.OffRamp = offRamp

	executor, err := datastore_utils.FindAndFormatRef(e.DataStore, inCfg.DefaultExecutor, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve executor ref on chain with selector %d: %w", chainSelector, err)
	}
	outCfg.DefaultExecutor = executor.Address

	defaultInboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, chainSelector, inCfg.DefaultInboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, err
	}
	outCfg.DefaultInboundCCVs = defaultInboundCCVs

	laneMandatedInboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, chainSelector, inCfg.LaneMandatedInboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, err
	}
	outCfg.LaneMandatedInboundCCVs = laneMandatedInboundCCVs

	defaultOutboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, chainSelector, inCfg.DefaultOutboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, err
	}
	outCfg.DefaultOutboundCCVs = defaultOutboundCCVs

	laneMandatedOutboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, chainSelector, inCfg.LaneMandatedOutboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, err
	}
	outCfg.LaneMandatedOutboundCCVs = laneMandatedOutboundCCVs

	return outCfg, nil
}

func resolveLocalContractsForTopologyChangeset(
	e deployment.Environment,
	chainSelector uint64,
	refs []datastore.AddressRef,
) ([]string, error) {
	resolved := make([]string, 0, len(refs))
	for _, ref := range refs {
		addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve contract ref on chain with selector %d: %w", chainSelector, err)
		}
		resolved = append(resolved, addr.Address)
	}
	return resolved, nil
}

// filterNOPsToCommitteeMembers returns only the NOPs whose aliases appear in at least one
// committee ChainConfig referenced by the input chains' CommitteeVerifiers. This prevents
// fetching signing keys for executor-only NOPs that will never be used as committee signers.
func filterNOPsToCommitteeMembers(nopTopology *offchain.NOPTopology, chains []PartialChainConfig) []offchain.NOPConfig {
	committeeAliases := make(map[string]struct{})
	for _, chainCfg := range chains {
		for _, cv := range chainCfg.CommitteeVerifiers {
			committee, ok := nopTopology.Committees[cv.CommitteeQualifier]
			if !ok {
				continue
			}
			for remoteSelector := range cv.RemoteChains {
				chainKey := strconv.FormatUint(remoteSelector, 10)
				if chainCommittee, ok := committee.ChainConfigs[chainKey]; ok {
					for _, alias := range chainCommittee.NOPAliases {
						committeeAliases[alias] = struct{}{}
					}
				}
			}
		}
	}
	filtered := make([]offchain.NOPConfig, 0, len(committeeAliases))
	for _, nop := range nopTopology.NOPs {
		if _, ok := committeeAliases[nop.Alias]; ok {
			filtered = append(filtered, nop)
		}
	}
	return filtered
}

// deriveFamiliesFromChains extracts the unique chain families referenced by the input chains.
func deriveFamiliesFromChains(chains []PartialChainConfig) []string {
	seen := make(map[string]struct{})
	for _, c := range chains {
		if family, err := chainsel.GetSelectorFamily(c.ChainSelector); err == nil {
			seen[family] = struct{}{}
		}
	}
	families := make([]string, 0, len(seen))
	for f := range seen {
		families = append(families, f)
	}
	return families
}
