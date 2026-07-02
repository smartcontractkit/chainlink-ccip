package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
)

// GrantTimelockAdminRole grants ADMIN_ROLE to account on timelock if it does not already hold it.
func GrantTimelockAdminRole(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	timelockAddr, account common.Address,
) error {
	timelock, err := LoadTimelockContract(timelockAddr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to load timelock contract %s: %w", timelockAddr, err)
	}
	hasRole, err := timelock.HasRole(nil, ops.ADMIN_ROLE.ID, account)
	if err != nil {
		return fmt.Errorf("failed to check whether %s is admin on timelock %s: %w", account, timelockAddr, err)
	}
	if hasRole {
		b.Logger.Infof("Timelock %s is already admin on Timelock %s on chain %s", account, timelockAddr, chain.Name)
		return nil
	}
	_, err = cldf_ops.ExecuteOperation(b, ops.OpGrantRoleTimelock, chain, contract.FunctionInput[ops.OpGrantRoleTimelockInput]{
		ChainSelector: chainSelector,
		Address:       timelockAddr,
		Args: ops.OpGrantRoleTimelockInput{
			RoleID:  ops.ADMIN_ROLE.ID,
			Account: account,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to grant admin role to %s on timelock %s on chain %d: %w", account, timelockAddr, chainSelector, err)
	}
	b.Logger.Infof("Granted Admin role on Timelock %s to %s on chain %s", timelockAddr, account, chain.Name)
	return nil
}

// RenounceDeployerTimelockAdmin renounces ADMIN_ROLE from the deployer on timelock if held.
func RenounceDeployerTimelockAdmin(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	timelockAddr common.Address,
) error {
	timelock, err := LoadTimelockContract(timelockAddr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to load timelock contract %s: %w", timelockAddr, err)
	}
	hasRole, err := timelock.HasRole(nil, ops.ADMIN_ROLE.ID, chain.DeployerKey.From)
	if err != nil {
		return fmt.Errorf("failed to check whether deployer %s is admin on timelock %s: %w", chain.DeployerKey.From, timelockAddr, err)
	}
	if !hasRole {
		b.Logger.Infof("Deployer is not admin on Timelock %s on chain %s, skipping renounce", timelockAddr, chain.Name)
		return nil
	}
	_, err = cldf_ops.ExecuteOperation(b, ops.OpRenounceRoleTimelock, chain, contract.FunctionInput[ops.OpRenounceRoleTimelockInput]{
		ChainSelector: chainSelector,
		Address:       timelockAddr,
		Args: ops.OpRenounceRoleTimelockInput{
			RoleID: ops.ADMIN_ROLE.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to renounce admin role on timelock %s on chain %d: %w", timelockAddr, chainSelector, err)
	}
	b.Logger.Infof("Renounced Admin role on Timelock %s on chain %s", timelockAddr, chain.Name)
	return nil
}

// TransferTimelockAdminTo grants ADMIN_ROLE to newAdmin and renounces it from the deployer.
func TransferTimelockAdminTo(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	timelockAddr, newAdmin common.Address,
) error {
	if err := GrantTimelockAdminRole(b, chain, chainSelector, timelockAddr, newAdmin); err != nil {
		return err
	}
	return RenounceDeployerTimelockAdmin(b, chain, chainSelector, timelockAddr)
}
