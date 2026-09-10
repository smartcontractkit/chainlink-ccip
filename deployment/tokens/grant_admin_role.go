package tokens

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type GrantTokenAdminRoleInput struct {
	ChainAdapterVersion *semver.Version             `yaml:"chainAdapterVersion" json:"chainAdapterVersion"`
	Grants              []GrantTokenAdminRoleConfig `yaml:"grants" json:"grants"`
	MCMS                mcms.Input                  `yaml:"mcms,omitempty" json:"mcms"`
}

type GrantTokenAdminRoleConfig struct {
	// ChainSelector identifies the chain that the token exists on
	ChainSelector uint64 `yaml:"chainSelector" json:"chainSelector"`

	// TokenRef is a reference to the token in the datastore. It is
	// expected that the token supports role-based access control.
	TokenRef datastore.AddressRef `yaml:"tokenRef" json:"tokenRef"`

	// AdminAddress is the address that will be granted the admin role
	// on the token. This field is required.
	AdminAddress string `yaml:"adminAddress" json:"adminAddress"`
}

type GrantTokenAdminRoleSequenceInput struct {
	ChainSelector   uint64
	TokenRef        datastore.AddressRef
	AdminAddress    string
	TimelockAddress string
}

func GrantTokenAdminRole() cldf.ChangeSetV2[GrantTokenAdminRoleInput] {
	return cldf.CreateChangeSet(
		grantTokenAdminRoleApply(GetTokenAdapterRegistry(), changesets.GetRegistry()),
		grantTokenAdminRoleVerify(GetTokenAdapterRegistry()),
	)
}

func grantTokenAdminRoleVerify(tokenRegistry *TokenAdapterRegistry) func(cldf.Environment, GrantTokenAdminRoleInput) error {
	return func(e cldf.Environment, cfg GrantTokenAdminRoleInput) error {
		if len(cfg.Grants) == 0 {
			return errors.New("at least one token admin role grant is required")
		}

		version := cfg.ChainAdapterVersion
		if version == nil {
			return errors.New("chain adapter version is required")
		}

		for i, grant := range cfg.Grants {
			selector := grant.ChainSelector
			if grant.TokenRef.ChainSelector != 0 && grant.TokenRef.ChainSelector != grant.ChainSelector {
				return fmt.Errorf("grant[%d]: chain selector mismatch in TokenRef: expected %d, got %d", i, grant.ChainSelector, grant.TokenRef.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(grant.TokenRef) {
				return fmt.Errorf("grant[%d]: token ref is required", i)
			}
			if grant.AdminAddress == "" {
				return fmt.Errorf("grant[%d]: admin address is required", i)
			}
			if !e.BlockChains.Exists(selector) {
				return fmt.Errorf("grant[%d]: chain selector %d not found in environment", i, selector)
			}

			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return fmt.Errorf("grant[%d]: invalid chain selector %d: %w", i, selector, err)
			}
			adapter, exists := tokenRegistry.GetTokenAdapter(family, version)
			if !exists {
				return fmt.Errorf("grant[%d]: no token adapter registered for chain family '%s' and version '%v'", i, family, version)
			}
			if _, ok := adapter.(TokenAdminRoleAdapter); !ok {
				return fmt.Errorf("grant[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role management", i, family, version)
			}
		}

		return nil
	}
}

func grantTokenAdminRoleApply(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, GrantTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg GrantTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		version := cfg.ChainAdapterVersion
		if version == nil {
			return cldf.ChangesetOutput{}, errors.New("chain adapter version is required")
		}

		for i, grant := range cfg.Grants {
			selector := grant.ChainSelector

			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: invalid chain selector %d: %w", i, selector, err)
			}
			tokenAdapter, ok := tokenRegistry.GetTokenAdapter(family, version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: no token adapter registered for chain family '%s' and version '%v'", i, family, version)
			}
			roleAdapter, ok := tokenAdapter.(TokenAdminRoleAdapter)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role management", i, family, version)
			}
			cleanRef, err := deploy.TryNormalizeAddressRef(selector, grant.TokenRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: failed to normalize token ref %s: %w", i, datastore_utils.SprintRef(grant.TokenRef), err)
			}
			tokenRef, err := datastore_utils.FindAndFormatRef(e.DataStore, cleanRef, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: failed to resolve token ref %s: %w", i, datastore_utils.SprintRef(cleanRef), err)
			}

			var timelockAddress string
			if mcmsReader, ok := mcmsRegistry.GetMCMSReader(family); ok {
				timelockRef, err := mcmsReader.GetTimelockRef(e, selector, cfg.MCMS)
				switch {
				case err != nil:
					e.Logger.Warnf("failed to resolve timelock address for grant[%d] on chain selector %d: %v", i, selector, err)
				case datastore_utils.IsAddressRefEmpty(timelockRef):
					e.Logger.Warnf("timelock ref is empty for grant[%d] on chain selector %d", i, selector)
				default:
					timelockAddress = timelockRef.Address
				}
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				roleAdapter.GrantTokenAdminRole(),
				e.BlockChains, GrantTokenAdminRoleSequenceInput{
					ChainSelector:   grant.ChainSelector,
					AdminAddress:    grant.AdminAddress,
					TimelockAddress: timelockAddress,
					TokenRef:        tokenRef,
				},
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("grant[%d]: failed to grant token admin role: %w", i, err)
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
