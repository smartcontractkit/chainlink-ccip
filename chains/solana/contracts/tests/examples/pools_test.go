package examples

import (
	"fmt"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	tokenpool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

// TestBaseTokenPoolHappyPath does basic happy path checks on the lock/release and burn/mint example pools
// more detailed token pool tests are handled by the test-token-pool which is used in the tokenpool_test.go and ccip_router_test.go
func TestBaseTokenPoolHappyPath(t *testing.T) {
	t.Parallel()

	// acting as "dumb" onramp & offramp, proxying calls to the pool that are signed by PDA
	test_ccip_invalid_receiver.SetProgramID(config.CcipInvalidReceiverProgram)
	rmn_remote.SetProgramID(config.RMNRemoteProgram)

	dumbRamp := config.CcipInvalidReceiverProgram
	allowedOfframpPDA, _ := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbRamp, dumbRamp)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	ctx := tests.Context(t)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)
	getBalance := func(account solana.PublicKey) string {
		balanceRes, err := solanaGoClient.GetTokenAccountBalance(ctx, account, config.DefaultCommitment)
		require.NoError(t, err)
		return balanceRes.Value.Amount
	}

	remotePool := tokenpool.RemoteAddress{Address: []byte{1, 2, 3}}
	remoteToken := tokenpool.RemoteAddress{Address: []byte{4, 5, 6}}

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{admin}, solanaGoClient, t)
		})

		t.Run("register offramp", func(t *testing.T) {
			ix, err := test_ccip_invalid_receiver.NewAddOfframpInstruction(
				config.EvmChainSelector,
				dumbRamp,
				allowedOfframpPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		})

		t.Run("RMN Remote", func(t *testing.T) {
			type ProgramData struct {
				DataType uint32
				Address  solana.PublicKey
			}
			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.RMNRemoteProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			ix, err := rmn_remote.NewInitializeInstruction(
				config.RMNRemoteConfigPDA,
				config.RMNRemoteCursesPDA,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.RMNRemoteProgram,
				programData.Address,
			).ValidateAndBuild()

			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)
		})
	})

	// test functionality with token & token-2022 standards
	for _, v := range []struct {
		tokenName    string
		tokenProgram solana.PublicKey
	}{
		{tokenName: "spl-token", tokenProgram: solana.TokenProgramID},
		{tokenName: "spl-token-2022", tokenProgram: config.Token2022Program},
	} {
		// test functionality with each pool type (burnmint & lockrelease)
		for _, p := range []struct {
			poolName    string
			poolProgram solana.PublicKey
		}{
			{poolName: "burnmint", poolProgram: config.CcipBasePoolBurnMint},
			{poolName: "lockrelease", poolProgram: config.CcipBasePoolLockRelease},
		} {
			t.Run(p.poolName+" "+v.tokenName, func(t *testing.T) {
				t.Parallel()

				decimals := uint8(0)
				amount := uint64(1000)
				poolProgram := p.poolProgram

				// for _, poolProgram := range []solana.PublicKey{config.CcipBasePoolBurnMint} {
				rampPoolSignerPDA, _, _ := state.FindExternalTokenPoolsSignerPDA(poolProgram, dumbRamp)

				mintPriv, err := solana.NewRandomPrivateKey()
				require.NoError(t, err)
				p, err := tokens.NewTokenPool(v.tokenProgram, poolProgram, mintPriv.PublicKey())
				require.NoError(t, err)
				mint := p.Mint

				t.Run("setup:token", func(t *testing.T) {
					// create token
					instructions, err := tokens.CreateToken(ctx, v.tokenProgram, mint, admin.PublicKey(), decimals, solanaGoClient, config.DefaultCommitment)
					require.NoError(t, err)

					// create admin associated token account
					createI, tokenAccount, err := tokens.CreateAssociatedTokenAccount(v.tokenProgram, mint, admin.PublicKey(), admin.PublicKey())
					require.NoError(t, err)
					p.User[admin.PublicKey()] = tokenAccount // set admin token account

					// mint tokens to admin
					mintToI, err := tokens.MintTo(amount, v.tokenProgram, mint, tokenAccount, admin.PublicKey())
					require.NoError(t, err)

					instructions = append(instructions, createI, mintToI)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, instructions, admin, config.DefaultCommitment, common.AddSigners(mintPriv))

					// validate
					outDec, outVal, err := tokens.TokenBalance(ctx, solanaGoClient, p.User[admin.PublicKey()], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, int(amount), outVal)
					require.Equal(t, decimals, outDec)
				})

				t.Run("pool:"+poolProgram.String(), func(t *testing.T) {
					poolConfig, err := tokens.TokenPoolConfigAddress(mint, poolProgram)
					require.NoError(t, err)
					poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
					require.NoError(t, err)
					createI, poolTokenAccount, err := tokens.CreateAssociatedTokenAccount(v.tokenProgram, mint, poolSigner, admin.PublicKey())
					require.NoError(t, err)

					t.Run("setup", func(t *testing.T) {
						type ProgramData struct {
							DataType uint32
							Address  solana.PublicKey
						}
						// get program data account
						data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, poolProgram, &rpc.GetAccountInfoOpts{
							Commitment: config.DefaultCommitment,
						})
						require.NoError(t, err)
						// Decode program data
						var programData ProgramData
						require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

						poolInitI, err := tokenpool.NewInitializeInstruction(dumbRamp, config.RMNRemoteProgram, poolConfig, mint, admin.PublicKey(), solana.SystemProgramID, poolProgram, programData.Address).ValidateAndBuild()
						require.NoError(t, err)

						// make pool mint_authority for token (required for burn/mint)
						authI, err := tokens.SetTokenMintAuthority(v.tokenProgram, poolSigner, mint, admin.PublicKey())
						require.NoError(t, err)

						// set pool config
						ixConfigure, err := tokenpool.NewInitChainRemoteConfigInstruction(
							config.EvmChainSelector,
							p.Mint,
							tokenpool.RemoteConfig{
								TokenAddress: remoteToken,
								Decimals:     9,
							},
							poolConfig,
							p.Chain[config.EvmChainSelector],
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, err)

						ixAppend, err := tokenpool.NewAppendRemotePoolAddressesInstruction(
							config.EvmChainSelector, p.Mint, []tokenpool.RemoteAddress{remotePool}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: poolInitI, Program: poolProgram},
							createI,
							authI,
							&tokens.TokenInstruction{Instruction: ixConfigure, Program: poolProgram},
							&tokens.TokenInstruction{Instruction: ixAppend, Program: poolProgram},
						}, admin, config.DefaultCommitment)

						// Test growing and shrinking allowlist (realloc test)
						a, _ := solana.NewRandomPrivateKey()
						b, _ := solana.NewRandomPrivateKey()

						ixGrowWithDuplicates, err := tokenpool.NewConfigureAllowListInstruction([]solana.PublicKey{a.PublicKey(), b.PublicKey(), b.PublicKey()}, false, poolConfig, mint, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{&tokens.TokenInstruction{Instruction: ixGrowWithDuplicates, Program: poolProgram}}, admin, rpc.CommitmentConfirmed, []string{"Key already existed in the allowlist"})

						ixGrow, err := tokenpool.NewConfigureAllowListInstruction([]solana.PublicKey{a.PublicKey(), b.PublicKey()}, false, poolConfig, mint, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{&tokens.TokenInstruction{Instruction: ixGrow, Program: poolProgram}}, admin, rpc.CommitmentConfirmed)

						ixShrink, err := tokenpool.NewRemoveFromAllowListInstruction([]solana.PublicKey{a.PublicKey(), b.PublicKey()}, poolConfig, mint, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{&tokens.TokenInstruction{Instruction: ixShrink, Program: poolProgram}}, admin, rpc.CommitmentConfirmed)

						// Shrinking fails now as the entries do not exist anymore
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{&tokens.TokenInstruction{Instruction: ixShrink, Program: poolProgram}}, admin, rpc.CommitmentConfirmed, []string{"Key did not exist in the allowlist"})
					})

					t.Run("Cannot re-initialize the state version", func(t *testing.T) {
						var state tokenpool.State
						err := common.GetAccountDataBorshInto(ctx, solanaGoClient, poolConfig, config.DefaultCommitment, &state)
						require.NoError(t, err)
						require.Equal(t, state.Version, uint8(1))

						ix, err := tokenpool.NewInitializeStateVersionInstruction(mint, poolConfig).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ix, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed, []string{"Invalid state version"})
					})

					t.Run("lockOrBurn", func(t *testing.T) {
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))

						transferI, err := tokens.TokenTransferChecked(amount, decimals, v.tokenProgram, p.User[admin.PublicKey()], mint, poolTokenAccount, admin.PublicKey(), solana.PublicKeySlice{})
						require.NoError(t, err)

						lbI, err := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
							test_ccip_invalid_receiver.LockOrBurnInV1{
								LocalToken:          mint,
								Amount:              amount,
								RemoteChainSelector: config.EvmChainSelector,
							},
							poolProgram,
							rampPoolSignerPDA,
							poolConfig,
							v.tokenProgram,
							mint,
							poolSigner,
							poolTokenAccount,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
							p.Chain[config.EvmChainSelector],
						).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI, lbI}, admin, config.DefaultCommitment)
						require.NotNil(t, res)

						// validate balances
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))
						expectedPoolBal := uint64(0)
						if poolProgram == config.CcipBasePoolLockRelease {
							expectedPoolBal = amount
						}
						require.Equal(t, fmt.Sprintf("%d", expectedPoolBal), getBalance(poolTokenAccount))
					})

					t.Run("releaseOrMint", func(t *testing.T) {
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))

						rmI, err := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
							test_ccip_invalid_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(amount * 1e9), // scale to proper decimals
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							},
							poolProgram,
							rampPoolSignerPDA,
							dumbRamp,
							allowedOfframpPDA,
							poolConfig,
							v.tokenProgram,
							mint,
							poolSigner,
							poolTokenAccount,
							p.Chain[config.EvmChainSelector],
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
							p.User[admin.PublicKey()],
						).ValidateAndBuild()
						require.NoError(t, err)

						cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{rmI}, admin, config.DefaultCommitment)

						// This validation is like a snapshot for gas consumption
						// Release or Mint CPI
						require.LessOrEqual(t, cu, uint32(110_000))

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
						require.NotNil(t, res)

						// validate balances
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})

					t.Run("rebalance", func(t *testing.T) {
						// test only relevant for lockrelease pool
						if poolProgram != config.CcipBasePoolLockRelease {
							return
						}

						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))

						approveIx, err := tokens.TokenApproveChecked(amount, decimals, v.tokenProgram, p.User[admin.PublicKey()], p.Mint, poolSigner, admin.PublicKey(), solana.PublicKeySlice{})
						require.NoError(t, err)

						acceptIx, err := tokenpool.NewSetCanAcceptLiquidityInstruction(true, poolConfig, mint, admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						provideIx, err := tokenpool.NewProvideLiquidityInstruction(amount, poolConfig, v.tokenProgram, p.Mint, poolSigner, poolTokenAccount, p.User[admin.PublicKey()], admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							approveIx,
							&tokens.TokenInstruction{Instruction: provideIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed, []string{"Liquidity not accepted"})

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							approveIx,
							&tokens.TokenInstruction{Instruction: acceptIx, Program: poolProgram},
							&tokens.TokenInstruction{Instruction: provideIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed)

						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(poolTokenAccount))

						withdrawIx, err := tokenpool.NewWithdrawLiquidityInstruction(amount, poolConfig, v.tokenProgram, p.Mint, poolSigner, poolTokenAccount, p.User[admin.PublicKey()], admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: withdrawIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed)

						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})
				})
			})
		}
	}
}
