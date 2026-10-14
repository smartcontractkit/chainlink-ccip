package changesets

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/gagliardetto/solana-go"
	"github.com/xssnick/tonutils-go/address"

	"github.com/smartcontractkit/mcms"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	solOffRamp "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/ccip_offramp"
	solRouter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/ccip_router"
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

type AddRemoteChainToRouterConfig struct {
	ChainSelector uint64
	// UpdatesByChain is a mapping of SVM chain selector -> remote chain selector -> remote chain config update
	UpdatesByChain map[uint64]*RouterConfig
	// Disallow mixing MCMS/non-MCMS per chain for simplicity.
	// (can still be achieved by calling this function multiple times)
	MCMS *cldfproposalutils.TimelockConfig
}

type RouterConfig struct {
	// if enabling AllowedSender -> it needs to be a complete list
	// onchain just clones what we pass in
	// and tooling does not handle upserts
	// so you have to clone what is in state, edit the list, and then pass into this changeset
	RouterDestinationConfig solRouter.DestChainConfig
	// inferred from onchain state
	IsUpdate bool
}

func (cfg *AddRemoteChainToRouterConfig) Validate(e cldf.Environment, state solanastateview.CCIPOnChainState) error {
	chain, ok := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if !ok {
		return fmt.Errorf("chain %d not found in environment", cfg.ChainSelector)
	}
	chainState := state.SolChains[cfg.ChainSelector]
	if err := chainState.ValidateRouterConfig(chain); err != nil {
		return err
	}

	if err := ValidateMCMSConfigSolana(e, cfg.MCMS, chain, chainState, solana.PublicKey{}, "", map[cldf.ContractType]bool{shared.Router: true}); err != nil {
		return err
	}
	routerProgramAddress, routerConfigPDA, _ := chainState.GetRouterInfo()
	var routerConfigAccount solRouter.Config
	// already validated that router config exists
	_ = chain.GetAccountDataBorshInto(context.Background(), routerConfigPDA, &routerConfigAccount)

	for remote, remoteConfig := range cfg.UpdatesByChain {
		if !e.BlockChains.Exists(remote) {
			return fmt.Errorf("remote chain %d is not supported", remote)
		}
		if remote == routerConfigAccount.SvmChainSelector {
			return fmt.Errorf("cannot add remote chain %d with same chain selector as current chain %d", remote, cfg.ChainSelector)
		}
		if err := validateRemoteOnRamp(e, remote); err != nil {
			return err
		}
		routerDestChainPDA, err := solState.FindDestChainStatePDA(remote, routerProgramAddress)
		if err != nil {
			return fmt.Errorf("failed to find dest chain state pda for remote chain %d: %w", remote, err)
		}
		var destChainStateAccount solRouter.DestChain
		err = chain.GetAccountDataBorshInto(context.Background(), routerDestChainPDA, &destChainStateAccount)
		if err == nil {
			e.Logger.Infow("remote chain already configured. setting as update", "remoteChainSel", remote)
			remoteConfig.IsUpdate = true
		}
	}
	return nil
}

// Adds new remote chain configurations
func AddRemoteChainToRouter(e cldf.Environment, cfg AddRemoteChainToRouterConfig) (cldf.ChangesetOutput, error) {
	s, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}

	if err := cfg.Validate(e, s); err != nil {
		return cldf.ChangesetOutput{}, err
	}

	ab := cldf.NewMemoryAddressBook()
	ds := datastore.NewMemoryDataStore()
	txns, err := doAddRemoteChainToRouter(e, s, cfg, ab, ds)
	if err != nil {
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
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
			AddressBook:           ab,
			DataStore:             ds,
		}, nil
	}

	return cldf.ChangesetOutput{
		AddressBook: ab,
		DataStore:   ds,
	}, nil
}

