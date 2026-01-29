package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
	"github.com/smartcontractkit/mcms/sdk"
	"github.com/smartcontractkit/mcms/sdk/evm"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
)

type SeqMCMSDeploymentCfg struct {
	ChainSelector     uint64
	ContractType      cldf.ContractType
	MCMConfig         *types.Config
	Label             *string
	Qualifier         *string
	ExistingAddresses []datastore.AddressRef
}

type SeqTransferMCMOwnershipToTimelockInput struct {
	ChainSelector uint64
	Contracts     []ops.OpTransferOwnershipInput
}

type SeqGrantAdminRoleOfTimelockToTimelockInput struct {
	ChainSelector           uint64
	TimelockAddress         common.Address
	NewAdminTimelockAddress common.Address
}

var SeqDeployMCMWithConfig = cldf_ops.NewSequence(
	"seq-deploy-mcm-with-config",
	semver.MustParse("1.0.0"),
	"Deploys MCM contract & sets config",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqMCMSDeploymentCfg) (output sequences.OnChainOutput, err error) {
		// Deploy MCM contracts
		var mcmAddr datastore.AddressRef
		switch in.ContractType {
		case utils.ProposerManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployProposerMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.ProposerManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		case utils.BypasserManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployBypasserMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.BypasserManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		case utils.CancellerManyChainMultisig:
			mcmAddr, err = contract.MaybeDeployContract(b, ops.OpDeployCancellerMCM, chain, contract.DeployInput[struct{}]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.CancellerManyChainMultisig, *semver.MustParse("1.0.0")),
			}, in.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported contract type for seq-deploy-mcm-with-config: %s", in.ContractType)
		}

		// Set config
		groupQuorums, groupParents, signerAddresses, signerGroups, err := sdk.ExtractSetConfigInputs(in.MCMConfig)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		_, err = cldf_ops.ExecuteOperation(b, ops.OpEVMSetConfigMCM, chain,
			contract.FunctionInput[ops.OpSetConfigMCMInput]{
				ChainSelector: in.ChainSelector,
				Address:       common.HexToAddress(mcmAddr.Address),
				Args: ops.OpSetConfigMCMInput{
					SignerAddresses: signerAddresses,
					SignerGroups:    signerGroups,
					GroupQuorums:    groupQuorums,
					GroupParents:    groupParents,
				},
			})

		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		b.Logger.Infof("Deployed %s at address %s on chain %s", in.ContractType, mcmAddr.Address, chain.Name)
		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{mcmAddr},
		}, nil
	},
)

