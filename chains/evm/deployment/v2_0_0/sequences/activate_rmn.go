package sequences

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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

// ActivateRMNInput deploys RMN 2.0.0 with the Ultra Fast Curse MCMS proposer as an initial
// curse admin, transfers RMN ownership to RMNMCMS, and points RMNProxy at the new RMN
// implementation (final step — makes fast curse live for CCIP consumers).
type ActivateRMNInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

// ActivateRMN runs the RMN v2 activation flow on a single EVM chain.
var ActivateRMN = cldf_ops.NewSequence(
	"activate-rmn",
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

		// 1. Deploy RMN 2.0.0 with the Ultra Fast Curse MCMS proposer as an initial curse admin.
		deployReport, err := cldf_ops.ExecuteSequence(b, DeployRMN, chain, DeployRMNInput{
			ChainSelector:     chain.Selector,
			ExistingAddresses: input.ExistingAddresses,
			Args: rmnops.ConstructorArgs{
				CurseAdmins: []common.Address{ultraFastCurseTimeLock},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RMN on chain %d: %w", chain.Selector, err)
		}
		if len(deployReport.Output.Addresses) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("RMN deploy produced no address on chain %d", chain.Selector)
		}
		rmnRef := deployReport.Output.Addresses[0]
		rmnAddr := common.HexToAddress(rmnRef.Address)
		output.Addresses = append(output.Addresses, rmnRef)

		// 2. Transfer RMN ownership to RMNMCMS timelock.
		ownershipBatchOps, err := transferRMNOwnershipToTimelock(b, chain, rmnAddr, rmnTimelock)
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
		Args: rmn_proxy.SetRMNArgs{
			RMN: rmnAddr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set RMN on RMNProxy on chain %d: %w", chain.Selector, err)
	}

	currentRMN, err := cldf_ops.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       rmnProxyAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify RMN on RMNProxy on chain %d: %w", chain.Selector, err)
	}
	if currentRMN.Output != rmnAddr {
		return nil, fmt.Errorf(
			"RMNProxy RMN on chain %d is %s after setRMN, expected %s",
			chain.Selector, currentRMN.Output.Hex(), rmnAddr.Hex(),
		)
	}

	var writes []contract.WriteOutput
	writes = append(writes, setRMNReport.Output)
	return writes, nil
}

func transferRMNOwnershipToTimelock(
	b cldf_ops.Bundle,
	chain evm.Chain,
	rmnAddr, rmnTimelock common.Address,
) ([]mcms_types.BatchOperation, error) {
	owner, ownable, err := mcms_seq.LoadOwnableContract(rmnAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to load RMN ownable contract %s: %w", rmnAddr.Hex(), err)
	}
	if owner == rmnTimelock {
		b.Logger.Infof("RMN %s on chain %d is already owned by RMNMCMS timelock %s",
			rmnAddr.Hex(), chain.Selector, rmnTimelock.Hex())
		return nil, nil
	}

	ownershipInput := mcms_ops.OpTransferOwnershipInput{
		ChainSelector:   chain.Selector,
		Address:         rmnAddr,
		ProposedOwner:   rmnTimelock,
		ContractType:    rmnops.ContractType,
		TimelockAddress: rmnTimelock,
	}
	deps := mcms_ops.OpEVMOwnershipDeps{
		Chain:    chain,
		OwnableC: ownable,
	}

	var pendingWrites []contract.WriteOutput
	transferReport, err := cldf_ops.ExecuteOperation(b, mcms_ops.OpTransferOwnership, deps, ownershipInput)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer RMN ownership to RMNMCMS on chain %d: %w", chain.Selector, err)
	}
	pendingWrites = append(pendingWrites, transferReport.Output)

	acceptReport, err := cldf_ops.ExecuteOperation(b, mcms_ops.OpAcceptOwnership, deps, ownershipInput)
	if err != nil {
		return nil, fmt.Errorf("failed to accept RMN ownership on RMNMCMS timelock on chain %d: %w", chain.Selector, err)
	}
	pendingWrites = append(pendingWrites, acceptReport.Output)

	if len(pendingWrites) == 0 {
		return nil, nil
	}
	batch, err := contract.NewBatchOperationFromWrites(pendingWrites)
	if err != nil {
		return nil, fmt.Errorf("failed to create batch operation from RMN ownership writes: %w", err)
	}
	return []mcms_types.BatchOperation{batch}, nil
}

func resolveTimelockAddress(refs []datastore.AddressRef, chainSelector uint64, qualifier string) (common.Address, error) {
	ref := datastore_utils.GetAddressRef(
		refs,
		chainSelector,
		common_utils.RBACTimelock,
		mcms_ops.MCMSVersion,
		qualifier,
	)
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
