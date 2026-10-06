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
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// mcmsQualifier returns the qualifier used to resolve the MCMS accounts, defaulting to CLLCCIP.
// Every MCMS account (MCM, timelock, access controllers) must come from the same qualifier: mixing
// them yields proposals the timelock rejects with InvalidAccessController.
func mcmsQualifier(input mcms_utils.Input) string {
	if input.Qualifier == "" {
		return common_utils.CLLQualifier
	}
	return input.Qualifier
}

// getMCMSAccountRef resolves a 1.6.0 MCMS account ref, failing if it is not in the datastore.
func getMCMSAccountRef(e deployment.Environment, chainSelector uint64, contractType deployment.ContractType, qualifier string) (cldf_datastore.AddressRef, error) {
	ref := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		contractType,
		common_utils.Version_1_6_0,
		qualifier,
	)
	if ref.Address == "" {
		return cldf_datastore.AddressRef{}, fmt.Errorf("%s with qualifier %q not found in datastore for chain %d", contractType, qualifier, chainSelector)
	}
	return ref, nil
}

func (a *SolanaAdapter) GetChainMetadata(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (mcms_types.ChainMetadata, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return mcms_types.ChainMetadata{}, fmt.Errorf("chain with selector %d not found in environment", chainSelector)
	}

	inspector := mcms_solana.NewInspector(chain.Client)
	qualifier := mcmsQualifier(input)

	var mcmType deployment.ContractType
	switch input.TimelockAction {
	case mcms_types.TimelockActionSchedule:
		mcmType = common_utils.ProposerManyChainMultisig
	case mcms_types.TimelockActionCancel:
		mcmType = common_utils.CancellerManyChainMultisig
	case mcms_types.TimelockActionBypass:
		mcmType = common_utils.BypasserManyChainMultisig
	default:
		return mcms_types.ChainMetadata{}, fmt.Errorf("unsupported timelock action %s for chain %d", input.TimelockAction, chainSelector)
	}
	mcmRef, err := getMCMSAccountRef(e, chainSelector, mcmType, qualifier)
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	id, seed, err := mcms_solana.ParseContractAddress(mcmRef.Address)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to parse %s address %s for chain %d: %w", mcmType, mcmRef.Address, chainSelector, err)
	}
	executor := mcms_solana.ContractAddress(
		id,
		seed,
	)
	opcount, err := inspector.GetOpCount(e.GetContext(), executor)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to get op count for chain %d: %w", chainSelector, err)
	}
	proposerAccount, err := getMCMSAccountRef(e, chainSelector, utils.ProposerAccessControllerAccount, qualifier)
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	cancellerAccount, err := getMCMSAccountRef(e, chainSelector, utils.CancellerAccessControllerAccount, qualifier)
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	bypasserAccount, err := getMCMSAccountRef(e, chainSelector, utils.BypasserAccessControllerAccount, qualifier)
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	metadata, err := mcms_solana.NewChainMetadata(
		opcount,
		id,
		seed,
		solana.MustPublicKeyFromBase58(proposerAccount.Address),
		solana.MustPublicKeyFromBase58(cancellerAccount.Address),
		solana.MustPublicKeyFromBase58(bypasserAccount.Address))
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to create Solana MCMS chain metadata for chain %d: %w", chainSelector, err)
	}
	return metadata, nil
}

func (a *SolanaAdapter) GetTimelockRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	return getMCMSAccountRef(e, chainSelector, common_utils.RBACTimelock, mcmsQualifier(input))
}

// GetMCMSRef returns the MCM program ref. The program is shared by every MCMS instance on the chain,
// so it is stored without a qualifier and the input qualifier does not apply.
func (a *SolanaAdapter) GetMCMSRef(e deployment.Environment, chainSelector uint64, _ mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	return getMCMSAccountRef(e, chainSelector, utils.McmProgramType, "")
}

func (a *SolanaAdapter) InitializeTimelockAddress(e deployment.Environment, input mcms.Input) error {
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
				// assume access controller will have all MCMS refs
				case utils.AccessControllerProgramType.String():
					report, err := transferAllMCMS(b, chain, in, true)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.BatchOps...)
				// assume rbac timelock will not have all MCMS refs
				case common_utils.RBACTimelock.String():
					report, err := transferAllMCMS(b, chain, in, false)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.BatchOps...)
				default:
					b.Logger.Debugf("unsupported contract type %s for ownership transfer via MCMS on Solana", contractRef.Type)
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
	return solana.MustPublicKeyFromBase58(in.CurrentOwner) == chain.DeployerKey.PublicKey(), nil
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
				// assume access controller will have all MCMS refs
				case utils.AccessControllerProgramType.String():
					report, err := acceptAllMCMS(b, chain, in, true)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.BatchOps...)
				// assume rbac timelock will not have all MCMS refs
				case common_utils.RBACTimelock.String():
					report, err := acceptAllMCMS(b, chain, in, false)
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership via MCMS on chain %d: %w", in.ChainSelector, err)
					}
					output.BatchOps = append(output.BatchOps, report.BatchOps...)
				default:
					b.Logger.Debugf("unsupported contract type %s for ownership transfer via MCMS on Solana", contractRef.Type)
				}
			}
			return output, nil
		})
}
