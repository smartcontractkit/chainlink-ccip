package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// RampConfigApplier implements deploy.RampConfigApplier for EVM 1.6.
// It applies imported onRamp/offRamp config (dynamic config and dest/source chain config) to the contracts.
// Static config is set at deployment and is not updated here.
type RampConfigApplier struct{}

// SequenceSetRampConfig returns a sequence that sets onRamp and offRamp dynamic config and dest/source chain config
// from the imported config output.
func (r RampConfigApplier) SequenceSetRampConfig() *cldf_ops.Sequence[deploy.SetRampConfigInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-config-applier:set-ramp-config",
		semver.MustParse("1.6.0"),
		"Sets onRamp and offRamp dynamic config and dest/source chain config from imported config",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.SetRampConfigInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}

			var onRampMeta *seq1_6.OnRampImportConfigSequenceOutput
			var offRampMeta *seq1_6.OffRampImportConfigSequenceOutput
			var onRampAddr, offRampAddr common.Address

			for _, c := range input.ImportOutput {
				if c.ChainSelector != input.ChainSelector {
					continue
				}
				if meta, ok := c.Metadata.(seq1_6.OnRampImportConfigSequenceOutput); ok {
					onRampMeta = &meta
					onRampAddr = common.HexToAddress(c.Address)
				}
				if meta, ok := c.Metadata.(seq1_6.OffRampImportConfigSequenceOutput); ok {
					offRampMeta = &meta
					offRampAddr = common.HexToAddress(c.Address)
				}
			}

			if onRampMeta == nil {
				return sequences.OnChainOutput{}, fmt.Errorf("onRamp import config metadata not found in import output for chain %d", input.ChainSelector)
			}
			if offRampMeta == nil {
				return sequences.OnChainOutput{}, fmt.Errorf("offRamp import config metadata not found in import output for chain %d", input.ChainSelector)
			}

			var writes []contract.WriteOutput

			// Set OnRamp dynamic config
			onRampDynamicReport, err := cldf_ops.ExecuteOperation(b, onrampops.SetDynamicConfig, chain, contract.FunctionInput[onrampops.DynamicConfig]{
				ChainSelector: input.ChainSelector,
				Address:       onRampAddr,
				Args:          onRampMeta.DynamicConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set onRamp dynamic config: %w", err)
			}
			writes = append(writes, onRampDynamicReport.Output)

			// Apply OnRamp dest chain config updates
			onRampDestArgs := make([]onrampops.DestChainConfigArgs, 0, len(onRampMeta.DestChainCfgs))
			for destSel, cfg := range onRampMeta.DestChainCfgs {
				onRampDestArgs = append(onRampDestArgs, onrampops.DestChainConfigArgs{
					DestChainSelector: destSel,
					Router:            cfg.Router,
					AllowlistEnabled:  cfg.AllowlistEnabled,
				})
			}
			if len(onRampDestArgs) > 0 {
				onRampDestReport, err := cldf_ops.ExecuteOperation(b, onrampops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onrampops.DestChainConfigArgs]{
					ChainSelector: input.ChainSelector,
					Address:       onRampAddr,
					Args:          onRampDestArgs,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply onRamp dest chain config: %w", err)
				}
				writes = append(writes, onRampDestReport.Output)
			}

			// Set OffRamp dynamic config
			offRampDynamicReport, err := cldf_ops.ExecuteOperation(b, offrampops.SetDynamicConfig, chain, contract.FunctionInput[offrampops.DynamicConfig]{
				ChainSelector: input.ChainSelector,
				Address:       offRampAddr,
				Args:          offRampMeta.DynamicConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set offRamp dynamic config: %w", err)
			}
			writes = append(writes, offRampDynamicReport.Output)

			// Apply OffRamp source chain config updates
			offRampSourceArgs := make([]offrampops.SourceChainConfigArgs, 0, len(offRampMeta.SourceChainCfgs))
			for srcSel, cfg := range offRampMeta.SourceChainCfgs {
				offRampSourceArgs = append(offRampSourceArgs, offrampops.SourceChainConfigArgs{
					SourceChainSelector:       srcSel,
					Router:                    cfg.Router,
					IsEnabled:                 cfg.IsEnabled,
					IsRMNVerificationDisabled: cfg.IsRMNVerificationDisabled,
					OnRamp:                    cfg.OnRamp,
				})
			}
			if len(offRampSourceArgs) > 0 {
				offRampSourceReport, err := cldf_ops.ExecuteOperation(b, offrampops.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offrampops.SourceChainConfigArgs]{
					ChainSelector: input.ChainSelector,
					Address:       offRampAddr,
					Args:          offRampSourceArgs,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply offRamp source chain config: %w", err)
				}
				writes = append(writes, offRampSourceReport.Output)
			}

			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build batch operation: %w", err)
			}

			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		},
	)
}
