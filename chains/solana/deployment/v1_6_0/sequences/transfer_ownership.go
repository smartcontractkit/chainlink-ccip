package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	feequoterops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) InitializeTimelockAddress(e deployment.Environment, input mcms.Input) error {
	solanaChains := e.BlockChains.SolanaChains()
	a.timelockAddr = make(map[uint64]solana.PublicKey)
	for sel := range solanaChains {
		signer := utils.GetTimelockSignerPDA(e.DataStore.Addresses().Filter(), input.TimelockAddressRef.Qualifier)
		a.timelockAddr[sel] = signer
	}
	return nil
}

func (a *SolanaAdapter) SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[deployops.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"seq-transfer-ownership-via-mcms",
		semver.MustParse("1.0.0"),
		"Transfers ownership of contracts via MCMS",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in deployops.TransferOwnershipPerChainInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.SolanaChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}

			for _, contractRef := range in.ContractRef {
				switch contractRef.Type.String() {
				case routerops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, routerops.TransferOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case offrampops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, offrampops.TransferOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case feequoterops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, feequoterops.TransferOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case rmnremoteops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, rmnremoteops.TransferOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("unsupported contract type %s for ownership transfer via MCMS on Solana", contractRef.Type)
				}
			}
			return output, nil
		})
}

func (a *SolanaAdapter) ShouldAcceptOwnershipWithTransferOwnership(e deployment.Environment, in deployops.TransferOwnershipPerChainInput) (bool, error) {
	chain, ok := e.BlockChains.SolanaChains()[in.ChainSelector]
	if !ok {
		return false, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
	}
	// Only accept ownership if the proposed owner is either the timelock or the deployer
	if solana.MustPublicKeyFromBase58(in.ProposedOwner) != a.timelockAddr[in.ChainSelector] && solana.MustPublicKeyFromBase58(in.ProposedOwner) != chain.DeployerKey.PublicKey() {
		return false, nil
	}
	return true, nil
}

func (a *SolanaAdapter) SequenceAcceptOwnership() *cldf_ops.Sequence[deployops.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"seq-accept-ownership",
		semver.MustParse("1.0.0"),
		"Accepts ownership of contracts",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in deployops.TransferOwnershipPerChainInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.SolanaChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}

			for _, contractRef := range in.ContractRef {
				switch contractRef.Type.String() {
				case routerops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, routerops.AcceptOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case offrampops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, offrampops.AcceptOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case feequoterops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, feequoterops.AcceptOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				case rmnremoteops.ContractType.String():
					report, err := cldf_ops.ExecuteOperation(b, rmnremoteops.AcceptOwnership, chain, utils.TransferOwnershipParams{
						Program:      solana.MustPublicKeyFromBase58(contractRef.Address),
						CurrentOwner: solana.MustPublicKeyFromBase58(in.CurrentOwner),
						NewOwner:     solana.MustPublicKeyFromBase58(in.ProposedOwner),
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("unsupported contract type %s for ownership transfer via MCMS on Solana", contractRef.Type)
				}
			}
			return output, nil
		})
}
