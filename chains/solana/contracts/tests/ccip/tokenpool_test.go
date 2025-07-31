package contracts

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"maps"
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

	// use the real program bindings, although interacting with the mock contract

	cctp_message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_message_transmitter"
	message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_message_transmitter"
	cctp_token_messenger_minter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_token_messenger_minter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/cctp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestTokenPool(t *testing.T) {
	t.Parallel()

	rmn_remote.SetProgramID(config.RMNRemoteProgram)
	test_token_pool.SetProgramID(config.CcipTokenPoolProgram)
	cctp_message_transmitter.SetProgramID(config.CctpMessageTransmitter)

	// acting as program wrapped by a token pool
	test_ccip_receiver.SetProgramID(config.CcipLogicReceiver)

	// acting as "dumb" onramp & "dump" offramp that just proxies the pool,
	// required for authnz in the pool but we don't want to test ramp internals here
	test_ccip_invalid_receiver.SetProgramID(config.CcipInvalidReceiverProgram)

	dumbRamp := config.CcipInvalidReceiverProgram
	dumbRampPoolSigner, _, _ := state.FindExternalTokenPoolsSignerPDA(config.CcipTokenPoolProgram, dumbRamp)

	admin := solana.MustPrivateKeyFromBase58("4whgxZhpxcArYWzM1iTmokruAzws9YVi2f9M7pWwchQniaFXBr1WGSGXgadeqHtiRooxNiPosdLj2g2ohbtkWtu5")
	anotherAdmin := solana.MustPrivateKeyFromBase58("3T3TwgX851KixNcZ2nJDftvXb5gMckHg8zzonKXpSthdDQEtttP4iY5VwetMRmoJkMDUPqSbrypFHVV2aC9FWHKE")
	user := solana.MustPrivateKeyFromBase58("iVu6VhFc44zTZx6LceYup18rQfxqFvKWFWzJtrNpFcWHCNnbmDfSMYFbDrj9R7RZon4t6YHHmFU8Fh6461PBvfC")

	ctx := tests.Context(t)

	allowedOfframpEvmPDA, err := state.FindAllowedOfframpPDA(config.EvmChainSelector, dumbRamp, dumbRamp)
	require.NoError(t, err)
	allowedOfframpSvmPDA, err := state.FindAllowedOfframpPDA(config.SvmChainSelector, dumbRamp, dumbRamp)
	require.NoError(t, err)

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, admin)
	getBalance := func(account solana.PublicKey) string {
		balanceRes, err := solanaGoClient.GetTokenAccountBalance(ctx, account, config.DefaultCommitment)
		require.NoError(t, err)
		return balanceRes.Value.Amount
	}

	var offrampLookupTable map[solana.PublicKey]solana.PublicKeySlice

	remotePool := test_token_pool.RemoteAddress{Address: []byte{1, 2, 3}}
	remoteToken := test_token_pool.RemoteAddress{Address: []byte{4, 5, 6}}

	t.Run("setup:funding", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{admin, anotherAdmin, user}, solanaGoClient, t)
	})

	t.Run("setup:allowed offramp", func(t *testing.T) {
		evmIx, err := test_ccip_invalid_receiver.NewAddOfframpInstruction(
			config.EvmChainSelector,
			dumbRamp,
			allowedOfframpEvmPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		svmIx, err := test_ccip_invalid_receiver.NewAddOfframpInstruction(
			config.SvmChainSelector,
			dumbRamp,
			allowedOfframpSvmPDA,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{evmIx, svmIx}, admin, config.DefaultCommitment)
	})

	t.Run("setup: (dumb) offramp lookup table", func(t *testing.T) {
		offrampLookupTableAddr, kErr := common.CreateLookupTable(ctx, solanaGoClient, admin)
		require.NoError(t, kErr)

		entries := solana.PublicKeySlice{
			dumbRamp,
			config.RMNRemoteProgram,
			config.RMNRemoteConfigPDA,
			config.RMNRemoteCursesPDA,
			solana.SystemProgramID,
			// hack just for the test to have it here - it shouldn't be in the table, but our tests are also
			// cursing in the same tx (to avoid concurrency side-effects) and thus need the extra space
			allowedOfframpSvmPDA,
		}

		offrampLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
			offrampLookupTableAddr: entries,
		}

		require.NoError(t, common.ExtendLookupTable(ctx, solanaGoClient, offrampLookupTableAddr, admin, entries))
		require.NoError(t, common.AwaitSlotChange(ctx, solanaGoClient))
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

						raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
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
						)
						lbI, err := raw.ValidateAndBuild()
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
							allowedOfframpEvmPDA,
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

						t.Run("Rate Limit Admin", func(t *testing.T) {
							// set invalid a new rate limit admin
							ixRateAdmin, err := test_token_pool.NewSetRateLimitAdminInstruction(
								p.Mint,
								user.PublicKey(),
								poolConfig,
								anotherAdmin.PublicKey(),
							).ValidateAndBuild()
							require.NoError(t, err)

							// test new rate limit admin
							ixRatesValid, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint,
								test_token_pool.RateLimitConfig{
									Enabled:  true,
									Capacity: amount,
									Rate:     1,
								}, test_token_pool.RateLimitConfig{
									Enabled:  false,
									Capacity: 0,
									Rate:     0,
								}, poolConfig, p.Chain[config.EvmChainSelector], user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
							require.NoError(t, err)

							// undo rate limit admin
							ixRateAdmin2, err := test_token_pool.NewSetRateLimitAdminInstruction(
								p.Mint,
								anotherAdmin.PublicKey(),
								poolConfig,
								anotherAdmin.PublicKey(),
							).ValidateAndBuild()
							require.NoError(t, err)

							// try to modify rate limit with invalid admin
							ixRates, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, p.Mint,
								test_token_pool.RateLimitConfig{
									Enabled:  true,
									Capacity: amount,
									Rate:     1,
								}, test_token_pool.RateLimitConfig{
									Enabled:  false,
									Capacity: 0,
									Rate:     0,
								}, poolConfig, p.Chain[config.EvmChainSelector], user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
							require.NoError(t, err)

							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ixRateAdmin, ixRatesValid, ixRateAdmin2, ixRates}, user, config.DefaultCommitment, []string{"SetChainRateLimit"}, common.AddSigners(anotherAdmin))
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

							raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
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
							)
							lbI, err := raw.ValidateAndBuild()
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
								allowedOfframpEvmPDA,
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

							raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
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
							)
							lbI, err := raw.ValidateAndBuild()
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
								allowedOfframpEvmPDA,
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
								allowedOfframpEvmPDA,
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
								allowedOfframpEvmPDA,
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
								allowedOfframpEvmPDA,
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
				allowedOfframpEvmPDA,
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

	t.Run("CCTP token pool", func(t *testing.T) {
		// do not run in parallel, as it checks the balances of the same users that are also used in the other suites

		usdcMintPriv, kErr := solana.NewRandomPrivateKey()
		require.NoError(t, kErr)
		usdcMint := usdcMintPriv.PublicKey()
		usdcDecimals := uint8(6)

		// Example domain, used as both local and remote domain, to test something similar to a
		// Solana <> Solana transfer. It has to be a valid domain due to used_nonces PDA derivation.
		// As of May 2025, valid domains are 0 to 10 (https://developers.circle.com/stablecoins/supported-domains)
		domain := uint32(5)

		messageTransmitter := cctp.GetMessageTransmitterPDAs(t)
		tokenMessengerMinter := cctp.GetTokenMessengerMinterPDAs(t, domain, usdcMint, usdcMint[:])
		cctpPool := cctp.GetTokenPoolPDAs(t, usdcMint)

		cctp_message_transmitter.SetProgramID(messageTransmitter.Program)
		cctp_token_messenger_minter.SetProgramID(tokenMessengerMinter.Program)
		cctp_token_pool.SetProgramID(config.CctpTokenPoolProgram)

		attesters, otherKeys, _ := testutils.GenerateSignersAndTransmitters(t, 1)
		attester := attesters[0]
		attesterSolana := solana.PublicKey(common.ToLeftPadded32Bytes(attester.Address[:]))
		tokenController := otherKeys[0]

		var userATA solana.PublicKey
		var adminATA solana.PublicKey

		var tpLookupTable map[solana.PublicKey]solana.PublicKeySlice
		var tpLookupTableAddr solana.PublicKey

		t.Run("Setup", func(t *testing.T) {
			type ProgramData struct {
				DataType uint32
				Address  solana.PublicKey
			}

			t.Run("CCTP token", func(t *testing.T) {
				// create USDC token
				instructions, err := tokens.CreateToken(ctx, solana.TokenProgramID, usdcMint, admin.PublicKey(), usdcDecimals, solanaGoClient, config.DefaultCommitment)
				require.NoError(t, err)

				// Admin: create associated token accounts & mint
				createAdminI, adminTokenAccount, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, usdcMint, admin.PublicKey(), admin.PublicKey())
				require.NoError(t, err)
				adminATA = adminTokenAccount
				mintToAdminI, err := tokens.MintTo(1000*1e6, solana.TokenProgramID, usdcMint, adminTokenAccount, admin.PublicKey())
				require.NoError(t, err)

				// User: create associated token account & mint
				createUserI, userTokenAccount, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, usdcMint, user.PublicKey(), admin.PublicKey())
				require.NoError(t, err)
				userATA = userTokenAccount
				mintToUserI, err := tokens.MintTo(100*1e6, solana.TokenProgramID, usdcMint, userTokenAccount, admin.PublicKey())
				require.NoError(t, err)

				instructions = append(instructions, createAdminI, mintToAdminI, createUserI, mintToUserI)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, instructions, admin, config.DefaultCommitment, common.AddSigners(usdcMintPriv))

				// validate admin for example
				outDec, outVal, err := tokens.TokenBalance(ctx, solanaGoClient, adminTokenAccount, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, int(1_000*1e6), outVal)
				require.Equal(t, usdcDecimals, outDec)
			})

			t.Run("CCTP token pool program", func(t *testing.T) {
				t.Parallel()

				t.Run("Create lookup table", func(t *testing.T) {
					tpLookupTableAddr, kErr = common.CreateLookupTable(ctx, solanaGoClient, admin)
					require.NoError(t, kErr)

					entries := solana.PublicKeySlice{
						tpLookupTableAddr,
						solana.SystemProgramID, // placeholder for the token admin registry, as there is none in the dumb ramp
						cctpPool.Program,
						cctpPool.State,
						cctpPool.TokenAccount,
						cctpPool.Signer,
						solana.TokenProgramID,
						usdcMint,
						solana.SystemProgramID, // placeholder for the fee token config, as there is none in the dumb ramp
						solana.SystemProgramID, // placeholder for the router signer, as there is none in the dumb ramp
						// -- CCTP custom entries --
						messageTransmitter.MessageTransmitter,
						tokenMessengerMinter.Program,
						solana.SystemProgramID,
						messageTransmitter.Program,
						tokenMessengerMinter.TokenMessenger,
						tokenMessengerMinter.TokenMinter,
						tokenMessengerMinter.LocalToken,
						tokenMessengerMinter.EventAuthority,
					}

					tpLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
						tpLookupTableAddr: entries,
					}

					require.NoError(t, common.ExtendLookupTable(ctx, solanaGoClient, tpLookupTableAddr, admin, entries))
					require.NoError(t, common.AwaitSlotChange(ctx, solanaGoClient))
				})

				t.Run("Initialize", func(t *testing.T) {
					// get program data account
					data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CctpTokenPoolProgram, &rpc.GetAccountInfoOpts{
						Commitment: config.DefaultCommitment,
					})
					require.NoError(t, err)
					// Decode program data
					var programData ProgramData
					require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

					poolGlobalInitI, err := cctp_token_pool.NewInitGlobalConfigInstruction(
						cctpPool.GlobalConfig,
						admin.PublicKey(),
						solana.SystemProgramID,
						cctpPool.Program,
						programData.Address,
					).ValidateAndBuild()
					require.NoError(t, err)

					poolInitI, err := cctp_token_pool.NewInitializeInstruction(
						dumbRamp,
						config.RMNRemoteProgram,
						cctpPool.State,
						usdcMint,
						admin.PublicKey(),
						solana.SystemProgramID,
						cctpPool.Program,
						programData.Address,
						cctpPool.GlobalConfig,
					).ValidateAndBuild()

					require.NoError(t, err)

					// set pool config
					ixConfigure, err := cctp_token_pool.NewInitChainRemoteConfigInstruction(
						config.SvmChainSelector,
						usdcMint,
						cctp_token_pool.RemoteConfig{
							TokenAddress: cctp_token_pool.RemoteAddress{
								Address: usdcMint.Bytes(),
							},
							Decimals:      usdcDecimals,
							PoolAddresses: []cctp_token_pool.RemoteAddress{},
						},
						cctpPool.State,
						cctpPool.SvmChainConfig,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					ixCctpConfigure, err := cctp_token_pool.NewEditChainRemoteConfigCctpInstruction(
						config.SvmChainSelector,
						usdcMint,
						cctp_token_pool.CctpChain{
							DomainId:          domain,
							DestinationCaller: cctpPool.Signer, // as it is svm<->svm, the caller is the pool signer
						},
						cctpPool.State,
						cctpPool.SvmChainConfig,
						admin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					ixAppend, err := cctp_token_pool.NewAppendRemotePoolAddressesInstruction(
						config.SvmChainSelector,
						usdcMint,
						[]cctp_token_pool.RemoteAddress{{Address: cctpPool.Signer.Bytes()}}, // when the source is Solana, the pool is identified by its signer
						cctpPool.State,
						cctpPool.SvmChainConfig,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// create pool token account
					createP, poolTokenAccount, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, usdcMint, cctpPool.Signer, admin.PublicKey())
					require.NoError(t, err)
					require.Equal(t, poolTokenAccount, cctpPool.TokenAccount)

					// submit tx with all instructions
					res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{poolGlobalInitI, poolInitI, ixConfigure, ixCctpConfigure, ixAppend, createP}, admin, config.DefaultCommitment)
					require.NotNil(t, res)

					// validate state
					var configAccount cctp_token_pool.State
					require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, cctpPool.State, config.DefaultCommitment, &configAccount))
					require.Equal(t, cctpPool.TokenAccount, configAccount.Config.PoolTokenAccount)

					// validate events
					var eventConfigured tokens.EventChainConfigured
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainConfigured", &eventConfigured, config.PrintEvents))
					require.Equal(t, config.SvmChainSelector, eventConfigured.ChainSelector)
					require.Equal(t, 0, len(eventConfigured.PoolAddresses))
					require.Equal(t, 0, len(eventConfigured.PreviousPoolAddresses))
					require.Equal(t, cctp_token_pool.RemoteAddress{Address: usdcMint.Bytes()}, eventConfigured.Token)
					require.Equal(t, 0, len(eventConfigured.PreviousToken.Address))
					require.Equal(t, usdcMint, eventConfigured.Mint)

					var eventAppended tokens.EventRemotePoolsAppended
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemotePoolsAppended", &eventAppended, config.PrintEvents))
					require.Equal(t, config.SvmChainSelector, eventAppended.ChainSelector)
					require.Equal(t, []cctp_token_pool.RemoteAddress{{Address: cctpPool.Signer.Bytes()}}, eventAppended.PoolAddresses)
					require.Equal(t, 0, len(eventAppended.PreviousPoolAddresses))
					require.Equal(t, usdcMint, eventAppended.Mint)

					var eventCctpEdit tokens.EventRemoteChainCctpConfigEdited
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainCctpConfigChanged", &eventCctpEdit, config.PrintEvents))
					require.Equal(t, domain, eventCctpEdit.Config.DomainId)
					require.Equal(t, cctpPool.Signer, eventCctpEdit.Config.DestinationCaller)
				})
			})

			t.Run("CCTP Message Transmitter program", func(t *testing.T) {
				t.Parallel()

				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, messageTransmitter.Program, &rpc.GetAccountInfoOpts{
					Commitment: config.DefaultCommitment,
				})
				require.NoError(t, err)
				// Decode program data
				var programData ProgramData
				require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

				ix, err := cctp_message_transmitter.NewInitializeInstruction(
					cctp_message_transmitter.InitializeParams{
						LocalDomain:        domain,
						Attester:           attesterSolana,
						MaxMessageBodySize: 1234,
						Version:            0,
					},
					admin.PublicKey(),
					admin.PublicKey(),
					messageTransmitter.MessageTransmitter,
					programData.Address,
					messageTransmitter.Program,
					solana.SystemProgramID,
					messageTransmitter.EventAuthority,
					messageTransmitter.Program,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			})

			t.Run("CCTP Token Messenger Minter program", func(t *testing.T) {
				t.Parallel()

				t.Run("Initialize", func(t *testing.T) {
					// get program data account
					data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, tokenMessengerMinter.Program, &rpc.GetAccountInfoOpts{
						Commitment: config.DefaultCommitment,
					})
					require.NoError(t, err)
					// Decode program data
					var programData ProgramData
					require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

					ix, err := cctp_token_messenger_minter.NewInitializeInstruction(
						cctp_token_messenger_minter.InitializeParams{
							TokenController:         tokenController.PublicKey(),
							LocalMessageTransmitter: messageTransmitter.Program,
							MessageBodyVersion:      0,
						},
						admin.PublicKey(),
						admin.PublicKey(),
						tokenMessengerMinter.AuthorityPda,
						tokenMessengerMinter.TokenMessenger,
						tokenMessengerMinter.TokenMinter,
						programData.Address,
						tokenMessengerMinter.Program,
						solana.SystemProgramID,
						tokenMessengerMinter.EventAuthority,
						tokenMessengerMinter.Program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				t.Run("Add remote token messenger", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewAddRemoteTokenMessengerInstruction(
						cctp_token_messenger_minter.AddRemoteTokenMessengerParams{
							Domain:         domain,
							TokenMessenger: tokenMessengerMinter.Program,
						},
						admin.PublicKey(),
						admin.PublicKey(),
						tokenMessengerMinter.TokenMessenger,
						tokenMessengerMinter.RemoteTokenMessenger,
						solana.SystemProgramID,
						tokenMessengerMinter.EventAuthority,
						tokenMessengerMinter.Program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				t.Run("Add local token", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewAddLocalTokenInstruction(
						cctp_token_messenger_minter.AddLocalTokenParams{},
						admin.PublicKey(),
						tokenController.PublicKey(),
						tokenMessengerMinter.TokenMinter,
						tokenMessengerMinter.LocalToken,
						tokenMessengerMinter.CustodyTokenAccount,
						usdcMint,
						solana.TokenProgramID,
						solana.SystemProgramID,
						tokenMessengerMinter.EventAuthority,
						tokenMessengerMinter.Program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(tokenController))
				})

				t.Run("Link token pair", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewLinkTokenPairInstruction(
						cctp_token_messenger_minter.LinkTokenPairParams{
							LocalToken:   tokenMessengerMinter.LocalToken,
							RemoteDomain: domain,
							RemoteToken:  usdcMint, // using Solana as remote domain, so using USDC mint as remote token
						},
						admin.PublicKey(),
						tokenController.PublicKey(),
						tokenMessengerMinter.TokenMinter,
						tokenMessengerMinter.TokenPair,
						solana.SystemProgramID,
						tokenMessengerMinter.EventAuthority,
						tokenMessengerMinter.Program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(tokenController))
				})

				t.Run("Set max burn amount per message", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewSetMaxBurnAmountPerMessageInstruction(
						cctp_token_messenger_minter.SetMaxBurnAmountPerMessageParams{
							BurnLimitPerMessage: 2 * 1e6, // 2 USDC, as it uses 6 decimals
						},
						tokenController.PublicKey(),
						tokenMessengerMinter.TokenMinter,
						tokenMessengerMinter.LocalToken,
						tokenMessengerMinter.EventAuthority,
						tokenMessengerMinter.Program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(tokenController))
				})
			})
		})

		t.Run("CCTP sanity checks", func(t *testing.T) {
			messageAmount := uint64(1123456) // 1.123456 USDC, as it uses 6 decimals

			var messageSent message_transmitter.MessageSent
			messageSentEventKeypair, err := solana.NewRandomPrivateKey()
			require.NoError(t, err)

			var attestation []byte

			t.Run("DepositForBurnWithCaller", func(t *testing.T) {
				ix, err := cctp_token_messenger_minter.NewDepositForBurnWithCallerInstruction(
					cctp_token_messenger_minter.DepositForBurnWithCallerParams{
						Amount:            messageAmount,
						DestinationDomain: domain,
						MintRecipient:     adminATA,
						DestinationCaller: admin.PublicKey(),
					},
					user.PublicKey(),
					user.PublicKey(),
					tokenMessengerMinter.AuthorityPda,
					userATA,
					messageTransmitter.MessageTransmitter,
					tokenMessengerMinter.TokenMessenger,
					tokenMessengerMinter.RemoteTokenMessenger,
					tokenMessengerMinter.TokenMinter,
					tokenMessengerMinter.LocalToken,
					usdcMint,
					messageSentEventKeypair.PublicKey(),
					messageTransmitter.Program,
					tokenMessengerMinter.Program,
					solana.TokenProgramID,
					solana.SystemProgramID,
					tokenMessengerMinter.EventAuthority,
					tokenMessengerMinter.Program,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, common.AddSigners(messageSentEventKeypair))

				returnedNonce, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, tokenMessengerMinter.Program.String(), binary.LittleEndian.Uint64)
				require.NoError(t, err)
				fmt.Println("Nonce:", returnedNonce)

				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, messageSentEventKeypair.PublicKey(), config.DefaultCommitment, &messageSent))
				fmt.Println("Message Sent Event Account Data:", messageSent)
				fmt.Println("Message Sent Event Bytes (hex):", hex.EncodeToString(messageSent.Message))
			})

			t.Run("ReceiveMessage", func(t *testing.T) {
				t.Run("Fund custody token account", func(t *testing.T) {
					ix, err := tokens.MintTo(messageAmount, solana.TokenProgramID, usdcMint, tokenMessengerMinter.CustodyTokenAccount, admin.PublicKey())
					require.NoError(t, err)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					decimals, amount, err := tokens.TokenBalance(ctx, solanaGoClient, tokenMessengerMinter.CustodyTokenAccount, config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, messageAmount, uint64(amount))
					require.Equal(t, usdcDecimals, decimals)
				})

				t.Run("Actually receive message", func(t *testing.T) {
					_, initial, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
					require.NoError(t, err)

					usedNoncesPDA := cctp.GetUsedNoncesPDA(t, messageSent.Message)

					attestation, err = ccip.AttestCCTP(messageSent.Message, attesters)
					fmt.Println("Attestation:", hex.EncodeToString(attestation))
					require.NoError(t, err)

					raw := message_transmitter.NewReceiveMessageInstruction(
						message_transmitter.ReceiveMessageParams{
							Message:     messageSent.Message,
							Attestation: attestation,
						},
						admin.PublicKey(), // payer
						admin.PublicKey(), // caller
						messageTransmitter.AuthorityPda,
						messageTransmitter.MessageTransmitter,
						usedNoncesPDA,
						tokenMessengerMinter.Program, // receiver
						solana.SystemProgramID,
						messageTransmitter.EventAuthority,
						messageTransmitter.Program,
					)

					raw.AccountMetaSlice = append(raw.AccountMetaSlice,
						solana.Meta(tokenMessengerMinter.TokenMessenger),
						solana.Meta(tokenMessengerMinter.RemoteTokenMessenger),
						solana.Meta(tokenMessengerMinter.TokenMinter).WRITE(),
						solana.Meta(tokenMessengerMinter.LocalToken).WRITE(),
						solana.Meta(tokenMessengerMinter.TokenPair),
						solana.Meta(adminATA).WRITE(), // user receiving the USDC
						solana.Meta(tokenMessengerMinter.CustodyTokenAccount).WRITE(),
						solana.Meta(solana.TokenProgramID),
						solana.Meta(tokenMessengerMinter.EventAuthority),
						solana.Meta(tokenMessengerMinter.Program))

					ix, err := raw.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					_, final, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
					require.NoError(t, err)

					require.Equal(t, uint64(initial)+messageAmount, uint64(final), "Admin should have received the USDC after receiving the message")
					fmt.Println("Final balance of admin's USDC ATA:", final)
				})
			})

			t.Run("Reclaim event account rent", func(t *testing.T) {
				initial, err := solanaGoClient.GetAccountInfoWithOpts(ctx, user.PublicKey(), &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)

				messageEventAccount, err := solanaGoClient.GetAccountInfoWithOpts(ctx, messageSentEventKeypair.PublicKey(), &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)

				reclaimIx, err := cctp_message_transmitter.NewReclaimEventAccountInstruction(
					cctp_message_transmitter.ReclaimEventAccountParams{
						Attestation: attestation,
					},
					user.PublicKey(),
					messageTransmitter.MessageTransmitter,
					messageSentEventKeypair.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{reclaimIx}, user, config.DefaultCommitment)
				require.NotNil(t, result)

				final, err := solanaGoClient.GetAccountInfoWithOpts(ctx, user.PublicKey(), &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)

				require.Equal(t, initial.Value.Lamports-result.Meta.Fee+messageEventAccount.Value.Lamports, final.Value.Lamports)
			})
		})

		t.Run("TypeVersion", func(t *testing.T) {
			t.Parallel()

			ix, err := cctp_token_pool.NewTypeVersionInstruction().ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			require.NotNil(t, result)

			output, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, config.CctpTokenPoolProgram.String(), func(b []byte) string {
				require.Len(t, b, int(binary.LittleEndian.Uint32(b[:4]))+4) // the first 4 bytes just encodes the length
				return string(b[4:])
			})
			require.NoError(t, err)
			// regex from https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
			semverRegex := "(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"
			require.Regexp(t, fmt.Sprintf("^%s %s$", "cctp-token-pool", semverRegex), output)
			fmt.Printf("Type Version: %s\n", output)
		})

		t.Run("E2E through pool (svm user -> svm admin)", func(t *testing.T) {
			t.Parallel()

			dumbRampCctpSigner, _, err := state.FindExternalTokenPoolsSignerPDA(config.CctpTokenPoolProgram, dumbRamp)
			require.NoError(t, err)

			remoteChainSelector := config.SvmChainSelector
			fakeCcipTotalNonce := uint64(1234567890)

			messageSentEventAddress, _, err := solana.FindProgramAddress([][]byte{
				[]byte("ccip_cctp_message_sent_event"),
				user.PublicKey().Bytes(),
				common.Uint64ToLE(remoteChainSelector),
				common.Uint64ToLE(fakeCcipTotalNonce),
			}, cctpPool.Program)
			require.NoError(t, err)

			messageAmount := uint64(1e5) // 0.1 USDC, as it uses 6 decimals
			var messageSentEventData message_transmitter.MessageSent
			var attestation []byte

			t.Run("Basic onramp", func(t *testing.T) {
				// fund pool signer with SOL so it can pay for the rent of the message sent event account
				fundPoolSignerIx, err := tokens.NativeTransfer(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), cctpPool.Signer)
				require.NoError(t, err)

				// fund pool signer with USDC, as the onramp would normally do it but we're using a mock ramp here
				transferI, err := tokens.TokenTransferChecked(messageAmount, usdcDecimals, solana.TokenProgramID, adminATA, usdcMint, cctpPool.TokenAccount, admin.PublicKey(), solana.PublicKeySlice{})
				require.NoError(t, err)

				lockOrBurnIn := cctp_token_pool.LockOrBurnInV1{
					LocalToken:          usdcMint,
					Amount:              messageAmount,
					RemoteChainSelector: config.SvmChainSelector,
					Receiver:            cctpPool.TokenAccount.Bytes(),
					OriginalSender:      user.PublicKey(),
					MsgTotalNonce:       fakeCcipTotalNonce,
				}

				dynamicAdditionalAccountMetas := []*solana.AccountMeta{
					solana.Meta(tokenMessengerMinter.AuthorityPda),
					solana.Meta(tokenMessengerMinter.RemoteTokenMessenger),
					solana.Meta(messageSentEventAddress).WRITE(),
				}

				additionalAccountMetas := []*solana.AccountMeta{
					// static ones, present in LUT
					solana.Meta(messageTransmitter.MessageTransmitter).WRITE(),
					solana.Meta(tokenMessengerMinter.Program),
					solana.Meta(solana.SystemProgramID),
					solana.Meta(messageTransmitter.Program),
					solana.Meta(tokenMessengerMinter.TokenMessenger),
					solana.Meta(tokenMessengerMinter.TokenMinter),
					solana.Meta(tokenMessengerMinter.LocalToken).WRITE(),
					solana.Meta(tokenMessengerMinter.EventAuthority),
				}
				additionalAccountMetas = append(additionalAccountMetas, dynamicAdditionalAccountMetas...)

				t.Run("Accounts derivation", func(t *testing.T) {
					accounts, tables := deriveCctpIxAccounts(ctx, t, solanaGoClient, admin, func(stage string, askWith []*solana.AccountMeta) RawIx {
						raw := cctp_token_pool.NewDeriveAccountsLockOrBurnTokensInstruction(
							stage,
							lockOrBurnIn,
						)
						raw.AccountMetaSlice = append(raw.AccountMetaSlice, askWith...)
						return raw
					})
					require.Equal(t, []solana.PublicKey{}, tables)
					require.Equal(t, dynamicAdditionalAccountMetas, accounts)
				})

				raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
					test_ccip_invalid_receiver.LockOrBurnInV1(lockOrBurnIn),
					cctpPool.Program,
					dumbRampCctpSigner,
					cctpPool.State,
					solana.TokenProgramID,
					usdcMint,
					cctpPool.Signer,
					cctpPool.TokenAccount,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
					cctpPool.SvmChainConfig,
				)

				raw.AccountMetaSlice = append(raw.AccountMetaSlice, additionalAccountMetas...)

				ix, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				t.Run("When there is a global curse, it fails", func(t *testing.T) {
					globalCurse := rmn_remote.CurseSubject{
						Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
					}
					curseIx, err := rmn_remote.NewCurseInstruction(
						globalCurse,
						config.RMNRemoteConfigPDA,
						admin.PublicKey(),
						config.RMNRemoteCursesPDA,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// submit curse and onramp in the same transaction, so there are no side-effects to other tests
					// as the tx is atomic and reverts
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{fundPoolSignerIx, transferI, curseIx, ix},
						admin, config.DefaultCommitment, []string{ccip.GloballyCursed_RmnRemoteError.String()})
				})

				t.Run("When there is a lane curse, it fails", func(t *testing.T) {
					svmCurse := rmn_remote.CurseSubject{}
					binary.LittleEndian.PutUint64(svmCurse.Value[:], config.SvmChainSelector)
					curseIx, err := rmn_remote.NewCurseInstruction(
						svmCurse,
						config.RMNRemoteConfigPDA,
						admin.PublicKey(),
						config.RMNRemoteCursesPDA,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// submit curse and onramp in the same transaction, so there are no side-effects to other tests
					// as the tx is atomic and reverts
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{fundPoolSignerIx, transferI, curseIx, ix},
						admin, config.DefaultCommitment, []string{ccip.SubjectCursed_RmnRemoteError.String()})
				})

				t.Run("When there is no curse, it succeeds", func(t *testing.T) {
					res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI, fundPoolSignerIx, ix}, admin, config.DefaultCommitment)
					require.NotNil(t, res)

					output, err := common.ExtractAnchorTypedReturnValue[cctp_token_pool.LockOrBurnOutV1](ctx, res.Meta.LogMessages, cctpPool.Program.String())
					require.NoError(t, err)
					outputSourceDomain := binary.BigEndian.Uint32(output.DestPoolData[60:64])
					require.Equal(t, outputSourceDomain, domain) // the source domain is Solana in this test

					require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, messageSentEventAddress, config.DefaultCommitment, &messageSentEventData))
					fmt.Println("Message Sent Event Data:", messageSentEventData)
					require.Equal(t, cctpPool.Signer, messageSentEventData.RentPayer)

					var ccipCctpMessageSentEvent ccip.EventCcipCctpMessageSent
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "CcipCctpMessageSentEvent", &ccipCctpMessageSentEvent, config.PrintEvents))
					require.Equal(t, user.PublicKey(), ccipCctpMessageSentEvent.OriginalSender)
					require.Equal(t, config.SvmChainSelector, ccipCctpMessageSentEvent.RemoteChainSelector)
					require.Equal(t, fakeCcipTotalNonce, ccipCctpMessageSentEvent.MsgTotalNonce)
					require.Equal(t, messageSentEventAddress, ccipCctpMessageSentEvent.EventAddress)
					require.Equal(t, messageSentEventData.Message, ccipCctpMessageSentEvent.MessageSentBytes)
					require.Equal(t, domain, ccipCctpMessageSentEvent.SourceDomain)

					// Check CCTP nonce in all ways it appears
					outputNonce := binary.BigEndian.Uint64(output.DestPoolData[24:32])                      // in pool returned dest data
					require.Equal(t, outputNonce, cctp.GetNonce(ccipCctpMessageSentEvent.MessageSentBytes)) // in event message bytes
					require.Equal(t, outputNonce, cctp.GetNonce(messageSentEventData.Message))              // in event account data
					require.Equal(t, outputNonce, ccipCctpMessageSentEvent.CctpNonce)                       // in the event field
				})
			})

			t.Run("Basic offramp", func(t *testing.T) {
				fundCustodyIx, err := tokens.MintTo(messageAmount, solana.TokenProgramID, usdcMint, tokenMessengerMinter.CustodyTokenAccount, admin.PublicKey())
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundCustodyIx}, admin, config.DefaultCommitment)

				decimals, amount, err := tokens.TokenBalance(ctx, solanaGoClient, tokenMessengerMinter.CustodyTokenAccount, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, messageAmount, uint64(amount))
				require.Equal(t, usdcDecimals, decimals)

				_, initial, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
				require.NoError(t, err)

				attestation, err = ccip.AttestCCTP(messageSentEventData.Message, attesters)
				require.NoError(t, err)
				offchainTokenData := cctp_token_pool.MessageAndAttestation{
					Message:     cctp_token_pool.CctpMessage{Data: messageSentEventData.Message},
					Attestation: attestation,
				}
				offchainTokenDataBuffer := new(bytes.Buffer)
				require.NoError(t, offchainTokenData.MarshalWithEncoder(bin.NewBorshEncoder(offchainTokenDataBuffer)))

				sourcePoolData := make([]byte, 64)
				binary.BigEndian.PutUint64(sourcePoolData[24:32], cctp.GetNonce(messageSentEventData.Message))
				binary.BigEndian.PutUint32(sourcePoolData[60:64], domain)

				releaseOrMintIn := cctp_token_pool.ReleaseOrMintInV1{
					OriginalSender:      cctp_token_pool.RemoteAddress{Address: user.PublicKey().Bytes()},
					RemoteChainSelector: config.SvmChainSelector,
					Receiver:            admin.PublicKey(),
					Amount:              tokens.ToLittleEndianU256(messageAmount),
					LocalToken:          usdcMint,
					SourcePoolAddress:   cctp_token_pool.RemoteAddress{Address: cctpPool.Signer.Bytes()}, // when the source is Solana, the pool is identified by its signer
					SourcePoolData:      sourcePoolData[:],
					OffchainTokenData:   offchainTokenDataBuffer.Bytes(),
				}

				dynamicAdditionalAccountMetas := []*solana.AccountMeta{
					// Accounts not in LUT
					solana.Meta(messageTransmitter.AuthorityPda),
					solana.Meta(messageTransmitter.EventAuthority),
					solana.Meta(tokenMessengerMinter.CustodyTokenAccount).WRITE(),
					solana.Meta(tokenMessengerMinter.RemoteTokenMessenger),
					solana.Meta(tokenMessengerMinter.TokenPair),
					solana.Meta(cctp.GetUsedNoncesPDA(t, messageSentEventData.Message)).WRITE(),
				}

				additionalAccountMetas := []*solana.AccountMeta{
					// Accounts present in LUT
					solana.Meta(messageTransmitter.MessageTransmitter),
					solana.Meta(tokenMessengerMinter.Program),
					solana.Meta(solana.SystemProgramID),
					solana.Meta(messageTransmitter.Program),
					solana.Meta(tokenMessengerMinter.TokenMessenger),
					solana.Meta(tokenMessengerMinter.TokenMinter).WRITE(),
					solana.Meta(tokenMessengerMinter.LocalToken).WRITE(),
					solana.Meta(tokenMessengerMinter.EventAuthority),
				}
				additionalAccountMetas = append(additionalAccountMetas, dynamicAdditionalAccountMetas...)

				t.Run("Accounts derivation", func(t *testing.T) {
					accounts, tables := deriveCctpIxAccounts(ctx, t, solanaGoClient, admin, func(stage string, askWith []*solana.AccountMeta) RawIx {
						raw := cctp_token_pool.NewDeriveAccountsReleaseOrMintTokensInstruction(
							stage,
							releaseOrMintIn,
						)
						raw.AccountMetaSlice = append(raw.AccountMetaSlice, askWith...)
						return raw
					})
					require.Equal(t, []solana.PublicKey{}, tables)
					require.Equal(t, dynamicAdditionalAccountMetas, accounts)
				})

				raw := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
					test_ccip_invalid_receiver.ReleaseOrMintInV1{
						OriginalSender:      releaseOrMintIn.OriginalSender.Address,
						RemoteChainSelector: releaseOrMintIn.RemoteChainSelector,
						Receiver:            releaseOrMintIn.Receiver,
						Amount:              releaseOrMintIn.Amount,
						LocalToken:          releaseOrMintIn.LocalToken,
						SourcePoolAddress:   releaseOrMintIn.SourcePoolAddress.Address,
						SourcePoolData:      releaseOrMintIn.SourcePoolData,
						OffchainTokenData:   releaseOrMintIn.OffchainTokenData,
					},
					cctpPool.Program,
					dumbRampCctpSigner,
					dumbRamp,
					allowedOfframpSvmPDA,
					cctpPool.State,
					solana.TokenProgramID,
					usdcMint,
					cctpPool.Signer,
					cctpPool.TokenAccount,
					cctpPool.SvmChainConfig,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
					adminATA,
				)

				raw.GetPoolSignerAccount().WRITE() // CCTP requires this, though other pools don't (which is why the bindings don't do it by default)

				raw.AccountMetaSlice = append(raw.AccountMetaSlice, additionalAccountMetas...)

				ix, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				// merge lookup tables into single map
				lookupTables := map[solana.PublicKey]solana.PublicKeySlice{}
				maps.Copy(lookupTables, offrampLookupTable)
				maps.Copy(lookupTables, tpLookupTable)

				t.Run("When there is a global curse, it fails", func(t *testing.T) {
					globalCurse := rmn_remote.CurseSubject{
						Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
					}
					curseIx, err := rmn_remote.NewCurseInstruction(
						globalCurse,
						config.RMNRemoteConfigPDA,
						admin.PublicKey(),
						config.RMNRemoteCursesPDA,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// submit curse and offramp in the same transaction, so there are no side-effects to other tests
					// as the tx is atomic and reverts
					testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{curseIx, ix},
						admin, config.DefaultCommitment, lookupTables, []string{ccip.GloballyCursed_RmnRemoteError.String()})
				})

				t.Run("When there is a lane curse, it fails", func(t *testing.T) {
					svmCurse := rmn_remote.CurseSubject{}
					binary.LittleEndian.PutUint64(svmCurse.Value[:], config.SvmChainSelector)
					curseIx, err := rmn_remote.NewCurseInstruction(
						svmCurse,
						config.RMNRemoteConfigPDA,
						admin.PublicKey(),
						config.RMNRemoteCursesPDA,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// submit curse and offramp in the same transaction, so there are no side-effects to other tests
					// as the tx is atomic and reverts
					testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{curseIx, ix},
						admin, config.DefaultCommitment, lookupTables, []string{ccip.SubjectCursed_RmnRemoteError.String()})
				})

				t.Run("When there is no curse, it succeeds", func(t *testing.T) {
					res := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{fundCustodyIx, ix},
						admin, config.DefaultCommitment, lookupTables)
					require.NotNil(t, res)

					_, final, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
					require.NoError(t, err)

					require.Equal(t, uint64(initial)+messageAmount, uint64(final), "Admin should have received the USDC after receiving the message")
				})
			})

			t.Run("Reclaim event account rent", func(t *testing.T) {
				poolSignerInfo, err := solanaGoClient.GetAccountInfoWithOpts(ctx, cctpPool.Signer, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				initialBalance := poolSignerInfo.Value.Lamports

				messageSentAccount, err := solanaGoClient.GetAccountInfoWithOpts(ctx, messageSentEventAddress, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				messageRent := messageSentAccount.Value.Lamports

				reclaimIx := cctp_token_pool.NewReclaimEventAccountInstruction(
					usdcMint,
					user.PublicKey(),
					remoteChainSelector,
					fakeCcipTotalNonce,
					attestation,
					cctpPool.State,
					cctpPool.Signer,
					messageSentEventAddress,
					messageTransmitter.MessageTransmitter,
					messageTransmitter.Program,
					admin.PublicKey(),
					solana.SystemProgramID,
				)

				ix, err := reclaimIx.ValidateAndBuild()
				require.NoError(t, err)

				// There's no fund manager, so funds cannot be reclaimed
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, []string{"Fund Manager is invalid"})

				fundManagerIxRaw := cctp_token_pool.NewSetFundManagerInstruction(admin.PublicKey(), cctpPool.State, usdcMint, admin.PublicKey())
				fundManagerIx, err := fundManagerIxRaw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundManagerIx, ix}, admin, config.DefaultCommitment)

				poolSignerInfoAfter, err := solanaGoClient.GetAccountInfoWithOpts(ctx, cctpPool.Signer, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				finalBalance := poolSignerInfoAfter.Value.Lamports

				require.Equal(t, finalBalance, initialBalance+messageRent)
			})

			t.Run("Reclaim funds from signer account", func(t *testing.T) {
				poolSignerInfoInitial, err := solanaGoClient.GetAccountInfoWithOpts(ctx, cctpPool.Signer, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				startingBalance := poolSignerInfoInitial.Value.Lamports
				require.LessOrEqual(t, startingBalance, 1*solana.LAMPORTS_PER_SOL)

				userAtaBalanceInitial, err := solanaGoClient.GetAccountInfoWithOpts(ctx, userATA, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				startingAtaBalance := userAtaBalanceInitial.Value.Lamports

				// Fund pool signer to exactly 1 SOL
				fundPoolSignerIx, err := tokens.NativeTransfer(1*solana.LAMPORTS_PER_SOL-startingBalance, admin.PublicKey(), cctpPool.Signer)
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundPoolSignerIx}, admin, config.DefaultCommitment)

				// Can't reclaim funds to invalid (unconfigured) receiver
				reclaimToInvalidIx, err := cctp_token_pool.NewReclaimFundsInstruction(solana.LAMPORTS_PER_SOL/5, cctpPool.State, usdcMint, cctpPool.Signer, userATA, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{reclaimToInvalidIx}, admin, config.DefaultCommitment, []string{"Invalid destination for funds reclaim"})

				// We set reclaim destination to user ATA
				reclaimDestinationConfigIx, err := cctp_token_pool.NewSetFundReclaimDestinationInstruction(userATA, cctpPool.State, usdcMint, admin.PublicKey()).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{reclaimDestinationConfigIx}, admin, config.DefaultCommitment)

				// Can't reclaim invalid values
				reclaimZeroIx, err := cctp_token_pool.NewReclaimFundsInstruction(0, cctpPool.State, usdcMint, cctpPool.Signer, userATA, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{reclaimZeroIx}, admin, config.DefaultCommitment, []string{"Invalid SOL amount"})

				reclaimExcessiveFundsIx, err := cctp_token_pool.NewReclaimFundsInstruction(2*solana.LAMPORTS_PER_SOL, cctpPool.State, usdcMint, cctpPool.Signer, userATA, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{reclaimExcessiveFundsIx}, admin, config.DefaultCommitment, []string{"Insufficient funds"})

				// reclaim limits are enforced
				setLimitIx, err := cctp_token_pool.NewSetMinimumSignerFundsInstruction(solana.LAMPORTS_PER_SOL/2, cctpPool.State, usdcMint, admin.PublicKey()).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{setLimitIx}, admin, config.DefaultCommitment)
				reclaimFundsOverLimitIx, err := cctp_token_pool.NewReclaimFundsInstruction(3*solana.LAMPORTS_PER_SOL/4, cctpPool.State, usdcMint, cctpPool.Signer, userATA, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{reclaimFundsOverLimitIx}, admin, config.DefaultCommitment, []string{"Insufficient funds"})

				// Otherwise, reclaiming succeeds
				validReclaimAmount := solana.LAMPORTS_PER_SOL / 4
				validReclaimIx, err := cctp_token_pool.NewReclaimFundsInstruction(validReclaimAmount, cctpPool.State, usdcMint, cctpPool.Signer, userATA, admin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{validReclaimIx}, admin, config.DefaultCommitment)

				poolSignerInfoAfter, err := solanaGoClient.GetAccountInfoWithOpts(ctx, cctpPool.Signer, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				finalBalance := poolSignerInfoAfter.Value.Lamports

				userAtaBalanceAfter, err := solanaGoClient.GetAccountInfoWithOpts(ctx, userATA, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				require.NoError(t, err)
				finalAtaBalance := userAtaBalanceAfter.Value.Lamports

				require.Equal(t, finalBalance, startingBalance-validReclaimAmount)
				require.Equal(t, finalAtaBalance, startingAtaBalance+validReclaimAmount)
			})
		})
	})
}

