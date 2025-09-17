package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmsevmsdk "github.com/smartcontractkit/mcms/sdk/evm"
	"github.com/smartcontractkit/mcms/types"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
)

type MCMSRole string

var (
	BypasserManyChainMultisig  datastore.ContractType = "BypasserManyChainMultiSig"
	CancellerManyChainMultisig datastore.ContractType = "CancellerManyChainMultiSig"
	ProposerManyChainMultisig  datastore.ContractType = "ProposerManyChainMultiSig"
	ManyChainMultisig          datastore.ContractType = "ManyChainMultiSig"
	RBACTimelock               datastore.ContractType = "RBACTimelock"
	CallProxy                  datastore.ContractType = "CallProxy"

	// roles
	ProposerRole  MCMSRole = "PROPOSER"
	BypasserRole  MCMSRole = "BYPASSER"
	CancellerRole MCMSRole = "CANCELLER"
)

func (r MCMSRole) String() string {
	return string(r)
}

type WithMCMS interface {
	MCMSConfig() *MCMSConfig
	TimelockAddressQualifier() string
	MCMSAddressQualifier() string
}

func ResolveMCMSParams[CFG WithMCMS](e cldf_deployment.Environment, cfg CFG) (MCMSParams, error) {
	if cfg.MCMSConfig() == nil {
		return MCMSParams{}, nil
	}
	// map of chain selector to timelock address
	timelockAddresses := make(map[types.ChainSelector]string)
	chainMetadata := make(map[types.ChainSelector]mcms_types.ChainMetadata)
	for sel := range e.BlockChains.EVMChains() {
		// find and format the MCMS related addresses
		timelockAddrs, err := datastore_utils.FindAndFormatEachRef(e.DataStore, []datastore.AddressRef{
			{
				ChainSelector: sel,
				Type:          datastore.ContractType(RBACTimelock),
				Version:       utils.Version1_0_0,
				Qualifier:     cfg.TimelockAddressQualifier(),
			},
		}, datastore_utils.ToEVMAddress)
		if err != nil {
			e.Logger.Errorf("failed to find timelock addresses for chain selector %d: %v", sel, err)
			return MCMSParams{}, err
		}
		if len(timelockAddrs) == 0 {
			e.Logger.Errorf("no timelock address found for chain selector %d with qualifier %s", sel, cfg.TimelockAddressQualifier())
			return MCMSParams{}, fmt.Errorf("no timelock address found for chain selector %d with qualifier %s", sel, cfg.TimelockAddressQualifier())
		}
		timelockAddresses[types.ChainSelector(sel)] = timelockAddrs[0].String()
		var mcmsAddr common.Address
		switch cfg.MCMSConfig().TimelockAction {
		case mcms_types.TimelockActionSchedule:
			mcmsAddr, err = findMCMS(e, sel, ProposerRole, ProposerManyChainMultisig, cfg.MCMSAddressQualifier())
			if err != nil {
				return MCMSParams{}, err
			}
			e.Logger.Infof("Using proposer MCMS %s to schedule on timelock %s for chain selector %d", mcmsAddr.String(), timelockAddrs[0].String(), sel)
		case mcms_types.TimelockActionBypass:
			mcmsAddr, err = findMCMS(e, sel, BypasserRole, BypasserManyChainMultisig, cfg.MCMSAddressQualifier())
			if err != nil {
				return MCMSParams{}, err
			}
			e.Logger.Infof("Using bypasser MCMS %s to bypass on timelock %s for chain selector %d", mcmsAddr.String(), timelockAddrs[0].String(), sel)
		case mcms_types.TimelockActionCancel:
			mcmsAddr, err = findMCMS(e, sel, CancellerRole, CancellerManyChainMultisig, cfg.MCMSAddressQualifier())
			if err != nil {
				return MCMSParams{}, err
			}
			e.Logger.Infof("Using canceller MCMS %s to cancel on timelock %s for chain selector %d", mcmsAddr.String(), timelockAddrs[0].String(), sel)
		}
		chainMetadata[types.ChainSelector(sel)], err = mcmsChainMetadata(e, sel, mcmsAddr)
		if err != nil {
			e.Logger.Errorf("failed to get chain metadata for chain selector %d: %v", sel, err)
			return MCMSParams{}, err
		}
	}

	return MCMSParams{
		MCMSConfig:        cfg.MCMSConfig(),
		TimelockAddresses: timelockAddresses,
		ChainMetadata:     chainMetadata,
	}, nil
}

