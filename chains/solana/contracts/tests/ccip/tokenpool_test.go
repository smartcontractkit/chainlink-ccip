package contracts

import (
	"encoding/binary"
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
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestTokenPool(t *testing.T) {
	t.Parallel()

	rmn_remote.SetProgramID(config.RMNRemoteProgram)
	test_token_pool.SetProgramID(config.CcipTokenPoolProgram)

	// acting as program wrapped by a token pool
	test_ccip_receiver.SetProgramID(config.CcipLogicReceiver)

	// acting as "dumb" offramp that just proxies the pool,
	// required for authnz in the pool but we don't want to test offramp internals here
	test_ccip_invalid_receiver.SetProgramID(config.CcipInvalidReceiverProgram)

	dumbRamp := config.CcipInvalidReceiverProgram
	dumbRampPoolSigner, _, _ := state.FindExternalTokenPoolsSignerPDA(config.CcipTokenPoolProgram, dumbRamp)

	admin := solana.MustPrivateKeyFromBase58("4whgxZhpxcArYWzM1iTmokruAzws9YVi2f9M7pWwchQniaFXBr1WGSGXgadeqHtiRooxNiPosdLj2g2ohbtkWtu5")
	anotherAdmin := solana.MustPrivateKeyFromBase58("3T3TwgX851KixNcZ2nJDftvXb5gMckHg8zzonKXpSthdDQEtttP4iY5VwetMRmoJkMDUPqSbrypFHVV2aC9FWHKE")
	user := solana.MustPrivateKeyFromBase58("iVu6VhFc44zTZx6LceYup18rQfxqFvKWFWzJtrNpFcWHCNnbmDfSMYFbDrj9R7RZon4t6YHHmFU8Fh6461PBvfC")

	ctx := tests.Context(t)

	allowedOfframpPDA, err := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbRamp, dumbRamp)
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)
	getBalance := func(account solana.PublicKey) string {
		balanceRes, err := solanaGoClient.GetTokenAccountBalance(ctx, account, config.DefaultCommitment)
		require.NoError(t, err)
		return balanceRes.Value.Amount
	}

	remotePool := test_token_pool.RemoteAddress{Address: []byte{1, 2, 3}}
	remoteToken := test_token_pool.RemoteAddress{Address: []byte{4, 5, 6}}

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anotherAdmin, user}, solanaGoClient, t)
	})

	t.Run("setup:allowed offramp", func(t *testing.T) {
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

	t.Run("setup: RMN Remote", func(t *testing.T) {
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

			for _, poolType := range []test_token_pool.PoolType{test_token_pool.LockAndRelease_PoolType, test_token_pool.BurnAndMint_PoolType} {
				mintPriv, err := solana.NewRandomPrivateKey()
				require.NoError(t, err)
				p, err := tokens.NewTokenPool(v.tokenProgram, config.CcipTokenPoolProgram, mintPriv.PublicKey())
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
						type ProgramData struct {
							DataType uint32
							Address  solana.PublicKey
						}
						// get program data account
						data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipTokenPoolProgram, &rpc.GetAccountInfoOpts{
							Commitment: config.DefaultCommitment,
						})
						require.NoError(t, err)
						// Decode program data
						var programData ProgramData
						require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

						poolInitI, err := test_token_pool.NewInitializeInstruction(
							poolType,
							dumbRamp,
							config.RMNRemoteProgram,
							poolConfig,
							mint,
							admin.PublicKey(),
							solana.SystemProgramID,
							config.CcipTokenPoolProgram,
							programData.Address,
						).ValidateAndBuild()

						require.NoError(t, err)

						// make pool mint_authority for token (required for burn/mint)
						authI, err := tokens.SetTokenMintAuthority(v.tokenProgram, poolSigner, mint, admin.PublicKey())
						require.NoError(t, err)

						// set rate limit
						ixRates, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint, test_token_pool.RateLimitConfig{
							Enabled:  true,
							Capacity: amount,
							Rate:     1, // slow refill
						}, test_token_pool.RateLimitConfig{
							Enabled:  false,
							Capacity: 0,
							Rate:     0,
						}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						// set pool config
						ixConfigure, err := test_token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, p.Mint, test_token_pool.RemoteConfig{
							TokenAddress: remoteToken,
							Decimals:     9,
						}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						ixAppend, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(
							config.EvmChainSelector, p.Mint, []test_token_pool.RemoteAddress{remotePool}, poolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{poolInitI, createI, authI, ixConfigure, ixAppend, ixRates}, admin, config.DefaultCommitment)
						require.NotNil(t, res)

						var configAccount test_token_pool.State
						require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, poolConfig, config.DefaultCommitment, &configAccount))
						require.Equal(t, poolTokenAccount, configAccount.Config.PoolTokenAccount)

						eventConfigured := tokens.EventChainConfigured{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainConfigured", &eventConfigured, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventConfigured.ChainSelector)
						require.Equal(t, 0, len(eventConfigured.PoolAddresses))
						require.Equal(t, 0, len(eventConfigured.PreviousPoolAddresses))
						require.Equal(t, remoteToken, eventConfigured.Token)
						require.Equal(t, 0, len(eventConfigured.PreviousToken.Address))
						require.Equal(t, mint, eventConfigured.Mint)

						eventAppended := tokens.EventRemotePoolsAppended{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemotePoolsAppended", &eventAppended, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventAppended.ChainSelector)
						require.Equal(t, []test_token_pool.RemoteAddress{remotePool}, eventAppended.PoolAddresses)
						require.Equal(t, 0, len(eventAppended.PreviousPoolAddresses))
						require.Equal(t, mint, eventAppended.Mint)

						eventRateLimit := tokens.EventRateLimitConfigured{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RateLimitConfigured", &eventRateLimit, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventRateLimit.ChainSelector)
						require.Equal(t, true, eventRateLimit.InboundRateLimit.Enabled)
						require.Equal(t, amount, eventRateLimit.InboundRateLimit.Capacity)
						require.Equal(t, uint64(1), eventRateLimit.InboundRateLimit.Rate)
						require.Equal(t, false, eventRateLimit.OutboundRateLimit.Enabled)
						require.Equal(t, uint64(0), eventRateLimit.OutboundRateLimit.Capacity)
						require.Equal(t, uint64(0), eventRateLimit.OutboundRateLimit.Rate)
						require.Equal(t, mint, eventRateLimit.Mint)
					})

					t.Run("admin:ownership", func(t *testing.T) {
						// successfully transfer ownership
						instruction, err := test_token_pool.NewTransferOwnershipInstruction(
							anotherAdmin.PublicKey(),
							poolConfig,
							mint,
							admin.PublicKey(),
						).ValidateAndBuild()
						require.NoError(t, err)
						result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, admin, config.DefaultCommitment)
						require.NotNil(t, result)

						// Successfully accept ownership
						// anotherAdmin becomes owner for remaining tests
						instruction, err = test_token_pool.NewAcceptOwnershipInstruction(
							poolConfig,
							mint,
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

						lbI, err := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
							test_ccip_invalid_receiver.LockOrBurnInV1{
								LocalToken:          mint,
								Amount:              amount,
								RemoteChainSelector: config.EvmChainSelector,
							},
							p.PoolProgram,
							dumbRampPoolSigner,
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

						event := tokens.EventBurnLock{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, tokens.MethodToEvent(lockOrBurn), &event))
						require.Equal(t, dumbRampPoolSigner, event.Sender)
						require.Equal(t, amount, event.Amount)
						require.Equal(t, mint, event.Mint)

						// validate balances
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))
						expectedPoolBal := uint64(0)
						if poolType == test_token_pool.LockAndRelease_PoolType {
							expectedPoolBal = amount
						}
						require.Equal(t, fmt.Sprintf("%d", expectedPoolBal), getBalance(poolTokenAccount))
					})

					t.Run(releaseOrMint, func(t *testing.T) {
						require.Equal(t, "0", getBalance(p.User[admin.PublicKey()]))

						rmI, err := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
							test_ccip_invalid_receiver.ReleaseOrMintInV1{
								LocalToken:          mint,
								SourcePoolAddress:   remotePool.Address,
								Amount:              tokens.ToLittleEndianU256(amount * 1e9), // scale to proper decimals
								Receiver:            admin.PublicKey(),
								RemoteChainSelector: config.EvmChainSelector,
							},
							config.CcipTokenPoolProgram,
							dumbRampPoolSigner,
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

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment)
						require.NotNil(t, res)

						event := tokens.EventMintRelease{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, tokens.MethodToEvent(releaseOrMint), &event))
						require.Equal(t, admin.PublicKey(), event.Recipient)
						require.Equal(t, poolSigner, event.Sender)
						require.Equal(t, amount, event.Amount)
						require.Equal(t, mint, event.Mint)

						// validate balances
						require.Equal(t, fmt.Sprintf("%d", amount), getBalance(p.User[admin.PublicKey()]))
						require.Equal(t, "0", getBalance(poolTokenAccount))
					})

					t.Run("invalid", func(t *testing.T) {
						t.Run("config", func(t *testing.T) {
							t.Parallel()

							cfgs := []struct {
								name   string
								c      test_token_pool.RateLimitConfig
								errStr string
							}{
								{
									name: "enabled-zero-rate",
									c: test_token_pool.RateLimitConfig{
										Enabled:  true,
										Capacity: 10,
										Rate:     0,
									},
									errStr: "invalid rate limit rate",
								},
								{
									name: "enabled-rate-larger-than-capacity",
									c: test_token_pool.RateLimitConfig{
										Enabled:  true,
										Capacity: 1,
										Rate:     100,
									},
									errStr: "invalid rate limit rate",
								},
								{
									name: "disabled-nonzero-rate",
									c: test_token_pool.RateLimitConfig{
										Enabled:  false,
										Capacity: 0,
										Rate:     100,
									},
									errStr: "disabled non-zero rate limit",
								},
								{
									name: "disabled-nonzero-capacity",
									c: test_token_pool.RateLimitConfig{
										Enabled:  false,
										Capacity: 10,
										Rate:     0,
									},
									errStr: "disabled non-zero rate limit",
								},
							}

							for _, cfg := range cfgs {
								t.Run(cfg.name, func(t *testing.T) {
									t.Parallel()

									ixRates, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint, cfg.c, test_token_pool.RateLimitConfig{}, poolConfig, p.Chain[config.EvmChainSelector], anotherAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
									require.NoError(t, err)

									testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ixRates}, anotherAdmin, config.DefaultCommitment, []string{cfg.errStr})
								})
							}
						})

						t.Run("globally cursed", func(t *testing.T) {
							globalCurse := rmn_remote.CurseSubject{
								Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
							}

							curseI, err := rmn_remote.NewCurseInstruction(
								globalCurse,
								config.RMNRemoteConfigPDA,
								admin.PublicKey(),
								config.RMNRemoteCursesPDA,
								solana.SystemProgramID,
							).ValidateAndBuild()
							require.NoError(t, err)

							// Do not alter global state, so don't just submit the curse instruction in a tx that succeeds,
							// as that may break parallel tests. Instead, submit the curse ix together with the pool ix
							// that fails, which reverts the entire tx and does not affect other tests.

							lbI, err := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
								test_ccip_invalid_receiver.LockOrBurnInV1{
									LocalToken:          mint,
									Amount:              amount,
									RemoteChainSelector: config.EvmChainSelector,
								},
								p.PoolProgram,
								dumbRampPoolSigner,
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

							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{curseI, lbI}, admin, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})

							rmI, err := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
								test_ccip_invalid_receiver.ReleaseOrMintInV1{
									LocalToken:          mint,
									SourcePoolAddress:   remotePool.Address,
									Amount:              tokens.ToLittleEndianU256(amount * 1e9), // scale to proper decimals
									Receiver:            admin.PublicKey(),
									RemoteChainSelector: config.EvmChainSelector,
								},
								config.CcipTokenPoolProgram,
								dumbRampPoolSigner,
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

							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{curseI, rmI}, admin, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
						})

						t.Run("subject cursed", func(t *testing.T) {
							evmCurse := rmn_remote.CurseSubject{}
							binary.LittleEndian.PutUint64(evmCurse.Value[:], config.EvmChainSelector)

							curseI, err := rmn_remote.NewCurseInstruction(
								evmCurse,
								config.RMNRemoteConfigPDA,
								admin.PublicKey(),
								config.RMNRemoteCursesPDA,
								solana.SystemProgramID,
							).ValidateAndBuild()
							require.NoError(t, err)

							// Do not alter global state, so don't just submit the curse instruction in a tx that succeeds,
							// as that may break parallel tests. Instead, submit the curse ix together with the pool ix
							// that fails, which reverts the entire tx and does not affect other tests.

							lbI, err := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
								test_ccip_invalid_receiver.LockOrBurnInV1{
									LocalToken:          mint,
									Amount:              amount,
									RemoteChainSelector: config.EvmChainSelector,
								},
								p.PoolProgram,
								dumbRampPoolSigner,
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

							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{curseI, lbI}, admin, config.DefaultCommitment, []string{"Error Code: SubjectCursed"})

							rmI, err := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
								test_ccip_invalid_receiver.ReleaseOrMintInV1{
									LocalToken:          mint,
									SourcePoolAddress:   remotePool.Address,
									Amount:              tokens.ToLittleEndianU256(amount * 1e9), // scale to proper decimals
									Receiver:            admin.PublicKey(),
									RemoteChainSelector: config.EvmChainSelector,
								},
								config.CcipTokenPoolProgram,
								dumbRampPoolSigner,
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

							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{curseI, rmI}, admin, config.DefaultCommitment, []string{"Error Code: SubjectCursed"})
						})

						t.Run("exceed-rate-limit", func(t *testing.T) {
							t.Parallel()

							// exceed capacity of bucket
							rmI, err := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
								test_ccip_invalid_receiver.ReleaseOrMintInV1{
									LocalToken:          mint,
									SourcePoolAddress:   remotePool.Address,
									Amount:              tokens.ToLittleEndianU256(amount * amount * 1e9),
									Receiver:            admin.PublicKey(),
									RemoteChainSelector: config.EvmChainSelector,
								},
								config.CcipTokenPoolProgram,
								dumbRampPoolSigner,
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
							rmI, err = test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
								test_ccip_invalid_receiver.ReleaseOrMintInV1{
									LocalToken:          mint,
									SourcePoolAddress:   remotePool.Address,
									Amount:              tokens.ToLittleEndianU256(amount * 1e9),
									Receiver:            admin.PublicKey(),
									RemoteChainSelector: config.EvmChainSelector,
								},
								config.CcipTokenPoolProgram,
								dumbRampPoolSigner,
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
							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{rmI, rmI}, user, config.DefaultCommitment, []string{"rate limit reached"})

							// pool should refill automatically, but slowly
							// small amount should pass
							time.Sleep(time.Second) // wait for refill
							rmI, err = test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
								test_ccip_invalid_receiver.ReleaseOrMintInV1{
									LocalToken:          mint,
									SourcePoolAddress:   remotePool.Address,
									Amount:              tokens.ToLittleEndianU256(1e9),
									Receiver:            admin.PublicKey(),
									RemoteChainSelector: config.EvmChainSelector,
								},
								config.CcipTokenPoolProgram,
								dumbRampPoolSigner,
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
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rmI}, user, config.DefaultCommitment)
						})
					})

					t.Run("closing", func(t *testing.T) {
						ixDelete, err := test_token_pool.NewDeleteChainConfigInstruction(config.EvmChainSelector, mint, poolConfig, p.Chain[config.EvmChainSelector], anotherAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
						require.NoError(t, err)

						ixRouterChange, err := test_token_pool.NewSetRouterInstruction(config.CcipRouterProgram, poolConfig, mint, anotherAdmin.PublicKey()).ValidateAndBuild()
						require.NoError(t, err)

						res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixDelete, ixRouterChange}, anotherAdmin, config.DefaultCommitment)
						require.NotNil(t, res)

						eventDelete := tokens.EventChainRemoved{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainRemoved", &eventDelete, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, eventDelete.ChainSelector)
						require.Equal(t, mint, eventDelete.Mint)

						eventRouter := tokens.EventRouterUpdated{}
						require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RouterUpdated", &eventRouter, config.PrintEvents))
						require.Equal(t, config.CcipRouterProgram, eventRouter.NewRouter)
						require.Equal(t, dumbRamp, eventRouter.OldRouter)
						require.Equal(t, mint, eventRouter.Mint)
					})
				})
			}
		})
	}

	// test functionality with arbitrary wrapped program
	t.Run("Wrapped", func(t *testing.T) {
		t.Parallel()
		mintPriv := solana.MustPrivateKeyFromBase58("5PMQ49JibQPVBFneTzixstoS2z888CoUgej1PoYgvmKXZcKw4b3Zd8vhCjKQcUDSjLnR9M1tUrzCCXLrPBZoqjJm")
		p, err := tokens.NewTokenPool(solana.TokenProgramID, config.CcipTokenPoolProgram, mintPriv.PublicKey())
		require.NoError(t, err)
		mint := p.Mint

		t.Run("setup:pool", func(t *testing.T) {
			var err error
			p.PoolConfig, err = tokens.TokenPoolConfigAddress(mint, config.CcipTokenPoolProgram)
			require.NoError(t, err)
			p.PoolSigner, err = tokens.TokenPoolSignerAddress(mint, config.CcipTokenPoolProgram)
			require.NoError(t, err)

			// create token
			instructions, err := tokens.CreateToken(ctx, solana.TokenProgramID, mint, admin.PublicKey(), 0, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, err)

			type ProgramData struct {
				DataType uint32
				Address  solana.PublicKey
			}
			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipTokenPoolProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)
			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			// create pool
			poolInitI, err := test_token_pool.NewInitializeInstruction(
				test_token_pool.Wrapped_PoolType,
				dumbRamp,
				config.RMNRemoteProgram,
				p.PoolConfig,
				mint,
				admin.PublicKey(),
				solana.SystemProgramID,
				config.CcipTokenPoolProgram,
				programData.Address,
			).ValidateAndBuild()
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
			configureI, err := test_token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, p.Mint, test_token_pool.RemoteConfig{
				TokenAddress: remoteToken,
			}, p.PoolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			// append remote pools
			appendI, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(
				config.EvmChainSelector, p.Mint, []test_token_pool.RemoteAddress{remotePool}, p.PoolConfig, p.Chain[config.EvmChainSelector], admin.PublicKey(), solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			res := testutils.SendAndConfirm(ctx, t, solanaGoClient, append(instructions, poolInitI, createR, createP, configureI, appendI), admin, config.DefaultCommitment, common.AddSigners(mintPriv))
			require.NotNil(t, res)
		})

		t.Run("burnOrLock", func(t *testing.T) {
			raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
				test_ccip_invalid_receiver.LockOrBurnInV1{LocalToken: mint, RemoteChainSelector: config.EvmChainSelector},
				p.PoolProgram,
				dumbRampPoolSigner,
				p.PoolConfig,
				solana.TokenProgramID,
				mint,
				p.PoolSigner,
				p.PoolTokenAccount,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
				p.Chain[config.EvmChainSelector],
			)
			raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.NewAccountMeta(config.CcipLogicReceiver, false, false))
			lbI, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{lbI}, admin, config.DefaultCommitment)
			require.NotNil(t, res)
			require.Contains(t, strings.Join(res.Meta.LogMessages, "\n"), "Called `ccip_token_lock_burn`")
		})

		t.Run("mintOrRelease", func(t *testing.T) {
			raw := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
				test_ccip_invalid_receiver.ReleaseOrMintInV1{
					LocalToken:          mint,
					SourcePoolAddress:   remotePool.Address,
					Receiver:            admin.PublicKey(),
					RemoteChainSelector: config.EvmChainSelector,
					Amount:              tokens.ToLittleEndianU256(1),
				},
				config.CcipTokenPoolProgram,
				dumbRampPoolSigner,
				dumbRamp,
				allowedOfframpPDA,
				p.PoolConfig,
				solana.TokenProgramID,
				mint,
				p.PoolSigner,
				p.PoolTokenAccount,
				p.Chain[config.EvmChainSelector],
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
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
