package changesets

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// OffRampSetSourceOnRampsEntry sets the full source-chain onramp whitelist on an EVM OffRamp.
// OffRamp.applySourceChainConfigUpdates replaces onRamps entirely; pass every address that should
// remain allowed (e.g. [new, old] during drain, [new] after drain).
type OffRampSetSourceOnRampsEntry struct {
	LocalChainSelector  uint64   `json:"localChainSelector" yaml:"localChainSelector"`
	SourceChainSelector uint64   `json:"sourceChainSelector" yaml:"sourceChainSelector"`
	OnRamps             []string `json:"onRamps" yaml:"onRamps"`
}

// OffRampSetSourceOnRampsInput is the durable-pipeline payload for offramp_set_source_onramps.
type OffRampSetSourceOnRampsInput struct {
	Updates []OffRampSetSourceOnRampsEntry `json:"updates" yaml:"updates"`
	MCMS    mcms.Input                     `json:"mcms" yaml:"mcms"`
}

// OffRampSetSourceOnRamps updates OffRamp source-chain config on EVM chains.
func OffRampSetSourceOnRamps(mcmsRegistry *cs_core.MCMSReaderRegistry) cldf.ChangeSetV2[OffRampSetSourceOnRampsInput] {
	return cldf.CreateChangeSet(
		makeOffRampSetSourceOnRampsApply(mcmsRegistry),
		makeOffRampSetSourceOnRampsVerify(),
	)
}

func makeOffRampSetSourceOnRampsVerify() func(cldf.Environment, OffRampSetSourceOnRampsInput) error {
	return func(_ cldf.Environment, cfg OffRampSetSourceOnRampsInput) error {
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
}

func makeOffRampSetSourceOnRampsApply(mcmsRegistry *cs_core.MCMSReaderRegistry) func(cldf.Environment, OffRampSetSourceOnRampsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg OffRampSetSourceOnRampsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		adapter := &evm_adapters.ChainFamilyAdapter{}

		for i, update := range cfg.Updates {
			family, err := chainsel.GetSelectorFamily(update.LocalChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: %w", i, err)
			}
			if family != chainsel.FamilyEVM {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: local chain %d must be EVM, got %q", i, update.LocalChainSelector, family)
			}

			chain, ok := e.BlockChains.EVMChains()[update.LocalChainSelector]
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: EVM chain %d not found in environment", i, update.LocalChainSelector)
			}

			offRampBytes, err := adapter.GetOffRampAddress(e.DataStore, update.LocalChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: resolve OffRamp on chain %d: %w", i, update.LocalChainSelector, err)
			}
			offRampAddr := common.BytesToAddress(offRampBytes)

			desiredOnRamps, err := parseOffRampSourceOnRampAddresses(update.OnRamps)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: %w", i, err)
			}

			currentReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       offRampAddr,
				Args:          update.SourceChainSelector,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: get source chain config for %d on OffRamp %s: %w",
					i, update.SourceChainSelector, offRampAddr, err)
			}
			current := currentReport.Output

			if sequences.UnorderedSliceEqual(current.OnRamps, desiredOnRamps, bytes.Equal) {
				e.Logger.Infow("OffRamp source onramp whitelist already matches desired state, skipping",
					"localChain", update.LocalChainSelector,
					"sourceChain", update.SourceChainSelector,
					"offRamp", offRampAddr.Hex(),
					"onRampCount", len(desiredOnRamps),
				)
				continue
			}

			desired := offramp.SourceChainConfigArgs{
				Router:              current.Router,
				SourceChainSelector: update.SourceChainSelector,
				IsEnabled:           current.IsEnabled,
				OnRamps:             desiredOnRamps,
				DefaultCCVs:         current.DefaultCCVs,
				LaneMandatedCCVs:    current.LaneMandatedCCVs,
			}

			offRampReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       offRampAddr,
				Args:          []offramp.SourceChainConfigArgs{desired},
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: apply source chain config on OffRamp %s: %w", i, offRampAddr, err)
			}

			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{offRampReport.Output})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("updates[%d]: build batch operation: %w", i, err)
			}
			batchOps = append(batchOps, batchOp)
		}

		return cs_core.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func parseOffRampSourceOnRampAddresses(addrs []string) ([][]byte, error) {
	out := make([][]byte, 0, len(addrs))
	seen := make(map[string]struct{}, len(addrs))
	for i, addr := range addrs {
		trimmed := strings.TrimSpace(addr)
		raw := common.FromHex(trimmed)
		if len(raw) == 0 {
			return nil, fmt.Errorf("onRamps[%d]: invalid hex address %q", i, addr)
		}
		var encoded []byte
		switch len(raw) {
		case 20:
			encoded = common.LeftPadBytes(raw, 32)
		case 32:
			encoded = raw
		default:
			return nil, fmt.Errorf("onRamps[%d]: address %q must be 20 or 32 bytes, got %d", i, addr, len(raw))
		}
		key := string(encoded)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, encoded)
	}
	if len(out) == 0 {
		return nil, errors.New("no valid onRamp addresses after parsing")
	}
	return out, nil
}
