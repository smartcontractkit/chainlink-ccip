package contracts

import (
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	// "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/redirecting_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	// "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestSetupRedirectingReceiver(t *testing.T) {
	devnetInfo, err := getDevnetInfo()
	require.NoError(t, err)

	ctx := tests.Context(t)
	client := rpc.New(devnetInfo.RPC)

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())

	offrampAddress, err := solana.PublicKeyFromBase58(devnetInfo.Offramp)
	require.NoError(t, err)

	offrampPDAs, err := getOfframpPDAs(offrampAddress)
	require.NoError(t, err)

	tokenMintForRedirect, err := solana.PublicKeyFromBase58(devnetInfo.TokenMintForRedirectTest)
	require.NoError(t, err)

	finalReceivingUser, err := solana.PublicKeyFromBase58(devnetInfo.FinalReceiverForRedirect)
	require.NoError(t, err)

	var referenceAddresses ccip_offramp.ReferenceAddresses

	t.Run("Read Reference Addresses", func(t *testing.T) {
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
		fmt.Printf("Reference Addresses: %+v\n", referenceAddresses)
	})

	redirectingReceiverAddress := solana.MustPublicKeyFromBase58(devnetInfo.RedirectingReceiver)

	redirecting_ccip_receiver.SetProgramID(redirectingReceiverAddress)

	statePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, redirectingReceiverAddress)
	require.NoError(t, err)
	tokenAdminPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("receiver_token_admin")}, redirectingReceiverAddress)
	require.NoError(t, err)

	t.Run("Initialize redirecting receiver", func(t *testing.T) {
		t.Skip()
		ix, err := redirecting_ccip_receiver.NewInitializeInstruction(
			referenceAddresses.Router,
			statePDA,
			tokenAdminPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Initialize receiver ATA for the current test token", func(t *testing.T) {
		t.Skip()
		ix, _, _ := tokens.CreateAssociatedTokenAccount(solana.Token2022ProgramID, tokenMintForRedirect, tokenAdminPDA, admin.PublicKey())
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
		fmt.Printf("Result: %s\n", result.Meta.LogMessages)
	})

	t.Run("Print required extra accounts", func(t *testing.T) {
		fmt.Printf("Account 4 (state): %s\n", statePDA)
		programATA, _, _ := tokens.FindAssociatedTokenAddress(solana.Token2022ProgramID, tokenMintForRedirect, tokenAdminPDA)
		fmt.Printf("Account 5 [WRITABLE] (program token account): %s\n", programATA)
		receiverATA, _, _ := tokens.FindAssociatedTokenAddress(solana.Token2022ProgramID, tokenMintForRedirect, finalReceivingUser)
		fmt.Printf("Account 6 [WRITABLE](destination token account): %s\n", receiverATA)
		fmt.Printf("Account 7 (token mint): %s\n", tokenMintForRedirect)
		fmt.Printf("Account 8 (token program): %s\n", solana.Token2022ProgramID)
		fmt.Printf("Account 9 (token admin): %s\n", tokenAdminPDA)
		fmt.Printf("\nIn command argument format:\n\n")
		fmt.Printf(" run . send-message --environment=\"solana-staging\" \\\n")
		fmt.Printf("--source \"$ETHEREUM_TESTNET_SEPOLIA_SELECTOR\" \\\n")
		fmt.Printf("--dest \"$SOLANA_DEVNET_SELECTOR\" \\\n")
		fmt.Printf("--receiver \"%s\" \\\n", redirectingReceiverAddress)
		fmt.Printf("--message \"%s\" \\\n", finalReceivingUser)
		fmt.Printf("--sol-accounts \"%s\" \\\n", statePDA)
		fmt.Printf("--sol-accounts \"%s\" \\\n", programATA)
		fmt.Printf("--sol-accounts \"%s\" \\\n", receiverATA)
		fmt.Printf("--sol-accounts \"%s\" \\\n", tokenMintForRedirect)
		fmt.Printf("--sol-accounts \"%s\" \\\n", solana.Token2022ProgramID)
		fmt.Printf("--sol-accounts \"%s\" \\\n", tokenAdminPDA)
		fmt.Printf("--sol-account-is-writable-bitmap 1 \\\n")
		fmt.Printf("--sol-account-is-writable-bitmap 2 \\\n")
		fmt.Printf("--sol-account-is-writable-bitmap 3 \\\n")
		// Note: The admin PDA, not the ATA, as offchain will derive it.
		fmt.Printf("--sol-token-receiver \"%s\" \\\n", tokenAdminPDA)
		fmt.Printf("--token-address \"<SEPOLIA_MINT_FOR_THIS_TOKEN>\"")
		fmt.Printf("--token-amount \"<AMOUNT>\"")

	})

	t.Log()
}
