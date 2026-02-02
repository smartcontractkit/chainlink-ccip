package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"

	mcms_types "github.com/smartcontractkit/mcms/types"

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OffRampApplySourceChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []offramp.OffRampSourceChainConfigArgs
}

type OffRampImportConfigSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	RemoteChains  []uint64
}

type OffRampImportConfigSequenceOutput struct {
	SourceChainCfgs map[uint64]offramp.OffRampSourceChainConfig
	StaticConfig    offramp.OffRampStaticConfig
	DynamicConfig   offramp.OffRampDynamicConfig
}

var (
	OffRampApplySourceChainConfigUpdatesSequence = operations.NewSequence(
		"OffRampApplySourceChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to source chain configurations stored on OffRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OffRampApplySourceChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, offrampops.OffRampApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.OffRampSourceChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OffRampApplySourceChainConfigUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := contract.NewBatchOperationFromWrites(writes)
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
				SourceChainCfgs: make(map[uint64]offramp.OffRampSourceChainConfig),
			}
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, offrampops.GetStaticConfig, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetStaticConfig on %s: %w", chain, err)
			}
			output.StaticConfig = report.Output
			out, err := operations.ExecuteOperation(b, offrampops.GetDynamicConfig, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetDynamicConfig on %s: %w", chain, err)
			}
			output.DynamicConfig = out.Output
			for _, remoteChain := range input.RemoteChains {
				report, err := operations.ExecuteOperation(b, offrampops.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       input.Address,
					Args:          remoteChain,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GetDestChainConfig for chain %d on %s: %w", remoteChain, chain, err)
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
