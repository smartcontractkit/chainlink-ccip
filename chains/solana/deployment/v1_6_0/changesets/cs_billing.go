package changesets

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/smartcontractkit/mcms"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
	solFeeQuoter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/fee_quoter"
	solState "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	solTokenUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

// ADD BILLING TOKEN
type BillingTokenConfig struct {
	ChainSelector uint64
	Config        solFeeQuoter.BillingTokenConfig
	MCMS          *cldfproposalutils.TimelockConfig

	// inferred from state
	IsUpdate bool
}

func (cfg *BillingTokenConfig) Validate(e cldf.Environment, state solanastateview.CCIPOnChainState) error {
	tokenPubKey := cfg.Config.Mint
	chainState := state.SolChains[cfg.ChainSelector]
	if err := chainState.CommonValidation(e, cfg.ChainSelector, tokenPubKey); err != nil {
		return err
	}
	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if err := chainState.ValidateFeeQuoterConfig(chain); err != nil {
		return err
	}
	if _, err := chainState.TokenToTokenProgram(tokenPubKey); err != nil {
		return err
	}
	if err := ValidateMCMSConfigSolana(e, cfg.MCMS, chain, chainState, solana.PublicKey{}, "", map[cldf.ContractType]bool{shared.FeeQuoter: true}); err != nil {
		return err
	}
	// check if already setup
	billingConfigPDA, _, err := solState.FindFqBillingTokenConfigPDA(tokenPubKey, chainState.FeeQuoter)
	if err != nil {
		return fmt.Errorf("failed to find billing token config pda (mint: %s, feeQuoter: %s): %w", tokenPubKey.String(), chainState.FeeQuoter.String(), err)
	}
	var token0ConfigAccount solFeeQuoter.BillingTokenConfigWrapper
	if err := chain.GetAccountDataBorshInto(context.Background(), billingConfigPDA, &token0ConfigAccount); err == nil {
		e.Logger.Infow("Billing token already exists. Configuring as update", "chainSelector", cfg.ChainSelector, "tokenPubKey", tokenPubKey.String())
		cfg.IsUpdate = true
	}
	return nil
}

func AddBillingToken(
	e cldf.Environment,
	chain cldf_solana.Chain,
	chainState solanastateview.CCIPChainState,
	billingTokenConfig solFeeQuoter.BillingTokenConfig,
	mcms *cldfproposalutils.TimelockConfig,
	isUpdate bool,
	feeQuoterAddress solana.PublicKey,
	routerAddress solana.PublicKey,
) ([]mcmsTypes.Transaction, error) {
	txns := make([]mcmsTypes.Transaction, 0)
	tokenPubKey := billingTokenConfig.Mint
	tokenBillingPDA, _, _ := solState.FindFqBillingTokenConfigPDA(tokenPubKey, feeQuoterAddress)
	// we dont need to handle test router here because we explicitly create this and token Receiver for test router
	billingSignerPDA, _, _ := solState.FindFeeBillingSignerPDA(routerAddress)
	tokenProgramID, _ := chainState.TokenToTokenProgram(tokenPubKey)
	tokenReceiver, _, _ := solTokenUtil.FindAssociatedTokenAddress(tokenProgramID, tokenPubKey, billingSignerPDA)
	feeQuoterConfigPDA, _, _ := solState.FindFqConfigPDA(feeQuoterAddress)
	feeQuoterUsingMCMS := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"")

	authority := GetAuthorityForIxn(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"",
	)
	var ixConfig solana.Instruction
	var err error
	if isUpdate {
		ixConfig, err = solFeeQuoter.NewUpdateBillingTokenConfigInstruction(
			billingTokenConfig,
			feeQuoterConfigPDA,
			tokenBillingPDA,
			authority,
		).ValidateAndBuild()
	} else {
		ixConfig, err = solFeeQuoter.NewAddBillingTokenConfigInstruction(
			billingTokenConfig,
			feeQuoterConfigPDA,
			tokenBillingPDA,
			tokenProgramID,
			tokenPubKey,
			tokenReceiver,
			authority, // ccip admin
			billingSignerPDA,
			ata.ProgramID,
			solana.SystemProgramID,
		).ValidateAndBuild()
	}
	if err != nil {
		return txns, fmt.Errorf("failed to generate instructions: %w", err)
	}
	if feeQuoterUsingMCMS {
		tx, err := BuildMCMSTxn(ixConfig, chainState.FeeQuoter.String(), shared.FeeQuoter)
		if err != nil {
			return txns, fmt.Errorf("failed to create transaction: %w", err)
		}
		txns = append(txns, *tx)
	} else {
		if err := chain.Confirm([]solana.Instruction{ixConfig}); err != nil {
			return txns, fmt.Errorf("failed to confirm instructions: %w", err)
		}
	}

	return txns, nil
}

