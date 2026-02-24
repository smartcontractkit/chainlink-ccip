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

// evmRampConfigSetArgs is the output of SequenceRampConfigInputCreation and input to SequenceSetRampConfig for EVM 1.6.
// Used as the concrete type when RampConfigSetArgs is any for deploy.RampConfigApplier[any].
type evmRampConfigSetArgs struct {
	ChainSelector        uint64
	OnRampAddr           common.Address
	OffRampAddr          common.Address
	OnRampDynamicConfig  onrampops.DynamicConfig
	OffRampDynamicConfig offrampops.DynamicConfig
	OnRampDestArgs       []onrampops.DestChainConfigArgs
	OffRampSourceArgs    []offrampops.SourceChainConfigArgs
}

// RampConfigApplier uses RampConfigSetArgs any so it implements deploy.RampConfigApplier[any] and can be registered directly.
type RampConfigApplier[RampConfigSetArgs any] struct{}

// SequenceRampConfigInputCreation parses consolidated contract metadata and produces ramp config set args.
func (r RampConfigApplier[RampConfigSetArgs]) SequenceRampConfigInputCreation() *cldf_ops.Sequence[deploy.SetRampConfigInput, RampConfigSetArgs, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-config-applier:input-creation",
		semver.MustParse("1.6.0"),
		"Creates ramp config set args from imported contract metadata",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.SetRampConfigInput) (RampConfigSetArgs, error) {
			var onRampMeta *seq1_6.OnRampImportConfigSequenceOutput
			var offRampMeta *seq1_6.OffRampImportConfigSequenceOutput
			var onRampAddr, offRampAddr common.Address

			for _, c := range input.ContractMeta {
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
				var zero RampConfigSetArgs
				return zero, fmt.Errorf("onRamp import config metadata not found in import output for chain %d", input.ChainSelector)
			}
			if offRampMeta == nil {
				var zero RampConfigSetArgs
				return zero, fmt.Errorf("offRamp import config metadata not found in import output for chain %d", input.ChainSelector)
			}

			onRampDestArgs := make([]onrampops.DestChainConfigArgs, 0, len(onRampMeta.DestChainCfgs))
			for destSel, cfg := range onRampMeta.DestChainCfgs {
				onRampDestArgs = append(onRampDestArgs, onrampops.DestChainConfigArgs{
					DestChainSelector: destSel,
					Router:            cfg.Router,
					AllowlistEnabled:  cfg.AllowlistEnabled,
				})
			}
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

			result := evmRampConfigSetArgs{
				ChainSelector:        input.ChainSelector,
				OnRampAddr:           onRampAddr,
				OffRampAddr:          offRampAddr,
				OnRampDynamicConfig:  onRampMeta.DynamicConfig,
				OffRampDynamicConfig: offRampMeta.DynamicConfig,
				OnRampDestArgs:       onRampDestArgs,
				OffRampSourceArgs:    offRampSourceArgs,
			}
			return any(result).(RampConfigSetArgs), nil
		},
	)
}

// SequenceSetRampConfig applies the ramp config set args to onRamp and offRamp contracts.
func (r RampConfigApplier[RampConfigSetArgs]) SequenceSetRampConfig() *cldf_ops.Sequence[RampConfigSetArgs, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-config-applier:set-ramp-config",
		semver.MustParse("1.6.0"),
		"Sets onRamp and offRamp dynamic config and dest/source chain config from created args",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, args RampConfigSetArgs) (sequences.OnChainOutput, error) {
			evmArgs, ok := any(args).(evmRampConfigSetArgs)
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("ramp config set args has unexpected type")
			}
			chain, ok := chains.EVMChains()[evmArgs.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", evmArgs.ChainSelector)
			}

			var writes []contract.WriteOutput

			// Set OnRamp dynamic config
			onRampDynamicReport, err := cldf_ops.ExecuteOperation(b, onrampops.SetDynamicConfig, chain, contract.FunctionInput[onrampops.DynamicConfig]{
				ChainSelector: evmArgs.ChainSelector,
				Address:       evmArgs.OnRampAddr,
				Args:          evmArgs.OnRampDynamicConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set onRamp dynamic config: %w", err)
			}
			writes = append(writes, onRampDynamicReport.Output)

			if len(evmArgs.OnRampDestArgs) > 0 {
				onRampDestReport, err := cldf_ops.ExecuteOperation(b, onrampops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onrampops.DestChainConfigArgs]{
					ChainSelector: evmArgs.ChainSelector,
					Address:       evmArgs.OnRampAddr,
					Args:          evmArgs.OnRampDestArgs,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply onRamp dest chain config: %w", err)
				}
				writes = append(writes, onRampDestReport.Output)
			}

			// Set OffRamp dynamic config
			offRampDynamicReport, err := cldf_ops.ExecuteOperation(b, offrampops.SetDynamicConfig, chain, contract.FunctionInput[offrampops.DynamicConfig]{
				ChainSelector: evmArgs.ChainSelector,
				Address:       evmArgs.OffRampAddr,
				Args:          evmArgs.OffRampDynamicConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set offRamp dynamic config: %w", err)
			}
			writes = append(writes, offRampDynamicReport.Output)

			if len(evmArgs.OffRampSourceArgs) > 0 {
				offRampSourceReport, err := cldf_ops.ExecuteOperation(b, offrampops.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offrampops.SourceChainConfigArgs]{
					ChainSelector: evmArgs.ChainSelector,
					Address:       evmArgs.OffRampAddr,
					Args:          evmArgs.OffRampSourceArgs,
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
