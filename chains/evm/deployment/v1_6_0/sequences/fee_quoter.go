package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type FeeQuoterApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []fqops.DestChainConfigArgs
}

type FeeQuoterUpdatePricesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fqops.PriceUpdates
}

type FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fqops.ApplyTokenTransferFeeConfigUpdatesArgs
}

var (
	FeeQuoterApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"FeeQuoterApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Apply updates to destination chain configs on the FeeQuoter 1.6.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
		report, err := operations.ExecuteOperation(b, fqops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fqops.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Address,
			Args:          input.UpdatesByChain,
		})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
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

	FeeQuoterUpdatePricesSequence = operations.NewSequence(
		"FeeQuoterUpdatePricesSequence",
		semver.MustParse("1.6.0"),
		"Update token and gas prices on FeeQuoter 1.6.0 contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterUpdatePricesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
		report, err := operations.ExecuteOperation(b, fqops.UpdatePrices, chain, contract.FunctionInput[fqops.PriceUpdates]{
			ChainSelector: chain.Selector,
			Address:       input.Address,
			Args:          input.UpdatesByChain,
		})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterUpdatePricesOp on %s: %w", chain, err)
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

	FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence = operations.NewSequence(
		"FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Update token transfer fee configs on FeeQuoter 1.6.0 contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
		report, err := operations.ExecuteOperation(b, fqops.ApplyTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[fqops.ApplyTokenTransferFeeConfigUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Address,
			Args:          input.UpdatesByChain,
		})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyTokenTransferFeeConfigUpdatesOp on %s: %w", chain, err)
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
)
