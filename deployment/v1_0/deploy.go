package v1_0

import (
	"fmt"
	"math/big"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type MCMSDeploymentConfig struct {
	Chains map[uint64]MCMSDeploymentConfigPerChain `json:"chains"`
}

type MCMSDeploymentConfigPerChain struct {
	Canceller        mcmstypes.Config `json:"canceller"`
	Bypasser         mcmstypes.Config `json:"bypasser"`
	Proposer         mcmstypes.Config `json:"proposer"`
	TimelockMinDelay *big.Int         `json:"timelockMinDelay"`
	Label            *string          `json:"label"`
	Qualifier        *string          `json:"qualifier"`
}

func DeployMCMS(deployerReg *DeployerRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig] {
	return cldf.CreateChangeSet(deployMCMSApply(deployerReg), deployMCMSVerify(deployerReg))
}

func deployMCMSVerify(_ *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) error {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) error {
		// TODO: implement
		return nil
	}
}

func deployMCMSApply(d *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		for selector, mcmsCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, MCMSVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, MCMSVersion.String())
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.DeployMCMS(), e.BlockChains,
				mcmsCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy MCMS on chain with selector %d: %w", selector, err)
			}

			reports = append(reports, deployReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			Build(mcms.Input{}) // for deployment, we don't need an MCMS proposal
	}
}
