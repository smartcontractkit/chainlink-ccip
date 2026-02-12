package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type SeqCurseInput struct {
	api.CurseInput
	Addr common.Address
}

var SeqCurse = cldf_ops.NewSequence(
	"seq-curse:rmn-remote",
	semver.MustParse("1.0.0"),
	"Cursing subjects with RMNRemote",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqCurseInput) (output sequences.OnChainOutput, err error) {
		// Use Curse0 which takes an array of subjects
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.Curse0, chain, contract.FunctionInput[[][16]byte]{
			Address:       in.Addr,
			ChainSelector: chain.Selector,
			Args:          in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects with RMNRemote at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}
		
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		output.BatchOps = append(output.BatchOps, batchOp)
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})

var SeqUncurse = cldf_ops.NewSequence(
	"seq-uncurse:rmn-remote",
	semver.MustParse("1.0.0"),
	"Uncursing subjects with RMNRemote",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqCurseInput) (output sequences.OnChainOutput, err error) {
		// Use Uncurse0 which takes an array of subjects
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.Uncurse0, chain, contract.FunctionInput[[][16]byte]{
			Address:       in.Addr,
			ChainSelector: chain.Selector,
			Args:          in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to uncurse subjects with RMNRemote at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}
		
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		output.BatchOps = append(output.BatchOps, batchOp)
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