func AddBillingTokenChangeset(e cldf.Environment, cfg BillingTokenConfig) (cldf.ChangesetOutput, error) {
	state, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}
	if err := cfg.Validate(e, state); err != nil {
		return cldf.ChangesetOutput{}, err
	}
	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	chainState := state.SolChains[cfg.ChainSelector]

	runSafely(func() {
		solFeeQuoter.SetProgramID(chainState.FeeQuoter)
	})

	txns, err := AddBillingToken(e, chain, chainState, cfg.Config, cfg.MCMS, cfg.IsUpdate, chainState.FeeQuoter, chainState.Router)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}

	tokenPubKey := cfg.Config.Mint
	tokenBillingPDA, _, _ := solState.FindFqBillingTokenConfigPDA(tokenPubKey, chainState.FeeQuoter)
	if err := extendLookupTable(e, chain, chainState.OffRamp, []solana.PublicKey{tokenBillingPDA}); err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to extend lookup table: %w", err)
	}
	e.Logger.Infow("Billing token added", "chainSelector", cfg.ChainSelector, "tokenPubKey", tokenPubKey.String())

	// create proposals for ixns
	if len(txns) > 0 {
		proposal, err := BuildProposalsForTxnsWithConfig(
			e, cfg.ChainSelector, "proposal to add billing token to Solana", cfg.MCMS, txns)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to build proposal: %w", err)
		}
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
		}, nil
	}

	return cldf.ChangesetOutput{}, nil
}

// ADD BILLING TOKEN FOR REMOTE CHAIN
type TokenTransferFeeForRemoteChainConfig struct {
	ChainSelector      uint64
	RemoteChainConfigs map[uint64]solFeeQuoter.TokenTransferFeeConfig
	TokenPubKey        solana.PublicKey
	MCMS               *cldfproposalutils.TimelockConfig
}

const MinDestBytesOverhead = 32

func (cfg TokenTransferFeeForRemoteChainConfig) Validate(e cldf.Environment, state solanastateview.CCIPOnChainState) error {
	tokenPubKey := cfg.TokenPubKey
	chainState := state.SolChains[cfg.ChainSelector]
	if err := chainState.CommonValidation(e, cfg.ChainSelector, tokenPubKey); err != nil {
		return err
	}
	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if err := chainState.ValidateFeeQuoterConfig(chain); err != nil {
		return fmt.Errorf("fee quoter validation failed: %w", err)
	}
	for _, config := range cfg.RemoteChainConfigs {
		if config.DestBytesOverhead < 32 {
			e.Logger.Infow("dest bytes overhead is less than minimum. Setting to minimum value",
				"destBytesOverhead", config.DestBytesOverhead,
				"minDestBytesOverhead", MinDestBytesOverhead)
			config.DestBytesOverhead = MinDestBytesOverhead
		}
		if config.MinFeeUsdcents > config.MaxFeeUsdcents {
			return fmt.Errorf("min fee %d cannot be greater than max fee %d", config.MinFeeUsdcents, config.MaxFeeUsdcents)
		}
	}

	return ValidateMCMSConfigSolana(e, cfg.MCMS, chain, chainState, solana.PublicKey{}, "", map[cldf.ContractType]bool{shared.FeeQuoter: true})
}

func AddTokenTransferFeeForRemoteChain(e cldf.Environment, cfg TokenTransferFeeForRemoteChainConfig) (cldf.ChangesetOutput, error) {
	state, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}
	if err := cfg.Validate(e, state); err != nil {
		return cldf.ChangesetOutput{}, err
	}

	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	chainState := state.SolChains[cfg.ChainSelector]
	tokenPubKey := cfg.TokenPubKey
	feeQuoterUsingMCMS := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"")

	authority := GetAuthorityForIxn(
		&e,
		chain,
		chainState,
		shared.FeeQuoter,
		solana.PublicKey{},
		"")
	runSafely(func() {
		solFeeQuoter.SetProgramID(chainState.FeeQuoter)
	})
	txns := make([]mcmsTypes.Transaction, 0)
	for remoteChainSelector, config := range cfg.RemoteChainConfigs {
		remoteBillingPDA, _, _ := solState.FindFqPerChainPerTokenConfigPDA(remoteChainSelector, tokenPubKey, chainState.FeeQuoter)

		ix, err := solFeeQuoter.NewSetTokenTransferFeeConfigInstruction(
			remoteChainSelector,
			tokenPubKey,
			config,
			chainState.FeeQuoterConfigPDA,
			remoteBillingPDA,
			authority,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to generate instructions: %w", err)
		}
		if !feeQuoterUsingMCMS {
			if err := chain.Confirm([]solana.Instruction{ix}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to confirm instructions: %w", err)
			}
		}
		if err := extendLookupTable(e, chain, chainState.OffRamp, []solana.PublicKey{remoteBillingPDA}); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to extend lookup table: %w", err)
		}

		e.Logger.Infow("Token billing set for remote chain", "chainSelector ", cfg.ChainSelector, "remoteChainSelector ", remoteChainSelector, "tokenPubKey", tokenPubKey.String())

		if feeQuoterUsingMCMS {
			tx, err := BuildMCMSTxn(ix, chainState.FeeQuoter.String(), shared.FeeQuoter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to create transaction: %w", err)
			}
			txns = append(txns, *tx)
		}
	}

	if len(txns) > 0 {
		proposal, err := BuildProposalsForTxnsWithConfig(
			e, cfg.ChainSelector, "proposal to set billing token for remote chain to Solana", cfg.MCMS, txns)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to build proposal: %w", err)
		}
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
		}, nil
	}

	return cldf.ChangesetOutput{}, nil
}
