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
	Name     string            `yaml:"name" json:"name"`
	Symbol   string            `yaml:"symbol" json:"symbol"`
	Decimals uint8             `yaml:"decimals" json:"decimals"`
	Supply   *big.Int          `yaml:"supply" json:"supply"`
	// SPLToken, ERC20, etc.
	Type     cldf.ContractType `yaml:"type" json:"type"`
	// not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingAddresses []datastore.AddressRef
}
type DeployTokenPoolInput struct {
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingAddresses []datastore.AddressRef
}
type RegisterTokenInput struct {
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingAddresses []datastore.AddressRef
}
type SetPoolInput struct {
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingAddresses []datastore.AddressRef
}
type UpdateAuthoritiesInput struct {
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	// not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingAddresses []datastore.AddressRef
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
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()
		for selector, input := range cfg.DeployTokenInputs {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, &TokenAdminRegistryVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			input.ExistingAddresses = e.DataStore.Addresses().Filter()
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.DeployToken(), e.BlockChains, input)
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
