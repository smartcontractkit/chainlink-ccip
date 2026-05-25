package tokens

import (
	"errors"
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RevokeTokenAdminRoleInput struct {
	Revocations []RevokeTokenAdminRoleConfig `yaml:"revocations" json:"revocations"`
	MCMS        mcms.Input                   `yaml:"mcms,omitempty" json:"mcms"`
}

type RevokeTokenAdminRoleConfig struct {
	ChainSelector uint64               `yaml:"chainSelector" json:"chainSelector"`
	TokenRef      datastore.AddressRef `yaml:"tokenRef" json:"tokenRef"`
	AdminAddress  string               `yaml:"adminAddress,omitempty" json:"adminAddress,omitempty"`
}

type RevokeTokenAdminRoleSequenceInput struct {
	ChainSelector   uint64
	TokenRef        datastore.AddressRef
	AdminAddress    string
	TimelockAddress string
}

func RevokeTokenAdminRole() cldf.ChangeSetV2[RevokeTokenAdminRoleInput] {
	return cldf.CreateChangeSet(
		revokeTokenAdminRoleApply(GetTokenAdapterRegistry(), changesets.GetRegistry()),
		revokeTokenAdminRoleVerify(GetTokenAdapterRegistry()),
	)
}

func revokeTokenAdminRoleVerify(tokenRegistry *TokenAdapterRegistry) func(cldf.Environment, RevokeTokenAdminRoleInput) error {
	return func(e cldf.Environment, cfg RevokeTokenAdminRoleInput) error {
		if len(cfg.Revocations) == 0 {
			return errors.New("at least one token admin role revocation is required")
		}

		for i, revocation := range cfg.Revocations {
			selector := revocation.ChainSelector
			version := cciputils.Version_1_0_0
			if revocation.TokenRef.ChainSelector != 0 && revocation.TokenRef.ChainSelector != revocation.ChainSelector {
				return fmt.Errorf("revocation[%d]: chain selector mismatch in TokenRef: expected %d, got %d", i, revocation.ChainSelector, revocation.TokenRef.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(revocation.TokenRef) {
				return fmt.Errorf("revocation[%d]: token ref is required", i)
			}
			if !e.BlockChains.Exists(selector) {
				return fmt.Errorf("revocation[%d]: chain selector %d not found in environment", i, selector)
			}

			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return fmt.Errorf("revocation[%d]: invalid chain selector %d: %w", i, selector, err)
			}
			adapter, exists := tokenRegistry.GetTokenAdapter(family, version)
			if !exists {
				return fmt.Errorf("revocation[%d]: no TokenPoolAdapter registered for chain family '%s' and version '%v'", i, family, version)
			}
			if _, ok := adapter.(TokenAdminRoleAdapter); !ok {
				return fmt.Errorf("revocation[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role revocation", i, family, version)
			}
		}

		return nil
	}
}

func revokeTokenAdminRoleApply(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, RevokeTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RevokeTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for i, revocation := range cfg.Revocations {
			selector := revocation.ChainSelector
			version := cciputils.Version_1_0_0

			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: invalid chain selector %d: %w", i, selector, err)
			}
			tokenAdapter, ok := tokenRegistry.GetTokenAdapter(family, version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: no TokenPoolAdapter registered for chain family '%s' and version '%v'", i, family, version)
			}
			roleAdapter, ok := tokenAdapter.(TokenAdminRoleAdapter)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role revocation", i, family, version)
			}
			tokenRef, err := ResolveTokenRef(e, tokenRegistry, selector, revocation.TokenRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: failed to resolve token ref: %w", i, err)
			}

			// NOTE: if after resolution, the timelock address is still empty, the adapter should fall back to the deployer key
			var timelockAddress string
			if mcmsReader, ok := mcmsRegistry.GetMCMSReader(family); ok {
				timelockRef, err := mcmsReader.GetTimelockRef(e, selector, cfg.MCMS)
				if err != nil || datastore_utils.IsAddressRefEmpty(timelockRef) {
					e.Logger.Warnf("failed to resolve timelock address for revocation[%d] on chain selector %d: %v", i, selector, err)
				} else {
					timelockAddress = timelockRef.Address
				}
			}

			adminAddress := revocation.AdminAddress
			if adminAddress == "" {
				adminAddress = timelockAddress
			}
			if revocation.AdminAddress == "" {
				e.Logger.Warnf("no MCMS reader registered for chain family '%s'; revocation[%d] will choose a sensible default", family, i)
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, roleAdapter.RevokeTokenAdminRole(), e.BlockChains, RevokeTokenAdminRoleSequenceInput{
				ChainSelector:   revocation.ChainSelector,
				TimelockAddress: timelockAddress,
				AdminAddress:    adminAddress,
				TokenRef:        tokenRef,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: failed to revoke token admin role: %w", i, err)
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
