package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployTokenInputConfig struct {
	DeployTokenPerChain map[uint64]DeployTokenInput `yaml:"deploy-token-per-chain" json:"deployTokenPerChain"`
	ChainAdapterVersion *semver.Version             `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	MCMS                mcms.Input                  `yaml:"mcms,omitempty" json:"mcms,omitempty"`
}

type tokenOwnershipInput struct {
	TokensByChain map[uint64][]tokenAdminInput `yaml:"tokens-by-chain" json:"tokensByChain"`
	MCMS          mcms.Input                   `yaml:"mcms,omitempty" json:"mcms,omitempty"`
}

type tokenAdminInput struct {
	TokenAddress string `yaml:"token-address" json:"tokenAddress"`
	AdminAddress string `yaml:"admin-address" json:"adminAddress"`
}

func DeployToken() cldf.ChangeSetV2[DeployTokenInputConfig] {
	return cldf.CreateChangeSet(deployTokenApply(), deployTokenVerify())
}

func deployTokenVerify() func(cldf.Environment, DeployTokenInputConfig) error {
	return func(e cldf.Environment, cfg DeployTokenInputConfig) error {
		tokenPoolRegistry := GetTokenAdapterRegistry()
		for selector, input := range cfg.DeployTokenPerChain {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return fmt.Errorf("not a valid selector: %v", err)
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
			if !exists {
				return fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			// deploy token
			input.ExistingDataStore = e.DataStore
			input.ChainSelector = selector
			err = tokenPoolAdapter.DeployTokenVerify(e, input)
			if err != nil {
				return fmt.Errorf("failed to verify deploy token input for chain selector %d: %w", selector, err)
			}
		}
		return nil
	}
}

func deployTokenApply() func(cldf.Environment, DeployTokenInputConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployTokenInputConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		// ds to collect all addresses created during this changeset
		// this gets passed as output
		ds := datastore.NewMemoryDataStore()
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()
		for selector, input := range cfg.DeployTokenPerChain {
			tmpDatastore := datastore.NewMemoryDataStore()
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			// deploy token
			input.ExistingDataStore = e.DataStore
			input.ChainSelector = selector
			timelockAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(common_utils.RBACTimelock),
				Qualifier:     cfg.MCMS.Qualifier,
			}, input.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("couldn't find the RBACTimelock "+
					"address in datastore for selector %v and qualifier %v %v", input.ChainSelector, cfg.MCMS.Qualifier, err)
			}
			// if token is deployed by CLL, set CCIP admin as RBACTimelock by default.
			// If input has CCIPAdmin and which is external address, set that address as CCIPAdmin
			// and we may not be able to register the token by CLL in that case.
			if input.CCIPAdmin == "" {
				input.CCIPAdmin = timelockAddr.Address
			}
			deployTokenReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, input)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, deployTokenReport.Output.BatchOps...)
			reports = append(reports, deployTokenReport.ExecutionReports...)
			for _, r := range deployTokenReport.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			err = tmpDatastore.Merge(e.DataStore)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
			e.DataStore = tmpDatastore.Seal()
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
