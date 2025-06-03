package contracts

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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

	// use the real program bindings, although interacting with the mock contract

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/burnmint_token_pool"
	cctp_message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_message_transmitter"
	message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_message_transmitter"
	cctp_token_messenger_minter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_token_messenger_minter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_invalid_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_token_pool"
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
	allowedOfframpSvmPDA, err := state.FindAllowedOfframpPDA(config.SvmChainSelector, dumbRamp, dumbRamp)
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
			t.Skip() // TODO

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
		t.Skip() // TODO

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
		t.Parallel()

		type MessageTransmitterPDAs struct {
			program,
			authorityPda,
			messageTransmitter,
			eventAuthority solana.PublicKey
		}

		getMessageTransmitterPDAs := func() MessageTransmitterPDAs {
			t.Helper()

			messageTransmitterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, config.CctpMessageTransmitter)
			require.NoError(t, err)
			eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, config.CctpMessageTransmitter)
			require.NoError(t, err)
			authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter_authority"), config.CctpTokenMessengerMinter.Bytes()}, config.CctpMessageTransmitter)
			require.NoError(t, err)

			return MessageTransmitterPDAs{
				program:            config.CctpMessageTransmitter,
				messageTransmitter: messageTransmitterPDA,
				authorityPda:       authorityPda,
				eventAuthority:     eventAuthority,
			}
		}

		type TokenMessengerMinterPDAs struct {
			program,
			authorityPda,
			tokenMessenger,
			remoteTokenMessenger,
			tokenMinter,
			localToken,
			tokenPair,
			custodyTokenAccount,
			eventAuthority solana.PublicKey
		}

		getTokenMessengerMinterPDAs := func(domain uint32, usdcMint solana.PublicKey) TokenMessengerMinterPDAs {
			t.Helper()

			authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("sender_authority")}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			tokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			remoteTokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), []byte(common.NumToStr(domain))}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			tokenMinterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			custodyTokenAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("custody"), usdcMint.Bytes()}, config.CctpTokenMessengerMinter)
			require.NoError(t, err)
			tokenPair, _, err := solana.FindProgramAddress([][]byte{[]byte("token_pair"), []byte(common.NumToStr(domain)), usdcMint[:]}, config.CctpTokenMessengerMinter) // faking that solana is again the remote domain
			require.NoError(t, err)
			localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcMint.Bytes()}, config.CctpTokenMessengerMinter)

			return TokenMessengerMinterPDAs{
				program:              config.CctpTokenMessengerMinter,
				authorityPda:         authorityPda,
				tokenMessenger:       tokenMessengerPDA,
				remoteTokenMessenger: remoteTokenMessengerPDA,
				tokenMinter:          tokenMinterPDA,
				eventAuthority:       eventAuthority,
				localToken:           localToken,
				tokenPair:            tokenPair,
				custodyTokenAccount:  custodyTokenAccount,
			}
		}

		type CctpTokenPoolPDAs struct {
			program,
			state,
			signer,
			tokenAccount,
			svmChainConfig solana.PublicKey
		}

		getCctpTokenPoolPDAs := func(usdcMint solana.PublicKey) CctpTokenPoolPDAs {
			t.Helper()

			statePDA, err := tokens.TokenPoolConfigAddress(usdcMint, config.CctpTokenPoolProgram)
			require.NoError(t, err)
			signerPDA, err := tokens.TokenPoolSignerAddress(usdcMint, config.CctpTokenPoolProgram)
			require.NoError(t, err)
			poolTokenAccount, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcMint, signerPDA)
			require.NoError(t, err)
			chainConfigPDA, _, err := tokens.TokenPoolChainConfigPDA(config.SvmChainSelector, usdcMint, config.CctpTokenPoolProgram)
			require.NoError(t, err)

			return CctpTokenPoolPDAs{
				program:        config.CctpTokenPoolProgram,
				state:          statePDA,
				signer:         signerPDA,
				tokenAccount:   poolTokenAccount,
				svmChainConfig: chainConfigPDA,
			}
		}

		usdcMintPriv, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)
		usdcMint := usdcMintPriv.PublicKey()
		usdcDecimals := uint8(6)

		// Example domain, used as both local and remote domain, to test something similar to a
		// Solana <> Solana transfer. It has to be a valid domain due to used_nonces PDA derivation.
		// As of May 2025, valid domains are 0 to 10 (https://developers.circle.com/stablecoins/supported-domains)
		domain := uint32(5)

		messageTransmitter := getMessageTransmitterPDAs()
		tokenMessengerMinter := getTokenMessengerMinterPDAs(domain, usdcMint)
		cctpPool := getCctpTokenPoolPDAs(usdcMint)

		cctp_message_transmitter.SetProgramID(messageTransmitter.program)
		cctp_token_messenger_minter.SetProgramID(tokenMessengerMinter.program)
		cctp_token_pool.SetProgramID(config.CctpTokenPoolProgram)

		attesters, otherKeys, _ := testutils.GenerateSignersAndTransmitters(t, 1)
		attester := attesters[0]
		attesterSolana := solana.PublicKey(common.ToLeftPadded32Bytes(attester.Address[:]))
		tokenController := otherKeys[0]

		var userATA solana.PublicKey
		var adminATA solana.PublicKey

		var tpLookupTable map[solana.PublicKey]solana.PublicKeySlice

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
					tpLookupTableAddr, err := common.CreateLookupTable(ctx, solanaGoClient, admin)
					require.NoError(t, err)

					entries := solana.PublicKeySlice{
						tpLookupTableAddr,
						solana.SystemProgramID, // placeholder for the token admin registry, as there is none in the dumb ramp
						cctpPool.program,
						cctpPool.state,
						cctpPool.tokenAccount,
						cctpPool.signer,
						solana.TokenProgramID,
						usdcMint,
						solana.SystemProgramID, // placeholder for the fee token config, as there is none in the dumb ramp
						solana.SystemProgramID, // placeholder for the router signer, as there is none in the dumb ramp
						// -- CCTP custom entries --
						messageTransmitter.messageTransmitter,
						messageTransmitter.eventAuthority,
						messageTransmitter.authorityPda,
						messageTransmitter.program,
						tokenMessengerMinter.authorityPda,
						tokenMessengerMinter.tokenMessenger,
						tokenMessengerMinter.tokenMinter,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.custodyTokenAccount,
						tokenMessengerMinter.localToken,
						tokenMessengerMinter.program,
					}

					tpLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
						tpLookupTableAddr: entries,
					}

					require.NoError(t, common.ExtendLookupTable(ctx, solanaGoClient, tpLookupTableAddr, admin, entries))
					common.AwaitSlotChange(ctx, solanaGoClient)
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

					poolInitI, err := cctp_token_pool.NewInitializeInstruction(
						dumbRamp,
						config.RMNRemoteProgram,
						cctpPool.state,
						usdcMint,
						admin.PublicKey(),
						solana.SystemProgramID,
						cctpPool.program,
						programData.Address,
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
						cctpPool.state,
						cctpPool.svmChainConfig,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					ixAppend, err := cctp_token_pool.NewAppendRemotePoolAddressesInstruction(
						config.SvmChainSelector,
						usdcMint,
						[]cctp_token_pool.RemoteAddress{{Address: cctpPool.program.Bytes()}}, // TODO check if this should be program address or pool state PDA
						cctpPool.state,
						cctpPool.svmChainConfig,
						admin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					// create pool token account
					createP, poolTokenAccount, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, usdcMint, cctpPool.signer, admin.PublicKey())
					require.NoError(t, err)
					require.Equal(t, poolTokenAccount, cctpPool.tokenAccount)

					// submit tx with all instructions
					res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{poolInitI, ixConfigure, ixAppend, createP}, admin, config.DefaultCommitment)
					require.NotNil(t, res)

					// validate state
					var configAccount burnmint_token_pool.State
					require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, cctpPool.state, config.DefaultCommitment, &configAccount))
					require.Equal(t, cctpPool.tokenAccount, configAccount.Config.PoolTokenAccount)

					// validate events
					// TODO check this is correct
					eventConfigured := tokens.EventChainConfigured{}
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainConfigured", &eventConfigured, config.PrintEvents))
					require.Equal(t, config.SvmChainSelector, eventConfigured.ChainSelector)
					require.Equal(t, 0, len(eventConfigured.PoolAddresses))
					require.Equal(t, 0, len(eventConfigured.PreviousPoolAddresses))
					require.Equal(t, cctp_token_pool.RemoteAddress{Address: usdcMint.Bytes()}, eventConfigured.Token)
					require.Equal(t, 0, len(eventConfigured.PreviousToken.Address))
					require.Equal(t, usdcMint, eventConfigured.Mint)

					eventAppended := tokens.EventRemotePoolsAppended{}
					require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemotePoolsAppended", &eventAppended, config.PrintEvents))
					require.Equal(t, config.SvmChainSelector, eventAppended.ChainSelector)
					require.Equal(t, []cctp_token_pool.RemoteAddress{{Address: cctpPool.program.Bytes()}}, eventAppended.PoolAddresses)
					require.Equal(t, 0, len(eventAppended.PreviousPoolAddresses))
					require.Equal(t, usdcMint, eventAppended.Mint)
				})
			})

			t.Run("CCTP Message Transmitter program", func(t *testing.T) {
				t.Parallel()

				// get program data account
				data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, messageTransmitter.program, &rpc.GetAccountInfoOpts{
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
					messageTransmitter.messageTransmitter,
					programData.Address,
					messageTransmitter.program,
					solana.SystemProgramID,
					messageTransmitter.eventAuthority,
					messageTransmitter.program,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
			})

			t.Run("CCTP Token Messenger Minter program", func(t *testing.T) {
				t.Parallel()

				t.Run("Initialize", func(t *testing.T) {
					// get program data account
					data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, tokenMessengerMinter.program, &rpc.GetAccountInfoOpts{
						Commitment: config.DefaultCommitment,
					})
					require.NoError(t, err)
					// Decode program data
					var programData ProgramData
					require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

					ix, err := cctp_token_messenger_minter.NewInitializeInstruction(
						cctp_token_messenger_minter.InitializeParams{
							TokenController:         tokenController.PublicKey(),
							LocalMessageTransmitter: messageTransmitter.program,
							MessageBodyVersion:      0,
						},
						admin.PublicKey(),
						admin.PublicKey(),
						tokenMessengerMinter.authorityPda,
						tokenMessengerMinter.tokenMessenger,
						tokenMessengerMinter.tokenMinter,
						programData.Address,
						tokenMessengerMinter.program,
						solana.SystemProgramID,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				t.Run("Add remote token messenger", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewAddRemoteTokenMessengerInstruction(
						cctp_token_messenger_minter.AddRemoteTokenMessengerParams{
							Domain:         domain,
							TokenMessenger: tokenMessengerMinter.program,
						},
						admin.PublicKey(),
						admin.PublicKey(),
						tokenMessengerMinter.tokenMessenger,
						tokenMessengerMinter.remoteTokenMessenger,
						solana.SystemProgramID,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)
				})

				t.Run("Add local token", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewAddLocalTokenInstruction(
						cctp_token_messenger_minter.AddLocalTokenParams{},
						admin.PublicKey(),
						tokenController.PublicKey(),
						tokenMessengerMinter.tokenMinter,
						tokenMessengerMinter.localToken,
						tokenMessengerMinter.custodyTokenAccount,
						usdcMint,
						solana.TokenProgramID,
						solana.SystemProgramID,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(tokenController))
				})

				t.Run("Link token pair", func(t *testing.T) {
					ix, err := cctp_token_messenger_minter.NewLinkTokenPairInstruction(
						cctp_token_messenger_minter.LinkTokenPairParams{
							LocalToken:   tokenMessengerMinter.localToken,
							RemoteDomain: domain,
							RemoteToken:  usdcMint, // using Solana as remote domain, so using USDC mint as remote token
						},
						admin.PublicKey(),
						tokenController.PublicKey(),
						tokenMessengerMinter.tokenMinter,
						tokenMessengerMinter.tokenPair,
						solana.SystemProgramID,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.program,
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
						tokenMessengerMinter.tokenMinter,
						tokenMessengerMinter.localToken,
						tokenMessengerMinter.eventAuthority,
						tokenMessengerMinter.program,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(tokenController))
				})

			})
		})

		getUsedNoncesPDA := func(t *testing.T, messageSent message_transmitter.MessageSent) solana.PublicKey {
			nonceBytes := messageSent.Message[12:20] // extract nonce from message, which is the 12th to 20th byte of the message
			nonce := binary.BigEndian.Uint64(nonceBytes)
			fmt.Println("Nonce for receiving message:", nonce)
			firstNonce := (nonce-1)/6400*6400 + 1 // round down to the first nonce that is a multiple of 6400
			fmt.Println("First nonce:", firstNonce)
			firstNoncePda, _, err := solana.FindProgramAddress(
				[][]byte{
					[]byte("used_nonces"),
					[]byte(common.NumToStr(domain)),
					[]byte(common.NumToStr(firstNonce)),
				},
				messageTransmitter.program,
			)
			require.NoError(t, err)
			return firstNoncePda
		}

		t.Run("CCTP sanity checks", func(t *testing.T) {
			messageAmount := uint64(1123456) // 1.123456 USDC, as it uses 6 decimals

			var messageSent message_transmitter.MessageSent

			t.Run("DepositForBurnWithCaller", func(t *testing.T) {
				messageSentEventKeypair, err := solana.NewRandomPrivateKey()

				ix, err := cctp_token_messenger_minter.NewDepositForBurnWithCallerInstruction(
					cctp_token_messenger_minter.DepositForBurnWithCallerParams{
						Amount:            messageAmount,
						DestinationDomain: domain,
						MintRecipient:     adminATA,
						DestinationCaller: admin.PublicKey(),
					},
					user.PublicKey(),
					user.PublicKey(),
					tokenMessengerMinter.authorityPda,
					userATA,
					messageTransmitter.messageTransmitter,
					tokenMessengerMinter.tokenMessenger,
					tokenMessengerMinter.remoteTokenMessenger,
					tokenMessengerMinter.tokenMinter,
					tokenMessengerMinter.localToken,
					usdcMint,
					messageSentEventKeypair.PublicKey(),
					messageTransmitter.program,
					tokenMessengerMinter.program,
					solana.TokenProgramID,
					solana.SystemProgramID,
					tokenMessengerMinter.eventAuthority,
					tokenMessengerMinter.program,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, common.AddSigners(messageSentEventKeypair))

				returnedNonce, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, tokenMessengerMinter.program.String(), binary.LittleEndian.Uint64)
				require.NoError(t, err)
				fmt.Println("Nonce:", returnedNonce)

				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, messageSentEventKeypair.PublicKey(), config.DefaultCommitment, &messageSent))
				fmt.Println("Message Sent Event Account Data:", messageSent)
				fmt.Println("Message Sent Event Bytes (hex):", hex.EncodeToString(messageSent.Message))
			})

			t.Run("ReceiveMessage", func(t *testing.T) {
				t.Run("Fund custody token account", func(t *testing.T) {
					ix, err := tokens.MintTo(messageAmount, solana.TokenProgramID, usdcMint, tokenMessengerMinter.custodyTokenAccount, admin.PublicKey())
					require.NoError(t, err)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					decimals, amount, err := tokens.TokenBalance(ctx, solanaGoClient, tokenMessengerMinter.custodyTokenAccount, config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, messageAmount, uint64(amount))
					require.Equal(t, usdcDecimals, decimals)
				})

				t.Run("Actually receive message", func(t *testing.T) {
					_, initial, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
					require.NoError(t, err)

					usedNoncesPDA := getUsedNoncesPDA(t, messageSent)

					attestation, err := ccip.AttestCCTP(messageSent.Message, attesters)
					fmt.Println("Attestation:", hex.EncodeToString(attestation))
					require.NoError(t, err)

					raw := message_transmitter.NewReceiveMessageInstruction(
						message_transmitter.ReceiveMessageParams{
							Message:     messageSent.Message,
							Attestation: attestation,
						},
						admin.PublicKey(), // payer
						admin.PublicKey(), // caller
						messageTransmitter.authorityPda,
						messageTransmitter.messageTransmitter,
						usedNoncesPDA,
						tokenMessengerMinter.program, // receiver
						solana.SystemProgramID,
						messageTransmitter.eventAuthority,
						messageTransmitter.program,
					)

					raw.AccountMetaSlice = append(raw.AccountMetaSlice,
						solana.Meta(tokenMessengerMinter.tokenMessenger),
						solana.Meta(tokenMessengerMinter.remoteTokenMessenger),
						solana.Meta(tokenMessengerMinter.tokenMinter).WRITE(),
						solana.Meta(tokenMessengerMinter.localToken).WRITE(),
						solana.Meta(tokenMessengerMinter.tokenPair),
						solana.Meta(adminATA).WRITE(), // user receiving the USDC
						solana.Meta(tokenMessengerMinter.custodyTokenAccount).WRITE(),
						solana.Meta(solana.TokenProgramID),
						solana.Meta(tokenMessengerMinter.eventAuthority),
						solana.Meta(tokenMessengerMinter.program))

					ix, err := raw.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, admin, config.DefaultCommitment)

					_, final, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)
					require.NoError(t, err)

					require.Equal(t, uint64(initial)+messageAmount, uint64(final), "Admin should have received the USDC after receiving the message")
					fmt.Println("Final balance of admin's USDC ATA:", final)
				})
			})
		})

		t.Run("TypeVersion", func(t *testing.T) {
			t.Parallel()

			ix, err := cctp_token_pool.NewTypeVersionInstruction(solana.SysVarClockPubkey).ValidateAndBuild()
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
			fakeCcipNonce := uint64(1234567890)

			messageSentEventAddress, _, err := solana.FindProgramAddress([][]byte{
				[]byte("cctp_message_sent_event"),
				user.PublicKey().Bytes(),
				common.Uint64ToLE(remoteChainSelector),
				common.Uint64ToLE(fakeCcipNonce),
			}, cctpPool.program)
			require.NoError(t, err)

			messageAmount := uint64(1e5) // 0.1 USDC, as it uses 6 decimals
			var messageSentEventData message_transmitter.MessageSent

			t.Run("Basic onramp", func(t *testing.T) {
				// fund pool signer with SOL so it can pay for the rent of the message sent event account
				fundPoolSignerIx, err := tokens.NativeTransfer(1*solana.LAMPORTS_PER_SOL, admin.PublicKey(), cctpPool.signer)
				require.NoError(t, err)

				// fund pool signer with USDC, as the onramp would normally do it but we're using a mock ramp here
				transferI, err := tokens.TokenTransferChecked(messageAmount, usdcDecimals, solana.TokenProgramID, adminATA, usdcMint, cctpPool.tokenAccount, admin.PublicKey(), solana.PublicKeySlice{})
				require.NoError(t, err)

				raw := test_ccip_invalid_receiver.NewPoolProxyLockOrBurnInstruction(
					test_ccip_invalid_receiver.LockOrBurnInV1{
						LocalToken:          usdcMint,
						Amount:              messageAmount,
						RemoteChainSelector: config.SvmChainSelector,
						Receiver:            adminATA.Bytes(), // TODO currently pools expect a receiver here then used to derive an ATA, not the ATA directly. But CCTP requires the ATA (somewhere at least). On EVM->SVM, this will require a change in the EVM contracts...
						OriginalSender:      user.PublicKey(),
						MsgNonce:            fakeCcipNonce,
					},
					cctpPool.program,
					dumbRampCctpSigner,
					cctpPool.state,
					solana.TokenProgramID,
					usdcMint,
					cctpPool.signer,
					cctpPool.tokenAccount,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
					cctpPool.svmChainConfig,
				)

				raw.AccountMetaSlice = append(raw.AccountMetaSlice,
					solana.Meta(tokenMessengerMinter.authorityPda),
					solana.Meta(messageTransmitter.messageTransmitter).WRITE(),
					solana.Meta(tokenMessengerMinter.tokenMessenger),
					solana.Meta(tokenMessengerMinter.tokenMinter),
					solana.Meta(tokenMessengerMinter.localToken).WRITE(),
					solana.Meta(messageTransmitter.program),
					solana.Meta(tokenMessengerMinter.program),
					solana.Meta(solana.SystemProgramID),
					solana.Meta(tokenMessengerMinter.eventAuthority),
					solana.Meta(tokenMessengerMinter.remoteTokenMessenger),
					solana.Meta(messageSentEventAddress).WRITE(),
				)

				ix, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				res := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferI, fundPoolSignerIx, ix}, admin, config.DefaultCommitment)
				require.NotNil(t, res)

				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, messageSentEventAddress, config.DefaultCommitment, &messageSentEventData))
				fmt.Println("Message Sent Event Data:", messageSentEventData)
				require.Equal(t, cctpPool.signer, messageSentEventData.RentPayer)

				messageSentEventAccInfo, err := solanaGoClient.GetAccountInfoWithOpts(ctx, messageSentEventAddress, &rpc.GetAccountInfoOpts{Commitment: config.DefaultCommitment})
				fmt.Println("Message Sent Event Account Info:", messageSentEventAccInfo)
				fmt.Println("Message Sent Event Account Info (Owner):", messageSentEventAccInfo.Value.Owner)
			})

			t.Run("Basic offramp", func(t *testing.T) {
				fundCustodyIx, err := tokens.MintTo(messageAmount, solana.TokenProgramID, usdcMint, tokenMessengerMinter.custodyTokenAccount, admin.PublicKey())
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{fundCustodyIx}, admin, config.DefaultCommitment)

				decimals, amount, err := tokens.TokenBalance(ctx, solanaGoClient, tokenMessengerMinter.custodyTokenAccount, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, messageAmount, uint64(amount))
				require.Equal(t, usdcDecimals, decimals)

				_, initial, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)

				attestation, err := ccip.AttestCCTP(messageSentEventData.Message, attesters)
				offchainTokenData := cctp_token_pool.MessageAndAttestation{
					Message:     cctp_token_pool.CctpMessage{Data: messageSentEventData.Message},
					Attestation: attestation,
				}
				offchainTokenDataBuffer := new(bytes.Buffer)
				require.NoError(t, offchainTokenData.MarshalWithEncoder(bin.NewBorshEncoder(offchainTokenDataBuffer)))

				raw := test_ccip_invalid_receiver.NewPoolProxyReleaseOrMintInstruction(
					test_ccip_invalid_receiver.ReleaseOrMintInV1{
						OriginalSender:      user.PublicKey().Bytes(),
						RemoteChainSelector: config.SvmChainSelector,
						Receiver:            admin.PublicKey(),
						Amount:              tokens.ToLittleEndianU256(messageAmount),
						LocalToken:          usdcMint,
						SourcePoolAddress:   cctpPool.program.Bytes(), // TODO should be the state PDA?
						SourcePoolData:      []byte{},
						OffchainTokenData:   offchainTokenDataBuffer.Bytes(),
					},
					cctpPool.program,
					dumbRampCctpSigner,
					dumbRamp,
					allowedOfframpSvmPDA,
					cctpPool.state,
					solana.TokenProgramID,
					usdcMint,
					cctpPool.signer,
					cctpPool.tokenAccount,
					cctpPool.svmChainConfig,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
					adminATA,
				)

				raw.GetPoolSignerAccount().WRITE()

				raw.AccountMetaSlice = append(raw.AccountMetaSlice,
					solana.Meta(messageTransmitter.authorityPda),
					solana.Meta(messageTransmitter.messageTransmitter),
					solana.Meta(tokenMessengerMinter.program),
					solana.Meta(solana.SystemProgramID),
					solana.Meta(messageTransmitter.eventAuthority),
					solana.Meta(messageTransmitter.program),
					solana.Meta(tokenMessengerMinter.tokenMessenger),
					solana.Meta(tokenMessengerMinter.tokenMinter).WRITE(),
					solana.Meta(tokenMessengerMinter.localToken).WRITE(),
					solana.Meta(tokenMessengerMinter.custodyTokenAccount).WRITE(),
					solana.Meta(tokenMessengerMinter.eventAuthority),
					solana.Meta(tokenMessengerMinter.remoteTokenMessenger),
					solana.Meta(tokenMessengerMinter.tokenPair),
					solana.Meta(getUsedNoncesPDA(t, messageSentEventData)).WRITE(),
				)

				ix, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				res := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{fundCustodyIx, ix}, admin, config.DefaultCommitment, tpLookupTable)
				require.NotNil(t, res)

				_, final, err := tokens.TokenBalance(ctx, solanaGoClient, adminATA, config.DefaultCommitment)

				require.Equal(t, uint64(initial)+messageAmount, uint64(final), "Admin should have received the USDC after receiving the message")
			})
		})
	})
}
