package changesets

import (
	"errors"
	"fmt"
	"strings"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// OffRampSetSourceOnRampsInput is the durable-pipeline payload for offramp_set_source_onramps.
type OffRampSetSourceOnRampsInput struct {
	Updates []adapters.OffRampSetSourceOnRampsEntry `json:"updates" yaml:"updates"`
	MCMS    mcms.Input                              `json:"mcms" yaml:"mcms"`
}

// OffRampSetSourceOnRamps updates OffRamp source-chain onramp whitelists via chain family adapters.
func OffRampSetSourceOnRamps(
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[OffRampSetSourceOnRampsInput] {
	return cldf.CreateChangeSet(
		makeOffRampSetSourceOnRampsApply(chainFamilyRegistry, mcmsRegistry),
		verifyOffRampSetSourceOnRampsInput,
	)
}

func verifyOffRampSetSourceOnRampsInput(_ cldf.Environment, cfg OffRampSetSourceOnRampsInput) error {
	if len(cfg.Updates) == 0 {
		return errors.New("at least one update is required")
	}
	for i, u := range cfg.Updates {
		if u.LocalChainSelector == 0 {
			return fmt.Errorf("updates[%d]: localChainSelector is required", i)
		}
		if u.SourceChainSelector == 0 {
			return fmt.Errorf("updates[%d]: sourceChainSelector is required", i)
		}
		if len(u.OnRamps) == 0 {
			return fmt.Errorf("updates[%d]: at least one onRamp is required", i)
		}
		for j, addr := range u.OnRamps {
			if strings.TrimSpace(addr) == "" {
				return fmt.Errorf("updates[%d].onRamps[%d]: address is required", i, j)
			}
		}
	}
	return nil
}

func makeOffRampSetSourceOnRampsApply(
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
) func(cldf.Environment, OffRampSetSourceOnRampsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg OffRampSetSourceOnRampsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)

		for i, update := range cfg.Updates {
			family, err := chainsel.GetSelectorFamily(update.LocalChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: %w", i, err)
			}

			adapter, ok := chainFamilyRegistry.GetChainFamily(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: no chain family adapter registered for %q", i, family)
			}

			setter, ok := adapter.(adapters.OffRampSourceOnRampSetter)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: chain family %q does not support offramp source onramp updates", i, family)
			}

			batchOp, skipped, err := setter.SetOffRampSourceOnRamps(e, update)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: %w", i, err)
			}
			if skipped {
				continue
			}
			batchOps = append(batchOps, *batchOp)
		}

		return changesetscore.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
