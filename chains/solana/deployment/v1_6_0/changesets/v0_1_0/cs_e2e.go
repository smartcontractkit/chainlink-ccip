package v0_1_0

import (
	"encoding/json"
	"errors"
	"fmt"

	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/changesets"

	proposeutils "github.com/smartcontractkit/cld-changesets/legacy/mcms/proposeutils"
	"github.com/smartcontractkit/mcms"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"
)

type E2ETokenPoolConfig struct {
	AddTokenPoolAndLookupTable        []AddTokenPoolAndLookupTableConfig
	RegisterTokenAdminRegistry        []RegisterTokenAdminRegistryConfig
	AcceptAdminRoleTokenAdminRegistry []AcceptAdminRoleTokenAdminRegistryConfig
	SetPool                           []SetPoolConfig
	RemoteChainTokenPool              []SetupTokenPoolForRemoteChainConfig // setup evm remote pools on solana
	// ConfigureTokenPoolContractsChangesets ran chainlink's EVM v1_5_1.ConfigureTokenPoolContractsChangeset
	// as the last step. That changeset is EVM code, which this module can't import, so the field is
	// kept only so existing inputs (where it is null) still decode; a non-empty value is rejected.
	// Configure the EVM side with the EVM changeset instead.
	ConfigureTokenPoolContractsChangesets []json.RawMessage
	MCMS                                  *cldfproposalutils.TimelockConfig // set it to aggregate all the proposals
}

func E2ETokenPool(e cldf.Environment, cfg E2ETokenPoolConfig) (cldf.ChangesetOutput, error) {
	finalOutput := cldf.ChangesetOutput{}
	finalOutput.AddressBook = cldf.NewMemoryAddressBook()
	addressBookToRemove := cldf.NewMemoryAddressBook()
	defer func(e cldf.Environment) {
		e.Logger.Info("SolanaE2ETokenPool changeset completed")
		e.Logger.Info("Final output: ", finalOutput.AddressBook)
	}(e)
	// if mcms config is not provided, use the mcms config from one of the other configs
	if cfg.MCMS == nil {
		switch {
		case len(cfg.RegisterTokenAdminRegistry) > 0 && cfg.RegisterTokenAdminRegistry[0].MCMS != nil:
			cfg.MCMS = cfg.RegisterTokenAdminRegistry[0].MCMS
		case len(cfg.AcceptAdminRoleTokenAdminRegistry) > 0 && cfg.AcceptAdminRoleTokenAdminRegistry[0].MCMS != nil:
			cfg.MCMS = cfg.AcceptAdminRoleTokenAdminRegistry[0].MCMS
		case len(cfg.SetPool) > 0 && cfg.SetPool[0].MCMS != nil:
			cfg.MCMS = cfg.SetPool[0].MCMS
		}
	}
	if len(cfg.ConfigureTokenPoolContractsChangesets) > 0 {
		return cldf.ChangesetOutput{}, errors.New("ConfigureTokenPoolContractsChangesets is no longer supported: configure the EVM token pools with the EVM changeset")
	}
	err := ProcessConfig(&e, cfg.AddTokenPoolAndLookupTable, AddTokenPoolAndLookupTable, &finalOutput, addressBookToRemove)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to add token pool and lookup table: %w", err)
	}
	err = ProcessConfig(&e, cfg.RemoteChainTokenPool, SetupTokenPoolForRemoteChain, &finalOutput, addressBookToRemove)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure remote chain token pool: %w", err)
	}
	err = ProcessConfig(&e, cfg.RegisterTokenAdminRegistry, RegisterTokenAdminRegistry, &finalOutput, addressBookToRemove)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to register token admin registry: %w", err)
	}
	err = ProcessConfig(&e, cfg.AcceptAdminRoleTokenAdminRegistry, AcceptAdminRoleTokenAdminRegistry, &finalOutput, addressBookToRemove)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to accept admin role: %w", err)
	}
	err = ProcessConfig(&e, cfg.SetPool, SetPool, &finalOutput, addressBookToRemove)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to set pool: %w", err)
	}
	err = AggregateAndCleanup(e, &finalOutput, addressBookToRemove, cfg.MCMS, "E2ETokenPool changeset")
	if err != nil {
		e.Logger.Error("failed to aggregate and cleanup: ", err)
	}

	return finalOutput, nil
}

func ProcessConfig[T any](
	e *cldf.Environment,
	configs []T,
	handler func(cldf.Environment, T) (cldf.ChangesetOutput, error),
	finalOutput *cldf.ChangesetOutput,
	tempRemoveBook cldf.AddressBook,
) error {
	for _, cfg := range configs {
		output, err := handler(*e, cfg)
		if err != nil {
			return err
		}
		err = changesets.MergeChangesetOutput(*e, finalOutput, output)
		if err != nil {
			return fmt.Errorf("failed to merge changeset output: %w", err)
		}

		if ab := output.AddressBook; ab != nil {
			if err := tempRemoveBook.Merge(ab); err != nil {
				return fmt.Errorf("failed to merge into temp: %w", err)
			}
		}
	}
	return nil
}

func AggregateAndCleanup(e cldf.Environment, finalOutput *cldf.ChangesetOutput, abToRemove cldf.AddressBook, cfg *cldfproposalutils.TimelockConfig, proposalDesc string) error {
	allProposals := finalOutput.MCMSTimelockProposals
	if len(allProposals) > 0 {
		state, err := solanastateview.LoadOnchainStateSolana(e)
		if err != nil {
			return fmt.Errorf("failed to load onchain state: %w", err)
		}
		proposal, err := proposeutils.AggregateProposalsV2(
			e, proposeutils.MCMSStates{
				// E2ETokenPool no longer runs EVM changesets, so every proposal it aggregates is
				// on Solana. Chainlink's version also passed the EVM MCMS state of every EVM chain.
				MCMSEVMState:    nil,
				MCMSSolanaState: state.SolanaMCMSStateByChain(e),
			},
			allProposals, proposalDesc, cfg,
		)
		if err != nil {
			return fmt.Errorf("failed to aggregate proposals: %w", err)
		}
		if proposal != nil {
			finalOutput.MCMSTimelockProposals = []mcms.TimelockProposal{*proposal}
		}
	}
	if addresses, err := abToRemove.Addresses(); err == nil && len(addresses) > 0 {
		if err := e.ExistingAddresses.Remove(abToRemove); err != nil {
			return fmt.Errorf("failed to remove temp address book: %w", err)
		}
	}
	return nil
}
