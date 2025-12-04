package tokenpools

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TODO: Check if this is correct
var TokenAdminRegistryVersion = *semver.MustParse("1.6.0")

// ConfigureTokensForTransfers returns a changeset that configures tokens on multiple chains for transfers with other chains.
func ManualRegistration(tokenRegistry *TokenPoolAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ManualRegistrationInput] {
	return cldf.CreateChangeSet(makeApply(tokenRegistry, mcmsRegistry), makeVerify(tokenRegistry, mcmsRegistry))
}

func makeVerify(_ *TokenPoolAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ManualRegistrationInput) error {
	return func(e cldf.Environment, cfg ManualRegistrationInput) error {
		// TODO: implement
		return nil
	}
}

func makeApply(tokenPoolRegistry *TokenPoolAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ManualRegistrationInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ManualRegistrationInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		family, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenPoolAdapter(family, &TokenAdminRegistryVersion)
		if !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
		}

		manualRegistrationReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tokenPoolAdapter.OnboardTokenPoolForSelfServe(), e.BlockChains, cfg)
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
