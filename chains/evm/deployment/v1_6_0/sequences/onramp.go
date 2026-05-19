package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	orbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OnRampApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []orbind.OnRampDestChainConfigArgs
}

type OnRampImportConfigSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	RemoteChains  []uint64
}

type OnRampImportConfigSequenceOutput struct {
	DestChainCfgs map[uint64]orbind.GetDestChainConfig
	StaticConfig  orbind.OnRampStaticConfig
	DynamicConfig orbind.OnRampDynamicConfig
}

var (
	OnRampApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"OnRampApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to destination chain configurations stored on OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]ops2contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			or, err := orbind.NewOnRamp(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind on ramp: %w", err)
			}
			report, err := operations.ExecuteOperation(b, onrampops.NewWriteApplyDestChainConfigUpdates(or), chain, ops2contract.FunctionInput[[]orbind.OnRampDestChainConfigArgs]{
				Args: input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := ops2contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})

	OnRampImportConfigSequence = operations.NewSequence(
		"OnRampImportConfigSequence",
		semver.MustParse("1.6.0"),
		"Imports OnRamp configuration from TokenAdminRegistry and FeeQuoter contracts into OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			or, err := orbind.NewOnRamp(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind on ramp: %w", err)
			}
			onRampDestConfigs := make(map[uint64]orbind.GetDestChainConfig)
			for _, remoteChain := range input.RemoteChains {
				report, err := operations.ExecuteOperation(b, onrampops.NewReadGetDestChainConfig(or), chain, ops2contract.FunctionInput[uint64]{
					Args: remoteChain,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for chain %d from OnRamp at %s on %s: %w", remoteChain, input.Address.String(), chain, err)
				}
				onRampDestConfigs[remoteChain] = report.Output
			}
			report, err := operations.ExecuteOperation(b, onrampops.NewReadGetStaticConfig(or), chain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get static config from OnRamp at %s on %s: %w", input.Address.String(), chain, err)
			}
			staticConfig := report.Output
			out, err := operations.ExecuteOperation(b, onrampops.NewReadGetDynamicConfig(or), chain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config from OnRamp at %s on %s: %w", input.Address.String(), chain, err)
			}
			dynamicConfig := out.Output
			contractMetadata := datastore.ContractMetadata{
				Address:       input.Address.Hex(),
				ChainSelector: chain.Selector,
				Metadata: OnRampImportConfigSequenceOutput{
					StaticConfig:  staticConfig,
					DynamicConfig: dynamicConfig,
					DestChainCfgs: onRampDestConfigs,
				},
			}
			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: []datastore.ContractMetadata{contractMetadata},
				},
			}, nil
		})
)
