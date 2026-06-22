package sequences

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/rmn"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ConfigureRMNCurseAdminsInput holds the parameters for updating authorized callers (curse admins) on an existing RMN contract.
type ConfigureRMNCurseAdminsInput struct {
	ChainSelector uint64
	RMNAddress    common.Address
	Args          rmnops.AuthorizedCallerArgs
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
		report, err := cldf_ops.ExecuteOperation(
			b, rmnops.ApplyAuthorizedCallerUpdates, chain,
			contract.FunctionInput[rmnops.AuthorizedCallerArgs]{
				ChainSelector: chain.Selector,
				Address:       input.RMNAddress,
				Args:          input.Args,
			})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to RMN(%s) on chain %d: %w",
				input.RMNAddress.Hex(), chain.Selector, err)
		}
		batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		output.BatchOps = []mcms_types.BatchOperation{batch}
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
		opOutput, err := cldf_ops.ExecuteOperation(b, rmnops.Curse0, chain, contract.FunctionInput[[][16]byte]{
			Address:       in.RMNAddress,
			ChainSelector: chain.Selector,
			Args:          in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to curse with RMN at %s on chain %d: %w", in.RMNAddress.String(), chain.Selector, err)
		}
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
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
		opOutput, err := cldf_ops.ExecuteOperation(b, rmnops.Uncurse0, chain, contract.FunctionInput[[][16]byte]{
			Address:       in.RMNAddress,
			ChainSelector: chain.Selector,
			Args:          in.Subjects,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to uncurse with RMN at %s on chain %d: %w", in.RMNAddress.String(), chain.Selector, err)
		}
		batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{opOutput.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})

// ActivateRMNInput deploys RMN 2.0.0 with the Ultra Fast Curse RBACTimelock as an initial
// curse admin, transfers RMN ownership to RMNMCMS, and points RMNProxy at the new RMN
// implementation (final step — makes fast curse live for CCIP consumers).
type ActivateRMNInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

// DeployAndActivateRMN runs the RMN v2 activation flow on a single EVM chain.
var DeployAndActivateRMN = cldf_ops.NewSequence(
	"deploy-and-activate-rmn",
	rmnops.Version,
	"Deploy RMN 2.0.0 with Ultra Fast Curse MCMS as curse admin, transfer ownership to RMNMCMS, and point RMNProxy at the new RMN",
	func(b cldf_ops.Bundle, chain evm.Chain, input ActivateRMNInput) (output sequences.OnChainOutput, err error) {
		if input.ChainSelector != chain.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match chain %d",
				input.ChainSelector, chain.Selector)
		}
		var writes []contract.WriteOutput

		rmnTimelock, err := resolveTimelockAddress(input.ExistingAddresses, chain.Selector, common_utils.RMNTimelockQualifier)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		ultraFastCurseTimeLock, err := resolveTimelockAddress(
			input.ExistingAddresses, chain.Selector, common_utils.UltraFastCurseMCMSQualifier)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// 1. Deploy RMN 2.0.0 with the Ultra Fast Curse RBACTimelock as an initial curse admin.
		rmnRef, err := contract.MaybeDeployContract(b, rmnops.Deploy, chain, contract.DeployInput[rmnops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmnops.ContractType, *rmnops.Version),
			ChainSelector:  chain.Selector,
			Args: rmnops.ConstructorArgs{
				CurseAdmins: []common.Address{ultraFastCurseTimeLock},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RMN on chain %d: %w", chain.Selector, err)
		}
		if rmnRef.Address == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("RMN address is empty after deploy on chain %d", chain.Selector)
		}
		rmnAddr := common.HexToAddress(rmnRef.Address)
		output.Addresses = append(output.Addresses, rmnRef)

		// 2. Transfer RMN ownership to RMNMCMS timelock.
		ownershipBatchOps, err := mcms_seq.TransferAndAcceptOwnership(b, chain, []mcms_ops.OpTransferOwnershipInput{
			{
				ChainSelector:   chain.Selector,
				Address:         rmnAddr,
				ProposedOwner:   rmnTimelock,
				ContractType:    rmnops.ContractType,
				TimelockAddress: rmnTimelock,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		output.BatchOps = append(output.BatchOps, ownershipBatchOps...)

		// 3. Point RMNProxy at the new RMN (makes the implementation live for CCIP).
		rmnProxyWrites, err := pointRMNProxyAtRMN(b, chain, input.ExistingAddresses, rmnAddr)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, rmnProxyWrites...)

		if len(writes) > 0 {
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			output.BatchOps = append(output.BatchOps, batch)
		}

		return output, nil
	},
)

func pointRMNProxyAtRMN(
	b cldf_ops.Bundle,
	chain evm.Chain,
	existingAddresses []datastore.AddressRef,
	rmnAddr common.Address,
) ([]contract.WriteOutput, error) {
	rmnProxyAddr, err := resolveRMNProxyAddress(existingAddresses, chain.Selector)
	if err != nil {
		return nil, err
	}
	setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, contract.FunctionInput[rmn_proxy.SetRMNArgs]{
		ChainSelector: chain.Selector,
		Address:       rmnProxyAddr,
		Args:          rmn_proxy.SetRMNArgs{RMN: rmnAddr},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set RMN on RMNProxy on chain %d: %w", chain.Selector, err)
	}
	return []contract.WriteOutput{setRMNReport.Output}, nil
}

func resolveTimelockAddress(refs []datastore.AddressRef, chainSelector uint64, qualifier string) (common.Address, error) {
	ref := datastore_utils.GetAddressRef(refs, chainSelector, common_utils.RBACTimelock, mcms_ops.MCMSVersion, qualifier)
	if ref.Address == "" {
		return common.Address{}, fmt.Errorf(
			"RBACTimelock with qualifier %q not found in existing addresses for chain %d",
			qualifier, chainSelector,
		)
	}
	addr, err := evm_datastore_utils.ToEVMAddress(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid RBACTimelock address for qualifier %q: %w", qualifier, err)
	}
	return addr, nil
}

func resolveRMNProxyAddress(refs []datastore.AddressRef, chainSelector uint64) (common.Address, error) {
	ds := datastore.NewMemoryDataStore()
	for _, ref := range refs {
		if addErr := ds.Addresses().Add(ref); addErr != nil && !errors.Is(addErr, datastore.ErrAddressRefExists) {
			return common.Address{}, fmt.Errorf("failed to add address ref to datastore: %w", addErr)
		}
	}
	addr, err := datastore_utils.FindAndFormatRef(
		ds.Seal(),
		datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		},
		chainSelector,
		evm_datastore_utils.ToEVMAddress,
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("RMNProxy not found for chain %d: %w", chainSelector, err)
	}
	return addr, nil
}
