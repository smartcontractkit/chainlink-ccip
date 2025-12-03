package utils

import (
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func GetAllMCMS(
	chain cldf_solana.Chain,
	qualifier string,
	existingAddresses []cldf_datastore.AddressRef) []cldf_datastore.AddressRef {
	accessControllerProgram := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		AccessControllerProgramType,
		common_utils.Version_1_6_0,
		"",
	)
	proposerAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		ProposerAccessControllerAccount,
		common_utils.Version_1_6_0,
		qualifier,
	)
	executorAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		ExecutorAccessControllerAccount,
		common_utils.Version_1_6_0,
		qualifier,
	)
	cancellerAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		CancellerAccessControllerAccount,
		common_utils.Version_1_6_0,
		qualifier,
	)
	bypasserAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		BypasserAccessControllerAccount,
		common_utils.Version_1_6_0,
		qualifier,
	)
	timelockProgram := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		common_utils.RBACTimelock,
		common_utils.Version_1_6_0,
		qualifier,
	)
	proposerMCMSAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		common_utils.ProposerManyChainMultisig,
		common_utils.Version_1_6_0,
		qualifier,
	)
	cancellerMCMSAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		common_utils.CancellerManyChainMultisig,
		common_utils.Version_1_6_0,
		qualifier,
	)
	bypasserMCMSAccount := datastore.GetAddressRef(
		existingAddresses,
		chain.ChainSelector(),
		common_utils.BypasserManyChainMultisig,
		common_utils.Version_1_6_0,
		qualifier,
	)
	return []cldf_datastore.AddressRef{
		accessControllerProgram,
		proposerAccount,
		executorAccount,
		cancellerAccount,
		bypasserAccount,
		timelockProgram,
		proposerMCMSAccount,
		cancellerMCMSAccount,
		bypasserMCMSAccount,
	}
}
