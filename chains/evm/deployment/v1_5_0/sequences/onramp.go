package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type OnRampSetTokenTransferFeeConfigSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain onramp.SetTokenTransferFeeConfigInput
}

var OnRampSetTokenTransferFeeConfigSequence = operations.NewSequence(
	"onramp:set-token-transfer-fee-config",
	semver.MustParse("1.5.0"),
	"Set token transfer fee config on the OnRamp 1.5.0 contract across multiple EVM chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampSetTokenTransferFeeConfigSequenceInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}
		report, err := operations.ExecuteOperation(b, onramp.OnRampSetTokenTransferFeeConfig, chain, contract.FunctionInput[onramp.SetTokenTransferFeeConfigInput]{
			ChainSelector: chain.Selector,
			Address:       input.Address,
			Args:          input.UpdatesByChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampSetTokenTransferFeeConfigOp on %s: %w", chain, err)
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