func doAddRemoteChainToRouter(
	e cldf.Environment,
	s solanastateview.CCIPOnChainState,
	cfg AddRemoteChainToRouterConfig,
	ab cldf.AddressBook,
	ds datastore.MutableDataStore,
) ([]mcmsTypes.Transaction, error) {
	txns := make([]mcmsTypes.Transaction, 0)
	chainSel := cfg.ChainSelector
	updates := cfg.UpdatesByChain
	chain := e.BlockChains.SolanaChains()[chainSel]
	chainState := s.SolChains[chainSel]
	ccipRouterID, routerConfigPDA, _ := s.SolChains[chainSel].GetRouterInfo()
	offRampID := s.SolChains[chainSel].OffRamp
	routerUsingMCMS := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.Router,
		solana.PublicKey{},
		"",
	)
	lookUpTableEntries := make([]solana.PublicKey, 0)
	authority := GetAuthorityForIxn(
		&e,
		chain,
		chainState,
		shared.Router,
		solana.PublicKey{},
		"",
	)

	for remoteChainSel, update := range updates {
		// verified while loading state
		routerRemoteStatePDA, _ := solState.FindDestChainStatePDA(remoteChainSel, ccipRouterID)
		allowedOffRampRemotePDA, _ := solState.FindAllowedOfframpPDA(remoteChainSel, offRampID, ccipRouterID)

		if update.IsUpdate {
			ix, err := solRouter.NewUpdateDestChainConfigInstruction(
				remoteChainSel,
				// TODO: this needs to be merged with what the user is sending in and whats their onchain.
				// right now, the user will have to send the final version of the config.
				update.RouterDestinationConfig,
				routerRemoteStatePDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate update router config instructions: %w", err)
			}
			routerIxData, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extra data payload from router update dest chain config instruction: %w", err)
			}
			// Manually create instruction rather than directly using the ix above
			// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
			routerIx := solana.NewInstruction(ccipRouterID, ix.Accounts(), routerIxData)
			e.Logger.Infow("update router config for remote chain", "remoteChainSel", remoteChainSel)
			if routerUsingMCMS {
				tx, err := BuildMCMSTxn(routerIx, ccipRouterID.String(), shared.Router)
				if err != nil {
					return txns, fmt.Errorf("failed to create update router config transaction: %w", err)
				}
				txns = append(txns, *tx)
			} else {
				err = chain.Confirm([]solana.Instruction{routerIx})
				if err != nil {
					return txns, fmt.Errorf("failed to confirm update router config instructions: %w", err)
				}
			}
		} else {
			// new remote chain
			lookUpTableEntries = append(lookUpTableEntries,
				routerRemoteStatePDA,
			)
			// generate instructions
			ix, err := solRouter.NewAddChainSelectorInstruction(
				remoteChainSel,
				update.RouterDestinationConfig,
				routerRemoteStatePDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate add router config instructions: %w", err)
			}
			routerIxData, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extra data payload from router update dest chain config instruction: %w", err)
			}
			// Manually create instruction rather than directly using the ix above
			// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
			routerIx := solana.NewInstruction(ccipRouterID, ix.Accounts(), routerIxData)
			e.Logger.Infow("add router config for remote chain", "remoteChainSel", remoteChainSel)
			ix, err = solRouter.NewAddOfframpInstruction(
				remoteChainSel,
				offRampID,
				allowedOffRampRemotePDA,
				routerConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate instructions: %w", err)
			}
			routerOfframpIxData, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extra data payload from router update dest chain config instruction: %w", err)
			}
			// Manually create instruction rather than directly using the ix above
			// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
			routerOfframpIx := solana.NewInstruction(ccipRouterID, ix.Accounts(), routerOfframpIxData)
			e.Logger.Infow("add offramp to router for remote chain", "remoteChainSel", remoteChainSel)
			if routerUsingMCMS {
				// build transactions if mcms
				for _, ix := range []solana.Instruction{routerIx, routerOfframpIx} {
					tx, err := BuildMCMSTxn(ix, ccipRouterID.String(), shared.Router)
					if err != nil {
						return txns, fmt.Errorf("failed to create add router config transaction: %w", err)
					}
					txns = append(txns, *tx)
				}
			} else {
				// confirm ixns if not mcms
				err = chain.Confirm([]solana.Instruction{routerIx, routerOfframpIx})
				if err != nil {
					return txns, fmt.Errorf("failed to confirm add router config instructions: %w", err)
				}
			}
			// add to address book
			tv := cldf.NewTypeAndVersion(shared.RemoteDest, shared.Version1_0_0)
			remoteChainSelStr := strconv.FormatUint(remoteChainSel, 10)
			tv.AddLabel(remoteChainSelStr)
			err = shared.RecordAddress(ab, ds, chainSel, routerRemoteStatePDA.String(), tv, remoteChainSelStr)
			if err != nil {
				return txns, fmt.Errorf("failed to save dest chain state to address book: %w", err)
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

type AddRemoteChainToOffRampConfig struct {
	ChainSelector uint64
	// UpdatesByChain is a mapping of SVM chain selector -> remote chain selector -> remote chain config update
	UpdatesByChain map[uint64]*OffRampConfig
	// Disallow mixing MCMS/non-MCMS per chain for simplicity.
	// (can still be achieved by calling this function multiple times)
	MCMS *cldfproposalutils.TimelockConfig
}

type OffRampConfig struct {
	// source
	EnabledAsSource bool
	// inferred from onchain state
	IsUpdate bool
}

func (cfg *AddRemoteChainToOffRampConfig) Validate(e cldf.Environment, state solanastateview.CCIPOnChainState) error {
	chain, ok := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if !ok {
		return fmt.Errorf("chain %d not found in environment", cfg.ChainSelector)
	}
	chainState := state.SolChains[cfg.ChainSelector]
	if err := chainState.ValidateOffRampConfig(chain); err != nil {
		return err
	}
	if err := ValidateMCMSConfigSolana(e, cfg.MCMS, chain, chainState, solana.PublicKey{}, "", map[cldf.ContractType]bool{shared.OffRamp: true}); err != nil {
		return err
	}

	for remote, remoteConfig := range cfg.UpdatesByChain {
		if !e.BlockChains.Exists(remote) {
			return fmt.Errorf("remote chain %d is not supported", remote)
		}
		if err := validateRemoteOnRamp(e, remote); err != nil {
			return err
		}
		offRampRemoteStatePDA, _, err := solState.FindOfframpSourceChainPDA(remote, chainState.OffRamp)
		if err != nil {
			return fmt.Errorf("failed to find dest chain state pda for remote chain %d: %w", remote, err)
		}
		var destChainStateAccount solOffRamp.SourceChain
		err = chain.GetAccountDataBorshInto(context.Background(), offRampRemoteStatePDA, &destChainStateAccount)
		if err == nil {
			e.Logger.Infow("remote chain already configured. setting as update", "remoteChainSel", remote)
			remoteConfig.IsUpdate = true
		}
	}
	return nil
}

// Adds new remote chain configurations
func AddRemoteChainToOffRamp(e cldf.Environment, cfg AddRemoteChainToOffRampConfig) (cldf.ChangesetOutput, error) {
	s, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}

	if err := cfg.Validate(e, s); err != nil {
		return cldf.ChangesetOutput{}, err
	}

	ab := cldf.NewMemoryAddressBook()
	ds := datastore.NewMemoryDataStore()
	txns, err := doAddRemoteChainToOffRamp(e, s, cfg, ab, ds)
	if err != nil {
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
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
			AddressBook:           ab,
			DataStore:             ds,
		}, nil
	}

	return cldf.ChangesetOutput{
		AddressBook: ab,
		DataStore:   ds,
	}, nil
}