// findMCMS finds the MCMS address for the given chain selector, role, and contract type.
// It first looks for addresses stored with type ManyChainMultisig , label PROPOSER/BYPASSER/CANCELLER and qualifier
// If not found, it falls back to type ProposerManyChainMultiSig/BypasserManyChainMultiSig/CancellerManyChainMultiSig and qualifier.
func findMCMS(e cldf_deployment.Environment, sel uint64, role MCMSRole, contractType datastore.ContractType, qualifier string) (common.Address, error) {
	// addresses could be stored with type ManyChainMultisig and label PROPOSER/BYPASSER/CANCELLER
	// or with type ProposerManyChainMultiSig/BypasserManyChainMultiSig/CancellerManyChainMultiSig and no label
	staticAddrs, err := datastore_utils.FindAndFormatEachRefIfFound(e.DataStore, []datastore.AddressRef{
		{
			ChainSelector: sel,
			Type:          ManyChainMultisig,
			Version:       utils.Version1_0_0,
			Qualifier:     qualifier,
			Labels:        datastore.NewLabelSet(role.String()),
		},
	}, datastore_utils.ToEVMAddress)
	if err != nil {
		e.Logger.Errorf("failed to find and format %s mcms address for chain selector %d: %v", contractType, sel, err)
		return common.Address{}, err
	}
	if len(staticAddrs) == 1 {
		return staticAddrs[0], nil
	} else if len(staticAddrs) > 1 {
		e.Logger.Errorf("multiple %s mcms addresses found for chain selector %d with qualifier %s and role %s", contractType, sel, qualifier, role)
		return common.Address{}, fmt.Errorf("multiple %s mcms addresses found for chain selector %d with qualifier %s and role %s", contractType, sel, qualifier, role)
	}
	// fallback to type ProposerManyChainMultiSig/BypasserManyChainMultiSig/CancellerManyChainMultiSig if not found
	staticAddrs, err = datastore_utils.FindAndFormatEachRefIfFound(e.DataStore, []datastore.AddressRef{
		{
			ChainSelector: sel,
			Type:          contractType,
			Version:       utils.Version1_0_0,
			Qualifier:     qualifier,
		},
	}, datastore_utils.ToEVMAddress)
	if err != nil {
		e.Logger.Errorf("failed to find and format %s mcms address for chain selector %d: %v", contractType, sel, err)
		return common.Address{}, err
	}
	if len(staticAddrs) == 0 {
		e.Logger.Errorf("no %s mcms address found for chain selector %d with qualifier %s", contractType, sel, qualifier)
		return common.Address{}, fmt.Errorf("no %s mcms address found for chain selector %d with qualifier %s", contractType, sel, qualifier)
	}
	return staticAddrs[0], nil
}

// mcmsChainMetadata retrieves the starting operation count for the given MCMS address on the specified chain.
func mcmsChainMetadata(e cldf_deployment.Environment, chain uint64, mcmsAddr common.Address) (mcms_types.ChainMetadata, error) {
	evmChains := e.BlockChains.EVMChains()
	if evmChains == nil {
		return mcms_types.ChainMetadata{}, fmt.Errorf("no EVM chains found in environment")
	}
	inspectorForChain := mcmsevmsdk.NewInspector(evmChains[chain].Client)
	opCount, err := inspectorForChain.GetOpCount(e.GetContext(), mcmsAddr.String())
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	return mcms_types.ChainMetadata{
		StartingOpCount: opCount,
		MCMAddress:      mcmsAddr.String(),
	}, nil
}