var SeqGrantAdminRoleOfTimelockToTimelock = cldf_ops.NewSequence(
	"seq-grant-admin-role-of-timelock-to-timelock",
	semver.MustParse("1.0.0"),
	"Grants admin role of specified timelock contract to the other specified timelock and renounces admin role of the deployer key",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqGrantAdminRoleOfTimelockToTimelockInput) (output sequences.OnChainOutput, err error) {
		// Load the Timelock contract
		timelock, err := LoadTimelockContract(in.TimelockAddress, chain.Client)
		if err != nil {
			b.Logger.Errorf("failed to load timelock contract %s: %v", in.TimelockAddress, err)
			return output, fmt.Errorf("error loading timclock contract %s: %w", in.TimelockAddress, err)
		}

		// Verify that admin of Timelock contract is the Deployer EOA
		callerHasRole, err := timelock.HasRole(nil, ops.ADMIN_ROLE.ID, chain.DeployerKey.From)
		if err != nil {
			b.Logger.Errorf("failed to check whether caller %s is admin on timelock contract %s: %v", chain.DeployerKey.From, in.TimelockAddress, err)
			return output, fmt.Errorf("failed to check whether caller %s is admin on timelock contract %s: %w", chain.DeployerKey.From, in.TimelockAddress, err)
		}
		if !callerHasRole {
			b.Logger.Errorf("caller %s is not admin on timelock contract %s: %v", chain.DeployerKey.From, in.TimelockAddress, err)
			return output, fmt.Errorf("caller %s is not admin on timelock contract %s: %w", chain.DeployerKey.From, in.TimelockAddress, err)
		}

		newAdminTimelockHasRole, err := timelock.HasRole(nil, ops.ADMIN_ROLE.ID, in.NewAdminTimelockAddress)
		if err != nil {
			b.Logger.Errorf("failed to check whether new timelock owner %s is admin on timelock contract %s: %v", in.NewAdminTimelockAddress, in.TimelockAddress, err)
			return output, fmt.Errorf("failed to check whether new timelock owner %s is admin on timelock contract %s: %w", in.NewAdminTimelockAddress, in.TimelockAddress, err)
		}

		// Grant admin role to new admin Timelock
		if !newAdminTimelockHasRole {
			_, err = cldf_ops.ExecuteOperation(b, ops.OpGrantRoleTimelock, chain, contract.FunctionInput[ops.OpGrantRoleTimelockInput]{
				ChainSelector: in.ChainSelector,
				Address:       in.TimelockAddress,
				Args: ops.OpGrantRoleTimelockInput{
					RoleID:  ops.ADMIN_ROLE.ID,
					Account: in.NewAdminTimelockAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to grant admin role to new admin timelock on chain %d: %w", in.ChainSelector, err)
			}
			b.Logger.Infof("Granted Admin role on Timelock %s to Timelock %s on chain %s", in.TimelockAddress, in.NewAdminTimelockAddress, chain)
		} else {
			b.Logger.Infof("Timelock %s is already admin on Timelock %s on chain %s", in.NewAdminTimelockAddress, in.TimelockAddress, chain)
		}

		// Renounce admin role from Deployer EOA
		_, err = cldf_ops.ExecuteOperation(b, ops.OpRenounceRoleTimelock, chain, contract.FunctionInput[ops.OpRenounceRoleTimelockInput]{
			ChainSelector: in.ChainSelector,
			Address:       in.TimelockAddress,
			Args: ops.OpRenounceRoleTimelockInput{
				RoleID: ops.ADMIN_ROLE.ID,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to renounce admin role on timelock contract on chain %d: %w", in.ChainSelector, err)
		}
		b.Logger.Infof("Renounced Admin role on Timelock %s on chain %s", in.TimelockAddress, chain)

		return sequences.OnChainOutput{}, nil
	})

var SeqTransferMCMOwnershipToTimelock = cldf_ops.NewSequence(
	"seq-transfer-mcm-ownership-to-timelock",
	semver.MustParse("1.0.0"),
	"Transfers ownership of MCM contract to the specified timelock",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqTransferMCMOwnershipToTimelockInput) (output sequences.OnChainOutput, err error) {
		for _, contractInput := range in.Contracts {
			contractAddr := contractInput.Address.Hex()
			owner, c, err := LoadOwnableContract(contractInput.Address, chain.Client)
			if err != nil {
				b.Logger.Errorf("failed to load ownable contract %s: %v", contractAddr, err)
				return output, fmt.Errorf("error loading ownable contract %s: %w", contractAddr, err)
			}
			if owner == contractInput.ProposedOwner {
				b.Logger.Infof("ownership of contract %s on chain %s is already set to %s, skipping transfer",
					contractAddr, chain.Name(), contractInput.ProposedOwner.Hex())
				continue
			}
			deps := ops.OpEVMOwnershipDeps{
				Chain:    chain,
				OwnableC: c,
			}
			report, err := cldf_ops.ExecuteOperation(b, ops.OpTransferOwnership, deps,
				ops.OpTransferOwnershipInput{
					ChainSelector:   in.ChainSelector,
					Address:         contractInput.Address,
					ProposedOwner:   contractInput.ProposedOwner,
					ContractType:    contractInput.ContractType,
					TimelockAddress: contractInput.TimelockAddress,
				})
			if err != nil {
				b.Logger.Errorf("failed to transfer ownership of contract %s on chain %s: %v", contractAddr, chain.Name(), err)
				return output, fmt.Errorf("error transferring ownership of contract %s on chain %s: %w", contractAddr, chain.Name(), err)
			}
			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}
		return output, nil
	})

var SeqAcceptMCMOwnershipFromTimelock = cldf_ops.NewSequence(
	"seq-accept-mcm-ownership-from-timelock",
	semver.MustParse("1.0.0"),
	"Accepts ownership of MCM contract from the specified timelock",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqTransferMCMOwnershipToTimelockInput) (output sequences.OnChainOutput, err error) {
		for _, contractInput := range in.Contracts {
			contractAddr := contractInput.Address.Hex()
			owner, c, err := LoadOwnableContract(contractInput.Address, chain.Client)
			if err != nil {
				b.Logger.Errorf("failed to load ownable contract %s: %v", contractAddr, err)
				return output, fmt.Errorf("error loading ownable contract %s: %w", contractAddr, err)
			}
			if owner == contractInput.ProposedOwner {
				b.Logger.Infof("ownership of contract %s on chain %s is already set to %s, skipping acceptance",
					contractAddr, chain.Name(), contractInput.ProposedOwner.Hex())
				continue
			}
			deps := ops.OpEVMOwnershipDeps{
				Chain:    chain,
				OwnableC: c,
			}
			report, err := cldf_ops.ExecuteOperation(b, ops.OpAcceptOwnership, deps,
				ops.OpTransferOwnershipInput{
					ChainSelector:   in.ChainSelector,
					Address:         contractInput.Address,
					ProposedOwner:   contractInput.ProposedOwner,
					ContractType:    contractInput.ContractType,
					TimelockAddress: contractInput.TimelockAddress,
				})
			if err != nil {
				b.Logger.Errorf("failed to accept ownership of contract %s on chain %s: %v", contractAddr, chain.Name(), err)
				return output, fmt.Errorf("error accepting ownership of contract %s on chain %s: %w", contractAddr, chain.Name(), err)
			}
			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}
		return output, nil
	})

func LoadOwnableContract(addr common.Address, client bind.ContractBackend) (common.Address, ops.OwnershipTranferable, error) {
	// Just using the ownership interface from here.
	c, err := burn_mint_erc677.NewBurnMintERC677(addr, client)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create contract: %w", err)
	}
	owner, err := c.Owner(nil)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to get owner of contract %s: %w", c.Address(), err)
	}

	return owner, c, nil
}

func LoadTimelockContract(addr common.Address, client bind.ContractBackend) (*bindings.RBACTimelock, error) {
	c, err := bindings.NewRBACTimelock(addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load timelock contract: %w", err)
	}

	return c, nil
}
