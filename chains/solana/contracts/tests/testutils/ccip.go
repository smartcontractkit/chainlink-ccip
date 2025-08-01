package testutils

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
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
	var re = regexp.MustCompile(`^TokenTransferStaticAccounts/\d+/0$`)
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
		tx := SimulateTransaction(ctx, t, solanaGoClient, []solana.Instruction{derive}, authority)
		derivation, err := common.ExtractAnchorTypedReturnValue[ccip_router.DeriveAccountsResponse](ctx, tx.Value.Logs, router.String())
		if err != nil {
			fmt.Printf("Error deriving accounts for stage %s: %v\n", stage, err)
			for _, line := range tx.Value.Logs {
				fmt.Println(line)
			}
		}
		require.NoError(t, err)
		fmt.Printf("Derive stage: %s. Len: %d\n", derivation.CurrentStage, len(derivation.AccountsToSave))

		isStartOfToken := re.MatchString(derivation.CurrentStage)
		if isStartOfToken {
			tokenIndices = append(tokenIndices, tokenIndex-byte(cap(ccip_router.NewCcipSendInstructionBuilder().AccountMetaSlice)))
		}
		tokenIndex += byte(len(derivation.AccountsToSave))

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
	offramp solana.PublicKey,
) (accounts []*solana.AccountMeta, lookUpTables []solana.PublicKey, tokenIndices []byte) {
	derivedAccounts := []*solana.AccountMeta{}
	askWith := []*solana.AccountMeta{}
	stage := "Start"
	tokenIndex := byte(0)
	var re = regexp.MustCompile(`^TokenTransferStaticAccounts/\d+/0$`)
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

		offrampConfigPDA, _, err := state.FindConfigPDA(offramp)
		require.NoError(t, err)

		deriveRaw := ccip_offramp.NewDeriveAccountsExecuteInstruction(
			params,
			stage,
			offrampConfigPDA,
		)
		deriveRaw.AccountMetaSlice = append(deriveRaw.AccountMetaSlice, askWith...)
		derive, err := deriveRaw.ValidateAndBuild()
		require.NoError(t, err)

		fmt.Printf("Stage: %s\n", stage)
		tx := SimulateTransaction(ctx, t, solanaGoClient, []solana.Instruction{derive}, transmitter)
		derivation, err := common.ExtractAnchorTypedReturnValue[ccip_offramp.DeriveAccountsResponse](ctx, tx.Value.Logs, offramp.String())
		require.NoError(t, err)

		isStartOfToken := re.MatchString(derivation.CurrentStage)
		if isStartOfToken {
			tokenIndices = append(tokenIndices, tokenIndex-byte(cap(ccip_offramp.NewExecuteInstructionBuilder().AccountMetaSlice)))
		}
		tokenIndex += byte(len(derivation.AccountsToSave))

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
