package sequences

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	feequoterops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	mcmsops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/mcms"
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

func (a *SolanaAdapter) GetChainMetadata(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (mcms_types.ChainMetadata, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return mcms_types.ChainMetadata{}, fmt.Errorf("chain with selector %d not found in environment", chainSelector)
	}
	mcmAddress := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.McmProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	proposerSeed := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		common_utils.ProposerManyChainMultisig,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	proposer := mcms_solana.ContractAddress(
		solana.MustPublicKeyFromBase58(mcmAddress.Address),
		mcms_solana.PDASeed([]byte(proposerSeed.Address)),
	)
	inspector := mcms_solana.NewInspector(chain.Client)
	opcount, err := inspector.GetOpCount(context.Background(), proposer)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to get op count for chain %d: %w", chainSelector, err)
	}

	var instanceSeed mcms_solana.PDASeed
	switch input.TimelockAction {
	case mcms_types.TimelockActionSchedule:
		ref := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.ProposerManyChainMultisig,
			common_utils.Version_1_6_0,
			input.Qualifier,
		)
		instanceSeed = mcms_solana.PDASeed([]byte(ref.Address))
	case mcms_types.TimelockActionCancel:
		ref := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.CancellerManyChainMultisig,
			common_utils.Version_1_6_0,
			input.Qualifier,
		)
		instanceSeed = mcms_solana.PDASeed([]byte(ref.Address))
	case mcms_types.TimelockActionBypass:
		ref := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.BypasserManyChainMultisig,
			common_utils.Version_1_6_0,
			input.Qualifier,
		)
		instanceSeed = mcms_solana.PDASeed([]byte(ref.Address))
	default:
		return mcms_types.ChainMetadata{}, fmt.Errorf("unsupported timelock action %s for chain %d", input.TimelockAction, chainSelector)
	}
	proposerAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		mcmsops.ProposerAccessControllerAccount,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	cancellerAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		mcmsops.CancellerAccessControllerAccount,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	bypasserAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		mcmsops.BypasserAccessControllerAccount,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	metadata, err := mcms_solana.NewChainMetadata(
		opcount,
		solana.MustPublicKeyFromBase58(mcmAddress.Address),
		instanceSeed,
		solana.MustPublicKeyFromBase58(proposerAccount.Address),
		solana.MustPublicKeyFromBase58(cancellerAccount.Address),
		solana.MustPublicKeyFromBase58(bypasserAccount.Address))
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to create Solana MCMS chain metadata for chain %d: %w", chainSelector, err)
	}
	return metadata, nil
}

func (a *SolanaAdapter) GetTimelockRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	timelockRef := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.TimelockCompositeAddress,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	return timelockRef, nil
}

func (a *SolanaAdapter) GetMCMSRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	mcmAddress := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.McmProgramType,
		common_utils.Version_1_6_0,
		input.Qualifier,
	)
	return mcmAddress, nil
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
