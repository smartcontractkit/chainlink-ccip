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

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	offbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OffRampApplySourceChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []offbind.OffRampSourceChainConfigArgs
}

type OffRampImportConfigSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	RemoteChains  []uint64
}

type OffRampImportConfigSequenceOutput struct {
	SourceChainCfgs map[uint64]offbind.OffRampSourceChainConfig
	StaticConfig    offbind.OffRampStaticConfig
	DynamicConfig   offbind.OffRampDynamicConfig
}

var (
	OffRampApplySourceChainConfigUpdatesSequence = operations.NewSequence(
		"OffRampApplySourceChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to source chain configurations stored on OffRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OffRampApplySourceChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]ops2contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			oor, err := offbind.NewOffRamp(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind off ramp: %w", err)
			}
			report, err := operations.ExecuteOperation(b, offrampops.NewWriteApplySourceChainConfigUpdates(oor), chain, ops2contract.FunctionInput[[]offbind.OffRampSourceChainConfigArgs]{
				Args: input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OffRampApplySourceChainConfigUpdatesOp on %s: %w", chain, err)
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

	OffRampImportConfigSequence = operations.NewSequence(
		"OffRampImportConfigSequence",
		semver.MustParse("1.6.0"),
		"Imports OffRamp contract configuration from multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OffRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			output := OffRampImportConfigSequenceOutput{
				SourceChainCfgs: make(map[uint64]offbind.OffRampSourceChainConfig),
			}
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			oor, err := offbind.NewOffRamp(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind off ramp: %w", err)
			}
			report, err := operations.ExecuteOperation(b, offrampops.NewReadGetStaticConfig(oor), chain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetStaticConfig on %s: %w", chain, err)
			}
			output.StaticConfig = report.Output
			out, err := operations.ExecuteOperation(b, offrampops.NewReadGetDynamicConfig(oor), chain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetDynamicConfig on %s: %w", chain, err)
			}
			output.DynamicConfig = out.Output
			for _, remoteChain := range input.RemoteChains {
				report, err := operations.ExecuteOperation(b, offrampops.NewReadGetSourceChainConfig(oor), chain, ops2contract.FunctionInput[uint64]{
					Args: remoteChain,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetSourceChainConfig for chain %d on %s: %w", remoteChain, chain, err)
				}
				output.SourceChainCfgs[remoteChain] = report.Output
			}
			contractMeta := datastore.ContractMetadata{
				Address:       input.Address.Hex(),
				ChainSelector: chain.Selector,
				Metadata:      output,
			}

			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: []datastore.ContractMetadata{contractMeta},
				},
			}, nil
		})
)
