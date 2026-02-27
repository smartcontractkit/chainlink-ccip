package lanes

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// DisableRemoteChainInput provides the local and remote chain info
// needed to disable a lane on the local chain.
type DisableRemoteChainInput struct {
	LocalChainSelector  uint64
	RemoteChainSelector uint64
	OnRamp              []byte
	OffRamp             []byte
	Router              []byte
	FeeQuoter           []byte
}

// DisableLanePair identifies two chains whose bidirectional lane should be disabled.
type DisableLanePair struct {
	ChainA  uint64
	ChainB  uint64
	Version *semver.Version
}

// DisableLaneConfig is the input for the DisableLane changeset.
type DisableLaneConfig struct {
	Lanes []DisableLanePair
	MCMS  mcms.Input
}

// DisableLane returns a changeset that disables bidirectional CCIP lanes.
func DisableLane(
	laneRegistry *LaneAdapterRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) cldf.ChangeSetV2[DisableLaneConfig] {
	return cldf.CreateChangeSet(
		makeDisableApply(laneRegistry, mcmsRegistry),
		makeDisableVerify(),
	)
}

func makeDisableVerify() func(cldf.Environment, DisableLaneConfig) error {
	return func(_ cldf.Environment, _ DisableLaneConfig) error {
		return nil
	}
}

func makeDisableApply(
	laneRegistry *LaneAdapterRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, DisableLaneConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DisableLaneConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, lane := range cfg.Lanes {
			chainAFamily, err := chain_selectors.GetSelectorFamily(lane.ChainA)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", lane.ChainA, err)
			}
			chainBFamily, err := chain_selectors.GetSelectorFamily(lane.ChainB)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", lane.ChainB, err)
			}

			chainAAdapter, exists := laneRegistry.GetLaneAdapter(chainAFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no LaneAdapter registered for chain family '%s' version %s", chainAFamily, lane.Version)
			}
			chainBAdapter, exists := laneRegistry.GetLaneAdapter(chainBFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no LaneAdapter registered for chain family '%s' version %s", chainBFamily, lane.Version)
			}

			chainADef := &ChainDefinition{Selector: lane.ChainA}
			chainBDef := &ChainDefinition{Selector: lane.ChainB}

			err = populateAddresses(e.DataStore, chainADef, chainAAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching addresses for chain %d: %w", lane.ChainA, err)
			}
			err = populateAddresses(e.DataStore, chainBDef, chainBAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching addresses for chain %d: %w", lane.ChainB, err)
			}

			type disablePair struct {
				local   *ChainDefinition
				remote  *ChainDefinition
				adapter LaneAdapter
			}
			for _, pair := range []disablePair{
				{local: chainADef, remote: chainBDef, adapter: chainAAdapter},
				{local: chainBDef, remote: chainADef, adapter: chainBAdapter},
			} {
				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					pair.adapter.DisableRemoteChain(),
					e.BlockChains,
					DisableRemoteChainInput{
						LocalChainSelector:  pair.local.Selector,
						RemoteChainSelector: pair.remote.Selector,
						OnRamp:              pair.local.OnRamp,
						OffRamp:             pair.local.OffRamp,
						Router:              pair.local.Router,
						FeeQuoter:           pair.local.FeeQuoter,
					},
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to disable remote chain %d on chain %d: %w", pair.remote.Selector, pair.local.Selector, err)
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
				reports = append(reports, report.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
