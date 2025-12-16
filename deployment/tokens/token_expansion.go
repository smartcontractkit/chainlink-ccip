package tokens

import (
	"fmt"
	"math/big"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type TokenExpansionInput struct {
	DeployTokenInputs       map[uint64]DeployTokenInput       `yaml:"deploy-token-inputs" json:"deployTokenInputs"`
	DeployTokenPoolInputs   map[uint64]DeployTokenPoolInput   `yaml:"deploy-token-pool-inputs" json:"deployTokenPoolInputs"`
	RegisterTokenInputs     map[uint64]RegisterTokenInput     `yaml:"register-token-inputs" json:"registerTokenInputs"`
	SetPoolInputs           map[uint64]SetPoolInput           `yaml:"set-pool-inputs" json:"setPoolInputs"`
	UpdateAuthoritiesInputs map[uint64]UpdateAuthoritiesInput `yaml:"update-authorities-inputs" json:"updateAuthoritiesInputs"`
	MCMS                    mcms.Input                        `yaml:"mcms,omitempty" json:"mcms,omitempty"`
}

type DeployTokenInput struct {
	Name     string   `yaml:"name" json:"name"`
	Symbol   string   `yaml:"symbol" json:"symbol"`
	Decimals uint8    `yaml:"decimals" json:"decimals"`
	Supply   *big.Int `yaml:"supply" json:"supply"`
	// list of addresses who may need special processing in order to send tokens
	// e.g. for Solana, addresses that need associated token accounts created
	Senders []string `yaml:"senders" json:"senders"`
	// SPLToken, ERC20, etc.
	Type cldf.ContractType `yaml:"type" json:"type"`
	// Solana Specific
	// private key in base58 encoding for vanity addresses
	TokenPrivKey string `yaml:"token-priv-key" json:"tokenPrivKey"`
	// if true, the freeze authority will be revoked on token creation
	// and it will be disabled FOREVER
	DisableFreezeAuthority bool `yaml:"disable-freeze-authority" json:"disableFreezeAuthority"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type DeployTokenPoolInput struct {
	RegisterTokenConfig RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type RegisterTokenInput struct {
	RegisterTokenConfig RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type SetPoolInput struct {
	RegisterTokenConfig RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}
type UpdateAuthoritiesInput struct {
	RegisterTokenConfig RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}

func TokenExpansion() cldf.ChangeSetV2[TokenExpansionInput] {
	return cldf.CreateChangeSet(tokenExpansionApply(), tokenExpansionVerify())
}

func tokenExpansionVerify() func(cldf.Environment, TokenExpansionInput) error {
	return func(e cldf.Environment, cfg TokenExpansionInput) error {
		// TODO: implement
		return nil
	}
}

func tokenExpansionApply() func(cldf.Environment, TokenExpansionInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg TokenExpansionInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()
		for selector, input := range cfg.DeployTokenInputs {
			tmpDatastore := datastore.NewMemoryDataStore()
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, &TokenAdminRegistryVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			input.ExistingDataStore = e.DataStore
			input.ChainSelector = selector
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, input)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
			for _, r := range report.Output.Addresses {
				if err := tmpDatastore.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %v on chain with selector %d to datastore: %w", r.Type, r.Version, r, r.ChainSelector, err)
				}
			}
			// merge tmpDatastore into e.DataStore
			tmpDatastore.Merge(e.DataStore)
			e.DataStore = tmpDatastore.Seal()
		}

		for selector, input := range cfg.DeployTokenPoolInputs {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, &TokenAdminRegistryVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			input.ExistingDataStore = e.DataStore
			input.ChainSelector = selector
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployTokenPoolForToken(), e.BlockChains, input)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", selector, err)
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
