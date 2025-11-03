package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var timelockAddr map[uint64]common.Address

type EVMTransferOwnershipAdapter struct {}

func (a *EVMTransferOwnershipAdapter) InitializeTimelockAddress(e deployment.Environment, input mcms.Input) error {
	evmChains := e.BlockChains.EVMChains()
	timelockAddr = make(map[uint64]common.Address)
	for sel := range evmChains {
		reader := &EVMMCMSReader{}
		timelockRef, err := reader.GetTimelockRef(e, sel, input)
		if err != nil {
			return fmt.Errorf("failed to get timelock ref for chain %d: %w", sel, err)
		}
		addr, err := datastore_utils.FindAndFormatRef(e.DataStore, timelockRef, sel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return fmt.Errorf("failed to find timelock address for chain %d: %w", sel, err)
		}
		timelockAddr[sel] = addr
	}
	return nil
}

func (a *EVMTransferOwnershipAdapter) SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-seq-transfer-ownership-via-mcms",
		semver.MustParse("1.0.0"),
		"Transfers ownership of contracts via MCMS",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.TransferOwnershipPerChainInput) (output sequences.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			seqInput := seq.SeqTransferMCMOwnershipToTimelockInput{
				ChainSelector: in.ChainSelector,
				Contracts:     make([]ops.OpTransferOwnershipInput, 0),
			}
			for _, contractRef := range in.ContractRef {
				timelockAddr, ok := timelockAddr[in.ChainSelector]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("timelock address not initialized for chain %d", in.ChainSelector)
				}
				seqInput.Contracts = append(seqInput.Contracts, ops.OpTransferOwnershipInput{
					ChainSelector:   in.ChainSelector,
					ProposedOwner:   common.HexToAddress(in.ProposedOwner),
					Address:         common.HexToAddress(contractRef.Address),
					ContractType:    deployment.ContractType(contractRef.Type),
					TimelockAddress: timelockAddr,
				})
			}
			report, err := cldf_ops.ExecuteSequence(b, seq.SeqTransferMCMOwnershipToTimelock, evmChain, seqInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
			return output, nil
		})
}

func (a *EVMTransferOwnershipAdapter) ShouldAcceptOwnershipWithTransferOwnership(e deployment.Environment, in api.TransferOwnershipPerChainInput) (bool, error) {
	chain, ok := e.BlockChains.EVMChains()[in.ChainSelector]
	if !ok {
		return false, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
	}
	// Only accept ownership if the proposed owner is either the timelock or the deployer
	if common.HexToAddress(in.ProposedOwner) != timelockAddr[in.ChainSelector] && common.HexToAddress(in.ProposedOwner) != chain.DeployerKey.From {
		return false, nil
	}
	return true, nil
}

func (a *EVMTransferOwnershipAdapter) SequenceAcceptOwnership() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-seq-accept-ownership",
		semver.MustParse("1.0.0"),
		"Accepts ownership of contracts",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.TransferOwnershipPerChainInput) (output sequences.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			seqInput := seq.SeqTransferMCMOwnershipToTimelockInput{
				ChainSelector: in.ChainSelector,
				Contracts:     make([]ops.OpTransferOwnershipInput, 0),
			}
			for _, contractRef := range in.ContractRef {
				timelockAddr, ok := timelockAddr[in.ChainSelector]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("timelock address not initialized for chain %d", in.ChainSelector)
				}
				seqInput.Contracts = append(seqInput.Contracts, ops.OpTransferOwnershipInput{
					ChainSelector:   in.ChainSelector,
					ProposedOwner:   common.HexToAddress(in.ProposedOwner),
					Address:         common.HexToAddress(contractRef.Address),
					ContractType:    deployment.ContractType(contractRef.Type),
					TimelockAddress: timelockAddr,
				})
			}
			report, err := cldf_ops.ExecuteSequence(b, seq.SeqAcceptMCMOwnershipFromTimelock, evmChain, seqInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership via MCMS on chain %d: %w", in.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
			return output, nil
		})
}
