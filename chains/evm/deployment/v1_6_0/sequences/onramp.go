package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
)

type OnRampApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []onramp.OnRampDestChainConfigArgs
}

type OnRampImportConfigSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	RemoteChains  []uint64
}

var (
	OnRampApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"OnRampApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to destination chain configurations stored on OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, onrampops.OnRampApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.OnRampDestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
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

	OnRampImportConfigSequence = operations.NewSequence(
		"OnRampImportConfigSequence",
		semver.MustParse("1.6.0"),
		"Imports OnRamp configuration from TokenAdminRegistry and FeeQuoter contracts into OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			onRampDestConfigs := make(map[uint64]onramp.GetDestChainConfig)
			for _, remoteChain := range input.RemoteChains {
				report, err := operations.ExecuteOperation(b, onrampops.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       input.Address,
					Args:          remoteChain,
				})
				onRampDestConfigs[remoteChain] = report.Output
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for chain %d from OnRamp at %s on %s: %w", remoteChain, input.Address.String(), chain, err)
				}
			}
		})
)
