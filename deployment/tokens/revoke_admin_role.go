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
			if !e.BlockChains.Exists(revocation.ChainSelector) {
				return fmt.Errorf("revocation[%d]: chain selector %d not found in environment", i, revocation.ChainSelector)
			}

			family, err := chain_selectors.GetSelectorFamily(revocation.ChainSelector)
			if err != nil {
				return fmt.Errorf("revocation[%d]: invalid chain selector %d: %w", i, revocation.ChainSelector, err)
			}
			adapter, exists := tokenRegistry.GetTokenAdapter(family, cciputils.Version_1_0_0)
			if !exists {
				return fmt.Errorf("revocation[%d]: no TokenPoolAdapter registered for chain family '%s' and version '%v'", i, family, cciputils.Version_1_0_0)
			}
			if _, ok := adapter.(TokenAdminRoleAdapter); !ok {
				return fmt.Errorf("revocation[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role revocation", i, family, cciputils.Version_1_0_0)
			}

			if _, err := resolveTokenAdminRoleRef(e, revocation); err != nil {
				return fmt.Errorf("revocation[%d]: failed to resolve token ref: %w", i, err)
			}
		}

		return nil
	}
}

func revokeTokenAdminRoleApply(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, RevokeTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RevokeTokenAdminRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		// Repeated Apply calls must re-read on-chain admin state instead of replaying
		// an earlier successful sequence report with the same input.
		opsBundle := cldf_ops.NewBundle(
			e.GetContext,
			e.Logger,
			cldf_ops.NewMemoryReporter(),
			cldf_ops.WithOperationRegistry(e.OperationsBundle.OperationRegistry),
		)

		for i, revocation := range cfg.Revocations {
			family, err := chain_selectors.GetSelectorFamily(revocation.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: invalid chain selector %d: %w", i, revocation.ChainSelector, err)
			}

			adapter, exists := tokenRegistry.GetTokenAdapter(family, cciputils.Version_1_0_0)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: no TokenPoolAdapter registered for chain family '%s' and version '%v'", i, family, cciputils.Version_1_0_0)
			}
			adminRoleAdapter, ok := adapter.(TokenAdminRoleAdapter)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: token adapter for chain family '%s' and version '%v' does not support token admin role revocation", i, family, cciputils.Version_1_0_0)
			}

			tokenRef, err := resolveTokenAdminRoleRef(e, revocation)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("revocation[%d]: failed to resolve token ref: %w", i, err)
			}

			var timelockAddress string
			if mcmsReader, ok := mcmsRegistry.GetMCMSReader(family); ok {
				timelockRef, err := mcmsReader.GetTimelockRef(e, revocation.ChainSelector, cfg.MCMS)
				if err != nil {
					e.Logger.Warnf("failed to resolve timelock address for revocation[%d] on chain selector %d: %v", i, revocation.ChainSelector, err)
				} else if !datastore_utils.IsAddressRefEmpty(timelockRef) {
					timelockAddress = timelockRef.Address
				}
			} else if revocation.AdminAddress == "" {
				e.Logger.Warnf("no MCMS reader registered for chain family '%s'; revocation[%d] will use the adapter default admin address", family, i)
			}

			adminAddress := revocation.AdminAddress
			if adminAddress == "" {
				adminAddress = timelockAddress
			}

			report, err := cldf_ops.ExecuteSequence(opsBundle, adminRoleAdapter.RevokeTokenAdminRole(), e.BlockChains, RevokeTokenAdminRoleSequenceInput{
				ChainSelector:   revocation.ChainSelector,
				TokenRef:        tokenRef,
				AdminAddress:    adminAddress,
				TimelockAddress: timelockAddress,
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

func resolveTokenAdminRoleRef(e cldf.Environment, revocation RevokeTokenAdminRoleConfig) (datastore.AddressRef, error) {
	if datastore_utils.IsAddressRefEmpty(revocation.TokenRef) {
		return datastore.AddressRef{}, errors.New("token ref is required")
	}
	if revocation.TokenRef.ChainSelector != 0 && revocation.TokenRef.ChainSelector != revocation.ChainSelector {
		return datastore.AddressRef{}, fmt.Errorf("token ref chain selector mismatch: expected %d, got %d", revocation.ChainSelector, revocation.TokenRef.ChainSelector)
	}

	tokenRef, err := TryNormalizeAddressRef(revocation.ChainSelector, revocation.TokenRef)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	tokenRef.ChainSelector = revocation.ChainSelector

	fullRef, err := datastore_utils.FindAndFormatRef(e.DataStore, tokenRef, revocation.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		return fullRef, nil
	}
	if tokenRef.Address != "" && tokenRef.Type != "" {
		return tokenRef, nil
	}

	return datastore.AddressRef{}, fmt.Errorf("token ref must resolve from datastore or include both address and type: %w", err)
}