func doAddRemoteChainToOffRamp(
	e cldf.Environment,
	s solanastateview.CCIPOnChainState,
	cfg AddRemoteChainToOffRampConfig,
	ab cldf.AddressBook,
	ds datastore.MutableDataStore,
) ([]mcmsTypes.Transaction, error) {
	txns := make([]mcmsTypes.Transaction, 0)
	chainSel := cfg.ChainSelector
	updates := cfg.UpdatesByChain
	chain := e.BlockChains.SolanaChains()[chainSel]
	chainState := s.SolChains[chainSel]
	offRampID := s.SolChains[chainSel].OffRamp
	offRampUsingMCMS := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.OffRamp,
		solana.PublicKey{},
		"",
	)
	lookUpTableEntries := make([]solana.PublicKey, 0)
	authority := GetAuthorityForIxn(
		&e,
		chain,
		chainState,
		shared.OffRamp,
		solana.PublicKey{},
		"",
	)

	for remoteChainSel, update := range updates {
		// verified while loading state
		offRampRemoteStatePDA, _, _ := solState.FindOfframpSourceChainPDA(remoteChainSel, offRampID)
		// offramp setup
		validSourceChainConfig, err := getSourceChainConfig(e, remoteChainSel, update.EnabledAsSource)
		if err != nil {
			return txns, fmt.Errorf("failed to get source chain config: %w", err)
		}

		var offRampIx solana.Instruction
		if update.IsUpdate {
			ix, err := solOffRamp.NewUpdateSourceChainConfigInstruction(
				remoteChainSel,
				validSourceChainConfig,
				offRampRemoteStatePDA,
				s.SolChains[chainSel].OffRampConfigPDA,
				authority,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate instructions: %w", err)
			}
			data, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extract data payload from offramp update source chain config instruction: %w", err)
			}
			// Manually create instruction rather than directly using the ix above
			// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
			offRampIx = solana.NewInstruction(offRampID, ix.Accounts(), data)
			e.Logger.Infow("update offramp config for remote chain", "remoteChainSel", remoteChainSel)
		} else {
			lookUpTableEntries = append(lookUpTableEntries,
				offRampRemoteStatePDA,
			)
			ix, err := solOffRamp.NewAddSourceChainInstruction(
				remoteChainSel,
				validSourceChainConfig,
				offRampRemoteStatePDA,
				s.SolChains[chainSel].OffRampConfigPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return txns, fmt.Errorf("failed to generate instructions: %w", err)
			}
			data, err := ix.Data()
			if err != nil {
				return txns, fmt.Errorf("failed to extract data payload from offramp add source chain config instruction: %w", err)
			}
			// Manually create instruction rather than directly using the ix above
			// Using the ix above requires setting the program ID in the binding directly which panics if called multiple times
			offRampIx = solana.NewInstruction(offRampID, ix.Accounts(), data)
			e.Logger.Infow("add offramp config for remote chain", "remoteChainSel", remoteChainSel)
			remoteChainSelStr := strconv.FormatUint(remoteChainSel, 10)
			tv := cldf.NewTypeAndVersion(shared.RemoteSource, shared.Version1_0_0)
			tv.AddLabel(remoteChainSelStr)
			err = shared.RecordAddress(ab, ds, chainSel, offRampRemoteStatePDA.String(), tv, remoteChainSelStr)
			if err != nil {
				return txns, fmt.Errorf("failed to save source chain state to address book: %w", err)
			}
		}

		if offRampUsingMCMS {
			tx, err := BuildMCMSTxn(offRampIx, offRampID.String(), shared.OffRamp)
			if err != nil {
				return txns, fmt.Errorf("failed to create transaction: %w", err)
			}
			txns = append(txns, *tx)
		} else {
			err = chain.Confirm([]solana.Instruction{offRampIx})
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

func getSourceChainConfig(e cldf.Environment, remoteChainSel uint64, enabledAsSource bool) (solOffRamp.SourceChainConfig, error) {
	var onRampAddress solOffRamp.OnRampAddress
	addressBytes, err := remoteOnRampAddressBytes(e, remoteChainSel)
	if err != nil {
		return solOffRamp.SourceChainConfig{}, err
	}
	copy(onRampAddress.Bytes[:], addressBytes)
	addressBytesLen := len(addressBytes)
	if addressBytesLen < 0 || addressBytesLen > math.MaxUint32 {
		return solOffRamp.SourceChainConfig{}, fmt.Errorf("address bytes length %d is outside valid uint32 range", addressBytesLen)
	}
	onRampAddress.Len = uint32(addressBytesLen)
	validSourceChainConfig := solOffRamp.SourceChainConfig{
		OnRamp:    onRampAddress,
		IsEnabled: enabledAsSource,
	}
	return validSourceChainConfig, nil
}

// remoteOnRampAddressRefs is the datastore ref each family's onramp address is read from in
// chainlink's stateview.CCIPOnChainState.GetOnRampAddressBytes.
var remoteOnRampAddressRefs = map[string]struct {
	contractType datastore.ContractType
	version      *semver.Version // nil matches any version
}{
	chain_selectors.FamilyEVM:    {datastore.ContractType(shared.OnRamp), semver.MustParse("1.6.0")},
	chain_selectors.FamilySolana: {datastore.ContractType(shared.Router), nil},
	chain_selectors.FamilyAptos:  {datastore.ContractType(shared.AptosCCIPType), nil},
	chain_selectors.FamilySui:    {"SuiOnRamp", nil},
	chain_selectors.FamilyTon:    {datastore.ContractType(shared.OnRamp), nil},
}

// remoteOnRampAddressBytes returns the remote chain's onramp address in the encoding chainlink's
// GetOnRampAddressBytes produced: the 20 EVM address bytes, the 32 bytes of the Solana router, the
// Aptos CCIP package and the Sui onramp package, and the TON onramp as 4 bytes of workchain followed
// by its 32-byte hash. It reads the datastore instead of every family's state, which the Solana
// module can't load. When several active refs match, the highest version wins, as in the state
// loaders that keep the latest program.
//
// The legacy changesets ignored this error and configured an empty onramp. It is returned here.
func remoteOnRampAddressBytes(e cldf.Environment, remote uint64) ([]byte, error) {
	family, err := chain_selectors.GetSelectorFamily(remote)
	if err != nil {
		return nil, err
	}
	want, ok := remoteOnRampAddressRefs[family]
	if !ok {
		return nil, fmt.Errorf("unsupported chain family %s", family)
	}
	if e.DataStore == nil {
		return nil, fmt.Errorf("datastore not available to read the onramp of chain %d", remote)
	}
	filters := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{
		datastore.AddressRefByChainSelector(remote),
		datastore.AddressRefByType(want.contractType),
	}
	if want.version != nil {
		filters = append(filters, datastore.AddressRefByVersion(want.version))
	}
	var latest *datastore.AddressRef
	for _, ref := range e.DataStore.Addresses().Filter(filters...) {
		if ref.Address == "" || ref.Labels.Contains(shared.SupersededLabel) {
			continue
		}
		switch {
		case latest == nil, ref.Version != nil && (latest.Version == nil || ref.Version.GreaterThan(latest.Version)):
			latest = &ref
		case ref.Address != latest.Address && (ref.Version == nil) == (latest.Version == nil) &&
			(ref.Version == nil || ref.Version.Equal(latest.Version)):
			return nil, fmt.Errorf("several active %s refs with the same version on chain %d; the onramp is ambiguous", want.contractType, remote)
		}
	}
	if latest == nil {
		return nil, fmt.Errorf("no onramp found in the datastore for %s chain %d", family, remote)
	}
	switch family {
	case chain_selectors.FamilyEVM:
		return evmAddressBytes(latest.Address)
	case chain_selectors.FamilySolana:
		key, err := solana.PublicKeyFromBase58(latest.Address)
		if err != nil {
			return nil, err
		}
		return key.Bytes(), nil
	case chain_selectors.FamilyAptos, chain_selectors.FamilySui:
		return moveAddressBytes(latest.Address)
	default: // chain_selectors.FamilyTon
		return tonRawAddressBytes(latest.Address)
	}
}

// moveAddressBytes parses an Aptos or Sui hex address, left-padding it to 32 bytes as
// aptos.AccountAddress.ParseStringRelaxed and Sui's NormalizeSuiAddress do.
func moveAddressBytes(address string) ([]byte, error) {
	h := strings.TrimPrefix(strings.TrimPrefix(address, "0x"), "0X")
	if len(h) == 0 || len(h) > 64 {
		return nil, fmt.Errorf("%q is not a 32-byte address", address)
	}
	b, err := hex.DecodeString(strings.Repeat("0", 64-len(h)) + h)
	if err != nil {
		return nil, fmt.Errorf("%q is not a 32-byte address: %w", address, err)
	}
	return b, nil
}

// tonRawAddressBytes parses a TON address as chainlink's TON state loader does and encodes it as
// chainlink-ton's codec.ToRawAddr does: the workchain as a big-endian int32, then the 32-byte
// account data.
func tonRawAddressBytes(addr string) ([]byte, error) {
	parsed, err := address.ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	if parsed.IsAddrNone() || parsed.Type() == address.ExtAddress || len(parsed.Data()) != 32 {
		return nil, fmt.Errorf("%q is not a standard TON address", addr)
	}
	raw := make([]byte, 36)
	binary.BigEndian.PutUint32(raw, uint32(parsed.Workchain())) //nolint:gosec // G115, as in ToRawAddr
	copy(raw[4:], parsed.Data())
	return raw, nil
}
