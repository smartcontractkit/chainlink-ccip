package examples

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	ccip_receiver "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestCcipReceiver(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	ccip_receiver.SetProgramID(config.CcipBaseReceiver)
	tokenAdmin, _, err := solana.FindProgramAddress([][]byte{[]byte("receiver_token_admin")}, config.CcipBaseReceiver)
	require.NoError(t, err)

	owner, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	receiverState, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, config.CcipBaseReceiver)
	require.NoError(t, err)

	solClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, owner)

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{owner, user}, solClient, t)
		})

		t.Run("initialize receiver", func(t *testing.T) {
			ix, err := ccip_receiver.NewInitializeInstruction(
				config.CcipRouterProgram,
				receiverState,
				tokenAdmin,
				owner.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, owner, rpc.CommitmentConfirmed)
		})
	})

	t.Run("enable source chain + source sender", func(t *testing.T) {
		approvedSenderPDA, err := state.FindApprovedSender(config.EvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
		require.NoError(t, err)
		ixApprove, err := ccip_receiver.NewApproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, owner.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixApprove}, owner, rpc.CommitmentConfirmed)

		t.Run("can disable and reenable", func(t *testing.T) {
			ixUnapprove, err := ccip_receiver.NewUnapproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, owner.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixUnapprove}, owner, rpc.CommitmentConfirmed)

			// ensure PDA closed
			_, err = solClient.GetAccountInfo(ctx, approvedSenderPDA)
			require.ErrorContains(t, err, "not found")

			ixApprove, err := ccip_receiver.NewApproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, owner.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixApprove}, owner, rpc.CommitmentConfirmed)
		})
	})

	t.Run("check ccip_receiver constraints", func(t *testing.T) {
		t.Run("invalid chain + sender", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, err := state.FindApprovedSender(config.SVMChainSelector, []byte{}, config.CcipBaseReceiver)
			require.NoError(t, err)

			ix, err := ccip_receiver.NewCcipReceiveInstruction(ccip_receiver.Any2SVMMessage{SourceChainSelector: config.SVMChainSelector}, user.PublicKey(), approvedSenderPDA, receiverState).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed, []string{"AccountNotInitialized"})
		})
		t.Run("invalid sender", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, err := state.FindApprovedSender(config.EvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
			require.NoError(t, err)
			ix, err := ccip_receiver.NewCcipReceiveInstruction(ccip_receiver.Any2SVMMessage{SourceChainSelector: config.EvmChainSelector, Sender: []byte{1, 2, 3}}, user.PublicKey(), approvedSenderPDA, receiverState).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed, []string{"Address is not router external execution PDA"})
		})
	})

	t.Run("token withdraw", func(t *testing.T) {
		// use token pool for address derivation & state management
		token, err := tokens.NewTokenPool(solana.TokenProgramID)
		require.NoError(t, err)

		ixs, ixErr := tokens.CreateToken(ctx, token.Program, token.Mint.PublicKey(), owner.PublicKey(), 0, solClient, config.DefaultCommitment)
		require.NoError(t, ixErr)

		ixAta, tokenAdminATA, err := tokens.CreateAssociatedTokenAccount(token.Program, token.Mint.PublicKey(), tokenAdmin, owner.PublicKey())
		require.NoError(t, err)
		ixAtaOwner, ownerATA, err := tokens.CreateAssociatedTokenAccount(token.Program, token.Mint.PublicKey(), owner.PublicKey(), owner.PublicKey())
		require.NoError(t, err)

		ixMintTo, mintErr := tokens.MintTo(123, token.Program, token.Mint.PublicKey(), tokenAdminATA, owner.PublicKey())
		require.NoError(t, mintErr)

		testutils.SendAndConfirm(ctx, t, solClient, append(ixs, ixAta, ixAtaOwner, ixMintTo), owner, rpc.CommitmentConfirmed, common.AddSigners(token.Mint))

		// withdraw
		_, initBal, err := tokens.TokenBalance(ctx, solClient, tokenAdminATA, config.DefaultCommitment)
		require.NoError(t, err)
		require.Equal(t, 123, initBal)
		_, initBalOwner, err := tokens.TokenBalance(ctx, solClient, ownerATA, config.DefaultCommitment)
		require.NoError(t, err)
		require.Equal(t, 0, initBalOwner)

		ix, err := ccip_receiver.NewWithdrawTokensInstruction(123, 0, receiverState, tokenAdminATA, ownerATA, token.Mint.PublicKey(), token.Program, tokenAdmin, owner.PublicKey()).ValidateAndBuild()
		testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, owner, rpc.CommitmentConfirmed)

		_, finalBal, err := tokens.TokenBalance(ctx, solClient, tokenAdminATA, config.DefaultCommitment)
		require.NoError(t, err)
		require.Equal(t, 0, finalBal)
		_, finalBalOwner, err := tokens.TokenBalance(ctx, solClient, ownerATA, config.DefaultCommitment)
		require.NoError(t, err)
		require.Equal(t, 123, finalBalOwner)
	})

}
