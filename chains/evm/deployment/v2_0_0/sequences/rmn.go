package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/rmn"
	rmnbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ConfigureRMNCurseAdminsInput holds the parameters for updating authorized callers (curse admins) on an existing RMN contract.
type ConfigureRMNCurseAdminsInput struct {
	ChainSelector uint64
	RMNAddress    common.Address
	Args          rmnbind.AuthorizedCallersAuthorizedCallerArgs
}

// ConfigureRMNCurseAdmins applies authorized caller (curse admin) updates to an already-deployed RMN contract.
var ConfigureRMNCurseAdmins = cldf_ops.NewSequence(
	"configure-rmn-curse-admins",
	rmnops.Version,
	"Applies authorized caller (curse admin) updates to the RMN contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureRMNCurseAdminsInput) (output sequences.OnChainOutput, err error) {
		if input.ChainSelector != chain.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match chain %d",
				input.ChainSelector, chain.Selector)
		}
		if len(input.Args.AddedCallers) == 0 && len(input.Args.RemovedCallers) == 0 {
			return sequences.OnChainOutput{}, nil
		}
		rmn, err := rmnbind.NewRMN(input.RMNAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind RMN: %w", err)
		}
		report, err := cldf_ops.ExecuteOperation(
			b, rmnops.NewWriteApplyAuthorizedCallerUpdates(rmn), chain,
			ops2contract.FunctionInput[rmnbind.AuthorizedCallersAuthorizedCallerArgs]{
				Args: input.Args,
			})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to RMN(%s) on chain %d: %w",
				input.RMNAddress.Hex(), chain.Selector, err)
		}
		batch, err := ops2contract.NewBatchOperationFromWrites([]ops2contract.WriteOutput{report.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		output.BatchOps = []mcms_types.BatchOperation{batch}
		return output, nil
	},
)

// DeployRMNInput deploys RMN (IRMN onchain implementation used behind RMNProxy / by verifiers).
type DeployRMNInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	Args              rmnops.ConstructorArgs
}

// DeployRMN deploys a new RMN contract or returns the existing address from the datastore.
var DeployRMN = cldf_ops.NewSequence(
	"deploy-rmn",
	rmnops.Version,
	"Deploys the RMN (curse / IRMN) contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployRMNInput) (output sequences.OnChainOutput, err error) {
		if input.ChainSelector != chain.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match chain %d",
				input.ChainSelector, chain.Selector)
		}

		rmnRef, err := ops2contract.MaybeDeployContract(b, rmnops.Deploy, chain, ops2contract.DeployInput[rmnops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmnops.ContractType, *rmnops.Version),
			Args:           input.Args,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RMN: %w", err)
		}
		if rmnRef.Address == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("RMN address is empty after deploy on chain %d", chain.Selector)
		}
		output.Addresses = append(output.Addresses, rmnRef)
		return output, nil
	},
)

// SeqCurseInput holds the parameters for cursing one or more subjects on an RMN v2.0.0 contract.
type SeqCurseInput struct {
	// ChainSelector is added to make the input distinct from other chains that have the same RMN address and Subjects
	// Without this distinction the sequence will trigger an cache hit in CLD and might not be executed.
	ChainSelector uint64
	RMNAddress    common.Address
	Subjects      [][16]byte
}

// SeqUncurseInput holds the parameters for uncursing one or more subjects on an RMN v2.0.0 contract.
type SeqUncurseInput struct {
	// ChainSelector is added to make the input distinct from other chains that have the same RMN address and Subjects
	// Without this distinction the sequence will trigger an cache hit in CLD and might not be executed.
	ChainSelector uint64
	RMNAddress    common.Address
	Subjects      [][16]byte
}

// RmnCurse curses one or more subjects on an RMN v2.0.0 contract.
var RmnCurse = cldf_ops.NewSequence(
	"rmn-curse",
	rmnops.Version,
	"Cursing subjects with RMN",
	func(b cldf_ops.Bundle, chain evm.Chain, in SeqCurseInput) (output sequences.OnChainOutput, err error) {
		rmn, err := rmnbind.NewRMN(in.RMNAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind RMN: %w", err)
		}
		opOutput, err := cldf_ops.ExecuteOperation(b, rmnops.NewWriteCurse0(rmn), chain, ops2contract.FunctionInput[[][16]byte]{
			Args: in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse with RMN at %s on chain %d: %w", in.RMNAddress.String(), chain.Selector, err)
		}
		batchOp, err := ops2contract.NewBatchOperationFromWrites([]ops2contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})

// RmnUncurse uncurses one or more subjects on an RMN v2.0.0 contract.
var RmnUncurse = cldf_ops.NewSequence(
	"rmn-uncurse",
	rmnops.Version,
	"Uncursing subjects with RMN",
	func(b cldf_ops.Bundle, chain evm.Chain, in SeqUncurseInput) (output sequences.OnChainOutput, err error) {
		rmn, err := rmnbind.NewRMN(in.RMNAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind RMN: %w", err)
		}
		opOutput, err := cldf_ops.ExecuteOperation(b, rmnops.NewWriteUncurse0(rmn), chain, ops2contract.FunctionInput[[][16]byte]{
			Args: in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to uncurse with RMN at %s on chain %d: %w", in.RMNAddress.String(), chain.Selector, err)
		}
		batchOp, err := ops2contract.NewBatchOperationFromWrites([]ops2contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
