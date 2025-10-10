package v1_0_0

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	sequtil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/v1_0"
)

type EVMDeployer struct{}

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
			// deploy timelock
			deployReport, err := cldf_ops.ExecuteOperation(b, ops.OpDeployTimelock, evmChain, contract.DeployInput[ops.OpDeployTimelockInput]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.RBACTimelock, *semver.MustParse("1.0.0")),
				Args: ops.OpDeployTimelockInput{
					TimelockMinDelay: in.TimelockMinDelay,
					Proposers:        []common.Address{common.HexToAddress(proposerAddr.Address)},
					Bypassers:        []common.Address{common.HexToAddress(bypasserAddr.Address)},
					Cancellers:       []common.Address{common.HexToAddress(cancellerAddr.Address)},
					Admin:            in.TimelockAdmin,
					// Add Executor later after call proxy is deployed
					Executors: []common.Address{},
				},
			})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy timelock on chain %d: %w", in.ChainSelector, err)
			}
			output.Addresses = append(output.Addresses, deployReport.Output)
			timelockAddr := deployReport.Output.Address
			b.Logger.Infof("Deployed Timelock at address %s on chain %s", timelockAddr, evmChain.Name)
			// deploy call proxy with timelock
			callProxyRep, err := cldf_ops.ExecuteOperation(b, ops.OpDeployCallProxy, evmChain, contract.DeployInput[ops.OpDeployCallProxyInput]{
				ChainSelector:  in.ChainSelector,
				Qualifier:      in.Qualifier,
				TypeAndVersion: cldf.NewTypeAndVersion(utils.CallProxy, *semver.MustParse("1.0.0")),
				Args: ops.OpDeployCallProxyInput{
					TimelockAddress: common.HexToAddress(timelockAddr),
				},
			})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to deploy call proxy on chain %d: %w", in.ChainSelector, err)
			}
			output.Addresses = append(output.Addresses, callProxyRep.Output)
			callProxyAddr := callProxyRep.Output.Address
			b.Logger.Infof("Deployed Call Proxy at address %s on chain %s", callProxyAddr, evmChain.Name)

			// now that call proxy is deployed, we can add it as executor to the timelock
			_, err = cldf_ops.ExecuteOperation(b, ops.OpGrantRoleTimelock, evmChain, contract.FunctionInput[ops.OpGrantRoleTimelockInput]{
				ChainSelector: in.ChainSelector,
				Address:       common.HexToAddress(timelockAddr),
				Args: ops.OpGrantRoleTimelockInput{
					RoleID:  ops.EXECUTOR_ROLE.ID,
					Account: common.HexToAddress(callProxyAddr),
				},
			})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to grant executor role to call proxy on timelock on chain %d: %w", in.ChainSelector, err)
			}
			b.Logger.Infof("Granted Executor role on Timelock %s to Call Proxy %s on chain %s", timelockAddr, callProxyAddr, evmChain.Name)
			return output, nil
		})
}
