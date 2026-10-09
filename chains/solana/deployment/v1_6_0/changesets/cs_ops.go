package changesets

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	"github.com/smartcontractkit/mcms"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

type SetUpgradeAuthorityConfig struct {
	ChainSelector         uint64
	NewUpgradeAuthority   solana.PublicKey
	SetAfterInitialDeploy bool                              // set all of the programs after the initial deploy
	SetOffRamp            bool                              // offramp not upgraded in place, so may need to set separately
	SetMCMSPrograms       bool                              // these all deploy at once so just set them all
	TransferKeys          []solana.PublicKey                // any keys not covered by the above e.g. partner programs
	MCMS                  *cldfproposalutils.TimelockConfig // if set, assumes current upgrade authority is the timelock
}

func SetUpgradeAuthorityChangeset(
	e cldf.Environment,
	config SetUpgradeAuthorityConfig,
) (cldf.ChangesetOutput, error) {
	chain := e.BlockChains.SolanaChains()[config.ChainSelector]
	state, err := solanastateview.LoadOnchainStateSolana(e)
	if err != nil {
		e.Logger.Errorw("Failed to load existing onchain state", "err", err)
		return cldf.ChangesetOutput{}, err
	}
	chainState, chainExists := state.SolChains[chain.Selector]
	if !chainExists {
		return cldf.ChangesetOutput{}, fmt.Errorf("chain %s not found in existing state, deploy the link token first", chain.String())
	}
	programs := make([]solana.PublicKey, 0)
	if config.SetAfterInitialDeploy {
		programs = append(programs, chainState.Router, chainState.FeeQuoter, chainState.RMNRemote, chainState.BurnMintTokenPools[shared.CLLMetadata], chainState.LockReleaseTokenPools[shared.CLLMetadata])
	}
	if config.SetOffRamp {
		programs = append(programs, chainState.OffRamp)
	}
	if config.SetMCMSPrograms {
		mcmState, err := loadMCMSStateIfDeployed(e, config.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to load onchain state: %w", err)
		}
		programs = append(programs, mcmState.AccessControllerProgram, mcmState.TimelockProgram, mcmState.McmProgram)
	}
	for _, transfer := range config.TransferKeys {
		if transfer.IsZero() {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get program address for chain %s", chain.String())
		}
		programs = append(programs, transfer)
	}
	// We do two loops here just to catch any errors before we get partway through the process
	for _, program := range programs {
		if program.IsZero() {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get program address for chain %s", chain.String())
		}
	}
	currentAuthority := chain.DeployerKey.PublicKey()
	if config.MCMS != nil {
		timelockSignerPDA, err := FetchTimelockSigner(e, chain.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get timelock signer: %w", err)
		}
		currentAuthority = timelockSignerPDA
	}
	e.Logger.Infow("Setting upgrade authority", "newUpgradeAuthority", config.NewUpgradeAuthority.String())
	mcmsTxns := make([]mcmsTypes.Transaction, 0)
	for _, programID := range programs {
		ixn := SetUpgradeAuthority(&e, programID, currentAuthority, config.NewUpgradeAuthority, false)
		if config.MCMS == nil {
			if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to confirm instructions: %w", err)
			}
		} else {
			tx, err := BuildMCMSTxn(
				ixn,
				solana.BPFLoaderUpgradeableProgramID.String(),
				cldf.ContractType(solana.BPFLoaderUpgradeableProgramID.String()),
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to create transaction: %w", err)
			}
			mcmsTxns = append(mcmsTxns, *tx)
		}
	}
	if len(mcmsTxns) > 0 {
		proposal, err := BuildProposalsForTxnsWithConfig(
			e, config.ChainSelector, "proposal to SetUpgradeAuthority in Solana", config.MCMS, mcmsTxns,
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

// SetUpgradeAuthority creates a transaction to set the upgrade authority for a program
func SetUpgradeAuthority(e *cldf.Environment, programID, currentUpgradeAuthority, newUpgradeAuthority solana.PublicKey, isBuffer bool) solana.Instruction {
	e.Logger.Infow("Setting upgrade authority", "programID", programID.String(), "currentUpgradeAuthority", currentUpgradeAuthority.String(), "newUpgradeAuthority", newUpgradeAuthority.String())
	// Buffers use the program account as the program data account
	programDataSlice := solana.NewAccountMeta(programID, true, false)
	if !isBuffer {
		// Actual program accounts use the program data account
		programDataAddress, _, _ := solana.FindProgramAddress([][]byte{programID.Bytes()}, solana.BPFLoaderUpgradeableProgramID)
		programDataSlice = solana.NewAccountMeta(programDataAddress, true, false)
	}

	keys := solana.AccountMetaSlice{
		programDataSlice, // Program account (writable)
		solana.NewAccountMeta(currentUpgradeAuthority, false, true), // Current upgrade authority (signer)
		solana.NewAccountMeta(newUpgradeAuthority, false, false),    // New upgrade authority
	}

	instruction := solana.NewInstruction(
		solana.BPFLoaderUpgradeableProgramID,
		keys,
		// https://github.com/solana-playground/solana-playground/blob/2998d4cf381aa319d26477c5d4e6d15059670a75/vscode/src/commands/deploy/bpf-upgradeable/bpf-upgradeable.ts#L72
		[]byte{4, 0, 0, 0}, // 4-byte SetAuthority instruction identifier
	)

	return instruction
}
