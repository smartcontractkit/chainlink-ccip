package deploy

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type MCMSDeploymentConfig struct {
	Chains         map[uint64]MCMSDeploymentConfigPerChain `json:"chains"`
	AdapterVersion *semver.Version                         `json:"adapterVersion"`
}

type MCMSDeploymentConfigPerChain struct {
	Canceller        mcmstypes.Config `json:"canceller"`
	Bypasser         mcmstypes.Config `json:"bypasser"`
	Proposer         mcmstypes.Config `json:"proposer"`
	TimelockMinDelay *big.Int         `json:"timelockMinDelay"`
	Label            *string          `json:"label"`
	Qualifier        *string          `json:"qualifier"`
	TimelockAdmin    common.Address   `json:"timelockAdmin"`
}

type MCMSDeploymentConfigPerChainWithAddress struct {
	MCMSDeploymentConfigPerChain
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

func DeployMCMS(deployerReg *DeployerRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig] {
	return cldf.CreateChangeSet(deployMCMSApply(deployerReg), deployMCMSVerify(deployerReg))
}

func deployMCMSVerify(_ *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) error {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) error {
		// TODO: implement
		if cfg.AdapterVersion == nil {
			return fmt.Errorf("adapter version is required for MCMS deployment verification")
		}
		return nil
	}
}

func deployMCMSApply(d *DeployerRegistry) func(cldf.Environment, MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		for selector, mcmsCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, cfg.AdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, cfg.AdapterVersion.String())
			}
			// find existing addresses for this chain from the env
			existingAddrs := d.ExistingAddressesForChain(e, selector)
			// create the sequence input
			seqCfg := MCMSDeploymentConfigPerChainWithAddress{
				MCMSDeploymentConfigPerChain: mcmsCfg,
				ExistingAddresses:            existingAddrs,
				ChainSelector:                selector,
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.DeployMCMS(), e.BlockChains,
				seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy MCMS on chain with selector %d: %w", selector, err)
			}

			for _, r := range deployReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			reports = append(reports, deployReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(mcms.Input{}) // for deployment, we don't need an MCMS proposal
	}
}
