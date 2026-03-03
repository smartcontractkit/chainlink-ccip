package deploytmp

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	mcmstypes "github.com/smartcontractkit/mcms/types"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

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

func UpdateMCMSConfig(deployerReg *DeployerTmpRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateMCMSConfigInput] {
	return cldf.CreateChangeSet(
		updateMCMSConfigApply(deployerReg, mcmsRegistry),
		updateMCMSConfigVerify(deployerReg, mcmsRegistry),
	)
}

func updateMCMSConfigVerify(_ *DeployerTmpRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateMCMSConfigInput) error {
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

func updateMCMSConfigApply(d *DeployerTmpRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateMCMSConfigInput) (cldf.ChangesetOutput, error) {
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
