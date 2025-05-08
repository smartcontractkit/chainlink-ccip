package contracts

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"maps"
	"math/big"
	"sort"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_sender"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestCCIPRouter(t *testing.T) {
	t.Parallel()

	ccip_router.SetProgramID(config.CcipRouterProgram)
	test_ccip_receiver.SetProgramID(config.CcipLogicReceiver)
	test_token_pool.SetProgramID(config.CcipTokenPoolProgram)
	fee_quoter.SetProgramID(config.FeeQuoterProgram)
	ccip_offramp.SetProgramID(config.CcipOfframpProgram)
	example_ccip_sender.SetProgramID(config.CcipBaseSender)
	rmn_remote.SetProgramID(config.RMNRemoteProgram)

	ctx := tests.Context(t)
	user := solana.MustPrivateKeyFromBase58("ZZdVf32Npuhci4u4ir2NW9491Y3FTv2Gwk41HMpvgJoh81UM42LcNqAN8SXapHfPcr61QP7sJj7K2mKHt7qFCoV")
	anotherUser := solana.MustPrivateKeyFromBase58("i9btAVgpmReUv9jH52xpPoYvtsv6XQSJrRGnLpTU4ArSP6E3Xa9aunyeT7n83QhMeLZMmRPnwY41xr7jFrTXAPR")
	tokenlessUser := solana.MustPrivateKeyFromBase58("4g8xCc96ox2ksCcv5VggqaenkSsnkUEe8WZBazLyTCMFnB2r2bJtdNK2QW9E6mojMmMGpGcSGuKFeDQbGLiCqM3n")
	legacyAdmin := solana.MustPrivateKeyFromBase58("5dYf8bzhFbmwTNyq8vAoeqe19vdjCbSEKkPrN6DQoJB2RwjyuVapCmU46pEBvbw7aeg5wBncTc8HQJnKZy2LYsmF")
	ccipAdmin := solana.MustPrivateKeyFromBase58("AmNCLrssMCEGUM5RWDao2v4J6ydB48FZBDye6x1o5Z1KR7QkZvuMgX185fZoESBotBtBZ2SiQVkoNwjmUD8ysge")
	token0PoolAdmin := solana.MustPrivateKeyFromBase58("2NqkEEvMWf5Y8aUTSZvGCfyqPi4KjBKJShXWH4BrTWVyfxBzJm22S1K4gtkgzHcAhStseHypRV7mKPsx5nVa9h2e")
	token1PoolAdmin := solana.MustPrivateKeyFromBase58("rAJULkVqwXHED22STrDdUPxfGqSi6gkHfSpgaYSTzp3X9MCqsYegEcWMJVZ5yQFw5H3mNdsBAJpR9xfdYRGfb7J")
	token2PoolAdmin := solana.MustPrivateKeyFromBase58("3UUqZ5xa3xv9fX1UJQyHsJtovE2gzmJUJjybizjrAvxh7NUmVyVQHUkJVkQwBKtVr5vLVCp1DWAeAzv46WzLoEmS")
	feeAggregator := solana.MustPrivateKeyFromBase58("NPchsbT3bkkziJPBUxto3eVTVKHcV4tou33NjRY8inArmzi7EXKBf5cC7MX47xqYMwkZGbkw7t55jCciCStSwNs")

	nonceEvmPDA, gerr := state.FindNoncePDA(config.EvmChainSelector, user.PublicKey(), config.CcipRouterProgram)
	require.NoError(t, gerr)
	nonceSvmPDA, gerr := state.FindNoncePDA(config.SvmChainSelector, user.PublicKey(), config.CcipRouterProgram)
	require.NoError(t, gerr)

	// billing
	type AccountsPerToken struct {
		name             string
		program          solana.PublicKey
		mint             solana.PublicKey
		billingATA       solana.PublicKey
		userATA          solana.PublicKey
		anotherUserATA   solana.PublicKey
		tokenlessUserATA solana.PublicKey
		feeAggregatorATA solana.PublicKey

		// fee quoter PDAs
		fqBillingConfigPDA solana.PublicKey
		fqEvmConfigPDA     solana.PublicKey
		// add other accounts as needed
	}
	wsol := AccountsPerToken{name: "WSOL (pre-2022)"}
	link22 := AccountsPerToken{name: "LINK sample token (2022)"}
	billingTokens := []*AccountsPerToken{&wsol, &link22}

	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, legacyAdmin)

	evmToken0Decimals := uint8(18)
	evmToken1Decimals := uint8(18)
	evmToken2Decimals := uint8(18)
	evmToken3Decimals := uint8(18)

	token0Decimals := uint8(9)
	token1Decimals := uint8(18)
	token2Decimals := uint8(9)
	link22Decimals := uint8(9) // Solana Decimals for Link Token2022 token

	// token addresses
	// Create link22 token, managed by "legacyAdmin" (not "ccipAdmin" who manages CCIP).
	// Random-generated key, but fixing it adds determinism to tests to make it easier to debug.
	linkMintPrivK := solana.MustPrivateKeyFromBase58("32YVeJArcWWWV96fztfkRQhohyFz5Hwno93AeGVrN4g2LuFyvwznrNd9A6tbvaTU6BuyBsynwJEMLre8vSy3CrVU")

	token0Mint := solana.MustPrivateKeyFromBase58("42uJJqZk4gFz6Q6ghMiaYrFdDapXhbufQdTCGJDMeyv2wN6wNBbXkBBPibF7xQQZemzRaDH66ouJmjfvWhPJKtQC")
	token0, gerr := tokens.NewTokenPool(config.Token2022Program, config.CcipTokenPoolProgram, token0Mint.PublicKey())
	require.NoError(t, gerr)
	token1Mint := solana.MustPrivateKeyFromBase58("5uBhsiup2KiXPNznVSZaFeuegvLXCtmVmuW4u7kfBbP3xSmT1BZgmapKYiCCwi8GUgwdSQniAah3rJaUJdhPprcB")
	token1, gerr := tokens.NewTokenPool(config.Token2022Program, config.CcipTokenPoolProgram, token1Mint.PublicKey())
	require.NoError(t, gerr)
	token2Mint := solana.MustPrivateKeyFromBase58("2b4zgrXRBDkAuhFMEUEaMJHXwPfPX5REmy3gA3Vxhj7efTLkZEBuq49mDSdQyCJFyG3KPRGt2PVoGF8VqhGUTo9")
	token2, gerr := tokens.NewTokenPool(config.Token2022Program, config.CcipTokenPoolProgram, token2Mint.PublicKey())
	require.NoError(t, gerr)
	linkPool, gerr := tokens.NewTokenPool(config.Token2022Program, config.CcipTokenPoolProgram, linkMintPrivK.PublicKey())
	require.NoError(t, gerr)

	signers, transmitters, getTransmitter := testutils.GenerateSignersAndTransmitters(t, config.MaxOracles)

	signerAddresses := [][20]byte{}
	transmitterPubKeys := []solana.PublicKey{}
	for _, v := range signers {
		signerAddresses = append(signerAddresses, v.Address)
	}
	for _, v := range transmitters {
		transmitterPubKeys = append(transmitterPubKeys, v.PublicKey())
	}
	// sort to match onchain sort
	sort.SliceStable(signerAddresses, func(i, j int) bool {
		return bytes.Compare(signerAddresses[i][:], signerAddresses[j][:]) == 1
	})
	sort.SliceStable(transmitterPubKeys, func(i, j int) bool {
		return bytes.Compare(transmitterPubKeys[i].Bytes(), transmitterPubKeys[j].Bytes()) == 1
	})

	getBalance := func(tokenAccount solana.PublicKey) uint64 {
		_, amount, berr := tokens.TokenBalance(ctx, solanaGoClient, tokenAccount, config.DefaultCommitment)
		require.NoError(t, berr)
		return uint64(amount)
	}

	getFqTokenConfigPDA := func(mint solana.PublicKey) solana.PublicKey {
		tokenConfigPda, _, _ := state.FindFqBillingTokenConfigPDA(mint, config.FeeQuoterProgram)
		return tokenConfigPda
	}

	getFqPerChainPerTokenConfigBillingPDA := func(mint solana.PublicKey) solana.PublicKey {
		tokenBillingPda, _, _ := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, mint, config.FeeQuoterProgram)
		return tokenBillingPda
	}

	onRampAddress := [64]byte{1, 2, 3}

	validSourceChainConfig := ccip_offramp.SourceChainConfig{
		OnRamp: ccip_offramp.OnRampAddress{
			Bytes: onRampAddress,
			Len:   3,
		},
		IsEnabled:                 true,
		IsRmnVerificationDisabled: true,
	}
	validFqDestChainConfig := fee_quoter.DestChainConfig{
		IsEnabled: true,

		LaneCodeVersion: fee_quoter.Default_CodeVersion,

		// minimal valid config
		DefaultTxGasLimit:           200000,
		MaxPerMsgGasLimit:           3000000,
		MaxDataBytes:                30000,
		MaxNumberOfTokensPerMsg:     5,
		DefaultTokenDestGasOverhead: 50000,
		ChainFamilySelector:         [4]uint8(config.EvmChainFamilySelector),

		DefaultTokenFeeUsdcents: 50,
		NetworkFeeUsdcents:      50,
	}
	// Small enough to fit in u160, big enough to not fall in the precompile space.
	validReceiverAddress := [32]byte{}
	validReceiverAddress[12] = 1

	emptyGenericExtraArgsV2 := []byte{}

	var ccipSendLookupTable map[solana.PublicKey]solana.PublicKeySlice
	var offrampLookupTable map[solana.PublicKey]solana.PublicKeySlice

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, append(
				transmitters,
				user,
				anotherUser,
				tokenlessUser,
				legacyAdmin,
				ccipAdmin,
				token0PoolAdmin,
				token1PoolAdmin,
				token2PoolAdmin,
				feeAggregator),
				solanaGoClient,
				t)
		})

		t.Run("Type version", func(t *testing.T) {
			testcases := []struct {
				ContractName              string
				Program                   solana.PublicKey
				NewTypeVersionInstruction solana.Instruction
			}{
				{"ccip-router", config.CcipRouterProgram, ccip_router.NewTypeVersionInstruction(solana.SysVarClockPubkey).Build()},
				{"fee-quoter", config.FeeQuoterProgram, fee_quoter.NewTypeVersionInstruction(solana.SysVarClockPubkey).Build()},
				{"ccip-offramp", config.CcipOfframpProgram, ccip_offramp.NewTypeVersionInstruction(solana.SysVarClockPubkey).Build()},
				{"rmn-remote", config.RMNRemoteProgram, rmn_remote.NewTypeVersionInstruction(solana.SysVarClockPubkey).Build()},
			}
			for _, testcase := range testcases {
				t.Run(testcase.ContractName, func(t *testing.T) {
					t.Parallel()

					result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{testcase.NewTypeVersionInstruction}, legacyAdmin, config.DefaultCommitment)
					require.NotNil(t, result)

					output, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, testcase.Program.String(), func(b []byte) string {
						require.Len(t, b, int(binary.LittleEndian.Uint32(b[:4]))+4) // the first 4 bytes just encodes the length
						return string(b[4:])
					})
					require.NoError(t, err)

					// regex from https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
					semverRegex := "(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"
					require.Regexp(t, fmt.Sprintf("^%s %s$", testcase.ContractName, semverRegex), output)
					fmt.Printf(testcase.ContractName+" Type Version: %s\n", output)
				})
			}
		})

		t.Run("receiver", func(t *testing.T) {
			instruction, ixErr := test_ccip_receiver.NewInitializeInstruction(
				config.CcipRouterProgram,
				config.ReceiverTargetAccountPDA,
				config.ReceiverExternalExecutionConfigPDA,
				user.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, ixErr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
		})

		t.Run("billing", func(t *testing.T) {
			//////////
			// WSOL //
			//////////

			wsolPDA, _, aerr := state.FindFqBillingTokenConfigPDA(solana.SolMint, config.FeeQuoterProgram)
			require.NoError(t, aerr)
			wsolReceiver, _, rerr := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, config.BillingSignerPDA)
			require.NoError(t, rerr)
			wsolEvmConfigPDA, _, perr := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, solana.SolMint, config.FeeQuoterProgram)
			require.NoError(t, perr)
			wsolUserATA, _, uerr := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, user.PublicKey())
			require.NoError(t, uerr)
			wsolAnotherUserATA, _, auerr := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, anotherUser.PublicKey())
			require.NoError(t, auerr)
			wsolTokenlessUserATA, _, tuerr := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, tokenlessUser.PublicKey())
			require.NoError(t, tuerr)
			wsolFeeAggregatorATA, _, fuerr := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, feeAggregator.PublicKey())
			require.NoError(t, fuerr)

			// persist the WSOL config for later use
			wsol.program = solana.TokenProgramID
			wsol.mint = solana.SolMint
			wsol.fqBillingConfigPDA = wsolPDA
			wsol.userATA = wsolUserATA
			wsol.anotherUserATA = wsolAnotherUserATA
			wsol.tokenlessUserATA = wsolTokenlessUserATA
			wsol.billingATA = wsolReceiver
			wsol.feeAggregatorATA = wsolFeeAggregatorATA
			wsol.fqEvmConfigPDA = wsolEvmConfigPDA

			///////////////
			// link22 //
			///////////////
			linkMint := linkMintPrivK.PublicKey()
			ixToken, terr := tokens.CreateToken(ctx, config.Token2022Program, linkMint, legacyAdmin.PublicKey(), link22Decimals, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, terr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ixToken, legacyAdmin, config.DefaultCommitment, common.AddSigners(linkMintPrivK))

			link22PDA, _, aerr := state.FindFqBillingTokenConfigPDA(linkMint, config.FeeQuoterProgram)
			require.NoError(t, aerr)
			link22EvmConfigPDA, _, puerr := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, linkMint, config.FeeQuoterProgram)
			require.NoError(t, puerr)
			link22Receiver, _, rerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, config.BillingSignerPDA)
			require.NoError(t, rerr)
			link22UserATA, _, uerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, user.PublicKey())
			require.NoError(t, uerr)
			link22AnotherUserATA, _, auerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, anotherUser.PublicKey())
			require.NoError(t, auerr)
			link22TokenlessUserATA, _, tuerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, tokenlessUser.PublicKey())
			require.NoError(t, tuerr)
			link22FeeAggregatorATA, _, fuerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, linkMint, feeAggregator.PublicKey())
			require.NoError(t, fuerr)

			// persist the link22 billing config for later use
			link22.program = config.Token2022Program
			link22.mint = linkMint
			link22.fqBillingConfigPDA = link22PDA
			link22.userATA = link22UserATA
			link22.anotherUserATA = link22AnotherUserATA
			link22.tokenlessUserATA = link22TokenlessUserATA
			link22.billingATA = link22Receiver
			link22.feeAggregatorATA = link22FeeAggregatorATA
			link22.fqEvmConfigPDA = link22EvmConfigPDA
		})

		t.Run("billing:funding_and_approvals", func(t *testing.T) {
			type Item struct {
				name       string
				user       solana.PrivateKey
				getATA     func(apt *AccountsPerToken) solana.PublicKey
				shouldFund bool
			}
			list := []Item{
				{
					name:       "user",
					user:       user,
					getATA:     func(apt *AccountsPerToken) solana.PublicKey { return apt.userATA },
					shouldFund: true,
				},
				{
					name:       "anotherUser",
					user:       anotherUser,
					getATA:     func(apt *AccountsPerToken) solana.PublicKey { return apt.anotherUserATA },
					shouldFund: true,
				},
				{
					name:       "tokenlessUser",
					user:       tokenlessUser,
					getATA:     func(apt *AccountsPerToken) solana.PublicKey { return apt.tokenlessUserATA },
					shouldFund: false, // do not fund tokenless user
				},
			}

			feeAggrAtaIx := make([]solana.Instruction, len(billingTokens))
			for i, token := range billingTokens {
				// create ATA for fee aggregator
				ixAtaFeeAggr, addrFeeAggr, uerr := tokens.CreateAssociatedTokenAccount(token.program, token.mint, feeAggregator.PublicKey(), feeAggregator.PublicKey())
				require.NoError(t, uerr)
				require.Equal(t, token.feeAggregatorATA, addrFeeAggr, fmt.Sprintf("ATA for feeAggregator and token %s", token.name))
				feeAggrAtaIx[i] = ixAtaFeeAggr
			}
			testutils.SendAndConfirm(ctx, t, solanaGoClient, feeAggrAtaIx, feeAggregator, config.DefaultCommitment)

			for _, it := range list {
				for _, token := range billingTokens {
					// create ATA for user
					ixAtaUser, addrUser, uerr := tokens.CreateAssociatedTokenAccount(token.program, token.mint, it.user.PublicKey(), it.user.PublicKey())
					require.NoError(t, uerr)
					require.Equal(t, it.getATA(token), addrUser, fmt.Sprintf("ATA for user %s and token %s", it.name, token.name))

					// Approve CCIP to transfer the user's token for billing
					ixApprove, aerr := tokens.TokenApproveChecked(1e9, 9, token.program, it.getATA(token), token.mint, config.BillingSignerPDA, it.user.PublicKey(), []solana.PublicKey{})
					require.NoError(t, aerr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixAtaUser, ixApprove}, it.user, config.DefaultCommitment)
				}

				if it.shouldFund {
					// fund user link22 (mint directly to user ATA)
					ixMint, merr := tokens.MintTo(1e9, link22.program, link22.mint, it.getATA(&link22), legacyAdmin.PublicKey())
					require.NoError(t, merr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixMint}, legacyAdmin, config.DefaultCommitment)

					// fund user WSOL (transfer SOL + syncNative)
					transferAmount := 1.0 * solana.LAMPORTS_PER_SOL
					ixTransfer, terr := tokens.NativeTransfer(wsol.program, transferAmount, it.user.PublicKey(), it.getATA(&wsol))
					require.NoError(t, terr)
					ixSync, serr := tokens.SyncNative(wsol.program, it.getATA(&wsol))
					require.NoError(t, serr)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixTransfer, ixSync}, it.user, config.DefaultCommitment)
				}
			}
		})

		t.Run("token", func(t *testing.T) {
			ix0, ixErr0 := tokens.CreateToken(ctx, token0.Program, token0.Mint, token0PoolAdmin.PublicKey(), token0Decimals, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, ixErr0)

			ix1, ixErr1 := tokens.CreateToken(ctx, token1.Program, token1.Mint, token1PoolAdmin.PublicKey(), token1Decimals, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, ixErr1)

			ix2, ixErr2 := tokens.CreateToken(ctx, token2.Program, token2.Mint, token2PoolAdmin.PublicKey(), token2Decimals, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, ixErr2)

			// mint tokens to user
			ixAta0, addr0, ataErr := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint, user.PublicKey(), token0PoolAdmin.PublicKey())
			require.NoError(t, ataErr)
			ixMintTo0, mintErr := tokens.MintTo(10000000, token0.Program, token0.Mint, addr0, token0PoolAdmin.PublicKey())
			require.NoError(t, mintErr)
			ixAta1, addr1, ataErr := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint, user.PublicKey(), token1PoolAdmin.PublicKey())
			require.NoError(t, ataErr)
			ixMintTo1, mintErr := tokens.MintTo(10000000, token1.Program, token1.Mint, addr1, token1PoolAdmin.PublicKey())
			require.NoError(t, mintErr)
			ixAta2, addr2, ataErr := tokens.CreateAssociatedTokenAccount(token2.Program, token2.Mint, user.PublicKey(), token2PoolAdmin.PublicKey())
			require.NoError(t, ataErr)
			ixMintTo2, mintErr := tokens.MintTo(10000000, token2.Program, token2.Mint, addr2, token2PoolAdmin.PublicKey())
			require.NoError(t, mintErr)

			// create ATA for receiver (receiver program address)
			ixAtaReceiver0, recAddr0, recErr := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint, config.ReceiverExternalExecutionConfigPDA, token0PoolAdmin.PublicKey())
			require.NoError(t, recErr)
			ixAtaReceiver1, recAddr1, recErr := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint, config.ReceiverExternalExecutionConfigPDA, token1PoolAdmin.PublicKey())
			require.NoError(t, recErr)
			ixAtaReceiver2, recAddr2, recErr := tokens.CreateAssociatedTokenAccount(token2.Program, token2.Mint, config.ReceiverExternalExecutionConfigPDA, token2PoolAdmin.PublicKey())
			require.NoError(t, recErr)
			ixAtaReceiverLink, recAddrLink, recErr := tokens.CreateAssociatedTokenAccount(link22.program, link22.mint, config.ReceiverExternalExecutionConfigPDA, legacyAdmin.PublicKey())
			require.NoError(t, recErr)

			token0.User[user.PublicKey()] = addr0
			token0.User[config.ReceiverExternalExecutionConfigPDA] = recAddr0
			token1.User[user.PublicKey()] = addr1
			token1.User[config.ReceiverExternalExecutionConfigPDA] = recAddr1
			token2.User[user.PublicKey()] = addr2
			token2.User[config.ReceiverExternalExecutionConfigPDA] = recAddr2
			linkPool.User[user.PublicKey()] = link22.userATA
			linkPool.User[config.ReceiverExternalExecutionConfigPDA] = recAddrLink

			ix0 = append(ix0, ixAta0, ixMintTo0, ixAtaReceiver0)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ix0, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token0Mint))
			ix1 = append(ix1, ixAta1, ixMintTo1, ixAtaReceiver1)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ix1, token1PoolAdmin, config.DefaultCommitment, common.AddSigners(token1Mint))
			ix2 = append(ix2, ixAta2, ixMintTo2, ixAtaReceiver2)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ix2, token2PoolAdmin, config.DefaultCommitment, common.AddSigners(token2Mint))

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixAtaReceiverLink}, legacyAdmin, config.DefaultCommitment)
		})

		t.Run("token-pool", func(t *testing.T) {
			token0.AdditionalAccounts = append(token0.AdditionalAccounts, solana.MemoProgramID) // add test additional accounts in pool interactions

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

			ixInit0, err := test_token_pool.NewInitializeInstruction(
				test_token_pool.BurnAndMint_PoolType,
				config.CcipRouterProgram,
				config.RMNRemoteProgram,
				token0.PoolConfig,
				token0.Mint,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipTokenPoolProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			ixInit1, err := test_token_pool.NewInitializeInstruction(
				test_token_pool.BurnAndMint_PoolType,
				config.CcipRouterProgram,
				config.RMNRemoteProgram,
				token1.PoolConfig,
				token1.Mint,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipTokenPoolProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			ixInit2, err := test_token_pool.NewInitializeInstruction(
				test_token_pool.BurnAndMint_PoolType,
				config.CcipRouterProgram,
				config.RMNRemoteProgram,
				token2.PoolConfig,
				token2.Mint,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipTokenPoolProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			ixInitLink, err := test_token_pool.NewInitializeInstruction(
				test_token_pool.BurnAndMint_PoolType,
				config.CcipRouterProgram,
				config.RMNRemoteProgram,
				linkPool.PoolConfig,
				linkPool.Mint,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipTokenPoolProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			// The pools are initially owned by the CCIP admin. For the purposes of this test, they're now transferred to the individual
			// token admins. In practice, these token pools would've been original instantiated under separate token pool programs
			// owned by each individual admin instead.
			ixTransfer0, err := test_token_pool.NewTransferOwnershipInstruction(
				token0PoolAdmin.PublicKey(),
				token0.PoolConfig,
				token0.Mint,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixTransfer1, err := test_token_pool.NewTransferOwnershipInstruction(
				token1PoolAdmin.PublicKey(),
				token1.PoolConfig,
				token1.Mint,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixTransfer2, err := test_token_pool.NewTransferOwnershipInstruction(
				token2PoolAdmin.PublicKey(),
				token2.PoolConfig,
				token2.Mint,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			ixAccept0, err := test_token_pool.NewAcceptOwnershipInstruction(
				token0.PoolConfig,
				token0.Mint,
				token0PoolAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixAccept1, err := test_token_pool.NewAcceptOwnershipInstruction(
				token1.PoolConfig,
				token1.Mint,
				token1PoolAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			ixAccept2, err := test_token_pool.NewAcceptOwnershipInstruction(
				token2.PoolConfig,
				token2.Mint,
				token2PoolAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			ixAta0, addr0, err := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint, token0.PoolSigner, token0PoolAdmin.PublicKey())
			require.NoError(t, err)
			token0.PoolTokenAccount = addr0
			token0.User[token0.PoolSigner] = token0.PoolTokenAccount
			ixAta1, addr1, err := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint, token1.PoolSigner, token0PoolAdmin.PublicKey())
			require.NoError(t, err)
			token1.PoolTokenAccount = addr1
			token1.User[token1.PoolSigner] = token1.PoolTokenAccount
			ixAta2, addr2, err := tokens.CreateAssociatedTokenAccount(token2.Program, token2.Mint, token2.PoolSigner, token0PoolAdmin.PublicKey())
			require.NoError(t, err)
			token2.PoolTokenAccount = addr2
			token2.User[token2.PoolSigner] = token2.PoolTokenAccount
			ixAtaLink, addrLink, err := tokens.CreateAssociatedTokenAccount(linkPool.Program, linkPool.Mint, linkPool.PoolSigner, legacyAdmin.PublicKey())
			require.NoError(t, err)
			linkPool.PoolTokenAccount = addrLink
			linkPool.User[linkPool.PoolSigner] = linkPool.PoolTokenAccount

			ixAuth, err := tokens.SetTokenMintAuthority(token0.Program, token0.PoolSigner, token0.Mint, token0PoolAdmin.PublicKey())
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixInit0, ixInit1, ixInit2, ixInitLink}, legacyAdmin, config.DefaultCommitment)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixTransfer0, ixTransfer1, ixTransfer2, ixAccept0, ixAccept1, ixAccept2}, legacyAdmin, config.DefaultCommitment, common.AddSigners(token0PoolAdmin, token1PoolAdmin, token2PoolAdmin))
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixAta0, ixAta1, ixAuth, ixAta2, ixAtaLink}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin, token2PoolAdmin, legacyAdmin))

			// Lookup Table for Tokens
			require.NoError(t, token0.SetupLookupTable(ctx, solanaGoClient, token0PoolAdmin))
			token0Entries := token0.ToTokenPoolEntries()
			require.NoError(t, token1.SetupLookupTable(ctx, solanaGoClient, token1PoolAdmin))
			token1Entries := token1.ToTokenPoolEntries()
			require.NoError(t, token2.SetupLookupTable(ctx, solanaGoClient, token2PoolAdmin))
			token2Entries := token2.ToTokenPoolEntries()
			require.NoError(t, linkPool.SetupLookupTable(ctx, solanaGoClient, legacyAdmin))
			link22Entries := linkPool.ToTokenPoolEntries()

			// Verify Lookup tables where correctly initialized
			lookupTableEntries0, err := common.GetAddressLookupTable(ctx, solanaGoClient, token0.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(token0Entries), len(lookupTableEntries0))
			require.Equal(t, token0Entries, lookupTableEntries0)

			lookupTableEntries1, err := common.GetAddressLookupTable(ctx, solanaGoClient, token1.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(token1Entries), len(lookupTableEntries1))
			require.Equal(t, token1Entries, lookupTableEntries1)

			lookupTableEntries2, err := common.GetAddressLookupTable(ctx, solanaGoClient, token2.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(token2Entries), len(lookupTableEntries2))
			require.Equal(t, token2Entries, lookupTableEntries2)

			lookupTableEntriesLink, err := common.GetAddressLookupTable(ctx, solanaGoClient, linkPool.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(link22Entries), len(lookupTableEntriesLink))
			require.Equal(t, link22Entries, lookupTableEntriesLink)
		})

		t.Run("Ccip Send address lookup table", func(t *testing.T) {
			// Create single Address Lookup Table, to be used in all ccip send tests.
			// Create it early in the test suite (a "setup" step) to let it warm up with more than enough time,
			// as otherwise it can slow down tests  for ~20 seconds.
			// It includes most accounts that are used, though not all of them are used at the same time (some are either/or).
			lookupEntries := []solana.PublicKey{
				config.RouterConfigPDA,
				nonceEvmPDA,
				nonceSvmPDA,
				config.EvmDestChainStatePDA,
				config.SvmDestChainStatePDA,
				solana.SystemProgramID,
				solana.TokenProgramID,
				solana.Token2022ProgramID,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				config.FqSvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				wsol.anotherUserATA,
				link22.fqBillingConfigPDA,
				link22.mint,
				link22.userATA,
				link22.billingATA,
			}
			lookupTableAddr, err := common.SetupLookupTable(ctx, solanaGoClient, legacyAdmin, lookupEntries)
			require.NoError(t, err)

			ccipSendLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
				lookupTableAddr: lookupEntries,
			}
		})

		t.Run("Offramp address lookup table", func(t *testing.T) {
			// Create single Address Lookup Table, to be used in all commit tests and some execution ones.
			// Create it early in the test suite (a "setup" step) to let it warm up with more than enough time,
			// as otherwise it can slow down tests  for ~20 seconds.

			lookupEntries := []solana.PublicKey{
				config.CcipOfframpProgram,
				config.OfframpConfigPDA,
				config.OfframpReferenceAddressesPDA,
				config.OfframpEvmSourceChainPDA,
				solana.SystemProgramID,
				solana.SysVarInstructionsPubkey,
				config.OfframpBillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteConfigPDA,
				config.RMNRemoteCursesPDA,

				// remaining accounts used on some price update
				config.FqEvmDestChainPDA,
				config.FqSvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				wsol.fqEvmConfigPDA,
				link22.fqBillingConfigPDA,
				link22.fqEvmConfigPDA,
			}
			lookupTableAddr, err := common.SetupLookupTable(ctx, solanaGoClient, legacyAdmin, lookupEntries)
			require.NoError(t, err)

			offrampLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
				lookupTableAddr: lookupEntries,
			}
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
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.RMNRemoteProgram,
				programData.Address,
			).ValidateAndBuild()

			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
		})
	})

	//////////////////////////
	// Config Account Tests //
	//////////////////////////

	t.Run("Config", func(t *testing.T) {
		type ProgramData struct {
			DataType uint32
			Address  solana.PublicKey
		}

		const invalidSVMChainSelector uint64 = 17

		t.Run("Router is initialized", func(t *testing.T) {
			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			tempFeeAggregator := anotherUser.PublicKey() // fee aggregator address, will be changed in later test
			instruction, err := ccip_router.NewInitializeInstruction(
				invalidSVMChainSelector,
				tempFeeAggregator,
				config.FeeQuoterProgram,
				token0.Mint, // to be changed in the next tests
				config.RMNRemoteProgram,
				config.RouterConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipRouterProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configSetEvent ccip.EventRouterConfigSet
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
			require.Equal(t, invalidSVMChainSelector, configSetEvent.SvmChainSelector)
			require.Equal(t, config.FeeQuoterProgram, configSetEvent.FeeQuoter)
			require.Equal(t, config.RMNRemoteProgram, configSetEvent.RMNRemote)
			require.Equal(t, token0.Mint, configSetEvent.LinkTokenMint) // to be changed in the next tests
			require.Equal(t, tempFeeAggregator, configSetEvent.FeeAggregator)

			// Fetch account data
			var configAccount ccip_router.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, uint64(17), configAccount.SvmChainSelector)
			require.Equal(t, config.FeeQuoterProgram, configAccount.FeeQuoter)
		})

		t.Run("Router: Update link mint", func(t *testing.T) {
			t.Run("When a non-admin tries to make the update, it fails", func(t *testing.T) {
				ix, err := ccip_router.NewSetLinkTokenMintInstruction(
					link22.mint,
					config.RouterConfigPDA,
					user.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
			})
			t.Run("When an admin tries to make the update, it succeeds", func(t *testing.T) {
				// Fetch account data
				var configAccount ccip_router.Config
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount))
				require.Equal(t, token0.Mint, configAccount.LinkTokenMint) // initially set to token0 mint

				ix, err := ccip_router.NewSetLinkTokenMintInstruction(
					link22.mint,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				// Check that the event was emitted with the updated value
				var configSetEvent ccip.EventRouterConfigSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
				require.Equal(t, link22.mint, configSetEvent.LinkTokenMint)

				// Check the onchain state
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount))
				require.Equal(t, link22.mint, configAccount.LinkTokenMint)
			})
		})

		t.Run("FeeQuoter is initialized", func(t *testing.T) {
			initialMaxFeeJuelsPerMsg := bin.Uint128{Lo: 100000000000000000, Hi: 0, Endianness: nil} // this gets updated in a later test

			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.FeeQuoterProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			ix, err := fee_quoter.NewInitializeInstruction(
				initialMaxFeeJuelsPerMsg,
				config.CcipRouterProgram,
				config.FqConfigPDA,
				token0.Mint, // to be changed in the next tests
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.FeeQuoterProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configSetEvent ccip.EventFeeQuoterConfigSet
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
			require.Equal(t, initialMaxFeeJuelsPerMsg, configSetEvent.MaxFeeJuelsPerMsg)
			require.Equal(t, token0.Mint, configSetEvent.LinkTokenMint) // to be changed in the next tests
			require.Equal(t, uint8(9), configSetEvent.LinkTokenDecimals)
			require.Equal(t, config.CcipRouterProgram, configSetEvent.Onramp)
			require.Equal(t, fee_quoter.V1_CodeVersion, configSetEvent.DefaultCodeVersion)

			// Fetch account data
			var fqConfig fee_quoter.Config
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &fqConfig))

			require.Equal(t, token0.Mint, fqConfig.LinkTokenMint) // to be changed in the next tests
			require.Equal(t, token0Decimals, fqConfig.LinkTokenLocalDecimals)
			require.Equal(t, initialMaxFeeJuelsPerMsg, fqConfig.MaxFeeJuelsPerMsg)
			require.Equal(t, legacyAdmin.PublicKey(), fqConfig.Owner)
			require.True(t, fqConfig.ProposedOwner.IsZero())
			require.Equal(t, config.CcipRouterProgram, fqConfig.Onramp)
		})

		t.Run("FeeQuoter: Update link mint", func(t *testing.T) {
			t.Run("When a non-admin tries to make the update, it fails", func(t *testing.T) {
				ix, err := fee_quoter.NewSetLinkTokenMintInstruction(
					config.FqConfigPDA,
					link22.mint,
					user.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
			})
			t.Run("When an admin tries to make the update, it succeeds", func(t *testing.T) {
				// Fetch account data
				var configAccount fee_quoter.Config
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &configAccount))
				require.Equal(t, token0.Mint, configAccount.LinkTokenMint)             // initially set to token0 mint
				require.Equal(t, token0Decimals, configAccount.LinkTokenLocalDecimals) // initially set to token0 mint decimals

				ix, err := fee_quoter.NewSetLinkTokenMintInstruction(
					config.FqConfigPDA,
					link22.mint,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				// Check that the event was emitted with the updated value
				var configSetEvent ccip.EventFeeQuoterConfigSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
				require.Equal(t, link22.mint, configSetEvent.LinkTokenMint)
				require.Equal(t, link22Decimals, configSetEvent.LinkTokenDecimals)

				// Check the onchain state
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &configAccount))
				require.Equal(t, link22.mint, configAccount.LinkTokenMint)
				require.Equal(t, link22Decimals, configAccount.LinkTokenLocalDecimals)
			})
		})

		t.Run("FeeQuoter: Update max fee juels per msg", func(t *testing.T) {
			defaultMaxFeeJuelsPerMsg := bin.Uint128{Lo: 300000000000000000, Hi: 0, Endianness: nil}

			t.Run("When a non-admin tries to make the update, it fails", func(t *testing.T) {
				ix, err := fee_quoter.NewSetMaxFeeJuelsPerMsgInstruction(
					defaultMaxFeeJuelsPerMsg,
					config.FqConfigPDA,
					user.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
			})
			t.Run("When an admin tries to make the update, it succeeds", func(t *testing.T) {
				// Check that initial value is different to what is going to be set
				var fqConfig fee_quoter.Config
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &fqConfig))
				require.NotEqual(t, defaultMaxFeeJuelsPerMsg, fqConfig.MaxFeeJuelsPerMsg)

				// Actually update it
				ix, err := fee_quoter.NewSetMaxFeeJuelsPerMsgInstruction(
					defaultMaxFeeJuelsPerMsg,
					config.FqConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				// Check that the event was emitted with the updated value
				var configSetEvent ccip.EventFeeQuoterConfigSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
				require.Equal(t, defaultMaxFeeJuelsPerMsg, configSetEvent.MaxFeeJuelsPerMsg)

				// Check that the final value is the expected one
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &fqConfig))
				require.Equal(t, defaultMaxFeeJuelsPerMsg, fqConfig.MaxFeeJuelsPerMsg)
			})
		})

		t.Run("FeeQuoter: Add offramp as price updater", func(t *testing.T) {
			testutils.AssertClosedAccount(ctx, t, solanaGoClient, config.FqAllowedPriceUpdaterOfframpPDA, config.DefaultCommitment)

			ix, err := fee_quoter.NewAddPriceUpdaterInstruction(
				config.OfframpBillingSignerPDA,
				config.FqAllowedPriceUpdaterOfframpPDA,
				config.FqConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)

			isClosed := common.IsClosedAccount(ctx, solanaGoClient, config.FqAllowedPriceUpdaterOfframpPDA, config.DefaultCommitment)
			require.False(t, isClosed)
		})

		t.Run("Offramp is initialized", func(t *testing.T) {
			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipOfframpProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			var lookupTableAddr solana.PublicKey
			for k := range offrampLookupTable { // there is only one entry
				lookupTableAddr = k
			}

			// Check that a user who isn't the upgrade authority cannot call this
			badInitIx, err := ccip_offramp.NewInitializeInstruction(
				config.OfframpReferenceAddressesPDA,
				config.CcipRouterProgram,
				config.FeeQuoterProgram,
				config.RMNRemoteProgram,
				lookupTableAddr,
				config.OfframpStatePDA,
				user.PublicKey(), // not the upgrade authority
				solana.SystemProgramID,
				config.CcipOfframpProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{badInitIx}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})

			badInitConfigIx, err := ccip_offramp.NewInitializeConfigInstruction(
				invalidSVMChainSelector,
				config.EnableExecutionAfter,
				config.OfframpConfigPDA,
				ccipAdmin.PublicKey(), // not the upgrade authority
				solana.SystemProgramID,
				config.CcipOfframpProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{badInitConfigIx}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})

			testReferenceAddresses := ccip_offramp.ReferenceAddresses{
				Router:             [32]byte{1},
				FeeQuoter:          [32]byte{2},
				OfframpLookupTable: [32]byte{3},
				RmnRemote:          [32]byte{4},
			}

			// Now, actually initialize the offramp
			initIx, err := ccip_offramp.NewInitializeInstruction(
				config.OfframpReferenceAddressesPDA,
				testReferenceAddresses.Router,             // will be updated in later test
				testReferenceAddresses.FeeQuoter,          // will be updated in later test
				testReferenceAddresses.RmnRemote,          // will be updated in later test
				testReferenceAddresses.OfframpLookupTable, // will be updated in later test
				config.OfframpStatePDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipOfframpProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			initConfigIx, err := ccip_offramp.NewInitializeConfigInstruction(
				invalidSVMChainSelector,
				config.EnableExecutionAfter,
				config.OfframpConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipOfframpProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{initIx, initConfigIx}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var refAddrEvent ccip.EventOfframpReferenceAddressesSet
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ReferenceAddressesSet", &refAddrEvent, config.PrintEvents))
			require.Equal(t, testReferenceAddresses.Router, refAddrEvent.Router)
			require.Equal(t, testReferenceAddresses.FeeQuoter, refAddrEvent.FeeQuoter)
			require.Equal(t, testReferenceAddresses.OfframpLookupTable, refAddrEvent.OfframpLookupTable)
			require.Equal(t, testReferenceAddresses.RmnRemote, refAddrEvent.RMNRemote)

			var configSetEvent ccip.EventOfframpConfigSet
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
			require.Equal(t, invalidSVMChainSelector, configSetEvent.SvmChainSelector)
			require.Equal(t, config.EnableExecutionAfter, configSetEvent.EnableManualExecutionAfter)

			// Fetch account data
			var offrampConfig ccip_offramp.Config
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpConfigPDA, config.DefaultCommitment, &offrampConfig))
			require.Equal(t, invalidSVMChainSelector, offrampConfig.SvmChainSelector)
			require.Equal(t, config.EnableExecutionAfter, offrampConfig.EnableManualExecutionAfter)
			require.Equal(t, legacyAdmin.PublicKey(), offrampConfig.Owner)
			require.Equal(t, solana.PublicKey{}, offrampConfig.ProposedOwner)

			// check price sequence start
			var state ccip_offramp.GlobalState
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpStatePDA, config.DefaultCommitment, &state))
			require.Equal(t, uint64(0), state.LatestPriceSequenceNumber)

			// check reference addresses
			var referenceAddresses ccip_offramp.ReferenceAddresses
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpReferenceAddressesPDA, config.DefaultCommitment, &referenceAddresses))
			require.Equal(t, testReferenceAddresses.FeeQuoter, referenceAddresses.FeeQuoter)
			require.Equal(t, testReferenceAddresses.OfframpLookupTable, referenceAddresses.OfframpLookupTable)
			require.Equal(t, testReferenceAddresses.Router, referenceAddresses.Router)
		})

		t.Run("Offramp: Update reference addresses", func(t *testing.T) {
			var lookupTableAddr solana.PublicKey
			for k := range offrampLookupTable { // there is only one entry
				lookupTableAddr = k
			}

			t.Run("When an unauthorized user tries to make the update, it fails", func(t *testing.T) {
				ix, err := ccip_offramp.NewUpdateReferenceAddressesInstruction(
					config.CcipRouterProgram,
					config.FeeQuoterProgram,
					lookupTableAddr,
					config.RMNRemoteProgram,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					user.PublicKey(), // unauthorized user here
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
			})

			t.Run("When the admin makes the update, it succeeds", func(t *testing.T) {
				ix, err := ccip_offramp.NewUpdateReferenceAddressesInstruction(
					config.CcipRouterProgram,
					config.FeeQuoterProgram,
					lookupTableAddr,
					config.RMNRemoteProgram,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				// check event
				var refAddrEvent ccip.EventOfframpReferenceAddressesSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ReferenceAddressesSet", &refAddrEvent, config.PrintEvents))
				require.Equal(t, config.CcipRouterProgram, refAddrEvent.Router)
				require.Equal(t, config.FeeQuoterProgram, refAddrEvent.FeeQuoter)
				require.Equal(t, config.RMNRemoteProgram, refAddrEvent.RMNRemote)
				require.Equal(t, lookupTableAddr, refAddrEvent.OfframpLookupTable)

				// check state
				var referenceAddresses ccip_offramp.ReferenceAddresses
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpReferenceAddressesPDA, config.DefaultCommitment, &referenceAddresses))
				require.Equal(t, config.FeeQuoterProgram, referenceAddresses.FeeQuoter)
				require.Equal(t, lookupTableAddr, referenceAddresses.OfframpLookupTable)
				require.Equal(t, config.CcipRouterProgram, referenceAddresses.Router)
				require.Equal(t, config.RMNRemoteProgram, referenceAddresses.RmnRemote)
			})
		})

		t.Run("Offramp permissions", func(t *testing.T) {
			instruction, err := ccip_router.NewAddOfframpInstruction(
				config.EvmChainSelector, config.CcipOfframpProgram, config.AllowedOfframpEvmPDA, config.RouterConfigPDA, legacyAdmin.PublicKey(), solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
		})

		t.Run("When admin updates the solana chain selector it's updated", func(t *testing.T) {
			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewUpdateSvmChainSelectorInstruction(
					config.SvmChainSelector,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var configSetEvent ccip.EventRouterConfigSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
				require.Equal(t, config.SvmChainSelector, configSetEvent.SvmChainSelector)

				var configAccount ccip_router.Config
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, config.SvmChainSelector, configAccount.SvmChainSelector)
			})

			t.Run("Offramp", func(t *testing.T) {
				instruction, err := ccip_offramp.NewUpdateSvmChainSelectorInstruction(
					config.SvmChainSelector,
					config.OfframpConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var configSetEvent ccip.EventOfframpConfigSet
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
				require.Equal(t, config.SvmChainSelector, configSetEvent.SvmChainSelector)

				var configAccount ccip_offramp.Config
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpConfigPDA, config.DefaultCommitment, &configAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, config.SvmChainSelector, configAccount.SvmChainSelector)
			})
		})

		type InvalidChainBillingInputTest struct {
			Name         string
			Selector     uint64
			Conf         fee_quoter.DestChainConfig
			SkipOnUpdate bool
			Error        string
		}
		invalidInputTests := []InvalidChainBillingInputTest{
			{
				Name:     "Zero DefaultTxGasLimit",
				Selector: config.EvmChainSelector,
				Conf: fee_quoter.DestChainConfig{
					DefaultTxGasLimit:   0,
					MaxPerMsgGasLimit:   validFqDestChainConfig.MaxPerMsgGasLimit,
					ChainFamilySelector: validFqDestChainConfig.ChainFamilySelector,
				},
				Error: ccip.ZeroGasLimit_FeeQuoterError.String(),
			},
			{
				Name:         "Zero DestChainSelector",
				Selector:     0,
				Conf:         validFqDestChainConfig,
				SkipOnUpdate: true, // as the 0-selector is invalid, the config account can never be initialized
				Error:        ccip.InvalidInputsChainSelector_FeeQuoterError.String(),
			},
			{
				Name:     "Zero ChainFamilySelector",
				Selector: config.EvmChainSelector,
				Conf: fee_quoter.DestChainConfig{
					DefaultTxGasLimit:   validFqDestChainConfig.DefaultTxGasLimit,
					MaxPerMsgGasLimit:   validFqDestChainConfig.MaxPerMsgGasLimit,
					ChainFamilySelector: [4]uint8{0, 0, 0, 0},
				},
				Error: ccip.InvalidChainFamilySelector_FeeQuoterError.String(),
			},
			{
				Name:     "DefaultTxGasLimit > MaxPerMsgGasLimit",
				Selector: config.EvmChainSelector,
				Conf: fee_quoter.DestChainConfig{
					DefaultTxGasLimit:   100,
					MaxPerMsgGasLimit:   1,
					ChainFamilySelector: validFqDestChainConfig.ChainFamilySelector,
				},
				Error: ccip.DefaultGasLimitExceedsMaximum_FeeQuoterError.String(),
			},
		}

		t.Run("Fee Quoter: When and admin adds a chain selector with invalid dest chain config, it fails", func(t *testing.T) {
			for _, test := range invalidInputTests {
				t.Run(test.Name, func(t *testing.T) {
					destChainPDA, _, derr := state.FindFqDestChainPDA(test.Selector, config.FeeQuoterProgram)
					require.NoError(t, derr)

					instruction, err := fee_quoter.NewAddDestChainInstruction(
						test.Selector,
						test.Conf, // here is the invalid dest config data
						config.FqConfigPDA,
						destChainPDA,
						legacyAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment, []string{"Error Code: " + test.Error})
					require.NotNil(t, result)
				})
			}
		})

		t.Run("When an unauthorized user tries to add a chain selector, it fails", func(t *testing.T) {
			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.EvmChainSelector,
					ccip_router.DestChainConfig{},
					config.EvmDestChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(), // not an admin
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_FeeQuoterError.String()})
				require.NotNil(t, result)
			})

			t.Run("Fee Quoter", func(t *testing.T) {
				instruction, err := fee_quoter.NewAddDestChainInstruction(
					config.EvmChainSelector,
					validFqDestChainConfig,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					user.PublicKey(), // not an admin
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_FeeQuoterError.String()})
				require.NotNil(t, result)
			})
		})

		t.Run("When admin adds a chain selector it's added on the list", func(t *testing.T) {
			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.EvmChainSelector,
					ccip_router.DestChainConfig{},
					config.EvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var destChainStateAccount ccip_router.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmDestChainStatePDA, config.DefaultCommitment, &destChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(0), destChainStateAccount.State.SequenceNumber)
				require.Equal(t, ccip_router.DestChainConfig{}, destChainStateAccount.Config)
			})

			t.Run("Fee Quoter", func(t *testing.T) {
				instruction, err := fee_quoter.NewAddDestChainInstruction(
					config.EvmChainSelector,
					validFqDestChainConfig,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var destChainAccount fee_quoter.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &destChainAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, fee_quoter.TimestampedPackedU224{}, destChainAccount.State.UsdPerUnitGas)
				require.Equal(t, validFqDestChainConfig, destChainAccount.Config)
			})

			t.Run("Offramp", func(t *testing.T) {
				instruction, err := ccip_offramp.NewAddSourceChainInstruction(
					config.EvmChainSelector,
					validSourceChainConfig,
					config.OfframpEvmSourceChainPDA,
					config.OfframpConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var sourceChainStateAccount ccip_offramp.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &sourceChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(1), sourceChainStateAccount.State.MinSeqNr)
				require.Equal(t, true, sourceChainStateAccount.Config.IsEnabled)
				onRampAddress := [64]byte{1, 2, 3}

				require.Equal(t, ccip_offramp.OnRampAddress{Bytes: onRampAddress, Len: 3}, sourceChainStateAccount.Config.OnRamp)
			})
		})

		t.Run("When admin adds another chain selector it's also added on the list", func(t *testing.T) {
			// Using another chain, solana as an example (which allows SVM -> SVM messages)

			// the router is the SVM onramp
			paddedCcipRouterProgram := common.ToPadded64Bytes(config.CcipRouterProgram.Bytes())
			onRampConfig := ccip_offramp.OnRampAddress{
				Bytes: paddedCcipRouterProgram,
				Len:   32,
			}

			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.SvmChainSelector,
					ccip_router.DestChainConfig{},
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var destChainStateAccount ccip_router.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.SvmDestChainStatePDA, config.DefaultCommitment, &destChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(0), destChainStateAccount.State.SequenceNumber)
			})

			t.Run("Fee Quoter", func(t *testing.T) {
				instruction, err := fee_quoter.NewAddDestChainInstruction(
					config.SvmChainSelector,
					fee_quoter.DestChainConfig{
						IsEnabled: true,
						// minimal valid config
						DefaultTxGasLimit:       1,
						MaxPerMsgGasLimit:       100,
						ChainFamilySelector:     [4]uint8(config.SvmChainFamilySelector),
						EnforceOutOfOrder:       true,
						MaxNumberOfTokensPerMsg: 10,
						MaxDataBytes:            100,
					},
					config.FqConfigPDA,
					config.FqSvmDestChainPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var destChainStateAccount fee_quoter.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqSvmDestChainPDA, config.DefaultCommitment, &destChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, fee_quoter.TimestampedPackedU224{}, destChainStateAccount.State.UsdPerUnitGas)
			})

			t.Run("Offramp", func(t *testing.T) {
				instruction, err := ccip_offramp.NewAddSourceChainInstruction(
					config.SvmChainSelector,
					ccip_offramp.SourceChainConfig{
						OnRamp:    onRampConfig,
						IsEnabled: true,
					},
					config.OfframpSvmSourceChainPDA,
					config.OfframpConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var sourceChainStateAccount ccip_offramp.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpSvmSourceChainPDA, config.DefaultCommitment, &sourceChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(1), sourceChainStateAccount.State.MinSeqNr)
				require.Equal(t, true, sourceChainStateAccount.Config.IsEnabled)
				require.Equal(t, ccip_offramp.OnRampAddress{Bytes: paddedCcipRouterProgram, Len: 32}, sourceChainStateAccount.Config.OnRamp)
			})
		})

		t.Run("Onramp lane ccip version", func(t *testing.T) {
			var initial ccip_router.DestChain
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.SvmDestChainStatePDA, config.DefaultCommitment, &initial))
			require.Equal(t, ccip_router.None_RestoreOnAction, initial.State.RestoreOnAction)

			t.Run("When a non-admin tries to bump the onramp ccip version for a dest chain, it fails", func(t *testing.T) {
				ix, err := ccip_router.NewBumpCcipVersionForDestChainInstruction(
					config.SvmChainSelector,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(), // not the admin
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			})

			t.Run("When an admin bumps the onramp ccip version for a dest chain, it succeeds", func(t *testing.T) {
				ix, err := ccip_router.NewBumpCcipVersionForDestChainInstruction(
					config.SvmChainSelector,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var routerDestChain ccip_router.DestChain
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.SvmDestChainStatePDA, config.DefaultCommitment, &routerDestChain))
				require.Equal(t, ccip_router.Rollback_RestoreOnAction, routerDestChain.State.RestoreOnAction)
			})

			t.Run("When a non-admin tries to rollback the onramp ccip version for a dest chain, it fails", func(t *testing.T) {
				ix, err := ccip_router.NewRollbackCcipVersionForDestChainInstruction(
					config.SvmChainSelector,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(), // not the admin
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			})

			t.Run("When an admin rollbacks the onramp ccip version for a dest chain, it succeeds", func(t *testing.T) {
				ix, err := ccip_router.NewRollbackCcipVersionForDestChainInstruction(
					config.SvmChainSelector,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var routerDestChain ccip_router.DestChain
				require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.SvmDestChainStatePDA, config.DefaultCommitment, &routerDestChain))
				require.Equal(t, ccip_router.Upgrade_RestoreOnAction, routerDestChain.State.RestoreOnAction)
			})

			t.Run("When an admin tries to rollback again the onramp ccip version for a dest chain, it fails", func(t *testing.T) {
				ix, err := ccip_router.NewRollbackCcipVersionForDestChainInstruction(
					config.SvmChainSelector,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.InvalidCcipVersionRollback_CcipRouterError.String()})
			})
		})

		t.Run("When a non-admin tries to disable the chain selector, it fails", func(t *testing.T) {
			t.Run("Offramp: Source", func(t *testing.T) {
				ix, err := ccip_offramp.NewDisableSourceChainSelectorInstruction(
					config.EvmChainSelector,
					config.OfframpEvmSourceChainPDA,
					config.OfframpConfigPDA,
					user.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})

			t.Run("Fee Quoter: Dest", func(t *testing.T) {
				ix, err := fee_quoter.NewDisableDestChainInstruction(
					config.EvmChainSelector,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					user.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})
		})

		t.Run("When an admin disables the chain selector, it is no longer enabled", func(t *testing.T) {
			t.Run("Offramp: Source", func(t *testing.T) {
				var initial ccip_offramp.SourceChain
				err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &initial)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, true, initial.Config.IsEnabled)

				ix, err := ccip_offramp.NewDisableSourceChainSelectorInstruction(
					config.EvmChainSelector,
					config.OfframpEvmSourceChainPDA,
					config.OfframpConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)

				var final ccip_offramp.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &final)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, false, final.Config.IsEnabled)
			})

			t.Run("Fee Quoter: Dest", func(t *testing.T) {
				var initial fee_quoter.DestChain
				err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &initial)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, true, initial.Config.IsEnabled)

				ix, err := fee_quoter.NewDisableDestChainInstruction(
					config.EvmChainSelector,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)

				var final fee_quoter.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &final)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, false, final.Config.IsEnabled)
			})
		})

		t.Run("Fee Quoter: When an admin tries to update the chain state with invalid destination chain config, it fails", func(t *testing.T) {
			for _, test := range invalidInputTests {
				if test.SkipOnUpdate {
					continue
				}
				t.Run(test.Name, func(t *testing.T) {
					destChainPDA, _, derr := state.FindFqDestChainPDA(test.Selector, config.FeeQuoterProgram)
					require.NoError(t, derr)
					instruction, err := fee_quoter.NewUpdateDestChainConfigInstruction(
						test.Selector,
						test.Conf,
						config.FqConfigPDA,
						destChainPDA,
						legacyAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment, []string{"Error Code: " + test.Error})
					require.NotNil(t, result)
				})
			}
		})

		t.Run("When an unauthorized user tries to update the chain state config, it fails", func(t *testing.T) {
			t.Run("Offramp Source", func(t *testing.T) {
				instruction, err := ccip_offramp.NewUpdateSourceChainConfigInstruction(
					config.EvmChainSelector,
					validSourceChainConfig,
					config.OfframpEvmSourceChainPDA,
					config.OfframpConfigPDA,
					user.PublicKey(), // unauthorized
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})

			t.Run("FeeQuoter: Dest", func(t *testing.T) {
				instruction, err := fee_quoter.NewUpdateDestChainConfigInstruction(
					config.EvmChainSelector,
					validFqDestChainConfig,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					user.PublicKey(), // unauthorized
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})
		})

		t.Run("When an admin updates the chain state config, it is configured", func(t *testing.T) {
			t.Run("Offramp: Source", func(t *testing.T) {
				var initialSource ccip_offramp.SourceChain
				serr := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &initialSource)
				require.NoError(t, serr, "failed to get account info")

				updated := initialSource.Config
				updated.IsEnabled = true
				require.NotEqual(t, initialSource.Config, updated) // at this point, onchain is disabled and we'll re-enable it

				instruction, err := ccip_offramp.NewUpdateSourceChainConfigInstruction(
					config.EvmChainSelector,
					updated,
					config.OfframpEvmSourceChainPDA,
					config.OfframpConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var final ccip_offramp.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &final)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, updated, final.Config)
			})

			t.Run("Fee Quoter: Dest", func(t *testing.T) {
				var initialDest fee_quoter.DestChain
				derr := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &initialDest)
				require.NoError(t, derr, "failed to get account info")

				updated := initialDest.Config
				updated.IsEnabled = true
				require.NotEqual(t, initialDest.Config, updated) // at this point, onchain is disabled and we'll re-enable it

				instruction, err := fee_quoter.NewUpdateDestChainConfigInstruction(
					config.EvmChainSelector,
					updated,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var final fee_quoter.DestChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &final)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, updated, final.Config)
			})
		})

		t.Run("When an unauthorized user tries to update the fee aggregator, it fails", func(t *testing.T) {
			instruction, err := ccip_router.NewUpdateFeeAggregatorInstruction(
				user.PublicKey(), // updating to some other address
				config.RouterConfigPDA,
				user.PublicKey(), // wrong user
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			require.NotNil(t, result)
		})

		t.Run("When an authorized user tries updates the fee aggregator, it succeeds", func(t *testing.T) {
			var configAccount ccip_router.Config
			err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")
			require.NotEqual(t, feeAggregator.PublicKey(), configAccount.FeeAggregator) // at this point, the fee aggregator is different

			instruction, err := ccip_router.NewUpdateFeeAggregatorInstruction(
				feeAggregator.PublicKey(), // updating to some other address
				config.RouterConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configSetEvent ccip.EventRouterConfigSet
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
			require.Equal(t, feeAggregator.PublicKey(), configSetEvent.FeeAggregator)

			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, feeAggregator.PublicKey(), configAccount.FeeAggregator) // now the fee aggregator is updated
		})

		t.Run("Can transfer ownership", func(t *testing.T) {
			// Fail to transfer ownership when not owner
			instruction, err := ccip_router.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RouterConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			require.NotNil(t, result)

			// successfully transfer ownership
			instruction, err = ccip_router.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RouterConfigPDA,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			transferEvent := ccip.OwnershipTransferRequested{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferRequested", &transferEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Fail to accept ownership when not proposed_owner
			instruction, err = ccip_router.NewAcceptOwnershipInstruction(
				config.RouterConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			require.NotNil(t, result)

			// Successfully accept ownership
			// ccipAdmin becomes owner for remaining tests
			instruction, err = ccip_router.NewAcceptOwnershipInstruction(
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
			acceptEvent := ccip.OwnershipTransferred{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferred", &acceptEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Current owner cannot propose self
			instruction, err = ccip_router.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.RedundantOwnerProposal_CcipRouterError.String()})
			require.NotNil(t, result)

			// Validate proposed set to 0-address
			var configAccount ccip_router.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})

		t.Run("Fee Quoter: Can transfer ownership", func(t *testing.T) {
			// Fail to transfer ownership when not owner
			instruction, err := fee_quoter.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.FqConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			require.NotNil(t, result)

			// successfully transfer ownership
			instruction, err = fee_quoter.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.FqConfigPDA,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			transferEvent := ccip.OwnershipTransferRequested{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferRequested", &transferEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Fail to accept ownership when not proposed_owner
			instruction, err = fee_quoter.NewAcceptOwnershipInstruction(
				config.FqConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})
			require.NotNil(t, result)

			// Successfully accept ownership
			// ccipAdmin becomes owner for remaining tests
			instruction, err = fee_quoter.NewAcceptOwnershipInstruction(
				config.FqConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
			acceptEvent := ccip.OwnershipTransferred{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferred", &acceptEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Current owner cannot propose self
			instruction, err = fee_quoter.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.FqConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.RedundantOwnerProposal_FeeQuoterError.String()})
			require.NotNil(t, result)

			// Validate proposed set to 0-address
			var configAccount fee_quoter.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})

		t.Run("Offramp: Can transfer ownership", func(t *testing.T) {
			// Fail to transfer ownership when not owner
			instruction, err := ccip_offramp.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.OfframpConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipOfframpError.String()})
			require.NotNil(t, result)

			// successfully transfer ownership
			instruction, err = ccip_offramp.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.OfframpConfigPDA,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			transferEvent := ccip.OwnershipTransferRequested{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferRequested", &transferEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Fail to accept ownership when not proposed_owner
			instruction, err = ccip_offramp.NewAcceptOwnershipInstruction(
				config.OfframpConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipOfframpError.String()})
			require.NotNil(t, result)

			// Successfully accept ownership
			// ccipAdmin becomes owner for remaining tests
			instruction, err = ccip_offramp.NewAcceptOwnershipInstruction(
				config.OfframpConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
			acceptEvent := ccip.OwnershipTransferred{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferred", &acceptEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Current owner cannot propose self
			instruction, err = ccip_offramp.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.OfframpConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.RedundantOwnerProposal_CcipOfframpError.String()})
			require.NotNil(t, result)

			// Validate proposed set to 0-address
			var configAccount ccip_offramp.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})
		t.Run("RMNRemote: Can transfer ownership", func(t *testing.T) {
			// Fail to transfer ownership when not owner
			instruction, err := rmn_remote.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RMNRemoteConfigPDA,
				config.RMNRemoteCursesPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_RmnRemoteError.String()})
			require.NotNil(t, result)

			// successfully transfer ownership
			instruction, err = rmn_remote.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RMNRemoteConfigPDA,
				config.RMNRemoteCursesPDA,
				legacyAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			transferEvent := ccip.OwnershipTransferRequested{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferRequested", &transferEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Fail to accept ownership when not proposed_owner
			instruction, err = rmn_remote.NewAcceptOwnershipInstruction(
				config.RMNRemoteConfigPDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_RmnRemoteError.String()})
			require.NotNil(t, result)

			// Successfully accept ownership
			// ccipAdmin becomes owner for remaining tests
			instruction, err = rmn_remote.NewAcceptOwnershipInstruction(
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
			acceptEvent := ccip.OwnershipTransferred{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "OwnershipTransferred", &acceptEvent, config.PrintEvents))
			require.Equal(t, legacyAdmin.PublicKey(), transferEvent.From)
			require.Equal(t, ccipAdmin.PublicKey(), transferEvent.To)

			// Current owner cannot propose self
			instruction, err = rmn_remote.NewTransferOwnershipInstruction(
				ccipAdmin.PublicKey(),
				config.RMNRemoteConfigPDA,
				config.RMNRemoteCursesPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.RedundantOwnerProposal_CcipOfframpError.String()})
			require.NotNil(t, result)

			// Validate proposed set to 0-address
			var configAccount rmn_remote.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, solana.PublicKey{}, configAccount.ProposedOwner)
		})
	})

	//////////////////////////
	// Billing Config Tests //
	//////////////////////////

	t.Run("Billing", func(t *testing.T) {
		t.Run("setup:fee_quoter:add_tokens", func(t *testing.T) {
			type TestToken struct {
				Config   fee_quoter.BillingTokenConfig
				Accounts AccountsPerToken
			}

			// Any nonzero timestamp is valid (for now)
			validTimestamp := int64(100)
			bigValue := [28]uint8{}
			bigNum, ok := new(big.Int).SetString("19816680000000000000", 10)
			require.True(t, ok)
			bigNum.FillBytes(bigValue[:])

			smallValue := [28]uint8{}
			smallNum, ok := new(big.Int).SetString("1981668000000000000", 10)
			require.True(t, ok)
			smallNum.FillBytes(smallValue[:])

			testTokens := []TestToken{
				{
					Accounts: wsol,
					Config: fee_quoter.BillingTokenConfig{
						Enabled: true,
						Mint:    solana.SolMint,
						UsdPerToken: fee_quoter.TimestampedPackedU224{
							Value:     smallValue,
							Timestamp: validTimestamp,
						},
						PremiumMultiplierWeiPerEth: 9000000,
					}},
				{
					Accounts: link22,
					Config: fee_quoter.BillingTokenConfig{
						Enabled: true,
						Mint:    link22.mint,
						UsdPerToken: fee_quoter.TimestampedPackedU224{
							Value:     bigValue,
							Timestamp: validTimestamp,
						},
						PremiumMultiplierWeiPerEth: 11000000,
					}},
			}

			for _, token := range testTokens {
				t.Run("add_"+token.Accounts.name, func(t *testing.T) {
					ixConfig, cerr := fee_quoter.NewAddBillingTokenConfigInstruction(
						token.Config,
						config.FqConfigPDA,
						token.Accounts.fqBillingConfigPDA,
						token.Accounts.program,
						token.Accounts.mint,
						token.Accounts.billingATA,
						ccipAdmin.PublicKey(),
						config.BillingSignerPDA,
						tokens.AssociatedTokenProgramID,
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, cerr)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, ccipAdmin, config.DefaultCommitment)
				})
			}
		})

		t.Run("FeeQuoter: Billing Token Config", func(t *testing.T) {
			pools := []tokens.TokenPool{token0, token1}

			for i, token := range pools {
				t.Run(fmt.Sprintf("token%d", i), func(t *testing.T) {
					t.Run("Pre-condition: Does not support token by default", func(t *testing.T) {
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint)
						var tokenConfigAccount fee_quoter.BillingTokenConfigWrapper
						err := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &tokenConfigAccount)
						require.EqualError(t, err, "not found")
					})

					t.Run("When admin adds token with valid input it is configured", func(t *testing.T) {
						// Any nonzero timestamp is valid (for now)
						validTimestamp := int64(100)
						value := [28]uint8{}
						big.NewInt(3e18).FillBytes(value[:])

						tokenConfig := fee_quoter.BillingTokenConfig{
							Enabled: true,
							Mint:    token.Mint,
							UsdPerToken: fee_quoter.TimestampedPackedU224{
								Timestamp: validTimestamp,
								Value:     value,
							},
							PremiumMultiplierWeiPerEth: 1,
						}

						tokenBillingPDA := getFqTokenConfigPDA(token.Mint)
						tokenReceiver, _, ferr := tokens.FindAssociatedTokenAddress(token.Program, token.Mint, config.BillingSignerPDA)
						require.NoError(t, ferr)

						ixConfig, cerr := fee_quoter.NewAddBillingTokenConfigInstruction(
							tokenConfig,
							config.FqConfigPDA,
							tokenBillingPDA,
							token.Program,
							token.Mint,
							tokenReceiver,
							ccipAdmin.PublicKey(),
							config.BillingSignerPDA,
							tokens.AssociatedTokenProgramID,
							solana.SystemProgramID,
						).ValidateAndBuild()
						require.NoError(t, cerr)

						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, ccipAdmin, config.DefaultCommitment)

						var tokenConfigAccount fee_quoter.BillingTokenConfigWrapper
						aerr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &tokenConfigAccount)
						require.NoError(t, aerr)

						require.Equal(t, tokenConfig, tokenConfigAccount.Config)
					})

					t.Run("When an unauthorized user updates token with correct configuration it fails", func(t *testing.T) {
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint)
						var initial fee_quoter.BillingTokenConfigWrapper
						ierr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &initial)
						require.NoError(t, ierr)

						tokenConfig := initial.Config
						tokenConfig.PremiumMultiplierWeiPerEth = initial.Config.PremiumMultiplierWeiPerEth*2 + 1 // updating something valid

						ixConfig, cerr := fee_quoter.NewUpdateBillingTokenConfigInstruction(tokenConfig, config.FqConfigPDA, tokenBillingPDA, legacyAdmin.PublicKey()).ValidateAndBuild() // wrong admin
						require.NoError(t, cerr)
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, legacyAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})

						var final fee_quoter.BillingTokenConfigWrapper
						ferr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &final)
						require.NoError(t, ferr)

						require.Equal(t, initial.Config, final.Config) // it was not updated, same values as initial
					})

					t.Run("When admin updates token it is updated", func(t *testing.T) {
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint)
						var initial fee_quoter.BillingTokenConfigWrapper
						ierr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &initial)
						require.NoError(t, ierr)

						tokenConfig := initial.Config
						tokenConfig.PremiumMultiplierWeiPerEth = initial.Config.PremiumMultiplierWeiPerEth*2 + 1 // updating something else

						ixConfig, cerr := fee_quoter.NewUpdateBillingTokenConfigInstruction(tokenConfig, config.FqConfigPDA, tokenBillingPDA, ccipAdmin.PublicKey()).ValidateAndBuild()
						require.NoError(t, cerr)
						result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, ccipAdmin, config.DefaultCommitment)

						// check event PremiumMultiplierWeiPerEthUpdated
						premiumMultiplierWeiPerEthUpdatedEvent := ccip.PremiumMultiplierWeiPerEthUpdated{}
						require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "PremiumMultiplierWeiPerEthUpdated", &premiumMultiplierWeiPerEthUpdatedEvent, config.PrintEvents))
						require.Equal(t, tokenConfig.Mint, premiumMultiplierWeiPerEthUpdatedEvent.Token)
						require.Equal(t, tokenConfig.PremiumMultiplierWeiPerEth, premiumMultiplierWeiPerEthUpdatedEvent.PremiumMultiplierWeiPerEth)

						var final fee_quoter.BillingTokenConfigWrapper
						ferr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, rpc.CommitmentProcessed, &final)
						require.NoError(t, ferr)

						require.NotEqual(t, initial.Config.PremiumMultiplierWeiPerEth, final.Config.PremiumMultiplierWeiPerEth) // it was updated
						require.Equal(t, tokenConfig.PremiumMultiplierWeiPerEth, final.Config.PremiumMultiplierWeiPerEth)
					})
				})
			}
		})
	})

	//////////////////////////
	//  setOcrConfig Tests  //
	//////////////////////////

	t.Run("Offramp: Config SetOcrConfig", func(t *testing.T) {
		t.Run("Successfully configures commit & execute DON ocr config for maximum signers and transmitters", func(t *testing.T) {
			// Check owner permissions
			instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
				0,
				ccip_offramp.Ocr3ConfigInfo{},
				[][20]byte{},
				[]solana.PublicKey{},
				config.OfframpConfigPDA,
				config.OfframpStatePDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Unauthorized_CcipRouterError.String()})

			inputs := []struct {
				plugin       testutils.OcrPlugin
				signers      [][20]byte
				transmitters []solana.PublicKey
				verifySig    uint8 // use as bool
			}{
				{
					testutils.OcrCommitPlugin,
					signerAddresses,
					transmitterPubKeys,
					1, // true
				},
				{
					testutils.OcrExecutePlugin,
					nil,
					transmitterPubKeys,
					0, // no sign verify needed for execute
				},
			}

			for _, v := range inputs {
				t.Run(v.plugin.String(), func(t *testing.T) {
					instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
						ccip_offramp.OcrPluginType(v.plugin),
						ccip_offramp.Ocr3ConfigInfo{
							ConfigDigest:                   config.ConfigDigest,
							F:                              config.OcrF,
							IsSignatureVerificationEnabled: v.verifySig,
						},
						v.signers,
						v.transmitters,
						config.OfframpConfigPDA,
						config.OfframpStatePDA,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
					require.NotNil(t, result)

					// Check event ConfigSet
					configSetEvent := ccip.EventConfigSet{}
					require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
					require.Equal(t, uint8(v.plugin), configSetEvent.OcrPluginType)
					require.Equal(t, config.ConfigDigest, configSetEvent.ConfigDigest)
					require.Equal(t, config.OcrF, configSetEvent.F)
					require.Equal(t, v.signers, configSetEvent.Signers)
					require.Equal(t, v.transmitters, configSetEvent.Transmitters)

					// check config state
					var configAccount ccip_offramp.Config
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpConfigPDA, config.DefaultCommitment, &configAccount)
					if err != nil {
						require.NoError(t, err, "failed to get account info")
					}
					pluginState := configAccount.Ocr3[v.plugin]
					require.Equal(t, config.ConfigDigest, pluginState.ConfigInfo.ConfigDigest)
					require.Equal(t, config.OcrF, pluginState.ConfigInfo.F)
					require.Equal(t, len(v.signers), int(pluginState.ConfigInfo.N))
					require.Equal(t, v.verifySig, pluginState.ConfigInfo.IsSignatureVerificationEnabled)
					require.Equal(t, config.MaxSignersAndTransmitters, len(pluginState.Signers))
					require.Equal(t, config.MaxSignersAndTransmitters, len(pluginState.Transmitters))
					for i := 0; i < config.MaxSignersAndTransmitters; i++ {
						// check signers (and zero values)
						signer := [20]byte{}
						if i < len(v.signers) {
							signer = v.signers[i]
						}
						require.Equal(t, signer, pluginState.Signers[i])

						// check transmitters (and zero values)
						transmitter := solana.PublicKey{}
						if i < len(v.transmitters) {
							transmitter = v.transmitters[i]
						}
						require.Equal(t, transmitter.Bytes(), pluginState.Transmitters[i][:])
					}
				})
			}
		})

		t.Run("SetOcrConfig edge cases", func(t *testing.T) {
			t.Run("It rejects an invalid plugin type", func(t *testing.T) {
				t.Parallel()
				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(100),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					transmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + common.InstructionDidNotDeserialize_AnchorError.String()})
			})

			t.Run("It rejects F = 0", func(t *testing.T) {
				t.Parallel()
				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            0,
					},
					signerAddresses,
					transmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigFMustBePositive_CcipOfframpError.String()})
			})

			t.Run("It rejects too many transmitters", func(t *testing.T) {
				t.Parallel()
				invalidTransmitters := make([]solana.PublicKey, config.MaxOracles+1)
				for i := range invalidTransmitters {
					invalidTransmitters[i] = getTransmitter().PublicKey()
				}
				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitters,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigTooManyTransmitters_CcipOfframpError.String()})
			})

			t.Run("It rejects too many signers", func(t *testing.T) {
				t.Parallel()
				invalidSigners := make([][20]byte, config.MaxOracles+1)
				for i := range invalidSigners {
					invalidSigners[i] = signerAddresses[0]
				}

				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSigners,
					transmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigTooManySigners_CcipOfframpError.String()})
			})

			t.Run("It rejects too high of F for signers", func(t *testing.T) {
				t.Parallel()
				invalidSigners := make([][20]byte, 1)
				invalidSigners[0] = signerAddresses[0]

				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSigners,
					transmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigFIsTooHigh_CcipOfframpError.String()})
			})

			t.Run("It rejects duplicate transmitters", func(t *testing.T) {
				t.Parallel()
				transmitter := getTransmitter().PublicKey()

				invalidTransmitters := make([]solana.PublicKey, 2)
				for i := range invalidTransmitters {
					invalidTransmitters[i] = transmitter
				}
				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitters,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigRepeatedOracle_CcipOfframpError.String()})
			})

			t.Run("It rejects duplicate signers", func(t *testing.T) {
				t.Parallel()
				repeatedSignerAddresses := [][20]byte{}
				for range signers {
					repeatedSignerAddresses = append(repeatedSignerAddresses, signers[0].Address)
				}
				oneTransmitter := []solana.PublicKey{transmitterPubKeys[0]}

				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					repeatedSignerAddresses,
					oneTransmitter,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3InvalidConfigRepeatedOracle_CcipOfframpError.String()})
			})

			t.Run("It rejects zero transmitter address", func(t *testing.T) {
				t.Parallel()
				invalidTransmitterPubKeys := []solana.PublicKey{transmitterPubKeys[0], common.ZeroAddress}

				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3OracleCannotBeZeroAddress_CcipOfframpError.String()})
			})

			t.Run("It rejects zero signer address", func(t *testing.T) {
				t.Parallel()
				invalidSignerAddresses := [][20]byte{{}}
				for _, v := range signers[1:] {
					invalidSignerAddresses = append(invalidSignerAddresses, v.Address)
				}
				instruction, err := ccip_offramp.NewSetOcrConfigInstruction(
					ccip_offramp.OcrPluginType(testutils.OcrCommitPlugin),
					ccip_offramp.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSignerAddresses,
					transmitterPubKeys,
					config.OfframpConfigPDA,
					config.OfframpStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3OracleCannotBeZeroAddress_CcipOfframpError.String()})
			})
		})
	})

	////////////////////////////////
	// Token Admin Registry Tests //
	////////////////////////////////

	t.Run("Token Admin Registry", func(t *testing.T) {
		t.Run("Token Admin Registry by Admin", func(t *testing.T) {
			t.Run("propose token admin registry as ccip admin", func(t *testing.T) {
				t.Run("When any user wants to set up the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						user.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the token admin registry, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						transmitter.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to set up the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("accept token admin registry as token admin", func(t *testing.T) {
				t.Run("When any user wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When proposed admin wants to accept the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("when admin is set, proposing a new one fails", func(t *testing.T) {
				instruction, err := ccip_router.NewCcipAdminOverridePendingAdministratorInstruction(
					token0PoolAdmin.PublicKey(),
					config.RouterConfigPDA,
					token0.AdminRegistryPDA,
					token0.Mint,
					ccipAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.InvalidTokenAdminRegistryProposedAdmin_CcipRouterError.String()})
			})

			t.Run("set pool", func(t *testing.T) {
				t.Run("When any user wants to set up the pool, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the pool, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						transmitter.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to set up the pool, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When setting pool to incorrect addresses in lookup table, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token1.PoolLookupTable, // accounts do not match the expected mint related accounts
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When Token Pool Admin wants to set up the pool, it succeeds", func(t *testing.T) {
					base := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					)

					base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(token0.PoolLookupTable))
					instruction, err := base.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token0.PoolLookupTable, tokenAdminRegistry.LookupTable)
				})

				t.Run("When Token Pool Admin wants to set up the pool again to zero, it is none", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						solana.PublicKey{},
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)

					// Rollback to previous state
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)
				})
			})

			t.Run("Transfer admin role for token admin registry", func(t *testing.T) {
				t.Run("When any user wants to transfer the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to transfer the token admin registry, it succeeds and permissions stay with no changes", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token0.PoolLookupTable, tokenAdminRegistry.LookupTable)

					// check if the admin is still the same
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// new one cant make changes yet
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When new admin accepts the token admin registry, it succeeds and permissions are updated", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token0.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token0.PoolLookupTable, tokenAdminRegistry.LookupTable)

					// check old admin can not make changes anymore
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})

					// new one can make changes now
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint,
						token0.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)
				})
			})
		})

		t.Run("Token Admin Registry by Mint Authority", func(t *testing.T) {
			t.Run("propose token admin registry as token mint authority", func(t *testing.T) {
				t.Run("When any user wants to set up the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						user.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the token admin registry, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						transmitter.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to set up the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When token mint_authority wants to set up the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1PoolAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("after admin is proposed but not accepted, overriding it succeeds", func(t *testing.T) {
				instruction, err := ccip_router.NewOwnerOverridePendingAdministratorInstruction(
					token2PoolAdmin.PublicKey(),
					config.RouterConfigPDA,
					token1.AdminRegistryPDA,
					token1.Mint,
					token1PoolAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

				tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
				require.NoError(t, err)
				require.Equal(t, uint8(1), tokenAdminRegistry.Version)
				require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
				require.Equal(t, token2PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
				require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)

				// We override it back to the previous proposal
				instruction, err = ccip_router.NewOwnerOverridePendingAdministratorInstruction(
					token1PoolAdmin.PublicKey(),
					config.RouterConfigPDA,
					token1.AdminRegistryPDA,
					token1.Mint,
					token1PoolAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

				tokenAdminRegistry = ccip_common.TokenAdminRegistry{}
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
				require.NoError(t, err)
				require.Equal(t, uint8(1), tokenAdminRegistry.Version)
				require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
				require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
				require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
			})

			t.Run("accept token admin registry as token admin", func(t *testing.T) {
				t.Run("When any user wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When proposed admin wants to accept the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("when admin is set, proposing a new one fails", func(t *testing.T) {
				instruction, err := ccip_router.NewOwnerOverridePendingAdministratorInstruction(
					token1PoolAdmin.PublicKey(),
					config.RouterConfigPDA,
					token1.AdminRegistryPDA,
					token1.Mint,
					token1PoolAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip.InvalidTokenAdminRegistryProposedAdmin_CcipRouterError.String()})
			})

			t.Run("set pool", func(t *testing.T) {
				t.Run("When Mint Authority wants to set up the pool, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					// transfer mint authority to pool once admin registry is set
					ixAuth, err := tokens.SetTokenMintAuthority(token1.Program, token1.PoolSigner, token1.Mint, token1PoolAdmin.PublicKey()) // TODO: Check this
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction, ixAuth}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token1.PoolLookupTable, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("Transfer admin role for token admin registry", func(t *testing.T) {
				t.Run("When invalid wants to transfer the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})
				t.Run("When mint authority wants to transfer the token admin registry, it succeeds and permissions stay with no changes", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token1.PoolLookupTable, tokenAdminRegistry.LookupTable)

					// check if the admin is still the same
					instruction, err = ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// new one cant make changes yet
					instruction, err = ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When new admin accepts the token admin registry, it succeeds and permissions are updated", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, token0PoolAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, token1.PoolLookupTable, tokenAdminRegistry.LookupTable)

					// check old admin can not make changes anymore
					instruction, err = ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})

					// new one can make changes now
					instruction, err = ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint,
						token1.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)
				})
			})

			t.Run("setup: Link pool", func(t *testing.T) {
				t.Run("propose link pool as token mint authority", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						legacyAdmin.PublicKey(),
						config.RouterConfigPDA,
						linkPool.AdminRegistryPDA,
						linkPool.Mint,
						legacyAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, linkPool.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
					require.Equal(t, legacyAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})

				t.Run("accept token admin registry as token admin", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						linkPool.AdminRegistryPDA,
						linkPool.Mint,
						legacyAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, linkPool.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, legacyAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})

				t.Run("set pool", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						linkPool.WritableIndexes,
						config.RouterConfigPDA,
						linkPool.AdminRegistryPDA,
						linkPool.Mint,
						linkPool.PoolLookupTable,
						legacyAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					// transfer mint authority to pool once admin registry is set
					ixAuth, err := tokens.SetTokenMintAuthority(token1.Program, linkPool.PoolSigner, linkPool.Mint, legacyAdmin.PublicKey()) // TODO: Check this
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction, ixAuth}, legacyAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, linkPool.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, legacyAdmin.PublicKey(), tokenAdminRegistry.Administrator)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, linkPool.PoolLookupTable, tokenAdminRegistry.LookupTable)
				})
			})
		})
	})

	//////////////////////////////
	// Token Pool Config Tests //
	/////////////////////////////
	t.Run("Token Pool Configuration", func(t *testing.T) {
		t.Run("RemoteConfig", func(t *testing.T) {
			for _, selector := range []uint64{config.EvmChainSelector, config.SvmChainSelector} {
				ix0, err := test_token_pool.NewInitChainRemoteConfigInstruction(selector, token0.Mint, base_token_pool.RemoteConfig{
					TokenAddress: base_token_pool.RemoteAddress{Address: config.EVMToken0AddressBytes},
					Decimals:     evmToken0Decimals,
				}, token0.PoolConfig, token0.Chain[selector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				PoolAddresses := []base_token_pool.RemoteAddress{{Address: []byte{4, 5, 6}}}
				ix1, err := test_token_pool.NewInitChainRemoteConfigInstruction(selector, token1.Mint, base_token_pool.RemoteConfig{
					PoolAddresses: []base_token_pool.RemoteAddress{},
					TokenAddress:  base_token_pool.RemoteAddress{Address: config.EVMToken1AddressBytes},
					Decimals:      evmToken1Decimals,
				}, token1.PoolConfig, token1.Chain[selector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				ix2, err := test_token_pool.NewInitChainRemoteConfigInstruction(selector, token2.Mint, base_token_pool.RemoteConfig{
					PoolAddresses: []base_token_pool.RemoteAddress{},
					TokenAddress:  base_token_pool.RemoteAddress{Address: config.EVMToken2AddressBytes},
					Decimals:      evmToken2Decimals,
				}, token2.PoolConfig, token2.Chain[selector], token2PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				ixLink, err := test_token_pool.NewInitChainRemoteConfigInstruction(selector, linkPool.Mint, base_token_pool.RemoteConfig{
					PoolAddresses: []base_token_pool.RemoteAddress{},
					TokenAddress:  base_token_pool.RemoteAddress{Address: config.EVMToken3AddressBytes},
					Decimals:      evmToken3Decimals,
				}, linkPool.PoolConfig, linkPool.Chain[selector], legacyAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)

				ix3, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(selector, token1.Mint, PoolAddresses, token1.PoolConfig, token1.Chain[selector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				ix4, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(selector, token2.Mint, PoolAddresses, token2.PoolConfig, token2.Chain[selector], token2PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)

				// Note: Splitted into two transactions to avoid exceeding the Solana transaction size limit
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1, ix2}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin, token2PoolAdmin, legacyAdmin))
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix3, ix4, ixLink}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin, token2PoolAdmin, legacyAdmin))
			}
		})

		t.Run("AppendRemotePools", func(t *testing.T) {
			ixEvm, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(config.EvmChainSelector, token0.Mint, []base_token_pool.RemoteAddress{{Address: []byte{1, 2, 3}}},
				token0.PoolConfig, token0.Chain[config.EvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ixSvm, err := test_token_pool.NewAppendRemotePoolAddressesInstruction(config.SvmChainSelector, token0.Mint, []base_token_pool.RemoteAddress{{Address: []byte{1, 2, 3}}},
				token0.PoolConfig, token0.Chain[config.SvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixEvm, ixSvm}, token0PoolAdmin, config.DefaultCommitment)
		})

		t.Run("RateLimit", func(t *testing.T) {
			ix0, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, token0.Mint, base_token_pool.RateLimitConfig{}, base_token_pool.RateLimitConfig{}, token0.PoolConfig, token0.Chain[config.EvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix1, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, token1.Mint, base_token_pool.RateLimitConfig{}, base_token_pool.RateLimitConfig{}, token1.PoolConfig, token1.Chain[config.EvmChainSelector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix2, err := test_token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, token2.Mint, base_token_pool.RateLimitConfig{}, base_token_pool.RateLimitConfig{}, token2.PoolConfig, token2.Chain[config.EvmChainSelector], token2PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1, ix2}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin, token2PoolAdmin))
		})

		t.Run("Billing", func(t *testing.T) {
			defaultTokenTransferFeeConfig := fee_quoter.TokenTransferFeeConfig{DestBytesOverhead: 32, MinFeeUsdcents: 0, MaxFeeUsdcents: 1}
			ix0, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint, defaultTokenTransferFeeConfig, config.FqConfigPDA, token0.Billing[config.EvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix1, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token1.Mint, defaultTokenTransferFeeConfig, config.FqConfigPDA, token1.Billing[config.EvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix2, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.SvmChainSelector, token0.Mint, defaultTokenTransferFeeConfig, config.FqConfigPDA, token0.Billing[config.SvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix3, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.SvmChainSelector, token1.Mint, defaultTokenTransferFeeConfig, config.FqConfigPDA, token1.Billing[config.SvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			// Deliberately not setting configuration for token2, as it exists to test cases where the configuration doesn't exist, given that it can't
			// be removed (only disabled) after it's initially set.

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1, ix2, ix3}, ccipAdmin, config.DefaultCommitment)

			// check TokenTransferFeeConfigUpdated events
			parsedEvents, perr := common.ParseMultipleEvents[ccip.TokenTransferFeeConfigUpdated](result.Meta.LogMessages, "TokenTransferFeeConfigUpdated", config.PrintEvents)
			require.NoError(t, perr)

			require.Equal(t, config.EvmChainSelector, parsedEvents[0].DestinationChainSelector)
			require.Equal(t, token0.Mint, parsedEvents[0].Token)
			require.Equal(t, defaultTokenTransferFeeConfig, parsedEvents[0].TokenTransferFeeConfig)

			require.Equal(t, config.EvmChainSelector, parsedEvents[1].DestinationChainSelector)
			require.Equal(t, token1.Mint, parsedEvents[1].Token)
			require.Equal(t, defaultTokenTransferFeeConfig, parsedEvents[1].TokenTransferFeeConfig)

			require.Equal(t, config.SvmChainSelector, parsedEvents[2].DestinationChainSelector)
			require.Equal(t, token0.Mint, parsedEvents[2].Token)
			require.Equal(t, defaultTokenTransferFeeConfig, parsedEvents[2].TokenTransferFeeConfig)

			require.Equal(t, config.SvmChainSelector, parsedEvents[3].DestinationChainSelector)
			require.Equal(t, token1.Mint, parsedEvents[3].Token)
			require.Equal(t, defaultTokenTransferFeeConfig, parsedEvents[3].TokenTransferFeeConfig)
		})

		// validate permissions for setting config
		t.Run("Permissions", func(t *testing.T) {
			t.Parallel()
			t.Run("Billing can only be set by CCIP admin", func(t *testing.T) {
				ix, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint, fee_quoter.TokenTransferFeeConfig{DestBytesOverhead: 32}, config.FqConfigPDA, token0.Billing[config.EvmChainSelector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, token1PoolAdmin, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
			})
		})
	})

	/////////////////////////////
	//   Manual Cursing Tests  //
	/////////////////////////////
	t.Run("Manual Cursing", func(t *testing.T) {
		t.Run("no curses by default", func(t *testing.T) {
			svmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(svmCurse.Value[:], config.SvmChainSelector)
			evmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(evmCurse.Value[:], config.EvmChainSelector)

			ix, err := rmn_remote.NewVerifyNotCursedInstruction(
				evmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				svmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment)
			require.NotNil(t, result)
		})

		t.Run("applying a global curse", func(t *testing.T) {
			globalCurse := rmn_remote.CurseSubject{
				Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			}

			ix, err := rmn_remote.NewCurseInstruction(
				globalCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var curses rmn_remote.Curses
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 1)
			require.Equal(t, curses.CursedSubjects[0], globalCurse)

			// All subjects are cursed now
			svmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(svmCurse.Value[:], config.SvmChainSelector)
			evmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(evmCurse.Value[:], config.EvmChainSelector)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				evmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
			require.NotNil(t, result)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				svmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
			require.NotNil(t, result)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				globalCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
			require.NotNil(t, result)
		})

		t.Run("ccip_send fails with a global curse active", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
			require.NotNil(t, result)
		})

		t.Run("commit fails with a global curse active", func(t *testing.T) {
			report := ccip_offramp.CommitInput{
				MerkleRoot:   nil,
				PriceUpdates: ccip_offramp.PriceUpdates{},
			}

			reportContext := ccip.NextCommitReportContext()

			sigs, err := ccip.SignCommitReport(reportContext, report, signers)
			require.NoError(t, err)

			transmitter := getTransmitter()

			raw := ccip_offramp.NewCommitPriceOnlyInstruction(
				reportContext,
				testutils.MustMarshalBorsh(t, report),
				sigs.Rs,
				sigs.Ss,
				sigs.RawVs,
				config.OfframpConfigPDA,
				config.OfframpReferenceAddressesPDA,
				transmitter.PublicKey(),
				solana.SystemProgramID,
				solana.SysVarInstructionsPubkey,
				config.OfframpBillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqAllowedPriceUpdaterOfframpPDA,
				config.FqConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)

			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
			require.NotNil(t, result)
		})

		t.Run("removing a global curse", func(t *testing.T) {
			globalCurse := rmn_remote.CurseSubject{
				Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			}

			ix, err := rmn_remote.NewUncurseInstruction(
				globalCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var curses rmn_remote.Curses
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 0)
		})

		t.Run("adding chain selector curses", func(t *testing.T) {
			svmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(svmCurse.Value[:], config.SvmChainSelector)
			evmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(evmCurse.Value[:], config.EvmChainSelector)

			ix, err := rmn_remote.NewCurseInstruction(
				svmCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var curses rmn_remote.Curses
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 1)
			require.Equal(t, curses.CursedSubjects[0], svmCurse)

			ix, err = rmn_remote.NewCurseInstruction(
				evmCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 2)
			require.Equal(t, curses.CursedSubjects[0], svmCurse)
			require.Equal(t, curses.CursedSubjects[1], evmCurse)

			ix, err = rmn_remote.NewUncurseInstruction(
				svmCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 1)
			require.Equal(t, curses.CursedSubjects[0], evmCurse)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				evmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: SubjectCursed"})
			require.NotNil(t, result)

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				svmCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			globalCurse := rmn_remote.CurseSubject{
				Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			}

			ix, err = rmn_remote.NewVerifyNotCursedInstruction(
				globalCurse,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
		})

		t.Run("ccip_send fails with a cursed destination", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: SubjectCursed"})
			require.NotNil(t, result)
		})

		t.Run("cleanup", func(t *testing.T) {
			evmCurse := rmn_remote.CurseSubject{}
			binary.LittleEndian.PutUint64(evmCurse.Value[:], config.EvmChainSelector)
			ix, err := rmn_remote.NewUncurseInstruction(
				evmCurse,
				config.RMNRemoteConfigPDA,
				ccipAdmin.PublicKey(),
				config.RMNRemoteCursesPDA,
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var curses rmn_remote.Curses
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RMNRemoteCursesPDA, config.DefaultCommitment, &curses)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, len(curses.CursedSubjects), 0)
		})
	})

	//////////////////////////
	//     getFee Tests     //
	//////////////////////////
	t.Run("getFee", func(t *testing.T) {
		t.Run("Fee is retrieved for a correctly formatted message", func(t *testing.T) {
			message := fee_quoter.SVM2AnyMessage{
				Receiver:  validReceiverAddress[:],
				FeeToken:  wsol.mint,
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, message, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, feeResult)
			fee, _ := common.ExtractTypedReturnValue(ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String(), binary.LittleEndian.Uint64)
			require.Greater(t, fee, uint64(0))
		})

		t.Run("Fee is retrieved for a correctly formatted message containing a nonnative token", func(t *testing.T) {
			message := fee_quoter.SVM2AnyMessage{
				Receiver:     validReceiverAddress[:],
				FeeToken:     wsol.mint,
				TokenAmounts: []fee_quoter.SVMTokenAmount{{Token: token0.Mint, Amount: 1}},
				ExtraArgs:    emptyGenericExtraArgsV2,
			}

			// Set some fees that will result in some appreciable change in the message fee
			billing := fee_quoter.TokenTransferFeeConfig{
				MinFeeUsdcents:    800,
				MaxFeeUsdcents:    1600,
				DeciBps:           0,
				DestGasOverhead:   100,
				DestBytesOverhead: 100,
				IsEnabled:         true,
			}
			token0BillingConfigPda := getFqTokenConfigPDA(token0.Mint)
			token0PerChainPerConfigPda := getFqPerChainPerTokenConfigBillingPDA(token0.Mint)
			ix, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint, billing, config.FqConfigPDA, token0PerChainPerConfigPda, ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)

			// check TokenTransferFeeConfigUpdated event
			tokenTransferFeeConfigUpdatedEvent := ccip.TokenTransferFeeConfigUpdated{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "TokenTransferFeeConfigUpdated", &tokenTransferFeeConfigUpdatedEvent, config.PrintEvents))

			require.Equal(t, config.EvmChainSelector, tokenTransferFeeConfigUpdatedEvent.DestinationChainSelector)
			require.Equal(t, token0.Mint, tokenTransferFeeConfigUpdatedEvent.Token)
			require.Equal(t, billing.MaxFeeUsdcents, tokenTransferFeeConfigUpdatedEvent.TokenTransferFeeConfig.MaxFeeUsdcents)

			raw := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, message, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
			raw.AccountMetaSlice.Append(solana.Meta(token0BillingConfigPda))
			raw.AccountMetaSlice.Append(solana.Meta(token0PerChainPerConfigPda))
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, feeResult)
			fee, _ := common.ExtractTypedReturnValue(ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String(), binary.LittleEndian.Uint64)
			require.Greater(t, fee, uint64(0))
		})

		t.Run("Fee is retrieved for a correctly formatted message containing a nonnative token with no billing config", func(t *testing.T) {
			message := fee_quoter.SVM2AnyMessage{
				Receiver:     validReceiverAddress[:],
				FeeToken:     wsol.mint,
				TokenAmounts: []fee_quoter.SVMTokenAmount{{Token: token2.Mint, Amount: 1}},
				ExtraArgs:    emptyGenericExtraArgsV2,
			}

			// Token 2 is not configured with any fee overrides
			token2BillingConfigPda := getFqTokenConfigPDA(token2.Mint)
			token2PerChainPerConfigPda := getFqPerChainPerTokenConfigBillingPDA(token2.Mint)

			raw := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, message, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
			raw.AccountMetaSlice.Append(solana.Meta(token2BillingConfigPda))
			raw.AccountMetaSlice.Append(solana.Meta(token2PerChainPerConfigPda))
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, feeResult)
			fee, _ := common.ExtractTypedReturnValue(ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String(), binary.LittleEndian.Uint64)
			require.Greater(t, fee, uint64(0))
		})

		t.Run("Cannot get fee for message with invalid address", func(t *testing.T) {
			// Bigger than u160
			tooBigAddress := [32]byte{}
			tooBigAddress[11] = 1

			// Falls within precompile region
			tooSmallAddress := [32]byte{}
			tooSmallAddress[31] = 1

			for _, address := range [][32]byte{tooBigAddress, tooSmallAddress} {
				message := fee_quoter.SVM2AnyMessage{
					Receiver:  address[:],
					FeeToken:  wsol.mint,
					ExtraArgs: emptyGenericExtraArgsV2,
				}

				raw := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, message, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
				instruction, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: InvalidEVMAddress"})
				require.NotNil(t, result)
			}
		})

		t.Run("Can retrieve fee through router CPI", func(t *testing.T) {
			message := ccip_router.SVM2AnyMessage{
				Receiver:  validReceiverAddress[:],
				FeeToken:  wsol.mint,
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewGetFeeInstruction(config.EvmChainSelector,
				message,
				config.RouterConfigPDA,
				config.EvmDestChainStatePDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
			)
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, feeResult)

			// Check first that FQ returned something valid to the router
			fqFee, err := common.ExtractAnchorTypedReturnValue[fee_quoter.GetFeeResult](ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String())
			require.NoError(t, err)
			require.Less(t, uint64(0), fqFee.Amount)
			require.Equal(t, wsol.mint, fqFee.Token)

			// Check that the router's response matches what FQ returned (with fewer fields)
			routerFee, err := common.ExtractAnchorTypedReturnValue[ccip_router.GetFeeResult](ctx, feeResult.Meta.LogMessages, config.CcipRouterProgram.String())
			require.NoError(t, err)
			require.Equal(t, fqFee.Token, routerFee.Token)
			require.Equal(t, fqFee.Amount, routerFee.Amount)
			require.Equal(t, fqFee.Juels, routerFee.Juels)
		})
	})

	//////////////////////////
	//    ccipSend Tests    //
	//////////////////////////

	t.Run("OnRamp ccipSend", func(t *testing.T) {
		t.Run("When sending to an invalid destination chain selector it fails", func(t *testing.T) {
			destinationChainSelector := uint64(189)
			destinationChainStatePDA, err := state.FindDestChainStatePDA(destinationChainSelector, config.CcipRouterProgram)
			require.NoError(t, err)
			fqDestChainPDA, _, err := state.FindFqDestChainPDA(destinationChainSelector, config.FeeQuoterProgram)
			require.NoError(t, err)
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				fqDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()

			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + common.AccountNotInitialized_AnchorError.String()})
			require.NotNil(t, result)
		})

		t.Run("When sending with an empty but enabled allowlist, it fails", func(t *testing.T) {
			var initialDestChain ccip_router.DestChain
			err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmDestChainStatePDA, config.DefaultCommitment, &initialDestChain)
			require.NoError(t, err, "failed to get account info")
			modifiedDestChain := initialDestChain
			modifiedDestChain.Config.AllowListEnabled = true

			updateDestChainIx, err := ccip_router.NewUpdateDestChainConfigInstruction(
				config.EvmChainSelector,
				modifiedDestChain.Config,
				config.EvmDestChainStatePDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{updateDestChainIx}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: SenderNotAllowed"})
			require.NotNil(t, result)

			// We now restore the config to keep the test state-neutral
			updateDestChainIx, err = ccip_router.NewUpdateDestChainConfigInstruction(
				config.EvmChainSelector,
				initialDestChain.Config,
				config.EvmDestChainStatePDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{updateDestChainIx}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
		})

		t.Run("When sending a Valid CCIP Message Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(1), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonceEvmPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(1), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(1), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, user.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: uint64(validFqDestChainConfig.DefaultTxGasLimit), Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit) // default gas limit
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)                                                    // default OOO Execution
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(1), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(1), ccipMessageSentEvent.Message.Header.Nonce)
			hash, err := ccip.HashSVMToAnyMessage(ccipMessageSentEvent.Message)
			require.NoError(t, err)
			require.Equal(t, hash, ccipMessageSentEvent.Message.Header.MessageId[:])
		})

		t.Run("When sending a CCIP Message with ExtraArgs overrides Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			trueValue := true
			message := ccip_router.SVM2AnyMessage{
				FeeToken: wsol.mint,
				Receiver: validReceiverAddress[:],
				Data:     []byte{4, 5, 6},
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.GenericExtraArgsV2{
					GasLimit:                 bin.Uint128{Lo: 99, Hi: 0},
					AllowOutOfOrderExecution: trueValue,
				}, ccip.GenericExtraArgsV2Tag),
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(2), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonceEvmPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(1), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(2), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, user.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: 99, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit) // check it's overwritten
			require.Equal(t, true, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)       // check it's overwritten
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(2), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(0), ccipMessageSentEvent.Message.Header.Nonce) // nonce is not incremented as it is OOO
		})

		t.Run("When sending a CCIP Message with only gas limit ExtraArgs overrides Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken: wsol.mint,
				Receiver: validReceiverAddress[:],
				Data:     []byte{4, 5, 6},
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.GenericExtraArgsV2{
					GasLimit: bin.Uint128{Lo: 99, Hi: 0},
				}, ccip.GenericExtraArgsV2Tag),
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(3), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonceEvmPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(2), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(3), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, user.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: 99, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit) // check it's overwritten
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)      // check it's default value
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(3), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(2), ccipMessageSentEvent.Message.Header.Nonce) // nonce is incremented
		})

		t.Run("When sending a CCIP Message with allow out of order ExtraArgs overrides Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken: wsol.mint,
				Receiver: validReceiverAddress[:],
				Data:     []byte{4, 5, 6},
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.GenericExtraArgsV2{
					AllowOutOfOrderExecution: true,
					GasLimit:                 bin.Uint128{Lo: 5000, Hi: 0},
				}, ccip.GenericExtraArgsV2Tag),
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(4), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonceEvmPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(2), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(4), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, user.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: 5000, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit) // default gas limit
			require.Equal(t, true, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)         // check it's overwritten
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(4), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(0), ccipMessageSentEvent.Message.Header.Nonce) // nonce is not incremented as it is OOO
		})

		t.Run("When gasLimit is set to zero, it overrides Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken: link22.mint,
				Receiver: validReceiverAddress[:],
				Data:     []byte{4, 5, 6},
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.GenericExtraArgsV2{
					GasLimit: bin.Uint128{Lo: 0, Hi: 0},
				}, ccip.GenericExtraArgsV2Tag),
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				link22.program,
				link22.mint,
				link22.userATA,
				link22.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				link22.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(5), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, nonceEvmPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(3), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(5), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, user.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: 0, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit)
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(5), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(3), ccipMessageSentEvent.Message.Header.Nonce)
		})

		t.Run("When sending a message with an invalid nonce account, it fails", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				anotherUser.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherUser, config.DefaultCommitment, []string{"Error Message: A seeds constraint was violated"})
			require.NotNil(t, result)
		})

		t.Run("When sending a message impersonating another user, it fails", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWithRPCError(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherUser, config.DefaultCommitment, []string{"Transaction signature verification failure"})
		})

		t.Run("When sending a message without flagging the user ATA as writable, it fails", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)

			// do NOT mark the user ATA as writable

			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWithRPCError(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.InvalidInputsAtaWritable_CcipRouterError.String()})
		})

		t.Run("When sending a message and paying with inconsistent fee token accounts, it fails", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA

			// These testcases are a quite a lot, this obviously blows up combinatorially and adds many seconds to the suite.
			// We can remove/reduce this, but I used it during development so for now I'm keeping them here
			for i, program := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.program }) {
				for j, mint := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.mint }) {
					for k, messageMint := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.mint }) {
						for l, billingConfigPDA := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.fqBillingConfigPDA }) {
							for m, userATA := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.userATA }) {
								for n, billingATA := range common.Map(billingTokens, func(t *AccountsPerToken) solana.PublicKey { return t.billingATA }) {
									if i == j && j == k && k == l && l == m && m == n {
										// skip cases where everything aligns well, which work
										continue
									}
									testName := fmt.Sprintf("when using program %v, mint %v, message mint %v, configPDA %v, userATA %v, billingATA %v", i, j, k, l, m, n)
									t.Run(testName, func(t *testing.T) {
										t.Parallel()
										raw := ccip_router.NewCcipSendInstruction(
											destinationChainSelector,
											ccip_router.SVM2AnyMessage{
												FeeToken:  messageMint,
												Receiver:  validReceiverAddress[:],
												Data:      []byte{4, 5, 6},
												ExtraArgs: emptyGenericExtraArgsV2,
											},
											[]byte{},
											config.RouterConfigPDA,
											destinationChainStatePDA,
											nonceEvmPDA,
											user.PublicKey(),
											solana.SystemProgramID,
											program,
											mint,
											userATA,
											billingATA,
											config.BillingSignerPDA,
											config.FeeQuoterProgram,
											config.FqConfigPDA,
											config.FqEvmDestChainPDA,
											billingConfigPDA,
											link22.fqBillingConfigPDA,
											config.RMNRemoteProgram,
											config.RMNRemoteCursesPDA,
											config.RMNRemoteConfigPDA,
										)
										raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
										instruction, err := raw.ValidateAndBuild()
										require.NoError(t, err)

										// Given the mixture of inputs, there can be different error types here, so just check that it fails but not each message
										testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{""})
									})
								}
							}
						}
					}
				}
			}
		})

		t.Run("When another user sending a Valid CCIP Message tries to pay with some else's tokens it fails", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  link22.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}
			anotherUserNonceEVMPDA, err := state.FindNoncePDA(config.EvmChainSelector, anotherUser.PublicKey(), config.CcipRouterProgram)
			require.NoError(t, err)

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				anotherUserNonceEVMPDA,
				anotherUser.PublicKey(),
				solana.SystemProgramID,
				link22.program,
				link22.mint,
				link22.userATA, // token account of a different user
				link22.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				link22.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherUser, config.DefaultCommitment, []string{ccip.InvalidInputsAtaAddress_CcipRouterError.String()})
		})

		t.Run("When another user sending a Valid CCIP Message Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  link22.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}
			anotherUserNonceEVMPDA, err := state.FindNoncePDA(config.EvmChainSelector, anotherUser.PublicKey(), config.CcipRouterProgram)
			require.NoError(t, err)

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				anotherUserNonceEVMPDA,
				anotherUser.PublicKey(),
				solana.SystemProgramID,
				link22.program,
				link22.mint,
				link22.anotherUserATA,
				link22.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				link22.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherUser, config.DefaultCommitment)
			require.NotNil(t, result)

			var chainStateAccount ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, destinationChainStatePDA, config.DefaultCommitment, &chainStateAccount)
			require.NoError(t, err, "failed to get account info")
			// Do not check source chain config, as it may have been updated by other tests in ccip offramp
			require.Equal(t, uint64(6), chainStateAccount.State.SequenceNumber)

			var nonceCounterAccount ccip_router.Nonce
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, anotherUserNonceEVMPDA, config.DefaultCommitment, &nonceCounterAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, uint64(1), nonceCounterAccount.Counter)

			ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
			require.Equal(t, uint64(21), ccipMessageSentEvent.DestinationChainSelector)
			require.Equal(t, uint64(6), ccipMessageSentEvent.SequenceNumber)
			require.Equal(t, anotherUser.PublicKey(), ccipMessageSentEvent.Message.Sender)
			require.Equal(t, validReceiverAddress[:], ccipMessageSentEvent.Message.Receiver)
			data := [3]uint8{4, 5, 6}
			require.Equal(t, data[:], ccipMessageSentEvent.Message.Data)
			require.Equal(t, bin.Uint128{Lo: uint64(validFqDestChainConfig.DefaultTxGasLimit), Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).GasLimit)
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.GenericExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.GenericExtraArgsV2Tag).AllowOutOfOrderExecution)
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(6), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(1), ccipMessageSentEvent.Message.Header.Nonce)
		})

		t.Run("token happy path", func(t *testing.T) {
			t.Run("single token", func(t *testing.T) {
				_, initSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint, config.DefaultCommitment)
				require.NoError(t, err)
				_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)

				destinationChainSelector := config.EvmChainSelector
				destinationChainStatePDA := config.EvmDestChainStatePDA
				message := ccip_router.SVM2AnyMessage{
					FeeToken: wsol.mint,
					Receiver: validReceiverAddress[:],
					Data:     []byte{4, 5, 6},
					TokenAmounts: []ccip_router.SVMTokenAmount{
						{
							Token:  token0.Mint,
							Amount: 1,
						},
					},
					ExtraArgs: emptyGenericExtraArgsV2,
				}

				userTokenAccount, ok := token0.User[user.PublicKey()]
				require.True(t, ok)

				base := ccip_router.NewCcipSendInstruction(
					destinationChainSelector,
					message,
					[]byte{0}, // starting indices for accounts
					config.RouterConfigPDA,
					destinationChainStatePDA,
					nonceEvmPDA,
					user.PublicKey(),
					solana.SystemProgramID,
					wsol.program,
					wsol.mint,
					wsol.userATA,
					wsol.billingATA,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					wsol.fqBillingConfigPDA,
					link22.fqBillingConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				base.GetFeeTokenUserAssociatedAccountAccount().WRITE()

				tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, userTokenAccount)
				require.NoError(t, err)
				base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas...)

				ix, err := base.ValidateAndBuild()
				require.NoError(t, err)

				ixApprove, err := tokens.TokenApproveChecked(1, token0Decimals, token0.Program, userTokenAccount, token0.Mint, config.BillingSignerPDA, user.PublicKey(), nil)
				require.NoError(t, err)

				result := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove, ix}, user, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(300_000))
				require.NotNil(t, result)

				// check CCIP event
				ccipMessageSentEvent := ccip.EventCCIPMessageSent{}
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, config.PrintEvents))
				require.Equal(t, 1, len(ccipMessageSentEvent.Message.TokenAmounts))
				ta := ccipMessageSentEvent.Message.TokenAmounts[0]

				require.Equal(t, wsol.mint, ccipMessageSentEvent.Message.FeeToken)
				require.Equal(t, tokens.ToLittleEndianU256(36333028), ccipMessageSentEvent.Message.FeeTokenAmount.LeBytes)
				// The difference is the ratio between the fee token value (wsol) and link token value (signified by link22 in these tests).
				// Since they have been configured in the test setup to differ by a factor of 10, so does the token amount and its value in juels.
				// Then, they differ again by 9 decimals due to the solana denomination being 1e9 divisions = 1 LINK, compared to 1e18 juels = 1 LINK
				// in EVM. Note how some precision is lost in the translation, because the price in solana is stored with respect to the native
				// decimals.
				require.Equal(t, tokens.ToLittleEndianU256(3633302000000000), ccipMessageSentEvent.Message.FeeValueJuels.LeBytes)
				require.Equal(t, token0.PoolConfig, ta.SourcePoolAddress)
				require.Equal(t, config.EVMToken0AddressBytes, ta.DestTokenAddress)
				// Local decimals are encoded in the extra data. By default, 9 decimals in Solana.
				expectedExtraData := make([]byte, 32)
				expectedExtraData[31] = token0Decimals
				require.Equal(t, 32, len(ta.ExtraData))
				require.Equal(t, expectedExtraData, ta.ExtraData)

				require.Equal(t, tokens.ToLittleEndianU256(1), ta.Amount.LeBytes)
				require.Equal(t, 4, len(ta.DestExecData))
				expectedDestExecData := make([]byte, 4)
				// Token0 billing had DestGasOverhead set to 100 during setup
				binary.BigEndian.PutUint32(expectedDestExecData[:], 100)
				require.Equal(t, expectedDestExecData, ta.DestExecData)

				// check pool event
				poolEvent := tokens.EventBurnLock{}
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "Burned", &poolEvent, config.PrintEvents))
				require.Equal(t, token0.RouterSigner, poolEvent.Sender)
				require.Equal(t, uint64(1), poolEvent.Amount)
				require.Equal(t, token0.Mint, poolEvent.Mint)

				// check balances
				_, currSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1, initSupply-currSupply) // burned amount
				_, currBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1, initBal-currBal) // burned amount
				_, poolBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.PoolTokenAccount, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 0, poolBal) // pool burned any sent to it
			})

			t.Run("two tokens", func(t *testing.T) {
				_, initBal0, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				_, initBal1, err := tokens.TokenBalance(ctx, solanaGoClient, token1.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)

				destinationChainSelector := config.EvmChainSelector
				destinationChainStatePDA := config.EvmDestChainStatePDA
				message := ccip_router.SVM2AnyMessage{
					FeeToken: wsol.mint,
					Receiver: validReceiverAddress[:],
					Data:     []byte{4, 5, 6},
					TokenAmounts: []ccip_router.SVMTokenAmount{
						{
							Token:  token0.Mint,
							Amount: 1,
						},
						{
							Token:  token1.Mint,
							Amount: 2,
						},
					},
					ExtraArgs: emptyGenericExtraArgsV2,
				}

				userTokenAccount0, ok := token0.User[user.PublicKey()]
				require.True(t, ok)
				userTokenAccount1, ok := token1.User[user.PublicKey()]
				require.True(t, ok)

				base := ccip_router.NewCcipSendInstruction(
					destinationChainSelector,
					message,
					[]byte{0, 14}, // starting indices for accounts
					config.RouterConfigPDA,
					destinationChainStatePDA,
					nonceEvmPDA,
					user.PublicKey(),
					solana.SystemProgramID,
					wsol.program,
					wsol.mint,
					wsol.userATA,
					wsol.billingATA,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
					config.FqEvmDestChainPDA,
					wsol.fqBillingConfigPDA,
					link22.fqBillingConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				base.GetFeeTokenUserAssociatedAccountAccount().WRITE()

				tokenMetas0, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, userTokenAccount0)
				require.NoError(t, err)
				base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas0...)
				tokenMetas1, addressTables1, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token1, userTokenAccount1)
				require.NoError(t, err)
				base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas1...)
				addressTables[token1.PoolLookupTable] = addressTables1[token1.PoolLookupTable]
				for k, v := range ccipSendLookupTable {
					addressTables[k] = v
				}

				ix, err := base.ValidateAndBuild()
				require.NoError(t, err)

				ixApprove0, err := tokens.TokenApproveChecked(1, token0Decimals, token0.Program, userTokenAccount0, token0.Mint, config.BillingSignerPDA, user.PublicKey(), nil)
				require.NoError(t, err)
				ixApprove1, err := tokens.TokenApproveChecked(2, token1Decimals, token1.Program, userTokenAccount1, token1.Mint, config.BillingSignerPDA, user.PublicKey(), nil)
				require.NoError(t, err)

				result := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove0, ixApprove1, ix}, user, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(800_000))
				require.NotNil(t, result)

				// check balances
				_, currBal0, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1, initBal0-currBal0) // burned amount
				_, currBal1, err := tokens.TokenBalance(ctx, solanaGoClient, token1.User[user.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 2, initBal1-currBal1) // burned amount
			})
		})

		t.Run("When sending with an enabled allowlist including the sender, it succeeds", func(t *testing.T) {
			var initialDestChain ccip_router.DestChain
			err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmDestChainStatePDA, config.DefaultCommitment, &initialDestChain)
			require.NoError(t, err, "failed to get account info")
			modifiedDestChain := initialDestChain
			modifiedDestChain.Config.AllowListEnabled = true
			modifiedDestChain.Config.AllowedSenders = []solana.PublicKey{
				user.PublicKey(),
				anotherUser.PublicKey(),
			}

			updateDestChainIx, err := ccip_router.NewUpdateDestChainConfigInstruction(
				config.EvmChainSelector,
				modifiedDestChain.Config,
				config.EvmDestChainStatePDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{updateDestChainIx}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var parsedDestChain ccip_router.DestChain
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmDestChainStatePDA, config.DefaultCommitment, &parsedDestChain)
			require.NoError(t, err, "failed to get account info")

			// This proves we're able to update the config with a dynamically sized element
			require.Equal(t, parsedDestChain.Config.AllowedSenders, modifiedDestChain.Config.AllowedSenders)

			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  wsol.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			raw := ccip_router.NewCcipSendInstruction(
				destinationChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				destinationChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				wsol.userATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			// We now restore the config to keep the test state-neutral
			updateDestChainIx, err = ccip_router.NewUpdateDestChainConfigInstruction(
				config.EvmChainSelector,
				initialDestChain.Config,
				config.EvmDestChainStatePDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{updateDestChainIx}, ccipAdmin, config.DefaultCommitment)
			require.NotNil(t, result)
		})

		t.Run("token pool accounts: validation", func(t *testing.T) {
			t.Parallel()
			// base transaction
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken: wsol.mint,
				Receiver: validReceiverAddress[:],
				Data:     []byte{4, 5, 6},
				TokenAmounts: []ccip_router.SVMTokenAmount{
					{
						Token:  token0.Mint,
						Amount: 1,
					},
				},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			userTokenAccount, ok := token0.User[user.PublicKey()]
			require.True(t, ok)

			inputs := []struct {
				name        string
				index       uint
				replaceWith *solana.AccountMeta // default to zero address
				errorStr    string
			}{
				{
					name:     "incorrect user token account",
					index:    0,
					errorStr: common.AccountOwnedByWrongProgram_AnchorError.String(),
				},
				{
					name:     "invalid billing config",
					index:    1,
					errorStr: common.ConstraintSeeds_AnchorError.String(),
				},
				{
					name:     "invalid token pool chain config",
					index:    2,
					errorStr: common.ConstraintSeeds_AnchorError.String(),
				},
				{
					name:     "pool config is not owned by pool program",
					index:    6,
					errorStr: common.ConstraintSeeds_AnchorError.String(),
				},
				{
					name:        "is pool config but for wrong token",
					index:       6,
					replaceWith: solana.Meta(token1.PoolConfig),
					errorStr:    common.ConstraintSeeds_AnchorError.String(),
				},
				{
					name:        "is pool config but missing write permissions",
					index:       6,
					replaceWith: solana.Meta(token0.PoolConfig),
					errorStr:    ccip.InvalidInputsLookupTableAccountWritable_CcipRouterError.String(),
				},
				{
					name:        "is pool lookup table but has write permissions",
					index:       3,
					replaceWith: solana.Meta(token0.PoolLookupTable).WRITE(),
					errorStr:    ccip.InvalidInputsLookupTableAccountWritable_CcipRouterError.String(),
				},
				{
					name:     "incorrect pool signer",
					index:    8,
					errorStr: ccip.InvalidInputsTokenAccounts_CcipRouterError.String(),
				},
				{
					name:     "invalid token program",
					index:    9,
					errorStr: common.InvalidProgramId_AnchorError.String(),
				},
				{
					name:     "incorrect pool token account",
					index:    7,
					errorStr: common.AccountOwnedByWrongProgram_AnchorError.String(),
				},
				{
					name:        "incorrect token pool lookup table",
					index:       3,
					replaceWith: solana.Meta(token1.PoolLookupTable),
					errorStr:    ccip.InvalidInputsLookupTableAccounts_CcipRouterError.String(),
				},
				{
					name:     "invalid fee token config",
					index:    11,
					errorStr: common.ConstraintSeeds_AnchorError.String(),
				},
				{
					name:     "extra accounts not in lookup table",
					index:    1_000, // large number to indicate append
					errorStr: ccip.InvalidInputsLookupTableAccounts_CcipRouterError.String(),
				},
				{
					name:     "remaining accounts mismatch",
					index:    13, // only works with token0
					errorStr: ccip.InvalidInputsLookupTableAccounts_CcipRouterError.String(),
				},
				{
					name:     "invalid token pool signer",
					index:    12,
					errorStr: common.ConstraintSeeds_AnchorError.String(),
				},
			}

			for _, in := range inputs {
				t.Run(in.name, func(t *testing.T) {
					t.Parallel()
					tx := ccip_router.NewCcipSendInstruction(
						destinationChainSelector,
						message,
						[]byte{0}, // starting indices for accounts
						config.RouterConfigPDA,
						destinationChainStatePDA,
						nonceEvmPDA,
						user.PublicKey(),
						solana.SystemProgramID,
						wsol.program,
						wsol.mint,
						wsol.userATA,
						wsol.billingATA,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
						config.FqEvmDestChainPDA,
						wsol.fqBillingConfigPDA,
						link22.fqBillingConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					tx.GetFeeTokenUserAssociatedAccountAccount().WRITE()

					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, userTokenAccount)
					require.NoError(t, err)
					// replace account meta with invalid account to trigger error or append
					if in.replaceWith == nil {
						in.replaceWith = solana.Meta(solana.PublicKey{}) // default 0 address
					}
					if in.index >= uint(len(tokenMetas)) {
						tokenMetas = append(tokenMetas, in.replaceWith)
					} else {
						tokenMetas[in.index] = in.replaceWith
					}

					tx.AccountMetaSlice = append(tx.AccountMetaSlice, tokenMetas...)
					ix, err := tx.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, addressTables, []string{in.errorStr})
				})
			}
		})

		toFqMsg := func(msg ccip_router.SVM2AnyMessage) fee_quoter.SVM2AnyMessage {
			fqTokenAmounts := make([]fee_quoter.SVMTokenAmount, len(msg.TokenAmounts))
			for i, ta := range msg.TokenAmounts {
				fqTokenAmounts[i] = fee_quoter.SVMTokenAmount{
					Token:  ta.Token,
					Amount: ta.Amount,
				}
			}
			return fee_quoter.SVM2AnyMessage{
				Receiver:     msg.Receiver,
				Data:         msg.Data,
				TokenAmounts: fqTokenAmounts,
				FeeToken:     msg.FeeToken,
				ExtraArgs:    msg.ExtraArgs,
			}
		}

		t.Run("When sending a Valid CCIP Message it bills the amount that getFee previously returned", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA

			for _, token := range billingTokens {
				t.Run("using "+token.name, func(t *testing.T) {
					message := ccip_router.SVM2AnyMessage{
						FeeToken:  token.mint,
						Receiver:  validReceiverAddress[:],
						Data:      []byte{4, 5, 6},
						ExtraArgs: emptyGenericExtraArgsV2,
					}
					fqMsg := toFqMsg(message)
					rawGetFeeIx := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, fqMsg, config.FqConfigPDA, config.FqEvmDestChainPDA, token.fqBillingConfigPDA, link22.fqBillingConfigPDA)
					ix, err := rawGetFeeIx.ValidateAndBuild()
					require.NoError(t, err)

					feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment)
					require.NotNil(t, feeResult)
					fmt.Println(feeResult.Meta.LogMessages)
					fee, err := common.ExtractAnchorTypedReturnValue[fee_quoter.GetFeeResult](ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String())
					require.NoError(t, err)
					require.Greater(t, fee.Amount, uint64(0))

					initialBalance := getBalance(token.billingATA)

					// ccipSend
					raw := ccip_router.NewCcipSendInstruction(
						destinationChainSelector,
						message,
						[]byte{},
						config.RouterConfigPDA,
						destinationChainStatePDA,
						nonceEvmPDA,
						user.PublicKey(),
						solana.SystemProgramID,
						token.program,
						token.mint,
						token.userATA,
						token.billingATA,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
						config.FqEvmDestChainPDA,
						token.fqBillingConfigPDA,
						link22.fqBillingConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
					instruction, err := raw.ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
					require.NotNil(t, result)

					finalBalance := getBalance(token.billingATA)

					// Check that the billing receiver account balance has increased by the fee amount
					require.Equal(t, fee.Amount, finalBalance-initialBalance)
				})
			}
		})

		t.Run("When sending a Valid CCIP Message but the user does not have enough funds of the fee token, it fails", func(t *testing.T) {
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  link22.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}

			noncePDA, err := state.FindNoncePDA(config.EvmChainSelector, tokenlessUser.PublicKey(), config.CcipRouterProgram)
			require.NoError(t, err)

			// ccipSend
			raw := ccip_router.NewCcipSendInstruction(
				config.EvmChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				config.EvmDestChainStatePDA,
				noncePDA,
				tokenlessUser.PublicKey(), // this user has 0 link22 balance, though they've approved the transfer
				solana.SystemProgramID,
				link22.program,
				link22.mint,
				link22.tokenlessUserATA,
				link22.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				link22.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, tokenlessUser, config.DefaultCommitment, []string{"insufficient funds"})
		})

		t.Run("When sending a valid CCIP message and paying in native SOL, it bills the same amount that getFee previously returned and it's accumulated as Wrapped SOL", func(t *testing.T) {
			getLamports := func(account solana.PublicKey) uint64 {
				out, err := solanaGoClient.GetBalance(ctx, account, rpc.CommitmentConfirmed)
				require.NoError(t, err)
				return out.Value
			}

			zeroPubkey := solana.PublicKeyFromBytes(make([]byte, 32))

			message := ccip_router.SVM2AnyMessage{
				FeeToken:  zeroPubkey, // will pay with native SOL
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyGenericExtraArgsV2,
			}
			fqMsg := toFqMsg(message)

			// getFee
			rawGetFeeIx := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, fqMsg, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
			ix, err := rawGetFeeIx.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SimulateTransaction(ctx, t, solanaGoClient, []solana.Instruction{ix}, user)
			require.NotNil(t, feeResult)
			fee, err := common.ExtractAnchorTypedReturnValue[fee_quoter.GetFeeResult](ctx, feeResult.Value.Logs, config.FeeQuoterProgram.String())
			require.NoError(t, err)
			require.Greater(t, fee.Amount, uint64(0))

			initialBalance := getBalance(wsol.billingATA)
			initialLamports := getLamports(user.PublicKey())

			// ccipSend
			raw := ccip_router.NewCcipSendInstruction(
				config.EvmChainSelector,
				message,
				[]byte{},
				config.RouterConfigPDA,
				config.EvmDestChainStatePDA,
				nonceEvmPDA,
				user.PublicKey(),
				solana.SystemProgramID,
				wsol.program,
				wsol.mint,
				zeroPubkey, // no user token account, because paying with native SOL
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)

			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
			require.NotNil(t, result)

			finalBalance := getBalance(wsol.billingATA)
			finalLamports := getLamports(user.PublicKey())

			// Check that the billing receiver account balance has increased by the fee amount
			require.Equal(t, fee.Amount, finalBalance-initialBalance)

			// Check that the user has paid for the tx cost and the ccip fee from their SOL
			require.Equal(t, fee.Amount+result.Meta.Fee, initialLamports-finalLamports)

			// although payment was in native SOL, this is considered equivalent to wsol
			// and the CCIP protocol expects an SPL token always
			var event ccip.EventCCIPMessageSent
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &event, config.PrintEvents))
			require.Equal(t, wsol.mint, event.Message.FeeToken)
		})
	})

	///////////////////////////
	// CCIP Sender Contract //
	//////////////////////////
	t.Run("Ccip Sender Contract", func(t *testing.T) {
		// addresses are not propagated as they are limited to the example sender only
		senderState, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, config.CcipBaseSender)
		require.NoError(t, err)
		senderPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_sender")}, config.CcipBaseSender)
		require.NoError(t, err)
		senderEvmDestChainConfigPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_chain_config"), common.Uint64ToLE(config.EvmChainSelector)}, config.CcipBaseSender)
		require.NoError(t, err)
		senderSvmDestChainConfigPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_chain_config"), common.Uint64ToLE(config.SvmChainSelector)}, config.CcipBaseSender)
		require.NoError(t, err)
		wsolATAIx, wsolSenderATA, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, solana.SolMint, senderPDA, user.PublicKey())
		require.NoError(t, err)
		link22ATAIx, link22SenderATA, err := tokens.CreateAssociatedTokenAccount(config.Token2022Program, link22.mint, senderPDA, user.PublicKey())
		require.NoError(t, err)
		senderEvmNoncePDA, err := state.FindNoncePDA(config.EvmChainSelector, senderPDA, config.CcipRouterProgram)
		require.NoError(t, err)
		senderSvmNoncePDA, err := state.FindNoncePDA(config.SvmChainSelector, senderPDA, config.CcipRouterProgram)
		require.NoError(t, err)

		token0ATAIx, token0SenderATA, err := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint, senderPDA, user.PublicKey())
		require.NoError(t, err)
		token1ATAIx, token1SenderATA, err := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint, senderPDA, user.PublicKey())
		require.NoError(t, err)

		t.Run("setup", func(t *testing.T) {
			initIx, err := example_ccip_sender.NewInitializeInstruction(config.CcipRouterProgram, senderState, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			evmDestChainIx, err := example_ccip_sender.NewInitChainConfigInstruction(config.EvmChainSelector, validReceiverAddress[:], emptyGenericExtraArgsV2, senderState, senderEvmDestChainConfigPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			svmDestChainIx, err := example_ccip_sender.NewInitChainConfigInstruction(config.SvmChainSelector, validReceiverAddress[:], testutils.MustSerializeExtraArgs(t, fee_quoter.SVMExtraArgsV1{
				AllowOutOfOrderExecution: true,
			}, ccip.SVMExtraArgsV1Tag), senderState, senderSvmDestChainConfigPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)

			transferSolIx, err := system.NewTransferInstruction(1_000_000_000, user.PublicKey(), senderPDA).ValidateAndBuild()
			require.NoError(t, err)

			approveLinkIx, err := tokens.TokenApproveChecked(1e9, 9, link22.program, link22.userATA, link22.mint, senderPDA, user.PublicKey(), nil)
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{transferSolIx, wsolATAIx, approveLinkIx, link22ATAIx, initIx, evmDestChainIx, svmDestChainIx, token0ATAIx, token1ATAIx}, user, rpc.CommitmentConfirmed)
		})

		feeConfig := []struct {
			name                                                                                          string
			feeToken, userATA, feeProgram, feeMint, feeSenderATA, feeBillingATA, feeTokenBillingConfigPDA solana.PublicKey
		}{
			{
				name:                     "sol",
				feeToken:                 solana.PublicKey{}, // native SOL
				userATA:                  wsol.userATA,
				feeProgram:               wsol.program,
				feeMint:                  wsol.mint,
				feeSenderATA:             wsolSenderATA,
				feeBillingATA:            wsol.billingATA,
				feeTokenBillingConfigPDA: wsol.fqBillingConfigPDA,
			},
			{
				name:                     "link",
				feeToken:                 link22.mint,
				userATA:                  link22.userATA,
				feeProgram:               link22.program,
				feeMint:                  link22.mint,
				feeSenderATA:             link22SenderATA,
				feeBillingATA:            link22.billingATA,
				feeTokenBillingConfigPDA: link22.fqBillingConfigPDA,
			},
		}

		chainConfig := []struct {
			destName                                               string
			chainSelector                                          uint64
			senderChainConfig, senderNonce, destChainPDA, fqConfig solana.PublicKey
		}{
			{
				destName:          "EVM",
				chainSelector:     config.EvmChainSelector,
				senderChainConfig: senderEvmDestChainConfigPDA,
				senderNonce:       senderEvmNoncePDA,
				destChainPDA:      config.EvmDestChainStatePDA,
				fqConfig:          config.FqEvmDestChainPDA,
			},
			{
				destName:          "SVM",
				chainSelector:     config.SvmChainSelector,
				senderChainConfig: senderSvmDestChainConfigPDA,
				senderNonce:       senderSvmNoncePDA,
				destChainPDA:      config.SvmDestChainStatePDA,
				fqConfig:          config.FqSvmDestChainPDA,
			},
		}

		for _, cc := range chainConfig {
			t.Run(fmt.Sprintf("SVM->%s", cc.destName), func(t *testing.T) {
				for _, fc := range feeConfig {
					t.Run(fmt.Sprintf("billing-%s/message_only", fc.name), func(t *testing.T) {
						ix, err := example_ccip_sender.NewCcipSendInstruction(
							cc.chainSelector,
							[]example_ccip_sender.SVMTokenAmount{}, // no tokens
							[]byte{1, 2, 3},                        // message data
							fc.feeToken,                            // empty fee token to indicate native SOL
							[]uint8{},
							senderState,
							cc.senderChainConfig,
							senderPDA,
							fc.userATA,
							user.PublicKey(),
							solana.SystemProgramID,
							config.CcipRouterProgram,
							config.RouterConfigPDA,
							cc.destChainPDA,
							cc.senderNonce,
							fc.feeProgram,
							fc.feeMint,
							fc.feeSenderATA,
							fc.feeBillingATA,
							config.BillingSignerPDA,
							config.FeeQuoterProgram,
							config.FqConfigPDA,
							cc.fqConfig,
							fc.feeTokenBillingConfigPDA,
							link22.fqBillingConfigPDA,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						).ValidateAndBuild()
						require.NoError(t, err)
						result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						require.NotNil(t, result)
						fmt.Printf("Logs: %s\n", result.Meta.LogMessages)
					})
				}

				// update config before running with tokens
				// SVM -> SVM with tokens requires different extraArgs than without tokens
				if cc.destName == "SVM" {
					svmDestChainWithTokensIx, err := example_ccip_sender.NewUpdateChainConfigInstruction(config.SvmChainSelector, validReceiverAddress[:], testutils.MustSerializeExtraArgs(t, fee_quoter.SVMExtraArgsV1{
						AllowOutOfOrderExecution: true,
						TokenReceiver:            validReceiverAddress,
					}, ccip.SVMExtraArgsV1Tag), senderState, senderSvmDestChainConfigPDA, user.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{svmDestChainWithTokensIx}, user, config.DefaultCommitment)
				}

				for _, fc := range feeConfig {
					t.Run(fmt.Sprintf("billing-%s/with_tokens", fc.name), func(t *testing.T) {
						t.Run("Happy path", func(t *testing.T) {
							base := example_ccip_sender.NewCcipSendInstruction(
								cc.chainSelector,
								[]example_ccip_sender.SVMTokenAmount{
									{
										Token:  token0.Mint,
										Amount: 1,
									},
									{
										Token:  token1.Mint,
										Amount: 2,
									},
								},
								[]byte{1, 2, 3}, // message data
								fc.feeToken,     // empty fee token to indicate native SOL
								[]uint8{2, 16},
								senderState,
								cc.senderChainConfig,
								senderPDA,
								fc.userATA,
								user.PublicKey(),
								solana.SystemProgramID,
								config.CcipRouterProgram,
								config.RouterConfigPDA,
								cc.destChainPDA,
								cc.senderNonce,
								fc.feeProgram,
								fc.feeMint,
								fc.feeSenderATA,
								fc.feeBillingATA,
								config.BillingSignerPDA,
								config.FeeQuoterProgram,
								config.FqConfigPDA,
								cc.fqConfig,
								fc.feeTokenBillingConfigPDA,
								link22.fqBillingConfigPDA,
								config.RMNRemoteProgram,
								config.RMNRemoteCursesPDA,
								config.RMNRemoteConfigPDA,
							)
							// pass user token accounts
							base.AccountMetaSlice = append(
								base.AccountMetaSlice,
								solana.Meta(token0.User[user.PublicKey()]).WRITE(),
								solana.Meta(token1.User[user.PublicKey()]).WRITE(),
							)

							// pass token pool accounts with the sender program ATA
							tokenMetas0, addressTables, err := tokens.ParseTokenLookupTableWithChain(ctx, solanaGoClient, token0, token0SenderATA, cc.chainSelector)
							require.NoError(t, err)
							base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas0...)
							tokenMetas1, addressTables1, err := tokens.ParseTokenLookupTableWithChain(ctx, solanaGoClient, token1, token1SenderATA, cc.chainSelector)
							require.NoError(t, err)
							base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas1...)
							addressTables[token1.PoolLookupTable] = addressTables1[token1.PoolLookupTable]
							for k, v := range ccipSendLookupTable {
								addressTables[k] = v
							}

							ix, err := base.ValidateAndBuild()
							require.NoError(t, err)

							ixApprove0, err := tokens.TokenApproveChecked(1, token0Decimals, token0.Program, token0.User[user.PublicKey()], token0.Mint, senderPDA, user.PublicKey(), nil)
							require.NoError(t, err)
							ixApprove1, err := tokens.TokenApproveChecked(2, token1Decimals, token1.Program, token1.User[user.PublicKey()], token1.Mint, senderPDA, user.PublicKey(), nil)
							require.NoError(t, err)

							testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove0, ixApprove1, ix}, user, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						})

						t.Run("When paying with "+fc.name+" and transferring link, it works", func(t *testing.T) {
							base := example_ccip_sender.NewCcipSendInstruction(
								cc.chainSelector,
								[]example_ccip_sender.SVMTokenAmount{
									{
										Token:  token0.Mint,
										Amount: 1,
									},
									{
										Token:  linkPool.Mint,
										Amount: 2,
									},
								},
								[]byte{1, 2, 3}, // message data
								fc.feeToken,     // empty fee token to indicate native SOL
								[]uint8{2, 16},
								senderState,
								cc.senderChainConfig,
								senderPDA,
								fc.userATA,
								user.PublicKey(),
								solana.SystemProgramID,
								config.CcipRouterProgram,
								config.RouterConfigPDA,
								cc.destChainPDA,
								cc.senderNonce,
								fc.feeProgram,
								fc.feeMint,
								fc.feeSenderATA,
								fc.feeBillingATA,
								config.BillingSignerPDA,
								config.FeeQuoterProgram,
								config.FqConfigPDA,
								cc.fqConfig,
								fc.feeTokenBillingConfigPDA,
								link22.fqBillingConfigPDA,
								config.RMNRemoteProgram,
								config.RMNRemoteCursesPDA,
								config.RMNRemoteConfigPDA,
							)
							// pass user token accounts
							base.AccountMetaSlice = append(
								base.AccountMetaSlice,
								solana.Meta(token0.User[user.PublicKey()]).WRITE(),
								solana.Meta(linkPool.User[user.PublicKey()]).WRITE(),
							)

							// pass token pool accounts with the sender program ATA
							tokenMetas0, addressTables, err := tokens.ParseTokenLookupTableWithChain(ctx, solanaGoClient, token0, token0SenderATA, cc.chainSelector)
							require.NoError(t, err)
							base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas0...)
							tokenMetasLink, addressTablesLink, err := tokens.ParseTokenLookupTableWithChain(ctx, solanaGoClient, linkPool, link22SenderATA, cc.chainSelector)
							require.NoError(t, err)
							base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetasLink...)
							addressTables[linkPool.PoolLookupTable] = addressTablesLink[linkPool.PoolLookupTable]
							for k, v := range ccipSendLookupTable {
								addressTables[k] = v
							}

							ix, err := base.ValidateAndBuild()
							require.NoError(t, err)

							ixApprove0, err := tokens.TokenApproveChecked(1, token0Decimals, token0.Program, token0.User[user.PublicKey()], token0.Mint, senderPDA, user.PublicKey(), nil)
							require.NoError(t, err)

							// Do NOT re-approve LINK transfer, as it was already approved in the setup step.
							// If we re-approve here just for the 2 tokens being sent, that will overwrite the delegated
							// amount, making future tests fail as the sender won't be allowed to transfer any more LINK

							testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove0, ix}, user, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						})
					})
				}
			})
		}

		t.Run("When transferring the same token twice, it fails", func(t *testing.T) {
			base := example_ccip_sender.NewCcipSendInstruction(
				config.EvmChainSelector,
				[]example_ccip_sender.SVMTokenAmount{
					{
						Token:  token0.Mint,
						Amount: 1,
					},
					{
						Token:  token0.Mint, // same token
						Amount: 2,
					},
				},
				[]byte{1, 2, 3},    // message data
				solana.PublicKey{}, // empty fee token to indicate native SOL
				[]uint8{2, 16},
				senderState,
				senderEvmDestChainConfigPDA,
				senderPDA,
				wsol.userATA,
				user.PublicKey(),
				solana.SystemProgramID,
				config.CcipRouterProgram,
				config.RouterConfigPDA,
				config.EvmDestChainStatePDA,
				senderEvmNoncePDA,
				solana.TokenProgramID,
				wsol.mint,
				wsolSenderATA,
				wsol.billingATA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,
				config.FqEvmDestChainPDA,
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.RMNRemoteProgram,
				config.RMNRemoteCursesPDA,
				config.RMNRemoteConfigPDA,
			)
			// pass user token accounts
			base.AccountMetaSlice = append(
				base.AccountMetaSlice,
				solana.Meta(token0.User[user.PublicKey()]).WRITE(),
				solana.Meta(token0.User[user.PublicKey()]).WRITE(),
			)

			// pass token pool accounts with the sender program ATA
			tokenMetas0, addressTables, err := tokens.ParseTokenLookupTableWithChain(ctx, solanaGoClient, token0, token0SenderATA, config.EvmChainSelector)
			require.NoError(t, err)
			base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas0...)
			base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas0...)
			for k, v := range ccipSendLookupTable {
				addressTables[k] = v
			}

			ix, err := base.ValidateAndBuild()
			require.NoError(t, err)

			ixApprove0, err := tokens.TokenApproveChecked(3, token0Decimals, token0.Program, token0.User[user.PublicKey()], token0.Mint, senderPDA, user.PublicKey(), nil)
			require.NoError(t, err)

			testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove0, ix}, user, config.DefaultCommitment, addressTables, []string{"TransferTokenDuplicated"})
		})
	})

	///////////////////////////
	// Withdraw billed funds //
	///////////////////////////
	t.Run("Withdraw billed funds", func(t *testing.T) {
		t.Run("Preconditions", func(t *testing.T) {
			require.Greater(t, getBalance(wsol.billingATA), uint64(0))
			require.Greater(t, getBalance(link22.billingATA), uint64(0))
		})

		t.Run("When an non-admin user tries to withdraw funds from a billing token account, it fails", func(t *testing.T) {
			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.billingATA,
				wsol.feeAggregatorATA,
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				user.PublicKey(), // wrong user here
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_CcipRouterError.String()})
		})

		t.Run("When withdrawing funds but sending them to the token account for a wrong token, it fails", func(t *testing.T) {
			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.billingATA,
				link22.feeAggregatorATA, // wrong token account
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip.InvalidInputsAtaAddress_CcipRouterError.String()})
		})

		t.Run("When withdrawing funds from a token account that does not belong to billing, it fails", func(t *testing.T) {
			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.userATA, // attempt to withdraw from user account instead of billingATA
				wsol.feeAggregatorATA,
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + common.ConstraintTokenOwner_AnchorError.String()})
		})

		t.Run("When withdrawing funds but sending them to a non-whitelisted token account, it fails", func(t *testing.T) {
			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.billingATA,
				wsol.userATA, // wrong destination, sending to user account instead of fee aggregator's
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip.InvalidInputsAtaAddress_CcipRouterError.String()})
		})

		t.Run("When trying to withdraw more funds than what's available, it fails", func(t *testing.T) {
			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				false,                         // withdraw all
				getBalance(wsol.billingATA)+1, // amount (more than what's available)
				wsol.mint,
				wsol.billingATA,
				wsol.feeAggregatorATA,
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip.InsufficientFunds_CcipRouterError.String()})
		})

		t.Run("When withdrawing a specific amount of funds, it succeeds", func(t *testing.T) {
			funds := getBalance(link22.billingATA)
			require.Greater(t, funds, uint64(0))

			initialAggrBalance := getBalance(link22.feeAggregatorATA)

			amount := uint64(2)

			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				false,  // withdraw all
				amount, // amount
				link22.mint,
				link22.billingATA,
				link22.feeAggregatorATA,
				link22.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)

			require.Equal(t, funds-amount, getBalance(link22.billingATA))                    // empty
			require.Equal(t, amount, getBalance(link22.feeAggregatorATA)-initialAggrBalance) // increased by exact amount
		})

		t.Run("When withdrawing all funds, it succeeds", func(t *testing.T) {
			funds := getBalance(wsol.billingATA)
			require.Greater(t, funds, uint64(0))

			initialAggrBalance := getBalance(wsol.feeAggregatorATA)

			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.billingATA,
				wsol.feeAggregatorATA,
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)

			require.Equal(t, uint64(0), getBalance(wsol.billingATA))                      // empty
			require.Equal(t, funds, getBalance(wsol.feeAggregatorATA)-initialAggrBalance) // increased by exact amount
		})

		t.Run("When withdrawing all funds but the accumulator account is already empty (no balance), it fails", func(t *testing.T) {
			funds := getBalance(wsol.billingATA)
			require.Equal(t, uint64(0), funds)

			ix, err := ccip_router.NewWithdrawBilledFundsInstruction(
				true,      // withdraw all
				uint64(0), // amount
				wsol.mint,
				wsol.billingATA,
				wsol.feeAggregatorATA,
				wsol.program,
				config.BillingSignerPDA,
				config.RouterConfigPDA,
				ccipAdmin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip.InsufficientFunds_CcipRouterError.String()})
		})
	})

	/////////////////////////////
	// FeeQuoter price updates //
	/////////////////////////////
	// These tests interact with FeeQuoter directly to try to update prices and the AllowedPriceUpdater list.
	// Later on, the offramp commit tests will further test the e2e commit including different cases of price updates.
	t.Run("FeeQuoter price updates", func(t *testing.T) {
		// re-using the legacy admin just in this suite because it has funds to submit transactions,
		// it is removed as a valid updater later in this same suite.
		testPriceUpdater := legacyAdmin
		testAllowedPriceUpdaterPDA, _, err := state.FindFqAllowedPriceUpdaterPDA(testPriceUpdater.PublicKey(), config.FeeQuoterProgram)
		require.NoError(t, err)

		t.Run("When a non-admin tries to add a price updater, it fails", func(t *testing.T) {
			ix, err := fee_quoter.NewAddPriceUpdaterInstruction(
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
				user.PublicKey(), // not admin
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_FeeQuoterError.String()})
		})

		t.Run("When an admin adds a price updater, it succeeds", func(t *testing.T) {
			ix, err := fee_quoter.NewAddPriceUpdaterInstruction(
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
		})

		t.Run("When the caller is not allowed to update the price, it fails", func(t *testing.T) {
			userAllowedPriceUpdaterPDA, _, err := state.FindFqAllowedPriceUpdaterPDA(user.PublicKey(), config.FeeQuoterProgram)
			require.NoError(t, err)

			ix, err := fee_quoter.NewUpdatePricesInstruction(
				[]fee_quoter.TokenPriceUpdate{},
				[]fee_quoter.GasPriceUpdate{},
				user.PublicKey(), // not allowed
				userAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.UnauthorizedPriceUpdater_FeeQuoterError.String()})
		})

		t.Run("When the caller is not allowed to update the price but tries to trick fee quoter to check an allowed caller, it fails", func(t *testing.T) {
			ix, err := fee_quoter.NewUpdatePricesInstruction(
				[]fee_quoter.TokenPriceUpdate{},
				[]fee_quoter.GasPriceUpdate{},
				user.PublicKey(),           // not allowed
				testAllowedPriceUpdaterPDA, // tricking the fee quoter to check the allowed price updater
				config.FqConfigPDA,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"AnchorError caused by account: allowed_price_updater. Error Code: " + common.ConstraintSeeds_AnchorError.String()})
		})

		t.Run("When the caller is allowed to update the price, it succeeds", func(t *testing.T) {
			// Check that initial onchain prices don't match the new ones
			tokenPrice := common.To28BytesBE(123)
			gasPrice := common.To28BytesBE(321)

			var tokenPriceAccount fee_quoter.BillingTokenConfigWrapper
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &tokenPriceAccount))
			require.NotEqual(t, tokenPrice, tokenPriceAccount.Config.UsdPerToken.Value)

			var gasPriceAccount fee_quoter.DestChain
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &gasPriceAccount))
			require.NotEqual(t, gasPrice, gasPriceAccount.State.UsdPerUnitGas.Value)

			// Update prices
			raw := fee_quoter.NewUpdatePricesInstruction(
				[]fee_quoter.TokenPriceUpdate{
					{
						SourceToken: wsol.mint,
						UsdPerToken: tokenPrice,
					},
				},
				[]fee_quoter.GasPriceUpdate{
					{
						DestChainSelector: config.EvmChainSelector,
						UsdPerUnitGas:     gasPrice,
					},
				},
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
			)
			raw.AccountMetaSlice.Append(solana.Meta(wsol.fqBillingConfigPDA).WRITE())
			raw.AccountMetaSlice.Append(solana.Meta(config.FqEvmDestChainPDA).WRITE())

			ix, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, testPriceUpdater, config.DefaultCommitment)

			// Check that final onchain prices match the new ones
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &tokenPriceAccount))
			require.Equal(t, tokenPrice, tokenPriceAccount.Config.UsdPerToken.Value)

			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &gasPriceAccount))
			require.Equal(t, gasPrice, gasPriceAccount.State.UsdPerUnitGas.Value)
		})

		t.Run("When the caller updates prices containing unregistered tokens, the call succeeds and the unregistered tokens are ignored", func(t *testing.T) {
			tokenPrice := common.To28BytesBE(123)
			anotherTokenPrice := common.To28BytesBE(234)

			var tokenPriceAccount fee_quoter.BillingTokenConfigWrapper
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &tokenPriceAccount))
			require.Equal(t, tokenPrice, tokenPriceAccount.Config.UsdPerToken.Value)

			// This token isn't registered as a billing token.
			require.True(t, common.IsClosedAccount(ctx, solanaGoClient, token2.Billing[config.EvmChainSelector], config.DefaultCommitment))

			// Update prices
			raw := fee_quoter.NewUpdatePricesInstruction(
				[]fee_quoter.TokenPriceUpdate{
					{
						SourceToken: wsol.mint,
						UsdPerToken: anotherTokenPrice,
					},
					{
						// This token will just be ignored.
						SourceToken: token2.Mint,
						UsdPerToken: anotherTokenPrice,
					},
				},
				[]fee_quoter.GasPriceUpdate{},
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
			)
			raw.AccountMetaSlice.Append(solana.Meta(wsol.fqBillingConfigPDA).WRITE())
			raw.AccountMetaSlice.Append(solana.Meta(token2.Billing[config.EvmChainSelector]).WRITE())

			ix, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, testPriceUpdater, config.DefaultCommitment)
			require.NotNil(t, result)

			updateIgnoredEvent := ccip.TokenPriceUpdateIgnored{}
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "TokenPriceUpdateIgnored", &updateIgnoredEvent, config.PrintEvents))
			require.Equal(t, updateIgnoredEvent.Token, token2.Mint)
			require.Equal(t, updateIgnoredEvent.Value, anotherTokenPrice)

			// Check that final onchain price matches the new one
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &tokenPriceAccount))
			require.Equal(t, anotherTokenPrice, tokenPriceAccount.Config.UsdPerToken.Value)

			// This token continues to be unregistered for billing.
			require.True(t, common.IsClosedAccount(ctx, solanaGoClient, token2.Billing[config.EvmChainSelector], config.DefaultCommitment))
		})

		t.Run("When a non-admin tries to remove a price updater, it fails", func(t *testing.T) {
			ix, err := fee_quoter.NewRemovePriceUpdaterInstruction(
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
				user.PublicKey(), // not admin
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip.Unauthorized_FeeQuoterError.String()})
		})

		t.Run("When an admin removes a price updater, it succeeds and the caller can no longer update prices", func(t *testing.T) {
			ix, err := fee_quoter.NewRemovePriceUpdaterInstruction(
				testPriceUpdater.PublicKey(),
				testAllowedPriceUpdaterPDA,
				config.FqConfigPDA,
				ccipAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
		})
	})

	//////////////////////////
	//        OffRamp       //
	//////////////////////////

	t.Run("OffRamp", func(t *testing.T) {
		//////////////////////////
		//     commit Tests     //
		//////////////////////////
		t.Run("Commit", func(t *testing.T) {
			currentMinSeqNr := uint64(1)

			oldReportContext := ccip.CreateReportContext(1) // use old sequence number

			type Comparator int
			const (
				Less Comparator = iota
				Equal
				Greater
			)

			t.Run("When committing a report with a valid source chain selector, merkle root and interval it succeeds", func(t *testing.T) {
				priceUpdatesCases := []struct {
					Name                    string
					PriceUpdates            ccip_offramp.PriceUpdates
					RemainingAccounts       []solana.PublicKey
					RunEventValidations     func(t *testing.T, tx *rpc.GetTransactionResult)
					RunStateValidations     func(t *testing.T)
					ReportContext           *[2][32]byte
					PriceSequenceComparator Comparator
					skipWithNoMerkleRoot    bool
				}{
					{
						Name:              "No price updates",
						PriceUpdates:      ccip_offramp.PriceUpdates{},
						RemainingAccounts: []solana.PublicKey{},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", nil, config.PrintEvents), "event not found")
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", nil, config.PrintEvents), "event not found")
						},
						RunStateValidations:     func(t *testing.T) {},
						PriceSequenceComparator: Greater, // it is a newer commit but with no price update
						skipWithNoMerkleRoot:    true,
					},
					{
						Name: "Single token price update",
						PriceUpdates: ccip_offramp.PriceUpdates{
							TokenPriceUpdates: []ccip_offramp.TokenPriceUpdate{{
								SourceToken: wsol.mint,
								UsdPerToken: common.To28BytesBE(1),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.OfframpStatePDA, wsol.fqBillingConfigPDA},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							// yes token update
							var update ccip.UsdPerTokenUpdated
							require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", &update, config.PrintEvents))
							require.Greater(t, update.Timestamp, int64(0)) // timestamp is set
							require.Equal(t, common.To28BytesBE(1), update.Value)

							// no gas updates
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", nil, config.PrintEvents), "event not found")
						},
						RunStateValidations: func(t *testing.T) {
							var tokenConfig fee_quoter.BillingTokenConfigWrapper
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &tokenConfig))
							require.Equal(t, common.To28BytesBE(1), tokenConfig.Config.UsdPerToken.Value)
							require.Greater(t, tokenConfig.Config.UsdPerToken.Timestamp, int64(0))
						},
						PriceSequenceComparator: Equal,
					},
					{
						Name: "Single gas price update on same chain as commit message",
						PriceUpdates: ccip_offramp.PriceUpdates{
							GasPriceUpdates: []ccip_offramp.GasPriceUpdate{{
								DestChainSelector: config.EvmChainSelector,
								UsdPerUnitGas:     common.To28BytesBE(1),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.OfframpStatePDA, config.FqEvmDestChainPDA},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							// no token updates
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", nil, config.PrintEvents), "event not found")

							// yes gas update
							var update ccip.UsdPerUnitGasUpdated
							require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", &update, config.PrintEvents))
							require.Greater(t, update.Timestamp, int64(0)) // timestamp is set
							require.Equal(t, common.To28BytesBE(1), update.Value)
						},
						RunStateValidations: func(t *testing.T) {
							var chainState fee_quoter.DestChain
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &chainState))
							require.Equal(t, common.To28BytesBE(1), chainState.State.UsdPerUnitGas.Value)
							require.Greater(t, chainState.State.UsdPerUnitGas.Timestamp, int64(0))
						},
						PriceSequenceComparator: Equal,
					},
					{
						Name: "Single gas price update on different chain (SVM) as commit message (EVM)",
						PriceUpdates: ccip_offramp.PriceUpdates{
							GasPriceUpdates: []ccip_offramp.GasPriceUpdate{{
								DestChainSelector: config.SvmChainSelector,
								UsdPerUnitGas:     common.To28BytesBE(2),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.OfframpStatePDA, config.FqSvmDestChainPDA},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							// no token updates
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", nil, config.PrintEvents), "event not found")

							// yes gas update
							var update ccip.UsdPerUnitGasUpdated
							require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", &update, config.PrintEvents))
							require.Greater(t, update.Timestamp, int64(0)) // timestamp is set
							require.Equal(t, common.To28BytesBE(2), update.Value)
						},
						RunStateValidations: func(t *testing.T) {
							var chainState fee_quoter.DestChain
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqSvmDestChainPDA, config.DefaultCommitment, &chainState))
							require.Equal(t, common.To28BytesBE(2), chainState.State.UsdPerUnitGas.Value)
							require.Greater(t, chainState.State.UsdPerUnitGas.Timestamp, int64(0))
						},
						PriceSequenceComparator: Equal,
					},
					{
						Name: "Multiple token & gas updates",
						PriceUpdates: ccip_offramp.PriceUpdates{
							TokenPriceUpdates: []ccip_offramp.TokenPriceUpdate{
								{SourceToken: wsol.mint, UsdPerToken: common.To28BytesBE(3)},
								{SourceToken: link22.mint, UsdPerToken: common.To28BytesBE(4)},
							},
							GasPriceUpdates: []ccip_offramp.GasPriceUpdate{
								{DestChainSelector: config.EvmChainSelector, UsdPerUnitGas: common.To28BytesBE(5)},
								{DestChainSelector: config.SvmChainSelector, UsdPerUnitGas: common.To28BytesBE(6)},
							},
						},
						RemainingAccounts: []solana.PublicKey{config.OfframpStatePDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA, config.FqEvmDestChainPDA, config.FqSvmDestChainPDA},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							// yes multiple token updates
							tokenUpdates, err := common.ParseMultipleEvents[ccip.UsdPerTokenUpdated](tx.Meta.LogMessages, "UsdPerTokenUpdated", config.PrintEvents)
							require.NoError(t, err)
							require.Len(t, tokenUpdates, 2)
							var eventWsol, eventLink22 bool
							for _, tokenUpdate := range tokenUpdates {
								switch tokenUpdate.Token {
								case wsol.mint:
									eventWsol = true
									require.Equal(t, common.To28BytesBE(3), tokenUpdate.Value)
								case link22.mint:
									eventLink22 = true
									require.Equal(t, common.To28BytesBE(4), tokenUpdate.Value)
								default:
									t.Fatalf("unexpected token update: %v", tokenUpdate)
								}
								require.Greater(t, tokenUpdate.Timestamp, int64(0)) // timestamp is set
							}
							require.True(t, eventWsol, "missing wsol update event")
							require.True(t, eventLink22, "missing link22 update event")

							// yes gas update
							gasUpdates, err := common.ParseMultipleEvents[ccip.UsdPerUnitGasUpdated](tx.Meta.LogMessages, "UsdPerUnitGasUpdated", config.PrintEvents)
							require.NoError(t, err)
							require.Len(t, gasUpdates, 2)
							var eventEvm, eventSVM bool
							for _, gasUpdate := range gasUpdates {
								switch gasUpdate.DestChain {
								case config.EvmChainSelector:
									eventEvm = true
									require.Equal(t, common.To28BytesBE(5), gasUpdate.Value)
								case config.SvmChainSelector:
									eventSVM = true
									require.Equal(t, common.To28BytesBE(6), gasUpdate.Value)
								default:
									t.Fatalf("unexpected gas update: %v", gasUpdate)
								}
								require.Greater(t, gasUpdate.Timestamp, int64(0)) // timestamp is set
							}
							require.True(t, eventEvm, "missing evm gas update event")
							require.True(t, eventSVM, "missing solana gas update event")
						},
						RunStateValidations: func(t *testing.T) {
							var wsolTokenConfig fee_quoter.BillingTokenConfigWrapper
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &wsolTokenConfig))
							require.Equal(t, common.To28BytesBE(3), wsolTokenConfig.Config.UsdPerToken.Value)
							require.Greater(t, wsolTokenConfig.Config.UsdPerToken.Timestamp, int64(0))

							var link22Config fee_quoter.BillingTokenConfigWrapper
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, link22.fqBillingConfigPDA, config.DefaultCommitment, &link22Config))
							require.Equal(t, common.To28BytesBE(4), link22Config.Config.UsdPerToken.Value)
							require.Greater(t, link22Config.Config.UsdPerToken.Timestamp, int64(0))

							var evmChainState fee_quoter.DestChain
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &evmChainState))
							require.Equal(t, common.To28BytesBE(5), evmChainState.State.UsdPerUnitGas.Value)
							require.Greater(t, evmChainState.State.UsdPerUnitGas.Timestamp, int64(0))

							var solanaChainState fee_quoter.DestChain
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqSvmDestChainPDA, config.DefaultCommitment, &solanaChainState))
							require.Equal(t, common.To28BytesBE(6), solanaChainState.State.UsdPerUnitGas.Value)
							require.Greater(t, solanaChainState.State.UsdPerUnitGas.Timestamp, int64(0))
						},
						PriceSequenceComparator: Equal,
					},
					{
						Name: "Valid price updates but old sequence number, so updates are ignored",
						PriceUpdates: ccip_offramp.PriceUpdates{
							TokenPriceUpdates: []ccip_offramp.TokenPriceUpdate{
								{SourceToken: wsol.mint, UsdPerToken: common.To28BytesBE(1)},
							},
							GasPriceUpdates: []ccip_offramp.GasPriceUpdate{
								{DestChainSelector: config.EvmChainSelector, UsdPerUnitGas: common.To28BytesBE(1)},
							},
						},
						RemainingAccounts: []solana.PublicKey{config.OfframpStatePDA, wsol.fqBillingConfigPDA, config.EvmDestChainStatePDA},
						ReportContext:     &oldReportContext,
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							// no events as updates are ignored (but commit is still accepted)
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", nil, config.PrintEvents), "event not found")
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", nil, config.PrintEvents), "event not found")
						},
						RunStateValidations: func(t *testing.T) {
							var wsolTokenConfig fee_quoter.BillingTokenConfigWrapper
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, wsol.fqBillingConfigPDA, config.DefaultCommitment, &wsolTokenConfig))
							// the price is NOT the one sent in this commit
							require.NotEqual(t, common.To28BytesBE(1), wsolTokenConfig.Config.UsdPerToken.Value)
							require.Greater(t, wsolTokenConfig.Config.UsdPerToken.Timestamp, int64(0))

							var evmChainState fee_quoter.DestChain
							require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &evmChainState))
							// the price is NOT the one sent in this commit
							require.NotEqual(t, common.To28BytesBE(1), evmChainState.State.UsdPerUnitGas.Value)
							require.Greater(t, evmChainState.State.UsdPerUnitGas.Timestamp, int64(0))
						},
						PriceSequenceComparator: Less, // it is an older commit, so price update is ignored and state remains ahead of this commit
					},
				}

				sequenceLength := uint64(5)

				for i, testcase := range priceUpdatesCases {
					t.Run(testcase.Name, func(t *testing.T) {
						msgAccounts := []solana.PublicKey{}
						_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{1, 2, 3, uint8(i)}, msgAccounts)
						rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
						require.NoError(t, err)

						minV := currentMinSeqNr
						maxV := currentMinSeqNr + sequenceLength - 1

						currentMinSeqNr = maxV + 1 // advance the outer sequence counter

						report := ccip_offramp.CommitInput{
							MerkleRoot: &ccip_offramp.MerkleRoot{
								SourceChainSelector: config.EvmChainSelector,
								OnRampAddress:       config.OnRampAddress,
								MinSeqNr:            minV,
								MaxSeqNr:            maxV,
								MerkleRoot:          root,
							},
							PriceUpdates: testcase.PriceUpdates,
						}

						var reportContext [2][32]byte
						var reportSequence uint64
						if testcase.ReportContext != nil {
							reportContext = *testcase.ReportContext
							reportSequence = ccip.ParseSequenceNumber(reportContext)
						} else {
							reportContext = ccip.NextCommitReportContext()
							reportSequence = ccip.ReportSequence()
						}

						sigs, err := ccip.SignCommitReport(reportContext, report, signers)
						require.NoError(t, err)

						transmitter := getTransmitter()

						raw := ccip_offramp.NewCommitInstruction(
							reportContext,
							testutils.MustMarshalBorsh(t, report),
							sigs.Rs,
							sigs.Ss,
							sigs.RawVs,
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							config.OfframpEvmSourceChainPDA,
							rootPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.OfframpBillingSignerPDA,
							config.FeeQuoterProgram,
							config.FqAllowedPriceUpdaterOfframpPDA,
							config.FqConfigPDA,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)

						for _, pubkey := range testcase.RemainingAccounts {
							raw.AccountMetaSlice.Append(solana.Meta(pubkey).WRITE())
						}

						instruction, err := raw.ValidateAndBuild()
						require.NoError(t, err)
						tx := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, offrampLookupTable, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))

						commitEvent := common.EventCommitReportAccepted{}
						require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent))
						require.Equal(t, config.EvmChainSelector, commitEvent.Report.SourceChainSelector)
						require.Equal(t, root, commitEvent.Report.MerkleRoot)
						require.Equal(t, minV, commitEvent.Report.MinSeqNr)
						require.Equal(t, maxV, commitEvent.Report.MaxSeqNr)

						transmittedEvent := ccip.EventTransmitted{}
						require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Transmitted", &transmittedEvent, config.PrintEvents))
						require.Equal(t, config.ConfigDigest, transmittedEvent.ConfigDigest)
						require.Equal(t, uint8(testutils.OcrCommitPlugin), transmittedEvent.OcrPluginType)
						require.Equal(t, reportSequence, transmittedEvent.SequenceNumber)

						var chainStateAccount ccip_offramp.SourceChain
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &chainStateAccount)
						require.NoError(t, err, "failed to get account info")
						require.Equal(t, currentMinSeqNr, chainStateAccount.State.MinSeqNr) // state now holds the "advanced outer" sequence number, which is the minimum for the next report
						// Do not check dest chain config, as it may have been updated by other tests in ccip onramp

						var rootAccount ccip_offramp.CommitReport
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
						require.NoError(t, err, "failed to get account info")
						require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)

						var globalState ccip_offramp.GlobalState
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpStatePDA, config.DefaultCommitment, &globalState)
						require.NoError(t, err)

						switch testcase.PriceSequenceComparator {
						case Less:
							require.Less(t, reportSequence, globalState.LatestPriceSequenceNumber)
						case Equal:
							require.Equal(t, reportSequence, globalState.LatestPriceSequenceNumber)
						case Greater:
							require.Greater(t, reportSequence, globalState.LatestPriceSequenceNumber)
						}

						testcase.RunEventValidations(t, tx)
						testcase.RunStateValidations(t)
					})
				}

				for _, testcase := range priceUpdatesCases {
					if testcase.skipWithNoMerkleRoot {
						continue
					}
					t.Run("prices only: "+testcase.Name, func(t *testing.T) {
						report := ccip_offramp.CommitInput{
							MerkleRoot:   nil,
							PriceUpdates: testcase.PriceUpdates,
						}

						var reportContext [2][32]byte
						if testcase.ReportContext != nil {
							reportContext = *testcase.ReportContext
						} else {
							reportContext = ccip.NextCommitReportContext()
						}

						sigs, err := ccip.SignCommitReport(reportContext, report, signers)
						require.NoError(t, err)

						transmitter := getTransmitter()

						raw := ccip_offramp.NewCommitPriceOnlyInstruction(
							reportContext,
							testutils.MustMarshalBorsh(t, report),
							sigs.Rs,
							sigs.Ss,
							sigs.RawVs,
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.OfframpBillingSignerPDA,
							config.FeeQuoterProgram,
							config.FqAllowedPriceUpdaterOfframpPDA,
							config.FqConfigPDA,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)

						for _, pubkey := range testcase.RemainingAccounts {
							raw.AccountMetaSlice.Append(solana.Meta(pubkey).WRITE())
						}

						instruction, err := raw.ValidateAndBuild()
						require.NoError(t, err)

						tx := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, offrampLookupTable, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						commitEvent := common.EventCommitReportAccepted{}
						require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent))

						require.Equal(t, testcase.PriceUpdates, commitEvent.PriceUpdates)
						require.Nil(t, commitEvent.Report)
						testcase.RunEventValidations(t, tx)
						testcase.RunStateValidations(t)
					})
				}
			})

			t.Run("Edge cases", func(t *testing.T) {
				t.Run("When committing a report with an invalid source chain selector it fails", func(t *testing.T) {
					t.Parallel()
					sourceChainSelector := uint64(34)
					sourceChainStatePDA, _, err := state.FindOfframpSourceChainPDA(sourceChainSelector, config.CcipOfframpProgram)
					require.NoError(t, err)
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, sourceChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindOfframpCommitReportPDA(sourceChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 4

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						sourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + common.AccountNotInitialized_AnchorError.String()})
				})

				t.Run("When committing a report with an invalid interval it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr - 2 // max lower than min

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.InvalidSequenceInterval_CcipOfframpError.String()})
				})

				t.Run("When committing a report with an interval size bigger than supported it fails", func(t *testing.T) {
					t.Parallel()
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 65 // max - min > 64

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.InvalidSequenceInterval_CcipOfframpError.String()})
				})

				t.Run("When committing a report with a zero merkle root it fails", func(t *testing.T) {
					t.Parallel()
					root := [32]byte{}
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr // max = min

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.InvalidProof_CcipOfframpError.String()})
				})

				t.Run("When committing a report with a repeated merkle root, it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{1, 2, 3, 1}, msgAccounts) // repeated root
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 4

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment,
						[]string{"Allocate: account Address", "already in use", "failed: custom program error: 0x0"})
				})

				t.Run("When committing a report with an invalid min interval, it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := uint64(8) // this is lower than expected
					maxV := uint64(10)

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.InvalidSequenceInterval_CcipOfframpError.String()})
				})

				t.Run("Invalid price updates", func(t *testing.T) {
					randomChain := uint64(123456) // just some non-existing chain
					randomChainPDA, _, err := state.FindFqDestChainPDA(randomChain, config.FeeQuoterProgram)
					require.NoError(t, err)

					testcases := []struct {
						Name              string
						Tokens            []solana.PublicKey
						GasChainSelectors []uint64
						AccountMetaSlice  solana.AccountMetaSlice
						ExpectedError     string
					}{
						{
							Name:              "with a gas price update for a chain that doesn't exist",
							GasChainSelectors: []uint64{randomChain},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(randomChainPDA).WRITE()},
							ExpectedError:     common.AccountNotInitialized_AnchorError.String(),
						},
						{
							Name:             "with a non-writable billing token config account",
							Tokens:           []solana.PublicKey{wsol.mint},
							AccountMetaSlice: solana.AccountMetaSlice{solana.Meta(wsol.fqBillingConfigPDA)}, // not writable
							ExpectedError:    ccip.InvalidInputsMissingWritable_CcipOfframpError.String(),
						},
						{
							// when the message source chain is the same as the chain whose gas is updated, then the same chain state is passed
							// in twice, in which case the resulting permissions are the sum of both instances. As only one is manually constructed here,
							// the other one is always writable (handled by the auto-generated code).
							Name:              "with a non-writable chain state account (different from the message source chain)",
							GasChainSelectors: []uint64{config.SvmChainSelector},                                 // the message source chain is EVM
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(config.SvmDestChainStatePDA)}, // not writable
							ExpectedError:     ccip.InvalidInputsDestChainStateAccount_FeeQuoterError.String(),
						},
						{
							Name:             "with the wrong billing token config account for a valid token",
							Tokens:           []solana.PublicKey{wsol.mint},
							AccountMetaSlice: solana.AccountMetaSlice{solana.Meta(link22.fqBillingConfigPDA).WRITE()}, // mismatch token
							ExpectedError:    ccip.InvalidInputsTokenConfigAccount_FeeQuoterError.String(),
						},
						{
							Name:              "with the wrong chain state account for a valid gas update",
							GasChainSelectors: []uint64{config.SvmChainSelector},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(config.EvmDestChainStatePDA).WRITE()}, // mismatch chain
							ExpectedError:     ccip.InvalidInputsDestChainStateAccount_FeeQuoterError.String(),
						},
						{
							Name:              "with too few accounts",
							Tokens:            []solana.PublicKey{wsol.mint},
							GasChainSelectors: []uint64{config.EvmChainSelector},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(wsol.fqBillingConfigPDA).WRITE()}, // missing chain state account
							ExpectedError:     ccip.InvalidInputsNumberOfAccounts_CcipOfframpError.String(),
						},
						// TODO right now I'm allowing sending too many remaining_accounts, but if we want to be restrictive with that we can add a test here
					}

					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{1, 2, 3}, msgAccounts)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					for _, testcase := range testcases {
						t.Run(testcase.Name, func(t *testing.T) {
							t.Parallel()

							priceUpdates := ccip_offramp.PriceUpdates{
								TokenPriceUpdates: make([]ccip_offramp.TokenPriceUpdate, len(testcase.Tokens)),
								GasPriceUpdates:   make([]ccip_offramp.GasPriceUpdate, len(testcase.GasChainSelectors)),
							}
							for i, token := range testcase.Tokens {
								priceUpdates.TokenPriceUpdates[i] = ccip_offramp.TokenPriceUpdate{
									SourceToken: token,
									UsdPerToken: common.To28BytesBE(uint64(i)),
								}
							}
							for i, chainSelector := range testcase.GasChainSelectors {
								priceUpdates.GasPriceUpdates[i] = ccip_offramp.GasPriceUpdate{
									DestChainSelector: chainSelector,
									UsdPerUnitGas:     common.To28BytesBE(uint64(i)),
								}
							}

							transmitter := getTransmitter()

							report := ccip_offramp.CommitInput{
								MerkleRoot: &ccip_offramp.MerkleRoot{
									SourceChainSelector: config.EvmChainSelector,
									OnRampAddress:       config.OnRampAddress,
									MinSeqNr:            currentMinSeqNr,
									MaxSeqNr:            currentMinSeqNr + 2,
									MerkleRoot:          root,
								},
								PriceUpdates: priceUpdates,
							}
							reportContext := ccip.NextCommitReportContext()
							sigs, err := ccip.SignCommitReport(reportContext, report, signers)
							require.NoError(t, err)

							raw := ccip_offramp.NewCommitInstruction(
								reportContext,
								testutils.MustMarshalBorsh(t, report),
								sigs.Rs,
								sigs.Ss,
								sigs.RawVs,
								config.OfframpConfigPDA,
								config.OfframpReferenceAddressesPDA,
								config.OfframpEvmSourceChainPDA,
								rootPDA,
								transmitter.PublicKey(),
								solana.SystemProgramID,
								solana.SysVarInstructionsPubkey,
								config.OfframpBillingSignerPDA,
								config.FeeQuoterProgram,
								config.FqAllowedPriceUpdaterOfframpPDA,
								config.FqConfigPDA,
								config.RMNRemoteProgram,
								config.RMNRemoteCursesPDA,
								config.RMNRemoteConfigPDA,
							)

							raw.AccountMetaSlice.Append(solana.Meta(config.OfframpStatePDA).WRITE())
							for _, meta := range testcase.AccountMetaSlice {
								raw.AccountMetaSlice.Append(meta)
							}

							instruction, err := raw.ValidateAndBuild()
							require.NoError(t, err)
							testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, offrampLookupTable, []string{testcase.ExpectedError}, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
						})
					}
				})
			})

			t.Run("When committing a report with the exact next interval, it succeeds", func(t *testing.T) {
				msgAccounts := []solana.PublicKey{}
				_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				minV := currentMinSeqNr
				maxV := currentMinSeqNr + 4

				currentMinSeqNr = maxV + 1 // advance the outer sequence counter as this will succeed

				report := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            minV,
						MaxSeqNr:            maxV,
						MerkleRoot:          root,
					},
				}
				reportContext := ccip.NextCommitReportContext()
				sigs, err := ccip.SignCommitReport(reportContext, report, signers)
				require.NoError(t, err)
				transmitter := getTransmitter()
				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, report),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				commitEvent := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent))
				require.Equal(t, config.EvmChainSelector, commitEvent.Report.SourceChainSelector)
				require.Equal(t, root, commitEvent.Report.MerkleRoot)
				require.Equal(t, minV, commitEvent.Report.MinSeqNr)
				require.Equal(t, maxV, commitEvent.Report.MaxSeqNr)

				transmittedEvent := ccip.EventTransmitted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Transmitted", &transmittedEvent, config.PrintEvents))
				require.Equal(t, config.ConfigDigest, transmittedEvent.ConfigDigest)
				require.Equal(t, uint8(testutils.OcrCommitPlugin), transmittedEvent.OcrPluginType)
				require.Equal(t, ccip.ReportSequence(), transmittedEvent.SequenceNumber)

				var chainStateAccount ccip_offramp.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpEvmSourceChainPDA, config.DefaultCommitment, &chainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, currentMinSeqNr, chainStateAccount.State.MinSeqNr)
				// Do not check dest chain config, as it may have been updated by other tests in ccip onramp

				var rootAccount ccip_offramp.CommitReport
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
				require.NoError(t, err, "failed to get account info")
				require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
			})

			t.Run("Ocr3Base::Transmit edge cases", func(t *testing.T) {
				t.Run("It rejects mismatch config digest", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					emptyReportContext := [2][32]byte{}

					instruction, err := ccip_offramp.NewCommitInstruction(
						emptyReportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ConfigDigestMismatch_CcipOfframpError.String()})
				})

				t.Run("It rejects unauthorized transmitter", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						user.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3UnauthorizedTransmitter_CcipOfframpError.String()})
				})

				t.Run("It rejects incorrect signature count", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					transmitter := getTransmitter()
					// remove signers
					sigs.Rs = sigs.Rs[1:]
					sigs.Ss = sigs.Ss[1:]

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3WrongNumberOfSignatures_CcipOfframpError.String()})
				})

				t.Run("It rejects invalid signature", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					sigs := ccip.Signatures{}
					transmitter := getTransmitter()

					instruction, err := ccip_offramp.NewCommitInstruction(
						ccip.NextCommitReportContext(),
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3WrongNumberOfSignatures_CcipOfframpError.String()})
				})

				t.Run("It rejects unauthorized signer", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)

					hash, err := ccip.HashCommitReport(reportContext, report)
					require.NoError(t, err)
					randomPrivateKey, err := secp256k1.GeneratePrivateKey()
					require.NoError(t, err)
					baseSig := ecdsa.SignCompact(randomPrivateKey, hash, false)
					sigs.RawVs[0] = baseSig[0] - 27 // key signs 27 or 28, but verification expects 0 or 1 (remove offset)
					sigs.Rs[0] = [32]byte(baseSig[1:33])
					sigs.Ss[0] = [32]byte(baseSig[33:65])

					transmitter := getTransmitter()

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3UnauthorizedSigner_CcipOfframpError.String()})
				})

				t.Run("It rejects duplicate signatures", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					reportContext := ccip.NextCommitReportContext()
					sigs, err := ccip.SignCommitReport(reportContext, report, signers)
					require.NoError(t, err)
					sigs.RawVs[0] = sigs.RawVs[1]
					sigs.Rs[0] = sigs.Rs[1]
					sigs.Ss[0] = sigs.Ss[1]
					transmitter := getTransmitter()

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3NonUniqueSignatures_CcipOfframpError.String()}, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT))
				})
			})
		})

		//////////////////////////
		//     execute Tests    //
		//////////////////////////

		t.Run("Execute", func(t *testing.T) {
			var executedSequenceNumber uint64
			reportContext := ccip.NextCommitReportContext() // reuse the same commit for all executions

			t.Run("When executing a report with merkle tree of size 1, it succeeds", func(t *testing.T) {
				transmitter := getTransmitter()

				sourceChainSelector := config.EvmChainSelector
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber
				executedSequenceNumber = sequenceNumber // persist this number as executed, for later tests

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)

				executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
				require.NoError(t, err)

				require.Equal(t, 2, len(executionEvents))
				require.Equal(t, config.EvmChainSelector, executionEvents[0].SourceChainSelector)
				require.Equal(t, sequenceNumber, executionEvents[0].SequenceNumber)
				require.Equal(t, hex.EncodeToString(message.Header.MessageId[:]), hex.EncodeToString(executionEvents[0].MessageID[:]))
				require.Equal(t, hex.EncodeToString(root[:]), hex.EncodeToString(executionEvents[0].MessageHash[:]))
				require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)

				require.Equal(t, config.EvmChainSelector, executionEvents[1].SourceChainSelector)
				require.Equal(t, sequenceNumber, executionEvents[1].SequenceNumber)
				require.Equal(t, hex.EncodeToString(message.Header.MessageId[:]), hex.EncodeToString(executionEvents[1].MessageID[:]))
				require.Equal(t, hex.EncodeToString(root[:]), hex.EncodeToString(executionEvents[1].MessageHash[:]))
				require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

				var rootAccount ccip_offramp.CommitReport
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
				require.NoError(t, err, "failed to get account info")
				require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
				require.Equal(t, bin.Uint128{Lo: 2, Hi: 0}, rootAccount.ExecutionStates)
				require.Equal(t, sequenceNumber, rootAccount.MinMsgNr)
				require.Equal(t, sequenceNumber, rootAccount.MaxMsgNr)
			})

			t.Run("Merkle root PDA can be closed iff it is successfully executed", func(t *testing.T) {
				transmitter := getTransmitter()

				sourceChainSelector := config.EvmChainSelector
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				// Validate that the CommitReportPDA cannot be closed before the execution is successful
				closeCommitPDAInstruction, err := ccip_offramp.NewCloseCommitReportAccountInstruction(
					commitReport.MerkleRoot.SourceChainSelector,
					commitReport.MerkleRoot.MerkleRoot[:],
					config.OfframpConfigPDA,
					rootPDA,
					config.OfframpReferenceAddressesPDA,
					wsol.mint,
					wsol.billingATA,
					config.BillingSignerPDA,
					wsol.program,
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{closeCommitPDAInstruction}, user, config.DefaultCommitment, []string{ccip.CommitReportHasPendingMessages_CcipOfframpError.String()})

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)

				// Validate that the execution is successful
				executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
				require.NoError(t, err)
				require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

				// Validate that the CommitReportPDA has rent
				balanceResult, err := solanaGoClient.GetBalance(ctx, rootPDA, config.DefaultCommitment)
				require.NoError(t, err)
				initCommitPDABalance := balanceResult.Value
				require.Greater(t, initCommitPDABalance, uint64(0))
				// initWsolBalance := getBalance(wsol.billingATA)
				initReceiverBalance, err := solanaGoClient.GetBalance(ctx, wsol.billingATA, config.DefaultCommitment)
				require.NoError(t, err)

				closeCommitPDAInstruction, err = ccip_offramp.NewCloseCommitReportAccountInstruction(
					commitReport.MerkleRoot.SourceChainSelector,
					commitReport.MerkleRoot.MerkleRoot[:],
					config.OfframpConfigPDA,
					rootPDA,
					config.OfframpReferenceAddressesPDA,
					wsol.mint,
					wsol.billingATA,
					config.BillingSignerPDA,
					wsol.program,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{closeCommitPDAInstruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting

				commitPDACloseEvent := ccip.EventCommitReportPDAClosed{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportPDAClosed", &commitPDACloseEvent, config.PrintEvents))

				require.Equal(t, commitReport.MerkleRoot.SourceChainSelector, commitPDACloseEvent.SourceChainSelector)
				require.Equal(t, hex.EncodeToString(root[:]), hex.EncodeToString(commitPDACloseEvent.MerkleRoot[:]))

				// Validate that the CommitReportPDA is closed
				testutils.AssertClosedAccount(ctx, t, solanaGoClient, rootPDA, config.DefaultCommitment)

				// Validate that the CommitReportPDA balance is 0
				balanceResult, err = solanaGoClient.GetBalance(ctx, rootPDA, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, uint64(0), balanceResult.Value)

				// Validate that the CommitReportPDA rent has been transferred to the billing account
				finalReceiverBalance, err := solanaGoClient.GetBalance(ctx, wsol.billingATA, config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, initReceiverBalance.Value+initCommitPDABalance, finalReceiverBalance.Value)
			})

			t.Run("When executing a report with not matching source chain selector in message, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				message.Header.SourceChainSelector = 89

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Message: Source chain selector not supported."})
			})

			t.Run("When executing a report with a wrong external execution signer PDA, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				invalidExecSignerPDA, _, err := state.FindExternalExecutionConfigPDA(config.CcipInvalidReceiverProgram, config.CcipOfframpProgram)
				require.NoError(t, err)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(invalidExecSignerPDA, false, false), // invalid signer here (wrong receiver program)
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip.InvalidInputsExternalExecutionSignerAccount_CcipOfframpError.String()})
			})

			t.Run("execute fails with a global curse active", func(t *testing.T) {
				transmitter := getTransmitter()

				sourceChainSelector := config.EvmChainSelector
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				globalCurse := rmn_remote.CurseSubject{
					Value: [16]uint8{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
				}

				// Curse is applied
				ix, err := rmn_remote.NewCurseInstruction(
					globalCurse,
					config.RMNRemoteConfigPDA,
					ccipAdmin.PublicKey(),
					config.RMNRemoteCursesPDA,
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
				require.NotNil(t, result)

				// Manually execution fails in the same way
				manual := ccip_offramp.NewManuallyExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					user.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				manual.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: GloballyCursed"})
				require.NotNil(t, result)

				// Curse is removed
				ix, err = rmn_remote.NewUncurseInstruction(
					globalCurse,
					config.RMNRemoteConfigPDA,
					ccipAdmin.PublicKey(),
					config.RMNRemoteCursesPDA,
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)
				require.NotNil(t, result)
			})

			t.Run("When executing a report with unsupported source chain selector account, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				unsupportedChainSelector := uint64(34)
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				unsupportedSourceChainPDA, _, err := state.FindOfframpSourceChainPDA(unsupportedChainSelector, config.CcipOfframpProgram)
				require.NoError(t, err)
				require.NoError(t, err)
				message.Header.SourceChainSelector = unsupportedChainSelector
				message.Header.SequenceNumber = 1

				instruction, err = ccip_offramp.NewAddSourceChainInstruction(
					unsupportedChainSelector,
					validSourceChainConfig,
					unsupportedSourceChainPDA,
					config.OfframpConfigPDA,
					ccipAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: unsupportedChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					unsupportedSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"AnchorError caused by account: commit_report. Error Code: " + common.ConstraintSeeds_AnchorError.String()})
			})

			// TODO review test case, code does not match the test name
			t.Run("When executing a report with incorrect solana chain selector, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.Header.DestChainSelector = 89 // invalid dest chain selector
				sequenceNumber := message.Header.SequenceNumber
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)

				require.NoError(t, err)
				root := [32]byte(hash)

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.UnsupportedDestinationChainSelector_CcipOfframpError.String()})
			})

			t.Run("When executing a report with nonexisting PDA for the Merkle Root, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Message: The program expected this account to be already initialized."})
			})

			t.Run("When executing a report for an already executed message, it is skipped", func(t *testing.T) {
				transmitter := getTransmitter()

				sourceChainSelector := config.EvmChainSelector

				message := ccip.CreateDefaultMessageWith(sourceChainSelector, executedSequenceNumber) // already executed seq number
				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)
				root := [32]byte(hash)

				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvent := ccip.EventSkippedAlreadyExecutedMessage{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "SkippedAlreadyExecutedMessage", &executionEvent, config.PrintEvents))

				require.NoError(t, err)
				require.NotNil(t, executionEvent)
				require.Equal(t, config.EvmChainSelector, executionEvent.SourceChainSelector)
				require.Equal(t, executedSequenceNumber, executionEvent.SequenceNumber)
			})

			t.Run("When executing a report for an already executed root, but not message, it succeeds", func(t *testing.T) {
				transmitter := getTransmitter()

				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message1, hash1 := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message2 := ccip.CreateDefaultMessageWith(config.EvmChainSelector, message1.Header.SequenceNumber+1)
				hash2, err := ccip.HashAnyToSVMMessage(message2, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				root, err := ccip.MerkleFrom([][32]byte{hash1, [32]byte(hash2)})
				require.NoError(t, err)

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message1.Header.SequenceNumber,
						MaxSeqNr:            message2.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport1 := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message2, // execute out of order
					Proofs:              [][32]uint8{hash1},
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport1),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvent := ccip.EventExecutionStateChanged{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

				executionReport2 := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message1,
					Proofs:              [][32]uint8{[32]byte(hash2)},
				}
				raw = ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport2),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
				require.NoError(t, err)

				require.Equal(t, 2, len(executionEvents))
				require.Equal(t, config.EvmChainSelector, executionEvents[0].SourceChainSelector)
				require.Equal(t, message1.Header.SequenceNumber, executionEvents[0].SequenceNumber)
				require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvents[0].MessageID[:]))
				require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvents[0].MessageHash[:]))
				require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)

				require.Equal(t, config.EvmChainSelector, executionEvents[1].SourceChainSelector)
				require.Equal(t, message1.Header.SequenceNumber, executionEvents[1].SequenceNumber)
				require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvents[1].MessageID[:]))
				require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvents[1].MessageHash[:]))
				require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

				var rootAccount ccip_offramp.CommitReport
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
				require.NoError(t, err, "failed to get account info")
				require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
				require.Equal(t, bin.Uint128{Lo: 10, Hi: 0}, rootAccount.ExecutionStates)
				require.Equal(t, message1.Header.SequenceNumber, rootAccount.MinMsgNr)
				require.Equal(t, message2.Header.SequenceNumber, rootAccount.MaxMsgNr)
			})

			t.Run("When executing a report that receiver program needs to init an account, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				stubAccountPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("counter")}, config.CcipInvalidReceiverProgram)
				msgAccounts := []solana.PublicKey{config.CcipInvalidReceiverProgram, stubAccountPDA, solana.SystemProgramID}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.TokenReceiver = stubAccountPDA
				sequenceNumber := message.Header.SequenceNumber
				message.ExtraArgs.IsWritableBitmap = 0
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)
				root := [32]byte(hash)

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				offrampExecSignerPDA, _, _ := state.FindExternalExecutionConfigPDA(config.CcipInvalidReceiverProgram, config.CcipOfframpProgram)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipInvalidReceiverProgram, false, false),
					solana.NewAccountMeta(offrampExecSignerPDA, true, false),
					solana.NewAccountMeta(stubAccountPDA, false, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)

				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				// failed ccipReceiver - init account requires mutable authority
				// ccipSigner is not a mutable account
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"writable privilege escalated", "Cross-program invocation with unauthorized signer or writable account"})
			})

			t.Run("When executing a report with an account with read permissions but sent as write, it uses the correct permissions", func(t *testing.T) {
				// This test is to check that the receiver program can execute a message with an account that is configured as readable but sent as writable
				// This is for a particular case when an account used in the message as read is also sent in the transaction (for example for the token pool program)
				// as writable. So, the CCIP Receiver program should be able to use the account as readable and NOT as writable.
				// All the accounts sent in the CPI to the CCIP Receiver must used the writable bitmap declared on source, but the check for that bitmap should be
				// as flexible to support using the same account with different permissions in the same message.
				transmitter := getTransmitter()

				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.ExtraArgs.IsWritableBitmap = ccip.GenerateBitMapForIndexes([]int{1})

				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				root := [32]byte(hash)

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message.Header.SequenceNumber,
						MaxSeqNr:            message.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false), // index 0 --> configured as readable
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),           // index 1
					solana.NewAccountMeta(solana.SystemProgramID, false, false),                   // index 2
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				// Fails in the CCIP Receiver contract as it expects the config.ReceiverExternalExecutionConfigPDA account to be writable, but it was sent as read only (the same way it was configured on source).
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Program log: Instruction: CcipReceive Program log: AnchorError caused by account: external_execution_config. Error Code: ConstraintMut. Error Number: 2000. Error Message: A mut constraint was violated."})
			})

			t.Run("message can be executed with empty Any2SVMRampMessage.Data", func(t *testing.T) {
				transmitter := getTransmitter()

				sourceChainSelector := config.EvmChainSelector
				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.Data = []byte{} // empty message data
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)
				root := [32]byte(hash)

				sequenceNumber := message.Header.SequenceNumber
				executedSequenceNumber = sequenceNumber // persist this number as executed, for later tests

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(computebudget.MAX_COMPUTE_UNIT_LIMIT)) // signature verification compute unit amounts can vary depending on sorting
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)

				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
			})

			t.Run("token happy path", func(t *testing.T) {
				t.Run("single token", func(t *testing.T) {
					_, initSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint, config.DefaultCommitment)
					require.NoError(t, err)
					_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)

					transmitter := getTransmitter()

					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint,
						// 1 token * 1e9 due to decimal differences (18 in EVM, 9 in SVM). This will result in 1 unit in SVM.
						Amount: ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1000000000)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}},
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{5},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpEvmPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					raw.AccountMetaSlice = append(
						raw.AccountMetaSlice,
						solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
						solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
						solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
						solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
						solana.NewAccountMeta(solana.SystemProgramID, false, false),
					)

					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					tx = testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(400_000))

					executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
					require.NoError(t, err)

					require.Equal(t, 2, len(executionEvents))
					require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)
					require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

					mintEvent := tokens.EventMintRelease{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Minted", &mintEvent, config.PrintEvents))
					require.Equal(t, config.ReceiverExternalExecutionConfigPDA, mintEvent.Recipient)
					require.Equal(t, token0.PoolSigner, mintEvent.Sender)
					require.Equal(t, uint64(1), mintEvent.Amount)
					require.Equal(t, token0.Mint, mintEvent.Mint)

					_, finalSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint, config.DefaultCommitment)
					require.NoError(t, err)
					_, finalBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, 1, finalSupply-initSupply)
					require.Equal(t, 1, finalBal-initBal)
				})

				t.Run("two tokens", func(t *testing.T) {
					_, initBal0, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					_, initBal1, err := tokens.TokenBalance(ctx, solanaGoClient, token1.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)

					transmitter := getTransmitter()

					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.ExtraArgs = ccip_offramp.Any2SVMRampExtraArgs{}
					message.Data = []byte{}
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint,
						// 1 token * 1e9 due to decimal differences (18 in EVM, 9 in SVM). This will result in 1 unit in SVM.
						Amount: ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1000000000)},
					}, {
						SourcePoolAddress: []byte{4, 5, 6},
						DestTokenAddress:  token1.Mint,
						// Token 2 has 18 decimals on both sides, no conversion will occur.
						Amount: ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(2)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}, {}},
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{0, 15},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpEvmPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)

					tokenMetas0, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas0...)

					tokenMetas1, addressTables1, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token1, token1.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token1.OfframpSigner))
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas1...)

					maps.Copy(addressTables, addressTables1)
					maps.Copy(addressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(400_000))

					// validate amounts
					_, finalBal0, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, 1, finalBal0-initBal0)
					_, finalBal1, err := tokens.TokenBalance(ctx, solanaGoClient, token1.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, 2, finalBal1-initBal1)
				})
			})

			t.Run("execute measure compute units", func(t *testing.T) {
				transmitter := getTransmitter()

				t.Run("Example Message Execution", func(t *testing.T) {
					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpEvmPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					raw.AccountMetaSlice = append(
						raw.AccountMetaSlice,
						solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
						solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
						solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
						solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
						solana.NewAccountMeta(solana.SystemProgramID, false, false),
					)

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					cu := testutils.GetRequiredCU(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)

					// This validation is like a snapshot for gas consumption
					require.LessOrEqual(t, cu, uint32(110_000))

					tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(cu))

					executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
					require.NoError(t, err)

					require.Equal(t, 2, len(executionEvents))
					require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)
					require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)
				})

				t.Run("1 Test Token Transfer", func(t *testing.T) {
					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.ExtraArgs = ccip_offramp.Any2SVMRampExtraArgs{}
					message.Data = []byte{}
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint,
						Amount:            ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}},
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{0},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpEvmPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)

					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)

					maps.Copy(addressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					cu := testutils.GetRequiredCUWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables)

					// This validation is like a snapshot for gas consumption
					// Execute: 1 Token Transfer + Message Execution
					require.LessOrEqual(t, cu, uint32(180_000))

					tx = testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(cu))

					executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
					require.NoError(t, err)

					require.Equal(t, 2, len(executionEvents))
					require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)
					require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

					mintEvent := tokens.EventMintRelease{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Minted", &mintEvent, config.PrintEvents))
					require.Equal(t, config.ReceiverExternalExecutionConfigPDA, mintEvent.Recipient)
					require.Equal(t, token0.Mint, mintEvent.Mint)
				})

				t.Run("1 Test Token Transfer + Example Message Execution", func(t *testing.T) {
					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint,
						Amount:            ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}},
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{5},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpEvmSourceChainPDA,
						rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpEvmPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					raw.AccountMetaSlice = append(
						raw.AccountMetaSlice,
						solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
						solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
						solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
						solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
						solana.NewAccountMeta(solana.SystemProgramID, false, false),
					)

					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)

					maps.Copy(addressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					cu := testutils.GetRequiredCUWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables)

					// This validation is like a snapshot for gas consumption
					// Execute: 1 Token Transfer + Message Execution
					require.LessOrEqual(t, cu, uint32(250_000))

					tx = testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(cu))

					executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
					require.NoError(t, err)

					require.Equal(t, 2, len(executionEvents))
					require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)
					require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

					mintEvent := tokens.EventMintRelease{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Minted", &mintEvent, config.PrintEvents))
					require.Equal(t, config.ReceiverExternalExecutionConfigPDA, mintEvent.Recipient)
					require.Equal(t, token0.Mint, mintEvent.Mint)
				})
			})

			t.Run("tokens other cases", func(t *testing.T) {
				type setupArgs struct {
					sourceChainSelector   uint64
					offrampSourceChainPDA solana.PublicKey
					onramp                []byte
					sequenceNumber        uint64 // use 0 if not overriding the default (which only tracks EVM seq num)
				}
				type setupResult struct {
					initSupply  int
					initBal     int
					message     ccip_offramp.Any2SVMRampMessage
					root        [32]byte
					rootPDA     solana.PublicKey
					transmitter solana.PrivateKey
				}
				testSetup := func(t *testing.T, args setupArgs) setupResult {
					_, initSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint, config.DefaultCommitment)
					require.NoError(t, err)
					_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)

					transmitter := getTransmitter()

					msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.Header.SourceChainSelector = args.sourceChainSelector
					if args.sequenceNumber != 0 {
						message.Header.SequenceNumber = args.sequenceNumber
					}
					require.NoError(t, err)
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint,
						Amount:            ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, args.onramp, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_offramp.CommitInput{
						MerkleRoot: &ccip_offramp.MerkleRoot{
							SourceChainSelector: args.sourceChainSelector,
							OnRampAddress:       args.onramp,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindOfframpCommitReportPDA(args.sourceChainSelector, root, config.CcipOfframpProgram)
					require.NoError(t, err)

					instruction, err := ccip_offramp.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						args.offrampSourceChainPDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.OfframpBillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqAllowedPriceUpdaterOfframpPDA,
						config.FqConfigPDA,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, offrampLookupTable, common.AddComputeUnitLimit(300_000))
					event := common.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

					return setupResult{
						initSupply,
						initBal,
						message,
						root,
						rootPDA,
						transmitter,
					}
				}

				t.Run("Router authorization of offramp for lanes", func(t *testing.T) {
					// Check SVM is an accepted source chain
					var sourceChain ccip_offramp.SourceChain
					require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.OfframpSvmSourceChainPDA, config.DefaultCommitment, &sourceChain))
					t.Log(sourceChain.Config)
					require.Equal(t, sourceChain.Config.OnRamp.Bytes[0:32], config.CcipRouterProgram[:])

					// Check offramp program is not registered in router as valid offramp for SVM<>SVM lane
					testutils.AssertClosedAccount(ctx, t, solanaGoClient, config.AllowedOfframpSvmPDA, config.DefaultCommitment)

					setup := testSetup(t, setupArgs{
						sourceChainSelector:   config.SvmChainSelector,
						offrampSourceChainPDA: config.OfframpSvmSourceChainPDA,
						onramp:                config.CcipRouterProgram.Bytes(),
						sequenceNumber:        1, // most utils assume EVM, but this is the first message from SVM
					})

					executionReport := ccip_offramp.ExecutionReportSingleChain{
						SourceChainSelector: config.SvmChainSelector,
						Message:             setup.message,
						OffchainTokenData:   [][]byte{{}},
						Proofs:              [][32]uint8{},
					}
					raw := ccip_offramp.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{5},
						config.OfframpConfigPDA,
						config.OfframpReferenceAddressesPDA,
						config.OfframpSvmSourceChainPDA,
						setup.rootPDA,
						config.CcipOfframpProgram,
						config.AllowedOfframpSvmPDA,
						setup.transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.RMNRemoteProgram,
						config.RMNRemoteCursesPDA,
						config.RMNRemoteConfigPDA,
					)
					raw.AccountMetaSlice = append(
						raw.AccountMetaSlice,
						solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
						solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
						solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
						solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
						solana.NewAccountMeta(solana.SystemProgramID, false, false),
					)

					raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					maps.Copy(addressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

					require.NoError(t, err)
					tokenMetas[1] = solana.Meta(token0.Billing[config.SvmChainSelector])       // overwrite field that our TokenPool utils assume will be for EVM
					tokenMetas[2] = solana.Meta(token0.Chain[config.SvmChainSelector]).WRITE() // overwrite field that our TokenPool utils assume will be for EVM
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)
					executeIx, err := raw.ValidateAndBuild()
					require.NoError(t, err)

					// It fails here as the offramp isn't registered in the router for that lane
					testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{executeIx}, setup.transmitter, config.DefaultCommitment, addressTables, []string{ccip.InvalidInputsAllowedOfframpAccount_CcipOfframpError.String()})

					// Allow that offramp for SVM
					allowIx, err := ccip_router.NewAddOfframpInstruction(
						config.SvmChainSelector,
						config.CcipOfframpProgram,
						config.AllowedOfframpSvmPDA,
						config.RouterConfigPDA,
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{allowIx}, ccipAdmin, config.DefaultCommitment)

					// Now the execute should succeed
					testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{executeIx}, setup.transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(250_000))
				})
			})

			t.Run("OffRamp Manual Execution: when executing a non-committed report, it fails", func(t *testing.T) {
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}

				raw := ccip_offramp.NewManuallyExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					[]byte{},
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					user.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"The program expected this account to be already initialized"})
			})

			t.Run("OffRamp Manual execution", func(t *testing.T) {
				transmitter := getTransmitter()

				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message1, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)

				hash1, err := ccip.HashAnyToSVMMessage(message1, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				message2 := ccip.CreateDefaultMessageWith(config.EvmChainSelector, message1.Header.SequenceNumber+1)
				hash2, err := ccip.HashAnyToSVMMessage(message2, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				root, err := ccip.MerkleFrom([][32]byte{[32]byte(hash1), [32]byte(hash2)})
				require.NoError(t, err)

				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message1.Header.SequenceNumber,
						MaxSeqNr:            message2.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)

				instruction, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				t.Run("Before elapsed time", func(t *testing.T) {
					t.Run("When user manually executing before the period of time has passed, it fails", func(t *testing.T) {
						executionReport := ccip_offramp.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_offramp.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							config.OfframpEvmSourceChainPDA,
							rootPDA,
							config.CcipOfframpProgram,
							config.AllowedOfframpEvmPDA,
							user.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						fmt.Printf("User: %s\n", user.PublicKey().String())
						fmt.Printf("Transmitter: %s\n", transmitter.PublicKey().String())

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip.ManualExecutionNotAllowed_CcipOfframpError.String()})
					})

					t.Run("When transmitter manually executing before the period of time has passed, it fails", func(t *testing.T) {
						executionReport := ccip_offramp.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_offramp.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							config.OfframpEvmSourceChainPDA,
							rootPDA,
							config.CcipOfframpProgram,
							config.AllowedOfframpEvmPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip.ManualExecutionNotAllowed_CcipOfframpError.String()})
					})
				})

				t.Run("Given the period of time has passed", func(t *testing.T) {
					newEnableManualExecutionAfter := int64(-1)

					instruction, err = ccip_offramp.NewUpdateEnableManualExecutionAfterInstruction(
						newEnableManualExecutionAfter,
						config.OfframpConfigPDA,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
					require.NotNil(t, result)

					var configSetEvent ccip.EventOfframpConfigSet
					require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "ConfigSet", &configSetEvent, config.PrintEvents))
					require.Equal(t, newEnableManualExecutionAfter, configSetEvent.EnableManualExecutionAfter)

					t.Run("When user manually executing after the period of time has passed, it succeeds", func(t *testing.T) {
						executionReport := ccip_offramp.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_offramp.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							config.OfframpEvmSourceChainPDA,
							rootPDA,
							config.CcipOfframpProgram,
							config.AllowedOfframpEvmPDA,
							user.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						t.Run("When the receiver is rejecting calls, it fails", func(t *testing.T) {
							// Make receiver reject calls
							rejectIx, err2 := test_ccip_receiver.NewSetRejectAllInstruction(
								true,
								config.ReceiverTargetAccountPDA,
								user.PublicKey(),
							).ValidateAndBuild()
							require.NoError(t, err2)
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{rejectIx}, user, config.DefaultCommitment)

							// Check that it fails
							testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"RejectAll"})

							acceptIx, err2 := test_ccip_receiver.NewSetRejectAllInstruction(
								false,
								config.ReceiverTargetAccountPDA,
								user.PublicKey(),
							).ValidateAndBuild()
							require.NoError(t, err2)

							// Make receiver accept calls again, to avoid disrupting following tests
							testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{acceptIx}, user, config.DefaultCommitment)
						})

						t.Run("When the receiver is accepting calls, it succeeds", func(t *testing.T) {
							tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)

							executionEvents, err2 := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
							require.NoError(t, err2)

							require.Equal(t, 2, len(executionEvents))
							require.Equal(t, config.EvmChainSelector, executionEvents[0].SourceChainSelector)
							require.Equal(t, message1.Header.SequenceNumber, executionEvents[0].SequenceNumber)
							require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvents[0].MessageID[:]))
							require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvents[0].MessageHash[:]))
							require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)

							require.Equal(t, config.EvmChainSelector, executionEvents[1].SourceChainSelector)
							require.Equal(t, message1.Header.SequenceNumber, executionEvents[1].SequenceNumber)
							require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvents[1].MessageID[:]))
							require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvents[1].MessageHash[:]))
							require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

							var rootAccount ccip_offramp.CommitReport
							err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
							require.NoError(t, err, "failed to get account info")
							require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
							require.Equal(t, bin.Uint128{Lo: 2, Hi: 0}, rootAccount.ExecutionStates)
							require.Equal(t, commitReport.MerkleRoot.MinSeqNr, rootAccount.MinMsgNr)
							require.Equal(t, commitReport.MerkleRoot.MaxSeqNr, rootAccount.MaxMsgNr)
						})
					})

					t.Run("When transmitter executing after the period of time has passed, it succeeds", func(t *testing.T) {
						executionReport := ccip_offramp.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message2,
							Proofs:              [][32]uint8{[32]byte(hash1)},
						}

						raw := ccip_offramp.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.OfframpConfigPDA,
							config.OfframpReferenceAddressesPDA,
							config.OfframpEvmSourceChainPDA,
							rootPDA,
							config.CcipOfframpProgram,
							config.AllowedOfframpEvmPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.RMNRemoteProgram,
							config.RMNRemoteCursesPDA,
							config.RMNRemoteConfigPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.OfframpReceiverExternalExecPDA, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
						executionEvents, err := common.ParseMultipleEvents[ccip.EventExecutionStateChanged](tx.Meta.LogMessages, "ExecutionStateChanged", config.PrintEvents)
						require.NoError(t, err)

						require.Equal(t, 2, len(executionEvents))
						require.Equal(t, config.EvmChainSelector, executionEvents[0].SourceChainSelector)
						require.Equal(t, message2.Header.SequenceNumber, executionEvents[0].SequenceNumber)
						require.Equal(t, hex.EncodeToString(message2.Header.MessageId[:]), hex.EncodeToString(executionEvents[0].MessageID[:]))
						require.Equal(t, hex.EncodeToString(hash2[:]), hex.EncodeToString(executionEvents[0].MessageHash[:]))
						require.Equal(t, ccip_offramp.InProgress_MessageExecutionState, executionEvents[0].State)

						require.Equal(t, config.EvmChainSelector, executionEvents[1].SourceChainSelector)
						require.Equal(t, message2.Header.SequenceNumber, executionEvents[1].SequenceNumber)
						require.Equal(t, hex.EncodeToString(message2.Header.MessageId[:]), hex.EncodeToString(executionEvents[1].MessageID[:]))
						require.Equal(t, hex.EncodeToString(hash2[:]), hex.EncodeToString(executionEvents[1].MessageHash[:]))
						require.Equal(t, ccip_offramp.Success_MessageExecutionState, executionEvents[1].State)

						var rootAccount ccip_offramp.CommitReport
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
						require.NoError(t, err, "failed to get account info")
						require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
						require.Equal(t, bin.Uint128{Lo: 10, Hi: 0}, rootAccount.ExecutionStates)
						require.Equal(t, commitReport.MerkleRoot.MinSeqNr, rootAccount.MinMsgNr)
						require.Equal(t, commitReport.MerkleRoot.MaxSeqNr, rootAccount.MaxMsgNr)
					})
				})
			})

			t.Run("uninitialized token account can be manually executed", func(t *testing.T) {
				// create new token receiver + find address (does not actually create account, just instruction)
				receiver := solana.MustPrivateKeyFromBase58("pYvqPnMDcx3hyE7jhSCAJLtnUeHzXp3aBm4yZ59mbz2Jw2ozW7BmBkLrMDBox17hn2mDsfHrNdR3PdvhGxaH9cB")
				ixATA, ata, err := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint, receiver.PublicKey(), legacyAdmin.PublicKey())
				require.NoError(t, err)
				token0.User[receiver.PublicKey()] = ata

				// create commit report ---------------------
				transmitter := getTransmitter()
				sourceChainSelector := config.EvmChainSelector
				msgAccounts := []solana.PublicKey{}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.TokenAmounts = []ccip_offramp.Any2SVMTokenTransfer{{
					SourcePoolAddress: []byte{1, 2, 3},
					DestTokenAddress:  token0.Mint,
					// 1 token * 1e9 due to decimal differences (18 in EVM, 9 in SVM). This will result in 1 unit in SVM.
					Amount: ccip_offramp.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1000000000)},
				}}
				message.TokenReceiver = receiver.PublicKey()
				rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				root := [32]byte(rootBytes)
				sequenceNumber := message.Header.SequenceNumber
				commitReport := ccip_offramp.CommitInput{
					MerkleRoot: &ccip_offramp.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindOfframpCommitReportPDA(config.EvmChainSelector, root, config.CcipOfframpProgram)
				require.NoError(t, err)
				execIx, err := ccip_offramp.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.OfframpBillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqAllowedPriceUpdaterOfframpPDA,
					config.FqConfigPDA,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{execIx}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := common.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEventCommitReportAccepted(tx.Meta.LogMessages, "CommitReportAccepted", &event))

				// try to execute report ----------------------
				// should fail because token account does not exist
				executionReport := ccip_offramp.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					OffchainTokenData:   [][]byte{{}},
					Proofs:              [][32]uint8{},
				}
				raw := ccip_offramp.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{0}, // only token transfer message
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[receiver.PublicKey()])
				require.NoError(t, err)
				raw.AccountMetaSlice = append(raw.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
				raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)
				// maps.Copy(addressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

				execIx, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{execIx}, transmitter, config.DefaultCommitment, addressTables, []string{"Error Code: " + common.AccountNotInitialized_AnchorError.String()})

				// create associated token account for user --------------------
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixATA}, legacyAdmin, config.DefaultCommitment)
				_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[receiver.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 0, initBal)

				// manual re-execution is successful -----------------------------------
				// NOTE: expects re-execution time to be instantaneous
				rawManual := ccip_offramp.NewManuallyExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					[]byte{0}, // only token transfer message
					config.OfframpConfigPDA,
					config.OfframpReferenceAddressesPDA,
					config.OfframpEvmSourceChainPDA,
					rootPDA,
					config.CcipOfframpProgram,
					config.AllowedOfframpEvmPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.RMNRemoteProgram,
					config.RMNRemoteCursesPDA,
					config.RMNRemoteConfigPDA,
				)

				manualTokenMetas, manualAddressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[receiver.PublicKey()])
				require.NoError(t, err)
				rawManual.AccountMetaSlice = append(rawManual.AccountMetaSlice, solana.Meta(token0.OfframpSigner))
				rawManual.AccountMetaSlice = append(rawManual.AccountMetaSlice, manualTokenMetas...)

				manualExecIx, err := rawManual.ValidateAndBuild()
				require.NoError(t, err)
				maps.Copy(manualAddressTables, offrampLookupTable) // commonly used ccip addresses - required otherwise tx is too large

				testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{manualExecIx}, legacyAdmin, config.DefaultCommitment, manualAddressTables, common.AddComputeUnitLimit(400_000))

				_, finalBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[receiver.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1, finalBal-initBal)
			})
		})
	})
}
