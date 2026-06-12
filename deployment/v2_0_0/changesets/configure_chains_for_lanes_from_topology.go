package changesets

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"k8s.io/utils/ptr"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

const defaultQualifier = "default"

// committeeVerifierRemoteChainInput is the internal per-remote CommitteeVerifier input before
// topology enrichment fills SignatureConfig.
type committeeVerifierRemoteChainInput struct {
	AllowlistEnabled          *bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               *uint16
	GasForVerification        *uint32
	PayloadSizeBytes          *uint16
}

type committeeVerifierInputConfig struct {
	CommitteeQualifier    string
	RemoteChains          map[uint64]committeeVerifierRemoteChainInput
	AllowedFinalityConfig *finality.Config
}

// PartialRemoteChainConfig is the internal per-remote chain input after lane expansion.
// Unset fields use adapter defaults, datastore resolution, or CCV auto-resolution.
type PartialRemoteChainConfig struct {
	AllowTrafficFrom          *bool
	DefaultExecutorQualifier  *string
	DefaultInboundCCVs        []datastore.AddressRef
	LaneMandatedInboundCCVs   []datastore.AddressRef
	DefaultOutboundCCVs       []datastore.AddressRef
	LaneMandatedOutboundCCVs  []datastore.AddressRef
	FeeQuoterDestChainConfig  adapters.FeeQuoterDestChainConfigOverrides
	ExecutorDestChainConfig   *adapters.ExecutorDestChainConfig
	BaseExecutionGasCost      *uint32
	TokenReceiverAllowed      *bool
	MessageNetworkFeeUSDCents *uint16
	TokenNetworkFeeUSDCents   *uint16
}

type partialChainConfig struct {
	ChainSelector      uint64
	CommitteeVerifiers []committeeVerifierInputConfig
	RemoteChains       map[uint64]PartialRemoteChainConfig
	FamilyExtras       map[string]any
}

// enrichedChainConfig is the internal representation after topology enrichment.
type enrichedChainConfig struct {
	ChainSelector      uint64
	CommitteeVerifiers []adapters.CommitteeVerifierConfig[datastore.AddressRef]
	RemoteChains       map[uint64]PartialRemoteChainConfig
	FamilyExtras       map[string]any
}

