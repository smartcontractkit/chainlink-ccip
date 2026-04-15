package changesets

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

const defaultQualifier = "default"

// CommitteeVerifierRemoteChainConfig is the user-facing input for a single remote chain on a
// CommitteeVerifier. It intentionally omits SignatureConfig because signers and threshold are
// derived automatically from the topology (NOPs × committee × chain) — the caller should never
// have to specify them manually.
type CommitteeVerifierRemoteChainConfig struct {
	AllowlistEnabled          *bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               *uint16
	GasForVerification        *uint32
	PayloadSizeBytes          *uint16
}

type CommitteeVerifierInputConfig struct {
	CommitteeQualifier    string
	RemoteChains          map[uint64]CommitteeVerifierRemoteChainConfig
	AllowedFinalityConfig *finality.Config `json:"allowedFinalityConfig" yaml:"allowedFinalityConfig"`
}

// FeeQuoterDestChainConfigOverrides provides lane-pair-specific overrides on top of the
// chain-family defaults returned by the remote adapter's GetDefaultFeeQuoterDestChainConfig.
// Nil fields are left at the adapter default; non-nil fields replace the default.
//
// DestGasOverhead is intentionally omitted — it is a LEGACY v2 field whose responsibility
// moved to OnRamp.BaseExecutionGasCost. ChainFamilySelector is also omitted because it is
// always auto-populated from the remote adapter.
type FeeQuoterDestChainConfigOverrides struct {
	OverrideExistingConfig      bool
	IsEnabled                   *bool
	MaxDataBytes                *uint32
	MaxPerMsgGasLimit           *uint32
	DestGasPerPayloadByteBase   *uint8
	DefaultTokenFeeUSDCents     *uint16
	DefaultTokenDestGasOverhead *uint32
	DefaultTxGasLimit           *uint32
	NetworkFeeUSDCents          *uint16
	LinkFeeMultiplierPercent    *uint8
	USDPerUnitGas               *big.Int
}

// PartialRemoteChainConfig is the user-facing input for a single remote chain. All fields
// are optional, but their empty values are resolved from different sources:
//   - adapter-backed remote-chain settings fall back to the remote chain family adapter's
//     defaults (via GetDefaultRemoteChainConfig and GetDefaultFeeQuoterDestChainConfig);
//   - contract addresses (OnRamp, OffRamp, Executor) are resolved automatically from the
//     datastore; and
//   - empty DefaultInboundCCVs / DefaultOutboundCCVs do not use adapter defaults and instead
//     auto-resolve via the committee verifier contract registry.
//
// An empty DefaultInboundCCVs or DefaultOutboundCCVs slice means "use the auto-resolved CCV
// set for this lane". LaneMandatedInboundCCVs / LaneMandatedOutboundCCVs are additional
// caller-specified CCVs that are enforced on top of that default resolution.
//
// Minimal usage: `remoteChainSelector: {}` — adapter-backed fields use their defaults,
// contracts are datastore-resolved, and CCVs are auto-resolved.
type PartialRemoteChainConfig struct {
	AllowTrafficFrom          *bool
	DefaultExecutorQualifier  string // defaults to "default" if empty
	DefaultInboundCCVs        []datastore.AddressRef
	LaneMandatedInboundCCVs   []datastore.AddressRef
	DefaultOutboundCCVs       []datastore.AddressRef
	LaneMandatedOutboundCCVs  []datastore.AddressRef
	FeeQuoterDestChainConfig  FeeQuoterDestChainConfigOverrides
	ExecutorDestChainConfig   *adapters.ExecutorDestChainConfig // nil = use adapter default
	BaseExecutionGasCost      *uint32                           // nil = use adapter default
	TokenReceiverAllowed      *bool
	MessageNetworkFeeUSDCents *uint16 // nil = use adapter default
	TokenNetworkFeeUSDCents   *uint16 // nil = use adapter default
}

