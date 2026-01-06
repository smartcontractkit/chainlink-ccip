package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type SeqUncurseInput struct {
	Addr     common.Address
	Requests []rmn_contract.RMNOwnerUnvoteToCurseRequest
}

type SeqCurseInput struct {
	api.CurseInput
	CurseID [16]byte
	Addr    common.Address
}

var SeqCurse = cldf_ops.NewSequence(
	"seq-curse:rmn",
	semver.MustParse("1.0.0"),
	"Cursing subjects with RMN",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqCurseInput) (output sequences.OnChainOutput, err error) {
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.Curse, chain, contract.FunctionInput[ops.CurseArgs]{
			Address:       in.Addr,
			ChainSelector: chain.Selector,
			Args: ops.CurseArgs{
				Subject: in.Subjects,
				CurseID: in.CurseID,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse with RMNRemote at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		output.BatchOps = append(output.BatchOps, batchOp)
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})

var SeqUncurse = cldf_ops.NewSequence(
	"seq-uncurse:rmn",
	semver.MustParse("1.0.0"),
	"Uncursing subjects with RMN",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqUncurseInput) (output sequences.OnChainOutput, err error) {
		opOutput, err := cldf_ops.ExecuteOperation(b, ops.Uncurse, chain, contract.FunctionInput[ops.UncurseArgs]{
			Address:       in.Addr,
			ChainSelector: chain.Selector,
			Args: ops.UncurseArgs{
				Requests: in.Requests,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse with RMN at %s on chain %d: %w", in.Addr.String(), chain.Selector, err)
		}
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		output.BatchOps = append(output.BatchOps, batchOp)
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
