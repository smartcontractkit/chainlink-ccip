package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	sequtil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
)

var Version = semver.MustParse("1.0.0")

type EVMDeployer struct{}

func (a *EVMDeployer) DeployChainContracts() *cldf_ops.Sequence[ccipapi.ContractDeploymentConfigPerChainWithAddress, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	// Not implemented for the 1.0.0 deployer
	return nil
}

func (a *EVMDeployer) SetOCR3Config() *cldf_ops.Sequence[ccipapi.SetOCR3ConfigInput, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	// Not implemented for the 1.0.0 deployer
	return nil
}

func (a *EVMDeployer) UpdateMCMSConfig() *cldf_ops.Sequence[ccipapi.UpdateMCMSConfigInputPerChainWithSelector, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"update-mcms-config",
		semver.MustParse("1.0.0"),
		"Updates MCMS Configs of the specified contracts with the specified configs",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in ccipapi.UpdateMCMSConfigInputPerChainWithSelector) (output sequtil.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequtil.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}

			// create sequence input
			seqInput := seq.SeqSetMCMSConfigInput{
				ChainSelector: in.ChainSelector,
				MCMConfig:     &in.MCMConfig,
				MCMContracts:  in.MCMContracts,
			}
			report, err := cldf_ops.ExecuteSequence(b, seq.SeqSetMCMSConfigs, evmChain, seqInput)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to update mcms config on chain %d: %w", in.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, report.Output.BatchOps...)

			return output, nil
		})
}

func (a *EVMDeployer) GrantAdminRoleToTimelock() *cldf_ops.Sequence[ccipapi.GrantAdminRoleToTimelockConfigPerChainWithSelector, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"grant-admin-role-of-timelock-to-timelock",
		semver.MustParse("1.0.0"),
		"Grants admin role of specified timelock contract to the other specified timelock and renounces admin role of the deployer key",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in ccipapi.GrantAdminRoleToTimelockConfigPerChainWithSelector) (output sequtil.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequtil.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}

			// create sequence input
			seqInput := seq.SeqGrantAdminRoleOfTimelockToTimelockInput{
				ChainSelector:           in.ChainSelector,
				TimelockAddress:         common.HexToAddress(in.TimelockToTransferRef.Address),
				NewAdminTimelockAddress: common.HexToAddress(in.NewAdminTimelockRef.Address),
			}
			report, err := cldf_ops.ExecuteSequence(b, seq.SeqGrantAdminRoleOfTimelockToTimelock, evmChain, seqInput)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to grant admin role to timelock on chain %d: %w", in.ChainSelector, err)
			}

			return report.Output, nil
		})
}