// PartialChainConfig describes the desired state for a single local chain. Well-known contract
// addresses (Router, OnRamp, FeeQuoter, OffRamp) are resolved automatically from the
// datastore via the chain family adapter. Each chain can reference a different set of remote
// chains, so per-source/dest configuration (e.g. different fee quoter parameters per
// destination) is supported naturally.
type PartialChainConfig struct {
	ChainSelector      uint64
	CommitteeVerifiers []CommitteeVerifierInputConfig
	RemoteChains       map[uint64]PartialRemoteChainConfig
	// FamilyExtras holds chain-family-specific configuration that the generic
	// changeset passes through opaquely to the family adapter's sequence.
	// All values must be serializable.
	FamilyExtras map[string]any
}

type ConfigureChainsForLanesFromTopologyConfig struct {
	Topology      *offchain.EnvironmentTopology
	Chains        []PartialChainConfig
	MCMS          mcms.Input
	UseTestRouter bool
}

// enrichedChainConfig is the internal representation of a chain config after the
// topology-level enrichment (signing keys, contract resolution). It mirrors PartialChainConfig
// but carries fully populated CommitteeVerifierConfig (with signers/threshold filled in).
type enrichedChainConfig struct {
	ChainSelector      uint64
	CommitteeVerifiers []adapters.CommitteeVerifierConfig[datastore.AddressRef]
	RemoteChains       map[uint64]PartialRemoteChainConfig
	FamilyExtras       map[string]any
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

			chains = append(chains, enrichedChainConfig{
				ChainSelector:      chainCfg.ChainSelector,
				CommitteeVerifiers: committeeVerifiers,
				RemoteChains:       chainCfg.RemoteChains,
				FamilyExtras:       chainCfg.FamilyExtras,
			})
		}

		return applyConfigureChains(e, chainFamilyRegistry, mcmsRegistry, committeeVerifierContractRegistry, chains, cfg.MCMS, cfg.UseTestRouter)
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
	if executorQualifier == "" {
		executorQualifier = defaultQualifier
	}
	executorAddr, err := localAdapter.ResolveExecutor(e.DataStore, localChainSelector, executorQualifier)
	if err != nil {
		return adapters.RemoteChainConfig[[]byte, string]{}, fmt.Errorf("failed to resolve executor (qualifier %q) on chain %d: %w", executorQualifier, localChainSelector, err)
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
		remoteAdapter.GetDefaultFeeQuoterDestChainConfig(),
		inCfg.FeeQuoterDestChainConfig,
	)
	fqConfig.ChainFamilySelector = remoteAdapter.GetChainFamilySelector()

	defaults := remoteAdapter.GetDefaultRemoteChainConfig()

	allowTrafficFrom := defaults.AllowTrafficFrom
	if inCfg.AllowTrafficFrom != nil {
		allowTrafficFrom = *inCfg.AllowTrafficFrom
	}

	executorConfig := defaults.ExecutorDestChainConfig
	if inCfg.ExecutorDestChainConfig != nil {
		executorConfig = *inCfg.ExecutorDestChainConfig
	}

	baseExecutionGasCost := defaults.BaseExecutionGasCost
	if inCfg.BaseExecutionGasCost != nil {
		baseExecutionGasCost = *inCfg.BaseExecutionGasCost
	}

	tokenReceiverAllowed := defaults.TokenReceiverAllowed
	if inCfg.TokenReceiverAllowed != nil {
		tokenReceiverAllowed = *inCfg.TokenReceiverAllowed
	}

	messageNetworkFeeUSDCents := defaults.MessageNetworkFeeUSDCents
	if inCfg.MessageNetworkFeeUSDCents != nil {
		messageNetworkFeeUSDCents = *inCfg.MessageNetworkFeeUSDCents
	}

	tokenNetworkFeeUSDCents := defaults.TokenNetworkFeeUSDCents
	if inCfg.TokenNetworkFeeUSDCents != nil {
		tokenNetworkFeeUSDCents = *inCfg.TokenNetworkFeeUSDCents
	}

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
		ExecutorDestChainConfig:   executorConfig,
		AddressBytesLength:        remoteAdapter.GetAddressBytesLength(),
		BaseExecutionGasCost:      baseExecutionGasCost,
		TokenReceiverAllowed:      &tokenReceiverAllowed,
		MessageNetworkFeeUSDCents: messageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   tokenNetworkFeeUSDCents,
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
	chainCfg PartialChainConfig,
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
	overrides CommitteeVerifierRemoteChainConfig,
	signatureConfig adapters.CommitteeVerifierSignatureQuorumConfig,
) adapters.CommitteeVerifierRemoteChainConfig {
	allowlistEnabled := defaults.AllowlistEnabled
	if overrides.AllowlistEnabled != nil {
		allowlistEnabled = *overrides.AllowlistEnabled
	}
	feeUSDCents := defaults.FeeUSDCents
	if overrides.FeeUSDCents != nil {
		feeUSDCents = *overrides.FeeUSDCents
	}
	gasForVerification := defaults.GasForVerification
	if overrides.GasForVerification != nil {
		gasForVerification = *overrides.GasForVerification
	}
	payloadSizeBytes := defaults.PayloadSizeBytes
	if overrides.PayloadSizeBytes != nil {
		payloadSizeBytes = *overrides.PayloadSizeBytes
	}
	return adapters.CommitteeVerifierRemoteChainConfig{
		AllowlistEnabled:          allowlistEnabled,
		AddedAllowlistedSenders:   overrides.AddedAllowlistedSenders,
		RemovedAllowlistedSenders: overrides.RemovedAllowlistedSenders,
		FeeUSDCents:               feeUSDCents,
		GasForVerification:        gasForVerification,
		PayloadSizeBytes:          payloadSizeBytes,
		SignatureConfig:           signatureConfig,
	}
}

