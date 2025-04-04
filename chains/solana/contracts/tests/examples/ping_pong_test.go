package examples

import (
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	bin "github.com/gagliardetto/binary"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ping_pong_demo"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

type ReusableAccounts struct {
	// Ping Pong accounts
	Config         solana.PublicKey
	NameVersion    solana.PublicKey
	FeeTokenATA    solana.PublicKey
	CcipSendSigner solana.PublicKey

	// Link-related accounts in other CCIP programs
	fqBillingConfig       solana.PublicKey
	routerBillingReceiver solana.PublicKey

	// "Dumb" offramp accounts
	offramp        solana.PublicKey
	allowedOfframp solana.PublicKey
	rampSigner     solana.PublicKey
	nonce          solana.PublicKey
}

func getReusableAccounts(t *testing.T, linkMint solana.PublicKey) ReusableAccounts {
	ppConfig, _, err := solana.FindProgramAddress([][]byte{[]byte("config")}, config.PingPongProgram)
	require.NoError(t, err)
	ppNameVersion, _, _ := solana.FindProgramAddress([][]byte{[]byte("name_version")}, config.PingPongProgram)
	require.NoError(t, err)
	ppCcipSendSigner, _, _ := solana.FindProgramAddress([][]byte{[]byte("ccip_send_signer")}, config.PingPongProgram)
	require.NoError(t, err)
	ppLinkAta, _, err := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, ppCcipSendSigner)
	require.NoError(t, err)

	fqBillingConfig, _, err := state.FindFqBillingTokenConfigPDA(linkMint, config.FeeQuoterProgram)
	require.NoError(t, err)
	routerBillingReceiver, _, err := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, config.BillingSignerPDA)
	require.NoError(t, err)

	dumbRamp := config.CcipInvalidReceiverProgram
	allowedOfframp, err := state.FindAllowedOfframpPDA(config.SvmChainSelector, dumbRamp, config.CcipRouterProgram)
	require.NoError(t, err)
	rampSigner, _, err := state.FindExternalExecutionConfigPDA(config.PingPongProgram, dumbRamp)
	require.NoError(t, err)
	nonce, err := state.FindNoncePDA(config.SvmChainSelector, ppCcipSendSigner, config.CcipRouterProgram)
	require.NoError(t, err)

	return ReusableAccounts{
		FeeTokenATA:           ppLinkAta,
		Config:                ppConfig,
		NameVersion:           ppNameVersion,
		CcipSendSigner:        ppCcipSendSigner,
		fqBillingConfig:       fqBillingConfig,
		routerBillingReceiver: routerBillingReceiver,
		offramp:               dumbRamp,
		allowedOfframp:        allowedOfframp,
		rampSigner:            rampSigner,
		nonce:                 nonce,
	}
}

