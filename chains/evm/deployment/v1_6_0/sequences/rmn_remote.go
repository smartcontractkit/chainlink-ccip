package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	rmnbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
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
		rmn, err := rmnbind.NewRMNRemote(in.Addr, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind RMNRemote: %w", err)
		}
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.NewWriteCurse0(rmn), chain, ops2contract.FunctionInput[[][16]byte]{
			Args: in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects with RMNRemote at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}

		batchOp, err := ops2contract.NewBatchOperationFromWrites([]ops2contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})

var SeqUncurse = cldf_ops.NewSequence(
	"seq-uncurse:rmn-remote",
	semver.MustParse("1.0.0"),
	"Uncursing subjects with RMNRemote",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqCurseInput) (output sequences.OnChainOutput, err error) {
		rmn, err := rmnbind.NewRMNRemote(in.Addr, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind RMNRemote: %w", err)
		}
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.NewWriteUncurse0(rmn), chain, ops2contract.FunctionInput[[][16]byte]{
			Args: in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to uncurse subjects with RMNRemote at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}

		batchOp, err := ops2contract.NewBatchOperationFromWrites([]ops2contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