// ConfigureChainsForLanesFromTopology is the canonical changeset for configuring CCIP 2.0
// lanes from bidirectional lane pairs and topology.
//
// The changeset has three phases:
//  1. Enrichment — for each chain, fetch signing keys from JD (if missing in topology),
//     look up CommitteeVerifier contract addresses from the registry, and derive the
//     signature quorum config (signers + threshold) from the topology's NOP membership.
//  2. Resolution — convert datastore.AddressRef pointers into resolved addresses. Remote
//     contract addresses (OnRamp, OffRamp on the remote chain) are resolved via the remote
//     chain family adapter (e.g. GetOnRampAddress / GetOffRampAddress), which handles any
//     required cross-family encoding (e.g. EVM 20-byte vs Solana 32-byte). Local contracts
//     are resolved to string addresses.
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
		if err := cfg.Topology.ValidateForEnvironment(e.Name, chainFamilyRegistry); err != nil {
			return fmt.Errorf("topology validation failed: %w", err)
		}
		if cfg.Topology.NOPTopology == nil || len(cfg.Topology.NOPTopology.Committees) == 0 {
			return fmt.Errorf("no committees defined in topology")
		}
		chains, err := expandLanesToPartialChainConfigs(cfg.Lanes, cfg.Topology.NOPTopology.Committees)
		if err != nil {
			return err
		}
		chains = filterPartialChainsToEnvironment(e.BlockChains, chains)
		if len(chains) == 0 {
			return fmt.Errorf("no lane chains are available in environment")
		}
		for _, chainCfg := range chains {
			if !slices.Contains(e.BlockChains.ListChainSelectors(), chainCfg.ChainSelector) {
				return fmt.Errorf("chain selector %d is not available in environment", chainCfg.ChainSelector)
			}
			for _, cv := range chainCfg.CommitteeVerifiers {
				if cv.AllowedFinalityConfig != nil {
					if err := cv.AllowedFinalityConfig.Validate(); err != nil {
						return fmt.Errorf("invalid AllowedFinalityConfig for committee verifier on chain %d: %w", chainCfg.ChainSelector, err)
					}
				}
			}
			if err := validateDefaultCCVsResolvable(chainCfg, committeeVerifierContractRegistry, e); err != nil {
				return err
			}
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) (deployment.ChangesetOutput, error) {
		chains, err := expandLanesToPartialChainConfigs(cfg.Lanes, cfg.Topology.NOPTopology.Committees)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
		chains = filterPartialChainsToEnvironment(e.BlockChains, chains)
		if len(chains) == 0 {
			return deployment.ChangesetOutput{}, fmt.Errorf("no lane chains are available in environment")
		}

		// ── Phase 1: Enrichment ──────────────────────────────────────────────────
		// Pre-fetch signing keys only for NOPs that are actually referenced as
		// committee signers. Executor-only NOPs are excluded to avoid unnecessary
		// JD calls that could fail if those NOPs don't exist in the node registry.
		targetFamilies := deriveFamiliesFromChains(chains)
		committeeNOPs := filterNOPsToCommitteeMembers(cfg.Topology.NOPTopology, chains)
		signingKeysByNOP, err := fetchSigningKeysForNOPsByFamilies(e, committeeNOPs, targetFamilies)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to fetch signing keys: %w", err)
		}

		// For each chain, populate the CommitteeVerifier config with signers/threshold
		// derived from the topology and resolve the verifier contract addresses.
		enriched := make([]enrichedChainConfig, 0, len(chains))
		for _, chainCfg := range chains {
			localFamily, err := chainsel.GetSelectorFamily(chainCfg.ChainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for local chain %d: %w", chainCfg.ChainSelector, err)
			}
			localAdapter, ok := chainFamilyRegistry.GetChainFamily(localFamily)
			if !ok {
				return deployment.ChangesetOutput{}, fmt.Errorf("no adapter registered for local chain family %q", localFamily)
			}

			committeeVerifiers := make([]adapters.CommitteeVerifierConfig[datastore.AddressRef], 0, len(chainCfg.CommitteeVerifiers))
			for _, verifierInput := range chainCfg.CommitteeVerifiers {
				remoteChains := make(map[uint64]adapters.CommitteeVerifierRemoteChainConfig, len(verifierInput.RemoteChains))
				for remoteSelector, remoteConfig := range verifierInput.RemoteChains {
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

					remoteFamily, err := chainsel.GetSelectorFamily(remoteSelector)
					if err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain %d: %w", remoteSelector, err)
					}
					remoteAdapter, ok := chainFamilyRegistry.GetChainFamily(remoteFamily)
					if !ok {
						return deployment.ChangesetOutput{}, fmt.Errorf("no adapter registered for remote chain family %q (remote %d)", remoteFamily, remoteSelector)
					}

					remoteChains[remoteSelector] = mergeCommitteeVerifierRemoteChainConfig(
						remoteAdapter.GetDefaultCommitteeVerifierRemoteChainConfig(),
						remoteConfig,
						*signatureConfig,
					)
				}

				contractAdapter, err := committeeVerifierContractRegistry.GetByChain(chainCfg.ChainSelector)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("no committee verifier contract adapter for chain %d: %w", chainCfg.ChainSelector, err)
				}
				contracts, err := contractAdapter.ResolveCommitteeVerifierContracts(e.DataStore, chainCfg.ChainSelector, verifierInput.CommitteeQualifier)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve committee verifier contracts for chain %d qualifier %q: %w", chainCfg.ChainSelector, verifierInput.CommitteeQualifier, err)
				}

				allowedFinality := localAdapter.GetDefaultFinalityConfig()
				if verifierInput.AllowedFinalityConfig != nil {
					allowedFinality = *verifierInput.AllowedFinalityConfig
				}
				if err := allowedFinality.Validate(); err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("invalid effective AllowedFinalityConfig for committee verifier on chain %d: %w", chainCfg.ChainSelector, err)
				}

				committeeVerifiers = append(committeeVerifiers, adapters.CommitteeVerifierConfig[datastore.AddressRef]{
					CommitteeVerifier:     contracts,
					RemoteChains:          remoteChains,
					AllowedFinalityConfig: allowedFinality,
				})
			}

			enriched = append(enriched, enrichedChainConfig{
				ChainSelector:      chainCfg.ChainSelector,
				CommitteeVerifiers: committeeVerifiers,
				RemoteChains:       chainCfg.RemoteChains,
				FamilyExtras:       chainCfg.FamilyExtras,
			})
		}

		return applyConfigureChains(e, chainFamilyRegistry, mcmsRegistry, committeeVerifierContractRegistry, enriched, cfg.MCMS, cfg.UseTestRouter())
	}

	return deployment.CreateChangeSet(apply, validate)
}

