package sequences

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

func init() {
	mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, &SolanaAdapter{})
}

type SolanaAdapter struct {
	timelockAddr map[uint64]solana.PublicKey
}

func (a *SolanaAdapter) GetRouterAddress(ds cldf_datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, cldf_datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          cldf_datastore.ContractType(router.ContractType),
		Version:       router.Version, // TODO import 1.6.1 version when it exists
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetRMNRemoteAddress(ds cldf_datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, cldf_datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          cldf_datastore.ContractType(rmnremoteops.ContractType),
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetChainMetadata(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (mcms_types.ChainMetadata, error) {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return mcms_types.ChainMetadata{}, fmt.Errorf("chain with selector %d not found in environment", chainSelector)
	}

	inspector := mcms_solana.NewInspector(chain.Client)

	var id solana.PublicKey
	var seed mcms_solana.PDASeed
	var err error
	switch input.TimelockAction {
	case mcms_types.TimelockActionSchedule:
		addr := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.ProposerManyChainMultisig,
			common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
			input.Qualifier,
		)
		id, seed, err = mcms_solana.ParseContractAddress(addr.Address)
		if err != nil {
			return mcms_types.ChainMetadata{}, fmt.Errorf("failed to parse proposer address %s for chain %d: %w", addr.Address, chainSelector, err)
		}
	case mcms_types.TimelockActionCancel:
		addr := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.CancellerManyChainMultisig,
			common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
			input.Qualifier,
		)
		id, seed, err = mcms_solana.ParseContractAddress(addr.Address)
		if err != nil {
			return mcms_types.ChainMetadata{}, fmt.Errorf("failed to parse address %s for chain %d: %w", addr.Address, chainSelector, err)
		}
	case mcms_types.TimelockActionBypass:
		addr := datastore.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSelector,
			common_utils.BypasserManyChainMultisig,
			common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
			input.Qualifier,
		)
		id, seed, err = mcms_solana.ParseContractAddress(addr.Address)
		if err != nil {
			return mcms_types.ChainMetadata{}, fmt.Errorf("failed to parse address %s for chain %d: %w", addr.Address, chainSelector, err)
		}
	default:
		return mcms_types.ChainMetadata{}, fmt.Errorf("unsupported timelock action %s for chain %d", input.TimelockAction, chainSelector)
	}
	executor := mcms_solana.ContractAddress(
		id,
		seed,
	)
	opcount, err := inspector.GetOpCount(e.GetContext(), executor)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to get op count for chain %d: %w", chainSelector, err)
	}
	proposerAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.ProposerAccessControllerAccount,
		common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
		input.Qualifier,
	)
	cancellerAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.CancellerAccessControllerAccount,
		common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
		input.Qualifier,
	)
	bypasserAccount := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.BypasserAccessControllerAccount,
		common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
		input.Qualifier,
	)
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
	ref := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
		input.Qualifier,
	)
	return ref, nil
}

func (a *SolanaAdapter) GetMCMSRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (cldf_datastore.AddressRef, error) {
	mcmAddress := datastore.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.McmProgramType,
		common_utils.Version_1_6_0, // TODO use 1.6.1 version when it exists
		input.Qualifier,
	)
	return mcmAddress, nil
}