type RawIx interface {
	ValidateAndBuild() (*cctp_token_pool.Instruction, error)
}

func deriveCctpIxAccounts(
	ctx context.Context,
	t *testing.T,
	solanaGoClient *rpc.Client,
	signer solana.PrivateKey,
	createRawIx func(stage string, askWith []*solana.AccountMeta) RawIx,
) (derivedAccounts []*solana.AccountMeta, lookUpTables []solana.PublicKey) {
	t.Helper()

	derivedAccounts = make([]*solana.AccountMeta, 0)
	lookUpTables = make([]solana.PublicKey, 0)

	askWith := []*solana.AccountMeta{}
	stage := "Start"
	for {
		deriveRaw := createRawIx(stage, askWith)
		derive, err := deriveRaw.ValidateAndBuild()
		require.NoError(t, err)
		tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{derive}, signer, config.DefaultCommitment)
		derivation, err := common.ExtractAnchorTypedReturnValue[cctp_token_pool.DeriveAccountsResponse](ctx, tx.Meta.LogMessages, config.CctpTokenPoolProgram.String())
		require.NoError(t, err)

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

		if len(derivation.LookUpTablesToSave) > 0 {
			lookUpTables = append(lookUpTables, derivation.LookUpTablesToSave...)
		}

		if len(derivation.NextStage) == 0 {
			return derivedAccounts, lookUpTables
		}
		stage = derivation.NextStage
	}
}
