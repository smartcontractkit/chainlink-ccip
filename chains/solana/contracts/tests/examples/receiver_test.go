package examples

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	bin "github.com/gagliardetto/binary"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_router"
	ccip_receiver "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/example_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestCcipReceiver(t *testing.T) {
	ctx := t.Context()

	ccip_router.SetProgramID(config.CcipRouterProgram)
	ccip_receiver.SetProgramID(config.CcipBaseReceiver)

	// invalid receiver here acts as a "dumb" offramp
	test_ccip_invalid_receiver.SetProgramID(config.CcipInvalidReceiverProgram)
	dumbOfframp := config.CcipInvalidReceiverProgram
	dumbOfframpSignerPDA, _, _ := state.FindExternalExecutionConfigPDA(config.CcipBaseReceiver, dumbOfframp)

	tokenAdmin, _, recAddErr := solana.FindProgramAddress([][]byte{[]byte("receiver_token_admin")}, config.CcipBaseReceiver)
	require.NoError(t, recAddErr)

	ccipAdmin := solana.MustPrivateKeyFromBase58("4D7Hw7YFWqN3jknCRuViYqxF3AKmYosQPnm3szmrR3bvnCPrxKchUCxfFWbQqMCb4oe7jfxynGmjFCTDSrPBdcUB")
	user := solana.MustPrivateKeyFromBase58("5VNkUFwLJ12f71vBMW3XWUfRUpMUnzBxXhPPePi8CzaSXfmQAC842BQtSDkBXR85q4pp6kR7DSiFWBVWGLbFTSoq")
	transmitter := solana.MustPrivateKeyFromBase58("3y3shDibTQ6NGGFDaCWJu6cfFNXje7Qb9uNWsLJqZ7sMANUugsWhLr5daVADhcceFcU2cMXPqL7r6oKr6eqUpQFP")
	invalidOfframp := solana.MustPrivateKeyFromBase58("DJkkQW479LLsWAxAik8kpjKAmd6xRRptqYt7eGRbftoFy3nLJRtCBh42yD2V1kqdg6Q5CWFtN84uS4oit3iAsa3")

	receiverState, _, stateErr := solana.FindProgramAddress([][]byte{[]byte("state")}, config.CcipBaseReceiver)
	require.NoError(t, stateErr)

	solClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, ccipAdmin)

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{ccipAdmin, user, transmitter, invalidOfframp}, solClient, t)
		})

		t.Run("router_setup", func(t *testing.T) {
			// get program data account
			data, accErr := solClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram, &rpc.GetAccountInfoOpts{
				Commitment: rpc.CommitmentConfirmed,
			})
			require.NoError(t, accErr)

			// Decode program data
			var programData struct {
				DataType uint32
				Address  solana.PublicKey
			}
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			feeAggregator := solana.MustPrivateKeyFromBase58("4mKsN4bLEPTQerRRCMALWMFKnkP1xiaC3rYCzcmEmgCu5yrf2eDCPH3jHbsaAg1giKKFwrxk9oUzVxHLYokS1QhN")
			linkTokenMint := solana.MustPrivateKeyFromBase58("2e6af6HmHgxmrv5dLVSqAzerPrLsjEJyyRATvjiBLPpahFv3wdE2NQqaHWjtb8WdVLrvoLchNLoHBr4KVC1GAxBC")

			ix, initErr := ccip_router.NewInitializeInstruction(
				1,
				feeAggregator.PublicKey(),
				config.FeeQuoterProgram,
				linkTokenMint.PublicKey(),
				config.RMNRemoteProgram,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipRouterProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, initErr)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, ccipAdmin, rpc.CommitmentConfirmed)
		})

		t.Run("initialize receiver", func(t *testing.T) {
			ix, initErr := ccip_receiver.NewInitializeInstruction(
				config.CcipRouterProgram,
				receiverState,
				tokenAdmin,
				user.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, initErr)

			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed)
		})

		t.Run("allow offramp in router", func(t *testing.T) {
			allowedOfframpPDA, pdaErr := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbOfframp, config.CcipRouterProgram)
			require.NoError(t, pdaErr)

			ix, offRampErr := ccip_router.NewAddOfframpInstruction(
				config.EvmChainSelector,
				dumbOfframp,
				allowedOfframpPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, offRampErr)

			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, ccipAdmin, rpc.CommitmentConfirmed)
		})
	})

	t.Run("enable source chain + source sender", func(t *testing.T) {
		approvedSenderPDA, findSenderErr := state.FindApprovedSender(config.EvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
		require.NoError(t, findSenderErr)
		ixApprove, appErr := ccip_receiver.NewApproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
		require.NoError(t, appErr)
		testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixApprove}, user, rpc.CommitmentConfirmed)

		t.Run("can disable and reenable", func(t *testing.T) {
			ixUnapprove, unappErr := ccip_receiver.NewUnapproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, unappErr)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixUnapprove}, user, rpc.CommitmentConfirmed)

			// ensure PDA closed
			_, pdaErr := solClient.GetAccountInfo(ctx, approvedSenderPDA)
			require.ErrorContains(t, pdaErr, "not found")

			ixApprove, appSenErr := ccip_receiver.NewApproveSenderInstruction(config.EvmChainSelector, []byte{1, 2, 3}, receiverState, approvedSenderPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, appSenErr)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixApprove}, user, rpc.CommitmentConfirmed)
		})
	})

	t.Run("check ccip_receiver constraints", func(t *testing.T) {
		allowedOfframpPDA, findOfframpErr := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbOfframp, config.CcipRouterProgram)
		require.NoError(t, findOfframpErr)

		t.Run("all valid", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, findSenderErr := state.FindApprovedSender(config.EvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
			require.NoError(t, findSenderErr)

			raw := test_ccip_invalid_receiver.NewReceiverProxyExecuteInstruction(
				test_ccip_invalid_receiver.Any2SVMMessage{SourceChainSelector: config.EvmChainSelector, Sender: []byte{1, 2, 3}},
				config.CcipBaseReceiver,
				dumbOfframpSignerPDA,
				dumbOfframp,
				allowedOfframpPDA,
			)

			raw.Append(solana.Meta(approvedSenderPDA))
			raw.Append(solana.Meta(receiverState))

			ix, validateErr := raw.ValidateAndBuild()
			require.NoError(t, validateErr)
			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, transmitter, rpc.CommitmentConfirmed)
		})

		t.Run("invalid caller (not offramp PDA)", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, findSenderErr := state.FindApprovedSender(config.EvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
			require.NoError(t, findSenderErr)

			dumbOfframpPDA, findOfframpErr := state.FindAllowedOfframpPDA(config.EvmChainSelector, transmitter.PublicKey(), config.CcipRouterProgram)
			require.NoError(t, findOfframpErr)
			testcases := []struct {
				Name              string
				OfframpProgram    solana.PublicKey
				AllowedOfframpPDA solana.PublicKey
			}{
				{"passing transmitter as program", transmitter.PublicKey(), dumbOfframpPDA},
				{"passing actual (dumb) offramp", dumbOfframp, dumbOfframpPDA},
			}

			for _, testcase := range testcases {
				t.Run(testcase.Name, func(t *testing.T) {
					t.Parallel()

					ix, validateErr := ccip_receiver.NewCcipReceiveInstruction( // calling the receiver directly, not through an offramp
						ccip_receiver.Any2SVMMessage{SourceChainSelector: config.EvmChainSelector, Sender: []byte{1, 2, 3}},
						transmitter.PublicKey(), // signing with the transmitter directly, not going through offramp
						testcase.OfframpProgram,
						testcase.AllowedOfframpPDA,
						approvedSenderPDA,
						receiverState,
					).ValidateAndBuild()
					require.NoError(t, validateErr)
					testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, transmitter, rpc.CommitmentConfirmed, []string{"Error Code: " + common.ConstraintSeeds_AnchorError.String()})
				})
			}
		})

		t.Run("invalid offramp for chain", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, findSenderErr := state.FindApprovedSender(config.SvmChainSelector, []byte{1, 2, 3}, config.CcipBaseReceiver)
			require.NoError(t, findSenderErr)

			raw := test_ccip_invalid_receiver.NewReceiverProxyExecuteInstruction(
				// sending from Svm instead of Evm. The offramp is not approved as such for that chain
				test_ccip_invalid_receiver.Any2SVMMessage{SourceChainSelector: config.SvmChainSelector, Sender: []byte{1, 2, 3}},
				config.CcipBaseReceiver,
				dumbOfframpSignerPDA,
				dumbOfframp,
				allowedOfframpPDA,
			)

			raw.Append(solana.Meta(approvedSenderPDA))
			raw.Append(solana.Meta(receiverState))

			ix, validateErr := raw.ValidateAndBuild()

			require.NoError(t, validateErr)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, transmitter, rpc.CommitmentConfirmed, []string{"Error Code: " + common.AccountNotInitialized_AnchorError.String()})
		})

		t.Run("invalid sender", func(t *testing.T) {
			t.Parallel()
			approvedSenderPDA, findSenderErr := state.FindApprovedSender(config.EvmChainSelector, []byte{3, 4, 5}, config.CcipBaseReceiver)
			require.NoError(t, findSenderErr)

			raw := test_ccip_invalid_receiver.NewReceiverProxyExecuteInstruction(
				test_ccip_invalid_receiver.Any2SVMMessage{SourceChainSelector: config.EvmChainSelector, Sender: []byte{3, 4, 5}},
				config.CcipBaseReceiver,
				dumbOfframpSignerPDA,
				dumbOfframp,
				allowedOfframpPDA,
			)

			raw.Append(solana.Meta(approvedSenderPDA))
			raw.Append(solana.Meta(receiverState))

			ix, validateErr := raw.ValidateAndBuild()

			require.NoError(t, validateErr)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, transmitter, rpc.CommitmentConfirmed, []string{"Error Code: " + common.AccountNotInitialized_AnchorError.String()})
		})
	})

	t.Run("token withdraw", func(t *testing.T) {
		// use token pool for address derivation & state management
		mint := solana.MustPrivateKeyFromBase58("4dD1x6rv1uLHKWCrYBY9WYa781YgNQGocVpqrS1EzfDQAq9TK4Vdyju6eLXicoSmjiGU9uZ9ExJHmC5GzwGoQUWD")
		token, tpErr := tokens.NewTokenPool(solana.TokenProgramID, config.CcipTokenPoolProgram, mint.PublicKey())
		require.NoError(t, tpErr)

		ixs, ixErr := tokens.CreateToken(ctx, token.Program, token.Mint, ccipAdmin.PublicKey(), 0, solClient, rpc.CommitmentConfirmed)
		require.NoError(t, ixErr)

		ixAta, tokenAdminATA, ataErr := tokens.CreateAssociatedTokenAccount(token.Program, token.Mint, tokenAdmin, ccipAdmin.PublicKey())
		require.NoError(t, ataErr)
		ixAtaOwner, ccipAdminATA, ataOwnerErr := tokens.CreateAssociatedTokenAccount(token.Program, token.Mint, ccipAdmin.PublicKey(), ccipAdmin.PublicKey())
		require.NoError(t, ataOwnerErr)

		ixMintTo, mintErr := tokens.MintTo(123, token.Program, token.Mint, tokenAdminATA, ccipAdmin.PublicKey())
		require.NoError(t, mintErr)

		testutils.SendAndConfirm(ctx, t, solClient, append(ixs, ixAta, ixAtaOwner, ixMintTo), ccipAdmin, rpc.CommitmentConfirmed, common.AddSigners(mint))

		// withdraw
		_, initBal, balErr := tokens.TokenBalance(ctx, solClient, tokenAdminATA, rpc.CommitmentConfirmed)
		require.NoError(t, balErr)
		require.Equal(t, 123, initBal)
		_, initBalOwner, balOwnerErr := tokens.TokenBalance(ctx, solClient, ccipAdminATA, rpc.CommitmentConfirmed)
		require.NoError(t, balOwnerErr)
		require.Equal(t, 0, initBalOwner)

		ix, withdrawErr := ccip_receiver.NewWithdrawTokensInstruction(123, 0, receiverState, tokenAdminATA, ccipAdminATA, token.Mint, token.Program, tokenAdmin, user.PublicKey()).ValidateAndBuild()
		require.NoError(t, withdrawErr)
		testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed)

		_, finalBal, balErr := tokens.TokenBalance(ctx, solClient, tokenAdminATA, rpc.CommitmentConfirmed)
		require.NoError(t, balErr)
		require.Equal(t, 0, finalBal)
		_, finalBalOwner, balOwnerErr := tokens.TokenBalance(ctx, solClient, ccipAdminATA, rpc.CommitmentConfirmed)
		require.NoError(t, balOwnerErr)
		require.Equal(t, 123, finalBalOwner)
	})
}
