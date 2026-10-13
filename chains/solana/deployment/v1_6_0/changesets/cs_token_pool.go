package changesets

import (
	"context"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"

	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/mcms"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	solBaseTokenPool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/cctp_token_pool"
	solTestTokenPool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/test_token_pool"
	solTokenUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

// append mcms txns generated from solanainstructions
func appendTxs(instructions []solana.Instruction, tokenPool solana.PublicKey, poolType cldf.ContractType, txns *[]mcmsTypes.Transaction) error {
	for _, ixn := range instructions {
		tx, err := BuildMCMSTxn(ixn, tokenPool.String(), poolType)
		if err != nil {
			return fmt.Errorf("failed to generate mcms txn: %w", err)
		}
		if tx == nil {
			return errors.New("mcms txn unexpectedly nil")
		}
		*txns = append(*txns, *tx)
	}
	return nil
}

// checks if the evmChainSelector is supported for the given token and pool type
func isSupportedChain(chain cldf_solana.Chain, solTokenPubKey, solPoolAddress solana.PublicKey, poolType cldf.ContractType, evmChainSelector uint64) (bool, solBaseTokenPool.BaseChain, error) {
	// check if this remote chain is already configured for this token
	remoteChainConfigPDA, _, err := solTokenUtil.TokenPoolChainConfigPDA(evmChainSelector, solTokenPubKey, solPoolAddress)
	if err != nil {
		return false, solBaseTokenPool.BaseChain{}, fmt.Errorf("failed to get token pool remote chain config pda (remoteSelector: %d, mint: %s, pool: %s): %w", evmChainSelector, solTokenPubKey.String(), solPoolAddress.String(), err)
	}
	var base solBaseTokenPool.BaseChain
	switch poolType {
	case shared.BurnMintTokenPool, shared.LockReleaseTokenPool:
		var remoteChainConfigAccount solTestTokenPool.ChainConfig
		err = chain.GetAccountDataBorshInto(context.Background(), remoteChainConfigPDA, &remoteChainConfigAccount)
		if err != nil { // not a supported chain for this combination of token and pool type
			return false, solBaseTokenPool.BaseChain{}, nil
		}
		base = remoteChainConfigAccount.Base
	case shared.CCTPTokenPool:
		var remoteChainConfigAccount cctp_token_pool.ChainConfig
		err = chain.GetAccountDataBorshInto(context.Background(), remoteChainConfigPDA, &remoteChainConfigAccount)
		if err != nil { // not a supported chain for this combination of token and pool type
			return false, solBaseTokenPool.BaseChain{}, nil
		}
		base = remoteChainConfigAccount.Base
	}
	return true, base, nil
}

type SyncDomainConfig struct {
	ChainSelector uint64
	// cctpChainConfigMap maps chain selectors to their associated CctpChainConfig
	CCTPChainConfigMap map[uint64]CctpChainConfig
	MCMS               *cldfproposalutils.TimelockConfig
}

type CctpChainConfig struct {
	Domain            uint32
	DestinationCaller solana.PublicKey
}

func (cfg SyncDomainConfig) Validate(e cldf.Environment, chainState solanastateview.CCIPChainState) error {
	// Validate map contains configs
	if len(cfg.CCTPChainConfigMap) == 0 {
		return errors.New("CCTP chain config map is empty")
	}
	// Validate USDC pool exists in state
	if chainState.CCTPTokenPool.IsZero() {
		return errors.New("CCTP token pool does not exist in state")
	}
	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]
	if err := solanastateview.ValidateOwnershipSolana(&e, chain, cfg.MCMS != nil, chainState.CCTPTokenPool, shared.CCTPTokenPool, chainState.USDCToken); err != nil {
		return fmt.Errorf("failed to validate ownership for cctp token pool: %w", err)
	}
	// Validate chain configs are initialized for each chain selector
	for chainSel := range cfg.CCTPChainConfigMap {
		supported, _, err := isSupportedChain(chain, chainState.USDCToken, chainState.CCTPTokenPool, shared.CCTPTokenPool, chainSel)
		if err != nil {
			return fmt.Errorf("failed to validate if remote chain %d is supported: %w", chainSel, err)
		}
		if !supported {
			return fmt.Errorf("chain config not initialized for selector %d", chainSel)
		}
	}
	return nil
}

// SyncDomain adds or removes CCTP domain configs from the Solana CCTP token pool
func SyncDomain(e cldf.Environment, cfg SyncDomainConfig) (cldf.ChangesetOutput, error) {
	e.Logger.Infow("Syncing USDC domains", "cfg", cfg)
	state, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		return cldf.ChangesetOutput{}, err
	}
	chainState := state.SolChains[cfg.ChainSelector]
	if err := cfg.Validate(e, chainState); err != nil {
		return cldf.ChangesetOutput{}, err
	}
	chain := e.BlockChains.SolanaChains()[cfg.ChainSelector]

	cctpTokenPool := chainState.CCTPTokenPool
	usdcToken := chainState.USDCToken

	useMcms := solanastateview.IsSolanaProgramOwnedByTimelock(
		&e,
		chain,
		chainState,
		shared.CCTPTokenPool,
		usdcToken,
		shared.CLLMetadata,
	)
	timelockSignerPDA, err := FetchTimelockSigner(e, cfg.ChainSelector)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to fetch timelock signer: %w", err)
	}
	var authority solana.PublicKey
	if useMcms {
		// If MCMS is used, the authority is the timelock signer PDA
		authority = timelockSignerPDA
	} else {
		// If MCMS is not used, the authority is the deployer key
		authority = chain.DeployerKey.PublicKey()
	}

	statePDA, err := solTokenUtil.TokenPoolConfigAddress(usdcToken, cctpTokenPool)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to calculate token pool config PDA: %w", err)
	}

	var txns []mcmsTypes.Transaction
	runSafely(func() {
		cctp_token_pool.SetProgramID(cctpTokenPool)
	})
	for remoteChainSel, cctpConfig := range cfg.CCTPChainConfigMap {
		e.Logger.Infow("Setting up USDC token pool CCTP config for remote chain", "remote_chain_selector", remoteChainSel, "domain", cctpConfig.Domain, "destination_caller", cctpConfig.DestinationCaller.String())

		chainConfigPDA, _, err := solTokenUtil.TokenPoolChainConfigPDA(remoteChainSel, usdcToken, cctpTokenPool)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to calculate token pool config PDA: %w", err)
		}

		ix, err := cctp_token_pool.NewEditChainRemoteConfigCctpInstruction(
			remoteChainSel,
			usdcToken,
			cctp_token_pool.CctpChain{
				DomainId:          cctpConfig.Domain,
				DestinationCaller: cctpConfig.DestinationCaller,
			},
			statePDA,
			chainConfigPDA,
			authority,
		).ValidateAndBuild()
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to generate instructions: %w", err)
		}
		if useMcms {
			err := appendTxs([]solana.Instruction{ix}, cctpTokenPool, shared.CCTPTokenPool, &txns)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to generate mcms txn: %w", err)
			}
		} else {
			if err := chain.Confirm([]solana.Instruction{ix}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to confirm instructions: %w", err)
			}
		}
	}

	if len(txns) > 0 {
		proposal, err := BuildProposalsForTxnsWithConfig(
			e, cfg.ChainSelector, "proposal to edit USDC token pool CCTP config in Solana", cfg.MCMS, txns,
		)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to build proposal: %w", err)
		}
		return cldf.ChangesetOutput{
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
		}, nil
	}

	return cldf.ChangesetOutput{}, nil
}