// applyConfigureChains performs Phase 2 (resolution) and Phase 3 (dispatch).
//
// Resolution: well-known local contracts (Router, OnRamp, FeeQuoter, OffRamp) are resolved
// automatically from the datastore via the chain family adapter. Remote contracts (OnRamp,
// OffRamp) are resolved via the remote family adapter. The Executor is resolved per-lane
// using the ExecutorQualifier.
//
// Dispatch: for each chain, the resolved config is passed to the family adapter's
// ConfigureChainForLanes sequence. The sequence produces BatchOps (MCMS proposal
// transactions) or executes immediately in deployer-key-owned mode. BatchOps from all
// chains are aggregated and returned via the OutputBuilder which handles MCMS wrapping.
func applyConfigureChains(
	e deployment.Environment,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
	chains []enrichedChainConfig,
	mcmsInput mcms.Input,
	useTestRouter bool,
) (deployment.ChangesetOutput, error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	ds := datastore.NewMemoryDataStore()

	for _, chainCfg := range chains {
		// ── Phase 2: Resolution ──────────────────────────────────────────────
		family, err := chainsel.GetSelectorFamily(chainCfg.ChainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainCfg.ChainSelector, err)
		}
		adapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return deployment.ChangesetOutput{}, fmt.Errorf("no adapter registered for chain family %q", family)
		}

		var routerBytes []byte
		if useTestRouter {
			routerBytes, err = adapter.GetTestRouter(e.DataStore, chainCfg.ChainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve test router on chain %d: %w", chainCfg.ChainSelector, err)
			}
		} else {
			routerBytes, err = adapter.GetRouterAddress(e.DataStore, chainCfg.ChainSelector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve router on chain %d: %w", chainCfg.ChainSelector, err)
			}
		}

		onRampBytes, err := adapter.GetOnRampAddress(e.DataStore, chainCfg.ChainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve onRamp on chain %d: %w", chainCfg.ChainSelector, err)
		}
		feeQuoterBytes, err := adapter.GetFQAddress(e.DataStore, chainCfg.ChainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve feeQuoter on chain %d: %w", chainCfg.ChainSelector, err)
		}
		offRampBytes, err := adapter.GetOffRampAddress(e.DataStore, chainCfg.ChainSelector)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve offRamp on chain %d: %w", chainCfg.ChainSelector, err)
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
				CommitteeVerifier:     contracts,
				RemoteChains:          verifier.RemoteChains,
				AllowedFinalityConfig: verifier.AllowedFinalityConfig,
			}
		}

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

			convertedRemoteConfig, err := resolveRemoteChainConfig(e, adapter, remoteAdapter, chainCfg.ChainSelector, remoteSelector, remoteChainCfg, committeeVerifierContractRegistry)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to process remote chain config for selector %d: %w", remoteSelector, err)
			}
			remoteChains[remoteSelector] = convertedRemoteConfig
		}

		// ── Phase 3: Dispatch ──────────────────────────────────────────────
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigureChainForLanes(), e.BlockChains, adapters.ConfigureChainForLanesInput{
			ChainSelector:       chainCfg.ChainSelector,
			AllowOnrampOverride: useTestRouter,
			Router:              routerBytes,
			OnRamp:              onRampBytes,
			CommitteeVerifiers:  committeeVerifiers,
			FeeQuoter:           feeQuoterBytes,
			OffRamp:             offRampBytes,
			RemoteChains:        remoteChains,
			FamilyExtras:        chainCfg.FamilyExtras,
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

// resolveRemoteChainConfig resolves a PartialRemoteChainConfig into the form expected by the
// sequence. Remote contracts (OnRamp, OffRamp) are resolved via the remote chain's adapter.
// Local contracts (Executor, CCVs) are resolved from the local chain's datastore/adapter.
// Fields not explicitly set in inCfg fall back to the remote adapter's GetDefaultRemoteChainConfig.
func resolveRemoteChainConfig(
	e deployment.Environment,
	localAdapter adapters.ChainFamily,
	remoteAdapter adapters.ChainFamily,
	localChainSelector uint64,
	remoteChainSelector uint64,
	inCfg PartialRemoteChainConfig,
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
) (adapters.RemoteChainConfig[[]byte, string], error) {
	remoteOnRampBytes, err := remoteAdapter.GetOnRampAddress(e.DataStore, remoteChainSelector)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve remote onRamp on chain %d: %w", remoteChainSelector, err)
	}
	remoteOffRampBytes, err := remoteAdapter.GetOffRampAddress(e.DataStore, remoteChainSelector)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve remote offRamp on chain %d: %w", remoteChainSelector, err)
	}

	executorQualifier := inCfg.DefaultExecutorQualifier
	if executorQualifier == nil {
		executorQualifier = ptr.To(defaultQualifier)
	}
	executorAddr, err := localAdapter.ResolveExecutor(e.DataStore, localChainSelector, *executorQualifier)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve executor (qualifier %q) on chain %d: %w", *executorQualifier, localChainSelector, err)
	}

	defaultInboundCCVs, err := resolveDefaultCCVs(e, localChainSelector, inCfg.DefaultInboundCCVs, committeeVerifierContractRegistry)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve default inbound CCVs on chain %d: %w", localChainSelector, err)
	}
	laneMandatedInboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, localChainSelector, inCfg.LaneMandatedInboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve lane-mandated inbound CCVs on chain %d: %w", localChainSelector, err)
	}
	defaultOutboundCCVs, err := resolveDefaultCCVs(e, localChainSelector, inCfg.DefaultOutboundCCVs, committeeVerifierContractRegistry)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve default outbound CCVs on chain %d: %w", localChainSelector, err)
	}
	laneMandatedOutboundCCVs, err := resolveLocalContractsForTopologyChangeset(e, localChainSelector, inCfg.LaneMandatedOutboundCCVs)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve lane-mandated outbound CCVs on chain %d: %w", localChainSelector, err)
	}

	fqConfig := mergeFeeQuoterDestChainConfig(
		remoteAdapter.GetDefaultFeeQuoterDestChainConfig(localChainSelector, remoteChainSelector, remoteAdapter.GetChainFamilySelector()),
		inCfg.FeeQuoterDestChainConfig,
	)
	// ChainFamilySelector is always authoritative from the remote adapter, regardless of
	// what the defaults or user-supplied overrides contain.
	fqConfig.ChainFamilySelector = remoteAdapter.GetChainFamilySelector()

	defaults := remoteAdapter.GetDefaultRemoteChainConfig(localChainSelector, remoteChainSelector)

	// Apply per-field user overrides on top of adapter defaults via coalesce.
	// Pointer fields: nil = keep adapter default; non-nil = use caller value.
	allowTrafficFrom := utils.Coalesce(inCfg.AllowTrafficFrom, defaults.AllowTrafficFrom)
	tokenReceiverAllowed := utils.Coalesce(inCfg.TokenReceiverAllowed, defaults.TokenReceiverAllowed)

	return adapters.RemoteChainConfig[[]byte, string]{
		AllowTrafficFrom:          &allowTrafficFrom,
		OnRamps:                   [][]byte{remoteOnRampBytes},
		OffRamp:                   remoteOffRampBytes,
		DefaultExecutor:           executorAddr,
		DefaultInboundCCVs:        defaultInboundCCVs,
		LaneMandatedInboundCCVs:   laneMandatedInboundCCVs,
		DefaultOutboundCCVs:       defaultOutboundCCVs,
		LaneMandatedOutboundCCVs:  laneMandatedOutboundCCVs,
		FeeQuoterDestChainConfig:  fqConfig,
		ExecutorDestChainConfig:   utils.Coalesce(inCfg.ExecutorDestChainConfig, defaults.ExecutorDestChainConfig),
		AddressBytesLength:        remoteAdapter.GetAddressBytesLength(),
		BaseExecutionGasCost:      utils.Coalesce(inCfg.BaseExecutionGasCost, defaults.BaseExecutionGasCost),
		TokenReceiverAllowed:      &tokenReceiverAllowed,
		MessageNetworkFeeUSDCents: utils.Coalesce(inCfg.MessageNetworkFeeUSDCents, defaults.MessageNetworkFeeUSDCents),
		TokenNetworkFeeUSDCents:   utils.Coalesce(inCfg.TokenNetworkFeeUSDCents, defaults.TokenNetworkFeeUSDCents),
	}, nil
}

