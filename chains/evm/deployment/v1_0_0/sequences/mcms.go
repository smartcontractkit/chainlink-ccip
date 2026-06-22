package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/sdk"
	"github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

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

type SeqGrantAdminRoleOfTimelockToTimelockInput struct {
	ChainSelector           uint64
	TimelockAddress         common.Address
	NewAdminTimelockAddress common.Address
}

type SeqSetMCMSConfigInput struct {
	ChainSelector uint64
	MCMConfig     *types.Config
	MCMContracts  []datastore.AddressRef
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

var SeqSetMCMSConfigs = cldf_ops.NewSequence(
	"seq-set-mcm-config",
	semver.MustParse("1.0.0"),
	"Sets config on previously deployed MCM contract",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqSetMCMSConfigInput) (output sequences.OnChainOutput, err error) {
		for _, mcmContract := range in.MCMContracts {
			// Set config on contract
			groupQuorums, groupParents, signerAddresses, signerGroups, err := sdk.ExtractSetConfigInputs(in.MCMConfig)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			report, err := cldf_ops.ExecuteOperation(b, ops.OpEVMSetConfigMCM, chain,
				contract.FunctionInput[ops.OpSetConfigMCMInput]{
					ChainSelector: in.ChainSelector,
					Address:       common.HexToAddress(mcmContract.Address),
					Args: ops.OpSetConfigMCMInput{
						SignerAddresses: signerAddresses,
						SignerGroups:    signerGroups,
						GroupQuorums:    groupQuorums,
						GroupParents:    groupParents,
					},
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update mcms config on chain %d for contract with address %s: %w",
					in.ChainSelector, mcmContract.Address, err)
			}
			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}

		return output, nil
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

		if err := TransferTimelockAdminTo(b, chain, in.ChainSelector, in.TimelockAddress, in.NewAdminTimelockAddress); err != nil {
			return sequences.OnChainOutput{}, err
		}

		return sequences.OnChainOutput{}, nil
	})

var SeqTransferMCMOwnershipToTimelock = cldf_ops.NewSequence(
	"seq-transfer-mcm-ownership-to-timelock",
	semver.MustParse("1.0.0"),
	"Transfers ownership of MCM contract to the specified timelock",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqTransferMCMOwnershipToTimelockInput) (output sequences.OnChainOutput, err error) {
		batchOps, err := TransferOwnership(b, chain, in.Contracts)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	})

var SeqAcceptMCMOwnershipFromTimelock = cldf_ops.NewSequence(
	"seq-accept-mcm-ownership-from-timelock",
	semver.MustParse("1.0.0"),
	"Accepts ownership of MCM contract from the specified timelock",
	func(b cldf_ops.Bundle, chain cldf_evm.Chain, in SeqTransferMCMOwnershipToTimelockInput) (output sequences.OnChainOutput, err error) {
		batchOps, err := AcceptOwnership(b, chain, in.Contracts)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	})
