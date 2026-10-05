package changesets

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/mcms"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	solFeeQuoter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/fee_quoter"
	solCommonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	solState "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"

	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

type AddRemoteChainToFeeQuoterConfig struct {
	ChainSelector uint64
	// UpdatesByChain is a mapping of SVM chain selector -> remote chain selector -> remote chain config update
	UpdatesByChain map[uint64]*FeeQuoterConfig
	// Disallow mixing MCMS/non-MCMS per chain for simplicity.
	// (can still be achieved by calling this function multiple times)
	MCMS *cldfproposalutils.TimelockConfig
}

type FeeQuoterConfig struct {
	FeeQuoterDestinationConfig solFeeQuoter.DestChainConfig
	// inferred from onchain state
	IsUpdate bool
}

func (cfg *AddRemoteChainToFeeQuoterConfig) Validate(e cldf.Environment, state solanastateview.CCIPOnChainState) error {
	chain, ok := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if !ok {
		return fmt.Errorf("chain %d not found in environment", cfg.ChainSelector)
	}
	chainState := state.SolChains[cfg.ChainSelector]
	if err := chainState.ValidateFeeQuoterConfig(chain); err != nil {
		return err
	}
	if err := ValidateMCMSConfigSolana(e, cfg.MCMS, chain, chainState, solana.PublicKey{}, "", map[cldf.ContractType]bool{shared.FeeQuoter: true}); err != nil {
		return err
	}
	for remote, remoteConfig := range cfg.UpdatesByChain {
		if !e.BlockChains.Exists(remote) {
			return fmt.Errorf("remote chain %d is not supported", remote)
		}
		if err := validateRemoteOnRamp(e, remote); err != nil {
			return err
		}
		fqRemoteChainPDA, _, err := solState.FindFqDestChainPDA(remote, chainState.FeeQuoter)
		if err != nil {
			return fmt.Errorf("failed to find dest chain state pda for remote chain %d: %w", remote, err)
		}
		var destChainStateAccount solFeeQuoter.DestChain
		err = chain.GetAccountDataBorshInto(context.Background(), fqRemoteChainPDA, &destChainStateAccount)
		if err == nil {
			e.Logger.Infow("remote chain already configured. setting as update", "remoteChainSel", remote)
			remoteConfig.IsUpdate = true
		}
	}
	return nil
}

// Adds new remote chain configurations
func AddRemoteChainToFeeQuoter(e cldf.Environment, cfg AddRemoteChainToFeeQuoterConfig) (cldf.ChangesetOutput, error) {
	s, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}

	if err := cfg.Validate(e, s); err != nil {
		return cldf.ChangesetOutput{}, err
	}

	ab := cldf.NewMemoryAddressBook()
	ds := datastore.NewMemoryDataStore()
	txns, err := doAddRemoteChainToFeeQuoter(e, s, cfg, ab)
	if err != nil {
		// skipped: doAddRemoteChainToFeeQuoter does not save any lane/multi-instance refs,
		// so the datastore needs no additional qualifier pass.
		return cldf.ChangesetOutput{
			AddressBook: ab,
			DataStore:   ds,
		}, err
	}

	// create proposals for ixns
	if len(txns) > 0 {
		proposal, err := BuildProposalsForTxnsWithConfig(
			e, cfg.ChainSelector, "proposal to add remote chains to Solana", cfg.MCMS, txns,
		)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to build proposal: %w", err)
		}
		// skipped: doAddRemoteChainToFeeQuoter does not save any lane/multi-instance refs,
		// so the datastore needs no additional qualifier pass.
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
			AddressBook:           ab,
			DataStore:             ds,
		}, nil
	}

	// skipped: doAddRemoteChainToFeeQuoter does not save any lane/multi-instance refs,
	// so the datastore needs no additional qualifier pass.
	return cldf.ChangesetOutput{
		AddressBook: ab,
		DataStore:   ds,
	}, nil
}

