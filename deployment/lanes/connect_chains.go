package lanes

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ConnectChains returns a changeset that configures CCIP lanes between chains using the provided lane and MCMS registries.
func ConnectChains(laneRegistry *LaneAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConnectChainsConfig] {
	return cldf.CreateChangeSet(makeApply(laneRegistry, mcmsRegistry), makeVerify(laneRegistry, mcmsRegistry))
}

func makeVerify(_ *LaneAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ConnectChainsConfig) error {
	return func(_ cldf.Environment, cfg ConnectChainsConfig) error {
		for i, lane := range cfg.Lanes {
			if err := validateChainDefinition(lane.ChainA); err != nil {
				return fmt.Errorf("lane %d ChainA: %w", i, err)
			}
			if err := validateChainDefinition(lane.ChainB); err != nil {
				return fmt.Errorf("lane %d ChainB: %w", i, err)
			}
		}
		return nil
	}
}

// validateChainDefinition rejects input where the caller has set fields that
// are populated programmatically by the changeset. Setting these is always a
// mistake — the values would be silently overwritten by populateAddresses.
func validateChainDefinition(def ChainDefinition) error {
	type check struct {
		name string
		set  bool
	}
	for _, c := range []check{
		{"OnRamp", len(def.OnRamp) > 0},
		{"OffRamp", len(def.OffRamp) > 0},
		{"Router", len(def.Router) > 0},
		{"FeeQuoter", len(def.FeeQuoter) > 0},
		{"FeeQuoterDestChainConfig", def.FeeQuoterDestChainConfig != (FeeQuoterDestChainConfig{})},
		{"FeeQuoterVersion", def.FeeQuoterVersion != nil},
	} {
		if c.set {
			return fmt.Errorf("field %q must not be set by the caller (populated programmatically)", c.name)
		}
	}
	return nil
}

func makeApply(laneRegistry *LaneAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		for _, lane := range cfg.Lanes {
			chainA, chainB := &lane.ChainA, &lane.ChainB
			chainAFamily, err := chain_selectors.GetSelectorFamily(chainA.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			chainAAdapter, exists := laneRegistry.GetLaneAdapter(chainAFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", chainAFamily)
			}
			chainBFamily, err := chain_selectors.GetSelectorFamily(chainB.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			chainBAdapter, exists := laneRegistry.GetLaneAdapter(chainBFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", chainBFamily)
			}
			err = populateAddresses(e.DataStore, chainA, chainAAdapter, lane.Version, lane.TestRouter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for src chain %d: %w", chainA.Selector, err)
			}
			err = populateAddresses(e.DataStore, chainB, chainBAdapter, lane.Version, lane.TestRouter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for dest chain %d: %w", chainB.Selector, err)
			}
			// to allow for type serialization, we apply fee quoter dest chain config overrides
			// at the lane level and then set them to nil before passing to the adapters
			chainA.FeeQuoterDestChainConfigOverrides = nil
			chainB.FeeQuoterDestChainConfigOverrides = nil
			type lanePair struct {
				src         *ChainDefinition
				dest        *ChainDefinition
				srcAdapter  LaneAdapter
				destAdapter LaneAdapter
			}
			for _, pair := range []lanePair{
				{src: chainA, dest: chainB, srcAdapter: chainAAdapter, destAdapter: chainBAdapter},
				{src: chainB, dest: chainA, srcAdapter: chainBAdapter, destAdapter: chainAAdapter},
			} {
				configureLaneReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, pair.srcAdapter.ConfigureLaneLegAsSource(), e.BlockChains, UpdateLanesInput{
					Source:       pair.src,
					Dest:         pair.dest,
					IsDisabled:   lane.IsDisabled,
					TestRouter:   lane.TestRouter,
					ExtraConfigs: lane.ExtraConfigs,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to lane leg as source with selector %d: %w", pair.src.Selector, err)
				}
				batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
				reports = append(reports, configureLaneReport.ExecutionReports...)
				for _, r := range configureLaneReport.Output.Addresses {
					if err := ds.Addresses().Add(r); err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
					}
				}
				// Write metadata to datastore
				err = sequences.WriteMetadataToDatastore(ds, configureLaneReport.Output.Metadata)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to write metadata to datastore: %w", err)
				}

				configureLaneReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, pair.destAdapter.ConfigureLaneLegAsDest(), e.BlockChains, UpdateLanesInput{
					Source:       pair.src,
					Dest:         pair.dest,
					IsDisabled:   lane.IsDisabled,
					TestRouter:   lane.TestRouter,
					ExtraConfigs: lane.ExtraConfigs,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure lane leg as dest with selector %d: %w", pair.dest.Selector, err)
				}
				batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
				reports = append(reports, configureLaneReport.ExecutionReports...)
				for _, r := range configureLaneReport.Output.Addresses {
					if err := ds.Addresses().Add(r); err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
					}
				}
				// Write metadata to datastore
				err = sequences.WriteMetadataToDatastore(ds, configureLaneReport.Output.Metadata)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to write metadata to datastore: %w", err)
				}
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func populateAddresses(ds datastore.DataStore, chainDef *ChainDefinition, adapter LaneAdapter, version *semver.Version, isTestRouter bool) error {
	var err error
	chainDef.OnRamp, err = adapter.GetOnRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching onramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.OffRamp, err = adapter.GetOffRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching offramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.FeeQuoter, err = adapter.GetFQAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching fee quoter address for chain %d: %w", chainDef.Selector, err)
	}
	if vp, ok := adapter.(FeeQuoterVersionProvider); ok {
		chainDef.FeeQuoterVersion, err = vp.GetFQVersion(ds, chainDef.FeeQuoter, chainDef.Selector)
		if err != nil {
			return fmt.Errorf("error fetching fee quoter version for chain %d: %w", chainDef.Selector, err)
		}
	}
	chainDef.Router, err = adapter.GetRouterAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching router address for chain %d: %w", chainDef.Selector, err)
	}
	if isTestRouter {
		if tp, ok := adapter.(TestRouterProvider); ok {
			chainDef.Router, err = tp.GetTestRouter(ds, chainDef.Selector)
			if err != nil {
				return fmt.Errorf("error fetching test router address for chain %d: %w", chainDef.Selector, err)
			}
		}
	}
	chainDef.FeeQuoterDestChainConfig = adapter.GetFeeQuoterDestChainConfig()
	if chainDef.FeeQuoterDestChainConfigOverrides != nil {
		(*chainDef.FeeQuoterDestChainConfigOverrides)(&chainDef.FeeQuoterDestChainConfig)
	}
	// TODO: should we also not populate gas price default as it will be used on updates? (see below)
	if chainDef.GasPrice == nil {
		chainDef.GasPrice = adapter.GetDefaultGasPrice()
	}
	// TODO: as this changeset is also used for updates, we should only populate token prices if they are not already set
	// (to avoid overwriting any on-chain live changes). This would need to only happen on the first run of the changeset
	// for a given lane, as currently the underlying adapter implementations always update with whats provided to them.
	//
	// if chainDef.TokenPrices == nil {
	// 	populateTokenPrices(ds, chainDef, adapter)
	// }

	// handle v2 separately
	return populateAddressesV2(ds, chainDef, adapter, version)
}

