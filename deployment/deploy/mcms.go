package deploy

import (
	"errors"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
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
	ContractVersion  string           `json:"contractVersion"`
}

type MCMSDeploymentConfigPerChainWithAddress struct {
	MCMSDeploymentConfigPerChain
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

type GrantAdminRoleToTimelockConfigPerChainWithSelector struct {
	GrantAdminRoleToTimelockConfigPerChain
	ChainSelector uint64
}

type GrantAdminRoleToTimelockConfigPerChain struct {
	TimelockToTransferRef datastore.AddressRef `json:"timelockToTransferRef"` // ref of the timelock that transfers its admin rights
	NewAdminTimelockRef   datastore.AddressRef `json:"newAdminTimelockRef"`   // ref of the timelock that will be granted admin
}

type GrantAdminRoleToTimelockConfig struct {
	Chains         map[uint64]GrantAdminRoleToTimelockConfigPerChain `json:"chains"`
	AdapterVersion *semver.Version                                   `json:"adapterVersion"`
}

type UpdateMCMSConfigInputPerChainWithSelector struct {
	UpdateMCMSConfigInputPerChain
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef // needed for Solana
}

type UpdateMCMSConfigInputPerChain struct {
	MCMConfig    mcmstypes.Config
	MCMContracts []datastore.AddressRef
}

type UpdateMCMSConfigInput struct {
	Chains         map[uint64]UpdateMCMSConfigInputPerChain `json:"chains"`
	AdapterVersion *semver.Version                          `json:"adapterVersion"`
	MCMS           mcms.Input                               `json:"mcms"`
}

func UpdateMCMSConfig(deployerReg *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateMCMSConfigInput] {
	return cldf.CreateChangeSet(
		updateMCMSConfigApply(deployerReg, mcmsRegistry),
		updateMCMSConfigVerify(deployerReg, mcmsRegistry),
	)
}

func updateMCMSConfigVerify(_ *DeployerRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateMCMSConfigInput) error {
	return func(e cldf.Environment, cfg UpdateMCMSConfigInput) error {
		if cfg.AdapterVersion == nil {
			return errors.New("adapter version is required")
		}

		// validate mcms input
		if err := cfg.MCMS.Validate(); err != nil {
			return err
		}

		validTypes := []string{utils.BypasserManyChainMultisig.String(), utils.ProposerManyChainMultisig.String(),
			utils.CancellerManyChainMultisig.String()}

		// validate each contract
		for _, chainCfg := range cfg.Chains {
			for _, contract := range chainCfg.MCMContracts {
				if !slices.Contains(validTypes, contract.Type.String()) {
					return errors.New("type of contract needs to be mcms")
				}
				if len(contract.Qualifier) == 0 {
					return errors.New("mcms contract qualifier cannot be empty")
				}
				if len(contract.Version.String()) == 0 {
					return errors.New("mcms contract version cannot be empty")
				}
			}
		}

		return nil
	}
}

func updateMCMSConfigApply(d *DeployerRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateMCMSConfigInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg UpdateMCMSConfigInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcmstypes.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		for selector, chainCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			deployer, exists := d.GetDeployer(family, cfg.AdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no deployer registered for chain family %s and version %s", family, cfg.AdapterVersion.String())
			}

			// If partial refs are provided, resolve to full refs
			mcmsContracts := []datastore.AddressRef{}
			for _, contract := range chainCfg.MCMContracts {
				mcmsQualifier := contract.Qualifier
				mcmsRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
					ChainSelector: selector,
					Type:          contract.Type,
					Version:       contract.Version,
					Qualifier:     mcmsQualifier,
				}, selector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to find mcms ref with qualifier %s on chain with selector %d", mcmsQualifier, selector)
				}

				mcmsContracts = append(mcmsContracts, mcmsRef)
			}

			// find existing addresses for this chain from the env
			existingAddrs := d.ExistingAddressesForChain(e, selector)

			// Call the set mcms config sequence
			seqCfg := UpdateMCMSConfigInputPerChainWithSelector{
				UpdateMCMSConfigInputPerChain: chainCfg,
				ChainSelector:                 selector,
				ExistingAddresses:             existingAddrs,
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployer.UpdateMCMSConfig(), e.BlockChains, seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to Update MCMS Config on chain with selector %d: %w", selector, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
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
			return errors.New("adapter version is required")
		}

		for _, chainCfg := range cfg.Chains {
			// validate timelock to transfer ref
			if chainCfg.TimelockToTransferRef.Type.String() != utils.RBACTimelock.String() {
				return errors.New("type of timelock to transfer must be rbactimelock")
			}
			if len(chainCfg.TimelockToTransferRef.Qualifier) == 0 {
				return errors.New("timelock to transfer qualifier cannot be empty")
			}
			if len(chainCfg.TimelockToTransferRef.Version.String()) == 0 {
				return errors.New("timelock to transfer version cannot be empty")
			}
			// validate new admin timelock ref
			if chainCfg.NewAdminTimelockRef.Type.String() != utils.RBACTimelock.String() {
				return errors.New("type of new admin timelock must be rbactimelock")
			}
			if len(chainCfg.NewAdminTimelockRef.Qualifier) == 0 {
				return errors.New("new admin timelock qualifier cannot be empty")
			}
			if len(chainCfg.NewAdminTimelockRef.Version.String()) == 0 {
				return errors.New("new admin timelock version cannot be empty")
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

			// If partial refs are provided, resolve to full refs
			timelockQualifier := chainCfg.TimelockToTransferRef.Qualifier
			timelockRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: selector,
				Type:          chainCfg.TimelockToTransferRef.Type,
				Version:       chainCfg.TimelockToTransferRef.Version,
				Qualifier:     timelockQualifier,
			}, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to find timelock ref with qualifier %s on chain with selector %d", timelockQualifier, selector)
			}
			chainCfg.TimelockToTransferRef = timelockRef

			newAdminTimelockQualifier := chainCfg.NewAdminTimelockRef.Qualifier
			newAdminTimelockRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: selector,
				Type:          chainCfg.NewAdminTimelockRef.Type,
				Version:       chainCfg.NewAdminTimelockRef.Version,
				Qualifier:     newAdminTimelockQualifier,
			}, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to find timelock ref with qualifier %s on chain with selector %d", newAdminTimelockQualifier, selector)
			}
			chainCfg.NewAdminTimelockRef = newAdminTimelockRef

			// Call the grant role sequence
			seqCfg := GrantAdminRoleToTimelockConfigPerChainWithSelector{
				GrantAdminRoleToTimelockConfigPerChain: chainCfg,
				ChainSelector:                          selector,
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
			return errors.New("adapter version is required for MCMS deployment verification")
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
		batchOps := make([]mcmstypes.BatchOperation, 0)
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