// resolveDefaultCCVs resolves CCVs either from explicit refs or by auto-resolving
// default lane CCV refs for the "default" qualifier (see CommitteeVerifierContractAdapter.GetCommitteeVerifierResolver).
func resolveDefaultCCVs(
	e deployment.Environment,
	chainSelector uint64,
	explicitRefs []datastore.AddressRef,
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
) ([]string, error) {
	if len(explicitRefs) > 0 {
		return resolveLocalContractsForTopologyChangeset(e, chainSelector, explicitRefs)
	}
	contractAdapter, err := committeeVerifierContractRegistry.GetByChain(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("no committee verifier contract adapter for chain %d: %w", chainSelector, err)
	}
	laneCCVRefs, err := contractAdapter.GetCommitteeVerifierResolver(e.DataStore, chainSelector, defaultQualifier)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve default lane CCV refs for %q qualifier on chain %d: %w", defaultQualifier, chainSelector, err)
	}
	return resolveLocalContractsForTopologyChangeset(e, chainSelector, laneCCVRefs)
}

// validateDefaultCCVsResolvable checks that every remote chain entry that relies on
// auto-resolved CCVs (i.e. does not provide explicit DefaultInboundCCVs or DefaultOutboundCCVs)
// can resolve default lane CCV refs for the "default" qualifier via the chain family's
// CommitteeVerifierContractAdapter (typically the verifier resolver only).
// This fails early in validation rather than silently producing a lane with no verifiers.
func validateDefaultCCVsResolvable(
	chainCfg partialChainConfig,
	registry *adapters.CommitteeVerifierContractRegistry,
	e deployment.Environment,
) error {
	needsAutoResolve := false
	for _, remote := range chainCfg.RemoteChains {
		if len(remote.DefaultInboundCCVs) == 0 || len(remote.DefaultOutboundCCVs) == 0 {
			needsAutoResolve = true
			break
		}
	}
	if !needsAutoResolve {
		return nil
	}
	contractAdapter, err := registry.GetByChain(chainCfg.ChainSelector)
	if err != nil {
		return fmt.Errorf("chain %d has remote chains that need auto-resolved CCVs but no committee verifier contract adapter is registered: %w",
			chainCfg.ChainSelector, err)
	}
	refs, err := contractAdapter.GetCommitteeVerifierResolver(e.DataStore, chainCfg.ChainSelector, defaultQualifier)
	if err != nil {
		return fmt.Errorf("chain %d has remote chains that need auto-resolved CCVs but failed to resolve %q qualifier default lane CCV refs: %w",
			chainCfg.ChainSelector, defaultQualifier, err)
	}
	if len(refs) == 0 {
		return fmt.Errorf("chain %d has remote chains that need auto-resolved CCVs but no default lane CCV refs for %q qualifier",
			chainCfg.ChainSelector, defaultQualifier)
	}
	return nil
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

func mergeCommitteeVerifierRemoteChainConfig(
	defaults adapters.CommitteeVerifierRemoteChainDefaults,
	overrides committeeVerifierRemoteChainInput,
	signatureConfig adapters.CommitteeVerifierSignatureQuorumConfig,
) adapters.CommitteeVerifierRemoteChainConfig {
	return adapters.CommitteeVerifierRemoteChainConfig{
		AllowlistEnabled:          utils.Coalesce(overrides.AllowlistEnabled, defaults.AllowlistEnabled),
		AddedAllowlistedSenders:   overrides.AddedAllowlistedSenders,
		RemovedAllowlistedSenders: overrides.RemovedAllowlistedSenders,
		FeeUSDCents:               utils.Coalesce(overrides.FeeUSDCents, defaults.FeeUSDCents),
		GasForVerification:        utils.Coalesce(overrides.GasForVerification, defaults.GasForVerification),
		PayloadSizeBytes:          utils.Coalesce(overrides.PayloadSizeBytes, defaults.PayloadSizeBytes),
		SignatureConfig:           signatureConfig,
	}
}

func mergeFeeQuoterDestChainConfig(
	defaults adapters.FeeQuoterDestChainConfigOverrides,
	overrides adapters.FeeQuoterDestChainConfigOverrides,
) adapters.FeeQuoterDestChainConfigOverrides {
	out := defaults
	out.OverrideExistingConfig = overrides.OverrideExistingConfig
	if overrides.IsEnabled != nil {
		out.IsEnabled = overrides.IsEnabled
	}
	if overrides.MaxDataBytes != nil {
		out.MaxDataBytes = overrides.MaxDataBytes
	}
	if overrides.MaxPerMsgGasLimit != nil {
		out.MaxPerMsgGasLimit = overrides.MaxPerMsgGasLimit
	}
	if overrides.DestGasPerPayloadByteBase != nil {
		out.DestGasPerPayloadByteBase = overrides.DestGasPerPayloadByteBase
	}
	if overrides.DefaultTokenFeeUSDCents != nil {
		out.DefaultTokenFeeUSDCents = overrides.DefaultTokenFeeUSDCents
	}
	if overrides.DefaultTokenDestGasOverhead != nil {
		out.DefaultTokenDestGasOverhead = overrides.DefaultTokenDestGasOverhead
	}
	if overrides.DefaultTxGasLimit != nil {
		out.DefaultTxGasLimit = overrides.DefaultTxGasLimit
	}
	if overrides.NetworkFeeUSDCents != nil {
		out.NetworkFeeUSDCents = overrides.NetworkFeeUSDCents
	}
	if overrides.LinkFeeMultiplierPercent != nil {
		out.LinkFeeMultiplierPercent = overrides.LinkFeeMultiplierPercent
	}
	if overrides.USDPerUnitGas != nil {
		out.USDPerUnitGas = new(big.Int).Set(overrides.USDPerUnitGas)
	}
	return out
}

// filterNOPsToCommitteeMembers returns only the NOPs whose aliases appear in at least one
// committee ChainConfig referenced by the input chains' CommitteeVerifiers. This prevents
// fetching signing keys for executor-only NOPs that will never be used as committee signers.
func filterNOPsToCommitteeMembers(nopTopology *offchain.NOPTopology, chains []partialChainConfig) []offchain.NOPConfig {
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

func deriveFamiliesFromChains(chains []partialChainConfig) []string {
	selectors := make([]uint64, 0, len(chains))
	for _, c := range chains {
		selectors = append(selectors, c.ChainSelector)
	}
	return deriveFamiliesFromSelectors(selectors)
}
