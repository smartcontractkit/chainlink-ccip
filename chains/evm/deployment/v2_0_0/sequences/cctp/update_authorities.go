package cctp

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_adapters_v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evm_ops_v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	evm_seq_v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// UpdateAuthorities transfers ownership of the CCTP contracts deployed by the changeset to
// the CLLCCIP MCMS timelock. Contracts that are not ownable, or that are already owned by
// the timelock, are skipped so the sequence is idempotent.
var UpdateAuthorities = cldf_ops.NewSequence(
	"cctp-chain:update-authorities",
	utils.Version_2_0_0,
	"Transfers ownership of CCTP contracts to the CLLCCIP MCMS timelock",
	func(b cldf_ops.Bundle, e *deployment.Environment, input adapters.UpdateAuthoritiesInput) (seq_core.OnChainOutput, error) {
		chain, ok := e.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return seq_core.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		timelockRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:      datastore.ContractType(utils.RBACTimelock),
			Version:   evm_ops_v1_0_0.MCMSVersion,
			Qualifier: utils.CLLQualifier,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to find CLLCCIP timelock ref for chain %d: %w", input.ChainSelector, err)
		}
		timelockAddr, err := evm_datastore_utils.ToEVMAddress(timelockRef)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to parse CLLCCIP timelock address for chain %d: %w", input.ChainSelector, err)
		}

		// Keep only contracts that currently need an ownership transfer: ownable contracts whose
		// owner is not already the timelock.
		refsToTransfer := make([]datastore.AddressRef, 0, len(input.ContractRefs))
		for _, ref := range input.ContractRefs {
			currentOwner, _, loadErr := evm_seq_v1_0_0.LoadOwnableContract(common.HexToAddress(ref.Address), chain.Client)
			if loadErr != nil {
				b.Logger.Debugf("skipping ownership transfer for non-ownable contract %s (%s) on chain %d: %v", ref.Address, ref.Type, input.ChainSelector, loadErr)
				continue
			}
			if currentOwner == timelockAddr {
				continue
			}
			refsToTransfer = append(refsToTransfer, ref)
		}
		if len(refsToTransfer) == 0 {
			return seq_core.OnChainOutput{}, nil
		}

		ownershipAdapter := &evm_adapters_v1_0_0.EVMTransferOwnershipAdapter{}
		if err := ownershipAdapter.InitializeTimelockAddress(*e, mcms.Input{Qualifier: utils.CLLQualifier}); err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to initialize timelock address for chain %d: %w", input.ChainSelector, err)
		}

		ownershipInput := deploy.TransferOwnershipPerChainInput{
			ChainSelector: chain.Selector,
			CurrentOwner:  chain.DeployerKey.From.Hex(),
			ProposedOwner: timelockAddr.Hex(),
			ContractRef:   refsToTransfer,
		}

		var result seq_core.OnChainOutput
		result, err = seq_core.RunAndMergeSequence(b, e.BlockChains, ownershipAdapter.SequenceTransferOwnershipViaMCMS(), ownershipInput, result)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to transfer ownership on chain %d: %w", input.ChainSelector, err)
		}
		result, err = seq_core.RunAndMergeSequence(b, e.BlockChains, ownershipAdapter.SequenceAcceptOwnership(), ownershipInput, result)
		if err != nil {
			return seq_core.OnChainOutput{}, fmt.Errorf("failed to accept ownership on chain %d: %w", input.ChainSelector, err)
		}
		return result, nil
	},
)
