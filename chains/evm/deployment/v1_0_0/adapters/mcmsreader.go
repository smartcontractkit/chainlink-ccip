package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	mcmsevmsdk "github.com/smartcontractkit/mcms/sdk/evm"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type EVMMCMSReader struct{}

func (r *EVMMCMSReader) GetChainMetadata(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (mcms_types.ChainMetadata, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return mcms_types.ChainMetadata{}, fmt.Errorf("EVM chain with selector %d not found", chainSelector)
	}
	inspector := mcmsevmsdk.NewInspector(chain.Client)
	mcmsAddr, err := r.GetMCMSRef(e, chainSelector, input)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to get MCMS address for chain %d: %w", chainSelector, err)
	}
	addrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, mcmsAddr, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to find MCMS address for chain %d: %w", chainSelector, err)
	}
	mcmAddr := addrRef.Hex()
	if mcmAddr == "" {
		return mcms_types.ChainMetadata{}, fmt.Errorf("MCMS address not found for chain %d", chainSelector)
	}
	// starting opCount
	opCount, err := inspector.GetOpCount(e.GetContext(), mcmAddr)
	if err != nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("failed to get opCount for MCMS at address %s on chain %d: %w", mcmAddr, chainSelector, err)
	}
	return mcms_types.ChainMetadata{
		StartingOpCount: opCount,
		MCMAddress:      mcmAddr,
	}, nil
}

func (r *EVMMCMSReader) GetTimelockRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (datastore.AddressRef, error) {
	timelockRef := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		utils.RBACTimelock,
		Version,
		input.Qualifier,
	)
	return timelockRef, nil
}

func (r *EVMMCMSReader) GetMCMSRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (datastore.AddressRef, error) {
	// find mcms address
	// populate contract type from TimelockAction
	var addrType datastore.ContractType
	switch input.TimelockAction {
	case mcms_types.TimelockActionSchedule:
		addrType = datastore.ContractType(utils.ProposerManyChainMultisig)
	case mcms_types.TimelockActionBypass:
		addrType = datastore.ContractType(utils.BypasserManyChainMultisig)
	case mcms_types.TimelockActionCancel:
		addrType = datastore.ContractType(utils.CancellerManyChainMultisig)
	default:
		return datastore.AddressRef{}, fmt.Errorf("unsupported timelock action type: %s", input.TimelockAction)
	}
	mcmAddress := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chainSelector,
		deployment.ContractType(addrType),
		Version,
		input.Qualifier,
	)
	return mcmAddress, nil
}
