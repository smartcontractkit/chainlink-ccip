package examples

import (
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	bin "github.com/gagliardetto/binary"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ping_pong_demo"
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

	return ReusableAccounts{
		FeeTokenATA:           ppLinkAta,
		Config:                ppConfig,
		NameVersion:           ppNameVersion,
		CcipSendSigner:        ppCcipSendSigner,
		fqBillingConfig:       fqBillingConfig,
		routerBillingReceiver: routerBillingReceiver,
	}
}

// Test basic happy path of the ping pong demo with "itself", meaning an SVM <-> SVM message
// from the ping pong program to itself.
func TestPingPong(t *testing.T) {
	t.Parallel()

	// acting as "dumb" onramp & offramp, proxying calls to the pool that are signed by PDA
	ccip_offramp.SetProgramID(config.CcipOfframpProgram)
	fee_quoter.SetProgramID(config.FeeQuoterProgram)
	ccip_router.SetProgramID(config.CcipRouterProgram)
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
					config.ExternalTokenPoolsSignerPDA,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin,
					config.DefaultCommitment)
			})

			t.Run("Offramp", func(t *testing.T) {
				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(
					ctx,
					config.CcipOfframpProgram,
					&rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment},
				)
				require.NoError(t, err)

				// Decode program data
				var programData ProgramData
				require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

				// Now, actually initialize the offramp
				initIx, err := ccip_offramp.NewInitializeInstruction(
					config.OfframpReferenceAddressesPDA,
					config.CcipRouterProgram,
					config.FeeQuoterProgram,
					config.RMNRemoteProgram,
					solana.PublicKey{}, // lookup table, not used in the tests, so just send an empty address,
					config.OfframpStatePDA,
					config.OfframpExternalExecutionConfigPDA,
					config.OfframpTokenPoolsSignerPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
					config.CcipOfframpProgram,
					programData.Address,
				).ValidateAndBuild()
				require.NoError(t, err)

				initConfigIx, err := ccip_offramp.NewInitializeConfigInstruction(
					config.SvmChainSelector,
					config.EnableExecutionAfter,
					config.OfframpConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
					config.CcipOfframpProgram,
					programData.Address,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initIx, initConfigIx}, admin,
					config.DefaultCommitment)
			})
		})

		t.Run("Other CCIP configs", func(t *testing.T) {
			t.Run("register offramp", func(t *testing.T) {
				routerIx, err := ccip_router.NewAddOfframpInstruction(
					config.SvmChainSelector,
					config.CcipOfframpProgram,
					config.AllowedOfframpSvmPDA,
					config.RouterConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				fqIx, err := fee_quoter.NewAddPriceUpdaterInstruction(
					config.OfframpBillingSignerPDA,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{routerIx, fqIx}, admin,
					config.DefaultCommitment)
			})

			t.Run("Register link as billing token", func(t *testing.T) {
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

			t.Run("Register SVM as source chain", func(t *testing.T) {
				var onRampAddress [64]byte
				copy(onRampAddress[:], config.CcipRouterProgram[:]) // the router is the SVM onramp

				instruction, err := ccip_offramp.NewAddSourceChainInstruction(
					config.SvmChainSelector,
					ccip_offramp.SourceChainConfig{
						OnRamp: ccip_offramp.OnRampAddress{
							Bytes: onRampAddress,
							Len:   32,
						},
						IsEnabled: true,
					},
					config.OfframpSvmSourceChainPDA,
					config.OfframpConfigPDA,
					admin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin,
					config.DefaultCommitment)
			})

			t.Run("Register SVM as dest chain", func(t *testing.T) {
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
		t.Run("Initialize", func(t *testing.T) {
			// TODO
		})

		t.Run("Fund PingPong LINK ATA", func(t *testing.T) {
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
	})
}
