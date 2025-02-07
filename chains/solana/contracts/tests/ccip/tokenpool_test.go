package contracts

import (
	"fmt"
	"strings"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestTokenPool(t *testing.T) {
	token_pool.SetProgramID(config.CcipTokenPoolProgram)
	ccip_router.SetProgramID(config.CcipRouterProgram)

	// acting as "dumb" offramp that just proxies the pool,
	// required for authnz in the pool but we don't want to test offramp internals here
	test_ccip_receiver.SetProgramID(config.CcipLogicReceiver)
	dumbOfframp := config.CcipLogicReceiver
	dumbOfframpPoolSigner, _, _ := solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer")}, dumbOfframp)

	admin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	anotherAdmin, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	ctx := tests.Context(t)

	allowedOfframpPDA, err := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbOfframp, config.CcipRouterProgram)
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)
	getBalance := func(account solana.PublicKey) string {
		balanceRes, err := solanaGoClient.GetTokenAccountBalance(ctx, account, config.DefaultCommitment)
		require.NoError(t, err)
		return balanceRes.Value.Amount
	}

	remotePool := token_pool.RemoteAddress{Address: []byte{1, 2, 3}}
	remoteToken := token_pool.RemoteAddress{Address: []byte{4, 5, 6}}

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anotherAdmin, user}, solanaGoClient, t)
	})

	t.Run("setup:router", func(t *testing.T) {
		// get program data account
		data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram, &rpc.GetAccountInfoOpts{
			Commitment: config.DefaultCommitment,
		})
		require.NoError(t, err)

		// Decode program data
		var programData struct {
			DataType uint32
			Address  solana.PublicKey
		}
		require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

		feeAggregator := solana.MustPrivateKeyFromBase58("4mKsN4bLEPTQerRRCMALWMFKnkP1xiaC3rYCzcmEmgCu5yrf2eDCPH3jHbsaAg1giKKFwrxk9oUzVxHLYokS1QhN")
		linkTokenMint := solana.MustPrivateKeyFromBase58("2e6af6HmHgxmrv5dLVSqAzerPrLsjEJyyRATvjiBLPpahFv3wdE2NQqaHWjtb8WdVLrvoLchNLoHBr4KVC1GAxBC")

		ix, err := ccip_router.NewInitializeInstruction(
			1,
			feeAggregator.PublicKey(),
			config.FeeQuoterProgram,
			linkTokenMint.PublicKey(),
			config.RouterConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
			config.CcipRouterProgram,
			programData.Address,
			config.ExternalTokenPoolsSignerPDA,
		).ValidateAndBuild()
		require.NoError(t, err)
		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
	})

	t.Run("setup:allowed offramp", func(t *testing.T) {
		ix, err := ccip_router.NewAddOfframpInstruction(
			config.EvmChainSelector,
			dumbOfframp,
			allowedOfframpPDA,
			config.RouterConfigPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
	})

	// test functionality with token & token-2022 standards
	for _, v := range []struct {
		tokenName    string
		tokenProgram solana.PublicKey
	}{
		{tokenName: "spl-token", tokenProgram: solana.TokenProgramID},
		{tokenName: "spl-token-2022", tokenProgram: config.Token2022Program},
	} {
		t.Run(v.tokenName, func(t *testing.T) {
			t.Parallel()
			decimals := uint8(0)
			amount := uint64(1000)

			for _, poolType := range []token_pool.PoolType{token_pool.LockAndRelease_PoolType, token_pool.BurnAndMint_PoolType} {
				p, err := tokens.NewTokenPool(v.tokenProgram)
				require.NoError(t, err)
				mint := p.Mint.PublicKey()

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
					testutils.SendAndConfirm(ctx, t, solanaGoClient, instructions, admin, config.DefaultCommitment, common.AddSigners(p.Mint))

					// validate
					outDec, outVal, err := tokens.TokenBalance(ctx, solanaGoClient, p.User[admin.PublicKey()], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, int(amount), outVal)
					require.Equal(t, decimals, outDec)
				})

				t.Run("pool:"+poolType.String(), func(t *testing.T) {
					poolConfig, err := tokens.TokenPoolConfigAddress(mint, config.CcipTokenPoolProgram)
					require.NoError(t, err)
					poolSigner, err := tokens.TokenPoolSignerAddress(mint, config.CcipTokenPoolProgram)
					require.NoError(t, err)
					createI, poolTokenAccount, err := tokens.CreateAssociatedTokenAccount(v.tokenProgram, mint, poolSigner, admin.PublicKey())
					require.NoError(t, err)

					// LockAndRelease => [Lock, Release]
					// BurnAndMint => [Burn, Mint]
					poolMethodName := strings.Split(poolType.String(), "And")
					require.Equal(t, 2, len(poolMethodName))
					lockOrBurn := poolMethodName[0]
					releaseOrMint := poolMethodName[1]

					t.Run("setup", func(t *testing.T) {
						poolInitI, err := token_pool.NewInitializeInstruction(poolType,
							admin.PublicKey(),
							config.CcipRouterProgram,
							poolConfig,
							mint,
							poolSigner,
							admin.PublicKey(),
							solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						// make pool mint_authority for token (required for burn/mint)
						authI, err := tokens.SetTokenMintAuthority(v.tokenProgram, poolSigner, mint, admin.PublicKey())
						require.NoError(t, err)

						// set rate limit
						ixRates, err := token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint.PublicKey(), token_pool.RateLimitConfig{
							Enabled:  true,
							Capacity: amount,
							Rate:     1, // slow refill
						}, token_pool.RateLimitConfig{
							Enabled:  false,
							Capacity: 0,
							Rate:     0,
						}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						// set pool config
						ixConfigure, err := token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, p.Mint.PublicKey(), token_pool.RemoteConfig{
							TokenAddress: remoteToken,
							Decimals:     9,
						}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						ixAppend, err := token_pool.NewAppendRemotePoolAddressesInstruction(
							config.EvmChainSelector, p.Mint.PublicKey(), []token_pool.RemoteAddress{remotePool}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{poolInitI, createI, authI, ixConfigure, ixAppend, ixRates}, admin, config.DefaultCommitment)
						require.NotNil(t, res)

						var configAccount token_pool.Config
						require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, poolConfig, config.DefaultCommitment, &configAccount))
						require.Equal(t, poolTokenAccount, configAccount.PoolTokenAccount)

						eventConfigured := tokens.EventChainConfigured{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainConfigured", &eventConfigured, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventConfigured.ChainSelector)
						require.Equal(t, 0, len(eventConfigured.PoolAddresses))
						require.Equal(t, 0, len(eventConfigured.PreviousPoolAddresses))
						require.Equal(t, remoteToken, eventConfigured.Token)
						require.Equal(t, 0, len(eventConfigured.PreviousToken.Address))

						eventAppended := tokens.EventRemotePoolsAppended{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemotePoolsAppended", &eventAppended, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventAppended.ChainSelector)
						require.Equal(t, []token_pool.RemoteAddress{remotePool}, eventAppended.PoolAddresses)
						require.Equal(t, 0, len(eventAppended.PreviousPoolAddresses))

						eventRateLimit := tokens.EventRateLimitConfigured{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RateLimitConfigured", &eventRateLimit, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventRateLimit.ChainSelector)
						require.Equal(t, true, eventRateLimit.InboundRateLimit.Enabled)
						require.Equal(t, amount, eventRateLimit.InboundRateLimit.Capacity)
						require.Equal(t, uint64(1), eventRateLimit.InboundRateLimit.Rate)
						require.Equal(t, false, eventRateLimit.OutboundRateLimit.Enabled)
						require.Equal(t, uint64(0), eventRateLimit.OutboundRateLimit.Capacity)
						require.Equal(t, uint64(0), eventRateLimit.OutboundRateLimit.Rate)
					})

					t.Run("admin:ownership", func(t *testing.T) {
						// successfully transfer ownership
						instruction, err := token_pool.NewTransferOwnershipInstruction(
							anotherAdmin.PublicKey(),
							poolConfig,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, err)
						result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
						require.NotNil(t, result)

						// Successfully accept ownership
						// anotherAdmin becomes owner for remaining tests
						instruction, err = token_pool.NewAcceptOwnershipInstruction(
							poolConfig,
							anotherAdmin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, err)
						result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherAdmin, config.DefaultCommitment)
						require.NotNil(t, result)
					})

					t.Run(lockOrBurn, func(t *testing.T) {
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))

						transferI, err := tokens.TokenTransferChecked(amount, decimals, v.tokenProgram, p.User[admin.PublicKey()], mint, poolTokenAccount, admin.PublicKey(), solana.PublicKeySlice{})
						require.NoError(t, err)

						lbI, err := token_pool.NewLockOrBurnTokensInstruction(token_pool.LockOrBurnInV1{
							LocalToken:          mint,
							Amount:              amount,
							RemoteChainSelector: config.EvmChainSelector,
						}, admin.PublicKey(), poolConfig, v.tokenProgram, mint, poolSigner, poolTokenAccount, p.Chain[config.EvmChainSelector]).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI, lbI}, admin, config.DefaultCommitment)
						require.NotNil(t, res)

						event := tokens.EventBurnLock{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, tokens.MethodToEvent(lockOrBurn), &event))
						require.Equal(t, admin.PublicKey(), event.Sender)
						require.Equal(t, amount, event.Amount)

						// validate balances
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))
						expectedPoolBal := uint64(0)
						if poolType == token_pool.LockAndRelease_PoolType {
							expectedPoolBal = amount
						}
						require.Equal(t, fmt.Sprintf("%d", expectedPoolBal), getBalance(poolTokenAccount))
					})

					t.Run(releaseOrMint, func(t *testing.T) {
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))

						rmI, err := test_ccip_receiver.NewPoolProxyReleaseOrMintInstruction(
							test_ccip_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(amount * 1e9), // scale to proper decimals
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							},
							config.CcipTokenPoolProgram,
							dumbOfframpPoolSigner,
							dumbOfframp,
							allowedOfframpPDA,
							poolConfig,
							v.tokenProgram,
							mint,
							poolSigner,
							poolTokenAccount,
							p.Chain[config.EvmChainSelector],
							p.User[admin.PublicKey()],
						).ValidateAndBuild()

						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment)
						require.NotNil(t, res)

						event := tokens.EventMintRelease{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, tokens.MethodToEvent(releaseOrMint), &event))
						require.Equal(t, admin.PublicKey(), event.Recipient)
						require.Equal(t, poolSigner, event.Sender)
						require.Equal(t, amount, event.Amount)

						// validate balances
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})

					t.Run("invalid", func(t *testing.T) {
						t.Run("config", func(t *testing.T) {
							t.Parallel()

							cfgs := []struct {
								name   string
								c      token_pool.RateLimitConfig
								errStr string
							}{
								{
									name: "enabled-zero-rate",
									c: token_pool.RateLimitConfig{
										Enabled:  true,
										Capacity: 10,
										Rate:     0,
									},
									errStr: "invalid rate limit rate",
								},
								{
									name: "enabled-rate-larger-than-capacity",
									c: token_pool.RateLimitConfig{
										Enabled:  true,
										Capacity: 1,
										Rate:     100,
									},
									errStr: "invalid rate limit rate",
								},
								{
									name: "disabled-nonzero-rate",
									c: token_pool.RateLimitConfig{
										Enabled:  false,
										Capacity: 0,
										Rate:     100,
									},
									errStr: "disabled non-zero rate limit",
								},
								{
									name: "disabled-nonzero-capacity",
									c: token_pool.RateLimitConfig{
										Enabled:  false,
										Capacity: 10,
										Rate:     0,
									},
									errStr: "disabled non-zero rate limit",
								},
							}

							for _, cfg := range cfgs {
								t.Run(cfg.name, func(t *testing.T) {
									// t.Parallel()

									ixRates, err := token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint.PublicKey(), cfg.c, token_pool.RateLimitConfig{}, poolConfig, p.Chain[config.EvmChainSelector], anotherAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
									require.NoError(t, err)

									testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ixRates}, anotherAdmin, config.DefaultCommitment, []string{cfg.errStr})
								})
							}
						})

						t.Run("exceed-rate-limit", func(t *testing.T) {
							t.Parallel()

							// exceed capacity of bucket
							rmI, err := test_ccip_receiver.NewPoolProxyReleaseOrMintInstruction(test_ccip_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(amount * amount * 1e9),
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							}, config.CcipTokenPoolProgram, dumbOfframpPoolSigner, dumbOfframp, allowedOfframpPDA, poolConfig, v.tokenProgram, mint, poolSigner, poolTokenAccount, p.Chain[config.EvmChainSelector], p.User[admin.PublicKey()]).ValidateAndBuild()
							require.NoError(t, err)
							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment, []string{"max capacity exceeded"})

							// exceed rate limit of transfer
							// request two release/mint of max capacity
							// if first does not exceed limit, the second one should
							transferI, err := tokens.TokenTransferChecked(amount,
								decimals,
								v.tokenProgram,
								p.User[admin.PublicKey()],
								mint,
								poolTokenAccount,
								admin.PublicKey(),
								solana.PublicKeySlice{}) // ensure pool is funded

							require.NoError(t, err)
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI}, admin, config.DefaultCommitment)
							rmI, err = test_ccip_receiver.NewPoolProxyReleaseOrMintInstruction(test_ccip_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(amount * 1e9),
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							}, config.CcipTokenPoolProgram, dumbOfframpPoolSigner, dumbOfframp, allowedOfframpPDA, poolConfig, v.tokenProgram, mint, poolSigner, poolTokenAccount, p.Chain[config.EvmChainSelector], p.User[admin.PublicKey()]).ValidateAndBuild()
							require.NoError(t, err)
							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{rmI, rmI}, user, config.DefaultCommitment, []string{"rate limit reached"})

							// pool should refill automatically, but slowly
							// small amount should pass
							time.Sleep(time.Second) // wait for refill
							rmI, err = test_ccip_receiver.NewPoolProxyReleaseOrMintInstruction(test_ccip_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(1e9),
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							}, config.CcipTokenPoolProgram, dumbOfframpPoolSigner, dumbOfframp, allowedOfframpPDA, poolConfig, v.tokenProgram, mint, poolSigner, poolTokenAccount, p.Chain[config.EvmChainSelector], p.User[admin.PublicKey()]).ValidateAndBuild()
							require.NoError(t, err)
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment)
						})
					})

					t.Run("closing", func(t *testing.T) {
						ixDelete, err := token_pool.NewDeleteChainConfigInstruction(config.EvmChainSelector, mint, poolConfig, p.Chain[config.EvmChainSelector], anotherAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						ixRouterChange, err := token_pool.NewSetRampAuthorityInstruction(config.ExternalTokenPoolsSignerPDA, poolConfig, anotherAdmin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixDelete, ixRouterChange}, anotherAdmin, config.DefaultCommitment)
						require.NotNil(t, res)

						eventDelete := tokens.EventChainRemoved{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainRemoved", &eventDelete, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventDelete.ChainSelector)

						eventRouter := tokens.EventRouterUpdated{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RampAuthorityUpdated", &eventRouter, config.PrintEvents))
						require.Equal(t, config.ExternalTokenPoolsSignerPDA, eventRouter.NewAuthority)
						require.Equal(t, admin.PublicKey(), eventRouter.OldAuthority)
					})
				})
			}
		})
	}

	// test functionality with arbitrary wrapped program
	t.Run("Wrapped", func(t *testing.T) {
		t.Parallel()
		p, err := tokens.NewTokenPool(solana.TokenProgramID)
		require.NoError(t, err)
		mint := p.Mint.PublicKey()

		t.Run("setup:pool", func(t *testing.T) {
			var err error
			p.PoolConfig, err = tokens.TokenPoolConfigAddress(mint, config.CcipTokenPoolProgram)
			require.NoError(t, err)
			p.PoolSigner, err = tokens.TokenPoolSignerAddress(mint, config.CcipTokenPoolProgram)
			require.NoError(t, err)

			// create token
			instructions, err := tokens.CreateToken(ctx, solana.TokenProgramID, mint, admin.PublicKey(), 0, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, err)

			// create pool
			poolInitI, err := token_pool.NewInitializeInstruction(token_pool.Wrapped_PoolType,
				admin.PublicKey(),
				config.CcipRouterProgram,
				p.PoolConfig,
				mint,
				p.PoolSigner,
				admin.PublicKey(),
				solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			// create admin receiver token account
			var createR solana.Instruction
			createR, p.User[admin.PublicKey()], err = tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, mint, admin.PublicKey(), admin.PublicKey())
			require.NoError(t, err)

			// create pool token account
			var createP solana.Instruction
			createP, p.PoolTokenAccount, err = tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, mint, p.PoolSigner, admin.PublicKey())
			require.NoError(t, err)

			// initialize pool config
			configureI, err := token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, p.Mint.PublicKey(), token_pool.RemoteConfig{
				TokenAddress: remoteToken,
			}, p.PoolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			// append remote pools
			appendI, err := token_pool.NewAppendRemotePoolAddressesInstruction(
				config.EvmChainSelector, p.Mint.PublicKey(), []token_pool.RemoteAddress{remotePool}, p.PoolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			res := testutils.SendAndConfirm(ctx, t, solanaGoClient, append(instructions, poolInitI, createR, createP, configureI, appendI), admin, config.DefaultCommitment, common.AddSigners(p.Mint))
			require.NotNil(t, res)
		})

		t.Run("burnOrLock", func(t *testing.T) {
			raw := token_pool.NewLockOrBurnTokensInstruction(token_pool.LockOrBurnInV1{LocalToken: mint, RemoteChainSelector: config.EvmChainSelector}, admin.PublicKey(), p.PoolConfig, solana.TokenProgramID, mint, p.PoolSigner, p.PoolTokenAccount, p.Chain[config.EvmChainSelector])
			raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.NewAccountMeta(config.CcipLogicReceiver, false, false))
			lbI, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{lbI}, admin, config.DefaultCommitment)
			require.NotNil(t, res)
			require.Contains(t, strings.Join(res.Meta.LogMessages, "\n"), "Called `ccip_token_lock_burn`")
		})

		t.Run("mintOrRelease", func(t *testing.T) {
			raw := test_ccip_receiver.NewPoolProxyReleaseOrMintInstruction(
				test_ccip_receiver.ReleaseOrMintInV1{
					LocalToken:          mint,
					SourcePoolAddress:   remotePool.Address,
					Receiver:            admin.PublicKey(),
					RemoteChainSelector: config.EvmChainSelector,
					Amount:              tokens.ToLittleEndianU256(1),
				},
				config.CcipTokenPoolProgram,
				dumbOfframpPoolSigner,
				dumbOfframp,
				allowedOfframpPDA,
				p.PoolConfig,
				solana.TokenProgramID,
				mint,
				p.PoolSigner,
				p.PoolTokenAccount,
				p.Chain[config.EvmChainSelector],
				p.User[admin.PublicKey()],
			)

			raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(config.CcipLogicReceiver))
			rmI, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment)
			require.NotNil(t, res)
			require.Contains(t, strings.Join(res.Meta.LogMessages, "\n"), "Called `ccip_token_release_mint`")
		})
	})
}