// Test basic happy path of the ping pong demo with "itself", meaning an SVM <-> SVM message
// from the ping pong program to itself.
func TestPingPong(t *testing.T) {
	t.Parallel()

	// acting as "dumb" offramp, proxying calls to the receiver that are signed by PDA
	test_ccip_invalid_receiver.SetProgramID(config.CcipInvalidReceiverProgram)
	fee_quoter.SetProgramID(config.FeeQuoterProgram)
	ccip_router.SetProgramID(config.CcipRouterProgram)
	rmn_remote.SetProgramID(config.RMNRemoteProgram)
	ping_pong_demo.SetProgramID(config.PingPongProgram)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	ctx := tests.Context(t)

	linkMintPrivK := solana.MustPrivateKeyFromBase58("32YVeJArcWWWV96fztfkRQhohyFz5Hwno93AeGVrN4g2LuFyvwznrNd9A6tbvaTU6BuyBsynwJEMLre8vSy3CrVU")
	linkMint := linkMintPrivK.PublicKey()

	reusableAccounts := getReusableAccounts(t, linkMint)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{admin}, solanaGoClient, t)
		})

		t.Run("Create LINK token", func(t *testing.T) {
			ixToken, terr := tokens.CreateToken(
				ctx,
				config.Token2022Program,
				linkMint,
				admin.PublicKey(),
				9,
				solanaGoClient,
				config.DefaultCommitment,
			)
			require.NoError(t, terr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ixToken, admin, config.DefaultCommitment,
				common.AddSigners(linkMintPrivK))
		})

		t.Run("Initializations", func(t *testing.T) {
			type ProgramData struct {
				DataType uint32
				Address  solana.PublicKey
			}

			t.Run("Fee Quoter", func(t *testing.T) {
				t.Parallel()

				defaultMaxFeeJuelsPerMsg := bin.Uint128{Lo: 300000000, Hi: 0, Endianness: nil}

				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.FeeQuoterProgram,
					&rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)

				// Decode program data
				var programData ProgramData
				require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

				ix, err := fee_quoter.NewInitializeInstruction(
					defaultMaxFeeJuelsPerMsg,
					config.CcipRouterProgram,
					config.FqConfigPDA,
					linkMint,
					admin.PublicKey(),
					solana.SystemProgramID,
					config.FeeQuoterProgram,
					programData.Address,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
					config.DefaultCommitment)
			})

			t.Run("Router", func(t *testing.T) {
				t.Parallel()

				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram,
					&rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)

				// Decode program data
				var programData ProgramData
				require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

				instruction, err := ccip_router.NewInitializeInstruction(
					config.SvmChainSelector,
					admin.PublicKey(),
					config.FeeQuoterProgram,
					linkMint,
					config.RMNRemoteProgram,
					config.RouterConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
					config.CcipRouterProgram,
					programData.Address,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin,
					config.DefaultCommitment)
			})

			t.Run("RMN", func(t *testing.T) {
				t.Parallel()

				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(
					ctx,
					config.RMNRemoteProgram,
					&rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment},
				)
				require.NoError(t, err)

				// Decode program data
				var programData ProgramData
				require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

				// Now, actually initialize the offramp
				initIx, err := rmn_remote.NewInitializeInstruction(
					config.RMNRemoteConfigPDA,
					config.RMNRemoteCursesPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
					config.RMNRemoteProgram,
					programData.Address,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initIx}, admin,
					config.DefaultCommitment)
			})
		})

		t.Run("Other CCIP configs", func(t *testing.T) {
			t.Run("register dumb offramp", func(t *testing.T) {
				t.Parallel()

				routerIx, err := ccip_router.NewAddOfframpInstruction(
					config.SvmChainSelector,
					reusableAccounts.offramp,
					reusableAccounts.allowedOfframp,
					config.RouterConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{routerIx}, admin,
					config.DefaultCommitment)
			})

			t.Run("Register link as billing token", func(t *testing.T) {
				t.Parallel()

				ix, cerr := fee_quoter.NewAddBillingTokenConfigInstruction(
					fee_quoter.BillingTokenConfig{
						Enabled: true,
						Mint:    linkMint,
						UsdPerToken: fee_quoter.TimestampedPackedU224{
							Value:     [28]uint8{1},
							Timestamp: 123,
						},
						PremiumMultiplierWeiPerEth: 0,
					},
					config.FqConfigPDA,
					reusableAccounts.fqBillingConfig,
					config.Token2022Program,
					linkMint,
					reusableAccounts.routerBillingReceiver,
					admin.PublicKey(),
					config.BillingSignerPDA,
					tokens.AssociatedTokenProgramID,
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, cerr)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
					config.DefaultCommitment)
			})

			t.Run("Register SVM as dest chain", func(t *testing.T) {
				t.Parallel()

				routerIx, err := ccip_router.NewAddChainSelectorInstruction(
					config.SvmChainSelector,
					ccip_router.DestChainConfig{},
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				fqIx, err := fee_quoter.NewAddDestChainInstruction(
					config.SvmChainSelector,
					fee_quoter.DestChainConfig{
						IsEnabled: true,

						LaneCodeVersion: fee_quoter.Default_CodeVersion,

						// minimal valid config
						DefaultTxGasLimit:           200000,
						MaxPerMsgGasLimit:           3000000,
						MaxDataBytes:                30000,
						MaxNumberOfTokensPerMsg:     5,
						DefaultTokenDestGasOverhead: 50000,
						ChainFamilySelector:         [4]uint8(config.SvmChainFamilySelector),

						DefaultTokenFeeUsdcents: 50,
						NetworkFeeUsdcents:      50,
					},
					config.FqConfigPDA,
					config.FqSvmDestChainPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{routerIx, fqIx}, admin,
					config.DefaultCommitment)
			})
		})
	})

	t.Run("Configure PingPong", func(t *testing.T) {
		t.Run("Initialize Config", func(t *testing.T) {
			extraArgs := testutils.MustSerializeExtraArgs(
				t,
				fee_quoter.SVMExtraArgsV1{
					ComputeUnits:             400_000,
					AccountIsWritableBitmap:  ccip.GenerateBitMapForIndexes([]int{1, 4, 7, 8, 9}),
					AllowOutOfOrderExecution: true,
					TokenReceiver:            [32]uint8{}, // none, no token transfer
					Accounts: [][32]uint8{
						reusableAccounts.Config,
						reusableAccounts.CcipSendSigner, // 1
						config.Token2022Program,
						linkMint,
						reusableAccounts.FeeTokenATA, // 4
						config.CcipRouterProgram,
						config.RouterConfigPDA,
						config.SvmDestChainStatePDA,            // 7
						reusableAccounts.nonce,                 // 8
						reusableAccounts.routerBillingReceiver, // 9
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
						config.FqSvmDestChainPDA,
						reusableAccounts.fqBillingConfig,
						reusableAccounts.fqBillingConfig,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
						solana.SystemProgramID,
					},
				},
				ccip.SVMExtraArgsV1Tag, // msg has SVM as destination
			)

			type ProgramData struct {
				DataType uint32
				Address  solana.PublicKey
			}
			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.PingPongProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			ix, err := ping_pong_demo.NewInitializeConfigInstruction(
				config.CcipRouterProgram,
				config.SvmChainSelector,
				config.PingPongProgram[:],
				true, // isPaused
				extraArgs,
				reusableAccounts.Config,
				linkMint,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.PingPongProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
				config.DefaultCommitment)
		})

		t.Run("Initialize", func(t *testing.T) {
			ix, err := ping_pong_demo.NewInitializeInstruction(
				reusableAccounts.Config,
				reusableAccounts.NameVersion,
				config.BillingSignerPDA,
				config.Token2022Program,
				linkMint,
				reusableAccounts.FeeTokenATA,
				reusableAccounts.CcipSendSigner,
				admin.PublicKey(),
				solana.SPLAssociatedTokenAccountProgramID,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
				config.DefaultCommitment)
		})

		t.Run("Fund PingPong LINK ATA", func(t *testing.T) {
			t.Parallel()

			fmt.Printf("Funding LINK ATA: %v\n", reusableAccounts.FeeTokenATA)
			fmt.Printf("Mint: %v\n", linkMint)
			decimals, supply, err := tokens.TokenBalance(ctx, solanaGoClient, reusableAccounts.FeeTokenATA, config.DefaultCommitment)
			require.NoError(t, err)
			fmt.Printf("Supply: %v, Decimals: %v\n", supply, decimals)

			ixMint, err := tokens.MintTo(
				1e9,
				config.Token2022Program,
				linkMint,
				reusableAccounts.FeeTokenATA,
				admin.PublicKey(),
			)
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixMint}, admin,
				config.DefaultCommitment)
		})

		t.Run("Fund PingPong CCIP Send Signer", func(t *testing.T) {
			t.Parallel()

			ix, err := system.NewTransferInstruction(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), reusableAccounts.CcipSendSigner).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
				config.DefaultCommitment)
		})
	})

	t.Run("PingPong", func(t *testing.T) {

		var latestSentMsg ccip_router.SVM2AnyRampMessage

		// Check that the CCIP Send event was emitted
		checkEventCCIPMessageSent := func(expectedSeqNr uint64, logs []string) {
			var event ccip.EventCCIPMessageSent
			require.NoError(t, common.ParseEvent(logs, "CCIPMessageSent", &event, config.PrintEvents))
			require.Equal(t, config.SvmChainSelector, event.DestinationChainSelector)
			require.Equal(t, expectedSeqNr, event.SequenceNumber)

			latestSentMsg = event.Message // persist for later reuse
		}

		t.Run("Start", func(t *testing.T) {
			ix, err := ping_pong_demo.NewStartPingPongInstruction(
				reusableAccounts.Config,
				admin.PublicKey(),
				reusableAccounts.CcipSendSigner,
				config.Token2022Program,
				linkMint,
				reusableAccounts.FeeTokenATA,
				config.CcipRouterProgram,
				config.RouterConfigPDA,
				config.SvmDestChainStatePDA,
				reusableAccounts.nonce,
				reusableAccounts.routerBillingReceiver,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqSvmDestChainPDA,
				reusableAccounts.fqBillingConfig,
				reusableAccounts.fqBillingConfig,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
				config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
			require.NotNil(t, result)

			checkEventCCIPMessageSent(1, result.Meta.LogMessages)
		})

		generateReceiveIx := func() solana.Instruction {
			var extraArgs fee_quoter.SVMExtraArgsV1
			testutils.MustDeserializeExtraArgs(t, &extraArgs, latestSentMsg.ExtraArgs, ccip.SVMExtraArgsV1Tag)

			raw := test_ccip_invalid_receiver.NewReceiverProxyExecuteInstruction(
				test_ccip_invalid_receiver.Any2SVMMessage{
					SourceChainSelector: latestSentMsg.Header.SourceChainSelector,
					MessageId:           latestSentMsg.Header.MessageId,
					Sender:              latestSentMsg.Sender[:],
					Data:                latestSentMsg.Data,
					TokenAmounts:        []example_ccip_receiver.SVMTokenAmount{},
				},
				solana.PublicKey(latestSentMsg.Receiver),
				reusableAccounts.rampSigner,
				reusableAccounts.offramp,
				reusableAccounts.allowedOfframp,
			)

			for i, account := range extraArgs.Accounts {
				meta := solana.Meta(account)
				if (extraArgs.AccountIsWritableBitmap & (uint64(1) << i)) != 0 {
					meta.WRITE()
				}
				raw.AccountMetaSlice.Append(meta)
			}

			ix, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			return ix
		}

		t.Run("Receive & respond (send)", func(t *testing.T) {
			ix := generateReceiveIx()
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin,
				config.DefaultCommitment, common.AddComputeUnitLimit(400_000))
			require.NotNil(t, result)

			checkEventCCIPMessageSent(latestSentMsg.Header.SequenceNumber+1, result.Meta.LogMessages)
		})

		t.Run("Pause", func(t *testing.T) {
			pauseIx, err := ping_pong_demo.NewSetPausedInstruction(
				true,
				reusableAccounts.Config,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			receiveIx := generateReceiveIx()
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{pauseIx, receiveIx}, admin,
				config.DefaultCommitment, common.AddComputeUnitLimit(400_000))
			require.NotNil(t, result)

			// Check that no CCIP Send event was emitted, as the program is paused
			var event ccip.EventCCIPMessageSent
			require.Error(t, common.NewEventNotFoundError("CCIPMessageSent"), common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &event, config.PrintEvents))
		})

		t.Run("Resume", func(t *testing.T) {
			resumeIx, err := ping_pong_demo.NewSetPausedInstruction(
				false,
				reusableAccounts.Config,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			receiveIx := generateReceiveIx()
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{resumeIx, receiveIx}, admin,
				config.DefaultCommitment, common.AddComputeUnitLimit(400_000))
			require.NotNil(t, result)

			checkEventCCIPMessageSent(latestSentMsg.Header.SequenceNumber+1, result.Meta.LogMessages)
		})
	})
}