func mergeFeeQuoterDestChainConfig(
	defaults adapters.FeeQuoterDestChainConfig,
	overrides FeeQuoterDestChainConfigOverrides,
) adapters.FeeQuoterDestChainConfig {
	defaults.OverrideExistingConfig = overrides.OverrideExistingConfig
	if overrides.IsEnabled != nil {
		defaults.IsEnabled = *overrides.IsEnabled
	}
	if overrides.MaxDataBytes != nil {
		defaults.MaxDataBytes = *overrides.MaxDataBytes
	}
	if overrides.MaxPerMsgGasLimit != nil {
		defaults.MaxPerMsgGasLimit = *overrides.MaxPerMsgGasLimit
	}
	if overrides.DestGasPerPayloadByteBase != nil {
		defaults.DestGasPerPayloadByteBase = *overrides.DestGasPerPayloadByteBase
	}
	if overrides.DefaultTokenFeeUSDCents != nil {
		defaults.DefaultTokenFeeUSDCents = *overrides.DefaultTokenFeeUSDCents
	}
	if overrides.DefaultTokenDestGasOverhead != nil {
		defaults.DefaultTokenDestGasOverhead = *overrides.DefaultTokenDestGasOverhead
	}
	if overrides.DefaultTxGasLimit != nil {
		defaults.DefaultTxGasLimit = *overrides.DefaultTxGasLimit
	}
	if overrides.NetworkFeeUSDCents != nil {
		defaults.NetworkFeeUSDCents = *overrides.NetworkFeeUSDCents
	}
	if overrides.LinkFeeMultiplierPercent != nil {
		defaults.LinkFeeMultiplierPercent = *overrides.LinkFeeMultiplierPercent
	}
	if overrides.USDPerUnitGas != nil {
		defaults.USDPerUnitGas = new(big.Int).Set(overrides.USDPerUnitGas)
	}
	return defaults
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

func deriveFamiliesFromChains(chains []PartialChainConfig) []string {
	selectors := make([]uint64, 0, len(chains))
	for _, c := range chains {
		selectors = append(selectors, c.ChainSelector)
	}
	return deriveFamiliesFromSelectors(selectors)
}
