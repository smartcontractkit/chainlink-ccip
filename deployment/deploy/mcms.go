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
	mcms_types "github.com/smartcontractkit/mcms/types"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type MCMSDeploymentConfig struct {
	Chains         map[uint64]MCMSDeploymentConfigPerChain `json:"chains"`
	AdapterVersion *semver.Version                         `json:"adapterVersion"`
	// specify MCMS in case we need it to actually perform the deployment
	// e.g. for transferring ownership post-deployment
	// e.g. Solana initializing a new MCMS instance owned by an existing MCMS
	MCMS mcms.Input `json:"mcms"`
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

type GrantAdminRoleToTimelockConfigPerChainWithAdminRef struct {
	GrantAdminRoleToTimelockConfigPerChain
	ChainSelector       uint64
	NewAdminTimelockRef datastore.AddressRef
}

type GrantAdminRoleToTimelockConfigPerChain struct {
	TimelockAddress           string `json:"timelockAddress"`           // address of timelock for which we would like to grant the role
	NewAdminTimelockVersion   string `json:"newAdminTimelockVersion"`   // version of timelock to which would like to grant admin role
	NewAdminTimelockQualifier string `json:"newAdminTimelockQualifier"` // qualifier of timelock to which would like to grant admin role
}

type GrantAdminRoleToTimelockConfig struct {
	Chains         map[uint64]GrantAdminRoleToTimelockConfigPerChain `json:"chains"`
	AdapterVersion *semver.Version                                   `json:"adapterVersion"`
}

func GrantAdminRoleToTimelock(deployerReg *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[GrantAdminRoleToTimelockConfig] {
	return cldf.CreateChangeSet(
		grantAdminRoleToTimelockApply(deployerReg, mcmsRegistry),
		grantAdminRoleToTimelockVerify(deployerReg, mcmsRegistry),
	)
}

func grantAdminRoleToTimelockVerify(_ *DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, GrantAdminRoleToTimelockConfig) error {
	return func(e cldf.Environment, cfg GrantAdminRoleToTimelockConfig) error {
		if cfg.AdapterVersion == nil {
			return fmt.Errorf("adapter version is required")
		}

		for _, chainCfg := range cfg.Chains {
			if len(chainCfg.TimelockAddress) == 0 {
				return fmt.Errorf("timelock address cannot be empty")
			}
			if len(chainCfg.NewAdminTimelockVersion) == 0 {
				return fmt.Errorf("new admin timelock version cannot be empty")
			}
			if len(chainCfg.NewAdminTimelockQualifier) == 0 {
				return fmt.Errorf("new admin timelock qualifier cannot be empty")
			}
		}

		return nil
	}
}

func grantAdminRoleToTimelockApply(d *DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, GrantAdminRoleToTimelockConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg GrantAdminRoleToTimelockConfig) (cldf.ChangesetOutput, error) {
		for selector, chainCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, cfg.AdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, cfg.AdapterVersion.String())
			}

			// Find new timelock admin ref
			timelockQualifier := chainCfg.NewAdminTimelockQualifier
			newAdminTimelockRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: selector,
				Type:          datastore.ContractType(utils.RBACTimelock),
				Version:       semver.MustParse(chainCfg.NewAdminTimelockVersion),
				Qualifier:     timelockQualifier,
			}, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to find timelock ref with qualifier %s on chain with selector %d", timelockQualifier, selector)
			}

			// Call the grant role sequence
			seqCfg := GrantAdminRoleToTimelockConfigPerChainWithAdminRef{
				GrantAdminRoleToTimelockConfigPerChain: chainCfg,
				ChainSelector:                          selector,
				NewAdminTimelockRef:                    newAdminTimelockRef,
			}

			_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.GrantAdminRoleToTimelock(), e.BlockChains, seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to Grant Admin Role to Timelock on chain with selector %d: %w", selector, err)
			}
		}

		return cldf.ChangesetOutput{}, nil
	}
}

func DeployMCMS(deployerReg *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig] {
	return cldf.CreateChangeSet(
		deployMCMSApply(deployerReg, mcmsRegistry, false),
		deployMCMSVerify(deployerReg, mcmsRegistry),
	)
}

func FinalizeDeployMCMS(deployerReg *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig] {
	return cldf.CreateChangeSet(
		deployMCMSApply(deployerReg, mcmsRegistry, true),
		deployMCMSVerify(deployerReg, mcmsRegistry))
}

func deployMCMSVerify(_ *DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, MCMSDeploymentConfig) error {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) error {
		// TODO: implement
		if cfg.AdapterVersion == nil {
			return fmt.Errorf("adapter version is required for MCMS deployment verification")
		}
		return nil
	}
}

func deployMCMSApply(
	d *DeployerRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
	finalize bool) func(cldf.Environment, MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MCMSDeploymentConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		batchOps := make([]mcms_types.BatchOperation, 0)
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
			var deployReport cldf_ops.SequenceReport[MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput]
			if finalize {
				deployReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.FinalizeDeployMCMS(), e.BlockChains,
					seqCfg)
			} else {
				deployReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.DeployMCMS(), e.BlockChains,
					seqCfg)
			}
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy MCMS on chain with selector %d: %w", selector, err)
			}

			for _, r := range deployReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			batchOps = append(batchOps, deployReport.Output.BatchOps...)
			reports = append(reports, deployReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
