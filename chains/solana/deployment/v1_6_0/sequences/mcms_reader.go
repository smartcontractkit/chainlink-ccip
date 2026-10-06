package sequences

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
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

// getMCMSAccountRef resolves a 1.6.0 MCMS account ref, failing if it is not in the datastore. Used when
// building a proposal, where a missing account would otherwise surface as an unusable proposal.
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

// GetTimelockRef returns an empty ref (and no error) when the timelock is not in the datastore: callers
// such as token expansion use that to detect chains without MCMS and fall back to non-timelock admins.
func (a *SolanaAdapter) GetTimelockRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	return datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		mcmsQualifier(input),
	), nil
}

// GetMCMSRef returns the MCM program ref, or an empty ref when it is not in the datastore. The program is
// shared by every MCMS instance on the chain, so it is stored without a qualifier and the input qualifier
// does not apply.
func (a *SolanaAdapter) GetMCMSRef(e deployment.Environment, chainSelector uint64, _ mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	return datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.McmProgramType,
		common_utils.Version_1_6_0,
		"",
	), nil
}
