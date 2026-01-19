package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ManualRegistrationInput struct {
	ChainSelector        uint64          `yaml:"chain-selector" json:"chainSelector"`
	ChainAdapterVersion  *semver.Version `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	ExistingAddresses    []datastore.AddressRef
	MCMS                 mcms.Input          `yaml:"mcms,omitempty" json:"mcms"`
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
}

type RegisterTokenConfig struct {
	TokenSymbol        string        `yaml:"token-symbol" json:"tokenSymbol"`
	ProposedOwner      string        `yaml:"proposed-owner" json:"proposedOwner"`
	TokenPoolQualifier string        `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType           string        `yaml:"pool-type" json:"poolType"`
	SVMExtraArgs       *SVMExtraArgs `yaml:"svm-extra-args,omitempty" json:"svmExtraArgs,omitempty"`
}

type SVMExtraArgs struct {
	CustomerMintAuthorities []solana.PublicKey `yaml:"customer-mint-authorities,omitempty" json:"customerMintAuthorities,omitempty"`
	SkipTokenPoolInit       bool               `yaml:"skip-token-pool-init" json:"skipTokenPoolInit"`
}

// ConfigureTokensForTransfers returns a changeset that configures tokens on multiple chains for transfers with other chains.
func ManualRegistration() cldf.ChangeSetV2[ManualRegistrationInput] {
	return cldf.CreateChangeSet(manualRegistrationApply(), manualRegistrationVerify())
}

func manualRegistrationVerify() func(cldf.Environment, ManualRegistrationInput) error {
	return func(e cldf.Environment, cfg ManualRegistrationInput) error {
		// TODO: implement
		return nil
	}
}

func manualRegistrationApply() func(cldf.Environment, ManualRegistrationInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ManualRegistrationInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		family, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
		if !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
		}

		manualRegistrationReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.ManualRegistration(), e.BlockChains, cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to manual register token and token pool %d: %w", cfg.ChainSelector, err)
		}
		batchOps = append(batchOps, manualRegistrationReport.Output.BatchOps...)
		reports = append(reports, manualRegistrationReport.ExecutionReports...)

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
