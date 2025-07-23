package examples

import (
	"fmt"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	burnmint_tokenpool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/burnmint_token_pool"
	tokenpool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

type PoolInfo struct {
	poolName    string
	poolProgram solana.PublicKey
}

var pools = []PoolInfo{
	{poolName: "lockrelease", poolProgram: config.CcipBasePoolLockRelease},
	{poolName: "burnmint", poolProgram: config.CcipBasePoolBurnMint},
}

type TokenInfo struct {
	tokenName    string
	tokenProgram solana.PublicKey
	multisig     bool
}

var tokenPrograms = []TokenInfo{
	{tokenName: "spl-token", tokenProgram: solana.TokenProgramID, multisig: false},
	{tokenName: "spl-token", tokenProgram: solana.TokenProgramID, multisig: true},
	{tokenName: "spl-token-2022", tokenProgram: config.Token2022Program, multisig: false},
	{tokenName: "spl-token-2022", tokenProgram: config.Token2022Program, multisig: true},
}

type ProgramData struct {
	DataType uint32
	Address  solana.PublicKey
}

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

	// test functionality with each pool type (burnmint & lockrelease)
	for _, p := range pools {

		var programData ProgramData
		var configPDA solana.PublicKey
		poolProgram := p.poolProgram

		t.Run("setup:tokenPool "+p.poolName, func(t *testing.T) {

			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, p.poolProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)
			// Decode program data
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			// Global Configuration
			configPDA, err = tokens.TokenPoolGlobalConfigPDA(p.poolProgram)
			require.NoError(t, err)

			ix, err := tokenpool.NewInitGlobalConfigInstruction(dumbRamp, config.RMNRemoteProgram, configPDA, admin.PublicKey(), solana.SystemProgramID, poolProgram, programData.Address).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
				&tokens.TokenInstruction{Instruction: ix, Program: p.poolProgram},
			}, admin, config.DefaultCommitment)

		})

		// test functionality with token & token-2022 standards
		for _, v := range tokenPrograms {

			t.Run(p.poolName+" "+v.tokenName, func(t *testing.T) {
				t.Parallel()

				decimals := uint8(0)
				amount := uint64(1000)

				// for _, poolProgram := range []solana.PublicKey{config.CcipBasePoolBurnMint} {
				rampPoolSignerPDA, _, _ := state.FindExternalTokenPoolsSignerPDA(poolProgram, dumbRamp)

				mintPriv, err := solana.NewRandomPrivateKey()
				require.NoError(t, err)
				tokenPool, err := tokens.NewTokenPool(v.tokenProgram, poolProgram, mintPriv.PublicKey())
				require.NoError(t, err)
				mint := tokenPool.Mint

				t.Run("setup:token", func(t *testing.T) {
					// create token
					instructions, err := tokens.CreateToken(ctx, v.tokenProgram, mint, admin.PublicKey(), decimals, solanaGoClient, config.DefaultCommitment)
					require.NoError(t, err)

					// create admin associated token account
					createI, tokenAccount, err := tokens.CreateAssociatedTokenAccount(v.tokenProgram, mint, admin.PublicKey(), admin.PublicKey())
					require.NoError(t, err)
					tokenPool.User[admin.PublicKey()] = tokenAccount // set admin token account

					// mint tokens to admin
					mintToI, err := tokens.MintTo(amount, v.tokenProgram, mint, tokenAccount, admin.PublicKey())
					require.NoError(t, err)

					instructions = append(instructions, createI, mintToI)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, instructions, admin, config.DefaultCommitment, common.AddSigners(mintPriv))

					// validate
					outDec, outVal, err := tokens.TokenBalance(ctx, solanaGoClient, tokenPool.User[admin.PublicKey()], config.DefaultCommitment)
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

						poolInitI, err := tokenpool.NewInitializeInstruction(poolConfig, mint, admin.PublicKey(), solana.SystemProgramID, poolProgram, programData.Address, configPDA).ValidateAndBuild()
						require.NoError(t, err)

						newMintAuthority := poolSigner

						if v.multisig {
							// create multisig
							multisig, err := solana.NewRandomPrivateKey()
							require.NoError(t, err)
							ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), v.tokenProgram, multisig.PublicKey(), 1, []solana.PublicKey{admin.PublicKey(), poolSigner}, solanaGoClient, config.DefaultCommitment)
							require.NoError(t, ixErrMsig)
							testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(multisig, admin))

							newMintAuthority = multisig.PublicKey()
							tokenPool.MintAuthorityMultisig = multisig.PublicKey() // set multisig as mint authority
						}

						// make pool mint_authority for token (required for burn/mint) poolsigner or multisig
						authI, err := tokens.SetTokenMintAuthority(v.tokenProgram, newMintAuthority, mint, admin.PublicKey())
						require.NoError(t, err)

						// set pool config
						ixConfigure, err := tokenpool.NewInitChainRemoteConfigInstruction(
							config.EvmChainSelector,
							tokenPool.Mint,
							tokenpool.RemoteConfig{
								TokenAddress: remoteToken,
								Decimals:     9,
							},
							poolConfig,
							tokenPool.Chain[config.EvmChainSelector],
							admin.PublicKey(),
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, err)

						ixAppend, err := tokenpool.NewAppendRemotePoolAddressesInstruction(
							config.EvmChainSelector, tokenPool.Mint, []tokenpool.RemoteAddress{remotePool}, poolConfig, tokenPool.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
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
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(tokenPool.User[admin.PublicKey()]))

						transferI, err := tokens.TokenTransferChecked(amount, decimals, v.tokenProgram, tokenPool.User[admin.PublicKey()], mint, poolTokenAccount, admin.PublicKey(), solana.PublicKeySlice{})
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
							tokenPool.Chain[config.EvmChainSelector],
						).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI, lbI}, admin, config.DefaultCommitment)
						require.NotNil(t, res)

						// validate balances
						require.Equal(t, "0", getBalance(tokenPool.User[admin.PublicKey()]))
						expectedPoolBal := uint64(0)
						if poolProgram == config.CcipBasePoolLockRelease {
							expectedPoolBal = amount
						}
						require.Equal(t, fmt.Sprintf("%d", expectedPoolBal), getBalance(poolTokenAccount))
					})

					t.Run("releaseOrMint", func(t *testing.T) {
						require.Equal(t, "0", getBalance(tokenPool.User[admin.PublicKey()]))

						raw := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
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
							tokenPool.Chain[config.EvmChainSelector],
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
							tokenPool.User[admin.PublicKey()],
						)

						if v.multisig {
							raw.AccountMetaSlice.Append(solana.Meta(tokenPool.MintAuthorityMultisig))
						}

						rmI, err := raw.ValidateAndBuild()
						require.NoError(t, err)

						cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{rmI}, admin, config.DefaultCommitment)

						// This validation is like a snapshot for gas consumption
						// Release or Mint CPI
						require.LessOrEqual(t, cu, uint32(110_000))

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, admin, config.DefaultCommitment, common.AddComputeUnitLimit(cu))
						require.NotNil(t, res)

						// validate balances
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(tokenPool.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})

					t.Run("rebalance", func(t *testing.T) {
						// test only relevant for lockrelease pool
						if poolProgram != config.CcipBasePoolLockRelease {
							return
						}

						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(tokenPool.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))

						approveIx, err := tokens.TokenApproveChecked(amount, decimals, v.tokenProgram, tokenPool.User[admin.PublicKey()], tokenPool.Mint, poolSigner, admin.PublicKey(), solana.PublicKeySlice{})
						require.NoError(t, err)

						acceptIx, err := tokenpool.NewSetCanAcceptLiquidityInstruction(true, poolConfig, mint, admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						setRebalancerIx, err := tokenpool.NewSetRebalancerInstruction(admin.PublicKey(), poolConfig, mint, admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						provideIx, err := tokenpool.NewProvideLiquidityInstruction(amount, poolConfig, v.tokenProgram, tokenPool.Mint, poolSigner, poolTokenAccount, tokenPool.User[admin.PublicKey()], admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							approveIx,
							&tokens.TokenInstruction{Instruction: setRebalancerIx, Program: poolProgram},
							&tokens.TokenInstruction{Instruction: provideIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed, []string{"Liquidity not accepted"})

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							approveIx,
							&tokens.TokenInstruction{Instruction: acceptIx, Program: poolProgram},
							&tokens.TokenInstruction{Instruction: setRebalancerIx, Program: poolProgram},
							&tokens.TokenInstruction{Instruction: provideIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed)

						require.Equal(t, "0", getBalance(tokenPool.User[admin.PublicKey()]))
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(poolTokenAccount))

						withdrawIx, err := tokenpool.NewWithdrawLiquidityInstruction(amount, poolConfig, v.tokenProgram, tokenPool.Mint, poolSigner, poolTokenAccount, tokenPool.User[admin.PublicKey()], admin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: withdrawIx, Program: poolProgram},
						}, admin, rpc.CommitmentConfirmed)

						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(tokenPool.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})
				})

				// If it wasn't a multisig, then we test the upgrade path to change the mint authority to a multisig
				if p.poolName == "burnmint" && !v.multisig {

					poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
					require.NoError(t, err)
					poolConfig, err := tokens.TokenPoolConfigAddress(mint, poolProgram)
					require.NoError(t, err)

					// create multisig
					multisig, err := solana.NewRandomPrivateKey()
					require.NoError(t, err)
					ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), v.tokenProgram, multisig.PublicKey(), 1, []solana.PublicKey{admin.PublicKey(), poolSigner}, solanaGoClient, config.DefaultCommitment)
					require.NoError(t, ixErrMsig)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(multisig, admin))

					t.Run("try to upgrade to invalid token program multisig", func(t *testing.T) {

						// create invalidMultisig
						invalidMultisig, err := solana.NewRandomPrivateKey()
						require.NoError(t, err)
						poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
						require.NoError(t, err)

						var invalidTokenProgram solana.PublicKey
						if v.tokenProgram == solana.TokenProgramID {
							invalidTokenProgram = solana.Token2022ProgramID
						} else {
							invalidTokenProgram = solana.TokenProgramID
						}
						ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), invalidTokenProgram, invalidMultisig.PublicKey(), 1, []solana.PublicKey{admin.PublicKey(), poolSigner}, solanaGoClient, config.DefaultCommitment)
						require.NoError(t, ixErrMsig)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(invalidMultisig, admin))

						newMintAuthority := invalidMultisig.PublicKey()
						tokenPool.MintAuthorityMultisig = newMintAuthority

						// set multisig as mint authority
						ixTransferMint, err := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							newMintAuthority,
							p.poolProgram,
							programData.Address,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment, []string{"InvalidMultisigOwner."})

						// Check that the mint authority was not changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(poolSigner), *mintAccount.MintAuthority)
					})

					t.Run("try to upgrade to invalid m configured multisig", func(t *testing.T) {

						// create invalidMultisig
						invalidMultisig, err := solana.NewRandomPrivateKey()
						require.NoError(t, err)
						poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
						require.NoError(t, err)

						ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), v.tokenProgram, invalidMultisig.PublicKey(), 2, []solana.PublicKey{admin.PublicKey(), poolSigner}, solanaGoClient, config.DefaultCommitment)
						require.NoError(t, ixErrMsig)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(invalidMultisig, admin))

						newMintAuthority := invalidMultisig.PublicKey()
						tokenPool.MintAuthorityMultisig = newMintAuthority

						// set multisig as mint authority
						ixTransferMint, err := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							newMintAuthority,
							p.poolProgram,
							programData.Address,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment, []string{"PoolSignerNotInMultisig"})

						// Check that the mint authority was not changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(poolSigner), *mintAccount.MintAuthority)
					})

					t.Run("try to upgrade to multisig without poolsigner", func(t *testing.T) {

						// create invalidMultisig
						invalidMultisig, err := solana.NewRandomPrivateKey()
						require.NoError(t, err)
						poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
						require.NoError(t, err)

						ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), v.tokenProgram, invalidMultisig.PublicKey(), 1, []solana.PublicKey{admin.PublicKey(), admin.PublicKey()}, solanaGoClient, config.DefaultCommitment)
						require.NoError(t, ixErrMsig)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(invalidMultisig, admin))

						newMintAuthority := invalidMultisig.PublicKey()
						tokenPool.MintAuthorityMultisig = newMintAuthority

						// set multisig as mint authority
						ixTransferMint, err := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							newMintAuthority,
							p.poolProgram,
							programData.Address,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment, []string{"PoolSignerNotInMultisig"})

						// Check that the mint authority was not changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(poolSigner), *mintAccount.MintAuthority)
					})
					t.Run("upgrade to valid multisig", func(t *testing.T) {

						newMintAuthority := multisig.PublicKey()
						tokenPool.MintAuthorityMultisig = newMintAuthority

						// set multisig as mint authority
						ixTransferMint, err := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							newMintAuthority,
							p.poolProgram,
							programData.Address,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment)

						// Check that the mint authority was changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(newMintAuthority), *mintAccount.MintAuthority)
					})
					t.Run("try to transfer back", func(t *testing.T) {

						poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
						require.NoError(t, err)
						poolConfig, err := tokens.TokenPoolConfigAddress(mint, poolProgram)
						require.NoError(t, err)

						// rolling back the change needs to be done manually, as the pool program does not support it
						ixTransferMint, err := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							poolSigner,
							p.poolProgram,
							programData.Address,
						).ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment, []string{"InvalidMultisigOwner"})

						// Check that the mint authority was not changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(multisig.PublicKey()), *mintAccount.MintAuthority)
					})
					t.Run("transfer to another multisig", func(t *testing.T) {

						poolSigner, err := tokens.TokenPoolSignerAddress(mint, poolProgram)
						require.NoError(t, err)
						poolConfig, err := tokens.TokenPoolConfigAddress(mint, poolProgram)
						require.NoError(t, err)

						// create multisig
						anotherMultisig, err := solana.NewRandomPrivateKey()
						require.NoError(t, err)
						ixMsig, ixErrMsig := tokens.CreateMultisig(ctx, admin.PublicKey(), v.tokenProgram, anotherMultisig.PublicKey(), 1, []solana.PublicKey{admin.PublicKey(), poolSigner}, solanaGoClient, config.DefaultCommitment)
						require.NoError(t, ixErrMsig)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, ixMsig, admin, config.DefaultCommitment, common.AddSigners(anotherMultisig, admin))

						raw := burnmint_tokenpool.NewTransferMintAuthorityToMultisigInstruction(
							poolConfig,
							tokenPool.Mint,
							v.tokenProgram,
							poolSigner,
							admin.PublicKey(),
							anotherMultisig.PublicKey(),
							p.poolProgram,
							programData.Address,
						)

						raw.AccountMetaSlice.Append(solana.Meta(multisig.PublicKey()))

						ixTransferMint, err := raw.ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
							&tokens.TokenInstruction{Instruction: ixTransferMint, Program: poolProgram},
						}, admin, config.DefaultCommitment)

						// Check that the mint authority was changed in the Mint Account
						mintAccount := token.Mint{}
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, mint, config.DefaultCommitment, &mintAccount)
						require.NoError(t, err)

						require.Equal(t, solana.PublicKey(anotherMultisig.PublicKey()), *mintAccount.MintAuthority)
					})
				}
			})
		}

		// Test self onboarding to the pool
		t.Run("self-onboard", func(t *testing.T) {

			// Enable self-served token pool onboarding
			ix, err := tokenpool.NewUpdateSelfServedAllowedInstruction(true, configPDA, admin.PublicKey(), poolProgram, programData.Address).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
				&tokens.TokenInstruction{Instruction: ix, Program: p.poolProgram},
			}, admin, config.DefaultCommitment)

			// Create a new token owner
			newOwner, err := solana.NewRandomPrivateKey()
			require.NoError(t, err)
			testutils.FundAccounts(ctx, []solana.PrivateKey{newOwner}, solanaGoClient, t)

			// Initialize a Token Pool in the CLL Token Pool Program for my own mint
			for _, v := range tokenPrograms {
				t.Run("self-onboard "+v.tokenName, func(t *testing.T) {
					t.Parallel()

					mintPriv, err := solana.NewRandomPrivateKey()
					require.NoError(t, err)
					p, err := tokens.NewTokenPool(v.tokenProgram, p.poolProgram, mintPriv.PublicKey())
					require.NoError(t, err)
					mint := p.Mint

					// create token
					instructions, err := tokens.CreateToken(ctx, v.tokenProgram, mint, newOwner.PublicKey(), uint8(6), solanaGoClient, config.DefaultCommitment)
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, instructions, newOwner, config.DefaultCommitment, common.AddSigners(mintPriv))

					poolConfig, err := tokens.TokenPoolConfigAddress(mint, p.PoolProgram)
					require.NoError(t, err)

					poolInitI, err := tokenpool.NewInitializeInstruction(poolConfig, mint, newOwner.PublicKey(), solana.SystemProgramID, p.PoolProgram, programData.Address, configPDA).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{
						&tokens.TokenInstruction{Instruction: poolInitI, Program: p.PoolProgram},
					}, newOwner, config.DefaultCommitment)

				})
			}
		})
	}
}