func (d *EVMDeployer) DeployMCMS() *cldf_ops.Sequence[ccipapi.MCMSDeploymentConfigPerChainWithAddress, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"deploy-mcms",
		semver.MustParse("1.0.0"),
		"Deploys all MCM contracts with config",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in ccipapi.MCMSDeploymentConfigPerChainWithAddress) (output sequtil.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequtil.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			// deploy and configure the proposer MCM
			seqInput := seq.SeqMCMSDeploymentCfg{
				ChainSelector:     in.ChainSelector,
				ContractType:      utils.ProposerManyChainMultisig,
				MCMConfig:         &in.Proposer,
				Qualifier:         in.Qualifier,
				Label:             in.Label,
				ExistingAddresses: in.ExistingAddresses,
			}
			report, err := cldf_ops.ExecuteSequence(b, seq.SeqDeployMCMWithConfig, evmChain, seqInput)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy and configure proposer MCM on chain %d: %w", in.ChainSelector, err)
			}
			if len(report.Output.Addresses) == 0 {
				return sequtil.OnChainOutput{}, fmt.Errorf("no proposer MCM address returned from deployment on chain %d", in.ChainSelector)
			}
			output.Addresses = append(output.Addresses, report.Output.Addresses...)
			proposerAddr := report.Output.Addresses[0]
			// deploy and configure the bypasser MCM
			seqInput = seq.SeqMCMSDeploymentCfg{
				ChainSelector:     in.ChainSelector,
				ContractType:      utils.BypasserManyChainMultisig,
				MCMConfig:         &in.Bypasser,
				Qualifier:         in.Qualifier,
				Label:             in.Label,
				ExistingAddresses: in.ExistingAddresses,
			}
			report, err = cldf_ops.ExecuteSequence(b, seq.SeqDeployMCMWithConfig, evmChain, seqInput)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy and configure bypasser MCM on chain %d: %w", in.ChainSelector, err)
			}
			if len(report.Output.Addresses) == 0 {
				return sequtil.OnChainOutput{}, fmt.Errorf("no bypasser MCM address returned from deployment on chain %d", in.ChainSelector)
			}
			output.Addresses = append(output.Addresses, report.Output.Addresses...)
			bypasserAddr := report.Output.Addresses[0]
			// deploy and configure the canceller MCM
			seqInput = seq.SeqMCMSDeploymentCfg{
				ChainSelector:     in.ChainSelector,
				ContractType:      utils.CancellerManyChainMultisig,
				MCMConfig:         &in.Canceller,
				Qualifier:         in.Qualifier,
				Label:             in.Label,
				ExistingAddresses: in.ExistingAddresses,
			}
			report, err = cldf_ops.ExecuteSequence(b, seq.SeqDeployMCMWithConfig, evmChain, seqInput)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy and configure canceller MCM on chain %d: %w", in.ChainSelector, err)
			}
			if len(report.Output.Addresses) == 0 {
				return sequtil.OnChainOutput{}, fmt.Errorf("no canceller MCM address returned from deployment on chain %d", in.ChainSelector)
			}
			output.Addresses = append(output.Addresses, report.Output.Addresses...)
			cancellerAddr := report.Output.Addresses[0]

			// deploy timelock — always use the deployer key as the initial admin
			// so we can manage roles immediately after deployment
			timelockAddr, err := contract.MaybeDeployContract(b, ops.OpDeployTimelock, evmChain, contract.DeployInput[ops.OpDeployTimelockInput]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.RBACTimelock, *ops.MCMSVersion),
				Args: ops.OpDeployTimelockInput{
					TimelockMinDelay: in.TimelockMinDelay,
					Proposers:        []common.Address{common.HexToAddress(proposerAddr.Address)},
					Bypassers:        []common.Address{common.HexToAddress(bypasserAddr.Address)},
					Cancellers:       []common.Address{common.HexToAddress(cancellerAddr.Address)},
					Admin:            evmChain.DeployerKey.From,
					// Add Executor later after call proxy is deployed
					Executors: []common.Address{},
				},
			}, in.ExistingAddresses)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy timelock on chain %d: %w", in.ChainSelector, err)
			}
			b.Logger.Infof("Deployed Timelock at address %s on chain %s", timelockAddr, evmChain.Name)
			// deploy call proxy with timelock
			callProxyAddr, err := contract.MaybeDeployContract(b, ops.OpDeployCallProxy, evmChain, contract.DeployInput[ops.OpDeployCallProxyInput]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.CallProxy, *ops.MCMSVersion),
				Args: ops.OpDeployCallProxyInput{
					TimelockAddress: common.HexToAddress(timelockAddr.Address),
				},
			}, in.ExistingAddresses)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy call proxy on chain %d: %w", in.ChainSelector, err)
			}
			b.Logger.Infof("Deployed Call Proxy at address %s on chain %s", callProxyAddr, evmChain.Name)
			output.Addresses = append(output.Addresses, timelockAddr, callProxyAddr)

			// now that call proxy is deployed, we can add it as executor to the timelock
			_, err = cldf_ops.ExecuteOperation(b, ops.OpGrantRoleTimelock, evmChain, contract.FunctionInput[ops.OpGrantRoleTimelockInput]{
				ChainSelector: in.ChainSelector,
				Address:       common.HexToAddress(timelockAddr.Address),
				Args: ops.OpGrantRoleTimelockInput{
					RoleID:  ops.EXECUTOR_ROLE.ID,
					Account: common.HexToAddress(callProxyAddr.Address),
				},
			})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to grant executor role to call proxy on timelock on chain %d: %w", in.ChainSelector, err)
			}
			b.Logger.Infof("Granted Executor role on Timelock %s to Call Proxy %s on chain %s", timelockAddr, callProxyAddr, evmChain.Name)

			// Determine the final ADMIN_ROLE holder for the timelock.
			// CLLCCIP is always self-governed (its timelock is its own admin).
			// Every other MCMS instance MUST use the existing CLLCCIP RBACTimelock as its
			// admin — if CLLCCIP has not been deployed yet, fail fast.
			isCLLCCIP := in.Qualifier != nil && *in.Qualifier == utils.CLLQualifier
			finalAdmin := common.HexToAddress(timelockAddr.Address) // default: self-governed (CLLCCIP case)
			if !isCLLCCIP {
				existingDS := cldf_datastore.NewMemoryDataStore()
				for _, ref := range in.ExistingAddresses {
					_ = existingDS.Addresses().Add(ref)
				}
				cllTimelockAddr, lookupErr := datastore_utils.FindAndFormatRef(
					existingDS.Seal(),
					cldf_datastore.AddressRef{
						Type:      cldf_datastore.ContractType(utils.RBACTimelock),
						Qualifier: utils.CLLQualifier,
					},
					in.ChainSelector,
					evm_datastore_utils.ToEVMAddress,
				)
				if lookupErr != nil || cllTimelockAddr == (common.Address{}) {
					qualifier := ""
					if in.Qualifier != nil {
						qualifier = *in.Qualifier
					}
					return sequtil.OnChainOutput{}, fmt.Errorf(
						"cannot deploy MCMS with qualifier %q on chain %d: CLLCCIP RBACTimelock must be deployed first (it will be set as admin)",
						qualifier, in.ChainSelector,
					)
				}
				finalAdmin = cllTimelockAddr
			}

			// If finalAdmin differs from the deployer key we need to transfer the ADMIN_ROLE.
			// RBACTimelock uses RBAC (not Ownable2Step), so:
			//   "transfer" = grantRole(ADMIN_ROLE, finalAdmin)
			//   "accept"   = renounceRole(ADMIN_ROLE) by the deployer — both happen in the same changeset.
			if finalAdmin != evmChain.DeployerKey.From {
				_, err = cldf_ops.ExecuteOperation(b, ops.OpGrantRoleTimelock, evmChain, contract.FunctionInput[ops.OpGrantRoleTimelockInput]{
					ChainSelector: in.ChainSelector,
					Address:       common.HexToAddress(timelockAddr.Address),
					Args: ops.OpGrantRoleTimelockInput{
						RoleID:  ops.ADMIN_ROLE.ID,
						Account: finalAdmin,
					},
				})
				if err != nil {
					return sequtil.OnChainOutput{}, fmt.Errorf("failed to grant admin role to %s on timelock on chain %d: %w", finalAdmin, in.ChainSelector, err)
				}
				b.Logger.Infof("Granted Admin role on Timelock %s to %s on chain %s", timelockAddr.Address, finalAdmin, evmChain.Name)

				_, err = cldf_ops.ExecuteOperation(b, ops.OpRenounceRoleTimelock, evmChain, contract.FunctionInput[ops.OpRenounceRoleTimelockInput]{
					ChainSelector: in.ChainSelector,
					Address:       common.HexToAddress(timelockAddr.Address),
					Args: ops.OpRenounceRoleTimelockInput{
						RoleID: ops.ADMIN_ROLE.ID,
					},
				})
				if err != nil {
					return sequtil.OnChainOutput{}, fmt.Errorf("failed to renounce admin role on timelock on chain %d: %w", in.ChainSelector, err)
				}
				b.Logger.Infof("Deployer renounced Admin role on Timelock %s on chain %s", timelockAddr.Address, evmChain.Name)
			}

			mcmOwnershipBatchOps, err := seq.TransferDeployedMCMsToTimelock(
				b,
				evmChain,
				in.ChainSelector,
				common.HexToAddress(timelockAddr.Address),
				[]cldf_datastore.AddressRef{proposerAddr, bypasserAddr, cancellerAddr},
			)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to transfer MCM ownership to timelock on chain %d: %w", in.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, mcmOwnershipBatchOps...)

			return output, nil
		})
}

func (d *EVMDeployer) FinalizeDeployMCMS() *cldf_ops.Sequence[ccipapi.MCMSDeploymentConfigPerChainWithAddress, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"finalize-deploy-mcms",
		semver.MustParse("1.0.0"),
		"On EVM, finalizing MCM deployment is a no-op",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in ccipapi.MCMSDeploymentConfigPerChainWithAddress) (output sequtil.OnChainOutput, err error) {
			return output, nil
		})
}