// This function is V2 fields only
func populateAddressesV2(ds datastore.DataStore, chainDef *ChainDefinition, adapter LaneAdapter, version *semver.Version) error {
	if version.LessThan(common_utils.Version_2_0_0) {
		return nil
	}
	committeeVerifiers := make([]CommitteeVerifierConfig[datastore.AddressRef], len(chainDef.CommitteeVerifiers))
	for i, verifier := range chainDef.CommitteeVerifiers {
		contracts := make([]datastore.AddressRef, 0, len(verifier.CommitteeVerifier))
		for _, contract := range verifier.CommitteeVerifier {
			contract, err := datastore_utils.FindAndFormatRef(ds, contract, contract.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return fmt.Errorf("failed to resolve CommitteeVerifier contract ref on chain with selector %d: %w", chainDef.Selector, err)
			}
			contracts = append(contracts, contract)
		}
		committeeVerifiers[i] = CommitteeVerifierConfig[datastore.AddressRef]{
			CommitteeVerifier: contracts,
			RemoteChains:      verifier.RemoteChains,
		}
	}
	chainDef.CommitteeVerifiers = committeeVerifiers

	executor, err := datastore_utils.FindAndFormatRef(ds, chainDef.DefaultExecutor, chainDef.DefaultExecutor.ChainSelector, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to resolve executor ref on chain with selector %d: %w", chainDef.Selector, err)
	}
	chainDef.DefaultExecutor = executor

	laneMandatedInboundCCVs := make([]datastore.AddressRef, 0, len(chainDef.LaneMandatedInboundCCVs))
	for _, ccv := range chainDef.LaneMandatedInboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(ds, ccv, ccv.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainDef.Selector, err)
		}
		laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, resolvedCCV)
	}
	chainDef.LaneMandatedInboundCCVs = laneMandatedInboundCCVs

	laneMandatedOutboundCCVs := make([]datastore.AddressRef, 0, len(chainDef.LaneMandatedOutboundCCVs))
	for _, ccv := range chainDef.LaneMandatedOutboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(ds, ccv, ccv.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainDef.Selector, err)
		}
		laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, resolvedCCV)
	}
	chainDef.LaneMandatedOutboundCCVs = laneMandatedOutboundCCVs

	defaultInboundCCVs := make([]datastore.AddressRef, 0, len(chainDef.DefaultInboundCCVs))
	for _, ccv := range chainDef.DefaultInboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(ds, ccv, ccv.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainDef.Selector, err)
		}
		defaultInboundCCVs = append(defaultInboundCCVs, resolvedCCV)
	}
	chainDef.DefaultInboundCCVs = defaultInboundCCVs

	defaultOutboundCCVs := make([]datastore.AddressRef, 0, len(chainDef.DefaultOutboundCCVs))
	for _, ccv := range chainDef.DefaultOutboundCCVs {
		resolvedCCV, err := datastore_utils.FindAndFormatRef(ds, ccv, ccv.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to resolve ccv ref on chain with selector %d: %w", chainDef.Selector, err)
		}
		defaultOutboundCCVs = append(defaultOutboundCCVs, resolvedCCV)
	}
	chainDef.DefaultOutboundCCVs = defaultOutboundCCVs
	return nil
}

func populateTokenPrices(ds datastore.DataStore, chainDef *ChainDefinition, adapter LaneAdapter) {
	var tokenPrices map[datastore.ContractType]*big.Int
	// Check if adapter implements TokenPriceProvider (optional interface)
	priceProvider, ok := adapter.(TokenPriceProvider)
	if !ok {
		return
	}

	// Get prices keyed by contract type
	tokenPrices = priceProvider.GetDefaultTokenPrices()

	// Resolve contract types to addresses
	addressPrices := make(map[string]*big.Int)
	for contractType, price := range tokenPrices {
		refs := ds.Addresses().Filter(
			datastore.AddressRefByType(contractType),
			datastore.AddressRefByChainSelector(chainDef.Selector),
		)
		for _, ref := range refs {
			addressPrices[ref.Address] = price
		}
	}

	chainDef.TokenPrices = addressPrices
}
