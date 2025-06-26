package testutils

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/stretchr/testify/require"
)

func DeriveSendAccounts(
	ctx context.Context,
	t *testing.T,
	authority solana.PrivateKey,
	message ccip_router.SVM2AnyMessage,
	destChainSelector uint64,
	solanaGoClient *rpc.Client,
	router solana.PublicKey,
) (accounts []*solana.AccountMeta, lookUpTables []solana.PublicKey, tokenIndices []uint8) {
	t.Helper()

	derivedAccounts := []*solana.AccountMeta{}
	askWith := []*solana.AccountMeta{}
	stage := "Start"
	tokenIndex := byte(0)
	routerConfigPDA, _, err := state.FindConfigPDA(router)
	require.NoError(t, err)
	derivingTokens := false

	for {
		params := ccip_router.DeriveAccountsCcipSendParams{
			DestChainSelector: destChainSelector,
			CcipSendCaller:    authority.PublicKey(),
			Message:           message,
		}

		deriveRaw := ccip_router.NewDeriveAccountsCcipSendInstruction(
			params,
			stage,
			routerConfigPDA,
		)
		deriveRaw.AccountMetaSlice = append(deriveRaw.AccountMetaSlice, askWith...)
		derive, err := deriveRaw.ValidateAndBuild()
		require.NoError(t, err)
		tx := SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{derive}, authority, config.DefaultCommitment)
		derivation, err := common.ExtractAnchorTypedReturnValue[ccip_router.DeriveAccountsResponse](ctx, tx.Meta.LogMessages, router.String())
		require.NoError(t, err)
		fmt.Printf("Derive stage: %s. Len: %d\n", derivation.CurrentStage, len(derivation.AccountsToSave))
		if derivation.CurrentStage == "TokenTransferStaticAccounts/0" {
			tokenIndices = append(tokenIndices, tokenIndex)
			derivingTokens = true
		}

		if derivingTokens {
			tokenIndex += byte(len(derivation.AccountsToSave))
		}

		for _, meta := range derivation.AccountsToSave {
			derivedAccounts = append(derivedAccounts, &solana.AccountMeta{
				PublicKey:  meta.Pubkey,
				IsWritable: meta.IsWritable,
				IsSigner:   meta.IsSigner,
			})
		}
		askWith = []*solana.AccountMeta{}
		for _, meta := range derivation.AskAgainWith {
			askWith = append(askWith, &solana.AccountMeta{
				PublicKey:  meta.Pubkey,
				IsWritable: meta.IsWritable,
				IsSigner:   meta.IsSigner,
			})
		}
		lookUpTables = append(lookUpTables, derivation.LookUpTablesToSave...)

		if len(derivation.NextStage) == 0 {
			return derivedAccounts, lookUpTables, tokenIndices
		}
		stage = derivation.NextStage
	}
}

func DeriveExecutionAccounts(
	ctx context.Context,
	t *testing.T,
	transmitter solana.PrivateKey,
	messagingAccounts []ccip_offramp.CcipAccountMeta,
	sourceChainSelector uint64,
	tokenTransferAndOffchainData []ccip_offramp.TokenTransferAndOffchainData,
	merkleRoot [32]uint8,
	bufferID []byte,
	tokenReceiver solana.PublicKey,
	solanaGoClient *rpc.Client,
) (accounts []*solana.AccountMeta, lookUpTables []solana.PublicKey, tokenIndices []byte) {
	derivedAccounts := []*solana.AccountMeta{}
	askWith := []*solana.AccountMeta{}
	stage := "Start"
	for {
		params := ccip_offramp.DeriveAccountsExecuteParams{
			ExecuteCaller:       transmitter.PublicKey(),
			MessageAccounts:     messagingAccounts,
			SourceChainSelector: sourceChainSelector,
			TokenTransfers:      tokenTransferAndOffchainData,
			MerkleRoot:          merkleRoot,
			BufferId:            bufferID,
			TokenReceiver:       tokenReceiver,
		}

		deriveRaw := ccip_offramp.NewDeriveAccountsExecuteInstruction(
			params,
			stage,
			config.OfframpConfigPDA,
		)
		deriveRaw.AccountMetaSlice = append(deriveRaw.AccountMetaSlice, askWith...)
		derive, err := deriveRaw.ValidateAndBuild()
		require.NoError(t, err)
		tx := SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{derive}, transmitter, config.DefaultCommitment)
		derivation, err := common.ExtractAnchorTypedReturnValue[ccip_offramp.DeriveAccountsResponse](ctx, tx.Meta.LogMessages, config.CcipOfframpProgram.String())
		require.NoError(t, err)

		if derivation.CurrentStage == "TokenTransferAccounts/Start" {
			// We offset the current index from the capacity of the default meta slice (the fixed accounts)
			tokenIndex := len(derivedAccounts) - cap(ccip_offramp.NewExecuteInstructionBuilder().AccountMetaSlice)
			tokenIndices = append(tokenIndices, byte(tokenIndex))
		}

		for _, meta := range derivation.AccountsToSave {
			derivedAccounts = append(derivedAccounts, &solana.AccountMeta{
				PublicKey:  meta.Pubkey,
				IsWritable: meta.IsWritable,
				IsSigner:   meta.IsSigner,
			})
		}

		askWith = []*solana.AccountMeta{}
		for _, meta := range derivation.AskAgainWith {
			askWith = append(askWith, &solana.AccountMeta{
				PublicKey:  meta.Pubkey,
				IsWritable: meta.IsWritable,
				IsSigner:   meta.IsSigner,
			})
		}

		lookUpTables = append(lookUpTables, derivation.LookUpTablesToSave...)

		if len(derivation.NextStage) == 0 {
			return derivedAccounts, lookUpTables, tokenIndices
		}
		stage = derivation.NextStage
	}
}