func doAddRemoteChainToFeeQuoter(
	e cldf.Environment,
	s solanastateview.CCIPOnChainState,
	cfg AddRemoteChainToFeeQuoterConfig,
	ab cldf.AddressBook,
) ([]mcmsTypes.Transaction, error) {
	txns := make([]mcmsTypes.Transaction, 0)
	chainSel := cfg.ChainSelector
	updates := cfg.UpdatesByChain
	chain := e.BlockChains.SolanaChains()[chainSel]
	chainState := s.SolChains[chainSel]
	feeQuoterID := s.SolChains[chainSel].FeeQuoter
	offRampID := s.SolChains[chainSel].OffRamp
	feeQuoterUsingMCMS := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"",
	)
	lookUpTableEntries := make([]solana.PublicKey, 0)
	authority := GetAuthorityForIxn(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"",
	)
	for remoteChainSel, update := range updates {
		// verified while loading state
		fqRemoteChainPDA, _, _ := solState.FindFqDestChainPDA(remoteChainSel, feeQuoterID)
		var feeQuoterIx solana.Instruction
		var err error
		if update.IsUpdate {
			ix, err := solFeeQuoter.NewUpdateDestChainConfigInstruction(
				remoteChainSel,
				// TODO: this needs to be merged with what the user is sending in and whats their onchain.
				// right now, the user will have to send the final version of the config.
				update.FeeQuoterDestinationConfig,
				s.SolChains[chainSel].FeeQuoterConfigPDA,
				fqRemoteChainPDA,
				authority,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate instructions: %w", err)
			}
			ixData, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extract data payload from fee quoter update dest chain config instruction: %w", err)
			}
			feeQuoterIx = solana.NewInstruction(feeQuoterID, ix.Accounts(), ixData)
			e.Logger.Infow("update fee quoter config for remote chain", "remoteChainSel", remoteChainSel)
		} else {
			lookUpTableEntries = append(lookUpTableEntries,
				fqRemoteChainPDA,
			)
			ix, err := solFeeQuoter.NewAddDestChainInstruction(
				remoteChainSel,
				update.FeeQuoterDestinationConfig,
				s.SolChains[chainSel].FeeQuoterConfigPDA,
				fqRemoteChainPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate instructions: %w", err)
			}
			ixData, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extract data payload from fee quoter add dest chain config instruction: %w", err)
			}
			feeQuoterIx = solana.NewInstruction(feeQuoterID, ix.Accounts(), ixData)
			e.Logger.Infow("add fee quoter config for remote chain", "remoteChainSel", remoteChainSel)
		}
		if feeQuoterUsingMCMS {
			tx, err := BuildMCMSTxn(feeQuoterIx, feeQuoterID.String(), shared.FeeQuoter)
			if err != nil {
				return txns, fmt.Errorf("failed to create transaction: %w", err)
			}
			txns = append(txns, *tx)
		} else {
			err = chain.Confirm([]solana.Instruction{feeQuoterIx})
			if err != nil {
				return txns, fmt.Errorf("failed to confirm instructions: %w", err)
			}
		}
	}

	if len(lookUpTableEntries) > 0 {
		err := extendLookupTable(e, chain, offRampID, lookUpTableEntries)
		if err != nil {
			return txns, fmt.Errorf("failed to extend lookup table: %w", err)
		}
	}

	return txns, nil
}

func extendLookupTable(e cldf.Environment, chain cldf_solana.Chain, offRampID solana.PublicKey, lookUpTableEntries []solana.PublicKey) error {
	addressLookupTable, err := solanastateview.FetchOfframpLookupTable(e.GetContext(), chain, offRampID)
	if err != nil {
		return fmt.Errorf("failed to get offramp reference addresses: %w", err)
	}

	addresses, err := solCommonUtil.GetAddressLookupTable(
		e.GetContext(),
		chain.Client,
		addressLookupTable,
	)
	if err != nil {
		return fmt.Errorf("failed to get address lookup table: %w", err)
	}

	// calculate diff and add new entries
	seen := make(map[solana.PublicKey]bool)
	toAdd := make([]solana.PublicKey, 0)
	for _, entry := range addresses {
		seen[entry] = true
	}
	for _, entry := range lookUpTableEntries {
		if _, ok := seen[entry]; !ok {
			toAdd = append(toAdd, entry)
		}
	}
	if len(toAdd) == 0 {
		e.Logger.Infow("no new entries to add to lookup table")
		return nil
	}

	e.Logger.Debugw("Populating lookup table", "keys", toAdd)
	if err := solCommonUtil.ExtendLookupTable(
		e.GetContext(),
		chain.Client,
		addressLookupTable,
		*chain.DeployerKey,
		toAdd,
	); err != nil {
		return fmt.Errorf("failed to extend lookup table: %w", err)
	}
	return nil
}

// remoteOnRampRefs is the datastore ref each family's onramp is loaded from in chainlink's
// stateview.CCIPOnChainState, which ValidateRamp(remote, OnRamp) checks.
var remoteOnRampRefs = map[string]struct {
	contractType datastore.ContractType
	version      *semver.Version // nil matches any version
	name         string
}{
	chain_selectors.FamilyEVM:    {datastore.ContractType(shared.OnRamp), semver.MustParse("1.6.0"), "onramp contract"},
	chain_selectors.FamilySolana: {datastore.ContractType(shared.Router), nil, "router contract"},
	chain_selectors.FamilyAptos:  {datastore.ContractType(shared.AptosCCIPType), nil, "ccip package"},
	chain_selectors.FamilySui:    {"SuiCCIP", nil, "ccip package"},
	chain_selectors.FamilyTon:    {datastore.ContractType(shared.Router), nil, "router contract"},
}

// validateRemoteOnRamp checks that the remote chain has an active onramp ref in the datastore.
// It replaces chainlink's state.ValidateRamp(remote, OnRamp), which needs every family's state;
// the Solana module only loads Solana state.
func validateRemoteOnRamp(e cldf.Environment, remote uint64) error {
	family, err := chain_selectors.GetSelectorFamily(remote)
	if err != nil {
		return err
	}
	want, ok := remoteOnRampRefs[family]
	if !ok {
		return fmt.Errorf("unknown chain family %s", family)
	}
	if e.DataStore == nil {
		return fmt.Errorf("datastore not available to validate remote chain %d", remote)
	}
	filters := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{
		datastore.AddressRefByChainSelector(remote),
		datastore.AddressRefByType(want.contractType),
	}
	if want.version != nil {
		filters = append(filters, datastore.AddressRefByVersion(want.version))
	}
	for _, ref := range e.DataStore.Addresses().Filter(filters...) {
		if ref.Address != "" && !ref.Labels.Contains(shared.SupersededLabel) {
			return nil
		}
	}
	return fmt.Errorf("%s does not exist on %s chain %d", want.name, family, remote)
}
